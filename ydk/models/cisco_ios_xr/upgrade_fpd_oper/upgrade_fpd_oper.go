// This module contains a collection of YANG definitions
// for Cisco IOS-XR upgrade-fpd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fpd: Field programmable device (FPD) operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-upgrade-fpd-oper fpd}", reflect.TypeOf(Fpd_{}))
    ydk.RegisterEntity("Cisco-IOS-XR-upgrade-fpd-oper:fpd", reflect.TypeOf(Fpd_{}))
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

// Fpd_
// Field programmable device (FPD) operational data
type Fpd_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of FPD supported nodes.
    Nodes Fpd__Nodes

    // FPD packages information.
    Packages Fpd__Packages
}

func (fpd_ *Fpd_) GetEntityData() *types.CommonEntityData {
    fpd_.EntityData.YFilter = fpd_.YFilter
    fpd_.EntityData.YangName = "fpd"
    fpd_.EntityData.BundleName = "cisco_ios_xr"
    fpd_.EntityData.ParentYangName = "Cisco-IOS-XR-upgrade-fpd-oper"
    fpd_.EntityData.SegmentPath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd"
    fpd_.EntityData.AbsolutePath = fpd_.EntityData.SegmentPath
    fpd_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fpd_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fpd_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fpd_.EntityData.Children = types.NewOrderedMap()
    fpd_.EntityData.Children.Append("nodes", types.YChild{"Nodes", &fpd_.Nodes})
    fpd_.EntityData.Children.Append("packages", types.YChild{"Packages", &fpd_.Packages})
    fpd_.EntityData.Leafs = types.NewOrderedMap()

    fpd_.EntityData.YListKeys = []string {}

    return &(fpd_.EntityData)
}

// Fpd__Nodes
// List of FPD supported nodes
type Fpd__Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of Fpd__Nodes_Node.
    Node []*Fpd__Nodes_Node
}

func (nodes *Fpd__Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "fpd"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Fpd__Nodes_Node
// Information about a particular node
type Fpd__Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // FPD information table.
    Devices Fpd__Nodes_Node_Devices
}

func (node *Fpd__Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("devices", types.YChild{"Devices", &node.Devices})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Fpd__Nodes_Node_Devices
// FPD information table
type Fpd__Nodes_Node_Devices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FPD information for a particular fpd type. The type is slice of
    // Fpd__Nodes_Node_Devices_Device.
    Device []*Fpd__Nodes_Node_Devices_Device
}

func (devices *Fpd__Nodes_Node_Devices) GetEntityData() *types.CommonEntityData {
    devices.EntityData.YFilter = devices.YFilter
    devices.EntityData.YangName = "devices"
    devices.EntityData.BundleName = "cisco_ios_xr"
    devices.EntityData.ParentYangName = "node"
    devices.EntityData.SegmentPath = "devices"
    devices.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/nodes/node/" + devices.EntityData.SegmentPath
    devices.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    devices.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    devices.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    devices.EntityData.Children = types.NewOrderedMap()
    devices.EntityData.Children.Append("device", types.YChild{"Device", nil})
    for i := range devices.Device {
        types.SetYListKey(devices.Device[i], i)
        devices.EntityData.Children.Append(types.GetSegmentPath(devices.Device[i]), types.YChild{"Device", devices.Device[i]})
    }
    devices.EntityData.Leafs = types.NewOrderedMap()

    devices.EntityData.YListKeys = []string {}

    return &(devices.EntityData)
}

// Fpd__Nodes_Node_Devices_Device
// FPD information for a particular fpd type
type Fpd__Nodes_Node_Devices_Device struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // FPD type. The type is Fpd.
    FpdType interface{}

    // Instance. The type is interface{} with range: 0..4294967295.
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

func (device *Fpd__Nodes_Node_Devices_Device) GetEntityData() *types.CommonEntityData {
    device.EntityData.YFilter = device.YFilter
    device.EntityData.YangName = "device"
    device.EntityData.BundleName = "cisco_ios_xr"
    device.EntityData.ParentYangName = "devices"
    device.EntityData.SegmentPath = "device" + types.AddNoKeyToken(device)
    device.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/nodes/node/devices/" + device.EntityData.SegmentPath
    device.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    device.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    device.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    device.EntityData.Children = types.NewOrderedMap()
    device.EntityData.Leafs = types.NewOrderedMap()
    device.EntityData.Leafs.Append("fpd-type", types.YLeaf{"FpdType", device.FpdType})
    device.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", device.Instance})
    device.EntityData.Leafs.Append("sub-type", types.YLeaf{"SubType", device.SubType})
    device.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", device.CardType})
    device.EntityData.Leafs.Append("hardware-version", types.YLeaf{"HardwareVersion", device.HardwareVersion})
    device.EntityData.Leafs.Append("software-version", types.YLeaf{"SoftwareVersion", device.SoftwareVersion})
    device.EntityData.Leafs.Append("is-upgrade-downgrade", types.YLeaf{"IsUpgradeDowngrade", device.IsUpgradeDowngrade})

    device.EntityData.YListKeys = []string {}

    return &(device.EntityData)
}

// Fpd__Packages
// FPD packages information
type Fpd__Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of packages. The type is slice of Fpd__Packages_AllPackage.
    AllPackage []*Fpd__Packages_AllPackage
}

func (packages *Fpd__Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "fpd"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/" + packages.EntityData.SegmentPath
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Children.Append("all-package", types.YChild{"AllPackage", nil})
    for i := range packages.AllPackage {
        types.SetYListKey(packages.AllPackage[i], i)
        packages.EntityData.Children.Append(types.GetSegmentPath(packages.AllPackage[i]), types.YChild{"AllPackage", packages.AllPackage[i]})
    }
    packages.EntityData.Leafs = types.NewOrderedMap()

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// Fpd__Packages_AllPackage
// List of packages
type Fpd__Packages_AllPackage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (allPackage *Fpd__Packages_AllPackage) GetEntityData() *types.CommonEntityData {
    allPackage.EntityData.YFilter = allPackage.YFilter
    allPackage.EntityData.YangName = "all-package"
    allPackage.EntityData.BundleName = "cisco_ios_xr"
    allPackage.EntityData.ParentYangName = "packages"
    allPackage.EntityData.SegmentPath = "all-package" + types.AddNoKeyToken(allPackage)
    allPackage.EntityData.AbsolutePath = "Cisco-IOS-XR-upgrade-fpd-oper:fpd/packages/" + allPackage.EntityData.SegmentPath
    allPackage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allPackage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allPackage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allPackage.EntityData.Children = types.NewOrderedMap()
    allPackage.EntityData.Leafs = types.NewOrderedMap()
    allPackage.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", allPackage.CardType})
    allPackage.EntityData.Leafs.Append("card-description", types.YLeaf{"CardDescription", allPackage.CardDescription})
    allPackage.EntityData.Leafs.Append("fpd-type", types.YLeaf{"FpdType", allPackage.FpdType})
    allPackage.EntityData.Leafs.Append("fpd-sub-type", types.YLeaf{"FpdSubType", allPackage.FpdSubType})
    allPackage.EntityData.Leafs.Append("software-version", types.YLeaf{"SoftwareVersion", allPackage.SoftwareVersion})
    allPackage.EntityData.Leafs.Append("minimum-required-software-version", types.YLeaf{"MinimumRequiredSoftwareVersion", allPackage.MinimumRequiredSoftwareVersion})
    allPackage.EntityData.Leafs.Append("minimum-required-hardware-version", types.YLeaf{"MinimumRequiredHardwareVersion", allPackage.MinimumRequiredHardwareVersion})

    allPackage.EntityData.YListKeys = []string {}

    return &(allPackage.EntityData)
}

