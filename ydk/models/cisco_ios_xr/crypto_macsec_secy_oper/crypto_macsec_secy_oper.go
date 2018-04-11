// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-secy package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec: Macsec operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_secy_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_secy_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-secy-oper macsec}", reflect.TypeOf(Macsec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-secy-oper:macsec", reflect.TypeOf(Macsec{}))
}

// Macsec
// Macsec operational data
type Macsec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Security Entity.
    Secy Macsec_Secy
}

func (macsec *Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-macsec-secy-oper"
    macsec.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-macsec-secy-oper:macsec"
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = make(map[string]types.YChild)
    macsec.EntityData.Children["secy"] = types.YChild{"Secy", &macsec.Secy}
    macsec.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsec.EntityData)
}

// Macsec_Secy
// MAC Security Entity
type Macsec_Secy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Security Data.
    Interfaces Macsec_Secy_Interfaces
}

func (secy *Macsec_Secy) GetEntityData() *types.CommonEntityData {
    secy.EntityData.YFilter = secy.YFilter
    secy.EntityData.YangName = "secy"
    secy.EntityData.BundleName = "cisco_ios_xr"
    secy.EntityData.ParentYangName = "macsec"
    secy.EntityData.SegmentPath = "secy"
    secy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secy.EntityData.Children = make(map[string]types.YChild)
    secy.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &secy.Interfaces}
    secy.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(secy.EntityData)
}

// Macsec_Secy_Interfaces
// MAC Security Data
type Macsec_Secy_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC Security Data for the Interface. The type is slice of
    // Macsec_Secy_Interfaces_Interface_.
    Interface_ []Macsec_Secy_Interfaces_Interface
}

func (interfaces *Macsec_Secy_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "secy"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Macsec_Secy_Interfaces_Interface
// MAC Security Data for the Interface
type Macsec_Secy_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // MACsec Stats.
    Stats Macsec_Secy_Interfaces_Interface_Stats
}

func (self *Macsec_Secy_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["stats"] = types.YChild{"Stats", &self.Stats}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats
// MACsec Stats
type Macsec_Secy_Interfaces_Interface_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface stats.
    IntfStats Macsec_Secy_Interfaces_Interface_Stats_IntfStats

    // Tx SC Stats.
    TxScStats Macsec_Secy_Interfaces_Interface_Stats_TxScStats

    // RX SC Stats List. The type is slice of
    // Macsec_Secy_Interfaces_Interface_Stats_RxScStats.
    RxScStats []Macsec_Secy_Interfaces_Interface_Stats_RxScStats
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "interface"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = make(map[string]types.YChild)
    stats.EntityData.Children["intf-stats"] = types.YChild{"IntfStats", &stats.IntfStats}
    stats.EntityData.Children["tx-sc-stats"] = types.YChild{"TxScStats", &stats.TxScStats}
    stats.EntityData.Children["rx-sc-stats"] = types.YChild{"RxScStats", nil}
    for i := range stats.RxScStats {
        stats.EntityData.Children[types.GetSegmentPath(&stats.RxScStats[i])] = types.YChild{"RxScStats", &stats.RxScStats[i]}
    }
    stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stats.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats_IntfStats
// Interface stats
type Macsec_Secy_Interfaces_Interface_Stats_IntfStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // InPktsUntagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUntagged interface{}

    // InPktsNoTag. The type is interface{} with range: 0..18446744073709551615.
    InPktsNoTag interface{}

    // InPktsBadTag. The type is interface{} with range: 0..18446744073709551615.
    InPktsBadTag interface{}

    // InPktsUnknownSCI. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnknownSci interface{}

    // InPktsNoSCI. The type is interface{} with range: 0..18446744073709551615.
    InPktsNoSci interface{}

    // InPktsOverrun. The type is interface{} with range: 0..18446744073709551615.
    InPktsOverrun interface{}

    // InOctetsValidated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // InOctetsDecrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}

    // OutPktsUntagged. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsUntagged interface{}

    // OutPktsTooLong. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // OutOctetsProtected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // OutOctetsEncrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetEntityData() *types.CommonEntityData {
    intfStats.EntityData.YFilter = intfStats.YFilter
    intfStats.EntityData.YangName = "intf-stats"
    intfStats.EntityData.BundleName = "cisco_ios_xr"
    intfStats.EntityData.ParentYangName = "stats"
    intfStats.EntityData.SegmentPath = "intf-stats"
    intfStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfStats.EntityData.Children = make(map[string]types.YChild)
    intfStats.EntityData.Leafs = make(map[string]types.YLeaf)
    intfStats.EntityData.Leafs["in-pkts-untagged"] = types.YLeaf{"InPktsUntagged", intfStats.InPktsUntagged}
    intfStats.EntityData.Leafs["in-pkts-no-tag"] = types.YLeaf{"InPktsNoTag", intfStats.InPktsNoTag}
    intfStats.EntityData.Leafs["in-pkts-bad-tag"] = types.YLeaf{"InPktsBadTag", intfStats.InPktsBadTag}
    intfStats.EntityData.Leafs["in-pkts-unknown-sci"] = types.YLeaf{"InPktsUnknownSci", intfStats.InPktsUnknownSci}
    intfStats.EntityData.Leafs["in-pkts-no-sci"] = types.YLeaf{"InPktsNoSci", intfStats.InPktsNoSci}
    intfStats.EntityData.Leafs["in-pkts-overrun"] = types.YLeaf{"InPktsOverrun", intfStats.InPktsOverrun}
    intfStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", intfStats.InOctetsValidated}
    intfStats.EntityData.Leafs["in-octets-decrypted"] = types.YLeaf{"InOctetsDecrypted", intfStats.InOctetsDecrypted}
    intfStats.EntityData.Leafs["out-pkts-untagged"] = types.YLeaf{"OutPktsUntagged", intfStats.OutPktsUntagged}
    intfStats.EntityData.Leafs["out-pkts-too-long"] = types.YLeaf{"OutPktsTooLong", intfStats.OutPktsTooLong}
    intfStats.EntityData.Leafs["out-octets-protected"] = types.YLeaf{"OutOctetsProtected", intfStats.OutOctetsProtected}
    intfStats.EntityData.Leafs["out-octets-encrypted"] = types.YLeaf{"OutOctetsEncrypted", intfStats.OutOctetsEncrypted}
    return &(intfStats.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats_TxScStats
// Tx SC Stats
type Macsec_Secy_Interfaces_Interface_Stats_TxScStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SCI. The type is interface{} with range: 0..18446744073709551615.
    TxSci interface{}

    // OutPktsProtected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // OutPktsEncrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // OutOctetsProtected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // OutOctetsEncrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}

    // OutPktsTooLong. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // tx sa stats. The type is slice of
    // Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat.
    TxsaStat []Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetEntityData() *types.CommonEntityData {
    txScStats.EntityData.YFilter = txScStats.YFilter
    txScStats.EntityData.YangName = "tx-sc-stats"
    txScStats.EntityData.BundleName = "cisco_ios_xr"
    txScStats.EntityData.ParentYangName = "stats"
    txScStats.EntityData.SegmentPath = "tx-sc-stats"
    txScStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txScStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txScStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txScStats.EntityData.Children = make(map[string]types.YChild)
    txScStats.EntityData.Children["txsa-stat"] = types.YChild{"TxsaStat", nil}
    for i := range txScStats.TxsaStat {
        txScStats.EntityData.Children[types.GetSegmentPath(&txScStats.TxsaStat[i])] = types.YChild{"TxsaStat", &txScStats.TxsaStat[i]}
    }
    txScStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txScStats.EntityData.Leafs["tx-sci"] = types.YLeaf{"TxSci", txScStats.TxSci}
    txScStats.EntityData.Leafs["out-pkts-protected"] = types.YLeaf{"OutPktsProtected", txScStats.OutPktsProtected}
    txScStats.EntityData.Leafs["out-pkts-encrypted"] = types.YLeaf{"OutPktsEncrypted", txScStats.OutPktsEncrypted}
    txScStats.EntityData.Leafs["out-octets-protected"] = types.YLeaf{"OutOctetsProtected", txScStats.OutOctetsProtected}
    txScStats.EntityData.Leafs["out-octets-encrypted"] = types.YLeaf{"OutOctetsEncrypted", txScStats.OutOctetsEncrypted}
    txScStats.EntityData.Leafs["out-pkts-too-long"] = types.YLeaf{"OutPktsTooLong", txScStats.OutPktsTooLong}
    return &(txScStats.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat
// tx sa stats
type Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OutPktsProtected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // OutPktsEncrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // NextPN. The type is interface{} with range: 0..18446744073709551615.
    NextPn interface{}
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetEntityData() *types.CommonEntityData {
    txsaStat.EntityData.YFilter = txsaStat.YFilter
    txsaStat.EntityData.YangName = "txsa-stat"
    txsaStat.EntityData.BundleName = "cisco_ios_xr"
    txsaStat.EntityData.ParentYangName = "tx-sc-stats"
    txsaStat.EntityData.SegmentPath = "txsa-stat"
    txsaStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txsaStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txsaStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txsaStat.EntityData.Children = make(map[string]types.YChild)
    txsaStat.EntityData.Leafs = make(map[string]types.YLeaf)
    txsaStat.EntityData.Leafs["out-pkts-protected"] = types.YLeaf{"OutPktsProtected", txsaStat.OutPktsProtected}
    txsaStat.EntityData.Leafs["out-pkts-encrypted"] = types.YLeaf{"OutPktsEncrypted", txsaStat.OutPktsEncrypted}
    txsaStat.EntityData.Leafs["next-pn"] = types.YLeaf{"NextPn", txsaStat.NextPn}
    return &(txsaStat.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats_RxScStats
// RX SC Stats List
type Macsec_Secy_Interfaces_Interface_Stats_RxScStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx SCI. The type is interface{} with range: 0..18446744073709551615.
    RxSci interface{}

    // InPktsUnchecked. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnchecked interface{}

    // InPktsDelayed. The type is interface{} with range: 0..18446744073709551615.
    InPktsDelayed interface{}

    // InPktsLate. The type is interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // InPktsOK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // InPktsInvalid. The type is interface{} with range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // InPktsNotValid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // InPktsNotUsingSA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // InPktsUnusedSA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // InPktsUntaggedHit. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUntaggedHit interface{}

    // InOctetsValidated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // InOctetsDecrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}

    // rxsa stats. The type is slice of
    // Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat.
    RxsaStat []Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetEntityData() *types.CommonEntityData {
    rxScStats.EntityData.YFilter = rxScStats.YFilter
    rxScStats.EntityData.YangName = "rx-sc-stats"
    rxScStats.EntityData.BundleName = "cisco_ios_xr"
    rxScStats.EntityData.ParentYangName = "stats"
    rxScStats.EntityData.SegmentPath = "rx-sc-stats"
    rxScStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxScStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxScStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxScStats.EntityData.Children = make(map[string]types.YChild)
    rxScStats.EntityData.Children["rxsa-stat"] = types.YChild{"RxsaStat", nil}
    for i := range rxScStats.RxsaStat {
        rxScStats.EntityData.Children[types.GetSegmentPath(&rxScStats.RxsaStat[i])] = types.YChild{"RxsaStat", &rxScStats.RxsaStat[i]}
    }
    rxScStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxScStats.EntityData.Leafs["rx-sci"] = types.YLeaf{"RxSci", rxScStats.RxSci}
    rxScStats.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxScStats.InPktsUnchecked}
    rxScStats.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxScStats.InPktsDelayed}
    rxScStats.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxScStats.InPktsLate}
    rxScStats.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxScStats.InPktsOk}
    rxScStats.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxScStats.InPktsInvalid}
    rxScStats.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxScStats.InPktsNotValid}
    rxScStats.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxScStats.InPktsNotUsingSa}
    rxScStats.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxScStats.InPktsUnusedSa}
    rxScStats.EntityData.Leafs["in-pkts-untagged-hit"] = types.YLeaf{"InPktsUntaggedHit", rxScStats.InPktsUntaggedHit}
    rxScStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxScStats.InOctetsValidated}
    rxScStats.EntityData.Leafs["in-octets-decrypted"] = types.YLeaf{"InOctetsDecrypted", rxScStats.InOctetsDecrypted}
    return &(rxScStats.EntityData)
}

// Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat
// rxsa stats
type Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // InPktsOK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // InPktsInvalid. The type is interface{} with range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // InPktsNotValid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // InPktsNotUsingSA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // InPktsUnusedSA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // NextPN. The type is interface{} with range: 0..18446744073709551615.
    NextPn interface{}
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetEntityData() *types.CommonEntityData {
    rxsaStat.EntityData.YFilter = rxsaStat.YFilter
    rxsaStat.EntityData.YangName = "rxsa-stat"
    rxsaStat.EntityData.BundleName = "cisco_ios_xr"
    rxsaStat.EntityData.ParentYangName = "rx-sc-stats"
    rxsaStat.EntityData.SegmentPath = "rxsa-stat"
    rxsaStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxsaStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxsaStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxsaStat.EntityData.Children = make(map[string]types.YChild)
    rxsaStat.EntityData.Leafs = make(map[string]types.YLeaf)
    rxsaStat.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxsaStat.InPktsOk}
    rxsaStat.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxsaStat.InPktsInvalid}
    rxsaStat.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxsaStat.InPktsNotValid}
    rxsaStat.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxsaStat.InPktsNotUsingSa}
    rxsaStat.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxsaStat.InPktsUnusedSa}
    rxsaStat.EntityData.Leafs["next-pn"] = types.YLeaf{"NextPn", rxsaStat.NextPn}
    return &(rxsaStat.EntityData)
}

