// This module contains a collection of YANG definitions
// for Cisco IOS-XR accounting package configuration.
// 
// This module contains definitions
// for the following management objects:
//   accounting: Global Accounting configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package accounting_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package accounting_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-accounting-cfg accounting}", reflect.TypeOf(Accounting{}))
    ydk.RegisterEntity("Cisco-IOS-XR-accounting-cfg:accounting", reflect.TypeOf(Accounting{}))
}

// Accounting
// Global Accounting configuration commands
type Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Accounting. The type is interface{}.
    Enable interface{}

    // Interfaces configuration.
    Interfaces Accounting_Interfaces
}

func (accounting *Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Accounting) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (accounting *Accounting) GetSegmentPath() string {
    return "Cisco-IOS-XR-accounting-cfg:accounting"
}

func (accounting *Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &accounting.Interfaces
    }
    return nil
}

func (accounting *Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &accounting.Interfaces
    return children
}

func (accounting *Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = accounting.Enable
    return leafs
}

func (accounting *Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Accounting) GetYangName() string { return "accounting" }

func (accounting *Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Accounting) GetParentYangName() string { return "Cisco-IOS-XR-accounting-cfg" }

// Accounting_Interfaces
// Interfaces configuration
type Accounting_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable accounting on Interfaces. The type is interface{}.
    Enable interface{}

    // Interfaces MPLS configuration.
    Mpls Accounting_Interfaces_Mpls

    // Interfaces Segment Routing configuration.
    SegmentRouting Accounting_Interfaces_SegmentRouting
}

func (interfaces *Accounting_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Accounting_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Accounting_Interfaces) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "mpls" { return "Mpls" }
    if yname == "segment-routing" { return "SegmentRouting" }
    return ""
}

func (interfaces *Accounting_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Accounting_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls" {
        return &interfaces.Mpls
    }
    if childYangName == "segment-routing" {
        return &interfaces.SegmentRouting
    }
    return nil
}

func (interfaces *Accounting_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls"] = &interfaces.Mpls
    children["segment-routing"] = &interfaces.SegmentRouting
    return children
}

func (interfaces *Accounting_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = interfaces.Enable
    return leafs
}

func (interfaces *Accounting_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Accounting_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Accounting_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Accounting_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Accounting_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Accounting_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Accounting_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Accounting_Interfaces) GetParentYangName() string { return "accounting" }

// Accounting_Interfaces_Mpls
// Interfaces MPLS configuration
type Accounting_Interfaces_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable accounting on MPLS. The type is interface{}.
    Enable interface{}

    // Enable accounting on MPLS IPv4 RSVP TE. The type is interface{}.
    EnableV4Rsvpte interface{}
}

func (mpls *Accounting_Interfaces_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *Accounting_Interfaces_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *Accounting_Interfaces_Mpls) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "enable-v4rsvpte" { return "EnableV4Rsvpte" }
    return ""
}

func (mpls *Accounting_Interfaces_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *Accounting_Interfaces_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mpls *Accounting_Interfaces_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mpls *Accounting_Interfaces_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mpls.Enable
    leafs["enable-v4rsvpte"] = mpls.EnableV4Rsvpte
    return leafs
}

func (mpls *Accounting_Interfaces_Mpls) GetBundleName() string { return "cisco_ios_xr" }

func (mpls *Accounting_Interfaces_Mpls) GetYangName() string { return "mpls" }

func (mpls *Accounting_Interfaces_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpls *Accounting_Interfaces_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpls *Accounting_Interfaces_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpls *Accounting_Interfaces_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *Accounting_Interfaces_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *Accounting_Interfaces_Mpls) GetParentYangName() string { return "interfaces" }

// Accounting_Interfaces_SegmentRouting
// Interfaces Segment Routing configuration
type Accounting_Interfaces_SegmentRouting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable accounting on Segment Routing. The type is interface{}.
    Enable interface{}

    // Enable accounting on Segment Routing MPLS IPv4. The type is interface{}.
    EnableMplsv4 interface{}

    // Enable accounting on Segment Routing MPLS IPv6. The type is interface{}.
    EnableMplsv6 interface{}
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetFilter() yfilter.YFilter { return segmentRouting.YFilter }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) SetFilter(yf yfilter.YFilter) { segmentRouting.YFilter = yf }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "enable-mplsv4" { return "EnableMplsv4" }
    if yname == "enable-mplsv6" { return "EnableMplsv6" }
    return ""
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetSegmentPath() string {
    return "segment-routing"
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = segmentRouting.Enable
    leafs["enable-mplsv4"] = segmentRouting.EnableMplsv4
    leafs["enable-mplsv6"] = segmentRouting.EnableMplsv6
    return leafs
}

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetBundleName() string { return "cisco_ios_xr" }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetYangName() string { return "segment-routing" }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) SetParent(parent types.Entity) { segmentRouting.parent = parent }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetParent() types.Entity { return segmentRouting.parent }

func (segmentRouting *Accounting_Interfaces_SegmentRouting) GetParentYangName() string { return "interfaces" }

