// This module contains a collection of YANG definitions for
// monitoring the Connectivity Fault Management protocol operation in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
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

// CfmLastClearedType represents Describes whether CFM stats have been cleared
type CfmLastClearedType string

const (
    // CFM stats have never been cleared
    CfmLastClearedType_never_cleared CfmLastClearedType = "never-cleared"

    // CFM stats have been cleared once before
    CfmLastClearedType_cleared_before CfmLastClearedType = "cleared-before"
)

// CfmStatistics
// Data nodes for CFM Statistics
type CfmStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CFM statistics.
    CfmMeps CfmStatistics_CfmMeps
}

func (cfmStatistics *CfmStatistics) GetEntityData() *types.CommonEntityData {
    cfmStatistics.EntityData.YFilter = cfmStatistics.YFilter
    cfmStatistics.EntityData.YangName = "cfm-statistics"
    cfmStatistics.EntityData.BundleName = "cisco_ios_xe"
    cfmStatistics.EntityData.ParentYangName = "Cisco-IOS-XE-cfm-oper"
    cfmStatistics.EntityData.SegmentPath = "Cisco-IOS-XE-cfm-oper:cfm-statistics"
    cfmStatistics.EntityData.AbsolutePath = cfmStatistics.EntityData.SegmentPath
    cfmStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfmStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfmStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfmStatistics.EntityData.Children = types.NewOrderedMap()
    cfmStatistics.EntityData.Children.Append("cfm-meps", types.YChild{"CfmMeps", &cfmStatistics.CfmMeps})
    cfmStatistics.EntityData.Leafs = types.NewOrderedMap()

    cfmStatistics.EntityData.YListKeys = []string {}

    return &(cfmStatistics.EntityData)
}

// CfmStatistics_CfmMeps
// CFM statistics
type CfmStatistics_CfmMeps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The list of MEP entries in the system. The type is slice of
    // CfmStatistics_CfmMeps_CfmMep.
    CfmMep []*CfmStatistics_CfmMeps_CfmMep
}

func (cfmMeps *CfmStatistics_CfmMeps) GetEntityData() *types.CommonEntityData {
    cfmMeps.EntityData.YFilter = cfmMeps.YFilter
    cfmMeps.EntityData.YangName = "cfm-meps"
    cfmMeps.EntityData.BundleName = "cisco_ios_xe"
    cfmMeps.EntityData.ParentYangName = "cfm-statistics"
    cfmMeps.EntityData.SegmentPath = "cfm-meps"
    cfmMeps.EntityData.AbsolutePath = "Cisco-IOS-XE-cfm-oper:cfm-statistics/" + cfmMeps.EntityData.SegmentPath
    cfmMeps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfmMeps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfmMeps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfmMeps.EntityData.Children = types.NewOrderedMap()
    cfmMeps.EntityData.Children.Append("cfm-mep", types.YChild{"CfmMep", nil})
    for i := range cfmMeps.CfmMep {
        cfmMeps.EntityData.Children.Append(types.GetSegmentPath(cfmMeps.CfmMep[i]), types.YChild{"CfmMep", cfmMeps.CfmMep[i]})
    }
    cfmMeps.EntityData.Leafs = types.NewOrderedMap()

    cfmMeps.EntityData.YListKeys = []string {}

    return &(cfmMeps.EntityData)
}

// CfmStatistics_CfmMeps_CfmMep
// The list of MEP entries in the system
type CfmStatistics_CfmMeps_CfmMep struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Info on when the stats were last cleared.
    LastCleared CfmStatistics_CfmMeps_CfmMep_LastCleared
}

func (cfmMep *CfmStatistics_CfmMeps_CfmMep) GetEntityData() *types.CommonEntityData {
    cfmMep.EntityData.YFilter = cfmMep.YFilter
    cfmMep.EntityData.YangName = "cfm-mep"
    cfmMep.EntityData.BundleName = "cisco_ios_xe"
    cfmMep.EntityData.ParentYangName = "cfm-meps"
    cfmMep.EntityData.SegmentPath = "cfm-mep" + types.AddKeyToken(cfmMep.DomainName, "domain-name") + types.AddKeyToken(cfmMep.MaName, "ma-name") + types.AddKeyToken(cfmMep.Mpid, "mpid")
    cfmMep.EntityData.AbsolutePath = "Cisco-IOS-XE-cfm-oper:cfm-statistics/cfm-meps/" + cfmMep.EntityData.SegmentPath
    cfmMep.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cfmMep.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cfmMep.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cfmMep.EntityData.Children = types.NewOrderedMap()
    cfmMep.EntityData.Children.Append("last-cleared", types.YChild{"LastCleared", &cfmMep.LastCleared})
    cfmMep.EntityData.Leafs = types.NewOrderedMap()
    cfmMep.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", cfmMep.DomainName})
    cfmMep.EntityData.Leafs.Append("ma-name", types.YLeaf{"MaName", cfmMep.MaName})
    cfmMep.EntityData.Leafs.Append("mpid", types.YLeaf{"Mpid", cfmMep.Mpid})
    cfmMep.EntityData.Leafs.Append("ccm-transmitted", types.YLeaf{"CcmTransmitted", cfmMep.CcmTransmitted})
    cfmMep.EntityData.Leafs.Append("ccm-seq-errors", types.YLeaf{"CcmSeqErrors", cfmMep.CcmSeqErrors})
    cfmMep.EntityData.Leafs.Append("ltr-unexpected", types.YLeaf{"LtrUnexpected", cfmMep.LtrUnexpected})
    cfmMep.EntityData.Leafs.Append("lbr-transmitted", types.YLeaf{"LbrTransmitted", cfmMep.LbrTransmitted})
    cfmMep.EntityData.Leafs.Append("lbr-seq-errors", types.YLeaf{"LbrSeqErrors", cfmMep.LbrSeqErrors})
    cfmMep.EntityData.Leafs.Append("lbr-received-ok", types.YLeaf{"LbrReceivedOk", cfmMep.LbrReceivedOk})
    cfmMep.EntityData.Leafs.Append("lbr-received-bad", types.YLeaf{"LbrReceivedBad", cfmMep.LbrReceivedBad})

    cfmMep.EntityData.YListKeys = []string {"DomainName", "MaName", "Mpid"}

    return &(cfmMep.EntityData)
}

// CfmStatistics_CfmMeps_CfmMep_LastCleared
// Info on when the stats were last cleared
type CfmStatistics_CfmMeps_CfmMep_LastCleared struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Never been cleared. The type is interface{}.
    Never interface{}

    // Date and time of the last time stats were cleared. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    Time interface{}
}

func (lastCleared *CfmStatistics_CfmMeps_CfmMep_LastCleared) GetEntityData() *types.CommonEntityData {
    lastCleared.EntityData.YFilter = lastCleared.YFilter
    lastCleared.EntityData.YangName = "last-cleared"
    lastCleared.EntityData.BundleName = "cisco_ios_xe"
    lastCleared.EntityData.ParentYangName = "cfm-mep"
    lastCleared.EntityData.SegmentPath = "last-cleared"
    lastCleared.EntityData.AbsolutePath = "Cisco-IOS-XE-cfm-oper:cfm-statistics/cfm-meps/cfm-mep/" + lastCleared.EntityData.SegmentPath
    lastCleared.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    lastCleared.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    lastCleared.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    lastCleared.EntityData.Children = types.NewOrderedMap()
    lastCleared.EntityData.Leafs = types.NewOrderedMap()
    lastCleared.EntityData.Leafs.Append("never", types.YLeaf{"Never", lastCleared.Never})
    lastCleared.EntityData.Leafs.Append("time", types.YLeaf{"Time", lastCleared.Time})

    lastCleared.EntityData.YListKeys = []string {}

    return &(lastCleared.EntityData)
}

