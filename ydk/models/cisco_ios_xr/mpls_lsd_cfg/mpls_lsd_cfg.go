// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-lsd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-lsd: MPLS LSD configuration data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

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

func (mplsLsd *MplsLsd) GetFilter() yfilter.YFilter { return mplsLsd.YFilter }

func (mplsLsd *MplsLsd) SetFilter(yf yfilter.YFilter) { mplsLsd.YFilter = yf }

func (mplsLsd *MplsLsd) GetGoName(yname string) string {
    if yname == "app-reg-delay-disable" { return "AppRegDelayDisable" }
    if yname == "mpls-entropy-label" { return "MplsEntropyLabel" }
    if yname == "mpls-ip-ttl-propagate-disable" { return "MplsIpTtlPropagateDisable" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "label-databases" { return "LabelDatabases" }
    return ""
}

func (mplsLsd *MplsLsd) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-lsd-cfg:mpls-lsd"
}

func (mplsLsd *MplsLsd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &mplsLsd.Ipv6
    }
    if childYangName == "ipv4" {
        return &mplsLsd.Ipv4
    }
    if childYangName == "label-databases" {
        return &mplsLsd.LabelDatabases
    }
    return nil
}

func (mplsLsd *MplsLsd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &mplsLsd.Ipv6
    children["ipv4"] = &mplsLsd.Ipv4
    children["label-databases"] = &mplsLsd.LabelDatabases
    return children
}

func (mplsLsd *MplsLsd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["app-reg-delay-disable"] = mplsLsd.AppRegDelayDisable
    leafs["mpls-entropy-label"] = mplsLsd.MplsEntropyLabel
    leafs["mpls-ip-ttl-propagate-disable"] = mplsLsd.MplsIpTtlPropagateDisable
    return leafs
}

func (mplsLsd *MplsLsd) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLsd *MplsLsd) GetYangName() string { return "mpls-lsd" }

func (mplsLsd *MplsLsd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLsd *MplsLsd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLsd *MplsLsd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLsd *MplsLsd) SetParent(parent types.Entity) { mplsLsd.parent = parent }

func (mplsLsd *MplsLsd) GetParent() types.Entity { return mplsLsd.parent }

func (mplsLsd *MplsLsd) GetParentYangName() string { return "Cisco-IOS-XR-mpls-lsd-cfg" }

// MplsLsd_Ipv6
// Configure IPv6 parameters
type MplsLsd_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of labels to pop upon MPLS IP TTL expiry. The type is interface{}
    // with range: 1..10.
    TtlExpirationPop interface{}
}

func (ipv6 *MplsLsd_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *MplsLsd_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *MplsLsd_Ipv6) GetGoName(yname string) string {
    if yname == "ttl-expiration-pop" { return "TtlExpirationPop" }
    return ""
}

func (ipv6 *MplsLsd_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *MplsLsd_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6 *MplsLsd_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6 *MplsLsd_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl-expiration-pop"] = ipv6.TtlExpirationPop
    return leafs
}

func (ipv6 *MplsLsd_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *MplsLsd_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *MplsLsd_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *MplsLsd_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *MplsLsd_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *MplsLsd_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *MplsLsd_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *MplsLsd_Ipv6) GetParentYangName() string { return "mpls-lsd" }

// MplsLsd_Ipv4
// Configure IPv4 parameters
type MplsLsd_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of labels to pop upon MPLS IP TTL expiry. The type is interface{}
    // with range: 1..10.
    TtlExpirationPop interface{}
}

func (ipv4 *MplsLsd_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *MplsLsd_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *MplsLsd_Ipv4) GetGoName(yname string) string {
    if yname == "ttl-expiration-pop" { return "TtlExpirationPop" }
    return ""
}

func (ipv4 *MplsLsd_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *MplsLsd_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *MplsLsd_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *MplsLsd_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl-expiration-pop"] = ipv4.TtlExpirationPop
    return leafs
}

func (ipv4 *MplsLsd_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *MplsLsd_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *MplsLsd_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *MplsLsd_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *MplsLsd_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *MplsLsd_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *MplsLsd_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *MplsLsd_Ipv4) GetParentYangName() string { return "mpls-lsd" }

// MplsLsd_LabelDatabases
// Table of label databases
type MplsLsd_LabelDatabases struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A label database. The type is slice of
    // MplsLsd_LabelDatabases_LabelDatabase.
    LabelDatabase []MplsLsd_LabelDatabases_LabelDatabase
}

func (labelDatabases *MplsLsd_LabelDatabases) GetFilter() yfilter.YFilter { return labelDatabases.YFilter }

func (labelDatabases *MplsLsd_LabelDatabases) SetFilter(yf yfilter.YFilter) { labelDatabases.YFilter = yf }

func (labelDatabases *MplsLsd_LabelDatabases) GetGoName(yname string) string {
    if yname == "label-database" { return "LabelDatabase" }
    return ""
}

func (labelDatabases *MplsLsd_LabelDatabases) GetSegmentPath() string {
    return "label-databases"
}

func (labelDatabases *MplsLsd_LabelDatabases) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-database" {
        for _, c := range labelDatabases.LabelDatabase {
            if labelDatabases.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLsd_LabelDatabases_LabelDatabase{}
        labelDatabases.LabelDatabase = append(labelDatabases.LabelDatabase, child)
        return &labelDatabases.LabelDatabase[len(labelDatabases.LabelDatabase)-1]
    }
    return nil
}

func (labelDatabases *MplsLsd_LabelDatabases) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelDatabases.LabelDatabase {
        children[labelDatabases.LabelDatabase[i].GetSegmentPath()] = &labelDatabases.LabelDatabase[i]
    }
    return children
}

func (labelDatabases *MplsLsd_LabelDatabases) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelDatabases *MplsLsd_LabelDatabases) GetBundleName() string { return "cisco_ios_xr" }

func (labelDatabases *MplsLsd_LabelDatabases) GetYangName() string { return "label-databases" }

func (labelDatabases *MplsLsd_LabelDatabases) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelDatabases *MplsLsd_LabelDatabases) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelDatabases *MplsLsd_LabelDatabases) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelDatabases *MplsLsd_LabelDatabases) SetParent(parent types.Entity) { labelDatabases.parent = parent }

func (labelDatabases *MplsLsd_LabelDatabases) GetParent() types.Entity { return labelDatabases.parent }

func (labelDatabases *MplsLsd_LabelDatabases) GetParentYangName() string { return "mpls-lsd" }

// MplsLsd_LabelDatabases_LabelDatabase
// A label database
type MplsLsd_LabelDatabases_LabelDatabase struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Label database identifier. The type is interface{}
    // with range: 0..4294967295.
    LabelDatabaseId interface{}

    // Label range.
    LabelRange MplsLsd_LabelDatabases_LabelDatabase_LabelRange
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetFilter() yfilter.YFilter { return labelDatabase.YFilter }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) SetFilter(yf yfilter.YFilter) { labelDatabase.YFilter = yf }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetGoName(yname string) string {
    if yname == "label-database-id" { return "LabelDatabaseId" }
    if yname == "label-range" { return "LabelRange" }
    return ""
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetSegmentPath() string {
    return "label-database" + "[label-database-id='" + fmt.Sprintf("%v", labelDatabase.LabelDatabaseId) + "']"
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-range" {
        return &labelDatabase.LabelRange
    }
    return nil
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-range"] = &labelDatabase.LabelRange
    return children
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-database-id"] = labelDatabase.LabelDatabaseId
    return leafs
}

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetBundleName() string { return "cisco_ios_xr" }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetYangName() string { return "label-database" }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) SetParent(parent types.Entity) { labelDatabase.parent = parent }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetParent() types.Entity { return labelDatabase.parent }

func (labelDatabase *MplsLsd_LabelDatabases_LabelDatabase) GetParentYangName() string { return "label-databases" }

// MplsLsd_LabelDatabases_LabelDatabase_LabelRange
// Label range
type MplsLsd_LabelDatabases_LabelDatabase_LabelRange struct {
    parent types.Entity
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

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetFilter() yfilter.YFilter { return labelRange.YFilter }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) SetFilter(yf yfilter.YFilter) { labelRange.YFilter = yf }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetGoName(yname string) string {
    if yname == "minvalue" { return "Minvalue" }
    if yname == "max-value" { return "MaxValue" }
    if yname == "min-static-value" { return "MinStaticValue" }
    if yname == "max-static-value" { return "MaxStaticValue" }
    return ""
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetSegmentPath() string {
    return "label-range"
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minvalue"] = labelRange.Minvalue
    leafs["max-value"] = labelRange.MaxValue
    leafs["min-static-value"] = labelRange.MinStaticValue
    leafs["max-static-value"] = labelRange.MaxStaticValue
    return leafs
}

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetBundleName() string { return "cisco_ios_xr" }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetYangName() string { return "label-range" }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) SetParent(parent types.Entity) { labelRange.parent = parent }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetParent() types.Entity { return labelRange.parent }

func (labelRange *MplsLsd_LabelDatabases_LabelDatabase_LabelRange) GetParentYangName() string { return "label-database" }

