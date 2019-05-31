// This module contains a collection of YANG definitions
// for Cisco IOS-XR ikev2 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ik-ev2: IKEv2 operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ikev2_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ikev2_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ikev2-oper ik-ev2}", reflect.TypeOf(IkEv2{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ikev2-oper:ik-ev2", reflect.TypeOf(IkEv2{}))
}

// IkEv2
// IKEv2 operational data
type IkEv2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node IKEv2 operational data.
    Nodes IkEv2_Nodes
}

func (ikEv2 *IkEv2) GetEntityData() *types.CommonEntityData {
    ikEv2.EntityData.YFilter = ikEv2.YFilter
    ikEv2.EntityData.YangName = "ik-ev2"
    ikEv2.EntityData.BundleName = "cisco_ios_xr"
    ikEv2.EntityData.ParentYangName = "Cisco-IOS-XR-ikev2-oper"
    ikEv2.EntityData.SegmentPath = "Cisco-IOS-XR-ikev2-oper:ik-ev2"
    ikEv2.EntityData.AbsolutePath = ikEv2.EntityData.SegmentPath
    ikEv2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ikEv2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ikEv2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ikEv2.EntityData.Children = types.NewOrderedMap()
    ikEv2.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ikEv2.Nodes})
    ikEv2.EntityData.Leafs = types.NewOrderedMap()

    ikEv2.EntityData.YListKeys = []string {}

    return &(ikEv2.EntityData)
}

// IkEv2_Nodes
// Per node IKEv2 operational data
type IkEv2_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 operational data for a particular node. The type is slice of
    // IkEv2_Nodes_Node.
    Node []*IkEv2_Nodes_Node
}

func (nodes *IkEv2_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ik-ev2"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// IkEv2_Nodes_Node
// IKEv2 operational data for a particular node
type IkEv2_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // IKEv2 Session data.
    Session IkEv2_Nodes_Node_Session

    // IKEv2 SA data.
    Sa IkEv2_Nodes_Node_Sa

    // IKEv2 policies on this node.
    Policies IkEv2_Nodes_Node_Policies

    // IKEv2 proposals on this node.
    Proposals IkEv2_Nodes_Node_Proposals

    // IKEv2 profiles on this node.
    Profiles IkEv2_Nodes_Node_Profiles

    // IKEv2 keyrings on this node.
    Keyrings IkEv2_Nodes_Node_Keyrings
}

func (node *IkEv2_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("session", types.YChild{"Session", &node.Session})
    node.EntityData.Children.Append("sa", types.YChild{"Sa", &node.Sa})
    node.EntityData.Children.Append("policies", types.YChild{"Policies", &node.Policies})
    node.EntityData.Children.Append("proposals", types.YChild{"Proposals", &node.Proposals})
    node.EntityData.Children.Append("profiles", types.YChild{"Profiles", &node.Profiles})
    node.EntityData.Children.Append("keyrings", types.YChild{"Keyrings", &node.Keyrings})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// IkEv2_Nodes_Node_Session
// IKEv2 Session data
type IkEv2_Nodes_Node_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session List. The type is slice of IkEv2_Nodes_Node_Session_Session.
    Session []*IkEv2_Nodes_Node_Session_Session
}

func (session *IkEv2_Nodes_Node_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "node"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range session.Session {
        types.SetYListKey(session.Session[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.Session[i]), types.YChild{"Session", session.Session[i]})
    }
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// IkEv2_Nodes_Node_Session_Session
// Session List
type IkEv2_Nodes_Node_Session_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Session Status. The type is string.
    SessionStatus interface{}

    // IKE Count. The type is interface{} with range: 0..4294967295.
    IkeCnt interface{}

    // Child Count. The type is interface{} with range: 0..4294967295.
    ChildCnt interface{}

    // SA List. The type is slice of IkEv2_Nodes_Node_Session_Session_Sa.
    Sa []*IkEv2_Nodes_Node_Session_Session_Sa

    // Child SA List. The type is slice of
    // IkEv2_Nodes_Node_Session_Session_ChildSa.
    ChildSa []*IkEv2_Nodes_Node_Session_Session_ChildSa
}

func (session *IkEv2_Nodes_Node_Session_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "session"
    session.EntityData.SegmentPath = "session" + types.AddNoKeyToken(session)
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/session/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("sa", types.YChild{"Sa", nil})
    for i := range session.Sa {
        types.SetYListKey(session.Sa[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.Sa[i]), types.YChild{"Sa", session.Sa[i]})
    }
    session.EntityData.Children.Append("child-sa", types.YChild{"ChildSa", nil})
    for i := range session.ChildSa {
        types.SetYListKey(session.ChildSa[i], i)
        session.EntityData.Children.Append(types.GetSegmentPath(session.ChildSa[i]), types.YChild{"ChildSa", session.ChildSa[i]})
    }
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("session-status", types.YLeaf{"SessionStatus", session.SessionStatus})
    session.EntityData.Leafs.Append("ike-cnt", types.YLeaf{"IkeCnt", session.IkeCnt})
    session.EntityData.Leafs.Append("child-cnt", types.YLeaf{"ChildCnt", session.ChildCnt})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// IkEv2_Nodes_Node_Session_Session_Sa
// SA List
type IkEv2_Nodes_Node_Session_Session_Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Local Addr Port. The type is string.
    LocalAddrPort interface{}

    // Remote Addr Port. The type is string.
    RemoteAddrPort interface{}

    // State. The type is string.
    State interface{}

    // Encryption. The type is string.
    Encr interface{}

    // Keysize. The type is interface{} with range: 0..4294967295.
    Keysize interface{}

    // PRF. The type is string.
    Prf interface{}

    // Hash. The type is string.
    Hash interface{}

    // DH Group. The type is interface{} with range: 0..4294967295.
    DhGroup interface{}

    // Auth Sign. The type is string.
    AuthSign interface{}

    // Auth Verify. The type is string.
    AuthVerify interface{}

    // Life-Active Time. The type is string.
    LifeActive interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Status Description. The type is string.
    StatusDesc interface{}

    // Local SPI. The type is string.
    LocalSpi interface{}

    // Remote SPI. The type is string.
    RemoteSpi interface{}

    // Local ID. The type is string.
    LocalId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Sa Initiator. The type is bool.
    SaInitiator interface{}
}

func (sa *IkEv2_Nodes_Node_Session_Session_Sa) GetEntityData() *types.CommonEntityData {
    sa.EntityData.YFilter = sa.YFilter
    sa.EntityData.YangName = "sa"
    sa.EntityData.BundleName = "cisco_ios_xr"
    sa.EntityData.ParentYangName = "session"
    sa.EntityData.SegmentPath = "sa" + types.AddNoKeyToken(sa)
    sa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/session/session/" + sa.EntityData.SegmentPath
    sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sa.EntityData.Children = types.NewOrderedMap()
    sa.EntityData.Leafs = types.NewOrderedMap()
    sa.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", sa.TunnelId})
    sa.EntityData.Leafs.Append("local-addr-port", types.YLeaf{"LocalAddrPort", sa.LocalAddrPort})
    sa.EntityData.Leafs.Append("remote-addr-port", types.YLeaf{"RemoteAddrPort", sa.RemoteAddrPort})
    sa.EntityData.Leafs.Append("state", types.YLeaf{"State", sa.State})
    sa.EntityData.Leafs.Append("encr", types.YLeaf{"Encr", sa.Encr})
    sa.EntityData.Leafs.Append("keysize", types.YLeaf{"Keysize", sa.Keysize})
    sa.EntityData.Leafs.Append("prf", types.YLeaf{"Prf", sa.Prf})
    sa.EntityData.Leafs.Append("hash", types.YLeaf{"Hash", sa.Hash})
    sa.EntityData.Leafs.Append("dh-group", types.YLeaf{"DhGroup", sa.DhGroup})
    sa.EntityData.Leafs.Append("auth-sign", types.YLeaf{"AuthSign", sa.AuthSign})
    sa.EntityData.Leafs.Append("auth-verify", types.YLeaf{"AuthVerify", sa.AuthVerify})
    sa.EntityData.Leafs.Append("life-active", types.YLeaf{"LifeActive", sa.LifeActive})
    sa.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", sa.SessionId})
    sa.EntityData.Leafs.Append("status-desc", types.YLeaf{"StatusDesc", sa.StatusDesc})
    sa.EntityData.Leafs.Append("local-spi", types.YLeaf{"LocalSpi", sa.LocalSpi})
    sa.EntityData.Leafs.Append("remote-spi", types.YLeaf{"RemoteSpi", sa.RemoteSpi})
    sa.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", sa.LocalId})
    sa.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", sa.RemoteId})
    sa.EntityData.Leafs.Append("sa-initiator", types.YLeaf{"SaInitiator", sa.SaInitiator})

    sa.EntityData.YListKeys = []string {}

    return &(sa.EntityData)
}

// IkEv2_Nodes_Node_Session_Session_ChildSa
// Child SA List
type IkEv2_Nodes_Node_Session_Session_ChildSa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Local Selector. The type is string.
    LocalSelector interface{}

    // Remote Selector. The type is string.
    RemoteSelector interface{}

    // ESP SPI In Out. The type is string.
    EspSpiInOut interface{}

    // Encryption. The type is string.
    Encr interface{}

    // Keysize. The type is interface{} with range: 0..4294967295.
    Keysize interface{}

    // ESP HMAC. The type is string.
    Hmac interface{}
}

func (childSa *IkEv2_Nodes_Node_Session_Session_ChildSa) GetEntityData() *types.CommonEntityData {
    childSa.EntityData.YFilter = childSa.YFilter
    childSa.EntityData.YangName = "child-sa"
    childSa.EntityData.BundleName = "cisco_ios_xr"
    childSa.EntityData.ParentYangName = "session"
    childSa.EntityData.SegmentPath = "child-sa" + types.AddNoKeyToken(childSa)
    childSa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/session/session/" + childSa.EntityData.SegmentPath
    childSa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    childSa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    childSa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    childSa.EntityData.Children = types.NewOrderedMap()
    childSa.EntityData.Leafs = types.NewOrderedMap()
    childSa.EntityData.Leafs.Append("local-selector", types.YLeaf{"LocalSelector", childSa.LocalSelector})
    childSa.EntityData.Leafs.Append("remote-selector", types.YLeaf{"RemoteSelector", childSa.RemoteSelector})
    childSa.EntityData.Leafs.Append("esp-spi-in-out", types.YLeaf{"EspSpiInOut", childSa.EspSpiInOut})
    childSa.EntityData.Leafs.Append("encr", types.YLeaf{"Encr", childSa.Encr})
    childSa.EntityData.Leafs.Append("keysize", types.YLeaf{"Keysize", childSa.Keysize})
    childSa.EntityData.Leafs.Append("hmac", types.YLeaf{"Hmac", childSa.Hmac})

    childSa.EntityData.YListKeys = []string {}

    return &(childSa.EntityData)
}

// IkEv2_Nodes_Node_Sa
// IKEv2 SA data
type IkEv2_Nodes_Node_Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 SA lookup on local IP.
    LocalV4 IkEv2_Nodes_Node_Sa_LocalV4

    // IKEv2 SA lookup on remote IP.
    RemoteV4 IkEv2_Nodes_Node_Sa_RemoteV4

    // IKEv2 SA all data.
    All IkEv2_Nodes_Node_Sa_All
}

func (sa *IkEv2_Nodes_Node_Sa) GetEntityData() *types.CommonEntityData {
    sa.EntityData.YFilter = sa.YFilter
    sa.EntityData.YangName = "sa"
    sa.EntityData.BundleName = "cisco_ios_xr"
    sa.EntityData.ParentYangName = "node"
    sa.EntityData.SegmentPath = "sa"
    sa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + sa.EntityData.SegmentPath
    sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sa.EntityData.Children = types.NewOrderedMap()
    sa.EntityData.Children.Append("local-v4", types.YChild{"LocalV4", &sa.LocalV4})
    sa.EntityData.Children.Append("remote-v4", types.YChild{"RemoteV4", &sa.RemoteV4})
    sa.EntityData.Children.Append("all", types.YChild{"All", &sa.All})
    sa.EntityData.Leafs = types.NewOrderedMap()

    sa.EntityData.YListKeys = []string {}

    return &(sa.EntityData)
}

// IkEv2_Nodes_Node_Sa_LocalV4
// IKEv2 SA lookup on local IP
type IkEv2_Nodes_Node_Sa_LocalV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 SA data based on address. The type is slice of
    // IkEv2_Nodes_Node_Sa_LocalV4_Ip.
    Ip []*IkEv2_Nodes_Node_Sa_LocalV4_Ip
}

func (localV4 *IkEv2_Nodes_Node_Sa_LocalV4) GetEntityData() *types.CommonEntityData {
    localV4.EntityData.YFilter = localV4.YFilter
    localV4.EntityData.YangName = "local-v4"
    localV4.EntityData.BundleName = "cisco_ios_xr"
    localV4.EntityData.ParentYangName = "sa"
    localV4.EntityData.SegmentPath = "local-v4"
    localV4.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/" + localV4.EntityData.SegmentPath
    localV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localV4.EntityData.Children = types.NewOrderedMap()
    localV4.EntityData.Children.Append("ip", types.YChild{"Ip", nil})
    for i := range localV4.Ip {
        localV4.EntityData.Children.Append(types.GetSegmentPath(localV4.Ip[i]), types.YChild{"Ip", localV4.Ip[i]})
    }
    localV4.EntityData.Leafs = types.NewOrderedMap()

    localV4.EntityData.YListKeys = []string {}

    return &(localV4.EntityData)
}

// IkEv2_Nodes_Node_Sa_LocalV4_Ip
// IKEv2 SA data based on address
type IkEv2_Nodes_Node_Sa_LocalV4_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // SA list. The type is slice of IkEv2_Nodes_Node_Sa_LocalV4_Ip_Sa.
    Sa []*IkEv2_Nodes_Node_Sa_LocalV4_Ip_Sa
}

func (ip *IkEv2_Nodes_Node_Sa_LocalV4_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "local-v4"
    ip.EntityData.SegmentPath = "ip" + types.AddKeyToken(ip.Address, "address")
    ip.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/local-v4/" + ip.EntityData.SegmentPath
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Children.Append("sa", types.YChild{"Sa", nil})
    for i := range ip.Sa {
        types.SetYListKey(ip.Sa[i], i)
        ip.EntityData.Children.Append(types.GetSegmentPath(ip.Sa[i]), types.YChild{"Sa", ip.Sa[i]})
    }
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("address", types.YLeaf{"Address", ip.Address})

    ip.EntityData.YListKeys = []string {"Address"}

    return &(ip.EntityData)
}

// IkEv2_Nodes_Node_Sa_LocalV4_Ip_Sa
// SA list
type IkEv2_Nodes_Node_Sa_LocalV4_Ip_Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Local Addr Port. The type is string.
    LocalAddrPort interface{}

    // Remote Addr Port. The type is string.
    RemoteAddrPort interface{}

    // State. The type is string.
    State interface{}

    // Encryption. The type is string.
    Encr interface{}

    // Keysize. The type is interface{} with range: 0..4294967295.
    Keysize interface{}

    // PRF. The type is string.
    Prf interface{}

    // Hash. The type is string.
    Hash interface{}

    // DH Group. The type is interface{} with range: 0..4294967295.
    DhGroup interface{}

    // Auth Sign. The type is string.
    AuthSign interface{}

    // Auth Verify. The type is string.
    AuthVerify interface{}

    // Life-Active Time. The type is string.
    LifeActive interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Status Description. The type is string.
    StatusDesc interface{}

    // Local SPI. The type is string.
    LocalSpi interface{}

    // Remote SPI. The type is string.
    RemoteSpi interface{}

    // Local ID. The type is string.
    LocalId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Sa Initiator. The type is bool.
    SaInitiator interface{}
}

func (sa *IkEv2_Nodes_Node_Sa_LocalV4_Ip_Sa) GetEntityData() *types.CommonEntityData {
    sa.EntityData.YFilter = sa.YFilter
    sa.EntityData.YangName = "sa"
    sa.EntityData.BundleName = "cisco_ios_xr"
    sa.EntityData.ParentYangName = "ip"
    sa.EntityData.SegmentPath = "sa" + types.AddNoKeyToken(sa)
    sa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/local-v4/ip/" + sa.EntityData.SegmentPath
    sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sa.EntityData.Children = types.NewOrderedMap()
    sa.EntityData.Leafs = types.NewOrderedMap()
    sa.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", sa.TunnelId})
    sa.EntityData.Leafs.Append("local-addr-port", types.YLeaf{"LocalAddrPort", sa.LocalAddrPort})
    sa.EntityData.Leafs.Append("remote-addr-port", types.YLeaf{"RemoteAddrPort", sa.RemoteAddrPort})
    sa.EntityData.Leafs.Append("state", types.YLeaf{"State", sa.State})
    sa.EntityData.Leafs.Append("encr", types.YLeaf{"Encr", sa.Encr})
    sa.EntityData.Leafs.Append("keysize", types.YLeaf{"Keysize", sa.Keysize})
    sa.EntityData.Leafs.Append("prf", types.YLeaf{"Prf", sa.Prf})
    sa.EntityData.Leafs.Append("hash", types.YLeaf{"Hash", sa.Hash})
    sa.EntityData.Leafs.Append("dh-group", types.YLeaf{"DhGroup", sa.DhGroup})
    sa.EntityData.Leafs.Append("auth-sign", types.YLeaf{"AuthSign", sa.AuthSign})
    sa.EntityData.Leafs.Append("auth-verify", types.YLeaf{"AuthVerify", sa.AuthVerify})
    sa.EntityData.Leafs.Append("life-active", types.YLeaf{"LifeActive", sa.LifeActive})
    sa.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", sa.SessionId})
    sa.EntityData.Leafs.Append("status-desc", types.YLeaf{"StatusDesc", sa.StatusDesc})
    sa.EntityData.Leafs.Append("local-spi", types.YLeaf{"LocalSpi", sa.LocalSpi})
    sa.EntityData.Leafs.Append("remote-spi", types.YLeaf{"RemoteSpi", sa.RemoteSpi})
    sa.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", sa.LocalId})
    sa.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", sa.RemoteId})
    sa.EntityData.Leafs.Append("sa-initiator", types.YLeaf{"SaInitiator", sa.SaInitiator})

    sa.EntityData.YListKeys = []string {}

    return &(sa.EntityData)
}

// IkEv2_Nodes_Node_Sa_RemoteV4
// IKEv2 SA lookup on remote IP
type IkEv2_Nodes_Node_Sa_RemoteV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 SA data based on address. The type is slice of
    // IkEv2_Nodes_Node_Sa_RemoteV4_Ip.
    Ip []*IkEv2_Nodes_Node_Sa_RemoteV4_Ip
}

func (remoteV4 *IkEv2_Nodes_Node_Sa_RemoteV4) GetEntityData() *types.CommonEntityData {
    remoteV4.EntityData.YFilter = remoteV4.YFilter
    remoteV4.EntityData.YangName = "remote-v4"
    remoteV4.EntityData.BundleName = "cisco_ios_xr"
    remoteV4.EntityData.ParentYangName = "sa"
    remoteV4.EntityData.SegmentPath = "remote-v4"
    remoteV4.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/" + remoteV4.EntityData.SegmentPath
    remoteV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteV4.EntityData.Children = types.NewOrderedMap()
    remoteV4.EntityData.Children.Append("ip", types.YChild{"Ip", nil})
    for i := range remoteV4.Ip {
        remoteV4.EntityData.Children.Append(types.GetSegmentPath(remoteV4.Ip[i]), types.YChild{"Ip", remoteV4.Ip[i]})
    }
    remoteV4.EntityData.Leafs = types.NewOrderedMap()

    remoteV4.EntityData.YListKeys = []string {}

    return &(remoteV4.EntityData)
}

// IkEv2_Nodes_Node_Sa_RemoteV4_Ip
// IKEv2 SA data based on address
type IkEv2_Nodes_Node_Sa_RemoteV4_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // SA list. The type is slice of IkEv2_Nodes_Node_Sa_RemoteV4_Ip_Sa.
    Sa []*IkEv2_Nodes_Node_Sa_RemoteV4_Ip_Sa
}

func (ip *IkEv2_Nodes_Node_Sa_RemoteV4_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "remote-v4"
    ip.EntityData.SegmentPath = "ip" + types.AddKeyToken(ip.Address, "address")
    ip.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/remote-v4/" + ip.EntityData.SegmentPath
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Children.Append("sa", types.YChild{"Sa", nil})
    for i := range ip.Sa {
        types.SetYListKey(ip.Sa[i], i)
        ip.EntityData.Children.Append(types.GetSegmentPath(ip.Sa[i]), types.YChild{"Sa", ip.Sa[i]})
    }
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("address", types.YLeaf{"Address", ip.Address})

    ip.EntityData.YListKeys = []string {"Address"}

    return &(ip.EntityData)
}

// IkEv2_Nodes_Node_Sa_RemoteV4_Ip_Sa
// SA list
type IkEv2_Nodes_Node_Sa_RemoteV4_Ip_Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Local Addr Port. The type is string.
    LocalAddrPort interface{}

    // Remote Addr Port. The type is string.
    RemoteAddrPort interface{}

    // State. The type is string.
    State interface{}

    // Encryption. The type is string.
    Encr interface{}

    // Keysize. The type is interface{} with range: 0..4294967295.
    Keysize interface{}

    // PRF. The type is string.
    Prf interface{}

    // Hash. The type is string.
    Hash interface{}

    // DH Group. The type is interface{} with range: 0..4294967295.
    DhGroup interface{}

    // Auth Sign. The type is string.
    AuthSign interface{}

    // Auth Verify. The type is string.
    AuthVerify interface{}

    // Life-Active Time. The type is string.
    LifeActive interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Status Description. The type is string.
    StatusDesc interface{}

    // Local SPI. The type is string.
    LocalSpi interface{}

    // Remote SPI. The type is string.
    RemoteSpi interface{}

    // Local ID. The type is string.
    LocalId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Sa Initiator. The type is bool.
    SaInitiator interface{}
}

func (sa *IkEv2_Nodes_Node_Sa_RemoteV4_Ip_Sa) GetEntityData() *types.CommonEntityData {
    sa.EntityData.YFilter = sa.YFilter
    sa.EntityData.YangName = "sa"
    sa.EntityData.BundleName = "cisco_ios_xr"
    sa.EntityData.ParentYangName = "ip"
    sa.EntityData.SegmentPath = "sa" + types.AddNoKeyToken(sa)
    sa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/remote-v4/ip/" + sa.EntityData.SegmentPath
    sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sa.EntityData.Children = types.NewOrderedMap()
    sa.EntityData.Leafs = types.NewOrderedMap()
    sa.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", sa.TunnelId})
    sa.EntityData.Leafs.Append("local-addr-port", types.YLeaf{"LocalAddrPort", sa.LocalAddrPort})
    sa.EntityData.Leafs.Append("remote-addr-port", types.YLeaf{"RemoteAddrPort", sa.RemoteAddrPort})
    sa.EntityData.Leafs.Append("state", types.YLeaf{"State", sa.State})
    sa.EntityData.Leafs.Append("encr", types.YLeaf{"Encr", sa.Encr})
    sa.EntityData.Leafs.Append("keysize", types.YLeaf{"Keysize", sa.Keysize})
    sa.EntityData.Leafs.Append("prf", types.YLeaf{"Prf", sa.Prf})
    sa.EntityData.Leafs.Append("hash", types.YLeaf{"Hash", sa.Hash})
    sa.EntityData.Leafs.Append("dh-group", types.YLeaf{"DhGroup", sa.DhGroup})
    sa.EntityData.Leafs.Append("auth-sign", types.YLeaf{"AuthSign", sa.AuthSign})
    sa.EntityData.Leafs.Append("auth-verify", types.YLeaf{"AuthVerify", sa.AuthVerify})
    sa.EntityData.Leafs.Append("life-active", types.YLeaf{"LifeActive", sa.LifeActive})
    sa.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", sa.SessionId})
    sa.EntityData.Leafs.Append("status-desc", types.YLeaf{"StatusDesc", sa.StatusDesc})
    sa.EntityData.Leafs.Append("local-spi", types.YLeaf{"LocalSpi", sa.LocalSpi})
    sa.EntityData.Leafs.Append("remote-spi", types.YLeaf{"RemoteSpi", sa.RemoteSpi})
    sa.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", sa.LocalId})
    sa.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", sa.RemoteId})
    sa.EntityData.Leafs.Append("sa-initiator", types.YLeaf{"SaInitiator", sa.SaInitiator})

    sa.EntityData.YListKeys = []string {}

    return &(sa.EntityData)
}

// IkEv2_Nodes_Node_Sa_All
// IKEv2 SA all data
type IkEv2_Nodes_Node_Sa_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SA list. The type is slice of IkEv2_Nodes_Node_Sa_All_Sa.
    Sa []*IkEv2_Nodes_Node_Sa_All_Sa
}

func (all *IkEv2_Nodes_Node_Sa_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "sa"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Children.Append("sa", types.YChild{"Sa", nil})
    for i := range all.Sa {
        types.SetYListKey(all.Sa[i], i)
        all.EntityData.Children.Append(types.GetSegmentPath(all.Sa[i]), types.YChild{"Sa", all.Sa[i]})
    }
    all.EntityData.Leafs = types.NewOrderedMap()

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// IkEv2_Nodes_Node_Sa_All_Sa
// SA list
type IkEv2_Nodes_Node_Sa_All_Sa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // Local Addr Port. The type is string.
    LocalAddrPort interface{}

    // Remote Addr Port. The type is string.
    RemoteAddrPort interface{}

    // State. The type is string.
    State interface{}

    // Encryption. The type is string.
    Encr interface{}

    // Keysize. The type is interface{} with range: 0..4294967295.
    Keysize interface{}

    // PRF. The type is string.
    Prf interface{}

    // Hash. The type is string.
    Hash interface{}

    // DH Group. The type is interface{} with range: 0..4294967295.
    DhGroup interface{}

    // Auth Sign. The type is string.
    AuthSign interface{}

    // Auth Verify. The type is string.
    AuthVerify interface{}

    // Life-Active Time. The type is string.
    LifeActive interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Status Description. The type is string.
    StatusDesc interface{}

    // Local SPI. The type is string.
    LocalSpi interface{}

    // Remote SPI. The type is string.
    RemoteSpi interface{}

    // Local ID. The type is string.
    LocalId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Sa Initiator. The type is bool.
    SaInitiator interface{}
}

func (sa *IkEv2_Nodes_Node_Sa_All_Sa) GetEntityData() *types.CommonEntityData {
    sa.EntityData.YFilter = sa.YFilter
    sa.EntityData.YangName = "sa"
    sa.EntityData.BundleName = "cisco_ios_xr"
    sa.EntityData.ParentYangName = "all"
    sa.EntityData.SegmentPath = "sa" + types.AddNoKeyToken(sa)
    sa.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/sa/all/" + sa.EntityData.SegmentPath
    sa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sa.EntityData.Children = types.NewOrderedMap()
    sa.EntityData.Leafs = types.NewOrderedMap()
    sa.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", sa.TunnelId})
    sa.EntityData.Leafs.Append("local-addr-port", types.YLeaf{"LocalAddrPort", sa.LocalAddrPort})
    sa.EntityData.Leafs.Append("remote-addr-port", types.YLeaf{"RemoteAddrPort", sa.RemoteAddrPort})
    sa.EntityData.Leafs.Append("state", types.YLeaf{"State", sa.State})
    sa.EntityData.Leafs.Append("encr", types.YLeaf{"Encr", sa.Encr})
    sa.EntityData.Leafs.Append("keysize", types.YLeaf{"Keysize", sa.Keysize})
    sa.EntityData.Leafs.Append("prf", types.YLeaf{"Prf", sa.Prf})
    sa.EntityData.Leafs.Append("hash", types.YLeaf{"Hash", sa.Hash})
    sa.EntityData.Leafs.Append("dh-group", types.YLeaf{"DhGroup", sa.DhGroup})
    sa.EntityData.Leafs.Append("auth-sign", types.YLeaf{"AuthSign", sa.AuthSign})
    sa.EntityData.Leafs.Append("auth-verify", types.YLeaf{"AuthVerify", sa.AuthVerify})
    sa.EntityData.Leafs.Append("life-active", types.YLeaf{"LifeActive", sa.LifeActive})
    sa.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", sa.SessionId})
    sa.EntityData.Leafs.Append("status-desc", types.YLeaf{"StatusDesc", sa.StatusDesc})
    sa.EntityData.Leafs.Append("local-spi", types.YLeaf{"LocalSpi", sa.LocalSpi})
    sa.EntityData.Leafs.Append("remote-spi", types.YLeaf{"RemoteSpi", sa.RemoteSpi})
    sa.EntityData.Leafs.Append("local-id", types.YLeaf{"LocalId", sa.LocalId})
    sa.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", sa.RemoteId})
    sa.EntityData.Leafs.Append("sa-initiator", types.YLeaf{"SaInitiator", sa.SaInitiator})

    sa.EntityData.YListKeys = []string {}

    return &(sa.EntityData)
}

// IkEv2_Nodes_Node_Policies
// IKEv2 policies on this node
type IkEv2_Nodes_Node_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 policy data. The type is slice of IkEv2_Nodes_Node_Policies_Policy.
    Policy []*IkEv2_Nodes_Node_Policies_Policy
}

func (policies *IkEv2_Nodes_Node_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "node"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + policies.EntityData.SegmentPath
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range policies.Policy {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.Policy[i]), types.YChild{"Policy", policies.Policy[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// IkEv2_Nodes_Node_Policies_Policy
// IKEv2 policy data
type IkEv2_Nodes_Node_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the IKEv2 policy. The type is string with
    // length: 1..32.
    Name interface{}

    // Name of the policy. The type is string.
    PolicyName interface{}

    // Match address of peer. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}

    // List of proposals. The type is slice of string.
    Proposal []interface{}
}

func (policy *IkEv2_Nodes_Node_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.Name, "name")
    policy.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/policies/" + policy.EntityData.SegmentPath
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("name", types.YLeaf{"Name", policy.Name})
    policy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policy.PolicyName})
    policy.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", policy.Addr})
    policy.EntityData.Leafs.Append("proposal", types.YLeaf{"Proposal", policy.Proposal})

    policy.EntityData.YListKeys = []string {"Name"}

    return &(policy.EntityData)
}

// IkEv2_Nodes_Node_Proposals
// IKEv2 proposals on this node
type IkEv2_Nodes_Node_Proposals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 proposal data. The type is slice of
    // IkEv2_Nodes_Node_Proposals_Proposal.
    Proposal []*IkEv2_Nodes_Node_Proposals_Proposal
}

func (proposals *IkEv2_Nodes_Node_Proposals) GetEntityData() *types.CommonEntityData {
    proposals.EntityData.YFilter = proposals.YFilter
    proposals.EntityData.YangName = "proposals"
    proposals.EntityData.BundleName = "cisco_ios_xr"
    proposals.EntityData.ParentYangName = "node"
    proposals.EntityData.SegmentPath = "proposals"
    proposals.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + proposals.EntityData.SegmentPath
    proposals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proposals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proposals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proposals.EntityData.Children = types.NewOrderedMap()
    proposals.EntityData.Children.Append("proposal", types.YChild{"Proposal", nil})
    for i := range proposals.Proposal {
        proposals.EntityData.Children.Append(types.GetSegmentPath(proposals.Proposal[i]), types.YChild{"Proposal", proposals.Proposal[i]})
    }
    proposals.EntityData.Leafs = types.NewOrderedMap()

    proposals.EntityData.YListKeys = []string {}

    return &(proposals.EntityData)
}

// IkEv2_Nodes_Node_Proposals_Proposal
// IKEv2 proposal data
type IkEv2_Nodes_Node_Proposals_Proposal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the IKEv2 proposal. The type is string
    // with length: 1..32.
    Name interface{}

    // Name of the proposal. The type is string.
    ProposalName interface{}

    // Encryption Algorithm. The type is slice of string.
    EncryptionAlg []interface{}

    // Integrity Algorithm. The type is slice of string.
    HashAlg []interface{}

    // PRF Algorithm. The type is slice of string.
    PrfAlg []interface{}

    // Group Algorithm. The type is slice of string.
    GroupAlg []interface{}
}

func (proposal *IkEv2_Nodes_Node_Proposals_Proposal) GetEntityData() *types.CommonEntityData {
    proposal.EntityData.YFilter = proposal.YFilter
    proposal.EntityData.YangName = "proposal"
    proposal.EntityData.BundleName = "cisco_ios_xr"
    proposal.EntityData.ParentYangName = "proposals"
    proposal.EntityData.SegmentPath = "proposal" + types.AddKeyToken(proposal.Name, "name")
    proposal.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/proposals/" + proposal.EntityData.SegmentPath
    proposal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proposal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proposal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proposal.EntityData.Children = types.NewOrderedMap()
    proposal.EntityData.Leafs = types.NewOrderedMap()
    proposal.EntityData.Leafs.Append("name", types.YLeaf{"Name", proposal.Name})
    proposal.EntityData.Leafs.Append("proposal-name", types.YLeaf{"ProposalName", proposal.ProposalName})
    proposal.EntityData.Leafs.Append("encryption-alg", types.YLeaf{"EncryptionAlg", proposal.EncryptionAlg})
    proposal.EntityData.Leafs.Append("hash-alg", types.YLeaf{"HashAlg", proposal.HashAlg})
    proposal.EntityData.Leafs.Append("prf-alg", types.YLeaf{"PrfAlg", proposal.PrfAlg})
    proposal.EntityData.Leafs.Append("group-alg", types.YLeaf{"GroupAlg", proposal.GroupAlg})

    proposal.EntityData.YListKeys = []string {"Name"}

    return &(proposal.EntityData)
}

// IkEv2_Nodes_Node_Profiles
// IKEv2 profiles on this node
type IkEv2_Nodes_Node_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 profile data. The type is slice of IkEv2_Nodes_Node_Profiles_Profile.
    Profile []*IkEv2_Nodes_Node_Profiles_Profile
}

func (profiles *IkEv2_Nodes_Node_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "node"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + profiles.EntityData.SegmentPath
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// IkEv2_Nodes_Node_Profiles_Profile
// IKEv2 profile data
type IkEv2_Nodes_Node_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the IKEv2 profile. The type is string with
    // length: 1..32.
    Name interface{}

    // Name of the Profile. The type is string.
    ProfileName interface{}

    // Name of the keyring. The type is string.
    KeyringName interface{}

    // Match Any Criteria. The type is bool.
    MatchAny interface{}

    // Lifetime. The type is interface{} with range: 0..4294967295.
    Lifetime interface{}

    // DPD Interval. The type is interface{} with range: 0..4294967295.
    DpdInterval interface{}

    // DPD Retry. The type is interface{} with range: 0..4294967295.
    DpdRetry interface{}

    // Match address of peer. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr []interface{}

    // Mask of peer. The type is slice of string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Mask []interface{}
}

func (profile *IkEv2_Nodes_Node_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.Name, "name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("name", types.YLeaf{"Name", profile.Name})
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("keyring-name", types.YLeaf{"KeyringName", profile.KeyringName})
    profile.EntityData.Leafs.Append("match-any", types.YLeaf{"MatchAny", profile.MatchAny})
    profile.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", profile.Lifetime})
    profile.EntityData.Leafs.Append("dpd-interval", types.YLeaf{"DpdInterval", profile.DpdInterval})
    profile.EntityData.Leafs.Append("dpd-retry", types.YLeaf{"DpdRetry", profile.DpdRetry})
    profile.EntityData.Leafs.Append("addr", types.YLeaf{"Addr", profile.Addr})
    profile.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", profile.Mask})

    profile.EntityData.YListKeys = []string {"Name"}

    return &(profile.EntityData)
}

// IkEv2_Nodes_Node_Keyrings
// IKEv2 keyrings on this node
type IkEv2_Nodes_Node_Keyrings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 keyring data. The type is slice of IkEv2_Nodes_Node_Keyrings_Keyring.
    Keyring []*IkEv2_Nodes_Node_Keyrings_Keyring
}

func (keyrings *IkEv2_Nodes_Node_Keyrings) GetEntityData() *types.CommonEntityData {
    keyrings.EntityData.YFilter = keyrings.YFilter
    keyrings.EntityData.YangName = "keyrings"
    keyrings.EntityData.BundleName = "cisco_ios_xr"
    keyrings.EntityData.ParentYangName = "node"
    keyrings.EntityData.SegmentPath = "keyrings"
    keyrings.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/" + keyrings.EntityData.SegmentPath
    keyrings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyrings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyrings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyrings.EntityData.Children = types.NewOrderedMap()
    keyrings.EntityData.Children.Append("keyring", types.YChild{"Keyring", nil})
    for i := range keyrings.Keyring {
        keyrings.EntityData.Children.Append(types.GetSegmentPath(keyrings.Keyring[i]), types.YChild{"Keyring", keyrings.Keyring[i]})
    }
    keyrings.EntityData.Leafs = types.NewOrderedMap()

    keyrings.EntityData.YListKeys = []string {}

    return &(keyrings.EntityData)
}

// IkEv2_Nodes_Node_Keyrings_Keyring
// IKEv2 keyring data
type IkEv2_Nodes_Node_Keyrings_Keyring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the IKEv2 keyring. The type is string with
    // length: 1..32.
    Name interface{}

    // Name of the Keyring. The type is string.
    KeyringName interface{}

    // List of peers. The type is slice of IkEv2_Nodes_Node_Keyrings_Keyring_Peer.
    Peer []*IkEv2_Nodes_Node_Keyrings_Keyring_Peer
}

func (keyring *IkEv2_Nodes_Node_Keyrings_Keyring) GetEntityData() *types.CommonEntityData {
    keyring.EntityData.YFilter = keyring.YFilter
    keyring.EntityData.YangName = "keyring"
    keyring.EntityData.BundleName = "cisco_ios_xr"
    keyring.EntityData.ParentYangName = "keyrings"
    keyring.EntityData.SegmentPath = "keyring" + types.AddKeyToken(keyring.Name, "name")
    keyring.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/keyrings/" + keyring.EntityData.SegmentPath
    keyring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyring.EntityData.Children = types.NewOrderedMap()
    keyring.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range keyring.Peer {
        types.SetYListKey(keyring.Peer[i], i)
        keyring.EntityData.Children.Append(types.GetSegmentPath(keyring.Peer[i]), types.YChild{"Peer", keyring.Peer[i]})
    }
    keyring.EntityData.Leafs = types.NewOrderedMap()
    keyring.EntityData.Leafs.Append("name", types.YLeaf{"Name", keyring.Name})
    keyring.EntityData.Leafs.Append("keyring-name", types.YLeaf{"KeyringName", keyring.KeyringName})

    keyring.EntityData.YListKeys = []string {"Name"}

    return &(keyring.EntityData)
}

// IkEv2_Nodes_Node_Keyrings_Keyring_Peer
// List of peers
type IkEv2_Nodes_Node_Keyrings_Keyring_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Name of the peer. The type is string.
    PeerName interface{}

    // IP Address of the peer. The type is string.
    IpAddress interface{}

    // Subnet mask of the peer. The type is string.
    Subnet interface{}

    // Local PSK. The type is string.
    LocalPsk interface{}

    // Remote PSK. The type is string.
    RemotePsk interface{}
}

func (peer *IkEv2_Nodes_Node_Keyrings_Keyring_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "keyring"
    peer.EntityData.SegmentPath = "peer" + types.AddNoKeyToken(peer)
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-oper:ik-ev2/nodes/node/keyrings/keyring/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-name", types.YLeaf{"PeerName", peer.PeerName})
    peer.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", peer.IpAddress})
    peer.EntityData.Leafs.Append("subnet", types.YLeaf{"Subnet", peer.Subnet})
    peer.EntityData.Leafs.Append("local-psk", types.YLeaf{"LocalPsk", peer.LocalPsk})
    peer.EntityData.Leafs.Append("remote-psk", types.YLeaf{"RemotePsk", peer.RemotePsk})

    peer.EntityData.YListKeys = []string {}

    return &(peer.EntityData)
}

