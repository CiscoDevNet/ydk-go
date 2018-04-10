// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-pl package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec-platform: MACSec operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_pl_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_pl_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-pl-oper macsec-platform}", reflect.TypeOf(MacsecPlatform{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-pl-oper:macsec-platform", reflect.TypeOf(MacsecPlatform{}))
}

// MacsecCard represents Macsec card
type MacsecCard string

const (
    // macsec none
    MacsecCard_macsec_none MacsecCard = "macsec-none"

    // macsec msfpga
    MacsecCard_macsec_msfpga MacsecCard = "macsec-msfpga"

    // macsec xlmsfpga
    MacsecCard_macsec_xlmsfpga MacsecCard = "macsec-xlmsfpga"

    // macsec apm
    MacsecCard_macsec_apm MacsecCard = "macsec-apm"
)

// MacsecPlatform
// MACSec operational data
type MacsecPlatform struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NodeTable for all the nodes.
    Nodes MacsecPlatform_Nodes
}

func (macsecPlatform *MacsecPlatform) GetEntityData() *types.CommonEntityData {
    macsecPlatform.EntityData.YFilter = macsecPlatform.YFilter
    macsecPlatform.EntityData.YangName = "macsec-platform"
    macsecPlatform.EntityData.BundleName = "cisco_ios_xr"
    macsecPlatform.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-macsec-pl-oper"
    macsecPlatform.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-macsec-pl-oper:macsec-platform"
    macsecPlatform.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecPlatform.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecPlatform.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecPlatform.EntityData.Children = make(map[string]types.YChild)
    macsecPlatform.EntityData.Children["nodes"] = types.YChild{"Nodes", &macsecPlatform.Nodes}
    macsecPlatform.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsecPlatform.EntityData)
}

// MacsecPlatform_Nodes
// NodeTable for all the nodes
type MacsecPlatform_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node where macsec interfaces exist. The type is slice of
    // MacsecPlatform_Nodes_Node.
    Node []MacsecPlatform_Nodes_Node
}

func (nodes *MacsecPlatform_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "macsec-platform"
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

// MacsecPlatform_Nodes_Node
// Node where macsec interfaces exist
type MacsecPlatform_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Table of Interfaces.
    Interfaces MacsecPlatform_Nodes_Node_Interfaces
}

func (node *MacsecPlatform_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces
// Table of Interfaces
type MacsecPlatform_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Where Macsec is configured. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_.
    Interface_ []MacsecPlatform_Nodes_Node_Interfaces_Interface
}

func (interfaces *MacsecPlatform_Nodes_Node_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "node"
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

// MacsecPlatform_Nodes_Node_Interfaces_Interface
// Interface Where Macsec is configured
type MacsecPlatform_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Value. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // The Hardware Statistics.
    HwStatistics MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics

    // Table of Hardware SAs.
    HwSas MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas

    // Table of Hardware Flows.
    HwFlowS MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS

    // The Software Statistics.
    SwStatistics MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics
}

func (self *MacsecPlatform_Nodes_Node_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["hw-statistics"] = types.YChild{"HwStatistics", &self.HwStatistics}
    self.EntityData.Children["hw-sas"] = types.YChild{"HwSas", &self.HwSas}
    self.EntityData.Children["hw-flow-s"] = types.YChild{"HwFlowS", &self.HwFlowS}
    self.EntityData.Children["sw-statistics"] = types.YChild{"SwStatistics", &self.SwStatistics}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics
// The Hardware Statistics
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext
}

func (hwStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics) GetEntityData() *types.CommonEntityData {
    hwStatistics.EntityData.YFilter = hwStatistics.YFilter
    hwStatistics.EntityData.YangName = "hw-statistics"
    hwStatistics.EntityData.BundleName = "cisco_ios_xr"
    hwStatistics.EntityData.ParentYangName = "interface"
    hwStatistics.EntityData.SegmentPath = "hw-statistics"
    hwStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwStatistics.EntityData.Children = make(map[string]types.YChild)
    hwStatistics.EntityData.Children["ext"] = types.YChild{"Ext", &hwStatistics.Ext}
    hwStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwStatistics.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type_ interface{}

    // MSFPGA Stats.
    MsfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats

    // XLFPGA Stats.
    XlfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats

    // ES200 Stats.
    Es200Stats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext) GetEntityData() *types.CommonEntityData {
    ext.EntityData.YFilter = ext.YFilter
    ext.EntityData.YangName = "ext"
    ext.EntityData.BundleName = "cisco_ios_xr"
    ext.EntityData.ParentYangName = "hw-statistics"
    ext.EntityData.SegmentPath = "ext"
    ext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ext.EntityData.Children = make(map[string]types.YChild)
    ext.EntityData.Children["msfpga-stats"] = types.YChild{"MsfpgaStats", &ext.MsfpgaStats}
    ext.EntityData.Children["xlfpga-stats"] = types.YChild{"XlfpgaStats", &ext.XlfpgaStats}
    ext.EntityData.Children["es200-stats"] = types.YChild{"Es200Stats", &ext.Es200Stats}
    ext.EntityData.Leafs = make(map[string]types.YLeaf)
    ext.EntityData.Leafs["type"] = types.YLeaf{"Type_", ext.Type_}
    return &(ext.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats
// MSFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats) GetEntityData() *types.CommonEntityData {
    msfpgaStats.EntityData.YFilter = msfpgaStats.YFilter
    msfpgaStats.EntityData.YangName = "msfpga-stats"
    msfpgaStats.EntityData.BundleName = "cisco_ios_xr"
    msfpgaStats.EntityData.ParentYangName = "ext"
    msfpgaStats.EntityData.SegmentPath = "msfpga-stats"
    msfpgaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    msfpgaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    msfpgaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    msfpgaStats.EntityData.Children = make(map[string]types.YChild)
    msfpgaStats.EntityData.Children["tx-sa-stats"] = types.YChild{"TxSaStats", &msfpgaStats.TxSaStats}
    msfpgaStats.EntityData.Children["rx-sa-stats"] = types.YChild{"RxSaStats", &msfpgaStats.RxSaStats}
    msfpgaStats.EntityData.Children["tx-interface-macsec-stats"] = types.YChild{"TxInterfaceMacsecStats", &msfpgaStats.TxInterfaceMacsecStats}
    msfpgaStats.EntityData.Children["rx-interface-macsec-stats"] = types.YChild{"RxInterfaceMacsecStats", &msfpgaStats.RxInterfaceMacsecStats}
    msfpgaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msfpgaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Pkts Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // Tx Pkts Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // Tx Octets Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxSaStats) GetEntityData() *types.CommonEntityData {
    txSaStats.EntityData.YFilter = txSaStats.YFilter
    txSaStats.EntityData.YangName = "tx-sa-stats"
    txSaStats.EntityData.BundleName = "cisco_ios_xr"
    txSaStats.EntityData.ParentYangName = "msfpga-stats"
    txSaStats.EntityData.SegmentPath = "tx-sa-stats"
    txSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSaStats.EntityData.Children = make(map[string]types.YChild)
    txSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txSaStats.EntityData.Leafs["out-pkts-protected"] = types.YLeaf{"OutPktsProtected", txSaStats.OutPktsProtected}
    txSaStats.EntityData.Leafs["out-pkts-encrypted"] = types.YLeaf{"OutPktsEncrypted", txSaStats.OutPktsEncrypted}
    txSaStats.EntityData.Leafs["out-octets-protected"] = types.YLeaf{"OutOctetsProtected", txSaStats.OutOctetsProtected}
    txSaStats.EntityData.Leafs["out-octets-encrypted"] = types.YLeaf{"OutOctetsEncrypted", txSaStats.OutOctetsEncrypted}
    return &(txSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // Rx Pkts Not Valid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsInvalid interface{}

    // Rx Pkts OK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // Rx Pkts Delayed. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsDelayed interface{}

    // Rx Pkts Late. The type is interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // Rx Pkts Unchecked. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnchecked interface{}

    // Rx Octets Validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxSaStats) GetEntityData() *types.CommonEntityData {
    rxSaStats.EntityData.YFilter = rxSaStats.YFilter
    rxSaStats.EntityData.YangName = "rx-sa-stats"
    rxSaStats.EntityData.BundleName = "cisco_ios_xr"
    rxSaStats.EntityData.ParentYangName = "msfpga-stats"
    rxSaStats.EntityData.SegmentPath = "rx-sa-stats"
    rxSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStats.EntityData.Children = make(map[string]types.YChild)
    rxSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStats.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxSaStats.InPktsUnusedSa}
    rxSaStats.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxSaStats.InPktsNotUsingSa}
    rxSaStats.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxSaStats.InPktsNotValid}
    rxSaStats.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxSaStats.InPktsInvalid}
    rxSaStats.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxSaStats.InPktsOk}
    rxSaStats.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxSaStats.InPktsDelayed}
    rxSaStats.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxSaStats.InPktsLate}
    rxSaStats.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxSaStats.InPktsUnchecked}
    rxSaStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxSaStats.InOctetsValidated}
    rxSaStats.EntityData.Leafs["in-octets-decrypted"] = types.YLeaf{"InOctetsDecrypted", rxSaStats.InOctetsDecrypted}
    return &(rxSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUncontrolled interface{}

    // Tx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUntagged interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktTooLong interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    txInterfaceMacsecStats.EntityData.YFilter = txInterfaceMacsecStats.YFilter
    txInterfaceMacsecStats.EntityData.YangName = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txInterfaceMacsecStats.EntityData.ParentYangName = "msfpga-stats"
    txInterfaceMacsecStats.EntityData.SegmentPath = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-uncontrolled"] = types.YLeaf{"OutPktUncontrolled", txInterfaceMacsecStats.OutPktUncontrolled}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-untagged"] = types.YLeaf{"OutPktUntagged", txInterfaceMacsecStats.OutPktUntagged}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-too-long"] = types.YLeaf{"OutPktTooLong", txInterfaceMacsecStats.OutPktTooLong}
    return &(txInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUntagged interface{}

    // Rx Pkts Notag. The type is interface{} with range: 0..18446744073709551615.
    InPktNotag interface{}

    // Rx Pkts Bad tag. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktBadTag interface{}

    // Rx Pkts No Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // Rx Pkts Unknown Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUnknownSci interface{}

    // Rx Pkts Tagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktTagged interface{}

    // Rx Pkts Over Run. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktOverrun interface{}

    // Rx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUncontrolled interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    rxInterfaceMacsecStats.EntityData.YFilter = rxInterfaceMacsecStats.YFilter
    rxInterfaceMacsecStats.EntityData.YangName = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxInterfaceMacsecStats.EntityData.ParentYangName = "msfpga-stats"
    rxInterfaceMacsecStats.EntityData.SegmentPath = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-untagged"] = types.YLeaf{"InPktUntagged", rxInterfaceMacsecStats.InPktUntagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-notag"] = types.YLeaf{"InPktNotag", rxInterfaceMacsecStats.InPktNotag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-bad-tag"] = types.YLeaf{"InPktBadTag", rxInterfaceMacsecStats.InPktBadTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-sci"] = types.YLeaf{"InPktNoSci", rxInterfaceMacsecStats.InPktNoSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-unknown-sci"] = types.YLeaf{"InPktUnknownSci", rxInterfaceMacsecStats.InPktUnknownSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-tagged"] = types.YLeaf{"InPktTagged", rxInterfaceMacsecStats.InPktTagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-overrun"] = types.YLeaf{"InPktOverrun", rxInterfaceMacsecStats.InPktOverrun}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-uncontrolled"] = types.YLeaf{"InPktUncontrolled", rxInterfaceMacsecStats.InPktUncontrolled}
    return &(rxInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats
// XLFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SC and SA Level Stats.
    MacsecTxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats

    // Rx SC and SA Level Stats.
    MacsecRxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats) GetEntityData() *types.CommonEntityData {
    xlfpgaStats.EntityData.YFilter = xlfpgaStats.YFilter
    xlfpgaStats.EntityData.YangName = "xlfpga-stats"
    xlfpgaStats.EntityData.BundleName = "cisco_ios_xr"
    xlfpgaStats.EntityData.ParentYangName = "ext"
    xlfpgaStats.EntityData.SegmentPath = "xlfpga-stats"
    xlfpgaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xlfpgaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xlfpgaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xlfpgaStats.EntityData.Children = make(map[string]types.YChild)
    xlfpgaStats.EntityData.Children["macsec-tx-stats"] = types.YChild{"MacsecTxStats", &xlfpgaStats.MacsecTxStats}
    xlfpgaStats.EntityData.Children["macsec-rx-stats"] = types.YChild{"MacsecRxStats", &xlfpgaStats.MacsecRxStats}
    xlfpgaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(xlfpgaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats
// Tx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedOctets interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    ScToolongPkts interface{}

    // Tx packets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedPkts interface{}

    // Tx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Tx Overrun Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Tx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Tx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Tx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Current Tx AN. The type is interface{} with range: 0..18446744073709551615.
    CurrentAn interface{}

    // Current Tx SA Encrypted Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SaEncryptedPkts interface{}
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetEntityData() *types.CommonEntityData {
    macsecTxStats.EntityData.YFilter = macsecTxStats.YFilter
    macsecTxStats.EntityData.YangName = "macsec-tx-stats"
    macsecTxStats.EntityData.BundleName = "cisco_ios_xr"
    macsecTxStats.EntityData.ParentYangName = "xlfpga-stats"
    macsecTxStats.EntityData.SegmentPath = "macsec-tx-stats"
    macsecTxStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecTxStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecTxStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecTxStats.EntityData.Children = make(map[string]types.YChild)
    macsecTxStats.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecTxStats.EntityData.Leafs["sc-encrypted-octets"] = types.YLeaf{"ScEncryptedOctets", macsecTxStats.ScEncryptedOctets}
    macsecTxStats.EntityData.Leafs["sc-toolong-pkts"] = types.YLeaf{"ScToolongPkts", macsecTxStats.ScToolongPkts}
    macsecTxStats.EntityData.Leafs["sc-encrypted-pkts"] = types.YLeaf{"ScEncryptedPkts", macsecTxStats.ScEncryptedPkts}
    macsecTxStats.EntityData.Leafs["sc-untagged-pkts"] = types.YLeaf{"ScUntaggedPkts", macsecTxStats.ScUntaggedPkts}
    macsecTxStats.EntityData.Leafs["sc-overrun-pkts"] = types.YLeaf{"ScOverrunPkts", macsecTxStats.ScOverrunPkts}
    macsecTxStats.EntityData.Leafs["sc-bypass-pkts"] = types.YLeaf{"ScBypassPkts", macsecTxStats.ScBypassPkts}
    macsecTxStats.EntityData.Leafs["sc-eapol-pkts"] = types.YLeaf{"ScEapolPkts", macsecTxStats.ScEapolPkts}
    macsecTxStats.EntityData.Leafs["sc-dropped-pkts"] = types.YLeaf{"ScDroppedPkts", macsecTxStats.ScDroppedPkts}
    macsecTxStats.EntityData.Leafs["current-an"] = types.YLeaf{"CurrentAn", macsecTxStats.CurrentAn}
    macsecTxStats.EntityData.Leafs["sa-encrypted-pkts"] = types.YLeaf{"SaEncryptedPkts", macsecTxStats.SaEncryptedPkts}
    return &(macsecTxStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats
// Rx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDecryptedOctets interface{}

    // Rx No Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoTagPkts interface{}

    // Rx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Rx Bad Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBadTagPkts interface{}

    // Rx Late Pkts. The type is interface{} with range: 0..18446744073709551615.
    ScLatePkts interface{}

    // Rx Delayed Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDelayedPkts interface{}

    // Rx Unchecked Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUncheckedPkts interface{}

    // Rx No SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoSciPkts interface{}

    // Rx Unknown SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnknownSciPkts interface{}

    // Rx Pkts Ok. The type is interface{} with range: 0..18446744073709551615.
    ScOkPkts interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotUsingPkts interface{}

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnusedPkts interface{}

    // Rx Not Valid Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotValidPkts interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    ScInvalidPkts interface{}

    // Rx Overrun Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Rx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Rx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Rx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Rx SA Level Stats. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat.
    RxSaStat []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetEntityData() *types.CommonEntityData {
    macsecRxStats.EntityData.YFilter = macsecRxStats.YFilter
    macsecRxStats.EntityData.YangName = "macsec-rx-stats"
    macsecRxStats.EntityData.BundleName = "cisco_ios_xr"
    macsecRxStats.EntityData.ParentYangName = "xlfpga-stats"
    macsecRxStats.EntityData.SegmentPath = "macsec-rx-stats"
    macsecRxStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecRxStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecRxStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecRxStats.EntityData.Children = make(map[string]types.YChild)
    macsecRxStats.EntityData.Children["rx-sa-stat"] = types.YChild{"RxSaStat", nil}
    for i := range macsecRxStats.RxSaStat {
        macsecRxStats.EntityData.Children[types.GetSegmentPath(&macsecRxStats.RxSaStat[i])] = types.YChild{"RxSaStat", &macsecRxStats.RxSaStat[i]}
    }
    macsecRxStats.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecRxStats.EntityData.Leafs["sc-decrypted-octets"] = types.YLeaf{"ScDecryptedOctets", macsecRxStats.ScDecryptedOctets}
    macsecRxStats.EntityData.Leafs["sc-no-tag-pkts"] = types.YLeaf{"ScNoTagPkts", macsecRxStats.ScNoTagPkts}
    macsecRxStats.EntityData.Leafs["sc-untagged-pkts"] = types.YLeaf{"ScUntaggedPkts", macsecRxStats.ScUntaggedPkts}
    macsecRxStats.EntityData.Leafs["sc-bad-tag-pkts"] = types.YLeaf{"ScBadTagPkts", macsecRxStats.ScBadTagPkts}
    macsecRxStats.EntityData.Leafs["sc-late-pkts"] = types.YLeaf{"ScLatePkts", macsecRxStats.ScLatePkts}
    macsecRxStats.EntityData.Leafs["sc-delayed-pkts"] = types.YLeaf{"ScDelayedPkts", macsecRxStats.ScDelayedPkts}
    macsecRxStats.EntityData.Leafs["sc-unchecked-pkts"] = types.YLeaf{"ScUncheckedPkts", macsecRxStats.ScUncheckedPkts}
    macsecRxStats.EntityData.Leafs["sc-no-sci-pkts"] = types.YLeaf{"ScNoSciPkts", macsecRxStats.ScNoSciPkts}
    macsecRxStats.EntityData.Leafs["sc-unknown-sci-pkts"] = types.YLeaf{"ScUnknownSciPkts", macsecRxStats.ScUnknownSciPkts}
    macsecRxStats.EntityData.Leafs["sc-ok-pkts"] = types.YLeaf{"ScOkPkts", macsecRxStats.ScOkPkts}
    macsecRxStats.EntityData.Leafs["sc-not-using-pkts"] = types.YLeaf{"ScNotUsingPkts", macsecRxStats.ScNotUsingPkts}
    macsecRxStats.EntityData.Leafs["sc-unused-pkts"] = types.YLeaf{"ScUnusedPkts", macsecRxStats.ScUnusedPkts}
    macsecRxStats.EntityData.Leafs["sc-not-valid-pkts"] = types.YLeaf{"ScNotValidPkts", macsecRxStats.ScNotValidPkts}
    macsecRxStats.EntityData.Leafs["sc-invalid-pkts"] = types.YLeaf{"ScInvalidPkts", macsecRxStats.ScInvalidPkts}
    macsecRxStats.EntityData.Leafs["sc-overrun-pkts"] = types.YLeaf{"ScOverrunPkts", macsecRxStats.ScOverrunPkts}
    macsecRxStats.EntityData.Leafs["sc-bypass-pkts"] = types.YLeaf{"ScBypassPkts", macsecRxStats.ScBypassPkts}
    macsecRxStats.EntityData.Leafs["sc-eapol-pkts"] = types.YLeaf{"ScEapolPkts", macsecRxStats.ScEapolPkts}
    macsecRxStats.EntityData.Leafs["sc-dropped-pkts"] = types.YLeaf{"ScDroppedPkts", macsecRxStats.ScDroppedPkts}
    return &(macsecRxStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
// Rx SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current Rx AN. The type is interface{} with range: 0..18446744073709551615.
    An interface{}

    // Rx Ok Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaOkPkts interface{}

    // Rx Pkts not using SA for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotUsingPkts interface{}

    // Rx Pkts Unused Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaUnusedPkts interface{}

    // Rx Not Valid Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotValidPkts interface{}

    // Rx Invalid Pkts for current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaInvalidPkts interface{}
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetEntityData() *types.CommonEntityData {
    rxSaStat.EntityData.YFilter = rxSaStat.YFilter
    rxSaStat.EntityData.YangName = "rx-sa-stat"
    rxSaStat.EntityData.BundleName = "cisco_ios_xr"
    rxSaStat.EntityData.ParentYangName = "macsec-rx-stats"
    rxSaStat.EntityData.SegmentPath = "rx-sa-stat"
    rxSaStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStat.EntityData.Children = make(map[string]types.YChild)
    rxSaStat.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStat.EntityData.Leafs["an"] = types.YLeaf{"An", rxSaStat.An}
    rxSaStat.EntityData.Leafs["sa-ok-pkts"] = types.YLeaf{"SaOkPkts", rxSaStat.SaOkPkts}
    rxSaStat.EntityData.Leafs["sa-not-using-pkts"] = types.YLeaf{"SaNotUsingPkts", rxSaStat.SaNotUsingPkts}
    rxSaStat.EntityData.Leafs["sa-unused-pkts"] = types.YLeaf{"SaUnusedPkts", rxSaStat.SaUnusedPkts}
    rxSaStat.EntityData.Leafs["sa-not-valid-pkts"] = types.YLeaf{"SaNotValidPkts", rxSaStat.SaNotValidPkts}
    rxSaStat.EntityData.Leafs["sa-invalid-pkts"] = types.YLeaf{"SaInvalidPkts", rxSaStat.SaInvalidPkts}
    return &(rxSaStat.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats
// ES200 Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats

    // Tx SC Macsec Stats.
    TxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats

    // Rx SC Macsec Stats.
    RxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats

    // Port level TX Stats.
    TxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats

    // Port level RX Stats.
    RxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats) GetEntityData() *types.CommonEntityData {
    es200Stats.EntityData.YFilter = es200Stats.YFilter
    es200Stats.EntityData.YangName = "es200-stats"
    es200Stats.EntityData.BundleName = "cisco_ios_xr"
    es200Stats.EntityData.ParentYangName = "ext"
    es200Stats.EntityData.SegmentPath = "es200-stats"
    es200Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es200Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es200Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es200Stats.EntityData.Children = make(map[string]types.YChild)
    es200Stats.EntityData.Children["tx-sa-stats"] = types.YChild{"TxSaStats", &es200Stats.TxSaStats}
    es200Stats.EntityData.Children["rx-sa-stats"] = types.YChild{"RxSaStats", &es200Stats.RxSaStats}
    es200Stats.EntityData.Children["tx-sc-macsec-stats"] = types.YChild{"TxScMacsecStats", &es200Stats.TxScMacsecStats}
    es200Stats.EntityData.Children["rx-sc-macsec-stats"] = types.YChild{"RxScMacsecStats", &es200Stats.RxScMacsecStats}
    es200Stats.EntityData.Children["tx-interface-macsec-stats"] = types.YChild{"TxInterfaceMacsecStats", &es200Stats.TxInterfaceMacsecStats}
    es200Stats.EntityData.Children["rx-interface-macsec-stats"] = types.YChild{"RxInterfaceMacsecStats", &es200Stats.RxInterfaceMacsecStats}
    es200Stats.EntityData.Children["tx-port-stats"] = types.YChild{"TxPortStats", &es200Stats.TxPortStats}
    es200Stats.EntityData.Children["rx-port-stats"] = types.YChild{"RxPortStats", &es200Stats.RxPortStats}
    es200Stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(es200Stats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // packets exceeding egress MTU. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected ?. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncryptedProtected1 interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxSaStats) GetEntityData() *types.CommonEntityData {
    txSaStats.EntityData.YFilter = txSaStats.YFilter
    txSaStats.EntityData.YangName = "tx-sa-stats"
    txSaStats.EntityData.BundleName = "cisco_ios_xr"
    txSaStats.EntityData.ParentYangName = "es200-stats"
    txSaStats.EntityData.SegmentPath = "tx-sa-stats"
    txSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSaStats.EntityData.Children = make(map[string]types.YChild)
    txSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txSaStats.EntityData.Leafs["out-pkts-too-long"] = types.YLeaf{"OutPktsTooLong", txSaStats.OutPktsTooLong}
    txSaStats.EntityData.Leafs["out-pkts-encrypted-protected"] = types.YLeaf{"OutPktsEncryptedProtected", txSaStats.OutPktsEncryptedProtected}
    txSaStats.EntityData.Leafs["out-octets-encrypted-protected1"] = types.YLeaf{"OutOctetsEncryptedProtected1", txSaStats.OutOctetsEncryptedProtected1}
    return &(txSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxSaStats) GetEntityData() *types.CommonEntityData {
    rxSaStats.EntityData.YFilter = rxSaStats.YFilter
    rxSaStats.EntityData.YangName = "rx-sa-stats"
    rxSaStats.EntityData.BundleName = "cisco_ios_xr"
    rxSaStats.EntityData.ParentYangName = "es200-stats"
    rxSaStats.EntityData.SegmentPath = "rx-sa-stats"
    rxSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStats.EntityData.Children = make(map[string]types.YChild)
    rxSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStats.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxSaStats.InPktsUnchecked}
    rxSaStats.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxSaStats.InPktsDelayed}
    rxSaStats.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxSaStats.InPktsLate}
    rxSaStats.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxSaStats.InPktsOk}
    rxSaStats.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxSaStats.InPktsInvalid}
    rxSaStats.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxSaStats.InPktsNotValid}
    rxSaStats.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxSaStats.InPktsNotUsingSa}
    rxSaStats.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxSaStats.InPktsUnusedSa}
    rxSaStats.EntityData.Leafs["in-octets-decrypted-validated1"] = types.YLeaf{"InOctetsDecryptedValidated1", rxSaStats.InOctetsDecryptedValidated1}
    rxSaStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxSaStats.InOctetsValidated}
    return &(rxSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats
// Tx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsSaNotInUse interface{}
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxScMacsecStats) GetEntityData() *types.CommonEntityData {
    txScMacsecStats.EntityData.YFilter = txScMacsecStats.YFilter
    txScMacsecStats.EntityData.YangName = "tx-sc-macsec-stats"
    txScMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txScMacsecStats.EntityData.ParentYangName = "es200-stats"
    txScMacsecStats.EntityData.SegmentPath = "tx-sc-macsec-stats"
    txScMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txScMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txScMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txScMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txScMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txScMacsecStats.EntityData.Leafs["out-pkts-sa-not-in-use"] = types.YLeaf{"OutPktsSaNotInUse", txScMacsecStats.OutPktsSaNotInUse}
    return &(txScMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats
// Rx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsSaNotInUse interface{}
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxScMacsecStats) GetEntityData() *types.CommonEntityData {
    rxScMacsecStats.EntityData.YFilter = rxScMacsecStats.YFilter
    rxScMacsecStats.EntityData.YangName = "rx-sc-macsec-stats"
    rxScMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxScMacsecStats.EntityData.ParentYangName = "es200-stats"
    rxScMacsecStats.EntityData.SegmentPath = "rx-sc-macsec-stats"
    rxScMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxScMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxScMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxScMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxScMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxScMacsecStats.EntityData.Leafs["in-pkts-sa-not-in-use"] = types.YLeaf{"InPktsSaNotInUse", rxScMacsecStats.InPktsSaNotInUse}
    return &(rxScMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // egress packet that is classified as control packet. The type is interface{}
    // with range: 0..18446744073709551615.
    OutPktCtrl interface{}

    // egress packet to go out untagged when protectFrames not set. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPktsUntagged interface{}

    // Octets tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsUnctrl interface{}

    // Octets tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCtrl interface{}

    // Octets tx on common port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCommon interface{}

    // Unicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsUnctrl interface{}

    // Unicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsCtrl interface{}

    // Multicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsUnctrl interface{}

    // Multicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsCtrl interface{}

    // Broadcast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsUnctrl interface{}

    // Broadcast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsClass interface{}

    // Packets dropped due to overflow in  processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsData interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    txInterfaceMacsecStats.EntityData.YFilter = txInterfaceMacsecStats.YFilter
    txInterfaceMacsecStats.EntityData.YangName = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txInterfaceMacsecStats.EntityData.ParentYangName = "es200-stats"
    txInterfaceMacsecStats.EntityData.SegmentPath = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txInterfaceMacsecStats.EntityData.Leafs["transform-error-pkts"] = types.YLeaf{"TransformErrorPkts", txInterfaceMacsecStats.TransformErrorPkts}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-ctrl"] = types.YLeaf{"OutPktCtrl", txInterfaceMacsecStats.OutPktCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkts-untagged"] = types.YLeaf{"OutPktsUntagged", txInterfaceMacsecStats.OutPktsUntagged}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-unctrl"] = types.YLeaf{"OutOctetsUnctrl", txInterfaceMacsecStats.OutOctetsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-ctrl"] = types.YLeaf{"OutOctetsCtrl", txInterfaceMacsecStats.OutOctetsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-common"] = types.YLeaf{"OutOctetsCommon", txInterfaceMacsecStats.OutOctetsCommon}
    txInterfaceMacsecStats.EntityData.Leafs["out-ucast-pkts-unctrl"] = types.YLeaf{"OutUcastPktsUnctrl", txInterfaceMacsecStats.OutUcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-ucast-pkts-ctrl"] = types.YLeaf{"OutUcastPktsCtrl", txInterfaceMacsecStats.OutUcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-mcast-pkts-unctrl"] = types.YLeaf{"OutMcastPktsUnctrl", txInterfaceMacsecStats.OutMcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-mcast-pkts-ctrl"] = types.YLeaf{"OutMcastPktsCtrl", txInterfaceMacsecStats.OutMcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-bcast-pkts-unctrl"] = types.YLeaf{"OutBcastPktsUnctrl", txInterfaceMacsecStats.OutBcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-bcast-pkts-ctrl"] = types.YLeaf{"OutBcastPktsCtrl", txInterfaceMacsecStats.OutBcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-drop-pkts-unctrl"] = types.YLeaf{"OutRxDropPktsUnctrl", txInterfaceMacsecStats.OutRxDropPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-drop-pkts-ctrl"] = types.YLeaf{"OutRxDropPktsCtrl", txInterfaceMacsecStats.OutRxDropPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-err-pkts-unctrl"] = types.YLeaf{"OutRxErrPktsUnctrl", txInterfaceMacsecStats.OutRxErrPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-err-pkts-ctrl"] = types.YLeaf{"OutRxErrPktsCtrl", txInterfaceMacsecStats.OutRxErrPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-drop-pkts-class"] = types.YLeaf{"OutDropPktsClass", txInterfaceMacsecStats.OutDropPktsClass}
    txInterfaceMacsecStats.EntityData.Leafs["out-drop-pkts-data"] = types.YLeaf{"OutDropPktsData", txInterfaceMacsecStats.OutDropPktsData}
    return &(txInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // ingress packet that is classified as control packet. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktCtrl interface{}

    // ingress packet untagged & validateFrames is strict. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktNoTag interface{}

    // ingress packet untagged & validateFrames is  !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUntagged interface{}

    // ingress frames received with an invalid MACSec tag or ICV                  
    // added with next one gives InPktsSCIMiss. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktBadTag interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is !strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnknownSci interface{}

    // ingress packets that are control or KaY packets. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktsTaggedCtrl interface{}

    // Octets rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsUnctrl interface{}

    // Octets rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsCtrl interface{}

    // Unicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsUnctrl interface{}

    // Unicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsCtrl interface{}

    // Multicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsUnctrl interface{}

    // Multicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsCtrl interface{}

    // Broadcast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsUnctrl interface{}

    // Broadcast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsClass interface{}

    // Packets dropped due to overflow in processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsData interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    rxInterfaceMacsecStats.EntityData.YFilter = rxInterfaceMacsecStats.YFilter
    rxInterfaceMacsecStats.EntityData.YangName = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxInterfaceMacsecStats.EntityData.ParentYangName = "es200-stats"
    rxInterfaceMacsecStats.EntityData.SegmentPath = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxInterfaceMacsecStats.EntityData.Leafs["transform-error-pkts"] = types.YLeaf{"TransformErrorPkts", rxInterfaceMacsecStats.TransformErrorPkts}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-ctrl"] = types.YLeaf{"InPktCtrl", rxInterfaceMacsecStats.InPktCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-tag"] = types.YLeaf{"InPktNoTag", rxInterfaceMacsecStats.InPktNoTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-untagged"] = types.YLeaf{"InPktsUntagged", rxInterfaceMacsecStats.InPktsUntagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-bad-tag"] = types.YLeaf{"InPktBadTag", rxInterfaceMacsecStats.InPktBadTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-sci"] = types.YLeaf{"InPktNoSci", rxInterfaceMacsecStats.InPktNoSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-unknown-sci"] = types.YLeaf{"InPktsUnknownSci", rxInterfaceMacsecStats.InPktsUnknownSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-tagged-ctrl"] = types.YLeaf{"InPktsTaggedCtrl", rxInterfaceMacsecStats.InPktsTaggedCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-octets-unctrl"] = types.YLeaf{"InOctetsUnctrl", rxInterfaceMacsecStats.InOctetsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-octets-ctrl"] = types.YLeaf{"InOctetsCtrl", rxInterfaceMacsecStats.InOctetsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-ucast-pkts-unctrl"] = types.YLeaf{"InUcastPktsUnctrl", rxInterfaceMacsecStats.InUcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-ucast-pkts-ctrl"] = types.YLeaf{"InUcastPktsCtrl", rxInterfaceMacsecStats.InUcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-mcast-pkts-unctrl"] = types.YLeaf{"InMcastPktsUnctrl", rxInterfaceMacsecStats.InMcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-mcast-pkts-ctrl"] = types.YLeaf{"InMcastPktsCtrl", rxInterfaceMacsecStats.InMcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-bcast-pkts-unctrl"] = types.YLeaf{"InBcastPktsUnctrl", rxInterfaceMacsecStats.InBcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-bcast-pkts-ctrl"] = types.YLeaf{"InBcastPktsCtrl", rxInterfaceMacsecStats.InBcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-drop-pkts-unctrl"] = types.YLeaf{"InRxDropPktsUnctrl", rxInterfaceMacsecStats.InRxDropPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-drop-pkts-ctrl"] = types.YLeaf{"InRxDropPktsCtrl", rxInterfaceMacsecStats.InRxDropPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-error-pkts-unctrl"] = types.YLeaf{"InRxErrorPktsUnctrl", rxInterfaceMacsecStats.InRxErrorPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-error-pkts-ctrl"] = types.YLeaf{"InRxErrorPktsCtrl", rxInterfaceMacsecStats.InRxErrorPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-drop-pkts-class"] = types.YLeaf{"InDropPktsClass", rxInterfaceMacsecStats.InDropPktsClass}
    rxInterfaceMacsecStats.EntityData.Leafs["in-drop-pkts-data"] = types.YLeaf{"InDropPktsData", rxInterfaceMacsecStats.InDropPktsData}
    return &(rxInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats
// Port level TX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_TxPortStats) GetEntityData() *types.CommonEntityData {
    txPortStats.EntityData.YFilter = txPortStats.YFilter
    txPortStats.EntityData.YangName = "tx-port-stats"
    txPortStats.EntityData.BundleName = "cisco_ios_xr"
    txPortStats.EntityData.ParentYangName = "es200-stats"
    txPortStats.EntityData.SegmentPath = "tx-port-stats"
    txPortStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txPortStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txPortStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txPortStats.EntityData.Children = make(map[string]types.YChild)
    txPortStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txPortStats.EntityData.Leafs["multi-flow-match"] = types.YLeaf{"MultiFlowMatch", txPortStats.MultiFlowMatch}
    txPortStats.EntityData.Leafs["parser-dropped"] = types.YLeaf{"ParserDropped", txPortStats.ParserDropped}
    txPortStats.EntityData.Leafs["flow-miss"] = types.YLeaf{"FlowMiss", txPortStats.FlowMiss}
    txPortStats.EntityData.Leafs["pkts-ctrl"] = types.YLeaf{"PktsCtrl", txPortStats.PktsCtrl}
    txPortStats.EntityData.Leafs["pkts-data"] = types.YLeaf{"PktsData", txPortStats.PktsData}
    txPortStats.EntityData.Leafs["pkts-dropped"] = types.YLeaf{"PktsDropped", txPortStats.PktsDropped}
    txPortStats.EntityData.Leafs["pkts-err-in"] = types.YLeaf{"PktsErrIn", txPortStats.PktsErrIn}
    return &(txPortStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats
// Port level RX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwStatistics_Ext_Es200Stats_RxPortStats) GetEntityData() *types.CommonEntityData {
    rxPortStats.EntityData.YFilter = rxPortStats.YFilter
    rxPortStats.EntityData.YangName = "rx-port-stats"
    rxPortStats.EntityData.BundleName = "cisco_ios_xr"
    rxPortStats.EntityData.ParentYangName = "es200-stats"
    rxPortStats.EntityData.SegmentPath = "rx-port-stats"
    rxPortStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxPortStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxPortStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxPortStats.EntityData.Children = make(map[string]types.YChild)
    rxPortStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxPortStats.EntityData.Leafs["multi-flow-match"] = types.YLeaf{"MultiFlowMatch", rxPortStats.MultiFlowMatch}
    rxPortStats.EntityData.Leafs["parser-dropped"] = types.YLeaf{"ParserDropped", rxPortStats.ParserDropped}
    rxPortStats.EntityData.Leafs["flow-miss"] = types.YLeaf{"FlowMiss", rxPortStats.FlowMiss}
    rxPortStats.EntityData.Leafs["pkts-ctrl"] = types.YLeaf{"PktsCtrl", rxPortStats.PktsCtrl}
    rxPortStats.EntityData.Leafs["pkts-data"] = types.YLeaf{"PktsData", rxPortStats.PktsData}
    rxPortStats.EntityData.Leafs["pkts-dropped"] = types.YLeaf{"PktsDropped", rxPortStats.PktsDropped}
    rxPortStats.EntityData.Leafs["pkts-err-in"] = types.YLeaf{"PktsErrIn", rxPortStats.PktsErrIn}
    return &(rxPortStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas
// Table of Hardware SAs
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hardware Security Association. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa.
    HwSa []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa
}

func (hwSas *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas) GetEntityData() *types.CommonEntityData {
    hwSas.EntityData.YFilter = hwSas.YFilter
    hwSas.EntityData.YangName = "hw-sas"
    hwSas.EntityData.BundleName = "cisco_ios_xr"
    hwSas.EntityData.ParentYangName = "interface"
    hwSas.EntityData.SegmentPath = "hw-sas"
    hwSas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwSas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwSas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwSas.EntityData.Children = make(map[string]types.YChild)
    hwSas.EntityData.Children["hw-sa"] = types.YChild{"HwSa", nil}
    for i := range hwSas.HwSa {
        hwSas.EntityData.Children[types.GetSegmentPath(&hwSas.HwSa[i])] = types.YChild{"HwSa", &hwSas.HwSa[i]}
    }
    hwSas.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwSas.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa
// Hardware Security Association
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SA ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SaId interface{}

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext
}

func (hwSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa) GetEntityData() *types.CommonEntityData {
    hwSa.EntityData.YFilter = hwSa.YFilter
    hwSa.EntityData.YangName = "hw-sa"
    hwSa.EntityData.BundleName = "cisco_ios_xr"
    hwSa.EntityData.ParentYangName = "hw-sas"
    hwSa.EntityData.SegmentPath = "hw-sa" + "[sa-id='" + fmt.Sprintf("%v", hwSa.SaId) + "']"
    hwSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwSa.EntityData.Children = make(map[string]types.YChild)
    hwSa.EntityData.Children["ext"] = types.YChild{"Ext", &hwSa.Ext}
    hwSa.EntityData.Leafs = make(map[string]types.YLeaf)
    hwSa.EntityData.Leafs["sa-id"] = types.YLeaf{"SaId", hwSa.SaId}
    return &(hwSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type_ interface{}

    // MSFPGA SA Information.
    MsfpgaSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa

    // XLFPGA SA Information.
    XlfpgaSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa

    // ES200 SA Information.
    Es200Sa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext) GetEntityData() *types.CommonEntityData {
    ext.EntityData.YFilter = ext.YFilter
    ext.EntityData.YangName = "ext"
    ext.EntityData.BundleName = "cisco_ios_xr"
    ext.EntityData.ParentYangName = "hw-sa"
    ext.EntityData.SegmentPath = "ext"
    ext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ext.EntityData.Children = make(map[string]types.YChild)
    ext.EntityData.Children["msfpga-sa"] = types.YChild{"MsfpgaSa", &ext.MsfpgaSa}
    ext.EntityData.Children["xlfpga-sa"] = types.YChild{"XlfpgaSa", &ext.XlfpgaSa}
    ext.EntityData.Children["es200-sa"] = types.YChild{"Es200Sa", &ext.Es200Sa}
    ext.EntityData.Leafs = make(map[string]types.YLeaf)
    ext.EntityData.Leafs["type"] = types.YLeaf{"Type_", ext.Type_}
    return &(ext.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa
// MSFPGA SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa

    // Rx SA Details.
    RxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa
}

func (msfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa) GetEntityData() *types.CommonEntityData {
    msfpgaSa.EntityData.YFilter = msfpgaSa.YFilter
    msfpgaSa.EntityData.YangName = "msfpga-sa"
    msfpgaSa.EntityData.BundleName = "cisco_ios_xr"
    msfpgaSa.EntityData.ParentYangName = "ext"
    msfpgaSa.EntityData.SegmentPath = "msfpga-sa"
    msfpgaSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    msfpgaSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    msfpgaSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    msfpgaSa.EntityData.Children = make(map[string]types.YChild)
    msfpgaSa.EntityData.Children["tx-sa"] = types.YChild{"TxSa", &msfpgaSa.TxSa}
    msfpgaSa.EntityData.Children["rx-sa"] = types.YChild{"RxSa", &msfpgaSa.RxSa}
    msfpgaSa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msfpgaSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SA Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // Crypto Algorithm. The type is interface{} with range: 0..255.
    CryptoAlgo interface{}

    // Key Length. The type is interface{} with range: 0..255.
    KeyLen interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // XPN EN. The type is interface{} with range: 0..255.
    Xpn interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Next Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    NextPn interface{}

    // Conf offset. The type is interface{} with range: 0..255.
    COffset interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Q bit. The type is bool.
    QBit interface{}

    // QQ bit. The type is bool.
    QqBit interface{}
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_TxSa) GetEntityData() *types.CommonEntityData {
    txSa.EntityData.YFilter = txSa.YFilter
    txSa.EntityData.YangName = "tx-sa"
    txSa.EntityData.BundleName = "cisco_ios_xr"
    txSa.EntityData.ParentYangName = "msfpga-sa"
    txSa.EntityData.SegmentPath = "tx-sa"
    txSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSa.EntityData.Children = make(map[string]types.YChild)
    txSa.EntityData.Leafs = make(map[string]types.YLeaf)
    txSa.EntityData.Leafs["sa-id"] = types.YLeaf{"SaId", txSa.SaId}
    txSa.EntityData.Leafs["valid"] = types.YLeaf{"Valid", txSa.Valid}
    txSa.EntityData.Leafs["is-egress"] = types.YLeaf{"IsEgress", txSa.IsEgress}
    txSa.EntityData.Leafs["crypto-algo"] = types.YLeaf{"CryptoAlgo", txSa.CryptoAlgo}
    txSa.EntityData.Leafs["key-len"] = types.YLeaf{"KeyLen", txSa.KeyLen}
    txSa.EntityData.Leafs["an"] = types.YLeaf{"An", txSa.An}
    txSa.EntityData.Leafs["xpn"] = types.YLeaf{"Xpn", txSa.Xpn}
    txSa.EntityData.Leafs["sci"] = types.YLeaf{"Sci", txSa.Sci}
    txSa.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", txSa.InUse}
    txSa.EntityData.Leafs["next-pn"] = types.YLeaf{"NextPn", txSa.NextPn}
    txSa.EntityData.Leafs["c-offset"] = types.YLeaf{"COffset", txSa.COffset}
    txSa.EntityData.Leafs["action"] = types.YLeaf{"Action", txSa.Action}
    txSa.EntityData.Leafs["q-bit"] = types.YLeaf{"QBit", txSa.QBit}
    txSa.EntityData.Leafs["qq-bit"] = types.YLeaf{"QqBit", txSa.QqBit}
    return &(txSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SA Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // Crypto Algorithm. The type is interface{} with range: 0..255.
    CryptoAlgo interface{}

    // Key Length. The type is interface{} with range: 0..255.
    KeyLen interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // XPN EN. The type is interface{} with range: 0..255.
    Xpn interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Next Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    NextPn interface{}

    // Conf offset. The type is interface{} with range: 0..255.
    COffset interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // Q bit. The type is bool.
    QBit interface{}

    // QQ bit. The type is bool.
    QqBit interface{}
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_MsfpgaSa_RxSa) GetEntityData() *types.CommonEntityData {
    rxSa.EntityData.YFilter = rxSa.YFilter
    rxSa.EntityData.YangName = "rx-sa"
    rxSa.EntityData.BundleName = "cisco_ios_xr"
    rxSa.EntityData.ParentYangName = "msfpga-sa"
    rxSa.EntityData.SegmentPath = "rx-sa"
    rxSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSa.EntityData.Children = make(map[string]types.YChild)
    rxSa.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSa.EntityData.Leafs["sa-id"] = types.YLeaf{"SaId", rxSa.SaId}
    rxSa.EntityData.Leafs["valid"] = types.YLeaf{"Valid", rxSa.Valid}
    rxSa.EntityData.Leafs["is-egress"] = types.YLeaf{"IsEgress", rxSa.IsEgress}
    rxSa.EntityData.Leafs["crypto-algo"] = types.YLeaf{"CryptoAlgo", rxSa.CryptoAlgo}
    rxSa.EntityData.Leafs["key-len"] = types.YLeaf{"KeyLen", rxSa.KeyLen}
    rxSa.EntityData.Leafs["an"] = types.YLeaf{"An", rxSa.An}
    rxSa.EntityData.Leafs["xpn"] = types.YLeaf{"Xpn", rxSa.Xpn}
    rxSa.EntityData.Leafs["sci"] = types.YLeaf{"Sci", rxSa.Sci}
    rxSa.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", rxSa.InUse}
    rxSa.EntityData.Leafs["next-pn"] = types.YLeaf{"NextPn", rxSa.NextPn}
    rxSa.EntityData.Leafs["c-offset"] = types.YLeaf{"COffset", rxSa.COffset}
    rxSa.EntityData.Leafs["action"] = types.YLeaf{"Action", rxSa.Action}
    rxSa.EntityData.Leafs["q-bit"] = types.YLeaf{"QBit", rxSa.QBit}
    rxSa.EntityData.Leafs["qq-bit"] = types.YLeaf{"QqBit", rxSa.QqBit}
    return &(rxSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa
// XLFPGA SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa

    // Rx SA Details.
    RxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa
}

func (xlfpgaSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa) GetEntityData() *types.CommonEntityData {
    xlfpgaSa.EntityData.YFilter = xlfpgaSa.YFilter
    xlfpgaSa.EntityData.YangName = "xlfpga-sa"
    xlfpgaSa.EntityData.BundleName = "cisco_ios_xr"
    xlfpgaSa.EntityData.ParentYangName = "ext"
    xlfpgaSa.EntityData.SegmentPath = "xlfpga-sa"
    xlfpgaSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xlfpgaSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xlfpgaSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xlfpgaSa.EntityData.Children = make(map[string]types.YChild)
    xlfpgaSa.EntityData.Children["tx-sa"] = types.YChild{"TxSa", &xlfpgaSa.TxSa}
    xlfpgaSa.EntityData.Children["rx-sa"] = types.YChild{"RxSa", &xlfpgaSa.RxSa}
    xlfpgaSa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(xlfpgaSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnable interface{}

    // Secure Mode - Must/Should. The type is interface{} with range: 0..255.
    SecureMode interface{}

    // Secure Channel ID. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Sec Tag Length(bytes) . The type is interface{} with range: 0..4294967295.
    // Units are byte.
    SectagLength interface{}

    // Cipher Suite Used. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..255.
    ConfidentialityOffset interface{}

    // FCS Error Config. The type is interface{} with range: 0..255.
    FcsErrCfg interface{}

    // Max Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNum interface{}

    // Association Number. The type is interface{} with range: 0..255.
    An interface{}

    // Initial Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    InitialPacketNumber interface{}

    // Short Secure Channel ID. The type is interface{} with range: 0..4294967295.
    Ssci interface{}

    // Current Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPacketNum interface{}

    // CRC Value. The type is interface{} with range: 0..4294967295.
    CrcValue interface{}
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_TxSa) GetEntityData() *types.CommonEntityData {
    txSa.EntityData.YFilter = txSa.YFilter
    txSa.EntityData.YangName = "tx-sa"
    txSa.EntityData.BundleName = "cisco_ios_xr"
    txSa.EntityData.ParentYangName = "xlfpga-sa"
    txSa.EntityData.SegmentPath = "tx-sa"
    txSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSa.EntityData.Children = make(map[string]types.YChild)
    txSa.EntityData.Leafs = make(map[string]types.YLeaf)
    txSa.EntityData.Leafs["protection-enable"] = types.YLeaf{"ProtectionEnable", txSa.ProtectionEnable}
    txSa.EntityData.Leafs["secure-mode"] = types.YLeaf{"SecureMode", txSa.SecureMode}
    txSa.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", txSa.SecureChannelId}
    txSa.EntityData.Leafs["sectag-length"] = types.YLeaf{"SectagLength", txSa.SectagLength}
    txSa.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", txSa.CipherSuite}
    txSa.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", txSa.ConfidentialityOffset}
    txSa.EntityData.Leafs["fcs-err-cfg"] = types.YLeaf{"FcsErrCfg", txSa.FcsErrCfg}
    txSa.EntityData.Leafs["max-packet-num"] = types.YLeaf{"MaxPacketNum", txSa.MaxPacketNum}
    txSa.EntityData.Leafs["an"] = types.YLeaf{"An", txSa.An}
    txSa.EntityData.Leafs["initial-packet-number"] = types.YLeaf{"InitialPacketNumber", txSa.InitialPacketNumber}
    txSa.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", txSa.Ssci}
    txSa.EntityData.Leafs["current-packet-num"] = types.YLeaf{"CurrentPacketNum", txSa.CurrentPacketNum}
    txSa.EntityData.Leafs["crc-value"] = types.YLeaf{"CrcValue", txSa.CrcValue}
    return &(txSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnable interface{}

    // Secure Mode - Must/Should. The type is interface{} with range:
    // 0..4294967295.
    SecureMode interface{}

    // Replay Protect Mode. The type is bool.
    ReplayProtectMode interface{}

    // Validation Mode. The type is interface{} with range: 0..4294967295.
    ValidationMode interface{}

    // Replay Window . The type is interface{} with range: 0..4294967295.
    ReplayWindow interface{}

    // Secure Channel ID. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Cipher Suite Used. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..255.
    ConfidentialityOffset interface{}

    // FCS Error Config. The type is interface{} with range: 0..4294967295.
    FcsErrCfg interface{}

    // Auth  Error Config. The type is interface{} with range: 0..4294967295.
    AuthErrCfg interface{}

    // Max Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNum interface{}

    // Num of AN's in Use. The type is interface{} with range: 0..4294967295.
    NumAnInUse interface{}

    // Association Number. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    An interface{}

    // Recent Association Num. The type is interface{} with range: 0..255.
    RecentAn interface{}

    // Untagged Pkts Detected. The type is bool.
    PktUntaggedDetected interface{}

    // Tagged Pkts Detected. The type is bool.
    PktTaggedDetected interface{}

    // Tagged Pkts Validated. The type is bool.
    PktTaggedValidated interface{}

    // Current Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPacketNum interface{}

    // Short Secure Channel ID. The type is slice of interface{} with range:
    // 0..4294967295.
    Ssci []interface{}

    // Lowest Acceptable Packet Number. The type is slice of interface{} with
    // range: 0..18446744073709551615.
    LowestAcceptablePacketNum []interface{}

    // Next expected Packet Number. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    NextExpectedPacketNum []interface{}

    // CRC Value. The type is slice of interface{} with range: 0..4294967295.
    CrcValue []interface{}
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_XlfpgaSa_RxSa) GetEntityData() *types.CommonEntityData {
    rxSa.EntityData.YFilter = rxSa.YFilter
    rxSa.EntityData.YangName = "rx-sa"
    rxSa.EntityData.BundleName = "cisco_ios_xr"
    rxSa.EntityData.ParentYangName = "xlfpga-sa"
    rxSa.EntityData.SegmentPath = "rx-sa"
    rxSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSa.EntityData.Children = make(map[string]types.YChild)
    rxSa.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSa.EntityData.Leafs["protection-enable"] = types.YLeaf{"ProtectionEnable", rxSa.ProtectionEnable}
    rxSa.EntityData.Leafs["secure-mode"] = types.YLeaf{"SecureMode", rxSa.SecureMode}
    rxSa.EntityData.Leafs["replay-protect-mode"] = types.YLeaf{"ReplayProtectMode", rxSa.ReplayProtectMode}
    rxSa.EntityData.Leafs["validation-mode"] = types.YLeaf{"ValidationMode", rxSa.ValidationMode}
    rxSa.EntityData.Leafs["replay-window"] = types.YLeaf{"ReplayWindow", rxSa.ReplayWindow}
    rxSa.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", rxSa.SecureChannelId}
    rxSa.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", rxSa.CipherSuite}
    rxSa.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", rxSa.ConfidentialityOffset}
    rxSa.EntityData.Leafs["fcs-err-cfg"] = types.YLeaf{"FcsErrCfg", rxSa.FcsErrCfg}
    rxSa.EntityData.Leafs["auth-err-cfg"] = types.YLeaf{"AuthErrCfg", rxSa.AuthErrCfg}
    rxSa.EntityData.Leafs["max-packet-num"] = types.YLeaf{"MaxPacketNum", rxSa.MaxPacketNum}
    rxSa.EntityData.Leafs["num-an-in-use"] = types.YLeaf{"NumAnInUse", rxSa.NumAnInUse}
    rxSa.EntityData.Leafs["an"] = types.YLeaf{"An", rxSa.An}
    rxSa.EntityData.Leafs["recent-an"] = types.YLeaf{"RecentAn", rxSa.RecentAn}
    rxSa.EntityData.Leafs["pkt-untagged-detected"] = types.YLeaf{"PktUntaggedDetected", rxSa.PktUntaggedDetected}
    rxSa.EntityData.Leafs["pkt-tagged-detected"] = types.YLeaf{"PktTaggedDetected", rxSa.PktTaggedDetected}
    rxSa.EntityData.Leafs["pkt-tagged-validated"] = types.YLeaf{"PktTaggedValidated", rxSa.PktTaggedValidated}
    rxSa.EntityData.Leafs["current-packet-num"] = types.YLeaf{"CurrentPacketNum", rxSa.CurrentPacketNum}
    rxSa.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", rxSa.Ssci}
    rxSa.EntityData.Leafs["lowest-acceptable-packet-num"] = types.YLeaf{"LowestAcceptablePacketNum", rxSa.LowestAcceptablePacketNum}
    rxSa.EntityData.Leafs["next-expected-packet-num"] = types.YLeaf{"NextExpectedPacketNum", rxSa.NextExpectedPacketNum}
    rxSa.EntityData.Leafs["crc-value"] = types.YLeaf{"CrcValue", rxSa.CrcValue}
    return &(rxSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa
// ES200 SA Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Details.
    TxSa MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa

    // Rx SA Details. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa.
    RxSa []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa
}

func (es200Sa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa) GetEntityData() *types.CommonEntityData {
    es200Sa.EntityData.YFilter = es200Sa.YFilter
    es200Sa.EntityData.YangName = "es200-sa"
    es200Sa.EntityData.BundleName = "cisco_ios_xr"
    es200Sa.EntityData.ParentYangName = "ext"
    es200Sa.EntityData.SegmentPath = "es200-sa"
    es200Sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es200Sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es200Sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es200Sa.EntityData.Children = make(map[string]types.YChild)
    es200Sa.EntityData.Children["tx-sa"] = types.YChild{"TxSa", &es200Sa.TxSa}
    es200Sa.EntityData.Children["rx-sa"] = types.YChild{"RxSa", nil}
    for i := range es200Sa.RxSa {
        es200Sa.EntityData.Children[types.GetSegmentPath(&es200Sa.RxSa[i])] = types.YChild{"RxSa", &es200Sa.RxSa[i]}
    }
    es200Sa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(es200Sa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa
// Tx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is structure valid. The type is bool.
    IsValid interface{}

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SC Number. The type is interface{} with range: 0..4294967295.
    ScNo interface{}

    // packets exceeding egress MTU. The type is interface{} with range: 0..255.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range: 0..255.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected. The type is interface{} with range: 0..255.
    OutOctetsEncryptedProtected1 interface{}

    // Initial Packet Number. The type is interface{} with range: 0..255.
    InitialPktNumber interface{}

    // Current packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    CurrentPktNumber interface{}

    // Maximum packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPktNumber interface{}

    // Xform Params.
    XformParams MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams
}

func (txSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa) GetEntityData() *types.CommonEntityData {
    txSa.EntityData.YFilter = txSa.YFilter
    txSa.EntityData.YangName = "tx-sa"
    txSa.EntityData.BundleName = "cisco_ios_xr"
    txSa.EntityData.ParentYangName = "es200-sa"
    txSa.EntityData.SegmentPath = "tx-sa"
    txSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSa.EntityData.Children = make(map[string]types.YChild)
    txSa.EntityData.Children["xform-params"] = types.YChild{"XformParams", &txSa.XformParams}
    txSa.EntityData.Leafs = make(map[string]types.YLeaf)
    txSa.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", txSa.IsValid}
    txSa.EntityData.Leafs["sa-id"] = types.YLeaf{"SaId", txSa.SaId}
    txSa.EntityData.Leafs["sc-no"] = types.YLeaf{"ScNo", txSa.ScNo}
    txSa.EntityData.Leafs["out-pkts-too-long"] = types.YLeaf{"OutPktsTooLong", txSa.OutPktsTooLong}
    txSa.EntityData.Leafs["out-pkts-encrypted-protected"] = types.YLeaf{"OutPktsEncryptedProtected", txSa.OutPktsEncryptedProtected}
    txSa.EntityData.Leafs["out-octets-encrypted-protected1"] = types.YLeaf{"OutOctetsEncryptedProtected1", txSa.OutOctetsEncryptedProtected1}
    txSa.EntityData.Leafs["initial-pkt-number"] = types.YLeaf{"InitialPktNumber", txSa.InitialPktNumber}
    txSa.EntityData.Leafs["current-pkt-number"] = types.YLeaf{"CurrentPktNumber", txSa.CurrentPktNumber}
    txSa.EntityData.Leafs["max-pkt-number"] = types.YLeaf{"MaxPktNumber", txSa.MaxPktNumber}
    return &(txSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams
//  Xform Params
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // range of pkt nos considered valid. The type is interface{} with range:
    // 0..4294967295.
    ReplayWinSize interface{}

    // Cryptographic algo used. The type is string.
    CryptAlgo interface{}

    // APM_TRUE if this is Egress Transform record, APM_FALSE otherwise. The type
    // is bool.
    IsEgressTr interface{}

    // AES Key length. The type is string.
    AesKeyLen interface{}

    // Association Number for egress. The type is interface{} with range: 0..255.
    AssocNum interface{}

    // TRUE if Seq Num is 64-bit, FALSE if it is 32-bit. The type is bool.
    IsSeqNum64Bit interface{}

    // TRUE to generate the authKey, so authKey in this struct not used           
    // APM_FALSE to use provided authKey. The type is bool.
    BgenAuthKey interface{}
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_TxSa_XformParams) GetEntityData() *types.CommonEntityData {
    xformParams.EntityData.YFilter = xformParams.YFilter
    xformParams.EntityData.YangName = "xform-params"
    xformParams.EntityData.BundleName = "cisco_ios_xr"
    xformParams.EntityData.ParentYangName = "tx-sa"
    xformParams.EntityData.SegmentPath = "xform-params"
    xformParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xformParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xformParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xformParams.EntityData.Children = make(map[string]types.YChild)
    xformParams.EntityData.Leafs = make(map[string]types.YLeaf)
    xformParams.EntityData.Leafs["replay-win-size"] = types.YLeaf{"ReplayWinSize", xformParams.ReplayWinSize}
    xformParams.EntityData.Leafs["crypt-algo"] = types.YLeaf{"CryptAlgo", xformParams.CryptAlgo}
    xformParams.EntityData.Leafs["is-egress-tr"] = types.YLeaf{"IsEgressTr", xformParams.IsEgressTr}
    xformParams.EntityData.Leafs["aes-key-len"] = types.YLeaf{"AesKeyLen", xformParams.AesKeyLen}
    xformParams.EntityData.Leafs["assoc-num"] = types.YLeaf{"AssocNum", xformParams.AssocNum}
    xformParams.EntityData.Leafs["is-seq-num64-bit"] = types.YLeaf{"IsSeqNum64Bit", xformParams.IsSeqNum64Bit}
    xformParams.EntityData.Leafs["bgen-auth-key"] = types.YLeaf{"BgenAuthKey", xformParams.BgenAuthKey}
    return &(xformParams.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa
// Rx SA Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is structure valid. The type is bool.
    IsValid interface{}

    // SA Index. The type is interface{} with range: 0..255.
    SaId interface{}

    // SC Number. The type is interface{} with range: 0..4294967295.
    ScNo interface{}

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..255.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..255.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..255.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range: 0..255.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..255.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..255.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..255.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use& validateFrames !strict. The type is
    // interface{} with range: 0..255.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range: 0..255.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range: 0..255.
    InOctetsValidated interface{}

    // Xform Params.
    XformParams MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams
}

func (rxSa *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa) GetEntityData() *types.CommonEntityData {
    rxSa.EntityData.YFilter = rxSa.YFilter
    rxSa.EntityData.YangName = "rx-sa"
    rxSa.EntityData.BundleName = "cisco_ios_xr"
    rxSa.EntityData.ParentYangName = "es200-sa"
    rxSa.EntityData.SegmentPath = "rx-sa"
    rxSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSa.EntityData.Children = make(map[string]types.YChild)
    rxSa.EntityData.Children["xform-params"] = types.YChild{"XformParams", &rxSa.XformParams}
    rxSa.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSa.EntityData.Leafs["is-valid"] = types.YLeaf{"IsValid", rxSa.IsValid}
    rxSa.EntityData.Leafs["sa-id"] = types.YLeaf{"SaId", rxSa.SaId}
    rxSa.EntityData.Leafs["sc-no"] = types.YLeaf{"ScNo", rxSa.ScNo}
    rxSa.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxSa.InPktsUnchecked}
    rxSa.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxSa.InPktsDelayed}
    rxSa.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxSa.InPktsLate}
    rxSa.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxSa.InPktsOk}
    rxSa.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxSa.InPktsInvalid}
    rxSa.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxSa.InPktsNotValid}
    rxSa.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxSa.InPktsNotUsingSa}
    rxSa.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxSa.InPktsUnusedSa}
    rxSa.EntityData.Leafs["in-octets-decrypted-validated1"] = types.YLeaf{"InOctetsDecryptedValidated1", rxSa.InOctetsDecryptedValidated1}
    rxSa.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxSa.InOctetsValidated}
    return &(rxSa.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams
//  Xform Params
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // range of pkt nos considered valid. The type is interface{} with range:
    // 0..4294967295.
    ReplayWinSize interface{}

    // Cryptographic algo used. The type is string.
    CryptAlgo interface{}

    // APM_TRUE if this is Egress Transform record, APM_FALSE otherwise. The type
    // is bool.
    IsEgressTr interface{}

    // AES Key length. The type is string.
    AesKeyLen interface{}

    // Association Number for egress. The type is interface{} with range: 0..255.
    AssocNum interface{}

    // TRUE if Seq Num is 64-bit, FALSE if it is 32-bit. The type is bool.
    IsSeqNum64Bit interface{}

    // TRUE to generate the authKey, so authKey in this struct not used           
    // APM_FALSE to use provided authKey. The type is bool.
    BgenAuthKey interface{}
}

func (xformParams *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwSas_HwSa_Ext_Es200Sa_RxSa_XformParams) GetEntityData() *types.CommonEntityData {
    xformParams.EntityData.YFilter = xformParams.YFilter
    xformParams.EntityData.YangName = "xform-params"
    xformParams.EntityData.BundleName = "cisco_ios_xr"
    xformParams.EntityData.ParentYangName = "rx-sa"
    xformParams.EntityData.SegmentPath = "xform-params"
    xformParams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xformParams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xformParams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xformParams.EntityData.Children = make(map[string]types.YChild)
    xformParams.EntityData.Leafs = make(map[string]types.YLeaf)
    xformParams.EntityData.Leafs["replay-win-size"] = types.YLeaf{"ReplayWinSize", xformParams.ReplayWinSize}
    xformParams.EntityData.Leafs["crypt-algo"] = types.YLeaf{"CryptAlgo", xformParams.CryptAlgo}
    xformParams.EntityData.Leafs["is-egress-tr"] = types.YLeaf{"IsEgressTr", xformParams.IsEgressTr}
    xformParams.EntityData.Leafs["aes-key-len"] = types.YLeaf{"AesKeyLen", xformParams.AesKeyLen}
    xformParams.EntityData.Leafs["assoc-num"] = types.YLeaf{"AssocNum", xformParams.AssocNum}
    xformParams.EntityData.Leafs["is-seq-num64-bit"] = types.YLeaf{"IsSeqNum64Bit", xformParams.IsSeqNum64Bit}
    xformParams.EntityData.Leafs["bgen-auth-key"] = types.YLeaf{"BgenAuthKey", xformParams.BgenAuthKey}
    return &(xformParams.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS
// Table of Hardware Flows
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hardware Flow. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow.
    HwFlow []MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow
}

func (hwFlowS *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS) GetEntityData() *types.CommonEntityData {
    hwFlowS.EntityData.YFilter = hwFlowS.YFilter
    hwFlowS.EntityData.YangName = "hw-flow-s"
    hwFlowS.EntityData.BundleName = "cisco_ios_xr"
    hwFlowS.EntityData.ParentYangName = "interface"
    hwFlowS.EntityData.SegmentPath = "hw-flow-s"
    hwFlowS.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwFlowS.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwFlowS.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwFlowS.EntityData.Children = make(map[string]types.YChild)
    hwFlowS.EntityData.Children["hw-flow"] = types.YChild{"HwFlow", nil}
    for i := range hwFlowS.HwFlow {
        hwFlowS.EntityData.Children[types.GetSegmentPath(&hwFlowS.HwFlow[i])] = types.YChild{"HwFlow", &hwFlowS.HwFlow[i]}
    }
    hwFlowS.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hwFlowS.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow
// Hardware Flow
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. FLOW ID. The type is interface{} with range:
    // -2147483648..2147483647.
    FlowId interface{}

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext
}

func (hwFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow) GetEntityData() *types.CommonEntityData {
    hwFlow.EntityData.YFilter = hwFlow.YFilter
    hwFlow.EntityData.YangName = "hw-flow"
    hwFlow.EntityData.BundleName = "cisco_ios_xr"
    hwFlow.EntityData.ParentYangName = "hw-flow-s"
    hwFlow.EntityData.SegmentPath = "hw-flow" + "[flow-id='" + fmt.Sprintf("%v", hwFlow.FlowId) + "']"
    hwFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwFlow.EntityData.Children = make(map[string]types.YChild)
    hwFlow.EntityData.Children["ext"] = types.YChild{"Ext", &hwFlow.Ext}
    hwFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    hwFlow.EntityData.Leafs["flow-id"] = types.YLeaf{"FlowId", hwFlow.FlowId}
    return &(hwFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type_ interface{}

    // MSFPGA Flow Information.
    MsfpgaFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow

    // ES200 Flow Information.
    Es200Flow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext) GetEntityData() *types.CommonEntityData {
    ext.EntityData.YFilter = ext.YFilter
    ext.EntityData.YangName = "ext"
    ext.EntityData.BundleName = "cisco_ios_xr"
    ext.EntityData.ParentYangName = "hw-flow"
    ext.EntityData.SegmentPath = "ext"
    ext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ext.EntityData.Children = make(map[string]types.YChild)
    ext.EntityData.Children["msfpga-flow"] = types.YChild{"MsfpgaFlow", &ext.MsfpgaFlow}
    ext.EntityData.Children["es200-flow"] = types.YChild{"Es200Flow", &ext.Es200Flow}
    ext.EntityData.Leafs = make(map[string]types.YLeaf)
    ext.EntityData.Leafs["type"] = types.YLeaf{"Type_", ext.Type_}
    return &(ext.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow
// MSFPGA Flow Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Flow Details.
    TxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow

    // Rx Flow Details.
    RxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow
}

func (msfpgaFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow) GetEntityData() *types.CommonEntityData {
    msfpgaFlow.EntityData.YFilter = msfpgaFlow.YFilter
    msfpgaFlow.EntityData.YangName = "msfpga-flow"
    msfpgaFlow.EntityData.BundleName = "cisco_ios_xr"
    msfpgaFlow.EntityData.ParentYangName = "ext"
    msfpgaFlow.EntityData.SegmentPath = "msfpga-flow"
    msfpgaFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    msfpgaFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    msfpgaFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    msfpgaFlow.EntityData.Children = make(map[string]types.YChild)
    msfpgaFlow.EntityData.Children["tx-flow"] = types.YChild{"TxFlow", &msfpgaFlow.TxFlow}
    msfpgaFlow.EntityData.Children["rx-flow"] = types.YChild{"RxFlow", &msfpgaFlow.RxFlow}
    msfpgaFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msfpgaFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow
// Tx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Index. The type is interface{} with range: 0..255.
    FlowId interface{}

    // Flow Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // If MAC SA in Use. The type is bool.
    SmacInuse interface{}

    // If MAC DA in Use. The type is bool.
    DmacInuse interface{}

    // Ether Type. The type is interface{} with range: 0..65535.
    Ethertype interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..65535.
    OuterVlan interface{}

    // Outer Vlan UserPri. The type is interface{} with range: 0..255.
    OuterVlanUp interface{}

    // Outer Vlan TPID. The type is interface{} with range: 0..65535.
    OuterVlanTpid interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..65535.
    InnerVlan interface{}

    // Inner Vlan UserPri. The type is interface{} with range: 0..255.
    InnerVlanUp interface{}

    // Inner Vlan TPID. The type is interface{} with range: 0..65535.
    InnerVlanTpid interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Source Port ChkEn. The type is bool.
    SourcePortChk interface{}

    // If SCI in use. The type is bool.
    SciInuse interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // Match Priority. The type is interface{} with range: 0..255.
    MatchPri interface{}

    // Is Control Pkt. The type is bool.
    IsCtrlPkt interface{}

    // Ctrl Pkt ChkEn. The type is bool.
    CtrlCheck interface{}

    // MatchUntagged. The type is bool.
    MatchUntagged interface{}

    // MatchTagged. The type is bool.
    MatchTagged interface{}

    // Match Bad Tag. The type is bool.
    MatchBadTag interface{}

    // MatchKaYTag. The type is bool.
    MatchKayTag interface{}

    // TCI V. The type is interface{} with range: 0..255.
    TciV interface{}

    // TCI ES. The type is interface{} with range: 0..255.
    TciEXr interface{}

    // TCI SC. The type is interface{} with range: 0..255.
    TciSc interface{}

    // TCI SCB. The type is interface{} with range: 0..255.
    TciScb interface{}

    // TCI E. The type is interface{} with range: 0..255.
    Tci interface{}

    // TCI C. The type is interface{} with range: 0..255.
    TciC interface{}

    // TCI AN. The type is interface{} with range: 0..255.
    TciAn interface{}

    // TciAnChkEn. The type is bool.
    TciAnChk interface{}

    // TciChkEn. The type is bool.
    TciChk interface{}

    // SAI. The type is interface{} with range: 0..4294967295.
    Sai interface{}

    // MAC SA. The type is slice of interface{} with range: 0..255.
    Macsa []interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_TxFlow) GetEntityData() *types.CommonEntityData {
    txFlow.EntityData.YFilter = txFlow.YFilter
    txFlow.EntityData.YangName = "tx-flow"
    txFlow.EntityData.BundleName = "cisco_ios_xr"
    txFlow.EntityData.ParentYangName = "msfpga-flow"
    txFlow.EntityData.SegmentPath = "tx-flow"
    txFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txFlow.EntityData.Children = make(map[string]types.YChild)
    txFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    txFlow.EntityData.Leafs["flow-id"] = types.YLeaf{"FlowId", txFlow.FlowId}
    txFlow.EntityData.Leafs["valid"] = types.YLeaf{"Valid", txFlow.Valid}
    txFlow.EntityData.Leafs["is-egress"] = types.YLeaf{"IsEgress", txFlow.IsEgress}
    txFlow.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", txFlow.InUse}
    txFlow.EntityData.Leafs["action"] = types.YLeaf{"Action", txFlow.Action}
    txFlow.EntityData.Leafs["smac-inuse"] = types.YLeaf{"SmacInuse", txFlow.SmacInuse}
    txFlow.EntityData.Leafs["dmac-inuse"] = types.YLeaf{"DmacInuse", txFlow.DmacInuse}
    txFlow.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", txFlow.Ethertype}
    txFlow.EntityData.Leafs["outer-vlan"] = types.YLeaf{"OuterVlan", txFlow.OuterVlan}
    txFlow.EntityData.Leafs["outer-vlan-up"] = types.YLeaf{"OuterVlanUp", txFlow.OuterVlanUp}
    txFlow.EntityData.Leafs["outer-vlan-tpid"] = types.YLeaf{"OuterVlanTpid", txFlow.OuterVlanTpid}
    txFlow.EntityData.Leafs["inner-vlan"] = types.YLeaf{"InnerVlan", txFlow.InnerVlan}
    txFlow.EntityData.Leafs["inner-vlan-up"] = types.YLeaf{"InnerVlanUp", txFlow.InnerVlanUp}
    txFlow.EntityData.Leafs["inner-vlan-tpid"] = types.YLeaf{"InnerVlanTpid", txFlow.InnerVlanTpid}
    txFlow.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", txFlow.SourcePort}
    txFlow.EntityData.Leafs["source-port-chk"] = types.YLeaf{"SourcePortChk", txFlow.SourcePortChk}
    txFlow.EntityData.Leafs["sci-inuse"] = types.YLeaf{"SciInuse", txFlow.SciInuse}
    txFlow.EntityData.Leafs["sci"] = types.YLeaf{"Sci", txFlow.Sci}
    txFlow.EntityData.Leafs["match-pri"] = types.YLeaf{"MatchPri", txFlow.MatchPri}
    txFlow.EntityData.Leafs["is-ctrl-pkt"] = types.YLeaf{"IsCtrlPkt", txFlow.IsCtrlPkt}
    txFlow.EntityData.Leafs["ctrl-check"] = types.YLeaf{"CtrlCheck", txFlow.CtrlCheck}
    txFlow.EntityData.Leafs["match-untagged"] = types.YLeaf{"MatchUntagged", txFlow.MatchUntagged}
    txFlow.EntityData.Leafs["match-tagged"] = types.YLeaf{"MatchTagged", txFlow.MatchTagged}
    txFlow.EntityData.Leafs["match-bad-tag"] = types.YLeaf{"MatchBadTag", txFlow.MatchBadTag}
    txFlow.EntityData.Leafs["match-kay-tag"] = types.YLeaf{"MatchKayTag", txFlow.MatchKayTag}
    txFlow.EntityData.Leafs["tci-v"] = types.YLeaf{"TciV", txFlow.TciV}
    txFlow.EntityData.Leafs["tci-e-xr"] = types.YLeaf{"TciEXr", txFlow.TciEXr}
    txFlow.EntityData.Leafs["tci-sc"] = types.YLeaf{"TciSc", txFlow.TciSc}
    txFlow.EntityData.Leafs["tci-scb"] = types.YLeaf{"TciScb", txFlow.TciScb}
    txFlow.EntityData.Leafs["tci"] = types.YLeaf{"Tci", txFlow.Tci}
    txFlow.EntityData.Leafs["tci-c"] = types.YLeaf{"TciC", txFlow.TciC}
    txFlow.EntityData.Leafs["tci-an"] = types.YLeaf{"TciAn", txFlow.TciAn}
    txFlow.EntityData.Leafs["tci-an-chk"] = types.YLeaf{"TciAnChk", txFlow.TciAnChk}
    txFlow.EntityData.Leafs["tci-chk"] = types.YLeaf{"TciChk", txFlow.TciChk}
    txFlow.EntityData.Leafs["sai"] = types.YLeaf{"Sai", txFlow.Sai}
    txFlow.EntityData.Leafs["macsa"] = types.YLeaf{"Macsa", txFlow.Macsa}
    txFlow.EntityData.Leafs["macda"] = types.YLeaf{"Macda", txFlow.Macda}
    return &(txFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow
// Rx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Index. The type is interface{} with range: 0..255.
    FlowId interface{}

    // Flow Validity. The type is bool.
    Valid interface{}

    // rx_tx direction. The type is bool.
    IsEgress interface{}

    // In Use. The type is bool.
    InUse interface{}

    // Action. The type is interface{} with range: 0..255.
    Action interface{}

    // If MAC SA in Use. The type is bool.
    SmacInuse interface{}

    // If MAC DA in Use. The type is bool.
    DmacInuse interface{}

    // Ether Type. The type is interface{} with range: 0..65535.
    Ethertype interface{}

    // Outer VLAN ID. The type is interface{} with range: 0..65535.
    OuterVlan interface{}

    // Outer Vlan UserPri. The type is interface{} with range: 0..255.
    OuterVlanUp interface{}

    // Outer Vlan TPID. The type is interface{} with range: 0..65535.
    OuterVlanTpid interface{}

    // Inner VLAN ID. The type is interface{} with range: 0..65535.
    InnerVlan interface{}

    // Inner Vlan UserPri. The type is interface{} with range: 0..255.
    InnerVlanUp interface{}

    // Inner Vlan TPID. The type is interface{} with range: 0..65535.
    InnerVlanTpid interface{}

    // Source Port. The type is interface{} with range: 0..4294967295.
    SourcePort interface{}

    // Source Port ChkEn. The type is bool.
    SourcePortChk interface{}

    // If SCI in use. The type is bool.
    SciInuse interface{}

    // SCI. The type is interface{} with range: 0..18446744073709551615.
    Sci interface{}

    // Match Priority. The type is interface{} with range: 0..255.
    MatchPri interface{}

    // Is Control Pkt. The type is bool.
    IsCtrlPkt interface{}

    // Ctrl Pkt ChkEn. The type is bool.
    CtrlCheck interface{}

    // MatchUntagged. The type is bool.
    MatchUntagged interface{}

    // MatchTagged. The type is bool.
    MatchTagged interface{}

    // Match Bad Tag. The type is bool.
    MatchBadTag interface{}

    // MatchKaYTag. The type is bool.
    MatchKayTag interface{}

    // TCI V. The type is interface{} with range: 0..255.
    TciV interface{}

    // TCI ES. The type is interface{} with range: 0..255.
    TciEXr interface{}

    // TCI SC. The type is interface{} with range: 0..255.
    TciSc interface{}

    // TCI SCB. The type is interface{} with range: 0..255.
    TciScb interface{}

    // TCI E. The type is interface{} with range: 0..255.
    Tci interface{}

    // TCI C. The type is interface{} with range: 0..255.
    TciC interface{}

    // TCI AN. The type is interface{} with range: 0..255.
    TciAn interface{}

    // TciAnChkEn. The type is bool.
    TciAnChk interface{}

    // TciChkEn. The type is bool.
    TciChk interface{}

    // SAI. The type is interface{} with range: 0..4294967295.
    Sai interface{}

    // MAC SA. The type is slice of interface{} with range: 0..255.
    Macsa []interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_MsfpgaFlow_RxFlow) GetEntityData() *types.CommonEntityData {
    rxFlow.EntityData.YFilter = rxFlow.YFilter
    rxFlow.EntityData.YangName = "rx-flow"
    rxFlow.EntityData.BundleName = "cisco_ios_xr"
    rxFlow.EntityData.ParentYangName = "msfpga-flow"
    rxFlow.EntityData.SegmentPath = "rx-flow"
    rxFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxFlow.EntityData.Children = make(map[string]types.YChild)
    rxFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    rxFlow.EntityData.Leafs["flow-id"] = types.YLeaf{"FlowId", rxFlow.FlowId}
    rxFlow.EntityData.Leafs["valid"] = types.YLeaf{"Valid", rxFlow.Valid}
    rxFlow.EntityData.Leafs["is-egress"] = types.YLeaf{"IsEgress", rxFlow.IsEgress}
    rxFlow.EntityData.Leafs["in-use"] = types.YLeaf{"InUse", rxFlow.InUse}
    rxFlow.EntityData.Leafs["action"] = types.YLeaf{"Action", rxFlow.Action}
    rxFlow.EntityData.Leafs["smac-inuse"] = types.YLeaf{"SmacInuse", rxFlow.SmacInuse}
    rxFlow.EntityData.Leafs["dmac-inuse"] = types.YLeaf{"DmacInuse", rxFlow.DmacInuse}
    rxFlow.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", rxFlow.Ethertype}
    rxFlow.EntityData.Leafs["outer-vlan"] = types.YLeaf{"OuterVlan", rxFlow.OuterVlan}
    rxFlow.EntityData.Leafs["outer-vlan-up"] = types.YLeaf{"OuterVlanUp", rxFlow.OuterVlanUp}
    rxFlow.EntityData.Leafs["outer-vlan-tpid"] = types.YLeaf{"OuterVlanTpid", rxFlow.OuterVlanTpid}
    rxFlow.EntityData.Leafs["inner-vlan"] = types.YLeaf{"InnerVlan", rxFlow.InnerVlan}
    rxFlow.EntityData.Leafs["inner-vlan-up"] = types.YLeaf{"InnerVlanUp", rxFlow.InnerVlanUp}
    rxFlow.EntityData.Leafs["inner-vlan-tpid"] = types.YLeaf{"InnerVlanTpid", rxFlow.InnerVlanTpid}
    rxFlow.EntityData.Leafs["source-port"] = types.YLeaf{"SourcePort", rxFlow.SourcePort}
    rxFlow.EntityData.Leafs["source-port-chk"] = types.YLeaf{"SourcePortChk", rxFlow.SourcePortChk}
    rxFlow.EntityData.Leafs["sci-inuse"] = types.YLeaf{"SciInuse", rxFlow.SciInuse}
    rxFlow.EntityData.Leafs["sci"] = types.YLeaf{"Sci", rxFlow.Sci}
    rxFlow.EntityData.Leafs["match-pri"] = types.YLeaf{"MatchPri", rxFlow.MatchPri}
    rxFlow.EntityData.Leafs["is-ctrl-pkt"] = types.YLeaf{"IsCtrlPkt", rxFlow.IsCtrlPkt}
    rxFlow.EntityData.Leafs["ctrl-check"] = types.YLeaf{"CtrlCheck", rxFlow.CtrlCheck}
    rxFlow.EntityData.Leafs["match-untagged"] = types.YLeaf{"MatchUntagged", rxFlow.MatchUntagged}
    rxFlow.EntityData.Leafs["match-tagged"] = types.YLeaf{"MatchTagged", rxFlow.MatchTagged}
    rxFlow.EntityData.Leafs["match-bad-tag"] = types.YLeaf{"MatchBadTag", rxFlow.MatchBadTag}
    rxFlow.EntityData.Leafs["match-kay-tag"] = types.YLeaf{"MatchKayTag", rxFlow.MatchKayTag}
    rxFlow.EntityData.Leafs["tci-v"] = types.YLeaf{"TciV", rxFlow.TciV}
    rxFlow.EntityData.Leafs["tci-e-xr"] = types.YLeaf{"TciEXr", rxFlow.TciEXr}
    rxFlow.EntityData.Leafs["tci-sc"] = types.YLeaf{"TciSc", rxFlow.TciSc}
    rxFlow.EntityData.Leafs["tci-scb"] = types.YLeaf{"TciScb", rxFlow.TciScb}
    rxFlow.EntityData.Leafs["tci"] = types.YLeaf{"Tci", rxFlow.Tci}
    rxFlow.EntityData.Leafs["tci-c"] = types.YLeaf{"TciC", rxFlow.TciC}
    rxFlow.EntityData.Leafs["tci-an"] = types.YLeaf{"TciAn", rxFlow.TciAn}
    rxFlow.EntityData.Leafs["tci-an-chk"] = types.YLeaf{"TciAnChk", rxFlow.TciAnChk}
    rxFlow.EntityData.Leafs["tci-chk"] = types.YLeaf{"TciChk", rxFlow.TciChk}
    rxFlow.EntityData.Leafs["sai"] = types.YLeaf{"Sai", rxFlow.Sai}
    rxFlow.EntityData.Leafs["macsa"] = types.YLeaf{"Macsa", rxFlow.Macsa}
    rxFlow.EntityData.Leafs["macda"] = types.YLeaf{"Macda", rxFlow.Macda}
    return &(rxFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow
// ES200 Flow Information
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Flow Details.
    TxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow

    // Rx Flow Details.
    RxFlow MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow
}

func (es200Flow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow) GetEntityData() *types.CommonEntityData {
    es200Flow.EntityData.YFilter = es200Flow.YFilter
    es200Flow.EntityData.YangName = "es200-flow"
    es200Flow.EntityData.BundleName = "cisco_ios_xr"
    es200Flow.EntityData.ParentYangName = "ext"
    es200Flow.EntityData.SegmentPath = "es200-flow"
    es200Flow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es200Flow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es200Flow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es200Flow.EntityData.Children = make(map[string]types.YChild)
    es200Flow.EntityData.Children["tx-flow"] = types.YChild{"TxFlow", &es200Flow.TxFlow}
    es200Flow.EntityData.Children["rx-flow"] = types.YChild{"RxFlow", &es200Flow.RxFlow}
    es200Flow.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(es200Flow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow
// Tx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Number. The type is interface{} with range: 0..4294967295.
    FlowNo interface{}

    // Is Flow Enabled. The type is bool.
    IsFlowEnabled interface{}

    // Parsed EtherType to match could be 0 if Ethertype should'nt                
    // be matched can be 0x88E5 for MACSec tag. The type is interface{} with
    // range: 0..65535.
    Ethertype interface{}

    // VLAN ID for outer tag use this when             only one tag should be
    // matched. The type is interface{} with range: 0..65535.
    OuterVlanId interface{}

    // VLAN User Priority for outer tag  use            this when only one tag
    // should be matched. The type is interface{} with range: 0..255.
    OuterVlanUserPri interface{}

    // VLAN ID for inner tag used when two              VLAN Tags should be
    // matched. The type is interface{} with range: 0..65535.
    InnerVlanId interface{}

    // VLAN User priority for inner tag use            when matching two VLAN
    // tags. The type is interface{} with range: 0..255.
    InnerVlanUserPri interface{}

    // SCI to be matched value required for            ingress only, pass NULL for
    // egress. The type is interface{} with range: 0..18446744073709551615.
    Psci interface{}

    // priority for match 0-15(highest) . The type is interface{} with range:
    // 0..255.
    MatchPriority interface{}

    // value of 'v' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciV interface{}

    // value of 'es' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciEXr interface{}

    // value of 'sc' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciSc interface{}

    // value of 'scb' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciScb interface{}

    // value of 'e' in TCI to match (1bit ). The type is interface{} with range:
    // 0..255.
    Tci interface{}

    // value of 'c' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciC interface{}

    // TCI bits will be checked only when this          bit is enabled. All the
    // values of TCI bits       are mandatory when TCI check is used. The type is
    // bool.
    TciChk interface{}

    // Type of packet. See ethMscCfyEPktType_e. The type is string.
    PktType interface{}

    // No. of MPLS or VLAN tags See ethMscCfyETagNum_e . The type is string.
    TagNum interface{}

    // Dei to match for innner Vlan tag. The type is bool.
    InnerVlanDei interface{}

    // Dei to match for outer Vlan tag. The type is bool.
    OuterVlanDei interface{}

    // Service Instance id . The type is interface{} with range: 0..4294967295.
    PbbSid interface{}

    // Backbone Vlan id . The type is interface{} with range: 0..4294967295.
    PbbBvid interface{}

    // pcp . The type is interface{} with range: 0..255.
    PbbPcp interface{}

    // dei . The type is interface{} with range: 0..255.
    PbbDei interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls1Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls1Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls1Bos interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls2Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls2Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls2Bos interface{}

    // Plain bits to compare. Max values:               untagged pkt - 40 bits
    // after EthType             1 VLAN tag - 24 bits after parsed EthType       
    // 2 VLAN tags- 8 bits after parsed EthType         1 MPLS tag - 32 bits after
    // 1st tag               2 MPLS tags- 8 bits following after 2nd          or
    // atmost 5th MPLS tag                           PBB - 16 bits after C-SA     
    // PBB with VLAN tag - 16 bits of VLAN tag . The type is interface{} with
    // range: 0..18446744073709551615. Units are bit.
    PlainBits interface{}

    // No. of bits used in plainBits. The type is interface{} with range: 0..255.
    PlainBitsSize interface{}

    // Force the pkt as control pkt irrepective         of the results of control
    // packet detector. The type is bool.
    ForceCtrl interface{}

    // Drop the packet. The type is bool.
    Drop interface{}

    // DA mask. The type is interface{} with range: 0..18446744073709551615.
    MaskDa interface{}

    // Parsed EtherType mask. The type is interface{} with range: 0..4294967295.
    MaskEthertype interface{}

    // Plain Bits mask. The type is interface{} with range:
    // 0..18446744073709551615.
    MaskPlainBits interface{}

    // Pkts matching the Flow. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowHits interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (txFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_TxFlow) GetEntityData() *types.CommonEntityData {
    txFlow.EntityData.YFilter = txFlow.YFilter
    txFlow.EntityData.YangName = "tx-flow"
    txFlow.EntityData.BundleName = "cisco_ios_xr"
    txFlow.EntityData.ParentYangName = "es200-flow"
    txFlow.EntityData.SegmentPath = "tx-flow"
    txFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txFlow.EntityData.Children = make(map[string]types.YChild)
    txFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    txFlow.EntityData.Leafs["flow-no"] = types.YLeaf{"FlowNo", txFlow.FlowNo}
    txFlow.EntityData.Leafs["is-flow-enabled"] = types.YLeaf{"IsFlowEnabled", txFlow.IsFlowEnabled}
    txFlow.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", txFlow.Ethertype}
    txFlow.EntityData.Leafs["outer-vlan-id"] = types.YLeaf{"OuterVlanId", txFlow.OuterVlanId}
    txFlow.EntityData.Leafs["outer-vlan-user-pri"] = types.YLeaf{"OuterVlanUserPri", txFlow.OuterVlanUserPri}
    txFlow.EntityData.Leafs["inner-vlan-id"] = types.YLeaf{"InnerVlanId", txFlow.InnerVlanId}
    txFlow.EntityData.Leafs["inner-vlan-user-pri"] = types.YLeaf{"InnerVlanUserPri", txFlow.InnerVlanUserPri}
    txFlow.EntityData.Leafs["psci"] = types.YLeaf{"Psci", txFlow.Psci}
    txFlow.EntityData.Leafs["match-priority"] = types.YLeaf{"MatchPriority", txFlow.MatchPriority}
    txFlow.EntityData.Leafs["tci-v"] = types.YLeaf{"TciV", txFlow.TciV}
    txFlow.EntityData.Leafs["tci-e-xr"] = types.YLeaf{"TciEXr", txFlow.TciEXr}
    txFlow.EntityData.Leafs["tci-sc"] = types.YLeaf{"TciSc", txFlow.TciSc}
    txFlow.EntityData.Leafs["tci-scb"] = types.YLeaf{"TciScb", txFlow.TciScb}
    txFlow.EntityData.Leafs["tci"] = types.YLeaf{"Tci", txFlow.Tci}
    txFlow.EntityData.Leafs["tci-c"] = types.YLeaf{"TciC", txFlow.TciC}
    txFlow.EntityData.Leafs["tci-chk"] = types.YLeaf{"TciChk", txFlow.TciChk}
    txFlow.EntityData.Leafs["pkt-type"] = types.YLeaf{"PktType", txFlow.PktType}
    txFlow.EntityData.Leafs["tag-num"] = types.YLeaf{"TagNum", txFlow.TagNum}
    txFlow.EntityData.Leafs["inner-vlan-dei"] = types.YLeaf{"InnerVlanDei", txFlow.InnerVlanDei}
    txFlow.EntityData.Leafs["outer-vlan-dei"] = types.YLeaf{"OuterVlanDei", txFlow.OuterVlanDei}
    txFlow.EntityData.Leafs["pbb-sid"] = types.YLeaf{"PbbSid", txFlow.PbbSid}
    txFlow.EntityData.Leafs["pbb-bvid"] = types.YLeaf{"PbbBvid", txFlow.PbbBvid}
    txFlow.EntityData.Leafs["pbb-pcp"] = types.YLeaf{"PbbPcp", txFlow.PbbPcp}
    txFlow.EntityData.Leafs["pbb-dei"] = types.YLeaf{"PbbDei", txFlow.PbbDei}
    txFlow.EntityData.Leafs["mpls1-label"] = types.YLeaf{"Mpls1Label", txFlow.Mpls1Label}
    txFlow.EntityData.Leafs["mpls1-exp"] = types.YLeaf{"Mpls1Exp", txFlow.Mpls1Exp}
    txFlow.EntityData.Leafs["mpls1-bos"] = types.YLeaf{"Mpls1Bos", txFlow.Mpls1Bos}
    txFlow.EntityData.Leafs["mpls2-label"] = types.YLeaf{"Mpls2Label", txFlow.Mpls2Label}
    txFlow.EntityData.Leafs["mpls2-exp"] = types.YLeaf{"Mpls2Exp", txFlow.Mpls2Exp}
    txFlow.EntityData.Leafs["mpls2-bos"] = types.YLeaf{"Mpls2Bos", txFlow.Mpls2Bos}
    txFlow.EntityData.Leafs["plain-bits"] = types.YLeaf{"PlainBits", txFlow.PlainBits}
    txFlow.EntityData.Leafs["plain-bits-size"] = types.YLeaf{"PlainBitsSize", txFlow.PlainBitsSize}
    txFlow.EntityData.Leafs["force-ctrl"] = types.YLeaf{"ForceCtrl", txFlow.ForceCtrl}
    txFlow.EntityData.Leafs["drop"] = types.YLeaf{"Drop", txFlow.Drop}
    txFlow.EntityData.Leafs["mask-da"] = types.YLeaf{"MaskDa", txFlow.MaskDa}
    txFlow.EntityData.Leafs["mask-ethertype"] = types.YLeaf{"MaskEthertype", txFlow.MaskEthertype}
    txFlow.EntityData.Leafs["mask-plain-bits"] = types.YLeaf{"MaskPlainBits", txFlow.MaskPlainBits}
    txFlow.EntityData.Leafs["flow-hits"] = types.YLeaf{"FlowHits", txFlow.FlowHits}
    txFlow.EntityData.Leafs["macda"] = types.YLeaf{"Macda", txFlow.Macda}
    return &(txFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow
// Rx Flow Details
type MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flow Number. The type is interface{} with range: 0..4294967295.
    FlowNo interface{}

    // Is Flow Enabled. The type is bool.
    IsFlowEnabled interface{}

    // Parsed EtherType to match could be 0 if Ethertype should'nt                
    // be matched can be 0x88E5 for MACSec tag. The type is interface{} with
    // range: 0..65535.
    Ethertype interface{}

    // VLAN ID for outer tag use this when             only one tag should be
    // matched. The type is interface{} with range: 0..65535.
    OuterVlanId interface{}

    // VLAN User Priority for outer tag  use            this when only one tag
    // should be matched. The type is interface{} with range: 0..255.
    OuterVlanUserPri interface{}

    // VLAN ID for inner tag used when two              VLAN Tags should be
    // matched. The type is interface{} with range: 0..65535.
    InnerVlanId interface{}

    // VLAN User priority for inner tag use            when matching two VLAN
    // tags. The type is interface{} with range: 0..255.
    InnerVlanUserPri interface{}

    // SCI to be matched value required for            ingress only, pass NULL for
    // egress. The type is interface{} with range: 0..18446744073709551615.
    Psci interface{}

    // priority for match 0-15(highest) . The type is interface{} with range:
    // 0..255.
    MatchPriority interface{}

    // value of 'v' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciV interface{}

    // value of 'es' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciEXr interface{}

    // value of 'sc' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciSc interface{}

    // value of 'scb' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciScb interface{}

    // value of 'e' in TCI to match (1bit ). The type is interface{} with range:
    // 0..255.
    Tci interface{}

    // value of 'c' in TCI to match (1bit) . The type is interface{} with range:
    // 0..255.
    TciC interface{}

    // TCI bits will be checked only when this          bit is enabled. All the
    // values of TCI bits       are mandatory when TCI check is used. The type is
    // bool.
    TciChk interface{}

    // Type of packet. See ethMscCfyEPktType_e. The type is string.
    PktType interface{}

    // No. of MPLS or VLAN tags See ethMscCfyETagNum_e . The type is string.
    TagNum interface{}

    // Dei to match for innner Vlan tag. The type is bool.
    InnerVlanDei interface{}

    // Dei to match for outer Vlan tag. The type is bool.
    OuterVlanDei interface{}

    // Service Instance id . The type is interface{} with range: 0..4294967295.
    PbbSid interface{}

    // Backbone Vlan id . The type is interface{} with range: 0..4294967295.
    PbbBvid interface{}

    // pcp . The type is interface{} with range: 0..255.
    PbbPcp interface{}

    // dei . The type is interface{} with range: 0..255.
    PbbDei interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls1Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls1Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls1Bos interface{}

    // label . The type is interface{} with range: 0..4294967295.
    Mpls2Label interface{}

    // exp . The type is interface{} with range: 0..255.
    Mpls2Exp interface{}

    // botton of stack . The type is interface{} with range: 0..255.
    Mpls2Bos interface{}

    // Plain bits to compare. Max values:               untagged pkt - 40 bits
    // after EthType             1 VLAN tag - 24 bits after parsed EthType       
    // 2 VLAN tags- 8 bits after parsed EthType         1 MPLS tag - 32 bits after
    // 1st tag               2 MPLS tags- 8 bits following after 2nd          or
    // atmost 5th MPLS tag                           PBB - 16 bits after C-SA     
    // PBB with VLAN tag - 16 bits of VLAN tag . The type is interface{} with
    // range: 0..18446744073709551615. Units are bit.
    PlainBits interface{}

    // No. of bits used in plainBits. The type is interface{} with range: 0..255.
    PlainBitsSize interface{}

    // Force the pkt as control pkt irrepective         of the results of control
    // packet detector. The type is bool.
    ForceCtrl interface{}

    // Drop the packet. The type is bool.
    Drop interface{}

    // DA mask. The type is interface{} with range: 0..18446744073709551615.
    MaskDa interface{}

    // Parsed EtherType mask. The type is interface{} with range: 0..4294967295.
    MaskEthertype interface{}

    // Plain Bits mask. The type is interface{} with range:
    // 0..18446744073709551615.
    MaskPlainBits interface{}

    // Pkts matching the Flow. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowHits interface{}

    // MAC DA. The type is slice of interface{} with range: 0..255.
    Macda []interface{}
}

func (rxFlow *MacsecPlatform_Nodes_Node_Interfaces_Interface_HwFlowS_HwFlow_Ext_Es200Flow_RxFlow) GetEntityData() *types.CommonEntityData {
    rxFlow.EntityData.YFilter = rxFlow.YFilter
    rxFlow.EntityData.YangName = "rx-flow"
    rxFlow.EntityData.BundleName = "cisco_ios_xr"
    rxFlow.EntityData.ParentYangName = "es200-flow"
    rxFlow.EntityData.SegmentPath = "rx-flow"
    rxFlow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxFlow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxFlow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxFlow.EntityData.Children = make(map[string]types.YChild)
    rxFlow.EntityData.Leafs = make(map[string]types.YLeaf)
    rxFlow.EntityData.Leafs["flow-no"] = types.YLeaf{"FlowNo", rxFlow.FlowNo}
    rxFlow.EntityData.Leafs["is-flow-enabled"] = types.YLeaf{"IsFlowEnabled", rxFlow.IsFlowEnabled}
    rxFlow.EntityData.Leafs["ethertype"] = types.YLeaf{"Ethertype", rxFlow.Ethertype}
    rxFlow.EntityData.Leafs["outer-vlan-id"] = types.YLeaf{"OuterVlanId", rxFlow.OuterVlanId}
    rxFlow.EntityData.Leafs["outer-vlan-user-pri"] = types.YLeaf{"OuterVlanUserPri", rxFlow.OuterVlanUserPri}
    rxFlow.EntityData.Leafs["inner-vlan-id"] = types.YLeaf{"InnerVlanId", rxFlow.InnerVlanId}
    rxFlow.EntityData.Leafs["inner-vlan-user-pri"] = types.YLeaf{"InnerVlanUserPri", rxFlow.InnerVlanUserPri}
    rxFlow.EntityData.Leafs["psci"] = types.YLeaf{"Psci", rxFlow.Psci}
    rxFlow.EntityData.Leafs["match-priority"] = types.YLeaf{"MatchPriority", rxFlow.MatchPriority}
    rxFlow.EntityData.Leafs["tci-v"] = types.YLeaf{"TciV", rxFlow.TciV}
    rxFlow.EntityData.Leafs["tci-e-xr"] = types.YLeaf{"TciEXr", rxFlow.TciEXr}
    rxFlow.EntityData.Leafs["tci-sc"] = types.YLeaf{"TciSc", rxFlow.TciSc}
    rxFlow.EntityData.Leafs["tci-scb"] = types.YLeaf{"TciScb", rxFlow.TciScb}
    rxFlow.EntityData.Leafs["tci"] = types.YLeaf{"Tci", rxFlow.Tci}
    rxFlow.EntityData.Leafs["tci-c"] = types.YLeaf{"TciC", rxFlow.TciC}
    rxFlow.EntityData.Leafs["tci-chk"] = types.YLeaf{"TciChk", rxFlow.TciChk}
    rxFlow.EntityData.Leafs["pkt-type"] = types.YLeaf{"PktType", rxFlow.PktType}
    rxFlow.EntityData.Leafs["tag-num"] = types.YLeaf{"TagNum", rxFlow.TagNum}
    rxFlow.EntityData.Leafs["inner-vlan-dei"] = types.YLeaf{"InnerVlanDei", rxFlow.InnerVlanDei}
    rxFlow.EntityData.Leafs["outer-vlan-dei"] = types.YLeaf{"OuterVlanDei", rxFlow.OuterVlanDei}
    rxFlow.EntityData.Leafs["pbb-sid"] = types.YLeaf{"PbbSid", rxFlow.PbbSid}
    rxFlow.EntityData.Leafs["pbb-bvid"] = types.YLeaf{"PbbBvid", rxFlow.PbbBvid}
    rxFlow.EntityData.Leafs["pbb-pcp"] = types.YLeaf{"PbbPcp", rxFlow.PbbPcp}
    rxFlow.EntityData.Leafs["pbb-dei"] = types.YLeaf{"PbbDei", rxFlow.PbbDei}
    rxFlow.EntityData.Leafs["mpls1-label"] = types.YLeaf{"Mpls1Label", rxFlow.Mpls1Label}
    rxFlow.EntityData.Leafs["mpls1-exp"] = types.YLeaf{"Mpls1Exp", rxFlow.Mpls1Exp}
    rxFlow.EntityData.Leafs["mpls1-bos"] = types.YLeaf{"Mpls1Bos", rxFlow.Mpls1Bos}
    rxFlow.EntityData.Leafs["mpls2-label"] = types.YLeaf{"Mpls2Label", rxFlow.Mpls2Label}
    rxFlow.EntityData.Leafs["mpls2-exp"] = types.YLeaf{"Mpls2Exp", rxFlow.Mpls2Exp}
    rxFlow.EntityData.Leafs["mpls2-bos"] = types.YLeaf{"Mpls2Bos", rxFlow.Mpls2Bos}
    rxFlow.EntityData.Leafs["plain-bits"] = types.YLeaf{"PlainBits", rxFlow.PlainBits}
    rxFlow.EntityData.Leafs["plain-bits-size"] = types.YLeaf{"PlainBitsSize", rxFlow.PlainBitsSize}
    rxFlow.EntityData.Leafs["force-ctrl"] = types.YLeaf{"ForceCtrl", rxFlow.ForceCtrl}
    rxFlow.EntityData.Leafs["drop"] = types.YLeaf{"Drop", rxFlow.Drop}
    rxFlow.EntityData.Leafs["mask-da"] = types.YLeaf{"MaskDa", rxFlow.MaskDa}
    rxFlow.EntityData.Leafs["mask-ethertype"] = types.YLeaf{"MaskEthertype", rxFlow.MaskEthertype}
    rxFlow.EntityData.Leafs["mask-plain-bits"] = types.YLeaf{"MaskPlainBits", rxFlow.MaskPlainBits}
    rxFlow.EntityData.Leafs["flow-hits"] = types.YLeaf{"FlowHits", rxFlow.FlowHits}
    rxFlow.EntityData.Leafs["macda"] = types.YLeaf{"Macda", rxFlow.Macda}
    return &(rxFlow.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics
// The Software Statistics
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ext.
    Ext MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext
}

func (swStatistics *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics) GetEntityData() *types.CommonEntityData {
    swStatistics.EntityData.YFilter = swStatistics.YFilter
    swStatistics.EntityData.YangName = "sw-statistics"
    swStatistics.EntityData.BundleName = "cisco_ios_xr"
    swStatistics.EntityData.ParentYangName = "interface"
    swStatistics.EntityData.SegmentPath = "sw-statistics"
    swStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    swStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    swStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    swStatistics.EntityData.Children = make(map[string]types.YChild)
    swStatistics.EntityData.Children["ext"] = types.YChild{"Ext", &swStatistics.Ext}
    swStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(swStatistics.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext
// ext
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // type. The type is MacsecCard.
    Type_ interface{}

    // MSFPGA Stats.
    MsfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats

    // XLFPGA Stats.
    XlfpgaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats

    // ES200 Stats.
    Es200Stats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats
}

func (ext *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext) GetEntityData() *types.CommonEntityData {
    ext.EntityData.YFilter = ext.YFilter
    ext.EntityData.YangName = "ext"
    ext.EntityData.BundleName = "cisco_ios_xr"
    ext.EntityData.ParentYangName = "sw-statistics"
    ext.EntityData.SegmentPath = "ext"
    ext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ext.EntityData.Children = make(map[string]types.YChild)
    ext.EntityData.Children["msfpga-stats"] = types.YChild{"MsfpgaStats", &ext.MsfpgaStats}
    ext.EntityData.Children["xlfpga-stats"] = types.YChild{"XlfpgaStats", &ext.XlfpgaStats}
    ext.EntityData.Children["es200-stats"] = types.YChild{"Es200Stats", &ext.Es200Stats}
    ext.EntityData.Leafs = make(map[string]types.YLeaf)
    ext.EntityData.Leafs["type"] = types.YLeaf{"Type_", ext.Type_}
    return &(ext.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats
// MSFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
}

func (msfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats) GetEntityData() *types.CommonEntityData {
    msfpgaStats.EntityData.YFilter = msfpgaStats.YFilter
    msfpgaStats.EntityData.YangName = "msfpga-stats"
    msfpgaStats.EntityData.BundleName = "cisco_ios_xr"
    msfpgaStats.EntityData.ParentYangName = "ext"
    msfpgaStats.EntityData.SegmentPath = "msfpga-stats"
    msfpgaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    msfpgaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    msfpgaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    msfpgaStats.EntityData.Children = make(map[string]types.YChild)
    msfpgaStats.EntityData.Children["tx-sa-stats"] = types.YChild{"TxSaStats", &msfpgaStats.TxSaStats}
    msfpgaStats.EntityData.Children["rx-sa-stats"] = types.YChild{"RxSaStats", &msfpgaStats.RxSaStats}
    msfpgaStats.EntityData.Children["tx-interface-macsec-stats"] = types.YChild{"TxInterfaceMacsecStats", &msfpgaStats.TxInterfaceMacsecStats}
    msfpgaStats.EntityData.Children["rx-interface-macsec-stats"] = types.YChild{"RxInterfaceMacsecStats", &msfpgaStats.RxInterfaceMacsecStats}
    msfpgaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(msfpgaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Pkts Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsProtected interface{}

    // Tx Pkts Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncrypted interface{}

    // Tx Octets Protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsProtected interface{}

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncrypted interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxSaStats) GetEntityData() *types.CommonEntityData {
    txSaStats.EntityData.YFilter = txSaStats.YFilter
    txSaStats.EntityData.YangName = "tx-sa-stats"
    txSaStats.EntityData.BundleName = "cisco_ios_xr"
    txSaStats.EntityData.ParentYangName = "msfpga-stats"
    txSaStats.EntityData.SegmentPath = "tx-sa-stats"
    txSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSaStats.EntityData.Children = make(map[string]types.YChild)
    txSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txSaStats.EntityData.Leafs["out-pkts-protected"] = types.YLeaf{"OutPktsProtected", txSaStats.OutPktsProtected}
    txSaStats.EntityData.Leafs["out-pkts-encrypted"] = types.YLeaf{"OutPktsEncrypted", txSaStats.OutPktsEncrypted}
    txSaStats.EntityData.Leafs["out-octets-protected"] = types.YLeaf{"OutOctetsProtected", txSaStats.OutOctetsProtected}
    txSaStats.EntityData.Leafs["out-octets-encrypted"] = types.YLeaf{"OutOctetsEncrypted", txSaStats.OutOctetsEncrypted}
    return &(txSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // Rx Pkts Not Valid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsNotValid interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsInvalid interface{}

    // Rx Pkts OK. The type is interface{} with range: 0..18446744073709551615.
    InPktsOk interface{}

    // Rx Pkts Delayed. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsDelayed interface{}

    // Rx Pkts Late. The type is interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // Rx Pkts Unchecked. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnchecked interface{}

    // Rx Octets Validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecrypted interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxSaStats) GetEntityData() *types.CommonEntityData {
    rxSaStats.EntityData.YFilter = rxSaStats.YFilter
    rxSaStats.EntityData.YangName = "rx-sa-stats"
    rxSaStats.EntityData.BundleName = "cisco_ios_xr"
    rxSaStats.EntityData.ParentYangName = "msfpga-stats"
    rxSaStats.EntityData.SegmentPath = "rx-sa-stats"
    rxSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStats.EntityData.Children = make(map[string]types.YChild)
    rxSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStats.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxSaStats.InPktsUnusedSa}
    rxSaStats.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxSaStats.InPktsNotUsingSa}
    rxSaStats.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxSaStats.InPktsNotValid}
    rxSaStats.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxSaStats.InPktsInvalid}
    rxSaStats.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxSaStats.InPktsOk}
    rxSaStats.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxSaStats.InPktsDelayed}
    rxSaStats.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxSaStats.InPktsLate}
    rxSaStats.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxSaStats.InPktsUnchecked}
    rxSaStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxSaStats.InOctetsValidated}
    rxSaStats.EntityData.Leafs["in-octets-decrypted"] = types.YLeaf{"InOctetsDecrypted", rxSaStats.InOctetsDecrypted}
    return &(rxSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUncontrolled interface{}

    // Tx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktUntagged interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktTooLong interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_TxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    txInterfaceMacsecStats.EntityData.YFilter = txInterfaceMacsecStats.YFilter
    txInterfaceMacsecStats.EntityData.YangName = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txInterfaceMacsecStats.EntityData.ParentYangName = "msfpga-stats"
    txInterfaceMacsecStats.EntityData.SegmentPath = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-uncontrolled"] = types.YLeaf{"OutPktUncontrolled", txInterfaceMacsecStats.OutPktUncontrolled}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-untagged"] = types.YLeaf{"OutPktUntagged", txInterfaceMacsecStats.OutPktUntagged}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-too-long"] = types.YLeaf{"OutPktTooLong", txInterfaceMacsecStats.OutPktTooLong}
    return &(txInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Pkts Untagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUntagged interface{}

    // Rx Pkts Notag. The type is interface{} with range: 0..18446744073709551615.
    InPktNotag interface{}

    // Rx Pkts Bad tag. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktBadTag interface{}

    // Rx Pkts No Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // Rx Pkts Unknown Sci. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUnknownSci interface{}

    // Rx Pkts Tagged. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktTagged interface{}

    // Rx Pkts Over Run. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktOverrun interface{}

    // Rx Pkts Uncontrolled. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktUncontrolled interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_MsfpgaStats_RxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    rxInterfaceMacsecStats.EntityData.YFilter = rxInterfaceMacsecStats.YFilter
    rxInterfaceMacsecStats.EntityData.YangName = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxInterfaceMacsecStats.EntityData.ParentYangName = "msfpga-stats"
    rxInterfaceMacsecStats.EntityData.SegmentPath = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-untagged"] = types.YLeaf{"InPktUntagged", rxInterfaceMacsecStats.InPktUntagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-notag"] = types.YLeaf{"InPktNotag", rxInterfaceMacsecStats.InPktNotag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-bad-tag"] = types.YLeaf{"InPktBadTag", rxInterfaceMacsecStats.InPktBadTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-sci"] = types.YLeaf{"InPktNoSci", rxInterfaceMacsecStats.InPktNoSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-unknown-sci"] = types.YLeaf{"InPktUnknownSci", rxInterfaceMacsecStats.InPktUnknownSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-tagged"] = types.YLeaf{"InPktTagged", rxInterfaceMacsecStats.InPktTagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-overrun"] = types.YLeaf{"InPktOverrun", rxInterfaceMacsecStats.InPktOverrun}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-uncontrolled"] = types.YLeaf{"InPktUncontrolled", rxInterfaceMacsecStats.InPktUncontrolled}
    return &(rxInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats
// XLFPGA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SC and SA Level Stats.
    MacsecTxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats

    // Rx SC and SA Level Stats.
    MacsecRxStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats
}

func (xlfpgaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats) GetEntityData() *types.CommonEntityData {
    xlfpgaStats.EntityData.YFilter = xlfpgaStats.YFilter
    xlfpgaStats.EntityData.YangName = "xlfpga-stats"
    xlfpgaStats.EntityData.BundleName = "cisco_ios_xr"
    xlfpgaStats.EntityData.ParentYangName = "ext"
    xlfpgaStats.EntityData.SegmentPath = "xlfpga-stats"
    xlfpgaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xlfpgaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xlfpgaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xlfpgaStats.EntityData.Children = make(map[string]types.YChild)
    xlfpgaStats.EntityData.Children["macsec-tx-stats"] = types.YChild{"MacsecTxStats", &xlfpgaStats.MacsecTxStats}
    xlfpgaStats.EntityData.Children["macsec-rx-stats"] = types.YChild{"MacsecRxStats", &xlfpgaStats.MacsecRxStats}
    xlfpgaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(xlfpgaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats
// Tx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx Octets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedOctets interface{}

    // Tx Pkts Too Long. The type is interface{} with range:
    // 0..18446744073709551615.
    ScToolongPkts interface{}

    // Tx packets Encrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEncryptedPkts interface{}

    // Tx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Tx Overrun Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Tx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Tx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Tx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Current Tx AN. The type is interface{} with range: 0..18446744073709551615.
    CurrentAn interface{}

    // Current Tx SA Encrypted Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    SaEncryptedPkts interface{}
}

func (macsecTxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecTxStats) GetEntityData() *types.CommonEntityData {
    macsecTxStats.EntityData.YFilter = macsecTxStats.YFilter
    macsecTxStats.EntityData.YangName = "macsec-tx-stats"
    macsecTxStats.EntityData.BundleName = "cisco_ios_xr"
    macsecTxStats.EntityData.ParentYangName = "xlfpga-stats"
    macsecTxStats.EntityData.SegmentPath = "macsec-tx-stats"
    macsecTxStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecTxStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecTxStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecTxStats.EntityData.Children = make(map[string]types.YChild)
    macsecTxStats.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecTxStats.EntityData.Leafs["sc-encrypted-octets"] = types.YLeaf{"ScEncryptedOctets", macsecTxStats.ScEncryptedOctets}
    macsecTxStats.EntityData.Leafs["sc-toolong-pkts"] = types.YLeaf{"ScToolongPkts", macsecTxStats.ScToolongPkts}
    macsecTxStats.EntityData.Leafs["sc-encrypted-pkts"] = types.YLeaf{"ScEncryptedPkts", macsecTxStats.ScEncryptedPkts}
    macsecTxStats.EntityData.Leafs["sc-untagged-pkts"] = types.YLeaf{"ScUntaggedPkts", macsecTxStats.ScUntaggedPkts}
    macsecTxStats.EntityData.Leafs["sc-overrun-pkts"] = types.YLeaf{"ScOverrunPkts", macsecTxStats.ScOverrunPkts}
    macsecTxStats.EntityData.Leafs["sc-bypass-pkts"] = types.YLeaf{"ScBypassPkts", macsecTxStats.ScBypassPkts}
    macsecTxStats.EntityData.Leafs["sc-eapol-pkts"] = types.YLeaf{"ScEapolPkts", macsecTxStats.ScEapolPkts}
    macsecTxStats.EntityData.Leafs["sc-dropped-pkts"] = types.YLeaf{"ScDroppedPkts", macsecTxStats.ScDroppedPkts}
    macsecTxStats.EntityData.Leafs["current-an"] = types.YLeaf{"CurrentAn", macsecTxStats.CurrentAn}
    macsecTxStats.EntityData.Leafs["sa-encrypted-pkts"] = types.YLeaf{"SaEncryptedPkts", macsecTxStats.SaEncryptedPkts}
    return &(macsecTxStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats
// Rx SC and SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx Octets Decrypted. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDecryptedOctets interface{}

    // Rx No Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoTagPkts interface{}

    // Rx Untagged Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUntaggedPkts interface{}

    // Rx Bad Tag Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBadTagPkts interface{}

    // Rx Late Pkts. The type is interface{} with range: 0..18446744073709551615.
    ScLatePkts interface{}

    // Rx Delayed Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDelayedPkts interface{}

    // Rx Unchecked Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUncheckedPkts interface{}

    // Rx No SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNoSciPkts interface{}

    // Rx Unknown SCI Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnknownSciPkts interface{}

    // Rx Pkts Ok. The type is interface{} with range: 0..18446744073709551615.
    ScOkPkts interface{}

    // Rx Pkts Not Using SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotUsingPkts interface{}

    // Rx Pkts Unused SA. The type is interface{} with range:
    // 0..18446744073709551615.
    ScUnusedPkts interface{}

    // Rx Not Valid Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScNotValidPkts interface{}

    // Rx Pkts Invalid. The type is interface{} with range:
    // 0..18446744073709551615.
    ScInvalidPkts interface{}

    // Rx Overrun Pkts. The type is interface{} with range:
    // 0..18446744073709551615.
    ScOverrunPkts interface{}

    // Rx Bypass Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScBypassPkts interface{}

    // Rx Eapol Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScEapolPkts interface{}

    // Rx Dropped Packets. The type is interface{} with range:
    // 0..18446744073709551615.
    ScDroppedPkts interface{}

    // Rx SA Level Stats. The type is slice of
    // MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat.
    RxSaStat []MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
}

func (macsecRxStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats) GetEntityData() *types.CommonEntityData {
    macsecRxStats.EntityData.YFilter = macsecRxStats.YFilter
    macsecRxStats.EntityData.YangName = "macsec-rx-stats"
    macsecRxStats.EntityData.BundleName = "cisco_ios_xr"
    macsecRxStats.EntityData.ParentYangName = "xlfpga-stats"
    macsecRxStats.EntityData.SegmentPath = "macsec-rx-stats"
    macsecRxStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsecRxStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsecRxStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsecRxStats.EntityData.Children = make(map[string]types.YChild)
    macsecRxStats.EntityData.Children["rx-sa-stat"] = types.YChild{"RxSaStat", nil}
    for i := range macsecRxStats.RxSaStat {
        macsecRxStats.EntityData.Children[types.GetSegmentPath(&macsecRxStats.RxSaStat[i])] = types.YChild{"RxSaStat", &macsecRxStats.RxSaStat[i]}
    }
    macsecRxStats.EntityData.Leafs = make(map[string]types.YLeaf)
    macsecRxStats.EntityData.Leafs["sc-decrypted-octets"] = types.YLeaf{"ScDecryptedOctets", macsecRxStats.ScDecryptedOctets}
    macsecRxStats.EntityData.Leafs["sc-no-tag-pkts"] = types.YLeaf{"ScNoTagPkts", macsecRxStats.ScNoTagPkts}
    macsecRxStats.EntityData.Leafs["sc-untagged-pkts"] = types.YLeaf{"ScUntaggedPkts", macsecRxStats.ScUntaggedPkts}
    macsecRxStats.EntityData.Leafs["sc-bad-tag-pkts"] = types.YLeaf{"ScBadTagPkts", macsecRxStats.ScBadTagPkts}
    macsecRxStats.EntityData.Leafs["sc-late-pkts"] = types.YLeaf{"ScLatePkts", macsecRxStats.ScLatePkts}
    macsecRxStats.EntityData.Leafs["sc-delayed-pkts"] = types.YLeaf{"ScDelayedPkts", macsecRxStats.ScDelayedPkts}
    macsecRxStats.EntityData.Leafs["sc-unchecked-pkts"] = types.YLeaf{"ScUncheckedPkts", macsecRxStats.ScUncheckedPkts}
    macsecRxStats.EntityData.Leafs["sc-no-sci-pkts"] = types.YLeaf{"ScNoSciPkts", macsecRxStats.ScNoSciPkts}
    macsecRxStats.EntityData.Leafs["sc-unknown-sci-pkts"] = types.YLeaf{"ScUnknownSciPkts", macsecRxStats.ScUnknownSciPkts}
    macsecRxStats.EntityData.Leafs["sc-ok-pkts"] = types.YLeaf{"ScOkPkts", macsecRxStats.ScOkPkts}
    macsecRxStats.EntityData.Leafs["sc-not-using-pkts"] = types.YLeaf{"ScNotUsingPkts", macsecRxStats.ScNotUsingPkts}
    macsecRxStats.EntityData.Leafs["sc-unused-pkts"] = types.YLeaf{"ScUnusedPkts", macsecRxStats.ScUnusedPkts}
    macsecRxStats.EntityData.Leafs["sc-not-valid-pkts"] = types.YLeaf{"ScNotValidPkts", macsecRxStats.ScNotValidPkts}
    macsecRxStats.EntityData.Leafs["sc-invalid-pkts"] = types.YLeaf{"ScInvalidPkts", macsecRxStats.ScInvalidPkts}
    macsecRxStats.EntityData.Leafs["sc-overrun-pkts"] = types.YLeaf{"ScOverrunPkts", macsecRxStats.ScOverrunPkts}
    macsecRxStats.EntityData.Leafs["sc-bypass-pkts"] = types.YLeaf{"ScBypassPkts", macsecRxStats.ScBypassPkts}
    macsecRxStats.EntityData.Leafs["sc-eapol-pkts"] = types.YLeaf{"ScEapolPkts", macsecRxStats.ScEapolPkts}
    macsecRxStats.EntityData.Leafs["sc-dropped-pkts"] = types.YLeaf{"ScDroppedPkts", macsecRxStats.ScDroppedPkts}
    return &(macsecRxStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat
// Rx SA Level Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current Rx AN. The type is interface{} with range: 0..18446744073709551615.
    An interface{}

    // Rx Ok Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaOkPkts interface{}

    // Rx Pkts not using SA for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotUsingPkts interface{}

    // Rx Pkts Unused Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaUnusedPkts interface{}

    // Rx Not Valid Pkts for Current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaNotValidPkts interface{}

    // Rx Invalid Pkts for current AN. The type is interface{} with range:
    // 0..18446744073709551615.
    SaInvalidPkts interface{}
}

func (rxSaStat *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_XlfpgaStats_MacsecRxStats_RxSaStat) GetEntityData() *types.CommonEntityData {
    rxSaStat.EntityData.YFilter = rxSaStat.YFilter
    rxSaStat.EntityData.YangName = "rx-sa-stat"
    rxSaStat.EntityData.BundleName = "cisco_ios_xr"
    rxSaStat.EntityData.ParentYangName = "macsec-rx-stats"
    rxSaStat.EntityData.SegmentPath = "rx-sa-stat"
    rxSaStat.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStat.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStat.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStat.EntityData.Children = make(map[string]types.YChild)
    rxSaStat.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStat.EntityData.Leafs["an"] = types.YLeaf{"An", rxSaStat.An}
    rxSaStat.EntityData.Leafs["sa-ok-pkts"] = types.YLeaf{"SaOkPkts", rxSaStat.SaOkPkts}
    rxSaStat.EntityData.Leafs["sa-not-using-pkts"] = types.YLeaf{"SaNotUsingPkts", rxSaStat.SaNotUsingPkts}
    rxSaStat.EntityData.Leafs["sa-unused-pkts"] = types.YLeaf{"SaUnusedPkts", rxSaStat.SaUnusedPkts}
    rxSaStat.EntityData.Leafs["sa-not-valid-pkts"] = types.YLeaf{"SaNotValidPkts", rxSaStat.SaNotValidPkts}
    rxSaStat.EntityData.Leafs["sa-invalid-pkts"] = types.YLeaf{"SaInvalidPkts", rxSaStat.SaInvalidPkts}
    return &(rxSaStat.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats
// ES200 Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx SA Stats.
    TxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats

    // Rx SA Stats.
    RxSaStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats

    // Tx SC Macsec Stats.
    TxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats

    // Rx SC Macsec Stats.
    RxScMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats

    // Tx interface Macsec Stats.
    TxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats

    // Rx interface Macsec Stats.
    RxInterfaceMacsecStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats

    // Port level TX Stats.
    TxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats

    // Port level RX Stats.
    RxPortStats MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats
}

func (es200Stats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats) GetEntityData() *types.CommonEntityData {
    es200Stats.EntityData.YFilter = es200Stats.YFilter
    es200Stats.EntityData.YangName = "es200-stats"
    es200Stats.EntityData.BundleName = "cisco_ios_xr"
    es200Stats.EntityData.ParentYangName = "ext"
    es200Stats.EntityData.SegmentPath = "es200-stats"
    es200Stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    es200Stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    es200Stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    es200Stats.EntityData.Children = make(map[string]types.YChild)
    es200Stats.EntityData.Children["tx-sa-stats"] = types.YChild{"TxSaStats", &es200Stats.TxSaStats}
    es200Stats.EntityData.Children["rx-sa-stats"] = types.YChild{"RxSaStats", &es200Stats.RxSaStats}
    es200Stats.EntityData.Children["tx-sc-macsec-stats"] = types.YChild{"TxScMacsecStats", &es200Stats.TxScMacsecStats}
    es200Stats.EntityData.Children["rx-sc-macsec-stats"] = types.YChild{"RxScMacsecStats", &es200Stats.RxScMacsecStats}
    es200Stats.EntityData.Children["tx-interface-macsec-stats"] = types.YChild{"TxInterfaceMacsecStats", &es200Stats.TxInterfaceMacsecStats}
    es200Stats.EntityData.Children["rx-interface-macsec-stats"] = types.YChild{"RxInterfaceMacsecStats", &es200Stats.RxInterfaceMacsecStats}
    es200Stats.EntityData.Children["tx-port-stats"] = types.YChild{"TxPortStats", &es200Stats.TxPortStats}
    es200Stats.EntityData.Children["rx-port-stats"] = types.YChild{"RxPortStats", &es200Stats.RxPortStats}
    es200Stats.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(es200Stats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats
// Tx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // packets exceeding egress MTU. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsTooLong interface{}

    // packets encrypted/protected. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsEncryptedProtected interface{}

    // octets1 encrypted/protected ?. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsEncryptedProtected1 interface{}
}

func (txSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxSaStats) GetEntityData() *types.CommonEntityData {
    txSaStats.EntityData.YFilter = txSaStats.YFilter
    txSaStats.EntityData.YangName = "tx-sa-stats"
    txSaStats.EntityData.BundleName = "cisco_ios_xr"
    txSaStats.EntityData.ParentYangName = "es200-stats"
    txSaStats.EntityData.SegmentPath = "tx-sa-stats"
    txSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txSaStats.EntityData.Children = make(map[string]types.YChild)
    txSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txSaStats.EntityData.Leafs["out-pkts-too-long"] = types.YLeaf{"OutPktsTooLong", txSaStats.OutPktsTooLong}
    txSaStats.EntityData.Leafs["out-pkts-encrypted-protected"] = types.YLeaf{"OutPktsEncryptedProtected", txSaStats.OutPktsEncryptedProtected}
    txSaStats.EntityData.Leafs["out-octets-encrypted-protected1"] = types.YLeaf{"OutOctetsEncryptedProtected1", txSaStats.OutOctetsEncryptedProtected1}
    return &(txSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats
// Rx SA Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // frame not valid & validateFrames disabled. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsUnchecked interface{}

    // PN of packet outside replay window & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsDelayed interface{}

    // PN of packet outside replay window & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsLate interface{}

    // packets with no error. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsOk interface{}

    // packet not valid & validateFrames !strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsInvalid interface{}

    // packet not valid & validateFrames strict. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktsNotValid interface{}

    // packet assigned to SA not in use & validateFrames strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsNotUsingSa interface{}

    // packet assigned to SA not in use & validateFrames !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUnusedSa interface{}

    // octets1 decrypted/validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsDecryptedValidated1 interface{}

    // octets validated. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsValidated interface{}
}

func (rxSaStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxSaStats) GetEntityData() *types.CommonEntityData {
    rxSaStats.EntityData.YFilter = rxSaStats.YFilter
    rxSaStats.EntityData.YangName = "rx-sa-stats"
    rxSaStats.EntityData.BundleName = "cisco_ios_xr"
    rxSaStats.EntityData.ParentYangName = "es200-stats"
    rxSaStats.EntityData.SegmentPath = "rx-sa-stats"
    rxSaStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxSaStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxSaStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxSaStats.EntityData.Children = make(map[string]types.YChild)
    rxSaStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxSaStats.EntityData.Leafs["in-pkts-unchecked"] = types.YLeaf{"InPktsUnchecked", rxSaStats.InPktsUnchecked}
    rxSaStats.EntityData.Leafs["in-pkts-delayed"] = types.YLeaf{"InPktsDelayed", rxSaStats.InPktsDelayed}
    rxSaStats.EntityData.Leafs["in-pkts-late"] = types.YLeaf{"InPktsLate", rxSaStats.InPktsLate}
    rxSaStats.EntityData.Leafs["in-pkts-ok"] = types.YLeaf{"InPktsOk", rxSaStats.InPktsOk}
    rxSaStats.EntityData.Leafs["in-pkts-invalid"] = types.YLeaf{"InPktsInvalid", rxSaStats.InPktsInvalid}
    rxSaStats.EntityData.Leafs["in-pkts-not-valid"] = types.YLeaf{"InPktsNotValid", rxSaStats.InPktsNotValid}
    rxSaStats.EntityData.Leafs["in-pkts-not-using-sa"] = types.YLeaf{"InPktsNotUsingSa", rxSaStats.InPktsNotUsingSa}
    rxSaStats.EntityData.Leafs["in-pkts-unused-sa"] = types.YLeaf{"InPktsUnusedSa", rxSaStats.InPktsUnusedSa}
    rxSaStats.EntityData.Leafs["in-octets-decrypted-validated1"] = types.YLeaf{"InOctetsDecryptedValidated1", rxSaStats.InOctetsDecryptedValidated1}
    rxSaStats.EntityData.Leafs["in-octets-validated"] = types.YLeaf{"InOctetsValidated", rxSaStats.InOctetsValidated}
    return &(rxSaStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats
// Tx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    OutPktsSaNotInUse interface{}
}

func (txScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxScMacsecStats) GetEntityData() *types.CommonEntityData {
    txScMacsecStats.EntityData.YFilter = txScMacsecStats.YFilter
    txScMacsecStats.EntityData.YangName = "tx-sc-macsec-stats"
    txScMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txScMacsecStats.EntityData.ParentYangName = "es200-stats"
    txScMacsecStats.EntityData.SegmentPath = "tx-sc-macsec-stats"
    txScMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txScMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txScMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txScMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txScMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txScMacsecStats.EntityData.Leafs["out-pkts-sa-not-in-use"] = types.YLeaf{"OutPktsSaNotInUse", txScMacsecStats.OutPktsSaNotInUse}
    return &(txScMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats
// Rx SC Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packets received with SA not in use. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsSaNotInUse interface{}
}

func (rxScMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxScMacsecStats) GetEntityData() *types.CommonEntityData {
    rxScMacsecStats.EntityData.YFilter = rxScMacsecStats.YFilter
    rxScMacsecStats.EntityData.YangName = "rx-sc-macsec-stats"
    rxScMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxScMacsecStats.EntityData.ParentYangName = "es200-stats"
    rxScMacsecStats.EntityData.SegmentPath = "rx-sc-macsec-stats"
    rxScMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxScMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxScMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxScMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxScMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxScMacsecStats.EntityData.Leafs["in-pkts-sa-not-in-use"] = types.YLeaf{"InPktsSaNotInUse", rxScMacsecStats.InPktsSaNotInUse}
    return &(rxScMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats
// Tx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // egress packet that is classified as control packet. The type is interface{}
    // with range: 0..18446744073709551615.
    OutPktCtrl interface{}

    // egress packet to go out untagged when protectFrames not set. The type is
    // interface{} with range: 0..18446744073709551615.
    OutPktsUntagged interface{}

    // Octets tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsUnctrl interface{}

    // Octets tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCtrl interface{}

    // Octets tx on common port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutOctetsCommon interface{}

    // Unicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsUnctrl interface{}

    // Unicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutUcastPktsCtrl interface{}

    // Multicast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsUnctrl interface{}

    // Multicast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutMcastPktsCtrl interface{}

    // Broadcast pkts tx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsUnctrl interface{}

    // Broadcast pkts tx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    OutBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    OutRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    OutRxErrPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsClass interface{}

    // Packets dropped due to overflow in  processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    OutDropPktsData interface{}
}

func (txInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    txInterfaceMacsecStats.EntityData.YFilter = txInterfaceMacsecStats.YFilter
    txInterfaceMacsecStats.EntityData.YangName = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    txInterfaceMacsecStats.EntityData.ParentYangName = "es200-stats"
    txInterfaceMacsecStats.EntityData.SegmentPath = "tx-interface-macsec-stats"
    txInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    txInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txInterfaceMacsecStats.EntityData.Leafs["transform-error-pkts"] = types.YLeaf{"TransformErrorPkts", txInterfaceMacsecStats.TransformErrorPkts}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkt-ctrl"] = types.YLeaf{"OutPktCtrl", txInterfaceMacsecStats.OutPktCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-pkts-untagged"] = types.YLeaf{"OutPktsUntagged", txInterfaceMacsecStats.OutPktsUntagged}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-unctrl"] = types.YLeaf{"OutOctetsUnctrl", txInterfaceMacsecStats.OutOctetsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-ctrl"] = types.YLeaf{"OutOctetsCtrl", txInterfaceMacsecStats.OutOctetsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-octets-common"] = types.YLeaf{"OutOctetsCommon", txInterfaceMacsecStats.OutOctetsCommon}
    txInterfaceMacsecStats.EntityData.Leafs["out-ucast-pkts-unctrl"] = types.YLeaf{"OutUcastPktsUnctrl", txInterfaceMacsecStats.OutUcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-ucast-pkts-ctrl"] = types.YLeaf{"OutUcastPktsCtrl", txInterfaceMacsecStats.OutUcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-mcast-pkts-unctrl"] = types.YLeaf{"OutMcastPktsUnctrl", txInterfaceMacsecStats.OutMcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-mcast-pkts-ctrl"] = types.YLeaf{"OutMcastPktsCtrl", txInterfaceMacsecStats.OutMcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-bcast-pkts-unctrl"] = types.YLeaf{"OutBcastPktsUnctrl", txInterfaceMacsecStats.OutBcastPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-bcast-pkts-ctrl"] = types.YLeaf{"OutBcastPktsCtrl", txInterfaceMacsecStats.OutBcastPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-drop-pkts-unctrl"] = types.YLeaf{"OutRxDropPktsUnctrl", txInterfaceMacsecStats.OutRxDropPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-drop-pkts-ctrl"] = types.YLeaf{"OutRxDropPktsCtrl", txInterfaceMacsecStats.OutRxDropPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-err-pkts-unctrl"] = types.YLeaf{"OutRxErrPktsUnctrl", txInterfaceMacsecStats.OutRxErrPktsUnctrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-rx-err-pkts-ctrl"] = types.YLeaf{"OutRxErrPktsCtrl", txInterfaceMacsecStats.OutRxErrPktsCtrl}
    txInterfaceMacsecStats.EntityData.Leafs["out-drop-pkts-class"] = types.YLeaf{"OutDropPktsClass", txInterfaceMacsecStats.OutDropPktsClass}
    txInterfaceMacsecStats.EntityData.Leafs["out-drop-pkts-data"] = types.YLeaf{"OutDropPktsData", txInterfaceMacsecStats.OutDropPktsData}
    return &(txInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats
// Rx interface Macsec Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // counter to count internal errors in the MACSec core. The type is
    // interface{} with range: 0..18446744073709551615.
    TransformErrorPkts interface{}

    // ingress packet that is classified as control packet. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktCtrl interface{}

    // ingress packet untagged & validateFrames is strict. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktNoTag interface{}

    // ingress packet untagged & validateFrames is  !strict. The type is
    // interface{} with range: 0..18446744073709551615.
    InPktsUntagged interface{}

    // ingress frames received with an invalid MACSec tag or ICV                  
    // added with next one gives InPktsSCIMiss. The type is interface{} with
    // range: 0..18446744073709551615.
    InPktBadTag interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktNoSci interface{}

    // correctly tagged ingress frames for which no valid SC found &              
    // validateFrames is !strict. The type is interface{} with range:
    // 0..18446744073709551615.
    InPktsUnknownSci interface{}

    // ingress packets that are control or KaY packets. The type is interface{}
    // with range: 0..18446744073709551615.
    InPktsTaggedCtrl interface{}

    // Octets rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsUnctrl interface{}

    // Octets rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InOctetsCtrl interface{}

    // Unicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsUnctrl interface{}

    // Unicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InUcastPktsCtrl interface{}

    // Multicast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsUnctrl interface{}

    // Multicast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InMcastPktsCtrl interface{}

    // Broadcast pkts rx on uncontrolled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsUnctrl interface{}

    // Broadcast pkts rx on controlled port. The type is interface{} with range:
    // 0..18446744073709551615.
    InBcastPktsCtrl interface{}

    // Control pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsUnctrl interface{}

    // Data pkts dropped due to overrun. The type is interface{} with range:
    // 0..18446744073709551615.
    InRxDropPktsCtrl interface{}

    // Control pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsUnctrl interface{}

    // Data pkts error-terminated due to overrun. The type is interface{} with
    // range: 0..18446744073709551615.
    InRxErrorPktsCtrl interface{}

    // Packets dropped due to overflow in classification pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsClass interface{}

    // Packets dropped due to overflow in processing pipeline. The type is
    // interface{} with range: 0..18446744073709551615.
    InDropPktsData interface{}
}

func (rxInterfaceMacsecStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxInterfaceMacsecStats) GetEntityData() *types.CommonEntityData {
    rxInterfaceMacsecStats.EntityData.YFilter = rxInterfaceMacsecStats.YFilter
    rxInterfaceMacsecStats.EntityData.YangName = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.BundleName = "cisco_ios_xr"
    rxInterfaceMacsecStats.EntityData.ParentYangName = "es200-stats"
    rxInterfaceMacsecStats.EntityData.SegmentPath = "rx-interface-macsec-stats"
    rxInterfaceMacsecStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxInterfaceMacsecStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxInterfaceMacsecStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxInterfaceMacsecStats.EntityData.Children = make(map[string]types.YChild)
    rxInterfaceMacsecStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxInterfaceMacsecStats.EntityData.Leafs["transform-error-pkts"] = types.YLeaf{"TransformErrorPkts", rxInterfaceMacsecStats.TransformErrorPkts}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-ctrl"] = types.YLeaf{"InPktCtrl", rxInterfaceMacsecStats.InPktCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-tag"] = types.YLeaf{"InPktNoTag", rxInterfaceMacsecStats.InPktNoTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-untagged"] = types.YLeaf{"InPktsUntagged", rxInterfaceMacsecStats.InPktsUntagged}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-bad-tag"] = types.YLeaf{"InPktBadTag", rxInterfaceMacsecStats.InPktBadTag}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkt-no-sci"] = types.YLeaf{"InPktNoSci", rxInterfaceMacsecStats.InPktNoSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-unknown-sci"] = types.YLeaf{"InPktsUnknownSci", rxInterfaceMacsecStats.InPktsUnknownSci}
    rxInterfaceMacsecStats.EntityData.Leafs["in-pkts-tagged-ctrl"] = types.YLeaf{"InPktsTaggedCtrl", rxInterfaceMacsecStats.InPktsTaggedCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-octets-unctrl"] = types.YLeaf{"InOctetsUnctrl", rxInterfaceMacsecStats.InOctetsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-octets-ctrl"] = types.YLeaf{"InOctetsCtrl", rxInterfaceMacsecStats.InOctetsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-ucast-pkts-unctrl"] = types.YLeaf{"InUcastPktsUnctrl", rxInterfaceMacsecStats.InUcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-ucast-pkts-ctrl"] = types.YLeaf{"InUcastPktsCtrl", rxInterfaceMacsecStats.InUcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-mcast-pkts-unctrl"] = types.YLeaf{"InMcastPktsUnctrl", rxInterfaceMacsecStats.InMcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-mcast-pkts-ctrl"] = types.YLeaf{"InMcastPktsCtrl", rxInterfaceMacsecStats.InMcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-bcast-pkts-unctrl"] = types.YLeaf{"InBcastPktsUnctrl", rxInterfaceMacsecStats.InBcastPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-bcast-pkts-ctrl"] = types.YLeaf{"InBcastPktsCtrl", rxInterfaceMacsecStats.InBcastPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-drop-pkts-unctrl"] = types.YLeaf{"InRxDropPktsUnctrl", rxInterfaceMacsecStats.InRxDropPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-drop-pkts-ctrl"] = types.YLeaf{"InRxDropPktsCtrl", rxInterfaceMacsecStats.InRxDropPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-error-pkts-unctrl"] = types.YLeaf{"InRxErrorPktsUnctrl", rxInterfaceMacsecStats.InRxErrorPktsUnctrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-rx-error-pkts-ctrl"] = types.YLeaf{"InRxErrorPktsCtrl", rxInterfaceMacsecStats.InRxErrorPktsCtrl}
    rxInterfaceMacsecStats.EntityData.Leafs["in-drop-pkts-class"] = types.YLeaf{"InDropPktsClass", rxInterfaceMacsecStats.InDropPktsClass}
    rxInterfaceMacsecStats.EntityData.Leafs["in-drop-pkts-data"] = types.YLeaf{"InDropPktsData", rxInterfaceMacsecStats.InDropPktsData}
    return &(rxInterfaceMacsecStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats
// Port level TX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (txPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_TxPortStats) GetEntityData() *types.CommonEntityData {
    txPortStats.EntityData.YFilter = txPortStats.YFilter
    txPortStats.EntityData.YangName = "tx-port-stats"
    txPortStats.EntityData.BundleName = "cisco_ios_xr"
    txPortStats.EntityData.ParentYangName = "es200-stats"
    txPortStats.EntityData.SegmentPath = "tx-port-stats"
    txPortStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txPortStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txPortStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txPortStats.EntityData.Children = make(map[string]types.YChild)
    txPortStats.EntityData.Leafs = make(map[string]types.YLeaf)
    txPortStats.EntityData.Leafs["multi-flow-match"] = types.YLeaf{"MultiFlowMatch", txPortStats.MultiFlowMatch}
    txPortStats.EntityData.Leafs["parser-dropped"] = types.YLeaf{"ParserDropped", txPortStats.ParserDropped}
    txPortStats.EntityData.Leafs["flow-miss"] = types.YLeaf{"FlowMiss", txPortStats.FlowMiss}
    txPortStats.EntityData.Leafs["pkts-ctrl"] = types.YLeaf{"PktsCtrl", txPortStats.PktsCtrl}
    txPortStats.EntityData.Leafs["pkts-data"] = types.YLeaf{"PktsData", txPortStats.PktsData}
    txPortStats.EntityData.Leafs["pkts-dropped"] = types.YLeaf{"PktsDropped", txPortStats.PktsDropped}
    txPortStats.EntityData.Leafs["pkts-err-in"] = types.YLeaf{"PktsErrIn", txPortStats.PktsErrIn}
    return &(txPortStats.EntityData)
}

// MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats
// Port level RX Stats
type MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pkts matching multiple flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    MultiFlowMatch interface{}

    // Pkts dropped by header parser as invalid. The type is interface{} with
    // range: 0..18446744073709551615.
    ParserDropped interface{}

    // Pkts matching none of flow entries. The type is interface{} with range:
    // 0..18446744073709551615.
    FlowMiss interface{}

    // Control pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsCtrl interface{}

    // Data pkts forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsData interface{}

    // Pkts dropped by classifier. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsDropped interface{}

    // Pkts received with an error indication. The type is interface{} with range:
    // 0..18446744073709551615.
    PktsErrIn interface{}
}

func (rxPortStats *MacsecPlatform_Nodes_Node_Interfaces_Interface_SwStatistics_Ext_Es200Stats_RxPortStats) GetEntityData() *types.CommonEntityData {
    rxPortStats.EntityData.YFilter = rxPortStats.YFilter
    rxPortStats.EntityData.YangName = "rx-port-stats"
    rxPortStats.EntityData.BundleName = "cisco_ios_xr"
    rxPortStats.EntityData.ParentYangName = "es200-stats"
    rxPortStats.EntityData.SegmentPath = "rx-port-stats"
    rxPortStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxPortStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxPortStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxPortStats.EntityData.Children = make(map[string]types.YChild)
    rxPortStats.EntityData.Leafs = make(map[string]types.YLeaf)
    rxPortStats.EntityData.Leafs["multi-flow-match"] = types.YLeaf{"MultiFlowMatch", rxPortStats.MultiFlowMatch}
    rxPortStats.EntityData.Leafs["parser-dropped"] = types.YLeaf{"ParserDropped", rxPortStats.ParserDropped}
    rxPortStats.EntityData.Leafs["flow-miss"] = types.YLeaf{"FlowMiss", rxPortStats.FlowMiss}
    rxPortStats.EntityData.Leafs["pkts-ctrl"] = types.YLeaf{"PktsCtrl", rxPortStats.PktsCtrl}
    rxPortStats.EntityData.Leafs["pkts-data"] = types.YLeaf{"PktsData", rxPortStats.PktsData}
    rxPortStats.EntityData.Leafs["pkts-dropped"] = types.YLeaf{"PktsDropped", rxPortStats.PktsDropped}
    rxPortStats.EntityData.Leafs["pkts-err-in"] = types.YLeaf{"PktsErrIn", rxPortStats.PktsErrIn}
    return &(rxPortStats.EntityData)
}

