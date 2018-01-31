// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-ssh package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ssh1: Crypto Secure Shell(SSH) data
//   ssh: ssh
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package crypto_ssh_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_ssh_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-ssh-oper ssh1}", reflect.TypeOf(Ssh1{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-ssh-oper:ssh1", reflect.TypeOf(Ssh1{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-ssh-oper ssh}", reflect.TypeOf(Ssh{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-ssh-oper:ssh", reflect.TypeOf(Ssh{}))
}

// KexName represents Different key-exchange(kex) algorithms
type KexName string

const (
    // Diffie-Hellman group 1 key exchange algorithm
    KexName_diffie_hellman_group1 KexName = "diffie-hellman-group1"

    // Diffie-Hellman group 14 key exchange algorithm
    KexName_diffie_hellman_group14 KexName = "diffie-hellman-group14"

    // Diffie-Hellman group 14 key exchange algorithm
    KexName_diffie_hellman_group15 KexName = "diffie-hellman-group15"

    // Diffie-Hellman group 16 key exchange algorithm
    KexName_diffie_hellman_group16 KexName = "diffie-hellman-group16"

    // Diffie-Hellman group 17 key exchange algorithm
    KexName_diffie_hellman_group17 KexName = "diffie-hellman-group17"

    // Diffie-Hellman key group 18 exchange algorithm
    KexName_diffie_hellman_group18 KexName = "diffie-hellman-group18"

    // Elliptical curve Diffie-Hellman prime 256 key
    // exchange algorithm
    KexName_ecdh_nistp256 KexName = "ecdh-nistp256"

    // Elliptical curve Diffie-Hellman prime 384 key
    // exchange algorithm
    KexName_ecdh_nistp384 KexName = "ecdh-nistp384"

    // Elliptical curve Diffie-Hellman prime 521
    // exchange algorithm
    KexName_ecdh_nistp521 KexName = "ecdh-nistp521"

    // Password authenticated key agreement algorithm
    KexName_password_authenticated KexName = "password-authenticated"
)

// Hostkey represents SSH session authentication types
type Hostkey string

const (
    // Algorithm type DSS
    Hostkey_ssh_dss Hostkey = "ssh-dss"

    // Algorithm type RSA
    Hostkey_ssh_rsa Hostkey = "ssh-rsa"
)

// Version represents SSH state versions
type Version string

const (
    // Version V2
    Version_v2 Version = "v2"

    // Version V1
    Version_v1 Version = "v1"
)

// Connection represents SSH channel connection types
type Connection string

const (
    // connection type not yet known
    Connection_undefined Connection = "undefined"

    // Interactive Shell
    Connection_shell Connection = "shell"

    // Remote Command Execution
    Connection_exec Connection = "exec"

    // Secure Copy
    Connection_scp Connection = "scp"

    // Secure File Transfer
    Connection_sftp_subsystem Connection = "sftp-subsystem"

    // Netconf Subsystem
    Connection_netconf_subsystem Connection = "netconf-subsystem"

    // TL1 Subsystem
    Connection_tl1_subsystem Connection = "tl1-subsystem"

    // Netconf XML Subsystem
    Connection_netconf_xml_subsystem Connection = "netconf-xml-subsystem"
)

// States represents SSH session states
type States string

const (
    // SSH Open
    States_open States = "open"

    // SSH version OK
    States_version_ok States = "version-ok"

    // Key exchange(KEX) init message exchanged
    States_key_exchange_initialize States = "key-exchange-initialize"

    // Diffie-Hellman(DH) secret is generated
    States_key_exchange_dh States = "key-exchange-dh"

    // New keys are received
    States_new_keys States = "new-keys"

    // Need more information to authenticate
    States_authenticate_information States = "authenticate-information"

    // The client successfully authenticated
    States_authenticated States = "authenticated"

    // Channel has been successfully opened
    States_channel_open States = "channel-open"

    // Allocated PTY
    States_pty_open States = "pty-open"

    // Opened an exec shell
    States_session_open States = "session-open"

    // Received rekey request
    States_rekey States = "rekey"

    // Session is suspended
    States_suspended States = "suspended"

    // Session has been closed
    States_session_closed States = "session-closed"
)

// Mac represents functions
type Mac string

const (
    // Hash-based Message Authentication Code(HMAC)
    // MD5 algorithm
    Mac_hmac_md5 Mac = "hmac-md5"

    // Hash-based Message Authentication Code(HMAC)
    // SHA1 algorithm
    Mac_hmac_sha1 Mac = "hmac-sha1"

    // Hash-based Message Authentication Code(HMAC)
    // SHA2-256 algorithm
    Mac_hmac_sha2_256 Mac = "hmac-sha2-256"

    // Hash-based Message Authentication Code(HMAC)
    // SHA2-512 algorithm
    Mac_hmac_sha2_512 Mac = "hmac-sha2-512"

    // AES GCM based Authentication Tag as MAC
    // algorithm
    Mac_aes_gcm Mac = "aes-gcm"
)

// Cipher represents SSH session in and out cipher standards
type Cipher string

const (
    // Advanced Encryption Standard(AES) 128 bits
    // cipher block chaining(CBC)
    Cipher_aes128_cbc Cipher = "aes128-cbc"

    // Advanced Encryption Standard(AES) 192 bits
    // cipher block chaining(CBC)
    Cipher_aes192_cbc Cipher = "aes192-cbc"

    // Advanced Encryption Standard(AES) 256 bits
    // cipher block chaining(CBC)
    Cipher_aes256_cbc Cipher = "aes256-cbc"

    // Triple Data Encryption Standard(DES) cipher
    // block chaining(CBC)
    Cipher_triple_des_cbc Cipher = "triple-des-cbc"

    // Advanced Encryption Standard(AES) 128 bits
    // counter mode (CTR)
    Cipher_aes128_ctr Cipher = "aes128-ctr"

    // Advanced Encryption Standard(AES) 192 bits
    // counter mode (CTR)
    Cipher_aes192_ctr Cipher = "aes192-ctr"

    // Advanced Encryption Standard(AES) 256 bits
    // counter mode (CTR)
    Cipher_aes256_ctr Cipher = "aes256-ctr"

    // Advanced Encryption Standard(AES) 128 bits GCM
    // mode (GCM)
    Cipher_aes128_gcm Cipher = "aes128-gcm"

    // Advanced Encryption Standard(AES) 256 bits GCM
    // mode (GCM)
    Cipher_aes256_gcm Cipher = "aes256-gcm"
)

// Authen represents SSH session authentication types
type Authen string

const (
    // Password
    Authen_password Authen = "password"

    // RSA public key encryption type
    Authen_rsa_public_key Authen = "rsa-public-key"

    // Keyboard interactive
    Authen_keyboard_interactive Authen = "keyboard-interactive"
)

// Ssh1
// Crypto Secure Shell(SSH) data
type Ssh1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // key exchange method data.
    Kex Ssh1_Kex
}

func (ssh1 *Ssh1) GetFilter() yfilter.YFilter { return ssh1.YFilter }

func (ssh1 *Ssh1) SetFilter(yf yfilter.YFilter) { ssh1.YFilter = yf }

func (ssh1 *Ssh1) GetGoName(yname string) string {
    if yname == "kex" { return "Kex" }
    return ""
}

func (ssh1 *Ssh1) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-ssh-oper:ssh1"
}

func (ssh1 *Ssh1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "kex" {
        return &ssh1.Kex
    }
    return nil
}

func (ssh1 *Ssh1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["kex"] = &ssh1.Kex
    return children
}

func (ssh1 *Ssh1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssh1 *Ssh1) GetBundleName() string { return "cisco_ios_xr" }

func (ssh1 *Ssh1) GetYangName() string { return "ssh1" }

func (ssh1 *Ssh1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssh1 *Ssh1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssh1 *Ssh1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssh1 *Ssh1) SetParent(parent types.Entity) { ssh1.parent = parent }

func (ssh1 *Ssh1) GetParent() types.Entity { return ssh1.parent }

func (ssh1 *Ssh1) GetParentYangName() string { return "Cisco-IOS-XR-crypto-ssh-oper" }

// Ssh1_Kex
// key exchange method data
type Ssh1_Kex struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific ssh session details.
    Nodes Ssh1_Kex_Nodes
}

func (kex *Ssh1_Kex) GetFilter() yfilter.YFilter { return kex.YFilter }

func (kex *Ssh1_Kex) SetFilter(yf yfilter.YFilter) { kex.YFilter = yf }

func (kex *Ssh1_Kex) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (kex *Ssh1_Kex) GetSegmentPath() string {
    return "kex"
}

func (kex *Ssh1_Kex) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &kex.Nodes
    }
    return nil
}

func (kex *Ssh1_Kex) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &kex.Nodes
    return children
}

func (kex *Ssh1_Kex) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (kex *Ssh1_Kex) GetBundleName() string { return "cisco_ios_xr" }

func (kex *Ssh1_Kex) GetYangName() string { return "kex" }

func (kex *Ssh1_Kex) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (kex *Ssh1_Kex) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (kex *Ssh1_Kex) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (kex *Ssh1_Kex) SetParent(parent types.Entity) { kex.parent = parent }

func (kex *Ssh1_Kex) GetParent() types.Entity { return kex.parent }

func (kex *Ssh1_Kex) GetParentYangName() string { return "ssh1" }

// Ssh1_Kex_Nodes
// Node-specific ssh session details
type Ssh1_Kex_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSH session details for a particular node. The type is slice of
    // Ssh1_Kex_Nodes_Node.
    Node []Ssh1_Kex_Nodes_Node
}

func (nodes *Ssh1_Kex_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ssh1_Kex_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ssh1_Kex_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ssh1_Kex_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ssh1_Kex_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh1_Kex_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ssh1_Kex_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ssh1_Kex_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ssh1_Kex_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ssh1_Kex_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ssh1_Kex_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ssh1_Kex_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ssh1_Kex_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ssh1_Kex_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ssh1_Kex_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ssh1_Kex_Nodes) GetParentYangName() string { return "kex" }

// Ssh1_Kex_Nodes_Node
// SSH session details for a particular node
type Ssh1_Kex_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // List of incoming sessions.
    IncomingSessions Ssh1_Kex_Nodes_Node_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh1_Kex_Nodes_Node_OutgoingConnections
}

func (node *Ssh1_Kex_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ssh1_Kex_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ssh1_Kex_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "incoming-sessions" { return "IncomingSessions" }
    if yname == "outgoing-connections" { return "OutgoingConnections" }
    return ""
}

func (node *Ssh1_Kex_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ssh1_Kex_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incoming-sessions" {
        return &node.IncomingSessions
    }
    if childYangName == "outgoing-connections" {
        return &node.OutgoingConnections
    }
    return nil
}

func (node *Ssh1_Kex_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["incoming-sessions"] = &node.IncomingSessions
    children["outgoing-connections"] = &node.OutgoingConnections
    return children
}

func (node *Ssh1_Kex_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ssh1_Kex_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ssh1_Kex_Nodes_Node) GetYangName() string { return "node" }

func (node *Ssh1_Kex_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ssh1_Kex_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ssh1_Kex_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ssh1_Kex_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ssh1_Kex_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ssh1_Kex_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ssh1_Kex_Nodes_Node_IncomingSessions
// List of incoming sessions
type Ssh1_Kex_Nodes_Node_IncomingSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo.
    SessionDetailInfo []Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetFilter() yfilter.YFilter { return incomingSessions.YFilter }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) SetFilter(yf yfilter.YFilter) { incomingSessions.YFilter = yf }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetGoName(yname string) string {
    if yname == "session-detail-info" { return "SessionDetailInfo" }
    return ""
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetSegmentPath() string {
    return "incoming-sessions"
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detail-info" {
        for _, c := range incomingSessions.SessionDetailInfo {
            if incomingSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo{}
        incomingSessions.SessionDetailInfo = append(incomingSessions.SessionDetailInfo, child)
        return &incomingSessions.SessionDetailInfo[len(incomingSessions.SessionDetailInfo)-1]
    }
    return nil
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incomingSessions.SessionDetailInfo {
        children[incomingSessions.SessionDetailInfo[i].GetSegmentPath()] = &incomingSessions.SessionDetailInfo[i]
    }
    return children
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetBundleName() string { return "cisco_ios_xr" }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetYangName() string { return "incoming-sessions" }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) SetParent(parent types.Entity) { incomingSessions.parent = parent }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetParent() types.Entity { return incomingSessions.parent }

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetParentYangName() string { return "node" }

// Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo
// session detail info
type Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Key exchange name. The type is KexName.
    KeyExchange interface{}

    // Host key algorithm. The type is Hostkey.
    PublicKey interface{}

    // In cipher algorithm. The type is Cipher.
    InCipher interface{}

    // Out cipher algorithm. The type is Cipher.
    OutCipher interface{}

    // In MAC. The type is Mac.
    InMac interface{}

    // Out MAC. The type is Mac.
    OutMac interface{}
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetFilter() yfilter.YFilter { return sessionDetailInfo.YFilter }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) SetFilter(yf yfilter.YFilter) { sessionDetailInfo.YFilter = yf }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "key-exchange" { return "KeyExchange" }
    if yname == "public-key" { return "PublicKey" }
    if yname == "in-cipher" { return "InCipher" }
    if yname == "out-cipher" { return "OutCipher" }
    if yname == "in-mac" { return "InMac" }
    if yname == "out-mac" { return "OutMac" }
    return ""
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetSegmentPath() string {
    return "session-detail-info"
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionDetailInfo.SessionId
    leafs["key-exchange"] = sessionDetailInfo.KeyExchange
    leafs["public-key"] = sessionDetailInfo.PublicKey
    leafs["in-cipher"] = sessionDetailInfo.InCipher
    leafs["out-cipher"] = sessionDetailInfo.OutCipher
    leafs["in-mac"] = sessionDetailInfo.InMac
    leafs["out-mac"] = sessionDetailInfo.OutMac
    return leafs
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetYangName() string { return "session-detail-info" }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) SetParent(parent types.Entity) { sessionDetailInfo.parent = parent }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetParent() types.Entity { return sessionDetailInfo.parent }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetParentYangName() string { return "incoming-sessions" }

// Ssh1_Kex_Nodes_Node_OutgoingConnections
// List of outgoing connections
type Ssh1_Kex_Nodes_Node_OutgoingConnections struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo.
    SessionDetailInfo []Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetFilter() yfilter.YFilter { return outgoingConnections.YFilter }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) SetFilter(yf yfilter.YFilter) { outgoingConnections.YFilter = yf }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetGoName(yname string) string {
    if yname == "session-detail-info" { return "SessionDetailInfo" }
    return ""
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetSegmentPath() string {
    return "outgoing-connections"
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detail-info" {
        for _, c := range outgoingConnections.SessionDetailInfo {
            if outgoingConnections.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo{}
        outgoingConnections.SessionDetailInfo = append(outgoingConnections.SessionDetailInfo, child)
        return &outgoingConnections.SessionDetailInfo[len(outgoingConnections.SessionDetailInfo)-1]
    }
    return nil
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outgoingConnections.SessionDetailInfo {
        children[outgoingConnections.SessionDetailInfo[i].GetSegmentPath()] = &outgoingConnections.SessionDetailInfo[i]
    }
    return children
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetBundleName() string { return "cisco_ios_xr" }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetYangName() string { return "outgoing-connections" }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) SetParent(parent types.Entity) { outgoingConnections.parent = parent }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetParent() types.Entity { return outgoingConnections.parent }

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetParentYangName() string { return "node" }

// Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo
// session detail info
type Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Key exchange name. The type is KexName.
    KeyExchange interface{}

    // Host key algorithm. The type is Hostkey.
    PublicKey interface{}

    // In cipher algorithm. The type is Cipher.
    InCipher interface{}

    // Out cipher algorithm. The type is Cipher.
    OutCipher interface{}

    // In MAC. The type is Mac.
    InMac interface{}

    // Out MAC. The type is Mac.
    OutMac interface{}
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetFilter() yfilter.YFilter { return sessionDetailInfo.YFilter }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) SetFilter(yf yfilter.YFilter) { sessionDetailInfo.YFilter = yf }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "key-exchange" { return "KeyExchange" }
    if yname == "public-key" { return "PublicKey" }
    if yname == "in-cipher" { return "InCipher" }
    if yname == "out-cipher" { return "OutCipher" }
    if yname == "in-mac" { return "InMac" }
    if yname == "out-mac" { return "OutMac" }
    return ""
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetSegmentPath() string {
    return "session-detail-info"
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionDetailInfo.SessionId
    leafs["key-exchange"] = sessionDetailInfo.KeyExchange
    leafs["public-key"] = sessionDetailInfo.PublicKey
    leafs["in-cipher"] = sessionDetailInfo.InCipher
    leafs["out-cipher"] = sessionDetailInfo.OutCipher
    leafs["in-mac"] = sessionDetailInfo.InMac
    leafs["out-mac"] = sessionDetailInfo.OutMac
    return leafs
}

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetYangName() string { return "session-detail-info" }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) SetParent(parent types.Entity) { sessionDetailInfo.parent = parent }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetParent() types.Entity { return sessionDetailInfo.parent }

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetParentYangName() string { return "outgoing-connections" }

// Ssh
// ssh
type Ssh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Crypto SSH session.
    Session Ssh_Session
}

func (ssh *Ssh) GetFilter() yfilter.YFilter { return ssh.YFilter }

func (ssh *Ssh) SetFilter(yf yfilter.YFilter) { ssh.YFilter = yf }

func (ssh *Ssh) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (ssh *Ssh) GetSegmentPath() string {
    return "Cisco-IOS-XR-crypto-ssh-oper:ssh"
}

func (ssh *Ssh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &ssh.Session
    }
    return nil
}

func (ssh *Ssh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &ssh.Session
    return children
}

func (ssh *Ssh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssh *Ssh) GetBundleName() string { return "cisco_ios_xr" }

func (ssh *Ssh) GetYangName() string { return "ssh" }

func (ssh *Ssh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssh *Ssh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssh *Ssh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssh *Ssh) SetParent(parent types.Entity) { ssh.parent = parent }

func (ssh *Ssh) GetParent() types.Entity { return ssh.parent }

func (ssh *Ssh) GetParentYangName() string { return "Cisco-IOS-XR-crypto-ssh-oper" }

// Ssh_Session
// Crypto SSH session
type Ssh_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSH session rekey information.
    Rekey Ssh_Session_Rekey

    // SSH session brief information.
    Brief Ssh_Session_Brief

    // SSH session detail information.
    Detail Ssh_Session_Detail
}

func (session *Ssh_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Ssh_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Ssh_Session) GetGoName(yname string) string {
    if yname == "rekey" { return "Rekey" }
    if yname == "brief" { return "Brief" }
    if yname == "detail" { return "Detail" }
    return ""
}

func (session *Ssh_Session) GetSegmentPath() string {
    return "session"
}

func (session *Ssh_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rekey" {
        return &session.Rekey
    }
    if childYangName == "brief" {
        return &session.Brief
    }
    if childYangName == "detail" {
        return &session.Detail
    }
    return nil
}

func (session *Ssh_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rekey"] = &session.Rekey
    children["brief"] = &session.Brief
    children["detail"] = &session.Detail
    return children
}

func (session *Ssh_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *Ssh_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Ssh_Session) GetYangName() string { return "session" }

func (session *Ssh_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Ssh_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Ssh_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Ssh_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Ssh_Session) GetParent() types.Entity { return session.parent }

func (session *Ssh_Session) GetParentYangName() string { return "ssh" }

// Ssh_Session_Rekey
// SSH session rekey information
type Ssh_Session_Rekey struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Rekey_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh_Session_Rekey_OutgoingConnections
}

func (rekey *Ssh_Session_Rekey) GetFilter() yfilter.YFilter { return rekey.YFilter }

func (rekey *Ssh_Session_Rekey) SetFilter(yf yfilter.YFilter) { rekey.YFilter = yf }

func (rekey *Ssh_Session_Rekey) GetGoName(yname string) string {
    if yname == "incoming-sessions" { return "IncomingSessions" }
    if yname == "outgoing-connections" { return "OutgoingConnections" }
    return ""
}

func (rekey *Ssh_Session_Rekey) GetSegmentPath() string {
    return "rekey"
}

func (rekey *Ssh_Session_Rekey) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incoming-sessions" {
        return &rekey.IncomingSessions
    }
    if childYangName == "outgoing-connections" {
        return &rekey.OutgoingConnections
    }
    return nil
}

func (rekey *Ssh_Session_Rekey) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["incoming-sessions"] = &rekey.IncomingSessions
    children["outgoing-connections"] = &rekey.OutgoingConnections
    return children
}

func (rekey *Ssh_Session_Rekey) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rekey *Ssh_Session_Rekey) GetBundleName() string { return "cisco_ios_xr" }

func (rekey *Ssh_Session_Rekey) GetYangName() string { return "rekey" }

func (rekey *Ssh_Session_Rekey) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rekey *Ssh_Session_Rekey) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rekey *Ssh_Session_Rekey) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rekey *Ssh_Session_Rekey) SetParent(parent types.Entity) { rekey.parent = parent }

func (rekey *Ssh_Session_Rekey) GetParent() types.Entity { return rekey.parent }

func (rekey *Ssh_Session_Rekey) GetParentYangName() string { return "session" }

// Ssh_Session_Rekey_IncomingSessions
// List of incoming sessions
type Ssh_Session_Rekey_IncomingSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session rekey info. The type is slice of
    // Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo.
    SessionRekeyInfo []Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetFilter() yfilter.YFilter { return incomingSessions.YFilter }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) SetFilter(yf yfilter.YFilter) { incomingSessions.YFilter = yf }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetGoName(yname string) string {
    if yname == "session-rekey-info" { return "SessionRekeyInfo" }
    return ""
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetSegmentPath() string {
    return "incoming-sessions"
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-rekey-info" {
        for _, c := range incomingSessions.SessionRekeyInfo {
            if incomingSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo{}
        incomingSessions.SessionRekeyInfo = append(incomingSessions.SessionRekeyInfo, child)
        return &incomingSessions.SessionRekeyInfo[len(incomingSessions.SessionRekeyInfo)-1]
    }
    return nil
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incomingSessions.SessionRekeyInfo {
        children[incomingSessions.SessionRekeyInfo[i].GetSegmentPath()] = &incomingSessions.SessionRekeyInfo[i]
    }
    return children
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetBundleName() string { return "cisco_ios_xr" }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetYangName() string { return "incoming-sessions" }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) SetParent(parent types.Entity) { incomingSessions.parent = parent }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetParent() types.Entity { return incomingSessions.parent }

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetParentYangName() string { return "rekey" }

// Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo
// session rekey info
type Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Session Rekey Count. The type is interface{} with range: 0..4294967295.
    SessionRekeyCount interface{}

    // Time To Rekey. The type is string.
    TimeToRekey interface{}

    // Volume To Rekey. The type is string.
    VolumeToRekey interface{}
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetFilter() yfilter.YFilter { return sessionRekeyInfo.YFilter }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) SetFilter(yf yfilter.YFilter) { sessionRekeyInfo.YFilter = yf }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "session-rekey-count" { return "SessionRekeyCount" }
    if yname == "time-to-rekey" { return "TimeToRekey" }
    if yname == "volume-to-rekey" { return "VolumeToRekey" }
    return ""
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetSegmentPath() string {
    return "session-rekey-info"
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionRekeyInfo.SessionId
    leafs["session-rekey-count"] = sessionRekeyInfo.SessionRekeyCount
    leafs["time-to-rekey"] = sessionRekeyInfo.TimeToRekey
    leafs["volume-to-rekey"] = sessionRekeyInfo.VolumeToRekey
    return leafs
}

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetYangName() string { return "session-rekey-info" }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) SetParent(parent types.Entity) { sessionRekeyInfo.parent = parent }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetParent() types.Entity { return sessionRekeyInfo.parent }

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetParentYangName() string { return "incoming-sessions" }

// Ssh_Session_Rekey_OutgoingConnections
// List of outgoing connections
type Ssh_Session_Rekey_OutgoingConnections struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session rekey info. The type is slice of
    // Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo.
    SessionRekeyInfo []Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetFilter() yfilter.YFilter { return outgoingConnections.YFilter }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) SetFilter(yf yfilter.YFilter) { outgoingConnections.YFilter = yf }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetGoName(yname string) string {
    if yname == "session-rekey-info" { return "SessionRekeyInfo" }
    return ""
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetSegmentPath() string {
    return "outgoing-connections"
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-rekey-info" {
        for _, c := range outgoingConnections.SessionRekeyInfo {
            if outgoingConnections.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo{}
        outgoingConnections.SessionRekeyInfo = append(outgoingConnections.SessionRekeyInfo, child)
        return &outgoingConnections.SessionRekeyInfo[len(outgoingConnections.SessionRekeyInfo)-1]
    }
    return nil
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outgoingConnections.SessionRekeyInfo {
        children[outgoingConnections.SessionRekeyInfo[i].GetSegmentPath()] = &outgoingConnections.SessionRekeyInfo[i]
    }
    return children
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetBundleName() string { return "cisco_ios_xr" }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetYangName() string { return "outgoing-connections" }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) SetParent(parent types.Entity) { outgoingConnections.parent = parent }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetParent() types.Entity { return outgoingConnections.parent }

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetParentYangName() string { return "rekey" }

// Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo
// session rekey info
type Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Session Rekey Count. The type is interface{} with range: 0..4294967295.
    SessionRekeyCount interface{}

    // Time To Rekey. The type is string.
    TimeToRekey interface{}

    // Volume To Rekey. The type is string.
    VolumeToRekey interface{}
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetFilter() yfilter.YFilter { return sessionRekeyInfo.YFilter }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) SetFilter(yf yfilter.YFilter) { sessionRekeyInfo.YFilter = yf }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "session-rekey-count" { return "SessionRekeyCount" }
    if yname == "time-to-rekey" { return "TimeToRekey" }
    if yname == "volume-to-rekey" { return "VolumeToRekey" }
    return ""
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetSegmentPath() string {
    return "session-rekey-info"
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionRekeyInfo.SessionId
    leafs["session-rekey-count"] = sessionRekeyInfo.SessionRekeyCount
    leafs["time-to-rekey"] = sessionRekeyInfo.TimeToRekey
    leafs["volume-to-rekey"] = sessionRekeyInfo.VolumeToRekey
    return leafs
}

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetYangName() string { return "session-rekey-info" }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) SetParent(parent types.Entity) { sessionRekeyInfo.parent = parent }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetParent() types.Entity { return sessionRekeyInfo.parent }

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetParentYangName() string { return "outgoing-connections" }

// Ssh_Session_Brief
// SSH session brief information
type Ssh_Session_Brief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Brief_IncomingSessions

    // List of outgoing sessions.
    OutgoingSessions Ssh_Session_Brief_OutgoingSessions
}

func (brief *Ssh_Session_Brief) GetFilter() yfilter.YFilter { return brief.YFilter }

func (brief *Ssh_Session_Brief) SetFilter(yf yfilter.YFilter) { brief.YFilter = yf }

func (brief *Ssh_Session_Brief) GetGoName(yname string) string {
    if yname == "incoming-sessions" { return "IncomingSessions" }
    if yname == "outgoing-sessions" { return "OutgoingSessions" }
    return ""
}

func (brief *Ssh_Session_Brief) GetSegmentPath() string {
    return "brief"
}

func (brief *Ssh_Session_Brief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incoming-sessions" {
        return &brief.IncomingSessions
    }
    if childYangName == "outgoing-sessions" {
        return &brief.OutgoingSessions
    }
    return nil
}

func (brief *Ssh_Session_Brief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["incoming-sessions"] = &brief.IncomingSessions
    children["outgoing-sessions"] = &brief.OutgoingSessions
    return children
}

func (brief *Ssh_Session_Brief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (brief *Ssh_Session_Brief) GetBundleName() string { return "cisco_ios_xr" }

func (brief *Ssh_Session_Brief) GetYangName() string { return "brief" }

func (brief *Ssh_Session_Brief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (brief *Ssh_Session_Brief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (brief *Ssh_Session_Brief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (brief *Ssh_Session_Brief) SetParent(parent types.Entity) { brief.parent = parent }

func (brief *Ssh_Session_Brief) GetParent() types.Entity { return brief.parent }

func (brief *Ssh_Session_Brief) GetParentYangName() string { return "session" }

// Ssh_Session_Brief_IncomingSessions
// List of incoming sessions
type Ssh_Session_Brief_IncomingSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session brief info. The type is slice of
    // Ssh_Session_Brief_IncomingSessions_SessionBriefInfo.
    SessionBriefInfo []Ssh_Session_Brief_IncomingSessions_SessionBriefInfo
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetFilter() yfilter.YFilter { return incomingSessions.YFilter }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) SetFilter(yf yfilter.YFilter) { incomingSessions.YFilter = yf }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetGoName(yname string) string {
    if yname == "session-brief-info" { return "SessionBriefInfo" }
    return ""
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetSegmentPath() string {
    return "incoming-sessions"
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-brief-info" {
        for _, c := range incomingSessions.SessionBriefInfo {
            if incomingSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Brief_IncomingSessions_SessionBriefInfo{}
        incomingSessions.SessionBriefInfo = append(incomingSessions.SessionBriefInfo, child)
        return &incomingSessions.SessionBriefInfo[len(incomingSessions.SessionBriefInfo)-1]
    }
    return nil
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incomingSessions.SessionBriefInfo {
        children[incomingSessions.SessionBriefInfo[i].GetSegmentPath()] = &incomingSessions.SessionBriefInfo[i]
    }
    return children
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetBundleName() string { return "cisco_ios_xr" }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetYangName() string { return "incoming-sessions" }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) SetParent(parent types.Entity) { incomingSessions.parent = parent }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetParent() types.Entity { return incomingSessions.parent }

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetParentYangName() string { return "brief" }

// Ssh_Session_Brief_IncomingSessions_SessionBriefInfo
// session brief info
type Ssh_Session_Brief_IncomingSessions_SessionBriefInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Channel ID. The type is interface{} with range: 0..4294967295.
    ChannelId interface{}

    // Boolean indicating whether line VTY line number is valid. The type is bool.
    VtyAssigned interface{}

    // VTY line number. The type is interface{} with range: 0..4294967295.
    VtyLineNumber interface{}

    // Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // SSH session state. The type is States.
    SessionState interface{}

    // User ID. The type is string.
    UserId interface{}

    // Host address. The type is string.
    HostAddress interface{}

    // SSH state version. The type is Version.
    Version interface{}

    // Authentication method. The type is Authen.
    AuthenticationType interface{}

    // Channel Connection Type. The type is Connection.
    ConnectionType interface{}
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetFilter() yfilter.YFilter { return sessionBriefInfo.YFilter }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) SetFilter(yf yfilter.YFilter) { sessionBriefInfo.YFilter = yf }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "channel-id" { return "ChannelId" }
    if yname == "vty-assigned" { return "VtyAssigned" }
    if yname == "vty-line-number" { return "VtyLineNumber" }
    if yname == "node-name" { return "NodeName" }
    if yname == "session-state" { return "SessionState" }
    if yname == "user-id" { return "UserId" }
    if yname == "host-address" { return "HostAddress" }
    if yname == "version" { return "Version" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "connection-type" { return "ConnectionType" }
    return ""
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetSegmentPath() string {
    return "session-brief-info"
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionBriefInfo.SessionId
    leafs["channel-id"] = sessionBriefInfo.ChannelId
    leafs["vty-assigned"] = sessionBriefInfo.VtyAssigned
    leafs["vty-line-number"] = sessionBriefInfo.VtyLineNumber
    leafs["node-name"] = sessionBriefInfo.NodeName
    leafs["session-state"] = sessionBriefInfo.SessionState
    leafs["user-id"] = sessionBriefInfo.UserId
    leafs["host-address"] = sessionBriefInfo.HostAddress
    leafs["version"] = sessionBriefInfo.Version
    leafs["authentication-type"] = sessionBriefInfo.AuthenticationType
    leafs["connection-type"] = sessionBriefInfo.ConnectionType
    return leafs
}

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetYangName() string { return "session-brief-info" }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) SetParent(parent types.Entity) { sessionBriefInfo.parent = parent }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetParent() types.Entity { return sessionBriefInfo.parent }

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetParentYangName() string { return "incoming-sessions" }

// Ssh_Session_Brief_OutgoingSessions
// List of outgoing sessions
type Ssh_Session_Brief_OutgoingSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session brief info. The type is slice of
    // Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo.
    SessionBriefInfo []Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetFilter() yfilter.YFilter { return outgoingSessions.YFilter }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) SetFilter(yf yfilter.YFilter) { outgoingSessions.YFilter = yf }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetGoName(yname string) string {
    if yname == "session-brief-info" { return "SessionBriefInfo" }
    return ""
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetSegmentPath() string {
    return "outgoing-sessions"
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-brief-info" {
        for _, c := range outgoingSessions.SessionBriefInfo {
            if outgoingSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo{}
        outgoingSessions.SessionBriefInfo = append(outgoingSessions.SessionBriefInfo, child)
        return &outgoingSessions.SessionBriefInfo[len(outgoingSessions.SessionBriefInfo)-1]
    }
    return nil
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outgoingSessions.SessionBriefInfo {
        children[outgoingSessions.SessionBriefInfo[i].GetSegmentPath()] = &outgoingSessions.SessionBriefInfo[i]
    }
    return children
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetBundleName() string { return "cisco_ios_xr" }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetYangName() string { return "outgoing-sessions" }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) SetParent(parent types.Entity) { outgoingSessions.parent = parent }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetParent() types.Entity { return outgoingSessions.parent }

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetParentYangName() string { return "brief" }

// Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo
// session brief info
type Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Channel ID. The type is interface{} with range: 0..4294967295.
    ChannelId interface{}

    // Boolean indicating whether line VTY line number is valid. The type is bool.
    VtyAssigned interface{}

    // VTY line number. The type is interface{} with range: 0..4294967295.
    VtyLineNumber interface{}

    // Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // SSH session state. The type is States.
    SessionState interface{}

    // User ID. The type is string.
    UserId interface{}

    // Host address. The type is string.
    HostAddress interface{}

    // SSH state version. The type is Version.
    Version interface{}

    // Authentication method. The type is Authen.
    AuthenticationType interface{}

    // Channel Connection Type. The type is Connection.
    ConnectionType interface{}
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetFilter() yfilter.YFilter { return sessionBriefInfo.YFilter }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) SetFilter(yf yfilter.YFilter) { sessionBriefInfo.YFilter = yf }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "channel-id" { return "ChannelId" }
    if yname == "vty-assigned" { return "VtyAssigned" }
    if yname == "vty-line-number" { return "VtyLineNumber" }
    if yname == "node-name" { return "NodeName" }
    if yname == "session-state" { return "SessionState" }
    if yname == "user-id" { return "UserId" }
    if yname == "host-address" { return "HostAddress" }
    if yname == "version" { return "Version" }
    if yname == "authentication-type" { return "AuthenticationType" }
    if yname == "connection-type" { return "ConnectionType" }
    return ""
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetSegmentPath() string {
    return "session-brief-info"
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionBriefInfo.SessionId
    leafs["channel-id"] = sessionBriefInfo.ChannelId
    leafs["vty-assigned"] = sessionBriefInfo.VtyAssigned
    leafs["vty-line-number"] = sessionBriefInfo.VtyLineNumber
    leafs["node-name"] = sessionBriefInfo.NodeName
    leafs["session-state"] = sessionBriefInfo.SessionState
    leafs["user-id"] = sessionBriefInfo.UserId
    leafs["host-address"] = sessionBriefInfo.HostAddress
    leafs["version"] = sessionBriefInfo.Version
    leafs["authentication-type"] = sessionBriefInfo.AuthenticationType
    leafs["connection-type"] = sessionBriefInfo.ConnectionType
    return leafs
}

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetYangName() string { return "session-brief-info" }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) SetParent(parent types.Entity) { sessionBriefInfo.parent = parent }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetParent() types.Entity { return sessionBriefInfo.parent }

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetParentYangName() string { return "outgoing-sessions" }

// Ssh_Session_Detail
// SSH session detail information
type Ssh_Session_Detail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Detail_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh_Session_Detail_OutgoingConnections
}

func (detail *Ssh_Session_Detail) GetFilter() yfilter.YFilter { return detail.YFilter }

func (detail *Ssh_Session_Detail) SetFilter(yf yfilter.YFilter) { detail.YFilter = yf }

func (detail *Ssh_Session_Detail) GetGoName(yname string) string {
    if yname == "incoming-sessions" { return "IncomingSessions" }
    if yname == "outgoing-connections" { return "OutgoingConnections" }
    return ""
}

func (detail *Ssh_Session_Detail) GetSegmentPath() string {
    return "detail"
}

func (detail *Ssh_Session_Detail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "incoming-sessions" {
        return &detail.IncomingSessions
    }
    if childYangName == "outgoing-connections" {
        return &detail.OutgoingConnections
    }
    return nil
}

func (detail *Ssh_Session_Detail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["incoming-sessions"] = &detail.IncomingSessions
    children["outgoing-connections"] = &detail.OutgoingConnections
    return children
}

func (detail *Ssh_Session_Detail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detail *Ssh_Session_Detail) GetBundleName() string { return "cisco_ios_xr" }

func (detail *Ssh_Session_Detail) GetYangName() string { return "detail" }

func (detail *Ssh_Session_Detail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detail *Ssh_Session_Detail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detail *Ssh_Session_Detail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detail *Ssh_Session_Detail) SetParent(parent types.Entity) { detail.parent = parent }

func (detail *Ssh_Session_Detail) GetParent() types.Entity { return detail.parent }

func (detail *Ssh_Session_Detail) GetParentYangName() string { return "session" }

// Ssh_Session_Detail_IncomingSessions
// List of incoming sessions
type Ssh_Session_Detail_IncomingSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh_Session_Detail_IncomingSessions_SessionDetailInfo.
    SessionDetailInfo []Ssh_Session_Detail_IncomingSessions_SessionDetailInfo
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetFilter() yfilter.YFilter { return incomingSessions.YFilter }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) SetFilter(yf yfilter.YFilter) { incomingSessions.YFilter = yf }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetGoName(yname string) string {
    if yname == "session-detail-info" { return "SessionDetailInfo" }
    return ""
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetSegmentPath() string {
    return "incoming-sessions"
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detail-info" {
        for _, c := range incomingSessions.SessionDetailInfo {
            if incomingSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Detail_IncomingSessions_SessionDetailInfo{}
        incomingSessions.SessionDetailInfo = append(incomingSessions.SessionDetailInfo, child)
        return &incomingSessions.SessionDetailInfo[len(incomingSessions.SessionDetailInfo)-1]
    }
    return nil
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range incomingSessions.SessionDetailInfo {
        children[incomingSessions.SessionDetailInfo[i].GetSegmentPath()] = &incomingSessions.SessionDetailInfo[i]
    }
    return children
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetBundleName() string { return "cisco_ios_xr" }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetYangName() string { return "incoming-sessions" }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) SetParent(parent types.Entity) { incomingSessions.parent = parent }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetParent() types.Entity { return incomingSessions.parent }

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetParentYangName() string { return "detail" }

// Ssh_Session_Detail_IncomingSessions_SessionDetailInfo
// session detail info
type Ssh_Session_Detail_IncomingSessions_SessionDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Key exchange name. The type is KexName.
    KeyExchange interface{}

    // Host key algorithm. The type is Hostkey.
    PublicKey interface{}

    // In cipher algorithm. The type is Cipher.
    InCipher interface{}

    // Out cipher algorithm. The type is Cipher.
    OutCipher interface{}

    // In MAC. The type is Mac.
    InMac interface{}

    // Out MAC. The type is Mac.
    OutMac interface{}
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetFilter() yfilter.YFilter { return sessionDetailInfo.YFilter }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) SetFilter(yf yfilter.YFilter) { sessionDetailInfo.YFilter = yf }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "key-exchange" { return "KeyExchange" }
    if yname == "public-key" { return "PublicKey" }
    if yname == "in-cipher" { return "InCipher" }
    if yname == "out-cipher" { return "OutCipher" }
    if yname == "in-mac" { return "InMac" }
    if yname == "out-mac" { return "OutMac" }
    return ""
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetSegmentPath() string {
    return "session-detail-info"
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionDetailInfo.SessionId
    leafs["key-exchange"] = sessionDetailInfo.KeyExchange
    leafs["public-key"] = sessionDetailInfo.PublicKey
    leafs["in-cipher"] = sessionDetailInfo.InCipher
    leafs["out-cipher"] = sessionDetailInfo.OutCipher
    leafs["in-mac"] = sessionDetailInfo.InMac
    leafs["out-mac"] = sessionDetailInfo.OutMac
    return leafs
}

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetYangName() string { return "session-detail-info" }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) SetParent(parent types.Entity) { sessionDetailInfo.parent = parent }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetParent() types.Entity { return sessionDetailInfo.parent }

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetParentYangName() string { return "incoming-sessions" }

// Ssh_Session_Detail_OutgoingConnections
// List of outgoing connections
type Ssh_Session_Detail_OutgoingConnections struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo.
    SessionDetailInfo []Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetFilter() yfilter.YFilter { return outgoingConnections.YFilter }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) SetFilter(yf yfilter.YFilter) { outgoingConnections.YFilter = yf }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetGoName(yname string) string {
    if yname == "session-detail-info" { return "SessionDetailInfo" }
    return ""
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetSegmentPath() string {
    return "outgoing-connections"
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-detail-info" {
        for _, c := range outgoingConnections.SessionDetailInfo {
            if outgoingConnections.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo{}
        outgoingConnections.SessionDetailInfo = append(outgoingConnections.SessionDetailInfo, child)
        return &outgoingConnections.SessionDetailInfo[len(outgoingConnections.SessionDetailInfo)-1]
    }
    return nil
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range outgoingConnections.SessionDetailInfo {
        children[outgoingConnections.SessionDetailInfo[i].GetSegmentPath()] = &outgoingConnections.SessionDetailInfo[i]
    }
    return children
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetBundleName() string { return "cisco_ios_xr" }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetYangName() string { return "outgoing-connections" }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) SetParent(parent types.Entity) { outgoingConnections.parent = parent }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetParent() types.Entity { return outgoingConnections.parent }

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetParentYangName() string { return "detail" }

// Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo
// session detail info
type Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Key exchange name. The type is KexName.
    KeyExchange interface{}

    // Host key algorithm. The type is Hostkey.
    PublicKey interface{}

    // In cipher algorithm. The type is Cipher.
    InCipher interface{}

    // Out cipher algorithm. The type is Cipher.
    OutCipher interface{}

    // In MAC. The type is Mac.
    InMac interface{}

    // Out MAC. The type is Mac.
    OutMac interface{}
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetFilter() yfilter.YFilter { return sessionDetailInfo.YFilter }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) SetFilter(yf yfilter.YFilter) { sessionDetailInfo.YFilter = yf }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "key-exchange" { return "KeyExchange" }
    if yname == "public-key" { return "PublicKey" }
    if yname == "in-cipher" { return "InCipher" }
    if yname == "out-cipher" { return "OutCipher" }
    if yname == "in-mac" { return "InMac" }
    if yname == "out-mac" { return "OutMac" }
    return ""
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetSegmentPath() string {
    return "session-detail-info"
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionDetailInfo.SessionId
    leafs["key-exchange"] = sessionDetailInfo.KeyExchange
    leafs["public-key"] = sessionDetailInfo.PublicKey
    leafs["in-cipher"] = sessionDetailInfo.InCipher
    leafs["out-cipher"] = sessionDetailInfo.OutCipher
    leafs["in-mac"] = sessionDetailInfo.InMac
    leafs["out-mac"] = sessionDetailInfo.OutMac
    return leafs
}

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetYangName() string { return "session-detail-info" }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) SetParent(parent types.Entity) { sessionDetailInfo.parent = parent }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetParent() types.Entity { return sessionDetailInfo.parent }

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetParentYangName() string { return "outgoing-connections" }

