// This module contains a collection of YANG definitions
// for Cisco IOS-XR atm-vcm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   atm-vcm: ATM VCM operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package atm_vcm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package atm_vcm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-atm-vcm-oper atm-vcm}", reflect.TypeOf(AtmVcm{}))
    ydk.RegisterEntity("Cisco-IOS-XR-atm-vcm-oper:atm-vcm", reflect.TypeOf(AtmVcm{}))
}

// Vc represents  ATM VC type
type Vc string

const (
    //  ATM Layer 3 VC type
    Vc_layer3_vc Vc = "layer3-vc"

    //  ATM Layer 2 VC type
    Vc_layer2_vc Vc = "layer2-vc"

    //  ATM Layer 2 VP type
    Vc_layer2_vp Vc = "layer2-vp"

    //  ATM type unknown
    Vc_vc_type_unknown Vc = "vc-type-unknown"
)

// VcEncap represents VC Encapsulation Type
type VcEncap string

const (
    // ILMI Encapsulation
    VcEncap_ilmi VcEncap = "ilmi"

    // QSAAL Encapsulation
    VcEncap_qsaal VcEncap = "qsaal"

    // SNAP Encapsulation
    VcEncap_snap VcEncap = "snap"

    // MUX Encapsulation
    VcEncap_mux VcEncap = "mux"

    // NLPID Encapsulation
    VcEncap_nlpid VcEncap = "nlpid"

    // F4OAM Encapsulation
    VcEncap_f4oam VcEncap = "f4oam"

    // AAL0 Encapsulation
    VcEncap_aal0 VcEncap = "aal0"

    // AAL5 Encapsulation
    VcEncap_aal5 VcEncap = "aal5"

    // Uknown (invalid) Encapsulation
    VcEncap_encap_unknown VcEncap = "encap-unknown"
)

// VcManageLevel represents ATM Class link manage level
type VcManageLevel string

const (
    // Managed
    VcManageLevel_manage VcManageLevel = "manage"

    // Not managed
    VcManageLevel_not_managed VcManageLevel = "not-managed"
)

// VcTestMode represents VC Test Mode Type
type VcTestMode string

const (
    // VC not in test mode
    VcTestMode_test_mode_none VcTestMode = "test-mode-none"

    // VC in test mode Loop
    VcTestMode_loop VcTestMode = "loop"

    // VC in test mode Reserved
    VcTestMode_reserved VcTestMode = "reserved"
)

// VpTrafShaping represents VP-Tunnel traffic shaping type
type VpTrafShaping string

const (
    // VP-Tunnel traffic shaping type CBR
    VpTrafShaping_vp_cbr VpTrafShaping = "vp-cbr"

    // VP-Tunnel traffic shaping type VBR-NR
    VpTrafShaping_vp_vbr_nrt VpTrafShaping = "vp-vbr-nrt"

    // VP-Tunnel traffic shaping type VBR-RT
    VpTrafShaping_vp_vbr_rt VpTrafShaping = "vp-vbr-rt"

    // VP-Tunnel traffic shaping type ABR
    VpTrafShaping_vp_abr VpTrafShaping = "vp-abr"

    // VP-Tunnel traffic shaping type UBR+
    VpTrafShaping_vp_ubr_plus VpTrafShaping = "vp-ubr-plus"

    // VP-Tunnel traffic shaping type UBR
    VpTrafShaping_vp_ubr VpTrafShaping = "vp-ubr"

    // VP-Tunnel traffic shaping type Unknown
    // (invalid)
    VpTrafShaping_vp_traf_shaping_unknown VpTrafShaping = "vp-traf-shaping-unknown"
)

// VcState represents VC State
type VcState string

const (
    // ATM VC State: only VC data structure
    // initialized   
    VcState_initialized VcState = "initialized"

    // ATM VC State: configuration currently being
    // changed
    VcState_modifying VcState = "modifying"

    // ATM VC State: configuration changed            
    VcState_modified VcState = "modified"

    // ATM VC State: command sent to hardware to
    // activate 
    VcState_activating VcState = "activating"

    // ATM VC State: activated in h/w or protocol     
    VcState_activated VcState = "activated"

    // ATM VC State: OAM/ILMI - yet to verify         
    VcState_not_verified VcState = "not-verified"

    // ATM VC State: Ready state                      
    VcState_ready VcState = "ready"

    // ATM VC State: command sent to h/w to deactivate
    VcState_deactivating VcState = "deactivating"

    // ATM VC State: inactive/not present in hardware 
    VcState_inactive VcState = "inactive"

    // ATM VC State: VC is being deleted              
    VcState_deleting VcState = "deleting"

    // ATM VC State: VC is already delete in hardware 
    VcState_deleted VcState = "deleted"

    // ATM VC State: Unknown(invalid)
    VcState_state_unknown VcState = "state-unknown"
)

// ClassLinkOamInheritLevel represents ATM VC-class inheritence level for class-link
type ClassLinkOamInheritLevel string

const (
    // Configured on VC
    ClassLinkOamInheritLevel_vc_configured_onvc ClassLinkOamInheritLevel = "vc-configured-onvc"

    // Class on VC
    ClassLinkOamInheritLevel_vc_class_onvc ClassLinkOamInheritLevel = "vc-class-onvc"

    // Class on sub-if
    ClassLinkOamInheritLevel_vc_class_on_sub_interface ClassLinkOamInheritLevel = "vc-class-on-sub-interface"

    // Class on main-if
    ClassLinkOamInheritLevel_vc_class_on_main_interface ClassLinkOamInheritLevel = "vc-class-on-main-interface"

    // Global default values
    ClassLinkOamInheritLevel_vc_global_default ClassLinkOamInheritLevel = "vc-global-default"

    // Unknown (invalid)
    ClassLinkOamInheritLevel_vc_inherit_level_unknown ClassLinkOamInheritLevel = "vc-inherit-level-unknown"
)

// VpState represents VP-Tunnel State
type VpState string

const (
    // VP-Tunnel State: Initialized
    VpState_vp_initialized VpState = "vp-initialized"

    // VP-Tunnel State: Modifying
    VpState_vp_modifying VpState = "vp-modifying"

    // VP-Tunnel State: Ready
    VpState_vp_ready VpState = "vp-ready"

    // VP-Tunnel State: Managed Down
    VpState_vp_mgd_down VpState = "vp-mgd-down"

    // VP-Tunnel State: Deleting
    VpState_vp_deleting VpState = "vp-deleting"

    // VP-Tunnel State: Deleted
    VpState_vp_deleted VpState = "vp-deleted"

    // VP-Tunnel State: Unknown
    VpState_vp_state_unknown VpState = "vp-state-unknown"
)

// VcTrafShaping represents VC traffic shaping type
type VcTrafShaping string

const (
    // VC traffic shaping type CBR
    VcTrafShaping_cbr VcTrafShaping = "cbr"

    // VC traffic shaping type VBR-NR
    VcTrafShaping_vbr_nrt VcTrafShaping = "vbr-nrt"

    // VC traffic shaping type VBR-RT
    VcTrafShaping_vbr_rt VcTrafShaping = "vbr-rt"

    // VC traffic shaping type ABR
    VcTrafShaping_abr VcTrafShaping = "abr"

    // VC traffic shaping type UBR+
    VcTrafShaping_ubr_plus VcTrafShaping = "ubr-plus"

    // VC traffic shaping type UBR
    VcTrafShaping_ubr VcTrafShaping = "ubr"

    // VC traffic shaping type Unknown (invalid)
    VcTrafShaping_traf_shaping_unknown VcTrafShaping = "traf-shaping-unknown"
)

// VcCellPackingMode represents ATM VC cell packing mode
type VcCellPackingMode string

const (
    // VP mode
    VcCellPackingMode_vp VcCellPackingMode = "vp"

    // VC mode
    VcCellPackingMode_vc VcCellPackingMode = "vc"

    // Port mode
    VcCellPackingMode_port_mode VcCellPackingMode = "port-mode"
)

// VcmPort represents ATM port type
type VcmPort string

const (
    //  Layer 2 ATM port type 
    VcmPort_port_type_layer_2 VcmPort = "port-type-layer-2"

    //  Layer 3 ATM port type 
    VcmPort_port_type_layer_3 VcmPort = "port-type-layer-3"

    //  ATM port type unknown 
    VcmPort_port_type_unknown VcmPort = "port-type-unknown"
)

// VcInheritLevel represents ATM vc-class inheritence level
type VcInheritLevel string

const (
    // ATM vc-class inherit level: Config of VC
    VcInheritLevel_directly_configured_onvc VcInheritLevel = "directly-configured-onvc"

    // ATM vc-class inherit level: Class of VC
    VcInheritLevel_vc_class_configured_onvc VcInheritLevel = "vc-class-configured-onvc"

    // ATM vc-class inherit level: Class of Sub-if
    VcInheritLevel_vc_class_configured_on_sub_interface VcInheritLevel = "vc-class-configured-on-sub-interface"

    // ATM vc-class inherit level: Class of Main-if
    VcInheritLevel_vc_class_configured_on_main_interface VcInheritLevel = "vc-class-configured-on-main-interface"

    // ATM vc-class inherit level: Global Default
    VcInheritLevel_global_default VcInheritLevel = "global-default"

    // ATM vc-class inherit level: Unknown (invalid)
    VcInheritLevel_unknown VcInheritLevel = "unknown"

    // ATM vc-class inherit level: Not supported on
    // this VC class
    VcInheritLevel_not_supported VcInheritLevel = "not-supported"
)

// AtmVcm
// ATM VCM operational data
type AtmVcm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains all the nodes.
    Nodes AtmVcm_Nodes
}

func (atmVcm *AtmVcm) GetEntityData() *types.CommonEntityData {
    atmVcm.EntityData.YFilter = atmVcm.YFilter
    atmVcm.EntityData.YangName = "atm-vcm"
    atmVcm.EntityData.BundleName = "cisco_ios_xr"
    atmVcm.EntityData.ParentYangName = "Cisco-IOS-XR-atm-vcm-oper"
    atmVcm.EntityData.SegmentPath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm"
    atmVcm.EntityData.AbsolutePath = atmVcm.EntityData.SegmentPath
    atmVcm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    atmVcm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    atmVcm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    atmVcm.EntityData.Children = types.NewOrderedMap()
    atmVcm.EntityData.Children.Append("nodes", types.YChild{"Nodes", &atmVcm.Nodes})
    atmVcm.EntityData.Leafs = types.NewOrderedMap()

    atmVcm.EntityData.YListKeys = []string {}

    return &(atmVcm.EntityData)
}

// AtmVcm_Nodes
// Contains all the nodes
type AtmVcm_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The node on which ATM Interfaces/VCs/VPs are located. The type is slice of
    // AtmVcm_Nodes_Node.
    Node []*AtmVcm_Nodes_Node
}

func (nodes *AtmVcm_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "atm-vcm"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/" + nodes.EntityData.SegmentPath
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

// AtmVcm_Nodes_Node
// The node on which ATM Interfaces/VCs/VPs are
// located
type AtmVcm_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Contains all VC information for node.
    Vcs AtmVcm_Nodes_Node_Vcs

    // Contains all cell packing information for node.
    CellPacks AtmVcm_Nodes_Node_CellPacks

    // Contains all L2 PVP information for node.
    Pvps AtmVcm_Nodes_Node_Pvps

    // Contains all class link information for node.
    ClassLinks AtmVcm_Nodes_Node_ClassLinks

    // Contains all Interface information for node.
    Interfaces AtmVcm_Nodes_Node_Interfaces

    // Contains all VP-tunnel information for node.
    VpTunnels AtmVcm_Nodes_Node_VpTunnels
}

func (node *AtmVcm_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("vcs", types.YChild{"Vcs", &node.Vcs})
    node.EntityData.Children.Append("cell-packs", types.YChild{"CellPacks", &node.CellPacks})
    node.EntityData.Children.Append("pvps", types.YChild{"Pvps", &node.Pvps})
    node.EntityData.Children.Append("class-links", types.YChild{"ClassLinks", &node.ClassLinks})
    node.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &node.Interfaces})
    node.EntityData.Children.Append("vp-tunnels", types.YChild{"VpTunnels", &node.VpTunnels})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// AtmVcm_Nodes_Node_Vcs
// Contains all VC information for node
type AtmVcm_Nodes_Node_Vcs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All VC information on a node. The type is slice of
    // AtmVcm_Nodes_Node_Vcs_Vc.
    Vc []*AtmVcm_Nodes_Node_Vcs_Vc
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetEntityData() *types.CommonEntityData {
    vcs.EntityData.YFilter = vcs.YFilter
    vcs.EntityData.YangName = "vcs"
    vcs.EntityData.BundleName = "cisco_ios_xr"
    vcs.EntityData.ParentYangName = "node"
    vcs.EntityData.SegmentPath = "vcs"
    vcs.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + vcs.EntityData.SegmentPath
    vcs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vcs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vcs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vcs.EntityData.Children = types.NewOrderedMap()
    vcs.EntityData.Children.Append("vc", types.YChild{"Vc", nil})
    for i := range vcs.Vc {
        vcs.EntityData.Children.Append(types.GetSegmentPath(vcs.Vc[i]), types.YChild{"Vc", vcs.Vc[i]})
    }
    vcs.EntityData.Leafs = types.NewOrderedMap()

    vcs.EntityData.YListKeys = []string {}

    return &(vcs.EntityData)
}

// AtmVcm_Nodes_Node_Vcs_Vc
// All VC information on a node
type AtmVcm_Nodes_Node_Vcs_Vc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: 0..4095.
    Vpi interface{}

    // VCI. The type is interface{} with range: 1..65535.
    Vci interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    MainInterface interface{}

    // Subinterface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SubInterface interface{}

    // VC Interfcace handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    VcInterface interface{}

    // VC VPI value. The type is interface{} with range: 0..65535.
    VpiXr interface{}

    // VC VCI value. The type is interface{} with range: 0..65535.
    VciXr interface{}

    // VC Type. The type is Vc.
    Type interface{}

    // Encapsulation type. The type is VcEncap.
    Encapsulation interface{}

    // ATM VC traffic shaping type. The type is VcTrafShaping.
    Shape interface{}

    // Peak cell rate in Kbps. The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    PeakCellRate interface{}

    // Sustained cell rate in Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    SustainedCellRate interface{}

    // Burst size in cells. The type is interface{} with range: 0..4294967295.
    BurstRate interface{}

    // Encapsulation inherit level - identifies if encapsulation is set to
    // default, configured on the VC, or inherited from the vcclass. The type is
    // VcInheritLevel.
    EncapsInheritLevel interface{}

    // Quality of Service inherit level - identifies if QoS is set to default,
    // configured on the VC, or inherited from the vcclass. The type is
    // VcInheritLevel.
    QosInheritLevel interface{}

    // Transmit MTU. The type is interface{} with range: 0..4294967295.
    TransmitMtu interface{}

    // Receive MTU. The type is interface{} with range: 0..4294967295.
    ReceiveMtu interface{}

    // VC on VP-tunnel flag. The type is bool.
    VcOnvpTunnel interface{}

    // VC on Point-to-point Sub-interface. The type is bool.
    VcOnP2pSubInterface interface{}

    // TRUE value indicates that the VC is operationally UP. The type is bool.
    OperStatus interface{}

    // TRUE value indicates that the VC is administratively UP. The type is bool.
    AminStatus interface{}

    // VC Internal state. The type is VcState.
    InternalState interface{}

    // Time when VC was last changed. The type is interface{} with range:
    // 0..4294967295.
    LastStateChangeTime interface{}

    // VC test mode. The type is VcTestMode.
    TestMode interface{}

    // Cell packing specific data.
    CellPackingData AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetEntityData() *types.CommonEntityData {
    vc.EntityData.YFilter = vc.YFilter
    vc.EntityData.YangName = "vc"
    vc.EntityData.BundleName = "cisco_ios_xr"
    vc.EntityData.ParentYangName = "vcs"
    vc.EntityData.SegmentPath = "vc" + types.AddKeyToken(vc.InterfaceName, "interface-name")
    vc.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/vcs/" + vc.EntityData.SegmentPath
    vc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vc.EntityData.Children = types.NewOrderedMap()
    vc.EntityData.Children.Append("cell-packing-data", types.YChild{"CellPackingData", &vc.CellPackingData})
    vc.EntityData.Leafs = types.NewOrderedMap()
    vc.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", vc.InterfaceName})
    vc.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", vc.Vpi})
    vc.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", vc.Vci})
    vc.EntityData.Leafs.Append("main-interface", types.YLeaf{"MainInterface", vc.MainInterface})
    vc.EntityData.Leafs.Append("sub-interface", types.YLeaf{"SubInterface", vc.SubInterface})
    vc.EntityData.Leafs.Append("vc-interface", types.YLeaf{"VcInterface", vc.VcInterface})
    vc.EntityData.Leafs.Append("vpi-xr", types.YLeaf{"VpiXr", vc.VpiXr})
    vc.EntityData.Leafs.Append("vci-xr", types.YLeaf{"VciXr", vc.VciXr})
    vc.EntityData.Leafs.Append("type", types.YLeaf{"Type", vc.Type})
    vc.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", vc.Encapsulation})
    vc.EntityData.Leafs.Append("shape", types.YLeaf{"Shape", vc.Shape})
    vc.EntityData.Leafs.Append("peak-cell-rate", types.YLeaf{"PeakCellRate", vc.PeakCellRate})
    vc.EntityData.Leafs.Append("sustained-cell-rate", types.YLeaf{"SustainedCellRate", vc.SustainedCellRate})
    vc.EntityData.Leafs.Append("burst-rate", types.YLeaf{"BurstRate", vc.BurstRate})
    vc.EntityData.Leafs.Append("encaps-inherit-level", types.YLeaf{"EncapsInheritLevel", vc.EncapsInheritLevel})
    vc.EntityData.Leafs.Append("qos-inherit-level", types.YLeaf{"QosInheritLevel", vc.QosInheritLevel})
    vc.EntityData.Leafs.Append("transmit-mtu", types.YLeaf{"TransmitMtu", vc.TransmitMtu})
    vc.EntityData.Leafs.Append("receive-mtu", types.YLeaf{"ReceiveMtu", vc.ReceiveMtu})
    vc.EntityData.Leafs.Append("vc-onvp-tunnel", types.YLeaf{"VcOnvpTunnel", vc.VcOnvpTunnel})
    vc.EntityData.Leafs.Append("vc-on-p2p-sub-interface", types.YLeaf{"VcOnP2pSubInterface", vc.VcOnP2pSubInterface})
    vc.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", vc.OperStatus})
    vc.EntityData.Leafs.Append("amin-status", types.YLeaf{"AminStatus", vc.AminStatus})
    vc.EntityData.Leafs.Append("internal-state", types.YLeaf{"InternalState", vc.InternalState})
    vc.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", vc.LastStateChangeTime})
    vc.EntityData.Leafs.Append("test-mode", types.YLeaf{"TestMode", vc.TestMode})

    vc.EntityData.YListKeys = []string {"InterfaceName"}

    return &(vc.EntityData)
}

// AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData
// Cell packing specific data
type AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local configuration of maximum number of cells to be packed per packet. The
    // type is interface{} with range: 0..65535.
    LocalMaxCellsPackedPerPacket interface{}

    // Negotiated value of maximum number of cells to be packed per packed. The
    // type is interface{} with range: 0..65535.
    NegotiatedMaxCellsPackedPerPacket interface{}

    // Maximum cell packing timeout inmicro seconds. The type is interface{} with
    // range: 0..65535. Units are microsecond.
    MaxCellPackedTimeout interface{}
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetEntityData() *types.CommonEntityData {
    cellPackingData.EntityData.YFilter = cellPackingData.YFilter
    cellPackingData.EntityData.YangName = "cell-packing-data"
    cellPackingData.EntityData.BundleName = "cisco_ios_xr"
    cellPackingData.EntityData.ParentYangName = "vc"
    cellPackingData.EntityData.SegmentPath = "cell-packing-data"
    cellPackingData.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/vcs/vc/" + cellPackingData.EntityData.SegmentPath
    cellPackingData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPackingData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPackingData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPackingData.EntityData.Children = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs.Append("local-max-cells-packed-per-packet", types.YLeaf{"LocalMaxCellsPackedPerPacket", cellPackingData.LocalMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("negotiated-max-cells-packed-per-packet", types.YLeaf{"NegotiatedMaxCellsPackedPerPacket", cellPackingData.NegotiatedMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("max-cell-packed-timeout", types.YLeaf{"MaxCellPackedTimeout", cellPackingData.MaxCellPackedTimeout})

    cellPackingData.EntityData.YListKeys = []string {}

    return &(cellPackingData.EntityData)
}

// AtmVcm_Nodes_Node_CellPacks
// Contains all cell packing information for node
type AtmVcm_Nodes_Node_CellPacks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All cell packing information on a node. The type is slice of
    // AtmVcm_Nodes_Node_CellPacks_CellPack.
    CellPack []*AtmVcm_Nodes_Node_CellPacks_CellPack
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetEntityData() *types.CommonEntityData {
    cellPacks.EntityData.YFilter = cellPacks.YFilter
    cellPacks.EntityData.YangName = "cell-packs"
    cellPacks.EntityData.BundleName = "cisco_ios_xr"
    cellPacks.EntityData.ParentYangName = "node"
    cellPacks.EntityData.SegmentPath = "cell-packs"
    cellPacks.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + cellPacks.EntityData.SegmentPath
    cellPacks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPacks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPacks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPacks.EntityData.Children = types.NewOrderedMap()
    cellPacks.EntityData.Children.Append("cell-pack", types.YChild{"CellPack", nil})
    for i := range cellPacks.CellPack {
        types.SetYListKey(cellPacks.CellPack[i], i)
        cellPacks.EntityData.Children.Append(types.GetSegmentPath(cellPacks.CellPack[i]), types.YChild{"CellPack", cellPacks.CellPack[i]})
    }
    cellPacks.EntityData.Leafs = types.NewOrderedMap()

    cellPacks.EntityData.YListKeys = []string {}

    return &(cellPacks.EntityData)
}

// AtmVcm_Nodes_Node_CellPacks_CellPack
// All cell packing information on a node
type AtmVcm_Nodes_Node_CellPacks_CellPack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // PCI. The type is interface{} with range: 0..4294967295.
    Pci interface{}

    // Sub-interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SubInterfaceName interface{}

    // ATM cell packing mode. The type is VcCellPackingMode.
    CellPackingMode interface{}

    // VPI. The type is interface{} with range: 0..4294967295.
    Vpi interface{}

    // VCI. The type is interface{} with range: 0..4294967295.
    Vci interface{}

    // Average cells/packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedAverageCellsPackets interface{}

    // Average cells/packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    SentCellsPackets interface{}

    // Cell packing specific data.
    CellPacking AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetEntityData() *types.CommonEntityData {
    cellPack.EntityData.YFilter = cellPack.YFilter
    cellPack.EntityData.YangName = "cell-pack"
    cellPack.EntityData.BundleName = "cisco_ios_xr"
    cellPack.EntityData.ParentYangName = "cell-packs"
    cellPack.EntityData.SegmentPath = "cell-pack" + types.AddNoKeyToken(cellPack)
    cellPack.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/cell-packs/" + cellPack.EntityData.SegmentPath
    cellPack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPack.EntityData.Children = types.NewOrderedMap()
    cellPack.EntityData.Children.Append("cell-packing", types.YChild{"CellPacking", &cellPack.CellPacking})
    cellPack.EntityData.Leafs = types.NewOrderedMap()
    cellPack.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", cellPack.InterfaceName})
    cellPack.EntityData.Leafs.Append("pci", types.YLeaf{"Pci", cellPack.Pci})
    cellPack.EntityData.Leafs.Append("sub-interface-name", types.YLeaf{"SubInterfaceName", cellPack.SubInterfaceName})
    cellPack.EntityData.Leafs.Append("cell-packing-mode", types.YLeaf{"CellPackingMode", cellPack.CellPackingMode})
    cellPack.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", cellPack.Vpi})
    cellPack.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", cellPack.Vci})
    cellPack.EntityData.Leafs.Append("received-average-cells-packets", types.YLeaf{"ReceivedAverageCellsPackets", cellPack.ReceivedAverageCellsPackets})
    cellPack.EntityData.Leafs.Append("sent-cells-packets", types.YLeaf{"SentCellsPackets", cellPack.SentCellsPackets})

    cellPack.EntityData.YListKeys = []string {}

    return &(cellPack.EntityData)
}

// AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking
// Cell packing specific data
type AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local configuration of maximum number of cells to be packed per packet. The
    // type is interface{} with range: 0..65535.
    LocalMaxCellsPackedPerPacket interface{}

    // Negotiated value of maximum number of cells to be packed per packed. The
    // type is interface{} with range: 0..65535.
    NegotiatedMaxCellsPackedPerPacket interface{}

    // Maximum cell packing timeout inmicro seconds. The type is interface{} with
    // range: 0..65535. Units are microsecond.
    MaxCellPackedTimeout interface{}
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetEntityData() *types.CommonEntityData {
    cellPacking.EntityData.YFilter = cellPacking.YFilter
    cellPacking.EntityData.YangName = "cell-packing"
    cellPacking.EntityData.BundleName = "cisco_ios_xr"
    cellPacking.EntityData.ParentYangName = "cell-pack"
    cellPacking.EntityData.SegmentPath = "cell-packing"
    cellPacking.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/cell-packs/cell-pack/" + cellPacking.EntityData.SegmentPath
    cellPacking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPacking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPacking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPacking.EntityData.Children = types.NewOrderedMap()
    cellPacking.EntityData.Leafs = types.NewOrderedMap()
    cellPacking.EntityData.Leafs.Append("local-max-cells-packed-per-packet", types.YLeaf{"LocalMaxCellsPackedPerPacket", cellPacking.LocalMaxCellsPackedPerPacket})
    cellPacking.EntityData.Leafs.Append("negotiated-max-cells-packed-per-packet", types.YLeaf{"NegotiatedMaxCellsPackedPerPacket", cellPacking.NegotiatedMaxCellsPackedPerPacket})
    cellPacking.EntityData.Leafs.Append("max-cell-packed-timeout", types.YLeaf{"MaxCellPackedTimeout", cellPacking.MaxCellPackedTimeout})

    cellPacking.EntityData.YListKeys = []string {}

    return &(cellPacking.EntityData)
}

// AtmVcm_Nodes_Node_Pvps
// Contains all L2 PVP information for node
type AtmVcm_Nodes_Node_Pvps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All L2 PVP information on a node. The type is slice of
    // AtmVcm_Nodes_Node_Pvps_Pvp.
    Pvp []*AtmVcm_Nodes_Node_Pvps_Pvp
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetEntityData() *types.CommonEntityData {
    pvps.EntityData.YFilter = pvps.YFilter
    pvps.EntityData.YangName = "pvps"
    pvps.EntityData.BundleName = "cisco_ios_xr"
    pvps.EntityData.ParentYangName = "node"
    pvps.EntityData.SegmentPath = "pvps"
    pvps.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + pvps.EntityData.SegmentPath
    pvps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pvps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pvps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pvps.EntityData.Children = types.NewOrderedMap()
    pvps.EntityData.Children.Append("pvp", types.YChild{"Pvp", nil})
    for i := range pvps.Pvp {
        pvps.EntityData.Children.Append(types.GetSegmentPath(pvps.Pvp[i]), types.YChild{"Pvp", pvps.Pvp[i]})
    }
    pvps.EntityData.Leafs = types.NewOrderedMap()

    pvps.EntityData.YListKeys = []string {}

    return &(pvps.EntityData)
}

// AtmVcm_Nodes_Node_Pvps_Pvp
// All L2 PVP information on a node
type AtmVcm_Nodes_Node_Pvps_Pvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: 0..4294967295.
    Vpi interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    MainInterface interface{}

    // Subinterface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SubInterface interface{}

    // VC Interfcace handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    VcInterface interface{}

    // VC VPI value. The type is interface{} with range: 0..65535.
    VpiXr interface{}

    // VC VCI value. The type is interface{} with range: 0..65535.
    VciXr interface{}

    // VC Type. The type is Vc.
    Type interface{}

    // Encapsulation type. The type is VcEncap.
    Encapsulation interface{}

    // ATM VC traffic shaping type. The type is VcTrafShaping.
    Shape interface{}

    // Peak cell rate in Kbps. The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    PeakCellRate interface{}

    // Sustained cell rate in Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    SustainedCellRate interface{}

    // Burst size in cells. The type is interface{} with range: 0..4294967295.
    BurstRate interface{}

    // Encapsulation inherit level - identifies if encapsulation is set to
    // default, configured on the VC, or inherited from the vcclass. The type is
    // VcInheritLevel.
    EncapsInheritLevel interface{}

    // Quality of Service inherit level - identifies if QoS is set to default,
    // configured on the VC, or inherited from the vcclass. The type is
    // VcInheritLevel.
    QosInheritLevel interface{}

    // Transmit MTU. The type is interface{} with range: 0..4294967295.
    TransmitMtu interface{}

    // Receive MTU. The type is interface{} with range: 0..4294967295.
    ReceiveMtu interface{}

    // VC on VP-tunnel flag. The type is bool.
    VcOnvpTunnel interface{}

    // VC on Point-to-point Sub-interface. The type is bool.
    VcOnP2pSubInterface interface{}

    // TRUE value indicates that the VC is operationally UP. The type is bool.
    OperStatus interface{}

    // TRUE value indicates that the VC is administratively UP. The type is bool.
    AminStatus interface{}

    // VC Internal state. The type is VcState.
    InternalState interface{}

    // Time when VC was last changed. The type is interface{} with range:
    // 0..4294967295.
    LastStateChangeTime interface{}

    // VC test mode. The type is VcTestMode.
    TestMode interface{}

    // Cell packing specific data.
    CellPackingData AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetEntityData() *types.CommonEntityData {
    pvp.EntityData.YFilter = pvp.YFilter
    pvp.EntityData.YangName = "pvp"
    pvp.EntityData.BundleName = "cisco_ios_xr"
    pvp.EntityData.ParentYangName = "pvps"
    pvp.EntityData.SegmentPath = "pvp" + types.AddKeyToken(pvp.InterfaceName, "interface-name")
    pvp.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/pvps/" + pvp.EntityData.SegmentPath
    pvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pvp.EntityData.Children = types.NewOrderedMap()
    pvp.EntityData.Children.Append("cell-packing-data", types.YChild{"CellPackingData", &pvp.CellPackingData})
    pvp.EntityData.Leafs = types.NewOrderedMap()
    pvp.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", pvp.InterfaceName})
    pvp.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", pvp.Vpi})
    pvp.EntityData.Leafs.Append("main-interface", types.YLeaf{"MainInterface", pvp.MainInterface})
    pvp.EntityData.Leafs.Append("sub-interface", types.YLeaf{"SubInterface", pvp.SubInterface})
    pvp.EntityData.Leafs.Append("vc-interface", types.YLeaf{"VcInterface", pvp.VcInterface})
    pvp.EntityData.Leafs.Append("vpi-xr", types.YLeaf{"VpiXr", pvp.VpiXr})
    pvp.EntityData.Leafs.Append("vci-xr", types.YLeaf{"VciXr", pvp.VciXr})
    pvp.EntityData.Leafs.Append("type", types.YLeaf{"Type", pvp.Type})
    pvp.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", pvp.Encapsulation})
    pvp.EntityData.Leafs.Append("shape", types.YLeaf{"Shape", pvp.Shape})
    pvp.EntityData.Leafs.Append("peak-cell-rate", types.YLeaf{"PeakCellRate", pvp.PeakCellRate})
    pvp.EntityData.Leafs.Append("sustained-cell-rate", types.YLeaf{"SustainedCellRate", pvp.SustainedCellRate})
    pvp.EntityData.Leafs.Append("burst-rate", types.YLeaf{"BurstRate", pvp.BurstRate})
    pvp.EntityData.Leafs.Append("encaps-inherit-level", types.YLeaf{"EncapsInheritLevel", pvp.EncapsInheritLevel})
    pvp.EntityData.Leafs.Append("qos-inherit-level", types.YLeaf{"QosInheritLevel", pvp.QosInheritLevel})
    pvp.EntityData.Leafs.Append("transmit-mtu", types.YLeaf{"TransmitMtu", pvp.TransmitMtu})
    pvp.EntityData.Leafs.Append("receive-mtu", types.YLeaf{"ReceiveMtu", pvp.ReceiveMtu})
    pvp.EntityData.Leafs.Append("vc-onvp-tunnel", types.YLeaf{"VcOnvpTunnel", pvp.VcOnvpTunnel})
    pvp.EntityData.Leafs.Append("vc-on-p2p-sub-interface", types.YLeaf{"VcOnP2pSubInterface", pvp.VcOnP2pSubInterface})
    pvp.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", pvp.OperStatus})
    pvp.EntityData.Leafs.Append("amin-status", types.YLeaf{"AminStatus", pvp.AminStatus})
    pvp.EntityData.Leafs.Append("internal-state", types.YLeaf{"InternalState", pvp.InternalState})
    pvp.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", pvp.LastStateChangeTime})
    pvp.EntityData.Leafs.Append("test-mode", types.YLeaf{"TestMode", pvp.TestMode})

    pvp.EntityData.YListKeys = []string {"InterfaceName"}

    return &(pvp.EntityData)
}

// AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData
// Cell packing specific data
type AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local configuration of maximum number of cells to be packed per packet. The
    // type is interface{} with range: 0..65535.
    LocalMaxCellsPackedPerPacket interface{}

    // Negotiated value of maximum number of cells to be packed per packed. The
    // type is interface{} with range: 0..65535.
    NegotiatedMaxCellsPackedPerPacket interface{}

    // Maximum cell packing timeout inmicro seconds. The type is interface{} with
    // range: 0..65535. Units are microsecond.
    MaxCellPackedTimeout interface{}
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetEntityData() *types.CommonEntityData {
    cellPackingData.EntityData.YFilter = cellPackingData.YFilter
    cellPackingData.EntityData.YangName = "cell-packing-data"
    cellPackingData.EntityData.BundleName = "cisco_ios_xr"
    cellPackingData.EntityData.ParentYangName = "pvp"
    cellPackingData.EntityData.SegmentPath = "cell-packing-data"
    cellPackingData.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/pvps/pvp/" + cellPackingData.EntityData.SegmentPath
    cellPackingData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPackingData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPackingData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPackingData.EntityData.Children = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs.Append("local-max-cells-packed-per-packet", types.YLeaf{"LocalMaxCellsPackedPerPacket", cellPackingData.LocalMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("negotiated-max-cells-packed-per-packet", types.YLeaf{"NegotiatedMaxCellsPackedPerPacket", cellPackingData.NegotiatedMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("max-cell-packed-timeout", types.YLeaf{"MaxCellPackedTimeout", cellPackingData.MaxCellPackedTimeout})

    cellPackingData.EntityData.YListKeys = []string {}

    return &(cellPackingData.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks
// Contains all class link information for node
type AtmVcm_Nodes_Node_ClassLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All ATM VC information on a node. The type is slice of
    // AtmVcm_Nodes_Node_ClassLinks_ClassLink.
    ClassLink []*AtmVcm_Nodes_Node_ClassLinks_ClassLink
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetEntityData() *types.CommonEntityData {
    classLinks.EntityData.YFilter = classLinks.YFilter
    classLinks.EntityData.YangName = "class-links"
    classLinks.EntityData.BundleName = "cisco_ios_xr"
    classLinks.EntityData.ParentYangName = "node"
    classLinks.EntityData.SegmentPath = "class-links"
    classLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + classLinks.EntityData.SegmentPath
    classLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classLinks.EntityData.Children = types.NewOrderedMap()
    classLinks.EntityData.Children.Append("class-link", types.YChild{"ClassLink", nil})
    for i := range classLinks.ClassLink {
        classLinks.EntityData.Children.Append(types.GetSegmentPath(classLinks.ClassLink[i]), types.YChild{"ClassLink", classLinks.ClassLink[i]})
    }
    classLinks.EntityData.Leafs = types.NewOrderedMap()

    classLinks.EntityData.YListKeys = []string {}

    return &(classLinks.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink
// All ATM VC information on a node
type AtmVcm_Nodes_Node_ClassLinks_ClassLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VPI. The type is interface{} with range:
    // 0..4294967295.
    Vpi interface{}

    // VCI. The type is interface{} with range: 0..4294967295.
    Vci interface{}

    // Sub-interface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SubInterfaceName interface{}

    // Not supported VC class.
    VcClassNotSupported AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported

    // Oam values for class link.
    OamConfig AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetEntityData() *types.CommonEntityData {
    classLink.EntityData.YFilter = classLink.YFilter
    classLink.EntityData.YangName = "class-link"
    classLink.EntityData.BundleName = "cisco_ios_xr"
    classLink.EntityData.ParentYangName = "class-links"
    classLink.EntityData.SegmentPath = "class-link" + types.AddKeyToken(classLink.Vpi, "vpi")
    classLink.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/" + classLink.EntityData.SegmentPath
    classLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classLink.EntityData.Children = types.NewOrderedMap()
    classLink.EntityData.Children.Append("vc-class-not-supported", types.YChild{"VcClassNotSupported", &classLink.VcClassNotSupported})
    classLink.EntityData.Children.Append("oam-config", types.YChild{"OamConfig", &classLink.OamConfig})
    classLink.EntityData.Leafs = types.NewOrderedMap()
    classLink.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", classLink.Vpi})
    classLink.EntityData.Leafs.Append("vci", types.YLeaf{"Vci", classLink.Vci})
    classLink.EntityData.Leafs.Append("sub-interface-name", types.YLeaf{"SubInterfaceName", classLink.SubInterfaceName})

    classLink.EntityData.YListKeys = []string {"Vpi"}

    return &(classLink.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported
// Not supported VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation type not supported. The type is VcEncap.
    EncapsulationNotSupported interface{}

    // NotSupportedInheritLevel. The type is VcInheritLevel.
    NotSupportedInheritLevel interface{}
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetEntityData() *types.CommonEntityData {
    vcClassNotSupported.EntityData.YFilter = vcClassNotSupported.YFilter
    vcClassNotSupported.EntityData.YangName = "vc-class-not-supported"
    vcClassNotSupported.EntityData.BundleName = "cisco_ios_xr"
    vcClassNotSupported.EntityData.ParentYangName = "class-link"
    vcClassNotSupported.EntityData.SegmentPath = "vc-class-not-supported"
    vcClassNotSupported.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/" + vcClassNotSupported.EntityData.SegmentPath
    vcClassNotSupported.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vcClassNotSupported.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vcClassNotSupported.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vcClassNotSupported.EntityData.Children = types.NewOrderedMap()
    vcClassNotSupported.EntityData.Leafs = types.NewOrderedMap()
    vcClassNotSupported.EntityData.Leafs.Append("encapsulation-not-supported", types.YLeaf{"EncapsulationNotSupported", vcClassNotSupported.EncapsulationNotSupported})
    vcClassNotSupported.EntityData.Leafs.Append("not-supported-inherit-level", types.YLeaf{"NotSupportedInheritLevel", vcClassNotSupported.NotSupportedInheritLevel})

    vcClassNotSupported.EntityData.YListKeys = []string {}

    return &(vcClassNotSupported.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig
// Oam values for class link
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Traffic Shaping detail of VC class.
    ClassLinkShaping AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping

    // Encapsulation details of VC class.
    ClassLinkEncapsulation AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation

    // OAM PVC details of VC class.
    OamPvc AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc

    // OAM Retry details of VC class.
    OamRetry AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry

    // AIS RDI details of a VC class.
    AisRdi AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetEntityData() *types.CommonEntityData {
    oamConfig.EntityData.YFilter = oamConfig.YFilter
    oamConfig.EntityData.YangName = "oam-config"
    oamConfig.EntityData.BundleName = "cisco_ios_xr"
    oamConfig.EntityData.ParentYangName = "class-link"
    oamConfig.EntityData.SegmentPath = "oam-config"
    oamConfig.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/" + oamConfig.EntityData.SegmentPath
    oamConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oamConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oamConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oamConfig.EntityData.Children = types.NewOrderedMap()
    oamConfig.EntityData.Children.Append("class-link-shaping", types.YChild{"ClassLinkShaping", &oamConfig.ClassLinkShaping})
    oamConfig.EntityData.Children.Append("class-link-encapsulation", types.YChild{"ClassLinkEncapsulation", &oamConfig.ClassLinkEncapsulation})
    oamConfig.EntityData.Children.Append("oam-pvc", types.YChild{"OamPvc", &oamConfig.OamPvc})
    oamConfig.EntityData.Children.Append("oam-retry", types.YChild{"OamRetry", &oamConfig.OamRetry})
    oamConfig.EntityData.Children.Append("ais-rdi", types.YChild{"AisRdi", &oamConfig.AisRdi})
    oamConfig.EntityData.Leafs = types.NewOrderedMap()

    oamConfig.EntityData.YListKeys = []string {}

    return &(oamConfig.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping
// Traffic Shaping detail of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ATM VC traffic shaping type. The type is VcTrafShaping.
    ShapingType interface{}

    // Peak output rate in Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    PeakOutputRate interface{}

    // Average output rate. The type is interface{} with range: 0..4294967295.
    AverageOutputRate interface{}

    // Burst output rate. The type is interface{} with range: 0..4294967295.
    BurstOutputRate interface{}

    // Shaping inherit level. The type is VcInheritLevel.
    ShapingInheritLevel interface{}
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetEntityData() *types.CommonEntityData {
    classLinkShaping.EntityData.YFilter = classLinkShaping.YFilter
    classLinkShaping.EntityData.YangName = "class-link-shaping"
    classLinkShaping.EntityData.BundleName = "cisco_ios_xr"
    classLinkShaping.EntityData.ParentYangName = "oam-config"
    classLinkShaping.EntityData.SegmentPath = "class-link-shaping"
    classLinkShaping.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/oam-config/" + classLinkShaping.EntityData.SegmentPath
    classLinkShaping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classLinkShaping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classLinkShaping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classLinkShaping.EntityData.Children = types.NewOrderedMap()
    classLinkShaping.EntityData.Leafs = types.NewOrderedMap()
    classLinkShaping.EntityData.Leafs.Append("shaping-type", types.YLeaf{"ShapingType", classLinkShaping.ShapingType})
    classLinkShaping.EntityData.Leafs.Append("peak-output-rate", types.YLeaf{"PeakOutputRate", classLinkShaping.PeakOutputRate})
    classLinkShaping.EntityData.Leafs.Append("average-output-rate", types.YLeaf{"AverageOutputRate", classLinkShaping.AverageOutputRate})
    classLinkShaping.EntityData.Leafs.Append("burst-output-rate", types.YLeaf{"BurstOutputRate", classLinkShaping.BurstOutputRate})
    classLinkShaping.EntityData.Leafs.Append("shaping-inherit-level", types.YLeaf{"ShapingInheritLevel", classLinkShaping.ShapingInheritLevel})

    classLinkShaping.EntityData.YListKeys = []string {}

    return &(classLinkShaping.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation
// Encapsulation details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encapsulation type. The type is VcEncap.
    EncapsulationType interface{}

    // Encapsulation inherit level. The type is VcInheritLevel.
    EncapsulationInheritLevel interface{}
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetEntityData() *types.CommonEntityData {
    classLinkEncapsulation.EntityData.YFilter = classLinkEncapsulation.YFilter
    classLinkEncapsulation.EntityData.YangName = "class-link-encapsulation"
    classLinkEncapsulation.EntityData.BundleName = "cisco_ios_xr"
    classLinkEncapsulation.EntityData.ParentYangName = "oam-config"
    classLinkEncapsulation.EntityData.SegmentPath = "class-link-encapsulation"
    classLinkEncapsulation.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/oam-config/" + classLinkEncapsulation.EntityData.SegmentPath
    classLinkEncapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classLinkEncapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classLinkEncapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classLinkEncapsulation.EntityData.Children = types.NewOrderedMap()
    classLinkEncapsulation.EntityData.Leafs = types.NewOrderedMap()
    classLinkEncapsulation.EntityData.Leafs.Append("encapsulation-type", types.YLeaf{"EncapsulationType", classLinkEncapsulation.EncapsulationType})
    classLinkEncapsulation.EntityData.Leafs.Append("encapsulation-inherit-level", types.YLeaf{"EncapsulationInheritLevel", classLinkEncapsulation.EncapsulationInheritLevel})

    classLinkEncapsulation.EntityData.YListKeys = []string {}

    return &(classLinkEncapsulation.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc
// OAM PVC details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Manage Level. The type is VcManageLevel.
    ManageLevel interface{}

    // PVC Frequency. The type is interface{} with range: 0..4294967295.
    PvcFrequency interface{}

    // Keep vc up. The type is bool.
    KeepVcUp interface{}

    // AIS RDI failure. The type is bool.
    AisRdiFailure interface{}

    // Manage inherit level. The type is ClassLinkOamInheritLevel.
    ManageInheritLevel interface{}
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetEntityData() *types.CommonEntityData {
    oamPvc.EntityData.YFilter = oamPvc.YFilter
    oamPvc.EntityData.YangName = "oam-pvc"
    oamPvc.EntityData.BundleName = "cisco_ios_xr"
    oamPvc.EntityData.ParentYangName = "oam-config"
    oamPvc.EntityData.SegmentPath = "oam-pvc"
    oamPvc.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/oam-config/" + oamPvc.EntityData.SegmentPath
    oamPvc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oamPvc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oamPvc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oamPvc.EntityData.Children = types.NewOrderedMap()
    oamPvc.EntityData.Leafs = types.NewOrderedMap()
    oamPvc.EntityData.Leafs.Append("manage-level", types.YLeaf{"ManageLevel", oamPvc.ManageLevel})
    oamPvc.EntityData.Leafs.Append("pvc-frequency", types.YLeaf{"PvcFrequency", oamPvc.PvcFrequency})
    oamPvc.EntityData.Leafs.Append("keep-vc-up", types.YLeaf{"KeepVcUp", oamPvc.KeepVcUp})
    oamPvc.EntityData.Leafs.Append("ais-rdi-failure", types.YLeaf{"AisRdiFailure", oamPvc.AisRdiFailure})
    oamPvc.EntityData.Leafs.Append("manage-inherit-level", types.YLeaf{"ManageInheritLevel", oamPvc.ManageInheritLevel})

    oamPvc.EntityData.YListKeys = []string {}

    return &(oamPvc.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry
// OAM Retry details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Retry Count. The type is interface{} with range: 0..4294967295.
    RetryUpCount interface{}

    // Down Count. The type is interface{} with range: 0..4294967295.
    DownCount interface{}

    // Retry frequency. The type is interface{} with range: 0..4294967295.
    RetryFrequency interface{}

    // Retry inherit level. The type is ClassLinkOamInheritLevel.
    RetryInheritLevel interface{}
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetEntityData() *types.CommonEntityData {
    oamRetry.EntityData.YFilter = oamRetry.YFilter
    oamRetry.EntityData.YangName = "oam-retry"
    oamRetry.EntityData.BundleName = "cisco_ios_xr"
    oamRetry.EntityData.ParentYangName = "oam-config"
    oamRetry.EntityData.SegmentPath = "oam-retry"
    oamRetry.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/oam-config/" + oamRetry.EntityData.SegmentPath
    oamRetry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oamRetry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oamRetry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oamRetry.EntityData.Children = types.NewOrderedMap()
    oamRetry.EntityData.Leafs = types.NewOrderedMap()
    oamRetry.EntityData.Leafs.Append("retry-up-count", types.YLeaf{"RetryUpCount", oamRetry.RetryUpCount})
    oamRetry.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", oamRetry.DownCount})
    oamRetry.EntityData.Leafs.Append("retry-frequency", types.YLeaf{"RetryFrequency", oamRetry.RetryFrequency})
    oamRetry.EntityData.Leafs.Append("retry-inherit-level", types.YLeaf{"RetryInheritLevel", oamRetry.RetryInheritLevel})

    oamRetry.EntityData.YListKeys = []string {}

    return &(oamRetry.EntityData)
}

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi
// AIS RDI details of a VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS RDI up count. The type is interface{} with range: 0..4294967295.
    AisRdiUpCount interface{}

    // Time (in seconds) with no AIS/RDI cells before declaring a VC as up. The
    // type is interface{} with range: 0..4294967295. Units are second.
    AisRdiUpTime interface{}

    // AIS RDI inherit level. The type is ClassLinkOamInheritLevel.
    AisRdiInheritLevel interface{}
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetEntityData() *types.CommonEntityData {
    aisRdi.EntityData.YFilter = aisRdi.YFilter
    aisRdi.EntityData.YangName = "ais-rdi"
    aisRdi.EntityData.BundleName = "cisco_ios_xr"
    aisRdi.EntityData.ParentYangName = "oam-config"
    aisRdi.EntityData.SegmentPath = "ais-rdi"
    aisRdi.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/class-links/class-link/oam-config/" + aisRdi.EntityData.SegmentPath
    aisRdi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aisRdi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aisRdi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aisRdi.EntityData.Children = types.NewOrderedMap()
    aisRdi.EntityData.Leafs = types.NewOrderedMap()
    aisRdi.EntityData.Leafs.Append("ais-rdi-up-count", types.YLeaf{"AisRdiUpCount", aisRdi.AisRdiUpCount})
    aisRdi.EntityData.Leafs.Append("ais-rdi-up-time", types.YLeaf{"AisRdiUpTime", aisRdi.AisRdiUpTime})
    aisRdi.EntityData.Leafs.Append("ais-rdi-inherit-level", types.YLeaf{"AisRdiInheritLevel", aisRdi.AisRdiInheritLevel})

    aisRdi.EntityData.YListKeys = []string {}

    return &(aisRdi.EntityData)
}

// AtmVcm_Nodes_Node_Interfaces
// Contains all Interface information for node
type AtmVcm_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ATM Interface data. The type is slice of
    // AtmVcm_Nodes_Node_Interfaces_Interface.
    Interface []*AtmVcm_Nodes_Node_Interfaces_Interface
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// AtmVcm_Nodes_Node_Interfaces_Interface
// ATM Interface data
type AtmVcm_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // ILMI VPI. The type is interface{} with range: 0..4294967295.
    IlmiVpi interface{}

    // ILMI VCI. The type is interface{} with range: 0..4294967295.
    IlmiVci interface{}

    // Number of PVC Failures. The type is interface{} with range: 0..4294967295.
    PvcFailures interface{}

    // Number of currently failing Layer 2 PVPs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer2pvPs interface{}

    // Number of currently failing Layer 2 PVCs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer2pvCs interface{}

    // Number of currently failing Layer 3 VP tunnels. The type is interface{}
    // with range: 0..4294967295.
    CurrentlyFailingLayer3vpTunnels interface{}

    // Number of currently failing Layer 3 PVCs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer3pvCs interface{}

    // If true, PVC failures trap is enabled. The type is bool.
    PvcFailuresTrapEnable interface{}

    // PVC trap notification interval. The type is interface{} with range:
    // 0..4294967295.
    PvcNotificationInterval interface{}

    // Number of Layer 2 PVPs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer2pvPs interface{}

    // Number of Layer 2 PVCs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer2pvCs interface{}

    // Number of Layer 3 VP tunnels configured. The type is interface{} with
    // range: 0..4294967295.
    ConfiguredLayer3vpTunnels interface{}

    // Number of Layer 3 PVCs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer3pvCs interface{}

    // ATM interface port type. The type is VcmPort.
    PortType interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    MainInterface interface{}

    // Number of L2 attachment circuits with the cell packing feature enabled on
    // this main interface. The type is interface{} with range: 0..65535.
    L2CellPackingCount interface{}

    // Cell packing specific information.
    CellPackingData AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("cell-packing-data", types.YChild{"CellPackingData", &self.CellPackingData})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("ilmi-vpi", types.YLeaf{"IlmiVpi", self.IlmiVpi})
    self.EntityData.Leafs.Append("ilmi-vci", types.YLeaf{"IlmiVci", self.IlmiVci})
    self.EntityData.Leafs.Append("pvc-failures", types.YLeaf{"PvcFailures", self.PvcFailures})
    self.EntityData.Leafs.Append("currently-failing-layer2pv-ps", types.YLeaf{"CurrentlyFailingLayer2pvPs", self.CurrentlyFailingLayer2pvPs})
    self.EntityData.Leafs.Append("currently-failing-layer2pv-cs", types.YLeaf{"CurrentlyFailingLayer2pvCs", self.CurrentlyFailingLayer2pvCs})
    self.EntityData.Leafs.Append("currently-failing-layer3vp-tunnels", types.YLeaf{"CurrentlyFailingLayer3vpTunnels", self.CurrentlyFailingLayer3vpTunnels})
    self.EntityData.Leafs.Append("currently-failing-layer3pv-cs", types.YLeaf{"CurrentlyFailingLayer3pvCs", self.CurrentlyFailingLayer3pvCs})
    self.EntityData.Leafs.Append("pvc-failures-trap-enable", types.YLeaf{"PvcFailuresTrapEnable", self.PvcFailuresTrapEnable})
    self.EntityData.Leafs.Append("pvc-notification-interval", types.YLeaf{"PvcNotificationInterval", self.PvcNotificationInterval})
    self.EntityData.Leafs.Append("configured-layer2pv-ps", types.YLeaf{"ConfiguredLayer2pvPs", self.ConfiguredLayer2pvPs})
    self.EntityData.Leafs.Append("configured-layer2pv-cs", types.YLeaf{"ConfiguredLayer2pvCs", self.ConfiguredLayer2pvCs})
    self.EntityData.Leafs.Append("configured-layer3vp-tunnels", types.YLeaf{"ConfiguredLayer3vpTunnels", self.ConfiguredLayer3vpTunnels})
    self.EntityData.Leafs.Append("configured-layer3pv-cs", types.YLeaf{"ConfiguredLayer3pvCs", self.ConfiguredLayer3pvCs})
    self.EntityData.Leafs.Append("port-type", types.YLeaf{"PortType", self.PortType})
    self.EntityData.Leafs.Append("main-interface", types.YLeaf{"MainInterface", self.MainInterface})
    self.EntityData.Leafs.Append("l2-cell-packing-count", types.YLeaf{"L2CellPackingCount", self.L2CellPackingCount})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData
// Cell packing specific information
type AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local configuration of maximum number of cells to be packed per packet. The
    // type is interface{} with range: 0..65535.
    LocalMaxCellsPackedPerPacket interface{}

    // Negotiated value of maximum number of cells to be packed per packed. The
    // type is interface{} with range: 0..65535.
    NegotiatedMaxCellsPackedPerPacket interface{}

    // Maximum cell packing timeout inmicro seconds. The type is interface{} with
    // range: 0..65535. Units are microsecond.
    MaxCellPackedTimeout interface{}
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetEntityData() *types.CommonEntityData {
    cellPackingData.EntityData.YFilter = cellPackingData.YFilter
    cellPackingData.EntityData.YangName = "cell-packing-data"
    cellPackingData.EntityData.BundleName = "cisco_ios_xr"
    cellPackingData.EntityData.ParentYangName = "interface"
    cellPackingData.EntityData.SegmentPath = "cell-packing-data"
    cellPackingData.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/interfaces/interface/" + cellPackingData.EntityData.SegmentPath
    cellPackingData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cellPackingData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cellPackingData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cellPackingData.EntityData.Children = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs = types.NewOrderedMap()
    cellPackingData.EntityData.Leafs.Append("local-max-cells-packed-per-packet", types.YLeaf{"LocalMaxCellsPackedPerPacket", cellPackingData.LocalMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("negotiated-max-cells-packed-per-packet", types.YLeaf{"NegotiatedMaxCellsPackedPerPacket", cellPackingData.NegotiatedMaxCellsPackedPerPacket})
    cellPackingData.EntityData.Leafs.Append("max-cell-packed-timeout", types.YLeaf{"MaxCellPackedTimeout", cellPackingData.MaxCellPackedTimeout})

    cellPackingData.EntityData.YListKeys = []string {}

    return &(cellPackingData.EntityData)
}

// AtmVcm_Nodes_Node_VpTunnels
// Contains all VP-tunnel information for node
type AtmVcm_Nodes_Node_VpTunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All VP-tunnel information on a node. The type is slice of
    // AtmVcm_Nodes_Node_VpTunnels_VpTunnel.
    VpTunnel []*AtmVcm_Nodes_Node_VpTunnels_VpTunnel
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetEntityData() *types.CommonEntityData {
    vpTunnels.EntityData.YFilter = vpTunnels.YFilter
    vpTunnels.EntityData.YangName = "vp-tunnels"
    vpTunnels.EntityData.BundleName = "cisco_ios_xr"
    vpTunnels.EntityData.ParentYangName = "node"
    vpTunnels.EntityData.SegmentPath = "vp-tunnels"
    vpTunnels.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/" + vpTunnels.EntityData.SegmentPath
    vpTunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpTunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpTunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpTunnels.EntityData.Children = types.NewOrderedMap()
    vpTunnels.EntityData.Children.Append("vp-tunnel", types.YChild{"VpTunnel", nil})
    for i := range vpTunnels.VpTunnel {
        vpTunnels.EntityData.Children.Append(types.GetSegmentPath(vpTunnels.VpTunnel[i]), types.YChild{"VpTunnel", vpTunnels.VpTunnel[i]})
    }
    vpTunnels.EntityData.Leafs = types.NewOrderedMap()

    vpTunnels.EntityData.YListKeys = []string {}

    return &(vpTunnels.EntityData)
}

// AtmVcm_Nodes_Node_VpTunnels_VpTunnel
// All VP-tunnel information on a node
type AtmVcm_Nodes_Node_VpTunnels_VpTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: 0..4294967295.
    Vpi interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    MainInterface interface{}

    // VP Interfcace handle. The type is string with pattern: [a-zA-Z0-9._/-]+.
    VpInterface interface{}

    // VP-Tunnel VPI value. The type is interface{} with range: 0..65535.
    VpiXr interface{}

    // ATM VP traffic shaping type. The type is VpTrafShaping.
    Shape interface{}

    // Peak cell rate in Kbps. The type is interface{} with range: 0..4294967295.
    // Units are kbit/s.
    PeakCellRate interface{}

    // Sustained cell rate in Kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    SustainedCellRate interface{}

    // Burst size in cells. The type is interface{} with range: 0..4294967295.
    BurstRate interface{}

    // F4OAM Enabled flag. The type is bool.
    F4oamEnabled interface{}

    // Number of Data PVCs under this VP-tunnel. The type is interface{} with
    // range: 0..4294967295.
    DataVcCount interface{}

    // TRUE value indicates that the VP is operationally UP. The type is bool.
    OperStatus interface{}

    // TRUE value indicates that the VP is administratively UP. The type is bool.
    AminStatus interface{}

    // Internal state. The type is VpState.
    InternalState interface{}

    // Time when VP-Tunnel state was last changed. The type is interface{} with
    // range: 0..4294967295.
    LastVpStateChangeTime interface{}
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetEntityData() *types.CommonEntityData {
    vpTunnel.EntityData.YFilter = vpTunnel.YFilter
    vpTunnel.EntityData.YangName = "vp-tunnel"
    vpTunnel.EntityData.BundleName = "cisco_ios_xr"
    vpTunnel.EntityData.ParentYangName = "vp-tunnels"
    vpTunnel.EntityData.SegmentPath = "vp-tunnel" + types.AddKeyToken(vpTunnel.InterfaceName, "interface-name")
    vpTunnel.EntityData.AbsolutePath = "Cisco-IOS-XR-atm-vcm-oper:atm-vcm/nodes/node/vp-tunnels/" + vpTunnel.EntityData.SegmentPath
    vpTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpTunnel.EntityData.Children = types.NewOrderedMap()
    vpTunnel.EntityData.Leafs = types.NewOrderedMap()
    vpTunnel.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", vpTunnel.InterfaceName})
    vpTunnel.EntityData.Leafs.Append("vpi", types.YLeaf{"Vpi", vpTunnel.Vpi})
    vpTunnel.EntityData.Leafs.Append("main-interface", types.YLeaf{"MainInterface", vpTunnel.MainInterface})
    vpTunnel.EntityData.Leafs.Append("vp-interface", types.YLeaf{"VpInterface", vpTunnel.VpInterface})
    vpTunnel.EntityData.Leafs.Append("vpi-xr", types.YLeaf{"VpiXr", vpTunnel.VpiXr})
    vpTunnel.EntityData.Leafs.Append("shape", types.YLeaf{"Shape", vpTunnel.Shape})
    vpTunnel.EntityData.Leafs.Append("peak-cell-rate", types.YLeaf{"PeakCellRate", vpTunnel.PeakCellRate})
    vpTunnel.EntityData.Leafs.Append("sustained-cell-rate", types.YLeaf{"SustainedCellRate", vpTunnel.SustainedCellRate})
    vpTunnel.EntityData.Leafs.Append("burst-rate", types.YLeaf{"BurstRate", vpTunnel.BurstRate})
    vpTunnel.EntityData.Leafs.Append("f4oam-enabled", types.YLeaf{"F4oamEnabled", vpTunnel.F4oamEnabled})
    vpTunnel.EntityData.Leafs.Append("data-vc-count", types.YLeaf{"DataVcCount", vpTunnel.DataVcCount})
    vpTunnel.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", vpTunnel.OperStatus})
    vpTunnel.EntityData.Leafs.Append("amin-status", types.YLeaf{"AminStatus", vpTunnel.AminStatus})
    vpTunnel.EntityData.Leafs.Append("internal-state", types.YLeaf{"InternalState", vpTunnel.InternalState})
    vpTunnel.EntityData.Leafs.Append("last-vp-state-change-time", types.YLeaf{"LastVpStateChangeTime", vpTunnel.LastVpStateChangeTime})

    vpTunnel.EntityData.YListKeys = []string {"InterfaceName"}

    return &(vpTunnel.EntityData)
}

