// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-macsec-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ncs1k-macsec-oper: Macsec data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs1k_macsec_ea_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs1k_macsec_ea_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-macsec-ea-oper ncs1k-macsec-oper}", reflect.TypeOf(Ncs1KMacsecOper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-macsec-ea-oper:ncs1k-macsec-oper", reflect.TypeOf(Ncs1KMacsecOper{}))
}

// Ncs1kCipherSuit represents Ncs1k cipher suit
type Ncs1kCipherSuit string

const (
    // GCM AES 256
    Ncs1kCipherSuit_gcm_aes_256 Ncs1kCipherSuit = "gcm-aes-256"

    // GCM AES 128
    Ncs1kCipherSuit_gcm_aes_128 Ncs1kCipherSuit = "gcm-aes-128"

    // GCM AES XPN 256
    Ncs1kCipherSuit_gcm_aes_xpn_256 Ncs1kCipherSuit = "gcm-aes-xpn-256"
)

// Ncs1KMacsecOper
// Macsec data
type Ncs1KMacsecOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All Macsec operational data.
    Ncs1KMacsecCtrlrNames Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetEntityData() *types.CommonEntityData {
    ncs1KMacsecOper.EntityData.YFilter = ncs1KMacsecOper.YFilter
    ncs1KMacsecOper.EntityData.YangName = "ncs1k-macsec-oper"
    ncs1KMacsecOper.EntityData.BundleName = "cisco_ios_xr"
    ncs1KMacsecOper.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-macsec-ea-oper"
    ncs1KMacsecOper.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-macsec-ea-oper:ncs1k-macsec-oper"
    ncs1KMacsecOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1KMacsecOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1KMacsecOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1KMacsecOper.EntityData.Children = make(map[string]types.YChild)
    ncs1KMacsecOper.EntityData.Children["ncs1k-macsec-ctrlr-names"] = types.YChild{"Ncs1KMacsecCtrlrNames", &ncs1KMacsecOper.Ncs1KMacsecCtrlrNames}
    ncs1KMacsecOper.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ncs1KMacsecOper.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames
// All Macsec operational data
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName.
    Ncs1KMacsecCtrlrName []Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetEntityData() *types.CommonEntityData {
    ncs1KMacsecCtrlrNames.EntityData.YFilter = ncs1KMacsecCtrlrNames.YFilter
    ncs1KMacsecCtrlrNames.EntityData.YangName = "ncs1k-macsec-ctrlr-names"
    ncs1KMacsecCtrlrNames.EntityData.BundleName = "cisco_ios_xr"
    ncs1KMacsecCtrlrNames.EntityData.ParentYangName = "ncs1k-macsec-oper"
    ncs1KMacsecCtrlrNames.EntityData.SegmentPath = "ncs1k-macsec-ctrlr-names"
    ncs1KMacsecCtrlrNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1KMacsecCtrlrNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1KMacsecCtrlrNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1KMacsecCtrlrNames.EntityData.Children = make(map[string]types.YChild)
    ncs1KMacsecCtrlrNames.EntityData.Children["ncs1k-macsec-ctrlr-name"] = types.YChild{"Ncs1KMacsecCtrlrName", nil}
    for i := range ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName {
        ncs1KMacsecCtrlrNames.EntityData.Children[types.GetSegmentPath(&ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName[i])] = types.YChild{"Ncs1KMacsecCtrlrName", &ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName[i]}
    }
    ncs1KMacsecCtrlrNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ncs1KMacsecCtrlrNames.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName
// Interface name
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    Name interface{}

    // controller data.
    Ncs1KStatusInfo Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetEntityData() *types.CommonEntityData {
    ncs1KMacsecCtrlrName.EntityData.YFilter = ncs1KMacsecCtrlrName.YFilter
    ncs1KMacsecCtrlrName.EntityData.YangName = "ncs1k-macsec-ctrlr-name"
    ncs1KMacsecCtrlrName.EntityData.BundleName = "cisco_ios_xr"
    ncs1KMacsecCtrlrName.EntityData.ParentYangName = "ncs1k-macsec-ctrlr-names"
    ncs1KMacsecCtrlrName.EntityData.SegmentPath = "ncs1k-macsec-ctrlr-name" + "[name='" + fmt.Sprintf("%v", ncs1KMacsecCtrlrName.Name) + "']"
    ncs1KMacsecCtrlrName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1KMacsecCtrlrName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1KMacsecCtrlrName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1KMacsecCtrlrName.EntityData.Children = make(map[string]types.YChild)
    ncs1KMacsecCtrlrName.EntityData.Children["ncs1k-status-info"] = types.YChild{"Ncs1KStatusInfo", &ncs1KMacsecCtrlrName.Ncs1KStatusInfo}
    ncs1KMacsecCtrlrName.EntityData.Leafs = make(map[string]types.YLeaf)
    ncs1KMacsecCtrlrName.EntityData.Leafs["name"] = types.YLeaf{"Name", ncs1KMacsecCtrlrName.Name}
    return &(ncs1KMacsecCtrlrName.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo
// controller data
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    ReplayWindowSize interface{}

    // Must Secure. The type is bool.
    MustSecure interface{}

    // Secure Mode. The type is interface{} with range: 0..4294967295.
    SecureMode interface{}

    // Encrypt Secure Channel Status.
    EncryptScStatus Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus

    // Decrypt Secure Channel Status.
    DecryptScStatus Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetEntityData() *types.CommonEntityData {
    ncs1KStatusInfo.EntityData.YFilter = ncs1KStatusInfo.YFilter
    ncs1KStatusInfo.EntityData.YangName = "ncs1k-status-info"
    ncs1KStatusInfo.EntityData.BundleName = "cisco_ios_xr"
    ncs1KStatusInfo.EntityData.ParentYangName = "ncs1k-macsec-ctrlr-name"
    ncs1KStatusInfo.EntityData.SegmentPath = "ncs1k-status-info"
    ncs1KStatusInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1KStatusInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1KStatusInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1KStatusInfo.EntityData.Children = make(map[string]types.YChild)
    ncs1KStatusInfo.EntityData.Children["encrypt-sc-status"] = types.YChild{"EncryptScStatus", &ncs1KStatusInfo.EncryptScStatus}
    ncs1KStatusInfo.EntityData.Children["decrypt-sc-status"] = types.YChild{"DecryptScStatus", &ncs1KStatusInfo.DecryptScStatus}
    ncs1KStatusInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    ncs1KStatusInfo.EntityData.Leafs["replay-window-size"] = types.YLeaf{"ReplayWindowSize", ncs1KStatusInfo.ReplayWindowSize}
    ncs1KStatusInfo.EntityData.Leafs["must-secure"] = types.YLeaf{"MustSecure", ncs1KStatusInfo.MustSecure}
    ncs1KStatusInfo.EntityData.Leafs["secure-mode"] = types.YLeaf{"SecureMode", ncs1KStatusInfo.SecureMode}
    return &(ncs1KStatusInfo.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus
// Encrypt Secure Channel Status
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is Ncs1kCipherSuit.
    CipherSuite interface{}

    // Initial Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    InitialPacketNumber interface{}

    // Secure Tag Length. The type is interface{} with range: 0..4294967295.
    SecureTagLength interface{}

    // Maximum Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation.
    ActiveAssociation []Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetEntityData() *types.CommonEntityData {
    encryptScStatus.EntityData.YFilter = encryptScStatus.YFilter
    encryptScStatus.EntityData.YangName = "encrypt-sc-status"
    encryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    encryptScStatus.EntityData.ParentYangName = "ncs1k-status-info"
    encryptScStatus.EntityData.SegmentPath = "encrypt-sc-status"
    encryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptScStatus.EntityData.Children = make(map[string]types.YChild)
    encryptScStatus.EntityData.Children["active-association"] = types.YChild{"ActiveAssociation", nil}
    for i := range encryptScStatus.ActiveAssociation {
        encryptScStatus.EntityData.Children[types.GetSegmentPath(&encryptScStatus.ActiveAssociation[i])] = types.YChild{"ActiveAssociation", &encryptScStatus.ActiveAssociation[i]}
    }
    encryptScStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptScStatus.EntityData.Leafs["protection-enabled"] = types.YLeaf{"ProtectionEnabled", encryptScStatus.ProtectionEnabled}
    encryptScStatus.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", encryptScStatus.SecureChannelId}
    encryptScStatus.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", encryptScStatus.ConfidentialityOffset}
    encryptScStatus.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", encryptScStatus.CipherSuite}
    encryptScStatus.EntityData.Leafs["initial-packet-number"] = types.YLeaf{"InitialPacketNumber", encryptScStatus.InitialPacketNumber}
    encryptScStatus.EntityData.Leafs["secure-tag-length"] = types.YLeaf{"SecureTagLength", encryptScStatus.SecureTagLength}
    encryptScStatus.EntityData.Leafs["max-packet-number"] = types.YLeaf{"MaxPacketNumber", encryptScStatus.MaxPacketNumber}
    encryptScStatus.EntityData.Leafs["recent-packet-number"] = types.YLeaf{"RecentPacketNumber", encryptScStatus.RecentPacketNumber}
    return &(encryptScStatus.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation
// Active Associations
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assocition Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Devive Association Number. The type is interface{} with range: 0..255.
    DeviceAssociationNumber interface{}

    // Short Secure Channel Id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}

    // Key Programmed Time. The type is string with length: 0..30.
    ProgrammedTime interface{}

    // 32bit CRC of Programmed Key. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    KeyCrc interface{}

    // XPN Salt. The type is slice of string with pattern: b'[0-9a-fA-F]{1,8}'.
    XpnSalt []interface{}
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "encrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = make(map[string]types.YChild)
    activeAssociation.EntityData.Leafs = make(map[string]types.YLeaf)
    activeAssociation.EntityData.Leafs["association-number"] = types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber}
    activeAssociation.EntityData.Leafs["device-association-number"] = types.YLeaf{"DeviceAssociationNumber", activeAssociation.DeviceAssociationNumber}
    activeAssociation.EntityData.Leafs["short-secure-channel-id"] = types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId}
    activeAssociation.EntityData.Leafs["programmed-time"] = types.YLeaf{"ProgrammedTime", activeAssociation.ProgrammedTime}
    activeAssociation.EntityData.Leafs["key-crc"] = types.YLeaf{"KeyCrc", activeAssociation.KeyCrc}
    activeAssociation.EntityData.Leafs["xpn-salt"] = types.YLeaf{"XpnSalt", activeAssociation.XpnSalt}
    return &(activeAssociation.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus
// Decrypt Secure Channel Status
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is interface{} with range:
    // 0..18446744073709551615.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is Ncs1kCipherSuit.
    CipherSuite interface{}

    // Initial Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    InitialPacketNumber interface{}

    // Secure Tag Length. The type is interface{} with range: 0..4294967295.
    SecureTagLength interface{}

    // Maximum Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation.
    ActiveAssociation []Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetEntityData() *types.CommonEntityData {
    decryptScStatus.EntityData.YFilter = decryptScStatus.YFilter
    decryptScStatus.EntityData.YangName = "decrypt-sc-status"
    decryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    decryptScStatus.EntityData.ParentYangName = "ncs1k-status-info"
    decryptScStatus.EntityData.SegmentPath = "decrypt-sc-status"
    decryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decryptScStatus.EntityData.Children = make(map[string]types.YChild)
    decryptScStatus.EntityData.Children["active-association"] = types.YChild{"ActiveAssociation", nil}
    for i := range decryptScStatus.ActiveAssociation {
        decryptScStatus.EntityData.Children[types.GetSegmentPath(&decryptScStatus.ActiveAssociation[i])] = types.YChild{"ActiveAssociation", &decryptScStatus.ActiveAssociation[i]}
    }
    decryptScStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    decryptScStatus.EntityData.Leafs["protection-enabled"] = types.YLeaf{"ProtectionEnabled", decryptScStatus.ProtectionEnabled}
    decryptScStatus.EntityData.Leafs["secure-channel-id"] = types.YLeaf{"SecureChannelId", decryptScStatus.SecureChannelId}
    decryptScStatus.EntityData.Leafs["confidentiality-offset"] = types.YLeaf{"ConfidentialityOffset", decryptScStatus.ConfidentialityOffset}
    decryptScStatus.EntityData.Leafs["cipher-suite"] = types.YLeaf{"CipherSuite", decryptScStatus.CipherSuite}
    decryptScStatus.EntityData.Leafs["initial-packet-number"] = types.YLeaf{"InitialPacketNumber", decryptScStatus.InitialPacketNumber}
    decryptScStatus.EntityData.Leafs["secure-tag-length"] = types.YLeaf{"SecureTagLength", decryptScStatus.SecureTagLength}
    decryptScStatus.EntityData.Leafs["max-packet-number"] = types.YLeaf{"MaxPacketNumber", decryptScStatus.MaxPacketNumber}
    decryptScStatus.EntityData.Leafs["recent-packet-number"] = types.YLeaf{"RecentPacketNumber", decryptScStatus.RecentPacketNumber}
    return &(decryptScStatus.EntityData)
}

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation
// Active Associations
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Assocition Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Devive Association Number. The type is interface{} with range: 0..255.
    DeviceAssociationNumber interface{}

    // Short Secure Channel Id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}

    // Key Programmed Time. The type is string with length: 0..30.
    ProgrammedTime interface{}

    // 32bit CRC of Programmed Key. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    KeyCrc interface{}

    // XPN Salt. The type is slice of string with pattern: b'[0-9a-fA-F]{1,8}'.
    XpnSalt []interface{}
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "decrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = make(map[string]types.YChild)
    activeAssociation.EntityData.Leafs = make(map[string]types.YLeaf)
    activeAssociation.EntityData.Leafs["association-number"] = types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber}
    activeAssociation.EntityData.Leafs["device-association-number"] = types.YLeaf{"DeviceAssociationNumber", activeAssociation.DeviceAssociationNumber}
    activeAssociation.EntityData.Leafs["short-secure-channel-id"] = types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId}
    activeAssociation.EntityData.Leafs["programmed-time"] = types.YLeaf{"ProgrammedTime", activeAssociation.ProgrammedTime}
    activeAssociation.EntityData.Leafs["key-crc"] = types.YLeaf{"KeyCrc", activeAssociation.KeyCrc}
    activeAssociation.EntityData.Leafs["xpn-salt"] = types.YLeaf{"XpnSalt", activeAssociation.XpnSalt}
    return &(activeAssociation.EntityData)
}

