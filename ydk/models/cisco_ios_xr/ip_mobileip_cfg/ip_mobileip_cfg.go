// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-mobileip package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mobile-ip: MobileIP configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Domain.
    Domains MobileIp_Domains

    // Table of LMA.
    Lmas MobileIp_Lmas
}

func (mobileIp *MobileIp) GetFilter() yfilter.YFilter { return mobileIp.YFilter }

func (mobileIp *MobileIp) SetFilter(yf yfilter.YFilter) { mobileIp.YFilter = yf }

func (mobileIp *MobileIp) GetGoName(yname string) string {
    if yname == "domains" { return "Domains" }
    if yname == "lmas" { return "Lmas" }
    return ""
}

func (mobileIp *MobileIp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-mobileip-cfg:mobile-ip"
}

func (mobileIp *MobileIp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "domains" {
        return &mobileIp.Domains
    }
    if childYangName == "lmas" {
        return &mobileIp.Lmas
    }
    return nil
}

func (mobileIp *MobileIp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["domains"] = &mobileIp.Domains
    children["lmas"] = &mobileIp.Lmas
    return children
}

func (mobileIp *MobileIp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobileIp *MobileIp) GetBundleName() string { return "cisco_ios_xr" }

func (mobileIp *MobileIp) GetYangName() string { return "mobile-ip" }

func (mobileIp *MobileIp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobileIp *MobileIp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobileIp *MobileIp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobileIp *MobileIp) SetParent(parent types.Entity) { mobileIp.parent = parent }

func (mobileIp *MobileIp) GetParent() types.Entity { return mobileIp.parent }

func (mobileIp *MobileIp) GetParentYangName() string { return "Cisco-IOS-XR-ip-mobileip-cfg" }

// MobileIp_Domains
// Table of Domain
type MobileIp_Domains struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PMIPv6 domain configuration. The type is slice of MobileIp_Domains_Domain.
    Domain []MobileIp_Domains_Domain
}

func (domains *MobileIp_Domains) GetFilter() yfilter.YFilter { return domains.YFilter }

func (domains *MobileIp_Domains) SetFilter(yf yfilter.YFilter) { domains.YFilter = yf }

func (domains *MobileIp_Domains) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    return ""
}

func (domains *MobileIp_Domains) GetSegmentPath() string {
    return "domains"
}

func (domains *MobileIp_Domains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "domain" {
        for _, c := range domains.Domain {
            if domains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Domains_Domain{}
        domains.Domain = append(domains.Domain, child)
        return &domains.Domain[len(domains.Domain)-1]
    }
    return nil
}

func (domains *MobileIp_Domains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range domains.Domain {
        children[domains.Domain[i].GetSegmentPath()] = &domains.Domain[i]
    }
    return children
}

func (domains *MobileIp_Domains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (domains *MobileIp_Domains) GetBundleName() string { return "cisco_ios_xr" }

func (domains *MobileIp_Domains) GetYangName() string { return "domains" }

func (domains *MobileIp_Domains) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domains *MobileIp_Domains) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domains *MobileIp_Domains) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domains *MobileIp_Domains) SetParent(parent types.Entity) { domains.parent = parent }

func (domains *MobileIp_Domains) GetParent() types.Entity { return domains.parent }

func (domains *MobileIp_Domains) GetParentYangName() string { return "mobile-ip" }

// MobileIp_Domains_Domain
// PMIPv6 domain configuration
type MobileIp_Domains_Domain struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (domain *MobileIp_Domains_Domain) GetFilter() yfilter.YFilter { return domain.YFilter }

func (domain *MobileIp_Domains_Domain) SetFilter(yf yfilter.YFilter) { domain.YFilter = yf }

func (domain *MobileIp_Domains_Domain) GetGoName(yname string) string {
    if yname == "domain-name" { return "DomainName" }
    if yname == "enable" { return "Enable" }
    if yname == "mags" { return "Mags" }
    if yname == "nais" { return "Nais" }
    if yname == "authenticate-option" { return "AuthenticateOption" }
    if yname == "lmas" { return "Lmas" }
    return ""
}

func (domain *MobileIp_Domains_Domain) GetSegmentPath() string {
    return "domain" + "[domain-name='" + fmt.Sprintf("%v", domain.DomainName) + "']"
}

func (domain *MobileIp_Domains_Domain) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mags" {
        return &domain.Mags
    }
    if childYangName == "nais" {
        return &domain.Nais
    }
    if childYangName == "authenticate-option" {
        return &domain.AuthenticateOption
    }
    if childYangName == "lmas" {
        return &domain.Lmas
    }
    return nil
}

func (domain *MobileIp_Domains_Domain) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mags"] = &domain.Mags
    children["nais"] = &domain.Nais
    children["authenticate-option"] = &domain.AuthenticateOption
    children["lmas"] = &domain.Lmas
    return children
}

func (domain *MobileIp_Domains_Domain) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-name"] = domain.DomainName
    leafs["enable"] = domain.Enable
    return leafs
}

func (domain *MobileIp_Domains_Domain) GetBundleName() string { return "cisco_ios_xr" }

func (domain *MobileIp_Domains_Domain) GetYangName() string { return "domain" }

func (domain *MobileIp_Domains_Domain) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (domain *MobileIp_Domains_Domain) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (domain *MobileIp_Domains_Domain) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (domain *MobileIp_Domains_Domain) SetParent(parent types.Entity) { domain.parent = parent }

func (domain *MobileIp_Domains_Domain) GetParent() types.Entity { return domain.parent }

func (domain *MobileIp_Domains_Domain) GetParentYangName() string { return "domains" }

// MobileIp_Domains_Domain_Mags
// Table of MAG
type MobileIp_Domains_Domain_Mags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAG within domain. The type is slice of MobileIp_Domains_Domain_Mags_Mag.
    Mag []MobileIp_Domains_Domain_Mags_Mag
}

func (mags *MobileIp_Domains_Domain_Mags) GetFilter() yfilter.YFilter { return mags.YFilter }

func (mags *MobileIp_Domains_Domain_Mags) SetFilter(yf yfilter.YFilter) { mags.YFilter = yf }

func (mags *MobileIp_Domains_Domain_Mags) GetGoName(yname string) string {
    if yname == "mag" { return "Mag" }
    return ""
}

func (mags *MobileIp_Domains_Domain_Mags) GetSegmentPath() string {
    return "mags"
}

func (mags *MobileIp_Domains_Domain_Mags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mag" {
        for _, c := range mags.Mag {
            if mags.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Domains_Domain_Mags_Mag{}
        mags.Mag = append(mags.Mag, child)
        return &mags.Mag[len(mags.Mag)-1]
    }
    return nil
}

func (mags *MobileIp_Domains_Domain_Mags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mags.Mag {
        children[mags.Mag[i].GetSegmentPath()] = &mags.Mag[i]
    }
    return children
}

func (mags *MobileIp_Domains_Domain_Mags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mags *MobileIp_Domains_Domain_Mags) GetBundleName() string { return "cisco_ios_xr" }

func (mags *MobileIp_Domains_Domain_Mags) GetYangName() string { return "mags" }

func (mags *MobileIp_Domains_Domain_Mags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mags *MobileIp_Domains_Domain_Mags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mags *MobileIp_Domains_Domain_Mags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mags *MobileIp_Domains_Domain_Mags) SetParent(parent types.Entity) { mags.parent = parent }

func (mags *MobileIp_Domains_Domain_Mags) GetParent() types.Entity { return mags.parent }

func (mags *MobileIp_Domains_Domain_Mags) GetParentYangName() string { return "domain" }

// MobileIp_Domains_Domain_Mags_Mag
// MAG within domain
type MobileIp_Domains_Domain_Mags_Mag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MAG Identifier. The type is string with length:
    // 1..125.
    MagName interface{}
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetFilter() yfilter.YFilter { return mag.YFilter }

func (mag *MobileIp_Domains_Domain_Mags_Mag) SetFilter(yf yfilter.YFilter) { mag.YFilter = yf }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetGoName(yname string) string {
    if yname == "mag-name" { return "MagName" }
    return ""
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetSegmentPath() string {
    return "mag" + "[mag-name='" + fmt.Sprintf("%v", mag.MagName) + "']"
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mag-name"] = mag.MagName
    return leafs
}

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetBundleName() string { return "cisco_ios_xr" }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetYangName() string { return "mag" }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mag *MobileIp_Domains_Domain_Mags_Mag) SetParent(parent types.Entity) { mag.parent = parent }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetParent() types.Entity { return mag.parent }

func (mag *MobileIp_Domains_Domain_Mags_Mag) GetParentYangName() string { return "mags" }

// MobileIp_Domains_Domain_Nais
// Table of NAI
type MobileIp_Domains_Domain_Nais struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Network access identifier or Realm. The type is slice of
    // MobileIp_Domains_Domain_Nais_Nai.
    Nai []MobileIp_Domains_Domain_Nais_Nai
}

func (nais *MobileIp_Domains_Domain_Nais) GetFilter() yfilter.YFilter { return nais.YFilter }

func (nais *MobileIp_Domains_Domain_Nais) SetFilter(yf yfilter.YFilter) { nais.YFilter = yf }

func (nais *MobileIp_Domains_Domain_Nais) GetGoName(yname string) string {
    if yname == "nai" { return "Nai" }
    return ""
}

func (nais *MobileIp_Domains_Domain_Nais) GetSegmentPath() string {
    return "nais"
}

func (nais *MobileIp_Domains_Domain_Nais) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nai" {
        for _, c := range nais.Nai {
            if nais.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Domains_Domain_Nais_Nai{}
        nais.Nai = append(nais.Nai, child)
        return &nais.Nai[len(nais.Nai)-1]
    }
    return nil
}

func (nais *MobileIp_Domains_Domain_Nais) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nais.Nai {
        children[nais.Nai[i].GetSegmentPath()] = &nais.Nai[i]
    }
    return children
}

func (nais *MobileIp_Domains_Domain_Nais) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nais *MobileIp_Domains_Domain_Nais) GetBundleName() string { return "cisco_ios_xr" }

func (nais *MobileIp_Domains_Domain_Nais) GetYangName() string { return "nais" }

func (nais *MobileIp_Domains_Domain_Nais) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nais *MobileIp_Domains_Domain_Nais) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nais *MobileIp_Domains_Domain_Nais) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nais *MobileIp_Domains_Domain_Nais) SetParent(parent types.Entity) { nais.parent = parent }

func (nais *MobileIp_Domains_Domain_Nais) GetParent() types.Entity { return nais.parent }

func (nais *MobileIp_Domains_Domain_Nais) GetParentYangName() string { return "domain" }

// MobileIp_Domains_Domain_Nais_Nai
// Network access identifier or Realm
type MobileIp_Domains_Domain_Nais_Nai struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetFilter() yfilter.YFilter { return nai.YFilter }

func (nai *MobileIp_Domains_Domain_Nais_Nai) SetFilter(yf yfilter.YFilter) { nai.YFilter = yf }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetGoName(yname string) string {
    if yname == "nai-name" { return "NaiName" }
    if yname == "lma" { return "Lma" }
    if yname == "apn" { return "Apn" }
    if yname == "customer" { return "Customer" }
    if yname == "service" { return "Service" }
    if yname == "network" { return "Network" }
    return ""
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetSegmentPath() string {
    return "nai" + "[nai-name='" + fmt.Sprintf("%v", nai.NaiName) + "']"
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nai-name"] = nai.NaiName
    leafs["lma"] = nai.Lma
    leafs["apn"] = nai.Apn
    leafs["customer"] = nai.Customer
    leafs["service"] = nai.Service
    leafs["network"] = nai.Network
    return leafs
}

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetBundleName() string { return "cisco_ios_xr" }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetYangName() string { return "nai" }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nai *MobileIp_Domains_Domain_Nais_Nai) SetParent(parent types.Entity) { nai.parent = parent }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetParent() types.Entity { return nai.parent }

func (nai *MobileIp_Domains_Domain_Nais_Nai) GetParentYangName() string { return "nais" }

// MobileIp_Domains_Domain_AuthenticateOption
// Authentication option between PMIPV6 entities
type MobileIp_Domains_Domain_AuthenticateOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetFilter() yfilter.YFilter { return authenticateOption.YFilter }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) SetFilter(yf yfilter.YFilter) { authenticateOption.YFilter = yf }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetGoName(yname string) string {
    if yname == "spi" { return "Spi" }
    if yname == "key" { return "Key" }
    return ""
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetSegmentPath() string {
    return "authenticate-option"
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spi"] = authenticateOption.Spi
    leafs["key"] = authenticateOption.Key
    return leafs
}

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetBundleName() string { return "cisco_ios_xr" }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetYangName() string { return "authenticate-option" }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) SetParent(parent types.Entity) { authenticateOption.parent = parent }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetParent() types.Entity { return authenticateOption.parent }

func (authenticateOption *MobileIp_Domains_Domain_AuthenticateOption) GetParentYangName() string { return "domain" }

// MobileIp_Domains_Domain_Lmas
// Table of LMA
type MobileIp_Domains_Domain_Lmas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LMA within domain. The type is slice of MobileIp_Domains_Domain_Lmas_Lma.
    Lma []MobileIp_Domains_Domain_Lmas_Lma
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetFilter() yfilter.YFilter { return lmas.YFilter }

func (lmas *MobileIp_Domains_Domain_Lmas) SetFilter(yf yfilter.YFilter) { lmas.YFilter = yf }

func (lmas *MobileIp_Domains_Domain_Lmas) GetGoName(yname string) string {
    if yname == "lma" { return "Lma" }
    return ""
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetSegmentPath() string {
    return "lmas"
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lma" {
        for _, c := range lmas.Lma {
            if lmas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Domains_Domain_Lmas_Lma{}
        lmas.Lma = append(lmas.Lma, child)
        return &lmas.Lma[len(lmas.Lma)-1]
    }
    return nil
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lmas.Lma {
        children[lmas.Lma[i].GetSegmentPath()] = &lmas.Lma[i]
    }
    return children
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lmas *MobileIp_Domains_Domain_Lmas) GetBundleName() string { return "cisco_ios_xr" }

func (lmas *MobileIp_Domains_Domain_Lmas) GetYangName() string { return "lmas" }

func (lmas *MobileIp_Domains_Domain_Lmas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmas *MobileIp_Domains_Domain_Lmas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmas *MobileIp_Domains_Domain_Lmas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmas *MobileIp_Domains_Domain_Lmas) SetParent(parent types.Entity) { lmas.parent = parent }

func (lmas *MobileIp_Domains_Domain_Lmas) GetParent() types.Entity { return lmas.parent }

func (lmas *MobileIp_Domains_Domain_Lmas) GetParentYangName() string { return "domain" }

// MobileIp_Domains_Domain_Lmas_Lma
// LMA within domain
type MobileIp_Domains_Domain_Lmas_Lma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LMA Identifier. The type is string with length:
    // 1..125.
    LmaName interface{}
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetFilter() yfilter.YFilter { return lma.YFilter }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) SetFilter(yf yfilter.YFilter) { lma.YFilter = yf }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetGoName(yname string) string {
    if yname == "lma-name" { return "LmaName" }
    return ""
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetSegmentPath() string {
    return "lma" + "[lma-name='" + fmt.Sprintf("%v", lma.LmaName) + "']"
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-name"] = lma.LmaName
    return leafs
}

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetBundleName() string { return "cisco_ios_xr" }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetYangName() string { return "lma" }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) SetParent(parent types.Entity) { lma.parent = parent }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetParent() types.Entity { return lma.parent }

func (lma *MobileIp_Domains_Domain_Lmas_Lma) GetParentYangName() string { return "lmas" }

// MobileIp_Lmas
// Table of LMA
type MobileIp_Lmas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PMIPv6 LMA configuration. The type is slice of MobileIp_Lmas_Lma.
    Lma []MobileIp_Lmas_Lma
}

func (lmas *MobileIp_Lmas) GetFilter() yfilter.YFilter { return lmas.YFilter }

func (lmas *MobileIp_Lmas) SetFilter(yf yfilter.YFilter) { lmas.YFilter = yf }

func (lmas *MobileIp_Lmas) GetGoName(yname string) string {
    if yname == "lma" { return "Lma" }
    return ""
}

func (lmas *MobileIp_Lmas) GetSegmentPath() string {
    return "lmas"
}

func (lmas *MobileIp_Lmas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lma" {
        for _, c := range lmas.Lma {
            if lmas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma{}
        lmas.Lma = append(lmas.Lma, child)
        return &lmas.Lma[len(lmas.Lma)-1]
    }
    return nil
}

func (lmas *MobileIp_Lmas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lmas.Lma {
        children[lmas.Lma[i].GetSegmentPath()] = &lmas.Lma[i]
    }
    return children
}

func (lmas *MobileIp_Lmas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lmas *MobileIp_Lmas) GetBundleName() string { return "cisco_ios_xr" }

func (lmas *MobileIp_Lmas) GetYangName() string { return "lmas" }

func (lmas *MobileIp_Lmas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmas *MobileIp_Lmas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmas *MobileIp_Lmas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmas *MobileIp_Lmas) SetParent(parent types.Entity) { lmas.parent = parent }

func (lmas *MobileIp_Lmas) GetParent() types.Entity { return lmas.parent }

func (lmas *MobileIp_Lmas) GetParentYangName() string { return "mobile-ip" }

// MobileIp_Lmas_Lma
// PMIPv6 LMA configuration
type MobileIp_Lmas_Lma struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

    // CN facing interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
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

func (lma *MobileIp_Lmas_Lma) GetFilter() yfilter.YFilter { return lma.YFilter }

func (lma *MobileIp_Lmas_Lma) SetFilter(yf yfilter.YFilter) { lma.YFilter = yf }

func (lma *MobileIp_Lmas_Lma) GetGoName(yname string) string {
    if yname == "lma-name" { return "LmaName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "generate" { return "Generate" }
    if yname == "mobile-route-ad" { return "MobileRouteAd" }
    if yname == "ani" { return "Ani" }
    if yname == "multipath" { return "Multipath" }
    if yname == "dynamic" { return "Dynamic" }
    if yname == "enforce" { return "Enforce" }
    if yname == "default-profile" { return "DefaultProfile" }
    if yname == "interface" { return "Interface" }
    if yname == "mobile-map" { return "MobileMap" }
    if yname == "pgw-subs-cont" { return "PgwSubsCont" }
    if yname == "binding-revocation-attributes" { return "BindingRevocationAttributes" }
    if yname == "rat-attributes" { return "RatAttributes" }
    if yname == "heart-beat-attributes" { return "HeartBeatAttributes" }
    if yname == "lmaipv6-addresses" { return "Lmaipv6Addresses" }
    if yname == "hnp" { return "Hnp" }
    if yname == "redistribute" { return "Redistribute" }
    if yname == "aaa" { return "Aaa" }
    if yname == "dscp" { return "Dscp" }
    if yname == "lmaipv4-addresses" { return "Lmaipv4Addresses" }
    if yname == "roles" { return "Roles" }
    if yname == "binding-attributes" { return "BindingAttributes" }
    if yname == "mags" { return "Mags" }
    if yname == "tunnel-attributes" { return "TunnelAttributes" }
    if yname == "services" { return "Services" }
    if yname == "networks" { return "Networks" }
    if yname == "replay-protection" { return "ReplayProtection" }
    return ""
}

func (lma *MobileIp_Lmas_Lma) GetSegmentPath() string {
    return "lma" + "[lma-name='" + fmt.Sprintf("%v", lma.LmaName) + "']" + "[domain-name='" + fmt.Sprintf("%v", lma.DomainName) + "']"
}

func (lma *MobileIp_Lmas_Lma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding-revocation-attributes" {
        return &lma.BindingRevocationAttributes
    }
    if childYangName == "rat-attributes" {
        return &lma.RatAttributes
    }
    if childYangName == "heart-beat-attributes" {
        return &lma.HeartBeatAttributes
    }
    if childYangName == "lmaipv6-addresses" {
        return &lma.Lmaipv6Addresses
    }
    if childYangName == "hnp" {
        return &lma.Hnp
    }
    if childYangName == "redistribute" {
        return &lma.Redistribute
    }
    if childYangName == "aaa" {
        return &lma.Aaa
    }
    if childYangName == "dscp" {
        return &lma.Dscp
    }
    if childYangName == "lmaipv4-addresses" {
        return &lma.Lmaipv4Addresses
    }
    if childYangName == "roles" {
        return &lma.Roles
    }
    if childYangName == "binding-attributes" {
        return &lma.BindingAttributes
    }
    if childYangName == "mags" {
        return &lma.Mags
    }
    if childYangName == "tunnel-attributes" {
        return &lma.TunnelAttributes
    }
    if childYangName == "services" {
        return &lma.Services
    }
    if childYangName == "networks" {
        return &lma.Networks
    }
    if childYangName == "replay-protection" {
        return &lma.ReplayProtection
    }
    return nil
}

func (lma *MobileIp_Lmas_Lma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["binding-revocation-attributes"] = &lma.BindingRevocationAttributes
    children["rat-attributes"] = &lma.RatAttributes
    children["heart-beat-attributes"] = &lma.HeartBeatAttributes
    children["lmaipv6-addresses"] = &lma.Lmaipv6Addresses
    children["hnp"] = &lma.Hnp
    children["redistribute"] = &lma.Redistribute
    children["aaa"] = &lma.Aaa
    children["dscp"] = &lma.Dscp
    children["lmaipv4-addresses"] = &lma.Lmaipv4Addresses
    children["roles"] = &lma.Roles
    children["binding-attributes"] = &lma.BindingAttributes
    children["mags"] = &lma.Mags
    children["tunnel-attributes"] = &lma.TunnelAttributes
    children["services"] = &lma.Services
    children["networks"] = &lma.Networks
    children["replay-protection"] = &lma.ReplayProtection
    return children
}

func (lma *MobileIp_Lmas_Lma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-name"] = lma.LmaName
    leafs["domain-name"] = lma.DomainName
    leafs["generate"] = lma.Generate
    leafs["mobile-route-ad"] = lma.MobileRouteAd
    leafs["ani"] = lma.Ani
    leafs["multipath"] = lma.Multipath
    leafs["dynamic"] = lma.Dynamic
    leafs["enforce"] = lma.Enforce
    leafs["default-profile"] = lma.DefaultProfile
    leafs["interface"] = lma.Interface
    leafs["mobile-map"] = lma.MobileMap
    leafs["pgw-subs-cont"] = lma.PgwSubsCont
    return leafs
}

func (lma *MobileIp_Lmas_Lma) GetBundleName() string { return "cisco_ios_xr" }

func (lma *MobileIp_Lmas_Lma) GetYangName() string { return "lma" }

func (lma *MobileIp_Lmas_Lma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lma *MobileIp_Lmas_Lma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lma *MobileIp_Lmas_Lma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lma *MobileIp_Lmas_Lma) SetParent(parent types.Entity) { lma.parent = parent }

func (lma *MobileIp_Lmas_Lma) GetParent() types.Entity { return lma.parent }

func (lma *MobileIp_Lmas_Lma) GetParentYangName() string { return "lmas" }

// MobileIp_Lmas_Lma_BindingRevocationAttributes
// LMA Binding Revocation Attributes
type MobileIp_Lmas_Lma_BindingRevocationAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of Retransmissons Allowed for BRI Message. The type is interface{}
    // with range: 1..10.
    Retry interface{}

    // Time to wait before Retransmitting BRI Message.
    Delay MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetFilter() yfilter.YFilter { return bindingRevocationAttributes.YFilter }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) SetFilter(yf yfilter.YFilter) { bindingRevocationAttributes.YFilter = yf }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetGoName(yname string) string {
    if yname == "retry" { return "Retry" }
    if yname == "delay" { return "Delay" }
    return ""
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetSegmentPath() string {
    return "binding-revocation-attributes"
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delay" {
        return &bindingRevocationAttributes.Delay
    }
    return nil
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delay"] = &bindingRevocationAttributes.Delay
    return children
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retry"] = bindingRevocationAttributes.Retry
    return leafs
}

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetYangName() string { return "binding-revocation-attributes" }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) SetParent(parent types.Entity) { bindingRevocationAttributes.parent = parent }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetParent() types.Entity { return bindingRevocationAttributes.parent }

func (bindingRevocationAttributes *MobileIp_Lmas_Lma_BindingRevocationAttributes) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay
// Time to wait before Retransmitting BRI
// Message
type MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify in millisec. The type is interface{} with range: 500..65535.
    BrMin interface{}

    // Specify in millisec. The type is interface{} with range: 500..65535.
    BrMax interface{}
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetGoName(yname string) string {
    if yname == "br-min" { return "BrMin" }
    if yname == "br-max" { return "BrMax" }
    return ""
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["br-min"] = delay.BrMin
    leafs["br-max"] = delay.BrMax
    return leafs
}

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetYangName() string { return "delay" }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetParent() types.Entity { return delay.parent }

func (delay *MobileIp_Lmas_Lma_BindingRevocationAttributes_Delay) GetParentYangName() string { return "binding-revocation-attributes" }

// MobileIp_Lmas_Lma_RatAttributes
// Radio access technology type config  this LMA
type MobileIp_Lmas_Lma_RatAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LMA rat type. The type is LmaRat.
    LmaRat interface{}

    // Specify the priority value. The type is interface{} with range: 1..255.
    PriorityValue interface{}
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetFilter() yfilter.YFilter { return ratAttributes.YFilter }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) SetFilter(yf yfilter.YFilter) { ratAttributes.YFilter = yf }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetGoName(yname string) string {
    if yname == "lma-rat" { return "LmaRat" }
    if yname == "priority-value" { return "PriorityValue" }
    return ""
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetSegmentPath() string {
    return "rat-attributes"
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-rat"] = ratAttributes.LmaRat
    leafs["priority-value"] = ratAttributes.PriorityValue
    return leafs
}

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetYangName() string { return "rat-attributes" }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) SetParent(parent types.Entity) { ratAttributes.parent = parent }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetParent() types.Entity { return ratAttributes.parent }

func (ratAttributes *MobileIp_Lmas_Lma_RatAttributes) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_HeartBeatAttributes
// heartbeat config for this LMA
type MobileIp_Lmas_Lma_HeartBeatAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the interval value in second. The type is interface{} with range:
    // 10..3600.
    Interval interface{}

    // Specify the retry value. The type is interface{} with range: 1..10.
    Retries interface{}

    // Specify the timeout value. The type is interface{} with range: 1..3600.
    Timeout interface{}
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetFilter() yfilter.YFilter { return heartBeatAttributes.YFilter }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) SetFilter(yf yfilter.YFilter) { heartBeatAttributes.YFilter = yf }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "retries" { return "Retries" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetSegmentPath() string {
    return "heart-beat-attributes"
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = heartBeatAttributes.Interval
    leafs["retries"] = heartBeatAttributes.Retries
    leafs["timeout"] = heartBeatAttributes.Timeout
    return leafs
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetYangName() string { return "heart-beat-attributes" }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) SetParent(parent types.Entity) { heartBeatAttributes.parent = parent }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetParent() types.Entity { return heartBeatAttributes.parent }

func (heartBeatAttributes *MobileIp_Lmas_Lma_HeartBeatAttributes) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Lmaipv6Addresses
// Table of LMAIPv6Address
type MobileIp_Lmas_Lma_Lmaipv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IPv6 address for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address.
    Lmaipv6Address []MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetFilter() yfilter.YFilter { return lmaipv6Addresses.YFilter }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) SetFilter(yf yfilter.YFilter) { lmaipv6Addresses.YFilter = yf }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetGoName(yname string) string {
    if yname == "lmaipv6-address" { return "Lmaipv6Address" }
    return ""
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetSegmentPath() string {
    return "lmaipv6-addresses"
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lmaipv6-address" {
        for _, c := range lmaipv6Addresses.Lmaipv6Address {
            if lmaipv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address{}
        lmaipv6Addresses.Lmaipv6Address = append(lmaipv6Addresses.Lmaipv6Address, child)
        return &lmaipv6Addresses.Lmaipv6Address[len(lmaipv6Addresses.Lmaipv6Address)-1]
    }
    return nil
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lmaipv6Addresses.Lmaipv6Address {
        children[lmaipv6Addresses.Lmaipv6Address[i].GetSegmentPath()] = &lmaipv6Addresses.Lmaipv6Address[i]
    }
    return children
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetYangName() string { return "lmaipv6-addresses" }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) SetParent(parent types.Entity) { lmaipv6Addresses.parent = parent }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetParent() types.Entity { return lmaipv6Addresses.parent }

func (lmaipv6Addresses *MobileIp_Lmas_Lma_Lmaipv6Addresses) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address
// Configure IPv6 address for this LMA
type MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LMA IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetFilter() yfilter.YFilter { return lmaipv6Address.YFilter }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) SetFilter(yf yfilter.YFilter) { lmaipv6Address.YFilter = yf }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetSegmentPath() string {
    return "lmaipv6-address" + "[address='" + fmt.Sprintf("%v", lmaipv6Address.Address) + "']"
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = lmaipv6Address.Address
    return leafs
}

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetYangName() string { return "lmaipv6-address" }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) SetParent(parent types.Entity) { lmaipv6Address.parent = parent }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetParent() types.Entity { return lmaipv6Address.parent }

func (lmaipv6Address *MobileIp_Lmas_Lma_Lmaipv6Addresses_Lmaipv6Address) GetParentYangName() string { return "lmaipv6-addresses" }

// MobileIp_Lmas_Lma_Hnp
// LMA HNP options
type MobileIp_Lmas_Lma_Hnp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // maximum HNPs allowed per MN interface. The type is interface{} with range:
    // 1..3.
    Maximum interface{}
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetFilter() yfilter.YFilter { return hnp.YFilter }

func (hnp *MobileIp_Lmas_Lma_Hnp) SetFilter(yf yfilter.YFilter) { hnp.YFilter = yf }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetGoName(yname string) string {
    if yname == "maximum" { return "Maximum" }
    return ""
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetSegmentPath() string {
    return "hnp"
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["maximum"] = hnp.Maximum
    return leafs
}

func (hnp *MobileIp_Lmas_Lma_Hnp) GetBundleName() string { return "cisco_ios_xr" }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetYangName() string { return "hnp" }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hnp *MobileIp_Lmas_Lma_Hnp) SetParent(parent types.Entity) { hnp.parent = parent }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetParent() types.Entity { return hnp.parent }

func (hnp *MobileIp_Lmas_Lma_Hnp) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Redistribute
// Redistribute routes
type MobileIp_Lmas_Lma_Redistribute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Redistribute type. The type is RedistType.
    RedistType interface{}

    // Redistribute sub-type. The type is RedistSubType.
    RedistSubType interface{}
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetFilter() yfilter.YFilter { return redistribute.YFilter }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) SetFilter(yf yfilter.YFilter) { redistribute.YFilter = yf }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetGoName(yname string) string {
    if yname == "redist-type" { return "RedistType" }
    if yname == "redist-sub-type" { return "RedistSubType" }
    return ""
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetSegmentPath() string {
    return "redistribute"
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["redist-type"] = redistribute.RedistType
    leafs["redist-sub-type"] = redistribute.RedistSubType
    return leafs
}

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetBundleName() string { return "cisco_ios_xr" }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetYangName() string { return "redistribute" }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) SetParent(parent types.Entity) { redistribute.parent = parent }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetParent() types.Entity { return redistribute.parent }

func (redistribute *MobileIp_Lmas_Lma_Redistribute) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Aaa
// AAA configuration for this LMA
type MobileIp_Lmas_Lma_Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA accounting for this LMA.
    Accounting MobileIp_Lmas_Lma_Aaa_Accounting
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *MobileIp_Lmas_Lma_Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetGoName(yname string) string {
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetSegmentPath() string {
    return "aaa"
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accounting" {
        return &aaa.Accounting
    }
    return nil
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accounting"] = &aaa.Accounting
    return children
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaa *MobileIp_Lmas_Lma_Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetYangName() string { return "aaa" }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *MobileIp_Lmas_Lma_Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *MobileIp_Lmas_Lma_Aaa) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Aaa_Accounting
// AAA accounting for this LMA
type MobileIp_Lmas_Lma_Aaa_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set constant integer. The type is interface{}.
    Enable interface{}

    // Interim acounting interval (in minutes). The type is interface{} with
    // range: 1..86400. Units are minute.
    InterimInterval interface{}
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interim-interval" { return "InterimInterval" }
    return ""
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = accounting.Enable
    leafs["interim-interval"] = accounting.InterimInterval
    return leafs
}

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetYangName() string { return "accounting" }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *MobileIp_Lmas_Lma_Aaa_Accounting) GetParentYangName() string { return "aaa" }

// MobileIp_Lmas_Lma_Dscp
// DSCP for packets originating from this LMA
type MobileIp_Lmas_Lma_Dscp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 1..63.
    Value interface{}

    // Set constant integer. The type is interface{}.
    Force interface{}
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetFilter() yfilter.YFilter { return dscp.YFilter }

func (dscp *MobileIp_Lmas_Lma_Dscp) SetFilter(yf yfilter.YFilter) { dscp.YFilter = yf }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "force" { return "Force" }
    return ""
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetSegmentPath() string {
    return "dscp"
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dscp.Value
    leafs["force"] = dscp.Force
    return leafs
}

func (dscp *MobileIp_Lmas_Lma_Dscp) GetBundleName() string { return "cisco_ios_xr" }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetYangName() string { return "dscp" }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscp *MobileIp_Lmas_Lma_Dscp) SetParent(parent types.Entity) { dscp.parent = parent }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetParent() types.Entity { return dscp.parent }

func (dscp *MobileIp_Lmas_Lma_Dscp) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Lmaipv4Addresses
// Table of LMAIPv4Address
type MobileIp_Lmas_Lma_Lmaipv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IPv4 address for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address.
    Lmaipv4Address []MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetFilter() yfilter.YFilter { return lmaipv4Addresses.YFilter }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) SetFilter(yf yfilter.YFilter) { lmaipv4Addresses.YFilter = yf }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetGoName(yname string) string {
    if yname == "lmaipv4-address" { return "Lmaipv4Address" }
    return ""
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetSegmentPath() string {
    return "lmaipv4-addresses"
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lmaipv4-address" {
        for _, c := range lmaipv4Addresses.Lmaipv4Address {
            if lmaipv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address{}
        lmaipv4Addresses.Lmaipv4Address = append(lmaipv4Addresses.Lmaipv4Address, child)
        return &lmaipv4Addresses.Lmaipv4Address[len(lmaipv4Addresses.Lmaipv4Address)-1]
    }
    return nil
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lmaipv4Addresses.Lmaipv4Address {
        children[lmaipv4Addresses.Lmaipv4Address[i].GetSegmentPath()] = &lmaipv4Addresses.Lmaipv4Address[i]
    }
    return children
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetYangName() string { return "lmaipv4-addresses" }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) SetParent(parent types.Entity) { lmaipv4Addresses.parent = parent }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetParent() types.Entity { return lmaipv4Addresses.parent }

func (lmaipv4Addresses *MobileIp_Lmas_Lma_Lmaipv4Addresses) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address
// Configure IPv4 address for this LMA
type MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LMA IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetFilter() yfilter.YFilter { return lmaipv4Address.YFilter }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) SetFilter(yf yfilter.YFilter) { lmaipv4Address.YFilter = yf }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetSegmentPath() string {
    return "lmaipv4-address" + "[address='" + fmt.Sprintf("%v", lmaipv4Address.Address) + "']"
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = lmaipv4Address.Address
    return leafs
}

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetYangName() string { return "lmaipv4-address" }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) SetParent(parent types.Entity) { lmaipv4Address.parent = parent }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetParent() types.Entity { return lmaipv4Address.parent }

func (lmaipv4Address *MobileIp_Lmas_Lma_Lmaipv4Addresses_Lmaipv4Address) GetParentYangName() string { return "lmaipv4-addresses" }

// MobileIp_Lmas_Lma_Roles
// Table of Role
type MobileIp_Lmas_Lma_Roles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Role of this LMA. The type is slice of MobileIp_Lmas_Lma_Roles_Role.
    Role []MobileIp_Lmas_Lma_Roles_Role
}

func (roles *MobileIp_Lmas_Lma_Roles) GetFilter() yfilter.YFilter { return roles.YFilter }

func (roles *MobileIp_Lmas_Lma_Roles) SetFilter(yf yfilter.YFilter) { roles.YFilter = yf }

func (roles *MobileIp_Lmas_Lma_Roles) GetGoName(yname string) string {
    if yname == "role" { return "Role" }
    return ""
}

func (roles *MobileIp_Lmas_Lma_Roles) GetSegmentPath() string {
    return "roles"
}

func (roles *MobileIp_Lmas_Lma_Roles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "role" {
        for _, c := range roles.Role {
            if roles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Roles_Role{}
        roles.Role = append(roles.Role, child)
        return &roles.Role[len(roles.Role)-1]
    }
    return nil
}

func (roles *MobileIp_Lmas_Lma_Roles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range roles.Role {
        children[roles.Role[i].GetSegmentPath()] = &roles.Role[i]
    }
    return children
}

func (roles *MobileIp_Lmas_Lma_Roles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (roles *MobileIp_Lmas_Lma_Roles) GetBundleName() string { return "cisco_ios_xr" }

func (roles *MobileIp_Lmas_Lma_Roles) GetYangName() string { return "roles" }

func (roles *MobileIp_Lmas_Lma_Roles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (roles *MobileIp_Lmas_Lma_Roles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (roles *MobileIp_Lmas_Lma_Roles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (roles *MobileIp_Lmas_Lma_Roles) SetParent(parent types.Entity) { roles.parent = parent }

func (roles *MobileIp_Lmas_Lma_Roles) GetParent() types.Entity { return roles.parent }

func (roles *MobileIp_Lmas_Lma_Roles) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Roles_Role
// Role of this LMA
type MobileIp_Lmas_Lma_Roles_Role struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LMA role mode. The type is LmaRole.
    LmaRole interface{}
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetFilter() yfilter.YFilter { return role.YFilter }

func (role *MobileIp_Lmas_Lma_Roles_Role) SetFilter(yf yfilter.YFilter) { role.YFilter = yf }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetGoName(yname string) string {
    if yname == "lma-role" { return "LmaRole" }
    return ""
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetSegmentPath() string {
    return "role" + "[lma-role='" + fmt.Sprintf("%v", role.LmaRole) + "']"
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-role"] = role.LmaRole
    return leafs
}

func (role *MobileIp_Lmas_Lma_Roles_Role) GetBundleName() string { return "cisco_ios_xr" }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetYangName() string { return "role" }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (role *MobileIp_Lmas_Lma_Roles_Role) SetParent(parent types.Entity) { role.parent = parent }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetParent() types.Entity { return role.parent }

func (role *MobileIp_Lmas_Lma_Roles_Role) GetParentYangName() string { return "roles" }

// MobileIp_Lmas_Lma_BindingAttributes
// LMA binding attributes
type MobileIp_Lmas_Lma_BindingAttributes struct {
    parent types.Entity
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

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetFilter() yfilter.YFilter { return bindingAttributes.YFilter }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) SetFilter(yf yfilter.YFilter) { bindingAttributes.YFilter = yf }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetGoName(yname string) string {
    if yname == "refresh-time" { return "RefreshTime" }
    if yname == "delete-wait-interval" { return "DeleteWaitInterval" }
    if yname == "create-wait-interval" { return "CreateWaitInterval" }
    if yname == "max-life-time" { return "MaxLifeTime" }
    if yname == "maximum" { return "Maximum" }
    return ""
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetSegmentPath() string {
    return "binding-attributes"
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["refresh-time"] = bindingAttributes.RefreshTime
    leafs["delete-wait-interval"] = bindingAttributes.DeleteWaitInterval
    leafs["create-wait-interval"] = bindingAttributes.CreateWaitInterval
    leafs["max-life-time"] = bindingAttributes.MaxLifeTime
    leafs["maximum"] = bindingAttributes.Maximum
    return leafs
}

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetYangName() string { return "binding-attributes" }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) SetParent(parent types.Entity) { bindingAttributes.parent = parent }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetParent() types.Entity { return bindingAttributes.parent }

func (bindingAttributes *MobileIp_Lmas_Lma_BindingAttributes) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Mags
// Table of MAG
type MobileIp_Lmas_Lma_Mags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAG within LMA. The type is slice of MobileIp_Lmas_Lma_Mags_Mag.
    Mag []MobileIp_Lmas_Lma_Mags_Mag
}

func (mags *MobileIp_Lmas_Lma_Mags) GetFilter() yfilter.YFilter { return mags.YFilter }

func (mags *MobileIp_Lmas_Lma_Mags) SetFilter(yf yfilter.YFilter) { mags.YFilter = yf }

func (mags *MobileIp_Lmas_Lma_Mags) GetGoName(yname string) string {
    if yname == "mag" { return "Mag" }
    return ""
}

func (mags *MobileIp_Lmas_Lma_Mags) GetSegmentPath() string {
    return "mags"
}

func (mags *MobileIp_Lmas_Lma_Mags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mag" {
        for _, c := range mags.Mag {
            if mags.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Mags_Mag{}
        mags.Mag = append(mags.Mag, child)
        return &mags.Mag[len(mags.Mag)-1]
    }
    return nil
}

func (mags *MobileIp_Lmas_Lma_Mags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mags.Mag {
        children[mags.Mag[i].GetSegmentPath()] = &mags.Mag[i]
    }
    return children
}

func (mags *MobileIp_Lmas_Lma_Mags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mags *MobileIp_Lmas_Lma_Mags) GetBundleName() string { return "cisco_ios_xr" }

func (mags *MobileIp_Lmas_Lma_Mags) GetYangName() string { return "mags" }

func (mags *MobileIp_Lmas_Lma_Mags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mags *MobileIp_Lmas_Lma_Mags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mags *MobileIp_Lmas_Lma_Mags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mags *MobileIp_Lmas_Lma_Mags) SetParent(parent types.Entity) { mags.parent = parent }

func (mags *MobileIp_Lmas_Lma_Mags) GetParent() types.Entity { return mags.parent }

func (mags *MobileIp_Lmas_Lma_Mags) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Mags_Mag
// MAG within LMA
type MobileIp_Lmas_Lma_Mags_Mag struct {
    parent types.Entity
    YFilter yfilter.YFilter

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
    // [a-zA-Z0-9./-]+.
    Tunnel interface{}

    // Authentication option between PMIPV6 entities.
    AuthenticateOption MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption

    // DSCP for packets originating from this MAG.
    Dscp MobileIp_Lmas_Lma_Mags_Mag_Dscp
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetFilter() yfilter.YFilter { return mag.YFilter }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) SetFilter(yf yfilter.YFilter) { mag.YFilter = yf }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetGoName(yname string) string {
    if yname == "mag-name" { return "MagName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "encap-option" { return "EncapOption" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "authenticate-option" { return "AuthenticateOption" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetSegmentPath() string {
    return "mag" + "[mag-name='" + fmt.Sprintf("%v", mag.MagName) + "']" + "[domain-name='" + fmt.Sprintf("%v", mag.DomainName) + "']"
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authenticate-option" {
        return &mag.AuthenticateOption
    }
    if childYangName == "dscp" {
        return &mag.Dscp
    }
    return nil
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authenticate-option"] = &mag.AuthenticateOption
    children["dscp"] = &mag.Dscp
    return children
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mag-name"] = mag.MagName
    leafs["domain-name"] = mag.DomainName
    leafs["encap-option"] = mag.EncapOption
    leafs["ipv4-address"] = mag.Ipv4Address
    leafs["ipv6-address"] = mag.Ipv6Address
    leafs["tunnel"] = mag.Tunnel
    return leafs
}

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetBundleName() string { return "cisco_ios_xr" }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetYangName() string { return "mag" }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) SetParent(parent types.Entity) { mag.parent = parent }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetParent() types.Entity { return mag.parent }

func (mag *MobileIp_Lmas_Lma_Mags_Mag) GetParentYangName() string { return "mags" }

// MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption
// Authentication option between PMIPV6
// entities
type MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetFilter() yfilter.YFilter { return authenticateOption.YFilter }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) SetFilter(yf yfilter.YFilter) { authenticateOption.YFilter = yf }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetGoName(yname string) string {
    if yname == "spi" { return "Spi" }
    if yname == "key" { return "Key" }
    return ""
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetSegmentPath() string {
    return "authenticate-option"
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spi"] = authenticateOption.Spi
    leafs["key"] = authenticateOption.Key
    return leafs
}

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetBundleName() string { return "cisco_ios_xr" }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetYangName() string { return "authenticate-option" }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) SetParent(parent types.Entity) { authenticateOption.parent = parent }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetParent() types.Entity { return authenticateOption.parent }

func (authenticateOption *MobileIp_Lmas_Lma_Mags_Mag_AuthenticateOption) GetParentYangName() string { return "mag" }

// MobileIp_Lmas_Lma_Mags_Mag_Dscp
// DSCP for packets originating from this MAG
type MobileIp_Lmas_Lma_Mags_Mag_Dscp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 1..63.
    Value interface{}

    // Set constant integer. The type is interface{}.
    Force interface{}
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetFilter() yfilter.YFilter { return dscp.YFilter }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) SetFilter(yf yfilter.YFilter) { dscp.YFilter = yf }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "force" { return "Force" }
    return ""
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetSegmentPath() string {
    return "dscp"
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = dscp.Value
    leafs["force"] = dscp.Force
    return leafs
}

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetBundleName() string { return "cisco_ios_xr" }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetYangName() string { return "dscp" }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) SetParent(parent types.Entity) { dscp.parent = parent }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetParent() types.Entity { return dscp.parent }

func (dscp *MobileIp_Lmas_Lma_Mags_Mag_Dscp) GetParentYangName() string { return "mag" }

// MobileIp_Lmas_Lma_TunnelAttributes
// tunnel attributes
type MobileIp_Lmas_Lma_TunnelAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // maximum transmission unit for tunnel. The type is interface{} with range:
    // 68..17916.
    Mtu interface{}

    // access list to apply for tunnel. The type is string with length: 1..125.
    Acl interface{}
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetFilter() yfilter.YFilter { return tunnelAttributes.YFilter }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) SetFilter(yf yfilter.YFilter) { tunnelAttributes.YFilter = yf }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetGoName(yname string) string {
    if yname == "mtu" { return "Mtu" }
    if yname == "acl" { return "Acl" }
    return ""
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetSegmentPath() string {
    return "tunnel-attributes"
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mtu"] = tunnelAttributes.Mtu
    leafs["acl"] = tunnelAttributes.Acl
    return leafs
}

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetYangName() string { return "tunnel-attributes" }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) SetParent(parent types.Entity) { tunnelAttributes.parent = parent }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetParent() types.Entity { return tunnelAttributes.parent }

func (tunnelAttributes *MobileIp_Lmas_Lma_TunnelAttributes) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Services
// Table of Service
type MobileIp_Lmas_Lma_Services struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Service of this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service.
    Service []MobileIp_Lmas_Lma_Services_Service
}

func (services *MobileIp_Lmas_Lma_Services) GetFilter() yfilter.YFilter { return services.YFilter }

func (services *MobileIp_Lmas_Lma_Services) SetFilter(yf yfilter.YFilter) { services.YFilter = yf }

func (services *MobileIp_Lmas_Lma_Services) GetGoName(yname string) string {
    if yname == "service" { return "Service" }
    return ""
}

func (services *MobileIp_Lmas_Lma_Services) GetSegmentPath() string {
    return "services"
}

func (services *MobileIp_Lmas_Lma_Services) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "service" {
        for _, c := range services.Service {
            if services.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service{}
        services.Service = append(services.Service, child)
        return &services.Service[len(services.Service)-1]
    }
    return nil
}

func (services *MobileIp_Lmas_Lma_Services) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range services.Service {
        children[services.Service[i].GetSegmentPath()] = &services.Service[i]
    }
    return children
}

func (services *MobileIp_Lmas_Lma_Services) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (services *MobileIp_Lmas_Lma_Services) GetBundleName() string { return "cisco_ios_xr" }

func (services *MobileIp_Lmas_Lma_Services) GetYangName() string { return "services" }

func (services *MobileIp_Lmas_Lma_Services) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (services *MobileIp_Lmas_Lma_Services) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (services *MobileIp_Lmas_Lma_Services) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (services *MobileIp_Lmas_Lma_Services) SetParent(parent types.Entity) { services.parent = parent }

func (services *MobileIp_Lmas_Lma_Services) GetParent() types.Entity { return services.parent }

func (services *MobileIp_Lmas_Lma_Services) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Services_Service
// Service of this LMA
type MobileIp_Lmas_Lma_Services_Service struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (service *MobileIp_Lmas_Lma_Services_Service) GetFilter() yfilter.YFilter { return service.YFilter }

func (service *MobileIp_Lmas_Lma_Services_Service) SetFilter(yf yfilter.YFilter) { service.YFilter = yf }

func (service *MobileIp_Lmas_Lma_Services_Service) GetGoName(yname string) string {
    if yname == "lma-service" { return "LmaService" }
    if yname == "mnp-customer" { return "MnpCustomer" }
    if yname == "mnp-ipv4-lmn" { return "MnpIpv4Lmn" }
    if yname == "mnp-ipv6-lmn" { return "MnpIpv6Lmn" }
    if yname == "mnp-lmn" { return "MnpLmn" }
    if yname == "ignore-home-address" { return "IgnoreHomeAddress" }
    if yname == "mnp-ipv4-customer" { return "MnpIpv4Customer" }
    if yname == "mnp-ipv6-customer" { return "MnpIpv6Customer" }
    if yname == "customers" { return "Customers" }
    return ""
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetSegmentPath() string {
    return "service" + "[lma-service='" + fmt.Sprintf("%v", service.LmaService) + "']"
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customers" {
        return &service.Customers
    }
    return nil
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["customers"] = &service.Customers
    return children
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-service"] = service.LmaService
    leafs["mnp-customer"] = service.MnpCustomer
    leafs["mnp-ipv4-lmn"] = service.MnpIpv4Lmn
    leafs["mnp-ipv6-lmn"] = service.MnpIpv6Lmn
    leafs["mnp-lmn"] = service.MnpLmn
    leafs["ignore-home-address"] = service.IgnoreHomeAddress
    leafs["mnp-ipv4-customer"] = service.MnpIpv4Customer
    leafs["mnp-ipv6-customer"] = service.MnpIpv6Customer
    return leafs
}

func (service *MobileIp_Lmas_Lma_Services_Service) GetBundleName() string { return "cisco_ios_xr" }

func (service *MobileIp_Lmas_Lma_Services_Service) GetYangName() string { return "service" }

func (service *MobileIp_Lmas_Lma_Services_Service) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (service *MobileIp_Lmas_Lma_Services_Service) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (service *MobileIp_Lmas_Lma_Services_Service) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (service *MobileIp_Lmas_Lma_Services_Service) SetParent(parent types.Entity) { service.parent = parent }

func (service *MobileIp_Lmas_Lma_Services_Service) GetParent() types.Entity { return service.parent }

func (service *MobileIp_Lmas_Lma_Services_Service) GetParentYangName() string { return "services" }

// MobileIp_Lmas_Lma_Services_Service_Customers
// Table of Customer
type MobileIp_Lmas_Lma_Services_Service_Customers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // customer configuration on this mobile local loop service. The type is slice
    // of MobileIp_Lmas_Lma_Services_Service_Customers_Customer.
    Customer []MobileIp_Lmas_Lma_Services_Service_Customers_Customer
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetFilter() yfilter.YFilter { return customers.YFilter }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) SetFilter(yf yfilter.YFilter) { customers.YFilter = yf }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetGoName(yname string) string {
    if yname == "customer" { return "Customer" }
    return ""
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetSegmentPath() string {
    return "customers"
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customer" {
        for _, c := range customers.Customer {
            if customers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service_Customers_Customer{}
        customers.Customer = append(customers.Customer, child)
        return &customers.Customer[len(customers.Customer)-1]
    }
    return nil
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range customers.Customer {
        children[customers.Customer[i].GetSegmentPath()] = &customers.Customer[i]
    }
    return children
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetBundleName() string { return "cisco_ios_xr" }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetYangName() string { return "customers" }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) SetParent(parent types.Entity) { customers.parent = parent }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetParent() types.Entity { return customers.parent }

func (customers *MobileIp_Lmas_Lma_Services_Service_Customers) GetParentYangName() string { return "service" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer
// customer configuration on this mobile local
// loop service
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetFilter() yfilter.YFilter { return customer.YFilter }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) SetFilter(yf yfilter.YFilter) { customer.YFilter = yf }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetGoName(yname string) string {
    if yname == "customer-name" { return "CustomerName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "mnp-customer" { return "MnpCustomer" }
    if yname == "mnp-ipv4-lmn" { return "MnpIpv4Lmn" }
    if yname == "mnp-ipv6-lmn" { return "MnpIpv6Lmn" }
    if yname == "mnp-lmn" { return "MnpLmn" }
    if yname == "mnp-ipv4-customer" { return "MnpIpv4Customer" }
    if yname == "mnp-ipv6-customer" { return "MnpIpv6Customer" }
    if yname == "mobile-route-ad" { return "MobileRouteAd" }
    if yname == "bandwidth-aggregate" { return "BandwidthAggregate" }
    if yname == "authenticate-option" { return "AuthenticateOption" }
    if yname == "heart-beat-attributes" { return "HeartBeatAttributes" }
    if yname == "transports" { return "Transports" }
    if yname == "network-attributes" { return "NetworkAttributes" }
    if yname == "gre-key" { return "GreKey" }
    if yname == "binding-attributes" { return "BindingAttributes" }
    return ""
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetSegmentPath() string {
    return "customer" + "[customer-name='" + fmt.Sprintf("%v", customer.CustomerName) + "']" + "[vrf-name='" + fmt.Sprintf("%v", customer.VrfName) + "']"
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authenticate-option" {
        return &customer.AuthenticateOption
    }
    if childYangName == "heart-beat-attributes" {
        return &customer.HeartBeatAttributes
    }
    if childYangName == "transports" {
        return &customer.Transports
    }
    if childYangName == "network-attributes" {
        return &customer.NetworkAttributes
    }
    if childYangName == "gre-key" {
        return &customer.GreKey
    }
    if childYangName == "binding-attributes" {
        return &customer.BindingAttributes
    }
    return nil
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authenticate-option"] = &customer.AuthenticateOption
    children["heart-beat-attributes"] = &customer.HeartBeatAttributes
    children["transports"] = &customer.Transports
    children["network-attributes"] = &customer.NetworkAttributes
    children["gre-key"] = &customer.GreKey
    children["binding-attributes"] = &customer.BindingAttributes
    return children
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["customer-name"] = customer.CustomerName
    leafs["vrf-name"] = customer.VrfName
    leafs["mnp-customer"] = customer.MnpCustomer
    leafs["mnp-ipv4-lmn"] = customer.MnpIpv4Lmn
    leafs["mnp-ipv6-lmn"] = customer.MnpIpv6Lmn
    leafs["mnp-lmn"] = customer.MnpLmn
    leafs["mnp-ipv4-customer"] = customer.MnpIpv4Customer
    leafs["mnp-ipv6-customer"] = customer.MnpIpv6Customer
    leafs["mobile-route-ad"] = customer.MobileRouteAd
    leafs["bandwidth-aggregate"] = customer.BandwidthAggregate
    return leafs
}

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetBundleName() string { return "cisco_ios_xr" }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetYangName() string { return "customer" }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) SetParent(parent types.Entity) { customer.parent = parent }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetParent() types.Entity { return customer.parent }

func (customer *MobileIp_Lmas_Lma_Services_Service_Customers_Customer) GetParentYangName() string { return "customers" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption
// Authentication option between PMIPV6
// entities
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SPI in hex value. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Spi interface{}

    // ASCII string. The type is string with length: 1..125.
    Key interface{}
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetFilter() yfilter.YFilter { return authenticateOption.YFilter }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) SetFilter(yf yfilter.YFilter) { authenticateOption.YFilter = yf }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetGoName(yname string) string {
    if yname == "spi" { return "Spi" }
    if yname == "key" { return "Key" }
    return ""
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetSegmentPath() string {
    return "authenticate-option"
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["spi"] = authenticateOption.Spi
    leafs["key"] = authenticateOption.Key
    return leafs
}

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetBundleName() string { return "cisco_ios_xr" }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetYangName() string { return "authenticate-option" }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) SetParent(parent types.Entity) { authenticateOption.parent = parent }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetParent() types.Entity { return authenticateOption.parent }

func (authenticateOption *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_AuthenticateOption) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes
// heartbeat config for this Customer
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the interval value in second. The type is interface{} with range:
    // 10..3600.
    Interval interface{}

    // Specify the retry value. The type is interface{} with range: 1..10.
    Retries interface{}

    // Specify the timeout value. The type is interface{} with range: 1..3600.
    Timeout interface{}
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetFilter() yfilter.YFilter { return heartBeatAttributes.YFilter }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) SetFilter(yf yfilter.YFilter) { heartBeatAttributes.YFilter = yf }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "retries" { return "Retries" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetSegmentPath() string {
    return "heart-beat-attributes"
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = heartBeatAttributes.Interval
    leafs["retries"] = heartBeatAttributes.Retries
    leafs["timeout"] = heartBeatAttributes.Timeout
    return leafs
}

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetYangName() string { return "heart-beat-attributes" }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) SetParent(parent types.Entity) { heartBeatAttributes.parent = parent }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetParent() types.Entity { return heartBeatAttributes.parent }

func (heartBeatAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_HeartBeatAttributes) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports
// Table of Transport
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customer transport attributes. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport.
    Transport []MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetFilter() yfilter.YFilter { return transports.YFilter }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) SetFilter(yf yfilter.YFilter) { transports.YFilter = yf }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetGoName(yname string) string {
    if yname == "transport" { return "Transport" }
    return ""
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetSegmentPath() string {
    return "transports"
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport" {
        for _, c := range transports.Transport {
            if transports.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport{}
        transports.Transport = append(transports.Transport, child)
        return &transports.Transport[len(transports.Transport)-1]
    }
    return nil
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range transports.Transport {
        children[transports.Transport[i].GetSegmentPath()] = &transports.Transport[i]
    }
    return children
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetBundleName() string { return "cisco_ios_xr" }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetYangName() string { return "transports" }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) SetParent(parent types.Entity) { transports.parent = parent }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetParent() types.Entity { return transports.parent }

func (transports *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport
// Customer transport attributes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..125.
    VrfName interface{}

    // Configure IPv4 address for this LMA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Configure IPv6 address for this LMA. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetFilter() yfilter.YFilter { return transport.YFilter }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) SetFilter(yf yfilter.YFilter) { transport.YFilter = yf }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetSegmentPath() string {
    return "transport" + "[vrf-name='" + fmt.Sprintf("%v", transport.VrfName) + "']"
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = transport.VrfName
    leafs["ipv4-address"] = transport.Ipv4Address
    leafs["ipv6-address"] = transport.Ipv6Address
    return leafs
}

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetBundleName() string { return "cisco_ios_xr" }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetYangName() string { return "transport" }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) SetParent(parent types.Entity) { transport.parent = parent }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetParent() types.Entity { return transport.parent }

func (transport *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_Transports_Transport) GetParentYangName() string { return "transports" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes
// network parameters for the customer
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // not authorize the network prefixes. The type is interface{}.
    Unauthorize interface{}

    // Table of Authorize.
    Authorizes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetFilter() yfilter.YFilter { return networkAttributes.YFilter }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) SetFilter(yf yfilter.YFilter) { networkAttributes.YFilter = yf }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetGoName(yname string) string {
    if yname == "unauthorize" { return "Unauthorize" }
    if yname == "authorizes" { return "Authorizes" }
    return ""
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetSegmentPath() string {
    return "network-attributes"
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authorizes" {
        return &networkAttributes.Authorizes
    }
    return nil
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["authorizes"] = &networkAttributes.Authorizes
    return children
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unauthorize"] = networkAttributes.Unauthorize
    return leafs
}

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetYangName() string { return "network-attributes" }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) SetParent(parent types.Entity) { networkAttributes.parent = parent }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetParent() types.Entity { return networkAttributes.parent }

func (networkAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes
// Table of Authorize
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // not authorize the network prefixes. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize.
    Authorize []MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetFilter() yfilter.YFilter { return authorizes.YFilter }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) SetFilter(yf yfilter.YFilter) { authorizes.YFilter = yf }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetGoName(yname string) string {
    if yname == "authorize" { return "Authorize" }
    return ""
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetSegmentPath() string {
    return "authorizes"
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "authorize" {
        for _, c := range authorizes.Authorize {
            if authorizes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize{}
        authorizes.Authorize = append(authorizes.Authorize, child)
        return &authorizes.Authorize[len(authorizes.Authorize)-1]
    }
    return nil
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range authorizes.Authorize {
        children[authorizes.Authorize[i].GetSegmentPath()] = &authorizes.Authorize[i]
    }
    return children
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetBundleName() string { return "cisco_ios_xr" }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetYangName() string { return "authorizes" }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) SetParent(parent types.Entity) { authorizes.parent = parent }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetParent() types.Entity { return authorizes.parent }

func (authorizes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes) GetParentYangName() string { return "network-attributes" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize
// not authorize the network prefixes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ASCII string. The type is string with length:
    // 1..125.
    Name interface{}

    // Pool configs for this network.
    PoolAttributes MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetFilter() yfilter.YFilter { return authorize.YFilter }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) SetFilter(yf yfilter.YFilter) { authorize.YFilter = yf }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "pool-attributes" { return "PoolAttributes" }
    return ""
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetSegmentPath() string {
    return "authorize" + "[name='" + fmt.Sprintf("%v", authorize.Name) + "']"
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool-attributes" {
        return &authorize.PoolAttributes
    }
    return nil
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pool-attributes"] = &authorize.PoolAttributes
    return children
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = authorize.Name
    return leafs
}

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetBundleName() string { return "cisco_ios_xr" }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetYangName() string { return "authorize" }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) SetParent(parent types.Entity) { authorize.parent = parent }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetParent() types.Entity { return authorize.parent }

func (authorize *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize) GetParentYangName() string { return "authorizes" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes
// Pool configs for this network
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // pool configs for the mobile nodes.
    MobileNode MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode

    // pool configs for the mobile network.
    MobileNetwork MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetFilter() yfilter.YFilter { return poolAttributes.YFilter }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) SetFilter(yf yfilter.YFilter) { poolAttributes.YFilter = yf }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetGoName(yname string) string {
    if yname == "mobile-node" { return "MobileNode" }
    if yname == "mobile-network" { return "MobileNetwork" }
    return ""
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetSegmentPath() string {
    return "pool-attributes"
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mobile-node" {
        return &poolAttributes.MobileNode
    }
    if childYangName == "mobile-network" {
        return &poolAttributes.MobileNetwork
    }
    return nil
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mobile-node"] = &poolAttributes.MobileNode
    children["mobile-network"] = &poolAttributes.MobileNetwork
    return children
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetYangName() string { return "pool-attributes" }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) SetParent(parent types.Entity) { poolAttributes.parent = parent }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetParent() types.Entity { return poolAttributes.parent }

func (poolAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes) GetParentYangName() string { return "authorize" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode
// pool configs for the mobile nodes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None.
    Ipv4Pool MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool

    // None.
    Ipv6Pool MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetFilter() yfilter.YFilter { return mobileNode.YFilter }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) SetFilter(yf yfilter.YFilter) { mobileNode.YFilter = yf }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetGoName(yname string) string {
    if yname == "ipv4-pool" { return "Ipv4Pool" }
    if yname == "ipv6-pool" { return "Ipv6Pool" }
    return ""
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetSegmentPath() string {
    return "mobile-node"
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-pool" {
        return &mobileNode.Ipv4Pool
    }
    if childYangName == "ipv6-pool" {
        return &mobileNode.Ipv6Pool
    }
    return nil
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-pool"] = &mobileNode.Ipv4Pool
    children["ipv6-pool"] = &mobileNode.Ipv6Pool
    return children
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetBundleName() string { return "cisco_ios_xr" }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetYangName() string { return "mobile-node" }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) SetParent(parent types.Entity) { mobileNode.parent = parent }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetParent() types.Entity { return mobileNode.parent }

func (mobileNode *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode) GetParentYangName() string { return "pool-attributes" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool
// None
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool IPv4 start address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetFilter() yfilter.YFilter { return ipv4Pool.YFilter }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) SetFilter(yf yfilter.YFilter) { ipv4Pool.YFilter = yf }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    return ""
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetSegmentPath() string {
    return "ipv4-pool"
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = ipv4Pool.StartAddress
    leafs["pool-prefix"] = ipv4Pool.PoolPrefix
    return leafs
}

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetYangName() string { return "ipv4-pool" }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) SetParent(parent types.Entity) { ipv4Pool.parent = parent }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetParent() types.Entity { return ipv4Pool.parent }

func (ipv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv4Pool) GetParentYangName() string { return "mobile-node" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool
// None
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool IPv6 start address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..62.
    PoolPrefix interface{}
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetFilter() yfilter.YFilter { return ipv6Pool.YFilter }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) SetFilter(yf yfilter.YFilter) { ipv6Pool.YFilter = yf }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    return ""
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetSegmentPath() string {
    return "ipv6-pool"
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = ipv6Pool.StartAddress
    leafs["pool-prefix"] = ipv6Pool.PoolPrefix
    return leafs
}

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetYangName() string { return "ipv6-pool" }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) SetParent(parent types.Entity) { ipv6Pool.parent = parent }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetParent() types.Entity { return ipv6Pool.parent }

func (ipv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNode_Ipv6Pool) GetParentYangName() string { return "mobile-node" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork
// pool configs for the mobile network
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of MRIPV6Pool.
    Mripv6Pools MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools

    // Table of MRIPV4Pool.
    Mripv4Pools MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetFilter() yfilter.YFilter { return mobileNetwork.YFilter }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) SetFilter(yf yfilter.YFilter) { mobileNetwork.YFilter = yf }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetGoName(yname string) string {
    if yname == "mripv6-pools" { return "Mripv6Pools" }
    if yname == "mripv4-pools" { return "Mripv4Pools" }
    return ""
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetSegmentPath() string {
    return "mobile-network"
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv6-pools" {
        return &mobileNetwork.Mripv6Pools
    }
    if childYangName == "mripv4-pools" {
        return &mobileNetwork.Mripv4Pools
    }
    return nil
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mripv6-pools"] = &mobileNetwork.Mripv6Pools
    children["mripv4-pools"] = &mobileNetwork.Mripv4Pools
    return children
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetYangName() string { return "mobile-network" }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) SetParent(parent types.Entity) { mobileNetwork.parent = parent }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetParent() types.Entity { return mobileNetwork.parent }

func (mobileNetwork *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork) GetParentYangName() string { return "pool-attributes" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools
// Table of MRIPV6Pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 pool. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool.
    Mripv6Pool []MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetFilter() yfilter.YFilter { return mripv6Pools.YFilter }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) SetFilter(yf yfilter.YFilter) { mripv6Pools.YFilter = yf }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetGoName(yname string) string {
    if yname == "mripv6-pool" { return "Mripv6Pool" }
    return ""
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetSegmentPath() string {
    return "mripv6-pools"
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv6-pool" {
        for _, c := range mripv6Pools.Mripv6Pool {
            if mripv6Pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool{}
        mripv6Pools.Mripv6Pool = append(mripv6Pools.Mripv6Pool, child)
        return &mripv6Pools.Mripv6Pool[len(mripv6Pools.Mripv6Pool)-1]
    }
    return nil
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mripv6Pools.Mripv6Pool {
        children[mripv6Pools.Mripv6Pool[i].GetSegmentPath()] = &mripv6Pools.Mripv6Pool[i]
    }
    return children
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetBundleName() string { return "cisco_ios_xr" }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetYangName() string { return "mripv6-pools" }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) SetParent(parent types.Entity) { mripv6Pools.parent = parent }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetParent() types.Entity { return mripv6Pools.parent }

func (mripv6Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools) GetParentYangName() string { return "mobile-network" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
// ipv6 pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Pool IPv6 start address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..64.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..64.
    NetworkPrefix interface{}
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetFilter() yfilter.YFilter { return mripv6Pool.YFilter }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) SetFilter(yf yfilter.YFilter) { mripv6Pool.YFilter = yf }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    if yname == "network-prefix" { return "NetworkPrefix" }
    return ""
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetSegmentPath() string {
    return "mripv6-pool" + "[start-address='" + fmt.Sprintf("%v", mripv6Pool.StartAddress) + "']"
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = mripv6Pool.StartAddress
    leafs["pool-prefix"] = mripv6Pool.PoolPrefix
    leafs["network-prefix"] = mripv6Pool.NetworkPrefix
    return leafs
}

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetBundleName() string { return "cisco_ios_xr" }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetYangName() string { return "mripv6-pool" }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) SetParent(parent types.Entity) { mripv6Pool.parent = parent }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetParent() types.Entity { return mripv6Pool.parent }

func (mripv6Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetParentYangName() string { return "mripv6-pools" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools
// Table of MRIPV4Pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 pool. The type is slice of
    // MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool.
    Mripv4Pool []MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetFilter() yfilter.YFilter { return mripv4Pools.YFilter }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) SetFilter(yf yfilter.YFilter) { mripv4Pools.YFilter = yf }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetGoName(yname string) string {
    if yname == "mripv4-pool" { return "Mripv4Pool" }
    return ""
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetSegmentPath() string {
    return "mripv4-pools"
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv4-pool" {
        for _, c := range mripv4Pools.Mripv4Pool {
            if mripv4Pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool{}
        mripv4Pools.Mripv4Pool = append(mripv4Pools.Mripv4Pool, child)
        return &mripv4Pools.Mripv4Pool[len(mripv4Pools.Mripv4Pool)-1]
    }
    return nil
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mripv4Pools.Mripv4Pool {
        children[mripv4Pools.Mripv4Pool[i].GetSegmentPath()] = &mripv4Pools.Mripv4Pool[i]
    }
    return children
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetBundleName() string { return "cisco_ios_xr" }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetYangName() string { return "mripv4-pools" }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) SetParent(parent types.Entity) { mripv4Pools.parent = parent }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetParent() types.Entity { return mripv4Pools.parent }

func (mripv4Pools *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools) GetParentYangName() string { return "mobile-network" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
// ipv4 pool
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Pool IPv4 start address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..32.
    NetworkPrefix interface{}
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetFilter() yfilter.YFilter { return mripv4Pool.YFilter }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) SetFilter(yf yfilter.YFilter) { mripv4Pool.YFilter = yf }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    if yname == "network-prefix" { return "NetworkPrefix" }
    return ""
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetSegmentPath() string {
    return "mripv4-pool" + "[start-address='" + fmt.Sprintf("%v", mripv4Pool.StartAddress) + "']"
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = mripv4Pool.StartAddress
    leafs["pool-prefix"] = mripv4Pool.PoolPrefix
    leafs["network-prefix"] = mripv4Pool.NetworkPrefix
    return leafs
}

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetBundleName() string { return "cisco_ios_xr" }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetYangName() string { return "mripv4-pool" }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) SetParent(parent types.Entity) { mripv4Pool.parent = parent }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetParent() types.Entity { return mripv4Pool.parent }

func (mripv4Pool *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_NetworkAttributes_Authorizes_Authorize_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetParentYangName() string { return "mripv4-pools" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey
// Customer specific GRE key
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // GRE key type. The type is GreKeyType.
    GreKeyType interface{}

    // GRE key value. The type is interface{} with range: 1..4294967295.
    GreKeyValue interface{}
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetFilter() yfilter.YFilter { return greKey.YFilter }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) SetFilter(yf yfilter.YFilter) { greKey.YFilter = yf }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetGoName(yname string) string {
    if yname == "gre-key-type" { return "GreKeyType" }
    if yname == "gre-key-value" { return "GreKeyValue" }
    return ""
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetSegmentPath() string {
    return "gre-key"
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["gre-key-type"] = greKey.GreKeyType
    leafs["gre-key-value"] = greKey.GreKeyValue
    return leafs
}

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetBundleName() string { return "cisco_ios_xr" }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetYangName() string { return "gre-key" }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) SetParent(parent types.Entity) { greKey.parent = parent }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetParent() types.Entity { return greKey.parent }

func (greKey *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_GreKey) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes
// Customer specific binding attributes
type MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Maximum bce lifetime permitted. The type is interface{} with range:
    // 10..65535. Units are second.
    MaxLifeTime interface{}
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetFilter() yfilter.YFilter { return bindingAttributes.YFilter }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) SetFilter(yf yfilter.YFilter) { bindingAttributes.YFilter = yf }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetGoName(yname string) string {
    if yname == "max-life-time" { return "MaxLifeTime" }
    return ""
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetSegmentPath() string {
    return "binding-attributes"
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-life-time"] = bindingAttributes.MaxLifeTime
    return leafs
}

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetYangName() string { return "binding-attributes" }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) SetParent(parent types.Entity) { bindingAttributes.parent = parent }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetParent() types.Entity { return bindingAttributes.parent }

func (bindingAttributes *MobileIp_Lmas_Lma_Services_Service_Customers_Customer_BindingAttributes) GetParentYangName() string { return "customer" }

// MobileIp_Lmas_Lma_Networks
// Table of Network
type MobileIp_Lmas_Lma_Networks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // network for this LMA. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network.
    Network []MobileIp_Lmas_Lma_Networks_Network
}

func (networks *MobileIp_Lmas_Lma_Networks) GetFilter() yfilter.YFilter { return networks.YFilter }

func (networks *MobileIp_Lmas_Lma_Networks) SetFilter(yf yfilter.YFilter) { networks.YFilter = yf }

func (networks *MobileIp_Lmas_Lma_Networks) GetGoName(yname string) string {
    if yname == "network" { return "Network" }
    return ""
}

func (networks *MobileIp_Lmas_Lma_Networks) GetSegmentPath() string {
    return "networks"
}

func (networks *MobileIp_Lmas_Lma_Networks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network" {
        for _, c := range networks.Network {
            if networks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Networks_Network{}
        networks.Network = append(networks.Network, child)
        return &networks.Network[len(networks.Network)-1]
    }
    return nil
}

func (networks *MobileIp_Lmas_Lma_Networks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networks.Network {
        children[networks.Network[i].GetSegmentPath()] = &networks.Network[i]
    }
    return children
}

func (networks *MobileIp_Lmas_Lma_Networks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networks *MobileIp_Lmas_Lma_Networks) GetBundleName() string { return "cisco_ios_xr" }

func (networks *MobileIp_Lmas_Lma_Networks) GetYangName() string { return "networks" }

func (networks *MobileIp_Lmas_Lma_Networks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (networks *MobileIp_Lmas_Lma_Networks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (networks *MobileIp_Lmas_Lma_Networks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (networks *MobileIp_Lmas_Lma_Networks) SetParent(parent types.Entity) { networks.parent = parent }

func (networks *MobileIp_Lmas_Lma_Networks) GetParent() types.Entity { return networks.parent }

func (networks *MobileIp_Lmas_Lma_Networks) GetParentYangName() string { return "lma" }

// MobileIp_Lmas_Lma_Networks_Network
// network for this LMA
type MobileIp_Lmas_Lma_Networks_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Network name. The type is string with length:
    // 1..125.
    LmaNetwork interface{}

    // Pool configs for this network.
    PoolAttributes MobileIp_Lmas_Lma_Networks_Network_PoolAttributes
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *MobileIp_Lmas_Lma_Networks_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetGoName(yname string) string {
    if yname == "lma-network" { return "LmaNetwork" }
    if yname == "pool-attributes" { return "PoolAttributes" }
    return ""
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetSegmentPath() string {
    return "network" + "[lma-network='" + fmt.Sprintf("%v", network.LmaNetwork) + "']"
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pool-attributes" {
        return &network.PoolAttributes
    }
    return nil
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pool-attributes"] = &network.PoolAttributes
    return children
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-network"] = network.LmaNetwork
    return leafs
}

func (network *MobileIp_Lmas_Lma_Networks_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetYangName() string { return "network" }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *MobileIp_Lmas_Lma_Networks_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetParent() types.Entity { return network.parent }

func (network *MobileIp_Lmas_Lma_Networks_Network) GetParentYangName() string { return "networks" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes
// Pool configs for this network
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // pool configs for the mobile nodes.
    MobileNode MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode

    // pool configs for the mobile network.
    MobileNetwork MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetFilter() yfilter.YFilter { return poolAttributes.YFilter }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) SetFilter(yf yfilter.YFilter) { poolAttributes.YFilter = yf }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetGoName(yname string) string {
    if yname == "mobile-node" { return "MobileNode" }
    if yname == "mobile-network" { return "MobileNetwork" }
    return ""
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetSegmentPath() string {
    return "pool-attributes"
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mobile-node" {
        return &poolAttributes.MobileNode
    }
    if childYangName == "mobile-network" {
        return &poolAttributes.MobileNetwork
    }
    return nil
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mobile-node"] = &poolAttributes.MobileNode
    children["mobile-network"] = &poolAttributes.MobileNetwork
    return children
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetYangName() string { return "pool-attributes" }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) SetParent(parent types.Entity) { poolAttributes.parent = parent }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetParent() types.Entity { return poolAttributes.parent }

func (poolAttributes *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes) GetParentYangName() string { return "network" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode
// pool configs for the mobile nodes
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None.
    Ipv4Pool MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool

    // None.
    Ipv6Pool MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetFilter() yfilter.YFilter { return mobileNode.YFilter }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) SetFilter(yf yfilter.YFilter) { mobileNode.YFilter = yf }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetGoName(yname string) string {
    if yname == "ipv4-pool" { return "Ipv4Pool" }
    if yname == "ipv6-pool" { return "Ipv6Pool" }
    return ""
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetSegmentPath() string {
    return "mobile-node"
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-pool" {
        return &mobileNode.Ipv4Pool
    }
    if childYangName == "ipv6-pool" {
        return &mobileNode.Ipv6Pool
    }
    return nil
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-pool"] = &mobileNode.Ipv4Pool
    children["ipv6-pool"] = &mobileNode.Ipv6Pool
    return children
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetBundleName() string { return "cisco_ios_xr" }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetYangName() string { return "mobile-node" }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) SetParent(parent types.Entity) { mobileNode.parent = parent }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetParent() types.Entity { return mobileNode.parent }

func (mobileNode *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode) GetParentYangName() string { return "pool-attributes" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool
// None
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool IPv4 start address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetFilter() yfilter.YFilter { return ipv4Pool.YFilter }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) SetFilter(yf yfilter.YFilter) { ipv4Pool.YFilter = yf }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    return ""
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetSegmentPath() string {
    return "ipv4-pool"
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = ipv4Pool.StartAddress
    leafs["pool-prefix"] = ipv4Pool.PoolPrefix
    return leafs
}

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetYangName() string { return "ipv4-pool" }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) SetParent(parent types.Entity) { ipv4Pool.parent = parent }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetParent() types.Entity { return ipv4Pool.parent }

func (ipv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv4Pool) GetParentYangName() string { return "mobile-node" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool
// None
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Pool IPv6 start address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..62.
    PoolPrefix interface{}
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetFilter() yfilter.YFilter { return ipv6Pool.YFilter }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) SetFilter(yf yfilter.YFilter) { ipv6Pool.YFilter = yf }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    return ""
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetSegmentPath() string {
    return "ipv6-pool"
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = ipv6Pool.StartAddress
    leafs["pool-prefix"] = ipv6Pool.PoolPrefix
    return leafs
}

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetYangName() string { return "ipv6-pool" }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) SetParent(parent types.Entity) { ipv6Pool.parent = parent }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetParent() types.Entity { return ipv6Pool.parent }

func (ipv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNode_Ipv6Pool) GetParentYangName() string { return "mobile-node" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork
// pool configs for the mobile network
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of MRIPV6Pool.
    Mripv6Pools MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools

    // Table of MRIPV4Pool.
    Mripv4Pools MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetFilter() yfilter.YFilter { return mobileNetwork.YFilter }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) SetFilter(yf yfilter.YFilter) { mobileNetwork.YFilter = yf }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetGoName(yname string) string {
    if yname == "mripv6-pools" { return "Mripv6Pools" }
    if yname == "mripv4-pools" { return "Mripv4Pools" }
    return ""
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetSegmentPath() string {
    return "mobile-network"
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv6-pools" {
        return &mobileNetwork.Mripv6Pools
    }
    if childYangName == "mripv4-pools" {
        return &mobileNetwork.Mripv4Pools
    }
    return nil
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mripv6-pools"] = &mobileNetwork.Mripv6Pools
    children["mripv4-pools"] = &mobileNetwork.Mripv4Pools
    return children
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetBundleName() string { return "cisco_ios_xr" }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetYangName() string { return "mobile-network" }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) SetParent(parent types.Entity) { mobileNetwork.parent = parent }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetParent() types.Entity { return mobileNetwork.parent }

func (mobileNetwork *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork) GetParentYangName() string { return "pool-attributes" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools
// Table of MRIPV6Pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv6 pool. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool.
    Mripv6Pool []MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetFilter() yfilter.YFilter { return mripv6Pools.YFilter }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) SetFilter(yf yfilter.YFilter) { mripv6Pools.YFilter = yf }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetGoName(yname string) string {
    if yname == "mripv6-pool" { return "Mripv6Pool" }
    return ""
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetSegmentPath() string {
    return "mripv6-pools"
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv6-pool" {
        for _, c := range mripv6Pools.Mripv6Pool {
            if mripv6Pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool{}
        mripv6Pools.Mripv6Pool = append(mripv6Pools.Mripv6Pool, child)
        return &mripv6Pools.Mripv6Pool[len(mripv6Pools.Mripv6Pool)-1]
    }
    return nil
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mripv6Pools.Mripv6Pool {
        children[mripv6Pools.Mripv6Pool[i].GetSegmentPath()] = &mripv6Pools.Mripv6Pool[i]
    }
    return children
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetBundleName() string { return "cisco_ios_xr" }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetYangName() string { return "mripv6-pools" }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) SetParent(parent types.Entity) { mripv6Pools.parent = parent }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetParent() types.Entity { return mripv6Pools.parent }

func (mripv6Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools) GetParentYangName() string { return "mobile-network" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool
// ipv6 pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Pool IPv6 start address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv6 Pool Prefix value. The type is interface{} with range: 8..64.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..64.
    NetworkPrefix interface{}
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetFilter() yfilter.YFilter { return mripv6Pool.YFilter }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) SetFilter(yf yfilter.YFilter) { mripv6Pool.YFilter = yf }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    if yname == "network-prefix" { return "NetworkPrefix" }
    return ""
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetSegmentPath() string {
    return "mripv6-pool" + "[start-address='" + fmt.Sprintf("%v", mripv6Pool.StartAddress) + "']"
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = mripv6Pool.StartAddress
    leafs["pool-prefix"] = mripv6Pool.PoolPrefix
    leafs["network-prefix"] = mripv6Pool.NetworkPrefix
    return leafs
}

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetBundleName() string { return "cisco_ios_xr" }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetYangName() string { return "mripv6-pool" }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) SetParent(parent types.Entity) { mripv6Pool.parent = parent }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetParent() types.Entity { return mripv6Pool.parent }

func (mripv6Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv6Pools_Mripv6Pool) GetParentYangName() string { return "mripv6-pools" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools
// Table of MRIPV4Pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 pool. The type is slice of
    // MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool.
    Mripv4Pool []MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetFilter() yfilter.YFilter { return mripv4Pools.YFilter }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) SetFilter(yf yfilter.YFilter) { mripv4Pools.YFilter = yf }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetGoName(yname string) string {
    if yname == "mripv4-pool" { return "Mripv4Pool" }
    return ""
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetSegmentPath() string {
    return "mripv4-pools"
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mripv4-pool" {
        for _, c := range mripv4Pools.Mripv4Pool {
            if mripv4Pools.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool{}
        mripv4Pools.Mripv4Pool = append(mripv4Pools.Mripv4Pool, child)
        return &mripv4Pools.Mripv4Pool[len(mripv4Pools.Mripv4Pool)-1]
    }
    return nil
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mripv4Pools.Mripv4Pool {
        children[mripv4Pools.Mripv4Pool[i].GetSegmentPath()] = &mripv4Pools.Mripv4Pool[i]
    }
    return children
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetBundleName() string { return "cisco_ios_xr" }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetYangName() string { return "mripv4-pools" }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) SetParent(parent types.Entity) { mripv4Pools.parent = parent }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetParent() types.Entity { return mripv4Pools.parent }

func (mripv4Pools *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools) GetParentYangName() string { return "mobile-network" }

// MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool
// ipv4 pool
type MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Pool IPv4 start address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StartAddress interface{}

    // IPv4 Pool Prefix value. The type is interface{} with range: 8..30.
    PoolPrefix interface{}

    // IPv4 Network Prefix value. The type is interface{} with range: 8..32.
    NetworkPrefix interface{}
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetFilter() yfilter.YFilter { return mripv4Pool.YFilter }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) SetFilter(yf yfilter.YFilter) { mripv4Pool.YFilter = yf }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetGoName(yname string) string {
    if yname == "start-address" { return "StartAddress" }
    if yname == "pool-prefix" { return "PoolPrefix" }
    if yname == "network-prefix" { return "NetworkPrefix" }
    return ""
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetSegmentPath() string {
    return "mripv4-pool" + "[start-address='" + fmt.Sprintf("%v", mripv4Pool.StartAddress) + "']"
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-address"] = mripv4Pool.StartAddress
    leafs["pool-prefix"] = mripv4Pool.PoolPrefix
    leafs["network-prefix"] = mripv4Pool.NetworkPrefix
    return leafs
}

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetBundleName() string { return "cisco_ios_xr" }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetYangName() string { return "mripv4-pool" }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) SetParent(parent types.Entity) { mripv4Pool.parent = parent }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetParent() types.Entity { return mripv4Pool.parent }

func (mripv4Pool *MobileIp_Lmas_Lma_Networks_Network_PoolAttributes_MobileNetwork_Mripv4Pools_Mripv4Pool) GetParentYangName() string { return "mripv4-pools" }

// MobileIp_Lmas_Lma_ReplayProtection
// Replay Protection Method
type MobileIp_Lmas_Lma_ReplayProtection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify timestamp window value in seconds. The type is interface{} with
    // range: 1..255. Units are second.
    TimestampWindow interface{}
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetFilter() yfilter.YFilter { return replayProtection.YFilter }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) SetFilter(yf yfilter.YFilter) { replayProtection.YFilter = yf }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetGoName(yname string) string {
    if yname == "timestamp-window" { return "TimestampWindow" }
    return ""
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetSegmentPath() string {
    return "replay-protection"
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timestamp-window"] = replayProtection.TimestampWindow
    return leafs
}

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetBundleName() string { return "cisco_ios_xr" }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetYangName() string { return "replay-protection" }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) SetParent(parent types.Entity) { replayProtection.parent = parent }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetParent() types.Entity { return replayProtection.parent }

func (replayProtection *MobileIp_Lmas_Lma_ReplayProtection) GetParentYangName() string { return "lma" }

