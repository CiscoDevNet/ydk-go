// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-autorp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   auto-rp: AutoRP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_autorp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_autorp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-autorp-oper auto-rp}", reflect.TypeOf(AutoRp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-autorp-oper:auto-rp", reflect.TypeOf(AutoRp{}))
}

// AutorpProtocolMode represents Autorp protocol mode
type AutorpProtocolMode string

const (
    // sparse
    AutorpProtocolMode_sparse AutorpProtocolMode = "sparse"

    // bidirectional
    AutorpProtocolMode_bidirectional AutorpProtocolMode = "bidirectional"
)

// AutoRp
// AutoRP operational data
type AutoRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Standby Process.
    Standby AutoRp_Standby

    // Active Process.
    Active AutoRp_Active
}

func (autoRp *AutoRp) GetEntityData() *types.CommonEntityData {
    autoRp.EntityData.YFilter = autoRp.YFilter
    autoRp.EntityData.YangName = "auto-rp"
    autoRp.EntityData.BundleName = "cisco_ios_xr"
    autoRp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-autorp-oper"
    autoRp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp"
    autoRp.EntityData.AbsolutePath = autoRp.EntityData.SegmentPath
    autoRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoRp.EntityData.Children = types.NewOrderedMap()
    autoRp.EntityData.Children.Append("standby", types.YChild{"Standby", &autoRp.Standby})
    autoRp.EntityData.Children.Append("active", types.YChild{"Active", &autoRp.Active})
    autoRp.EntityData.Leafs = types.NewOrderedMap()

    autoRp.EntityData.YListKeys = []string {}

    return &(autoRp.EntityData)
}

// AutoRp_Standby
// Standby Process
type AutoRp_Standby struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate RP.
    CandidateRp AutoRp_Standby_CandidateRp

    // AutoRP Mapping Agent Table.
    MappingAgent AutoRp_Standby_MappingAgent
}

func (standby *AutoRp_Standby) GetEntityData() *types.CommonEntityData {
    standby.EntityData.YFilter = standby.YFilter
    standby.EntityData.YangName = "standby"
    standby.EntityData.BundleName = "cisco_ios_xr"
    standby.EntityData.ParentYangName = "auto-rp"
    standby.EntityData.SegmentPath = "standby"
    standby.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/" + standby.EntityData.SegmentPath
    standby.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    standby.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    standby.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    standby.EntityData.Children = types.NewOrderedMap()
    standby.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", &standby.CandidateRp})
    standby.EntityData.Children.Append("mapping-agent", types.YChild{"MappingAgent", &standby.MappingAgent})
    standby.EntityData.Leafs = types.NewOrderedMap()

    standby.EntityData.YListKeys = []string {}

    return &(standby.EntityData)
}

// AutoRp_Standby_CandidateRp
// AutoRP Candidate RP
type AutoRp_Standby_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate Traffic Counters.
    Traffic AutoRp_Standby_CandidateRp_Traffic

    // AutoRP Candidate RP Table.
    Rps AutoRp_Standby_CandidateRp_Rps
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "standby"
    candidateRp.EntityData.SegmentPath = "candidate-rp"
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Children.Append("traffic", types.YChild{"Traffic", &candidateRp.Traffic})
    candidateRp.EntityData.Children.Append("rps", types.YChild{"Rps", &candidateRp.Rps})
    candidateRp.EntityData.Leafs = types.NewOrderedMap()

    candidateRp.EntityData.YListKeys = []string {}

    return &(candidateRp.EntityData)
}

// AutoRp_Standby_CandidateRp_Traffic
// AutoRP Candidate Traffic Counters
type AutoRp_Standby_CandidateRp_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "candidate-rp"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/candidate-rp/" + traffic.EntityData.SegmentPath
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("active-sent-packets", types.YLeaf{"ActiveSentPackets", traffic.ActiveSentPackets})
    traffic.EntityData.Leafs.Append("standby-sent-packets", types.YLeaf{"StandbySentPackets", traffic.StandbySentPackets})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// AutoRp_Standby_CandidateRp_Rps
// AutoRP Candidate RP Table
type AutoRp_Standby_CandidateRp_Rps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate RP Entry. The type is slice of
    // AutoRp_Standby_CandidateRp_Rps_Rp.
    Rp []*AutoRp_Standby_CandidateRp_Rps_Rp
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetEntityData() *types.CommonEntityData {
    rps.EntityData.YFilter = rps.YFilter
    rps.EntityData.YangName = "rps"
    rps.EntityData.BundleName = "cisco_ios_xr"
    rps.EntityData.ParentYangName = "candidate-rp"
    rps.EntityData.SegmentPath = "rps"
    rps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/candidate-rp/" + rps.EntityData.SegmentPath
    rps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rps.EntityData.Children = types.NewOrderedMap()
    rps.EntityData.Children.Append("rp", types.YChild{"Rp", nil})
    for i := range rps.Rp {
        types.SetYListKey(rps.Rp[i], i)
        rps.EntityData.Children.Append(types.GetSegmentPath(rps.Rp[i]), types.YChild{"Rp", rps.Rp[i]})
    }
    rps.EntityData.Leafs = types.NewOrderedMap()

    rps.EntityData.YListKeys = []string {}

    return &(rps.EntityData)
}

// AutoRp_Standby_CandidateRp_Rps_Rp
// AutoRP Candidate RP Entry
type AutoRp_Standby_CandidateRp_Rps_Rp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // ACL Name. The type is string.
    AccessListName interface{}

    // Candidate RP IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CandidateRpAddress interface{}

    // TTL. The type is interface{} with range: -2147483648..2147483647.
    Ttl interface{}

    // Announce Period. The type is interface{} with range:
    // -2147483648..2147483647.
    AnnouncePeriod interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolModeXr interface{}
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetEntityData() *types.CommonEntityData {
    rp.EntityData.YFilter = rp.YFilter
    rp.EntityData.YangName = "rp"
    rp.EntityData.BundleName = "cisco_ios_xr"
    rp.EntityData.ParentYangName = "rps"
    rp.EntityData.SegmentPath = "rp" + types.AddNoKeyToken(rp)
    rp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/candidate-rp/rps/" + rp.EntityData.SegmentPath
    rp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rp.EntityData.Children = types.NewOrderedMap()
    rp.EntityData.Leafs = types.NewOrderedMap()
    rp.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", rp.InterfaceName})
    rp.EntityData.Leafs.Append("protocol-mode", types.YLeaf{"ProtocolMode", rp.ProtocolMode})
    rp.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", rp.AccessListName})
    rp.EntityData.Leafs.Append("candidate-rp-address", types.YLeaf{"CandidateRpAddress", rp.CandidateRpAddress})
    rp.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", rp.Ttl})
    rp.EntityData.Leafs.Append("announce-period", types.YLeaf{"AnnouncePeriod", rp.AnnouncePeriod})
    rp.EntityData.Leafs.Append("protocol-mode-xr", types.YLeaf{"ProtocolModeXr", rp.ProtocolModeXr})

    rp.EntityData.YListKeys = []string {}

    return &(rp.EntityData)
}

// AutoRp_Standby_MappingAgent
// AutoRP Mapping Agent Table
type AutoRp_Standby_MappingAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Traffic Counters.
    Traffic AutoRp_Standby_MappingAgent_Traffic

    // AutoRP Mapping Agent Table Entries.
    RpAddresses AutoRp_Standby_MappingAgent_RpAddresses

    // AutoRP Mapping Agent Summary Information.
    Summary AutoRp_Standby_MappingAgent_Summary
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetEntityData() *types.CommonEntityData {
    mappingAgent.EntityData.YFilter = mappingAgent.YFilter
    mappingAgent.EntityData.YangName = "mapping-agent"
    mappingAgent.EntityData.BundleName = "cisco_ios_xr"
    mappingAgent.EntityData.ParentYangName = "standby"
    mappingAgent.EntityData.SegmentPath = "mapping-agent"
    mappingAgent.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/" + mappingAgent.EntityData.SegmentPath
    mappingAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingAgent.EntityData.Children = types.NewOrderedMap()
    mappingAgent.EntityData.Children.Append("traffic", types.YChild{"Traffic", &mappingAgent.Traffic})
    mappingAgent.EntityData.Children.Append("rp-addresses", types.YChild{"RpAddresses", &mappingAgent.RpAddresses})
    mappingAgent.EntityData.Children.Append("summary", types.YChild{"Summary", &mappingAgent.Summary})
    mappingAgent.EntityData.Leafs = types.NewOrderedMap()

    mappingAgent.EntityData.YListKeys = []string {}

    return &(mappingAgent.EntityData)
}

// AutoRp_Standby_MappingAgent_Traffic
// AutoRP Mapping Agent Traffic Counters
type AutoRp_Standby_MappingAgent_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}

    // Number of packets received in active role. The type is interface{} with
    // range: 0..4294967295.
    ActiveReceivedPackets interface{}

    // Number of packets dropped in receive path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbyReceivedPackets interface{}
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "mapping-agent"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/mapping-agent/" + traffic.EntityData.SegmentPath
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("active-sent-packets", types.YLeaf{"ActiveSentPackets", traffic.ActiveSentPackets})
    traffic.EntityData.Leafs.Append("standby-sent-packets", types.YLeaf{"StandbySentPackets", traffic.StandbySentPackets})
    traffic.EntityData.Leafs.Append("active-received-packets", types.YLeaf{"ActiveReceivedPackets", traffic.ActiveReceivedPackets})
    traffic.EntityData.Leafs.Append("standby-received-packets", types.YLeaf{"StandbyReceivedPackets", traffic.StandbyReceivedPackets})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// AutoRp_Standby_MappingAgent_RpAddresses
// AutoRP Mapping Agent Table Entries
type AutoRp_Standby_MappingAgent_RpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Entry. The type is slice of
    // AutoRp_Standby_MappingAgent_RpAddresses_RpAddress.
    RpAddress []*AutoRp_Standby_MappingAgent_RpAddresses_RpAddress
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetEntityData() *types.CommonEntityData {
    rpAddresses.EntityData.YFilter = rpAddresses.YFilter
    rpAddresses.EntityData.YangName = "rp-addresses"
    rpAddresses.EntityData.BundleName = "cisco_ios_xr"
    rpAddresses.EntityData.ParentYangName = "mapping-agent"
    rpAddresses.EntityData.SegmentPath = "rp-addresses"
    rpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/mapping-agent/" + rpAddresses.EntityData.SegmentPath
    rpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpAddresses.EntityData.Children = types.NewOrderedMap()
    rpAddresses.EntityData.Children.Append("rp-address", types.YChild{"RpAddress", nil})
    for i := range rpAddresses.RpAddress {
        rpAddresses.EntityData.Children.Append(types.GetSegmentPath(rpAddresses.RpAddress[i]), types.YChild{"RpAddress", rpAddresses.RpAddress[i]})
    }
    rpAddresses.EntityData.Leafs = types.NewOrderedMap()

    rpAddresses.EntityData.YListKeys = []string {}

    return &(rpAddresses.EntityData)
}

// AutoRp_Standby_MappingAgent_RpAddresses_RpAddress
// AutoRP Mapping Agent Entry
type AutoRp_Standby_MappingAgent_RpAddresses_RpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RpAddress interface{}

    // Candidate-RP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RpAddressXr interface{}

    // Time for expiration in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ExpiryTime interface{}

    // PIM version of the CRP. The type is interface{} with range: 0..255.
    PimVersion interface{}

    // Array of ranges. The type is slice of
    // AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range.
    Range []*AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetEntityData() *types.CommonEntityData {
    rpAddress.EntityData.YFilter = rpAddress.YFilter
    rpAddress.EntityData.YangName = "rp-address"
    rpAddress.EntityData.BundleName = "cisco_ios_xr"
    rpAddress.EntityData.ParentYangName = "rp-addresses"
    rpAddress.EntityData.SegmentPath = "rp-address" + types.AddKeyToken(rpAddress.RpAddress, "rp-address")
    rpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/mapping-agent/rp-addresses/" + rpAddress.EntityData.SegmentPath
    rpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpAddress.EntityData.Children = types.NewOrderedMap()
    rpAddress.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range rpAddress.Range {
        types.SetYListKey(rpAddress.Range[i], i)
        rpAddress.EntityData.Children.Append(types.GetSegmentPath(rpAddress.Range[i]), types.YChild{"Range", rpAddress.Range[i]})
    }
    rpAddress.EntityData.Leafs = types.NewOrderedMap()
    rpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", rpAddress.RpAddress})
    rpAddress.EntityData.Leafs.Append("rp-address-xr", types.YLeaf{"RpAddressXr", rpAddress.RpAddressXr})
    rpAddress.EntityData.Leafs.Append("expiry-time", types.YLeaf{"ExpiryTime", rpAddress.ExpiryTime})
    rpAddress.EntityData.Leafs.Append("pim-version", types.YLeaf{"PimVersion", rpAddress.PimVersion})

    rpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(rpAddress.EntityData)
}

// AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range
// Array of ranges
type AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Prefix of the range. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length of the range. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolMode interface{}

    // Is this entry advertised ?. The type is bool.
    IsAdvertised interface{}

    // Source of the entry. The type is interface{} with range: 0..255.
    CreateType interface{}

    // Checkpoint object id. The type is interface{} with range: 0..4294967295.
    CheckPointObjectId interface{}

    // Uptime in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Uptime interface{}
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "rp-address"
    self.EntityData.SegmentPath = "range" + types.AddNoKeyToken(self)
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/mapping-agent/rp-addresses/rp-address/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("protocol-mode", types.YLeaf{"ProtocolMode", self.ProtocolMode})
    self.EntityData.Leafs.Append("is-advertised", types.YLeaf{"IsAdvertised", self.IsAdvertised})
    self.EntityData.Leafs.Append("create-type", types.YLeaf{"CreateType", self.CreateType})
    self.EntityData.Leafs.Append("check-point-object-id", types.YLeaf{"CheckPointObjectId", self.CheckPointObjectId})
    self.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", self.Uptime})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// AutoRp_Standby_MappingAgent_Summary
// AutoRP Mapping Agent Summary Information
type AutoRp_Standby_MappingAgent_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is maximum enforcement disabled ?. The type is bool.
    IsMaximumDisabled interface{}

    // Maximum group to RP mapping entries allowed. The type is interface{} with
    // range: 0..4294967295.
    CacheLimit interface{}

    // Number of group to RP mapping entries in the cache. The type is interface{}
    // with range: 0..4294967295.
    CacheCount interface{}
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "mapping-agent"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/standby/mapping-agent/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("is-maximum-disabled", types.YLeaf{"IsMaximumDisabled", summary.IsMaximumDisabled})
    summary.EntityData.Leafs.Append("cache-limit", types.YLeaf{"CacheLimit", summary.CacheLimit})
    summary.EntityData.Leafs.Append("cache-count", types.YLeaf{"CacheCount", summary.CacheCount})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// AutoRp_Active
// Active Process
type AutoRp_Active struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate RP.
    CandidateRp AutoRp_Active_CandidateRp

    // AutoRP Mapping Agent Table.
    MappingAgent AutoRp_Active_MappingAgent
}

func (active *AutoRp_Active) GetEntityData() *types.CommonEntityData {
    active.EntityData.YFilter = active.YFilter
    active.EntityData.YangName = "active"
    active.EntityData.BundleName = "cisco_ios_xr"
    active.EntityData.ParentYangName = "auto-rp"
    active.EntityData.SegmentPath = "active"
    active.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/" + active.EntityData.SegmentPath
    active.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    active.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    active.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    active.EntityData.Children = types.NewOrderedMap()
    active.EntityData.Children.Append("candidate-rp", types.YChild{"CandidateRp", &active.CandidateRp})
    active.EntityData.Children.Append("mapping-agent", types.YChild{"MappingAgent", &active.MappingAgent})
    active.EntityData.Leafs = types.NewOrderedMap()

    active.EntityData.YListKeys = []string {}

    return &(active.EntityData)
}

// AutoRp_Active_CandidateRp
// AutoRP Candidate RP
type AutoRp_Active_CandidateRp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate Traffic Counters.
    Traffic AutoRp_Active_CandidateRp_Traffic

    // AutoRP Candidate RP Table.
    Rps AutoRp_Active_CandidateRp_Rps
}

func (candidateRp *AutoRp_Active_CandidateRp) GetEntityData() *types.CommonEntityData {
    candidateRp.EntityData.YFilter = candidateRp.YFilter
    candidateRp.EntityData.YangName = "candidate-rp"
    candidateRp.EntityData.BundleName = "cisco_ios_xr"
    candidateRp.EntityData.ParentYangName = "active"
    candidateRp.EntityData.SegmentPath = "candidate-rp"
    candidateRp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/" + candidateRp.EntityData.SegmentPath
    candidateRp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateRp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateRp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateRp.EntityData.Children = types.NewOrderedMap()
    candidateRp.EntityData.Children.Append("traffic", types.YChild{"Traffic", &candidateRp.Traffic})
    candidateRp.EntityData.Children.Append("rps", types.YChild{"Rps", &candidateRp.Rps})
    candidateRp.EntityData.Leafs = types.NewOrderedMap()

    candidateRp.EntityData.YListKeys = []string {}

    return &(candidateRp.EntityData)
}

// AutoRp_Active_CandidateRp_Traffic
// AutoRP Candidate Traffic Counters
type AutoRp_Active_CandidateRp_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "candidate-rp"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/candidate-rp/" + traffic.EntityData.SegmentPath
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("active-sent-packets", types.YLeaf{"ActiveSentPackets", traffic.ActiveSentPackets})
    traffic.EntityData.Leafs.Append("standby-sent-packets", types.YLeaf{"StandbySentPackets", traffic.StandbySentPackets})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// AutoRp_Active_CandidateRp_Rps
// AutoRP Candidate RP Table
type AutoRp_Active_CandidateRp_Rps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Candidate RP Entry. The type is slice of
    // AutoRp_Active_CandidateRp_Rps_Rp.
    Rp []*AutoRp_Active_CandidateRp_Rps_Rp
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetEntityData() *types.CommonEntityData {
    rps.EntityData.YFilter = rps.YFilter
    rps.EntityData.YangName = "rps"
    rps.EntityData.BundleName = "cisco_ios_xr"
    rps.EntityData.ParentYangName = "candidate-rp"
    rps.EntityData.SegmentPath = "rps"
    rps.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/candidate-rp/" + rps.EntityData.SegmentPath
    rps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rps.EntityData.Children = types.NewOrderedMap()
    rps.EntityData.Children.Append("rp", types.YChild{"Rp", nil})
    for i := range rps.Rp {
        types.SetYListKey(rps.Rp[i], i)
        rps.EntityData.Children.Append(types.GetSegmentPath(rps.Rp[i]), types.YChild{"Rp", rps.Rp[i]})
    }
    rps.EntityData.Leafs = types.NewOrderedMap()

    rps.EntityData.YListKeys = []string {}

    return &(rps.EntityData)
}

// AutoRp_Active_CandidateRp_Rps_Rp
// AutoRP Candidate RP Entry
type AutoRp_Active_CandidateRp_Rps_Rp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface Name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // ACL Name. The type is string.
    AccessListName interface{}

    // Candidate RP IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CandidateRpAddress interface{}

    // TTL. The type is interface{} with range: -2147483648..2147483647.
    Ttl interface{}

    // Announce Period. The type is interface{} with range:
    // -2147483648..2147483647.
    AnnouncePeriod interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolModeXr interface{}
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetEntityData() *types.CommonEntityData {
    rp.EntityData.YFilter = rp.YFilter
    rp.EntityData.YangName = "rp"
    rp.EntityData.BundleName = "cisco_ios_xr"
    rp.EntityData.ParentYangName = "rps"
    rp.EntityData.SegmentPath = "rp" + types.AddNoKeyToken(rp)
    rp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/candidate-rp/rps/" + rp.EntityData.SegmentPath
    rp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rp.EntityData.Children = types.NewOrderedMap()
    rp.EntityData.Leafs = types.NewOrderedMap()
    rp.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", rp.InterfaceName})
    rp.EntityData.Leafs.Append("protocol-mode", types.YLeaf{"ProtocolMode", rp.ProtocolMode})
    rp.EntityData.Leafs.Append("access-list-name", types.YLeaf{"AccessListName", rp.AccessListName})
    rp.EntityData.Leafs.Append("candidate-rp-address", types.YLeaf{"CandidateRpAddress", rp.CandidateRpAddress})
    rp.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", rp.Ttl})
    rp.EntityData.Leafs.Append("announce-period", types.YLeaf{"AnnouncePeriod", rp.AnnouncePeriod})
    rp.EntityData.Leafs.Append("protocol-mode-xr", types.YLeaf{"ProtocolModeXr", rp.ProtocolModeXr})

    rp.EntityData.YListKeys = []string {}

    return &(rp.EntityData)
}

// AutoRp_Active_MappingAgent
// AutoRP Mapping Agent Table
type AutoRp_Active_MappingAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Traffic Counters.
    Traffic AutoRp_Active_MappingAgent_Traffic

    // AutoRP Mapping Agent Table Entries.
    RpAddresses AutoRp_Active_MappingAgent_RpAddresses

    // AutoRP Mapping Agent Summary Information.
    Summary AutoRp_Active_MappingAgent_Summary
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetEntityData() *types.CommonEntityData {
    mappingAgent.EntityData.YFilter = mappingAgent.YFilter
    mappingAgent.EntityData.YangName = "mapping-agent"
    mappingAgent.EntityData.BundleName = "cisco_ios_xr"
    mappingAgent.EntityData.ParentYangName = "active"
    mappingAgent.EntityData.SegmentPath = "mapping-agent"
    mappingAgent.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/" + mappingAgent.EntityData.SegmentPath
    mappingAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mappingAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mappingAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mappingAgent.EntityData.Children = types.NewOrderedMap()
    mappingAgent.EntityData.Children.Append("traffic", types.YChild{"Traffic", &mappingAgent.Traffic})
    mappingAgent.EntityData.Children.Append("rp-addresses", types.YChild{"RpAddresses", &mappingAgent.RpAddresses})
    mappingAgent.EntityData.Children.Append("summary", types.YChild{"Summary", &mappingAgent.Summary})
    mappingAgent.EntityData.Leafs = types.NewOrderedMap()

    mappingAgent.EntityData.YListKeys = []string {}

    return &(mappingAgent.EntityData)
}

// AutoRp_Active_MappingAgent_Traffic
// AutoRP Mapping Agent Traffic Counters
type AutoRp_Active_MappingAgent_Traffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}

    // Number of packets received in active role. The type is interface{} with
    // range: 0..4294967295.
    ActiveReceivedPackets interface{}

    // Number of packets dropped in receive path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbyReceivedPackets interface{}
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetEntityData() *types.CommonEntityData {
    traffic.EntityData.YFilter = traffic.YFilter
    traffic.EntityData.YangName = "traffic"
    traffic.EntityData.BundleName = "cisco_ios_xr"
    traffic.EntityData.ParentYangName = "mapping-agent"
    traffic.EntityData.SegmentPath = "traffic"
    traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/mapping-agent/" + traffic.EntityData.SegmentPath
    traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traffic.EntityData.Children = types.NewOrderedMap()
    traffic.EntityData.Leafs = types.NewOrderedMap()
    traffic.EntityData.Leafs.Append("active-sent-packets", types.YLeaf{"ActiveSentPackets", traffic.ActiveSentPackets})
    traffic.EntityData.Leafs.Append("standby-sent-packets", types.YLeaf{"StandbySentPackets", traffic.StandbySentPackets})
    traffic.EntityData.Leafs.Append("active-received-packets", types.YLeaf{"ActiveReceivedPackets", traffic.ActiveReceivedPackets})
    traffic.EntityData.Leafs.Append("standby-received-packets", types.YLeaf{"StandbyReceivedPackets", traffic.StandbyReceivedPackets})

    traffic.EntityData.YListKeys = []string {}

    return &(traffic.EntityData)
}

// AutoRp_Active_MappingAgent_RpAddresses
// AutoRP Mapping Agent Table Entries
type AutoRp_Active_MappingAgent_RpAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Entry. The type is slice of
    // AutoRp_Active_MappingAgent_RpAddresses_RpAddress.
    RpAddress []*AutoRp_Active_MappingAgent_RpAddresses_RpAddress
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetEntityData() *types.CommonEntityData {
    rpAddresses.EntityData.YFilter = rpAddresses.YFilter
    rpAddresses.EntityData.YangName = "rp-addresses"
    rpAddresses.EntityData.BundleName = "cisco_ios_xr"
    rpAddresses.EntityData.ParentYangName = "mapping-agent"
    rpAddresses.EntityData.SegmentPath = "rp-addresses"
    rpAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/mapping-agent/" + rpAddresses.EntityData.SegmentPath
    rpAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpAddresses.EntityData.Children = types.NewOrderedMap()
    rpAddresses.EntityData.Children.Append("rp-address", types.YChild{"RpAddress", nil})
    for i := range rpAddresses.RpAddress {
        rpAddresses.EntityData.Children.Append(types.GetSegmentPath(rpAddresses.RpAddress[i]), types.YChild{"RpAddress", rpAddresses.RpAddress[i]})
    }
    rpAddresses.EntityData.Leafs = types.NewOrderedMap()

    rpAddresses.EntityData.YListKeys = []string {}

    return &(rpAddresses.EntityData)
}

// AutoRp_Active_MappingAgent_RpAddresses_RpAddress
// AutoRP Mapping Agent Entry
type AutoRp_Active_MappingAgent_RpAddresses_RpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RpAddress interface{}

    // Candidate-RP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RpAddressXr interface{}

    // Time for expiration in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ExpiryTime interface{}

    // PIM version of the CRP. The type is interface{} with range: 0..255.
    PimVersion interface{}

    // Array of ranges. The type is slice of
    // AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range.
    Range []*AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetEntityData() *types.CommonEntityData {
    rpAddress.EntityData.YFilter = rpAddress.YFilter
    rpAddress.EntityData.YangName = "rp-address"
    rpAddress.EntityData.BundleName = "cisco_ios_xr"
    rpAddress.EntityData.ParentYangName = "rp-addresses"
    rpAddress.EntityData.SegmentPath = "rp-address" + types.AddKeyToken(rpAddress.RpAddress, "rp-address")
    rpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/mapping-agent/rp-addresses/" + rpAddress.EntityData.SegmentPath
    rpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rpAddress.EntityData.Children = types.NewOrderedMap()
    rpAddress.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range rpAddress.Range {
        types.SetYListKey(rpAddress.Range[i], i)
        rpAddress.EntityData.Children.Append(types.GetSegmentPath(rpAddress.Range[i]), types.YChild{"Range", rpAddress.Range[i]})
    }
    rpAddress.EntityData.Leafs = types.NewOrderedMap()
    rpAddress.EntityData.Leafs.Append("rp-address", types.YLeaf{"RpAddress", rpAddress.RpAddress})
    rpAddress.EntityData.Leafs.Append("rp-address-xr", types.YLeaf{"RpAddressXr", rpAddress.RpAddressXr})
    rpAddress.EntityData.Leafs.Append("expiry-time", types.YLeaf{"ExpiryTime", rpAddress.ExpiryTime})
    rpAddress.EntityData.Leafs.Append("pim-version", types.YLeaf{"PimVersion", rpAddress.PimVersion})

    rpAddress.EntityData.YListKeys = []string {"RpAddress"}

    return &(rpAddress.EntityData)
}

// AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range
// Array of ranges
type AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Prefix of the range. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length of the range. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolMode interface{}

    // Is this entry advertised ?. The type is bool.
    IsAdvertised interface{}

    // Source of the entry. The type is interface{} with range: 0..255.
    CreateType interface{}

    // Checkpoint object id. The type is interface{} with range: 0..4294967295.
    CheckPointObjectId interface{}

    // Uptime in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    Uptime interface{}
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "rp-address"
    self.EntityData.SegmentPath = "range" + types.AddNoKeyToken(self)
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/mapping-agent/rp-addresses/rp-address/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("protocol-mode", types.YLeaf{"ProtocolMode", self.ProtocolMode})
    self.EntityData.Leafs.Append("is-advertised", types.YLeaf{"IsAdvertised", self.IsAdvertised})
    self.EntityData.Leafs.Append("create-type", types.YLeaf{"CreateType", self.CreateType})
    self.EntityData.Leafs.Append("check-point-object-id", types.YLeaf{"CheckPointObjectId", self.CheckPointObjectId})
    self.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", self.Uptime})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// AutoRp_Active_MappingAgent_Summary
// AutoRP Mapping Agent Summary Information
type AutoRp_Active_MappingAgent_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is maximum enforcement disabled ?. The type is bool.
    IsMaximumDisabled interface{}

    // Maximum group to RP mapping entries allowed. The type is interface{} with
    // range: 0..4294967295.
    CacheLimit interface{}

    // Number of group to RP mapping entries in the cache. The type is interface{}
    // with range: 0..4294967295.
    CacheCount interface{}
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "mapping-agent"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp/active/mapping-agent/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("is-maximum-disabled", types.YLeaf{"IsMaximumDisabled", summary.IsMaximumDisabled})
    summary.EntityData.Leafs.Append("cache-limit", types.YLeaf{"CacheLimit", summary.CacheLimit})
    summary.EntityData.Leafs.Append("cache-count", types.YLeaf{"CacheCount", summary.CacheCount})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

