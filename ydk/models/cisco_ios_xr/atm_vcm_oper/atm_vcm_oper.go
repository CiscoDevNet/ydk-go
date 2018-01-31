// This module contains a collection of YANG definitions
// for Cisco IOS-XR atm-vcm package operational data.
// 
// This module contains definitions
// for the following management objects:
//   atm-vcm: ATM VCM operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains all the nodes.
    Nodes AtmVcm_Nodes
}

func (atmVcm *AtmVcm) GetFilter() yfilter.YFilter { return atmVcm.YFilter }

func (atmVcm *AtmVcm) SetFilter(yf yfilter.YFilter) { atmVcm.YFilter = yf }

func (atmVcm *AtmVcm) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (atmVcm *AtmVcm) GetSegmentPath() string {
    return "Cisco-IOS-XR-atm-vcm-oper:atm-vcm"
}

func (atmVcm *AtmVcm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &atmVcm.Nodes
    }
    return nil
}

func (atmVcm *AtmVcm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &atmVcm.Nodes
    return children
}

func (atmVcm *AtmVcm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmVcm *AtmVcm) GetBundleName() string { return "cisco_ios_xr" }

func (atmVcm *AtmVcm) GetYangName() string { return "atm-vcm" }

func (atmVcm *AtmVcm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (atmVcm *AtmVcm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (atmVcm *AtmVcm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (atmVcm *AtmVcm) SetParent(parent types.Entity) { atmVcm.parent = parent }

func (atmVcm *AtmVcm) GetParent() types.Entity { return atmVcm.parent }

func (atmVcm *AtmVcm) GetParentYangName() string { return "Cisco-IOS-XR-atm-vcm-oper" }

// AtmVcm_Nodes
// Contains all the nodes
type AtmVcm_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The node on which ATM Interfaces/VCs/VPs are located. The type is slice of
    // AtmVcm_Nodes_Node.
    Node []AtmVcm_Nodes_Node
}

func (nodes *AtmVcm_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *AtmVcm_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *AtmVcm_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *AtmVcm_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *AtmVcm_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *AtmVcm_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *AtmVcm_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *AtmVcm_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *AtmVcm_Nodes) GetYangName() string { return "nodes" }

func (nodes *AtmVcm_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *AtmVcm_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *AtmVcm_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *AtmVcm_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *AtmVcm_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *AtmVcm_Nodes) GetParentYangName() string { return "atm-vcm" }

// AtmVcm_Nodes_Node
// The node on which ATM Interfaces/VCs/VPs are
// located
type AtmVcm_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (node *AtmVcm_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *AtmVcm_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *AtmVcm_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "vcs" { return "Vcs" }
    if yname == "cell-packs" { return "CellPacks" }
    if yname == "pvps" { return "Pvps" }
    if yname == "class-links" { return "ClassLinks" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "vp-tunnels" { return "VpTunnels" }
    return ""
}

func (node *AtmVcm_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *AtmVcm_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vcs" {
        return &node.Vcs
    }
    if childYangName == "cell-packs" {
        return &node.CellPacks
    }
    if childYangName == "pvps" {
        return &node.Pvps
    }
    if childYangName == "class-links" {
        return &node.ClassLinks
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "vp-tunnels" {
        return &node.VpTunnels
    }
    return nil
}

func (node *AtmVcm_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vcs"] = &node.Vcs
    children["cell-packs"] = &node.CellPacks
    children["pvps"] = &node.Pvps
    children["class-links"] = &node.ClassLinks
    children["interfaces"] = &node.Interfaces
    children["vp-tunnels"] = &node.VpTunnels
    return children
}

func (node *AtmVcm_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *AtmVcm_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *AtmVcm_Nodes_Node) GetYangName() string { return "node" }

func (node *AtmVcm_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *AtmVcm_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *AtmVcm_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *AtmVcm_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *AtmVcm_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *AtmVcm_Nodes_Node) GetParentYangName() string { return "nodes" }

// AtmVcm_Nodes_Node_Vcs
// Contains all VC information for node
type AtmVcm_Nodes_Node_Vcs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All VC information on a node. The type is slice of
    // AtmVcm_Nodes_Node_Vcs_Vc.
    Vc []AtmVcm_Nodes_Node_Vcs_Vc
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetFilter() yfilter.YFilter { return vcs.YFilter }

func (vcs *AtmVcm_Nodes_Node_Vcs) SetFilter(yf yfilter.YFilter) { vcs.YFilter = yf }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetGoName(yname string) string {
    if yname == "vc" { return "Vc" }
    return ""
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetSegmentPath() string {
    return "vcs"
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vc" {
        for _, c := range vcs.Vc {
            if vcs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_Vcs_Vc{}
        vcs.Vc = append(vcs.Vc, child)
        return &vcs.Vc[len(vcs.Vc)-1]
    }
    return nil
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vcs.Vc {
        children[vcs.Vc[i].GetSegmentPath()] = &vcs.Vc[i]
    }
    return children
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vcs *AtmVcm_Nodes_Node_Vcs) GetBundleName() string { return "cisco_ios_xr" }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetYangName() string { return "vcs" }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vcs *AtmVcm_Nodes_Node_Vcs) SetParent(parent types.Entity) { vcs.parent = parent }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetParent() types.Entity { return vcs.parent }

func (vcs *AtmVcm_Nodes_Node_Vcs) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_Vcs_Vc
// All VC information on a node
type AtmVcm_Nodes_Node_Vcs_Vc struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: 0..4095.
    Vpi interface{}

    // VCI. The type is interface{} with range: 1..65535.
    Vci interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    MainInterface interface{}

    // Subinterface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    SubInterface interface{}

    // VC Interfcace handle. The type is string with pattern: [a-zA-Z0-9./-]+.
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
    VcOnP2PSubInterface interface{}

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

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetFilter() yfilter.YFilter { return vc.YFilter }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) SetFilter(yf yfilter.YFilter) { vc.YFilter = yf }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "vpi" { return "Vpi" }
    if yname == "vci" { return "Vci" }
    if yname == "main-interface" { return "MainInterface" }
    if yname == "sub-interface" { return "SubInterface" }
    if yname == "vc-interface" { return "VcInterface" }
    if yname == "vpi-xr" { return "VpiXr" }
    if yname == "vci-xr" { return "VciXr" }
    if yname == "type" { return "Type" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "shape" { return "Shape" }
    if yname == "peak-cell-rate" { return "PeakCellRate" }
    if yname == "sustained-cell-rate" { return "SustainedCellRate" }
    if yname == "burst-rate" { return "BurstRate" }
    if yname == "encaps-inherit-level" { return "EncapsInheritLevel" }
    if yname == "qos-inherit-level" { return "QosInheritLevel" }
    if yname == "transmit-mtu" { return "TransmitMtu" }
    if yname == "receive-mtu" { return "ReceiveMtu" }
    if yname == "vc-onvp-tunnel" { return "VcOnvpTunnel" }
    if yname == "vc-on-p2p-sub-interface" { return "VcOnP2PSubInterface" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "amin-status" { return "AminStatus" }
    if yname == "internal-state" { return "InternalState" }
    if yname == "last-state-change-time" { return "LastStateChangeTime" }
    if yname == "test-mode" { return "TestMode" }
    if yname == "cell-packing-data" { return "CellPackingData" }
    return ""
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetSegmentPath() string {
    return "vc" + "[interface-name='" + fmt.Sprintf("%v", vc.InterfaceName) + "']"
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cell-packing-data" {
        return &vc.CellPackingData
    }
    return nil
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cell-packing-data"] = &vc.CellPackingData
    return children
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = vc.InterfaceName
    leafs["vpi"] = vc.Vpi
    leafs["vci"] = vc.Vci
    leafs["main-interface"] = vc.MainInterface
    leafs["sub-interface"] = vc.SubInterface
    leafs["vc-interface"] = vc.VcInterface
    leafs["vpi-xr"] = vc.VpiXr
    leafs["vci-xr"] = vc.VciXr
    leafs["type"] = vc.Type
    leafs["encapsulation"] = vc.Encapsulation
    leafs["shape"] = vc.Shape
    leafs["peak-cell-rate"] = vc.PeakCellRate
    leafs["sustained-cell-rate"] = vc.SustainedCellRate
    leafs["burst-rate"] = vc.BurstRate
    leafs["encaps-inherit-level"] = vc.EncapsInheritLevel
    leafs["qos-inherit-level"] = vc.QosInheritLevel
    leafs["transmit-mtu"] = vc.TransmitMtu
    leafs["receive-mtu"] = vc.ReceiveMtu
    leafs["vc-onvp-tunnel"] = vc.VcOnvpTunnel
    leafs["vc-on-p2p-sub-interface"] = vc.VcOnP2PSubInterface
    leafs["oper-status"] = vc.OperStatus
    leafs["amin-status"] = vc.AminStatus
    leafs["internal-state"] = vc.InternalState
    leafs["last-state-change-time"] = vc.LastStateChangeTime
    leafs["test-mode"] = vc.TestMode
    return leafs
}

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetBundleName() string { return "cisco_ios_xr" }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetYangName() string { return "vc" }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) SetParent(parent types.Entity) { vc.parent = parent }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetParent() types.Entity { return vc.parent }

func (vc *AtmVcm_Nodes_Node_Vcs_Vc) GetParentYangName() string { return "vcs" }

// AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData
// Cell packing specific data
type AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData struct {
    parent types.Entity
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

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetFilter() yfilter.YFilter { return cellPackingData.YFilter }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) SetFilter(yf yfilter.YFilter) { cellPackingData.YFilter = yf }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetGoName(yname string) string {
    if yname == "local-max-cells-packed-per-packet" { return "LocalMaxCellsPackedPerPacket" }
    if yname == "negotiated-max-cells-packed-per-packet" { return "NegotiatedMaxCellsPackedPerPacket" }
    if yname == "max-cell-packed-timeout" { return "MaxCellPackedTimeout" }
    return ""
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetSegmentPath() string {
    return "cell-packing-data"
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-max-cells-packed-per-packet"] = cellPackingData.LocalMaxCellsPackedPerPacket
    leafs["negotiated-max-cells-packed-per-packet"] = cellPackingData.NegotiatedMaxCellsPackedPerPacket
    leafs["max-cell-packed-timeout"] = cellPackingData.MaxCellPackedTimeout
    return leafs
}

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetBundleName() string { return "cisco_ios_xr" }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetYangName() string { return "cell-packing-data" }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) SetParent(parent types.Entity) { cellPackingData.parent = parent }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetParent() types.Entity { return cellPackingData.parent }

func (cellPackingData *AtmVcm_Nodes_Node_Vcs_Vc_CellPackingData) GetParentYangName() string { return "vc" }

// AtmVcm_Nodes_Node_CellPacks
// Contains all cell packing information for node
type AtmVcm_Nodes_Node_CellPacks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All cell packing information on a node. The type is slice of
    // AtmVcm_Nodes_Node_CellPacks_CellPack.
    CellPack []AtmVcm_Nodes_Node_CellPacks_CellPack
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetFilter() yfilter.YFilter { return cellPacks.YFilter }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) SetFilter(yf yfilter.YFilter) { cellPacks.YFilter = yf }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetGoName(yname string) string {
    if yname == "cell-pack" { return "CellPack" }
    return ""
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetSegmentPath() string {
    return "cell-packs"
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cell-pack" {
        for _, c := range cellPacks.CellPack {
            if cellPacks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_CellPacks_CellPack{}
        cellPacks.CellPack = append(cellPacks.CellPack, child)
        return &cellPacks.CellPack[len(cellPacks.CellPack)-1]
    }
    return nil
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cellPacks.CellPack {
        children[cellPacks.CellPack[i].GetSegmentPath()] = &cellPacks.CellPack[i]
    }
    return children
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetBundleName() string { return "cisco_ios_xr" }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetYangName() string { return "cell-packs" }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) SetParent(parent types.Entity) { cellPacks.parent = parent }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetParent() types.Entity { return cellPacks.parent }

func (cellPacks *AtmVcm_Nodes_Node_CellPacks) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_CellPacks_CellPack
// All cell packing information on a node
type AtmVcm_Nodes_Node_CellPacks_CellPack struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // PCI. The type is interface{} with range: -2147483648..2147483647.
    Pci interface{}

    // Sub-interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetFilter() yfilter.YFilter { return cellPack.YFilter }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) SetFilter(yf yfilter.YFilter) { cellPack.YFilter = yf }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "pci" { return "Pci" }
    if yname == "sub-interface-name" { return "SubInterfaceName" }
    if yname == "cell-packing-mode" { return "CellPackingMode" }
    if yname == "vpi" { return "Vpi" }
    if yname == "vci" { return "Vci" }
    if yname == "received-average-cells-packets" { return "ReceivedAverageCellsPackets" }
    if yname == "sent-cells-packets" { return "SentCellsPackets" }
    if yname == "cell-packing" { return "CellPacking" }
    return ""
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetSegmentPath() string {
    return "cell-pack"
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cell-packing" {
        return &cellPack.CellPacking
    }
    return nil
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cell-packing"] = &cellPack.CellPacking
    return children
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = cellPack.InterfaceName
    leafs["pci"] = cellPack.Pci
    leafs["sub-interface-name"] = cellPack.SubInterfaceName
    leafs["cell-packing-mode"] = cellPack.CellPackingMode
    leafs["vpi"] = cellPack.Vpi
    leafs["vci"] = cellPack.Vci
    leafs["received-average-cells-packets"] = cellPack.ReceivedAverageCellsPackets
    leafs["sent-cells-packets"] = cellPack.SentCellsPackets
    return leafs
}

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetBundleName() string { return "cisco_ios_xr" }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetYangName() string { return "cell-pack" }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) SetParent(parent types.Entity) { cellPack.parent = parent }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetParent() types.Entity { return cellPack.parent }

func (cellPack *AtmVcm_Nodes_Node_CellPacks_CellPack) GetParentYangName() string { return "cell-packs" }

// AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking
// Cell packing specific data
type AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking struct {
    parent types.Entity
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

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetFilter() yfilter.YFilter { return cellPacking.YFilter }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) SetFilter(yf yfilter.YFilter) { cellPacking.YFilter = yf }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetGoName(yname string) string {
    if yname == "local-max-cells-packed-per-packet" { return "LocalMaxCellsPackedPerPacket" }
    if yname == "negotiated-max-cells-packed-per-packet" { return "NegotiatedMaxCellsPackedPerPacket" }
    if yname == "max-cell-packed-timeout" { return "MaxCellPackedTimeout" }
    return ""
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetSegmentPath() string {
    return "cell-packing"
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-max-cells-packed-per-packet"] = cellPacking.LocalMaxCellsPackedPerPacket
    leafs["negotiated-max-cells-packed-per-packet"] = cellPacking.NegotiatedMaxCellsPackedPerPacket
    leafs["max-cell-packed-timeout"] = cellPacking.MaxCellPackedTimeout
    return leafs
}

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetBundleName() string { return "cisco_ios_xr" }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetYangName() string { return "cell-packing" }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) SetParent(parent types.Entity) { cellPacking.parent = parent }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetParent() types.Entity { return cellPacking.parent }

func (cellPacking *AtmVcm_Nodes_Node_CellPacks_CellPack_CellPacking) GetParentYangName() string { return "cell-pack" }

// AtmVcm_Nodes_Node_Pvps
// Contains all L2 PVP information for node
type AtmVcm_Nodes_Node_Pvps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All L2 PVP information on a node. The type is slice of
    // AtmVcm_Nodes_Node_Pvps_Pvp.
    Pvp []AtmVcm_Nodes_Node_Pvps_Pvp
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetFilter() yfilter.YFilter { return pvps.YFilter }

func (pvps *AtmVcm_Nodes_Node_Pvps) SetFilter(yf yfilter.YFilter) { pvps.YFilter = yf }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetGoName(yname string) string {
    if yname == "pvp" { return "Pvp" }
    return ""
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetSegmentPath() string {
    return "pvps"
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pvp" {
        for _, c := range pvps.Pvp {
            if pvps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_Pvps_Pvp{}
        pvps.Pvp = append(pvps.Pvp, child)
        return &pvps.Pvp[len(pvps.Pvp)-1]
    }
    return nil
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pvps.Pvp {
        children[pvps.Pvp[i].GetSegmentPath()] = &pvps.Pvp[i]
    }
    return children
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pvps *AtmVcm_Nodes_Node_Pvps) GetBundleName() string { return "cisco_ios_xr" }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetYangName() string { return "pvps" }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pvps *AtmVcm_Nodes_Node_Pvps) SetParent(parent types.Entity) { pvps.parent = parent }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetParent() types.Entity { return pvps.parent }

func (pvps *AtmVcm_Nodes_Node_Pvps) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_Pvps_Pvp
// All L2 PVP information on a node
type AtmVcm_Nodes_Node_Pvps_Pvp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: -2147483648..2147483647.
    Vpi interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    MainInterface interface{}

    // Subinterface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    SubInterface interface{}

    // VC Interfcace handle. The type is string with pattern: [a-zA-Z0-9./-]+.
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
    VcOnP2PSubInterface interface{}

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

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetFilter() yfilter.YFilter { return pvp.YFilter }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) SetFilter(yf yfilter.YFilter) { pvp.YFilter = yf }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "vpi" { return "Vpi" }
    if yname == "main-interface" { return "MainInterface" }
    if yname == "sub-interface" { return "SubInterface" }
    if yname == "vc-interface" { return "VcInterface" }
    if yname == "vpi-xr" { return "VpiXr" }
    if yname == "vci-xr" { return "VciXr" }
    if yname == "type" { return "Type" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "shape" { return "Shape" }
    if yname == "peak-cell-rate" { return "PeakCellRate" }
    if yname == "sustained-cell-rate" { return "SustainedCellRate" }
    if yname == "burst-rate" { return "BurstRate" }
    if yname == "encaps-inherit-level" { return "EncapsInheritLevel" }
    if yname == "qos-inherit-level" { return "QosInheritLevel" }
    if yname == "transmit-mtu" { return "TransmitMtu" }
    if yname == "receive-mtu" { return "ReceiveMtu" }
    if yname == "vc-onvp-tunnel" { return "VcOnvpTunnel" }
    if yname == "vc-on-p2p-sub-interface" { return "VcOnP2PSubInterface" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "amin-status" { return "AminStatus" }
    if yname == "internal-state" { return "InternalState" }
    if yname == "last-state-change-time" { return "LastStateChangeTime" }
    if yname == "test-mode" { return "TestMode" }
    if yname == "cell-packing-data" { return "CellPackingData" }
    return ""
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetSegmentPath() string {
    return "pvp" + "[interface-name='" + fmt.Sprintf("%v", pvp.InterfaceName) + "']"
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cell-packing-data" {
        return &pvp.CellPackingData
    }
    return nil
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cell-packing-data"] = &pvp.CellPackingData
    return children
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = pvp.InterfaceName
    leafs["vpi"] = pvp.Vpi
    leafs["main-interface"] = pvp.MainInterface
    leafs["sub-interface"] = pvp.SubInterface
    leafs["vc-interface"] = pvp.VcInterface
    leafs["vpi-xr"] = pvp.VpiXr
    leafs["vci-xr"] = pvp.VciXr
    leafs["type"] = pvp.Type
    leafs["encapsulation"] = pvp.Encapsulation
    leafs["shape"] = pvp.Shape
    leafs["peak-cell-rate"] = pvp.PeakCellRate
    leafs["sustained-cell-rate"] = pvp.SustainedCellRate
    leafs["burst-rate"] = pvp.BurstRate
    leafs["encaps-inherit-level"] = pvp.EncapsInheritLevel
    leafs["qos-inherit-level"] = pvp.QosInheritLevel
    leafs["transmit-mtu"] = pvp.TransmitMtu
    leafs["receive-mtu"] = pvp.ReceiveMtu
    leafs["vc-onvp-tunnel"] = pvp.VcOnvpTunnel
    leafs["vc-on-p2p-sub-interface"] = pvp.VcOnP2PSubInterface
    leafs["oper-status"] = pvp.OperStatus
    leafs["amin-status"] = pvp.AminStatus
    leafs["internal-state"] = pvp.InternalState
    leafs["last-state-change-time"] = pvp.LastStateChangeTime
    leafs["test-mode"] = pvp.TestMode
    return leafs
}

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetBundleName() string { return "cisco_ios_xr" }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetYangName() string { return "pvp" }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) SetParent(parent types.Entity) { pvp.parent = parent }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetParent() types.Entity { return pvp.parent }

func (pvp *AtmVcm_Nodes_Node_Pvps_Pvp) GetParentYangName() string { return "pvps" }

// AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData
// Cell packing specific data
type AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData struct {
    parent types.Entity
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

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetFilter() yfilter.YFilter { return cellPackingData.YFilter }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) SetFilter(yf yfilter.YFilter) { cellPackingData.YFilter = yf }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetGoName(yname string) string {
    if yname == "local-max-cells-packed-per-packet" { return "LocalMaxCellsPackedPerPacket" }
    if yname == "negotiated-max-cells-packed-per-packet" { return "NegotiatedMaxCellsPackedPerPacket" }
    if yname == "max-cell-packed-timeout" { return "MaxCellPackedTimeout" }
    return ""
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetSegmentPath() string {
    return "cell-packing-data"
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-max-cells-packed-per-packet"] = cellPackingData.LocalMaxCellsPackedPerPacket
    leafs["negotiated-max-cells-packed-per-packet"] = cellPackingData.NegotiatedMaxCellsPackedPerPacket
    leafs["max-cell-packed-timeout"] = cellPackingData.MaxCellPackedTimeout
    return leafs
}

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetBundleName() string { return "cisco_ios_xr" }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetYangName() string { return "cell-packing-data" }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) SetParent(parent types.Entity) { cellPackingData.parent = parent }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetParent() types.Entity { return cellPackingData.parent }

func (cellPackingData *AtmVcm_Nodes_Node_Pvps_Pvp_CellPackingData) GetParentYangName() string { return "pvp" }

// AtmVcm_Nodes_Node_ClassLinks
// Contains all class link information for node
type AtmVcm_Nodes_Node_ClassLinks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All ATM VC information on a node. The type is slice of
    // AtmVcm_Nodes_Node_ClassLinks_ClassLink.
    ClassLink []AtmVcm_Nodes_Node_ClassLinks_ClassLink
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetFilter() yfilter.YFilter { return classLinks.YFilter }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) SetFilter(yf yfilter.YFilter) { classLinks.YFilter = yf }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetGoName(yname string) string {
    if yname == "class-link" { return "ClassLink" }
    return ""
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetSegmentPath() string {
    return "class-links"
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-link" {
        for _, c := range classLinks.ClassLink {
            if classLinks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_ClassLinks_ClassLink{}
        classLinks.ClassLink = append(classLinks.ClassLink, child)
        return &classLinks.ClassLink[len(classLinks.ClassLink)-1]
    }
    return nil
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classLinks.ClassLink {
        children[classLinks.ClassLink[i].GetSegmentPath()] = &classLinks.ClassLink[i]
    }
    return children
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetBundleName() string { return "cisco_ios_xr" }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetYangName() string { return "class-links" }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) SetParent(parent types.Entity) { classLinks.parent = parent }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetParent() types.Entity { return classLinks.parent }

func (classLinks *AtmVcm_Nodes_Node_ClassLinks) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink
// All ATM VC information on a node
type AtmVcm_Nodes_Node_ClassLinks_ClassLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VPI. The type is interface{} with range:
    // -2147483648..2147483647.
    Vpi interface{}

    // VCI. The type is interface{} with range: -2147483648..2147483647.
    Vci interface{}

    // Sub-interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    SubInterfaceName interface{}

    // Not supported VC class.
    VcClassNotSupported AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported

    // Oam values for class link.
    OamConfig AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetFilter() yfilter.YFilter { return classLink.YFilter }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) SetFilter(yf yfilter.YFilter) { classLink.YFilter = yf }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetGoName(yname string) string {
    if yname == "vpi" { return "Vpi" }
    if yname == "vci" { return "Vci" }
    if yname == "sub-interface-name" { return "SubInterfaceName" }
    if yname == "vc-class-not-supported" { return "VcClassNotSupported" }
    if yname == "oam-config" { return "OamConfig" }
    return ""
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetSegmentPath() string {
    return "class-link" + "[vpi='" + fmt.Sprintf("%v", classLink.Vpi) + "']"
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vc-class-not-supported" {
        return &classLink.VcClassNotSupported
    }
    if childYangName == "oam-config" {
        return &classLink.OamConfig
    }
    return nil
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vc-class-not-supported"] = &classLink.VcClassNotSupported
    children["oam-config"] = &classLink.OamConfig
    return children
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpi"] = classLink.Vpi
    leafs["vci"] = classLink.Vci
    leafs["sub-interface-name"] = classLink.SubInterfaceName
    return leafs
}

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetBundleName() string { return "cisco_ios_xr" }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetYangName() string { return "class-link" }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) SetParent(parent types.Entity) { classLink.parent = parent }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetParent() types.Entity { return classLink.parent }

func (classLink *AtmVcm_Nodes_Node_ClassLinks_ClassLink) GetParentYangName() string { return "class-links" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported
// Not supported VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation type not supported. The type is VcEncap.
    EncapsulationNotSupported interface{}

    // NotSupportedInheritLevel. The type is VcInheritLevel.
    NotSupportedInheritLevel interface{}
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetFilter() yfilter.YFilter { return vcClassNotSupported.YFilter }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) SetFilter(yf yfilter.YFilter) { vcClassNotSupported.YFilter = yf }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetGoName(yname string) string {
    if yname == "encapsulation-not-supported" { return "EncapsulationNotSupported" }
    if yname == "not-supported-inherit-level" { return "NotSupportedInheritLevel" }
    return ""
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetSegmentPath() string {
    return "vc-class-not-supported"
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation-not-supported"] = vcClassNotSupported.EncapsulationNotSupported
    leafs["not-supported-inherit-level"] = vcClassNotSupported.NotSupportedInheritLevel
    return leafs
}

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetBundleName() string { return "cisco_ios_xr" }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetYangName() string { return "vc-class-not-supported" }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) SetParent(parent types.Entity) { vcClassNotSupported.parent = parent }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetParent() types.Entity { return vcClassNotSupported.parent }

func (vcClassNotSupported *AtmVcm_Nodes_Node_ClassLinks_ClassLink_VcClassNotSupported) GetParentYangName() string { return "class-link" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig
// Oam values for class link
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig struct {
    parent types.Entity
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

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetFilter() yfilter.YFilter { return oamConfig.YFilter }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) SetFilter(yf yfilter.YFilter) { oamConfig.YFilter = yf }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetGoName(yname string) string {
    if yname == "class-link-shaping" { return "ClassLinkShaping" }
    if yname == "class-link-encapsulation" { return "ClassLinkEncapsulation" }
    if yname == "oam-pvc" { return "OamPvc" }
    if yname == "oam-retry" { return "OamRetry" }
    if yname == "ais-rdi" { return "AisRdi" }
    return ""
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetSegmentPath() string {
    return "oam-config"
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-link-shaping" {
        return &oamConfig.ClassLinkShaping
    }
    if childYangName == "class-link-encapsulation" {
        return &oamConfig.ClassLinkEncapsulation
    }
    if childYangName == "oam-pvc" {
        return &oamConfig.OamPvc
    }
    if childYangName == "oam-retry" {
        return &oamConfig.OamRetry
    }
    if childYangName == "ais-rdi" {
        return &oamConfig.AisRdi
    }
    return nil
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["class-link-shaping"] = &oamConfig.ClassLinkShaping
    children["class-link-encapsulation"] = &oamConfig.ClassLinkEncapsulation
    children["oam-pvc"] = &oamConfig.OamPvc
    children["oam-retry"] = &oamConfig.OamRetry
    children["ais-rdi"] = &oamConfig.AisRdi
    return children
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetBundleName() string { return "cisco_ios_xr" }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetYangName() string { return "oam-config" }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) SetParent(parent types.Entity) { oamConfig.parent = parent }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetParent() types.Entity { return oamConfig.parent }

func (oamConfig *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig) GetParentYangName() string { return "class-link" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping
// Traffic Shaping detail of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping struct {
    parent types.Entity
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

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetFilter() yfilter.YFilter { return classLinkShaping.YFilter }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) SetFilter(yf yfilter.YFilter) { classLinkShaping.YFilter = yf }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetGoName(yname string) string {
    if yname == "shaping-type" { return "ShapingType" }
    if yname == "peak-output-rate" { return "PeakOutputRate" }
    if yname == "average-output-rate" { return "AverageOutputRate" }
    if yname == "burst-output-rate" { return "BurstOutputRate" }
    if yname == "shaping-inherit-level" { return "ShapingInheritLevel" }
    return ""
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetSegmentPath() string {
    return "class-link-shaping"
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["shaping-type"] = classLinkShaping.ShapingType
    leafs["peak-output-rate"] = classLinkShaping.PeakOutputRate
    leafs["average-output-rate"] = classLinkShaping.AverageOutputRate
    leafs["burst-output-rate"] = classLinkShaping.BurstOutputRate
    leafs["shaping-inherit-level"] = classLinkShaping.ShapingInheritLevel
    return leafs
}

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetBundleName() string { return "cisco_ios_xr" }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetYangName() string { return "class-link-shaping" }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) SetParent(parent types.Entity) { classLinkShaping.parent = parent }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetParent() types.Entity { return classLinkShaping.parent }

func (classLinkShaping *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkShaping) GetParentYangName() string { return "oam-config" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation
// Encapsulation details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Encapsulation type. The type is VcEncap.
    EncapsulationType interface{}

    // Encapsulation inherit level. The type is VcInheritLevel.
    EncapsulationInheritLevel interface{}
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetFilter() yfilter.YFilter { return classLinkEncapsulation.YFilter }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) SetFilter(yf yfilter.YFilter) { classLinkEncapsulation.YFilter = yf }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetGoName(yname string) string {
    if yname == "encapsulation-type" { return "EncapsulationType" }
    if yname == "encapsulation-inherit-level" { return "EncapsulationInheritLevel" }
    return ""
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetSegmentPath() string {
    return "class-link-encapsulation"
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["encapsulation-type"] = classLinkEncapsulation.EncapsulationType
    leafs["encapsulation-inherit-level"] = classLinkEncapsulation.EncapsulationInheritLevel
    return leafs
}

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetBundleName() string { return "cisco_ios_xr" }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetYangName() string { return "class-link-encapsulation" }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) SetParent(parent types.Entity) { classLinkEncapsulation.parent = parent }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetParent() types.Entity { return classLinkEncapsulation.parent }

func (classLinkEncapsulation *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_ClassLinkEncapsulation) GetParentYangName() string { return "oam-config" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc
// OAM PVC details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc struct {
    parent types.Entity
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

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetFilter() yfilter.YFilter { return oamPvc.YFilter }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) SetFilter(yf yfilter.YFilter) { oamPvc.YFilter = yf }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetGoName(yname string) string {
    if yname == "manage-level" { return "ManageLevel" }
    if yname == "pvc-frequency" { return "PvcFrequency" }
    if yname == "keep-vc-up" { return "KeepVcUp" }
    if yname == "ais-rdi-failure" { return "AisRdiFailure" }
    if yname == "manage-inherit-level" { return "ManageInheritLevel" }
    return ""
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetSegmentPath() string {
    return "oam-pvc"
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["manage-level"] = oamPvc.ManageLevel
    leafs["pvc-frequency"] = oamPvc.PvcFrequency
    leafs["keep-vc-up"] = oamPvc.KeepVcUp
    leafs["ais-rdi-failure"] = oamPvc.AisRdiFailure
    leafs["manage-inherit-level"] = oamPvc.ManageInheritLevel
    return leafs
}

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetBundleName() string { return "cisco_ios_xr" }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetYangName() string { return "oam-pvc" }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) SetParent(parent types.Entity) { oamPvc.parent = parent }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetParent() types.Entity { return oamPvc.parent }

func (oamPvc *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamPvc) GetParentYangName() string { return "oam-config" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry
// OAM Retry details of VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry struct {
    parent types.Entity
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

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetFilter() yfilter.YFilter { return oamRetry.YFilter }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) SetFilter(yf yfilter.YFilter) { oamRetry.YFilter = yf }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetGoName(yname string) string {
    if yname == "retry-up-count" { return "RetryUpCount" }
    if yname == "down-count" { return "DownCount" }
    if yname == "retry-frequency" { return "RetryFrequency" }
    if yname == "retry-inherit-level" { return "RetryInheritLevel" }
    return ""
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetSegmentPath() string {
    return "oam-retry"
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retry-up-count"] = oamRetry.RetryUpCount
    leafs["down-count"] = oamRetry.DownCount
    leafs["retry-frequency"] = oamRetry.RetryFrequency
    leafs["retry-inherit-level"] = oamRetry.RetryInheritLevel
    return leafs
}

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetBundleName() string { return "cisco_ios_xr" }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetYangName() string { return "oam-retry" }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) SetParent(parent types.Entity) { oamRetry.parent = parent }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetParent() types.Entity { return oamRetry.parent }

func (oamRetry *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_OamRetry) GetParentYangName() string { return "oam-config" }

// AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi
// AIS RDI details of a VC class
type AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS RDI up count. The type is interface{} with range: 0..4294967295.
    AisRdiUpCount interface{}

    // Time (in seconds) with no AIS/RDI cells before declaring a VC as up. The
    // type is interface{} with range: 0..4294967295. Units are second.
    AisRdiUpTime interface{}

    // AIS RDI inherit level. The type is ClassLinkOamInheritLevel.
    AisRdiInheritLevel interface{}
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetFilter() yfilter.YFilter { return aisRdi.YFilter }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) SetFilter(yf yfilter.YFilter) { aisRdi.YFilter = yf }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetGoName(yname string) string {
    if yname == "ais-rdi-up-count" { return "AisRdiUpCount" }
    if yname == "ais-rdi-up-time" { return "AisRdiUpTime" }
    if yname == "ais-rdi-inherit-level" { return "AisRdiInheritLevel" }
    return ""
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetSegmentPath() string {
    return "ais-rdi"
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ais-rdi-up-count"] = aisRdi.AisRdiUpCount
    leafs["ais-rdi-up-time"] = aisRdi.AisRdiUpTime
    leafs["ais-rdi-inherit-level"] = aisRdi.AisRdiInheritLevel
    return leafs
}

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetBundleName() string { return "cisco_ios_xr" }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetYangName() string { return "ais-rdi" }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) SetParent(parent types.Entity) { aisRdi.parent = parent }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetParent() types.Entity { return aisRdi.parent }

func (aisRdi *AtmVcm_Nodes_Node_ClassLinks_ClassLink_OamConfig_AisRdi) GetParentYangName() string { return "oam-config" }

// AtmVcm_Nodes_Node_Interfaces
// Contains all Interface information for node
type AtmVcm_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ATM Interface data. The type is slice of
    // AtmVcm_Nodes_Node_Interfaces_Interface.
    Interface []AtmVcm_Nodes_Node_Interfaces_Interface
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *AtmVcm_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_Interfaces_Interface
// ATM Interface data
type AtmVcm_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // ILMI VPI. The type is interface{} with range: 0..4294967295.
    IlmiVpi interface{}

    // ILMI VCI. The type is interface{} with range: 0..4294967295.
    IlmiVci interface{}

    // Number of PVC Failures. The type is interface{} with range: 0..4294967295.
    PvcFailures interface{}

    // Number of currently failing Layer 2 PVPs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer2PvPs interface{}

    // Number of currently failing Layer 2 PVCs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer2PvCs interface{}

    // Number of currently failing Layer 3 VP tunnels. The type is interface{}
    // with range: 0..4294967295.
    CurrentlyFailingLayer3VpTunnels interface{}

    // Number of currently failing Layer 3 PVCs. The type is interface{} with
    // range: 0..4294967295.
    CurrentlyFailingLayer3PvCs interface{}

    // If true, PVC failures trap is enabled. The type is bool.
    PvcFailuresTrapEnable interface{}

    // PVC trap notification interval. The type is interface{} with range:
    // 0..4294967295.
    PvcNotificationInterval interface{}

    // Number of Layer 2 PVPs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer2PvPs interface{}

    // Number of Layer 2 PVCs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer2PvCs interface{}

    // Number of Layer 3 VP tunnels configured. The type is interface{} with
    // range: 0..4294967295.
    ConfiguredLayer3VpTunnels interface{}

    // Number of Layer 3 PVCs configured. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredLayer3PvCs interface{}

    // ATM interface port type. The type is VcmPort.
    PortType interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    MainInterface interface{}

    // Number of L2 attachment circuits with the cell packing feature enabled on
    // this main interface. The type is interface{} with range: 0..65535.
    L2CellPackingCount interface{}

    // Cell packing specific information.
    CellPackingData AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ilmi-vpi" { return "IlmiVpi" }
    if yname == "ilmi-vci" { return "IlmiVci" }
    if yname == "pvc-failures" { return "PvcFailures" }
    if yname == "currently-failing-layer2pv-ps" { return "CurrentlyFailingLayer2PvPs" }
    if yname == "currently-failing-layer2pv-cs" { return "CurrentlyFailingLayer2PvCs" }
    if yname == "currently-failing-layer3vp-tunnels" { return "CurrentlyFailingLayer3VpTunnels" }
    if yname == "currently-failing-layer3pv-cs" { return "CurrentlyFailingLayer3PvCs" }
    if yname == "pvc-failures-trap-enable" { return "PvcFailuresTrapEnable" }
    if yname == "pvc-notification-interval" { return "PvcNotificationInterval" }
    if yname == "configured-layer2pv-ps" { return "ConfiguredLayer2PvPs" }
    if yname == "configured-layer2pv-cs" { return "ConfiguredLayer2PvCs" }
    if yname == "configured-layer3vp-tunnels" { return "ConfiguredLayer3VpTunnels" }
    if yname == "configured-layer3pv-cs" { return "ConfiguredLayer3PvCs" }
    if yname == "port-type" { return "PortType" }
    if yname == "main-interface" { return "MainInterface" }
    if yname == "l2-cell-packing-count" { return "L2CellPackingCount" }
    if yname == "cell-packing-data" { return "CellPackingData" }
    return ""
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cell-packing-data" {
        return &self.CellPackingData
    }
    return nil
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cell-packing-data"] = &self.CellPackingData
    return children
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["ilmi-vpi"] = self.IlmiVpi
    leafs["ilmi-vci"] = self.IlmiVci
    leafs["pvc-failures"] = self.PvcFailures
    leafs["currently-failing-layer2pv-ps"] = self.CurrentlyFailingLayer2PvPs
    leafs["currently-failing-layer2pv-cs"] = self.CurrentlyFailingLayer2PvCs
    leafs["currently-failing-layer3vp-tunnels"] = self.CurrentlyFailingLayer3VpTunnels
    leafs["currently-failing-layer3pv-cs"] = self.CurrentlyFailingLayer3PvCs
    leafs["pvc-failures-trap-enable"] = self.PvcFailuresTrapEnable
    leafs["pvc-notification-interval"] = self.PvcNotificationInterval
    leafs["configured-layer2pv-ps"] = self.ConfiguredLayer2PvPs
    leafs["configured-layer2pv-cs"] = self.ConfiguredLayer2PvCs
    leafs["configured-layer3vp-tunnels"] = self.ConfiguredLayer3VpTunnels
    leafs["configured-layer3pv-cs"] = self.ConfiguredLayer3PvCs
    leafs["port-type"] = self.PortType
    leafs["main-interface"] = self.MainInterface
    leafs["l2-cell-packing-count"] = self.L2CellPackingCount
    return leafs
}

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *AtmVcm_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData
// Cell packing specific information
type AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData struct {
    parent types.Entity
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

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetFilter() yfilter.YFilter { return cellPackingData.YFilter }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) SetFilter(yf yfilter.YFilter) { cellPackingData.YFilter = yf }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetGoName(yname string) string {
    if yname == "local-max-cells-packed-per-packet" { return "LocalMaxCellsPackedPerPacket" }
    if yname == "negotiated-max-cells-packed-per-packet" { return "NegotiatedMaxCellsPackedPerPacket" }
    if yname == "max-cell-packed-timeout" { return "MaxCellPackedTimeout" }
    return ""
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetSegmentPath() string {
    return "cell-packing-data"
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-max-cells-packed-per-packet"] = cellPackingData.LocalMaxCellsPackedPerPacket
    leafs["negotiated-max-cells-packed-per-packet"] = cellPackingData.NegotiatedMaxCellsPackedPerPacket
    leafs["max-cell-packed-timeout"] = cellPackingData.MaxCellPackedTimeout
    return leafs
}

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetBundleName() string { return "cisco_ios_xr" }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetYangName() string { return "cell-packing-data" }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) SetParent(parent types.Entity) { cellPackingData.parent = parent }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetParent() types.Entity { return cellPackingData.parent }

func (cellPackingData *AtmVcm_Nodes_Node_Interfaces_Interface_CellPackingData) GetParentYangName() string { return "interface" }

// AtmVcm_Nodes_Node_VpTunnels
// Contains all VP-tunnel information for node
type AtmVcm_Nodes_Node_VpTunnels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All VP-tunnel information on a node. The type is slice of
    // AtmVcm_Nodes_Node_VpTunnels_VpTunnel.
    VpTunnel []AtmVcm_Nodes_Node_VpTunnels_VpTunnel
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetFilter() yfilter.YFilter { return vpTunnels.YFilter }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) SetFilter(yf yfilter.YFilter) { vpTunnels.YFilter = yf }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetGoName(yname string) string {
    if yname == "vp-tunnel" { return "VpTunnel" }
    return ""
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetSegmentPath() string {
    return "vp-tunnels"
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vp-tunnel" {
        for _, c := range vpTunnels.VpTunnel {
            if vpTunnels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AtmVcm_Nodes_Node_VpTunnels_VpTunnel{}
        vpTunnels.VpTunnel = append(vpTunnels.VpTunnel, child)
        return &vpTunnels.VpTunnel[len(vpTunnels.VpTunnel)-1]
    }
    return nil
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vpTunnels.VpTunnel {
        children[vpTunnels.VpTunnel[i].GetSegmentPath()] = &vpTunnels.VpTunnel[i]
    }
    return children
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetBundleName() string { return "cisco_ios_xr" }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetYangName() string { return "vp-tunnels" }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) SetParent(parent types.Entity) { vpTunnels.parent = parent }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetParent() types.Entity { return vpTunnels.parent }

func (vpTunnels *AtmVcm_Nodes_Node_VpTunnels) GetParentYangName() string { return "node" }

// AtmVcm_Nodes_Node_VpTunnels_VpTunnel
// All VP-tunnel information on a node
type AtmVcm_Nodes_Node_VpTunnels_VpTunnel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // VPI. The type is interface{} with range: -2147483648..2147483647.
    Vpi interface{}

    // Main Interface handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    MainInterface interface{}

    // VP Interfcace handle. The type is string with pattern: [a-zA-Z0-9./-]+.
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
    F4OamEnabled interface{}

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

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetFilter() yfilter.YFilter { return vpTunnel.YFilter }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) SetFilter(yf yfilter.YFilter) { vpTunnel.YFilter = yf }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "vpi" { return "Vpi" }
    if yname == "main-interface" { return "MainInterface" }
    if yname == "vp-interface" { return "VpInterface" }
    if yname == "vpi-xr" { return "VpiXr" }
    if yname == "shape" { return "Shape" }
    if yname == "peak-cell-rate" { return "PeakCellRate" }
    if yname == "sustained-cell-rate" { return "SustainedCellRate" }
    if yname == "burst-rate" { return "BurstRate" }
    if yname == "f4oam-enabled" { return "F4OamEnabled" }
    if yname == "data-vc-count" { return "DataVcCount" }
    if yname == "oper-status" { return "OperStatus" }
    if yname == "amin-status" { return "AminStatus" }
    if yname == "internal-state" { return "InternalState" }
    if yname == "last-vp-state-change-time" { return "LastVpStateChangeTime" }
    return ""
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetSegmentPath() string {
    return "vp-tunnel" + "[interface-name='" + fmt.Sprintf("%v", vpTunnel.InterfaceName) + "']"
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = vpTunnel.InterfaceName
    leafs["vpi"] = vpTunnel.Vpi
    leafs["main-interface"] = vpTunnel.MainInterface
    leafs["vp-interface"] = vpTunnel.VpInterface
    leafs["vpi-xr"] = vpTunnel.VpiXr
    leafs["shape"] = vpTunnel.Shape
    leafs["peak-cell-rate"] = vpTunnel.PeakCellRate
    leafs["sustained-cell-rate"] = vpTunnel.SustainedCellRate
    leafs["burst-rate"] = vpTunnel.BurstRate
    leafs["f4oam-enabled"] = vpTunnel.F4OamEnabled
    leafs["data-vc-count"] = vpTunnel.DataVcCount
    leafs["oper-status"] = vpTunnel.OperStatus
    leafs["amin-status"] = vpTunnel.AminStatus
    leafs["internal-state"] = vpTunnel.InternalState
    leafs["last-vp-state-change-time"] = vpTunnel.LastVpStateChangeTime
    return leafs
}

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetBundleName() string { return "cisco_ios_xr" }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetYangName() string { return "vp-tunnel" }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) SetParent(parent types.Entity) { vpTunnel.parent = parent }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetParent() types.Entity { return vpTunnel.parent }

func (vpTunnel *AtmVcm_Nodes_Node_VpTunnels_VpTunnel) GetParentYangName() string { return "vp-tunnels" }

