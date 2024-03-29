// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-ntp package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   ntp: NTP admin operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_ntp_admin_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_ntp_admin_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-ntp-admin-oper ntp}", reflect.TypeOf(Ntp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-ntp-admin-oper:ntp", reflect.TypeOf(Ntp{}))
}

// NtpPeerStatus represents Type of peer status
type NtpPeerStatus string

const (
    //    reject
    NtpPeerStatus_ntp_ctl_pst_sel_reject NtpPeerStatus = "ntp-ctl-pst-sel-reject"

    //  x falsetick
    NtpPeerStatus_ntp_ctl_pst_sel_sane NtpPeerStatus = "ntp-ctl-pst-sel-sane"

    //  . excess 
    NtpPeerStatus_ntp_ctl_pst_sel_correct NtpPeerStatus = "ntp-ctl-pst-sel-correct"

    //  - outlyer
    NtpPeerStatus_ntp_ctl_pst_sel_selcand NtpPeerStatus = "ntp-ctl-pst-sel-selcand"

    //  + candidate
    NtpPeerStatus_ntp_ctl_pst_sel_sync_cand NtpPeerStatus = "ntp-ctl-pst-sel-sync-cand"

    //  # selected
    NtpPeerStatus_ntp_ctl_pst_sel_distsys_peer NtpPeerStatus = "ntp-ctl-pst-sel-distsys-peer"

    //  * sys peer
    NtpPeerStatus_ntp_ctl_pst_sel_sys_peer NtpPeerStatus = "ntp-ctl-pst-sel-sys-peer"

    //  o pps peer
    NtpPeerStatus_ntp_ctl_pst_sel_pps NtpPeerStatus = "ntp-ctl-pst-sel-pps"
)

// NtpMode represents Type of mode
type NtpMode string

const (
    // Unspecified probably old NTP version
    NtpMode_ntp_mode_unspec NtpMode = "ntp-mode-unspec"

    // Symmetric active
    NtpMode_ntp_mode_symetric_active NtpMode = "ntp-mode-symetric-active"

    // Symmetric passive
    NtpMode_ntp_mode_symetric_passive NtpMode = "ntp-mode-symetric-passive"

    // Client mode
    NtpMode_ntp_mode_client NtpMode = "ntp-mode-client"

    // Server mode
    NtpMode_ntp_mode_server NtpMode = "ntp-mode-server"

    // Broadcast mode
    NtpMode_ntp_mode_xcast_server NtpMode = "ntp-mode-xcast-server"

    // Control mode packet
    NtpMode_ntp_mode_control NtpMode = "ntp-mode-control"

    // Implementation defined function
    NtpMode_ntp_mode_private NtpMode = "ntp-mode-private"

    // A broadcast client mode
    NtpMode_ntp_mode_xcast_client NtpMode = "ntp-mode-xcast-client"
)

// ClockUpdateNode represents Mode of Clock Update
type ClockUpdateNode string

const (
    //  clock is never updated
    ClockUpdateNode_clk_never_updated ClockUpdateNode = "clk-never-updated"

    //  clock is updated
    ClockUpdateNode_clk_updated ClockUpdateNode = "clk-updated"

    //  clock has no update info
    ClockUpdateNode_clk_no_update_info ClockUpdateNode = "clk-no-update-info"
)

// NtpLoopFilterState represents Loop filter state
type NtpLoopFilterState string

const (
    //  never set
    NtpLoopFilterState_ntp_loop_flt_n_set NtpLoopFilterState = "ntp-loop-flt-n-set"

    //  drift set from file
    NtpLoopFilterState_ntp_loop_flt_f_set NtpLoopFilterState = "ntp-loop-flt-f-set"

    //  spike
    NtpLoopFilterState_ntp_loop_flt_spik NtpLoopFilterState = "ntp-loop-flt-spik"

    //  drift being measured
    NtpLoopFilterState_ntp_loop_flt_freq NtpLoopFilterState = "ntp-loop-flt-freq"

    //  normal controlled loop
    NtpLoopFilterState_ntp_loop_flt_sync NtpLoopFilterState = "ntp-loop-flt-sync"

    //  unknown
    NtpLoopFilterState_ntp_loop_flt_unkn NtpLoopFilterState = "ntp-loop-flt-unkn"
)

// NtpLeap represents Type of leap
type NtpLeap string

const (
    // Normal, no leap second warning
    NtpLeap_ntp_leap_no_warning NtpLeap = "ntp-leap-no-warning"

    // Last minute of day has 61 seconds
    NtpLeap_ntp_leap_addse_cond NtpLeap = "ntp-leap-addse-cond"

    // Last minute of day has 59 seconds
    NtpLeap_ntp_leap_delse_cond NtpLeap = "ntp-leap-delse-cond"

    // Overload, clock is free running
    NtpLeap_ntp_leap_not_in_sync NtpLeap = "ntp-leap-not-in-sync"
)

// Ntp
// NTP admin operational data
type Ntp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rack-specific NTP operational data.
    Racks Ntp_Racks
}

func (ntp *Ntp) GetEntityData() *types.CommonEntityData {
    ntp.EntityData.YFilter = ntp.YFilter
    ntp.EntityData.YangName = "ntp"
    ntp.EntityData.BundleName = "cisco_ios_xr"
    ntp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-ntp-admin-oper"
    ntp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp"
    ntp.EntityData.AbsolutePath = ntp.EntityData.SegmentPath
    ntp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntp.EntityData.Children = types.NewOrderedMap()
    ntp.EntityData.Children.Append("racks", types.YChild{"Racks", &ntp.Racks})
    ntp.EntityData.Leafs = types.NewOrderedMap()

    ntp.EntityData.YListKeys = []string {}

    return &(ntp.EntityData)
}

// Ntp_Racks
// Rack-specific NTP operational data
type Ntp_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular rack. The type is slice of
    // Ntp_Racks_Rack.
    Rack []*Ntp_Racks_Rack
}

func (racks *Ntp_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "ntp"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/" + racks.EntityData.SegmentPath
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = types.NewOrderedMap()
    racks.EntityData.Children.Append("rack", types.YChild{"Rack", nil})
    for i := range racks.Rack {
        racks.EntityData.Children.Append(types.GetSegmentPath(racks.Rack[i]), types.YChild{"Rack", racks.Rack[i]})
    }
    racks.EntityData.Leafs = types.NewOrderedMap()

    racks.EntityData.YListKeys = []string {}

    return &(racks.EntityData)
}

// Ntp_Racks_Rack
// NTP operational data for a particular rack
type Ntp_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The rack number. The type is interface{} with
    // range: 0..4294967295.
    Number interface{}

    // Node-specific NTP operational data.
    Slots Ntp_Racks_Rack_Slots
}

func (rack *Ntp_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + types.AddKeyToken(rack.Number, "number")
    rack.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/" + rack.EntityData.SegmentPath
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = types.NewOrderedMap()
    rack.EntityData.Children.Append("slots", types.YChild{"Slots", &rack.Slots})
    rack.EntityData.Leafs = types.NewOrderedMap()
    rack.EntityData.Leafs.Append("number", types.YLeaf{"Number", rack.Number})

    rack.EntityData.YListKeys = []string {"Number"}

    return &(rack.EntityData)
}

// Ntp_Racks_Rack_Slots
// Node-specific NTP operational data
type Ntp_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular slot. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot.
    Slot []*Ntp_Racks_Rack_Slots_Slot
}

func (slots *Ntp_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot
// NTP operational data for a particular slot
type Ntp_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The slot number. The type is interface{} with
    // range: 0..4294967295.
    Number interface{}

    // Instance-specific NTP operational data.
    Instances Ntp_Racks_Rack_Slots_Slot_Instances
}

func (slot *Ntp_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Number, "number")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("instances", types.YChild{"Instances", &slot.Instances})
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("number", types.YLeaf{"Number", slot.Number})

    slot.EntityData.YListKeys = []string {"Number"}

    return &(slot.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances
// Instance-specific NTP operational data
type Ntp_Racks_Rack_Slots_Slot_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular instance. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []*Ntp_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Ntp_Racks_Rack_Slots_Slot_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "slot"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/" + instances.EntityData.SegmentPath
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = types.NewOrderedMap()
    instances.EntityData.Children.Append("instance", types.YChild{"Instance", nil})
    for i := range instances.Instance {
        instances.EntityData.Children.Append(types.GetSegmentPath(instances.Instance[i]), types.YChild{"Instance", instances.Instance[i]})
    }
    instances.EntityData.Leafs = types.NewOrderedMap()

    instances.EntityData.YListKeys = []string {}

    return &(instances.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance
// NTP operational data for a particular
// instance
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The instance number. The type is interface{} with
    // range: 0..4294967295.
    Number interface{}

    // Status of NTP peer(s).
    Status Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status

    // NTP Associations information.
    Associations Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations

    // NTP Associations Detail information.
    AssociationsDetail Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail
}

func (instance *Ntp_Racks_Rack_Slots_Slot_Instances_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "instances"
    instance.EntityData.SegmentPath = "instance" + types.AddKeyToken(instance.Number, "number")
    instance.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/" + instance.EntityData.SegmentPath
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Children.Append("status", types.YChild{"Status", &instance.Status})
    instance.EntityData.Children.Append("associations", types.YChild{"Associations", &instance.Associations})
    instance.EntityData.Children.Append("associations-detail", types.YChild{"AssociationsDetail", &instance.AssociationsDetail})
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("number", types.YLeaf{"Number", instance.Number})

    instance.EntityData.YListKeys = []string {"Number"}

    return &(instance.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status
// Status of NTP peer(s)
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Peer dispersion. The type is string.
    SysDispersion interface{}

    // Clock offset. The type is string.
    SysOffset interface{}

    // Clock period in nanosecs. The type is interface{} with range:
    // 0..4294967295. Units are nanosecond.
    ClockPeriod interface{}

    // leap. The type is NtpLeap.
    SysLeap interface{}

    // Precision. The type is interface{} with range: -128..127.
    SysPrecision interface{}

    // Stratum. The type is interface{} with range: 0..255.
    SysStratum interface{}

    // Reference clock ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SysRefId interface{}

    // Root delay. The type is string.
    SysRootDelay interface{}

    // Root dispersion. The type is string.
    SysRootDispersion interface{}

    // Loop Filter State. The type is NtpLoopFilterState.
    LoopFilterState interface{}

    // Peer poll interval. The type is interface{} with range: 0..255.
    PollInterval interface{}

    // Is clock updated. The type is ClockUpdateNode.
    IsUpdated interface{}

    // Last Update. The type is interface{} with range: -2147483648..2147483647.
    LastUpdate interface{}

    // Reference time.
    SysRefTime Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime

    // System Drift.
    SysDrift Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift
}

func (status *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "instance"
    status.EntityData.SegmentPath = "status"
    status.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/" + status.EntityData.SegmentPath
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = types.NewOrderedMap()
    status.EntityData.Children.Append("sys-ref-time", types.YChild{"SysRefTime", &status.SysRefTime})
    status.EntityData.Children.Append("sys-drift", types.YChild{"SysDrift", &status.SysDrift})
    status.EntityData.Leafs = types.NewOrderedMap()
    status.EntityData.Leafs.Append("is-ntp-enabled", types.YLeaf{"IsNtpEnabled", status.IsNtpEnabled})
    status.EntityData.Leafs.Append("sys-dispersion", types.YLeaf{"SysDispersion", status.SysDispersion})
    status.EntityData.Leafs.Append("sys-offset", types.YLeaf{"SysOffset", status.SysOffset})
    status.EntityData.Leafs.Append("clock-period", types.YLeaf{"ClockPeriod", status.ClockPeriod})
    status.EntityData.Leafs.Append("sys-leap", types.YLeaf{"SysLeap", status.SysLeap})
    status.EntityData.Leafs.Append("sys-precision", types.YLeaf{"SysPrecision", status.SysPrecision})
    status.EntityData.Leafs.Append("sys-stratum", types.YLeaf{"SysStratum", status.SysStratum})
    status.EntityData.Leafs.Append("sys-ref-id", types.YLeaf{"SysRefId", status.SysRefId})
    status.EntityData.Leafs.Append("sys-root-delay", types.YLeaf{"SysRootDelay", status.SysRootDelay})
    status.EntityData.Leafs.Append("sys-root-dispersion", types.YLeaf{"SysRootDispersion", status.SysRootDispersion})
    status.EntityData.Leafs.Append("loop-filter-state", types.YLeaf{"LoopFilterState", status.LoopFilterState})
    status.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", status.PollInterval})
    status.EntityData.Leafs.Append("is-updated", types.YLeaf{"IsUpdated", status.IsUpdated})
    status.EntityData.Leafs.Append("last-update", types.YLeaf{"LastUpdate", status.LastUpdate})

    status.EntityData.YListKeys = []string {}

    return &(status.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime
// Reference time
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_FracSecs
}

func (sysRefTime *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime) GetEntityData() *types.CommonEntityData {
    sysRefTime.EntityData.YFilter = sysRefTime.YFilter
    sysRefTime.EntityData.YangName = "sys-ref-time"
    sysRefTime.EntityData.BundleName = "cisco_ios_xr"
    sysRefTime.EntityData.ParentYangName = "status"
    sysRefTime.EntityData.SegmentPath = "sys-ref-time"
    sysRefTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/" + sysRefTime.EntityData.SegmentPath
    sysRefTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysRefTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysRefTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysRefTime.EntityData.Children = types.NewOrderedMap()
    sysRefTime.EntityData.Children.Append("sec", types.YChild{"Sec", &sysRefTime.Sec})
    sysRefTime.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &sysRefTime.FracSecs})
    sysRefTime.EntityData.Leafs = types.NewOrderedMap()

    sysRefTime.EntityData.YListKeys = []string {}

    return &(sysRefTime.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "sys-ref-time"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/sys-ref-time/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysRefTime_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "sys-ref-time"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/sys-ref-time/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift
// System Drift
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_FracSecs
}

func (sysDrift *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift) GetEntityData() *types.CommonEntityData {
    sysDrift.EntityData.YFilter = sysDrift.YFilter
    sysDrift.EntityData.YangName = "sys-drift"
    sysDrift.EntityData.BundleName = "cisco_ios_xr"
    sysDrift.EntityData.ParentYangName = "status"
    sysDrift.EntityData.SegmentPath = "sys-drift"
    sysDrift.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/" + sysDrift.EntityData.SegmentPath
    sysDrift.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysDrift.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysDrift.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysDrift.EntityData.Children = types.NewOrderedMap()
    sysDrift.EntityData.Children.Append("sec", types.YChild{"Sec", &sysDrift.Sec})
    sysDrift.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &sysDrift.FracSecs})
    sysDrift.EntityData.Leafs = types.NewOrderedMap()

    sysDrift.EntityData.YListKeys = []string {}

    return &(sysDrift.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "sys-drift"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/sys-drift/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Status_SysDrift_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "sys-drift"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/status/sys-drift/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations
// NTP Associations information
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Leap. The type is NtpLeap.
    SysLeap interface{}

    // Peer info. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo.
    PeerSummaryInfo []*Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo
}

func (associations *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations) GetEntityData() *types.CommonEntityData {
    associations.EntityData.YFilter = associations.YFilter
    associations.EntityData.YangName = "associations"
    associations.EntityData.BundleName = "cisco_ios_xr"
    associations.EntityData.ParentYangName = "instance"
    associations.EntityData.SegmentPath = "associations"
    associations.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/" + associations.EntityData.SegmentPath
    associations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associations.EntityData.Children = types.NewOrderedMap()
    associations.EntityData.Children.Append("peer-summary-info", types.YChild{"PeerSummaryInfo", nil})
    for i := range associations.PeerSummaryInfo {
        types.SetYListKey(associations.PeerSummaryInfo[i], i)
        associations.EntityData.Children.Append(types.GetSegmentPath(associations.PeerSummaryInfo[i]), types.YChild{"PeerSummaryInfo", associations.PeerSummaryInfo[i]})
    }
    associations.EntityData.Leafs = types.NewOrderedMap()
    associations.EntityData.Leafs.Append("is-ntp-enabled", types.YLeaf{"IsNtpEnabled", associations.IsNtpEnabled})
    associations.EntityData.Leafs.Append("sys-leap", types.YLeaf{"SysLeap", associations.SysLeap})

    associations.EntityData.YListKeys = []string {}

    return &(associations.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo
// Peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Time since last frame received (-1=none). The type is interface{} with
    // range: -2147483648..2147483647.
    TimeSince interface{}

    // Common peer info.
    PeerInfoCommon Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo_PeerInfoCommon
}

func (peerSummaryInfo *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo) GetEntityData() *types.CommonEntityData {
    peerSummaryInfo.EntityData.YFilter = peerSummaryInfo.YFilter
    peerSummaryInfo.EntityData.YangName = "peer-summary-info"
    peerSummaryInfo.EntityData.BundleName = "cisco_ios_xr"
    peerSummaryInfo.EntityData.ParentYangName = "associations"
    peerSummaryInfo.EntityData.SegmentPath = "peer-summary-info" + types.AddNoKeyToken(peerSummaryInfo)
    peerSummaryInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations/" + peerSummaryInfo.EntityData.SegmentPath
    peerSummaryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerSummaryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerSummaryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerSummaryInfo.EntityData.Children = types.NewOrderedMap()
    peerSummaryInfo.EntityData.Children.Append("peer-info-common", types.YChild{"PeerInfoCommon", &peerSummaryInfo.PeerInfoCommon})
    peerSummaryInfo.EntityData.Leafs = types.NewOrderedMap()
    peerSummaryInfo.EntityData.Leafs.Append("time-since", types.YLeaf{"TimeSince", peerSummaryInfo.TimeSince})

    peerSummaryInfo.EntityData.YListKeys = []string {}

    return &(peerSummaryInfo.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo_PeerInfoCommon
// Common peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo_PeerInfoCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association mode with this peer. The type is NtpMode.
    HostMode interface{}

    // Is configured. The type is bool.
    IsConfigured interface{}

    // Peer Address. The type is string.
    Address interface{}

    // Peer reference ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferenceId interface{}

    // Host poll. The type is interface{} with range: 0..255.
    HostPoll interface{}

    // Reachability. The type is interface{} with range: 0..255.
    Reachability interface{}

    // Peer stratum. The type is interface{} with range: 0..255.
    Stratum interface{}

    // Peer status. The type is NtpPeerStatus.
    Status interface{}

    // Peer delay. The type is string.
    Delay interface{}

    // Peer offset. The type is string.
    Offset interface{}

    // Peer dispersion. The type is string.
    Dispersion interface{}

    // Indicates whether this is syspeer. The type is bool.
    IsSysPeer interface{}
}

func (peerInfoCommon *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo_PeerInfoCommon) GetEntityData() *types.CommonEntityData {
    peerInfoCommon.EntityData.YFilter = peerInfoCommon.YFilter
    peerInfoCommon.EntityData.YangName = "peer-info-common"
    peerInfoCommon.EntityData.BundleName = "cisco_ios_xr"
    peerInfoCommon.EntityData.ParentYangName = "peer-summary-info"
    peerInfoCommon.EntityData.SegmentPath = "peer-info-common"
    peerInfoCommon.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations/peer-summary-info/" + peerInfoCommon.EntityData.SegmentPath
    peerInfoCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfoCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfoCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfoCommon.EntityData.Children = types.NewOrderedMap()
    peerInfoCommon.EntityData.Leafs = types.NewOrderedMap()
    peerInfoCommon.EntityData.Leafs.Append("host-mode", types.YLeaf{"HostMode", peerInfoCommon.HostMode})
    peerInfoCommon.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", peerInfoCommon.IsConfigured})
    peerInfoCommon.EntityData.Leafs.Append("address", types.YLeaf{"Address", peerInfoCommon.Address})
    peerInfoCommon.EntityData.Leafs.Append("reference-id", types.YLeaf{"ReferenceId", peerInfoCommon.ReferenceId})
    peerInfoCommon.EntityData.Leafs.Append("host-poll", types.YLeaf{"HostPoll", peerInfoCommon.HostPoll})
    peerInfoCommon.EntityData.Leafs.Append("reachability", types.YLeaf{"Reachability", peerInfoCommon.Reachability})
    peerInfoCommon.EntityData.Leafs.Append("stratum", types.YLeaf{"Stratum", peerInfoCommon.Stratum})
    peerInfoCommon.EntityData.Leafs.Append("status", types.YLeaf{"Status", peerInfoCommon.Status})
    peerInfoCommon.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", peerInfoCommon.Delay})
    peerInfoCommon.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", peerInfoCommon.Offset})
    peerInfoCommon.EntityData.Leafs.Append("dispersion", types.YLeaf{"Dispersion", peerInfoCommon.Dispersion})
    peerInfoCommon.EntityData.Leafs.Append("is-sys-peer", types.YLeaf{"IsSysPeer", peerInfoCommon.IsSysPeer})

    peerInfoCommon.EntityData.YListKeys = []string {}

    return &(peerInfoCommon.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail
// NTP Associations Detail information
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Leap. The type is NtpLeap.
    SysLeap interface{}

    // Peer info. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo.
    PeerDetailInfo []*Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo
}

func (associationsDetail *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail) GetEntityData() *types.CommonEntityData {
    associationsDetail.EntityData.YFilter = associationsDetail.YFilter
    associationsDetail.EntityData.YangName = "associations-detail"
    associationsDetail.EntityData.BundleName = "cisco_ios_xr"
    associationsDetail.EntityData.ParentYangName = "instance"
    associationsDetail.EntityData.SegmentPath = "associations-detail"
    associationsDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/" + associationsDetail.EntityData.SegmentPath
    associationsDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationsDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationsDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationsDetail.EntityData.Children = types.NewOrderedMap()
    associationsDetail.EntityData.Children.Append("peer-detail-info", types.YChild{"PeerDetailInfo", nil})
    for i := range associationsDetail.PeerDetailInfo {
        types.SetYListKey(associationsDetail.PeerDetailInfo[i], i)
        associationsDetail.EntityData.Children.Append(types.GetSegmentPath(associationsDetail.PeerDetailInfo[i]), types.YChild{"PeerDetailInfo", associationsDetail.PeerDetailInfo[i]})
    }
    associationsDetail.EntityData.Leafs = types.NewOrderedMap()
    associationsDetail.EntityData.Leafs.Append("is-ntp-enabled", types.YLeaf{"IsNtpEnabled", associationsDetail.IsNtpEnabled})
    associationsDetail.EntityData.Leafs.Append("sys-leap", types.YLeaf{"SysLeap", associationsDetail.SysLeap})

    associationsDetail.EntityData.YListKeys = []string {}

    return &(associationsDetail.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo
// Peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Leap. The type is NtpLeap.
    Leap interface{}

    // Peer's association mode. The type is NtpMode.
    PeerMode interface{}

    // Peer poll interval. The type is interface{} with range: 0..255.
    PollInterval interface{}

    // Is refclock. The type is bool.
    IsRefClock interface{}

    // Is authenticated. The type is bool.
    IsAuthenticated interface{}

    // Root delay. The type is string.
    RootDelay interface{}

    // Root dispersion. The type is string.
    RootDispersion interface{}

    // Synch distance. The type is string.
    SynchDistance interface{}

    // Precision. The type is interface{} with range: -128..127.
    Precision interface{}

    // NTP version. The type is interface{} with range: 0..255.
    Version interface{}

    // Index into filter shift register. The type is interface{} with range:
    // 0..4294967295.
    FilterIndex interface{}

    // Cast Flags. The type is interface{} with range: 0..255.
    CastFlags interface{}

    // Common peer info.
    PeerInfoCommon Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_PeerInfoCommon

    // Reference time.
    RefTime Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime

    // Originate timestamp.
    OriginateTime Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime

    // Receive timestamp.
    ReceiveTime Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime

    // Transmit timestamp.
    TransmitTime Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime

    // Filter Details. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail.
    FilterDetail []*Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail
}

func (peerDetailInfo *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "associations-detail"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info" + types.AddNoKeyToken(peerDetailInfo)
    peerDetailInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/" + peerDetailInfo.EntityData.SegmentPath
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = types.NewOrderedMap()
    peerDetailInfo.EntityData.Children.Append("peer-info-common", types.YChild{"PeerInfoCommon", &peerDetailInfo.PeerInfoCommon})
    peerDetailInfo.EntityData.Children.Append("ref-time", types.YChild{"RefTime", &peerDetailInfo.RefTime})
    peerDetailInfo.EntityData.Children.Append("originate-time", types.YChild{"OriginateTime", &peerDetailInfo.OriginateTime})
    peerDetailInfo.EntityData.Children.Append("receive-time", types.YChild{"ReceiveTime", &peerDetailInfo.ReceiveTime})
    peerDetailInfo.EntityData.Children.Append("transmit-time", types.YChild{"TransmitTime", &peerDetailInfo.TransmitTime})
    peerDetailInfo.EntityData.Children.Append("filter-detail", types.YChild{"FilterDetail", nil})
    for i := range peerDetailInfo.FilterDetail {
        types.SetYListKey(peerDetailInfo.FilterDetail[i], i)
        peerDetailInfo.EntityData.Children.Append(types.GetSegmentPath(peerDetailInfo.FilterDetail[i]), types.YChild{"FilterDetail", peerDetailInfo.FilterDetail[i]})
    }
    peerDetailInfo.EntityData.Leafs = types.NewOrderedMap()
    peerDetailInfo.EntityData.Leafs.Append("leap", types.YLeaf{"Leap", peerDetailInfo.Leap})
    peerDetailInfo.EntityData.Leafs.Append("peer-mode", types.YLeaf{"PeerMode", peerDetailInfo.PeerMode})
    peerDetailInfo.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", peerDetailInfo.PollInterval})
    peerDetailInfo.EntityData.Leafs.Append("is-ref-clock", types.YLeaf{"IsRefClock", peerDetailInfo.IsRefClock})
    peerDetailInfo.EntityData.Leafs.Append("is-authenticated", types.YLeaf{"IsAuthenticated", peerDetailInfo.IsAuthenticated})
    peerDetailInfo.EntityData.Leafs.Append("root-delay", types.YLeaf{"RootDelay", peerDetailInfo.RootDelay})
    peerDetailInfo.EntityData.Leafs.Append("root-dispersion", types.YLeaf{"RootDispersion", peerDetailInfo.RootDispersion})
    peerDetailInfo.EntityData.Leafs.Append("synch-distance", types.YLeaf{"SynchDistance", peerDetailInfo.SynchDistance})
    peerDetailInfo.EntityData.Leafs.Append("precision", types.YLeaf{"Precision", peerDetailInfo.Precision})
    peerDetailInfo.EntityData.Leafs.Append("version", types.YLeaf{"Version", peerDetailInfo.Version})
    peerDetailInfo.EntityData.Leafs.Append("filter-index", types.YLeaf{"FilterIndex", peerDetailInfo.FilterIndex})
    peerDetailInfo.EntityData.Leafs.Append("cast-flags", types.YLeaf{"CastFlags", peerDetailInfo.CastFlags})

    peerDetailInfo.EntityData.YListKeys = []string {}

    return &(peerDetailInfo.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_PeerInfoCommon
// Common peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_PeerInfoCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association mode with this peer. The type is NtpMode.
    HostMode interface{}

    // Is configured. The type is bool.
    IsConfigured interface{}

    // Peer Address. The type is string.
    Address interface{}

    // Peer reference ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferenceId interface{}

    // Host poll. The type is interface{} with range: 0..255.
    HostPoll interface{}

    // Reachability. The type is interface{} with range: 0..255.
    Reachability interface{}

    // Peer stratum. The type is interface{} with range: 0..255.
    Stratum interface{}

    // Peer status. The type is NtpPeerStatus.
    Status interface{}

    // Peer delay. The type is string.
    Delay interface{}

    // Peer offset. The type is string.
    Offset interface{}

    // Peer dispersion. The type is string.
    Dispersion interface{}

    // Indicates whether this is syspeer. The type is bool.
    IsSysPeer interface{}
}

func (peerInfoCommon *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetEntityData() *types.CommonEntityData {
    peerInfoCommon.EntityData.YFilter = peerInfoCommon.YFilter
    peerInfoCommon.EntityData.YangName = "peer-info-common"
    peerInfoCommon.EntityData.BundleName = "cisco_ios_xr"
    peerInfoCommon.EntityData.ParentYangName = "peer-detail-info"
    peerInfoCommon.EntityData.SegmentPath = "peer-info-common"
    peerInfoCommon.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + peerInfoCommon.EntityData.SegmentPath
    peerInfoCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfoCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfoCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfoCommon.EntityData.Children = types.NewOrderedMap()
    peerInfoCommon.EntityData.Leafs = types.NewOrderedMap()
    peerInfoCommon.EntityData.Leafs.Append("host-mode", types.YLeaf{"HostMode", peerInfoCommon.HostMode})
    peerInfoCommon.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", peerInfoCommon.IsConfigured})
    peerInfoCommon.EntityData.Leafs.Append("address", types.YLeaf{"Address", peerInfoCommon.Address})
    peerInfoCommon.EntityData.Leafs.Append("reference-id", types.YLeaf{"ReferenceId", peerInfoCommon.ReferenceId})
    peerInfoCommon.EntityData.Leafs.Append("host-poll", types.YLeaf{"HostPoll", peerInfoCommon.HostPoll})
    peerInfoCommon.EntityData.Leafs.Append("reachability", types.YLeaf{"Reachability", peerInfoCommon.Reachability})
    peerInfoCommon.EntityData.Leafs.Append("stratum", types.YLeaf{"Stratum", peerInfoCommon.Stratum})
    peerInfoCommon.EntityData.Leafs.Append("status", types.YLeaf{"Status", peerInfoCommon.Status})
    peerInfoCommon.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", peerInfoCommon.Delay})
    peerInfoCommon.EntityData.Leafs.Append("offset", types.YLeaf{"Offset", peerInfoCommon.Offset})
    peerInfoCommon.EntityData.Leafs.Append("dispersion", types.YLeaf{"Dispersion", peerInfoCommon.Dispersion})
    peerInfoCommon.EntityData.Leafs.Append("is-sys-peer", types.YLeaf{"IsSysPeer", peerInfoCommon.IsSysPeer})

    peerInfoCommon.EntityData.YListKeys = []string {}

    return &(peerInfoCommon.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime
// Reference time
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs
}

func (refTime *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime) GetEntityData() *types.CommonEntityData {
    refTime.EntityData.YFilter = refTime.YFilter
    refTime.EntityData.YangName = "ref-time"
    refTime.EntityData.BundleName = "cisco_ios_xr"
    refTime.EntityData.ParentYangName = "peer-detail-info"
    refTime.EntityData.SegmentPath = "ref-time"
    refTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + refTime.EntityData.SegmentPath
    refTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    refTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    refTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    refTime.EntityData.Children = types.NewOrderedMap()
    refTime.EntityData.Children.Append("sec", types.YChild{"Sec", &refTime.Sec})
    refTime.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &refTime.FracSecs})
    refTime.EntityData.Leafs = types.NewOrderedMap()

    refTime.EntityData.YListKeys = []string {}

    return &(refTime.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "ref-time"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/ref-time/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "ref-time"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/ref-time/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime
// Originate timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs
}

func (originateTime *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime) GetEntityData() *types.CommonEntityData {
    originateTime.EntityData.YFilter = originateTime.YFilter
    originateTime.EntityData.YangName = "originate-time"
    originateTime.EntityData.BundleName = "cisco_ios_xr"
    originateTime.EntityData.ParentYangName = "peer-detail-info"
    originateTime.EntityData.SegmentPath = "originate-time"
    originateTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + originateTime.EntityData.SegmentPath
    originateTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    originateTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    originateTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    originateTime.EntityData.Children = types.NewOrderedMap()
    originateTime.EntityData.Children.Append("sec", types.YChild{"Sec", &originateTime.Sec})
    originateTime.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &originateTime.FracSecs})
    originateTime.EntityData.Leafs = types.NewOrderedMap()

    originateTime.EntityData.YListKeys = []string {}

    return &(originateTime.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "originate-time"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/originate-time/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "originate-time"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/originate-time/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime
// Receive timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs
}

func (receiveTime *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetEntityData() *types.CommonEntityData {
    receiveTime.EntityData.YFilter = receiveTime.YFilter
    receiveTime.EntityData.YangName = "receive-time"
    receiveTime.EntityData.BundleName = "cisco_ios_xr"
    receiveTime.EntityData.ParentYangName = "peer-detail-info"
    receiveTime.EntityData.SegmentPath = "receive-time"
    receiveTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + receiveTime.EntityData.SegmentPath
    receiveTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveTime.EntityData.Children = types.NewOrderedMap()
    receiveTime.EntityData.Children.Append("sec", types.YChild{"Sec", &receiveTime.Sec})
    receiveTime.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &receiveTime.FracSecs})
    receiveTime.EntityData.Leafs = types.NewOrderedMap()

    receiveTime.EntityData.YListKeys = []string {}

    return &(receiveTime.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "receive-time"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/receive-time/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "receive-time"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/receive-time/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime
// Transmit timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs
}

func (transmitTime *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime) GetEntityData() *types.CommonEntityData {
    transmitTime.EntityData.YFilter = transmitTime.YFilter
    transmitTime.EntityData.YangName = "transmit-time"
    transmitTime.EntityData.BundleName = "cisco_ios_xr"
    transmitTime.EntityData.ParentYangName = "peer-detail-info"
    transmitTime.EntityData.SegmentPath = "transmit-time"
    transmitTime.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + transmitTime.EntityData.SegmentPath
    transmitTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitTime.EntityData.Children = types.NewOrderedMap()
    transmitTime.EntityData.Children.Append("sec", types.YChild{"Sec", &transmitTime.Sec})
    transmitTime.EntityData.Children.Append("frac-secs", types.YChild{"FracSecs", &transmitTime.FracSecs})
    transmitTime.EntityData.Leafs = types.NewOrderedMap()

    transmitTime.EntityData.YListKeys = []string {}

    return &(transmitTime.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetEntityData() *types.CommonEntityData {
    sec.EntityData.YFilter = sec.YFilter
    sec.EntityData.YangName = "sec"
    sec.EntityData.BundleName = "cisco_ios_xr"
    sec.EntityData.ParentYangName = "transmit-time"
    sec.EntityData.SegmentPath = "sec"
    sec.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/transmit-time/" + sec.EntityData.SegmentPath
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = types.NewOrderedMap()
    sec.EntityData.Leafs = types.NewOrderedMap()
    sec.EntityData.Leafs.Append("int", types.YLeaf{"Int", sec.Int})

    sec.EntityData.YListKeys = []string {}

    return &(sec.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetEntityData() *types.CommonEntityData {
    fracSecs.EntityData.YFilter = fracSecs.YFilter
    fracSecs.EntityData.YangName = "frac-secs"
    fracSecs.EntityData.BundleName = "cisco_ios_xr"
    fracSecs.EntityData.ParentYangName = "transmit-time"
    fracSecs.EntityData.SegmentPath = "frac-secs"
    fracSecs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/transmit-time/" + fracSecs.EntityData.SegmentPath
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = types.NewOrderedMap()
    fracSecs.EntityData.Leafs = types.NewOrderedMap()
    fracSecs.EntityData.Leafs.Append("frac", types.YLeaf{"Frac", fracSecs.Frac})

    fracSecs.EntityData.YListKeys = []string {}

    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail
// Filter Details
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // filter delay. The type is string.
    FilterDelay interface{}

    // filter offset. The type is string.
    FilterOffset interface{}

    // filter disp. The type is string.
    FilterDisp interface{}
}

func (filterDetail *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail) GetEntityData() *types.CommonEntityData {
    filterDetail.EntityData.YFilter = filterDetail.YFilter
    filterDetail.EntityData.YangName = "filter-detail"
    filterDetail.EntityData.BundleName = "cisco_ios_xr"
    filterDetail.EntityData.ParentYangName = "peer-detail-info"
    filterDetail.EntityData.SegmentPath = "filter-detail" + types.AddNoKeyToken(filterDetail)
    filterDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-ntp-admin-oper:ntp/racks/rack/slots/slot/instances/instance/associations-detail/peer-detail-info/" + filterDetail.EntityData.SegmentPath
    filterDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterDetail.EntityData.Children = types.NewOrderedMap()
    filterDetail.EntityData.Leafs = types.NewOrderedMap()
    filterDetail.EntityData.Leafs.Append("filter-delay", types.YLeaf{"FilterDelay", filterDetail.FilterDelay})
    filterDetail.EntityData.Leafs.Append("filter-offset", types.YLeaf{"FilterOffset", filterDetail.FilterOffset})
    filterDetail.EntityData.Leafs.Append("filter-disp", types.YLeaf{"FilterDisp", filterDetail.FilterDisp})

    filterDetail.EntityData.YListKeys = []string {}

    return &(filterDetail.EntityData)
}

