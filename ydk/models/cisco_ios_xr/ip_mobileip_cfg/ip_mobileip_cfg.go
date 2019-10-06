// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-mobileip package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mobile-ip: MobileIP configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_mobileip_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_mobileip_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-mobileip-cfg mobile-ip}", reflect.TypeOf(MobileIp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip", reflect.TypeOf(MobileIp{}))
}

// ServiceType represents Service type
type ServiceType string

const (
    // ipv4 service type
    ServiceType_ipv4 ServiceType = "ipv4"

    // ipv6 service type
    ServiceType_ipv6 ServiceType = "ipv6"

    // dual service type
    ServiceType_dual ServiceType = "dual"
)

// LmaService represents Lma service
type LmaService string

const (
    // Wireless Private Routing service
    LmaService_service_mll LmaService = "service-mll"
)

// EncapOpt represents Encap opt
type EncapOpt string

const (
    // GRE IPv4 tunnel encap
    EncapOpt_greipv4 EncapOpt = "greipv4"

    // GRE IPv6 tunnel encap
    EncapOpt_greipv6 EncapOpt = "greipv6"

    // mGRE IPv4 tunnel encap
    EncapOpt_mgreipv4 EncapOpt = "mgreipv4"

    // mGRE IPv6 tunnel encap
    EncapOpt_mgreipv6 EncapOpt = "mgreipv6"
)

// RedistSubType represents Redist sub type
type RedistSubType string

const (
    // Redistribute HoA/HNP host prefix routes
    RedistSubType_host_prefix RedistSubType = "host-prefix"

    // Disable redistribution of HoA/HNP host and pool
    // refix routes
    RedistSubType_disable RedistSubType = "disable"
)

// LmaRole represents Lma role
type LmaRole string

const (
    // 3GMA mode
    LmaRole_Y_3gma LmaRole = "3gma"
)

// RedistType represents Redist type
type RedistType string

const (
    // Redistribute HoA/HNP routes
    RedistType_home_address RedistType = "home-address"
)

// GreKeyType represents Gre key type
type GreKeyType string

const (
    // Symmetric GRE Key (same Uplink and Downlink
    // key)
    GreKeyType_symmetric GreKeyType = "symmetric"
)

// LmaRat represents Lma rat
type LmaRat string

const (
    // VIRTUAL rat
    LmaRat_virtual LmaRat = "virtual"

    // PPP rat
    LmaRat_ppp LmaRat = "ppp"

    // ETHERNET rat
    LmaRat_ethernet LmaRat = "ethernet"

    // WLAN rat
    LmaRat_wlan LmaRat = "wlan"

    // WIMAX rat
    LmaRat_wi_max LmaRat = "wi-max"

    // 3GPP_GERAN rat
    LmaRat_Y_3gppgeran LmaRat = "3gppgeran"

    // 3GPP_UTRAN rat
    LmaRat_Y_3gpputran LmaRat = "3gpputran"

    // 3GPP_E_UTRAN rat
    LmaRat_Y_3gppeutran LmaRat = "3gppeutran"

    // 3GPP2_EHRPD rat
    LmaRat_Y_3gpp2ehrpd LmaRat = "3gpp2ehrpd"

    // 3GPP2_HRPD rat
    LmaRat_Y_3gpp2hrpd LmaRat = "3gpp2hrpd"

    // 3GPP2_1RTT rat
    LmaRat_Y_3gpp21rtt LmaRat = "3gpp21rtt"

    // 3GPP2_UMB rat
    LmaRat_Y_3gpp2umb LmaRat = "3gpp2umb"
)

// MobileIp
// MobileIP configuration
type MobileIp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Domain.
    Domains MobileIp_Domains

    // Table of LMA.
    Lmas MobileIp_Lmas
}

func (mobileIp *MobileIp) GetEntityData() *types.CommonEntityData {
    mobileIp.EntityData.YFilter = mobileIp.YFilter
    mobileIp.EntityData.YangName = "mobile-ip"
    mobileIp.EntityData.BundleName = "cisco_ios_xr"
    mobileIp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-mobileip-cfg"
    mobileIp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip"
    mobileIp.EntityData.AbsolutePath = mobileIp.EntityData.SegmentPath
    mobileIp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobileIp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobileIp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobileIp.EntityData.Children = types.NewOrderedMap()
    mobileIp.EntityData.Children.Append("domains", types.YChild{"Domains", &mobileIp.Domains})
    mobileIp.EntityData.Children.Append("lmas", types.YChild{"Lmas", &mobileIp.Lmas})
    mobileIp.EntityData.Leafs = types.NewOrderedMap()

    mobileIp.EntityData.YListKeys = []string {}

    return &(mobileIp.EntityData)
}

// MobileIp_Domains
// Table of Domain
type MobileIp_Domains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PMIPv6 domain configuration. The type is slice of MobileIp_Domains_Domain.
    Domain []*MobileIp_Domains_Domain
}

func (domains *MobileIp_Domains) GetEntityData() *types.CommonEntityData {
    domains.EntityData.YFilter = domains.YFilter
    domains.EntityData.YangName = "domains"
    domains.EntityData.BundleName = "cisco_ios_xr"
    domains.EntityData.ParentYangName = "mobile-ip"
    domains.EntityData.SegmentPath = "domains"
    domains.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/" + domains.EntityData.SegmentPath
    domains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domains.EntityData.Children = types.NewOrderedMap()
    domains.EntityData.Children.Append("domain", types.YChild{"Domain", nil})
    for i := range domains.Domain {
        domains.EntityData.Children.Append(types.GetSegmentPath(domains.Domain[i]), types.YChild{"Domain", domains.Domain[i]})
    }
    domains.EntityData.Leafs = types.NewOrderedMap()

    domains.EntityData.YListKeys = []string {}

    return &(domains.EntityData)
}

// MobileIp_Domains_Domain
// PMIPv6 domain configuration
type MobileIp_Domains_Domain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Domain Name. The type is string with length:
    // 1..125.
    DomainName interface{}

    // Enable PMIPv6 domain configuration. Deletion of this object also causes
    // deletion of all associated objects under Domain. The type is interface{}.
    Enable interface{}

    // Table of MAG.
    Mags MobileIp_Domains_Domain_Mags

    // Table of NAI.
    Nais MobileIp_Domains_Domain_Nais

    // Authentication option between PMIPV6 entities.
    AuthenticateOption MobileIp_Domains_Domain_AuthenticateOption

    // Table of LMA.
    Lmas MobileIp_Domains_Domain_Lmas
}

func (domain *MobileIp_Domains_Domain) GetEntityData() *types.CommonEntityData {
    domain.EntityData.YFilter = domain.YFilter
    domain.EntityData.YangName = "domain"
    domain.EntityData.BundleName = "cisco_ios_xr"
    domain.EntityData.ParentYangName = "domains"
    domain.EntityData.SegmentPath = "domain" + types.AddKeyToken(domain.DomainName, "domain-name")
    domain.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/" + domain.EntityData.SegmentPath
    domain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domain.EntityData.Children = types.NewOrderedMap()
    domain.EntityData.Children.Append("mags", types.YChild{"Mags", &domain.Mags})
    domain.EntityData.Children.Append("nais", types.YChild{"Nais", &domain.Nais})
    domain.EntityData.Children.Append("authenticate-option", types.YChild{"AuthenticateOption", &domain.AuthenticateOption})
    domain.EntityData.Children.Append("lmas", types.YChild{"Lmas", &domain.Lmas})
    domain.EntityData.Leafs = types.NewOrderedMap()
    domain.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", domain.DomainName})
    domain.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", domain.Enable})

    domain.EntityData.YListKeys = []string {"DomainName"}

    return &(domain.EntityData)
}

// MobileIp_Domains_Domain_Mags
// Table of MAG
type MobileIp_Domains_Domain_Mags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAG within domain. The type is slice of MobileIp_Domains_Domain_Mags_Mag.
    Mag []*MobileIp_Domains_Domain_Mags_Mag
}

func (mags *MobileIp_Domains_Domain_Mags) GetEntityData() *types.CommonEntityData {
    mags.EntityData.YFilter = mags.YFilter
    mags.EntityData.YangName = "mags"
    mags.EntityData.BundleName = "cisco_ios_xr"
    mags.EntityData.ParentYangName = "domain"
    mags.EntityData.SegmentPath = "mags"
    mags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/" + mags.EntityData.SegmentPath
    mags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mags.EntityData.Children = types.NewOrderedMap()
    mags.EntityData.Children.Append("mag", types.YChild{"Mag", nil})
    for i := range mags.Mag {
        mags.EntityData.Children.Append(types.GetSegmentPath(mags.Mag[i]), types.YChild{"Mag", mags.Mag[i]})
    }
    mags.EntityData.Leafs = types.NewOrderedMap()

    mags.EntityData.YListKeys = []string {}

    return &(mags.EntityData)
}

// MobileIp_Domains_Domain_Mags_Mag
// MAG within domain
type MobileIp_Domains_Domain_Mags_Mag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MAG Identifier. The type is string with length:
    // 1..125.
    MagName interface{}
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetEntityData() *types.CommonEntityData {
    mag.EntityData.YFilter = mag.YFilter
    mag.EntityData.YangName = "mag"
    mag.EntityData.BundleName = "cisco_ios_xr"
    mag.EntityData.ParentYangName = "mags"
    mag.EntityData.SegmentPath = "mag" + types.AddKeyToken(mag.MagName, "mag-name")
    mag.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/mags/" + mag.EntityData.SegmentPath
    mag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mag.EntityData.Children = types.NewOrderedMap()
    mag.EntityData.Leafs = types.NewOrderedMap()
    mag.EntityData.Leafs.Append("mag-name", types.YLeaf{"MagName", mag.MagName})

    mag.EntityData.YListKeys = []string {"MagName"}

    return &(mag.EntityData)
}

// MobileIp_Domains_Domain_Nais
// Table of NAI
type MobileIp_Domains_Domain_Nais struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Network access identifier or Realm. The type is slice of
    // MobileIp_Domains_Domain_Nais_Nai.
    Nai []*MobileIp_Domains_Domain_Nais_Nai
}

func (nais *MobileIp_Domains_Domain_Nais) GetEntityData() *types.CommonEntityData {
    nais.EntityData.YFilter = nais.YFilter
    nais.EntityData.YangName = "nais"
    nais.EntityData.BundleName = "cisco_ios_xr"
    nais.EntityData.ParentYangName = "domain"
    nais.EntityData.SegmentPath = "nais"
    nais.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/" + nais.EntityData.SegmentPath
    nais.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nais.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nais.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nais.EntityData.Children = types.NewOrderedMap()
    nais.EntityData.Children.Append("nai", types.YChild{"Nai", nil})
    for i := range nais.Nai {
        nais.EntityData.Children.Append(types.GetSegmentPath(nais.Nai[i]), types.YChild{"Nai", nais.Nai[i]})
    }
    nais.EntityData.Leafs = types.NewOrderedMap()

    nais.EntityData.YListKeys = []string {}

    return &(nais.EntityData)
}

// MobileIp_Domains_Domain_Nais_Nai
// Network access identifier or Realm
type MobileIp_Domains_Domain_Nais_Nai struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MN Identifier. The type is string with length:
    // 1..125.
    NaiName interface{}

    // LMA for this MN. The type is string with length: 1..125.
    Lma interface{}

    // Access point network for this MN. The type is string with length: 1..125.
    Apn interface{}

    // Customer name for this MN. The type is string with length: 1..125.
    Customer interface{}

    // Service type for this MN. The type is ServiceType.
    Service interface{}

    // Network name (Address pool) for this MN. The type is string with length:
    // 1..125.
    Network interface{}
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetEntityData() *types.CommonEntityData {
    nai.EntityData.YFilter = nai.YFilter
    nai.EntityData.YangName = "nai"
    nai.EntityData.BundleName = "cisco_ios_xr"
    nai.EntityData.ParentYangName = "nais"
    nai.EntityData.SegmentPath = "nai" + types.AddKeyToken(nai.NaiName, "nai-name")
    nai.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/nais/" + nai.EntityData.SegmentPath
    nai.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nai.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nai.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nai.EntityData.Children = types.NewOrderedMap()
    nai.EntityData.Leafs = types.NewOrderedMap()
    nai.EntityData.Leafs.Append("nai-name", types.YLeaf{"NaiName", nai.NaiName})
    nai.EntityData.Leafs.Append("lma", types.YLeaf{"Lma", nai.Lma})
    nai.EntityData.Leafs.Append("apn", types.YLeaf{"Apn", nai.Apn})
    nai.EntityData.Leafs.Append("customer", types.YLeaf{"Customer", nai.Customer})
    nai.EntityData.Leafs.Append("service", types.YLeaf{"Service", nai.Service})
    nai.EntityData.Leafs.Append("network", types.YLeaf{"Network", nai.Network})

    nai.EntityData.YListKeys = []string {"NaiName"}

    return &(nai.EntityData)
}

// MobileIp_Domains_Domain_AuthenticateOption
// Authentication option between PMIPV6 entities
type MobileIp_Domains_Domain_AuthenticateOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetEntityData() *types.CommonEntityData {
    authenticateOption.EntityData.YFilter = authenticateOption.YFilter
    authenticateOption.EntityData.YangName = "authenticate-option"
    authenticateOption.EntityData.BundleName = "cisco_ios_xr"
    authenticateOption.EntityData.ParentYangName = "domain"
    authenticateOption.EntityData.SegmentPath = "authenticate-option"
    authenticateOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/" + authenticateOption.EntityData.SegmentPath
    authenticateOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticateOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticateOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticateOption.EntityData.Children = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authenticateOption.Spi})
    authenticateOption.EntityData.Leafs.Append("key", types.YLeaf{"Key", authenticateOption.Key})

    authenticateOption.EntityData.YListKeys = []string {}

    return &(authenticateOption.EntityData)
}

// MobileIp_Domains_Domain_Lmas
// Table of LMA
type MobileIp_Domains_Domain_Lmas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LMA within domain. The type is slice of MobileIp_Domains_Domain_Lmas_Lma.
    Lma []*MobileIp_Domains_Domain_Lmas_Lma
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetEntityData() *types.CommonEntityData {
    lmas.EntityData.YFilter = lmas.YFilter
    lmas.EntityData.YangName = "lmas"
    lmas.EntityData.BundleName = "cisco_ios_xr"
    lmas.EntityData.ParentYangName = "domain"
    lmas.EntityData.SegmentPath = "lmas"
    lmas.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/" + lmas.EntityData.SegmentPath
    lmas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmas.EntityData.Children = types.NewOrderedMap()
    lmas.EntityData.Children.Append("lma", types.YChild{"Lma", nil})
    for i := range lmas.Lma {
        lmas.EntityData.Children.Append(types.GetSegmentPath(lmas.Lma[i]), types.YChild{"Lma", lmas.Lma[i]})
    }
    lmas.EntityData.Leafs = types.NewOrderedMap()

    lmas.EntityData.YListKeys = []string {}

    return &(lmas.EntityData)
}

// MobileIp_Domains_Domain_Lmas_Lma
// LMA within domain
type MobileIp_Domains_Domain_Lmas_Lma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA Identifier. The type is string with length:
    // 1..125.
    LmaName interface{}
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetEntityData() *types.CommonEntityData {
    lma.EntityData.YFilter = lma.YFilter
    lma.EntityData.YangName = "lma"
    lma.EntityData.BundleName = "cisco_ios_xr"
    lma.EntityData.ParentYangName = "lmas"
    lma.EntityData.SegmentPath = "lma" + types.AddKeyToken(lma.LmaName, "lma-name")
    lma.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/domains/domain/lmas/" + lma.EntityData.SegmentPath
    lma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lma.EntityData.Children = types.NewOrderedMap()
    lma.EntityData.Leafs = types.NewOrderedMap()
    lma.EntityData.Leafs.Append("lma-name", types.YLeaf{"LmaName", lma.LmaName})

    lma.EntityData.YListKeys = []string {"LmaName"}

    return &(lma.EntityData)
}

// MobileIp_Lmas
// Table of LMA
type MobileIp_Lmas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PMIPv6 LMA configuration. The type is slice of MobileIp_Lmas_Lma.
    Lma []*MobileIp_Lmas_Lma
}

func (lmas *MobileIp_Lmas) GetEntityData() *types.CommonEntityData {
    lmas.EntityData.YFilter = lmas.YFilter
    lmas.EntityData.YangName = "lmas"
    lmas.EntityData.BundleName = "cisco_ios_xr"
    lmas.EntityData.ParentYangName = "mobile-ip"
    lmas.EntityData.SegmentPath = "lmas"
    lmas.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/" + lmas.EntityData.SegmentPath
    lmas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmas.EntityData.Children = types.NewOrderedMap()
    lmas.EntityData.Children.Append("lma", types.YChild{"Lma", nil})
    for i := range lmas.Lma {
        lmas.EntityData.Children.Append(types.GetSegmentPath(lmas.Lma[i]), types.YChild{"Lma", lmas.Lma[i]})
    }
    lmas.EntityData.Leafs = types.NewOrderedMap()

    lmas.EntityData.YListKeys = []string {}

    return &(lmas.EntityData)
}

// MobileIp_Lmas_Lma
// PMIPv6 LMA configuration
type MobileIp_Lmas_Lma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA name. The type is string with length: 1..125.
    LmaName interface{}

    // This attribute is a key. Domain name. The type is string with length:
    // 1..125.
    DomainName interface{}

    // generate gre key for LMA bindings. The type is interface{}.
    Generate interface{}

    // Specify the Admin Distance value. The type is interface{} with range:
    // 1..254.
    MobileRouteAd interface{}

    // enable ani option processing in LMA. The type is interface{}.
    Ani interface{}

    // enable multipath support for LMA. The type is interface{}.
    Multipath interface{}

    // enable dynamic mag learning for LMA. The type is interface{}.
    Dynamic interface{}

    // enforce heartbeat values to MAG. The type is interface{}.
    Enforce interface{}

    // Default MN profile for LMA. The type is string with length: 1..125.
    DefaultProfile interface{}

    // CN facing interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Mobile Map for this LMA. The type is string with length: 1..125.
    MobileMap interface{}

    // Feature related to interface with PGW. The type is interface{}.
    PgwSubsCont interface{}

    // LMA Binding Revocation Attributes.
    BindingRevocationAttributes MobileIp_Lmas_Lma_BindingRevocationAttributes

    // Radio access technology type config  this LMA.
    RatAttributes MobileIp_Lmas_Lma_RatAttributes

    // heartbeat config for this LMA.
    HeartBeatAttributes MobileIp_Lmas_Lma_HeartBeatAttributes

    // Table of LMAIPv6Address.
    Lmaipv6Addresses MobileIp_Lmas_Lma_Lmaipv6Addresses

    // LMA HNP options.
    Hnp MobileIp_Lmas_Lma_Hnp

    // Redistribute routes.
    Redistribute MobileIp_Lmas_Lma_Redistribute

    // AAA configuration for this LMA.
    Aaa MobileIp_Lmas_Lma_Aaa

    // DSCP for packets originating from this LMA.
    Dscp MobileIp_Lmas_Lma_Dscp

    // Table of LMAIPv4Address.
    Lmaipv4Addresses MobileIp_Lmas_Lma_Lmaipv4Addresses

    // Table of Role.
    Roles MobileIp_Lmas_Lma_Roles

    // LMA binding attributes.
    BindingAttributes MobileIp_Lmas_Lma_BindingAttributes

    // Table of MAG.
    Mags MobileIp_Lmas_Lma_Mags

    // tunnel attributes.
    TunnelAttributes MobileIp_Lmas_Lma_TunnelAttributes

    // Table of Service.
    Services MobileIp_Lmas_Lma_Services

    // Table of Network.
    Networks MobileIp_Lmas_Lma_Networks

    // Replay Protection Method.
    ReplayProtection MobileIp_Lmas_Lma_ReplayProtection
}

func (lma *MobileIp_Lmas_Lma) GetEntityData() *types.CommonEntityData {
    lma.EntityData.YFilter = lma.YFilter
    lma.EntityData.YangName = "lma"
    lma.EntityData.BundleName = "cisco_ios_xr"
    lma.EntityData.ParentYangName = "lmas"
    lma.EntityData.SegmentPath = "lma" + types.AddKeyToken(lma.LmaName, "lma-name") + types.AddKeyToken(lma.DomainName, "domain-name")
    lma.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/" + lma.EntityData.SegmentPath
    lma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lma.EntityData.Children = types.NewOrderedMap()
    lma.EntityData.Children.Append("binding-revocation-attributes", types.YChild{"BindingRevocationAttributes", &lma.BindingRevocationAttributes})
    lma.EntityData.Children.Append("rat-attributes", types.YChild{"RatAttributes", &lma.RatAttributes})
    lma.EntityData.Children.Append("heart-beat-attributes", types.YChild{"HeartBeatAttributes", &lma.HeartBeatAttributes})
    lma.EntityData.Children.Append("lmaipv6-addresses", types.YChild{"Lmaipv6Addresses", &lma.Lmaipv6Addresses})
    lma.EntityData.Children.Append("hnp", types.YChild{"Hnp", &lma.Hnp})
    lma.EntityData.Children.Append("redistribute", types.YChild{"Redistribute", &lma.Redistribute})
    lma.EntityData.Children.Append("aaa", types.YChild{"Aaa", &lma.Aaa})
    lma.EntityData.Children.Append("dscp", types.YChild{"Dscp", &lma.Dscp})
    lma.EntityData.Children.Append("lmaipv4-addresses", types.YChild{"Lmaipv4Addresses", &lma.Lmaipv4Addresses})
    lma.EntityData.Children.Append("roles", types.YChild{"Roles", &lma.Roles})
    lma.EntityData.Children.Append("binding-attributes", types.YChild{"BindingAttributes", &lma.BindingAttributes})
    lma.EntityData.Children.Append("mags", types.YChild{"Mags", &lma.Mags})
    lma.EntityData.Children.Append("tunnel-attributes", types.YChild{"TunnelAttributes", &lma.TunnelAttributes})
    lma.EntityData.Children.Append("services", types.YChild{"Services", &lma.Services})
    lma.EntityData.Children.Append("networks", types.YChild{"Networks", &lma.Networks})
    lma.EntityData.Children.Append("replay-protection", types.YChild{"ReplayProtection", &lma.ReplayProtection})
    lma.EntityData.Leafs = types.NewOrderedMap()
    lma.EntityData.Leafs.Append("lma-name", types.YLeaf{"LmaName", lma.LmaName})
    lma.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", lma.DomainName})
    lma.EntityData.Leafs.Append("generate", types.YLeaf{"Generate", lma.Generate})
    lma.EntityData.Leafs.Append("mobile-route-ad", types.YLeaf{"MobileRouteAd", lma.MobileRouteAd})
    lma.EntityData.Leafs.Append("ani", types.YLeaf{"Ani", lma.Ani})
    lma.EntityData.Leafs.Append("multipath", types.YLeaf{"Multipath", lma.Multipath})
    lma.EntityData.Leafs.Append("dynamic", types.YLeaf{"Dynamic", lma.Dynamic})
    lma.EntityData.Leafs.Append("enforce", types.YLeaf{"Enforce", lma.Enforce})
    lma.EntityData.Leafs.Append("default-profile", types.YLeaf{"DefaultProfile", lma.DefaultProfile})
    lma.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", lma.Interface})
    lma.EntityData.Leafs.Append("mobile-map", types.YLeaf{"MobileMap", lma.MobileMap})
    lma.EntityData.Leafs.Append("pgw-subs-cont", types.YLeaf{"PgwSubsCont", lma.PgwSubsCont})

    lma.EntityData.YListKeys = []string {"LmaName", "DomainName"}

    return &(lma.EntityData)
}

// MobileIp_Lmas_Lma_BindingRevocationAttributes
// LMA Binding Revocation Attributes
type MobileIp_Lmas_Lma_BindingRevocationAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Retransmissons Allowed for BRI Message. The type is interface{}
    // with range: 1..10.
    Retry interface{}

    // Time to wait before Retransmitting BRI Message.
    Delay MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetEntityData() *types.CommonEntityData {
    bindingRevocationAttributes.EntityData.YFilter = bindingRevocationAttributes.YFilter
    bindingRevocationAttributes.EntityData.YangName = "binding-revocation-attributes"
    bindingRevocationAttributes.EntityData.BundleName = "cisco_ios_xr"
    bindingRevocationAttributes.EntityData.ParentYangName = "lma"
    bindingRevocationAttributes.EntityData.SegmentPath = "binding-revocation-attributes"
    bindingRevocationAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + bindingRevocationAttributes.EntityData.SegmentPath
    bindingRevocationAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingRevocationAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingRevocationAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingRevocationAttributes.EntityData.Children = types.NewOrderedMap()
    bindingRevocationAttributes.EntityData.Children.Append("delay", types.YChild{"Delay", &bindingRevocationAttributes.Delay})
    bindingRevocationAttributes.EntityData.Leafs = types.NewOrderedMap()
    bindingRevocationAttributes.EntityData.Leafs.Append("retry", types.YLeaf{"Retry", bindingRevocationAttributes.Retry})

    bindingRevocationAttributes.EntityData.YListKeys = []string {}

    return &(bindingRevocationAttributes.EntityData)
}

// MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay
// Time to wait before Retransmitting BRI
// Message
type MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify in millisec. The type is interface{} with range: 500..65535.
    BrMin interface{}

    // Specify in millisec. The type is interface{} with range: 500..65535.
    BrMax interface{}
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "binding-revocation-attributes"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/binding-revocation-attributes/" + delay.EntityData.SegmentPath
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = types.NewOrderedMap()
    delay.EntityData.Leafs = types.NewOrderedMap()
    delay.EntityData.Leafs.Append("br-min", types.YLeaf{"BrMin", delay.BrMin})
    delay.EntityData.Leafs.Append("br-max", types.YLeaf{"BrMax", delay.BrMax})

    delay.EntityData.YListKeys = []string {}

    return &(delay.EntityData)
}

// MobileIp_Lmas_Lma_RatAttributes
// Radio access technology type config  this LMA
type MobileIp_Lmas_Lma_RatAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LMA rat type. The type is LmaRat.
    LmaRat interface{}

    // Specify the priority value. The type is interface{} with range: 1..255.
    PriorityValue interface{}
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetEntityData() *types.CommonEntityData {
    ratAttributes.EntityData.YFilter = ratAttributes.YFilter
    ratAttributes.EntityData.YangName = "rat-attributes"
    ratAttributes.EntityData.BundleName = "cisco_ios_xr"
    ratAttributes.EntityData.ParentYangName = "lma"
    ratAttributes.EntityData.SegmentPath = "rat-attributes"
    ratAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + ratAttributes.EntityData.SegmentPath
    ratAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ratAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ratAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ratAttributes.EntityData.Children = types.NewOrderedMap()
    ratAttributes.EntityData.Leafs = types.NewOrderedMap()
    ratAttributes.EntityData.Leafs.Append("lma-rat", types.YLeaf{"LmaRat", ratAttributes.LmaRat})
    ratAttributes.EntityData.Leafs.Append("priority-value", types.YLeaf{"PriorityValue", ratAttributes.PriorityValue})

    ratAttributes.EntityData.YListKeys = []string {}

    return &(ratAttributes.EntityData)
}

// MobileIp_Lmas_Lma_HeartBeatAttributes
// heartbeat config for this LMA
type MobileIp_Lmas_Lma_HeartBeatAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the interval value in second. The type is interface{} with range:
    // 10..3600.
    Interval interface{}

    // Specify the retry value. The type is interface{} with range: 1..10.
    Retries interface{}

    // Specify the timeout value. The type is interface{} with range: 1..3600.
    Timeout interface{}
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetEntityData() *types.CommonEntityData {
    heartBeatAttributes.EntityData.YFilter = heartBeatAttributes.YFilter
    heartBeatAttributes.EntityData.YangName = "heart-beat-attributes"
    heartBeatAttributes.EntityData.BundleName = "cisco_ios_xr"
    heartBeatAttributes.EntityData.ParentYangName = "lma"
    heartBeatAttributes.EntityData.SegmentPath = "heart-beat-attributes"
    heartBeatAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + heartBeatAttributes.EntityData.SegmentPath
    heartBeatAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    heartBeatAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    heartBeatAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    heartBeatAttributes.EntityData.Children = types.NewOrderedMap()
    heartBeatAttributes.EntityData.Leafs = types.NewOrderedMap()
    heartBeatAttributes.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", heartBeatAttributes.Interval})
    heartBeatAttributes.EntityData.Leafs.Append("retries", types.YLeaf{"Retries", heartBeatAttributes.Retries})
    heartBeatAttributes.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", heartBeatAttributes.Timeout})

    heartBeatAttributes.EntityData.YListKeys = []string {}

    return &(heartBeatAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Lmaipv6Addresses
// Table of LMAIPv6Address
type MobileIp_Lmas_Lma_Lmaipv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IPv6 address for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address.
    Lmaipv6Address []*MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetEntityData() *types.CommonEntityData {
    lmaipv6Addresses.EntityData.YFilter = lmaipv6Addresses.YFilter
    lmaipv6Addresses.EntityData.YangName = "lmaipv6-addresses"
    lmaipv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    lmaipv6Addresses.EntityData.ParentYangName = "lma"
    lmaipv6Addresses.EntityData.SegmentPath = "lmaipv6-addresses"
    lmaipv6Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + lmaipv6Addresses.EntityData.SegmentPath
    lmaipv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmaipv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmaipv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmaipv6Addresses.EntityData.Children = types.NewOrderedMap()
    lmaipv6Addresses.EntityData.Children.Append("lmaipv6-address", types.YChild{"Lmaipv6Address", nil})
    for i := range lmaipv6Addresses.Lmaipv6Address {
        lmaipv6Addresses.EntityData.Children.Append(types.GetSegmentPath(lmaipv6Addresses.Lmaipv6Address[i]), types.YChild{"Lmaipv6Address", lmaipv6Addresses.Lmaipv6Address[i]})
    }
    lmaipv6Addresses.EntityData.Leafs = types.NewOrderedMap()

    lmaipv6Addresses.EntityData.YListKeys = []string {}

    return &(lmaipv6Addresses.EntityData)
}

// MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address
// Configure IPv6 address for this LMA
type MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetEntityData() *types.CommonEntityData {
    lmaipv6Address.EntityData.YFilter = lmaipv6Address.YFilter
    lmaipv6Address.EntityData.YangName = "lmaipv6-address"
    lmaipv6Address.EntityData.BundleName = "cisco_ios_xr"
    lmaipv6Address.EntityData.ParentYangName = "lmaipv6-addresses"
    lmaipv6Address.EntityData.SegmentPath = "lmaipv6-address" + types.AddKeyToken(lmaipv6Address.Address, "address")
    lmaipv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/lmaipv6-addresses/" + lmaipv6Address.EntityData.SegmentPath
    lmaipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmaipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmaipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmaipv6Address.EntityData.Children = types.NewOrderedMap()
    lmaipv6Address.EntityData.Leafs = types.NewOrderedMap()
    lmaipv6Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", lmaipv6Address.Address})

    lmaipv6Address.EntityData.YListKeys = []string {"Address"}

    return &(lmaipv6Address.EntityData)
}

// MobileIp_Lmas_Lma_Hnp
// LMA HNP options
type MobileIp_Lmas_Lma_Hnp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // maximum HNPs allowed per MN interface. The type is interface{} with range:
    // 1..3.
    Maximum interface{}
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetEntityData() *types.CommonEntityData {
    hnp.EntityData.YFilter = hnp.YFilter
    hnp.EntityData.YangName = "hnp"
    hnp.EntityData.BundleName = "cisco_ios_xr"
    hnp.EntityData.ParentYangName = "lma"
    hnp.EntityData.SegmentPath = "hnp"
    hnp.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + hnp.EntityData.SegmentPath
    hnp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hnp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hnp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hnp.EntityData.Children = types.NewOrderedMap()
    hnp.EntityData.Leafs = types.NewOrderedMap()
    hnp.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", hnp.Maximum})

    hnp.EntityData.YListKeys = []string {}

    return &(hnp.EntityData)
}

// MobileIp_Lmas_Lma_Redistribute
// Redistribute routes
type MobileIp_Lmas_Lma_Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute type. The type is RedistType.
    RedistType interface{}

    // Redistribute sub-type. The type is RedistSubType.
    RedistSubType interface{}
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetEntityData() *types.CommonEntityData {
    redistribute.EntityData.YFilter = redistribute.YFilter
    redistribute.EntityData.YangName = "redistribute"
    redistribute.EntityData.BundleName = "cisco_ios_xr"
    redistribute.EntityData.ParentYangName = "lma"
    redistribute.EntityData.SegmentPath = "redistribute"
    redistribute.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + redistribute.EntityData.SegmentPath
    redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribute.EntityData.Children = types.NewOrderedMap()
    redistribute.EntityData.Leafs = types.NewOrderedMap()
    redistribute.EntityData.Leafs.Append("redist-type", types.YLeaf{"RedistType", redistribute.RedistType})
    redistribute.EntityData.Leafs.Append("redist-sub-type", types.YLeaf{"RedistSubType", redistribute.RedistSubType})

    redistribute.EntityData.YListKeys = []string {}

    return &(redistribute.EntityData)
}

// MobileIp_Lmas_Lma_Aaa
// AAA configuration for this LMA
type MobileIp_Lmas_Lma_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA accounting for this LMA.
    Accounting MobileIp_Lmas_Lma_Aaa_Accounting
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "lma"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + aaa.EntityData.SegmentPath
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("accounting", types.YChild{"Accounting", &aaa.Accounting})
    aaa.EntityData.Leafs = types.NewOrderedMap()

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// MobileIp_Lmas_Lma_Aaa_Accounting
// AAA accounting for this LMA
type MobileIp_Lmas_Lma_Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set constant integer. The type is interface{}.
    Enable interface{}

    // Interim acounting interval (in minutes). The type is interface{} with
    // range: 1..86400. Units are minute.
    InterimInterval interface{}
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/aaa/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", accounting.Enable})
    accounting.EntityData.Leafs.Append("interim-interval", types.YLeaf{"InterimInterval", accounting.InterimInterval})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// MobileIp_Lmas_Lma_Dscp
// DSCP for packets originating from this LMA
type MobileIp_Lmas_Lma_Dscp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 1..63.
    Value interface{}

    // Set constant integer. The type is interface{}.
    Force interface{}
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetEntityData() *types.CommonEntityData {
    dscp.EntityData.YFilter = dscp.YFilter
    dscp.EntityData.YangName = "dscp"
    dscp.EntityData.BundleName = "cisco_ios_xr"
    dscp.EntityData.ParentYangName = "lma"
    dscp.EntityData.SegmentPath = "dscp"
    dscp.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + dscp.EntityData.SegmentPath
    dscp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscp.EntityData.Children = types.NewOrderedMap()
    dscp.EntityData.Leafs = types.NewOrderedMap()
    dscp.EntityData.Leafs.Append("value", types.YLeaf{"Value", dscp.Value})
    dscp.EntityData.Leafs.Append("force", types.YLeaf{"Force", dscp.Force})

    dscp.EntityData.YListKeys = []string {}

    return &(dscp.EntityData)
}

// MobileIp_Lmas_Lma_Lmaipv4Addresses
// Table of LMAIPv4Address
type MobileIp_Lmas_Lma_Lmaipv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure IPv4 address for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address.
    Lmaipv4Address []*MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetEntityData() *types.CommonEntityData {
    lmaipv4Addresses.EntityData.YFilter = lmaipv4Addresses.YFilter
    lmaipv4Addresses.EntityData.YangName = "lmaipv4-addresses"
    lmaipv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    lmaipv4Addresses.EntityData.ParentYangName = "lma"
    lmaipv4Addresses.EntityData.SegmentPath = "lmaipv4-addresses"
    lmaipv4Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + lmaipv4Addresses.EntityData.SegmentPath
    lmaipv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmaipv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmaipv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmaipv4Addresses.EntityData.Children = types.NewOrderedMap()
    lmaipv4Addresses.EntityData.Children.Append("lmaipv4-address", types.YChild{"Lmaipv4Address", nil})
    for i := range lmaipv4Addresses.Lmaipv4Address {
        lmaipv4Addresses.EntityData.Children.Append(types.GetSegmentPath(lmaipv4Addresses.Lmaipv4Address[i]), types.YChild{"Lmaipv4Address", lmaipv4Addresses.Lmaipv4Address[i]})
    }
    lmaipv4Addresses.EntityData.Leafs = types.NewOrderedMap()

    lmaipv4Addresses.EntityData.YListKeys = []string {}

    return &(lmaipv4Addresses.EntityData)
}

// MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address
// Configure IPv4 address for this LMA
type MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetEntityData() *types.CommonEntityData {
    lmaipv4Address.EntityData.YFilter = lmaipv4Address.YFilter
    lmaipv4Address.EntityData.YangName = "lmaipv4-address"
    lmaipv4Address.EntityData.BundleName = "cisco_ios_xr"
    lmaipv4Address.EntityData.ParentYangName = "lmaipv4-addresses"
    lmaipv4Address.EntityData.SegmentPath = "lmaipv4-address" + types.AddKeyToken(lmaipv4Address.Address, "address")
    lmaipv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/lmaipv4-addresses/" + lmaipv4Address.EntityData.SegmentPath
    lmaipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmaipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmaipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmaipv4Address.EntityData.Children = types.NewOrderedMap()
    lmaipv4Address.EntityData.Leafs = types.NewOrderedMap()
    lmaipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", lmaipv4Address.Address})

    lmaipv4Address.EntityData.YListKeys = []string {"Address"}

    return &(lmaipv4Address.EntityData)
}

// MobileIp_Lmas_Lma_Roles
// Table of Role
type MobileIp_Lmas_Lma_Roles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Role of this LMA. The type is slice of MobileIp_Lmas_Lma_Roles_Role.
    Role []*MobileIp_Lmas_Lma_Roles_Role
}

func (roles *MobileIp_Lmas_Lma_Roles) GetEntityData() *types.CommonEntityData {
    roles.EntityData.YFilter = roles.YFilter
    roles.EntityData.YangName = "roles"
    roles.EntityData.BundleName = "cisco_ios_xr"
    roles.EntityData.ParentYangName = "lma"
    roles.EntityData.SegmentPath = "roles"
    roles.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + roles.EntityData.SegmentPath
    roles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    roles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    roles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    roles.EntityData.Children = types.NewOrderedMap()
    roles.EntityData.Children.Append("role", types.YChild{"Role", nil})
    for i := range roles.Role {
        roles.EntityData.Children.Append(types.GetSegmentPath(roles.Role[i]), types.YChild{"Role", roles.Role[i]})
    }
    roles.EntityData.Leafs = types.NewOrderedMap()

    roles.EntityData.YListKeys = []string {}

    return &(roles.EntityData)
}

// MobileIp_Lmas_Lma_Roles_Role
// Role of this LMA
type MobileIp_Lmas_Lma_Roles_Role struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA role mode. The type is LmaRole.
    LmaRole interface{}
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetEntityData() *types.CommonEntityData {
    role.EntityData.YFilter = role.YFilter
    role.EntityData.YangName = "role"
    role.EntityData.BundleName = "cisco_ios_xr"
    role.EntityData.ParentYangName = "roles"
    role.EntityData.SegmentPath = "role" + types.AddKeyToken(role.LmaRole, "lma-role")
    role.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/roles/" + role.EntityData.SegmentPath
    role.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    role.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    role.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    role.EntityData.Children = types.NewOrderedMap()
    role.EntityData.Leafs = types.NewOrderedMap()
    role.EntityData.Leafs.Append("lma-role", types.YLeaf{"LmaRole", role.LmaRole})

    role.EntityData.YListKeys = []string {"LmaRole"}

    return &(role.EntityData)
}

// MobileIp_Lmas_Lma_BindingAttributes
// LMA binding attributes
type MobileIp_Lmas_Lma_BindingAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify in seconds, (multiples of 4). The type is interface{} with range:
    // 4..65000. Units are second.
    RefreshTime interface{}

    // bce delete wait time interval. The type is interface{} with range:
    // 100..65535.
    DeleteWaitInterval interface{}

    // bce create wait time interval. The type is interface{} with range:
    // 100..65535.
    CreateWaitInterval interface{}

    // Maximum bce lifetime permitted. The type is interface{} with range:
    // 10..65535. Units are second.
    MaxLifeTime interface{}

    // Specify max. number of bindings. The type is interface{} with range:
    // 1..128000.
    Maximum interface{}
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetEntityData() *types.CommonEntityData {
    bindingAttributes.EntityData.YFilter = bindingAttributes.YFilter
    bindingAttributes.EntityData.YangName = "binding-attributes"
    bindingAttributes.EntityData.BundleName = "cisco_ios_xr"
    bindingAttributes.EntityData.ParentYangName = "lma"
    bindingAttributes.EntityData.SegmentPath = "binding-attributes"
    bindingAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + bindingAttributes.EntityData.SegmentPath
    bindingAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingAttributes.EntityData.Children = types.NewOrderedMap()
    bindingAttributes.EntityData.Leafs = types.NewOrderedMap()
    bindingAttributes.EntityData.Leafs.Append("refresh-time", types.YLeaf{"RefreshTime", bindingAttributes.RefreshTime})
    bindingAttributes.EntityData.Leafs.Append("delete-wait-interval", types.YLeaf{"DeleteWaitInterval", bindingAttributes.DeleteWaitInterval})
    bindingAttributes.EntityData.Leafs.Append("create-wait-interval", types.YLeaf{"CreateWaitInterval", bindingAttributes.CreateWaitInterval})
    bindingAttributes.EntityData.Leafs.Append("max-life-time", types.YLeaf{"MaxLifeTime", bindingAttributes.MaxLifeTime})
    bindingAttributes.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", bindingAttributes.Maximum})

    bindingAttributes.EntityData.YListKeys = []string {}

    return &(bindingAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Mags
// Table of MAG
type MobileIp_Lmas_Lma_Mags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAG within LMA. The type is slice of MobileIp_Lmas_Lma_Mags_Mag.
    Mag []*MobileIp_Lmas_Lma_Mags_Mag
}

func (mags *MobileIp_Lmas_Lma_Mags) GetEntityData() *types.CommonEntityData {
    mags.EntityData.YFilter = mags.YFilter
    mags.EntityData.YangName = "mags"
    mags.EntityData.BundleName = "cisco_ios_xr"
    mags.EntityData.ParentYangName = "lma"
    mags.EntityData.SegmentPath = "mags"
    mags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + mags.EntityData.SegmentPath
    mags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mags.EntityData.Children = types.NewOrderedMap()
    mags.EntityData.Children.Append("mag", types.YChild{"Mag", nil})
    for i := range mags.Mag {
        mags.EntityData.Children.Append(types.GetSegmentPath(mags.Mag[i]), types.YChild{"Mag", mags.Mag[i]})
    }
    mags.EntityData.Leafs = types.NewOrderedMap()

    mags.EntityData.YListKeys = []string {}

    return &(mags.EntityData)
}

// MobileIp_Lmas_Lma_Mags_Mag
// MAG within LMA
type MobileIp_Lmas_Lma_Mags_Mag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MAG identifier. The type is string with length:
    // 1..125.
    MagName interface{}

    // This attribute is a key. Domain name. The type is string with length:
    // 1..125.
    DomainName interface{}

    // Encapsulation option between PMIPV6 entities. The type is EncapOpt.
    EncapOption interface{}

    // Configure IPv4 address for this MAG. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Configure IPv6 address for this MAG. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // static tunnel for this peer MAG. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Tunnel interface{}

    // Authentication option between PMIPV6 entities.
    AuthenticateOption MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption

    // DSCP for packets originating from this MAG.
    Dscp MobileIp_Lmas_Lma_Mags_Mag_Dscp
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetEntityData() *types.CommonEntityData {
    mag.EntityData.YFilter = mag.YFilter
    mag.EntityData.YangName = "mag"
    mag.EntityData.BundleName = "cisco_ios_xr"
    mag.EntityData.ParentYangName = "mags"
    mag.EntityData.SegmentPath = "mag" + types.AddKeyToken(mag.MagName, "mag-name") + types.AddKeyToken(mag.DomainName, "domain-name")
    mag.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/mags/" + mag.EntityData.SegmentPath
    mag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mag.EntityData.Children = types.NewOrderedMap()
    mag.EntityData.Children.Append("authenticate-option", types.YChild{"AuthenticateOption", &mag.AuthenticateOption})
    mag.EntityData.Children.Append("dscp", types.YChild{"Dscp", &mag.Dscp})
    mag.EntityData.Leafs = types.NewOrderedMap()
    mag.EntityData.Leafs.Append("mag-name", types.YLeaf{"MagName", mag.MagName})
    mag.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", mag.DomainName})
    mag.EntityData.Leafs.Append("encap-option", types.YLeaf{"EncapOption", mag.EncapOption})
    mag.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", mag.Ipv4Address})
    mag.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", mag.Ipv6Address})
    mag.EntityData.Leafs.Append("tunnel", types.YLeaf{"Tunnel", mag.Tunnel})

    mag.EntityData.YListKeys = []string {"MagName", "DomainName"}

    return &(mag.EntityData)
}

// MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption
// Authentication option between PMIPV6
// entities
type MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetEntityData() *types.CommonEntityData {
    authenticateOption.EntityData.YFilter = authenticateOption.YFilter
    authenticateOption.EntityData.YangName = "authenticate-option"
    authenticateOption.EntityData.BundleName = "cisco_ios_xr"
    authenticateOption.EntityData.ParentYangName = "mag"
    authenticateOption.EntityData.SegmentPath = "authenticate-option"
    authenticateOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/mags/mag/" + authenticateOption.EntityData.SegmentPath
    authenticateOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticateOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticateOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticateOption.EntityData.Children = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authenticateOption.Spi})
    authenticateOption.EntityData.Leafs.Append("key", types.YLeaf{"Key", authenticateOption.Key})

    authenticateOption.EntityData.YListKeys = []string {}

    return &(authenticateOption.EntityData)
}

// MobileIp_Lmas_Lma_Mags_Mag_Dscp
// DSCP for packets originating from this MAG
type MobileIp_Lmas_Lma_Mags_Mag_Dscp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 1..63.
    Value interface{}

    // Set constant integer. The type is interface{}.
    Force interface{}
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetEntityData() *types.CommonEntityData {
    dscp.EntityData.YFilter = dscp.YFilter
    dscp.EntityData.YangName = "dscp"
    dscp.EntityData.BundleName = "cisco_ios_xr"
    dscp.EntityData.ParentYangName = "mag"
    dscp.EntityData.SegmentPath = "dscp"
    dscp.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/mags/mag/" + dscp.EntityData.SegmentPath
    dscp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dscp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dscp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dscp.EntityData.Children = types.NewOrderedMap()
    dscp.EntityData.Leafs = types.NewOrderedMap()
    dscp.EntityData.Leafs.Append("value", types.YLeaf{"Value", dscp.Value})
    dscp.EntityData.Leafs.Append("force", types.YLeaf{"Force", dscp.Force})

    dscp.EntityData.YListKeys = []string {}

    return &(dscp.EntityData)
}

// MobileIp_Lmas_Lma_TunnelAttributes
// tunnel attributes
type MobileIp_Lmas_Lma_TunnelAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // maximum transmission unit for tunnel. The type is interface{} with range:
    // 68..17916.
    Mtu interface{}

    // access list to apply for tunnel. The type is string with length: 1..125.
    Acl interface{}
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetEntityData() *types.CommonEntityData {
    tunnelAttributes.EntityData.YFilter = tunnelAttributes.YFilter
    tunnelAttributes.EntityData.YangName = "tunnel-attributes"
    tunnelAttributes.EntityData.BundleName = "cisco_ios_xr"
    tunnelAttributes.EntityData.ParentYangName = "lma"
    tunnelAttributes.EntityData.SegmentPath = "tunnel-attributes"
    tunnelAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + tunnelAttributes.EntityData.SegmentPath
    tunnelAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelAttributes.EntityData.Children = types.NewOrderedMap()
    tunnelAttributes.EntityData.Leafs = types.NewOrderedMap()
    tunnelAttributes.EntityData.Leafs.Append("mtu", types.YLeaf{"Mtu", tunnelAttributes.Mtu})
    tunnelAttributes.EntityData.Leafs.Append("acl", types.YLeaf{"Acl", tunnelAttributes.Acl})

    tunnelAttributes.EntityData.YListKeys = []string {}

    return &(tunnelAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Services
// Table of Service
type MobileIp_Lmas_Lma_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Service of this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service.
    Service []*MobileIp_Lmas_Lma_Services_Service
}

func (services *MobileIp_Lmas_Lma_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "lma"
    services.EntityData.SegmentPath = "services"
    services.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + services.EntityData.SegmentPath
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = types.NewOrderedMap()
    services.EntityData.Children.Append("service", types.YChild{"Service", nil})
    for i := range services.Service {
        services.EntityData.Children.Append(types.GetSegmentPath(services.Service[i]), types.YChild{"Service", services.Service[i]})
    }
    services.EntityData.Leafs = types.NewOrderedMap()

    services.EntityData.YListKeys = []string {}

    return &(services.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service
// Service of this LMA
type MobileIp_Lmas_Lma_Services_Service struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LMA service mode. The type is LmaService.
    LmaService interface{}

    // mnp limit config for all customer's. The type is interface{} with range:
    // 1..4000000.
    MnpCustomer interface{}

    // mnp limit config for all logical mn's. The type is interface{} with range:
    // 1..32.
    MnpIpv4Lmn interface{}

    // mnp limit config for all logical mn's. The type is interface{} with range:
    // 1..32.
    MnpIpv6Lmn interface{}

    // mnp limit config for all logical mn's. The type is interface{} with range:
    // 1..32.
    MnpLmn interface{}

    // Ignore HoA/HNP option. The type is interface{}.
    IgnoreHomeAddress interface{}

    // mnp limit config for all customer's. The type is interface{} with range:
    // 1..4000000.
    MnpIpv4Customer interface{}

    // mnp limit config for all customer's. The type is interface{} with range:
    // 1..4000000.
    MnpIpv6Customer interface{}

    // Table of Customer.
    Customers MobileIp_Lmas_Lma_Services_Service_Customers
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetEntityData() *types.CommonEntityData {
    service.EntityData.YFilter = service.YFilter
    service.EntityData.YangName = "service"
    service.EntityData.BundleName = "cisco_ios_xr"
    service.EntityData.ParentYangName = "services"
    service.EntityData.SegmentPath = "service" + types.AddKeyToken(service.LmaService, "lma-service")
    service.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/" + service.EntityData.SegmentPath
    service.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    service.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    service.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    service.EntityData.Children = types.NewOrderedMap()
    service.EntityData.Children.Append("customers", types.YChild{"Customers", &service.Customers})
    service.EntityData.Leafs = types.NewOrderedMap()
    service.EntityData.Leafs.Append("lma-service", types.YLeaf{"LmaService", service.LmaService})
    service.EntityData.Leafs.Append("mnp-customer", types.YLeaf{"MnpCustomer", service.MnpCustomer})
    service.EntityData.Leafs.Append("mnp-ipv4-lmn", types.YLeaf{"MnpIpv4Lmn", service.MnpIpv4Lmn})
    service.EntityData.Leafs.Append("mnp-ipv6-lmn", types.YLeaf{"MnpIpv6Lmn", service.MnpIpv6Lmn})
    service.EntityData.Leafs.Append("mnp-lmn", types.YLeaf{"MnpLmn", service.MnpLmn})
    service.EntityData.Leafs.Append("ignore-home-address", types.YLeaf{"IgnoreHomeAddress", service.IgnoreHomeAddress})
    service.EntityData.Leafs.Append("mnp-ipv4-customer", types.YLeaf{"MnpIpv4Customer", service.MnpIpv4Customer})
    service.EntityData.Leafs.Append("mnp-ipv6-customer", types.YLeaf{"MnpIpv6Customer", service.MnpIpv6Customer})

    service.EntityData.YListKeys = []string {"LmaService"}

    return &(service.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers
// Table of Customer
type MobileIp_Lmas_Lma_Services_Service_Customers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // customer configuration on this mobile local loop service. The type is slice
    // of MobileIp_Lmas_Lma_Services_Service_Customers_Customer.
    Customer []*MobileIp_Lmas_Lma_Services_Service_Customers_Customer
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetEntityData() *types.CommonEntityData {
    customers.EntityData.YFilter = customers.YFilter
    customers.EntityData.YangName = "customers"
    customers.EntityData.BundleName = "cisco_ios_xr"
    customers.EntityData.ParentYangName = "service"
    customers.EntityData.SegmentPath = "customers"
    customers.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/" + customers.EntityData.SegmentPath
    customers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customers.EntityData.Children = types.NewOrderedMap()
    customers.EntityData.Children.Append("customer", types.YChild{"Customer", nil})
    for i := range customers.Customer {
        customers.EntityData.Children.Append(types.GetSegmentPath(customers.Customer[i]), types.YChild{"Customer", customers.Customer[i]})
    }
    customers.EntityData.Leafs = types.NewOrderedMap()

    customers.EntityData.YListKeys = []string {}

    return &(customers.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer
// customer configuration on this mobile local
// loop service
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Customer name. The type is string with length:
    // 1..32.
    CustomerName interface{}

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // mnp limit config for customer. The type is interface{} with range:
    // 1..4000000.
    MnpCustomer interface{}

    // mnp limit config for customer specific logical mn. The type is interface{}
    // with range: 1..32.
    MnpIpv4Lmn interface{}

    // mnp limit config for customer specific logical mn. The type is interface{}
    // with range: 1..32.
    MnpIpv6Lmn interface{}

    // mnp limit config for customer specific logical mn. The type is interface{}
    // with range: 1..32.
    MnpLmn interface{}

    // mnp limit config for customer. The type is interface{} with range:
    // 1..4000000.
    MnpIpv4Customer interface{}

    // mnp limit config for customer. The type is interface{} with range:
    // 1..4000000.
    MnpIpv6Customer interface{}

    // Specify the Admin Distance value. The type is interface{} with range:
    // 1..254.
    MobileRouteAd interface{}

    // Aggregate bandwidth across all logical MNs. The type is interface{} with
    // range: 1..4294967295. Units are kbit/s.
    BandwidthAggregate interface{}

    // Authentication option between PMIPV6 entities.
    AuthenticateOption MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption

    // heartbeat config for this Customer.
    HeartBeatAttributes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes

    // Table of Transport.
    Transports MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports

    // network parameters for the customer.
    NetworkAttributes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes

    // Customer specific GRE key.
    GreKey MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey

    // Customer specific binding attributes.
    BindingAttributes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetEntityData() *types.CommonEntityData {
    customer.EntityData.YFilter = customer.YFilter
    customer.EntityData.YangName = "customer"
    customer.EntityData.BundleName = "cisco_ios_xr"
    customer.EntityData.ParentYangName = "customers"
    customer.EntityData.SegmentPath = "customer" + types.AddKeyToken(customer.CustomerName, "customer-name") + types.AddKeyToken(customer.VrfName, "vrf-name")
    customer.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/" + customer.EntityData.SegmentPath
    customer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customer.EntityData.Children = types.NewOrderedMap()
    customer.EntityData.Children.Append("authenticate-option", types.YChild{"AuthenticateOption", &customer.AuthenticateOption})
    customer.EntityData.Children.Append("heart-beat-attributes", types.YChild{"HeartBeatAttributes", &customer.HeartBeatAttributes})
    customer.EntityData.Children.Append("transports", types.YChild{"Transports", &customer.Transports})
    customer.EntityData.Children.Append("network-attributes", types.YChild{"NetworkAttributes", &customer.NetworkAttributes})
    customer.EntityData.Children.Append("gre-key", types.YChild{"GreKey", &customer.GreKey})
    customer.EntityData.Children.Append("binding-attributes", types.YChild{"BindingAttributes", &customer.BindingAttributes})
    customer.EntityData.Leafs = types.NewOrderedMap()
    customer.EntityData.Leafs.Append("customer-name", types.YLeaf{"CustomerName", customer.CustomerName})
    customer.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", customer.VrfName})
    customer.EntityData.Leafs.Append("mnp-customer", types.YLeaf{"MnpCustomer", customer.MnpCustomer})
    customer.EntityData.Leafs.Append("mnp-ipv4-lmn", types.YLeaf{"MnpIpv4Lmn", customer.MnpIpv4Lmn})
    customer.EntityData.Leafs.Append("mnp-ipv6-lmn", types.YLeaf{"MnpIpv6Lmn", customer.MnpIpv6Lmn})
    customer.EntityData.Leafs.Append("mnp-lmn", types.YLeaf{"MnpLmn", customer.MnpLmn})
    customer.EntityData.Leafs.Append("mnp-ipv4-customer", types.YLeaf{"MnpIpv4Customer", customer.MnpIpv4Customer})
    customer.EntityData.Leafs.Append("mnp-ipv6-customer", types.YLeaf{"MnpIpv6Customer", customer.MnpIpv6Customer})
    customer.EntityData.Leafs.Append("mobile-route-ad", types.YLeaf{"MobileRouteAd", customer.MobileRouteAd})
    customer.EntityData.Leafs.Append("bandwidth-aggregate", types.YLeaf{"BandwidthAggregate", customer.BandwidthAggregate})

    customer.EntityData.YListKeys = []string {"CustomerName", "VrfName"}

    return &(customer.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption
// Authentication option between PMIPV6
// entities
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetEntityData() *types.CommonEntityData {
    authenticateOption.EntityData.YFilter = authenticateOption.YFilter
    authenticateOption.EntityData.YangName = "authenticate-option"
    authenticateOption.EntityData.BundleName = "cisco_ios_xr"
    authenticateOption.EntityData.ParentYangName = "customer"
    authenticateOption.EntityData.SegmentPath = "authenticate-option"
    authenticateOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + authenticateOption.EntityData.SegmentPath
    authenticateOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticateOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticateOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticateOption.EntityData.Children = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs = types.NewOrderedMap()
    authenticateOption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authenticateOption.Spi})
    authenticateOption.EntityData.Leafs.Append("key", types.YLeaf{"Key", authenticateOption.Key})

    authenticateOption.EntityData.YListKeys = []string {}

    return &(authenticateOption.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes
// heartbeat config for this Customer
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the interval value in second. The type is interface{} with range:
    // 10..3600.
    Interval interface{}

    // Specify the retry value. The type is interface{} with range: 1..10.
    Retries interface{}

    // Specify the timeout value. The type is interface{} with range: 1..3600.
    Timeout interface{}
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetEntityData() *types.CommonEntityData {
    heartBeatAttributes.EntityData.YFilter = heartBeatAttributes.YFilter
    heartBeatAttributes.EntityData.YangName = "heart-beat-attributes"
    heartBeatAttributes.EntityData.BundleName = "cisco_ios_xr"
    heartBeatAttributes.EntityData.ParentYangName = "customer"
    heartBeatAttributes.EntityData.SegmentPath = "heart-beat-attributes"
    heartBeatAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + heartBeatAttributes.EntityData.SegmentPath
    heartBeatAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    heartBeatAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    heartBeatAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    heartBeatAttributes.EntityData.Children = types.NewOrderedMap()
    heartBeatAttributes.EntityData.Leafs = types.NewOrderedMap()
    heartBeatAttributes.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", heartBeatAttributes.Interval})
    heartBeatAttributes.EntityData.Leafs.Append("retries", types.YLeaf{"Retries", heartBeatAttributes.Retries})
    heartBeatAttributes.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", heartBeatAttributes.Timeout})

    heartBeatAttributes.EntityData.YListKeys = []string {}

    return &(heartBeatAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports
// Table of Transport
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Customer transport attributes. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport.
    Transport []*MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetEntityData() *types.CommonEntityData {
    transports.EntityData.YFilter = transports.YFilter
    transports.EntityData.YangName = "transports"
    transports.EntityData.BundleName = "cisco_ios_xr"
    transports.EntityData.ParentYangName = "customer"
    transports.EntityData.SegmentPath = "transports"
    transports.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + transports.EntityData.SegmentPath
    transports.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transports.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transports.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transports.EntityData.Children = types.NewOrderedMap()
    transports.EntityData.Children.Append("transport", types.YChild{"Transport", nil})
    for i := range transports.Transport {
        transports.EntityData.Children.Append(types.GetSegmentPath(transports.Transport[i]), types.YChild{"Transport", transports.Transport[i]})
    }
    transports.EntityData.Leafs = types.NewOrderedMap()

    transports.EntityData.YListKeys = []string {}

    return &(transports.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport
// Customer transport attributes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with length: 1..125.
    VrfName interface{}

    // Configure IPv4 address for this LMA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Configure IPv6 address for this LMA. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xr"
    transport.EntityData.ParentYangName = "transports"
    transport.EntityData.SegmentPath = "transport" + types.AddKeyToken(transport.VrfName, "vrf-name")
    transport.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/transports/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Leafs = types.NewOrderedMap()
    transport.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", transport.VrfName})
    transport.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", transport.Ipv4Address})
    transport.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", transport.Ipv6Address})

    transport.EntityData.YListKeys = []string {"VrfName"}

    return &(transport.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes
// network parameters for the customer
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // not authorize the network prefixes. The type is interface{}.
    Unauthorize interface{}

    // Table of Authorize.
    Authorizes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetEntityData() *types.CommonEntityData {
    networkAttributes.EntityData.YFilter = networkAttributes.YFilter
    networkAttributes.EntityData.YangName = "network-attributes"
    networkAttributes.EntityData.BundleName = "cisco_ios_xr"
    networkAttributes.EntityData.ParentYangName = "customer"
    networkAttributes.EntityData.SegmentPath = "network-attributes"
    networkAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + networkAttributes.EntityData.SegmentPath
    networkAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networkAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networkAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networkAttributes.EntityData.Children = types.NewOrderedMap()
    networkAttributes.EntityData.Children.Append("authorizes", types.YChild{"Authorizes", &networkAttributes.Authorizes})
    networkAttributes.EntityData.Leafs = types.NewOrderedMap()
    networkAttributes.EntityData.Leafs.Append("unauthorize", types.YLeaf{"Unauthorize", networkAttributes.Unauthorize})

    networkAttributes.EntityData.YListKeys = []string {}

    return &(networkAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes
// Table of Authorize
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // not authorize the network prefixes. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize.
    Authorize []*MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetEntityData() *types.CommonEntityData {
    authorizes.EntityData.YFilter = authorizes.YFilter
    authorizes.EntityData.YangName = "authorizes"
    authorizes.EntityData.BundleName = "cisco_ios_xr"
    authorizes.EntityData.ParentYangName = "network-attributes"
    authorizes.EntityData.SegmentPath = "authorizes"
    authorizes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/" + authorizes.EntityData.SegmentPath
    authorizes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizes.EntityData.Children = types.NewOrderedMap()
    authorizes.EntityData.Children.Append("authorize", types.YChild{"Authorize", nil})
    for i := range authorizes.Authorize {
        authorizes.EntityData.Children.Append(types.GetSegmentPath(authorizes.Authorize[i]), types.YChild{"Authorize", authorizes.Authorize[i]})
    }
    authorizes.EntityData.Leafs = types.NewOrderedMap()

    authorizes.EntityData.YListKeys = []string {}

    return &(authorizes.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize
// not authorize the network prefixes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ASCII string. The type is string with length:
    // 1..125.
    Name interface{}

    // Pool configs for this network.
    PoolAttributes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetEntityData() *types.CommonEntityData {
    authorize.EntityData.YFilter = authorize.YFilter
    authorize.EntityData.YangName = "authorize"
    authorize.EntityData.BundleName = "cisco_ios_xr"
    authorize.EntityData.ParentYangName = "authorizes"
    authorize.EntityData.SegmentPath = "authorize" + types.AddKeyToken(authorize.Name, "name")
    authorize.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/" + authorize.EntityData.SegmentPath
    authorize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorize.EntityData.Children = types.NewOrderedMap()
    authorize.EntityData.Children.Append("pool-attributes", types.YChild{"PoolAttributes", &authorize.PoolAttributes})
    authorize.EntityData.Leafs = types.NewOrderedMap()
    authorize.EntityData.Leafs.Append("name", types.YLeaf{"Name", authorize.Name})

    authorize.EntityData.YListKeys = []string {"Name"}

    return &(authorize.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes
// Pool configs for this network
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // pool configs for the mobile nodes.
    MobileNode MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode

    // pool configs for the mobile network.
    MobileNetwork MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetEntityData() *types.CommonEntityData {
    poolAttributes.EntityData.YFilter = poolAttributes.YFilter
    poolAttributes.EntityData.YangName = "pool-attributes"
    poolAttributes.EntityData.BundleName = "cisco_ios_xr"
    poolAttributes.EntityData.ParentYangName = "authorize"
    poolAttributes.EntityData.SegmentPath = "pool-attributes"
    poolAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/" + poolAttributes.EntityData.SegmentPath
    poolAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolAttributes.EntityData.Children = types.NewOrderedMap()
    poolAttributes.EntityData.Children.Append("mobile-node", types.YChild{"MobileNode", &poolAttributes.MobileNode})
    poolAttributes.EntityData.Children.Append("mobile-network", types.YChild{"MobileNetwork", &poolAttributes.MobileNetwork})
    poolAttributes.EntityData.Leafs = types.NewOrderedMap()

    poolAttributes.EntityData.YListKeys = []string {}

    return &(poolAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode
// pool configs for the mobile nodes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None.
    Ipv4Pool MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool

    // None.
    Ipv6Pool MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetEntityData() *types.CommonEntityData {
    mobileNode.EntityData.YFilter = mobileNode.YFilter
    mobileNode.EntityData.YangName = "mobile-node"
    mobileNode.EntityData.BundleName = "cisco_ios_xr"
    mobileNode.EntityData.ParentYangName = "pool-attributes"
    mobileNode.EntityData.SegmentPath = "mobile-node"
    mobileNode.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/" + mobileNode.EntityData.SegmentPath
    mobileNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobileNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobileNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobileNode.EntityData.Children = types.NewOrderedMap()
    mobileNode.EntityData.Children.Append("ipv4-pool", types.YChild{"Ipv4Pool", &mobileNode.Ipv4Pool})
    mobileNode.EntityData.Children.Append("ipv6-pool", types.YChild{"Ipv6Pool", &mobileNode.Ipv6Pool})
    mobileNode.EntityData.Leafs = types.NewOrderedMap()

    mobileNode.EntityData.YListKeys = []string {}

    return &(mobileNode.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool
// None
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool IPv4 start address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetEntityData() *types.CommonEntityData {
    ipv4Pool.EntityData.YFilter = ipv4Pool.YFilter
    ipv4Pool.EntityData.YangName = "ipv4-pool"
    ipv4Pool.EntityData.BundleName = "cisco_ios_xr"
    ipv4Pool.EntityData.ParentYangName = "mobile-node"
    ipv4Pool.EntityData.SegmentPath = "ipv4-pool"
    ipv4Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-node/" + ipv4Pool.EntityData.SegmentPath
    ipv4Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Pool.EntityData.Children = types.NewOrderedMap()
    ipv4Pool.EntityData.Leafs = types.NewOrderedMap()
    ipv4Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", ipv4Pool.StartAddress})
    ipv4Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", ipv4Pool.PoolPrefix})

    ipv4Pool.EntityData.YListKeys = []string {}

    return &(ipv4Pool.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool
// None
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool IPv6 start address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..62.
    PoolPrefix interface{}
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetEntityData() *types.CommonEntityData {
    ipv6Pool.EntityData.YFilter = ipv6Pool.YFilter
    ipv6Pool.EntityData.YangName = "ipv6-pool"
    ipv6Pool.EntityData.BundleName = "cisco_ios_xr"
    ipv6Pool.EntityData.ParentYangName = "mobile-node"
    ipv6Pool.EntityData.SegmentPath = "ipv6-pool"
    ipv6Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-node/" + ipv6Pool.EntityData.SegmentPath
    ipv6Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Pool.EntityData.Children = types.NewOrderedMap()
    ipv6Pool.EntityData.Leafs = types.NewOrderedMap()
    ipv6Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", ipv6Pool.StartAddress})
    ipv6Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", ipv6Pool.PoolPrefix})

    ipv6Pool.EntityData.YListKeys = []string {}

    return &(ipv6Pool.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork
// pool configs for the mobile network
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of MRIPV6Pool.
    Mripv6Pools MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools

    // Table of MRIPV4Pool.
    Mripv4Pools MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetEntityData() *types.CommonEntityData {
    mobileNetwork.EntityData.YFilter = mobileNetwork.YFilter
    mobileNetwork.EntityData.YangName = "mobile-network"
    mobileNetwork.EntityData.BundleName = "cisco_ios_xr"
    mobileNetwork.EntityData.ParentYangName = "pool-attributes"
    mobileNetwork.EntityData.SegmentPath = "mobile-network"
    mobileNetwork.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/" + mobileNetwork.EntityData.SegmentPath
    mobileNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobileNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobileNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobileNetwork.EntityData.Children = types.NewOrderedMap()
    mobileNetwork.EntityData.Children.Append("mripv6-pools", types.YChild{"Mripv6Pools", &mobileNetwork.Mripv6Pools})
    mobileNetwork.EntityData.Children.Append("mripv4-pools", types.YChild{"Mripv4Pools", &mobileNetwork.Mripv4Pools})
    mobileNetwork.EntityData.Leafs = types.NewOrderedMap()

    mobileNetwork.EntityData.YListKeys = []string {}

    return &(mobileNetwork.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools
// Table of MRIPV6Pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 pool. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool.
    Mripv6Pool []*MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetEntityData() *types.CommonEntityData {
    mripv6Pools.EntityData.YFilter = mripv6Pools.YFilter
    mripv6Pools.EntityData.YangName = "mripv6-pools"
    mripv6Pools.EntityData.BundleName = "cisco_ios_xr"
    mripv6Pools.EntityData.ParentYangName = "mobile-network"
    mripv6Pools.EntityData.SegmentPath = "mripv6-pools"
    mripv6Pools.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-network/" + mripv6Pools.EntityData.SegmentPath
    mripv6Pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv6Pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv6Pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv6Pools.EntityData.Children = types.NewOrderedMap()
    mripv6Pools.EntityData.Children.Append("mripv6-pool", types.YChild{"Mripv6Pool", nil})
    for i := range mripv6Pools.Mripv6Pool {
        mripv6Pools.EntityData.Children.Append(types.GetSegmentPath(mripv6Pools.Mripv6Pool[i]), types.YChild{"Mripv6Pool", mripv6Pools.Mripv6Pool[i]})
    }
    mripv6Pools.EntityData.Leafs = types.NewOrderedMap()

    mripv6Pools.EntityData.YListKeys = []string {}

    return &(mripv6Pools.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
// ipv6 pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Pool IPv6 start address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..64.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..64.
    NetworkPrefix interface{}
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetEntityData() *types.CommonEntityData {
    mripv6Pool.EntityData.YFilter = mripv6Pool.YFilter
    mripv6Pool.EntityData.YangName = "mripv6-pool"
    mripv6Pool.EntityData.BundleName = "cisco_ios_xr"
    mripv6Pool.EntityData.ParentYangName = "mripv6-pools"
    mripv6Pool.EntityData.SegmentPath = "mripv6-pool" + types.AddKeyToken(mripv6Pool.StartAddress, "start-address")
    mripv6Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-network/mripv6-pools/" + mripv6Pool.EntityData.SegmentPath
    mripv6Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv6Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv6Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv6Pool.EntityData.Children = types.NewOrderedMap()
    mripv6Pool.EntityData.Leafs = types.NewOrderedMap()
    mripv6Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", mripv6Pool.StartAddress})
    mripv6Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", mripv6Pool.PoolPrefix})
    mripv6Pool.EntityData.Leafs.Append("network-prefix", types.YLeaf{"NetworkPrefix", mripv6Pool.NetworkPrefix})

    mripv6Pool.EntityData.YListKeys = []string {"StartAddress"}

    return &(mripv6Pool.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools
// Table of MRIPV4Pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 pool. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool.
    Mripv4Pool []*MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetEntityData() *types.CommonEntityData {
    mripv4Pools.EntityData.YFilter = mripv4Pools.YFilter
    mripv4Pools.EntityData.YangName = "mripv4-pools"
    mripv4Pools.EntityData.BundleName = "cisco_ios_xr"
    mripv4Pools.EntityData.ParentYangName = "mobile-network"
    mripv4Pools.EntityData.SegmentPath = "mripv4-pools"
    mripv4Pools.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-network/" + mripv4Pools.EntityData.SegmentPath
    mripv4Pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv4Pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv4Pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv4Pools.EntityData.Children = types.NewOrderedMap()
    mripv4Pools.EntityData.Children.Append("mripv4-pool", types.YChild{"Mripv4Pool", nil})
    for i := range mripv4Pools.Mripv4Pool {
        mripv4Pools.EntityData.Children.Append(types.GetSegmentPath(mripv4Pools.Mripv4Pool[i]), types.YChild{"Mripv4Pool", mripv4Pools.Mripv4Pool[i]})
    }
    mripv4Pools.EntityData.Leafs = types.NewOrderedMap()

    mripv4Pools.EntityData.YListKeys = []string {}

    return &(mripv4Pools.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
// ipv4 pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Pool IPv4 start address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..32.
    NetworkPrefix interface{}
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetEntityData() *types.CommonEntityData {
    mripv4Pool.EntityData.YFilter = mripv4Pool.YFilter
    mripv4Pool.EntityData.YangName = "mripv4-pool"
    mripv4Pool.EntityData.BundleName = "cisco_ios_xr"
    mripv4Pool.EntityData.ParentYangName = "mripv4-pools"
    mripv4Pool.EntityData.SegmentPath = "mripv4-pool" + types.AddKeyToken(mripv4Pool.StartAddress, "start-address")
    mripv4Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/network-attributes/authorizes/authorize/pool-attributes/mobile-network/mripv4-pools/" + mripv4Pool.EntityData.SegmentPath
    mripv4Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv4Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv4Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv4Pool.EntityData.Children = types.NewOrderedMap()
    mripv4Pool.EntityData.Leafs = types.NewOrderedMap()
    mripv4Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", mripv4Pool.StartAddress})
    mripv4Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", mripv4Pool.PoolPrefix})
    mripv4Pool.EntityData.Leafs.Append("network-prefix", types.YLeaf{"NetworkPrefix", mripv4Pool.NetworkPrefix})

    mripv4Pool.EntityData.YListKeys = []string {"StartAddress"}

    return &(mripv4Pool.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey
// Customer specific GRE key
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GRE key type. The type is GreKeyType.
    GreKeyType interface{}

    // GRE key value. The type is interface{} with range: 1..4294967295.
    GreKeyValue interface{}
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetEntityData() *types.CommonEntityData {
    greKey.EntityData.YFilter = greKey.YFilter
    greKey.EntityData.YangName = "gre-key"
    greKey.EntityData.BundleName = "cisco_ios_xr"
    greKey.EntityData.ParentYangName = "customer"
    greKey.EntityData.SegmentPath = "gre-key"
    greKey.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + greKey.EntityData.SegmentPath
    greKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    greKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    greKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    greKey.EntityData.Children = types.NewOrderedMap()
    greKey.EntityData.Leafs = types.NewOrderedMap()
    greKey.EntityData.Leafs.Append("gre-key-type", types.YLeaf{"GreKeyType", greKey.GreKeyType})
    greKey.EntityData.Leafs.Append("gre-key-value", types.YLeaf{"GreKeyValue", greKey.GreKeyValue})

    greKey.EntityData.YListKeys = []string {}

    return &(greKey.EntityData)
}

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes
// Customer specific binding attributes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum bce lifetime permitted. The type is interface{} with range:
    // 10..65535. Units are second.
    MaxLifeTime interface{}
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetEntityData() *types.CommonEntityData {
    bindingAttributes.EntityData.YFilter = bindingAttributes.YFilter
    bindingAttributes.EntityData.YangName = "binding-attributes"
    bindingAttributes.EntityData.BundleName = "cisco_ios_xr"
    bindingAttributes.EntityData.ParentYangName = "customer"
    bindingAttributes.EntityData.SegmentPath = "binding-attributes"
    bindingAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/services/service/customers/customer/" + bindingAttributes.EntityData.SegmentPath
    bindingAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingAttributes.EntityData.Children = types.NewOrderedMap()
    bindingAttributes.EntityData.Leafs = types.NewOrderedMap()
    bindingAttributes.EntityData.Leafs.Append("max-life-time", types.YLeaf{"MaxLifeTime", bindingAttributes.MaxLifeTime})

    bindingAttributes.EntityData.YListKeys = []string {}

    return &(bindingAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Networks
// Table of Network
type MobileIp_Lmas_Lma_Networks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // network for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network.
    Network []*MobileIp_Lmas_Lma_Networks_Network
}

func (networks *MobileIp_Lmas_Lma_Networks) GetEntityData() *types.CommonEntityData {
    networks.EntityData.YFilter = networks.YFilter
    networks.EntityData.YangName = "networks"
    networks.EntityData.BundleName = "cisco_ios_xr"
    networks.EntityData.ParentYangName = "lma"
    networks.EntityData.SegmentPath = "networks"
    networks.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + networks.EntityData.SegmentPath
    networks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    networks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    networks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    networks.EntityData.Children = types.NewOrderedMap()
    networks.EntityData.Children.Append("network", types.YChild{"Network", nil})
    for i := range networks.Network {
        networks.EntityData.Children.Append(types.GetSegmentPath(networks.Network[i]), types.YChild{"Network", networks.Network[i]})
    }
    networks.EntityData.Leafs = types.NewOrderedMap()

    networks.EntityData.YListKeys = []string {}

    return &(networks.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network
// network for this LMA
type MobileIp_Lmas_Lma_Networks_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Network name. The type is string with length:
    // 1..125.
    LmaNetwork interface{}

    // Pool configs for this network.
    PoolAttributes MobileIp_Lmas_Lma_Networks_Network_PoolAttributes
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "networks"
    network.EntityData.SegmentPath = "network" + types.AddKeyToken(network.LmaNetwork, "lma-network")
    network.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/" + network.EntityData.SegmentPath
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Children.Append("pool-attributes", types.YChild{"PoolAttributes", &network.PoolAttributes})
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("lma-network", types.YLeaf{"LmaNetwork", network.LmaNetwork})

    network.EntityData.YListKeys = []string {"LmaNetwork"}

    return &(network.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes
// Pool configs for this network
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // pool configs for the mobile nodes.
    MobileNode MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode

    // pool configs for the mobile network.
    MobileNetwork MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetEntityData() *types.CommonEntityData {
    poolAttributes.EntityData.YFilter = poolAttributes.YFilter
    poolAttributes.EntityData.YangName = "pool-attributes"
    poolAttributes.EntityData.BundleName = "cisco_ios_xr"
    poolAttributes.EntityData.ParentYangName = "network"
    poolAttributes.EntityData.SegmentPath = "pool-attributes"
    poolAttributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/" + poolAttributes.EntityData.SegmentPath
    poolAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    poolAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    poolAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    poolAttributes.EntityData.Children = types.NewOrderedMap()
    poolAttributes.EntityData.Children.Append("mobile-node", types.YChild{"MobileNode", &poolAttributes.MobileNode})
    poolAttributes.EntityData.Children.Append("mobile-network", types.YChild{"MobileNetwork", &poolAttributes.MobileNetwork})
    poolAttributes.EntityData.Leafs = types.NewOrderedMap()

    poolAttributes.EntityData.YListKeys = []string {}

    return &(poolAttributes.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode
// pool configs for the mobile nodes
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None.
    Ipv4Pool MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool

    // None.
    Ipv6Pool MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetEntityData() *types.CommonEntityData {
    mobileNode.EntityData.YFilter = mobileNode.YFilter
    mobileNode.EntityData.YangName = "mobile-node"
    mobileNode.EntityData.BundleName = "cisco_ios_xr"
    mobileNode.EntityData.ParentYangName = "pool-attributes"
    mobileNode.EntityData.SegmentPath = "mobile-node"
    mobileNode.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/" + mobileNode.EntityData.SegmentPath
    mobileNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobileNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobileNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobileNode.EntityData.Children = types.NewOrderedMap()
    mobileNode.EntityData.Children.Append("ipv4-pool", types.YChild{"Ipv4Pool", &mobileNode.Ipv4Pool})
    mobileNode.EntityData.Children.Append("ipv6-pool", types.YChild{"Ipv6Pool", &mobileNode.Ipv6Pool})
    mobileNode.EntityData.Leafs = types.NewOrderedMap()

    mobileNode.EntityData.YListKeys = []string {}

    return &(mobileNode.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool
// None
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool IPv4 start address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetEntityData() *types.CommonEntityData {
    ipv4Pool.EntityData.YFilter = ipv4Pool.YFilter
    ipv4Pool.EntityData.YangName = "ipv4-pool"
    ipv4Pool.EntityData.BundleName = "cisco_ios_xr"
    ipv4Pool.EntityData.ParentYangName = "mobile-node"
    ipv4Pool.EntityData.SegmentPath = "ipv4-pool"
    ipv4Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-node/" + ipv4Pool.EntityData.SegmentPath
    ipv4Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Pool.EntityData.Children = types.NewOrderedMap()
    ipv4Pool.EntityData.Leafs = types.NewOrderedMap()
    ipv4Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", ipv4Pool.StartAddress})
    ipv4Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", ipv4Pool.PoolPrefix})

    ipv4Pool.EntityData.YListKeys = []string {}

    return &(ipv4Pool.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool
// None
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pool IPv6 start address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..62.
    PoolPrefix interface{}
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetEntityData() *types.CommonEntityData {
    ipv6Pool.EntityData.YFilter = ipv6Pool.YFilter
    ipv6Pool.EntityData.YangName = "ipv6-pool"
    ipv6Pool.EntityData.BundleName = "cisco_ios_xr"
    ipv6Pool.EntityData.ParentYangName = "mobile-node"
    ipv6Pool.EntityData.SegmentPath = "ipv6-pool"
    ipv6Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-node/" + ipv6Pool.EntityData.SegmentPath
    ipv6Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Pool.EntityData.Children = types.NewOrderedMap()
    ipv6Pool.EntityData.Leafs = types.NewOrderedMap()
    ipv6Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", ipv6Pool.StartAddress})
    ipv6Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", ipv6Pool.PoolPrefix})

    ipv6Pool.EntityData.YListKeys = []string {}

    return &(ipv6Pool.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork
// pool configs for the mobile network
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of MRIPV6Pool.
    Mripv6Pools MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools

    // Table of MRIPV4Pool.
    Mripv4Pools MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetEntityData() *types.CommonEntityData {
    mobileNetwork.EntityData.YFilter = mobileNetwork.YFilter
    mobileNetwork.EntityData.YangName = "mobile-network"
    mobileNetwork.EntityData.BundleName = "cisco_ios_xr"
    mobileNetwork.EntityData.ParentYangName = "pool-attributes"
    mobileNetwork.EntityData.SegmentPath = "mobile-network"
    mobileNetwork.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/" + mobileNetwork.EntityData.SegmentPath
    mobileNetwork.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobileNetwork.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobileNetwork.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobileNetwork.EntityData.Children = types.NewOrderedMap()
    mobileNetwork.EntityData.Children.Append("mripv6-pools", types.YChild{"Mripv6Pools", &mobileNetwork.Mripv6Pools})
    mobileNetwork.EntityData.Children.Append("mripv4-pools", types.YChild{"Mripv4Pools", &mobileNetwork.Mripv4Pools})
    mobileNetwork.EntityData.Leafs = types.NewOrderedMap()

    mobileNetwork.EntityData.YListKeys = []string {}

    return &(mobileNetwork.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools
// Table of MRIPV6Pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv6 pool. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool.
    Mripv6Pool []*MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetEntityData() *types.CommonEntityData {
    mripv6Pools.EntityData.YFilter = mripv6Pools.YFilter
    mripv6Pools.EntityData.YangName = "mripv6-pools"
    mripv6Pools.EntityData.BundleName = "cisco_ios_xr"
    mripv6Pools.EntityData.ParentYangName = "mobile-network"
    mripv6Pools.EntityData.SegmentPath = "mripv6-pools"
    mripv6Pools.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-network/" + mripv6Pools.EntityData.SegmentPath
    mripv6Pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv6Pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv6Pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv6Pools.EntityData.Children = types.NewOrderedMap()
    mripv6Pools.EntityData.Children.Append("mripv6-pool", types.YChild{"Mripv6Pool", nil})
    for i := range mripv6Pools.Mripv6Pool {
        mripv6Pools.EntityData.Children.Append(types.GetSegmentPath(mripv6Pools.Mripv6Pool[i]), types.YChild{"Mripv6Pool", mripv6Pools.Mripv6Pool[i]})
    }
    mripv6Pools.EntityData.Leafs = types.NewOrderedMap()

    mripv6Pools.EntityData.YListKeys = []string {}

    return &(mripv6Pools.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
// ipv6 pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Pool IPv6 start address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..64.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..64.
    NetworkPrefix interface{}
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetEntityData() *types.CommonEntityData {
    mripv6Pool.EntityData.YFilter = mripv6Pool.YFilter
    mripv6Pool.EntityData.YangName = "mripv6-pool"
    mripv6Pool.EntityData.BundleName = "cisco_ios_xr"
    mripv6Pool.EntityData.ParentYangName = "mripv6-pools"
    mripv6Pool.EntityData.SegmentPath = "mripv6-pool" + types.AddKeyToken(mripv6Pool.StartAddress, "start-address")
    mripv6Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-network/mripv6-pools/" + mripv6Pool.EntityData.SegmentPath
    mripv6Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv6Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv6Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv6Pool.EntityData.Children = types.NewOrderedMap()
    mripv6Pool.EntityData.Leafs = types.NewOrderedMap()
    mripv6Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", mripv6Pool.StartAddress})
    mripv6Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", mripv6Pool.PoolPrefix})
    mripv6Pool.EntityData.Leafs.Append("network-prefix", types.YLeaf{"NetworkPrefix", mripv6Pool.NetworkPrefix})

    mripv6Pool.EntityData.YListKeys = []string {"StartAddress"}

    return &(mripv6Pool.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools
// Table of MRIPV4Pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 pool. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool.
    Mripv4Pool []*MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetEntityData() *types.CommonEntityData {
    mripv4Pools.EntityData.YFilter = mripv4Pools.YFilter
    mripv4Pools.EntityData.YangName = "mripv4-pools"
    mripv4Pools.EntityData.BundleName = "cisco_ios_xr"
    mripv4Pools.EntityData.ParentYangName = "mobile-network"
    mripv4Pools.EntityData.SegmentPath = "mripv4-pools"
    mripv4Pools.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-network/" + mripv4Pools.EntityData.SegmentPath
    mripv4Pools.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv4Pools.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv4Pools.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv4Pools.EntityData.Children = types.NewOrderedMap()
    mripv4Pools.EntityData.Children.Append("mripv4-pool", types.YChild{"Mripv4Pool", nil})
    for i := range mripv4Pools.Mripv4Pool {
        mripv4Pools.EntityData.Children.Append(types.GetSegmentPath(mripv4Pools.Mripv4Pool[i]), types.YChild{"Mripv4Pool", mripv4Pools.Mripv4Pool[i]})
    }
    mripv4Pools.EntityData.Leafs = types.NewOrderedMap()

    mripv4Pools.EntityData.YListKeys = []string {}

    return &(mripv4Pools.EntityData)
}

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
// ipv4 pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Pool IPv4 start address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..32.
    NetworkPrefix interface{}
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetEntityData() *types.CommonEntityData {
    mripv4Pool.EntityData.YFilter = mripv4Pool.YFilter
    mripv4Pool.EntityData.YangName = "mripv4-pool"
    mripv4Pool.EntityData.BundleName = "cisco_ios_xr"
    mripv4Pool.EntityData.ParentYangName = "mripv4-pools"
    mripv4Pool.EntityData.SegmentPath = "mripv4-pool" + types.AddKeyToken(mripv4Pool.StartAddress, "start-address")
    mripv4Pool.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/networks/network/pool-attributes/mobile-network/mripv4-pools/" + mripv4Pool.EntityData.SegmentPath
    mripv4Pool.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mripv4Pool.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mripv4Pool.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mripv4Pool.EntityData.Children = types.NewOrderedMap()
    mripv4Pool.EntityData.Leafs = types.NewOrderedMap()
    mripv4Pool.EntityData.Leafs.Append("start-address", types.YLeaf{"StartAddress", mripv4Pool.StartAddress})
    mripv4Pool.EntityData.Leafs.Append("pool-prefix", types.YLeaf{"PoolPrefix", mripv4Pool.PoolPrefix})
    mripv4Pool.EntityData.Leafs.Append("network-prefix", types.YLeaf{"NetworkPrefix", mripv4Pool.NetworkPrefix})

    mripv4Pool.EntityData.YListKeys = []string {"StartAddress"}

    return &(mripv4Pool.EntityData)
}

// MobileIp_Lmas_Lma_ReplayProtection
// Replay Protection Method
type MobileIp_Lmas_Lma_ReplayProtection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify timestamp window value in seconds. The type is interface{} with
    // range: 1..255. Units are second.
    TimestampWindow interface{}
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetEntityData() *types.CommonEntityData {
    replayProtection.EntityData.YFilter = replayProtection.YFilter
    replayProtection.EntityData.YangName = "replay-protection"
    replayProtection.EntityData.BundleName = "cisco_ios_xr"
    replayProtection.EntityData.ParentYangName = "lma"
    replayProtection.EntityData.SegmentPath = "replay-protection"
    replayProtection.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip/lmas/lma/" + replayProtection.EntityData.SegmentPath
    replayProtection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replayProtection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replayProtection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replayProtection.EntityData.Children = types.NewOrderedMap()
    replayProtection.EntityData.Leafs = types.NewOrderedMap()
    replayProtection.EntityData.Leafs.Append("timestamp-window", types.YLeaf{"TimestampWindow", replayProtection.TimestampWindow})

    replayProtection.EntityData.YListKeys = []string {}

    return &(replayProtection.EntityData)
}

