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

// Version represents SSH state versions
type Version string

const (
    // Version V2
    Version_v2 Version = "v2"

    // Version V1
    Version_v1 Version = "v1"
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

// Hostkey represents SSH session authentication types
type Hostkey string

const (
    // Algorithm type DSS
    Hostkey_ssh_dss Hostkey = "ssh-dss"

    // Algorithm type RSA
    Hostkey_ssh_rsa Hostkey = "ssh-rsa"
)

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

// Ssh1
// Crypto Secure Shell(SSH) data
type Ssh1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // key exchange method data.
    Kex Ssh1_Kex
}

func (ssh1 *Ssh1) GetEntityData() *types.CommonEntityData {
    ssh1.EntityData.YFilter = ssh1.YFilter
    ssh1.EntityData.YangName = "ssh1"
    ssh1.EntityData.BundleName = "cisco_ios_xr"
    ssh1.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-ssh-oper"
    ssh1.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-ssh-oper:ssh1"
    ssh1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssh1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssh1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssh1.EntityData.Children = make(map[string]types.YChild)
    ssh1.EntityData.Children["kex"] = types.YChild{"Kex", &ssh1.Kex}
    ssh1.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ssh1.EntityData)
}

// Ssh1_Kex
// key exchange method data
type Ssh1_Kex struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific ssh session details.
    Nodes Ssh1_Kex_Nodes
}

func (kex *Ssh1_Kex) GetEntityData() *types.CommonEntityData {
    kex.EntityData.YFilter = kex.YFilter
    kex.EntityData.YangName = "kex"
    kex.EntityData.BundleName = "cisco_ios_xr"
    kex.EntityData.ParentYangName = "ssh1"
    kex.EntityData.SegmentPath = "kex"
    kex.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    kex.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    kex.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    kex.EntityData.Children = make(map[string]types.YChild)
    kex.EntityData.Children["nodes"] = types.YChild{"Nodes", &kex.Nodes}
    kex.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(kex.EntityData)
}

// Ssh1_Kex_Nodes
// Node-specific ssh session details
type Ssh1_Kex_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSH session details for a particular node. The type is slice of
    // Ssh1_Kex_Nodes_Node.
    Node []Ssh1_Kex_Nodes_Node
}

func (nodes *Ssh1_Kex_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "kex"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Ssh1_Kex_Nodes_Node
// SSH session details for a particular node
type Ssh1_Kex_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // List of incoming sessions.
    IncomingSessions Ssh1_Kex_Nodes_Node_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh1_Kex_Nodes_Node_OutgoingConnections
}

func (node *Ssh1_Kex_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["incoming-sessions"] = types.YChild{"IncomingSessions", &node.IncomingSessions}
    node.EntityData.Children["outgoing-connections"] = types.YChild{"OutgoingConnections", &node.OutgoingConnections}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Ssh1_Kex_Nodes_Node_IncomingSessions
// List of incoming sessions
type Ssh1_Kex_Nodes_Node_IncomingSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo.
    SessionDetailInfo []Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo
}

func (incomingSessions *Ssh1_Kex_Nodes_Node_IncomingSessions) GetEntityData() *types.CommonEntityData {
    incomingSessions.EntityData.YFilter = incomingSessions.YFilter
    incomingSessions.EntityData.YangName = "incoming-sessions"
    incomingSessions.EntityData.BundleName = "cisco_ios_xr"
    incomingSessions.EntityData.ParentYangName = "node"
    incomingSessions.EntityData.SegmentPath = "incoming-sessions"
    incomingSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingSessions.EntityData.Children = make(map[string]types.YChild)
    incomingSessions.EntityData.Children["session-detail-info"] = types.YChild{"SessionDetailInfo", nil}
    for i := range incomingSessions.SessionDetailInfo {
        incomingSessions.EntityData.Children[types.GetSegmentPath(&incomingSessions.SessionDetailInfo[i])] = types.YChild{"SessionDetailInfo", &incomingSessions.SessionDetailInfo[i]}
    }
    incomingSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incomingSessions.EntityData)
}

// Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo
// session detail info
type Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo struct {
    EntityData types.CommonEntityData
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

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_IncomingSessions_SessionDetailInfo) GetEntityData() *types.CommonEntityData {
    sessionDetailInfo.EntityData.YFilter = sessionDetailInfo.YFilter
    sessionDetailInfo.EntityData.YangName = "session-detail-info"
    sessionDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailInfo.EntityData.ParentYangName = "incoming-sessions"
    sessionDetailInfo.EntityData.SegmentPath = "session-detail-info"
    sessionDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailInfo.EntityData.Children = make(map[string]types.YChild)
    sessionDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDetailInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionDetailInfo.SessionId}
    sessionDetailInfo.EntityData.Leafs["key-exchange"] = types.YLeaf{"KeyExchange", sessionDetailInfo.KeyExchange}
    sessionDetailInfo.EntityData.Leafs["public-key"] = types.YLeaf{"PublicKey", sessionDetailInfo.PublicKey}
    sessionDetailInfo.EntityData.Leafs["in-cipher"] = types.YLeaf{"InCipher", sessionDetailInfo.InCipher}
    sessionDetailInfo.EntityData.Leafs["out-cipher"] = types.YLeaf{"OutCipher", sessionDetailInfo.OutCipher}
    sessionDetailInfo.EntityData.Leafs["in-mac"] = types.YLeaf{"InMac", sessionDetailInfo.InMac}
    sessionDetailInfo.EntityData.Leafs["out-mac"] = types.YLeaf{"OutMac", sessionDetailInfo.OutMac}
    return &(sessionDetailInfo.EntityData)
}

// Ssh1_Kex_Nodes_Node_OutgoingConnections
// List of outgoing connections
type Ssh1_Kex_Nodes_Node_OutgoingConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo.
    SessionDetailInfo []Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo
}

func (outgoingConnections *Ssh1_Kex_Nodes_Node_OutgoingConnections) GetEntityData() *types.CommonEntityData {
    outgoingConnections.EntityData.YFilter = outgoingConnections.YFilter
    outgoingConnections.EntityData.YangName = "outgoing-connections"
    outgoingConnections.EntityData.BundleName = "cisco_ios_xr"
    outgoingConnections.EntityData.ParentYangName = "node"
    outgoingConnections.EntityData.SegmentPath = "outgoing-connections"
    outgoingConnections.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outgoingConnections.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outgoingConnections.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outgoingConnections.EntityData.Children = make(map[string]types.YChild)
    outgoingConnections.EntityData.Children["session-detail-info"] = types.YChild{"SessionDetailInfo", nil}
    for i := range outgoingConnections.SessionDetailInfo {
        outgoingConnections.EntityData.Children[types.GetSegmentPath(&outgoingConnections.SessionDetailInfo[i])] = types.YChild{"SessionDetailInfo", &outgoingConnections.SessionDetailInfo[i]}
    }
    outgoingConnections.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outgoingConnections.EntityData)
}

// Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo
// session detail info
type Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo struct {
    EntityData types.CommonEntityData
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

func (sessionDetailInfo *Ssh1_Kex_Nodes_Node_OutgoingConnections_SessionDetailInfo) GetEntityData() *types.CommonEntityData {
    sessionDetailInfo.EntityData.YFilter = sessionDetailInfo.YFilter
    sessionDetailInfo.EntityData.YangName = "session-detail-info"
    sessionDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailInfo.EntityData.ParentYangName = "outgoing-connections"
    sessionDetailInfo.EntityData.SegmentPath = "session-detail-info"
    sessionDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailInfo.EntityData.Children = make(map[string]types.YChild)
    sessionDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDetailInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionDetailInfo.SessionId}
    sessionDetailInfo.EntityData.Leafs["key-exchange"] = types.YLeaf{"KeyExchange", sessionDetailInfo.KeyExchange}
    sessionDetailInfo.EntityData.Leafs["public-key"] = types.YLeaf{"PublicKey", sessionDetailInfo.PublicKey}
    sessionDetailInfo.EntityData.Leafs["in-cipher"] = types.YLeaf{"InCipher", sessionDetailInfo.InCipher}
    sessionDetailInfo.EntityData.Leafs["out-cipher"] = types.YLeaf{"OutCipher", sessionDetailInfo.OutCipher}
    sessionDetailInfo.EntityData.Leafs["in-mac"] = types.YLeaf{"InMac", sessionDetailInfo.InMac}
    sessionDetailInfo.EntityData.Leafs["out-mac"] = types.YLeaf{"OutMac", sessionDetailInfo.OutMac}
    return &(sessionDetailInfo.EntityData)
}

// Ssh
// ssh
type Ssh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Crypto SSH session.
    Session Ssh_Session
}

func (ssh *Ssh) GetEntityData() *types.CommonEntityData {
    ssh.EntityData.YFilter = ssh.YFilter
    ssh.EntityData.YangName = "ssh"
    ssh.EntityData.BundleName = "cisco_ios_xr"
    ssh.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-ssh-oper"
    ssh.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-ssh-oper:ssh"
    ssh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssh.EntityData.Children = make(map[string]types.YChild)
    ssh.EntityData.Children["session"] = types.YChild{"Session", &ssh.Session}
    ssh.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ssh.EntityData)
}

// Ssh_Session
// Crypto SSH session
type Ssh_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSH session rekey information.
    Rekey Ssh_Session_Rekey

    // SSH session brief information.
    Brief Ssh_Session_Brief

    // SSH session detail information.
    Detail Ssh_Session_Detail
}

func (session *Ssh_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "ssh"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["rekey"] = types.YChild{"Rekey", &session.Rekey}
    session.EntityData.Children["brief"] = types.YChild{"Brief", &session.Brief}
    session.EntityData.Children["detail"] = types.YChild{"Detail", &session.Detail}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// Ssh_Session_Rekey
// SSH session rekey information
type Ssh_Session_Rekey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Rekey_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh_Session_Rekey_OutgoingConnections
}

func (rekey *Ssh_Session_Rekey) GetEntityData() *types.CommonEntityData {
    rekey.EntityData.YFilter = rekey.YFilter
    rekey.EntityData.YangName = "rekey"
    rekey.EntityData.BundleName = "cisco_ios_xr"
    rekey.EntityData.ParentYangName = "session"
    rekey.EntityData.SegmentPath = "rekey"
    rekey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rekey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rekey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rekey.EntityData.Children = make(map[string]types.YChild)
    rekey.EntityData.Children["incoming-sessions"] = types.YChild{"IncomingSessions", &rekey.IncomingSessions}
    rekey.EntityData.Children["outgoing-connections"] = types.YChild{"OutgoingConnections", &rekey.OutgoingConnections}
    rekey.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rekey.EntityData)
}

// Ssh_Session_Rekey_IncomingSessions
// List of incoming sessions
type Ssh_Session_Rekey_IncomingSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session rekey info. The type is slice of
    // Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo.
    SessionRekeyInfo []Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo
}

func (incomingSessions *Ssh_Session_Rekey_IncomingSessions) GetEntityData() *types.CommonEntityData {
    incomingSessions.EntityData.YFilter = incomingSessions.YFilter
    incomingSessions.EntityData.YangName = "incoming-sessions"
    incomingSessions.EntityData.BundleName = "cisco_ios_xr"
    incomingSessions.EntityData.ParentYangName = "rekey"
    incomingSessions.EntityData.SegmentPath = "incoming-sessions"
    incomingSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingSessions.EntityData.Children = make(map[string]types.YChild)
    incomingSessions.EntityData.Children["session-rekey-info"] = types.YChild{"SessionRekeyInfo", nil}
    for i := range incomingSessions.SessionRekeyInfo {
        incomingSessions.EntityData.Children[types.GetSegmentPath(&incomingSessions.SessionRekeyInfo[i])] = types.YChild{"SessionRekeyInfo", &incomingSessions.SessionRekeyInfo[i]}
    }
    incomingSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incomingSessions.EntityData)
}

// Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo
// session rekey info
type Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo struct {
    EntityData types.CommonEntityData
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

func (sessionRekeyInfo *Ssh_Session_Rekey_IncomingSessions_SessionRekeyInfo) GetEntityData() *types.CommonEntityData {
    sessionRekeyInfo.EntityData.YFilter = sessionRekeyInfo.YFilter
    sessionRekeyInfo.EntityData.YangName = "session-rekey-info"
    sessionRekeyInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionRekeyInfo.EntityData.ParentYangName = "incoming-sessions"
    sessionRekeyInfo.EntityData.SegmentPath = "session-rekey-info"
    sessionRekeyInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionRekeyInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionRekeyInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionRekeyInfo.EntityData.Children = make(map[string]types.YChild)
    sessionRekeyInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionRekeyInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionRekeyInfo.SessionId}
    sessionRekeyInfo.EntityData.Leafs["session-rekey-count"] = types.YLeaf{"SessionRekeyCount", sessionRekeyInfo.SessionRekeyCount}
    sessionRekeyInfo.EntityData.Leafs["time-to-rekey"] = types.YLeaf{"TimeToRekey", sessionRekeyInfo.TimeToRekey}
    sessionRekeyInfo.EntityData.Leafs["volume-to-rekey"] = types.YLeaf{"VolumeToRekey", sessionRekeyInfo.VolumeToRekey}
    return &(sessionRekeyInfo.EntityData)
}

// Ssh_Session_Rekey_OutgoingConnections
// List of outgoing connections
type Ssh_Session_Rekey_OutgoingConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session rekey info. The type is slice of
    // Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo.
    SessionRekeyInfo []Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo
}

func (outgoingConnections *Ssh_Session_Rekey_OutgoingConnections) GetEntityData() *types.CommonEntityData {
    outgoingConnections.EntityData.YFilter = outgoingConnections.YFilter
    outgoingConnections.EntityData.YangName = "outgoing-connections"
    outgoingConnections.EntityData.BundleName = "cisco_ios_xr"
    outgoingConnections.EntityData.ParentYangName = "rekey"
    outgoingConnections.EntityData.SegmentPath = "outgoing-connections"
    outgoingConnections.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outgoingConnections.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outgoingConnections.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outgoingConnections.EntityData.Children = make(map[string]types.YChild)
    outgoingConnections.EntityData.Children["session-rekey-info"] = types.YChild{"SessionRekeyInfo", nil}
    for i := range outgoingConnections.SessionRekeyInfo {
        outgoingConnections.EntityData.Children[types.GetSegmentPath(&outgoingConnections.SessionRekeyInfo[i])] = types.YChild{"SessionRekeyInfo", &outgoingConnections.SessionRekeyInfo[i]}
    }
    outgoingConnections.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outgoingConnections.EntityData)
}

// Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo
// session rekey info
type Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo struct {
    EntityData types.CommonEntityData
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

func (sessionRekeyInfo *Ssh_Session_Rekey_OutgoingConnections_SessionRekeyInfo) GetEntityData() *types.CommonEntityData {
    sessionRekeyInfo.EntityData.YFilter = sessionRekeyInfo.YFilter
    sessionRekeyInfo.EntityData.YangName = "session-rekey-info"
    sessionRekeyInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionRekeyInfo.EntityData.ParentYangName = "outgoing-connections"
    sessionRekeyInfo.EntityData.SegmentPath = "session-rekey-info"
    sessionRekeyInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionRekeyInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionRekeyInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionRekeyInfo.EntityData.Children = make(map[string]types.YChild)
    sessionRekeyInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionRekeyInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionRekeyInfo.SessionId}
    sessionRekeyInfo.EntityData.Leafs["session-rekey-count"] = types.YLeaf{"SessionRekeyCount", sessionRekeyInfo.SessionRekeyCount}
    sessionRekeyInfo.EntityData.Leafs["time-to-rekey"] = types.YLeaf{"TimeToRekey", sessionRekeyInfo.TimeToRekey}
    sessionRekeyInfo.EntityData.Leafs["volume-to-rekey"] = types.YLeaf{"VolumeToRekey", sessionRekeyInfo.VolumeToRekey}
    return &(sessionRekeyInfo.EntityData)
}

// Ssh_Session_Brief
// SSH session brief information
type Ssh_Session_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Brief_IncomingSessions

    // List of outgoing sessions.
    OutgoingSessions Ssh_Session_Brief_OutgoingSessions
}

func (brief *Ssh_Session_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "session"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = make(map[string]types.YChild)
    brief.EntityData.Children["incoming-sessions"] = types.YChild{"IncomingSessions", &brief.IncomingSessions}
    brief.EntityData.Children["outgoing-sessions"] = types.YChild{"OutgoingSessions", &brief.OutgoingSessions}
    brief.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(brief.EntityData)
}

// Ssh_Session_Brief_IncomingSessions
// List of incoming sessions
type Ssh_Session_Brief_IncomingSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session brief info. The type is slice of
    // Ssh_Session_Brief_IncomingSessions_SessionBriefInfo.
    SessionBriefInfo []Ssh_Session_Brief_IncomingSessions_SessionBriefInfo
}

func (incomingSessions *Ssh_Session_Brief_IncomingSessions) GetEntityData() *types.CommonEntityData {
    incomingSessions.EntityData.YFilter = incomingSessions.YFilter
    incomingSessions.EntityData.YangName = "incoming-sessions"
    incomingSessions.EntityData.BundleName = "cisco_ios_xr"
    incomingSessions.EntityData.ParentYangName = "brief"
    incomingSessions.EntityData.SegmentPath = "incoming-sessions"
    incomingSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingSessions.EntityData.Children = make(map[string]types.YChild)
    incomingSessions.EntityData.Children["session-brief-info"] = types.YChild{"SessionBriefInfo", nil}
    for i := range incomingSessions.SessionBriefInfo {
        incomingSessions.EntityData.Children[types.GetSegmentPath(&incomingSessions.SessionBriefInfo[i])] = types.YChild{"SessionBriefInfo", &incomingSessions.SessionBriefInfo[i]}
    }
    incomingSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incomingSessions.EntityData)
}

// Ssh_Session_Brief_IncomingSessions_SessionBriefInfo
// session brief info
type Ssh_Session_Brief_IncomingSessions_SessionBriefInfo struct {
    EntityData types.CommonEntityData
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
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (sessionBriefInfo *Ssh_Session_Brief_IncomingSessions_SessionBriefInfo) GetEntityData() *types.CommonEntityData {
    sessionBriefInfo.EntityData.YFilter = sessionBriefInfo.YFilter
    sessionBriefInfo.EntityData.YangName = "session-brief-info"
    sessionBriefInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionBriefInfo.EntityData.ParentYangName = "incoming-sessions"
    sessionBriefInfo.EntityData.SegmentPath = "session-brief-info"
    sessionBriefInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionBriefInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionBriefInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionBriefInfo.EntityData.Children = make(map[string]types.YChild)
    sessionBriefInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionBriefInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionBriefInfo.SessionId}
    sessionBriefInfo.EntityData.Leafs["channel-id"] = types.YLeaf{"ChannelId", sessionBriefInfo.ChannelId}
    sessionBriefInfo.EntityData.Leafs["vty-assigned"] = types.YLeaf{"VtyAssigned", sessionBriefInfo.VtyAssigned}
    sessionBriefInfo.EntityData.Leafs["vty-line-number"] = types.YLeaf{"VtyLineNumber", sessionBriefInfo.VtyLineNumber}
    sessionBriefInfo.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", sessionBriefInfo.NodeName}
    sessionBriefInfo.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", sessionBriefInfo.SessionState}
    sessionBriefInfo.EntityData.Leafs["user-id"] = types.YLeaf{"UserId", sessionBriefInfo.UserId}
    sessionBriefInfo.EntityData.Leafs["host-address"] = types.YLeaf{"HostAddress", sessionBriefInfo.HostAddress}
    sessionBriefInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", sessionBriefInfo.Version}
    sessionBriefInfo.EntityData.Leafs["authentication-type"] = types.YLeaf{"AuthenticationType", sessionBriefInfo.AuthenticationType}
    sessionBriefInfo.EntityData.Leafs["connection-type"] = types.YLeaf{"ConnectionType", sessionBriefInfo.ConnectionType}
    return &(sessionBriefInfo.EntityData)
}

// Ssh_Session_Brief_OutgoingSessions
// List of outgoing sessions
type Ssh_Session_Brief_OutgoingSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session brief info. The type is slice of
    // Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo.
    SessionBriefInfo []Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo
}

func (outgoingSessions *Ssh_Session_Brief_OutgoingSessions) GetEntityData() *types.CommonEntityData {
    outgoingSessions.EntityData.YFilter = outgoingSessions.YFilter
    outgoingSessions.EntityData.YangName = "outgoing-sessions"
    outgoingSessions.EntityData.BundleName = "cisco_ios_xr"
    outgoingSessions.EntityData.ParentYangName = "brief"
    outgoingSessions.EntityData.SegmentPath = "outgoing-sessions"
    outgoingSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outgoingSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outgoingSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outgoingSessions.EntityData.Children = make(map[string]types.YChild)
    outgoingSessions.EntityData.Children["session-brief-info"] = types.YChild{"SessionBriefInfo", nil}
    for i := range outgoingSessions.SessionBriefInfo {
        outgoingSessions.EntityData.Children[types.GetSegmentPath(&outgoingSessions.SessionBriefInfo[i])] = types.YChild{"SessionBriefInfo", &outgoingSessions.SessionBriefInfo[i]}
    }
    outgoingSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outgoingSessions.EntityData)
}

// Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo
// session brief info
type Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo struct {
    EntityData types.CommonEntityData
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
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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

func (sessionBriefInfo *Ssh_Session_Brief_OutgoingSessions_SessionBriefInfo) GetEntityData() *types.CommonEntityData {
    sessionBriefInfo.EntityData.YFilter = sessionBriefInfo.YFilter
    sessionBriefInfo.EntityData.YangName = "session-brief-info"
    sessionBriefInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionBriefInfo.EntityData.ParentYangName = "outgoing-sessions"
    sessionBriefInfo.EntityData.SegmentPath = "session-brief-info"
    sessionBriefInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionBriefInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionBriefInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionBriefInfo.EntityData.Children = make(map[string]types.YChild)
    sessionBriefInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionBriefInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionBriefInfo.SessionId}
    sessionBriefInfo.EntityData.Leafs["channel-id"] = types.YLeaf{"ChannelId", sessionBriefInfo.ChannelId}
    sessionBriefInfo.EntityData.Leafs["vty-assigned"] = types.YLeaf{"VtyAssigned", sessionBriefInfo.VtyAssigned}
    sessionBriefInfo.EntityData.Leafs["vty-line-number"] = types.YLeaf{"VtyLineNumber", sessionBriefInfo.VtyLineNumber}
    sessionBriefInfo.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", sessionBriefInfo.NodeName}
    sessionBriefInfo.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", sessionBriefInfo.SessionState}
    sessionBriefInfo.EntityData.Leafs["user-id"] = types.YLeaf{"UserId", sessionBriefInfo.UserId}
    sessionBriefInfo.EntityData.Leafs["host-address"] = types.YLeaf{"HostAddress", sessionBriefInfo.HostAddress}
    sessionBriefInfo.EntityData.Leafs["version"] = types.YLeaf{"Version", sessionBriefInfo.Version}
    sessionBriefInfo.EntityData.Leafs["authentication-type"] = types.YLeaf{"AuthenticationType", sessionBriefInfo.AuthenticationType}
    sessionBriefInfo.EntityData.Leafs["connection-type"] = types.YLeaf{"ConnectionType", sessionBriefInfo.ConnectionType}
    return &(sessionBriefInfo.EntityData)
}

// Ssh_Session_Detail
// SSH session detail information
type Ssh_Session_Detail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of incoming sessions.
    IncomingSessions Ssh_Session_Detail_IncomingSessions

    // List of outgoing connections.
    OutgoingConnections Ssh_Session_Detail_OutgoingConnections
}

func (detail *Ssh_Session_Detail) GetEntityData() *types.CommonEntityData {
    detail.EntityData.YFilter = detail.YFilter
    detail.EntityData.YangName = "detail"
    detail.EntityData.BundleName = "cisco_ios_xr"
    detail.EntityData.ParentYangName = "session"
    detail.EntityData.SegmentPath = "detail"
    detail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detail.EntityData.Children = make(map[string]types.YChild)
    detail.EntityData.Children["incoming-sessions"] = types.YChild{"IncomingSessions", &detail.IncomingSessions}
    detail.EntityData.Children["outgoing-connections"] = types.YChild{"OutgoingConnections", &detail.OutgoingConnections}
    detail.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detail.EntityData)
}

// Ssh_Session_Detail_IncomingSessions
// List of incoming sessions
type Ssh_Session_Detail_IncomingSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh_Session_Detail_IncomingSessions_SessionDetailInfo.
    SessionDetailInfo []Ssh_Session_Detail_IncomingSessions_SessionDetailInfo
}

func (incomingSessions *Ssh_Session_Detail_IncomingSessions) GetEntityData() *types.CommonEntityData {
    incomingSessions.EntityData.YFilter = incomingSessions.YFilter
    incomingSessions.EntityData.YangName = "incoming-sessions"
    incomingSessions.EntityData.BundleName = "cisco_ios_xr"
    incomingSessions.EntityData.ParentYangName = "detail"
    incomingSessions.EntityData.SegmentPath = "incoming-sessions"
    incomingSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incomingSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incomingSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incomingSessions.EntityData.Children = make(map[string]types.YChild)
    incomingSessions.EntityData.Children["session-detail-info"] = types.YChild{"SessionDetailInfo", nil}
    for i := range incomingSessions.SessionDetailInfo {
        incomingSessions.EntityData.Children[types.GetSegmentPath(&incomingSessions.SessionDetailInfo[i])] = types.YChild{"SessionDetailInfo", &incomingSessions.SessionDetailInfo[i]}
    }
    incomingSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(incomingSessions.EntityData)
}

// Ssh_Session_Detail_IncomingSessions_SessionDetailInfo
// session detail info
type Ssh_Session_Detail_IncomingSessions_SessionDetailInfo struct {
    EntityData types.CommonEntityData
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

func (sessionDetailInfo *Ssh_Session_Detail_IncomingSessions_SessionDetailInfo) GetEntityData() *types.CommonEntityData {
    sessionDetailInfo.EntityData.YFilter = sessionDetailInfo.YFilter
    sessionDetailInfo.EntityData.YangName = "session-detail-info"
    sessionDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailInfo.EntityData.ParentYangName = "incoming-sessions"
    sessionDetailInfo.EntityData.SegmentPath = "session-detail-info"
    sessionDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailInfo.EntityData.Children = make(map[string]types.YChild)
    sessionDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDetailInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionDetailInfo.SessionId}
    sessionDetailInfo.EntityData.Leafs["key-exchange"] = types.YLeaf{"KeyExchange", sessionDetailInfo.KeyExchange}
    sessionDetailInfo.EntityData.Leafs["public-key"] = types.YLeaf{"PublicKey", sessionDetailInfo.PublicKey}
    sessionDetailInfo.EntityData.Leafs["in-cipher"] = types.YLeaf{"InCipher", sessionDetailInfo.InCipher}
    sessionDetailInfo.EntityData.Leafs["out-cipher"] = types.YLeaf{"OutCipher", sessionDetailInfo.OutCipher}
    sessionDetailInfo.EntityData.Leafs["in-mac"] = types.YLeaf{"InMac", sessionDetailInfo.InMac}
    sessionDetailInfo.EntityData.Leafs["out-mac"] = types.YLeaf{"OutMac", sessionDetailInfo.OutMac}
    return &(sessionDetailInfo.EntityData)
}

// Ssh_Session_Detail_OutgoingConnections
// List of outgoing connections
type Ssh_Session_Detail_OutgoingConnections struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session detail info. The type is slice of
    // Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo.
    SessionDetailInfo []Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo
}

func (outgoingConnections *Ssh_Session_Detail_OutgoingConnections) GetEntityData() *types.CommonEntityData {
    outgoingConnections.EntityData.YFilter = outgoingConnections.YFilter
    outgoingConnections.EntityData.YangName = "outgoing-connections"
    outgoingConnections.EntityData.BundleName = "cisco_ios_xr"
    outgoingConnections.EntityData.ParentYangName = "detail"
    outgoingConnections.EntityData.SegmentPath = "outgoing-connections"
    outgoingConnections.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outgoingConnections.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outgoingConnections.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outgoingConnections.EntityData.Children = make(map[string]types.YChild)
    outgoingConnections.EntityData.Children["session-detail-info"] = types.YChild{"SessionDetailInfo", nil}
    for i := range outgoingConnections.SessionDetailInfo {
        outgoingConnections.EntityData.Children[types.GetSegmentPath(&outgoingConnections.SessionDetailInfo[i])] = types.YChild{"SessionDetailInfo", &outgoingConnections.SessionDetailInfo[i]}
    }
    outgoingConnections.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(outgoingConnections.EntityData)
}

// Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo
// session detail info
type Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo struct {
    EntityData types.CommonEntityData
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

func (sessionDetailInfo *Ssh_Session_Detail_OutgoingConnections_SessionDetailInfo) GetEntityData() *types.CommonEntityData {
    sessionDetailInfo.EntityData.YFilter = sessionDetailInfo.YFilter
    sessionDetailInfo.EntityData.YangName = "session-detail-info"
    sessionDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    sessionDetailInfo.EntityData.ParentYangName = "outgoing-connections"
    sessionDetailInfo.EntityData.SegmentPath = "session-detail-info"
    sessionDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetailInfo.EntityData.Children = make(map[string]types.YChild)
    sessionDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDetailInfo.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", sessionDetailInfo.SessionId}
    sessionDetailInfo.EntityData.Leafs["key-exchange"] = types.YLeaf{"KeyExchange", sessionDetailInfo.KeyExchange}
    sessionDetailInfo.EntityData.Leafs["public-key"] = types.YLeaf{"PublicKey", sessionDetailInfo.PublicKey}
    sessionDetailInfo.EntityData.Leafs["in-cipher"] = types.YLeaf{"InCipher", sessionDetailInfo.InCipher}
    sessionDetailInfo.EntityData.Leafs["out-cipher"] = types.YLeaf{"OutCipher", sessionDetailInfo.OutCipher}
    sessionDetailInfo.EntityData.Leafs["in-mac"] = types.YLeaf{"InMac", sessionDetailInfo.InMac}
    sessionDetailInfo.EntityData.Leafs["out-mac"] = types.YLeaf{"OutMac", sessionDetailInfo.OutMac}
    return &(sessionDetailInfo.EntityData)
}

