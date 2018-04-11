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

// EthernetFeatures
// Ethernet Features Configuration
type EthernetFeatures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Egress Filtering Configuration.
    EgressFiltering EthernetFeatures_EgressFiltering

    // CFM global configuration.
    Cfm EthernetFeatures_Cfm

    // Ethernet Link OAM Global Configuration.
    EtherLinkOam EthernetFeatures_EtherLinkOam
}

func (ethernetFeatures *EthernetFeatures) GetEntityData() *types.CommonEntityData {
    ethernetFeatures.EntityData.YFilter = ethernetFeatures.YFilter
    ethernetFeatures.EntityData.YangName = "ethernet-features"
    ethernetFeatures.EntityData.BundleName = "cisco_ios_xr"
    ethernetFeatures.EntityData.ParentYangName = "Cisco-IOS-XR-l2-eth-infra-cfg"
    ethernetFeatures.EntityData.SegmentPath = "Cisco-IOS-XR-l2-eth-infra-cfg:ethernet-features"
    ethernetFeatures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetFeatures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetFeatures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetFeatures.EntityData.Children = make(map[string]types.YChild)
    ethernetFeatures.EntityData.Children["egress-filtering"] = types.YChild{"EgressFiltering", &ethernetFeatures.EgressFiltering}
    ethernetFeatures.EntityData.Children["Cisco-IOS-XR-ethernet-cfm-cfg:cfm"] = types.YChild{"Cfm", &ethernetFeatures.Cfm}
    ethernetFeatures.EntityData.Children["Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam"] = types.YChild{"EtherLinkOam", &ethernetFeatures.EtherLinkOam}
    ethernetFeatures.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetFeatures.EntityData)
}

// EthernetFeatures_EgressFiltering
// Egress Filtering Configuration
type EthernetFeatures_EgressFiltering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether Egress Filtering is on by default. The type is interface{}.
    EgressFilteringDefaultOn interface{}
}

func (egressFiltering *EthernetFeatures_EgressFiltering) GetEntityData() *types.CommonEntityData {
    egressFiltering.EntityData.YFilter = egressFiltering.YFilter
    egressFiltering.EntityData.YangName = "egress-filtering"
    egressFiltering.EntityData.BundleName = "cisco_ios_xr"
    egressFiltering.EntityData.ParentYangName = "ethernet-features"
    egressFiltering.EntityData.SegmentPath = "egress-filtering"
    egressFiltering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressFiltering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressFiltering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressFiltering.EntityData.Children = make(map[string]types.YChild)
    egressFiltering.EntityData.Leafs = make(map[string]types.YLeaf)
    egressFiltering.EntityData.Leafs["egress-filtering-default-on"] = types.YLeaf{"EgressFilteringDefaultOn", egressFiltering.EgressFilteringDefaultOn}
    return &(egressFiltering.EntityData)
}

// EthernetFeatures_Cfm
// CFM global configuration
type EthernetFeatures_Cfm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable processing of Ethernet SLA packets on nV Satellite devices. The
    // type is interface{}.
    NvSatelliteSlaProcessingDisable interface{}

    // Traceroute Cache Configuration.
    TracerouteCache EthernetFeatures_Cfm_TracerouteCache

    // Domain-specific global configuration.
    Domains EthernetFeatures_Cfm_Domains
}

func (cfm *EthernetFeatures_Cfm) GetEntityData() *types.CommonEntityData {
    cfm.EntityData.YFilter = cfm.YFilter
    cfm.EntityData.YangName = "cfm"
    cfm.EntityData.BundleName = "cisco_ios_xr"
    cfm.EntityData.ParentYangName = "ethernet-features"
    cfm.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-cfm-cfg:cfm"
    cfm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfm.EntityData.Children = make(map[string]types.YChild)
    cfm.EntityData.Children["traceroute-cache"] = types.YChild{"TracerouteCache", &cfm.TracerouteCache}
    cfm.EntityData.Children["domains"] = types.YChild{"Domains", &cfm.Domains}
    cfm.EntityData.Leafs = make(map[string]types.YLeaf)
    cfm.EntityData.Leafs["nv-satellite-sla-processing-disable"] = types.YLeaf{"NvSatelliteSlaProcessingDisable", cfm.NvSatelliteSlaProcessingDisable}
    return &(cfm.EntityData)
}

// EthernetFeatures_Cfm_TracerouteCache
// Traceroute Cache Configuration
type EthernetFeatures_Cfm_TracerouteCache struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hold Time in minutes. The type is interface{} with range: 1..525600. The
    // default value is 100.
    HoldTime interface{}

    // Cache Size limit (number of replies). The type is interface{} with range:
    // 1..4294967295. The default value is 100.
    CacheSize interface{}
}

func (tracerouteCache *EthernetFeatures_Cfm_TracerouteCache) GetEntityData() *types.CommonEntityData {
    tracerouteCache.EntityData.YFilter = tracerouteCache.YFilter
    tracerouteCache.EntityData.YangName = "traceroute-cache"
    tracerouteCache.EntityData.BundleName = "cisco_ios_xr"
    tracerouteCache.EntityData.ParentYangName = "cfm"
    tracerouteCache.EntityData.SegmentPath = "traceroute-cache"
    tracerouteCache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracerouteCache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracerouteCache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracerouteCache.EntityData.Children = make(map[string]types.YChild)
    tracerouteCache.EntityData.Leafs = make(map[string]types.YLeaf)
    tracerouteCache.EntityData.Leafs["hold-time"] = types.YLeaf{"HoldTime", tracerouteCache.HoldTime}
    tracerouteCache.EntityData.Leafs["cache-size"] = types.YLeaf{"CacheSize", tracerouteCache.CacheSize}
    return &(tracerouteCache.EntityData)
}

// EthernetFeatures_Cfm_Domains
// Domain-specific global configuration
type EthernetFeatures_Cfm_Domains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular Maintenance Domain. The type is slice of
    // EthernetFeatures_Cfm_Domains_Domain.
    Domain []EthernetFeatures_Cfm_Domains_Domain
}

func (domains *EthernetFeatures_Cfm_Domains) GetEntityData() *types.CommonEntityData {
    domains.EntityData.YFilter = domains.YFilter
    domains.EntityData.YangName = "domains"
    domains.EntityData.BundleName = "cisco_ios_xr"
    domains.EntityData.ParentYangName = "cfm"
    domains.EntityData.SegmentPath = "domains"
    domains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domains.EntityData.Children = make(map[string]types.YChild)
    domains.EntityData.Children["domain"] = types.YChild{"Domain", nil}
    for i := range domains.Domain {
        domains.EntityData.Children[types.GetSegmentPath(&domains.Domain[i])] = types.YChild{"Domain", &domains.Domain[i]}
    }
    domains.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(domains.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain
// Configuration for a particular Maintenance
// Domain
type EthernetFeatures_Cfm_Domains_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Maintenance Domain. The type is string with
    // length: 1..79.
    Domain interface{}

    // Service-specific global configuration.
    Services EthernetFeatures_Cfm_Domains_Domain_Services

    // Fundamental properties of the domain.
    DomainProperties EthernetFeatures_Cfm_Domains_Domain_DomainProperties
}

func (domain *EthernetFeatures_Cfm_Domains_Domain) GetEntityData() *types.CommonEntityData {
    domain.EntityData.YFilter = domain.YFilter
    domain.EntityData.YangName = "domain"
    domain.EntityData.BundleName = "cisco_ios_xr"
    domain.EntityData.ParentYangName = "domains"
    domain.EntityData.SegmentPath = "domain" + "[domain='" + fmt.Sprintf("%v", domain.Domain) + "']"
    domain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domain.EntityData.Children = make(map[string]types.YChild)
    domain.EntityData.Children["services"] = types.YChild{"Services", &domain.Services}
    domain.EntityData.Children["domain-properties"] = types.YChild{"DomainProperties", &domain.DomainProperties}
    domain.EntityData.Leafs = make(map[string]types.YLeaf)
    domain.EntityData.Leafs["domain"] = types.YLeaf{"Domain", domain.Domain}
    return &(domain.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services
// Service-specific global configuration
type EthernetFeatures_Cfm_Domains_Domain_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular Service (Maintenance Association). The type
    // is slice of EthernetFeatures_Cfm_Domains_Domain_Services_Service.
    Service []EthernetFeatures_Cfm_Domains_Domain_Services_Service
}

func (services *EthernetFeatures_Cfm_Domains_Domain_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "domain"
    services.EntityData.SegmentPath = "services"
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = make(map[string]types.YChild)
    services.EntityData.Children["service"] = types.YChild{"Service", nil}
    for i := range services.Service {
        services.EntityData.Children[types.GetSegmentPath(&services.Service[i])] = types.YChild{"Service", &services.Service[i]}
    }
    services.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(services.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service
// Configuration for a particular Service
// (Maintenance Association)
type EthernetFeatures_Cfm_Domains_Domain_Services_Service struct {
    EntityData types.CommonEntityData
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

func (service *EthernetFeatures_Cfm_Domains_Domain_Services_Service) GetEntityData() *types.CommonEntityData {
    service.EntityData.YFilter = service.YFilter
    service.EntityData.YangName = "service"
    service.EntityData.BundleName = "cisco_ios_xr"
    service.EntityData.ParentYangName = "services"
    service.EntityData.SegmentPath = "service" + "[service='" + fmt.Sprintf("%v", service.Service) + "']"
    service.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    service.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    service.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    service.EntityData.Children = make(map[string]types.YChild)
    service.EntityData.Children["efd2"] = types.YChild{"Efd2", &service.Efd2}
    service.EntityData.Children["continuity-check-interval"] = types.YChild{"ContinuityCheckInterval", &service.ContinuityCheckInterval}
    service.EntityData.Children["mip-auto-creation"] = types.YChild{"MipAutoCreation", &service.MipAutoCreation}
    service.EntityData.Children["ais"] = types.YChild{"Ais", &service.Ais}
    service.EntityData.Children["cross-check"] = types.YChild{"CrossCheck", &service.CrossCheck}
    service.EntityData.Children["service-properties"] = types.YChild{"ServiceProperties", &service.ServiceProperties}
    service.EntityData.Leafs = make(map[string]types.YLeaf)
    service.EntityData.Leafs["service"] = types.YLeaf{"Service", service.Service}
    service.EntityData.Leafs["maximum-meps"] = types.YLeaf{"MaximumMeps", service.MaximumMeps}
    service.EntityData.Leafs["log-cross-check-errors"] = types.YLeaf{"LogCrossCheckErrors", service.LogCrossCheckErrors}
    service.EntityData.Leafs["continuity-check-archive-hold-time"] = types.YLeaf{"ContinuityCheckArchiveHoldTime", service.ContinuityCheckArchiveHoldTime}
    service.EntityData.Leafs["tags"] = types.YLeaf{"Tags", service.Tags}
    service.EntityData.Leafs["log-continuity-check-state-changes"] = types.YLeaf{"LogContinuityCheckStateChanges", service.LogContinuityCheckStateChanges}
    service.EntityData.Leafs["log-efd"] = types.YLeaf{"LogEfd", service.LogEfd}
    service.EntityData.Leafs["continuity-check-auto-traceroute"] = types.YLeaf{"ContinuityCheckAutoTraceroute", service.ContinuityCheckAutoTraceroute}
    service.EntityData.Leafs["log-continuity-check-errors"] = types.YLeaf{"LogContinuityCheckErrors", service.LogContinuityCheckErrors}
    service.EntityData.Leafs["log-ais"] = types.YLeaf{"LogAis", service.LogAis}
    return &(service.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2
// Enable EFD to bring down ports when MEPs
// detect errors
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable EFD. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Enable protection switching notifications. The type is interface{}.
    ProtectionSwitchingEnable interface{}
}

func (efd2 *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Efd2) GetEntityData() *types.CommonEntityData {
    efd2.EntityData.YFilter = efd2.YFilter
    efd2.EntityData.YangName = "efd2"
    efd2.EntityData.BundleName = "cisco_ios_xr"
    efd2.EntityData.ParentYangName = "service"
    efd2.EntityData.SegmentPath = "efd2"
    efd2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    efd2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    efd2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    efd2.EntityData.Children = make(map[string]types.YChild)
    efd2.EntityData.Leafs = make(map[string]types.YLeaf)
    efd2.EntityData.Leafs["enable"] = types.YLeaf{"Enable", efd2.Enable}
    efd2.EntityData.Leafs["protection-switching-enable"] = types.YLeaf{"ProtectionSwitchingEnable", efd2.ProtectionSwitchingEnable}
    return &(efd2.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval
// Continuity Check Interval and Loss
// Threshold.  Configuring the interval
// enables Continuity Check.
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CCM Interval. The type is CfmCcmInterval. This attribute is mandatory.
    CcmInterval interface{}

    // Loss Threshold (default 3). The type is interface{} with range: 2..255.
    LossThreshold interface{}
}

func (continuityCheckInterval *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ContinuityCheckInterval) GetEntityData() *types.CommonEntityData {
    continuityCheckInterval.EntityData.YFilter = continuityCheckInterval.YFilter
    continuityCheckInterval.EntityData.YangName = "continuity-check-interval"
    continuityCheckInterval.EntityData.BundleName = "cisco_ios_xr"
    continuityCheckInterval.EntityData.ParentYangName = "service"
    continuityCheckInterval.EntityData.SegmentPath = "continuity-check-interval"
    continuityCheckInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    continuityCheckInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    continuityCheckInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    continuityCheckInterval.EntityData.Children = make(map[string]types.YChild)
    continuityCheckInterval.EntityData.Leafs = make(map[string]types.YLeaf)
    continuityCheckInterval.EntityData.Leafs["ccm-interval"] = types.YLeaf{"CcmInterval", continuityCheckInterval.CcmInterval}
    continuityCheckInterval.EntityData.Leafs["loss-threshold"] = types.YLeaf{"LossThreshold", continuityCheckInterval.LossThreshold}
    return &(continuityCheckInterval.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation
// MIP Auto-creation Policy
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MIP Auto-creation Policy. The type is CfmMipPolicy. This attribute is
    // mandatory.
    MipPolicy interface{}

    // Enable CCM Learning at MIPs in this service. The type is interface{}.
    CcmLearningEnable interface{}
}

func (mipAutoCreation *EthernetFeatures_Cfm_Domains_Domain_Services_Service_MipAutoCreation) GetEntityData() *types.CommonEntityData {
    mipAutoCreation.EntityData.YFilter = mipAutoCreation.YFilter
    mipAutoCreation.EntityData.YangName = "mip-auto-creation"
    mipAutoCreation.EntityData.BundleName = "cisco_ios_xr"
    mipAutoCreation.EntityData.ParentYangName = "service"
    mipAutoCreation.EntityData.SegmentPath = "mip-auto-creation"
    mipAutoCreation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mipAutoCreation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mipAutoCreation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mipAutoCreation.EntityData.Children = make(map[string]types.YChild)
    mipAutoCreation.EntityData.Leafs = make(map[string]types.YLeaf)
    mipAutoCreation.EntityData.Leafs["mip-policy"] = types.YLeaf{"MipPolicy", mipAutoCreation.MipPolicy}
    mipAutoCreation.EntityData.Leafs["ccm-learning-enable"] = types.YLeaf{"CcmLearningEnable", mipAutoCreation.CcmLearningEnable}
    return &(mipAutoCreation.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais
// Service specific AIS configuration
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS transmission configuration.
    Transmission EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission
}

func (ais *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais) GetEntityData() *types.CommonEntityData {
    ais.EntityData.YFilter = ais.YFilter
    ais.EntityData.YangName = "ais"
    ais.EntityData.BundleName = "cisco_ios_xr"
    ais.EntityData.ParentYangName = "service"
    ais.EntityData.SegmentPath = "ais"
    ais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ais.EntityData.Children = make(map[string]types.YChild)
    ais.EntityData.Children["transmission"] = types.YChild{"Transmission", &ais.Transmission}
    ais.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ais.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission
// AIS transmission configuration
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AIS Interval. The type is CfmAisInterval.
    AisInterval interface{}

    // Class of Service bits. The type is interface{} with range: 0..7.
    Cos interface{}
}

func (transmission *EthernetFeatures_Cfm_Domains_Domain_Services_Service_Ais_Transmission) GetEntityData() *types.CommonEntityData {
    transmission.EntityData.YFilter = transmission.YFilter
    transmission.EntityData.YangName = "transmission"
    transmission.EntityData.BundleName = "cisco_ios_xr"
    transmission.EntityData.ParentYangName = "ais"
    transmission.EntityData.SegmentPath = "transmission"
    transmission.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmission.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmission.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmission.EntityData.Children = make(map[string]types.YChild)
    transmission.EntityData.Leafs = make(map[string]types.YLeaf)
    transmission.EntityData.Leafs["ais-interval"] = types.YLeaf{"AisInterval", transmission.AisInterval}
    transmission.EntityData.Leafs["cos"] = types.YLeaf{"Cos", transmission.Cos}
    return &(transmission.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck
// Cross-check configuration
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable automatic MEP cross-check. The type is interface{}.
    Auto interface{}

    // Cross-check MEPs.
    CrossCheckMeps EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps
}

func (crossCheck *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck) GetEntityData() *types.CommonEntityData {
    crossCheck.EntityData.YFilter = crossCheck.YFilter
    crossCheck.EntityData.YangName = "cross-check"
    crossCheck.EntityData.BundleName = "cisco_ios_xr"
    crossCheck.EntityData.ParentYangName = "service"
    crossCheck.EntityData.SegmentPath = "cross-check"
    crossCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossCheck.EntityData.Children = make(map[string]types.YChild)
    crossCheck.EntityData.Children["cross-check-meps"] = types.YChild{"CrossCheckMeps", &crossCheck.CrossCheckMeps}
    crossCheck.EntityData.Leafs = make(map[string]types.YLeaf)
    crossCheck.EntityData.Leafs["auto"] = types.YLeaf{"Auto", crossCheck.Auto}
    return &(crossCheck.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps
// Cross-check MEPs
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MEP ID and optional MAC Address for Cross-check. The type is slice of
    // EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep.
    CrossCheckMep []EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep
}

func (crossCheckMeps *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps) GetEntityData() *types.CommonEntityData {
    crossCheckMeps.EntityData.YFilter = crossCheckMeps.YFilter
    crossCheckMeps.EntityData.YangName = "cross-check-meps"
    crossCheckMeps.EntityData.BundleName = "cisco_ios_xr"
    crossCheckMeps.EntityData.ParentYangName = "cross-check"
    crossCheckMeps.EntityData.SegmentPath = "cross-check-meps"
    crossCheckMeps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossCheckMeps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossCheckMeps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossCheckMeps.EntityData.Children = make(map[string]types.YChild)
    crossCheckMeps.EntityData.Children["cross-check-mep"] = types.YChild{"CrossCheckMep", nil}
    for i := range crossCheckMeps.CrossCheckMep {
        crossCheckMeps.EntityData.Children[types.GetSegmentPath(&crossCheckMeps.CrossCheckMep[i])] = types.YChild{"CrossCheckMep", &crossCheckMeps.CrossCheckMep[i]}
    }
    crossCheckMeps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(crossCheckMeps.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep
// MEP ID and optional MAC Address for
// Cross-check
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MEP ID. The type is interface{} with range:
    // 1..8191.
    MepId interface{}

    // MAC Address is specified. The type is interface{}.
    EnableMacAddress interface{}

    // MAC Address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}
}

func (crossCheckMep *EthernetFeatures_Cfm_Domains_Domain_Services_Service_CrossCheck_CrossCheckMeps_CrossCheckMep) GetEntityData() *types.CommonEntityData {
    crossCheckMep.EntityData.YFilter = crossCheckMep.YFilter
    crossCheckMep.EntityData.YangName = "cross-check-mep"
    crossCheckMep.EntityData.BundleName = "cisco_ios_xr"
    crossCheckMep.EntityData.ParentYangName = "cross-check-meps"
    crossCheckMep.EntityData.SegmentPath = "cross-check-mep" + "[mep-id='" + fmt.Sprintf("%v", crossCheckMep.MepId) + "']"
    crossCheckMep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    crossCheckMep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    crossCheckMep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    crossCheckMep.EntityData.Children = make(map[string]types.YChild)
    crossCheckMep.EntityData.Leafs = make(map[string]types.YLeaf)
    crossCheckMep.EntityData.Leafs["mep-id"] = types.YLeaf{"MepId", crossCheckMep.MepId}
    crossCheckMep.EntityData.Leafs["enable-mac-address"] = types.YLeaf{"EnableMacAddress", crossCheckMep.EnableMacAddress}
    crossCheckMep.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", crossCheckMep.MacAddress}
    return &(crossCheckMep.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties
// Fundamental properties of the service
// (maintenance association)
// This type is a presence type.
type EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties struct {
    EntityData types.CommonEntityData
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

func (serviceProperties *EthernetFeatures_Cfm_Domains_Domain_Services_Service_ServiceProperties) GetEntityData() *types.CommonEntityData {
    serviceProperties.EntityData.YFilter = serviceProperties.YFilter
    serviceProperties.EntityData.YangName = "service-properties"
    serviceProperties.EntityData.BundleName = "cisco_ios_xr"
    serviceProperties.EntityData.ParentYangName = "service"
    serviceProperties.EntityData.SegmentPath = "service-properties"
    serviceProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceProperties.EntityData.Children = make(map[string]types.YChild)
    serviceProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceProperties.EntityData.Leafs["service-type"] = types.YLeaf{"ServiceType", serviceProperties.ServiceType}
    serviceProperties.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", serviceProperties.GroupName}
    serviceProperties.EntityData.Leafs["switching-name"] = types.YLeaf{"SwitchingName", serviceProperties.SwitchingName}
    serviceProperties.EntityData.Leafs["ce-id"] = types.YLeaf{"CeId", serviceProperties.CeId}
    serviceProperties.EntityData.Leafs["remote-ce-id"] = types.YLeaf{"RemoteCeId", serviceProperties.RemoteCeId}
    serviceProperties.EntityData.Leafs["evi"] = types.YLeaf{"Evi", serviceProperties.Evi}
    serviceProperties.EntityData.Leafs["short-ma-name-format"] = types.YLeaf{"ShortMaNameFormat", serviceProperties.ShortMaNameFormat}
    serviceProperties.EntityData.Leafs["short-ma-name-string"] = types.YLeaf{"ShortMaNameString", serviceProperties.ShortMaNameString}
    serviceProperties.EntityData.Leafs["short-ma-name-number"] = types.YLeaf{"ShortMaNameNumber", serviceProperties.ShortMaNameNumber}
    serviceProperties.EntityData.Leafs["short-ma-name-oui"] = types.YLeaf{"ShortMaNameOui", serviceProperties.ShortMaNameOui}
    serviceProperties.EntityData.Leafs["short-ma-name-vpn-index"] = types.YLeaf{"ShortMaNameVpnIndex", serviceProperties.ShortMaNameVpnIndex}
    serviceProperties.EntityData.Leafs["short-ma-name-icc"] = types.YLeaf{"ShortMaNameIcc", serviceProperties.ShortMaNameIcc}
    serviceProperties.EntityData.Leafs["short-ma-name-umc"] = types.YLeaf{"ShortMaNameUmc", serviceProperties.ShortMaNameUmc}
    return &(serviceProperties.EntityData)
}

// EthernetFeatures_Cfm_Domains_Domain_DomainProperties
// Fundamental properties of the domain
type EthernetFeatures_Cfm_Domains_Domain_DomainProperties struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maintenance Domain Level. The type is interface{} with range: 0..7.
    Level interface{}

    // Maintenance Domain ID Format. The type is CfmMdidFormat.
    MdidFormat interface{}

    // MAC Address, if MDID Format is MACAddress. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MdidMacAddress interface{}

    // Unsigned 16-bit Interger, if MDID Format is MACAddress. The type is
    // interface{} with range: 0..65535.
    MdidNumber interface{}

    // String MDID, if MDID format is String or DNSLike. The type is string with
    // length: 1..43.
    MdidString interface{}
}

func (domainProperties *EthernetFeatures_Cfm_Domains_Domain_DomainProperties) GetEntityData() *types.CommonEntityData {
    domainProperties.EntityData.YFilter = domainProperties.YFilter
    domainProperties.EntityData.YangName = "domain-properties"
    domainProperties.EntityData.BundleName = "cisco_ios_xr"
    domainProperties.EntityData.ParentYangName = "domain"
    domainProperties.EntityData.SegmentPath = "domain-properties"
    domainProperties.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainProperties.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainProperties.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainProperties.EntityData.Children = make(map[string]types.YChild)
    domainProperties.EntityData.Leafs = make(map[string]types.YLeaf)
    domainProperties.EntityData.Leafs["level"] = types.YLeaf{"Level", domainProperties.Level}
    domainProperties.EntityData.Leafs["mdid-format"] = types.YLeaf{"MdidFormat", domainProperties.MdidFormat}
    domainProperties.EntityData.Leafs["mdid-mac-address"] = types.YLeaf{"MdidMacAddress", domainProperties.MdidMacAddress}
    domainProperties.EntityData.Leafs["mdid-number"] = types.YLeaf{"MdidNumber", domainProperties.MdidNumber}
    domainProperties.EntityData.Leafs["mdid-string"] = types.YLeaf{"MdidString", domainProperties.MdidString}
    return &(domainProperties.EntityData)
}

// EthernetFeatures_EtherLinkOam
// Ethernet Link OAM Global Configuration
type EthernetFeatures_EtherLinkOam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Ethernet Link OAM profiles.
    Profiles EthernetFeatures_EtherLinkOam_Profiles
}

func (etherLinkOam *EthernetFeatures_EtherLinkOam) GetEntityData() *types.CommonEntityData {
    etherLinkOam.EntityData.YFilter = etherLinkOam.YFilter
    etherLinkOam.EntityData.YangName = "ether-link-oam"
    etherLinkOam.EntityData.BundleName = "cisco_ios_xr"
    etherLinkOam.EntityData.ParentYangName = "ethernet-features"
    etherLinkOam.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-link-oam-cfg:ether-link-oam"
    etherLinkOam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etherLinkOam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etherLinkOam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etherLinkOam.EntityData.Children = make(map[string]types.YChild)
    etherLinkOam.EntityData.Children["profiles"] = types.YChild{"Profiles", &etherLinkOam.Profiles}
    etherLinkOam.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(etherLinkOam.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles
// Table of Ethernet Link OAM profiles
type EthernetFeatures_EtherLinkOam_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the profile. The type is slice of
    // EthernetFeatures_EtherLinkOam_Profiles_Profile.
    Profile []EthernetFeatures_EtherLinkOam_Profiles_Profile
}

func (profiles *EthernetFeatures_EtherLinkOam_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ether-link-oam"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = make(map[string]types.YChild)
    profiles.EntityData.Children["profile"] = types.YChild{"Profile", nil}
    for i := range profiles.Profile {
        profiles.EntityData.Children[types.GetSegmentPath(&profiles.Profile[i])] = types.YChild{"Profile", &profiles.Profile[i]}
    }
    profiles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(profiles.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile
// Name of the profile
type EthernetFeatures_EtherLinkOam_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (profile *EthernetFeatures_EtherLinkOam_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + "[profile='" + fmt.Sprintf("%v", profile.Profile) + "']"
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = make(map[string]types.YChild)
    profile.EntityData.Children["action"] = types.YChild{"Action", &profile.Action}
    profile.EntityData.Children["require-remote"] = types.YChild{"RequireRemote", &profile.RequireRemote}
    profile.EntityData.Children["link-monitoring"] = types.YChild{"LinkMonitoring", &profile.LinkMonitoring}
    profile.EntityData.Leafs = make(map[string]types.YLeaf)
    profile.EntityData.Leafs["profile"] = types.YLeaf{"Profile", profile.Profile}
    profile.EntityData.Leafs["mib-retrieval"] = types.YLeaf{"MibRetrieval", profile.MibRetrieval}
    profile.EntityData.Leafs["udlf"] = types.YLeaf{"Udlf", profile.Udlf}
    profile.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", profile.HelloInterval}
    profile.EntityData.Leafs["mode"] = types.YLeaf{"Mode", profile.Mode}
    profile.EntityData.Leafs["remote-loopback"] = types.YLeaf{"RemoteLoopback", profile.RemoteLoopback}
    profile.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", profile.Timeout}
    return &(profile.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_Action
// Configure action parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_Action struct {
    EntityData types.CommonEntityData
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

func (action *EthernetFeatures_EtherLinkOam_Profiles_Profile_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "profile"
    action.EntityData.SegmentPath = "action"
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = make(map[string]types.YChild)
    action.EntityData.Leafs = make(map[string]types.YLeaf)
    action.EntityData.Leafs["dying-gasp"] = types.YLeaf{"DyingGasp", action.DyingGasp}
    action.EntityData.Leafs["session-up"] = types.YLeaf{"SessionUp", action.SessionUp}
    action.EntityData.Leafs["critical-event"] = types.YLeaf{"CriticalEvent", action.CriticalEvent}
    action.EntityData.Leafs["session-down"] = types.YLeaf{"SessionDown", action.SessionDown}
    action.EntityData.Leafs["discovery-timeout"] = types.YLeaf{"DiscoveryTimeout", action.DiscoveryTimeout}
    action.EntityData.Leafs["high-threshold"] = types.YLeaf{"HighThreshold", action.HighThreshold}
    action.EntityData.Leafs["capabilities-conflict"] = types.YLeaf{"CapabilitiesConflict", action.CapabilitiesConflict}
    action.EntityData.Leafs["remote-loopback"] = types.YLeaf{"RemoteLoopback", action.RemoteLoopback}
    action.EntityData.Leafs["link-fault"] = types.YLeaf{"LinkFault", action.LinkFault}
    action.EntityData.Leafs["wiring-conflict"] = types.YLeaf{"WiringConflict", action.WiringConflict}
    return &(action.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote
// Configure remote requirement parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote struct {
    EntityData types.CommonEntityData
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

func (requireRemote *EthernetFeatures_EtherLinkOam_Profiles_Profile_RequireRemote) GetEntityData() *types.CommonEntityData {
    requireRemote.EntityData.YFilter = requireRemote.YFilter
    requireRemote.EntityData.YangName = "require-remote"
    requireRemote.EntityData.BundleName = "cisco_ios_xr"
    requireRemote.EntityData.ParentYangName = "profile"
    requireRemote.EntityData.SegmentPath = "require-remote"
    requireRemote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requireRemote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requireRemote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requireRemote.EntityData.Children = make(map[string]types.YChild)
    requireRemote.EntityData.Leafs = make(map[string]types.YLeaf)
    requireRemote.EntityData.Leafs["mib-retrieval"] = types.YLeaf{"MibRetrieval", requireRemote.MibRetrieval}
    requireRemote.EntityData.Leafs["mode"] = types.YLeaf{"Mode", requireRemote.Mode}
    requireRemote.EntityData.Leafs["remote-loopback"] = types.YLeaf{"RemoteLoopback", requireRemote.RemoteLoopback}
    requireRemote.EntityData.Leafs["link-monitoring"] = types.YLeaf{"LinkMonitoring", requireRemote.LinkMonitoring}
    return &(requireRemote.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring
// Configure link monitor parameters
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring struct {
    EntityData types.CommonEntityData
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

func (linkMonitoring *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring) GetEntityData() *types.CommonEntityData {
    linkMonitoring.EntityData.YFilter = linkMonitoring.YFilter
    linkMonitoring.EntityData.YangName = "link-monitoring"
    linkMonitoring.EntityData.BundleName = "cisco_ios_xr"
    linkMonitoring.EntityData.ParentYangName = "profile"
    linkMonitoring.EntityData.SegmentPath = "link-monitoring"
    linkMonitoring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkMonitoring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkMonitoring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkMonitoring.EntityData.Children = make(map[string]types.YChild)
    linkMonitoring.EntityData.Children["symbol-period"] = types.YChild{"SymbolPeriod", &linkMonitoring.SymbolPeriod}
    linkMonitoring.EntityData.Children["frame-period"] = types.YChild{"FramePeriod", &linkMonitoring.FramePeriod}
    linkMonitoring.EntityData.Children["frame-seconds"] = types.YChild{"FrameSeconds", &linkMonitoring.FrameSeconds}
    linkMonitoring.EntityData.Children["frame"] = types.YChild{"Frame", &linkMonitoring.Frame}
    linkMonitoring.EntityData.Leafs = make(map[string]types.YLeaf)
    linkMonitoring.EntityData.Leafs["monitoring"] = types.YLeaf{"Monitoring", linkMonitoring.Monitoring}
    return &(linkMonitoring.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod
// Symbol-period event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Window size configuration for symbol-period events.
    Window EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window

    // Threshold configuration for symbol-period events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold
}

func (symbolPeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod) GetEntityData() *types.CommonEntityData {
    symbolPeriod.EntityData.YFilter = symbolPeriod.YFilter
    symbolPeriod.EntityData.YangName = "symbol-period"
    symbolPeriod.EntityData.BundleName = "cisco_ios_xr"
    symbolPeriod.EntityData.ParentYangName = "link-monitoring"
    symbolPeriod.EntityData.SegmentPath = "symbol-period"
    symbolPeriod.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    symbolPeriod.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    symbolPeriod.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    symbolPeriod.EntityData.Children = make(map[string]types.YChild)
    symbolPeriod.EntityData.Children["window"] = types.YChild{"Window", &symbolPeriod.Window}
    symbolPeriod.EntityData.Children["threshold"] = types.YChild{"Threshold", &symbolPeriod.Threshold}
    symbolPeriod.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(symbolPeriod.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window
// Window size configuration for symbol-period
// events
// This type is a presence type.
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window struct {
    EntityData types.CommonEntityData
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

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Window) GetEntityData() *types.CommonEntityData {
    window.EntityData.YFilter = window.YFilter
    window.EntityData.YangName = "window"
    window.EntityData.BundleName = "cisco_ios_xr"
    window.EntityData.ParentYangName = "symbol-period"
    window.EntityData.SegmentPath = "window"
    window.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    window.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    window.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    window.EntityData.Children = make(map[string]types.YChild)
    window.EntityData.Leafs = make(map[string]types.YLeaf)
    window.EntityData.Leafs["window"] = types.YLeaf{"Window", window.Window}
    window.EntityData.Leafs["units"] = types.YLeaf{"Units", window.Units}
    window.EntityData.Leafs["multiplier"] = types.YLeaf{"Multiplier", window.Multiplier}
    return &(window.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold
// Threshold configuration for symbol-period
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_SymbolPeriod_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "symbol-period"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["threshold-low"] = types.YLeaf{"ThresholdLow", threshold.ThresholdLow}
    threshold.EntityData.Leafs["threshold-high"] = types.YLeaf{"ThresholdHigh", threshold.ThresholdHigh}
    threshold.EntityData.Leafs["units"] = types.YLeaf{"Units", threshold.Units}
    threshold.EntityData.Leafs["multiplier-low"] = types.YLeaf{"MultiplierLow", threshold.MultiplierLow}
    threshold.EntityData.Leafs["multiplier-high"] = types.YLeaf{"MultiplierHigh", threshold.MultiplierHigh}
    return &(threshold.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod
// Frame-period event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Window size configuration for frame-period events.
    Window EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window

    // Threshold configuration for frame-period events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold
}

func (framePeriod *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod) GetEntityData() *types.CommonEntityData {
    framePeriod.EntityData.YFilter = framePeriod.YFilter
    framePeriod.EntityData.YangName = "frame-period"
    framePeriod.EntityData.BundleName = "cisco_ios_xr"
    framePeriod.EntityData.ParentYangName = "link-monitoring"
    framePeriod.EntityData.SegmentPath = "frame-period"
    framePeriod.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    framePeriod.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    framePeriod.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    framePeriod.EntityData.Children = make(map[string]types.YChild)
    framePeriod.EntityData.Children["window"] = types.YChild{"Window", &framePeriod.Window}
    framePeriod.EntityData.Children["threshold"] = types.YChild{"Threshold", &framePeriod.Threshold}
    framePeriod.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(framePeriod.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window
// Window size configuration for frame-period
// events
// This type is a presence type.
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window struct {
    EntityData types.CommonEntityData
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

func (window *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Window) GetEntityData() *types.CommonEntityData {
    window.EntityData.YFilter = window.YFilter
    window.EntityData.YangName = "window"
    window.EntityData.BundleName = "cisco_ios_xr"
    window.EntityData.ParentYangName = "frame-period"
    window.EntityData.SegmentPath = "window"
    window.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    window.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    window.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    window.EntityData.Children = make(map[string]types.YChild)
    window.EntityData.Leafs = make(map[string]types.YLeaf)
    window.EntityData.Leafs["window"] = types.YLeaf{"Window", window.Window}
    window.EntityData.Leafs["units"] = types.YLeaf{"Units", window.Units}
    window.EntityData.Leafs["multiplier"] = types.YLeaf{"Multiplier", window.Multiplier}
    return &(window.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold
// Threshold configuration for frame-period
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FramePeriod_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "frame-period"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["threshold-low"] = types.YLeaf{"ThresholdLow", threshold.ThresholdLow}
    threshold.EntityData.Leafs["threshold-high"] = types.YLeaf{"ThresholdHigh", threshold.ThresholdHigh}
    threshold.EntityData.Leafs["units"] = types.YLeaf{"Units", threshold.Units}
    threshold.EntityData.Leafs["multiplier-low"] = types.YLeaf{"MultiplierLow", threshold.MultiplierLow}
    threshold.EntityData.Leafs["multiplier-high"] = types.YLeaf{"MultiplierHigh", threshold.MultiplierHigh}
    return &(threshold.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds
// Frame-seconds event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Window size configuration for frame-seconds events. The type is interface{}
    // with range: 10000..900000. Units are millisecond. The default value is
    // 60000.
    Window interface{}

    // Threshold configuration for frame-seconds events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold
}

func (frameSeconds *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds) GetEntityData() *types.CommonEntityData {
    frameSeconds.EntityData.YFilter = frameSeconds.YFilter
    frameSeconds.EntityData.YangName = "frame-seconds"
    frameSeconds.EntityData.BundleName = "cisco_ios_xr"
    frameSeconds.EntityData.ParentYangName = "link-monitoring"
    frameSeconds.EntityData.SegmentPath = "frame-seconds"
    frameSeconds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frameSeconds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frameSeconds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frameSeconds.EntityData.Children = make(map[string]types.YChild)
    frameSeconds.EntityData.Children["threshold"] = types.YChild{"Threshold", &frameSeconds.Threshold}
    frameSeconds.EntityData.Leafs = make(map[string]types.YLeaf)
    frameSeconds.EntityData.Leafs["window"] = types.YLeaf{"Window", frameSeconds.Window}
    return &(frameSeconds.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold
// Threshold configuration for frame-seconds
// events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The low threshold for frame-seconds events. The type is interface{} with
    // range: 1..900. Units are second. The default value is 1.
    ThresholdLow interface{}

    // The high threshold for frame-seconds events. The type is interface{} with
    // range: 1..900. Units are second.
    ThresholdHigh interface{}
}

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_FrameSeconds_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "frame-seconds"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["threshold-low"] = types.YLeaf{"ThresholdLow", threshold.ThresholdLow}
    threshold.EntityData.Leafs["threshold-high"] = types.YLeaf{"ThresholdHigh", threshold.ThresholdHigh}
    return &(threshold.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame
// Frame event configuration
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Window size configuration for frame events. The type is interface{} with
    // range: 1000..60000. Units are millisecond. The default value is 1000.
    Window interface{}

    // Threshold configuration for frame events.
    Threshold EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold
}

func (frame *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame) GetEntityData() *types.CommonEntityData {
    frame.EntityData.YFilter = frame.YFilter
    frame.EntityData.YangName = "frame"
    frame.EntityData.BundleName = "cisco_ios_xr"
    frame.EntityData.ParentYangName = "link-monitoring"
    frame.EntityData.SegmentPath = "frame"
    frame.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frame.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frame.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frame.EntityData.Children = make(map[string]types.YChild)
    frame.EntityData.Children["threshold"] = types.YChild{"Threshold", &frame.Threshold}
    frame.EntityData.Leafs = make(map[string]types.YLeaf)
    frame.EntityData.Leafs["window"] = types.YLeaf{"Window", frame.Window}
    return &(frame.EntityData)
}

// EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold
// Threshold configuration for frame events
type EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *EthernetFeatures_EtherLinkOam_Profiles_Profile_LinkMonitoring_Frame_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "frame"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["threshold-low"] = types.YLeaf{"ThresholdLow", threshold.ThresholdLow}
    threshold.EntityData.Leafs["threshold-high"] = types.YLeaf{"ThresholdHigh", threshold.ThresholdHigh}
    threshold.EntityData.Leafs["multiplier-low"] = types.YLeaf{"MultiplierLow", threshold.MultiplierLow}
    threshold.EntityData.Leafs["multiplier-high"] = types.YLeaf{"MultiplierHigh", threshold.MultiplierHigh}
    return &(threshold.EntityData)
}

