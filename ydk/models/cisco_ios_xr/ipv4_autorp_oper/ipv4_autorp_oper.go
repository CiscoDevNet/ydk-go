// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-autorp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   auto-rp: AutoRP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Standby Process.
    Standby AutoRp_Standby

    // Active Process.
    Active AutoRp_Active
}

func (autoRp *AutoRp) GetFilter() yfilter.YFilter { return autoRp.YFilter }

func (autoRp *AutoRp) SetFilter(yf yfilter.YFilter) { autoRp.YFilter = yf }

func (autoRp *AutoRp) GetGoName(yname string) string {
    if yname == "standby" { return "Standby" }
    if yname == "active" { return "Active" }
    return ""
}

func (autoRp *AutoRp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-autorp-oper:auto-rp"
}

func (autoRp *AutoRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "standby" {
        return &autoRp.Standby
    }
    if childYangName == "active" {
        return &autoRp.Active
    }
    return nil
}

func (autoRp *AutoRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["standby"] = &autoRp.Standby
    children["active"] = &autoRp.Active
    return children
}

func (autoRp *AutoRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (autoRp *AutoRp) GetBundleName() string { return "cisco_ios_xr" }

func (autoRp *AutoRp) GetYangName() string { return "auto-rp" }

func (autoRp *AutoRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (autoRp *AutoRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (autoRp *AutoRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (autoRp *AutoRp) SetParent(parent types.Entity) { autoRp.parent = parent }

func (autoRp *AutoRp) GetParent() types.Entity { return autoRp.parent }

func (autoRp *AutoRp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-autorp-oper" }

// AutoRp_Standby
// Standby Process
type AutoRp_Standby struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate RP.
    CandidateRp AutoRp_Standby_CandidateRp

    // AutoRP Mapping Agent Table.
    MappingAgent AutoRp_Standby_MappingAgent
}

func (standby *AutoRp_Standby) GetFilter() yfilter.YFilter { return standby.YFilter }

func (standby *AutoRp_Standby) SetFilter(yf yfilter.YFilter) { standby.YFilter = yf }

func (standby *AutoRp_Standby) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    if yname == "mapping-agent" { return "MappingAgent" }
    return ""
}

func (standby *AutoRp_Standby) GetSegmentPath() string {
    return "standby"
}

func (standby *AutoRp_Standby) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        return &standby.CandidateRp
    }
    if childYangName == "mapping-agent" {
        return &standby.MappingAgent
    }
    return nil
}

func (standby *AutoRp_Standby) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-rp"] = &standby.CandidateRp
    children["mapping-agent"] = &standby.MappingAgent
    return children
}

func (standby *AutoRp_Standby) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (standby *AutoRp_Standby) GetBundleName() string { return "cisco_ios_xr" }

func (standby *AutoRp_Standby) GetYangName() string { return "standby" }

func (standby *AutoRp_Standby) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (standby *AutoRp_Standby) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (standby *AutoRp_Standby) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (standby *AutoRp_Standby) SetParent(parent types.Entity) { standby.parent = parent }

func (standby *AutoRp_Standby) GetParent() types.Entity { return standby.parent }

func (standby *AutoRp_Standby) GetParentYangName() string { return "auto-rp" }

// AutoRp_Standby_CandidateRp
// AutoRP Candidate RP
type AutoRp_Standby_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate Traffic Counters.
    Traffic AutoRp_Standby_CandidateRp_Traffic

    // AutoRP Candidate RP Table.
    Rps AutoRp_Standby_CandidateRp_Rps
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *AutoRp_Standby_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *AutoRp_Standby_CandidateRp) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    if yname == "rps" { return "Rps" }
    return ""
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetSegmentPath() string {
    return "candidate-rp"
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &candidateRp.Traffic
    }
    if childYangName == "rps" {
        return &candidateRp.Rps
    }
    return nil
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &candidateRp.Traffic
    children["rps"] = &candidateRp.Rps
    return children
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRp *AutoRp_Standby_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *AutoRp_Standby_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *AutoRp_Standby_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *AutoRp_Standby_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *AutoRp_Standby_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *AutoRp_Standby_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *AutoRp_Standby_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *AutoRp_Standby_CandidateRp) GetParentYangName() string { return "standby" }

// AutoRp_Standby_CandidateRp_Traffic
// AutoRP Candidate Traffic Counters
type AutoRp_Standby_CandidateRp_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetGoName(yname string) string {
    if yname == "active-sent-packets" { return "ActiveSentPackets" }
    if yname == "standby-sent-packets" { return "StandbySentPackets" }
    return ""
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-sent-packets"] = traffic.ActiveSentPackets
    leafs["standby-sent-packets"] = traffic.StandbySentPackets
    return leafs
}

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetYangName() string { return "traffic" }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *AutoRp_Standby_CandidateRp_Traffic) GetParentYangName() string { return "candidate-rp" }

// AutoRp_Standby_CandidateRp_Rps
// AutoRP Candidate RP Table
type AutoRp_Standby_CandidateRp_Rps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate RP Entry. The type is slice of
    // AutoRp_Standby_CandidateRp_Rps_Rp.
    Rp []AutoRp_Standby_CandidateRp_Rps_Rp
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetFilter() yfilter.YFilter { return rps.YFilter }

func (rps *AutoRp_Standby_CandidateRp_Rps) SetFilter(yf yfilter.YFilter) { rps.YFilter = yf }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetGoName(yname string) string {
    if yname == "rp" { return "Rp" }
    return ""
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetSegmentPath() string {
    return "rps"
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rp" {
        for _, c := range rps.Rp {
            if rps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Standby_CandidateRp_Rps_Rp{}
        rps.Rp = append(rps.Rp, child)
        return &rps.Rp[len(rps.Rp)-1]
    }
    return nil
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rps.Rp {
        children[rps.Rp[i].GetSegmentPath()] = &rps.Rp[i]
    }
    return children
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rps *AutoRp_Standby_CandidateRp_Rps) GetBundleName() string { return "cisco_ios_xr" }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetYangName() string { return "rps" }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rps *AutoRp_Standby_CandidateRp_Rps) SetParent(parent types.Entity) { rps.parent = parent }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetParent() types.Entity { return rps.parent }

func (rps *AutoRp_Standby_CandidateRp_Rps) GetParentYangName() string { return "candidate-rp" }

// AutoRp_Standby_CandidateRp_Rps_Rp
// AutoRP Candidate RP Entry
type AutoRp_Standby_CandidateRp_Rps_Rp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // ACL Name. The type is string.
    AccessListName interface{}

    // Candidate RP IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CandidateRpAddress interface{}

    // TTL. The type is interface{} with range: -2147483648..2147483647.
    Ttl interface{}

    // Announce Period. The type is interface{} with range:
    // -2147483648..2147483647.
    AnnouncePeriod interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolModeXr interface{}
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetFilter() yfilter.YFilter { return rp.YFilter }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) SetFilter(yf yfilter.YFilter) { rp.YFilter = yf }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "protocol-mode" { return "ProtocolMode" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "candidate-rp-address" { return "CandidateRpAddress" }
    if yname == "ttl" { return "Ttl" }
    if yname == "announce-period" { return "AnnouncePeriod" }
    if yname == "protocol-mode-xr" { return "ProtocolModeXr" }
    return ""
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetSegmentPath() string {
    return "rp"
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = rp.InterfaceName
    leafs["protocol-mode"] = rp.ProtocolMode
    leafs["access-list-name"] = rp.AccessListName
    leafs["candidate-rp-address"] = rp.CandidateRpAddress
    leafs["ttl"] = rp.Ttl
    leafs["announce-period"] = rp.AnnouncePeriod
    leafs["protocol-mode-xr"] = rp.ProtocolModeXr
    return leafs
}

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetBundleName() string { return "cisco_ios_xr" }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetYangName() string { return "rp" }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) SetParent(parent types.Entity) { rp.parent = parent }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetParent() types.Entity { return rp.parent }

func (rp *AutoRp_Standby_CandidateRp_Rps_Rp) GetParentYangName() string { return "rps" }

// AutoRp_Standby_MappingAgent
// AutoRP Mapping Agent Table
type AutoRp_Standby_MappingAgent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Traffic Counters.
    Traffic AutoRp_Standby_MappingAgent_Traffic

    // AutoRP Mapping Agent Table Entries.
    RpAddresses AutoRp_Standby_MappingAgent_RpAddresses

    // AutoRP Mapping Agent Summary Information.
    Summary AutoRp_Standby_MappingAgent_Summary
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetFilter() yfilter.YFilter { return mappingAgent.YFilter }

func (mappingAgent *AutoRp_Standby_MappingAgent) SetFilter(yf yfilter.YFilter) { mappingAgent.YFilter = yf }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    if yname == "rp-addresses" { return "RpAddresses" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetSegmentPath() string {
    return "mapping-agent"
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &mappingAgent.Traffic
    }
    if childYangName == "rp-addresses" {
        return &mappingAgent.RpAddresses
    }
    if childYangName == "summary" {
        return &mappingAgent.Summary
    }
    return nil
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &mappingAgent.Traffic
    children["rp-addresses"] = &mappingAgent.RpAddresses
    children["summary"] = &mappingAgent.Summary
    return children
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mappingAgent *AutoRp_Standby_MappingAgent) GetBundleName() string { return "cisco_ios_xr" }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetYangName() string { return "mapping-agent" }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingAgent *AutoRp_Standby_MappingAgent) SetParent(parent types.Entity) { mappingAgent.parent = parent }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetParent() types.Entity { return mappingAgent.parent }

func (mappingAgent *AutoRp_Standby_MappingAgent) GetParentYangName() string { return "standby" }

// AutoRp_Standby_MappingAgent_Traffic
// AutoRP Mapping Agent Traffic Counters
type AutoRp_Standby_MappingAgent_Traffic struct {
    parent types.Entity
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

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetGoName(yname string) string {
    if yname == "active-sent-packets" { return "ActiveSentPackets" }
    if yname == "standby-sent-packets" { return "StandbySentPackets" }
    if yname == "active-received-packets" { return "ActiveReceivedPackets" }
    if yname == "standby-received-packets" { return "StandbyReceivedPackets" }
    return ""
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-sent-packets"] = traffic.ActiveSentPackets
    leafs["standby-sent-packets"] = traffic.StandbySentPackets
    leafs["active-received-packets"] = traffic.ActiveReceivedPackets
    leafs["standby-received-packets"] = traffic.StandbyReceivedPackets
    return leafs
}

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetYangName() string { return "traffic" }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *AutoRp_Standby_MappingAgent_Traffic) GetParentYangName() string { return "mapping-agent" }

// AutoRp_Standby_MappingAgent_RpAddresses
// AutoRP Mapping Agent Table Entries
type AutoRp_Standby_MappingAgent_RpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Entry. The type is slice of
    // AutoRp_Standby_MappingAgent_RpAddresses_RpAddress.
    RpAddress []AutoRp_Standby_MappingAgent_RpAddresses_RpAddress
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetFilter() yfilter.YFilter { return rpAddresses.YFilter }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) SetFilter(yf yfilter.YFilter) { rpAddresses.YFilter = yf }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    return ""
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetSegmentPath() string {
    return "rp-addresses"
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rp-address" {
        for _, c := range rpAddresses.RpAddress {
            if rpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Standby_MappingAgent_RpAddresses_RpAddress{}
        rpAddresses.RpAddress = append(rpAddresses.RpAddress, child)
        return &rpAddresses.RpAddress[len(rpAddresses.RpAddress)-1]
    }
    return nil
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rpAddresses.RpAddress {
        children[rpAddresses.RpAddress[i].GetSegmentPath()] = &rpAddresses.RpAddress[i]
    }
    return children
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetYangName() string { return "rp-addresses" }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) SetParent(parent types.Entity) { rpAddresses.parent = parent }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetParent() types.Entity { return rpAddresses.parent }

func (rpAddresses *AutoRp_Standby_MappingAgent_RpAddresses) GetParentYangName() string { return "mapping-agent" }

// AutoRp_Standby_MappingAgent_RpAddresses_RpAddress
// AutoRP Mapping Agent Entry
type AutoRp_Standby_MappingAgent_RpAddresses_RpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Candidate-RP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpAddressXr interface{}

    // Time for expiration in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ExpiryTime interface{}

    // PIM version of the CRP. The type is interface{} with range: 0..255.
    PimVersion interface{}

    // Array of ranges. The type is slice of
    // AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range.
    Range []AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetFilter() yfilter.YFilter { return rpAddress.YFilter }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) SetFilter(yf yfilter.YFilter) { rpAddress.YFilter = yf }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "rp-address-xr" { return "RpAddressXr" }
    if yname == "expiry-time" { return "ExpiryTime" }
    if yname == "pim-version" { return "PimVersion" }
    if yname == "range" { return "Range" }
    return ""
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetSegmentPath() string {
    return "rp-address" + "[rp-address='" + fmt.Sprintf("%v", rpAddress.RpAddress) + "']"
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "range" {
        for _, c := range rpAddress.Range {
            if rpAddress.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range{}
        rpAddress.Range = append(rpAddress.Range, child)
        return &rpAddress.Range[len(rpAddress.Range)-1]
    }
    return nil
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rpAddress.Range {
        children[rpAddress.Range[i].GetSegmentPath()] = &rpAddress.Range[i]
    }
    return children
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = rpAddress.RpAddress
    leafs["rp-address-xr"] = rpAddress.RpAddressXr
    leafs["expiry-time"] = rpAddress.ExpiryTime
    leafs["pim-version"] = rpAddress.PimVersion
    return leafs
}

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetYangName() string { return "rp-address" }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) SetParent(parent types.Entity) { rpAddress.parent = parent }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetParent() types.Entity { return rpAddress.parent }

func (rpAddress *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress) GetParentYangName() string { return "rp-addresses" }

// AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range
// Array of ranges
type AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix of the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "protocol-mode" { return "ProtocolMode" }
    if yname == "is-advertised" { return "IsAdvertised" }
    if yname == "create-type" { return "CreateType" }
    if yname == "check-point-object-id" { return "CheckPointObjectId" }
    if yname == "uptime" { return "Uptime" }
    return ""
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetSegmentPath() string {
    return "range"
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = self.Prefix
    leafs["prefix-length"] = self.PrefixLength
    leafs["protocol-mode"] = self.ProtocolMode
    leafs["is-advertised"] = self.IsAdvertised
    leafs["create-type"] = self.CreateType
    leafs["check-point-object-id"] = self.CheckPointObjectId
    leafs["uptime"] = self.Uptime
    return leafs
}

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetBundleName() string { return "cisco_ios_xr" }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetYangName() string { return "range" }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetParent() types.Entity { return self.parent }

func (self *AutoRp_Standby_MappingAgent_RpAddresses_RpAddress_Range) GetParentYangName() string { return "rp-address" }

// AutoRp_Standby_MappingAgent_Summary
// AutoRP Mapping Agent Summary Information
type AutoRp_Standby_MappingAgent_Summary struct {
    parent types.Entity
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

func (summary *AutoRp_Standby_MappingAgent_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *AutoRp_Standby_MappingAgent_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetGoName(yname string) string {
    if yname == "is-maximum-disabled" { return "IsMaximumDisabled" }
    if yname == "cache-limit" { return "CacheLimit" }
    if yname == "cache-count" { return "CacheCount" }
    return ""
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-maximum-disabled"] = summary.IsMaximumDisabled
    leafs["cache-limit"] = summary.CacheLimit
    leafs["cache-count"] = summary.CacheCount
    return leafs
}

func (summary *AutoRp_Standby_MappingAgent_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetYangName() string { return "summary" }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *AutoRp_Standby_MappingAgent_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetParent() types.Entity { return summary.parent }

func (summary *AutoRp_Standby_MappingAgent_Summary) GetParentYangName() string { return "mapping-agent" }

// AutoRp_Active
// Active Process
type AutoRp_Active struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate RP.
    CandidateRp AutoRp_Active_CandidateRp

    // AutoRP Mapping Agent Table.
    MappingAgent AutoRp_Active_MappingAgent
}

func (active *AutoRp_Active) GetFilter() yfilter.YFilter { return active.YFilter }

func (active *AutoRp_Active) SetFilter(yf yfilter.YFilter) { active.YFilter = yf }

func (active *AutoRp_Active) GetGoName(yname string) string {
    if yname == "candidate-rp" { return "CandidateRp" }
    if yname == "mapping-agent" { return "MappingAgent" }
    return ""
}

func (active *AutoRp_Active) GetSegmentPath() string {
    return "active"
}

func (active *AutoRp_Active) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "candidate-rp" {
        return &active.CandidateRp
    }
    if childYangName == "mapping-agent" {
        return &active.MappingAgent
    }
    return nil
}

func (active *AutoRp_Active) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["candidate-rp"] = &active.CandidateRp
    children["mapping-agent"] = &active.MappingAgent
    return children
}

func (active *AutoRp_Active) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (active *AutoRp_Active) GetBundleName() string { return "cisco_ios_xr" }

func (active *AutoRp_Active) GetYangName() string { return "active" }

func (active *AutoRp_Active) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (active *AutoRp_Active) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (active *AutoRp_Active) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (active *AutoRp_Active) SetParent(parent types.Entity) { active.parent = parent }

func (active *AutoRp_Active) GetParent() types.Entity { return active.parent }

func (active *AutoRp_Active) GetParentYangName() string { return "auto-rp" }

// AutoRp_Active_CandidateRp
// AutoRP Candidate RP
type AutoRp_Active_CandidateRp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate Traffic Counters.
    Traffic AutoRp_Active_CandidateRp_Traffic

    // AutoRP Candidate RP Table.
    Rps AutoRp_Active_CandidateRp_Rps
}

func (candidateRp *AutoRp_Active_CandidateRp) GetFilter() yfilter.YFilter { return candidateRp.YFilter }

func (candidateRp *AutoRp_Active_CandidateRp) SetFilter(yf yfilter.YFilter) { candidateRp.YFilter = yf }

func (candidateRp *AutoRp_Active_CandidateRp) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    if yname == "rps" { return "Rps" }
    return ""
}

func (candidateRp *AutoRp_Active_CandidateRp) GetSegmentPath() string {
    return "candidate-rp"
}

func (candidateRp *AutoRp_Active_CandidateRp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &candidateRp.Traffic
    }
    if childYangName == "rps" {
        return &candidateRp.Rps
    }
    return nil
}

func (candidateRp *AutoRp_Active_CandidateRp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &candidateRp.Traffic
    children["rps"] = &candidateRp.Rps
    return children
}

func (candidateRp *AutoRp_Active_CandidateRp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (candidateRp *AutoRp_Active_CandidateRp) GetBundleName() string { return "cisco_ios_xr" }

func (candidateRp *AutoRp_Active_CandidateRp) GetYangName() string { return "candidate-rp" }

func (candidateRp *AutoRp_Active_CandidateRp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (candidateRp *AutoRp_Active_CandidateRp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (candidateRp *AutoRp_Active_CandidateRp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (candidateRp *AutoRp_Active_CandidateRp) SetParent(parent types.Entity) { candidateRp.parent = parent }

func (candidateRp *AutoRp_Active_CandidateRp) GetParent() types.Entity { return candidateRp.parent }

func (candidateRp *AutoRp_Active_CandidateRp) GetParentYangName() string { return "active" }

// AutoRp_Active_CandidateRp_Traffic
// AutoRP Candidate Traffic Counters
type AutoRp_Active_CandidateRp_Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of packets sent in active role. The type is interface{} with range:
    // 0..4294967295.
    ActiveSentPackets interface{}

    // Number of packets dropped in send path in standby role. The type is
    // interface{} with range: 0..4294967295.
    StandbySentPackets interface{}
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *AutoRp_Active_CandidateRp_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetGoName(yname string) string {
    if yname == "active-sent-packets" { return "ActiveSentPackets" }
    if yname == "standby-sent-packets" { return "StandbySentPackets" }
    return ""
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-sent-packets"] = traffic.ActiveSentPackets
    leafs["standby-sent-packets"] = traffic.StandbySentPackets
    return leafs
}

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetYangName() string { return "traffic" }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *AutoRp_Active_CandidateRp_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *AutoRp_Active_CandidateRp_Traffic) GetParentYangName() string { return "candidate-rp" }

// AutoRp_Active_CandidateRp_Rps
// AutoRP Candidate RP Table
type AutoRp_Active_CandidateRp_Rps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Candidate RP Entry. The type is slice of
    // AutoRp_Active_CandidateRp_Rps_Rp.
    Rp []AutoRp_Active_CandidateRp_Rps_Rp
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetFilter() yfilter.YFilter { return rps.YFilter }

func (rps *AutoRp_Active_CandidateRp_Rps) SetFilter(yf yfilter.YFilter) { rps.YFilter = yf }

func (rps *AutoRp_Active_CandidateRp_Rps) GetGoName(yname string) string {
    if yname == "rp" { return "Rp" }
    return ""
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetSegmentPath() string {
    return "rps"
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rp" {
        for _, c := range rps.Rp {
            if rps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Active_CandidateRp_Rps_Rp{}
        rps.Rp = append(rps.Rp, child)
        return &rps.Rp[len(rps.Rp)-1]
    }
    return nil
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rps.Rp {
        children[rps.Rp[i].GetSegmentPath()] = &rps.Rp[i]
    }
    return children
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rps *AutoRp_Active_CandidateRp_Rps) GetBundleName() string { return "cisco_ios_xr" }

func (rps *AutoRp_Active_CandidateRp_Rps) GetYangName() string { return "rps" }

func (rps *AutoRp_Active_CandidateRp_Rps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rps *AutoRp_Active_CandidateRp_Rps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rps *AutoRp_Active_CandidateRp_Rps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rps *AutoRp_Active_CandidateRp_Rps) SetParent(parent types.Entity) { rps.parent = parent }

func (rps *AutoRp_Active_CandidateRp_Rps) GetParent() types.Entity { return rps.parent }

func (rps *AutoRp_Active_CandidateRp_Rps) GetParentYangName() string { return "candidate-rp" }

// AutoRp_Active_CandidateRp_Rps_Rp
// AutoRP Candidate RP Entry
type AutoRp_Active_CandidateRp_Rps_Rp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Protocol Mode. The type is AutoRpProtocolMode.
    ProtocolMode interface{}

    // ACL Name. The type is string.
    AccessListName interface{}

    // Candidate RP IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CandidateRpAddress interface{}

    // TTL. The type is interface{} with range: -2147483648..2147483647.
    Ttl interface{}

    // Announce Period. The type is interface{} with range:
    // -2147483648..2147483647.
    AnnouncePeriod interface{}

    // Protocol Mode. The type is AutorpProtocolMode.
    ProtocolModeXr interface{}
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetFilter() yfilter.YFilter { return rp.YFilter }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) SetFilter(yf yfilter.YFilter) { rp.YFilter = yf }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "protocol-mode" { return "ProtocolMode" }
    if yname == "access-list-name" { return "AccessListName" }
    if yname == "candidate-rp-address" { return "CandidateRpAddress" }
    if yname == "ttl" { return "Ttl" }
    if yname == "announce-period" { return "AnnouncePeriod" }
    if yname == "protocol-mode-xr" { return "ProtocolModeXr" }
    return ""
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetSegmentPath() string {
    return "rp"
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = rp.InterfaceName
    leafs["protocol-mode"] = rp.ProtocolMode
    leafs["access-list-name"] = rp.AccessListName
    leafs["candidate-rp-address"] = rp.CandidateRpAddress
    leafs["ttl"] = rp.Ttl
    leafs["announce-period"] = rp.AnnouncePeriod
    leafs["protocol-mode-xr"] = rp.ProtocolModeXr
    return leafs
}

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetBundleName() string { return "cisco_ios_xr" }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetYangName() string { return "rp" }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) SetParent(parent types.Entity) { rp.parent = parent }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetParent() types.Entity { return rp.parent }

func (rp *AutoRp_Active_CandidateRp_Rps_Rp) GetParentYangName() string { return "rps" }

// AutoRp_Active_MappingAgent
// AutoRP Mapping Agent Table
type AutoRp_Active_MappingAgent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Traffic Counters.
    Traffic AutoRp_Active_MappingAgent_Traffic

    // AutoRP Mapping Agent Table Entries.
    RpAddresses AutoRp_Active_MappingAgent_RpAddresses

    // AutoRP Mapping Agent Summary Information.
    Summary AutoRp_Active_MappingAgent_Summary
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetFilter() yfilter.YFilter { return mappingAgent.YFilter }

func (mappingAgent *AutoRp_Active_MappingAgent) SetFilter(yf yfilter.YFilter) { mappingAgent.YFilter = yf }

func (mappingAgent *AutoRp_Active_MappingAgent) GetGoName(yname string) string {
    if yname == "traffic" { return "Traffic" }
    if yname == "rp-addresses" { return "RpAddresses" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetSegmentPath() string {
    return "mapping-agent"
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "traffic" {
        return &mappingAgent.Traffic
    }
    if childYangName == "rp-addresses" {
        return &mappingAgent.RpAddresses
    }
    if childYangName == "summary" {
        return &mappingAgent.Summary
    }
    return nil
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["traffic"] = &mappingAgent.Traffic
    children["rp-addresses"] = &mappingAgent.RpAddresses
    children["summary"] = &mappingAgent.Summary
    return children
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mappingAgent *AutoRp_Active_MappingAgent) GetBundleName() string { return "cisco_ios_xr" }

func (mappingAgent *AutoRp_Active_MappingAgent) GetYangName() string { return "mapping-agent" }

func (mappingAgent *AutoRp_Active_MappingAgent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mappingAgent *AutoRp_Active_MappingAgent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mappingAgent *AutoRp_Active_MappingAgent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mappingAgent *AutoRp_Active_MappingAgent) SetParent(parent types.Entity) { mappingAgent.parent = parent }

func (mappingAgent *AutoRp_Active_MappingAgent) GetParent() types.Entity { return mappingAgent.parent }

func (mappingAgent *AutoRp_Active_MappingAgent) GetParentYangName() string { return "active" }

// AutoRp_Active_MappingAgent_Traffic
// AutoRP Mapping Agent Traffic Counters
type AutoRp_Active_MappingAgent_Traffic struct {
    parent types.Entity
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

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetFilter() yfilter.YFilter { return traffic.YFilter }

func (traffic *AutoRp_Active_MappingAgent_Traffic) SetFilter(yf yfilter.YFilter) { traffic.YFilter = yf }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetGoName(yname string) string {
    if yname == "active-sent-packets" { return "ActiveSentPackets" }
    if yname == "standby-sent-packets" { return "StandbySentPackets" }
    if yname == "active-received-packets" { return "ActiveReceivedPackets" }
    if yname == "standby-received-packets" { return "StandbyReceivedPackets" }
    return ""
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetSegmentPath() string {
    return "traffic"
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["active-sent-packets"] = traffic.ActiveSentPackets
    leafs["standby-sent-packets"] = traffic.StandbySentPackets
    leafs["active-received-packets"] = traffic.ActiveReceivedPackets
    leafs["standby-received-packets"] = traffic.StandbyReceivedPackets
    return leafs
}

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetYangName() string { return "traffic" }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (traffic *AutoRp_Active_MappingAgent_Traffic) SetParent(parent types.Entity) { traffic.parent = parent }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetParent() types.Entity { return traffic.parent }

func (traffic *AutoRp_Active_MappingAgent_Traffic) GetParentYangName() string { return "mapping-agent" }

// AutoRp_Active_MappingAgent_RpAddresses
// AutoRP Mapping Agent Table Entries
type AutoRp_Active_MappingAgent_RpAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AutoRP Mapping Agent Entry. The type is slice of
    // AutoRp_Active_MappingAgent_RpAddresses_RpAddress.
    RpAddress []AutoRp_Active_MappingAgent_RpAddresses_RpAddress
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetFilter() yfilter.YFilter { return rpAddresses.YFilter }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) SetFilter(yf yfilter.YFilter) { rpAddresses.YFilter = yf }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    return ""
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetSegmentPath() string {
    return "rp-addresses"
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rp-address" {
        for _, c := range rpAddresses.RpAddress {
            if rpAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Active_MappingAgent_RpAddresses_RpAddress{}
        rpAddresses.RpAddress = append(rpAddresses.RpAddress, child)
        return &rpAddresses.RpAddress[len(rpAddresses.RpAddress)-1]
    }
    return nil
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rpAddresses.RpAddress {
        children[rpAddresses.RpAddress[i].GetSegmentPath()] = &rpAddresses.RpAddress[i]
    }
    return children
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetYangName() string { return "rp-addresses" }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) SetParent(parent types.Entity) { rpAddresses.parent = parent }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetParent() types.Entity { return rpAddresses.parent }

func (rpAddresses *AutoRp_Active_MappingAgent_RpAddresses) GetParentYangName() string { return "mapping-agent" }

// AutoRp_Active_MappingAgent_RpAddresses_RpAddress
// AutoRP Mapping Agent Entry
type AutoRp_Active_MappingAgent_RpAddresses_RpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. RP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpAddress interface{}

    // Candidate-RP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RpAddressXr interface{}

    // Time for expiration in seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    ExpiryTime interface{}

    // PIM version of the CRP. The type is interface{} with range: 0..255.
    PimVersion interface{}

    // Array of ranges. The type is slice of
    // AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range.
    Range []AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetFilter() yfilter.YFilter { return rpAddress.YFilter }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) SetFilter(yf yfilter.YFilter) { rpAddress.YFilter = yf }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetGoName(yname string) string {
    if yname == "rp-address" { return "RpAddress" }
    if yname == "rp-address-xr" { return "RpAddressXr" }
    if yname == "expiry-time" { return "ExpiryTime" }
    if yname == "pim-version" { return "PimVersion" }
    if yname == "range" { return "Range" }
    return ""
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetSegmentPath() string {
    return "rp-address" + "[rp-address='" + fmt.Sprintf("%v", rpAddress.RpAddress) + "']"
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "range" {
        for _, c := range rpAddress.Range {
            if rpAddress.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range{}
        rpAddress.Range = append(rpAddress.Range, child)
        return &rpAddress.Range[len(rpAddress.Range)-1]
    }
    return nil
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rpAddress.Range {
        children[rpAddress.Range[i].GetSegmentPath()] = &rpAddress.Range[i]
    }
    return children
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rp-address"] = rpAddress.RpAddress
    leafs["rp-address-xr"] = rpAddress.RpAddressXr
    leafs["expiry-time"] = rpAddress.ExpiryTime
    leafs["pim-version"] = rpAddress.PimVersion
    return leafs
}

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetYangName() string { return "rp-address" }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) SetParent(parent types.Entity) { rpAddress.parent = parent }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetParent() types.Entity { return rpAddress.parent }

func (rpAddress *AutoRp_Active_MappingAgent_RpAddresses_RpAddress) GetParentYangName() string { return "rp-addresses" }

// AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range
// Array of ranges
type AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix of the range. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetGoName(yname string) string {
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "protocol-mode" { return "ProtocolMode" }
    if yname == "is-advertised" { return "IsAdvertised" }
    if yname == "create-type" { return "CreateType" }
    if yname == "check-point-object-id" { return "CheckPointObjectId" }
    if yname == "uptime" { return "Uptime" }
    return ""
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetSegmentPath() string {
    return "range"
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix"] = self.Prefix
    leafs["prefix-length"] = self.PrefixLength
    leafs["protocol-mode"] = self.ProtocolMode
    leafs["is-advertised"] = self.IsAdvertised
    leafs["create-type"] = self.CreateType
    leafs["check-point-object-id"] = self.CheckPointObjectId
    leafs["uptime"] = self.Uptime
    return leafs
}

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetBundleName() string { return "cisco_ios_xr" }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetYangName() string { return "range" }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) SetParent(parent types.Entity) { self.parent = parent }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetParent() types.Entity { return self.parent }

func (self *AutoRp_Active_MappingAgent_RpAddresses_RpAddress_Range) GetParentYangName() string { return "rp-address" }

// AutoRp_Active_MappingAgent_Summary
// AutoRP Mapping Agent Summary Information
type AutoRp_Active_MappingAgent_Summary struct {
    parent types.Entity
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

func (summary *AutoRp_Active_MappingAgent_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *AutoRp_Active_MappingAgent_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *AutoRp_Active_MappingAgent_Summary) GetGoName(yname string) string {
    if yname == "is-maximum-disabled" { return "IsMaximumDisabled" }
    if yname == "cache-limit" { return "CacheLimit" }
    if yname == "cache-count" { return "CacheCount" }
    return ""
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-maximum-disabled"] = summary.IsMaximumDisabled
    leafs["cache-limit"] = summary.CacheLimit
    leafs["cache-count"] = summary.CacheCount
    return leafs
}

func (summary *AutoRp_Active_MappingAgent_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *AutoRp_Active_MappingAgent_Summary) GetYangName() string { return "summary" }

func (summary *AutoRp_Active_MappingAgent_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *AutoRp_Active_MappingAgent_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *AutoRp_Active_MappingAgent_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *AutoRp_Active_MappingAgent_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *AutoRp_Active_MappingAgent_Summary) GetParent() types.Entity { return summary.parent }

func (summary *AutoRp_Active_MappingAgent_Summary) GetParentYangName() string { return "mapping-agent" }

