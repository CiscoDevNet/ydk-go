// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the top level container for
// all hardware devices managed in Sysadmin.
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_controllers_asr9k

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_controllers_asr9k"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-controllers-ASR9K controller}", reflect.TypeOf(Controller{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-controllers-asr9k:controller", reflect.TypeOf(Controller{}))
}

// Controller
type Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Switch_ Controller_Switch
}

func (controller *Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-controllers-asr9k"
    controller.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-controllers-asr9k:controller"
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = make(map[string]types.YChild)
    controller.EntityData.Children["switch"] = types.YChild{"Switch_", &controller.Switch_}
    controller.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(controller.EntityData)
}

// Controller_Switch
type Controller_Switch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Control Ethernet switch operational data.
    Oper Controller_Switch_Oper
}

func (self *Controller_Switch) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "switch"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "controller"
    self.EntityData.SegmentPath = "switch"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["oper"] = types.YChild{"Oper", &self.Oper}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(self.EntityData)
}

// Controller_Switch_Oper
// Control Ethernet switch operational data.
type Controller_Switch_Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Reachable Controller_Switch_Oper_Reachable

    
    Summary Controller_Switch_Oper_Summary

    
    Statistics Controller_Switch_Oper_Statistics

    
    Mac Controller_Switch_Oper_Mac

    
    Bridge Controller_Switch_Oper_Bridge

    
    Fdb Controller_Switch_Oper_Fdb

    
    Vlan Controller_Switch_Oper_Vlan

    
    Esd Controller_Switch_Oper_Esd

    
    MgmtAgent Controller_Switch_Oper_MgmtAgent

    
    Sdr Controller_Switch_Oper_Sdr

    
    PortState Controller_Switch_Oper_PortState

    
    Trunk Controller_Switch_Oper_Trunk

    
    SwitchDebugCont Controller_Switch_Oper_SwitchDebugCont
}

func (oper *Controller_Switch_Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "switch"
    oper.EntityData.SegmentPath = "oper"
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = make(map[string]types.YChild)
    oper.EntityData.Children["reachable"] = types.YChild{"Reachable", &oper.Reachable}
    oper.EntityData.Children["summary"] = types.YChild{"Summary", &oper.Summary}
    oper.EntityData.Children["statistics"] = types.YChild{"Statistics", &oper.Statistics}
    oper.EntityData.Children["mac"] = types.YChild{"Mac", &oper.Mac}
    oper.EntityData.Children["bridge"] = types.YChild{"Bridge", &oper.Bridge}
    oper.EntityData.Children["fdb"] = types.YChild{"Fdb", &oper.Fdb}
    oper.EntityData.Children["vlan"] = types.YChild{"Vlan", &oper.Vlan}
    oper.EntityData.Children["esd"] = types.YChild{"Esd", &oper.Esd}
    oper.EntityData.Children["mgmt-agent"] = types.YChild{"MgmtAgent", &oper.MgmtAgent}
    oper.EntityData.Children["sdr"] = types.YChild{"Sdr", &oper.Sdr}
    oper.EntityData.Children["port-state"] = types.YChild{"PortState", &oper.PortState}
    oper.EntityData.Children["trunk"] = types.YChild{"Trunk", &oper.Trunk}
    oper.EntityData.Children["switch-debug-cont"] = types.YChild{"SwitchDebugCont", &oper.SwitchDebugCont}
    oper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(oper.EntityData)
}

// Controller_Switch_Oper_Reachable
type Controller_Switch_Oper_Reachable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Reachable_Location.
    Location []Controller_Switch_Oper_Reachable_Location
}

func (reachable *Controller_Switch_Oper_Reachable) GetEntityData() *types.CommonEntityData {
    reachable.EntityData.YFilter = reachable.YFilter
    reachable.EntityData.YangName = "reachable"
    reachable.EntityData.BundleName = "cisco_ios_xr"
    reachable.EntityData.ParentYangName = "oper"
    reachable.EntityData.SegmentPath = "reachable"
    reachable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reachable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reachable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reachable.EntityData.Children = make(map[string]types.YChild)
    reachable.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range reachable.Location {
        reachable.EntityData.Children[types.GetSegmentPath(&reachable.Location[i])] = types.YChild{"Location", &reachable.Location[i]}
    }
    reachable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(reachable.EntityData)
}

// Controller_Switch_Oper_Reachable_Location
type Controller_Switch_Oper_Reachable_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}
}

func (location *Controller_Switch_Oper_Reachable_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "reachable"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Summary
type Controller_Switch_Oper_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Summary_Location.
    Location []Controller_Switch_Oper_Summary_Location
}

func (summary *Controller_Switch_Oper_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "oper"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range summary.Location {
        summary.EntityData.Children[types.GetSegmentPath(&summary.Location[i])] = types.YChild{"Location", &summary.Location[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// Controller_Switch_Oper_Summary_Location
type Controller_Switch_Oper_Summary_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Rack serial number. The type is string.
    SerialNum interface{}

    // The type is slice of Controller_Switch_Oper_Summary_Location_PortIter.
    PortIter []Controller_Switch_Oper_Summary_Location_PortIter
}

func (location *Controller_Switch_Oper_Summary_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["serial-num"] = types.YLeaf{"SerialNum", location.SerialNum}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Summary_Location_PortIter
type Controller_Switch_Oper_Summary_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // Physical port state. The type is EsdmaSwitchPortState.
    PhysState interface{}

    // Administrative port state. The type is EsdmaSwitchPortState.
    AdminState interface{}

    // Indicates the port speed in bits per second. The type is string.
    PortSpeed interface{}

    // Protocol invoked port state. The type is MlapStateEnum.
    ProtocolState interface{}

    // Indicates whether this port is allowed to forward traffic. The type is
    // SwitchForwardingState.
    Forwarding interface{}

    // Indicates what this port connects to. The type is string.
    ConnectsTo interface{}
}

func (portIter *Controller_Switch_Oper_Summary_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    portIter.EntityData.Leafs["phys-state"] = types.YLeaf{"PhysState", portIter.PhysState}
    portIter.EntityData.Leafs["admin-state"] = types.YLeaf{"AdminState", portIter.AdminState}
    portIter.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", portIter.PortSpeed}
    portIter.EntityData.Leafs["protocol-state"] = types.YLeaf{"ProtocolState", portIter.ProtocolState}
    portIter.EntityData.Leafs["forwarding"] = types.YLeaf{"Forwarding", portIter.Forwarding}
    portIter.EntityData.Leafs["connects-to"] = types.YLeaf{"ConnectsTo", portIter.ConnectsTo}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Statistics
type Controller_Switch_Oper_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    SummaryStatistics Controller_Switch_Oper_Statistics_SummaryStatistics

    
    Detail Controller_Switch_Oper_Statistics_Detail
}

func (statistics *Controller_Switch_Oper_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "oper"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["summary-statistics"] = types.YChild{"SummaryStatistics", &statistics.SummaryStatistics}
    statistics.EntityData.Children["detail"] = types.YChild{"Detail", &statistics.Detail}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Controller_Switch_Oper_Statistics_SummaryStatistics
type Controller_Switch_Oper_Statistics_SummaryStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Statistics_SummaryStatistics_Location.
    Location []Controller_Switch_Oper_Statistics_SummaryStatistics_Location
}

func (summaryStatistics *Controller_Switch_Oper_Statistics_SummaryStatistics) GetEntityData() *types.CommonEntityData {
    summaryStatistics.EntityData.YFilter = summaryStatistics.YFilter
    summaryStatistics.EntityData.YangName = "summary-statistics"
    summaryStatistics.EntityData.BundleName = "cisco_ios_xr"
    summaryStatistics.EntityData.ParentYangName = "statistics"
    summaryStatistics.EntityData.SegmentPath = "summary-statistics"
    summaryStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryStatistics.EntityData.Children = make(map[string]types.YChild)
    summaryStatistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range summaryStatistics.Location {
        summaryStatistics.EntityData.Children[types.GetSegmentPath(&summaryStatistics.Location[i])] = types.YChild{"Location", &summaryStatistics.Location[i]}
    }
    summaryStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summaryStatistics.EntityData)
}

// Controller_Switch_Oper_Statistics_SummaryStatistics_Location
type Controller_Switch_Oper_Statistics_SummaryStatistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Rack serial number. The type is string.
    SerialNum interface{}

    // The type is slice of
    // Controller_Switch_Oper_Statistics_SummaryStatistics_Location_PortIter.
    PortIter []Controller_Switch_Oper_Statistics_SummaryStatistics_Location_PortIter
}

func (location *Controller_Switch_Oper_Statistics_SummaryStatistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "summary-statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["serial-num"] = types.YLeaf{"SerialNum", location.SerialNum}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Statistics_SummaryStatistics_Location_PortIter
type Controller_Switch_Oper_Statistics_SummaryStatistics_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // Physical port state. The type is EsdmaSwitchPortState.
    PhysState interface{}

    // Physical port state changes. The type is interface{} with range:
    // 0..4294967295.
    StateChanges interface{}

    // Packets transmitted on this switch port. The type is interface{} with
    // range: 0..18446744073709551615.
    SwSumTxPackets interface{}

    // Indicates the port speed in bits per second. The type is interface{} with
    // range: 0..18446744073709551615.
    SwSumRxPackets interface{}

    // Indicates the number of transmitted packets that had an error or were
    // dropped by the policer. The type is interface{} with range:
    // 0..18446744073709551615.
    SwSumTxDropsErrors interface{}

    // Indicates the number of received packets that had an error or were dropped
    // by the policer. The type is interface{} with range:
    // 0..18446744073709551615.
    SwSumRxDropsErrors interface{}

    // Indicates what this port connects to. The type is string.
    ConnectsTo interface{}
}

func (portIter *Controller_Switch_Oper_Statistics_SummaryStatistics_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    portIter.EntityData.Leafs["phys-state"] = types.YLeaf{"PhysState", portIter.PhysState}
    portIter.EntityData.Leafs["state-changes"] = types.YLeaf{"StateChanges", portIter.StateChanges}
    portIter.EntityData.Leafs["sw-sum-tx-packets"] = types.YLeaf{"SwSumTxPackets", portIter.SwSumTxPackets}
    portIter.EntityData.Leafs["sw-sum-rx-packets"] = types.YLeaf{"SwSumRxPackets", portIter.SwSumRxPackets}
    portIter.EntityData.Leafs["sw-sum-tx-drops-errors"] = types.YLeaf{"SwSumTxDropsErrors", portIter.SwSumTxDropsErrors}
    portIter.EntityData.Leafs["sw-sum-rx-drops-errors"] = types.YLeaf{"SwSumRxDropsErrors", portIter.SwSumRxDropsErrors}
    portIter.EntityData.Leafs["connects-to"] = types.YLeaf{"ConnectsTo", portIter.ConnectsTo}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Statistics_Detail
type Controller_Switch_Oper_Statistics_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Statistics_Detail_Location.
    Location []Controller_Switch_Oper_Statistics_Detail_Location
}

func (detail *Controller_Switch_Oper_Statistics_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "statistics"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range detail.Location {
        detail.EntityData.Children[types.GetSegmentPath(&detail.Location[i])] = types.YChild{"Location", &detail.Location[i]}
    }
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// Controller_Switch_Oper_Statistics_Detail_Location
type Controller_Switch_Oper_Statistics_Detail_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Statistics_Detail_Location_PortIter.
    PortIter []Controller_Switch_Oper_Statistics_Detail_Location_PortIter
}

func (location *Controller_Switch_Oper_Statistics_Detail_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "detail"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Statistics_Detail_Location_PortIter
type Controller_Switch_Oper_Statistics_Detail_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // Physical port state. The type is EsdmaSwitchPortState.
    PhysState interface{}

    // Indicates the port speed in bits per second. The type is string.
    PortSpeed interface{}

    // Indicates what this port connects to. The type is string.
    ConnectsTo interface{}

    
    Counters Controller_Switch_Oper_Statistics_Detail_Location_PortIter_Counters
}

func (portIter *Controller_Switch_Oper_Statistics_Detail_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["counters"] = types.YChild{"Counters", &portIter.Counters}
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    portIter.EntityData.Leafs["phys-state"] = types.YLeaf{"PhysState", portIter.PhysState}
    portIter.EntityData.Leafs["port-speed"] = types.YLeaf{"PortSpeed", portIter.PortSpeed}
    portIter.EntityData.Leafs["connects-to"] = types.YLeaf{"ConnectsTo", portIter.ConnectsTo}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Statistics_Detail_Location_PortIter_Counters
type Controller_Switch_Oper_Statistics_Detail_Location_PortIter_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxUcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxMcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxBcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxFlowControl interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxGoodOctets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxBadOctets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxFifoOverrun interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxUndersize interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxFragments interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxOversize interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxJabber interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxBadCrc interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxCollisions interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxPolicingDrops interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxUcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxMcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxBcastPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxFlowControl interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxGoodOctets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxDeferred interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxFifoUnrun interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxMultCollision interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxExcessCollision interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxLateCollisions interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxPolicingDrops interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetTxqDrops interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets64 interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets65127 interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets128255 interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets256511 interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets5121023 interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    SwDetRxtxPackets1024Max interface{}
}

func (counters *Controller_Switch_Oper_Statistics_Detail_Location_PortIter_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "port-iter"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    counters.EntityData.Leafs["sw-det-rx-ucast-packets"] = types.YLeaf{"SwDetRxUcastPackets", counters.SwDetRxUcastPackets}
    counters.EntityData.Leafs["sw-det-rx-mcast-packets"] = types.YLeaf{"SwDetRxMcastPackets", counters.SwDetRxMcastPackets}
    counters.EntityData.Leafs["sw-det-rx-bcast-packets"] = types.YLeaf{"SwDetRxBcastPackets", counters.SwDetRxBcastPackets}
    counters.EntityData.Leafs["sw-det-rx-flow-control"] = types.YLeaf{"SwDetRxFlowControl", counters.SwDetRxFlowControl}
    counters.EntityData.Leafs["sw-det-rx-good-octets"] = types.YLeaf{"SwDetRxGoodOctets", counters.SwDetRxGoodOctets}
    counters.EntityData.Leafs["sw-det-rx-bad-octets"] = types.YLeaf{"SwDetRxBadOctets", counters.SwDetRxBadOctets}
    counters.EntityData.Leafs["sw-det-rx-fifo-overrun"] = types.YLeaf{"SwDetRxFifoOverrun", counters.SwDetRxFifoOverrun}
    counters.EntityData.Leafs["sw-det-rx-undersize"] = types.YLeaf{"SwDetRxUndersize", counters.SwDetRxUndersize}
    counters.EntityData.Leafs["sw-det-rx-fragments"] = types.YLeaf{"SwDetRxFragments", counters.SwDetRxFragments}
    counters.EntityData.Leafs["sw-det-rx-oversize"] = types.YLeaf{"SwDetRxOversize", counters.SwDetRxOversize}
    counters.EntityData.Leafs["sw-det-rx-jabber"] = types.YLeaf{"SwDetRxJabber", counters.SwDetRxJabber}
    counters.EntityData.Leafs["sw-det-rx-errors"] = types.YLeaf{"SwDetRxErrors", counters.SwDetRxErrors}
    counters.EntityData.Leafs["sw-det-rx-bad-crc"] = types.YLeaf{"SwDetRxBadCrc", counters.SwDetRxBadCrc}
    counters.EntityData.Leafs["sw-det-rx-collisions"] = types.YLeaf{"SwDetRxCollisions", counters.SwDetRxCollisions}
    counters.EntityData.Leafs["sw-det-rx-policing-drops"] = types.YLeaf{"SwDetRxPolicingDrops", counters.SwDetRxPolicingDrops}
    counters.EntityData.Leafs["sw-det-tx-ucast-packets"] = types.YLeaf{"SwDetTxUcastPackets", counters.SwDetTxUcastPackets}
    counters.EntityData.Leafs["sw-det-tx-mcast-packets"] = types.YLeaf{"SwDetTxMcastPackets", counters.SwDetTxMcastPackets}
    counters.EntityData.Leafs["sw-det-tx-bcast-packets"] = types.YLeaf{"SwDetTxBcastPackets", counters.SwDetTxBcastPackets}
    counters.EntityData.Leafs["sw-det-tx-flow-control"] = types.YLeaf{"SwDetTxFlowControl", counters.SwDetTxFlowControl}
    counters.EntityData.Leafs["sw-det-tx-good-octets"] = types.YLeaf{"SwDetTxGoodOctets", counters.SwDetTxGoodOctets}
    counters.EntityData.Leafs["sw-det-tx-deferred"] = types.YLeaf{"SwDetTxDeferred", counters.SwDetTxDeferred}
    counters.EntityData.Leafs["sw-det-tx-fifo-unrun"] = types.YLeaf{"SwDetTxFifoUnrun", counters.SwDetTxFifoUnrun}
    counters.EntityData.Leafs["sw-det-tx-mult-collision"] = types.YLeaf{"SwDetTxMultCollision", counters.SwDetTxMultCollision}
    counters.EntityData.Leafs["sw-det-tx-excess-collision"] = types.YLeaf{"SwDetTxExcessCollision", counters.SwDetTxExcessCollision}
    counters.EntityData.Leafs["sw-det-tx-late-collisions"] = types.YLeaf{"SwDetTxLateCollisions", counters.SwDetTxLateCollisions}
    counters.EntityData.Leafs["sw-det-tx-policing-drops"] = types.YLeaf{"SwDetTxPolicingDrops", counters.SwDetTxPolicingDrops}
    counters.EntityData.Leafs["sw-det-txq-drops"] = types.YLeaf{"SwDetTxqDrops", counters.SwDetTxqDrops}
    counters.EntityData.Leafs["sw-det-rxtx-packets-64"] = types.YLeaf{"SwDetRxtxPackets64", counters.SwDetRxtxPackets64}
    counters.EntityData.Leafs["sw-det-rxtx-packets-65-127"] = types.YLeaf{"SwDetRxtxPackets65127", counters.SwDetRxtxPackets65127}
    counters.EntityData.Leafs["sw-det-rxtx-packets-128-255"] = types.YLeaf{"SwDetRxtxPackets128255", counters.SwDetRxtxPackets128255}
    counters.EntityData.Leafs["sw-det-rxtx-packets-256-511"] = types.YLeaf{"SwDetRxtxPackets256511", counters.SwDetRxtxPackets256511}
    counters.EntityData.Leafs["sw-det-rxtx-packets-512-1023"] = types.YLeaf{"SwDetRxtxPackets5121023", counters.SwDetRxtxPackets5121023}
    counters.EntityData.Leafs["sw-det-rxtx-packets-1024-max"] = types.YLeaf{"SwDetRxtxPackets1024Max", counters.SwDetRxtxPackets1024Max}
    return &(counters.EntityData)
}

// Controller_Switch_Oper_Mac
type Controller_Switch_Oper_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    MacStatistics Controller_Switch_Oper_Mac_MacStatistics
}

func (mac *Controller_Switch_Oper_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "oper"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["mac-statistics"] = types.YChild{"MacStatistics", &mac.MacStatistics}
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mac.EntityData)
}

// Controller_Switch_Oper_Mac_MacStatistics
type Controller_Switch_Oper_Mac_MacStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Mac_MacStatistics_Location.
    Location []Controller_Switch_Oper_Mac_MacStatistics_Location
}

func (macStatistics *Controller_Switch_Oper_Mac_MacStatistics) GetEntityData() *types.CommonEntityData {
    macStatistics.EntityData.YFilter = macStatistics.YFilter
    macStatistics.EntityData.YangName = "mac-statistics"
    macStatistics.EntityData.BundleName = "cisco_ios_xr"
    macStatistics.EntityData.ParentYangName = "mac"
    macStatistics.EntityData.SegmentPath = "mac-statistics"
    macStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macStatistics.EntityData.Children = make(map[string]types.YChild)
    macStatistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range macStatistics.Location {
        macStatistics.EntityData.Children[types.GetSegmentPath(&macStatistics.Location[i])] = types.YChild{"Location", &macStatistics.Location[i]}
    }
    macStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macStatistics.EntityData)
}

// Controller_Switch_Oper_Mac_MacStatistics_Location
type Controller_Switch_Oper_Mac_MacStatistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch MAC information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter.
    PortIter []Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter
}

func (location *Controller_Switch_Oper_Mac_MacStatistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "mac-statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter
type Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // The type is slice of
    // Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter_MacEntry.
    MacEntry []Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter_MacEntry
}

func (portIter *Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["mac-entry"] = types.YChild{"MacEntry", nil}
    for i := range portIter.MacEntry {
        portIter.EntityData.Children[types.GetSegmentPath(&portIter.MacEntry[i])] = types.YChild{"MacEntry", &portIter.MacEntry[i]}
    }
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter_MacEntry
type Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter_MacEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // The type is string.
    BaseReg interface{}

    // The type is string.
    Desc interface{}

    // The type is string.
    Value interface{}
}

func (macEntry *Controller_Switch_Oper_Mac_MacStatistics_Location_PortIter_MacEntry) GetEntityData() *types.CommonEntityData {
    macEntry.EntityData.YFilter = macEntry.YFilter
    macEntry.EntityData.YangName = "mac-entry"
    macEntry.EntityData.BundleName = "cisco_ios_xr"
    macEntry.EntityData.ParentYangName = "port-iter"
    macEntry.EntityData.SegmentPath = "mac-entry" + "[id='" + fmt.Sprintf("%v", macEntry.Id) + "']"
    macEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macEntry.EntityData.Children = make(map[string]types.YChild)
    macEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    macEntry.EntityData.Leafs["id"] = types.YLeaf{"Id", macEntry.Id}
    macEntry.EntityData.Leafs["base-reg"] = types.YLeaf{"BaseReg", macEntry.BaseReg}
    macEntry.EntityData.Leafs["desc"] = types.YLeaf{"Desc", macEntry.Desc}
    macEntry.EntityData.Leafs["value"] = types.YLeaf{"Value", macEntry.Value}
    return &(macEntry.EntityData)
}

// Controller_Switch_Oper_Bridge
type Controller_Switch_Oper_Bridge struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Statistics Controller_Switch_Oper_Bridge_Statistics
}

func (bridge *Controller_Switch_Oper_Bridge) GetEntityData() *types.CommonEntityData {
    bridge.EntityData.YFilter = bridge.YFilter
    bridge.EntityData.YangName = "bridge"
    bridge.EntityData.BundleName = "cisco_ios_xr"
    bridge.EntityData.ParentYangName = "oper"
    bridge.EntityData.SegmentPath = "bridge"
    bridge.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridge.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridge.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridge.EntityData.Children = make(map[string]types.YChild)
    bridge.EntityData.Children["statistics"] = types.YChild{"Statistics", &bridge.Statistics}
    bridge.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bridge.EntityData)
}

// Controller_Switch_Oper_Bridge_Statistics
type Controller_Switch_Oper_Bridge_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Bridge_Statistics_Location.
    Location []Controller_Switch_Oper_Bridge_Statistics_Location
}

func (statistics *Controller_Switch_Oper_Bridge_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "bridge"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range statistics.Location {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Location[i])] = types.YChild{"Location", &statistics.Location[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Controller_Switch_Oper_Bridge_Statistics_Location
type Controller_Switch_Oper_Bridge_Statistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Bridge_Statistics_Location_IngressSetId.
    IngressSetId []Controller_Switch_Oper_Bridge_Statistics_Location_IngressSetId

    // The type is slice of
    // Controller_Switch_Oper_Bridge_Statistics_Location_EgressSetId.
    EgressSetId []Controller_Switch_Oper_Bridge_Statistics_Location_EgressSetId
}

func (location *Controller_Switch_Oper_Bridge_Statistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["ingress-set-id"] = types.YChild{"IngressSetId", nil}
    for i := range location.IngressSetId {
        location.EntityData.Children[types.GetSegmentPath(&location.IngressSetId[i])] = types.YChild{"IngressSetId", &location.IngressSetId[i]}
    }
    location.EntityData.Children["egress-set-id"] = types.YChild{"EgressSetId", nil}
    for i := range location.EgressSetId {
        location.EntityData.Children[types.GetSegmentPath(&location.EgressSetId[i])] = types.YChild{"EgressSetId", &location.EgressSetId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Bridge_Statistics_Location_IngressSetId
type Controller_Switch_Oper_Bridge_Statistics_Location_IngressSetId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    IngressSet interface{}

    // The type is string.
    IngressSetName interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressFrames interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressVlanDiscards interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressSecurityDiscards interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    IngressOtherDiscards interface{}
}

func (ingressSetId *Controller_Switch_Oper_Bridge_Statistics_Location_IngressSetId) GetEntityData() *types.CommonEntityData {
    ingressSetId.EntityData.YFilter = ingressSetId.YFilter
    ingressSetId.EntityData.YangName = "ingress-set-id"
    ingressSetId.EntityData.BundleName = "cisco_ios_xr"
    ingressSetId.EntityData.ParentYangName = "location"
    ingressSetId.EntityData.SegmentPath = "ingress-set-id" + "[ingress-set='" + fmt.Sprintf("%v", ingressSetId.IngressSet) + "']"
    ingressSetId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ingressSetId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ingressSetId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ingressSetId.EntityData.Children = make(map[string]types.YChild)
    ingressSetId.EntityData.Leafs = make(map[string]types.YLeaf)
    ingressSetId.EntityData.Leafs["ingress-set"] = types.YLeaf{"IngressSet", ingressSetId.IngressSet}
    ingressSetId.EntityData.Leafs["ingress-set-name"] = types.YLeaf{"IngressSetName", ingressSetId.IngressSetName}
    ingressSetId.EntityData.Leafs["ingress-frames"] = types.YLeaf{"IngressFrames", ingressSetId.IngressFrames}
    ingressSetId.EntityData.Leafs["ingress-vlan-discards"] = types.YLeaf{"IngressVlanDiscards", ingressSetId.IngressVlanDiscards}
    ingressSetId.EntityData.Leafs["ingress-security-discards"] = types.YLeaf{"IngressSecurityDiscards", ingressSetId.IngressSecurityDiscards}
    ingressSetId.EntityData.Leafs["ingress-other-discards"] = types.YLeaf{"IngressOtherDiscards", ingressSetId.IngressOtherDiscards}
    return &(ingressSetId.EntityData)
}

// Controller_Switch_Oper_Bridge_Statistics_Location_EgressSetId
type Controller_Switch_Oper_Bridge_Statistics_Location_EgressSetId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    EgressSet interface{}

    // The type is string.
    EgressSetName interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressUcastFrames interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressMcastFrames interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressBcastFrames interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressDiscardedFrames interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressTxqCongestion interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressCtrlPackets interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    EgressOtherDrops interface{}
}

func (egressSetId *Controller_Switch_Oper_Bridge_Statistics_Location_EgressSetId) GetEntityData() *types.CommonEntityData {
    egressSetId.EntityData.YFilter = egressSetId.YFilter
    egressSetId.EntityData.YangName = "egress-set-id"
    egressSetId.EntityData.BundleName = "cisco_ios_xr"
    egressSetId.EntityData.ParentYangName = "location"
    egressSetId.EntityData.SegmentPath = "egress-set-id" + "[egress-set='" + fmt.Sprintf("%v", egressSetId.EgressSet) + "']"
    egressSetId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    egressSetId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    egressSetId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    egressSetId.EntityData.Children = make(map[string]types.YChild)
    egressSetId.EntityData.Leafs = make(map[string]types.YLeaf)
    egressSetId.EntityData.Leafs["egress-set"] = types.YLeaf{"EgressSet", egressSetId.EgressSet}
    egressSetId.EntityData.Leafs["egress-set-name"] = types.YLeaf{"EgressSetName", egressSetId.EgressSetName}
    egressSetId.EntityData.Leafs["egress-ucast-frames"] = types.YLeaf{"EgressUcastFrames", egressSetId.EgressUcastFrames}
    egressSetId.EntityData.Leafs["egress-mcast-frames"] = types.YLeaf{"EgressMcastFrames", egressSetId.EgressMcastFrames}
    egressSetId.EntityData.Leafs["egress-bcast-frames"] = types.YLeaf{"EgressBcastFrames", egressSetId.EgressBcastFrames}
    egressSetId.EntityData.Leafs["egress-discarded-frames"] = types.YLeaf{"EgressDiscardedFrames", egressSetId.EgressDiscardedFrames}
    egressSetId.EntityData.Leafs["egress-txq-congestion"] = types.YLeaf{"EgressTxqCongestion", egressSetId.EgressTxqCongestion}
    egressSetId.EntityData.Leafs["egress-ctrl-packets"] = types.YLeaf{"EgressCtrlPackets", egressSetId.EgressCtrlPackets}
    egressSetId.EntityData.Leafs["egress-other-drops"] = types.YLeaf{"EgressOtherDrops", egressSetId.EgressOtherDrops}
    return &(egressSetId.EntityData)
}

// Controller_Switch_Oper_Fdb
type Controller_Switch_Oper_Fdb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Vlan Controller_Switch_Oper_Fdb_Vlan

    
    Mac Controller_Switch_Oper_Fdb_Mac

    
    Port Controller_Switch_Oper_Fdb_Port

    
    Statistics Controller_Switch_Oper_Fdb_Statistics

    
    SwitchFdbCommon Controller_Switch_Oper_Fdb_SwitchFdbCommon
}

func (fdb *Controller_Switch_Oper_Fdb) GetEntityData() *types.CommonEntityData {
    fdb.EntityData.YFilter = fdb.YFilter
    fdb.EntityData.YangName = "fdb"
    fdb.EntityData.BundleName = "cisco_ios_xr"
    fdb.EntityData.ParentYangName = "oper"
    fdb.EntityData.SegmentPath = "fdb"
    fdb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdb.EntityData.Children = make(map[string]types.YChild)
    fdb.EntityData.Children["vlan"] = types.YChild{"Vlan", &fdb.Vlan}
    fdb.EntityData.Children["mac"] = types.YChild{"Mac", &fdb.Mac}
    fdb.EntityData.Children["port"] = types.YChild{"Port", &fdb.Port}
    fdb.EntityData.Children["statistics"] = types.YChild{"Statistics", &fdb.Statistics}
    fdb.EntityData.Children["switch-fdb-common"] = types.YChild{"SwitchFdbCommon", &fdb.SwitchFdbCommon}
    fdb.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fdb.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan
type Controller_Switch_Oper_Fdb_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Fdb_Vlan_VlanIter.
    VlanIter []Controller_Switch_Oper_Fdb_Vlan_VlanIter
}

func (vlan *Controller_Switch_Oper_Fdb_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "fdb"
    vlan.EntityData.SegmentPath = "vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = make(map[string]types.YChild)
    vlan.EntityData.Children["vlan-iter"] = types.YChild{"VlanIter", nil}
    for i := range vlan.VlanIter {
        vlan.EntityData.Children[types.GetSegmentPath(&vlan.VlanIter[i])] = types.YChild{"VlanIter", &vlan.VlanIter[i]}
    }
    vlan.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlan.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan_VlanIter
type Controller_Switch_Oper_Fdb_Vlan_VlanIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4095.
    Vlan interface{}

    
    SwitchFdbCommon Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon
}

func (vlanIter *Controller_Switch_Oper_Fdb_Vlan_VlanIter) GetEntityData() *types.CommonEntityData {
    vlanIter.EntityData.YFilter = vlanIter.YFilter
    vlanIter.EntityData.YangName = "vlan-iter"
    vlanIter.EntityData.BundleName = "cisco_ios_xr"
    vlanIter.EntityData.ParentYangName = "vlan"
    vlanIter.EntityData.SegmentPath = "vlan-iter" + "[vlan='" + fmt.Sprintf("%v", vlanIter.Vlan) + "']"
    vlanIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanIter.EntityData.Children = make(map[string]types.YChild)
    vlanIter.EntityData.Children["switch-fdb-common"] = types.YChild{"SwitchFdbCommon", &vlanIter.SwitchFdbCommon}
    vlanIter.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanIter.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanIter.Vlan}
    return &(vlanIter.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon
type Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location.
    Location []Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location
}

func (switchFdbCommon *Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon) GetEntityData() *types.CommonEntityData {
    switchFdbCommon.EntityData.YFilter = switchFdbCommon.YFilter
    switchFdbCommon.EntityData.YangName = "switch-fdb-common"
    switchFdbCommon.EntityData.BundleName = "cisco_ios_xr"
    switchFdbCommon.EntityData.ParentYangName = "vlan-iter"
    switchFdbCommon.EntityData.SegmentPath = "switch-fdb-common"
    switchFdbCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchFdbCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchFdbCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchFdbCommon.EntityData.Children = make(map[string]types.YChild)
    switchFdbCommon.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range switchFdbCommon.Location {
        switchFdbCommon.EntityData.Children[types.GetSegmentPath(&switchFdbCommon.Location[i])] = types.YChild{"Location", &switchFdbCommon.Location[i]}
    }
    switchFdbCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(switchFdbCommon.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location
type Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch FDB information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Number of FDB entries in the table. The type is interface{} with range:
    // 0..4294967295.
    NumEntries interface{}

    // FDB entries contain an entry from the trunk. The type is interface{} with
    // range: 0..255.
    HasTrunkEntry interface{}

    // Message displayed when an FDB entry contains an entry for a trunk member
    // port. The type is string.
    TrunkEntryMessage interface{}

    
    FdbBlock Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock
}

func (location *Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "switch-fdb-common"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fdb-block"] = types.YChild{"FdbBlock", &location.FdbBlock}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["num-entries"] = types.YLeaf{"NumEntries", location.NumEntries}
    location.EntityData.Leafs["has-trunk-entry"] = types.YLeaf{"HasTrunkEntry", location.HasTrunkEntry}
    location.EntityData.Leafs["trunk-entry-message"] = types.YLeaf{"TrunkEntryMessage", location.TrunkEntryMessage}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock
type Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry.
    FdbEntry []Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
}

func (fdbBlock *Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock) GetEntityData() *types.CommonEntityData {
    fdbBlock.EntityData.YFilter = fdbBlock.YFilter
    fdbBlock.EntityData.YangName = "fdb-block"
    fdbBlock.EntityData.BundleName = "cisco_ios_xr"
    fdbBlock.EntityData.ParentYangName = "location"
    fdbBlock.EntityData.SegmentPath = "fdb-block"
    fdbBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbBlock.EntityData.Children = make(map[string]types.YChild)
    fdbBlock.EntityData.Children["fdb-entry"] = types.YChild{"FdbEntry", nil}
    for i := range fdbBlock.FdbEntry {
        fdbBlock.EntityData.Children[types.GetSegmentPath(&fdbBlock.FdbEntry[i])] = types.YChild{"FdbEntry", &fdbBlock.FdbEntry[i]}
    }
    fdbBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fdbBlock.EntityData)
}

// Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
type Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    FdbIndex interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    FdbMacAddr interface{}

    // The type is interface{} with range: 0..4095.
    FdbVlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry_FdbVlan
    FdbVlanHex interface{}

    // Switch port MAC address learned on. The type is interface{} with range:
    // 0..127.
    FdbPort interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbTrapEntry interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbStaticEntry interface{}

    // The type is slice of interface{} with range: 0..255.
    FdbSyncedCores []interface{}
}

func (fdbEntry *Controller_Switch_Oper_Fdb_Vlan_VlanIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry) GetEntityData() *types.CommonEntityData {
    fdbEntry.EntityData.YFilter = fdbEntry.YFilter
    fdbEntry.EntityData.YangName = "fdb-entry"
    fdbEntry.EntityData.BundleName = "cisco_ios_xr"
    fdbEntry.EntityData.ParentYangName = "fdb-block"
    fdbEntry.EntityData.SegmentPath = "fdb-entry" + "[fdb-index='" + fmt.Sprintf("%v", fdbEntry.FdbIndex) + "']"
    fdbEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbEntry.EntityData.Children = make(map[string]types.YChild)
    fdbEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    fdbEntry.EntityData.Leafs["fdb-index"] = types.YLeaf{"FdbIndex", fdbEntry.FdbIndex}
    fdbEntry.EntityData.Leafs["fdb-mac-addr"] = types.YLeaf{"FdbMacAddr", fdbEntry.FdbMacAddr}
    fdbEntry.EntityData.Leafs["fdb-vlan"] = types.YLeaf{"FdbVlan", fdbEntry.FdbVlan}
    fdbEntry.EntityData.Leafs["fdb-vlan-hex"] = types.YLeaf{"FdbVlanHex", fdbEntry.FdbVlanHex}
    fdbEntry.EntityData.Leafs["fdb-port"] = types.YLeaf{"FdbPort", fdbEntry.FdbPort}
    fdbEntry.EntityData.Leafs["fdb-trap-entry"] = types.YLeaf{"FdbTrapEntry", fdbEntry.FdbTrapEntry}
    fdbEntry.EntityData.Leafs["fdb-static-entry"] = types.YLeaf{"FdbStaticEntry", fdbEntry.FdbStaticEntry}
    fdbEntry.EntityData.Leafs["fdb-synced-cores"] = types.YLeaf{"FdbSyncedCores", fdbEntry.FdbSyncedCores}
    return &(fdbEntry.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac
type Controller_Switch_Oper_Fdb_Mac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Fdb_Mac_MacIter.
    MacIter []Controller_Switch_Oper_Fdb_Mac_MacIter
}

func (mac *Controller_Switch_Oper_Fdb_Mac) GetEntityData() *types.CommonEntityData {
    mac.EntityData.YFilter = mac.YFilter
    mac.EntityData.YangName = "mac"
    mac.EntityData.BundleName = "cisco_ios_xr"
    mac.EntityData.ParentYangName = "fdb"
    mac.EntityData.SegmentPath = "mac"
    mac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mac.EntityData.Children = make(map[string]types.YChild)
    mac.EntityData.Children["mac-iter"] = types.YChild{"MacIter", nil}
    for i := range mac.MacIter {
        mac.EntityData.Children[types.GetSegmentPath(&mac.MacIter[i])] = types.YChild{"MacIter", &mac.MacIter[i]}
    }
    mac.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mac.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac_MacIter
type Controller_Switch_Oper_Fdb_Mac_MacIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    Mac interface{}

    
    SwitchFdbCommon Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon
}

func (macIter *Controller_Switch_Oper_Fdb_Mac_MacIter) GetEntityData() *types.CommonEntityData {
    macIter.EntityData.YFilter = macIter.YFilter
    macIter.EntityData.YangName = "mac-iter"
    macIter.EntityData.BundleName = "cisco_ios_xr"
    macIter.EntityData.ParentYangName = "mac"
    macIter.EntityData.SegmentPath = "mac-iter" + "[mac='" + fmt.Sprintf("%v", macIter.Mac) + "']"
    macIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macIter.EntityData.Children = make(map[string]types.YChild)
    macIter.EntityData.Children["switch-fdb-common"] = types.YChild{"SwitchFdbCommon", &macIter.SwitchFdbCommon}
    macIter.EntityData.Leafs = make(map[string]types.YLeaf)
    macIter.EntityData.Leafs["mac"] = types.YLeaf{"Mac", macIter.Mac}
    return &(macIter.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon
type Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location.
    Location []Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location
}

func (switchFdbCommon *Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon) GetEntityData() *types.CommonEntityData {
    switchFdbCommon.EntityData.YFilter = switchFdbCommon.YFilter
    switchFdbCommon.EntityData.YangName = "switch-fdb-common"
    switchFdbCommon.EntityData.BundleName = "cisco_ios_xr"
    switchFdbCommon.EntityData.ParentYangName = "mac-iter"
    switchFdbCommon.EntityData.SegmentPath = "switch-fdb-common"
    switchFdbCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchFdbCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchFdbCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchFdbCommon.EntityData.Children = make(map[string]types.YChild)
    switchFdbCommon.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range switchFdbCommon.Location {
        switchFdbCommon.EntityData.Children[types.GetSegmentPath(&switchFdbCommon.Location[i])] = types.YChild{"Location", &switchFdbCommon.Location[i]}
    }
    switchFdbCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(switchFdbCommon.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location
type Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch FDB information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Number of FDB entries in the table. The type is interface{} with range:
    // 0..4294967295.
    NumEntries interface{}

    // FDB entries contain an entry from the trunk. The type is interface{} with
    // range: 0..255.
    HasTrunkEntry interface{}

    // Message displayed when an FDB entry contains an entry for a trunk member
    // port. The type is string.
    TrunkEntryMessage interface{}

    
    FdbBlock Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock
}

func (location *Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "switch-fdb-common"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fdb-block"] = types.YChild{"FdbBlock", &location.FdbBlock}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["num-entries"] = types.YLeaf{"NumEntries", location.NumEntries}
    location.EntityData.Leafs["has-trunk-entry"] = types.YLeaf{"HasTrunkEntry", location.HasTrunkEntry}
    location.EntityData.Leafs["trunk-entry-message"] = types.YLeaf{"TrunkEntryMessage", location.TrunkEntryMessage}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock
type Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry.
    FdbEntry []Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
}

func (fdbBlock *Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock) GetEntityData() *types.CommonEntityData {
    fdbBlock.EntityData.YFilter = fdbBlock.YFilter
    fdbBlock.EntityData.YangName = "fdb-block"
    fdbBlock.EntityData.BundleName = "cisco_ios_xr"
    fdbBlock.EntityData.ParentYangName = "location"
    fdbBlock.EntityData.SegmentPath = "fdb-block"
    fdbBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbBlock.EntityData.Children = make(map[string]types.YChild)
    fdbBlock.EntityData.Children["fdb-entry"] = types.YChild{"FdbEntry", nil}
    for i := range fdbBlock.FdbEntry {
        fdbBlock.EntityData.Children[types.GetSegmentPath(&fdbBlock.FdbEntry[i])] = types.YChild{"FdbEntry", &fdbBlock.FdbEntry[i]}
    }
    fdbBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fdbBlock.EntityData)
}

// Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
type Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    FdbIndex interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    FdbMacAddr interface{}

    // The type is interface{} with range: 0..4095.
    FdbVlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry_FdbVlan
    FdbVlanHex interface{}

    // Switch port MAC address learned on. The type is interface{} with range:
    // 0..127.
    FdbPort interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbTrapEntry interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbStaticEntry interface{}

    // The type is slice of interface{} with range: 0..255.
    FdbSyncedCores []interface{}
}

func (fdbEntry *Controller_Switch_Oper_Fdb_Mac_MacIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry) GetEntityData() *types.CommonEntityData {
    fdbEntry.EntityData.YFilter = fdbEntry.YFilter
    fdbEntry.EntityData.YangName = "fdb-entry"
    fdbEntry.EntityData.BundleName = "cisco_ios_xr"
    fdbEntry.EntityData.ParentYangName = "fdb-block"
    fdbEntry.EntityData.SegmentPath = "fdb-entry" + "[fdb-index='" + fmt.Sprintf("%v", fdbEntry.FdbIndex) + "']"
    fdbEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbEntry.EntityData.Children = make(map[string]types.YChild)
    fdbEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    fdbEntry.EntityData.Leafs["fdb-index"] = types.YLeaf{"FdbIndex", fdbEntry.FdbIndex}
    fdbEntry.EntityData.Leafs["fdb-mac-addr"] = types.YLeaf{"FdbMacAddr", fdbEntry.FdbMacAddr}
    fdbEntry.EntityData.Leafs["fdb-vlan"] = types.YLeaf{"FdbVlan", fdbEntry.FdbVlan}
    fdbEntry.EntityData.Leafs["fdb-vlan-hex"] = types.YLeaf{"FdbVlanHex", fdbEntry.FdbVlanHex}
    fdbEntry.EntityData.Leafs["fdb-port"] = types.YLeaf{"FdbPort", fdbEntry.FdbPort}
    fdbEntry.EntityData.Leafs["fdb-trap-entry"] = types.YLeaf{"FdbTrapEntry", fdbEntry.FdbTrapEntry}
    fdbEntry.EntityData.Leafs["fdb-static-entry"] = types.YLeaf{"FdbStaticEntry", fdbEntry.FdbStaticEntry}
    fdbEntry.EntityData.Leafs["fdb-synced-cores"] = types.YLeaf{"FdbSyncedCores", fdbEntry.FdbSyncedCores}
    return &(fdbEntry.EntityData)
}

// Controller_Switch_Oper_Fdb_Port
type Controller_Switch_Oper_Fdb_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Fdb_Port_PortIter.
    PortIter []Controller_Switch_Oper_Fdb_Port_PortIter
}

func (port *Controller_Switch_Oper_Fdb_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "fdb"
    port.EntityData.SegmentPath = "port"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range port.PortIter {
        port.EntityData.Children[types.GetSegmentPath(&port.PortIter[i])] = types.YChild{"PortIter", &port.PortIter[i]}
    }
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(port.EntityData)
}

// Controller_Switch_Oper_Fdb_Port_PortIter
type Controller_Switch_Oper_Fdb_Port_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port used for filtering. The type is
    // interface{} with range: 0..127.
    Port interface{}

    
    SwitchFdbCommon Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon
}

func (portIter *Controller_Switch_Oper_Fdb_Port_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "port"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["switch-fdb-common"] = types.YChild{"SwitchFdbCommon", &portIter.SwitchFdbCommon}
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon
type Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location.
    Location []Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location
}

func (switchFdbCommon *Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon) GetEntityData() *types.CommonEntityData {
    switchFdbCommon.EntityData.YFilter = switchFdbCommon.YFilter
    switchFdbCommon.EntityData.YangName = "switch-fdb-common"
    switchFdbCommon.EntityData.BundleName = "cisco_ios_xr"
    switchFdbCommon.EntityData.ParentYangName = "port-iter"
    switchFdbCommon.EntityData.SegmentPath = "switch-fdb-common"
    switchFdbCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchFdbCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchFdbCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchFdbCommon.EntityData.Children = make(map[string]types.YChild)
    switchFdbCommon.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range switchFdbCommon.Location {
        switchFdbCommon.EntityData.Children[types.GetSegmentPath(&switchFdbCommon.Location[i])] = types.YChild{"Location", &switchFdbCommon.Location[i]}
    }
    switchFdbCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(switchFdbCommon.EntityData)
}

// Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location
type Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch FDB information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Number of FDB entries in the table. The type is interface{} with range:
    // 0..4294967295.
    NumEntries interface{}

    // FDB entries contain an entry from the trunk. The type is interface{} with
    // range: 0..255.
    HasTrunkEntry interface{}

    // Message displayed when an FDB entry contains an entry for a trunk member
    // port. The type is string.
    TrunkEntryMessage interface{}

    
    FdbBlock Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock
}

func (location *Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "switch-fdb-common"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fdb-block"] = types.YChild{"FdbBlock", &location.FdbBlock}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["num-entries"] = types.YLeaf{"NumEntries", location.NumEntries}
    location.EntityData.Leafs["has-trunk-entry"] = types.YLeaf{"HasTrunkEntry", location.HasTrunkEntry}
    location.EntityData.Leafs["trunk-entry-message"] = types.YLeaf{"TrunkEntryMessage", location.TrunkEntryMessage}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock
type Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry.
    FdbEntry []Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
}

func (fdbBlock *Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock) GetEntityData() *types.CommonEntityData {
    fdbBlock.EntityData.YFilter = fdbBlock.YFilter
    fdbBlock.EntityData.YangName = "fdb-block"
    fdbBlock.EntityData.BundleName = "cisco_ios_xr"
    fdbBlock.EntityData.ParentYangName = "location"
    fdbBlock.EntityData.SegmentPath = "fdb-block"
    fdbBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbBlock.EntityData.Children = make(map[string]types.YChild)
    fdbBlock.EntityData.Children["fdb-entry"] = types.YChild{"FdbEntry", nil}
    for i := range fdbBlock.FdbEntry {
        fdbBlock.EntityData.Children[types.GetSegmentPath(&fdbBlock.FdbEntry[i])] = types.YChild{"FdbEntry", &fdbBlock.FdbEntry[i]}
    }
    fdbBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fdbBlock.EntityData)
}

// Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry
type Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    FdbIndex interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    FdbMacAddr interface{}

    // The type is interface{} with range: 0..4095.
    FdbVlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry_FdbVlan
    FdbVlanHex interface{}

    // Switch port MAC address learned on. The type is interface{} with range:
    // 0..127.
    FdbPort interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbTrapEntry interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbStaticEntry interface{}

    // The type is slice of interface{} with range: 0..255.
    FdbSyncedCores []interface{}
}

func (fdbEntry *Controller_Switch_Oper_Fdb_Port_PortIter_SwitchFdbCommon_Location_FdbBlock_FdbEntry) GetEntityData() *types.CommonEntityData {
    fdbEntry.EntityData.YFilter = fdbEntry.YFilter
    fdbEntry.EntityData.YangName = "fdb-entry"
    fdbEntry.EntityData.BundleName = "cisco_ios_xr"
    fdbEntry.EntityData.ParentYangName = "fdb-block"
    fdbEntry.EntityData.SegmentPath = "fdb-entry" + "[fdb-index='" + fmt.Sprintf("%v", fdbEntry.FdbIndex) + "']"
    fdbEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbEntry.EntityData.Children = make(map[string]types.YChild)
    fdbEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    fdbEntry.EntityData.Leafs["fdb-index"] = types.YLeaf{"FdbIndex", fdbEntry.FdbIndex}
    fdbEntry.EntityData.Leafs["fdb-mac-addr"] = types.YLeaf{"FdbMacAddr", fdbEntry.FdbMacAddr}
    fdbEntry.EntityData.Leafs["fdb-vlan"] = types.YLeaf{"FdbVlan", fdbEntry.FdbVlan}
    fdbEntry.EntityData.Leafs["fdb-vlan-hex"] = types.YLeaf{"FdbVlanHex", fdbEntry.FdbVlanHex}
    fdbEntry.EntityData.Leafs["fdb-port"] = types.YLeaf{"FdbPort", fdbEntry.FdbPort}
    fdbEntry.EntityData.Leafs["fdb-trap-entry"] = types.YLeaf{"FdbTrapEntry", fdbEntry.FdbTrapEntry}
    fdbEntry.EntityData.Leafs["fdb-static-entry"] = types.YLeaf{"FdbStaticEntry", fdbEntry.FdbStaticEntry}
    fdbEntry.EntityData.Leafs["fdb-synced-cores"] = types.YLeaf{"FdbSyncedCores", fdbEntry.FdbSyncedCores}
    return &(fdbEntry.EntityData)
}

// Controller_Switch_Oper_Fdb_Statistics
type Controller_Switch_Oper_Fdb_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Fdb_Statistics_Location.
    Location []Controller_Switch_Oper_Fdb_Statistics_Location
}

func (statistics *Controller_Switch_Oper_Fdb_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "fdb"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range statistics.Location {
        statistics.EntityData.Children[types.GetSegmentPath(&statistics.Location[i])] = types.YChild{"Location", &statistics.Location[i]}
    }
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Controller_Switch_Oper_Fdb_Statistics_Location
type Controller_Switch_Oper_Fdb_Statistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    
    CounterInfo Controller_Switch_Oper_Fdb_Statistics_Location_CounterInfo

    // The type is slice of Controller_Switch_Oper_Fdb_Statistics_Location_CoreId.
    CoreId []Controller_Switch_Oper_Fdb_Statistics_Location_CoreId
}

func (location *Controller_Switch_Oper_Fdb_Statistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["counter-info"] = types.YChild{"CounterInfo", &location.CounterInfo}
    location.EntityData.Children["core-id"] = types.YChild{"CoreId", nil}
    for i := range location.CoreId {
        location.EntityData.Children[types.GetSegmentPath(&location.CoreId[i])] = types.YChild{"CoreId", &location.CoreId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Fdb_Statistics_Location_CounterInfo
type Controller_Switch_Oper_Fdb_Statistics_Location_CounterInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    FdbShadowEntries interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbMaxShadowEntries interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbMaxHashChain interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbEntriesAdded interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbEntriesDeleted interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbEntriesUpdated interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbFlushes interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAddressUpdates interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbNewAddresses interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAgedUpdates interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbTransplantUpdates interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbForwardingUpdates interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAddressInsertErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAddressUpdateErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbMemoryErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAllocationErrors interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbAddressUpdatesQueued interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbAddressQueueFull interface{}

    // The type is interface{} with range: 0..4294967295.
    FdbForwardingUpdatesQueued interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbForwardingQueueFull interface{}
}

func (counterInfo *Controller_Switch_Oper_Fdb_Statistics_Location_CounterInfo) GetEntityData() *types.CommonEntityData {
    counterInfo.EntityData.YFilter = counterInfo.YFilter
    counterInfo.EntityData.YangName = "counter-info"
    counterInfo.EntityData.BundleName = "cisco_ios_xr"
    counterInfo.EntityData.ParentYangName = "location"
    counterInfo.EntityData.SegmentPath = "counter-info"
    counterInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counterInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counterInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counterInfo.EntityData.Children = make(map[string]types.YChild)
    counterInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    counterInfo.EntityData.Leafs["fdb-shadow-entries"] = types.YLeaf{"FdbShadowEntries", counterInfo.FdbShadowEntries}
    counterInfo.EntityData.Leafs["fdb-max-shadow-entries"] = types.YLeaf{"FdbMaxShadowEntries", counterInfo.FdbMaxShadowEntries}
    counterInfo.EntityData.Leafs["fdb-max-hash-chain"] = types.YLeaf{"FdbMaxHashChain", counterInfo.FdbMaxHashChain}
    counterInfo.EntityData.Leafs["fdb-entries-added"] = types.YLeaf{"FdbEntriesAdded", counterInfo.FdbEntriesAdded}
    counterInfo.EntityData.Leafs["fdb-entries-deleted"] = types.YLeaf{"FdbEntriesDeleted", counterInfo.FdbEntriesDeleted}
    counterInfo.EntityData.Leafs["fdb-entries-updated"] = types.YLeaf{"FdbEntriesUpdated", counterInfo.FdbEntriesUpdated}
    counterInfo.EntityData.Leafs["fdb-flushes"] = types.YLeaf{"FdbFlushes", counterInfo.FdbFlushes}
    counterInfo.EntityData.Leafs["fdb-address-updates"] = types.YLeaf{"FdbAddressUpdates", counterInfo.FdbAddressUpdates}
    counterInfo.EntityData.Leafs["fdb-new-addresses"] = types.YLeaf{"FdbNewAddresses", counterInfo.FdbNewAddresses}
    counterInfo.EntityData.Leafs["fdb-aged-updates"] = types.YLeaf{"FdbAgedUpdates", counterInfo.FdbAgedUpdates}
    counterInfo.EntityData.Leafs["fdb-transplant-updates"] = types.YLeaf{"FdbTransplantUpdates", counterInfo.FdbTransplantUpdates}
    counterInfo.EntityData.Leafs["fdb-forwarding-updates"] = types.YLeaf{"FdbForwardingUpdates", counterInfo.FdbForwardingUpdates}
    counterInfo.EntityData.Leafs["fdb-address-insert-errors"] = types.YLeaf{"FdbAddressInsertErrors", counterInfo.FdbAddressInsertErrors}
    counterInfo.EntityData.Leafs["fdb-address-update-errors"] = types.YLeaf{"FdbAddressUpdateErrors", counterInfo.FdbAddressUpdateErrors}
    counterInfo.EntityData.Leafs["fdb-memory-errors"] = types.YLeaf{"FdbMemoryErrors", counterInfo.FdbMemoryErrors}
    counterInfo.EntityData.Leafs["fdb-allocation-errors"] = types.YLeaf{"FdbAllocationErrors", counterInfo.FdbAllocationErrors}
    counterInfo.EntityData.Leafs["fdb-address-updates-queued"] = types.YLeaf{"FdbAddressUpdatesQueued", counterInfo.FdbAddressUpdatesQueued}
    counterInfo.EntityData.Leafs["fdb-address-queue-full"] = types.YLeaf{"FdbAddressQueueFull", counterInfo.FdbAddressQueueFull}
    counterInfo.EntityData.Leafs["fdb-forwarding-updates-queued"] = types.YLeaf{"FdbForwardingUpdatesQueued", counterInfo.FdbForwardingUpdatesQueued}
    counterInfo.EntityData.Leafs["fdb-forwarding-queue-full"] = types.YLeaf{"FdbForwardingQueueFull", counterInfo.FdbForwardingQueueFull}
    return &(counterInfo.EntityData)
}

// Controller_Switch_Oper_Fdb_Statistics_Location_CoreId
type Controller_Switch_Oper_Fdb_Statistics_Location_CoreId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range:
    // -2147483648..2147483647.
    Core interface{}

    // The type is interface{} with range: 0..4294967295.
    CoreEntries interface{}

    // The type is interface{} with range: 0..4294967295.
    CoreStaticEntries interface{}
}

func (coreId *Controller_Switch_Oper_Fdb_Statistics_Location_CoreId) GetEntityData() *types.CommonEntityData {
    coreId.EntityData.YFilter = coreId.YFilter
    coreId.EntityData.YangName = "core-id"
    coreId.EntityData.BundleName = "cisco_ios_xr"
    coreId.EntityData.ParentYangName = "location"
    coreId.EntityData.SegmentPath = "core-id" + "[core='" + fmt.Sprintf("%v", coreId.Core) + "']"
    coreId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    coreId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    coreId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    coreId.EntityData.Children = make(map[string]types.YChild)
    coreId.EntityData.Leafs = make(map[string]types.YLeaf)
    coreId.EntityData.Leafs["core"] = types.YLeaf{"Core", coreId.Core}
    coreId.EntityData.Leafs["core-entries"] = types.YLeaf{"CoreEntries", coreId.CoreEntries}
    coreId.EntityData.Leafs["core-static-entries"] = types.YLeaf{"CoreStaticEntries", coreId.CoreStaticEntries}
    return &(coreId.EntityData)
}

// Controller_Switch_Oper_Fdb_SwitchFdbCommon
type Controller_Switch_Oper_Fdb_SwitchFdbCommon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location.
    Location []Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location
}

func (switchFdbCommon *Controller_Switch_Oper_Fdb_SwitchFdbCommon) GetEntityData() *types.CommonEntityData {
    switchFdbCommon.EntityData.YFilter = switchFdbCommon.YFilter
    switchFdbCommon.EntityData.YangName = "switch-fdb-common"
    switchFdbCommon.EntityData.BundleName = "cisco_ios_xr"
    switchFdbCommon.EntityData.ParentYangName = "fdb"
    switchFdbCommon.EntityData.SegmentPath = "switch-fdb-common"
    switchFdbCommon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchFdbCommon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchFdbCommon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchFdbCommon.EntityData.Children = make(map[string]types.YChild)
    switchFdbCommon.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range switchFdbCommon.Location {
        switchFdbCommon.EntityData.Children[types.GetSegmentPath(&switchFdbCommon.Location[i])] = types.YChild{"Location", &switchFdbCommon.Location[i]}
    }
    switchFdbCommon.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(switchFdbCommon.EntityData)
}

// Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location
type Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch FDB information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Number of FDB entries in the table. The type is interface{} with range:
    // 0..4294967295.
    NumEntries interface{}

    // FDB entries contain an entry from the trunk. The type is interface{} with
    // range: 0..255.
    HasTrunkEntry interface{}

    // Message displayed when an FDB entry contains an entry for a trunk member
    // port. The type is string.
    TrunkEntryMessage interface{}

    
    FdbBlock Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock
}

func (location *Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "switch-fdb-common"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["fdb-block"] = types.YChild{"FdbBlock", &location.FdbBlock}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["num-entries"] = types.YLeaf{"NumEntries", location.NumEntries}
    location.EntityData.Leafs["has-trunk-entry"] = types.YLeaf{"HasTrunkEntry", location.HasTrunkEntry}
    location.EntityData.Leafs["trunk-entry-message"] = types.YLeaf{"TrunkEntryMessage", location.TrunkEntryMessage}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock
type Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry.
    FdbEntry []Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry
}

func (fdbBlock *Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock) GetEntityData() *types.CommonEntityData {
    fdbBlock.EntityData.YFilter = fdbBlock.YFilter
    fdbBlock.EntityData.YangName = "fdb-block"
    fdbBlock.EntityData.BundleName = "cisco_ios_xr"
    fdbBlock.EntityData.ParentYangName = "location"
    fdbBlock.EntityData.SegmentPath = "fdb-block"
    fdbBlock.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbBlock.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbBlock.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbBlock.EntityData.Children = make(map[string]types.YChild)
    fdbBlock.EntityData.Children["fdb-entry"] = types.YChild{"FdbEntry", nil}
    for i := range fdbBlock.FdbEntry {
        fdbBlock.EntityData.Children[types.GetSegmentPath(&fdbBlock.FdbEntry[i])] = types.YChild{"FdbEntry", &fdbBlock.FdbEntry[i]}
    }
    fdbBlock.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fdbBlock.EntityData)
}

// Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry
type Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    FdbIndex interface{}

    // The type is string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    FdbMacAddr interface{}

    // The type is interface{} with range: 0..4095.
    FdbVlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry_FdbVlan
    FdbVlanHex interface{}

    // Switch port MAC address learned on. The type is interface{} with range:
    // 0..127.
    FdbPort interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbTrapEntry interface{}

    // The type is EsdmaSwitchYesNoEnum.
    FdbStaticEntry interface{}

    // The type is slice of interface{} with range: 0..255.
    FdbSyncedCores []interface{}
}

func (fdbEntry *Controller_Switch_Oper_Fdb_SwitchFdbCommon_Location_FdbBlock_FdbEntry) GetEntityData() *types.CommonEntityData {
    fdbEntry.EntityData.YFilter = fdbEntry.YFilter
    fdbEntry.EntityData.YangName = "fdb-entry"
    fdbEntry.EntityData.BundleName = "cisco_ios_xr"
    fdbEntry.EntityData.ParentYangName = "fdb-block"
    fdbEntry.EntityData.SegmentPath = "fdb-entry" + "[fdb-index='" + fmt.Sprintf("%v", fdbEntry.FdbIndex) + "']"
    fdbEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fdbEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fdbEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fdbEntry.EntityData.Children = make(map[string]types.YChild)
    fdbEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    fdbEntry.EntityData.Leafs["fdb-index"] = types.YLeaf{"FdbIndex", fdbEntry.FdbIndex}
    fdbEntry.EntityData.Leafs["fdb-mac-addr"] = types.YLeaf{"FdbMacAddr", fdbEntry.FdbMacAddr}
    fdbEntry.EntityData.Leafs["fdb-vlan"] = types.YLeaf{"FdbVlan", fdbEntry.FdbVlan}
    fdbEntry.EntityData.Leafs["fdb-vlan-hex"] = types.YLeaf{"FdbVlanHex", fdbEntry.FdbVlanHex}
    fdbEntry.EntityData.Leafs["fdb-port"] = types.YLeaf{"FdbPort", fdbEntry.FdbPort}
    fdbEntry.EntityData.Leafs["fdb-trap-entry"] = types.YLeaf{"FdbTrapEntry", fdbEntry.FdbTrapEntry}
    fdbEntry.EntityData.Leafs["fdb-static-entry"] = types.YLeaf{"FdbStaticEntry", fdbEntry.FdbStaticEntry}
    fdbEntry.EntityData.Leafs["fdb-synced-cores"] = types.YLeaf{"FdbSyncedCores", fdbEntry.FdbSyncedCores}
    return &(fdbEntry.EntityData)
}

// Controller_Switch_Oper_Vlan
type Controller_Switch_Oper_Vlan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Rules Controller_Switch_Oper_Vlan_Rules

    
    Information Controller_Switch_Oper_Vlan_Information

    
    VlanDetail Controller_Switch_Oper_Vlan_VlanDetail

    
    Membership Controller_Switch_Oper_Vlan_Membership
}

func (vlan *Controller_Switch_Oper_Vlan) GetEntityData() *types.CommonEntityData {
    vlan.EntityData.YFilter = vlan.YFilter
    vlan.EntityData.YangName = "vlan"
    vlan.EntityData.BundleName = "cisco_ios_xr"
    vlan.EntityData.ParentYangName = "oper"
    vlan.EntityData.SegmentPath = "vlan"
    vlan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlan.EntityData.Children = make(map[string]types.YChild)
    vlan.EntityData.Children["rules"] = types.YChild{"Rules", &vlan.Rules}
    vlan.EntityData.Children["information"] = types.YChild{"Information", &vlan.Information}
    vlan.EntityData.Children["vlan-detail"] = types.YChild{"VlanDetail", &vlan.VlanDetail}
    vlan.EntityData.Children["membership"] = types.YChild{"Membership", &vlan.Membership}
    vlan.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlan.EntityData)
}

// Controller_Switch_Oper_Vlan_Rules
type Controller_Switch_Oper_Vlan_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Vlan_Rules_Location.
    Location []Controller_Switch_Oper_Vlan_Rules_Location
}

func (rules *Controller_Switch_Oper_Vlan_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "vlan"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = make(map[string]types.YChild)
    rules.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rules.Location {
        rules.EntityData.Children[types.GetSegmentPath(&rules.Location[i])] = types.YChild{"Location", &rules.Location[i]}
    }
    rules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rules.EntityData)
}

// Controller_Switch_Oper_Vlan_Rules_Location
type Controller_Switch_Oper_Vlan_Rules_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest to display the
    // VLAN configuration for. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of Controller_Switch_Oper_Vlan_Rules_Location_PortIter.
    PortIter []Controller_Switch_Oper_Vlan_Rules_Location_PortIter
}

func (location *Controller_Switch_Oper_Vlan_Rules_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rules"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Vlan_Rules_Location_PortIter
type Controller_Switch_Oper_Vlan_Rules_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId.
    VlanId []Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId
}

func (portIter *Controller_Switch_Oper_Vlan_Rules_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["vlan-id"] = types.YChild{"VlanId", nil}
    for i := range portIter.VlanId {
        portIter.EntityData.Children[types.GetSegmentPath(&portIter.VlanId[i])] = types.YChild{"VlanId", &portIter.VlanId[i]}
    }
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId
type Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This VLAN represents a VLAN membership for this
    // switch port. Multiple rules may exist to support this VLAN. The type is
    // interface{} with range: 0..4095.
    Vlan interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId_RuleId.
    RuleId []Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId_RuleId
}

func (vlanId *Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "port-iter"
    vlanId.EntityData.SegmentPath = "vlan-id" + "[vlan='" + fmt.Sprintf("%v", vlanId.Vlan) + "']"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = make(map[string]types.YChild)
    vlanId.EntityData.Children["rule-id"] = types.YChild{"RuleId", nil}
    for i := range vlanId.RuleId {
        vlanId.EntityData.Children[types.GetSegmentPath(&vlanId.RuleId[i])] = types.YChild{"RuleId", &vlanId.RuleId[i]}
    }
    vlanId.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanId.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanId.Vlan}
    return &(vlanId.EntityData)
}

// Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId_RuleId
type Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId_RuleId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet direction this rule applies to. The type is SwitchDataDirectionEnum.
    Direction interface{}

    // The type is SwitchTableTypeEnum.
    MatchTable interface{}

    // The type is SwitchMatchTypeEnum.
    MatchType interface{}

    // The type is interface{} with range: 0..4095.
    MatchVlanId interface{}

    // The type is SwitchActionTypeEnum.
    Action interface{}

    // The type is interface{} with range: 0..4095.
    ActionVlanId interface{}
}

func (ruleId *Controller_Switch_Oper_Vlan_Rules_Location_PortIter_VlanId_RuleId) GetEntityData() *types.CommonEntityData {
    ruleId.EntityData.YFilter = ruleId.YFilter
    ruleId.EntityData.YangName = "rule-id"
    ruleId.EntityData.BundleName = "cisco_ios_xr"
    ruleId.EntityData.ParentYangName = "vlan-id"
    ruleId.EntityData.SegmentPath = "rule-id"
    ruleId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleId.EntityData.Children = make(map[string]types.YChild)
    ruleId.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleId.EntityData.Leafs["direction"] = types.YLeaf{"Direction", ruleId.Direction}
    ruleId.EntityData.Leafs["match-table"] = types.YLeaf{"MatchTable", ruleId.MatchTable}
    ruleId.EntityData.Leafs["match-type"] = types.YLeaf{"MatchType", ruleId.MatchType}
    ruleId.EntityData.Leafs["match-vlan-id"] = types.YLeaf{"MatchVlanId", ruleId.MatchVlanId}
    ruleId.EntityData.Leafs["action"] = types.YLeaf{"Action", ruleId.Action}
    ruleId.EntityData.Leafs["action-vlan-id"] = types.YLeaf{"ActionVlanId", ruleId.ActionVlanId}
    return &(ruleId.EntityData)
}

// Controller_Switch_Oper_Vlan_Information
type Controller_Switch_Oper_Vlan_Information struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Vlan_Information_Location.
    Location []Controller_Switch_Oper_Vlan_Information_Location

    
    Summary Controller_Switch_Oper_Vlan_Information_Summary
}

func (information *Controller_Switch_Oper_Vlan_Information) GetEntityData() *types.CommonEntityData {
    information.EntityData.YFilter = information.YFilter
    information.EntityData.YangName = "information"
    information.EntityData.BundleName = "cisco_ios_xr"
    information.EntityData.ParentYangName = "vlan"
    information.EntityData.SegmentPath = "information"
    information.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    information.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    information.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    information.EntityData.Children = make(map[string]types.YChild)
    information.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range information.Location {
        information.EntityData.Children[types.GetSegmentPath(&information.Location[i])] = types.YChild{"Location", &information.Location[i]}
    }
    information.EntityData.Children["summary"] = types.YChild{"Summary", &information.Summary}
    information.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(information.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Location
type Controller_Switch_Oper_Vlan_Information_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to list the switch VLAN information for. The
    // type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Rack serial number. The type is string.
    SerialNum interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Information_Location_SdrId.
    SdrId []Controller_Switch_Oper_Vlan_Information_Location_SdrId
}

func (location *Controller_Switch_Oper_Vlan_Information_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "information"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range location.SdrId {
        location.EntityData.Children[types.GetSegmentPath(&location.SdrId[i])] = types.YChild{"SdrId", &location.SdrId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["serial-num"] = types.YLeaf{"SerialNum", location.SerialNum}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Location_SdrId
type Controller_Switch_Oper_Vlan_Information_Location_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    // The type is string.
    EsdSdrName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId.
    VlanId []Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId
}

func (sdrId *Controller_Switch_Oper_Vlan_Information_Location_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "location"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["vlan-id"] = types.YChild{"VlanId", nil}
    for i := range sdrId.VlanId {
        sdrId.EntityData.Children[types.GetSegmentPath(&sdrId.VlanId[i])] = types.YChild{"VlanId", &sdrId.VlanId[i]}
    }
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    sdrId.EntityData.Leafs["esd-sdr-name"] = types.YLeaf{"EsdSdrName", sdrId.EsdSdrName}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId
type Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4095.
    Vlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId_Vlan
    VlanHex interface{}

    // The type is string.
    VlanUse interface{}
}

func (vlanId *Controller_Switch_Oper_Vlan_Information_Location_SdrId_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "sdr-id"
    vlanId.EntityData.SegmentPath = "vlan-id" + "[vlan='" + fmt.Sprintf("%v", vlanId.Vlan) + "']"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = make(map[string]types.YChild)
    vlanId.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanId.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanId.Vlan}
    vlanId.EntityData.Leafs["vlan-hex"] = types.YLeaf{"VlanHex", vlanId.VlanHex}
    vlanId.EntityData.Leafs["vlan-use"] = types.YLeaf{"VlanUse", vlanId.VlanUse}
    return &(vlanId.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Summary
type Controller_Switch_Oper_Vlan_Information_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Vlan_Information_Summary_SdrId.
    SdrId []Controller_Switch_Oper_Vlan_Information_Summary_SdrId
}

func (summary *Controller_Switch_Oper_Vlan_Information_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "information"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range summary.SdrId {
        summary.EntityData.Children[types.GetSegmentPath(&summary.SdrId[i])] = types.YChild{"SdrId", &summary.SdrId[i]}
    }
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Summary_SdrId
type Controller_Switch_Oper_Vlan_Information_Summary_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    // The type is string.
    EsdSdrName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId.
    VlanId []Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId
}

func (sdrId *Controller_Switch_Oper_Vlan_Information_Summary_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "summary"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["vlan-id"] = types.YChild{"VlanId", nil}
    for i := range sdrId.VlanId {
        sdrId.EntityData.Children[types.GetSegmentPath(&sdrId.VlanId[i])] = types.YChild{"VlanId", &sdrId.VlanId[i]}
    }
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    sdrId.EntityData.Leafs["esd-sdr-name"] = types.YLeaf{"EsdSdrName", sdrId.EsdSdrName}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId
type Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4095.
    Vlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId_Vlan
    VlanHex interface{}

    // The type is string.
    VlanUse interface{}
}

func (vlanId *Controller_Switch_Oper_Vlan_Information_Summary_SdrId_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "sdr-id"
    vlanId.EntityData.SegmentPath = "vlan-id" + "[vlan='" + fmt.Sprintf("%v", vlanId.Vlan) + "']"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = make(map[string]types.YChild)
    vlanId.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanId.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanId.Vlan}
    vlanId.EntityData.Leafs["vlan-hex"] = types.YLeaf{"VlanHex", vlanId.VlanHex}
    vlanId.EntityData.Leafs["vlan-use"] = types.YLeaf{"VlanUse", vlanId.VlanUse}
    return &(vlanId.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail
type Controller_Switch_Oper_Vlan_VlanDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Vlan_VlanDetail_VlanId.
    VlanId []Controller_Switch_Oper_Vlan_VlanDetail_VlanId
}

func (vlanDetail *Controller_Switch_Oper_Vlan_VlanDetail) GetEntityData() *types.CommonEntityData {
    vlanDetail.EntityData.YFilter = vlanDetail.YFilter
    vlanDetail.EntityData.YangName = "vlan-detail"
    vlanDetail.EntityData.BundleName = "cisco_ios_xr"
    vlanDetail.EntityData.ParentYangName = "vlan"
    vlanDetail.EntityData.SegmentPath = "vlan-detail"
    vlanDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanDetail.EntityData.Children = make(map[string]types.YChild)
    vlanDetail.EntityData.Children["vlan-id"] = types.YChild{"VlanId", nil}
    for i := range vlanDetail.VlanId {
        vlanDetail.EntityData.Children[types.GetSegmentPath(&vlanDetail.VlanId[i])] = types.YChild{"VlanId", &vlanDetail.VlanId[i]}
    }
    vlanDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vlanDetail.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail_VlanId
type Controller_Switch_Oper_Vlan_VlanDetail_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4095.
    Vlan interface{}

    
    Rules Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules
}

func (vlanId *Controller_Switch_Oper_Vlan_VlanDetail_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "vlan-detail"
    vlanId.EntityData.SegmentPath = "vlan-id" + "[vlan='" + fmt.Sprintf("%v", vlanId.Vlan) + "']"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = make(map[string]types.YChild)
    vlanId.EntityData.Children["rules"] = types.YChild{"Rules", &vlanId.Rules}
    vlanId.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanId.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanId.Vlan}
    return &(vlanId.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules
type Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location.
    Location []Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location
}

func (rules *Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "vlan-id"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = make(map[string]types.YChild)
    rules.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range rules.Location {
        rules.EntityData.Children[types.GetSegmentPath(&rules.Location[i])] = types.YChild{"Location", &rules.Location[i]}
    }
    rules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rules.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location
type Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter.
    PortIter []Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter
}

func (location *Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "rules"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter
type Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter_RuleId.
    RuleId []Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter_RuleId
}

func (portIter *Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["rule-id"] = types.YChild{"RuleId", nil}
    for i := range portIter.RuleId {
        portIter.EntityData.Children[types.GetSegmentPath(&portIter.RuleId[i])] = types.YChild{"RuleId", &portIter.RuleId[i]}
    }
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter_RuleId
type Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter_RuleId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packet direction this rule applies to. The type is SwitchDataDirectionEnum.
    Direction interface{}

    // The type is SwitchTableTypeEnum.
    MatchTable interface{}

    // The type is SwitchMatchTypeEnum.
    MatchType interface{}

    // The type is interface{} with range: 0..4095.
    MatchVlanId interface{}

    // The type is SwitchActionTypeEnum.
    Action interface{}

    // The type is interface{} with range: 0..4095.
    ActionVlanId interface{}
}

func (ruleId *Controller_Switch_Oper_Vlan_VlanDetail_VlanId_Rules_Location_PortIter_RuleId) GetEntityData() *types.CommonEntityData {
    ruleId.EntityData.YFilter = ruleId.YFilter
    ruleId.EntityData.YangName = "rule-id"
    ruleId.EntityData.BundleName = "cisco_ios_xr"
    ruleId.EntityData.ParentYangName = "port-iter"
    ruleId.EntityData.SegmentPath = "rule-id"
    ruleId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleId.EntityData.Children = make(map[string]types.YChild)
    ruleId.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleId.EntityData.Leafs["direction"] = types.YLeaf{"Direction", ruleId.Direction}
    ruleId.EntityData.Leafs["match-table"] = types.YLeaf{"MatchTable", ruleId.MatchTable}
    ruleId.EntityData.Leafs["match-type"] = types.YLeaf{"MatchType", ruleId.MatchType}
    ruleId.EntityData.Leafs["match-vlan-id"] = types.YLeaf{"MatchVlanId", ruleId.MatchVlanId}
    ruleId.EntityData.Leafs["action"] = types.YLeaf{"Action", ruleId.Action}
    ruleId.EntityData.Leafs["action-vlan-id"] = types.YLeaf{"ActionVlanId", ruleId.ActionVlanId}
    return &(ruleId.EntityData)
}

// Controller_Switch_Oper_Vlan_Membership
type Controller_Switch_Oper_Vlan_Membership struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Vlan_Membership_Location.
    Location []Controller_Switch_Oper_Vlan_Membership_Location
}

func (membership *Controller_Switch_Oper_Vlan_Membership) GetEntityData() *types.CommonEntityData {
    membership.EntityData.YFilter = membership.YFilter
    membership.EntityData.YangName = "membership"
    membership.EntityData.BundleName = "cisco_ios_xr"
    membership.EntityData.ParentYangName = "vlan"
    membership.EntityData.SegmentPath = "membership"
    membership.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    membership.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    membership.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    membership.EntityData.Children = make(map[string]types.YChild)
    membership.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range membership.Location {
        membership.EntityData.Children[types.GetSegmentPath(&membership.Location[i])] = types.YChild{"Location", &membership.Location[i]}
    }
    membership.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(membership.EntityData)
}

// Controller_Switch_Oper_Vlan_Membership_Location
type Controller_Switch_Oper_Vlan_Membership_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest to display the
    // VLAN configuration for. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Vlan_Membership_Location_VlanId.
    VlanId []Controller_Switch_Oper_Vlan_Membership_Location_VlanId
}

func (location *Controller_Switch_Oper_Vlan_Membership_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "membership"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["vlan-id"] = types.YChild{"VlanId", nil}
    for i := range location.VlanId {
        location.EntityData.Children[types.GetSegmentPath(&location.VlanId[i])] = types.YChild{"VlanId", &location.VlanId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Vlan_Membership_Location_VlanId
type Controller_Switch_Oper_Vlan_Membership_Location_VlanId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4095.
    Vlan interface{}

    // The type is string with range: 0..4095. Refers to
    // sysadmin_controllers_asr9k.Controller_Switch_Oper_Vlan_Membership_Location_VlanId_Vlan
    VlanHex interface{}

    // The type is slice of interface{} with range: 0..255.
    Port []interface{}
}

func (vlanId *Controller_Switch_Oper_Vlan_Membership_Location_VlanId) GetEntityData() *types.CommonEntityData {
    vlanId.EntityData.YFilter = vlanId.YFilter
    vlanId.EntityData.YangName = "vlan-id"
    vlanId.EntityData.BundleName = "cisco_ios_xr"
    vlanId.EntityData.ParentYangName = "location"
    vlanId.EntityData.SegmentPath = "vlan-id" + "[vlan='" + fmt.Sprintf("%v", vlanId.Vlan) + "']"
    vlanId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vlanId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vlanId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vlanId.EntityData.Children = make(map[string]types.YChild)
    vlanId.EntityData.Leafs = make(map[string]types.YLeaf)
    vlanId.EntityData.Leafs["vlan"] = types.YLeaf{"Vlan", vlanId.Vlan}
    vlanId.EntityData.Leafs["vlan-hex"] = types.YLeaf{"VlanHex", vlanId.VlanHex}
    vlanId.EntityData.Leafs["port"] = types.YLeaf{"Port", vlanId.Port}
    return &(vlanId.EntityData)
}

// Controller_Switch_Oper_Esd
type Controller_Switch_Oper_Esd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable process instance names. The type is slice of
    // Controller_Switch_Oper_Esd_Instance.
    Instance []Controller_Switch_Oper_Esd_Instance
}

func (esd *Controller_Switch_Oper_Esd) GetEntityData() *types.CommonEntityData {
    esd.EntityData.YFilter = esd.YFilter
    esd.EntityData.YangName = "esd"
    esd.EntityData.BundleName = "cisco_ios_xr"
    esd.EntityData.ParentYangName = "oper"
    esd.EntityData.SegmentPath = "esd"
    esd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esd.EntityData.Children = make(map[string]types.YChild)
    esd.EntityData.Children["instance"] = types.YChild{"Instance", nil}
    for i := range esd.Instance {
        esd.EntityData.Children[types.GetSegmentPath(&esd.Instance[i])] = types.YChild{"Instance", &esd.Instance[i]}
    }
    esd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(esd.EntityData)
}

// Controller_Switch_Oper_Esd_Instance
// show traceable process instance names
type Controller_Switch_Oper_Esd_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    ProcessName interface{}

    // show traceable processes. The type is slice of
    // Controller_Switch_Oper_Esd_Instance_Trace.
    Trace []Controller_Switch_Oper_Esd_Instance_Trace
}

func (instance *Controller_Switch_Oper_Esd_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "esd"
    instance.EntityData.SegmentPath = "instance" + "[process_name='" + fmt.Sprintf("%v", instance.ProcessName) + "']"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range instance.Trace {
        instance.EntityData.Children[types.GetSegmentPath(&instance.Trace[i])] = types.YChild{"Trace", &instance.Trace[i]}
    }
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["process_name"] = types.YLeaf{"ProcessName", instance.ProcessName}
    return &(instance.EntityData)
}

// Controller_Switch_Oper_Esd_Instance_Trace
// show traceable processes
type Controller_Switch_Oper_Esd_Instance_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Controller_Switch_Oper_Esd_Instance_Trace_Location.
    Location []Controller_Switch_Oper_Esd_Instance_Trace_Location
}

func (trace *Controller_Switch_Oper_Esd_Instance_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "instance"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Controller_Switch_Oper_Esd_Instance_Trace_Location
type Controller_Switch_Oper_Esd_Instance_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions.
    AllOptions []Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions
}

func (location *Controller_Switch_Oper_Esd_Instance_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions
type Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions_TraceBlocks
type Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Controller_Switch_Oper_Esd_Instance_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

// Controller_Switch_Oper_MgmtAgent
type Controller_Switch_Oper_MgmtAgent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Esdma Controller_Switch_Oper_MgmtAgent_Esdma

    
    Connections Controller_Switch_Oper_MgmtAgent_Connections
}

func (mgmtAgent *Controller_Switch_Oper_MgmtAgent) GetEntityData() *types.CommonEntityData {
    mgmtAgent.EntityData.YFilter = mgmtAgent.YFilter
    mgmtAgent.EntityData.YangName = "mgmt-agent"
    mgmtAgent.EntityData.BundleName = "cisco_ios_xr"
    mgmtAgent.EntityData.ParentYangName = "oper"
    mgmtAgent.EntityData.SegmentPath = "mgmt-agent"
    mgmtAgent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mgmtAgent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mgmtAgent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mgmtAgent.EntityData.Children = make(map[string]types.YChild)
    mgmtAgent.EntityData.Children["esdma"] = types.YChild{"Esdma", &mgmtAgent.Esdma}
    mgmtAgent.EntityData.Children["connections"] = types.YChild{"Connections", &mgmtAgent.Connections}
    mgmtAgent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mgmtAgent.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Esdma
type Controller_Switch_Oper_MgmtAgent_Esdma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // show traceable processes. The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Esdma_Trace.
    Trace []Controller_Switch_Oper_MgmtAgent_Esdma_Trace
}

func (esdma *Controller_Switch_Oper_MgmtAgent_Esdma) GetEntityData() *types.CommonEntityData {
    esdma.EntityData.YFilter = esdma.YFilter
    esdma.EntityData.YangName = "esdma"
    esdma.EntityData.BundleName = "cisco_ios_xr"
    esdma.EntityData.ParentYangName = "mgmt-agent"
    esdma.EntityData.SegmentPath = "esdma"
    esdma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esdma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esdma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esdma.EntityData.Children = make(map[string]types.YChild)
    esdma.EntityData.Children["trace"] = types.YChild{"Trace", nil}
    for i := range esdma.Trace {
        esdma.EntityData.Children[types.GetSegmentPath(&esdma.Trace[i])] = types.YChild{"Trace", &esdma.Trace[i]}
    }
    esdma.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(esdma.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Esdma_Trace
// show traceable processes
type Controller_Switch_Oper_MgmtAgent_Esdma_Trace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Buffer interface{}

    // The type is slice of Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location.
    Location []Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location
}

func (trace *Controller_Switch_Oper_MgmtAgent_Esdma_Trace) GetEntityData() *types.CommonEntityData {
    trace.EntityData.YFilter = trace.YFilter
    trace.EntityData.YangName = "trace"
    trace.EntityData.BundleName = "cisco_ios_xr"
    trace.EntityData.ParentYangName = "esdma"
    trace.EntityData.SegmentPath = "trace" + "[buffer='" + fmt.Sprintf("%v", trace.Buffer) + "']"
    trace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trace.EntityData.Children = make(map[string]types.YChild)
    trace.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trace.Location {
        trace.EntityData.Children[types.GetSegmentPath(&trace.Location[i])] = types.YChild{"Location", &trace.Location[i]}
    }
    trace.EntityData.Leafs = make(map[string]types.YLeaf)
    trace.EntityData.Leafs["buffer"] = types.YLeaf{"Buffer", trace.Buffer}
    return &(trace.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location
type Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    LocationName interface{}

    // The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions.
    AllOptions []Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions
}

func (location *Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trace"
    location.EntityData.SegmentPath = "location" + "[location_name='" + fmt.Sprintf("%v", location.LocationName) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["all-options"] = types.YChild{"AllOptions", nil}
    for i := range location.AllOptions {
        location.EntityData.Children[types.GetSegmentPath(&location.AllOptions[i])] = types.YChild{"AllOptions", &location.AllOptions[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["location_name"] = types.YLeaf{"LocationName", location.LocationName}
    return &(location.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions
type Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Option interface{}

    // The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions_TraceBlocks.
    TraceBlocks []Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions_TraceBlocks
}

func (allOptions *Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions) GetEntityData() *types.CommonEntityData {
    allOptions.EntityData.YFilter = allOptions.YFilter
    allOptions.EntityData.YangName = "all-options"
    allOptions.EntityData.BundleName = "cisco_ios_xr"
    allOptions.EntityData.ParentYangName = "location"
    allOptions.EntityData.SegmentPath = "all-options" + "[option='" + fmt.Sprintf("%v", allOptions.Option) + "']"
    allOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allOptions.EntityData.Children = make(map[string]types.YChild)
    allOptions.EntityData.Children["trace-blocks"] = types.YChild{"TraceBlocks", nil}
    for i := range allOptions.TraceBlocks {
        allOptions.EntityData.Children[types.GetSegmentPath(&allOptions.TraceBlocks[i])] = types.YChild{"TraceBlocks", &allOptions.TraceBlocks[i]}
    }
    allOptions.EntityData.Leafs = make(map[string]types.YLeaf)
    allOptions.EntityData.Leafs["option"] = types.YLeaf{"Option", allOptions.Option}
    return &(allOptions.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions_TraceBlocks
type Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions_TraceBlocks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trace output block. The type is string.
    Data interface{}
}

func (traceBlocks *Controller_Switch_Oper_MgmtAgent_Esdma_Trace_Location_AllOptions_TraceBlocks) GetEntityData() *types.CommonEntityData {
    traceBlocks.EntityData.YFilter = traceBlocks.YFilter
    traceBlocks.EntityData.YangName = "trace-blocks"
    traceBlocks.EntityData.BundleName = "cisco_ios_xr"
    traceBlocks.EntityData.ParentYangName = "all-options"
    traceBlocks.EntityData.SegmentPath = "trace-blocks"
    traceBlocks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBlocks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBlocks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBlocks.EntityData.Children = make(map[string]types.YChild)
    traceBlocks.EntityData.Leafs = make(map[string]types.YLeaf)
    traceBlocks.EntityData.Leafs["data"] = types.YLeaf{"Data", traceBlocks.Data}
    return &(traceBlocks.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Connections
type Controller_Switch_Oper_MgmtAgent_Connections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..4294967295.
    SdrNmNumClients interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdClientNumClients interface{}

    // The type is interface{} with range: 0..4294967295.
    MlapClientNumClients interface{}

    
    EsdmaInfo Controller_Switch_Oper_MgmtAgent_Connections_EsdmaInfo

    // The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Connections_SdrNmClientId.
    SdrNmClientId []Controller_Switch_Oper_MgmtAgent_Connections_SdrNmClientId

    // The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Connections_EsdClientId.
    EsdClientId []Controller_Switch_Oper_MgmtAgent_Connections_EsdClientId

    // The type is slice of
    // Controller_Switch_Oper_MgmtAgent_Connections_MlapClientId.
    MlapClientId []Controller_Switch_Oper_MgmtAgent_Connections_MlapClientId
}

func (connections *Controller_Switch_Oper_MgmtAgent_Connections) GetEntityData() *types.CommonEntityData {
    connections.EntityData.YFilter = connections.YFilter
    connections.EntityData.YangName = "connections"
    connections.EntityData.BundleName = "cisco_ios_xr"
    connections.EntityData.ParentYangName = "mgmt-agent"
    connections.EntityData.SegmentPath = "connections"
    connections.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connections.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connections.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connections.EntityData.Children = make(map[string]types.YChild)
    connections.EntityData.Children["esdma-info"] = types.YChild{"EsdmaInfo", &connections.EsdmaInfo}
    connections.EntityData.Children["sdr-nm-client-id"] = types.YChild{"SdrNmClientId", nil}
    for i := range connections.SdrNmClientId {
        connections.EntityData.Children[types.GetSegmentPath(&connections.SdrNmClientId[i])] = types.YChild{"SdrNmClientId", &connections.SdrNmClientId[i]}
    }
    connections.EntityData.Children["esd-client-id"] = types.YChild{"EsdClientId", nil}
    for i := range connections.EsdClientId {
        connections.EntityData.Children[types.GetSegmentPath(&connections.EsdClientId[i])] = types.YChild{"EsdClientId", &connections.EsdClientId[i]}
    }
    connections.EntityData.Children["mlap-client-id"] = types.YChild{"MlapClientId", nil}
    for i := range connections.MlapClientId {
        connections.EntityData.Children[types.GetSegmentPath(&connections.MlapClientId[i])] = types.YChild{"MlapClientId", &connections.MlapClientId[i]}
    }
    connections.EntityData.Leafs = make(map[string]types.YLeaf)
    connections.EntityData.Leafs["sdr-nm-num-clients"] = types.YLeaf{"SdrNmNumClients", connections.SdrNmNumClients}
    connections.EntityData.Leafs["esd-client-num-clients"] = types.YLeaf{"EsdClientNumClients", connections.EsdClientNumClients}
    connections.EntityData.Leafs["mlap-client-num-clients"] = types.YLeaf{"MlapClientNumClients", connections.MlapClientNumClients}
    return &(connections.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Connections_EsdmaInfo
type Controller_Switch_Oper_MgmtAgent_Connections_EsdmaInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    ActiveEsdmaAddress interface{}
}

func (esdmaInfo *Controller_Switch_Oper_MgmtAgent_Connections_EsdmaInfo) GetEntityData() *types.CommonEntityData {
    esdmaInfo.EntityData.YFilter = esdmaInfo.YFilter
    esdmaInfo.EntityData.YangName = "esdma-info"
    esdmaInfo.EntityData.BundleName = "cisco_ios_xr"
    esdmaInfo.EntityData.ParentYangName = "connections"
    esdmaInfo.EntityData.SegmentPath = "esdma-info"
    esdmaInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esdmaInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esdmaInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esdmaInfo.EntityData.Children = make(map[string]types.YChild)
    esdmaInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    esdmaInfo.EntityData.Leafs["active-esdma-address"] = types.YLeaf{"ActiveEsdmaAddress", esdmaInfo.ActiveEsdmaAddress}
    return &(esdmaInfo.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Connections_SdrNmClientId
type Controller_Switch_Oper_MgmtAgent_Connections_SdrNmClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SdrNmIpAddress interface{}

    // The type is interface{} with range: 0..65535.
    SdrNmPortNum interface{}
}

func (sdrNmClientId *Controller_Switch_Oper_MgmtAgent_Connections_SdrNmClientId) GetEntityData() *types.CommonEntityData {
    sdrNmClientId.EntityData.YFilter = sdrNmClientId.YFilter
    sdrNmClientId.EntityData.YangName = "sdr-nm-client-id"
    sdrNmClientId.EntityData.BundleName = "cisco_ios_xr"
    sdrNmClientId.EntityData.ParentYangName = "connections"
    sdrNmClientId.EntityData.SegmentPath = "sdr-nm-client-id"
    sdrNmClientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrNmClientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrNmClientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrNmClientId.EntityData.Children = make(map[string]types.YChild)
    sdrNmClientId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrNmClientId.EntityData.Leafs["sdr-nm-ip-address"] = types.YLeaf{"SdrNmIpAddress", sdrNmClientId.SdrNmIpAddress}
    sdrNmClientId.EntityData.Leafs["sdr-nm-port-num"] = types.YLeaf{"SdrNmPortNum", sdrNmClientId.SdrNmPortNum}
    return &(sdrNmClientId.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Connections_EsdClientId
type Controller_Switch_Oper_MgmtAgent_Connections_EsdClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    EsdClientIpAddress interface{}

    // The type is interface{} with range: 0..65535.
    EsdClientPortNum interface{}

    // The type is string.
    EsdClientLocation interface{}

    // The type is EsdmaSwitchYesNoEnum.
    EsdClientSwitchInfoCached interface{}

    // The type is EsdmaSwitchYesNoEnum.
    EsdClientSdrInfoCached interface{}
}

func (esdClientId *Controller_Switch_Oper_MgmtAgent_Connections_EsdClientId) GetEntityData() *types.CommonEntityData {
    esdClientId.EntityData.YFilter = esdClientId.YFilter
    esdClientId.EntityData.YangName = "esd-client-id"
    esdClientId.EntityData.BundleName = "cisco_ios_xr"
    esdClientId.EntityData.ParentYangName = "connections"
    esdClientId.EntityData.SegmentPath = "esd-client-id"
    esdClientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esdClientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esdClientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esdClientId.EntityData.Children = make(map[string]types.YChild)
    esdClientId.EntityData.Leafs = make(map[string]types.YLeaf)
    esdClientId.EntityData.Leafs["esd-client-ip-address"] = types.YLeaf{"EsdClientIpAddress", esdClientId.EsdClientIpAddress}
    esdClientId.EntityData.Leafs["esd-client-port-num"] = types.YLeaf{"EsdClientPortNum", esdClientId.EsdClientPortNum}
    esdClientId.EntityData.Leafs["esd-client-location"] = types.YLeaf{"EsdClientLocation", esdClientId.EsdClientLocation}
    esdClientId.EntityData.Leafs["esd-client-switch-info-cached"] = types.YLeaf{"EsdClientSwitchInfoCached", esdClientId.EsdClientSwitchInfoCached}
    esdClientId.EntityData.Leafs["esd-client-sdr-info-cached"] = types.YLeaf{"EsdClientSdrInfoCached", esdClientId.EsdClientSdrInfoCached}
    return &(esdClientId.EntityData)
}

// Controller_Switch_Oper_MgmtAgent_Connections_MlapClientId
type Controller_Switch_Oper_MgmtAgent_Connections_MlapClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    MlapClientIpAddress interface{}

    // The type is interface{} with range: 0..65535.
    MlapClientPortNum interface{}

    // The type is string.
    MlapClientLocation interface{}

    // The type is EsdmaSwitchYesNoEnum.
    MlapClientSwitchInfoCached interface{}
}

func (mlapClientId *Controller_Switch_Oper_MgmtAgent_Connections_MlapClientId) GetEntityData() *types.CommonEntityData {
    mlapClientId.EntityData.YFilter = mlapClientId.YFilter
    mlapClientId.EntityData.YangName = "mlap-client-id"
    mlapClientId.EntityData.BundleName = "cisco_ios_xr"
    mlapClientId.EntityData.ParentYangName = "connections"
    mlapClientId.EntityData.SegmentPath = "mlap-client-id"
    mlapClientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mlapClientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mlapClientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mlapClientId.EntityData.Children = make(map[string]types.YChild)
    mlapClientId.EntityData.Leafs = make(map[string]types.YLeaf)
    mlapClientId.EntityData.Leafs["mlap-client-ip-address"] = types.YLeaf{"MlapClientIpAddress", mlapClientId.MlapClientIpAddress}
    mlapClientId.EntityData.Leafs["mlap-client-port-num"] = types.YLeaf{"MlapClientPortNum", mlapClientId.MlapClientPortNum}
    mlapClientId.EntityData.Leafs["mlap-client-location"] = types.YLeaf{"MlapClientLocation", mlapClientId.MlapClientLocation}
    mlapClientId.EntityData.Leafs["mlap-client-switch-info-cached"] = types.YLeaf{"MlapClientSwitchInfoCached", mlapClientId.MlapClientSwitchInfoCached}
    return &(mlapClientId.EntityData)
}

// Controller_Switch_Oper_Sdr
type Controller_Switch_Oper_Sdr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    PortStatistics Controller_Switch_Oper_Sdr_PortStatistics

    
    GlobalStatistics Controller_Switch_Oper_Sdr_GlobalStatistics

    
    Policers Controller_Switch_Oper_Sdr_Policers

    
    SdrDetail Controller_Switch_Oper_Sdr_SdrDetail
}

func (sdr *Controller_Switch_Oper_Sdr) GetEntityData() *types.CommonEntityData {
    sdr.EntityData.YFilter = sdr.YFilter
    sdr.EntityData.YangName = "sdr"
    sdr.EntityData.BundleName = "cisco_ios_xr"
    sdr.EntityData.ParentYangName = "oper"
    sdr.EntityData.SegmentPath = "sdr"
    sdr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdr.EntityData.Children = make(map[string]types.YChild)
    sdr.EntityData.Children["port-statistics"] = types.YChild{"PortStatistics", &sdr.PortStatistics}
    sdr.EntityData.Children["global-statistics"] = types.YChild{"GlobalStatistics", &sdr.GlobalStatistics}
    sdr.EntityData.Children["policers"] = types.YChild{"Policers", &sdr.Policers}
    sdr.EntityData.Children["sdr-detail"] = types.YChild{"SdrDetail", &sdr.SdrDetail}
    sdr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdr.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics
type Controller_Switch_Oper_Sdr_PortStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Sdr_PortStatistics_Location.
    Location []Controller_Switch_Oper_Sdr_PortStatistics_Location
}

func (portStatistics *Controller_Switch_Oper_Sdr_PortStatistics) GetEntityData() *types.CommonEntityData {
    portStatistics.EntityData.YFilter = portStatistics.YFilter
    portStatistics.EntityData.YangName = "port-statistics"
    portStatistics.EntityData.BundleName = "cisco_ios_xr"
    portStatistics.EntityData.ParentYangName = "sdr"
    portStatistics.EntityData.SegmentPath = "port-statistics"
    portStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStatistics.EntityData.Children = make(map[string]types.YChild)
    portStatistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range portStatistics.Location {
        portStatistics.EntityData.Children[types.GetSegmentPath(&portStatistics.Location[i])] = types.YChild{"Location", &portStatistics.Location[i]}
    }
    portStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portStatistics.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics_Location
type Controller_Switch_Oper_Sdr_PortStatistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest to display the
    // SDR configuration for. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter.
    PortIter []Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter
}

func (location *Controller_Switch_Oper_Sdr_PortStatistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "port-statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter
type Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId.
    SdrId []Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId
}

func (portIter *Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range portIter.SdrId {
        portIter.EntityData.Children[types.GetSegmentPath(&portIter.SdrId[i])] = types.YChild{"SdrId", &portIter.SdrId[i]}
    }
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId
type Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    // The type is string.
    EsdSdrName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId.
    TrafficTypeId []Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId
}

func (sdrId *Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "port-iter"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["traffic-type-id"] = types.YChild{"TrafficTypeId", nil}
    for i := range sdrId.TrafficTypeId {
        sdrId.EntityData.Children[types.GetSegmentPath(&sdrId.TrafficTypeId[i])] = types.YChild{"TrafficTypeId", &sdrId.TrafficTypeId[i]}
    }
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    sdrId.EntityData.Leafs["esd-sdr-name"] = types.YLeaf{"EsdSdrName", sdrId.EsdSdrName}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId
type Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaSdrTrafficType.
    TrafficType interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId_DirectionId.
    DirectionId []Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId_DirectionId
}

func (trafficTypeId *Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId) GetEntityData() *types.CommonEntityData {
    trafficTypeId.EntityData.YFilter = trafficTypeId.YFilter
    trafficTypeId.EntityData.YangName = "traffic-type-id"
    trafficTypeId.EntityData.BundleName = "cisco_ios_xr"
    trafficTypeId.EntityData.ParentYangName = "sdr-id"
    trafficTypeId.EntityData.SegmentPath = "traffic-type-id" + "[traffic-type='" + fmt.Sprintf("%v", trafficTypeId.TrafficType) + "']"
    trafficTypeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficTypeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficTypeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficTypeId.EntityData.Children = make(map[string]types.YChild)
    trafficTypeId.EntityData.Children["direction-id"] = types.YChild{"DirectionId", nil}
    for i := range trafficTypeId.DirectionId {
        trafficTypeId.EntityData.Children[types.GetSegmentPath(&trafficTypeId.DirectionId[i])] = types.YChild{"DirectionId", &trafficTypeId.DirectionId[i]}
    }
    trafficTypeId.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficTypeId.EntityData.Leafs["traffic-type"] = types.YLeaf{"TrafficType", trafficTypeId.TrafficType}
    return &(trafficTypeId.EntityData)
}

// Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId_DirectionId
type Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId_DirectionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Packet direction this rule applies to. The type is
    // SwitchDataDirectionEnum.
    Direction interface{}

    // Counts packets within the committed information rate for all traffic
    // classes on this SDR. The type is interface{} with range:
    // 0..18446744073709551615.
    GreenPackets interface{}

    // Counts packets above the committed information rate, but within the excess
    // information rate for all traffic classes on this SDR. The type is
    // interface{} with range: 0..18446744073709551615.
    YellowPackets interface{}

    // Counts packets above the excess information rate for all traffic classes on
    // this SDR. Generally, these packets are dropped. The type is interface{}
    // with range: 0..18446744073709551615.
    RedPackets interface{}
}

func (directionId *Controller_Switch_Oper_Sdr_PortStatistics_Location_PortIter_SdrId_TrafficTypeId_DirectionId) GetEntityData() *types.CommonEntityData {
    directionId.EntityData.YFilter = directionId.YFilter
    directionId.EntityData.YangName = "direction-id"
    directionId.EntityData.BundleName = "cisco_ios_xr"
    directionId.EntityData.ParentYangName = "traffic-type-id"
    directionId.EntityData.SegmentPath = "direction-id" + "[direction='" + fmt.Sprintf("%v", directionId.Direction) + "']"
    directionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directionId.EntityData.Children = make(map[string]types.YChild)
    directionId.EntityData.Leafs = make(map[string]types.YLeaf)
    directionId.EntityData.Leafs["direction"] = types.YLeaf{"Direction", directionId.Direction}
    directionId.EntityData.Leafs["green-packets"] = types.YLeaf{"GreenPackets", directionId.GreenPackets}
    directionId.EntityData.Leafs["yellow-packets"] = types.YLeaf{"YellowPackets", directionId.YellowPackets}
    directionId.EntityData.Leafs["red-packets"] = types.YLeaf{"RedPackets", directionId.RedPackets}
    return &(directionId.EntityData)
}

// Controller_Switch_Oper_Sdr_GlobalStatistics
type Controller_Switch_Oper_Sdr_GlobalStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Sdr_GlobalStatistics_Location.
    Location []Controller_Switch_Oper_Sdr_GlobalStatistics_Location
}

func (globalStatistics *Controller_Switch_Oper_Sdr_GlobalStatistics) GetEntityData() *types.CommonEntityData {
    globalStatistics.EntityData.YFilter = globalStatistics.YFilter
    globalStatistics.EntityData.YangName = "global-statistics"
    globalStatistics.EntityData.BundleName = "cisco_ios_xr"
    globalStatistics.EntityData.ParentYangName = "sdr"
    globalStatistics.EntityData.SegmentPath = "global-statistics"
    globalStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalStatistics.EntityData.Children = make(map[string]types.YChild)
    globalStatistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range globalStatistics.Location {
        globalStatistics.EntityData.Children[types.GetSegmentPath(&globalStatistics.Location[i])] = types.YChild{"Location", &globalStatistics.Location[i]}
    }
    globalStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalStatistics.EntityData)
}

// Controller_Switch_Oper_Sdr_GlobalStatistics_Location
type Controller_Switch_Oper_Sdr_GlobalStatistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest to display the
    // SDR configuration for. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId.
    SdrId []Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId
}

func (location *Controller_Switch_Oper_Sdr_GlobalStatistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "global-statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range location.SdrId {
        location.EntityData.Children[types.GetSegmentPath(&location.SdrId[i])] = types.YChild{"SdrId", &location.SdrId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId
type Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    // The type is string.
    EsdSdrName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId.
    TrafficTypeId []Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId
}

func (sdrId *Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "location"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["traffic-type-id"] = types.YChild{"TrafficTypeId", nil}
    for i := range sdrId.TrafficTypeId {
        sdrId.EntityData.Children[types.GetSegmentPath(&sdrId.TrafficTypeId[i])] = types.YChild{"TrafficTypeId", &sdrId.TrafficTypeId[i]}
    }
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    sdrId.EntityData.Leafs["esd-sdr-name"] = types.YLeaf{"EsdSdrName", sdrId.EsdSdrName}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId
type Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaSdrTrafficType.
    TrafficType interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId_TrafficClassId.
    TrafficClassId []Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId_TrafficClassId
}

func (trafficTypeId *Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId) GetEntityData() *types.CommonEntityData {
    trafficTypeId.EntityData.YFilter = trafficTypeId.YFilter
    trafficTypeId.EntityData.YangName = "traffic-type-id"
    trafficTypeId.EntityData.BundleName = "cisco_ios_xr"
    trafficTypeId.EntityData.ParentYangName = "sdr-id"
    trafficTypeId.EntityData.SegmentPath = "traffic-type-id" + "[traffic-type='" + fmt.Sprintf("%v", trafficTypeId.TrafficType) + "']"
    trafficTypeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficTypeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficTypeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficTypeId.EntityData.Children = make(map[string]types.YChild)
    trafficTypeId.EntityData.Children["traffic-class-id"] = types.YChild{"TrafficClassId", nil}
    for i := range trafficTypeId.TrafficClassId {
        trafficTypeId.EntityData.Children[types.GetSegmentPath(&trafficTypeId.TrafficClassId[i])] = types.YChild{"TrafficClassId", &trafficTypeId.TrafficClassId[i]}
    }
    trafficTypeId.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficTypeId.EntityData.Leafs["traffic-type"] = types.YLeaf{"TrafficType", trafficTypeId.TrafficType}
    return &(trafficTypeId.EntityData)
}

// Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId_TrafficClassId
type Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId_TrafficClassId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Traffic class (0->7) for these statistics. The
    // type is interface{} with range: -1..7.
    Tc interface{}

    // Counts packets within the committed information rate for a traffic class on
    // this VLAN. The type is interface{} with range: 0..18446744073709551615.
    GreenPackets interface{}

    // Counts packets above the committed information rate, but within the excess
    // information rate for a traffic class on this VLAN. The type is interface{}
    // with range: 0..18446744073709551615.
    YellowPackets interface{}

    // Counts packets above the excess information rate for a traffic class on
    // this VLAN. Generally, these packets are dropped. The type is interface{}
    // with range: 0..18446744073709551615.
    RedPackets interface{}
}

func (trafficClassId *Controller_Switch_Oper_Sdr_GlobalStatistics_Location_SdrId_TrafficTypeId_TrafficClassId) GetEntityData() *types.CommonEntityData {
    trafficClassId.EntityData.YFilter = trafficClassId.YFilter
    trafficClassId.EntityData.YangName = "traffic-class-id"
    trafficClassId.EntityData.BundleName = "cisco_ios_xr"
    trafficClassId.EntityData.ParentYangName = "traffic-type-id"
    trafficClassId.EntityData.SegmentPath = "traffic-class-id" + "[tc='" + fmt.Sprintf("%v", trafficClassId.Tc) + "']"
    trafficClassId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficClassId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficClassId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficClassId.EntityData.Children = make(map[string]types.YChild)
    trafficClassId.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficClassId.EntityData.Leafs["tc"] = types.YLeaf{"Tc", trafficClassId.Tc}
    trafficClassId.EntityData.Leafs["green-packets"] = types.YLeaf{"GreenPackets", trafficClassId.GreenPackets}
    trafficClassId.EntityData.Leafs["yellow-packets"] = types.YLeaf{"YellowPackets", trafficClassId.YellowPackets}
    trafficClassId.EntityData.Leafs["red-packets"] = types.YLeaf{"RedPackets", trafficClassId.RedPackets}
    return &(trafficClassId.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers
type Controller_Switch_Oper_Sdr_Policers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Sdr_Policers_Location.
    Location []Controller_Switch_Oper_Sdr_Policers_Location
}

func (policers *Controller_Switch_Oper_Sdr_Policers) GetEntityData() *types.CommonEntityData {
    policers.EntityData.YFilter = policers.YFilter
    policers.EntityData.YangName = "policers"
    policers.EntityData.BundleName = "cisco_ios_xr"
    policers.EntityData.ParentYangName = "sdr"
    policers.EntityData.SegmentPath = "policers"
    policers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policers.EntityData.Children = make(map[string]types.YChild)
    policers.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range policers.Location {
        policers.EntityData.Children[types.GetSegmentPath(&policers.Location[i])] = types.YChild{"Location", &policers.Location[i]}
    }
    policers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(policers.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers_Location
type Controller_Switch_Oper_Sdr_Policers_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    
    EsdPolicerStatus Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus

    // The type is slice of Controller_Switch_Oper_Sdr_Policers_Location_SdrId.
    SdrId []Controller_Switch_Oper_Sdr_Policers_Location_SdrId
}

func (location *Controller_Switch_Oper_Sdr_Policers_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "policers"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["esd-policer-status"] = types.YChild{"EsdPolicerStatus", &location.EsdPolicerStatus}
    location.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range location.SdrId {
        location.EntityData.Children[types.GetSegmentPath(&location.SdrId[i])] = types.YChild{"SdrId", &location.SdrId[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus
type Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    IndentGroup Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus_IndentGroup
}

func (esdPolicerStatus *Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus) GetEntityData() *types.CommonEntityData {
    esdPolicerStatus.EntityData.YFilter = esdPolicerStatus.YFilter
    esdPolicerStatus.EntityData.YangName = "esd-policer-status"
    esdPolicerStatus.EntityData.BundleName = "cisco_ios_xr"
    esdPolicerStatus.EntityData.ParentYangName = "location"
    esdPolicerStatus.EntityData.SegmentPath = "esd-policer-status"
    esdPolicerStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esdPolicerStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esdPolicerStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esdPolicerStatus.EntityData.Children = make(map[string]types.YChild)
    esdPolicerStatus.EntityData.Children["indent-group"] = types.YChild{"IndentGroup", &esdPolicerStatus.IndentGroup}
    esdPolicerStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(esdPolicerStatus.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus_IndentGroup
type Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus_IndentGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is EsdmaSwitchYesNoEnum.
    EsdPortPolicingEnabled interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdPortCommittedBurstSize interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdPortPeakBurstSize interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdPortPolicerMru interface{}

    // The type is EsdmaSwitchYesNoEnum.
    EsdGlobalPolicingEnabled interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdGlobalCommittedBurstSize interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdGlobalPeakBurstSize interface{}

    // The type is interface{} with range: 0..4294967295.
    EsdGlobalPolicerMru interface{}
}

func (indentGroup *Controller_Switch_Oper_Sdr_Policers_Location_EsdPolicerStatus_IndentGroup) GetEntityData() *types.CommonEntityData {
    indentGroup.EntityData.YFilter = indentGroup.YFilter
    indentGroup.EntityData.YangName = "indent-group"
    indentGroup.EntityData.BundleName = "cisco_ios_xr"
    indentGroup.EntityData.ParentYangName = "esd-policer-status"
    indentGroup.EntityData.SegmentPath = "indent-group"
    indentGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indentGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indentGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indentGroup.EntityData.Children = make(map[string]types.YChild)
    indentGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    indentGroup.EntityData.Leafs["esd-port-policing-enabled"] = types.YLeaf{"EsdPortPolicingEnabled", indentGroup.EsdPortPolicingEnabled}
    indentGroup.EntityData.Leafs["esd-port-committed-burst-size"] = types.YLeaf{"EsdPortCommittedBurstSize", indentGroup.EsdPortCommittedBurstSize}
    indentGroup.EntityData.Leafs["esd-port-peak-burst-size"] = types.YLeaf{"EsdPortPeakBurstSize", indentGroup.EsdPortPeakBurstSize}
    indentGroup.EntityData.Leafs["esd-port-policer-mru"] = types.YLeaf{"EsdPortPolicerMru", indentGroup.EsdPortPolicerMru}
    indentGroup.EntityData.Leafs["esd-global-policing-enabled"] = types.YLeaf{"EsdGlobalPolicingEnabled", indentGroup.EsdGlobalPolicingEnabled}
    indentGroup.EntityData.Leafs["esd-global-committed-burst-size"] = types.YLeaf{"EsdGlobalCommittedBurstSize", indentGroup.EsdGlobalCommittedBurstSize}
    indentGroup.EntityData.Leafs["esd-global-peak-burst-size"] = types.YLeaf{"EsdGlobalPeakBurstSize", indentGroup.EsdGlobalPeakBurstSize}
    indentGroup.EntityData.Leafs["esd-global-policer-mru"] = types.YLeaf{"EsdGlobalPolicerMru", indentGroup.EsdGlobalPolicerMru}
    return &(indentGroup.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers_Location_SdrId
type Controller_Switch_Oper_Sdr_Policers_Location_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    // The type is string.
    EsdSdrName interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCir interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrPir interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_Policers_Location_SdrId_EsdSdrCosTypeIter.
    EsdSdrCosTypeIter []Controller_Switch_Oper_Sdr_Policers_Location_SdrId_EsdSdrCosTypeIter
}

func (sdrId *Controller_Switch_Oper_Sdr_Policers_Location_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "location"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["esd-sdr-cos-type-iter"] = types.YChild{"EsdSdrCosTypeIter", nil}
    for i := range sdrId.EsdSdrCosTypeIter {
        sdrId.EntityData.Children[types.GetSegmentPath(&sdrId.EsdSdrCosTypeIter[i])] = types.YChild{"EsdSdrCosTypeIter", &sdrId.EsdSdrCosTypeIter[i]}
    }
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    sdrId.EntityData.Leafs["esd-sdr-name"] = types.YLeaf{"EsdSdrName", sdrId.EsdSdrName}
    sdrId.EntityData.Leafs["esd-sdr-cir"] = types.YLeaf{"EsdSdrCir", sdrId.EsdSdrCir}
    sdrId.EntityData.Leafs["esd-sdr-pir"] = types.YLeaf{"EsdSdrPir", sdrId.EsdSdrPir}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Sdr_Policers_Location_SdrId_EsdSdrCosTypeIter
type Controller_Switch_Oper_Sdr_Policers_Location_SdrId_EsdSdrCosTypeIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdCirEirType.
    EsdSdrCosType interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos0 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos1 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos2 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos3 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos4 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos5 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos6 interface{}

    // The type is interface{} with range: 0..100.
    EsdSdrCos7 interface{}
}

func (esdSdrCosTypeIter *Controller_Switch_Oper_Sdr_Policers_Location_SdrId_EsdSdrCosTypeIter) GetEntityData() *types.CommonEntityData {
    esdSdrCosTypeIter.EntityData.YFilter = esdSdrCosTypeIter.YFilter
    esdSdrCosTypeIter.EntityData.YangName = "esd-sdr-cos-type-iter"
    esdSdrCosTypeIter.EntityData.BundleName = "cisco_ios_xr"
    esdSdrCosTypeIter.EntityData.ParentYangName = "sdr-id"
    esdSdrCosTypeIter.EntityData.SegmentPath = "esd-sdr-cos-type-iter" + "[esd-sdr-cos-type='" + fmt.Sprintf("%v", esdSdrCosTypeIter.EsdSdrCosType) + "']"
    esdSdrCosTypeIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    esdSdrCosTypeIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    esdSdrCosTypeIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    esdSdrCosTypeIter.EntityData.Children = make(map[string]types.YChild)
    esdSdrCosTypeIter.EntityData.Leafs = make(map[string]types.YLeaf)
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-type"] = types.YLeaf{"EsdSdrCosType", esdSdrCosTypeIter.EsdSdrCosType}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-0"] = types.YLeaf{"EsdSdrCos0", esdSdrCosTypeIter.EsdSdrCos0}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-1"] = types.YLeaf{"EsdSdrCos1", esdSdrCosTypeIter.EsdSdrCos1}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-2"] = types.YLeaf{"EsdSdrCos2", esdSdrCosTypeIter.EsdSdrCos2}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-3"] = types.YLeaf{"EsdSdrCos3", esdSdrCosTypeIter.EsdSdrCos3}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-4"] = types.YLeaf{"EsdSdrCos4", esdSdrCosTypeIter.EsdSdrCos4}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-5"] = types.YLeaf{"EsdSdrCos5", esdSdrCosTypeIter.EsdSdrCos5}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-6"] = types.YLeaf{"EsdSdrCos6", esdSdrCosTypeIter.EsdSdrCos6}
    esdSdrCosTypeIter.EntityData.Leafs["esd-sdr-cos-7"] = types.YLeaf{"EsdSdrCos7", esdSdrCosTypeIter.EsdSdrCos7}
    return &(esdSdrCosTypeIter.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail
type Controller_Switch_Oper_Sdr_SdrDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Sdr_SdrDetail_SdrId.
    SdrId []Controller_Switch_Oper_Sdr_SdrDetail_SdrId
}

func (sdrDetail *Controller_Switch_Oper_Sdr_SdrDetail) GetEntityData() *types.CommonEntityData {
    sdrDetail.EntityData.YFilter = sdrDetail.YFilter
    sdrDetail.EntityData.YangName = "sdr-detail"
    sdrDetail.EntityData.BundleName = "cisco_ios_xr"
    sdrDetail.EntityData.ParentYangName = "sdr"
    sdrDetail.EntityData.SegmentPath = "sdr-detail"
    sdrDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrDetail.EntityData.Children = make(map[string]types.YChild)
    sdrDetail.EntityData.Children["sdr-id"] = types.YChild{"SdrId", nil}
    for i := range sdrDetail.SdrId {
        sdrDetail.EntityData.Children[types.GetSegmentPath(&sdrDetail.SdrId[i])] = types.YChild{"SdrId", &sdrDetail.SdrId[i]}
    }
    sdrDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sdrDetail.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Sdr interface{}

    
    PortStatistics Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics
}

func (sdrId *Controller_Switch_Oper_Sdr_SdrDetail_SdrId) GetEntityData() *types.CommonEntityData {
    sdrId.EntityData.YFilter = sdrId.YFilter
    sdrId.EntityData.YangName = "sdr-id"
    sdrId.EntityData.BundleName = "cisco_ios_xr"
    sdrId.EntityData.ParentYangName = "sdr-detail"
    sdrId.EntityData.SegmentPath = "sdr-id" + "[sdr='" + fmt.Sprintf("%v", sdrId.Sdr) + "']"
    sdrId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sdrId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sdrId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sdrId.EntityData.Children = make(map[string]types.YChild)
    sdrId.EntityData.Children["port-statistics"] = types.YChild{"PortStatistics", &sdrId.PortStatistics}
    sdrId.EntityData.Leafs = make(map[string]types.YLeaf)
    sdrId.EntityData.Leafs["sdr"] = types.YLeaf{"Sdr", sdrId.Sdr}
    return &(sdrId.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location.
    Location []Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location
}

func (portStatistics *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics) GetEntityData() *types.CommonEntityData {
    portStatistics.EntityData.YFilter = portStatistics.YFilter
    portStatistics.EntityData.YangName = "port-statistics"
    portStatistics.EntityData.BundleName = "cisco_ios_xr"
    portStatistics.EntityData.ParentYangName = "sdr-id"
    portStatistics.EntityData.SegmentPath = "port-statistics"
    portStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portStatistics.EntityData.Children = make(map[string]types.YChild)
    portStatistics.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range portStatistics.Location {
        portStatistics.EntityData.Children[types.GetSegmentPath(&portStatistics.Location[i])] = types.YChild{"Location", &portStatistics.Location[i]}
    }
    portStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portStatistics.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter.
    PortIter []Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter
}

func (location *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "port-statistics"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId.
    TrafficTypeId []Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId
}

func (portIter *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Children["traffic-type-id"] = types.YChild{"TrafficTypeId", nil}
    for i := range portIter.TrafficTypeId {
        portIter.EntityData.Children[types.GetSegmentPath(&portIter.TrafficTypeId[i])] = types.YChild{"TrafficTypeId", &portIter.TrafficTypeId[i]}
    }
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaSdrTrafficType.
    TrafficType interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId.
    DirectionId []Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId
}

func (trafficTypeId *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId) GetEntityData() *types.CommonEntityData {
    trafficTypeId.EntityData.YFilter = trafficTypeId.YFilter
    trafficTypeId.EntityData.YangName = "traffic-type-id"
    trafficTypeId.EntityData.BundleName = "cisco_ios_xr"
    trafficTypeId.EntityData.ParentYangName = "port-iter"
    trafficTypeId.EntityData.SegmentPath = "traffic-type-id" + "[traffic-type='" + fmt.Sprintf("%v", trafficTypeId.TrafficType) + "']"
    trafficTypeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficTypeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficTypeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficTypeId.EntityData.Children = make(map[string]types.YChild)
    trafficTypeId.EntityData.Children["direction-id"] = types.YChild{"DirectionId", nil}
    for i := range trafficTypeId.DirectionId {
        trafficTypeId.EntityData.Children[types.GetSegmentPath(&trafficTypeId.DirectionId[i])] = types.YChild{"DirectionId", &trafficTypeId.DirectionId[i]}
    }
    trafficTypeId.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficTypeId.EntityData.Leafs["traffic-type"] = types.YLeaf{"TrafficType", trafficTypeId.TrafficType}
    return &(trafficTypeId.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Packet direction this rule applies to. The type is
    // SwitchDataDirectionEnum.
    Direction interface{}

    // The type is slice of
    // Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId_TrafficClassId.
    TrafficClassId []Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId_TrafficClassId
}

func (directionId *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId) GetEntityData() *types.CommonEntityData {
    directionId.EntityData.YFilter = directionId.YFilter
    directionId.EntityData.YangName = "direction-id"
    directionId.EntityData.BundleName = "cisco_ios_xr"
    directionId.EntityData.ParentYangName = "traffic-type-id"
    directionId.EntityData.SegmentPath = "direction-id" + "[direction='" + fmt.Sprintf("%v", directionId.Direction) + "']"
    directionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    directionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    directionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    directionId.EntityData.Children = make(map[string]types.YChild)
    directionId.EntityData.Children["traffic-class-id"] = types.YChild{"TrafficClassId", nil}
    for i := range directionId.TrafficClassId {
        directionId.EntityData.Children[types.GetSegmentPath(&directionId.TrafficClassId[i])] = types.YChild{"TrafficClassId", &directionId.TrafficClassId[i]}
    }
    directionId.EntityData.Leafs = make(map[string]types.YLeaf)
    directionId.EntityData.Leafs["direction"] = types.YLeaf{"Direction", directionId.Direction}
    return &(directionId.EntityData)
}

// Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId_TrafficClassId
type Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId_TrafficClassId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Traffic class (0->7) for these statistics. The
    // type is interface{} with range: -1..7.
    Tc interface{}

    // Counts packets within the committed information rate for a traffic class on
    // this VLAN. The type is interface{} with range: 0..18446744073709551615.
    GreenPackets interface{}

    // Counts packets above the committed information rate, but within the excess
    // information rate for a traffic class on this VLAN. The type is interface{}
    // with range: 0..18446744073709551615.
    YellowPackets interface{}

    // Counts packets above the excess information rate for a traffic class on
    // this VLAN. Generally, these packets are dropped. The type is interface{}
    // with range: 0..18446744073709551615.
    RedPackets interface{}
}

func (trafficClassId *Controller_Switch_Oper_Sdr_SdrDetail_SdrId_PortStatistics_Location_PortIter_TrafficTypeId_DirectionId_TrafficClassId) GetEntityData() *types.CommonEntityData {
    trafficClassId.EntityData.YFilter = trafficClassId.YFilter
    trafficClassId.EntityData.YangName = "traffic-class-id"
    trafficClassId.EntityData.BundleName = "cisco_ios_xr"
    trafficClassId.EntityData.ParentYangName = "direction-id"
    trafficClassId.EntityData.SegmentPath = "traffic-class-id" + "[tc='" + fmt.Sprintf("%v", trafficClassId.Tc) + "']"
    trafficClassId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trafficClassId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trafficClassId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trafficClassId.EntityData.Children = make(map[string]types.YChild)
    trafficClassId.EntityData.Leafs = make(map[string]types.YLeaf)
    trafficClassId.EntityData.Leafs["tc"] = types.YLeaf{"Tc", trafficClassId.Tc}
    trafficClassId.EntityData.Leafs["green-packets"] = types.YLeaf{"GreenPackets", trafficClassId.GreenPackets}
    trafficClassId.EntityData.Leafs["yellow-packets"] = types.YLeaf{"YellowPackets", trafficClassId.YellowPackets}
    trafficClassId.EntityData.Leafs["red-packets"] = types.YLeaf{"RedPackets", trafficClassId.RedPackets}
    return &(trafficClassId.EntityData)
}

// Controller_Switch_Oper_PortState
type Controller_Switch_Oper_PortState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_PortState_Location.
    Location []Controller_Switch_Oper_PortState_Location
}

func (portState *Controller_Switch_Oper_PortState) GetEntityData() *types.CommonEntityData {
    portState.EntityData.YFilter = portState.YFilter
    portState.EntityData.YangName = "port-state"
    portState.EntityData.BundleName = "cisco_ios_xr"
    portState.EntityData.ParentYangName = "oper"
    portState.EntityData.SegmentPath = "port-state"
    portState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portState.EntityData.Children = make(map[string]types.YChild)
    portState.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range portState.Location {
        portState.EntityData.Children[types.GetSegmentPath(&portState.Location[i])] = types.YChild{"Location", &portState.Location[i]}
    }
    portState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(portState.EntityData)
}

// Controller_Switch_Oper_PortState_Location
type Controller_Switch_Oper_PortState_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card that owns a switch of interest. The type is
    // EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // The type is slice of Controller_Switch_Oper_PortState_Location_PortIter.
    PortIter []Controller_Switch_Oper_PortState_Location_PortIter
}

func (location *Controller_Switch_Oper_PortState_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "port-state"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["port-iter"] = types.YChild{"PortIter", nil}
    for i := range location.PortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.PortIter[i])] = types.YChild{"PortIter", &location.PortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_PortState_Location_PortIter
type Controller_Switch_Oper_PortState_Location_PortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switch port. The type is interface{} with range:
    // 0..127.
    Port interface{}
}

func (portIter *Controller_Switch_Oper_PortState_Location_PortIter) GetEntityData() *types.CommonEntityData {
    portIter.EntityData.YFilter = portIter.YFilter
    portIter.EntityData.YangName = "port-iter"
    portIter.EntityData.BundleName = "cisco_ios_xr"
    portIter.EntityData.ParentYangName = "location"
    portIter.EntityData.SegmentPath = "port-iter" + "[port='" + fmt.Sprintf("%v", portIter.Port) + "']"
    portIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    portIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    portIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    portIter.EntityData.Children = make(map[string]types.YChild)
    portIter.EntityData.Leafs = make(map[string]types.YLeaf)
    portIter.EntityData.Leafs["port"] = types.YLeaf{"Port", portIter.Port}
    return &(portIter.EntityData)
}

// Controller_Switch_Oper_Trunk
type Controller_Switch_Oper_Trunk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Controller_Switch_Oper_Trunk_Location.
    Location []Controller_Switch_Oper_Trunk_Location
}

func (trunk *Controller_Switch_Oper_Trunk) GetEntityData() *types.CommonEntityData {
    trunk.EntityData.YFilter = trunk.YFilter
    trunk.EntityData.YangName = "trunk"
    trunk.EntityData.BundleName = "cisco_ios_xr"
    trunk.EntityData.ParentYangName = "oper"
    trunk.EntityData.SegmentPath = "trunk"
    trunk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunk.EntityData.Children = make(map[string]types.YChild)
    trunk.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range trunk.Location {
        trunk.EntityData.Children[types.GetSegmentPath(&trunk.Location[i])] = types.YChild{"Location", &trunk.Location[i]}
    }
    trunk.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trunk.EntityData)
}

// Controller_Switch_Oper_Trunk_Location
type Controller_Switch_Oper_Trunk_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rack to display the switch trunk group information
    // for. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. Card to display the switch trunk group information
    // for. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. Switch type to display the switch trunk group
    // information for. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    // Number of member ports in the trunk group. The type is interface{} with
    // range: 0..4294967295.
    TrunkMemberCount interface{}

    // Name of the trunk group. The type is string.
    TrunkName interface{}

    // The type is slice of
    // Controller_Switch_Oper_Trunk_Location_TrunkMemberPortIter.
    TrunkMemberPortIter []Controller_Switch_Oper_Trunk_Location_TrunkMemberPortIter
}

func (location *Controller_Switch_Oper_Trunk_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "trunk"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["trunk-member-port-iter"] = types.YChild{"TrunkMemberPortIter", nil}
    for i := range location.TrunkMemberPortIter {
        location.EntityData.Children[types.GetSegmentPath(&location.TrunkMemberPortIter[i])] = types.YChild{"TrunkMemberPortIter", &location.TrunkMemberPortIter[i]}
    }
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    location.EntityData.Leafs["trunk-member-count"] = types.YLeaf{"TrunkMemberCount", location.TrunkMemberCount}
    location.EntityData.Leafs["trunk-name"] = types.YLeaf{"TrunkName", location.TrunkName}
    return &(location.EntityData)
}

// Controller_Switch_Oper_Trunk_Location_TrunkMemberPortIter
type Controller_Switch_Oper_Trunk_Location_TrunkMemberPortIter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Trunk Member Port. The type is interface{} with
    // range: 0..4294967295.
    TrunkMemberPort interface{}

    // Trunk Member Status. The type is EsdmaTrunkMemberStatus.
    TrunkMemberStatus interface{}
}

func (trunkMemberPortIter *Controller_Switch_Oper_Trunk_Location_TrunkMemberPortIter) GetEntityData() *types.CommonEntityData {
    trunkMemberPortIter.EntityData.YFilter = trunkMemberPortIter.YFilter
    trunkMemberPortIter.EntityData.YangName = "trunk-member-port-iter"
    trunkMemberPortIter.EntityData.BundleName = "cisco_ios_xr"
    trunkMemberPortIter.EntityData.ParentYangName = "location"
    trunkMemberPortIter.EntityData.SegmentPath = "trunk-member-port-iter" + "[trunk-member-port='" + fmt.Sprintf("%v", trunkMemberPortIter.TrunkMemberPort) + "']"
    trunkMemberPortIter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trunkMemberPortIter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trunkMemberPortIter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trunkMemberPortIter.EntityData.Children = make(map[string]types.YChild)
    trunkMemberPortIter.EntityData.Leafs = make(map[string]types.YLeaf)
    trunkMemberPortIter.EntityData.Leafs["trunk-member-port"] = types.YLeaf{"TrunkMemberPort", trunkMemberPortIter.TrunkMemberPort}
    trunkMemberPortIter.EntityData.Leafs["trunk-member-status"] = types.YLeaf{"TrunkMemberStatus", trunkMemberPortIter.TrunkMemberStatus}
    return &(trunkMemberPortIter.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont
type Controller_Switch_Oper_SwitchDebugCont struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Debug Controller_Switch_Oper_SwitchDebugCont_Debug
}

func (switchDebugCont *Controller_Switch_Oper_SwitchDebugCont) GetEntityData() *types.CommonEntityData {
    switchDebugCont.EntityData.YFilter = switchDebugCont.YFilter
    switchDebugCont.EntityData.YangName = "switch-debug-cont"
    switchDebugCont.EntityData.BundleName = "cisco_ios_xr"
    switchDebugCont.EntityData.ParentYangName = "oper"
    switchDebugCont.EntityData.SegmentPath = "switch-debug-cont"
    switchDebugCont.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchDebugCont.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchDebugCont.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchDebugCont.EntityData.Children = make(map[string]types.YChild)
    switchDebugCont.EntityData.Children["debug"] = types.YChild{"Debug", &switchDebugCont.Debug}
    switchDebugCont.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(switchDebugCont.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont_Debug
type Controller_Switch_Oper_SwitchDebugCont_Debug struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Counters Controller_Switch_Oper_SwitchDebugCont_Debug_Counters
}

func (debug *Controller_Switch_Oper_SwitchDebugCont_Debug) GetEntityData() *types.CommonEntityData {
    debug.EntityData.YFilter = debug.YFilter
    debug.EntityData.YangName = "debug"
    debug.EntityData.BundleName = "cisco_ios_xr"
    debug.EntityData.ParentYangName = "switch-debug-cont"
    debug.EntityData.SegmentPath = "debug"
    debug.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    debug.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    debug.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    debug.EntityData.Children = make(map[string]types.YChild)
    debug.EntityData.Children["counters"] = types.YChild{"Counters", &debug.Counters}
    debug.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(debug.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont_Debug_Counters
type Controller_Switch_Oper_SwitchDebugCont_Debug_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location.
    Location []Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location
}

func (counters *Controller_Switch_Oper_SwitchDebugCont_Debug_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "debug"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = make(map[string]types.YChild)
    counters.EntityData.Children["location"] = types.YChild{"Location", nil}
    for i := range counters.Location {
        counters.EntityData.Children[types.GetSegmentPath(&counters.Location[i])] = types.YChild{"Location", &counters.Location[i]}
    }
    counters.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(counters.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location
type Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is EsdmaRackNumEnum.
    Rack interface{}

    // This attribute is a key. The type is EsdmaCpu.
    Card interface{}

    // This attribute is a key. The type is EsdmaSwitchTypeEnum.
    SwitchId interface{}

    
    Counters Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters_
}

func (location *Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location) GetEntityData() *types.CommonEntityData {
    location.EntityData.YFilter = location.YFilter
    location.EntityData.YangName = "location"
    location.EntityData.BundleName = "cisco_ios_xr"
    location.EntityData.ParentYangName = "counters"
    location.EntityData.SegmentPath = "location" + "[rack='" + fmt.Sprintf("%v", location.Rack) + "']" + "[card='" + fmt.Sprintf("%v", location.Card) + "']" + "[switch-id='" + fmt.Sprintf("%v", location.SwitchId) + "']"
    location.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    location.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    location.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    location.EntityData.Children = make(map[string]types.YChild)
    location.EntityData.Children["counters"] = types.YChild{"Counters", &location.Counters}
    location.EntityData.Leafs = make(map[string]types.YLeaf)
    location.EntityData.Leafs["rack"] = types.YLeaf{"Rack", location.Rack}
    location.EntityData.Leafs["card"] = types.YLeaf{"Card", location.Card}
    location.EntityData.Leafs["switch-id"] = types.YLeaf{"SwitchId", location.SwitchId}
    return &(location.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters_
type Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is bool.
    PhyPollingEnabled interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TxThreadWdogCnt interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    RxThreadWdogCnt interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TaskLockLongestWaitTime interface{}

    // The type is interface{} with range: 0..4294967295.
    TaskLockLongestWaitEvent interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TaskLockLongestHeldTime interface{}

    // The type is interface{} with range: 0..4294967295.
    TaskLockLongestHeldEvent interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    TaskUnlockLongestWaitTime interface{}

    // The type is interface{} with range: 0..4294967295.
    TaskUnlockLongestWaitEvent interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaMaxRxDequeuedPerInt interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaRxPacketsDequeued interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaRxPacketDequeueErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketsQueued interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketsCompleted interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketNoMsgErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketMsgTooBigErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketNoBufferErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketQueueErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxPacketCompletionErrors interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaMaxTxFreedPerInt interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxBufAllocCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    DmaTxBufFreeCount interface{}

    // The type is slice of
    // Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters__SwitchCore.
    SwitchCore []Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters__SwitchCore
}

func (counters_ *Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters_) GetEntityData() *types.CommonEntityData {
    counters_.EntityData.YFilter = counters_.YFilter
    counters_.EntityData.YangName = "counters"
    counters_.EntityData.BundleName = "cisco_ios_xr"
    counters_.EntityData.ParentYangName = "location"
    counters_.EntityData.SegmentPath = "counters"
    counters_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters_.EntityData.Children = make(map[string]types.YChild)
    counters_.EntityData.Children["switch-core"] = types.YChild{"SwitchCore", nil}
    for i := range counters_.SwitchCore {
        counters_.EntityData.Children[types.GetSegmentPath(&counters_.SwitchCore[i])] = types.YChild{"SwitchCore", &counters_.SwitchCore[i]}
    }
    counters_.EntityData.Leafs = make(map[string]types.YLeaf)
    counters_.EntityData.Leafs["phy-polling-enabled"] = types.YLeaf{"PhyPollingEnabled", counters_.PhyPollingEnabled}
    counters_.EntityData.Leafs["tx-thread-wdog-cnt"] = types.YLeaf{"TxThreadWdogCnt", counters_.TxThreadWdogCnt}
    counters_.EntityData.Leafs["rx-thread-wdog-cnt"] = types.YLeaf{"RxThreadWdogCnt", counters_.RxThreadWdogCnt}
    counters_.EntityData.Leafs["task-lock-longest-wait-time"] = types.YLeaf{"TaskLockLongestWaitTime", counters_.TaskLockLongestWaitTime}
    counters_.EntityData.Leafs["task-lock-longest-wait-event"] = types.YLeaf{"TaskLockLongestWaitEvent", counters_.TaskLockLongestWaitEvent}
    counters_.EntityData.Leafs["task-lock-longest-held-time"] = types.YLeaf{"TaskLockLongestHeldTime", counters_.TaskLockLongestHeldTime}
    counters_.EntityData.Leafs["task-lock-longest-held-event"] = types.YLeaf{"TaskLockLongestHeldEvent", counters_.TaskLockLongestHeldEvent}
    counters_.EntityData.Leafs["task-unlock-longest-wait-time"] = types.YLeaf{"TaskUnlockLongestWaitTime", counters_.TaskUnlockLongestWaitTime}
    counters_.EntityData.Leafs["task-unlock-longest-wait-event"] = types.YLeaf{"TaskUnlockLongestWaitEvent", counters_.TaskUnlockLongestWaitEvent}
    counters_.EntityData.Leafs["dma-max-rx-dequeued-per-int"] = types.YLeaf{"DmaMaxRxDequeuedPerInt", counters_.DmaMaxRxDequeuedPerInt}
    counters_.EntityData.Leafs["dma-rx-packets-dequeued"] = types.YLeaf{"DmaRxPacketsDequeued", counters_.DmaRxPacketsDequeued}
    counters_.EntityData.Leafs["dma-rx-packet-dequeue-errors"] = types.YLeaf{"DmaRxPacketDequeueErrors", counters_.DmaRxPacketDequeueErrors}
    counters_.EntityData.Leafs["dma-tx-packets-queued"] = types.YLeaf{"DmaTxPacketsQueued", counters_.DmaTxPacketsQueued}
    counters_.EntityData.Leafs["dma-tx-packets-completed"] = types.YLeaf{"DmaTxPacketsCompleted", counters_.DmaTxPacketsCompleted}
    counters_.EntityData.Leafs["dma-tx-packet-no-msg-errors"] = types.YLeaf{"DmaTxPacketNoMsgErrors", counters_.DmaTxPacketNoMsgErrors}
    counters_.EntityData.Leafs["dma-tx-packet-msg-too-big-errors"] = types.YLeaf{"DmaTxPacketMsgTooBigErrors", counters_.DmaTxPacketMsgTooBigErrors}
    counters_.EntityData.Leafs["dma-tx-packet-no-buffer-errors"] = types.YLeaf{"DmaTxPacketNoBufferErrors", counters_.DmaTxPacketNoBufferErrors}
    counters_.EntityData.Leafs["dma-tx-packet-queue-errors"] = types.YLeaf{"DmaTxPacketQueueErrors", counters_.DmaTxPacketQueueErrors}
    counters_.EntityData.Leafs["dma-tx-packet-completion-errors"] = types.YLeaf{"DmaTxPacketCompletionErrors", counters_.DmaTxPacketCompletionErrors}
    counters_.EntityData.Leafs["dma-max-tx-freed-per-int"] = types.YLeaf{"DmaMaxTxFreedPerInt", counters_.DmaMaxTxFreedPerInt}
    counters_.EntityData.Leafs["dma-tx-buf-alloc-count"] = types.YLeaf{"DmaTxBufAllocCount", counters_.DmaTxBufAllocCount}
    counters_.EntityData.Leafs["dma-tx-buf-free-count"] = types.YLeaf{"DmaTxBufFreeCount", counters_.DmaTxBufFreeCount}
    return &(counters_.EntityData)
}

// Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters__SwitchCore
type Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters__SwitchCore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..255.
    Core interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    MsiCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    AerCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    HpCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    WdogCount interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    CoreTaskLockLongestWaitTime interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    CoreTaskLockLongestHeldTime interface{}

    // The type is interface{} with range: 0..18446744073709551615.
    CoreTaskUnlockLongestWaitTime interface{}
}

func (switchCore *Controller_Switch_Oper_SwitchDebugCont_Debug_Counters_Location_Counters__SwitchCore) GetEntityData() *types.CommonEntityData {
    switchCore.EntityData.YFilter = switchCore.YFilter
    switchCore.EntityData.YangName = "switch-core"
    switchCore.EntityData.BundleName = "cisco_ios_xr"
    switchCore.EntityData.ParentYangName = "counters"
    switchCore.EntityData.SegmentPath = "switch-core" + "[core='" + fmt.Sprintf("%v", switchCore.Core) + "']"
    switchCore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchCore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchCore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchCore.EntityData.Children = make(map[string]types.YChild)
    switchCore.EntityData.Leafs = make(map[string]types.YLeaf)
    switchCore.EntityData.Leafs["core"] = types.YLeaf{"Core", switchCore.Core}
    switchCore.EntityData.Leafs["msi-count"] = types.YLeaf{"MsiCount", switchCore.MsiCount}
    switchCore.EntityData.Leafs["aer-count"] = types.YLeaf{"AerCount", switchCore.AerCount}
    switchCore.EntityData.Leafs["hp-count"] = types.YLeaf{"HpCount", switchCore.HpCount}
    switchCore.EntityData.Leafs["wdog-count"] = types.YLeaf{"WdogCount", switchCore.WdogCount}
    switchCore.EntityData.Leafs["core-task-lock-longest-wait-time"] = types.YLeaf{"CoreTaskLockLongestWaitTime", switchCore.CoreTaskLockLongestWaitTime}
    switchCore.EntityData.Leafs["core-task-lock-longest-held-time"] = types.YLeaf{"CoreTaskLockLongestHeldTime", switchCore.CoreTaskLockLongestHeldTime}
    switchCore.EntityData.Leafs["core-task-unlock-longest-wait-time"] = types.YLeaf{"CoreTaskUnlockLongestWaitTime", switchCore.CoreTaskUnlockLongestWaitTime}
    return &(switchCore.EntityData)
}
