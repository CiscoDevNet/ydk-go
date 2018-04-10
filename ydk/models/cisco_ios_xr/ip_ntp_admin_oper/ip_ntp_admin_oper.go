// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-ntp package
// admin-plane operational data.
// 
// This module contains definitions
// for the following management objects:
//   ntp: NTP admin operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    ntp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntp.EntityData.Children = make(map[string]types.YChild)
    ntp.EntityData.Children["racks"] = types.YChild{"Racks", &ntp.Racks}
    ntp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ntp.EntityData)
}

// Ntp_Racks
// Rack-specific NTP operational data
type Ntp_Racks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular rack. The type is slice of
    // Ntp_Racks_Rack.
    Rack []Ntp_Racks_Rack
}

func (racks *Ntp_Racks) GetEntityData() *types.CommonEntityData {
    racks.EntityData.YFilter = racks.YFilter
    racks.EntityData.YangName = "racks"
    racks.EntityData.BundleName = "cisco_ios_xr"
    racks.EntityData.ParentYangName = "ntp"
    racks.EntityData.SegmentPath = "racks"
    racks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    racks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    racks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    racks.EntityData.Children = make(map[string]types.YChild)
    racks.EntityData.Children["rack"] = types.YChild{"Rack", nil}
    for i := range racks.Rack {
        racks.EntityData.Children[types.GetSegmentPath(&racks.Rack[i])] = types.YChild{"Rack", &racks.Rack[i]}
    }
    racks.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(racks.EntityData)
}

// Ntp_Racks_Rack
// NTP operational data for a particular rack
type Ntp_Racks_Rack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The rack number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Node-specific NTP operational data.
    Slots Ntp_Racks_Rack_Slots
}

func (rack *Ntp_Racks_Rack) GetEntityData() *types.CommonEntityData {
    rack.EntityData.YFilter = rack.YFilter
    rack.EntityData.YangName = "rack"
    rack.EntityData.BundleName = "cisco_ios_xr"
    rack.EntityData.ParentYangName = "racks"
    rack.EntityData.SegmentPath = "rack" + "[number='" + fmt.Sprintf("%v", rack.Number) + "']"
    rack.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rack.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rack.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rack.EntityData.Children = make(map[string]types.YChild)
    rack.EntityData.Children["slots"] = types.YChild{"Slots", &rack.Slots}
    rack.EntityData.Leafs = make(map[string]types.YLeaf)
    rack.EntityData.Leafs["number"] = types.YLeaf{"Number", rack.Number}
    return &(rack.EntityData)
}

// Ntp_Racks_Rack_Slots
// Node-specific NTP operational data
type Ntp_Racks_Rack_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular slot. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot.
    Slot []Ntp_Racks_Rack_Slots_Slot
}

func (slots *Ntp_Racks_Rack_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "rack"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = make(map[string]types.YChild)
    slots.EntityData.Children["slot"] = types.YChild{"Slot", nil}
    for i := range slots.Slot {
        slots.EntityData.Children[types.GetSegmentPath(&slots.Slot[i])] = types.YChild{"Slot", &slots.Slot[i]}
    }
    slots.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slots.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot
// NTP operational data for a particular slot
type Ntp_Racks_Rack_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The slot number. The type is interface{} with
    // range: -2147483648..2147483647.
    Number interface{}

    // Instance-specific NTP operational data.
    Instances Ntp_Racks_Rack_Slots_Slot_Instances
}

func (slot *Ntp_Racks_Rack_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + "[number='" + fmt.Sprintf("%v", slot.Number) + "']"
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = make(map[string]types.YChild)
    slot.EntityData.Children["instances"] = types.YChild{"Instances", &slot.Instances}
    slot.EntityData.Leafs = make(map[string]types.YLeaf)
    slot.EntityData.Leafs["number"] = types.YLeaf{"Number", slot.Number}
    return &(slot.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances
// Instance-specific NTP operational data
type Ntp_Racks_Rack_Slots_Slot_Instances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NTP operational data for a particular instance. The type is slice of
    // Ntp_Racks_Rack_Slots_Slot_Instances_Instance.
    Instance []Ntp_Racks_Rack_Slots_Slot_Instances_Instance
}

func (instances *Ntp_Racks_Rack_Slots_Slot_Instances) GetEntityData() *types.CommonEntityData {
    instances.EntityData.YFilter = instances.YFilter
    instances.EntityData.YangName = "instances"
    instances.EntityData.BundleName = "cisco_ios_xr"
    instances.EntityData.ParentYangName = "slot"
    instances.EntityData.SegmentPath = "instances"
    instances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instances.EntityData.Children = make(map[string]types.YChild)
    instances.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range instances.Instance {
        instances.EntityData.Children[types.GetSegmentPath(&instances.Instance[i])] = types.YChild{"Instance", &instances.Instance[i]}
    }
    instances.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(instances.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance
// NTP operational data for a particular
// instance
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The instance number. The type is interface{} with
    // range: -2147483648..2147483647.
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
    instance.EntityData.SegmentPath = "instance" + "[number='" + fmt.Sprintf("%v", instance.Number) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["status"] = types.YChild{"Status", &instance.Status}
    instance.EntityData.Children["associations"] = types.YChild{"Associations", &instance.Associations}
    instance.EntityData.Children["associations-detail"] = types.YChild{"AssociationsDetail", &instance.AssociationsDetail}
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["number"] = types.YLeaf{"Number", instance.Number}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Children["sys-ref-time"] = types.YChild{"SysRefTime", &status.SysRefTime}
    status.EntityData.Children["sys-drift"] = types.YChild{"SysDrift", &status.SysDrift}
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["is-ntp-enabled"] = types.YLeaf{"IsNtpEnabled", status.IsNtpEnabled}
    status.EntityData.Leafs["sys-dispersion"] = types.YLeaf{"SysDispersion", status.SysDispersion}
    status.EntityData.Leafs["sys-offset"] = types.YLeaf{"SysOffset", status.SysOffset}
    status.EntityData.Leafs["clock-period"] = types.YLeaf{"ClockPeriod", status.ClockPeriod}
    status.EntityData.Leafs["sys-leap"] = types.YLeaf{"SysLeap", status.SysLeap}
    status.EntityData.Leafs["sys-precision"] = types.YLeaf{"SysPrecision", status.SysPrecision}
    status.EntityData.Leafs["sys-stratum"] = types.YLeaf{"SysStratum", status.SysStratum}
    status.EntityData.Leafs["sys-ref-id"] = types.YLeaf{"SysRefId", status.SysRefId}
    status.EntityData.Leafs["sys-root-delay"] = types.YLeaf{"SysRootDelay", status.SysRootDelay}
    status.EntityData.Leafs["sys-root-dispersion"] = types.YLeaf{"SysRootDispersion", status.SysRootDispersion}
    status.EntityData.Leafs["loop-filter-state"] = types.YLeaf{"LoopFilterState", status.LoopFilterState}
    status.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", status.PollInterval}
    status.EntityData.Leafs["is-updated"] = types.YLeaf{"IsUpdated", status.IsUpdated}
    status.EntityData.Leafs["last-update"] = types.YLeaf{"LastUpdate", status.LastUpdate}
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
    sysRefTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysRefTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysRefTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysRefTime.EntityData.Children = make(map[string]types.YChild)
    sysRefTime.EntityData.Children["sec"] = types.YChild{"Sec", &sysRefTime.Sec}
    sysRefTime.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &sysRefTime.FracSecs}
    sysRefTime.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
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
    sysDrift.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sysDrift.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sysDrift.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sysDrift.EntityData.Children = make(map[string]types.YChild)
    sysDrift.EntityData.Children["sec"] = types.YChild{"Sec", &sysDrift.Sec}
    sysDrift.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &sysDrift.FracSecs}
    sysDrift.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
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
    PeerSummaryInfo []Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo
}

func (associations *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations) GetEntityData() *types.CommonEntityData {
    associations.EntityData.YFilter = associations.YFilter
    associations.EntityData.YangName = "associations"
    associations.EntityData.BundleName = "cisco_ios_xr"
    associations.EntityData.ParentYangName = "instance"
    associations.EntityData.SegmentPath = "associations"
    associations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associations.EntityData.Children = make(map[string]types.YChild)
    associations.EntityData.Children["peer-summary-info"] = types.YChild{"PeerSummaryInfo", nil}
    for i := range associations.PeerSummaryInfo {
        associations.EntityData.Children[types.GetSegmentPath(&associations.PeerSummaryInfo[i])] = types.YChild{"PeerSummaryInfo", &associations.PeerSummaryInfo[i]}
    }
    associations.EntityData.Leafs = make(map[string]types.YLeaf)
    associations.EntityData.Leafs["is-ntp-enabled"] = types.YLeaf{"IsNtpEnabled", associations.IsNtpEnabled}
    associations.EntityData.Leafs["sys-leap"] = types.YLeaf{"SysLeap", associations.SysLeap}
    return &(associations.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo
// Peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_Associations_PeerSummaryInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    peerSummaryInfo.EntityData.SegmentPath = "peer-summary-info"
    peerSummaryInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerSummaryInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerSummaryInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerSummaryInfo.EntityData.Children = make(map[string]types.YChild)
    peerSummaryInfo.EntityData.Children["peer-info-common"] = types.YChild{"PeerInfoCommon", &peerSummaryInfo.PeerInfoCommon}
    peerSummaryInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerSummaryInfo.EntityData.Leafs["time-since"] = types.YLeaf{"TimeSince", peerSummaryInfo.TimeSince}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    peerInfoCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfoCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfoCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfoCommon.EntityData.Children = make(map[string]types.YChild)
    peerInfoCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    peerInfoCommon.EntityData.Leafs["host-mode"] = types.YLeaf{"HostMode", peerInfoCommon.HostMode}
    peerInfoCommon.EntityData.Leafs["is-configured"] = types.YLeaf{"IsConfigured", peerInfoCommon.IsConfigured}
    peerInfoCommon.EntityData.Leafs["address"] = types.YLeaf{"Address", peerInfoCommon.Address}
    peerInfoCommon.EntityData.Leafs["reference-id"] = types.YLeaf{"ReferenceId", peerInfoCommon.ReferenceId}
    peerInfoCommon.EntityData.Leafs["host-poll"] = types.YLeaf{"HostPoll", peerInfoCommon.HostPoll}
    peerInfoCommon.EntityData.Leafs["reachability"] = types.YLeaf{"Reachability", peerInfoCommon.Reachability}
    peerInfoCommon.EntityData.Leafs["stratum"] = types.YLeaf{"Stratum", peerInfoCommon.Stratum}
    peerInfoCommon.EntityData.Leafs["status"] = types.YLeaf{"Status", peerInfoCommon.Status}
    peerInfoCommon.EntityData.Leafs["delay"] = types.YLeaf{"Delay", peerInfoCommon.Delay}
    peerInfoCommon.EntityData.Leafs["offset"] = types.YLeaf{"Offset", peerInfoCommon.Offset}
    peerInfoCommon.EntityData.Leafs["dispersion"] = types.YLeaf{"Dispersion", peerInfoCommon.Dispersion}
    peerInfoCommon.EntityData.Leafs["is-sys-peer"] = types.YLeaf{"IsSysPeer", peerInfoCommon.IsSysPeer}
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
    PeerDetailInfo []Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo
}

func (associationsDetail *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail) GetEntityData() *types.CommonEntityData {
    associationsDetail.EntityData.YFilter = associationsDetail.YFilter
    associationsDetail.EntityData.YangName = "associations-detail"
    associationsDetail.EntityData.BundleName = "cisco_ios_xr"
    associationsDetail.EntityData.ParentYangName = "instance"
    associationsDetail.EntityData.SegmentPath = "associations-detail"
    associationsDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationsDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationsDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationsDetail.EntityData.Children = make(map[string]types.YChild)
    associationsDetail.EntityData.Children["peer-detail-info"] = types.YChild{"PeerDetailInfo", nil}
    for i := range associationsDetail.PeerDetailInfo {
        associationsDetail.EntityData.Children[types.GetSegmentPath(&associationsDetail.PeerDetailInfo[i])] = types.YChild{"PeerDetailInfo", &associationsDetail.PeerDetailInfo[i]}
    }
    associationsDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    associationsDetail.EntityData.Leafs["is-ntp-enabled"] = types.YLeaf{"IsNtpEnabled", associationsDetail.IsNtpEnabled}
    associationsDetail.EntityData.Leafs["sys-leap"] = types.YLeaf{"SysLeap", associationsDetail.SysLeap}
    return &(associationsDetail.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo
// Peer info
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    FilterDetail []Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail
}

func (peerDetailInfo *Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "associations-detail"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info"
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = make(map[string]types.YChild)
    peerDetailInfo.EntityData.Children["peer-info-common"] = types.YChild{"PeerInfoCommon", &peerDetailInfo.PeerInfoCommon}
    peerDetailInfo.EntityData.Children["ref-time"] = types.YChild{"RefTime", &peerDetailInfo.RefTime}
    peerDetailInfo.EntityData.Children["originate-time"] = types.YChild{"OriginateTime", &peerDetailInfo.OriginateTime}
    peerDetailInfo.EntityData.Children["receive-time"] = types.YChild{"ReceiveTime", &peerDetailInfo.ReceiveTime}
    peerDetailInfo.EntityData.Children["transmit-time"] = types.YChild{"TransmitTime", &peerDetailInfo.TransmitTime}
    peerDetailInfo.EntityData.Children["filter-detail"] = types.YChild{"FilterDetail", nil}
    for i := range peerDetailInfo.FilterDetail {
        peerDetailInfo.EntityData.Children[types.GetSegmentPath(&peerDetailInfo.FilterDetail[i])] = types.YChild{"FilterDetail", &peerDetailInfo.FilterDetail[i]}
    }
    peerDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerDetailInfo.EntityData.Leafs["leap"] = types.YLeaf{"Leap", peerDetailInfo.Leap}
    peerDetailInfo.EntityData.Leafs["peer-mode"] = types.YLeaf{"PeerMode", peerDetailInfo.PeerMode}
    peerDetailInfo.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", peerDetailInfo.PollInterval}
    peerDetailInfo.EntityData.Leafs["is-ref-clock"] = types.YLeaf{"IsRefClock", peerDetailInfo.IsRefClock}
    peerDetailInfo.EntityData.Leafs["is-authenticated"] = types.YLeaf{"IsAuthenticated", peerDetailInfo.IsAuthenticated}
    peerDetailInfo.EntityData.Leafs["root-delay"] = types.YLeaf{"RootDelay", peerDetailInfo.RootDelay}
    peerDetailInfo.EntityData.Leafs["root-dispersion"] = types.YLeaf{"RootDispersion", peerDetailInfo.RootDispersion}
    peerDetailInfo.EntityData.Leafs["synch-distance"] = types.YLeaf{"SynchDistance", peerDetailInfo.SynchDistance}
    peerDetailInfo.EntityData.Leafs["precision"] = types.YLeaf{"Precision", peerDetailInfo.Precision}
    peerDetailInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", peerDetailInfo.Version}
    peerDetailInfo.EntityData.Leafs["filter-index"] = types.YLeaf{"FilterIndex", peerDetailInfo.FilterIndex}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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
    peerInfoCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfoCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfoCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfoCommon.EntityData.Children = make(map[string]types.YChild)
    peerInfoCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    peerInfoCommon.EntityData.Leafs["host-mode"] = types.YLeaf{"HostMode", peerInfoCommon.HostMode}
    peerInfoCommon.EntityData.Leafs["is-configured"] = types.YLeaf{"IsConfigured", peerInfoCommon.IsConfigured}
    peerInfoCommon.EntityData.Leafs["address"] = types.YLeaf{"Address", peerInfoCommon.Address}
    peerInfoCommon.EntityData.Leafs["reference-id"] = types.YLeaf{"ReferenceId", peerInfoCommon.ReferenceId}
    peerInfoCommon.EntityData.Leafs["host-poll"] = types.YLeaf{"HostPoll", peerInfoCommon.HostPoll}
    peerInfoCommon.EntityData.Leafs["reachability"] = types.YLeaf{"Reachability", peerInfoCommon.Reachability}
    peerInfoCommon.EntityData.Leafs["stratum"] = types.YLeaf{"Stratum", peerInfoCommon.Stratum}
    peerInfoCommon.EntityData.Leafs["status"] = types.YLeaf{"Status", peerInfoCommon.Status}
    peerInfoCommon.EntityData.Leafs["delay"] = types.YLeaf{"Delay", peerInfoCommon.Delay}
    peerInfoCommon.EntityData.Leafs["offset"] = types.YLeaf{"Offset", peerInfoCommon.Offset}
    peerInfoCommon.EntityData.Leafs["dispersion"] = types.YLeaf{"Dispersion", peerInfoCommon.Dispersion}
    peerInfoCommon.EntityData.Leafs["is-sys-peer"] = types.YLeaf{"IsSysPeer", peerInfoCommon.IsSysPeer}
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
    refTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    refTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    refTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    refTime.EntityData.Children = make(map[string]types.YChild)
    refTime.EntityData.Children["sec"] = types.YChild{"Sec", &refTime.Sec}
    refTime.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &refTime.FracSecs}
    refTime.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
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
    originateTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    originateTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    originateTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    originateTime.EntityData.Children = make(map[string]types.YChild)
    originateTime.EntityData.Children["sec"] = types.YChild{"Sec", &originateTime.Sec}
    originateTime.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &originateTime.FracSecs}
    originateTime.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
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
    receiveTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveTime.EntityData.Children = make(map[string]types.YChild)
    receiveTime.EntityData.Children["sec"] = types.YChild{"Sec", &receiveTime.Sec}
    receiveTime.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &receiveTime.FracSecs}
    receiveTime.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
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
    transmitTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitTime.EntityData.Children = make(map[string]types.YChild)
    transmitTime.EntityData.Children["sec"] = types.YChild{"Sec", &transmitTime.Sec}
    transmitTime.EntityData.Children["frac-secs"] = types.YChild{"FracSecs", &transmitTime.FracSecs}
    transmitTime.EntityData.Leafs = make(map[string]types.YLeaf)
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
    sec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sec.EntityData.Children = make(map[string]types.YChild)
    sec.EntityData.Leafs = make(map[string]types.YLeaf)
    sec.EntityData.Leafs["int"] = types.YLeaf{"Int", sec.Int}
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
    fracSecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fracSecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fracSecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fracSecs.EntityData.Children = make(map[string]types.YChild)
    fracSecs.EntityData.Leafs = make(map[string]types.YLeaf)
    fracSecs.EntityData.Leafs["frac"] = types.YLeaf{"Frac", fracSecs.Frac}
    return &(fracSecs.EntityData)
}

// Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail
// Filter Details
type Ntp_Racks_Rack_Slots_Slot_Instances_Instance_AssociationsDetail_PeerDetailInfo_FilterDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    filterDetail.EntityData.SegmentPath = "filter-detail"
    filterDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterDetail.EntityData.Children = make(map[string]types.YChild)
    filterDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    filterDetail.EntityData.Leafs["filter-delay"] = types.YLeaf{"FilterDelay", filterDetail.FilterDelay}
    filterDetail.EntityData.Leafs["filter-offset"] = types.YLeaf{"FilterOffset", filterDetail.FilterOffset}
    filterDetail.EntityData.Leafs["filter-disp"] = types.YLeaf{"FilterDisp", filterDetail.FilterDisp}
    return &(filterDetail.EntityData)
}

