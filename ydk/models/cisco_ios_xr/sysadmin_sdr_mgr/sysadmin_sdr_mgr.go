// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the SDR-SM support config for SDR
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_sdr_mgr

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_sdr_mgr"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sdr-mgr sdr-config}", reflect.TypeOf(SdrConfig{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config", reflect.TypeOf(SdrConfig{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sdr-mgr sdr-manager}", reflect.TypeOf(SdrManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager", reflect.TypeOf(SdrManager{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sdr-mgr sdr-operation}", reflect.TypeOf(SdrOperation{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-operation", reflect.TypeOf(SdrOperation{}))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-sdr-mgr private-sdr}", reflect.TypeOf(PrivateSdr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-sdr-mgr:private-sdr", reflect.TypeOf(PrivateSdr{}))
}

// CardType represents The List of supported Card Types
type CardType string

const (
    CardType_RP CardType = "RP"

    CardType_LC CardType = "LC"

    CardType_CC CardType = "CC"
)

// VmReloadReason represents The List of vm reload reasons
type VmReloadReason string

const (
    VmReloadReason_CARD_OFFLINE VmReloadReason = "CARD_OFFLINE"

    VmReloadReason_CARD_SHUTDOWN VmReloadReason = "CARD_SHUTDOWN"

    VmReloadReason_ALL_VM_RELOAD VmReloadReason = "ALL_VM_RELOAD"

    VmReloadReason_VM_REQUESTED_GRACEFUL_RELOAD VmReloadReason = "VM_REQUESTED_GRACEFUL_RELOAD"

    VmReloadReason_VM_REQUESTED_UNGRACEFUL_RELOAD VmReloadReason = "VM_REQUESTED_UNGRACEFUL_RELOAD"

    VmReloadReason_SDR_CLI_REQUESTED VmReloadReason = "SDR_CLI_REQUESTED"

    VmReloadReason_SDR_VCPU_VMEM_CHANGED VmReloadReason = "SDR_VCPU_VMEM_CHANGED"

    VmReloadReason_SDR_HEARTBEAT_FAILURE VmReloadReason = "SDR_HEARTBEAT_FAILURE"

    VmReloadReason_FIRST_BOOT VmReloadReason = "FIRST_BOOT"

    VmReloadReason_SMU VmReloadReason = "SMU"

    VmReloadReason_REASON_UNKNOWN VmReloadReason = "REASON_UNKNOWN"
)

// SdrConfig
type SdrConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Add/Edit a Secure Domain Router by name. The type is slice of
    // SdrConfig_Sdr.
    Sdr []*SdrConfig_Sdr
}

func (sdrConfig *SdrConfig) GetEntityData() *types.CommonEntityData {
    sdrConfig.EntityData.YFilter = sdrConfig.YFilter
    sdrConfig.EntityData.YangName = "sdr-config"
    sdrConfig.EntityData.BundleName = "cisco_ios_xr"
    sdrConfig.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdr-mgr"
    sdrConfig.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config"
    sdrConfig.EntityData.AbsolutePath = sdrConfig.EntityData.SegmentPath
    sdrConfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrConfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrConfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrConfig.EntityData.Children = types.NewOrderedMap()
    sdrConfig.EntityData.Children.Append("sdr", types.YChild{"Sdr", nil})
    for i := range sdrConfig.Sdr {
        sdrConfig.EntityData.Children.Append(types.GetSegmentPath(sdrConfig.Sdr[i]), types.YChild{"Sdr", sdrConfig.Sdr[i]})
    }
    sdrConfig.EntityData.Leafs = types.NewOrderedMap()

    sdrConfig.EntityData.YListKeys = []string {}

    return &(sdrConfig.EntityData)
}

// SdrConfig_Sdr
// Add/Edit a Secure Domain Router by name
type SdrConfig_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Secure Domain Router , 30 max
    // characters. The type is string with pattern: b'[a-zA-Z0-9_-]{1,30}'.
    Name interface{}

    // List of the initial image and packages for the Secure Domain Router. The
    // type is string.
    InitialImage interface{}

    // Amount of time between lead down to declare SDR down. The type is
    // interface{} with range: 0..4294967295.
    LeadDownDelta interface{}

    // Setting for pairing mode. The type is PairingMode.
    PairingMode interface{}

    // ISSU flag. Once disabled, ISSU won't be performed for this SDR. The type is
    // Issu.
    Issu interface{}

    // Edit resources for a Secure Domain Router.
    Resources SdrConfig_Sdr_Resources

    // Enter list of nodes' location to add to this LR. The type is slice of
    // SdrConfig_Sdr_Location.
    Location []*SdrConfig_Sdr_Location

    
    Action SdrConfig_Sdr_Action

    
    Detail SdrConfig_Sdr_Detail

    
    RebootHistory SdrConfig_Sdr_RebootHistory

    
    Nodes SdrConfig_Sdr_Nodes

    
    Pairing2 SdrConfig_Sdr_Pairing2

    // Add/Edit a RP Pairing by name. The type is slice of SdrConfig_Sdr_Pairing.
    Pairing []*SdrConfig_Sdr_Pairing
}

func (sdr *SdrConfig_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "sdr-config"
    sdr.EntityData.SegmentPath = "sdr" + types.AddKeyToken(sdr.Name, "name")
    sdr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/" + sdr.EntityData.SegmentPath
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = types.NewOrderedMap()
    sdr.EntityData.Children.Append("resources", types.YChild{"Resources", &sdr.Resources})
    sdr.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range sdr.Location {
        sdr.EntityData.Children.Append(types.GetSegmentPath(sdr.Location[i]), types.YChild{"Location", sdr.Location[i]})
    }
    sdr.EntityData.Children.Append("Action", types.YChild{"Action", &sdr.Action})
    sdr.EntityData.Children.Append("detail", types.YChild{"Detail", &sdr.Detail})
    sdr.EntityData.Children.Append("reboot-history", types.YChild{"RebootHistory", &sdr.RebootHistory})
    sdr.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sdr.Nodes})
    sdr.EntityData.Children.Append("pairing2", types.YChild{"Pairing2", &sdr.Pairing2})
    sdr.EntityData.Children.Append("pairing", types.YChild{"Pairing", nil})
    for i := range sdr.Pairing {
        sdr.EntityData.Children.Append(types.GetSegmentPath(sdr.Pairing[i]), types.YChild{"Pairing", sdr.Pairing[i]})
    }
    sdr.EntityData.Leafs = types.NewOrderedMap()
    sdr.EntityData.Leafs.Append("name", types.YLeaf{"Name", sdr.Name})
    sdr.EntityData.Leafs.Append("initial-image", types.YLeaf{"InitialImage", sdr.InitialImage})
    sdr.EntityData.Leafs.Append("lead_down_delta", types.YLeaf{"LeadDownDelta", sdr.LeadDownDelta})
    sdr.EntityData.Leafs.Append("pairing-mode", types.YLeaf{"PairingMode", sdr.PairingMode})
    sdr.EntityData.Leafs.Append("issu", types.YLeaf{"Issu", sdr.Issu})

    sdr.EntityData.YListKeys = []string {"Name"}

    return &(sdr.EntityData)
}

// SdrConfig_Sdr_Resources
// Edit resources for a Secure Domain Router
type SdrConfig_Sdr_Resources struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fgids for a Secure Domain Router. The type is interface{} with range:
    // 25000..524288.
    Fgid interface{}

    // Management External VLAN for Secure Domain Router. The type is interface{}
    // with range: 2..4094.
    MgmtExtVlan interface{}

    // The type is slice of SdrConfig_Sdr_Resources_CardType.
    CardType []*SdrConfig_Sdr_Resources_CardType
}

func (resources *SdrConfig_Sdr_Resources) GetEntityData() *types.CommonEntityData {
    resources.EntityData.YFilter = resources.YFilter
    resources.EntityData.YangName = "resources"
    resources.EntityData.BundleName = "cisco_ios_xr"
    resources.EntityData.ParentYangName = "sdr"
    resources.EntityData.SegmentPath = "resources"
    resources.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + resources.EntityData.SegmentPath
    resources.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    resources.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    resources.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    resources.EntityData.Children = types.NewOrderedMap()
    resources.EntityData.Children.Append("card-type", types.YChild{"CardType", nil})
    for i := range resources.CardType {
        resources.EntityData.Children.Append(types.GetSegmentPath(resources.CardType[i]), types.YChild{"CardType", resources.CardType[i]})
    }
    resources.EntityData.Leafs = types.NewOrderedMap()
    resources.EntityData.Leafs.Append("fgid", types.YLeaf{"Fgid", resources.Fgid})
    resources.EntityData.Leafs.Append("mgmt_ext_vlan", types.YLeaf{"MgmtExtVlan", resources.MgmtExtVlan})

    resources.EntityData.YListKeys = []string {}

    return &(resources.EntityData)
}

// SdrConfig_Sdr_Resources_CardType
type SdrConfig_Sdr_Resources_CardType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Card Type. The type is CardType.
    Type interface{}

    // VM Memory Size in units of [GB]. The type is interface{} with range:
    // 1..128.
    VmMemory interface{}

    // VM Number of CPUs. The type is interface{} with range: 1..128.
    VmCpu interface{}
}

func (cardType *SdrConfig_Sdr_Resources_CardType) GetEntityData() *types.CommonEntityData {
    cardType.EntityData.YFilter = cardType.YFilter
    cardType.EntityData.YangName = "card-type"
    cardType.EntityData.BundleName = "cisco_ios_xr"
    cardType.EntityData.ParentYangName = "resources"
    cardType.EntityData.SegmentPath = "card-type" + types.AddKeyToken(cardType.Type, "type")
    cardType.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/resources/" + cardType.EntityData.SegmentPath
    cardType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardType.EntityData.Children = types.NewOrderedMap()
    cardType.EntityData.Leafs = types.NewOrderedMap()
    cardType.EntityData.Leafs.Append("type", types.YLeaf{"Type", cardType.Type})
    cardType.EntityData.Leafs.Append("vm-memory", types.YLeaf{"VmMemory", cardType.VmMemory})
    cardType.EntityData.Leafs.Append("vm-cpu", types.YLeaf{"VmCpu", cardType.VmCpu})

    cardType.EntityData.YListKeys = []string {"Type"}

    return &(cardType.EntityData)
}

// SdrConfig_Sdr_Location
// Enter list of nodes' location to add to this LR
type SdrConfig_Sdr_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Enter location or all. The type is string with
    // pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/(([rR][pP]|[lL][cC]|[cC][bB])?\\d{1,2}))(/[cC][pP][uU]0)?|all'.
    NodeLocation interface{}

    // Enter list of slices to add to this SDR. The type is slice of interface{}
    // with range: 0..4.
    Slice []interface{}
}

func (location *SdrConfig_Sdr_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "sdr"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("slice", types.YLeaf{"Slice", location.Slice})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_Action
type SdrConfig_Sdr_Action struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrConfig_Sdr_Action_Location.
    Location []*SdrConfig_Sdr_Action_Location
}

func (action *SdrConfig_Sdr_Action) GetEntityData() *types.CommonEntityData {
    action.EntityData.YFilter = action.YFilter
    action.EntityData.YangName = "Action"
    action.EntityData.BundleName = "cisco_ios_xr"
    action.EntityData.ParentYangName = "sdr"
    action.EntityData.SegmentPath = "Action"
    action.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + action.EntityData.SegmentPath
    action.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    action.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    action.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    action.EntityData.Children = types.NewOrderedMap()
    action.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range action.Location {
        action.EntityData.Children.Append(types.GetSegmentPath(action.Location[i]), types.YChild{"Location", action.Location[i]})
    }
    action.EntityData.Leafs = types.NewOrderedMap()

    action.EntityData.YListKeys = []string {}

    return &(action.EntityData)
}

// SdrConfig_Sdr_Action_Location
type SdrConfig_Sdr_Action_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Enter location or all. The type is string with
    // pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/((([rR]([sS]){0,1}[pP])|[cC][bB])?\\d{1,2})/[V][M](0?[0-9]|1[1-5]))?|all'.
    NodeLocation interface{}
}

func (location *SdrConfig_Sdr_Action_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "Action"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/Action/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_Detail
type SdrConfig_Sdr_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrConfig_Sdr_Detail_Location.
    Location []*SdrConfig_Sdr_Detail_Location
}

func (detail *SdrConfig_Sdr_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "sdr"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + detail.EntityData.SegmentPath
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = types.NewOrderedMap()
    detail.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range detail.Location {
        detail.EntityData.Children.Append(types.GetSegmentPath(detail.Location[i]), types.YChild{"Location", detail.Location[i]})
    }
    detail.EntityData.Leafs = types.NewOrderedMap()

    detail.EntityData.YListKeys = []string {}

    return &(detail.EntityData)
}

// SdrConfig_Sdr_Detail_Location
type SdrConfig_Sdr_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/((([rR]([sS]){0,1}[pP])|[cC][bB])?\\d{1,2})/[V][M](0?[0-9]|1[1-5]))?'.
    NodeLocation interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // IP address of VM. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // MAC address of VM. The type is string.
    MacAddress interface{}

    // The type is string.
    BootPart interface{}

    // The type is string.
    DataPart interface{}

    // The type is string.
    BigDisk interface{}

    // The type is interface{} with range: 0..4294967295.
    VmId interface{}

    // The type is interface{} with range: 0..4294967295.
    Vmcpu interface{}

    // The type is interface{} with range: 0..4294967295.
    Vmmemory interface{}

    // The type is string.
    CardType interface{}

    // The type is string.
    CardSerial interface{}

    // The type is string.
    RackType interface{}

    // Slice numbers if configured. The type is one of the following types: slice
    // of int with range: 0..255, or slice of string.
    Slice []interface{}

    // The type is string.
    ChassisSerial interface{}

    // The type is string.
    HwVersion interface{}

    // The type is string.
    MgmtExtVlan interface{}

    // State of VM. The type is string.
    State interface{}

    // The type is string.
    StartTime interface{}

    // Number of times rebooted since first boot. The type is interface{} with
    // range: 0..4294967295.
    RebootCount interface{}

    // Number of times rebooted since lasr card reload. The type is interface{}
    // with range: 0..4294967295.
    RhCount interface{}

    // The type is slice of SdrConfig_Sdr_Detail_Location_RebootHist1.
    RebootHist1 []*SdrConfig_Sdr_Detail_Location_RebootHist1
}

func (location *SdrConfig_Sdr_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/detail/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("reboot_hist1", types.YChild{"RebootHist1", nil})
    for i := range location.RebootHist1 {
        location.EntityData.Children.Append(types.GetSegmentPath(location.RebootHist1[i]), types.YChild{"RebootHist1", location.RebootHist1[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", location.SdrId})
    location.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", location.IpAddr})
    location.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", location.MacAddress})
    location.EntityData.Leafs.Append("boot_part", types.YLeaf{"BootPart", location.BootPart})
    location.EntityData.Leafs.Append("data_part", types.YLeaf{"DataPart", location.DataPart})
    location.EntityData.Leafs.Append("big_disk", types.YLeaf{"BigDisk", location.BigDisk})
    location.EntityData.Leafs.Append("vm_id", types.YLeaf{"VmId", location.VmId})
    location.EntityData.Leafs.Append("vmcpu", types.YLeaf{"Vmcpu", location.Vmcpu})
    location.EntityData.Leafs.Append("vmmemory", types.YLeaf{"Vmmemory", location.Vmmemory})
    location.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", location.CardType})
    location.EntityData.Leafs.Append("card_serial", types.YLeaf{"CardSerial", location.CardSerial})
    location.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", location.RackType})
    location.EntityData.Leafs.Append("slice", types.YLeaf{"Slice", location.Slice})
    location.EntityData.Leafs.Append("chassis_serial", types.YLeaf{"ChassisSerial", location.ChassisSerial})
    location.EntityData.Leafs.Append("hw_version", types.YLeaf{"HwVersion", location.HwVersion})
    location.EntityData.Leafs.Append("mgmt_ext_vlan", types.YLeaf{"MgmtExtVlan", location.MgmtExtVlan})
    location.EntityData.Leafs.Append("state", types.YLeaf{"State", location.State})
    location.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", location.StartTime})
    location.EntityData.Leafs.Append("reboot_count", types.YLeaf{"RebootCount", location.RebootCount})
    location.EntityData.Leafs.Append("rh_count", types.YLeaf{"RhCount", location.RhCount})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_Detail_Location_RebootHist1
type SdrConfig_Sdr_Detail_Location_RebootHist1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // The type is string.
    Time interface{}

    // Reason for reload. The type is string.
    Reason interface{}
}

func (rebootHist1 *SdrConfig_Sdr_Detail_Location_RebootHist1) GetEntityData() *types.CommonEntityData {
    rebootHist1.EntityData.YFilter = rebootHist1.YFilter
    rebootHist1.EntityData.YangName = "reboot_hist1"
    rebootHist1.EntityData.BundleName = "cisco_ios_xr"
    rebootHist1.EntityData.ParentYangName = "location"
    rebootHist1.EntityData.SegmentPath = "reboot_hist1" + types.AddKeyToken(rebootHist1.Count, "count")
    rebootHist1.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/detail/location/" + rebootHist1.EntityData.SegmentPath
    rebootHist1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHist1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHist1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHist1.EntityData.Children = types.NewOrderedMap()
    rebootHist1.EntityData.Leafs = types.NewOrderedMap()
    rebootHist1.EntityData.Leafs.Append("count", types.YLeaf{"Count", rebootHist1.Count})
    rebootHist1.EntityData.Leafs.Append("Time", types.YLeaf{"Time", rebootHist1.Time})
    rebootHist1.EntityData.Leafs.Append("Reason", types.YLeaf{"Reason", rebootHist1.Reason})

    rebootHist1.EntityData.YListKeys = []string {"Count"}

    return &(rebootHist1.EntityData)
}

// SdrConfig_Sdr_RebootHistory
type SdrConfig_Sdr_RebootHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Reverse SdrConfig_Sdr_RebootHistory_Reverse

    
    DefaultDisp SdrConfig_Sdr_RebootHistory_DefaultDisp
}

func (rebootHistory *SdrConfig_Sdr_RebootHistory) GetEntityData() *types.CommonEntityData {
    rebootHistory.EntityData.YFilter = rebootHistory.YFilter
    rebootHistory.EntityData.YangName = "reboot-history"
    rebootHistory.EntityData.BundleName = "cisco_ios_xr"
    rebootHistory.EntityData.ParentYangName = "sdr"
    rebootHistory.EntityData.SegmentPath = "reboot-history"
    rebootHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + rebootHistory.EntityData.SegmentPath
    rebootHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHistory.EntityData.Children = types.NewOrderedMap()
    rebootHistory.EntityData.Children.Append("reverse", types.YChild{"Reverse", &rebootHistory.Reverse})
    rebootHistory.EntityData.Children.Append("default-disp", types.YChild{"DefaultDisp", &rebootHistory.DefaultDisp})
    rebootHistory.EntityData.Leafs = types.NewOrderedMap()

    rebootHistory.EntityData.YListKeys = []string {}

    return &(rebootHistory.EntityData)
}

// SdrConfig_Sdr_RebootHistory_Reverse
type SdrConfig_Sdr_RebootHistory_Reverse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrConfig_Sdr_RebootHistory_Reverse_Location.
    Location []*SdrConfig_Sdr_RebootHistory_Reverse_Location
}

func (reverse *SdrConfig_Sdr_RebootHistory_Reverse) GetEntityData() *types.CommonEntityData {
    reverse.EntityData.YFilter = reverse.YFilter
    reverse.EntityData.YangName = "reverse"
    reverse.EntityData.BundleName = "cisco_ios_xr"
    reverse.EntityData.ParentYangName = "reboot-history"
    reverse.EntityData.SegmentPath = "reverse"
    reverse.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/" + reverse.EntityData.SegmentPath
    reverse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reverse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reverse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reverse.EntityData.Children = types.NewOrderedMap()
    reverse.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range reverse.Location {
        reverse.EntityData.Children.Append(types.GetSegmentPath(reverse.Location[i]), types.YChild{"Location", reverse.Location[i]})
    }
    reverse.EntityData.Leafs = types.NewOrderedMap()

    reverse.EntityData.YListKeys = []string {}

    return &(reverse.EntityData)
}

// SdrConfig_Sdr_RebootHistory_Reverse_Location
type SdrConfig_Sdr_RebootHistory_Reverse_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/((([rR]([sS]){0,1}[pP])|[cC][bB])?\\d{1,2})/[V][M](0?[0-9]|1[1-5]))?'.
    NodeLocation interface{}

    // Number of times rebooted since first boot. The type is interface{} with
    // range: 0..4294967295.
    RebootCount interface{}

    // Number of times rebooted since last card reload. The type is interface{}
    // with range: 0..4294967295.
    RhCount interface{}

    // The type is slice of
    // SdrConfig_Sdr_RebootHistory_Reverse_Location_RebootHist2.
    RebootHist2 []*SdrConfig_Sdr_RebootHistory_Reverse_Location_RebootHist2
}

func (location *SdrConfig_Sdr_RebootHistory_Reverse_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reverse"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/reverse/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("reboot_hist2", types.YChild{"RebootHist2", nil})
    for i := range location.RebootHist2 {
        location.EntityData.Children.Append(types.GetSegmentPath(location.RebootHist2[i]), types.YChild{"RebootHist2", location.RebootHist2[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("reboot_count", types.YLeaf{"RebootCount", location.RebootCount})
    location.EntityData.Leafs.Append("rh_count", types.YLeaf{"RhCount", location.RhCount})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_RebootHistory_Reverse_Location_RebootHist2
type SdrConfig_Sdr_RebootHistory_Reverse_Location_RebootHist2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // The type is string.
    Time interface{}

    // Reason for reload. The type is string.
    Reason interface{}
}

func (rebootHist2 *SdrConfig_Sdr_RebootHistory_Reverse_Location_RebootHist2) GetEntityData() *types.CommonEntityData {
    rebootHist2.EntityData.YFilter = rebootHist2.YFilter
    rebootHist2.EntityData.YangName = "reboot_hist2"
    rebootHist2.EntityData.BundleName = "cisco_ios_xr"
    rebootHist2.EntityData.ParentYangName = "location"
    rebootHist2.EntityData.SegmentPath = "reboot_hist2" + types.AddKeyToken(rebootHist2.Count, "count")
    rebootHist2.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/reverse/location/" + rebootHist2.EntityData.SegmentPath
    rebootHist2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHist2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHist2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHist2.EntityData.Children = types.NewOrderedMap()
    rebootHist2.EntityData.Leafs = types.NewOrderedMap()
    rebootHist2.EntityData.Leafs.Append("count", types.YLeaf{"Count", rebootHist2.Count})
    rebootHist2.EntityData.Leafs.Append("Time", types.YLeaf{"Time", rebootHist2.Time})
    rebootHist2.EntityData.Leafs.Append("Reason", types.YLeaf{"Reason", rebootHist2.Reason})

    rebootHist2.EntityData.YListKeys = []string {"Count"}

    return &(rebootHist2.EntityData)
}

// SdrConfig_Sdr_RebootHistory_DefaultDisp
type SdrConfig_Sdr_RebootHistory_DefaultDisp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrConfig_Sdr_RebootHistory_DefaultDisp_Location.
    Location []*SdrConfig_Sdr_RebootHistory_DefaultDisp_Location
}

func (defaultDisp *SdrConfig_Sdr_RebootHistory_DefaultDisp) GetEntityData() *types.CommonEntityData {
    defaultDisp.EntityData.YFilter = defaultDisp.YFilter
    defaultDisp.EntityData.YangName = "default-disp"
    defaultDisp.EntityData.BundleName = "cisco_ios_xr"
    defaultDisp.EntityData.ParentYangName = "reboot-history"
    defaultDisp.EntityData.SegmentPath = "default-disp"
    defaultDisp.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/" + defaultDisp.EntityData.SegmentPath
    defaultDisp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultDisp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultDisp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultDisp.EntityData.Children = types.NewOrderedMap()
    defaultDisp.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range defaultDisp.Location {
        defaultDisp.EntityData.Children.Append(types.GetSegmentPath(defaultDisp.Location[i]), types.YChild{"Location", defaultDisp.Location[i]})
    }
    defaultDisp.EntityData.Leafs = types.NewOrderedMap()

    defaultDisp.EntityData.YListKeys = []string {}

    return &(defaultDisp.EntityData)
}

// SdrConfig_Sdr_RebootHistory_DefaultDisp_Location
type SdrConfig_Sdr_RebootHistory_DefaultDisp_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/((([rR]([sS]){0,1}[pP])|[cC][bB])?\\d{1,2})/[V][M](0?[0-9]|1[1-5]))?'.
    NodeLocation interface{}

    // Number of times rebooted since first boot. The type is interface{} with
    // range: 0..4294967295.
    RebootCount interface{}

    // Number of times rebooted since last card reload. The type is interface{}
    // with range: 0..4294967295.
    RhCount interface{}

    // The type is slice of
    // SdrConfig_Sdr_RebootHistory_DefaultDisp_Location_RebootHist2.
    RebootHist2 []*SdrConfig_Sdr_RebootHistory_DefaultDisp_Location_RebootHist2
}

func (location *SdrConfig_Sdr_RebootHistory_DefaultDisp_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "default-disp"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/default-disp/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("reboot_hist2", types.YChild{"RebootHist2", nil})
    for i := range location.RebootHist2 {
        location.EntityData.Children.Append(types.GetSegmentPath(location.RebootHist2[i]), types.YChild{"RebootHist2", location.RebootHist2[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("reboot_count", types.YLeaf{"RebootCount", location.RebootCount})
    location.EntityData.Leafs.Append("rh_count", types.YLeaf{"RhCount", location.RhCount})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_RebootHistory_DefaultDisp_Location_RebootHist2
type SdrConfig_Sdr_RebootHistory_DefaultDisp_Location_RebootHist2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Count interface{}

    // The type is string.
    Time interface{}

    // Reason for reload. The type is string.
    Reason interface{}
}

func (rebootHist2 *SdrConfig_Sdr_RebootHistory_DefaultDisp_Location_RebootHist2) GetEntityData() *types.CommonEntityData {
    rebootHist2.EntityData.YFilter = rebootHist2.YFilter
    rebootHist2.EntityData.YangName = "reboot_hist2"
    rebootHist2.EntityData.BundleName = "cisco_ios_xr"
    rebootHist2.EntityData.ParentYangName = "location"
    rebootHist2.EntityData.SegmentPath = "reboot_hist2" + types.AddKeyToken(rebootHist2.Count, "count")
    rebootHist2.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/reboot-history/default-disp/location/" + rebootHist2.EntityData.SegmentPath
    rebootHist2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rebootHist2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rebootHist2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rebootHist2.EntityData.Children = types.NewOrderedMap()
    rebootHist2.EntityData.Leafs = types.NewOrderedMap()
    rebootHist2.EntityData.Leafs.Append("count", types.YLeaf{"Count", rebootHist2.Count})
    rebootHist2.EntityData.Leafs.Append("Time", types.YLeaf{"Time", rebootHist2.Time})
    rebootHist2.EntityData.Leafs.Append("Reason", types.YLeaf{"Reason", rebootHist2.Reason})

    rebootHist2.EntityData.YListKeys = []string {"Count"}

    return &(rebootHist2.EntityData)
}

// SdrConfig_Sdr_Nodes
type SdrConfig_Sdr_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrConfig_Sdr_Nodes_Location.
    Location []*SdrConfig_Sdr_Nodes_Location
}

func (nodes *SdrConfig_Sdr_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "sdr"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nodes.Location {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Location[i]), types.YChild{"Location", nodes.Location[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// SdrConfig_Sdr_Nodes_Location
type SdrConfig_Sdr_Nodes_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/((([rR]([sS]){0,1}[pP])|[cC][bB])?\\d{1,2})/[V][M](0?[0-9]|1[1-5]))?'.
    NodeLocation interface{}

    // The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // IP address of VM. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // MAC address of VM. The type is string.
    MacAddress interface{}

    // State of VM. The type is string.
    State interface{}

    // The type is string.
    StartTime interface{}

    // Reason for last reload. The type is string.
    ReloadReason interface{}

    // Number of times rebooted since first boot. The type is interface{} with
    // range: 0..4294967295.
    RebootCount interface{}

    // Number of times rebooted since first boot. The type is interface{} with
    // range: 0..4294967295.
    RhCount interface{}
}

func (location *SdrConfig_Sdr_Nodes_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nodes"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/nodes/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", location.SdrId})
    location.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", location.IpAddr})
    location.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", location.MacAddress})
    location.EntityData.Leafs.Append("state", types.YLeaf{"State", location.State})
    location.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", location.StartTime})
    location.EntityData.Leafs.Append("reload_reason", types.YLeaf{"ReloadReason", location.ReloadReason})
    location.EntityData.Leafs.Append("reboot_count", types.YLeaf{"RebootCount", location.RebootCount})
    location.EntityData.Leafs.Append("rh_count", types.YLeaf{"RhCount", location.RhCount})

    location.EntityData.YListKeys = []string {"NodeLocation"}

    return &(location.EntityData)
}

// SdrConfig_Sdr_Pairing2
type SdrConfig_Sdr_Pairing2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mode of Pairing. The type is string.
    PairingMode interface{}

    
    Sdrlead SdrConfig_Sdr_Pairing2_Sdrlead

    // The type is slice of SdrConfig_Sdr_Pairing2_Pairing.
    Pairing []*SdrConfig_Sdr_Pairing2_Pairing
}

func (pairing2 *SdrConfig_Sdr_Pairing2) GetEntityData() *types.CommonEntityData {
    pairing2.EntityData.YFilter = pairing2.YFilter
    pairing2.EntityData.YangName = "pairing2"
    pairing2.EntityData.BundleName = "cisco_ios_xr"
    pairing2.EntityData.ParentYangName = "sdr"
    pairing2.EntityData.SegmentPath = "pairing2"
    pairing2.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + pairing2.EntityData.SegmentPath
    pairing2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pairing2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pairing2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pairing2.EntityData.Children = types.NewOrderedMap()
    pairing2.EntityData.Children.Append("sdrlead", types.YChild{"Sdrlead", &pairing2.Sdrlead})
    pairing2.EntityData.Children.Append("pairing", types.YChild{"Pairing", nil})
    for i := range pairing2.Pairing {
        pairing2.EntityData.Children.Append(types.GetSegmentPath(pairing2.Pairing[i]), types.YChild{"Pairing", pairing2.Pairing[i]})
    }
    pairing2.EntityData.Leafs = types.NewOrderedMap()
    pairing2.EntityData.Leafs.Append("pairing-mode", types.YLeaf{"PairingMode", pairing2.PairingMode})

    pairing2.EntityData.YListKeys = []string {}

    return &(pairing2.EntityData)
}

// SdrConfig_Sdr_Pairing2_Sdrlead
type SdrConfig_Sdr_Pairing2_Sdrlead struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Rp1 interface{}

    // The type is string.
    Rp2 interface{}
}

func (sdrlead *SdrConfig_Sdr_Pairing2_Sdrlead) GetEntityData() *types.CommonEntityData {
    sdrlead.EntityData.YFilter = sdrlead.YFilter
    sdrlead.EntityData.YangName = "sdrlead"
    sdrlead.EntityData.BundleName = "cisco_ios_xr"
    sdrlead.EntityData.ParentYangName = "pairing2"
    sdrlead.EntityData.SegmentPath = "sdrlead"
    sdrlead.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/pairing2/" + sdrlead.EntityData.SegmentPath
    sdrlead.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrlead.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrlead.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrlead.EntityData.Children = types.NewOrderedMap()
    sdrlead.EntityData.Leafs = types.NewOrderedMap()
    sdrlead.EntityData.Leafs.Append("rp1", types.YLeaf{"Rp1", sdrlead.Rp1})
    sdrlead.EntityData.Leafs.Append("rp2", types.YLeaf{"Rp2", sdrlead.Rp2})

    sdrlead.EntityData.YListKeys = []string {}

    return &(sdrlead.EntityData)
}

// SdrConfig_Sdr_Pairing2_Pairing
type SdrConfig_Sdr_Pairing2_Pairing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is string.
    Rp1 interface{}

    // The type is string.
    Rp2 interface{}
}

func (pairing *SdrConfig_Sdr_Pairing2_Pairing) GetEntityData() *types.CommonEntityData {
    pairing.EntityData.YFilter = pairing.YFilter
    pairing.EntityData.YangName = "pairing"
    pairing.EntityData.BundleName = "cisco_ios_xr"
    pairing.EntityData.ParentYangName = "pairing2"
    pairing.EntityData.SegmentPath = "pairing" + types.AddKeyToken(pairing.Name, "name")
    pairing.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/pairing2/" + pairing.EntityData.SegmentPath
    pairing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pairing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pairing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pairing.EntityData.Children = types.NewOrderedMap()
    pairing.EntityData.Leafs = types.NewOrderedMap()
    pairing.EntityData.Leafs.Append("name", types.YLeaf{"Name", pairing.Name})
    pairing.EntityData.Leafs.Append("rp1", types.YLeaf{"Rp1", pairing.Rp1})
    pairing.EntityData.Leafs.Append("rp2", types.YLeaf{"Rp2", pairing.Rp2})

    pairing.EntityData.YListKeys = []string {"Name"}

    return &(pairing.EntityData)
}

// SdrConfig_Sdr_Pairing
// Add/Edit a RP Pairing by name
type SdrConfig_Sdr_Pairing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with pattern:
    // b'[a-zA-Z0-9_-]{1,64}'.
    Name interface{}

    // Enter RP Node location. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/(([rR][pP]|[cC][bB])\\d{1,2}))(/[cC][pP][uU]0)?'.
    // This attribute is mandatory.
    Rp1 interface{}

    // Enter RP Node location. The type is string with pattern:
    // b'((0?[0-9]|1[1-5]|[bB]\\d)/(([rR][pP]|[cC][bB])\\d{1,2}))(/[cC][pP][uU]0)?'.
    // This attribute is mandatory.
    Rp2 interface{}
}

func (pairing *SdrConfig_Sdr_Pairing) GetEntityData() *types.CommonEntityData {
    pairing.EntityData.YFilter = pairing.YFilter
    pairing.EntityData.YangName = "pairing"
    pairing.EntityData.BundleName = "cisco_ios_xr"
    pairing.EntityData.ParentYangName = "sdr"
    pairing.EntityData.SegmentPath = "pairing" + types.AddKeyToken(pairing.Name, "name")
    pairing.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-config/sdr/" + pairing.EntityData.SegmentPath
    pairing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pairing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pairing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pairing.EntityData.Children = types.NewOrderedMap()
    pairing.EntityData.Leafs = types.NewOrderedMap()
    pairing.EntityData.Leafs.Append("name", types.YLeaf{"Name", pairing.Name})
    pairing.EntityData.Leafs.Append("rp1", types.YLeaf{"Rp1", pairing.Rp1})
    pairing.EntityData.Leafs.Append("rp2", types.YLeaf{"Rp2", pairing.Rp2})

    pairing.EntityData.YListKeys = []string {"Name"}

    return &(pairing.EntityData)
}

// SdrConfig_Sdr_Issu represents ISSU flag. Once disabled, ISSU won't be performed for this SDR.
type SdrConfig_Sdr_Issu string

const (
    SdrConfig_Sdr_Issu_disable SdrConfig_Sdr_Issu = "disable"
)

// SdrConfig_Sdr_PairingMode represents Setting for pairing mode
type SdrConfig_Sdr_PairingMode string

const (
    SdrConfig_Sdr_PairingMode_intra_rack SdrConfig_Sdr_PairingMode = "intra-rack"

    SdrConfig_Sdr_PairingMode_inter_rack SdrConfig_Sdr_PairingMode = "inter-rack"
)

// SdrManager
type SdrManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SdrMgr SdrManager_SdrMgr
}

func (sdrManager *SdrManager) GetEntityData() *types.CommonEntityData {
    sdrManager.EntityData.YFilter = sdrManager.YFilter
    sdrManager.EntityData.YangName = "sdr-manager"
    sdrManager.EntityData.BundleName = "cisco_ios_xr"
    sdrManager.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdr-mgr"
    sdrManager.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager"
    sdrManager.EntityData.AbsolutePath = sdrManager.EntityData.SegmentPath
    sdrManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrManager.EntityData.Children = types.NewOrderedMap()
    sdrManager.EntityData.Children.Append("sdr_mgr", types.YChild{"SdrMgr", &sdrManager.SdrMgr})
    sdrManager.EntityData.Leafs = types.NewOrderedMap()

    sdrManager.EntityData.YListKeys = []string {}

    return &(sdrManager.EntityData)
}

// SdrManager_SdrMgr
type SdrManager_SdrMgr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of SdrManager_SdrMgr_Trace.
    Trace []*SdrManager_SdrMgr_Trace
}

func (sdrMgr *SdrManager_SdrMgr) GetEntityData() *types.CommonEntityData {
    sdrMgr.EntityData.YFilter = sdrMgr.YFilter
    sdrMgr.EntityData.YangName = "sdr_mgr"
    sdrMgr.EntityData.BundleName = "cisco_ios_xr"
    sdrMgr.EntityData.ParentYangName = "sdr-manager"
    sdrMgr.EntityData.SegmentPath = "sdr_mgr"
    sdrMgr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager/" + sdrMgr.EntityData.SegmentPath
    sdrMgr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrMgr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrMgr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrMgr.EntityData.Children = types.NewOrderedMap()
    sdrMgr.EntityData.Children.Append("trace", types.YChild{"Trace", nil})
    for i := range sdrMgr.Trace {
        sdrMgr.EntityData.Children.Append(types.GetSegmentPath(sdrMgr.Trace[i]), types.YChild{"Trace", sdrMgr.Trace[i]})
    }
    sdrMgr.EntityData.Leafs = types.NewOrderedMap()

    sdrMgr.EntityData.YListKeys = []string {}

    return &(sdrMgr.EntityData)
}

// SdrManager_SdrMgr_Trace
// show traceable processes
type SdrManager_SdrMgr_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of SdrManager_SdrMgr_Trace_Location.
    Location []*SdrManager_SdrMgr_Trace_Location
}

func (trace *SdrManager_SdrMgr_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "sdr_mgr"
    trace.EntityData.SegmentPath = "trace" + types.AddKeyToken(trace.Buffer, "buffer")
    trace.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager/sdr_mgr/" + trace.EntityData.SegmentPath
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = types.NewOrderedMap()
    trace.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range trace.Location {
        trace.EntityData.Children.Append(types.GetSegmentPath(trace.Location[i]), types.YChild{"Location", trace.Location[i]})
    }
    trace.EntityData.Leafs = types.NewOrderedMap()
    trace.EntityData.Leafs.Append("buffer", types.YLeaf{"Buffer", trace.Buffer})

    trace.EntityData.YListKeys = []string {"Buffer"}

    return &(trace.EntityData)
}

// SdrManager_SdrMgr_Trace_Location
type SdrManager_SdrMgr_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of SdrManager_SdrMgr_Trace_Location_AllOptions.
    AllOptions []*SdrManager_SdrMgr_Trace_Location_AllOptions
}

func (location *SdrManager_SdrMgr_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.LocationName, "location_name")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager/sdr_mgr/trace/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Children.Append("all-options", types.YChild{"AllOptions", nil})
    for i := range location.AllOptions {
        location.EntityData.Children.Append(types.GetSegmentPath(location.AllOptions[i]), types.YChild{"AllOptions", location.AllOptions[i]})
    }
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("location_name", types.YLeaf{"LocationName", location.LocationName})

    location.EntityData.YListKeys = []string {"LocationName"}

    return &(location.EntityData)
}

// SdrManager_SdrMgr_Trace_Location_AllOptions
type SdrManager_SdrMgr_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // SdrManager_SdrMgr_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []*SdrManager_SdrMgr_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *SdrManager_SdrMgr_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + types.AddKeyToken(allOptions.Option, "option")
    allOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager/sdr_mgr/trace/location/" + allOptions.EntityData.SegmentPath
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = types.NewOrderedMap()
    allOptions.EntityData.Children.Append("trace-blocks", types.YChild{"TraceBlocks", nil})
    for i := range allOptions.TraceBlocks {
        types.SetYListKey(allOptions.TraceBlocks[i], i)
        allOptions.EntityData.Children.Append(types.GetSegmentPath(allOptions.TraceBlocks[i]), types.YChild{"TraceBlocks", allOptions.TraceBlocks[i]})
    }
    allOptions.EntityData.Leafs = types.NewOrderedMap()
    allOptions.EntityData.Leafs.Append("option", types.YLeaf{"Option", allOptions.Option})

    allOptions.EntityData.YListKeys = []string {"Option"}

    return &(allOptions.EntityData)
}

// SdrManager_SdrMgr_Trace_Location_AllOptions_TraceBlocks
type SdrManager_SdrMgr_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *SdrManager_SdrMgr_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks" + types.AddNoKeyToken(traceBlocks)
    traceBlocks.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-manager/sdr_mgr/trace/location/all-options/" + traceBlocks.EntityData.SegmentPath
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs = types.NewOrderedMap()
    traceBlocks.EntityData.Leafs.Append("data", types.YLeaf{"Data", traceBlocks.Data})

    traceBlocks.EntityData.YListKeys = []string {}

    return &(traceBlocks.EntityData)
}

// SdrOperation
type SdrOperation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SDR. The type is slice of SdrOperation_Sdr.
    Sdr []*SdrOperation_Sdr
}

func (sdrOperation *SdrOperation) GetEntityData() *types.CommonEntityData {
    sdrOperation.EntityData.YFilter = sdrOperation.YFilter
    sdrOperation.EntityData.YangName = "sdr-operation"
    sdrOperation.EntityData.BundleName = "cisco_ios_xr"
    sdrOperation.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdr-mgr"
    sdrOperation.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-operation"
    sdrOperation.EntityData.AbsolutePath = sdrOperation.EntityData.SegmentPath
    sdrOperation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrOperation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrOperation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrOperation.EntityData.Children = types.NewOrderedMap()
    sdrOperation.EntityData.Children.Append("sdr", types.YChild{"Sdr", nil})
    for i := range sdrOperation.Sdr {
        sdrOperation.EntityData.Children.Append(types.GetSegmentPath(sdrOperation.Sdr[i]), types.YChild{"Sdr", sdrOperation.Sdr[i]})
    }
    sdrOperation.EntityData.Leafs = types.NewOrderedMap()

    sdrOperation.EntityData.YListKeys = []string {}

    return &(sdrOperation.EntityData)
}

// SdrOperation_Sdr
// SDR
type SdrOperation_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the Secure Domain Router, 30 max
    // characters. The type is string with pattern: b'[a-zA-Z0-9_-]{1,30}'.
    Name interface{}

    
    Nodes SdrOperation_Sdr_Nodes
}

func (sdr *SdrOperation_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "sdr-operation"
    sdr.EntityData.SegmentPath = "sdr" + types.AddKeyToken(sdr.Name, "name")
    sdr.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-operation/" + sdr.EntityData.SegmentPath
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = types.NewOrderedMap()
    sdr.EntityData.Children.Append("nodes", types.YChild{"Nodes", &sdr.Nodes})
    sdr.EntityData.Leafs = types.NewOrderedMap()
    sdr.EntityData.Leafs.Append("name", types.YLeaf{"Name", sdr.Name})

    sdr.EntityData.YListKeys = []string {"Name"}

    return &(sdr.EntityData)
}

// SdrOperation_Sdr_Nodes
type SdrOperation_Sdr_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of SdrOperation_Sdr_Nodes_Location.
    Location []*SdrOperation_Sdr_Nodes_Location
}

func (nodes *SdrOperation_Sdr_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "sdr"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-operation/sdr/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("location", types.YChild{"Location", nil})
    for i := range nodes.Location {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Location[i]), types.YChild{"Location", nodes.Location[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// SdrOperation_Sdr_Nodes_Location
type SdrOperation_Sdr_Nodes_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    NodeLocation interface{}

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    NodeType interface{}

    // Progress of Operation. The type is string.
    Progress interface{}

    // State of Operation. The type is string.
    State interface{}
}

func (location *SdrOperation_Sdr_Nodes_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "nodes"
    location.EntityData.SegmentPath = "location" + types.AddKeyToken(location.NodeLocation, "node-location") + types.AddKeyToken(location.NodeType, "node-type")
    location.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:sdr-operation/sdr/nodes/" + location.EntityData.SegmentPath
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = types.NewOrderedMap()
    location.EntityData.Leafs = types.NewOrderedMap()
    location.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", location.NodeLocation})
    location.EntityData.Leafs.Append("node-type", types.YLeaf{"NodeType", location.NodeType})
    location.EntityData.Leafs.Append("progress", types.YLeaf{"Progress", location.Progress})
    location.EntityData.Leafs.Append("state", types.YLeaf{"State", location.State})

    location.EntityData.YListKeys = []string {"NodeLocation", "NodeType"}

    return &(location.EntityData)
}

// PrivateSdr
type PrivateSdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of PrivateSdr_SdrName.
    SdrName []*PrivateSdr_SdrName
}

func (privateSdr *PrivateSdr) GetEntityData() *types.CommonEntityData {
    privateSdr.EntityData.YFilter = privateSdr.YFilter
    privateSdr.EntityData.YangName = "private-sdr"
    privateSdr.EntityData.BundleName = "cisco_ios_xr"
    privateSdr.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-sdr-mgr"
    privateSdr.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-sdr-mgr:private-sdr"
    privateSdr.EntityData.AbsolutePath = privateSdr.EntityData.SegmentPath
    privateSdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateSdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateSdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateSdr.EntityData.Children = types.NewOrderedMap()
    privateSdr.EntityData.Children.Append("sdr-name", types.YChild{"SdrName", nil})
    for i := range privateSdr.SdrName {
        privateSdr.EntityData.Children.Append(types.GetSegmentPath(privateSdr.SdrName[i]), types.YChild{"SdrName", privateSdr.SdrName[i]})
    }
    privateSdr.EntityData.Leafs = types.NewOrderedMap()

    privateSdr.EntityData.YListKeys = []string {}

    return &(privateSdr.EntityData)
}

// PrivateSdr_SdrName
type PrivateSdr_SdrName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The type is interface{} with range: 0..4294967295.
    LeadRack0 interface{}

    // The type is interface{} with range: 0..4294967295.
    LeadRack1 interface{}

    // The type is slice of PrivateSdr_SdrName_Pairing.
    Pairing []*PrivateSdr_SdrName_Pairing
}

func (sdrName *PrivateSdr_SdrName) GetEntityData() *types.CommonEntityData {
    sdrName.EntityData.YFilter = sdrName.YFilter
    sdrName.EntityData.YangName = "sdr-name"
    sdrName.EntityData.BundleName = "cisco_ios_xr"
    sdrName.EntityData.ParentYangName = "private-sdr"
    sdrName.EntityData.SegmentPath = "sdr-name" + types.AddKeyToken(sdrName.Name, "name")
    sdrName.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:private-sdr/" + sdrName.EntityData.SegmentPath
    sdrName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrName.EntityData.Children = types.NewOrderedMap()
    sdrName.EntityData.Children.Append("pairing", types.YChild{"Pairing", nil})
    for i := range sdrName.Pairing {
        sdrName.EntityData.Children.Append(types.GetSegmentPath(sdrName.Pairing[i]), types.YChild{"Pairing", sdrName.Pairing[i]})
    }
    sdrName.EntityData.Leafs = types.NewOrderedMap()
    sdrName.EntityData.Leafs.Append("name", types.YLeaf{"Name", sdrName.Name})
    sdrName.EntityData.Leafs.Append("id", types.YLeaf{"Id", sdrName.Id})
    sdrName.EntityData.Leafs.Append("lead_rack0", types.YLeaf{"LeadRack0", sdrName.LeadRack0})
    sdrName.EntityData.Leafs.Append("lead_rack1", types.YLeaf{"LeadRack1", sdrName.LeadRack1})

    sdrName.EntityData.YListKeys = []string {"Name"}

    return &(sdrName.EntityData)
}

// PrivateSdr_SdrName_Pairing
type PrivateSdr_SdrName_Pairing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Num interface{}

    // The type is bool.
    SecondExist interface{}

    // The type is interface{} with range: 0..4294967295.
    Rp1Rack interface{}

    // The type is interface{} with range: 0..4294967295.
    Rp1Slot interface{}

    // The type is interface{} with range: 0..4294967295.
    Rp2Rack interface{}

    // The type is interface{} with range: 0..4294967295.
    Rp2Slot interface{}
}

func (pairing *PrivateSdr_SdrName_Pairing) GetEntityData() *types.CommonEntityData {
    pairing.EntityData.YFilter = pairing.YFilter
    pairing.EntityData.YangName = "pairing"
    pairing.EntityData.BundleName = "cisco_ios_xr"
    pairing.EntityData.ParentYangName = "sdr-name"
    pairing.EntityData.SegmentPath = "pairing" + types.AddKeyToken(pairing.Num, "num")
    pairing.EntityData.AbsolutePath = "Cisco-IOS-XR-sysadmin-sdr-mgr:private-sdr/sdr-name/" + pairing.EntityData.SegmentPath
    pairing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pairing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pairing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pairing.EntityData.Children = types.NewOrderedMap()
    pairing.EntityData.Leafs = types.NewOrderedMap()
    pairing.EntityData.Leafs.Append("num", types.YLeaf{"Num", pairing.Num})
    pairing.EntityData.Leafs.Append("second_exist", types.YLeaf{"SecondExist", pairing.SecondExist})
    pairing.EntityData.Leafs.Append("rp1_rack", types.YLeaf{"Rp1Rack", pairing.Rp1Rack})
    pairing.EntityData.Leafs.Append("rp1_slot", types.YLeaf{"Rp1Slot", pairing.Rp1Slot})
    pairing.EntityData.Leafs.Append("rp2_rack", types.YLeaf{"Rp2Rack", pairing.Rp2Rack})
    pairing.EntityData.Leafs.Append("rp2_slot", types.YLeaf{"Rp2Slot", pairing.Rp2Slot})

    pairing.EntityData.YListKeys = []string {"Num"}

    return &(pairing.EntityData)
}

