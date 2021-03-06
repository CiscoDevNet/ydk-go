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

    // Traffic Engineering configuration data.
    TrafficEngineering Sr_TrafficEngineering
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
    sr.EntityData.Children.Append("Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering", types.YChild{"TrafficEngineering", &sr.TrafficEngineering})
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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
    // [a-zA-Z0-9._/-]+.
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    // SRLB block name. The type is string with pattern: (srlb_default). This
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

// Sr_TrafficEngineering
// Traffic Engineering configuration data
type Sr_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use TE-latency algorithm. The type is interface{}.
    TeLatency interface{}

    // Maximum SID Depth Configuration. The type is interface{} with range:
    // 1..255.
    MaximumSidDepth interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // SR-TE out-of-resources handling configuration.
    OutOfResources Sr_TrafficEngineering_OutOfResources

    // On-demand color configuration.
    OnDemandColors Sr_TrafficEngineering_OnDemandColors

    // Segment-lists configuration.
    Segments Sr_TrafficEngineering_Segments

    // Logging configuration.
    Logging Sr_TrafficEngineering_Logging

    // SR-TE timers configuration.
    Timers Sr_TrafficEngineering_Timers

    // Binding sid rules.
    BindingSidRules Sr_TrafficEngineering_BindingSidRules

    // Policy configuration.
    Policies Sr_TrafficEngineering_Policies

    // SR-TE interfaces.
    SrteInterfaces Sr_TrafficEngineering_SrteInterfaces

    // Path Computation Client.
    Pcc Sr_TrafficEngineering_Pcc

    // Affinity-map configuration.
    AffinityMaps Sr_TrafficEngineering_AffinityMaps
}

func (trafficEngineering *Sr_TrafficEngineering) GetEntityData() *types.CommonEntityData {
    trafficEngineering.EntityData.YFilter = trafficEngineering.YFilter
    trafficEngineering.EntityData.YangName = "traffic-engineering"
    trafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    trafficEngineering.EntityData.ParentYangName = "sr"
    trafficEngineering.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering"
    trafficEngineering.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/" + trafficEngineering.EntityData.SegmentPath
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = types.NewOrderedMap()
    trafficEngineering.EntityData.Children.Append("out-of-resources", types.YChild{"OutOfResources", &trafficEngineering.OutOfResources})
    trafficEngineering.EntityData.Children.Append("on-demand-colors", types.YChild{"OnDemandColors", &trafficEngineering.OnDemandColors})
    trafficEngineering.EntityData.Children.Append("segments", types.YChild{"Segments", &trafficEngineering.Segments})
    trafficEngineering.EntityData.Children.Append("logging", types.YChild{"Logging", &trafficEngineering.Logging})
    trafficEngineering.EntityData.Children.Append("timers", types.YChild{"Timers", &trafficEngineering.Timers})
    trafficEngineering.EntityData.Children.Append("binding-sid-rules", types.YChild{"BindingSidRules", &trafficEngineering.BindingSidRules})
    trafficEngineering.EntityData.Children.Append("policies", types.YChild{"Policies", &trafficEngineering.Policies})
    trafficEngineering.EntityData.Children.Append("srte-interfaces", types.YChild{"SrteInterfaces", &trafficEngineering.SrteInterfaces})
    trafficEngineering.EntityData.Children.Append("pcc", types.YChild{"Pcc", &trafficEngineering.Pcc})
    trafficEngineering.EntityData.Children.Append("affinity-maps", types.YChild{"AffinityMaps", &trafficEngineering.AffinityMaps})
    trafficEngineering.EntityData.Leafs = types.NewOrderedMap()
    trafficEngineering.EntityData.Leafs.Append("te-latency", types.YLeaf{"TeLatency", trafficEngineering.TeLatency})
    trafficEngineering.EntityData.Leafs.Append("maximum-sid-depth", types.YLeaf{"MaximumSidDepth", trafficEngineering.MaximumSidDepth})
    trafficEngineering.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", trafficEngineering.Enable})

    trafficEngineering.EntityData.YListKeys = []string {}

    return &(trafficEngineering.EntityData)
}

// Sr_TrafficEngineering_OutOfResources
// SR-TE out-of-resources handling configuration
type Sr_TrafficEngineering_OutOfResources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum total number of LSP path operations in a single batch. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    MaximumPathsBatch interface{}

    // Maximum total number of LSP paths that can be created. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    MaximumPaths interface{}
}

func (outOfResources *Sr_TrafficEngineering_OutOfResources) GetEntityData() *types.CommonEntityData {
    outOfResources.EntityData.YFilter = outOfResources.YFilter
    outOfResources.EntityData.YangName = "out-of-resources"
    outOfResources.EntityData.BundleName = "cisco_ios_xr"
    outOfResources.EntityData.ParentYangName = "traffic-engineering"
    outOfResources.EntityData.SegmentPath = "out-of-resources"
    outOfResources.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + outOfResources.EntityData.SegmentPath
    outOfResources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outOfResources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outOfResources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outOfResources.EntityData.Children = types.NewOrderedMap()
    outOfResources.EntityData.Leafs = types.NewOrderedMap()
    outOfResources.EntityData.Leafs.Append("maximum-paths-batch", types.YLeaf{"MaximumPathsBatch", outOfResources.MaximumPathsBatch})
    outOfResources.EntityData.Leafs.Append("maximum-paths", types.YLeaf{"MaximumPaths", outOfResources.MaximumPaths})

    outOfResources.EntityData.YListKeys = []string {}

    return &(outOfResources.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors
// On-demand color configuration
type Sr_TrafficEngineering_OnDemandColors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // On-demand color configuration. The type is slice of
    // Sr_TrafficEngineering_OnDemandColors_OnDemandColor.
    OnDemandColor []*Sr_TrafficEngineering_OnDemandColors_OnDemandColor
}

func (onDemandColors *Sr_TrafficEngineering_OnDemandColors) GetEntityData() *types.CommonEntityData {
    onDemandColors.EntityData.YFilter = onDemandColors.YFilter
    onDemandColors.EntityData.YangName = "on-demand-colors"
    onDemandColors.EntityData.BundleName = "cisco_ios_xr"
    onDemandColors.EntityData.ParentYangName = "traffic-engineering"
    onDemandColors.EntityData.SegmentPath = "on-demand-colors"
    onDemandColors.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + onDemandColors.EntityData.SegmentPath
    onDemandColors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColors.EntityData.Children = types.NewOrderedMap()
    onDemandColors.EntityData.Children.Append("on-demand-color", types.YChild{"OnDemandColor", nil})
    for i := range onDemandColors.OnDemandColor {
        onDemandColors.EntityData.Children.Append(types.GetSegmentPath(onDemandColors.OnDemandColor[i]), types.YChild{"OnDemandColor", onDemandColors.OnDemandColor[i]})
    }
    onDemandColors.EntityData.Leafs = types.NewOrderedMap()

    onDemandColors.EntityData.YListKeys = []string {}

    return &(onDemandColors.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor
// On-demand color configuration
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Color. The type is interface{} with range:
    // 1..4294967295.
    Color interface{}

    // The value of the bandwidth reserved by this policy in kbps. The type is
    // interface{} with range: 1..4294967295.
    Bandwidth interface{}

    // Maximum SID Depth Configuration. The type is interface{} with range:
    // 1..255.
    MaximumSidDepth interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Dynamic MPLS path properties.
    OnDemandColorDynMpls Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls
}

func (onDemandColor *Sr_TrafficEngineering_OnDemandColors_OnDemandColor) GetEntityData() *types.CommonEntityData {
    onDemandColor.EntityData.YFilter = onDemandColor.YFilter
    onDemandColor.EntityData.YangName = "on-demand-color"
    onDemandColor.EntityData.BundleName = "cisco_ios_xr"
    onDemandColor.EntityData.ParentYangName = "on-demand-colors"
    onDemandColor.EntityData.SegmentPath = "on-demand-color" + types.AddKeyToken(onDemandColor.Color, "color")
    onDemandColor.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/" + onDemandColor.EntityData.SegmentPath
    onDemandColor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColor.EntityData.Children = types.NewOrderedMap()
    onDemandColor.EntityData.Children.Append("on-demand-color-dyn-mpls", types.YChild{"OnDemandColorDynMpls", &onDemandColor.OnDemandColorDynMpls})
    onDemandColor.EntityData.Leafs = types.NewOrderedMap()
    onDemandColor.EntityData.Leafs.Append("color", types.YLeaf{"Color", onDemandColor.Color})
    onDemandColor.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", onDemandColor.Bandwidth})
    onDemandColor.EntityData.Leafs.Append("maximum-sid-depth", types.YLeaf{"MaximumSidDepth", onDemandColor.MaximumSidDepth})
    onDemandColor.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", onDemandColor.Enable})

    onDemandColor.EntityData.YListKeys = []string {"Color"}

    return &(onDemandColor.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls
// Dynamic MPLS path properties
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix-SID algorithm. The type is interface{} with range: 128..255.
    OnDemandColorDynMplsFlexAlgorithm interface{}

    // Dynamic MPLS path properties submode Enable. The type is interface{}.
    Enable interface{}

    // Metric type.
    OnDemandColorDynMplsMetric Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric

    // Use Path Computation Element.
    OnDemandColorDynMplsPce Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsPce

    // Disjoint path.
    DisjointPath Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath

    // SR path computation and verification affinity rules.
    AffinityRules Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules
}

func (onDemandColorDynMpls *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls) GetEntityData() *types.CommonEntityData {
    onDemandColorDynMpls.EntityData.YFilter = onDemandColorDynMpls.YFilter
    onDemandColorDynMpls.EntityData.YangName = "on-demand-color-dyn-mpls"
    onDemandColorDynMpls.EntityData.BundleName = "cisco_ios_xr"
    onDemandColorDynMpls.EntityData.ParentYangName = "on-demand-color"
    onDemandColorDynMpls.EntityData.SegmentPath = "on-demand-color-dyn-mpls"
    onDemandColorDynMpls.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/" + onDemandColorDynMpls.EntityData.SegmentPath
    onDemandColorDynMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMpls.EntityData.Children = types.NewOrderedMap()
    onDemandColorDynMpls.EntityData.Children.Append("on-demand-color-dyn-mpls-metric", types.YChild{"OnDemandColorDynMplsMetric", &onDemandColorDynMpls.OnDemandColorDynMplsMetric})
    onDemandColorDynMpls.EntityData.Children.Append("on-demand-color-dyn-mpls-pce", types.YChild{"OnDemandColorDynMplsPce", &onDemandColorDynMpls.OnDemandColorDynMplsPce})
    onDemandColorDynMpls.EntityData.Children.Append("disjoint-path", types.YChild{"DisjointPath", &onDemandColorDynMpls.DisjointPath})
    onDemandColorDynMpls.EntityData.Children.Append("affinity-rules", types.YChild{"AffinityRules", &onDemandColorDynMpls.AffinityRules})
    onDemandColorDynMpls.EntityData.Leafs = types.NewOrderedMap()
    onDemandColorDynMpls.EntityData.Leafs.Append("on-demand-color-dyn-mpls-flex-algorithm", types.YLeaf{"OnDemandColorDynMplsFlexAlgorithm", onDemandColorDynMpls.OnDemandColorDynMplsFlexAlgorithm})
    onDemandColorDynMpls.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", onDemandColorDynMpls.Enable})

    onDemandColorDynMpls.EntityData.YListKeys = []string {}

    return &(onDemandColorDynMpls.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric
// Metric type
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric Type. The type is XtcMetric.
    MetricType interface{}

    // Metric submode Enable. The type is interface{}.
    Enable interface{}

    // Metric Margin.
    MetricMargin Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric_MetricMargin
}

func (onDemandColorDynMplsMetric *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric) GetEntityData() *types.CommonEntityData {
    onDemandColorDynMplsMetric.EntityData.YFilter = onDemandColorDynMplsMetric.YFilter
    onDemandColorDynMplsMetric.EntityData.YangName = "on-demand-color-dyn-mpls-metric"
    onDemandColorDynMplsMetric.EntityData.BundleName = "cisco_ios_xr"
    onDemandColorDynMplsMetric.EntityData.ParentYangName = "on-demand-color-dyn-mpls"
    onDemandColorDynMplsMetric.EntityData.SegmentPath = "on-demand-color-dyn-mpls-metric"
    onDemandColorDynMplsMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/" + onDemandColorDynMplsMetric.EntityData.SegmentPath
    onDemandColorDynMplsMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMplsMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMplsMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMplsMetric.EntityData.Children = types.NewOrderedMap()
    onDemandColorDynMplsMetric.EntityData.Children.Append("metric-margin", types.YChild{"MetricMargin", &onDemandColorDynMplsMetric.MetricMargin})
    onDemandColorDynMplsMetric.EntityData.Leafs = types.NewOrderedMap()
    onDemandColorDynMplsMetric.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", onDemandColorDynMplsMetric.MetricType})
    onDemandColorDynMplsMetric.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", onDemandColorDynMplsMetric.Enable})

    onDemandColorDynMplsMetric.EntityData.YListKeys = []string {}

    return &(onDemandColorDynMplsMetric.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric_MetricMargin
// Metric Margin
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric_MetricMargin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric margin type. The type is XtcMetricValue.
    ValueType interface{}

    // Absolute metric value. The type is interface{} with range: 0..2147483647.
    AbsoluteValue interface{}

    // Relative metric value. The type is interface{} with range: 0..2147483647.
    RelativeValue interface{}
}

func (metricMargin *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric_MetricMargin) GetEntityData() *types.CommonEntityData {
    metricMargin.EntityData.YFilter = metricMargin.YFilter
    metricMargin.EntityData.YangName = "metric-margin"
    metricMargin.EntityData.BundleName = "cisco_ios_xr"
    metricMargin.EntityData.ParentYangName = "on-demand-color-dyn-mpls-metric"
    metricMargin.EntityData.SegmentPath = "metric-margin"
    metricMargin.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/on-demand-color-dyn-mpls-metric/" + metricMargin.EntityData.SegmentPath
    metricMargin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricMargin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricMargin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricMargin.EntityData.Children = types.NewOrderedMap()
    metricMargin.EntityData.Leafs = types.NewOrderedMap()
    metricMargin.EntityData.Leafs.Append("value-type", types.YLeaf{"ValueType", metricMargin.ValueType})
    metricMargin.EntityData.Leafs.Append("absolute-value", types.YLeaf{"AbsoluteValue", metricMargin.AbsoluteValue})
    metricMargin.EntityData.Leafs.Append("relative-value", types.YLeaf{"RelativeValue", metricMargin.RelativeValue})

    metricMargin.EntityData.YListKeys = []string {}

    return &(metricMargin.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsPce
// Use Path Computation Element
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsPce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE submode Enable. The type is interface{}.
    Enable interface{}
}

func (onDemandColorDynMplsPce *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsPce) GetEntityData() *types.CommonEntityData {
    onDemandColorDynMplsPce.EntityData.YFilter = onDemandColorDynMplsPce.YFilter
    onDemandColorDynMplsPce.EntityData.YangName = "on-demand-color-dyn-mpls-pce"
    onDemandColorDynMplsPce.EntityData.BundleName = "cisco_ios_xr"
    onDemandColorDynMplsPce.EntityData.ParentYangName = "on-demand-color-dyn-mpls"
    onDemandColorDynMplsPce.EntityData.SegmentPath = "on-demand-color-dyn-mpls-pce"
    onDemandColorDynMplsPce.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/" + onDemandColorDynMplsPce.EntityData.SegmentPath
    onDemandColorDynMplsPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMplsPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMplsPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMplsPce.EntityData.Children = types.NewOrderedMap()
    onDemandColorDynMplsPce.EntityData.Leafs = types.NewOrderedMap()
    onDemandColorDynMplsPce.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", onDemandColorDynMplsPce.Enable})

    onDemandColorDynMplsPce.EntityData.YListKeys = []string {}

    return &(onDemandColorDynMplsPce.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath
// Disjoint path
// This type is a presence type.
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Group ID. The type is interface{} with range: 1..65535. This attribute is
    // mandatory.
    GroupId interface{}

    // Disjointness Type. The type is XtcDisjointness. This attribute is
    // mandatory.
    DisjointnessType interface{}

    // Sub ID. The type is interface{} with range: 1..65535.
    SubId interface{}
}

func (disjointPath *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath) GetEntityData() *types.CommonEntityData {
    disjointPath.EntityData.YFilter = disjointPath.YFilter
    disjointPath.EntityData.YangName = "disjoint-path"
    disjointPath.EntityData.BundleName = "cisco_ios_xr"
    disjointPath.EntityData.ParentYangName = "on-demand-color-dyn-mpls"
    disjointPath.EntityData.SegmentPath = "disjoint-path"
    disjointPath.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/" + disjointPath.EntityData.SegmentPath
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = types.NewOrderedMap()
    disjointPath.EntityData.Leafs = types.NewOrderedMap()
    disjointPath.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", disjointPath.GroupId})
    disjointPath.EntityData.Leafs.Append("disjointness-type", types.YLeaf{"DisjointnessType", disjointPath.DisjointnessType})
    disjointPath.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", disjointPath.SubId})

    disjointPath.EntityData.YListKeys = []string {}

    return &(disjointPath.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules
// SR path computation and verification affinity
// rules
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SR path computation and verification affinity rule. The type is slice of
    // Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule.
    AffinityRule []*Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule
}

func (affinityRules *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules) GetEntityData() *types.CommonEntityData {
    affinityRules.EntityData.YFilter = affinityRules.YFilter
    affinityRules.EntityData.YangName = "affinity-rules"
    affinityRules.EntityData.BundleName = "cisco_ios_xr"
    affinityRules.EntityData.ParentYangName = "on-demand-color-dyn-mpls"
    affinityRules.EntityData.SegmentPath = "affinity-rules"
    affinityRules.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/" + affinityRules.EntityData.SegmentPath
    affinityRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRules.EntityData.Children = types.NewOrderedMap()
    affinityRules.EntityData.Children.Append("affinity-rule", types.YChild{"AffinityRule", nil})
    for i := range affinityRules.AffinityRule {
        affinityRules.EntityData.Children.Append(types.GetSegmentPath(affinityRules.AffinityRule[i]), types.YChild{"AffinityRule", affinityRules.AffinityRule[i]})
    }
    affinityRules.EntityData.Leafs = types.NewOrderedMap()

    affinityRules.EntityData.YListKeys = []string {}

    return &(affinityRules.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule
// SR path computation and verification affinity
// rule
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Affinity rule type. The type is XtcAffinityRule.
    Rule interface{}

    // Affinity rule name. The type is slice of
    // Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule_AffinityName.
    AffinityName []*Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule_AffinityName
}

func (affinityRule *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule) GetEntityData() *types.CommonEntityData {
    affinityRule.EntityData.YFilter = affinityRule.YFilter
    affinityRule.EntityData.YangName = "affinity-rule"
    affinityRule.EntityData.BundleName = "cisco_ios_xr"
    affinityRule.EntityData.ParentYangName = "affinity-rules"
    affinityRule.EntityData.SegmentPath = "affinity-rule" + types.AddKeyToken(affinityRule.Rule, "rule")
    affinityRule.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/affinity-rules/" + affinityRule.EntityData.SegmentPath
    affinityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRule.EntityData.Children = types.NewOrderedMap()
    affinityRule.EntityData.Children.Append("affinity-name", types.YChild{"AffinityName", nil})
    for i := range affinityRule.AffinityName {
        affinityRule.EntityData.Children.Append(types.GetSegmentPath(affinityRule.AffinityName[i]), types.YChild{"AffinityName", affinityRule.AffinityName[i]})
    }
    affinityRule.EntityData.Leafs = types.NewOrderedMap()
    affinityRule.EntityData.Leafs.Append("rule", types.YLeaf{"Rule", affinityRule.Rule})

    affinityRule.EntityData.YListKeys = []string {"Rule"}

    return &(affinityRule.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule_AffinityName
// Affinity rule name
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule_AffinityName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Affinity name. The type is string with length:
    // 1..32.
    AffinityName interface{}
}

func (affinityName *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_AffinityRules_AffinityRule_AffinityName) GetEntityData() *types.CommonEntityData {
    affinityName.EntityData.YFilter = affinityName.YFilter
    affinityName.EntityData.YangName = "affinity-name"
    affinityName.EntityData.BundleName = "cisco_ios_xr"
    affinityName.EntityData.ParentYangName = "affinity-rule"
    affinityName.EntityData.SegmentPath = "affinity-name" + types.AddKeyToken(affinityName.AffinityName, "affinity-name")
    affinityName.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/on-demand-colors/on-demand-color/on-demand-color-dyn-mpls/affinity-rules/affinity-rule/" + affinityName.EntityData.SegmentPath
    affinityName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityName.EntityData.Children = types.NewOrderedMap()
    affinityName.EntityData.Leafs = types.NewOrderedMap()
    affinityName.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityName.AffinityName})

    affinityName.EntityData.YListKeys = []string {"AffinityName"}

    return &(affinityName.EntityData)
}

// Sr_TrafficEngineering_Segments
// Segment-lists configuration
type Sr_TrafficEngineering_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment-list configuration. The type is slice of
    // Sr_TrafficEngineering_Segments_Segment.
    Segment []*Sr_TrafficEngineering_Segments_Segment
}

func (segments *Sr_TrafficEngineering_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "traffic-engineering"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + segments.EntityData.SegmentPath
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Children.Append("segment", types.YChild{"Segment", nil})
    for i := range segments.Segment {
        segments.EntityData.Children.Append(types.GetSegmentPath(segments.Segment[i]), types.YChild{"Segment", segments.Segment[i]})
    }
    segments.EntityData.Leafs = types.NewOrderedMap()

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment
// Segment-list configuration
type Sr_TrafficEngineering_Segments_Segment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    PathName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Segments/hops configuration for given Segment-list.
    Segments Sr_TrafficEngineering_Segments_Segment_Segments
}

func (segment *Sr_TrafficEngineering_Segments_Segment) GetEntityData() *types.CommonEntityData {
    segment.EntityData.YFilter = segment.YFilter
    segment.EntityData.YangName = "segment"
    segment.EntityData.BundleName = "cisco_ios_xr"
    segment.EntityData.ParentYangName = "segments"
    segment.EntityData.SegmentPath = "segment" + types.AddKeyToken(segment.PathName, "path-name")
    segment.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/segments/" + segment.EntityData.SegmentPath
    segment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment.EntityData.Children = types.NewOrderedMap()
    segment.EntityData.Children.Append("segments", types.YChild{"Segments", &segment.Segments})
    segment.EntityData.Leafs = types.NewOrderedMap()
    segment.EntityData.Leafs.Append("path-name", types.YLeaf{"PathName", segment.PathName})
    segment.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", segment.Enable})

    segment.EntityData.YListKeys = []string {"PathName"}

    return &(segment.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment_Segments
// Segments/hops configuration for given
// Segment-list
type Sr_TrafficEngineering_Segments_Segment_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Segment/hop at the index. The type is slice of
    // Sr_TrafficEngineering_Segments_Segment_Segments_Segment.
    Segment []*Sr_TrafficEngineering_Segments_Segment_Segments_Segment
}

func (segments *Sr_TrafficEngineering_Segments_Segment_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "segment"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/segments/segment/" + segments.EntityData.SegmentPath
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Children.Append("segment", types.YChild{"Segment", nil})
    for i := range segments.Segment {
        segments.EntityData.Children.Append(types.GetSegmentPath(segments.Segment[i]), types.YChild{"Segment", segments.Segment[i]})
    }
    segments.EntityData.Leafs = types.NewOrderedMap()

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment_Segments_Segment
// Configure Segment/hop at the index
type Sr_TrafficEngineering_Segments_Segment_Segments_Segment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Segment index. The type is interface{} with range:
    // 1..65535.
    SegmentIndex interface{}

    // Segment/hop type. The type is XtcSegment.
    SegmentType interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575.
    MplsLabel interface{}
}

func (segment *Sr_TrafficEngineering_Segments_Segment_Segments_Segment) GetEntityData() *types.CommonEntityData {
    segment.EntityData.YFilter = segment.YFilter
    segment.EntityData.YangName = "segment"
    segment.EntityData.BundleName = "cisco_ios_xr"
    segment.EntityData.ParentYangName = "segments"
    segment.EntityData.SegmentPath = "segment" + types.AddKeyToken(segment.SegmentIndex, "segment-index")
    segment.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/segments/segment/segments/" + segment.EntityData.SegmentPath
    segment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment.EntityData.Children = types.NewOrderedMap()
    segment.EntityData.Leafs = types.NewOrderedMap()
    segment.EntityData.Leafs.Append("segment-index", types.YLeaf{"SegmentIndex", segment.SegmentIndex})
    segment.EntityData.Leafs.Append("segment-type", types.YLeaf{"SegmentType", segment.SegmentType})
    segment.EntityData.Leafs.Append("address", types.YLeaf{"Address", segment.Address})
    segment.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", segment.MplsLabel})

    segment.EntityData.YListKeys = []string {"SegmentIndex"}

    return &(segment.EntityData)
}

// Sr_TrafficEngineering_Logging
// Logging configuration
// This type is a presence type.
type Sr_TrafficEngineering_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable logging for policy status. The type is interface{}. This attribute
    // is mandatory.
    PolicyStatus interface{}
}

func (logging *Sr_TrafficEngineering_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "traffic-engineering"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("policy-status", types.YLeaf{"PolicyStatus", logging.PolicyStatus})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Sr_TrafficEngineering_Timers
// SR-TE timers configuration
type Sr_TrafficEngineering_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay before cleaning up candidate paths. The type is interface{} with
    // range: 0..86400. The default value is 120.
    CandidatePathCleanupDelay interface{}

    // Delay before cleaning up previous path. The type is interface{} with range:
    // 0..300. The default value is 10.
    CleanupDelay interface{}

    // Delay before switching to a reoptimized path. The type is interface{} with
    // range: 0..300. The default value is 10.
    InstallDelay interface{}

    // How often to perform periodic reoptimization of policies. The type is
    // interface{} with range: 0..86400. The default value is 600.
    PeriodicReoptimization interface{}
}

func (timers *Sr_TrafficEngineering_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "traffic-engineering"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Leafs = types.NewOrderedMap()
    timers.EntityData.Leafs.Append("candidate-path-cleanup-delay", types.YLeaf{"CandidatePathCleanupDelay", timers.CandidatePathCleanupDelay})
    timers.EntityData.Leafs.Append("cleanup-delay", types.YLeaf{"CleanupDelay", timers.CleanupDelay})
    timers.EntityData.Leafs.Append("install-delay", types.YLeaf{"InstallDelay", timers.InstallDelay})
    timers.EntityData.Leafs.Append("periodic-reoptimization", types.YLeaf{"PeriodicReoptimization", timers.PeriodicReoptimization})

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Sr_TrafficEngineering_BindingSidRules
// Binding sid rules
type Sr_TrafficEngineering_BindingSidRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding sid explicit options.
    Explicit Sr_TrafficEngineering_BindingSidRules_Explicit

    // Dynamic binding SID options.
    DynamicBindingSidRules Sr_TrafficEngineering_BindingSidRules_DynamicBindingSidRules
}

func (bindingSidRules *Sr_TrafficEngineering_BindingSidRules) GetEntityData() *types.CommonEntityData {
    bindingSidRules.EntityData.YFilter = bindingSidRules.YFilter
    bindingSidRules.EntityData.YangName = "binding-sid-rules"
    bindingSidRules.EntityData.BundleName = "cisco_ios_xr"
    bindingSidRules.EntityData.ParentYangName = "traffic-engineering"
    bindingSidRules.EntityData.SegmentPath = "binding-sid-rules"
    bindingSidRules.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + bindingSidRules.EntityData.SegmentPath
    bindingSidRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSidRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSidRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSidRules.EntityData.Children = types.NewOrderedMap()
    bindingSidRules.EntityData.Children.Append("explicit", types.YChild{"Explicit", &bindingSidRules.Explicit})
    bindingSidRules.EntityData.Children.Append("dynamic-binding-sid-rules", types.YChild{"DynamicBindingSidRules", &bindingSidRules.DynamicBindingSidRules})
    bindingSidRules.EntityData.Leafs = types.NewOrderedMap()

    bindingSidRules.EntityData.YListKeys = []string {}

    return &(bindingSidRules.EntityData)
}

// Sr_TrafficEngineering_BindingSidRules_Explicit
// Binding sid explicit options
// This type is a presence type.
type Sr_TrafficEngineering_BindingSidRules_Explicit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Binding sid explicit rule. The type is XtcBindingSidexplicitRule. This
    // attribute is mandatory.
    Rule interface{}
}

func (explicit *Sr_TrafficEngineering_BindingSidRules_Explicit) GetEntityData() *types.CommonEntityData {
    explicit.EntityData.YFilter = explicit.YFilter
    explicit.EntityData.YangName = "explicit"
    explicit.EntityData.BundleName = "cisco_ios_xr"
    explicit.EntityData.ParentYangName = "binding-sid-rules"
    explicit.EntityData.SegmentPath = "explicit"
    explicit.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/binding-sid-rules/" + explicit.EntityData.SegmentPath
    explicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicit.EntityData.Children = types.NewOrderedMap()
    explicit.EntityData.Leafs = types.NewOrderedMap()
    explicit.EntityData.Leafs.Append("rule", types.YLeaf{"Rule", explicit.Rule})

    explicit.EntityData.YListKeys = []string {}

    return &(explicit.EntityData)
}

// Sr_TrafficEngineering_BindingSidRules_DynamicBindingSidRules
// Dynamic binding SID options
// This type is a presence type.
type Sr_TrafficEngineering_BindingSidRules_DynamicBindingSidRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Binding SID dynamic rule. The type is XtcBindingSidDynamicRule. This
    // attribute is mandatory.
    DynamicBindingSidRule interface{}
}

func (dynamicBindingSidRules *Sr_TrafficEngineering_BindingSidRules_DynamicBindingSidRules) GetEntityData() *types.CommonEntityData {
    dynamicBindingSidRules.EntityData.YFilter = dynamicBindingSidRules.YFilter
    dynamicBindingSidRules.EntityData.YangName = "dynamic-binding-sid-rules"
    dynamicBindingSidRules.EntityData.BundleName = "cisco_ios_xr"
    dynamicBindingSidRules.EntityData.ParentYangName = "binding-sid-rules"
    dynamicBindingSidRules.EntityData.SegmentPath = "dynamic-binding-sid-rules"
    dynamicBindingSidRules.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/binding-sid-rules/" + dynamicBindingSidRules.EntityData.SegmentPath
    dynamicBindingSidRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicBindingSidRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicBindingSidRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicBindingSidRules.EntityData.Children = types.NewOrderedMap()
    dynamicBindingSidRules.EntityData.Leafs = types.NewOrderedMap()
    dynamicBindingSidRules.EntityData.Leafs.Append("dynamic-binding-sid-rule", types.YLeaf{"DynamicBindingSidRule", dynamicBindingSidRules.DynamicBindingSidRule})

    dynamicBindingSidRules.EntityData.YListKeys = []string {}

    return &(dynamicBindingSidRules.EntityData)
}

// Sr_TrafficEngineering_Policies
// Policy configuration
type Sr_TrafficEngineering_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy.
    Policy []*Sr_TrafficEngineering_Policies_Policy
}

func (policies *Sr_TrafficEngineering_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "traffic-engineering"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + policies.EntityData.SegmentPath
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range policies.Policy {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.Policy[i]), types.YChild{"Policy", policies.Policy[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy
// Policy configuration
type Sr_TrafficEngineering_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Policy name. The type is string with length:
    // 1..59.
    PolicyName interface{}

    // Forward class associated with the policy. The type is interface{} with
    // range: 1..7.
    ForwardClass interface{}

    // IPv6 disable. The type is interface{}.
    Ipv6Disable interface{}

    // Administratively shutdown policy. The type is interface{}.
    Shutdown interface{}

    // The value of the bandwidth reserved by this policy in kbps. The type is
    // interface{} with range: 1..4294967295.
    Bandwidth interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Steering options for the policy.
    Steering Sr_TrafficEngineering_Policies_Policy_Steering

    // Binding Segment ID.
    BindingSid Sr_TrafficEngineering_Policies_Policy_BindingSid

    // Color and endpoint of a policyColor, EndPointType, Endpoint.
    PolicyColorEndpoint Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint

    // Autoroute configuration.
    AutoRoute Sr_TrafficEngineering_Policies_Policy_AutoRoute

    // Policy candidate-paths configuration.
    CandidatePaths Sr_TrafficEngineering_Policies_Policy_CandidatePaths
}

func (policy *Sr_TrafficEngineering_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.PolicyName, "policy-name")
    policy.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/" + policy.EntityData.SegmentPath
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Children.Append("steering", types.YChild{"Steering", &policy.Steering})
    policy.EntityData.Children.Append("binding-sid", types.YChild{"BindingSid", &policy.BindingSid})
    policy.EntityData.Children.Append("policy-color-endpoint", types.YChild{"PolicyColorEndpoint", &policy.PolicyColorEndpoint})
    policy.EntityData.Children.Append("auto-route", types.YChild{"AutoRoute", &policy.AutoRoute})
    policy.EntityData.Children.Append("candidate-paths", types.YChild{"CandidatePaths", &policy.CandidatePaths})
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policy.PolicyName})
    policy.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", policy.ForwardClass})
    policy.EntityData.Leafs.Append("ipv6-disable", types.YLeaf{"Ipv6Disable", policy.Ipv6Disable})
    policy.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", policy.Shutdown})
    policy.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", policy.Bandwidth})
    policy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", policy.Enable})

    policy.EntityData.YListKeys = []string {"PolicyName"}

    return &(policy.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_Steering
// Steering options for the policy
type Sr_TrafficEngineering_Policies_Policy_Steering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application table that steering options need to be applied.
    Applications Sr_TrafficEngineering_Policies_Policy_Steering_Applications
}

func (steering *Sr_TrafficEngineering_Policies_Policy_Steering) GetEntityData() *types.CommonEntityData {
    steering.EntityData.YFilter = steering.YFilter
    steering.EntityData.YangName = "steering"
    steering.EntityData.BundleName = "cisco_ios_xr"
    steering.EntityData.ParentYangName = "policy"
    steering.EntityData.SegmentPath = "steering"
    steering.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/" + steering.EntityData.SegmentPath
    steering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    steering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    steering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    steering.EntityData.Children = types.NewOrderedMap()
    steering.EntityData.Children.Append("applications", types.YChild{"Applications", &steering.Applications})
    steering.EntityData.Leafs = types.NewOrderedMap()

    steering.EntityData.YListKeys = []string {}

    return &(steering.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_Steering_Applications
// Application table that steering options need
// to be applied
type Sr_TrafficEngineering_Policies_Policy_Steering_Applications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application that steering options need to be applied. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_Steering_Applications_Application.
    Application []*Sr_TrafficEngineering_Policies_Policy_Steering_Applications_Application
}

func (applications *Sr_TrafficEngineering_Policies_Policy_Steering_Applications) GetEntityData() *types.CommonEntityData {
    applications.EntityData.YFilter = applications.YFilter
    applications.EntityData.YangName = "applications"
    applications.EntityData.BundleName = "cisco_ios_xr"
    applications.EntityData.ParentYangName = "steering"
    applications.EntityData.SegmentPath = "applications"
    applications.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/steering/" + applications.EntityData.SegmentPath
    applications.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    applications.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    applications.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    applications.EntityData.Children = types.NewOrderedMap()
    applications.EntityData.Children.Append("application", types.YChild{"Application", nil})
    for i := range applications.Application {
        applications.EntityData.Children.Append(types.GetSegmentPath(applications.Application[i]), types.YChild{"Application", applications.Application[i]})
    }
    applications.EntityData.Leafs = types.NewOrderedMap()

    applications.EntityData.YListKeys = []string {}

    return &(applications.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_Steering_Applications_Application
// Application that steering options need to
// be applied
type Sr_TrafficEngineering_Policies_Policy_Steering_Applications_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Steering application. The type is
    // XtcSteeringApplication.
    Application interface{}

    // Disable all steering services. The type is interface{}.
    Disable interface{}
}

func (application *Sr_TrafficEngineering_Policies_Policy_Steering_Applications_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "applications"
    application.EntityData.SegmentPath = "application" + types.AddKeyToken(application.Application, "application")
    application.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/steering/applications/" + application.EntityData.SegmentPath
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = types.NewOrderedMap()
    application.EntityData.Leafs = types.NewOrderedMap()
    application.EntityData.Leafs.Append("application", types.YLeaf{"Application", application.Application})
    application.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", application.Disable})

    application.EntityData.YListKeys = []string {"Application"}

    return &(application.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_BindingSid
// Binding Segment ID
type Sr_TrafficEngineering_Policies_Policy_BindingSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding SID type. The type is XtcBindingSid.
    BindingSidType interface{}

    // MPLS Label. The type is interface{} with range: 16..1048575.
    MplsLabel interface{}
}

func (bindingSid *Sr_TrafficEngineering_Policies_Policy_BindingSid) GetEntityData() *types.CommonEntityData {
    bindingSid.EntityData.YFilter = bindingSid.YFilter
    bindingSid.EntityData.YangName = "binding-sid"
    bindingSid.EntityData.BundleName = "cisco_ios_xr"
    bindingSid.EntityData.ParentYangName = "policy"
    bindingSid.EntityData.SegmentPath = "binding-sid"
    bindingSid.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/" + bindingSid.EntityData.SegmentPath
    bindingSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSid.EntityData.Children = types.NewOrderedMap()
    bindingSid.EntityData.Leafs = types.NewOrderedMap()
    bindingSid.EntityData.Leafs.Append("binding-sid-type", types.YLeaf{"BindingSidType", bindingSid.BindingSidType})
    bindingSid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", bindingSid.MplsLabel})

    bindingSid.EntityData.YListKeys = []string {}

    return &(bindingSid.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint
// Color and endpoint of a policyColor,
// EndPointType, Endpoint
// This type is a presence type.
type Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Color. The type is interface{} with range: 1..4294967295. This attribute is
    // mandatory.
    Color interface{}

    // End point type. The type is XtcEndPoint. This attribute is mandatory.
    EndPointType interface{}

    // End point address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    EndPointAddress interface{}
}

func (policyColorEndpoint *Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint) GetEntityData() *types.CommonEntityData {
    policyColorEndpoint.EntityData.YFilter = policyColorEndpoint.YFilter
    policyColorEndpoint.EntityData.YangName = "policy-color-endpoint"
    policyColorEndpoint.EntityData.BundleName = "cisco_ios_xr"
    policyColorEndpoint.EntityData.ParentYangName = "policy"
    policyColorEndpoint.EntityData.SegmentPath = "policy-color-endpoint"
    policyColorEndpoint.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/" + policyColorEndpoint.EntityData.SegmentPath
    policyColorEndpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyColorEndpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyColorEndpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyColorEndpoint.EntityData.Children = types.NewOrderedMap()
    policyColorEndpoint.EntityData.Leafs = types.NewOrderedMap()
    policyColorEndpoint.EntityData.Leafs.Append("color", types.YLeaf{"Color", policyColorEndpoint.Color})
    policyColorEndpoint.EntityData.Leafs.Append("end-point-type", types.YLeaf{"EndPointType", policyColorEndpoint.EndPointType})
    policyColorEndpoint.EntityData.Leafs.Append("end-point-address", types.YLeaf{"EndPointAddress", policyColorEndpoint.EndPointAddress})

    policyColorEndpoint.EntityData.YListKeys = []string {}

    return &(policyColorEndpoint.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_AutoRoute
// Autoroute configuration
type Sr_TrafficEngineering_Policies_Policy_AutoRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Force autoroute policy to be safe for carrying SR labelled traffic. The
    // type is interface{}.
    ForceSrInclude interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Autoroute metric.
    AutoRouteMetric Sr_TrafficEngineering_Policies_Policy_AutoRoute_AutoRouteMetric

    // Autoroute include prefix table configuration.
    IncludePrefixes Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes
}

func (autoRoute *Sr_TrafficEngineering_Policies_Policy_AutoRoute) GetEntityData() *types.CommonEntityData {
    autoRoute.EntityData.YFilter = autoRoute.YFilter
    autoRoute.EntityData.YangName = "auto-route"
    autoRoute.EntityData.BundleName = "cisco_ios_xr"
    autoRoute.EntityData.ParentYangName = "policy"
    autoRoute.EntityData.SegmentPath = "auto-route"
    autoRoute.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/" + autoRoute.EntityData.SegmentPath
    autoRoute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRoute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRoute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRoute.EntityData.Children = types.NewOrderedMap()
    autoRoute.EntityData.Children.Append("auto-route-metric", types.YChild{"AutoRouteMetric", &autoRoute.AutoRouteMetric})
    autoRoute.EntityData.Children.Append("include-prefixes", types.YChild{"IncludePrefixes", &autoRoute.IncludePrefixes})
    autoRoute.EntityData.Leafs = types.NewOrderedMap()
    autoRoute.EntityData.Leafs.Append("force-sr-include", types.YLeaf{"ForceSrInclude", autoRoute.ForceSrInclude})
    autoRoute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoRoute.Enable})

    autoRoute.EntityData.YListKeys = []string {}

    return &(autoRoute.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_AutoRoute_AutoRouteMetric
// Autoroute metric
type Sr_TrafficEngineering_Policies_Policy_AutoRoute_AutoRouteMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric type. The type is XtcAutoRouteMetric.
    AutorouteMetricType interface{}

    // Autoroute relative metric. The type is interface{} with range: -10..10.
    MetricRelativeValue interface{}

    // Autoroute constant metric. The type is interface{} with range:
    // 1..2147483647.
    MetricConstantValue interface{}
}

func (autoRouteMetric *Sr_TrafficEngineering_Policies_Policy_AutoRoute_AutoRouteMetric) GetEntityData() *types.CommonEntityData {
    autoRouteMetric.EntityData.YFilter = autoRouteMetric.YFilter
    autoRouteMetric.EntityData.YangName = "auto-route-metric"
    autoRouteMetric.EntityData.BundleName = "cisco_ios_xr"
    autoRouteMetric.EntityData.ParentYangName = "auto-route"
    autoRouteMetric.EntityData.SegmentPath = "auto-route-metric"
    autoRouteMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/auto-route/" + autoRouteMetric.EntityData.SegmentPath
    autoRouteMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRouteMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRouteMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRouteMetric.EntityData.Children = types.NewOrderedMap()
    autoRouteMetric.EntityData.Leafs = types.NewOrderedMap()
    autoRouteMetric.EntityData.Leafs.Append("autoroute-metric-type", types.YLeaf{"AutorouteMetricType", autoRouteMetric.AutorouteMetricType})
    autoRouteMetric.EntityData.Leafs.Append("metric-relative-value", types.YLeaf{"MetricRelativeValue", autoRouteMetric.MetricRelativeValue})
    autoRouteMetric.EntityData.Leafs.Append("metric-constant-value", types.YLeaf{"MetricConstantValue", autoRouteMetric.MetricConstantValue})

    autoRouteMetric.EntityData.YListKeys = []string {}

    return &(autoRouteMetric.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes
// Autoroute include prefix table configuration
type Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autoroute IP prefix to include. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes_IncludePrefix.
    IncludePrefix []*Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes_IncludePrefix
}

func (includePrefixes *Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes) GetEntityData() *types.CommonEntityData {
    includePrefixes.EntityData.YFilter = includePrefixes.YFilter
    includePrefixes.EntityData.YangName = "include-prefixes"
    includePrefixes.EntityData.BundleName = "cisco_ios_xr"
    includePrefixes.EntityData.ParentYangName = "auto-route"
    includePrefixes.EntityData.SegmentPath = "include-prefixes"
    includePrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/auto-route/" + includePrefixes.EntityData.SegmentPath
    includePrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includePrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includePrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includePrefixes.EntityData.Children = types.NewOrderedMap()
    includePrefixes.EntityData.Children.Append("include-prefix", types.YChild{"IncludePrefix", nil})
    for i := range includePrefixes.IncludePrefix {
        includePrefixes.EntityData.Children.Append(types.GetSegmentPath(includePrefixes.IncludePrefix[i]), types.YChild{"IncludePrefix", includePrefixes.IncludePrefix[i]})
    }
    includePrefixes.EntityData.Leafs = types.NewOrderedMap()

    includePrefixes.EntityData.YListKeys = []string {}

    return &(includePrefixes.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes_IncludePrefix
// Autoroute IP prefix to include
type Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes_IncludePrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address family type. The type is XtcAddressFamily.
    AfType interface{}

    // This attribute is a key. Autoroute prefix IP address. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixAddress interface{}

    // This attribute is a key. Autoroute IP prefix length. The type is
    // interface{} with range: 0..32.
    PrefixLength interface{}
}

func (includePrefix *Sr_TrafficEngineering_Policies_Policy_AutoRoute_IncludePrefixes_IncludePrefix) GetEntityData() *types.CommonEntityData {
    includePrefix.EntityData.YFilter = includePrefix.YFilter
    includePrefix.EntityData.YangName = "include-prefix"
    includePrefix.EntityData.BundleName = "cisco_ios_xr"
    includePrefix.EntityData.ParentYangName = "include-prefixes"
    includePrefix.EntityData.SegmentPath = "include-prefix" + types.AddKeyToken(includePrefix.AfType, "af-type") + types.AddKeyToken(includePrefix.PrefixAddress, "prefix-address") + types.AddKeyToken(includePrefix.PrefixLength, "prefix-length")
    includePrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/auto-route/include-prefixes/" + includePrefix.EntityData.SegmentPath
    includePrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includePrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includePrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includePrefix.EntityData.Children = types.NewOrderedMap()
    includePrefix.EntityData.Leafs = types.NewOrderedMap()
    includePrefix.EntityData.Leafs.Append("af-type", types.YLeaf{"AfType", includePrefix.AfType})
    includePrefix.EntityData.Leafs.Append("prefix-address", types.YLeaf{"PrefixAddress", includePrefix.PrefixAddress})
    includePrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", includePrefix.PrefixLength})

    includePrefix.EntityData.YListKeys = []string {"AfType", "PrefixAddress", "PrefixLength"}

    return &(includePrefix.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths
// Policy candidate-paths configuration
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // Policy path-option preference table.
    Preferences Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences
}

func (candidatePaths *Sr_TrafficEngineering_Policies_Policy_CandidatePaths) GetEntityData() *types.CommonEntityData {
    candidatePaths.EntityData.YFilter = candidatePaths.YFilter
    candidatePaths.EntityData.YangName = "candidate-paths"
    candidatePaths.EntityData.BundleName = "cisco_ios_xr"
    candidatePaths.EntityData.ParentYangName = "policy"
    candidatePaths.EntityData.SegmentPath = "candidate-paths"
    candidatePaths.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/" + candidatePaths.EntityData.SegmentPath
    candidatePaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidatePaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidatePaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidatePaths.EntityData.Children = types.NewOrderedMap()
    candidatePaths.EntityData.Children.Append("preferences", types.YChild{"Preferences", &candidatePaths.Preferences})
    candidatePaths.EntityData.Leafs = types.NewOrderedMap()
    candidatePaths.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", candidatePaths.Enable})

    candidatePaths.EntityData.YListKeys = []string {}

    return &(candidatePaths.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences
// Policy path-option preference table
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy path-option preference entry. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference.
    Preference []*Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference
}

func (preferences *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences) GetEntityData() *types.CommonEntityData {
    preferences.EntityData.YFilter = preferences.YFilter
    preferences.EntityData.YangName = "preferences"
    preferences.EntityData.BundleName = "cisco_ios_xr"
    preferences.EntityData.ParentYangName = "candidate-paths"
    preferences.EntityData.SegmentPath = "preferences"
    preferences.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/" + preferences.EntityData.SegmentPath
    preferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preferences.EntityData.Children = types.NewOrderedMap()
    preferences.EntityData.Children.Append("preference", types.YChild{"Preference", nil})
    for i := range preferences.Preference {
        preferences.EntityData.Children.Append(types.GetSegmentPath(preferences.Preference[i]), types.YChild{"Preference", preferences.Preference[i]})
    }
    preferences.EntityData.Leafs = types.NewOrderedMap()

    preferences.EntityData.YListKeys = []string {}

    return &(preferences.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference
// Policy path-option preference entry
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Path-option preference. The type is interface{}
    // with range: 1..65535.
    PathIndex interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // SR path computation and verification constraints.
    Constraints Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints

    // Policy path-option preference configuration.
    PathInfos Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos
}

func (preference *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference) GetEntityData() *types.CommonEntityData {
    preference.EntityData.YFilter = preference.YFilter
    preference.EntityData.YangName = "preference"
    preference.EntityData.BundleName = "cisco_ios_xr"
    preference.EntityData.ParentYangName = "preferences"
    preference.EntityData.SegmentPath = "preference" + types.AddKeyToken(preference.PathIndex, "path-index")
    preference.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/" + preference.EntityData.SegmentPath
    preference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preference.EntityData.Children = types.NewOrderedMap()
    preference.EntityData.Children.Append("constraints", types.YChild{"Constraints", &preference.Constraints})
    preference.EntityData.Children.Append("path-infos", types.YChild{"PathInfos", &preference.PathInfos})
    preference.EntityData.Leafs = types.NewOrderedMap()
    preference.EntityData.Leafs.Append("path-index", types.YLeaf{"PathIndex", preference.PathIndex})
    preference.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", preference.Enable})

    preference.EntityData.YListKeys = []string {"PathIndex"}

    return &(preference.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints
// SR path computation and verification
// constraints
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disjoint path.
    DisjointPath Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_DisjointPath

    // SR path computation segment specific rules.
    SegmentRules Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_SegmentRules

    // SR path computation and verification affinity rules.
    AffinityRules Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules
}

func (constraints *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints) GetEntityData() *types.CommonEntityData {
    constraints.EntityData.YFilter = constraints.YFilter
    constraints.EntityData.YangName = "constraints"
    constraints.EntityData.BundleName = "cisco_ios_xr"
    constraints.EntityData.ParentYangName = "preference"
    constraints.EntityData.SegmentPath = "constraints"
    constraints.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/" + constraints.EntityData.SegmentPath
    constraints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    constraints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    constraints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    constraints.EntityData.Children = types.NewOrderedMap()
    constraints.EntityData.Children.Append("disjoint-path", types.YChild{"DisjointPath", &constraints.DisjointPath})
    constraints.EntityData.Children.Append("segment-rules", types.YChild{"SegmentRules", &constraints.SegmentRules})
    constraints.EntityData.Children.Append("affinity-rules", types.YChild{"AffinityRules", &constraints.AffinityRules})
    constraints.EntityData.Leafs = types.NewOrderedMap()

    constraints.EntityData.YListKeys = []string {}

    return &(constraints.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_DisjointPath
// Disjoint path
// This type is a presence type.
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Group ID. The type is interface{} with range: 1..65535. This attribute is
    // mandatory.
    GroupId interface{}

    // Disjointness Type. The type is XtcDisjointness. This attribute is
    // mandatory.
    DisjointnessType interface{}

    // Sub ID. The type is interface{} with range: 1..65535.
    SubId interface{}
}

func (disjointPath *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_DisjointPath) GetEntityData() *types.CommonEntityData {
    disjointPath.EntityData.YFilter = disjointPath.YFilter
    disjointPath.EntityData.YangName = "disjoint-path"
    disjointPath.EntityData.BundleName = "cisco_ios_xr"
    disjointPath.EntityData.ParentYangName = "constraints"
    disjointPath.EntityData.SegmentPath = "disjoint-path"
    disjointPath.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/constraints/" + disjointPath.EntityData.SegmentPath
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = types.NewOrderedMap()
    disjointPath.EntityData.Leafs = types.NewOrderedMap()
    disjointPath.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", disjointPath.GroupId})
    disjointPath.EntityData.Leafs.Append("disjointness-type", types.YLeaf{"DisjointnessType", disjointPath.DisjointnessType})
    disjointPath.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", disjointPath.SubId})

    disjointPath.EntityData.YListKeys = []string {}

    return &(disjointPath.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_SegmentRules
// SR path computation segment specific
// rules
// This type is a presence type.
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_SegmentRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Prefix-SID algorithm. The type is interface{} with range: 128..255. This
    // attribute is mandatory.
    SidAlgorithm interface{}
}

func (segmentRules *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_SegmentRules) GetEntityData() *types.CommonEntityData {
    segmentRules.EntityData.YFilter = segmentRules.YFilter
    segmentRules.EntityData.YangName = "segment-rules"
    segmentRules.EntityData.BundleName = "cisco_ios_xr"
    segmentRules.EntityData.ParentYangName = "constraints"
    segmentRules.EntityData.SegmentPath = "segment-rules"
    segmentRules.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/constraints/" + segmentRules.EntityData.SegmentPath
    segmentRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRules.EntityData.Children = types.NewOrderedMap()
    segmentRules.EntityData.Leafs = types.NewOrderedMap()
    segmentRules.EntityData.Leafs.Append("sid-algorithm", types.YLeaf{"SidAlgorithm", segmentRules.SidAlgorithm})

    segmentRules.EntityData.YListKeys = []string {}

    return &(segmentRules.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules
// SR path computation and verification affinity
// rules
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SR path computation and verification affinity rule. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule.
    AffinityRule []*Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule
}

func (affinityRules *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules) GetEntityData() *types.CommonEntityData {
    affinityRules.EntityData.YFilter = affinityRules.YFilter
    affinityRules.EntityData.YangName = "affinity-rules"
    affinityRules.EntityData.BundleName = "cisco_ios_xr"
    affinityRules.EntityData.ParentYangName = "constraints"
    affinityRules.EntityData.SegmentPath = "affinity-rules"
    affinityRules.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/constraints/" + affinityRules.EntityData.SegmentPath
    affinityRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRules.EntityData.Children = types.NewOrderedMap()
    affinityRules.EntityData.Children.Append("affinity-rule", types.YChild{"AffinityRule", nil})
    for i := range affinityRules.AffinityRule {
        affinityRules.EntityData.Children.Append(types.GetSegmentPath(affinityRules.AffinityRule[i]), types.YChild{"AffinityRule", affinityRules.AffinityRule[i]})
    }
    affinityRules.EntityData.Leafs = types.NewOrderedMap()

    affinityRules.EntityData.YListKeys = []string {}

    return &(affinityRules.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule
// SR path computation and verification affinity
// rule
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Affinity rule type. The type is XtcAffinityRule.
    Rule interface{}

    // Affinity rule name. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule_AffinityName.
    AffinityName []*Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule_AffinityName
}

func (affinityRule *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule) GetEntityData() *types.CommonEntityData {
    affinityRule.EntityData.YFilter = affinityRule.YFilter
    affinityRule.EntityData.YangName = "affinity-rule"
    affinityRule.EntityData.BundleName = "cisco_ios_xr"
    affinityRule.EntityData.ParentYangName = "affinity-rules"
    affinityRule.EntityData.SegmentPath = "affinity-rule" + types.AddKeyToken(affinityRule.Rule, "rule")
    affinityRule.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/constraints/affinity-rules/" + affinityRule.EntityData.SegmentPath
    affinityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRule.EntityData.Children = types.NewOrderedMap()
    affinityRule.EntityData.Children.Append("affinity-name", types.YChild{"AffinityName", nil})
    for i := range affinityRule.AffinityName {
        affinityRule.EntityData.Children.Append(types.GetSegmentPath(affinityRule.AffinityName[i]), types.YChild{"AffinityName", affinityRule.AffinityName[i]})
    }
    affinityRule.EntityData.Leafs = types.NewOrderedMap()
    affinityRule.EntityData.Leafs.Append("rule", types.YLeaf{"Rule", affinityRule.Rule})

    affinityRule.EntityData.YListKeys = []string {"Rule"}

    return &(affinityRule.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule_AffinityName
// Affinity rule name
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule_AffinityName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Affinity name. The type is string with length:
    // 1..32.
    AffinityName interface{}
}

func (affinityName *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule_AffinityName) GetEntityData() *types.CommonEntityData {
    affinityName.EntityData.YFilter = affinityName.YFilter
    affinityName.EntityData.YangName = "affinity-name"
    affinityName.EntityData.BundleName = "cisco_ios_xr"
    affinityName.EntityData.ParentYangName = "affinity-rule"
    affinityName.EntityData.SegmentPath = "affinity-name" + types.AddKeyToken(affinityName.AffinityName, "affinity-name")
    affinityName.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/constraints/affinity-rules/affinity-rule/" + affinityName.EntityData.SegmentPath
    affinityName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityName.EntityData.Children = types.NewOrderedMap()
    affinityName.EntityData.Leafs = types.NewOrderedMap()
    affinityName.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityName.AffinityName})

    affinityName.EntityData.YListKeys = []string {"AffinityName"}

    return &(affinityName.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos
// Policy path-option preference
// configuration
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo.
    PathInfo []*Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
}

func (pathInfos *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos) GetEntityData() *types.CommonEntityData {
    pathInfos.EntityData.YFilter = pathInfos.YFilter
    pathInfos.EntityData.YangName = "path-infos"
    pathInfos.EntityData.BundleName = "cisco_ios_xr"
    pathInfos.EntityData.ParentYangName = "preference"
    pathInfos.EntityData.SegmentPath = "path-infos"
    pathInfos.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/" + pathInfos.EntityData.SegmentPath
    pathInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfos.EntityData.Children = types.NewOrderedMap()
    pathInfos.EntityData.Children.Append("path-info", types.YChild{"PathInfo", nil})
    for i := range pathInfos.PathInfo {
        pathInfos.EntityData.Children.Append(types.GetSegmentPath(pathInfos.PathInfo[i]), types.YChild{"PathInfo", pathInfos.PathInfo[i]})
    }
    pathInfos.EntityData.Leafs = types.NewOrderedMap()

    pathInfos.EntityData.YListKeys = []string {}

    return &(pathInfos.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
// Policy configuration
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Path-option type. The type is XtcPath.
    Type interface{}

    // This attribute is a key. Type of dynamic path to be computed. The type is
    // XtcPathHop.
    HopType interface{}

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    SegmentListName interface{}

    // Path-option weight. The type is interface{} with range: 1..4294967295.
    Weight interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Metric configuration, valid only for dynamic path-options.
    Metric Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric

    // Path Computation Element Protocol.
    Pcep Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pcep
}

func (pathInfo *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "path-infos"
    pathInfo.EntityData.SegmentPath = "path-info" + types.AddKeyToken(pathInfo.Type, "type") + types.AddKeyToken(pathInfo.HopType, "hop-type") + types.AddKeyToken(pathInfo.SegmentListName, "segment-list-name")
    pathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/path-infos/" + pathInfo.EntityData.SegmentPath
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = types.NewOrderedMap()
    pathInfo.EntityData.Children.Append("metric", types.YChild{"Metric", &pathInfo.Metric})
    pathInfo.EntityData.Children.Append("pcep", types.YChild{"Pcep", &pathInfo.Pcep})
    pathInfo.EntityData.Leafs = types.NewOrderedMap()
    pathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathInfo.Type})
    pathInfo.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", pathInfo.HopType})
    pathInfo.EntityData.Leafs.Append("segment-list-name", types.YLeaf{"SegmentListName", pathInfo.SegmentListName})
    pathInfo.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", pathInfo.Weight})
    pathInfo.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathInfo.Enable})

    pathInfo.EntityData.YListKeys = []string {"Type", "HopType", "SegmentListName"}

    return &(pathInfo.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric
// Metric configuration, valid only for
// dynamic path-options
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of SIDs. The type is interface{} with range: 1..255.
    SidLimit interface{}

    // Metric type. The type is XtcMetric.
    MetricType interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Metric Margin.
    MetricMargin Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_MetricMargin
}

func (metric *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "path-info"
    metric.EntityData.SegmentPath = "metric"
    metric.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/path-infos/path-info/" + metric.EntityData.SegmentPath
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = types.NewOrderedMap()
    metric.EntityData.Children.Append("metric-margin", types.YChild{"MetricMargin", &metric.MetricMargin})
    metric.EntityData.Leafs = types.NewOrderedMap()
    metric.EntityData.Leafs.Append("sid-limit", types.YLeaf{"SidLimit", metric.SidLimit})
    metric.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", metric.MetricType})
    metric.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", metric.Enable})

    metric.EntityData.YListKeys = []string {}

    return &(metric.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_MetricMargin
// Metric Margin
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_MetricMargin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric margin type. The type is XtcMetricValue.
    ValueType interface{}

    // Absolute metric value. The type is interface{} with range: 0..2147483647.
    AbsoluteValue interface{}

    // Relative metric value. The type is interface{} with range: 0..2147483647.
    RelativeValue interface{}
}

func (metricMargin *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_MetricMargin) GetEntityData() *types.CommonEntityData {
    metricMargin.EntityData.YFilter = metricMargin.YFilter
    metricMargin.EntityData.YangName = "metric-margin"
    metricMargin.EntityData.BundleName = "cisco_ios_xr"
    metricMargin.EntityData.ParentYangName = "metric"
    metricMargin.EntityData.SegmentPath = "metric-margin"
    metricMargin.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/path-infos/path-info/metric/" + metricMargin.EntityData.SegmentPath
    metricMargin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metricMargin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metricMargin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metricMargin.EntityData.Children = types.NewOrderedMap()
    metricMargin.EntityData.Leafs = types.NewOrderedMap()
    metricMargin.EntityData.Leafs.Append("value-type", types.YLeaf{"ValueType", metricMargin.ValueType})
    metricMargin.EntityData.Leafs.Append("absolute-value", types.YLeaf{"AbsoluteValue", metricMargin.AbsoluteValue})
    metricMargin.EntityData.Leafs.Append("relative-value", types.YLeaf{"RelativeValue", metricMargin.RelativeValue})

    metricMargin.EntityData.YListKeys = []string {}

    return &(metricMargin.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pcep
// Path Computation Element Protocol
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pcep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}
}

func (pcep *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pcep) GetEntityData() *types.CommonEntityData {
    pcep.EntityData.YFilter = pcep.YFilter
    pcep.EntityData.YangName = "pcep"
    pcep.EntityData.BundleName = "cisco_ios_xr"
    pcep.EntityData.ParentYangName = "path-info"
    pcep.EntityData.SegmentPath = "pcep"
    pcep.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/policies/policy/candidate-paths/preferences/preference/path-infos/path-info/" + pcep.EntityData.SegmentPath
    pcep.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcep.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcep.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcep.EntityData.Children = types.NewOrderedMap()
    pcep.EntityData.Leafs = types.NewOrderedMap()
    pcep.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pcep.Enable})

    pcep.EntityData.YListKeys = []string {}

    return &(pcep.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces
// SR-TE interfaces
type Sr_TrafficEngineering_SrteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SR-TE interface. The type is slice of
    // Sr_TrafficEngineering_SrteInterfaces_SrteInterface.
    SrteInterface []*Sr_TrafficEngineering_SrteInterfaces_SrteInterface
}

func (srteInterfaces *Sr_TrafficEngineering_SrteInterfaces) GetEntityData() *types.CommonEntityData {
    srteInterfaces.EntityData.YFilter = srteInterfaces.YFilter
    srteInterfaces.EntityData.YangName = "srte-interfaces"
    srteInterfaces.EntityData.BundleName = "cisco_ios_xr"
    srteInterfaces.EntityData.ParentYangName = "traffic-engineering"
    srteInterfaces.EntityData.SegmentPath = "srte-interfaces"
    srteInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + srteInterfaces.EntityData.SegmentPath
    srteInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srteInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srteInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srteInterfaces.EntityData.Children = types.NewOrderedMap()
    srteInterfaces.EntityData.Children.Append("srte-interface", types.YChild{"SrteInterface", nil})
    for i := range srteInterfaces.SrteInterface {
        srteInterfaces.EntityData.Children.Append(types.GetSegmentPath(srteInterfaces.SrteInterface[i]), types.YChild{"SrteInterface", srteInterfaces.SrteInterface[i]})
    }
    srteInterfaces.EntityData.Leafs = types.NewOrderedMap()

    srteInterfaces.EntityData.YListKeys = []string {}

    return &(srteInterfaces.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface
// SR-TE interface
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. SR-TE Interface name. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    SrteInterfaceName interface{}

    // Interface TE metric configuration. The type is interface{} with range:
    // 0..2147483647.
    InterfaceMetric interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Set user defined interface attribute flags.
    InterfaceAffinities Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities
}

func (srteInterface *Sr_TrafficEngineering_SrteInterfaces_SrteInterface) GetEntityData() *types.CommonEntityData {
    srteInterface.EntityData.YFilter = srteInterface.YFilter
    srteInterface.EntityData.YangName = "srte-interface"
    srteInterface.EntityData.BundleName = "cisco_ios_xr"
    srteInterface.EntityData.ParentYangName = "srte-interfaces"
    srteInterface.EntityData.SegmentPath = "srte-interface" + types.AddKeyToken(srteInterface.SrteInterfaceName, "srte-interface-name")
    srteInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/srte-interfaces/" + srteInterface.EntityData.SegmentPath
    srteInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srteInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srteInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srteInterface.EntityData.Children = types.NewOrderedMap()
    srteInterface.EntityData.Children.Append("interface-affinities", types.YChild{"InterfaceAffinities", &srteInterface.InterfaceAffinities})
    srteInterface.EntityData.Leafs = types.NewOrderedMap()
    srteInterface.EntityData.Leafs.Append("srte-interface-name", types.YLeaf{"SrteInterfaceName", srteInterface.SrteInterfaceName})
    srteInterface.EntityData.Leafs.Append("interface-metric", types.YLeaf{"InterfaceMetric", srteInterface.InterfaceMetric})
    srteInterface.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", srteInterface.Enable})

    srteInterface.EntityData.YListKeys = []string {"SrteInterfaceName"}

    return &(srteInterface.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities
// Set user defined interface attribute flags
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set user defined interface attribute flags. The type is slice of
    // Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity.
    InterfaceAffinity []*Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity
}

func (interfaceAffinities *Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities) GetEntityData() *types.CommonEntityData {
    interfaceAffinities.EntityData.YFilter = interfaceAffinities.YFilter
    interfaceAffinities.EntityData.YangName = "interface-affinities"
    interfaceAffinities.EntityData.BundleName = "cisco_ios_xr"
    interfaceAffinities.EntityData.ParentYangName = "srte-interface"
    interfaceAffinities.EntityData.SegmentPath = "interface-affinities"
    interfaceAffinities.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/srte-interfaces/srte-interface/" + interfaceAffinities.EntityData.SegmentPath
    interfaceAffinities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAffinities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAffinities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAffinities.EntityData.Children = types.NewOrderedMap()
    interfaceAffinities.EntityData.Children.Append("interface-affinity", types.YChild{"InterfaceAffinity", nil})
    for i := range interfaceAffinities.InterfaceAffinity {
        interfaceAffinities.EntityData.Children.Append(types.GetSegmentPath(interfaceAffinities.InterfaceAffinity[i]), types.YChild{"InterfaceAffinity", interfaceAffinities.InterfaceAffinity[i]})
    }
    interfaceAffinities.EntityData.Leafs = types.NewOrderedMap()

    interfaceAffinities.EntityData.YListKeys = []string {}

    return &(interfaceAffinities.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity
// Set user defined interface attribute flags
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface affinity names. The type is string with
    // length: 1..32.
    AffinityName interface{}
}

func (interfaceAffinity *Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity) GetEntityData() *types.CommonEntityData {
    interfaceAffinity.EntityData.YFilter = interfaceAffinity.YFilter
    interfaceAffinity.EntityData.YangName = "interface-affinity"
    interfaceAffinity.EntityData.BundleName = "cisco_ios_xr"
    interfaceAffinity.EntityData.ParentYangName = "interface-affinities"
    interfaceAffinity.EntityData.SegmentPath = "interface-affinity" + types.AddKeyToken(interfaceAffinity.AffinityName, "affinity-name")
    interfaceAffinity.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/srte-interfaces/srte-interface/interface-affinities/" + interfaceAffinity.EntityData.SegmentPath
    interfaceAffinity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAffinity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAffinity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAffinity.EntityData.Children = types.NewOrderedMap()
    interfaceAffinity.EntityData.Leafs = types.NewOrderedMap()
    interfaceAffinity.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", interfaceAffinity.AffinityName})

    interfaceAffinity.EntityData.YListKeys = []string {"AffinityName"}

    return &(interfaceAffinity.EntityData)
}

// Sr_TrafficEngineering_Pcc
// Path Computation Client
type Sr_TrafficEngineering_Pcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amount of time after which the peer can declare this session down, if no
    // PCEP message has been received. The type is interface{} with range: 1..255.
    DeadTimerInterval interface{}

    // Enable PCC centric model, where PCC only allows the lowest precedence PCE
    // to initiate policies. The type is interface{}.
    PccCentric interface{}

    // Report all local SR policies to connected PCEP peers. The type is
    // interface{}.
    ReportAll interface{}

    // Maximum time between two consecutive PCEP messages sent by this node. The
    // type is interface{} with range: 0..255.
    KeepaliveTimerInterval interface{}

    // Amount of time a PCE Initiated policy can remain orphan. The type is
    // interface{} with range: 15..14400.
    InitiatedStateInterval interface{}

    // Local source IP address to use on PCEP sessions. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // PCC Enable. The type is interface{}.
    Enable interface{}

    // Amount of time that a policy will be owned by a PCE after that PCE has gone
    // down. The type is interface{} with range: 10..180.
    InitiatedOrphanInterval interface{}

    // The maximum time delegated SR-TE policies can remain up without an active
    // connection to a PCE. The type is interface{} with range: 0..3600.
    DelegationTimeout interface{}

    // PCE peer configuration.
    PcePeers Sr_TrafficEngineering_Pcc_PcePeers

    // PCE peer configuration.
    PceAddresses Sr_TrafficEngineering_Pcc_PceAddresses
}

func (pcc *Sr_TrafficEngineering_Pcc) GetEntityData() *types.CommonEntityData {
    pcc.EntityData.YFilter = pcc.YFilter
    pcc.EntityData.YangName = "pcc"
    pcc.EntityData.BundleName = "cisco_ios_xr"
    pcc.EntityData.ParentYangName = "traffic-engineering"
    pcc.EntityData.SegmentPath = "pcc"
    pcc.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + pcc.EntityData.SegmentPath
    pcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcc.EntityData.Children = types.NewOrderedMap()
    pcc.EntityData.Children.Append("pce-peers", types.YChild{"PcePeers", &pcc.PcePeers})
    pcc.EntityData.Children.Append("pce-addresses", types.YChild{"PceAddresses", &pcc.PceAddresses})
    pcc.EntityData.Leafs = types.NewOrderedMap()
    pcc.EntityData.Leafs.Append("dead-timer-interval", types.YLeaf{"DeadTimerInterval", pcc.DeadTimerInterval})
    pcc.EntityData.Leafs.Append("pcc-centric", types.YLeaf{"PccCentric", pcc.PccCentric})
    pcc.EntityData.Leafs.Append("report-all", types.YLeaf{"ReportAll", pcc.ReportAll})
    pcc.EntityData.Leafs.Append("keepalive-timer-interval", types.YLeaf{"KeepaliveTimerInterval", pcc.KeepaliveTimerInterval})
    pcc.EntityData.Leafs.Append("initiated-state-interval", types.YLeaf{"InitiatedStateInterval", pcc.InitiatedStateInterval})
    pcc.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", pcc.SourceAddress})
    pcc.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pcc.Enable})
    pcc.EntityData.Leafs.Append("initiated-orphan-interval", types.YLeaf{"InitiatedOrphanInterval", pcc.InitiatedOrphanInterval})
    pcc.EntityData.Leafs.Append("delegation-timeout", types.YLeaf{"DelegationTimeout", pcc.DelegationTimeout})

    pcc.EntityData.YListKeys = []string {}

    return &(pcc.EntityData)
}

// Sr_TrafficEngineering_Pcc_PcePeers
// PCE peer configuration
type Sr_TrafficEngineering_Pcc_PcePeers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer. The type is slice of Sr_TrafficEngineering_Pcc_PcePeers_PcePeer.
    PcePeer []*Sr_TrafficEngineering_Pcc_PcePeers_PcePeer
}

func (pcePeers *Sr_TrafficEngineering_Pcc_PcePeers) GetEntityData() *types.CommonEntityData {
    pcePeers.EntityData.YFilter = pcePeers.YFilter
    pcePeers.EntityData.YangName = "pce-peers"
    pcePeers.EntityData.BundleName = "cisco_ios_xr"
    pcePeers.EntityData.ParentYangName = "pcc"
    pcePeers.EntityData.SegmentPath = "pce-peers"
    pcePeers.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/pcc/" + pcePeers.EntityData.SegmentPath
    pcePeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcePeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcePeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcePeers.EntityData.Children = types.NewOrderedMap()
    pcePeers.EntityData.Children.Append("pce-peer", types.YChild{"PcePeer", nil})
    for i := range pcePeers.PcePeer {
        pcePeers.EntityData.Children.Append(types.GetSegmentPath(pcePeers.PcePeer[i]), types.YChild{"PcePeer", pcePeers.PcePeer[i]})
    }
    pcePeers.EntityData.Leafs = types.NewOrderedMap()

    pcePeers.EntityData.YListKeys = []string {}

    return &(pcePeers.EntityData)
}

// Sr_TrafficEngineering_Pcc_PcePeers_PcePeer
// PCE peer
type Sr_TrafficEngineering_Pcc_PcePeers_PcePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Remote PCE address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PceAddress interface{}

    // PCC Peer MD5 Password. The type is string with length: 1..25.
    Password interface{}

    // PCC Peer Keychain. The type is string with length: 1..32.
    Keychain interface{}

    // PCC Peer Enable. The type is interface{}.
    Enable interface{}

    // Precedence value of this PCE. The type is interface{} with range: 0..255.
    Precedence interface{}
}

func (pcePeer *Sr_TrafficEngineering_Pcc_PcePeers_PcePeer) GetEntityData() *types.CommonEntityData {
    pcePeer.EntityData.YFilter = pcePeer.YFilter
    pcePeer.EntityData.YangName = "pce-peer"
    pcePeer.EntityData.BundleName = "cisco_ios_xr"
    pcePeer.EntityData.ParentYangName = "pce-peers"
    pcePeer.EntityData.SegmentPath = "pce-peer" + types.AddKeyToken(pcePeer.PceAddress, "pce-address")
    pcePeer.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/pcc/pce-peers/" + pcePeer.EntityData.SegmentPath
    pcePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcePeer.EntityData.Children = types.NewOrderedMap()
    pcePeer.EntityData.Leafs = types.NewOrderedMap()
    pcePeer.EntityData.Leafs.Append("pce-address", types.YLeaf{"PceAddress", pcePeer.PceAddress})
    pcePeer.EntityData.Leafs.Append("password", types.YLeaf{"Password", pcePeer.Password})
    pcePeer.EntityData.Leafs.Append("keychain", types.YLeaf{"Keychain", pcePeer.Keychain})
    pcePeer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pcePeer.Enable})
    pcePeer.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", pcePeer.Precedence})

    pcePeer.EntityData.YListKeys = []string {"PceAddress"}

    return &(pcePeer.EntityData)
}

// Sr_TrafficEngineering_Pcc_PceAddresses
// PCE peer configuration
type Sr_TrafficEngineering_Pcc_PceAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer address. The type is slice of
    // Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress.
    PceAddress []*Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress
}

func (pceAddresses *Sr_TrafficEngineering_Pcc_PceAddresses) GetEntityData() *types.CommonEntityData {
    pceAddresses.EntityData.YFilter = pceAddresses.YFilter
    pceAddresses.EntityData.YangName = "pce-addresses"
    pceAddresses.EntityData.BundleName = "cisco_ios_xr"
    pceAddresses.EntityData.ParentYangName = "pcc"
    pceAddresses.EntityData.SegmentPath = "pce-addresses"
    pceAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/pcc/" + pceAddresses.EntityData.SegmentPath
    pceAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceAddresses.EntityData.Children = types.NewOrderedMap()
    pceAddresses.EntityData.Children.Append("pce-address", types.YChild{"PceAddress", nil})
    for i := range pceAddresses.PceAddress {
        pceAddresses.EntityData.Children.Append(types.GetSegmentPath(pceAddresses.PceAddress[i]), types.YChild{"PceAddress", pceAddresses.PceAddress[i]})
    }
    pceAddresses.EntityData.Leafs = types.NewOrderedMap()

    pceAddresses.EntityData.YListKeys = []string {}

    return &(pceAddresses.EntityData)
}

// Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress
// PCE peer address
type Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Remote PCE address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PceAddress interface{}

    // Precedence value of this PCE. The type is interface{} with range: 0..255.
    // This attribute is mandatory.
    Precedence interface{}
}

func (pceAddress *Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress) GetEntityData() *types.CommonEntityData {
    pceAddress.EntityData.YFilter = pceAddress.YFilter
    pceAddress.EntityData.YangName = "pce-address"
    pceAddress.EntityData.BundleName = "cisco_ios_xr"
    pceAddress.EntityData.ParentYangName = "pce-addresses"
    pceAddress.EntityData.SegmentPath = "pce-address" + types.AddKeyToken(pceAddress.PceAddress, "pce-address")
    pceAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/pcc/pce-addresses/" + pceAddress.EntityData.SegmentPath
    pceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceAddress.EntityData.Children = types.NewOrderedMap()
    pceAddress.EntityData.Leafs = types.NewOrderedMap()
    pceAddress.EntityData.Leafs.Append("pce-address", types.YLeaf{"PceAddress", pceAddress.PceAddress})
    pceAddress.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", pceAddress.Precedence})

    pceAddress.EntityData.YListKeys = []string {"PceAddress"}

    return &(pceAddress.EntityData)
}

// Sr_TrafficEngineering_AffinityMaps
// Affinity-map configuration
type Sr_TrafficEngineering_AffinityMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity-map entry. The type is slice of
    // Sr_TrafficEngineering_AffinityMaps_AffinityMap.
    AffinityMap []*Sr_TrafficEngineering_AffinityMaps_AffinityMap
}

func (affinityMaps *Sr_TrafficEngineering_AffinityMaps) GetEntityData() *types.CommonEntityData {
    affinityMaps.EntityData.YFilter = affinityMaps.YFilter
    affinityMaps.EntityData.YangName = "affinity-maps"
    affinityMaps.EntityData.BundleName = "cisco_ios_xr"
    affinityMaps.EntityData.ParentYangName = "traffic-engineering"
    affinityMaps.EntityData.SegmentPath = "affinity-maps"
    affinityMaps.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/" + affinityMaps.EntityData.SegmentPath
    affinityMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMaps.EntityData.Children = types.NewOrderedMap()
    affinityMaps.EntityData.Children.Append("affinity-map", types.YChild{"AffinityMap", nil})
    for i := range affinityMaps.AffinityMap {
        affinityMaps.EntityData.Children.Append(types.GetSegmentPath(affinityMaps.AffinityMap[i]), types.YChild{"AffinityMap", affinityMaps.AffinityMap[i]})
    }
    affinityMaps.EntityData.Leafs = types.NewOrderedMap()

    affinityMaps.EntityData.YListKeys = []string {}

    return &(affinityMaps.EntityData)
}

// Sr_TrafficEngineering_AffinityMaps_AffinityMap
// Affinity-map entry
type Sr_TrafficEngineering_AffinityMaps_AffinityMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Affinity-map bit-position. The type is string with
    // length: 1..32.
    AffinityName interface{}

    // Affinity-map bit-position. The type is interface{} with range: 0..31. This
    // attribute is mandatory.
    BitPosition interface{}
}

func (affinityMap *Sr_TrafficEngineering_AffinityMaps_AffinityMap) GetEntityData() *types.CommonEntityData {
    affinityMap.EntityData.YFilter = affinityMap.YFilter
    affinityMap.EntityData.YangName = "affinity-map"
    affinityMap.EntityData.BundleName = "cisco_ios_xr"
    affinityMap.EntityData.ParentYangName = "affinity-maps"
    affinityMap.EntityData.SegmentPath = "affinity-map" + types.AddKeyToken(affinityMap.AffinityName, "affinity-name")
    affinityMap.EntityData.AbsolutePath = "Cisco-IOS-XR-segment-routing-ms-cfg:sr/Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering/affinity-maps/" + affinityMap.EntityData.SegmentPath
    affinityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMap.EntityData.Children = types.NewOrderedMap()
    affinityMap.EntityData.Leafs = types.NewOrderedMap()
    affinityMap.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityMap.AffinityName})
    affinityMap.EntityData.Leafs.Append("bit-position", types.YLeaf{"BitPosition", affinityMap.BitPosition})

    affinityMap.EntityData.YListKeys = []string {"AffinityName"}

    return &(affinityMap.EntityData)
}

