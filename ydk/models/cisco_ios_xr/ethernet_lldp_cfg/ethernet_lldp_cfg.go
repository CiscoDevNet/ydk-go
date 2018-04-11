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
    EntityData types.CommonEntityData
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

    // Enable or disable LLDP on Mgmt interfaces as well globally. The type is
    // bool. The default value is false.
    EnableMgmtintf interface{}

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

func (lldp *Lldp) GetEntityData() *types.CommonEntityData {
    lldp.EntityData.YFilter = lldp.YFilter
    lldp.EntityData.YangName = "lldp"
    lldp.EntityData.BundleName = "cisco_ios_xr"
    lldp.EntityData.ParentYangName = "Cisco-IOS-XR-ethernet-lldp-cfg"
    lldp.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-lldp-cfg:lldp"
    lldp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lldp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lldp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lldp.EntityData.Children = make(map[string]types.YChild)
    lldp.EntityData.Children["tlv-select"] = types.YChild{"TlvSelect", &lldp.TlvSelect}
    lldp.EntityData.Leafs = make(map[string]types.YLeaf)
    lldp.EntityData.Leafs["holdtime"] = types.YLeaf{"Holdtime", lldp.Holdtime}
    lldp.EntityData.Leafs["extended-show-width"] = types.YLeaf{"ExtendedShowWidth", lldp.ExtendedShowWidth}
    lldp.EntityData.Leafs["enable-subintf"] = types.YLeaf{"EnableSubintf", lldp.EnableSubintf}
    lldp.EntityData.Leafs["enable-mgmtintf"] = types.YLeaf{"EnableMgmtintf", lldp.EnableMgmtintf}
    lldp.EntityData.Leafs["timer"] = types.YLeaf{"Timer", lldp.Timer}
    lldp.EntityData.Leafs["reinit"] = types.YLeaf{"Reinit", lldp.Reinit}
    lldp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", lldp.Enable}
    return &(lldp.EntityData)
}

// Lldp_TlvSelect
// Selection of LLDP TLVs to disable
// This type is a presence type.
type Lldp_TlvSelect struct {
    EntityData types.CommonEntityData
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

func (tlvSelect *Lldp_TlvSelect) GetEntityData() *types.CommonEntityData {
    tlvSelect.EntityData.YFilter = tlvSelect.YFilter
    tlvSelect.EntityData.YangName = "tlv-select"
    tlvSelect.EntityData.BundleName = "cisco_ios_xr"
    tlvSelect.EntityData.ParentYangName = "lldp"
    tlvSelect.EntityData.SegmentPath = "tlv-select"
    tlvSelect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tlvSelect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tlvSelect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tlvSelect.EntityData.Children = make(map[string]types.YChild)
    tlvSelect.EntityData.Children["system-name"] = types.YChild{"SystemName", &tlvSelect.SystemName}
    tlvSelect.EntityData.Children["port-description"] = types.YChild{"PortDescription", &tlvSelect.PortDescription}
    tlvSelect.EntityData.Children["system-description"] = types.YChild{"SystemDescription", &tlvSelect.SystemDescription}
    tlvSelect.EntityData.Children["system-capabilities"] = types.YChild{"SystemCapabilities", &tlvSelect.SystemCapabilities}
    tlvSelect.EntityData.Children["management-address"] = types.YChild{"ManagementAddress", &tlvSelect.ManagementAddress}
    tlvSelect.EntityData.Leafs = make(map[string]types.YLeaf)
    tlvSelect.EntityData.Leafs["tlv-select-enter"] = types.YLeaf{"TlvSelectEnter", tlvSelect.TlvSelectEnter}
    return &(tlvSelect.EntityData)
}

// Lldp_TlvSelect_SystemName
// System Name TLV
type Lldp_TlvSelect_SystemName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable System Name TLV. The type is bool. The default value is false.
    Disable interface{}
}

func (systemName *Lldp_TlvSelect_SystemName) GetEntityData() *types.CommonEntityData {
    systemName.EntityData.YFilter = systemName.YFilter
    systemName.EntityData.YangName = "system-name"
    systemName.EntityData.BundleName = "cisco_ios_xr"
    systemName.EntityData.ParentYangName = "tlv-select"
    systemName.EntityData.SegmentPath = "system-name"
    systemName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemName.EntityData.Children = make(map[string]types.YChild)
    systemName.EntityData.Leafs = make(map[string]types.YLeaf)
    systemName.EntityData.Leafs["disable"] = types.YLeaf{"Disable", systemName.Disable}
    return &(systemName.EntityData)
}

// Lldp_TlvSelect_PortDescription
// Port Description TLV
type Lldp_TlvSelect_PortDescription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable Port Description TLV. The type is bool. The default value is false.
    Disable interface{}
}

func (portDescription *Lldp_TlvSelect_PortDescription) GetEntityData() *types.CommonEntityData {
    portDescription.EntityData.YFilter = portDescription.YFilter
    portDescription.EntityData.YangName = "port-description"
    portDescription.EntityData.BundleName = "cisco_ios_xr"
    portDescription.EntityData.ParentYangName = "tlv-select"
    portDescription.EntityData.SegmentPath = "port-description"
    portDescription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portDescription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portDescription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portDescription.EntityData.Children = make(map[string]types.YChild)
    portDescription.EntityData.Leafs = make(map[string]types.YLeaf)
    portDescription.EntityData.Leafs["disable"] = types.YLeaf{"Disable", portDescription.Disable}
    return &(portDescription.EntityData)
}

// Lldp_TlvSelect_SystemDescription
// System Description TLV
type Lldp_TlvSelect_SystemDescription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable System Description TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (systemDescription *Lldp_TlvSelect_SystemDescription) GetEntityData() *types.CommonEntityData {
    systemDescription.EntityData.YFilter = systemDescription.YFilter
    systemDescription.EntityData.YangName = "system-description"
    systemDescription.EntityData.BundleName = "cisco_ios_xr"
    systemDescription.EntityData.ParentYangName = "tlv-select"
    systemDescription.EntityData.SegmentPath = "system-description"
    systemDescription.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemDescription.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemDescription.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemDescription.EntityData.Children = make(map[string]types.YChild)
    systemDescription.EntityData.Leafs = make(map[string]types.YLeaf)
    systemDescription.EntityData.Leafs["disable"] = types.YLeaf{"Disable", systemDescription.Disable}
    return &(systemDescription.EntityData)
}

// Lldp_TlvSelect_SystemCapabilities
// System Capabilities TLV
type Lldp_TlvSelect_SystemCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable System Capabilities TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (systemCapabilities *Lldp_TlvSelect_SystemCapabilities) GetEntityData() *types.CommonEntityData {
    systemCapabilities.EntityData.YFilter = systemCapabilities.YFilter
    systemCapabilities.EntityData.YangName = "system-capabilities"
    systemCapabilities.EntityData.BundleName = "cisco_ios_xr"
    systemCapabilities.EntityData.ParentYangName = "tlv-select"
    systemCapabilities.EntityData.SegmentPath = "system-capabilities"
    systemCapabilities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemCapabilities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemCapabilities.EntityData.Children = make(map[string]types.YChild)
    systemCapabilities.EntityData.Leafs = make(map[string]types.YLeaf)
    systemCapabilities.EntityData.Leafs["disable"] = types.YLeaf{"Disable", systemCapabilities.Disable}
    return &(systemCapabilities.EntityData)
}

// Lldp_TlvSelect_ManagementAddress
// Management Address TLV
type Lldp_TlvSelect_ManagementAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // disable Management Address TLV. The type is bool. The default value is
    // false.
    Disable interface{}
}

func (managementAddress *Lldp_TlvSelect_ManagementAddress) GetEntityData() *types.CommonEntityData {
    managementAddress.EntityData.YFilter = managementAddress.YFilter
    managementAddress.EntityData.YangName = "management-address"
    managementAddress.EntityData.BundleName = "cisco_ios_xr"
    managementAddress.EntityData.ParentYangName = "tlv-select"
    managementAddress.EntityData.SegmentPath = "management-address"
    managementAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    managementAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    managementAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    managementAddress.EntityData.Children = make(map[string]types.YChild)
    managementAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    managementAddress.EntityData.Leafs["disable"] = types.YLeaf{"Disable", managementAddress.Disable}
    return &(managementAddress.EntityData)
}

