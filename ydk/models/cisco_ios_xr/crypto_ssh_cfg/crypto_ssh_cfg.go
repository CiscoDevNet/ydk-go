// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-ssh package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ssh: Secure Shell configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_ssh_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_ssh_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-ssh-cfg ssh}", reflect.TypeOf(Ssh{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-ssh-cfg:ssh", reflect.TypeOf(Ssh{}))
}

// Ssh
// Secure Shell configuration
type Ssh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Provide SSH client service.
    Client Ssh_Client

    // Provide SSH server service.
    Server Ssh_Server
}

func (ssh *Ssh) GetEntityData() *types.CommonEntityData {
    ssh.EntityData.YFilter = ssh.YFilter
    ssh.EntityData.YangName = "ssh"
    ssh.EntityData.BundleName = "cisco_ios_xr"
    ssh.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-ssh-cfg"
    ssh.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-ssh-cfg:ssh"
    ssh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssh.EntityData.Children = make(map[string]types.YChild)
    ssh.EntityData.Children["client"] = types.YChild{"Client", &ssh.Client}
    ssh.EntityData.Children["server"] = types.YChild{"Server", &ssh.Server}
    ssh.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ssh.EntityData)
}

// Ssh_Client
// Provide SSH client service
type Ssh_Client struct {
    EntityData types.CommonEntityData
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
    // range: 30..1440. Units are minute. The default value is 60.
    RekeyTime interface{}

    // Source interface for ssh client sessions. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    SourceInterface interface{}

    // Cisco sshd DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}

    // Cisco ssh algorithms.
    ClientAlgo Ssh_Client_ClientAlgo

    // clientenable.
    ClientEnable Ssh_Client_ClientEnable
}

func (client *Ssh_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "ssh"
    client.EntityData.SegmentPath = "client"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Children["client-algo"] = types.YChild{"ClientAlgo", &client.ClientAlgo}
    client.EntityData.Children["client-enable"] = types.YChild{"ClientEnable", &client.ClientEnable}
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["rekey-volume"] = types.YLeaf{"RekeyVolume", client.RekeyVolume}
    client.EntityData.Leafs["host-public-key"] = types.YLeaf{"HostPublicKey", client.HostPublicKey}
    client.EntityData.Leafs["client-vrf"] = types.YLeaf{"ClientVrf", client.ClientVrf}
    client.EntityData.Leafs["rekey-time"] = types.YLeaf{"RekeyTime", client.RekeyTime}
    client.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", client.SourceInterface}
    client.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", client.Dscp}
    return &(client.EntityData)
}

// Ssh_Client_ClientAlgo
// Cisco ssh algorithms
type Ssh_Client_ClientAlgo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key exchange algorithm.
    KeyExchange Ssh_Client_ClientAlgo_KeyExchange
}

func (clientAlgo *Ssh_Client_ClientAlgo) GetEntityData() *types.CommonEntityData {
    clientAlgo.EntityData.YFilter = clientAlgo.YFilter
    clientAlgo.EntityData.YangName = "client-algo"
    clientAlgo.EntityData.BundleName = "cisco_ios_xr"
    clientAlgo.EntityData.ParentYangName = "client"
    clientAlgo.EntityData.SegmentPath = "client-algo"
    clientAlgo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientAlgo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientAlgo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientAlgo.EntityData.Children = make(map[string]types.YChild)
    clientAlgo.EntityData.Children["key-exchange"] = types.YChild{"KeyExchange", &clientAlgo.KeyExchange}
    clientAlgo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clientAlgo.EntityData)
}

// Ssh_Client_ClientAlgo_KeyExchange
// Key exchange algorithm
// This type is a presence type.
type Ssh_Client_ClientAlgo_KeyExchange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key exchange algorithm. The type is string with length: 1..32. This
    // attribute is mandatory.
    KexAlgo1St interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo2Nd interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo3Rd interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo4Th interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo5Th interface{}
}

func (keyExchange *Ssh_Client_ClientAlgo_KeyExchange) GetEntityData() *types.CommonEntityData {
    keyExchange.EntityData.YFilter = keyExchange.YFilter
    keyExchange.EntityData.YangName = "key-exchange"
    keyExchange.EntityData.BundleName = "cisco_ios_xr"
    keyExchange.EntityData.ParentYangName = "client-algo"
    keyExchange.EntityData.SegmentPath = "key-exchange"
    keyExchange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyExchange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyExchange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyExchange.EntityData.Children = make(map[string]types.YChild)
    keyExchange.EntityData.Leafs = make(map[string]types.YLeaf)
    keyExchange.EntityData.Leafs["kex-algo1st"] = types.YLeaf{"KexAlgo1St", keyExchange.KexAlgo1St}
    keyExchange.EntityData.Leafs["kex-algo2nd"] = types.YLeaf{"KexAlgo2Nd", keyExchange.KexAlgo2Nd}
    keyExchange.EntityData.Leafs["kex-algo3rd"] = types.YLeaf{"KexAlgo3Rd", keyExchange.KexAlgo3Rd}
    keyExchange.EntityData.Leafs["kex-algo4th"] = types.YLeaf{"KexAlgo4Th", keyExchange.KexAlgo4Th}
    keyExchange.EntityData.Leafs["kex-algo5th"] = types.YLeaf{"KexAlgo5Th", keyExchange.KexAlgo5Th}
    return &(keyExchange.EntityData)
}

// Ssh_Client_ClientEnable
// clientenable
type Ssh_Client_ClientEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // clientcipher.
    ClientCipher Ssh_Client_ClientEnable_ClientCipher
}

func (clientEnable *Ssh_Client_ClientEnable) GetEntityData() *types.CommonEntityData {
    clientEnable.EntityData.YFilter = clientEnable.YFilter
    clientEnable.EntityData.YangName = "client-enable"
    clientEnable.EntityData.BundleName = "cisco_ios_xr"
    clientEnable.EntityData.ParentYangName = "client"
    clientEnable.EntityData.SegmentPath = "client-enable"
    clientEnable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientEnable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientEnable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientEnable.EntityData.Children = make(map[string]types.YChild)
    clientEnable.EntityData.Children["client-cipher"] = types.YChild{"ClientCipher", &clientEnable.ClientCipher}
    clientEnable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clientEnable.EntityData)
}

// Ssh_Client_ClientEnable_ClientCipher
// clientcipher
type Ssh_Client_ClientEnable_ClientCipher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable AES-CBC ciphers for client. The type is bool. The default value is
    // false.
    Aescbc interface{}
}

func (clientCipher *Ssh_Client_ClientEnable_ClientCipher) GetEntityData() *types.CommonEntityData {
    clientCipher.EntityData.YFilter = clientCipher.YFilter
    clientCipher.EntityData.YangName = "client-cipher"
    clientCipher.EntityData.BundleName = "cisco_ios_xr"
    clientCipher.EntityData.ParentYangName = "client-enable"
    clientCipher.EntityData.SegmentPath = "client-cipher"
    clientCipher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientCipher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientCipher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientCipher.EntityData.Children = make(map[string]types.YChild)
    clientCipher.EntityData.Leafs = make(map[string]types.YLeaf)
    clientCipher.EntityData.Leafs["aescbc"] = types.YLeaf{"Aescbc", clientCipher.Aescbc}
    return &(clientCipher.EntityData)
}

// Ssh_Server
// Provide SSH server service
type Ssh_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure volume-based rekey for SSH. The type is interface{} with range:
    // 1024..4095. The default value is 1024.
    RekeyVolume interface{}

    // Cisco sshd session-limit of service requests. The type is interface{} with
    // range: 1..110.
    SessionLimit interface{}

    // port number on which ssh service to be started for netconf. The type is
    // interface{} with range: 1..65535. The default value is 830.
    Netconf interface{}

    // Cisco sshd force protocol version 2 only. The type is interface{}.
    V2 interface{}

    // Time Period in minutes, defalut 60. The type is interface{} with range:
    // 30..1440. Units are minute. The default value is 60.
    RekeyTime interface{}

    // Enable ssh server logging. The type is interface{}.
    Logging interface{}

    // Cisco sshd rate-limit of service requests. The type is interface{} with
    // range: 1..600. The default value is 60.
    RateLimit interface{}

    // Timeout value between 5-120 seconds defalut 30. The type is interface{}
    // with range: 5..120. Units are second. The default value is 30.
    Timeout interface{}

    // Cisco sshd DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}

    // disable.
    Disable Ssh_Server_Disable

    // enable.
    Enable Ssh_Server_Enable

    // Cisco sshd VRF name.
    VrfTable Ssh_Server_VrfTable

    // Cisco ssh algorithms.
    ServerAlgo Ssh_Server_ServerAlgo

    // Capability.
    Capability Ssh_Server_Capability

    // Cisco sshd Netconf VRF name.
    NetconfVrfTable Ssh_Server_NetconfVrfTable
}

func (server *Ssh_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "ssh"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Children["disable"] = types.YChild{"Disable", &server.Disable}
    server.EntityData.Children["enable"] = types.YChild{"Enable", &server.Enable}
    server.EntityData.Children["vrf-table"] = types.YChild{"VrfTable", &server.VrfTable}
    server.EntityData.Children["server-algo"] = types.YChild{"ServerAlgo", &server.ServerAlgo}
    server.EntityData.Children["capability"] = types.YChild{"Capability", &server.Capability}
    server.EntityData.Children["netconf-vrf-table"] = types.YChild{"NetconfVrfTable", &server.NetconfVrfTable}
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["rekey-volume"] = types.YLeaf{"RekeyVolume", server.RekeyVolume}
    server.EntityData.Leafs["session-limit"] = types.YLeaf{"SessionLimit", server.SessionLimit}
    server.EntityData.Leafs["netconf"] = types.YLeaf{"Netconf", server.Netconf}
    server.EntityData.Leafs["v2"] = types.YLeaf{"V2", server.V2}
    server.EntityData.Leafs["rekey-time"] = types.YLeaf{"RekeyTime", server.RekeyTime}
    server.EntityData.Leafs["logging"] = types.YLeaf{"Logging", server.Logging}
    server.EntityData.Leafs["rate-limit"] = types.YLeaf{"RateLimit", server.RateLimit}
    server.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", server.Timeout}
    server.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", server.Dscp}
    return &(server.EntityData)
}

// Ssh_Server_Disable
// disable
type Ssh_Server_Disable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // hmac.
    Hmac Ssh_Server_Disable_Hmac
}

func (disable *Ssh_Server_Disable) GetEntityData() *types.CommonEntityData {
    disable.EntityData.YFilter = disable.YFilter
    disable.EntityData.YangName = "disable"
    disable.EntityData.BundleName = "cisco_ios_xr"
    disable.EntityData.ParentYangName = "server"
    disable.EntityData.SegmentPath = "disable"
    disable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disable.EntityData.Children = make(map[string]types.YChild)
    disable.EntityData.Children["hmac"] = types.YChild{"Hmac", &disable.Hmac}
    disable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(disable.EntityData)
}

// Ssh_Server_Disable_Hmac
// hmac
type Ssh_Server_Disable_Hmac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable Hmac-sha2-512 negotiation. The type is bool. The default value is
    // false.
    HmacSha512 interface{}
}

func (hmac *Ssh_Server_Disable_Hmac) GetEntityData() *types.CommonEntityData {
    hmac.EntityData.YFilter = hmac.YFilter
    hmac.EntityData.YangName = "hmac"
    hmac.EntityData.BundleName = "cisco_ios_xr"
    hmac.EntityData.ParentYangName = "disable"
    hmac.EntityData.SegmentPath = "hmac"
    hmac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hmac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hmac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hmac.EntityData.Children = make(map[string]types.YChild)
    hmac.EntityData.Leafs = make(map[string]types.YLeaf)
    hmac.EntityData.Leafs["hmac-sha512"] = types.YLeaf{"HmacSha512", hmac.HmacSha512}
    return &(hmac.EntityData)
}

// Ssh_Server_Enable
// enable
type Ssh_Server_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // cipher.
    Cipher Ssh_Server_Enable_Cipher
}

func (enable *Ssh_Server_Enable) GetEntityData() *types.CommonEntityData {
    enable.EntityData.YFilter = enable.YFilter
    enable.EntityData.YangName = "enable"
    enable.EntityData.BundleName = "cisco_ios_xr"
    enable.EntityData.ParentYangName = "server"
    enable.EntityData.SegmentPath = "enable"
    enable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    enable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    enable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    enable.EntityData.Children = make(map[string]types.YChild)
    enable.EntityData.Children["cipher"] = types.YChild{"Cipher", &enable.Cipher}
    enable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(enable.EntityData)
}

// Ssh_Server_Enable_Cipher
// cipher
type Ssh_Server_Enable_Cipher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable AES-CBC ciphers. The type is bool. The default value is false.
    Aescbc interface{}
}

func (cipher *Ssh_Server_Enable_Cipher) GetEntityData() *types.CommonEntityData {
    cipher.EntityData.YFilter = cipher.YFilter
    cipher.EntityData.YangName = "cipher"
    cipher.EntityData.BundleName = "cisco_ios_xr"
    cipher.EntityData.ParentYangName = "enable"
    cipher.EntityData.SegmentPath = "cipher"
    cipher.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cipher.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cipher.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cipher.EntityData.Children = make(map[string]types.YChild)
    cipher.EntityData.Leafs = make(map[string]types.YLeaf)
    cipher.EntityData.Leafs["aescbc"] = types.YLeaf{"Aescbc", cipher.Aescbc}
    return &(cipher.EntityData)
}

// Ssh_Server_VrfTable
// Cisco sshd VRF name
type Ssh_Server_VrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Ssh_Server_VrfTable_Vrf.
    Vrf []Ssh_Server_VrfTable_Vrf
}

func (vrfTable *Ssh_Server_VrfTable) GetEntityData() *types.CommonEntityData {
    vrfTable.EntityData.YFilter = vrfTable.YFilter
    vrfTable.EntityData.YangName = "vrf-table"
    vrfTable.EntityData.BundleName = "cisco_ios_xr"
    vrfTable.EntityData.ParentYangName = "server"
    vrfTable.EntityData.SegmentPath = "vrf-table"
    vrfTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfTable.EntityData.Children = make(map[string]types.YChild)
    vrfTable.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfTable.Vrf {
        vrfTable.EntityData.Children[types.GetSegmentPath(&vrfTable.Vrf[i])] = types.YChild{"Vrf", &vrfTable.Vrf[i]}
    }
    vrfTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfTable.EntityData)
}

// Ssh_Server_VrfTable_Vrf
// Enter VRF name
type Ssh_Server_VrfTable_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *Ssh_Server_VrfTable_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrf-table"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vrf.Enable}
    vrf.EntityData.Leafs["ipv4-access-list"] = types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList}
    vrf.EntityData.Leafs["ipv6-access-list"] = types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList}
    return &(vrf.EntityData)
}

// Ssh_Server_ServerAlgo
// Cisco ssh algorithms
type Ssh_Server_ServerAlgo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key exchange algorithm.
    KeyExchange Ssh_Server_ServerAlgo_KeyExchange
}

func (serverAlgo *Ssh_Server_ServerAlgo) GetEntityData() *types.CommonEntityData {
    serverAlgo.EntityData.YFilter = serverAlgo.YFilter
    serverAlgo.EntityData.YangName = "server-algo"
    serverAlgo.EntityData.BundleName = "cisco_ios_xr"
    serverAlgo.EntityData.ParentYangName = "server"
    serverAlgo.EntityData.SegmentPath = "server-algo"
    serverAlgo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverAlgo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverAlgo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverAlgo.EntityData.Children = make(map[string]types.YChild)
    serverAlgo.EntityData.Children["key-exchange"] = types.YChild{"KeyExchange", &serverAlgo.KeyExchange}
    serverAlgo.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serverAlgo.EntityData)
}

// Ssh_Server_ServerAlgo_KeyExchange
// Key exchange algorithm
// This type is a presence type.
type Ssh_Server_ServerAlgo_KeyExchange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key exchange algorithm. The type is string with length: 1..32. This
    // attribute is mandatory.
    KexAlgo1St interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo2Nd interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo3Rd interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo4Th interface{}

    // key exchange algorithm. The type is string with length: 1..32.
    KexAlgo5Th interface{}
}

func (keyExchange *Ssh_Server_ServerAlgo_KeyExchange) GetEntityData() *types.CommonEntityData {
    keyExchange.EntityData.YFilter = keyExchange.YFilter
    keyExchange.EntityData.YangName = "key-exchange"
    keyExchange.EntityData.BundleName = "cisco_ios_xr"
    keyExchange.EntityData.ParentYangName = "server-algo"
    keyExchange.EntityData.SegmentPath = "key-exchange"
    keyExchange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyExchange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyExchange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyExchange.EntityData.Children = make(map[string]types.YChild)
    keyExchange.EntityData.Leafs = make(map[string]types.YLeaf)
    keyExchange.EntityData.Leafs["kex-algo1st"] = types.YLeaf{"KexAlgo1St", keyExchange.KexAlgo1St}
    keyExchange.EntityData.Leafs["kex-algo2nd"] = types.YLeaf{"KexAlgo2Nd", keyExchange.KexAlgo2Nd}
    keyExchange.EntityData.Leafs["kex-algo3rd"] = types.YLeaf{"KexAlgo3Rd", keyExchange.KexAlgo3Rd}
    keyExchange.EntityData.Leafs["kex-algo4th"] = types.YLeaf{"KexAlgo4Th", keyExchange.KexAlgo4Th}
    keyExchange.EntityData.Leafs["kex-algo5th"] = types.YLeaf{"KexAlgo5Th", keyExchange.KexAlgo5Th}
    return &(keyExchange.EntityData)
}

// Ssh_Server_Capability
// Capability
type Ssh_Server_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Netconf-XML stack on port 22. The type is bool. The default value is
    // false.
    NetconfXml interface{}
}

func (capability *Ssh_Server_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "cisco_ios_xr"
    capability.EntityData.ParentYangName = "server"
    capability.EntityData.SegmentPath = "capability"
    capability.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capability.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capability.EntityData.Children = make(map[string]types.YChild)
    capability.EntityData.Leafs = make(map[string]types.YLeaf)
    capability.EntityData.Leafs["netconf-xml"] = types.YLeaf{"NetconfXml", capability.NetconfXml}
    return &(capability.EntityData)
}

// Ssh_Server_NetconfVrfTable
// Cisco sshd Netconf VRF name
type Ssh_Server_NetconfVrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Ssh_Server_NetconfVrfTable_Vrf.
    Vrf []Ssh_Server_NetconfVrfTable_Vrf
}

func (netconfVrfTable *Ssh_Server_NetconfVrfTable) GetEntityData() *types.CommonEntityData {
    netconfVrfTable.EntityData.YFilter = netconfVrfTable.YFilter
    netconfVrfTable.EntityData.YangName = "netconf-vrf-table"
    netconfVrfTable.EntityData.BundleName = "cisco_ios_xr"
    netconfVrfTable.EntityData.ParentYangName = "server"
    netconfVrfTable.EntityData.SegmentPath = "netconf-vrf-table"
    netconfVrfTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfVrfTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfVrfTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfVrfTable.EntityData.Children = make(map[string]types.YChild)
    netconfVrfTable.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range netconfVrfTable.Vrf {
        netconfVrfTable.EntityData.Children[types.GetSegmentPath(&netconfVrfTable.Vrf[i])] = types.YChild{"Vrf", &netconfVrfTable.Vrf[i]}
    }
    netconfVrfTable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(netconfVrfTable.EntityData)
}

// Ssh_Server_NetconfVrfTable_Vrf
// Enter VRF name
type Ssh_Server_NetconfVrfTable_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *Ssh_Server_NetconfVrfTable_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "netconf-vrf-table"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vrf.Enable}
    vrf.EntityData.Leafs["ipv4-access-list"] = types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList}
    vrf.EntityData.Leafs["ipv6-access-list"] = types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList}
    return &(vrf.EntityData)
}

