// This module contains a collection of YANG definitions
// for NAT operational data.
// Copyright (c) 2017 by Cisco Systems, Inc.
// All rights reserved.
package nat_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package nat_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-nat-oper nat-data}", reflect.TypeOf(NatData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-nat-oper:nat-data", reflect.TypeOf(NatData{}))
}

// NatData
// NAT statistics
type NatData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global NAT statistics.
    IpNatStatistics NatData_IpNatStatistics

    // IP NAT translations. The type is slice of NatData_IpNatTranslation.
    IpNatTranslation []*NatData_IpNatTranslation
}

func (natData *NatData) GetEntityData() *types.CommonEntityData {
    natData.EntityData.YFilter = natData.YFilter
    natData.EntityData.YangName = "nat-data"
    natData.EntityData.BundleName = "cisco_ios_xe"
    natData.EntityData.ParentYangName = "Cisco-IOS-XE-nat-oper"
    natData.EntityData.SegmentPath = "Cisco-IOS-XE-nat-oper:nat-data"
    natData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    natData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    natData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    natData.EntityData.Children = types.NewOrderedMap()
    natData.EntityData.Children.Append("ip-nat-statistics", types.YChild{"IpNatStatistics", &natData.IpNatStatistics})
    natData.EntityData.Children.Append("ip-nat-translation", types.YChild{"IpNatTranslation", nil})
    for i := range natData.IpNatTranslation {
        natData.EntityData.Children.Append(types.GetSegmentPath(natData.IpNatTranslation[i]), types.YChild{"IpNatTranslation", natData.IpNatTranslation[i]})
    }
    natData.EntityData.Leafs = types.NewOrderedMap()

    natData.EntityData.YListKeys = []string {}

    return &(natData.EntityData)
}

// NatData_IpNatStatistics
// Global NAT statistics
// This type is a presence type.
type NatData_IpNatStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Indicates if the NAT feature has been initialized. The type is bool.
    Initialized interface{}

    // Total translations. The type is interface{} with range:
    // 0..18446744073709551615.
    Entries interface{}

    // Total static translations. The type is interface{} with range:
    // 0..18446744073709551615.
    Statics interface{}

    // Sorted static translations by domain. The type is slice of interface{} with
    // range: 0..18446744073709551615.
    StaticsSorted []interface{}

    // Total flows. The type is interface{} with range: 0..18446744073709551615.
    Flows interface{}

    // Number of inside interfaces. The type is interface{} with range:
    // 0..18446744073709551615.
    Insides interface{}

    // Number of outside interfaces. The type is interface{} with range:
    // 0..18446744073709551615.
    Outsides interface{}

    // Number of entries which timed out . The type is interface{} with range:
    // 0..18446744073709551615.
    EntryTimeouts interface{}

    // Successful searches with matching NAT session. The type is interface{} with
    // range: 0..18446744073709551615.
    Hits interface{}

    // Unsuccessful searches without matching NAT session. The type is interface{}
    // with range: 0..18446744073709551615.
    Misses interface{}

    // Translated in interrupt switching. The type is interface{} with range:
    // 0..18446744073709551615.
    InterruptSwitched interface{}

    // Packets punted to process. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsPunted interface{}

    // Counter for saved fragment packets. The type is interface{} with range:
    // 0..18446744073709551615.
    FragPakCount interface{}

    // Dropped pool stats from platform. The type is interface{} with range:
    // 0..18446744073709551615.
    PoolStatsDrop interface{}

    // Dropped mapping stats from platform. The type is interface{} with range:
    // 0..18446744073709551615.
    MappingStatsDrop interface{}

    // Counter for port block alloc req fails. The type is interface{} with range:
    // 0..18446744073709551615.
    PortlistReqFail interface{}

    // Counter for add ipalias fails drops. The type is interface{} with range:
    // 0..18446744073709551615.
    IpaliasAddFail interface{}

    // Counter for add limit_entry fails drops. The type is interface{} with
    // range: 0..18446744073709551615.
    LimitEntryAddFail interface{}

    // Counter for NAT inside->outside drops. The type is interface{} with range:
    // 0..18446744073709551615.
    In2outDrops interface{}

    // Counter for NAT outside->inside drops. The type is interface{} with range:
    // 0..18446744073709551615.
    Out2inDrops interface{}

    // MIB counter for address binds. The type is interface{} with range:
    // 0..4294967295.
    MibAddrBinds interface{}

    // MIB counter for address port binds. The type is interface{} with range:
    // 0..4294967295.
    MibAddportBinds interface{}
}

func (ipNatStatistics *NatData_IpNatStatistics) GetEntityData() *types.CommonEntityData {
    ipNatStatistics.EntityData.YFilter = ipNatStatistics.YFilter
    ipNatStatistics.EntityData.YangName = "ip-nat-statistics"
    ipNatStatistics.EntityData.BundleName = "cisco_ios_xe"
    ipNatStatistics.EntityData.ParentYangName = "nat-data"
    ipNatStatistics.EntityData.SegmentPath = "ip-nat-statistics"
    ipNatStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNatStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNatStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNatStatistics.EntityData.Children = types.NewOrderedMap()
    ipNatStatistics.EntityData.Leafs = types.NewOrderedMap()
    ipNatStatistics.EntityData.Leafs.Append("initialized", types.YLeaf{"Initialized", ipNatStatistics.Initialized})
    ipNatStatistics.EntityData.Leafs.Append("entries", types.YLeaf{"Entries", ipNatStatistics.Entries})
    ipNatStatistics.EntityData.Leafs.Append("statics", types.YLeaf{"Statics", ipNatStatistics.Statics})
    ipNatStatistics.EntityData.Leafs.Append("statics-sorted", types.YLeaf{"StaticsSorted", ipNatStatistics.StaticsSorted})
    ipNatStatistics.EntityData.Leafs.Append("flows", types.YLeaf{"Flows", ipNatStatistics.Flows})
    ipNatStatistics.EntityData.Leafs.Append("insides", types.YLeaf{"Insides", ipNatStatistics.Insides})
    ipNatStatistics.EntityData.Leafs.Append("outsides", types.YLeaf{"Outsides", ipNatStatistics.Outsides})
    ipNatStatistics.EntityData.Leafs.Append("entry-timeouts", types.YLeaf{"EntryTimeouts", ipNatStatistics.EntryTimeouts})
    ipNatStatistics.EntityData.Leafs.Append("hits", types.YLeaf{"Hits", ipNatStatistics.Hits})
    ipNatStatistics.EntityData.Leafs.Append("misses", types.YLeaf{"Misses", ipNatStatistics.Misses})
    ipNatStatistics.EntityData.Leafs.Append("interrupt-switched", types.YLeaf{"InterruptSwitched", ipNatStatistics.InterruptSwitched})
    ipNatStatistics.EntityData.Leafs.Append("packets-punted", types.YLeaf{"PacketsPunted", ipNatStatistics.PacketsPunted})
    ipNatStatistics.EntityData.Leafs.Append("frag-pak-count", types.YLeaf{"FragPakCount", ipNatStatistics.FragPakCount})
    ipNatStatistics.EntityData.Leafs.Append("pool-stats-drop", types.YLeaf{"PoolStatsDrop", ipNatStatistics.PoolStatsDrop})
    ipNatStatistics.EntityData.Leafs.Append("mapping-stats-drop", types.YLeaf{"MappingStatsDrop", ipNatStatistics.MappingStatsDrop})
    ipNatStatistics.EntityData.Leafs.Append("portlist-req-fail", types.YLeaf{"PortlistReqFail", ipNatStatistics.PortlistReqFail})
    ipNatStatistics.EntityData.Leafs.Append("ipalias-add-fail", types.YLeaf{"IpaliasAddFail", ipNatStatistics.IpaliasAddFail})
    ipNatStatistics.EntityData.Leafs.Append("limit-entry-add-fail", types.YLeaf{"LimitEntryAddFail", ipNatStatistics.LimitEntryAddFail})
    ipNatStatistics.EntityData.Leafs.Append("in2out-drops", types.YLeaf{"In2outDrops", ipNatStatistics.In2outDrops})
    ipNatStatistics.EntityData.Leafs.Append("out2in-drops", types.YLeaf{"Out2inDrops", ipNatStatistics.Out2inDrops})
    ipNatStatistics.EntityData.Leafs.Append("mib-addr-binds", types.YLeaf{"MibAddrBinds", ipNatStatistics.MibAddrBinds})
    ipNatStatistics.EntityData.Leafs.Append("mib-addport-binds", types.YLeaf{"MibAddportBinds", ipNatStatistics.MibAddportBinds})

    ipNatStatistics.EntityData.YListKeys = []string {}

    return &(ipNatStatistics.EntityData)
}

// NatData_IpNatTranslation
// IP NAT translations
type NatData_IpNatTranslation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Inside local address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    InsideLocalAddr interface{}

    // This attribute is a key. Outside local address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OutsideLocalAddr interface{}

    // This attribute is a key. Inside local port. The type is interface{} with
    // range: 0..65535.
    InsideLocalPort interface{}

    // This attribute is a key. Outside local port. The type is interface{} with
    // range: 0..65535.
    OutsideLocalPort interface{}

    // This attribute is a key. VRF ID. The type is interface{} with range:
    // 0..65535.
    Vrfid interface{}

    // This attribute is a key. Protocol. The type is interface{} with range:
    // 0..255.
    Protocol interface{}

    // Inside global address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    InsideGlobalAddr interface{}

    // Outside global address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    OutsideGlobalAddr interface{}

    // Inside global port. The type is interface{} with range: 0..65535.
    InsideGlobalPort interface{}

    // Outside global port. The type is interface{} with range: 0..65535.
    OutsideGlobalPort interface{}

    // Translation flags. The type is interface{} with range: 0..4294967295.
    Flags interface{}

    // Application type. The type is interface{} with range: 0..255.
    ApplicationType interface{}
}

func (ipNatTranslation *NatData_IpNatTranslation) GetEntityData() *types.CommonEntityData {
    ipNatTranslation.EntityData.YFilter = ipNatTranslation.YFilter
    ipNatTranslation.EntityData.YangName = "ip-nat-translation"
    ipNatTranslation.EntityData.BundleName = "cisco_ios_xe"
    ipNatTranslation.EntityData.ParentYangName = "nat-data"
    ipNatTranslation.EntityData.SegmentPath = "ip-nat-translation" + types.AddKeyToken(ipNatTranslation.InsideLocalAddr, "inside-local-addr") + types.AddKeyToken(ipNatTranslation.OutsideLocalAddr, "outside-local-addr") + types.AddKeyToken(ipNatTranslation.InsideLocalPort, "inside-local-port") + types.AddKeyToken(ipNatTranslation.OutsideLocalPort, "outside-local-port") + types.AddKeyToken(ipNatTranslation.Vrfid, "vrfid") + types.AddKeyToken(ipNatTranslation.Protocol, "protocol")
    ipNatTranslation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipNatTranslation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipNatTranslation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipNatTranslation.EntityData.Children = types.NewOrderedMap()
    ipNatTranslation.EntityData.Leafs = types.NewOrderedMap()
    ipNatTranslation.EntityData.Leafs.Append("inside-local-addr", types.YLeaf{"InsideLocalAddr", ipNatTranslation.InsideLocalAddr})
    ipNatTranslation.EntityData.Leafs.Append("outside-local-addr", types.YLeaf{"OutsideLocalAddr", ipNatTranslation.OutsideLocalAddr})
    ipNatTranslation.EntityData.Leafs.Append("inside-local-port", types.YLeaf{"InsideLocalPort", ipNatTranslation.InsideLocalPort})
    ipNatTranslation.EntityData.Leafs.Append("outside-local-port", types.YLeaf{"OutsideLocalPort", ipNatTranslation.OutsideLocalPort})
    ipNatTranslation.EntityData.Leafs.Append("vrfid", types.YLeaf{"Vrfid", ipNatTranslation.Vrfid})
    ipNatTranslation.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", ipNatTranslation.Protocol})
    ipNatTranslation.EntityData.Leafs.Append("inside-global-addr", types.YLeaf{"InsideGlobalAddr", ipNatTranslation.InsideGlobalAddr})
    ipNatTranslation.EntityData.Leafs.Append("outside-global-addr", types.YLeaf{"OutsideGlobalAddr", ipNatTranslation.OutsideGlobalAddr})
    ipNatTranslation.EntityData.Leafs.Append("inside-global-port", types.YLeaf{"InsideGlobalPort", ipNatTranslation.InsideGlobalPort})
    ipNatTranslation.EntityData.Leafs.Append("outside-global-port", types.YLeaf{"OutsideGlobalPort", ipNatTranslation.OutsideGlobalPort})
    ipNatTranslation.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", ipNatTranslation.Flags})
    ipNatTranslation.EntityData.Leafs.Append("application-type", types.YLeaf{"ApplicationType", ipNatTranslation.ApplicationType})

    ipNatTranslation.EntityData.YListKeys = []string {"InsideLocalAddr", "OutsideLocalAddr", "InsideLocalPort", "OutsideLocalPort", "Vrfid", "Protocol"}

    return &(ipNatTranslation.EntityData)
}

