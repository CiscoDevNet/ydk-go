// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-lsd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-lsd: MPLS LSD configuration data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package mpls_lsd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_lsd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-lsd-cfg mpls-lsd}", reflect.TypeOf(MplsLsd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd", reflect.TypeOf(MplsLsd{}))
}

// MplsIpTtlPropagateDisable represents Mpls ip ttl propagate disable
type MplsIpTtlPropagateDisable string

const (
    // Disable IP TTL propagation for all MPLS packets
    MplsIpTtlPropagateDisable_all MplsIpTtlPropagateDisable = "all"

    // Disable IP TTL propagation for only forwarded
    // MPLS packets
    MplsIpTtlPropagateDisable_forward MplsIpTtlPropagateDisable = "forward"

    // Disable IP TTL propagation for only locally
    // generated MPLS packets
    MplsIpTtlPropagateDisable_local MplsIpTtlPropagateDisable = "local"
)

// MplsLsd
// MPLS LSD configuration data
type MplsLsd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Multiply the MPLS LSD Ltrace buffer length. The type is interface{} with
    // range: 2..5.
    LtraceMultiplier interface{}

    // Disable LSD application reg delay. The type is interface{}.
    AppRegDelayDisable interface{}

    // Enable MPLS Entropy Label. The type is interface{}.
    MplsEntropyLabel interface{}

    // Disable Propagation of IP TTL onto the label stack. The type is
    // MplsIpTtlPropagateDisable.
    MplsIpTtlPropagateDisable interface{}

    // Configure IPv6 parameters.
    Ipv6 MplsLsd_Ipv6

    // Configure IPv4 parameters.
    Ipv4 MplsLsd_Ipv4

    // Table of label databases.
    LabelDatabases MplsLsd_LabelDatabases
}

func (mplsLsd *MplsLsd) GetEntityData() *types.CommonEntityData {
    mplsLsd.EntityData.YFilter = mplsLsd.YFilter
    mplsLsd.EntityData.YangName = "mpls-lsd"
    mplsLsd.EntityData.BundleName = "cisco_ios_xr"
    mplsLsd.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-lsd-cfg"
    mplsLsd.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd"
    mplsLsd.EntityData.AbsolutePath = mplsLsd.EntityData.SegmentPath
    mplsLsd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLsd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLsd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLsd.EntityData.Children = types.NewOrderedMap()
    mplsLsd.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &mplsLsd.Ipv6})
    mplsLsd.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &mplsLsd.Ipv4})
    mplsLsd.EntityData.Children.Append("label-databases", types.YChild{"LabelDatabases", &mplsLsd.LabelDatabases})
    mplsLsd.EntityData.Leafs = types.NewOrderedMap()
    mplsLsd.EntityData.Leafs.Append("ltrace-multiplier", types.YLeaf{"LtraceMultiplier", mplsLsd.LtraceMultiplier})
    mplsLsd.EntityData.Leafs.Append("app-reg-delay-disable", types.YLeaf{"AppRegDelayDisable", mplsLsd.AppRegDelayDisable})
    mplsLsd.EntityData.Leafs.Append("mpls-entropy-label", types.YLeaf{"MplsEntropyLabel", mplsLsd.MplsEntropyLabel})
    mplsLsd.EntityData.Leafs.Append("mpls-ip-ttl-propagate-disable", types.YLeaf{"MplsIpTtlPropagateDisable", mplsLsd.MplsIpTtlPropagateDisable})

    mplsLsd.EntityData.YListKeys = []string {}

    return &(mplsLsd.EntityData)
}

// MplsLsd_Ipv6
// Configure IPv6 parameters
type MplsLsd_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of labels to pop upon MPLS IP TTL expiry. The type is interface{}
    // with range: 1..10.
    TtlExpirationPop interface{}
}

func (ipv6 *MplsLsd_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "mpls-lsd"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("ttl-expiration-pop", types.YLeaf{"TtlExpirationPop", ipv6.TtlExpirationPop})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// MplsLsd_Ipv4
// Configure IPv4 parameters
type MplsLsd_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of labels to pop upon MPLS IP TTL expiry. The type is interface{}
    // with range: 1..10.
    TtlExpirationPop interface{}
}

func (ipv4 *MplsLsd_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "mpls-lsd"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("ttl-expiration-pop", types.YLeaf{"TtlExpirationPop", ipv4.TtlExpirationPop})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// MplsLsd_LabelDatabases
// Table of label databases
type MplsLsd_LabelDatabases struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A label database. The type is slice of
    // MplsLsd_LabelDatabases_LabelDatabase.
    LabelDatabase []*MplsLsd_LabelDatabases_LabelDatabase
}

func (labelDatabases *MplsLsd_LabelDatabases) GetEntityData() *types.CommonEntityData {
    labelDatabases.EntityData.YFilter = labelDatabases.YFilter
    labelDatabases.EntityData.YangName = "label-databases"
    labelDatabases.EntityData.BundleName = "cisco_ios_xr"
    labelDatabases.EntityData.ParentYangName = "mpls-lsd"
    labelDatabases.EntityData.SegmentPath = "label-databases"
    labelDatabases.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd/" + labelDatabases.EntityData.SegmentPath
    labelDatabases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelDatabases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelDatabases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelDatabases.EntityData.Children = types.NewOrderedMap()
    labelDatabases.EntityData.Children.Append("label-database", types.YChild{"LabelDatabase", nil})
    for i := range labelDatabases.LabelDatabase {
        labelDatabases.EntityData.Children.Append(types.GetSegmentPath(labelDatabases.LabelDatabase[i]), types.YChild{"LabelDatabase", labelDatabases.LabelDatabase[i]})
    }
    labelDatabases.EntityData.Leafs = types.NewOrderedMap()

    labelDatabases.EntityData.YListKeys = []string {}

    return &(labelDatabases.EntityData)
}

// MplsLsd_LabelDatabases_LabelDatabase
// A label database
type MplsLsd_LabelDatabases_LabelDatabase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Label database identifier. The type is interface{}
    // with range: 0..0.
    LabelDatabaseId interface{}

    // Label range.
    LabelRange MplsLsd_LabelDatabases_LabelDatabase_LabelRange
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetEntityData() *types.CommonEntityData {
    labelDatabase.EntityData.YFilter = labelDatabase.YFilter
    labelDatabase.EntityData.YangName = "label-database"
    labelDatabase.EntityData.BundleName = "cisco_ios_xr"
    labelDatabase.EntityData.ParentYangName = "label-databases"
    labelDatabase.EntityData.SegmentPath = "label-database" + types.AddKeyToken(labelDatabase.LabelDatabaseId, "label-database-id")
    labelDatabase.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd/label-databases/" + labelDatabase.EntityData.SegmentPath
    labelDatabase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelDatabase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelDatabase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelDatabase.EntityData.Children = types.NewOrderedMap()
    labelDatabase.EntityData.Children.Append("label-range", types.YChild{"LabelRange", &labelDatabase.LabelRange})
    labelDatabase.EntityData.Leafs = types.NewOrderedMap()
    labelDatabase.EntityData.Leafs.Append("label-database-id", types.YLeaf{"LabelDatabaseId", labelDatabase.LabelDatabaseId})

    labelDatabase.EntityData.YListKeys = []string {"LabelDatabaseId"}

    return &(labelDatabase.EntityData)
}

// MplsLsd_LabelDatabases_LabelDatabase_LabelRange
// Label range
type MplsLsd_LabelDatabases_LabelDatabase_LabelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum label value. The type is interface{} with range: 16000..1048575.
    Minvalue interface{}

    // Maximum label value. The type is interface{} with range: 16000..1048575.
    MaxValue interface{}

    // Minimum static label value. The type is interface{} with range: 0..1048575.
    MinStaticValue interface{}

    // Maximum static label value. The type is interface{} with range: 0..1048575.
    MaxStaticValue interface{}
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetEntityData() *types.CommonEntityData {
    labelRange.EntityData.YFilter = labelRange.YFilter
    labelRange.EntityData.YangName = "label-range"
    labelRange.EntityData.BundleName = "cisco_ios_xr"
    labelRange.EntityData.ParentYangName = "label-database"
    labelRange.EntityData.SegmentPath = "label-range"
    labelRange.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd/label-databases/label-database/" + labelRange.EntityData.SegmentPath
    labelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelRange.EntityData.Children = types.NewOrderedMap()
    labelRange.EntityData.Leafs = types.NewOrderedMap()
    labelRange.EntityData.Leafs.Append("minvalue", types.YLeaf{"Minvalue", labelRange.Minvalue})
    labelRange.EntityData.Leafs.Append("max-value", types.YLeaf{"MaxValue", labelRange.MaxValue})
    labelRange.EntityData.Leafs.Append("min-static-value", types.YLeaf{"MinStaticValue", labelRange.MinStaticValue})
    labelRange.EntityData.Leafs.Append("max-static-value", types.YLeaf{"MaxStaticValue", labelRange.MaxStaticValue})

    labelRange.EntityData.YListKeys = []string {}

    return &(labelRange.EntityData)
}

