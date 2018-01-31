// This module contains a collection of YANG definitions
// for Cisco IOS-XR l2-eth-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ethernet-features: Ethernet Features Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-l2vpn-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package l2_eth_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package l2_eth_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-l2-eth-infra-cfg ethernet-features}", reflect.TypeOf(EthernetFeatures{}))
    ydk.RegisterEntity("Cisco-IOS-XR-l2-eth-infra-cfg:ethernet-features", reflect.TypeOf(EthernetFeatures{}))
}

// EgressFiltering represents Egress filtering
type EgressFiltering string

const (
    // Strict Egress Filtering
    EgressFiltering_egress_filtering_type_strict EgressFiltering = "egress-filtering-type-strict"

    // Egress Filtering Disabled
    EgressFiltering_egress_filtering_type_disable EgressFiltering = "egress-filtering-type-disable"

    // Default Egress Filtering Behavior
    EgressFiltering_egress_filtering_type_default EgressFiltering = "egress-filtering-type-default"
)

// L2ProtocolName represents L2 protocol name
type L2ProtocolName string

const (
    // CDP
    L2ProtocolName_cdp L2ProtocolName = "cdp"

    // STP
    L2ProtocolName_stp L2ProtocolName = "stp"

    // VTP
    L2ProtocolName_vtp L2ProtocolName = "vtp"

    // PVST+
    L2ProtocolName_pvst L2ProtocolName = "pvst"

    // CDP, PVST+, STP, and VTP
    L2ProtocolName_cpsv L2ProtocolName = "cpsv"
)

// Filtering represents Filtering
type Filtering string

const (
    // C-Vlan ingress frame filtering (Table 8-1 of
    // 802.1ad standard)
    Filtering_filtering_type_dot1q Filtering = "filtering-type-dot1q"

    // S-Vlan ingress frame filtering (Table 8-2 of
    // 802.1ad standard)
    Filtering_filtering_type_dot1ad Filtering = "filtering-type-dot1ad"
)

// L2ProtocolMode represents L2 protocol mode
type L2ProtocolMode string

const (
    // Forward packets transparently
    L2ProtocolMode_forward L2ProtocolMode = "forward"

    // Drop the protocol's packets
    L2ProtocolMode_drop L2ProtocolMode = "drop"

    // Tunnel ingress frames, untunnel egress frames
    L2ProtocolMode_tunnel L2ProtocolMode = "tunnel"

    // Tunnel egress frames, untunnel ingress frames
    L2ProtocolMode_reverse_tunnel L2ProtocolMode = "reverse-tunnel"
)

// EthernetFeatures
// Ethernet Features Configuration
type EthernetFeatures struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Egress Filtering Configuration.
    EgressFiltering EthernetFeatures_EgressFiltering

    // CFM global configuration.
    Cfm EthernetFeatures_Cfm

    // Ethernet Link OAM Global Configuration.
    EtherLinkOam EthernetFeatures_EtherLinkOam
}

func (ethernetFeatures *EthernetFeatures) GetFilter() yfilter.YFilter { return ethernetFeatures.YFilter }

func (ethernetFeatures *EthernetFeatures) SetFilter(yf yfilter.YFilter) { ethernetFeatures.YFilter = yf }

func (ethernetFeatures *EthernetFeatures) GetGoName(yname string) string {
    if yname == "egress-filtering" { return "EgressFiltering" }
    if yname == "Cisco-IOS-XR-ethernet-cfm-cfg:cfm" { return "Cfm" }
    if yname == "Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam" { return "EtherLinkOam" }
    return ""
}

func (ethernetFeatures *EthernetFeatures) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2-eth-infra-cfg:ethernet-features"
}

func (ethernetFeatures *EthernetFeatures) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "egress-filtering" {
        return &ethernetFeatures.EgressFiltering
    }
    if childYangName == "Cisco-IOS-XR-ethernet-cfm-cfg:cfm" {
        return &ethernetFeatures.Cfm
    }
    if childYangName == "Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam" {
        return &ethernetFeatures.EtherLinkOam
    }
    return nil
}

func (ethernetFeatures *EthernetFeatures) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["egress-filtering"] = &ethernetFeatures.EgressFiltering
    children["Cisco-IOS-XR-ethernet-cfm-cfg:cfm"] = &ethernetFeatures.Cfm
    children["Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam"] = &ethernetFeatures.EtherLinkOam
    return children
}

func (ethernetFeatures *EthernetFeatures) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ethernetFeatures *EthernetFeatures) GetBundleName() string { return "cisco_ios_xr" }

func (ethernetFeatures *EthernetFeatures) GetYangName() string { return "ethernet-features" }

func (ethernetFeatures *EthernetFeatures) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ethernetFeatures *EthernetFeatures) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ethernetFeatures *EthernetFeatures) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ethernetFeatures *EthernetFeatures) SetParent(parent types.Entity) { ethernetFeatures.parent = parent }

func (ethernetFeatures *EthernetFeatures) GetParent() types.Entity { return ethernetFeatures.parent }

func (ethernetFeatures *EthernetFeatures) GetParentYangName() string { return "Cisco-IOS-XR-l2-eth-infra-cfg" }

// EthernetFeatures_EgressFiltering
// Egress Filtering Configuration
type EthernetFeatures_EgressFiltering struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether Egress Filtering is on by default. The type is interface{}.
    EgressFilteringDefaultOn interface{}
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetFilter() yfilter.YFilter { return egressFiltering.YFilter }

func (egressFiltering *EthernetFeatures_EgressFiltering) SetFilter(yf yfilter.YFilter) { egressFiltering.YFilter = yf }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetGoName(yname string) string {
    if yname == "egress-filtering-default-on" { return "EgressFilteringDefaultOn" }
    return ""
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetSegmentPath() string {
    return "egress-filtering"
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["egress-filtering-default-on"] = egressFiltering.EgressFilteringDefaultOn
    return leafs
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetBundleName() string { return "cisco_ios_xr" }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetYangName() string { return "egress-filtering" }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (egressFiltering *EthernetFeatures_EgressFiltering) SetParent(parent types.Entity) { egressFiltering.parent = parent }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetParent() types.Entity { return egressFiltering.parent }

func (egressFiltering *EthernetFeatures_EgressFiltering) GetParentYangName() string { return "ethernet-features" }

// EthernetFeatures_Cfm
// CFM global configuration
type EthernetFeatures_Cfm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable processing of Ethernet SLA packets on nV Satellite devices. The
    // type is interface{}.
    NvSatelliteSlaProcessingDisable interface{}

    // Traceroute Cache Configuration.
    TracerouteCache EthernetFeatures_Cfm_TracerouteCache

    // Domain-specific global configuration.
    Domains EthernetFeatures_Cfm_Domains
}

func (cfm *EthernetFeatures_Cfm) GetFilter() yfilter.YFilter { return cfm.YFilter }

func (cfm *EthernetFeatures_Cfm) SetFilter(yf yfilter.YFilter) { cfm.YFilter = yf }

func (cfm *EthernetFeatures_Cfm) GetGoName(yname string) string {
    if yname == "nv-satellite-sla-processing-disable" { return "NvSatelliteSlaProcessingDisable" }
    if yname == "traceroute-cache" { return "TracerouteCache" }
    if yname == "domains" { return "Domains" }
    return ""
}

func (cfm *EthernetFeatures_Cfm) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-cfm-cfg:cfm"
}

func (cfm *EthernetFeatures_Cfm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traceroute-cache" {
        return &cfm.TracerouteCache
    }
    if childYangName == "domains" {
        return &cfm.Domains
    }
    return nil
}

func (cfm *EthernetFeatures_Cfm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traceroute-cache"] = &cfm.TracerouteCache
    children["domains"] = &cfm.Domains
    return children
}

func (cfm *EthernetFeatures_Cfm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nv-satellite-sla-processing-disable"] = cfm.NvSatelliteSlaProcessingDisable
    return leafs
}

func (cfm *EthernetFeatures_Cfm) GetBundleName() string { return "cisco_ios_xr" }

func (cfm *EthernetFeatures_Cfm) GetYangName() string { return "cfm" }

func (cfm *EthernetFeatures_Cfm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfm *EthernetFeatures_Cfm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfm *EthernetFeatures_Cfm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfm *EthernetFeatures_Cfm) SetParent(parent types.Entity) { cfm.parent = parent }

func (cfm *EthernetFeatures_Cfm) GetParent() types.Entity { return cfm.parent }

func (cfm *EthernetFeatures_Cfm) GetParentYangName() string { return "ethernet-features" }

// EthernetFeatures_Cfm_TracerouteCache
// Traceroute Cache Configuration
type EthernetFeatures_Cfm_TracerouteCache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hold Time in minutes. The type is interface{} with range: 1..525600. The
    // default value is 100.
    HoldTime interface{}

    // Cache Size limit (number of replies). The type is interface{} with range:
    // 1..4294967295. The default value is 100.
    CacheSize interface{}
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetFilter() yfilter.YFilter { return tracerouteCache.YFilter }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) SetFilter(yf yfilter.YFilter) { tracerouteCache.YFilter = yf }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetGoName(yname string) string {
    if yname == "hold-time" { return "HoldTime" }
    if yname == "cache-size" { return "CacheSize" }
    return ""
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetSegmentPath() string {
    return "traceroute-cache"
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hold-time"] = tracerouteCache.HoldTime
    leafs["cache-size"] = tracerouteCache.CacheSize
    return leafs
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetBundleName() string { return "cisco_ios_xr" }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetYangName() string { return "traceroute-cache" }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) SetParent(parent types.Entity) { tracerouteCache.parent = parent }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetParent() types.Entity { return tracerouteCache.parent }

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetParentYangName() string { return "cfm" }

// EthernetFeatures_Cfm_Domains
// Domain-specific global configuration
type EthernetFeatures_Cfm_Domains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular Maintenance Domain. The type is slice of
    // EthernetFeatures_Cfm_Domains_Domain.
    Domain []EthernetFeatures_Cfm_Domains_Domain
}

func (domains *EthernetFeatures_Cfm_Domains) GetFilter() yfilter.YFilter { return domains.YFilter }

func (domains *EthernetFeatures_Cfm_Domains) SetFilter(yf yfilter.YFilter) { domains.YFilter = yf }

func (domains *EthernetFeatures_Cfm_Domains) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    return ""
}

func (domains *EthernetFeatures_Cfm_Domains) GetSegmentPath() string {
    return "domains"
}

func (domains *EthernetFeatures_Cfm_Domains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "domain" {
        for _, c := range domains.Domain {
            if domains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetFeatures_Cfm_Domains_Domain{}
        domains.Domain = append(domains.Domain, child)
        return &domains.Domain[len(domains.Domain)-1]
    }
    return nil
}

func (domains *EthernetFeatures_Cfm_Domains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range domains.Domain {
        children[domains.Domain[i].GetSegmentPath()] = &domains.Domain[i]
    }
    return children
}

func (domains *EthernetFeatures_Cfm_Domains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (domains *EthernetFeatures_Cfm_Domains) GetBundleName() string { return "cisco_ios_xr" }

func (domains *EthernetFeatures_Cfm_Domains) GetYangName() string { return "domains" }

func (domains *EthernetFeatures_Cfm_Domains) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domains *EthernetFeatures_Cfm_Domains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domains *EthernetFeatures_Cfm_Domains) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domains *EthernetFeatures_Cfm_Domains) SetParent(parent types.Entity) { domains.parent = parent }

func (domains *EthernetFeatures_Cfm_Domains) GetParent() types.Entity { return domains.parent }

func (domains *EthernetFeatures_Cfm_Domains) GetParentYangName() string { return "cfm" }

// EthernetFeatures_Cfm_Domains_Domain
// Configuration for a particular Maintenance
// Domain
type EthernetFeatures_Cfm_Domains_Domain struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // Service-specific global configuration.
    Services EthernetFeatures_Cfm_Domains_Domain_Services

    // Fundamental properties of the domain.
    DomainProperties EthernetFeatures_Cfm_Domains_Domain_DomainProperties
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetFilter() yfilter.YFilter { return domain.YFilter }

func (domain *EthernetFeatures_Cfm_Domains_Domain) SetFilter(yf yfilter.YFilter) { domain.YFilter = yf }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "services" { return "Services" }
    if yname == "domain-properties" { return "DomainProperties" }
    return ""
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetSegmentPath() string {
    return "domain" + "[domain='" + fmt.Sprintf("%v", domain.Domain) + "']"
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "services" {
        return &domain.Services
    }
    if childYangName == "domain-properties" {
        return &domain.DomainProperties
    }
    return nil
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["services"] = &domain.Services
    children["domain-properties"] = &domain.DomainProperties
    return children
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = domain.Domain
    return leafs
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetBundleName() string { return "cisco_ios_xr" }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetYangName() string { return "domain" }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domain *EthernetFeatures_Cfm_Domains_Domain) SetParent(parent types.Entity) { domain.parent = parent }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetParent() types.Entity { return domain.parent }

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetParentYangName() string { return "domains" }

// EthernetFeatures_Cfm_Domains_Domain_Services
// Service-specific global configuration
type EthernetFeatures_Cfm_Domains_Domain_Services struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configuration for a particular Service (Maintenance Association). The type
    // is slice of EthernetFeatures_Cfm_Domains_Domain_Services_Service.
    Service []EthernetFeatures_Cfm_Domains_Domain_Services_Service
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetFilter() yfilter.YFilter { return services.YFilter }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) SetFilter(yf yfilter.YFilter) { services.YFilter = yf }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetGoName(yname string) string {
    if yname == "service" { return "Service" }
    return ""
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetSegmentPath() string {
    return "services"
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service" {
        for _, c := range services.Service {
            if services.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetFeatures_Cfm_Domains_Domain_Services_Service{}
        services.Service = append(services.Service, child)
        return &services.Service[len(services.Service)-1]
    }
    return nil
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range services.Service {
        children[services.Service[i].GetSegmentPath()] = &services.Service[i]
    }
    return children
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetBundleName() string { return "cisco_ios_xr" }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetYangName() string { return "services" }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) SetParent(parent types.Entity) { services.parent = parent }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetParent() types.Entity { return services.parent }

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetParentYangName() string { return "domain" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service
// Configuration for a particular Service
// (Maintenance Association)
type EthernetFeatures_Cfm_Domains_Domain_Services_Service struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Service (Maintenance Association). The type is
    // string with length: 1..79.
    Service interface{}

    // Limit on the number of MEPs in the service. The type is interface{} with
    // range: 2..8190. The default value is 100.
    MaximumMeps interface{}

    // Log Cross-check Errors detected for peer MEPs. The type is interface{}.
    LogCrossCheckErrors interface{}

    // How long to store information for peer MEPs that have timed out. The type
    // is interface{} with range: 1..65535. The default value is 100.
    ContinuityCheckArchiveHoldTime interface{}

    // The number of tags to use when sending CFM packets from up MEPs in this
    // Service. The type is interface{} with range: 0..4294967295.
    Tags interface{}

    // Log peer MEPs state changes. The type is interface{}.
    LogContinuityCheckStateChanges interface{}

    // Enable logging. The type is interface{}.
    LogEfd interface{}

    // Automatically trigger a traceroute when a peer MEP times out. The type is
    // interface{}.
    ContinuityCheckAutoTraceroute interface{}

    // Log CCM Errors detected for peer MEPs. The type is interface{}.
    LogContinuityCheckErrors interface{}

    // Log receipt of AIS and LCK messages. The type is interface{}.
    LogAis interface{}

    // Enable EFD to bring down ports when MEPs detect errors.
    Efd2 EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2

    // Continuity Check Interval and Loss Threshold.  Configuring the interval
    // enables Continuity Check.
    ContinuityCheckInterval EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval

    // MIP Auto-creation Policy.
    MipAutoCreation EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation

    // Service specific AIS configuration.
    Ais EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais

    // Cross-check configuration.
    CrossCheck EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck

    // Fundamental properties of the service (maintenance association).
    ServiceProperties EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetFilter() yfilter.YFilter { return service.YFilter }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) SetFilter(yf yfilter.YFilter) { service.YFilter = yf }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetGoName(yname string) string {
    if yname == "service" { return "Service" }
    if yname == "maximum-meps" { return "MaximumMeps" }
    if yname == "log-cross-check-errors" { return "LogCrossCheckErrors" }
    if yname == "continuity-check-archive-hold-time" { return "ContinuityCheckArchiveHoldTime" }
    if yname == "tags" { return "Tags" }
    if yname == "log-continuity-check-state-changes" { return "LogContinuityCheckStateChanges" }
    if yname == "log-efd" { return "LogEfd" }
    if yname == "continuity-check-auto-traceroute" { return "ContinuityCheckAutoTraceroute" }
    if yname == "log-continuity-check-errors" { return "LogContinuityCheckErrors" }
    if yname == "log-ais" { return "LogAis" }
    if yname == "efd2" { return "Efd2" }
    if yname == "continuity-check-interval" { return "ContinuityCheckInterval" }
    if yname == "mip-auto-creation" { return "MipAutoCreation" }
    if yname == "ais" { return "Ais" }
    if yname == "cross-check" { return "CrossCheck" }
    if yname == "service-properties" { return "ServiceProperties" }
    return ""
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetSegmentPath() string {
    return "service" + "[service='" + fmt.Sprintf("%v", service.Service) + "']"
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "efd2" {
        return &service.Efd2
    }
    if childYangName == "continuity-check-interval" {
        return &service.ContinuityCheckInterval
    }
    if childYangName == "mip-auto-creation" {
        return &service.MipAutoCreation
    }
    if childYangName == "ais" {
        return &service.Ais
    }
    if childYangName == "cross-check" {
        return &service.CrossCheck
    }
    if childYangName == "service-properties" {
        return &service.ServiceProperties
    }
    return nil
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["efd2"] = &service.Efd2
    children["continuity-check-interval"] = &service.ContinuityCheckInterval
    children["mip-auto-creation"] = &service.MipAutoCreation
    children["ais"] = &service.Ais
    children["cross-check"] = &service.CrossCheck
    children["service-properties"] = &service.ServiceProperties
    return children
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service"] = service.Service
    leafs["maximum-meps"] = service.MaximumMeps
    leafs["log-cross-check-errors"] = service.LogCrossCheckErrors
    leafs["continuity-check-archive-hold-time"] = service.ContinuityCheckArchiveHoldTime
    leafs["tags"] = service.Tags
    leafs["log-continuity-check-state-changes"] = service.LogContinuityCheckStateChanges
    leafs["log-efd"] = service.LogEfd
    leafs["continuity-check-auto-traceroute"] = service.ContinuityCheckAutoTraceroute
    leafs["log-continuity-check-errors"] = service.LogContinuityCheckErrors
    leafs["log-ais"] = service.LogAis
    return leafs
}

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetBundleName() string { return "cisco_ios_xr" }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetYangName() string { return "service" }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) SetParent(parent types.Entity) { service.parent = parent }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetParent() types.Entity { return service.parent }

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetParentYangName() string { return "services" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2
// Enable EFD to bring down ports when MEPs
// detect errors
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable EFD. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Enable protection switching notifications. The type is interface{}.
    ProtectionSwitchingEnable interface{}
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetFilter() yfilter.YFilter { return efd2.YFilter }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) SetFilter(yf yfilter.YFilter) { efd2.YFilter = yf }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "protection-switching-enable" { return "ProtectionSwitchingEnable" }
    return ""
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetSegmentPath() string {
    return "efd2"
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = efd2.Enable
    leafs["protection-switching-enable"] = efd2.ProtectionSwitchingEnable
    return leafs
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetBundleName() string { return "cisco_ios_xr" }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetYangName() string { return "efd2" }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) SetParent(parent types.Entity) { efd2.parent = parent }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetParent() types.Entity { return efd2.parent }

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval
// Continuity Check Interval and Loss
// Threshold.  Configuring the interval
// enables Continuity Check.
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // CCM Interval. The type is CfmCcmInterval. This attribute is mandatory.
    CcmInterval interface{}

    // Loss Threshold (default 3). The type is interface{} with range: 2..255.
    LossThreshold interface{}
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetFilter() yfilter.YFilter { return continuityCheckInterval.YFilter }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) SetFilter(yf yfilter.YFilter) { continuityCheckInterval.YFilter = yf }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetGoName(yname string) string {
    if yname == "ccm-interval" { return "CcmInterval" }
    if yname == "loss-threshold" { return "LossThreshold" }
    return ""
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetSegmentPath() string {
    return "continuity-check-interval"
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccm-interval"] = continuityCheckInterval.CcmInterval
    leafs["loss-threshold"] = continuityCheckInterval.LossThreshold
    return leafs
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetBundleName() string { return "cisco_ios_xr" }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetYangName() string { return "continuity-check-interval" }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) SetParent(parent types.Entity) { continuityCheckInterval.parent = parent }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetParent() types.Entity { return continuityCheckInterval.parent }

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation
// MIP Auto-creation Policy
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MIP Auto-creation Policy. The type is CfmMipPolicy. This attribute is
    // mandatory.
    MipPolicy interface{}

    // Enable CCM Learning at MIPs in this service. The type is interface{}.
    CcmLearningEnable interface{}
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetFilter() yfilter.YFilter { return mipAutoCreation.YFilter }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) SetFilter(yf yfilter.YFilter) { mipAutoCreation.YFilter = yf }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetGoName(yname string) string {
    if yname == "mip-policy" { return "MipPolicy" }
    if yname == "ccm-learning-enable" { return "CcmLearningEnable" }
    return ""
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetSegmentPath() string {
    return "mip-auto-creation"
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mip-policy"] = mipAutoCreation.MipPolicy
    leafs["ccm-learning-enable"] = mipAutoCreation.CcmLearningEnable
    return leafs
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetBundleName() string { return "cisco_ios_xr" }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetYangName() string { return "mip-auto-creation" }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) SetParent(parent types.Entity) { mipAutoCreation.parent = parent }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetParent() types.Entity { return mipAutoCreation.parent }

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais
// Service specific AIS configuration
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS transmission configuration.
    Transmission EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetFilter() yfilter.YFilter { return ais.YFilter }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) SetFilter(yf yfilter.YFilter) { ais.YFilter = yf }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetGoName(yname string) string {
    if yname == "transmission" { return "Transmission" }
    return ""
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetSegmentPath() string {
    return "ais"
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transmission" {
        return &ais.Transmission
    }
    return nil
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transmission"] = &ais.Transmission
    return children
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetBundleName() string { return "cisco_ios_xr" }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetYangName() string { return "ais" }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) SetParent(parent types.Entity) { ais.parent = parent }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetParent() types.Entity { return ais.parent }

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission
// AIS transmission configuration
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AIS Interval. The type is CfmAisInterval.
    AisInterval interface{}

    // Class of Service bits. The type is interface{} with range: 0..7.
    Cos interface{}
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetFilter() yfilter.YFilter { return transmission.YFilter }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) SetFilter(yf yfilter.YFilter) { transmission.YFilter = yf }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetGoName(yname string) string {
    if yname == "ais-interval" { return "AisInterval" }
    if yname == "cos" { return "Cos" }
    return ""
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetSegmentPath() string {
    return "transmission"
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ais-interval"] = transmission.AisInterval
    leafs["cos"] = transmission.Cos
    return leafs
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetBundleName() string { return "cisco_ios_xr" }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetYangName() string { return "transmission" }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) SetParent(parent types.Entity) { transmission.parent = parent }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetParent() types.Entity { return transmission.parent }

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetParentYangName() string { return "ais" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck
// Cross-check configuration
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable automatic MEP cross-check. The type is interface{}.
    Auto interface{}

    // Cross-check MEPs.
    CrossCheckMeps EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetFilter() yfilter.YFilter { return crossCheck.YFilter }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) SetFilter(yf yfilter.YFilter) { crossCheck.YFilter = yf }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetGoName(yname string) string {
    if yname == "auto" { return "Auto" }
    if yname == "cross-check-meps" { return "CrossCheckMeps" }
    return ""
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetSegmentPath() string {
    return "cross-check"
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cross-check-meps" {
        return &crossCheck.CrossCheckMeps
    }
    return nil
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cross-check-meps"] = &crossCheck.CrossCheckMeps
    return children
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["auto"] = crossCheck.Auto
    return leafs
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetBundleName() string { return "cisco_ios_xr" }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetYangName() string { return "cross-check" }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) SetParent(parent types.Entity) { crossCheck.parent = parent }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetParent() types.Entity { return crossCheck.parent }

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps
// Cross-check MEPs
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MEP ID and optional MAC Address for Cross-check. The type is slice of
    // EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep.
    CrossCheckMep []EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetFilter() yfilter.YFilter { return crossCheckMeps.YFilter }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) SetFilter(yf yfilter.YFilter) { crossCheckMeps.YFilter = yf }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetGoName(yname string) string {
    if yname == "cross-check-mep" { return "CrossCheckMep" }
    return ""
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetSegmentPath() string {
    return "cross-check-meps"
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cross-check-mep" {
        for _, c := range crossCheckMeps.CrossCheckMep {
            if crossCheckMeps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep{}
        crossCheckMeps.CrossCheckMep = append(crossCheckMeps.CrossCheckMep, child)
        return &crossCheckMeps.CrossCheckMep[len(crossCheckMeps.CrossCheckMep)-1]
    }
    return nil
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range crossCheckMeps.CrossCheckMep {
        children[crossCheckMeps.CrossCheckMep[i].GetSegmentPath()] = &crossCheckMeps.CrossCheckMep[i]
    }
    return children
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetBundleName() string { return "cisco_ios_xr" }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetYangName() string { return "cross-check-meps" }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) SetParent(parent types.Entity) { crossCheckMeps.parent = parent }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetParent() types.Entity { return crossCheckMeps.parent }

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetParentYangName() string { return "cross-check" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep
// MEP ID and optional MAC Address for
// Cross-check
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // MAC Address is specified. The type is interface{}.
    EnableMacAddress interface{}

    // MAC Address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetFilter() yfilter.YFilter { return crossCheckMep.YFilter }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) SetFilter(yf yfilter.YFilter) { crossCheckMep.YFilter = yf }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetGoName(yname string) string {
    if yname == "mep-id" { return "MepId" }
    if yname == "enable-mac-address" { return "EnableMacAddress" }
    if yname == "mac-address" { return "MacAddress" }
    return ""
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetSegmentPath() string {
    return "cross-check-mep" + "[mep-id='" + fmt.Sprintf("%v", crossCheckMep.MepId) + "']"
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mep-id"] = crossCheckMep.MepId
    leafs["enable-mac-address"] = crossCheckMep.EnableMacAddress
    leafs["mac-address"] = crossCheckMep.MacAddress
    return leafs
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetBundleName() string { return "cisco_ios_xr" }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetYangName() string { return "cross-check-mep" }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) SetParent(parent types.Entity) { crossCheckMep.parent = parent }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetParent() types.Entity { return crossCheckMep.parent }

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetParentYangName() string { return "cross-check-meps" }

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties
// Fundamental properties of the service
// (maintenance association)
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of Service. The type is CfmService. This attribute is mandatory.
    ServiceType interface{}

    // Bridge Group or Cross-connect Group, if Service Type is BridgeDomain or
    // CrossConnect. The type is string.
    GroupName interface{}

    // Bridge Domain or Cross-connect name, if Service Type is BridgeDomain or
    // CrossConnect. The type is string.
    SwitchingName interface{}

    // Local Customer Edge Identifier. The type is interface{} with range:
    // 1..16384.
    CeId interface{}

    // Remote Customer Edge Identifier. The type is interface{} with range:
    // 1..16384.
    RemoteCeId interface{}

    // EVPN ID. The type is interface{} with range: 1..65534.
    Evi interface{}

    // Short MA Name Format. The type is CfmShortMaNameFormat.
    ShortMaNameFormat interface{}

    // String Short MA Name, if format is String. The type is string with length:
    // 1..45.
    ShortMaNameString interface{}

    // Numeric Short MA Name, if format is VlanID or Number. The type is
    // interface{} with range: 0..65535.
    ShortMaNameNumber interface{}

    // VPN OUI, if Short MA Name format is VPN_ID. The type is interface{} with
    // range: 0..16777215.
    ShortMaNameOui interface{}

    // VPN Index, if Short MA Name format is VPN_ID. The type is interface{} with
    // range: -2147483648..2147483647.
    ShortMaNameVpnIndex interface{}

    // ITU Carrier Code (ICC), if format is ICCBased. The type is string with
    // length: 1..6.
    ShortMaNameIcc interface{}

    // Unique MEG ID Code (UMC), if format is ICCBased. The type is string with
    // length: 1..12.
    ShortMaNameUmc interface{}
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetFilter() yfilter.YFilter { return serviceProperties.YFilter }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) SetFilter(yf yfilter.YFilter) { serviceProperties.YFilter = yf }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetGoName(yname string) string {
    if yname == "service-type" { return "ServiceType" }
    if yname == "group-name" { return "GroupName" }
    if yname == "switching-name" { return "SwitchingName" }
    if yname == "ce-id" { return "CeId" }
    if yname == "remote-ce-id" { return "RemoteCeId" }
    if yname == "evi" { return "Evi" }
    if yname == "short-ma-name-format" { return "ShortMaNameFormat" }
    if yname == "short-ma-name-string" { return "ShortMaNameString" }
    if yname == "short-ma-name-number" { return "ShortMaNameNumber" }
    if yname == "short-ma-name-oui" { return "ShortMaNameOui" }
    if yname == "short-ma-name-vpn-index" { return "ShortMaNameVpnIndex" }
    if yname == "short-ma-name-icc" { return "ShortMaNameIcc" }
    if yname == "short-ma-name-umc" { return "ShortMaNameUmc" }
    return ""
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetSegmentPath() string {
    return "service-properties"
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["service-type"] = serviceProperties.ServiceType
    leafs["group-name"] = serviceProperties.GroupName
    leafs["switching-name"] = serviceProperties.SwitchingName
    leafs["ce-id"] = serviceProperties.CeId
    leafs["remote-ce-id"] = serviceProperties.RemoteCeId
    leafs["evi"] = serviceProperties.Evi
    leafs["short-ma-name-format"] = serviceProperties.ShortMaNameFormat
    leafs["short-ma-name-string"] = serviceProperties.ShortMaNameString
    leafs["short-ma-name-number"] = serviceProperties.ShortMaNameNumber
    leafs["short-ma-name-oui"] = serviceProperties.ShortMaNameOui
    leafs["short-ma-name-vpn-index"] = serviceProperties.ShortMaNameVpnIndex
    leafs["short-ma-name-icc"] = serviceProperties.ShortMaNameIcc
    leafs["short-ma-name-umc"] = serviceProperties.ShortMaNameUmc
    return leafs
}

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetBundleName() string { return "cisco_ios_xr" }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetYangName() string { return "service-properties" }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) SetParent(parent types.Entity) { serviceProperties.parent = parent }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetParent() types.Entity { return serviceProperties.parent }

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetParentYangName() string { return "service" }

// EthernetFeatures_Cfm_Domains_Domain_DomainProperties
// Fundamental properties of the domain
type EthernetFeatures_Cfm_Domains_Domain_DomainProperties struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maintenance Domain Level. The type is interface{} with range: 0..7.
    Level interface{}

    // Maintenance Domain ID Format. The type is CfmMdidFormat.
    MdidFormat interface{}

    // MAC Address, if MDID Format is MACAddress. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MdidMacAddress interface{}

    // Unsigned 16-bit Interger, if MDID Format is MACAddress. The type is
    // interface{} with range: 0..65535.
    MdidNumber interface{}

    // String MDID, if MDID format is String or DNSLike. The type is string with
    // length: 1..43.
    MdidString interface{}
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetFilter() yfilter.YFilter { return domainProperties.YFilter }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) SetFilter(yf yfilter.YFilter) { domainProperties.YFilter = yf }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetGoName(yname string) string {
    if yname == "level" { return "Level" }
    if yname == "mdid-format" { return "MdidFormat" }
    if yname == "mdid-mac-address" { return "MdidMacAddress" }
    if yname == "mdid-number" { return "MdidNumber" }
    if yname == "mdid-string" { return "MdidString" }
    return ""
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetSegmentPath() string {
    return "domain-properties"
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level"] = domainProperties.Level
    leafs["mdid-format"] = domainProperties.MdidFormat
    leafs["mdid-mac-address"] = domainProperties.MdidMacAddress
    leafs["mdid-number"] = domainProperties.MdidNumber
    leafs["mdid-string"] = domainProperties.MdidString
    return leafs
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetBundleName() string { return "cisco_ios_xr" }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetYangName() string { return "domain-properties" }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) SetParent(parent types.Entity) { domainProperties.parent = parent }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetParent() types.Entity { return domainProperties.parent }

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetParentYangName() string { return "domain" }

// EthernetFeatures_EtherLinkOam
// Ethernet Link OAM Global Configuration
type EthernetFeatures_EtherLinkOam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Ethernet Link OAM profiles.
    Profiles EthernetFeatures_EtherLinkOam_Profiles
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetFilter() yfilter.YFilter { return etherLinkOam.YFilter }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) SetFilter(yf yfilter.YFilter) { etherLinkOam.YFilter = yf }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetGoName(yname string) string {
    if yname == "profiles" { return "Profiles" }
    return ""
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam"
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profiles" {
        return &etherLinkOam.Profiles
    }
    return nil
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profiles"] = &etherLinkOam.Profiles
    return children
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetBundleName() string { return "cisco_ios_xr" }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetYangName() string { return "ether-link-oam" }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) SetParent(parent types.Entity) { etherLinkOam.parent = parent }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetParent() types.Entity { return etherLinkOam.parent }

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetParentYangName() string { return "ethernet-features" }

// EthernetFeatures_EtherLinkOam_Profiles
// Table of Ethernet Link OAM profiles
type EthernetFeatures_EtherLinkOam_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the profile. The type is slice of
    // EthernetFeatures_EtherLinkOam_Profiles_Profile.
    Profile []EthernetFeatures_EtherLinkOam_Profiles_Profile
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EthernetFeatures_EtherLinkOam_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetYangName() string { return "profiles" }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetParentYangName() string { return "ether-link-oam" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile
// Name of the profile
type EthernetFeatures_EtherLinkOam_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Profile interface{}

    // Enable or disable MIB retrieval support. The type is bool.
    MibRetrieval interface{}

    // Enable or disable uni-directional link-fault detection support. The type is
    // bool.
    Udlf interface{}

    // Possible Ethernet Link OAM hello intervals. The type is
    // EtherLinkOamHelloIntervalEnum.
    HelloInterval interface{}

    // Set the OAM mode to passive. The type is EtherLinkOamModeEnum.
    Mode interface{}

    // Enable or disable remote loopback support. The type is bool.
    RemoteLoopback interface{}

    // Connection timeout period in number of lost heartbeats. The type is
    // interface{} with range: 2..30.
    Timeout interface{}

    // Configure action parameters.
    Action EthernetFeatures_EtherLinkOam_Profiles_Profile_Action

    // Configure remote requirement parameters.
    RequireRemote EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote

    // Configure link monitor parameters.
    LinkMonitoring EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    if yname == "mib-retrieval" { return "MibRetrieval" }
    if yname == "udlf" { return "Udlf" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "mode" { return "Mode" }
    if yname == "remote-loopback" { return "RemoteLoopback" }
    if yname == "timeout" { return "Timeout" }
    if yname == "action" { return "Action" }
    if yname == "require-remote" { return "RequireRemote" }
    if yname == "link-monitoring" { return "LinkMonitoring" }
    return ""
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile='" + fmt.Sprintf("%v", profile.Profile) + "']"
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "action" {
        return &profile.Action
    }
    if childYangName == "require-remote" {
        return &profile.RequireRemote
    }
    if childYangName == "link-monitoring" {
        return &profile.LinkMonitoring
    }
    return nil
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["action"] = &profile.Action
    children["require-remote"] = &profile.RequireRemote
    children["link-monitoring"] = &profile.LinkMonitoring
    return children
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = profile.Profile
    leafs["mib-retrieval"] = profile.MibRetrieval
    leafs["udlf"] = profile.Udlf
    leafs["hello-interval"] = profile.HelloInterval
    leafs["mode"] = profile.Mode
    leafs["remote-loopback"] = profile.RemoteLoopback
    leafs["timeout"] = profile.Timeout
    return leafs
}

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetParentYangName() string { return "profiles" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_Action
// Configure action parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_Action struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action to perform when a dying gasp occurs. The type is
    // EtherLinkOamEventActionEnum.
    DyingGasp interface{}

    // Action to perform when a session comes up. The type is
    // EtherLinkOamEventActionPrimEnum.
    SessionUp interface{}

    // Action to perform when a critical event occurs. The type is
    // EtherLinkOamEventActionEnum.
    CriticalEvent interface{}

    // Action to perform when a session goes down. The type is
    // EtherLinkOamEventActionEnumEfd.
    SessionDown interface{}

    // Action to perform when discovery timeout occurs. The type is
    // EtherLinkOamEventActionEnumEfd.
    DiscoveryTimeout interface{}

    // Action to perform when a high-threshold event occurs. The type is
    // EtherLinkOamEventActionEnum.
    HighThreshold interface{}

    // Action to perform when a capabilities conflict occurs. The type is
    // EtherLinkOamEventActionEnumEfd.
    CapabilitiesConflict interface{}

    // Action to perform when remote loopback is entered or exited. The type is
    // EtherLinkOamEventActionPrimEnum.
    RemoteLoopback interface{}

    // Action to perform when a link fault message is received. The type is
    // EtherLinkOamEventActionEnumEfd.
    LinkFault interface{}

    // Action to perform when a wiring conflict occurs. The type is
    // EtherLinkOamEventActionEnumEfd.
    WiringConflict interface{}
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetFilter() yfilter.YFilter { return action.YFilter }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) SetFilter(yf yfilter.YFilter) { action.YFilter = yf }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetGoName(yname string) string {
    if yname == "dying-gasp" { return "DyingGasp" }
    if yname == "session-up" { return "SessionUp" }
    if yname == "critical-event" { return "CriticalEvent" }
    if yname == "session-down" { return "SessionDown" }
    if yname == "discovery-timeout" { return "DiscoveryTimeout" }
    if yname == "high-threshold" { return "HighThreshold" }
    if yname == "capabilities-conflict" { return "CapabilitiesConflict" }
    if yname == "remote-loopback" { return "RemoteLoopback" }
    if yname == "link-fault" { return "LinkFault" }
    if yname == "wiring-conflict" { return "WiringConflict" }
    return ""
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetSegmentPath() string {
    return "action"
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dying-gasp"] = action.DyingGasp
    leafs["session-up"] = action.SessionUp
    leafs["critical-event"] = action.CriticalEvent
    leafs["session-down"] = action.SessionDown
    leafs["discovery-timeout"] = action.DiscoveryTimeout
    leafs["high-threshold"] = action.HighThreshold
    leafs["capabilities-conflict"] = action.CapabilitiesConflict
    leafs["remote-loopback"] = action.RemoteLoopback
    leafs["link-fault"] = action.LinkFault
    leafs["wiring-conflict"] = action.WiringConflict
    return leafs
}

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetBundleName() string { return "cisco_ios_xr" }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetYangName() string { return "action" }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) SetParent(parent types.Entity) { action.parent = parent }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetParent() types.Entity { return action.parent }

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetParentYangName() string { return "profile" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote
// Configure remote requirement parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable MIB retrieval requirement. The type is bool.
    MibRetrieval interface{}

    // Possible required OAM modes. The type is EtherLinkOamRequireModeEnum.
    Mode interface{}

    // Enable or disable remote loopback requirement. The type is bool.
    RemoteLoopback interface{}

    // Enable or disable link monitoring requirement. The type is bool.
    LinkMonitoring interface{}
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetFilter() yfilter.YFilter { return requireRemote.YFilter }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) SetFilter(yf yfilter.YFilter) { requireRemote.YFilter = yf }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetGoName(yname string) string {
    if yname == "mib-retrieval" { return "MibRetrieval" }
    if yname == "mode" { return "Mode" }
    if yname == "remote-loopback" { return "RemoteLoopback" }
    if yname == "link-monitoring" { return "LinkMonitoring" }
    return ""
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetSegmentPath() string {
    return "require-remote"
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mib-retrieval"] = requireRemote.MibRetrieval
    leafs["mode"] = requireRemote.Mode
    leafs["remote-loopback"] = requireRemote.RemoteLoopback
    leafs["link-monitoring"] = requireRemote.LinkMonitoring
    return leafs
}

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetBundleName() string { return "cisco_ios_xr" }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetYangName() string { return "require-remote" }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) SetParent(parent types.Entity) { requireRemote.parent = parent }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetParent() types.Entity { return requireRemote.parent }

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetParentYangName() string { return "profile" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring
// Configure link monitor parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable monitoring. The type is bool.
    Monitoring interface{}

    // Symbol-period event configuration.
    SymbolPeriod EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod

    // Frame-period event configuration.
    FramePeriod EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod

    // Frame-seconds event configuration.
    FrameSeconds EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds

    // Frame event configuration.
    Frame EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetFilter() yfilter.YFilter { return linkMonitoring.YFilter }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) SetFilter(yf yfilter.YFilter) { linkMonitoring.YFilter = yf }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetGoName(yname string) string {
    if yname == "monitoring" { return "Monitoring" }
    if yname == "symbol-period" { return "SymbolPeriod" }
    if yname == "frame-period" { return "FramePeriod" }
    if yname == "frame-seconds" { return "FrameSeconds" }
    if yname == "frame" { return "Frame" }
    return ""
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetSegmentPath() string {
    return "link-monitoring"
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "symbol-period" {
        return &linkMonitoring.SymbolPeriod
    }
    if childYangName == "frame-period" {
        return &linkMonitoring.FramePeriod
    }
    if childYangName == "frame-seconds" {
        return &linkMonitoring.FrameSeconds
    }
    if childYangName == "frame" {
        return &linkMonitoring.Frame
    }
    return nil
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["symbol-period"] = &linkMonitoring.SymbolPeriod
    children["frame-period"] = &linkMonitoring.FramePeriod
    children["frame-seconds"] = &linkMonitoring.FrameSeconds
    children["frame"] = &linkMonitoring.Frame
    return children
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["monitoring"] = linkMonitoring.Monitoring
    return leafs
}

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetBundleName() string { return "cisco_ios_xr" }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetYangName() string { return "link-monitoring" }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) SetParent(parent types.Entity) { linkMonitoring.parent = parent }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetParent() types.Entity { return linkMonitoring.parent }

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetParentYangName() string { return "profile" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod
// Symbol-period event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Window size configuration for symbol-period events.
    Window EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window

    // Threshold configuration for symbol-period events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetFilter() yfilter.YFilter { return symbolPeriod.YFilter }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) SetFilter(yf yfilter.YFilter) { symbolPeriod.YFilter = yf }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetSegmentPath() string {
    return "symbol-period"
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "window" {
        return &symbolPeriod.Window
    }
    if childYangName == "threshold" {
        return &symbolPeriod.Threshold
    }
    return nil
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["window"] = &symbolPeriod.Window
    children["threshold"] = &symbolPeriod.Threshold
    return children
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetBundleName() string { return "cisco_ios_xr" }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetYangName() string { return "symbol-period" }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) SetParent(parent types.Entity) { symbolPeriod.parent = parent }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetParent() types.Entity { return symbolPeriod.parent }

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetParentYangName() string { return "link-monitoring" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window
// Window size configuration for symbol-period
// events
// This type is a presence type.
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of the symbol-period window. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    Window interface{}

    // Units to use for this window. The type is
    // EtherLinkOamWindowUnitsSymbolsEnum. This attribute is mandatory.
    Units interface{}

    // The multiplier to use for this window (only valid if 'Units' is Symbols and
    // treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum. The default value is none.
    Multiplier interface{}
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetFilter() yfilter.YFilter { return window.YFilter }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) SetFilter(yf yfilter.YFilter) { window.YFilter = yf }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "units" { return "Units" }
    if yname == "multiplier" { return "Multiplier" }
    return ""
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetSegmentPath() string {
    return "window"
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["window"] = window.Window
    leafs["units"] = window.Units
    leafs["multiplier"] = window.Multiplier
    return leafs
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetBundleName() string { return "cisco_ios_xr" }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetYangName() string { return "window" }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) SetParent(parent types.Entity) { window.parent = parent }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetParent() types.Entity { return window.parent }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetParentYangName() string { return "symbol-period" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold
// Threshold configuration for symbol-period
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The low threshold for symbol-period. The type is interface{} with range:
    // 1..4294967295. The default value is 1.
    ThresholdLow interface{}

    // The high threshold for symbol-period. The type is interface{} with range:
    // 1..4294967295.
    ThresholdHigh interface{}

    // The units to use for these thresholds. The type is
    // EtherLinkOamThresholdUnitsSymbolsEnum. The default value is symbols.
    Units interface{}

    // The multiplier to use for the low threshold (only valid if 'Units' is
    // Symbols and treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum. The default value is none.
    MultiplierLow interface{}

    // The multiplier to use for the high threshold (only valid if 'Units' is
    // Symbols and treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum.
    MultiplierHigh interface{}
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetGoName(yname string) string {
    if yname == "threshold-low" { return "ThresholdLow" }
    if yname == "threshold-high" { return "ThresholdHigh" }
    if yname == "units" { return "Units" }
    if yname == "multiplier-low" { return "MultiplierLow" }
    if yname == "multiplier-high" { return "MultiplierHigh" }
    return ""
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-low"] = threshold.ThresholdLow
    leafs["threshold-high"] = threshold.ThresholdHigh
    leafs["units"] = threshold.Units
    leafs["multiplier-low"] = threshold.MultiplierLow
    leafs["multiplier-high"] = threshold.MultiplierHigh
    return leafs
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetYangName() string { return "threshold" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetParentYangName() string { return "symbol-period" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod
// Frame-period event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Window size configuration for frame-period events.
    Window EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window

    // Threshold configuration for frame-period events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetFilter() yfilter.YFilter { return framePeriod.YFilter }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) SetFilter(yf yfilter.YFilter) { framePeriod.YFilter = yf }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetSegmentPath() string {
    return "frame-period"
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "window" {
        return &framePeriod.Window
    }
    if childYangName == "threshold" {
        return &framePeriod.Threshold
    }
    return nil
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["window"] = &framePeriod.Window
    children["threshold"] = &framePeriod.Threshold
    return children
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetBundleName() string { return "cisco_ios_xr" }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetYangName() string { return "frame-period" }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) SetParent(parent types.Entity) { framePeriod.parent = parent }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetParent() types.Entity { return framePeriod.parent }

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetParentYangName() string { return "link-monitoring" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window
// Window size configuration for frame-period
// events
// This type is a presence type.
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Size of the frame-period window. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    Window interface{}

    // The units to use for this window. The type is
    // EtherLinkOamWindowUnitsFramesEnum. This attribute is mandatory.
    Units interface{}

    // The multiplier to use for this window (only valid if 'Units' is Frames and
    // treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum. The default value is none.
    Multiplier interface{}
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetFilter() yfilter.YFilter { return window.YFilter }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) SetFilter(yf yfilter.YFilter) { window.YFilter = yf }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "units" { return "Units" }
    if yname == "multiplier" { return "Multiplier" }
    return ""
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetSegmentPath() string {
    return "window"
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["window"] = window.Window
    leafs["units"] = window.Units
    leafs["multiplier"] = window.Multiplier
    return leafs
}

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetBundleName() string { return "cisco_ios_xr" }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetYangName() string { return "window" }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) SetParent(parent types.Entity) { window.parent = parent }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetParent() types.Entity { return window.parent }

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetParentYangName() string { return "frame-period" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold
// Threshold configuration for frame-period
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The low threshold for frame-period events. The type is interface{} with
    // range: 1..4294967295. The default value is 1.
    ThresholdLow interface{}

    // The high threshold for frame-period events. The type is interface{} with
    // range: 1..4294967295.
    ThresholdHigh interface{}

    // The units to use for these thresholds. The type is
    // EtherLinkOamThresholdUnitsFramesEnum. The default value is ppm.
    Units interface{}

    // The multiplier to use for the low threshold (only valid if 'Units' is
    // Frames and treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum. The default value is none.
    MultiplierLow interface{}

    // The multiplier to use for the high threshold (only valid if 'Units' is
    // Frames and treated as 1 if unspecified). The type is
    // EtherLinkOamThresholdWindowMultiplierEnum.
    MultiplierHigh interface{}
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetGoName(yname string) string {
    if yname == "threshold-low" { return "ThresholdLow" }
    if yname == "threshold-high" { return "ThresholdHigh" }
    if yname == "units" { return "Units" }
    if yname == "multiplier-low" { return "MultiplierLow" }
    if yname == "multiplier-high" { return "MultiplierHigh" }
    return ""
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-low"] = threshold.ThresholdLow
    leafs["threshold-high"] = threshold.ThresholdHigh
    leafs["units"] = threshold.Units
    leafs["multiplier-low"] = threshold.MultiplierLow
    leafs["multiplier-high"] = threshold.MultiplierHigh
    return leafs
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetYangName() string { return "threshold" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetParentYangName() string { return "frame-period" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds
// Frame-seconds event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Window size configuration for frame-seconds events. The type is interface{}
    // with range: 10000..900000. Units are millisecond. The default value is
    // 60000.
    Window interface{}

    // Threshold configuration for frame-seconds events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetFilter() yfilter.YFilter { return frameSeconds.YFilter }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) SetFilter(yf yfilter.YFilter) { frameSeconds.YFilter = yf }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetSegmentPath() string {
    return "frame-seconds"
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &frameSeconds.Threshold
    }
    return nil
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &frameSeconds.Threshold
    return children
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["window"] = frameSeconds.Window
    return leafs
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetBundleName() string { return "cisco_ios_xr" }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetYangName() string { return "frame-seconds" }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) SetParent(parent types.Entity) { frameSeconds.parent = parent }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetParent() types.Entity { return frameSeconds.parent }

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetParentYangName() string { return "link-monitoring" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold
// Threshold configuration for frame-seconds
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The low threshold for frame-seconds events. The type is interface{} with
    // range: 1..900. Units are second. The default value is 1.
    ThresholdLow interface{}

    // The high threshold for frame-seconds events. The type is interface{} with
    // range: 1..900. Units are second.
    ThresholdHigh interface{}
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetGoName(yname string) string {
    if yname == "threshold-low" { return "ThresholdLow" }
    if yname == "threshold-high" { return "ThresholdHigh" }
    return ""
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-low"] = threshold.ThresholdLow
    leafs["threshold-high"] = threshold.ThresholdHigh
    return leafs
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetYangName() string { return "threshold" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetParentYangName() string { return "frame-seconds" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame
// Frame event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Window size configuration for frame events. The type is interface{} with
    // range: 1000..60000. Units are millisecond. The default value is 1000.
    Window interface{}

    // Threshold configuration for frame events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetFilter() yfilter.YFilter { return frame.YFilter }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) SetFilter(yf yfilter.YFilter) { frame.YFilter = yf }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetGoName(yname string) string {
    if yname == "window" { return "Window" }
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetSegmentPath() string {
    return "frame"
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &frame.Threshold
    }
    return nil
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &frame.Threshold
    return children
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["window"] = frame.Window
    return leafs
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetBundleName() string { return "cisco_ios_xr" }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetYangName() string { return "frame" }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) SetParent(parent types.Entity) { frame.parent = parent }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetParent() types.Entity { return frame.parent }

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetParentYangName() string { return "link-monitoring" }

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold
// Threshold configuration for frame events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The low threshold for frame events. The type is interface{} with range:
    // 1..4294967295. The default value is 1.
    ThresholdLow interface{}

    // The high threshold for frame events. The type is interface{} with range:
    // 1..4294967295.
    ThresholdHigh interface{}

    // The multiplier to use for the low threshold (treated as 1 if unspecified).
    // The type is EtherLinkOamThresholdWindowMultiplierEnum. The default value is
    // none.
    MultiplierLow interface{}

    // The multiplier to use for the high threshold (treated as 1 if unspecified).
    // The type is EtherLinkOamThresholdWindowMultiplierEnum.
    MultiplierHigh interface{}
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetGoName(yname string) string {
    if yname == "threshold-low" { return "ThresholdLow" }
    if yname == "threshold-high" { return "ThresholdHigh" }
    if yname == "multiplier-low" { return "MultiplierLow" }
    if yname == "multiplier-high" { return "MultiplierHigh" }
    return ""
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold-low"] = threshold.ThresholdLow
    leafs["threshold-high"] = threshold.ThresholdHigh
    leafs["multiplier-low"] = threshold.MultiplierLow
    leafs["multiplier-high"] = threshold.MultiplierHigh
    return leafs
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetYangName() string { return "threshold" }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetParentYangName() string { return "frame" }

