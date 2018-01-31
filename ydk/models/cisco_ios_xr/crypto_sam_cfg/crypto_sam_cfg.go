// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-sam package configuration.
// 
// This module contains definitions
// for the following management objects:
//   crypto: Crypto configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_sam_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_sam_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-sam-cfg crypto}", reflect.TypeOf(Crypto{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-sam-cfg:crypto", reflect.TypeOf(Crypto{}))
}

// CryptoSamAction represents Crypto sam action
type CryptoSamAction string

const (
    // To respond YES to the SAM prompt
    CryptoSamAction_proceed CryptoSamAction = "proceed"

    // To respond NO to the SAM prompt
    CryptoSamAction_terminate CryptoSamAction = "terminate"
)

// Crypto
// Crypto configuration
type Crypto struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Software Authentication Manager (SAM) Config.
    Sam Crypto_Sam

    // Secure Shell configuration.
    Ssh Crypto_Ssh
}

func (crypto *Crypto) GetFilter() yfilter.YFilter { return crypto.YFilter }

func (crypto *Crypto) SetFilter(yf yfilter.YFilter) { crypto.YFilter = yf }

func (crypto *Crypto) GetGoName(yname string) string {
    if yname == "sam" { return "Sam" }
    if yname == "Cisco-IOS-XR-crypto-ssh-cfg:ssh" { return "Ssh" }
    return ""
}

func (crypto *Crypto) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-sam-cfg:crypto"
}

func (crypto *Crypto) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sam" {
        return &crypto.Sam
    }
    if childYangName == "Cisco-IOS-XR-crypto-ssh-cfg:ssh" {
        return &crypto.Ssh
    }
    return nil
}

func (crypto *Crypto) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sam"] = &crypto.Sam
    children["Cisco-IOS-XR-crypto-ssh-cfg:ssh"] = &crypto.Ssh
    return children
}

func (crypto *Crypto) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (crypto *Crypto) GetBundleName() string { return "cisco_ios_xr" }

func (crypto *Crypto) GetYangName() string { return "crypto" }

func (crypto *Crypto) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (crypto *Crypto) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (crypto *Crypto) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (crypto *Crypto) SetParent(parent types.Entity) { crypto.parent = parent }

func (crypto *Crypto) GetParent() types.Entity { return crypto.parent }

func (crypto *Crypto) GetParentYangName() string { return "Cisco-IOS-XR-crypto-sam-cfg" }

// Crypto_Sam
// Software Authentication Manager (SAM) Config
type Crypto_Sam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set prompt interval at reboot time.
    PromptInterval Crypto_Sam_PromptInterval
}

func (sam *Crypto_Sam) GetFilter() yfilter.YFilter { return sam.YFilter }

func (sam *Crypto_Sam) SetFilter(yf yfilter.YFilter) { sam.YFilter = yf }

func (sam *Crypto_Sam) GetGoName(yname string) string {
    if yname == "prompt-interval" { return "PromptInterval" }
    return ""
}

func (sam *Crypto_Sam) GetSegmentPath() string {
    return "sam"
}

func (sam *Crypto_Sam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prompt-interval" {
        return &sam.PromptInterval
    }
    return nil
}

func (sam *Crypto_Sam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prompt-interval"] = &sam.PromptInterval
    return children
}

func (sam *Crypto_Sam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sam *Crypto_Sam) GetBundleName() string { return "cisco_ios_xr" }

func (sam *Crypto_Sam) GetYangName() string { return "sam" }

func (sam *Crypto_Sam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sam *Crypto_Sam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sam *Crypto_Sam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sam *Crypto_Sam) SetParent(parent types.Entity) { sam.parent = parent }

func (sam *Crypto_Sam) GetParent() types.Entity { return sam.parent }

func (sam *Crypto_Sam) GetParentYangName() string { return "crypto" }

// Crypto_Sam_PromptInterval
// Set prompt interval at reboot time
// This type is a presence type.
type Crypto_Sam_PromptInterval struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Respond to SAM prompt either Proceed/Terminate. The type is
    // CryptoSamAction. This attribute is mandatory.
    Action interface{}

    // Prompt time from 0 - 300 seconds. The type is interface{} with range:
    // 0..300. This attribute is mandatory. Units are second.
    PromptTime interface{}
}

func (promptInterval *Crypto_Sam_PromptInterval) GetFilter() yfilter.YFilter { return promptInterval.YFilter }

func (promptInterval *Crypto_Sam_PromptInterval) SetFilter(yf yfilter.YFilter) { promptInterval.YFilter = yf }

func (promptInterval *Crypto_Sam_PromptInterval) GetGoName(yname string) string {
    if yname == "action" { return "Action" }
    if yname == "prompt-time" { return "PromptTime" }
    return ""
}

func (promptInterval *Crypto_Sam_PromptInterval) GetSegmentPath() string {
    return "prompt-interval"
}

func (promptInterval *Crypto_Sam_PromptInterval) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (promptInterval *Crypto_Sam_PromptInterval) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (promptInterval *Crypto_Sam_PromptInterval) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action"] = promptInterval.Action
    leafs["prompt-time"] = promptInterval.PromptTime
    return leafs
}

func (promptInterval *Crypto_Sam_PromptInterval) GetBundleName() string { return "cisco_ios_xr" }

func (promptInterval *Crypto_Sam_PromptInterval) GetYangName() string { return "prompt-interval" }

func (promptInterval *Crypto_Sam_PromptInterval) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (promptInterval *Crypto_Sam_PromptInterval) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (promptInterval *Crypto_Sam_PromptInterval) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (promptInterval *Crypto_Sam_PromptInterval) SetParent(parent types.Entity) { promptInterval.parent = parent }

func (promptInterval *Crypto_Sam_PromptInterval) GetParent() types.Entity { return promptInterval.parent }

func (promptInterval *Crypto_Sam_PromptInterval) GetParentYangName() string { return "sam" }

// Crypto_Ssh
// Secure Shell configuration
type Crypto_Ssh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Provide SSH client service.
    Client Crypto_Ssh_Client

    // Provide SSH server service.
    Server Crypto_Ssh_Server
}

func (ssh *Crypto_Ssh) GetFilter() yfilter.YFilter { return ssh.YFilter }

func (ssh *Crypto_Ssh) SetFilter(yf yfilter.YFilter) { ssh.YFilter = yf }

func (ssh *Crypto_Ssh) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    if yname == "server" { return "Server" }
    return ""
}

func (ssh *Crypto_Ssh) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-ssh-cfg:ssh"
}

func (ssh *Crypto_Ssh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        return &ssh.Client
    }
    if childYangName == "server" {
        return &ssh.Server
    }
    return nil
}

func (ssh *Crypto_Ssh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client"] = &ssh.Client
    children["server"] = &ssh.Server
    return children
}

func (ssh *Crypto_Ssh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssh *Crypto_Ssh) GetBundleName() string { return "cisco_ios_xr" }

func (ssh *Crypto_Ssh) GetYangName() string { return "ssh" }

func (ssh *Crypto_Ssh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssh *Crypto_Ssh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssh *Crypto_Ssh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssh *Crypto_Ssh) SetParent(parent types.Entity) { ssh.parent = parent }

func (ssh *Crypto_Ssh) GetParent() types.Entity { return ssh.parent }

func (ssh *Crypto_Ssh) GetParentYangName() string { return "crypto" }

// Crypto_Ssh_Client
// Provide SSH client service
type Crypto_Ssh_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure client volume-based rekey for SSH. The type is interface{} with
    // range: 1024..4095. The default value is 1024.
    RekeyVolume interface{}

    // Filename - where to store known host file. The type is string.
    HostPublicKey interface{}

    // Source interface VRF for ssh client sessions. The type is string with
    // length: 1..32.
    ClientVrf interface{}

    // Configure client time-based rekey for SSH. The type is interface{} with
    // range: 30..1440. The default value is 60.
    RekeyTime interface{}

    // Source interface for ssh client sessions. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Cisco sshd DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}

    // clientenable.
    ClientEnable Crypto_Ssh_Client_ClientEnable
}

func (client *Crypto_Ssh_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *Crypto_Ssh_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *Crypto_Ssh_Client) GetGoName(yname string) string {
    if yname == "rekey-volume" { return "RekeyVolume" }
    if yname == "host-public-key" { return "HostPublicKey" }
    if yname == "client-vrf" { return "ClientVrf" }
    if yname == "rekey-time" { return "RekeyTime" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "dscp" { return "Dscp" }
    if yname == "client-enable" { return "ClientEnable" }
    return ""
}

func (client *Crypto_Ssh_Client) GetSegmentPath() string {
    return "client"
}

func (client *Crypto_Ssh_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-enable" {
        return &client.ClientEnable
    }
    return nil
}

func (client *Crypto_Ssh_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client-enable"] = &client.ClientEnable
    return children
}

func (client *Crypto_Ssh_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rekey-volume"] = client.RekeyVolume
    leafs["host-public-key"] = client.HostPublicKey
    leafs["client-vrf"] = client.ClientVrf
    leafs["rekey-time"] = client.RekeyTime
    leafs["source-interface"] = client.SourceInterface
    leafs["dscp"] = client.Dscp
    return leafs
}

func (client *Crypto_Ssh_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *Crypto_Ssh_Client) GetYangName() string { return "client" }

func (client *Crypto_Ssh_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *Crypto_Ssh_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *Crypto_Ssh_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *Crypto_Ssh_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *Crypto_Ssh_Client) GetParent() types.Entity { return client.parent }

func (client *Crypto_Ssh_Client) GetParentYangName() string { return "ssh" }

// Crypto_Ssh_Client_ClientEnable
// clientenable
type Crypto_Ssh_Client_ClientEnable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // clientcipher.
    ClientCipher Crypto_Ssh_Client_ClientEnable_ClientCipher
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetFilter() yfilter.YFilter { return clientEnable.YFilter }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) SetFilter(yf yfilter.YFilter) { clientEnable.YFilter = yf }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetGoName(yname string) string {
    if yname == "client-cipher" { return "ClientCipher" }
    return ""
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetSegmentPath() string {
    return "client-enable"
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client-cipher" {
        return &clientEnable.ClientCipher
    }
    return nil
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["client-cipher"] = &clientEnable.ClientCipher
    return children
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetBundleName() string { return "cisco_ios_xr" }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetYangName() string { return "client-enable" }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) SetParent(parent types.Entity) { clientEnable.parent = parent }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetParent() types.Entity { return clientEnable.parent }

func (clientEnable *Crypto_Ssh_Client_ClientEnable) GetParentYangName() string { return "client" }

// Crypto_Ssh_Client_ClientEnable_ClientCipher
// clientcipher
type Crypto_Ssh_Client_ClientEnable_ClientCipher struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable AES-CBC ciphers for client. The type is bool. The default value is
    // false.
    Aescbc interface{}
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetFilter() yfilter.YFilter { return clientCipher.YFilter }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) SetFilter(yf yfilter.YFilter) { clientCipher.YFilter = yf }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetGoName(yname string) string {
    if yname == "aescbc" { return "Aescbc" }
    return ""
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetSegmentPath() string {
    return "client-cipher"
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aescbc"] = clientCipher.Aescbc
    return leafs
}

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetBundleName() string { return "cisco_ios_xr" }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetYangName() string { return "client-cipher" }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) SetParent(parent types.Entity) { clientCipher.parent = parent }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetParent() types.Entity { return clientCipher.parent }

func (clientCipher *Crypto_Ssh_Client_ClientEnable_ClientCipher) GetParentYangName() string { return "client-enable" }

// Crypto_Ssh_Server
// Provide SSH server service
type Crypto_Ssh_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure volume-based rekey for SSH. The type is interface{} with range:
    // 1024..4095. The default value is 1024.
    RekeyVolume interface{}

    // Cisco sshd session-limit of service requests. The type is interface{} with
    // range: 1..1024.
    SessionLimit interface{}

    // port number on which ssh service to be started for netconf. The type is
    // interface{} with range: 1..65535. The default value is 830.
    Netconf interface{}

    // Cisco sshd force protocol version 2 only. The type is interface{}.
    V2 interface{}

    // Time Period in minutes, defalut 60. The type is interface{} with range:
    // 30..1440. The default value is 60.
    RekeyTime interface{}

    // Enable ssh server logging. The type is interface{}.
    Logging interface{}

    // Cisco sshd rate-limit of service requests. The type is interface{} with
    // range: 1..600. The default value is 60.
    RateLimit interface{}

    // Timeout value between 5-120 seconds defalut 30. The type is interface{}
    // with range: 5..120. The default value is 30.
    Timeout interface{}

    // Cisco sshd DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}

    // disable.
    Disable Crypto_Ssh_Server_Disable

    // enable.
    Enable Crypto_Ssh_Server_Enable

    // Cisco sshd VRF name.
    VrfTable Crypto_Ssh_Server_VrfTable

    // Capability.
    Capability Crypto_Ssh_Server_Capability

    // Cisco sshd Netconf VRF name.
    NetconfVrfTable Crypto_Ssh_Server_NetconfVrfTable
}

func (server *Crypto_Ssh_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Crypto_Ssh_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Crypto_Ssh_Server) GetGoName(yname string) string {
    if yname == "rekey-volume" { return "RekeyVolume" }
    if yname == "session-limit" { return "SessionLimit" }
    if yname == "netconf" { return "Netconf" }
    if yname == "v2" { return "V2" }
    if yname == "rekey-time" { return "RekeyTime" }
    if yname == "logging" { return "Logging" }
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "timeout" { return "Timeout" }
    if yname == "dscp" { return "Dscp" }
    if yname == "disable" { return "Disable" }
    if yname == "enable" { return "Enable" }
    if yname == "vrf-table" { return "VrfTable" }
    if yname == "capability" { return "Capability" }
    if yname == "netconf-vrf-table" { return "NetconfVrfTable" }
    return ""
}

func (server *Crypto_Ssh_Server) GetSegmentPath() string {
    return "server"
}

func (server *Crypto_Ssh_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "disable" {
        return &server.Disable
    }
    if childYangName == "enable" {
        return &server.Enable
    }
    if childYangName == "vrf-table" {
        return &server.VrfTable
    }
    if childYangName == "capability" {
        return &server.Capability
    }
    if childYangName == "netconf-vrf-table" {
        return &server.NetconfVrfTable
    }
    return nil
}

func (server *Crypto_Ssh_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["disable"] = &server.Disable
    children["enable"] = &server.Enable
    children["vrf-table"] = &server.VrfTable
    children["capability"] = &server.Capability
    children["netconf-vrf-table"] = &server.NetconfVrfTable
    return children
}

func (server *Crypto_Ssh_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rekey-volume"] = server.RekeyVolume
    leafs["session-limit"] = server.SessionLimit
    leafs["netconf"] = server.Netconf
    leafs["v2"] = server.V2
    leafs["rekey-time"] = server.RekeyTime
    leafs["logging"] = server.Logging
    leafs["rate-limit"] = server.RateLimit
    leafs["timeout"] = server.Timeout
    leafs["dscp"] = server.Dscp
    return leafs
}

func (server *Crypto_Ssh_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Crypto_Ssh_Server) GetYangName() string { return "server" }

func (server *Crypto_Ssh_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Crypto_Ssh_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Crypto_Ssh_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Crypto_Ssh_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Crypto_Ssh_Server) GetParent() types.Entity { return server.parent }

func (server *Crypto_Ssh_Server) GetParentYangName() string { return "ssh" }

// Crypto_Ssh_Server_Disable
// disable
type Crypto_Ssh_Server_Disable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // hmac.
    Hmac Crypto_Ssh_Server_Disable_Hmac
}

func (disable *Crypto_Ssh_Server_Disable) GetFilter() yfilter.YFilter { return disable.YFilter }

func (disable *Crypto_Ssh_Server_Disable) SetFilter(yf yfilter.YFilter) { disable.YFilter = yf }

func (disable *Crypto_Ssh_Server_Disable) GetGoName(yname string) string {
    if yname == "hmac" { return "Hmac" }
    return ""
}

func (disable *Crypto_Ssh_Server_Disable) GetSegmentPath() string {
    return "disable"
}

func (disable *Crypto_Ssh_Server_Disable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hmac" {
        return &disable.Hmac
    }
    return nil
}

func (disable *Crypto_Ssh_Server_Disable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hmac"] = &disable.Hmac
    return children
}

func (disable *Crypto_Ssh_Server_Disable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (disable *Crypto_Ssh_Server_Disable) GetBundleName() string { return "cisco_ios_xr" }

func (disable *Crypto_Ssh_Server_Disable) GetYangName() string { return "disable" }

func (disable *Crypto_Ssh_Server_Disable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (disable *Crypto_Ssh_Server_Disable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (disable *Crypto_Ssh_Server_Disable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (disable *Crypto_Ssh_Server_Disable) SetParent(parent types.Entity) { disable.parent = parent }

func (disable *Crypto_Ssh_Server_Disable) GetParent() types.Entity { return disable.parent }

func (disable *Crypto_Ssh_Server_Disable) GetParentYangName() string { return "server" }

// Crypto_Ssh_Server_Disable_Hmac
// hmac
type Crypto_Ssh_Server_Disable_Hmac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable Hmac-sha2-512 negotiation. The type is bool. The default value is
    // false.
    HmacSha512 interface{}
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetFilter() yfilter.YFilter { return hmac.YFilter }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) SetFilter(yf yfilter.YFilter) { hmac.YFilter = yf }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetGoName(yname string) string {
    if yname == "hmac-sha512" { return "HmacSha512" }
    return ""
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetSegmentPath() string {
    return "hmac"
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha512"] = hmac.HmacSha512
    return leafs
}

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetBundleName() string { return "cisco_ios_xr" }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetYangName() string { return "hmac" }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) SetParent(parent types.Entity) { hmac.parent = parent }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetParent() types.Entity { return hmac.parent }

func (hmac *Crypto_Ssh_Server_Disable_Hmac) GetParentYangName() string { return "disable" }

// Crypto_Ssh_Server_Enable
// enable
type Crypto_Ssh_Server_Enable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // cipher.
    Cipher Crypto_Ssh_Server_Enable_Cipher
}

func (enable *Crypto_Ssh_Server_Enable) GetFilter() yfilter.YFilter { return enable.YFilter }

func (enable *Crypto_Ssh_Server_Enable) SetFilter(yf yfilter.YFilter) { enable.YFilter = yf }

func (enable *Crypto_Ssh_Server_Enable) GetGoName(yname string) string {
    if yname == "cipher" { return "Cipher" }
    return ""
}

func (enable *Crypto_Ssh_Server_Enable) GetSegmentPath() string {
    return "enable"
}

func (enable *Crypto_Ssh_Server_Enable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipher" {
        return &enable.Cipher
    }
    return nil
}

func (enable *Crypto_Ssh_Server_Enable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipher"] = &enable.Cipher
    return children
}

func (enable *Crypto_Ssh_Server_Enable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (enable *Crypto_Ssh_Server_Enable) GetBundleName() string { return "cisco_ios_xr" }

func (enable *Crypto_Ssh_Server_Enable) GetYangName() string { return "enable" }

func (enable *Crypto_Ssh_Server_Enable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (enable *Crypto_Ssh_Server_Enable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (enable *Crypto_Ssh_Server_Enable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (enable *Crypto_Ssh_Server_Enable) SetParent(parent types.Entity) { enable.parent = parent }

func (enable *Crypto_Ssh_Server_Enable) GetParent() types.Entity { return enable.parent }

func (enable *Crypto_Ssh_Server_Enable) GetParentYangName() string { return "server" }

// Crypto_Ssh_Server_Enable_Cipher
// cipher
type Crypto_Ssh_Server_Enable_Cipher struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable AES-CBC ciphers. The type is bool. The default value is false.
    Aescbc interface{}
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetFilter() yfilter.YFilter { return cipher.YFilter }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) SetFilter(yf yfilter.YFilter) { cipher.YFilter = yf }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetGoName(yname string) string {
    if yname == "aescbc" { return "Aescbc" }
    return ""
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetSegmentPath() string {
    return "cipher"
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aescbc"] = cipher.Aescbc
    return leafs
}

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetBundleName() string { return "cisco_ios_xr" }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetYangName() string { return "cipher" }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) SetParent(parent types.Entity) { cipher.parent = parent }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetParent() types.Entity { return cipher.parent }

func (cipher *Crypto_Ssh_Server_Enable_Cipher) GetParentYangName() string { return "enable" }

// Crypto_Ssh_Server_VrfTable
// Cisco sshd VRF name
type Crypto_Ssh_Server_VrfTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Crypto_Ssh_Server_VrfTable_Vrf.
    Vrf []Crypto_Ssh_Server_VrfTable_Vrf
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetFilter() yfilter.YFilter { return vrfTable.YFilter }

func (vrfTable *Crypto_Ssh_Server_VrfTable) SetFilter(yf yfilter.YFilter) { vrfTable.YFilter = yf }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetSegmentPath() string {
    return "vrf-table"
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfTable.Vrf {
            if vrfTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Crypto_Ssh_Server_VrfTable_Vrf{}
        vrfTable.Vrf = append(vrfTable.Vrf, child)
        return &vrfTable.Vrf[len(vrfTable.Vrf)-1]
    }
    return nil
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfTable.Vrf {
        children[vrfTable.Vrf[i].GetSegmentPath()] = &vrfTable.Vrf[i]
    }
    return children
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetBundleName() string { return "cisco_ios_xr" }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetYangName() string { return "vrf-table" }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfTable *Crypto_Ssh_Server_VrfTable) SetParent(parent types.Entity) { vrfTable.parent = parent }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetParent() types.Entity { return vrfTable.parent }

func (vrfTable *Crypto_Ssh_Server_VrfTable) GetParentYangName() string { return "server" }

// Crypto_Ssh_Server_VrfTable_Vrf
// Enter VRF name
type Crypto_Ssh_Server_VrfTable_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Enter VRF name. The type is string with length:
    // 1..32.
    VrfName interface{}

    // Enable to use VRF. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // SSH v4 access-list name. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // SSH v6 access-list name. The type is string with length: 1..32.
    Ipv6AccessList interface{}
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "ipv4-access-list" { return "Ipv4AccessList" }
    if yname == "ipv6-access-list" { return "Ipv6AccessList" }
    return ""
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    leafs["ipv4-access-list"] = vrf.Ipv4AccessList
    leafs["ipv6-access-list"] = vrf.Ipv6AccessList
    return leafs
}

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetYangName() string { return "vrf" }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Crypto_Ssh_Server_VrfTable_Vrf) GetParentYangName() string { return "vrf-table" }

// Crypto_Ssh_Server_Capability
// Capability
type Crypto_Ssh_Server_Capability struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Netconf-XML stack on port 22. The type is bool. The default value is
    // false.
    NetconfXml interface{}
}

func (capability *Crypto_Ssh_Server_Capability) GetFilter() yfilter.YFilter { return capability.YFilter }

func (capability *Crypto_Ssh_Server_Capability) SetFilter(yf yfilter.YFilter) { capability.YFilter = yf }

func (capability *Crypto_Ssh_Server_Capability) GetGoName(yname string) string {
    if yname == "netconf-xml" { return "NetconfXml" }
    return ""
}

func (capability *Crypto_Ssh_Server_Capability) GetSegmentPath() string {
    return "capability"
}

func (capability *Crypto_Ssh_Server_Capability) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capability *Crypto_Ssh_Server_Capability) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capability *Crypto_Ssh_Server_Capability) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["netconf-xml"] = capability.NetconfXml
    return leafs
}

func (capability *Crypto_Ssh_Server_Capability) GetBundleName() string { return "cisco_ios_xr" }

func (capability *Crypto_Ssh_Server_Capability) GetYangName() string { return "capability" }

func (capability *Crypto_Ssh_Server_Capability) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capability *Crypto_Ssh_Server_Capability) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capability *Crypto_Ssh_Server_Capability) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capability *Crypto_Ssh_Server_Capability) SetParent(parent types.Entity) { capability.parent = parent }

func (capability *Crypto_Ssh_Server_Capability) GetParent() types.Entity { return capability.parent }

func (capability *Crypto_Ssh_Server_Capability) GetParentYangName() string { return "server" }

// Crypto_Ssh_Server_NetconfVrfTable
// Cisco sshd Netconf VRF name
type Crypto_Ssh_Server_NetconfVrfTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Crypto_Ssh_Server_NetconfVrfTable_Vrf.
    Vrf []Crypto_Ssh_Server_NetconfVrfTable_Vrf
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetFilter() yfilter.YFilter { return netconfVrfTable.YFilter }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) SetFilter(yf yfilter.YFilter) { netconfVrfTable.YFilter = yf }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetSegmentPath() string {
    return "netconf-vrf-table"
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range netconfVrfTable.Vrf {
            if netconfVrfTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Crypto_Ssh_Server_NetconfVrfTable_Vrf{}
        netconfVrfTable.Vrf = append(netconfVrfTable.Vrf, child)
        return &netconfVrfTable.Vrf[len(netconfVrfTable.Vrf)-1]
    }
    return nil
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range netconfVrfTable.Vrf {
        children[netconfVrfTable.Vrf[i].GetSegmentPath()] = &netconfVrfTable.Vrf[i]
    }
    return children
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetBundleName() string { return "cisco_ios_xr" }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetYangName() string { return "netconf-vrf-table" }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) SetParent(parent types.Entity) { netconfVrfTable.parent = parent }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetParent() types.Entity { return netconfVrfTable.parent }

func (netconfVrfTable *Crypto_Ssh_Server_NetconfVrfTable) GetParentYangName() string { return "server" }

// Crypto_Ssh_Server_NetconfVrfTable_Vrf
// Enter VRF name
type Crypto_Ssh_Server_NetconfVrfTable_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Enter VRF name. The type is string with length:
    // 1..32.
    VrfName interface{}

    // Enable to use VRF. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // SSH v4 access-list name. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // SSH v6 access-list name. The type is string with length: 1..32.
    Ipv6AccessList interface{}
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "ipv4-access-list" { return "Ipv4AccessList" }
    if yname == "ipv6-access-list" { return "Ipv6AccessList" }
    return ""
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    leafs["ipv4-access-list"] = vrf.Ipv4AccessList
    leafs["ipv6-access-list"] = vrf.Ipv6AccessList
    return leafs
}

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetYangName() string { return "vrf" }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Crypto_Ssh_Server_NetconfVrfTable_Vrf) GetParentYangName() string { return "netconf-vrf-table" }

