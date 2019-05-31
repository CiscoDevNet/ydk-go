// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sr: Segment Routing
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_ms_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_ms_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-cfg sr}", reflect.TypeOf(Sr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-cfg:sr", reflect.TypeOf(Sr{}))
}

// SrmsMiFlag represents Srms mi flag
type SrmsMiFlag string

const (
    // Disable flag
    SrmsMiFlag_disable SrmsMiFlag = "disable"

    // Enable flag
    SrmsMiFlag_enable SrmsMiFlag = "enable"
)

// SrmsAddressFamily represents Srms address family
type SrmsAddressFamily string

const (
    // IP version 4
    SrmsAddressFamily_ipv4 SrmsAddressFamily = "ipv4"

    // IP version 6
    SrmsAddressFamily_ipv6 SrmsAddressFamily = "ipv6"
)

// SidTypeList represents Sid type list
type SidTypeList string

const (
    // Absolute SID
    SidTypeList_absolute SidTypeList = "absolute"

    // Index SID
    SidTypeList_index SidTypeList = "index"
)

// Sr
// Segment Routing
type Sr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // enable SR. The type is interface{}.
    Enable interface{}

    // Segment Routing Local Block of Labels.
    LocalBlock Sr_LocalBlock

    // Mapping Server.
    Mappings Sr_Mappings

    // Segment Routing Adjacency SID.
    AdjacencySid Sr_AdjacencySid

    // Global Block Segment Routing.
    GlobalBlock Sr_GlobalBlock

    // Segment Routing with IPv6 dataplane.
    Srv6 Sr_Srv6
}

func (sr *Sr) GetEntityData() *types.CommonEntityData {
    sr.EntityData.YFilter = sr.YFilter
    sr.EntityData.YangName = "sr"
    sr.EntityData.BundleName = "cisco_ios_xr"
    sr.EntityData.ParentYangName = "Cisco-IOS-XR-segment-routing-ms-cfg"
    sr.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr"
    sr.EntityData.AbsolutePath = sr.EntityData.SegmentPath
    sr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sr.EntityData.Children = types.NewOrderedMap()
    sr.EntityData.Children.Append("local-block", types.YChild{"LocalBlock", &sr.LocalBlock})
    sr.EntityData.Children.Append("mappings", types.YChild{"Mappings", &sr.Mappings})
    sr.EntityData.Children.Append("adjacency-sid", types.YChild{"AdjacencySid", &sr.AdjacencySid})
    sr.EntityData.Children.Append("global-block", types.YChild{"GlobalBlock", &sr.GlobalBlock})
    sr.EntityData.Children.Append("Cisco-IOS-XR-segment-routing-srv6-cfg:srv6", types.YChild{"Srv6", &sr.Srv6})
    sr.EntityData.Leafs = types.NewOrderedMap()
    sr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", sr.Enable})

    sr.EntityData.YListKeys = []string {}

    return &(sr.EntityData)
}

// Sr_LocalBlock
// Segment Routing Local Block of Labels
// This type is a presence type.
type Sr_LocalBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // SRLB Lower Bound. The type is interface{} with range: 15000..1048574. This
    // attribute is mandatory.
    LowerBound interface{}

    // SRLB Upper Bound. The type is interface{} with range: 15001..1048575. This
    // attribute is mandatory.
    UpperBound interface{}
}

func (localBlock *Sr_LocalBlock) GetEntityData() *types.CommonEntityData {
    localBlock.EntityData.YFilter = localBlock.YFilter
    localBlock.EntityData.YangName = "local-block"
    localBlock.EntityData.BundleName = "cisco_ios_xr"
    localBlock.EntityData.ParentYangName = "sr"
    localBlock.EntityData.SegmentPath = "local-block"
    localBlock.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + localBlock.EntityData.SegmentPath
    localBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localBlock.EntityData.Children = types.NewOrderedMap()
    localBlock.EntityData.Leafs = types.NewOrderedMap()
    localBlock.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", localBlock.LowerBound})
    localBlock.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", localBlock.UpperBound})

    localBlock.EntityData.YListKeys = []string {}

    return &(localBlock.EntityData)
}

// Sr_Mappings
// Mapping Server
type Sr_Mappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP prefix to SID mapping. The type is slice of Sr_Mappings_Mapping.
    Mapping []*Sr_Mappings_Mapping
}

func (mappings *Sr_Mappings) GetEntityData() *types.CommonEntityData {
    mappings.EntityData.YFilter = mappings.YFilter
    mappings.EntityData.YangName = "mappings"
    mappings.EntityData.BundleName = "cisco_ios_xr"
    mappings.EntityData.ParentYangName = "sr"
    mappings.EntityData.SegmentPath = "mappings"
    mappings.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + mappings.EntityData.SegmentPath
    mappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappings.EntityData.Children = types.NewOrderedMap()
    mappings.EntityData.Children.Append("mapping", types.YChild{"Mapping", nil})
    for i := range mappings.Mapping {
        mappings.EntityData.Children.Append(types.GetSegmentPath(mappings.Mapping[i]), types.YChild{"Mapping", mappings.Mapping[i]})
    }
    mappings.EntityData.Leafs = types.NewOrderedMap()

    mappings.EntityData.YListKeys = []string {}

    return &(mappings.EntityData)
}

// Sr_Mappings_Mapping
// IP prefix to SID mapping
type Sr_Mappings_Mapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family. The type is SrmsAddressFamily.
    Af interface{}

    // This attribute is a key. IP prefix. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // This attribute is a key. Mask. The type is interface{} with range: 1..128.
    Mask interface{}

    // Start of SID index range. The type is interface{} with range: 0..1048575.
    SidStart interface{}

    // Range (number of SIDs). The type is interface{} with range: 0..1048575.
    SidRange interface{}

    // Enable/Disable Attached flag. The type is SrmsMiFlag.
    FlagAttached interface{}
}

func (mapping *Sr_Mappings_Mapping) GetEntityData() *types.CommonEntityData {
    mapping.EntityData.YFilter = mapping.YFilter
    mapping.EntityData.YangName = "mapping"
    mapping.EntityData.BundleName = "cisco_ios_xr"
    mapping.EntityData.ParentYangName = "mappings"
    mapping.EntityData.SegmentPath = "mapping" + types.AddKeyToken(mapping.Af, "af") + types.AddKeyToken(mapping.Ip, "ip") + types.AddKeyToken(mapping.Mask, "mask")
    mapping.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/mappings/" + mapping.EntityData.SegmentPath
    mapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mapping.EntityData.Children = types.NewOrderedMap()
    mapping.EntityData.Leafs = types.NewOrderedMap()
    mapping.EntityData.Leafs.Append("af", types.YLeaf{"Af", mapping.Af})
    mapping.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", mapping.Ip})
    mapping.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", mapping.Mask})
    mapping.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", mapping.SidStart})
    mapping.EntityData.Leafs.Append("sid-range", types.YLeaf{"SidRange", mapping.SidRange})
    mapping.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", mapping.FlagAttached})

    mapping.EntityData.YListKeys = []string {"Af", "Ip", "Mask"}

    return &(mapping.EntityData)
}

// Sr_AdjacencySid
// Segment Routing Adjacency SID
type Sr_AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface Table.
    Interfaces Sr_AdjacencySid_Interfaces
}

func (adjacencySid *Sr_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "sr"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + adjacencySid.EntityData.SegmentPath
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &adjacencySid.Interfaces})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()

    adjacencySid.EntityData.YListKeys = []string {}

    return &(adjacencySid.EntityData)
}

// Sr_AdjacencySid_Interfaces
// Segment Routing Adjacency SID Interface Table
type Sr_AdjacencySid_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface. The type is slice of
    // Sr_AdjacencySid_Interfaces_Interface.
    Interface []*Sr_AdjacencySid_Interfaces_Interface
}

func (interfaces *Sr_AdjacencySid_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "adjacency-sid"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface
// Segment Routing Adjacency SID Interface
type Sr_AdjacencySid_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    Interface interface{}

    // Segment Routing Adjacency SID Interface Address Family Table.
    AddressFamilies Sr_AdjacencySid_Interfaces_Interface_AddressFamilies
}

func (self *Sr_AdjacencySid_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("address-families", types.YChild{"AddressFamilies", &self.AddressFamilies})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})

    self.EntityData.YListKeys = []string {"Interface"}

    return &(self.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies
// Segment Routing Adjacency SID Interface
// Address Family Table
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface Address Family. The type is slice
    // of Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily.
    AddressFamily []*Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily
}

func (addressFamilies *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies) GetEntityData() *types.CommonEntityData {
    addressFamilies.EntityData.YFilter = addressFamilies.YFilter
    addressFamilies.EntityData.YangName = "address-families"
    addressFamilies.EntityData.BundleName = "cisco_ios_xr"
    addressFamilies.EntityData.ParentYangName = "interface"
    addressFamilies.EntityData.SegmentPath = "address-families"
    addressFamilies.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/interface/" + addressFamilies.EntityData.SegmentPath
    addressFamilies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilies.EntityData.Children = types.NewOrderedMap()
    addressFamilies.EntityData.Children.Append("address-family", types.YChild{"AddressFamily", nil})
    for i := range addressFamilies.AddressFamily {
        addressFamilies.EntityData.Children.Append(types.GetSegmentPath(addressFamilies.AddressFamily[i]), types.YChild{"AddressFamily", addressFamilies.AddressFamily[i]})
    }
    addressFamilies.EntityData.Leafs = types.NewOrderedMap()

    addressFamilies.EntityData.YListKeys = []string {}

    return &(addressFamilies.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily
// Segment Routing Adjacency SID Interface
// Address Family
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address Family. The type is SrmsAddressFamily.
    AddressFamily interface{}

    // Segment Routing Adjacency SID Interface Address Family NextHop Table.
    NextHops Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops
}

func (addressFamily *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xr"
    addressFamily.EntityData.ParentYangName = "address-families"
    addressFamily.EntityData.SegmentPath = "address-family" + types.AddKeyToken(addressFamily.AddressFamily, "address-family")
    addressFamily.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/interface/address-families/" + addressFamily.EntityData.SegmentPath
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = types.NewOrderedMap()
    addressFamily.EntityData.Children.Append("next-hops", types.YChild{"NextHops", &addressFamily.NextHops})
    addressFamily.EntityData.Leafs = types.NewOrderedMap()
    addressFamily.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", addressFamily.AddressFamily})

    addressFamily.EntityData.YListKeys = []string {"AddressFamily"}

    return &(addressFamily.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops
// Segment Routing Adjacency SID Interface
// Address Family NextHop Table
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface Address Family NextHop, use a
    // single ANYADDR (0.0.0.0 or ::) NextHop for point to point links. The type
    // is slice of
    // Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop.
    NextHop []*Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop
}

func (nextHops *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops) GetEntityData() *types.CommonEntityData {
    nextHops.EntityData.YFilter = nextHops.YFilter
    nextHops.EntityData.YangName = "next-hops"
    nextHops.EntityData.BundleName = "cisco_ios_xr"
    nextHops.EntityData.ParentYangName = "address-family"
    nextHops.EntityData.SegmentPath = "next-hops"
    nextHops.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/interface/address-families/address-family/" + nextHops.EntityData.SegmentPath
    nextHops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHops.EntityData.Children = types.NewOrderedMap()
    nextHops.EntityData.Children.Append("next-hop", types.YChild{"NextHop", nil})
    for i := range nextHops.NextHop {
        nextHops.EntityData.Children.Append(types.GetSegmentPath(nextHops.NextHop[i]), types.YChild{"NextHop", nextHops.NextHop[i]})
    }
    nextHops.EntityData.Leafs = types.NewOrderedMap()

    nextHops.EntityData.YListKeys = []string {}

    return &(nextHops.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop
// Segment Routing Adjacency SID Interface
// Address Family NextHop, use a single
// ANYADDR (0.0.0.0 or ::) NextHop for point
// to point links
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. NextHop IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // L2 Adjacency SID type and value.
    L2AdjacencySid Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop_L2AdjacencySid
}

func (nextHop *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xr"
    nextHop.EntityData.ParentYangName = "next-hops"
    nextHop.EntityData.SegmentPath = "next-hop" + types.AddKeyToken(nextHop.IpAddr, "ip-addr")
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/interface/address-families/address-family/next-hops/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Children.Append("l2-adjacency-sid", types.YChild{"L2AdjacencySid", &nextHop.L2AdjacencySid})
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", nextHop.IpAddr})

    nextHop.EntityData.YListKeys = []string {"IpAddr"}

    return &(nextHop.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop_L2AdjacencySid
// L2 Adjacency SID type and value
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop_L2AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is SidTypeList.
    SidType interface{}

    // SID value. The type is interface{} with range: 15000..1048575.
    AbsoluteSid interface{}

    // SID value. The type is interface{} with range: 0..1048575.
    IndexSid interface{}

    // SRLB block name. The type is string with pattern: b'(srlb_default)'. This
    // attribute is mandatory.
    Srlb interface{}
}

func (l2AdjacencySid *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop_L2AdjacencySid) GetEntityData() *types.CommonEntityData {
    l2AdjacencySid.EntityData.YFilter = l2AdjacencySid.YFilter
    l2AdjacencySid.EntityData.YangName = "l2-adjacency-sid"
    l2AdjacencySid.EntityData.BundleName = "cisco_ios_xr"
    l2AdjacencySid.EntityData.ParentYangName = "next-hop"
    l2AdjacencySid.EntityData.SegmentPath = "l2-adjacency-sid"
    l2AdjacencySid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/adjacency-sid/interfaces/interface/address-families/address-family/next-hops/next-hop/" + l2AdjacencySid.EntityData.SegmentPath
    l2AdjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2AdjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2AdjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2AdjacencySid.EntityData.Children = types.NewOrderedMap()
    l2AdjacencySid.EntityData.Leafs = types.NewOrderedMap()
    l2AdjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", l2AdjacencySid.SidType})
    l2AdjacencySid.EntityData.Leafs.Append("absolute-sid", types.YLeaf{"AbsoluteSid", l2AdjacencySid.AbsoluteSid})
    l2AdjacencySid.EntityData.Leafs.Append("index-sid", types.YLeaf{"IndexSid", l2AdjacencySid.IndexSid})
    l2AdjacencySid.EntityData.Leafs.Append("srlb", types.YLeaf{"Srlb", l2AdjacencySid.Srlb})

    l2AdjacencySid.EntityData.YListKeys = []string {}

    return &(l2AdjacencySid.EntityData)
}

// Sr_GlobalBlock
// Global Block Segment Routing
// This type is a presence type.
type Sr_GlobalBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // SRGB Lower Bound. The type is interface{} with range: 16000..1048574. This
    // attribute is mandatory.
    LowerBound interface{}

    // SRGB Upper Bound. The type is interface{} with range: 16001..1048575. This
    // attribute is mandatory.
    UpperBound interface{}
}

func (globalBlock *Sr_GlobalBlock) GetEntityData() *types.CommonEntityData {
    globalBlock.EntityData.YFilter = globalBlock.YFilter
    globalBlock.EntityData.YangName = "global-block"
    globalBlock.EntityData.BundleName = "cisco_ios_xr"
    globalBlock.EntityData.ParentYangName = "sr"
    globalBlock.EntityData.SegmentPath = "global-block"
    globalBlock.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + globalBlock.EntityData.SegmentPath
    globalBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalBlock.EntityData.Children = types.NewOrderedMap()
    globalBlock.EntityData.Leafs = types.NewOrderedMap()
    globalBlock.EntityData.Leafs.Append("lower-bound", types.YLeaf{"LowerBound", globalBlock.LowerBound})
    globalBlock.EntityData.Leafs.Append("upper-bound", types.YLeaf{"UpperBound", globalBlock.UpperBound})

    globalBlock.EntityData.YListKeys = []string {}

    return &(globalBlock.EntityData)
}

// Sr_Srv6
// Segment Routing with IPv6 dataplane
type Sr_Srv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SRv6. The type is interface{}.
    Enable interface{}

    // Configure SID holdtime for a stale/freed SID. The type is interface{} with
    // range: 0..60.
    SidHoldtime interface{}

    // Enable logging.
    Logging Sr_Srv6_Logging

    // Configure SRv6 locators parameters.
    Locators Sr_Srv6_Locators

    // Configure encapsulation related parameters.
    Encapsulation Sr_Srv6_Encapsulation
}

func (srv6 *Sr_Srv6) GetEntityData() *types.CommonEntityData {
    srv6.EntityData.YFilter = srv6.YFilter
    srv6.EntityData.YangName = "srv6"
    srv6.EntityData.BundleName = "cisco_ios_xr"
    srv6.EntityData.ParentYangName = "sr"
    srv6.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-srv6-cfg:srv6"
    srv6.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + srv6.EntityData.SegmentPath
    srv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srv6.EntityData.Children = types.NewOrderedMap()
    srv6.EntityData.Children.Append("logging", types.YChild{"Logging", &srv6.Logging})
    srv6.EntityData.Children.Append("locators", types.YChild{"Locators", &srv6.Locators})
    srv6.EntityData.Children.Append("encapsulation", types.YChild{"Encapsulation", &srv6.Encapsulation})
    srv6.EntityData.Leafs = types.NewOrderedMap()
    srv6.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", srv6.Enable})
    srv6.EntityData.Leafs.Append("sid-holdtime", types.YLeaf{"SidHoldtime", srv6.SidHoldtime})

    srv6.EntityData.YListKeys = []string {}

    return &(srv6.EntityData)
}

// Sr_Srv6_Logging
// Enable logging
type Sr_Srv6_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable logging for locator status changes. The type is interface{}.
    LocatorStatus interface{}
}

func (logging *Sr_Srv6_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "srv6"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("locator-status", types.YLeaf{"LocatorStatus", logging.LocatorStatus})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Sr_Srv6_Locators
// Configure SRv6 locators parameters
type Sr_Srv6_Locators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure SRv6 table of locators.
    Locators Sr_Srv6_Locators_Locators
}

func (locators *Sr_Srv6_Locators) GetEntityData() *types.CommonEntityData {
    locators.EntityData.YFilter = locators.YFilter
    locators.EntityData.YangName = "locators"
    locators.EntityData.BundleName = "cisco_ios_xr"
    locators.EntityData.ParentYangName = "srv6"
    locators.EntityData.SegmentPath = "locators"
    locators.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/" + locators.EntityData.SegmentPath
    locators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locators.EntityData.Children = types.NewOrderedMap()
    locators.EntityData.Children.Append("locators", types.YChild{"Locators", &locators.Locators})
    locators.EntityData.Leafs = types.NewOrderedMap()

    locators.EntityData.YListKeys = []string {}

    return &(locators.EntityData)
}

// Sr_Srv6_Locators_Locators
// Configure SRv6 table of locators
type Sr_Srv6_Locators_Locators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a SRv6 locator. The type is slice of
    // Sr_Srv6_Locators_Locators_Locator.
    Locator []*Sr_Srv6_Locators_Locators_Locator
}

func (locators *Sr_Srv6_Locators_Locators) GetEntityData() *types.CommonEntityData {
    locators.EntityData.YFilter = locators.YFilter
    locators.EntityData.YangName = "locators"
    locators.EntityData.BundleName = "cisco_ios_xr"
    locators.EntityData.ParentYangName = "locators"
    locators.EntityData.SegmentPath = "locators"
    locators.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/locators/" + locators.EntityData.SegmentPath
    locators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locators.EntityData.Children = types.NewOrderedMap()
    locators.EntityData.Children.Append("locator", types.YChild{"Locator", nil})
    for i := range locators.Locator {
        locators.EntityData.Children.Append(types.GetSegmentPath(locators.Locator[i]), types.YChild{"Locator", locators.Locator[i]})
    }
    locators.EntityData.Leafs = types.NewOrderedMap()

    locators.EntityData.YListKeys = []string {}

    return &(locators.EntityData)
}

// Sr_Srv6_Locators_Locators_Locator
// Configure a SRv6 locator
type Sr_Srv6_Locators_Locators_Locator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Locator name. The type is string with length:
    // 1..58.
    Name interface{}

    // Enable a SRv6 locator. The type is interface{}.
    LocatorEnable interface{}

    // Specify locator prefix value.
    Prefix Sr_Srv6_Locators_Locators_Locator_Prefix
}

func (locator *Sr_Srv6_Locators_Locators_Locator) GetEntityData() *types.CommonEntityData {
    locator.EntityData.YFilter = locator.YFilter
    locator.EntityData.YangName = "locator"
    locator.EntityData.BundleName = "cisco_ios_xr"
    locator.EntityData.ParentYangName = "locators"
    locator.EntityData.SegmentPath = "locator" + types.AddKeyToken(locator.Name, "name")
    locator.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/locators/locators/" + locator.EntityData.SegmentPath
    locator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    locator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    locator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    locator.EntityData.Children = types.NewOrderedMap()
    locator.EntityData.Children.Append("prefix", types.YChild{"Prefix", &locator.Prefix})
    locator.EntityData.Leafs = types.NewOrderedMap()
    locator.EntityData.Leafs.Append("name", types.YLeaf{"Name", locator.Name})
    locator.EntityData.Leafs.Append("locator-enable", types.YLeaf{"LocatorEnable", locator.LocatorEnable})

    locator.EntityData.YListKeys = []string {"Name"}

    return &(locator.EntityData)
}

// Sr_Srv6_Locators_Locators_Locator_Prefix
// Specify locator prefix value
type Sr_Srv6_Locators_Locators_Locator_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Prefix. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: 32..112.
    PrefixLength interface{}
}

func (prefix *Sr_Srv6_Locators_Locators_Locator_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "locator"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/locators/locators/locator/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", prefix.Prefix})
    prefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefix.PrefixLength})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// Sr_Srv6_Encapsulation
// Configure encapsulation related parameters
type Sr_Srv6_Encapsulation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a source address. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Configure IPv6 Hop-Limit options.
    HopLimit Sr_Srv6_Encapsulation_HopLimit
}

func (encapsulation *Sr_Srv6_Encapsulation) GetEntityData() *types.CommonEntityData {
    encapsulation.EntityData.YFilter = encapsulation.YFilter
    encapsulation.EntityData.YangName = "encapsulation"
    encapsulation.EntityData.BundleName = "cisco_ios_xr"
    encapsulation.EntityData.ParentYangName = "srv6"
    encapsulation.EntityData.SegmentPath = "encapsulation"
    encapsulation.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/" + encapsulation.EntityData.SegmentPath
    encapsulation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encapsulation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encapsulation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encapsulation.EntityData.Children = types.NewOrderedMap()
    encapsulation.EntityData.Children.Append("hop-limit", types.YChild{"HopLimit", &encapsulation.HopLimit})
    encapsulation.EntityData.Leafs = types.NewOrderedMap()
    encapsulation.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", encapsulation.SourceAddress})

    encapsulation.EntityData.YListKeys = []string {}

    return &(encapsulation.EntityData)
}

// Sr_Srv6_Encapsulation_HopLimit
// Configure IPv6 Hop-Limit options
type Sr_Srv6_Encapsulation_HopLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hop-Limit config option. The type is Srv6EncapsulationHopLimitOption.
    Option interface{}

    // Count for Hop-limit. The type is interface{} with range: 1..255.
    Value interface{}
}

func (hopLimit *Sr_Srv6_Encapsulation_HopLimit) GetEntityData() *types.CommonEntityData {
    hopLimit.EntityData.YFilter = hopLimit.YFilter
    hopLimit.EntityData.YangName = "hop-limit"
    hopLimit.EntityData.BundleName = "cisco_ios_xr"
    hopLimit.EntityData.ParentYangName = "encapsulation"
    hopLimit.EntityData.SegmentPath = "hop-limit"
    hopLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-segment-routing-srv6-cfg:srv6/encapsulation/" + hopLimit.EntityData.SegmentPath
    hopLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hopLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hopLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hopLimit.EntityData.Children = types.NewOrderedMap()
    hopLimit.EntityData.Leafs = types.NewOrderedMap()
    hopLimit.EntityData.Leafs.Append("option", types.YLeaf{"Option", hopLimit.Option})
    hopLimit.EntityData.Leafs.Append("value", types.YLeaf{"Value", hopLimit.Value})

    hopLimit.EntityData.YListKeys = []string {}

    return &(hopLimit.EntityData)
}

