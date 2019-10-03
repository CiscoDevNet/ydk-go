// This module contains a collection of YANG definitions
// for Cisco IOS-XR icpe-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   nv-satellite-global: nV Satellite Global configuration
//   nv-satellites: nv satellites
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-rgmgr-cfg,
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package icpe_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package icpe_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-icpe-infra-cfg nv-satellite-global}", reflect.TypeOf(NvSatelliteGlobal{}))
    ydk.RegisterEntity("Cisco-IOS-XR-icpe-infra-cfg:nv-satellite-global", reflect.TypeOf(NvSatelliteGlobal{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-icpe-infra-cfg nv-satellites}", reflect.TypeOf(NvSatellites{}))
    ydk.RegisterEntity("Cisco-IOS-XR-icpe-infra-cfg:nv-satellites", reflect.TypeOf(NvSatellites{}))
}

// NvSatelliteGlobal
// nV Satellite Global configuration
type NvSatelliteGlobal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Chassis MAC address.
    ChassisMac NvSatelliteGlobal_ChassisMac
}

func (nvSatelliteGlobal *NvSatelliteGlobal) GetEntityData() *types.CommonEntityData {
    nvSatelliteGlobal.EntityData.YFilter = nvSatelliteGlobal.YFilter
    nvSatelliteGlobal.EntityData.YangName = "nv-satellite-global"
    nvSatelliteGlobal.EntityData.BundleName = "cisco_ios_xr"
    nvSatelliteGlobal.EntityData.ParentYangName = "Cisco-IOS-XR-icpe-infra-cfg"
    nvSatelliteGlobal.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite-global"
    nvSatelliteGlobal.EntityData.AbsolutePath = nvSatelliteGlobal.EntityData.SegmentPath
    nvSatelliteGlobal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatelliteGlobal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatelliteGlobal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatelliteGlobal.EntityData.Children = types.NewOrderedMap()
    nvSatelliteGlobal.EntityData.Children.Append("chassis-mac", types.YChild{"ChassisMac", &nvSatelliteGlobal.ChassisMac})
    nvSatelliteGlobal.EntityData.Leafs = types.NewOrderedMap()

    nvSatelliteGlobal.EntityData.YListKeys = []string {}

    return &(nvSatelliteGlobal.EntityData)
}

// NvSatelliteGlobal_ChassisMac
// Chassis MAC address
type NvSatelliteGlobal_ChassisMac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac1 interface{}

    // Second two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac2 interface{}

    // Third two bytes of MAC address. The type is interface{} with range:
    // 0..2147483647. Units are byte.
    Mac3 interface{}
}

func (chassisMac *NvSatelliteGlobal_ChassisMac) GetEntityData() *types.CommonEntityData {
    chassisMac.EntityData.YFilter = chassisMac.YFilter
    chassisMac.EntityData.YangName = "chassis-mac"
    chassisMac.EntityData.BundleName = "cisco_ios_xr"
    chassisMac.EntityData.ParentYangName = "nv-satellite-global"
    chassisMac.EntityData.SegmentPath = "chassis-mac"
    chassisMac.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellite-global/" + chassisMac.EntityData.SegmentPath
    chassisMac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    chassisMac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    chassisMac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    chassisMac.EntityData.Children = types.NewOrderedMap()
    chassisMac.EntityData.Leafs = types.NewOrderedMap()
    chassisMac.EntityData.Leafs.Append("mac1", types.YLeaf{"Mac1", chassisMac.Mac1})
    chassisMac.EntityData.Leafs.Append("mac2", types.YLeaf{"Mac2", chassisMac.Mac2})
    chassisMac.EntityData.Leafs.Append("mac3", types.YLeaf{"Mac3", chassisMac.Mac3})

    chassisMac.EntityData.YListKeys = []string {}

    return &(chassisMac.EntityData)
}

// NvSatellites
// nv satellites
type NvSatellites struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite Configuration. The type is slice of NvSatellites_NvSatellite.
    NvSatellite []*NvSatellites_NvSatellite
}

func (nvSatellites *NvSatellites) GetEntityData() *types.CommonEntityData {
    nvSatellites.EntityData.YFilter = nvSatellites.YFilter
    nvSatellites.EntityData.YangName = "nv-satellites"
    nvSatellites.EntityData.BundleName = "cisco_ios_xr"
    nvSatellites.EntityData.ParentYangName = "Cisco-IOS-XR-icpe-infra-cfg"
    nvSatellites.EntityData.SegmentPath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites"
    nvSatellites.EntityData.AbsolutePath = nvSatellites.EntityData.SegmentPath
    nvSatellites.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellites.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellites.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellites.EntityData.Children = types.NewOrderedMap()
    nvSatellites.EntityData.Children.Append("nv-satellite", types.YChild{"NvSatellite", nil})
    for i := range nvSatellites.NvSatellite {
        nvSatellites.EntityData.Children.Append(types.GetSegmentPath(nvSatellites.NvSatellite[i]), types.YChild{"NvSatellite", nvSatellites.NvSatellite[i]})
    }
    nvSatellites.EntityData.Leafs = types.NewOrderedMap()

    nvSatellites.EntityData.YListKeys = []string {}

    return &(nvSatellites.EntityData)
}

// NvSatellites_NvSatellite
// Satellite Configuration
type NvSatellites_NvSatellite struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Satellite ID. The type is interface{} with range:
    // 100..65534.
    SatelliteId interface{}

    // VRF for Satellite IP Address. The type is string.
    Vrf interface{}

    // Discovery timeout warning for the satellite. The type is interface{} with
    // range: 0..4294967295.
    TimeoutWarning interface{}

    // Satellite Name. The type is string.
    DeviceName interface{}

    // Satellite Description. The type is string.
    Description interface{}

    // Satellite Type. The type is string. This attribute is mandatory.
    Type interface{}

    // Enable. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Discovery timeout for the satellite. The type is interface{} with range:
    // 0..4294967295.
    DiscTimeout interface{}

    // Timer (in seconds) for delaying switchback in a dual home setup. The type
    // is interface{} with range: 0..4294967295. Units are second.
    DelayedSwitchback interface{}

    // Satellite Serial Number. The type is string.
    SerialNumber interface{}

    // Encrypted password for the Satellite. The type is string with pattern:
    // (!.+)|([^!].+).
    Secret interface{}

    // Satellite IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Satellite auto-upgrade capability.
    UpgradeOnConnect NvSatellites_NvSatellite_UpgradeOnConnect

    // Enable interfaces on the satellite to be used as fabric ports table.
    CandidateFabricPorts NvSatellites_NvSatellite_CandidateFabricPorts

    // Satellite User.
    ConnectionInfo NvSatellites_NvSatellite_ConnectionInfo

    // Redundancy submode.
    Redundancy NvSatellites_NvSatellite_Redundancy
}

func (nvSatellite *NvSatellites_NvSatellite) GetEntityData() *types.CommonEntityData {
    nvSatellite.EntityData.YFilter = nvSatellite.YFilter
    nvSatellite.EntityData.YangName = "nv-satellite"
    nvSatellite.EntityData.BundleName = "cisco_ios_xr"
    nvSatellite.EntityData.ParentYangName = "nv-satellites"
    nvSatellite.EntityData.SegmentPath = "nv-satellite" + types.AddKeyToken(nvSatellite.SatelliteId, "satellite-id")
    nvSatellite.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/" + nvSatellite.EntityData.SegmentPath
    nvSatellite.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nvSatellite.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nvSatellite.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nvSatellite.EntityData.Children = types.NewOrderedMap()
    nvSatellite.EntityData.Children.Append("upgrade-on-connect", types.YChild{"UpgradeOnConnect", &nvSatellite.UpgradeOnConnect})
    nvSatellite.EntityData.Children.Append("candidate-fabric-ports", types.YChild{"CandidateFabricPorts", &nvSatellite.CandidateFabricPorts})
    nvSatellite.EntityData.Children.Append("connection-info", types.YChild{"ConnectionInfo", &nvSatellite.ConnectionInfo})
    nvSatellite.EntityData.Children.Append("redundancy", types.YChild{"Redundancy", &nvSatellite.Redundancy})
    nvSatellite.EntityData.Leafs = types.NewOrderedMap()
    nvSatellite.EntityData.Leafs.Append("satellite-id", types.YLeaf{"SatelliteId", nvSatellite.SatelliteId})
    nvSatellite.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", nvSatellite.Vrf})
    nvSatellite.EntityData.Leafs.Append("timeout-warning", types.YLeaf{"TimeoutWarning", nvSatellite.TimeoutWarning})
    nvSatellite.EntityData.Leafs.Append("device-name", types.YLeaf{"DeviceName", nvSatellite.DeviceName})
    nvSatellite.EntityData.Leafs.Append("description", types.YLeaf{"Description", nvSatellite.Description})
    nvSatellite.EntityData.Leafs.Append("type", types.YLeaf{"Type", nvSatellite.Type})
    nvSatellite.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", nvSatellite.Enable})
    nvSatellite.EntityData.Leafs.Append("disc-timeout", types.YLeaf{"DiscTimeout", nvSatellite.DiscTimeout})
    nvSatellite.EntityData.Leafs.Append("delayed-switchback", types.YLeaf{"DelayedSwitchback", nvSatellite.DelayedSwitchback})
    nvSatellite.EntityData.Leafs.Append("serial-number", types.YLeaf{"SerialNumber", nvSatellite.SerialNumber})
    nvSatellite.EntityData.Leafs.Append("secret", types.YLeaf{"Secret", nvSatellite.Secret})
    nvSatellite.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", nvSatellite.IpAddress})

    nvSatellite.EntityData.YListKeys = []string {"SatelliteId"}

    return &(nvSatellite.EntityData)
}

// NvSatellites_NvSatellite_UpgradeOnConnect
// Satellite auto-upgrade capability
type NvSatellites_NvSatellite_UpgradeOnConnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // When to upgrade the satellite. The type is ConnectType.
    ConnectType interface{}

    // Reference name. The type is string.
    Reference interface{}

    // Image Reference name. The type is string.
    ImageReference interface{}
}

func (upgradeOnConnect *NvSatellites_NvSatellite_UpgradeOnConnect) GetEntityData() *types.CommonEntityData {
    upgradeOnConnect.EntityData.YFilter = upgradeOnConnect.YFilter
    upgradeOnConnect.EntityData.YangName = "upgrade-on-connect"
    upgradeOnConnect.EntityData.BundleName = "cisco_ios_xr"
    upgradeOnConnect.EntityData.ParentYangName = "nv-satellite"
    upgradeOnConnect.EntityData.SegmentPath = "upgrade-on-connect"
    upgradeOnConnect.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/nv-satellite/" + upgradeOnConnect.EntityData.SegmentPath
    upgradeOnConnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    upgradeOnConnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    upgradeOnConnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    upgradeOnConnect.EntityData.Children = types.NewOrderedMap()
    upgradeOnConnect.EntityData.Leafs = types.NewOrderedMap()
    upgradeOnConnect.EntityData.Leafs.Append("connect-type", types.YLeaf{"ConnectType", upgradeOnConnect.ConnectType})
    upgradeOnConnect.EntityData.Leafs.Append("reference", types.YLeaf{"Reference", upgradeOnConnect.Reference})
    upgradeOnConnect.EntityData.Leafs.Append("image-reference", types.YLeaf{"ImageReference", upgradeOnConnect.ImageReference})

    upgradeOnConnect.EntityData.YListKeys = []string {}

    return &(upgradeOnConnect.EntityData)
}

// NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType represents When to upgrade the satellite
type NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType string

const (
    // Satellite Upgrade on Connection
    NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType_on_connection NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType = "on-connection"

    // Satellite Upgrade on First Connection after
    // configuration or host boot
    NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType_on_first_connection NvSatellites_NvSatellite_UpgradeOnConnect_ConnectType = "on-first-connection"
)

// NvSatellites_NvSatellite_CandidateFabricPorts
// Enable interfaces on the satellite to be used
// as fabric ports table
type NvSatellites_NvSatellite_CandidateFabricPorts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable interfaces on the satellite to be used as fabric ports. The type is
    // slice of NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort.
    CandidateFabricPort []*NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort
}

func (candidateFabricPorts *NvSatellites_NvSatellite_CandidateFabricPorts) GetEntityData() *types.CommonEntityData {
    candidateFabricPorts.EntityData.YFilter = candidateFabricPorts.YFilter
    candidateFabricPorts.EntityData.YangName = "candidate-fabric-ports"
    candidateFabricPorts.EntityData.BundleName = "cisco_ios_xr"
    candidateFabricPorts.EntityData.ParentYangName = "nv-satellite"
    candidateFabricPorts.EntityData.SegmentPath = "candidate-fabric-ports"
    candidateFabricPorts.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/nv-satellite/" + candidateFabricPorts.EntityData.SegmentPath
    candidateFabricPorts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateFabricPorts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateFabricPorts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateFabricPorts.EntityData.Children = types.NewOrderedMap()
    candidateFabricPorts.EntityData.Children.Append("candidate-fabric-port", types.YChild{"CandidateFabricPort", nil})
    for i := range candidateFabricPorts.CandidateFabricPort {
        candidateFabricPorts.EntityData.Children.Append(types.GetSegmentPath(candidateFabricPorts.CandidateFabricPort[i]), types.YChild{"CandidateFabricPort", candidateFabricPorts.CandidateFabricPort[i]})
    }
    candidateFabricPorts.EntityData.Leafs = types.NewOrderedMap()

    candidateFabricPorts.EntityData.YListKeys = []string {}

    return &(candidateFabricPorts.EntityData)
}

// NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort
// Enable interfaces on the satellite to be used
// as fabric ports
type NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PortType interface{}

    // This attribute is a key. Slot. The type is interface{} with range: 0..8.
    Slot interface{}

    // This attribute is a key. Sub slot. The type is interface{} with range:
    // 0..8.
    SubSlot interface{}

    // Port range. The type is string. This attribute is mandatory.
    PortRange interface{}
}

func (candidateFabricPort *NvSatellites_NvSatellite_CandidateFabricPorts_CandidateFabricPort) GetEntityData() *types.CommonEntityData {
    candidateFabricPort.EntityData.YFilter = candidateFabricPort.YFilter
    candidateFabricPort.EntityData.YangName = "candidate-fabric-port"
    candidateFabricPort.EntityData.BundleName = "cisco_ios_xr"
    candidateFabricPort.EntityData.ParentYangName = "candidate-fabric-ports"
    candidateFabricPort.EntityData.SegmentPath = "candidate-fabric-port" + types.AddKeyToken(candidateFabricPort.PortType, "port-type") + types.AddKeyToken(candidateFabricPort.Slot, "slot") + types.AddKeyToken(candidateFabricPort.SubSlot, "sub-slot")
    candidateFabricPort.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/nv-satellite/candidate-fabric-ports/" + candidateFabricPort.EntityData.SegmentPath
    candidateFabricPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateFabricPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateFabricPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateFabricPort.EntityData.Children = types.NewOrderedMap()
    candidateFabricPort.EntityData.Leafs = types.NewOrderedMap()
    candidateFabricPort.EntityData.Leafs.Append("port-type", types.YLeaf{"PortType", candidateFabricPort.PortType})
    candidateFabricPort.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", candidateFabricPort.Slot})
    candidateFabricPort.EntityData.Leafs.Append("sub-slot", types.YLeaf{"SubSlot", candidateFabricPort.SubSlot})
    candidateFabricPort.EntityData.Leafs.Append("port-range", types.YLeaf{"PortRange", candidateFabricPort.PortRange})

    candidateFabricPort.EntityData.YListKeys = []string {"PortType", "Slot", "SubSlot"}

    return &(candidateFabricPort.EntityData)
}

// NvSatellites_NvSatellite_ConnectionInfo
// Satellite User
type NvSatellites_NvSatellite_ConnectionInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Satellite Username. The type is string.
    Username interface{}

    // Encrypted password for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}
}

func (connectionInfo *NvSatellites_NvSatellite_ConnectionInfo) GetEntityData() *types.CommonEntityData {
    connectionInfo.EntityData.YFilter = connectionInfo.YFilter
    connectionInfo.EntityData.YangName = "connection-info"
    connectionInfo.EntityData.BundleName = "cisco_ios_xr"
    connectionInfo.EntityData.ParentYangName = "nv-satellite"
    connectionInfo.EntityData.SegmentPath = "connection-info"
    connectionInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/nv-satellite/" + connectionInfo.EntityData.SegmentPath
    connectionInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionInfo.EntityData.Children = types.NewOrderedMap()
    connectionInfo.EntityData.Leafs = types.NewOrderedMap()
    connectionInfo.EntityData.Leafs.Append("username", types.YLeaf{"Username", connectionInfo.Username})
    connectionInfo.EntityData.Leafs.Append("password", types.YLeaf{"Password", connectionInfo.Password})

    connectionInfo.EntityData.YListKeys = []string {}

    return &(connectionInfo.EntityData)
}

// NvSatellites_NvSatellite_Redundancy
// Redundancy submode
type NvSatellites_NvSatellite_Redundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Priority for this host for the given satellite. The type is interface{}
    // with range: 0..255.
    HostPriority interface{}
}

func (redundancy *NvSatellites_NvSatellite_Redundancy) GetEntityData() *types.CommonEntityData {
    redundancy.EntityData.YFilter = redundancy.YFilter
    redundancy.EntityData.YangName = "redundancy"
    redundancy.EntityData.BundleName = "cisco_ios_xr"
    redundancy.EntityData.ParentYangName = "nv-satellite"
    redundancy.EntityData.SegmentPath = "redundancy"
    redundancy.EntityData.AbsolutePath = "Cisco-IOS-XR-icpe-infra-cfg:nv-satellites/nv-satellite/" + redundancy.EntityData.SegmentPath
    redundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redundancy.EntityData.Children = types.NewOrderedMap()
    redundancy.EntityData.Leafs = types.NewOrderedMap()
    redundancy.EntityData.Leafs.Append("host-priority", types.YLeaf{"HostPriority", redundancy.HostPriority})

    redundancy.EntityData.YListKeys = []string {}

    return &(redundancy.EntityData)
}

