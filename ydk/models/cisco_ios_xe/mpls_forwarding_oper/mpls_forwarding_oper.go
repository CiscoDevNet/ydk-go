// This module contains a collection of YANG definitions
// for mpls forwarding operational data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package mpls_forwarding_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_forwarding_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-forwarding-oper mpls-forwarding-oper-data}", reflect.TypeOf(MplsForwardingOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data", reflect.TypeOf(MplsForwardingOperData{}))
}

// ForwardingNextHopType represents Describes supported next hop types
type ForwardingNextHopType string

const (
    // Next hop is via a P2P link
    ForwardingNextHopType_point2point ForwardingNextHopType = "point2point"

    // Next hop IP address
    ForwardingNextHopType_next_hop_ip_address ForwardingNextHopType = "next-hop-ip-address"
)

// OutgoingInterfaceDescriptionType represents Describes outgoing interface information types
type OutgoingInterfaceDescriptionType string

const (
    // Forwarding path's outgoing interface type
    OutgoingInterfaceDescriptionType_interface_type OutgoingInterfaceDescriptionType = "interface-type"

    // Forwarding path's outgoing interface value
    OutgoingInterfaceDescriptionType_interface_value OutgoingInterfaceDescriptionType = "interface-value"
)

// ConnectionInfoType represents Describes connection information types
type ConnectionInfoType string

const (
    // Connection information string
    ConnectionInfoType_information_string ConnectionInfoType = "information-string"

    // IP prefix for the connection
    ConnectionInfoType_ip_prefix ConnectionInfoType = "ip-prefix"
)

// OutgoingInterfaceType represents Describes supported outgoing interface types
type OutgoingInterfaceType string

const (
    // Outgoing interface is not found
    OutgoingInterfaceType_drop OutgoingInterfaceType = "drop"

    // Punt packets to cpu
    OutgoingInterfaceType_punt OutgoingInterfaceType = "punt"

    // Aggregate interface (e.g., used for VPN label allocation)
    OutgoingInterfaceType_aggregate OutgoingInterfaceType = "aggregate"

    // Exception
    OutgoingInterfaceType_exception OutgoingInterfaceType = "exception"

    // No outgoing interface
    OutgoingInterfaceType_none OutgoingInterfaceType = "none"
)

// OutgoingLabelType represents Describes supported outgoing label types
type OutgoingLabelType string

const (
    // A label is not present
    OutgoingLabelType_no_label OutgoingLabelType = "no-label"

    // Label is poped
    OutgoingLabelType_pop_label OutgoingLabelType = "pop-label"

    // IPv4 explicit NULL label is present
    OutgoingLabelType_ipv4_explicit_null OutgoingLabelType = "ipv4-explicit-null"

    // IPv6 explicit NULL label is present
    OutgoingLabelType_ipv6_explicit_null OutgoingLabelType = "ipv6-explicit-null"

    // A regular MPLS label is present
    OutgoingLabelType_regular_label OutgoingLabelType = "regular-label"

    // An invalid label is present
    OutgoingLabelType_invalid_label OutgoingLabelType = "invalid-label"
)

// MplsForwardingOperData
// MPLS forwarding operatinal data
type MplsForwardingOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local label information. The type is slice of
    // MplsForwardingOperData_MplsLocalLabel.
    MplsLocalLabel []*MplsForwardingOperData_MplsLocalLabel

    // Statistics for forwarding paths of the local labels. The type is slice of
    // MplsForwardingOperData_MplsLocalLabelStatistics.
    MplsLocalLabelStatistics []*MplsForwardingOperData_MplsLocalLabelStatistics
}

func (mplsForwardingOperData *MplsForwardingOperData) GetEntityData() *types.CommonEntityData {
    mplsForwardingOperData.EntityData.YFilter = mplsForwardingOperData.YFilter
    mplsForwardingOperData.EntityData.YangName = "mpls-forwarding-oper-data"
    mplsForwardingOperData.EntityData.BundleName = "cisco_ios_xe"
    mplsForwardingOperData.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-forwarding-oper"
    mplsForwardingOperData.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data"
    mplsForwardingOperData.EntityData.AbsolutePath = mplsForwardingOperData.EntityData.SegmentPath
    mplsForwardingOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsForwardingOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsForwardingOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsForwardingOperData.EntityData.Children = types.NewOrderedMap()
    mplsForwardingOperData.EntityData.Children.Append("mpls-local-label", types.YChild{"MplsLocalLabel", nil})
    for i := range mplsForwardingOperData.MplsLocalLabel {
        mplsForwardingOperData.EntityData.Children.Append(types.GetSegmentPath(mplsForwardingOperData.MplsLocalLabel[i]), types.YChild{"MplsLocalLabel", mplsForwardingOperData.MplsLocalLabel[i]})
    }
    mplsForwardingOperData.EntityData.Children.Append("mpls-local-label-statistics", types.YChild{"MplsLocalLabelStatistics", nil})
    for i := range mplsForwardingOperData.MplsLocalLabelStatistics {
        mplsForwardingOperData.EntityData.Children.Append(types.GetSegmentPath(mplsForwardingOperData.MplsLocalLabelStatistics[i]), types.YChild{"MplsLocalLabelStatistics", mplsForwardingOperData.MplsLocalLabelStatistics[i]})
    }
    mplsForwardingOperData.EntityData.Leafs = types.NewOrderedMap()

    mplsForwardingOperData.EntityData.YListKeys = []string {}

    return &(mplsForwardingOperData.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel
// Local label information
type MplsForwardingOperData_MplsLocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Value of local label. The type is interface{} with
    // range: 0..4294967295.
    LocalLabel interface{}

    // Connection information for the local label.
    ConnectionInformation MplsForwardingOperData_MplsLocalLabel_ConnectionInformation

    // ECMP paths for the local label. The type is slice of
    // MplsForwardingOperData_MplsLocalLabel_ForwardingPaths.
    ForwardingPaths []*MplsForwardingOperData_MplsLocalLabel_ForwardingPaths
}

func (mplsLocalLabel *MplsForwardingOperData_MplsLocalLabel) GetEntityData() *types.CommonEntityData {
    mplsLocalLabel.EntityData.YFilter = mplsLocalLabel.YFilter
    mplsLocalLabel.EntityData.YangName = "mpls-local-label"
    mplsLocalLabel.EntityData.BundleName = "cisco_ios_xe"
    mplsLocalLabel.EntityData.ParentYangName = "mpls-forwarding-oper-data"
    mplsLocalLabel.EntityData.SegmentPath = "mpls-local-label" + types.AddKeyToken(mplsLocalLabel.LocalLabel, "local-label")
    mplsLocalLabel.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/" + mplsLocalLabel.EntityData.SegmentPath
    mplsLocalLabel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLocalLabel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLocalLabel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLocalLabel.EntityData.Children = types.NewOrderedMap()
    mplsLocalLabel.EntityData.Children.Append("connection-information", types.YChild{"ConnectionInformation", &mplsLocalLabel.ConnectionInformation})
    mplsLocalLabel.EntityData.Children.Append("forwarding-paths", types.YChild{"ForwardingPaths", nil})
    for i := range mplsLocalLabel.ForwardingPaths {
        types.SetYListKey(mplsLocalLabel.ForwardingPaths[i], i)
        mplsLocalLabel.EntityData.Children.Append(types.GetSegmentPath(mplsLocalLabel.ForwardingPaths[i]), types.YChild{"ForwardingPaths", mplsLocalLabel.ForwardingPaths[i]})
    }
    mplsLocalLabel.EntityData.Leafs = types.NewOrderedMap()
    mplsLocalLabel.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", mplsLocalLabel.LocalLabel})

    mplsLocalLabel.EntityData.YListKeys = []string {"LocalLabel"}

    return &(mplsLocalLabel.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel_ConnectionInformation
// Connection information for the local label
type MplsForwardingOperData_MplsLocalLabel_ConnectionInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Connection information string. The type is string.
    InfoStr interface{}

    // IP prefix for the connection. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    IpPrefix interface{}
}

func (connectionInformation *MplsForwardingOperData_MplsLocalLabel_ConnectionInformation) GetEntityData() *types.CommonEntityData {
    connectionInformation.EntityData.YFilter = connectionInformation.YFilter
    connectionInformation.EntityData.YangName = "connection-information"
    connectionInformation.EntityData.BundleName = "cisco_ios_xe"
    connectionInformation.EntityData.ParentYangName = "mpls-local-label"
    connectionInformation.EntityData.SegmentPath = "connection-information"
    connectionInformation.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/mpls-local-label/" + connectionInformation.EntityData.SegmentPath
    connectionInformation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    connectionInformation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    connectionInformation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    connectionInformation.EntityData.Children = types.NewOrderedMap()
    connectionInformation.EntityData.Leafs = types.NewOrderedMap()
    connectionInformation.EntityData.Leafs.Append("info-str", types.YLeaf{"InfoStr", connectionInformation.InfoStr})
    connectionInformation.EntityData.Leafs.Append("ip-prefix", types.YLeaf{"IpPrefix", connectionInformation.IpPrefix})

    connectionInformation.EntityData.YListKeys = []string {}

    return &(connectionInformation.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel_ForwardingPaths
// ECMP paths for the local label
type MplsForwardingOperData_MplsLocalLabel_ForwardingPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Outgoing interface information.
    OutgoingInterface MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingInterface

    // Outgoing label information.
    OutgoingLabel MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingLabel

    // Next hop information.
    NextHop MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_NextHop
}

func (forwardingPaths *MplsForwardingOperData_MplsLocalLabel_ForwardingPaths) GetEntityData() *types.CommonEntityData {
    forwardingPaths.EntityData.YFilter = forwardingPaths.YFilter
    forwardingPaths.EntityData.YangName = "forwarding-paths"
    forwardingPaths.EntityData.BundleName = "cisco_ios_xe"
    forwardingPaths.EntityData.ParentYangName = "mpls-local-label"
    forwardingPaths.EntityData.SegmentPath = "forwarding-paths" + types.AddNoKeyToken(forwardingPaths)
    forwardingPaths.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/mpls-local-label/" + forwardingPaths.EntityData.SegmentPath
    forwardingPaths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingPaths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingPaths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingPaths.EntityData.Children = types.NewOrderedMap()
    forwardingPaths.EntityData.Children.Append("outgoing-interface", types.YChild{"OutgoingInterface", &forwardingPaths.OutgoingInterface})
    forwardingPaths.EntityData.Children.Append("outgoing-label", types.YChild{"OutgoingLabel", &forwardingPaths.OutgoingLabel})
    forwardingPaths.EntityData.Children.Append("next-hop", types.YChild{"NextHop", &forwardingPaths.NextHop})
    forwardingPaths.EntityData.Leafs = types.NewOrderedMap()

    forwardingPaths.EntityData.YListKeys = []string {}

    return &(forwardingPaths.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingInterface
// Outgoing interface information
type MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outgoing interface type. The type is OutgoingInterfaceType.
    InterfaceType interface{}

    // Outgoing interface value. The type is string.
    InterfaceValue interface{}
}

func (outgoingInterface *MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingInterface) GetEntityData() *types.CommonEntityData {
    outgoingInterface.EntityData.YFilter = outgoingInterface.YFilter
    outgoingInterface.EntityData.YangName = "outgoing-interface"
    outgoingInterface.EntityData.BundleName = "cisco_ios_xe"
    outgoingInterface.EntityData.ParentYangName = "forwarding-paths"
    outgoingInterface.EntityData.SegmentPath = "outgoing-interface"
    outgoingInterface.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/mpls-local-label/forwarding-paths/" + outgoingInterface.EntityData.SegmentPath
    outgoingInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    outgoingInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    outgoingInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    outgoingInterface.EntityData.Children = types.NewOrderedMap()
    outgoingInterface.EntityData.Leafs = types.NewOrderedMap()
    outgoingInterface.EntityData.Leafs.Append("interface-type", types.YLeaf{"InterfaceType", outgoingInterface.InterfaceType})
    outgoingInterface.EntityData.Leafs.Append("interface-value", types.YLeaf{"InterfaceValue", outgoingInterface.InterfaceValue})

    outgoingInterface.EntityData.YListKeys = []string {}

    return &(outgoingInterface.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingLabel
// Outgoing label information
type MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True if a label is not present. The type is bool.
    NoLabel interface{}

    // Pop label value. The type is interface{} with range: 0..4294967295.
    PopLabel interface{}

    // IPv4 explicit null label value. The type is interface{} with range:
    // 0..4294967295.
    Ipv4ExplicitNull interface{}

    // IPv6 explicit null label value. The type is interface{} with range:
    // 0..4294967295.
    Ipv6ExplicitNull interface{}

    // Regular label value. The type is interface{} with range: 0..4294967295.
    RegularLabel interface{}

    // True if a label with a value that is outside IETF acceptable label range is
    // present. The type is bool.
    InvalidLabel interface{}
}

func (outgoingLabel *MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_OutgoingLabel) GetEntityData() *types.CommonEntityData {
    outgoingLabel.EntityData.YFilter = outgoingLabel.YFilter
    outgoingLabel.EntityData.YangName = "outgoing-label"
    outgoingLabel.EntityData.BundleName = "cisco_ios_xe"
    outgoingLabel.EntityData.ParentYangName = "forwarding-paths"
    outgoingLabel.EntityData.SegmentPath = "outgoing-label"
    outgoingLabel.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/mpls-local-label/forwarding-paths/" + outgoingLabel.EntityData.SegmentPath
    outgoingLabel.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    outgoingLabel.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    outgoingLabel.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    outgoingLabel.EntityData.Children = types.NewOrderedMap()
    outgoingLabel.EntityData.Leafs = types.NewOrderedMap()
    outgoingLabel.EntityData.Leafs.Append("no-label", types.YLeaf{"NoLabel", outgoingLabel.NoLabel})
    outgoingLabel.EntityData.Leafs.Append("pop-label", types.YLeaf{"PopLabel", outgoingLabel.PopLabel})
    outgoingLabel.EntityData.Leafs.Append("ipv4-explicit-null", types.YLeaf{"Ipv4ExplicitNull", outgoingLabel.Ipv4ExplicitNull})
    outgoingLabel.EntityData.Leafs.Append("ipv6-explicit-null", types.YLeaf{"Ipv6ExplicitNull", outgoingLabel.Ipv6ExplicitNull})
    outgoingLabel.EntityData.Leafs.Append("regular-label", types.YLeaf{"RegularLabel", outgoingLabel.RegularLabel})
    outgoingLabel.EntityData.Leafs.Append("invalid-label", types.YLeaf{"InvalidLabel", outgoingLabel.InvalidLabel})

    outgoingLabel.EntityData.YListKeys = []string {}

    return &(outgoingLabel.EntityData)
}

// MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_NextHop
// Next hop information
type MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_NextHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True if next hop is via a p2p link. The type is bool.
    P2p interface{}

    // Next hop ip address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (nextHop *MplsForwardingOperData_MplsLocalLabel_ForwardingPaths_NextHop) GetEntityData() *types.CommonEntityData {
    nextHop.EntityData.YFilter = nextHop.YFilter
    nextHop.EntityData.YangName = "next-hop"
    nextHop.EntityData.BundleName = "cisco_ios_xe"
    nextHop.EntityData.ParentYangName = "forwarding-paths"
    nextHop.EntityData.SegmentPath = "next-hop"
    nextHop.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/mpls-local-label/forwarding-paths/" + nextHop.EntityData.SegmentPath
    nextHop.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nextHop.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nextHop.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nextHop.EntityData.Children = types.NewOrderedMap()
    nextHop.EntityData.Leafs = types.NewOrderedMap()
    nextHop.EntityData.Leafs.Append("p2p", types.YLeaf{"P2p", nextHop.P2p})
    nextHop.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", nextHop.IpAddress})

    nextHop.EntityData.YListKeys = []string {}

    return &(nextHop.EntityData)
}

// MplsForwardingOperData_MplsLocalLabelStatistics
// Statistics for forwarding paths of the local labels
type MplsForwardingOperData_MplsLocalLabelStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Value of local label. The type is interface{} with
    // range: 0..4294967295.
    LocalLabel interface{}

    // This attribute is a key. Path index among ECMP paths for the same local
    // label. The type is interface{} with range: 0..255.
    ForwardingPathIndex interface{}

    // Number of bytes switched to a particular ECMP path. The type is interface{}
    // with range: 0..18446744073709551615.
    LabelSwitchedBytes interface{}

    // Number of packets switched to a particular ECMP path. The type is
    // interface{} with range: 0..18446744073709551615.
    LabelSwitchedPackets interface{}
}

func (mplsLocalLabelStatistics *MplsForwardingOperData_MplsLocalLabelStatistics) GetEntityData() *types.CommonEntityData {
    mplsLocalLabelStatistics.EntityData.YFilter = mplsLocalLabelStatistics.YFilter
    mplsLocalLabelStatistics.EntityData.YangName = "mpls-local-label-statistics"
    mplsLocalLabelStatistics.EntityData.BundleName = "cisco_ios_xe"
    mplsLocalLabelStatistics.EntityData.ParentYangName = "mpls-forwarding-oper-data"
    mplsLocalLabelStatistics.EntityData.SegmentPath = "mpls-local-label-statistics" + types.AddKeyToken(mplsLocalLabelStatistics.LocalLabel, "local-label") + types.AddKeyToken(mplsLocalLabelStatistics.ForwardingPathIndex, "forwarding-path-index")
    mplsLocalLabelStatistics.EntityData.AbsolutePath = "Cisco-IOS-XE-mpls-forwarding-oper:mpls-forwarding-oper-data/" + mplsLocalLabelStatistics.EntityData.SegmentPath
    mplsLocalLabelStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLocalLabelStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLocalLabelStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLocalLabelStatistics.EntityData.Children = types.NewOrderedMap()
    mplsLocalLabelStatistics.EntityData.Leafs = types.NewOrderedMap()
    mplsLocalLabelStatistics.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", mplsLocalLabelStatistics.LocalLabel})
    mplsLocalLabelStatistics.EntityData.Leafs.Append("forwarding-path-index", types.YLeaf{"ForwardingPathIndex", mplsLocalLabelStatistics.ForwardingPathIndex})
    mplsLocalLabelStatistics.EntityData.Leafs.Append("label-switched-bytes", types.YLeaf{"LabelSwitchedBytes", mplsLocalLabelStatistics.LabelSwitchedBytes})
    mplsLocalLabelStatistics.EntityData.Leafs.Append("label-switched-packets", types.YLeaf{"LabelSwitchedPackets", mplsLocalLabelStatistics.LabelSwitchedPackets})

    mplsLocalLabelStatistics.EntityData.YListKeys = []string {"LocalLabel", "ForwardingPathIndex"}

    return &(mplsLocalLabelStatistics.EntityData)
}

