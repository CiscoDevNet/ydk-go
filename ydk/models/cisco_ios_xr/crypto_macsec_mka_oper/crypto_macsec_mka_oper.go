// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-macsec-mka package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec: Macsec operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// MacsecCipherSuite represents Macsec cipher suite
type MacsecCipherSuite string

const (
    // Invalid MACsec cipher
    MacsecCipherSuite_cipher_suite_none MacsecCipherSuite = "cipher-suite-none"

    // 128 bit GCM_AES MACsec cipher suite
    MacsecCipherSuite_cipher_suite_gcm_aes_128 MacsecCipherSuite = "cipher-suite-gcm-aes-128"

    // 256 bit GCM_AES MACsec cipher suite
    MacsecCipherSuite_cipher_suite_gcm_aes_256 MacsecCipherSuite = "cipher-suite-gcm-aes-256"

    // 128 bit GCM_AES MACsec XPN cipher suite
    MacsecCipherSuite_cipher_suite_gcm_aes_128_xpn MacsecCipherSuite = "cipher-suite-gcm-aes-128-xpn"

    // 256 bit GCM_AES MACsec XPN cipher suite
    MacsecCipherSuite_cipher_suite_gcm_aes_256_xpn MacsecCipherSuite = "cipher-suite-gcm-aes-256-xpn"
)

// MkaAuthenticationMode represents Mka authentication mode
type MkaAuthenticationMode string

const (
    // Invalid authentication mode
    MkaAuthenticationMode_auth_mode_invalid MkaAuthenticationMode = "auth-mode-invalid"

    // Preshared Key
    MkaAuthenticationMode_auth_mode_psk MkaAuthenticationMode = "auth-mode-psk"

    // EAP
    MkaAuthenticationMode_auth_mode_eap MkaAuthenticationMode = "auth-mode-eap"
)

// MacsecServicePort represents Macsec service port
type MacsecServicePort string

const (
    // Macsec Service not enabled
    MacsecServicePort_macsec_service_port_none MacsecServicePort = "macsec-service-port-none"

    // Macsec Service Encryption Port
    MacsecServicePort_macsec_service_port_encryption MacsecServicePort = "macsec-service-port-encryption"

    // Macsec Service Decryption Port
    MacsecServicePort_macsec_service_port_decryption MacsecServicePort = "macsec-service-port-decryption"
)

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
    macsec.EntityData.AbsolutePath = macsec.EntityData.SegmentPath
    macsec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macsec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macsec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macsec.EntityData.Children = types.NewOrderedMap()
    macsec.EntityData.Children.Append("mka", types.YChild{"Mka", &macsec.Mka})
    macsec.EntityData.Leafs = types.NewOrderedMap()

    macsec.EntityData.YListKeys = []string {}

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
    mka.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/" + mka.EntityData.SegmentPath
    mka.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mka.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mka.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mka.EntityData.Children = types.NewOrderedMap()
    mka.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &mka.Interfaces})
    mka.EntityData.Leafs = types.NewOrderedMap()

    mka.EntityData.YListKeys = []string {}

    return &(mka.EntityData)
}

// Macsec_Mka_Interfaces
// MKA Data
type Macsec_Mka_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MKA Data for the Interface. The type is slice of
    // Macsec_Mka_Interfaces_Interface.
    Interface []*Macsec_Mka_Interfaces_Interface
}

func (interfaces *Macsec_Mka_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "mka"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Macsec_Mka_Interfaces_Interface
// MKA Data for the Interface
type Macsec_Mka_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Name interface{}

    // MKA Session Data.
    Session Macsec_Mka_Interfaces_Interface_Session

    // MKA Interface Summary Data.
    Info Macsec_Mka_Interfaces_Interface_Info
}

func (self *Macsec_Mka_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Name, "name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("session", types.YChild{"Session", &self.Session})
    self.EntityData.Children.Append("info", types.YChild{"Info", &self.Info})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("name", types.YLeaf{"Name", self.Name})

    self.EntityData.YListKeys = []string {"Name"}

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
    Ca []*Macsec_Mka_Interfaces_Interface_Session_Ca
}

func (session *Macsec_Mka_Interfaces_Interface_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "interface"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("session-summary", types.YChild{"SessionSummary", &session.SessionSummary})
    session.EntityData.Children.Append("vp", types.YChild{"Vp", &session.Vp})
    session.EntityData.Children.Append("ca", types.YChild{"Ca", nil})
    for i := range session.Ca {
        types.SetYListKey(session.Ca[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.Ca[i]), types.YChild{"Ca", session.Ca[i]})
    }
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

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

    // MKA Cipher Suite. The type is string.
    MkaCipherSuite interface{}

    // configured cipher suite. The type is string.
    ConfiguredMacSecCipherSuite interface{}

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
    sessionSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/" + sessionSummary.EntityData.SegmentPath
    sessionSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionSummary.EntityData.Children = types.NewOrderedMap()
    sessionSummary.EntityData.Children.Append("outer-tag", types.YChild{"OuterTag", &sessionSummary.OuterTag})
    sessionSummary.EntityData.Children.Append("inner-tag", types.YChild{"InnerTag", &sessionSummary.InnerTag})
    sessionSummary.EntityData.Leafs = types.NewOrderedMap()
    sessionSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sessionSummary.InterfaceName})
    sessionSummary.EntityData.Leafs.Append("inherited-policy", types.YLeaf{"InheritedPolicy", sessionSummary.InheritedPolicy})
    sessionSummary.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", sessionSummary.Policy})
    sessionSummary.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", sessionSummary.Priority})
    sessionSummary.EntityData.Leafs.Append("my-mac", types.YLeaf{"MyMac", sessionSummary.MyMac})
    sessionSummary.EntityData.Leafs.Append("delay-protection", types.YLeaf{"DelayProtection", sessionSummary.DelayProtection})
    sessionSummary.EntityData.Leafs.Append("replay-protect", types.YLeaf{"ReplayProtect", sessionSummary.ReplayProtect})
    sessionSummary.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", sessionSummary.WindowSize})
    sessionSummary.EntityData.Leafs.Append("include-icv-indicator", types.YLeaf{"IncludeIcvIndicator", sessionSummary.IncludeIcvIndicator})
    sessionSummary.EntityData.Leafs.Append("confidentiality-offset", types.YLeaf{"ConfidentialityOffset", sessionSummary.ConfidentialityOffset})
    sessionSummary.EntityData.Leafs.Append("algo-agility", types.YLeaf{"AlgoAgility", sessionSummary.AlgoAgility})
    sessionSummary.EntityData.Leafs.Append("capability", types.YLeaf{"Capability", sessionSummary.Capability})
    sessionSummary.EntityData.Leafs.Append("mka-cipher-suite", types.YLeaf{"MkaCipherSuite", sessionSummary.MkaCipherSuite})
    sessionSummary.EntityData.Leafs.Append("configured-mac-sec-cipher-suite", types.YLeaf{"ConfiguredMacSecCipherSuite", sessionSummary.ConfiguredMacSecCipherSuite})
    sessionSummary.EntityData.Leafs.Append("mac-sec-desired", types.YLeaf{"MacSecDesired", sessionSummary.MacSecDesired})

    sessionSummary.EntityData.YListKeys = []string {}

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
    outerTag.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/session-summary/" + outerTag.EntityData.SegmentPath
    outerTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outerTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outerTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outerTag.EntityData.Children = types.NewOrderedMap()
    outerTag.EntityData.Leafs = types.NewOrderedMap()
    outerTag.EntityData.Leafs.Append("ether-type", types.YLeaf{"EtherType", outerTag.EtherType})
    outerTag.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", outerTag.Priority})
    outerTag.EntityData.Leafs.Append("cfi", types.YLeaf{"Cfi", outerTag.Cfi})
    outerTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", outerTag.VlanId})

    outerTag.EntityData.YListKeys = []string {}

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
    innerTag.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/session-summary/" + innerTag.EntityData.SegmentPath
    innerTag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    innerTag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    innerTag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    innerTag.EntityData.Children = types.NewOrderedMap()
    innerTag.EntityData.Leafs = types.NewOrderedMap()
    innerTag.EntityData.Leafs.Append("ether-type", types.YLeaf{"EtherType", innerTag.EtherType})
    innerTag.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", innerTag.Priority})
    innerTag.EntityData.Leafs.Append("cfi", types.YLeaf{"Cfi", innerTag.Cfi})
    innerTag.EntityData.Leafs.Append("vlan-id", types.YLeaf{"VlanId", innerTag.VlanId})

    innerTag.EntityData.YListKeys = []string {}

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

    // SAK Cipher Suite. The type is MacsecCipherSuite.
    MacsecCipherSuite interface{}

    // SSCI of the Local TxSC. The type is interface{} with range: 0..4294967295.
    Ssci interface{}

    // Next SAK Rekey time in Sec. The type is string.
    TimeToSakRekey interface{}

    // Fallback Keepalive. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive.
    FallbackKeepalive []*Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
}

func (vp *Macsec_Mka_Interfaces_Interface_Session_Vp) GetEntityData() *types.CommonEntityData {
    vp.EntityData.YFilter = vp.YFilter
    vp.EntityData.YangName = "vp"
    vp.EntityData.BundleName = "cisco_ios_xr"
    vp.EntityData.ParentYangName = "session"
    vp.EntityData.SegmentPath = "vp"
    vp.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/" + vp.EntityData.SegmentPath
    vp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vp.EntityData.Children = types.NewOrderedMap()
    vp.EntityData.Children.Append("fallback-keepalive", types.YChild{"FallbackKeepalive", nil})
    for i := range vp.FallbackKeepalive {
        types.SetYListKey(vp.FallbackKeepalive[i], i)
        vp.EntityData.Children.Append(types.GetSegmentPath(vp.FallbackKeepalive[i]), types.YChild{"FallbackKeepalive", vp.FallbackKeepalive[i]})
    }
    vp.EntityData.Leafs = types.NewOrderedMap()
    vp.EntityData.Leafs.Append("my-sci", types.YLeaf{"MySci", vp.MySci})
    vp.EntityData.Leafs.Append("virtual-port-id", types.YLeaf{"VirtualPortId", vp.VirtualPortId})
    vp.EntityData.Leafs.Append("latest-rx", types.YLeaf{"LatestRx", vp.LatestRx})
    vp.EntityData.Leafs.Append("latest-tx", types.YLeaf{"LatestTx", vp.LatestTx})
    vp.EntityData.Leafs.Append("latest-an", types.YLeaf{"LatestAn", vp.LatestAn})
    vp.EntityData.Leafs.Append("latest-ki", types.YLeaf{"LatestKi", vp.LatestKi})
    vp.EntityData.Leafs.Append("latest-kn", types.YLeaf{"LatestKn", vp.LatestKn})
    vp.EntityData.Leafs.Append("old-rx", types.YLeaf{"OldRx", vp.OldRx})
    vp.EntityData.Leafs.Append("old-tx", types.YLeaf{"OldTx", vp.OldTx})
    vp.EntityData.Leafs.Append("old-an", types.YLeaf{"OldAn", vp.OldAn})
    vp.EntityData.Leafs.Append("old-ki", types.YLeaf{"OldKi", vp.OldKi})
    vp.EntityData.Leafs.Append("old-kn", types.YLeaf{"OldKn", vp.OldKn})
    vp.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", vp.WaitTime})
    vp.EntityData.Leafs.Append("retire-time", types.YLeaf{"RetireTime", vp.RetireTime})
    vp.EntityData.Leafs.Append("macsec-cipher-suite", types.YLeaf{"MacsecCipherSuite", vp.MacsecCipherSuite})
    vp.EntityData.Leafs.Append("ssci", types.YLeaf{"Ssci", vp.Ssci})
    vp.EntityData.Leafs.Append("time-to-sak-rekey", types.YLeaf{"TimeToSakRekey", vp.TimeToSakRekey})

    vp.EntityData.YListKeys = []string {}

    return &(vp.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive
// Fallback Keepalive
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    fallbackKeepalive.EntityData.SegmentPath = "fallback-keepalive" + types.AddNoKeyToken(fallbackKeepalive)
    fallbackKeepalive.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/vp/" + fallbackKeepalive.EntityData.SegmentPath
    fallbackKeepalive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fallbackKeepalive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fallbackKeepalive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fallbackKeepalive.EntityData.Children = types.NewOrderedMap()
    fallbackKeepalive.EntityData.Children.Append("peers-status", types.YChild{"PeersStatus", &fallbackKeepalive.PeersStatus})
    fallbackKeepalive.EntityData.Leafs = types.NewOrderedMap()
    fallbackKeepalive.EntityData.Leafs.Append("ckn", types.YLeaf{"Ckn", fallbackKeepalive.Ckn})
    fallbackKeepalive.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", fallbackKeepalive.Mi})
    fallbackKeepalive.EntityData.Leafs.Append("mn", types.YLeaf{"Mn", fallbackKeepalive.Mn})

    fallbackKeepalive.EntityData.YListKeys = []string {}

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
    Peer []*Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus) GetEntityData() *types.CommonEntityData {
    peersStatus.EntityData.YFilter = peersStatus.YFilter
    peersStatus.EntityData.YangName = "peers-status"
    peersStatus.EntityData.BundleName = "cisco_ios_xr"
    peersStatus.EntityData.ParentYangName = "fallback-keepalive"
    peersStatus.EntityData.SegmentPath = "peers-status"
    peersStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/vp/fallback-keepalive/" + peersStatus.EntityData.SegmentPath
    peersStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peersStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peersStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peersStatus.EntityData.Children = types.NewOrderedMap()
    peersStatus.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peersStatus.Peer {
        types.SetYListKey(peersStatus.Peer[i], i)
        peersStatus.EntityData.Children.Append(types.GetSegmentPath(peersStatus.Peer[i]), types.YChild{"Peer", peersStatus.Peer[i]})
    }
    peersStatus.EntityData.Leafs = types.NewOrderedMap()
    peersStatus.EntityData.Leafs.Append("tx-mkpdu-timestamp", types.YLeaf{"TxMkpduTimestamp", peersStatus.TxMkpduTimestamp})
    peersStatus.EntityData.Leafs.Append("peer-count", types.YLeaf{"PeerCount", peersStatus.PeerCount})

    peersStatus.EntityData.YListKeys = []string {}

    return &(peersStatus.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Vp_FallbackKeepalive_PeersStatus_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    peer.EntityData.SegmentPath = "peer" + types.AddNoKeyToken(peer)
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/vp/fallback-keepalive/peers-status/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("peer-data", types.YChild{"PeerData", &peer.PeerData})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("sci", types.YLeaf{"Sci", peer.Sci})

    peer.EntityData.YListKeys = []string {}

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
    peerData.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/vp/fallback-keepalive/peers-status/peer/" + peerData.EntityData.SegmentPath
    peerData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerData.EntityData.Children = types.NewOrderedMap()
    peerData.EntityData.Leafs = types.NewOrderedMap()
    peerData.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", peerData.Mi})
    peerData.EntityData.Leafs.Append("icv-status", types.YLeaf{"IcvStatus", peerData.IcvStatus})
    peerData.EntityData.Leafs.Append("icv-check-timestamp", types.YLeaf{"IcvCheckTimestamp", peerData.IcvCheckTimestamp})

    peerData.EntityData.YListKeys = []string {}

    return &(peerData.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca
// CA List for a Session
type Macsec_Mka_Interfaces_Interface_Session_Ca struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    LivePeer []*Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer

    // Potential Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer.
    PotentialPeer []*Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer

    // Dormant Peer List. The type is slice of
    // Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer.
    DormantPeer []*Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
}

func (ca *Macsec_Mka_Interfaces_Interface_Session_Ca) GetEntityData() *types.CommonEntityData {
    ca.EntityData.YFilter = ca.YFilter
    ca.EntityData.YangName = "ca"
    ca.EntityData.BundleName = "cisco_ios_xr"
    ca.EntityData.ParentYangName = "session"
    ca.EntityData.SegmentPath = "ca" + types.AddNoKeyToken(ca)
    ca.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/" + ca.EntityData.SegmentPath
    ca.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ca.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ca.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ca.EntityData.Children = types.NewOrderedMap()
    ca.EntityData.Children.Append("peers-status", types.YChild{"PeersStatus", &ca.PeersStatus})
    ca.EntityData.Children.Append("live-peer", types.YChild{"LivePeer", nil})
    for i := range ca.LivePeer {
        types.SetYListKey(ca.LivePeer[i], i)
        ca.EntityData.Children.Append(types.GetSegmentPath(ca.LivePeer[i]), types.YChild{"LivePeer", ca.LivePeer[i]})
    }
    ca.EntityData.Children.Append("potential-peer", types.YChild{"PotentialPeer", nil})
    for i := range ca.PotentialPeer {
        types.SetYListKey(ca.PotentialPeer[i], i)
        ca.EntityData.Children.Append(types.GetSegmentPath(ca.PotentialPeer[i]), types.YChild{"PotentialPeer", ca.PotentialPeer[i]})
    }
    ca.EntityData.Children.Append("dormant-peer", types.YChild{"DormantPeer", nil})
    for i := range ca.DormantPeer {
        types.SetYListKey(ca.DormantPeer[i], i)
        ca.EntityData.Children.Append(types.GetSegmentPath(ca.DormantPeer[i]), types.YChild{"DormantPeer", ca.DormantPeer[i]})
    }
    ca.EntityData.Leafs = types.NewOrderedMap()
    ca.EntityData.Leafs.Append("is-key-server", types.YLeaf{"IsKeyServer", ca.IsKeyServer})
    ca.EntityData.Leafs.Append("status", types.YLeaf{"Status", ca.Status})
    ca.EntityData.Leafs.Append("num-live-peers", types.YLeaf{"NumLivePeers", ca.NumLivePeers})
    ca.EntityData.Leafs.Append("first-ca", types.YLeaf{"FirstCa", ca.FirstCa})
    ca.EntityData.Leafs.Append("peer-sci", types.YLeaf{"PeerSci", ca.PeerSci})
    ca.EntityData.Leafs.Append("num-live-peers-responded", types.YLeaf{"NumLivePeersResponded", ca.NumLivePeersResponded})
    ca.EntityData.Leafs.Append("ckn", types.YLeaf{"Ckn", ca.Ckn})
    ca.EntityData.Leafs.Append("my-mi", types.YLeaf{"MyMi", ca.MyMi})
    ca.EntityData.Leafs.Append("my-mn", types.YLeaf{"MyMn", ca.MyMn})
    ca.EntityData.Leafs.Append("authenticator", types.YLeaf{"Authenticator", ca.Authenticator})
    ca.EntityData.Leafs.Append("status-description", types.YLeaf{"StatusDescription", ca.StatusDescription})
    ca.EntityData.Leafs.Append("authentication-mode", types.YLeaf{"AuthenticationMode", ca.AuthenticationMode})
    ca.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", ca.KeyChain})

    ca.EntityData.YListKeys = []string {}

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
    Peer []*Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
}

func (peersStatus *Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus) GetEntityData() *types.CommonEntityData {
    peersStatus.EntityData.YFilter = peersStatus.YFilter
    peersStatus.EntityData.YangName = "peers-status"
    peersStatus.EntityData.BundleName = "cisco_ios_xr"
    peersStatus.EntityData.ParentYangName = "ca"
    peersStatus.EntityData.SegmentPath = "peers-status"
    peersStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/" + peersStatus.EntityData.SegmentPath
    peersStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peersStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peersStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peersStatus.EntityData.Children = types.NewOrderedMap()
    peersStatus.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peersStatus.Peer {
        types.SetYListKey(peersStatus.Peer[i], i)
        peersStatus.EntityData.Children.Append(types.GetSegmentPath(peersStatus.Peer[i]), types.YChild{"Peer", peersStatus.Peer[i]})
    }
    peersStatus.EntityData.Leafs = types.NewOrderedMap()
    peersStatus.EntityData.Leafs.Append("tx-mkpdu-timestamp", types.YLeaf{"TxMkpduTimestamp", peersStatus.TxMkpduTimestamp})
    peersStatus.EntityData.Leafs.Append("peer-count", types.YLeaf{"PeerCount", peersStatus.PeerCount})

    peersStatus.EntityData.YListKeys = []string {}

    return &(peersStatus.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer
// Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PeersStatus_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    peer.EntityData.SegmentPath = "peer" + types.AddNoKeyToken(peer)
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/peers-status/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("peer-data", types.YChild{"PeerData", &peer.PeerData})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("sci", types.YLeaf{"Sci", peer.Sci})

    peer.EntityData.YListKeys = []string {}

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
    peerData.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/peers-status/peer/" + peerData.EntityData.SegmentPath
    peerData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerData.EntityData.Children = types.NewOrderedMap()
    peerData.EntityData.Leafs = types.NewOrderedMap()
    peerData.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", peerData.Mi})
    peerData.EntityData.Leafs.Append("icv-status", types.YLeaf{"IcvStatus", peerData.IcvStatus})
    peerData.EntityData.Leafs.Append("icv-check-timestamp", types.YLeaf{"IcvCheckTimestamp", peerData.IcvCheckTimestamp})

    peerData.EntityData.YListKeys = []string {}

    return &(peerData.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer
// Live Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_LivePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    livePeer.EntityData.SegmentPath = "live-peer" + types.AddNoKeyToken(livePeer)
    livePeer.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/" + livePeer.EntityData.SegmentPath
    livePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    livePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    livePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    livePeer.EntityData.Children = types.NewOrderedMap()
    livePeer.EntityData.Leafs = types.NewOrderedMap()
    livePeer.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", livePeer.Mi})
    livePeer.EntityData.Leafs.Append("sci", types.YLeaf{"Sci", livePeer.Sci})
    livePeer.EntityData.Leafs.Append("mn", types.YLeaf{"Mn", livePeer.Mn})
    livePeer.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", livePeer.Priority})
    livePeer.EntityData.Leafs.Append("ssci", types.YLeaf{"Ssci", livePeer.Ssci})

    livePeer.EntityData.YListKeys = []string {}

    return &(livePeer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer
// Potential Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_PotentialPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    potentialPeer.EntityData.SegmentPath = "potential-peer" + types.AddNoKeyToken(potentialPeer)
    potentialPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/" + potentialPeer.EntityData.SegmentPath
    potentialPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    potentialPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    potentialPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    potentialPeer.EntityData.Children = types.NewOrderedMap()
    potentialPeer.EntityData.Leafs = types.NewOrderedMap()
    potentialPeer.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", potentialPeer.Mi})
    potentialPeer.EntityData.Leafs.Append("sci", types.YLeaf{"Sci", potentialPeer.Sci})
    potentialPeer.EntityData.Leafs.Append("mn", types.YLeaf{"Mn", potentialPeer.Mn})
    potentialPeer.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", potentialPeer.Priority})
    potentialPeer.EntityData.Leafs.Append("ssci", types.YLeaf{"Ssci", potentialPeer.Ssci})

    potentialPeer.EntityData.YListKeys = []string {}

    return &(potentialPeer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer
// Dormant Peer List
type Macsec_Mka_Interfaces_Interface_Session_Ca_DormantPeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    dormantPeer.EntityData.SegmentPath = "dormant-peer" + types.AddNoKeyToken(dormantPeer)
    dormantPeer.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/session/ca/" + dormantPeer.EntityData.SegmentPath
    dormantPeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dormantPeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dormantPeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dormantPeer.EntityData.Children = types.NewOrderedMap()
    dormantPeer.EntityData.Leafs = types.NewOrderedMap()
    dormantPeer.EntityData.Leafs.Append("mi", types.YLeaf{"Mi", dormantPeer.Mi})
    dormantPeer.EntityData.Leafs.Append("sci", types.YLeaf{"Sci", dormantPeer.Sci})
    dormantPeer.EntityData.Leafs.Append("mn", types.YLeaf{"Mn", dormantPeer.Mn})
    dormantPeer.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", dormantPeer.Priority})
    dormantPeer.EntityData.Leafs.Append("ssci", types.YLeaf{"Ssci", dormantPeer.Ssci})

    dormantPeer.EntityData.YListKeys = []string {}

    return &(dormantPeer.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Info
// MKA Interface Summary Data
type Macsec_Mka_Interfaces_Interface_Info struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MKA Interface Summary.
    InterfaceSummary Macsec_Mka_Interfaces_Interface_Info_InterfaceSummary
}

func (info *Macsec_Mka_Interfaces_Interface_Info) GetEntityData() *types.CommonEntityData {
    info.EntityData.YFilter = info.YFilter
    info.EntityData.YangName = "info"
    info.EntityData.BundleName = "cisco_ios_xr"
    info.EntityData.ParentYangName = "interface"
    info.EntityData.SegmentPath = "info"
    info.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/" + info.EntityData.SegmentPath
    info.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    info.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    info.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    info.EntityData.Children = types.NewOrderedMap()
    info.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", &info.InterfaceSummary})
    info.EntityData.Leafs = types.NewOrderedMap()

    info.EntityData.YListKeys = []string {}

    return &(info.EntityData)
}

// Macsec_Mka_Interfaces_Interface_Info_InterfaceSummary
// MKA Interface Summary
type Macsec_Mka_Interfaces_Interface_Info_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // macsec configured interface. The type is string.
    InterfaceName interface{}

    // Short Name String. The type is string.
    ShortName interface{}

    // Name  of the Key chain. The type is string.
    KeyChain interface{}

    // Policy name. The type is string.
    Policy interface{}

    // Is macsec-service port or not. The type is bool.
    MacsecSvcPort interface{}

    // Macsec-service Encryption / Decryption port. The type is MacsecServicePort.
    MacsecSvcPortType interface{}

    // Macsec Service paired port Short Name String. The type is string.
    SvcportShortName interface{}

    // MKA authentication mode. The type is MkaAuthenticationMode.
    MkaMode interface{}

    // fallback Keychain name. The type is string.
    FallbackKeychain interface{}

    // MacsecShutdown. The type is bool.
    MacsecShutdown interface{}
}

func (interfaceSummary *Macsec_Mka_Interfaces_Interface_Info_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "info"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.AbsolutePath = "Cisco-IOS-XR-crypto-macsec-mka-oper:macsec/mka/interfaces/interface/info/" + interfaceSummary.EntityData.SegmentPath
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName})
    interfaceSummary.EntityData.Leafs.Append("short-name", types.YLeaf{"ShortName", interfaceSummary.ShortName})
    interfaceSummary.EntityData.Leafs.Append("key-chain", types.YLeaf{"KeyChain", interfaceSummary.KeyChain})
    interfaceSummary.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", interfaceSummary.Policy})
    interfaceSummary.EntityData.Leafs.Append("macsec-svc-port", types.YLeaf{"MacsecSvcPort", interfaceSummary.MacsecSvcPort})
    interfaceSummary.EntityData.Leafs.Append("macsec-svc-port-type", types.YLeaf{"MacsecSvcPortType", interfaceSummary.MacsecSvcPortType})
    interfaceSummary.EntityData.Leafs.Append("svcport-short-name", types.YLeaf{"SvcportShortName", interfaceSummary.SvcportShortName})
    interfaceSummary.EntityData.Leafs.Append("mka-mode", types.YLeaf{"MkaMode", interfaceSummary.MkaMode})
    interfaceSummary.EntityData.Leafs.Append("fallback-keychain", types.YLeaf{"FallbackKeychain", interfaceSummary.FallbackKeychain})
    interfaceSummary.EntityData.Leafs.Append("macsec-shutdown", types.YLeaf{"MacsecShutdown", interfaceSummary.MacsecShutdown})

    interfaceSummary.EntityData.YListKeys = []string {}

    return &(interfaceSummary.EntityData)
}

