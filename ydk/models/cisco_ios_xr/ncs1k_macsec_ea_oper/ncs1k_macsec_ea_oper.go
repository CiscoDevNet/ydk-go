// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs1k-macsec-ea package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ncs1k-macsec-oper: Macsec data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs1k-macsec-ea-oper ncs1k-macsec-oper}", reflect.TypeOf(Ncs1kMacsecOper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs1k-macsec-ea-oper:ncs1k-macsec-oper", reflect.TypeOf(Ncs1kMacsecOper{}))
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

// Ncs1kMacsecOper
// Macsec data
type Ncs1kMacsecOper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All Macsec operational data.
    Ncs1kMacsecCtrlrNames Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames
}

func (ncs1kMacsecOper *Ncs1kMacsecOper) GetEntityData() *types.CommonEntityData {
    ncs1kMacsecOper.EntityData.YFilter = ncs1kMacsecOper.YFilter
    ncs1kMacsecOper.EntityData.YangName = "ncs1k-macsec-oper"
    ncs1kMacsecOper.EntityData.BundleName = "cisco_ios_xr"
    ncs1kMacsecOper.EntityData.ParentYangName = "Cisco-IOS-XR-ncs1k-macsec-ea-oper"
    ncs1kMacsecOper.EntityData.SegmentPath = "Cisco-IOS-XR-ncs1k-macsec-ea-oper:ncs1k-macsec-oper"
    ncs1kMacsecOper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1kMacsecOper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1kMacsecOper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1kMacsecOper.EntityData.Children = types.NewOrderedMap()
    ncs1kMacsecOper.EntityData.Children.Append("ncs1k-macsec-ctrlr-names", types.YChild{"Ncs1kMacsecCtrlrNames", &ncs1kMacsecOper.Ncs1kMacsecCtrlrNames})
    ncs1kMacsecOper.EntityData.Leafs = types.NewOrderedMap()

    ncs1kMacsecOper.EntityData.YListKeys = []string {}

    return &(ncs1kMacsecOper.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames
// All Macsec operational data
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName.
    Ncs1kMacsecCtrlrName []*Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName
}

func (ncs1kMacsecCtrlrNames *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames) GetEntityData() *types.CommonEntityData {
    ncs1kMacsecCtrlrNames.EntityData.YFilter = ncs1kMacsecCtrlrNames.YFilter
    ncs1kMacsecCtrlrNames.EntityData.YangName = "ncs1k-macsec-ctrlr-names"
    ncs1kMacsecCtrlrNames.EntityData.BundleName = "cisco_ios_xr"
    ncs1kMacsecCtrlrNames.EntityData.ParentYangName = "ncs1k-macsec-oper"
    ncs1kMacsecCtrlrNames.EntityData.SegmentPath = "ncs1k-macsec-ctrlr-names"
    ncs1kMacsecCtrlrNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1kMacsecCtrlrNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1kMacsecCtrlrNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1kMacsecCtrlrNames.EntityData.Children = types.NewOrderedMap()
    ncs1kMacsecCtrlrNames.EntityData.Children.Append("ncs1k-macsec-ctrlr-name", types.YChild{"Ncs1kMacsecCtrlrName", nil})
    for i := range ncs1kMacsecCtrlrNames.Ncs1kMacsecCtrlrName {
        ncs1kMacsecCtrlrNames.EntityData.Children.Append(types.GetSegmentPath(ncs1kMacsecCtrlrNames.Ncs1kMacsecCtrlrName[i]), types.YChild{"Ncs1kMacsecCtrlrName", ncs1kMacsecCtrlrNames.Ncs1kMacsecCtrlrName[i]})
    }
    ncs1kMacsecCtrlrNames.EntityData.Leafs = types.NewOrderedMap()

    ncs1kMacsecCtrlrNames.EntityData.YListKeys = []string {}

    return &(ncs1kMacsecCtrlrNames.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName
// Interface name
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    Name interface{}

    // controller data.
    Ncs1kStatusInfo Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo
}

func (ncs1kMacsecCtrlrName *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName) GetEntityData() *types.CommonEntityData {
    ncs1kMacsecCtrlrName.EntityData.YFilter = ncs1kMacsecCtrlrName.YFilter
    ncs1kMacsecCtrlrName.EntityData.YangName = "ncs1k-macsec-ctrlr-name"
    ncs1kMacsecCtrlrName.EntityData.BundleName = "cisco_ios_xr"
    ncs1kMacsecCtrlrName.EntityData.ParentYangName = "ncs1k-macsec-ctrlr-names"
    ncs1kMacsecCtrlrName.EntityData.SegmentPath = "ncs1k-macsec-ctrlr-name" + types.AddKeyToken(ncs1kMacsecCtrlrName.Name, "name")
    ncs1kMacsecCtrlrName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1kMacsecCtrlrName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1kMacsecCtrlrName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1kMacsecCtrlrName.EntityData.Children = types.NewOrderedMap()
    ncs1kMacsecCtrlrName.EntityData.Children.Append("ncs1k-status-info", types.YChild{"Ncs1kStatusInfo", &ncs1kMacsecCtrlrName.Ncs1kStatusInfo})
    ncs1kMacsecCtrlrName.EntityData.Leafs = types.NewOrderedMap()
    ncs1kMacsecCtrlrName.EntityData.Leafs.Append("name", types.YLeaf{"Name", ncs1kMacsecCtrlrName.Name})

    ncs1kMacsecCtrlrName.EntityData.YListKeys = []string {"Name"}

    return &(ncs1kMacsecCtrlrName.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo
// controller data
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    ReplayWindowSize interface{}

    // Must Secure. The type is bool.
    MustSecure interface{}

    // Secure Mode. The type is interface{} with range: 0..4294967295.
    SecureMode interface{}

    // Encrypt Secure Channel Status.
    EncryptScStatus Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus

    // Decrypt Secure Channel Status.
    DecryptScStatus Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus
}

func (ncs1kStatusInfo *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo) GetEntityData() *types.CommonEntityData {
    ncs1kStatusInfo.EntityData.YFilter = ncs1kStatusInfo.YFilter
    ncs1kStatusInfo.EntityData.YangName = "ncs1k-status-info"
    ncs1kStatusInfo.EntityData.BundleName = "cisco_ios_xr"
    ncs1kStatusInfo.EntityData.ParentYangName = "ncs1k-macsec-ctrlr-name"
    ncs1kStatusInfo.EntityData.SegmentPath = "ncs1k-status-info"
    ncs1kStatusInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncs1kStatusInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncs1kStatusInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncs1kStatusInfo.EntityData.Children = types.NewOrderedMap()
    ncs1kStatusInfo.EntityData.Children.Append("encrypt-sc-status", types.YChild{"EncryptScStatus", &ncs1kStatusInfo.EncryptScStatus})
    ncs1kStatusInfo.EntityData.Children.Append("decrypt-sc-status", types.YChild{"DecryptScStatus", &ncs1kStatusInfo.DecryptScStatus})
    ncs1kStatusInfo.EntityData.Leafs = types.NewOrderedMap()
    ncs1kStatusInfo.EntityData.Leafs.Append("replay-window-size", types.YLeaf{"ReplayWindowSize", ncs1kStatusInfo.ReplayWindowSize})
    ncs1kStatusInfo.EntityData.Leafs.Append("must-secure", types.YLeaf{"MustSecure", ncs1kStatusInfo.MustSecure})
    ncs1kStatusInfo.EntityData.Leafs.Append("secure-mode", types.YLeaf{"SecureMode", ncs1kStatusInfo.SecureMode})

    ncs1kStatusInfo.EntityData.YListKeys = []string {}

    return &(ncs1kStatusInfo.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus
// Encrypt Secure Channel Status
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus struct {
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
    // Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus_ActiveAssociation.
    ActiveAssociation []*Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus_ActiveAssociation
}

func (encryptScStatus *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus) GetEntityData() *types.CommonEntityData {
    encryptScStatus.EntityData.YFilter = encryptScStatus.YFilter
    encryptScStatus.EntityData.YangName = "encrypt-sc-status"
    encryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    encryptScStatus.EntityData.ParentYangName = "ncs1k-status-info"
    encryptScStatus.EntityData.SegmentPath = "encrypt-sc-status"
    encryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptScStatus.EntityData.Children = types.NewOrderedMap()
    encryptScStatus.EntityData.Children.Append("active-association", types.YChild{"ActiveAssociation", nil})
    for i := range encryptScStatus.ActiveAssociation {
        encryptScStatus.EntityData.Children.Append(types.GetSegmentPath(encryptScStatus.ActiveAssociation[i]), types.YChild{"ActiveAssociation", encryptScStatus.ActiveAssociation[i]})
    }
    encryptScStatus.EntityData.Leafs = types.NewOrderedMap()
    encryptScStatus.EntityData.Leafs.Append("protection-enabled", types.YLeaf{"ProtectionEnabled", encryptScStatus.ProtectionEnabled})
    encryptScStatus.EntityData.Leafs.Append("secure-channel-id", types.YLeaf{"SecureChannelId", encryptScStatus.SecureChannelId})
    encryptScStatus.EntityData.Leafs.Append("confidentiality-offset", types.YLeaf{"ConfidentialityOffset", encryptScStatus.ConfidentialityOffset})
    encryptScStatus.EntityData.Leafs.Append("cipher-suite", types.YLeaf{"CipherSuite", encryptScStatus.CipherSuite})
    encryptScStatus.EntityData.Leafs.Append("initial-packet-number", types.YLeaf{"InitialPacketNumber", encryptScStatus.InitialPacketNumber})
    encryptScStatus.EntityData.Leafs.Append("secure-tag-length", types.YLeaf{"SecureTagLength", encryptScStatus.SecureTagLength})
    encryptScStatus.EntityData.Leafs.Append("max-packet-number", types.YLeaf{"MaxPacketNumber", encryptScStatus.MaxPacketNumber})
    encryptScStatus.EntityData.Leafs.Append("recent-packet-number", types.YLeaf{"RecentPacketNumber", encryptScStatus.RecentPacketNumber})

    encryptScStatus.EntityData.YListKeys = []string {}

    return &(encryptScStatus.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus_ActiveAssociation
// Active Associations
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus_ActiveAssociation struct {
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
    // [0-9a-fA-F]{1,8}.
    KeyCrc interface{}

    // XPN Salt. The type is slice of string with pattern: [0-9a-fA-F]{1,8}.
    XpnSalt []interface{}
}

func (activeAssociation *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_EncryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "encrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = types.NewOrderedMap()
    activeAssociation.EntityData.Leafs = types.NewOrderedMap()
    activeAssociation.EntityData.Leafs.Append("association-number", types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber})
    activeAssociation.EntityData.Leafs.Append("device-association-number", types.YLeaf{"DeviceAssociationNumber", activeAssociation.DeviceAssociationNumber})
    activeAssociation.EntityData.Leafs.Append("short-secure-channel-id", types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId})
    activeAssociation.EntityData.Leafs.Append("programmed-time", types.YLeaf{"ProgrammedTime", activeAssociation.ProgrammedTime})
    activeAssociation.EntityData.Leafs.Append("key-crc", types.YLeaf{"KeyCrc", activeAssociation.KeyCrc})
    activeAssociation.EntityData.Leafs.Append("xpn-salt", types.YLeaf{"XpnSalt", activeAssociation.XpnSalt})

    activeAssociation.EntityData.YListKeys = []string {}

    return &(activeAssociation.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus
// Decrypt Secure Channel Status
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus struct {
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
    // Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus_ActiveAssociation.
    ActiveAssociation []*Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus_ActiveAssociation
}

func (decryptScStatus *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus) GetEntityData() *types.CommonEntityData {
    decryptScStatus.EntityData.YFilter = decryptScStatus.YFilter
    decryptScStatus.EntityData.YangName = "decrypt-sc-status"
    decryptScStatus.EntityData.BundleName = "cisco_ios_xr"
    decryptScStatus.EntityData.ParentYangName = "ncs1k-status-info"
    decryptScStatus.EntityData.SegmentPath = "decrypt-sc-status"
    decryptScStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    decryptScStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    decryptScStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    decryptScStatus.EntityData.Children = types.NewOrderedMap()
    decryptScStatus.EntityData.Children.Append("active-association", types.YChild{"ActiveAssociation", nil})
    for i := range decryptScStatus.ActiveAssociation {
        decryptScStatus.EntityData.Children.Append(types.GetSegmentPath(decryptScStatus.ActiveAssociation[i]), types.YChild{"ActiveAssociation", decryptScStatus.ActiveAssociation[i]})
    }
    decryptScStatus.EntityData.Leafs = types.NewOrderedMap()
    decryptScStatus.EntityData.Leafs.Append("protection-enabled", types.YLeaf{"ProtectionEnabled", decryptScStatus.ProtectionEnabled})
    decryptScStatus.EntityData.Leafs.Append("secure-channel-id", types.YLeaf{"SecureChannelId", decryptScStatus.SecureChannelId})
    decryptScStatus.EntityData.Leafs.Append("confidentiality-offset", types.YLeaf{"ConfidentialityOffset", decryptScStatus.ConfidentialityOffset})
    decryptScStatus.EntityData.Leafs.Append("cipher-suite", types.YLeaf{"CipherSuite", decryptScStatus.CipherSuite})
    decryptScStatus.EntityData.Leafs.Append("initial-packet-number", types.YLeaf{"InitialPacketNumber", decryptScStatus.InitialPacketNumber})
    decryptScStatus.EntityData.Leafs.Append("secure-tag-length", types.YLeaf{"SecureTagLength", decryptScStatus.SecureTagLength})
    decryptScStatus.EntityData.Leafs.Append("max-packet-number", types.YLeaf{"MaxPacketNumber", decryptScStatus.MaxPacketNumber})
    decryptScStatus.EntityData.Leafs.Append("recent-packet-number", types.YLeaf{"RecentPacketNumber", decryptScStatus.RecentPacketNumber})

    decryptScStatus.EntityData.YListKeys = []string {}

    return &(decryptScStatus.EntityData)
}

// Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus_ActiveAssociation
// Active Associations
type Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus_ActiveAssociation struct {
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
    // [0-9a-fA-F]{1,8}.
    KeyCrc interface{}

    // XPN Salt. The type is slice of string with pattern: [0-9a-fA-F]{1,8}.
    XpnSalt []interface{}
}

func (activeAssociation *Ncs1kMacsecOper_Ncs1kMacsecCtrlrNames_Ncs1kMacsecCtrlrName_Ncs1kStatusInfo_DecryptScStatus_ActiveAssociation) GetEntityData() *types.CommonEntityData {
    activeAssociation.EntityData.YFilter = activeAssociation.YFilter
    activeAssociation.EntityData.YangName = "active-association"
    activeAssociation.EntityData.BundleName = "cisco_ios_xr"
    activeAssociation.EntityData.ParentYangName = "decrypt-sc-status"
    activeAssociation.EntityData.SegmentPath = "active-association"
    activeAssociation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeAssociation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeAssociation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeAssociation.EntityData.Children = types.NewOrderedMap()
    activeAssociation.EntityData.Leafs = types.NewOrderedMap()
    activeAssociation.EntityData.Leafs.Append("association-number", types.YLeaf{"AssociationNumber", activeAssociation.AssociationNumber})
    activeAssociation.EntityData.Leafs.Append("device-association-number", types.YLeaf{"DeviceAssociationNumber", activeAssociation.DeviceAssociationNumber})
    activeAssociation.EntityData.Leafs.Append("short-secure-channel-id", types.YLeaf{"ShortSecureChannelId", activeAssociation.ShortSecureChannelId})
    activeAssociation.EntityData.Leafs.Append("programmed-time", types.YLeaf{"ProgrammedTime", activeAssociation.ProgrammedTime})
    activeAssociation.EntityData.Leafs.Append("key-crc", types.YLeaf{"KeyCrc", activeAssociation.KeyCrc})
    activeAssociation.EntityData.Leafs.Append("xpn-salt", types.YLeaf{"XpnSalt", activeAssociation.XpnSalt})

    activeAssociation.EntityData.YListKeys = []string {}

    return &(activeAssociation.EntityData)
}

