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

type Inbound struct {
}

func (id Inbound) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:inbound"
}

type Direction struct {
}

func (id Direction) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:direction"
}

type Outbound struct {
}

func (id Outbound) String() string {
	return "Cisco-IOS-XE-diffserv-target-oper:outbound"
}

// DiffservInterfacesState
// Interface configuration parameters.
type DiffservInterfacesState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of configured interfaces on the device. The type is slice of
    // DiffservInterfacesState_DiffservInterface.
    DiffservInterface []DiffservInterfacesState_DiffservInterface
}

func (diffservInterfacesState *DiffservInterfacesState) GetFilter() yfilter.YFilter { return diffservInterfacesState.YFilter }

func (diffservInterfacesState *DiffservInterfacesState) SetFilter(yf yfilter.YFilter) { diffservInterfacesState.YFilter = yf }

func (diffservInterfacesState *DiffservInterfacesState) GetGoName(yname string) string {
    if yname == "diffserv-interface" { return "DiffservInterface" }
    return ""
}

func (diffservInterfacesState *DiffservInterfacesState) GetSegmentPath() string {
    return "Cisco-IOS-XE-diffserv-target-oper:diffserv-interfaces-state"
}

func (diffservInterfacesState *DiffservInterfacesState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffserv-interface" {
        for _, c := range diffservInterfacesState.DiffservInterface {
            if diffservInterfacesState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DiffservInterfacesState_DiffservInterface{}
        diffservInterfacesState.DiffservInterface = append(diffservInterfacesState.DiffservInterface, child)
        return &diffservInterfacesState.DiffservInterface[len(diffservInterfacesState.DiffservInterface)-1]
    }
    return nil
}

func (diffservInterfacesState *DiffservInterfacesState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservInterfacesState.DiffservInterface {
        children[diffservInterfacesState.DiffservInterface[i].GetSegmentPath()] = &diffservInterfacesState.DiffservInterface[i]
    }
    return children
}

func (diffservInterfacesState *DiffservInterfacesState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diffservInterfacesState *DiffservInterfacesState) GetBundleName() string { return "cisco_ios_xe" }

func (diffservInterfacesState *DiffservInterfacesState) GetYangName() string { return "diffserv-interfaces-state" }

func (diffservInterfacesState *DiffservInterfacesState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservInterfacesState *DiffservInterfacesState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservInterfacesState *DiffservInterfacesState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservInterfacesState *DiffservInterfacesState) SetParent(parent types.Entity) { diffservInterfacesState.parent = parent }

func (diffservInterfacesState *DiffservInterfacesState) GetParent() types.Entity { return diffservInterfacesState.parent }

func (diffservInterfacesState *DiffservInterfacesState) GetParentYangName() string { return "Cisco-IOS-XE-diffserv-target-oper" }

// DiffservInterfacesState_DiffservInterface
// The list of configured interfaces on the device.
type DiffservInterfacesState_DiffservInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string.
    Name interface{}

    // policy target for inbound or outbound direction. The type is slice of
    // DiffservInterfacesState_DiffservInterface_DiffservTargetEntry.
    DiffservTargetEntry []DiffservInterfacesState_DiffservInterface_DiffservTargetEntry
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetFilter() yfilter.YFilter { return diffservInterface.YFilter }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) SetFilter(yf yfilter.YFilter) { diffservInterface.YFilter = yf }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "diffserv-target-entry" { return "DiffservTargetEntry" }
    return ""
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetSegmentPath() string {
    return "diffserv-interface" + "[name='" + fmt.Sprintf("%v", diffservInterface.Name) + "']"
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffserv-target-entry" {
        for _, c := range diffservInterface.DiffservTargetEntry {
            if diffservInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DiffservInterfacesState_DiffservInterface_DiffservTargetEntry{}
        diffservInterface.DiffservTargetEntry = append(diffservInterface.DiffservTargetEntry, child)
        return &diffservInterface.DiffservTargetEntry[len(diffservInterface.DiffservTargetEntry)-1]
    }
    return nil
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservInterface.DiffservTargetEntry {
        children[diffservInterface.DiffservTargetEntry[i].GetSegmentPath()] = &diffservInterface.DiffservTargetEntry[i]
    }
    return children
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = diffservInterface.Name
    return leafs
}

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetBundleName() string { return "cisco_ios_xe" }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetYangName() string { return "diffserv-interface" }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) SetParent(parent types.Entity) { diffservInterface.parent = parent }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetParent() types.Entity { return diffservInterface.parent }

func (diffservInterface *DiffservInterfacesState_DiffservInterface) GetParentYangName() string { return "diffserv-interfaces-state" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry
// policy target for inbound or outbound direction
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry struct {
    parent types.Entity
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

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetFilter() yfilter.YFilter { return diffservTargetEntry.YFilter }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) SetFilter(yf yfilter.YFilter) { diffservTargetEntry.YFilter = yf }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetGoName(yname string) string {
    if yname == "direction" { return "Direction" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "diffserv-target-classifier-statistics" { return "DiffservTargetClassifierStatistics" }
    return ""
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetSegmentPath() string {
    return "diffserv-target-entry" + "[direction='" + fmt.Sprintf("%v", diffservTargetEntry.Direction) + "']" + "[policy-name='" + fmt.Sprintf("%v", diffservTargetEntry.PolicyName) + "']"
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diffserv-target-classifier-statistics" {
        for _, c := range diffservTargetEntry.DiffservTargetClassifierStatistics {
            if diffservTargetEntry.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics{}
        diffservTargetEntry.DiffservTargetClassifierStatistics = append(diffservTargetEntry.DiffservTargetClassifierStatistics, child)
        return &diffservTargetEntry.DiffservTargetClassifierStatistics[len(diffservTargetEntry.DiffservTargetClassifierStatistics)-1]
    }
    return nil
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diffservTargetEntry.DiffservTargetClassifierStatistics {
        children[diffservTargetEntry.DiffservTargetClassifierStatistics[i].GetSegmentPath()] = &diffservTargetEntry.DiffservTargetClassifierStatistics[i]
    }
    return children
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["direction"] = diffservTargetEntry.Direction
    leafs["policy-name"] = diffservTargetEntry.PolicyName
    return leafs
}

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetBundleName() string { return "cisco_ios_xe" }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetYangName() string { return "diffserv-target-entry" }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) SetParent(parent types.Entity) { diffservTargetEntry.parent = parent }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetParent() types.Entity { return diffservTargetEntry.parent }

func (diffservTargetEntry *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry) GetParentYangName() string { return "diffserv-interface" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics
// Statistics for each Classifier Entry in a Policy
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics struct {
    parent types.Entity
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

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetFilter() yfilter.YFilter { return diffservTargetClassifierStatistics.YFilter }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) SetFilter(yf yfilter.YFilter) { diffservTargetClassifierStatistics.YFilter = yf }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetGoName(yname string) string {
    if yname == "classifier-entry-name" { return "ClassifierEntryName" }
    if yname == "parent-path" { return "ParentPath" }
    if yname == "classifier-entry-statistics" { return "ClassifierEntryStatistics" }
    if yname == "meter-statistics" { return "MeterStatistics" }
    if yname == "queuing-statistics" { return "QueuingStatistics" }
    return ""
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetSegmentPath() string {
    return "diffserv-target-classifier-statistics" + "[classifier-entry-name='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ClassifierEntryName) + "']" + "[parent-path='" + fmt.Sprintf("%v", diffservTargetClassifierStatistics.ParentPath) + "']"
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "classifier-entry-statistics" {
        return &diffservTargetClassifierStatistics.ClassifierEntryStatistics
    }
    if childYangName == "meter-statistics" {
        for _, c := range diffservTargetClassifierStatistics.MeterStatistics {
            if diffservTargetClassifierStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics{}
        diffservTargetClassifierStatistics.MeterStatistics = append(diffservTargetClassifierStatistics.MeterStatistics, child)
        return &diffservTargetClassifierStatistics.MeterStatistics[len(diffservTargetClassifierStatistics.MeterStatistics)-1]
    }
    if childYangName == "queuing-statistics" {
        return &diffservTargetClassifierStatistics.QueuingStatistics
    }
    return nil
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["classifier-entry-statistics"] = &diffservTargetClassifierStatistics.ClassifierEntryStatistics
    for i := range diffservTargetClassifierStatistics.MeterStatistics {
        children[diffservTargetClassifierStatistics.MeterStatistics[i].GetSegmentPath()] = &diffservTargetClassifierStatistics.MeterStatistics[i]
    }
    children["queuing-statistics"] = &diffservTargetClassifierStatistics.QueuingStatistics
    return children
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classifier-entry-name"] = diffservTargetClassifierStatistics.ClassifierEntryName
    leafs["parent-path"] = diffservTargetClassifierStatistics.ParentPath
    return leafs
}

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetYangName() string { return "diffserv-target-classifier-statistics" }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) SetParent(parent types.Entity) { diffservTargetClassifierStatistics.parent = parent }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetParent() types.Entity { return diffservTargetClassifierStatistics.parent }

func (diffservTargetClassifierStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics) GetParentYangName() string { return "diffserv-target-entry" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics
// 
// This group defines the classifier filter statistics of 
// each classifier entry
//        
// 
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics struct {
    parent types.Entity
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

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetFilter() yfilter.YFilter { return classifierEntryStatistics.YFilter }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) SetFilter(yf yfilter.YFilter) { classifierEntryStatistics.YFilter = yf }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetGoName(yname string) string {
    if yname == "classified-pkts" { return "ClassifiedPkts" }
    if yname == "classified-bytes" { return "ClassifiedBytes" }
    if yname == "classified-rate" { return "ClassifiedRate" }
    return ""
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetSegmentPath() string {
    return "classifier-entry-statistics"
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["classified-pkts"] = classifierEntryStatistics.ClassifiedPkts
    leafs["classified-bytes"] = classifierEntryStatistics.ClassifiedBytes
    leafs["classified-rate"] = classifierEntryStatistics.ClassifiedRate
    return leafs
}

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetYangName() string { return "classifier-entry-statistics" }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) SetParent(parent types.Entity) { classifierEntryStatistics.parent = parent }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetParent() types.Entity { return classifierEntryStatistics.parent }

func (classifierEntryStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_ClassifierEntryStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics
// Meter statistics
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics struct {
    parent types.Entity
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

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetFilter() yfilter.YFilter { return meterStatistics.YFilter }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) SetFilter(yf yfilter.YFilter) { meterStatistics.YFilter = yf }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetGoName(yname string) string {
    if yname == "meter-id" { return "MeterId" }
    if yname == "meter-succeed-pkts" { return "MeterSucceedPkts" }
    if yname == "meter-succeed-bytes" { return "MeterSucceedBytes" }
    if yname == "meter-failed-pkts" { return "MeterFailedPkts" }
    if yname == "meter-failed-bytes" { return "MeterFailedBytes" }
    return ""
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetSegmentPath() string {
    return "meter-statistics" + "[meter-id='" + fmt.Sprintf("%v", meterStatistics.MeterId) + "']"
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["meter-id"] = meterStatistics.MeterId
    leafs["meter-succeed-pkts"] = meterStatistics.MeterSucceedPkts
    leafs["meter-succeed-bytes"] = meterStatistics.MeterSucceedBytes
    leafs["meter-failed-pkts"] = meterStatistics.MeterFailedPkts
    leafs["meter-failed-bytes"] = meterStatistics.MeterFailedBytes
    return leafs
}

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetYangName() string { return "meter-statistics" }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) SetParent(parent types.Entity) { meterStatistics.parent = parent }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetParent() types.Entity { return meterStatistics.parent }

func (meterStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_MeterStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics
// queue related statistics 
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics struct {
    parent types.Entity
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

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetFilter() yfilter.YFilter { return queuingStatistics.YFilter }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) SetFilter(yf yfilter.YFilter) { queuingStatistics.YFilter = yf }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetGoName(yname string) string {
    if yname == "output-pkts" { return "OutputPkts" }
    if yname == "output-bytes" { return "OutputBytes" }
    if yname == "queue-size-pkts" { return "QueueSizePkts" }
    if yname == "queue-size-bytes" { return "QueueSizeBytes" }
    if yname == "drop-pkts" { return "DropPkts" }
    if yname == "drop-bytes" { return "DropBytes" }
    if yname == "wred-stats" { return "WredStats" }
    return ""
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetSegmentPath() string {
    return "queuing-statistics"
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred-stats" {
        return &queuingStatistics.WredStats
    }
    return nil
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wred-stats"] = &queuingStatistics.WredStats
    return children
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["output-pkts"] = queuingStatistics.OutputPkts
    leafs["output-bytes"] = queuingStatistics.OutputBytes
    leafs["queue-size-pkts"] = queuingStatistics.QueueSizePkts
    leafs["queue-size-bytes"] = queuingStatistics.QueueSizeBytes
    leafs["drop-pkts"] = queuingStatistics.DropPkts
    leafs["drop-bytes"] = queuingStatistics.DropBytes
    return leafs
}

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetYangName() string { return "queuing-statistics" }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) SetParent(parent types.Entity) { queuingStatistics.parent = parent }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetParent() types.Entity { return queuingStatistics.parent }

func (queuingStatistics *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics) GetParentYangName() string { return "diffserv-target-classifier-statistics" }

// DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats
// Container for WRED statistics
type DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Early drop packets . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropPkts interface{}

    // Early drop bytes . The type is interface{} with range:
    // 0..18446744073709551615.
    EarlyDropBytes interface{}
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetFilter() yfilter.YFilter { return wredStats.YFilter }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) SetFilter(yf yfilter.YFilter) { wredStats.YFilter = yf }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetGoName(yname string) string {
    if yname == "early-drop-pkts" { return "EarlyDropPkts" }
    if yname == "early-drop-bytes" { return "EarlyDropBytes" }
    return ""
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetSegmentPath() string {
    return "wred-stats"
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["early-drop-pkts"] = wredStats.EarlyDropPkts
    leafs["early-drop-bytes"] = wredStats.EarlyDropBytes
    return leafs
}

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetBundleName() string { return "cisco_ios_xe" }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetYangName() string { return "wred-stats" }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) SetParent(parent types.Entity) { wredStats.parent = parent }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetParent() types.Entity { return wredStats.parent }

func (wredStats *DiffservInterfacesState_DiffservInterface_DiffservTargetEntry_DiffservTargetClassifierStatistics_QueuingStatistics_WredStats) GetParentYangName() string { return "queuing-statistics" }

