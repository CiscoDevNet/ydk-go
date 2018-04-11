// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sr: Segment Routing
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    sr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sr.EntityData.Children = make(map[string]types.YChild)
    sr.EntityData.Children["local-block"] = types.YChild{"LocalBlock", &sr.LocalBlock}
    sr.EntityData.Children["mappings"] = types.YChild{"Mappings", &sr.Mappings}
    sr.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", &sr.AdjacencySid}
    sr.EntityData.Children["global-block"] = types.YChild{"GlobalBlock", &sr.GlobalBlock}
    sr.EntityData.Children["Cisco-IOS-XR-infra-xtc-agent-cfg:traffic-engineering"] = types.YChild{"TrafficEngineering", &sr.TrafficEngineering}
    sr.EntityData.Leafs = make(map[string]types.YLeaf)
    sr.EntityData.Leafs["enable"] = types.YLeaf{"Enable", sr.Enable}
    return &(sr.EntityData)
}

// Sr_LocalBlock
// Segment Routing Local Block of Labels
// This type is a presence type.
type Sr_LocalBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    localBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localBlock.EntityData.Children = make(map[string]types.YChild)
    localBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    localBlock.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", localBlock.LowerBound}
    localBlock.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", localBlock.UpperBound}
    return &(localBlock.EntityData)
}

// Sr_Mappings
// Mapping Server
type Sr_Mappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP prefix to SID mapping. The type is slice of Sr_Mappings_Mapping.
    Mapping []Sr_Mappings_Mapping
}

func (mappings *Sr_Mappings) GetEntityData() *types.CommonEntityData {
    mappings.EntityData.YFilter = mappings.YFilter
    mappings.EntityData.YangName = "mappings"
    mappings.EntityData.BundleName = "cisco_ios_xr"
    mappings.EntityData.ParentYangName = "sr"
    mappings.EntityData.SegmentPath = "mappings"
    mappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappings.EntityData.Children = make(map[string]types.YChild)
    mappings.EntityData.Children["mapping"] = types.YChild{"Mapping", nil}
    for i := range mappings.Mapping {
        mappings.EntityData.Children[types.GetSegmentPath(&mappings.Mapping[i])] = types.YChild{"Mapping", &mappings.Mapping[i]}
    }
    mappings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mappings.EntityData)
}

// Sr_Mappings_Mapping
// IP prefix to SID mapping
type Sr_Mappings_Mapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Af interface{}

    // This attribute is a key. IP prefix. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Ip interface{}

    // This attribute is a key. Mask. The type is interface{} with range:
    // -2147483648..2147483647.
    Mask interface{}

    // Start of SID index range. The type is interface{} with range: 0..1048575.
    SidStart interface{}

    // Range (number of SIDs). The type is interface{} with range:
    // -2147483648..2147483647.
    SidRange interface{}

    // Enable/Disable Attached flag. The type is SrmsMiFlag.
    FlagAttached interface{}
}

func (mapping *Sr_Mappings_Mapping) GetEntityData() *types.CommonEntityData {
    mapping.EntityData.YFilter = mapping.YFilter
    mapping.EntityData.YangName = "mapping"
    mapping.EntityData.BundleName = "cisco_ios_xr"
    mapping.EntityData.ParentYangName = "mappings"
    mapping.EntityData.SegmentPath = "mapping" + "[af='" + fmt.Sprintf("%v", mapping.Af) + "']" + "[ip='" + fmt.Sprintf("%v", mapping.Ip) + "']" + "[mask='" + fmt.Sprintf("%v", mapping.Mask) + "']"
    mapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mapping.EntityData.Children = make(map[string]types.YChild)
    mapping.EntityData.Leafs = make(map[string]types.YLeaf)
    mapping.EntityData.Leafs["af"] = types.YLeaf{"Af", mapping.Af}
    mapping.EntityData.Leafs["ip"] = types.YLeaf{"Ip", mapping.Ip}
    mapping.EntityData.Leafs["mask"] = types.YLeaf{"Mask", mapping.Mask}
    mapping.EntityData.Leafs["sid-start"] = types.YLeaf{"SidStart", mapping.SidStart}
    mapping.EntityData.Leafs["sid-range"] = types.YLeaf{"SidRange", mapping.SidRange}
    mapping.EntityData.Leafs["flag-attached"] = types.YLeaf{"FlagAttached", mapping.FlagAttached}
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
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &adjacencySid.Interfaces}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(adjacencySid.EntityData)
}

// Sr_AdjacencySid_Interfaces
// Segment Routing Adjacency SID Interface Table
type Sr_AdjacencySid_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface. The type is slice of
    // Sr_AdjacencySid_Interfaces_Interface_.
    Interface_ []Sr_AdjacencySid_Interfaces_Interface
}

func (interfaces *Sr_AdjacencySid_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "adjacency-sid"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface
// Segment Routing Adjacency SID Interface
type Sr_AdjacencySid_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Interface_ interface{}

    // Segment Routing Adjacency SID Interface Address Family Table.
    AddressFamilies Sr_AdjacencySid_Interfaces_Interface_AddressFamilies
}

func (self *Sr_AdjacencySid_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface='" + fmt.Sprintf("%v", self.Interface_) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["address-families"] = types.YChild{"AddressFamilies", &self.AddressFamilies}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", self.Interface_}
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
    AddressFamily []Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily
}

func (addressFamilies *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies) GetEntityData() *types.CommonEntityData {
    addressFamilies.EntityData.YFilter = addressFamilies.YFilter
    addressFamilies.EntityData.YangName = "address-families"
    addressFamilies.EntityData.BundleName = "cisco_ios_xr"
    addressFamilies.EntityData.ParentYangName = "interface"
    addressFamilies.EntityData.SegmentPath = "address-families"
    addressFamilies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilies.EntityData.Children = make(map[string]types.YChild)
    addressFamilies.EntityData.Children["address-family"] = types.YChild{"AddressFamily", nil}
    for i := range addressFamilies.AddressFamily {
        addressFamilies.EntityData.Children[types.GetSegmentPath(&addressFamilies.AddressFamily[i])] = types.YChild{"AddressFamily", &addressFamilies.AddressFamily[i]}
    }
    addressFamilies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilies.EntityData)
}

// Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily
// Segment Routing Adjacency SID Interface
// Address Family
type Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    addressFamily.EntityData.SegmentPath = "address-family" + "[address-family='" + fmt.Sprintf("%v", addressFamily.AddressFamily) + "']"
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = make(map[string]types.YChild)
    addressFamily.EntityData.Children["next-hops"] = types.YChild{"NextHops", &addressFamily.NextHops}
    addressFamily.EntityData.Leafs = make(map[string]types.YLeaf)
    addressFamily.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", addressFamily.AddressFamily}
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
    NextHop []Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops_NextHop
}

func (nextHops *Sr_AdjacencySid_Interfaces_Interface_AddressFamilies_AddressFamily_NextHops) GetEntityData() *types.CommonEntityData {
    nextHops.EntityData.YFilter = nextHops.YFilter
    nextHops.EntityData.YangName = "next-hops"
    nextHops.EntityData.BundleName = "cisco_ios_xr"
    nextHops.EntityData.ParentYangName = "address-family"
    nextHops.EntityData.SegmentPath = "next-hops"
    nextHops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHops.EntityData.Children = make(map[string]types.YChild)
    nextHops.EntityData.Children["next-hop"] = types.YChild{"NextHop", nil}
    for i := range nextHops.NextHop {
        nextHops.EntityData.Children[types.GetSegmentPath(&nextHops.NextHop[i])] = types.YChild{"NextHop", &nextHops.NextHop[i]}
    }
    nextHops.EntityData.Leafs = make(map[string]types.YLeaf)
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
    nextHop.EntityData.SegmentPath = "next-hop" + "[ip-addr='" + fmt.Sprintf("%v", nextHop.IpAddr) + "']"
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nextHop.EntityData.Children = make(map[string]types.YChild)
    nextHop.EntityData.Children["l2-adjacency-sid"] = types.YChild{"L2AdjacencySid", &nextHop.L2AdjacencySid}
    nextHop.EntityData.Leafs = make(map[string]types.YLeaf)
    nextHop.EntityData.Leafs["ip-addr"] = types.YLeaf{"IpAddr", nextHop.IpAddr}
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
    l2AdjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2AdjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2AdjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2AdjacencySid.EntityData.Children = make(map[string]types.YChild)
    l2AdjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    l2AdjacencySid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", l2AdjacencySid.SidType}
    l2AdjacencySid.EntityData.Leafs["absolute-sid"] = types.YLeaf{"AbsoluteSid", l2AdjacencySid.AbsoluteSid}
    l2AdjacencySid.EntityData.Leafs["index-sid"] = types.YLeaf{"IndexSid", l2AdjacencySid.IndexSid}
    l2AdjacencySid.EntityData.Leafs["srlb"] = types.YLeaf{"Srlb", l2AdjacencySid.Srlb}
    return &(l2AdjacencySid.EntityData)
}

// Sr_GlobalBlock
// Global Block Segment Routing
// This type is a presence type.
type Sr_GlobalBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    globalBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalBlock.EntityData.Children = make(map[string]types.YChild)
    globalBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    globalBlock.EntityData.Leafs["lower-bound"] = types.YLeaf{"LowerBound", globalBlock.LowerBound}
    globalBlock.EntityData.Leafs["upper-bound"] = types.YLeaf{"UpperBound", globalBlock.UpperBound}
    return &(globalBlock.EntityData)
}

// Sr_TrafficEngineering
// Traffic Engineering configuration data
type Sr_TrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // On-demand color configuration.
    OnDemandColors Sr_TrafficEngineering_OnDemandColors

    // Segment-lists configuration.
    Segments Sr_TrafficEngineering_Segments

    // Logging configuration.
    Logging Sr_TrafficEngineering_Logging

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
    trafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficEngineering.EntityData.Children = make(map[string]types.YChild)
    trafficEngineering.EntityData.Children["on-demand-colors"] = types.YChild{"OnDemandColors", &trafficEngineering.OnDemandColors}
    trafficEngineering.EntityData.Children["segments"] = types.YChild{"Segments", &trafficEngineering.Segments}
    trafficEngineering.EntityData.Children["logging"] = types.YChild{"Logging", &trafficEngineering.Logging}
    trafficEngineering.EntityData.Children["binding-sid-rules"] = types.YChild{"BindingSidRules", &trafficEngineering.BindingSidRules}
    trafficEngineering.EntityData.Children["policies"] = types.YChild{"Policies", &trafficEngineering.Policies}
    trafficEngineering.EntityData.Children["srte-interfaces"] = types.YChild{"SrteInterfaces", &trafficEngineering.SrteInterfaces}
    trafficEngineering.EntityData.Children["pcc"] = types.YChild{"Pcc", &trafficEngineering.Pcc}
    trafficEngineering.EntityData.Children["affinity-maps"] = types.YChild{"AffinityMaps", &trafficEngineering.AffinityMaps}
    trafficEngineering.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficEngineering.EntityData.Leafs["enable"] = types.YLeaf{"Enable", trafficEngineering.Enable}
    return &(trafficEngineering.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors
// On-demand color configuration
type Sr_TrafficEngineering_OnDemandColors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // On-demand color configuration. The type is slice of
    // Sr_TrafficEngineering_OnDemandColors_OnDemandColor.
    OnDemandColor []Sr_TrafficEngineering_OnDemandColors_OnDemandColor
}

func (onDemandColors *Sr_TrafficEngineering_OnDemandColors) GetEntityData() *types.CommonEntityData {
    onDemandColors.EntityData.YFilter = onDemandColors.YFilter
    onDemandColors.EntityData.YangName = "on-demand-colors"
    onDemandColors.EntityData.BundleName = "cisco_ios_xr"
    onDemandColors.EntityData.ParentYangName = "traffic-engineering"
    onDemandColors.EntityData.SegmentPath = "on-demand-colors"
    onDemandColors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColors.EntityData.Children = make(map[string]types.YChild)
    onDemandColors.EntityData.Children["on-demand-color"] = types.YChild{"OnDemandColor", nil}
    for i := range onDemandColors.OnDemandColor {
        onDemandColors.EntityData.Children[types.GetSegmentPath(&onDemandColors.OnDemandColor[i])] = types.YChild{"OnDemandColor", &onDemandColors.OnDemandColor[i]}
    }
    onDemandColors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(onDemandColors.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor
// On-demand color configuration
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Color. The type is interface{} with range:
    // 1..4294967295.
    Color interface{}

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
    onDemandColor.EntityData.SegmentPath = "on-demand-color" + "[color='" + fmt.Sprintf("%v", onDemandColor.Color) + "']"
    onDemandColor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColor.EntityData.Children = make(map[string]types.YChild)
    onDemandColor.EntityData.Children["on-demand-color-dyn-mpls"] = types.YChild{"OnDemandColorDynMpls", &onDemandColor.OnDemandColorDynMpls}
    onDemandColor.EntityData.Leafs = make(map[string]types.YLeaf)
    onDemandColor.EntityData.Leafs["color"] = types.YLeaf{"Color", onDemandColor.Color}
    onDemandColor.EntityData.Leafs["enable"] = types.YLeaf{"Enable", onDemandColor.Enable}
    return &(onDemandColor.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls
// Dynamic MPLS path properties
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Dynamic MPLS path properties submode Enable. The type is interface{}.
    Enable interface{}

    // Metric type.
    OnDemandColorDynMplsMetric Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric

    // Use Path Computation Element.
    OnDemandColorDynMplsPce Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsPce

    // Disjoint path.
    DisjointPath Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath
}

func (onDemandColorDynMpls *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls) GetEntityData() *types.CommonEntityData {
    onDemandColorDynMpls.EntityData.YFilter = onDemandColorDynMpls.YFilter
    onDemandColorDynMpls.EntityData.YangName = "on-demand-color-dyn-mpls"
    onDemandColorDynMpls.EntityData.BundleName = "cisco_ios_xr"
    onDemandColorDynMpls.EntityData.ParentYangName = "on-demand-color"
    onDemandColorDynMpls.EntityData.SegmentPath = "on-demand-color-dyn-mpls"
    onDemandColorDynMpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMpls.EntityData.Children = make(map[string]types.YChild)
    onDemandColorDynMpls.EntityData.Children["on-demand-color-dyn-mpls-metric"] = types.YChild{"OnDemandColorDynMplsMetric", &onDemandColorDynMpls.OnDemandColorDynMplsMetric}
    onDemandColorDynMpls.EntityData.Children["on-demand-color-dyn-mpls-pce"] = types.YChild{"OnDemandColorDynMplsPce", &onDemandColorDynMpls.OnDemandColorDynMplsPce}
    onDemandColorDynMpls.EntityData.Children["disjoint-path"] = types.YChild{"DisjointPath", &onDemandColorDynMpls.DisjointPath}
    onDemandColorDynMpls.EntityData.Leafs = make(map[string]types.YLeaf)
    onDemandColorDynMpls.EntityData.Leafs["enable"] = types.YLeaf{"Enable", onDemandColorDynMpls.Enable}
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
}

func (onDemandColorDynMplsMetric *Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_OnDemandColorDynMplsMetric) GetEntityData() *types.CommonEntityData {
    onDemandColorDynMplsMetric.EntityData.YFilter = onDemandColorDynMplsMetric.YFilter
    onDemandColorDynMplsMetric.EntityData.YangName = "on-demand-color-dyn-mpls-metric"
    onDemandColorDynMplsMetric.EntityData.BundleName = "cisco_ios_xr"
    onDemandColorDynMplsMetric.EntityData.ParentYangName = "on-demand-color-dyn-mpls"
    onDemandColorDynMplsMetric.EntityData.SegmentPath = "on-demand-color-dyn-mpls-metric"
    onDemandColorDynMplsMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMplsMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMplsMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMplsMetric.EntityData.Children = make(map[string]types.YChild)
    onDemandColorDynMplsMetric.EntityData.Leafs = make(map[string]types.YLeaf)
    onDemandColorDynMplsMetric.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", onDemandColorDynMplsMetric.MetricType}
    onDemandColorDynMplsMetric.EntityData.Leafs["enable"] = types.YLeaf{"Enable", onDemandColorDynMplsMetric.Enable}
    return &(onDemandColorDynMplsMetric.EntityData)
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
    onDemandColorDynMplsPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColorDynMplsPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColorDynMplsPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColorDynMplsPce.EntityData.Children = make(map[string]types.YChild)
    onDemandColorDynMplsPce.EntityData.Leafs = make(map[string]types.YLeaf)
    onDemandColorDynMplsPce.EntityData.Leafs["enable"] = types.YLeaf{"Enable", onDemandColorDynMplsPce.Enable}
    return &(onDemandColorDynMplsPce.EntityData)
}

// Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath
// Disjoint path
// This type is a presence type.
type Sr_TrafficEngineering_OnDemandColors_OnDemandColor_OnDemandColorDynMpls_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = make(map[string]types.YChild)
    disjointPath.EntityData.Leafs = make(map[string]types.YLeaf)
    disjointPath.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", disjointPath.GroupId}
    disjointPath.EntityData.Leafs["disjointness-type"] = types.YLeaf{"DisjointnessType", disjointPath.DisjointnessType}
    disjointPath.EntityData.Leafs["sub-id"] = types.YLeaf{"SubId", disjointPath.SubId}
    return &(disjointPath.EntityData)
}

// Sr_TrafficEngineering_Segments
// Segment-lists configuration
type Sr_TrafficEngineering_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment-list configuration. The type is slice of
    // Sr_TrafficEngineering_Segments_Segment.
    Segment []Sr_TrafficEngineering_Segments_Segment
}

func (segments *Sr_TrafficEngineering_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "traffic-engineering"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = make(map[string]types.YChild)
    segments.EntityData.Children["segment"] = types.YChild{"Segment", nil}
    for i := range segments.Segment {
        segments.EntityData.Children[types.GetSegmentPath(&segments.Segment[i])] = types.YChild{"Segment", &segments.Segment[i]}
    }
    segments.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(segments.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment
// Segment-list configuration
type Sr_TrafficEngineering_Segments_Segment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    PathName interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Segments/hops configuration for given Segment-list.
    Segments Sr_TrafficEngineering_Segments_Segment_Segments_
}

func (segment *Sr_TrafficEngineering_Segments_Segment) GetEntityData() *types.CommonEntityData {
    segment.EntityData.YFilter = segment.YFilter
    segment.EntityData.YangName = "segment"
    segment.EntityData.BundleName = "cisco_ios_xr"
    segment.EntityData.ParentYangName = "segments"
    segment.EntityData.SegmentPath = "segment" + "[path-name='" + fmt.Sprintf("%v", segment.PathName) + "']"
    segment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment.EntityData.Children = make(map[string]types.YChild)
    segment.EntityData.Children["segments"] = types.YChild{"Segments", &segment.Segments}
    segment.EntityData.Leafs = make(map[string]types.YLeaf)
    segment.EntityData.Leafs["path-name"] = types.YLeaf{"PathName", segment.PathName}
    segment.EntityData.Leafs["enable"] = types.YLeaf{"Enable", segment.Enable}
    return &(segment.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment_Segments_
// Segments/hops configuration for given
// Segment-list
type Sr_TrafficEngineering_Segments_Segment_Segments_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Segment/hop at the index. The type is slice of
    // Sr_TrafficEngineering_Segments_Segment_Segments__Segment.
    Segment []Sr_TrafficEngineering_Segments_Segment_Segments__Segment_
}

func (segments_ *Sr_TrafficEngineering_Segments_Segment_Segments_) GetEntityData() *types.CommonEntityData {
    segments_.EntityData.YFilter = segments_.YFilter
    segments_.EntityData.YangName = "segments"
    segments_.EntityData.BundleName = "cisco_ios_xr"
    segments_.EntityData.ParentYangName = "segment"
    segments_.EntityData.SegmentPath = "segments"
    segments_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments_.EntityData.Children = make(map[string]types.YChild)
    segments_.EntityData.Children["segment"] = types.YChild{"Segment", nil}
    for i := range segments_.Segment {
        segments_.EntityData.Children[types.GetSegmentPath(&segments_.Segment[i])] = types.YChild{"Segment", &segments_.Segment[i]}
    }
    segments_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(segments_.EntityData)
}

// Sr_TrafficEngineering_Segments_Segment_Segments__Segment_
// Configure Segment/hop at the index
type Sr_TrafficEngineering_Segments_Segment_Segments__Segment_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Segment index. The type is interface{} with range:
    // 1..65535.
    SegmentIndex interface{}

    // Segment/hop type. The type is XtcSegment.
    SegmentType interface{}

    // IPv4 Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // MPLS Label. The type is interface{} with range: 0..1048575.
    MplsLabel interface{}
}

func (segment_ *Sr_TrafficEngineering_Segments_Segment_Segments__Segment_) GetEntityData() *types.CommonEntityData {
    segment_.EntityData.YFilter = segment_.YFilter
    segment_.EntityData.YangName = "segment"
    segment_.EntityData.BundleName = "cisco_ios_xr"
    segment_.EntityData.ParentYangName = "segments"
    segment_.EntityData.SegmentPath = "segment" + "[segment-index='" + fmt.Sprintf("%v", segment_.SegmentIndex) + "']"
    segment_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segment_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segment_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segment_.EntityData.Children = make(map[string]types.YChild)
    segment_.EntityData.Leafs = make(map[string]types.YLeaf)
    segment_.EntityData.Leafs["segment-index"] = types.YLeaf{"SegmentIndex", segment_.SegmentIndex}
    segment_.EntityData.Leafs["segment-type"] = types.YLeaf{"SegmentType", segment_.SegmentType}
    segment_.EntityData.Leafs["address"] = types.YLeaf{"Address", segment_.Address}
    segment_.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", segment_.MplsLabel}
    return &(segment_.EntityData)
}

// Sr_TrafficEngineering_Logging
// Logging configuration
type Sr_TrafficEngineering_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable logging for policy status. The type is interface{}.
    PolicyStatus interface{}
}

func (logging *Sr_TrafficEngineering_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "traffic-engineering"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    logging.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", logging.PolicyStatus}
    return &(logging.EntityData)
}

// Sr_TrafficEngineering_BindingSidRules
// Binding sid rules
type Sr_TrafficEngineering_BindingSidRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding sid explicit options.
    Explicit Sr_TrafficEngineering_BindingSidRules_Explicit
}

func (bindingSidRules *Sr_TrafficEngineering_BindingSidRules) GetEntityData() *types.CommonEntityData {
    bindingSidRules.EntityData.YFilter = bindingSidRules.YFilter
    bindingSidRules.EntityData.YangName = "binding-sid-rules"
    bindingSidRules.EntityData.BundleName = "cisco_ios_xr"
    bindingSidRules.EntityData.ParentYangName = "traffic-engineering"
    bindingSidRules.EntityData.SegmentPath = "binding-sid-rules"
    bindingSidRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSidRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSidRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSidRules.EntityData.Children = make(map[string]types.YChild)
    bindingSidRules.EntityData.Children["explicit"] = types.YChild{"Explicit", &bindingSidRules.Explicit}
    bindingSidRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bindingSidRules.EntityData)
}

// Sr_TrafficEngineering_BindingSidRules_Explicit
// Binding sid explicit options
type Sr_TrafficEngineering_BindingSidRules_Explicit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding sid explicit rule. The type is XtcBindingSidexplicitRule.
    Rule interface{}
}

func (explicit *Sr_TrafficEngineering_BindingSidRules_Explicit) GetEntityData() *types.CommonEntityData {
    explicit.EntityData.YFilter = explicit.YFilter
    explicit.EntityData.YangName = "explicit"
    explicit.EntityData.BundleName = "cisco_ios_xr"
    explicit.EntityData.ParentYangName = "binding-sid-rules"
    explicit.EntityData.SegmentPath = "explicit"
    explicit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    explicit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    explicit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    explicit.EntityData.Children = make(map[string]types.YChild)
    explicit.EntityData.Leafs = make(map[string]types.YLeaf)
    explicit.EntityData.Leafs["rule"] = types.YLeaf{"Rule", explicit.Rule}
    return &(explicit.EntityData)
}

// Sr_TrafficEngineering_Policies
// Policy configuration
type Sr_TrafficEngineering_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy.
    Policy []Sr_TrafficEngineering_Policies_Policy
}

func (policies *Sr_TrafficEngineering_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "traffic-engineering"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = make(map[string]types.YChild)
    policies.EntityData.Children["policy"] = types.YChild{"Policy", nil}
    for i := range policies.Policy {
        policies.EntityData.Children[types.GetSegmentPath(&policies.Policy[i])] = types.YChild{"Policy", &policies.Policy[i]}
    }
    policies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policies.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy
// Policy configuration
type Sr_TrafficEngineering_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Policy name. The type is string with length:
    // 1..128.
    PolicyName interface{}

    // Administratively shutdown policy. The type is interface{}.
    Shutdown interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Binding Segment ID.
    BindingSid Sr_TrafficEngineering_Policies_Policy_BindingSid

    // Color and endpoint of a policyColor, EndPointType, Endpoint.
    PolicyColorEndpoint Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint

    // Policy candidate-paths configuration.
    CandidatePaths Sr_TrafficEngineering_Policies_Policy_CandidatePaths
}

func (policy *Sr_TrafficEngineering_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + "[policy-name='" + fmt.Sprintf("%v", policy.PolicyName) + "']"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = make(map[string]types.YChild)
    policy.EntityData.Children["binding-sid"] = types.YChild{"BindingSid", &policy.BindingSid}
    policy.EntityData.Children["policy-color-endpoint"] = types.YChild{"PolicyColorEndpoint", &policy.PolicyColorEndpoint}
    policy.EntityData.Children["candidate-paths"] = types.YChild{"CandidatePaths", &policy.CandidatePaths}
    policy.EntityData.Leafs = make(map[string]types.YLeaf)
    policy.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policy.PolicyName}
    policy.EntityData.Leafs["shutdown"] = types.YLeaf{"Shutdown", policy.Shutdown}
    policy.EntityData.Leafs["enable"] = types.YLeaf{"Enable", policy.Enable}
    return &(policy.EntityData)
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
    bindingSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSid.EntityData.Children = make(map[string]types.YChild)
    bindingSid.EntityData.Leafs = make(map[string]types.YLeaf)
    bindingSid.EntityData.Leafs["binding-sid-type"] = types.YLeaf{"BindingSidType", bindingSid.BindingSidType}
    bindingSid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", bindingSid.MplsLabel}
    return &(bindingSid.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint
// Color and endpoint of a policyColor,
// EndPointType, Endpoint
// This type is a presence type.
type Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Color. The type is interface{} with range: 1..4294967295. This attribute is
    // mandatory.
    Color interface{}

    // End point type. The type is XtcEndPoint. This attribute is mandatory.
    EndPointType interface{}

    // End point address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    EndPointAddress interface{}
}

func (policyColorEndpoint *Sr_TrafficEngineering_Policies_Policy_PolicyColorEndpoint) GetEntityData() *types.CommonEntityData {
    policyColorEndpoint.EntityData.YFilter = policyColorEndpoint.YFilter
    policyColorEndpoint.EntityData.YangName = "policy-color-endpoint"
    policyColorEndpoint.EntityData.BundleName = "cisco_ios_xr"
    policyColorEndpoint.EntityData.ParentYangName = "policy"
    policyColorEndpoint.EntityData.SegmentPath = "policy-color-endpoint"
    policyColorEndpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyColorEndpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyColorEndpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyColorEndpoint.EntityData.Children = make(map[string]types.YChild)
    policyColorEndpoint.EntityData.Leafs = make(map[string]types.YLeaf)
    policyColorEndpoint.EntityData.Leafs["color"] = types.YLeaf{"Color", policyColorEndpoint.Color}
    policyColorEndpoint.EntityData.Leafs["end-point-type"] = types.YLeaf{"EndPointType", policyColorEndpoint.EndPointType}
    policyColorEndpoint.EntityData.Leafs["end-point-address"] = types.YLeaf{"EndPointAddress", policyColorEndpoint.EndPointAddress}
    return &(policyColorEndpoint.EntityData)
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
    candidatePaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidatePaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidatePaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidatePaths.EntityData.Children = make(map[string]types.YChild)
    candidatePaths.EntityData.Children["preferences"] = types.YChild{"Preferences", &candidatePaths.Preferences}
    candidatePaths.EntityData.Leafs = make(map[string]types.YLeaf)
    candidatePaths.EntityData.Leafs["enable"] = types.YLeaf{"Enable", candidatePaths.Enable}
    return &(candidatePaths.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences
// Policy path-option preference table
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy path-option preference entry. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference.
    Preference []Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference
}

func (preferences *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences) GetEntityData() *types.CommonEntityData {
    preferences.EntityData.YFilter = preferences.YFilter
    preferences.EntityData.YangName = "preferences"
    preferences.EntityData.BundleName = "cisco_ios_xr"
    preferences.EntityData.ParentYangName = "candidate-paths"
    preferences.EntityData.SegmentPath = "preferences"
    preferences.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preferences.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preferences.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preferences.EntityData.Children = make(map[string]types.YChild)
    preferences.EntityData.Children["preference"] = types.YChild{"Preference", nil}
    for i := range preferences.Preference {
        preferences.EntityData.Children[types.GetSegmentPath(&preferences.Preference[i])] = types.YChild{"Preference", &preferences.Preference[i]}
    }
    preferences.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(preferences.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference
// Policy path-option preference entry
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    preference.EntityData.SegmentPath = "preference" + "[path-index='" + fmt.Sprintf("%v", preference.PathIndex) + "']"
    preference.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    preference.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    preference.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    preference.EntityData.Children = make(map[string]types.YChild)
    preference.EntityData.Children["constraints"] = types.YChild{"Constraints", &preference.Constraints}
    preference.EntityData.Children["path-infos"] = types.YChild{"PathInfos", &preference.PathInfos}
    preference.EntityData.Leafs = make(map[string]types.YLeaf)
    preference.EntityData.Leafs["path-index"] = types.YLeaf{"PathIndex", preference.PathIndex}
    preference.EntityData.Leafs["enable"] = types.YLeaf{"Enable", preference.Enable}
    return &(preference.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints
// SR path computation and verification
// constraints
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}

    // SR path computation and verification affinity rules.
    AffinityRules Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules
}

func (constraints *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints) GetEntityData() *types.CommonEntityData {
    constraints.EntityData.YFilter = constraints.YFilter
    constraints.EntityData.YangName = "constraints"
    constraints.EntityData.BundleName = "cisco_ios_xr"
    constraints.EntityData.ParentYangName = "preference"
    constraints.EntityData.SegmentPath = "constraints"
    constraints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    constraints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    constraints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    constraints.EntityData.Children = make(map[string]types.YChild)
    constraints.EntityData.Children["affinity-rules"] = types.YChild{"AffinityRules", &constraints.AffinityRules}
    constraints.EntityData.Leafs = make(map[string]types.YLeaf)
    constraints.EntityData.Leafs["enable"] = types.YLeaf{"Enable", constraints.Enable}
    return &(constraints.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules
// SR path computation and verification
// affinity rules
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SR path computation and verification affinity rule. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule.
    AffinityRule []Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule
}

func (affinityRules *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules) GetEntityData() *types.CommonEntityData {
    affinityRules.EntityData.YFilter = affinityRules.YFilter
    affinityRules.EntityData.YangName = "affinity-rules"
    affinityRules.EntityData.BundleName = "cisco_ios_xr"
    affinityRules.EntityData.ParentYangName = "constraints"
    affinityRules.EntityData.SegmentPath = "affinity-rules"
    affinityRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRules.EntityData.Children = make(map[string]types.YChild)
    affinityRules.EntityData.Children["affinity-rule"] = types.YChild{"AffinityRule", nil}
    for i := range affinityRules.AffinityRule {
        affinityRules.EntityData.Children[types.GetSegmentPath(&affinityRules.AffinityRule[i])] = types.YChild{"AffinityRule", &affinityRules.AffinityRule[i]}
    }
    affinityRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(affinityRules.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule
// SR path computation and verification
// affinity rule
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path-option preference. The type is
    // XtcAffinityRule.
    Rule interface{}

    // This attribute is a key. The color. The type is string with length: 1..128.
    Color interface{}
}

func (affinityRule *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_Constraints_AffinityRules_AffinityRule) GetEntityData() *types.CommonEntityData {
    affinityRule.EntityData.YFilter = affinityRule.YFilter
    affinityRule.EntityData.YangName = "affinity-rule"
    affinityRule.EntityData.BundleName = "cisco_ios_xr"
    affinityRule.EntityData.ParentYangName = "affinity-rules"
    affinityRule.EntityData.SegmentPath = "affinity-rule" + "[rule='" + fmt.Sprintf("%v", affinityRule.Rule) + "']" + "[color='" + fmt.Sprintf("%v", affinityRule.Color) + "']"
    affinityRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityRule.EntityData.Children = make(map[string]types.YChild)
    affinityRule.EntityData.Leafs = make(map[string]types.YLeaf)
    affinityRule.EntityData.Leafs["rule"] = types.YLeaf{"Rule", affinityRule.Rule}
    affinityRule.EntityData.Leafs["color"] = types.YLeaf{"Color", affinityRule.Color}
    return &(affinityRule.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos
// Policy path-option preference
// configuration
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy configuration. The type is slice of
    // Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo.
    PathInfo []Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
}

func (pathInfos *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos) GetEntityData() *types.CommonEntityData {
    pathInfos.EntityData.YFilter = pathInfos.YFilter
    pathInfos.EntityData.YangName = "path-infos"
    pathInfos.EntityData.BundleName = "cisco_ios_xr"
    pathInfos.EntityData.ParentYangName = "preference"
    pathInfos.EntityData.SegmentPath = "path-infos"
    pathInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfos.EntityData.Children = make(map[string]types.YChild)
    pathInfos.EntityData.Children["path-info"] = types.YChild{"PathInfo", nil}
    for i := range pathInfos.PathInfo {
        pathInfos.EntityData.Children[types.GetSegmentPath(&pathInfos.PathInfo[i])] = types.YChild{"PathInfo", &pathInfos.PathInfo[i]}
    }
    pathInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pathInfos.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo
// Policy configuration
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path-option type. The type is XtcPath.
    Type_ interface{}

    // This attribute is a key. Type of dynamic path to be computed. The type is
    // XtcPathHop.
    HopType interface{}

    // This attribute is a key. Segment-list name. The type is string with length:
    // 1..128.
    SegmentListName interface{}

    // Path-option weight. The type is interface{} with range:
    // -2147483648..2147483647.
    Weight interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // PCE configuration for the candidate-path . Valid only for dynamic
    // candidate-paths .
    Pce Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pce

    // Metric configuration, valid only for dynamic path-options.
    Metric Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric
}

func (pathInfo *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "path-infos"
    pathInfo.EntityData.SegmentPath = "path-info" + "[type='" + fmt.Sprintf("%v", pathInfo.Type_) + "']" + "[hop-type='" + fmt.Sprintf("%v", pathInfo.HopType) + "']" + "[segment-list-name='" + fmt.Sprintf("%v", pathInfo.SegmentListName) + "']"
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = make(map[string]types.YChild)
    pathInfo.EntityData.Children["pce"] = types.YChild{"Pce", &pathInfo.Pce}
    pathInfo.EntityData.Children["metric"] = types.YChild{"Metric", &pathInfo.Metric}
    pathInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    pathInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", pathInfo.Type_}
    pathInfo.EntityData.Leafs["hop-type"] = types.YLeaf{"HopType", pathInfo.HopType}
    pathInfo.EntityData.Leafs["segment-list-name"] = types.YLeaf{"SegmentListName", pathInfo.SegmentListName}
    pathInfo.EntityData.Leafs["weight"] = types.YLeaf{"Weight", pathInfo.Weight}
    pathInfo.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pathInfo.Enable}
    return &(pathInfo.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pce
// PCE configuration for the candidate-path
// . Valid only for dynamic candidate-paths
// .
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True only. The type is interface{}.
    Enable interface{}
}

func (pce *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Pce) GetEntityData() *types.CommonEntityData {
    pce.EntityData.YFilter = pce.YFilter
    pce.EntityData.YangName = "pce"
    pce.EntityData.BundleName = "cisco_ios_xr"
    pce.EntityData.ParentYangName = "path-info"
    pce.EntityData.SegmentPath = "pce"
    pce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pce.EntityData.Children = make(map[string]types.YChild)
    pce.EntityData.Leafs = make(map[string]types.YLeaf)
    pce.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pce.Enable}
    return &(pce.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric
// Metric configuration, valid only for
// dynamic path-options
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of SIDs. The type is interface{} with range: 0..255.
    SidLimit interface{}

    // Metric type. The type is XtcMetric.
    MetricType interface{}

    // True only. The type is interface{}.
    Enable interface{}

    // Metric type.
    Margin Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_Margin
}

func (metric *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "path-info"
    metric.EntityData.SegmentPath = "metric"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = make(map[string]types.YChild)
    metric.EntityData.Children["margin"] = types.YChild{"Margin", &metric.Margin}
    metric.EntityData.Leafs = make(map[string]types.YLeaf)
    metric.EntityData.Leafs["sid-limit"] = types.YLeaf{"SidLimit", metric.SidLimit}
    metric.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", metric.MetricType}
    metric.EntityData.Leafs["enable"] = types.YLeaf{"Enable", metric.Enable}
    return &(metric.EntityData)
}

// Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_Margin
// Metric type
type Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_Margin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Metric value type. The type is XtcMetricValue.
    ValueType interface{}

    // Absolute metric value. The type is interface{} with range: 0..2147483647.
    AbsoluteValue interface{}

    // Relative metric value. The type is interface{} with range: 0..100.
    RelativeValue interface{}
}

func (margin *Sr_TrafficEngineering_Policies_Policy_CandidatePaths_Preferences_Preference_PathInfos_PathInfo_Metric_Margin) GetEntityData() *types.CommonEntityData {
    margin.EntityData.YFilter = margin.YFilter
    margin.EntityData.YangName = "margin"
    margin.EntityData.BundleName = "cisco_ios_xr"
    margin.EntityData.ParentYangName = "metric"
    margin.EntityData.SegmentPath = "margin"
    margin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    margin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    margin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    margin.EntityData.Children = make(map[string]types.YChild)
    margin.EntityData.Leafs = make(map[string]types.YLeaf)
    margin.EntityData.Leafs["value-type"] = types.YLeaf{"ValueType", margin.ValueType}
    margin.EntityData.Leafs["absolute-value"] = types.YLeaf{"AbsoluteValue", margin.AbsoluteValue}
    margin.EntityData.Leafs["relative-value"] = types.YLeaf{"RelativeValue", margin.RelativeValue}
    return &(margin.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces
// SR-TE interfaces
type Sr_TrafficEngineering_SrteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SR-TE interface. The type is slice of
    // Sr_TrafficEngineering_SrteInterfaces_SrteInterface.
    SrteInterface []Sr_TrafficEngineering_SrteInterfaces_SrteInterface
}

func (srteInterfaces *Sr_TrafficEngineering_SrteInterfaces) GetEntityData() *types.CommonEntityData {
    srteInterfaces.EntityData.YFilter = srteInterfaces.YFilter
    srteInterfaces.EntityData.YangName = "srte-interfaces"
    srteInterfaces.EntityData.BundleName = "cisco_ios_xr"
    srteInterfaces.EntityData.ParentYangName = "traffic-engineering"
    srteInterfaces.EntityData.SegmentPath = "srte-interfaces"
    srteInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srteInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srteInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srteInterfaces.EntityData.Children = make(map[string]types.YChild)
    srteInterfaces.EntityData.Children["srte-interface"] = types.YChild{"SrteInterface", nil}
    for i := range srteInterfaces.SrteInterface {
        srteInterfaces.EntityData.Children[types.GetSegmentPath(&srteInterfaces.SrteInterface[i])] = types.YChild{"SrteInterface", &srteInterfaces.SrteInterface[i]}
    }
    srteInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(srteInterfaces.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface
// SR-TE interface
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SR-TE Interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    SrteInterfaceName interface{}

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
    srteInterface.EntityData.SegmentPath = "srte-interface" + "[srte-interface-name='" + fmt.Sprintf("%v", srteInterface.SrteInterfaceName) + "']"
    srteInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srteInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srteInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srteInterface.EntityData.Children = make(map[string]types.YChild)
    srteInterface.EntityData.Children["interface-affinities"] = types.YChild{"InterfaceAffinities", &srteInterface.InterfaceAffinities}
    srteInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    srteInterface.EntityData.Leafs["srte-interface-name"] = types.YLeaf{"SrteInterfaceName", srteInterface.SrteInterfaceName}
    srteInterface.EntityData.Leafs["enable"] = types.YLeaf{"Enable", srteInterface.Enable}
    return &(srteInterface.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities
// Set user defined interface attribute flags
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set user defined interface attribute flags. The type is slice of
    // Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity.
    InterfaceAffinity []Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity
}

func (interfaceAffinities *Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities) GetEntityData() *types.CommonEntityData {
    interfaceAffinities.EntityData.YFilter = interfaceAffinities.YFilter
    interfaceAffinities.EntityData.YangName = "interface-affinities"
    interfaceAffinities.EntityData.BundleName = "cisco_ios_xr"
    interfaceAffinities.EntityData.ParentYangName = "srte-interface"
    interfaceAffinities.EntityData.SegmentPath = "interface-affinities"
    interfaceAffinities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAffinities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAffinities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAffinities.EntityData.Children = make(map[string]types.YChild)
    interfaceAffinities.EntityData.Children["interface-affinity"] = types.YChild{"InterfaceAffinity", nil}
    for i := range interfaceAffinities.InterfaceAffinity {
        interfaceAffinities.EntityData.Children[types.GetSegmentPath(&interfaceAffinities.InterfaceAffinity[i])] = types.YChild{"InterfaceAffinity", &interfaceAffinities.InterfaceAffinity[i]}
    }
    interfaceAffinities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceAffinities.EntityData)
}

// Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity
// Set user defined interface attribute flags
type Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface affinity colors. The type is string with
    // length: 1..32.
    Color interface{}
}

func (interfaceAffinity *Sr_TrafficEngineering_SrteInterfaces_SrteInterface_InterfaceAffinities_InterfaceAffinity) GetEntityData() *types.CommonEntityData {
    interfaceAffinity.EntityData.YFilter = interfaceAffinity.YFilter
    interfaceAffinity.EntityData.YangName = "interface-affinity"
    interfaceAffinity.EntityData.BundleName = "cisco_ios_xr"
    interfaceAffinity.EntityData.ParentYangName = "interface-affinities"
    interfaceAffinity.EntityData.SegmentPath = "interface-affinity" + "[color='" + fmt.Sprintf("%v", interfaceAffinity.Color) + "']"
    interfaceAffinity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceAffinity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceAffinity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceAffinity.EntityData.Children = make(map[string]types.YChild)
    interfaceAffinity.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceAffinity.EntityData.Leafs["color"] = types.YLeaf{"Color", interfaceAffinity.Color}
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

    // Report all local SR policies to connected PCEP peers. The type is
    // interface{}.
    ReportAll interface{}

    // Maximum time between two consecutive PCEP messages sent by this node. The
    // type is interface{} with range: 0..255.
    KeepaliveTimerInterval interface{}

    // Local source IP address to use on PCEP sessions. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // PCC Enable. The type is interface{}.
    Enable interface{}

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
    pcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcc.EntityData.Children = make(map[string]types.YChild)
    pcc.EntityData.Children["pce-peers"] = types.YChild{"PcePeers", &pcc.PcePeers}
    pcc.EntityData.Children["pce-addresses"] = types.YChild{"PceAddresses", &pcc.PceAddresses}
    pcc.EntityData.Leafs = make(map[string]types.YLeaf)
    pcc.EntityData.Leafs["dead-timer-interval"] = types.YLeaf{"DeadTimerInterval", pcc.DeadTimerInterval}
    pcc.EntityData.Leafs["report-all"] = types.YLeaf{"ReportAll", pcc.ReportAll}
    pcc.EntityData.Leafs["keepalive-timer-interval"] = types.YLeaf{"KeepaliveTimerInterval", pcc.KeepaliveTimerInterval}
    pcc.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", pcc.SourceAddress}
    pcc.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pcc.Enable}
    pcc.EntityData.Leafs["delegation-timeout"] = types.YLeaf{"DelegationTimeout", pcc.DelegationTimeout}
    return &(pcc.EntityData)
}

// Sr_TrafficEngineering_Pcc_PcePeers
// PCE peer configuration
type Sr_TrafficEngineering_Pcc_PcePeers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer. The type is slice of Sr_TrafficEngineering_Pcc_PcePeers_PcePeer.
    PcePeer []Sr_TrafficEngineering_Pcc_PcePeers_PcePeer
}

func (pcePeers *Sr_TrafficEngineering_Pcc_PcePeers) GetEntityData() *types.CommonEntityData {
    pcePeers.EntityData.YFilter = pcePeers.YFilter
    pcePeers.EntityData.YangName = "pce-peers"
    pcePeers.EntityData.BundleName = "cisco_ios_xr"
    pcePeers.EntityData.ParentYangName = "pcc"
    pcePeers.EntityData.SegmentPath = "pce-peers"
    pcePeers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcePeers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcePeers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcePeers.EntityData.Children = make(map[string]types.YChild)
    pcePeers.EntityData.Children["pce-peer"] = types.YChild{"PcePeer", nil}
    for i := range pcePeers.PcePeer {
        pcePeers.EntityData.Children[types.GetSegmentPath(&pcePeers.PcePeer[i])] = types.YChild{"PcePeer", &pcePeers.PcePeer[i]}
    }
    pcePeers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcePeers.EntityData)
}

// Sr_TrafficEngineering_Pcc_PcePeers_PcePeer
// PCE peer
type Sr_TrafficEngineering_Pcc_PcePeers_PcePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Remote PCE address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PceAddress interface{}

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
    pcePeer.EntityData.SegmentPath = "pce-peer" + "[pce-address='" + fmt.Sprintf("%v", pcePeer.PceAddress) + "']"
    pcePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcePeer.EntityData.Children = make(map[string]types.YChild)
    pcePeer.EntityData.Leafs = make(map[string]types.YLeaf)
    pcePeer.EntityData.Leafs["pce-address"] = types.YLeaf{"PceAddress", pcePeer.PceAddress}
    pcePeer.EntityData.Leafs["enable"] = types.YLeaf{"Enable", pcePeer.Enable}
    pcePeer.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", pcePeer.Precedence}
    return &(pcePeer.EntityData)
}

// Sr_TrafficEngineering_Pcc_PceAddresses
// PCE peer configuration
type Sr_TrafficEngineering_Pcc_PceAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer address. The type is slice of
    // Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress.
    PceAddress []Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress
}

func (pceAddresses *Sr_TrafficEngineering_Pcc_PceAddresses) GetEntityData() *types.CommonEntityData {
    pceAddresses.EntityData.YFilter = pceAddresses.YFilter
    pceAddresses.EntityData.YangName = "pce-addresses"
    pceAddresses.EntityData.BundleName = "cisco_ios_xr"
    pceAddresses.EntityData.ParentYangName = "pcc"
    pceAddresses.EntityData.SegmentPath = "pce-addresses"
    pceAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceAddresses.EntityData.Children = make(map[string]types.YChild)
    pceAddresses.EntityData.Children["pce-address"] = types.YChild{"PceAddress", nil}
    for i := range pceAddresses.PceAddress {
        pceAddresses.EntityData.Children[types.GetSegmentPath(&pceAddresses.PceAddress[i])] = types.YChild{"PceAddress", &pceAddresses.PceAddress[i]}
    }
    pceAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pceAddresses.EntityData)
}

// Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress
// PCE peer address
type Sr_TrafficEngineering_Pcc_PceAddresses_PceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Remote PCE address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    pceAddress.EntityData.SegmentPath = "pce-address" + "[pce-address='" + fmt.Sprintf("%v", pceAddress.PceAddress) + "']"
    pceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceAddress.EntityData.Children = make(map[string]types.YChild)
    pceAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    pceAddress.EntityData.Leafs["pce-address"] = types.YLeaf{"PceAddress", pceAddress.PceAddress}
    pceAddress.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", pceAddress.Precedence}
    return &(pceAddress.EntityData)
}

// Sr_TrafficEngineering_AffinityMaps
// Affinity-map configuration
type Sr_TrafficEngineering_AffinityMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity-map entry. The type is slice of
    // Sr_TrafficEngineering_AffinityMaps_AffinityMap.
    AffinityMap []Sr_TrafficEngineering_AffinityMaps_AffinityMap
}

func (affinityMaps *Sr_TrafficEngineering_AffinityMaps) GetEntityData() *types.CommonEntityData {
    affinityMaps.EntityData.YFilter = affinityMaps.YFilter
    affinityMaps.EntityData.YangName = "affinity-maps"
    affinityMaps.EntityData.BundleName = "cisco_ios_xr"
    affinityMaps.EntityData.ParentYangName = "traffic-engineering"
    affinityMaps.EntityData.SegmentPath = "affinity-maps"
    affinityMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMaps.EntityData.Children = make(map[string]types.YChild)
    affinityMaps.EntityData.Children["affinity-map"] = types.YChild{"AffinityMap", nil}
    for i := range affinityMaps.AffinityMap {
        affinityMaps.EntityData.Children[types.GetSegmentPath(&affinityMaps.AffinityMap[i])] = types.YChild{"AffinityMap", &affinityMaps.AffinityMap[i]}
    }
    affinityMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(affinityMaps.EntityData)
}

// Sr_TrafficEngineering_AffinityMaps_AffinityMap
// Affinity-map entry
type Sr_TrafficEngineering_AffinityMaps_AffinityMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Affinity-map bit-position. The type is string with
    // length: 1..32.
    Color interface{}

    // Affinity-map bit-position. The type is interface{} with range: 0..31. This
    // attribute is mandatory.
    BitPosition interface{}
}

func (affinityMap *Sr_TrafficEngineering_AffinityMaps_AffinityMap) GetEntityData() *types.CommonEntityData {
    affinityMap.EntityData.YFilter = affinityMap.YFilter
    affinityMap.EntityData.YangName = "affinity-map"
    affinityMap.EntityData.BundleName = "cisco_ios_xr"
    affinityMap.EntityData.ParentYangName = "affinity-maps"
    affinityMap.EntityData.SegmentPath = "affinity-map" + "[color='" + fmt.Sprintf("%v", affinityMap.Color) + "']"
    affinityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMap.EntityData.Children = make(map[string]types.YChild)
    affinityMap.EntityData.Leafs = make(map[string]types.YLeaf)
    affinityMap.EntityData.Leafs["color"] = types.YLeaf{"Color", affinityMap.Color}
    affinityMap.EntityData.Leafs["bit-position"] = types.YLeaf{"BitPosition", affinityMap.BitPosition}
    return &(affinityMap.EntityData)
}

