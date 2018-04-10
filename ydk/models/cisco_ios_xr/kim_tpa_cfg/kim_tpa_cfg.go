// This module contains a collection of YANG definitions
// for Cisco IOS-XR kim-tpa package configuration.
// 
// This module contains definitions
// for the following management objects:
//   tpa: tpa configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package kim_tpa_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package kim_tpa_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-kim-tpa-cfg tpa}", reflect.TypeOf(Tpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-kim-tpa-cfg:tpa", reflect.TypeOf(Tpa{}))
}

// Tpa
// tpa configuration commands
type Tpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF container.
    VrfNames Tpa_VrfNames

    // Third party app logging.
    Logging Tpa_Logging

    // Statistics.
    Statistics Tpa_Statistics
}

func (tpa *Tpa) GetEntityData() *types.CommonEntityData {
    tpa.EntityData.YFilter = tpa.YFilter
    tpa.EntityData.YangName = "tpa"
    tpa.EntityData.BundleName = "cisco_ios_xr"
    tpa.EntityData.ParentYangName = "Cisco-IOS-XR-kim-tpa-cfg"
    tpa.EntityData.SegmentPath = "Cisco-IOS-XR-kim-tpa-cfg:tpa"
    tpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tpa.EntityData.Children = make(map[string]types.YChild)
    tpa.EntityData.Children["vrf-names"] = types.YChild{"VrfNames", &tpa.VrfNames}
    tpa.EntityData.Children["logging"] = types.YChild{"Logging", &tpa.Logging}
    tpa.EntityData.Children["statistics"] = types.YChild{"Statistics", &tpa.Statistics}
    tpa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tpa.EntityData)
}

// Tpa_VrfNames
// VRF container
type Tpa_VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Tpa_VrfNames_VrfName.
    VrfName []Tpa_VrfNames_VrfName
}

func (vrfNames *Tpa_VrfNames) GetEntityData() *types.CommonEntityData {
    vrfNames.EntityData.YFilter = vrfNames.YFilter
    vrfNames.EntityData.YangName = "vrf-names"
    vrfNames.EntityData.BundleName = "cisco_ios_xr"
    vrfNames.EntityData.ParentYangName = "tpa"
    vrfNames.EntityData.SegmentPath = "vrf-names"
    vrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfNames.EntityData.Children = make(map[string]types.YChild)
    vrfNames.EntityData.Children["vrf-name"] = types.YChild{"VrfName", nil}
    for i := range vrfNames.VrfName {
        vrfNames.EntityData.Children[types.GetSegmentPath(&vrfNames.VrfName[i])] = types.YChild{"VrfName", &vrfNames.VrfName[i]}
    }
    vrfNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfNames.EntityData)
}

// Tpa_VrfNames_VrfName
// VRF name
type Tpa_VrfNames_VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // EastWest container.
    EastWestNames Tpa_VrfNames_VrfName_EastWestNames

    // Address family.
    AddressFamily Tpa_VrfNames_VrfName_AddressFamily
}

func (vrfName *Tpa_VrfNames_VrfName) GetEntityData() *types.CommonEntityData {
    vrfName.EntityData.YFilter = vrfName.YFilter
    vrfName.EntityData.YangName = "vrf-name"
    vrfName.EntityData.BundleName = "cisco_ios_xr"
    vrfName.EntityData.ParentYangName = "vrf-names"
    vrfName.EntityData.SegmentPath = "vrf-name" + "[vrf-name='" + fmt.Sprintf("%v", vrfName.VrfName) + "']"
    vrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfName.EntityData.Children = make(map[string]types.YChild)
    vrfName.EntityData.Children["east-west-names"] = types.YChild{"EastWestNames", &vrfName.EastWestNames}
    vrfName.EntityData.Children["address-family"] = types.YChild{"AddressFamily", &vrfName.AddressFamily}
    vrfName.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfName.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfName.VrfName}
    return &(vrfName.EntityData)
}

// Tpa_VrfNames_VrfName_EastWestNames
// EastWest container
type Tpa_VrfNames_VrfName_EastWestNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // East West interface. The type is slice of
    // Tpa_VrfNames_VrfName_EastWestNames_EastWestName.
    EastWestName []Tpa_VrfNames_VrfName_EastWestNames_EastWestName
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetEntityData() *types.CommonEntityData {
    eastWestNames.EntityData.YFilter = eastWestNames.YFilter
    eastWestNames.EntityData.YangName = "east-west-names"
    eastWestNames.EntityData.BundleName = "cisco_ios_xr"
    eastWestNames.EntityData.ParentYangName = "vrf-name"
    eastWestNames.EntityData.SegmentPath = "east-west-names"
    eastWestNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eastWestNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eastWestNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eastWestNames.EntityData.Children = make(map[string]types.YChild)
    eastWestNames.EntityData.Children["east-west-name"] = types.YChild{"EastWestName", nil}
    for i := range eastWestNames.EastWestName {
        eastWestNames.EntityData.Children[types.GetSegmentPath(&eastWestNames.EastWestName[i])] = types.YChild{"EastWestName", &eastWestNames.EastWestName[i]}
    }
    eastWestNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eastWestNames.EntityData)
}

// Tpa_VrfNames_VrfName_EastWestNames_EastWestName
// East West interface
type Tpa_VrfNames_VrfName_EastWestNames_EastWestName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    EastWestName interface{}

    // VRF name. The type is string.
    Vrf interface{}

    // Interface. The type is string.
    Interface_ interface{}
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetEntityData() *types.CommonEntityData {
    eastWestName.EntityData.YFilter = eastWestName.YFilter
    eastWestName.EntityData.YangName = "east-west-name"
    eastWestName.EntityData.BundleName = "cisco_ios_xr"
    eastWestName.EntityData.ParentYangName = "east-west-names"
    eastWestName.EntityData.SegmentPath = "east-west-name" + "[east-west-name='" + fmt.Sprintf("%v", eastWestName.EastWestName) + "']"
    eastWestName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eastWestName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eastWestName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eastWestName.EntityData.Children = make(map[string]types.YChild)
    eastWestName.EntityData.Leafs = make(map[string]types.YLeaf)
    eastWestName.EntityData.Leafs["east-west-name"] = types.YLeaf{"EastWestName", eastWestName.EastWestName}
    eastWestName.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", eastWestName.Vrf}
    eastWestName.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", eastWestName.Interface_}
    return &(eastWestName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily
// Address family
type Tpa_VrfNames_VrfName_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 configuration.
    Ipv6 Tpa_VrfNames_VrfName_AddressFamily_Ipv6

    // IPv4 configuration.
    Ipv4 Tpa_VrfNames_VrfName_AddressFamily_Ipv4
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xr"
    addressFamily.EntityData.ParentYangName = "vrf-name"
    addressFamily.EntityData.SegmentPath = "address-family"
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = make(map[string]types.YChild)
    addressFamily.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &addressFamily.Ipv6}
    addressFamily.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &addressFamily.Ipv4}
    addressFamily.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamily.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6
// IPv6 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface name for source address. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UpdateSource interface{}
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "address-family"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["default-route"] = types.YLeaf{"DefaultRoute", ipv6.DefaultRoute}
    ipv6.EntityData.Leafs["update-source"] = types.YLeaf{"UpdateSource", ipv6.UpdateSource}
    return &(ipv6.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4
// IPv4 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface name for source address. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    UpdateSource interface{}
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "address-family"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["default-route"] = types.YLeaf{"DefaultRoute", ipv4.DefaultRoute}
    ipv4.EntityData.Leafs["update-source"] = types.YLeaf{"UpdateSource", ipv4.UpdateSource}
    return &(ipv4.EntityData)
}

// Tpa_Logging
// Third party app logging
type Tpa_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // KIM logging.
    Kim Tpa_Logging_Kim
}

func (logging *Tpa_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "tpa"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Children["kim"] = types.YChild{"Kim", &logging.Kim}
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(logging.EntityData)
}

// Tpa_Logging_Kim
// KIM logging
type Tpa_Logging_Kim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How many log rotation files to keep. The type is interface{} with range:
    // -2147483648..2147483647.
    RotationMax interface{}

    // Size in Kilobytes of the log file. The type is interface{} with range:
    // -2147483648..2147483647. Units are kilobyte.
    FileSizeMaxKb interface{}
}

func (kim *Tpa_Logging_Kim) GetEntityData() *types.CommonEntityData {
    kim.EntityData.YFilter = kim.YFilter
    kim.EntityData.YangName = "kim"
    kim.EntityData.BundleName = "cisco_ios_xr"
    kim.EntityData.ParentYangName = "logging"
    kim.EntityData.SegmentPath = "kim"
    kim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    kim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    kim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    kim.EntityData.Children = make(map[string]types.YChild)
    kim.EntityData.Leafs = make(map[string]types.YLeaf)
    kim.EntityData.Leafs["rotation-max"] = types.YLeaf{"RotationMax", kim.RotationMax}
    kim.EntityData.Leafs["file-size-max-kb"] = types.YLeaf{"FileSizeMaxKb", kim.FileSizeMaxKb}
    return &(kim.EntityData)
}

// Tpa_Statistics
// Statistics
type Tpa_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How many interface events to record. The type is interface{} with range:
    // -2147483648..2147483647.
    MaxIntfEvents interface{}

    // How many LPTS events to record. The type is interface{} with range:
    // -2147483648..2147483647.
    MaxLptsEvents interface{}

    // Statistics update frequency into Linux. The type is interface{} with range:
    // -2147483648..2147483647. Units are second.
    StatisticsUpdateFrequency interface{}
}

func (statistics *Tpa_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "tpa"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    statistics.EntityData.Leafs["max-intf-events"] = types.YLeaf{"MaxIntfEvents", statistics.MaxIntfEvents}
    statistics.EntityData.Leafs["max-lpts-events"] = types.YLeaf{"MaxLptsEvents", statistics.MaxLptsEvents}
    statistics.EntityData.Leafs["statistics-update-frequency"] = types.YLeaf{"StatisticsUpdateFrequency", statistics.StatisticsUpdateFrequency}
    return &(statistics.EntityData)
}

