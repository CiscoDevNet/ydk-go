// This YANG module describes the configuration and operational
// model for bridge domain.
// 
// Terms and Acronyms
//   AC : Attachment Circuits
// 
//   BD : Bridge Domain
// 
//   BCB : Backbone Core Bridge
// 
//   BEB : Backbone Edge Bridge
// 
//   B-MAC : Backbone MAC Address
// 
//   CE : Customer Edge
// 
//   C-MAC : Customer/Client MAC Address
// 
//   DHCP : Dynamic Host Configuration Protocol
// 
//   DAI : Dynamic ARP Inspection
// 
//   EVC : Ethernet Virtual Circuit
// 
//   IGMP : Internet Group Management Protocol
// 
//   IPSG : IP Source Guard
// 
//   L2VPN : Layer 2 Virtual Private Network
// 
//   MLD : Multicast Listener Discovery
// 
//   PBB : Provider Backbone Bridge
// 
//   VLAN : Virtual Local Area Network
// 
package cisco_bridge_domain

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_bridge_domain"))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:cisco-bridge-domain bridge-domain-config}", reflect.TypeOf(BridgeDomainConfig{}))
    ydk.RegisterEntity("cisco-bridge-domain:bridge-domain-config", reflect.TypeOf(BridgeDomainConfig{}))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:cisco-bridge-domain bridge-domain-state}", reflect.TypeOf(BridgeDomainState{}))
    ydk.RegisterEntity("cisco-bridge-domain:bridge-domain-state", reflect.TypeOf(BridgeDomainState{}))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:cisco-bridge-domain clear-bridge-domain}", reflect.TypeOf(ClearBridgeDomain{}))
    ydk.RegisterEntity("cisco-bridge-domain:clear-bridge-domain", reflect.TypeOf(ClearBridgeDomain{}))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:cisco-bridge-domain clear-mac-address}", reflect.TypeOf(ClearMacAddress{}))
    ydk.RegisterEntity("cisco-bridge-domain:clear-mac-address", reflect.TypeOf(ClearMacAddress{}))
    ydk.RegisterEntity("{urn:cisco:params:xml:ns:yang:cisco-bridge-domain create-parameterized-bridge-domains}", reflect.TypeOf(CreateParameterizedBridgeDomains{}))
    ydk.RegisterEntity("cisco-bridge-domain:create-parameterized-bridge-domains", reflect.TypeOf(CreateParameterizedBridgeDomains{}))
}

// BridgeDomainStateType represents Bridge domain states.
type BridgeDomainStateType string

const (
    // Bridge domain is operationally Up.
    BridgeDomainStateType_up BridgeDomainStateType = "up"

    // Bridge domain is operationally Down.
    BridgeDomainStateType_down BridgeDomainStateType = "down"

    // Bridge domain is shutdown by administrator.
    BridgeDomainStateType_admin_down BridgeDomainStateType = "admin-down"
)

// BridgeDomainConfig
// This container defines overall configuration data for bridge
// -domains on a network device.
type BridgeDomainConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global configurations for bridge-domains.
    Global BridgeDomainConfig_Global

    // Collection of bridge-groups.  A Bridge-group is logical grouping construct
    // for bridge domains. It defines grouping of bridge-domains under a named
    // bridge-group. For example all bridge-domains belonging to a single customer
    // can be grouped under a bridge -group.
    BridgeGroups BridgeDomainConfig_BridgeGroups

    // Collection of bridge domains.
    BridgeDomains BridgeDomainConfig_BridgeDomains
}

func (bridgeDomainConfig *BridgeDomainConfig) GetEntityData() *types.CommonEntityData {
    bridgeDomainConfig.EntityData.YFilter = bridgeDomainConfig.YFilter
    bridgeDomainConfig.EntityData.YangName = "bridge-domain-config"
    bridgeDomainConfig.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomainConfig.EntityData.ParentYangName = "cisco-bridge-domain"
    bridgeDomainConfig.EntityData.SegmentPath = "cisco-bridge-domain:bridge-domain-config"
    bridgeDomainConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomainConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomainConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomainConfig.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainConfig.EntityData.Children["global"] = types.YChild{"Global", &bridgeDomainConfig.Global}
    bridgeDomainConfig.EntityData.Children["bridge-groups"] = types.YChild{"BridgeGroups", &bridgeDomainConfig.BridgeGroups}
    bridgeDomainConfig.EntityData.Children["bridge-domains"] = types.YChild{"BridgeDomains", &bridgeDomainConfig.BridgeDomains}
    bridgeDomainConfig.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainConfig.EntityData)
}

// BridgeDomainConfig_Global
// Global configurations for bridge-domains.
type BridgeDomainConfig_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If this leaf is set to true, then it enables the emission of
    // 'bd-state-notification'; otherwise these notifications are not emitted. The
    // type is bool.
    BdStateNotificationEnabled interface{}

    // This leaf defines the maximum number of 'bd-state- notification' that can
    // be emitted from the device per second. The type is interface{} with range:
    // 0..4294967295.
    BdStateNotificationRate interface{}

    // Provider Backbone Bridging (PBB) related global configurations.
    Pbb BridgeDomainConfig_Global_Pbb
}

func (global *BridgeDomainConfig_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xe"
    global.EntityData.ParentYangName = "bridge-domain-config"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Children["pbb"] = types.YChild{"Pbb", &global.Pbb}
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["bd-state-notification-enabled"] = types.YLeaf{"BdStateNotificationEnabled", global.BdStateNotificationEnabled}
    global.EntityData.Leafs["bd-state-notification-rate"] = types.YLeaf{"BdStateNotificationRate", global.BdStateNotificationRate}
    return &(global.EntityData)
}

// BridgeDomainConfig_Global_Pbb
// Provider Backbone Bridging (PBB) related global
// configurations.
type BridgeDomainConfig_Global_Pbb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Backbone source mac address configuration for Provider Backbone Bridging
    // (PBB). The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    BackboneSrcMac interface{}
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetEntityData() *types.CommonEntityData {
    pbb.EntityData.YFilter = pbb.YFilter
    pbb.EntityData.YangName = "pbb"
    pbb.EntityData.BundleName = "cisco_ios_xe"
    pbb.EntityData.ParentYangName = "global"
    pbb.EntityData.SegmentPath = "pbb"
    pbb.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pbb.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pbb.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pbb.EntityData.Children = make(map[string]types.YChild)
    pbb.EntityData.Leafs = make(map[string]types.YLeaf)
    pbb.EntityData.Leafs["backbone-src-mac"] = types.YLeaf{"BackboneSrcMac", pbb.BackboneSrcMac}
    return &(pbb.EntityData)
}

// BridgeDomainConfig_BridgeGroups
// Collection of bridge-groups.
// 
// A Bridge-group is logical grouping construct for bridge
// domains. It defines grouping of bridge-domains under a
// named bridge-group. For example all bridge-domains
// belonging to a single customer can be grouped under a bridge
// -group
type BridgeDomainConfig_BridgeGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge-group configuration. The type is slice of
    // BridgeDomainConfig_BridgeGroups_BridgeGroup.
    BridgeGroup []BridgeDomainConfig_BridgeGroups_BridgeGroup
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetEntityData() *types.CommonEntityData {
    bridgeGroups.EntityData.YFilter = bridgeGroups.YFilter
    bridgeGroups.EntityData.YangName = "bridge-groups"
    bridgeGroups.EntityData.BundleName = "cisco_ios_xe"
    bridgeGroups.EntityData.ParentYangName = "bridge-domain-config"
    bridgeGroups.EntityData.SegmentPath = "bridge-groups"
    bridgeGroups.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeGroups.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeGroups.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeGroups.EntityData.Children = make(map[string]types.YChild)
    bridgeGroups.EntityData.Children["bridge-group"] = types.YChild{"BridgeGroup", nil}
    for i := range bridgeGroups.BridgeGroup {
        bridgeGroups.EntityData.Children[types.GetSegmentPath(&bridgeGroups.BridgeGroup[i])] = types.YChild{"BridgeGroup", &bridgeGroups.BridgeGroup[i]}
    }
    bridgeGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeGroups.EntityData)
}

// BridgeDomainConfig_BridgeGroups_BridgeGroup
// Bridge-group configuration.
type BridgeDomainConfig_BridgeGroups_BridgeGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge-group name. The type is string with length:
    // 1..32.
    Name interface{}
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetEntityData() *types.CommonEntityData {
    bridgeGroup.EntityData.YFilter = bridgeGroup.YFilter
    bridgeGroup.EntityData.YangName = "bridge-group"
    bridgeGroup.EntityData.BundleName = "cisco_ios_xe"
    bridgeGroup.EntityData.ParentYangName = "bridge-groups"
    bridgeGroup.EntityData.SegmentPath = "bridge-group" + "[name='" + fmt.Sprintf("%v", bridgeGroup.Name) + "']"
    bridgeGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeGroup.EntityData.Children = make(map[string]types.YChild)
    bridgeGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeGroup.EntityData.Leafs["name"] = types.YLeaf{"Name", bridgeGroup.Name}
    return &(bridgeGroup.EntityData)
}

// BridgeDomainConfig_BridgeDomains
// Collection of bridge domains.
type BridgeDomainConfig_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Definition of a bridge-domain. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain.
    BridgeDomain []BridgeDomainConfig_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-config"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomains.EntityData.Children = make(map[string]types.YChild)
    bridgeDomains.EntityData.Children["bridge-domain"] = types.YChild{"BridgeDomain", nil}
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children[types.GetSegmentPath(&bridgeDomains.BridgeDomain[i])] = types.YChild{"BridgeDomain", &bridgeDomains.BridgeDomain[i]}
    }
    bridgeDomains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomains.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain
// Definition of a bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain name or number. The type is string.
    Id interface{}

    // Reference to bridge-group name. If bridge-groups are supported, referred
    // bridge-group MUST be created first. The type is string with length: 1..32.
    // Refers to
    // cisco_bridge_domain.BridgeDomainConfig_BridgeGroups_BridgeGroup_Name This
    // attribute is mandatory.
    BridgeGroup interface{}

    // This leaf represents configured admin status of the bridge domain. The type
    // is bool. The default value is true.
    Enabled interface{}

    // Enable/disable bridge-domain status change notification.  If true, any
    // change in bridge-domain operational status will be notified to client via
    // 'bd-status-change' notification. The type is bool.
    BdStatusChangeNotification interface{}

    // The MTU size for bridge domain. All bridge domain members must have same
    // MTU size to be operational in the domain. The type is interface{} with
    // range: 46..65535.
    Mtu interface{}

    // Flood modes for optimization. The type is FloodingMode.
    FloodingMode interface{}

    // Collection of bridge-domain members.
    Members BridgeDomainConfig_BridgeDomains_BridgeDomain_Members

    // MAC features for bridge domain.
    Mac BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac

    // Dynamic ARP Inspection (DAI) configurations.
    DynamicArpInspection BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection

    // IP source guard (IPSG) configurations.
    IpSourceGuard BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard

    // A collection of storm control threshold and action configurations.
    StormControl BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl

    // Enable IGMP snooping.
    IgmpSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping

    // Enable MLD snooping.
    MldSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping

    // Enable DHCP IPv4 snooping.
    DhcpIpv4Snooping BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + "[id='" + fmt.Sprintf("%v", bridgeDomain.Id) + "']"
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = make(map[string]types.YChild)
    bridgeDomain.EntityData.Children["members"] = types.YChild{"Members", &bridgeDomain.Members}
    bridgeDomain.EntityData.Children["mac"] = types.YChild{"Mac", &bridgeDomain.Mac}
    bridgeDomain.EntityData.Children["dynamic-arp-inspection"] = types.YChild{"DynamicArpInspection", &bridgeDomain.DynamicArpInspection}
    bridgeDomain.EntityData.Children["ip-source-guard"] = types.YChild{"IpSourceGuard", &bridgeDomain.IpSourceGuard}
    bridgeDomain.EntityData.Children["storm-control"] = types.YChild{"StormControl", &bridgeDomain.StormControl}
    bridgeDomain.EntityData.Children["igmp-snooping"] = types.YChild{"IgmpSnooping", &bridgeDomain.IgmpSnooping}
    bridgeDomain.EntityData.Children["mld-snooping"] = types.YChild{"MldSnooping", &bridgeDomain.MldSnooping}
    bridgeDomain.EntityData.Children["dhcp-ipv4-snooping"] = types.YChild{"DhcpIpv4Snooping", &bridgeDomain.DhcpIpv4Snooping}
    bridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomain.EntityData.Leafs["id"] = types.YLeaf{"Id", bridgeDomain.Id}
    bridgeDomain.EntityData.Leafs["bridge-group"] = types.YLeaf{"BridgeGroup", bridgeDomain.BridgeGroup}
    bridgeDomain.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", bridgeDomain.Enabled}
    bridgeDomain.EntityData.Leafs["bd-status-change-notification"] = types.YLeaf{"BdStatusChangeNotification", bridgeDomain.BdStatusChangeNotification}
    bridgeDomain.EntityData.Leafs["mtu"] = types.YLeaf{"Mtu", bridgeDomain.Mtu}
    bridgeDomain.EntityData.Leafs["flooding-mode"] = types.YLeaf{"FloodingMode", bridgeDomain.FloodingMode}
    return &(bridgeDomain.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Attachment circuits for current bridge-domain. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember.
    AcMember []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember

    // List of Virtual Forrwarding Interfaces for current bridge-domain. The type
    // is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember.
    VfiMember []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember

    // Collection of access pseudowire members.  A Pseudowires can be a regular
    // interface with ifType 'ifPwType' or it can represented as a non-interface
    // context. This container provides model definition for both types of the
    // pseudowires.
    AccessPwMember BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "cisco_ios_xe"
    members.EntityData.ParentYangName = "bridge-domain"
    members.EntityData.SegmentPath = "members"
    members.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    members.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    members.EntityData.Children = make(map[string]types.YChild)
    members.EntityData.Children["ac-member"] = types.YChild{"AcMember", nil}
    for i := range members.AcMember {
        members.EntityData.Children[types.GetSegmentPath(&members.AcMember[i])] = types.YChild{"AcMember", &members.AcMember[i]}
    }
    members.EntityData.Children["vfi-member"] = types.YChild{"VfiMember", nil}
    for i := range members.VfiMember {
        members.EntityData.Children[types.GetSegmentPath(&members.VfiMember[i])] = types.YChild{"VfiMember", &members.VfiMember[i]}
    }
    members.EntityData.Children["access-pw-member"] = types.YChild{"AccessPwMember", &members.AccessPwMember}
    members.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(members.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember
// List of Attachment circuits for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an attchment circuit interface
    // instance which is configured to be part of this bridge-domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Bridge domain aggregates attachment circuits (ACs) and pseudowires (PWs) in
    // one or more groups called Split Horizon Groups. When applied to bridge
    // domains, Split Horizon refers to the flooding and forwarding behavior
    // between members of a Split Horizon group. In general, frames received on
    // one member of a split horizon group are not flooded out to the other
    // members.
    SplitHorizonGroup BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup

    // MAC features for bridge domain.
    Mac BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac

    // Enable IGMP snooping.
    IgmpSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping

    // Enable MLD snooping.
    MldSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping

    // Enable DHCP IPv4 snooping.
    DhcpIpv4Snooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping

    // Flooding configurations.
    Flooding BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding

    // A collection of storm control threshold and action configurations.
    StormControl BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl

    // Dynamic ARP Inspection (DAI) configurations.
    DynamicArpInspection BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection

    // IP source guard (IPSG) configurations.
    IpSourceGuard BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetEntityData() *types.CommonEntityData {
    acMember.EntityData.YFilter = acMember.YFilter
    acMember.EntityData.YangName = "ac-member"
    acMember.EntityData.BundleName = "cisco_ios_xe"
    acMember.EntityData.ParentYangName = "members"
    acMember.EntityData.SegmentPath = "ac-member" + "[interface='" + fmt.Sprintf("%v", acMember.Interface_) + "']"
    acMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    acMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    acMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    acMember.EntityData.Children = make(map[string]types.YChild)
    acMember.EntityData.Children["split-horizon-group"] = types.YChild{"SplitHorizonGroup", &acMember.SplitHorizonGroup}
    acMember.EntityData.Children["mac"] = types.YChild{"Mac", &acMember.Mac}
    acMember.EntityData.Children["igmp-snooping"] = types.YChild{"IgmpSnooping", &acMember.IgmpSnooping}
    acMember.EntityData.Children["mld-snooping"] = types.YChild{"MldSnooping", &acMember.MldSnooping}
    acMember.EntityData.Children["dhcp-ipv4-snooping"] = types.YChild{"DhcpIpv4Snooping", &acMember.DhcpIpv4Snooping}
    acMember.EntityData.Children["flooding"] = types.YChild{"Flooding", &acMember.Flooding}
    acMember.EntityData.Children["storm-control"] = types.YChild{"StormControl", &acMember.StormControl}
    acMember.EntityData.Children["dynamic-arp-inspection"] = types.YChild{"DynamicArpInspection", &acMember.DynamicArpInspection}
    acMember.EntityData.Children["ip-source-guard"] = types.YChild{"IpSourceGuard", &acMember.IpSourceGuard}
    acMember.EntityData.Leafs = make(map[string]types.YLeaf)
    acMember.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", acMember.Interface_}
    return &(acMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup
// Bridge domain aggregates attachment circuits (ACs) and
// pseudowires (PWs) in one or more groups called Split Horizon
// Groups. When applied to bridge domains, Split Horizon refers
// to the flooding and forwarding behavior between members of a
// Split Horizon group. In general, frames received on one
// member of a split horizon group are not flooded out to the
// other members.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon group number for bridge domain member. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Id interface{}
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    splitHorizonGroup.EntityData.YFilter = splitHorizonGroup.YFilter
    splitHorizonGroup.EntityData.YangName = "split-horizon-group"
    splitHorizonGroup.EntityData.BundleName = "cisco_ios_xe"
    splitHorizonGroup.EntityData.ParentYangName = "ac-member"
    splitHorizonGroup.EntityData.SegmentPath = "split-horizon-group"
    splitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    splitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    splitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    splitHorizonGroup.EntityData.Children = make(map[string]types.YChild)
    splitHorizonGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    splitHorizonGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", splitHorizonGroup.Id}
    return &(splitHorizonGroup.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable disable mac learning. The type is bool. The default value is true.
    LearningEnabled interface{}

    // MAC table learning limit.
    Limit BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit

    // MAC aging configurations.
    Aging BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging

    // Port down event.
    PortDown BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown

    // MAC secure parameters.
    Secure BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xe"
    mac.EntityData.ParentYangName = "ac-member"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["limit"] = types.YChild{"Limit", &mac.Limit}
    mac.EntityData.Children["aging"] = types.YChild{"Aging", &mac.Aging}
    mac.EntityData.Children["port-down"] = types.YChild{"PortDown", &mac.PortDown}
    mac.EntityData.Children["secure"] = types.YChild{"Secure", &mac.Secure}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["learning-enabled"] = types.YLeaf{"LearningEnabled", mac.LearningEnabled}
    return &(mac.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifNoneNotifSnmpTrapNotifSyslogNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetEntityData() *types.CommonEntityData {
    limit.EntityData.YFilter = limit.YFilter
    limit.EntityData.YangName = "limit"
    limit.EntityData.BundleName = "cisco_ios_xe"
    limit.EntityData.ParentYangName = "mac"
    limit.EntityData.SegmentPath = "limit"
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = make(map[string]types.YChild)
    limit.EntityData.Leafs = make(map[string]types.YLeaf)
    limit.EntityData.Leafs["maximum"] = types.YLeaf{"Maximum", limit.Maximum}
    limit.EntityData.Leafs["action"] = types.YLeaf{"Action", limit.Action}
    limit.EntityData.Leafs["notification"] = types.YLeaf{"Notification", limit.Notification}
    return &(limit.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type_ interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = make(map[string]types.YChild)
    aging.EntityData.Leafs = make(map[string]types.YLeaf)
    aging.EntityData.Leafs["time"] = types.YLeaf{"Time", aging.Time}
    aging.EntityData.Leafs["type"] = types.YLeaf{"Type_", aging.Type_}
    return &(aging.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetEntityData() *types.CommonEntityData {
    portDown.EntityData.YFilter = portDown.YFilter
    portDown.EntityData.YangName = "port-down"
    portDown.EntityData.BundleName = "cisco_ios_xe"
    portDown.EntityData.ParentYangName = "mac"
    portDown.EntityData.SegmentPath = "port-down"
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = make(map[string]types.YChild)
    portDown.EntityData.Leafs = make(map[string]types.YLeaf)
    portDown.EntityData.Leafs["flush"] = types.YLeaf{"Flush", portDown.Flush}
    return &(portDown.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure
// MAC secure parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable mac secure feature. The type is bool.
    Enabled interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetEntityData() *types.CommonEntityData {
    secure.EntityData.YFilter = secure.YFilter
    secure.EntityData.YangName = "secure"
    secure.EntityData.BundleName = "cisco_ios_xe"
    secure.EntityData.ParentYangName = "mac"
    secure.EntityData.SegmentPath = "secure"
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = make(map[string]types.YChild)
    secure.EntityData.Leafs = make(map[string]types.YLeaf)
    secure.EntityData.Leafs["action"] = types.YLeaf{"Action", secure.Action}
    secure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", secure.Logging}
    secure.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", secure.Enabled}
    return &(secure.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping
// Enable IGMP snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetEntityData() *types.CommonEntityData {
    igmpSnooping.EntityData.YFilter = igmpSnooping.YFilter
    igmpSnooping.EntityData.YangName = "igmp-snooping"
    igmpSnooping.EntityData.BundleName = "cisco_ios_xe"
    igmpSnooping.EntityData.ParentYangName = "ac-member"
    igmpSnooping.EntityData.SegmentPath = "igmp-snooping"
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = make(map[string]types.YChild)
    igmpSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    igmpSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", igmpSnooping.ProfileName}
    return &(igmpSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetEntityData() *types.CommonEntityData {
    mldSnooping.EntityData.YFilter = mldSnooping.YFilter
    mldSnooping.EntityData.YangName = "mld-snooping"
    mldSnooping.EntityData.BundleName = "cisco_ios_xe"
    mldSnooping.EntityData.ParentYangName = "ac-member"
    mldSnooping.EntityData.SegmentPath = "mld-snooping"
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = make(map[string]types.YChild)
    mldSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    mldSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", mldSnooping.ProfileName}
    return &(mldSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetEntityData() *types.CommonEntityData {
    dhcpIpv4Snooping.EntityData.YFilter = dhcpIpv4Snooping.YFilter
    dhcpIpv4Snooping.EntityData.YangName = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.BundleName = "cisco_ios_xe"
    dhcpIpv4Snooping.EntityData.ParentYangName = "ac-member"
    dhcpIpv4Snooping.EntityData.SegmentPath = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = make(map[string]types.YChild)
    dhcpIpv4Snooping.EntityData.Leafs = make(map[string]types.YLeaf)
    dhcpIpv4Snooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName}
    return &(dhcpIpv4Snooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "ac-member"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = make(map[string]types.YChild)
    flooding.EntityData.Leafs = make(map[string]types.YLeaf)
    flooding.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", flooding.Disabled}
    flooding.EntityData.Leafs["disabled-unknown-unicast"] = types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast}
    return &(flooding.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionDropActionSnmpTrapActionShutdown.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "ac-member"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = make(map[string]types.YChild)
    stormControl.EntityData.Children["thresholds"] = types.YChild{"Thresholds", nil}
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children[types.GetSegmentPath(&stormControl.Thresholds[i])] = types.YChild{"Thresholds", &stormControl.Thresholds[i]}
    }
    stormControl.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControl.EntityData.Leafs["action"] = types.YLeaf{"Action", stormControl.Action}
    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type for
    // which an administrator desires to configure storm control. The type is
    // EthTrafficClass.
    TrafficClass interface{}

    // Traffic threshold value. Unit of the value is indicated by leaf 'unit'. The
    // type is interface{} with range: 0..4294967295. This attribute is mandatory.
    Value interface{}

    // This enumeration define unit of the traffic threshold value. The type is
    // Unit. This attribute is mandatory.
    Unit interface{}
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xe"
    thresholds.EntityData.ParentYangName = "storm-control"
    thresholds.EntityData.SegmentPath = "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = make(map[string]types.YChild)
    thresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholds.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", thresholds.TrafficClass}
    thresholds.EntityData.Leafs["value"] = types.YLeaf{"Value", thresholds.Value}
    thresholds.EntityData.Leafs["unit"] = types.YLeaf{"Unit", thresholds.Unit}
    return &(thresholds.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit represents value.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit string

const (
    // Bits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit_bps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit = "bps"

    // Kilobits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit_kbps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit = "kbps"

    // Packets per seconds.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit_pps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds_Unit = "pps"
)

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection
// Dynamic ARP Inspection (DAI) configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable DAI logging. The type is bool.
    Logging interface{}

    // Enable or disable Dynamic ARP Inspection. The type is bool.
    Enable interface{}

    // Enable address validation.
    AddressValidation BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetEntityData() *types.CommonEntityData {
    dynamicArpInspection.EntityData.YFilter = dynamicArpInspection.YFilter
    dynamicArpInspection.EntityData.YangName = "dynamic-arp-inspection"
    dynamicArpInspection.EntityData.BundleName = "cisco_ios_xe"
    dynamicArpInspection.EntityData.ParentYangName = "ac-member"
    dynamicArpInspection.EntityData.SegmentPath = "dynamic-arp-inspection"
    dynamicArpInspection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dynamicArpInspection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dynamicArpInspection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dynamicArpInspection.EntityData.Children = make(map[string]types.YChild)
    dynamicArpInspection.EntityData.Children["address-validation"] = types.YChild{"AddressValidation", &dynamicArpInspection.AddressValidation}
    dynamicArpInspection.EntityData.Leafs = make(map[string]types.YLeaf)
    dynamicArpInspection.EntityData.Leafs["logging"] = types.YLeaf{"Logging", dynamicArpInspection.Logging}
    dynamicArpInspection.EntityData.Leafs["enable"] = types.YLeaf{"Enable", dynamicArpInspection.Enable}
    return &(dynamicArpInspection.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match Destination MAC Address. The type is interface{}.
    DstMac interface{}

    // Match Source MAC Address. The type is interface{}.
    SrcMac interface{}

    // Match IPv4 Address. The type is interface{}.
    Ipv4 interface{}
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetEntityData() *types.CommonEntityData {
    addressValidation.EntityData.YFilter = addressValidation.YFilter
    addressValidation.EntityData.YangName = "address-validation"
    addressValidation.EntityData.BundleName = "cisco_ios_xe"
    addressValidation.EntityData.ParentYangName = "dynamic-arp-inspection"
    addressValidation.EntityData.SegmentPath = "address-validation"
    addressValidation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressValidation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressValidation.EntityData.Children = make(map[string]types.YChild)
    addressValidation.EntityData.Leafs = make(map[string]types.YLeaf)
    addressValidation.EntityData.Leafs["dst-mac"] = types.YLeaf{"DstMac", addressValidation.DstMac}
    addressValidation.EntityData.Leafs["src-mac"] = types.YLeaf{"SrcMac", addressValidation.SrcMac}
    addressValidation.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", addressValidation.Ipv4}
    return &(addressValidation.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard
// IP source guard (IPSG) configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPSG logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable IP source guard feature. The type is bool.
    Enable interface{}
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xe"
    ipSourceGuard.EntityData.ParentYangName = "ac-member"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSourceGuard.EntityData.Children = make(map[string]types.YChild)
    ipSourceGuard.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSourceGuard.EntityData.Leafs["logging"] = types.YLeaf{"Logging", ipSourceGuard.Logging}
    ipSourceGuard.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ipSourceGuard.Enable}
    return &(ipSourceGuard.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember
// List of Virtual Forrwarding Interfaces for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an Virtual Forwarding Interface
    // instance which is configured to be part of this bridge-domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface_ interface{}
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetEntityData() *types.CommonEntityData {
    vfiMember.EntityData.YFilter = vfiMember.YFilter
    vfiMember.EntityData.YangName = "vfi-member"
    vfiMember.EntityData.BundleName = "cisco_ios_xe"
    vfiMember.EntityData.ParentYangName = "members"
    vfiMember.EntityData.SegmentPath = "vfi-member" + "[interface='" + fmt.Sprintf("%v", vfiMember.Interface_) + "']"
    vfiMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vfiMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vfiMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vfiMember.EntityData.Children = make(map[string]types.YChild)
    vfiMember.EntityData.Leafs = make(map[string]types.YLeaf)
    vfiMember.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", vfiMember.Interface_}
    return &(vfiMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember
// Collection of access pseudowire members.
// 
// A Pseudowires can be a regular interface with ifType
// 'ifPwType' or it can represented as a non-interface
// context. This container provides model definition for
// both types of the pseudowires.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of interface based access pseudowires for current bridge-domain. The
    // type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember.
    AccessPwIfMember []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember

    // Collection of neighbor specification based pseudo-wires. The type is slice
    // of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec.
    PwNeighborSpec []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetEntityData() *types.CommonEntityData {
    accessPwMember.EntityData.YFilter = accessPwMember.YFilter
    accessPwMember.EntityData.YangName = "access-pw-member"
    accessPwMember.EntityData.BundleName = "cisco_ios_xe"
    accessPwMember.EntityData.ParentYangName = "members"
    accessPwMember.EntityData.SegmentPath = "access-pw-member"
    accessPwMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwMember.EntityData.Children = make(map[string]types.YChild)
    accessPwMember.EntityData.Children["access-pw-if-member"] = types.YChild{"AccessPwIfMember", nil}
    for i := range accessPwMember.AccessPwIfMember {
        accessPwMember.EntityData.Children[types.GetSegmentPath(&accessPwMember.AccessPwIfMember[i])] = types.YChild{"AccessPwIfMember", &accessPwMember.AccessPwIfMember[i]}
    }
    accessPwMember.EntityData.Children["pw-neighbor-spec"] = types.YChild{"PwNeighborSpec", nil}
    for i := range accessPwMember.PwNeighborSpec {
        accessPwMember.EntityData.Children[types.GetSegmentPath(&accessPwMember.PwNeighborSpec[i])] = types.YChild{"PwNeighborSpec", &accessPwMember.PwNeighborSpec[i]}
    }
    accessPwMember.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessPwMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember
// List of interface based access pseudowires for
// current bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an access pseudo-wire interface
    // instance which is configured to be part of this bridge domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface_ interface{}
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetEntityData() *types.CommonEntityData {
    accessPwIfMember.EntityData.YFilter = accessPwIfMember.YFilter
    accessPwIfMember.EntityData.YangName = "access-pw-if-member"
    accessPwIfMember.EntityData.BundleName = "cisco_ios_xe"
    accessPwIfMember.EntityData.ParentYangName = "access-pw-member"
    accessPwIfMember.EntityData.SegmentPath = "access-pw-if-member" + "[interface='" + fmt.Sprintf("%v", accessPwIfMember.Interface_) + "']"
    accessPwIfMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwIfMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwIfMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwIfMember.EntityData.Children = make(map[string]types.YChild)
    accessPwIfMember.EntityData.Leafs = make(map[string]types.YLeaf)
    accessPwIfMember.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", accessPwIfMember.Interface_}
    return &(accessPwIfMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec
// Collection of neighbor specification based
// pseudo-wires.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 address of the neighbor. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborIpAddress interface{}

    // This attribute is a key. Pseudowire VC ID. The type is interface{} with
    // range: 1..4294967295.
    VcId interface{}

    // Reference to a pseudowire template. The type is string. Refers to
    // cisco_pw.PseudowireConfig_PwTemplates_PwTemplate_Name
    PwClassTemplate interface{}

    // Encapsulation configuration for this neighbor. The type is one of the
    // following: PwEncapMpls.
    EncapType interface{}

    // Specify a tag for a VLAN ID for the pseudowire. The type is interface{}
    // with range: 1..4094.
    TagImposeVlan interface{}

    // The local source IPv6 address. Note this should only be configured when
    // neighbor address is IPv6 type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceIpv6 interface{}

    // Statically configured labels, signalling should be none.
    StaticLabel BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel

    // Bridge domain aggregates attachment circuits (ACs) and pseudowires (PWs) in
    // one or more groups called Split Horizon Groups. When applied to bridge
    // domains, Split Horizon refers to the flooding and forwarding behavior
    // between members of a Split Horizon group. In general, frames received on
    // one member of a split horizon group are not flooded out to the other
    // members.
    SplitHorizonGroup BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup

    // MAC features for bridge domain.
    Mac BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac

    // Enable IGMP snooping.
    IgmpSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping

    // Enable MLD snooping.
    MldSnooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping

    // Enable DHCP IPv4 snooping.
    DhcpIpv4Snooping BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping

    // Flooding configurations.
    Flooding BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding

    // A collection of storm control threshold and action configurations.
    StormControl BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl

    // Backup pseudo-wire.
    Backup BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetEntityData() *types.CommonEntityData {
    pwNeighborSpec.EntityData.YFilter = pwNeighborSpec.YFilter
    pwNeighborSpec.EntityData.YangName = "pw-neighbor-spec"
    pwNeighborSpec.EntityData.BundleName = "cisco_ios_xe"
    pwNeighborSpec.EntityData.ParentYangName = "access-pw-member"
    pwNeighborSpec.EntityData.SegmentPath = "pw-neighbor-spec" + "[neighbor-ip-address='" + fmt.Sprintf("%v", pwNeighborSpec.NeighborIpAddress) + "']" + "[vc-id='" + fmt.Sprintf("%v", pwNeighborSpec.VcId) + "']"
    pwNeighborSpec.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwNeighborSpec.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwNeighborSpec.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwNeighborSpec.EntityData.Children = make(map[string]types.YChild)
    pwNeighborSpec.EntityData.Children["static-label"] = types.YChild{"StaticLabel", &pwNeighborSpec.StaticLabel}
    pwNeighborSpec.EntityData.Children["split-horizon-group"] = types.YChild{"SplitHorizonGroup", &pwNeighborSpec.SplitHorizonGroup}
    pwNeighborSpec.EntityData.Children["mac"] = types.YChild{"Mac", &pwNeighborSpec.Mac}
    pwNeighborSpec.EntityData.Children["igmp-snooping"] = types.YChild{"IgmpSnooping", &pwNeighborSpec.IgmpSnooping}
    pwNeighborSpec.EntityData.Children["mld-snooping"] = types.YChild{"MldSnooping", &pwNeighborSpec.MldSnooping}
    pwNeighborSpec.EntityData.Children["dhcp-ipv4-snooping"] = types.YChild{"DhcpIpv4Snooping", &pwNeighborSpec.DhcpIpv4Snooping}
    pwNeighborSpec.EntityData.Children["flooding"] = types.YChild{"Flooding", &pwNeighborSpec.Flooding}
    pwNeighborSpec.EntityData.Children["storm-control"] = types.YChild{"StormControl", &pwNeighborSpec.StormControl}
    pwNeighborSpec.EntityData.Children["backup"] = types.YChild{"Backup", &pwNeighborSpec.Backup}
    pwNeighborSpec.EntityData.Leafs = make(map[string]types.YLeaf)
    pwNeighborSpec.EntityData.Leafs["neighbor-ip-address"] = types.YLeaf{"NeighborIpAddress", pwNeighborSpec.NeighborIpAddress}
    pwNeighborSpec.EntityData.Leafs["vc-id"] = types.YLeaf{"VcId", pwNeighborSpec.VcId}
    pwNeighborSpec.EntityData.Leafs["pw-class-template"] = types.YLeaf{"PwClassTemplate", pwNeighborSpec.PwClassTemplate}
    pwNeighborSpec.EntityData.Leafs["encap-type"] = types.YLeaf{"EncapType", pwNeighborSpec.EncapType}
    pwNeighborSpec.EntityData.Leafs["tag-impose-vlan"] = types.YLeaf{"TagImposeVlan", pwNeighborSpec.TagImposeVlan}
    pwNeighborSpec.EntityData.Leafs["source-ipv6"] = types.YLeaf{"SourceIpv6", pwNeighborSpec.SourceIpv6}
    return &(pwNeighborSpec.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel
// Statically configured labels, signalling should be none
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local MPLS label ID. The type is interface{} with range: 16..1048575.
    LocalLabel interface{}

    // Remote MPLS label ID. The type is interface{} with range: 16..1048575.
    RemoteLabel interface{}
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetEntityData() *types.CommonEntityData {
    staticLabel.EntityData.YFilter = staticLabel.YFilter
    staticLabel.EntityData.YangName = "static-label"
    staticLabel.EntityData.BundleName = "cisco_ios_xe"
    staticLabel.EntityData.ParentYangName = "pw-neighbor-spec"
    staticLabel.EntityData.SegmentPath = "static-label"
    staticLabel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    staticLabel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    staticLabel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    staticLabel.EntityData.Children = make(map[string]types.YChild)
    staticLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    staticLabel.EntityData.Leafs["local-label"] = types.YLeaf{"LocalLabel", staticLabel.LocalLabel}
    staticLabel.EntityData.Leafs["remote-label"] = types.YLeaf{"RemoteLabel", staticLabel.RemoteLabel}
    return &(staticLabel.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup
// Bridge domain aggregates attachment circuits (ACs) and
// pseudowires (PWs) in one or more groups called Split Horizon
// Groups. When applied to bridge domains, Split Horizon refers
// to the flooding and forwarding behavior between members of a
// Split Horizon group. In general, frames received on one
// member of a split horizon group are not flooded out to the
// other members.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Split Horizon group number for bridge domain member. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Id interface{}
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetEntityData() *types.CommonEntityData {
    splitHorizonGroup.EntityData.YFilter = splitHorizonGroup.YFilter
    splitHorizonGroup.EntityData.YangName = "split-horizon-group"
    splitHorizonGroup.EntityData.BundleName = "cisco_ios_xe"
    splitHorizonGroup.EntityData.ParentYangName = "pw-neighbor-spec"
    splitHorizonGroup.EntityData.SegmentPath = "split-horizon-group"
    splitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    splitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    splitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    splitHorizonGroup.EntityData.Children = make(map[string]types.YChild)
    splitHorizonGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    splitHorizonGroup.EntityData.Leafs["id"] = types.YLeaf{"Id", splitHorizonGroup.Id}
    return &(splitHorizonGroup.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable disable mac learning. The type is bool. The default value is true.
    LearningEnabled interface{}

    // MAC table learning limit.
    Limit BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit

    // MAC aging configurations.
    Aging BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging

    // Port down event.
    PortDown BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown

    // MAC secure parameters.
    Secure BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xe"
    mac.EntityData.ParentYangName = "pw-neighbor-spec"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["limit"] = types.YChild{"Limit", &mac.Limit}
    mac.EntityData.Children["aging"] = types.YChild{"Aging", &mac.Aging}
    mac.EntityData.Children["port-down"] = types.YChild{"PortDown", &mac.PortDown}
    mac.EntityData.Children["secure"] = types.YChild{"Secure", &mac.Secure}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["learning-enabled"] = types.YLeaf{"LearningEnabled", mac.LearningEnabled}
    return &(mac.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifNoneNotifSnmpTrapNotifSyslogNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetEntityData() *types.CommonEntityData {
    limit.EntityData.YFilter = limit.YFilter
    limit.EntityData.YangName = "limit"
    limit.EntityData.BundleName = "cisco_ios_xe"
    limit.EntityData.ParentYangName = "mac"
    limit.EntityData.SegmentPath = "limit"
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = make(map[string]types.YChild)
    limit.EntityData.Leafs = make(map[string]types.YLeaf)
    limit.EntityData.Leafs["maximum"] = types.YLeaf{"Maximum", limit.Maximum}
    limit.EntityData.Leafs["action"] = types.YLeaf{"Action", limit.Action}
    limit.EntityData.Leafs["notification"] = types.YLeaf{"Notification", limit.Notification}
    return &(limit.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type_ interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = make(map[string]types.YChild)
    aging.EntityData.Leafs = make(map[string]types.YLeaf)
    aging.EntityData.Leafs["time"] = types.YLeaf{"Time", aging.Time}
    aging.EntityData.Leafs["type"] = types.YLeaf{"Type_", aging.Type_}
    return &(aging.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetEntityData() *types.CommonEntityData {
    portDown.EntityData.YFilter = portDown.YFilter
    portDown.EntityData.YangName = "port-down"
    portDown.EntityData.BundleName = "cisco_ios_xe"
    portDown.EntityData.ParentYangName = "mac"
    portDown.EntityData.SegmentPath = "port-down"
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = make(map[string]types.YChild)
    portDown.EntityData.Leafs = make(map[string]types.YLeaf)
    portDown.EntityData.Leafs["flush"] = types.YLeaf{"Flush", portDown.Flush}
    return &(portDown.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure
// MAC secure parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable mac secure feature. The type is bool.
    Enabled interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetEntityData() *types.CommonEntityData {
    secure.EntityData.YFilter = secure.YFilter
    secure.EntityData.YangName = "secure"
    secure.EntityData.BundleName = "cisco_ios_xe"
    secure.EntityData.ParentYangName = "mac"
    secure.EntityData.SegmentPath = "secure"
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = make(map[string]types.YChild)
    secure.EntityData.Leafs = make(map[string]types.YLeaf)
    secure.EntityData.Leafs["action"] = types.YLeaf{"Action", secure.Action}
    secure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", secure.Logging}
    secure.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", secure.Enabled}
    return &(secure.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping
// Enable IGMP snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetEntityData() *types.CommonEntityData {
    igmpSnooping.EntityData.YFilter = igmpSnooping.YFilter
    igmpSnooping.EntityData.YangName = "igmp-snooping"
    igmpSnooping.EntityData.BundleName = "cisco_ios_xe"
    igmpSnooping.EntityData.ParentYangName = "pw-neighbor-spec"
    igmpSnooping.EntityData.SegmentPath = "igmp-snooping"
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = make(map[string]types.YChild)
    igmpSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    igmpSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", igmpSnooping.ProfileName}
    return &(igmpSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetEntityData() *types.CommonEntityData {
    mldSnooping.EntityData.YFilter = mldSnooping.YFilter
    mldSnooping.EntityData.YangName = "mld-snooping"
    mldSnooping.EntityData.BundleName = "cisco_ios_xe"
    mldSnooping.EntityData.ParentYangName = "pw-neighbor-spec"
    mldSnooping.EntityData.SegmentPath = "mld-snooping"
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = make(map[string]types.YChild)
    mldSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    mldSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", mldSnooping.ProfileName}
    return &(mldSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetEntityData() *types.CommonEntityData {
    dhcpIpv4Snooping.EntityData.YFilter = dhcpIpv4Snooping.YFilter
    dhcpIpv4Snooping.EntityData.YangName = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.BundleName = "cisco_ios_xe"
    dhcpIpv4Snooping.EntityData.ParentYangName = "pw-neighbor-spec"
    dhcpIpv4Snooping.EntityData.SegmentPath = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = make(map[string]types.YChild)
    dhcpIpv4Snooping.EntityData.Leafs = make(map[string]types.YLeaf)
    dhcpIpv4Snooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName}
    return &(dhcpIpv4Snooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "pw-neighbor-spec"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = make(map[string]types.YChild)
    flooding.EntityData.Leafs = make(map[string]types.YLeaf)
    flooding.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", flooding.Disabled}
    flooding.EntityData.Leafs["disabled-unknown-unicast"] = types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast}
    return &(flooding.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionDropActionSnmpTrapActionShutdown.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "pw-neighbor-spec"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = make(map[string]types.YChild)
    stormControl.EntityData.Children["thresholds"] = types.YChild{"Thresholds", nil}
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children[types.GetSegmentPath(&stormControl.Thresholds[i])] = types.YChild{"Thresholds", &stormControl.Thresholds[i]}
    }
    stormControl.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControl.EntityData.Leafs["action"] = types.YLeaf{"Action", stormControl.Action}
    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type for
    // which an administrator desires to configure storm control. The type is
    // EthTrafficClass.
    TrafficClass interface{}

    // Traffic threshold value. Unit of the value is indicated by leaf 'unit'. The
    // type is interface{} with range: 0..4294967295. This attribute is mandatory.
    Value interface{}

    // This enumeration define unit of the traffic threshold value. The type is
    // Unit. This attribute is mandatory.
    Unit interface{}
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xe"
    thresholds.EntityData.ParentYangName = "storm-control"
    thresholds.EntityData.SegmentPath = "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = make(map[string]types.YChild)
    thresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholds.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", thresholds.TrafficClass}
    thresholds.EntityData.Leafs["value"] = types.YLeaf{"Value", thresholds.Value}
    thresholds.EntityData.Leafs["unit"] = types.YLeaf{"Unit", thresholds.Unit}
    return &(thresholds.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit represents value.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit string

const (
    // Bits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit_bps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit = "bps"

    // Kilobits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit_kbps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit = "kbps"

    // Packets per seconds.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit_pps BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds_Unit = "pps"
)

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup
// Backup pseudo-wire.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 or IPv6 address of the neighbor. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NeighborIpAddress interface{}

    // Pseudowire VC ID. The type is interface{} with range: 1..4294967295.
    VcId interface{}

    // Reference to a pseudowire template. The type is string. Refers to
    // cisco_pw.PseudowireConfig_PwTemplates_PwTemplate_Name
    PwClassTemplate interface{}
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetEntityData() *types.CommonEntityData {
    backup.EntityData.YFilter = backup.YFilter
    backup.EntityData.YangName = "backup"
    backup.EntityData.BundleName = "cisco_ios_xe"
    backup.EntityData.ParentYangName = "pw-neighbor-spec"
    backup.EntityData.SegmentPath = "backup"
    backup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    backup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    backup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    backup.EntityData.Children = make(map[string]types.YChild)
    backup.EntityData.Leafs = make(map[string]types.YLeaf)
    backup.EntityData.Leafs["neighbor-ip-address"] = types.YLeaf{"NeighborIpAddress", backup.NeighborIpAddress}
    backup.EntityData.Leafs["vc-id"] = types.YLeaf{"VcId", backup.VcId}
    backup.EntityData.Leafs["pw-class-template"] = types.YLeaf{"PwClassTemplate", backup.PwClassTemplate}
    return &(backup.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable disable mac learning. The type is bool. The default value is true.
    LearningEnabled interface{}

    // MAC table learning limit.
    Limit BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit

    // MAC aging configurations.
    Aging BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging

    // Port down event.
    PortDown BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown

    // Flooding configurations.
    Flooding BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding

    // MAC secure parameters.
    Secure BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure

    // Static mac address list parameters.
    Static BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xe"
    mac.EntityData.ParentYangName = "bridge-domain"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["limit"] = types.YChild{"Limit", &mac.Limit}
    mac.EntityData.Children["aging"] = types.YChild{"Aging", &mac.Aging}
    mac.EntityData.Children["port-down"] = types.YChild{"PortDown", &mac.PortDown}
    mac.EntityData.Children["flooding"] = types.YChild{"Flooding", &mac.Flooding}
    mac.EntityData.Children["secure"] = types.YChild{"Secure", &mac.Secure}
    mac.EntityData.Children["static"] = types.YChild{"Static", &mac.Static}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    mac.EntityData.Leafs["learning-enabled"] = types.YLeaf{"LearningEnabled", mac.LearningEnabled}
    return &(mac.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifNoneNotifSnmpTrapNotifSyslogNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetEntityData() *types.CommonEntityData {
    limit.EntityData.YFilter = limit.YFilter
    limit.EntityData.YangName = "limit"
    limit.EntityData.BundleName = "cisco_ios_xe"
    limit.EntityData.ParentYangName = "mac"
    limit.EntityData.SegmentPath = "limit"
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = make(map[string]types.YChild)
    limit.EntityData.Leafs = make(map[string]types.YLeaf)
    limit.EntityData.Leafs["maximum"] = types.YLeaf{"Maximum", limit.Maximum}
    limit.EntityData.Leafs["action"] = types.YLeaf{"Action", limit.Action}
    limit.EntityData.Leafs["notification"] = types.YLeaf{"Notification", limit.Notification}
    return &(limit.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type_ interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = make(map[string]types.YChild)
    aging.EntityData.Leafs = make(map[string]types.YLeaf)
    aging.EntityData.Leafs["time"] = types.YLeaf{"Time", aging.Time}
    aging.EntityData.Leafs["type"] = types.YLeaf{"Type_", aging.Type_}
    return &(aging.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetEntityData() *types.CommonEntityData {
    portDown.EntityData.YFilter = portDown.YFilter
    portDown.EntityData.YangName = "port-down"
    portDown.EntityData.BundleName = "cisco_ios_xe"
    portDown.EntityData.ParentYangName = "mac"
    portDown.EntityData.SegmentPath = "port-down"
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = make(map[string]types.YChild)
    portDown.EntityData.Leafs = make(map[string]types.YLeaf)
    portDown.EntityData.Leafs["flush"] = types.YLeaf{"Flush", portDown.Flush}
    return &(portDown.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "mac"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = make(map[string]types.YChild)
    flooding.EntityData.Leafs = make(map[string]types.YLeaf)
    flooding.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", flooding.Disabled}
    flooding.EntityData.Leafs["disabled-unknown-unicast"] = types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast}
    return &(flooding.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure
// MAC secure parameters.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetEntityData() *types.CommonEntityData {
    secure.EntityData.YFilter = secure.YFilter
    secure.EntityData.YangName = "secure"
    secure.EntityData.BundleName = "cisco_ios_xe"
    secure.EntityData.ParentYangName = "mac"
    secure.EntityData.SegmentPath = "secure"
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = make(map[string]types.YChild)
    secure.EntityData.Leafs = make(map[string]types.YLeaf)
    secure.EntityData.Leafs["action"] = types.YLeaf{"Action", secure.Action}
    secure.EntityData.Leafs["logging"] = types.YLeaf{"Logging", secure.Logging}
    return &(secure.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static
// Static mac address list parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address entry. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses.
    MacAddresses []BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xe"
    static.EntityData.ParentYangName = "mac"
    static.EntityData.SegmentPath = "static"
    static.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    static.EntityData.Children = make(map[string]types.YChild)
    static.EntityData.Children["mac-addresses"] = types.YChild{"MacAddresses", nil}
    for i := range static.MacAddresses {
        static.EntityData.Children[types.GetSegmentPath(&static.MacAddresses[i])] = types.YChild{"MacAddresses", &static.MacAddresses[i]}
    }
    static.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(static.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
// MAC address entry.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddr interface{}

    // Drop packet. The type is bool. This attribute is mandatory.
    Drop interface{}
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetEntityData() *types.CommonEntityData {
    macAddresses.EntityData.YFilter = macAddresses.YFilter
    macAddresses.EntityData.YangName = "mac-addresses"
    macAddresses.EntityData.BundleName = "cisco_ios_xe"
    macAddresses.EntityData.ParentYangName = "static"
    macAddresses.EntityData.SegmentPath = "mac-addresses" + "[mac-addr='" + fmt.Sprintf("%v", macAddresses.MacAddr) + "']"
    macAddresses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    macAddresses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    macAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    macAddresses.EntityData.Children = make(map[string]types.YChild)
    macAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    macAddresses.EntityData.Leafs["mac-addr"] = types.YLeaf{"MacAddr", macAddresses.MacAddr}
    macAddresses.EntityData.Leafs["drop"] = types.YLeaf{"Drop", macAddresses.Drop}
    return &(macAddresses.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection
// Dynamic ARP Inspection (DAI) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable DAI logging. The type is bool.
    Logging interface{}

    // Enable address validation.
    AddressValidation BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetEntityData() *types.CommonEntityData {
    dynamicArpInspection.EntityData.YFilter = dynamicArpInspection.YFilter
    dynamicArpInspection.EntityData.YangName = "dynamic-arp-inspection"
    dynamicArpInspection.EntityData.BundleName = "cisco_ios_xe"
    dynamicArpInspection.EntityData.ParentYangName = "bridge-domain"
    dynamicArpInspection.EntityData.SegmentPath = "dynamic-arp-inspection"
    dynamicArpInspection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dynamicArpInspection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dynamicArpInspection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dynamicArpInspection.EntityData.Children = make(map[string]types.YChild)
    dynamicArpInspection.EntityData.Children["address-validation"] = types.YChild{"AddressValidation", &dynamicArpInspection.AddressValidation}
    dynamicArpInspection.EntityData.Leafs = make(map[string]types.YLeaf)
    dynamicArpInspection.EntityData.Leafs["logging"] = types.YLeaf{"Logging", dynamicArpInspection.Logging}
    return &(dynamicArpInspection.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match Destination MAC Address. The type is interface{}.
    DstMac interface{}

    // Match Source MAC Address. The type is interface{}.
    SrcMac interface{}

    // Match IPv4 Address. The type is interface{}.
    Ipv4 interface{}
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetEntityData() *types.CommonEntityData {
    addressValidation.EntityData.YFilter = addressValidation.YFilter
    addressValidation.EntityData.YangName = "address-validation"
    addressValidation.EntityData.BundleName = "cisco_ios_xe"
    addressValidation.EntityData.ParentYangName = "dynamic-arp-inspection"
    addressValidation.EntityData.SegmentPath = "address-validation"
    addressValidation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressValidation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressValidation.EntityData.Children = make(map[string]types.YChild)
    addressValidation.EntityData.Leafs = make(map[string]types.YLeaf)
    addressValidation.EntityData.Leafs["dst-mac"] = types.YLeaf{"DstMac", addressValidation.DstMac}
    addressValidation.EntityData.Leafs["src-mac"] = types.YLeaf{"SrcMac", addressValidation.SrcMac}
    addressValidation.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", addressValidation.Ipv4}
    return &(addressValidation.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard
// IP source guard (IPSG) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IPSG logging. The type is bool. The default value is false.
    Logging interface{}
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xe"
    ipSourceGuard.EntityData.ParentYangName = "bridge-domain"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSourceGuard.EntityData.Children = make(map[string]types.YChild)
    ipSourceGuard.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSourceGuard.EntityData.Leafs["logging"] = types.YLeaf{"Logging", ipSourceGuard.Logging}
    return &(ipSourceGuard.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionDropActionSnmpTrapActionShutdown.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "bridge-domain"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = make(map[string]types.YChild)
    stormControl.EntityData.Children["thresholds"] = types.YChild{"Thresholds", nil}
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children[types.GetSegmentPath(&stormControl.Thresholds[i])] = types.YChild{"Thresholds", &stormControl.Thresholds[i]}
    }
    stormControl.EntityData.Leafs = make(map[string]types.YLeaf)
    stormControl.EntityData.Leafs["action"] = types.YLeaf{"Action", stormControl.Action}
    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type for
    // which an administrator desires to configure storm control. The type is
    // EthTrafficClass.
    TrafficClass interface{}

    // Traffic threshold value. Unit of the value is indicated by leaf 'unit'. The
    // type is interface{} with range: 0..4294967295. This attribute is mandatory.
    Value interface{}

    // This enumeration define unit of the traffic threshold value. The type is
    // Unit. This attribute is mandatory.
    Unit interface{}
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetEntityData() *types.CommonEntityData {
    thresholds.EntityData.YFilter = thresholds.YFilter
    thresholds.EntityData.YangName = "thresholds"
    thresholds.EntityData.BundleName = "cisco_ios_xe"
    thresholds.EntityData.ParentYangName = "storm-control"
    thresholds.EntityData.SegmentPath = "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = make(map[string]types.YChild)
    thresholds.EntityData.Leafs = make(map[string]types.YLeaf)
    thresholds.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", thresholds.TrafficClass}
    thresholds.EntityData.Leafs["value"] = types.YLeaf{"Value", thresholds.Value}
    thresholds.EntityData.Leafs["unit"] = types.YLeaf{"Unit", thresholds.Unit}
    return &(thresholds.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit represents value.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit string

const (
    // Bits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit_bps BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit = "bps"

    // Kilobits per second.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit_kbps BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit = "kbps"

    // Packets per seconds.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit_pps BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds_Unit = "pps"
)

// BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping
// Enable IGMP snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string.
    ProfileName interface{}

    // Disable IGMP snooping. The type is interface{}.
    Disabled interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetEntityData() *types.CommonEntityData {
    igmpSnooping.EntityData.YFilter = igmpSnooping.YFilter
    igmpSnooping.EntityData.YangName = "igmp-snooping"
    igmpSnooping.EntityData.BundleName = "cisco_ios_xe"
    igmpSnooping.EntityData.ParentYangName = "bridge-domain"
    igmpSnooping.EntityData.SegmentPath = "igmp-snooping"
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = make(map[string]types.YChild)
    igmpSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    igmpSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", igmpSnooping.ProfileName}
    igmpSnooping.EntityData.Leafs["disabled"] = types.YLeaf{"Disabled", igmpSnooping.Disabled}
    return &(igmpSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetEntityData() *types.CommonEntityData {
    mldSnooping.EntityData.YFilter = mldSnooping.YFilter
    mldSnooping.EntityData.YangName = "mld-snooping"
    mldSnooping.EntityData.BundleName = "cisco_ios_xe"
    mldSnooping.EntityData.ParentYangName = "bridge-domain"
    mldSnooping.EntityData.SegmentPath = "mld-snooping"
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = make(map[string]types.YChild)
    mldSnooping.EntityData.Leafs = make(map[string]types.YLeaf)
    mldSnooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", mldSnooping.ProfileName}
    return &(mldSnooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetEntityData() *types.CommonEntityData {
    dhcpIpv4Snooping.EntityData.YFilter = dhcpIpv4Snooping.YFilter
    dhcpIpv4Snooping.EntityData.YangName = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.BundleName = "cisco_ios_xe"
    dhcpIpv4Snooping.EntityData.ParentYangName = "bridge-domain"
    dhcpIpv4Snooping.EntityData.SegmentPath = "dhcp-ipv4-snooping"
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = make(map[string]types.YChild)
    dhcpIpv4Snooping.EntityData.Leafs = make(map[string]types.YLeaf)
    dhcpIpv4Snooping.EntityData.Leafs["profile-name"] = types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName}
    return &(dhcpIpv4Snooping.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode represents Flood modes for optimization.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode string

const (
    // Flood mode optimized for convergence.
    BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode_convergence_optimized BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode = "convergence-optimized"

    // Flood mode optimized for resiliency
    BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode_resilience_optimized BridgeDomainConfig_BridgeDomains_BridgeDomain_FloodingMode = "resilience-optimized"
)

// BridgeDomainState
// This container defines bridge-domain operational data.
type BridgeDomainState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This container defines system capabilities for bridge domain.
    SystemCapabilities BridgeDomainState_SystemCapabilities

    // This container defines module capabilities for bridge domain.
    ModuleCapabilities BridgeDomainState_ModuleCapabilities

    // Bridge domain state data.
    BridgeDomains BridgeDomainState_BridgeDomains

    // This list contains mac-address entries for bridge domains. The type is
    // slice of BridgeDomainState_MacTable.
    MacTable []BridgeDomainState_MacTable
}

func (bridgeDomainState *BridgeDomainState) GetEntityData() *types.CommonEntityData {
    bridgeDomainState.EntityData.YFilter = bridgeDomainState.YFilter
    bridgeDomainState.EntityData.YangName = "bridge-domain-state"
    bridgeDomainState.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomainState.EntityData.ParentYangName = "cisco-bridge-domain"
    bridgeDomainState.EntityData.SegmentPath = "cisco-bridge-domain:bridge-domain-state"
    bridgeDomainState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomainState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomainState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomainState.EntityData.Children = make(map[string]types.YChild)
    bridgeDomainState.EntityData.Children["system-capabilities"] = types.YChild{"SystemCapabilities", &bridgeDomainState.SystemCapabilities}
    bridgeDomainState.EntityData.Children["module-capabilities"] = types.YChild{"ModuleCapabilities", &bridgeDomainState.ModuleCapabilities}
    bridgeDomainState.EntityData.Children["bridge-domains"] = types.YChild{"BridgeDomains", &bridgeDomainState.BridgeDomains}
    bridgeDomainState.EntityData.Children["mac-table"] = types.YChild{"MacTable", nil}
    for i := range bridgeDomainState.MacTable {
        bridgeDomainState.EntityData.Children[types.GetSegmentPath(&bridgeDomainState.MacTable[i])] = types.YChild{"MacTable", &bridgeDomainState.MacTable[i]}
    }
    bridgeDomainState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomainState.EntityData)
}

// BridgeDomainState_SystemCapabilities
// This container defines system capabilities for bridge
// domain.
type BridgeDomainState_SystemCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of bridge-domains suported. The type is interface{} with
    // range: 0..4294967295.
    MaxBd interface{}

    // Maximum number of attachment circuits per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxAcPerBd interface{}

    // Maximum number of access pseudowires per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxPwPerBd interface{}

    // Maximum number of virtual forwarding instances per bridge-domains. The type
    // is interface{} with range: 0..4294967295.
    MaxVfiPerBd interface{}

    // Maximum number of Split Horizon groups per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxShGroupPerBd interface{}

    // Maximum number of Interflex interfaces per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxInterflexIfPerBd interface{}
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetEntityData() *types.CommonEntityData {
    systemCapabilities.EntityData.YFilter = systemCapabilities.YFilter
    systemCapabilities.EntityData.YangName = "system-capabilities"
    systemCapabilities.EntityData.BundleName = "cisco_ios_xe"
    systemCapabilities.EntityData.ParentYangName = "bridge-domain-state"
    systemCapabilities.EntityData.SegmentPath = "system-capabilities"
    systemCapabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    systemCapabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    systemCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    systemCapabilities.EntityData.Children = make(map[string]types.YChild)
    systemCapabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    systemCapabilities.EntityData.Leafs["max-bd"] = types.YLeaf{"MaxBd", systemCapabilities.MaxBd}
    systemCapabilities.EntityData.Leafs["max-ac-per-bd"] = types.YLeaf{"MaxAcPerBd", systemCapabilities.MaxAcPerBd}
    systemCapabilities.EntityData.Leafs["max-pw-per-bd"] = types.YLeaf{"MaxPwPerBd", systemCapabilities.MaxPwPerBd}
    systemCapabilities.EntityData.Leafs["max-vfi-per-bd"] = types.YLeaf{"MaxVfiPerBd", systemCapabilities.MaxVfiPerBd}
    systemCapabilities.EntityData.Leafs["max-sh-group-per-bd"] = types.YLeaf{"MaxShGroupPerBd", systemCapabilities.MaxShGroupPerBd}
    systemCapabilities.EntityData.Leafs["max-interflex-if-per-bd"] = types.YLeaf{"MaxInterflexIfPerBd", systemCapabilities.MaxInterflexIfPerBd}
    return &(systemCapabilities.EntityData)
}

// BridgeDomainState_ModuleCapabilities
// This container defines module capabilities for bridge
// domain.
type BridgeDomainState_ModuleCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of capabillity statements for hardware module in the system. The
    // type is slice of BridgeDomainState_ModuleCapabilities_Modules.
    Modules []BridgeDomainState_ModuleCapabilities_Modules
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetEntityData() *types.CommonEntityData {
    moduleCapabilities.EntityData.YFilter = moduleCapabilities.YFilter
    moduleCapabilities.EntityData.YangName = "module-capabilities"
    moduleCapabilities.EntityData.BundleName = "cisco_ios_xe"
    moduleCapabilities.EntityData.ParentYangName = "bridge-domain-state"
    moduleCapabilities.EntityData.SegmentPath = "module-capabilities"
    moduleCapabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    moduleCapabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    moduleCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    moduleCapabilities.EntityData.Children = make(map[string]types.YChild)
    moduleCapabilities.EntityData.Children["modules"] = types.YChild{"Modules", nil}
    for i := range moduleCapabilities.Modules {
        moduleCapabilities.EntityData.Children[types.GetSegmentPath(&moduleCapabilities.Modules[i])] = types.YChild{"Modules", &moduleCapabilities.Modules[i]}
    }
    moduleCapabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(moduleCapabilities.EntityData)
}

// BridgeDomainState_ModuleCapabilities_Modules
// Collection of capabillity statements for hardware
// module in the system.
type BridgeDomainState_ModuleCapabilities_Modules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the hardware module such as linecards, for
    // which capability is described. The type is string.
    Name interface{}

    // Maximum number of MAC addresses per bridge-domains supported by this
    // module. The type is interface{} with range: 0..4294967295.
    MaxMacPerBd interface{}

    // Maximum number of PBB Edge type bridge-domains supported by this module.
    // The type is interface{} with range: 0..4294967295.
    MaxPddEdgeBd interface{}

    // Maximum number of bridge-domains suported. The type is interface{} with
    // range: 0..4294967295.
    MaxBd interface{}

    // Maximum number of attachment circuits per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxAcPerBd interface{}

    // Maximum number of access pseudowires per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxPwPerBd interface{}

    // Maximum number of virtual forwarding instances per bridge-domains. The type
    // is interface{} with range: 0..4294967295.
    MaxVfiPerBd interface{}

    // Maximum number of Split Horizon groups per bridge-domains. The type is
    // interface{} with range: 0..4294967295.
    MaxShGroupPerBd interface{}
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetEntityData() *types.CommonEntityData {
    modules.EntityData.YFilter = modules.YFilter
    modules.EntityData.YangName = "modules"
    modules.EntityData.BundleName = "cisco_ios_xe"
    modules.EntityData.ParentYangName = "module-capabilities"
    modules.EntityData.SegmentPath = "modules" + "[name='" + fmt.Sprintf("%v", modules.Name) + "']"
    modules.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    modules.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    modules.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    modules.EntityData.Children = make(map[string]types.YChild)
    modules.EntityData.Leafs = make(map[string]types.YLeaf)
    modules.EntityData.Leafs["name"] = types.YLeaf{"Name", modules.Name}
    modules.EntityData.Leafs["max-mac-per-bd"] = types.YLeaf{"MaxMacPerBd", modules.MaxMacPerBd}
    modules.EntityData.Leafs["max-pdd-edge-bd"] = types.YLeaf{"MaxPddEdgeBd", modules.MaxPddEdgeBd}
    modules.EntityData.Leafs["max-bd"] = types.YLeaf{"MaxBd", modules.MaxBd}
    modules.EntityData.Leafs["max-ac-per-bd"] = types.YLeaf{"MaxAcPerBd", modules.MaxAcPerBd}
    modules.EntityData.Leafs["max-pw-per-bd"] = types.YLeaf{"MaxPwPerBd", modules.MaxPwPerBd}
    modules.EntityData.Leafs["max-vfi-per-bd"] = types.YLeaf{"MaxVfiPerBd", modules.MaxVfiPerBd}
    modules.EntityData.Leafs["max-sh-group-per-bd"] = types.YLeaf{"MaxShGroupPerBd", modules.MaxShGroupPerBd}
    return &(modules.EntityData)
}

// BridgeDomainState_BridgeDomains
// Bridge domain state data.
type BridgeDomainState_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of bridge-domain state data. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain.
    BridgeDomain []BridgeDomainState_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-state"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomains.EntityData.Children = make(map[string]types.YChild)
    bridgeDomains.EntityData.Children["bridge-domain"] = types.YChild{"BridgeDomain", nil}
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children[types.GetSegmentPath(&bridgeDomains.BridgeDomain[i])] = types.YChild{"BridgeDomain", &bridgeDomains.BridgeDomain[i]}
    }
    bridgeDomains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridgeDomains.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain
// Collection of bridge-domain state data.
type BridgeDomainState_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain name or number. The type is string.
    Id interface{}

    // Bridge domain operational/admin status. The type is BridgeDomainStateType.
    // This attribute is mandatory.
    BdState interface{}

    // System time when this bridge-domain was created. The type is interface{}
    // with range: 0..4294967295.
    CreateTime interface{}

    // Number of consecutive ticks since bridge-domain status was changed last
    // time. The type is interface{} with range: 0..4294967295.
    LastStatusChange interface{}

    // This leaf indicates if MAC address limit has been reached. The type is
    // bool.
    MacLimitReached interface{}

    // Point to Mutipoint pseudowire state. The type is bool.
    P2MpPwDisabled interface{}

    // Collection of bridge-domain members.
    Members BridgeDomainState_BridgeDomains_BridgeDomain_Members
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + "[id='" + fmt.Sprintf("%v", bridgeDomain.Id) + "']"
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = make(map[string]types.YChild)
    bridgeDomain.EntityData.Children["members"] = types.YChild{"Members", &bridgeDomain.Members}
    bridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomain.EntityData.Leafs["id"] = types.YLeaf{"Id", bridgeDomain.Id}
    bridgeDomain.EntityData.Leafs["bd-state"] = types.YLeaf{"BdState", bridgeDomain.BdState}
    bridgeDomain.EntityData.Leafs["create-time"] = types.YLeaf{"CreateTime", bridgeDomain.CreateTime}
    bridgeDomain.EntityData.Leafs["last-status-change"] = types.YLeaf{"LastStatusChange", bridgeDomain.LastStatusChange}
    bridgeDomain.EntityData.Leafs["mac-limit-reached"] = types.YLeaf{"MacLimitReached", bridgeDomain.MacLimitReached}
    bridgeDomain.EntityData.Leafs["p2mp-pw-disabled"] = types.YLeaf{"P2MpPwDisabled", bridgeDomain.P2MpPwDisabled}
    return &(bridgeDomain.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of attachment circuits for this bridge domains. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember.
    AcMember []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember

    // Reference to an instance of Bridge domain Virtual Forwarding Instance (VFI)
    // name. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember.
    VfiMember []BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember

    // Collection of access pseudowire members of the bridge domain. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember.
    AccessPwMember []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "cisco_ios_xe"
    members.EntityData.ParentYangName = "bridge-domain"
    members.EntityData.SegmentPath = "members"
    members.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    members.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    members.EntityData.Children = make(map[string]types.YChild)
    members.EntityData.Children["ac-member"] = types.YChild{"AcMember", nil}
    for i := range members.AcMember {
        members.EntityData.Children[types.GetSegmentPath(&members.AcMember[i])] = types.YChild{"AcMember", &members.AcMember[i]}
    }
    members.EntityData.Children["vfi-member"] = types.YChild{"VfiMember", nil}
    for i := range members.VfiMember {
        members.EntityData.Children[types.GetSegmentPath(&members.VfiMember[i])] = types.YChild{"VfiMember", &members.VfiMember[i]}
    }
    members.EntityData.Children["access-pw-member"] = types.YChild{"AccessPwMember", nil}
    for i := range members.AccessPwMember {
        members.EntityData.Children[types.GetSegmentPath(&members.AccessPwMember[i])] = types.YChild{"AccessPwMember", &members.AccessPwMember[i]}
    }
    members.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(members.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember
// List of attachment circuits for this bridge domains
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an instance of Bridge domain
    // attachment circuit (AC) interface name. The type is string. Refers to
    // ietf_interfaces.InterfacesState_Interface_Name
    Interface_ interface{}

    // Number of static MAC address configured on this bridge-domain member
    // interface. The type is interface{} with range: 0..4294967295.
    StaticMacCount interface{}

    // Dynamic ARP Inspection (DAI) statistics.
    DaiStats BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats

    // IPSG stats.
    IpsgStats BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats

    // Storm control statistics.
    StormControl BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetEntityData() *types.CommonEntityData {
    acMember.EntityData.YFilter = acMember.YFilter
    acMember.EntityData.YangName = "ac-member"
    acMember.EntityData.BundleName = "cisco_ios_xe"
    acMember.EntityData.ParentYangName = "members"
    acMember.EntityData.SegmentPath = "ac-member" + "[interface='" + fmt.Sprintf("%v", acMember.Interface_) + "']"
    acMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    acMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    acMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    acMember.EntityData.Children = make(map[string]types.YChild)
    acMember.EntityData.Children["dai-stats"] = types.YChild{"DaiStats", &acMember.DaiStats}
    acMember.EntityData.Children["ipsg-stats"] = types.YChild{"IpsgStats", &acMember.IpsgStats}
    acMember.EntityData.Children["storm-control"] = types.YChild{"StormControl", &acMember.StormControl}
    acMember.EntityData.Leafs = make(map[string]types.YLeaf)
    acMember.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", acMember.Interface_}
    acMember.EntityData.Leafs["static-mac-count"] = types.YLeaf{"StaticMacCount", acMember.StaticMacCount}
    return &(acMember.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats
// Dynamic ARP Inspection (DAI) statistics.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets dropped by interface due to DAI actions. The type is
    // interface{} with range: 0..18446744073709551615.
    PacketDrops interface{}

    // Number of bytes dropped by interface due to DAI actions. The type is
    // interface{} with range: 0..18446744073709551615.
    ByteDrops interface{}
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetEntityData() *types.CommonEntityData {
    daiStats.EntityData.YFilter = daiStats.YFilter
    daiStats.EntityData.YangName = "dai-stats"
    daiStats.EntityData.BundleName = "cisco_ios_xe"
    daiStats.EntityData.ParentYangName = "ac-member"
    daiStats.EntityData.SegmentPath = "dai-stats"
    daiStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daiStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daiStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daiStats.EntityData.Children = make(map[string]types.YChild)
    daiStats.EntityData.Leafs = make(map[string]types.YLeaf)
    daiStats.EntityData.Leafs["packet-drops"] = types.YLeaf{"PacketDrops", daiStats.PacketDrops}
    daiStats.EntityData.Leafs["byte-drops"] = types.YLeaf{"ByteDrops", daiStats.ByteDrops}
    return &(daiStats.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats
// IPSG stats.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets dropped by interface due to IPSG actions. The type is
    // interface{} with range: 0..18446744073709551615.
    PacketDrops interface{}

    // Number of bytes dropped by interface due to IPSG actions. The type is
    // interface{} with range: 0..18446744073709551615.
    ByteDrops interface{}
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetEntityData() *types.CommonEntityData {
    ipsgStats.EntityData.YFilter = ipsgStats.YFilter
    ipsgStats.EntityData.YangName = "ipsg-stats"
    ipsgStats.EntityData.BundleName = "cisco_ios_xe"
    ipsgStats.EntityData.ParentYangName = "ac-member"
    ipsgStats.EntityData.SegmentPath = "ipsg-stats"
    ipsgStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsgStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsgStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsgStats.EntityData.Children = make(map[string]types.YChild)
    ipsgStats.EntityData.Leafs = make(map[string]types.YLeaf)
    ipsgStats.EntityData.Leafs["packet-drops"] = types.YLeaf{"PacketDrops", ipsgStats.PacketDrops}
    ipsgStats.EntityData.Leafs["byte-drops"] = types.YLeaf{"ByteDrops", ipsgStats.ByteDrops}
    return &(ipsgStats.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl
// Storm control statistics.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of packet drop statistics for ethernet traffic clasess. The type
    // is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter.
    DropCounter []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "ac-member"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = make(map[string]types.YChild)
    stormControl.EntityData.Children["drop-counter"] = types.YChild{"DropCounter", nil}
    for i := range stormControl.DropCounter {
        stormControl.EntityData.Children[types.GetSegmentPath(&stormControl.DropCounter[i])] = types.YChild{"DropCounter", &stormControl.DropCounter[i]}
    }
    stormControl.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stormControl.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
// Collection of packet drop statistics for ethernet traffic
// clasess.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ethernet traffic class i.e. broadcast, multicast
    // or unknown unicast. The type is EthTrafficClass.
    TrafficClass interface{}

    // The total number of dropped packets due to storm control violations. The
    // type is interface{} with range: 0..18446744073709551615.
    PacketDrops interface{}

    // The total number of bytes of traffic dropped due to storm control
    // violations. The type is interface{} with range: 0..18446744073709551615.
    OctateDrops interface{}
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetEntityData() *types.CommonEntityData {
    dropCounter.EntityData.YFilter = dropCounter.YFilter
    dropCounter.EntityData.YangName = "drop-counter"
    dropCounter.EntityData.BundleName = "cisco_ios_xe"
    dropCounter.EntityData.ParentYangName = "storm-control"
    dropCounter.EntityData.SegmentPath = "drop-counter" + "[traffic-class='" + fmt.Sprintf("%v", dropCounter.TrafficClass) + "']"
    dropCounter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dropCounter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dropCounter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dropCounter.EntityData.Children = make(map[string]types.YChild)
    dropCounter.EntityData.Leafs = make(map[string]types.YLeaf)
    dropCounter.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", dropCounter.TrafficClass}
    dropCounter.EntityData.Leafs["packet-drops"] = types.YLeaf{"PacketDrops", dropCounter.PacketDrops}
    dropCounter.EntityData.Leafs["octate-drops"] = types.YLeaf{"OctateDrops", dropCounter.OctateDrops}
    return &(dropCounter.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember
// Reference to an instance of Bridge domain Virtual
// Forwarding Instance (VFI) name.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain memeber interface name. The type is
    // string. Refers to ietf_interfaces.InterfacesState_Interface_Name
    Interface_ interface{}

    // Flooding operational status.
    Flooding BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetEntityData() *types.CommonEntityData {
    vfiMember.EntityData.YFilter = vfiMember.YFilter
    vfiMember.EntityData.YangName = "vfi-member"
    vfiMember.EntityData.BundleName = "cisco_ios_xe"
    vfiMember.EntityData.ParentYangName = "members"
    vfiMember.EntityData.SegmentPath = "vfi-member" + "[interface='" + fmt.Sprintf("%v", vfiMember.Interface_) + "']"
    vfiMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vfiMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vfiMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vfiMember.EntityData.Children = make(map[string]types.YChild)
    vfiMember.EntityData.Children["flooding"] = types.YChild{"Flooding", &vfiMember.Flooding}
    vfiMember.EntityData.Leafs = make(map[string]types.YLeaf)
    vfiMember.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", vfiMember.Interface_}
    return &(vfiMember.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding
// Flooding operational status
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status.
    Status []BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "vfi-member"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = make(map[string]types.YChild)
    flooding.EntityData.Children["status"] = types.YChild{"Status", nil}
    for i := range flooding.Status {
        flooding.EntityData.Children[types.GetSegmentPath(&flooding.Status[i])] = types.YChild{"Status", &flooding.Status[i]}
    }
    flooding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flooding.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type. The
    // type is EthTrafficClass.
    TrafficClass interface{}

    // This leaf indicates if flooding is enabled for corresponding traffic class.
    // The type is bool.
    Enabled interface{}
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xe"
    status.EntityData.ParentYangName = "flooding"
    status.EntityData.SegmentPath = "status" + "[traffic-class='" + fmt.Sprintf("%v", status.TrafficClass) + "']"
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", status.TrafficClass}
    status.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", status.Enabled}
    return &(status.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember
// Collection of access pseudowire members of the bridge
// domain.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to peer ip address of a pseudowire
    // instance. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    VcPeerAddress interface{}

    // This attribute is a key. Reference to vc-id of a pseudowire instance. The
    // type is string with range: 0..4294967295. Refers to
    // cisco_pw.PseudowireState_Pseudowires_VcId
    VcId interface{}

    // Flooding operational status.
    Flooding BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetEntityData() *types.CommonEntityData {
    accessPwMember.EntityData.YFilter = accessPwMember.YFilter
    accessPwMember.EntityData.YangName = "access-pw-member"
    accessPwMember.EntityData.BundleName = "cisco_ios_xe"
    accessPwMember.EntityData.ParentYangName = "members"
    accessPwMember.EntityData.SegmentPath = "access-pw-member" + "[vc-peer-address='" + fmt.Sprintf("%v", accessPwMember.VcPeerAddress) + "']" + "[vc-id='" + fmt.Sprintf("%v", accessPwMember.VcId) + "']"
    accessPwMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwMember.EntityData.Children = make(map[string]types.YChild)
    accessPwMember.EntityData.Children["flooding"] = types.YChild{"Flooding", &accessPwMember.Flooding}
    accessPwMember.EntityData.Leafs = make(map[string]types.YLeaf)
    accessPwMember.EntityData.Leafs["vc-peer-address"] = types.YLeaf{"VcPeerAddress", accessPwMember.VcPeerAddress}
    accessPwMember.EntityData.Leafs["vc-id"] = types.YLeaf{"VcId", accessPwMember.VcId}
    return &(accessPwMember.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding
// Flooding operational status
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status.
    Status []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "access-pw-member"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = make(map[string]types.YChild)
    flooding.EntityData.Children["status"] = types.YChild{"Status", nil}
    for i := range flooding.Status {
        flooding.EntityData.Children[types.GetSegmentPath(&flooding.Status[i])] = types.YChild{"Status", &flooding.Status[i]}
    }
    flooding.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(flooding.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type. The
    // type is EthTrafficClass.
    TrafficClass interface{}

    // This leaf indicates if flooding is enabled for corresponding traffic class.
    // The type is bool.
    Enabled interface{}
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xe"
    status.EntityData.ParentYangName = "flooding"
    status.EntityData.SegmentPath = "status" + "[traffic-class='" + fmt.Sprintf("%v", status.TrafficClass) + "']"
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["traffic-class"] = types.YLeaf{"TrafficClass", status.TrafficClass}
    status.EntityData.Leafs["enabled"] = types.YLeaf{"Enabled", status.Enabled}
    return &(status.EntityData)
}

// BridgeDomainState_MacTable
// This list contains mac-address entries for bridge
// domains.
type BridgeDomainState_MacTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge-domain name where MAC entry is learnt. The
    // type is string.
    BdId interface{}

    // This attribute is a key. MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // MAC address type. The type is MacType.
    MacType interface{}

    // Reference to an interface instance where MAC  address is learnt. The type
    // is string. Refers to ietf_interfaces.Interfaces_Interface_Name This
    // attribute is mandatory.
    Interface_ interface{}

    // Secure MAC address. The type is bool.
    SecureMac interface{}

    // TBD ?NTFY?. The type is bool.
    NtfyMac interface{}

    // Time since mac address was learnt on the interface. The type is interface{}
    // with range: 0..4294967295.
    Age interface{}

    // Linecard / Module where mac address was learnt. The type is string.
    Location interface{}
}

func (macTable *BridgeDomainState_MacTable) GetEntityData() *types.CommonEntityData {
    macTable.EntityData.YFilter = macTable.YFilter
    macTable.EntityData.YangName = "mac-table"
    macTable.EntityData.BundleName = "cisco_ios_xe"
    macTable.EntityData.ParentYangName = "bridge-domain-state"
    macTable.EntityData.SegmentPath = "mac-table" + "[bd-id='" + fmt.Sprintf("%v", macTable.BdId) + "']" + "[mac-address='" + fmt.Sprintf("%v", macTable.MacAddress) + "']"
    macTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    macTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    macTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    macTable.EntityData.Children = make(map[string]types.YChild)
    macTable.EntityData.Leafs = make(map[string]types.YLeaf)
    macTable.EntityData.Leafs["bd-id"] = types.YLeaf{"BdId", macTable.BdId}
    macTable.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", macTable.MacAddress}
    macTable.EntityData.Leafs["mac-type"] = types.YLeaf{"MacType", macTable.MacType}
    macTable.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", macTable.Interface_}
    macTable.EntityData.Leafs["secure-mac"] = types.YLeaf{"SecureMac", macTable.SecureMac}
    macTable.EntityData.Leafs["ntfy-mac"] = types.YLeaf{"NtfyMac", macTable.NtfyMac}
    macTable.EntityData.Leafs["age"] = types.YLeaf{"Age", macTable.Age}
    macTable.EntityData.Leafs["location"] = types.YLeaf{"Location", macTable.Location}
    return &(macTable.EntityData)
}

// BridgeDomainState_MacTable_MacType represents MAC address type.
type BridgeDomainState_MacTable_MacType string

const (
    // MAC address is configured statically.
    BridgeDomainState_MacTable_MacType_static BridgeDomainState_MacTable_MacType = "static"

    // MAC address is learnt dynamicaly.
    BridgeDomainState_MacTable_MacType_dynamic BridgeDomainState_MacTable_MacType = "dynamic"
)

// ClearBridgeDomain
// Clear mac-address table for bridge-domain and allows a bridge
// to forward again after it was put in shutdown state as a
// result of exceeding the configured MAC limit.
type ClearBridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearBridgeDomain_Input

    
    Output ClearBridgeDomain_Output
}

func (clearBridgeDomain *ClearBridgeDomain) GetEntityData() *types.CommonEntityData {
    clearBridgeDomain.EntityData.YFilter = clearBridgeDomain.YFilter
    clearBridgeDomain.EntityData.YangName = "clear-bridge-domain"
    clearBridgeDomain.EntityData.BundleName = "cisco_ios_xe"
    clearBridgeDomain.EntityData.ParentYangName = "cisco-bridge-domain"
    clearBridgeDomain.EntityData.SegmentPath = "cisco-bridge-domain:clear-bridge-domain"
    clearBridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearBridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearBridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearBridgeDomain.EntityData.Children = make(map[string]types.YChild)
    clearBridgeDomain.EntityData.Children["input"] = types.YChild{"Input", &clearBridgeDomain.Input}
    clearBridgeDomain.EntityData.Children["output"] = types.YChild{"Output", &clearBridgeDomain.Output}
    clearBridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearBridgeDomain.EntityData)
}

// ClearBridgeDomain_Input
type ClearBridgeDomain_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear all bridge-domains configured on the device. The type is interface{}.
    All interface{}

    // Clear a single bridge-domain. The type is string.
    BdId interface{}

    // Clears all bridge-domains under this bridge-group. The type is string.
    BgId interface{}
}

func (input *ClearBridgeDomain_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "clear-bridge-domain"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["all"] = types.YLeaf{"All", input.All}
    input.EntityData.Leafs["bd-id"] = types.YLeaf{"BdId", input.BdId}
    input.EntityData.Leafs["bg-id"] = types.YLeaf{"BgId", input.BgId}
    return &(input.EntityData)
}

// ClearBridgeDomain_Output
type ClearBridgeDomain_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *ClearBridgeDomain_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "clear-bridge-domain"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["errstr"] = types.YLeaf{"Errstr", output.Errstr}
    return &(output.EntityData)
}

// ClearMacAddress
// This RPC allows to clear dynamically learnt mac-address
// entries from the mac-address table.
type ClearMacAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearMacAddress_Input

    
    Output ClearMacAddress_Output
}

func (clearMacAddress *ClearMacAddress) GetEntityData() *types.CommonEntityData {
    clearMacAddress.EntityData.YFilter = clearMacAddress.YFilter
    clearMacAddress.EntityData.YangName = "clear-mac-address"
    clearMacAddress.EntityData.BundleName = "cisco_ios_xe"
    clearMacAddress.EntityData.ParentYangName = "cisco-bridge-domain"
    clearMacAddress.EntityData.SegmentPath = "cisco-bridge-domain:clear-mac-address"
    clearMacAddress.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearMacAddress.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearMacAddress.EntityData.Children = make(map[string]types.YChild)
    clearMacAddress.EntityData.Children["input"] = types.YChild{"Input", &clearMacAddress.Input}
    clearMacAddress.EntityData.Children["output"] = types.YChild{"Output", &clearMacAddress.Output}
    clearMacAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clearMacAddress.EntityData)
}

// ClearMacAddress_Input
type ClearMacAddress_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to an interface. Clear mac-addesses learnt by by this interface.
    // The type is string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface_ interface{}

    // Clear a specific mac-address entry from the mac-table. The type is string
    // with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Clear mac-address entries for given bridge-domain(s).
    BridgeDomain ClearMacAddress_Input_BridgeDomain
}

func (input *ClearMacAddress_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "clear-mac-address"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["bridge-domain"] = types.YChild{"BridgeDomain", &input.BridgeDomain}
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", input.Interface_}
    input.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", input.MacAddress}
    return &(input.EntityData)
}

// ClearMacAddress_Input_BridgeDomain
// Clear mac-address entries for given bridge-domain(s).
type ClearMacAddress_Input_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bridge-domain identifier, clear all mac-address entries dynamically learnt
    // on this bridge-domain. The type is string. This attribute is mandatory.
    BdId interface{}

    // Bridge-group identifier, clear all mac-address entries from all
    // bridge-domains under this bridge-group. The type is string. This attribute
    // is mandatory.
    BgId interface{}
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomain.EntityData.ParentYangName = "input"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain"
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = make(map[string]types.YChild)
    bridgeDomain.EntityData.Leafs = make(map[string]types.YLeaf)
    bridgeDomain.EntityData.Leafs["bd-id"] = types.YLeaf{"BdId", bridgeDomain.BdId}
    bridgeDomain.EntityData.Leafs["bg-id"] = types.YLeaf{"BgId", bridgeDomain.BgId}
    return &(bridgeDomain.EntityData)
}

// ClearMacAddress_Output
type ClearMacAddress_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *ClearMacAddress_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "clear-mac-address"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["errstr"] = types.YLeaf{"Errstr", output.Errstr}
    return &(output.EntityData)
}

// CreateParameterizedBridgeDomains
// Create bridge domains automatically from user defined
// parameters.
type CreateParameterizedBridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CreateParameterizedBridgeDomains_Input

    
    Output CreateParameterizedBridgeDomains_Output
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetEntityData() *types.CommonEntityData {
    createParameterizedBridgeDomains.EntityData.YFilter = createParameterizedBridgeDomains.YFilter
    createParameterizedBridgeDomains.EntityData.YangName = "create-parameterized-bridge-domains"
    createParameterizedBridgeDomains.EntityData.BundleName = "cisco_ios_xe"
    createParameterizedBridgeDomains.EntityData.ParentYangName = "cisco-bridge-domain"
    createParameterizedBridgeDomains.EntityData.SegmentPath = "cisco-bridge-domain:create-parameterized-bridge-domains"
    createParameterizedBridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    createParameterizedBridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    createParameterizedBridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    createParameterizedBridgeDomains.EntityData.Children = make(map[string]types.YChild)
    createParameterizedBridgeDomains.EntityData.Children["input"] = types.YChild{"Input", &createParameterizedBridgeDomains.Input}
    createParameterizedBridgeDomains.EntityData.Children["output"] = types.YChild{"Output", &createParameterizedBridgeDomains.Output}
    createParameterizedBridgeDomains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(createParameterizedBridgeDomains.EntityData)
}

// CreateParameterizedBridgeDomains_Input
type CreateParameterizedBridgeDomains_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Parameter for automatic bridge domain creation. The type is Parameter.
    Parameter interface{}

    // Bridge-domain member interface. The type is slice of
    // CreateParameterizedBridgeDomains_Input_Member.
    Member []CreateParameterizedBridgeDomains_Input_Member
}

func (input *CreateParameterizedBridgeDomains_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "create-parameterized-bridge-domains"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = make(map[string]types.YChild)
    input.EntityData.Children["member"] = types.YChild{"Member", nil}
    for i := range input.Member {
        input.EntityData.Children[types.GetSegmentPath(&input.Member[i])] = types.YChild{"Member", &input.Member[i]}
    }
    input.EntityData.Leafs = make(map[string]types.YLeaf)
    input.EntityData.Leafs["parameter"] = types.YLeaf{"Parameter", input.Parameter}
    return &(input.EntityData)
}

// CreateParameterizedBridgeDomains_Input_Member
// Bridge-domain member interface.
type CreateParameterizedBridgeDomains_Input_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an interface instance which is
    // configured to be part of this bridge domain. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface_ interface{}
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "cisco_ios_xe"
    member.EntityData.ParentYangName = "input"
    member.EntityData.SegmentPath = "member" + "[interface='" + fmt.Sprintf("%v", member.Interface_) + "']"
    member.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    member.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    member.EntityData.Children = make(map[string]types.YChild)
    member.EntityData.Leafs = make(map[string]types.YLeaf)
    member.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", member.Interface_}
    return &(member.EntityData)
}

// CreateParameterizedBridgeDomains_Input_Parameter represents Parameter for automatic bridge domain creation.
type CreateParameterizedBridgeDomains_Input_Parameter string

const (
    // Create bridge-domains from vlan encapsulations of the
    // member interfaces.
    CreateParameterizedBridgeDomains_Input_Parameter_vlan CreateParameterizedBridgeDomains_Input_Parameter = "vlan"
)

// CreateParameterizedBridgeDomains_Output
type CreateParameterizedBridgeDomains_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *CreateParameterizedBridgeDomains_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "create-parameterized-bridge-domains"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = make(map[string]types.YChild)
    output.EntityData.Leafs = make(map[string]types.YLeaf)
    output.EntityData.Leafs["errstr"] = types.YLeaf{"Errstr", output.Errstr}
    return &(output.EntityData)
}

