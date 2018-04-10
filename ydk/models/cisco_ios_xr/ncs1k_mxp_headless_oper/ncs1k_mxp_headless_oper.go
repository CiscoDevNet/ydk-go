// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp-headless package operational data.
// 
// This module contains definitions
// for the following management objects:
//   headless-func-data: Information related to headless
//     functionality
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_mxp_headless_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_mxp_headless_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-mxp-headless-oper headless-func-data}", reflect.TypeOf(HeadlessFuncData{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data", reflect.TypeOf(HeadlessFuncData{}))
}

// MxpOtnPrbsStatus represents Mxp otn prbs status
type MxpOtnPrbsStatus string

const (
    // Locked
    MxpOtnPrbsStatus_locked MxpOtnPrbsStatus = "locked"

    // Not Locked
    MxpOtnPrbsStatus_not_locked MxpOtnPrbsStatus = "not-locked"

    // Not Applicable
    MxpOtnPrbsStatus_not_applicable MxpOtnPrbsStatus = "not-applicable"

    // Locked With Errors
    MxpOtnPrbsStatus_locked_with_errors MxpOtnPrbsStatus = "locked-with-errors"
)

// HeadlessFuncData
// Information related to headless functionality
type HeadlessFuncData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OTN Statistics collected during last headless operation.
    OtnPortNames HeadlessFuncData_OtnPortNames

    // Ethernet Statistics collected during last headless operation.
    EthernetPortNames HeadlessFuncData_EthernetPortNames
}

func (headlessFuncData *HeadlessFuncData) GetEntityData() *types.CommonEntityData {
    headlessFuncData.EntityData.YFilter = headlessFuncData.YFilter
    headlessFuncData.EntityData.YangName = "headless-func-data"
    headlessFuncData.EntityData.BundleName = "cisco_ios_xr"
    headlessFuncData.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-mxp-headless-oper"
    headlessFuncData.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data"
    headlessFuncData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    headlessFuncData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    headlessFuncData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    headlessFuncData.EntityData.Children = make(map[string]types.YChild)
    headlessFuncData.EntityData.Children["otn-port-names"] = types.YChild{"OtnPortNames", &headlessFuncData.OtnPortNames}
    headlessFuncData.EntityData.Children["ethernet-port-names"] = types.YChild{"EthernetPortNames", &headlessFuncData.EthernetPortNames}
    headlessFuncData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(headlessFuncData.EntityData)
}

// HeadlessFuncData_OtnPortNames
// OTN Statistics collected during last headless
// operation
type HeadlessFuncData_OtnPortNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // port Name. The type is slice of HeadlessFuncData_OtnPortNames_OtnPortName.
    OtnPortName []HeadlessFuncData_OtnPortNames_OtnPortName
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetEntityData() *types.CommonEntityData {
    otnPortNames.EntityData.YFilter = otnPortNames.YFilter
    otnPortNames.EntityData.YangName = "otn-port-names"
    otnPortNames.EntityData.BundleName = "cisco_ios_xr"
    otnPortNames.EntityData.ParentYangName = "headless-func-data"
    otnPortNames.EntityData.SegmentPath = "otn-port-names"
    otnPortNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPortNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPortNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPortNames.EntityData.Children = make(map[string]types.YChild)
    otnPortNames.EntityData.Children["otn-port-name"] = types.YChild{"OtnPortName", nil}
    for i := range otnPortNames.OtnPortName {
        otnPortNames.EntityData.Children[types.GetSegmentPath(&otnPortNames.OtnPortName[i])] = types.YChild{"OtnPortName", &otnPortNames.OtnPortName[i]}
    }
    otnPortNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(otnPortNames.EntityData)
}

// HeadlessFuncData_OtnPortNames_OtnPortName
// port Name
type HeadlessFuncData_OtnPortNames_OtnPortName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Started Stateful. The type is bool.
    StartedStateful interface{}

    // Headless Start Time. The type is string with length: 0..64.
    HeadlessStartTime interface{}

    // Headless End Time. The type is string with length: 0..64.
    HeadlessEndTime interface{}

    // OTN statistics.
    OtnStatistics HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics

    // PRBS Statistics.
    PrbsStatistics HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics
}

func (otnPortName *HeadlessFuncData_OtnPortNames_OtnPortName) GetEntityData() *types.CommonEntityData {
    otnPortName.EntityData.YFilter = otnPortName.YFilter
    otnPortName.EntityData.YangName = "otn-port-name"
    otnPortName.EntityData.BundleName = "cisco_ios_xr"
    otnPortName.EntityData.ParentYangName = "otn-port-names"
    otnPortName.EntityData.SegmentPath = "otn-port-name" + "[name='" + fmt.Sprintf("%v", otnPortName.Name) + "']"
    otnPortName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPortName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPortName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPortName.EntityData.Children = make(map[string]types.YChild)
    otnPortName.EntityData.Children["otn-statistics"] = types.YChild{"OtnStatistics", &otnPortName.OtnStatistics}
    otnPortName.EntityData.Children["prbs-statistics"] = types.YChild{"PrbsStatistics", &otnPortName.PrbsStatistics}
    otnPortName.EntityData.Leafs = make(map[string]types.YLeaf)
    otnPortName.EntityData.Leafs["name"] = types.YLeaf{"Name", otnPortName.Name}
    otnPortName.EntityData.Leafs["started-stateful"] = types.YLeaf{"StartedStateful", otnPortName.StartedStateful}
    otnPortName.EntityData.Leafs["headless-start-time"] = types.YLeaf{"HeadlessStartTime", otnPortName.HeadlessStartTime}
    otnPortName.EntityData.Leafs["headless-end-time"] = types.YLeaf{"HeadlessEndTime", otnPortName.HeadlessEndTime}
    return &(otnPortName.EntityData)
}

// HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics
// OTN statistics
type HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SmBip. The type is interface{} with range: 0..18446744073709551615.
    SmBip interface{}

    // SmBei. The type is interface{} with range: 0..18446744073709551615.
    SmBei interface{}

    // FecEc. The type is interface{} with range: 0..18446744073709551615.
    FecEc interface{}

    // FecUc. The type is interface{} with range: 0..18446744073709551615.
    FecUc interface{}
}

func (otnStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_OtnStatistics) GetEntityData() *types.CommonEntityData {
    otnStatistics.EntityData.YFilter = otnStatistics.YFilter
    otnStatistics.EntityData.YangName = "otn-statistics"
    otnStatistics.EntityData.BundleName = "cisco_ios_xr"
    otnStatistics.EntityData.ParentYangName = "otn-port-name"
    otnStatistics.EntityData.SegmentPath = "otn-statistics"
    otnStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnStatistics.EntityData.Children = make(map[string]types.YChild)
    otnStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    otnStatistics.EntityData.Leafs["sm-bip"] = types.YLeaf{"SmBip", otnStatistics.SmBip}
    otnStatistics.EntityData.Leafs["sm-bei"] = types.YLeaf{"SmBei", otnStatistics.SmBei}
    otnStatistics.EntityData.Leafs["fec-ec"] = types.YLeaf{"FecEc", otnStatistics.FecEc}
    otnStatistics.EntityData.Leafs["fec-uc"] = types.YLeaf{"FecUc", otnStatistics.FecUc}
    return &(otnStatistics.EntityData)
}

// HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics
// PRBS Statistics
type HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EBC. The type is interface{} with range: 0..18446744073709551615.
    Ebc interface{}

    // SyncStatus. The type is MxpOtnPrbsStatus.
    SyncStatus interface{}
}

func (prbsStatistics *HeadlessFuncData_OtnPortNames_OtnPortName_PrbsStatistics) GetEntityData() *types.CommonEntityData {
    prbsStatistics.EntityData.YFilter = prbsStatistics.YFilter
    prbsStatistics.EntityData.YangName = "prbs-statistics"
    prbsStatistics.EntityData.BundleName = "cisco_ios_xr"
    prbsStatistics.EntityData.ParentYangName = "otn-port-name"
    prbsStatistics.EntityData.SegmentPath = "prbs-statistics"
    prbsStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbsStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbsStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbsStatistics.EntityData.Children = make(map[string]types.YChild)
    prbsStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    prbsStatistics.EntityData.Leafs["ebc"] = types.YLeaf{"Ebc", prbsStatistics.Ebc}
    prbsStatistics.EntityData.Leafs["sync-status"] = types.YLeaf{"SyncStatus", prbsStatistics.SyncStatus}
    return &(prbsStatistics.EntityData)
}

// HeadlessFuncData_EthernetPortNames
// Ethernet Statistics collected during last
// headless operation
type HeadlessFuncData_EthernetPortNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Port Name. The type is slice of
    // HeadlessFuncData_EthernetPortNames_EthernetPortName.
    EthernetPortName []HeadlessFuncData_EthernetPortNames_EthernetPortName
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetEntityData() *types.CommonEntityData {
    ethernetPortNames.EntityData.YFilter = ethernetPortNames.YFilter
    ethernetPortNames.EntityData.YangName = "ethernet-port-names"
    ethernetPortNames.EntityData.BundleName = "cisco_ios_xr"
    ethernetPortNames.EntityData.ParentYangName = "headless-func-data"
    ethernetPortNames.EntityData.SegmentPath = "ethernet-port-names"
    ethernetPortNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetPortNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetPortNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetPortNames.EntityData.Children = make(map[string]types.YChild)
    ethernetPortNames.EntityData.Children["ethernet-port-name"] = types.YChild{"EthernetPortName", nil}
    for i := range ethernetPortNames.EthernetPortName {
        ethernetPortNames.EntityData.Children[types.GetSegmentPath(&ethernetPortNames.EthernetPortName[i])] = types.YChild{"EthernetPortName", &ethernetPortNames.EthernetPortName[i]}
    }
    ethernetPortNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ethernetPortNames.EntityData)
}

// HeadlessFuncData_EthernetPortNames_EthernetPortName
// Port Name
type HeadlessFuncData_EthernetPortNames_EthernetPortName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // Started Stateful. The type is bool.
    StartedStateful interface{}

    // Headless Start Time. The type is string with length: 0..64.
    HeadlessStartTime interface{}

    // Headless End Time. The type is string with length: 0..64.
    HeadlessEndTime interface{}

    // Ether Statistics.
    EtherStatistics HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics
}

func (ethernetPortName *HeadlessFuncData_EthernetPortNames_EthernetPortName) GetEntityData() *types.CommonEntityData {
    ethernetPortName.EntityData.YFilter = ethernetPortName.YFilter
    ethernetPortName.EntityData.YangName = "ethernet-port-name"
    ethernetPortName.EntityData.BundleName = "cisco_ios_xr"
    ethernetPortName.EntityData.ParentYangName = "ethernet-port-names"
    ethernetPortName.EntityData.SegmentPath = "ethernet-port-name" + "[name='" + fmt.Sprintf("%v", ethernetPortName.Name) + "']"
    ethernetPortName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetPortName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetPortName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetPortName.EntityData.Children = make(map[string]types.YChild)
    ethernetPortName.EntityData.Children["ether-statistics"] = types.YChild{"EtherStatistics", &ethernetPortName.EtherStatistics}
    ethernetPortName.EntityData.Leafs = make(map[string]types.YLeaf)
    ethernetPortName.EntityData.Leafs["name"] = types.YLeaf{"Name", ethernetPortName.Name}
    ethernetPortName.EntityData.Leafs["started-stateful"] = types.YLeaf{"StartedStateful", ethernetPortName.StartedStateful}
    ethernetPortName.EntityData.Leafs["headless-start-time"] = types.YLeaf{"HeadlessStartTime", ethernetPortName.HeadlessStartTime}
    ethernetPortName.EntityData.Leafs["headless-end-time"] = types.YLeaf{"HeadlessEndTime", ethernetPortName.HeadlessEndTime}
    return &(ethernetPortName.EntityData)
}

// HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics
// Ether Statistics
type HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxPktsOverSized. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsOverSized interface{}

    // RxPktsBadFcs. The type is interface{} with range: 0..18446744073709551615.
    RxPktsBadFcs interface{}

    // RxErrorJabbers. The type is interface{} with range:
    // 0..18446744073709551615.
    RxErrorJabbers interface{}

    // RxPktsMulticast. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsMulticast interface{}

    // RxPktsBroadcast. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsBroadcast interface{}

    // RxPktsUnderSized. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPktsUnderSized interface{}

    // RxPackets. The type is interface{} with range: 0..18446744073709551615.
    RxPackets interface{}

    // RxTotalBytes. The type is interface{} with range: 0..18446744073709551615.
    RxTotalBytes interface{}

    // RxBytesGood. The type is interface{} with range: 0..18446744073709551615.
    RxBytesGood interface{}

    // RxPktsGood. The type is interface{} with range: 0..18446744073709551615.
    RxPktsGood interface{}

    // TxBytesGood. The type is interface{} with range: 0..18446744073709551615.
    TxBytesGood interface{}

    // TxPktsGood. The type is interface{} with range: 0..18446744073709551615.
    TxPktsGood interface{}

    // RxRecvFragments. The type is interface{} with range:
    // 0..18446744073709551615.
    RxRecvFragments interface{}

    // RxPkts64Bytes. The type is interface{} with range: 0..18446744073709551615.
    RxPkts64Bytes interface{}

    // RxPkts65To127Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts65To127Bytes interface{}

    // RxPkts128to255Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts128To255Bytes interface{}

    // RxPkts256To511Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts256To511Bytes interface{}

    // RxPkts512To1023Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts512To1023Bytes interface{}

    // RxPkts1024To1518Bytes. The type is interface{} with range:
    // 0..18446744073709551615.
    RxPkts1024To1518Bytes interface{}

    // RxPktsUnicast. The type is interface{} with range: 0..18446744073709551615.
    RxPktsUnicast interface{}

    // TxPackets. The type is interface{} with range: 0..18446744073709551615.
    TxPackets interface{}

    // TxTotalBytes. The type is interface{} with range: 0..18446744073709551615.
    TxTotalBytes interface{}

    // TxPktsUnderSized. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPktsUnderSized interface{}

    // TxPktsOverSized. The type is interface{} with range:
    // 0..18446744073709551615.
    TxPktsOverSized interface{}

    // TxFragments. The type is interface{} with range: 0..18446744073709551615.
    TxFragments interface{}

    // TxJabber. The type is interface{} with range: 0..18446744073709551615.
    TxJabber interface{}

    // TxBadFCS. The type is interface{} with range: 0..18446744073709551615.
    TxBadFcs interface{}

    // RxPktDrop. The type is interface{} with range: 0..18446744073709551615.
    RxPktDrop interface{}

    // RxPause. The type is interface{} with range: 0..18446744073709551615.
    RxPause interface{}

    // TxPause. The type is interface{} with range: 0..18446744073709551615.
    TxPause interface{}

    // RxLldpPkt. The type is interface{} with range: 0..18446744073709551615.
    RxLldpPkt interface{}

    // Rx8021QPkt. The type is interface{} with range: 0..18446744073709551615.
    Rx8021QPkt interface{}
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetEntityData() *types.CommonEntityData {
    etherStatistics.EntityData.YFilter = etherStatistics.YFilter
    etherStatistics.EntityData.YangName = "ether-statistics"
    etherStatistics.EntityData.BundleName = "cisco_ios_xr"
    etherStatistics.EntityData.ParentYangName = "ethernet-port-name"
    etherStatistics.EntityData.SegmentPath = "ether-statistics"
    etherStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etherStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etherStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etherStatistics.EntityData.Children = make(map[string]types.YChild)
    etherStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    etherStatistics.EntityData.Leafs["rx-pkts-over-sized"] = types.YLeaf{"RxPktsOverSized", etherStatistics.RxPktsOverSized}
    etherStatistics.EntityData.Leafs["rx-pkts-bad-fcs"] = types.YLeaf{"RxPktsBadFcs", etherStatistics.RxPktsBadFcs}
    etherStatistics.EntityData.Leafs["rx-error-jabbers"] = types.YLeaf{"RxErrorJabbers", etherStatistics.RxErrorJabbers}
    etherStatistics.EntityData.Leafs["rx-pkts-multicast"] = types.YLeaf{"RxPktsMulticast", etherStatistics.RxPktsMulticast}
    etherStatistics.EntityData.Leafs["rx-pkts-broadcast"] = types.YLeaf{"RxPktsBroadcast", etherStatistics.RxPktsBroadcast}
    etherStatistics.EntityData.Leafs["rx-pkts-under-sized"] = types.YLeaf{"RxPktsUnderSized", etherStatistics.RxPktsUnderSized}
    etherStatistics.EntityData.Leafs["rx-packets"] = types.YLeaf{"RxPackets", etherStatistics.RxPackets}
    etherStatistics.EntityData.Leafs["rx-total-bytes"] = types.YLeaf{"RxTotalBytes", etherStatistics.RxTotalBytes}
    etherStatistics.EntityData.Leafs["rx-bytes-good"] = types.YLeaf{"RxBytesGood", etherStatistics.RxBytesGood}
    etherStatistics.EntityData.Leafs["rx-pkts-good"] = types.YLeaf{"RxPktsGood", etherStatistics.RxPktsGood}
    etherStatistics.EntityData.Leafs["tx-bytes-good"] = types.YLeaf{"TxBytesGood", etherStatistics.TxBytesGood}
    etherStatistics.EntityData.Leafs["tx-pkts-good"] = types.YLeaf{"TxPktsGood", etherStatistics.TxPktsGood}
    etherStatistics.EntityData.Leafs["rx-recv-fragments"] = types.YLeaf{"RxRecvFragments", etherStatistics.RxRecvFragments}
    etherStatistics.EntityData.Leafs["rx-pkts64-bytes"] = types.YLeaf{"RxPkts64Bytes", etherStatistics.RxPkts64Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts65-to127-bytes"] = types.YLeaf{"RxPkts65To127Bytes", etherStatistics.RxPkts65To127Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts128to255-bytes"] = types.YLeaf{"RxPkts128To255Bytes", etherStatistics.RxPkts128To255Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts256-to511-bytes"] = types.YLeaf{"RxPkts256To511Bytes", etherStatistics.RxPkts256To511Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts512-to1023-bytes"] = types.YLeaf{"RxPkts512To1023Bytes", etherStatistics.RxPkts512To1023Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts1024-to1518-bytes"] = types.YLeaf{"RxPkts1024To1518Bytes", etherStatistics.RxPkts1024To1518Bytes}
    etherStatistics.EntityData.Leafs["rx-pkts-unicast"] = types.YLeaf{"RxPktsUnicast", etherStatistics.RxPktsUnicast}
    etherStatistics.EntityData.Leafs["tx-packets"] = types.YLeaf{"TxPackets", etherStatistics.TxPackets}
    etherStatistics.EntityData.Leafs["tx-total-bytes"] = types.YLeaf{"TxTotalBytes", etherStatistics.TxTotalBytes}
    etherStatistics.EntityData.Leafs["tx-pkts-under-sized"] = types.YLeaf{"TxPktsUnderSized", etherStatistics.TxPktsUnderSized}
    etherStatistics.EntityData.Leafs["tx-pkts-over-sized"] = types.YLeaf{"TxPktsOverSized", etherStatistics.TxPktsOverSized}
    etherStatistics.EntityData.Leafs["tx-fragments"] = types.YLeaf{"TxFragments", etherStatistics.TxFragments}
    etherStatistics.EntityData.Leafs["tx-jabber"] = types.YLeaf{"TxJabber", etherStatistics.TxJabber}
    etherStatistics.EntityData.Leafs["tx-bad-fcs"] = types.YLeaf{"TxBadFcs", etherStatistics.TxBadFcs}
    etherStatistics.EntityData.Leafs["rx-pkt-drop"] = types.YLeaf{"RxPktDrop", etherStatistics.RxPktDrop}
    etherStatistics.EntityData.Leafs["rx-pause"] = types.YLeaf{"RxPause", etherStatistics.RxPause}
    etherStatistics.EntityData.Leafs["tx-pause"] = types.YLeaf{"TxPause", etherStatistics.TxPause}
    etherStatistics.EntityData.Leafs["rx-lldp-pkt"] = types.YLeaf{"RxLldpPkt", etherStatistics.RxLldpPkt}
    etherStatistics.EntityData.Leafs["rx8021q-pkt"] = types.YLeaf{"Rx8021QPkt", etherStatistics.Rx8021QPkt}
    return &(etherStatistics.EntityData)
}

