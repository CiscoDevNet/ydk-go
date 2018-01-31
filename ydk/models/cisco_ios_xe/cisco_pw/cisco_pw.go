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

type PwVcTypeEther struct {
}

func (id PwVcTypeEther) String() string {
	return "cisco-pw:pw-vc-type-ether"
}

type PwLbEthSrcDstMac struct {
}

func (id PwLbEthSrcDstMac) String() string {
	return "cisco-pw:pw-lb-eth-src-dst-mac"
}

type PwSignalingProtocolType struct {
}

func (id PwSignalingProtocolType) String() string {
	return "cisco-pw:pw-signaling-protocol-type"
}

type PwLoadBalanceType struct {
}

func (id PwLoadBalanceType) String() string {
	return "cisco-pw:pw-load-balance-type"
}

type PwSequencingTransmit struct {
}

func (id PwSequencingTransmit) String() string {
	return "cisco-pw:pw-sequencing-transmit"
}

type PwVcTypeVlanPassthrough struct {
}

func (id PwVcTypeVlanPassthrough) String() string {
	return "cisco-pw:pw-vc-type-vlan-passthrough"
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

type PwLbEthDstMac struct {
}

func (id PwLbEthDstMac) String() string {
	return "cisco-pw:pw-lb-eth-dst-mac"
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

type PwSequencingReceive struct {
}

func (id PwSequencingReceive) String() string {
	return "cisco-pw:pw-sequencing-receive"
}

type PwLbEthSrcMac struct {
}

func (id PwLbEthSrcMac) String() string {
	return "cisco-pw:pw-lb-eth-src-mac"
}

type PwLbEthernetType struct {
}

func (id PwLbEthernetType) String() string {
	return "cisco-pw:pw-lb-ethernet-type"
}

type PwSequencingType struct {
}

func (id PwSequencingType) String() string {
	return "cisco-pw:pw-sequencing-type"
}

type PwSignalingProtocolLdp struct {
}

func (id PwSignalingProtocolLdp) String() string {
	return "cisco-pw:pw-signaling-protocol-ldp"
}

type PwSequencingBoth struct {
}

func (id PwSequencingBoth) String() string {
	return "cisco-pw:pw-sequencing-both"
}

type PwVcTypeVlan struct {
}

func (id PwVcTypeVlan) String() string {
	return "cisco-pw:pw-vc-type-vlan"
}

type PwLbIpType struct {
}

func (id PwLbIpType) String() string {
	return "cisco-pw:pw-lb-ip-type"
}

type PwSignalingProtocolNone struct {
}

func (id PwSignalingProtocolNone) String() string {
	return "cisco-pw:pw-signaling-protocol-none"
}

type PwSignalingProtocolBgp struct {
}

func (id PwSignalingProtocolBgp) String() string {
	return "cisco-pw:pw-signaling-protocol-bgp"
}

type PwLbIpSrcIp struct {
}

func (id PwLbIpSrcIp) String() string {
	return "cisco-pw:pw-lb-ip-src-ip"
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Global configurations related to pseudowires.
    Global PseudowireConfig_Global

    // Contains list of all pw-template definitions. Also called PW Class (XR) and
    // Port Profile (NX-OS).
    PwTemplates PseudowireConfig_PwTemplates

    // List of pseudowire static oam classes.
    PwStaticOamClasses PseudowireConfig_PwStaticOamClasses
}

func (pseudowireConfig *PseudowireConfig) GetFilter() yfilter.YFilter { return pseudowireConfig.YFilter }

func (pseudowireConfig *PseudowireConfig) SetFilter(yf yfilter.YFilter) { pseudowireConfig.YFilter = yf }

func (pseudowireConfig *PseudowireConfig) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "pw-templates" { return "PwTemplates" }
    if yname == "pw-static-oam-classes" { return "PwStaticOamClasses" }
    return ""
}

func (pseudowireConfig *PseudowireConfig) GetSegmentPath() string {
    return "cisco-pw:pseudowire-config"
}

func (pseudowireConfig *PseudowireConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &pseudowireConfig.Global
    }
    if childYangName == "pw-templates" {
        return &pseudowireConfig.PwTemplates
    }
    if childYangName == "pw-static-oam-classes" {
        return &pseudowireConfig.PwStaticOamClasses
    }
    return nil
}

func (pseudowireConfig *PseudowireConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &pseudowireConfig.Global
    children["pw-templates"] = &pseudowireConfig.PwTemplates
    children["pw-static-oam-classes"] = &pseudowireConfig.PwStaticOamClasses
    return children
}

func (pseudowireConfig *PseudowireConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pseudowireConfig *PseudowireConfig) GetBundleName() string { return "cisco_ios_xe" }

func (pseudowireConfig *PseudowireConfig) GetYangName() string { return "pseudowire-config" }

func (pseudowireConfig *PseudowireConfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pseudowireConfig *PseudowireConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pseudowireConfig *PseudowireConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pseudowireConfig *PseudowireConfig) SetParent(parent types.Entity) { pseudowireConfig.parent = parent }

func (pseudowireConfig *PseudowireConfig) GetParent() types.Entity { return pseudowireConfig.parent }

func (pseudowireConfig *PseudowireConfig) GetParentYangName() string { return "cisco-pw" }

// PseudowireConfig_Global
// Global configurations related to pseudowires.
type PseudowireConfig_Global struct {
    parent types.Entity
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

func (global *PseudowireConfig_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *PseudowireConfig_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *PseudowireConfig_Global) GetGoName(yname string) string {
    if yname == "pw-grouping" { return "PwGrouping" }
    if yname == "pw-oam-refresh-transmit" { return "PwOamRefreshTransmit" }
    if yname == "pw-status" { return "PwStatus" }
    if yname == "predictive-redundancy" { return "PredictiveRedundancy" }
    if yname == "vc-state-notification-enabled" { return "VcStateNotificationEnabled" }
    if yname == "vc-state-notification-batch-size" { return "VcStateNotificationBatchSize" }
    if yname == "vc-state-notification-rate" { return "VcStateNotificationRate" }
    return ""
}

func (global *PseudowireConfig_Global) GetSegmentPath() string {
    return "global"
}

func (global *PseudowireConfig_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (global *PseudowireConfig_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (global *PseudowireConfig_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pw-grouping"] = global.PwGrouping
    leafs["pw-oam-refresh-transmit"] = global.PwOamRefreshTransmit
    leafs["pw-status"] = global.PwStatus
    leafs["predictive-redundancy"] = global.PredictiveRedundancy
    leafs["vc-state-notification-enabled"] = global.VcStateNotificationEnabled
    leafs["vc-state-notification-batch-size"] = global.VcStateNotificationBatchSize
    leafs["vc-state-notification-rate"] = global.VcStateNotificationRate
    return leafs
}

func (global *PseudowireConfig_Global) GetBundleName() string { return "cisco_ios_xe" }

func (global *PseudowireConfig_Global) GetYangName() string { return "global" }

func (global *PseudowireConfig_Global) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (global *PseudowireConfig_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (global *PseudowireConfig_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (global *PseudowireConfig_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *PseudowireConfig_Global) GetParent() types.Entity { return global.parent }

func (global *PseudowireConfig_Global) GetParentYangName() string { return "pseudowire-config" }

// PseudowireConfig_PwTemplates
// Contains list of all pw-template definitions.
// Also called PW Class (XR) and Port Profile (NX-OS)
type PseudowireConfig_PwTemplates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pseudowire template list. The type is slice of
    // PseudowireConfig_PwTemplates_PwTemplate.
    PwTemplate []PseudowireConfig_PwTemplates_PwTemplate
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetFilter() yfilter.YFilter { return pwTemplates.YFilter }

func (pwTemplates *PseudowireConfig_PwTemplates) SetFilter(yf yfilter.YFilter) { pwTemplates.YFilter = yf }

func (pwTemplates *PseudowireConfig_PwTemplates) GetGoName(yname string) string {
    if yname == "pw-template" { return "PwTemplate" }
    return ""
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetSegmentPath() string {
    return "pw-templates"
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pw-template" {
        for _, c := range pwTemplates.PwTemplate {
            if pwTemplates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PseudowireConfig_PwTemplates_PwTemplate{}
        pwTemplates.PwTemplate = append(pwTemplates.PwTemplate, child)
        return &pwTemplates.PwTemplate[len(pwTemplates.PwTemplate)-1]
    }
    return nil
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pwTemplates.PwTemplate {
        children[pwTemplates.PwTemplate[i].GetSegmentPath()] = &pwTemplates.PwTemplate[i]
    }
    return children
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pwTemplates *PseudowireConfig_PwTemplates) GetBundleName() string { return "cisco_ios_xe" }

func (pwTemplates *PseudowireConfig_PwTemplates) GetYangName() string { return "pw-templates" }

func (pwTemplates *PseudowireConfig_PwTemplates) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pwTemplates *PseudowireConfig_PwTemplates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pwTemplates *PseudowireConfig_PwTemplates) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pwTemplates *PseudowireConfig_PwTemplates) SetParent(parent types.Entity) { pwTemplates.parent = parent }

func (pwTemplates *PseudowireConfig_PwTemplates) GetParent() types.Entity { return pwTemplates.parent }

func (pwTemplates *PseudowireConfig_PwTemplates) GetParentYangName() string { return "pseudowire-config" }

// PseudowireConfig_PwTemplates_PwTemplate
// Pseudowire template list.
type PseudowireConfig_PwTemplates_PwTemplate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. PW Template name. The type is string.
    Name interface{}

    // Encapsulation used for PW. The type is one of the following: PwEncapMpls.
    Encapsulation interface{}

    // Use control word in the MPLS PW header. The type is bool.
    ControlWord interface{}

    // Signaling protocol to use. The type is one of the following:
    // PwSignalingProtocolLdpPwSignalingProtocolNonePwSignalingProtocolBgp.
    SignalingProtocol interface{}

    // Type of VC in the PW. The type is one of the following:
    // PwVcTypeEtherPwVcTypeVlanPassthroughPwVcTypeVlan.
    VcType interface{}

    // Send switching TLV. The type is bool.
    SwitchingTlv interface{}

    // The local source IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetFilter() yfilter.YFilter { return pwTemplate.YFilter }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) SetFilter(yf yfilter.YFilter) { pwTemplate.YFilter = yf }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "encapsulation" { return "Encapsulation" }
    if yname == "control-word" { return "ControlWord" }
    if yname == "signaling-protocol" { return "SignalingProtocol" }
    if yname == "vc-type" { return "VcType" }
    if yname == "switching-tlv" { return "SwitchingTlv" }
    if yname == "source-ip" { return "SourceIp" }
    if yname == "tag-rewrite-ingress-vlan" { return "TagRewriteIngressVlan" }
    if yname == "mac-withdraw" { return "MacWithdraw" }
    if yname == "load-balance" { return "LoadBalance" }
    if yname == "preferred-path" { return "PreferredPath" }
    if yname == "sequencing" { return "Sequencing" }
    if yname == "vccv" { return "Vccv" }
    if yname == "switchover-delay" { return "SwitchoverDelay" }
    if yname == "status" { return "Status" }
    if yname == "port-profile-spec" { return "PortProfileSpec" }
    return ""
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetSegmentPath() string {
    return "pw-template" + "[name='" + fmt.Sprintf("%v", pwTemplate.Name) + "']"
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "load-balance" {
        return &pwTemplate.LoadBalance
    }
    if childYangName == "preferred-path" {
        return &pwTemplate.PreferredPath
    }
    if childYangName == "sequencing" {
        return &pwTemplate.Sequencing
    }
    if childYangName == "vccv" {
        return &pwTemplate.Vccv
    }
    if childYangName == "switchover-delay" {
        return &pwTemplate.SwitchoverDelay
    }
    if childYangName == "status" {
        return &pwTemplate.Status
    }
    if childYangName == "port-profile-spec" {
        return &pwTemplate.PortProfileSpec
    }
    return nil
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["load-balance"] = &pwTemplate.LoadBalance
    children["preferred-path"] = &pwTemplate.PreferredPath
    children["sequencing"] = &pwTemplate.Sequencing
    children["vccv"] = &pwTemplate.Vccv
    children["switchover-delay"] = &pwTemplate.SwitchoverDelay
    children["status"] = &pwTemplate.Status
    children["port-profile-spec"] = &pwTemplate.PortProfileSpec
    return children
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = pwTemplate.Name
    leafs["encapsulation"] = pwTemplate.Encapsulation
    leafs["control-word"] = pwTemplate.ControlWord
    leafs["signaling-protocol"] = pwTemplate.SignalingProtocol
    leafs["vc-type"] = pwTemplate.VcType
    leafs["switching-tlv"] = pwTemplate.SwitchingTlv
    leafs["source-ip"] = pwTemplate.SourceIp
    leafs["tag-rewrite-ingress-vlan"] = pwTemplate.TagRewriteIngressVlan
    leafs["mac-withdraw"] = pwTemplate.MacWithdraw
    return leafs
}

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetBundleName() string { return "cisco_ios_xe" }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetYangName() string { return "pw-template" }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) SetParent(parent types.Entity) { pwTemplate.parent = parent }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetParent() types.Entity { return pwTemplate.parent }

func (pwTemplate *PseudowireConfig_PwTemplates_PwTemplate) GetParentYangName() string { return "pw-templates" }

// PseudowireConfig_PwTemplates_PwTemplate_LoadBalance
// Load balancing mechanism.
type PseudowireConfig_PwTemplates_PwTemplate_LoadBalance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ethernet mac address based load balancing. The type is one of the
    // following: PwLbEthSrcDstMacPwLbEthDstMacPwLbEthSrcMac. The default value is
    // pw-lb-eth-src-dst-mac.
    Ethernet interface{}

    // IP address based load balancing. The type is one of the following:
    // PwLbIpSrcDstIpPwLbIpSrcIp.
    Ip interface{}

    // Enable Flow Aware Label (FAT) PW - the capability to carry flow label on
    // PW.
    FlowLabel PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetFilter() yfilter.YFilter { return loadBalance.YFilter }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) SetFilter(yf yfilter.YFilter) { loadBalance.YFilter = yf }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetGoName(yname string) string {
    if yname == "ethernet" { return "Ethernet" }
    if yname == "ip" { return "Ip" }
    if yname == "flow-label" { return "FlowLabel" }
    return ""
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetSegmentPath() string {
    return "load-balance"
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flow-label" {
        return &loadBalance.FlowLabel
    }
    return nil
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flow-label"] = &loadBalance.FlowLabel
    return children
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ethernet"] = loadBalance.Ethernet
    leafs["ip"] = loadBalance.Ip
    return leafs
}

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetBundleName() string { return "cisco_ios_xe" }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetYangName() string { return "load-balance" }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) SetParent(parent types.Entity) { loadBalance.parent = parent }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetParent() types.Entity { return loadBalance.parent }

func (loadBalance *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel
// Enable Flow Aware Label (FAT) PW - the capability to
// carry flow label on PW
type PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel struct {
    parent types.Entity
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

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetFilter() yfilter.YFilter { return flowLabel.YFilter }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) SetFilter(yf yfilter.YFilter) { flowLabel.YFilter = yf }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "tlv-code-17" { return "TlvCode17" }
    if yname == "static" { return "Static" }
    return ""
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetSegmentPath() string {
    return "flow-label"
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = flowLabel.Direction
    leafs["tlv-code-17"] = flowLabel.TlvCode17
    leafs["static"] = flowLabel.Static
    return leafs
}

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetBundleName() string { return "cisco_ios_xe" }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetYangName() string { return "flow-label" }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) SetParent(parent types.Entity) { flowLabel.parent = parent }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetParent() types.Entity { return flowLabel.parent }

func (flowLabel *PseudowireConfig_PwTemplates_PwTemplate_LoadBalance_FlowLabel) GetParentYangName() string { return "load-balance" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to a tunnel interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // TODO. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // TODO. The type is string.
    Hostname interface{}

    // Disable fall back to alternative route. The type is bool.
    DisableFallback interface{}
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetFilter() yfilter.YFilter { return preferredPath.YFilter }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) SetFilter(yf yfilter.YFilter) { preferredPath.YFilter = yf }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "address" { return "Address" }
    if yname == "hostname" { return "Hostname" }
    if yname == "disable-fallback" { return "DisableFallback" }
    return ""
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetSegmentPath() string {
    return "preferred-path"
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = preferredPath.Interface
    leafs["address"] = preferredPath.Address
    leafs["hostname"] = preferredPath.Hostname
    leafs["disable-fallback"] = preferredPath.DisableFallback
    return leafs
}

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetBundleName() string { return "cisco_ios_xe" }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetYangName() string { return "preferred-path" }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) SetParent(parent types.Entity) { preferredPath.parent = parent }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetParent() types.Entity { return preferredPath.parent }

func (preferredPath *PseudowireConfig_PwTemplates_PwTemplate_PreferredPath) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_Sequencing
// Sequencing options
type PseudowireConfig_PwTemplates_PwTemplate_Sequencing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TODO. The type is one of the following:
    // PwSequencingTransmitPwSequencingReceivePwSequencingBoth.
    Direction interface{}

    // TODO. The type is interface{} with range: 5..65535.
    Resync interface{}
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetFilter() yfilter.YFilter { return sequencing.YFilter }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) SetFilter(yf yfilter.YFilter) { sequencing.YFilter = yf }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "resync" { return "Resync" }
    return ""
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetSegmentPath() string {
    return "sequencing"
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = sequencing.Direction
    leafs["resync"] = sequencing.Resync
    return leafs
}

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetBundleName() string { return "cisco_ios_xe" }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetYangName() string { return "sequencing" }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) SetParent(parent types.Entity) { sequencing.parent = parent }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetParent() types.Entity { return sequencing.parent }

func (sequencing *PseudowireConfig_PwTemplates_PwTemplate_Sequencing) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_Vccv
// VCCV configuration
type PseudowireConfig_PwTemplates_PwTemplate_Vccv struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable VCCV verification type. The type is bool.
    ControlWord interface{}
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetFilter() yfilter.YFilter { return vccv.YFilter }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) SetFilter(yf yfilter.YFilter) { vccv.YFilter = yf }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetGoName(yname string) string {
    if yname == "control-word" { return "ControlWord" }
    return ""
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetSegmentPath() string {
    return "vccv"
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["control-word"] = vccv.ControlWord
    return leafs
}

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetBundleName() string { return "cisco_ios_xe" }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetYangName() string { return "vccv" }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) SetParent(parent types.Entity) { vccv.parent = parent }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetParent() types.Entity { return vccv.parent }

func (vccv *PseudowireConfig_PwTemplates_PwTemplate_Vccv) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay
// Timer configuration related to pseudowire redundancy
// switchover and restoring to primary
type PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay struct {
    parent types.Entity
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

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetFilter() yfilter.YFilter { return switchoverDelay.YFilter }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) SetFilter(yf yfilter.YFilter) { switchoverDelay.YFilter = yf }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetGoName(yname string) string {
    if yname == "switchover-timer" { return "SwitchoverTimer" }
    if yname == "timer" { return "Timer" }
    if yname == "never" { return "Never" }
    return ""
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetSegmentPath() string {
    return "switchover-delay"
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switchover-timer"] = switchoverDelay.SwitchoverTimer
    leafs["timer"] = switchoverDelay.Timer
    leafs["never"] = switchoverDelay.Never
    return leafs
}

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetBundleName() string { return "cisco_ios_xe" }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetYangName() string { return "switchover-delay" }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) SetParent(parent types.Entity) { switchoverDelay.parent = parent }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetParent() types.Entity { return switchoverDelay.parent }

func (switchoverDelay *PseudowireConfig_PwTemplates_PwTemplate_SwitchoverDelay) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_Status
// TODO
type PseudowireConfig_PwTemplates_PwTemplate_Status struct {
    parent types.Entity
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

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetGoName(yname string) string {
    if yname == "decoupled" { return "Decoupled" }
    if yname == "disable" { return "Disable" }
    if yname == "peer-topo-dual-homed" { return "PeerTopoDualHomed" }
    if yname == "route-watch-disable" { return "RouteWatchDisable" }
    if yname == "redundancy-master" { return "RedundancyMaster" }
    return ""
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetSegmentPath() string {
    return "status"
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["decoupled"] = status.Decoupled
    leafs["disable"] = status.Disable
    leafs["peer-topo-dual-homed"] = status.PeerTopoDualHomed
    leafs["route-watch-disable"] = status.RouteWatchDisable
    leafs["redundancy-master"] = status.RedundancyMaster
    return leafs
}

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetBundleName() string { return "cisco_ios_xe" }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetYangName() string { return "status" }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetParent() types.Entity { return status.parent }

func (status *PseudowireConfig_PwTemplates_PwTemplate_Status) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec
// Pseudowire port profile configurations.
type PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec struct {
    parent types.Entity
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

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetFilter() yfilter.YFilter { return portProfileSpec.YFilter }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) SetFilter(yf yfilter.YFilter) { portProfileSpec.YFilter = yf }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetGoName(yname string) string {
    if yname == "description" { return "Description" }
    if yname == "shutdown" { return "Shutdown" }
    if yname == "shut-force" { return "ShutForce" }
    if yname == "mtu" { return "Mtu" }
    if yname == "max-ports" { return "MaxPorts" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetSegmentPath() string {
    return "port-profile-spec"
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["description"] = portProfileSpec.Description
    leafs["shutdown"] = portProfileSpec.Shutdown
    leafs["shut-force"] = portProfileSpec.ShutForce
    leafs["mtu"] = portProfileSpec.Mtu
    leafs["max-ports"] = portProfileSpec.MaxPorts
    leafs["enabled"] = portProfileSpec.Enabled
    return leafs
}

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetBundleName() string { return "cisco_ios_xe" }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetYangName() string { return "port-profile-spec" }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) SetParent(parent types.Entity) { portProfileSpec.parent = parent }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetParent() types.Entity { return portProfileSpec.parent }

func (portProfileSpec *PseudowireConfig_PwTemplates_PwTemplate_PortProfileSpec) GetParentYangName() string { return "pw-template" }

// PseudowireConfig_PwStaticOamClasses
// List of pseudowire static oam classes.
type PseudowireConfig_PwStaticOamClasses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pseudowire static oam class configuration. The type is slice of
    // PseudowireConfig_PwStaticOamClasses_PwStaticOamClass.
    PwStaticOamClass []PseudowireConfig_PwStaticOamClasses_PwStaticOamClass
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetFilter() yfilter.YFilter { return pwStaticOamClasses.YFilter }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) SetFilter(yf yfilter.YFilter) { pwStaticOamClasses.YFilter = yf }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetGoName(yname string) string {
    if yname == "pw-static-oam-class" { return "PwStaticOamClass" }
    return ""
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetSegmentPath() string {
    return "pw-static-oam-classes"
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pw-static-oam-class" {
        for _, c := range pwStaticOamClasses.PwStaticOamClass {
            if pwStaticOamClasses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PseudowireConfig_PwStaticOamClasses_PwStaticOamClass{}
        pwStaticOamClasses.PwStaticOamClass = append(pwStaticOamClasses.PwStaticOamClass, child)
        return &pwStaticOamClasses.PwStaticOamClass[len(pwStaticOamClasses.PwStaticOamClass)-1]
    }
    return nil
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pwStaticOamClasses.PwStaticOamClass {
        children[pwStaticOamClasses.PwStaticOamClass[i].GetSegmentPath()] = &pwStaticOamClasses.PwStaticOamClass[i]
    }
    return children
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetBundleName() string { return "cisco_ios_xe" }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetYangName() string { return "pw-static-oam-classes" }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) SetParent(parent types.Entity) { pwStaticOamClasses.parent = parent }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetParent() types.Entity { return pwStaticOamClasses.parent }

func (pwStaticOamClasses *PseudowireConfig_PwStaticOamClasses) GetParentYangName() string { return "pseudowire-config" }

// PseudowireConfig_PwStaticOamClasses_PwStaticOamClass
// Pseudowire static oam class configuration
type PseudowireConfig_PwStaticOamClasses_PwStaticOamClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetFilter() yfilter.YFilter { return pwStaticOamClass.YFilter }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) SetFilter(yf yfilter.YFilter) { pwStaticOamClass.YFilter = yf }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "ack" { return "Ack" }
    if yname == "keepalive" { return "Keepalive" }
    if yname == "timeout-refresh-send" { return "TimeoutRefreshSend" }
    if yname == "timeout-refresh-ack" { return "TimeoutRefreshAck" }
    return ""
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetSegmentPath() string {
    return "pw-static-oam-class" + "[name='" + fmt.Sprintf("%v", pwStaticOamClass.Name) + "']"
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = pwStaticOamClass.Name
    leafs["ack"] = pwStaticOamClass.Ack
    leafs["keepalive"] = pwStaticOamClass.Keepalive
    leafs["timeout-refresh-send"] = pwStaticOamClass.TimeoutRefreshSend
    leafs["timeout-refresh-ack"] = pwStaticOamClass.TimeoutRefreshAck
    return leafs
}

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetBundleName() string { return "cisco_ios_xe" }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetYangName() string { return "pw-static-oam-class" }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) SetParent(parent types.Entity) { pwStaticOamClass.parent = parent }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetParent() types.Entity { return pwStaticOamClass.parent }

func (pwStaticOamClass *PseudowireConfig_PwStaticOamClasses_PwStaticOamClass) GetParentYangName() string { return "pw-static-oam-classes" }

// PseudowireState
// Contains the operational state for all the pseudowire
// instances in the switch, no matter what L2VPN service
// created them.
type PseudowireState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each list element contains the state for a pseudowire instance. The list
    // can be optionally keyed by any combination of vc-id, peer-address, etc.
    // Additional filtering of the list by the agent may be performed upon request
    // by the client using subtree filtering as described in RFC 6020 Section 6.
    // The type is slice of PseudowireState_Pseudowires.
    Pseudowires []PseudowireState_Pseudowires
}

func (pseudowireState *PseudowireState) GetFilter() yfilter.YFilter { return pseudowireState.YFilter }

func (pseudowireState *PseudowireState) SetFilter(yf yfilter.YFilter) { pseudowireState.YFilter = yf }

func (pseudowireState *PseudowireState) GetGoName(yname string) string {
    if yname == "pseudowires" { return "Pseudowires" }
    return ""
}

func (pseudowireState *PseudowireState) GetSegmentPath() string {
    return "cisco-pw:pseudowire-state"
}

func (pseudowireState *PseudowireState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pseudowires" {
        for _, c := range pseudowireState.Pseudowires {
            if pseudowireState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PseudowireState_Pseudowires{}
        pseudowireState.Pseudowires = append(pseudowireState.Pseudowires, child)
        return &pseudowireState.Pseudowires[len(pseudowireState.Pseudowires)-1]
    }
    return nil
}

func (pseudowireState *PseudowireState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pseudowireState.Pseudowires {
        children[pseudowireState.Pseudowires[i].GetSegmentPath()] = &pseudowireState.Pseudowires[i]
    }
    return children
}

func (pseudowireState *PseudowireState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pseudowireState *PseudowireState) GetBundleName() string { return "cisco_ios_xe" }

func (pseudowireState *PseudowireState) GetYangName() string { return "pseudowire-state" }

func (pseudowireState *PseudowireState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pseudowireState *PseudowireState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pseudowireState *PseudowireState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pseudowireState *PseudowireState) SetParent(parent types.Entity) { pseudowireState.parent = parent }

func (pseudowireState *PseudowireState) GetParent() types.Entity { return pseudowireState.parent }

func (pseudowireState *PseudowireState) GetParentYangName() string { return "cisco-pw" }

// PseudowireState_Pseudowires
// Each list element contains the state for a pseudowire
// instance. The list can be optionally keyed by any
// combination of vc-id, peer-address, etc.
// Additional filtering of the list by the agent may be
// performed upon request by the client using subtree
// filtering as described in RFC 6020 Section 6.
type PseudowireState_Pseudowires struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This object contains the value of the peer node
    // address of the PW/PE protocol entity. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // following: PwVcTypeEtherPwVcTypeVlanPassthroughPwVcTypeVlan.
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

func (pseudowires *PseudowireState_Pseudowires) GetFilter() yfilter.YFilter { return pseudowires.YFilter }

func (pseudowires *PseudowireState_Pseudowires) SetFilter(yf yfilter.YFilter) { pseudowires.YFilter = yf }

func (pseudowires *PseudowireState_Pseudowires) GetGoName(yname string) string {
    if yname == "vc-peer-address" { return "VcPeerAddress" }
    if yname == "vc-id" { return "VcId" }
    if yname == "vc-owner-type" { return "VcOwnerType" }
    if yname == "vc-name" { return "VcName" }
    if yname == "vc-index" { return "VcIndex" }
    if yname == "vc-type" { return "VcType" }
    if yname == "vc-owner-name" { return "VcOwnerName" }
    if yname == "vc-psn-type" { return "VcPsnType" }
    if yname == "vc-local-group-id" { return "VcLocalGroupId" }
    if yname == "vc-control-word" { return "VcControlWord" }
    if yname == "vc-local-if-mtu" { return "VcLocalIfMtu" }
    if yname == "vc-remote-group-id" { return "VcRemoteGroupId" }
    if yname == "vc-remote-control-word" { return "VcRemoteControlWord" }
    if yname == "vc-remote-if-mtu" { return "VcRemoteIfMtu" }
    if yname == "vc-outbound-label" { return "VcOutboundLabel" }
    if yname == "vc-inbound-label" { return "VcInboundLabel" }
    if yname == "vc-oper-status" { return "VcOperStatus" }
    if yname == "vc-inbound-oper-status" { return "VcInboundOperStatus" }
    if yname == "vc-outbound-oper-status" { return "VcOutboundOperStatus" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (pseudowires *PseudowireState_Pseudowires) GetSegmentPath() string {
    return "pseudowires" + "[vc-peer-address='" + fmt.Sprintf("%v", pseudowires.VcPeerAddress) + "']" + "[vc-id='" + fmt.Sprintf("%v", pseudowires.VcId) + "']" + "[vc-owner-type='" + fmt.Sprintf("%v", pseudowires.VcOwnerType) + "']" + "[vc-name='" + fmt.Sprintf("%v", pseudowires.VcName) + "']" + "[vc-index='" + fmt.Sprintf("%v", pseudowires.VcIndex) + "']"
}

func (pseudowires *PseudowireState_Pseudowires) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &pseudowires.Statistics
    }
    return nil
}

func (pseudowires *PseudowireState_Pseudowires) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &pseudowires.Statistics
    return children
}

func (pseudowires *PseudowireState_Pseudowires) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vc-peer-address"] = pseudowires.VcPeerAddress
    leafs["vc-id"] = pseudowires.VcId
    leafs["vc-owner-type"] = pseudowires.VcOwnerType
    leafs["vc-name"] = pseudowires.VcName
    leafs["vc-index"] = pseudowires.VcIndex
    leafs["vc-type"] = pseudowires.VcType
    leafs["vc-owner-name"] = pseudowires.VcOwnerName
    leafs["vc-psn-type"] = pseudowires.VcPsnType
    leafs["vc-local-group-id"] = pseudowires.VcLocalGroupId
    leafs["vc-control-word"] = pseudowires.VcControlWord
    leafs["vc-local-if-mtu"] = pseudowires.VcLocalIfMtu
    leafs["vc-remote-group-id"] = pseudowires.VcRemoteGroupId
    leafs["vc-remote-control-word"] = pseudowires.VcRemoteControlWord
    leafs["vc-remote-if-mtu"] = pseudowires.VcRemoteIfMtu
    leafs["vc-outbound-label"] = pseudowires.VcOutboundLabel
    leafs["vc-inbound-label"] = pseudowires.VcInboundLabel
    leafs["vc-oper-status"] = pseudowires.VcOperStatus
    leafs["vc-inbound-oper-status"] = pseudowires.VcInboundOperStatus
    leafs["vc-outbound-oper-status"] = pseudowires.VcOutboundOperStatus
    return leafs
}

func (pseudowires *PseudowireState_Pseudowires) GetBundleName() string { return "cisco_ios_xe" }

func (pseudowires *PseudowireState_Pseudowires) GetYangName() string { return "pseudowires" }

func (pseudowires *PseudowireState_Pseudowires) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pseudowires *PseudowireState_Pseudowires) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pseudowires *PseudowireState_Pseudowires) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pseudowires *PseudowireState_Pseudowires) SetParent(parent types.Entity) { pseudowires.parent = parent }

func (pseudowires *PseudowireState_Pseudowires) GetParent() types.Entity { return pseudowires.parent }

func (pseudowires *PseudowireState_Pseudowires) GetParentYangName() string { return "pseudowire-state" }

// PseudowireState_Pseudowires_Statistics
// A collection of pseudowire-related statistics objects
type PseudowireState_Pseudowires_Statistics struct {
    parent types.Entity
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
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}). This
    // attribute is mandatory.
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

func (statistics *PseudowireState_Pseudowires_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *PseudowireState_Pseudowires_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *PseudowireState_Pseudowires_Statistics) GetGoName(yname string) string {
    if yname == "vc-create-time" { return "VcCreateTime" }
    if yname == "vc-up-time" { return "VcUpTime" }
    if yname == "discontinuity-time" { return "DiscontinuityTime" }
    if yname == "in-octets" { return "InOctets" }
    if yname == "in-pkts" { return "InPkts" }
    if yname == "in-errors" { return "InErrors" }
    if yname == "out-octets" { return "OutOctets" }
    if yname == "out-pkts" { return "OutPkts" }
    if yname == "out-errors" { return "OutErrors" }
    return ""
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vc-create-time"] = statistics.VcCreateTime
    leafs["vc-up-time"] = statistics.VcUpTime
    leafs["discontinuity-time"] = statistics.DiscontinuityTime
    leafs["in-octets"] = statistics.InOctets
    leafs["in-pkts"] = statistics.InPkts
    leafs["in-errors"] = statistics.InErrors
    leafs["out-octets"] = statistics.OutOctets
    leafs["out-pkts"] = statistics.OutPkts
    leafs["out-errors"] = statistics.OutErrors
    return leafs
}

func (statistics *PseudowireState_Pseudowires_Statistics) GetBundleName() string { return "cisco_ios_xe" }

func (statistics *PseudowireState_Pseudowires_Statistics) GetYangName() string { return "statistics" }

func (statistics *PseudowireState_Pseudowires_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (statistics *PseudowireState_Pseudowires_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (statistics *PseudowireState_Pseudowires_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (statistics *PseudowireState_Pseudowires_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *PseudowireState_Pseudowires_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *PseudowireState_Pseudowires_Statistics) GetParentYangName() string { return "pseudowires" }

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

