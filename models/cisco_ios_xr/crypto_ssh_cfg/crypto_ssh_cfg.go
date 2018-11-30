// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-ssh package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ssh: Secure Shell configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    ssh.EntityData.Children = types.NewOrderedMap()
    ssh.EntityData.Children.Append("client", types.YChild{"Client", &ssh.Client})
    ssh.EntityData.Children.Append("server", types.YChild{"Server", &ssh.Server})
    ssh.EntityData.Leafs = types.NewOrderedMap()

    ssh.EntityData.YListKeys = []string {}

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
    // [a-zA-Z0-9._/-]+.
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

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("client-algo", types.YChild{"ClientAlgo", &client.ClientAlgo})
    client.EntityData.Children.Append("client-enable", types.YChild{"ClientEnable", &client.ClientEnable})
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("rekey-volume", types.YLeaf{"RekeyVolume", client.RekeyVolume})
    client.EntityData.Leafs.Append("host-public-key", types.YLeaf{"HostPublicKey", client.HostPublicKey})
    client.EntityData.Leafs.Append("client-vrf", types.YLeaf{"ClientVrf", client.ClientVrf})
    client.EntityData.Leafs.Append("rekey-time", types.YLeaf{"RekeyTime", client.RekeyTime})
    client.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", client.SourceInterface})
    client.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", client.Dscp})

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// Ssh_Client_ClientAlgo
// Cisco ssh algorithms
type Ssh_Client_ClientAlgo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key exchange algorithm.
    KeyExchanges Ssh_Client_ClientAlgo_KeyExchanges
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

    clientAlgo.EntityData.Children = types.NewOrderedMap()
    clientAlgo.EntityData.Children.Append("key-exchanges", types.YChild{"KeyExchanges", &clientAlgo.KeyExchanges})
    clientAlgo.EntityData.Leafs = types.NewOrderedMap()

    clientAlgo.EntityData.YListKeys = []string {}

    return &(clientAlgo.EntityData)
}

// Ssh_Client_ClientAlgo_KeyExchanges
// Key exchange algorithm
type Ssh_Client_ClientAlgo_KeyExchanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key exchange algorithm. The type is slice of string with length: 1..32.
    KeyExchange []interface{}
}

func (keyExchanges *Ssh_Client_ClientAlgo_KeyExchanges) GetEntityData() *types.CommonEntityData {
    keyExchanges.EntityData.YFilter = keyExchanges.YFilter
    keyExchanges.EntityData.YangName = "key-exchanges"
    keyExchanges.EntityData.BundleName = "cisco_ios_xr"
    keyExchanges.EntityData.ParentYangName = "client-algo"
    keyExchanges.EntityData.SegmentPath = "key-exchanges"
    keyExchanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyExchanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyExchanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyExchanges.EntityData.Children = types.NewOrderedMap()
    keyExchanges.EntityData.Leafs = types.NewOrderedMap()
    keyExchanges.EntityData.Leafs.Append("key-exchange", types.YLeaf{"KeyExchange", keyExchanges.KeyExchange})

    keyExchanges.EntityData.YListKeys = []string {}

    return &(keyExchanges.EntityData)
}

// Ssh_Client_ClientEnable
// clientenable
type Ssh_Client_ClientEnable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable AES-CBC and 3DES_CBC for ssh client.
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

    clientEnable.EntityData.Children = types.NewOrderedMap()
    clientEnable.EntityData.Children.Append("client-cipher", types.YChild{"ClientCipher", &clientEnable.ClientCipher})
    clientEnable.EntityData.Leafs = types.NewOrderedMap()

    clientEnable.EntityData.YListKeys = []string {}

    return &(clientEnable.EntityData)
}

// Ssh_Client_ClientEnable_ClientCipher
// Enable AES-CBC and 3DES_CBC for ssh client
type Ssh_Client_ClientEnable_ClientCipher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable AES-CBC ciphers. The type is bool. The default value is false.
    AesCbc interface{}

    // Enable 3DES-CBC cipher. The type is bool. The default value is false.
    TripledesCbc interface{}
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

    clientCipher.EntityData.Children = types.NewOrderedMap()
    clientCipher.EntityData.Leafs = types.NewOrderedMap()
    clientCipher.EntityData.Leafs.Append("aes-cbc", types.YLeaf{"AesCbc", clientCipher.AesCbc})
    clientCipher.EntityData.Leafs.Append("tripledes-cbc", types.YLeaf{"TripledesCbc", clientCipher.TripledesCbc})

    clientCipher.EntityData.YListKeys = []string {}

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

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("disable", types.YChild{"Disable", &server.Disable})
    server.EntityData.Children.Append("enable", types.YChild{"Enable", &server.Enable})
    server.EntityData.Children.Append("vrf-table", types.YChild{"VrfTable", &server.VrfTable})
    server.EntityData.Children.Append("server-algo", types.YChild{"ServerAlgo", &server.ServerAlgo})
    server.EntityData.Children.Append("capability", types.YChild{"Capability", &server.Capability})
    server.EntityData.Children.Append("netconf-vrf-table", types.YChild{"NetconfVrfTable", &server.NetconfVrfTable})
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("rekey-volume", types.YLeaf{"RekeyVolume", server.RekeyVolume})
    server.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", server.SessionLimit})
    server.EntityData.Leafs.Append("netconf", types.YLeaf{"Netconf", server.Netconf})
    server.EntityData.Leafs.Append("v2", types.YLeaf{"V2", server.V2})
    server.EntityData.Leafs.Append("rekey-time", types.YLeaf{"RekeyTime", server.RekeyTime})
    server.EntityData.Leafs.Append("logging", types.YLeaf{"Logging", server.Logging})
    server.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", server.RateLimit})
    server.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", server.Timeout})
    server.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", server.Dscp})

    server.EntityData.YListKeys = []string {}

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

    disable.EntityData.Children = types.NewOrderedMap()
    disable.EntityData.Children.Append("hmac", types.YChild{"Hmac", &disable.Hmac})
    disable.EntityData.Leafs = types.NewOrderedMap()

    disable.EntityData.YListKeys = []string {}

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

    hmac.EntityData.Children = types.NewOrderedMap()
    hmac.EntityData.Leafs = types.NewOrderedMap()
    hmac.EntityData.Leafs.Append("hmac-sha512", types.YLeaf{"HmacSha512", hmac.HmacSha512})

    hmac.EntityData.YListKeys = []string {}

    return &(hmac.EntityData)
}

// Ssh_Server_Enable
// enable
type Ssh_Server_Enable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable AES-CBC and 3DES-CBC ciphers.
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

    enable.EntityData.Children = types.NewOrderedMap()
    enable.EntityData.Children.Append("cipher", types.YChild{"Cipher", &enable.Cipher})
    enable.EntityData.Leafs = types.NewOrderedMap()

    enable.EntityData.YListKeys = []string {}

    return &(enable.EntityData)
}

// Ssh_Server_Enable_Cipher
// Enable AES-CBC and 3DES-CBC ciphers
type Ssh_Server_Enable_Cipher struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable aes-cbc ciphers. The type is bool. The default value is false.
    AesCbc interface{}

    // Enable 3des-cbc cipher. The type is bool. The default value is false.
    TripledesCbc interface{}
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

    cipher.EntityData.Children = types.NewOrderedMap()
    cipher.EntityData.Leafs = types.NewOrderedMap()
    cipher.EntityData.Leafs.Append("aes-cbc", types.YLeaf{"AesCbc", cipher.AesCbc})
    cipher.EntityData.Leafs.Append("tripledes-cbc", types.YLeaf{"TripledesCbc", cipher.TripledesCbc})

    cipher.EntityData.YListKeys = []string {}

    return &(cipher.EntityData)
}

// Ssh_Server_VrfTable
// Cisco sshd VRF name
type Ssh_Server_VrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Ssh_Server_VrfTable_Vrf.
    Vrf []*Ssh_Server_VrfTable_Vrf
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

    vrfTable.EntityData.Children = types.NewOrderedMap()
    vrfTable.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfTable.Vrf {
        vrfTable.EntityData.Children.Append(types.GetSegmentPath(vrfTable.Vrf[i]), types.YChild{"Vrf", vrfTable.Vrf[i]})
    }
    vrfTable.EntityData.Leafs = types.NewOrderedMap()

    vrfTable.EntityData.YListKeys = []string {}

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
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrf.Enable})
    vrf.EntityData.Leafs.Append("ipv4-access-list", types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList})
    vrf.EntityData.Leafs.Append("ipv6-access-list", types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ssh_Server_ServerAlgo
// Cisco ssh algorithms
type Ssh_Server_ServerAlgo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key exchange algorithm.
    KeyExchanges Ssh_Server_ServerAlgo_KeyExchanges
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

    serverAlgo.EntityData.Children = types.NewOrderedMap()
    serverAlgo.EntityData.Children.Append("key-exchanges", types.YChild{"KeyExchanges", &serverAlgo.KeyExchanges})
    serverAlgo.EntityData.Leafs = types.NewOrderedMap()

    serverAlgo.EntityData.YListKeys = []string {}

    return &(serverAlgo.EntityData)
}

// Ssh_Server_ServerAlgo_KeyExchanges
// Key exchange algorithm
type Ssh_Server_ServerAlgo_KeyExchanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key exchange algorithm. The type is slice of string with length: 1..32.
    KeyExchange []interface{}
}

func (keyExchanges *Ssh_Server_ServerAlgo_KeyExchanges) GetEntityData() *types.CommonEntityData {
    keyExchanges.EntityData.YFilter = keyExchanges.YFilter
    keyExchanges.EntityData.YangName = "key-exchanges"
    keyExchanges.EntityData.BundleName = "cisco_ios_xr"
    keyExchanges.EntityData.ParentYangName = "server-algo"
    keyExchanges.EntityData.SegmentPath = "key-exchanges"
    keyExchanges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyExchanges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyExchanges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyExchanges.EntityData.Children = types.NewOrderedMap()
    keyExchanges.EntityData.Leafs = types.NewOrderedMap()
    keyExchanges.EntityData.Leafs.Append("key-exchange", types.YLeaf{"KeyExchange", keyExchanges.KeyExchange})

    keyExchanges.EntityData.YListKeys = []string {}

    return &(keyExchanges.EntityData)
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

    capability.EntityData.Children = types.NewOrderedMap()
    capability.EntityData.Leafs = types.NewOrderedMap()
    capability.EntityData.Leafs.Append("netconf-xml", types.YLeaf{"NetconfXml", capability.NetconfXml})

    capability.EntityData.YListKeys = []string {}

    return &(capability.EntityData)
}

// Ssh_Server_NetconfVrfTable
// Cisco sshd Netconf VRF name
type Ssh_Server_NetconfVrfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter VRF name. The type is slice of Ssh_Server_NetconfVrfTable_Vrf.
    Vrf []*Ssh_Server_NetconfVrfTable_Vrf
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

    netconfVrfTable.EntityData.Children = types.NewOrderedMap()
    netconfVrfTable.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range netconfVrfTable.Vrf {
        netconfVrfTable.EntityData.Children.Append(types.GetSegmentPath(netconfVrfTable.Vrf[i]), types.YChild{"Vrf", netconfVrfTable.Vrf[i]})
    }
    netconfVrfTable.EntityData.Leafs = types.NewOrderedMap()

    netconfVrfTable.EntityData.YListKeys = []string {}

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
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrf.Enable})
    vrf.EntityData.Leafs.Append("ipv4-access-list", types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList})
    vrf.EntityData.Leafs.Append("ipv6-access-list", types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

