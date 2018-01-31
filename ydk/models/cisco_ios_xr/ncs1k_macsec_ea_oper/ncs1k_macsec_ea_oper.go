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
    parent types.Entity
    YFilter yfilter.YFilter

    // All Macsec operational data.
    Ncs1KMacsecCtrlrNames Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetFilter() yfilter.YFilter { return ncs1KMacsecOper.YFilter }

func (ncs1KMacsecOper *Ncs1KMacsecOper) SetFilter(yf yfilter.YFilter) { ncs1KMacsecOper.YFilter = yf }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetGoName(yname string) string {
    if yname == "ncs1k-macsec-ctrlr-names" { return "Ncs1KMacsecCtrlrNames" }
    return ""
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs1k-macsec-ea-oper:ncs1k-macsec-oper"
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ncs1k-macsec-ctrlr-names" {
        return &ncs1KMacsecOper.Ncs1KMacsecCtrlrNames
    }
    return nil
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ncs1k-macsec-ctrlr-names"] = &ncs1KMacsecOper.Ncs1KMacsecCtrlrNames
    return children
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetBundleName() string { return "cisco_ios_xr" }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetYangName() string { return "ncs1k-macsec-oper" }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncs1KMacsecOper *Ncs1KMacsecOper) SetParent(parent types.Entity) { ncs1KMacsecOper.parent = parent }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetParent() types.Entity { return ncs1KMacsecOper.parent }

func (ncs1KMacsecOper *Ncs1KMacsecOper) GetParentYangName() string { return "Cisco-IOS-XR-ncs1k-macsec-ea-oper" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames
// All Macsec operational data
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is slice of
    // Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName.
    Ncs1KMacsecCtrlrName []Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetFilter() yfilter.YFilter { return ncs1KMacsecCtrlrNames.YFilter }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) SetFilter(yf yfilter.YFilter) { ncs1KMacsecCtrlrNames.YFilter = yf }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetGoName(yname string) string {
    if yname == "ncs1k-macsec-ctrlr-name" { return "Ncs1KMacsecCtrlrName" }
    return ""
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetSegmentPath() string {
    return "ncs1k-macsec-ctrlr-names"
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ncs1k-macsec-ctrlr-name" {
        for _, c := range ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName {
            if ncs1KMacsecCtrlrNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName{}
        ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName = append(ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName, child)
        return &ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName[len(ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName)-1]
    }
    return nil
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName {
        children[ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName[i].GetSegmentPath()] = &ncs1KMacsecCtrlrNames.Ncs1KMacsecCtrlrName[i]
    }
    return children
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetBundleName() string { return "cisco_ios_xr" }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetYangName() string { return "ncs1k-macsec-ctrlr-names" }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) SetParent(parent types.Entity) { ncs1KMacsecCtrlrNames.parent = parent }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetParent() types.Entity { return ncs1KMacsecCtrlrNames.parent }

func (ncs1KMacsecCtrlrNames *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames) GetParentYangName() string { return "ncs1k-macsec-oper" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName
// Interface name
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // controller data.
    Ncs1KStatusInfo Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetFilter() yfilter.YFilter { return ncs1KMacsecCtrlrName.YFilter }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) SetFilter(yf yfilter.YFilter) { ncs1KMacsecCtrlrName.YFilter = yf }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "ncs1k-status-info" { return "Ncs1KStatusInfo" }
    return ""
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetSegmentPath() string {
    return "ncs1k-macsec-ctrlr-name" + "[name='" + fmt.Sprintf("%v", ncs1KMacsecCtrlrName.Name) + "']"
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ncs1k-status-info" {
        return &ncs1KMacsecCtrlrName.Ncs1KStatusInfo
    }
    return nil
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ncs1k-status-info"] = &ncs1KMacsecCtrlrName.Ncs1KStatusInfo
    return children
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ncs1KMacsecCtrlrName.Name
    return leafs
}

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetBundleName() string { return "cisco_ios_xr" }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetYangName() string { return "ncs1k-macsec-ctrlr-name" }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) SetParent(parent types.Entity) { ncs1KMacsecCtrlrName.parent = parent }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetParent() types.Entity { return ncs1KMacsecCtrlrName.parent }

func (ncs1KMacsecCtrlrName *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName) GetParentYangName() string { return "ncs1k-macsec-ctrlr-names" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo
// controller data
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo struct {
    parent types.Entity
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

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetFilter() yfilter.YFilter { return ncs1KStatusInfo.YFilter }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) SetFilter(yf yfilter.YFilter) { ncs1KStatusInfo.YFilter = yf }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetGoName(yname string) string {
    if yname == "replay-window-size" { return "ReplayWindowSize" }
    if yname == "must-secure" { return "MustSecure" }
    if yname == "secure-mode" { return "SecureMode" }
    if yname == "encrypt-sc-status" { return "EncryptScStatus" }
    if yname == "decrypt-sc-status" { return "DecryptScStatus" }
    return ""
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetSegmentPath() string {
    return "ncs1k-status-info"
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypt-sc-status" {
        return &ncs1KStatusInfo.EncryptScStatus
    }
    if childYangName == "decrypt-sc-status" {
        return &ncs1KStatusInfo.DecryptScStatus
    }
    return nil
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encrypt-sc-status"] = &ncs1KStatusInfo.EncryptScStatus
    children["decrypt-sc-status"] = &ncs1KStatusInfo.DecryptScStatus
    return children
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["replay-window-size"] = ncs1KStatusInfo.ReplayWindowSize
    leafs["must-secure"] = ncs1KStatusInfo.MustSecure
    leafs["secure-mode"] = ncs1KStatusInfo.SecureMode
    return leafs
}

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetYangName() string { return "ncs1k-status-info" }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) SetParent(parent types.Entity) { ncs1KStatusInfo.parent = parent }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetParent() types.Entity { return ncs1KStatusInfo.parent }

func (ncs1KStatusInfo *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo) GetParentYangName() string { return "ncs1k-macsec-ctrlr-name" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus
// Encrypt Secure Channel Status
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus struct {
    parent types.Entity
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

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetFilter() yfilter.YFilter { return encryptScStatus.YFilter }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) SetFilter(yf yfilter.YFilter) { encryptScStatus.YFilter = yf }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetGoName(yname string) string {
    if yname == "protection-enabled" { return "ProtectionEnabled" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "initial-packet-number" { return "InitialPacketNumber" }
    if yname == "secure-tag-length" { return "SecureTagLength" }
    if yname == "max-packet-number" { return "MaxPacketNumber" }
    if yname == "recent-packet-number" { return "RecentPacketNumber" }
    if yname == "active-association" { return "ActiveAssociation" }
    return ""
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetSegmentPath() string {
    return "encrypt-sc-status"
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-association" {
        for _, c := range encryptScStatus.ActiveAssociation {
            if encryptScStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation{}
        encryptScStatus.ActiveAssociation = append(encryptScStatus.ActiveAssociation, child)
        return &encryptScStatus.ActiveAssociation[len(encryptScStatus.ActiveAssociation)-1]
    }
    return nil
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptScStatus.ActiveAssociation {
        children[encryptScStatus.ActiveAssociation[i].GetSegmentPath()] = &encryptScStatus.ActiveAssociation[i]
    }
    return children
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enabled"] = encryptScStatus.ProtectionEnabled
    leafs["secure-channel-id"] = encryptScStatus.SecureChannelId
    leafs["confidentiality-offset"] = encryptScStatus.ConfidentialityOffset
    leafs["cipher-suite"] = encryptScStatus.CipherSuite
    leafs["initial-packet-number"] = encryptScStatus.InitialPacketNumber
    leafs["secure-tag-length"] = encryptScStatus.SecureTagLength
    leafs["max-packet-number"] = encryptScStatus.MaxPacketNumber
    leafs["recent-packet-number"] = encryptScStatus.RecentPacketNumber
    return leafs
}

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetBundleName() string { return "cisco_ios_xr" }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetYangName() string { return "encrypt-sc-status" }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) SetParent(parent types.Entity) { encryptScStatus.parent = parent }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetParent() types.Entity { return encryptScStatus.parent }

func (encryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus) GetParentYangName() string { return "ncs1k-status-info" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation
// Active Associations
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation struct {
    parent types.Entity
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

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetFilter() yfilter.YFilter { return activeAssociation.YFilter }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) SetFilter(yf yfilter.YFilter) { activeAssociation.YFilter = yf }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetGoName(yname string) string {
    if yname == "association-number" { return "AssociationNumber" }
    if yname == "device-association-number" { return "DeviceAssociationNumber" }
    if yname == "short-secure-channel-id" { return "ShortSecureChannelId" }
    if yname == "programmed-time" { return "ProgrammedTime" }
    if yname == "key-crc" { return "KeyCrc" }
    if yname == "xpn-salt" { return "XpnSalt" }
    return ""
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetSegmentPath() string {
    return "active-association"
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-number"] = activeAssociation.AssociationNumber
    leafs["device-association-number"] = activeAssociation.DeviceAssociationNumber
    leafs["short-secure-channel-id"] = activeAssociation.ShortSecureChannelId
    leafs["programmed-time"] = activeAssociation.ProgrammedTime
    leafs["key-crc"] = activeAssociation.KeyCrc
    leafs["xpn-salt"] = activeAssociation.XpnSalt
    return leafs
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetBundleName() string { return "cisco_ios_xr" }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetYangName() string { return "active-association" }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) SetParent(parent types.Entity) { activeAssociation.parent = parent }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetParent() types.Entity { return activeAssociation.parent }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_EncryptScStatus_ActiveAssociation) GetParentYangName() string { return "encrypt-sc-status" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus
// Decrypt Secure Channel Status
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus struct {
    parent types.Entity
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

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetFilter() yfilter.YFilter { return decryptScStatus.YFilter }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) SetFilter(yf yfilter.YFilter) { decryptScStatus.YFilter = yf }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetGoName(yname string) string {
    if yname == "protection-enabled" { return "ProtectionEnabled" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "initial-packet-number" { return "InitialPacketNumber" }
    if yname == "secure-tag-length" { return "SecureTagLength" }
    if yname == "max-packet-number" { return "MaxPacketNumber" }
    if yname == "recent-packet-number" { return "RecentPacketNumber" }
    if yname == "active-association" { return "ActiveAssociation" }
    return ""
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetSegmentPath() string {
    return "decrypt-sc-status"
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-association" {
        for _, c := range decryptScStatus.ActiveAssociation {
            if decryptScStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation{}
        decryptScStatus.ActiveAssociation = append(decryptScStatus.ActiveAssociation, child)
        return &decryptScStatus.ActiveAssociation[len(decryptScStatus.ActiveAssociation)-1]
    }
    return nil
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range decryptScStatus.ActiveAssociation {
        children[decryptScStatus.ActiveAssociation[i].GetSegmentPath()] = &decryptScStatus.ActiveAssociation[i]
    }
    return children
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enabled"] = decryptScStatus.ProtectionEnabled
    leafs["secure-channel-id"] = decryptScStatus.SecureChannelId
    leafs["confidentiality-offset"] = decryptScStatus.ConfidentialityOffset
    leafs["cipher-suite"] = decryptScStatus.CipherSuite
    leafs["initial-packet-number"] = decryptScStatus.InitialPacketNumber
    leafs["secure-tag-length"] = decryptScStatus.SecureTagLength
    leafs["max-packet-number"] = decryptScStatus.MaxPacketNumber
    leafs["recent-packet-number"] = decryptScStatus.RecentPacketNumber
    return leafs
}

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetBundleName() string { return "cisco_ios_xr" }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetYangName() string { return "decrypt-sc-status" }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) SetParent(parent types.Entity) { decryptScStatus.parent = parent }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetParent() types.Entity { return decryptScStatus.parent }

func (decryptScStatus *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus) GetParentYangName() string { return "ncs1k-status-info" }

// Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation
// Active Associations
type Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation struct {
    parent types.Entity
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

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetFilter() yfilter.YFilter { return activeAssociation.YFilter }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) SetFilter(yf yfilter.YFilter) { activeAssociation.YFilter = yf }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetGoName(yname string) string {
    if yname == "association-number" { return "AssociationNumber" }
    if yname == "device-association-number" { return "DeviceAssociationNumber" }
    if yname == "short-secure-channel-id" { return "ShortSecureChannelId" }
    if yname == "programmed-time" { return "ProgrammedTime" }
    if yname == "key-crc" { return "KeyCrc" }
    if yname == "xpn-salt" { return "XpnSalt" }
    return ""
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetSegmentPath() string {
    return "active-association"
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-number"] = activeAssociation.AssociationNumber
    leafs["device-association-number"] = activeAssociation.DeviceAssociationNumber
    leafs["short-secure-channel-id"] = activeAssociation.ShortSecureChannelId
    leafs["programmed-time"] = activeAssociation.ProgrammedTime
    leafs["key-crc"] = activeAssociation.KeyCrc
    leafs["xpn-salt"] = activeAssociation.XpnSalt
    return leafs
}

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetBundleName() string { return "cisco_ios_xr" }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetYangName() string { return "active-association" }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) SetParent(parent types.Entity) { activeAssociation.parent = parent }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetParent() types.Entity { return activeAssociation.parent }

func (activeAssociation *Ncs1KMacsecOper_Ncs1KMacsecCtrlrNames_Ncs1KMacsecCtrlrName_Ncs1KStatusInfo_DecryptScStatus_ActiveAssociation) GetParentYangName() string { return "decrypt-sc-status" }

