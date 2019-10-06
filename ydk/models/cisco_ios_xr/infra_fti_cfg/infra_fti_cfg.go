// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-fti package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dci-fabric-interconnect: Configure FTI
//     parameters/sub-parameters
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_fti_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_fti_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-fti-cfg dci-fabric-interconnect}", reflect.TypeOf(DciFabricInterconnect{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect", reflect.TypeOf(DciFabricInterconnect{}))
}

// DciFabricInterconnect
// Configure FTI parameters/sub-parameters
type DciFabricInterconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Identity (Loopback IP address)<x.x.x.x>. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Identity interface{}

    // Configure fabric parameters.
    Fabrics DciFabricInterconnect_Fabrics

    // Configure Auto Config Pool parameters.
    Acp DciFabricInterconnect_Acp
}

func (dciFabricInterconnect *DciFabricInterconnect) GetEntityData() *types.CommonEntityData {
    dciFabricInterconnect.EntityData.YFilter = dciFabricInterconnect.YFilter
    dciFabricInterconnect.EntityData.YangName = "dci-fabric-interconnect"
    dciFabricInterconnect.EntityData.BundleName = "cisco_ios_xr"
    dciFabricInterconnect.EntityData.ParentYangName = "Cisco-IOS-XR-infra-fti-cfg"
    dciFabricInterconnect.EntityData.SegmentPath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect"
    dciFabricInterconnect.EntityData.AbsolutePath = dciFabricInterconnect.EntityData.SegmentPath
    dciFabricInterconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dciFabricInterconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dciFabricInterconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dciFabricInterconnect.EntityData.Children = types.NewOrderedMap()
    dciFabricInterconnect.EntityData.Children.Append("fabrics", types.YChild{"Fabrics", &dciFabricInterconnect.Fabrics})
    dciFabricInterconnect.EntityData.Children.Append("acp", types.YChild{"Acp", &dciFabricInterconnect.Acp})
    dciFabricInterconnect.EntityData.Leafs = types.NewOrderedMap()
    dciFabricInterconnect.EntityData.Leafs.Append("identity", types.YLeaf{"Identity", dciFabricInterconnect.Identity})

    dciFabricInterconnect.EntityData.YListKeys = []string {}

    return &(dciFabricInterconnect.EntityData)
}

// DciFabricInterconnect_Fabrics
// Configure fabric parameters
type DciFabricInterconnect_Fabrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter fabric identifier. The type is slice of
    // DciFabricInterconnect_Fabrics_Fabric.
    Fabric []*DciFabricInterconnect_Fabrics_Fabric
}

func (fabrics *DciFabricInterconnect_Fabrics) GetEntityData() *types.CommonEntityData {
    fabrics.EntityData.YFilter = fabrics.YFilter
    fabrics.EntityData.YangName = "fabrics"
    fabrics.EntityData.BundleName = "cisco_ios_xr"
    fabrics.EntityData.ParentYangName = "dci-fabric-interconnect"
    fabrics.EntityData.SegmentPath = "fabrics"
    fabrics.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/" + fabrics.EntityData.SegmentPath
    fabrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabrics.EntityData.Children = types.NewOrderedMap()
    fabrics.EntityData.Children.Append("fabric", types.YChild{"Fabric", nil})
    for i := range fabrics.Fabric {
        fabrics.EntityData.Children.Append(types.GetSegmentPath(fabrics.Fabric[i]), types.YChild{"Fabric", fabrics.Fabric[i]})
    }
    fabrics.EntityData.Leafs = types.NewOrderedMap()

    fabrics.EntityData.YListKeys = []string {}

    return &(fabrics.EntityData)
}

// DciFabricInterconnect_Fabrics_Fabric
// Enter fabric identifier
type DciFabricInterconnect_Fabrics_Fabric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. fabric identifier. The type is interface{} with
    // range: 1000..9999.
    Id1 interface{}

    // Disabled or Path to certificate. The type is string.
    Ssl interface{}

    // Enter Opflex peer info.
    Controllers DciFabricInterconnect_Fabrics_Fabric_Controllers
}

func (fabric *DciFabricInterconnect_Fabrics_Fabric) GetEntityData() *types.CommonEntityData {
    fabric.EntityData.YFilter = fabric.YFilter
    fabric.EntityData.YangName = "fabric"
    fabric.EntityData.BundleName = "cisco_ios_xr"
    fabric.EntityData.ParentYangName = "fabrics"
    fabric.EntityData.SegmentPath = "fabric" + types.AddKeyToken(fabric.Id1, "id1")
    fabric.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/fabrics/" + fabric.EntityData.SegmentPath
    fabric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabric.EntityData.Children = types.NewOrderedMap()
    fabric.EntityData.Children.Append("controllers", types.YChild{"Controllers", &fabric.Controllers})
    fabric.EntityData.Leafs = types.NewOrderedMap()
    fabric.EntityData.Leafs.Append("id1", types.YLeaf{"Id1", fabric.Id1})
    fabric.EntityData.Leafs.Append("ssl", types.YLeaf{"Ssl", fabric.Ssl})

    fabric.EntityData.YListKeys = []string {"Id1"}

    return &(fabric.EntityData)
}

// DciFabricInterconnect_Fabrics_Fabric_Controllers
// Enter Opflex peer info
type DciFabricInterconnect_Fabrics_Fabric_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter Spine IP address. The type is slice of
    // DciFabricInterconnect_Fabrics_Fabric_Controllers_Controller.
    Controller []*DciFabricInterconnect_Fabrics_Fabric_Controllers_Controller
}

func (controllers *DciFabricInterconnect_Fabrics_Fabric_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "fabric"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/fabrics/fabric/" + controllers.EntityData.SegmentPath
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

// DciFabricInterconnect_Fabrics_Fabric_Controllers_Controller
// Enter Spine IP address
type DciFabricInterconnect_Fabrics_Fabric_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Enter Spine IP address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ip1 interface{}
}

func (controller *DciFabricInterconnect_Fabrics_Fabric_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.Ip1, "ip1")
    controller.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/fabrics/fabric/controllers/" + controller.EntityData.SegmentPath
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("ip1", types.YLeaf{"Ip1", controller.Ip1})

    controller.EntityData.YListKeys = []string {"Ip1"}

    return &(controller.EntityData)
}

// DciFabricInterconnect_Acp
// Configure Auto Config Pool parameters
type DciFabricInterconnect_Acp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify NVE interface id. The type is interface{} with range:
    // 0..4294967295.
    NveId interface{}

    // Specify BGP AS number. The type is interface{} with range: 0..4294967295.
    BgpAs interface{}

    // Specify Bridge-group name. The type is string.
    BgName interface{}

    // Specify BD pool range.
    BdRange DciFabricInterconnect_Acp_BdRange

    // Specify VNI pool range.
    VniRange DciFabricInterconnect_Acp_VniRange

    // Specify BVI pool range.
    BviRange DciFabricInterconnect_Acp_BviRange

    // Configure local VRF parameters.
    Vrfs DciFabricInterconnect_Acp_Vrfs
}

func (acp *DciFabricInterconnect_Acp) GetEntityData() *types.CommonEntityData {
    acp.EntityData.YFilter = acp.YFilter
    acp.EntityData.YangName = "acp"
    acp.EntityData.BundleName = "cisco_ios_xr"
    acp.EntityData.ParentYangName = "dci-fabric-interconnect"
    acp.EntityData.SegmentPath = "acp"
    acp.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/" + acp.EntityData.SegmentPath
    acp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acp.EntityData.Children = types.NewOrderedMap()
    acp.EntityData.Children.Append("bd-range", types.YChild{"BdRange", &acp.BdRange})
    acp.EntityData.Children.Append("vni-range", types.YChild{"VniRange", &acp.VniRange})
    acp.EntityData.Children.Append("bvi-range", types.YChild{"BviRange", &acp.BviRange})
    acp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &acp.Vrfs})
    acp.EntityData.Leafs = types.NewOrderedMap()
    acp.EntityData.Leafs.Append("nve-id", types.YLeaf{"NveId", acp.NveId})
    acp.EntityData.Leafs.Append("bgp-as", types.YLeaf{"BgpAs", acp.BgpAs})
    acp.EntityData.Leafs.Append("bg-name", types.YLeaf{"BgName", acp.BgName})

    acp.EntityData.YListKeys = []string {}

    return &(acp.EntityData)
}

// DciFabricInterconnect_Acp_BdRange
// Specify BD pool range
type DciFabricInterconnect_Acp_BdRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BD Range:min value. The type is interface{} with range: 1..4000.
    BdMin interface{}

    // BD Range:max value. The type is interface{} with range: 0..4294967295.
    BdMax interface{}
}

func (bdRange *DciFabricInterconnect_Acp_BdRange) GetEntityData() *types.CommonEntityData {
    bdRange.EntityData.YFilter = bdRange.YFilter
    bdRange.EntityData.YangName = "bd-range"
    bdRange.EntityData.BundleName = "cisco_ios_xr"
    bdRange.EntityData.ParentYangName = "acp"
    bdRange.EntityData.SegmentPath = "bd-range"
    bdRange.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/acp/" + bdRange.EntityData.SegmentPath
    bdRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bdRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bdRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bdRange.EntityData.Children = types.NewOrderedMap()
    bdRange.EntityData.Leafs = types.NewOrderedMap()
    bdRange.EntityData.Leafs.Append("bd-min", types.YLeaf{"BdMin", bdRange.BdMin})
    bdRange.EntityData.Leafs.Append("bd-max", types.YLeaf{"BdMax", bdRange.BdMax})

    bdRange.EntityData.YListKeys = []string {}

    return &(bdRange.EntityData)
}

// DciFabricInterconnect_Acp_VniRange
// Specify VNI pool range
type DciFabricInterconnect_Acp_VniRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VNI Range:min value. The type is interface{} with range: 1..4000.
    VniMin interface{}

    // VNI Range:max value. The type is interface{} with range: 0..4294967295.
    VniMax interface{}
}

func (vniRange *DciFabricInterconnect_Acp_VniRange) GetEntityData() *types.CommonEntityData {
    vniRange.EntityData.YFilter = vniRange.YFilter
    vniRange.EntityData.YangName = "vni-range"
    vniRange.EntityData.BundleName = "cisco_ios_xr"
    vniRange.EntityData.ParentYangName = "acp"
    vniRange.EntityData.SegmentPath = "vni-range"
    vniRange.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/acp/" + vniRange.EntityData.SegmentPath
    vniRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vniRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vniRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vniRange.EntityData.Children = types.NewOrderedMap()
    vniRange.EntityData.Leafs = types.NewOrderedMap()
    vniRange.EntityData.Leafs.Append("vni-min", types.YLeaf{"VniMin", vniRange.VniMin})
    vniRange.EntityData.Leafs.Append("vni-max", types.YLeaf{"VniMax", vniRange.VniMax})

    vniRange.EntityData.YListKeys = []string {}

    return &(vniRange.EntityData)
}

// DciFabricInterconnect_Acp_BviRange
// Specify BVI pool range
type DciFabricInterconnect_Acp_BviRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BVI Range:min value. The type is interface{} with range: 1..4000.
    BviMin interface{}

    // BVI Range:max value. The type is interface{} with range: 0..4294967295.
    BviMax interface{}
}

func (bviRange *DciFabricInterconnect_Acp_BviRange) GetEntityData() *types.CommonEntityData {
    bviRange.EntityData.YFilter = bviRange.YFilter
    bviRange.EntityData.YangName = "bvi-range"
    bviRange.EntityData.BundleName = "cisco_ios_xr"
    bviRange.EntityData.ParentYangName = "acp"
    bviRange.EntityData.SegmentPath = "bvi-range"
    bviRange.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/acp/" + bviRange.EntityData.SegmentPath
    bviRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bviRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bviRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bviRange.EntityData.Children = types.NewOrderedMap()
    bviRange.EntityData.Leafs = types.NewOrderedMap()
    bviRange.EntityData.Leafs.Append("bvi-min", types.YLeaf{"BviMin", bviRange.BviMin})
    bviRange.EntityData.Leafs.Append("bvi-max", types.YLeaf{"BviMax", bviRange.BviMax})

    bviRange.EntityData.YListKeys = []string {}

    return &(bviRange.EntityData)
}

// DciFabricInterconnect_Acp_Vrfs
// Configure local VRF parameters
type DciFabricInterconnect_Acp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // vrf name. The type is slice of DciFabricInterconnect_Acp_Vrfs_Vrf.
    Vrf []*DciFabricInterconnect_Acp_Vrfs_Vrf
}

func (vrfs *DciFabricInterconnect_Acp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "acp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/acp/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// DciFabricInterconnect_Acp_Vrfs_Vrf
// vrf name
type DciFabricInterconnect_Acp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. vrf name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // BVI override IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    BviVrfIp interface{}
}

func (vrf *DciFabricInterconnect_Acp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-fti-cfg:dci-fabric-interconnect/acp/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("bvi-vrf-ip", types.YLeaf{"BviVrfIp", vrf.BviVrfIp})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

