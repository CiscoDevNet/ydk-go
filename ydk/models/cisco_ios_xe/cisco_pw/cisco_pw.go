// This YANG module describes the configuration and operational
// model of Psuedowire on PE router
// 
// Terms and Acronyms
//    AC    : Attachment Circuit
//    BD    : Bridge Domain
//    L2VPN : Layer 2 Virtual Private Network
//    PE    : Provider Edge
//    PW    : Pseudowire
//    VFI   : Virtual Forwarding Instance
//    VPLS  : Virtual Private LAN Service
//    VLAN  : Virtual Local Area Network
// 
package cisco_pw

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_pw"))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:pw pseudowire-config}", reflect.TypeOf(PseudowireConfig{}))
    ydk.RegisterEntity("cisco-pw:pseudowire-config", reflect.TypeOf(PseudowireConfig{}))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:pw pseudowire-state}", reflect.TypeOf(PseudowireState{}))
    ydk.RegisterEntity("cisco-pw:pseudowire-state", reflect.TypeOf(PseudowireState{}))
}

type PwEncapsulationType struct {
}

func (id PwEncapsulationType) String() string {
	return "cisco-pw:pw-encapsulation-type"
}

type PwEncapMpls struct {
}

func (id PwEncapMpls) String() string {
	return "cisco-pw:pw-encap-mpls"
}

type PwVcType struct {
}

func (id PwVcType) String() string {
	return "cisco-pw:pw-vc-type"
}

type PwVcTypeEther struct {
}

func (id PwVcTypeEther) String() string {
	return "cisco-pw:pw-vc-type-ether"
}

type PwVcTypeVlan struct {
}

func (id PwVcTypeVlan) String() string {
	return "cisco-pw:pw-vc-type-vlan"
}

type PwVcTypeVlanPassthrough struct {
}

func (id PwVcTypeVlanPassthrough) String() string {
	return "cisco-pw:pw-vc-type-vlan-passthrough"
}

type PwLoadBalanceType struct {
}

func (id PwLoadBalanceType) String() string {
	return "cisco-pw:pw-load-balance-type"
}

type PwLbEthernetType struct {
}

func (id PwLbEthernetType) String() string {
	return "cisco-pw:pw-lb-ethernet-type"
}

type PwLbEthSrcMac struct {
}

func (id PwLbEthSrcMac) String() string {
	return "cisco-pw:pw-lb-eth-src-mac"
}

type PwLbEthDstMac struct {
}

func (id PwLbEthDstMac) String() string {
	return "cisco-pw:pw-lb-eth-dst-mac"
}

type PwLbEthSrcDstMac struct {
}

func (id PwLbEthSrcDstMac) String() string {
	return "cisco-pw:pw-lb-eth-src-dst-mac"
}

type PwLbIpType struct {
}

func (id PwLbIpType) String() string {
	return "cisco-pw:pw-lb-ip-type"
}

type PwLbIpSrcIp struct {
}

func (id PwLbIpSrcIp) String() string {
	return "cisco-pw:pw-lb-ip-src-ip"
}

type PwLbIpDstIp struct {
}

func (id PwLbIpDstIp) String() string {
	return "cisco-pw:pw-lb-ip-dst-ip"
}

type PwLbIpSrcDstIp struct {
}

func (id PwLbIpSrcDstIp) String() string {
	return "cisco-pw:pw-lb-ip-src-dst-ip"
}

type PwSignalingProtocolType struct {
}

func (id PwSignalingProtocolType) String() string {
	return "cisco-pw:pw-signaling-protocol-type"
}

type PwSignalingProtocolNone struct {
}

func (id PwSignalingProtocolNone) String() string {
	return "cisco-pw:pw-signaling-protocol-none"
}

type PwSignalingProtocolLdp struct {
}

func (id PwSignalingProtocolLdp) String() string {
	return "cisco-pw:pw-signaling-protocol-ldp"
}

type PwSignalingProtocolBgp struct {
}

func (id PwSignalingProtocolBgp) String() string {
	return "cisco-pw:pw-signaling-protocol-bgp"
}

type PwSequencingType struct {
}

func (id PwSequencingType) String() string {
	return "cisco-pw:pw-sequencing-type"
}

type PwSequencingReceive struct {
}

func (id PwSequencingReceive) String() string {
	return "cisco-pw:pw-sequencing-receive"
}

type PwSequencingTransmit struct {
}

func (id PwSequencingTransmit) String() string {
	return "cisco-pw:pw-sequencing-transmit"
}

type PwSequencingBoth struct {
}

func (id PwSequencingBoth) String() string {
	return "cisco-pw:pw-sequencing-both"
}

// PwOperStateType represents Indicates the operational status of the PW VC
type PwOperStateType string

const (
    // Pseudowire is established, i.e., it is operational
    // and available to carry packets
    PwOperStateType_up PwOperStateType = "up"

    // Pseudowire is idle, i.e., it is administratively up
    // but network connectivity with the neighbor is not
    // established
    PwOperStateType_down PwOperStateType = "down"

    // Pseudowire can take over traffic for another
    // pseudowire
    PwOperStateType_cold_standby PwOperStateType = "cold-standby"

    // Pseudowire is available to immediately take over
    // traffic from another pseudowire
    PwOperStateType_hot_standby PwOperStateType = "hot-standby"

    // Pseudowire was previously in active state and is
    // in the process of being restored to active state
    PwOperStateType_recovering PwOperStateType = "recovering"

    // Pseudowire no longer has the required hardware
    // resources to be able to offer packet service to
    // the system
    PwOperStateType_no_hardware PwOperStateType = "no-hardware"

    // Pseudowire is not completely provisioned
    PwOperStateType_unresolved PwOperStateType = "unresolved"

    // Pseudowire has received local configuration
    PwOperStateType_provisioned PwOperStateType = "provisioned"

    // Pseudowire neighbor has sent its pseudowire
    // binding information
    PwOperStateType_remote_standby PwOperStateType = "remote-standby"

    // Pseudowire is provisioned, got a local label, and
    // its local binding has been sent to the neighbor
    PwOperStateType_local_ready PwOperStateType = "local-ready"

    // Pseudowire is locally provisioned, local binding
    // has been sent, and remote binding was received
    // from the neighbor
    PwOperStateType_all_ready PwOperStateType = "all-ready"
)

// PseudowireConfig
// Pseudowire configuration data.
type PseudowireConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global configurations related to pseudowires.
    Global PseudowireConfig_Global

    // Contains list of all pw-template definitions. Also called PW Class (XR) and
    // Port Profile (NX-OS).
    PwTemplates PseudowireConfig_PwTemplates

    // List of pseudowire static oam classes.
    PwStaticOamClasses PseudowireConfig_PwStaticOamClasses
}

func (pseudowireConfig *PseudowireConfig) GetEntityData() *types.CommonEntityData {
    pseudowireConfig.EntityData.YFilter = pseudowireConfig.YFilter
    pseudowireConfig.EntityData.YangName = "pseudowire-config"
    pseudowireConfig.EntityData.BundleName = "cisco_ios_xe"
    pseudowireConfig.EntityData.ParentYangName = "cisco-pw"
    pseudowireConfig.EntityData.SegmentPath = "cisco-pw:pseudowire-config"
    pseudowireConfig.EntityData.AbsolutePath = pseudowireConfig.EntityData.SegmentPath
    pseudowireConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pseudowireConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pseudowireConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pseudowireConfig.EntityData.Children = types.NewOrderedMap()
    pseudowireConfig.EntityData.Children.Append("global", types.YChild{"Global", &pseudowireConfig.Global})
    pseudowireConfig.EntityData.Children.Append("pw-templates", types.YChild{"PwTemplates", &pseudowireConfig.PwTemplates})
    pseudowireConfig.EntityData.Children.Append("pw-static-oam-classes", types.YChild{"PwStaticOamClasses", &pseudowireConfig.PwStaticOamClasses})
    pseudowireConfig.EntityData.Leafs = types.NewOrderedMap()

    pseudowireConfig.EntityData.YListKeys = []string {}

    return &(pseudowireConfig.EntityData)
}

// PseudowireConfig_Global
// Global configurations related to pseudowires.
type PseudowireConfig_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable pw-grouping.  When pseudowires (PW) are established, each PW is
    // assigned a group ID that is common for all PWs created from the same
    // physical port. Hence, when the physical port becomes non-functional or is
    // deleted, L2VPN sends a single message to advertise the status change of all
    // PWs belonging to the group. A single L2VPN signal thus avoids a lot of
    // processing and loss in reactivity. The type is bool. The default value is
    // false.
    PwGrouping interface{}

    // Set pseudowire oam transmit refresh time (in seconds). The type is
    // interface{} with range: 1..4095.
    PwOamRefreshTransmit interface{}

    // Enable pw-status. The type is bool.
    PwStatus interface{}

    // Enable predictive redundancy. The type is bool.
    PredictiveRedundancy interface{}

    // If this leaf is set to true, then it enables the emission of
    // 'vc-state-notification'; otherwise these notifications are not emitted. The
    // type is bool.
    VcStateNotificationEnabled interface{}

    // 'vc-state-notification' allows batching of pseudowires in order to reduce
    // number of notifications generated from the device. All pseudowires in a
    // batched notification MUST share same state at the same time.  This leaf
    // defines the maximum number of VCs that can be batched in
    // vc-state-notification. The type is interface{} with range: 0..4294967295.
    VcStateNotificationBatchSize interface{}

    // This leaf defines the maximum number of VC state change notifications that
    // can be emitted from the device per second. The type is interface{} with
    // range: 0..4294967295.
    VcStateNotificationRate interface{}
}

func (global *PseudowireConfig_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xe"
    global.EntityData.ParentYangName = "pseudowire-config"
    global.EntityData.SegmentPath = "global"
    global.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("pw-grouping", types.YLeaf{"PwGrouping", global.PwGrouping})
    global.EntityData.Leafs.Append("pw-oam-refresh-transmit", types.YLeaf{"PwOamRefreshTransmit", global.PwOamRefreshTransmit})
    global.EntityData.Leafs.Append("pw-status", types.YLeaf{"PwStatus", global.PwStatus})
    global.EntityData.Leafs.Append("predictive-redundancy", types.YLeaf{"PredictiveRedundancy", global.PredictiveRedundancy})
    global.EntityData.Leafs.Append("vc-state-notification-enabled", types.YLeaf{"VcStateNotificationEnabled", global.VcStateNotificationEnabled})
    global.EntityData.Leafs.Append("vc-state-notification-batch-size", types.YLeaf{"VcStateNotificationBatchSize", global.VcStateNotificationBatchSize})
    global.EntityData.Leafs.Append("vc-state-notification-rate", types.YLeaf{"VcStateNotificationRate", global.VcStateNotificationRate})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// PseudowireConfig_PwTemplates
// Contains list of all pw-template definitions.
// Also called PW Class (XR) and Port Profile (NX-OS)
type PseudowireConfig_PwTemplates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire template list. The type is slice of
    // PseudowireConfig_PwTemplates_PwTemplate.
    PwTemplate []*PseudowireConfig_PwTemplates_PwTemplate
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetEntityData() *types.CommonEntityData {
    pwTemplates.EntityData.YFilter = pwTemplates.YFilter
    pwTemplates.EntityData.YangName = "pw-templates"
    pwTemplates.EntityData.BundleName = "cisco_ios_xe"
    pwTemplates.EntityData.ParentYangName = "pseudowire-config"
    pwTemplates.EntityData.SegmentPath = "pw-templates"
    pwTemplates.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/" + pwTemplates.EntityData.SegmentPath
    pwTemplates.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwTemplates.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwTemplates.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwTemplates.EntityData.Children = types.NewOrderedMap()
    pwTemplates.EntityData.Children.Append("pw-template", types.YChild{"PwTemplate", nil})
    for i := range pwTemplates.PwTemplate {
        pwTemplates.EntityData.Children.Append(types.GetSegmentPath(pwTemplates.PwTemplate[i]), types.YChild{"PwTemplate", pwTemplates.PwTemplate[i]})
    }
    pwTemplates.EntityData.Leafs = types.NewOrderedMap()

    pwTemplates.EntityData.YListKeys = []string {}

    return &(pwTemplates.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate
// Pseudowire template list.
type PseudowireConfig_PwTemplates_PwTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. PW Template name. The type is string.
    Name interface{}

    // Encapsulation used for PW. The type is one of the following: PwEncapMpls.
    Encapsulation interface{}

    // Use control word in the MPLS PW header. The type is bool.
    ControlWord interface{}

    // Signaling protocol to use. The type is one of the following:
    // PwSignalingProtocolNonePwSignalingProtocolLdpPwSignalingProtocolBgp.
    SignalingProtocol interface{}

    // Type of VC in the PW. The type is one of the following:
    // PwVcTypeEtherPwVcTypeVlanPwVcTypeVlanPassthrough.
    VcType interface{}

    // Send switching TLV. The type is bool.
    SwitchingTlv interface{}

    // The local source IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceIp interface{}

    // Configure ingress tag rewrite vlan. The type is interface{} with range:
    // 1..4094.
    TagRewriteIngressVlan interface{}

    // Send Mac-withdraw message when PW becomes active. The type is bool. The
    // default value is false.
    MacWithdraw interface{}

    // Load balancing mechanism.
    LoadBalance PseudowireConfig_PwTemplates_PwTemplate_LoadBalance

    // Preferred path for the PW.
    PreferredPath PseudowireConfig_PwTemplates_PwTemplate_PreferredPath

    // Sequencing options.
    Sequencing PseudowireConfig_PwTemplates_PwTemplate_Sequencing

    // VCCV configuration.
    Vccv PseudowireConfig_PwTemplates_PwTemplate_Vccv

    // Timer configuration related to pseudowire redundancy switchover and
    // restoring to primary.
    SwitchoverDelay PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay

    // TODO.
    Status PseudowireConfig_PwTemplates_PwTemplate_Status

    // Pseudowire port profile configurations.
    PortProfileSpec PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetEntityData() *types.CommonEntityData {
    pwTemplate.EntityData.YFilter = pwTemplate.YFilter
    pwTemplate.EntityData.YangName = "pw-template"
    pwTemplate.EntityData.BundleName = "cisco_ios_xe"
    pwTemplate.EntityData.ParentYangName = "pw-templates"
    pwTemplate.EntityData.SegmentPath = "pw-template" + types.AddKeyToken(pwTemplate.Name, "name")
    pwTemplate.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/" + pwTemplate.EntityData.SegmentPath
    pwTemplate.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwTemplate.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwTemplate.EntityData.Children = types.NewOrderedMap()
    pwTemplate.EntityData.Children.Append("load-balance", types.YChild{"LoadBalance", &pwTemplate.LoadBalance})
    pwTemplate.EntityData.Children.Append("preferred-path", types.YChild{"PreferredPath", &pwTemplate.PreferredPath})
    pwTemplate.EntityData.Children.Append("sequencing", types.YChild{"Sequencing", &pwTemplate.Sequencing})
    pwTemplate.EntityData.Children.Append("vccv", types.YChild{"Vccv", &pwTemplate.Vccv})
    pwTemplate.EntityData.Children.Append("switchover-delay", types.YChild{"SwitchoverDelay", &pwTemplate.SwitchoverDelay})
    pwTemplate.EntityData.Children.Append("status", types.YChild{"Status", &pwTemplate.Status})
    pwTemplate.EntityData.Children.Append("port-profile-spec", types.YChild{"PortProfileSpec", &pwTemplate.PortProfileSpec})
    pwTemplate.EntityData.Leafs = types.NewOrderedMap()
    pwTemplate.EntityData.Leafs.Append("name", types.YLeaf{"Name", pwTemplate.Name})
    pwTemplate.EntityData.Leafs.Append("encapsulation", types.YLeaf{"Encapsulation", pwTemplate.Encapsulation})
    pwTemplate.EntityData.Leafs.Append("control-word", types.YLeaf{"ControlWord", pwTemplate.ControlWord})
    pwTemplate.EntityData.Leafs.Append("signaling-protocol", types.YLeaf{"SignalingProtocol", pwTemplate.SignalingProtocol})
    pwTemplate.EntityData.Leafs.Append("vc-type", types.YLeaf{"VcType", pwTemplate.VcType})
    pwTemplate.EntityData.Leafs.Append("switching-tlv", types.YLeaf{"SwitchingTlv", pwTemplate.SwitchingTlv})
    pwTemplate.EntityData.Leafs.Append("source-ip", types.YLeaf{"SourceIp", pwTemplate.SourceIp})
    pwTemplate.EntityData.Leafs.Append("tag-rewrite-ingress-vlan", types.YLeaf{"TagRewriteIngressVlan", pwTemplate.TagRewriteIngressVlan})
    pwTemplate.EntityData.Leafs.Append("mac-withdraw", types.YLeaf{"MacWithdraw", pwTemplate.MacWithdraw})

    pwTemplate.EntityData.YListKeys = []string {"Name"}

    return &(pwTemplate.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_LoadBalance
// Load balancing mechanism.
type PseudowireConfig_PwTemplates_PwTemplate_LoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ethernet mac address based load balancing. The type is one of the
    // following: PwLbEthSrcMacPwLbEthDstMacPwLbEthSrcDstMac. The default value is
    // pw-lb-eth-src-dst-mac.
    Ethernet interface{}

    // IP address based load balancing. The type is one of the following:
    // PwLbIpSrcIpPwLbIpSrcDstIp.
    Ip interface{}

    // Enable Flow Aware Label (FAT) PW - the capability to carry flow label on
    // PW.
    FlowLabel PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetEntityData() *types.CommonEntityData {
    loadBalance.EntityData.YFilter = loadBalance.YFilter
    loadBalance.EntityData.YangName = "load-balance"
    loadBalance.EntityData.BundleName = "cisco_ios_xe"
    loadBalance.EntityData.ParentYangName = "pw-template"
    loadBalance.EntityData.SegmentPath = "load-balance"
    loadBalance.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + loadBalance.EntityData.SegmentPath
    loadBalance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    loadBalance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    loadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    loadBalance.EntityData.Children = types.NewOrderedMap()
    loadBalance.EntityData.Children.Append("flow-label", types.YChild{"FlowLabel", &loadBalance.FlowLabel})
    loadBalance.EntityData.Leafs = types.NewOrderedMap()
    loadBalance.EntityData.Leafs.Append("ethernet", types.YLeaf{"Ethernet", loadBalance.Ethernet})
    loadBalance.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", loadBalance.Ip})

    loadBalance.EntityData.YListKeys = []string {}

    return &(loadBalance.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel
// Enable Flow Aware Label (FAT) PW - the capability to
// carry flow label on PW
type PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Directions to enable Flow Aware Label PW. The type is Direction.
    Direction interface{}

    // Carry code 0x17 as Flow Aware Label (FAT) PW signalled sub-tlv id. The type
    // is bool.
    TlvCode17 interface{}

    // Use statically configured flow label on traffic flowing on the PW. The type
    // is bool. The default value is false.
    Static interface{}
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetEntityData() *types.CommonEntityData {
    flowLabel.EntityData.YFilter = flowLabel.YFilter
    flowLabel.EntityData.YangName = "flow-label"
    flowLabel.EntityData.BundleName = "cisco_ios_xe"
    flowLabel.EntityData.ParentYangName = "load-balance"
    flowLabel.EntityData.SegmentPath = "flow-label"
    flowLabel.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/load-balance/" + flowLabel.EntityData.SegmentPath
    flowLabel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flowLabel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flowLabel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flowLabel.EntityData.Children = types.NewOrderedMap()
    flowLabel.EntityData.Leafs = types.NewOrderedMap()
    flowLabel.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", flowLabel.Direction})
    flowLabel.EntityData.Leafs.Append("tlv-code-17", types.YLeaf{"TlvCode17", flowLabel.TlvCode17})
    flowLabel.EntityData.Leafs.Append("static", types.YLeaf{"Static", flowLabel.Static})

    flowLabel.EntityData.YListKeys = []string {}

    return &(flowLabel.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction represents Directions to enable Flow Aware Label PW
type PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction string

const (
    // TODO
    PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction_transmit PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction = "transmit"

    // TODO
    PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction_receive PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction = "receive"

    // TODO
    PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction_both PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel_Direction = "both"
)

// PseudowireConfig_PwTemplates_PwTemplate_PreferredPath
// Preferred path for the PW
type PseudowireConfig_PwTemplates_PwTemplate_PreferredPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to a tunnel interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // TODO. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // TODO. The type is string.
    Hostname interface{}

    // Disable fall back to alternative route. The type is bool.
    DisableFallback interface{}
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetEntityData() *types.CommonEntityData {
    preferredPath.EntityData.YFilter = preferredPath.YFilter
    preferredPath.EntityData.YangName = "preferred-path"
    preferredPath.EntityData.BundleName = "cisco_ios_xe"
    preferredPath.EntityData.ParentYangName = "pw-template"
    preferredPath.EntityData.SegmentPath = "preferred-path"
    preferredPath.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + preferredPath.EntityData.SegmentPath
    preferredPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    preferredPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    preferredPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    preferredPath.EntityData.Children = types.NewOrderedMap()
    preferredPath.EntityData.Leafs = types.NewOrderedMap()
    preferredPath.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", preferredPath.Interface})
    preferredPath.EntityData.Leafs.Append("address", types.YLeaf{"Address", preferredPath.Address})
    preferredPath.EntityData.Leafs.Append("hostname", types.YLeaf{"Hostname", preferredPath.Hostname})
    preferredPath.EntityData.Leafs.Append("disable-fallback", types.YLeaf{"DisableFallback", preferredPath.DisableFallback})

    preferredPath.EntityData.YListKeys = []string {}

    return &(preferredPath.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_Sequencing
// Sequencing options
type PseudowireConfig_PwTemplates_PwTemplate_Sequencing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TODO. The type is one of the following:
    // PwSequencingReceivePwSequencingTransmitPwSequencingBoth.
    Direction interface{}

    // TODO. The type is interface{} with range: 5..65535.
    Resync interface{}
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetEntityData() *types.CommonEntityData {
    sequencing.EntityData.YFilter = sequencing.YFilter
    sequencing.EntityData.YangName = "sequencing"
    sequencing.EntityData.BundleName = "cisco_ios_xe"
    sequencing.EntityData.ParentYangName = "pw-template"
    sequencing.EntityData.SegmentPath = "sequencing"
    sequencing.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + sequencing.EntityData.SegmentPath
    sequencing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sequencing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sequencing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sequencing.EntityData.Children = types.NewOrderedMap()
    sequencing.EntityData.Leafs = types.NewOrderedMap()
    sequencing.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", sequencing.Direction})
    sequencing.EntityData.Leafs.Append("resync", types.YLeaf{"Resync", sequencing.Resync})

    sequencing.EntityData.YListKeys = []string {}

    return &(sequencing.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_Vccv
// VCCV configuration
type PseudowireConfig_PwTemplates_PwTemplate_Vccv struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable VCCV verification type. The type is bool.
    ControlWord interface{}
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetEntityData() *types.CommonEntityData {
    vccv.EntityData.YFilter = vccv.YFilter
    vccv.EntityData.YangName = "vccv"
    vccv.EntityData.BundleName = "cisco_ios_xe"
    vccv.EntityData.ParentYangName = "pw-template"
    vccv.EntityData.SegmentPath = "vccv"
    vccv.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + vccv.EntityData.SegmentPath
    vccv.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vccv.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vccv.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vccv.EntityData.Children = types.NewOrderedMap()
    vccv.EntityData.Leafs = types.NewOrderedMap()
    vccv.EntityData.Leafs.Append("control-word", types.YLeaf{"ControlWord", vccv.ControlWord})

    vccv.EntityData.YListKeys = []string {}

    return &(vccv.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay
// Timer configuration related to pseudowire redundancy
// switchover and restoring to primary
type PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies how long the backup pseudowire should wait before taking over.
    // The type is interface{} with range: 0..120.
    SwitchoverTimer interface{}

    // Specifies how long the primary pseudowire should wait after it becomes
    // active to take over for the backup pseudowire. The type is interface{} with
    // range: 0..180.
    Timer interface{}

    // The primary pseudowire never takes over for the backup. The type is bool.
    // The default value is false.
    Never interface{}
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetEntityData() *types.CommonEntityData {
    switchoverDelay.EntityData.YFilter = switchoverDelay.YFilter
    switchoverDelay.EntityData.YangName = "switchover-delay"
    switchoverDelay.EntityData.BundleName = "cisco_ios_xe"
    switchoverDelay.EntityData.ParentYangName = "pw-template"
    switchoverDelay.EntityData.SegmentPath = "switchover-delay"
    switchoverDelay.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + switchoverDelay.EntityData.SegmentPath
    switchoverDelay.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    switchoverDelay.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    switchoverDelay.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    switchoverDelay.EntityData.Children = types.NewOrderedMap()
    switchoverDelay.EntityData.Leafs = types.NewOrderedMap()
    switchoverDelay.EntityData.Leafs.Append("switchover-timer", types.YLeaf{"SwitchoverTimer", switchoverDelay.SwitchoverTimer})
    switchoverDelay.EntityData.Leafs.Append("timer", types.YLeaf{"Timer", switchoverDelay.Timer})
    switchoverDelay.EntityData.Leafs.Append("never", types.YLeaf{"Never", switchoverDelay.Never})

    switchoverDelay.EntityData.YListKeys = []string {}

    return &(switchoverDelay.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_Status
// TODO
type PseudowireConfig_PwTemplates_PwTemplate_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reflect standby status of the attachment circuit as up on the pseudowire.
    // The type is bool. The default value is false.
    Decoupled interface{}

    // Do not send pseudowire status to peer. The type is bool. The default value
    // is false.
    Disable interface{}

    // Our peer(s) are participating in a redundant solution with some form of
    // redundancy protocol running between the peer routers. Only one of the
    // remote peers will advertise a status of UP at a time. The other will
    // advertise standby. Change our configuration so we can send a status of UP
    // on both active and redundant pseudowires. The type is bool.
    PeerTopoDualHomed interface{}

    // Disable listening for routing events to trigger redundancy status changes.
    // The type is bool. The default value is false.
    RouteWatchDisable interface{}

    // Make the PE as master to enables Pseudowire Preferential Forwarding feature
    // to display the status of  the active and backup pseudowires. The type is
    // bool. The default value is false.
    RedundancyMaster interface{}
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xe"
    status.EntityData.ParentYangName = "pw-template"
    status.EntityData.SegmentPath = "status"
    status.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("decoupled", types.YLeaf{"Decoupled", status.Decoupled})
    status.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", status.Disable})
    status.EntityData.Leafs.Append("peer-topo-dual-homed", types.YLeaf{"PeerTopoDualHomed", status.PeerTopoDualHomed})
    status.EntityData.Leafs.Append("route-watch-disable", types.YLeaf{"RouteWatchDisable", status.RouteWatchDisable})
    status.EntityData.Leafs.Append("redundancy-master", types.YLeaf{"RedundancyMaster", status.RedundancyMaster})

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec
// Pseudowire port profile configurations.
type PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Description string for the port profile. The type is string.
    Description interface{}

    // Shut down this template. The type is bool. The default value is false.
    Shutdown interface{}

    // Force shut down this port profile. The type is bool. The default value is
    // false.
    ShutForce interface{}

    // MTU of the port. The type is interface{} with range: 0..4294967295.
    Mtu interface{}

    // Maximum number of ports. The type is interface{} with range: 1..16384.
    MaxPorts interface{}

    // Enable this port profile. The type is bool. The default value is false.
    Enabled interface{}
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetEntityData() *types.CommonEntityData {
    portProfileSpec.EntityData.YFilter = portProfileSpec.YFilter
    portProfileSpec.EntityData.YangName = "port-profile-spec"
    portProfileSpec.EntityData.BundleName = "cisco_ios_xe"
    portProfileSpec.EntityData.ParentYangName = "pw-template"
    portProfileSpec.EntityData.SegmentPath = "port-profile-spec"
    portProfileSpec.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-templates/pw-template/" + portProfileSpec.EntityData.SegmentPath
    portProfileSpec.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portProfileSpec.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portProfileSpec.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portProfileSpec.EntityData.Children = types.NewOrderedMap()
    portProfileSpec.EntityData.Leafs = types.NewOrderedMap()
    portProfileSpec.EntityData.Leafs.Append("description", types.YLeaf{"Description", portProfileSpec.Description})
    portProfileSpec.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", portProfileSpec.Shutdown})
    portProfileSpec.EntityData.Leafs.Append("shut-force", types.YLeaf{"ShutForce", portProfileSpec.ShutForce})
    portProfileSpec.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", portProfileSpec.Mtu})
    portProfileSpec.EntityData.Leafs.Append("max-ports", types.YLeaf{"MaxPorts", portProfileSpec.MaxPorts})
    portProfileSpec.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", portProfileSpec.Enabled})

    portProfileSpec.EntityData.YListKeys = []string {}

    return &(portProfileSpec.EntityData)
}

// PseudowireConfig_PwStaticOamClasses
// List of pseudowire static oam classes.
type PseudowireConfig_PwStaticOamClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pseudowire static oam class configuration. The type is slice of
    // PseudowireConfig_PwStaticOamClasses_PwStaticOamClass.
    PwStaticOamClass []*PseudowireConfig_PwStaticOamClasses_PwStaticOamClass
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetEntityData() *types.CommonEntityData {
    pwStaticOamClasses.EntityData.YFilter = pwStaticOamClasses.YFilter
    pwStaticOamClasses.EntityData.YangName = "pw-static-oam-classes"
    pwStaticOamClasses.EntityData.BundleName = "cisco_ios_xe"
    pwStaticOamClasses.EntityData.ParentYangName = "pseudowire-config"
    pwStaticOamClasses.EntityData.SegmentPath = "pw-static-oam-classes"
    pwStaticOamClasses.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/" + pwStaticOamClasses.EntityData.SegmentPath
    pwStaticOamClasses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwStaticOamClasses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwStaticOamClasses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwStaticOamClasses.EntityData.Children = types.NewOrderedMap()
    pwStaticOamClasses.EntityData.Children.Append("pw-static-oam-class", types.YChild{"PwStaticOamClass", nil})
    for i := range pwStaticOamClasses.PwStaticOamClass {
        pwStaticOamClasses.EntityData.Children.Append(types.GetSegmentPath(pwStaticOamClasses.PwStaticOamClass[i]), types.YChild{"PwStaticOamClass", pwStaticOamClasses.PwStaticOamClass[i]})
    }
    pwStaticOamClasses.EntityData.Leafs = types.NewOrderedMap()

    pwStaticOamClasses.EntityData.YListKeys = []string {}

    return &(pwStaticOamClasses.EntityData)
}

// PseudowireConfig_PwStaticOamClasses_PwStaticOamClass
// Pseudowire static oam class configuration
type PseudowireConfig_PwStaticOamClasses_PwStaticOamClass struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OAM class name. The type is string.
    Name interface{}

    // Enable ack. The type is bool.
    Ack interface{}

    // Keepalive in seconds. The type is interface{} with range: 1..4095. The
    // default value is 600.
    Keepalive interface{}

    // Refresh timeout in seconds. The type is interface{} with range: 1..4095.
    // The default value is 30.
    TimeoutRefreshSend interface{}

    // Ack timeout in seconds. The type is interface{} with range: 1..4095. The
    // default value is 600.
    TimeoutRefreshAck interface{}
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetEntityData() *types.CommonEntityData {
    pwStaticOamClass.EntityData.YFilter = pwStaticOamClass.YFilter
    pwStaticOamClass.EntityData.YangName = "pw-static-oam-class"
    pwStaticOamClass.EntityData.BundleName = "cisco_ios_xe"
    pwStaticOamClass.EntityData.ParentYangName = "pw-static-oam-classes"
    pwStaticOamClass.EntityData.SegmentPath = "pw-static-oam-class" + types.AddKeyToken(pwStaticOamClass.Name, "name")
    pwStaticOamClass.EntityData.AbsolutePath = "cisco-pw:pseudowire-config/pw-static-oam-classes/" + pwStaticOamClass.EntityData.SegmentPath
    pwStaticOamClass.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwStaticOamClass.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwStaticOamClass.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwStaticOamClass.EntityData.Children = types.NewOrderedMap()
    pwStaticOamClass.EntityData.Leafs = types.NewOrderedMap()
    pwStaticOamClass.EntityData.Leafs.Append("name", types.YLeaf{"Name", pwStaticOamClass.Name})
    pwStaticOamClass.EntityData.Leafs.Append("ack", types.YLeaf{"Ack", pwStaticOamClass.Ack})
    pwStaticOamClass.EntityData.Leafs.Append("keepalive", types.YLeaf{"Keepalive", pwStaticOamClass.Keepalive})
    pwStaticOamClass.EntityData.Leafs.Append("timeout-refresh-send", types.YLeaf{"TimeoutRefreshSend", pwStaticOamClass.TimeoutRefreshSend})
    pwStaticOamClass.EntityData.Leafs.Append("timeout-refresh-ack", types.YLeaf{"TimeoutRefreshAck", pwStaticOamClass.TimeoutRefreshAck})

    pwStaticOamClass.EntityData.YListKeys = []string {"Name"}

    return &(pwStaticOamClass.EntityData)
}

// PseudowireState
// Contains the operational state for all the pseudowire
// instances in the switch, no matter what L2VPN service
// created them.
type PseudowireState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each list element contains the state for a pseudowire instance. The list
    // can be optionally keyed by any combination of vc-id, peer-address, etc.
    // Additional filtering of the list by the agent may be performed upon request
    // by the client using subtree filtering as described in RFC 6020 Section 6.
    // The type is slice of PseudowireState_Pseudowires.
    Pseudowires []*PseudowireState_Pseudowires
}

func (pseudowireState *PseudowireState) GetEntityData() *types.CommonEntityData {
    pseudowireState.EntityData.YFilter = pseudowireState.YFilter
    pseudowireState.EntityData.YangName = "pseudowire-state"
    pseudowireState.EntityData.BundleName = "cisco_ios_xe"
    pseudowireState.EntityData.ParentYangName = "cisco-pw"
    pseudowireState.EntityData.SegmentPath = "cisco-pw:pseudowire-state"
    pseudowireState.EntityData.AbsolutePath = pseudowireState.EntityData.SegmentPath
    pseudowireState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pseudowireState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pseudowireState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pseudowireState.EntityData.Children = types.NewOrderedMap()
    pseudowireState.EntityData.Children.Append("pseudowires", types.YChild{"Pseudowires", nil})
    for i := range pseudowireState.Pseudowires {
        pseudowireState.EntityData.Children.Append(types.GetSegmentPath(pseudowireState.Pseudowires[i]), types.YChild{"Pseudowires", pseudowireState.Pseudowires[i]})
    }
    pseudowireState.EntityData.Leafs = types.NewOrderedMap()

    pseudowireState.EntityData.YListKeys = []string {}

    return &(pseudowireState.EntityData)
}

// PseudowireState_Pseudowires
// Each list element contains the state for a pseudowire
// instance. The list can be optionally keyed by any
// combination of vc-id, peer-address, etc.
// Additional filtering of the list by the agent may be
// performed upon request by the client using subtree
// filtering as described in RFC 6020 Section 6.
type PseudowireState_Pseudowires struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object contains the value of the peer node
    // address of the PW/PE protocol entity. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VcPeerAddress interface{}

    // This attribute is a key. Used to distinguish between pseudowires going to
    // the same peer node. The type is interface{} with range: 0..4294967295.
    VcId interface{}

    // This attribute is a key. Indicates the service responsible for establishing
    // this VC. The type is VcOwnerType.
    VcOwnerType interface{}

    // This attribute is a key. The canonical name assigned to the VC. The name
    // may be autogenerated by the device; this Yang model does not currently
    // support direct configuration of this name. The type is string.
    VcName interface{}

    // This attribute is a key. Locally-unique ID within the device for this
    // pseudowire. The type is interface{} with range: 0..4294967295.
    VcIndex interface{}

    // Indicates the service to be carried over this VC. The type is one of the
    // following: PwVcTypeEtherPwVcTypeVlanPwVcTypeVlanPassthrough.
    VcType interface{}

    // Name of the L2VPN service instance that created the pseudowire VC. The type
    // is string.
    VcOwnerName interface{}

    // Indicates the type of packet-switched network that carries this VC. The
    // type is VcPsnType.
    VcPsnType interface{}

    // Used to identify which local pseudowire group this pseudowire belongs to.
    // The type is interface{} with range: 0..4294967295.
    VcLocalGroupId interface{}

    // Indicates if the control word is sent with each packet by the local node.
    // The type is bool.
    VcControlWord interface{}

    // If not zero, this represents the locally supported MTU size over the
    // interface (or the virtual interface) associated with the VC. The type is
    // interface{} with range: 0..4294967295.
    VcLocalIfMtu interface{}

    // If not zero, indicates the pseudowire group to which this pseudowire
    // belongs on the peer node. The type is interface{} with range:
    // 0..4294967295.
    VcRemoteGroupId interface{}

    // Indicates whether MPLS control words are used by the pseudowire PSN
    // service. The type is VcRemoteControlWord.
    VcRemoteControlWord interface{}

    // The remote interface MTU as (optionally) received from the remote node via
    // the signaling protocol. Should be zero if this parameter is not available
    // or not used. The type is interface{} with range: 0..4294967295.
    VcRemoteIfMtu interface{}

    // The VC label used in the outbound direction (i.e. toward the PSN). Example:
    // for MPLS PSN, it represents the 20 bits of VC tag, for L2TP it represent
    // the 32 bits Session ID. If the label is not yet known (signaling in
    // procesS), the object should return a value of 0xFFFFFFFF. The type is
    // interface{} with range: 0..4294967295.
    VcOutboundLabel interface{}

    // The VC label used in the inbound direction (i.e. packets received from the
    // PSN). Example: for MPLS PSN, it represents the 20 bits of VC tag, for L2TP
    // it represents the 32 bits Session ID. If the label is not yet known
    // (signaling in process), the object should return a value of 0xFFFFFFFF. The
    // type is interface{} with range: 0..4294967295.
    VcInboundLabel interface{}

    // Indicates the actual combined operational status of this VC. It is 'up' if
    // both vc-inbound-oper-status and vc-outbound-oper-status are in 'up' state.
    // For all other values, if the VCs in both directions are of the same value
    // it reflects that value, otherwise it is set to the most severe status out
    // of the two statuses. The order of severance from most severe to less severe
    // is: unknown, notPresent, down, lowerLayerDown, dormant, testing, up. The
    // operator may consult the per direction oper-status for fault isolation per
    // direction. The type is PwOperStateType.
    VcOperStatus interface{}

    // Indicates the actual operational status of this VC in the  inbound
    // direction. The type is PwOperStateType.
    VcInboundOperStatus interface{}

    // Indicates the actual operational status of this VC in the outbound
    // direction. The type is PwOperStateType.
    VcOutboundOperStatus interface{}

    // A collection of pseudowire-related statistics objects.
    Statistics PseudowireState_Pseudowires_Statistics
}

func (pseudowires *PseudowireState_Pseudowires) GetEntityData() *types.CommonEntityData {
    pseudowires.EntityData.YFilter = pseudowires.YFilter
    pseudowires.EntityData.YangName = "pseudowires"
    pseudowires.EntityData.BundleName = "cisco_ios_xe"
    pseudowires.EntityData.ParentYangName = "pseudowire-state"
    pseudowires.EntityData.SegmentPath = "pseudowires" + types.AddKeyToken(pseudowires.VcPeerAddress, "vc-peer-address") + types.AddKeyToken(pseudowires.VcId, "vc-id") + types.AddKeyToken(pseudowires.VcOwnerType, "vc-owner-type") + types.AddKeyToken(pseudowires.VcName, "vc-name") + types.AddKeyToken(pseudowires.VcIndex, "vc-index")
    pseudowires.EntityData.AbsolutePath = "cisco-pw:pseudowire-state/" + pseudowires.EntityData.SegmentPath
    pseudowires.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pseudowires.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pseudowires.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pseudowires.EntityData.Children = types.NewOrderedMap()
    pseudowires.EntityData.Children.Append("statistics", types.YChild{"Statistics", &pseudowires.Statistics})
    pseudowires.EntityData.Leafs = types.NewOrderedMap()
    pseudowires.EntityData.Leafs.Append("vc-peer-address", types.YLeaf{"VcPeerAddress", pseudowires.VcPeerAddress})
    pseudowires.EntityData.Leafs.Append("vc-id", types.YLeaf{"VcId", pseudowires.VcId})
    pseudowires.EntityData.Leafs.Append("vc-owner-type", types.YLeaf{"VcOwnerType", pseudowires.VcOwnerType})
    pseudowires.EntityData.Leafs.Append("vc-name", types.YLeaf{"VcName", pseudowires.VcName})
    pseudowires.EntityData.Leafs.Append("vc-index", types.YLeaf{"VcIndex", pseudowires.VcIndex})
    pseudowires.EntityData.Leafs.Append("vc-type", types.YLeaf{"VcType", pseudowires.VcType})
    pseudowires.EntityData.Leafs.Append("vc-owner-name", types.YLeaf{"VcOwnerName", pseudowires.VcOwnerName})
    pseudowires.EntityData.Leafs.Append("vc-psn-type", types.YLeaf{"VcPsnType", pseudowires.VcPsnType})
    pseudowires.EntityData.Leafs.Append("vc-local-group-id", types.YLeaf{"VcLocalGroupId", pseudowires.VcLocalGroupId})
    pseudowires.EntityData.Leafs.Append("vc-control-word", types.YLeaf{"VcControlWord", pseudowires.VcControlWord})
    pseudowires.EntityData.Leafs.Append("vc-local-if-mtu", types.YLeaf{"VcLocalIfMtu", pseudowires.VcLocalIfMtu})
    pseudowires.EntityData.Leafs.Append("vc-remote-group-id", types.YLeaf{"VcRemoteGroupId", pseudowires.VcRemoteGroupId})
    pseudowires.EntityData.Leafs.Append("vc-remote-control-word", types.YLeaf{"VcRemoteControlWord", pseudowires.VcRemoteControlWord})
    pseudowires.EntityData.Leafs.Append("vc-remote-if-mtu", types.YLeaf{"VcRemoteIfMtu", pseudowires.VcRemoteIfMtu})
    pseudowires.EntityData.Leafs.Append("vc-outbound-label", types.YLeaf{"VcOutboundLabel", pseudowires.VcOutboundLabel})
    pseudowires.EntityData.Leafs.Append("vc-inbound-label", types.YLeaf{"VcInboundLabel", pseudowires.VcInboundLabel})
    pseudowires.EntityData.Leafs.Append("vc-oper-status", types.YLeaf{"VcOperStatus", pseudowires.VcOperStatus})
    pseudowires.EntityData.Leafs.Append("vc-inbound-oper-status", types.YLeaf{"VcInboundOperStatus", pseudowires.VcInboundOperStatus})
    pseudowires.EntityData.Leafs.Append("vc-outbound-oper-status", types.YLeaf{"VcOutboundOperStatus", pseudowires.VcOutboundOperStatus})

    pseudowires.EntityData.YListKeys = []string {"VcPeerAddress", "VcId", "VcOwnerType", "VcName", "VcIndex"}

    return &(pseudowires.EntityData)
}

// PseudowireState_Pseudowires_Statistics
// A collection of pseudowire-related statistics objects
type PseudowireState_Pseudowires_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // System time when this VC was created. The type is interface{} with range:
    // 0..4294967295.
    VcCreateTime interface{}

    // Number of consecutive ticks this VC has been 'up' in both directions
    // together. The type is interface{} with range: 0..4294967295.
    VcUpTime interface{}

    // The time on the most recent occasion at which any of this pseudowire's
    // counters sufferred discontinuity. If no such discontinuities have occurred
    // since the last re-initialization of the local management subsystem, then
    // this node contains the time the local management subsystem re-initialized
    // itself. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    // This attribute is mandatory.
    DiscontinuityTime interface{}

    // The total number of octets received on this pseudowire.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctets interface{}

    // The total number of packets received on this pseudowire.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InPkts interface{}

    // The total number of inbound packets that contained errors.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    InErrors interface{}

    // The total number of octets sent on this pseudowire.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctets interface{}

    // The total number of packets sent on this pseudowire.  Discontinuities in
    // the value of this counter can occur at re-initialization of the management
    // system and at other times as indicated by the value of
    // 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPkts interface{}

    // The total number of outbound packets that could not be sent on this
    // pseudowire.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system and at other times as indicated
    // by the value of 'discontinuity-time'. The type is interface{} with range:
    // 0..18446744073709551615.
    OutErrors interface{}
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xe"
    statistics.EntityData.ParentYangName = "pseudowires"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "cisco-pw:pseudowire-state/pseudowires/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("vc-create-time", types.YLeaf{"VcCreateTime", statistics.VcCreateTime})
    statistics.EntityData.Leafs.Append("vc-up-time", types.YLeaf{"VcUpTime", statistics.VcUpTime})
    statistics.EntityData.Leafs.Append("discontinuity-time", types.YLeaf{"DiscontinuityTime", statistics.DiscontinuityTime})
    statistics.EntityData.Leafs.Append("in-octets", types.YLeaf{"InOctets", statistics.InOctets})
    statistics.EntityData.Leafs.Append("in-pkts", types.YLeaf{"InPkts", statistics.InPkts})
    statistics.EntityData.Leafs.Append("in-errors", types.YLeaf{"InErrors", statistics.InErrors})
    statistics.EntityData.Leafs.Append("out-octets", types.YLeaf{"OutOctets", statistics.OutOctets})
    statistics.EntityData.Leafs.Append("out-pkts", types.YLeaf{"OutPkts", statistics.OutPkts})
    statistics.EntityData.Leafs.Append("out-errors", types.YLeaf{"OutErrors", statistics.OutErrors})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// PseudowireState_Pseudowires_VcOwnerType represents this VC
type PseudowireState_Pseudowires_VcOwnerType string

const (
    // Created under VPWS context
    PseudowireState_Pseudowires_VcOwnerType_vpws PseudowireState_Pseudowires_VcOwnerType = "vpws"

    // Created under VFI context
    PseudowireState_Pseudowires_VcOwnerType_vpls_vfi PseudowireState_Pseudowires_VcOwnerType = "vpls-vfi"

    // Created under bridge domain context
    PseudowireState_Pseudowires_VcOwnerType_vpls_bridge_domain PseudowireState_Pseudowires_VcOwnerType = "vpls-bridge-domain"

    // Created by manual pseudowire interface
    // configuration
    PseudowireState_Pseudowires_VcOwnerType_interface_ PseudowireState_Pseudowires_VcOwnerType = "interface"
)

// PseudowireState_Pseudowires_VcPsnType represents that carries this VC
type PseudowireState_Pseudowires_VcPsnType string

const (
    // MPLS encapsulation
    PseudowireState_Pseudowires_VcPsnType_mpls PseudowireState_Pseudowires_VcPsnType = "mpls"

    // L2TP encapsulation
    PseudowireState_Pseudowires_VcPsnType_l2tp PseudowireState_Pseudowires_VcPsnType = "l2tp"

    // IP core
    PseudowireState_Pseudowires_VcPsnType_ip PseudowireState_Pseudowires_VcPsnType = "ip"

    // MPLS encapsulation over IP core
    PseudowireState_Pseudowires_VcPsnType_mpls_over_ip PseudowireState_Pseudowires_VcPsnType = "mpls-over-ip"

    // GRE encapsulation
    PseudowireState_Pseudowires_VcPsnType_gre PseudowireState_Pseudowires_VcPsnType = "gre"

    // Some other encapsulation
    PseudowireState_Pseudowires_VcPsnType_other PseudowireState_Pseudowires_VcPsnType = "other"
)

// PseudowireState_Pseudowires_VcRemoteControlWord represents pseudowire PSN service
type PseudowireState_Pseudowires_VcRemoteControlWord string

const (
    // Peer sends control word
    PseudowireState_Pseudowires_VcRemoteControlWord_noControlWord PseudowireState_Pseudowires_VcRemoteControlWord = "noControlWord"

    // Peer does not send control word
    PseudowireState_Pseudowires_VcRemoteControlWord_withControlWord PseudowireState_Pseudowires_VcRemoteControlWord = "withControlWord"

    // Have not received indication yet if peer sends
    // control word
    PseudowireState_Pseudowires_VcRemoteControlWord_notYetKnown PseudowireState_Pseudowires_VcRemoteControlWord = "notYetKnown"
)

