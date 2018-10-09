// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-ma package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-network-global: IPv4 network global configuration data
//   subscriber-pta: subscriber pta
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

func (ipv4NetworkGlobal *Ipv4NetworkGlobal) GetEntityData() *types.CommonEntityData {
    ipv4NetworkGlobal.EntityData.YFilter = ipv4NetworkGlobal.YFilter
    ipv4NetworkGlobal.EntityData.YangName = "ipv4-network-global"
    ipv4NetworkGlobal.EntityData.BundleName = "cisco_ios_xr"
    ipv4NetworkGlobal.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ma-cfg"
    ipv4NetworkGlobal.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-cfg:ipv4-network-global"
    ipv4NetworkGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4NetworkGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4NetworkGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4NetworkGlobal.EntityData.Children = types.NewOrderedMap()
    ipv4NetworkGlobal.EntityData.Children.Append("unnumbered", types.YChild{"Unnumbered", &ipv4NetworkGlobal.Unnumbered})
    ipv4NetworkGlobal.EntityData.Children.Append("qppb", types.YChild{"Qppb", &ipv4NetworkGlobal.Qppb})
    ipv4NetworkGlobal.EntityData.Leafs = types.NewOrderedMap()
    ipv4NetworkGlobal.EntityData.Leafs.Append("source-route", types.YLeaf{"SourceRoute", ipv4NetworkGlobal.SourceRoute})
    ipv4NetworkGlobal.EntityData.Leafs.Append("reassemble-max-packets", types.YLeaf{"ReassembleMaxPackets", ipv4NetworkGlobal.ReassembleMaxPackets})
    ipv4NetworkGlobal.EntityData.Leafs.Append("reassemble-time-out", types.YLeaf{"ReassembleTimeOut", ipv4NetworkGlobal.ReassembleTimeOut})

    ipv4NetworkGlobal.EntityData.YListKeys = []string {}

    return &(ipv4NetworkGlobal.EntityData)
}

// Ipv4NetworkGlobal_Unnumbered
// Enable IPv4 processing without an explicit
// address
type Ipv4NetworkGlobal_Unnumbered struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure MPLS routing protocol parameters.
    Mpls Ipv4NetworkGlobal_Unnumbered_Mpls
}

func (unnumbered *Ipv4NetworkGlobal_Unnumbered) GetEntityData() *types.CommonEntityData {
    unnumbered.EntityData.YFilter = unnumbered.YFilter
    unnumbered.EntityData.YangName = "unnumbered"
    unnumbered.EntityData.BundleName = "cisco_ios_xr"
    unnumbered.EntityData.ParentYangName = "ipv4-network-global"
    unnumbered.EntityData.SegmentPath = "unnumbered"
    unnumbered.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    unnumbered.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    unnumbered.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    unnumbered.EntityData.Children = types.NewOrderedMap()
    unnumbered.EntityData.Children.Append("mpls", types.YChild{"Mpls", &unnumbered.Mpls})
    unnumbered.EntityData.Leafs = types.NewOrderedMap()

    unnumbered.EntityData.YListKeys = []string {}

    return &(unnumbered.EntityData)
}

// Ipv4NetworkGlobal_Unnumbered_Mpls
// Configure MPLS routing protocol parameters
type Ipv4NetworkGlobal_Unnumbered_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 commands for MPLS Traffic Engineering.
    Te Ipv4NetworkGlobal_Unnumbered_Mpls_Te
}

func (mpls *Ipv4NetworkGlobal_Unnumbered_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xr"
    mpls.EntityData.ParentYangName = "unnumbered"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Children.Append("te", types.YChild{"Te", &mpls.Te})
    mpls.EntityData.Leafs = types.NewOrderedMap()

    mpls.EntityData.YListKeys = []string {}

    return &(mpls.EntityData)
}

// Ipv4NetworkGlobal_Unnumbered_Mpls_Te
// IPv4 commands for MPLS Traffic Engineering
type Ipv4NetworkGlobal_Unnumbered_Mpls_Te struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IP processing without an explicit address on MPLS Traffic-Eng. The
    // type is string.
    Interface interface{}
}

func (te *Ipv4NetworkGlobal_Unnumbered_Mpls_Te) GetEntityData() *types.CommonEntityData {
    te.EntityData.YFilter = te.YFilter
    te.EntityData.YangName = "te"
    te.EntityData.BundleName = "cisco_ios_xr"
    te.EntityData.ParentYangName = "mpls"
    te.EntityData.SegmentPath = "te"
    te.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    te.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    te.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    te.EntityData.Children = types.NewOrderedMap()
    te.EntityData.Leafs = types.NewOrderedMap()
    te.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", te.Interface})

    te.EntityData.YListKeys = []string {}

    return &(te.EntityData)
}

// Ipv4NetworkGlobal_Qppb
// QPPB
type Ipv4NetworkGlobal_Qppb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QPPB configuration on source. The type is Ipv4Qppb.
    Source interface{}

    // QPPB configuration on destination. The type is Ipv4Qppb.
    Destination interface{}
}

func (qppb *Ipv4NetworkGlobal_Qppb) GetEntityData() *types.CommonEntityData {
    qppb.EntityData.YFilter = qppb.YFilter
    qppb.EntityData.YangName = "qppb"
    qppb.EntityData.BundleName = "cisco_ios_xr"
    qppb.EntityData.ParentYangName = "ipv4-network-global"
    qppb.EntityData.SegmentPath = "qppb"
    qppb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qppb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qppb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qppb.EntityData.Children = types.NewOrderedMap()
    qppb.EntityData.Leafs = types.NewOrderedMap()
    qppb.EntityData.Leafs.Append("source", types.YLeaf{"Source", qppb.Source})
    qppb.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", qppb.Destination})

    qppb.EntityData.YListKeys = []string {}

    return &(qppb.EntityData)
}

// SubscriberPta
// subscriber pta
type SubscriberPta struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP MSS Adjust (bytes). The type is interface{} with range: 1280..1536.
    // Units are byte.
    TcpMssAdjust interface{}
}

func (subscriberPta *SubscriberPta) GetEntityData() *types.CommonEntityData {
    subscriberPta.EntityData.YFilter = subscriberPta.YFilter
    subscriberPta.EntityData.YangName = "subscriber-pta"
    subscriberPta.EntityData.BundleName = "cisco_ios_xr"
    subscriberPta.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-ma-cfg"
    subscriberPta.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ma-cfg:subscriber-pta"
    subscriberPta.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberPta.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberPta.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberPta.EntityData.Children = types.NewOrderedMap()
    subscriberPta.EntityData.Leafs = types.NewOrderedMap()
    subscriberPta.EntityData.Leafs.Append("tcp-mss-adjust", types.YLeaf{"TcpMssAdjust", subscriberPta.TcpMssAdjust})

    subscriberPta.EntityData.YListKeys = []string {}

    return &(subscriberPta.EntityData)
}

