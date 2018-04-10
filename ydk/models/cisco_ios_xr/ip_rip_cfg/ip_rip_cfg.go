// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-rip package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rip: RIP configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_rip_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_rip_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-rip-cfg rip}", reflect.TypeOf(Rip{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-rip-cfg:rip", reflect.TypeOf(Rip{}))
}

// RipExtCommunity represents Rip ext community
type RipExtCommunity string

const (
    // AS:nn format
    RipExtCommunity_as RipExtCommunity = "as"

    // IPV4Address:nn format
    RipExtCommunity_ipv4_address RipExtCommunity = "ipv4-address"

    // 4-byte ASN format
    RipExtCommunity_four_byte_as RipExtCommunity = "four-byte-as"
)

// DefaultRedistRoute represents Default redist route
type DefaultRedistRoute string

const (
    // All routes
    DefaultRedistRoute_all DefaultRedistRoute = "all"
)

// DefaultInformationOption represents Default information option
type DefaultInformationOption string

const (
    // Always
    DefaultInformationOption_always DefaultInformationOption = "always"

    // Use route-policy
    DefaultInformationOption_policy DefaultInformationOption = "policy"
)

// IsisRedistRoute represents Isis redist route
type IsisRedistRoute string

const (
    // Level-1 routes only
    IsisRedistRoute_level1 IsisRedistRoute = "level1"

    // Level-1 routes only
    IsisRedistRoute_level2 IsisRedistRoute = "level2"

    // Level-1 and level-2 routes
    IsisRedistRoute_level1_and2 IsisRedistRoute = "level1-and2"
)

// BgpRedistRoute represents Bgp redist route
type BgpRedistRoute string

const (
    // All routes
    BgpRedistRoute_all BgpRedistRoute = "all"

    // Internal routes only
    BgpRedistRoute_internal BgpRedistRoute = "internal"

    // External routes only
    BgpRedistRoute_external BgpRedistRoute = "external"

    // Local routes only
    BgpRedistRoute_local BgpRedistRoute = "local"
)

// RipAuthMode represents Rip auth mode
type RipAuthMode string

const (
    // Text mode
    RipAuthMode_text RipAuthMode = "text"

    // MD5 mode
    RipAuthMode_md5 RipAuthMode = "md5"
)

// Rip
// RIP configuration
type Rip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP configuration for Default VRF.
    DefaultVrf Rip_DefaultVrf

    // VRF related RIP configuration.
    Vrfs Rip_Vrfs
}

func (rip *Rip) GetEntityData() *types.CommonEntityData {
    rip.EntityData.YFilter = rip.YFilter
    rip.EntityData.YangName = "rip"
    rip.EntityData.BundleName = "cisco_ios_xr"
    rip.EntityData.ParentYangName = "Cisco-IOS-XR-ip-rip-cfg"
    rip.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rip-cfg:rip"
    rip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rip.EntityData.Children = make(map[string]types.YChild)
    rip.EntityData.Children["default-vrf"] = types.YChild{"DefaultVrf", &rip.DefaultVrf}
    rip.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &rip.Vrfs}
    rip.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rip.EntityData)
}

// Rip_DefaultVrf
// RIP configuration for Default VRF
type Rip_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Starts RIP configuration for Default VRF. The type is interface{}.
    Enable interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Administrative distance. The type is interface{} with range: 0..255. The
    // default value is 120.
    Distance interface{}

    // Default metric of redistributed routes. The type is interface{} with range:
    // 0..16.
    DefaultMetric interface{}

    // Inter-packet delay for RIP updates. The type is interface{} with range:
    // 8..50. Units are millisecond.
    OutputDelay interface{}

    // Enable automatic network number summarization. The type is interface{}.
    AutoSummary interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Disable validation of source address of routing updates. The type is
    // interface{}.
    ValidateSourceDisable interface{}

    // Maximum number of paths allowed per route. The type is interface{} with
    // range: 1..128. The default value is 4.
    MaximumPaths interface{}

    // Enable Cisco Non Stop Forwarding. The type is interface{}.
    Nsf interface{}

    // Route Policy for inbbound routing updates. The type is string.
    PolicyIn interface{}

    // Controls default information origination.
    DefaultInformation Rip_DefaultVrf_DefaultInformation

    // Redistribute information from another routing protocol.
    Redistribution Rip_DefaultVrf_Redistribution

    // Table of IP specific administrative distances.
    IpDistances Rip_DefaultVrf_IpDistances

    // Table of RIP interfaces.
    Interfaces Rip_DefaultVrf_Interfaces

    // Configure RIP Neighbors.
    Neighbors Rip_DefaultVrf_Neighbors

    // Various routing timers.
    Timers Rip_DefaultVrf_Timers
}

func (defaultVrf *Rip_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "rip"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = make(map[string]types.YChild)
    defaultVrf.EntityData.Children["default-information"] = types.YChild{"DefaultInformation", &defaultVrf.DefaultInformation}
    defaultVrf.EntityData.Children["redistribution"] = types.YChild{"Redistribution", &defaultVrf.Redistribution}
    defaultVrf.EntityData.Children["ip-distances"] = types.YChild{"IpDistances", &defaultVrf.IpDistances}
    defaultVrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &defaultVrf.Interfaces}
    defaultVrf.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &defaultVrf.Neighbors}
    defaultVrf.EntityData.Children["timers"] = types.YChild{"Timers", &defaultVrf.Timers}
    defaultVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultVrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", defaultVrf.Enable}
    defaultVrf.EntityData.Leafs["broadcast-for-v2"] = types.YLeaf{"BroadcastForV2", defaultVrf.BroadcastForV2}
    defaultVrf.EntityData.Leafs["distance"] = types.YLeaf{"Distance", defaultVrf.Distance}
    defaultVrf.EntityData.Leafs["default-metric"] = types.YLeaf{"DefaultMetric", defaultVrf.DefaultMetric}
    defaultVrf.EntityData.Leafs["output-delay"] = types.YLeaf{"OutputDelay", defaultVrf.OutputDelay}
    defaultVrf.EntityData.Leafs["auto-summary"] = types.YLeaf{"AutoSummary", defaultVrf.AutoSummary}
    defaultVrf.EntityData.Leafs["policy-out"] = types.YLeaf{"PolicyOut", defaultVrf.PolicyOut}
    defaultVrf.EntityData.Leafs["validate-source-disable"] = types.YLeaf{"ValidateSourceDisable", defaultVrf.ValidateSourceDisable}
    defaultVrf.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", defaultVrf.MaximumPaths}
    defaultVrf.EntityData.Leafs["nsf"] = types.YLeaf{"Nsf", defaultVrf.Nsf}
    defaultVrf.EntityData.Leafs["policy-in"] = types.YLeaf{"PolicyIn", defaultVrf.PolicyIn}
    return &(defaultVrf.EntityData)
}

// Rip_DefaultVrf_DefaultInformation
// Controls default information origination
// This type is a presence type.
type Rip_DefaultVrf_DefaultInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy name. The type is string.
    RoutePolicyName interface{}

    // Origination option. The type is DefaultInformationOption. This attribute is
    // mandatory.
    Option interface{}
}

func (defaultInformation *Rip_DefaultVrf_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "default-vrf"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = make(map[string]types.YChild)
    defaultInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultInformation.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", defaultInformation.RoutePolicyName}
    defaultInformation.EntityData.Leafs["option"] = types.YLeaf{"Option", defaultInformation.Option}
    return &(defaultInformation.EntityData)
}

// Rip_DefaultVrf_Redistribution
// Redistribute information from another routing
// protocol
type Rip_DefaultVrf_Redistribution struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute connected routes.
    Connected Rip_DefaultVrf_Redistribution_Connected

    // Redistribute BGP routes.
    Bgps Rip_DefaultVrf_Redistribution_Bgps

    // Redistribute IS-IS routes.
    Isises Rip_DefaultVrf_Redistribution_Isises

    // Redistribute EIGRP routes.
    EigrpS Rip_DefaultVrf_Redistribution_EigrpS

    // Redistribute static routes.
    Static Rip_DefaultVrf_Redistribution_Static

    // Redistribute OSPF routes.
    Ospfs Rip_DefaultVrf_Redistribution_Ospfs
}

func (redistribution *Rip_DefaultVrf_Redistribution) GetEntityData() *types.CommonEntityData {
    redistribution.EntityData.YFilter = redistribution.YFilter
    redistribution.EntityData.YangName = "redistribution"
    redistribution.EntityData.BundleName = "cisco_ios_xr"
    redistribution.EntityData.ParentYangName = "default-vrf"
    redistribution.EntityData.SegmentPath = "redistribution"
    redistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribution.EntityData.Children = make(map[string]types.YChild)
    redistribution.EntityData.Children["connected"] = types.YChild{"Connected", &redistribution.Connected}
    redistribution.EntityData.Children["bgps"] = types.YChild{"Bgps", &redistribution.Bgps}
    redistribution.EntityData.Children["isises"] = types.YChild{"Isises", &redistribution.Isises}
    redistribution.EntityData.Children["eigrp-s"] = types.YChild{"EigrpS", &redistribution.EigrpS}
    redistribution.EntityData.Children["static"] = types.YChild{"Static", &redistribution.Static}
    redistribution.EntityData.Children["ospfs"] = types.YChild{"Ospfs", &redistribution.Ospfs}
    redistribution.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redistribution.EntityData)
}

// Rip_DefaultVrf_Redistribution_Connected
// Redistribute connected routes
// This type is a presence type.
type Rip_DefaultVrf_Redistribution_Connected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (connected *Rip_DefaultVrf_Redistribution_Connected) GetEntityData() *types.CommonEntityData {
    connected.EntityData.YFilter = connected.YFilter
    connected.EntityData.YangName = "connected"
    connected.EntityData.BundleName = "cisco_ios_xr"
    connected.EntityData.ParentYangName = "redistribution"
    connected.EntityData.SegmentPath = "connected"
    connected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connected.EntityData.Children = make(map[string]types.YChild)
    connected.EntityData.Leafs = make(map[string]types.YLeaf)
    connected.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", connected.RoutePolicyName}
    connected.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", connected.RouteType}
    return &(connected.EntityData)
}

// Rip_DefaultVrf_Redistribution_Bgps
// Redistribute BGP routes
type Rip_DefaultVrf_Redistribution_Bgps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous system number. The type is slice of
    // Rip_DefaultVrf_Redistribution_Bgps_Bgp.
    Bgp []Rip_DefaultVrf_Redistribution_Bgps_Bgp
}

func (bgps *Rip_DefaultVrf_Redistribution_Bgps) GetEntityData() *types.CommonEntityData {
    bgps.EntityData.YFilter = bgps.YFilter
    bgps.EntityData.YangName = "bgps"
    bgps.EntityData.BundleName = "cisco_ios_xr"
    bgps.EntityData.ParentYangName = "redistribution"
    bgps.EntityData.SegmentPath = "bgps"
    bgps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgps.EntityData.Children = make(map[string]types.YChild)
    bgps.EntityData.Children["bgp"] = types.YChild{"Bgp", nil}
    for i := range bgps.Bgp {
        bgps.EntityData.Children[types.GetSegmentPath(&bgps.Bgp[i])] = types.YChild{"Bgp", &bgps.Bgp[i]}
    }
    bgps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgps.EntityData)
}

// Rip_DefaultVrf_Redistribution_Bgps_Bgp
// Autonomous system number
type Rip_DefaultVrf_Redistribution_Bgps_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Higher 16 bits of 4-byte BGP AS number. The type
    // is interface{} with range: 0..65535.
    Asnxx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS number. The type is
    // interface{} with range: 0..4294967295.
    Asnyy interface{}

    // Route Policy name. The type is string.
    Policy interface{}

    // Route type. The type is BgpRedistRoute.
    Type_ interface{}
}

func (bgp *Rip_DefaultVrf_Redistribution_Bgps_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "bgps"
    bgp.EntityData.SegmentPath = "bgp" + "[asnxx='" + fmt.Sprintf("%v", bgp.Asnxx) + "']" + "[asnyy='" + fmt.Sprintf("%v", bgp.Asnyy) + "']"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["asnxx"] = types.YLeaf{"Asnxx", bgp.Asnxx}
    bgp.EntityData.Leafs["asnyy"] = types.YLeaf{"Asnyy", bgp.Asnyy}
    bgp.EntityData.Leafs["policy"] = types.YLeaf{"Policy", bgp.Policy}
    bgp.EntityData.Leafs["type"] = types.YLeaf{"Type_", bgp.Type_}
    return &(bgp.EntityData)
}

// Rip_DefaultVrf_Redistribution_Isises
// Redistribute IS-IS routes
type Rip_DefaultVrf_Redistribution_Isises struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute IS-IS routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_Isises_Isis.
    Isis []Rip_DefaultVrf_Redistribution_Isises_Isis
}

func (isises *Rip_DefaultVrf_Redistribution_Isises) GetEntityData() *types.CommonEntityData {
    isises.EntityData.YFilter = isises.YFilter
    isises.EntityData.YangName = "isises"
    isises.EntityData.BundleName = "cisco_ios_xr"
    isises.EntityData.ParentYangName = "redistribution"
    isises.EntityData.SegmentPath = "isises"
    isises.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isises.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isises.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isises.EntityData.Children = make(map[string]types.YChild)
    isises.EntityData.Children["isis"] = types.YChild{"Isis", nil}
    for i := range isises.Isis {
        isises.EntityData.Children[types.GetSegmentPath(&isises.Isis[i])] = types.YChild{"Isis", &isises.Isis[i]}
    }
    isises.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(isises.EntityData)
}

// Rip_DefaultVrf_Redistribution_Isises_Isis
// Redistribute IS-IS routes
type Rip_DefaultVrf_Redistribution_Isises_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IS-IS instance name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    IsisName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is IsisRedistRoute.
    RouteType interface{}
}

func (isis *Rip_DefaultVrf_Redistribution_Isises_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "isises"
    isis.EntityData.SegmentPath = "isis" + "[isis-name='" + fmt.Sprintf("%v", isis.IsisName) + "']"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["isis-name"] = types.YLeaf{"IsisName", isis.IsisName}
    isis.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", isis.RoutePolicyName}
    isis.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", isis.RouteType}
    return &(isis.EntityData)
}

// Rip_DefaultVrf_Redistribution_EigrpS
// Redistribute EIGRP routes
type Rip_DefaultVrf_Redistribution_EigrpS struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute EIGRP routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_EigrpS_Eigrp.
    Eigrp []Rip_DefaultVrf_Redistribution_EigrpS_Eigrp
}

func (eigrpS *Rip_DefaultVrf_Redistribution_EigrpS) GetEntityData() *types.CommonEntityData {
    eigrpS.EntityData.YFilter = eigrpS.YFilter
    eigrpS.EntityData.YangName = "eigrp-s"
    eigrpS.EntityData.BundleName = "cisco_ios_xr"
    eigrpS.EntityData.ParentYangName = "redistribution"
    eigrpS.EntityData.SegmentPath = "eigrp-s"
    eigrpS.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpS.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpS.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpS.EntityData.Children = make(map[string]types.YChild)
    eigrpS.EntityData.Children["eigrp"] = types.YChild{"Eigrp", nil}
    for i := range eigrpS.Eigrp {
        eigrpS.EntityData.Children[types.GetSegmentPath(&eigrpS.Eigrp[i])] = types.YChild{"Eigrp", &eigrpS.Eigrp[i]}
    }
    eigrpS.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eigrpS.EntityData)
}

// Rip_DefaultVrf_Redistribution_EigrpS_Eigrp
// Redistribute EIGRP routes
type Rip_DefaultVrf_Redistribution_EigrpS_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Autonomous system number. The type is interface{}
    // with range: 1..65535.
    As interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (eigrp *Rip_DefaultVrf_Redistribution_EigrpS_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "eigrp-s"
    eigrp.EntityData.SegmentPath = "eigrp" + "[as='" + fmt.Sprintf("%v", eigrp.As) + "']"
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = make(map[string]types.YChild)
    eigrp.EntityData.Leafs = make(map[string]types.YLeaf)
    eigrp.EntityData.Leafs["as"] = types.YLeaf{"As", eigrp.As}
    eigrp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName}
    eigrp.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", eigrp.RouteType}
    return &(eigrp.EntityData)
}

// Rip_DefaultVrf_Redistribution_Static
// Redistribute static routes
// This type is a presence type.
type Rip_DefaultVrf_Redistribution_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (static *Rip_DefaultVrf_Redistribution_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xr"
    static.EntityData.ParentYangName = "redistribution"
    static.EntityData.SegmentPath = "static"
    static.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    static.EntityData.Children = make(map[string]types.YChild)
    static.EntityData.Leafs = make(map[string]types.YLeaf)
    static.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", static.RoutePolicyName}
    static.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", static.RouteType}
    return &(static.EntityData)
}

// Rip_DefaultVrf_Redistribution_Ospfs
// Redistribute OSPF routes
type Rip_DefaultVrf_Redistribution_Ospfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute OSPF routes. The type is slice of
    // Rip_DefaultVrf_Redistribution_Ospfs_Ospf.
    Ospf []Rip_DefaultVrf_Redistribution_Ospfs_Ospf
}

func (ospfs *Rip_DefaultVrf_Redistribution_Ospfs) GetEntityData() *types.CommonEntityData {
    ospfs.EntityData.YFilter = ospfs.YFilter
    ospfs.EntityData.YangName = "ospfs"
    ospfs.EntityData.BundleName = "cisco_ios_xr"
    ospfs.EntityData.ParentYangName = "redistribution"
    ospfs.EntityData.SegmentPath = "ospfs"
    ospfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfs.EntityData.Children = make(map[string]types.YChild)
    ospfs.EntityData.Children["ospf"] = types.YChild{"Ospf", nil}
    for i := range ospfs.Ospf {
        ospfs.EntityData.Children[types.GetSegmentPath(&ospfs.Ospf[i])] = types.YChild{"Ospf", &ospfs.Ospf[i]}
    }
    ospfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfs.EntityData)
}

// Rip_DefaultVrf_Redistribution_Ospfs_Ospf
// Redistribute OSPF routes
type Rip_DefaultVrf_Redistribution_Ospfs_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID for the OSPF instance. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    OspfName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Internal routes. The type is bool.
    Internal interface{}

    // External routes. The type is bool.
    External interface{}

    // External route type. The type is interface{} with range: 0..2.
    ExternalType interface{}

    // NSSA External routes. The type is bool.
    NssaExternal interface{}

    // NSSA External route type. The type is interface{} with range: 0..2.
    NssaExternalType interface{}
}

func (ospf *Rip_DefaultVrf_Redistribution_Ospfs_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "ospfs"
    ospf.EntityData.SegmentPath = "ospf" + "[ospf-name='" + fmt.Sprintf("%v", ospf.OspfName) + "']"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["ospf-name"] = types.YLeaf{"OspfName", ospf.OspfName}
    ospf.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", ospf.RoutePolicyName}
    ospf.EntityData.Leafs["internal"] = types.YLeaf{"Internal", ospf.Internal}
    ospf.EntityData.Leafs["external"] = types.YLeaf{"External", ospf.External}
    ospf.EntityData.Leafs["external-type"] = types.YLeaf{"ExternalType", ospf.ExternalType}
    ospf.EntityData.Leafs["nssa-external"] = types.YLeaf{"NssaExternal", ospf.NssaExternal}
    ospf.EntityData.Leafs["nssa-external-type"] = types.YLeaf{"NssaExternalType", ospf.NssaExternalType}
    return &(ospf.EntityData)
}

// Rip_DefaultVrf_IpDistances
// Table of IP specific administrative distances
type Rip_DefaultVrf_IpDistances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP specific administrative distance. The type is slice of
    // Rip_DefaultVrf_IpDistances_IpDistance.
    IpDistance []Rip_DefaultVrf_IpDistances_IpDistance
}

func (ipDistances *Rip_DefaultVrf_IpDistances) GetEntityData() *types.CommonEntityData {
    ipDistances.EntityData.YFilter = ipDistances.YFilter
    ipDistances.EntityData.YangName = "ip-distances"
    ipDistances.EntityData.BundleName = "cisco_ios_xr"
    ipDistances.EntityData.ParentYangName = "default-vrf"
    ipDistances.EntityData.SegmentPath = "ip-distances"
    ipDistances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDistances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDistances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDistances.EntityData.Children = make(map[string]types.YChild)
    ipDistances.EntityData.Children["ip-distance"] = types.YChild{"IpDistance", nil}
    for i := range ipDistances.IpDistance {
        ipDistances.EntityData.Children[types.GetSegmentPath(&ipDistances.IpDistance[i])] = types.YChild{"IpDistance", &ipDistances.IpDistance[i]}
    }
    ipDistances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipDistances.EntityData)
}

// Rip_DefaultVrf_IpDistances_IpDistance
// IP specific administrative distance
type Rip_DefaultVrf_IpDistances_IpDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP Source address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. IP address mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Netmask interface{}

    // Administrative distance. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (ipDistance *Rip_DefaultVrf_IpDistances_IpDistance) GetEntityData() *types.CommonEntityData {
    ipDistance.EntityData.YFilter = ipDistance.YFilter
    ipDistance.EntityData.YangName = "ip-distance"
    ipDistance.EntityData.BundleName = "cisco_ios_xr"
    ipDistance.EntityData.ParentYangName = "ip-distances"
    ipDistance.EntityData.SegmentPath = "ip-distance" + "[address='" + fmt.Sprintf("%v", ipDistance.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", ipDistance.Netmask) + "']"
    ipDistance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDistance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDistance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDistance.EntityData.Children = make(map[string]types.YChild)
    ipDistance.EntityData.Leafs = make(map[string]types.YLeaf)
    ipDistance.EntityData.Leafs["address"] = types.YLeaf{"Address", ipDistance.Address}
    ipDistance.EntityData.Leafs["netmask"] = types.YLeaf{"Netmask", ipDistance.Netmask}
    ipDistance.EntityData.Leafs["distance"] = types.YLeaf{"Distance", ipDistance.Distance}
    return &(ipDistance.EntityData)
}

// Rip_DefaultVrf_Interfaces
// Table of RIP interfaces
type Rip_DefaultVrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP interface name. The type is slice of
    // Rip_DefaultVrf_Interfaces_Interface_.
    Interface_ []Rip_DefaultVrf_Interfaces_Interface
}

func (interfaces *Rip_DefaultVrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "default-vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface
// RIP interface name
type Rip_DefaultVrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Enable poison reverse. The type is interface{}.
    PoisonReverse interface{}

    // Suppress routing updates on this interface. The type is interface{}.
    Passive interface{}

    // Starts RIP interface configuration. The type is interface{}.
    Enable interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Accept RIP updates with metric 0. The type is interface{}.
    AcceptMetricZero interface{}

    // Route Policy for inbound routing updates. The type is string.
    PolicyIn interface{}

    // Disable split horizon. The type is interface{}.
    SplitHorizonDisable interface{}

    // Authentication keychain and mode.
    Authentication Rip_DefaultVrf_Interfaces_Interface_Authentication

    // SOO community for prefixes learned over this interface.
    SiteOfOrigin Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin

    // RIP versions supported for receiving advertisements.
    ReceiveVersion Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion

    // RIP versions supported for sending advertisements.
    SendVersion Rip_DefaultVrf_Interfaces_Interface_SendVersion
}

func (self *Rip_DefaultVrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["authentication"] = types.YChild{"Authentication", &self.Authentication}
    self.EntityData.Children["site-of-origin"] = types.YChild{"SiteOfOrigin", &self.SiteOfOrigin}
    self.EntityData.Children["receive-version"] = types.YChild{"ReceiveVersion", &self.ReceiveVersion}
    self.EntityData.Children["send-version"] = types.YChild{"SendVersion", &self.SendVersion}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["broadcast-for-v2"] = types.YLeaf{"BroadcastForV2", self.BroadcastForV2}
    self.EntityData.Leafs["poison-reverse"] = types.YLeaf{"PoisonReverse", self.PoisonReverse}
    self.EntityData.Leafs["passive"] = types.YLeaf{"Passive", self.Passive}
    self.EntityData.Leafs["enable"] = types.YLeaf{"Enable", self.Enable}
    self.EntityData.Leafs["policy-out"] = types.YLeaf{"PolicyOut", self.PolicyOut}
    self.EntityData.Leafs["accept-metric-zero"] = types.YLeaf{"AcceptMetricZero", self.AcceptMetricZero}
    self.EntityData.Leafs["policy-in"] = types.YLeaf{"PolicyIn", self.PolicyIn}
    self.EntityData.Leafs["split-horizon-disable"] = types.YLeaf{"SplitHorizonDisable", self.SplitHorizonDisable}
    return &(self.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_Authentication
// Authentication keychain and mode
// This type is a presence type.
type Rip_DefaultVrf_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of keychain. The type is string. This attribute is mandatory.
    Keychain interface{}

    // Authentication mode. The type is RipAuthMode. This attribute is mandatory.
    Mode interface{}
}

func (authentication *Rip_DefaultVrf_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["keychain"] = types.YLeaf{"Keychain", authentication.Keychain}
    authentication.EntityData.Leafs["mode"] = types.YLeaf{"Mode", authentication.Mode}
    return &(authentication.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin
// SOO community for prefixes learned over this
// interface
type Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Extended community type. The type is RipExtCommunity.
    Type_ interface{}

    // AS Number for AS:nn format. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // 32 bit value for AS:nn format. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address for IPV4Address:nn format. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // 16bit value for IPV4Address:nn format. The type is interface{} with range:
    // 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Rip_DefaultVrf_Interfaces_Interface_SiteOfOrigin) GetEntityData() *types.CommonEntityData {
    siteOfOrigin.EntityData.YFilter = siteOfOrigin.YFilter
    siteOfOrigin.EntityData.YangName = "site-of-origin"
    siteOfOrigin.EntityData.BundleName = "cisco_ios_xr"
    siteOfOrigin.EntityData.ParentYangName = "interface"
    siteOfOrigin.EntityData.SegmentPath = "site-of-origin"
    siteOfOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siteOfOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siteOfOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siteOfOrigin.EntityData.Children = make(map[string]types.YChild)
    siteOfOrigin.EntityData.Leafs = make(map[string]types.YLeaf)
    siteOfOrigin.EntityData.Leafs["type"] = types.YLeaf{"Type_", siteOfOrigin.Type_}
    siteOfOrigin.EntityData.Leafs["as-xx"] = types.YLeaf{"AsXx", siteOfOrigin.AsXx}
    siteOfOrigin.EntityData.Leafs["as-yy"] = types.YLeaf{"AsYy", siteOfOrigin.AsYy}
    siteOfOrigin.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", siteOfOrigin.AsIndex}
    siteOfOrigin.EntityData.Leafs["address"] = types.YLeaf{"Address", siteOfOrigin.Address}
    siteOfOrigin.EntityData.Leafs["address-index"] = types.YLeaf{"AddressIndex", siteOfOrigin.AddressIndex}
    return &(siteOfOrigin.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion
// RIP versions supported for receiving
// advertisements
type Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (receiveVersion *Rip_DefaultVrf_Interfaces_Interface_ReceiveVersion) GetEntityData() *types.CommonEntityData {
    receiveVersion.EntityData.YFilter = receiveVersion.YFilter
    receiveVersion.EntityData.YangName = "receive-version"
    receiveVersion.EntityData.BundleName = "cisco_ios_xr"
    receiveVersion.EntityData.ParentYangName = "interface"
    receiveVersion.EntityData.SegmentPath = "receive-version"
    receiveVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveVersion.EntityData.Children = make(map[string]types.YChild)
    receiveVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    receiveVersion.EntityData.Leafs["version1"] = types.YLeaf{"Version1", receiveVersion.Version1}
    receiveVersion.EntityData.Leafs["version2"] = types.YLeaf{"Version2", receiveVersion.Version2}
    return &(receiveVersion.EntityData)
}

// Rip_DefaultVrf_Interfaces_Interface_SendVersion
// RIP versions supported for sending
// advertisements
type Rip_DefaultVrf_Interfaces_Interface_SendVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (sendVersion *Rip_DefaultVrf_Interfaces_Interface_SendVersion) GetEntityData() *types.CommonEntityData {
    sendVersion.EntityData.YFilter = sendVersion.YFilter
    sendVersion.EntityData.YangName = "send-version"
    sendVersion.EntityData.BundleName = "cisco_ios_xr"
    sendVersion.EntityData.ParentYangName = "interface"
    sendVersion.EntityData.SegmentPath = "send-version"
    sendVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendVersion.EntityData.Children = make(map[string]types.YChild)
    sendVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    sendVersion.EntityData.Leafs["version1"] = types.YLeaf{"Version1", sendVersion.Version1}
    sendVersion.EntityData.Leafs["version2"] = types.YLeaf{"Version2", sendVersion.Version2}
    return &(sendVersion.EntityData)
}

// Rip_DefaultVrf_Neighbors
// Configure RIP Neighbors
type Rip_DefaultVrf_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor address. The type is slice of Rip_DefaultVrf_Neighbors_Neighbor.
    Neighbor []Rip_DefaultVrf_Neighbors_Neighbor
}

func (neighbors *Rip_DefaultVrf_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "default-vrf"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Rip_DefaultVrf_Neighbors_Neighbor
// Neighbor address
type Rip_DefaultVrf_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}
}

func (neighbor *Rip_DefaultVrf_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", neighbor.NeighborAddress}
    return &(neighbor.EntityData)
}

// Rip_DefaultVrf_Timers
// Various routing timers
// This type is a presence type.
type Rip_DefaultVrf_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between updates. The type is interface{} with range: 5..50000.
    // This attribute is mandatory.
    UpdateTimer interface{}

    // Invalid. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    InvalidTimer interface{}

    // Holddown. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    HolddownTimer interface{}

    // Flush. The type is interface{} with range: 16..250000. This attribute is
    // mandatory.
    FlushTimer interface{}
}

func (timers *Rip_DefaultVrf_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "default-vrf"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", timers.UpdateTimer}
    timers.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", timers.InvalidTimer}
    timers.EntityData.Leafs["holddown-timer"] = types.YLeaf{"HolddownTimer", timers.HolddownTimer}
    timers.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", timers.FlushTimer}
    return &(timers.EntityData)
}

// Rip_Vrfs
// VRF related RIP configuration
type Rip_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP configuration for a particular VRF. The type is slice of Rip_Vrfs_Vrf.
    Vrf []Rip_Vrfs_Vrf
}

func (vrfs *Rip_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "rip"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Rip_Vrfs_Vrf
// RIP configuration for a particular VRF
type Rip_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Starts RIP configuration for a particular VRF. The type is interface{}.
    Enable interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Administrative distance. The type is interface{} with range: 0..255. The
    // default value is 120.
    Distance interface{}

    // Default metric of redistributed routes. The type is interface{} with range:
    // 0..16.
    DefaultMetric interface{}

    // Inter-packet delay for RIP updates. The type is interface{} with range:
    // 8..50. Units are millisecond.
    OutputDelay interface{}

    // Enable automatic network number summarization. The type is interface{}.
    AutoSummary interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Disable validation of source address of routing updates. The type is
    // interface{}.
    ValidateSourceDisable interface{}

    // Maximum number of paths allowed per route. The type is interface{} with
    // range: 1..128. The default value is 4.
    MaximumPaths interface{}

    // Enable Cisco Non Stop Forwarding. The type is interface{}.
    Nsf interface{}

    // Route Policy for inbbound routing updates. The type is string.
    PolicyIn interface{}

    // Controls default information origination.
    DefaultInformation Rip_Vrfs_Vrf_DefaultInformation

    // Redistribute information from another routing protocol.
    Redistribution Rip_Vrfs_Vrf_Redistribution

    // Table of IP specific administrative distances.
    IpDistances Rip_Vrfs_Vrf_IpDistances

    // Table of RIP interfaces.
    Interfaces Rip_Vrfs_Vrf_Interfaces

    // Configure RIP Neighbors.
    Neighbors Rip_Vrfs_Vrf_Neighbors

    // Various routing timers.
    Timers Rip_Vrfs_Vrf_Timers
}

func (vrf *Rip_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["default-information"] = types.YChild{"DefaultInformation", &vrf.DefaultInformation}
    vrf.EntityData.Children["redistribution"] = types.YChild{"Redistribution", &vrf.Redistribution}
    vrf.EntityData.Children["ip-distances"] = types.YChild{"IpDistances", &vrf.IpDistances}
    vrf.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &vrf.Interfaces}
    vrf.EntityData.Children["neighbors"] = types.YChild{"Neighbors", &vrf.Neighbors}
    vrf.EntityData.Children["timers"] = types.YChild{"Timers", &vrf.Timers}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vrf.Enable}
    vrf.EntityData.Leafs["broadcast-for-v2"] = types.YLeaf{"BroadcastForV2", vrf.BroadcastForV2}
    vrf.EntityData.Leafs["distance"] = types.YLeaf{"Distance", vrf.Distance}
    vrf.EntityData.Leafs["default-metric"] = types.YLeaf{"DefaultMetric", vrf.DefaultMetric}
    vrf.EntityData.Leafs["output-delay"] = types.YLeaf{"OutputDelay", vrf.OutputDelay}
    vrf.EntityData.Leafs["auto-summary"] = types.YLeaf{"AutoSummary", vrf.AutoSummary}
    vrf.EntityData.Leafs["policy-out"] = types.YLeaf{"PolicyOut", vrf.PolicyOut}
    vrf.EntityData.Leafs["validate-source-disable"] = types.YLeaf{"ValidateSourceDisable", vrf.ValidateSourceDisable}
    vrf.EntityData.Leafs["maximum-paths"] = types.YLeaf{"MaximumPaths", vrf.MaximumPaths}
    vrf.EntityData.Leafs["nsf"] = types.YLeaf{"Nsf", vrf.Nsf}
    vrf.EntityData.Leafs["policy-in"] = types.YLeaf{"PolicyIn", vrf.PolicyIn}
    return &(vrf.EntityData)
}

// Rip_Vrfs_Vrf_DefaultInformation
// Controls default information origination
// This type is a presence type.
type Rip_Vrfs_Vrf_DefaultInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route policy name. The type is string.
    RoutePolicyName interface{}

    // Origination option. The type is DefaultInformationOption. This attribute is
    // mandatory.
    Option interface{}
}

func (defaultInformation *Rip_Vrfs_Vrf_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "vrf"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = make(map[string]types.YChild)
    defaultInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultInformation.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", defaultInformation.RoutePolicyName}
    defaultInformation.EntityData.Leafs["option"] = types.YLeaf{"Option", defaultInformation.Option}
    return &(defaultInformation.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution
// Redistribute information from another routing
// protocol
type Rip_Vrfs_Vrf_Redistribution struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute connected routes.
    Connected Rip_Vrfs_Vrf_Redistribution_Connected

    // Redistribute BGP routes.
    Bgps Rip_Vrfs_Vrf_Redistribution_Bgps

    // Redistribute IS-IS routes.
    Isises Rip_Vrfs_Vrf_Redistribution_Isises

    // Redistribute EIGRP routes.
    EigrpS Rip_Vrfs_Vrf_Redistribution_EigrpS

    // Redistribute static routes.
    Static Rip_Vrfs_Vrf_Redistribution_Static

    // Redistribute OSPF routes.
    Ospfs Rip_Vrfs_Vrf_Redistribution_Ospfs
}

func (redistribution *Rip_Vrfs_Vrf_Redistribution) GetEntityData() *types.CommonEntityData {
    redistribution.EntityData.YFilter = redistribution.YFilter
    redistribution.EntityData.YangName = "redistribution"
    redistribution.EntityData.BundleName = "cisco_ios_xr"
    redistribution.EntityData.ParentYangName = "vrf"
    redistribution.EntityData.SegmentPath = "redistribution"
    redistribution.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribution.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribution.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribution.EntityData.Children = make(map[string]types.YChild)
    redistribution.EntityData.Children["connected"] = types.YChild{"Connected", &redistribution.Connected}
    redistribution.EntityData.Children["bgps"] = types.YChild{"Bgps", &redistribution.Bgps}
    redistribution.EntityData.Children["isises"] = types.YChild{"Isises", &redistribution.Isises}
    redistribution.EntityData.Children["eigrp-s"] = types.YChild{"EigrpS", &redistribution.EigrpS}
    redistribution.EntityData.Children["static"] = types.YChild{"Static", &redistribution.Static}
    redistribution.EntityData.Children["ospfs"] = types.YChild{"Ospfs", &redistribution.Ospfs}
    redistribution.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(redistribution.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Connected
// Redistribute connected routes
// This type is a presence type.
type Rip_Vrfs_Vrf_Redistribution_Connected struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (connected *Rip_Vrfs_Vrf_Redistribution_Connected) GetEntityData() *types.CommonEntityData {
    connected.EntityData.YFilter = connected.YFilter
    connected.EntityData.YangName = "connected"
    connected.EntityData.BundleName = "cisco_ios_xr"
    connected.EntityData.ParentYangName = "redistribution"
    connected.EntityData.SegmentPath = "connected"
    connected.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connected.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connected.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connected.EntityData.Children = make(map[string]types.YChild)
    connected.EntityData.Leafs = make(map[string]types.YLeaf)
    connected.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", connected.RoutePolicyName}
    connected.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", connected.RouteType}
    return &(connected.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Bgps
// Redistribute BGP routes
type Rip_Vrfs_Vrf_Redistribution_Bgps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous system number. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp.
    Bgp []Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp
}

func (bgps *Rip_Vrfs_Vrf_Redistribution_Bgps) GetEntityData() *types.CommonEntityData {
    bgps.EntityData.YFilter = bgps.YFilter
    bgps.EntityData.YangName = "bgps"
    bgps.EntityData.BundleName = "cisco_ios_xr"
    bgps.EntityData.ParentYangName = "redistribution"
    bgps.EntityData.SegmentPath = "bgps"
    bgps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgps.EntityData.Children = make(map[string]types.YChild)
    bgps.EntityData.Children["bgp"] = types.YChild{"Bgp", nil}
    for i := range bgps.Bgp {
        bgps.EntityData.Children[types.GetSegmentPath(&bgps.Bgp[i])] = types.YChild{"Bgp", &bgps.Bgp[i]}
    }
    bgps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgps.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp
// Autonomous system number
type Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Higher 16 bits of 4-byte BGP AS number. The type
    // is interface{} with range: 0..65535.
    Asnxx interface{}

    // This attribute is a key. 2-byte or 4-byte BGP AS number. The type is
    // interface{} with range: 0..4294967295.
    Asnyy interface{}

    // Route Policy name. The type is string.
    Policy interface{}

    // Route type. The type is BgpRedistRoute.
    Type_ interface{}
}

func (bgp *Rip_Vrfs_Vrf_Redistribution_Bgps_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "bgps"
    bgp.EntityData.SegmentPath = "bgp" + "[asnxx='" + fmt.Sprintf("%v", bgp.Asnxx) + "']" + "[asnyy='" + fmt.Sprintf("%v", bgp.Asnyy) + "']"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["asnxx"] = types.YLeaf{"Asnxx", bgp.Asnxx}
    bgp.EntityData.Leafs["asnyy"] = types.YLeaf{"Asnyy", bgp.Asnyy}
    bgp.EntityData.Leafs["policy"] = types.YLeaf{"Policy", bgp.Policy}
    bgp.EntityData.Leafs["type"] = types.YLeaf{"Type_", bgp.Type_}
    return &(bgp.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Isises
// Redistribute IS-IS routes
type Rip_Vrfs_Vrf_Redistribution_Isises struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute IS-IS routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Isises_Isis.
    Isis []Rip_Vrfs_Vrf_Redistribution_Isises_Isis
}

func (isises *Rip_Vrfs_Vrf_Redistribution_Isises) GetEntityData() *types.CommonEntityData {
    isises.EntityData.YFilter = isises.YFilter
    isises.EntityData.YangName = "isises"
    isises.EntityData.BundleName = "cisco_ios_xr"
    isises.EntityData.ParentYangName = "redistribution"
    isises.EntityData.SegmentPath = "isises"
    isises.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isises.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isises.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isises.EntityData.Children = make(map[string]types.YChild)
    isises.EntityData.Children["isis"] = types.YChild{"Isis", nil}
    for i := range isises.Isis {
        isises.EntityData.Children[types.GetSegmentPath(&isises.Isis[i])] = types.YChild{"Isis", &isises.Isis[i]}
    }
    isises.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(isises.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Isises_Isis
// Redistribute IS-IS routes
type Rip_Vrfs_Vrf_Redistribution_Isises_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IS-IS instance name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    IsisName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is IsisRedistRoute.
    RouteType interface{}
}

func (isis *Rip_Vrfs_Vrf_Redistribution_Isises_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "isises"
    isis.EntityData.SegmentPath = "isis" + "[isis-name='" + fmt.Sprintf("%v", isis.IsisName) + "']"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["isis-name"] = types.YLeaf{"IsisName", isis.IsisName}
    isis.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", isis.RoutePolicyName}
    isis.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", isis.RouteType}
    return &(isis.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_EigrpS
// Redistribute EIGRP routes
type Rip_Vrfs_Vrf_Redistribution_EigrpS struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute EIGRP routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp.
    Eigrp []Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp
}

func (eigrpS *Rip_Vrfs_Vrf_Redistribution_EigrpS) GetEntityData() *types.CommonEntityData {
    eigrpS.EntityData.YFilter = eigrpS.YFilter
    eigrpS.EntityData.YangName = "eigrp-s"
    eigrpS.EntityData.BundleName = "cisco_ios_xr"
    eigrpS.EntityData.ParentYangName = "redistribution"
    eigrpS.EntityData.SegmentPath = "eigrp-s"
    eigrpS.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrpS.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrpS.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrpS.EntityData.Children = make(map[string]types.YChild)
    eigrpS.EntityData.Children["eigrp"] = types.YChild{"Eigrp", nil}
    for i := range eigrpS.Eigrp {
        eigrpS.EntityData.Children[types.GetSegmentPath(&eigrpS.Eigrp[i])] = types.YChild{"Eigrp", &eigrpS.Eigrp[i]}
    }
    eigrpS.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(eigrpS.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp
// Redistribute EIGRP routes
type Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Autonomous system number. The type is interface{}
    // with range: 1..65535.
    As interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (eigrp *Rip_Vrfs_Vrf_Redistribution_EigrpS_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "eigrp-s"
    eigrp.EntityData.SegmentPath = "eigrp" + "[as='" + fmt.Sprintf("%v", eigrp.As) + "']"
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = make(map[string]types.YChild)
    eigrp.EntityData.Leafs = make(map[string]types.YLeaf)
    eigrp.EntityData.Leafs["as"] = types.YLeaf{"As", eigrp.As}
    eigrp.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName}
    eigrp.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", eigrp.RouteType}
    return &(eigrp.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Static
// Redistribute static routes
// This type is a presence type.
type Rip_Vrfs_Vrf_Redistribution_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Route type. The type is DefaultRedistRoute.
    RouteType interface{}
}

func (static *Rip_Vrfs_Vrf_Redistribution_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xr"
    static.EntityData.ParentYangName = "redistribution"
    static.EntityData.SegmentPath = "static"
    static.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    static.EntityData.Children = make(map[string]types.YChild)
    static.EntityData.Leafs = make(map[string]types.YLeaf)
    static.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", static.RoutePolicyName}
    static.EntityData.Leafs["route-type"] = types.YLeaf{"RouteType", static.RouteType}
    return &(static.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Ospfs
// Redistribute OSPF routes
type Rip_Vrfs_Vrf_Redistribution_Ospfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute OSPF routes. The type is slice of
    // Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf.
    Ospf []Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf
}

func (ospfs *Rip_Vrfs_Vrf_Redistribution_Ospfs) GetEntityData() *types.CommonEntityData {
    ospfs.EntityData.YFilter = ospfs.YFilter
    ospfs.EntityData.YangName = "ospfs"
    ospfs.EntityData.BundleName = "cisco_ios_xr"
    ospfs.EntityData.ParentYangName = "redistribution"
    ospfs.EntityData.SegmentPath = "ospfs"
    ospfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfs.EntityData.Children = make(map[string]types.YChild)
    ospfs.EntityData.Children["ospf"] = types.YChild{"Ospf", nil}
    for i := range ospfs.Ospf {
        ospfs.EntityData.Children[types.GetSegmentPath(&ospfs.Ospf[i])] = types.YChild{"Ospf", &ospfs.Ospf[i]}
    }
    ospfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ospfs.EntityData)
}

// Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf
// Redistribute OSPF routes
type Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process ID for the OSPF instance. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    OspfName interface{}

    // Route Policy name. The type is string.
    RoutePolicyName interface{}

    // Internal routes. The type is bool.
    Internal interface{}

    // External routes. The type is bool.
    External interface{}

    // External route type. The type is interface{} with range: 0..2.
    ExternalType interface{}

    // NSSA External routes. The type is bool.
    NssaExternal interface{}

    // NSSA External route type. The type is interface{} with range: 0..2.
    NssaExternalType interface{}
}

func (ospf *Rip_Vrfs_Vrf_Redistribution_Ospfs_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "ospfs"
    ospf.EntityData.SegmentPath = "ospf" + "[ospf-name='" + fmt.Sprintf("%v", ospf.OspfName) + "']"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["ospf-name"] = types.YLeaf{"OspfName", ospf.OspfName}
    ospf.EntityData.Leafs["route-policy-name"] = types.YLeaf{"RoutePolicyName", ospf.RoutePolicyName}
    ospf.EntityData.Leafs["internal"] = types.YLeaf{"Internal", ospf.Internal}
    ospf.EntityData.Leafs["external"] = types.YLeaf{"External", ospf.External}
    ospf.EntityData.Leafs["external-type"] = types.YLeaf{"ExternalType", ospf.ExternalType}
    ospf.EntityData.Leafs["nssa-external"] = types.YLeaf{"NssaExternal", ospf.NssaExternal}
    ospf.EntityData.Leafs["nssa-external-type"] = types.YLeaf{"NssaExternalType", ospf.NssaExternalType}
    return &(ospf.EntityData)
}

// Rip_Vrfs_Vrf_IpDistances
// Table of IP specific administrative distances
type Rip_Vrfs_Vrf_IpDistances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP specific administrative distance. The type is slice of
    // Rip_Vrfs_Vrf_IpDistances_IpDistance.
    IpDistance []Rip_Vrfs_Vrf_IpDistances_IpDistance
}

func (ipDistances *Rip_Vrfs_Vrf_IpDistances) GetEntityData() *types.CommonEntityData {
    ipDistances.EntityData.YFilter = ipDistances.YFilter
    ipDistances.EntityData.YangName = "ip-distances"
    ipDistances.EntityData.BundleName = "cisco_ios_xr"
    ipDistances.EntityData.ParentYangName = "vrf"
    ipDistances.EntityData.SegmentPath = "ip-distances"
    ipDistances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDistances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDistances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDistances.EntityData.Children = make(map[string]types.YChild)
    ipDistances.EntityData.Children["ip-distance"] = types.YChild{"IpDistance", nil}
    for i := range ipDistances.IpDistance {
        ipDistances.EntityData.Children[types.GetSegmentPath(&ipDistances.IpDistance[i])] = types.YChild{"IpDistance", &ipDistances.IpDistance[i]}
    }
    ipDistances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipDistances.EntityData)
}

// Rip_Vrfs_Vrf_IpDistances_IpDistance
// IP specific administrative distance
type Rip_Vrfs_Vrf_IpDistances_IpDistance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP Source address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. IP address mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Netmask interface{}

    // Administrative distance. The type is interface{} with range: 0..255. This
    // attribute is mandatory.
    Distance interface{}
}

func (ipDistance *Rip_Vrfs_Vrf_IpDistances_IpDistance) GetEntityData() *types.CommonEntityData {
    ipDistance.EntityData.YFilter = ipDistance.YFilter
    ipDistance.EntityData.YangName = "ip-distance"
    ipDistance.EntityData.BundleName = "cisco_ios_xr"
    ipDistance.EntityData.ParentYangName = "ip-distances"
    ipDistance.EntityData.SegmentPath = "ip-distance" + "[address='" + fmt.Sprintf("%v", ipDistance.Address) + "']" + "[netmask='" + fmt.Sprintf("%v", ipDistance.Netmask) + "']"
    ipDistance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDistance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDistance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDistance.EntityData.Children = make(map[string]types.YChild)
    ipDistance.EntityData.Leafs = make(map[string]types.YLeaf)
    ipDistance.EntityData.Leafs["address"] = types.YLeaf{"Address", ipDistance.Address}
    ipDistance.EntityData.Leafs["netmask"] = types.YLeaf{"Netmask", ipDistance.Netmask}
    ipDistance.EntityData.Leafs["distance"] = types.YLeaf{"Distance", ipDistance.Distance}
    return &(ipDistance.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces
// Table of RIP interfaces
type Rip_Vrfs_Vrf_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RIP interface name. The type is slice of
    // Rip_Vrfs_Vrf_Interfaces_Interface_.
    Interface_ []Rip_Vrfs_Vrf_Interfaces_Interface
}

func (interfaces *Rip_Vrfs_Vrf_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrf"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface
// RIP interface name
type Rip_Vrfs_Vrf_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Send RIP v2 output packets to broadcast address. The type is interface{}.
    BroadcastForV2 interface{}

    // Enable poison reverse. The type is interface{}.
    PoisonReverse interface{}

    // Suppress routing updates on this interface. The type is interface{}.
    Passive interface{}

    // Starts RIP interface configuration. The type is interface{}.
    Enable interface{}

    // Route Policy for outbound routing updates. The type is string.
    PolicyOut interface{}

    // Accept RIP updates with metric 0. The type is interface{}.
    AcceptMetricZero interface{}

    // Route Policy for inbound routing updates. The type is string.
    PolicyIn interface{}

    // Disable split horizon. The type is interface{}.
    SplitHorizonDisable interface{}

    // Authentication keychain and mode.
    Authentication Rip_Vrfs_Vrf_Interfaces_Interface_Authentication

    // SOO community for prefixes learned over this interface.
    SiteOfOrigin Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin

    // RIP versions supported for receiving advertisements.
    ReceiveVersion Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion

    // RIP versions supported for sending advertisements.
    SendVersion Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion
}

func (self *Rip_Vrfs_Vrf_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["authentication"] = types.YChild{"Authentication", &self.Authentication}
    self.EntityData.Children["site-of-origin"] = types.YChild{"SiteOfOrigin", &self.SiteOfOrigin}
    self.EntityData.Children["receive-version"] = types.YChild{"ReceiveVersion", &self.ReceiveVersion}
    self.EntityData.Children["send-version"] = types.YChild{"SendVersion", &self.SendVersion}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["broadcast-for-v2"] = types.YLeaf{"BroadcastForV2", self.BroadcastForV2}
    self.EntityData.Leafs["poison-reverse"] = types.YLeaf{"PoisonReverse", self.PoisonReverse}
    self.EntityData.Leafs["passive"] = types.YLeaf{"Passive", self.Passive}
    self.EntityData.Leafs["enable"] = types.YLeaf{"Enable", self.Enable}
    self.EntityData.Leafs["policy-out"] = types.YLeaf{"PolicyOut", self.PolicyOut}
    self.EntityData.Leafs["accept-metric-zero"] = types.YLeaf{"AcceptMetricZero", self.AcceptMetricZero}
    self.EntityData.Leafs["policy-in"] = types.YLeaf{"PolicyIn", self.PolicyIn}
    self.EntityData.Leafs["split-horizon-disable"] = types.YLeaf{"SplitHorizonDisable", self.SplitHorizonDisable}
    return &(self.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_Authentication
// Authentication keychain and mode
// This type is a presence type.
type Rip_Vrfs_Vrf_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of keychain. The type is string. This attribute is mandatory.
    Keychain interface{}

    // Authentication mode. The type is RipAuthMode. This attribute is mandatory.
    Mode interface{}
}

func (authentication *Rip_Vrfs_Vrf_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["keychain"] = types.YLeaf{"Keychain", authentication.Keychain}
    authentication.EntityData.Leafs["mode"] = types.YLeaf{"Mode", authentication.Mode}
    return &(authentication.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin
// SOO community for prefixes learned over this
// interface
type Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Extended community type. The type is RipExtCommunity.
    Type_ interface{}

    // AS Number for AS:nn format. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // 32 bit value for AS:nn format. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // AS Number Index. The type is interface{} with range: 0..4294967295.
    AsIndex interface{}

    // IPV4 address for IPV4Address:nn format. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // 16bit value for IPV4Address:nn format. The type is interface{} with range:
    // 0..65535.
    AddressIndex interface{}
}

func (siteOfOrigin *Rip_Vrfs_Vrf_Interfaces_Interface_SiteOfOrigin) GetEntityData() *types.CommonEntityData {
    siteOfOrigin.EntityData.YFilter = siteOfOrigin.YFilter
    siteOfOrigin.EntityData.YangName = "site-of-origin"
    siteOfOrigin.EntityData.BundleName = "cisco_ios_xr"
    siteOfOrigin.EntityData.ParentYangName = "interface"
    siteOfOrigin.EntityData.SegmentPath = "site-of-origin"
    siteOfOrigin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    siteOfOrigin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    siteOfOrigin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    siteOfOrigin.EntityData.Children = make(map[string]types.YChild)
    siteOfOrigin.EntityData.Leafs = make(map[string]types.YLeaf)
    siteOfOrigin.EntityData.Leafs["type"] = types.YLeaf{"Type_", siteOfOrigin.Type_}
    siteOfOrigin.EntityData.Leafs["as-xx"] = types.YLeaf{"AsXx", siteOfOrigin.AsXx}
    siteOfOrigin.EntityData.Leafs["as-yy"] = types.YLeaf{"AsYy", siteOfOrigin.AsYy}
    siteOfOrigin.EntityData.Leafs["as-index"] = types.YLeaf{"AsIndex", siteOfOrigin.AsIndex}
    siteOfOrigin.EntityData.Leafs["address"] = types.YLeaf{"Address", siteOfOrigin.Address}
    siteOfOrigin.EntityData.Leafs["address-index"] = types.YLeaf{"AddressIndex", siteOfOrigin.AddressIndex}
    return &(siteOfOrigin.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion
// RIP versions supported for receiving
// advertisements
type Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (receiveVersion *Rip_Vrfs_Vrf_Interfaces_Interface_ReceiveVersion) GetEntityData() *types.CommonEntityData {
    receiveVersion.EntityData.YFilter = receiveVersion.YFilter
    receiveVersion.EntityData.YangName = "receive-version"
    receiveVersion.EntityData.BundleName = "cisco_ios_xr"
    receiveVersion.EntityData.ParentYangName = "interface"
    receiveVersion.EntityData.SegmentPath = "receive-version"
    receiveVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveVersion.EntityData.Children = make(map[string]types.YChild)
    receiveVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    receiveVersion.EntityData.Leafs["version1"] = types.YLeaf{"Version1", receiveVersion.Version1}
    receiveVersion.EntityData.Leafs["version2"] = types.YLeaf{"Version2", receiveVersion.Version2}
    return &(receiveVersion.EntityData)
}

// Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion
// RIP versions supported for sending
// advertisements
type Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Support RIP version 1. The type is bool.
    Version1 interface{}

    // Support RIP version 2. The type is bool. The default value is true.
    Version2 interface{}
}

func (sendVersion *Rip_Vrfs_Vrf_Interfaces_Interface_SendVersion) GetEntityData() *types.CommonEntityData {
    sendVersion.EntityData.YFilter = sendVersion.YFilter
    sendVersion.EntityData.YangName = "send-version"
    sendVersion.EntityData.BundleName = "cisco_ios_xr"
    sendVersion.EntityData.ParentYangName = "interface"
    sendVersion.EntityData.SegmentPath = "send-version"
    sendVersion.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendVersion.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendVersion.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendVersion.EntityData.Children = make(map[string]types.YChild)
    sendVersion.EntityData.Leafs = make(map[string]types.YLeaf)
    sendVersion.EntityData.Leafs["version1"] = types.YLeaf{"Version1", sendVersion.Version1}
    sendVersion.EntityData.Leafs["version2"] = types.YLeaf{"Version2", sendVersion.Version2}
    return &(sendVersion.EntityData)
}

// Rip_Vrfs_Vrf_Neighbors
// Configure RIP Neighbors
type Rip_Vrfs_Vrf_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Neighbor address. The type is slice of Rip_Vrfs_Vrf_Neighbors_Neighbor.
    Neighbor []Rip_Vrfs_Vrf_Neighbors_Neighbor
}

func (neighbors *Rip_Vrfs_Vrf_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "vrf"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = make(map[string]types.YChild)
    neighbors.EntityData.Children["neighbor"] = types.YChild{"Neighbor", nil}
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children[types.GetSegmentPath(&neighbors.Neighbor[i])] = types.YChild{"Neighbor", &neighbors.Neighbor[i]}
    }
    neighbors.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(neighbors.EntityData)
}

// Rip_Vrfs_Vrf_Neighbors_Neighbor
// Neighbor address
type Rip_Vrfs_Vrf_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NeighborAddress interface{}
}

func (neighbor *Rip_Vrfs_Vrf_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + "[neighbor-address='" + fmt.Sprintf("%v", neighbor.NeighborAddress) + "']"
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = make(map[string]types.YChild)
    neighbor.EntityData.Leafs = make(map[string]types.YLeaf)
    neighbor.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", neighbor.NeighborAddress}
    return &(neighbor.EntityData)
}

// Rip_Vrfs_Vrf_Timers
// Various routing timers
// This type is a presence type.
type Rip_Vrfs_Vrf_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between updates. The type is interface{} with range: 5..50000.
    // This attribute is mandatory.
    UpdateTimer interface{}

    // Invalid. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    InvalidTimer interface{}

    // Holddown. The type is interface{} with range: 15..200000. This attribute is
    // mandatory.
    HolddownTimer interface{}

    // Flush. The type is interface{} with range: 16..250000. This attribute is
    // mandatory.
    FlushTimer interface{}
}

func (timers *Rip_Vrfs_Vrf_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "vrf"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["update-timer"] = types.YLeaf{"UpdateTimer", timers.UpdateTimer}
    timers.EntityData.Leafs["invalid-timer"] = types.YLeaf{"InvalidTimer", timers.InvalidTimer}
    timers.EntityData.Leafs["holddown-timer"] = types.YLeaf{"HolddownTimer", timers.HolddownTimer}
    timers.EntityData.Leafs["flush-timer"] = types.YLeaf{"FlushTimer", timers.FlushTimer}
    return &(timers.EntityData)
}

