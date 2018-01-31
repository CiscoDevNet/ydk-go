// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-mobileip package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pmipv6: Proxy Mobile IPv6
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_mobileip_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_mobileip_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-mobileip-oper pmipv6}", reflect.TypeOf(Pmipv6{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-mobileip-oper:pmipv6", reflect.TypeOf(Pmipv6{}))
}

// Pmipv6Role represents PMIPV6 Role Types
type Pmipv6Role string

const (
    // WLAN
    Pmipv6Role_wlan Pmipv6Role = "wlan"

    // 3GPP
    Pmipv6Role_gpp Pmipv6Role = "gpp"

    // LTE
    Pmipv6Role_lte Pmipv6Role = "lte"

    // WiMAX
    Pmipv6Role_wi_max Pmipv6Role = "wi-max"

    // 3GMA
    Pmipv6Role_gma Pmipv6Role = "gma"

    // MAX Role
    Pmipv6Role_rmax Pmipv6Role = "rmax"
)

// Pmipv6Encap represents ENCAP Types
type Pmipv6Encap string

const (
    // None
    Pmipv6Encap_none Pmipv6Encap = "none"

    // IPV6 Tunnel
    Pmipv6Encap_ipv6 Pmipv6Encap = "ipv6"

    // IPV6 in IPV4 Tunnel
    Pmipv6Encap_ipv6_ipv4 Pmipv6Encap = "ipv6-ipv4"

    // IPV6 in IPV4 UDP Tunnel
    Pmipv6Encap_ipv6_udp Pmipv6Encap = "ipv6-udp"

    // GRE IPV4 Tunnel
    Pmipv6Encap_gre_ipv4 Pmipv6Encap = "gre-ipv4"

    // GRE IPV6 Tunnel
    Pmipv6Encap_gre_ipv6 Pmipv6Encap = "gre-ipv6"

    // GRE Tunnel
    Pmipv6Encap_gre Pmipv6Encap = "gre"

    // MGRE IPV4 Tunnel
    Pmipv6Encap_mgre_ipv4 Pmipv6Encap = "mgre-ipv4"

    // MGRE IPV6 Tunnel
    Pmipv6Encap_mgre_ipv6 Pmipv6Encap = "mgre-ipv6"

    // MIP UDP Tunnel
    Pmipv6Encap_mip_udp Pmipv6Encap = "mip-udp"

    // MIP MUDP Tunnel
    Pmipv6Encap_mip_mudp Pmipv6Encap = "mip-mudp"

    // MAX Encap Type
    Pmipv6Encap_max Pmipv6Encap = "max"
)

// Pmipv6Addr represents Address Types
type Pmipv6Addr string

const (
    // None
    Pmipv6Addr_none Pmipv6Addr = "none"

    // IPV4 Address
    Pmipv6Addr_ipv4 Pmipv6Addr = "ipv4"

    // IPV6 Address
    Pmipv6Addr_ipv6 Pmipv6Addr = "ipv6"

    // Both IPV4 and IPV6 Address
    Pmipv6Addr_pmipv6_addr_ipv4_ipv6 Pmipv6Addr = "pmipv6-addr-ipv4-ipv6"
)

// Pmipv6
// Proxy Mobile IPv6
type Pmipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None.
    Lma Pmipv6_Lma
}

func (pmipv6 *Pmipv6) GetFilter() yfilter.YFilter { return pmipv6.YFilter }

func (pmipv6 *Pmipv6) SetFilter(yf yfilter.YFilter) { pmipv6.YFilter = yf }

func (pmipv6 *Pmipv6) GetGoName(yname string) string {
    if yname == "lma" { return "Lma" }
    return ""
}

func (pmipv6 *Pmipv6) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-mobileip-oper:pmipv6"
}

func (pmipv6 *Pmipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lma" {
        return &pmipv6.Lma
    }
    return nil
}

func (pmipv6 *Pmipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lma"] = &pmipv6.Lma
    return children
}

func (pmipv6 *Pmipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pmipv6 *Pmipv6) GetBundleName() string { return "cisco_ios_xr" }

func (pmipv6 *Pmipv6) GetYangName() string { return "pmipv6" }

func (pmipv6 *Pmipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pmipv6 *Pmipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pmipv6 *Pmipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pmipv6 *Pmipv6) SetParent(parent types.Entity) { pmipv6.parent = parent }

func (pmipv6 *Pmipv6) GetParent() types.Entity { return pmipv6.parent }

func (pmipv6 *Pmipv6) GetParentYangName() string { return "Cisco-IOS-XR-ip-mobileip-oper" }

// Pmipv6_Lma
// None
type Pmipv6_Lma struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None.
    Statistics Pmipv6_Lma_Statistics

    // Table of Binding.
    Bindings Pmipv6_Lma_Bindings

    // Table of Heartbeat.
    Heartbeats Pmipv6_Lma_Heartbeats

    // Global Configuration Variables.
    ConfigVariables Pmipv6_Lma_ConfigVariables
}

func (lma *Pmipv6_Lma) GetFilter() yfilter.YFilter { return lma.YFilter }

func (lma *Pmipv6_Lma) SetFilter(yf yfilter.YFilter) { lma.YFilter = yf }

func (lma *Pmipv6_Lma) GetGoName(yname string) string {
    if yname == "statistics" { return "Statistics" }
    if yname == "bindings" { return "Bindings" }
    if yname == "heartbeats" { return "Heartbeats" }
    if yname == "config-variables" { return "ConfigVariables" }
    return ""
}

func (lma *Pmipv6_Lma) GetSegmentPath() string {
    return "lma"
}

func (lma *Pmipv6_Lma) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &lma.Statistics
    }
    if childYangName == "bindings" {
        return &lma.Bindings
    }
    if childYangName == "heartbeats" {
        return &lma.Heartbeats
    }
    if childYangName == "config-variables" {
        return &lma.ConfigVariables
    }
    return nil
}

func (lma *Pmipv6_Lma) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &lma.Statistics
    children["bindings"] = &lma.Bindings
    children["heartbeats"] = &lma.Heartbeats
    children["config-variables"] = &lma.ConfigVariables
    return children
}

func (lma *Pmipv6_Lma) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lma *Pmipv6_Lma) GetBundleName() string { return "cisco_ios_xr" }

func (lma *Pmipv6_Lma) GetYangName() string { return "lma" }

func (lma *Pmipv6_Lma) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lma *Pmipv6_Lma) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lma *Pmipv6_Lma) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lma *Pmipv6_Lma) SetParent(parent types.Entity) { lma.parent = parent }

func (lma *Pmipv6_Lma) GetParent() types.Entity { return lma.parent }

func (lma *Pmipv6_Lma) GetParentYangName() string { return "pmipv6" }

// Pmipv6_Lma_Statistics
// None
type Pmipv6_Lma_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of CustomerStatistics.
    CustomerStatistics Pmipv6_Lma_Statistics_CustomerStatistics

    // LMA License Statistics.
    License Pmipv6_Lma_Statistics_License

    // Global Statistics.
    Global Pmipv6_Lma_Statistics_Global

    // Table of MAGStatistics.
    MagStatistics Pmipv6_Lma_Statistics_MagStatistics
}

func (statistics *Pmipv6_Lma_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Pmipv6_Lma_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Pmipv6_Lma_Statistics) GetGoName(yname string) string {
    if yname == "customer-statistics" { return "CustomerStatistics" }
    if yname == "license" { return "License" }
    if yname == "global" { return "Global" }
    if yname == "mag-statistics" { return "MagStatistics" }
    return ""
}

func (statistics *Pmipv6_Lma_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Pmipv6_Lma_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customer-statistics" {
        return &statistics.CustomerStatistics
    }
    if childYangName == "license" {
        return &statistics.License
    }
    if childYangName == "global" {
        return &statistics.Global
    }
    if childYangName == "mag-statistics" {
        return &statistics.MagStatistics
    }
    return nil
}

func (statistics *Pmipv6_Lma_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["customer-statistics"] = &statistics.CustomerStatistics
    children["license"] = &statistics.License
    children["global"] = &statistics.Global
    children["mag-statistics"] = &statistics.MagStatistics
    return children
}

func (statistics *Pmipv6_Lma_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Pmipv6_Lma_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Pmipv6_Lma_Statistics) GetYangName() string { return "statistics" }

func (statistics *Pmipv6_Lma_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Pmipv6_Lma_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Pmipv6_Lma_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Pmipv6_Lma_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Pmipv6_Lma_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Pmipv6_Lma_Statistics) GetParentYangName() string { return "lma" }

// Pmipv6_Lma_Statistics_CustomerStatistics
// Table of CustomerStatistics
type Pmipv6_Lma_Statistics_CustomerStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customer statistics. The type is slice of
    // Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic.
    CustomerStatistic []Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetFilter() yfilter.YFilter { return customerStatistics.YFilter }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) SetFilter(yf yfilter.YFilter) { customerStatistics.YFilter = yf }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetGoName(yname string) string {
    if yname == "customer-statistic" { return "CustomerStatistic" }
    return ""
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetSegmentPath() string {
    return "customer-statistics"
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customer-statistic" {
        for _, c := range customerStatistics.CustomerStatistic {
            if customerStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic{}
        customerStatistics.CustomerStatistic = append(customerStatistics.CustomerStatistic, child)
        return &customerStatistics.CustomerStatistic[len(customerStatistics.CustomerStatistic)-1]
    }
    return nil
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range customerStatistics.CustomerStatistic {
        children[customerStatistics.CustomerStatistic[i].GetSegmentPath()] = &customerStatistics.CustomerStatistic[i]
    }
    return children
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetYangName() string { return "customer-statistics" }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) SetParent(parent types.Entity) { customerStatistics.parent = parent }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetParent() types.Entity { return customerStatistics.parent }

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetParentYangName() string { return "statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic
// Customer statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Customer Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    CustomerName interface{}

    // LMA Identifier. The type is string.
    LmaIdentifier interface{}

    // Count of Bindings. The type is interface{} with range: 0..4294967295.
    BceCount interface{}

    // Count of Handoffs. The type is interface{} with range: 0..4294967295.
    HandoffCount interface{}

    // Count of IPv4 Mobile Node Prefixes. The type is interface{} with range:
    // 0..4294967295.
    Ipv4MnpCount interface{}

    // Count of IPv6 Mobile Node Prefixes. The type is interface{} with range:
    // 0..4294967295.
    Ipv6MnpCount interface{}

    // LMA Protocol Statistics.
    ProtocolStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics

    // LMA Accounting Statistics.
    AccountingStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetFilter() yfilter.YFilter { return customerStatistic.YFilter }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) SetFilter(yf yfilter.YFilter) { customerStatistic.YFilter = yf }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetGoName(yname string) string {
    if yname == "customer-name" { return "CustomerName" }
    if yname == "lma-identifier" { return "LmaIdentifier" }
    if yname == "bce-count" { return "BceCount" }
    if yname == "handoff-count" { return "HandoffCount" }
    if yname == "ipv4-mnp-count" { return "Ipv4MnpCount" }
    if yname == "ipv6-mnp-count" { return "Ipv6MnpCount" }
    if yname == "protocol-statistics" { return "ProtocolStatistics" }
    if yname == "accounting-statistics" { return "AccountingStatistics" }
    return ""
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetSegmentPath() string {
    return "customer-statistic" + "[customer-name='" + fmt.Sprintf("%v", customerStatistic.CustomerName) + "']"
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol-statistics" {
        return &customerStatistic.ProtocolStatistics
    }
    if childYangName == "accounting-statistics" {
        return &customerStatistic.AccountingStatistics
    }
    return nil
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol-statistics"] = &customerStatistic.ProtocolStatistics
    children["accounting-statistics"] = &customerStatistic.AccountingStatistics
    return children
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["customer-name"] = customerStatistic.CustomerName
    leafs["lma-identifier"] = customerStatistic.LmaIdentifier
    leafs["bce-count"] = customerStatistic.BceCount
    leafs["handoff-count"] = customerStatistic.HandoffCount
    leafs["ipv4-mnp-count"] = customerStatistic.Ipv4MnpCount
    leafs["ipv6-mnp-count"] = customerStatistic.Ipv6MnpCount
    return leafs
}

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetYangName() string { return "customer-statistic" }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) SetParent(parent types.Entity) { customerStatistic.parent = parent }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetParent() types.Entity { return customerStatistic.parent }

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetParentYangName() string { return "customer-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBU Receive Statistics.
    PbuReceiveStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics

    // PBA Send Statistics.
    PbaSendStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics

    // PBRI Send Statistics.
    PbriSendStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics

    // PBRI Receive Statistics.
    PbriReceiveStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics

    // PBRA Send Statistics.
    PbraSendStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics

    // PBRA Receive Statistics.
    PbraReceiveStatistics Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetFilter() yfilter.YFilter { return protocolStatistics.YFilter }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) SetFilter(yf yfilter.YFilter) { protocolStatistics.YFilter = yf }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetGoName(yname string) string {
    if yname == "pbu-receive-statistics" { return "PbuReceiveStatistics" }
    if yname == "pba-send-statistics" { return "PbaSendStatistics" }
    if yname == "pbri-send-statistics" { return "PbriSendStatistics" }
    if yname == "pbri-receive-statistics" { return "PbriReceiveStatistics" }
    if yname == "pbra-send-statistics" { return "PbraSendStatistics" }
    if yname == "pbra-receive-statistics" { return "PbraReceiveStatistics" }
    return ""
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetSegmentPath() string {
    return "protocol-statistics"
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbu-receive-statistics" {
        return &protocolStatistics.PbuReceiveStatistics
    }
    if childYangName == "pba-send-statistics" {
        return &protocolStatistics.PbaSendStatistics
    }
    if childYangName == "pbri-send-statistics" {
        return &protocolStatistics.PbriSendStatistics
    }
    if childYangName == "pbri-receive-statistics" {
        return &protocolStatistics.PbriReceiveStatistics
    }
    if childYangName == "pbra-send-statistics" {
        return &protocolStatistics.PbraSendStatistics
    }
    if childYangName == "pbra-receive-statistics" {
        return &protocolStatistics.PbraReceiveStatistics
    }
    return nil
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbu-receive-statistics"] = &protocolStatistics.PbuReceiveStatistics
    children["pba-send-statistics"] = &protocolStatistics.PbaSendStatistics
    children["pbri-send-statistics"] = &protocolStatistics.PbriSendStatistics
    children["pbri-receive-statistics"] = &protocolStatistics.PbriReceiveStatistics
    children["pbra-send-statistics"] = &protocolStatistics.PbraSendStatistics
    children["pbra-receive-statistics"] = &protocolStatistics.PbraReceiveStatistics
    return children
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetYangName() string { return "protocol-statistics" }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) SetParent(parent types.Entity) { protocolStatistics.parent = parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetParent() types.Entity { return protocolStatistics.parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetParentYangName() string { return "customer-statistic" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetFilter() yfilter.YFilter { return pbuReceiveStatistics.YFilter }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbuReceiveStatistics.YFilter = yf }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbu-count" { return "PbuCount" }
    if yname == "pbu-drop-count" { return "PbuDropCount" }
    return ""
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetSegmentPath() string {
    return "pbu-receive-statistics"
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbu-count"] = pbuReceiveStatistics.PbuCount
    leafs["pbu-drop-count"] = pbuReceiveStatistics.PbuDropCount
    return leafs
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetYangName() string { return "pbu-receive-statistics" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) SetParent(parent types.Entity) { pbuReceiveStatistics.parent = parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetParent() types.Entity { return pbuReceiveStatistics.parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBAs. The type is interface{} with range: 0..18446744073709551615.
    PbaCount interface{}

    // Count of PBAs dropped. The type is interface{} with range: 0..4294967295.
    PbaDropCount interface{}

    // Count of Status Code - Binding Update accepted. The type is interface{}
    // with range: 0..4294967295.
    AcceptedCount interface{}

    // Count of Status Code - Last BA status code sent. The type is interface{}
    // with range: 0..4294967295.
    UnknownCount interface{}

    // Count of Status Code - Reason unspecified. The type is interface{} with
    // range: 0..4294967295.
    UnspecifiedFailureCount interface{}

    // Count of Status Code - Administratively prohibited. The type is interface{}
    // with range: 0..4294967295.
    AdminFailureCount interface{}

    // Count of Status Code - Insufficient resources. The type is interface{} with
    // range: 0..4294967295.
    ResourceFailureCount interface{}

    // Count of Status Code - Home registration not supported. The type is
    // interface{} with range: 0..4294967295.
    HomeRegFailureCount interface{}

    // Count of Status Code - Not home subnet. The type is interface{} with range:
    // 0..4294967295.
    HomeSubnetFailureCount interface{}

    // Count of Status Code - Sequence number out of window. The type is
    // interface{} with range: 0..4294967295.
    BadSequenceFailureCount interface{}

    // Count of Status Code - Registration type change. The type is interface{}
    // with range: 0..4294967295.
    RegTypeFailureCount interface{}

    // Count of Status Code - Auth Fail. The type is interface{} with range:
    // 0..4294967295.
    AuthenFailureCount interface{}

    // Count of Status Code - Proxy Registration not enabled. The type is
    // interface{} with range: 0..4294967295.
    ProxyRegNotEnabledCount interface{}

    // Count of Status Code - Not LMA for this Mobile Node. The type is
    // interface{} with range: 0..4294967295.
    NotLmaForThisMnCount interface{}

    // Count of Status Code - MAG not auth.for proxyreg. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForProxyRegCount interface{}

    // Count of Status Code - Not authorized for HNP. The type is interface{} with
    // range: 0..4294967295.
    NoAuthorForHnpCount interface{}

    // Count of Status Code - Invalid timestamp value. The type is interface{}
    // with range: 0..4294967295.
    TimestampMismatchCount interface{}

    // Count of Status Code - Timestamp lower than previous accepted. The type is
    // interface{} with range: 0..4294967295.
    TimestampLowerThanPreviousAcceptedCount interface{}

    // Count of Status Code - Missing Home Network Prefix option. The type is
    // interface{} with range: 0..4294967295.
    MissingHnpOptCount interface{}

    // Count of Status Code - Recevied HNPs do not match with BCE. The type is
    // interface{} with range: 0..4294967295.
    ReceivedHnpsDoNotMatchBceHnpsCount interface{}

    // Count of Status Code - Missing MN identifier option. The type is
    // interface{} with range: 0..4294967295.
    MissingMnIdOptCount interface{}

    // Count of Status Code - Missing Handoff Indicator. The type is interface{}
    // with range: 0..4294967295.
    MissingHiOptCount interface{}

    // Count of Status Code - Missing ATT option. The type is interface{} with
    // range: 0..4294967295.
    MissingAccessTechTypeOptCount interface{}

    // Count of Status Code - Not authorized for IPv4 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv4MobilityCount interface{}

    // Count of Status Code - Not authorized for IPv4 HoA. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForIpv4HoaCount interface{}

    // Count of Status Code - Not authorized for IPv6 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv6MobilityCount interface{}

    // Count of Status Code - Multiple IPv4 HoA not supported. The type is
    // interface{} with range: 0..4294967295.
    MultipleIpv4HoANotSupportedCount interface{}

    // Count of Status Code - GRE Key option is required. The type is interface{}
    // with range: 0..4294967295.
    GreKeyOptRequiredCount interface{}
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetFilter() yfilter.YFilter { return pbaSendStatistics.YFilter }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) SetFilter(yf yfilter.YFilter) { pbaSendStatistics.YFilter = yf }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetGoName(yname string) string {
    if yname == "pba-count" { return "PbaCount" }
    if yname == "pba-drop-count" { return "PbaDropCount" }
    if yname == "accepted-count" { return "AcceptedCount" }
    if yname == "unknown-count" { return "UnknownCount" }
    if yname == "unspecified-failure-count" { return "UnspecifiedFailureCount" }
    if yname == "admin-failure-count" { return "AdminFailureCount" }
    if yname == "resource-failure-count" { return "ResourceFailureCount" }
    if yname == "home-reg-failure-count" { return "HomeRegFailureCount" }
    if yname == "home-subnet-failure-count" { return "HomeSubnetFailureCount" }
    if yname == "bad-sequence-failure-count" { return "BadSequenceFailureCount" }
    if yname == "reg-type-failure-count" { return "RegTypeFailureCount" }
    if yname == "authen-failure-count" { return "AuthenFailureCount" }
    if yname == "proxy-reg-not-enabled-count" { return "ProxyRegNotEnabledCount" }
    if yname == "not-lma-for-this-mn-count" { return "NotLmaForThisMnCount" }
    if yname == "no-author-for-proxy-reg-count" { return "NoAuthorForProxyRegCount" }
    if yname == "no-author-for-hnp-count" { return "NoAuthorForHnpCount" }
    if yname == "timestamp-mismatch-count" { return "TimestampMismatchCount" }
    if yname == "timestamp-lower-than-previous-accepted-count" { return "TimestampLowerThanPreviousAcceptedCount" }
    if yname == "missing-hnp-opt-count" { return "MissingHnpOptCount" }
    if yname == "received-hnps-do-not-match-bce-hnps-count" { return "ReceivedHnpsDoNotMatchBceHnpsCount" }
    if yname == "missing-mn-id-opt-count" { return "MissingMnIdOptCount" }
    if yname == "missing-hi-opt-count" { return "MissingHiOptCount" }
    if yname == "missing-access-tech-type-opt-count" { return "MissingAccessTechTypeOptCount" }
    if yname == "no-author-for-ipv4-mobility-count" { return "NoAuthorForIpv4MobilityCount" }
    if yname == "no-author-for-ipv4-hoa-count" { return "NoAuthorForIpv4HoaCount" }
    if yname == "no-author-for-ipv6-mobility-count" { return "NoAuthorForIpv6MobilityCount" }
    if yname == "multiple-ipv4-ho-a-not-supported-count" { return "MultipleIpv4HoANotSupportedCount" }
    if yname == "gre-key-opt-required-count" { return "GreKeyOptRequiredCount" }
    return ""
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetSegmentPath() string {
    return "pba-send-statistics"
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pba-count"] = pbaSendStatistics.PbaCount
    leafs["pba-drop-count"] = pbaSendStatistics.PbaDropCount
    leafs["accepted-count"] = pbaSendStatistics.AcceptedCount
    leafs["unknown-count"] = pbaSendStatistics.UnknownCount
    leafs["unspecified-failure-count"] = pbaSendStatistics.UnspecifiedFailureCount
    leafs["admin-failure-count"] = pbaSendStatistics.AdminFailureCount
    leafs["resource-failure-count"] = pbaSendStatistics.ResourceFailureCount
    leafs["home-reg-failure-count"] = pbaSendStatistics.HomeRegFailureCount
    leafs["home-subnet-failure-count"] = pbaSendStatistics.HomeSubnetFailureCount
    leafs["bad-sequence-failure-count"] = pbaSendStatistics.BadSequenceFailureCount
    leafs["reg-type-failure-count"] = pbaSendStatistics.RegTypeFailureCount
    leafs["authen-failure-count"] = pbaSendStatistics.AuthenFailureCount
    leafs["proxy-reg-not-enabled-count"] = pbaSendStatistics.ProxyRegNotEnabledCount
    leafs["not-lma-for-this-mn-count"] = pbaSendStatistics.NotLmaForThisMnCount
    leafs["no-author-for-proxy-reg-count"] = pbaSendStatistics.NoAuthorForProxyRegCount
    leafs["no-author-for-hnp-count"] = pbaSendStatistics.NoAuthorForHnpCount
    leafs["timestamp-mismatch-count"] = pbaSendStatistics.TimestampMismatchCount
    leafs["timestamp-lower-than-previous-accepted-count"] = pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount
    leafs["missing-hnp-opt-count"] = pbaSendStatistics.MissingHnpOptCount
    leafs["received-hnps-do-not-match-bce-hnps-count"] = pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount
    leafs["missing-mn-id-opt-count"] = pbaSendStatistics.MissingMnIdOptCount
    leafs["missing-hi-opt-count"] = pbaSendStatistics.MissingHiOptCount
    leafs["missing-access-tech-type-opt-count"] = pbaSendStatistics.MissingAccessTechTypeOptCount
    leafs["no-author-for-ipv4-mobility-count"] = pbaSendStatistics.NoAuthorForIpv4MobilityCount
    leafs["no-author-for-ipv4-hoa-count"] = pbaSendStatistics.NoAuthorForIpv4HoaCount
    leafs["no-author-for-ipv6-mobility-count"] = pbaSendStatistics.NoAuthorForIpv6MobilityCount
    leafs["multiple-ipv4-ho-a-not-supported-count"] = pbaSendStatistics.MultipleIpv4HoANotSupportedCount
    leafs["gre-key-opt-required-count"] = pbaSendStatistics.GreKeyOptRequiredCount
    return leafs
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetYangName() string { return "pba-send-statistics" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) SetParent(parent types.Entity) { pbaSendStatistics.parent = parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetParent() types.Entity { return pbaSendStatistics.parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetFilter() yfilter.YFilter { return pbriSendStatistics.YFilter }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) SetFilter(yf yfilter.YFilter) { pbriSendStatistics.YFilter = yf }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetSegmentPath() string {
    return "pbri-send-statistics"
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriSendStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriSendStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriSendStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriSendStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriSendStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriSendStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriSendStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriSendStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriSendStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriSendStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriSendStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriSendStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetYangName() string { return "pbri-send-statistics" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) SetParent(parent types.Entity) { pbriSendStatistics.parent = parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetParent() types.Entity { return pbriSendStatistics.parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetFilter() yfilter.YFilter { return pbriReceiveStatistics.YFilter }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbriReceiveStatistics.YFilter = yf }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetSegmentPath() string {
    return "pbri-receive-statistics"
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriReceiveStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriReceiveStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriReceiveStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriReceiveStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriReceiveStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriReceiveStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriReceiveStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriReceiveStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriReceiveStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriReceiveStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriReceiveStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriReceiveStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetYangName() string { return "pbri-receive-statistics" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) SetParent(parent types.Entity) { pbriReceiveStatistics.parent = parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetParent() types.Entity { return pbriReceiveStatistics.parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetFilter() yfilter.YFilter { return pbraSendStatistics.YFilter }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) SetFilter(yf yfilter.YFilter) { pbraSendStatistics.YFilter = yf }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetSegmentPath() string {
    return "pbra-send-statistics"
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraSendStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraSendStatistics.PbraDropCount
    leafs["success-count"] = pbraSendStatistics.SuccessCount
    leafs["partial-success-count"] = pbraSendStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraSendStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraSendStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraSendStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraSendStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraSendStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraSendStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraSendStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraSendStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetYangName() string { return "pbra-send-statistics" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) SetParent(parent types.Entity) { pbraSendStatistics.parent = parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetParent() types.Entity { return pbraSendStatistics.parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetFilter() yfilter.YFilter { return pbraReceiveStatistics.YFilter }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbraReceiveStatistics.YFilter = yf }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetSegmentPath() string {
    return "pbra-receive-statistics"
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraReceiveStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraReceiveStatistics.PbraDropCount
    leafs["success-count"] = pbraReceiveStatistics.SuccessCount
    leafs["partial-success-count"] = pbraReceiveStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraReceiveStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraReceiveStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraReceiveStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraReceiveStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraReceiveStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraReceiveStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraReceiveStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraReceiveStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetYangName() string { return "pbra-receive-statistics" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) SetParent(parent types.Entity) { pbraReceiveStatistics.parent = parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetParent() types.Entity { return pbraReceiveStatistics.parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics
// LMA Accounting Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of Accounting Start Records Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AccountingStartSentCount interface{}

    // Count of Accounting Update Records Sent. The type is interface{} with
    // range: 0..18446744073709551615.
    AccountingUpdateSentCount interface{}

    // Count of Accounting Stop Records Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AccountingStopSentCount interface{}
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetFilter() yfilter.YFilter { return accountingStatistics.YFilter }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) SetFilter(yf yfilter.YFilter) { accountingStatistics.YFilter = yf }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetGoName(yname string) string {
    if yname == "accounting-start-sent-count" { return "AccountingStartSentCount" }
    if yname == "accounting-update-sent-count" { return "AccountingUpdateSentCount" }
    if yname == "accounting-stop-sent-count" { return "AccountingStopSentCount" }
    return ""
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetSegmentPath() string {
    return "accounting-statistics"
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accounting-start-sent-count"] = accountingStatistics.AccountingStartSentCount
    leafs["accounting-update-sent-count"] = accountingStatistics.AccountingUpdateSentCount
    leafs["accounting-stop-sent-count"] = accountingStatistics.AccountingStopSentCount
    return leafs
}

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetYangName() string { return "accounting-statistics" }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) SetParent(parent types.Entity) { accountingStatistics.parent = parent }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetParent() types.Entity { return accountingStatistics.parent }

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetParentYangName() string { return "customer-statistic" }

// Pmipv6_Lma_Statistics_License
// LMA License Statistics
type Pmipv6_Lma_Statistics_License struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LMA Identifier. The type is string.
    LmaIdentifier interface{}

    // Instantaneous Count of Bindings. The type is interface{} with range:
    // 0..4294967295.
    BceCount interface{}

    // Peak Count of Bindings. The type is interface{} with range: 0..4294967295.
    PeakBceCount interface{}

    // Timestamp when the Peak Count of Bindings was reset. The type is
    // interface{} with range: 0..4294967295.
    PeakBceCountResetTimestamp interface{}
}

func (license *Pmipv6_Lma_Statistics_License) GetFilter() yfilter.YFilter { return license.YFilter }

func (license *Pmipv6_Lma_Statistics_License) SetFilter(yf yfilter.YFilter) { license.YFilter = yf }

func (license *Pmipv6_Lma_Statistics_License) GetGoName(yname string) string {
    if yname == "lma-identifier" { return "LmaIdentifier" }
    if yname == "bce-count" { return "BceCount" }
    if yname == "peak-bce-count" { return "PeakBceCount" }
    if yname == "peak-bce-count-reset-timestamp" { return "PeakBceCountResetTimestamp" }
    return ""
}

func (license *Pmipv6_Lma_Statistics_License) GetSegmentPath() string {
    return "license"
}

func (license *Pmipv6_Lma_Statistics_License) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (license *Pmipv6_Lma_Statistics_License) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (license *Pmipv6_Lma_Statistics_License) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-identifier"] = license.LmaIdentifier
    leafs["bce-count"] = license.BceCount
    leafs["peak-bce-count"] = license.PeakBceCount
    leafs["peak-bce-count-reset-timestamp"] = license.PeakBceCountResetTimestamp
    return leafs
}

func (license *Pmipv6_Lma_Statistics_License) GetBundleName() string { return "cisco_ios_xr" }

func (license *Pmipv6_Lma_Statistics_License) GetYangName() string { return "license" }

func (license *Pmipv6_Lma_Statistics_License) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (license *Pmipv6_Lma_Statistics_License) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (license *Pmipv6_Lma_Statistics_License) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (license *Pmipv6_Lma_Statistics_License) SetParent(parent types.Entity) { license.parent = parent }

func (license *Pmipv6_Lma_Statistics_License) GetParent() types.Entity { return license.parent }

func (license *Pmipv6_Lma_Statistics_License) GetParentYangName() string { return "statistics" }

// Pmipv6_Lma_Statistics_Global
// Global Statistics
type Pmipv6_Lma_Statistics_Global struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LMA Identifier. The type is string.
    LmaIdentifier interface{}

    // Count of Bindings. The type is interface{} with range: 0..4294967295.
    BceCount interface{}

    // Count of Handoffs. The type is interface{} with range: 0..4294967295.
    HandoffCount interface{}

    // Count of Single Tenants. The type is interface{} with range: 0..4294967295.
    SingleTenantCount interface{}

    // Count of Multi Tenants. The type is interface{} with range: 0..4294967295.
    MultiTenantCount interface{}

    // Packet Statistics.
    PacketStatistics Pmipv6_Lma_Statistics_Global_PacketStatistics

    // LMA Protocol Statistics.
    ProtocolStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics

    // LMA Accounting Statistics.
    AccountingStatistics Pmipv6_Lma_Statistics_Global_AccountingStatistics
}

func (global *Pmipv6_Lma_Statistics_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Pmipv6_Lma_Statistics_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Pmipv6_Lma_Statistics_Global) GetGoName(yname string) string {
    if yname == "lma-identifier" { return "LmaIdentifier" }
    if yname == "bce-count" { return "BceCount" }
    if yname == "handoff-count" { return "HandoffCount" }
    if yname == "single-tenant-count" { return "SingleTenantCount" }
    if yname == "multi-tenant-count" { return "MultiTenantCount" }
    if yname == "packet-statistics" { return "PacketStatistics" }
    if yname == "protocol-statistics" { return "ProtocolStatistics" }
    if yname == "accounting-statistics" { return "AccountingStatistics" }
    return ""
}

func (global *Pmipv6_Lma_Statistics_Global) GetSegmentPath() string {
    return "global"
}

func (global *Pmipv6_Lma_Statistics_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-statistics" {
        return &global.PacketStatistics
    }
    if childYangName == "protocol-statistics" {
        return &global.ProtocolStatistics
    }
    if childYangName == "accounting-statistics" {
        return &global.AccountingStatistics
    }
    return nil
}

func (global *Pmipv6_Lma_Statistics_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-statistics"] = &global.PacketStatistics
    children["protocol-statistics"] = &global.ProtocolStatistics
    children["accounting-statistics"] = &global.AccountingStatistics
    return children
}

func (global *Pmipv6_Lma_Statistics_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lma-identifier"] = global.LmaIdentifier
    leafs["bce-count"] = global.BceCount
    leafs["handoff-count"] = global.HandoffCount
    leafs["single-tenant-count"] = global.SingleTenantCount
    leafs["multi-tenant-count"] = global.MultiTenantCount
    return leafs
}

func (global *Pmipv6_Lma_Statistics_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Pmipv6_Lma_Statistics_Global) GetYangName() string { return "global" }

func (global *Pmipv6_Lma_Statistics_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Pmipv6_Lma_Statistics_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Pmipv6_Lma_Statistics_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Pmipv6_Lma_Statistics_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Pmipv6_Lma_Statistics_Global) GetParent() types.Entity { return global.parent }

func (global *Pmipv6_Lma_Statistics_Global) GetParentYangName() string { return "statistics" }

// Pmipv6_Lma_Statistics_Global_PacketStatistics
// Packet Statistics
type Pmipv6_Lma_Statistics_Global_PacketStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Checksumm errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ChecksumErrors interface{}

    // Drop count of sent packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SendDrops interface{}

    // Drop count of received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveDrops interface{}

    // Count of received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Count of sent packets. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsSent interface{}

    // Drop count of IPv6 sent packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SendDropsIpv6 interface{}

    // Drop count of IPv6 received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveDropsIpv6 interface{}

    // Count of IPv6 received packets. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceivedIpv6 interface{}

    // Count of IPv6 sent packets. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsSentIpv6 interface{}
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetFilter() yfilter.YFilter { return packetStatistics.YFilter }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) SetFilter(yf yfilter.YFilter) { packetStatistics.YFilter = yf }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetGoName(yname string) string {
    if yname == "checksum-errors" { return "ChecksumErrors" }
    if yname == "send-drops" { return "SendDrops" }
    if yname == "receive-drops" { return "ReceiveDrops" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "send-drops-ipv6" { return "SendDropsIpv6" }
    if yname == "receive-drops-ipv6" { return "ReceiveDropsIpv6" }
    if yname == "packets-received-ipv6" { return "PacketsReceivedIpv6" }
    if yname == "packets-sent-ipv6" { return "PacketsSentIpv6" }
    return ""
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetSegmentPath() string {
    return "packet-statistics"
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["checksum-errors"] = packetStatistics.ChecksumErrors
    leafs["send-drops"] = packetStatistics.SendDrops
    leafs["receive-drops"] = packetStatistics.ReceiveDrops
    leafs["packets-received"] = packetStatistics.PacketsReceived
    leafs["packets-sent"] = packetStatistics.PacketsSent
    leafs["send-drops-ipv6"] = packetStatistics.SendDropsIpv6
    leafs["receive-drops-ipv6"] = packetStatistics.ReceiveDropsIpv6
    leafs["packets-received-ipv6"] = packetStatistics.PacketsReceivedIpv6
    leafs["packets-sent-ipv6"] = packetStatistics.PacketsSentIpv6
    return leafs
}

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetYangName() string { return "packet-statistics" }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) SetParent(parent types.Entity) { packetStatistics.parent = parent }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetParent() types.Entity { return packetStatistics.parent }

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetParentYangName() string { return "global" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBU Receive Statistics.
    PbuReceiveStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics

    // PBA Send Statistics.
    PbaSendStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics

    // PBRI Send Statistics.
    PbriSendStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics

    // PBRI Receive Statistics.
    PbriReceiveStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics

    // PBRA Send Statistics.
    PbraSendStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics

    // PBRA Receive Statistics.
    PbraReceiveStatistics Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetFilter() yfilter.YFilter { return protocolStatistics.YFilter }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) SetFilter(yf yfilter.YFilter) { protocolStatistics.YFilter = yf }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetGoName(yname string) string {
    if yname == "pbu-receive-statistics" { return "PbuReceiveStatistics" }
    if yname == "pba-send-statistics" { return "PbaSendStatistics" }
    if yname == "pbri-send-statistics" { return "PbriSendStatistics" }
    if yname == "pbri-receive-statistics" { return "PbriReceiveStatistics" }
    if yname == "pbra-send-statistics" { return "PbraSendStatistics" }
    if yname == "pbra-receive-statistics" { return "PbraReceiveStatistics" }
    return ""
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetSegmentPath() string {
    return "protocol-statistics"
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbu-receive-statistics" {
        return &protocolStatistics.PbuReceiveStatistics
    }
    if childYangName == "pba-send-statistics" {
        return &protocolStatistics.PbaSendStatistics
    }
    if childYangName == "pbri-send-statistics" {
        return &protocolStatistics.PbriSendStatistics
    }
    if childYangName == "pbri-receive-statistics" {
        return &protocolStatistics.PbriReceiveStatistics
    }
    if childYangName == "pbra-send-statistics" {
        return &protocolStatistics.PbraSendStatistics
    }
    if childYangName == "pbra-receive-statistics" {
        return &protocolStatistics.PbraReceiveStatistics
    }
    return nil
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbu-receive-statistics"] = &protocolStatistics.PbuReceiveStatistics
    children["pba-send-statistics"] = &protocolStatistics.PbaSendStatistics
    children["pbri-send-statistics"] = &protocolStatistics.PbriSendStatistics
    children["pbri-receive-statistics"] = &protocolStatistics.PbriReceiveStatistics
    children["pbra-send-statistics"] = &protocolStatistics.PbraSendStatistics
    children["pbra-receive-statistics"] = &protocolStatistics.PbraReceiveStatistics
    return children
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetYangName() string { return "protocol-statistics" }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) SetParent(parent types.Entity) { protocolStatistics.parent = parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetParent() types.Entity { return protocolStatistics.parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetParentYangName() string { return "global" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetFilter() yfilter.YFilter { return pbuReceiveStatistics.YFilter }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbuReceiveStatistics.YFilter = yf }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbu-count" { return "PbuCount" }
    if yname == "pbu-drop-count" { return "PbuDropCount" }
    return ""
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetSegmentPath() string {
    return "pbu-receive-statistics"
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbu-count"] = pbuReceiveStatistics.PbuCount
    leafs["pbu-drop-count"] = pbuReceiveStatistics.PbuDropCount
    return leafs
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetYangName() string { return "pbu-receive-statistics" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) SetParent(parent types.Entity) { pbuReceiveStatistics.parent = parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetParent() types.Entity { return pbuReceiveStatistics.parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBAs. The type is interface{} with range: 0..18446744073709551615.
    PbaCount interface{}

    // Count of PBAs dropped. The type is interface{} with range: 0..4294967295.
    PbaDropCount interface{}

    // Count of Status Code - Binding Update accepted. The type is interface{}
    // with range: 0..4294967295.
    AcceptedCount interface{}

    // Count of Status Code - Last BA status code sent. The type is interface{}
    // with range: 0..4294967295.
    UnknownCount interface{}

    // Count of Status Code - Reason unspecified. The type is interface{} with
    // range: 0..4294967295.
    UnspecifiedFailureCount interface{}

    // Count of Status Code - Administratively prohibited. The type is interface{}
    // with range: 0..4294967295.
    AdminFailureCount interface{}

    // Count of Status Code - Insufficient resources. The type is interface{} with
    // range: 0..4294967295.
    ResourceFailureCount interface{}

    // Count of Status Code - Home registration not supported. The type is
    // interface{} with range: 0..4294967295.
    HomeRegFailureCount interface{}

    // Count of Status Code - Not home subnet. The type is interface{} with range:
    // 0..4294967295.
    HomeSubnetFailureCount interface{}

    // Count of Status Code - Sequence number out of window. The type is
    // interface{} with range: 0..4294967295.
    BadSequenceFailureCount interface{}

    // Count of Status Code - Registration type change. The type is interface{}
    // with range: 0..4294967295.
    RegTypeFailureCount interface{}

    // Count of Status Code - Auth Fail. The type is interface{} with range:
    // 0..4294967295.
    AuthenFailureCount interface{}

    // Count of Status Code - Proxy Registration not enabled. The type is
    // interface{} with range: 0..4294967295.
    ProxyRegNotEnabledCount interface{}

    // Count of Status Code - Not LMA for this Mobile Node. The type is
    // interface{} with range: 0..4294967295.
    NotLmaForThisMnCount interface{}

    // Count of Status Code - MAG not auth.for proxyreg. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForProxyRegCount interface{}

    // Count of Status Code - Not authorized for HNP. The type is interface{} with
    // range: 0..4294967295.
    NoAuthorForHnpCount interface{}

    // Count of Status Code - Invalid timestamp value. The type is interface{}
    // with range: 0..4294967295.
    TimestampMismatchCount interface{}

    // Count of Status Code - Timestamp lower than previous accepted. The type is
    // interface{} with range: 0..4294967295.
    TimestampLowerThanPreviousAcceptedCount interface{}

    // Count of Status Code - Missing Home Network Prefix option. The type is
    // interface{} with range: 0..4294967295.
    MissingHnpOptCount interface{}

    // Count of Status Code - Recevied HNPs do not match with BCE. The type is
    // interface{} with range: 0..4294967295.
    ReceivedHnpsDoNotMatchBceHnpsCount interface{}

    // Count of Status Code - Missing MN identifier option. The type is
    // interface{} with range: 0..4294967295.
    MissingMnIdOptCount interface{}

    // Count of Status Code - Missing Handoff Indicator. The type is interface{}
    // with range: 0..4294967295.
    MissingHiOptCount interface{}

    // Count of Status Code - Missing ATT option. The type is interface{} with
    // range: 0..4294967295.
    MissingAccessTechTypeOptCount interface{}

    // Count of Status Code - Not authorized for IPv4 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv4MobilityCount interface{}

    // Count of Status Code - Not authorized for IPv4 HoA. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForIpv4HoaCount interface{}

    // Count of Status Code - Not authorized for IPv6 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv6MobilityCount interface{}

    // Count of Status Code - Multiple IPv4 HoA not supported. The type is
    // interface{} with range: 0..4294967295.
    MultipleIpv4HoANotSupportedCount interface{}

    // Count of Status Code - GRE Key option is required. The type is interface{}
    // with range: 0..4294967295.
    GreKeyOptRequiredCount interface{}
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetFilter() yfilter.YFilter { return pbaSendStatistics.YFilter }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) SetFilter(yf yfilter.YFilter) { pbaSendStatistics.YFilter = yf }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetGoName(yname string) string {
    if yname == "pba-count" { return "PbaCount" }
    if yname == "pba-drop-count" { return "PbaDropCount" }
    if yname == "accepted-count" { return "AcceptedCount" }
    if yname == "unknown-count" { return "UnknownCount" }
    if yname == "unspecified-failure-count" { return "UnspecifiedFailureCount" }
    if yname == "admin-failure-count" { return "AdminFailureCount" }
    if yname == "resource-failure-count" { return "ResourceFailureCount" }
    if yname == "home-reg-failure-count" { return "HomeRegFailureCount" }
    if yname == "home-subnet-failure-count" { return "HomeSubnetFailureCount" }
    if yname == "bad-sequence-failure-count" { return "BadSequenceFailureCount" }
    if yname == "reg-type-failure-count" { return "RegTypeFailureCount" }
    if yname == "authen-failure-count" { return "AuthenFailureCount" }
    if yname == "proxy-reg-not-enabled-count" { return "ProxyRegNotEnabledCount" }
    if yname == "not-lma-for-this-mn-count" { return "NotLmaForThisMnCount" }
    if yname == "no-author-for-proxy-reg-count" { return "NoAuthorForProxyRegCount" }
    if yname == "no-author-for-hnp-count" { return "NoAuthorForHnpCount" }
    if yname == "timestamp-mismatch-count" { return "TimestampMismatchCount" }
    if yname == "timestamp-lower-than-previous-accepted-count" { return "TimestampLowerThanPreviousAcceptedCount" }
    if yname == "missing-hnp-opt-count" { return "MissingHnpOptCount" }
    if yname == "received-hnps-do-not-match-bce-hnps-count" { return "ReceivedHnpsDoNotMatchBceHnpsCount" }
    if yname == "missing-mn-id-opt-count" { return "MissingMnIdOptCount" }
    if yname == "missing-hi-opt-count" { return "MissingHiOptCount" }
    if yname == "missing-access-tech-type-opt-count" { return "MissingAccessTechTypeOptCount" }
    if yname == "no-author-for-ipv4-mobility-count" { return "NoAuthorForIpv4MobilityCount" }
    if yname == "no-author-for-ipv4-hoa-count" { return "NoAuthorForIpv4HoaCount" }
    if yname == "no-author-for-ipv6-mobility-count" { return "NoAuthorForIpv6MobilityCount" }
    if yname == "multiple-ipv4-ho-a-not-supported-count" { return "MultipleIpv4HoANotSupportedCount" }
    if yname == "gre-key-opt-required-count" { return "GreKeyOptRequiredCount" }
    return ""
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetSegmentPath() string {
    return "pba-send-statistics"
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pba-count"] = pbaSendStatistics.PbaCount
    leafs["pba-drop-count"] = pbaSendStatistics.PbaDropCount
    leafs["accepted-count"] = pbaSendStatistics.AcceptedCount
    leafs["unknown-count"] = pbaSendStatistics.UnknownCount
    leafs["unspecified-failure-count"] = pbaSendStatistics.UnspecifiedFailureCount
    leafs["admin-failure-count"] = pbaSendStatistics.AdminFailureCount
    leafs["resource-failure-count"] = pbaSendStatistics.ResourceFailureCount
    leafs["home-reg-failure-count"] = pbaSendStatistics.HomeRegFailureCount
    leafs["home-subnet-failure-count"] = pbaSendStatistics.HomeSubnetFailureCount
    leafs["bad-sequence-failure-count"] = pbaSendStatistics.BadSequenceFailureCount
    leafs["reg-type-failure-count"] = pbaSendStatistics.RegTypeFailureCount
    leafs["authen-failure-count"] = pbaSendStatistics.AuthenFailureCount
    leafs["proxy-reg-not-enabled-count"] = pbaSendStatistics.ProxyRegNotEnabledCount
    leafs["not-lma-for-this-mn-count"] = pbaSendStatistics.NotLmaForThisMnCount
    leafs["no-author-for-proxy-reg-count"] = pbaSendStatistics.NoAuthorForProxyRegCount
    leafs["no-author-for-hnp-count"] = pbaSendStatistics.NoAuthorForHnpCount
    leafs["timestamp-mismatch-count"] = pbaSendStatistics.TimestampMismatchCount
    leafs["timestamp-lower-than-previous-accepted-count"] = pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount
    leafs["missing-hnp-opt-count"] = pbaSendStatistics.MissingHnpOptCount
    leafs["received-hnps-do-not-match-bce-hnps-count"] = pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount
    leafs["missing-mn-id-opt-count"] = pbaSendStatistics.MissingMnIdOptCount
    leafs["missing-hi-opt-count"] = pbaSendStatistics.MissingHiOptCount
    leafs["missing-access-tech-type-opt-count"] = pbaSendStatistics.MissingAccessTechTypeOptCount
    leafs["no-author-for-ipv4-mobility-count"] = pbaSendStatistics.NoAuthorForIpv4MobilityCount
    leafs["no-author-for-ipv4-hoa-count"] = pbaSendStatistics.NoAuthorForIpv4HoaCount
    leafs["no-author-for-ipv6-mobility-count"] = pbaSendStatistics.NoAuthorForIpv6MobilityCount
    leafs["multiple-ipv4-ho-a-not-supported-count"] = pbaSendStatistics.MultipleIpv4HoANotSupportedCount
    leafs["gre-key-opt-required-count"] = pbaSendStatistics.GreKeyOptRequiredCount
    return leafs
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetYangName() string { return "pba-send-statistics" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) SetParent(parent types.Entity) { pbaSendStatistics.parent = parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetParent() types.Entity { return pbaSendStatistics.parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetFilter() yfilter.YFilter { return pbriSendStatistics.YFilter }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) SetFilter(yf yfilter.YFilter) { pbriSendStatistics.YFilter = yf }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetSegmentPath() string {
    return "pbri-send-statistics"
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriSendStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriSendStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriSendStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriSendStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriSendStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriSendStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriSendStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriSendStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriSendStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriSendStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriSendStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriSendStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetYangName() string { return "pbri-send-statistics" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) SetParent(parent types.Entity) { pbriSendStatistics.parent = parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetParent() types.Entity { return pbriSendStatistics.parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetFilter() yfilter.YFilter { return pbriReceiveStatistics.YFilter }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbriReceiveStatistics.YFilter = yf }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetSegmentPath() string {
    return "pbri-receive-statistics"
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriReceiveStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriReceiveStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriReceiveStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriReceiveStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriReceiveStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriReceiveStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriReceiveStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriReceiveStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriReceiveStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriReceiveStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriReceiveStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriReceiveStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetYangName() string { return "pbri-receive-statistics" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) SetParent(parent types.Entity) { pbriReceiveStatistics.parent = parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetParent() types.Entity { return pbriReceiveStatistics.parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetFilter() yfilter.YFilter { return pbraSendStatistics.YFilter }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) SetFilter(yf yfilter.YFilter) { pbraSendStatistics.YFilter = yf }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetSegmentPath() string {
    return "pbra-send-statistics"
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraSendStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraSendStatistics.PbraDropCount
    leafs["success-count"] = pbraSendStatistics.SuccessCount
    leafs["partial-success-count"] = pbraSendStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraSendStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraSendStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraSendStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraSendStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraSendStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraSendStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraSendStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraSendStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetYangName() string { return "pbra-send-statistics" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) SetParent(parent types.Entity) { pbraSendStatistics.parent = parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetParent() types.Entity { return pbraSendStatistics.parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetFilter() yfilter.YFilter { return pbraReceiveStatistics.YFilter }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbraReceiveStatistics.YFilter = yf }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetSegmentPath() string {
    return "pbra-receive-statistics"
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraReceiveStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraReceiveStatistics.PbraDropCount
    leafs["success-count"] = pbraReceiveStatistics.SuccessCount
    leafs["partial-success-count"] = pbraReceiveStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraReceiveStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraReceiveStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraReceiveStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraReceiveStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraReceiveStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraReceiveStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraReceiveStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraReceiveStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetYangName() string { return "pbra-receive-statistics" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) SetParent(parent types.Entity) { pbraReceiveStatistics.parent = parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetParent() types.Entity { return pbraReceiveStatistics.parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_Global_AccountingStatistics
// LMA Accounting Statistics
type Pmipv6_Lma_Statistics_Global_AccountingStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of Accounting Start Records Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AccountingStartSentCount interface{}

    // Count of Accounting Update Records Sent. The type is interface{} with
    // range: 0..18446744073709551615.
    AccountingUpdateSentCount interface{}

    // Count of Accounting Stop Records Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AccountingStopSentCount interface{}
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetFilter() yfilter.YFilter { return accountingStatistics.YFilter }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) SetFilter(yf yfilter.YFilter) { accountingStatistics.YFilter = yf }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetGoName(yname string) string {
    if yname == "accounting-start-sent-count" { return "AccountingStartSentCount" }
    if yname == "accounting-update-sent-count" { return "AccountingUpdateSentCount" }
    if yname == "accounting-stop-sent-count" { return "AccountingStopSentCount" }
    return ""
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetSegmentPath() string {
    return "accounting-statistics"
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["accounting-start-sent-count"] = accountingStatistics.AccountingStartSentCount
    leafs["accounting-update-sent-count"] = accountingStatistics.AccountingUpdateSentCount
    leafs["accounting-stop-sent-count"] = accountingStatistics.AccountingStopSentCount
    return leafs
}

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetYangName() string { return "accounting-statistics" }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) SetParent(parent types.Entity) { accountingStatistics.parent = parent }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetParent() types.Entity { return accountingStatistics.parent }

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetParentYangName() string { return "global" }

// Pmipv6_Lma_Statistics_MagStatistics
// Table of MAGStatistics
type Pmipv6_Lma_Statistics_MagStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer MAG statistics. The type is slice of
    // Pmipv6_Lma_Statistics_MagStatistics_MagStatistic.
    MagStatistic []Pmipv6_Lma_Statistics_MagStatistics_MagStatistic
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetFilter() yfilter.YFilter { return magStatistics.YFilter }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) SetFilter(yf yfilter.YFilter) { magStatistics.YFilter = yf }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetGoName(yname string) string {
    if yname == "mag-statistic" { return "MagStatistic" }
    return ""
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetSegmentPath() string {
    return "mag-statistics"
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mag-statistic" {
        for _, c := range magStatistics.MagStatistic {
            if magStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Statistics_MagStatistics_MagStatistic{}
        magStatistics.MagStatistic = append(magStatistics.MagStatistic, child)
        return &magStatistics.MagStatistic[len(magStatistics.MagStatistic)-1]
    }
    return nil
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range magStatistics.MagStatistic {
        children[magStatistics.MagStatistic[i].GetSegmentPath()] = &magStatistics.MagStatistic[i]
    }
    return children
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetYangName() string { return "mag-statistics" }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) SetParent(parent types.Entity) { magStatistics.parent = parent }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetParent() types.Entity { return magStatistics.parent }

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetParentYangName() string { return "statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic
// Peer MAG statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer MAG Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MagName interface{}

    // LMA Identifier. The type is string.
    LmaIdentifier interface{}

    // LMA Protocol Statistics.
    ProtocolStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetFilter() yfilter.YFilter { return magStatistic.YFilter }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) SetFilter(yf yfilter.YFilter) { magStatistic.YFilter = yf }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetGoName(yname string) string {
    if yname == "mag-name" { return "MagName" }
    if yname == "lma-identifier" { return "LmaIdentifier" }
    if yname == "protocol-statistics" { return "ProtocolStatistics" }
    return ""
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetSegmentPath() string {
    return "mag-statistic" + "[mag-name='" + fmt.Sprintf("%v", magStatistic.MagName) + "']"
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "protocol-statistics" {
        return &magStatistic.ProtocolStatistics
    }
    return nil
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["protocol-statistics"] = &magStatistic.ProtocolStatistics
    return children
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mag-name"] = magStatistic.MagName
    leafs["lma-identifier"] = magStatistic.LmaIdentifier
    return leafs
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetYangName() string { return "mag-statistic" }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) SetParent(parent types.Entity) { magStatistic.parent = parent }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetParent() types.Entity { return magStatistic.parent }

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetParentYangName() string { return "mag-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBU Receive Statistics.
    PbuReceiveStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics

    // PBA Send Statistics.
    PbaSendStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics

    // PBRI Send Statistics.
    PbriSendStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics

    // PBRI Receive Statistics.
    PbriReceiveStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics

    // PBRA Send Statistics.
    PbraSendStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics

    // PBRA Receive Statistics.
    PbraReceiveStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetFilter() yfilter.YFilter { return protocolStatistics.YFilter }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) SetFilter(yf yfilter.YFilter) { protocolStatistics.YFilter = yf }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetGoName(yname string) string {
    if yname == "pbu-receive-statistics" { return "PbuReceiveStatistics" }
    if yname == "pba-send-statistics" { return "PbaSendStatistics" }
    if yname == "pbri-send-statistics" { return "PbriSendStatistics" }
    if yname == "pbri-receive-statistics" { return "PbriReceiveStatistics" }
    if yname == "pbra-send-statistics" { return "PbraSendStatistics" }
    if yname == "pbra-receive-statistics" { return "PbraReceiveStatistics" }
    return ""
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetSegmentPath() string {
    return "protocol-statistics"
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbu-receive-statistics" {
        return &protocolStatistics.PbuReceiveStatistics
    }
    if childYangName == "pba-send-statistics" {
        return &protocolStatistics.PbaSendStatistics
    }
    if childYangName == "pbri-send-statistics" {
        return &protocolStatistics.PbriSendStatistics
    }
    if childYangName == "pbri-receive-statistics" {
        return &protocolStatistics.PbriReceiveStatistics
    }
    if childYangName == "pbra-send-statistics" {
        return &protocolStatistics.PbraSendStatistics
    }
    if childYangName == "pbra-receive-statistics" {
        return &protocolStatistics.PbraReceiveStatistics
    }
    return nil
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbu-receive-statistics"] = &protocolStatistics.PbuReceiveStatistics
    children["pba-send-statistics"] = &protocolStatistics.PbaSendStatistics
    children["pbri-send-statistics"] = &protocolStatistics.PbriSendStatistics
    children["pbri-receive-statistics"] = &protocolStatistics.PbriReceiveStatistics
    children["pbra-send-statistics"] = &protocolStatistics.PbraSendStatistics
    children["pbra-receive-statistics"] = &protocolStatistics.PbraReceiveStatistics
    return children
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetYangName() string { return "protocol-statistics" }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) SetParent(parent types.Entity) { protocolStatistics.parent = parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetParent() types.Entity { return protocolStatistics.parent }

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetParentYangName() string { return "mag-statistic" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetFilter() yfilter.YFilter { return pbuReceiveStatistics.YFilter }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbuReceiveStatistics.YFilter = yf }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbu-count" { return "PbuCount" }
    if yname == "pbu-drop-count" { return "PbuDropCount" }
    return ""
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetSegmentPath() string {
    return "pbu-receive-statistics"
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbu-count"] = pbuReceiveStatistics.PbuCount
    leafs["pbu-drop-count"] = pbuReceiveStatistics.PbuDropCount
    return leafs
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetYangName() string { return "pbu-receive-statistics" }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) SetParent(parent types.Entity) { pbuReceiveStatistics.parent = parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetParent() types.Entity { return pbuReceiveStatistics.parent }

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBAs. The type is interface{} with range: 0..18446744073709551615.
    PbaCount interface{}

    // Count of PBAs dropped. The type is interface{} with range: 0..4294967295.
    PbaDropCount interface{}

    // Count of Status Code - Binding Update accepted. The type is interface{}
    // with range: 0..4294967295.
    AcceptedCount interface{}

    // Count of Status Code - Last BA status code sent. The type is interface{}
    // with range: 0..4294967295.
    UnknownCount interface{}

    // Count of Status Code - Reason unspecified. The type is interface{} with
    // range: 0..4294967295.
    UnspecifiedFailureCount interface{}

    // Count of Status Code - Administratively prohibited. The type is interface{}
    // with range: 0..4294967295.
    AdminFailureCount interface{}

    // Count of Status Code - Insufficient resources. The type is interface{} with
    // range: 0..4294967295.
    ResourceFailureCount interface{}

    // Count of Status Code - Home registration not supported. The type is
    // interface{} with range: 0..4294967295.
    HomeRegFailureCount interface{}

    // Count of Status Code - Not home subnet. The type is interface{} with range:
    // 0..4294967295.
    HomeSubnetFailureCount interface{}

    // Count of Status Code - Sequence number out of window. The type is
    // interface{} with range: 0..4294967295.
    BadSequenceFailureCount interface{}

    // Count of Status Code - Registration type change. The type is interface{}
    // with range: 0..4294967295.
    RegTypeFailureCount interface{}

    // Count of Status Code - Auth Fail. The type is interface{} with range:
    // 0..4294967295.
    AuthenFailureCount interface{}

    // Count of Status Code - Proxy Registration not enabled. The type is
    // interface{} with range: 0..4294967295.
    ProxyRegNotEnabledCount interface{}

    // Count of Status Code - Not LMA for this Mobile Node. The type is
    // interface{} with range: 0..4294967295.
    NotLmaForThisMnCount interface{}

    // Count of Status Code - MAG not auth.for proxyreg. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForProxyRegCount interface{}

    // Count of Status Code - Not authorized for HNP. The type is interface{} with
    // range: 0..4294967295.
    NoAuthorForHnpCount interface{}

    // Count of Status Code - Invalid timestamp value. The type is interface{}
    // with range: 0..4294967295.
    TimestampMismatchCount interface{}

    // Count of Status Code - Timestamp lower than previous accepted. The type is
    // interface{} with range: 0..4294967295.
    TimestampLowerThanPreviousAcceptedCount interface{}

    // Count of Status Code - Missing Home Network Prefix option. The type is
    // interface{} with range: 0..4294967295.
    MissingHnpOptCount interface{}

    // Count of Status Code - Recevied HNPs do not match with BCE. The type is
    // interface{} with range: 0..4294967295.
    ReceivedHnpsDoNotMatchBceHnpsCount interface{}

    // Count of Status Code - Missing MN identifier option. The type is
    // interface{} with range: 0..4294967295.
    MissingMnIdOptCount interface{}

    // Count of Status Code - Missing Handoff Indicator. The type is interface{}
    // with range: 0..4294967295.
    MissingHiOptCount interface{}

    // Count of Status Code - Missing ATT option. The type is interface{} with
    // range: 0..4294967295.
    MissingAccessTechTypeOptCount interface{}

    // Count of Status Code - Not authorized for IPv4 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv4MobilityCount interface{}

    // Count of Status Code - Not authorized for IPv4 HoA. The type is interface{}
    // with range: 0..4294967295.
    NoAuthorForIpv4HoaCount interface{}

    // Count of Status Code - Not authorized for IPv6 mobility. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForIpv6MobilityCount interface{}

    // Count of Status Code - Multiple IPv4 HoA not supported. The type is
    // interface{} with range: 0..4294967295.
    MultipleIpv4HoANotSupportedCount interface{}

    // Count of Status Code - GRE Key option is required. The type is interface{}
    // with range: 0..4294967295.
    GreKeyOptRequiredCount interface{}
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetFilter() yfilter.YFilter { return pbaSendStatistics.YFilter }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) SetFilter(yf yfilter.YFilter) { pbaSendStatistics.YFilter = yf }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetGoName(yname string) string {
    if yname == "pba-count" { return "PbaCount" }
    if yname == "pba-drop-count" { return "PbaDropCount" }
    if yname == "accepted-count" { return "AcceptedCount" }
    if yname == "unknown-count" { return "UnknownCount" }
    if yname == "unspecified-failure-count" { return "UnspecifiedFailureCount" }
    if yname == "admin-failure-count" { return "AdminFailureCount" }
    if yname == "resource-failure-count" { return "ResourceFailureCount" }
    if yname == "home-reg-failure-count" { return "HomeRegFailureCount" }
    if yname == "home-subnet-failure-count" { return "HomeSubnetFailureCount" }
    if yname == "bad-sequence-failure-count" { return "BadSequenceFailureCount" }
    if yname == "reg-type-failure-count" { return "RegTypeFailureCount" }
    if yname == "authen-failure-count" { return "AuthenFailureCount" }
    if yname == "proxy-reg-not-enabled-count" { return "ProxyRegNotEnabledCount" }
    if yname == "not-lma-for-this-mn-count" { return "NotLmaForThisMnCount" }
    if yname == "no-author-for-proxy-reg-count" { return "NoAuthorForProxyRegCount" }
    if yname == "no-author-for-hnp-count" { return "NoAuthorForHnpCount" }
    if yname == "timestamp-mismatch-count" { return "TimestampMismatchCount" }
    if yname == "timestamp-lower-than-previous-accepted-count" { return "TimestampLowerThanPreviousAcceptedCount" }
    if yname == "missing-hnp-opt-count" { return "MissingHnpOptCount" }
    if yname == "received-hnps-do-not-match-bce-hnps-count" { return "ReceivedHnpsDoNotMatchBceHnpsCount" }
    if yname == "missing-mn-id-opt-count" { return "MissingMnIdOptCount" }
    if yname == "missing-hi-opt-count" { return "MissingHiOptCount" }
    if yname == "missing-access-tech-type-opt-count" { return "MissingAccessTechTypeOptCount" }
    if yname == "no-author-for-ipv4-mobility-count" { return "NoAuthorForIpv4MobilityCount" }
    if yname == "no-author-for-ipv4-hoa-count" { return "NoAuthorForIpv4HoaCount" }
    if yname == "no-author-for-ipv6-mobility-count" { return "NoAuthorForIpv6MobilityCount" }
    if yname == "multiple-ipv4-ho-a-not-supported-count" { return "MultipleIpv4HoANotSupportedCount" }
    if yname == "gre-key-opt-required-count" { return "GreKeyOptRequiredCount" }
    return ""
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetSegmentPath() string {
    return "pba-send-statistics"
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pba-count"] = pbaSendStatistics.PbaCount
    leafs["pba-drop-count"] = pbaSendStatistics.PbaDropCount
    leafs["accepted-count"] = pbaSendStatistics.AcceptedCount
    leafs["unknown-count"] = pbaSendStatistics.UnknownCount
    leafs["unspecified-failure-count"] = pbaSendStatistics.UnspecifiedFailureCount
    leafs["admin-failure-count"] = pbaSendStatistics.AdminFailureCount
    leafs["resource-failure-count"] = pbaSendStatistics.ResourceFailureCount
    leafs["home-reg-failure-count"] = pbaSendStatistics.HomeRegFailureCount
    leafs["home-subnet-failure-count"] = pbaSendStatistics.HomeSubnetFailureCount
    leafs["bad-sequence-failure-count"] = pbaSendStatistics.BadSequenceFailureCount
    leafs["reg-type-failure-count"] = pbaSendStatistics.RegTypeFailureCount
    leafs["authen-failure-count"] = pbaSendStatistics.AuthenFailureCount
    leafs["proxy-reg-not-enabled-count"] = pbaSendStatistics.ProxyRegNotEnabledCount
    leafs["not-lma-for-this-mn-count"] = pbaSendStatistics.NotLmaForThisMnCount
    leafs["no-author-for-proxy-reg-count"] = pbaSendStatistics.NoAuthorForProxyRegCount
    leafs["no-author-for-hnp-count"] = pbaSendStatistics.NoAuthorForHnpCount
    leafs["timestamp-mismatch-count"] = pbaSendStatistics.TimestampMismatchCount
    leafs["timestamp-lower-than-previous-accepted-count"] = pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount
    leafs["missing-hnp-opt-count"] = pbaSendStatistics.MissingHnpOptCount
    leafs["received-hnps-do-not-match-bce-hnps-count"] = pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount
    leafs["missing-mn-id-opt-count"] = pbaSendStatistics.MissingMnIdOptCount
    leafs["missing-hi-opt-count"] = pbaSendStatistics.MissingHiOptCount
    leafs["missing-access-tech-type-opt-count"] = pbaSendStatistics.MissingAccessTechTypeOptCount
    leafs["no-author-for-ipv4-mobility-count"] = pbaSendStatistics.NoAuthorForIpv4MobilityCount
    leafs["no-author-for-ipv4-hoa-count"] = pbaSendStatistics.NoAuthorForIpv4HoaCount
    leafs["no-author-for-ipv6-mobility-count"] = pbaSendStatistics.NoAuthorForIpv6MobilityCount
    leafs["multiple-ipv4-ho-a-not-supported-count"] = pbaSendStatistics.MultipleIpv4HoANotSupportedCount
    leafs["gre-key-opt-required-count"] = pbaSendStatistics.GreKeyOptRequiredCount
    return leafs
}

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetYangName() string { return "pba-send-statistics" }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) SetParent(parent types.Entity) { pbaSendStatistics.parent = parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetParent() types.Entity { return pbaSendStatistics.parent }

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetFilter() yfilter.YFilter { return pbriSendStatistics.YFilter }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) SetFilter(yf yfilter.YFilter) { pbriSendStatistics.YFilter = yf }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetSegmentPath() string {
    return "pbri-send-statistics"
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriSendStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriSendStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriSendStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriSendStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriSendStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriSendStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriSendStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriSendStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriSendStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriSendStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriSendStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriSendStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetYangName() string { return "pbri-send-statistics" }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) SetParent(parent types.Entity) { pbriSendStatistics.parent = parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetParent() types.Entity { return pbriSendStatistics.parent }

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRIs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbriCount interface{}

    // Count of PBRIs dropped. The type is interface{} with range: 0..4294967295.
    PbriDropCount interface{}

    // Count of Revoc Trigger - Unspecified. The type is interface{} with range:
    // 0..4294967295.
    UnspecifiedCount interface{}

    // Count of Revoc Trigger - Administrative Reason. The type is interface{}
    // with range: 0..4294967295.
    AdminReasonCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Same ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverSameAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Different ATT. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverDifferentAttCount interface{}

    // Count of Revoc Trigger - Inter MAG Handover Unknown. The type is
    // interface{} with range: 0..4294967295.
    MagHandoverUnknownCount interface{}

    // Count of Revoc Trigger - User Init Session Terminatation. The type is
    // interface{} with range: 0..4294967295.
    UserSessionTerminationCount interface{}

    // Count of Revoc Trigger - Access Network Session Termination. The type is
    // interface{} with range: 0..4294967295.
    NetworkSessionTerminationCount interface{}

    // Count of Revoc Trigger - Possible Out-of-Sync BCE State. The type is
    // interface{} with range: 0..4294967295.
    OutOfSyncBceStateCount interface{}

    // Count of Revoc Trigger - Per-Peer Policy. The type is interface{} with
    // range: 0..4294967295.
    PerPeerPolicyCount interface{}

    // Count of Revoc Trigger - Revoking Mobility Node Local Policy. The type is
    // interface{} with range: 0..4294967295.
    RevokingMnLocalPolicyCount interface{}
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetFilter() yfilter.YFilter { return pbriReceiveStatistics.YFilter }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbriReceiveStatistics.YFilter = yf }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbri-count" { return "PbriCount" }
    if yname == "pbri-drop-count" { return "PbriDropCount" }
    if yname == "unspecified-count" { return "UnspecifiedCount" }
    if yname == "admin-reason-count" { return "AdminReasonCount" }
    if yname == "mag-handover-same-att-count" { return "MagHandoverSameAttCount" }
    if yname == "mag-handover-different-att-count" { return "MagHandoverDifferentAttCount" }
    if yname == "mag-handover-unknown-count" { return "MagHandoverUnknownCount" }
    if yname == "user-session-termination-count" { return "UserSessionTerminationCount" }
    if yname == "network-session-termination-count" { return "NetworkSessionTerminationCount" }
    if yname == "out-of-sync-bce-state-count" { return "OutOfSyncBceStateCount" }
    if yname == "per-peer-policy-count" { return "PerPeerPolicyCount" }
    if yname == "revoking-mn-local-policy-count" { return "RevokingMnLocalPolicyCount" }
    return ""
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetSegmentPath() string {
    return "pbri-receive-statistics"
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbri-count"] = pbriReceiveStatistics.PbriCount
    leafs["pbri-drop-count"] = pbriReceiveStatistics.PbriDropCount
    leafs["unspecified-count"] = pbriReceiveStatistics.UnspecifiedCount
    leafs["admin-reason-count"] = pbriReceiveStatistics.AdminReasonCount
    leafs["mag-handover-same-att-count"] = pbriReceiveStatistics.MagHandoverSameAttCount
    leafs["mag-handover-different-att-count"] = pbriReceiveStatistics.MagHandoverDifferentAttCount
    leafs["mag-handover-unknown-count"] = pbriReceiveStatistics.MagHandoverUnknownCount
    leafs["user-session-termination-count"] = pbriReceiveStatistics.UserSessionTerminationCount
    leafs["network-session-termination-count"] = pbriReceiveStatistics.NetworkSessionTerminationCount
    leafs["out-of-sync-bce-state-count"] = pbriReceiveStatistics.OutOfSyncBceStateCount
    leafs["per-peer-policy-count"] = pbriReceiveStatistics.PerPeerPolicyCount
    leafs["revoking-mn-local-policy-count"] = pbriReceiveStatistics.RevokingMnLocalPolicyCount
    return leafs
}

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetYangName() string { return "pbri-receive-statistics" }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) SetParent(parent types.Entity) { pbriReceiveStatistics.parent = parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetParent() types.Entity { return pbriReceiveStatistics.parent }

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetFilter() yfilter.YFilter { return pbraSendStatistics.YFilter }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) SetFilter(yf yfilter.YFilter) { pbraSendStatistics.YFilter = yf }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetSegmentPath() string {
    return "pbra-send-statistics"
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraSendStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraSendStatistics.PbraDropCount
    leafs["success-count"] = pbraSendStatistics.SuccessCount
    leafs["partial-success-count"] = pbraSendStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraSendStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraSendStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraSendStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraSendStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraSendStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraSendStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraSendStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraSendStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetYangName() string { return "pbra-send-statistics" }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) SetParent(parent types.Entity) { pbraSendStatistics.parent = parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetParent() types.Entity { return pbraSendStatistics.parent }

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Count of PBRAs. The type is interface{} with range:
    // 0..18446744073709551615.
    PbraCount interface{}

    // Count of PBRAs dropped. The type is interface{} with range: 0..4294967295.
    PbraDropCount interface{}

    // Count of Revoc Status - Success. The type is interface{} with range:
    // 0..4294967295.
    SuccessCount interface{}

    // Count of Revoc Status - Partial Success. The type is interface{} with
    // range: 0..4294967295.
    PartialSuccessCount interface{}

    // Count of Revoc Status - Binding Does Not Exist. The type is interface{}
    // with range: 0..4294967295.
    NoBindingCount interface{}

    // Count of Revoc Status - IPv4 Home Address Option Required. The type is
    // interface{} with range: 0..4294967295.
    HoaRequiredCount interface{}

    // Count of Revoc Status - Global Revocation NOT Authorized. The type is
    // interface{} with range: 0..4294967295.
    NoAuthorForGlobalRevocCount interface{}

    // Count of Revoc Status - Revoked Mobile Node Identity Required. The type is
    // interface{} with range: 0..4294967295.
    MnIdentityRequiredCount interface{}

    // Count of Revoc Status - Revocation Failed - MN is Attached. The type is
    // interface{} with range: 0..4294967295.
    MnAttachedCount interface{}

    // Count of Revoc Status - Revocation Trigger NOT supported. The type is
    // interface{} with range: 0..4294967295.
    UnknownRevocTriggerCount interface{}

    // Count of Revoc Status - Revocation Function NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    RevocFunctionNotSupportedCount interface{}

    // Count of Revoc Status - Proxy Binding Revocation NOT Supported. The type is
    // interface{} with range: 0..4294967295.
    PbrNotSupportedCount interface{}
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetFilter() yfilter.YFilter { return pbraReceiveStatistics.YFilter }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) SetFilter(yf yfilter.YFilter) { pbraReceiveStatistics.YFilter = yf }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetGoName(yname string) string {
    if yname == "pbra-count" { return "PbraCount" }
    if yname == "pbra-drop-count" { return "PbraDropCount" }
    if yname == "success-count" { return "SuccessCount" }
    if yname == "partial-success-count" { return "PartialSuccessCount" }
    if yname == "no-binding-count" { return "NoBindingCount" }
    if yname == "hoa-required-count" { return "HoaRequiredCount" }
    if yname == "no-author-for-global-revoc-count" { return "NoAuthorForGlobalRevocCount" }
    if yname == "mn-identity-required-count" { return "MnIdentityRequiredCount" }
    if yname == "mn-attached-count" { return "MnAttachedCount" }
    if yname == "unknown-revoc-trigger-count" { return "UnknownRevocTriggerCount" }
    if yname == "revoc-function-not-supported-count" { return "RevocFunctionNotSupportedCount" }
    if yname == "pbr-not-supported-count" { return "PbrNotSupportedCount" }
    return ""
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetSegmentPath() string {
    return "pbra-receive-statistics"
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pbra-count"] = pbraReceiveStatistics.PbraCount
    leafs["pbra-drop-count"] = pbraReceiveStatistics.PbraDropCount
    leafs["success-count"] = pbraReceiveStatistics.SuccessCount
    leafs["partial-success-count"] = pbraReceiveStatistics.PartialSuccessCount
    leafs["no-binding-count"] = pbraReceiveStatistics.NoBindingCount
    leafs["hoa-required-count"] = pbraReceiveStatistics.HoaRequiredCount
    leafs["no-author-for-global-revoc-count"] = pbraReceiveStatistics.NoAuthorForGlobalRevocCount
    leafs["mn-identity-required-count"] = pbraReceiveStatistics.MnIdentityRequiredCount
    leafs["mn-attached-count"] = pbraReceiveStatistics.MnAttachedCount
    leafs["unknown-revoc-trigger-count"] = pbraReceiveStatistics.UnknownRevocTriggerCount
    leafs["revoc-function-not-supported-count"] = pbraReceiveStatistics.RevocFunctionNotSupportedCount
    leafs["pbr-not-supported-count"] = pbraReceiveStatistics.PbrNotSupportedCount
    return leafs
}

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetYangName() string { return "pbra-receive-statistics" }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) SetParent(parent types.Entity) { pbraReceiveStatistics.parent = parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetParent() types.Entity { return pbraReceiveStatistics.parent }

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetParentYangName() string { return "protocol-statistics" }

// Pmipv6_Lma_Bindings
// Table of Binding
type Pmipv6_Lma_Bindings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Binding Parameters. The type is slice of Pmipv6_Lma_Bindings_Binding.
    Binding []Pmipv6_Lma_Bindings_Binding
}

func (bindings *Pmipv6_Lma_Bindings) GetFilter() yfilter.YFilter { return bindings.YFilter }

func (bindings *Pmipv6_Lma_Bindings) SetFilter(yf yfilter.YFilter) { bindings.YFilter = yf }

func (bindings *Pmipv6_Lma_Bindings) GetGoName(yname string) string {
    if yname == "binding" { return "Binding" }
    return ""
}

func (bindings *Pmipv6_Lma_Bindings) GetSegmentPath() string {
    return "bindings"
}

func (bindings *Pmipv6_Lma_Bindings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding" {
        for _, c := range bindings.Binding {
            if bindings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Bindings_Binding{}
        bindings.Binding = append(bindings.Binding, child)
        return &bindings.Binding[len(bindings.Binding)-1]
    }
    return nil
}

func (bindings *Pmipv6_Lma_Bindings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bindings.Binding {
        children[bindings.Binding[i].GetSegmentPath()] = &bindings.Binding[i]
    }
    return children
}

func (bindings *Pmipv6_Lma_Bindings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bindings *Pmipv6_Lma_Bindings) GetBundleName() string { return "cisco_ios_xr" }

func (bindings *Pmipv6_Lma_Bindings) GetYangName() string { return "bindings" }

func (bindings *Pmipv6_Lma_Bindings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bindings *Pmipv6_Lma_Bindings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bindings *Pmipv6_Lma_Bindings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bindings *Pmipv6_Lma_Bindings) SetParent(parent types.Entity) { bindings.parent = parent }

func (bindings *Pmipv6_Lma_Bindings) GetParent() types.Entity { return bindings.parent }

func (bindings *Pmipv6_Lma_Bindings) GetParentYangName() string { return "lma" }

// Pmipv6_Lma_Bindings_Binding
// Binding Parameters
type Pmipv6_Lma_Bindings_Binding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer MAG ID. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MagName interface{}

    // NAI String. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    NaiString interface{}

    // IMSI String. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ImsiString interface{}

    // Customer String. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CustomerName interface{}

    // Mobile Node Identifier. The type is string.
    Mnnai interface{}

    // Customer name. The type is string.
    CustomerNameXr interface{}

    // Link Layer Identifier. The type is string.
    Llid interface{}

    // Peer Identifier. The type is string.
    PeerId interface{}

    // Access Interface. The type is string.
    Phyintf interface{}

    // Tunnel Interface. The type is string.
    Tunnel interface{}

    // State Name. The type is string.
    State interface{}

    // Access Point Network. The type is string.
    Apn interface{}

    // MN ATT. The type is interface{} with range: 0..65535.
    Att interface{}

    // MN HOA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Hoa interface{}

    // MN Default Router. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Dflt interface{}

    // Life Time of Binding. The type is interface{} with range: 0..4294967295.
    Lifetime interface{}

    // Life Time Remaining. The type is interface{} with range: 0..4294967295.
    Liferem interface{}

    // Refresh Time of Binding. The type is interface{} with range: 0..4294967295.
    Refresh interface{}

    // Refresh Time Remaining. The type is interface{} with range: 0..4294967295.
    RefreshRem interface{}

    // Prefix Length. The type is interface{} with range: 0..255.
    PrefixLen interface{}

    // HNP count. The type is interface{} with range: 0..255.
    NumHnps interface{}

    // COA count. The type is interface{} with range: 0..255.
    NumCoa interface{}

    // IPv4 DMNP count. The type is interface{} with range: 0..255.
    NumDmnpV4 interface{}

    // IPv6 DMNP count. The type is interface{} with range: 0..255.
    NumDmnpV6 interface{}

    // MN Home Network Prefixes. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Hnps interface{}

    // Ignore HoA/HNP. The type is bool.
    IgnoreHomeAddress interface{}

    // Upstream GRE Key. The type is interface{} with range: 0..4294967295.
    UpStreamGrekey interface{}

    // DownStream GRE Key. The type is interface{} with range: 0..4294967295.
    DownStreamGrekey interface{}

    // VRF ID of Access Interface. The type is interface{} with range:
    // 0..4294967295.
    Vrfid interface{}

    // COA entries. The type is slice of Pmipv6_Lma_Bindings_Binding_Coa.
    Coa []Pmipv6_Lma_Bindings_Binding_Coa

    // IPv4 DMNP prefixes. The type is slice of
    // Pmipv6_Lma_Bindings_Binding_DmnpV4.
    DmnpV4 []Pmipv6_Lma_Bindings_Binding_DmnpV4

    // IPv6 DMNP prefixes. The type is slice of
    // Pmipv6_Lma_Bindings_Binding_DmnpV6.
    DmnpV6 []Pmipv6_Lma_Bindings_Binding_DmnpV6
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *Pmipv6_Lma_Bindings_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *Pmipv6_Lma_Bindings_Binding) GetGoName(yname string) string {
    if yname == "mag-name" { return "MagName" }
    if yname == "nai-string" { return "NaiString" }
    if yname == "imsi-string" { return "ImsiString" }
    if yname == "customer-name" { return "CustomerName" }
    if yname == "mnnai" { return "Mnnai" }
    if yname == "customer-name-xr" { return "CustomerNameXr" }
    if yname == "llid" { return "Llid" }
    if yname == "peer-id" { return "PeerId" }
    if yname == "phyintf" { return "Phyintf" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "state" { return "State" }
    if yname == "apn" { return "Apn" }
    if yname == "att" { return "Att" }
    if yname == "hoa" { return "Hoa" }
    if yname == "dflt" { return "Dflt" }
    if yname == "lifetime" { return "Lifetime" }
    if yname == "liferem" { return "Liferem" }
    if yname == "refresh" { return "Refresh" }
    if yname == "refresh-rem" { return "RefreshRem" }
    if yname == "prefix-len" { return "PrefixLen" }
    if yname == "num-hnps" { return "NumHnps" }
    if yname == "num-coa" { return "NumCoa" }
    if yname == "num-dmnp-v4" { return "NumDmnpV4" }
    if yname == "num-dmnp-v6" { return "NumDmnpV6" }
    if yname == "hnps" { return "Hnps" }
    if yname == "ignore-home-address" { return "IgnoreHomeAddress" }
    if yname == "up-stream-grekey" { return "UpStreamGrekey" }
    if yname == "down-stream-grekey" { return "DownStreamGrekey" }
    if yname == "vrfid" { return "Vrfid" }
    if yname == "coa" { return "Coa" }
    if yname == "dmnp-v4" { return "DmnpV4" }
    if yname == "dmnp-v6" { return "DmnpV6" }
    return ""
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetSegmentPath() string {
    return "binding"
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "coa" {
        for _, c := range binding.Coa {
            if binding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Bindings_Binding_Coa{}
        binding.Coa = append(binding.Coa, child)
        return &binding.Coa[len(binding.Coa)-1]
    }
    if childYangName == "dmnp-v4" {
        for _, c := range binding.DmnpV4 {
            if binding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Bindings_Binding_DmnpV4{}
        binding.DmnpV4 = append(binding.DmnpV4, child)
        return &binding.DmnpV4[len(binding.DmnpV4)-1]
    }
    if childYangName == "dmnp-v6" {
        for _, c := range binding.DmnpV6 {
            if binding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Bindings_Binding_DmnpV6{}
        binding.DmnpV6 = append(binding.DmnpV6, child)
        return &binding.DmnpV6[len(binding.DmnpV6)-1]
    }
    return nil
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range binding.Coa {
        children[binding.Coa[i].GetSegmentPath()] = &binding.Coa[i]
    }
    for i := range binding.DmnpV4 {
        children[binding.DmnpV4[i].GetSegmentPath()] = &binding.DmnpV4[i]
    }
    for i := range binding.DmnpV6 {
        children[binding.DmnpV6[i].GetSegmentPath()] = &binding.DmnpV6[i]
    }
    return children
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mag-name"] = binding.MagName
    leafs["nai-string"] = binding.NaiString
    leafs["imsi-string"] = binding.ImsiString
    leafs["customer-name"] = binding.CustomerName
    leafs["mnnai"] = binding.Mnnai
    leafs["customer-name-xr"] = binding.CustomerNameXr
    leafs["llid"] = binding.Llid
    leafs["peer-id"] = binding.PeerId
    leafs["phyintf"] = binding.Phyintf
    leafs["tunnel"] = binding.Tunnel
    leafs["state"] = binding.State
    leafs["apn"] = binding.Apn
    leafs["att"] = binding.Att
    leafs["hoa"] = binding.Hoa
    leafs["dflt"] = binding.Dflt
    leafs["lifetime"] = binding.Lifetime
    leafs["liferem"] = binding.Liferem
    leafs["refresh"] = binding.Refresh
    leafs["refresh-rem"] = binding.RefreshRem
    leafs["prefix-len"] = binding.PrefixLen
    leafs["num-hnps"] = binding.NumHnps
    leafs["num-coa"] = binding.NumCoa
    leafs["num-dmnp-v4"] = binding.NumDmnpV4
    leafs["num-dmnp-v6"] = binding.NumDmnpV6
    leafs["hnps"] = binding.Hnps
    leafs["ignore-home-address"] = binding.IgnoreHomeAddress
    leafs["up-stream-grekey"] = binding.UpStreamGrekey
    leafs["down-stream-grekey"] = binding.DownStreamGrekey
    leafs["vrfid"] = binding.Vrfid
    return leafs
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetBundleName() string { return "cisco_ios_xr" }

func (binding *Pmipv6_Lma_Bindings_Binding) GetYangName() string { return "binding" }

func (binding *Pmipv6_Lma_Bindings_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (binding *Pmipv6_Lma_Bindings_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (binding *Pmipv6_Lma_Bindings_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (binding *Pmipv6_Lma_Bindings_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *Pmipv6_Lma_Bindings_Binding) GetParent() types.Entity { return binding.parent }

func (binding *Pmipv6_Lma_Bindings_Binding) GetParentYangName() string { return "bindings" }

// Pmipv6_Lma_Bindings_Binding_Coa
// COA entries
type Pmipv6_Lma_Bindings_Binding_Coa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Layer Identifier. The type is string.
    Llid interface{}

    // Peer Name. The type is string.
    PeerName interface{}

    // Tunnel Interface. The type is string.
    Tunnel interface{}

    // Egress Label. The type is string.
    ELabel interface{}

    // Label Color. The type is string.
    Color interface{}

    // Roaming Intf. The type is string.
    RoaMinTf interface{}

    // COA STATE. The type is string.
    Pstate interface{}

    // MSISDN. The type is string.
    Msisdn interface{}

    // IMSI or IMSI NAI. The type is string.
    Imsi interface{}

    // CDMA NAI. The type is string.
    CdmaNai interface{}

    // Subscriber APN on PWG. The type is string.
    PgwApn interface{}

    // Subscriber Transport VRF on PGW. The type is string.
    PgwTransVrf interface{}

    // MN ATT. The type is interface{} with range: 0..65535.
    Att interface{}

    // Life Time of coa. The type is interface{} with range: 0..4294967295.
    Lifetime interface{}

    // Life Time remain of coa. The type is interface{} with range: 0..4294967295.
    LifetimeRemaining interface{}

    // refresh Time of coa. The type is interface{} with range: 0..4294967295.
    Refresh interface{}

    // refresh Time remain of coa. The type is interface{} with range:
    // 0..4294967295.
    RefreshRem interface{}

    // down key for coa tunnel. The type is interface{} with range: 0..4294967295.
    Dnkey interface{}

    // up key for coa tunnel. The type is interface{} with range: 0..4294967295.
    Upkey interface{}

    // IPv4 CoA. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CoaV4 interface{}

    // IPv6 CoA. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    CoaV6 interface{}
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetFilter() yfilter.YFilter { return coa.YFilter }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) SetFilter(yf yfilter.YFilter) { coa.YFilter = yf }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetGoName(yname string) string {
    if yname == "llid" { return "Llid" }
    if yname == "peer-name" { return "PeerName" }
    if yname == "tunnel" { return "Tunnel" }
    if yname == "e-label" { return "ELabel" }
    if yname == "color" { return "Color" }
    if yname == "roa-min-tf" { return "RoaMinTf" }
    if yname == "pstate" { return "Pstate" }
    if yname == "msisdn" { return "Msisdn" }
    if yname == "imsi" { return "Imsi" }
    if yname == "cdma-nai" { return "CdmaNai" }
    if yname == "pgw-apn" { return "PgwApn" }
    if yname == "pgw-trans-vrf" { return "PgwTransVrf" }
    if yname == "att" { return "Att" }
    if yname == "lifetime" { return "Lifetime" }
    if yname == "lifetime-remaining" { return "LifetimeRemaining" }
    if yname == "refresh" { return "Refresh" }
    if yname == "refresh-rem" { return "RefreshRem" }
    if yname == "dnkey" { return "Dnkey" }
    if yname == "upkey" { return "Upkey" }
    if yname == "coa-v4" { return "CoaV4" }
    if yname == "coa-v6" { return "CoaV6" }
    return ""
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetSegmentPath() string {
    return "coa"
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["llid"] = coa.Llid
    leafs["peer-name"] = coa.PeerName
    leafs["tunnel"] = coa.Tunnel
    leafs["e-label"] = coa.ELabel
    leafs["color"] = coa.Color
    leafs["roa-min-tf"] = coa.RoaMinTf
    leafs["pstate"] = coa.Pstate
    leafs["msisdn"] = coa.Msisdn
    leafs["imsi"] = coa.Imsi
    leafs["cdma-nai"] = coa.CdmaNai
    leafs["pgw-apn"] = coa.PgwApn
    leafs["pgw-trans-vrf"] = coa.PgwTransVrf
    leafs["att"] = coa.Att
    leafs["lifetime"] = coa.Lifetime
    leafs["lifetime-remaining"] = coa.LifetimeRemaining
    leafs["refresh"] = coa.Refresh
    leafs["refresh-rem"] = coa.RefreshRem
    leafs["dnkey"] = coa.Dnkey
    leafs["upkey"] = coa.Upkey
    leafs["coa-v4"] = coa.CoaV4
    leafs["coa-v6"] = coa.CoaV6
    return leafs
}

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetBundleName() string { return "cisco_ios_xr" }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetYangName() string { return "coa" }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) SetParent(parent types.Entity) { coa.parent = parent }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetParent() types.Entity { return coa.parent }

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetParentYangName() string { return "binding" }

// Pmipv6_Lma_Bindings_Binding_DmnpV4
// IPv4 DMNP prefixes
type Pmipv6_Lma_Bindings_Binding_DmnpV4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 prefix length. The type is interface{} with range: 0..255.
    Pfxlen interface{}

    // IPv4 prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetFilter() yfilter.YFilter { return dmnpV4.YFilter }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) SetFilter(yf yfilter.YFilter) { dmnpV4.YFilter = yf }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetGoName(yname string) string {
    if yname == "pfxlen" { return "Pfxlen" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetSegmentPath() string {
    return "dmnp-v4"
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pfxlen"] = dmnpV4.Pfxlen
    leafs["prefix"] = dmnpV4.Prefix
    return leafs
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetBundleName() string { return "cisco_ios_xr" }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetYangName() string { return "dmnp-v4" }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) SetParent(parent types.Entity) { dmnpV4.parent = parent }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetParent() types.Entity { return dmnpV4.parent }

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetParentYangName() string { return "binding" }

// Pmipv6_Lma_Bindings_Binding_DmnpV6
// IPv6 DMNP prefixes
type Pmipv6_Lma_Bindings_Binding_DmnpV6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 prefix length. The type is interface{} with range: 0..255.
    Pfxlen interface{}

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetFilter() yfilter.YFilter { return dmnpV6.YFilter }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) SetFilter(yf yfilter.YFilter) { dmnpV6.YFilter = yf }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetGoName(yname string) string {
    if yname == "pfxlen" { return "Pfxlen" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetSegmentPath() string {
    return "dmnp-v6"
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pfxlen"] = dmnpV6.Pfxlen
    leafs["prefix"] = dmnpV6.Prefix
    return leafs
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetBundleName() string { return "cisco_ios_xr" }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetYangName() string { return "dmnp-v6" }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) SetParent(parent types.Entity) { dmnpV6.parent = parent }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetParent() types.Entity { return dmnpV6.parent }

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetParentYangName() string { return "binding" }

// Pmipv6_Lma_Heartbeats
// Table of Heartbeat
type Pmipv6_Lma_Heartbeats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Heartbeat information. The type is slice of
    // Pmipv6_Lma_Heartbeats_Heartbeat.
    Heartbeat []Pmipv6_Lma_Heartbeats_Heartbeat
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetFilter() yfilter.YFilter { return heartbeats.YFilter }

func (heartbeats *Pmipv6_Lma_Heartbeats) SetFilter(yf yfilter.YFilter) { heartbeats.YFilter = yf }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetGoName(yname string) string {
    if yname == "heartbeat" { return "Heartbeat" }
    return ""
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetSegmentPath() string {
    return "heartbeats"
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "heartbeat" {
        for _, c := range heartbeats.Heartbeat {
            if heartbeats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_Heartbeats_Heartbeat{}
        heartbeats.Heartbeat = append(heartbeats.Heartbeat, child)
        return &heartbeats.Heartbeat[len(heartbeats.Heartbeat)-1]
    }
    return nil
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range heartbeats.Heartbeat {
        children[heartbeats.Heartbeat[i].GetSegmentPath()] = &heartbeats.Heartbeat[i]
    }
    return children
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetBundleName() string { return "cisco_ios_xr" }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetYangName() string { return "heartbeats" }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (heartbeats *Pmipv6_Lma_Heartbeats) SetParent(parent types.Entity) { heartbeats.parent = parent }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetParent() types.Entity { return heartbeats.parent }

func (heartbeats *Pmipv6_Lma_Heartbeats) GetParentYangName() string { return "lma" }

// Pmipv6_Lma_Heartbeats_Heartbeat
// Heartbeat information
type Pmipv6_Lma_Heartbeats_Heartbeat struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 or IPv6 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddr interface{}

    // VRF Name. The type is string.
    Vrf interface{}

    // Customer Name. The type is string.
    CustomerName interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Destination Port. The type is interface{} with range: 0..4294967295.
    DestinationPort interface{}

    // Source IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceIpv4Address interface{}

    // Destination IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationIpv4Address interface{}

    // Source IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceIpv6Address interface{}

    // Destination IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationIpv6Address interface{}

    // Path Status. The type is bool.
    Status interface{}

    // IPv6 Path. The type is bool.
    Ipv6Path interface{}
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetFilter() yfilter.YFilter { return heartbeat.YFilter }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) SetFilter(yf yfilter.YFilter) { heartbeat.YFilter = yf }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetGoName(yname string) string {
    if yname == "peer-addr" { return "PeerAddr" }
    if yname == "vrf" { return "Vrf" }
    if yname == "customer-name" { return "CustomerName" }
    if yname == "source-port" { return "SourcePort" }
    if yname == "destination-port" { return "DestinationPort" }
    if yname == "source-ipv4-address" { return "SourceIpv4Address" }
    if yname == "destination-ipv4-address" { return "DestinationIpv4Address" }
    if yname == "source-ipv6-address" { return "SourceIpv6Address" }
    if yname == "destination-ipv6-address" { return "DestinationIpv6Address" }
    if yname == "status" { return "Status" }
    if yname == "ipv6-path" { return "Ipv6Path" }
    return ""
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetSegmentPath() string {
    return "heartbeat" + "[peer-addr='" + fmt.Sprintf("%v", heartbeat.PeerAddr) + "']"
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-addr"] = heartbeat.PeerAddr
    leafs["vrf"] = heartbeat.Vrf
    leafs["customer-name"] = heartbeat.CustomerName
    leafs["source-port"] = heartbeat.SourcePort
    leafs["destination-port"] = heartbeat.DestinationPort
    leafs["source-ipv4-address"] = heartbeat.SourceIpv4Address
    leafs["destination-ipv4-address"] = heartbeat.DestinationIpv4Address
    leafs["source-ipv6-address"] = heartbeat.SourceIpv6Address
    leafs["destination-ipv6-address"] = heartbeat.DestinationIpv6Address
    leafs["status"] = heartbeat.Status
    leafs["ipv6-path"] = heartbeat.Ipv6Path
    return leafs
}

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetBundleName() string { return "cisco_ios_xr" }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetYangName() string { return "heartbeat" }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) SetParent(parent types.Entity) { heartbeat.parent = parent }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetParent() types.Entity { return heartbeat.parent }

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetParentYangName() string { return "heartbeats" }

// Pmipv6_Lma_ConfigVariables
// Global Configuration Variables
type Pmipv6_Lma_ConfigVariables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of CustomerVariables.
    CustomerVariables Pmipv6_Lma_ConfigVariables_CustomerVariables

    // Global Configuration Variables.
    GlobalVariables Pmipv6_Lma_ConfigVariables_GlobalVariables
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetFilter() yfilter.YFilter { return configVariables.YFilter }

func (configVariables *Pmipv6_Lma_ConfigVariables) SetFilter(yf yfilter.YFilter) { configVariables.YFilter = yf }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetGoName(yname string) string {
    if yname == "customer-variables" { return "CustomerVariables" }
    if yname == "global-variables" { return "GlobalVariables" }
    return ""
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetSegmentPath() string {
    return "config-variables"
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customer-variables" {
        return &configVariables.CustomerVariables
    }
    if childYangName == "global-variables" {
        return &configVariables.GlobalVariables
    }
    return nil
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["customer-variables"] = &configVariables.CustomerVariables
    children["global-variables"] = &configVariables.GlobalVariables
    return children
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetBundleName() string { return "cisco_ios_xr" }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetYangName() string { return "config-variables" }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configVariables *Pmipv6_Lma_ConfigVariables) SetParent(parent types.Entity) { configVariables.parent = parent }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetParent() types.Entity { return configVariables.parent }

func (configVariables *Pmipv6_Lma_ConfigVariables) GetParentYangName() string { return "lma" }

// Pmipv6_Lma_ConfigVariables_CustomerVariables
// Table of CustomerVariables
type Pmipv6_Lma_ConfigVariables_CustomerVariables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customer name string. The type is slice of
    // Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable.
    CustomerVariable []Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetFilter() yfilter.YFilter { return customerVariables.YFilter }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) SetFilter(yf yfilter.YFilter) { customerVariables.YFilter = yf }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetGoName(yname string) string {
    if yname == "customer-variable" { return "CustomerVariable" }
    return ""
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetSegmentPath() string {
    return "customer-variables"
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "customer-variable" {
        for _, c := range customerVariables.CustomerVariable {
            if customerVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable{}
        customerVariables.CustomerVariable = append(customerVariables.CustomerVariable, child)
        return &customerVariables.CustomerVariable[len(customerVariables.CustomerVariable)-1]
    }
    return nil
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range customerVariables.CustomerVariable {
        children[customerVariables.CustomerVariable[i].GetSegmentPath()] = &customerVariables.CustomerVariable[i]
    }
    return children
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetBundleName() string { return "cisco_ios_xr" }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetYangName() string { return "customer-variables" }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) SetParent(parent types.Entity) { customerVariables.parent = parent }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetParent() types.Entity { return customerVariables.parent }

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetParentYangName() string { return "config-variables" }

// Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable
// Customer name string
type Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Customer name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    CustomerName interface{}

    // Customer Name. The type is string.
    CustName interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Authentication Option. The type is bool.
    AuthOption interface{}

    // MLL service parameters.
    MllService Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetFilter() yfilter.YFilter { return customerVariable.YFilter }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) SetFilter(yf yfilter.YFilter) { customerVariable.YFilter = yf }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetGoName(yname string) string {
    if yname == "customer-name" { return "CustomerName" }
    if yname == "cust-name" { return "CustName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "auth-option" { return "AuthOption" }
    if yname == "mll-service" { return "MllService" }
    return ""
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetSegmentPath() string {
    return "customer-variable" + "[customer-name='" + fmt.Sprintf("%v", customerVariable.CustomerName) + "']"
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mll-service" {
        return &customerVariable.MllService
    }
    return nil
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mll-service"] = &customerVariable.MllService
    return children
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["customer-name"] = customerVariable.CustomerName
    leafs["cust-name"] = customerVariable.CustName
    leafs["vrf-name"] = customerVariable.VrfName
    leafs["auth-option"] = customerVariable.AuthOption
    return leafs
}

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetBundleName() string { return "cisco_ios_xr" }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetYangName() string { return "customer-variable" }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) SetParent(parent types.Entity) { customerVariable.parent = parent }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetParent() types.Entity { return customerVariable.parent }

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetParentYangName() string { return "customer-variables" }

// Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService
// MLL service parameters
type Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore Home Address. The type is bool.
    IgnoreHoa interface{}

    // Max IPv4 prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpIpv4LmnMax interface{}

    // Max IPv6 prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpIpv6LmnMax interface{}

    // Max prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpLmnMax interface{}

    // Max IPv4 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv4CustMax interface{}

    // Max IPv6 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv6CustMax interface{}

    // Max prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpCustMax interface{}

    // Current IPv4 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv4CustCur interface{}

    // Current IPv6 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv6CustCur interface{}
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetFilter() yfilter.YFilter { return mllService.YFilter }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) SetFilter(yf yfilter.YFilter) { mllService.YFilter = yf }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetGoName(yname string) string {
    if yname == "ignore-hoa" { return "IgnoreHoa" }
    if yname == "mnp-ipv4-lmn-max" { return "MnpIpv4LmnMax" }
    if yname == "mnp-ipv6-lmn-max" { return "MnpIpv6LmnMax" }
    if yname == "mnp-lmn-max" { return "MnpLmnMax" }
    if yname == "mnp-ipv4-cust-max" { return "MnpIpv4CustMax" }
    if yname == "mnp-ipv6-cust-max" { return "MnpIpv6CustMax" }
    if yname == "mnp-cust-max" { return "MnpCustMax" }
    if yname == "mnp-ipv4-cust-cur" { return "MnpIpv4CustCur" }
    if yname == "mnp-ipv6-cust-cur" { return "MnpIpv6CustCur" }
    return ""
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetSegmentPath() string {
    return "mll-service"
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ignore-hoa"] = mllService.IgnoreHoa
    leafs["mnp-ipv4-lmn-max"] = mllService.MnpIpv4LmnMax
    leafs["mnp-ipv6-lmn-max"] = mllService.MnpIpv6LmnMax
    leafs["mnp-lmn-max"] = mllService.MnpLmnMax
    leafs["mnp-ipv4-cust-max"] = mllService.MnpIpv4CustMax
    leafs["mnp-ipv6-cust-max"] = mllService.MnpIpv6CustMax
    leafs["mnp-cust-max"] = mllService.MnpCustMax
    leafs["mnp-ipv4-cust-cur"] = mllService.MnpIpv4CustCur
    leafs["mnp-ipv6-cust-cur"] = mllService.MnpIpv6CustCur
    return leafs
}

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetBundleName() string { return "cisco_ios_xr" }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetYangName() string { return "mll-service" }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) SetParent(parent types.Entity) { mllService.parent = parent }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetParent() types.Entity { return mllService.parent }

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetParentYangName() string { return "customer-variable" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables
// Global Configuration Variables
type Pmipv6_Lma_ConfigVariables_GlobalVariables struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain Name. The type is string.
    Domain interface{}

    // Self ID. The type is string.
    Selfid interface{}

    // APN Name. The type is string.
    ApnName interface{}

    // Role Type. The type is Pmipv6Role.
    Role interface{}

    // Number of Networks/Intf. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // Number of Peers. The type is interface{} with range: 0..4294967295.
    Peers interface{}

    // Number of Customers. The type is interface{} with range: 0..4294967295.
    Customers interface{}

    // Number of Networks. The type is interface{} with range: 0..4294967295.
    NumNetwork interface{}

    // Discover MN Detachment. The type is bool.
    DiscoverMn interface{}

    // Local Routing. The type is bool.
    LocalRouting interface{}

    // AAA Accounting. The type is bool.
    AaaAccounting interface{}

    // Default MN Enabled. The type is bool.
    DefaultMn interface{}

    // APN Present. The type is bool.
    Apn interface{}

    // Learn MAG. The type is bool.
    LearnMag interface{}

    // Session Manager. The type is bool.
    SessionMgr interface{}

    // Service. The type is interface{} with range: 0..255.
    Service interface{}

    // Default MN Profile Name. The type is string.
    Profile interface{}

    // Discover Detach Period. The type is interface{} with range: 0..4294967295.
    Ddp interface{}

    // Discover Detach Timeout. The type is interface{} with range: 0..4294967295.
    Ddt interface{}

    // Discover Detach Retries. The type is interface{} with range: 0..255.
    Ddr interface{}

    // Domain Parameters.
    Parameters Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters

    // MLL service parameters.
    MllService Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService

    // MAG Access List. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf.
    Intf []Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf

    // Peer Parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer.
    Peer []Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer

    // LMA Network Parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Network.
    Network []Pmipv6_Lma_ConfigVariables_GlobalVariables_Network

    // Customer parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust.
    Cust []Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetFilter() yfilter.YFilter { return globalVariables.YFilter }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) SetFilter(yf yfilter.YFilter) { globalVariables.YFilter = yf }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetGoName(yname string) string {
    if yname == "domain" { return "Domain" }
    if yname == "selfid" { return "Selfid" }
    if yname == "apn-name" { return "ApnName" }
    if yname == "role" { return "Role" }
    if yname == "count" { return "Count" }
    if yname == "peers" { return "Peers" }
    if yname == "customers" { return "Customers" }
    if yname == "num-network" { return "NumNetwork" }
    if yname == "discover-mn" { return "DiscoverMn" }
    if yname == "local-routing" { return "LocalRouting" }
    if yname == "aaa-accounting" { return "AaaAccounting" }
    if yname == "default-mn" { return "DefaultMn" }
    if yname == "apn" { return "Apn" }
    if yname == "learn-mag" { return "LearnMag" }
    if yname == "session-mgr" { return "SessionMgr" }
    if yname == "service" { return "Service" }
    if yname == "profile" { return "Profile" }
    if yname == "ddp" { return "Ddp" }
    if yname == "ddt" { return "Ddt" }
    if yname == "ddr" { return "Ddr" }
    if yname == "parameters" { return "Parameters" }
    if yname == "mll-service" { return "MllService" }
    if yname == "intf" { return "Intf" }
    if yname == "peer" { return "Peer" }
    if yname == "network" { return "Network" }
    if yname == "cust" { return "Cust" }
    return ""
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetSegmentPath() string {
    return "global-variables"
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "parameters" {
        return &globalVariables.Parameters
    }
    if childYangName == "mll-service" {
        return &globalVariables.MllService
    }
    if childYangName == "intf" {
        for _, c := range globalVariables.Intf {
            if globalVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf{}
        globalVariables.Intf = append(globalVariables.Intf, child)
        return &globalVariables.Intf[len(globalVariables.Intf)-1]
    }
    if childYangName == "peer" {
        for _, c := range globalVariables.Peer {
            if globalVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer{}
        globalVariables.Peer = append(globalVariables.Peer, child)
        return &globalVariables.Peer[len(globalVariables.Peer)-1]
    }
    if childYangName == "network" {
        for _, c := range globalVariables.Network {
            if globalVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_ConfigVariables_GlobalVariables_Network{}
        globalVariables.Network = append(globalVariables.Network, child)
        return &globalVariables.Network[len(globalVariables.Network)-1]
    }
    if childYangName == "cust" {
        for _, c := range globalVariables.Cust {
            if globalVariables.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust{}
        globalVariables.Cust = append(globalVariables.Cust, child)
        return &globalVariables.Cust[len(globalVariables.Cust)-1]
    }
    return nil
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["parameters"] = &globalVariables.Parameters
    children["mll-service"] = &globalVariables.MllService
    for i := range globalVariables.Intf {
        children[globalVariables.Intf[i].GetSegmentPath()] = &globalVariables.Intf[i]
    }
    for i := range globalVariables.Peer {
        children[globalVariables.Peer[i].GetSegmentPath()] = &globalVariables.Peer[i]
    }
    for i := range globalVariables.Network {
        children[globalVariables.Network[i].GetSegmentPath()] = &globalVariables.Network[i]
    }
    for i := range globalVariables.Cust {
        children[globalVariables.Cust[i].GetSegmentPath()] = &globalVariables.Cust[i]
    }
    return children
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain"] = globalVariables.Domain
    leafs["selfid"] = globalVariables.Selfid
    leafs["apn-name"] = globalVariables.ApnName
    leafs["role"] = globalVariables.Role
    leafs["count"] = globalVariables.Count
    leafs["peers"] = globalVariables.Peers
    leafs["customers"] = globalVariables.Customers
    leafs["num-network"] = globalVariables.NumNetwork
    leafs["discover-mn"] = globalVariables.DiscoverMn
    leafs["local-routing"] = globalVariables.LocalRouting
    leafs["aaa-accounting"] = globalVariables.AaaAccounting
    leafs["default-mn"] = globalVariables.DefaultMn
    leafs["apn"] = globalVariables.Apn
    leafs["learn-mag"] = globalVariables.LearnMag
    leafs["session-mgr"] = globalVariables.SessionMgr
    leafs["service"] = globalVariables.Service
    leafs["profile"] = globalVariables.Profile
    leafs["ddp"] = globalVariables.Ddp
    leafs["ddt"] = globalVariables.Ddt
    leafs["ddr"] = globalVariables.Ddr
    return leafs
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetBundleName() string { return "cisco_ios_xr" }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetYangName() string { return "global-variables" }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) SetParent(parent types.Entity) { globalVariables.parent = parent }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetParent() types.Entity { return globalVariables.parent }

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetParentYangName() string { return "config-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters
// Domain Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timestamp method in use. The type is bool.
    Timestamp interface{}

    // Timestamp Validity Window. The type is interface{} with range:
    // 0..18446744073709551615.
    Window interface{}

    // Authentication Option. The type is bool.
    AuthOption interface{}

    // BCE Registration Lifetime. The type is interface{} with range:
    // 0..4294967295.
    RegTime interface{}

    // BCE Refresh Time. The type is interface{} with range: 0..4294967295.
    RefTime interface{}

    // Refresh Retransmit Init. The type is interface{} with range: 0..65535.
    Retx interface{}

    // Refresh Retransmit Max. The type is interface{} with range: 0..65535.
    RetMax interface{}

    // BRI Init Delay time. The type is interface{} with range: 0..65535.
    BriInit interface{}

    // BRI Max Retries. The type is interface{} with range: 0..65535.
    BriRetries interface{}

    // BRI Max Delay time. The type is interface{} with range: 0..65535.
    BriMax interface{}

    // Allowed Max. Bindings. The type is interface{} with range: 0..4294967295.
    MaxBindings interface{}

    // Allowed HNPs per MN Intf. The type is interface{} with range: 0..255.
    Hnp interface{}

    // Encapsulation Type. The type is Pmipv6Encap.
    Encap interface{}

    // BCE Delete Hold Timer. The type is interface{} with range: 0..65535.
    DeleteTime interface{}

    // BCE Create Wait Timer. The type is interface{} with range: 0..65535.
    CreateTime interface{}

    // Upstream GRE Key. The type is interface{} with range: 0..4294967295.
    UpGrekey interface{}

    // Downstream GRE Key. The type is interface{} with range: 0..4294967295.
    DownGrekey interface{}

    // Self Identifier.
    SelfId Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetFilter() yfilter.YFilter { return parameters.YFilter }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) SetFilter(yf yfilter.YFilter) { parameters.YFilter = yf }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetGoName(yname string) string {
    if yname == "timestamp" { return "Timestamp" }
    if yname == "window" { return "Window" }
    if yname == "auth-option" { return "AuthOption" }
    if yname == "reg-time" { return "RegTime" }
    if yname == "ref-time" { return "RefTime" }
    if yname == "retx" { return "Retx" }
    if yname == "ret-max" { return "RetMax" }
    if yname == "bri-init" { return "BriInit" }
    if yname == "bri-retries" { return "BriRetries" }
    if yname == "bri-max" { return "BriMax" }
    if yname == "max-bindings" { return "MaxBindings" }
    if yname == "hnp" { return "Hnp" }
    if yname == "encap" { return "Encap" }
    if yname == "delete-time" { return "DeleteTime" }
    if yname == "create-time" { return "CreateTime" }
    if yname == "up-grekey" { return "UpGrekey" }
    if yname == "down-grekey" { return "DownGrekey" }
    if yname == "self-id" { return "SelfId" }
    return ""
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetSegmentPath() string {
    return "parameters"
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "self-id" {
        return &parameters.SelfId
    }
    return nil
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["self-id"] = &parameters.SelfId
    return children
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timestamp"] = parameters.Timestamp
    leafs["window"] = parameters.Window
    leafs["auth-option"] = parameters.AuthOption
    leafs["reg-time"] = parameters.RegTime
    leafs["ref-time"] = parameters.RefTime
    leafs["retx"] = parameters.Retx
    leafs["ret-max"] = parameters.RetMax
    leafs["bri-init"] = parameters.BriInit
    leafs["bri-retries"] = parameters.BriRetries
    leafs["bri-max"] = parameters.BriMax
    leafs["max-bindings"] = parameters.MaxBindings
    leafs["hnp"] = parameters.Hnp
    leafs["encap"] = parameters.Encap
    leafs["delete-time"] = parameters.DeleteTime
    leafs["create-time"] = parameters.CreateTime
    leafs["up-grekey"] = parameters.UpGrekey
    leafs["down-grekey"] = parameters.DownGrekey
    return leafs
}

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetBundleName() string { return "cisco_ios_xr" }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetYangName() string { return "parameters" }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) SetParent(parent types.Entity) { parameters.parent = parent }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetParent() types.Entity { return parameters.parent }

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetParentYangName() string { return "global-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId
// Self Identifier
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Identifier of PMIP Node. The type is string.
    Entity interface{}

    // IPV4 or IPV6 or Both. The type is Pmipv6Addr.
    AddrType interface{}

    // IPV6 address of LMA/MAG. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // IPV4 addrress of LMA/MAG. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetFilter() yfilter.YFilter { return selfId.YFilter }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) SetFilter(yf yfilter.YFilter) { selfId.YFilter = yf }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetGoName(yname string) string {
    if yname == "entity" { return "Entity" }
    if yname == "addr-type" { return "AddrType" }
    if yname == "address" { return "Address" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    return ""
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetSegmentPath() string {
    return "self-id"
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entity"] = selfId.Entity
    leafs["addr-type"] = selfId.AddrType
    leafs["address"] = selfId.Address
    leafs["ipv4-address"] = selfId.Ipv4Address
    return leafs
}

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetBundleName() string { return "cisco_ios_xr" }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetYangName() string { return "self-id" }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) SetParent(parent types.Entity) { selfId.parent = parent }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetParent() types.Entity { return selfId.parent }

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetParentYangName() string { return "parameters" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService
// MLL service parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ignore Home Address. The type is bool.
    IgnoreHoa interface{}

    // Max IPv4 prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpIpv4LmnMax interface{}

    // Max IPv6 prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpIpv6LmnMax interface{}

    // Max prefixes per LMN. The type is interface{} with range: 0..65535.
    MnpLmnMax interface{}

    // Max IPv4 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv4CustMax interface{}

    // Max IPv6 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv6CustMax interface{}

    // Max prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpCustMax interface{}

    // Current IPv4 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv4CustCur interface{}

    // Current IPv6 prefixes per Customer. The type is interface{} with range:
    // 0..4294967295.
    MnpIpv6CustCur interface{}
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetFilter() yfilter.YFilter { return mllService.YFilter }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) SetFilter(yf yfilter.YFilter) { mllService.YFilter = yf }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetGoName(yname string) string {
    if yname == "ignore-hoa" { return "IgnoreHoa" }
    if yname == "mnp-ipv4-lmn-max" { return "MnpIpv4LmnMax" }
    if yname == "mnp-ipv6-lmn-max" { return "MnpIpv6LmnMax" }
    if yname == "mnp-lmn-max" { return "MnpLmnMax" }
    if yname == "mnp-ipv4-cust-max" { return "MnpIpv4CustMax" }
    if yname == "mnp-ipv6-cust-max" { return "MnpIpv6CustMax" }
    if yname == "mnp-cust-max" { return "MnpCustMax" }
    if yname == "mnp-ipv4-cust-cur" { return "MnpIpv4CustCur" }
    if yname == "mnp-ipv6-cust-cur" { return "MnpIpv6CustCur" }
    return ""
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetSegmentPath() string {
    return "mll-service"
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ignore-hoa"] = mllService.IgnoreHoa
    leafs["mnp-ipv4-lmn-max"] = mllService.MnpIpv4LmnMax
    leafs["mnp-ipv6-lmn-max"] = mllService.MnpIpv6LmnMax
    leafs["mnp-lmn-max"] = mllService.MnpLmnMax
    leafs["mnp-ipv4-cust-max"] = mllService.MnpIpv4CustMax
    leafs["mnp-ipv6-cust-max"] = mllService.MnpIpv6CustMax
    leafs["mnp-cust-max"] = mllService.MnpCustMax
    leafs["mnp-ipv4-cust-cur"] = mllService.MnpIpv4CustCur
    leafs["mnp-ipv6-cust-cur"] = mllService.MnpIpv6CustCur
    return leafs
}

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetBundleName() string { return "cisco_ios_xr" }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetYangName() string { return "mll-service" }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) SetParent(parent types.Entity) { mllService.parent = parent }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetParent() types.Entity { return mllService.parent }

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetParentYangName() string { return "global-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf
// MAG Access List
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // APN Present. The type is bool.
    Apn interface{}

    // Access Interface Name. The type is string.
    Interface interface{}

    // APN Name. The type is string.
    ApnName interface{}
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetFilter() yfilter.YFilter { return intf.YFilter }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) SetFilter(yf yfilter.YFilter) { intf.YFilter = yf }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetGoName(yname string) string {
    if yname == "apn" { return "Apn" }
    if yname == "interface" { return "Interface" }
    if yname == "apn-name" { return "ApnName" }
    return ""
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetSegmentPath() string {
    return "intf"
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["apn"] = intf.Apn
    leafs["interface"] = intf.Interface
    leafs["apn-name"] = intf.ApnName
    return leafs
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetBundleName() string { return "cisco_ios_xr" }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetYangName() string { return "intf" }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) SetParent(parent types.Entity) { intf.parent = parent }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetParent() types.Entity { return intf.parent }

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetParentYangName() string { return "global-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer
// Peer Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer Name. The type is string.
    Peer interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Peer static tunnel intf. The type is string.
    Interface interface{}

    // Encapsulation Type. The type is Pmipv6Encap.
    Encap interface{}

    // Authentication Option. The type is bool.
    Auth interface{}

    // VRF Present. The type is bool.
    Vrf interface{}

    // Static tunnel Present. The type is bool.
    Statictunnel interface{}
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetGoName(yname string) string {
    if yname == "peer" { return "Peer" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "interface" { return "Interface" }
    if yname == "encap" { return "Encap" }
    if yname == "auth" { return "Auth" }
    if yname == "vrf" { return "Vrf" }
    if yname == "statictunnel" { return "Statictunnel" }
    return ""
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer"] = peer.Peer
    leafs["vrf-name"] = peer.VrfName
    leafs["interface"] = peer.Interface
    leafs["encap"] = peer.Encap
    leafs["auth"] = peer.Auth
    leafs["vrf"] = peer.Vrf
    leafs["statictunnel"] = peer.Statictunnel
    return leafs
}

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetYangName() string { return "peer" }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetParentYangName() string { return "global-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Network
// LMA Network Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Network struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV4 pool Present. The type is bool.
    V4Pool interface{}

    // IPV6 pool Present. The type is bool.
    V6Pool interface{}

    // Network Name. The type is string.
    Network interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // v4 prefix len. The type is interface{} with range: 0..255.
    V4PfxLen interface{}

    // v6 prefix len. The type is interface{} with range: 0..255.
    V6PfxLen interface{}

    // num of mrnet. The type is interface{} with range: 0..255.
    Mrnet interface{}
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetFilter() yfilter.YFilter { return network.YFilter }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) SetFilter(yf yfilter.YFilter) { network.YFilter = yf }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetGoName(yname string) string {
    if yname == "v4pool" { return "V4Pool" }
    if yname == "v6pool" { return "V6Pool" }
    if yname == "network" { return "Network" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "v4pfx-len" { return "V4PfxLen" }
    if yname == "v6pfx-len" { return "V6PfxLen" }
    if yname == "mrnet" { return "Mrnet" }
    return ""
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetSegmentPath() string {
    return "network"
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["v4pool"] = network.V4Pool
    leafs["v6pool"] = network.V6Pool
    leafs["network"] = network.Network
    leafs["ipv4"] = network.Ipv4
    leafs["ipv6"] = network.Ipv6
    leafs["v4pfx-len"] = network.V4PfxLen
    leafs["v6pfx-len"] = network.V6PfxLen
    leafs["mrnet"] = network.Mrnet
    return leafs
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetBundleName() string { return "cisco_ios_xr" }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetYangName() string { return "network" }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) SetParent(parent types.Entity) { network.parent = parent }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetParent() types.Entity { return network.parent }

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetParentYangName() string { return "global-variables" }

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust
// Customer parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Customer Present. The type is bool.
    Cust interface{}

    // Customer VRF Present. The type is bool.
    Vrf interface{}

    // Transport VRF Present. The type is bool.
    TVrf interface{}

    // Authentication Option. The type is bool.
    AuthOption interface{}

    // HeartBeat Option. The type is bool.
    HeartBeat interface{}

    // BCE Registration Lifetime. The type is interface{} with range:
    // 0..4294967295.
    RegTime interface{}

    // CUSTOMER Name. The type is string.
    CustName interface{}

    // VRF Name. The type is string.
    VrfName interface{}

    // Transport VRF Name. The type is string.
    TVrfName interface{}
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetFilter() yfilter.YFilter { return cust.YFilter }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) SetFilter(yf yfilter.YFilter) { cust.YFilter = yf }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetGoName(yname string) string {
    if yname == "cust" { return "Cust" }
    if yname == "vrf" { return "Vrf" }
    if yname == "t-vrf" { return "TVrf" }
    if yname == "auth-option" { return "AuthOption" }
    if yname == "heart-beat" { return "HeartBeat" }
    if yname == "reg-time" { return "RegTime" }
    if yname == "cust-name" { return "CustName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "t-vrf-name" { return "TVrfName" }
    return ""
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetSegmentPath() string {
    return "cust"
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cust"] = cust.Cust
    leafs["vrf"] = cust.Vrf
    leafs["t-vrf"] = cust.TVrf
    leafs["auth-option"] = cust.AuthOption
    leafs["heart-beat"] = cust.HeartBeat
    leafs["reg-time"] = cust.RegTime
    leafs["cust-name"] = cust.CustName
    leafs["vrf-name"] = cust.VrfName
    leafs["t-vrf-name"] = cust.TVrfName
    return leafs
}

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetBundleName() string { return "cisco_ios_xr" }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetYangName() string { return "cust" }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) SetParent(parent types.Entity) { cust.parent = parent }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetParent() types.Entity { return cust.parent }

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetParentYangName() string { return "global-variables" }

