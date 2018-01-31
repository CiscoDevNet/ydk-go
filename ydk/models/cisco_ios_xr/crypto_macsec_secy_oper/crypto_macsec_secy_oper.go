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
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Security Entity.
    Secy Macsec_Secy
}

func (macsec *Macsec) GetFilter() yfilter.YFilter { return macsec.YFilter }

func (macsec *Macsec) SetFilter(yf yfilter.YFilter) { macsec.YFilter = yf }

func (macsec *Macsec) GetGoName(yname string) string {
    if yname == "secy" { return "Secy" }
    return ""
}

func (macsec *Macsec) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-macsec-secy-oper:macsec"
}

func (macsec *Macsec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secy" {
        return &macsec.Secy
    }
    return nil
}

func (macsec *Macsec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["secy"] = &macsec.Secy
    return children
}

func (macsec *Macsec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsec *Macsec) GetBundleName() string { return "cisco_ios_xr" }

func (macsec *Macsec) GetYangName() string { return "macsec" }

func (macsec *Macsec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsec *Macsec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsec *Macsec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsec *Macsec) SetParent(parent types.Entity) { macsec.parent = parent }

func (macsec *Macsec) GetParent() types.Entity { return macsec.parent }

func (macsec *Macsec) GetParentYangName() string { return "Cisco-IOS-XR-crypto-macsec-secy-oper" }

// Macsec_Secy
// MAC Security Entity
type Macsec_Secy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Security Data.
    Interfaces Macsec_Secy_Interfaces
}

func (secy *Macsec_Secy) GetFilter() yfilter.YFilter { return secy.YFilter }

func (secy *Macsec_Secy) SetFilter(yf yfilter.YFilter) { secy.YFilter = yf }

func (secy *Macsec_Secy) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (secy *Macsec_Secy) GetSegmentPath() string {
    return "secy"
}

func (secy *Macsec_Secy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &secy.Interfaces
    }
    return nil
}

func (secy *Macsec_Secy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &secy.Interfaces
    return children
}

func (secy *Macsec_Secy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secy *Macsec_Secy) GetBundleName() string { return "cisco_ios_xr" }

func (secy *Macsec_Secy) GetYangName() string { return "secy" }

func (secy *Macsec_Secy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secy *Macsec_Secy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secy *Macsec_Secy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secy *Macsec_Secy) SetParent(parent types.Entity) { secy.parent = parent }

func (secy *Macsec_Secy) GetParent() types.Entity { return secy.parent }

func (secy *Macsec_Secy) GetParentYangName() string { return "macsec" }

// Macsec_Secy_Interfaces
// MAC Security Data
type Macsec_Secy_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MAC Security Data for the Interface. The type is slice of
    // Macsec_Secy_Interfaces_Interface.
    Interface []Macsec_Secy_Interfaces_Interface
}

func (interfaces *Macsec_Secy_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Macsec_Secy_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Macsec_Secy_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Macsec_Secy_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Macsec_Secy_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Secy_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Macsec_Secy_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Macsec_Secy_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Macsec_Secy_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Macsec_Secy_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Macsec_Secy_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Macsec_Secy_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Macsec_Secy_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Macsec_Secy_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Macsec_Secy_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Macsec_Secy_Interfaces) GetParentYangName() string { return "secy" }

// Macsec_Secy_Interfaces_Interface
// MAC Security Data for the Interface
type Macsec_Secy_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // MACsec Stats.
    Stats Macsec_Secy_Interfaces_Interface_Stats
}

func (self *Macsec_Secy_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Macsec_Secy_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Macsec_Secy_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "stats" { return "Stats" }
    return ""
}

func (self *Macsec_Secy_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
}

func (self *Macsec_Secy_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats" {
        return &self.Stats
    }
    return nil
}

func (self *Macsec_Secy_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats"] = &self.Stats
    return children
}

func (self *Macsec_Secy_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = self.Name
    return leafs
}

func (self *Macsec_Secy_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Macsec_Secy_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Macsec_Secy_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Macsec_Secy_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Macsec_Secy_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Macsec_Secy_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Macsec_Secy_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Macsec_Secy_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Macsec_Secy_Interfaces_Interface_Stats
// MACsec Stats
type Macsec_Secy_Interfaces_Interface_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface stats.
    IntfStats Macsec_Secy_Interfaces_Interface_Stats_IntfStats

    // Tx SC Stats.
    TxScStats Macsec_Secy_Interfaces_Interface_Stats_TxScStats

    // RX SC Stats List. The type is slice of
    // Macsec_Secy_Interfaces_Interface_Stats_RxScStats.
    RxScStats []Macsec_Secy_Interfaces_Interface_Stats_RxScStats
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetGoName(yname string) string {
    if yname == "intf-stats" { return "IntfStats" }
    if yname == "tx-sc-stats" { return "TxScStats" }
    if yname == "rx-sc-stats" { return "RxScStats" }
    return ""
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intf-stats" {
        return &stats.IntfStats
    }
    if childYangName == "tx-sc-stats" {
        return &stats.TxScStats
    }
    if childYangName == "rx-sc-stats" {
        for _, c := range stats.RxScStats {
            if stats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Secy_Interfaces_Interface_Stats_RxScStats{}
        stats.RxScStats = append(stats.RxScStats, child)
        return &stats.RxScStats[len(stats.RxScStats)-1]
    }
    return nil
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intf-stats"] = &stats.IntfStats
    children["tx-sc-stats"] = &stats.TxScStats
    for i := range stats.RxScStats {
        children[stats.RxScStats[i].GetSegmentPath()] = &stats.RxScStats[i]
    }
    return children
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetYangName() string { return "stats" }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetParent() types.Entity { return stats.parent }

func (stats *Macsec_Secy_Interfaces_Interface_Stats) GetParentYangName() string { return "interface" }

// Macsec_Secy_Interfaces_Interface_Stats_IntfStats
// Interface stats
type Macsec_Secy_Interfaces_Interface_Stats_IntfStats struct {
    parent types.Entity
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

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetFilter() yfilter.YFilter { return intfStats.YFilter }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) SetFilter(yf yfilter.YFilter) { intfStats.YFilter = yf }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetGoName(yname string) string {
    if yname == "in-pkts-untagged" { return "InPktsUntagged" }
    if yname == "in-pkts-no-tag" { return "InPktsNoTag" }
    if yname == "in-pkts-bad-tag" { return "InPktsBadTag" }
    if yname == "in-pkts-unknown-sci" { return "InPktsUnknownSci" }
    if yname == "in-pkts-no-sci" { return "InPktsNoSci" }
    if yname == "in-pkts-overrun" { return "InPktsOverrun" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    if yname == "in-octets-decrypted" { return "InOctetsDecrypted" }
    if yname == "out-pkts-untagged" { return "OutPktsUntagged" }
    if yname == "out-pkts-too-long" { return "OutPktsTooLong" }
    if yname == "out-octets-protected" { return "OutOctetsProtected" }
    if yname == "out-octets-encrypted" { return "OutOctetsEncrypted" }
    return ""
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetSegmentPath() string {
    return "intf-stats"
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-untagged"] = intfStats.InPktsUntagged
    leafs["in-pkts-no-tag"] = intfStats.InPktsNoTag
    leafs["in-pkts-bad-tag"] = intfStats.InPktsBadTag
    leafs["in-pkts-unknown-sci"] = intfStats.InPktsUnknownSci
    leafs["in-pkts-no-sci"] = intfStats.InPktsNoSci
    leafs["in-pkts-overrun"] = intfStats.InPktsOverrun
    leafs["in-octets-validated"] = intfStats.InOctetsValidated
    leafs["in-octets-decrypted"] = intfStats.InOctetsDecrypted
    leafs["out-pkts-untagged"] = intfStats.OutPktsUntagged
    leafs["out-pkts-too-long"] = intfStats.OutPktsTooLong
    leafs["out-octets-protected"] = intfStats.OutOctetsProtected
    leafs["out-octets-encrypted"] = intfStats.OutOctetsEncrypted
    return leafs
}

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetBundleName() string { return "cisco_ios_xr" }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetYangName() string { return "intf-stats" }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) SetParent(parent types.Entity) { intfStats.parent = parent }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetParent() types.Entity { return intfStats.parent }

func (intfStats *Macsec_Secy_Interfaces_Interface_Stats_IntfStats) GetParentYangName() string { return "stats" }

// Macsec_Secy_Interfaces_Interface_Stats_TxScStats
// Tx SC Stats
type Macsec_Secy_Interfaces_Interface_Stats_TxScStats struct {
    parent types.Entity
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

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetFilter() yfilter.YFilter { return txScStats.YFilter }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) SetFilter(yf yfilter.YFilter) { txScStats.YFilter = yf }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetGoName(yname string) string {
    if yname == "tx-sci" { return "TxSci" }
    if yname == "out-pkts-protected" { return "OutPktsProtected" }
    if yname == "out-pkts-encrypted" { return "OutPktsEncrypted" }
    if yname == "out-octets-protected" { return "OutOctetsProtected" }
    if yname == "out-octets-encrypted" { return "OutOctetsEncrypted" }
    if yname == "out-pkts-too-long" { return "OutPktsTooLong" }
    if yname == "txsa-stat" { return "TxsaStat" }
    return ""
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetSegmentPath() string {
    return "tx-sc-stats"
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "txsa-stat" {
        for _, c := range txScStats.TxsaStat {
            if txScStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat{}
        txScStats.TxsaStat = append(txScStats.TxsaStat, child)
        return &txScStats.TxsaStat[len(txScStats.TxsaStat)-1]
    }
    return nil
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range txScStats.TxsaStat {
        children[txScStats.TxsaStat[i].GetSegmentPath()] = &txScStats.TxsaStat[i]
    }
    return children
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-sci"] = txScStats.TxSci
    leafs["out-pkts-protected"] = txScStats.OutPktsProtected
    leafs["out-pkts-encrypted"] = txScStats.OutPktsEncrypted
    leafs["out-octets-protected"] = txScStats.OutOctetsProtected
    leafs["out-octets-encrypted"] = txScStats.OutOctetsEncrypted
    leafs["out-pkts-too-long"] = txScStats.OutPktsTooLong
    return leafs
}

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetBundleName() string { return "cisco_ios_xr" }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetYangName() string { return "tx-sc-stats" }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) SetParent(parent types.Entity) { txScStats.parent = parent }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetParent() types.Entity { return txScStats.parent }

func (txScStats *Macsec_Secy_Interfaces_Interface_Stats_TxScStats) GetParentYangName() string { return "stats" }

// Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat
// tx sa stats
type Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat struct {
    parent types.Entity
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

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetFilter() yfilter.YFilter { return txsaStat.YFilter }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) SetFilter(yf yfilter.YFilter) { txsaStat.YFilter = yf }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetGoName(yname string) string {
    if yname == "out-pkts-protected" { return "OutPktsProtected" }
    if yname == "out-pkts-encrypted" { return "OutPktsEncrypted" }
    if yname == "next-pn" { return "NextPn" }
    return ""
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetSegmentPath() string {
    return "txsa-stat"
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-pkts-protected"] = txsaStat.OutPktsProtected
    leafs["out-pkts-encrypted"] = txsaStat.OutPktsEncrypted
    leafs["next-pn"] = txsaStat.NextPn
    return leafs
}

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetBundleName() string { return "cisco_ios_xr" }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetYangName() string { return "txsa-stat" }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) SetParent(parent types.Entity) { txsaStat.parent = parent }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetParent() types.Entity { return txsaStat.parent }

func (txsaStat *Macsec_Secy_Interfaces_Interface_Stats_TxScStats_TxsaStat) GetParentYangName() string { return "tx-sc-stats" }

// Macsec_Secy_Interfaces_Interface_Stats_RxScStats
// RX SC Stats List
type Macsec_Secy_Interfaces_Interface_Stats_RxScStats struct {
    parent types.Entity
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

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetFilter() yfilter.YFilter { return rxScStats.YFilter }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) SetFilter(yf yfilter.YFilter) { rxScStats.YFilter = yf }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetGoName(yname string) string {
    if yname == "rx-sci" { return "RxSci" }
    if yname == "in-pkts-unchecked" { return "InPktsUnchecked" }
    if yname == "in-pkts-delayed" { return "InPktsDelayed" }
    if yname == "in-pkts-late" { return "InPktsLate" }
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "in-pkts-untagged-hit" { return "InPktsUntaggedHit" }
    if yname == "in-octets-validated" { return "InOctetsValidated" }
    if yname == "in-octets-decrypted" { return "InOctetsDecrypted" }
    if yname == "rxsa-stat" { return "RxsaStat" }
    return ""
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetSegmentPath() string {
    return "rx-sc-stats"
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rxsa-stat" {
        for _, c := range rxScStats.RxsaStat {
            if rxScStats.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat{}
        rxScStats.RxsaStat = append(rxScStats.RxsaStat, child)
        return &rxScStats.RxsaStat[len(rxScStats.RxsaStat)-1]
    }
    return nil
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rxScStats.RxsaStat {
        children[rxScStats.RxsaStat[i].GetSegmentPath()] = &rxScStats.RxsaStat[i]
    }
    return children
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rx-sci"] = rxScStats.RxSci
    leafs["in-pkts-unchecked"] = rxScStats.InPktsUnchecked
    leafs["in-pkts-delayed"] = rxScStats.InPktsDelayed
    leafs["in-pkts-late"] = rxScStats.InPktsLate
    leafs["in-pkts-ok"] = rxScStats.InPktsOk
    leafs["in-pkts-invalid"] = rxScStats.InPktsInvalid
    leafs["in-pkts-not-valid"] = rxScStats.InPktsNotValid
    leafs["in-pkts-not-using-sa"] = rxScStats.InPktsNotUsingSa
    leafs["in-pkts-unused-sa"] = rxScStats.InPktsUnusedSa
    leafs["in-pkts-untagged-hit"] = rxScStats.InPktsUntaggedHit
    leafs["in-octets-validated"] = rxScStats.InOctetsValidated
    leafs["in-octets-decrypted"] = rxScStats.InOctetsDecrypted
    return leafs
}

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetBundleName() string { return "cisco_ios_xr" }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetYangName() string { return "rx-sc-stats" }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) SetParent(parent types.Entity) { rxScStats.parent = parent }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetParent() types.Entity { return rxScStats.parent }

func (rxScStats *Macsec_Secy_Interfaces_Interface_Stats_RxScStats) GetParentYangName() string { return "stats" }

// Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat
// rxsa stats
type Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat struct {
    parent types.Entity
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

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetFilter() yfilter.YFilter { return rxsaStat.YFilter }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) SetFilter(yf yfilter.YFilter) { rxsaStat.YFilter = yf }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetGoName(yname string) string {
    if yname == "in-pkts-ok" { return "InPktsOk" }
    if yname == "in-pkts-invalid" { return "InPktsInvalid" }
    if yname == "in-pkts-not-valid" { return "InPktsNotValid" }
    if yname == "in-pkts-not-using-sa" { return "InPktsNotUsingSa" }
    if yname == "in-pkts-unused-sa" { return "InPktsUnusedSa" }
    if yname == "next-pn" { return "NextPn" }
    return ""
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetSegmentPath() string {
    return "rxsa-stat"
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-pkts-ok"] = rxsaStat.InPktsOk
    leafs["in-pkts-invalid"] = rxsaStat.InPktsInvalid
    leafs["in-pkts-not-valid"] = rxsaStat.InPktsNotValid
    leafs["in-pkts-not-using-sa"] = rxsaStat.InPktsNotUsingSa
    leafs["in-pkts-unused-sa"] = rxsaStat.InPktsUnusedSa
    leafs["next-pn"] = rxsaStat.NextPn
    return leafs
}

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetBundleName() string { return "cisco_ios_xr" }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetYangName() string { return "rxsa-stat" }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) SetParent(parent types.Entity) { rxsaStat.parent = parent }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetParent() types.Entity { return rxsaStat.parent }

func (rxsaStat *Macsec_Secy_Interfaces_Interface_Stats_RxScStats_RxsaStat) GetParentYangName() string { return "rx-sc-stats" }

