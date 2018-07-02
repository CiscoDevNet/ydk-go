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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // None.
    Lma Pmipv6_Lma
}

func (pmipv6 *Pmipv6) GetEntityData() *types.CommonEntityData {
    pmipv6.EntityData.YFilter = pmipv6.YFilter
    pmipv6.EntityData.YangName = "pmipv6"
    pmipv6.EntityData.BundleName = "cisco_ios_xr"
    pmipv6.EntityData.ParentYangName = "Cisco-IOS-XR-ip-mobileip-oper"
    pmipv6.EntityData.SegmentPath = "Cisco-IOS-XR-ip-mobileip-oper:pmipv6"
    pmipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pmipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pmipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pmipv6.EntityData.Children = types.NewOrderedMap()
    pmipv6.EntityData.Children.Append("lma", types.YChild{"Lma", &pmipv6.Lma})
    pmipv6.EntityData.Leafs = types.NewOrderedMap()

    pmipv6.EntityData.YListKeys = []string {}

    return &(pmipv6.EntityData)
}

// Pmipv6_Lma
// None
type Pmipv6_Lma struct {
    EntityData types.CommonEntityData
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

func (lma *Pmipv6_Lma) GetEntityData() *types.CommonEntityData {
    lma.EntityData.YFilter = lma.YFilter
    lma.EntityData.YangName = "lma"
    lma.EntityData.BundleName = "cisco_ios_xr"
    lma.EntityData.ParentYangName = "pmipv6"
    lma.EntityData.SegmentPath = "lma"
    lma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lma.EntityData.Children = types.NewOrderedMap()
    lma.EntityData.Children.Append("statistics", types.YChild{"Statistics", &lma.Statistics})
    lma.EntityData.Children.Append("bindings", types.YChild{"Bindings", &lma.Bindings})
    lma.EntityData.Children.Append("heartbeats", types.YChild{"Heartbeats", &lma.Heartbeats})
    lma.EntityData.Children.Append("config-variables", types.YChild{"ConfigVariables", &lma.ConfigVariables})
    lma.EntityData.Leafs = types.NewOrderedMap()

    lma.EntityData.YListKeys = []string {}

    return &(lma.EntityData)
}

// Pmipv6_Lma_Statistics
// None
type Pmipv6_Lma_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *Pmipv6_Lma_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "lma"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("customer-statistics", types.YChild{"CustomerStatistics", &statistics.CustomerStatistics})
    statistics.EntityData.Children.Append("license", types.YChild{"License", &statistics.License})
    statistics.EntityData.Children.Append("global", types.YChild{"Global", &statistics.Global})
    statistics.EntityData.Children.Append("mag-statistics", types.YChild{"MagStatistics", &statistics.MagStatistics})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics
// Table of CustomerStatistics
type Pmipv6_Lma_Statistics_CustomerStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Customer statistics. The type is slice of
    // Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic.
    CustomerStatistic []*Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic
}

func (customerStatistics *Pmipv6_Lma_Statistics_CustomerStatistics) GetEntityData() *types.CommonEntityData {
    customerStatistics.EntityData.YFilter = customerStatistics.YFilter
    customerStatistics.EntityData.YangName = "customer-statistics"
    customerStatistics.EntityData.BundleName = "cisco_ios_xr"
    customerStatistics.EntityData.ParentYangName = "statistics"
    customerStatistics.EntityData.SegmentPath = "customer-statistics"
    customerStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customerStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customerStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customerStatistics.EntityData.Children = types.NewOrderedMap()
    customerStatistics.EntityData.Children.Append("customer-statistic", types.YChild{"CustomerStatistic", nil})
    for i := range customerStatistics.CustomerStatistic {
        customerStatistics.EntityData.Children.Append(types.GetSegmentPath(customerStatistics.CustomerStatistic[i]), types.YChild{"CustomerStatistic", customerStatistics.CustomerStatistic[i]})
    }
    customerStatistics.EntityData.Leafs = types.NewOrderedMap()

    customerStatistics.EntityData.YListKeys = []string {}

    return &(customerStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic
// Customer statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic struct {
    EntityData types.CommonEntityData
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

func (customerStatistic *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic) GetEntityData() *types.CommonEntityData {
    customerStatistic.EntityData.YFilter = customerStatistic.YFilter
    customerStatistic.EntityData.YangName = "customer-statistic"
    customerStatistic.EntityData.BundleName = "cisco_ios_xr"
    customerStatistic.EntityData.ParentYangName = "customer-statistics"
    customerStatistic.EntityData.SegmentPath = "customer-statistic" + types.AddKeyToken(customerStatistic.CustomerName, "customer-name")
    customerStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customerStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customerStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customerStatistic.EntityData.Children = types.NewOrderedMap()
    customerStatistic.EntityData.Children.Append("protocol-statistics", types.YChild{"ProtocolStatistics", &customerStatistic.ProtocolStatistics})
    customerStatistic.EntityData.Children.Append("accounting-statistics", types.YChild{"AccountingStatistics", &customerStatistic.AccountingStatistics})
    customerStatistic.EntityData.Leafs = types.NewOrderedMap()
    customerStatistic.EntityData.Leafs.Append("customer-name", types.YLeaf{"CustomerName", customerStatistic.CustomerName})
    customerStatistic.EntityData.Leafs.Append("lma-identifier", types.YLeaf{"LmaIdentifier", customerStatistic.LmaIdentifier})
    customerStatistic.EntityData.Leafs.Append("bce-count", types.YLeaf{"BceCount", customerStatistic.BceCount})
    customerStatistic.EntityData.Leafs.Append("handoff-count", types.YLeaf{"HandoffCount", customerStatistic.HandoffCount})
    customerStatistic.EntityData.Leafs.Append("ipv4-mnp-count", types.YLeaf{"Ipv4MnpCount", customerStatistic.Ipv4MnpCount})
    customerStatistic.EntityData.Leafs.Append("ipv6-mnp-count", types.YLeaf{"Ipv6MnpCount", customerStatistic.Ipv6MnpCount})

    customerStatistic.EntityData.YListKeys = []string {"CustomerName"}

    return &(customerStatistic.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics struct {
    EntityData types.CommonEntityData
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

func (protocolStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics) GetEntityData() *types.CommonEntityData {
    protocolStatistics.EntityData.YFilter = protocolStatistics.YFilter
    protocolStatistics.EntityData.YangName = "protocol-statistics"
    protocolStatistics.EntityData.BundleName = "cisco_ios_xr"
    protocolStatistics.EntityData.ParentYangName = "customer-statistic"
    protocolStatistics.EntityData.SegmentPath = "protocol-statistics"
    protocolStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolStatistics.EntityData.Children = types.NewOrderedMap()
    protocolStatistics.EntityData.Children.Append("pbu-receive-statistics", types.YChild{"PbuReceiveStatistics", &protocolStatistics.PbuReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pba-send-statistics", types.YChild{"PbaSendStatistics", &protocolStatistics.PbaSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-send-statistics", types.YChild{"PbriSendStatistics", &protocolStatistics.PbriSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-receive-statistics", types.YChild{"PbriReceiveStatistics", &protocolStatistics.PbriReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-send-statistics", types.YChild{"PbraSendStatistics", &protocolStatistics.PbraSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-receive-statistics", types.YChild{"PbraReceiveStatistics", &protocolStatistics.PbraReceiveStatistics})
    protocolStatistics.EntityData.Leafs = types.NewOrderedMap()

    protocolStatistics.EntityData.YListKeys = []string {}

    return &(protocolStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbuReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbuReceiveStatistics.EntityData.YFilter = pbuReceiveStatistics.YFilter
    pbuReceiveStatistics.EntityData.YangName = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbuReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbuReceiveStatistics.EntityData.SegmentPath = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbuReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbuReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbuReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-count", types.YLeaf{"PbuCount", pbuReceiveStatistics.PbuCount})
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-drop-count", types.YLeaf{"PbuDropCount", pbuReceiveStatistics.PbuDropCount})

    pbuReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbuReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbaSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbaSendStatistics) GetEntityData() *types.CommonEntityData {
    pbaSendStatistics.EntityData.YFilter = pbaSendStatistics.YFilter
    pbaSendStatistics.EntityData.YangName = "pba-send-statistics"
    pbaSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbaSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbaSendStatistics.EntityData.SegmentPath = "pba-send-statistics"
    pbaSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbaSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbaSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbaSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs.Append("pba-count", types.YLeaf{"PbaCount", pbaSendStatistics.PbaCount})
    pbaSendStatistics.EntityData.Leafs.Append("pba-drop-count", types.YLeaf{"PbaDropCount", pbaSendStatistics.PbaDropCount})
    pbaSendStatistics.EntityData.Leafs.Append("accepted-count", types.YLeaf{"AcceptedCount", pbaSendStatistics.AcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", pbaSendStatistics.UnknownCount})
    pbaSendStatistics.EntityData.Leafs.Append("unspecified-failure-count", types.YLeaf{"UnspecifiedFailureCount", pbaSendStatistics.UnspecifiedFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("admin-failure-count", types.YLeaf{"AdminFailureCount", pbaSendStatistics.AdminFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("resource-failure-count", types.YLeaf{"ResourceFailureCount", pbaSendStatistics.ResourceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-reg-failure-count", types.YLeaf{"HomeRegFailureCount", pbaSendStatistics.HomeRegFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-subnet-failure-count", types.YLeaf{"HomeSubnetFailureCount", pbaSendStatistics.HomeSubnetFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("bad-sequence-failure-count", types.YLeaf{"BadSequenceFailureCount", pbaSendStatistics.BadSequenceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("reg-type-failure-count", types.YLeaf{"RegTypeFailureCount", pbaSendStatistics.RegTypeFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("authen-failure-count", types.YLeaf{"AuthenFailureCount", pbaSendStatistics.AuthenFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("proxy-reg-not-enabled-count", types.YLeaf{"ProxyRegNotEnabledCount", pbaSendStatistics.ProxyRegNotEnabledCount})
    pbaSendStatistics.EntityData.Leafs.Append("not-lma-for-this-mn-count", types.YLeaf{"NotLmaForThisMnCount", pbaSendStatistics.NotLmaForThisMnCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-proxy-reg-count", types.YLeaf{"NoAuthorForProxyRegCount", pbaSendStatistics.NoAuthorForProxyRegCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-hnp-count", types.YLeaf{"NoAuthorForHnpCount", pbaSendStatistics.NoAuthorForHnpCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-mismatch-count", types.YLeaf{"TimestampMismatchCount", pbaSendStatistics.TimestampMismatchCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-lower-than-previous-accepted-count", types.YLeaf{"TimestampLowerThanPreviousAcceptedCount", pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hnp-opt-count", types.YLeaf{"MissingHnpOptCount", pbaSendStatistics.MissingHnpOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("received-hnps-do-not-match-bce-hnps-count", types.YLeaf{"ReceivedHnpsDoNotMatchBceHnpsCount", pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-mn-id-opt-count", types.YLeaf{"MissingMnIdOptCount", pbaSendStatistics.MissingMnIdOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hi-opt-count", types.YLeaf{"MissingHiOptCount", pbaSendStatistics.MissingHiOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-access-tech-type-opt-count", types.YLeaf{"MissingAccessTechTypeOptCount", pbaSendStatistics.MissingAccessTechTypeOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-mobility-count", types.YLeaf{"NoAuthorForIpv4MobilityCount", pbaSendStatistics.NoAuthorForIpv4MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-hoa-count", types.YLeaf{"NoAuthorForIpv4HoaCount", pbaSendStatistics.NoAuthorForIpv4HoaCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv6-mobility-count", types.YLeaf{"NoAuthorForIpv6MobilityCount", pbaSendStatistics.NoAuthorForIpv6MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("multiple-ipv4-ho-a-not-supported-count", types.YLeaf{"MultipleIpv4HoANotSupportedCount", pbaSendStatistics.MultipleIpv4HoANotSupportedCount})
    pbaSendStatistics.EntityData.Leafs.Append("gre-key-opt-required-count", types.YLeaf{"GreKeyOptRequiredCount", pbaSendStatistics.GreKeyOptRequiredCount})

    pbaSendStatistics.EntityData.YListKeys = []string {}

    return &(pbaSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriSendStatistics) GetEntityData() *types.CommonEntityData {
    pbriSendStatistics.EntityData.YFilter = pbriSendStatistics.YFilter
    pbriSendStatistics.EntityData.YangName = "pbri-send-statistics"
    pbriSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriSendStatistics.EntityData.SegmentPath = "pbri-send-statistics"
    pbriSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriSendStatistics.PbriCount})
    pbriSendStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriSendStatistics.PbriDropCount})
    pbriSendStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriSendStatistics.UnspecifiedCount})
    pbriSendStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriSendStatistics.AdminReasonCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriSendStatistics.MagHandoverSameAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriSendStatistics.MagHandoverDifferentAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriSendStatistics.MagHandoverUnknownCount})
    pbriSendStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriSendStatistics.UserSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriSendStatistics.NetworkSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriSendStatistics.OutOfSyncBceStateCount})
    pbriSendStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriSendStatistics.PerPeerPolicyCount})
    pbriSendStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriSendStatistics.RevokingMnLocalPolicyCount})

    pbriSendStatistics.EntityData.YListKeys = []string {}

    return &(pbriSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbriReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbriReceiveStatistics.EntityData.YFilter = pbriReceiveStatistics.YFilter
    pbriReceiveStatistics.EntityData.YangName = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriReceiveStatistics.EntityData.SegmentPath = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriReceiveStatistics.PbriCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriReceiveStatistics.PbriDropCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriReceiveStatistics.UnspecifiedCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriReceiveStatistics.AdminReasonCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriReceiveStatistics.MagHandoverSameAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriReceiveStatistics.MagHandoverDifferentAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriReceiveStatistics.MagHandoverUnknownCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriReceiveStatistics.UserSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriReceiveStatistics.NetworkSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriReceiveStatistics.OutOfSyncBceStateCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriReceiveStatistics.PerPeerPolicyCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriReceiveStatistics.RevokingMnLocalPolicyCount})

    pbriReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbriReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraSendStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraSendStatistics) GetEntityData() *types.CommonEntityData {
    pbraSendStatistics.EntityData.YFilter = pbraSendStatistics.YFilter
    pbraSendStatistics.EntityData.YangName = "pbra-send-statistics"
    pbraSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraSendStatistics.EntityData.SegmentPath = "pbra-send-statistics"
    pbraSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraSendStatistics.PbraCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraSendStatistics.PbraDropCount})
    pbraSendStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraSendStatistics.SuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraSendStatistics.PartialSuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraSendStatistics.NoBindingCount})
    pbraSendStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraSendStatistics.HoaRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraSendStatistics.NoAuthorForGlobalRevocCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraSendStatistics.MnIdentityRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraSendStatistics.MnAttachedCount})
    pbraSendStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraSendStatistics.UnknownRevocTriggerCount})
    pbraSendStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraSendStatistics.RevocFunctionNotSupportedCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraSendStatistics.PbrNotSupportedCount})

    pbraSendStatistics.EntityData.YListKeys = []string {}

    return &(pbraSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_ProtocolStatistics_PbraReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbraReceiveStatistics.EntityData.YFilter = pbraReceiveStatistics.YFilter
    pbraReceiveStatistics.EntityData.YangName = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraReceiveStatistics.EntityData.SegmentPath = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraReceiveStatistics.PbraCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraReceiveStatistics.PbraDropCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraReceiveStatistics.SuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraReceiveStatistics.PartialSuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraReceiveStatistics.NoBindingCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraReceiveStatistics.HoaRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraReceiveStatistics.NoAuthorForGlobalRevocCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraReceiveStatistics.MnIdentityRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraReceiveStatistics.MnAttachedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraReceiveStatistics.UnknownRevocTriggerCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraReceiveStatistics.RevocFunctionNotSupportedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraReceiveStatistics.PbrNotSupportedCount})

    pbraReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbraReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics
// LMA Accounting Statistics
type Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics struct {
    EntityData types.CommonEntityData
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

func (accountingStatistics *Pmipv6_Lma_Statistics_CustomerStatistics_CustomerStatistic_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "customer-statistic"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs.Append("accounting-start-sent-count", types.YLeaf{"AccountingStartSentCount", accountingStatistics.AccountingStartSentCount})
    accountingStatistics.EntityData.Leafs.Append("accounting-update-sent-count", types.YLeaf{"AccountingUpdateSentCount", accountingStatistics.AccountingUpdateSentCount})
    accountingStatistics.EntityData.Leafs.Append("accounting-stop-sent-count", types.YLeaf{"AccountingStopSentCount", accountingStatistics.AccountingStopSentCount})

    accountingStatistics.EntityData.YListKeys = []string {}

    return &(accountingStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_License
// LMA License Statistics
type Pmipv6_Lma_Statistics_License struct {
    EntityData types.CommonEntityData
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

func (license *Pmipv6_Lma_Statistics_License) GetEntityData() *types.CommonEntityData {
    license.EntityData.YFilter = license.YFilter
    license.EntityData.YangName = "license"
    license.EntityData.BundleName = "cisco_ios_xr"
    license.EntityData.ParentYangName = "statistics"
    license.EntityData.SegmentPath = "license"
    license.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    license.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    license.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    license.EntityData.Children = types.NewOrderedMap()
    license.EntityData.Leafs = types.NewOrderedMap()
    license.EntityData.Leafs.Append("lma-identifier", types.YLeaf{"LmaIdentifier", license.LmaIdentifier})
    license.EntityData.Leafs.Append("bce-count", types.YLeaf{"BceCount", license.BceCount})
    license.EntityData.Leafs.Append("peak-bce-count", types.YLeaf{"PeakBceCount", license.PeakBceCount})
    license.EntityData.Leafs.Append("peak-bce-count-reset-timestamp", types.YLeaf{"PeakBceCountResetTimestamp", license.PeakBceCountResetTimestamp})

    license.EntityData.YListKeys = []string {}

    return &(license.EntityData)
}

// Pmipv6_Lma_Statistics_Global
// Global Statistics
type Pmipv6_Lma_Statistics_Global struct {
    EntityData types.CommonEntityData
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

func (global *Pmipv6_Lma_Statistics_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "statistics"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = types.NewOrderedMap()
    global.EntityData.Children.Append("packet-statistics", types.YChild{"PacketStatistics", &global.PacketStatistics})
    global.EntityData.Children.Append("protocol-statistics", types.YChild{"ProtocolStatistics", &global.ProtocolStatistics})
    global.EntityData.Children.Append("accounting-statistics", types.YChild{"AccountingStatistics", &global.AccountingStatistics})
    global.EntityData.Leafs = types.NewOrderedMap()
    global.EntityData.Leafs.Append("lma-identifier", types.YLeaf{"LmaIdentifier", global.LmaIdentifier})
    global.EntityData.Leafs.Append("bce-count", types.YLeaf{"BceCount", global.BceCount})
    global.EntityData.Leafs.Append("handoff-count", types.YLeaf{"HandoffCount", global.HandoffCount})
    global.EntityData.Leafs.Append("single-tenant-count", types.YLeaf{"SingleTenantCount", global.SingleTenantCount})
    global.EntityData.Leafs.Append("multi-tenant-count", types.YLeaf{"MultiTenantCount", global.MultiTenantCount})

    global.EntityData.YListKeys = []string {}

    return &(global.EntityData)
}

// Pmipv6_Lma_Statistics_Global_PacketStatistics
// Packet Statistics
type Pmipv6_Lma_Statistics_Global_PacketStatistics struct {
    EntityData types.CommonEntityData
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

func (packetStatistics *Pmipv6_Lma_Statistics_Global_PacketStatistics) GetEntityData() *types.CommonEntityData {
    packetStatistics.EntityData.YFilter = packetStatistics.YFilter
    packetStatistics.EntityData.YangName = "packet-statistics"
    packetStatistics.EntityData.BundleName = "cisco_ios_xr"
    packetStatistics.EntityData.ParentYangName = "global"
    packetStatistics.EntityData.SegmentPath = "packet-statistics"
    packetStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetStatistics.EntityData.Children = types.NewOrderedMap()
    packetStatistics.EntityData.Leafs = types.NewOrderedMap()
    packetStatistics.EntityData.Leafs.Append("checksum-errors", types.YLeaf{"ChecksumErrors", packetStatistics.ChecksumErrors})
    packetStatistics.EntityData.Leafs.Append("send-drops", types.YLeaf{"SendDrops", packetStatistics.SendDrops})
    packetStatistics.EntityData.Leafs.Append("receive-drops", types.YLeaf{"ReceiveDrops", packetStatistics.ReceiveDrops})
    packetStatistics.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", packetStatistics.PacketsReceived})
    packetStatistics.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", packetStatistics.PacketsSent})
    packetStatistics.EntityData.Leafs.Append("send-drops-ipv6", types.YLeaf{"SendDropsIpv6", packetStatistics.SendDropsIpv6})
    packetStatistics.EntityData.Leafs.Append("receive-drops-ipv6", types.YLeaf{"ReceiveDropsIpv6", packetStatistics.ReceiveDropsIpv6})
    packetStatistics.EntityData.Leafs.Append("packets-received-ipv6", types.YLeaf{"PacketsReceivedIpv6", packetStatistics.PacketsReceivedIpv6})
    packetStatistics.EntityData.Leafs.Append("packets-sent-ipv6", types.YLeaf{"PacketsSentIpv6", packetStatistics.PacketsSentIpv6})

    packetStatistics.EntityData.YListKeys = []string {}

    return &(packetStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics struct {
    EntityData types.CommonEntityData
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

func (protocolStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics) GetEntityData() *types.CommonEntityData {
    protocolStatistics.EntityData.YFilter = protocolStatistics.YFilter
    protocolStatistics.EntityData.YangName = "protocol-statistics"
    protocolStatistics.EntityData.BundleName = "cisco_ios_xr"
    protocolStatistics.EntityData.ParentYangName = "global"
    protocolStatistics.EntityData.SegmentPath = "protocol-statistics"
    protocolStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolStatistics.EntityData.Children = types.NewOrderedMap()
    protocolStatistics.EntityData.Children.Append("pbu-receive-statistics", types.YChild{"PbuReceiveStatistics", &protocolStatistics.PbuReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pba-send-statistics", types.YChild{"PbaSendStatistics", &protocolStatistics.PbaSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-send-statistics", types.YChild{"PbriSendStatistics", &protocolStatistics.PbriSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-receive-statistics", types.YChild{"PbriReceiveStatistics", &protocolStatistics.PbriReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-send-statistics", types.YChild{"PbraSendStatistics", &protocolStatistics.PbraSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-receive-statistics", types.YChild{"PbraReceiveStatistics", &protocolStatistics.PbraReceiveStatistics})
    protocolStatistics.EntityData.Leafs = types.NewOrderedMap()

    protocolStatistics.EntityData.YListKeys = []string {}

    return &(protocolStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbuReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbuReceiveStatistics.EntityData.YFilter = pbuReceiveStatistics.YFilter
    pbuReceiveStatistics.EntityData.YangName = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbuReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbuReceiveStatistics.EntityData.SegmentPath = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbuReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbuReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbuReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-count", types.YLeaf{"PbuCount", pbuReceiveStatistics.PbuCount})
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-drop-count", types.YLeaf{"PbuDropCount", pbuReceiveStatistics.PbuDropCount})

    pbuReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbuReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbaSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbaSendStatistics) GetEntityData() *types.CommonEntityData {
    pbaSendStatistics.EntityData.YFilter = pbaSendStatistics.YFilter
    pbaSendStatistics.EntityData.YangName = "pba-send-statistics"
    pbaSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbaSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbaSendStatistics.EntityData.SegmentPath = "pba-send-statistics"
    pbaSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbaSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbaSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbaSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs.Append("pba-count", types.YLeaf{"PbaCount", pbaSendStatistics.PbaCount})
    pbaSendStatistics.EntityData.Leafs.Append("pba-drop-count", types.YLeaf{"PbaDropCount", pbaSendStatistics.PbaDropCount})
    pbaSendStatistics.EntityData.Leafs.Append("accepted-count", types.YLeaf{"AcceptedCount", pbaSendStatistics.AcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", pbaSendStatistics.UnknownCount})
    pbaSendStatistics.EntityData.Leafs.Append("unspecified-failure-count", types.YLeaf{"UnspecifiedFailureCount", pbaSendStatistics.UnspecifiedFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("admin-failure-count", types.YLeaf{"AdminFailureCount", pbaSendStatistics.AdminFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("resource-failure-count", types.YLeaf{"ResourceFailureCount", pbaSendStatistics.ResourceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-reg-failure-count", types.YLeaf{"HomeRegFailureCount", pbaSendStatistics.HomeRegFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-subnet-failure-count", types.YLeaf{"HomeSubnetFailureCount", pbaSendStatistics.HomeSubnetFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("bad-sequence-failure-count", types.YLeaf{"BadSequenceFailureCount", pbaSendStatistics.BadSequenceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("reg-type-failure-count", types.YLeaf{"RegTypeFailureCount", pbaSendStatistics.RegTypeFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("authen-failure-count", types.YLeaf{"AuthenFailureCount", pbaSendStatistics.AuthenFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("proxy-reg-not-enabled-count", types.YLeaf{"ProxyRegNotEnabledCount", pbaSendStatistics.ProxyRegNotEnabledCount})
    pbaSendStatistics.EntityData.Leafs.Append("not-lma-for-this-mn-count", types.YLeaf{"NotLmaForThisMnCount", pbaSendStatistics.NotLmaForThisMnCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-proxy-reg-count", types.YLeaf{"NoAuthorForProxyRegCount", pbaSendStatistics.NoAuthorForProxyRegCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-hnp-count", types.YLeaf{"NoAuthorForHnpCount", pbaSendStatistics.NoAuthorForHnpCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-mismatch-count", types.YLeaf{"TimestampMismatchCount", pbaSendStatistics.TimestampMismatchCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-lower-than-previous-accepted-count", types.YLeaf{"TimestampLowerThanPreviousAcceptedCount", pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hnp-opt-count", types.YLeaf{"MissingHnpOptCount", pbaSendStatistics.MissingHnpOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("received-hnps-do-not-match-bce-hnps-count", types.YLeaf{"ReceivedHnpsDoNotMatchBceHnpsCount", pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-mn-id-opt-count", types.YLeaf{"MissingMnIdOptCount", pbaSendStatistics.MissingMnIdOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hi-opt-count", types.YLeaf{"MissingHiOptCount", pbaSendStatistics.MissingHiOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-access-tech-type-opt-count", types.YLeaf{"MissingAccessTechTypeOptCount", pbaSendStatistics.MissingAccessTechTypeOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-mobility-count", types.YLeaf{"NoAuthorForIpv4MobilityCount", pbaSendStatistics.NoAuthorForIpv4MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-hoa-count", types.YLeaf{"NoAuthorForIpv4HoaCount", pbaSendStatistics.NoAuthorForIpv4HoaCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv6-mobility-count", types.YLeaf{"NoAuthorForIpv6MobilityCount", pbaSendStatistics.NoAuthorForIpv6MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("multiple-ipv4-ho-a-not-supported-count", types.YLeaf{"MultipleIpv4HoANotSupportedCount", pbaSendStatistics.MultipleIpv4HoANotSupportedCount})
    pbaSendStatistics.EntityData.Leafs.Append("gre-key-opt-required-count", types.YLeaf{"GreKeyOptRequiredCount", pbaSendStatistics.GreKeyOptRequiredCount})

    pbaSendStatistics.EntityData.YListKeys = []string {}

    return &(pbaSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriSendStatistics) GetEntityData() *types.CommonEntityData {
    pbriSendStatistics.EntityData.YFilter = pbriSendStatistics.YFilter
    pbriSendStatistics.EntityData.YangName = "pbri-send-statistics"
    pbriSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriSendStatistics.EntityData.SegmentPath = "pbri-send-statistics"
    pbriSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriSendStatistics.PbriCount})
    pbriSendStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriSendStatistics.PbriDropCount})
    pbriSendStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriSendStatistics.UnspecifiedCount})
    pbriSendStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriSendStatistics.AdminReasonCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriSendStatistics.MagHandoverSameAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriSendStatistics.MagHandoverDifferentAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriSendStatistics.MagHandoverUnknownCount})
    pbriSendStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriSendStatistics.UserSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriSendStatistics.NetworkSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriSendStatistics.OutOfSyncBceStateCount})
    pbriSendStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriSendStatistics.PerPeerPolicyCount})
    pbriSendStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriSendStatistics.RevokingMnLocalPolicyCount})

    pbriSendStatistics.EntityData.YListKeys = []string {}

    return &(pbriSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbriReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbriReceiveStatistics.EntityData.YFilter = pbriReceiveStatistics.YFilter
    pbriReceiveStatistics.EntityData.YangName = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriReceiveStatistics.EntityData.SegmentPath = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriReceiveStatistics.PbriCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriReceiveStatistics.PbriDropCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriReceiveStatistics.UnspecifiedCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriReceiveStatistics.AdminReasonCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriReceiveStatistics.MagHandoverSameAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriReceiveStatistics.MagHandoverDifferentAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriReceiveStatistics.MagHandoverUnknownCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriReceiveStatistics.UserSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriReceiveStatistics.NetworkSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriReceiveStatistics.OutOfSyncBceStateCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriReceiveStatistics.PerPeerPolicyCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriReceiveStatistics.RevokingMnLocalPolicyCount})

    pbriReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbriReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraSendStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraSendStatistics) GetEntityData() *types.CommonEntityData {
    pbraSendStatistics.EntityData.YFilter = pbraSendStatistics.YFilter
    pbraSendStatistics.EntityData.YangName = "pbra-send-statistics"
    pbraSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraSendStatistics.EntityData.SegmentPath = "pbra-send-statistics"
    pbraSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraSendStatistics.PbraCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraSendStatistics.PbraDropCount})
    pbraSendStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraSendStatistics.SuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraSendStatistics.PartialSuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraSendStatistics.NoBindingCount})
    pbraSendStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraSendStatistics.HoaRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraSendStatistics.NoAuthorForGlobalRevocCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraSendStatistics.MnIdentityRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraSendStatistics.MnAttachedCount})
    pbraSendStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraSendStatistics.UnknownRevocTriggerCount})
    pbraSendStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraSendStatistics.RevocFunctionNotSupportedCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraSendStatistics.PbrNotSupportedCount})

    pbraSendStatistics.EntityData.YListKeys = []string {}

    return &(pbraSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_Global_ProtocolStatistics_PbraReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbraReceiveStatistics.EntityData.YFilter = pbraReceiveStatistics.YFilter
    pbraReceiveStatistics.EntityData.YangName = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraReceiveStatistics.EntityData.SegmentPath = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraReceiveStatistics.PbraCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraReceiveStatistics.PbraDropCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraReceiveStatistics.SuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraReceiveStatistics.PartialSuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraReceiveStatistics.NoBindingCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraReceiveStatistics.HoaRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraReceiveStatistics.NoAuthorForGlobalRevocCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraReceiveStatistics.MnIdentityRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraReceiveStatistics.MnAttachedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraReceiveStatistics.UnknownRevocTriggerCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraReceiveStatistics.RevocFunctionNotSupportedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraReceiveStatistics.PbrNotSupportedCount})

    pbraReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbraReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_Global_AccountingStatistics
// LMA Accounting Statistics
type Pmipv6_Lma_Statistics_Global_AccountingStatistics struct {
    EntityData types.CommonEntityData
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

func (accountingStatistics *Pmipv6_Lma_Statistics_Global_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "global"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs = types.NewOrderedMap()
    accountingStatistics.EntityData.Leafs.Append("accounting-start-sent-count", types.YLeaf{"AccountingStartSentCount", accountingStatistics.AccountingStartSentCount})
    accountingStatistics.EntityData.Leafs.Append("accounting-update-sent-count", types.YLeaf{"AccountingUpdateSentCount", accountingStatistics.AccountingUpdateSentCount})
    accountingStatistics.EntityData.Leafs.Append("accounting-stop-sent-count", types.YLeaf{"AccountingStopSentCount", accountingStatistics.AccountingStopSentCount})

    accountingStatistics.EntityData.YListKeys = []string {}

    return &(accountingStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics
// Table of MAGStatistics
type Pmipv6_Lma_Statistics_MagStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer MAG statistics. The type is slice of
    // Pmipv6_Lma_Statistics_MagStatistics_MagStatistic.
    MagStatistic []*Pmipv6_Lma_Statistics_MagStatistics_MagStatistic
}

func (magStatistics *Pmipv6_Lma_Statistics_MagStatistics) GetEntityData() *types.CommonEntityData {
    magStatistics.EntityData.YFilter = magStatistics.YFilter
    magStatistics.EntityData.YangName = "mag-statistics"
    magStatistics.EntityData.BundleName = "cisco_ios_xr"
    magStatistics.EntityData.ParentYangName = "statistics"
    magStatistics.EntityData.SegmentPath = "mag-statistics"
    magStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    magStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    magStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    magStatistics.EntityData.Children = types.NewOrderedMap()
    magStatistics.EntityData.Children.Append("mag-statistic", types.YChild{"MagStatistic", nil})
    for i := range magStatistics.MagStatistic {
        magStatistics.EntityData.Children.Append(types.GetSegmentPath(magStatistics.MagStatistic[i]), types.YChild{"MagStatistic", magStatistics.MagStatistic[i]})
    }
    magStatistics.EntityData.Leafs = types.NewOrderedMap()

    magStatistics.EntityData.YListKeys = []string {}

    return &(magStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic
// Peer MAG statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer MAG Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    MagName interface{}

    // LMA Identifier. The type is string.
    LmaIdentifier interface{}

    // LMA Protocol Statistics.
    ProtocolStatistics Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics
}

func (magStatistic *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic) GetEntityData() *types.CommonEntityData {
    magStatistic.EntityData.YFilter = magStatistic.YFilter
    magStatistic.EntityData.YangName = "mag-statistic"
    magStatistic.EntityData.BundleName = "cisco_ios_xr"
    magStatistic.EntityData.ParentYangName = "mag-statistics"
    magStatistic.EntityData.SegmentPath = "mag-statistic" + types.AddKeyToken(magStatistic.MagName, "mag-name")
    magStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    magStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    magStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    magStatistic.EntityData.Children = types.NewOrderedMap()
    magStatistic.EntityData.Children.Append("protocol-statistics", types.YChild{"ProtocolStatistics", &magStatistic.ProtocolStatistics})
    magStatistic.EntityData.Leafs = types.NewOrderedMap()
    magStatistic.EntityData.Leafs.Append("mag-name", types.YLeaf{"MagName", magStatistic.MagName})
    magStatistic.EntityData.Leafs.Append("lma-identifier", types.YLeaf{"LmaIdentifier", magStatistic.LmaIdentifier})

    magStatistic.EntityData.YListKeys = []string {"MagName"}

    return &(magStatistic.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics
// LMA Protocol Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics struct {
    EntityData types.CommonEntityData
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

func (protocolStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics) GetEntityData() *types.CommonEntityData {
    protocolStatistics.EntityData.YFilter = protocolStatistics.YFilter
    protocolStatistics.EntityData.YangName = "protocol-statistics"
    protocolStatistics.EntityData.BundleName = "cisco_ios_xr"
    protocolStatistics.EntityData.ParentYangName = "mag-statistic"
    protocolStatistics.EntityData.SegmentPath = "protocol-statistics"
    protocolStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protocolStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protocolStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protocolStatistics.EntityData.Children = types.NewOrderedMap()
    protocolStatistics.EntityData.Children.Append("pbu-receive-statistics", types.YChild{"PbuReceiveStatistics", &protocolStatistics.PbuReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pba-send-statistics", types.YChild{"PbaSendStatistics", &protocolStatistics.PbaSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-send-statistics", types.YChild{"PbriSendStatistics", &protocolStatistics.PbriSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbri-receive-statistics", types.YChild{"PbriReceiveStatistics", &protocolStatistics.PbriReceiveStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-send-statistics", types.YChild{"PbraSendStatistics", &protocolStatistics.PbraSendStatistics})
    protocolStatistics.EntityData.Children.Append("pbra-receive-statistics", types.YChild{"PbraReceiveStatistics", &protocolStatistics.PbraReceiveStatistics})
    protocolStatistics.EntityData.Leafs = types.NewOrderedMap()

    protocolStatistics.EntityData.YListKeys = []string {}

    return &(protocolStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics
// PBU Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of PBUs. The type is interface{} with range: 0..18446744073709551615.
    PbuCount interface{}

    // Count of PBUs Dropped. The type is interface{} with range: 0..4294967295.
    PbuDropCount interface{}
}

func (pbuReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbuReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbuReceiveStatistics.EntityData.YFilter = pbuReceiveStatistics.YFilter
    pbuReceiveStatistics.EntityData.YangName = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbuReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbuReceiveStatistics.EntityData.SegmentPath = "pbu-receive-statistics"
    pbuReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbuReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbuReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbuReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-count", types.YLeaf{"PbuCount", pbuReceiveStatistics.PbuCount})
    pbuReceiveStatistics.EntityData.Leafs.Append("pbu-drop-count", types.YLeaf{"PbuDropCount", pbuReceiveStatistics.PbuDropCount})

    pbuReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbuReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics
// PBA Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbaSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbaSendStatistics) GetEntityData() *types.CommonEntityData {
    pbaSendStatistics.EntityData.YFilter = pbaSendStatistics.YFilter
    pbaSendStatistics.EntityData.YangName = "pba-send-statistics"
    pbaSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbaSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbaSendStatistics.EntityData.SegmentPath = "pba-send-statistics"
    pbaSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbaSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbaSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbaSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbaSendStatistics.EntityData.Leafs.Append("pba-count", types.YLeaf{"PbaCount", pbaSendStatistics.PbaCount})
    pbaSendStatistics.EntityData.Leafs.Append("pba-drop-count", types.YLeaf{"PbaDropCount", pbaSendStatistics.PbaDropCount})
    pbaSendStatistics.EntityData.Leafs.Append("accepted-count", types.YLeaf{"AcceptedCount", pbaSendStatistics.AcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", pbaSendStatistics.UnknownCount})
    pbaSendStatistics.EntityData.Leafs.Append("unspecified-failure-count", types.YLeaf{"UnspecifiedFailureCount", pbaSendStatistics.UnspecifiedFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("admin-failure-count", types.YLeaf{"AdminFailureCount", pbaSendStatistics.AdminFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("resource-failure-count", types.YLeaf{"ResourceFailureCount", pbaSendStatistics.ResourceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-reg-failure-count", types.YLeaf{"HomeRegFailureCount", pbaSendStatistics.HomeRegFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("home-subnet-failure-count", types.YLeaf{"HomeSubnetFailureCount", pbaSendStatistics.HomeSubnetFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("bad-sequence-failure-count", types.YLeaf{"BadSequenceFailureCount", pbaSendStatistics.BadSequenceFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("reg-type-failure-count", types.YLeaf{"RegTypeFailureCount", pbaSendStatistics.RegTypeFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("authen-failure-count", types.YLeaf{"AuthenFailureCount", pbaSendStatistics.AuthenFailureCount})
    pbaSendStatistics.EntityData.Leafs.Append("proxy-reg-not-enabled-count", types.YLeaf{"ProxyRegNotEnabledCount", pbaSendStatistics.ProxyRegNotEnabledCount})
    pbaSendStatistics.EntityData.Leafs.Append("not-lma-for-this-mn-count", types.YLeaf{"NotLmaForThisMnCount", pbaSendStatistics.NotLmaForThisMnCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-proxy-reg-count", types.YLeaf{"NoAuthorForProxyRegCount", pbaSendStatistics.NoAuthorForProxyRegCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-hnp-count", types.YLeaf{"NoAuthorForHnpCount", pbaSendStatistics.NoAuthorForHnpCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-mismatch-count", types.YLeaf{"TimestampMismatchCount", pbaSendStatistics.TimestampMismatchCount})
    pbaSendStatistics.EntityData.Leafs.Append("timestamp-lower-than-previous-accepted-count", types.YLeaf{"TimestampLowerThanPreviousAcceptedCount", pbaSendStatistics.TimestampLowerThanPreviousAcceptedCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hnp-opt-count", types.YLeaf{"MissingHnpOptCount", pbaSendStatistics.MissingHnpOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("received-hnps-do-not-match-bce-hnps-count", types.YLeaf{"ReceivedHnpsDoNotMatchBceHnpsCount", pbaSendStatistics.ReceivedHnpsDoNotMatchBceHnpsCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-mn-id-opt-count", types.YLeaf{"MissingMnIdOptCount", pbaSendStatistics.MissingMnIdOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-hi-opt-count", types.YLeaf{"MissingHiOptCount", pbaSendStatistics.MissingHiOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("missing-access-tech-type-opt-count", types.YLeaf{"MissingAccessTechTypeOptCount", pbaSendStatistics.MissingAccessTechTypeOptCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-mobility-count", types.YLeaf{"NoAuthorForIpv4MobilityCount", pbaSendStatistics.NoAuthorForIpv4MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv4-hoa-count", types.YLeaf{"NoAuthorForIpv4HoaCount", pbaSendStatistics.NoAuthorForIpv4HoaCount})
    pbaSendStatistics.EntityData.Leafs.Append("no-author-for-ipv6-mobility-count", types.YLeaf{"NoAuthorForIpv6MobilityCount", pbaSendStatistics.NoAuthorForIpv6MobilityCount})
    pbaSendStatistics.EntityData.Leafs.Append("multiple-ipv4-ho-a-not-supported-count", types.YLeaf{"MultipleIpv4HoANotSupportedCount", pbaSendStatistics.MultipleIpv4HoANotSupportedCount})
    pbaSendStatistics.EntityData.Leafs.Append("gre-key-opt-required-count", types.YLeaf{"GreKeyOptRequiredCount", pbaSendStatistics.GreKeyOptRequiredCount})

    pbaSendStatistics.EntityData.YListKeys = []string {}

    return &(pbaSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics
// PBRI Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriSendStatistics) GetEntityData() *types.CommonEntityData {
    pbriSendStatistics.EntityData.YFilter = pbriSendStatistics.YFilter
    pbriSendStatistics.EntityData.YangName = "pbri-send-statistics"
    pbriSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriSendStatistics.EntityData.SegmentPath = "pbri-send-statistics"
    pbriSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriSendStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriSendStatistics.PbriCount})
    pbriSendStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriSendStatistics.PbriDropCount})
    pbriSendStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriSendStatistics.UnspecifiedCount})
    pbriSendStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriSendStatistics.AdminReasonCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriSendStatistics.MagHandoverSameAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriSendStatistics.MagHandoverDifferentAttCount})
    pbriSendStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriSendStatistics.MagHandoverUnknownCount})
    pbriSendStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriSendStatistics.UserSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriSendStatistics.NetworkSessionTerminationCount})
    pbriSendStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriSendStatistics.OutOfSyncBceStateCount})
    pbriSendStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriSendStatistics.PerPeerPolicyCount})
    pbriSendStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriSendStatistics.RevokingMnLocalPolicyCount})

    pbriSendStatistics.EntityData.YListKeys = []string {}

    return &(pbriSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics
// PBRI Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbriReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbriReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbriReceiveStatistics.EntityData.YFilter = pbriReceiveStatistics.YFilter
    pbriReceiveStatistics.EntityData.YangName = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbriReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbriReceiveStatistics.EntityData.SegmentPath = "pbri-receive-statistics"
    pbriReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbriReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbriReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbriReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-count", types.YLeaf{"PbriCount", pbriReceiveStatistics.PbriCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("pbri-drop-count", types.YLeaf{"PbriDropCount", pbriReceiveStatistics.PbriDropCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("unspecified-count", types.YLeaf{"UnspecifiedCount", pbriReceiveStatistics.UnspecifiedCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("admin-reason-count", types.YLeaf{"AdminReasonCount", pbriReceiveStatistics.AdminReasonCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-same-att-count", types.YLeaf{"MagHandoverSameAttCount", pbriReceiveStatistics.MagHandoverSameAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-different-att-count", types.YLeaf{"MagHandoverDifferentAttCount", pbriReceiveStatistics.MagHandoverDifferentAttCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("mag-handover-unknown-count", types.YLeaf{"MagHandoverUnknownCount", pbriReceiveStatistics.MagHandoverUnknownCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("user-session-termination-count", types.YLeaf{"UserSessionTerminationCount", pbriReceiveStatistics.UserSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("network-session-termination-count", types.YLeaf{"NetworkSessionTerminationCount", pbriReceiveStatistics.NetworkSessionTerminationCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("out-of-sync-bce-state-count", types.YLeaf{"OutOfSyncBceStateCount", pbriReceiveStatistics.OutOfSyncBceStateCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("per-peer-policy-count", types.YLeaf{"PerPeerPolicyCount", pbriReceiveStatistics.PerPeerPolicyCount})
    pbriReceiveStatistics.EntityData.Leafs.Append("revoking-mn-local-policy-count", types.YLeaf{"RevokingMnLocalPolicyCount", pbriReceiveStatistics.RevokingMnLocalPolicyCount})

    pbriReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbriReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics
// PBRA Send Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraSendStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraSendStatistics) GetEntityData() *types.CommonEntityData {
    pbraSendStatistics.EntityData.YFilter = pbraSendStatistics.YFilter
    pbraSendStatistics.EntityData.YangName = "pbra-send-statistics"
    pbraSendStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraSendStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraSendStatistics.EntityData.SegmentPath = "pbra-send-statistics"
    pbraSendStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraSendStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraSendStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraSendStatistics.EntityData.Children = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraSendStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraSendStatistics.PbraCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraSendStatistics.PbraDropCount})
    pbraSendStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraSendStatistics.SuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraSendStatistics.PartialSuccessCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraSendStatistics.NoBindingCount})
    pbraSendStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraSendStatistics.HoaRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraSendStatistics.NoAuthorForGlobalRevocCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraSendStatistics.MnIdentityRequiredCount})
    pbraSendStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraSendStatistics.MnAttachedCount})
    pbraSendStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraSendStatistics.UnknownRevocTriggerCount})
    pbraSendStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraSendStatistics.RevocFunctionNotSupportedCount})
    pbraSendStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraSendStatistics.PbrNotSupportedCount})

    pbraSendStatistics.EntityData.YListKeys = []string {}

    return &(pbraSendStatistics.EntityData)
}

// Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics
// PBRA Receive Statistics
type Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics struct {
    EntityData types.CommonEntityData
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

func (pbraReceiveStatistics *Pmipv6_Lma_Statistics_MagStatistics_MagStatistic_ProtocolStatistics_PbraReceiveStatistics) GetEntityData() *types.CommonEntityData {
    pbraReceiveStatistics.EntityData.YFilter = pbraReceiveStatistics.YFilter
    pbraReceiveStatistics.EntityData.YangName = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbraReceiveStatistics.EntityData.ParentYangName = "protocol-statistics"
    pbraReceiveStatistics.EntityData.SegmentPath = "pbra-receive-statistics"
    pbraReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbraReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbraReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbraReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-count", types.YLeaf{"PbraCount", pbraReceiveStatistics.PbraCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbra-drop-count", types.YLeaf{"PbraDropCount", pbraReceiveStatistics.PbraDropCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("success-count", types.YLeaf{"SuccessCount", pbraReceiveStatistics.SuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("partial-success-count", types.YLeaf{"PartialSuccessCount", pbraReceiveStatistics.PartialSuccessCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-binding-count", types.YLeaf{"NoBindingCount", pbraReceiveStatistics.NoBindingCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("hoa-required-count", types.YLeaf{"HoaRequiredCount", pbraReceiveStatistics.HoaRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("no-author-for-global-revoc-count", types.YLeaf{"NoAuthorForGlobalRevocCount", pbraReceiveStatistics.NoAuthorForGlobalRevocCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-identity-required-count", types.YLeaf{"MnIdentityRequiredCount", pbraReceiveStatistics.MnIdentityRequiredCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("mn-attached-count", types.YLeaf{"MnAttachedCount", pbraReceiveStatistics.MnAttachedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("unknown-revoc-trigger-count", types.YLeaf{"UnknownRevocTriggerCount", pbraReceiveStatistics.UnknownRevocTriggerCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("revoc-function-not-supported-count", types.YLeaf{"RevocFunctionNotSupportedCount", pbraReceiveStatistics.RevocFunctionNotSupportedCount})
    pbraReceiveStatistics.EntityData.Leafs.Append("pbr-not-supported-count", types.YLeaf{"PbrNotSupportedCount", pbraReceiveStatistics.PbrNotSupportedCount})

    pbraReceiveStatistics.EntityData.YListKeys = []string {}

    return &(pbraReceiveStatistics.EntityData)
}

// Pmipv6_Lma_Bindings
// Table of Binding
type Pmipv6_Lma_Bindings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding Parameters. The type is slice of Pmipv6_Lma_Bindings_Binding.
    Binding []*Pmipv6_Lma_Bindings_Binding
}

func (bindings *Pmipv6_Lma_Bindings) GetEntityData() *types.CommonEntityData {
    bindings.EntityData.YFilter = bindings.YFilter
    bindings.EntityData.YangName = "bindings"
    bindings.EntityData.BundleName = "cisco_ios_xr"
    bindings.EntityData.ParentYangName = "lma"
    bindings.EntityData.SegmentPath = "bindings"
    bindings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindings.EntityData.Children = types.NewOrderedMap()
    bindings.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range bindings.Binding {
        bindings.EntityData.Children.Append(types.GetSegmentPath(bindings.Binding[i]), types.YChild{"Binding", bindings.Binding[i]})
    }
    bindings.EntityData.Leafs = types.NewOrderedMap()

    bindings.EntityData.YListKeys = []string {}

    return &(bindings.EntityData)
}

// Pmipv6_Lma_Bindings_Binding
// Binding Parameters
type Pmipv6_Lma_Bindings_Binding struct {
    EntityData types.CommonEntityData
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
    Coa []*Pmipv6_Lma_Bindings_Binding_Coa

    // IPv4 DMNP prefixes. The type is slice of
    // Pmipv6_Lma_Bindings_Binding_DmnpV4.
    DmnpV4 []*Pmipv6_Lma_Bindings_Binding_DmnpV4

    // IPv6 DMNP prefixes. The type is slice of
    // Pmipv6_Lma_Bindings_Binding_DmnpV6.
    DmnpV6 []*Pmipv6_Lma_Bindings_Binding_DmnpV6
}

func (binding *Pmipv6_Lma_Bindings_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xr"
    binding.EntityData.ParentYangName = "bindings"
    binding.EntityData.SegmentPath = "binding"
    binding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Children.Append("coa", types.YChild{"Coa", nil})
    for i := range binding.Coa {
        binding.EntityData.Children.Append(types.GetSegmentPath(binding.Coa[i]), types.YChild{"Coa", binding.Coa[i]})
    }
    binding.EntityData.Children.Append("dmnp-v4", types.YChild{"DmnpV4", nil})
    for i := range binding.DmnpV4 {
        binding.EntityData.Children.Append(types.GetSegmentPath(binding.DmnpV4[i]), types.YChild{"DmnpV4", binding.DmnpV4[i]})
    }
    binding.EntityData.Children.Append("dmnp-v6", types.YChild{"DmnpV6", nil})
    for i := range binding.DmnpV6 {
        binding.EntityData.Children.Append(types.GetSegmentPath(binding.DmnpV6[i]), types.YChild{"DmnpV6", binding.DmnpV6[i]})
    }
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("mag-name", types.YLeaf{"MagName", binding.MagName})
    binding.EntityData.Leafs.Append("nai-string", types.YLeaf{"NaiString", binding.NaiString})
    binding.EntityData.Leafs.Append("imsi-string", types.YLeaf{"ImsiString", binding.ImsiString})
    binding.EntityData.Leafs.Append("customer-name", types.YLeaf{"CustomerName", binding.CustomerName})
    binding.EntityData.Leafs.Append("mnnai", types.YLeaf{"Mnnai", binding.Mnnai})
    binding.EntityData.Leafs.Append("customer-name-xr", types.YLeaf{"CustomerNameXr", binding.CustomerNameXr})
    binding.EntityData.Leafs.Append("llid", types.YLeaf{"Llid", binding.Llid})
    binding.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", binding.PeerId})
    binding.EntityData.Leafs.Append("phyintf", types.YLeaf{"Phyintf", binding.Phyintf})
    binding.EntityData.Leafs.Append("tunnel", types.YLeaf{"Tunnel", binding.Tunnel})
    binding.EntityData.Leafs.Append("state", types.YLeaf{"State", binding.State})
    binding.EntityData.Leafs.Append("apn", types.YLeaf{"Apn", binding.Apn})
    binding.EntityData.Leafs.Append("att", types.YLeaf{"Att", binding.Att})
    binding.EntityData.Leafs.Append("hoa", types.YLeaf{"Hoa", binding.Hoa})
    binding.EntityData.Leafs.Append("dflt", types.YLeaf{"Dflt", binding.Dflt})
    binding.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", binding.Lifetime})
    binding.EntityData.Leafs.Append("liferem", types.YLeaf{"Liferem", binding.Liferem})
    binding.EntityData.Leafs.Append("refresh", types.YLeaf{"Refresh", binding.Refresh})
    binding.EntityData.Leafs.Append("refresh-rem", types.YLeaf{"RefreshRem", binding.RefreshRem})
    binding.EntityData.Leafs.Append("prefix-len", types.YLeaf{"PrefixLen", binding.PrefixLen})
    binding.EntityData.Leafs.Append("num-hnps", types.YLeaf{"NumHnps", binding.NumHnps})
    binding.EntityData.Leafs.Append("num-coa", types.YLeaf{"NumCoa", binding.NumCoa})
    binding.EntityData.Leafs.Append("num-dmnp-v4", types.YLeaf{"NumDmnpV4", binding.NumDmnpV4})
    binding.EntityData.Leafs.Append("num-dmnp-v6", types.YLeaf{"NumDmnpV6", binding.NumDmnpV6})
    binding.EntityData.Leafs.Append("hnps", types.YLeaf{"Hnps", binding.Hnps})
    binding.EntityData.Leafs.Append("ignore-home-address", types.YLeaf{"IgnoreHomeAddress", binding.IgnoreHomeAddress})
    binding.EntityData.Leafs.Append("up-stream-grekey", types.YLeaf{"UpStreamGrekey", binding.UpStreamGrekey})
    binding.EntityData.Leafs.Append("down-stream-grekey", types.YLeaf{"DownStreamGrekey", binding.DownStreamGrekey})
    binding.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", binding.Vrfid})

    binding.EntityData.YListKeys = []string {}

    return &(binding.EntityData)
}

// Pmipv6_Lma_Bindings_Binding_Coa
// COA entries
type Pmipv6_Lma_Bindings_Binding_Coa struct {
    EntityData types.CommonEntityData
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

func (coa *Pmipv6_Lma_Bindings_Binding_Coa) GetEntityData() *types.CommonEntityData {
    coa.EntityData.YFilter = coa.YFilter
    coa.EntityData.YangName = "coa"
    coa.EntityData.BundleName = "cisco_ios_xr"
    coa.EntityData.ParentYangName = "binding"
    coa.EntityData.SegmentPath = "coa"
    coa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coa.EntityData.Children = types.NewOrderedMap()
    coa.EntityData.Leafs = types.NewOrderedMap()
    coa.EntityData.Leafs.Append("llid", types.YLeaf{"Llid", coa.Llid})
    coa.EntityData.Leafs.Append("peer-name", types.YLeaf{"PeerName", coa.PeerName})
    coa.EntityData.Leafs.Append("tunnel", types.YLeaf{"Tunnel", coa.Tunnel})
    coa.EntityData.Leafs.Append("e-label", types.YLeaf{"ELabel", coa.ELabel})
    coa.EntityData.Leafs.Append("color", types.YLeaf{"Color", coa.Color})
    coa.EntityData.Leafs.Append("roa-min-tf", types.YLeaf{"RoaMinTf", coa.RoaMinTf})
    coa.EntityData.Leafs.Append("pstate", types.YLeaf{"Pstate", coa.Pstate})
    coa.EntityData.Leafs.Append("msisdn", types.YLeaf{"Msisdn", coa.Msisdn})
    coa.EntityData.Leafs.Append("imsi", types.YLeaf{"Imsi", coa.Imsi})
    coa.EntityData.Leafs.Append("cdma-nai", types.YLeaf{"CdmaNai", coa.CdmaNai})
    coa.EntityData.Leafs.Append("pgw-apn", types.YLeaf{"PgwApn", coa.PgwApn})
    coa.EntityData.Leafs.Append("pgw-trans-vrf", types.YLeaf{"PgwTransVrf", coa.PgwTransVrf})
    coa.EntityData.Leafs.Append("att", types.YLeaf{"Att", coa.Att})
    coa.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", coa.Lifetime})
    coa.EntityData.Leafs.Append("lifetime-remaining", types.YLeaf{"LifetimeRemaining", coa.LifetimeRemaining})
    coa.EntityData.Leafs.Append("refresh", types.YLeaf{"Refresh", coa.Refresh})
    coa.EntityData.Leafs.Append("refresh-rem", types.YLeaf{"RefreshRem", coa.RefreshRem})
    coa.EntityData.Leafs.Append("dnkey", types.YLeaf{"Dnkey", coa.Dnkey})
    coa.EntityData.Leafs.Append("upkey", types.YLeaf{"Upkey", coa.Upkey})
    coa.EntityData.Leafs.Append("coa-v4", types.YLeaf{"CoaV4", coa.CoaV4})
    coa.EntityData.Leafs.Append("coa-v6", types.YLeaf{"CoaV6", coa.CoaV6})

    coa.EntityData.YListKeys = []string {}

    return &(coa.EntityData)
}

// Pmipv6_Lma_Bindings_Binding_DmnpV4
// IPv4 DMNP prefixes
type Pmipv6_Lma_Bindings_Binding_DmnpV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 prefix length. The type is interface{} with range: 0..255.
    Pfxlen interface{}

    // IPv4 prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Prefix interface{}
}

func (dmnpV4 *Pmipv6_Lma_Bindings_Binding_DmnpV4) GetEntityData() *types.CommonEntityData {
    dmnpV4.EntityData.YFilter = dmnpV4.YFilter
    dmnpV4.EntityData.YangName = "dmnp-v4"
    dmnpV4.EntityData.BundleName = "cisco_ios_xr"
    dmnpV4.EntityData.ParentYangName = "binding"
    dmnpV4.EntityData.SegmentPath = "dmnp-v4"
    dmnpV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dmnpV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dmnpV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dmnpV4.EntityData.Children = types.NewOrderedMap()
    dmnpV4.EntityData.Leafs = types.NewOrderedMap()
    dmnpV4.EntityData.Leafs.Append("pfxlen", types.YLeaf{"Pfxlen", dmnpV4.Pfxlen})
    dmnpV4.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", dmnpV4.Prefix})

    dmnpV4.EntityData.YListKeys = []string {}

    return &(dmnpV4.EntityData)
}

// Pmipv6_Lma_Bindings_Binding_DmnpV6
// IPv6 DMNP prefixes
type Pmipv6_Lma_Bindings_Binding_DmnpV6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 prefix length. The type is interface{} with range: 0..255.
    Pfxlen interface{}

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}
}

func (dmnpV6 *Pmipv6_Lma_Bindings_Binding_DmnpV6) GetEntityData() *types.CommonEntityData {
    dmnpV6.EntityData.YFilter = dmnpV6.YFilter
    dmnpV6.EntityData.YangName = "dmnp-v6"
    dmnpV6.EntityData.BundleName = "cisco_ios_xr"
    dmnpV6.EntityData.ParentYangName = "binding"
    dmnpV6.EntityData.SegmentPath = "dmnp-v6"
    dmnpV6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dmnpV6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dmnpV6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dmnpV6.EntityData.Children = types.NewOrderedMap()
    dmnpV6.EntityData.Leafs = types.NewOrderedMap()
    dmnpV6.EntityData.Leafs.Append("pfxlen", types.YLeaf{"Pfxlen", dmnpV6.Pfxlen})
    dmnpV6.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", dmnpV6.Prefix})

    dmnpV6.EntityData.YListKeys = []string {}

    return &(dmnpV6.EntityData)
}

// Pmipv6_Lma_Heartbeats
// Table of Heartbeat
type Pmipv6_Lma_Heartbeats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Heartbeat information. The type is slice of
    // Pmipv6_Lma_Heartbeats_Heartbeat.
    Heartbeat []*Pmipv6_Lma_Heartbeats_Heartbeat
}

func (heartbeats *Pmipv6_Lma_Heartbeats) GetEntityData() *types.CommonEntityData {
    heartbeats.EntityData.YFilter = heartbeats.YFilter
    heartbeats.EntityData.YangName = "heartbeats"
    heartbeats.EntityData.BundleName = "cisco_ios_xr"
    heartbeats.EntityData.ParentYangName = "lma"
    heartbeats.EntityData.SegmentPath = "heartbeats"
    heartbeats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    heartbeats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    heartbeats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    heartbeats.EntityData.Children = types.NewOrderedMap()
    heartbeats.EntityData.Children.Append("heartbeat", types.YChild{"Heartbeat", nil})
    for i := range heartbeats.Heartbeat {
        heartbeats.EntityData.Children.Append(types.GetSegmentPath(heartbeats.Heartbeat[i]), types.YChild{"Heartbeat", heartbeats.Heartbeat[i]})
    }
    heartbeats.EntityData.Leafs = types.NewOrderedMap()

    heartbeats.EntityData.YListKeys = []string {}

    return &(heartbeats.EntityData)
}

// Pmipv6_Lma_Heartbeats_Heartbeat
// Heartbeat information
type Pmipv6_Lma_Heartbeats_Heartbeat struct {
    EntityData types.CommonEntityData
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

func (heartbeat *Pmipv6_Lma_Heartbeats_Heartbeat) GetEntityData() *types.CommonEntityData {
    heartbeat.EntityData.YFilter = heartbeat.YFilter
    heartbeat.EntityData.YangName = "heartbeat"
    heartbeat.EntityData.BundleName = "cisco_ios_xr"
    heartbeat.EntityData.ParentYangName = "heartbeats"
    heartbeat.EntityData.SegmentPath = "heartbeat" + types.AddKeyToken(heartbeat.PeerAddr, "peer-addr")
    heartbeat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    heartbeat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    heartbeat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    heartbeat.EntityData.Children = types.NewOrderedMap()
    heartbeat.EntityData.Leafs = types.NewOrderedMap()
    heartbeat.EntityData.Leafs.Append("peer-addr", types.YLeaf{"PeerAddr", heartbeat.PeerAddr})
    heartbeat.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", heartbeat.Vrf})
    heartbeat.EntityData.Leafs.Append("customer-name", types.YLeaf{"CustomerName", heartbeat.CustomerName})
    heartbeat.EntityData.Leafs.Append("source-port", types.YLeaf{"SourcePort", heartbeat.SourcePort})
    heartbeat.EntityData.Leafs.Append("destination-port", types.YLeaf{"DestinationPort", heartbeat.DestinationPort})
    heartbeat.EntityData.Leafs.Append("source-ipv4-address", types.YLeaf{"SourceIpv4Address", heartbeat.SourceIpv4Address})
    heartbeat.EntityData.Leafs.Append("destination-ipv4-address", types.YLeaf{"DestinationIpv4Address", heartbeat.DestinationIpv4Address})
    heartbeat.EntityData.Leafs.Append("source-ipv6-address", types.YLeaf{"SourceIpv6Address", heartbeat.SourceIpv6Address})
    heartbeat.EntityData.Leafs.Append("destination-ipv6-address", types.YLeaf{"DestinationIpv6Address", heartbeat.DestinationIpv6Address})
    heartbeat.EntityData.Leafs.Append("status", types.YLeaf{"Status", heartbeat.Status})
    heartbeat.EntityData.Leafs.Append("ipv6-path", types.YLeaf{"Ipv6Path", heartbeat.Ipv6Path})

    heartbeat.EntityData.YListKeys = []string {"PeerAddr"}

    return &(heartbeat.EntityData)
}

// Pmipv6_Lma_ConfigVariables
// Global Configuration Variables
type Pmipv6_Lma_ConfigVariables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of CustomerVariables.
    CustomerVariables Pmipv6_Lma_ConfigVariables_CustomerVariables

    // Global Configuration Variables.
    GlobalVariables Pmipv6_Lma_ConfigVariables_GlobalVariables
}

func (configVariables *Pmipv6_Lma_ConfigVariables) GetEntityData() *types.CommonEntityData {
    configVariables.EntityData.YFilter = configVariables.YFilter
    configVariables.EntityData.YangName = "config-variables"
    configVariables.EntityData.BundleName = "cisco_ios_xr"
    configVariables.EntityData.ParentYangName = "lma"
    configVariables.EntityData.SegmentPath = "config-variables"
    configVariables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configVariables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configVariables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configVariables.EntityData.Children = types.NewOrderedMap()
    configVariables.EntityData.Children.Append("customer-variables", types.YChild{"CustomerVariables", &configVariables.CustomerVariables})
    configVariables.EntityData.Children.Append("global-variables", types.YChild{"GlobalVariables", &configVariables.GlobalVariables})
    configVariables.EntityData.Leafs = types.NewOrderedMap()

    configVariables.EntityData.YListKeys = []string {}

    return &(configVariables.EntityData)
}

// Pmipv6_Lma_ConfigVariables_CustomerVariables
// Table of CustomerVariables
type Pmipv6_Lma_ConfigVariables_CustomerVariables struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Customer name string. The type is slice of
    // Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable.
    CustomerVariable []*Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable
}

func (customerVariables *Pmipv6_Lma_ConfigVariables_CustomerVariables) GetEntityData() *types.CommonEntityData {
    customerVariables.EntityData.YFilter = customerVariables.YFilter
    customerVariables.EntityData.YangName = "customer-variables"
    customerVariables.EntityData.BundleName = "cisco_ios_xr"
    customerVariables.EntityData.ParentYangName = "config-variables"
    customerVariables.EntityData.SegmentPath = "customer-variables"
    customerVariables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customerVariables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customerVariables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customerVariables.EntityData.Children = types.NewOrderedMap()
    customerVariables.EntityData.Children.Append("customer-variable", types.YChild{"CustomerVariable", nil})
    for i := range customerVariables.CustomerVariable {
        customerVariables.EntityData.Children.Append(types.GetSegmentPath(customerVariables.CustomerVariable[i]), types.YChild{"CustomerVariable", customerVariables.CustomerVariable[i]})
    }
    customerVariables.EntityData.Leafs = types.NewOrderedMap()

    customerVariables.EntityData.YListKeys = []string {}

    return &(customerVariables.EntityData)
}

// Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable
// Customer name string
type Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable struct {
    EntityData types.CommonEntityData
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

func (customerVariable *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable) GetEntityData() *types.CommonEntityData {
    customerVariable.EntityData.YFilter = customerVariable.YFilter
    customerVariable.EntityData.YangName = "customer-variable"
    customerVariable.EntityData.BundleName = "cisco_ios_xr"
    customerVariable.EntityData.ParentYangName = "customer-variables"
    customerVariable.EntityData.SegmentPath = "customer-variable" + types.AddKeyToken(customerVariable.CustomerName, "customer-name")
    customerVariable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    customerVariable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    customerVariable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    customerVariable.EntityData.Children = types.NewOrderedMap()
    customerVariable.EntityData.Children.Append("mll-service", types.YChild{"MllService", &customerVariable.MllService})
    customerVariable.EntityData.Leafs = types.NewOrderedMap()
    customerVariable.EntityData.Leafs.Append("customer-name", types.YLeaf{"CustomerName", customerVariable.CustomerName})
    customerVariable.EntityData.Leafs.Append("cust-name", types.YLeaf{"CustName", customerVariable.CustName})
    customerVariable.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", customerVariable.VrfName})
    customerVariable.EntityData.Leafs.Append("auth-option", types.YLeaf{"AuthOption", customerVariable.AuthOption})

    customerVariable.EntityData.YListKeys = []string {"CustomerName"}

    return &(customerVariable.EntityData)
}

// Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService
// MLL service parameters
type Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService struct {
    EntityData types.CommonEntityData
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

func (mllService *Pmipv6_Lma_ConfigVariables_CustomerVariables_CustomerVariable_MllService) GetEntityData() *types.CommonEntityData {
    mllService.EntityData.YFilter = mllService.YFilter
    mllService.EntityData.YangName = "mll-service"
    mllService.EntityData.BundleName = "cisco_ios_xr"
    mllService.EntityData.ParentYangName = "customer-variable"
    mllService.EntityData.SegmentPath = "mll-service"
    mllService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mllService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mllService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mllService.EntityData.Children = types.NewOrderedMap()
    mllService.EntityData.Leafs = types.NewOrderedMap()
    mllService.EntityData.Leafs.Append("ignore-hoa", types.YLeaf{"IgnoreHoa", mllService.IgnoreHoa})
    mllService.EntityData.Leafs.Append("mnp-ipv4-lmn-max", types.YLeaf{"MnpIpv4LmnMax", mllService.MnpIpv4LmnMax})
    mllService.EntityData.Leafs.Append("mnp-ipv6-lmn-max", types.YLeaf{"MnpIpv6LmnMax", mllService.MnpIpv6LmnMax})
    mllService.EntityData.Leafs.Append("mnp-lmn-max", types.YLeaf{"MnpLmnMax", mllService.MnpLmnMax})
    mllService.EntityData.Leafs.Append("mnp-ipv4-cust-max", types.YLeaf{"MnpIpv4CustMax", mllService.MnpIpv4CustMax})
    mllService.EntityData.Leafs.Append("mnp-ipv6-cust-max", types.YLeaf{"MnpIpv6CustMax", mllService.MnpIpv6CustMax})
    mllService.EntityData.Leafs.Append("mnp-cust-max", types.YLeaf{"MnpCustMax", mllService.MnpCustMax})
    mllService.EntityData.Leafs.Append("mnp-ipv4-cust-cur", types.YLeaf{"MnpIpv4CustCur", mllService.MnpIpv4CustCur})
    mllService.EntityData.Leafs.Append("mnp-ipv6-cust-cur", types.YLeaf{"MnpIpv6CustCur", mllService.MnpIpv6CustCur})

    mllService.EntityData.YListKeys = []string {}

    return &(mllService.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables
// Global Configuration Variables
type Pmipv6_Lma_ConfigVariables_GlobalVariables struct {
    EntityData types.CommonEntityData
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
    Intf []*Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf

    // Peer Parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer.
    Peer []*Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer

    // LMA Network Parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Network.
    Network []*Pmipv6_Lma_ConfigVariables_GlobalVariables_Network

    // Customer parameters. The type is slice of
    // Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust.
    Cust []*Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust
}

func (globalVariables *Pmipv6_Lma_ConfigVariables_GlobalVariables) GetEntityData() *types.CommonEntityData {
    globalVariables.EntityData.YFilter = globalVariables.YFilter
    globalVariables.EntityData.YangName = "global-variables"
    globalVariables.EntityData.BundleName = "cisco_ios_xr"
    globalVariables.EntityData.ParentYangName = "config-variables"
    globalVariables.EntityData.SegmentPath = "global-variables"
    globalVariables.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalVariables.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalVariables.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalVariables.EntityData.Children = types.NewOrderedMap()
    globalVariables.EntityData.Children.Append("parameters", types.YChild{"Parameters", &globalVariables.Parameters})
    globalVariables.EntityData.Children.Append("mll-service", types.YChild{"MllService", &globalVariables.MllService})
    globalVariables.EntityData.Children.Append("intf", types.YChild{"Intf", nil})
    for i := range globalVariables.Intf {
        globalVariables.EntityData.Children.Append(types.GetSegmentPath(globalVariables.Intf[i]), types.YChild{"Intf", globalVariables.Intf[i]})
    }
    globalVariables.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range globalVariables.Peer {
        globalVariables.EntityData.Children.Append(types.GetSegmentPath(globalVariables.Peer[i]), types.YChild{"Peer", globalVariables.Peer[i]})
    }
    globalVariables.EntityData.Children.Append("network", types.YChild{"Network", nil})
    for i := range globalVariables.Network {
        globalVariables.EntityData.Children.Append(types.GetSegmentPath(globalVariables.Network[i]), types.YChild{"Network", globalVariables.Network[i]})
    }
    globalVariables.EntityData.Children.Append("cust", types.YChild{"Cust", nil})
    for i := range globalVariables.Cust {
        globalVariables.EntityData.Children.Append(types.GetSegmentPath(globalVariables.Cust[i]), types.YChild{"Cust", globalVariables.Cust[i]})
    }
    globalVariables.EntityData.Leafs = types.NewOrderedMap()
    globalVariables.EntityData.Leafs.Append("domain", types.YLeaf{"Domain", globalVariables.Domain})
    globalVariables.EntityData.Leafs.Append("selfid", types.YLeaf{"Selfid", globalVariables.Selfid})
    globalVariables.EntityData.Leafs.Append("apn-name", types.YLeaf{"ApnName", globalVariables.ApnName})
    globalVariables.EntityData.Leafs.Append("role", types.YLeaf{"Role", globalVariables.Role})
    globalVariables.EntityData.Leafs.Append("count", types.YLeaf{"Count", globalVariables.Count})
    globalVariables.EntityData.Leafs.Append("peers", types.YLeaf{"Peers", globalVariables.Peers})
    globalVariables.EntityData.Leafs.Append("customers", types.YLeaf{"Customers", globalVariables.Customers})
    globalVariables.EntityData.Leafs.Append("num-network", types.YLeaf{"NumNetwork", globalVariables.NumNetwork})
    globalVariables.EntityData.Leafs.Append("discover-mn", types.YLeaf{"DiscoverMn", globalVariables.DiscoverMn})
    globalVariables.EntityData.Leafs.Append("local-routing", types.YLeaf{"LocalRouting", globalVariables.LocalRouting})
    globalVariables.EntityData.Leafs.Append("aaa-accounting", types.YLeaf{"AaaAccounting", globalVariables.AaaAccounting})
    globalVariables.EntityData.Leafs.Append("default-mn", types.YLeaf{"DefaultMn", globalVariables.DefaultMn})
    globalVariables.EntityData.Leafs.Append("apn", types.YLeaf{"Apn", globalVariables.Apn})
    globalVariables.EntityData.Leafs.Append("learn-mag", types.YLeaf{"LearnMag", globalVariables.LearnMag})
    globalVariables.EntityData.Leafs.Append("session-mgr", types.YLeaf{"SessionMgr", globalVariables.SessionMgr})
    globalVariables.EntityData.Leafs.Append("service", types.YLeaf{"Service", globalVariables.Service})
    globalVariables.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", globalVariables.Profile})
    globalVariables.EntityData.Leafs.Append("ddp", types.YLeaf{"Ddp", globalVariables.Ddp})
    globalVariables.EntityData.Leafs.Append("ddt", types.YLeaf{"Ddt", globalVariables.Ddt})
    globalVariables.EntityData.Leafs.Append("ddr", types.YLeaf{"Ddr", globalVariables.Ddr})

    globalVariables.EntityData.YListKeys = []string {}

    return &(globalVariables.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters
// Domain Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters struct {
    EntityData types.CommonEntityData
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

func (parameters *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "global-variables"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Children.Append("self-id", types.YChild{"SelfId", &parameters.SelfId})
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", parameters.Timestamp})
    parameters.EntityData.Leafs.Append("window", types.YLeaf{"Window", parameters.Window})
    parameters.EntityData.Leafs.Append("auth-option", types.YLeaf{"AuthOption", parameters.AuthOption})
    parameters.EntityData.Leafs.Append("reg-time", types.YLeaf{"RegTime", parameters.RegTime})
    parameters.EntityData.Leafs.Append("ref-time", types.YLeaf{"RefTime", parameters.RefTime})
    parameters.EntityData.Leafs.Append("retx", types.YLeaf{"Retx", parameters.Retx})
    parameters.EntityData.Leafs.Append("ret-max", types.YLeaf{"RetMax", parameters.RetMax})
    parameters.EntityData.Leafs.Append("bri-init", types.YLeaf{"BriInit", parameters.BriInit})
    parameters.EntityData.Leafs.Append("bri-retries", types.YLeaf{"BriRetries", parameters.BriRetries})
    parameters.EntityData.Leafs.Append("bri-max", types.YLeaf{"BriMax", parameters.BriMax})
    parameters.EntityData.Leafs.Append("max-bindings", types.YLeaf{"MaxBindings", parameters.MaxBindings})
    parameters.EntityData.Leafs.Append("hnp", types.YLeaf{"Hnp", parameters.Hnp})
    parameters.EntityData.Leafs.Append("encap", types.YLeaf{"Encap", parameters.Encap})
    parameters.EntityData.Leafs.Append("delete-time", types.YLeaf{"DeleteTime", parameters.DeleteTime})
    parameters.EntityData.Leafs.Append("create-time", types.YLeaf{"CreateTime", parameters.CreateTime})
    parameters.EntityData.Leafs.Append("up-grekey", types.YLeaf{"UpGrekey", parameters.UpGrekey})
    parameters.EntityData.Leafs.Append("down-grekey", types.YLeaf{"DownGrekey", parameters.DownGrekey})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId
// Self Identifier
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId struct {
    EntityData types.CommonEntityData
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

func (selfId *Pmipv6_Lma_ConfigVariables_GlobalVariables_Parameters_SelfId) GetEntityData() *types.CommonEntityData {
    selfId.EntityData.YFilter = selfId.YFilter
    selfId.EntityData.YangName = "self-id"
    selfId.EntityData.BundleName = "cisco_ios_xr"
    selfId.EntityData.ParentYangName = "parameters"
    selfId.EntityData.SegmentPath = "self-id"
    selfId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selfId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selfId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selfId.EntityData.Children = types.NewOrderedMap()
    selfId.EntityData.Leafs = types.NewOrderedMap()
    selfId.EntityData.Leafs.Append("entity", types.YLeaf{"Entity", selfId.Entity})
    selfId.EntityData.Leafs.Append("addr-type", types.YLeaf{"AddrType", selfId.AddrType})
    selfId.EntityData.Leafs.Append("address", types.YLeaf{"Address", selfId.Address})
    selfId.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", selfId.Ipv4Address})

    selfId.EntityData.YListKeys = []string {}

    return &(selfId.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService
// MLL service parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService struct {
    EntityData types.CommonEntityData
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

func (mllService *Pmipv6_Lma_ConfigVariables_GlobalVariables_MllService) GetEntityData() *types.CommonEntityData {
    mllService.EntityData.YFilter = mllService.YFilter
    mllService.EntityData.YangName = "mll-service"
    mllService.EntityData.BundleName = "cisco_ios_xr"
    mllService.EntityData.ParentYangName = "global-variables"
    mllService.EntityData.SegmentPath = "mll-service"
    mllService.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mllService.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mllService.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mllService.EntityData.Children = types.NewOrderedMap()
    mllService.EntityData.Leafs = types.NewOrderedMap()
    mllService.EntityData.Leafs.Append("ignore-hoa", types.YLeaf{"IgnoreHoa", mllService.IgnoreHoa})
    mllService.EntityData.Leafs.Append("mnp-ipv4-lmn-max", types.YLeaf{"MnpIpv4LmnMax", mllService.MnpIpv4LmnMax})
    mllService.EntityData.Leafs.Append("mnp-ipv6-lmn-max", types.YLeaf{"MnpIpv6LmnMax", mllService.MnpIpv6LmnMax})
    mllService.EntityData.Leafs.Append("mnp-lmn-max", types.YLeaf{"MnpLmnMax", mllService.MnpLmnMax})
    mllService.EntityData.Leafs.Append("mnp-ipv4-cust-max", types.YLeaf{"MnpIpv4CustMax", mllService.MnpIpv4CustMax})
    mllService.EntityData.Leafs.Append("mnp-ipv6-cust-max", types.YLeaf{"MnpIpv6CustMax", mllService.MnpIpv6CustMax})
    mllService.EntityData.Leafs.Append("mnp-cust-max", types.YLeaf{"MnpCustMax", mllService.MnpCustMax})
    mllService.EntityData.Leafs.Append("mnp-ipv4-cust-cur", types.YLeaf{"MnpIpv4CustCur", mllService.MnpIpv4CustCur})
    mllService.EntityData.Leafs.Append("mnp-ipv6-cust-cur", types.YLeaf{"MnpIpv6CustCur", mllService.MnpIpv6CustCur})

    mllService.EntityData.YListKeys = []string {}

    return &(mllService.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf
// MAG Access List
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // APN Present. The type is bool.
    Apn interface{}

    // Access Interface Name. The type is string.
    Interface interface{}

    // APN Name. The type is string.
    ApnName interface{}
}

func (intf *Pmipv6_Lma_ConfigVariables_GlobalVariables_Intf) GetEntityData() *types.CommonEntityData {
    intf.EntityData.YFilter = intf.YFilter
    intf.EntityData.YangName = "intf"
    intf.EntityData.BundleName = "cisco_ios_xr"
    intf.EntityData.ParentYangName = "global-variables"
    intf.EntityData.SegmentPath = "intf"
    intf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intf.EntityData.Children = types.NewOrderedMap()
    intf.EntityData.Leafs = types.NewOrderedMap()
    intf.EntityData.Leafs.Append("apn", types.YLeaf{"Apn", intf.Apn})
    intf.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", intf.Interface})
    intf.EntityData.Leafs.Append("apn-name", types.YLeaf{"ApnName", intf.ApnName})

    intf.EntityData.YListKeys = []string {}

    return &(intf.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer
// Peer Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer struct {
    EntityData types.CommonEntityData
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

func (peer *Pmipv6_Lma_ConfigVariables_GlobalVariables_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "global-variables"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer", types.YLeaf{"Peer", peer.Peer})
    peer.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", peer.VrfName})
    peer.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", peer.Interface})
    peer.EntityData.Leafs.Append("encap", types.YLeaf{"Encap", peer.Encap})
    peer.EntityData.Leafs.Append("auth", types.YLeaf{"Auth", peer.Auth})
    peer.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", peer.Vrf})
    peer.EntityData.Leafs.Append("statictunnel", types.YLeaf{"Statictunnel", peer.Statictunnel})

    peer.EntityData.YListKeys = []string {}

    return &(peer.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Network
// LMA Network Parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Network struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV4 pool Present. The type is bool.
    V4pool interface{}

    // IPV6 pool Present. The type is bool.
    V6pool interface{}

    // Network Name. The type is string.
    Network interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}

    // v4 prefix len. The type is interface{} with range: 0..255.
    V4pfxLen interface{}

    // v6 prefix len. The type is interface{} with range: 0..255.
    V6pfxLen interface{}

    // num of mrnet. The type is interface{} with range: 0..255.
    Mrnet interface{}
}

func (network *Pmipv6_Lma_ConfigVariables_GlobalVariables_Network) GetEntityData() *types.CommonEntityData {
    network.EntityData.YFilter = network.YFilter
    network.EntityData.YangName = "network"
    network.EntityData.BundleName = "cisco_ios_xr"
    network.EntityData.ParentYangName = "global-variables"
    network.EntityData.SegmentPath = "network"
    network.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    network.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    network.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    network.EntityData.Children = types.NewOrderedMap()
    network.EntityData.Leafs = types.NewOrderedMap()
    network.EntityData.Leafs.Append("v4pool", types.YLeaf{"V4pool", network.V4pool})
    network.EntityData.Leafs.Append("v6pool", types.YLeaf{"V6pool", network.V6pool})
    network.EntityData.Leafs.Append("network", types.YLeaf{"Network", network.Network})
    network.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", network.Ipv4})
    network.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", network.Ipv6})
    network.EntityData.Leafs.Append("v4pfx-len", types.YLeaf{"V4pfxLen", network.V4pfxLen})
    network.EntityData.Leafs.Append("v6pfx-len", types.YLeaf{"V6pfxLen", network.V6pfxLen})
    network.EntityData.Leafs.Append("mrnet", types.YLeaf{"Mrnet", network.Mrnet})

    network.EntityData.YListKeys = []string {}

    return &(network.EntityData)
}

// Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust
// Customer parameters
type Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust struct {
    EntityData types.CommonEntityData
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

func (cust *Pmipv6_Lma_ConfigVariables_GlobalVariables_Cust) GetEntityData() *types.CommonEntityData {
    cust.EntityData.YFilter = cust.YFilter
    cust.EntityData.YangName = "cust"
    cust.EntityData.BundleName = "cisco_ios_xr"
    cust.EntityData.ParentYangName = "global-variables"
    cust.EntityData.SegmentPath = "cust"
    cust.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cust.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cust.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cust.EntityData.Children = types.NewOrderedMap()
    cust.EntityData.Leafs = types.NewOrderedMap()
    cust.EntityData.Leafs.Append("cust", types.YLeaf{"Cust", cust.Cust})
    cust.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", cust.Vrf})
    cust.EntityData.Leafs.Append("t-vrf", types.YLeaf{"TVrf", cust.TVrf})
    cust.EntityData.Leafs.Append("auth-option", types.YLeaf{"AuthOption", cust.AuthOption})
    cust.EntityData.Leafs.Append("heart-beat", types.YLeaf{"HeartBeat", cust.HeartBeat})
    cust.EntityData.Leafs.Append("reg-time", types.YLeaf{"RegTime", cust.RegTime})
    cust.EntityData.Leafs.Append("cust-name", types.YLeaf{"CustName", cust.CustName})
    cust.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", cust.VrfName})
    cust.EntityData.Leafs.Append("t-vrf-name", types.YLeaf{"TVrfName", cust.TVrfName})

    cust.EntityData.YListKeys = []string {}

    return &(cust.EntityData)
}

