// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-ma package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-network-global: IPv4 network global configuration data
//   subscriber-pta: subscriber pta
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_ma_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_ma_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ma-cfg ipv4-network-global}", reflect.TypeOf(Ipv4NetworkGlobal{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ma-cfg:ipv4-network-global", reflect.TypeOf(Ipv4NetworkGlobal{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-ma-cfg subscriber-pta}", reflect.TypeOf(SubscriberPta{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-ma-cfg:subscriber-pta", reflect.TypeOf(SubscriberPta{}))
}

// Ipv4Qppb represents Ipv4 qppb
type Ipv4Qppb string

const (
    // No QPPB configuration
    Ipv4Qppb_none Ipv4Qppb = "none"

    // Enable ip-precedence based QPPB
    Ipv4Qppb_ip_prec Ipv4Qppb = "ip-prec"

    // Enable qos-group based QPPB
    Ipv4Qppb_qos_grp Ipv4Qppb = "qos-grp"

    // Enable both ip-precedence and qos-group based
    // QPPB
    Ipv4Qppb_both Ipv4Qppb = "both"
)

// Ipv4NetworkGlobal
// IPv4 network global configuration data
type Ipv4NetworkGlobal struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The flag for enabling whether to process packets with source routing header
    // options. The type is bool. The default value is true.
    SourceRoute interface{}

    // Percentage of total packets available in the system. The type is
    // interface{} with range: 1..50. Units are percentage.
    ReassembleMaxPackets interface{}

    // Number of seconds a reassembly queue will hold before timeout. The type is
    // interface{} with range: 1..120. Units are second.
    ReassembleTimeOut interface{}

    // Enable IPv4 processing without an explicit address.
    Unnumbered Ipv4NetworkGlobal_Unnumbered

    // QPPB.
    Qppb Ipv4NetworkGlobal_Qppb
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetFilter() yfilter.YFilter { return ipv4NetworkGlobal.YFilter }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) SetFilter(yf yfilter.YFilter) { ipv4NetworkGlobal.YFilter = yf }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetGoName(yname string) string {
    if yname == "source-route" { return "SourceRoute" }
    if yname == "reassemble-max-packets" { return "ReassembleMaxPackets" }
    if yname == "reassemble-time-out" { return "ReassembleTimeOut" }
    if yname == "unnumbered" { return "Unnumbered" }
    if yname == "qppb" { return "Qppb" }
    return ""
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-cfg:ipv4-network-global"
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "unnumbered" {
        return &ipv4NetworkGlobal.Unnumbered
    }
    if childYangName == "qppb" {
        return &ipv4NetworkGlobal.Qppb
    }
    return nil
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["unnumbered"] = &ipv4NetworkGlobal.Unnumbered
    children["qppb"] = &ipv4NetworkGlobal.Qppb
    return children
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-route"] = ipv4NetworkGlobal.SourceRoute
    leafs["reassemble-max-packets"] = ipv4NetworkGlobal.ReassembleMaxPackets
    leafs["reassemble-time-out"] = ipv4NetworkGlobal.ReassembleTimeOut
    return leafs
}

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetYangName() string { return "ipv4-network-global" }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) SetParent(parent types.Entity) { ipv4NetworkGlobal.parent = parent }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetParent() types.Entity { return ipv4NetworkGlobal.parent }

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ma-cfg" }

// Ipv4NetworkGlobal_Unnumbered
// Enable IPv4 processing without an explicit
// address
type Ipv4NetworkGlobal_Unnumbered struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure MPLS routing protocol parameters.
    Mpls Ipv4NetworkGlobal_Unnumbered_Mpls
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetFilter() yfilter.YFilter { return unnumbered.YFilter }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) SetFilter(yf yfilter.YFilter) { unnumbered.YFilter = yf }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetGoName(yname string) string {
    if yname == "mpls" { return "Mpls" }
    return ""
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetSegmentPath() string {
    return "unnumbered"
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls" {
        return &unnumbered.Mpls
    }
    return nil
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls"] = &unnumbered.Mpls
    return children
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetBundleName() string { return "cisco_ios_xr" }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetYangName() string { return "unnumbered" }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) SetParent(parent types.Entity) { unnumbered.parent = parent }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetParent() types.Entity { return unnumbered.parent }

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetParentYangName() string { return "ipv4-network-global" }

// Ipv4NetworkGlobal_Unnumbered_Mpls
// Configure MPLS routing protocol parameters
type Ipv4NetworkGlobal_Unnumbered_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv4 commands for MPLS Traffic Engineering.
    Te Ipv4NetworkGlobal_Unnumbered_Mpls_Te
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetGoName(yname string) string {
    if yname == "te" { return "Te" }
    return ""
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "te" {
        return &mpls.Te
    }
    return nil
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["te"] = &mpls.Te
    return children
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetYangName() string { return "mpls" }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetParentYangName() string { return "unnumbered" }

// Ipv4NetworkGlobal_Unnumbered_Mpls_Te
// IPv4 commands for MPLS Traffic Engineering
type Ipv4NetworkGlobal_Unnumbered_Mpls_Te struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable IP processing without an explicit address on MPLS Traffic-Eng. The
    // type is string.
    Interface interface{}
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetFilter() yfilter.YFilter { return te.YFilter }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) SetFilter(yf yfilter.YFilter) { te.YFilter = yf }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetSegmentPath() string {
    return "te"
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = te.Interface
    return leafs
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetBundleName() string { return "cisco_ios_xr" }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetYangName() string { return "te" }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) SetParent(parent types.Entity) { te.parent = parent }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetParent() types.Entity { return te.parent }

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetParentYangName() string { return "mpls" }

// Ipv4NetworkGlobal_Qppb
// QPPB
type Ipv4NetworkGlobal_Qppb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QPPB configuration on source. The type is Ipv4Qppb.
    Source interface{}

    // QPPB configuration on destination. The type is Ipv4Qppb.
    Destination interface{}
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetFilter() yfilter.YFilter { return qppb.YFilter }

func (qppb *Ipv4NetworkGlobal_Qppb) SetFilter(yf yfilter.YFilter) { qppb.YFilter = yf }

func (qppb *Ipv4NetworkGlobal_Qppb) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    return ""
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetSegmentPath() string {
    return "qppb"
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source"] = qppb.Source
    leafs["destination"] = qppb.Destination
    return leafs
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetBundleName() string { return "cisco_ios_xr" }

func (qppb *Ipv4NetworkGlobal_Qppb) GetYangName() string { return "qppb" }

func (qppb *Ipv4NetworkGlobal_Qppb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qppb *Ipv4NetworkGlobal_Qppb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qppb *Ipv4NetworkGlobal_Qppb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qppb *Ipv4NetworkGlobal_Qppb) SetParent(parent types.Entity) { qppb.parent = parent }

func (qppb *Ipv4NetworkGlobal_Qppb) GetParent() types.Entity { return qppb.parent }

func (qppb *Ipv4NetworkGlobal_Qppb) GetParentYangName() string { return "ipv4-network-global" }

// SubscriberPta
// subscriber pta
type SubscriberPta struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP MSS Adjust (bytes). The type is interface{} with range: 1280..1536.
    // Units are byte.
    TcpMssAdjust interface{}
}

func (subscriberPta *SubscriberPta) GetFilter() yfilter.YFilter { return subscriberPta.YFilter }

func (subscriberPta *SubscriberPta) SetFilter(yf yfilter.YFilter) { subscriberPta.YFilter = yf }

func (subscriberPta *SubscriberPta) GetGoName(yname string) string {
    if yname == "tcp-mss-adjust" { return "TcpMssAdjust" }
    return ""
}

func (subscriberPta *SubscriberPta) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ma-cfg:subscriber-pta"
}

func (subscriberPta *SubscriberPta) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (subscriberPta *SubscriberPta) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (subscriberPta *SubscriberPta) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-mss-adjust"] = subscriberPta.TcpMssAdjust
    return leafs
}

func (subscriberPta *SubscriberPta) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberPta *SubscriberPta) GetYangName() string { return "subscriber-pta" }

func (subscriberPta *SubscriberPta) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberPta *SubscriberPta) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberPta *SubscriberPta) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberPta *SubscriberPta) SetParent(parent types.Entity) { subscriberPta.parent = parent }

func (subscriberPta *SubscriberPta) GetParent() types.Entity { return subscriberPta.parent }

func (subscriberPta *SubscriberPta) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-ma-cfg" }

