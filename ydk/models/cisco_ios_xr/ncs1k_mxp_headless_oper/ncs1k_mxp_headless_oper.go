// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-mxp-headless package operational data.
// 
// This module contains definitions
// for the following management objects:
//   headless-func-data: Information related to headless
//     functionality
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    headlessFuncData.EntityData.AbsolutePath = headlessFuncData.EntityData.SegmentPath
    headlessFuncData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    headlessFuncData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    headlessFuncData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    headlessFuncData.EntityData.Children = types.NewOrderedMap()
    headlessFuncData.EntityData.Children.Append("otn-port-names", types.YChild{"OtnPortNames", &headlessFuncData.OtnPortNames})
    headlessFuncData.EntityData.Children.Append("ethernet-port-names", types.YChild{"EthernetPortNames", &headlessFuncData.EthernetPortNames})
    headlessFuncData.EntityData.Leafs = types.NewOrderedMap()

    headlessFuncData.EntityData.YListKeys = []string {}

    return &(headlessFuncData.EntityData)
}

// HeadlessFuncData_OtnPortNames
// OTN Statistics collected during last headless
// operation
type HeadlessFuncData_OtnPortNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // port Name. The type is slice of HeadlessFuncData_OtnPortNames_OtnPortName.
    OtnPortName []*HeadlessFuncData_OtnPortNames_OtnPortName
}

func (otnPortNames *HeadlessFuncData_OtnPortNames) GetEntityData() *types.CommonEntityData {
    otnPortNames.EntityData.YFilter = otnPortNames.YFilter
    otnPortNames.EntityData.YangName = "otn-port-names"
    otnPortNames.EntityData.BundleName = "cisco_ios_xr"
    otnPortNames.EntityData.ParentYangName = "headless-func-data"
    otnPortNames.EntityData.SegmentPath = "otn-port-names"
    otnPortNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/" + otnPortNames.EntityData.SegmentPath
    otnPortNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPortNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPortNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPortNames.EntityData.Children = types.NewOrderedMap()
    otnPortNames.EntityData.Children.Append("otn-port-name", types.YChild{"OtnPortName", nil})
    for i := range otnPortNames.OtnPortName {
        otnPortNames.EntityData.Children.Append(types.GetSegmentPath(otnPortNames.OtnPortName[i]), types.YChild{"OtnPortName", otnPortNames.OtnPortName[i]})
    }
    otnPortNames.EntityData.Leafs = types.NewOrderedMap()

    otnPortNames.EntityData.YListKeys = []string {}

    return &(otnPortNames.EntityData)
}

// HeadlessFuncData_OtnPortNames_OtnPortName
// port Name
type HeadlessFuncData_OtnPortNames_OtnPortName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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
    otnPortName.EntityData.SegmentPath = "otn-port-name" + types.AddKeyToken(otnPortName.Name, "name")
    otnPortName.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/otn-port-names/" + otnPortName.EntityData.SegmentPath
    otnPortName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPortName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPortName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPortName.EntityData.Children = types.NewOrderedMap()
    otnPortName.EntityData.Children.Append("otn-statistics", types.YChild{"OtnStatistics", &otnPortName.OtnStatistics})
    otnPortName.EntityData.Children.Append("prbs-statistics", types.YChild{"PrbsStatistics", &otnPortName.PrbsStatistics})
    otnPortName.EntityData.Leafs = types.NewOrderedMap()
    otnPortName.EntityData.Leafs.Append("name", types.YLeaf{"Name", otnPortName.Name})
    otnPortName.EntityData.Leafs.Append("started-stateful", types.YLeaf{"StartedStateful", otnPortName.StartedStateful})
    otnPortName.EntityData.Leafs.Append("headless-start-time", types.YLeaf{"HeadlessStartTime", otnPortName.HeadlessStartTime})
    otnPortName.EntityData.Leafs.Append("headless-end-time", types.YLeaf{"HeadlessEndTime", otnPortName.HeadlessEndTime})

    otnPortName.EntityData.YListKeys = []string {"Name"}

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
    otnStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/otn-port-names/otn-port-name/" + otnStatistics.EntityData.SegmentPath
    otnStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnStatistics.EntityData.Children = types.NewOrderedMap()
    otnStatistics.EntityData.Leafs = types.NewOrderedMap()
    otnStatistics.EntityData.Leafs.Append("sm-bip", types.YLeaf{"SmBip", otnStatistics.SmBip})
    otnStatistics.EntityData.Leafs.Append("sm-bei", types.YLeaf{"SmBei", otnStatistics.SmBei})
    otnStatistics.EntityData.Leafs.Append("fec-ec", types.YLeaf{"FecEc", otnStatistics.FecEc})
    otnStatistics.EntityData.Leafs.Append("fec-uc", types.YLeaf{"FecUc", otnStatistics.FecUc})

    otnStatistics.EntityData.YListKeys = []string {}

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
    prbsStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/otn-port-names/otn-port-name/" + prbsStatistics.EntityData.SegmentPath
    prbsStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prbsStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prbsStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prbsStatistics.EntityData.Children = types.NewOrderedMap()
    prbsStatistics.EntityData.Leafs = types.NewOrderedMap()
    prbsStatistics.EntityData.Leafs.Append("ebc", types.YLeaf{"Ebc", prbsStatistics.Ebc})
    prbsStatistics.EntityData.Leafs.Append("sync-status", types.YLeaf{"SyncStatus", prbsStatistics.SyncStatus})

    prbsStatistics.EntityData.YListKeys = []string {}

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
    EthernetPortName []*HeadlessFuncData_EthernetPortNames_EthernetPortName
}

func (ethernetPortNames *HeadlessFuncData_EthernetPortNames) GetEntityData() *types.CommonEntityData {
    ethernetPortNames.EntityData.YFilter = ethernetPortNames.YFilter
    ethernetPortNames.EntityData.YangName = "ethernet-port-names"
    ethernetPortNames.EntityData.BundleName = "cisco_ios_xr"
    ethernetPortNames.EntityData.ParentYangName = "headless-func-data"
    ethernetPortNames.EntityData.SegmentPath = "ethernet-port-names"
    ethernetPortNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/" + ethernetPortNames.EntityData.SegmentPath
    ethernetPortNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetPortNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetPortNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetPortNames.EntityData.Children = types.NewOrderedMap()
    ethernetPortNames.EntityData.Children.Append("ethernet-port-name", types.YChild{"EthernetPortName", nil})
    for i := range ethernetPortNames.EthernetPortName {
        ethernetPortNames.EntityData.Children.Append(types.GetSegmentPath(ethernetPortNames.EthernetPortName[i]), types.YChild{"EthernetPortName", ethernetPortNames.EthernetPortName[i]})
    }
    ethernetPortNames.EntityData.Leafs = types.NewOrderedMap()

    ethernetPortNames.EntityData.YListKeys = []string {}

    return &(ethernetPortNames.EntityData)
}

// HeadlessFuncData_EthernetPortNames_EthernetPortName
// Port Name
type HeadlessFuncData_EthernetPortNames_EthernetPortName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
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
    ethernetPortName.EntityData.SegmentPath = "ethernet-port-name" + types.AddKeyToken(ethernetPortName.Name, "name")
    ethernetPortName.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/ethernet-port-names/" + ethernetPortName.EntityData.SegmentPath
    ethernetPortName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ethernetPortName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ethernetPortName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ethernetPortName.EntityData.Children = types.NewOrderedMap()
    ethernetPortName.EntityData.Children.Append("ether-statistics", types.YChild{"EtherStatistics", &ethernetPortName.EtherStatistics})
    ethernetPortName.EntityData.Leafs = types.NewOrderedMap()
    ethernetPortName.EntityData.Leafs.Append("name", types.YLeaf{"Name", ethernetPortName.Name})
    ethernetPortName.EntityData.Leafs.Append("started-stateful", types.YLeaf{"StartedStateful", ethernetPortName.StartedStateful})
    ethernetPortName.EntityData.Leafs.Append("headless-start-time", types.YLeaf{"HeadlessStartTime", ethernetPortName.HeadlessStartTime})
    ethernetPortName.EntityData.Leafs.Append("headless-end-time", types.YLeaf{"HeadlessEndTime", ethernetPortName.HeadlessEndTime})

    ethernetPortName.EntityData.YListKeys = []string {"Name"}

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
    RxPkts128to255Bytes interface{}

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
    Rx8021qPkt interface{}
}

func (etherStatistics *HeadlessFuncData_EthernetPortNames_EthernetPortName_EtherStatistics) GetEntityData() *types.CommonEntityData {
    etherStatistics.EntityData.YFilter = etherStatistics.YFilter
    etherStatistics.EntityData.YangName = "ether-statistics"
    etherStatistics.EntityData.BundleName = "cisco_ios_xr"
    etherStatistics.EntityData.ParentYangName = "ethernet-port-name"
    etherStatistics.EntityData.SegmentPath = "ether-statistics"
    etherStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ncs1k-mxp-headless-oper:headless-func-data/ethernet-port-names/ethernet-port-name/" + etherStatistics.EntityData.SegmentPath
    etherStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    etherStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    etherStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    etherStatistics.EntityData.Children = types.NewOrderedMap()
    etherStatistics.EntityData.Leafs = types.NewOrderedMap()
    etherStatistics.EntityData.Leafs.Append("rx-pkts-over-sized", types.YLeaf{"RxPktsOverSized", etherStatistics.RxPktsOverSized})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-bad-fcs", types.YLeaf{"RxPktsBadFcs", etherStatistics.RxPktsBadFcs})
    etherStatistics.EntityData.Leafs.Append("rx-error-jabbers", types.YLeaf{"RxErrorJabbers", etherStatistics.RxErrorJabbers})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-multicast", types.YLeaf{"RxPktsMulticast", etherStatistics.RxPktsMulticast})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-broadcast", types.YLeaf{"RxPktsBroadcast", etherStatistics.RxPktsBroadcast})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-under-sized", types.YLeaf{"RxPktsUnderSized", etherStatistics.RxPktsUnderSized})
    etherStatistics.EntityData.Leafs.Append("rx-packets", types.YLeaf{"RxPackets", etherStatistics.RxPackets})
    etherStatistics.EntityData.Leafs.Append("rx-total-bytes", types.YLeaf{"RxTotalBytes", etherStatistics.RxTotalBytes})
    etherStatistics.EntityData.Leafs.Append("rx-bytes-good", types.YLeaf{"RxBytesGood", etherStatistics.RxBytesGood})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-good", types.YLeaf{"RxPktsGood", etherStatistics.RxPktsGood})
    etherStatistics.EntityData.Leafs.Append("tx-bytes-good", types.YLeaf{"TxBytesGood", etherStatistics.TxBytesGood})
    etherStatistics.EntityData.Leafs.Append("tx-pkts-good", types.YLeaf{"TxPktsGood", etherStatistics.TxPktsGood})
    etherStatistics.EntityData.Leafs.Append("rx-recv-fragments", types.YLeaf{"RxRecvFragments", etherStatistics.RxRecvFragments})
    etherStatistics.EntityData.Leafs.Append("rx-pkts64-bytes", types.YLeaf{"RxPkts64Bytes", etherStatistics.RxPkts64Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts65-to127-bytes", types.YLeaf{"RxPkts65To127Bytes", etherStatistics.RxPkts65To127Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts128to255-bytes", types.YLeaf{"RxPkts128to255Bytes", etherStatistics.RxPkts128to255Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts256-to511-bytes", types.YLeaf{"RxPkts256To511Bytes", etherStatistics.RxPkts256To511Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts512-to1023-bytes", types.YLeaf{"RxPkts512To1023Bytes", etherStatistics.RxPkts512To1023Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts1024-to1518-bytes", types.YLeaf{"RxPkts1024To1518Bytes", etherStatistics.RxPkts1024To1518Bytes})
    etherStatistics.EntityData.Leafs.Append("rx-pkts-unicast", types.YLeaf{"RxPktsUnicast", etherStatistics.RxPktsUnicast})
    etherStatistics.EntityData.Leafs.Append("tx-packets", types.YLeaf{"TxPackets", etherStatistics.TxPackets})
    etherStatistics.EntityData.Leafs.Append("tx-total-bytes", types.YLeaf{"TxTotalBytes", etherStatistics.TxTotalBytes})
    etherStatistics.EntityData.Leafs.Append("tx-pkts-under-sized", types.YLeaf{"TxPktsUnderSized", etherStatistics.TxPktsUnderSized})
    etherStatistics.EntityData.Leafs.Append("tx-pkts-over-sized", types.YLeaf{"TxPktsOverSized", etherStatistics.TxPktsOverSized})
    etherStatistics.EntityData.Leafs.Append("tx-fragments", types.YLeaf{"TxFragments", etherStatistics.TxFragments})
    etherStatistics.EntityData.Leafs.Append("tx-jabber", types.YLeaf{"TxJabber", etherStatistics.TxJabber})
    etherStatistics.EntityData.Leafs.Append("tx-bad-fcs", types.YLeaf{"TxBadFcs", etherStatistics.TxBadFcs})
    etherStatistics.EntityData.Leafs.Append("rx-pkt-drop", types.YLeaf{"RxPktDrop", etherStatistics.RxPktDrop})
    etherStatistics.EntityData.Leafs.Append("rx-pause", types.YLeaf{"RxPause", etherStatistics.RxPause})
    etherStatistics.EntityData.Leafs.Append("tx-pause", types.YLeaf{"TxPause", etherStatistics.TxPause})
    etherStatistics.EntityData.Leafs.Append("rx-lldp-pkt", types.YLeaf{"RxLldpPkt", etherStatistics.RxLldpPkt})
    etherStatistics.EntityData.Leafs.Append("rx8021q-pkt", types.YLeaf{"Rx8021qPkt", etherStatistics.Rx8021qPkt})

    etherStatistics.EntityData.YListKeys = []string {}

    return &(etherStatistics.EntityData)
}

