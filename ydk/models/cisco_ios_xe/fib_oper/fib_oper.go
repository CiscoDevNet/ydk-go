// This module contains a collection of YANG definitions
// for IOS-XE FIB operational data.
package fib_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fib_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-fib-oper fib-oper-data}", reflect.TypeOf(FibOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-fib-oper:fib-oper-data", reflect.TypeOf(FibOperData{}))
}

// FibPathType represents Type of FIB path used
type FibPathType string

const (
    // Unknown FIB path type
    FibPathType_fib_path_type_unknown FibPathType = "fib-path-type-unknown"

    // Receive FIB path type
    FibPathType_fib_path_type_receive FibPathType = "fib-path-type-receive"

    // Connected FIB path type
    FibPathType_fib_path_type_connected FibPathType = "fib-path-type-connected"

    // Attached Prefix FIB path type
    FibPathType_fib_path_type_attached_prefix FibPathType = "fib-path-type-attached-prefix"

    // Attached Host FIB path type
    FibPathType_fib_path_type_attached_host FibPathType = "fib-path-type-attached-host"

    // Attached Nexthop FIB path type
    FibPathType_fib_path_type_attached_nexthop FibPathType = "fib-path-type-attached-nexthop"

    // Recursive FIB path type
    FibPathType_fib_path_type_recursive FibPathType = "fib-path-type-recursive"

    // Adjacency Prefix FIB path type
    FibPathType_fib_path_type_adjacency_prefix FibPathType = "fib-path-type-adjacency-prefix"

    // Special Prefix FIB path type
    FibPathType_fib_path_type_special_prefix FibPathType = "fib-path-type-special-prefix"
)

// FibAddressFamily represents FIB Address Family Types
type FibAddressFamily string

const (
    // Unknown Address Family
    FibAddressFamily_fib_addr_fam_unknown FibAddressFamily = "fib-addr-fam-unknown"

    // IPv4 Address Family
    FibAddressFamily_fib_addr_fam_ipv4 FibAddressFamily = "fib-addr-fam-ipv4"

    // IPv6 Address Family
    FibAddressFamily_fib_addr_fam_ipv6 FibAddressFamily = "fib-addr-fam-ipv6"
)

// EncapsulationHeaderType represents Types of header for packet encapsulation
type EncapsulationHeaderType string

const (
    // Unknown encapsulation header type
    EncapsulationHeaderType_encap_hdr_type_unknown EncapsulationHeaderType = "encap-hdr-type-unknown"

    // GRE encapsulation header type
    EncapsulationHeaderType_encap_hdr_type_gre EncapsulationHeaderType = "encap-hdr-type-gre"

    // IPv4 encapsulation header type
    EncapsulationHeaderType_encap_hdr_type_ipv4 EncapsulationHeaderType = "encap-hdr-type-ipv4"

    // IPv6 encapsulation header type
    EncapsulationHeaderType_encap_hdr_type_ipv6 EncapsulationHeaderType = "encap-hdr-type-ipv6"

    // MPLS encapsulation header type
    EncapsulationHeaderType_encap_hdr_type_mpls EncapsulationHeaderType = "encap-hdr-type-mpls"
)

// FibOperData
// This module contains a collection of YANG definitions for
// monitoring the operation of IOS-XE CEF.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
type FibOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIB Network Instances. The type is slice of FibOperData_FibNiEntry.
    FibNiEntry []*FibOperData_FibNiEntry
}

func (fibOperData *FibOperData) GetEntityData() *types.CommonEntityData {
    fibOperData.EntityData.YFilter = fibOperData.YFilter
    fibOperData.EntityData.YangName = "fib-oper-data"
    fibOperData.EntityData.BundleName = "cisco_ios_xe"
    fibOperData.EntityData.ParentYangName = "Cisco-IOS-XE-fib-oper"
    fibOperData.EntityData.SegmentPath = "Cisco-IOS-XE-fib-oper:fib-oper-data"
    fibOperData.EntityData.AbsolutePath = fibOperData.EntityData.SegmentPath
    fibOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fibOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fibOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fibOperData.EntityData.Children = types.NewOrderedMap()
    fibOperData.EntityData.Children.Append("fib-ni-entry", types.YChild{"FibNiEntry", nil})
    for i := range fibOperData.FibNiEntry {
        fibOperData.EntityData.Children.Append(types.GetSegmentPath(fibOperData.FibNiEntry[i]), types.YChild{"FibNiEntry", fibOperData.FibNiEntry[i]})
    }
    fibOperData.EntityData.Leafs = types.NewOrderedMap()

    fibOperData.EntityData.YListKeys = []string {}

    return &(fibOperData.EntityData)
}

// FibOperData_FibNiEntry
// FIB Network Instances
type FibOperData_FibNiEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Instance Name. The type is string.
    InstanceName interface{}

    // Address Family. The type is FibAddressFamily.
    Af interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    NumPfx interface{}

    // Number of forwarding prefixes. The type is interface{} with range:
    // 0..4294967295.
    NumPfxFwd interface{}

    // Number of non-forwarding prefixes. The type is interface{} with range:
    // 0..4294967295.
    NumPfxNonFwd interface{}

    // List of FIB entries. The type is slice of
    // FibOperData_FibNiEntry_FibEntries.
    FibEntries []*FibOperData_FibNiEntry_FibEntries
}

func (fibNiEntry *FibOperData_FibNiEntry) GetEntityData() *types.CommonEntityData {
    fibNiEntry.EntityData.YFilter = fibNiEntry.YFilter
    fibNiEntry.EntityData.YangName = "fib-ni-entry"
    fibNiEntry.EntityData.BundleName = "cisco_ios_xe"
    fibNiEntry.EntityData.ParentYangName = "fib-oper-data"
    fibNiEntry.EntityData.SegmentPath = "fib-ni-entry" + types.AddKeyToken(fibNiEntry.InstanceName, "instance-name")
    fibNiEntry.EntityData.AbsolutePath = "Cisco-IOS-XE-fib-oper:fib-oper-data/" + fibNiEntry.EntityData.SegmentPath
    fibNiEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fibNiEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fibNiEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fibNiEntry.EntityData.Children = types.NewOrderedMap()
    fibNiEntry.EntityData.Children.Append("fib-entries", types.YChild{"FibEntries", nil})
    for i := range fibNiEntry.FibEntries {
        fibNiEntry.EntityData.Children.Append(types.GetSegmentPath(fibNiEntry.FibEntries[i]), types.YChild{"FibEntries", fibNiEntry.FibEntries[i]})
    }
    fibNiEntry.EntityData.Leafs = types.NewOrderedMap()
    fibNiEntry.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", fibNiEntry.InstanceName})
    fibNiEntry.EntityData.Leafs.Append("af", types.YLeaf{"Af", fibNiEntry.Af})
    fibNiEntry.EntityData.Leafs.Append("num-pfx", types.YLeaf{"NumPfx", fibNiEntry.NumPfx})
    fibNiEntry.EntityData.Leafs.Append("num-pfx-fwd", types.YLeaf{"NumPfxFwd", fibNiEntry.NumPfxFwd})
    fibNiEntry.EntityData.Leafs.Append("num-pfx-non-fwd", types.YLeaf{"NumPfxNonFwd", fibNiEntry.NumPfxNonFwd})

    fibNiEntry.EntityData.YListKeys = []string {"InstanceName"}

    return &(fibNiEntry.EntityData)
}

// FibOperData_FibNiEntry_FibEntries
// List of FIB entries
type FibOperData_FibNiEntry_FibEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    IpAddr interface{}

    // Instance Name. The type is string.
    InstanceName interface{}

    // Address Family. The type is FibAddressFamily.
    Af interface{}

    // Number of Paths available. The type is interface{} with range: 0..255.
    NumPaths interface{}

    // Packets forwarded through this entry. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsForwarded interface{}

    // Octets forwarded through this entry. The type is interface{} with range:
    // 0..18446744073709551615.
    OctetsForwarded interface{}

    // List of FIB next-hop entries. The type is slice of
    // FibOperData_FibNiEntry_FibEntries_FibNexthopEntries.
    FibNexthopEntries []*FibOperData_FibNiEntry_FibEntries_FibNexthopEntries
}

func (fibEntries *FibOperData_FibNiEntry_FibEntries) GetEntityData() *types.CommonEntityData {
    fibEntries.EntityData.YFilter = fibEntries.YFilter
    fibEntries.EntityData.YangName = "fib-entries"
    fibEntries.EntityData.BundleName = "cisco_ios_xe"
    fibEntries.EntityData.ParentYangName = "fib-ni-entry"
    fibEntries.EntityData.SegmentPath = "fib-entries" + types.AddKeyToken(fibEntries.IpAddr, "ip-addr")
    fibEntries.EntityData.AbsolutePath = "Cisco-IOS-XE-fib-oper:fib-oper-data/fib-ni-entry/" + fibEntries.EntityData.SegmentPath
    fibEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fibEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fibEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fibEntries.EntityData.Children = types.NewOrderedMap()
    fibEntries.EntityData.Children.Append("fib-nexthop-entries", types.YChild{"FibNexthopEntries", nil})
    for i := range fibEntries.FibNexthopEntries {
        fibEntries.EntityData.Children.Append(types.GetSegmentPath(fibEntries.FibNexthopEntries[i]), types.YChild{"FibNexthopEntries", fibEntries.FibNexthopEntries[i]})
    }
    fibEntries.EntityData.Leafs = types.NewOrderedMap()
    fibEntries.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", fibEntries.IpAddr})
    fibEntries.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", fibEntries.InstanceName})
    fibEntries.EntityData.Leafs.Append("af", types.YLeaf{"Af", fibEntries.Af})
    fibEntries.EntityData.Leafs.Append("num-paths", types.YLeaf{"NumPaths", fibEntries.NumPaths})
    fibEntries.EntityData.Leafs.Append("packets-forwarded", types.YLeaf{"PacketsForwarded", fibEntries.PacketsForwarded})
    fibEntries.EntityData.Leafs.Append("octets-forwarded", types.YLeaf{"OctetsForwarded", fibEntries.OctetsForwarded})

    fibEntries.EntityData.YListKeys = []string {"IpAddr"}

    return &(fibEntries.EntityData)
}

// FibOperData_FibNiEntry_FibEntries_FibNexthopEntries
// List of FIB next-hop entries
type FibOperData_FibNiEntry_FibEntries_FibNexthopEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Nexthop IP Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    NhAddr interface{}

    // Unique Next-hop Path Index. The type is interface{} with range:
    // 0..4294967295.
    Index interface{}

    // Address Family. The type is FibAddressFamily.
    Af interface{}

    // Output Interface Name. The type is string.
    Ifname interface{}

    // FIB path type. The type is FibPathType.
    PathType interface{}

    // Unique Next-hop Path Index. The type is interface{} with range:
    // 0..4294967295.
    PathId interface{}

    // Next-hop weight. The type is interface{} with range: 0..255.
    Weight interface{}

    // Encap Header Type. The type is EncapsulationHeaderType.
    Encap interface{}

    // Decap Header Type. The type is EncapsulationHeaderType.
    Decap interface{}

    // Resolved Nexthop IP Address. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    ResolvedNhAddr interface{}
}

func (fibNexthopEntries *FibOperData_FibNiEntry_FibEntries_FibNexthopEntries) GetEntityData() *types.CommonEntityData {
    fibNexthopEntries.EntityData.YFilter = fibNexthopEntries.YFilter
    fibNexthopEntries.EntityData.YangName = "fib-nexthop-entries"
    fibNexthopEntries.EntityData.BundleName = "cisco_ios_xe"
    fibNexthopEntries.EntityData.ParentYangName = "fib-entries"
    fibNexthopEntries.EntityData.SegmentPath = "fib-nexthop-entries" + types.AddKeyToken(fibNexthopEntries.NhAddr, "nh-addr")
    fibNexthopEntries.EntityData.AbsolutePath = "Cisco-IOS-XE-fib-oper:fib-oper-data/fib-ni-entry/fib-entries/" + fibNexthopEntries.EntityData.SegmentPath
    fibNexthopEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fibNexthopEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fibNexthopEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fibNexthopEntries.EntityData.Children = types.NewOrderedMap()
    fibNexthopEntries.EntityData.Leafs = types.NewOrderedMap()
    fibNexthopEntries.EntityData.Leafs.Append("nh-addr", types.YLeaf{"NhAddr", fibNexthopEntries.NhAddr})
    fibNexthopEntries.EntityData.Leafs.Append("index", types.YLeaf{"Index", fibNexthopEntries.Index})
    fibNexthopEntries.EntityData.Leafs.Append("af", types.YLeaf{"Af", fibNexthopEntries.Af})
    fibNexthopEntries.EntityData.Leafs.Append("ifname", types.YLeaf{"Ifname", fibNexthopEntries.Ifname})
    fibNexthopEntries.EntityData.Leafs.Append("path-type", types.YLeaf{"PathType", fibNexthopEntries.PathType})
    fibNexthopEntries.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", fibNexthopEntries.PathId})
    fibNexthopEntries.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", fibNexthopEntries.Weight})
    fibNexthopEntries.EntityData.Leafs.Append("encap", types.YLeaf{"Encap", fibNexthopEntries.Encap})
    fibNexthopEntries.EntityData.Leafs.Append("decap", types.YLeaf{"Decap", fibNexthopEntries.Decap})
    fibNexthopEntries.EntityData.Leafs.Append("resolved-nh-addr", types.YLeaf{"ResolvedNhAddr", fibNexthopEntries.ResolvedNhAddr})

    fibNexthopEntries.EntityData.YListKeys = []string {"NhAddr"}

    return &(fibNexthopEntries.EntityData)
}

