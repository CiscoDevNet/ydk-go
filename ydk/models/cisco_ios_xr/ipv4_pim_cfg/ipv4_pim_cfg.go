// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-pim package configuration.
// 
// This module contains definitions
// for the following management objects:
//   pim: PIM configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_pim_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_pim_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-pim-cfg pim}", reflect.TypeOf(Pim{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-pim-cfg:pim", reflect.TypeOf(Pim{}))
}

// PimProtocolMode represents Pim protocol mode
type PimProtocolMode string

const (
    // Sparse Mode
    PimProtocolMode_sm PimProtocolMode = "sm"

    // Bidirectional
    PimProtocolMode_bidir PimProtocolMode = "bidir"
)

// PimMultipath represents Pim multipath
type PimMultipath string

const (
    // Enable PIM multipath
    PimMultipath_enable PimMultipath = "enable"

    // Enable PIM multipath with interface based
    // hashing
    PimMultipath_interface_hash PimMultipath = "interface-hash"

    // Enable PIM multipath with source based hashing
    PimMultipath_source_hash PimMultipath = "source-hash"

    // Enable PIM multipath with source next-hop
    // hashing
    PimMultipath_source_next_hop_hash PimMultipath = "source-next-hop-hash"

    // Enable PIM multipath with source, group based
    // hashing
    PimMultipath_source_group_hash PimMultipath = "source-group-hash"
)

// Pim
// PIM configuration
// This type is a presence type.
type Pim struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs Pim_Vrfs

    // Default Context.
    DefaultContext Pim_DefaultContext
}

func (pim *Pim) GetFilter() yfilter.YFilter { return pim.YFilter }

func (pim *Pim) SetFilter(yf yfilter.YFilter) { pim.YFilter = yf }

func (pim *Pim) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "default-context" { return "DefaultContext" }
    return ""
}

func (pim *Pim) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-pim-cfg:pim"
}

func (pim *Pim) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &pim.Vrfs
    }
    if childYangName == "default-context" {
        return &pim.DefaultContext
    }
    return nil
}

func (pim *Pim) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &pim.Vrfs
    children["default-context"] = &pim.DefaultContext
    return children
}

func (pim *Pim) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pim *Pim) GetBundleName() string { return "cisco_ios_xr" }

func (pim *Pim) GetYangName() string { return "pim" }

func (pim *Pim) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pim *Pim) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pim *Pim) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pim *Pim) SetParent(parent types.Entity) { pim.parent = parent }

func (pim *Pim) GetParent() types.Entity { return pim.parent }

func (pim *Pim) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-pim-cfg" }

// Pim_Vrfs
// VRF table
type Pim_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Pim_Vrfs_Vrf.
    Vrf []Pim_Vrfs_Vrf
}

func (vrfs *Pim_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Pim_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Pim_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Pim_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Pim_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Pim_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Pim_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Pim_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Pim_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Pim_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Pim_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Pim_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Pim_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Pim_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Pim_Vrfs) GetParentYangName() string { return "pim" }

// Pim_Vrfs_Vrf
// VRF name
type Pim_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPV4 commands.
    Ipv4 Pim_Vrfs_Vrf_Ipv4

    // IPV6 commands.
    Ipv6 Pim_Vrfs_Vrf_Ipv6
}

func (vrf *Pim_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Pim_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Pim_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (vrf *Pim_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Pim_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    if childYangName == "ipv6" {
        return &vrf.Ipv6
    }
    return nil
}

func (vrf *Pim_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &vrf.Ipv4
    children["ipv6"] = &vrf.Ipv6
    return children
}

func (vrf *Pim_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Pim_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Pim_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Pim_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Pim_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Pim_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Pim_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Pim_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Pim_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Pim_Vrfs_Vrf_Ipv4
// IPV4 commands
type Pim_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Disable Rendezvous Point discovery through the AutoRP protocol. The type is
    // interface{}.
    AutoRpDisable interface{}

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_Vrfs_Vrf_Ipv4_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_Vrfs_Vrf_Ipv4_Rpf

    // Configure PIM State Limits.
    Maximum Pim_Vrfs_Vrf_Ipv4_Maximum

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable

    // Configure IP Multicast SSM.
    Ssm Pim_Vrfs_Vrf_Ipv4_Ssm

    // Inject Explicit PIM RPF Vector Proxy's.
    Injects Pim_Vrfs_Vrf_Ipv4_Injects

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_Vrfs_Vrf_Ipv4_Bsr

    // Multicast Only Fast Re-Route.
    Mofrr Pim_Vrfs_Vrf_Ipv4_Mofrr

    // Inject PIM RPF Vector Proxy's.
    Paths Pim_Vrfs_Vrf_Ipv4_Paths

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_Vrfs_Vrf_Ipv4_AllowRp

    // Configure convergence parameters.
    Convergence Pim_Vrfs_Vrf_Ipv4_Convergence

    // Interface-level Configuration.
    Interfaces Pim_Vrfs_Vrf_Ipv4_Interfaces
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "neighbor-check-on-receive" { return "NeighborCheckOnReceive" }
    if yname == "old-register-checksum" { return "OldRegisterChecksum" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "spt-threshold-infinity" { return "SptThresholdInfinity" }
    if yname == "log-neighbor-changes" { return "LogNeighborChanges" }
    if yname == "register-source" { return "RegisterSource" }
    if yname == "accept-register" { return "AcceptRegister" }
    if yname == "suppress-rpf-prunes" { return "SuppressRpfPrunes" }
    if yname == "ssm-allow-override" { return "SsmAllowOverride" }
    if yname == "multipath" { return "Multipath" }
    if yname == "rp-static-deny" { return "RpStaticDeny" }
    if yname == "suppress-data-registers" { return "SuppressDataRegisters" }
    if yname == "neighbor-check-on-send" { return "NeighborCheckOnSend" }
    if yname == "auto-rp-disable" { return "AutoRpDisable" }
    if yname == "sparse-mode-rp-addresses" { return "SparseModeRpAddresses" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "rpf" { return "Rpf" }
    if yname == "maximum" { return "Maximum" }
    if yname == "sg-expiry-timer" { return "SgExpiryTimer" }
    if yname == "rpf-vector-enable" { return "RpfVectorEnable" }
    if yname == "ssm" { return "Ssm" }
    if yname == "injects" { return "Injects" }
    if yname == "bidir-rp-addresses" { return "BidirRpAddresses" }
    if yname == "bsr" { return "Bsr" }
    if yname == "mofrr" { return "Mofrr" }
    if yname == "paths" { return "Paths" }
    if yname == "allow-rp" { return "AllowRp" }
    if yname == "convergence" { return "Convergence" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-addresses" {
        return &ipv4.SparseModeRpAddresses
    }
    if childYangName == "inheritable-defaults" {
        return &ipv4.InheritableDefaults
    }
    if childYangName == "rpf" {
        return &ipv4.Rpf
    }
    if childYangName == "maximum" {
        return &ipv4.Maximum
    }
    if childYangName == "sg-expiry-timer" {
        return &ipv4.SgExpiryTimer
    }
    if childYangName == "rpf-vector-enable" {
        return &ipv4.RpfVectorEnable
    }
    if childYangName == "ssm" {
        return &ipv4.Ssm
    }
    if childYangName == "injects" {
        return &ipv4.Injects
    }
    if childYangName == "bidir-rp-addresses" {
        return &ipv4.BidirRpAddresses
    }
    if childYangName == "bsr" {
        return &ipv4.Bsr
    }
    if childYangName == "mofrr" {
        return &ipv4.Mofrr
    }
    if childYangName == "paths" {
        return &ipv4.Paths
    }
    if childYangName == "allow-rp" {
        return &ipv4.AllowRp
    }
    if childYangName == "convergence" {
        return &ipv4.Convergence
    }
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    return nil
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sparse-mode-rp-addresses"] = &ipv4.SparseModeRpAddresses
    children["inheritable-defaults"] = &ipv4.InheritableDefaults
    children["rpf"] = &ipv4.Rpf
    children["maximum"] = &ipv4.Maximum
    children["sg-expiry-timer"] = &ipv4.SgExpiryTimer
    children["rpf-vector-enable"] = &ipv4.RpfVectorEnable
    children["ssm"] = &ipv4.Ssm
    children["injects"] = &ipv4.Injects
    children["bidir-rp-addresses"] = &ipv4.BidirRpAddresses
    children["bsr"] = &ipv4.Bsr
    children["mofrr"] = &ipv4.Mofrr
    children["paths"] = &ipv4.Paths
    children["allow-rp"] = &ipv4.AllowRp
    children["convergence"] = &ipv4.Convergence
    children["interfaces"] = &ipv4.Interfaces
    return children
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-check-on-receive"] = ipv4.NeighborCheckOnReceive
    leafs["old-register-checksum"] = ipv4.OldRegisterChecksum
    leafs["neighbor-filter"] = ipv4.NeighborFilter
    leafs["spt-threshold-infinity"] = ipv4.SptThresholdInfinity
    leafs["log-neighbor-changes"] = ipv4.LogNeighborChanges
    leafs["register-source"] = ipv4.RegisterSource
    leafs["accept-register"] = ipv4.AcceptRegister
    leafs["suppress-rpf-prunes"] = ipv4.SuppressRpfPrunes
    leafs["ssm-allow-override"] = ipv4.SsmAllowOverride
    leafs["multipath"] = ipv4.Multipath
    leafs["rp-static-deny"] = ipv4.RpStaticDeny
    leafs["suppress-data-registers"] = ipv4.SuppressDataRegisters
    leafs["neighbor-check-on-send"] = ipv4.NeighborCheckOnSend
    leafs["auto-rp-disable"] = ipv4.AutoRpDisable
    return leafs
}

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Pim_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

// Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetFilter() yfilter.YFilter { return sparseModeRpAddresses.YFilter }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) SetFilter(yf yfilter.YFilter) { sparseModeRpAddresses.YFilter = yf }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetGoName(yname string) string {
    if yname == "sparse-mode-rp-address" { return "SparseModeRpAddress" }
    return ""
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetSegmentPath() string {
    return "sparse-mode-rp-addresses"
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-address" {
        for _, c := range sparseModeRpAddresses.SparseModeRpAddress {
            if sparseModeRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress{}
        sparseModeRpAddresses.SparseModeRpAddress = append(sparseModeRpAddresses.SparseModeRpAddress, child)
        return &sparseModeRpAddresses.SparseModeRpAddress[len(sparseModeRpAddresses.SparseModeRpAddress)-1]
    }
    return nil
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        children[sparseModeRpAddresses.SparseModeRpAddress[i].GetSegmentPath()] = &sparseModeRpAddresses.SparseModeRpAddress[i]
    }
    return children
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetYangName() string { return "sparse-mode-rp-addresses" }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) SetParent(parent types.Entity) { sparseModeRpAddresses.parent = parent }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetParent() types.Entity { return sparseModeRpAddresses.parent }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetFilter() yfilter.YFilter { return sparseModeRpAddress.YFilter }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) SetFilter(yf yfilter.YFilter) { sparseModeRpAddress.YFilter = yf }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetSegmentPath() string {
    return "sparse-mode-rp-address" + "[rp-address='" + fmt.Sprintf("%v", sparseModeRpAddress.RpAddress) + "']"
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = sparseModeRpAddress.RpAddress
    leafs["access-list-name"] = sparseModeRpAddress.AccessListName
    leafs["auto-rp-override"] = sparseModeRpAddress.AutoRpOverride
    return leafs
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetYangName() string { return "sparse-mode-rp-address" }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) SetParent(parent types.Entity) { sparseModeRpAddress.parent = parent }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetParent() types.Entity { return sparseModeRpAddress.parent }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetParentYangName() string { return "sparse-mode-rp-addresses" }

// Pim_Vrfs_Vrf_Ipv4_InheritableDefaults
// Inheritable defaults
type Pim_Vrfs_Vrf_Ipv4_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetGoName(yname string) string {
    if yname == "convergence-timeout" { return "ConvergenceTimeout" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    return ""
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["convergence-timeout"] = inheritableDefaults.ConvergenceTimeout
    leafs["hello-interval"] = inheritableDefaults.HelloInterval
    leafs["propagation-delay"] = inheritableDefaults.PropagationDelay
    leafs["dr-priority"] = inheritableDefaults.DrPriority
    leafs["join-prune-mtu"] = inheritableDefaults.JoinPruneMtu
    leafs["jp-interval"] = inheritableDefaults.JpInterval
    leafs["override-interval"] = inheritableDefaults.OverrideInterval
    return leafs
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv4_InheritableDefaults) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Rpf
// Configure RPF options
type Pim_Vrfs_Vrf_Ipv4_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy"] = rpf.RoutePolicy
    return leafs
}

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetYangName() string { return "rpf" }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Pim_Vrfs_Vrf_Ipv4_Rpf) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Maximum
// Configure PIM State Limits
type Pim_Vrfs_Vrf_Ipv4_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_Vrfs_Vrf_Ipv4_Maximum_Routes
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetGoName(yname string) string {
    if yname == "group-mappings-auto-rp" { return "GroupMappingsAutoRp" }
    if yname == "bsr-group-mappings" { return "BsrGroupMappings" }
    if yname == "register-states" { return "RegisterStates" }
    if yname == "route-interfaces" { return "RouteInterfaces" }
    if yname == "bsr-candidate-rp-cache" { return "BsrCandidateRpCache" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-mappings-auto-rp" {
        return &maximum.GroupMappingsAutoRp
    }
    if childYangName == "bsr-group-mappings" {
        return &maximum.BsrGroupMappings
    }
    if childYangName == "register-states" {
        return &maximum.RegisterStates
    }
    if childYangName == "route-interfaces" {
        return &maximum.RouteInterfaces
    }
    if childYangName == "bsr-candidate-rp-cache" {
        return &maximum.BsrCandidateRpCache
    }
    if childYangName == "routes" {
        return &maximum.Routes
    }
    return nil
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-mappings-auto-rp"] = &maximum.GroupMappingsAutoRp
    children["bsr-group-mappings"] = &maximum.BsrGroupMappings
    children["register-states"] = &maximum.RegisterStates
    children["route-interfaces"] = &maximum.RouteInterfaces
    children["bsr-candidate-rp-cache"] = &maximum.BsrCandidateRpCache
    children["routes"] = &maximum.Routes
    return children
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetYangName() string { return "maximum" }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Pim_Vrfs_Vrf_Ipv4_Maximum) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetFilter() yfilter.YFilter { return groupMappingsAutoRp.YFilter }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { groupMappingsAutoRp.YFilter = yf }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-group-ranges-auto-rp" { return "MaximumGroupRangesAutoRp" }
    if yname == "threshold-group-ranges-auto-rp" { return "ThresholdGroupRangesAutoRp" }
    return ""
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetSegmentPath() string {
    return "group-mappings-auto-rp"
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-group-ranges-auto-rp"] = groupMappingsAutoRp.MaximumGroupRangesAutoRp
    leafs["threshold-group-ranges-auto-rp"] = groupMappingsAutoRp.ThresholdGroupRangesAutoRp
    return leafs
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetYangName() string { return "group-mappings-auto-rp" }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) SetParent(parent types.Entity) { groupMappingsAutoRp.parent = parent }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetParent() types.Entity { return groupMappingsAutoRp.parent }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv4_Maximum_GroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetFilter() yfilter.YFilter { return bsrGroupMappings.YFilter }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGroupMappings.YFilter = yf }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-group-ranges" { return "BsrMaximumGroupRanges" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetSegmentPath() string {
    return "bsr-group-mappings"
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-group-ranges"] = bsrGroupMappings.BsrMaximumGroupRanges
    leafs["warning-threshold"] = bsrGroupMappings.WarningThreshold
    return leafs
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetYangName() string { return "bsr-group-mappings" }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) SetParent(parent types.Entity) { bsrGroupMappings.parent = parent }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetParent() types.Entity { return bsrGroupMappings.parent }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetFilter() yfilter.YFilter { return registerStates.YFilter }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) SetFilter(yf yfilter.YFilter) { registerStates.YFilter = yf }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetSegmentPath() string {
    return "register-states"
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = registerStates.MaximumRegisterStates
    leafs["warning-threshold"] = registerStates.WarningThreshold
    return leafs
}

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetYangName() string { return "register-states" }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) SetParent(parent types.Entity) { registerStates.parent = parent }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetParent() types.Entity { return registerStates.parent }

func (registerStates *Pim_Vrfs_Vrf_Ipv4_Maximum_RegisterStates) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetFilter() yfilter.YFilter { return routeInterfaces.YFilter }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) SetFilter(yf yfilter.YFilter) { routeInterfaces.YFilter = yf }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetSegmentPath() string {
    return "route-interfaces"
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = routeInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = routeInterfaces.WarningThreshold
    return leafs
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetYangName() string { return "route-interfaces" }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) SetParent(parent types.Entity) { routeInterfaces.parent = parent }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetParent() types.Entity { return routeInterfaces.parent }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv4_Maximum_RouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetFilter() yfilter.YFilter { return bsrCandidateRpCache.YFilter }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrCandidateRpCache.YFilter = yf }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-candidate-rp-cache" { return "BsrMaximumCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetSegmentPath() string {
    return "bsr-candidate-rp-cache"
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-candidate-rp-cache"] = bsrCandidateRpCache.BsrMaximumCandidateRpCache
    leafs["warning-threshold"] = bsrCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetYangName() string { return "bsr-candidate-rp-cache" }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) SetParent(parent types.Entity) { bsrCandidateRpCache.parent = parent }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetParent() types.Entity { return bsrCandidateRpCache.parent }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv4_Maximum_BsrCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Maximum_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = routes.MaximumRoutes
    leafs["warning-threshold"] = routes.WarningThreshold
    return leafs
}

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetYangName() string { return "routes" }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Pim_Vrfs_Vrf_Ipv4_Maximum_Routes) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetFilter() yfilter.YFilter { return sgExpiryTimer.YFilter }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) SetFilter(yf yfilter.YFilter) { sgExpiryTimer.YFilter = yf }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetSegmentPath() string {
    return "sg-expiry-timer"
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = sgExpiryTimer.Interval
    leafs["access-list-name"] = sgExpiryTimer.AccessListName
    return leafs
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetBundleName() string { return "cisco_ios_xr" }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetYangName() string { return "sg-expiry-timer" }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) SetParent(parent types.Entity) { sgExpiryTimer.parent = parent }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetParent() types.Entity { return sgExpiryTimer.parent }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv4_SgExpiryTimer) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetFilter() yfilter.YFilter { return rpfVectorEnable.YFilter }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) SetFilter(yf yfilter.YFilter) { rpfVectorEnable.YFilter = yf }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-ebgp" { return "AllowEbgp" }
    if yname == "disable-ibgp" { return "DisableIbgp" }
    return ""
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetSegmentPath() string {
    return "rpf-vector-enable"
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpfVectorEnable.Enable
    leafs["allow-ebgp"] = rpfVectorEnable.AllowEbgp
    leafs["disable-ibgp"] = rpfVectorEnable.DisableIbgp
    return leafs
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetBundleName() string { return "cisco_ios_xr" }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetYangName() string { return "rpf-vector-enable" }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) SetParent(parent types.Entity) { rpfVectorEnable.parent = parent }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetParent() types.Entity { return rpfVectorEnable.parent }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv4_RpfVectorEnable) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Ssm
// Configure IP Multicast SSM
type Pim_Vrfs_Vrf_Ipv4_Ssm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetFilter() yfilter.YFilter { return ssm.YFilter }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) SetFilter(yf yfilter.YFilter) { ssm.YFilter = yf }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "range" { return "Range" }
    return ""
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetSegmentPath() string {
    return "ssm"
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = ssm.Disable
    leafs["range"] = ssm.Range
    return leafs
}

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetBundleName() string { return "cisco_ios_xr" }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetYangName() string { return "ssm" }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) SetParent(parent types.Entity) { ssm.parent = parent }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetParent() types.Entity { return ssm.parent }

func (ssm *Pim_Vrfs_Vrf_Ipv4_Ssm) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Injects
// Inject Explicit PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Injects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inject Explicit PIM RPF Vector Proxy's. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Injects_Inject.
    Inject []Pim_Vrfs_Vrf_Ipv4_Injects_Inject
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetFilter() yfilter.YFilter { return injects.YFilter }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) SetFilter(yf yfilter.YFilter) { injects.YFilter = yf }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetGoName(yname string) string {
    if yname == "inject" { return "Inject" }
    return ""
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetSegmentPath() string {
    return "injects"
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inject" {
        for _, c := range injects.Inject {
            if injects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Injects_Inject{}
        injects.Inject = append(injects.Inject, child)
        return &injects.Inject[len(injects.Inject)-1]
    }
    return nil
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range injects.Inject {
        children[injects.Inject[i].GetSegmentPath()] = &injects.Inject[i]
    }
    return children
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetBundleName() string { return "cisco_ios_xr" }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetYangName() string { return "injects" }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) SetParent(parent types.Entity) { injects.parent = parent }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetParent() types.Entity { return injects.parent }

func (injects *Pim_Vrfs_Vrf_Ipv4_Injects) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Injects_Inject
// Inject Explicit PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Injects_Inject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetFilter() yfilter.YFilter { return inject.YFilter }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) SetFilter(yf yfilter.YFilter) { inject.YFilter = yf }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "rpf-proxy-address" { return "RpfProxyAddress" }
    return ""
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetSegmentPath() string {
    return "inject" + "[source-address='" + fmt.Sprintf("%v", inject.SourceAddress) + "']" + "[prefix-length='" + fmt.Sprintf("%v", inject.PrefixLength) + "']"
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = inject.SourceAddress
    leafs["prefix-length"] = inject.PrefixLength
    leafs["rpf-proxy-address"] = inject.RpfProxyAddress
    return leafs
}

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetBundleName() string { return "cisco_ios_xr" }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetYangName() string { return "inject" }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) SetParent(parent types.Entity) { inject.parent = parent }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetParent() types.Entity { return inject.parent }

func (inject *Pim_Vrfs_Vrf_Ipv4_Injects_Inject) GetParentYangName() string { return "injects" }

// Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetFilter() yfilter.YFilter { return bidirRpAddresses.YFilter }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) SetFilter(yf yfilter.YFilter) { bidirRpAddresses.YFilter = yf }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetGoName(yname string) string {
    if yname == "bidir-rp-address" { return "BidirRpAddress" }
    return ""
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetSegmentPath() string {
    return "bidir-rp-addresses"
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bidir-rp-address" {
        for _, c := range bidirRpAddresses.BidirRpAddress {
            if bidirRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress{}
        bidirRpAddresses.BidirRpAddress = append(bidirRpAddresses.BidirRpAddress, child)
        return &bidirRpAddresses.BidirRpAddress[len(bidirRpAddresses.BidirRpAddress)-1]
    }
    return nil
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bidirRpAddresses.BidirRpAddress {
        children[bidirRpAddresses.BidirRpAddress[i].GetSegmentPath()] = &bidirRpAddresses.BidirRpAddress[i]
    }
    return children
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetYangName() string { return "bidir-rp-addresses" }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) SetParent(parent types.Entity) { bidirRpAddresses.parent = parent }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetParent() types.Entity { return bidirRpAddresses.parent }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetFilter() yfilter.YFilter { return bidirRpAddress.YFilter }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) SetFilter(yf yfilter.YFilter) { bidirRpAddress.YFilter = yf }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetSegmentPath() string {
    return "bidir-rp-address" + "[rp-address='" + fmt.Sprintf("%v", bidirRpAddress.RpAddress) + "']"
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = bidirRpAddress.RpAddress
    leafs["access-list-name"] = bidirRpAddress.AccessListName
    leafs["auto-rp-override"] = bidirRpAddress.AutoRpOverride
    return leafs
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetYangName() string { return "bidir-rp-address" }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) SetParent(parent types.Entity) { bidirRpAddress.parent = parent }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetParent() types.Entity { return bidirRpAddress.parent }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv4_BidirRpAddresses_BidirRpAddress) GetParentYangName() string { return "bidir-rp-addresses" }

// Pim_Vrfs_Vrf_Ipv4_Bsr
// PIM BSR configuration
type Pim_Vrfs_Vrf_Ipv4_Bsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetFilter() yfilter.YFilter { return bsr.YFilter }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) SetFilter(yf yfilter.YFilter) { bsr.YFilter = yf }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetGoName(yname string) string {
    if yname == "candidate-bsr" { return "CandidateBsr" }
    if yname == "candidate-rps" { return "CandidateRps" }
    return ""
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetSegmentPath() string {
    return "bsr"
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-bsr" {
        return &bsr.CandidateBsr
    }
    if childYangName == "candidate-rps" {
        return &bsr.CandidateRps
    }
    return nil
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-bsr"] = &bsr.CandidateBsr
    children["candidate-rps"] = &bsr.CandidateRps
    return children
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetBundleName() string { return "cisco_ios_xr" }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetYangName() string { return "bsr" }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) SetParent(parent types.Entity) { bsr.parent = parent }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetParent() types.Entity { return bsr.parent }

func (bsr *Pim_Vrfs_Vrf_Ipv4_Bsr) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BSR Address configured. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..32. The default value is 30.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetFilter() yfilter.YFilter { return candidateBsr.YFilter }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) SetFilter(yf yfilter.YFilter) { candidateBsr.YFilter = yf }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetSegmentPath() string {
    return "candidate-bsr"
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateBsr.Address
    leafs["prefix-length"] = candidateBsr.PrefixLength
    leafs["priority"] = candidateBsr.Priority
    return leafs
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetBundleName() string { return "cisco_ios_xr" }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetYangName() string { return "candidate-bsr" }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) SetParent(parent types.Entity) { candidateBsr.parent = parent }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetParent() types.Entity { return candidateBsr.parent }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateBsr) GetParentYangName() string { return "bsr" }

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps
// PIM RP configuration
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp.
    CandidateRp []Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetFilter() yfilter.YFilter { return candidateRps.YFilter }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) SetFilter(yf yfilter.YFilter) { candidateRps.YFilter = yf }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    return ""
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetSegmentPath() string {
    return "candidate-rps"
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        for _, c := range candidateRps.CandidateRp {
            if candidateRps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp{}
        candidateRps.CandidateRp = append(candidateRps.CandidateRp, child)
        return &candidateRps.CandidateRp[len(candidateRps.CandidateRp)-1]
    }
    return nil
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateRps.CandidateRp {
        children[candidateRps.CandidateRp[i].GetSegmentPath()] = &candidateRps.CandidateRp[i]
    }
    return children
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetYangName() string { return "candidate-rps" }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) SetParent(parent types.Entity) { candidateRps.parent = parent }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetParent() types.Entity { return candidateRps.parent }

func (candidateRps *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps) GetParentYangName() string { return "bsr" }

// Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mode" { return "Mode" }
    if yname == "group-list" { return "GroupList" }
    if yname == "priority" { return "Priority" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetSegmentPath() string {
    return "candidate-rp" + "[address='" + fmt.Sprintf("%v", candidateRp.Address) + "']" + "[mode='" + fmt.Sprintf("%v", candidateRp.Mode) + "']"
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateRp.Address
    leafs["mode"] = candidateRp.Mode
    leafs["group-list"] = candidateRp.GroupList
    leafs["priority"] = candidateRp.Priority
    leafs["interval"] = candidateRp.Interval
    return leafs
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *Pim_Vrfs_Vrf_Ipv4_Bsr_CandidateRps_CandidateRp) GetParentYangName() string { return "candidate-rps" }

// Pim_Vrfs_Vrf_Ipv4_Mofrr
// Multicast Only Fast Re-Route
type Pim_Vrfs_Vrf_Ipv4_Mofrr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifying SG that should do RIB MOFRR. The type is string with
    // length: 1..64.
    Rib interface{}

    // Non-revertive Multicast Only Fast Re-Route. The type is interface{}.
    NonRevertive interface{}

    // Enable Multicast Only FRR. The type is interface{}.
    Enable interface{}

    // Access-list specifying SG that should do FLOW MOFRR. The type is string
    // with length: 1..64.
    Flow interface{}

    // Clone multicast joins.
    CloneJoins Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins

    // Clone multicast traffic.
    CloneSources Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetFilter() yfilter.YFilter { return mofrr.YFilter }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) SetFilter(yf yfilter.YFilter) { mofrr.YFilter = yf }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetGoName(yname string) string {
    if yname == "rib" { return "Rib" }
    if yname == "non-revertive" { return "NonRevertive" }
    if yname == "enable" { return "Enable" }
    if yname == "flow" { return "Flow" }
    if yname == "clone-joins" { return "CloneJoins" }
    if yname == "clone-sources" { return "CloneSources" }
    return ""
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetSegmentPath() string {
    return "mofrr"
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-joins" {
        return &mofrr.CloneJoins
    }
    if childYangName == "clone-sources" {
        return &mofrr.CloneSources
    }
    return nil
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clone-joins"] = &mofrr.CloneJoins
    children["clone-sources"] = &mofrr.CloneSources
    return children
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rib"] = mofrr.Rib
    leafs["non-revertive"] = mofrr.NonRevertive
    leafs["enable"] = mofrr.Enable
    leafs["flow"] = mofrr.Flow
    return leafs
}

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetBundleName() string { return "cisco_ios_xr" }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetYangName() string { return "mofrr" }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) SetParent(parent types.Entity) { mofrr.parent = parent }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetParent() types.Entity { return mofrr.parent }

func (mofrr *Pim_Vrfs_Vrf_Ipv4_Mofrr) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins
// Clone multicast joins
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clone S,G joins as S1,G joins and S2,G joins. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin.
    CloneJoin []Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetFilter() yfilter.YFilter { return cloneJoins.YFilter }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) SetFilter(yf yfilter.YFilter) { cloneJoins.YFilter = yf }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetGoName(yname string) string {
    if yname == "clone-join" { return "CloneJoin" }
    return ""
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetSegmentPath() string {
    return "clone-joins"
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-join" {
        for _, c := range cloneJoins.CloneJoin {
            if cloneJoins.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin{}
        cloneJoins.CloneJoin = append(cloneJoins.CloneJoin, child)
        return &cloneJoins.CloneJoin[len(cloneJoins.CloneJoin)-1]
    }
    return nil
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cloneJoins.CloneJoin {
        children[cloneJoins.CloneJoin[i].GetSegmentPath()] = &cloneJoins.CloneJoin[i]
    }
    return children
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetBundleName() string { return "cisco_ios_xr" }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetYangName() string { return "clone-joins" }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) SetParent(parent types.Entity) { cloneJoins.parent = parent }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetParent() types.Entity { return cloneJoins.parent }

func (cloneJoins *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins) GetParentYangName() string { return "mofrr" }

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin
// Clone S,G joins as S1,G joins and S2,G joins
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetFilter() yfilter.YFilter { return cloneJoin.YFilter }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) SetFilter(yf yfilter.YFilter) { cloneJoin.YFilter = yf }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    if yname == "primary" { return "Primary" }
    if yname == "backup" { return "Backup" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetSegmentPath() string {
    return "clone-join" + "[source='" + fmt.Sprintf("%v", cloneJoin.Source) + "']" + "[primary='" + fmt.Sprintf("%v", cloneJoin.Primary) + "']" + "[backup='" + fmt.Sprintf("%v", cloneJoin.Backup) + "']" + "[prefix-length='" + fmt.Sprintf("%v", cloneJoin.PrefixLength) + "']"
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = cloneJoin.Source
    leafs["primary"] = cloneJoin.Primary
    leafs["backup"] = cloneJoin.Backup
    leafs["prefix-length"] = cloneJoin.PrefixLength
    return leafs
}

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetBundleName() string { return "cisco_ios_xr" }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetYangName() string { return "clone-join" }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) SetParent(parent types.Entity) { cloneJoin.parent = parent }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetParent() types.Entity { return cloneJoin.parent }

func (cloneJoin *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneJoins_CloneJoin) GetParentYangName() string { return "clone-joins" }

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources
// Clone multicast traffic
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clone S,G traffic as S1,G traffic and S2,G traffic. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource.
    CloneSource []Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetFilter() yfilter.YFilter { return cloneSources.YFilter }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) SetFilter(yf yfilter.YFilter) { cloneSources.YFilter = yf }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetGoName(yname string) string {
    if yname == "clone-source" { return "CloneSource" }
    return ""
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetSegmentPath() string {
    return "clone-sources"
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-source" {
        for _, c := range cloneSources.CloneSource {
            if cloneSources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource{}
        cloneSources.CloneSource = append(cloneSources.CloneSource, child)
        return &cloneSources.CloneSource[len(cloneSources.CloneSource)-1]
    }
    return nil
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cloneSources.CloneSource {
        children[cloneSources.CloneSource[i].GetSegmentPath()] = &cloneSources.CloneSource[i]
    }
    return children
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetBundleName() string { return "cisco_ios_xr" }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetYangName() string { return "clone-sources" }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) SetParent(parent types.Entity) { cloneSources.parent = parent }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetParent() types.Entity { return cloneSources.parent }

func (cloneSources *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources) GetParentYangName() string { return "mofrr" }

// Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource
// Clone S,G traffic as S1,G traffic and S2,G
// traffic
type Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetFilter() yfilter.YFilter { return cloneSource.YFilter }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) SetFilter(yf yfilter.YFilter) { cloneSource.YFilter = yf }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    if yname == "primary" { return "Primary" }
    if yname == "backup" { return "Backup" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetSegmentPath() string {
    return "clone-source" + "[source='" + fmt.Sprintf("%v", cloneSource.Source) + "']" + "[primary='" + fmt.Sprintf("%v", cloneSource.Primary) + "']" + "[backup='" + fmt.Sprintf("%v", cloneSource.Backup) + "']" + "[prefix-length='" + fmt.Sprintf("%v", cloneSource.PrefixLength) + "']"
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = cloneSource.Source
    leafs["primary"] = cloneSource.Primary
    leafs["backup"] = cloneSource.Backup
    leafs["prefix-length"] = cloneSource.PrefixLength
    return leafs
}

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetBundleName() string { return "cisco_ios_xr" }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetYangName() string { return "clone-source" }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) SetParent(parent types.Entity) { cloneSource.parent = parent }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetParent() types.Entity { return cloneSource.parent }

func (cloneSource *Pim_Vrfs_Vrf_Ipv4_Mofrr_CloneSources_CloneSource) GetParentYangName() string { return "clone-sources" }

// Pim_Vrfs_Vrf_Ipv4_Paths
// Inject PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inject PIM RPF Vector Proxy's. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Paths_Path.
    Path []Pim_Vrfs_Vrf_Ipv4_Paths_Path
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetYangName() string { return "paths" }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Pim_Vrfs_Vrf_Ipv4_Paths) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Paths_Path
// Inject PIM RPF Vector Proxy's
type Pim_Vrfs_Vrf_Ipv4_Paths_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "rpf-proxy-address" { return "RpfProxyAddress" }
    return ""
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetSegmentPath() string {
    return "path" + "[source-address='" + fmt.Sprintf("%v", path.SourceAddress) + "']" + "[prefix-length='" + fmt.Sprintf("%v", path.PrefixLength) + "']"
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = path.SourceAddress
    leafs["prefix-length"] = path.PrefixLength
    leafs["rpf-proxy-address"] = path.RpfProxyAddress
    return leafs
}

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetYangName() string { return "path" }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *Pim_Vrfs_Vrf_Ipv4_Paths_Path) GetParentYangName() string { return "paths" }

// Pim_Vrfs_Vrf_Ipv4_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_AllowRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetFilter() yfilter.YFilter { return allowRp.YFilter }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) SetFilter(yf yfilter.YFilter) { allowRp.YFilter = yf }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetGoName(yname string) string {
    if yname == "rp-list-name" { return "RpListName" }
    if yname == "group-list-name" { return "GroupListName" }
    return ""
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetSegmentPath() string {
    return "allow-rp"
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-list-name"] = allowRp.RpListName
    leafs["group-list-name"] = allowRp.GroupListName
    return leafs
}

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetBundleName() string { return "cisco_ios_xr" }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetYangName() string { return "allow-rp" }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) SetParent(parent types.Entity) { allowRp.parent = parent }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetParent() types.Entity { return allowRp.parent }

func (allowRp *Pim_Vrfs_Vrf_Ipv4_AllowRp) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Convergence
// Configure convergence parameters
type Pim_Vrfs_Vrf_Ipv4_Convergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetFilter() yfilter.YFilter { return convergence.YFilter }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) SetFilter(yf yfilter.YFilter) { convergence.YFilter = yf }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetGoName(yname string) string {
    if yname == "rpf-conflict-join-delay" { return "RpfConflictJoinDelay" }
    if yname == "link-down-prune-delay" { return "LinkDownPruneDelay" }
    return ""
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetSegmentPath() string {
    return "convergence"
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rpf-conflict-join-delay"] = convergence.RpfConflictJoinDelay
    leafs["link-down-prune-delay"] = convergence.LinkDownPruneDelay
    return leafs
}

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetBundleName() string { return "cisco_ios_xr" }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetYangName() string { return "convergence" }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) SetParent(parent types.Entity) { convergence.parent = parent }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetParent() types.Entity { return convergence.parent }

func (convergence *Pim_Vrfs_Vrf_Ipv4_Convergence) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Interfaces
// Interface-level Configuration
type Pim_Vrfs_Vrf_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface.
    Interface []Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pim_Vrfs_Vrf_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface
// The name of the interface
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "bsr-border" { return "BsrBorder" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "interface-enable" { return "InterfaceEnable" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "bfd" { return "Bfd" }
    return ""
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-routes" {
        return &self.MaximumRoutes
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    return nil
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-routes"] = &self.MaximumRoutes
    children["bfd"] = &self.Bfd
    return children
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    leafs["neighbor-filter"] = self.NeighborFilter
    leafs["hello-interval"] = self.HelloInterval
    leafs["bsr-border"] = self.BsrBorder
    leafs["propagation-delay"] = self.PropagationDelay
    leafs["dr-priority"] = self.DrPriority
    leafs["join-prune-mtu"] = self.JoinPruneMtu
    leafs["interface-enable"] = self.InterfaceEnable
    leafs["jp-interval"] = self.JpInterval
    leafs["override-interval"] = self.OverrideInterval
    return leafs
}

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetFilter() yfilter.YFilter { return maximumRoutes.YFilter }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) SetFilter(yf yfilter.YFilter) { maximumRoutes.YFilter = yf }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetSegmentPath() string {
    return "maximum-routes"
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = maximumRoutes.Maximum
    leafs["warning-threshold"] = maximumRoutes.WarningThreshold
    leafs["access-list-name"] = maximumRoutes.AccessListName
    return leafs
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetYangName() string { return "maximum-routes" }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) SetParent(parent types.Entity) { maximumRoutes.parent = parent }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetParent() types.Entity { return maximumRoutes.parent }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_MaximumRoutes) GetParentYangName() string { return "interface" }

// Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd
// BFD configuration
type Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    if yname == "interval" { return "Interval" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    leafs["interval"] = bfd.Interval
    leafs["enable"] = bfd.Enable
    return leafs
}

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Pim_Vrfs_Vrf_Ipv4_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Pim_Vrfs_Vrf_Ipv6
// IPV6 commands
type Pim_Vrfs_Vrf_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Set Embedded RP processing support. The type is interface{}.
    EmbeddedRpDisable interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_Vrfs_Vrf_Ipv6_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_Vrfs_Vrf_Ipv6_Rpf

    // Configure PIM State Limits.
    Maximum Pim_Vrfs_Vrf_Ipv6_Maximum

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable

    // Configure IP Multicast SSM.
    Ssm Pim_Vrfs_Vrf_Ipv6_Ssm

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_Vrfs_Vrf_Ipv6_Bsr

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_Vrfs_Vrf_Ipv6_AllowRp

    // Set Embedded RP processing support.
    EmbeddedRpAddresses Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses

    // Configure convergence parameters.
    Convergence Pim_Vrfs_Vrf_Ipv6_Convergence

    // Interface-level Configuration.
    Interfaces Pim_Vrfs_Vrf_Ipv6_Interfaces
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetGoName(yname string) string {
    if yname == "neighbor-check-on-receive" { return "NeighborCheckOnReceive" }
    if yname == "old-register-checksum" { return "OldRegisterChecksum" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "spt-threshold-infinity" { return "SptThresholdInfinity" }
    if yname == "log-neighbor-changes" { return "LogNeighborChanges" }
    if yname == "register-source" { return "RegisterSource" }
    if yname == "accept-register" { return "AcceptRegister" }
    if yname == "embedded-rp-disable" { return "EmbeddedRpDisable" }
    if yname == "suppress-rpf-prunes" { return "SuppressRpfPrunes" }
    if yname == "ssm-allow-override" { return "SsmAllowOverride" }
    if yname == "multipath" { return "Multipath" }
    if yname == "rp-static-deny" { return "RpStaticDeny" }
    if yname == "suppress-data-registers" { return "SuppressDataRegisters" }
    if yname == "neighbor-check-on-send" { return "NeighborCheckOnSend" }
    if yname == "sparse-mode-rp-addresses" { return "SparseModeRpAddresses" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "rpf" { return "Rpf" }
    if yname == "maximum" { return "Maximum" }
    if yname == "sg-expiry-timer" { return "SgExpiryTimer" }
    if yname == "rpf-vector-enable" { return "RpfVectorEnable" }
    if yname == "ssm" { return "Ssm" }
    if yname == "bidir-rp-addresses" { return "BidirRpAddresses" }
    if yname == "bsr" { return "Bsr" }
    if yname == "allow-rp" { return "AllowRp" }
    if yname == "embedded-rp-addresses" { return "EmbeddedRpAddresses" }
    if yname == "convergence" { return "Convergence" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-addresses" {
        return &ipv6.SparseModeRpAddresses
    }
    if childYangName == "inheritable-defaults" {
        return &ipv6.InheritableDefaults
    }
    if childYangName == "rpf" {
        return &ipv6.Rpf
    }
    if childYangName == "maximum" {
        return &ipv6.Maximum
    }
    if childYangName == "sg-expiry-timer" {
        return &ipv6.SgExpiryTimer
    }
    if childYangName == "rpf-vector-enable" {
        return &ipv6.RpfVectorEnable
    }
    if childYangName == "ssm" {
        return &ipv6.Ssm
    }
    if childYangName == "bidir-rp-addresses" {
        return &ipv6.BidirRpAddresses
    }
    if childYangName == "bsr" {
        return &ipv6.Bsr
    }
    if childYangName == "allow-rp" {
        return &ipv6.AllowRp
    }
    if childYangName == "embedded-rp-addresses" {
        return &ipv6.EmbeddedRpAddresses
    }
    if childYangName == "convergence" {
        return &ipv6.Convergence
    }
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    return nil
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sparse-mode-rp-addresses"] = &ipv6.SparseModeRpAddresses
    children["inheritable-defaults"] = &ipv6.InheritableDefaults
    children["rpf"] = &ipv6.Rpf
    children["maximum"] = &ipv6.Maximum
    children["sg-expiry-timer"] = &ipv6.SgExpiryTimer
    children["rpf-vector-enable"] = &ipv6.RpfVectorEnable
    children["ssm"] = &ipv6.Ssm
    children["bidir-rp-addresses"] = &ipv6.BidirRpAddresses
    children["bsr"] = &ipv6.Bsr
    children["allow-rp"] = &ipv6.AllowRp
    children["embedded-rp-addresses"] = &ipv6.EmbeddedRpAddresses
    children["convergence"] = &ipv6.Convergence
    children["interfaces"] = &ipv6.Interfaces
    return children
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-check-on-receive"] = ipv6.NeighborCheckOnReceive
    leafs["old-register-checksum"] = ipv6.OldRegisterChecksum
    leafs["neighbor-filter"] = ipv6.NeighborFilter
    leafs["spt-threshold-infinity"] = ipv6.SptThresholdInfinity
    leafs["log-neighbor-changes"] = ipv6.LogNeighborChanges
    leafs["register-source"] = ipv6.RegisterSource
    leafs["accept-register"] = ipv6.AcceptRegister
    leafs["embedded-rp-disable"] = ipv6.EmbeddedRpDisable
    leafs["suppress-rpf-prunes"] = ipv6.SuppressRpfPrunes
    leafs["ssm-allow-override"] = ipv6.SsmAllowOverride
    leafs["multipath"] = ipv6.Multipath
    leafs["rp-static-deny"] = ipv6.RpStaticDeny
    leafs["suppress-data-registers"] = ipv6.SuppressDataRegisters
    leafs["neighbor-check-on-send"] = ipv6.NeighborCheckOnSend
    return leafs
}

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Pim_Vrfs_Vrf_Ipv6) GetParentYangName() string { return "vrf" }

// Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetFilter() yfilter.YFilter { return sparseModeRpAddresses.YFilter }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) SetFilter(yf yfilter.YFilter) { sparseModeRpAddresses.YFilter = yf }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetGoName(yname string) string {
    if yname == "sparse-mode-rp-address" { return "SparseModeRpAddress" }
    return ""
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetSegmentPath() string {
    return "sparse-mode-rp-addresses"
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-address" {
        for _, c := range sparseModeRpAddresses.SparseModeRpAddress {
            if sparseModeRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress{}
        sparseModeRpAddresses.SparseModeRpAddress = append(sparseModeRpAddresses.SparseModeRpAddress, child)
        return &sparseModeRpAddresses.SparseModeRpAddress[len(sparseModeRpAddresses.SparseModeRpAddress)-1]
    }
    return nil
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        children[sparseModeRpAddresses.SparseModeRpAddress[i].GetSegmentPath()] = &sparseModeRpAddresses.SparseModeRpAddress[i]
    }
    return children
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetYangName() string { return "sparse-mode-rp-addresses" }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) SetParent(parent types.Entity) { sparseModeRpAddresses.parent = parent }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetParent() types.Entity { return sparseModeRpAddresses.parent }

func (sparseModeRpAddresses *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetFilter() yfilter.YFilter { return sparseModeRpAddress.YFilter }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) SetFilter(yf yfilter.YFilter) { sparseModeRpAddress.YFilter = yf }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetSegmentPath() string {
    return "sparse-mode-rp-address" + "[rp-address='" + fmt.Sprintf("%v", sparseModeRpAddress.RpAddress) + "']"
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = sparseModeRpAddress.RpAddress
    leafs["access-list-name"] = sparseModeRpAddress.AccessListName
    leafs["auto-rp-override"] = sparseModeRpAddress.AutoRpOverride
    return leafs
}

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetYangName() string { return "sparse-mode-rp-address" }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) SetParent(parent types.Entity) { sparseModeRpAddress.parent = parent }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetParent() types.Entity { return sparseModeRpAddress.parent }

func (sparseModeRpAddress *Pim_Vrfs_Vrf_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetParentYangName() string { return "sparse-mode-rp-addresses" }

// Pim_Vrfs_Vrf_Ipv6_InheritableDefaults
// Inheritable defaults
type Pim_Vrfs_Vrf_Ipv6_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetGoName(yname string) string {
    if yname == "convergence-timeout" { return "ConvergenceTimeout" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    return ""
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["convergence-timeout"] = inheritableDefaults.ConvergenceTimeout
    leafs["hello-interval"] = inheritableDefaults.HelloInterval
    leafs["propagation-delay"] = inheritableDefaults.PropagationDelay
    leafs["dr-priority"] = inheritableDefaults.DrPriority
    leafs["join-prune-mtu"] = inheritableDefaults.JoinPruneMtu
    leafs["jp-interval"] = inheritableDefaults.JpInterval
    leafs["override-interval"] = inheritableDefaults.OverrideInterval
    return leafs
}

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Pim_Vrfs_Vrf_Ipv6_InheritableDefaults) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Rpf
// Configure RPF options
type Pim_Vrfs_Vrf_Ipv6_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy"] = rpf.RoutePolicy
    return leafs
}

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetYangName() string { return "rpf" }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Pim_Vrfs_Vrf_Ipv6_Rpf) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Maximum
// Configure PIM State Limits
type Pim_Vrfs_Vrf_Ipv6_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_Vrfs_Vrf_Ipv6_Maximum_Routes
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetGoName(yname string) string {
    if yname == "group-mappings-auto-rp" { return "GroupMappingsAutoRp" }
    if yname == "bsr-group-mappings" { return "BsrGroupMappings" }
    if yname == "register-states" { return "RegisterStates" }
    if yname == "route-interfaces" { return "RouteInterfaces" }
    if yname == "bsr-candidate-rp-cache" { return "BsrCandidateRpCache" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group-mappings-auto-rp" {
        return &maximum.GroupMappingsAutoRp
    }
    if childYangName == "bsr-group-mappings" {
        return &maximum.BsrGroupMappings
    }
    if childYangName == "register-states" {
        return &maximum.RegisterStates
    }
    if childYangName == "route-interfaces" {
        return &maximum.RouteInterfaces
    }
    if childYangName == "bsr-candidate-rp-cache" {
        return &maximum.BsrCandidateRpCache
    }
    if childYangName == "routes" {
        return &maximum.Routes
    }
    return nil
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["group-mappings-auto-rp"] = &maximum.GroupMappingsAutoRp
    children["bsr-group-mappings"] = &maximum.BsrGroupMappings
    children["register-states"] = &maximum.RegisterStates
    children["route-interfaces"] = &maximum.RouteInterfaces
    children["bsr-candidate-rp-cache"] = &maximum.BsrCandidateRpCache
    children["routes"] = &maximum.Routes
    return children
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetYangName() string { return "maximum" }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Pim_Vrfs_Vrf_Ipv6_Maximum) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetFilter() yfilter.YFilter { return groupMappingsAutoRp.YFilter }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { groupMappingsAutoRp.YFilter = yf }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-group-ranges-auto-rp" { return "MaximumGroupRangesAutoRp" }
    if yname == "threshold-group-ranges-auto-rp" { return "ThresholdGroupRangesAutoRp" }
    return ""
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetSegmentPath() string {
    return "group-mappings-auto-rp"
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-group-ranges-auto-rp"] = groupMappingsAutoRp.MaximumGroupRangesAutoRp
    leafs["threshold-group-ranges-auto-rp"] = groupMappingsAutoRp.ThresholdGroupRangesAutoRp
    return leafs
}

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetYangName() string { return "group-mappings-auto-rp" }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) SetParent(parent types.Entity) { groupMappingsAutoRp.parent = parent }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetParent() types.Entity { return groupMappingsAutoRp.parent }

func (groupMappingsAutoRp *Pim_Vrfs_Vrf_Ipv6_Maximum_GroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetFilter() yfilter.YFilter { return bsrGroupMappings.YFilter }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGroupMappings.YFilter = yf }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-group-ranges" { return "BsrMaximumGroupRanges" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetSegmentPath() string {
    return "bsr-group-mappings"
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-group-ranges"] = bsrGroupMappings.BsrMaximumGroupRanges
    leafs["warning-threshold"] = bsrGroupMappings.WarningThreshold
    return leafs
}

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetYangName() string { return "bsr-group-mappings" }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) SetParent(parent types.Entity) { bsrGroupMappings.parent = parent }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetParent() types.Entity { return bsrGroupMappings.parent }

func (bsrGroupMappings *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetFilter() yfilter.YFilter { return registerStates.YFilter }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) SetFilter(yf yfilter.YFilter) { registerStates.YFilter = yf }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetSegmentPath() string {
    return "register-states"
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = registerStates.MaximumRegisterStates
    leafs["warning-threshold"] = registerStates.WarningThreshold
    return leafs
}

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetYangName() string { return "register-states" }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) SetParent(parent types.Entity) { registerStates.parent = parent }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetParent() types.Entity { return registerStates.parent }

func (registerStates *Pim_Vrfs_Vrf_Ipv6_Maximum_RegisterStates) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetFilter() yfilter.YFilter { return routeInterfaces.YFilter }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) SetFilter(yf yfilter.YFilter) { routeInterfaces.YFilter = yf }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetSegmentPath() string {
    return "route-interfaces"
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = routeInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = routeInterfaces.WarningThreshold
    return leafs
}

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetYangName() string { return "route-interfaces" }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) SetParent(parent types.Entity) { routeInterfaces.parent = parent }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetParent() types.Entity { return routeInterfaces.parent }

func (routeInterfaces *Pim_Vrfs_Vrf_Ipv6_Maximum_RouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetFilter() yfilter.YFilter { return bsrCandidateRpCache.YFilter }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrCandidateRpCache.YFilter = yf }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-candidate-rp-cache" { return "BsrMaximumCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetSegmentPath() string {
    return "bsr-candidate-rp-cache"
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-candidate-rp-cache"] = bsrCandidateRpCache.BsrMaximumCandidateRpCache
    leafs["warning-threshold"] = bsrCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetYangName() string { return "bsr-candidate-rp-cache" }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) SetParent(parent types.Entity) { bsrCandidateRpCache.parent = parent }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetParent() types.Entity { return bsrCandidateRpCache.parent }

func (bsrCandidateRpCache *Pim_Vrfs_Vrf_Ipv6_Maximum_BsrCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Maximum_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = routes.MaximumRoutes
    leafs["warning-threshold"] = routes.WarningThreshold
    return leafs
}

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetYangName() string { return "routes" }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Pim_Vrfs_Vrf_Ipv6_Maximum_Routes) GetParentYangName() string { return "maximum" }

// Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetFilter() yfilter.YFilter { return sgExpiryTimer.YFilter }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) SetFilter(yf yfilter.YFilter) { sgExpiryTimer.YFilter = yf }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetSegmentPath() string {
    return "sg-expiry-timer"
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = sgExpiryTimer.Interval
    leafs["access-list-name"] = sgExpiryTimer.AccessListName
    return leafs
}

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetBundleName() string { return "cisco_ios_xr" }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetYangName() string { return "sg-expiry-timer" }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) SetParent(parent types.Entity) { sgExpiryTimer.parent = parent }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetParent() types.Entity { return sgExpiryTimer.parent }

func (sgExpiryTimer *Pim_Vrfs_Vrf_Ipv6_SgExpiryTimer) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetFilter() yfilter.YFilter { return rpfVectorEnable.YFilter }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) SetFilter(yf yfilter.YFilter) { rpfVectorEnable.YFilter = yf }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-ebgp" { return "AllowEbgp" }
    if yname == "disable-ibgp" { return "DisableIbgp" }
    return ""
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetSegmentPath() string {
    return "rpf-vector-enable"
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpfVectorEnable.Enable
    leafs["allow-ebgp"] = rpfVectorEnable.AllowEbgp
    leafs["disable-ibgp"] = rpfVectorEnable.DisableIbgp
    return leafs
}

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetBundleName() string { return "cisco_ios_xr" }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetYangName() string { return "rpf-vector-enable" }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) SetParent(parent types.Entity) { rpfVectorEnable.parent = parent }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetParent() types.Entity { return rpfVectorEnable.parent }

func (rpfVectorEnable *Pim_Vrfs_Vrf_Ipv6_RpfVectorEnable) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Ssm
// Configure IP Multicast SSM
type Pim_Vrfs_Vrf_Ipv6_Ssm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetFilter() yfilter.YFilter { return ssm.YFilter }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) SetFilter(yf yfilter.YFilter) { ssm.YFilter = yf }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "range" { return "Range" }
    return ""
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetSegmentPath() string {
    return "ssm"
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = ssm.Disable
    leafs["range"] = ssm.Range
    return leafs
}

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetBundleName() string { return "cisco_ios_xr" }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetYangName() string { return "ssm" }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) SetParent(parent types.Entity) { ssm.parent = parent }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetParent() types.Entity { return ssm.parent }

func (ssm *Pim_Vrfs_Vrf_Ipv6_Ssm) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetFilter() yfilter.YFilter { return bidirRpAddresses.YFilter }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) SetFilter(yf yfilter.YFilter) { bidirRpAddresses.YFilter = yf }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetGoName(yname string) string {
    if yname == "bidir-rp-address" { return "BidirRpAddress" }
    return ""
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetSegmentPath() string {
    return "bidir-rp-addresses"
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bidir-rp-address" {
        for _, c := range bidirRpAddresses.BidirRpAddress {
            if bidirRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress{}
        bidirRpAddresses.BidirRpAddress = append(bidirRpAddresses.BidirRpAddress, child)
        return &bidirRpAddresses.BidirRpAddress[len(bidirRpAddresses.BidirRpAddress)-1]
    }
    return nil
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bidirRpAddresses.BidirRpAddress {
        children[bidirRpAddresses.BidirRpAddress[i].GetSegmentPath()] = &bidirRpAddresses.BidirRpAddress[i]
    }
    return children
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetYangName() string { return "bidir-rp-addresses" }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) SetParent(parent types.Entity) { bidirRpAddresses.parent = parent }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetParent() types.Entity { return bidirRpAddresses.parent }

func (bidirRpAddresses *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetFilter() yfilter.YFilter { return bidirRpAddress.YFilter }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) SetFilter(yf yfilter.YFilter) { bidirRpAddress.YFilter = yf }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetSegmentPath() string {
    return "bidir-rp-address" + "[rp-address='" + fmt.Sprintf("%v", bidirRpAddress.RpAddress) + "']"
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = bidirRpAddress.RpAddress
    leafs["access-list-name"] = bidirRpAddress.AccessListName
    leafs["auto-rp-override"] = bidirRpAddress.AutoRpOverride
    return leafs
}

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetYangName() string { return "bidir-rp-address" }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) SetParent(parent types.Entity) { bidirRpAddress.parent = parent }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetParent() types.Entity { return bidirRpAddress.parent }

func (bidirRpAddress *Pim_Vrfs_Vrf_Ipv6_BidirRpAddresses_BidirRpAddress) GetParentYangName() string { return "bidir-rp-addresses" }

// Pim_Vrfs_Vrf_Ipv6_Bsr
// PIM BSR configuration
type Pim_Vrfs_Vrf_Ipv6_Bsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetFilter() yfilter.YFilter { return bsr.YFilter }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) SetFilter(yf yfilter.YFilter) { bsr.YFilter = yf }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetGoName(yname string) string {
    if yname == "candidate-bsr" { return "CandidateBsr" }
    if yname == "candidate-rps" { return "CandidateRps" }
    return ""
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetSegmentPath() string {
    return "bsr"
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-bsr" {
        return &bsr.CandidateBsr
    }
    if childYangName == "candidate-rps" {
        return &bsr.CandidateRps
    }
    return nil
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-bsr"] = &bsr.CandidateBsr
    children["candidate-rps"] = &bsr.CandidateRps
    return children
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetBundleName() string { return "cisco_ios_xr" }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetYangName() string { return "bsr" }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) SetParent(parent types.Entity) { bsr.parent = parent }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetParent() types.Entity { return bsr.parent }

func (bsr *Pim_Vrfs_Vrf_Ipv6_Bsr) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BSR Address configured. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..128. The default value is 126.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetFilter() yfilter.YFilter { return candidateBsr.YFilter }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) SetFilter(yf yfilter.YFilter) { candidateBsr.YFilter = yf }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetSegmentPath() string {
    return "candidate-bsr"
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateBsr.Address
    leafs["prefix-length"] = candidateBsr.PrefixLength
    leafs["priority"] = candidateBsr.Priority
    return leafs
}

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetBundleName() string { return "cisco_ios_xr" }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetYangName() string { return "candidate-bsr" }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) SetParent(parent types.Entity) { candidateBsr.parent = parent }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetParent() types.Entity { return candidateBsr.parent }

func (candidateBsr *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateBsr) GetParentYangName() string { return "bsr" }

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps
// PIM RP configuration
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp.
    CandidateRp []Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetFilter() yfilter.YFilter { return candidateRps.YFilter }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) SetFilter(yf yfilter.YFilter) { candidateRps.YFilter = yf }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    return ""
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetSegmentPath() string {
    return "candidate-rps"
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        for _, c := range candidateRps.CandidateRp {
            if candidateRps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp{}
        candidateRps.CandidateRp = append(candidateRps.CandidateRp, child)
        return &candidateRps.CandidateRp[len(candidateRps.CandidateRp)-1]
    }
    return nil
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateRps.CandidateRp {
        children[candidateRps.CandidateRp[i].GetSegmentPath()] = &candidateRps.CandidateRp[i]
    }
    return children
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetYangName() string { return "candidate-rps" }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) SetParent(parent types.Entity) { candidateRps.parent = parent }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetParent() types.Entity { return candidateRps.parent }

func (candidateRps *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps) GetParentYangName() string { return "bsr" }

// Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mode" { return "Mode" }
    if yname == "group-list" { return "GroupList" }
    if yname == "priority" { return "Priority" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetSegmentPath() string {
    return "candidate-rp" + "[address='" + fmt.Sprintf("%v", candidateRp.Address) + "']" + "[mode='" + fmt.Sprintf("%v", candidateRp.Mode) + "']"
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateRp.Address
    leafs["mode"] = candidateRp.Mode
    leafs["group-list"] = candidateRp.GroupList
    leafs["priority"] = candidateRp.Priority
    leafs["interval"] = candidateRp.Interval
    return leafs
}

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *Pim_Vrfs_Vrf_Ipv6_Bsr_CandidateRps_CandidateRp) GetParentYangName() string { return "candidate-rps" }

// Pim_Vrfs_Vrf_Ipv6_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_AllowRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetFilter() yfilter.YFilter { return allowRp.YFilter }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) SetFilter(yf yfilter.YFilter) { allowRp.YFilter = yf }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetGoName(yname string) string {
    if yname == "rp-list-name" { return "RpListName" }
    if yname == "group-list-name" { return "GroupListName" }
    return ""
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetSegmentPath() string {
    return "allow-rp"
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-list-name"] = allowRp.RpListName
    leafs["group-list-name"] = allowRp.GroupListName
    return leafs
}

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetBundleName() string { return "cisco_ios_xr" }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetYangName() string { return "allow-rp" }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) SetParent(parent types.Entity) { allowRp.parent = parent }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetParent() types.Entity { return allowRp.parent }

func (allowRp *Pim_Vrfs_Vrf_Ipv6_AllowRp) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses
// Set Embedded RP processing support
type Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set Embedded RP processing support. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress.
    EmbeddedRpAddress []Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetFilter() yfilter.YFilter { return embeddedRpAddresses.YFilter }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) SetFilter(yf yfilter.YFilter) { embeddedRpAddresses.YFilter = yf }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetGoName(yname string) string {
    if yname == "embedded-rp-address" { return "EmbeddedRpAddress" }
    return ""
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetSegmentPath() string {
    return "embedded-rp-addresses"
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "embedded-rp-address" {
        for _, c := range embeddedRpAddresses.EmbeddedRpAddress {
            if embeddedRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress{}
        embeddedRpAddresses.EmbeddedRpAddress = append(embeddedRpAddresses.EmbeddedRpAddress, child)
        return &embeddedRpAddresses.EmbeddedRpAddress[len(embeddedRpAddresses.EmbeddedRpAddress)-1]
    }
    return nil
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range embeddedRpAddresses.EmbeddedRpAddress {
        children[embeddedRpAddresses.EmbeddedRpAddress[i].GetSegmentPath()] = &embeddedRpAddresses.EmbeddedRpAddress[i]
    }
    return children
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetYangName() string { return "embedded-rp-addresses" }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) SetParent(parent types.Entity) { embeddedRpAddresses.parent = parent }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetParent() types.Entity { return embeddedRpAddresses.parent }

func (embeddedRpAddresses *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
// Set Embedded RP processing support
type Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of the Rendezvous Point. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64. This attribute is mandatory.
    AccessListName interface{}
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetFilter() yfilter.YFilter { return embeddedRpAddress.YFilter }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) SetFilter(yf yfilter.YFilter) { embeddedRpAddress.YFilter = yf }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetSegmentPath() string {
    return "embedded-rp-address" + "[rp-address='" + fmt.Sprintf("%v", embeddedRpAddress.RpAddress) + "']"
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = embeddedRpAddress.RpAddress
    leafs["access-list-name"] = embeddedRpAddress.AccessListName
    return leafs
}

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetYangName() string { return "embedded-rp-address" }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) SetParent(parent types.Entity) { embeddedRpAddress.parent = parent }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetParent() types.Entity { return embeddedRpAddress.parent }

func (embeddedRpAddress *Pim_Vrfs_Vrf_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetParentYangName() string { return "embedded-rp-addresses" }

// Pim_Vrfs_Vrf_Ipv6_Convergence
// Configure convergence parameters
type Pim_Vrfs_Vrf_Ipv6_Convergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetFilter() yfilter.YFilter { return convergence.YFilter }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) SetFilter(yf yfilter.YFilter) { convergence.YFilter = yf }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetGoName(yname string) string {
    if yname == "rpf-conflict-join-delay" { return "RpfConflictJoinDelay" }
    if yname == "link-down-prune-delay" { return "LinkDownPruneDelay" }
    return ""
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetSegmentPath() string {
    return "convergence"
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rpf-conflict-join-delay"] = convergence.RpfConflictJoinDelay
    leafs["link-down-prune-delay"] = convergence.LinkDownPruneDelay
    return leafs
}

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetBundleName() string { return "cisco_ios_xr" }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetYangName() string { return "convergence" }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) SetParent(parent types.Entity) { convergence.parent = parent }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetParent() types.Entity { return convergence.parent }

func (convergence *Pim_Vrfs_Vrf_Ipv6_Convergence) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Interfaces
// Interface-level Configuration
type Pim_Vrfs_Vrf_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface.
    Interface []Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pim_Vrfs_Vrf_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface
// The name of the interface
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "bsr-border" { return "BsrBorder" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "interface-enable" { return "InterfaceEnable" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "bfd" { return "Bfd" }
    return ""
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-routes" {
        return &self.MaximumRoutes
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    return nil
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-routes"] = &self.MaximumRoutes
    children["bfd"] = &self.Bfd
    return children
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    leafs["neighbor-filter"] = self.NeighborFilter
    leafs["hello-interval"] = self.HelloInterval
    leafs["bsr-border"] = self.BsrBorder
    leafs["propagation-delay"] = self.PropagationDelay
    leafs["dr-priority"] = self.DrPriority
    leafs["join-prune-mtu"] = self.JoinPruneMtu
    leafs["interface-enable"] = self.InterfaceEnable
    leafs["jp-interval"] = self.JpInterval
    leafs["override-interval"] = self.OverrideInterval
    return leafs
}

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetFilter() yfilter.YFilter { return maximumRoutes.YFilter }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) SetFilter(yf yfilter.YFilter) { maximumRoutes.YFilter = yf }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetSegmentPath() string {
    return "maximum-routes"
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = maximumRoutes.Maximum
    leafs["warning-threshold"] = maximumRoutes.WarningThreshold
    leafs["access-list-name"] = maximumRoutes.AccessListName
    return leafs
}

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetYangName() string { return "maximum-routes" }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) SetParent(parent types.Entity) { maximumRoutes.parent = parent }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetParent() types.Entity { return maximumRoutes.parent }

func (maximumRoutes *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_MaximumRoutes) GetParentYangName() string { return "interface" }

// Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd
// BFD configuration
type Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    if yname == "interval" { return "Interval" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    leafs["interval"] = bfd.Interval
    leafs["enable"] = bfd.Enable
    return leafs
}

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Pim_Vrfs_Vrf_Ipv6_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Pim_DefaultContext
// Default Context
// This type is a presence type.
type Pim_DefaultContext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6 commands.
    Ipv6 Pim_DefaultContext_Ipv6

    // IPV4 commands.
    Ipv4 Pim_DefaultContext_Ipv4
}

func (defaultContext *Pim_DefaultContext) GetFilter() yfilter.YFilter { return defaultContext.YFilter }

func (defaultContext *Pim_DefaultContext) SetFilter(yf yfilter.YFilter) { defaultContext.YFilter = yf }

func (defaultContext *Pim_DefaultContext) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (defaultContext *Pim_DefaultContext) GetSegmentPath() string {
    return "default-context"
}

func (defaultContext *Pim_DefaultContext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &defaultContext.Ipv6
    }
    if childYangName == "ipv4" {
        return &defaultContext.Ipv4
    }
    return nil
}

func (defaultContext *Pim_DefaultContext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &defaultContext.Ipv6
    children["ipv4"] = &defaultContext.Ipv4
    return children
}

func (defaultContext *Pim_DefaultContext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultContext *Pim_DefaultContext) GetBundleName() string { return "cisco_ios_xr" }

func (defaultContext *Pim_DefaultContext) GetYangName() string { return "default-context" }

func (defaultContext *Pim_DefaultContext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultContext *Pim_DefaultContext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultContext *Pim_DefaultContext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultContext *Pim_DefaultContext) SetParent(parent types.Entity) { defaultContext.parent = parent }

func (defaultContext *Pim_DefaultContext) GetParent() types.Entity { return defaultContext.parent }

func (defaultContext *Pim_DefaultContext) GetParentYangName() string { return "pim" }

// Pim_DefaultContext_Ipv6
// IPV6 commands
type Pim_DefaultContext_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Set Embedded RP processing support. The type is interface{}.
    EmbeddedRpDisable interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Interface-level Configuration.
    Interfaces Pim_DefaultContext_Ipv6_Interfaces

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_DefaultContext_Ipv6_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_DefaultContext_Ipv6_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_DefaultContext_Ipv6_Rpf

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_DefaultContext_Ipv6_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_DefaultContext_Ipv6_RpfVectorEnable

    // Configure Non-stop forwarding (NSF) options.
    Nsf Pim_DefaultContext_Ipv6_Nsf

    // Configure PIM State Limits.
    Maximum Pim_DefaultContext_Ipv6_Maximum

    // Configure IP Multicast SSM.
    Ssm Pim_DefaultContext_Ipv6_Ssm

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_DefaultContext_Ipv6_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_DefaultContext_Ipv6_Bsr

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_DefaultContext_Ipv6_AllowRp

    // Set Embedded RP processing support.
    EmbeddedRpAddresses Pim_DefaultContext_Ipv6_EmbeddedRpAddresses

    // Configure convergence parameters.
    Convergence Pim_DefaultContext_Ipv6_Convergence
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Pim_DefaultContext_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Pim_DefaultContext_Ipv6) GetGoName(yname string) string {
    if yname == "neighbor-check-on-receive" { return "NeighborCheckOnReceive" }
    if yname == "old-register-checksum" { return "OldRegisterChecksum" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "spt-threshold-infinity" { return "SptThresholdInfinity" }
    if yname == "log-neighbor-changes" { return "LogNeighborChanges" }
    if yname == "register-source" { return "RegisterSource" }
    if yname == "accept-register" { return "AcceptRegister" }
    if yname == "embedded-rp-disable" { return "EmbeddedRpDisable" }
    if yname == "suppress-rpf-prunes" { return "SuppressRpfPrunes" }
    if yname == "ssm-allow-override" { return "SsmAllowOverride" }
    if yname == "multipath" { return "Multipath" }
    if yname == "rp-static-deny" { return "RpStaticDeny" }
    if yname == "suppress-data-registers" { return "SuppressDataRegisters" }
    if yname == "neighbor-check-on-send" { return "NeighborCheckOnSend" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "sparse-mode-rp-addresses" { return "SparseModeRpAddresses" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "rpf" { return "Rpf" }
    if yname == "sg-expiry-timer" { return "SgExpiryTimer" }
    if yname == "rpf-vector-enable" { return "RpfVectorEnable" }
    if yname == "nsf" { return "Nsf" }
    if yname == "maximum" { return "Maximum" }
    if yname == "ssm" { return "Ssm" }
    if yname == "bidir-rp-addresses" { return "BidirRpAddresses" }
    if yname == "bsr" { return "Bsr" }
    if yname == "allow-rp" { return "AllowRp" }
    if yname == "embedded-rp-addresses" { return "EmbeddedRpAddresses" }
    if yname == "convergence" { return "Convergence" }
    return ""
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    if childYangName == "sparse-mode-rp-addresses" {
        return &ipv6.SparseModeRpAddresses
    }
    if childYangName == "inheritable-defaults" {
        return &ipv6.InheritableDefaults
    }
    if childYangName == "rpf" {
        return &ipv6.Rpf
    }
    if childYangName == "sg-expiry-timer" {
        return &ipv6.SgExpiryTimer
    }
    if childYangName == "rpf-vector-enable" {
        return &ipv6.RpfVectorEnable
    }
    if childYangName == "nsf" {
        return &ipv6.Nsf
    }
    if childYangName == "maximum" {
        return &ipv6.Maximum
    }
    if childYangName == "ssm" {
        return &ipv6.Ssm
    }
    if childYangName == "bidir-rp-addresses" {
        return &ipv6.BidirRpAddresses
    }
    if childYangName == "bsr" {
        return &ipv6.Bsr
    }
    if childYangName == "allow-rp" {
        return &ipv6.AllowRp
    }
    if childYangName == "embedded-rp-addresses" {
        return &ipv6.EmbeddedRpAddresses
    }
    if childYangName == "convergence" {
        return &ipv6.Convergence
    }
    return nil
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &ipv6.Interfaces
    children["sparse-mode-rp-addresses"] = &ipv6.SparseModeRpAddresses
    children["inheritable-defaults"] = &ipv6.InheritableDefaults
    children["rpf"] = &ipv6.Rpf
    children["sg-expiry-timer"] = &ipv6.SgExpiryTimer
    children["rpf-vector-enable"] = &ipv6.RpfVectorEnable
    children["nsf"] = &ipv6.Nsf
    children["maximum"] = &ipv6.Maximum
    children["ssm"] = &ipv6.Ssm
    children["bidir-rp-addresses"] = &ipv6.BidirRpAddresses
    children["bsr"] = &ipv6.Bsr
    children["allow-rp"] = &ipv6.AllowRp
    children["embedded-rp-addresses"] = &ipv6.EmbeddedRpAddresses
    children["convergence"] = &ipv6.Convergence
    return children
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-check-on-receive"] = ipv6.NeighborCheckOnReceive
    leafs["old-register-checksum"] = ipv6.OldRegisterChecksum
    leafs["neighbor-filter"] = ipv6.NeighborFilter
    leafs["spt-threshold-infinity"] = ipv6.SptThresholdInfinity
    leafs["log-neighbor-changes"] = ipv6.LogNeighborChanges
    leafs["register-source"] = ipv6.RegisterSource
    leafs["accept-register"] = ipv6.AcceptRegister
    leafs["embedded-rp-disable"] = ipv6.EmbeddedRpDisable
    leafs["suppress-rpf-prunes"] = ipv6.SuppressRpfPrunes
    leafs["ssm-allow-override"] = ipv6.SsmAllowOverride
    leafs["multipath"] = ipv6.Multipath
    leafs["rp-static-deny"] = ipv6.RpStaticDeny
    leafs["suppress-data-registers"] = ipv6.SuppressDataRegisters
    leafs["neighbor-check-on-send"] = ipv6.NeighborCheckOnSend
    return leafs
}

func (ipv6 *Pim_DefaultContext_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Pim_DefaultContext_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Pim_DefaultContext_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Pim_DefaultContext_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Pim_DefaultContext_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Pim_DefaultContext_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Pim_DefaultContext_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Pim_DefaultContext_Ipv6) GetParentYangName() string { return "default-context" }

// Pim_DefaultContext_Ipv6_Interfaces
// Interface-level Configuration
type Pim_DefaultContext_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_DefaultContext_Ipv6_Interfaces_Interface.
    Interface []Pim_DefaultContext_Ipv6_Interfaces_Interface
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pim_DefaultContext_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Interfaces_Interface
// The name of the interface
type Pim_DefaultContext_Ipv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "bsr-border" { return "BsrBorder" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "interface-enable" { return "InterfaceEnable" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "bfd" { return "Bfd" }
    return ""
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "maximum-routes" {
        return &self.MaximumRoutes
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    return nil
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["maximum-routes"] = &self.MaximumRoutes
    children["bfd"] = &self.Bfd
    return children
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    leafs["neighbor-filter"] = self.NeighborFilter
    leafs["hello-interval"] = self.HelloInterval
    leafs["bsr-border"] = self.BsrBorder
    leafs["propagation-delay"] = self.PropagationDelay
    leafs["dr-priority"] = self.DrPriority
    leafs["join-prune-mtu"] = self.JoinPruneMtu
    leafs["interface-enable"] = self.InterfaceEnable
    leafs["jp-interval"] = self.JpInterval
    leafs["override-interval"] = self.OverrideInterval
    return leafs
}

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pim_DefaultContext_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetFilter() yfilter.YFilter { return maximumRoutes.YFilter }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) SetFilter(yf yfilter.YFilter) { maximumRoutes.YFilter = yf }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetSegmentPath() string {
    return "maximum-routes"
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = maximumRoutes.Maximum
    leafs["warning-threshold"] = maximumRoutes.WarningThreshold
    leafs["access-list-name"] = maximumRoutes.AccessListName
    return leafs
}

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetYangName() string { return "maximum-routes" }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) SetParent(parent types.Entity) { maximumRoutes.parent = parent }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetParent() types.Entity { return maximumRoutes.parent }

func (maximumRoutes *Pim_DefaultContext_Ipv6_Interfaces_Interface_MaximumRoutes) GetParentYangName() string { return "interface" }

// Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd
// BFD configuration
type Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    if yname == "interval" { return "Interval" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    leafs["interval"] = bfd.Interval
    leafs["enable"] = bfd.Enable
    return leafs
}

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Pim_DefaultContext_Ipv6_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Pim_DefaultContext_Ipv6_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_DefaultContext_Ipv6_SparseModeRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetFilter() yfilter.YFilter { return sparseModeRpAddresses.YFilter }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) SetFilter(yf yfilter.YFilter) { sparseModeRpAddresses.YFilter = yf }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetGoName(yname string) string {
    if yname == "sparse-mode-rp-address" { return "SparseModeRpAddress" }
    return ""
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetSegmentPath() string {
    return "sparse-mode-rp-addresses"
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-address" {
        for _, c := range sparseModeRpAddresses.SparseModeRpAddress {
            if sparseModeRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress{}
        sparseModeRpAddresses.SparseModeRpAddress = append(sparseModeRpAddresses.SparseModeRpAddress, child)
        return &sparseModeRpAddresses.SparseModeRpAddress[len(sparseModeRpAddresses.SparseModeRpAddress)-1]
    }
    return nil
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        children[sparseModeRpAddresses.SparseModeRpAddress[i].GetSegmentPath()] = &sparseModeRpAddresses.SparseModeRpAddress[i]
    }
    return children
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetYangName() string { return "sparse-mode-rp-addresses" }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) SetParent(parent types.Entity) { sparseModeRpAddresses.parent = parent }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetParent() types.Entity { return sparseModeRpAddresses.parent }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv6_SparseModeRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetFilter() yfilter.YFilter { return sparseModeRpAddress.YFilter }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) SetFilter(yf yfilter.YFilter) { sparseModeRpAddress.YFilter = yf }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetSegmentPath() string {
    return "sparse-mode-rp-address" + "[rp-address='" + fmt.Sprintf("%v", sparseModeRpAddress.RpAddress) + "']"
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = sparseModeRpAddress.RpAddress
    leafs["access-list-name"] = sparseModeRpAddress.AccessListName
    leafs["auto-rp-override"] = sparseModeRpAddress.AutoRpOverride
    return leafs
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetYangName() string { return "sparse-mode-rp-address" }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) SetParent(parent types.Entity) { sparseModeRpAddress.parent = parent }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetParent() types.Entity { return sparseModeRpAddress.parent }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv6_SparseModeRpAddresses_SparseModeRpAddress) GetParentYangName() string { return "sparse-mode-rp-addresses" }

// Pim_DefaultContext_Ipv6_InheritableDefaults
// Inheritable defaults
type Pim_DefaultContext_Ipv6_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetGoName(yname string) string {
    if yname == "convergence-timeout" { return "ConvergenceTimeout" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    return ""
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["convergence-timeout"] = inheritableDefaults.ConvergenceTimeout
    leafs["hello-interval"] = inheritableDefaults.HelloInterval
    leafs["propagation-delay"] = inheritableDefaults.PropagationDelay
    leafs["dr-priority"] = inheritableDefaults.DrPriority
    leafs["join-prune-mtu"] = inheritableDefaults.JoinPruneMtu
    leafs["jp-interval"] = inheritableDefaults.JpInterval
    leafs["override-interval"] = inheritableDefaults.OverrideInterval
    return leafs
}

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Pim_DefaultContext_Ipv6_InheritableDefaults) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Rpf
// Configure RPF options
type Pim_DefaultContext_Ipv6_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy"] = rpf.RoutePolicy
    return leafs
}

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetYangName() string { return "rpf" }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Pim_DefaultContext_Ipv6_Rpf) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_DefaultContext_Ipv6_SgExpiryTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetFilter() yfilter.YFilter { return sgExpiryTimer.YFilter }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) SetFilter(yf yfilter.YFilter) { sgExpiryTimer.YFilter = yf }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetSegmentPath() string {
    return "sg-expiry-timer"
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = sgExpiryTimer.Interval
    leafs["access-list-name"] = sgExpiryTimer.AccessListName
    return leafs
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetBundleName() string { return "cisco_ios_xr" }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetYangName() string { return "sg-expiry-timer" }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) SetParent(parent types.Entity) { sgExpiryTimer.parent = parent }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetParent() types.Entity { return sgExpiryTimer.parent }

func (sgExpiryTimer *Pim_DefaultContext_Ipv6_SgExpiryTimer) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_DefaultContext_Ipv6_RpfVectorEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetFilter() yfilter.YFilter { return rpfVectorEnable.YFilter }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) SetFilter(yf yfilter.YFilter) { rpfVectorEnable.YFilter = yf }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-ebgp" { return "AllowEbgp" }
    if yname == "disable-ibgp" { return "DisableIbgp" }
    return ""
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetSegmentPath() string {
    return "rpf-vector-enable"
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpfVectorEnable.Enable
    leafs["allow-ebgp"] = rpfVectorEnable.AllowEbgp
    leafs["disable-ibgp"] = rpfVectorEnable.DisableIbgp
    return leafs
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetBundleName() string { return "cisco_ios_xr" }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetYangName() string { return "rpf-vector-enable" }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) SetParent(parent types.Entity) { rpfVectorEnable.parent = parent }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetParent() types.Entity { return rpfVectorEnable.parent }

func (rpfVectorEnable *Pim_DefaultContext_Ipv6_RpfVectorEnable) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Nsf
// Configure Non-stop forwarding (NSF) options
type Pim_DefaultContext_Ipv6_Nsf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Override default maximum lifetime for PIM NSF mode. The type is interface{}
    // with range: 10..600. Units are second.
    Lifetime interface{}
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetFilter() yfilter.YFilter { return nsf.YFilter }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) SetFilter(yf yfilter.YFilter) { nsf.YFilter = yf }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetGoName(yname string) string {
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetSegmentPath() string {
    return "nsf"
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lifetime"] = nsf.Lifetime
    return leafs
}

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetBundleName() string { return "cisco_ios_xr" }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetYangName() string { return "nsf" }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) SetParent(parent types.Entity) { nsf.parent = parent }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetParent() types.Entity { return nsf.parent }

func (nsf *Pim_DefaultContext_Ipv6_Nsf) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Maximum
// Configure PIM State Limits
type Pim_DefaultContext_Ipv6_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalLowPriorityPacketQueue interface{}

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalHighPriorityPacketQueue interface{}

    // Override default global maximum and threshold for PIM group mapping ranges
    // from BSR.
    BsrGlobalGroupMappings Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings

    // Override default maximum for number of routes.
    GlobalRoutes Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes

    // Maximum for number of group mappings from autorp mapping agent.
    GlobalGroupMappingsAutoRp Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp

    // Override default global maximum and threshold for C-RP set in BSR.
    BsrGlobalCandidateRpCache Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache

    // Override default maximum for number of sparse-mode source registers.
    GlobalRegisterStates Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates

    // Override default maximum for number of route-interfaces.
    GlobalRouteInterfaces Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_DefaultContext_Ipv6_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_DefaultContext_Ipv6_Maximum_Routes
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetGoName(yname string) string {
    if yname == "global-low-priority-packet-queue" { return "GlobalLowPriorityPacketQueue" }
    if yname == "global-high-priority-packet-queue" { return "GlobalHighPriorityPacketQueue" }
    if yname == "bsr-global-group-mappings" { return "BsrGlobalGroupMappings" }
    if yname == "global-routes" { return "GlobalRoutes" }
    if yname == "global-group-mappings-auto-rp" { return "GlobalGroupMappingsAutoRp" }
    if yname == "bsr-global-candidate-rp-cache" { return "BsrGlobalCandidateRpCache" }
    if yname == "global-register-states" { return "GlobalRegisterStates" }
    if yname == "global-route-interfaces" { return "GlobalRouteInterfaces" }
    if yname == "group-mappings-auto-rp" { return "GroupMappingsAutoRp" }
    if yname == "bsr-group-mappings" { return "BsrGroupMappings" }
    if yname == "register-states" { return "RegisterStates" }
    if yname == "route-interfaces" { return "RouteInterfaces" }
    if yname == "bsr-candidate-rp-cache" { return "BsrCandidateRpCache" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bsr-global-group-mappings" {
        return &maximum.BsrGlobalGroupMappings
    }
    if childYangName == "global-routes" {
        return &maximum.GlobalRoutes
    }
    if childYangName == "global-group-mappings-auto-rp" {
        return &maximum.GlobalGroupMappingsAutoRp
    }
    if childYangName == "bsr-global-candidate-rp-cache" {
        return &maximum.BsrGlobalCandidateRpCache
    }
    if childYangName == "global-register-states" {
        return &maximum.GlobalRegisterStates
    }
    if childYangName == "global-route-interfaces" {
        return &maximum.GlobalRouteInterfaces
    }
    if childYangName == "group-mappings-auto-rp" {
        return &maximum.GroupMappingsAutoRp
    }
    if childYangName == "bsr-group-mappings" {
        return &maximum.BsrGroupMappings
    }
    if childYangName == "register-states" {
        return &maximum.RegisterStates
    }
    if childYangName == "route-interfaces" {
        return &maximum.RouteInterfaces
    }
    if childYangName == "bsr-candidate-rp-cache" {
        return &maximum.BsrCandidateRpCache
    }
    if childYangName == "routes" {
        return &maximum.Routes
    }
    return nil
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bsr-global-group-mappings"] = &maximum.BsrGlobalGroupMappings
    children["global-routes"] = &maximum.GlobalRoutes
    children["global-group-mappings-auto-rp"] = &maximum.GlobalGroupMappingsAutoRp
    children["bsr-global-candidate-rp-cache"] = &maximum.BsrGlobalCandidateRpCache
    children["global-register-states"] = &maximum.GlobalRegisterStates
    children["global-route-interfaces"] = &maximum.GlobalRouteInterfaces
    children["group-mappings-auto-rp"] = &maximum.GroupMappingsAutoRp
    children["bsr-group-mappings"] = &maximum.BsrGroupMappings
    children["register-states"] = &maximum.RegisterStates
    children["route-interfaces"] = &maximum.RouteInterfaces
    children["bsr-candidate-rp-cache"] = &maximum.BsrCandidateRpCache
    children["routes"] = &maximum.Routes
    return children
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-low-priority-packet-queue"] = maximum.GlobalLowPriorityPacketQueue
    leafs["global-high-priority-packet-queue"] = maximum.GlobalHighPriorityPacketQueue
    return leafs
}

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetYangName() string { return "maximum" }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Pim_DefaultContext_Ipv6_Maximum) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings
// Override default global maximum and threshold
// for PIM group mapping ranges from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global Maximum number of PIM group mapping ranges from BSR. The type is
    // interface{} with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalGroupMappings interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetFilter() yfilter.YFilter { return bsrGlobalGroupMappings.YFilter }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGlobalGroupMappings.YFilter = yf }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-global-group-mappings" { return "BsrMaximumGlobalGroupMappings" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetSegmentPath() string {
    return "bsr-global-group-mappings"
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-global-group-mappings"] = bsrGlobalGroupMappings.BsrMaximumGlobalGroupMappings
    leafs["warning-threshold"] = bsrGlobalGroupMappings.WarningThreshold
    return leafs
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetYangName() string { return "bsr-global-group-mappings" }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) SetParent(parent types.Entity) { bsrGlobalGroupMappings.parent = parent }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetParent() types.Entity { return bsrGlobalGroupMappings.parent }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetFilter() yfilter.YFilter { return globalRoutes.YFilter }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) SetFilter(yf yfilter.YFilter) { globalRoutes.YFilter = yf }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetSegmentPath() string {
    return "global-routes"
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = globalRoutes.MaximumRoutes
    leafs["warning-threshold"] = globalRoutes.WarningThreshold
    return leafs
}

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetYangName() string { return "global-routes" }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) SetParent(parent types.Entity) { globalRoutes.parent = parent }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetParent() types.Entity { return globalRoutes.parent }

func (globalRoutes *Pim_DefaultContext_Ipv6_Maximum_GlobalRoutes) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp
// Maximum for number of group mappings from
// autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGlobalGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGlobalGroupRangesAutoRp interface{}
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetFilter() yfilter.YFilter { return globalGroupMappingsAutoRp.YFilter }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { globalGroupMappingsAutoRp.YFilter = yf }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-global-group-ranges-auto-rp" { return "MaximumGlobalGroupRangesAutoRp" }
    if yname == "threshold-global-group-ranges-auto-rp" { return "ThresholdGlobalGroupRangesAutoRp" }
    return ""
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetSegmentPath() string {
    return "global-group-mappings-auto-rp"
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-global-group-ranges-auto-rp"] = globalGroupMappingsAutoRp.MaximumGlobalGroupRangesAutoRp
    leafs["threshold-global-group-ranges-auto-rp"] = globalGroupMappingsAutoRp.ThresholdGlobalGroupRangesAutoRp
    return leafs
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetYangName() string { return "global-group-mappings-auto-rp" }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) SetParent(parent types.Entity) { globalGroupMappingsAutoRp.parent = parent }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetParent() types.Entity { return globalGroupMappingsAutoRp.parent }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GlobalGroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache
// Override default global maximum and threshold
// for C-RP set in BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global Maximum number of PIM C-RP Sets from BSR. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetFilter() yfilter.YFilter { return bsrGlobalCandidateRpCache.YFilter }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrGlobalCandidateRpCache.YFilter = yf }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-global-candidate-rp-cache" { return "BsrMaximumGlobalCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetSegmentPath() string {
    return "bsr-global-candidate-rp-cache"
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-global-candidate-rp-cache"] = bsrGlobalCandidateRpCache.BsrMaximumGlobalCandidateRpCache
    leafs["warning-threshold"] = bsrGlobalCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetYangName() string { return "bsr-global-candidate-rp-cache" }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) SetParent(parent types.Entity) { bsrGlobalCandidateRpCache.parent = parent }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetParent() types.Entity { return bsrGlobalCandidateRpCache.parent }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrGlobalCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetFilter() yfilter.YFilter { return globalRegisterStates.YFilter }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) SetFilter(yf yfilter.YFilter) { globalRegisterStates.YFilter = yf }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetSegmentPath() string {
    return "global-register-states"
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = globalRegisterStates.MaximumRegisterStates
    leafs["warning-threshold"] = globalRegisterStates.WarningThreshold
    return leafs
}

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetYangName() string { return "global-register-states" }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) SetParent(parent types.Entity) { globalRegisterStates.parent = parent }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetParent() types.Entity { return globalRegisterStates.parent }

func (globalRegisterStates *Pim_DefaultContext_Ipv6_Maximum_GlobalRegisterStates) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetFilter() yfilter.YFilter { return globalRouteInterfaces.YFilter }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) SetFilter(yf yfilter.YFilter) { globalRouteInterfaces.YFilter = yf }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetSegmentPath() string {
    return "global-route-interfaces"
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = globalRouteInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = globalRouteInterfaces.WarningThreshold
    return leafs
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetYangName() string { return "global-route-interfaces" }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) SetParent(parent types.Entity) { globalRouteInterfaces.parent = parent }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetParent() types.Entity { return globalRouteInterfaces.parent }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv6_Maximum_GlobalRouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetFilter() yfilter.YFilter { return groupMappingsAutoRp.YFilter }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { groupMappingsAutoRp.YFilter = yf }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-group-ranges-auto-rp" { return "MaximumGroupRangesAutoRp" }
    if yname == "threshold-group-ranges-auto-rp" { return "ThresholdGroupRangesAutoRp" }
    return ""
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetSegmentPath() string {
    return "group-mappings-auto-rp"
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-group-ranges-auto-rp"] = groupMappingsAutoRp.MaximumGroupRangesAutoRp
    leafs["threshold-group-ranges-auto-rp"] = groupMappingsAutoRp.ThresholdGroupRangesAutoRp
    return leafs
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetYangName() string { return "group-mappings-auto-rp" }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) SetParent(parent types.Entity) { groupMappingsAutoRp.parent = parent }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetParent() types.Entity { return groupMappingsAutoRp.parent }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv6_Maximum_GroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetFilter() yfilter.YFilter { return bsrGroupMappings.YFilter }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGroupMappings.YFilter = yf }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-group-ranges" { return "BsrMaximumGroupRanges" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetSegmentPath() string {
    return "bsr-group-mappings"
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-group-ranges"] = bsrGroupMappings.BsrMaximumGroupRanges
    leafs["warning-threshold"] = bsrGroupMappings.WarningThreshold
    return leafs
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetYangName() string { return "bsr-group-mappings" }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) SetParent(parent types.Entity) { bsrGroupMappings.parent = parent }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetParent() types.Entity { return bsrGroupMappings.parent }

func (bsrGroupMappings *Pim_DefaultContext_Ipv6_Maximum_BsrGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_RegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetFilter() yfilter.YFilter { return registerStates.YFilter }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) SetFilter(yf yfilter.YFilter) { registerStates.YFilter = yf }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetSegmentPath() string {
    return "register-states"
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = registerStates.MaximumRegisterStates
    leafs["warning-threshold"] = registerStates.WarningThreshold
    return leafs
}

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetYangName() string { return "register-states" }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) SetParent(parent types.Entity) { registerStates.parent = parent }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetParent() types.Entity { return registerStates.parent }

func (registerStates *Pim_DefaultContext_Ipv6_Maximum_RegisterStates) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetFilter() yfilter.YFilter { return routeInterfaces.YFilter }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) SetFilter(yf yfilter.YFilter) { routeInterfaces.YFilter = yf }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetSegmentPath() string {
    return "route-interfaces"
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = routeInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = routeInterfaces.WarningThreshold
    return leafs
}

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetYangName() string { return "route-interfaces" }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) SetParent(parent types.Entity) { routeInterfaces.parent = parent }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetParent() types.Entity { return routeInterfaces.parent }

func (routeInterfaces *Pim_DefaultContext_Ipv6_Maximum_RouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetFilter() yfilter.YFilter { return bsrCandidateRpCache.YFilter }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrCandidateRpCache.YFilter = yf }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-candidate-rp-cache" { return "BsrMaximumCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetSegmentPath() string {
    return "bsr-candidate-rp-cache"
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-candidate-rp-cache"] = bsrCandidateRpCache.BsrMaximumCandidateRpCache
    leafs["warning-threshold"] = bsrCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetYangName() string { return "bsr-candidate-rp-cache" }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) SetParent(parent types.Entity) { bsrCandidateRpCache.parent = parent }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetParent() types.Entity { return bsrCandidateRpCache.parent }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv6_Maximum_BsrCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Maximum_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = routes.MaximumRoutes
    leafs["warning-threshold"] = routes.WarningThreshold
    return leafs
}

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetYangName() string { return "routes" }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Pim_DefaultContext_Ipv6_Maximum_Routes) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv6_Ssm
// Configure IP Multicast SSM
type Pim_DefaultContext_Ipv6_Ssm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetFilter() yfilter.YFilter { return ssm.YFilter }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) SetFilter(yf yfilter.YFilter) { ssm.YFilter = yf }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "range" { return "Range" }
    return ""
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetSegmentPath() string {
    return "ssm"
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = ssm.Disable
    leafs["range"] = ssm.Range
    return leafs
}

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetBundleName() string { return "cisco_ios_xr" }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetYangName() string { return "ssm" }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) SetParent(parent types.Entity) { ssm.parent = parent }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetParent() types.Entity { return ssm.parent }

func (ssm *Pim_DefaultContext_Ipv6_Ssm) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_DefaultContext_Ipv6_BidirRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetFilter() yfilter.YFilter { return bidirRpAddresses.YFilter }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) SetFilter(yf yfilter.YFilter) { bidirRpAddresses.YFilter = yf }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetGoName(yname string) string {
    if yname == "bidir-rp-address" { return "BidirRpAddress" }
    return ""
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetSegmentPath() string {
    return "bidir-rp-addresses"
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bidir-rp-address" {
        for _, c := range bidirRpAddresses.BidirRpAddress {
            if bidirRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress{}
        bidirRpAddresses.BidirRpAddress = append(bidirRpAddresses.BidirRpAddress, child)
        return &bidirRpAddresses.BidirRpAddress[len(bidirRpAddresses.BidirRpAddress)-1]
    }
    return nil
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bidirRpAddresses.BidirRpAddress {
        children[bidirRpAddresses.BidirRpAddress[i].GetSegmentPath()] = &bidirRpAddresses.BidirRpAddress[i]
    }
    return children
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetYangName() string { return "bidir-rp-addresses" }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) SetParent(parent types.Entity) { bidirRpAddresses.parent = parent }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetParent() types.Entity { return bidirRpAddresses.parent }

func (bidirRpAddresses *Pim_DefaultContext_Ipv6_BidirRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetFilter() yfilter.YFilter { return bidirRpAddress.YFilter }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) SetFilter(yf yfilter.YFilter) { bidirRpAddress.YFilter = yf }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetSegmentPath() string {
    return "bidir-rp-address" + "[rp-address='" + fmt.Sprintf("%v", bidirRpAddress.RpAddress) + "']"
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = bidirRpAddress.RpAddress
    leafs["access-list-name"] = bidirRpAddress.AccessListName
    leafs["auto-rp-override"] = bidirRpAddress.AutoRpOverride
    return leafs
}

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetYangName() string { return "bidir-rp-address" }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) SetParent(parent types.Entity) { bidirRpAddress.parent = parent }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetParent() types.Entity { return bidirRpAddress.parent }

func (bidirRpAddress *Pim_DefaultContext_Ipv6_BidirRpAddresses_BidirRpAddress) GetParentYangName() string { return "bidir-rp-addresses" }

// Pim_DefaultContext_Ipv6_Bsr
// PIM BSR configuration
type Pim_DefaultContext_Ipv6_Bsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_DefaultContext_Ipv6_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_DefaultContext_Ipv6_Bsr_CandidateRps
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetFilter() yfilter.YFilter { return bsr.YFilter }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) SetFilter(yf yfilter.YFilter) { bsr.YFilter = yf }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetGoName(yname string) string {
    if yname == "candidate-bsr" { return "CandidateBsr" }
    if yname == "candidate-rps" { return "CandidateRps" }
    return ""
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetSegmentPath() string {
    return "bsr"
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-bsr" {
        return &bsr.CandidateBsr
    }
    if childYangName == "candidate-rps" {
        return &bsr.CandidateRps
    }
    return nil
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-bsr"] = &bsr.CandidateBsr
    children["candidate-rps"] = &bsr.CandidateRps
    return children
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetBundleName() string { return "cisco_ios_xr" }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetYangName() string { return "bsr" }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) SetParent(parent types.Entity) { bsr.parent = parent }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetParent() types.Entity { return bsr.parent }

func (bsr *Pim_DefaultContext_Ipv6_Bsr) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_DefaultContext_Ipv6_Bsr_CandidateBsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BSR Address configured. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..128. The default value is 126.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetFilter() yfilter.YFilter { return candidateBsr.YFilter }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) SetFilter(yf yfilter.YFilter) { candidateBsr.YFilter = yf }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetSegmentPath() string {
    return "candidate-bsr"
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateBsr.Address
    leafs["prefix-length"] = candidateBsr.PrefixLength
    leafs["priority"] = candidateBsr.Priority
    return leafs
}

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetBundleName() string { return "cisco_ios_xr" }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetYangName() string { return "candidate-bsr" }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) SetParent(parent types.Entity) { candidateBsr.parent = parent }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetParent() types.Entity { return candidateBsr.parent }

func (candidateBsr *Pim_DefaultContext_Ipv6_Bsr_CandidateBsr) GetParentYangName() string { return "bsr" }

// Pim_DefaultContext_Ipv6_Bsr_CandidateRps
// PIM RP configuration
type Pim_DefaultContext_Ipv6_Bsr_CandidateRps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp.
    CandidateRp []Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetFilter() yfilter.YFilter { return candidateRps.YFilter }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) SetFilter(yf yfilter.YFilter) { candidateRps.YFilter = yf }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    return ""
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetSegmentPath() string {
    return "candidate-rps"
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        for _, c := range candidateRps.CandidateRp {
            if candidateRps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp{}
        candidateRps.CandidateRp = append(candidateRps.CandidateRp, child)
        return &candidateRps.CandidateRp[len(candidateRps.CandidateRp)-1]
    }
    return nil
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateRps.CandidateRp {
        children[candidateRps.CandidateRp[i].GetSegmentPath()] = &candidateRps.CandidateRp[i]
    }
    return children
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetYangName() string { return "candidate-rps" }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) SetParent(parent types.Entity) { candidateRps.parent = parent }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetParent() types.Entity { return candidateRps.parent }

func (candidateRps *Pim_DefaultContext_Ipv6_Bsr_CandidateRps) GetParentYangName() string { return "bsr" }

// Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mode" { return "Mode" }
    if yname == "group-list" { return "GroupList" }
    if yname == "priority" { return "Priority" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetSegmentPath() string {
    return "candidate-rp" + "[address='" + fmt.Sprintf("%v", candidateRp.Address) + "']" + "[mode='" + fmt.Sprintf("%v", candidateRp.Mode) + "']"
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateRp.Address
    leafs["mode"] = candidateRp.Mode
    leafs["group-list"] = candidateRp.GroupList
    leafs["priority"] = candidateRp.Priority
    leafs["interval"] = candidateRp.Interval
    return leafs
}

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *Pim_DefaultContext_Ipv6_Bsr_CandidateRps_CandidateRp) GetParentYangName() string { return "candidate-rps" }

// Pim_DefaultContext_Ipv6_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_DefaultContext_Ipv6_AllowRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetFilter() yfilter.YFilter { return allowRp.YFilter }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) SetFilter(yf yfilter.YFilter) { allowRp.YFilter = yf }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetGoName(yname string) string {
    if yname == "rp-list-name" { return "RpListName" }
    if yname == "group-list-name" { return "GroupListName" }
    return ""
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetSegmentPath() string {
    return "allow-rp"
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-list-name"] = allowRp.RpListName
    leafs["group-list-name"] = allowRp.GroupListName
    return leafs
}

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetBundleName() string { return "cisco_ios_xr" }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetYangName() string { return "allow-rp" }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) SetParent(parent types.Entity) { allowRp.parent = parent }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetParent() types.Entity { return allowRp.parent }

func (allowRp *Pim_DefaultContext_Ipv6_AllowRp) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_EmbeddedRpAddresses
// Set Embedded RP processing support
type Pim_DefaultContext_Ipv6_EmbeddedRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set Embedded RP processing support. The type is slice of
    // Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress.
    EmbeddedRpAddress []Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetFilter() yfilter.YFilter { return embeddedRpAddresses.YFilter }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) SetFilter(yf yfilter.YFilter) { embeddedRpAddresses.YFilter = yf }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetGoName(yname string) string {
    if yname == "embedded-rp-address" { return "EmbeddedRpAddress" }
    return ""
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetSegmentPath() string {
    return "embedded-rp-addresses"
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "embedded-rp-address" {
        for _, c := range embeddedRpAddresses.EmbeddedRpAddress {
            if embeddedRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress{}
        embeddedRpAddresses.EmbeddedRpAddress = append(embeddedRpAddresses.EmbeddedRpAddress, child)
        return &embeddedRpAddresses.EmbeddedRpAddress[len(embeddedRpAddresses.EmbeddedRpAddress)-1]
    }
    return nil
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range embeddedRpAddresses.EmbeddedRpAddress {
        children[embeddedRpAddresses.EmbeddedRpAddress[i].GetSegmentPath()] = &embeddedRpAddresses.EmbeddedRpAddress[i]
    }
    return children
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetYangName() string { return "embedded-rp-addresses" }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) SetParent(parent types.Entity) { embeddedRpAddresses.parent = parent }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetParent() types.Entity { return embeddedRpAddresses.parent }

func (embeddedRpAddresses *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress
// Set Embedded RP processing support
type Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of the Rendezvous Point. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64. This attribute is mandatory.
    AccessListName interface{}
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetFilter() yfilter.YFilter { return embeddedRpAddress.YFilter }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) SetFilter(yf yfilter.YFilter) { embeddedRpAddress.YFilter = yf }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetSegmentPath() string {
    return "embedded-rp-address" + "[rp-address='" + fmt.Sprintf("%v", embeddedRpAddress.RpAddress) + "']"
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = embeddedRpAddress.RpAddress
    leafs["access-list-name"] = embeddedRpAddress.AccessListName
    return leafs
}

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetYangName() string { return "embedded-rp-address" }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) SetParent(parent types.Entity) { embeddedRpAddress.parent = parent }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetParent() types.Entity { return embeddedRpAddress.parent }

func (embeddedRpAddress *Pim_DefaultContext_Ipv6_EmbeddedRpAddresses_EmbeddedRpAddress) GetParentYangName() string { return "embedded-rp-addresses" }

// Pim_DefaultContext_Ipv6_Convergence
// Configure convergence parameters
type Pim_DefaultContext_Ipv6_Convergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetFilter() yfilter.YFilter { return convergence.YFilter }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) SetFilter(yf yfilter.YFilter) { convergence.YFilter = yf }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetGoName(yname string) string {
    if yname == "rpf-conflict-join-delay" { return "RpfConflictJoinDelay" }
    if yname == "link-down-prune-delay" { return "LinkDownPruneDelay" }
    return ""
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetSegmentPath() string {
    return "convergence"
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rpf-conflict-join-delay"] = convergence.RpfConflictJoinDelay
    leafs["link-down-prune-delay"] = convergence.LinkDownPruneDelay
    return leafs
}

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetBundleName() string { return "cisco_ios_xr" }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetYangName() string { return "convergence" }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) SetParent(parent types.Entity) { convergence.parent = parent }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetParent() types.Entity { return convergence.parent }

func (convergence *Pim_DefaultContext_Ipv6_Convergence) GetParentYangName() string { return "ipv6" }

// Pim_DefaultContext_Ipv4
// IPV4 commands
type Pim_DefaultContext_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable PIM neighbor checking when receiving PIM messages. The type is
    // interface{}.
    NeighborCheckOnReceive interface{}

    // Generate registers compatible with older IOS versions. The type is
    // interface{}.
    OldRegisterChecksum interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Configure threshold of infinity for switching to SPT on last-hop. The type
    // is string.
    SptThresholdInfinity interface{}

    // PIM neighbor state change logging is turned on if configured. The type is
    // interface{}.
    LogNeighborChanges interface{}

    // Source address to use for register messages. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    RegisterSource interface{}

    // Access-list which specifies unauthorized sources. The type is string with
    // length: 1..64.
    AcceptRegister interface{}

    // Suppress prunes triggered as a result of RPF changes. The type is
    // interface{}.
    SuppressRpfPrunes interface{}

    // Allow SSM ranges to be overridden by more specific ranges. The type is
    // interface{}.
    SsmAllowOverride interface{}

    // Enable equal-cost multipath routing. The type is PimMultipath.
    Multipath interface{}

    // Configure static RP deny range. The type is string with length: 1..64.
    RpStaticDeny interface{}

    // Suppress data registers after initial state setup. The type is interface{}.
    SuppressDataRegisters interface{}

    // Enable PIM neighbor checking when sending join-prunes. The type is
    // interface{}.
    NeighborCheckOnSend interface{}

    // Disable Rendezvous Point discovery through the AutoRP protocol. The type is
    // interface{}.
    AutoRpDisable interface{}

    // Configure RPF-redirect feature.
    RpfRedirect Pim_DefaultContext_Ipv4_RpfRedirect

    // Interface-level Configuration.
    Interfaces Pim_DefaultContext_Ipv4_Interfaces

    // Configure Candidate-RPs.
    AutoRpCandidateRps Pim_DefaultContext_Ipv4_AutoRpCandidateRps

    // Configure AutoRP Mapping Agent.
    AutoRpMappingAgent Pim_DefaultContext_Ipv4_AutoRpMappingAgent

    // Configure Sparse-Mode Rendezvous Point.
    SparseModeRpAddresses Pim_DefaultContext_Ipv4_SparseModeRpAddresses

    // Inheritable defaults.
    InheritableDefaults Pim_DefaultContext_Ipv4_InheritableDefaults

    // Configure RPF options.
    Rpf Pim_DefaultContext_Ipv4_Rpf

    // Configure expiry timer for S,G routes.
    SgExpiryTimer Pim_DefaultContext_Ipv4_SgExpiryTimer

    // Enable PIM RPF Vector Proxy's.
    RpfVectorEnable Pim_DefaultContext_Ipv4_RpfVectorEnable

    // Configure Non-stop forwarding (NSF) options.
    Nsf Pim_DefaultContext_Ipv4_Nsf

    // Configure PIM State Limits.
    Maximum Pim_DefaultContext_Ipv4_Maximum

    // Configure IP Multicast SSM.
    Ssm Pim_DefaultContext_Ipv4_Ssm

    // Inject Explicit PIM RPF Vector Proxy's.
    Injects Pim_DefaultContext_Ipv4_Injects

    // Configure Bidirectional PIM Rendezvous Point.
    BidirRpAddresses Pim_DefaultContext_Ipv4_BidirRpAddresses

    // PIM BSR configuration.
    Bsr Pim_DefaultContext_Ipv4_Bsr

    // Multicast Only Fast Re-Route.
    Mofrr Pim_DefaultContext_Ipv4_Mofrr

    // Inject PIM RPF Vector Proxy's.
    Paths Pim_DefaultContext_Ipv4_Paths

    // Enable allow-rp filtering for SM joins.
    AllowRp Pim_DefaultContext_Ipv4_AllowRp

    // Configure convergence parameters.
    Convergence Pim_DefaultContext_Ipv4_Convergence
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Pim_DefaultContext_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Pim_DefaultContext_Ipv4) GetGoName(yname string) string {
    if yname == "neighbor-check-on-receive" { return "NeighborCheckOnReceive" }
    if yname == "old-register-checksum" { return "OldRegisterChecksum" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "spt-threshold-infinity" { return "SptThresholdInfinity" }
    if yname == "log-neighbor-changes" { return "LogNeighborChanges" }
    if yname == "register-source" { return "RegisterSource" }
    if yname == "accept-register" { return "AcceptRegister" }
    if yname == "suppress-rpf-prunes" { return "SuppressRpfPrunes" }
    if yname == "ssm-allow-override" { return "SsmAllowOverride" }
    if yname == "multipath" { return "Multipath" }
    if yname == "rp-static-deny" { return "RpStaticDeny" }
    if yname == "suppress-data-registers" { return "SuppressDataRegisters" }
    if yname == "neighbor-check-on-send" { return "NeighborCheckOnSend" }
    if yname == "auto-rp-disable" { return "AutoRpDisable" }
    if yname == "rpf-redirect" { return "RpfRedirect" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "auto-rp-candidate-rps" { return "AutoRpCandidateRps" }
    if yname == "auto-rp-mapping-agent" { return "AutoRpMappingAgent" }
    if yname == "sparse-mode-rp-addresses" { return "SparseModeRpAddresses" }
    if yname == "inheritable-defaults" { return "InheritableDefaults" }
    if yname == "rpf" { return "Rpf" }
    if yname == "sg-expiry-timer" { return "SgExpiryTimer" }
    if yname == "rpf-vector-enable" { return "RpfVectorEnable" }
    if yname == "nsf" { return "Nsf" }
    if yname == "maximum" { return "Maximum" }
    if yname == "ssm" { return "Ssm" }
    if yname == "injects" { return "Injects" }
    if yname == "bidir-rp-addresses" { return "BidirRpAddresses" }
    if yname == "bsr" { return "Bsr" }
    if yname == "mofrr" { return "Mofrr" }
    if yname == "paths" { return "Paths" }
    if yname == "allow-rp" { return "AllowRp" }
    if yname == "convergence" { return "Convergence" }
    return ""
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rpf-redirect" {
        return &ipv4.RpfRedirect
    }
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    if childYangName == "auto-rp-candidate-rps" {
        return &ipv4.AutoRpCandidateRps
    }
    if childYangName == "auto-rp-mapping-agent" {
        return &ipv4.AutoRpMappingAgent
    }
    if childYangName == "sparse-mode-rp-addresses" {
        return &ipv4.SparseModeRpAddresses
    }
    if childYangName == "inheritable-defaults" {
        return &ipv4.InheritableDefaults
    }
    if childYangName == "rpf" {
        return &ipv4.Rpf
    }
    if childYangName == "sg-expiry-timer" {
        return &ipv4.SgExpiryTimer
    }
    if childYangName == "rpf-vector-enable" {
        return &ipv4.RpfVectorEnable
    }
    if childYangName == "nsf" {
        return &ipv4.Nsf
    }
    if childYangName == "maximum" {
        return &ipv4.Maximum
    }
    if childYangName == "ssm" {
        return &ipv4.Ssm
    }
    if childYangName == "injects" {
        return &ipv4.Injects
    }
    if childYangName == "bidir-rp-addresses" {
        return &ipv4.BidirRpAddresses
    }
    if childYangName == "bsr" {
        return &ipv4.Bsr
    }
    if childYangName == "mofrr" {
        return &ipv4.Mofrr
    }
    if childYangName == "paths" {
        return &ipv4.Paths
    }
    if childYangName == "allow-rp" {
        return &ipv4.AllowRp
    }
    if childYangName == "convergence" {
        return &ipv4.Convergence
    }
    return nil
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rpf-redirect"] = &ipv4.RpfRedirect
    children["interfaces"] = &ipv4.Interfaces
    children["auto-rp-candidate-rps"] = &ipv4.AutoRpCandidateRps
    children["auto-rp-mapping-agent"] = &ipv4.AutoRpMappingAgent
    children["sparse-mode-rp-addresses"] = &ipv4.SparseModeRpAddresses
    children["inheritable-defaults"] = &ipv4.InheritableDefaults
    children["rpf"] = &ipv4.Rpf
    children["sg-expiry-timer"] = &ipv4.SgExpiryTimer
    children["rpf-vector-enable"] = &ipv4.RpfVectorEnable
    children["nsf"] = &ipv4.Nsf
    children["maximum"] = &ipv4.Maximum
    children["ssm"] = &ipv4.Ssm
    children["injects"] = &ipv4.Injects
    children["bidir-rp-addresses"] = &ipv4.BidirRpAddresses
    children["bsr"] = &ipv4.Bsr
    children["mofrr"] = &ipv4.Mofrr
    children["paths"] = &ipv4.Paths
    children["allow-rp"] = &ipv4.AllowRp
    children["convergence"] = &ipv4.Convergence
    return children
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-check-on-receive"] = ipv4.NeighborCheckOnReceive
    leafs["old-register-checksum"] = ipv4.OldRegisterChecksum
    leafs["neighbor-filter"] = ipv4.NeighborFilter
    leafs["spt-threshold-infinity"] = ipv4.SptThresholdInfinity
    leafs["log-neighbor-changes"] = ipv4.LogNeighborChanges
    leafs["register-source"] = ipv4.RegisterSource
    leafs["accept-register"] = ipv4.AcceptRegister
    leafs["suppress-rpf-prunes"] = ipv4.SuppressRpfPrunes
    leafs["ssm-allow-override"] = ipv4.SsmAllowOverride
    leafs["multipath"] = ipv4.Multipath
    leafs["rp-static-deny"] = ipv4.RpStaticDeny
    leafs["suppress-data-registers"] = ipv4.SuppressDataRegisters
    leafs["neighbor-check-on-send"] = ipv4.NeighborCheckOnSend
    leafs["auto-rp-disable"] = ipv4.AutoRpDisable
    return leafs
}

func (ipv4 *Pim_DefaultContext_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Pim_DefaultContext_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Pim_DefaultContext_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Pim_DefaultContext_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Pim_DefaultContext_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Pim_DefaultContext_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Pim_DefaultContext_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Pim_DefaultContext_Ipv4) GetParentYangName() string { return "default-context" }

// Pim_DefaultContext_Ipv4_RpfRedirect
// Configure RPF-redirect feature
type Pim_DefaultContext_Ipv4_RpfRedirect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetFilter() yfilter.YFilter { return rpfRedirect.YFilter }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) SetFilter(yf yfilter.YFilter) { rpfRedirect.YFilter = yf }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetSegmentPath() string {
    return "rpf-redirect"
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy"] = rpfRedirect.RoutePolicy
    return leafs
}

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetBundleName() string { return "cisco_ios_xr" }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetYangName() string { return "rpf-redirect" }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) SetParent(parent types.Entity) { rpfRedirect.parent = parent }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetParent() types.Entity { return rpfRedirect.parent }

func (rpfRedirect *Pim_DefaultContext_Ipv4_RpfRedirect) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Interfaces
// Interface-level Configuration
type Pim_DefaultContext_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Pim_DefaultContext_Ipv4_Interfaces_Interface.
    Interface []Pim_DefaultContext_Ipv4_Interfaces_Interface
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Pim_DefaultContext_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Interfaces_Interface
// The name of the interface
type Pim_DefaultContext_Ipv4_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enter PIM Interface processing. The type is interface{}.
    Enable interface{}

    // Access-list of neighbors to be filtered. The type is string with length:
    // 1..64.
    NeighborFilter interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // BSR Border configuration for Interface. The type is bool.
    BsrBorder interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Enable PIM processing on the interface. The type is bool.
    InterfaceEnable interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}

    // Configure RPF-redirect bundle for interface. Applicable for IPv4 only.
    RedirectBundle Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle

    // Maximum number of allowed routes for this interface.
    MaximumRoutes Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes

    // BFD configuration.
    Bfd Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "enable" { return "Enable" }
    if yname == "neighbor-filter" { return "NeighborFilter" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "bsr-border" { return "BsrBorder" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "interface-enable" { return "InterfaceEnable" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    if yname == "redirect-bundle" { return "RedirectBundle" }
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "bfd" { return "Bfd" }
    return ""
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redirect-bundle" {
        return &self.RedirectBundle
    }
    if childYangName == "maximum-routes" {
        return &self.MaximumRoutes
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    return nil
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["redirect-bundle"] = &self.RedirectBundle
    children["maximum-routes"] = &self.MaximumRoutes
    children["bfd"] = &self.Bfd
    return children
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["enable"] = self.Enable
    leafs["neighbor-filter"] = self.NeighborFilter
    leafs["hello-interval"] = self.HelloInterval
    leafs["bsr-border"] = self.BsrBorder
    leafs["propagation-delay"] = self.PropagationDelay
    leafs["dr-priority"] = self.DrPriority
    leafs["join-prune-mtu"] = self.JoinPruneMtu
    leafs["interface-enable"] = self.InterfaceEnable
    leafs["jp-interval"] = self.JpInterval
    leafs["override-interval"] = self.OverrideInterval
    return leafs
}

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Pim_DefaultContext_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle
// Configure RPF-redirect bundle for interface.
// Applicable for IPv4 only
type Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bundle name. The type is string with length: 1..32.
    BundleName interface{}

    // Interface bandwidth in Kbps. The type is interface{} with range:
    // 0..100000000. Units are kbit/s.
    InterfaceBandwidth interface{}

    // Threshold bandwidth in Kbps. The type is interface{} with range:
    // 0..100000000. Units are kbit/s.
    ThresholdBandwidth interface{}
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetFilter() yfilter.YFilter { return redirectBundle.YFilter }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) SetFilter(yf yfilter.YFilter) { redirectBundle.YFilter = yf }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetGoName(yname string) string {
    if yname == "bundle-name" { return "BundleName" }
    if yname == "interface-bandwidth" { return "InterfaceBandwidth" }
    if yname == "threshold-bandwidth" { return "ThresholdBandwidth" }
    return ""
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetSegmentPath() string {
    return "redirect-bundle"
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-name"] = redirectBundle.BundleName
    leafs["interface-bandwidth"] = redirectBundle.InterfaceBandwidth
    leafs["threshold-bandwidth"] = redirectBundle.ThresholdBandwidth
    return leafs
}

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetBundleName() string { return "cisco_ios_xr" }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetYangName() string { return "redirect-bundle" }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) SetParent(parent types.Entity) { redirectBundle.parent = parent }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetParent() types.Entity { return redirectBundle.parent }

func (redirectBundle *Pim_DefaultContext_Ipv4_Interfaces_Interface_RedirectBundle) GetParentYangName() string { return "interface" }

// Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes
// Maximum number of allowed routes for this
// interface
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of routes for this interface. The type is interface{} with
    // range: 1..1100000. This attribute is mandatory.
    Maximum interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000.
    WarningThreshold interface{}

    // Access-list to account for. The type is string with length: 1..64.
    AccessListName interface{}
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetFilter() yfilter.YFilter { return maximumRoutes.YFilter }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) SetFilter(yf yfilter.YFilter) { maximumRoutes.YFilter = yf }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetSegmentPath() string {
    return "maximum-routes"
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = maximumRoutes.Maximum
    leafs["warning-threshold"] = maximumRoutes.WarningThreshold
    leafs["access-list-name"] = maximumRoutes.AccessListName
    return leafs
}

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetYangName() string { return "maximum-routes" }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) SetParent(parent types.Entity) { maximumRoutes.parent = parent }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetParent() types.Entity { return maximumRoutes.parent }

func (maximumRoutes *Pim_DefaultContext_Ipv4_Interfaces_Interface_MaximumRoutes) GetParentYangName() string { return "interface" }

// Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd
// BFD configuration
type Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by PIM. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by PIM. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // TRUE to enable BFD. FALSE to disable and to prevent inheritance from a
    // parent. The type is bool.
    Enable interface{}
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    if yname == "interval" { return "Interval" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    leafs["interval"] = bfd.Interval
    leafs["enable"] = bfd.Enable
    return leafs
}

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Pim_DefaultContext_Ipv4_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Pim_DefaultContext_Ipv4_AutoRpCandidateRps
// Configure Candidate-RPs
type Pim_DefaultContext_Ipv4_AutoRpCandidateRps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifications for a Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp.
    AutoRpCandidateRp []Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetFilter() yfilter.YFilter { return autoRpCandidateRps.YFilter }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) SetFilter(yf yfilter.YFilter) { autoRpCandidateRps.YFilter = yf }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetGoName(yname string) string {
    if yname == "auto-rp-candidate-rp" { return "AutoRpCandidateRp" }
    return ""
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetSegmentPath() string {
    return "auto-rp-candidate-rps"
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "auto-rp-candidate-rp" {
        for _, c := range autoRpCandidateRps.AutoRpCandidateRp {
            if autoRpCandidateRps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp{}
        autoRpCandidateRps.AutoRpCandidateRp = append(autoRpCandidateRps.AutoRpCandidateRp, child)
        return &autoRpCandidateRps.AutoRpCandidateRp[len(autoRpCandidateRps.AutoRpCandidateRp)-1]
    }
    return nil
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range autoRpCandidateRps.AutoRpCandidateRp {
        children[autoRpCandidateRps.AutoRpCandidateRp[i].GetSegmentPath()] = &autoRpCandidateRps.AutoRpCandidateRp[i]
    }
    return children
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetBundleName() string { return "cisco_ios_xr" }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetYangName() string { return "auto-rp-candidate-rps" }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) SetParent(parent types.Entity) { autoRpCandidateRps.parent = parent }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetParent() types.Entity { return autoRpCandidateRps.parent }

func (autoRpCandidateRps *Pim_DefaultContext_Ipv4_AutoRpCandidateRps) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp
// Specifications for a Candidate-RP
type Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface from which Candidate-RP packets will be
    // sourced. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // This attribute is a key. Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // TTL in Hops. The type is interface{} with range: 1..255. This attribute is
    // mandatory.
    Ttl interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64. The default value is 224-4.
    AccessListName interface{}

    // Time between announcements <in seconds> . The type is interface{} with
    // range: 1..600. Units are second. The default value is 60.
    AnnouncePeriod interface{}
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetFilter() yfilter.YFilter { return autoRpCandidateRp.YFilter }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) SetFilter(yf yfilter.YFilter) { autoRpCandidateRp.YFilter = yf }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "protocol-mode" { return "ProtocolMode" }
    if yname == "ttl" { return "Ttl" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "announce-period" { return "AnnouncePeriod" }
    return ""
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetSegmentPath() string {
    return "auto-rp-candidate-rp" + "[interface-name='" + fmt.Sprintf("%v", autoRpCandidateRp.InterfaceName) + "']" + "[protocol-mode='" + fmt.Sprintf("%v", autoRpCandidateRp.ProtocolMode) + "']"
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = autoRpCandidateRp.InterfaceName
    leafs["protocol-mode"] = autoRpCandidateRp.ProtocolMode
    leafs["ttl"] = autoRpCandidateRp.Ttl
    leafs["access-list-name"] = autoRpCandidateRp.AccessListName
    leafs["announce-period"] = autoRpCandidateRp.AnnouncePeriod
    return leafs
}

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetYangName() string { return "auto-rp-candidate-rp" }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) SetParent(parent types.Entity) { autoRpCandidateRp.parent = parent }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetParent() types.Entity { return autoRpCandidateRp.parent }

func (autoRpCandidateRp *Pim_DefaultContext_Ipv4_AutoRpCandidateRps_AutoRpCandidateRp) GetParentYangName() string { return "auto-rp-candidate-rps" }

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent
// Configure AutoRP Mapping Agent
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifications for Mapping Agent configured on this box.
    Parameters Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters

    // Mapping Agent cache size limit.
    CacheLimit Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetFilter() yfilter.YFilter { return autoRpMappingAgent.YFilter }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) SetFilter(yf yfilter.YFilter) { autoRpMappingAgent.YFilter = yf }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetGoName(yname string) string {
    if yname == "parameters" { return "Parameters" }
    if yname == "cache-limit" { return "CacheLimit" }
    return ""
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetSegmentPath() string {
    return "auto-rp-mapping-agent"
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "parameters" {
        return &autoRpMappingAgent.Parameters
    }
    if childYangName == "cache-limit" {
        return &autoRpMappingAgent.CacheLimit
    }
    return nil
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["parameters"] = &autoRpMappingAgent.Parameters
    children["cache-limit"] = &autoRpMappingAgent.CacheLimit
    return children
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetBundleName() string { return "cisco_ios_xr" }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetYangName() string { return "auto-rp-mapping-agent" }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) SetParent(parent types.Entity) { autoRpMappingAgent.parent = parent }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetParent() types.Entity { return autoRpMappingAgent.parent }

func (autoRpMappingAgent *Pim_DefaultContext_Ipv4_AutoRpMappingAgent) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters
// Specifications for Mapping Agent configured
// on this box
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface from which mapping packets will be sourced . The type is string
    // with pattern: [a-zA-Z0-9./-]+. This attribute is mandatory.
    InterfaceName interface{}

    // TTL in Hops. The type is interface{} with range: 1..255. This attribute is
    // mandatory.
    Ttl interface{}

    // Time between discovery messages <in seconds>. The type is interface{} with
    // range: 1..600. Units are second. The default value is 60.
    AnnouncePeriod interface{}
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetFilter() yfilter.YFilter { return parameters.YFilter }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) SetFilter(yf yfilter.YFilter) { parameters.YFilter = yf }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ttl" { return "Ttl" }
    if yname == "announce-period" { return "AnnouncePeriod" }
    return ""
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetSegmentPath() string {
    return "parameters"
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = parameters.InterfaceName
    leafs["ttl"] = parameters.Ttl
    leafs["announce-period"] = parameters.AnnouncePeriod
    return leafs
}

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetBundleName() string { return "cisco_ios_xr" }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetYangName() string { return "parameters" }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) SetParent(parent types.Entity) { parameters.parent = parent }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetParent() types.Entity { return parameters.parent }

func (parameters *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_Parameters) GetParentYangName() string { return "auto-rp-mapping-agent" }

// Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit
// Mapping Agent cache size limit
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of mapping cache entries. The type is interface{} with
    // range: 1..1000. This attribute is mandatory.
    MaximumCacheEntry interface{}

    // Warning threshold number of cache entries. The type is interface{} with
    // range: 1..1000. The default value is 450.
    ThresholdCacheEntry interface{}
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetFilter() yfilter.YFilter { return cacheLimit.YFilter }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) SetFilter(yf yfilter.YFilter) { cacheLimit.YFilter = yf }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetGoName(yname string) string {
    if yname == "maximum-cache-entry" { return "MaximumCacheEntry" }
    if yname == "threshold-cache-entry" { return "ThresholdCacheEntry" }
    return ""
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetSegmentPath() string {
    return "cache-limit"
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-cache-entry"] = cacheLimit.MaximumCacheEntry
    leafs["threshold-cache-entry"] = cacheLimit.ThresholdCacheEntry
    return leafs
}

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetBundleName() string { return "cisco_ios_xr" }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetYangName() string { return "cache-limit" }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) SetParent(parent types.Entity) { cacheLimit.parent = parent }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetParent() types.Entity { return cacheLimit.parent }

func (cacheLimit *Pim_DefaultContext_Ipv4_AutoRpMappingAgent_CacheLimit) GetParentYangName() string { return "auto-rp-mapping-agent" }

// Pim_DefaultContext_Ipv4_SparseModeRpAddresses
// Configure Sparse-Mode Rendezvous Point
type Pim_DefaultContext_Ipv4_SparseModeRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress.
    SparseModeRpAddress []Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetFilter() yfilter.YFilter { return sparseModeRpAddresses.YFilter }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) SetFilter(yf yfilter.YFilter) { sparseModeRpAddresses.YFilter = yf }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetGoName(yname string) string {
    if yname == "sparse-mode-rp-address" { return "SparseModeRpAddress" }
    return ""
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetSegmentPath() string {
    return "sparse-mode-rp-addresses"
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sparse-mode-rp-address" {
        for _, c := range sparseModeRpAddresses.SparseModeRpAddress {
            if sparseModeRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress{}
        sparseModeRpAddresses.SparseModeRpAddress = append(sparseModeRpAddresses.SparseModeRpAddress, child)
        return &sparseModeRpAddresses.SparseModeRpAddress[len(sparseModeRpAddresses.SparseModeRpAddress)-1]
    }
    return nil
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sparseModeRpAddresses.SparseModeRpAddress {
        children[sparseModeRpAddresses.SparseModeRpAddress[i].GetSegmentPath()] = &sparseModeRpAddresses.SparseModeRpAddress[i]
    }
    return children
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetYangName() string { return "sparse-mode-rp-addresses" }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) SetParent(parent types.Entity) { sparseModeRpAddresses.parent = parent }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetParent() types.Entity { return sparseModeRpAddresses.parent }

func (sparseModeRpAddresses *Pim_DefaultContext_Ipv4_SparseModeRpAddresses) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a  given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetFilter() yfilter.YFilter { return sparseModeRpAddress.YFilter }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) SetFilter(yf yfilter.YFilter) { sparseModeRpAddress.YFilter = yf }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetSegmentPath() string {
    return "sparse-mode-rp-address" + "[rp-address='" + fmt.Sprintf("%v", sparseModeRpAddress.RpAddress) + "']"
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = sparseModeRpAddress.RpAddress
    leafs["access-list-name"] = sparseModeRpAddress.AccessListName
    leafs["auto-rp-override"] = sparseModeRpAddress.AutoRpOverride
    return leafs
}

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetYangName() string { return "sparse-mode-rp-address" }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) SetParent(parent types.Entity) { sparseModeRpAddress.parent = parent }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetParent() types.Entity { return sparseModeRpAddress.parent }

func (sparseModeRpAddress *Pim_DefaultContext_Ipv4_SparseModeRpAddresses_SparseModeRpAddress) GetParentYangName() string { return "sparse-mode-rp-addresses" }

// Pim_DefaultContext_Ipv4_InheritableDefaults
// Inheritable defaults
type Pim_DefaultContext_Ipv4_InheritableDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Convergency timeout in seconds. The type is interface{} with range:
    // 1800..2400. Units are second.
    ConvergenceTimeout interface{}

    // Hello interval in seconds. The type is interface{} with range: 1..3600.
    // Units are second.
    HelloInterval interface{}

    // Propagation delay in milli seconds. The type is interface{} with range:
    // 100..32767. Units are millisecond.
    PropagationDelay interface{}

    // Hello DR priority, preference given to larger value. The type is
    // interface{} with range: 0..4294967295.
    DrPriority interface{}

    // Join-Prune MTU in Bytes. The type is interface{} with range: 576..65535.
    // Units are byte.
    JoinPruneMtu interface{}

    // Join-Prune interval in seconds. The type is interface{} with range:
    // 10..600. Units are second.
    JpInterval interface{}

    // Override interval in milliseconds. The type is interface{} with range:
    // 400..65535. Units are millisecond.
    OverrideInterval interface{}
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetFilter() yfilter.YFilter { return inheritableDefaults.YFilter }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) SetFilter(yf yfilter.YFilter) { inheritableDefaults.YFilter = yf }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetGoName(yname string) string {
    if yname == "convergence-timeout" { return "ConvergenceTimeout" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "propagation-delay" { return "PropagationDelay" }
    if yname == "dr-priority" { return "DrPriority" }
    if yname == "join-prune-mtu" { return "JoinPruneMtu" }
    if yname == "jp-interval" { return "JpInterval" }
    if yname == "override-interval" { return "OverrideInterval" }
    return ""
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetSegmentPath() string {
    return "inheritable-defaults"
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["convergence-timeout"] = inheritableDefaults.ConvergenceTimeout
    leafs["hello-interval"] = inheritableDefaults.HelloInterval
    leafs["propagation-delay"] = inheritableDefaults.PropagationDelay
    leafs["dr-priority"] = inheritableDefaults.DrPriority
    leafs["join-prune-mtu"] = inheritableDefaults.JoinPruneMtu
    leafs["jp-interval"] = inheritableDefaults.JpInterval
    leafs["override-interval"] = inheritableDefaults.OverrideInterval
    return leafs
}

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetYangName() string { return "inheritable-defaults" }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) SetParent(parent types.Entity) { inheritableDefaults.parent = parent }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetParent() types.Entity { return inheritableDefaults.parent }

func (inheritableDefaults *Pim_DefaultContext_Ipv4_InheritableDefaults) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Rpf
// Configure RPF options
type Pim_DefaultContext_Ipv4_Rpf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Route policy to select RPF topology. The type is string with length: 1..64.
    RoutePolicy interface{}
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetFilter() yfilter.YFilter { return rpf.YFilter }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) SetFilter(yf yfilter.YFilter) { rpf.YFilter = yf }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetGoName(yname string) string {
    if yname == "route-policy" { return "RoutePolicy" }
    return ""
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetSegmentPath() string {
    return "rpf"
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route-policy"] = rpf.RoutePolicy
    return leafs
}

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetBundleName() string { return "cisco_ios_xr" }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetYangName() string { return "rpf" }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) SetParent(parent types.Entity) { rpf.parent = parent }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetParent() types.Entity { return rpf.parent }

func (rpf *Pim_DefaultContext_Ipv4_Rpf) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_SgExpiryTimer
// Configure expiry timer for S,G routes
type Pim_DefaultContext_Ipv4_SgExpiryTimer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // (S,G) expiry time in seconds. The type is interface{} with range:
    // 40..57600. Units are second.
    Interval interface{}

    // Access-list of applicable S,G routes. The type is string with length:
    // 1..64.
    AccessListName interface{}
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetFilter() yfilter.YFilter { return sgExpiryTimer.YFilter }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) SetFilter(yf yfilter.YFilter) { sgExpiryTimer.YFilter = yf }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "access-list-name" { return "AccessListName" }
    return ""
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetSegmentPath() string {
    return "sg-expiry-timer"
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = sgExpiryTimer.Interval
    leafs["access-list-name"] = sgExpiryTimer.AccessListName
    return leafs
}

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetBundleName() string { return "cisco_ios_xr" }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetYangName() string { return "sg-expiry-timer" }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) SetParent(parent types.Entity) { sgExpiryTimer.parent = parent }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetParent() types.Entity { return sgExpiryTimer.parent }

func (sgExpiryTimer *Pim_DefaultContext_Ipv4_SgExpiryTimer) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_RpfVectorEnable
// Enable PIM RPF Vector Proxy's
// This type is a presence type.
type Pim_DefaultContext_Ipv4_RpfVectorEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF Vector is turned on if configured. The type is interface{}. This
    // attribute is mandatory.
    Enable interface{}

    // Allow RPF Vector origination over eBGP sessions. The type is interface{}.
    AllowEbgp interface{}

    // Disable RPF Vector origination over iBGP sessions. The type is interface{}.
    DisableIbgp interface{}
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetFilter() yfilter.YFilter { return rpfVectorEnable.YFilter }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) SetFilter(yf yfilter.YFilter) { rpfVectorEnable.YFilter = yf }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-ebgp" { return "AllowEbgp" }
    if yname == "disable-ibgp" { return "DisableIbgp" }
    return ""
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetSegmentPath() string {
    return "rpf-vector-enable"
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rpfVectorEnable.Enable
    leafs["allow-ebgp"] = rpfVectorEnable.AllowEbgp
    leafs["disable-ibgp"] = rpfVectorEnable.DisableIbgp
    return leafs
}

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetBundleName() string { return "cisco_ios_xr" }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetYangName() string { return "rpf-vector-enable" }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) SetParent(parent types.Entity) { rpfVectorEnable.parent = parent }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetParent() types.Entity { return rpfVectorEnable.parent }

func (rpfVectorEnable *Pim_DefaultContext_Ipv4_RpfVectorEnable) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Nsf
// Configure Non-stop forwarding (NSF) options
type Pim_DefaultContext_Ipv4_Nsf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Override default maximum lifetime for PIM NSF mode. The type is interface{}
    // with range: 10..600. Units are second.
    Lifetime interface{}
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetFilter() yfilter.YFilter { return nsf.YFilter }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) SetFilter(yf yfilter.YFilter) { nsf.YFilter = yf }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetGoName(yname string) string {
    if yname == "lifetime" { return "Lifetime" }
    return ""
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetSegmentPath() string {
    return "nsf"
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lifetime"] = nsf.Lifetime
    return leafs
}

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetBundleName() string { return "cisco_ios_xr" }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetYangName() string { return "nsf" }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) SetParent(parent types.Entity) { nsf.parent = parent }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetParent() types.Entity { return nsf.parent }

func (nsf *Pim_DefaultContext_Ipv4_Nsf) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Maximum
// Configure PIM State Limits
type Pim_DefaultContext_Ipv4_Maximum struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalLowPriorityPacketQueue interface{}

    // Maximum packet queue size in bytes. The type is interface{} with range:
    // 0..2147483648. Units are byte.
    GlobalHighPriorityPacketQueue interface{}

    // Override default global maximum and threshold for PIM group mapping ranges
    // from BSR.
    BsrGlobalGroupMappings Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings

    // Override default maximum for number of routes.
    GlobalRoutes Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes

    // Maximum for number of group mappings from autorp mapping agent.
    GlobalGroupMappingsAutoRp Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp

    // Override default global maximum and threshold for C-RP set in BSR.
    BsrGlobalCandidateRpCache Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache

    // Override default maximum for number of sparse-mode source registers.
    GlobalRegisterStates Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates

    // Override default maximum for number of route-interfaces.
    GlobalRouteInterfaces Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces

    // Override default maximum for number of group mappings from autorp mapping
    // agent.
    GroupMappingsAutoRp Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp

    // Override default maximum and threshold for number of group mappings from
    // BSR.
    BsrGroupMappings Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings

    // Override default maximum for number of sparse-mode source registers.
    RegisterStates Pim_DefaultContext_Ipv4_Maximum_RegisterStates

    // Override default maximum for number of route-interfaces.
    RouteInterfaces Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces

    // Override default maximum and threshold for BSR C-RP cache setting.
    BsrCandidateRpCache Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache

    // Override default maximum for number of routes.
    Routes Pim_DefaultContext_Ipv4_Maximum_Routes
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetFilter() yfilter.YFilter { return maximum.YFilter }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) SetFilter(yf yfilter.YFilter) { maximum.YFilter = yf }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetGoName(yname string) string {
    if yname == "global-low-priority-packet-queue" { return "GlobalLowPriorityPacketQueue" }
    if yname == "global-high-priority-packet-queue" { return "GlobalHighPriorityPacketQueue" }
    if yname == "bsr-global-group-mappings" { return "BsrGlobalGroupMappings" }
    if yname == "global-routes" { return "GlobalRoutes" }
    if yname == "global-group-mappings-auto-rp" { return "GlobalGroupMappingsAutoRp" }
    if yname == "bsr-global-candidate-rp-cache" { return "BsrGlobalCandidateRpCache" }
    if yname == "global-register-states" { return "GlobalRegisterStates" }
    if yname == "global-route-interfaces" { return "GlobalRouteInterfaces" }
    if yname == "group-mappings-auto-rp" { return "GroupMappingsAutoRp" }
    if yname == "bsr-group-mappings" { return "BsrGroupMappings" }
    if yname == "register-states" { return "RegisterStates" }
    if yname == "route-interfaces" { return "RouteInterfaces" }
    if yname == "bsr-candidate-rp-cache" { return "BsrCandidateRpCache" }
    if yname == "routes" { return "Routes" }
    return ""
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetSegmentPath() string {
    return "maximum"
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bsr-global-group-mappings" {
        return &maximum.BsrGlobalGroupMappings
    }
    if childYangName == "global-routes" {
        return &maximum.GlobalRoutes
    }
    if childYangName == "global-group-mappings-auto-rp" {
        return &maximum.GlobalGroupMappingsAutoRp
    }
    if childYangName == "bsr-global-candidate-rp-cache" {
        return &maximum.BsrGlobalCandidateRpCache
    }
    if childYangName == "global-register-states" {
        return &maximum.GlobalRegisterStates
    }
    if childYangName == "global-route-interfaces" {
        return &maximum.GlobalRouteInterfaces
    }
    if childYangName == "group-mappings-auto-rp" {
        return &maximum.GroupMappingsAutoRp
    }
    if childYangName == "bsr-group-mappings" {
        return &maximum.BsrGroupMappings
    }
    if childYangName == "register-states" {
        return &maximum.RegisterStates
    }
    if childYangName == "route-interfaces" {
        return &maximum.RouteInterfaces
    }
    if childYangName == "bsr-candidate-rp-cache" {
        return &maximum.BsrCandidateRpCache
    }
    if childYangName == "routes" {
        return &maximum.Routes
    }
    return nil
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bsr-global-group-mappings"] = &maximum.BsrGlobalGroupMappings
    children["global-routes"] = &maximum.GlobalRoutes
    children["global-group-mappings-auto-rp"] = &maximum.GlobalGroupMappingsAutoRp
    children["bsr-global-candidate-rp-cache"] = &maximum.BsrGlobalCandidateRpCache
    children["global-register-states"] = &maximum.GlobalRegisterStates
    children["global-route-interfaces"] = &maximum.GlobalRouteInterfaces
    children["group-mappings-auto-rp"] = &maximum.GroupMappingsAutoRp
    children["bsr-group-mappings"] = &maximum.BsrGroupMappings
    children["register-states"] = &maximum.RegisterStates
    children["route-interfaces"] = &maximum.RouteInterfaces
    children["bsr-candidate-rp-cache"] = &maximum.BsrCandidateRpCache
    children["routes"] = &maximum.Routes
    return children
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-low-priority-packet-queue"] = maximum.GlobalLowPriorityPacketQueue
    leafs["global-high-priority-packet-queue"] = maximum.GlobalHighPriorityPacketQueue
    return leafs
}

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetBundleName() string { return "cisco_ios_xr" }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetYangName() string { return "maximum" }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) SetParent(parent types.Entity) { maximum.parent = parent }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetParent() types.Entity { return maximum.parent }

func (maximum *Pim_DefaultContext_Ipv4_Maximum) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings
// Override default global maximum and threshold
// for PIM group mapping ranges from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global Maximum number of PIM group mapping ranges from BSR. The type is
    // interface{} with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalGroupMappings interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetFilter() yfilter.YFilter { return bsrGlobalGroupMappings.YFilter }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGlobalGroupMappings.YFilter = yf }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-global-group-mappings" { return "BsrMaximumGlobalGroupMappings" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetSegmentPath() string {
    return "bsr-global-group-mappings"
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-global-group-mappings"] = bsrGlobalGroupMappings.BsrMaximumGlobalGroupMappings
    leafs["warning-threshold"] = bsrGlobalGroupMappings.WarningThreshold
    return leafs
}

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetYangName() string { return "bsr-global-group-mappings" }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) SetParent(parent types.Entity) { bsrGlobalGroupMappings.parent = parent }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetParent() types.Entity { return bsrGlobalGroupMappings.parent }

func (bsrGlobalGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetFilter() yfilter.YFilter { return globalRoutes.YFilter }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) SetFilter(yf yfilter.YFilter) { globalRoutes.YFilter = yf }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetSegmentPath() string {
    return "global-routes"
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = globalRoutes.MaximumRoutes
    leafs["warning-threshold"] = globalRoutes.WarningThreshold
    return leafs
}

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetBundleName() string { return "cisco_ios_xr" }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetYangName() string { return "global-routes" }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) SetParent(parent types.Entity) { globalRoutes.parent = parent }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetParent() types.Entity { return globalRoutes.parent }

func (globalRoutes *Pim_DefaultContext_Ipv4_Maximum_GlobalRoutes) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp
// Maximum for number of group mappings from
// autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGlobalGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGlobalGroupRangesAutoRp interface{}
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetFilter() yfilter.YFilter { return globalGroupMappingsAutoRp.YFilter }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { globalGroupMappingsAutoRp.YFilter = yf }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-global-group-ranges-auto-rp" { return "MaximumGlobalGroupRangesAutoRp" }
    if yname == "threshold-global-group-ranges-auto-rp" { return "ThresholdGlobalGroupRangesAutoRp" }
    return ""
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetSegmentPath() string {
    return "global-group-mappings-auto-rp"
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-global-group-ranges-auto-rp"] = globalGroupMappingsAutoRp.MaximumGlobalGroupRangesAutoRp
    leafs["threshold-global-group-ranges-auto-rp"] = globalGroupMappingsAutoRp.ThresholdGlobalGroupRangesAutoRp
    return leafs
}

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetYangName() string { return "global-group-mappings-auto-rp" }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) SetParent(parent types.Entity) { globalGroupMappingsAutoRp.parent = parent }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetParent() types.Entity { return globalGroupMappingsAutoRp.parent }

func (globalGroupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GlobalGroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache
// Override default global maximum and threshold
// for C-RP set in BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Global Maximum number of PIM C-RP Sets from BSR. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    BsrMaximumGlobalCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetFilter() yfilter.YFilter { return bsrGlobalCandidateRpCache.YFilter }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrGlobalCandidateRpCache.YFilter = yf }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-global-candidate-rp-cache" { return "BsrMaximumGlobalCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetSegmentPath() string {
    return "bsr-global-candidate-rp-cache"
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-global-candidate-rp-cache"] = bsrGlobalCandidateRpCache.BsrMaximumGlobalCandidateRpCache
    leafs["warning-threshold"] = bsrGlobalCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetYangName() string { return "bsr-global-candidate-rp-cache" }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) SetParent(parent types.Entity) { bsrGlobalCandidateRpCache.parent = parent }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetParent() types.Entity { return bsrGlobalCandidateRpCache.parent }

func (bsrGlobalCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrGlobalCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetFilter() yfilter.YFilter { return globalRegisterStates.YFilter }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) SetFilter(yf yfilter.YFilter) { globalRegisterStates.YFilter = yf }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetSegmentPath() string {
    return "global-register-states"
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = globalRegisterStates.MaximumRegisterStates
    leafs["warning-threshold"] = globalRegisterStates.WarningThreshold
    return leafs
}

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetYangName() string { return "global-register-states" }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) SetParent(parent types.Entity) { globalRegisterStates.parent = parent }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetParent() types.Entity { return globalRegisterStates.parent }

func (globalRegisterStates *Pim_DefaultContext_Ipv4_Maximum_GlobalRegisterStates) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetFilter() yfilter.YFilter { return globalRouteInterfaces.YFilter }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) SetFilter(yf yfilter.YFilter) { globalRouteInterfaces.YFilter = yf }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetSegmentPath() string {
    return "global-route-interfaces"
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = globalRouteInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = globalRouteInterfaces.WarningThreshold
    return leafs
}

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetYangName() string { return "global-route-interfaces" }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) SetParent(parent types.Entity) { globalRouteInterfaces.parent = parent }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetParent() types.Entity { return globalRouteInterfaces.parent }

func (globalRouteInterfaces *Pim_DefaultContext_Ipv4_Maximum_GlobalRouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp
// Override default maximum for number of group
// mappings from autorp mapping agent
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from autorp. The type is interface{}
    // with range: 1..10000. This attribute is mandatory.
    MaximumGroupRangesAutoRp interface{}

    // Warning threshold number of PIM group mappings from autorp. The type is
    // interface{} with range: 1..10000. The default value is 450.
    ThresholdGroupRangesAutoRp interface{}
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetFilter() yfilter.YFilter { return groupMappingsAutoRp.YFilter }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) SetFilter(yf yfilter.YFilter) { groupMappingsAutoRp.YFilter = yf }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetGoName(yname string) string {
    if yname == "maximum-group-ranges-auto-rp" { return "MaximumGroupRangesAutoRp" }
    if yname == "threshold-group-ranges-auto-rp" { return "ThresholdGroupRangesAutoRp" }
    return ""
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetSegmentPath() string {
    return "group-mappings-auto-rp"
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-group-ranges-auto-rp"] = groupMappingsAutoRp.MaximumGroupRangesAutoRp
    leafs["threshold-group-ranges-auto-rp"] = groupMappingsAutoRp.ThresholdGroupRangesAutoRp
    return leafs
}

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetYangName() string { return "group-mappings-auto-rp" }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) SetParent(parent types.Entity) { groupMappingsAutoRp.parent = parent }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetParent() types.Entity { return groupMappingsAutoRp.parent }

func (groupMappingsAutoRp *Pim_DefaultContext_Ipv4_Maximum_GroupMappingsAutoRp) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings
// Override default maximum and threshold for
// number of group mappings from BSR
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM group mappings from BSR. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumGroupRanges interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 500.
    WarningThreshold interface{}
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetFilter() yfilter.YFilter { return bsrGroupMappings.YFilter }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) SetFilter(yf yfilter.YFilter) { bsrGroupMappings.YFilter = yf }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetGoName(yname string) string {
    if yname == "bsr-maximum-group-ranges" { return "BsrMaximumGroupRanges" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetSegmentPath() string {
    return "bsr-group-mappings"
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-group-ranges"] = bsrGroupMappings.BsrMaximumGroupRanges
    leafs["warning-threshold"] = bsrGroupMappings.WarningThreshold
    return leafs
}

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetBundleName() string { return "cisco_ios_xr" }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetYangName() string { return "bsr-group-mappings" }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) SetParent(parent types.Entity) { bsrGroupMappings.parent = parent }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetParent() types.Entity { return bsrGroupMappings.parent }

func (bsrGroupMappings *Pim_DefaultContext_Ipv4_Maximum_BsrGroupMappings) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_RegisterStates
// Override default maximum for number of
// sparse-mode source registers
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_RegisterStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM Sparse-Mode register states. The type is interface{}
    // with range: 0..75000. This attribute is mandatory.
    MaximumRegisterStates interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 0..75000. The default value is 20000.
    WarningThreshold interface{}
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetFilter() yfilter.YFilter { return registerStates.YFilter }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) SetFilter(yf yfilter.YFilter) { registerStates.YFilter = yf }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetGoName(yname string) string {
    if yname == "maximum-register-states" { return "MaximumRegisterStates" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetSegmentPath() string {
    return "register-states"
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-register-states"] = registerStates.MaximumRegisterStates
    leafs["warning-threshold"] = registerStates.WarningThreshold
    return leafs
}

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetBundleName() string { return "cisco_ios_xr" }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetYangName() string { return "register-states" }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) SetParent(parent types.Entity) { registerStates.parent = parent }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetParent() types.Entity { return registerStates.parent }

func (registerStates *Pim_DefaultContext_Ipv4_Maximum_RegisterStates) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces
// Override default maximum for number of
// route-interfaces
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM route-interfaces. The type is interface{} with range:
    // 1..1100000. This attribute is mandatory.
    MaximumRouteInterfaces interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..1100000. The default value is 300000.
    WarningThreshold interface{}
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetFilter() yfilter.YFilter { return routeInterfaces.YFilter }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) SetFilter(yf yfilter.YFilter) { routeInterfaces.YFilter = yf }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetGoName(yname string) string {
    if yname == "maximum-route-interfaces" { return "MaximumRouteInterfaces" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetSegmentPath() string {
    return "route-interfaces"
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-route-interfaces"] = routeInterfaces.MaximumRouteInterfaces
    leafs["warning-threshold"] = routeInterfaces.WarningThreshold
    return leafs
}

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetYangName() string { return "route-interfaces" }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) SetParent(parent types.Entity) { routeInterfaces.parent = parent }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetParent() types.Entity { return routeInterfaces.parent }

func (routeInterfaces *Pim_DefaultContext_Ipv4_Maximum_RouteInterfaces) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache
// Override default maximum and threshold for BSR
// C-RP cache setting
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of BSR C-RP cache setting. The type is interface{} with
    // range: 1..10000. This attribute is mandatory.
    BsrMaximumCandidateRpCache interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..10000. The default value is 100.
    WarningThreshold interface{}
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetFilter() yfilter.YFilter { return bsrCandidateRpCache.YFilter }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) SetFilter(yf yfilter.YFilter) { bsrCandidateRpCache.YFilter = yf }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetGoName(yname string) string {
    if yname == "bsr-maximum-candidate-rp-cache" { return "BsrMaximumCandidateRpCache" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetSegmentPath() string {
    return "bsr-candidate-rp-cache"
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bsr-maximum-candidate-rp-cache"] = bsrCandidateRpCache.BsrMaximumCandidateRpCache
    leafs["warning-threshold"] = bsrCandidateRpCache.WarningThreshold
    return leafs
}

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetBundleName() string { return "cisco_ios_xr" }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetYangName() string { return "bsr-candidate-rp-cache" }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) SetParent(parent types.Entity) { bsrCandidateRpCache.parent = parent }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetParent() types.Entity { return bsrCandidateRpCache.parent }

func (bsrCandidateRpCache *Pim_DefaultContext_Ipv4_Maximum_BsrCandidateRpCache) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Maximum_Routes
// Override default maximum for number of routes
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Maximum_Routes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of PIM routes. The type is interface{} with range:
    // 1..200000. This attribute is mandatory.
    MaximumRoutes interface{}

    // Set threshold to print warning. The type is interface{} with range:
    // 1..200000. The default value is 100000.
    WarningThreshold interface{}
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetFilter() yfilter.YFilter { return routes.YFilter }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) SetFilter(yf yfilter.YFilter) { routes.YFilter = yf }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetGoName(yname string) string {
    if yname == "maximum-routes" { return "MaximumRoutes" }
    if yname == "warning-threshold" { return "WarningThreshold" }
    return ""
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetSegmentPath() string {
    return "routes"
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum-routes"] = routes.MaximumRoutes
    leafs["warning-threshold"] = routes.WarningThreshold
    return leafs
}

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetBundleName() string { return "cisco_ios_xr" }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetYangName() string { return "routes" }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) SetParent(parent types.Entity) { routes.parent = parent }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetParent() types.Entity { return routes.parent }

func (routes *Pim_DefaultContext_Ipv4_Maximum_Routes) GetParentYangName() string { return "maximum" }

// Pim_DefaultContext_Ipv4_Ssm
// Configure IP Multicast SSM
type Pim_DefaultContext_Ipv4_Ssm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if SSM is disabled on this router. The type is bool. The default value
    // is false.
    Disable interface{}

    // Access list of groups enabled with SSM. The type is string with length:
    // 1..64.
    Range interface{}
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetFilter() yfilter.YFilter { return ssm.YFilter }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) SetFilter(yf yfilter.YFilter) { ssm.YFilter = yf }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    if yname == "range" { return "Range" }
    return ""
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetSegmentPath() string {
    return "ssm"
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = ssm.Disable
    leafs["range"] = ssm.Range
    return leafs
}

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetBundleName() string { return "cisco_ios_xr" }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetYangName() string { return "ssm" }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) SetParent(parent types.Entity) { ssm.parent = parent }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetParent() types.Entity { return ssm.parent }

func (ssm *Pim_DefaultContext_Ipv4_Ssm) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Injects
// Inject Explicit PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Injects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inject Explicit PIM RPF Vector Proxy's. The type is slice of
    // Pim_DefaultContext_Ipv4_Injects_Inject.
    Inject []Pim_DefaultContext_Ipv4_Injects_Inject
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetFilter() yfilter.YFilter { return injects.YFilter }

func (injects *Pim_DefaultContext_Ipv4_Injects) SetFilter(yf yfilter.YFilter) { injects.YFilter = yf }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetGoName(yname string) string {
    if yname == "inject" { return "Inject" }
    return ""
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetSegmentPath() string {
    return "injects"
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inject" {
        for _, c := range injects.Inject {
            if injects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Injects_Inject{}
        injects.Inject = append(injects.Inject, child)
        return &injects.Inject[len(injects.Inject)-1]
    }
    return nil
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range injects.Inject {
        children[injects.Inject[i].GetSegmentPath()] = &injects.Inject[i]
    }
    return children
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (injects *Pim_DefaultContext_Ipv4_Injects) GetBundleName() string { return "cisco_ios_xr" }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetYangName() string { return "injects" }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (injects *Pim_DefaultContext_Ipv4_Injects) SetParent(parent types.Entity) { injects.parent = parent }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetParent() types.Entity { return injects.parent }

func (injects *Pim_DefaultContext_Ipv4_Injects) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Injects_Inject
// Inject Explicit PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Injects_Inject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetFilter() yfilter.YFilter { return inject.YFilter }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) SetFilter(yf yfilter.YFilter) { inject.YFilter = yf }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "rpf-proxy-address" { return "RpfProxyAddress" }
    return ""
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetSegmentPath() string {
    return "inject" + "[source-address='" + fmt.Sprintf("%v", inject.SourceAddress) + "']" + "[prefix-length='" + fmt.Sprintf("%v", inject.PrefixLength) + "']"
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = inject.SourceAddress
    leafs["prefix-length"] = inject.PrefixLength
    leafs["rpf-proxy-address"] = inject.RpfProxyAddress
    return leafs
}

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetBundleName() string { return "cisco_ios_xr" }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetYangName() string { return "inject" }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) SetParent(parent types.Entity) { inject.parent = parent }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetParent() types.Entity { return inject.parent }

func (inject *Pim_DefaultContext_Ipv4_Injects_Inject) GetParentYangName() string { return "injects" }

// Pim_DefaultContext_Ipv4_BidirRpAddresses
// Configure Bidirectional PIM Rendezvous Point
type Pim_DefaultContext_Ipv4_BidirRpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Rendezvous Point. The type is slice of
    // Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress.
    BidirRpAddress []Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetFilter() yfilter.YFilter { return bidirRpAddresses.YFilter }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) SetFilter(yf yfilter.YFilter) { bidirRpAddresses.YFilter = yf }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetGoName(yname string) string {
    if yname == "bidir-rp-address" { return "BidirRpAddress" }
    return ""
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetSegmentPath() string {
    return "bidir-rp-addresses"
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bidir-rp-address" {
        for _, c := range bidirRpAddresses.BidirRpAddress {
            if bidirRpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress{}
        bidirRpAddresses.BidirRpAddress = append(bidirRpAddresses.BidirRpAddress, child)
        return &bidirRpAddresses.BidirRpAddress[len(bidirRpAddresses.BidirRpAddress)-1]
    }
    return nil
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bidirRpAddresses.BidirRpAddress {
        children[bidirRpAddresses.BidirRpAddress[i].GetSegmentPath()] = &bidirRpAddresses.BidirRpAddress[i]
    }
    return children
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetYangName() string { return "bidir-rp-addresses" }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) SetParent(parent types.Entity) { bidirRpAddresses.parent = parent }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetParent() types.Entity { return bidirRpAddresses.parent }

func (bidirRpAddresses *Pim_DefaultContext_Ipv4_BidirRpAddresses) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress
// Address of the Rendezvous Point
type Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP address of Rendezvous Point. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Access list of groups that should map to a given RP. The type is string
    // with length: 1..64.
    AccessListName interface{}

    // TRUE Indicates if static RP config overrides AutoRP and BSR. The type is
    // bool.
    AutoRpOverride interface{}
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetFilter() yfilter.YFilter { return bidirRpAddress.YFilter }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) SetFilter(yf yfilter.YFilter) { bidirRpAddress.YFilter = yf }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "auto-rp-override" { return "AutoRpOverride" }
    return ""
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetSegmentPath() string {
    return "bidir-rp-address" + "[rp-address='" + fmt.Sprintf("%v", bidirRpAddress.RpAddress) + "']"
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = bidirRpAddress.RpAddress
    leafs["access-list-name"] = bidirRpAddress.AccessListName
    leafs["auto-rp-override"] = bidirRpAddress.AutoRpOverride
    return leafs
}

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetYangName() string { return "bidir-rp-address" }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) SetParent(parent types.Entity) { bidirRpAddress.parent = parent }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetParent() types.Entity { return bidirRpAddress.parent }

func (bidirRpAddress *Pim_DefaultContext_Ipv4_BidirRpAddresses_BidirRpAddress) GetParentYangName() string { return "bidir-rp-addresses" }

// Pim_DefaultContext_Ipv4_Bsr
// PIM BSR configuration
type Pim_DefaultContext_Ipv4_Bsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIM Candidate BSR configuration.
    CandidateBsr Pim_DefaultContext_Ipv4_Bsr_CandidateBsr

    // PIM RP configuration.
    CandidateRps Pim_DefaultContext_Ipv4_Bsr_CandidateRps
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetFilter() yfilter.YFilter { return bsr.YFilter }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) SetFilter(yf yfilter.YFilter) { bsr.YFilter = yf }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetGoName(yname string) string {
    if yname == "candidate-bsr" { return "CandidateBsr" }
    if yname == "candidate-rps" { return "CandidateRps" }
    return ""
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetSegmentPath() string {
    return "bsr"
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-bsr" {
        return &bsr.CandidateBsr
    }
    if childYangName == "candidate-rps" {
        return &bsr.CandidateRps
    }
    return nil
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-bsr"] = &bsr.CandidateBsr
    children["candidate-rps"] = &bsr.CandidateRps
    return children
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetBundleName() string { return "cisco_ios_xr" }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetYangName() string { return "bsr" }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) SetParent(parent types.Entity) { bsr.parent = parent }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetParent() types.Entity { return bsr.parent }

func (bsr *Pim_DefaultContext_Ipv4_Bsr) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Bsr_CandidateBsr
// PIM Candidate BSR configuration
// This type is a presence type.
type Pim_DefaultContext_Ipv4_Bsr_CandidateBsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BSR Address configured. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // Hash Mask Length for this candidate BSR. The type is interface{} with
    // range: 0..32. The default value is 30.
    PrefixLength interface{}

    // Priority of the Candidate BSR. The type is interface{} with range: 1..255.
    // The default value is 1.
    Priority interface{}
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetFilter() yfilter.YFilter { return candidateBsr.YFilter }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) SetFilter(yf yfilter.YFilter) { candidateBsr.YFilter = yf }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetSegmentPath() string {
    return "candidate-bsr"
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateBsr.Address
    leafs["prefix-length"] = candidateBsr.PrefixLength
    leafs["priority"] = candidateBsr.Priority
    return leafs
}

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetBundleName() string { return "cisco_ios_xr" }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetYangName() string { return "candidate-bsr" }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) SetParent(parent types.Entity) { candidateBsr.parent = parent }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetParent() types.Entity { return candidateBsr.parent }

func (candidateBsr *Pim_DefaultContext_Ipv4_Bsr_CandidateBsr) GetParentYangName() string { return "bsr" }

// Pim_DefaultContext_Ipv4_Bsr_CandidateRps
// PIM RP configuration
type Pim_DefaultContext_Ipv4_Bsr_CandidateRps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of PIM SM BSR Candidate-RP. The type is slice of
    // Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp.
    CandidateRp []Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetFilter() yfilter.YFilter { return candidateRps.YFilter }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) SetFilter(yf yfilter.YFilter) { candidateRps.YFilter = yf }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    return ""
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetSegmentPath() string {
    return "candidate-rps"
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        for _, c := range candidateRps.CandidateRp {
            if candidateRps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp{}
        candidateRps.CandidateRp = append(candidateRps.CandidateRp, child)
        return &candidateRps.CandidateRp[len(candidateRps.CandidateRp)-1]
    }
    return nil
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range candidateRps.CandidateRp {
        children[candidateRps.CandidateRp[i].GetSegmentPath()] = &candidateRps.CandidateRp[i]
    }
    return children
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetYangName() string { return "candidate-rps" }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) SetParent(parent types.Entity) { candidateRps.parent = parent }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetParent() types.Entity { return candidateRps.parent }

func (candidateRps *Pim_DefaultContext_Ipv4_Bsr_CandidateRps) GetParentYangName() string { return "bsr" }

// Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp
// Address of PIM SM BSR Candidate-RP
type Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address of Candidate-RP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. SM or Bidir. The type is PimProtocolMode.
    Mode interface{}

    // Access-list specifying the group range for the Candidate-RP. The type is
    // string with length: 1..64.
    GroupList interface{}

    // Priority of the CRP. The type is interface{} with range: 1..255. The
    // default value is 192.
    Priority interface{}

    // Advertisement interval. The type is interface{} with range: 30..600. The
    // default value is 60.
    Interval interface{}
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "mode" { return "Mode" }
    if yname == "group-list" { return "GroupList" }
    if yname == "priority" { return "Priority" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetSegmentPath() string {
    return "candidate-rp" + "[address='" + fmt.Sprintf("%v", candidateRp.Address) + "']" + "[mode='" + fmt.Sprintf("%v", candidateRp.Mode) + "']"
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = candidateRp.Address
    leafs["mode"] = candidateRp.Mode
    leafs["group-list"] = candidateRp.GroupList
    leafs["priority"] = candidateRp.Priority
    leafs["interval"] = candidateRp.Interval
    return leafs
}

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *Pim_DefaultContext_Ipv4_Bsr_CandidateRps_CandidateRp) GetParentYangName() string { return "candidate-rps" }

// Pim_DefaultContext_Ipv4_Mofrr
// Multicast Only Fast Re-Route
type Pim_DefaultContext_Ipv4_Mofrr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifying SG that should do RIB MOFRR. The type is string with
    // length: 1..64.
    Rib interface{}

    // Non-revertive Multicast Only Fast Re-Route. The type is interface{}.
    NonRevertive interface{}

    // Enable Multicast Only FRR. The type is interface{}.
    Enable interface{}

    // Access-list specifying SG that should do FLOW MOFRR. The type is string
    // with length: 1..64.
    Flow interface{}

    // Clone multicast joins.
    CloneJoins Pim_DefaultContext_Ipv4_Mofrr_CloneJoins

    // Clone multicast traffic.
    CloneSources Pim_DefaultContext_Ipv4_Mofrr_CloneSources
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetFilter() yfilter.YFilter { return mofrr.YFilter }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) SetFilter(yf yfilter.YFilter) { mofrr.YFilter = yf }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetGoName(yname string) string {
    if yname == "rib" { return "Rib" }
    if yname == "non-revertive" { return "NonRevertive" }
    if yname == "enable" { return "Enable" }
    if yname == "flow" { return "Flow" }
    if yname == "clone-joins" { return "CloneJoins" }
    if yname == "clone-sources" { return "CloneSources" }
    return ""
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetSegmentPath() string {
    return "mofrr"
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-joins" {
        return &mofrr.CloneJoins
    }
    if childYangName == "clone-sources" {
        return &mofrr.CloneSources
    }
    return nil
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clone-joins"] = &mofrr.CloneJoins
    children["clone-sources"] = &mofrr.CloneSources
    return children
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rib"] = mofrr.Rib
    leafs["non-revertive"] = mofrr.NonRevertive
    leafs["enable"] = mofrr.Enable
    leafs["flow"] = mofrr.Flow
    return leafs
}

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetBundleName() string { return "cisco_ios_xr" }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetYangName() string { return "mofrr" }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) SetParent(parent types.Entity) { mofrr.parent = parent }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetParent() types.Entity { return mofrr.parent }

func (mofrr *Pim_DefaultContext_Ipv4_Mofrr) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Mofrr_CloneJoins
// Clone multicast joins
type Pim_DefaultContext_Ipv4_Mofrr_CloneJoins struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clone S,G joins as S1,G joins and S2,G joins. The type is slice of
    // Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin.
    CloneJoin []Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetFilter() yfilter.YFilter { return cloneJoins.YFilter }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) SetFilter(yf yfilter.YFilter) { cloneJoins.YFilter = yf }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetGoName(yname string) string {
    if yname == "clone-join" { return "CloneJoin" }
    return ""
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetSegmentPath() string {
    return "clone-joins"
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-join" {
        for _, c := range cloneJoins.CloneJoin {
            if cloneJoins.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin{}
        cloneJoins.CloneJoin = append(cloneJoins.CloneJoin, child)
        return &cloneJoins.CloneJoin[len(cloneJoins.CloneJoin)-1]
    }
    return nil
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cloneJoins.CloneJoin {
        children[cloneJoins.CloneJoin[i].GetSegmentPath()] = &cloneJoins.CloneJoin[i]
    }
    return children
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetBundleName() string { return "cisco_ios_xr" }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetYangName() string { return "clone-joins" }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) SetParent(parent types.Entity) { cloneJoins.parent = parent }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetParent() types.Entity { return cloneJoins.parent }

func (cloneJoins *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins) GetParentYangName() string { return "mofrr" }

// Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin
// Clone S,G joins as S1,G joins and S2,G joins
type Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetFilter() yfilter.YFilter { return cloneJoin.YFilter }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) SetFilter(yf yfilter.YFilter) { cloneJoin.YFilter = yf }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    if yname == "primary" { return "Primary" }
    if yname == "backup" { return "Backup" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetSegmentPath() string {
    return "clone-join" + "[source='" + fmt.Sprintf("%v", cloneJoin.Source) + "']" + "[primary='" + fmt.Sprintf("%v", cloneJoin.Primary) + "']" + "[backup='" + fmt.Sprintf("%v", cloneJoin.Backup) + "']" + "[prefix-length='" + fmt.Sprintf("%v", cloneJoin.PrefixLength) + "']"
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = cloneJoin.Source
    leafs["primary"] = cloneJoin.Primary
    leafs["backup"] = cloneJoin.Backup
    leafs["prefix-length"] = cloneJoin.PrefixLength
    return leafs
}

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetBundleName() string { return "cisco_ios_xr" }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetYangName() string { return "clone-join" }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) SetParent(parent types.Entity) { cloneJoin.parent = parent }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetParent() types.Entity { return cloneJoin.parent }

func (cloneJoin *Pim_DefaultContext_Ipv4_Mofrr_CloneJoins_CloneJoin) GetParentYangName() string { return "clone-joins" }

// Pim_DefaultContext_Ipv4_Mofrr_CloneSources
// Clone multicast traffic
type Pim_DefaultContext_Ipv4_Mofrr_CloneSources struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clone S,G traffic as S1,G traffic and S2,G traffic. The type is slice of
    // Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource.
    CloneSource []Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetFilter() yfilter.YFilter { return cloneSources.YFilter }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) SetFilter(yf yfilter.YFilter) { cloneSources.YFilter = yf }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetGoName(yname string) string {
    if yname == "clone-source" { return "CloneSource" }
    return ""
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetSegmentPath() string {
    return "clone-sources"
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clone-source" {
        for _, c := range cloneSources.CloneSource {
            if cloneSources.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource{}
        cloneSources.CloneSource = append(cloneSources.CloneSource, child)
        return &cloneSources.CloneSource[len(cloneSources.CloneSource)-1]
    }
    return nil
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cloneSources.CloneSource {
        children[cloneSources.CloneSource[i].GetSegmentPath()] = &cloneSources.CloneSource[i]
    }
    return children
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetBundleName() string { return "cisco_ios_xr" }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetYangName() string { return "clone-sources" }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) SetParent(parent types.Entity) { cloneSources.parent = parent }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetParent() types.Entity { return cloneSources.parent }

func (cloneSources *Pim_DefaultContext_Ipv4_Mofrr_CloneSources) GetParentYangName() string { return "mofrr" }

// Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource
// Clone S,G traffic as S1,G traffic and S2,G
// traffic
type Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Original source address (S). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Primary cloned address (S1). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Primary interface{}

    // This attribute is a key. Backup cloned address (S2). The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Backup interface{}

    // This attribute is a key. Mask length. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetFilter() yfilter.YFilter { return cloneSource.YFilter }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) SetFilter(yf yfilter.YFilter) { cloneSource.YFilter = yf }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    if yname == "primary" { return "Primary" }
    if yname == "backup" { return "Backup" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetSegmentPath() string {
    return "clone-source" + "[source='" + fmt.Sprintf("%v", cloneSource.Source) + "']" + "[primary='" + fmt.Sprintf("%v", cloneSource.Primary) + "']" + "[backup='" + fmt.Sprintf("%v", cloneSource.Backup) + "']" + "[prefix-length='" + fmt.Sprintf("%v", cloneSource.PrefixLength) + "']"
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = cloneSource.Source
    leafs["primary"] = cloneSource.Primary
    leafs["backup"] = cloneSource.Backup
    leafs["prefix-length"] = cloneSource.PrefixLength
    return leafs
}

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetBundleName() string { return "cisco_ios_xr" }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetYangName() string { return "clone-source" }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) SetParent(parent types.Entity) { cloneSource.parent = parent }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetParent() types.Entity { return cloneSource.parent }

func (cloneSource *Pim_DefaultContext_Ipv4_Mofrr_CloneSources_CloneSource) GetParentYangName() string { return "clone-sources" }

// Pim_DefaultContext_Ipv4_Paths
// Inject PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Inject PIM RPF Vector Proxy's. The type is slice of
    // Pim_DefaultContext_Ipv4_Paths_Path.
    Path []Pim_DefaultContext_Ipv4_Paths_Path
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *Pim_DefaultContext_Ipv4_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pim_DefaultContext_Ipv4_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *Pim_DefaultContext_Ipv4_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetYangName() string { return "paths" }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *Pim_DefaultContext_Ipv4_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetParent() types.Entity { return paths.parent }

func (paths *Pim_DefaultContext_Ipv4_Paths) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Paths_Path
// Inject PIM RPF Vector Proxy's
type Pim_DefaultContext_Ipv4_Paths_Path struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Masklen. The type is interface{} with range:
    // 0..32.
    PrefixLength interface{}

    // RPF Proxy Address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpfProxyAddress []interface{}
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "rpf-proxy-address" { return "RpfProxyAddress" }
    return ""
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetSegmentPath() string {
    return "path" + "[source-address='" + fmt.Sprintf("%v", path.SourceAddress) + "']" + "[prefix-length='" + fmt.Sprintf("%v", path.PrefixLength) + "']"
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = path.SourceAddress
    leafs["prefix-length"] = path.PrefixLength
    leafs["rpf-proxy-address"] = path.RpfProxyAddress
    return leafs
}

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetYangName() string { return "path" }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *Pim_DefaultContext_Ipv4_Paths_Path) GetParentYangName() string { return "paths" }

// Pim_DefaultContext_Ipv4_AllowRp
// Enable allow-rp filtering for SM joins
// This type is a presence type.
type Pim_DefaultContext_Ipv4_AllowRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access-list specifiying applicable RPs. The type is string with length:
    // 1..64.
    RpListName interface{}

    // Access-list specifiying applicable groups. The type is string with length:
    // 1..64.
    GroupListName interface{}
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetFilter() yfilter.YFilter { return allowRp.YFilter }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) SetFilter(yf yfilter.YFilter) { allowRp.YFilter = yf }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetGoName(yname string) string {
    if yname == "rp-list-name" { return "RpListName" }
    if yname == "group-list-name" { return "GroupListName" }
    return ""
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetSegmentPath() string {
    return "allow-rp"
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-list-name"] = allowRp.RpListName
    leafs["group-list-name"] = allowRp.GroupListName
    return leafs
}

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetBundleName() string { return "cisco_ios_xr" }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetYangName() string { return "allow-rp" }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) SetParent(parent types.Entity) { allowRp.parent = parent }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetParent() types.Entity { return allowRp.parent }

func (allowRp *Pim_DefaultContext_Ipv4_AllowRp) GetParentYangName() string { return "ipv4" }

// Pim_DefaultContext_Ipv4_Convergence
// Configure convergence parameters
type Pim_DefaultContext_Ipv4_Convergence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Dampen first join if RPF path is through one of the downstream neighbor.
    // The type is interface{} with range: 0..15. Units are second.
    RpfConflictJoinDelay interface{}

    // Delay prunes if route join state transitions to not-joined on link down.
    // The type is interface{} with range: 0..60. Units are second.
    LinkDownPruneDelay interface{}
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetFilter() yfilter.YFilter { return convergence.YFilter }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) SetFilter(yf yfilter.YFilter) { convergence.YFilter = yf }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetGoName(yname string) string {
    if yname == "rpf-conflict-join-delay" { return "RpfConflictJoinDelay" }
    if yname == "link-down-prune-delay" { return "LinkDownPruneDelay" }
    return ""
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetSegmentPath() string {
    return "convergence"
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rpf-conflict-join-delay"] = convergence.RpfConflictJoinDelay
    leafs["link-down-prune-delay"] = convergence.LinkDownPruneDelay
    return leafs
}

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetBundleName() string { return "cisco_ios_xr" }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetYangName() string { return "convergence" }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) SetParent(parent types.Entity) { convergence.parent = parent }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetParent() types.Entity { return convergence.parent }

func (convergence *Pim_DefaultContext_Ipv4_Convergence) GetParentYangName() string { return "ipv4" }

