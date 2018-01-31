// This module contains a collection of YANG definitions
// for Cisco IOS-XR upgrade-fpd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fpd: Field programmable device (FPD) operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package upgrade_fpd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package upgrade_fpd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-upgrade-fpd-oper fpd}", reflect.TypeOf(Fpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-upgrade-fpd-oper:fpd", reflect.TypeOf(Fpd{}))
}

// FpdSub1 represents FPD sub types
type FpdSub1 string

const (
    // FPGA device
    FpdSub1_fpga1 FpdSub1 = "fpga1"

    // ROMMON device
    FpdSub1_rommon FpdSub1 = "rommon"

    // ROMMONA device
    FpdSub1_rommona FpdSub1 = "rommona"

    // Fabric loader
    FpdSub1_fabric_loader FpdSub1 = "fabric-loader"

    // FPGA device
    FpdSub1_fpga2 FpdSub1 = "fpga2"

    // FPGA device
    FpdSub1_fpga3 FpdSub1 = "fpga3"

    // FPGA device
    FpdSub1_fpga4 FpdSub1 = "fpga4"

    // FPGA device
    FpdSub1_fpga5 FpdSub1 = "fpga5"

    // FPGA device
    FpdSub1_fpga6 FpdSub1 = "fpga6"

    // FPGA device
    FpdSub1_fpga7 FpdSub1 = "fpga7"

    // FPGA device
    FpdSub1_fpga8 FpdSub1 = "fpga8"

    // FPGA device
    FpdSub1_fpga9 FpdSub1 = "fpga9"

    // FPGA device
    FpdSub1_fpga10 FpdSub1 = "fpga10"

    // FPGA device
    FpdSub1_fpga11 FpdSub1 = "fpga11"

    // FPGA device
    FpdSub1_fpga12 FpdSub1 = "fpga12"

    // FPGA device
    FpdSub1_fpga13 FpdSub1 = "fpga13"

    // FPGA device
    FpdSub1_fpga14 FpdSub1 = "fpga14"

    // CPLD device
    FpdSub1_cpld1 FpdSub1 = "cpld1"

    // CPLD device
    FpdSub1_cpld2 FpdSub1 = "cpld2"

    // CPLD device
    FpdSub1_cpld3 FpdSub1 = "cpld3"

    // CPLD device
    FpdSub1_cpld4 FpdSub1 = "cpld4"

    // CPLD device
    FpdSub1_cpld5 FpdSub1 = "cpld5"

    // CPLD device
    FpdSub1_cpld6 FpdSub1 = "cpld6"

    // CAN bus controller
    FpdSub1_cbc FpdSub1 = "cbc"

    // HSBI image
    FpdSub1_hsbi FpdSub1 = "hsbi"

    // Fabric Tx POD
    FpdSub1_txpod FpdSub1 = "txpod"

    // Fabric Rx POD
    FpdSub1_rxpod FpdSub1 = "rxpod"

    // IBMC
    FpdSub1_ibmc FpdSub1 = "ibmc"

    // FSBL
    FpdSub1_fsbl FpdSub1 = "fsbl"

    // Linux firmware
    FpdSub1_lnx FpdSub1 = "lnx"

    // FPGA device
    FpdSub1_fpga15 FpdSub1 = "fpga15"

    // FPGA device
    FpdSub1_fpga16 FpdSub1 = "fpga16"

    // FC FSBL
    FpdSub1_fc_fsbl FpdSub1 = "fc-fsbl"

    // FC linux firmware
    FpdSub1_fc_lnx FpdSub1 = "fc-lnx"
)

// Fpd represents Fpd
type Fpd string

const (
    // SPA class of fpd
    Fpd_spa Fpd = "spa"

    // Linecard class of fpd
    Fpd_lc Fpd = "lc"

    // SAM class of fpd
    Fpd_sam Fpd = "sam"
)

// Fpd1 represents FPD types
type Fpd1 string

const (
    // Shared port adapter
    Fpd1_spa Fpd1 = "spa"

    // Line card
    Fpd1_lc Fpd1 = "lc"

    // Service acceleration module
    Fpd1_sam Fpd1 = "sam"
)

// FpdSub represents Fpd sub
type FpdSub string

const (
    // FPGA device
    FpdSub_fpga1 FpdSub = "fpga1"

    // ROMMON device
    FpdSub_rommon FpdSub = "rommon"

    // ROMMON device #A
    FpdSub_rommona FpdSub = "rommona"

    // Fabric loader
    FpdSub_fabldr FpdSub = "fabldr"

    // FPGA device #2
    FpdSub_fpga2 FpdSub = "fpga2"

    // FPGA device #3
    FpdSub_fpga3 FpdSub = "fpga3"

    // FPGA device #4
    FpdSub_fpga4 FpdSub = "fpga4"

    // FPGA device #5
    FpdSub_fpga5 FpdSub = "fpga5"

    // FPGA device #6
    FpdSub_fpga6 FpdSub = "fpga6"

    // FPGA device #7
    FpdSub_fpga7 FpdSub = "fpga7"

    // FPGA device #8
    FpdSub_fpga8 FpdSub = "fpga8"

    // FPGA device #9
    FpdSub_fpga9 FpdSub = "fpga9"

    // FPGA device #10
    FpdSub_fpga10 FpdSub = "fpga10"

    // FPGA device #11
    FpdSub_fpga11 FpdSub = "fpga11"

    // FPGA device #12
    FpdSub_fpga12 FpdSub = "fpga12"

    // FPGA device #13
    FpdSub_fpga13 FpdSub = "fpga13"

    // FPGA device #14
    FpdSub_fpga14 FpdSub = "fpga14"

    // CPLD device #1
    FpdSub_cpld1 FpdSub = "cpld1"

    // CPLD device #2
    FpdSub_cpld2 FpdSub = "cpld2"

    // CPLD device #3
    FpdSub_cpld3 FpdSub = "cpld3"

    // CPLD device #4
    FpdSub_cpld4 FpdSub = "cpld4"

    // CPLD device #5
    FpdSub_cpld5 FpdSub = "cpld5"

    // CPLD device #6
    FpdSub_cpld6 FpdSub = "cpld6"

    // Can bus controller
    FpdSub_cbc FpdSub = "cbc"

    // HSBI image
    FpdSub_hsbi FpdSub = "hsbi"

    // Fabric Tx POD
    FpdSub_txpod FpdSub = "txpod"

    // Fabric Rx POD
    FpdSub_rxpod FpdSub = "rxpod"

    // IBMC
    FpdSub_ibmc FpdSub = "ibmc"

    // FSBL
    FpdSub_fsbl FpdSub = "fsbl"

    // Linux firmware
    FpdSub_lnx FpdSub = "lnx"

    // FPGA device #15
    FpdSub_fpga15 FpdSub = "fpga15"

    // FPGA device #16
    FpdSub_fpga16 FpdSub = "fpga16"

    // FC FSBL
    FpdSub_fc_fsbl FpdSub = "fc-fsbl"

    // FC linux firmware
    FpdSub_fc_lnx FpdSub = "fc-lnx"
)

// Fpd
// Field programmable device (FPD) operational data
type Fpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of FPD supported nodes.
    Nodes Fpd_Nodes

    // FPD packages information.
    Packages Fpd_Packages
}

func (fpd *Fpd) GetFilter() yfilter.YFilter { return fpd.YFilter }

func (fpd *Fpd) SetFilter(yf yfilter.YFilter) { fpd.YFilter = yf }

func (fpd *Fpd) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "packages" { return "Packages" }
    return ""
}

func (fpd *Fpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-upgrade-fpd-oper:fpd"
}

func (fpd *Fpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &fpd.Nodes
    }
    if childYangName == "packages" {
        return &fpd.Packages
    }
    return nil
}

func (fpd *Fpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &fpd.Nodes
    children["packages"] = &fpd.Packages
    return children
}

func (fpd *Fpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fpd *Fpd) GetBundleName() string { return "cisco_ios_xr" }

func (fpd *Fpd) GetYangName() string { return "fpd" }

func (fpd *Fpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fpd *Fpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fpd *Fpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fpd *Fpd) SetParent(parent types.Entity) { fpd.parent = parent }

func (fpd *Fpd) GetParent() types.Entity { return fpd.parent }

func (fpd *Fpd) GetParentYangName() string { return "Cisco-IOS-XR-upgrade-fpd-oper" }

// Fpd_Nodes
// List of FPD supported nodes
type Fpd_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of Fpd_Nodes_Node.
    Node []Fpd_Nodes_Node
}

func (nodes *Fpd_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Fpd_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Fpd_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Fpd_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Fpd_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fpd_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Fpd_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Fpd_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Fpd_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Fpd_Nodes) GetYangName() string { return "nodes" }

func (nodes *Fpd_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Fpd_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Fpd_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Fpd_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Fpd_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Fpd_Nodes) GetParentYangName() string { return "fpd" }

// Fpd_Nodes_Node
// Information about a particular node
type Fpd_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // FPD information table.
    Devices Fpd_Nodes_Node_Devices
}

func (node *Fpd_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Fpd_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Fpd_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "devices" { return "Devices" }
    return ""
}

func (node *Fpd_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Fpd_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "devices" {
        return &node.Devices
    }
    return nil
}

func (node *Fpd_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["devices"] = &node.Devices
    return children
}

func (node *Fpd_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Fpd_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Fpd_Nodes_Node) GetYangName() string { return "node" }

func (node *Fpd_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Fpd_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Fpd_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Fpd_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Fpd_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Fpd_Nodes_Node) GetParentYangName() string { return "nodes" }

// Fpd_Nodes_Node_Devices
// FPD information table
type Fpd_Nodes_Node_Devices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FPD information for a particular fpd type. The type is slice of
    // Fpd_Nodes_Node_Devices_Device.
    Device []Fpd_Nodes_Node_Devices_Device
}

func (devices *Fpd_Nodes_Node_Devices) GetFilter() yfilter.YFilter { return devices.YFilter }

func (devices *Fpd_Nodes_Node_Devices) SetFilter(yf yfilter.YFilter) { devices.YFilter = yf }

func (devices *Fpd_Nodes_Node_Devices) GetGoName(yname string) string {
    if yname == "device" { return "Device" }
    return ""
}

func (devices *Fpd_Nodes_Node_Devices) GetSegmentPath() string {
    return "devices"
}

func (devices *Fpd_Nodes_Node_Devices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device" {
        for _, c := range devices.Device {
            if devices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fpd_Nodes_Node_Devices_Device{}
        devices.Device = append(devices.Device, child)
        return &devices.Device[len(devices.Device)-1]
    }
    return nil
}

func (devices *Fpd_Nodes_Node_Devices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range devices.Device {
        children[devices.Device[i].GetSegmentPath()] = &devices.Device[i]
    }
    return children
}

func (devices *Fpd_Nodes_Node_Devices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (devices *Fpd_Nodes_Node_Devices) GetBundleName() string { return "cisco_ios_xr" }

func (devices *Fpd_Nodes_Node_Devices) GetYangName() string { return "devices" }

func (devices *Fpd_Nodes_Node_Devices) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (devices *Fpd_Nodes_Node_Devices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (devices *Fpd_Nodes_Node_Devices) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (devices *Fpd_Nodes_Node_Devices) SetParent(parent types.Entity) { devices.parent = parent }

func (devices *Fpd_Nodes_Node_Devices) GetParent() types.Entity { return devices.parent }

func (devices *Fpd_Nodes_Node_Devices) GetParentYangName() string { return "node" }

// Fpd_Nodes_Node_Devices_Device
// FPD information for a particular fpd type
type Fpd_Nodes_Node_Devices_Device struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FPD type. The type is Fpd.
    FpdType interface{}

    // Instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // FPD sub type. The type is FpdSub.
    SubType interface{}

    // Card type containing FPD. The type is string.
    CardType interface{}

    // FPD hardware version inX.Y format. X-Major version, Y-Minor version. The
    // type is string.
    HardwareVersion interface{}

    // FPD software version in X.Y format X-Major version, Y-Minor version Note:
    // 'Unknown' is returned in case the software version of the FPD cannot be
    // determined. The type is string.
    SoftwareVersion interface{}

    // If true, upgrade or downgrade. The type is bool.
    IsUpgradeDowngrade interface{}
}

func (device *Fpd_Nodes_Node_Devices_Device) GetFilter() yfilter.YFilter { return device.YFilter }

func (device *Fpd_Nodes_Node_Devices_Device) SetFilter(yf yfilter.YFilter) { device.YFilter = yf }

func (device *Fpd_Nodes_Node_Devices_Device) GetGoName(yname string) string {
    if yname == "fpd-type" { return "FpdType" }
    if yname == "instance" { return "Instance" }
    if yname == "sub-type" { return "SubType" }
    if yname == "card-type" { return "CardType" }
    if yname == "hardware-version" { return "HardwareVersion" }
    if yname == "software-version" { return "SoftwareVersion" }
    if yname == "is-upgrade-downgrade" { return "IsUpgradeDowngrade" }
    return ""
}

func (device *Fpd_Nodes_Node_Devices_Device) GetSegmentPath() string {
    return "device"
}

func (device *Fpd_Nodes_Node_Devices_Device) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (device *Fpd_Nodes_Node_Devices_Device) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (device *Fpd_Nodes_Node_Devices_Device) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["fpd-type"] = device.FpdType
    leafs["instance"] = device.Instance
    leafs["sub-type"] = device.SubType
    leafs["card-type"] = device.CardType
    leafs["hardware-version"] = device.HardwareVersion
    leafs["software-version"] = device.SoftwareVersion
    leafs["is-upgrade-downgrade"] = device.IsUpgradeDowngrade
    return leafs
}

func (device *Fpd_Nodes_Node_Devices_Device) GetBundleName() string { return "cisco_ios_xr" }

func (device *Fpd_Nodes_Node_Devices_Device) GetYangName() string { return "device" }

func (device *Fpd_Nodes_Node_Devices_Device) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (device *Fpd_Nodes_Node_Devices_Device) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (device *Fpd_Nodes_Node_Devices_Device) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (device *Fpd_Nodes_Node_Devices_Device) SetParent(parent types.Entity) { device.parent = parent }

func (device *Fpd_Nodes_Node_Devices_Device) GetParent() types.Entity { return device.parent }

func (device *Fpd_Nodes_Node_Devices_Device) GetParentYangName() string { return "devices" }

// Fpd_Packages
// FPD packages information
type Fpd_Packages struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of packages. The type is slice of Fpd_Packages_AllPackage.
    AllPackage []Fpd_Packages_AllPackage
}

func (packages *Fpd_Packages) GetFilter() yfilter.YFilter { return packages.YFilter }

func (packages *Fpd_Packages) SetFilter(yf yfilter.YFilter) { packages.YFilter = yf }

func (packages *Fpd_Packages) GetGoName(yname string) string {
    if yname == "all-package" { return "AllPackage" }
    return ""
}

func (packages *Fpd_Packages) GetSegmentPath() string {
    return "packages"
}

func (packages *Fpd_Packages) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-package" {
        for _, c := range packages.AllPackage {
            if packages.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fpd_Packages_AllPackage{}
        packages.AllPackage = append(packages.AllPackage, child)
        return &packages.AllPackage[len(packages.AllPackage)-1]
    }
    return nil
}

func (packages *Fpd_Packages) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range packages.AllPackage {
        children[packages.AllPackage[i].GetSegmentPath()] = &packages.AllPackage[i]
    }
    return children
}

func (packages *Fpd_Packages) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (packages *Fpd_Packages) GetBundleName() string { return "cisco_ios_xr" }

func (packages *Fpd_Packages) GetYangName() string { return "packages" }

func (packages *Fpd_Packages) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packages *Fpd_Packages) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packages *Fpd_Packages) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packages *Fpd_Packages) SetParent(parent types.Entity) { packages.parent = parent }

func (packages *Fpd_Packages) GetParent() types.Entity { return packages.parent }

func (packages *Fpd_Packages) GetParentYangName() string { return "fpd" }

// Fpd_Packages_AllPackage
// List of packages
type Fpd_Packages_AllPackage struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Card type containing FPD. The type is string.
    CardType interface{}

    // Card description. The type is string.
    CardDescription interface{}

    // FPD type. The type is Fpd1.
    FpdType interface{}

    // FPD sub type. The type is FpdSub1.
    FpdSubType interface{}

    // FPD software version in X.Y format X-Major version, Y-Minor version Note:
    // 'Unknown' is returned in case the software version of the FPD cannot be
    // determined. The type is string.
    SoftwareVersion interface{}

    // Minimum required FPD software version in X.Y format X-Major version,
    // Y-Minor version Note: 'Unknown' is returned in case the software version of
    // the FPD cannot be determined. The type is string.
    MinimumRequiredSoftwareVersion interface{}

    // Minimum required FPD hardware version in X.Y format X-Major version,
    // Y-Minor version . The type is string.
    MinimumRequiredHardwareVersion interface{}
}

func (allPackage *Fpd_Packages_AllPackage) GetFilter() yfilter.YFilter { return allPackage.YFilter }

func (allPackage *Fpd_Packages_AllPackage) SetFilter(yf yfilter.YFilter) { allPackage.YFilter = yf }

func (allPackage *Fpd_Packages_AllPackage) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "card-description" { return "CardDescription" }
    if yname == "fpd-type" { return "FpdType" }
    if yname == "fpd-sub-type" { return "FpdSubType" }
    if yname == "software-version" { return "SoftwareVersion" }
    if yname == "minimum-required-software-version" { return "MinimumRequiredSoftwareVersion" }
    if yname == "minimum-required-hardware-version" { return "MinimumRequiredHardwareVersion" }
    return ""
}

func (allPackage *Fpd_Packages_AllPackage) GetSegmentPath() string {
    return "all-package"
}

func (allPackage *Fpd_Packages_AllPackage) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allPackage *Fpd_Packages_AllPackage) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allPackage *Fpd_Packages_AllPackage) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = allPackage.CardType
    leafs["card-description"] = allPackage.CardDescription
    leafs["fpd-type"] = allPackage.FpdType
    leafs["fpd-sub-type"] = allPackage.FpdSubType
    leafs["software-version"] = allPackage.SoftwareVersion
    leafs["minimum-required-software-version"] = allPackage.MinimumRequiredSoftwareVersion
    leafs["minimum-required-hardware-version"] = allPackage.MinimumRequiredHardwareVersion
    return leafs
}

func (allPackage *Fpd_Packages_AllPackage) GetBundleName() string { return "cisco_ios_xr" }

func (allPackage *Fpd_Packages_AllPackage) GetYangName() string { return "all-package" }

func (allPackage *Fpd_Packages_AllPackage) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allPackage *Fpd_Packages_AllPackage) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allPackage *Fpd_Packages_AllPackage) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allPackage *Fpd_Packages_AllPackage) SetParent(parent types.Entity) { allPackage.parent = parent }

func (allPackage *Fpd_Packages_AllPackage) GetParent() types.Entity { return allPackage.parent }

func (allPackage *Fpd_Packages_AllPackage) GetParentYangName() string { return "packages" }

