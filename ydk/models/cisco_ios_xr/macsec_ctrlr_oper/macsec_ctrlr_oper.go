// This module contains a collection of YANG definitions
// for Cisco IOS-XR macsec-ctrlr package operational data.
// 
// This module contains definitions
// for the following management objects:
//   macsec-ctrlr-oper: Macsec controller data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package macsec_ctrlr_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package macsec_ctrlr_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-macsec-ctrlr-oper macsec-ctrlr-oper}", reflect.TypeOf(MacsecCtrlrOper{}))
    ydk.RegisterEntity("Cisco-IOS-XR-macsec-ctrlr-oper:macsec-ctrlr-oper", reflect.TypeOf(MacsecCtrlrOper{}))
}

// MacsecCtrlrCiphersuit represents Macsec ctrlr ciphersuit
type MacsecCtrlrCiphersuit string

const (
    // GCM AES 256
    MacsecCtrlrCiphersuit_gcm_aes_256 MacsecCtrlrCiphersuit = "gcm-aes-256"

    // GCM AES 128
    MacsecCtrlrCiphersuit_gcm_aes_128 MacsecCtrlrCiphersuit = "gcm-aes-128"

    // GCM AES XPN 256
    MacsecCtrlrCiphersuit_gcm_aes_xpn_256 MacsecCtrlrCiphersuit = "gcm-aes-xpn-256"
)

// MacsecCtrlrState represents Macsec ctrlr state
type MacsecCtrlrState string

const (
    // Up
    MacsecCtrlrState_macsec_ctrlr_state_up MacsecCtrlrState = "macsec-ctrlr-state-up"

    // Down
    MacsecCtrlrState_macsec_ctrlr_state_down MacsecCtrlrState = "macsec-ctrlr-state-down"

    // Administratively Down
    MacsecCtrlrState_macsec_ctrlr_state_admin_down MacsecCtrlrState = "macsec-ctrlr-state-admin-down"
)

// MacsecCtrlrOper
// Macsec controller data
type MacsecCtrlrOper struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All Macsec Controller Port operational data.
    MacsecCtrlrPorts MacsecCtrlrOper_MacsecCtrlrPorts
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetFilter() yfilter.YFilter { return macsecCtrlrOper.YFilter }

func (macsecCtrlrOper *MacsecCtrlrOper) SetFilter(yf yfilter.YFilter) { macsecCtrlrOper.YFilter = yf }

func (macsecCtrlrOper *MacsecCtrlrOper) GetGoName(yname string) string {
    if yname == "macsec-ctrlr-ports" { return "MacsecCtrlrPorts" }
    return ""
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetSegmentPath() string {
    return "Cisco-IOS-XR-macsec-ctrlr-oper:macsec-ctrlr-oper"
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec-ctrlr-ports" {
        return &macsecCtrlrOper.MacsecCtrlrPorts
    }
    return nil
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macsec-ctrlr-ports"] = &macsecCtrlrOper.MacsecCtrlrPorts
    return children
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsecCtrlrOper *MacsecCtrlrOper) GetBundleName() string { return "cisco_ios_xr" }

func (macsecCtrlrOper *MacsecCtrlrOper) GetYangName() string { return "macsec-ctrlr-oper" }

func (macsecCtrlrOper *MacsecCtrlrOper) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecCtrlrOper *MacsecCtrlrOper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecCtrlrOper *MacsecCtrlrOper) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecCtrlrOper *MacsecCtrlrOper) SetParent(parent types.Entity) { macsecCtrlrOper.parent = parent }

func (macsecCtrlrOper *MacsecCtrlrOper) GetParent() types.Entity { return macsecCtrlrOper.parent }

func (macsecCtrlrOper *MacsecCtrlrOper) GetParentYangName() string { return "Cisco-IOS-XR-macsec-ctrlr-oper" }

// MacsecCtrlrOper_MacsecCtrlrPorts
// All Macsec Controller Port operational data
type MacsecCtrlrOper_MacsecCtrlrPorts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Controller name. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort.
    MacsecCtrlrPort []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetFilter() yfilter.YFilter { return macsecCtrlrPorts.YFilter }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) SetFilter(yf yfilter.YFilter) { macsecCtrlrPorts.YFilter = yf }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetGoName(yname string) string {
    if yname == "macsec-ctrlr-port" { return "MacsecCtrlrPort" }
    return ""
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetSegmentPath() string {
    return "macsec-ctrlr-ports"
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec-ctrlr-port" {
        for _, c := range macsecCtrlrPorts.MacsecCtrlrPort {
            if macsecCtrlrPorts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort{}
        macsecCtrlrPorts.MacsecCtrlrPort = append(macsecCtrlrPorts.MacsecCtrlrPort, child)
        return &macsecCtrlrPorts.MacsecCtrlrPort[len(macsecCtrlrPorts.MacsecCtrlrPort)-1]
    }
    return nil
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range macsecCtrlrPorts.MacsecCtrlrPort {
        children[macsecCtrlrPorts.MacsecCtrlrPort[i].GetSegmentPath()] = &macsecCtrlrPorts.MacsecCtrlrPort[i]
    }
    return children
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetBundleName() string { return "cisco_ios_xr" }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetYangName() string { return "macsec-ctrlr-ports" }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) SetParent(parent types.Entity) { macsecCtrlrPorts.parent = parent }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetParent() types.Entity { return macsecCtrlrPorts.parent }

func (macsecCtrlrPorts *MacsecCtrlrOper_MacsecCtrlrPorts) GetParentYangName() string { return "macsec-ctrlr-oper" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort
// Controller name
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Port name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    Name interface{}

    // Macsec Controller operational data.
    MacsecCtrlrInfo MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetFilter() yfilter.YFilter { return macsecCtrlrPort.YFilter }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) SetFilter(yf yfilter.YFilter) { macsecCtrlrPort.YFilter = yf }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "macsec-ctrlr-info" { return "MacsecCtrlrInfo" }
    return ""
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetSegmentPath() string {
    return "macsec-ctrlr-port" + "[name='" + fmt.Sprintf("%v", macsecCtrlrPort.Name) + "']"
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "macsec-ctrlr-info" {
        return &macsecCtrlrPort.MacsecCtrlrInfo
    }
    return nil
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["macsec-ctrlr-info"] = &macsecCtrlrPort.MacsecCtrlrInfo
    return children
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = macsecCtrlrPort.Name
    return leafs
}

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetBundleName() string { return "cisco_ios_xr" }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetYangName() string { return "macsec-ctrlr-port" }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) SetParent(parent types.Entity) { macsecCtrlrPort.parent = parent }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetParent() types.Entity { return macsecCtrlrPort.parent }

func (macsecCtrlrPort *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort) GetParentYangName() string { return "macsec-ctrlr-ports" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo
// Macsec Controller operational data
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // State. The type is MacsecCtrlrState.
    State interface{}

    // Replay Window Size. The type is interface{} with range: 0..4294967295.
    ReplayWindowSize interface{}

    // Must Secure. The type is bool.
    MustSecure interface{}

    // Secure Mode. The type is interface{} with range: 0..4294967295.
    SecureMode interface{}

    // Encrypt Secure Channel Status.
    EncryptScStatus MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus

    // Decrypt Secure Channel Status.
    DecryptScStatus MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetFilter() yfilter.YFilter { return macsecCtrlrInfo.YFilter }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) SetFilter(yf yfilter.YFilter) { macsecCtrlrInfo.YFilter = yf }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "replay-window-size" { return "ReplayWindowSize" }
    if yname == "must-secure" { return "MustSecure" }
    if yname == "secure-mode" { return "SecureMode" }
    if yname == "encrypt-sc-status" { return "EncryptScStatus" }
    if yname == "decrypt-sc-status" { return "DecryptScStatus" }
    return ""
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetSegmentPath() string {
    return "macsec-ctrlr-info"
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypt-sc-status" {
        return &macsecCtrlrInfo.EncryptScStatus
    }
    if childYangName == "decrypt-sc-status" {
        return &macsecCtrlrInfo.DecryptScStatus
    }
    return nil
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encrypt-sc-status"] = &macsecCtrlrInfo.EncryptScStatus
    children["decrypt-sc-status"] = &macsecCtrlrInfo.DecryptScStatus
    return children
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = macsecCtrlrInfo.State
    leafs["replay-window-size"] = macsecCtrlrInfo.ReplayWindowSize
    leafs["must-secure"] = macsecCtrlrInfo.MustSecure
    leafs["secure-mode"] = macsecCtrlrInfo.SecureMode
    return leafs
}

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetBundleName() string { return "cisco_ios_xr" }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetYangName() string { return "macsec-ctrlr-info" }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) SetParent(parent types.Entity) { macsecCtrlrInfo.parent = parent }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetParent() types.Entity { return macsecCtrlrInfo.parent }

func (macsecCtrlrInfo *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo) GetParentYangName() string { return "macsec-ctrlr-port" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus
// Encrypt Secure Channel Status
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is string.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is MacsecCtrlrCiphersuit.
    CipherSuite interface{}

    // Max packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation.
    ActiveAssociation []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetFilter() yfilter.YFilter { return encryptScStatus.YFilter }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) SetFilter(yf yfilter.YFilter) { encryptScStatus.YFilter = yf }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetGoName(yname string) string {
    if yname == "protection-enabled" { return "ProtectionEnabled" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "max-packet-number" { return "MaxPacketNumber" }
    if yname == "recent-packet-number" { return "RecentPacketNumber" }
    if yname == "active-association" { return "ActiveAssociation" }
    return ""
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetSegmentPath() string {
    return "encrypt-sc-status"
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-association" {
        for _, c := range encryptScStatus.ActiveAssociation {
            if encryptScStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation{}
        encryptScStatus.ActiveAssociation = append(encryptScStatus.ActiveAssociation, child)
        return &encryptScStatus.ActiveAssociation[len(encryptScStatus.ActiveAssociation)-1]
    }
    return nil
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptScStatus.ActiveAssociation {
        children[encryptScStatus.ActiveAssociation[i].GetSegmentPath()] = &encryptScStatus.ActiveAssociation[i]
    }
    return children
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enabled"] = encryptScStatus.ProtectionEnabled
    leafs["secure-channel-id"] = encryptScStatus.SecureChannelId
    leafs["confidentiality-offset"] = encryptScStatus.ConfidentialityOffset
    leafs["cipher-suite"] = encryptScStatus.CipherSuite
    leafs["max-packet-number"] = encryptScStatus.MaxPacketNumber
    leafs["recent-packet-number"] = encryptScStatus.RecentPacketNumber
    return leafs
}

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetBundleName() string { return "cisco_ios_xr" }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetYangName() string { return "encrypt-sc-status" }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) SetParent(parent types.Entity) { encryptScStatus.parent = parent }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetParent() types.Entity { return encryptScStatus.parent }

func (encryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus) GetParentYangName() string { return "macsec-ctrlr-info" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation
// Active Associations
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Short secure channel id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetFilter() yfilter.YFilter { return activeAssociation.YFilter }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) SetFilter(yf yfilter.YFilter) { activeAssociation.YFilter = yf }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetGoName(yname string) string {
    if yname == "association-number" { return "AssociationNumber" }
    if yname == "short-secure-channel-id" { return "ShortSecureChannelId" }
    return ""
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetSegmentPath() string {
    return "active-association"
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-number"] = activeAssociation.AssociationNumber
    leafs["short-secure-channel-id"] = activeAssociation.ShortSecureChannelId
    return leafs
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetBundleName() string { return "cisco_ios_xr" }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetYangName() string { return "active-association" }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) SetParent(parent types.Entity) { activeAssociation.parent = parent }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetParent() types.Entity { return activeAssociation.parent }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_EncryptScStatus_ActiveAssociation) GetParentYangName() string { return "encrypt-sc-status" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus
// Decrypt Secure Channel Status
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protection Enabled. The type is bool.
    ProtectionEnabled interface{}

    // Secure Channel Id. The type is string.
    SecureChannelId interface{}

    // Confidentiality offset. The type is interface{} with range: 0..4294967295.
    ConfidentialityOffset interface{}

    // Cipher Suite. The type is MacsecCtrlrCiphersuit.
    CipherSuite interface{}

    // Max packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxPacketNumber interface{}

    // Recent Packet Number. The type is interface{} with range:
    // 0..18446744073709551615.
    RecentPacketNumber interface{}

    // Active Associations. The type is slice of
    // MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation.
    ActiveAssociation []MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetFilter() yfilter.YFilter { return decryptScStatus.YFilter }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) SetFilter(yf yfilter.YFilter) { decryptScStatus.YFilter = yf }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetGoName(yname string) string {
    if yname == "protection-enabled" { return "ProtectionEnabled" }
    if yname == "secure-channel-id" { return "SecureChannelId" }
    if yname == "confidentiality-offset" { return "ConfidentialityOffset" }
    if yname == "cipher-suite" { return "CipherSuite" }
    if yname == "max-packet-number" { return "MaxPacketNumber" }
    if yname == "recent-packet-number" { return "RecentPacketNumber" }
    if yname == "active-association" { return "ActiveAssociation" }
    return ""
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetSegmentPath() string {
    return "decrypt-sc-status"
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "active-association" {
        for _, c := range decryptScStatus.ActiveAssociation {
            if decryptScStatus.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation{}
        decryptScStatus.ActiveAssociation = append(decryptScStatus.ActiveAssociation, child)
        return &decryptScStatus.ActiveAssociation[len(decryptScStatus.ActiveAssociation)-1]
    }
    return nil
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range decryptScStatus.ActiveAssociation {
        children[decryptScStatus.ActiveAssociation[i].GetSegmentPath()] = &decryptScStatus.ActiveAssociation[i]
    }
    return children
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protection-enabled"] = decryptScStatus.ProtectionEnabled
    leafs["secure-channel-id"] = decryptScStatus.SecureChannelId
    leafs["confidentiality-offset"] = decryptScStatus.ConfidentialityOffset
    leafs["cipher-suite"] = decryptScStatus.CipherSuite
    leafs["max-packet-number"] = decryptScStatus.MaxPacketNumber
    leafs["recent-packet-number"] = decryptScStatus.RecentPacketNumber
    return leafs
}

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetBundleName() string { return "cisco_ios_xr" }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetYangName() string { return "decrypt-sc-status" }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) SetParent(parent types.Entity) { decryptScStatus.parent = parent }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetParent() types.Entity { return decryptScStatus.parent }

func (decryptScStatus *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus) GetParentYangName() string { return "macsec-ctrlr-info" }

// MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation
// Active Associations
type MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association Number. The type is interface{} with range: 0..255.
    AssociationNumber interface{}

    // Short secure channel id. The type is interface{} with range: 0..4294967295.
    ShortSecureChannelId interface{}
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetFilter() yfilter.YFilter { return activeAssociation.YFilter }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) SetFilter(yf yfilter.YFilter) { activeAssociation.YFilter = yf }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetGoName(yname string) string {
    if yname == "association-number" { return "AssociationNumber" }
    if yname == "short-secure-channel-id" { return "ShortSecureChannelId" }
    return ""
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetSegmentPath() string {
    return "active-association"
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-number"] = activeAssociation.AssociationNumber
    leafs["short-secure-channel-id"] = activeAssociation.ShortSecureChannelId
    return leafs
}

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetBundleName() string { return "cisco_ios_xr" }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetYangName() string { return "active-association" }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) SetParent(parent types.Entity) { activeAssociation.parent = parent }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetParent() types.Entity { return activeAssociation.parent }

func (activeAssociation *MacsecCtrlrOper_MacsecCtrlrPorts_MacsecCtrlrPort_MacsecCtrlrInfo_DecryptScStatus_ActiveAssociation) GetParentYangName() string { return "decrypt-sc-status" }

