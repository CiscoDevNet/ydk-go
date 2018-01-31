// This module contains a collection of YANG definitions
// for Cisco IOS-XR ethernet-lldp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lldp: Enable LLDP, or configure global LLDP subcommands
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ethernet_lldp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ethernet_lldp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ethernet-lldp-cfg lldp}", reflect.TypeOf(Lldp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ethernet-lldp-cfg:lldp", reflect.TypeOf(Lldp{}))
}

// Lldp
// Enable LLDP, or configure global LLDP subcommands
type Lldp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Length  of time  (in sec) that receiver must keep this packet. The type is
    // interface{} with range: 0..65535.
    Holdtime interface{}

    // Enable or disable LLDP Show LLDP Neighbor Extended Width. The type is bool.
    // The default value is false.
    ExtendedShowWidth interface{}

    // Enable or disable LLDP on Sub-interfaces as well globally. The type is
    // bool. The default value is false.
    EnableSubintf interface{}

    // Specify the rate at which LLDP packets are sent (in sec). The type is
    // interface{} with range: 5..65534. The default value is 30.
    Timer interface{}

    // Delay (in sec) for LLDP initialization on any interface. The type is
    // interface{} with range: 2..5. The default value is 2.
    Reinit interface{}

    // Enable or disable LLDP globally. The type is bool. The default value is
    // false.
    Enable interface{}

    // Selection of LLDP TLVs to disable.
    TlvSelect Lldp_TlvSelect
}

func (lldp *Lldp) GetFilter() yfilter.YFilter { return lldp.YFilter }

func (lldp *Lldp) SetFilter(yf yfilter.YFilter) { lldp.YFilter = yf }

func (lldp *Lldp) GetGoName(yname string) string {
    if yname == "holdtime" { return "Holdtime" }
    if yname == "extended-show-width" { return "ExtendedShowWidth" }
    if yname == "enable-subintf" { return "EnableSubintf" }
    if yname == "timer" { return "Timer" }
    if yname == "reinit" { return "Reinit" }
    if yname == "enable" { return "Enable" }
    if yname == "tlv-select" { return "TlvSelect" }
    return ""
}

func (lldp *Lldp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-lldp-cfg:lldp"
}

func (lldp *Lldp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tlv-select" {
        return &lldp.TlvSelect
    }
    return nil
}

func (lldp *Lldp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tlv-select"] = &lldp.TlvSelect
    return children
}

func (lldp *Lldp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["holdtime"] = lldp.Holdtime
    leafs["extended-show-width"] = lldp.ExtendedShowWidth
    leafs["enable-subintf"] = lldp.EnableSubintf
    leafs["timer"] = lldp.Timer
    leafs["reinit"] = lldp.Reinit
    leafs["enable"] = lldp.Enable
    return leafs
}

func (lldp *Lldp) GetBundleName() string { return "cisco_ios_xr" }

func (lldp *Lldp) GetYangName() string { return "lldp" }

func (lldp *Lldp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lldp *Lldp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lldp *Lldp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lldp *Lldp) SetParent(parent types.Entity) { lldp.parent = parent }

func (lldp *Lldp) GetParent() types.Entity { return lldp.parent }

func (lldp *Lldp) GetParentYangName() string { return "Cisco-IOS-XR-ethernet-lldp-cfg" }

// Lldp_TlvSelect
// Selection of LLDP TLVs to disable
// This type is a presence type.
type Lldp_TlvSelect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // enter lldp tlv-select submode. The type is bool. This attribute is
    // mandatory.
    TlvSelectEnter interface{}

    // System Name TLV.
    SystemName Lldp_TlvSelect_SystemName

    // Port Description TLV.
    PortDescription Lldp_TlvSelect_PortDescription

    // System Description TLV.
    SystemDescription Lldp_TlvSelect_SystemDescription

    // System Capabilities TLV.
    SystemCapabilities Lldp_TlvSelect_SystemCapabilities

    // Management Address TLV.
    ManagementAddress Lldp_TlvSelect_ManagementAddress
}

func (tlvSelect *Lldp_TlvSelect) GetFilter() yfilter.YFilter { return tlvSelect.YFilter }

func (tlvSelect *Lldp_TlvSelect) SetFilter(yf yfilter.YFilter) { tlvSelect.YFilter = yf }

func (tlvSelect *Lldp_TlvSelect) GetGoName(yname string) string {
    if yname == "tlv-select-enter" { return "TlvSelectEnter" }
    if yname == "system-name" { return "SystemName" }
    if yname == "port-description" { return "PortDescription" }
    if yname == "system-description" { return "SystemDescription" }
    if yname == "system-capabilities" { return "SystemCapabilities" }
    if yname == "management-address" { return "ManagementAddress" }
    return ""
}

func (tlvSelect *Lldp_TlvSelect) GetSegmentPath() string {
    return "tlv-select"
}

func (tlvSelect *Lldp_TlvSelect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "system-name" {
        return &tlvSelect.SystemName
    }
    if childYangName == "port-description" {
        return &tlvSelect.PortDescription
    }
    if childYangName == "system-description" {
        return &tlvSelect.SystemDescription
    }
    if childYangName == "system-capabilities" {
        return &tlvSelect.SystemCapabilities
    }
    if childYangName == "management-address" {
        return &tlvSelect.ManagementAddress
    }
    return nil
}

func (tlvSelect *Lldp_TlvSelect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["system-name"] = &tlvSelect.SystemName
    children["port-description"] = &tlvSelect.PortDescription
    children["system-description"] = &tlvSelect.SystemDescription
    children["system-capabilities"] = &tlvSelect.SystemCapabilities
    children["management-address"] = &tlvSelect.ManagementAddress
    return children
}

func (tlvSelect *Lldp_TlvSelect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tlv-select-enter"] = tlvSelect.TlvSelectEnter
    return leafs
}

func (tlvSelect *Lldp_TlvSelect) GetBundleName() string { return "cisco_ios_xr" }

func (tlvSelect *Lldp_TlvSelect) GetYangName() string { return "tlv-select" }

func (tlvSelect *Lldp_TlvSelect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tlvSelect *Lldp_TlvSelect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tlvSelect *Lldp_TlvSelect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tlvSelect *Lldp_TlvSelect) SetParent(parent types.Entity) { tlvSelect.parent = parent }

func (tlvSelect *Lldp_TlvSelect) GetParent() types.Entity { return tlvSelect.parent }

func (tlvSelect *Lldp_TlvSelect) GetParentYangName() string { return "lldp" }

// Lldp_TlvSelect_SystemName
// System Name TLV
type Lldp_TlvSelect_SystemName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable System Name TLV. The type is bool. The default value is false.
    Disable interface{}
}

func (systemName *Lldp_TlvSelect_SystemName) GetFilter() yfilter.YFilter { return systemName.YFilter }

func (systemName *Lldp_TlvSelect_SystemName) SetFilter(yf yfilter.YFilter) { systemName.YFilter = yf }

func (systemName *Lldp_TlvSelect_SystemName) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (systemName *Lldp_TlvSelect_SystemName) GetSegmentPath() string {
    return "system-name"
}

func (systemName *Lldp_TlvSelect_SystemName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemName *Lldp_TlvSelect_SystemName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemName *Lldp_TlvSelect_SystemName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = systemName.Disable
    return leafs
}

func (systemName *Lldp_TlvSelect_SystemName) GetBundleName() string { return "cisco_ios_xr" }

func (systemName *Lldp_TlvSelect_SystemName) GetYangName() string { return "system-name" }

func (systemName *Lldp_TlvSelect_SystemName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemName *Lldp_TlvSelect_SystemName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemName *Lldp_TlvSelect_SystemName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemName *Lldp_TlvSelect_SystemName) SetParent(parent types.Entity) { systemName.parent = parent }

func (systemName *Lldp_TlvSelect_SystemName) GetParent() types.Entity { return systemName.parent }

func (systemName *Lldp_TlvSelect_SystemName) GetParentYangName() string { return "tlv-select" }

// Lldp_TlvSelect_PortDescription
// Port Description TLV
type Lldp_TlvSelect_PortDescription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable Port Description TLV. The type is bool. The default value is false.
    Disable interface{}
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetFilter() yfilter.YFilter { return portDescription.YFilter }

func (portDescription *Lldp_TlvSelect_PortDescription) SetFilter(yf yfilter.YFilter) { portDescription.YFilter = yf }

func (portDescription *Lldp_TlvSelect_PortDescription) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetSegmentPath() string {
    return "port-description"
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = portDescription.Disable
    return leafs
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetBundleName() string { return "cisco_ios_xr" }

func (portDescription *Lldp_TlvSelect_PortDescription) GetYangName() string { return "port-description" }

func (portDescription *Lldp_TlvSelect_PortDescription) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portDescription *Lldp_TlvSelect_PortDescription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portDescription *Lldp_TlvSelect_PortDescription) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portDescription *Lldp_TlvSelect_PortDescription) SetParent(parent types.Entity) { portDescription.parent = parent }

func (portDescription *Lldp_TlvSelect_PortDescription) GetParent() types.Entity { return portDescription.parent }

func (portDescription *Lldp_TlvSelect_PortDescription) GetParentYangName() string { return "tlv-select" }

// Lldp_TlvSelect_SystemDescription
// System Description TLV
type Lldp_TlvSelect_SystemDescription struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable System Description TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetFilter() yfilter.YFilter { return systemDescription.YFilter }

func (systemDescription *Lldp_TlvSelect_SystemDescription) SetFilter(yf yfilter.YFilter) { systemDescription.YFilter = yf }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetSegmentPath() string {
    return "system-description"
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = systemDescription.Disable
    return leafs
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetBundleName() string { return "cisco_ios_xr" }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetYangName() string { return "system-description" }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemDescription *Lldp_TlvSelect_SystemDescription) SetParent(parent types.Entity) { systemDescription.parent = parent }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetParent() types.Entity { return systemDescription.parent }

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetParentYangName() string { return "tlv-select" }

// Lldp_TlvSelect_SystemCapabilities
// System Capabilities TLV
type Lldp_TlvSelect_SystemCapabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable System Capabilities TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetFilter() yfilter.YFilter { return systemCapabilities.YFilter }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) SetFilter(yf yfilter.YFilter) { systemCapabilities.YFilter = yf }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetSegmentPath() string {
    return "system-capabilities"
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = systemCapabilities.Disable
    return leafs
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetBundleName() string { return "cisco_ios_xr" }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetYangName() string { return "system-capabilities" }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) SetParent(parent types.Entity) { systemCapabilities.parent = parent }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetParent() types.Entity { return systemCapabilities.parent }

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetParentYangName() string { return "tlv-select" }

// Lldp_TlvSelect_ManagementAddress
// Management Address TLV
type Lldp_TlvSelect_ManagementAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // disable Management Address TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetFilter() yfilter.YFilter { return managementAddress.YFilter }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) SetFilter(yf yfilter.YFilter) { managementAddress.YFilter = yf }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetSegmentPath() string {
    return "management-address"
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = managementAddress.Disable
    return leafs
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetBundleName() string { return "cisco_ios_xr" }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetYangName() string { return "management-address" }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) SetParent(parent types.Entity) { managementAddress.parent = parent }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetParent() types.Entity { return managementAddress.parent }

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetParentYangName() string { return "tlv-select" }

