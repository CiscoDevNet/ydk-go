// This module contains a collection of YANG definitions for
// monitoring memory usage of processes in a Network Element.Copyright (c) 2016-2017 by Cisco Systems, Inc.All rights reserved.
package cfm_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cfm_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-cfm-oper cfm-statistics}", reflect.TypeOf(CfmStatistics{}))
    ydk.RegisterEntity("Cisco-IOS-XE-cfm-oper:cfm-statistics", reflect.TypeOf(CfmStatistics{}))
}

// CfmLastClearedType
type CfmLastClearedType string

const (
    CfmLastClearedType_never_cleared CfmLastClearedType = "never-cleared"

    CfmLastClearedType_cleared_before CfmLastClearedType = "cleared-before"
)

// CfmStatistics
// Data nodes for CFM Statistics.
type CfmStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    CfmMeps CfmStatistics_CfmMeps
}

func (cfmStatistics *CfmStatistics) GetFilter() yfilter.YFilter { return cfmStatistics.YFilter }

func (cfmStatistics *CfmStatistics) SetFilter(yf yfilter.YFilter) { cfmStatistics.YFilter = yf }

func (cfmStatistics *CfmStatistics) GetGoName(yname string) string {
    if yname == "cfm-meps" { return "CfmMeps" }
    return ""
}

func (cfmStatistics *CfmStatistics) GetSegmentPath() string {
    return "Cisco-IOS-XE-cfm-oper:cfm-statistics"
}

func (cfmStatistics *CfmStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cfm-meps" {
        return &cfmStatistics.CfmMeps
    }
    return nil
}

func (cfmStatistics *CfmStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cfm-meps"] = &cfmStatistics.CfmMeps
    return children
}

func (cfmStatistics *CfmStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cfmStatistics *CfmStatistics) GetBundleName() string { return "cisco_ios_xe" }

func (cfmStatistics *CfmStatistics) GetYangName() string { return "cfm-statistics" }

func (cfmStatistics *CfmStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfmStatistics *CfmStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfmStatistics *CfmStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfmStatistics *CfmStatistics) SetParent(parent types.Entity) { cfmStatistics.parent = parent }

func (cfmStatistics *CfmStatistics) GetParent() types.Entity { return cfmStatistics.parent }

func (cfmStatistics *CfmStatistics) GetParentYangName() string { return "Cisco-IOS-XE-cfm-oper" }

// CfmStatistics_CfmMeps
type CfmStatistics_CfmMeps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The list of MEP entries in the system. The type is slice of
    // CfmStatistics_CfmMeps_CfmMep.
    CfmMep []CfmStatistics_CfmMeps_CfmMep
}

func (cfmMeps *CfmStatistics_CfmMeps) GetFilter() yfilter.YFilter { return cfmMeps.YFilter }

func (cfmMeps *CfmStatistics_CfmMeps) SetFilter(yf yfilter.YFilter) { cfmMeps.YFilter = yf }

func (cfmMeps *CfmStatistics_CfmMeps) GetGoName(yname string) string {
    if yname == "cfm-mep" { return "CfmMep" }
    return ""
}

func (cfmMeps *CfmStatistics_CfmMeps) GetSegmentPath() string {
    return "cfm-meps"
}

func (cfmMeps *CfmStatistics_CfmMeps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cfm-mep" {
        for _, c := range cfmMeps.CfmMep {
            if cfmMeps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CfmStatistics_CfmMeps_CfmMep{}
        cfmMeps.CfmMep = append(cfmMeps.CfmMep, child)
        return &cfmMeps.CfmMep[len(cfmMeps.CfmMep)-1]
    }
    return nil
}

func (cfmMeps *CfmStatistics_CfmMeps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cfmMeps.CfmMep {
        children[cfmMeps.CfmMep[i].GetSegmentPath()] = &cfmMeps.CfmMep[i]
    }
    return children
}

func (cfmMeps *CfmStatistics_CfmMeps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cfmMeps *CfmStatistics_CfmMeps) GetBundleName() string { return "cisco_ios_xe" }

func (cfmMeps *CfmStatistics_CfmMeps) GetYangName() string { return "cfm-meps" }

func (cfmMeps *CfmStatistics_CfmMeps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfmMeps *CfmStatistics_CfmMeps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfmMeps *CfmStatistics_CfmMeps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfmMeps *CfmStatistics_CfmMeps) SetParent(parent types.Entity) { cfmMeps.parent = parent }

func (cfmMeps *CfmStatistics_CfmMeps) GetParent() types.Entity { return cfmMeps.parent }

func (cfmMeps *CfmStatistics_CfmMeps) GetParentYangName() string { return "cfm-statistics" }

// CfmStatistics_CfmMeps_CfmMep
// The list of MEP entries in the system.
type CfmStatistics_CfmMeps_CfmMep struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the Domain corresponding the the MEP.
    // The type is string.
    DomainName interface{}

    // This attribute is a key. The name of the MA corresponding the the MEP. The
    // type is string.
    MaName interface{}

    // This attribute is a key. ID of the MEP. The type is interface{} with range:
    // 0..4294967295.
    Mpid interface{}

    // The number of CCMs transmitted from the local MEP. The type is interface{}
    // with range: 0..18446744073709551615.
    CcmTransmitted interface{}

    // The number of CCM sequence number errors detected. The type is interface{}
    // with range: 0..18446744073709551615.
    CcmSeqErrors interface{}

    // The number of unexpected linktrace reply packets  received at this MEP. The
    // type is interface{} with range: 0..18446744073709551615.
    LtrUnexpected interface{}

    // The number of loopback reply packets transmitted from the local MEP. The
    // type is interface{} with range: 0..18446744073709551615.
    LbrTransmitted interface{}

    // The number of loopback reply packets received  with sequence number errors.
    // The type is interface{} with range: 0..18446744073709551615.
    LbrSeqErrors interface{}

    // The number of valid loopback reply packets received. The type is
    // interface{} with range: 0..18446744073709551615.
    LbrReceivedOk interface{}

    // The number of loopback reply packets received  with corrupted data pattern.
    // The type is interface{} with range: 0..18446744073709551615.
    LbrReceivedBad interface{}

    
    LastCleared CfmStatistics_CfmMeps_CfmMep_LastCleared
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetFilter() yfilter.YFilter { return cfmMep.YFilter }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) SetFilter(yf yfilter.YFilter) { cfmMep.YFilter = yf }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetGoName(yname string) string {
    if yname == "domain-name" { return "DomainName" }
    if yname == "ma-name" { return "MaName" }
    if yname == "mpid" { return "Mpid" }
    if yname == "ccm-transmitted" { return "CcmTransmitted" }
    if yname == "ccm-seq-errors" { return "CcmSeqErrors" }
    if yname == "ltr-unexpected" { return "LtrUnexpected" }
    if yname == "lbr-transmitted" { return "LbrTransmitted" }
    if yname == "lbr-seq-errors" { return "LbrSeqErrors" }
    if yname == "lbr-received-ok" { return "LbrReceivedOk" }
    if yname == "lbr-received-bad" { return "LbrReceivedBad" }
    if yname == "last-cleared" { return "LastCleared" }
    return ""
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetSegmentPath() string {
    return "cfm-mep" + "[domain-name='" + fmt.Sprintf("%v", cfmMep.DomainName) + "']" + "[ma-name='" + fmt.Sprintf("%v", cfmMep.MaName) + "']" + "[mpid='" + fmt.Sprintf("%v", cfmMep.Mpid) + "']"
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "last-cleared" {
        return &cfmMep.LastCleared
    }
    return nil
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["last-cleared"] = &cfmMep.LastCleared
    return children
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-name"] = cfmMep.DomainName
    leafs["ma-name"] = cfmMep.MaName
    leafs["mpid"] = cfmMep.Mpid
    leafs["ccm-transmitted"] = cfmMep.CcmTransmitted
    leafs["ccm-seq-errors"] = cfmMep.CcmSeqErrors
    leafs["ltr-unexpected"] = cfmMep.LtrUnexpected
    leafs["lbr-transmitted"] = cfmMep.LbrTransmitted
    leafs["lbr-seq-errors"] = cfmMep.LbrSeqErrors
    leafs["lbr-received-ok"] = cfmMep.LbrReceivedOk
    leafs["lbr-received-bad"] = cfmMep.LbrReceivedBad
    return leafs
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetBundleName() string { return "cisco_ios_xe" }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetYangName() string { return "cfm-mep" }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) SetParent(parent types.Entity) { cfmMep.parent = parent }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetParent() types.Entity { return cfmMep.parent }

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetParentYangName() string { return "cfm-meps" }

// CfmStatistics_CfmMeps_CfmMep_LastCleared
type CfmStatistics_CfmMeps_CfmMep_LastCleared struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is interface{}.
    Never interface{}

    // The type is string.
    Time interface{}
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetFilter() yfilter.YFilter { return lastCleared.YFilter }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) SetFilter(yf yfilter.YFilter) { lastCleared.YFilter = yf }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetGoName(yname string) string {
    if yname == "never" { return "Never" }
    if yname == "time" { return "Time" }
    return ""
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetSegmentPath() string {
    return "last-cleared"
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["never"] = lastCleared.Never
    leafs["time"] = lastCleared.Time
    return leafs
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetBundleName() string { return "cisco_ios_xe" }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetYangName() string { return "last-cleared" }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) SetParent(parent types.Entity) { lastCleared.parent = parent }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetParent() types.Entity { return lastCleared.parent }

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetParentYangName() string { return "cfm-mep" }

