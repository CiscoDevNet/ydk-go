// This module contains a collection of YANG definitions
// for Cisco IOS-XR segment-routing-ms package operational data.
// 
// This module contains definitions
// for the following management objects:
//   srms: Segment Routing Mapping Server operational data
//   srlb: srlb
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package segment_routing_ms_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package segment_routing_ms_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-oper srms}", reflect.TypeOf(Srms{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-oper:srms", reflect.TypeOf(Srms{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-segment-routing-ms-oper srlb}", reflect.TypeOf(Srlb{}))
    ydk.RegisterEntity("Cisco-IOS-XR-segment-routing-ms-oper:srlb", reflect.TypeOf(Srlb{}))
}

// SrmsAf represents Srms af
type SrmsAf string

const (
    // None
    SrmsAf_none SrmsAf = "none"

    // IPv4
    SrmsAf_ipv4 SrmsAf = "ipv4"

    // IPv6
    SrmsAf_ipv6 SrmsAf = "ipv6"
)

// SrmsMiFlagEB represents Srms mi flag e b
type SrmsMiFlagEB string

const (
    // False
    SrmsMiFlagEB_false_ SrmsMiFlagEB = "false"

    // True
    SrmsMiFlagEB_true_ SrmsMiFlagEB = "true"
)

// SrmsMiAfEB represents Srms mi af e b
type SrmsMiAfEB string

const (
    // None
    SrmsMiAfEB_none SrmsMiAfEB = "none"

    // IPv4
    SrmsMiAfEB_ipv4 SrmsMiAfEB = "ipv4"

    // IPv6
    SrmsMiAfEB_ipv6 SrmsMiAfEB = "ipv6"
)

// SrmsMiSrcEB represents Srms mi src e b
type SrmsMiSrcEB string

const (
    // None
    SrmsMiSrcEB_none SrmsMiSrcEB = "none"

    // Local
    SrmsMiSrcEB_local SrmsMiSrcEB = "local"

    // Remote
    SrmsMiSrcEB_remote SrmsMiSrcEB = "remote"
)

// SidTypeEnum represents Sid type enum
type SidTypeEnum string

const (
    // Absolute SID
    SidTypeEnum_absolute SidTypeEnum = "absolute"

    // Index SID
    SidTypeEnum_index SidTypeEnum = "index"
)

// Srms
// Segment Routing Mapping Server operational data
type Srms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP prefix to SID mappings.
    Mapping Srms_Mapping

    // Adjacency SID.
    AdjacencySid Srms_AdjacencySid

    // Policy operational data.
    Policy Srms_Policy
}

func (srms *Srms) GetEntityData() *types.CommonEntityData {
    srms.EntityData.YFilter = srms.YFilter
    srms.EntityData.YangName = "srms"
    srms.EntityData.BundleName = "cisco_ios_xr"
    srms.EntityData.ParentYangName = "Cisco-IOS-XR-segment-routing-ms-oper"
    srms.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-ms-oper:srms"
    srms.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srms.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srms.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srms.EntityData.Children = types.NewOrderedMap()
    srms.EntityData.Children.Append("mapping", types.YChild{"Mapping", &srms.Mapping})
    srms.EntityData.Children.Append("adjacency-sid", types.YChild{"AdjacencySid", &srms.AdjacencySid})
    srms.EntityData.Children.Append("policy", types.YChild{"Policy", &srms.Policy})
    srms.EntityData.Leafs = types.NewOrderedMap()

    srms.EntityData.YListKeys = []string {}

    return &(srms.EntityData)
}

// Srms_Mapping
// IP prefix to SID mappings
type Srms_Mapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 prefix to SID mappings.
    MappingIpv4 Srms_Mapping_MappingIpv4

    // IPv6 prefix to SID mappings.
    MappingIpv6 Srms_Mapping_MappingIpv6
}

func (mapping *Srms_Mapping) GetEntityData() *types.CommonEntityData {
    mapping.EntityData.YFilter = mapping.YFilter
    mapping.EntityData.YangName = "mapping"
    mapping.EntityData.BundleName = "cisco_ios_xr"
    mapping.EntityData.ParentYangName = "srms"
    mapping.EntityData.SegmentPath = "mapping"
    mapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mapping.EntityData.Children = types.NewOrderedMap()
    mapping.EntityData.Children.Append("mapping-ipv4", types.YChild{"MappingIpv4", &mapping.MappingIpv4})
    mapping.EntityData.Children.Append("mapping-ipv6", types.YChild{"MappingIpv6", &mapping.MappingIpv6})
    mapping.EntityData.Leafs = types.NewOrderedMap()

    mapping.EntityData.YListKeys = []string {}

    return &(mapping.EntityData)
}

// Srms_Mapping_MappingIpv4
// IPv4 prefix to SID mappings
type Srms_Mapping_MappingIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP prefix to SID mapping item. It's not possible to list all of the IP
    // prefix to SID mappings, as the set of valid prefixes could be very large.
    // Instead, SID map information must be retrieved individually for each prefix
    // of interest. The type is slice of Srms_Mapping_MappingIpv4_MappingMi.
    MappingMi []*Srms_Mapping_MappingIpv4_MappingMi
}

func (mappingIpv4 *Srms_Mapping_MappingIpv4) GetEntityData() *types.CommonEntityData {
    mappingIpv4.EntityData.YFilter = mappingIpv4.YFilter
    mappingIpv4.EntityData.YangName = "mapping-ipv4"
    mappingIpv4.EntityData.BundleName = "cisco_ios_xr"
    mappingIpv4.EntityData.ParentYangName = "mapping"
    mappingIpv4.EntityData.SegmentPath = "mapping-ipv4"
    mappingIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingIpv4.EntityData.Children = types.NewOrderedMap()
    mappingIpv4.EntityData.Children.Append("mapping-mi", types.YChild{"MappingMi", nil})
    for i := range mappingIpv4.MappingMi {
        mappingIpv4.EntityData.Children.Append(types.GetSegmentPath(mappingIpv4.MappingMi[i]), types.YChild{"MappingMi", mappingIpv4.MappingMi[i]})
    }
    mappingIpv4.EntityData.Leafs = types.NewOrderedMap()

    mappingIpv4.EntityData.YListKeys = []string {}

    return &(mappingIpv4.EntityData)
}

// Srms_Mapping_MappingIpv4_MappingMi
// IP prefix to SID mapping item. It's not possible
// to list all of the IP prefix to SID mappings, as
// the set of valid prefixes could be very large.
// Instead, SID map information must be retrieved
// individually for each prefix of interest.
type Srms_Mapping_MappingIpv4_MappingMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ip interface{}

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Mapping_MappingIpv4_MappingMi_Addr
}

func (mappingMi *Srms_Mapping_MappingIpv4_MappingMi) GetEntityData() *types.CommonEntityData {
    mappingMi.EntityData.YFilter = mappingMi.YFilter
    mappingMi.EntityData.YangName = "mapping-mi"
    mappingMi.EntityData.BundleName = "cisco_ios_xr"
    mappingMi.EntityData.ParentYangName = "mapping-ipv4"
    mappingMi.EntityData.SegmentPath = "mapping-mi"
    mappingMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingMi.EntityData.Children = types.NewOrderedMap()
    mappingMi.EntityData.Children.Append("addr", types.YChild{"Addr", &mappingMi.Addr})
    mappingMi.EntityData.Leafs = types.NewOrderedMap()
    mappingMi.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", mappingMi.Ip})
    mappingMi.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", mappingMi.Prefix})
    mappingMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", mappingMi.Src})
    mappingMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", mappingMi.Router})
    mappingMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", mappingMi.Area})
    mappingMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", mappingMi.PrefixXr})
    mappingMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", mappingMi.SidStart})
    mappingMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", mappingMi.SidCount})
    mappingMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", mappingMi.LastPrefix})
    mappingMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", mappingMi.LastSidIndex})
    mappingMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", mappingMi.FlagAttached})

    mappingMi.EntityData.YListKeys = []string {}

    return &(mappingMi.EntityData)
}

// Srms_Mapping_MappingIpv4_MappingMi_Addr
// addr
type Srms_Mapping_MappingIpv4_MappingMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Mapping_MappingIpv4_MappingMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "mapping-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srms_Mapping_MappingIpv6
// IPv6 prefix to SID mappings
type Srms_Mapping_MappingIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP prefix to SID mapping item. It's not possible to list all of the IP
    // prefix to SID mappings, as the set of valid prefixes could be very large.
    // Instead, SID map information must be retrieved individually for each prefix
    // of interest. The type is slice of Srms_Mapping_MappingIpv6_MappingMi.
    MappingMi []*Srms_Mapping_MappingIpv6_MappingMi
}

func (mappingIpv6 *Srms_Mapping_MappingIpv6) GetEntityData() *types.CommonEntityData {
    mappingIpv6.EntityData.YFilter = mappingIpv6.YFilter
    mappingIpv6.EntityData.YangName = "mapping-ipv6"
    mappingIpv6.EntityData.BundleName = "cisco_ios_xr"
    mappingIpv6.EntityData.ParentYangName = "mapping"
    mappingIpv6.EntityData.SegmentPath = "mapping-ipv6"
    mappingIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingIpv6.EntityData.Children = types.NewOrderedMap()
    mappingIpv6.EntityData.Children.Append("mapping-mi", types.YChild{"MappingMi", nil})
    for i := range mappingIpv6.MappingMi {
        mappingIpv6.EntityData.Children.Append(types.GetSegmentPath(mappingIpv6.MappingMi[i]), types.YChild{"MappingMi", mappingIpv6.MappingMi[i]})
    }
    mappingIpv6.EntityData.Leafs = types.NewOrderedMap()

    mappingIpv6.EntityData.YListKeys = []string {}

    return &(mappingIpv6.EntityData)
}

// Srms_Mapping_MappingIpv6_MappingMi
// IP prefix to SID mapping item. It's not possible
// to list all of the IP prefix to SID mappings, as
// the set of valid prefixes could be very large.
// Instead, SID map information must be retrieved
// individually for each prefix of interest.
type Srms_Mapping_MappingIpv6_MappingMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Ip interface{}

    // Prefix. The type is interface{} with range: 0..4294967295.
    Prefix interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Mapping_MappingIpv6_MappingMi_Addr
}

func (mappingMi *Srms_Mapping_MappingIpv6_MappingMi) GetEntityData() *types.CommonEntityData {
    mappingMi.EntityData.YFilter = mappingMi.YFilter
    mappingMi.EntityData.YangName = "mapping-mi"
    mappingMi.EntityData.BundleName = "cisco_ios_xr"
    mappingMi.EntityData.ParentYangName = "mapping-ipv6"
    mappingMi.EntityData.SegmentPath = "mapping-mi"
    mappingMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingMi.EntityData.Children = types.NewOrderedMap()
    mappingMi.EntityData.Children.Append("addr", types.YChild{"Addr", &mappingMi.Addr})
    mappingMi.EntityData.Leafs = types.NewOrderedMap()
    mappingMi.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", mappingMi.Ip})
    mappingMi.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", mappingMi.Prefix})
    mappingMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", mappingMi.Src})
    mappingMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", mappingMi.Router})
    mappingMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", mappingMi.Area})
    mappingMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", mappingMi.PrefixXr})
    mappingMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", mappingMi.SidStart})
    mappingMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", mappingMi.SidCount})
    mappingMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", mappingMi.LastPrefix})
    mappingMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", mappingMi.LastSidIndex})
    mappingMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", mappingMi.FlagAttached})

    mappingMi.EntityData.YListKeys = []string {}

    return &(mappingMi.EntityData)
}

// Srms_Mapping_MappingIpv6_MappingMi_Addr
// addr
type Srms_Mapping_MappingIpv6_MappingMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Mapping_MappingIpv6_MappingMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "mapping-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srms_AdjacencySid
// Adjacency SID
type Srms_AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // L2 Adjacency Option.
    L2Adjacency Srms_AdjacencySid_L2Adjacency
}

func (adjacencySid *Srms_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "srms"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("l2-adjacency", types.YChild{"L2Adjacency", &adjacencySid.L2Adjacency})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()

    adjacencySid.EntityData.YListKeys = []string {}

    return &(adjacencySid.EntityData)
}

// Srms_AdjacencySid_L2Adjacency
// L2 Adjacency Option
type Srms_AdjacencySid_L2Adjacency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface directory.
    Interfaces Srms_AdjacencySid_L2Adjacency_Interfaces
}

func (l2Adjacency *Srms_AdjacencySid_L2Adjacency) GetEntityData() *types.CommonEntityData {
    l2Adjacency.EntityData.YFilter = l2Adjacency.YFilter
    l2Adjacency.EntityData.YangName = "l2-adjacency"
    l2Adjacency.EntityData.BundleName = "cisco_ios_xr"
    l2Adjacency.EntityData.ParentYangName = "adjacency-sid"
    l2Adjacency.EntityData.SegmentPath = "l2-adjacency"
    l2Adjacency.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Adjacency.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Adjacency.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Adjacency.EntityData.Children = types.NewOrderedMap()
    l2Adjacency.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &l2Adjacency.Interfaces})
    l2Adjacency.EntityData.Leafs = types.NewOrderedMap()

    l2Adjacency.EntityData.YListKeys = []string {}

    return &(l2Adjacency.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces
// Interface directory
type Srms_AdjacencySid_L2Adjacency_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Routing Adjacency SID Interface. The type is slice of
    // Srms_AdjacencySid_L2Adjacency_Interfaces_Interface.
    Interface []*Srms_AdjacencySid_L2Adjacency_Interfaces_Interface
}

func (interfaces *Srms_AdjacencySid_L2Adjacency_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "l2-adjacency"
    interfaces.EntityData.SegmentPath = "interfaces"
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

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface
// Segment Routing Adjacency SID Interface
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // address family container.
    AddressFamily Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily
}

func (self *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("address-family", types.YChild{"AddressFamily", &self.AddressFamily})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily
// address family container
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP version 4.
    Ipv4 Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4

    // IP version 6.
    Ipv6 Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6
}

func (addressFamily *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xr"
    addressFamily.EntityData.ParentYangName = "interface"
    addressFamily.EntityData.SegmentPath = "address-family"
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = types.NewOrderedMap()
    addressFamily.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &addressFamily.Ipv4})
    addressFamily.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &addressFamily.Ipv6})
    addressFamily.EntityData.Leafs = types.NewOrderedMap()

    addressFamily.EntityData.YListKeys = []string {}

    return &(addressFamily.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4
// IP version 4
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID record. The type is slice of
    // Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord.
    SidRecord []*Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord
}

func (ipv4 *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "address-family"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("sid-record", types.YChild{"SidRecord", nil})
    for i := range ipv4.SidRecord {
        ipv4.EntityData.Children.Append(types.GetSegmentPath(ipv4.SidRecord[i]), types.YChild{"SidRecord", ipv4.SidRecord[i]})
    }
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord
// SID record
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is SidTypeEnum.
    SidType interface{}

    // SID value. The type is interface{} with range: 0..4294967295.
    SidValue interface{}

    // Interface name. The type is string with length: 0..64.
    InterfaceName interface{}

    // SID Value. The type is interface{} with range: 0..4294967295.
    SidValueXr interface{}

    // SID type. The type is interface{} with range: 0..4294967295.
    SidTypeXr interface{}

    // Interface address family. The type is interface{} with range:
    // 0..4294967295.
    AddressFamily interface{}

    // Has nexthop. The type is bool.
    HasNexthop interface{}

    // Interface count. The type is interface{} with range:
    // -2147483648..2147483647.
    InterfaceCount interface{}

    // Interface delete count. The type is interface{} with range:
    // -2147483648..2147483647.
    InterfaceDeleteCount interface{}

    // Nexthop address.
    NexthopAddress Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord_NexthopAddress
}

func (sidRecord *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord) GetEntityData() *types.CommonEntityData {
    sidRecord.EntityData.YFilter = sidRecord.YFilter
    sidRecord.EntityData.YangName = "sid-record"
    sidRecord.EntityData.BundleName = "cisco_ios_xr"
    sidRecord.EntityData.ParentYangName = "ipv4"
    sidRecord.EntityData.SegmentPath = "sid-record"
    sidRecord.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidRecord.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidRecord.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidRecord.EntityData.Children = types.NewOrderedMap()
    sidRecord.EntityData.Children.Append("nexthop-address", types.YChild{"NexthopAddress", &sidRecord.NexthopAddress})
    sidRecord.EntityData.Leafs = types.NewOrderedMap()
    sidRecord.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", sidRecord.SidType})
    sidRecord.EntityData.Leafs.Append("sid-value", types.YLeaf{"SidValue", sidRecord.SidValue})
    sidRecord.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sidRecord.InterfaceName})
    sidRecord.EntityData.Leafs.Append("sid-value-xr", types.YLeaf{"SidValueXr", sidRecord.SidValueXr})
    sidRecord.EntityData.Leafs.Append("sid-type-xr", types.YLeaf{"SidTypeXr", sidRecord.SidTypeXr})
    sidRecord.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", sidRecord.AddressFamily})
    sidRecord.EntityData.Leafs.Append("has-nexthop", types.YLeaf{"HasNexthop", sidRecord.HasNexthop})
    sidRecord.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", sidRecord.InterfaceCount})
    sidRecord.EntityData.Leafs.Append("interface-delete-count", types.YLeaf{"InterfaceDeleteCount", sidRecord.InterfaceDeleteCount})

    sidRecord.EntityData.YListKeys = []string {}

    return &(sidRecord.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord_NexthopAddress
// Nexthop address
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord_NexthopAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsAf.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (nexthopAddress *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv4_SidRecord_NexthopAddress) GetEntityData() *types.CommonEntityData {
    nexthopAddress.EntityData.YFilter = nexthopAddress.YFilter
    nexthopAddress.EntityData.YangName = "nexthop-address"
    nexthopAddress.EntityData.BundleName = "cisco_ios_xr"
    nexthopAddress.EntityData.ParentYangName = "sid-record"
    nexthopAddress.EntityData.SegmentPath = "nexthop-address"
    nexthopAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthopAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthopAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthopAddress.EntityData.Children = types.NewOrderedMap()
    nexthopAddress.EntityData.Leafs = types.NewOrderedMap()
    nexthopAddress.EntityData.Leafs.Append("af", types.YLeaf{"Af", nexthopAddress.Af})
    nexthopAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nexthopAddress.Ipv4})
    nexthopAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nexthopAddress.Ipv6})

    nexthopAddress.EntityData.YListKeys = []string {}

    return &(nexthopAddress.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6
// IP version 6
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID record. The type is slice of
    // Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord.
    SidRecord []*Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord
}

func (ipv6 *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "address-family"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("sid-record", types.YChild{"SidRecord", nil})
    for i := range ipv6.SidRecord {
        ipv6.EntityData.Children.Append(types.GetSegmentPath(ipv6.SidRecord[i]), types.YChild{"SidRecord", ipv6.SidRecord[i]})
    }
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord
// SID record
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is SidTypeEnum.
    SidType interface{}

    // SID value. The type is interface{} with range: 0..4294967295.
    SidValue interface{}

    // Interface name. The type is string with length: 0..64.
    InterfaceName interface{}

    // SID Value. The type is interface{} with range: 0..4294967295.
    SidValueXr interface{}

    // SID type. The type is interface{} with range: 0..4294967295.
    SidTypeXr interface{}

    // Interface address family. The type is interface{} with range:
    // 0..4294967295.
    AddressFamily interface{}

    // Has nexthop. The type is bool.
    HasNexthop interface{}

    // Interface count. The type is interface{} with range:
    // -2147483648..2147483647.
    InterfaceCount interface{}

    // Interface delete count. The type is interface{} with range:
    // -2147483648..2147483647.
    InterfaceDeleteCount interface{}

    // Nexthop address.
    NexthopAddress Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord_NexthopAddress
}

func (sidRecord *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord) GetEntityData() *types.CommonEntityData {
    sidRecord.EntityData.YFilter = sidRecord.YFilter
    sidRecord.EntityData.YangName = "sid-record"
    sidRecord.EntityData.BundleName = "cisco_ios_xr"
    sidRecord.EntityData.ParentYangName = "ipv6"
    sidRecord.EntityData.SegmentPath = "sid-record"
    sidRecord.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidRecord.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidRecord.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidRecord.EntityData.Children = types.NewOrderedMap()
    sidRecord.EntityData.Children.Append("nexthop-address", types.YChild{"NexthopAddress", &sidRecord.NexthopAddress})
    sidRecord.EntityData.Leafs = types.NewOrderedMap()
    sidRecord.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", sidRecord.SidType})
    sidRecord.EntityData.Leafs.Append("sid-value", types.YLeaf{"SidValue", sidRecord.SidValue})
    sidRecord.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sidRecord.InterfaceName})
    sidRecord.EntityData.Leafs.Append("sid-value-xr", types.YLeaf{"SidValueXr", sidRecord.SidValueXr})
    sidRecord.EntityData.Leafs.Append("sid-type-xr", types.YLeaf{"SidTypeXr", sidRecord.SidTypeXr})
    sidRecord.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", sidRecord.AddressFamily})
    sidRecord.EntityData.Leafs.Append("has-nexthop", types.YLeaf{"HasNexthop", sidRecord.HasNexthop})
    sidRecord.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", sidRecord.InterfaceCount})
    sidRecord.EntityData.Leafs.Append("interface-delete-count", types.YLeaf{"InterfaceDeleteCount", sidRecord.InterfaceDeleteCount})

    sidRecord.EntityData.YListKeys = []string {}

    return &(sidRecord.EntityData)
}

// Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord_NexthopAddress
// Nexthop address
type Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord_NexthopAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsAf.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (nexthopAddress *Srms_AdjacencySid_L2Adjacency_Interfaces_Interface_AddressFamily_Ipv6_SidRecord_NexthopAddress) GetEntityData() *types.CommonEntityData {
    nexthopAddress.EntityData.YFilter = nexthopAddress.YFilter
    nexthopAddress.EntityData.YangName = "nexthop-address"
    nexthopAddress.EntityData.BundleName = "cisco_ios_xr"
    nexthopAddress.EntityData.ParentYangName = "sid-record"
    nexthopAddress.EntityData.SegmentPath = "nexthop-address"
    nexthopAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthopAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthopAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthopAddress.EntityData.Children = types.NewOrderedMap()
    nexthopAddress.EntityData.Leafs = types.NewOrderedMap()
    nexthopAddress.EntityData.Leafs.Append("af", types.YLeaf{"Af", nexthopAddress.Af})
    nexthopAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", nexthopAddress.Ipv4})
    nexthopAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", nexthopAddress.Ipv6})

    nexthopAddress.EntityData.YListKeys = []string {}

    return &(nexthopAddress.EntityData)
}

// Srms_Policy
// Policy operational data
type Srms_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 policy operational data.
    PolicyIpv4 Srms_Policy_PolicyIpv4

    // IPv6 policy operational data.
    PolicyIpv6 Srms_Policy_PolicyIpv6
}

func (policy *Srms_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "srms"
    policy.EntityData.SegmentPath = "policy"
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Children.Append("policy-ipv4", types.YChild{"PolicyIpv4", &policy.PolicyIpv4})
    policy.EntityData.Children.Append("policy-ipv6", types.YChild{"PolicyIpv6", &policy.PolicyIpv6})
    policy.EntityData.Leafs = types.NewOrderedMap()

    policy.EntityData.YListKeys = []string {}

    return &(policy.EntityData)
}

// Srms_Policy_PolicyIpv4
// IPv4 policy operational data
type Srms_Policy_PolicyIpv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 backup policy operational data.
    PolicyIpv4Backup Srms_Policy_PolicyIpv4_PolicyIpv4Backup

    // IPv4 active policy operational data.
    PolicyIpv4Active Srms_Policy_PolicyIpv4_PolicyIpv4Active
}

func (policyIpv4 *Srms_Policy_PolicyIpv4) GetEntityData() *types.CommonEntityData {
    policyIpv4.EntityData.YFilter = policyIpv4.YFilter
    policyIpv4.EntityData.YangName = "policy-ipv4"
    policyIpv4.EntityData.BundleName = "cisco_ios_xr"
    policyIpv4.EntityData.ParentYangName = "policy"
    policyIpv4.EntityData.SegmentPath = "policy-ipv4"
    policyIpv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv4.EntityData.Children = types.NewOrderedMap()
    policyIpv4.EntityData.Children.Append("policy-ipv4-backup", types.YChild{"PolicyIpv4Backup", &policyIpv4.PolicyIpv4Backup})
    policyIpv4.EntityData.Children.Append("policy-ipv4-active", types.YChild{"PolicyIpv4Active", &policyIpv4.PolicyIpv4Active})
    policyIpv4.EntityData.Leafs = types.NewOrderedMap()

    policyIpv4.EntityData.YListKeys = []string {}

    return &(policyIpv4.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup
// IPv4 backup policy operational data
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi.
    PolicyMi []*Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi
}

func (policyIpv4Backup *Srms_Policy_PolicyIpv4_PolicyIpv4Backup) GetEntityData() *types.CommonEntityData {
    policyIpv4Backup.EntityData.YFilter = policyIpv4Backup.YFilter
    policyIpv4Backup.EntityData.YangName = "policy-ipv4-backup"
    policyIpv4Backup.EntityData.BundleName = "cisco_ios_xr"
    policyIpv4Backup.EntityData.ParentYangName = "policy-ipv4"
    policyIpv4Backup.EntityData.SegmentPath = "policy-ipv4-backup"
    policyIpv4Backup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv4Backup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv4Backup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv4Backup.EntityData.Children = types.NewOrderedMap()
    policyIpv4Backup.EntityData.Children.Append("policy-mi", types.YChild{"PolicyMi", nil})
    for i := range policyIpv4Backup.PolicyMi {
        policyIpv4Backup.EntityData.Children.Append(types.GetSegmentPath(policyIpv4Backup.PolicyMi[i]), types.YChild{"PolicyMi", policyIpv4Backup.PolicyMi[i]})
    }
    policyIpv4Backup.EntityData.Leafs = types.NewOrderedMap()

    policyIpv4Backup.EntityData.YListKeys = []string {}

    return &(policyIpv4Backup.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi) GetEntityData() *types.CommonEntityData {
    policyMi.EntityData.YFilter = policyMi.YFilter
    policyMi.EntityData.YangName = "policy-mi"
    policyMi.EntityData.BundleName = "cisco_ios_xr"
    policyMi.EntityData.ParentYangName = "policy-ipv4-backup"
    policyMi.EntityData.SegmentPath = "policy-mi" + types.AddKeyToken(policyMi.MiId, "mi-id")
    policyMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMi.EntityData.Children = types.NewOrderedMap()
    policyMi.EntityData.Children.Append("addr", types.YChild{"Addr", &policyMi.Addr})
    policyMi.EntityData.Leafs = types.NewOrderedMap()
    policyMi.EntityData.Leafs.Append("mi-id", types.YLeaf{"MiId", policyMi.MiId})
    policyMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", policyMi.Src})
    policyMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", policyMi.Router})
    policyMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", policyMi.Area})
    policyMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", policyMi.PrefixXr})
    policyMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", policyMi.SidStart})
    policyMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", policyMi.SidCount})
    policyMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", policyMi.LastPrefix})
    policyMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", policyMi.LastSidIndex})
    policyMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", policyMi.FlagAttached})

    policyMi.EntityData.YListKeys = []string {"MiId"}

    return &(policyMi.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Backup_PolicyMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "policy-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Active
// IPv4 active policy operational data
type Srms_Policy_PolicyIpv4_PolicyIpv4Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi.
    PolicyMi []*Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi
}

func (policyIpv4Active *Srms_Policy_PolicyIpv4_PolicyIpv4Active) GetEntityData() *types.CommonEntityData {
    policyIpv4Active.EntityData.YFilter = policyIpv4Active.YFilter
    policyIpv4Active.EntityData.YangName = "policy-ipv4-active"
    policyIpv4Active.EntityData.BundleName = "cisco_ios_xr"
    policyIpv4Active.EntityData.ParentYangName = "policy-ipv4"
    policyIpv4Active.EntityData.SegmentPath = "policy-ipv4-active"
    policyIpv4Active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv4Active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv4Active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv4Active.EntityData.Children = types.NewOrderedMap()
    policyIpv4Active.EntityData.Children.Append("policy-mi", types.YChild{"PolicyMi", nil})
    for i := range policyIpv4Active.PolicyMi {
        policyIpv4Active.EntityData.Children.Append(types.GetSegmentPath(policyIpv4Active.PolicyMi[i]), types.YChild{"PolicyMi", policyIpv4Active.PolicyMi[i]})
    }
    policyIpv4Active.EntityData.Leafs = types.NewOrderedMap()

    policyIpv4Active.EntityData.YListKeys = []string {}

    return &(policyIpv4Active.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi) GetEntityData() *types.CommonEntityData {
    policyMi.EntityData.YFilter = policyMi.YFilter
    policyMi.EntityData.YangName = "policy-mi"
    policyMi.EntityData.BundleName = "cisco_ios_xr"
    policyMi.EntityData.ParentYangName = "policy-ipv4-active"
    policyMi.EntityData.SegmentPath = "policy-mi" + types.AddKeyToken(policyMi.MiId, "mi-id")
    policyMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMi.EntityData.Children = types.NewOrderedMap()
    policyMi.EntityData.Children.Append("addr", types.YChild{"Addr", &policyMi.Addr})
    policyMi.EntityData.Leafs = types.NewOrderedMap()
    policyMi.EntityData.Leafs.Append("mi-id", types.YLeaf{"MiId", policyMi.MiId})
    policyMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", policyMi.Src})
    policyMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", policyMi.Router})
    policyMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", policyMi.Area})
    policyMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", policyMi.PrefixXr})
    policyMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", policyMi.SidStart})
    policyMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", policyMi.SidCount})
    policyMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", policyMi.LastPrefix})
    policyMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", policyMi.LastSidIndex})
    policyMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", policyMi.FlagAttached})

    policyMi.EntityData.YListKeys = []string {"MiId"}

    return &(policyMi.EntityData)
}

// Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv4_PolicyIpv4Active_PolicyMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "policy-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srms_Policy_PolicyIpv6
// IPv6 policy operational data
type Srms_Policy_PolicyIpv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 backup policy operational data.
    PolicyIpv6Backup Srms_Policy_PolicyIpv6_PolicyIpv6Backup

    // IPv6 active policy operational data.
    PolicyIpv6Active Srms_Policy_PolicyIpv6_PolicyIpv6Active
}

func (policyIpv6 *Srms_Policy_PolicyIpv6) GetEntityData() *types.CommonEntityData {
    policyIpv6.EntityData.YFilter = policyIpv6.YFilter
    policyIpv6.EntityData.YangName = "policy-ipv6"
    policyIpv6.EntityData.BundleName = "cisco_ios_xr"
    policyIpv6.EntityData.ParentYangName = "policy"
    policyIpv6.EntityData.SegmentPath = "policy-ipv6"
    policyIpv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv6.EntityData.Children = types.NewOrderedMap()
    policyIpv6.EntityData.Children.Append("policy-ipv6-backup", types.YChild{"PolicyIpv6Backup", &policyIpv6.PolicyIpv6Backup})
    policyIpv6.EntityData.Children.Append("policy-ipv6-active", types.YChild{"PolicyIpv6Active", &policyIpv6.PolicyIpv6Active})
    policyIpv6.EntityData.Leafs = types.NewOrderedMap()

    policyIpv6.EntityData.YListKeys = []string {}

    return &(policyIpv6.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup
// IPv6 backup policy operational data
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi.
    PolicyMi []*Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi
}

func (policyIpv6Backup *Srms_Policy_PolicyIpv6_PolicyIpv6Backup) GetEntityData() *types.CommonEntityData {
    policyIpv6Backup.EntityData.YFilter = policyIpv6Backup.YFilter
    policyIpv6Backup.EntityData.YangName = "policy-ipv6-backup"
    policyIpv6Backup.EntityData.BundleName = "cisco_ios_xr"
    policyIpv6Backup.EntityData.ParentYangName = "policy-ipv6"
    policyIpv6Backup.EntityData.SegmentPath = "policy-ipv6-backup"
    policyIpv6Backup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv6Backup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv6Backup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv6Backup.EntityData.Children = types.NewOrderedMap()
    policyIpv6Backup.EntityData.Children.Append("policy-mi", types.YChild{"PolicyMi", nil})
    for i := range policyIpv6Backup.PolicyMi {
        policyIpv6Backup.EntityData.Children.Append(types.GetSegmentPath(policyIpv6Backup.PolicyMi[i]), types.YChild{"PolicyMi", policyIpv6Backup.PolicyMi[i]})
    }
    policyIpv6Backup.EntityData.Leafs = types.NewOrderedMap()

    policyIpv6Backup.EntityData.YListKeys = []string {}

    return &(policyIpv6Backup.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi) GetEntityData() *types.CommonEntityData {
    policyMi.EntityData.YFilter = policyMi.YFilter
    policyMi.EntityData.YangName = "policy-mi"
    policyMi.EntityData.BundleName = "cisco_ios_xr"
    policyMi.EntityData.ParentYangName = "policy-ipv6-backup"
    policyMi.EntityData.SegmentPath = "policy-mi" + types.AddKeyToken(policyMi.MiId, "mi-id")
    policyMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMi.EntityData.Children = types.NewOrderedMap()
    policyMi.EntityData.Children.Append("addr", types.YChild{"Addr", &policyMi.Addr})
    policyMi.EntityData.Leafs = types.NewOrderedMap()
    policyMi.EntityData.Leafs.Append("mi-id", types.YLeaf{"MiId", policyMi.MiId})
    policyMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", policyMi.Src})
    policyMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", policyMi.Router})
    policyMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", policyMi.Area})
    policyMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", policyMi.PrefixXr})
    policyMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", policyMi.SidStart})
    policyMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", policyMi.SidCount})
    policyMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", policyMi.LastPrefix})
    policyMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", policyMi.LastSidIndex})
    policyMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", policyMi.FlagAttached})

    policyMi.EntityData.YListKeys = []string {"MiId"}

    return &(policyMi.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Backup_PolicyMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "policy-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Active
// IPv6 active policy operational data
type Srms_Policy_PolicyIpv6_PolicyIpv6Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mapping Item. The type is slice of
    // Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi.
    PolicyMi []*Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi
}

func (policyIpv6Active *Srms_Policy_PolicyIpv6_PolicyIpv6Active) GetEntityData() *types.CommonEntityData {
    policyIpv6Active.EntityData.YFilter = policyIpv6Active.YFilter
    policyIpv6Active.EntityData.YangName = "policy-ipv6-active"
    policyIpv6Active.EntityData.BundleName = "cisco_ios_xr"
    policyIpv6Active.EntityData.ParentYangName = "policy-ipv6"
    policyIpv6Active.EntityData.SegmentPath = "policy-ipv6-active"
    policyIpv6Active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIpv6Active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIpv6Active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIpv6Active.EntityData.Children = types.NewOrderedMap()
    policyIpv6Active.EntityData.Children.Append("policy-mi", types.YChild{"PolicyMi", nil})
    for i := range policyIpv6Active.PolicyMi {
        policyIpv6Active.EntityData.Children.Append(types.GetSegmentPath(policyIpv6Active.PolicyMi[i]), types.YChild{"PolicyMi", policyIpv6Active.PolicyMi[i]})
    }
    policyIpv6Active.EntityData.Leafs = types.NewOrderedMap()

    policyIpv6Active.EntityData.YListKeys = []string {}

    return &(policyIpv6Active.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi
// Mapping Item
type Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Mapping Item ID (0, 1, 2, ...). The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    MiId interface{}

    // src. The type is SrmsMiSrcEB.
    Src interface{}

    // Router ID. The type is string with length: 0..30.
    Router interface{}

    // Area (OSPF) or Level (ISIS). The type is string with length: 0..30.
    Area interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixXr interface{}

    // Starting SID. The type is interface{} with range: 0..4294967295.
    SidStart interface{}

    // SID range. The type is interface{} with range: 0..4294967295.
    SidCount interface{}

    // Last IP Prefix. The type is string with length: 0..50.
    LastPrefix interface{}

    // Last SID Index. The type is interface{} with range: 0..4294967295.
    LastSidIndex interface{}

    // Attached flag. The type is SrmsMiFlagEB.
    FlagAttached interface{}

    // addr.
    Addr Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr
}

func (policyMi *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi) GetEntityData() *types.CommonEntityData {
    policyMi.EntityData.YFilter = policyMi.YFilter
    policyMi.EntityData.YangName = "policy-mi"
    policyMi.EntityData.BundleName = "cisco_ios_xr"
    policyMi.EntityData.ParentYangName = "policy-ipv6-active"
    policyMi.EntityData.SegmentPath = "policy-mi" + types.AddKeyToken(policyMi.MiId, "mi-id")
    policyMi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyMi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyMi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyMi.EntityData.Children = types.NewOrderedMap()
    policyMi.EntityData.Children.Append("addr", types.YChild{"Addr", &policyMi.Addr})
    policyMi.EntityData.Leafs = types.NewOrderedMap()
    policyMi.EntityData.Leafs.Append("mi-id", types.YLeaf{"MiId", policyMi.MiId})
    policyMi.EntityData.Leafs.Append("src", types.YLeaf{"Src", policyMi.Src})
    policyMi.EntityData.Leafs.Append("router", types.YLeaf{"Router", policyMi.Router})
    policyMi.EntityData.Leafs.Append("area", types.YLeaf{"Area", policyMi.Area})
    policyMi.EntityData.Leafs.Append("prefix-xr", types.YLeaf{"PrefixXr", policyMi.PrefixXr})
    policyMi.EntityData.Leafs.Append("sid-start", types.YLeaf{"SidStart", policyMi.SidStart})
    policyMi.EntityData.Leafs.Append("sid-count", types.YLeaf{"SidCount", policyMi.SidCount})
    policyMi.EntityData.Leafs.Append("last-prefix", types.YLeaf{"LastPrefix", policyMi.LastPrefix})
    policyMi.EntityData.Leafs.Append("last-sid-index", types.YLeaf{"LastSidIndex", policyMi.LastSidIndex})
    policyMi.EntityData.Leafs.Append("flag-attached", types.YLeaf{"FlagAttached", policyMi.FlagAttached})

    policyMi.EntityData.YListKeys = []string {"MiId"}

    return &(policyMi.EntityData)
}

// Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr
// addr
type Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AF. The type is SrmsMiAfEB.
    Af interface{}

    // IPv4. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (addr *Srms_Policy_PolicyIpv6_PolicyIpv6Active_PolicyMi_Addr) GetEntityData() *types.CommonEntityData {
    addr.EntityData.YFilter = addr.YFilter
    addr.EntityData.YangName = "addr"
    addr.EntityData.BundleName = "cisco_ios_xr"
    addr.EntityData.ParentYangName = "policy-mi"
    addr.EntityData.SegmentPath = "addr"
    addr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addr.EntityData.Children = types.NewOrderedMap()
    addr.EntityData.Leafs = types.NewOrderedMap()
    addr.EntityData.Leafs.Append("af", types.YLeaf{"Af", addr.Af})
    addr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", addr.Ipv4})
    addr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", addr.Ipv6})

    addr.EntityData.YListKeys = []string {}

    return &(addr.EntityData)
}

// Srlb
// srlb
type Srlb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLB Inconsistencies.
    SrlbInconsistency Srlb_SrlbInconsistency
}

func (srlb *Srlb) GetEntityData() *types.CommonEntityData {
    srlb.EntityData.YFilter = srlb.YFilter
    srlb.EntityData.YangName = "srlb"
    srlb.EntityData.BundleName = "cisco_ios_xr"
    srlb.EntityData.ParentYangName = "Cisco-IOS-XR-segment-routing-ms-oper"
    srlb.EntityData.SegmentPath = "Cisco-IOS-XR-segment-routing-ms-oper:srlb"
    srlb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlb.EntityData.Children = types.NewOrderedMap()
    srlb.EntityData.Children.Append("srlb-inconsistency", types.YChild{"SrlbInconsistency", &srlb.SrlbInconsistency})
    srlb.EntityData.Leafs = types.NewOrderedMap()

    srlb.EntityData.YListKeys = []string {}

    return &(srlb.EntityData)
}

// Srlb_SrlbInconsistency
// SRLB Inconsistencies
type Srlb_SrlbInconsistency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start label of Segment Routing Local Block range. The type is interface{}
    // with range: 0..4294967295.
    StartSrlbRange interface{}

    // End label of Segment Routing Local Block range. The type is interface{}
    // with range: 0..4294967295.
    EndSrlbRange interface{}
}

func (srlbInconsistency *Srlb_SrlbInconsistency) GetEntityData() *types.CommonEntityData {
    srlbInconsistency.EntityData.YFilter = srlbInconsistency.YFilter
    srlbInconsistency.EntityData.YangName = "srlb-inconsistency"
    srlbInconsistency.EntityData.BundleName = "cisco_ios_xr"
    srlbInconsistency.EntityData.ParentYangName = "srlb"
    srlbInconsistency.EntityData.SegmentPath = "srlb-inconsistency"
    srlbInconsistency.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlbInconsistency.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlbInconsistency.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlbInconsistency.EntityData.Children = types.NewOrderedMap()
    srlbInconsistency.EntityData.Leafs = types.NewOrderedMap()
    srlbInconsistency.EntityData.Leafs.Append("start-srlb-range", types.YLeaf{"StartSrlbRange", srlbInconsistency.StartSrlbRange})
    srlbInconsistency.EntityData.Leafs.Append("end-srlb-range", types.YLeaf{"EndSrlbRange", srlbInconsistency.EndSrlbRange})

    srlbInconsistency.EntityData.YListKeys = []string {}

    return &(srlbInconsistency.EntityData)
}

