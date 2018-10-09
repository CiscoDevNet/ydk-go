// This module contains a collection of YANG
// definitions for VRRP operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package vrrp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package vrrp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-vrrp-oper vrrp-oper-data}", reflect.TypeOf(VrrpOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-vrrp-oper:vrrp-oper-data", reflect.TypeOf(VrrpOperData{}))
}

// MasterReason represents Indicates why this router became master of the VRRP group
type MasterReason string

const (
    // Router is not in master state
    MasterReason_reason_not_master MasterReason = "reason-not-master"

    // Won the Master election due to higher priority
    MasterReason_reason_priority MasterReason = "reason-priority"

    // Preempted a lower priority Master router
    MasterReason_reason_preempt MasterReason = "reason-preempt"

    // Master router stopped responding
    MasterReason_reason_master_no_response MasterReason = "reason-master-no-response"
)

// VrrpProtoState represents VRRP group state
type VrrpProtoState string

const (
    // Indicates that the virtual router group is waiting for a startup event
    VrrpProtoState_proto_state_init VrrpProtoState = "proto-state-init"

    // Indicates that the virtual router is monitoring the availability of a master
    VrrpProtoState_proto_state_backup VrrpProtoState = "proto-state-backup"

    // Indicates that the virtual router is forwarding packets for IP addresses that are associated with this router
    VrrpProtoState_proto_state_master VrrpProtoState = "proto-state-master"

    // Indicates that the virtual router is under recovery
    VrrpProtoState_proto_state_recover VrrpProtoState = "proto-state-recover"
)

// ProtoVersion represents VRRP protocol version
type ProtoVersion string

const (
    ProtoVersion_vrrp_v3 ProtoVersion = "vrrp-v3"
)

// TrackState represents Indicates whether the track is resolved
type TrackState string

const (
    // Track is resolved
    TrackState_vrrp_track_state_resolved TrackState = "vrrp-track-state-resolved"

    // Track is unresolved
    TrackState_vrrp_track_state_unresolved TrackState = "vrrp-track-state-unresolved"
)

// OmpStateUpdown represents Indicates the state of the Overlay Management Protocol tracking
type OmpStateUpdown string

const (
    // Indicates OMP track is up
    OmpStateUpdown_omp_up OmpStateUpdown = "omp-up"

    // Indicates OMP track is down
    OmpStateUpdown_omp_down OmpStateUpdown = "omp-down"
)

// VrrpOperData
// VRRP operational data
type VrrpOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRRP operational state. The type is slice of VrrpOperData_VrrpOperState.
    VrrpOperState []*VrrpOperData_VrrpOperState
}

func (vrrpOperData *VrrpOperData) GetEntityData() *types.CommonEntityData {
    vrrpOperData.EntityData.YFilter = vrrpOperData.YFilter
    vrrpOperData.EntityData.YangName = "vrrp-oper-data"
    vrrpOperData.EntityData.BundleName = "cisco_ios_xe"
    vrrpOperData.EntityData.ParentYangName = "Cisco-IOS-XE-vrrp-oper"
    vrrpOperData.EntityData.SegmentPath = "Cisco-IOS-XE-vrrp-oper:vrrp-oper-data"
    vrrpOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vrrpOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vrrpOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vrrpOperData.EntityData.Children = types.NewOrderedMap()
    vrrpOperData.EntityData.Children.Append("vrrp-oper-state", types.YChild{"VrrpOperState", nil})
    for i := range vrrpOperData.VrrpOperState {
        vrrpOperData.EntityData.Children.Append(types.GetSegmentPath(vrrpOperData.VrrpOperState[i]), types.YChild{"VrrpOperState", vrrpOperData.VrrpOperState[i]})
    }
    vrrpOperData.EntityData.Leafs = types.NewOrderedMap()

    vrrpOperData.EntityData.YListKeys = []string {}

    return &(vrrpOperData.EntityData)
}

// VrrpOperData_VrrpOperState
// VRRP operational state
type VrrpOperData_VrrpOperState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IfIndex for the interface on which VRRP group is
    // hosted. The type is interface{} with range: 0..4294967295.
    IfNumber interface{}

    // This attribute is a key. VRRP group number. The type is interface{} with
    // range: 0..4294967295.
    GroupId interface{}

    // This attribute is a key. Address family of VRRP group. IPv4 or IPv6 are the
    // two valid values. The type is AddrType.
    AddrType interface{}

    // VRRP version. The type is ProtoVersion.
    Version interface{}

    // Primary Virtual IP address for the VRRP group. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    VirtualIp interface{}

    // Name for the interface on which VRRP group is hosted. The type is string.
    IfName interface{}

    // VRRP group state. The type is VrrpProtoState.
    VrrpState interface{}

    // Virtual MAC address for the VRRP group. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMac interface{}

    // IP address of the Master router for the VRRP group. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    MasterIp interface{}

    // Whether the router owns the VRRP Primary virtual IP address. When Interface
    // IP address and VRRP Primary virtual IP address are the same for this
    // router, its priority is bumped up to 255. The type is bool.
    IsOwner interface{}

    // Specifies the priority value used for VRRP master election process. Valid
    // values are 0 to 255. 0 is used by a current master router to gracefully
    // retire from the vrrp group. 255 indicates the master router also owns the
    // VRRP virtual IP address. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Time interval between hello packets sent by the master router in
    // milliseconds. The type is interface{} with range: 0..4294967295.
    AdvertisementTimer interface{}

    // Time after which a backup router declares the current master to be down in
    // milliseconds. The type is interface{} with range: 0..4294967295.
    MasterDownTimer interface{}

    // Time to skew Master Down Interval in milliseconds. The type is interface{}
    // with range: 0..4294967295.
    SkewTime interface{}

    // Controls whether a higher priority virtual router will preempt a lower
    // priority master. The type is bool.
    Preempt interface{}

    // Number of master transitions that have happened in the VRRP group. The type
    // is interface{} with range: 0..4294967295.
    MasterTransitions interface{}

    // Indicates why this router became master of the VRRP group. The type is
    // MasterReason.
    NewMasterReason interface{}

    // Time when state of the VRRP group last changed. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastStateChangeTime interface{}

    // Total number of VRRP packets that arrived with advertisement interval
    // different to the configured value. The type is interface{} with range:
    // 0..4294967295.
    AdvIntervalErrors interface{}

    // Total number of VRRP packets received by the virtual router with IPv4 TTL
    // (for VRRP over IPv4) or IPv6 Hop Limit (for VRRP over IPv6) not equal to
    // 255. The type is interface{} with range: 0..4294967295.
    IpTtlErrors interface{}

    // Total number of VRRP packets received with priority 0. The type is
    // interface{} with range: 0..4294967295.
    RcvdPriZeroPak interface{}

    // Total number of VRRP packets sent with priority 0. The type is interface{}
    // with range: 0..4294967295.
    SentPriZeroPak interface{}

    // Total number of VRRP packets received with invalid Type field. The type is
    // interface{} with range: 0..4294967295.
    RcvdInvalidTypePak interface{}

    // Total number of VRRP packets received with address not matching the address
    // list locally configured on the router. The type is interface{} with range:
    // 0..4294967295.
    AddrListErrors interface{}

    // Total number of VRRP packets received with length less that VRRP header
    // length. The type is interface{} with range: 0..4294967295.
    PakLenErrors interface{}

    // Indicates the last time when a discontinuity happened in gathering
    // statistics. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    DiscontinuityTime interface{}

    // Total number of VRRP packets sent. The type is interface{} with range:
    // 0..4294967295.
    AdvertisementSent interface{}

    // Total number of VRRP packets received. The type is interface{} with range:
    // 0..4294967295.
    AdvertisementRcvd interface{}

    // Indicates the state of the Overlay Management protocol tracking. The type
    // is OmpStateUpdown.
    OmpState interface{}

    // Contains the list of secondary address configured on the group. The type is
    // one of the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SecondaryVipAddresses []interface{}

    // Status of list of tracking objects in the group. The type is slice of
    // VrrpOperData_VrrpOperState_TrackList.
    TrackList []*VrrpOperData_VrrpOperState_TrackList
}

func (vrrpOperState *VrrpOperData_VrrpOperState) GetEntityData() *types.CommonEntityData {
    vrrpOperState.EntityData.YFilter = vrrpOperState.YFilter
    vrrpOperState.EntityData.YangName = "vrrp-oper-state"
    vrrpOperState.EntityData.BundleName = "cisco_ios_xe"
    vrrpOperState.EntityData.ParentYangName = "vrrp-oper-data"
    vrrpOperState.EntityData.SegmentPath = "vrrp-oper-state" + types.AddKeyToken(vrrpOperState.IfNumber, "if-number") + types.AddKeyToken(vrrpOperState.GroupId, "group-id") + types.AddKeyToken(vrrpOperState.AddrType, "addr-type")
    vrrpOperState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vrrpOperState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vrrpOperState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vrrpOperState.EntityData.Children = types.NewOrderedMap()
    vrrpOperState.EntityData.Children.Append("track-list", types.YChild{"TrackList", nil})
    for i := range vrrpOperState.TrackList {
        vrrpOperState.EntityData.Children.Append(types.GetSegmentPath(vrrpOperState.TrackList[i]), types.YChild{"TrackList", vrrpOperState.TrackList[i]})
    }
    vrrpOperState.EntityData.Leafs = types.NewOrderedMap()
    vrrpOperState.EntityData.Leafs.Append("if-number", types.YLeaf{"IfNumber", vrrpOperState.IfNumber})
    vrrpOperState.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", vrrpOperState.GroupId})
    vrrpOperState.EntityData.Leafs.Append("addr-type", types.YLeaf{"AddrType", vrrpOperState.AddrType})
    vrrpOperState.EntityData.Leafs.Append("version", types.YLeaf{"Version", vrrpOperState.Version})
    vrrpOperState.EntityData.Leafs.Append("virtual-ip", types.YLeaf{"VirtualIp", vrrpOperState.VirtualIp})
    vrrpOperState.EntityData.Leafs.Append("if-name", types.YLeaf{"IfName", vrrpOperState.IfName})
    vrrpOperState.EntityData.Leafs.Append("vrrp-state", types.YLeaf{"VrrpState", vrrpOperState.VrrpState})
    vrrpOperState.EntityData.Leafs.Append("virtual-mac", types.YLeaf{"VirtualMac", vrrpOperState.VirtualMac})
    vrrpOperState.EntityData.Leafs.Append("master-ip", types.YLeaf{"MasterIp", vrrpOperState.MasterIp})
    vrrpOperState.EntityData.Leafs.Append("is-owner", types.YLeaf{"IsOwner", vrrpOperState.IsOwner})
    vrrpOperState.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", vrrpOperState.Priority})
    vrrpOperState.EntityData.Leafs.Append("advertisement-timer", types.YLeaf{"AdvertisementTimer", vrrpOperState.AdvertisementTimer})
    vrrpOperState.EntityData.Leafs.Append("master-down-timer", types.YLeaf{"MasterDownTimer", vrrpOperState.MasterDownTimer})
    vrrpOperState.EntityData.Leafs.Append("skew-time", types.YLeaf{"SkewTime", vrrpOperState.SkewTime})
    vrrpOperState.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", vrrpOperState.Preempt})
    vrrpOperState.EntityData.Leafs.Append("master-transitions", types.YLeaf{"MasterTransitions", vrrpOperState.MasterTransitions})
    vrrpOperState.EntityData.Leafs.Append("new-master-reason", types.YLeaf{"NewMasterReason", vrrpOperState.NewMasterReason})
    vrrpOperState.EntityData.Leafs.Append("last-state-change-time", types.YLeaf{"LastStateChangeTime", vrrpOperState.LastStateChangeTime})
    vrrpOperState.EntityData.Leafs.Append("adv-interval-errors", types.YLeaf{"AdvIntervalErrors", vrrpOperState.AdvIntervalErrors})
    vrrpOperState.EntityData.Leafs.Append("ip-ttl-errors", types.YLeaf{"IpTtlErrors", vrrpOperState.IpTtlErrors})
    vrrpOperState.EntityData.Leafs.Append("rcvd-pri-zero-pak", types.YLeaf{"RcvdPriZeroPak", vrrpOperState.RcvdPriZeroPak})
    vrrpOperState.EntityData.Leafs.Append("sent-pri-zero-pak", types.YLeaf{"SentPriZeroPak", vrrpOperState.SentPriZeroPak})
    vrrpOperState.EntityData.Leafs.Append("rcvd-invalid-type-pak", types.YLeaf{"RcvdInvalidTypePak", vrrpOperState.RcvdInvalidTypePak})
    vrrpOperState.EntityData.Leafs.Append("addr-list-errors", types.YLeaf{"AddrListErrors", vrrpOperState.AddrListErrors})
    vrrpOperState.EntityData.Leafs.Append("pak-len-errors", types.YLeaf{"PakLenErrors", vrrpOperState.PakLenErrors})
    vrrpOperState.EntityData.Leafs.Append("discontinuity-time", types.YLeaf{"DiscontinuityTime", vrrpOperState.DiscontinuityTime})
    vrrpOperState.EntityData.Leafs.Append("advertisement-sent", types.YLeaf{"AdvertisementSent", vrrpOperState.AdvertisementSent})
    vrrpOperState.EntityData.Leafs.Append("advertisement-rcvd", types.YLeaf{"AdvertisementRcvd", vrrpOperState.AdvertisementRcvd})
    vrrpOperState.EntityData.Leafs.Append("omp-state", types.YLeaf{"OmpState", vrrpOperState.OmpState})
    vrrpOperState.EntityData.Leafs.Append("secondary-vip-addresses", types.YLeaf{"SecondaryVipAddresses", vrrpOperState.SecondaryVipAddresses})

    vrrpOperState.EntityData.YListKeys = []string {"IfNumber", "GroupId", "AddrType"}

    return &(vrrpOperState.EntityData)
}

// VrrpOperData_VrrpOperState_TrackList
// Status of list of tracking objects in the group
type VrrpOperData_VrrpOperState_TrackList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the tracking object. The type is string.
    TrackName interface{}

    // State of the tracking object. The type is TrackState.
    TrackObjState interface{}
}

func (trackList *VrrpOperData_VrrpOperState_TrackList) GetEntityData() *types.CommonEntityData {
    trackList.EntityData.YFilter = trackList.YFilter
    trackList.EntityData.YangName = "track-list"
    trackList.EntityData.BundleName = "cisco_ios_xe"
    trackList.EntityData.ParentYangName = "vrrp-oper-state"
    trackList.EntityData.SegmentPath = "track-list"
    trackList.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    trackList.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    trackList.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    trackList.EntityData.Children = types.NewOrderedMap()
    trackList.EntityData.Leafs = types.NewOrderedMap()
    trackList.EntityData.Leafs.Append("track-name", types.YLeaf{"TrackName", trackList.TrackName})
    trackList.EntityData.Leafs.Append("track-obj-state", types.YLeaf{"TrackObjState", trackList.TrackObjState})

    trackList.EntityData.YListKeys = []string {}

    return &(trackList.EntityData)
}

