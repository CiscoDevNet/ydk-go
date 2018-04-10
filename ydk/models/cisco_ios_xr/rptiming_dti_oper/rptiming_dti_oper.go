// This module contains a collection of YANG definitions
// for Cisco IOS-XR rptiming-dti package operational data.
// 
// This module contains definitions
// for the following management objects:
//   dti-controller: DTI interface controller status and
//     configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package rptiming_dti_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package rptiming_dti_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-rptiming-dti-oper dti-controller}", reflect.TypeOf(DtiController{}))
    ydk.RegisterEntity("Cisco-IOS-XR-rptiming-dti-oper:dti-controller", reflect.TypeOf(DtiController{}))
}

// DtiController
// DTI interface controller status and configuration
type DtiController struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes applicable to DTI controller.
    Nodes DtiController_Nodes
}

func (dtiController *DtiController) GetEntityData() *types.CommonEntityData {
    dtiController.EntityData.YFilter = dtiController.YFilter
    dtiController.EntityData.YangName = "dti-controller"
    dtiController.EntityData.BundleName = "cisco_ios_xr"
    dtiController.EntityData.ParentYangName = "Cisco-IOS-XR-rptiming-dti-oper"
    dtiController.EntityData.SegmentPath = "Cisco-IOS-XR-rptiming-dti-oper:dti-controller"
    dtiController.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dtiController.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dtiController.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dtiController.EntityData.Children = make(map[string]types.YChild)
    dtiController.EntityData.Children["nodes"] = types.YChild{"Nodes", &dtiController.Nodes}
    dtiController.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(dtiController.EntityData)
}

// DtiController_Nodes
// List of nodes applicable to DTI controller
type DtiController_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DTI operational data for a single node. The type is slice of
    // DtiController_Nodes_Node.
    Node []DtiController_Nodes_Node
}

func (nodes *DtiController_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "dti-controller"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// DtiController_Nodes_Node
// DTI operational data for a single node
type DtiController_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node Name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Display DTI client status.
    Client DtiController_Nodes_Node_Client

    // Display DTI input port status.
    Port DtiController_Nodes_Node_Port

    // Display DTI time-of-day status.
    TimeOfDay DtiController_Nodes_Node_TimeOfDay
}

func (node *DtiController_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["client"] = types.YChild{"Client", &node.Client}
    node.EntityData.Children["port"] = types.YChild{"Port", &node.Port}
    node.EntityData.Children["time-of-day"] = types.YChild{"TimeOfDay", &node.TimeOfDay}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// DtiController_Nodes_Node_Client
// Display DTI client status
type DtiController_Nodes_Node_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // timestamp comparator enable. The type is string with length: 0..50.
    TimestampComparatorEnable interface{}

    // register write enable. The type is string with length: 0..50.
    RegisterWriteEnable interface{}

    // revertive mode enable. The type is string with length: 0..50.
    RevertiveModeEnable interface{}

    // port mode select. The type is string with length: 0..50.
    PortModeSelect interface{}

    // force freerun. The type is string with length: 0..50.
    ForceFreerun interface{}

    // reference select port. The type is string with length: 0..50.
    ReferenceSelectPort interface{}

    // timestamp sync detected. The type is string with length: 0..50.
    TimestampSyncDetected interface{}

    // 10Mhz reference detected. The type is string with length: 0..50.
    Reference10MhzDetected interface{}

    // active input port. The type is string with length: 0..50.
    ActiveInputPort interface{}

    // client state. The type is string with length: 0..50.
    ClientState interface{}
}

func (client *DtiController_Nodes_Node_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "node"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["timestamp-comparator-enable"] = types.YLeaf{"TimestampComparatorEnable", client.TimestampComparatorEnable}
    client.EntityData.Leafs["register-write-enable"] = types.YLeaf{"RegisterWriteEnable", client.RegisterWriteEnable}
    client.EntityData.Leafs["revertive-mode-enable"] = types.YLeaf{"RevertiveModeEnable", client.RevertiveModeEnable}
    client.EntityData.Leafs["port-mode-select"] = types.YLeaf{"PortModeSelect", client.PortModeSelect}
    client.EntityData.Leafs["force-freerun"] = types.YLeaf{"ForceFreerun", client.ForceFreerun}
    client.EntityData.Leafs["reference-select-port"] = types.YLeaf{"ReferenceSelectPort", client.ReferenceSelectPort}
    client.EntityData.Leafs["timestamp-sync-detected"] = types.YLeaf{"TimestampSyncDetected", client.TimestampSyncDetected}
    client.EntityData.Leafs["reference10mhz-detected"] = types.YLeaf{"Reference10MhzDetected", client.Reference10MhzDetected}
    client.EntityData.Leafs["active-input-port"] = types.YLeaf{"ActiveInputPort", client.ActiveInputPort}
    client.EntityData.Leafs["client-state"] = types.YLeaf{"ClientState", client.ClientState}
    return &(client.EntityData)
}

// DtiController_Nodes_Node_Port
// Display DTI input port status
type DtiController_Nodes_Node_Port struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // port1 frame error rate greater than 5 percent. The type is string with
    // length: 0..50. Units are percentage.
    Port1FrErrRateGreater5Per interface{}

    // port1 frame error rate greater than 2 percent. The type is string with
    // length: 0..50. Units are percentage.
    Port1FrErrRateGreater2Per interface{}

    // port1 DTI signal detected. The type is string with length: 0..50.
    Port1DtiSignalDetected interface{}

    // port1 server timing source. The type is string with length: 0..50.
    Port1ServerTimingSource interface{}

    // port1 server type. The type is string with length: 0..50.
    Port1ServerType interface{}

    // port1 server clock type. The type is string with length: 0..50.
    Port1ServerClockType interface{}

    // port1 server state. The type is string with length: 0..50.
    Port1ServerState interface{}

    // port1 client performance stable. The type is string with length: 0..50.
    Port1ClientPerfStable interface{}

    // port1 cable advance valid. The type is string with length: 0..50.
    Port1CableAdvanceValid interface{}

    // port2 frame error rate greater than 5 percent. The type is string with
    // length: 0..50. Units are percentage.
    Port2FrErrRateGreater5Per interface{}

    // port2 frame error rate greater than 2 percent. The type is string with
    // length: 0..50. Units are percentage.
    Port2FrErrRateGreater2Per interface{}

    // port2 DTI signal detected. The type is string with length: 0..50.
    Port2DtiSignalDetected interface{}

    // port2 server timing source. The type is string with length: 0..50.
    Port2ServerTimingSource interface{}

    // port2 server type. The type is string with length: 0..50.
    Port2ServerType interface{}

    // port2 server clock type. The type is string with length: 0..50.
    Port2ServerClockType interface{}

    // port2 server state. The type is string with length: 0..50.
    Port2ServerState interface{}

    // port2 client performance stable. The type is string with length: 0..50.
    Port2ClientPerfStable interface{}

    // port2 cable advance valid. The type is string with length: 0..50.
    Port2CableAdvanceValid interface{}
}

func (port *DtiController_Nodes_Node_Port) GetEntityData() *types.CommonEntityData {
    port.EntityData.YFilter = port.YFilter
    port.EntityData.YangName = "port"
    port.EntityData.BundleName = "cisco_ios_xr"
    port.EntityData.ParentYangName = "node"
    port.EntityData.SegmentPath = "port"
    port.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    port.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    port.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    port.EntityData.Children = make(map[string]types.YChild)
    port.EntityData.Leafs = make(map[string]types.YLeaf)
    port.EntityData.Leafs["port1-fr-err-rate-greater5-per"] = types.YLeaf{"Port1FrErrRateGreater5Per", port.Port1FrErrRateGreater5Per}
    port.EntityData.Leafs["port1-fr-err-rate-greater2-per"] = types.YLeaf{"Port1FrErrRateGreater2Per", port.Port1FrErrRateGreater2Per}
    port.EntityData.Leafs["port1-dti-signal-detected"] = types.YLeaf{"Port1DtiSignalDetected", port.Port1DtiSignalDetected}
    port.EntityData.Leafs["port1-server-timing-source"] = types.YLeaf{"Port1ServerTimingSource", port.Port1ServerTimingSource}
    port.EntityData.Leafs["port1-server-type"] = types.YLeaf{"Port1ServerType", port.Port1ServerType}
    port.EntityData.Leafs["port1-server-clock-type"] = types.YLeaf{"Port1ServerClockType", port.Port1ServerClockType}
    port.EntityData.Leafs["port1-server-state"] = types.YLeaf{"Port1ServerState", port.Port1ServerState}
    port.EntityData.Leafs["port1-client-perf-stable"] = types.YLeaf{"Port1ClientPerfStable", port.Port1ClientPerfStable}
    port.EntityData.Leafs["port1-cable-advance-valid"] = types.YLeaf{"Port1CableAdvanceValid", port.Port1CableAdvanceValid}
    port.EntityData.Leafs["port2-fr-err-rate-greater5-per"] = types.YLeaf{"Port2FrErrRateGreater5Per", port.Port2FrErrRateGreater5Per}
    port.EntityData.Leafs["port2-fr-err-rate-greater2-per"] = types.YLeaf{"Port2FrErrRateGreater2Per", port.Port2FrErrRateGreater2Per}
    port.EntityData.Leafs["port2-dti-signal-detected"] = types.YLeaf{"Port2DtiSignalDetected", port.Port2DtiSignalDetected}
    port.EntityData.Leafs["port2-server-timing-source"] = types.YLeaf{"Port2ServerTimingSource", port.Port2ServerTimingSource}
    port.EntityData.Leafs["port2-server-type"] = types.YLeaf{"Port2ServerType", port.Port2ServerType}
    port.EntityData.Leafs["port2-server-clock-type"] = types.YLeaf{"Port2ServerClockType", port.Port2ServerClockType}
    port.EntityData.Leafs["port2-server-state"] = types.YLeaf{"Port2ServerState", port.Port2ServerState}
    port.EntityData.Leafs["port2-client-perf-stable"] = types.YLeaf{"Port2ClientPerfStable", port.Port2ClientPerfStable}
    port.EntityData.Leafs["port2-cable-advance-valid"] = types.YLeaf{"Port2CableAdvanceValid", port.Port2CableAdvanceValid}
    return &(port.EntityData)
}

// DtiController_Nodes_Node_TimeOfDay
// Display DTI time-of-day status
type DtiController_Nodes_Node_TimeOfDay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // port1 tod message mode. The type is string with length: 0..50.
    Port1TodMessageMode interface{}

    // port1 tod state. The type is string with length: 0..50.
    Port1TodState interface{}

    // port1 tod time setting mode. The type is string with length: 0..50.
    Port1TodTimeSettingMode interface{}

    // port1 gps seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Port1GpsSeconds interface{}

    // port1 leap seconds. The type is interface{} with range: 0..255. Units are
    // second.
    Port1LeapSeconds interface{}

    // port1 calendar time valid. The type is string with length: 0..50.
    Port1CalendarTimeValid interface{}

    // port1 modified julian date. The type is string with length: 0..50.
    Port1ModifiedJulianDate interface{}

    // port1 date. The type is string with length: 0..50.
    Port1Date interface{}

    // port1 time. The type is string with length: 0..50.
    Port1Time interface{}

    // port1 time zone offset. The type is string with length: 0..50.
    Port1TimeZoneOffset interface{}

    // port1 leap second indicator. The type is interface{} with range: 0..255.
    Port1LeapSecondIndicator interface{}

    // port2 tod message mode. The type is string with length: 0..50.
    Port2TodMessageMode interface{}

    // port2 tod state. The type is string with length: 0..50.
    Port2TodState interface{}

    // port2 tod time setting mode. The type is string with length: 0..50.
    Port2TodTimeSettingMode interface{}

    // port2 gps seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    Port2GpsSeconds interface{}

    // port2 leap seconds. The type is interface{} with range: 0..255. Units are
    // second.
    Port2LeapSeconds interface{}

    // port2 calendar time valid. The type is string with length: 0..50.
    Port2CalendarTimeValid interface{}

    // port2 modified julian date. The type is string with length: 0..50.
    Port2ModifiedJulianDate interface{}

    // port2 date. The type is string with length: 0..50.
    Port2Date interface{}

    // port2 time. The type is string with length: 0..50.
    Port2Time interface{}

    // port2 time zone offset. The type is string with length: 0..50.
    Port2TimeZoneOffset interface{}

    // port2 leap second indicator. The type is interface{} with range: 0..255.
    Port2LeapSecondIndicator interface{}
}

func (timeOfDay *DtiController_Nodes_Node_TimeOfDay) GetEntityData() *types.CommonEntityData {
    timeOfDay.EntityData.YFilter = timeOfDay.YFilter
    timeOfDay.EntityData.YangName = "time-of-day"
    timeOfDay.EntityData.BundleName = "cisco_ios_xr"
    timeOfDay.EntityData.ParentYangName = "node"
    timeOfDay.EntityData.SegmentPath = "time-of-day"
    timeOfDay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeOfDay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeOfDay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeOfDay.EntityData.Children = make(map[string]types.YChild)
    timeOfDay.EntityData.Leafs = make(map[string]types.YLeaf)
    timeOfDay.EntityData.Leafs["port1-tod-message-mode"] = types.YLeaf{"Port1TodMessageMode", timeOfDay.Port1TodMessageMode}
    timeOfDay.EntityData.Leafs["port1-tod-state"] = types.YLeaf{"Port1TodState", timeOfDay.Port1TodState}
    timeOfDay.EntityData.Leafs["port1-tod-time-setting-mode"] = types.YLeaf{"Port1TodTimeSettingMode", timeOfDay.Port1TodTimeSettingMode}
    timeOfDay.EntityData.Leafs["port1-gps-seconds"] = types.YLeaf{"Port1GpsSeconds", timeOfDay.Port1GpsSeconds}
    timeOfDay.EntityData.Leafs["port1-leap-seconds"] = types.YLeaf{"Port1LeapSeconds", timeOfDay.Port1LeapSeconds}
    timeOfDay.EntityData.Leafs["port1-calendar-time-valid"] = types.YLeaf{"Port1CalendarTimeValid", timeOfDay.Port1CalendarTimeValid}
    timeOfDay.EntityData.Leafs["port1-modified-julian-date"] = types.YLeaf{"Port1ModifiedJulianDate", timeOfDay.Port1ModifiedJulianDate}
    timeOfDay.EntityData.Leafs["port1-date"] = types.YLeaf{"Port1Date", timeOfDay.Port1Date}
    timeOfDay.EntityData.Leafs["port1-time"] = types.YLeaf{"Port1Time", timeOfDay.Port1Time}
    timeOfDay.EntityData.Leafs["port1-time-zone-offset"] = types.YLeaf{"Port1TimeZoneOffset", timeOfDay.Port1TimeZoneOffset}
    timeOfDay.EntityData.Leafs["port1-leap-second-indicator"] = types.YLeaf{"Port1LeapSecondIndicator", timeOfDay.Port1LeapSecondIndicator}
    timeOfDay.EntityData.Leafs["port2-tod-message-mode"] = types.YLeaf{"Port2TodMessageMode", timeOfDay.Port2TodMessageMode}
    timeOfDay.EntityData.Leafs["port2-tod-state"] = types.YLeaf{"Port2TodState", timeOfDay.Port2TodState}
    timeOfDay.EntityData.Leafs["port2-tod-time-setting-mode"] = types.YLeaf{"Port2TodTimeSettingMode", timeOfDay.Port2TodTimeSettingMode}
    timeOfDay.EntityData.Leafs["port2-gps-seconds"] = types.YLeaf{"Port2GpsSeconds", timeOfDay.Port2GpsSeconds}
    timeOfDay.EntityData.Leafs["port2-leap-seconds"] = types.YLeaf{"Port2LeapSeconds", timeOfDay.Port2LeapSeconds}
    timeOfDay.EntityData.Leafs["port2-calendar-time-valid"] = types.YLeaf{"Port2CalendarTimeValid", timeOfDay.Port2CalendarTimeValid}
    timeOfDay.EntityData.Leafs["port2-modified-julian-date"] = types.YLeaf{"Port2ModifiedJulianDate", timeOfDay.Port2ModifiedJulianDate}
    timeOfDay.EntityData.Leafs["port2-date"] = types.YLeaf{"Port2Date", timeOfDay.Port2Date}
    timeOfDay.EntityData.Leafs["port2-time"] = types.YLeaf{"Port2Time", timeOfDay.Port2Time}
    timeOfDay.EntityData.Leafs["port2-time-zone-offset"] = types.YLeaf{"Port2TimeZoneOffset", timeOfDay.Port2TimeZoneOffset}
    timeOfDay.EntityData.Leafs["port2-leap-second-indicator"] = types.YLeaf{"Port2LeapSecondIndicator", timeOfDay.Port2LeapSecondIndicator}
    return &(timeOfDay.EntityData)
}

