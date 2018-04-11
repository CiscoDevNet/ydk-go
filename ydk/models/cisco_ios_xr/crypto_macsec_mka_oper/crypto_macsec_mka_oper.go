// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-mka package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec: Macsec operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_macsec_mka_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_macsec_mka_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-macsec-mka-oper macsec}", reflect.TypeOf(Macsec{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-macsec-mka-oper:macsec", reflect.TypeOf(Macsec{}))
}

// Macsec
// Macsec operational data
type Macsec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MKA Data.
    Mka Macsec_Mka
}

func (macsec *Macsec) GetEntityData() *types.CommonEntityData {
    macsec.EntityData.YFilter = macsec.YFilter
    macsec.EntityData.YangName = "macsec"
    macsec.EntityData.BundleName = "cisco_ios_xr"
    macsec.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-macsec-mka-oper"
    macsec.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec"
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = make(map[string]types.YChild)
    macsec.EntityData.Children["mka"] = types.YChild{"Mka", &macsec.Mka}
    macsec.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macsec.EntityData)
}

// Macsec_Mka
// MKA Data
type Macsec_Mka struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MKA Data.
    Interfaces Macsec_Mka_Interfaces
}

func (mka *Macsec_Mka) GetEntityData() *types.CommonEntityData {
    mka.EntityData.YFilter = mka.YFilter
    mka.EntityData.YangName = "mka"
    mka.EntityData.BundleName = "cisco_ios_xr"
    mka.EntityData.ParentYangName = "macsec"
    mka.EntityData.SegmentPath = "mka"
    mka.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mka.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mka.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mka.EntityData.Children = make(map[string]types.YChild)
    mka.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &mka.Interfaces}
    mka.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mka.EntityData)
}

// Macsec_Mka_Interfaces
// MKA Data
type Macsec_Mka_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MKA Data for the Interface. The type is slice of
    // Macsec_Mka_Interfaces_Interface_.
    Interface_ []Macsec_Mka_Interfaces_Interface
}

func (interfaces *Macsec_Mka_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "mka"
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

// Macsec_Mka_Interfaces_Interface
// MKA Data for the Interface
type Macsec_Mka_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // MKA Session Data.
    Session Macsec_Mka_Interfaces_Interface_Session
}

func (self *Macsec_Mka_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[name='" + fmt.Sprintf("%v", self.Name) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["session"] = types.YChild{"Session", &self.Session}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["name"] = types.YLeaf{"Name", self.Name}
    return &(self.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session
// MKA Session Data
type Macsec_Mka_Interfaces_Interface_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session summary.
    SessionSummary Macsec_Mka_Interfaces_Interface_Session_SessionSummary

    // Virtual Pointer Info.
    Vp Macsec_Mka_Interfaces_Interface_Session_Vp

    // CA List for a Session. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca.
    Ca []Macsec_Mka_Interfaces_Interface_Session_Ca
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "interface"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["session-summary"] = types.YChild{"SessionSummary", &session.SessionSummary}
    session.EntityData.Children["vp"] = types.YChild{"Vp", &session.Vp}
    session.EntityData.Children["ca"] = types.YChild{"Ca", nil}
    for i := range session.Ca {
        session.EntityData.Children[types.GetSegmentPath(&session.Ca[i])] = types.YChild{"Ca", &session.Ca[i]}
    }
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary
// Session summary
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macsec configured interface. The type is string.
    InterfaceName interface{}

    // Is Inherited Policy. The type is bool.
    InheritedPolicy interface{}

    // Policy Name. The type is string.
    Policy interface{}

    // Key Server Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // My MAC. The type is string.
    MyMac interface{}

    // Delay Protect. The type is bool.
    DelayProtection interface{}

    // Replay Protect. The type is bool.
    ReplayProtect interface{}

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    WindowSize interface{}

    // IncludeICVIndicator. The type is bool.
    IncludeIcvIndicator interface{}

    // Confidentiality Offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Alogorithm Agility. The type is interface{} with range: 0..4294967295.
    AlgoAgility interface{}

    // MACSec Capability. The type is interface{} with range: 0..4294967295.
    Capability interface{}

    // Cipher String. The type is string.
    CipherStr interface{}

    // MACSec Desired. The type is bool.
    MacSecDesired interface{}

    // VLAN Outer TAG.
    OuterTag Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag

    // VLAN Inner TAG.
    InnerTag Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag
}

func (sessionSummary *Macsec_Mka_Interfaces_Interface_Session_SessionSummary) GetEntityData() *types.CommonEntityData {
    sessionSummary.EntityData.YFilter = sessionSummary.YFilter
    sessionSummary.EntityData.YangName = "session-summary"
    sessionSummary.EntityData.BundleName = "cisco_ios_xr"
    sessionSummary.EntityData.ParentYangName = "session"
    sessionSummary.EntityData.SegmentPath = "session-summary"
    sessionSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionSummary.EntityData.Children = make(map[string]types.YChild)
    sessionSummary.EntityData.Children["outer-tag"] = types.YChild{"OuterTag", &sessionSummary.OuterTag}
    sessionSummary.EntityData.Children["inner-tag"] = types.YChild{"InnerTag", &sessionSummary.InnerTag}
    sessionSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", sessionSummary.InterfaceName}
    sessionSummary.EntityData.Leafs["inherited-policy"] = types.YLeaf{"InheritedPolicy", sessionSummary.InheritedPolicy}
    sessionSummary.EntityData.Leafs["policy"] = types.YLeaf{"Policy", sessionSummary.Policy}
    sessionSummary.EntityData.Leafs["priority"] = types.YLeaf{"Priority", sessionSummary.Priority}
    sessionSummary.EntityData.Leafs["my-mac"] = types.YLeaf{"MyMac", sessionSummary.MyMac}
    sessionSummary.EntityData.Leafs["delay-protection"] = types.YLeaf{"DelayProtection", sessionSummary.DelayProtection}
    sessionSummary.EntityData.Leafs["replay-protect"] = types.YLeaf{"ReplayProtect", sessionSummary.ReplayProtect}
    sessionSummary.EntityData.Leafs["window-size"] = types.YLeaf{"WindowSize", sessionSummary.WindowSize}
    sessionSummary.EntityData.Leafs["include-icv-indicator"] = types.YLeaf{"IncludeIcvIndicator", sessionSummary.IncludeIcvIndicator}
    sessionSummary.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", sessionSummary.ConfidentialityOffset}
    sessionSummary.EntityData.Leafs["algo-agility"] = types.YLeaf{"AlgoAgility", sessionSummary.AlgoAgility}
    sessionSummary.EntityData.Leafs["capability"] = types.YLeaf{"Capability", sessionSummary.Capability}
    sessionSummary.EntityData.Leafs["cipher-str"] = types.YLeaf{"CipherStr", sessionSummary.CipherStr}
    sessionSummary.EntityData.Leafs["mac-sec-desired"] = types.YLeaf{"MacSecDesired", sessionSummary.MacSecDesired}
    return &(sessionSummary.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag
// VLAN Outer TAG
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EtherType. The type is interface{} with range: 0..65535.
    EtherType interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Cannonical Format Identifier. The type is interface{} with range: 0..255.
    Cfi interface{}

    // Vlan Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (outerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_OuterTag) GetEntityData() *types.CommonEntityData {
    outerTag.EntityData.YFilter = outerTag.YFilter
    outerTag.EntityData.YangName = "outer-tag"
    outerTag.EntityData.BundleName = "cisco_ios_xr"
    outerTag.EntityData.ParentYangName = "session-summary"
    outerTag.EntityData.SegmentPath = "outer-tag"
    outerTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerTag.EntityData.Children = make(map[string]types.YChild)
    outerTag.EntityData.Leafs = make(map[string]types.YLeaf)
    outerTag.EntityData.Leafs["ether-type"] = types.YLeaf{"EtherType", outerTag.EtherType}
    outerTag.EntityData.Leafs["priority"] = types.YLeaf{"Priority", outerTag.Priority}
    outerTag.EntityData.Leafs["cfi"] = types.YLeaf{"Cfi", outerTag.Cfi}
    outerTag.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", outerTag.VlanId}
    return &(outerTag.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag
// VLAN Inner TAG
type Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // EtherType. The type is interface{} with range: 0..65535.
    EtherType interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Cannonical Format Identifier. The type is interface{} with range: 0..255.
    Cfi interface{}

    // Vlan Id. The type is interface{} with range: 0..65535.
    VlanId interface{}
}

func (innerTag *Macsec_Mka_Interfaces_Interface_Session_SessionSummary_InnerTag) GetEntityData() *types.CommonEntityData {
    innerTag.EntityData.YFilter = innerTag.YFilter
    innerTag.EntityData.YangName = "inner-tag"
    innerTag.EntityData.BundleName = "cisco_ios_xr"
    innerTag.EntityData.ParentYangName = "session-summary"
    innerTag.EntityData.SegmentPath = "inner-tag"
    innerTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerTag.EntityData.Children = make(map[string]types.YChild)
    innerTag.EntityData.Leafs = make(map[string]types.YLeaf)
    innerTag.EntityData.Leafs["ether-type"] = types.YLeaf{"EtherType", innerTag.EtherType}
    innerTag.EntityData.Leafs["priority"] = types.YLeaf{"Priority", innerTag.Priority}
    innerTag.EntityData.Leafs["cfi"] = types.YLeaf{"Cfi", innerTag.Cfi}
    innerTag.EntityData.Leafs["vlan-id"] = types.YLeaf{"VlanId", innerTag.VlanId}
    return &(innerTag.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp
// Virtual Pointer Info
type Macsec_Mka_Interfaces_Interface_Session_Vp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local SCI(MAC). The type is string.
    MySci interface{}

    // Virtual Port ID. The type is interface{} with range: 0..4294967295.
    VirtualPortId interface{}

    // Latest Rx status. The type is bool.
    LatestRx interface{}

    // Latest Tx status. The type is bool.
    LatestTx interface{}

    // Latest SAK AN. The type is interface{} with range: 0..4294967295.
    LatestAn interface{}

    // Latest SAK KI. The type is string.
    LatestKi interface{}

    // Latest SAK KN. The type is interface{} with range: 0..4294967295.
    LatestKn interface{}

    // Old Rx status. The type is bool.
    OldRx interface{}

    // Old Tx status. The type is bool.
    OldTx interface{}

    // Old SAK AN. The type is interface{} with range: 0..4294967295.
    OldAn interface{}

    // Old SAK KI. The type is string.
    OldKi interface{}

    // Old SAK KN. The type is interface{} with range: 0..4294967295.
    OldKn interface{}

    // SAK Transmit Wait Time. The type is interface{} with range: 0..4294967295.
    WaitTime interface{}

    // SAK Retire time. The type is interface{} with range: 0..4294967295.
    RetireTime interface{}

    // SAK Cipher Suite. The type is interface{} with range: 0..4294967295.
    CipherSuite interface{}

    // SSCI of the Local TxSC. The type is interface{} with range: 0..4294967295.
    Ssci interface{}

    // Next SAK Rekey time in Sec. The type is string.
    TimeToSakRekey interface{}

    // Fallback Keepalive. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive.
    FallbackKeepalive []Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetEntityData() *types.CommonEntityData {
    vp.EntityData.YFilter = vp.YFilter
    vp.EntityData.YangName = "vp"
    vp.EntityData.BundleName = "cisco_ios_xr"
    vp.EntityData.ParentYangName = "session"
    vp.EntityData.SegmentPath = "vp"
    vp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vp.EntityData.Children = make(map[string]types.YChild)
    vp.EntityData.Children["fallback-keepalive"] = types.YChild{"FallbackKeepalive", nil}
    for i := range vp.FallbackKeepalive {
        vp.EntityData.Children[types.GetSegmentPath(&vp.FallbackKeepalive[i])] = types.YChild{"FallbackKeepalive", &vp.FallbackKeepalive[i]}
    }
    vp.EntityData.Leafs = make(map[string]types.YLeaf)
    vp.EntityData.Leafs["my-sci"] = types.YLeaf{"MySci", vp.MySci}
    vp.EntityData.Leafs["virtual-port-id"] = types.YLeaf{"VirtualPortId", vp.VirtualPortId}
    vp.EntityData.Leafs["latest-rx"] = types.YLeaf{"LatestRx", vp.LatestRx}
    vp.EntityData.Leafs["latest-tx"] = types.YLeaf{"LatestTx", vp.LatestTx}
    vp.EntityData.Leafs["latest-an"] = types.YLeaf{"LatestAn", vp.LatestAn}
    vp.EntityData.Leafs["latest-ki"] = types.YLeaf{"LatestKi", vp.LatestKi}
    vp.EntityData.Leafs["latest-kn"] = types.YLeaf{"LatestKn", vp.LatestKn}
    vp.EntityData.Leafs["old-rx"] = types.YLeaf{"OldRx", vp.OldRx}
    vp.EntityData.Leafs["old-tx"] = types.YLeaf{"OldTx", vp.OldTx}
    vp.EntityData.Leafs["old-an"] = types.YLeaf{"OldAn", vp.OldAn}
    vp.EntityData.Leafs["old-ki"] = types.YLeaf{"OldKi", vp.OldKi}
    vp.EntityData.Leafs["old-kn"] = types.YLeaf{"OldKn", vp.OldKn}
    vp.EntityData.Leafs["wait-time"] = types.YLeaf{"WaitTime", vp.WaitTime}
    vp.EntityData.Leafs["retire-time"] = types.YLeaf{"RetireTime", vp.RetireTime}
    vp.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", vp.CipherSuite}
    vp.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", vp.Ssci}
    vp.EntityData.Leafs["time-to-sak-rekey"] = types.YLeaf{"TimeToSakRekey", vp.TimeToSakRekey}
    return &(vp.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
// Fallback Keepalive
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CKN. The type is string with length: 0..65.
    Ckn interface{}

    // Member Identifier. The type is string with length: 0..25.
    Mi interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // Peers Status.
    PeersStatus Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus
}

func (fallbackKeepalive *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive) GetEntityData() *types.CommonEntityData {
    fallbackKeepalive.EntityData.YFilter = fallbackKeepalive.YFilter
    fallbackKeepalive.EntityData.YangName = "fallback-keepalive"
    fallbackKeepalive.EntityData.BundleName = "cisco_ios_xr"
    fallbackKeepalive.EntityData.ParentYangName = "vp"
    fallbackKeepalive.EntityData.SegmentPath = "fallback-keepalive"
    fallbackKeepalive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fallbackKeepalive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fallbackKeepalive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fallbackKeepalive.EntityData.Children = make(map[string]types.YChild)
    fallbackKeepalive.EntityData.Children["peers-status"] = types.YChild{"PeersStatus", &fallbackKeepalive.PeersStatus}
    fallbackKeepalive.EntityData.Leafs = make(map[string]types.YLeaf)
    fallbackKeepalive.EntityData.Leafs["ckn"] = types.YLeaf{"Ckn", fallbackKeepalive.Ckn}
    fallbackKeepalive.EntityData.Leafs["mi"] = types.YLeaf{"Mi", fallbackKeepalive.Mi}
    fallbackKeepalive.EntityData.Leafs["mn"] = types.YLeaf{"Mn", fallbackKeepalive.Mn}
    return &(fallbackKeepalive.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus
// Peers Status
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx MKPDU Timestamp. The type is string with length: 0..128.
    TxMkpduTimestamp interface{}

    // Peer Count. The type is interface{} with range: 0..4294967295.
    PeerCount interface{}

    // Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer.
    Peer []Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetEntityData() *types.CommonEntityData {
    peersStatus.EntityData.YFilter = peersStatus.YFilter
    peersStatus.EntityData.YangName = "peers-status"
    peersStatus.EntityData.BundleName = "cisco_ios_xr"
    peersStatus.EntityData.ParentYangName = "fallback-keepalive"
    peersStatus.EntityData.SegmentPath = "peers-status"
    peersStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peersStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peersStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peersStatus.EntityData.Children = make(map[string]types.YChild)
    peersStatus.EntityData.Children["peer"] = types.YChild{"Peer", nil}
    for i := range peersStatus.Peer {
        peersStatus.EntityData.Children[types.GetSegmentPath(&peersStatus.Peer[i])] = types.YChild{"Peer", &peersStatus.Peer[i]}
    }
    peersStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    peersStatus.EntityData.Leafs["tx-mkpdu-timestamp"] = types.YLeaf{"TxMkpduTimestamp", peersStatus.TxMkpduTimestamp}
    peersStatus.EntityData.Leafs["peer-count"] = types.YLeaf{"PeerCount", peersStatus.PeerCount}
    return &(peersStatus.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx SCI. The type is string with length: 0..17.
    Sci interface{}

    // Peer Status Data.
    PeerData Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers-status"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = make(map[string]types.YChild)
    peer.EntityData.Children["peer-data"] = types.YChild{"PeerData", &peer.PeerData}
    peer.EntityData.Leafs = make(map[string]types.YLeaf)
    peer.EntityData.Leafs["sci"] = types.YLeaf{"Sci", peer.Sci}
    return &(peer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData
// Peer Status Data
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member ID. The type is string with length: 0..25.
    Mi interface{}

    // ICV Status. The type is string with length: 0..10.
    IcvStatus interface{}

    // ICV Check Timestamp. The type is string with length: 0..128.
    IcvCheckTimestamp interface{}
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer_PeerData) GetEntityData() *types.CommonEntityData {
    peerData.EntityData.YFilter = peerData.YFilter
    peerData.EntityData.YangName = "peer-data"
    peerData.EntityData.BundleName = "cisco_ios_xr"
    peerData.EntityData.ParentYangName = "peer"
    peerData.EntityData.SegmentPath = "peer-data"
    peerData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerData.EntityData.Children = make(map[string]types.YChild)
    peerData.EntityData.Leafs = make(map[string]types.YLeaf)
    peerData.EntityData.Leafs["mi"] = types.YLeaf{"Mi", peerData.Mi}
    peerData.EntityData.Leafs["icv-status"] = types.YLeaf{"IcvStatus", peerData.IcvStatus}
    peerData.EntityData.Leafs["icv-check-timestamp"] = types.YLeaf{"IcvCheckTimestamp", peerData.IcvCheckTimestamp}
    return &(peerData.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca
// CA List for a Session
type Macsec_Mka_Interfaces_Interface_Session_Ca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is Key Server. The type is bool.
    IsKeyServer interface{}

    // Session Status [Secured/Not Secured]. The type is interface{} with range:
    // 0..4294967295.
    Status interface{}

    // Number of Live Peers. The type is interface{} with range: 0..4294967295.
    NumLivePeers interface{}

    // Is First CA. The type is bool.
    FirstCa interface{}

    // Peer SCI(MAC). The type is string.
    PeerSci interface{}

    // Number of Live Peers responded. The type is interface{} with range:
    // 0..4294967295.
    NumLivePeersResponded interface{}

    // CKN. The type is string.
    Ckn interface{}

    // Member Identifier. The type is string.
    MyMi interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    MyMn interface{}

    // authenticator. The type is bool.
    Authenticator interface{}

    // Status Description. The type is string.
    StatusDescription interface{}

    // CA Authentication Mode :PRIMARY-PSK/FALLBACK-PSK/EAP. The type is string.
    AuthenticationMode interface{}

    // Key Chain name. The type is string.
    KeyChain interface{}

    // Peers Status.
    PeersStatus Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus

    // Live Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer.
    LivePeer []Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer

    // Potential Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer.
    PotentialPeer []Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer

    // Dormant Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer.
    DormantPeer []Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetEntityData() *types.CommonEntityData {
    ca.EntityData.YFilter = ca.YFilter
    ca.EntityData.YangName = "ca"
    ca.EntityData.BundleName = "cisco_ios_xr"
    ca.EntityData.ParentYangName = "session"
    ca.EntityData.SegmentPath = "ca"
    ca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ca.EntityData.Children = make(map[string]types.YChild)
    ca.EntityData.Children["peers-status"] = types.YChild{"PeersStatus", &ca.PeersStatus}
    ca.EntityData.Children["live-peer"] = types.YChild{"LivePeer", nil}
    for i := range ca.LivePeer {
        ca.EntityData.Children[types.GetSegmentPath(&ca.LivePeer[i])] = types.YChild{"LivePeer", &ca.LivePeer[i]}
    }
    ca.EntityData.Children["potential-peer"] = types.YChild{"PotentialPeer", nil}
    for i := range ca.PotentialPeer {
        ca.EntityData.Children[types.GetSegmentPath(&ca.PotentialPeer[i])] = types.YChild{"PotentialPeer", &ca.PotentialPeer[i]}
    }
    ca.EntityData.Children["dormant-peer"] = types.YChild{"DormantPeer", nil}
    for i := range ca.DormantPeer {
        ca.EntityData.Children[types.GetSegmentPath(&ca.DormantPeer[i])] = types.YChild{"DormantPeer", &ca.DormantPeer[i]}
    }
    ca.EntityData.Leafs = make(map[string]types.YLeaf)
    ca.EntityData.Leafs["is-key-server"] = types.YLeaf{"IsKeyServer", ca.IsKeyServer}
    ca.EntityData.Leafs["status"] = types.YLeaf{"Status", ca.Status}
    ca.EntityData.Leafs["num-live-peers"] = types.YLeaf{"NumLivePeers", ca.NumLivePeers}
    ca.EntityData.Leafs["first-ca"] = types.YLeaf{"FirstCa", ca.FirstCa}
    ca.EntityData.Leafs["peer-sci"] = types.YLeaf{"PeerSci", ca.PeerSci}
    ca.EntityData.Leafs["num-live-peers-responded"] = types.YLeaf{"NumLivePeersResponded", ca.NumLivePeersResponded}
    ca.EntityData.Leafs["ckn"] = types.YLeaf{"Ckn", ca.Ckn}
    ca.EntityData.Leafs["my-mi"] = types.YLeaf{"MyMi", ca.MyMi}
    ca.EntityData.Leafs["my-mn"] = types.YLeaf{"MyMn", ca.MyMn}
    ca.EntityData.Leafs["authenticator"] = types.YLeaf{"Authenticator", ca.Authenticator}
    ca.EntityData.Leafs["status-description"] = types.YLeaf{"StatusDescription", ca.StatusDescription}
    ca.EntityData.Leafs["authentication-mode"] = types.YLeaf{"AuthenticationMode", ca.AuthenticationMode}
    ca.EntityData.Leafs["key-chain"] = types.YLeaf{"KeyChain", ca.KeyChain}
    return &(ca.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus
// Peers Status
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tx MKPDU Timestamp. The type is string with length: 0..128.
    TxMkpduTimestamp interface{}

    // Peer Count. The type is interface{} with range: 0..4294967295.
    PeerCount interface{}

    // Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer.
    Peer []Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetEntityData() *types.CommonEntityData {
    peersStatus.EntityData.YFilter = peersStatus.YFilter
    peersStatus.EntityData.YangName = "peers-status"
    peersStatus.EntityData.BundleName = "cisco_ios_xr"
    peersStatus.EntityData.ParentYangName = "ca"
    peersStatus.EntityData.SegmentPath = "peers-status"
    peersStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peersStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peersStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peersStatus.EntityData.Children = make(map[string]types.YChild)
    peersStatus.EntityData.Children["peer"] = types.YChild{"Peer", nil}
    for i := range peersStatus.Peer {
        peersStatus.EntityData.Children[types.GetSegmentPath(&peersStatus.Peer[i])] = types.YChild{"Peer", &peersStatus.Peer[i]}
    }
    peersStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    peersStatus.EntityData.Leafs["tx-mkpdu-timestamp"] = types.YLeaf{"TxMkpduTimestamp", peersStatus.TxMkpduTimestamp}
    peersStatus.EntityData.Leafs["peer-count"] = types.YLeaf{"PeerCount", peersStatus.PeerCount}
    return &(peersStatus.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rx SCI. The type is string with length: 0..17.
    Sci interface{}

    // Peer Status Data.
    PeerData Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData
}

func (peer *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers-status"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = make(map[string]types.YChild)
    peer.EntityData.Children["peer-data"] = types.YChild{"PeerData", &peer.PeerData}
    peer.EntityData.Leafs = make(map[string]types.YLeaf)
    peer.EntityData.Leafs["sci"] = types.YLeaf{"Sci", peer.Sci}
    return &(peer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData
// Peer Status Data
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member ID. The type is string with length: 0..25.
    Mi interface{}

    // ICV Status. The type is string with length: 0..10.
    IcvStatus interface{}

    // ICV Check Timestamp. The type is string with length: 0..128.
    IcvCheckTimestamp interface{}
}

func (peerData *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer_PeerData) GetEntityData() *types.CommonEntityData {
    peerData.EntityData.YFilter = peerData.YFilter
    peerData.EntityData.YangName = "peer-data"
    peerData.EntityData.BundleName = "cisco_ios_xr"
    peerData.EntityData.ParentYangName = "peer"
    peerData.EntityData.SegmentPath = "peer-data"
    peerData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerData.EntityData.Children = make(map[string]types.YChild)
    peerData.EntityData.Leafs = make(map[string]types.YLeaf)
    peerData.EntityData.Leafs["mi"] = types.YLeaf{"Mi", peerData.Mi}
    peerData.EntityData.Leafs["icv-status"] = types.YLeaf{"IcvStatus", peerData.IcvStatus}
    peerData.EntityData.Leafs["icv-check-timestamp"] = types.YLeaf{"IcvCheckTimestamp", peerData.IcvCheckTimestamp}
    return &(peerData.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer
// Live Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (livePeer *Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer) GetEntityData() *types.CommonEntityData {
    livePeer.EntityData.YFilter = livePeer.YFilter
    livePeer.EntityData.YangName = "live-peer"
    livePeer.EntityData.BundleName = "cisco_ios_xr"
    livePeer.EntityData.ParentYangName = "ca"
    livePeer.EntityData.SegmentPath = "live-peer"
    livePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    livePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    livePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    livePeer.EntityData.Children = make(map[string]types.YChild)
    livePeer.EntityData.Leafs = make(map[string]types.YLeaf)
    livePeer.EntityData.Leafs["mi"] = types.YLeaf{"Mi", livePeer.Mi}
    livePeer.EntityData.Leafs["sci"] = types.YLeaf{"Sci", livePeer.Sci}
    livePeer.EntityData.Leafs["mn"] = types.YLeaf{"Mn", livePeer.Mn}
    livePeer.EntityData.Leafs["priority"] = types.YLeaf{"Priority", livePeer.Priority}
    livePeer.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", livePeer.Ssci}
    return &(livePeer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer
// Potential Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (potentialPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer) GetEntityData() *types.CommonEntityData {
    potentialPeer.EntityData.YFilter = potentialPeer.YFilter
    potentialPeer.EntityData.YangName = "potential-peer"
    potentialPeer.EntityData.BundleName = "cisco_ios_xr"
    potentialPeer.EntityData.ParentYangName = "ca"
    potentialPeer.EntityData.SegmentPath = "potential-peer"
    potentialPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    potentialPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    potentialPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    potentialPeer.EntityData.Children = make(map[string]types.YChild)
    potentialPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    potentialPeer.EntityData.Leafs["mi"] = types.YLeaf{"Mi", potentialPeer.Mi}
    potentialPeer.EntityData.Leafs["sci"] = types.YLeaf{"Sci", potentialPeer.Sci}
    potentialPeer.EntityData.Leafs["mn"] = types.YLeaf{"Mn", potentialPeer.Mn}
    potentialPeer.EntityData.Leafs["priority"] = types.YLeaf{"Priority", potentialPeer.Priority}
    potentialPeer.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", potentialPeer.Ssci}
    return &(potentialPeer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
// Dormant Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member ID. The type is string.
    Mi interface{}

    // Rx SCI. The type is string.
    Sci interface{}

    // Message Number. The type is interface{} with range: 0..4294967295.
    Mn interface{}

    // KS Priority. The type is interface{} with range: 0..4294967295.
    Priority interface{}

    // Peer SSCI. The type is interface{} with range: 0..4294967295.
    Ssci interface{}
}

func (dormantPeer *Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer) GetEntityData() *types.CommonEntityData {
    dormantPeer.EntityData.YFilter = dormantPeer.YFilter
    dormantPeer.EntityData.YangName = "dormant-peer"
    dormantPeer.EntityData.BundleName = "cisco_ios_xr"
    dormantPeer.EntityData.ParentYangName = "ca"
    dormantPeer.EntityData.SegmentPath = "dormant-peer"
    dormantPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dormantPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dormantPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dormantPeer.EntityData.Children = make(map[string]types.YChild)
    dormantPeer.EntityData.Leafs = make(map[string]types.YLeaf)
    dormantPeer.EntityData.Leafs["mi"] = types.YLeaf{"Mi", dormantPeer.Mi}
    dormantPeer.EntityData.Leafs["sci"] = types.YLeaf{"Sci", dormantPeer.Sci}
    dormantPeer.EntityData.Leafs["mn"] = types.YLeaf{"Mn", dormantPeer.Mn}
    dormantPeer.EntityData.Leafs["priority"] = types.YLeaf{"Priority", dormantPeer.Priority}
    dormantPeer.EntityData.Leafs["ssci"] = types.YLeaf{"Ssci", dormantPeer.Ssci}
    return &(dormantPeer.EntityData)
}

