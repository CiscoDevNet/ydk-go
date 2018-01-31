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
    parent types.Entity
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

func (bridgeDomainConfig *BridgeDomainConfig) GetFilter() yfilter.YFilter { return bridgeDomainConfig.YFilter }

func (bridgeDomainConfig *BridgeDomainConfig) SetFilter(yf yfilter.YFilter) { bridgeDomainConfig.YFilter = yf }

func (bridgeDomainConfig *BridgeDomainConfig) GetGoName(yname string) string {
    if yname == "global" { return "Global" }
    if yname == "bridge-groups" { return "BridgeGroups" }
    if yname == "bridge-domains" { return "BridgeDomains" }
    return ""
}

func (bridgeDomainConfig *BridgeDomainConfig) GetSegmentPath() string {
    return "cisco-bridge-domain:bridge-domain-config"
}

func (bridgeDomainConfig *BridgeDomainConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global" {
        return &bridgeDomainConfig.Global
    }
    if childYangName == "bridge-groups" {
        return &bridgeDomainConfig.BridgeGroups
    }
    if childYangName == "bridge-domains" {
        return &bridgeDomainConfig.BridgeDomains
    }
    return nil
}

func (bridgeDomainConfig *BridgeDomainConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global"] = &bridgeDomainConfig.Global
    children["bridge-groups"] = &bridgeDomainConfig.BridgeGroups
    children["bridge-domains"] = &bridgeDomainConfig.BridgeDomains
    return children
}

func (bridgeDomainConfig *BridgeDomainConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bridgeDomainConfig *BridgeDomainConfig) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomainConfig *BridgeDomainConfig) GetYangName() string { return "bridge-domain-config" }

func (bridgeDomainConfig *BridgeDomainConfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomainConfig *BridgeDomainConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomainConfig *BridgeDomainConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomainConfig *BridgeDomainConfig) SetParent(parent types.Entity) { bridgeDomainConfig.parent = parent }

func (bridgeDomainConfig *BridgeDomainConfig) GetParent() types.Entity { return bridgeDomainConfig.parent }

func (bridgeDomainConfig *BridgeDomainConfig) GetParentYangName() string { return "cisco-bridge-domain" }

// BridgeDomainConfig_Global
// Global configurations for bridge-domains.
type BridgeDomainConfig_Global struct {
    parent types.Entity
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

func (global *BridgeDomainConfig_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *BridgeDomainConfig_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *BridgeDomainConfig_Global) GetGoName(yname string) string {
    if yname == "bd-state-notification-enabled" { return "BdStateNotificationEnabled" }
    if yname == "bd-state-notification-rate" { return "BdStateNotificationRate" }
    if yname == "pbb" { return "Pbb" }
    return ""
}

func (global *BridgeDomainConfig_Global) GetSegmentPath() string {
    return "global"
}

func (global *BridgeDomainConfig_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbb" {
        return &global.Pbb
    }
    return nil
}

func (global *BridgeDomainConfig_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbb"] = &global.Pbb
    return children
}

func (global *BridgeDomainConfig_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bd-state-notification-enabled"] = global.BdStateNotificationEnabled
    leafs["bd-state-notification-rate"] = global.BdStateNotificationRate
    return leafs
}

func (global *BridgeDomainConfig_Global) GetBundleName() string { return "cisco_ios_xe" }

func (global *BridgeDomainConfig_Global) GetYangName() string { return "global" }

func (global *BridgeDomainConfig_Global) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (global *BridgeDomainConfig_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (global *BridgeDomainConfig_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (global *BridgeDomainConfig_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *BridgeDomainConfig_Global) GetParent() types.Entity { return global.parent }

func (global *BridgeDomainConfig_Global) GetParentYangName() string { return "bridge-domain-config" }

// BridgeDomainConfig_Global_Pbb
// Provider Backbone Bridging (PBB) related global
// configurations.
type BridgeDomainConfig_Global_Pbb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Backbone source mac address configuration for Provider Backbone Bridging
    // (PBB). The type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    BackboneSrcMac interface{}
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetFilter() yfilter.YFilter { return pbb.YFilter }

func (pbb *BridgeDomainConfig_Global_Pbb) SetFilter(yf yfilter.YFilter) { pbb.YFilter = yf }

func (pbb *BridgeDomainConfig_Global_Pbb) GetGoName(yname string) string {
    if yname == "backbone-src-mac" { return "BackboneSrcMac" }
    return ""
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetSegmentPath() string {
    return "pbb"
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["backbone-src-mac"] = pbb.BackboneSrcMac
    return leafs
}

func (pbb *BridgeDomainConfig_Global_Pbb) GetBundleName() string { return "cisco_ios_xe" }

func (pbb *BridgeDomainConfig_Global_Pbb) GetYangName() string { return "pbb" }

func (pbb *BridgeDomainConfig_Global_Pbb) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pbb *BridgeDomainConfig_Global_Pbb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pbb *BridgeDomainConfig_Global_Pbb) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pbb *BridgeDomainConfig_Global_Pbb) SetParent(parent types.Entity) { pbb.parent = parent }

func (pbb *BridgeDomainConfig_Global_Pbb) GetParent() types.Entity { return pbb.parent }

func (pbb *BridgeDomainConfig_Global_Pbb) GetParentYangName() string { return "global" }

// BridgeDomainConfig_BridgeGroups
// Collection of bridge-groups.
// 
// A Bridge-group is logical grouping construct for bridge
// domains. It defines grouping of bridge-domains under a
// named bridge-group. For example all bridge-domains
// belonging to a single customer can be grouped under a bridge
// -group
type BridgeDomainConfig_BridgeGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bridge-group configuration. The type is slice of
    // BridgeDomainConfig_BridgeGroups_BridgeGroup.
    BridgeGroup []BridgeDomainConfig_BridgeGroups_BridgeGroup
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetFilter() yfilter.YFilter { return bridgeGroups.YFilter }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) SetFilter(yf yfilter.YFilter) { bridgeGroups.YFilter = yf }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetGoName(yname string) string {
    if yname == "bridge-group" { return "BridgeGroup" }
    return ""
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetSegmentPath() string {
    return "bridge-groups"
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bridge-group" {
        for _, c := range bridgeGroups.BridgeGroup {
            if bridgeGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeGroups_BridgeGroup{}
        bridgeGroups.BridgeGroup = append(bridgeGroups.BridgeGroup, child)
        return &bridgeGroups.BridgeGroup[len(bridgeGroups.BridgeGroup)-1]
    }
    return nil
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bridgeGroups.BridgeGroup {
        children[bridgeGroups.BridgeGroup[i].GetSegmentPath()] = &bridgeGroups.BridgeGroup[i]
    }
    return children
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetYangName() string { return "bridge-groups" }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) SetParent(parent types.Entity) { bridgeGroups.parent = parent }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetParent() types.Entity { return bridgeGroups.parent }

func (bridgeGroups *BridgeDomainConfig_BridgeGroups) GetParentYangName() string { return "bridge-domain-config" }

// BridgeDomainConfig_BridgeGroups_BridgeGroup
// Bridge-group configuration.
type BridgeDomainConfig_BridgeGroups_BridgeGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge-group name. The type is string with length:
    // 1..32.
    Name interface{}
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetFilter() yfilter.YFilter { return bridgeGroup.YFilter }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) SetFilter(yf yfilter.YFilter) { bridgeGroup.YFilter = yf }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetSegmentPath() string {
    return "bridge-group" + "[name='" + fmt.Sprintf("%v", bridgeGroup.Name) + "']"
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = bridgeGroup.Name
    return leafs
}

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetYangName() string { return "bridge-group" }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) SetParent(parent types.Entity) { bridgeGroup.parent = parent }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetParent() types.Entity { return bridgeGroup.parent }

func (bridgeGroup *BridgeDomainConfig_BridgeGroups_BridgeGroup) GetParentYangName() string { return "bridge-groups" }

// BridgeDomainConfig_BridgeDomains
// Collection of bridge domains.
type BridgeDomainConfig_BridgeDomains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Definition of a bridge-domain. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain.
    BridgeDomain []BridgeDomainConfig_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetFilter() yfilter.YFilter { return bridgeDomains.YFilter }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) SetFilter(yf yfilter.YFilter) { bridgeDomains.YFilter = yf }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetGoName(yname string) string {
    if yname == "bridge-domain" { return "BridgeDomain" }
    return ""
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetSegmentPath() string {
    return "bridge-domains"
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bridge-domain" {
        for _, c := range bridgeDomains.BridgeDomain {
            if bridgeDomains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain{}
        bridgeDomains.BridgeDomain = append(bridgeDomains.BridgeDomain, child)
        return &bridgeDomains.BridgeDomain[len(bridgeDomains.BridgeDomain)-1]
    }
    return nil
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bridgeDomains.BridgeDomain {
        children[bridgeDomains.BridgeDomain[i].GetSegmentPath()] = &bridgeDomains.BridgeDomain[i]
    }
    return children
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetYangName() string { return "bridge-domains" }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) SetParent(parent types.Entity) { bridgeDomains.parent = parent }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetParent() types.Entity { return bridgeDomains.parent }

func (bridgeDomains *BridgeDomainConfig_BridgeDomains) GetParentYangName() string { return "bridge-domain-config" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain
// Definition of a bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain struct {
    parent types.Entity
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

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetFilter() yfilter.YFilter { return bridgeDomain.YFilter }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) SetFilter(yf yfilter.YFilter) { bridgeDomain.YFilter = yf }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "bridge-group" { return "BridgeGroup" }
    if yname == "enabled" { return "Enabled" }
    if yname == "bd-status-change-notification" { return "BdStatusChangeNotification" }
    if yname == "mtu" { return "Mtu" }
    if yname == "flooding-mode" { return "FloodingMode" }
    if yname == "members" { return "Members" }
    if yname == "mac" { return "Mac" }
    if yname == "dynamic-arp-inspection" { return "DynamicArpInspection" }
    if yname == "ip-source-guard" { return "IpSourceGuard" }
    if yname == "storm-control" { return "StormControl" }
    if yname == "igmp-snooping" { return "IgmpSnooping" }
    if yname == "mld-snooping" { return "MldSnooping" }
    if yname == "dhcp-ipv4-snooping" { return "DhcpIpv4Snooping" }
    return ""
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetSegmentPath() string {
    return "bridge-domain" + "[id='" + fmt.Sprintf("%v", bridgeDomain.Id) + "']"
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "members" {
        return &bridgeDomain.Members
    }
    if childYangName == "mac" {
        return &bridgeDomain.Mac
    }
    if childYangName == "dynamic-arp-inspection" {
        return &bridgeDomain.DynamicArpInspection
    }
    if childYangName == "ip-source-guard" {
        return &bridgeDomain.IpSourceGuard
    }
    if childYangName == "storm-control" {
        return &bridgeDomain.StormControl
    }
    if childYangName == "igmp-snooping" {
        return &bridgeDomain.IgmpSnooping
    }
    if childYangName == "mld-snooping" {
        return &bridgeDomain.MldSnooping
    }
    if childYangName == "dhcp-ipv4-snooping" {
        return &bridgeDomain.DhcpIpv4Snooping
    }
    return nil
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["members"] = &bridgeDomain.Members
    children["mac"] = &bridgeDomain.Mac
    children["dynamic-arp-inspection"] = &bridgeDomain.DynamicArpInspection
    children["ip-source-guard"] = &bridgeDomain.IpSourceGuard
    children["storm-control"] = &bridgeDomain.StormControl
    children["igmp-snooping"] = &bridgeDomain.IgmpSnooping
    children["mld-snooping"] = &bridgeDomain.MldSnooping
    children["dhcp-ipv4-snooping"] = &bridgeDomain.DhcpIpv4Snooping
    return children
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = bridgeDomain.Id
    leafs["bridge-group"] = bridgeDomain.BridgeGroup
    leafs["enabled"] = bridgeDomain.Enabled
    leafs["bd-status-change-notification"] = bridgeDomain.BdStatusChangeNotification
    leafs["mtu"] = bridgeDomain.Mtu
    leafs["flooding-mode"] = bridgeDomain.FloodingMode
    return leafs
}

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetYangName() string { return "bridge-domain" }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) SetParent(parent types.Entity) { bridgeDomain.parent = parent }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetParent() types.Entity { return bridgeDomain.parent }

func (bridgeDomain *BridgeDomainConfig_BridgeDomains_BridgeDomain) GetParentYangName() string { return "bridge-domains" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members struct {
    parent types.Entity
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

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetFilter() yfilter.YFilter { return members.YFilter }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) SetFilter(yf yfilter.YFilter) { members.YFilter = yf }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetGoName(yname string) string {
    if yname == "ac-member" { return "AcMember" }
    if yname == "vfi-member" { return "VfiMember" }
    if yname == "access-pw-member" { return "AccessPwMember" }
    return ""
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetSegmentPath() string {
    return "members"
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ac-member" {
        for _, c := range members.AcMember {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember{}
        members.AcMember = append(members.AcMember, child)
        return &members.AcMember[len(members.AcMember)-1]
    }
    if childYangName == "vfi-member" {
        for _, c := range members.VfiMember {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember{}
        members.VfiMember = append(members.VfiMember, child)
        return &members.VfiMember[len(members.VfiMember)-1]
    }
    if childYangName == "access-pw-member" {
        return &members.AccessPwMember
    }
    return nil
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range members.AcMember {
        children[members.AcMember[i].GetSegmentPath()] = &members.AcMember[i]
    }
    for i := range members.VfiMember {
        children[members.VfiMember[i].GetSegmentPath()] = &members.VfiMember[i]
    }
    children["access-pw-member"] = &members.AccessPwMember
    return children
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetBundleName() string { return "cisco_ios_xe" }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetYangName() string { return "members" }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) SetParent(parent types.Entity) { members.parent = parent }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetParent() types.Entity { return members.parent }

func (members *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember
// List of Attachment circuits for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetFilter() yfilter.YFilter { return acMember.YFilter }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) SetFilter(yf yfilter.YFilter) { acMember.YFilter = yf }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "split-horizon-group" { return "SplitHorizonGroup" }
    if yname == "mac" { return "Mac" }
    if yname == "igmp-snooping" { return "IgmpSnooping" }
    if yname == "mld-snooping" { return "MldSnooping" }
    if yname == "dhcp-ipv4-snooping" { return "DhcpIpv4Snooping" }
    if yname == "flooding" { return "Flooding" }
    if yname == "storm-control" { return "StormControl" }
    if yname == "dynamic-arp-inspection" { return "DynamicArpInspection" }
    if yname == "ip-source-guard" { return "IpSourceGuard" }
    return ""
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetSegmentPath() string {
    return "ac-member" + "[interface='" + fmt.Sprintf("%v", acMember.Interface) + "']"
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "split-horizon-group" {
        return &acMember.SplitHorizonGroup
    }
    if childYangName == "mac" {
        return &acMember.Mac
    }
    if childYangName == "igmp-snooping" {
        return &acMember.IgmpSnooping
    }
    if childYangName == "mld-snooping" {
        return &acMember.MldSnooping
    }
    if childYangName == "dhcp-ipv4-snooping" {
        return &acMember.DhcpIpv4Snooping
    }
    if childYangName == "flooding" {
        return &acMember.Flooding
    }
    if childYangName == "storm-control" {
        return &acMember.StormControl
    }
    if childYangName == "dynamic-arp-inspection" {
        return &acMember.DynamicArpInspection
    }
    if childYangName == "ip-source-guard" {
        return &acMember.IpSourceGuard
    }
    return nil
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["split-horizon-group"] = &acMember.SplitHorizonGroup
    children["mac"] = &acMember.Mac
    children["igmp-snooping"] = &acMember.IgmpSnooping
    children["mld-snooping"] = &acMember.MldSnooping
    children["dhcp-ipv4-snooping"] = &acMember.DhcpIpv4Snooping
    children["flooding"] = &acMember.Flooding
    children["storm-control"] = &acMember.StormControl
    children["dynamic-arp-inspection"] = &acMember.DynamicArpInspection
    children["ip-source-guard"] = &acMember.IpSourceGuard
    return children
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = acMember.Interface
    return leafs
}

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetBundleName() string { return "cisco_ios_xe" }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetYangName() string { return "ac-member" }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) SetParent(parent types.Entity) { acMember.parent = parent }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetParent() types.Entity { return acMember.parent }

func (acMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember) GetParentYangName() string { return "members" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Split Horizon group number for bridge domain member. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Id interface{}
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetFilter() yfilter.YFilter { return splitHorizonGroup.YFilter }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) SetFilter(yf yfilter.YFilter) { splitHorizonGroup.YFilter = yf }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    return ""
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetSegmentPath() string {
    return "split-horizon-group"
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = splitHorizonGroup.Id
    return leafs
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetBundleName() string { return "cisco_ios_xe" }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetYangName() string { return "split-horizon-group" }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) SetParent(parent types.Entity) { splitHorizonGroup.parent = parent }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetParent() types.Entity { return splitHorizonGroup.parent }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_SplitHorizonGroup) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac struct {
    parent types.Entity
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

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetGoName(yname string) string {
    if yname == "learning-enabled" { return "LearningEnabled" }
    if yname == "limit" { return "Limit" }
    if yname == "aging" { return "Aging" }
    if yname == "port-down" { return "PortDown" }
    if yname == "secure" { return "Secure" }
    return ""
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limit" {
        return &mac.Limit
    }
    if childYangName == "aging" {
        return &mac.Aging
    }
    if childYangName == "port-down" {
        return &mac.PortDown
    }
    if childYangName == "secure" {
        return &mac.Secure
    }
    return nil
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["limit"] = &mac.Limit
    children["aging"] = &mac.Aging
    children["port-down"] = &mac.PortDown
    children["secure"] = &mac.Secure
    return children
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["learning-enabled"] = mac.LearningEnabled
    return leafs
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetBundleName() string { return "cisco_ios_xe" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetYangName() string { return "mac" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetParent() types.Entity { return mac.parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifSyslogNotifSnmpTrapNotifNoneNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetFilter() yfilter.YFilter { return limit.YFilter }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) SetFilter(yf yfilter.YFilter) { limit.YFilter = yf }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "action" { return "Action" }
    if yname == "notification" { return "Notification" }
    return ""
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetSegmentPath() string {
    return "limit"
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = limit.Maximum
    leafs["action"] = limit.Action
    leafs["notification"] = limit.Notification
    return leafs
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetBundleName() string { return "cisco_ios_xe" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetYangName() string { return "limit" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) SetParent(parent types.Entity) { limit.parent = parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetParent() types.Entity { return limit.parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Limit) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetFilter() yfilter.YFilter { return aging.YFilter }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) SetFilter(yf yfilter.YFilter) { aging.YFilter = yf }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetGoName(yname string) string {
    if yname == "time" { return "Time" }
    if yname == "type" { return "Type" }
    return ""
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetSegmentPath() string {
    return "aging"
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time"] = aging.Time
    leafs["type"] = aging.Type
    return leafs
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetBundleName() string { return "cisco_ios_xe" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetYangName() string { return "aging" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) SetParent(parent types.Entity) { aging.parent = parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetParent() types.Entity { return aging.parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Aging) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetFilter() yfilter.YFilter { return portDown.YFilter }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) SetFilter(yf yfilter.YFilter) { portDown.YFilter = yf }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetGoName(yname string) string {
    if yname == "flush" { return "Flush" }
    return ""
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetSegmentPath() string {
    return "port-down"
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flush"] = portDown.Flush
    return leafs
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetBundleName() string { return "cisco_ios_xe" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetYangName() string { return "port-down" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) SetParent(parent types.Entity) { portDown.parent = parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetParent() types.Entity { return portDown.parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_PortDown) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure
// MAC secure parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable mac secure feature. The type is bool.
    Enabled interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetFilter() yfilter.YFilter { return secure.YFilter }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) SetFilter(yf yfilter.YFilter) { secure.YFilter = yf }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "logging" { return "Logging" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetSegmentPath() string {
    return "secure"
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = secure.Action
    leafs["logging"] = secure.Logging
    leafs["enabled"] = secure.Enabled
    return leafs
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetBundleName() string { return "cisco_ios_xe" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetYangName() string { return "secure" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) SetParent(parent types.Entity) { secure.parent = parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetParent() types.Entity { return secure.parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Mac_Secure) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping
// Enable IGMP snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetFilter() yfilter.YFilter { return igmpSnooping.YFilter }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) SetFilter(yf yfilter.YFilter) { igmpSnooping.YFilter = yf }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetSegmentPath() string {
    return "igmp-snooping"
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = igmpSnooping.ProfileName
    return leafs
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetYangName() string { return "igmp-snooping" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) SetParent(parent types.Entity) { igmpSnooping.parent = parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetParent() types.Entity { return igmpSnooping.parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IgmpSnooping) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetFilter() yfilter.YFilter { return mldSnooping.YFilter }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) SetFilter(yf yfilter.YFilter) { mldSnooping.YFilter = yf }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetSegmentPath() string {
    return "mld-snooping"
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = mldSnooping.ProfileName
    return leafs
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetYangName() string { return "mld-snooping" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) SetParent(parent types.Entity) { mldSnooping.parent = parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetParent() types.Entity { return mldSnooping.parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_MldSnooping) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetFilter() yfilter.YFilter { return dhcpIpv4Snooping.YFilter }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) SetFilter(yf yfilter.YFilter) { dhcpIpv4Snooping.YFilter = yf }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetSegmentPath() string {
    return "dhcp-ipv4-snooping"
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = dhcpIpv4Snooping.ProfileName
    return leafs
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetBundleName() string { return "cisco_ios_xe" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetYangName() string { return "dhcp-ipv4-snooping" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) SetParent(parent types.Entity) { dhcpIpv4Snooping.parent = parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetParent() types.Entity { return dhcpIpv4Snooping.parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DhcpIpv4Snooping) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetFilter() yfilter.YFilter { return flooding.YFilter }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) SetFilter(yf yfilter.YFilter) { flooding.YFilter = yf }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetGoName(yname string) string {
    if yname == "disabled" { return "Disabled" }
    if yname == "disabled-unknown-unicast" { return "DisabledUnknownUnicast" }
    return ""
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetSegmentPath() string {
    return "flooding"
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disabled"] = flooding.Disabled
    leafs["disabled-unknown-unicast"] = flooding.DisabledUnknownUnicast
    return leafs
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetBundleName() string { return "cisco_ios_xe" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetYangName() string { return "flooding" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) SetParent(parent types.Entity) { flooding.parent = parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetParent() types.Entity { return flooding.parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_Flooding) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionShutdownActionSnmpTrapActionDrop.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetFilter() yfilter.YFilter { return stormControl.YFilter }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) SetFilter(yf yfilter.YFilter) { stormControl.YFilter = yf }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "thresholds" { return "Thresholds" }
    return ""
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetSegmentPath() string {
    return "storm-control"
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thresholds" {
        for _, c := range stormControl.Thresholds {
            if stormControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds{}
        stormControl.Thresholds = append(stormControl.Thresholds, child)
        return &stormControl.Thresholds[len(stormControl.Thresholds)-1]
    }
    return nil
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stormControl.Thresholds {
        children[stormControl.Thresholds[i].GetSegmentPath()] = &stormControl.Thresholds[i]
    }
    return children
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = stormControl.Action
    return leafs
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetBundleName() string { return "cisco_ios_xe" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetYangName() string { return "storm-control" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) SetParent(parent types.Entity) { stormControl.parent = parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetParent() types.Entity { return stormControl.parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds struct {
    parent types.Entity
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

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetFilter() yfilter.YFilter { return thresholds.YFilter }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) SetFilter(yf yfilter.YFilter) { thresholds.YFilter = yf }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetSegmentPath() string {
    return "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = thresholds.TrafficClass
    leafs["value"] = thresholds.Value
    leafs["unit"] = thresholds.Unit
    return leafs
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetBundleName() string { return "cisco_ios_xe" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetYangName() string { return "thresholds" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) SetParent(parent types.Entity) { thresholds.parent = parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetParent() types.Entity { return thresholds.parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_Thresholds) GetParentYangName() string { return "storm-control" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable DAI logging. The type is bool.
    Logging interface{}

    // Enable or disable Dynamic ARP Inspection. The type is bool.
    Enable interface{}

    // Enable address validation.
    AddressValidation BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetFilter() yfilter.YFilter { return dynamicArpInspection.YFilter }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) SetFilter(yf yfilter.YFilter) { dynamicArpInspection.YFilter = yf }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    if yname == "enable" { return "Enable" }
    if yname == "address-validation" { return "AddressValidation" }
    return ""
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetSegmentPath() string {
    return "dynamic-arp-inspection"
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-validation" {
        return &dynamicArpInspection.AddressValidation
    }
    return nil
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-validation"] = &dynamicArpInspection.AddressValidation
    return children
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging"] = dynamicArpInspection.Logging
    leafs["enable"] = dynamicArpInspection.Enable
    return leafs
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetBundleName() string { return "cisco_ios_xe" }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetYangName() string { return "dynamic-arp-inspection" }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) SetParent(parent types.Entity) { dynamicArpInspection.parent = parent }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetParent() types.Entity { return dynamicArpInspection.parent }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match Destination MAC Address. The type is interface{}.
    DstMac interface{}

    // Match Source MAC Address. The type is interface{}.
    SrcMac interface{}

    // Match IPv4 Address. The type is interface{}.
    Ipv4 interface{}
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetFilter() yfilter.YFilter { return addressValidation.YFilter }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) SetFilter(yf yfilter.YFilter) { addressValidation.YFilter = yf }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetGoName(yname string) string {
    if yname == "dst-mac" { return "DstMac" }
    if yname == "src-mac" { return "SrcMac" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetSegmentPath() string {
    return "address-validation"
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dst-mac"] = addressValidation.DstMac
    leafs["src-mac"] = addressValidation.SrcMac
    leafs["ipv4"] = addressValidation.Ipv4
    return leafs
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetBundleName() string { return "cisco_ios_xe" }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetYangName() string { return "address-validation" }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) SetParent(parent types.Entity) { addressValidation.parent = parent }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetParent() types.Entity { return addressValidation.parent }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_DynamicArpInspection_AddressValidation) GetParentYangName() string { return "dynamic-arp-inspection" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard
// IP source guard (IPSG) configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable IPSG logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable IP source guard feature. The type is bool.
    Enable interface{}
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetFilter() yfilter.YFilter { return ipSourceGuard.YFilter }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) SetFilter(yf yfilter.YFilter) { ipSourceGuard.YFilter = yf }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetSegmentPath() string {
    return "ip-source-guard"
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging"] = ipSourceGuard.Logging
    leafs["enable"] = ipSourceGuard.Enable
    return leafs
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetBundleName() string { return "cisco_ios_xe" }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetYangName() string { return "ip-source-guard" }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) SetParent(parent types.Entity) { ipSourceGuard.parent = parent }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetParent() types.Entity { return ipSourceGuard.parent }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AcMember_IpSourceGuard) GetParentYangName() string { return "ac-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember
// List of Virtual Forrwarding Interfaces for current
// bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an Virtual Forwarding Interface
    // instance which is configured to be part of this bridge-domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetFilter() yfilter.YFilter { return vfiMember.YFilter }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) SetFilter(yf yfilter.YFilter) { vfiMember.YFilter = yf }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetSegmentPath() string {
    return "vfi-member" + "[interface='" + fmt.Sprintf("%v", vfiMember.Interface) + "']"
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = vfiMember.Interface
    return leafs
}

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetBundleName() string { return "cisco_ios_xe" }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetYangName() string { return "vfi-member" }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) SetParent(parent types.Entity) { vfiMember.parent = parent }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetParent() types.Entity { return vfiMember.parent }

func (vfiMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_VfiMember) GetParentYangName() string { return "members" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember
// Collection of access pseudowire members.
// 
// A Pseudowires can be a regular interface with ifType
// 'ifPwType' or it can represented as a non-interface
// context. This container provides model definition for
// both types of the pseudowires.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember struct {
    parent types.Entity
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

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetFilter() yfilter.YFilter { return accessPwMember.YFilter }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) SetFilter(yf yfilter.YFilter) { accessPwMember.YFilter = yf }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetGoName(yname string) string {
    if yname == "access-pw-if-member" { return "AccessPwIfMember" }
    if yname == "pw-neighbor-spec" { return "PwNeighborSpec" }
    return ""
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetSegmentPath() string {
    return "access-pw-member"
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-pw-if-member" {
        for _, c := range accessPwMember.AccessPwIfMember {
            if accessPwMember.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember{}
        accessPwMember.AccessPwIfMember = append(accessPwMember.AccessPwIfMember, child)
        return &accessPwMember.AccessPwIfMember[len(accessPwMember.AccessPwIfMember)-1]
    }
    if childYangName == "pw-neighbor-spec" {
        for _, c := range accessPwMember.PwNeighborSpec {
            if accessPwMember.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec{}
        accessPwMember.PwNeighborSpec = append(accessPwMember.PwNeighborSpec, child)
        return &accessPwMember.PwNeighborSpec[len(accessPwMember.PwNeighborSpec)-1]
    }
    return nil
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessPwMember.AccessPwIfMember {
        children[accessPwMember.AccessPwIfMember[i].GetSegmentPath()] = &accessPwMember.AccessPwIfMember[i]
    }
    for i := range accessPwMember.PwNeighborSpec {
        children[accessPwMember.PwNeighborSpec[i].GetSegmentPath()] = &accessPwMember.PwNeighborSpec[i]
    }
    return children
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetBundleName() string { return "cisco_ios_xe" }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetYangName() string { return "access-pw-member" }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) SetParent(parent types.Entity) { accessPwMember.parent = parent }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetParent() types.Entity { return accessPwMember.parent }

func (accessPwMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetParentYangName() string { return "members" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember
// List of interface based access pseudowires for
// current bridge-domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an access pseudo-wire interface
    // instance which is configured to be part of this bridge domain. The type is
    // string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetFilter() yfilter.YFilter { return accessPwIfMember.YFilter }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) SetFilter(yf yfilter.YFilter) { accessPwIfMember.YFilter = yf }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetSegmentPath() string {
    return "access-pw-if-member" + "[interface='" + fmt.Sprintf("%v", accessPwIfMember.Interface) + "']"
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = accessPwIfMember.Interface
    return leafs
}

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetBundleName() string { return "cisco_ios_xe" }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetYangName() string { return "access-pw-if-member" }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) SetParent(parent types.Entity) { accessPwIfMember.parent = parent }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetParent() types.Entity { return accessPwIfMember.parent }

func (accessPwIfMember *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_AccessPwIfMember) GetParentYangName() string { return "access-pw-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec
// Collection of neighbor specification based
// pseudo-wires.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 address of the neighbor. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetFilter() yfilter.YFilter { return pwNeighborSpec.YFilter }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) SetFilter(yf yfilter.YFilter) { pwNeighborSpec.YFilter = yf }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetGoName(yname string) string {
    if yname == "neighbor-ip-address" { return "NeighborIpAddress" }
    if yname == "vc-id" { return "VcId" }
    if yname == "pw-class-template" { return "PwClassTemplate" }
    if yname == "encap-type" { return "EncapType" }
    if yname == "tag-impose-vlan" { return "TagImposeVlan" }
    if yname == "source-ipv6" { return "SourceIpv6" }
    if yname == "static-label" { return "StaticLabel" }
    if yname == "split-horizon-group" { return "SplitHorizonGroup" }
    if yname == "mac" { return "Mac" }
    if yname == "igmp-snooping" { return "IgmpSnooping" }
    if yname == "mld-snooping" { return "MldSnooping" }
    if yname == "dhcp-ipv4-snooping" { return "DhcpIpv4Snooping" }
    if yname == "flooding" { return "Flooding" }
    if yname == "storm-control" { return "StormControl" }
    if yname == "backup" { return "Backup" }
    return ""
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetSegmentPath() string {
    return "pw-neighbor-spec" + "[neighbor-ip-address='" + fmt.Sprintf("%v", pwNeighborSpec.NeighborIpAddress) + "']" + "[vc-id='" + fmt.Sprintf("%v", pwNeighborSpec.VcId) + "']"
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-label" {
        return &pwNeighborSpec.StaticLabel
    }
    if childYangName == "split-horizon-group" {
        return &pwNeighborSpec.SplitHorizonGroup
    }
    if childYangName == "mac" {
        return &pwNeighborSpec.Mac
    }
    if childYangName == "igmp-snooping" {
        return &pwNeighborSpec.IgmpSnooping
    }
    if childYangName == "mld-snooping" {
        return &pwNeighborSpec.MldSnooping
    }
    if childYangName == "dhcp-ipv4-snooping" {
        return &pwNeighborSpec.DhcpIpv4Snooping
    }
    if childYangName == "flooding" {
        return &pwNeighborSpec.Flooding
    }
    if childYangName == "storm-control" {
        return &pwNeighborSpec.StormControl
    }
    if childYangName == "backup" {
        return &pwNeighborSpec.Backup
    }
    return nil
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-label"] = &pwNeighborSpec.StaticLabel
    children["split-horizon-group"] = &pwNeighborSpec.SplitHorizonGroup
    children["mac"] = &pwNeighborSpec.Mac
    children["igmp-snooping"] = &pwNeighborSpec.IgmpSnooping
    children["mld-snooping"] = &pwNeighborSpec.MldSnooping
    children["dhcp-ipv4-snooping"] = &pwNeighborSpec.DhcpIpv4Snooping
    children["flooding"] = &pwNeighborSpec.Flooding
    children["storm-control"] = &pwNeighborSpec.StormControl
    children["backup"] = &pwNeighborSpec.Backup
    return children
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-ip-address"] = pwNeighborSpec.NeighborIpAddress
    leafs["vc-id"] = pwNeighborSpec.VcId
    leafs["pw-class-template"] = pwNeighborSpec.PwClassTemplate
    leafs["encap-type"] = pwNeighborSpec.EncapType
    leafs["tag-impose-vlan"] = pwNeighborSpec.TagImposeVlan
    leafs["source-ipv6"] = pwNeighborSpec.SourceIpv6
    return leafs
}

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetBundleName() string { return "cisco_ios_xe" }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetYangName() string { return "pw-neighbor-spec" }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) SetParent(parent types.Entity) { pwNeighborSpec.parent = parent }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetParent() types.Entity { return pwNeighborSpec.parent }

func (pwNeighborSpec *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec) GetParentYangName() string { return "access-pw-member" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel
// Statically configured labels, signalling should be none
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local MPLS label ID. The type is interface{} with range: 16..1048575.
    LocalLabel interface{}

    // Remote MPLS label ID. The type is interface{} with range: 16..1048575.
    RemoteLabel interface{}
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetFilter() yfilter.YFilter { return staticLabel.YFilter }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) SetFilter(yf yfilter.YFilter) { staticLabel.YFilter = yf }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    if yname == "remote-label" { return "RemoteLabel" }
    return ""
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetSegmentPath() string {
    return "static-label"
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label"] = staticLabel.LocalLabel
    leafs["remote-label"] = staticLabel.RemoteLabel
    return leafs
}

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetBundleName() string { return "cisco_ios_xe" }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetYangName() string { return "static-label" }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) SetParent(parent types.Entity) { staticLabel.parent = parent }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetParent() types.Entity { return staticLabel.parent }

func (staticLabel *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StaticLabel) GetParentYangName() string { return "pw-neighbor-spec" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Split Horizon group number for bridge domain member. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Id interface{}
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetFilter() yfilter.YFilter { return splitHorizonGroup.YFilter }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) SetFilter(yf yfilter.YFilter) { splitHorizonGroup.YFilter = yf }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    return ""
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetSegmentPath() string {
    return "split-horizon-group"
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = splitHorizonGroup.Id
    return leafs
}

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetBundleName() string { return "cisco_ios_xe" }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetYangName() string { return "split-horizon-group" }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) SetParent(parent types.Entity) { splitHorizonGroup.parent = parent }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetParent() types.Entity { return splitHorizonGroup.parent }

func (splitHorizonGroup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_SplitHorizonGroup) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac struct {
    parent types.Entity
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

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetGoName(yname string) string {
    if yname == "learning-enabled" { return "LearningEnabled" }
    if yname == "limit" { return "Limit" }
    if yname == "aging" { return "Aging" }
    if yname == "port-down" { return "PortDown" }
    if yname == "secure" { return "Secure" }
    return ""
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limit" {
        return &mac.Limit
    }
    if childYangName == "aging" {
        return &mac.Aging
    }
    if childYangName == "port-down" {
        return &mac.PortDown
    }
    if childYangName == "secure" {
        return &mac.Secure
    }
    return nil
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["limit"] = &mac.Limit
    children["aging"] = &mac.Aging
    children["port-down"] = &mac.PortDown
    children["secure"] = &mac.Secure
    return children
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["learning-enabled"] = mac.LearningEnabled
    return leafs
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetBundleName() string { return "cisco_ios_xe" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetYangName() string { return "mac" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetParent() types.Entity { return mac.parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifSyslogNotifSnmpTrapNotifNoneNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetFilter() yfilter.YFilter { return limit.YFilter }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) SetFilter(yf yfilter.YFilter) { limit.YFilter = yf }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "action" { return "Action" }
    if yname == "notification" { return "Notification" }
    return ""
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetSegmentPath() string {
    return "limit"
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = limit.Maximum
    leafs["action"] = limit.Action
    leafs["notification"] = limit.Notification
    return leafs
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetBundleName() string { return "cisco_ios_xe" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetYangName() string { return "limit" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) SetParent(parent types.Entity) { limit.parent = parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetParent() types.Entity { return limit.parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Limit) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetFilter() yfilter.YFilter { return aging.YFilter }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) SetFilter(yf yfilter.YFilter) { aging.YFilter = yf }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetGoName(yname string) string {
    if yname == "time" { return "Time" }
    if yname == "type" { return "Type" }
    return ""
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetSegmentPath() string {
    return "aging"
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time"] = aging.Time
    leafs["type"] = aging.Type
    return leafs
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetBundleName() string { return "cisco_ios_xe" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetYangName() string { return "aging" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) SetParent(parent types.Entity) { aging.parent = parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetParent() types.Entity { return aging.parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Aging) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetFilter() yfilter.YFilter { return portDown.YFilter }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) SetFilter(yf yfilter.YFilter) { portDown.YFilter = yf }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetGoName(yname string) string {
    if yname == "flush" { return "Flush" }
    return ""
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetSegmentPath() string {
    return "port-down"
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flush"] = portDown.Flush
    return leafs
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetBundleName() string { return "cisco_ios_xe" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetYangName() string { return "port-down" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) SetParent(parent types.Entity) { portDown.parent = parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetParent() types.Entity { return portDown.parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_PortDown) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure
// MAC secure parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}

    // Enable or disable mac secure feature. The type is bool.
    Enabled interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetFilter() yfilter.YFilter { return secure.YFilter }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) SetFilter(yf yfilter.YFilter) { secure.YFilter = yf }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "logging" { return "Logging" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetSegmentPath() string {
    return "secure"
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = secure.Action
    leafs["logging"] = secure.Logging
    leafs["enabled"] = secure.Enabled
    return leafs
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetBundleName() string { return "cisco_ios_xe" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetYangName() string { return "secure" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) SetParent(parent types.Entity) { secure.parent = parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetParent() types.Entity { return secure.parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Mac_Secure) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping
// Enable IGMP snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetFilter() yfilter.YFilter { return igmpSnooping.YFilter }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) SetFilter(yf yfilter.YFilter) { igmpSnooping.YFilter = yf }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetSegmentPath() string {
    return "igmp-snooping"
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = igmpSnooping.ProfileName
    return leafs
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetYangName() string { return "igmp-snooping" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) SetParent(parent types.Entity) { igmpSnooping.parent = parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetParent() types.Entity { return igmpSnooping.parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_IgmpSnooping) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetFilter() yfilter.YFilter { return mldSnooping.YFilter }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) SetFilter(yf yfilter.YFilter) { mldSnooping.YFilter = yf }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetSegmentPath() string {
    return "mld-snooping"
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = mldSnooping.ProfileName
    return leafs
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetYangName() string { return "mld-snooping" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) SetParent(parent types.Entity) { mldSnooping.parent = parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetParent() types.Entity { return mldSnooping.parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_MldSnooping) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetFilter() yfilter.YFilter { return dhcpIpv4Snooping.YFilter }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) SetFilter(yf yfilter.YFilter) { dhcpIpv4Snooping.YFilter = yf }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetSegmentPath() string {
    return "dhcp-ipv4-snooping"
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = dhcpIpv4Snooping.ProfileName
    return leafs
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetBundleName() string { return "cisco_ios_xe" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetYangName() string { return "dhcp-ipv4-snooping" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) SetParent(parent types.Entity) { dhcpIpv4Snooping.parent = parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetParent() types.Entity { return dhcpIpv4Snooping.parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_DhcpIpv4Snooping) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetFilter() yfilter.YFilter { return flooding.YFilter }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) SetFilter(yf yfilter.YFilter) { flooding.YFilter = yf }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetGoName(yname string) string {
    if yname == "disabled" { return "Disabled" }
    if yname == "disabled-unknown-unicast" { return "DisabledUnknownUnicast" }
    return ""
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetSegmentPath() string {
    return "flooding"
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disabled"] = flooding.Disabled
    leafs["disabled-unknown-unicast"] = flooding.DisabledUnknownUnicast
    return leafs
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetBundleName() string { return "cisco_ios_xe" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetYangName() string { return "flooding" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) SetParent(parent types.Entity) { flooding.parent = parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetParent() types.Entity { return flooding.parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Flooding) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionShutdownActionSnmpTrapActionDrop.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetFilter() yfilter.YFilter { return stormControl.YFilter }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) SetFilter(yf yfilter.YFilter) { stormControl.YFilter = yf }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "thresholds" { return "Thresholds" }
    return ""
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetSegmentPath() string {
    return "storm-control"
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thresholds" {
        for _, c := range stormControl.Thresholds {
            if stormControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds{}
        stormControl.Thresholds = append(stormControl.Thresholds, child)
        return &stormControl.Thresholds[len(stormControl.Thresholds)-1]
    }
    return nil
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stormControl.Thresholds {
        children[stormControl.Thresholds[i].GetSegmentPath()] = &stormControl.Thresholds[i]
    }
    return children
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = stormControl.Action
    return leafs
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetBundleName() string { return "cisco_ios_xe" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetYangName() string { return "storm-control" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) SetParent(parent types.Entity) { stormControl.parent = parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetParent() types.Entity { return stormControl.parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds struct {
    parent types.Entity
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

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetFilter() yfilter.YFilter { return thresholds.YFilter }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) SetFilter(yf yfilter.YFilter) { thresholds.YFilter = yf }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetSegmentPath() string {
    return "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = thresholds.TrafficClass
    leafs["value"] = thresholds.Value
    leafs["unit"] = thresholds.Unit
    return leafs
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetBundleName() string { return "cisco_ios_xe" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetYangName() string { return "thresholds" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) SetParent(parent types.Entity) { thresholds.parent = parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetParent() types.Entity { return thresholds.parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_StormControl_Thresholds) GetParentYangName() string { return "storm-control" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 or IPv6 address of the neighbor. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NeighborIpAddress interface{}

    // Pseudowire VC ID. The type is interface{} with range: 1..4294967295.
    VcId interface{}

    // Reference to a pseudowire template. The type is string. Refers to
    // cisco_pw.PseudowireConfig_PwTemplates_PwTemplate_Name
    PwClassTemplate interface{}
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetFilter() yfilter.YFilter { return backup.YFilter }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) SetFilter(yf yfilter.YFilter) { backup.YFilter = yf }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetGoName(yname string) string {
    if yname == "neighbor-ip-address" { return "NeighborIpAddress" }
    if yname == "vc-id" { return "VcId" }
    if yname == "pw-class-template" { return "PwClassTemplate" }
    return ""
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetSegmentPath() string {
    return "backup"
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["neighbor-ip-address"] = backup.NeighborIpAddress
    leafs["vc-id"] = backup.VcId
    leafs["pw-class-template"] = backup.PwClassTemplate
    return leafs
}

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetBundleName() string { return "cisco_ios_xe" }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetYangName() string { return "backup" }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) SetParent(parent types.Entity) { backup.parent = parent }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetParent() types.Entity { return backup.parent }

func (backup *BridgeDomainConfig_BridgeDomains_BridgeDomain_Members_AccessPwMember_PwNeighborSpec_Backup) GetParentYangName() string { return "pw-neighbor-spec" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac
// MAC features for bridge domain.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac struct {
    parent types.Entity
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

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetGoName(yname string) string {
    if yname == "learning-enabled" { return "LearningEnabled" }
    if yname == "limit" { return "Limit" }
    if yname == "aging" { return "Aging" }
    if yname == "port-down" { return "PortDown" }
    if yname == "flooding" { return "Flooding" }
    if yname == "secure" { return "Secure" }
    if yname == "static" { return "Static" }
    return ""
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "limit" {
        return &mac.Limit
    }
    if childYangName == "aging" {
        return &mac.Aging
    }
    if childYangName == "port-down" {
        return &mac.PortDown
    }
    if childYangName == "flooding" {
        return &mac.Flooding
    }
    if childYangName == "secure" {
        return &mac.Secure
    }
    if childYangName == "static" {
        return &mac.Static
    }
    return nil
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["limit"] = &mac.Limit
    children["aging"] = &mac.Aging
    children["port-down"] = &mac.PortDown
    children["flooding"] = &mac.Flooding
    children["secure"] = &mac.Secure
    children["static"] = &mac.Static
    return children
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["learning-enabled"] = mac.LearningEnabled
    return leafs
}

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetBundleName() string { return "cisco_ios_xe" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetYangName() string { return "mac" }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetParent() types.Entity { return mac.parent }

func (mac *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit
// MAC table learning limit.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum number of mac addresses that can be learnt. The type is interface{}
    // with range: 0..4294967295.
    Maximum interface{}

    // MAC limit violation actions. The type is MacLimitAction.
    Action interface{}

    // MAC limit violation notifications. The type is one of the following:
    // NotifSyslogNotifSnmpTrapNotifNoneNotifSyslogAndSnmpTrap.
    Notification interface{}
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetFilter() yfilter.YFilter { return limit.YFilter }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) SetFilter(yf yfilter.YFilter) { limit.YFilter = yf }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    if yname == "action" { return "Action" }
    if yname == "notification" { return "Notification" }
    return ""
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetSegmentPath() string {
    return "limit"
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = limit.Maximum
    leafs["action"] = limit.Action
    leafs["notification"] = limit.Notification
    return leafs
}

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetBundleName() string { return "cisco_ios_xe" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetYangName() string { return "limit" }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) SetParent(parent types.Entity) { limit.parent = parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetParent() types.Entity { return limit.parent }

func (limit *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Limit) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging
// MAC aging configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The timeout period in seconds for aging out dynamically learned forwarding
    // information. The type is interface{} with range: 0..4294967295. Units are
    // seconds. The default value is 300.
    Time interface{}

    // MAC aging type. The type is MacAgingType.
    Type interface{}
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetFilter() yfilter.YFilter { return aging.YFilter }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) SetFilter(yf yfilter.YFilter) { aging.YFilter = yf }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetGoName(yname string) string {
    if yname == "time" { return "Time" }
    if yname == "type" { return "Type" }
    return ""
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetSegmentPath() string {
    return "aging"
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time"] = aging.Time
    leafs["type"] = aging.Type
    return leafs
}

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetBundleName() string { return "cisco_ios_xe" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetYangName() string { return "aging" }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) SetParent(parent types.Entity) { aging.parent = parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetParent() types.Entity { return aging.parent }

func (aging *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Aging) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown
// Port down event
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable mac table flush when port moves to down state. The type is
    // bool. The default value is true.
    Flush interface{}
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetFilter() yfilter.YFilter { return portDown.YFilter }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) SetFilter(yf yfilter.YFilter) { portDown.YFilter = yf }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetGoName(yname string) string {
    if yname == "flush" { return "Flush" }
    return ""
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetSegmentPath() string {
    return "port-down"
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flush"] = portDown.Flush
    return leafs
}

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetBundleName() string { return "cisco_ios_xe" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetYangName() string { return "port-down" }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) SetParent(parent types.Entity) { portDown.parent = parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetParent() types.Entity { return portDown.parent }

func (portDown *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_PortDown) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding
// Flooding configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable flooding. The type is interface{}.
    Disabled interface{}

    // Disable unknown unicast flooding. The type is interface{}.
    DisabledUnknownUnicast interface{}
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetFilter() yfilter.YFilter { return flooding.YFilter }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) SetFilter(yf yfilter.YFilter) { flooding.YFilter = yf }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetGoName(yname string) string {
    if yname == "disabled" { return "Disabled" }
    if yname == "disabled-unknown-unicast" { return "DisabledUnknownUnicast" }
    return ""
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetSegmentPath() string {
    return "flooding"
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disabled"] = flooding.Disabled
    leafs["disabled-unknown-unicast"] = flooding.DisabledUnknownUnicast
    return leafs
}

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetBundleName() string { return "cisco_ios_xe" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetYangName() string { return "flooding" }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) SetParent(parent types.Entity) { flooding.parent = parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetParent() types.Entity { return flooding.parent }

func (flooding *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Flooding) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure
// MAC secure parameters.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC secure action for violating packets. The type is MacSecureAction. The
    // default value is restrict.
    Action interface{}

    // Enable/Disable logging. The type is bool. The default value is false.
    Logging interface{}
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetFilter() yfilter.YFilter { return secure.YFilter }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) SetFilter(yf yfilter.YFilter) { secure.YFilter = yf }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "logging" { return "Logging" }
    return ""
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetSegmentPath() string {
    return "secure"
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = secure.Action
    leafs["logging"] = secure.Logging
    return leafs
}

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetBundleName() string { return "cisco_ios_xe" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetYangName() string { return "secure" }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) SetParent(parent types.Entity) { secure.parent = parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetParent() types.Entity { return secure.parent }

func (secure *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Secure) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static
// Static mac address list parameters.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC address entry. The type is slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses.
    MacAddresses []BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetGoName(yname string) string {
    if yname == "mac-addresses" { return "MacAddresses" }
    return ""
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetSegmentPath() string {
    return "static"
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-addresses" {
        for _, c := range static.MacAddresses {
            if static.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses{}
        static.MacAddresses = append(static.MacAddresses, child)
        return &static.MacAddresses[len(static.MacAddresses)-1]
    }
    return nil
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range static.MacAddresses {
        children[static.MacAddresses[i].GetSegmentPath()] = &static.MacAddresses[i]
    }
    return children
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetBundleName() string { return "cisco_ios_xe" }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetYangName() string { return "static" }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetParent() types.Entity { return static.parent }

func (static *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static) GetParentYangName() string { return "mac" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses
// MAC address entry.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Static MAC address. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddr interface{}

    // Drop packet. The type is bool. This attribute is mandatory.
    Drop interface{}
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetFilter() yfilter.YFilter { return macAddresses.YFilter }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) SetFilter(yf yfilter.YFilter) { macAddresses.YFilter = yf }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetGoName(yname string) string {
    if yname == "mac-addr" { return "MacAddr" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetSegmentPath() string {
    return "mac-addresses" + "[mac-addr='" + fmt.Sprintf("%v", macAddresses.MacAddr) + "']"
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-addr"] = macAddresses.MacAddr
    leafs["drop"] = macAddresses.Drop
    return leafs
}

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetBundleName() string { return "cisco_ios_xe" }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetYangName() string { return "mac-addresses" }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) SetParent(parent types.Entity) { macAddresses.parent = parent }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetParent() types.Entity { return macAddresses.parent }

func (macAddresses *BridgeDomainConfig_BridgeDomains_BridgeDomain_Mac_Static_MacAddresses) GetParentYangName() string { return "static" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection
// Dynamic ARP Inspection (DAI) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable DAI logging. The type is bool.
    Logging interface{}

    // Enable address validation.
    AddressValidation BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetFilter() yfilter.YFilter { return dynamicArpInspection.YFilter }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) SetFilter(yf yfilter.YFilter) { dynamicArpInspection.YFilter = yf }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    if yname == "address-validation" { return "AddressValidation" }
    return ""
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetSegmentPath() string {
    return "dynamic-arp-inspection"
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-validation" {
        return &dynamicArpInspection.AddressValidation
    }
    return nil
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address-validation"] = &dynamicArpInspection.AddressValidation
    return children
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging"] = dynamicArpInspection.Logging
    return leafs
}

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetBundleName() string { return "cisco_ios_xe" }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetYangName() string { return "dynamic-arp-inspection" }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) SetParent(parent types.Entity) { dynamicArpInspection.parent = parent }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetParent() types.Entity { return dynamicArpInspection.parent }

func (dynamicArpInspection *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation
// Enable address validation.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match Destination MAC Address. The type is interface{}.
    DstMac interface{}

    // Match Source MAC Address. The type is interface{}.
    SrcMac interface{}

    // Match IPv4 Address. The type is interface{}.
    Ipv4 interface{}
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetFilter() yfilter.YFilter { return addressValidation.YFilter }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) SetFilter(yf yfilter.YFilter) { addressValidation.YFilter = yf }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetGoName(yname string) string {
    if yname == "dst-mac" { return "DstMac" }
    if yname == "src-mac" { return "SrcMac" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetSegmentPath() string {
    return "address-validation"
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dst-mac"] = addressValidation.DstMac
    leafs["src-mac"] = addressValidation.SrcMac
    leafs["ipv4"] = addressValidation.Ipv4
    return leafs
}

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetBundleName() string { return "cisco_ios_xe" }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetYangName() string { return "address-validation" }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) SetParent(parent types.Entity) { addressValidation.parent = parent }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetParent() types.Entity { return addressValidation.parent }

func (addressValidation *BridgeDomainConfig_BridgeDomains_BridgeDomain_DynamicArpInspection_AddressValidation) GetParentYangName() string { return "dynamic-arp-inspection" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard
// IP source guard (IPSG) configurations.
// This type is a presence type.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable IPSG logging. The type is bool. The default value is false.
    Logging interface{}
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetFilter() yfilter.YFilter { return ipSourceGuard.YFilter }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) SetFilter(yf yfilter.YFilter) { ipSourceGuard.YFilter = yf }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    return ""
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetSegmentPath() string {
    return "ip-source-guard"
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["logging"] = ipSourceGuard.Logging
    return leafs
}

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetBundleName() string { return "cisco_ios_xe" }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetYangName() string { return "ip-source-guard" }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) SetParent(parent types.Entity) { ipSourceGuard.parent = parent }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetParent() types.Entity { return ipSourceGuard.parent }

func (ipSourceGuard *BridgeDomainConfig_BridgeDomains_BridgeDomain_IpSourceGuard) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl
// A collection of storm control threshold and action
// configurations.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This leaf represents the storm control action taken when the traffic of a
    // particular type exceeds the configured threshold values. The type is one of
    // the following: ActionShutdownActionSnmpTrapActionDrop.
    Action interface{}

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds.
    Thresholds []BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetFilter() yfilter.YFilter { return stormControl.YFilter }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) SetFilter(yf yfilter.YFilter) { stormControl.YFilter = yf }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "thresholds" { return "Thresholds" }
    return ""
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetSegmentPath() string {
    return "storm-control"
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thresholds" {
        for _, c := range stormControl.Thresholds {
            if stormControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds{}
        stormControl.Thresholds = append(stormControl.Thresholds, child)
        return &stormControl.Thresholds[len(stormControl.Thresholds)-1]
    }
    return nil
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stormControl.Thresholds {
        children[stormControl.Thresholds[i].GetSegmentPath()] = &stormControl.Thresholds[i]
    }
    return children
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = stormControl.Action
    return leafs
}

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetBundleName() string { return "cisco_ios_xe" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetYangName() string { return "storm-control" }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) SetParent(parent types.Entity) { stormControl.parent = parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetParent() types.Entity { return stormControl.parent }

func (stormControl *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds
// A collection of storm control threshold configuration
// entries.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds struct {
    parent types.Entity
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

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetFilter() yfilter.YFilter { return thresholds.YFilter }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) SetFilter(yf yfilter.YFilter) { thresholds.YFilter = yf }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetSegmentPath() string {
    return "thresholds" + "[traffic-class='" + fmt.Sprintf("%v", thresholds.TrafficClass) + "']"
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = thresholds.TrafficClass
    leafs["value"] = thresholds.Value
    leafs["unit"] = thresholds.Unit
    return leafs
}

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetBundleName() string { return "cisco_ios_xe" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetYangName() string { return "thresholds" }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) SetParent(parent types.Entity) { thresholds.parent = parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetParent() types.Entity { return thresholds.parent }

func (thresholds *BridgeDomainConfig_BridgeDomains_BridgeDomain_StormControl_Thresholds) GetParentYangName() string { return "storm-control" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // IGMP snooping profile name. The type is string.
    ProfileName interface{}

    // Disable IGMP snooping. The type is interface{}.
    Disabled interface{}
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetFilter() yfilter.YFilter { return igmpSnooping.YFilter }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) SetFilter(yf yfilter.YFilter) { igmpSnooping.YFilter = yf }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "disabled" { return "Disabled" }
    return ""
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetSegmentPath() string {
    return "igmp-snooping"
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = igmpSnooping.ProfileName
    leafs["disabled"] = igmpSnooping.Disabled
    return leafs
}

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetYangName() string { return "igmp-snooping" }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) SetParent(parent types.Entity) { igmpSnooping.parent = parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetParent() types.Entity { return igmpSnooping.parent }

func (igmpSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_IgmpSnooping) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping
// Enable MLD snooping
type BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MLD snooping profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetFilter() yfilter.YFilter { return mldSnooping.YFilter }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) SetFilter(yf yfilter.YFilter) { mldSnooping.YFilter = yf }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetSegmentPath() string {
    return "mld-snooping"
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = mldSnooping.ProfileName
    return leafs
}

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetBundleName() string { return "cisco_ios_xe" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetYangName() string { return "mld-snooping" }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) SetParent(parent types.Entity) { mldSnooping.parent = parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetParent() types.Entity { return mldSnooping.parent }

func (mldSnooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_MldSnooping) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping
// Enable DHCP IPv4 snooping.
type BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv4 snooping profile name. The type is string. This attribute is
    // mandatory.
    ProfileName interface{}
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetFilter() yfilter.YFilter { return dhcpIpv4Snooping.YFilter }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) SetFilter(yf yfilter.YFilter) { dhcpIpv4Snooping.YFilter = yf }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    return ""
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetSegmentPath() string {
    return "dhcp-ipv4-snooping"
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = dhcpIpv4Snooping.ProfileName
    return leafs
}

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetBundleName() string { return "cisco_ios_xe" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetYangName() string { return "dhcp-ipv4-snooping" }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) SetParent(parent types.Entity) { dhcpIpv4Snooping.parent = parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetParent() types.Entity { return dhcpIpv4Snooping.parent }

func (dhcpIpv4Snooping *BridgeDomainConfig_BridgeDomains_BridgeDomain_DhcpIpv4Snooping) GetParentYangName() string { return "bridge-domain" }

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
    parent types.Entity
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

func (bridgeDomainState *BridgeDomainState) GetFilter() yfilter.YFilter { return bridgeDomainState.YFilter }

func (bridgeDomainState *BridgeDomainState) SetFilter(yf yfilter.YFilter) { bridgeDomainState.YFilter = yf }

func (bridgeDomainState *BridgeDomainState) GetGoName(yname string) string {
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "module-capabilities" { return "ModuleCapabilities" }
    if yname == "bridge-domains" { return "BridgeDomains" }
    if yname == "mac-table" { return "MacTable" }
    return ""
}

func (bridgeDomainState *BridgeDomainState) GetSegmentPath() string {
    return "cisco-bridge-domain:bridge-domain-state"
}

func (bridgeDomainState *BridgeDomainState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "system-capabilities" {
        return &bridgeDomainState.SystemCapabilities
    }
    if childYangName == "module-capabilities" {
        return &bridgeDomainState.ModuleCapabilities
    }
    if childYangName == "bridge-domains" {
        return &bridgeDomainState.BridgeDomains
    }
    if childYangName == "mac-table" {
        for _, c := range bridgeDomainState.MacTable {
            if bridgeDomainState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_MacTable{}
        bridgeDomainState.MacTable = append(bridgeDomainState.MacTable, child)
        return &bridgeDomainState.MacTable[len(bridgeDomainState.MacTable)-1]
    }
    return nil
}

func (bridgeDomainState *BridgeDomainState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["system-capabilities"] = &bridgeDomainState.SystemCapabilities
    children["module-capabilities"] = &bridgeDomainState.ModuleCapabilities
    children["bridge-domains"] = &bridgeDomainState.BridgeDomains
    for i := range bridgeDomainState.MacTable {
        children[bridgeDomainState.MacTable[i].GetSegmentPath()] = &bridgeDomainState.MacTable[i]
    }
    return children
}

func (bridgeDomainState *BridgeDomainState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bridgeDomainState *BridgeDomainState) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomainState *BridgeDomainState) GetYangName() string { return "bridge-domain-state" }

func (bridgeDomainState *BridgeDomainState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomainState *BridgeDomainState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomainState *BridgeDomainState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomainState *BridgeDomainState) SetParent(parent types.Entity) { bridgeDomainState.parent = parent }

func (bridgeDomainState *BridgeDomainState) GetParent() types.Entity { return bridgeDomainState.parent }

func (bridgeDomainState *BridgeDomainState) GetParentYangName() string { return "cisco-bridge-domain" }

// BridgeDomainState_SystemCapabilities
// This container defines system capabilities for bridge
// domain.
type BridgeDomainState_SystemCapabilities struct {
    parent types.Entity
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

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetFilter() yfilter.YFilter { return systemCapabilities.YFilter }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) SetFilter(yf yfilter.YFilter) { systemCapabilities.YFilter = yf }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetGoName(yname string) string {
    if yname == "max-bd" { return "MaxBd" }
    if yname == "max-ac-per-bd" { return "MaxAcPerBd" }
    if yname == "max-pw-per-bd" { return "MaxPwPerBd" }
    if yname == "max-vfi-per-bd" { return "MaxVfiPerBd" }
    if yname == "max-sh-group-per-bd" { return "MaxShGroupPerBd" }
    if yname == "max-interflex-if-per-bd" { return "MaxInterflexIfPerBd" }
    return ""
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetSegmentPath() string {
    return "system-capabilities"
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-bd"] = systemCapabilities.MaxBd
    leafs["max-ac-per-bd"] = systemCapabilities.MaxAcPerBd
    leafs["max-pw-per-bd"] = systemCapabilities.MaxPwPerBd
    leafs["max-vfi-per-bd"] = systemCapabilities.MaxVfiPerBd
    leafs["max-sh-group-per-bd"] = systemCapabilities.MaxShGroupPerBd
    leafs["max-interflex-if-per-bd"] = systemCapabilities.MaxInterflexIfPerBd
    return leafs
}

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetBundleName() string { return "cisco_ios_xe" }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetYangName() string { return "system-capabilities" }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) SetParent(parent types.Entity) { systemCapabilities.parent = parent }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetParent() types.Entity { return systemCapabilities.parent }

func (systemCapabilities *BridgeDomainState_SystemCapabilities) GetParentYangName() string { return "bridge-domain-state" }

// BridgeDomainState_ModuleCapabilities
// This container defines module capabilities for bridge
// domain.
type BridgeDomainState_ModuleCapabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of capabillity statements for hardware module in the system. The
    // type is slice of BridgeDomainState_ModuleCapabilities_Modules.
    Modules []BridgeDomainState_ModuleCapabilities_Modules
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetFilter() yfilter.YFilter { return moduleCapabilities.YFilter }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) SetFilter(yf yfilter.YFilter) { moduleCapabilities.YFilter = yf }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetGoName(yname string) string {
    if yname == "modules" { return "Modules" }
    return ""
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetSegmentPath() string {
    return "module-capabilities"
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "modules" {
        for _, c := range moduleCapabilities.Modules {
            if moduleCapabilities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_ModuleCapabilities_Modules{}
        moduleCapabilities.Modules = append(moduleCapabilities.Modules, child)
        return &moduleCapabilities.Modules[len(moduleCapabilities.Modules)-1]
    }
    return nil
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range moduleCapabilities.Modules {
        children[moduleCapabilities.Modules[i].GetSegmentPath()] = &moduleCapabilities.Modules[i]
    }
    return children
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetBundleName() string { return "cisco_ios_xe" }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetYangName() string { return "module-capabilities" }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) SetParent(parent types.Entity) { moduleCapabilities.parent = parent }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetParent() types.Entity { return moduleCapabilities.parent }

func (moduleCapabilities *BridgeDomainState_ModuleCapabilities) GetParentYangName() string { return "bridge-domain-state" }

// BridgeDomainState_ModuleCapabilities_Modules
// Collection of capabillity statements for hardware
// module in the system.
type BridgeDomainState_ModuleCapabilities_Modules struct {
    parent types.Entity
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

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetFilter() yfilter.YFilter { return modules.YFilter }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) SetFilter(yf yfilter.YFilter) { modules.YFilter = yf }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "max-mac-per-bd" { return "MaxMacPerBd" }
    if yname == "max-pdd-edge-bd" { return "MaxPddEdgeBd" }
    if yname == "max-bd" { return "MaxBd" }
    if yname == "max-ac-per-bd" { return "MaxAcPerBd" }
    if yname == "max-pw-per-bd" { return "MaxPwPerBd" }
    if yname == "max-vfi-per-bd" { return "MaxVfiPerBd" }
    if yname == "max-sh-group-per-bd" { return "MaxShGroupPerBd" }
    return ""
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetSegmentPath() string {
    return "modules" + "[name='" + fmt.Sprintf("%v", modules.Name) + "']"
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = modules.Name
    leafs["max-mac-per-bd"] = modules.MaxMacPerBd
    leafs["max-pdd-edge-bd"] = modules.MaxPddEdgeBd
    leafs["max-bd"] = modules.MaxBd
    leafs["max-ac-per-bd"] = modules.MaxAcPerBd
    leafs["max-pw-per-bd"] = modules.MaxPwPerBd
    leafs["max-vfi-per-bd"] = modules.MaxVfiPerBd
    leafs["max-sh-group-per-bd"] = modules.MaxShGroupPerBd
    return leafs
}

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetBundleName() string { return "cisco_ios_xe" }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetYangName() string { return "modules" }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) SetParent(parent types.Entity) { modules.parent = parent }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetParent() types.Entity { return modules.parent }

func (modules *BridgeDomainState_ModuleCapabilities_Modules) GetParentYangName() string { return "module-capabilities" }

// BridgeDomainState_BridgeDomains
// Bridge domain state data.
type BridgeDomainState_BridgeDomains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of bridge-domain state data. The type is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain.
    BridgeDomain []BridgeDomainState_BridgeDomains_BridgeDomain
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetFilter() yfilter.YFilter { return bridgeDomains.YFilter }

func (bridgeDomains *BridgeDomainState_BridgeDomains) SetFilter(yf yfilter.YFilter) { bridgeDomains.YFilter = yf }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetGoName(yname string) string {
    if yname == "bridge-domain" { return "BridgeDomain" }
    return ""
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetSegmentPath() string {
    return "bridge-domains"
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bridge-domain" {
        for _, c := range bridgeDomains.BridgeDomain {
            if bridgeDomains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain{}
        bridgeDomains.BridgeDomain = append(bridgeDomains.BridgeDomain, child)
        return &bridgeDomains.BridgeDomain[len(bridgeDomains.BridgeDomain)-1]
    }
    return nil
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bridgeDomains.BridgeDomain {
        children[bridgeDomains.BridgeDomain[i].GetSegmentPath()] = &bridgeDomains.BridgeDomain[i]
    }
    return children
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetYangName() string { return "bridge-domains" }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomains *BridgeDomainState_BridgeDomains) SetParent(parent types.Entity) { bridgeDomains.parent = parent }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetParent() types.Entity { return bridgeDomains.parent }

func (bridgeDomains *BridgeDomainState_BridgeDomains) GetParentYangName() string { return "bridge-domain-state" }

// BridgeDomainState_BridgeDomains_BridgeDomain
// Collection of bridge-domain state data.
type BridgeDomainState_BridgeDomains_BridgeDomain struct {
    parent types.Entity
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

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetFilter() yfilter.YFilter { return bridgeDomain.YFilter }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) SetFilter(yf yfilter.YFilter) { bridgeDomain.YFilter = yf }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "bd-state" { return "BdState" }
    if yname == "create-time" { return "CreateTime" }
    if yname == "last-status-change" { return "LastStatusChange" }
    if yname == "mac-limit-reached" { return "MacLimitReached" }
    if yname == "p2mp-pw-disabled" { return "P2MpPwDisabled" }
    if yname == "members" { return "Members" }
    return ""
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetSegmentPath() string {
    return "bridge-domain" + "[id='" + fmt.Sprintf("%v", bridgeDomain.Id) + "']"
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "members" {
        return &bridgeDomain.Members
    }
    return nil
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["members"] = &bridgeDomain.Members
    return children
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = bridgeDomain.Id
    leafs["bd-state"] = bridgeDomain.BdState
    leafs["create-time"] = bridgeDomain.CreateTime
    leafs["last-status-change"] = bridgeDomain.LastStatusChange
    leafs["mac-limit-reached"] = bridgeDomain.MacLimitReached
    leafs["p2mp-pw-disabled"] = bridgeDomain.P2MpPwDisabled
    return leafs
}

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetYangName() string { return "bridge-domain" }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) SetParent(parent types.Entity) { bridgeDomain.parent = parent }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetParent() types.Entity { return bridgeDomain.parent }

func (bridgeDomain *BridgeDomainState_BridgeDomains_BridgeDomain) GetParentYangName() string { return "bridge-domains" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members
// Collection of bridge-domain members.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members struct {
    parent types.Entity
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

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetFilter() yfilter.YFilter { return members.YFilter }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) SetFilter(yf yfilter.YFilter) { members.YFilter = yf }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetGoName(yname string) string {
    if yname == "ac-member" { return "AcMember" }
    if yname == "vfi-member" { return "VfiMember" }
    if yname == "access-pw-member" { return "AccessPwMember" }
    return ""
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetSegmentPath() string {
    return "members"
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ac-member" {
        for _, c := range members.AcMember {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember{}
        members.AcMember = append(members.AcMember, child)
        return &members.AcMember[len(members.AcMember)-1]
    }
    if childYangName == "vfi-member" {
        for _, c := range members.VfiMember {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember{}
        members.VfiMember = append(members.VfiMember, child)
        return &members.VfiMember[len(members.VfiMember)-1]
    }
    if childYangName == "access-pw-member" {
        for _, c := range members.AccessPwMember {
            if members.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember{}
        members.AccessPwMember = append(members.AccessPwMember, child)
        return &members.AccessPwMember[len(members.AccessPwMember)-1]
    }
    return nil
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range members.AcMember {
        children[members.AcMember[i].GetSegmentPath()] = &members.AcMember[i]
    }
    for i := range members.VfiMember {
        children[members.VfiMember[i].GetSegmentPath()] = &members.VfiMember[i]
    }
    for i := range members.AccessPwMember {
        children[members.AccessPwMember[i].GetSegmentPath()] = &members.AccessPwMember[i]
    }
    return children
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetBundleName() string { return "cisco_ios_xe" }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetYangName() string { return "members" }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) SetParent(parent types.Entity) { members.parent = parent }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetParent() types.Entity { return members.parent }

func (members *BridgeDomainState_BridgeDomains_BridgeDomain_Members) GetParentYangName() string { return "bridge-domain" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember
// List of attachment circuits for this bridge domains
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetFilter() yfilter.YFilter { return acMember.YFilter }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) SetFilter(yf yfilter.YFilter) { acMember.YFilter = yf }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "static-mac-count" { return "StaticMacCount" }
    if yname == "dai-stats" { return "DaiStats" }
    if yname == "ipsg-stats" { return "IpsgStats" }
    if yname == "storm-control" { return "StormControl" }
    return ""
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetSegmentPath() string {
    return "ac-member" + "[interface='" + fmt.Sprintf("%v", acMember.Interface) + "']"
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dai-stats" {
        return &acMember.DaiStats
    }
    if childYangName == "ipsg-stats" {
        return &acMember.IpsgStats
    }
    if childYangName == "storm-control" {
        return &acMember.StormControl
    }
    return nil
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dai-stats"] = &acMember.DaiStats
    children["ipsg-stats"] = &acMember.IpsgStats
    children["storm-control"] = &acMember.StormControl
    return children
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = acMember.Interface
    leafs["static-mac-count"] = acMember.StaticMacCount
    return leafs
}

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetBundleName() string { return "cisco_ios_xe" }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetYangName() string { return "ac-member" }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) SetParent(parent types.Entity) { acMember.parent = parent }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetParent() types.Entity { return acMember.parent }

func (acMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember) GetParentYangName() string { return "members" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats
// Dynamic ARP Inspection (DAI) statistics.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets dropped by interface due to DAI actions. The type is
    // interface{} with range: 0..18446744073709551615.
    PacketDrops interface{}

    // Number of bytes dropped by interface due to DAI actions. The type is
    // interface{} with range: 0..18446744073709551615.
    ByteDrops interface{}
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetFilter() yfilter.YFilter { return daiStats.YFilter }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) SetFilter(yf yfilter.YFilter) { daiStats.YFilter = yf }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetGoName(yname string) string {
    if yname == "packet-drops" { return "PacketDrops" }
    if yname == "byte-drops" { return "ByteDrops" }
    return ""
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetSegmentPath() string {
    return "dai-stats"
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-drops"] = daiStats.PacketDrops
    leafs["byte-drops"] = daiStats.ByteDrops
    return leafs
}

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetBundleName() string { return "cisco_ios_xe" }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetYangName() string { return "dai-stats" }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) SetParent(parent types.Entity) { daiStats.parent = parent }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetParent() types.Entity { return daiStats.parent }

func (daiStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_DaiStats) GetParentYangName() string { return "ac-member" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats
// IPSG stats.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets dropped by interface due to IPSG actions. The type is
    // interface{} with range: 0..18446744073709551615.
    PacketDrops interface{}

    // Number of bytes dropped by interface due to IPSG actions. The type is
    // interface{} with range: 0..18446744073709551615.
    ByteDrops interface{}
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetFilter() yfilter.YFilter { return ipsgStats.YFilter }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) SetFilter(yf yfilter.YFilter) { ipsgStats.YFilter = yf }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetGoName(yname string) string {
    if yname == "packet-drops" { return "PacketDrops" }
    if yname == "byte-drops" { return "ByteDrops" }
    return ""
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetSegmentPath() string {
    return "ipsg-stats"
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["packet-drops"] = ipsgStats.PacketDrops
    leafs["byte-drops"] = ipsgStats.ByteDrops
    return leafs
}

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetBundleName() string { return "cisco_ios_xe" }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetYangName() string { return "ipsg-stats" }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) SetParent(parent types.Entity) { ipsgStats.parent = parent }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetParent() types.Entity { return ipsgStats.parent }

func (ipsgStats *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_IpsgStats) GetParentYangName() string { return "ac-member" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl
// Storm control statistics.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Collection of packet drop statistics for ethernet traffic clasess. The type
    // is slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter.
    DropCounter []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetFilter() yfilter.YFilter { return stormControl.YFilter }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) SetFilter(yf yfilter.YFilter) { stormControl.YFilter = yf }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetGoName(yname string) string {
    if yname == "drop-counter" { return "DropCounter" }
    return ""
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetSegmentPath() string {
    return "storm-control"
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "drop-counter" {
        for _, c := range stormControl.DropCounter {
            if stormControl.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter{}
        stormControl.DropCounter = append(stormControl.DropCounter, child)
        return &stormControl.DropCounter[len(stormControl.DropCounter)-1]
    }
    return nil
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range stormControl.DropCounter {
        children[stormControl.DropCounter[i].GetSegmentPath()] = &stormControl.DropCounter[i]
    }
    return children
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetBundleName() string { return "cisco_ios_xe" }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetYangName() string { return "storm-control" }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) SetParent(parent types.Entity) { stormControl.parent = parent }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetParent() types.Entity { return stormControl.parent }

func (stormControl *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl) GetParentYangName() string { return "ac-member" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter
// Collection of packet drop statistics for ethernet traffic
// clasess.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter struct {
    parent types.Entity
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

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetFilter() yfilter.YFilter { return dropCounter.YFilter }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) SetFilter(yf yfilter.YFilter) { dropCounter.YFilter = yf }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "packet-drops" { return "PacketDrops" }
    if yname == "octate-drops" { return "OctateDrops" }
    return ""
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetSegmentPath() string {
    return "drop-counter" + "[traffic-class='" + fmt.Sprintf("%v", dropCounter.TrafficClass) + "']"
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = dropCounter.TrafficClass
    leafs["packet-drops"] = dropCounter.PacketDrops
    leafs["octate-drops"] = dropCounter.OctateDrops
    return leafs
}

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetBundleName() string { return "cisco_ios_xe" }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetYangName() string { return "drop-counter" }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) SetParent(parent types.Entity) { dropCounter.parent = parent }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetParent() types.Entity { return dropCounter.parent }

func (dropCounter *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AcMember_StormControl_DropCounter) GetParentYangName() string { return "storm-control" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember
// Reference to an instance of Bridge domain Virtual
// Forwarding Instance (VFI) name.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge domain memeber interface name. The type is
    // string. Refers to ietf_interfaces.InterfacesState_Interface_Name
    Interface interface{}

    // Flooding operational status.
    Flooding BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetFilter() yfilter.YFilter { return vfiMember.YFilter }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) SetFilter(yf yfilter.YFilter) { vfiMember.YFilter = yf }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "flooding" { return "Flooding" }
    return ""
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetSegmentPath() string {
    return "vfi-member" + "[interface='" + fmt.Sprintf("%v", vfiMember.Interface) + "']"
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flooding" {
        return &vfiMember.Flooding
    }
    return nil
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flooding"] = &vfiMember.Flooding
    return children
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = vfiMember.Interface
    return leafs
}

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetBundleName() string { return "cisco_ios_xe" }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetYangName() string { return "vfi-member" }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) SetParent(parent types.Entity) { vfiMember.parent = parent }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetParent() types.Entity { return vfiMember.parent }

func (vfiMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember) GetParentYangName() string { return "members" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding
// Flooding operational status
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status.
    Status []BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetFilter() yfilter.YFilter { return flooding.YFilter }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) SetFilter(yf yfilter.YFilter) { flooding.YFilter = yf }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetGoName(yname string) string {
    if yname == "status" { return "Status" }
    return ""
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetSegmentPath() string {
    return "flooding"
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "status" {
        for _, c := range flooding.Status {
            if flooding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status{}
        flooding.Status = append(flooding.Status, child)
        return &flooding.Status[len(flooding.Status)-1]
    }
    return nil
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flooding.Status {
        children[flooding.Status[i].GetSegmentPath()] = &flooding.Status[i]
    }
    return children
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetBundleName() string { return "cisco_ios_xe" }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetYangName() string { return "flooding" }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) SetParent(parent types.Entity) { flooding.parent = parent }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetParent() types.Entity { return flooding.parent }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding) GetParentYangName() string { return "vfi-member" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type. The
    // type is EthTrafficClass.
    TrafficClass interface{}

    // This leaf indicates if flooding is enabled for corresponding traffic class.
    // The type is bool.
    Enabled interface{}
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetSegmentPath() string {
    return "status" + "[traffic-class='" + fmt.Sprintf("%v", status.TrafficClass) + "']"
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = status.TrafficClass
    leafs["enabled"] = status.Enabled
    return leafs
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetBundleName() string { return "cisco_ios_xe" }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetYangName() string { return "status" }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetParent() types.Entity { return status.parent }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_VfiMember_Flooding_Status) GetParentYangName() string { return "flooding" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember
// Collection of access pseudowire members of the bridge
// domain.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to peer ip address of a pseudowire
    // instance. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VcPeerAddress interface{}

    // This attribute is a key. Reference to vc-id of a pseudowire instance. The
    // type is string with range: 0..4294967295. Refers to
    // cisco_pw.PseudowireState_Pseudowires_VcId
    VcId interface{}

    // Flooding operational status.
    Flooding BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetFilter() yfilter.YFilter { return accessPwMember.YFilter }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) SetFilter(yf yfilter.YFilter) { accessPwMember.YFilter = yf }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetGoName(yname string) string {
    if yname == "vc-peer-address" { return "VcPeerAddress" }
    if yname == "vc-id" { return "VcId" }
    if yname == "flooding" { return "Flooding" }
    return ""
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetSegmentPath() string {
    return "access-pw-member" + "[vc-peer-address='" + fmt.Sprintf("%v", accessPwMember.VcPeerAddress) + "']" + "[vc-id='" + fmt.Sprintf("%v", accessPwMember.VcId) + "']"
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flooding" {
        return &accessPwMember.Flooding
    }
    return nil
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flooding"] = &accessPwMember.Flooding
    return children
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vc-peer-address"] = accessPwMember.VcPeerAddress
    leafs["vc-id"] = accessPwMember.VcId
    return leafs
}

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetBundleName() string { return "cisco_ios_xe" }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetYangName() string { return "access-pw-member" }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) SetParent(parent types.Entity) { accessPwMember.parent = parent }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetParent() types.Entity { return accessPwMember.parent }

func (accessPwMember *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember) GetParentYangName() string { return "members" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding
// Flooding operational status
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A collection of storm control threshold configuration entries. The type is
    // slice of
    // BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status.
    Status []BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetFilter() yfilter.YFilter { return flooding.YFilter }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) SetFilter(yf yfilter.YFilter) { flooding.YFilter = yf }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetGoName(yname string) string {
    if yname == "status" { return "Status" }
    return ""
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetSegmentPath() string {
    return "flooding"
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "status" {
        for _, c := range flooding.Status {
            if flooding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status{}
        flooding.Status = append(flooding.Status, child)
        return &flooding.Status[len(flooding.Status)-1]
    }
    return nil
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flooding.Status {
        children[flooding.Status[i].GetSegmentPath()] = &flooding.Status[i]
    }
    return children
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetBundleName() string { return "cisco_ios_xe" }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetYangName() string { return "flooding" }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) SetParent(parent types.Entity) { flooding.parent = parent }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetParent() types.Entity { return flooding.parent }

func (flooding *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding) GetParentYangName() string { return "access-pw-member" }

// BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status
// A collection of storm control threshold configuration
// entries.
type BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf identifies a ethernet traffic type. The
    // type is EthTrafficClass.
    TrafficClass interface{}

    // This leaf indicates if flooding is enabled for corresponding traffic class.
    // The type is bool.
    Enabled interface{}
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetGoName(yname string) string {
    if yname == "traffic-class" { return "TrafficClass" }
    if yname == "enabled" { return "Enabled" }
    return ""
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetSegmentPath() string {
    return "status" + "[traffic-class='" + fmt.Sprintf("%v", status.TrafficClass) + "']"
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["traffic-class"] = status.TrafficClass
    leafs["enabled"] = status.Enabled
    return leafs
}

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetBundleName() string { return "cisco_ios_xe" }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetYangName() string { return "status" }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetParent() types.Entity { return status.parent }

func (status *BridgeDomainState_BridgeDomains_BridgeDomain_Members_AccessPwMember_Flooding_Status) GetParentYangName() string { return "flooding" }

// BridgeDomainState_MacTable
// This list contains mac-address entries for bridge
// domains.
type BridgeDomainState_MacTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bridge-domain name where MAC entry is learnt. The
    // type is string.
    BdId interface{}

    // This attribute is a key. MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
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

func (macTable *BridgeDomainState_MacTable) GetFilter() yfilter.YFilter { return macTable.YFilter }

func (macTable *BridgeDomainState_MacTable) SetFilter(yf yfilter.YFilter) { macTable.YFilter = yf }

func (macTable *BridgeDomainState_MacTable) GetGoName(yname string) string {
    if yname == "bd-id" { return "BdId" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "mac-type" { return "MacType" }
    if yname == "interface" { return "Interface" }
    if yname == "secure-mac" { return "SecureMac" }
    if yname == "ntfy-mac" { return "NtfyMac" }
    if yname == "age" { return "Age" }
    if yname == "location" { return "Location" }
    return ""
}

func (macTable *BridgeDomainState_MacTable) GetSegmentPath() string {
    return "mac-table" + "[bd-id='" + fmt.Sprintf("%v", macTable.BdId) + "']" + "[mac-address='" + fmt.Sprintf("%v", macTable.MacAddress) + "']"
}

func (macTable *BridgeDomainState_MacTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macTable *BridgeDomainState_MacTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macTable *BridgeDomainState_MacTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bd-id"] = macTable.BdId
    leafs["mac-address"] = macTable.MacAddress
    leafs["mac-type"] = macTable.MacType
    leafs["interface"] = macTable.Interface
    leafs["secure-mac"] = macTable.SecureMac
    leafs["ntfy-mac"] = macTable.NtfyMac
    leafs["age"] = macTable.Age
    leafs["location"] = macTable.Location
    return leafs
}

func (macTable *BridgeDomainState_MacTable) GetBundleName() string { return "cisco_ios_xe" }

func (macTable *BridgeDomainState_MacTable) GetYangName() string { return "mac-table" }

func (macTable *BridgeDomainState_MacTable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (macTable *BridgeDomainState_MacTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (macTable *BridgeDomainState_MacTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (macTable *BridgeDomainState_MacTable) SetParent(parent types.Entity) { macTable.parent = parent }

func (macTable *BridgeDomainState_MacTable) GetParent() types.Entity { return macTable.parent }

func (macTable *BridgeDomainState_MacTable) GetParentYangName() string { return "bridge-domain-state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearBridgeDomain_Input

    
    Output ClearBridgeDomain_Output
}

func (clearBridgeDomain *ClearBridgeDomain) GetFilter() yfilter.YFilter { return clearBridgeDomain.YFilter }

func (clearBridgeDomain *ClearBridgeDomain) SetFilter(yf yfilter.YFilter) { clearBridgeDomain.YFilter = yf }

func (clearBridgeDomain *ClearBridgeDomain) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (clearBridgeDomain *ClearBridgeDomain) GetSegmentPath() string {
    return "cisco-bridge-domain:clear-bridge-domain"
}

func (clearBridgeDomain *ClearBridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearBridgeDomain.Input
    }
    if childYangName == "output" {
        return &clearBridgeDomain.Output
    }
    return nil
}

func (clearBridgeDomain *ClearBridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearBridgeDomain.Input
    children["output"] = &clearBridgeDomain.Output
    return children
}

func (clearBridgeDomain *ClearBridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearBridgeDomain *ClearBridgeDomain) GetBundleName() string { return "cisco_ios_xe" }

func (clearBridgeDomain *ClearBridgeDomain) GetYangName() string { return "clear-bridge-domain" }

func (clearBridgeDomain *ClearBridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (clearBridgeDomain *ClearBridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (clearBridgeDomain *ClearBridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (clearBridgeDomain *ClearBridgeDomain) SetParent(parent types.Entity) { clearBridgeDomain.parent = parent }

func (clearBridgeDomain *ClearBridgeDomain) GetParent() types.Entity { return clearBridgeDomain.parent }

func (clearBridgeDomain *ClearBridgeDomain) GetParentYangName() string { return "cisco-bridge-domain" }

// ClearBridgeDomain_Input
type ClearBridgeDomain_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear all bridge-domains configured on the device. The type is interface{}.
    All interface{}

    // Clear a single bridge-domain. The type is string.
    BdId interface{}

    // Clears all bridge-domains under this bridge-group. The type is string.
    BgId interface{}
}

func (input *ClearBridgeDomain_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearBridgeDomain_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearBridgeDomain_Input) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "bd-id" { return "BdId" }
    if yname == "bg-id" { return "BgId" }
    return ""
}

func (input *ClearBridgeDomain_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearBridgeDomain_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *ClearBridgeDomain_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *ClearBridgeDomain_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = input.All
    leafs["bd-id"] = input.BdId
    leafs["bg-id"] = input.BgId
    return leafs
}

func (input *ClearBridgeDomain_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *ClearBridgeDomain_Input) GetYangName() string { return "input" }

func (input *ClearBridgeDomain_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *ClearBridgeDomain_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *ClearBridgeDomain_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *ClearBridgeDomain_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearBridgeDomain_Input) GetParent() types.Entity { return input.parent }

func (input *ClearBridgeDomain_Input) GetParentYangName() string { return "clear-bridge-domain" }

// ClearBridgeDomain_Output
type ClearBridgeDomain_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *ClearBridgeDomain_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *ClearBridgeDomain_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *ClearBridgeDomain_Output) GetGoName(yname string) string {
    if yname == "errstr" { return "Errstr" }
    return ""
}

func (output *ClearBridgeDomain_Output) GetSegmentPath() string {
    return "output"
}

func (output *ClearBridgeDomain_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *ClearBridgeDomain_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *ClearBridgeDomain_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["errstr"] = output.Errstr
    return leafs
}

func (output *ClearBridgeDomain_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *ClearBridgeDomain_Output) GetYangName() string { return "output" }

func (output *ClearBridgeDomain_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *ClearBridgeDomain_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *ClearBridgeDomain_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *ClearBridgeDomain_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *ClearBridgeDomain_Output) GetParent() types.Entity { return output.parent }

func (output *ClearBridgeDomain_Output) GetParentYangName() string { return "clear-bridge-domain" }

// ClearMacAddress
// This RPC allows to clear dynamically learnt mac-address
// entries from the mac-address table.
type ClearMacAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearMacAddress_Input

    
    Output ClearMacAddress_Output
}

func (clearMacAddress *ClearMacAddress) GetFilter() yfilter.YFilter { return clearMacAddress.YFilter }

func (clearMacAddress *ClearMacAddress) SetFilter(yf yfilter.YFilter) { clearMacAddress.YFilter = yf }

func (clearMacAddress *ClearMacAddress) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (clearMacAddress *ClearMacAddress) GetSegmentPath() string {
    return "cisco-bridge-domain:clear-mac-address"
}

func (clearMacAddress *ClearMacAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearMacAddress.Input
    }
    if childYangName == "output" {
        return &clearMacAddress.Output
    }
    return nil
}

func (clearMacAddress *ClearMacAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearMacAddress.Input
    children["output"] = &clearMacAddress.Output
    return children
}

func (clearMacAddress *ClearMacAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearMacAddress *ClearMacAddress) GetBundleName() string { return "cisco_ios_xe" }

func (clearMacAddress *ClearMacAddress) GetYangName() string { return "clear-mac-address" }

func (clearMacAddress *ClearMacAddress) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (clearMacAddress *ClearMacAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (clearMacAddress *ClearMacAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (clearMacAddress *ClearMacAddress) SetParent(parent types.Entity) { clearMacAddress.parent = parent }

func (clearMacAddress *ClearMacAddress) GetParent() types.Entity { return clearMacAddress.parent }

func (clearMacAddress *ClearMacAddress) GetParentYangName() string { return "cisco-bridge-domain" }

// ClearMacAddress_Input
type ClearMacAddress_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reference to an interface. Clear mac-addesses learnt by by this interface.
    // The type is string. Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Clear a specific mac-address entry from the mac-table. The type is string
    // with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Clear mac-address entries for given bridge-domain(s).
    BridgeDomain ClearMacAddress_Input_BridgeDomain
}

func (input *ClearMacAddress_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearMacAddress_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearMacAddress_Input) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "mac-address" { return "MacAddress" }
    if yname == "bridge-domain" { return "BridgeDomain" }
    return ""
}

func (input *ClearMacAddress_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearMacAddress_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bridge-domain" {
        return &input.BridgeDomain
    }
    return nil
}

func (input *ClearMacAddress_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bridge-domain"] = &input.BridgeDomain
    return children
}

func (input *ClearMacAddress_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = input.Interface
    leafs["mac-address"] = input.MacAddress
    return leafs
}

func (input *ClearMacAddress_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *ClearMacAddress_Input) GetYangName() string { return "input" }

func (input *ClearMacAddress_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *ClearMacAddress_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *ClearMacAddress_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *ClearMacAddress_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearMacAddress_Input) GetParent() types.Entity { return input.parent }

func (input *ClearMacAddress_Input) GetParentYangName() string { return "clear-mac-address" }

// ClearMacAddress_Input_BridgeDomain
// Clear mac-address entries for given bridge-domain(s).
type ClearMacAddress_Input_BridgeDomain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bridge-domain identifier, clear all mac-address entries dynamically learnt
    // on this bridge-domain. The type is string. This attribute is mandatory.
    BdId interface{}

    // Bridge-group identifier, clear all mac-address entries from all
    // bridge-domains under this bridge-group. The type is string. This attribute
    // is mandatory.
    BgId interface{}
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetFilter() yfilter.YFilter { return bridgeDomain.YFilter }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) SetFilter(yf yfilter.YFilter) { bridgeDomain.YFilter = yf }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetGoName(yname string) string {
    if yname == "bd-id" { return "BdId" }
    if yname == "bg-id" { return "BgId" }
    return ""
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetSegmentPath() string {
    return "bridge-domain"
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bd-id"] = bridgeDomain.BdId
    leafs["bg-id"] = bridgeDomain.BgId
    return leafs
}

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetBundleName() string { return "cisco_ios_xe" }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetYangName() string { return "bridge-domain" }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) SetParent(parent types.Entity) { bridgeDomain.parent = parent }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetParent() types.Entity { return bridgeDomain.parent }

func (bridgeDomain *ClearMacAddress_Input_BridgeDomain) GetParentYangName() string { return "input" }

// ClearMacAddress_Output
type ClearMacAddress_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *ClearMacAddress_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *ClearMacAddress_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *ClearMacAddress_Output) GetGoName(yname string) string {
    if yname == "errstr" { return "Errstr" }
    return ""
}

func (output *ClearMacAddress_Output) GetSegmentPath() string {
    return "output"
}

func (output *ClearMacAddress_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *ClearMacAddress_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *ClearMacAddress_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["errstr"] = output.Errstr
    return leafs
}

func (output *ClearMacAddress_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *ClearMacAddress_Output) GetYangName() string { return "output" }

func (output *ClearMacAddress_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *ClearMacAddress_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *ClearMacAddress_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *ClearMacAddress_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *ClearMacAddress_Output) GetParent() types.Entity { return output.parent }

func (output *ClearMacAddress_Output) GetParentYangName() string { return "clear-mac-address" }

// CreateParameterizedBridgeDomains
// Create bridge domains automatically from user defined
// parameters.
type CreateParameterizedBridgeDomains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CreateParameterizedBridgeDomains_Input

    
    Output CreateParameterizedBridgeDomains_Output
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetFilter() yfilter.YFilter { return createParameterizedBridgeDomains.YFilter }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) SetFilter(yf yfilter.YFilter) { createParameterizedBridgeDomains.YFilter = yf }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetSegmentPath() string {
    return "cisco-bridge-domain:create-parameterized-bridge-domains"
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &createParameterizedBridgeDomains.Input
    }
    if childYangName == "output" {
        return &createParameterizedBridgeDomains.Output
    }
    return nil
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &createParameterizedBridgeDomains.Input
    children["output"] = &createParameterizedBridgeDomains.Output
    return children
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetBundleName() string { return "cisco_ios_xe" }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetYangName() string { return "create-parameterized-bridge-domains" }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) SetParent(parent types.Entity) { createParameterizedBridgeDomains.parent = parent }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetParent() types.Entity { return createParameterizedBridgeDomains.parent }

func (createParameterizedBridgeDomains *CreateParameterizedBridgeDomains) GetParentYangName() string { return "cisco-bridge-domain" }

// CreateParameterizedBridgeDomains_Input
type CreateParameterizedBridgeDomains_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Parameter for automatic bridge domain creation. The type is Parameter.
    Parameter interface{}

    // Bridge-domain member interface. The type is slice of
    // CreateParameterizedBridgeDomains_Input_Member.
    Member []CreateParameterizedBridgeDomains_Input_Member
}

func (input *CreateParameterizedBridgeDomains_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CreateParameterizedBridgeDomains_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CreateParameterizedBridgeDomains_Input) GetGoName(yname string) string {
    if yname == "parameter" { return "Parameter" }
    if yname == "member" { return "Member" }
    return ""
}

func (input *CreateParameterizedBridgeDomains_Input) GetSegmentPath() string {
    return "input"
}

func (input *CreateParameterizedBridgeDomains_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member" {
        for _, c := range input.Member {
            if input.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CreateParameterizedBridgeDomains_Input_Member{}
        input.Member = append(input.Member, child)
        return &input.Member[len(input.Member)-1]
    }
    return nil
}

func (input *CreateParameterizedBridgeDomains_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range input.Member {
        children[input.Member[i].GetSegmentPath()] = &input.Member[i]
    }
    return children
}

func (input *CreateParameterizedBridgeDomains_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["parameter"] = input.Parameter
    return leafs
}

func (input *CreateParameterizedBridgeDomains_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *CreateParameterizedBridgeDomains_Input) GetYangName() string { return "input" }

func (input *CreateParameterizedBridgeDomains_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *CreateParameterizedBridgeDomains_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *CreateParameterizedBridgeDomains_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *CreateParameterizedBridgeDomains_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CreateParameterizedBridgeDomains_Input) GetParent() types.Entity { return input.parent }

func (input *CreateParameterizedBridgeDomains_Input) GetParentYangName() string { return "create-parameterized-bridge-domains" }

// CreateParameterizedBridgeDomains_Input_Member
// Bridge-domain member interface.
type CreateParameterizedBridgeDomains_Input_Member struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Reference to an interface instance which is
    // configured to be part of this bridge domain. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetFilter() yfilter.YFilter { return member.YFilter }

func (member *CreateParameterizedBridgeDomains_Input_Member) SetFilter(yf yfilter.YFilter) { member.YFilter = yf }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetSegmentPath() string {
    return "member" + "[interface='" + fmt.Sprintf("%v", member.Interface) + "']"
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = member.Interface
    return leafs
}

func (member *CreateParameterizedBridgeDomains_Input_Member) GetBundleName() string { return "cisco_ios_xe" }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetYangName() string { return "member" }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (member *CreateParameterizedBridgeDomains_Input_Member) SetParent(parent types.Entity) { member.parent = parent }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetParent() types.Entity { return member.parent }

func (member *CreateParameterizedBridgeDomains_Input_Member) GetParentYangName() string { return "input" }

// CreateParameterizedBridgeDomains_Input_Parameter represents Parameter for automatic bridge domain creation.
type CreateParameterizedBridgeDomains_Input_Parameter string

const (
    // Create bridge-domains from vlan encapsulations of the
    // member interfaces.
    CreateParameterizedBridgeDomains_Input_Parameter_vlan CreateParameterizedBridgeDomains_Input_Parameter = "vlan"
)

// CreateParameterizedBridgeDomains_Output
type CreateParameterizedBridgeDomains_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error message from the device if RPC failed. The type is string.
    Errstr interface{}
}

func (output *CreateParameterizedBridgeDomains_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *CreateParameterizedBridgeDomains_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *CreateParameterizedBridgeDomains_Output) GetGoName(yname string) string {
    if yname == "errstr" { return "Errstr" }
    return ""
}

func (output *CreateParameterizedBridgeDomains_Output) GetSegmentPath() string {
    return "output"
}

func (output *CreateParameterizedBridgeDomains_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *CreateParameterizedBridgeDomains_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *CreateParameterizedBridgeDomains_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["errstr"] = output.Errstr
    return leafs
}

func (output *CreateParameterizedBridgeDomains_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *CreateParameterizedBridgeDomains_Output) GetYangName() string { return "output" }

func (output *CreateParameterizedBridgeDomains_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *CreateParameterizedBridgeDomains_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *CreateParameterizedBridgeDomains_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *CreateParameterizedBridgeDomains_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *CreateParameterizedBridgeDomains_Output) GetParent() types.Entity { return output.parent }

func (output *CreateParameterizedBridgeDomains_Output) GetParentYangName() string { return "create-parameterized-bridge-domains" }

