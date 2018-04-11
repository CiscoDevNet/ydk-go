// This module contains a collection of YANG definitions for
// Diffserv operational dataCopyright (c) 2017 by Cisco Systems, Inc.All rights reserved.
package diffserv_target_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package diffserv_target_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-diffserv-target-oper diffserv-interfaces-state}", reflect.TypeOf(DiffservInterfacesState{}))
    ydk.RegisterEntity("Cisco-IOS-XE-diffserv-target-oper:diffserv-interfaces-state", reflect.TypeOf(DiffservInterfacesState{}))
}

type Direction struct {
}

func (id Direction) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:direction"
}

type Inbound struct {
}

func (id Inbound) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:inbound"
}

type Outbound struct {
}

func (id Outbound) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:outbound"
}

// DiffservInterfacesState
// Interface configuration parameters.
type DiffservInterfacesState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of configured interfaces on the device. The type is slice of
    // DiffservInterfacesState_DiffservInterface.
    DiffservInterface []DiffservInterfacesState_DiffservInterface
}

func (diffservInterfacesState *DiffservInterfacesState) GetEntityData() *types.CommonEntityData {
    diffservInterfacesState.EntityData.YFilter = diffservInterfacesState.YFilter
    diffservInterfacesState.EntityData.YangName = "diffserv-interfaces-state"
    diffservInterfacesState.EntityData.BundleName = "cisco_ios_xe"
    diffservInterfacesState.EntityData.ParentYangName = "Cisco-IOS-XE-diffserv-target-oper"
    diffservInterfacesState.EntityData.SegmentPath = "Cisco-IOS-XE-diffserv-target-oper:diffserv-interfaces-state"
    diffservInterfacesState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservInterfacesState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservInterfacesState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservInterfacesState.EntityData.Children = make(map[string]types.YChild)
    diffservInterfacesState.EntityData.Children["diffserv-interface"] = types.YChild{"DiffservInterface", nil}
    for i := range diffservInterfacesState.DiffservInterface {
        diffservInterfacesState.EntityData.Children[types.GetSegmentPath(&diffservInterfacesState.DiffservInterface[i])] = types.YChild{"DiffservInterface", &diffservInterfacesState.DiffservInterface[i]}
    }
    diffservInterfacesState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diffservInterfacesState.EntityData)
}

// DiffservInterfacesState_DiffservInterface
// The list of configured interfaces on the device.
type DiffservInterfacesState_DiffservInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string.
    Name interface{}

    // policy target for inbound or outbound direction. The type is slice of
    // DiffservInterfacesState_DiffservInterface_DiffservTargetEntry.
    DiffservTargetEntry []DiffservInterfacesState_DiffservInterface_DiffservTargetEntry
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetEntityData() *types.CommonEntityData {
    diffservInterface.EntityData.YFilter = diffservInterface.YFilter
    diffservInterface.EntityData.YangName = "diffserv-interface"
    diffservInterface.EntityData.BundleName = "cisco_ios_xe"
    diffservInterface.EntityData.ParentYangName = "diffserv-interfaces-state"
    diffservInterface.EntityData.SegmentPath = "diffserv-interface" + "[name='" + fmt.Sprintf("%v", diffservInterface.Name) + "']"
    diffservInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservInterface.EntityData.Children = make(map[string]types.YChild)
    diffservInterface.EntityData.Children["diffserv-target-entry"] = types.YChild{"DiffservTargetEntry", nil}
    for i := range diffservInterface.DiffservTargetEntry {
        diffservInterface.EntityData.Children[types.GetSegmentPath(&diffservInterface.DiffservTargetEntry[i])] = types.YChild{"DiffservTargetEntry", &diffservInterface.DiffservTargetEntry[i]}
    }
    diffservInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservInterface.EntityData.Leafs["name"] = types.YLeaf{"Name", diffservInterface.Name}
    return &(diffservInterface.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry
// policy target for inbound or outbound direction
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Direction fo the traffic flow either inbound or
    // outbound. The type is one of the following: InboundOutbound.
    Direction interface{}

    // This attribute is a key. Policy entry name. The type is string.
    PolicyName interface{}

    // Statistics for each Classifier Entry in a Policy. The type is slice of
    // DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics.
    DiffservTargetClassifierStatistics []DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetEntityData() *types.CommonEntityData {
    diffservTargetEntry.EntityData.YFilter = diffservTargetEntry.YFilter
    diffservTargetEntry.EntityData.YangName = "diffserv-target-entry"
    diffservTargetEntry.EntityData.BundleName = "cisco_ios_xe"
    diffservTargetEntry.EntityData.ParentYangName = "diffserv-interface"
    diffservTargetEntry.EntityData.SegmentPath = "diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
    diffservTargetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservTargetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservTargetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservTargetEntry.EntityData.Children = make(map[string]types.YChild)
    diffservTargetEntry.EntityData.Children["diffserv-target-classifier-statistics"] = types.YChild{"DiffservTargetClassifierStatistics", nil}
    for i := range diffservTargetEntry.DiffservTargetClassifierStatistics {
        diffservTargetEntry.EntityData.Children[types.GetSegmentPath(&diffservTargetEntry.DiffservTargetClassifierStatistics[i])] = types.YChild{"DiffservTargetClassifierStatistics", &diffservTargetEntry.DiffservTargetClassifierStatistics[i]}
    }
    diffservTargetEntry.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservTargetEntry.EntityData.Leafs["direction"] = types.YLeaf{"Direction", diffservTargetEntry.Direction}
    diffservTargetEntry.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", diffservTargetEntry.PolicyName}
    return &(diffservTargetEntry.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics
// Statistics for each Classifier Entry in a Policy
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Classifier Entry Name. The type is string.
    ClassifierEntryName interface{}

    // This attribute is a key. Path of the Classifier Entry in a hierarchial
    // policy . The type is string.
    ParentPath interface{}

    // This group defines the classifier filter statistics of  each classifier
    // entry         .
    ClassifierEntryStatistics DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics

    // Meter statistics. The type is slice of
    // DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics.
    MeterStatistics []DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics

    // queue related statistics .
    QueuingStatistics DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetEntityData() *types.CommonEntityData {
    diffservTargetClassifierStatistics.EntityData.YFilter = diffservTargetClassifierStatistics.YFilter
    diffservTargetClassifierStatistics.EntityData.YangName = "diffserv-target-classifier-statistics"
    diffservTargetClassifierStatistics.EntityData.BundleName = "cisco_ios_xe"
    diffservTargetClassifierStatistics.EntityData.ParentYangName = "diffserv-target-entry"
    diffservTargetClassifierStatistics.EntityData.SegmentPath = "diffserv-target-classifier-statistics" + "[classifier-entry-name='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ClassifierEntryName) + "']" + "[parent-path='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ParentPath) + "']"
    diffservTargetClassifierStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    diffservTargetClassifierStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    diffservTargetClassifierStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    diffservTargetClassifierStatistics.EntityData.Children = make(map[string]types.YChild)
    diffservTargetClassifierStatistics.EntityData.Children["classifier-entry-statistics"] = types.YChild{"ClassifierEntryStatistics", &diffservTargetClassifierStatistics.ClassifierEntryStatistics}
    diffservTargetClassifierStatistics.EntityData.Children["meter-statistics"] = types.YChild{"MeterStatistics", nil}
    for i := range diffservTargetClassifierStatistics.MeterStatistics {
        diffservTargetClassifierStatistics.EntityData.Children[types.GetSegmentPath(&diffservTargetClassifierStatistics.MeterStatistics[i])] = types.YChild{"MeterStatistics", &diffservTargetClassifierStatistics.MeterStatistics[i]}
    }
    diffservTargetClassifierStatistics.EntityData.Children["queuing-statistics"] = types.YChild{"QueuingStatistics", &diffservTargetClassifierStatistics.QueuingStatistics}
    diffservTargetClassifierStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    diffservTargetClassifierStatistics.EntityData.Leafs["classifier-entry-name"] = types.YLeaf{"ClassifierEntryName", diffservTargetClassifierStatistics.ClassifierEntryName}
    diffservTargetClassifierStatistics.EntityData.Leafs["parent-path"] = types.YLeaf{"ParentPath", diffservTargetClassifierStatistics.ParentPath}
    return &(diffservTargetClassifierStatistics.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics
// 
// This group defines the classifier filter statistics of 
// each classifier entry
//        
// 
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of total packets which filtered  to the classifier-entry. The type
    // is interface{} with range: 0..18446744073709551615.
    ClassifiedPkts interface{}

    // Number of total bytes which filtered   to the classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615.
    ClassifiedBytes interface{}

    // Rate of average data flow through the   classifier-entry. The type is
    // interface{} with range: 0..18446744073709551615. Units are bits-per-second.
    ClassifiedRate interface{}
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetEntityData() *types.CommonEntityData {
    classifierEntryStatistics.EntityData.YFilter = classifierEntryStatistics.YFilter
    classifierEntryStatistics.EntityData.YangName = "classifier-entry-statistics"
    classifierEntryStatistics.EntityData.BundleName = "cisco_ios_xe"
    classifierEntryStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    classifierEntryStatistics.EntityData.SegmentPath = "classifier-entry-statistics"
    classifierEntryStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    classifierEntryStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    classifierEntryStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    classifierEntryStatistics.EntityData.Children = make(map[string]types.YChild)
    classifierEntryStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    classifierEntryStatistics.EntityData.Leafs["classified-pkts"] = types.YLeaf{"ClassifiedPkts", classifierEntryStatistics.ClassifiedPkts}
    classifierEntryStatistics.EntityData.Leafs["classified-bytes"] = types.YLeaf{"ClassifiedBytes", classifierEntryStatistics.ClassifiedBytes}
    classifierEntryStatistics.EntityData.Leafs["classified-rate"] = types.YLeaf{"ClassifiedRate", classifierEntryStatistics.ClassifiedRate}
    return &(classifierEntryStatistics.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics
// Meter statistics
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Meter Identifier. The type is interface{} with
    // range: 0..65535.
    MeterId interface{}

    // Number of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedPkts interface{}

    // Bytes of packets which succeed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterSucceedBytes interface{}

    // Number of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedPkts interface{}

    // Bytes of packets which failed the meter. The type is interface{} with
    // range: 0..18446744073709551615.
    MeterFailedBytes interface{}
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetEntityData() *types.CommonEntityData {
    meterStatistics.EntityData.YFilter = meterStatistics.YFilter
    meterStatistics.EntityData.YangName = "meter-statistics"
    meterStatistics.EntityData.BundleName = "cisco_ios_xe"
    meterStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    meterStatistics.EntityData.SegmentPath = "meter-statistics" + "[meter-id='" + fmt.Sprintf("%v", meterStatistics.MeterId) + "']"
    meterStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    meterStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    meterStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    meterStatistics.EntityData.Children = make(map[string]types.YChild)
    meterStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    meterStatistics.EntityData.Leafs["meter-id"] = types.YLeaf{"MeterId", meterStatistics.MeterId}
    meterStatistics.EntityData.Leafs["meter-succeed-pkts"] = types.YLeaf{"MeterSucceedPkts", meterStatistics.MeterSucceedPkts}
    meterStatistics.EntityData.Leafs["meter-succeed-bytes"] = types.YLeaf{"MeterSucceedBytes", meterStatistics.MeterSucceedBytes}
    meterStatistics.EntityData.Leafs["meter-failed-pkts"] = types.YLeaf{"MeterFailedPkts", meterStatistics.MeterFailedPkts}
    meterStatistics.EntityData.Leafs["meter-failed-bytes"] = types.YLeaf{"MeterFailedBytes", meterStatistics.MeterFailedBytes}
    return &(meterStatistics.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
// queue related statistics 
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of packets transmitted from queue . The type is interface{} with
    // range: 0..18446744073709551615.
    OutputPkts interface{}

    // Number of bytes transmitted from queue . The type is interface{} with
    // range: 0..18446744073709551615.
    OutputBytes interface{}

    // Number of packets currently buffered . The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizePkts interface{}

    // Number of bytes currently buffered . The type is interface{} with range:
    // 0..18446744073709551615.
    QueueSizeBytes interface{}

    // Total number of packets dropped . The type is interface{} with range:
    // 0..18446744073709551615.
    DropPkts interface{}

    // Total number of bytes dropped . The type is interface{} with range:
    // 0..18446744073709551615.
    DropBytes interface{}

    // Container for WRED statistics.
    WredStats DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetEntityData() *types.CommonEntityData {
    queuingStatistics.EntityData.YFilter = queuingStatistics.YFilter
    queuingStatistics.EntityData.YangName = "queuing-statistics"
    queuingStatistics.EntityData.BundleName = "cisco_ios_xe"
    queuingStatistics.EntityData.ParentYangName = "diffserv-target-classifier-statistics"
    queuingStatistics.EntityData.SegmentPath = "queuing-statistics"
    queuingStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    queuingStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    queuingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    queuingStatistics.EntityData.Children = make(map[string]types.YChild)
    queuingStatistics.EntityData.Children["wred-stats"] = types.YChild{"WredStats", &queuingStatistics.WredStats}
    queuingStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    queuingStatistics.EntityData.Leafs["output-pkts"] = types.YLeaf{"OutputPkts", queuingStatistics.OutputPkts}
    queuingStatistics.EntityData.Leafs["output-bytes"] = types.YLeaf{"OutputBytes", queuingStatistics.OutputBytes}
    queuingStatistics.EntityData.Leafs["queue-size-pkts"] = types.YLeaf{"QueueSizePkts", queuingStatistics.QueueSizePkts}
    queuingStatistics.EntityData.Leafs["queue-size-bytes"] = types.YLeaf{"QueueSizeBytes", queuingStatistics.QueueSizeBytes}
    queuingStatistics.EntityData.Leafs["drop-pkts"] = types.YLeaf{"DropPkts", queuingStatistics.DropPkts}
    queuingStatistics.EntityData.Leafs["drop-bytes"] = types.YLeaf{"DropBytes", queuingStatistics.DropBytes}
    return &(queuingStatistics.EntityData)
}

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
// Container for WRED statistics
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Early drop packets . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropPkts interface{}

    // Early drop bytes . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropBytes interface{}
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetEntityData() *types.CommonEntityData {
    wredStats.EntityData.YFilter = wredStats.YFilter
    wredStats.EntityData.YangName = "wred-stats"
    wredStats.EntityData.BundleName = "cisco_ios_xe"
    wredStats.EntityData.ParentYangName = "queuing-statistics"
    wredStats.EntityData.SegmentPath = "wred-stats"
    wredStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    wredStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    wredStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    wredStats.EntityData.Children = make(map[string]types.YChild)
    wredStats.EntityData.Leafs = make(map[string]types.YLeaf)
    wredStats.EntityData.Leafs["early-drop-pkts"] = types.YLeaf{"EarlyDropPkts", wredStats.EarlyDropPkts}
    wredStats.EntityData.Leafs["early-drop-bytes"] = types.YLeaf{"EarlyDropBytes", wredStats.EarlyDropBytes}
    return &(wredStats.EntityData)
}

