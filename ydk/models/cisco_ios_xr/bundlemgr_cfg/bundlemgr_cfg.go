// This module contains a collection of YANG definitions
// for Cisco IOS-XR bundlemgr package configuration.
// 
// This module contains definitions
// for the following management objects:
//   lacp: Link Aggregation Control Protocol commands
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-rgmgr-cfg,
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package bundlemgr_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bundlemgr_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-bundlemgr-cfg lacp}", reflect.TypeOf(Lacp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-bundlemgr-cfg:lacp", reflect.TypeOf(Lacp{}))
}

// BfdMode represents Bfd mode
type BfdMode string

const (
    // BFD mode not configured on per-bundle basis
    BfdMode_no_cfg BfdMode = "no-cfg"

    // BFD mode Cisco
    BfdMode_cisco BfdMode = "cisco"

    // BFD mode IETF
    BfdMode_ietf BfdMode = "ietf"
)

// ChurnLogging represents Churn logging
type ChurnLogging string

const (
    // Logging for actor churn only
    ChurnLogging_actor ChurnLogging = "actor"

    // Logging for partner churn only
    ChurnLogging_partner ChurnLogging = "partner"

    // Logging for actor and partner churn
    ChurnLogging_both ChurnLogging = "both"
)

// BundleMode represents Bundle mode
type BundleMode string

const (
    // On
    BundleMode_on BundleMode = "on"

    // Active
    BundleMode_active BundleMode = "active"

    // Passive
    BundleMode_passive BundleMode = "passive"
)

// BundleMinimumBandwidthRange represents Bundle minimum bandwidth range
type BundleMinimumBandwidthRange string

const (
    // None
    BundleMinimumBandwidthRange_none BundleMinimumBandwidthRange = "none"

    // kbps
    BundleMinimumBandwidthRange_kbps BundleMinimumBandwidthRange = "kbps"

    // mbps
    BundleMinimumBandwidthRange_mbps BundleMinimumBandwidthRange = "mbps"

    // gbps
    BundleMinimumBandwidthRange_gbps BundleMinimumBandwidthRange = "gbps"
)

// PeriodShortEnum represents Period short enum
type PeriodShortEnum string

const (
    // Use the standard LACP short period (1s)
    PeriodShortEnum_true PeriodShortEnum = "true"
)

// BundleCiscoExtTypes represents Bundle cisco ext types
type BundleCiscoExtTypes string

const (
    // LON signaling disabled
    BundleCiscoExtTypes_lon_signaling_off BundleCiscoExtTypes = "lon-signaling-off"

    // LON signaling enabled
    BundleCiscoExtTypes_lon_signaling_on BundleCiscoExtTypes = "lon-signaling-on"
)

// BundleMaximumActiveLinksMode represents Bundle maximum active links mode
type BundleMaximumActiveLinksMode string

const (
    // Default
    BundleMaximumActiveLinksMode_default_ BundleMaximumActiveLinksMode = "default"

    // Hot standby
    BundleMaximumActiveLinksMode_hot_standby BundleMaximumActiveLinksMode = "hot-standby"
)

// MlacpSwitchover represents Mlacp switchover
type MlacpSwitchover string

const (
    // Brute force shutdown
    MlacpSwitchover_brute_force MlacpSwitchover = "brute-force"

    // Revertive behavior
    MlacpSwitchover_revertive MlacpSwitchover = "revertive"
)

// BundleLoadBalance represents Bundle load balance
type BundleLoadBalance string

const (
    // Default hash function used
    BundleLoadBalance_default_ BundleLoadBalance = "default"

    // Send all traffic for this EFP over an
    // automatically selected member
    BundleLoadBalance_efp_auto BundleLoadBalance = "efp-auto"

    // Send all traffic for this EFP over the member
    // corresponding to the specified hash function
    BundleLoadBalance_efp_value BundleLoadBalance = "efp-value"

    // Load balance according to source IP address
    BundleLoadBalance_source_ip BundleLoadBalance = "source-ip"

    // Load balance according to detination IP address
    BundleLoadBalance_destination_ip BundleLoadBalance = "destination-ip"
)

// BundlePortActivity represents Bundle port activity
type BundlePortActivity string

const (
    // On
    BundlePortActivity_on BundlePortActivity = "on"

    // Active
    BundlePortActivity_active BundlePortActivity = "active"

    // Passive
    BundlePortActivity_passive BundlePortActivity = "passive"

    // Inherit
    BundlePortActivity_inherit BundlePortActivity = "inherit"
)

// MlacpMaximizeParameter represents Mlacp maximize parameter
type MlacpMaximizeParameter string

const (
    // Maximize the number of operational links
    MlacpMaximizeParameter_links MlacpMaximizeParameter = "links"

    // Maximize the operational bandwidth
    MlacpMaximizeParameter_bandwidth MlacpMaximizeParameter = "bandwidth"
)

// BundlePeriod represents Bundle period
type BundlePeriod string

const (
    // Use the standard LACP short period (1s)
    BundlePeriod_true BundlePeriod = "true"
)

// Lacp
// Link Aggregation Control Protocol commands
type Lacp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unique identifier for this system. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    SystemMac interface{}

    // Priority for this system. Lower value is higher priority. The type is
    // interface{} with range: 1..65535. The default value is 32768.
    SystemPriority interface{}
}

func (lacp *Lacp) GetEntityData() *types.CommonEntityData {
    lacp.EntityData.YFilter = lacp.YFilter
    lacp.EntityData.YangName = "lacp"
    lacp.EntityData.BundleName = "cisco_ios_xr"
    lacp.EntityData.ParentYangName = "Cisco-IOS-XR-bundlemgr-cfg"
    lacp.EntityData.SegmentPath = "Cisco-IOS-XR-bundlemgr-cfg:lacp"
    lacp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lacp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lacp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lacp.EntityData.Children = make(map[string]types.YChild)
    lacp.EntityData.Leafs = make(map[string]types.YLeaf)
    lacp.EntityData.Leafs["system-mac"] = types.YLeaf{"SystemMac", lacp.SystemMac}
    lacp.EntityData.Leafs["system-priority"] = types.YLeaf{"SystemPriority", lacp.SystemPriority}
    return &(lacp.EntityData)
}

