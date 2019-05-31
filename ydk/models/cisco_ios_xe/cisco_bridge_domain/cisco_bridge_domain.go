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
    bridgeDomainConfig.EntityData.AbsolutePath = bridgeDomainConfig.EntityData.SegmentPath
    bridgeDomainConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomainConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomainConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomainConfig.EntityData.Children = types.NewOrderedMap()
    bridgeDomainConfig.EntityData.Children.Append("global", types.YChild{"Global", &bridgeDomainConfig.Global})
    bridgeDomainConfig.EntityData.Children.Append("bridge-groups", types.YChild{"BridgeGroups", &bridgeDomainConfig.BridgeGroups})
    bridgeDomainConfig.EntityData.Children.Append("bridge-domains", types.YChild{"BridgeDomains", &bridgeDomainConfig.BridgeDomains})
    bridgeDomainConfig.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainConfig.EntityData.YListKeys = []string {}

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
    global.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/" + global.EntityData.SegmentPath
    global.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("pbb", types.YChild{"Pbb", &global.Pbb})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("bd-state-notification-enabled", types.YLeaf{"BdStateNotificationEnabled", global.BdStateNotificationEnabled})
    global.EntityData.Leafs.Append("bd-state-notification-rate", types.YLeaf{"BdStateNotificationRate", global.BdStateNotificationRate})

    global.EntityData.YListKeys = []string {}

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
    pbb.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/global/" + pbb.EntityData.SegmentPath
    pbb.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pbb.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pbb.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pbb.EntityData.Children = types.NewOrderedMap()
    pbb.EntityData.Leafs = types.NewOrderedMap()
    pbb.EntityData.Leafs.Append("backbone-src-mac", types.YLeaf{"BackboneSrcMac", pbb.BackboneSrcMac})

    pbb.EntityData.YListKeys = []string {}

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
    BridgeGroup []*BridgeDomainConfig_BridgeGroups_BridgeGroup
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetEntityData() *types.CommonEntityData {
    bridgeGroups.EntityData.YFilter = bridgeGroups.YFilter
    bridgeGroups.EntityData.YangName = "bridge-groups"
    bridgeGroups.EntityData.BundleName = "cisco_ios_xe"
    bridgeGroups.EntityData.ParentYangName = "bridge-domain-config"
    bridgeGroups.EntityData.SegmentPath = "bridge-groups"
    bridgeGroups.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/" + bridgeGroups.EntityData.SegmentPath
    bridgeGroups.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeGroups.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeGroups.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeGroups.EntityData.Children = types.NewOrderedMap()
    bridgeGroups.EntityData.Children.Append("bridge-group", types.YChild{"BridgeGroup", nil})
    for i := range bridgeGroups.BridgeGroup {
        bridgeGroups.EntityData.Children.Append(types.GetSegmentPath(bridgeGroups.BridgeGroup[i]), types.YChild{"BridgeGroup", bridgeGroups.BridgeGroup[i]})
    }
    bridgeGroups.EntityData.Leafs = types.NewOrderedMap()

    bridgeGroups.EntityData.YListKeys = []string {}

    return &(bridgeGroups.EntityData)
}

// BridgeDomainConfig_BridgeGroups_BridgeGroup
// Bridge-group configuration.
type BridgeDomainConfig_BridgeGroups_BridgeGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Bridge-group name. The type is string with length:
    // 1..32.
    Name interface{}
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetEntityData() *types.CommonEntityData {
    bridgeGroup.EntityData.YFilter = bridgeGroup.YFilter
    bridgeGroup.EntityData.YangName = "bridge-group"
    bridgeGroup.EntityData.BundleName = "cisco_ios_xe"
    bridgeGroup.EntityData.ParentYangName = "bridge-groups"
    bridgeGroup.EntityData.SegmentPath = "bridge-group" + types.AddKeyToken(bridgeGroup.Name, "name")
    bridgeGroup.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-groups/" + bridgeGroup.EntityData.SegmentPath
    bridgeGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeGroup.EntityData.Children = types.NewOrderedMap()
    bridgeGroup.EntityData.Leafs = types.NewOrderedMap()
    bridgeGroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", bridgeGroup.Name})

    bridgeGroup.EntityData.YListKeys = []string {"Name"}

    return &(bridgeGroup.EntityData)
}

// BridgeDomainConfig_BridgeDomains
// Collection of bridge domains.
type BridgeDomainConfig_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Definition of a bridge-domain. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain.
    BridgeDomain []*BridgeDomainConfig_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-config"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/" + bridgeDomains.EntityData.SegmentPath
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomains.EntityData.Children = types.NewOrderedMap()
    bridgeDomains.EntityData.Children.Append("bridge-domain", types.YChild{"BridgeDomain", nil})
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children.Append(types.GetSegmentPath(bridgeDomains.BridgeDomain[i]), types.YChild{"BridgeDomain", bridgeDomains.BridgeDomain[i]})
    }
    bridgeDomains.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomains.EntityData.YListKeys = []string {}

    return &(bridgeDomains.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain
// Definition of a bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + types.AddKeyToken(bridgeDomain.Id, "id")
    bridgeDomain.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/" + bridgeDomain.EntityData.SegmentPath
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = types.NewOrderedMap()
    bridgeDomain.EntityData.Children.Append("members", types.YChild{"Members", &bridgeDomain.Members})
    bridgeDomain.EntityData.Children.Append("mac", types.YChild{"Mac", &bridgeDomain.Mac})
    bridgeDomain.EntityData.Children.Append("dynamic-arp-inspection", types.YChild{"DynamicArpInspection", &bridgeDomain.DynamicArpInspection})
    bridgeDomain.EntityData.Children.Append("ip-source-guard", types.YChild{"IpSourceGuard", &bridgeDomain.IpSourceGuard})
    bridgeDomain.EntityData.Children.Append("storm-control", types.YChild{"StormControl", &bridgeDomain.StormControl})
    bridgeDomain.EntityData.Children.Append("igmp-snooping", types.YChild{"IgmpSnooping", &bridgeDomain.IgmpSnooping})
    bridgeDomain.EntityData.Children.Append("mld-snooping", types.YChild{"MldSnooping", &bridgeDomain.MldSnooping})
    bridgeDomain.EntityData.Children.Append("dhcp-ipv4-snooping", types.YChild{"DhcpIpv4Snooping", &bridgeDomain.DhcpIpv4Snooping})
    bridgeDomain.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs.Append("id", types.YLeaf{"Id", bridgeDomain.Id})
    bridgeDomain.EntityData.Leafs.Append("bridge-group", types.YLeaf{"BridgeGroup", bridgeDomain.BridgeGroup})
    bridgeDomain.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", bridgeDomain.Enabled})
    bridgeDomain.EntityData.Leafs.Append("bd-status-change-notification", types.YLeaf{"BdStatusChangeNotification", bridgeDomain.BdStatusChangeNotification})
    bridgeDomain.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", bridgeDomain.Mtu})
    bridgeDomain.EntityData.Leafs.Append("flooding-mode", types.YLeaf{"FloodingMode", bridgeDomain.FloodingMode})

    bridgeDomain.EntityData.YListKeys = []string {"Id"}

    return &(bridgeDomain.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of Attachment circuits for current bridge-domain. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember.
    AcMember []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember

    // List of Virtual Forrwarding Interfaces for current bridge-domain. The type
    // is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember.
    VfiMember []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember

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
    members.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + members.EntityData.SegmentPath
    members.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    members.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    members.EntityData.Children = types.NewOrderedMap()
    members.EntityData.Children.Append("ac-member", types.YChild{"AcMember", nil})
    for i := range members.AcMember {
        members.EntityData.Children.Append(types.GetSegmentPath(members.AcMember[i]), types.YChild{"AcMember", members.AcMember[i]})
    }
    members.EntityData.Children.Append("vfi-member", types.YChild{"VfiMember", nil})
    for i := range members.VfiMember {
        members.EntityData.Children.Append(types.GetSegmentPath(members.VfiMember[i]), types.YChild{"VfiMember", members.VfiMember[i]})
    }
    members.EntityData.Children.Append("access-pw-member", types.YChild{"AccessPwMember", &members.AccessPwMember})
    members.EntityData.Leafs = types.NewOrderedMap()

    members.EntityData.YListKeys = []string {}

    return &(members.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember
// List of Attachment circuits for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to an attchment circuit interface
    // instance which is configured to be part of this bridge-domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

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
    acMember.EntityData.SegmentPath = "ac-member" + types.AddKeyToken(acMember.Interface, "interface")
    acMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/" + acMember.EntityData.SegmentPath
    acMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    acMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    acMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    acMember.EntityData.Children = types.NewOrderedMap()
    acMember.EntityData.Children.Append("split-horizon-group", types.YChild{"SplitHorizonGroup", &acMember.SplitHorizonGroup})
    acMember.EntityData.Children.Append("mac", types.YChild{"Mac", &acMember.Mac})
    acMember.EntityData.Children.Append("igmp-snooping", types.YChild{"IgmpSnooping", &acMember.IgmpSnooping})
    acMember.EntityData.Children.Append("mld-snooping", types.YChild{"MldSnooping", &acMember.MldSnooping})
    acMember.EntityData.Children.Append("dhcp-ipv4-snooping", types.YChild{"DhcpIpv4Snooping", &acMember.DhcpIpv4Snooping})
    acMember.EntityData.Children.Append("flooding", types.YChild{"Flooding", &acMember.Flooding})
    acMember.EntityData.Children.Append("storm-control", types.YChild{"StormControl", &acMember.StormControl})
    acMember.EntityData.Children.Append("dynamic-arp-inspection", types.YChild{"DynamicArpInspection", &acMember.DynamicArpInspection})
    acMember.EntityData.Children.Append("ip-source-guard", types.YChild{"IpSourceGuard", &acMember.IpSourceGuard})
    acMember.EntityData.Leafs = types.NewOrderedMap()
    acMember.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", acMember.Interface})

    acMember.EntityData.YListKeys = []string {"Interface"}

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
    YPresence bool

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
    splitHorizonGroup.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + splitHorizonGroup.EntityData.SegmentPath
    splitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    splitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    splitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    splitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    splitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    splitHorizonGroup.EntityData.Leafs.Append("id", types.YLeaf{"Id", splitHorizonGroup.Id})

    splitHorizonGroup.EntityData.YListKeys = []string {}

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
    mac.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + mac.EntityData.SegmentPath
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("limit", types.YChild{"Limit", &mac.Limit})
    mac.EntityData.Children.Append("aging", types.YChild{"Aging", &mac.Aging})
    mac.EntityData.Children.Append("port-down", types.YChild{"PortDown", &mac.PortDown})
    mac.EntityData.Children.Append("secure", types.YChild{"Secure", &mac.Secure})
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("learning-enabled", types.YLeaf{"LearningEnabled", mac.LearningEnabled})

    mac.EntityData.YListKeys = []string {}

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
    limit.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/mac/" + limit.EntityData.SegmentPath
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = types.NewOrderedMap()
    limit.EntityData.Leafs = types.NewOrderedMap()
    limit.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", limit.Maximum})
    limit.EntityData.Leafs.Append("action", types.YLeaf{"Action", limit.Action})
    limit.EntityData.Leafs.Append("notification", types.YLeaf{"Notification", limit.Notification})

    limit.EntityData.YListKeys = []string {}

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
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/mac/" + aging.EntityData.SegmentPath
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = types.NewOrderedMap()
    aging.EntityData.Leafs = types.NewOrderedMap()
    aging.EntityData.Leafs.Append("time", types.YLeaf{"Time", aging.Time})
    aging.EntityData.Leafs.Append("type", types.YLeaf{"Type", aging.Type})

    aging.EntityData.YListKeys = []string {}

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
    portDown.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/mac/" + portDown.EntityData.SegmentPath
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = types.NewOrderedMap()
    portDown.EntityData.Leafs = types.NewOrderedMap()
    portDown.EntityData.Leafs.Append("flush", types.YLeaf{"Flush", portDown.Flush})

    portDown.EntityData.YListKeys = []string {}

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
    secure.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/mac/" + secure.EntityData.SegmentPath
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = types.NewOrderedMap()
    secure.EntityData.Leafs = types.NewOrderedMap()
    secure.EntityData.Leafs.Append("action", types.YLeaf{"Action", secure.Action})
    secure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", secure.Logging})
    secure.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", secure.Enabled})

    secure.EntityData.YListKeys = []string {}

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
    igmpSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + igmpSnooping.EntityData.SegmentPath
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", igmpSnooping.ProfileName})

    igmpSnooping.EntityData.YListKeys = []string {}

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
    mldSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + mldSnooping.EntityData.SegmentPath
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", mldSnooping.ProfileName})

    mldSnooping.EntityData.YListKeys = []string {}

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
    dhcpIpv4Snooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + dhcpIpv4Snooping.EntityData.SegmentPath
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName})

    dhcpIpv4Snooping.EntityData.YListKeys = []string {}

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
    flooding.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + flooding.EntityData.SegmentPath
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = types.NewOrderedMap()
    flooding.EntityData.Leafs = types.NewOrderedMap()
    flooding.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", flooding.Disabled})
    flooding.EntityData.Leafs.Append("disabled-unknown-unicast", types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast})

    flooding.EntityData.YListKeys = []string {}

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
    Thresholds []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "ac-member"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + stormControl.EntityData.SegmentPath
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = types.NewOrderedMap()
    stormControl.EntityData.Children.Append("thresholds", types.YChild{"Thresholds", nil})
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children.Append(types.GetSegmentPath(stormControl.Thresholds[i]), types.YChild{"Thresholds", stormControl.Thresholds[i]})
    }
    stormControl.EntityData.Leafs = types.NewOrderedMap()
    stormControl.EntityData.Leafs.Append("action", types.YLeaf{"Action", stormControl.Action})

    stormControl.EntityData.YListKeys = []string {}

    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    thresholds.EntityData.SegmentPath = "thresholds" + types.AddKeyToken(thresholds.TrafficClass, "traffic-class")
    thresholds.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/storm-control/" + thresholds.EntityData.SegmentPath
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = types.NewOrderedMap()
    thresholds.EntityData.Leafs = types.NewOrderedMap()
    thresholds.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", thresholds.TrafficClass})
    thresholds.EntityData.Leafs.Append("value", types.YLeaf{"Value", thresholds.Value})
    thresholds.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", thresholds.Unit})

    thresholds.EntityData.YListKeys = []string {"TrafficClass"}

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
    dynamicArpInspection.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + dynamicArpInspection.EntityData.SegmentPath
    dynamicArpInspection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dynamicArpInspection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dynamicArpInspection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dynamicArpInspection.EntityData.Children = types.NewOrderedMap()
    dynamicArpInspection.EntityData.Children.Append("address-validation", types.YChild{"AddressValidation", &dynamicArpInspection.AddressValidation})
    dynamicArpInspection.EntityData.Leafs = types.NewOrderedMap()
    dynamicArpInspection.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", dynamicArpInspection.Logging})
    dynamicArpInspection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", dynamicArpInspection.Enable})

    dynamicArpInspection.EntityData.YListKeys = []string {}

    return &(dynamicArpInspection.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    addressValidation.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/dynamic-arp-inspection/" + addressValidation.EntityData.SegmentPath
    addressValidation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressValidation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressValidation.EntityData.Children = types.NewOrderedMap()
    addressValidation.EntityData.Leafs = types.NewOrderedMap()
    addressValidation.EntityData.Leafs.Append("dst-mac", types.YLeaf{"DstMac", addressValidation.DstMac})
    addressValidation.EntityData.Leafs.Append("src-mac", types.YLeaf{"SrcMac", addressValidation.SrcMac})
    addressValidation.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addressValidation.Ipv4})

    addressValidation.EntityData.YListKeys = []string {}

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
    ipSourceGuard.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/ac-member/" + ipSourceGuard.EntityData.SegmentPath
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSourceGuard.EntityData.Children = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", ipSourceGuard.Logging})
    ipSourceGuard.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipSourceGuard.Enable})

    ipSourceGuard.EntityData.YListKeys = []string {}

    return &(ipSourceGuard.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember
// List of Virtual Forrwarding Interfaces for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to an Virtual Forwarding Interface
    // instance which is configured to be part of this bridge-domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetEntityData() *types.CommonEntityData {
    vfiMember.EntityData.YFilter = vfiMember.YFilter
    vfiMember.EntityData.YangName = "vfi-member"
    vfiMember.EntityData.BundleName = "cisco_ios_xe"
    vfiMember.EntityData.ParentYangName = "members"
    vfiMember.EntityData.SegmentPath = "vfi-member" + types.AddKeyToken(vfiMember.Interface, "interface")
    vfiMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/" + vfiMember.EntityData.SegmentPath
    vfiMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vfiMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vfiMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vfiMember.EntityData.Children = types.NewOrderedMap()
    vfiMember.EntityData.Leafs = types.NewOrderedMap()
    vfiMember.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", vfiMember.Interface})

    vfiMember.EntityData.YListKeys = []string {"Interface"}

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
    AccessPwIfMember []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember

    // Collection of neighbor specification based pseudo-wires. The type is slice
    // of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec.
    PwNeighborSpec []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetEntityData() *types.CommonEntityData {
    accessPwMember.EntityData.YFilter = accessPwMember.YFilter
    accessPwMember.EntityData.YangName = "access-pw-member"
    accessPwMember.EntityData.BundleName = "cisco_ios_xe"
    accessPwMember.EntityData.ParentYangName = "members"
    accessPwMember.EntityData.SegmentPath = "access-pw-member"
    accessPwMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/" + accessPwMember.EntityData.SegmentPath
    accessPwMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwMember.EntityData.Children = types.NewOrderedMap()
    accessPwMember.EntityData.Children.Append("access-pw-if-member", types.YChild{"AccessPwIfMember", nil})
    for i := range accessPwMember.AccessPwIfMember {
        accessPwMember.EntityData.Children.Append(types.GetSegmentPath(accessPwMember.AccessPwIfMember[i]), types.YChild{"AccessPwIfMember", accessPwMember.AccessPwIfMember[i]})
    }
    accessPwMember.EntityData.Children.Append("pw-neighbor-spec", types.YChild{"PwNeighborSpec", nil})
    for i := range accessPwMember.PwNeighborSpec {
        accessPwMember.EntityData.Children.Append(types.GetSegmentPath(accessPwMember.PwNeighborSpec[i]), types.YChild{"PwNeighborSpec", accessPwMember.PwNeighborSpec[i]})
    }
    accessPwMember.EntityData.Leafs = types.NewOrderedMap()

    accessPwMember.EntityData.YListKeys = []string {}

    return &(accessPwMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember
// List of interface based access pseudowires for
// current bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to an access pseudo-wire interface
    // instance which is configured to be part of this bridge domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetEntityData() *types.CommonEntityData {
    accessPwIfMember.EntityData.YFilter = accessPwIfMember.YFilter
    accessPwIfMember.EntityData.YangName = "access-pw-if-member"
    accessPwIfMember.EntityData.BundleName = "cisco_ios_xe"
    accessPwIfMember.EntityData.ParentYangName = "access-pw-member"
    accessPwIfMember.EntityData.SegmentPath = "access-pw-if-member" + types.AddKeyToken(accessPwIfMember.Interface, "interface")
    accessPwIfMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/" + accessPwIfMember.EntityData.SegmentPath
    accessPwIfMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwIfMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwIfMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwIfMember.EntityData.Children = types.NewOrderedMap()
    accessPwIfMember.EntityData.Leafs = types.NewOrderedMap()
    accessPwIfMember.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", accessPwIfMember.Interface})

    accessPwIfMember.EntityData.YListKeys = []string {"Interface"}

    return &(accessPwIfMember.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec
// Collection of neighbor specification based
// pseudo-wires.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    pwNeighborSpec.EntityData.SegmentPath = "pw-neighbor-spec" + types.AddKeyToken(pwNeighborSpec.NeighborIpAddress, "neighbor-ip-address") + types.AddKeyToken(pwNeighborSpec.VcId, "vc-id")
    pwNeighborSpec.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/" + pwNeighborSpec.EntityData.SegmentPath
    pwNeighborSpec.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pwNeighborSpec.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pwNeighborSpec.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pwNeighborSpec.EntityData.Children = types.NewOrderedMap()
    pwNeighborSpec.EntityData.Children.Append("static-label", types.YChild{"StaticLabel", &pwNeighborSpec.StaticLabel})
    pwNeighborSpec.EntityData.Children.Append("split-horizon-group", types.YChild{"SplitHorizonGroup", &pwNeighborSpec.SplitHorizonGroup})
    pwNeighborSpec.EntityData.Children.Append("mac", types.YChild{"Mac", &pwNeighborSpec.Mac})
    pwNeighborSpec.EntityData.Children.Append("igmp-snooping", types.YChild{"IgmpSnooping", &pwNeighborSpec.IgmpSnooping})
    pwNeighborSpec.EntityData.Children.Append("mld-snooping", types.YChild{"MldSnooping", &pwNeighborSpec.MldSnooping})
    pwNeighborSpec.EntityData.Children.Append("dhcp-ipv4-snooping", types.YChild{"DhcpIpv4Snooping", &pwNeighborSpec.DhcpIpv4Snooping})
    pwNeighborSpec.EntityData.Children.Append("flooding", types.YChild{"Flooding", &pwNeighborSpec.Flooding})
    pwNeighborSpec.EntityData.Children.Append("storm-control", types.YChild{"StormControl", &pwNeighborSpec.StormControl})
    pwNeighborSpec.EntityData.Children.Append("backup", types.YChild{"Backup", &pwNeighborSpec.Backup})
    pwNeighborSpec.EntityData.Leafs = types.NewOrderedMap()
    pwNeighborSpec.EntityData.Leafs.Append("neighbor-ip-address", types.YLeaf{"NeighborIpAddress", pwNeighborSpec.NeighborIpAddress})
    pwNeighborSpec.EntityData.Leafs.Append("vc-id", types.YLeaf{"VcId", pwNeighborSpec.VcId})
    pwNeighborSpec.EntityData.Leafs.Append("pw-class-template", types.YLeaf{"PwClassTemplate", pwNeighborSpec.PwClassTemplate})
    pwNeighborSpec.EntityData.Leafs.Append("encap-type", types.YLeaf{"EncapType", pwNeighborSpec.EncapType})
    pwNeighborSpec.EntityData.Leafs.Append("tag-impose-vlan", types.YLeaf{"TagImposeVlan", pwNeighborSpec.TagImposeVlan})
    pwNeighborSpec.EntityData.Leafs.Append("source-ipv6", types.YLeaf{"SourceIpv6", pwNeighborSpec.SourceIpv6})

    pwNeighborSpec.EntityData.YListKeys = []string {"NeighborIpAddress", "VcId"}

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
    staticLabel.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + staticLabel.EntityData.SegmentPath
    staticLabel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    staticLabel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    staticLabel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    staticLabel.EntityData.Children = types.NewOrderedMap()
    staticLabel.EntityData.Leafs = types.NewOrderedMap()
    staticLabel.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", staticLabel.LocalLabel})
    staticLabel.EntityData.Leafs.Append("remote-label", types.YLeaf{"RemoteLabel", staticLabel.RemoteLabel})

    staticLabel.EntityData.YListKeys = []string {}

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
    YPresence bool

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
    splitHorizonGroup.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + splitHorizonGroup.EntityData.SegmentPath
    splitHorizonGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    splitHorizonGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    splitHorizonGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    splitHorizonGroup.EntityData.Children = types.NewOrderedMap()
    splitHorizonGroup.EntityData.Leafs = types.NewOrderedMap()
    splitHorizonGroup.EntityData.Leafs.Append("id", types.YLeaf{"Id", splitHorizonGroup.Id})

    splitHorizonGroup.EntityData.YListKeys = []string {}

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
    mac.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + mac.EntityData.SegmentPath
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("limit", types.YChild{"Limit", &mac.Limit})
    mac.EntityData.Children.Append("aging", types.YChild{"Aging", &mac.Aging})
    mac.EntityData.Children.Append("port-down", types.YChild{"PortDown", &mac.PortDown})
    mac.EntityData.Children.Append("secure", types.YChild{"Secure", &mac.Secure})
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("learning-enabled", types.YLeaf{"LearningEnabled", mac.LearningEnabled})

    mac.EntityData.YListKeys = []string {}

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
    limit.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/mac/" + limit.EntityData.SegmentPath
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = types.NewOrderedMap()
    limit.EntityData.Leafs = types.NewOrderedMap()
    limit.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", limit.Maximum})
    limit.EntityData.Leafs.Append("action", types.YLeaf{"Action", limit.Action})
    limit.EntityData.Leafs.Append("notification", types.YLeaf{"Notification", limit.Notification})

    limit.EntityData.YListKeys = []string {}

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
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/mac/" + aging.EntityData.SegmentPath
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = types.NewOrderedMap()
    aging.EntityData.Leafs = types.NewOrderedMap()
    aging.EntityData.Leafs.Append("time", types.YLeaf{"Time", aging.Time})
    aging.EntityData.Leafs.Append("type", types.YLeaf{"Type", aging.Type})

    aging.EntityData.YListKeys = []string {}

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
    portDown.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/mac/" + portDown.EntityData.SegmentPath
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = types.NewOrderedMap()
    portDown.EntityData.Leafs = types.NewOrderedMap()
    portDown.EntityData.Leafs.Append("flush", types.YLeaf{"Flush", portDown.Flush})

    portDown.EntityData.YListKeys = []string {}

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
    secure.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/mac/" + secure.EntityData.SegmentPath
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = types.NewOrderedMap()
    secure.EntityData.Leafs = types.NewOrderedMap()
    secure.EntityData.Leafs.Append("action", types.YLeaf{"Action", secure.Action})
    secure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", secure.Logging})
    secure.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", secure.Enabled})

    secure.EntityData.YListKeys = []string {}

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
    igmpSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + igmpSnooping.EntityData.SegmentPath
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", igmpSnooping.ProfileName})

    igmpSnooping.EntityData.YListKeys = []string {}

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
    mldSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + mldSnooping.EntityData.SegmentPath
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", mldSnooping.ProfileName})

    mldSnooping.EntityData.YListKeys = []string {}

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
    dhcpIpv4Snooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + dhcpIpv4Snooping.EntityData.SegmentPath
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName})

    dhcpIpv4Snooping.EntityData.YListKeys = []string {}

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
    flooding.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + flooding.EntityData.SegmentPath
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = types.NewOrderedMap()
    flooding.EntityData.Leafs = types.NewOrderedMap()
    flooding.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", flooding.Disabled})
    flooding.EntityData.Leafs.Append("disabled-unknown-unicast", types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast})

    flooding.EntityData.YListKeys = []string {}

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
    Thresholds []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "pw-neighbor-spec"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + stormControl.EntityData.SegmentPath
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = types.NewOrderedMap()
    stormControl.EntityData.Children.Append("thresholds", types.YChild{"Thresholds", nil})
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children.Append(types.GetSegmentPath(stormControl.Thresholds[i]), types.YChild{"Thresholds", stormControl.Thresholds[i]})
    }
    stormControl.EntityData.Leafs = types.NewOrderedMap()
    stormControl.EntityData.Leafs.Append("action", types.YLeaf{"Action", stormControl.Action})

    stormControl.EntityData.YListKeys = []string {}

    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    thresholds.EntityData.SegmentPath = "thresholds" + types.AddKeyToken(thresholds.TrafficClass, "traffic-class")
    thresholds.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/storm-control/" + thresholds.EntityData.SegmentPath
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = types.NewOrderedMap()
    thresholds.EntityData.Leafs = types.NewOrderedMap()
    thresholds.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", thresholds.TrafficClass})
    thresholds.EntityData.Leafs.Append("value", types.YLeaf{"Value", thresholds.Value})
    thresholds.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", thresholds.Unit})

    thresholds.EntityData.YListKeys = []string {"TrafficClass"}

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
    backup.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/members/access-pw-member/pw-neighbor-spec/" + backup.EntityData.SegmentPath
    backup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    backup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    backup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    backup.EntityData.Children = types.NewOrderedMap()
    backup.EntityData.Leafs = types.NewOrderedMap()
    backup.EntityData.Leafs.Append("neighbor-ip-address", types.YLeaf{"NeighborIpAddress", backup.NeighborIpAddress})
    backup.EntityData.Leafs.Append("vc-id", types.YLeaf{"VcId", backup.VcId})
    backup.EntityData.Leafs.Append("pw-class-template", types.YLeaf{"PwClassTemplate", backup.PwClassTemplate})

    backup.EntityData.YListKeys = []string {}

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
    mac.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + mac.EntityData.SegmentPath
    mac.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mac.EntityData.Children = types.NewOrderedMap()
    mac.EntityData.Children.Append("limit", types.YChild{"Limit", &mac.Limit})
    mac.EntityData.Children.Append("aging", types.YChild{"Aging", &mac.Aging})
    mac.EntityData.Children.Append("port-down", types.YChild{"PortDown", &mac.PortDown})
    mac.EntityData.Children.Append("flooding", types.YChild{"Flooding", &mac.Flooding})
    mac.EntityData.Children.Append("secure", types.YChild{"Secure", &mac.Secure})
    mac.EntityData.Children.Append("static", types.YChild{"Static", &mac.Static})
    mac.EntityData.Leafs = types.NewOrderedMap()
    mac.EntityData.Leafs.Append("learning-enabled", types.YLeaf{"LearningEnabled", mac.LearningEnabled})

    mac.EntityData.YListKeys = []string {}

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
    limit.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + limit.EntityData.SegmentPath
    limit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    limit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    limit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    limit.EntityData.Children = types.NewOrderedMap()
    limit.EntityData.Leafs = types.NewOrderedMap()
    limit.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", limit.Maximum})
    limit.EntityData.Leafs.Append("action", types.YLeaf{"Action", limit.Action})
    limit.EntityData.Leafs.Append("notification", types.YLeaf{"Notification", limit.Notification})

    limit.EntityData.YListKeys = []string {}

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
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetEntityData() *types.CommonEntityData {
    aging.EntityData.YFilter = aging.YFilter
    aging.EntityData.YangName = "aging"
    aging.EntityData.BundleName = "cisco_ios_xe"
    aging.EntityData.ParentYangName = "mac"
    aging.EntityData.SegmentPath = "aging"
    aging.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + aging.EntityData.SegmentPath
    aging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aging.EntityData.Children = types.NewOrderedMap()
    aging.EntityData.Leafs = types.NewOrderedMap()
    aging.EntityData.Leafs.Append("time", types.YLeaf{"Time", aging.Time})
    aging.EntityData.Leafs.Append("type", types.YLeaf{"Type", aging.Type})

    aging.EntityData.YListKeys = []string {}

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
    portDown.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + portDown.EntityData.SegmentPath
    portDown.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    portDown.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    portDown.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    portDown.EntityData.Children = types.NewOrderedMap()
    portDown.EntityData.Leafs = types.NewOrderedMap()
    portDown.EntityData.Leafs.Append("flush", types.YLeaf{"Flush", portDown.Flush})

    portDown.EntityData.YListKeys = []string {}

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
    flooding.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + flooding.EntityData.SegmentPath
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = types.NewOrderedMap()
    flooding.EntityData.Leafs = types.NewOrderedMap()
    flooding.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", flooding.Disabled})
    flooding.EntityData.Leafs.Append("disabled-unknown-unicast", types.YLeaf{"DisabledUnknownUnicast", flooding.DisabledUnknownUnicast})

    flooding.EntityData.YListKeys = []string {}

    return &(flooding.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure
// MAC secure parameters.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    secure.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + secure.EntityData.SegmentPath
    secure.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    secure.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    secure.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    secure.EntityData.Children = types.NewOrderedMap()
    secure.EntityData.Leafs = types.NewOrderedMap()
    secure.EntityData.Leafs.Append("action", types.YLeaf{"Action", secure.Action})
    secure.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", secure.Logging})

    secure.EntityData.YListKeys = []string {}

    return &(secure.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static
// Static mac address list parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address entry. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses.
    MacAddresses []*BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xe"
    static.EntityData.ParentYangName = "mac"
    static.EntityData.SegmentPath = "static"
    static.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/" + static.EntityData.SegmentPath
    static.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    static.EntityData.Children = types.NewOrderedMap()
    static.EntityData.Children.Append("mac-addresses", types.YChild{"MacAddresses", nil})
    for i := range static.MacAddresses {
        static.EntityData.Children.Append(types.GetSegmentPath(static.MacAddresses[i]), types.YChild{"MacAddresses", static.MacAddresses[i]})
    }
    static.EntityData.Leafs = types.NewOrderedMap()

    static.EntityData.YListKeys = []string {}

    return &(static.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
// MAC address entry.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    macAddresses.EntityData.SegmentPath = "mac-addresses" + types.AddKeyToken(macAddresses.MacAddr, "mac-addr")
    macAddresses.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/mac/static/" + macAddresses.EntityData.SegmentPath
    macAddresses.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    macAddresses.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    macAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    macAddresses.EntityData.Children = types.NewOrderedMap()
    macAddresses.EntityData.Leafs = types.NewOrderedMap()
    macAddresses.EntityData.Leafs.Append("mac-addr", types.YLeaf{"MacAddr", macAddresses.MacAddr})
    macAddresses.EntityData.Leafs.Append("drop", types.YLeaf{"Drop", macAddresses.Drop})

    macAddresses.EntityData.YListKeys = []string {"MacAddr"}

    return &(macAddresses.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection
// Dynamic ARP Inspection (DAI) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    dynamicArpInspection.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + dynamicArpInspection.EntityData.SegmentPath
    dynamicArpInspection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dynamicArpInspection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dynamicArpInspection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dynamicArpInspection.EntityData.Children = types.NewOrderedMap()
    dynamicArpInspection.EntityData.Children.Append("address-validation", types.YChild{"AddressValidation", &dynamicArpInspection.AddressValidation})
    dynamicArpInspection.EntityData.Leafs = types.NewOrderedMap()
    dynamicArpInspection.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", dynamicArpInspection.Logging})

    dynamicArpInspection.EntityData.YListKeys = []string {}

    return &(dynamicArpInspection.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    addressValidation.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/dynamic-arp-inspection/" + addressValidation.EntityData.SegmentPath
    addressValidation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressValidation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressValidation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressValidation.EntityData.Children = types.NewOrderedMap()
    addressValidation.EntityData.Leafs = types.NewOrderedMap()
    addressValidation.EntityData.Leafs.Append("dst-mac", types.YLeaf{"DstMac", addressValidation.DstMac})
    addressValidation.EntityData.Leafs.Append("src-mac", types.YLeaf{"SrcMac", addressValidation.SrcMac})
    addressValidation.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addressValidation.Ipv4})

    addressValidation.EntityData.YListKeys = []string {}

    return &(addressValidation.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard
// IP source guard (IPSG) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable IPSG logging. The type is bool. The default value is false.
    Logging interface{}
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetEntityData() *types.CommonEntityData {
    ipSourceGuard.EntityData.YFilter = ipSourceGuard.YFilter
    ipSourceGuard.EntityData.YangName = "ip-source-guard"
    ipSourceGuard.EntityData.BundleName = "cisco_ios_xe"
    ipSourceGuard.EntityData.ParentYangName = "bridge-domain"
    ipSourceGuard.EntityData.SegmentPath = "ip-source-guard"
    ipSourceGuard.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + ipSourceGuard.EntityData.SegmentPath
    ipSourceGuard.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSourceGuard.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSourceGuard.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSourceGuard.EntityData.Children = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs = types.NewOrderedMap()
    ipSourceGuard.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", ipSourceGuard.Logging})

    ipSourceGuard.EntityData.YListKeys = []string {}

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
    Thresholds []*BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "bridge-domain"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + stormControl.EntityData.SegmentPath
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = types.NewOrderedMap()
    stormControl.EntityData.Children.Append("thresholds", types.YChild{"Thresholds", nil})
    for i := range stormControl.Thresholds {
        stormControl.EntityData.Children.Append(types.GetSegmentPath(stormControl.Thresholds[i]), types.YChild{"Thresholds", stormControl.Thresholds[i]})
    }
    stormControl.EntityData.Leafs = types.NewOrderedMap()
    stormControl.EntityData.Leafs.Append("action", types.YLeaf{"Action", stormControl.Action})

    stormControl.EntityData.YListKeys = []string {}

    return &(stormControl.EntityData)
}

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    thresholds.EntityData.SegmentPath = "thresholds" + types.AddKeyToken(thresholds.TrafficClass, "traffic-class")
    thresholds.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/storm-control/" + thresholds.EntityData.SegmentPath
    thresholds.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    thresholds.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    thresholds.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    thresholds.EntityData.Children = types.NewOrderedMap()
    thresholds.EntityData.Leafs = types.NewOrderedMap()
    thresholds.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", thresholds.TrafficClass})
    thresholds.EntityData.Leafs.Append("value", types.YLeaf{"Value", thresholds.Value})
    thresholds.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", thresholds.Unit})

    thresholds.EntityData.YListKeys = []string {"TrafficClass"}

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
    igmpSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + igmpSnooping.EntityData.SegmentPath
    igmpSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igmpSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igmpSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igmpSnooping.EntityData.Children = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs = types.NewOrderedMap()
    igmpSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", igmpSnooping.ProfileName})
    igmpSnooping.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", igmpSnooping.Disabled})

    igmpSnooping.EntityData.YListKeys = []string {}

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
    mldSnooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + mldSnooping.EntityData.SegmentPath
    mldSnooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mldSnooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mldSnooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mldSnooping.EntityData.Children = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs = types.NewOrderedMap()
    mldSnooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", mldSnooping.ProfileName})

    mldSnooping.EntityData.YListKeys = []string {}

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
    dhcpIpv4Snooping.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-config/bridge-domains/bridge-domain/" + dhcpIpv4Snooping.EntityData.SegmentPath
    dhcpIpv4Snooping.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dhcpIpv4Snooping.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dhcpIpv4Snooping.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dhcpIpv4Snooping.EntityData.Children = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs = types.NewOrderedMap()
    dhcpIpv4Snooping.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", dhcpIpv4Snooping.ProfileName})

    dhcpIpv4Snooping.EntityData.YListKeys = []string {}

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
    MacTable []*BridgeDomainState_MacTable
}

func (bridgeDomainState *BridgeDomainState) GetEntityData() *types.CommonEntityData {
    bridgeDomainState.EntityData.YFilter = bridgeDomainState.YFilter
    bridgeDomainState.EntityData.YangName = "bridge-domain-state"
    bridgeDomainState.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomainState.EntityData.ParentYangName = "cisco-bridge-domain"
    bridgeDomainState.EntityData.SegmentPath = "cisco-bridge-domain:bridge-domain-state"
    bridgeDomainState.EntityData.AbsolutePath = bridgeDomainState.EntityData.SegmentPath
    bridgeDomainState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomainState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomainState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomainState.EntityData.Children = types.NewOrderedMap()
    bridgeDomainState.EntityData.Children.Append("system-capabilities", types.YChild{"SystemCapabilities", &bridgeDomainState.SystemCapabilities})
    bridgeDomainState.EntityData.Children.Append("module-capabilities", types.YChild{"ModuleCapabilities", &bridgeDomainState.ModuleCapabilities})
    bridgeDomainState.EntityData.Children.Append("bridge-domains", types.YChild{"BridgeDomains", &bridgeDomainState.BridgeDomains})
    bridgeDomainState.EntityData.Children.Append("mac-table", types.YChild{"MacTable", nil})
    for i := range bridgeDomainState.MacTable {
        bridgeDomainState.EntityData.Children.Append(types.GetSegmentPath(bridgeDomainState.MacTable[i]), types.YChild{"MacTable", bridgeDomainState.MacTable[i]})
    }
    bridgeDomainState.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomainState.EntityData.YListKeys = []string {}

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
    systemCapabilities.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/" + systemCapabilities.EntityData.SegmentPath
    systemCapabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    systemCapabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    systemCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    systemCapabilities.EntityData.Children = types.NewOrderedMap()
    systemCapabilities.EntityData.Leafs = types.NewOrderedMap()
    systemCapabilities.EntityData.Leafs.Append("max-bd", types.YLeaf{"MaxBd", systemCapabilities.MaxBd})
    systemCapabilities.EntityData.Leafs.Append("max-ac-per-bd", types.YLeaf{"MaxAcPerBd", systemCapabilities.MaxAcPerBd})
    systemCapabilities.EntityData.Leafs.Append("max-pw-per-bd", types.YLeaf{"MaxPwPerBd", systemCapabilities.MaxPwPerBd})
    systemCapabilities.EntityData.Leafs.Append("max-vfi-per-bd", types.YLeaf{"MaxVfiPerBd", systemCapabilities.MaxVfiPerBd})
    systemCapabilities.EntityData.Leafs.Append("max-sh-group-per-bd", types.YLeaf{"MaxShGroupPerBd", systemCapabilities.MaxShGroupPerBd})
    systemCapabilities.EntityData.Leafs.Append("max-interflex-if-per-bd", types.YLeaf{"MaxInterflexIfPerBd", systemCapabilities.MaxInterflexIfPerBd})

    systemCapabilities.EntityData.YListKeys = []string {}

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
    Modules []*BridgeDomainState_ModuleCapabilities_Modules
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetEntityData() *types.CommonEntityData {
    moduleCapabilities.EntityData.YFilter = moduleCapabilities.YFilter
    moduleCapabilities.EntityData.YangName = "module-capabilities"
    moduleCapabilities.EntityData.BundleName = "cisco_ios_xe"
    moduleCapabilities.EntityData.ParentYangName = "bridge-domain-state"
    moduleCapabilities.EntityData.SegmentPath = "module-capabilities"
    moduleCapabilities.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/" + moduleCapabilities.EntityData.SegmentPath
    moduleCapabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    moduleCapabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    moduleCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    moduleCapabilities.EntityData.Children = types.NewOrderedMap()
    moduleCapabilities.EntityData.Children.Append("modules", types.YChild{"Modules", nil})
    for i := range moduleCapabilities.Modules {
        moduleCapabilities.EntityData.Children.Append(types.GetSegmentPath(moduleCapabilities.Modules[i]), types.YChild{"Modules", moduleCapabilities.Modules[i]})
    }
    moduleCapabilities.EntityData.Leafs = types.NewOrderedMap()

    moduleCapabilities.EntityData.YListKeys = []string {}

    return &(moduleCapabilities.EntityData)
}

// BridgeDomainState_ModuleCapabilities_Modules
// Collection of capabillity statements for hardware
// module in the system.
type BridgeDomainState_ModuleCapabilities_Modules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    modules.EntityData.SegmentPath = "modules" + types.AddKeyToken(modules.Name, "name")
    modules.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/module-capabilities/" + modules.EntityData.SegmentPath
    modules.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    modules.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    modules.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    modules.EntityData.Children = types.NewOrderedMap()
    modules.EntityData.Leafs = types.NewOrderedMap()
    modules.EntityData.Leafs.Append("name", types.YLeaf{"Name", modules.Name})
    modules.EntityData.Leafs.Append("max-mac-per-bd", types.YLeaf{"MaxMacPerBd", modules.MaxMacPerBd})
    modules.EntityData.Leafs.Append("max-pdd-edge-bd", types.YLeaf{"MaxPddEdgeBd", modules.MaxPddEdgeBd})
    modules.EntityData.Leafs.Append("max-bd", types.YLeaf{"MaxBd", modules.MaxBd})
    modules.EntityData.Leafs.Append("max-ac-per-bd", types.YLeaf{"MaxAcPerBd", modules.MaxAcPerBd})
    modules.EntityData.Leafs.Append("max-pw-per-bd", types.YLeaf{"MaxPwPerBd", modules.MaxPwPerBd})
    modules.EntityData.Leafs.Append("max-vfi-per-bd", types.YLeaf{"MaxVfiPerBd", modules.MaxVfiPerBd})
    modules.EntityData.Leafs.Append("max-sh-group-per-bd", types.YLeaf{"MaxShGroupPerBd", modules.MaxShGroupPerBd})

    modules.EntityData.YListKeys = []string {"Name"}

    return &(modules.EntityData)
}

// BridgeDomainState_BridgeDomains
// Bridge domain state data.
type BridgeDomainState_BridgeDomains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Collection of bridge-domain state data. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain.
    BridgeDomain []*BridgeDomainState_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetEntityData() *types.CommonEntityData {
    bridgeDomains.EntityData.YFilter = bridgeDomains.YFilter
    bridgeDomains.EntityData.YangName = "bridge-domains"
    bridgeDomains.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomains.EntityData.ParentYangName = "bridge-domain-state"
    bridgeDomains.EntityData.SegmentPath = "bridge-domains"
    bridgeDomains.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/" + bridgeDomains.EntityData.SegmentPath
    bridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomains.EntityData.Children = types.NewOrderedMap()
    bridgeDomains.EntityData.Children.Append("bridge-domain", types.YChild{"BridgeDomain", nil})
    for i := range bridgeDomains.BridgeDomain {
        bridgeDomains.EntityData.Children.Append(types.GetSegmentPath(bridgeDomains.BridgeDomain[i]), types.YChild{"BridgeDomain", bridgeDomains.BridgeDomain[i]})
    }
    bridgeDomains.EntityData.Leafs = types.NewOrderedMap()

    bridgeDomains.EntityData.YListKeys = []string {}

    return &(bridgeDomains.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain
// Collection of bridge-domain state data.
type BridgeDomainState_BridgeDomains_BridgeDomain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    P2mpPwDisabled interface{}

    // Collection of bridge-domain members.
    Members BridgeDomainState_BridgeDomains_BridgeDomain_Members
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetEntityData() *types.CommonEntityData {
    bridgeDomain.EntityData.YFilter = bridgeDomain.YFilter
    bridgeDomain.EntityData.YangName = "bridge-domain"
    bridgeDomain.EntityData.BundleName = "cisco_ios_xe"
    bridgeDomain.EntityData.ParentYangName = "bridge-domains"
    bridgeDomain.EntityData.SegmentPath = "bridge-domain" + types.AddKeyToken(bridgeDomain.Id, "id")
    bridgeDomain.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/" + bridgeDomain.EntityData.SegmentPath
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = types.NewOrderedMap()
    bridgeDomain.EntityData.Children.Append("members", types.YChild{"Members", &bridgeDomain.Members})
    bridgeDomain.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs.Append("id", types.YLeaf{"Id", bridgeDomain.Id})
    bridgeDomain.EntityData.Leafs.Append("bd-state", types.YLeaf{"BdState", bridgeDomain.BdState})
    bridgeDomain.EntityData.Leafs.Append("create-time", types.YLeaf{"CreateTime", bridgeDomain.CreateTime})
    bridgeDomain.EntityData.Leafs.Append("last-status-change", types.YLeaf{"LastStatusChange", bridgeDomain.LastStatusChange})
    bridgeDomain.EntityData.Leafs.Append("mac-limit-reached", types.YLeaf{"MacLimitReached", bridgeDomain.MacLimitReached})
    bridgeDomain.EntityData.Leafs.Append("p2mp-pw-disabled", types.YLeaf{"P2mpPwDisabled", bridgeDomain.P2mpPwDisabled})

    bridgeDomain.EntityData.YListKeys = []string {"Id"}

    return &(bridgeDomain.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of attachment circuits for this bridge domains. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember.
    AcMember []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember

    // Reference to an instance of Bridge domain Virtual Forwarding Instance (VFI)
    // name. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember.
    VfiMember []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember

    // Collection of access pseudowire members of the bridge domain. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember.
    AccessPwMember []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetEntityData() *types.CommonEntityData {
    members.EntityData.YFilter = members.YFilter
    members.EntityData.YangName = "members"
    members.EntityData.BundleName = "cisco_ios_xe"
    members.EntityData.ParentYangName = "bridge-domain"
    members.EntityData.SegmentPath = "members"
    members.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/" + members.EntityData.SegmentPath
    members.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    members.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    members.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    members.EntityData.Children = types.NewOrderedMap()
    members.EntityData.Children.Append("ac-member", types.YChild{"AcMember", nil})
    for i := range members.AcMember {
        members.EntityData.Children.Append(types.GetSegmentPath(members.AcMember[i]), types.YChild{"AcMember", members.AcMember[i]})
    }
    members.EntityData.Children.Append("vfi-member", types.YChild{"VfiMember", nil})
    for i := range members.VfiMember {
        members.EntityData.Children.Append(types.GetSegmentPath(members.VfiMember[i]), types.YChild{"VfiMember", members.VfiMember[i]})
    }
    members.EntityData.Children.Append("access-pw-member", types.YChild{"AccessPwMember", nil})
    for i := range members.AccessPwMember {
        members.EntityData.Children.Append(types.GetSegmentPath(members.AccessPwMember[i]), types.YChild{"AccessPwMember", members.AccessPwMember[i]})
    }
    members.EntityData.Leafs = types.NewOrderedMap()

    members.EntityData.YListKeys = []string {}

    return &(members.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember
// List of attachment circuits for this bridge domains
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to an instance of Bridge domain
    // attachment circuit (AC) interface name. The type is string. Refers to
    // ietf_interfaces.InterfacesState_Interface_Name
    Interface interface{}

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
    acMember.EntityData.SegmentPath = "ac-member" + types.AddKeyToken(acMember.Interface, "interface")
    acMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/" + acMember.EntityData.SegmentPath
    acMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    acMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    acMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    acMember.EntityData.Children = types.NewOrderedMap()
    acMember.EntityData.Children.Append("dai-stats", types.YChild{"DaiStats", &acMember.DaiStats})
    acMember.EntityData.Children.Append("ipsg-stats", types.YChild{"IpsgStats", &acMember.IpsgStats})
    acMember.EntityData.Children.Append("storm-control", types.YChild{"StormControl", &acMember.StormControl})
    acMember.EntityData.Leafs = types.NewOrderedMap()
    acMember.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", acMember.Interface})
    acMember.EntityData.Leafs.Append("static-mac-count", types.YLeaf{"StaticMacCount", acMember.StaticMacCount})

    acMember.EntityData.YListKeys = []string {"Interface"}

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
    daiStats.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/ac-member/" + daiStats.EntityData.SegmentPath
    daiStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    daiStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    daiStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    daiStats.EntityData.Children = types.NewOrderedMap()
    daiStats.EntityData.Leafs = types.NewOrderedMap()
    daiStats.EntityData.Leafs.Append("packet-drops", types.YLeaf{"PacketDrops", daiStats.PacketDrops})
    daiStats.EntityData.Leafs.Append("byte-drops", types.YLeaf{"ByteDrops", daiStats.ByteDrops})

    daiStats.EntityData.YListKeys = []string {}

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
    ipsgStats.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/ac-member/" + ipsgStats.EntityData.SegmentPath
    ipsgStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsgStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsgStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsgStats.EntityData.Children = types.NewOrderedMap()
    ipsgStats.EntityData.Leafs = types.NewOrderedMap()
    ipsgStats.EntityData.Leafs.Append("packet-drops", types.YLeaf{"PacketDrops", ipsgStats.PacketDrops})
    ipsgStats.EntityData.Leafs.Append("byte-drops", types.YLeaf{"ByteDrops", ipsgStats.ByteDrops})

    ipsgStats.EntityData.YListKeys = []string {}

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
    DropCounter []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetEntityData() *types.CommonEntityData {
    stormControl.EntityData.YFilter = stormControl.YFilter
    stormControl.EntityData.YangName = "storm-control"
    stormControl.EntityData.BundleName = "cisco_ios_xe"
    stormControl.EntityData.ParentYangName = "ac-member"
    stormControl.EntityData.SegmentPath = "storm-control"
    stormControl.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/ac-member/" + stormControl.EntityData.SegmentPath
    stormControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stormControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stormControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stormControl.EntityData.Children = types.NewOrderedMap()
    stormControl.EntityData.Children.Append("drop-counter", types.YChild{"DropCounter", nil})
    for i := range stormControl.DropCounter {
        stormControl.EntityData.Children.Append(types.GetSegmentPath(stormControl.DropCounter[i]), types.YChild{"DropCounter", stormControl.DropCounter[i]})
    }
    stormControl.EntityData.Leafs = types.NewOrderedMap()

    stormControl.EntityData.YListKeys = []string {}

    return &(stormControl.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
// Collection of packet drop statistics for ethernet traffic
// clasess.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    dropCounter.EntityData.SegmentPath = "drop-counter" + types.AddKeyToken(dropCounter.TrafficClass, "traffic-class")
    dropCounter.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/ac-member/storm-control/" + dropCounter.EntityData.SegmentPath
    dropCounter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dropCounter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dropCounter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dropCounter.EntityData.Children = types.NewOrderedMap()
    dropCounter.EntityData.Leafs = types.NewOrderedMap()
    dropCounter.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", dropCounter.TrafficClass})
    dropCounter.EntityData.Leafs.Append("packet-drops", types.YLeaf{"PacketDrops", dropCounter.PacketDrops})
    dropCounter.EntityData.Leafs.Append("octate-drops", types.YLeaf{"OctateDrops", dropCounter.OctateDrops})

    dropCounter.EntityData.YListKeys = []string {"TrafficClass"}

    return &(dropCounter.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember
// Reference to an instance of Bridge domain Virtual
// Forwarding Instance (VFI) name.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Bridge domain memeber interface name. The type is
    // string. Refers to ietf_interfaces.InterfacesState_Interface_Name
    Interface interface{}

    // Flooding operational status.
    Flooding BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetEntityData() *types.CommonEntityData {
    vfiMember.EntityData.YFilter = vfiMember.YFilter
    vfiMember.EntityData.YangName = "vfi-member"
    vfiMember.EntityData.BundleName = "cisco_ios_xe"
    vfiMember.EntityData.ParentYangName = "members"
    vfiMember.EntityData.SegmentPath = "vfi-member" + types.AddKeyToken(vfiMember.Interface, "interface")
    vfiMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/" + vfiMember.EntityData.SegmentPath
    vfiMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vfiMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vfiMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vfiMember.EntityData.Children = types.NewOrderedMap()
    vfiMember.EntityData.Children.Append("flooding", types.YChild{"Flooding", &vfiMember.Flooding})
    vfiMember.EntityData.Leafs = types.NewOrderedMap()
    vfiMember.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", vfiMember.Interface})

    vfiMember.EntityData.YListKeys = []string {"Interface"}

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
    Status []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "vfi-member"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/vfi-member/" + flooding.EntityData.SegmentPath
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = types.NewOrderedMap()
    flooding.EntityData.Children.Append("status", types.YChild{"Status", nil})
    for i := range flooding.Status {
        flooding.EntityData.Children.Append(types.GetSegmentPath(flooding.Status[i]), types.YChild{"Status", flooding.Status[i]})
    }
    flooding.EntityData.Leafs = types.NewOrderedMap()

    flooding.EntityData.YListKeys = []string {}

    return &(flooding.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    status.EntityData.SegmentPath = "status" + types.AddKeyToken(status.TrafficClass, "traffic-class")
    status.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/vfi-member/flooding/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", status.TrafficClass})
    status.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", status.Enabled})

    status.EntityData.YListKeys = []string {"TrafficClass"}

    return &(status.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember
// Collection of access pseudowire members of the bridge
// domain.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    accessPwMember.EntityData.SegmentPath = "access-pw-member" + types.AddKeyToken(accessPwMember.VcPeerAddress, "vc-peer-address") + types.AddKeyToken(accessPwMember.VcId, "vc-id")
    accessPwMember.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/" + accessPwMember.EntityData.SegmentPath
    accessPwMember.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accessPwMember.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accessPwMember.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accessPwMember.EntityData.Children = types.NewOrderedMap()
    accessPwMember.EntityData.Children.Append("flooding", types.YChild{"Flooding", &accessPwMember.Flooding})
    accessPwMember.EntityData.Leafs = types.NewOrderedMap()
    accessPwMember.EntityData.Leafs.Append("vc-peer-address", types.YLeaf{"VcPeerAddress", accessPwMember.VcPeerAddress})
    accessPwMember.EntityData.Leafs.Append("vc-id", types.YLeaf{"VcId", accessPwMember.VcId})

    accessPwMember.EntityData.YListKeys = []string {"VcPeerAddress", "VcId"}

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
    Status []*BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetEntityData() *types.CommonEntityData {
    flooding.EntityData.YFilter = flooding.YFilter
    flooding.EntityData.YangName = "flooding"
    flooding.EntityData.BundleName = "cisco_ios_xe"
    flooding.EntityData.ParentYangName = "access-pw-member"
    flooding.EntityData.SegmentPath = "flooding"
    flooding.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/access-pw-member/" + flooding.EntityData.SegmentPath
    flooding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    flooding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    flooding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    flooding.EntityData.Children = types.NewOrderedMap()
    flooding.EntityData.Children.Append("status", types.YChild{"Status", nil})
    for i := range flooding.Status {
        flooding.EntityData.Children.Append(types.GetSegmentPath(flooding.Status[i]), types.YChild{"Status", flooding.Status[i]})
    }
    flooding.EntityData.Leafs = types.NewOrderedMap()

    flooding.EntityData.YListKeys = []string {}

    return &(flooding.EntityData)
}

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    status.EntityData.SegmentPath = "status" + types.AddKeyToken(status.TrafficClass, "traffic-class")
    status.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/bridge-domains/bridge-domain/members/access-pw-member/flooding/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("traffic-class", types.YLeaf{"TrafficClass", status.TrafficClass})
    status.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", status.Enabled})

    status.EntityData.YListKeys = []string {"TrafficClass"}

    return &(status.EntityData)
}

// BridgeDomainState_MacTable
// This list contains mac-address entries for bridge
// domains.
type BridgeDomainState_MacTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    Interface interface{}

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
    macTable.EntityData.SegmentPath = "mac-table" + types.AddKeyToken(macTable.BdId, "bd-id") + types.AddKeyToken(macTable.MacAddress, "mac-address")
    macTable.EntityData.AbsolutePath = "cisco-bridge-domain:bridge-domain-state/" + macTable.EntityData.SegmentPath
    macTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    macTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    macTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    macTable.EntityData.Children = types.NewOrderedMap()
    macTable.EntityData.Leafs = types.NewOrderedMap()
    macTable.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", macTable.BdId})
    macTable.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", macTable.MacAddress})
    macTable.EntityData.Leafs.Append("mac-type", types.YLeaf{"MacType", macTable.MacType})
    macTable.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", macTable.Interface})
    macTable.EntityData.Leafs.Append("secure-mac", types.YLeaf{"SecureMac", macTable.SecureMac})
    macTable.EntityData.Leafs.Append("ntfy-mac", types.YLeaf{"NtfyMac", macTable.NtfyMac})
    macTable.EntityData.Leafs.Append("age", types.YLeaf{"Age", macTable.Age})
    macTable.EntityData.Leafs.Append("location", types.YLeaf{"Location", macTable.Location})

    macTable.EntityData.YListKeys = []string {"BdId", "MacAddress"}

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
    clearBridgeDomain.EntityData.AbsolutePath = clearBridgeDomain.EntityData.SegmentPath
    clearBridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearBridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearBridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearBridgeDomain.EntityData.Children = types.NewOrderedMap()
    clearBridgeDomain.EntityData.Children.Append("input", types.YChild{"Input", &clearBridgeDomain.Input})
    clearBridgeDomain.EntityData.Children.Append("output", types.YChild{"Output", &clearBridgeDomain.Output})
    clearBridgeDomain.EntityData.Leafs = types.NewOrderedMap()

    clearBridgeDomain.EntityData.YListKeys = []string {}

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
    input.EntityData.AbsolutePath = "cisco-bridge-domain:clear-bridge-domain/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("all", types.YLeaf{"All", input.All})
    input.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", input.BdId})
    input.EntityData.Leafs.Append("bg-id", types.YLeaf{"BgId", input.BgId})

    input.EntityData.YListKeys = []string {}

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
    output.EntityData.AbsolutePath = "cisco-bridge-domain:clear-bridge-domain/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("errstr", types.YLeaf{"Errstr", output.Errstr})

    output.EntityData.YListKeys = []string {}

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
    clearMacAddress.EntityData.AbsolutePath = clearMacAddress.EntityData.SegmentPath
    clearMacAddress.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearMacAddress.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearMacAddress.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearMacAddress.EntityData.Children = types.NewOrderedMap()
    clearMacAddress.EntityData.Children.Append("input", types.YChild{"Input", &clearMacAddress.Input})
    clearMacAddress.EntityData.Children.Append("output", types.YChild{"Output", &clearMacAddress.Output})
    clearMacAddress.EntityData.Leafs = types.NewOrderedMap()

    clearMacAddress.EntityData.YListKeys = []string {}

    return &(clearMacAddress.EntityData)
}

// ClearMacAddress_Input
type ClearMacAddress_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reference to an interface. Clear mac-addesses learnt by by this interface.
    // The type is string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

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
    input.EntityData.AbsolutePath = "cisco-bridge-domain:clear-mac-address/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("bridge-domain", types.YChild{"BridgeDomain", &input.BridgeDomain})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", input.Interface})
    input.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", input.MacAddress})

    input.EntityData.YListKeys = []string {}

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
    bridgeDomain.EntityData.AbsolutePath = "cisco-bridge-domain:clear-mac-address/input/" + bridgeDomain.EntityData.SegmentPath
    bridgeDomain.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bridgeDomain.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bridgeDomain.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bridgeDomain.EntityData.Children = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs = types.NewOrderedMap()
    bridgeDomain.EntityData.Leafs.Append("bd-id", types.YLeaf{"BdId", bridgeDomain.BdId})
    bridgeDomain.EntityData.Leafs.Append("bg-id", types.YLeaf{"BgId", bridgeDomain.BgId})

    bridgeDomain.EntityData.YListKeys = []string {}

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
    output.EntityData.AbsolutePath = "cisco-bridge-domain:clear-mac-address/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("errstr", types.YLeaf{"Errstr", output.Errstr})

    output.EntityData.YListKeys = []string {}

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
    createParameterizedBridgeDomains.EntityData.AbsolutePath = createParameterizedBridgeDomains.EntityData.SegmentPath
    createParameterizedBridgeDomains.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    createParameterizedBridgeDomains.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    createParameterizedBridgeDomains.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    createParameterizedBridgeDomains.EntityData.Children = types.NewOrderedMap()
    createParameterizedBridgeDomains.EntityData.Children.Append("input", types.YChild{"Input", &createParameterizedBridgeDomains.Input})
    createParameterizedBridgeDomains.EntityData.Children.Append("output", types.YChild{"Output", &createParameterizedBridgeDomains.Output})
    createParameterizedBridgeDomains.EntityData.Leafs = types.NewOrderedMap()

    createParameterizedBridgeDomains.EntityData.YListKeys = []string {}

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
    Member []*CreateParameterizedBridgeDomains_Input_Member
}

func (input *CreateParameterizedBridgeDomains_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "create-parameterized-bridge-domains"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "cisco-bridge-domain:create-parameterized-bridge-domains/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("member", types.YChild{"Member", nil})
    for i := range input.Member {
        input.EntityData.Children.Append(types.GetSegmentPath(input.Member[i]), types.YChild{"Member", input.Member[i]})
    }
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("parameter", types.YLeaf{"Parameter", input.Parameter})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CreateParameterizedBridgeDomains_Input_Member
// Bridge-domain member interface.
type CreateParameterizedBridgeDomains_Input_Member struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Reference to an interface instance which is
    // configured to be part of this bridge domain. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetEntityData() *types.CommonEntityData {
    member.EntityData.YFilter = member.YFilter
    member.EntityData.YangName = "member"
    member.EntityData.BundleName = "cisco_ios_xe"
    member.EntityData.ParentYangName = "input"
    member.EntityData.SegmentPath = "member" + types.AddKeyToken(member.Interface, "interface")
    member.EntityData.AbsolutePath = "cisco-bridge-domain:create-parameterized-bridge-domains/input/" + member.EntityData.SegmentPath
    member.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    member.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    member.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    member.EntityData.Children = types.NewOrderedMap()
    member.EntityData.Leafs = types.NewOrderedMap()
    member.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", member.Interface})

    member.EntityData.YListKeys = []string {"Interface"}

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
    output.EntityData.AbsolutePath = "cisco-bridge-domain:create-parameterized-bridge-domains/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("errstr", types.YLeaf{"Errstr", output.Errstr})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

