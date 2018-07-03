// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-netconf package configuration.
// 
// This module contains definitions
// for the following management objects:
//   netconf-yang: NETCONF YANG configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_netconf_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_netconf_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-netconf-cfg netconf-yang}", reflect.TypeOf(NetconfYang{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-netconf-cfg:netconf-yang", reflect.TypeOf(NetconfYang{}))
}

// NetconfYang
// NETCONF YANG configuration commands
type NetconfYang struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NETCONF YANG agent configuration commands.
    Agent NetconfYang_Agent
}

func (netconfYang *NetconfYang) GetEntityData() *types.CommonEntityData {
    netconfYang.EntityData.YFilter = netconfYang.YFilter
    netconfYang.EntityData.YangName = "netconf-yang"
    netconfYang.EntityData.BundleName = "cisco_ios_xr"
    netconfYang.EntityData.ParentYangName = "Cisco-IOS-XR-man-netconf-cfg"
    netconfYang.EntityData.SegmentPath = "Cisco-IOS-XR-man-netconf-cfg:netconf-yang"
    netconfYang.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconfYang.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconfYang.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconfYang.EntityData.Children = types.NewOrderedMap()
    netconfYang.EntityData.Children.Append("agent", types.YChild{"Agent", &netconfYang.Agent})
    netconfYang.EntityData.Leafs = types.NewOrderedMap()

    netconfYang.EntityData.YListKeys = []string {}

    return &(netconfYang.EntityData)
}

// NetconfYang_Agent
// NETCONF YANG agent configuration commands
type NetconfYang_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of bytes to process per sec. The type is interface{} with range:
    // 4096..4294967295. Units are byte.
    RateLimit interface{}

    // NETCONF YANG agent over SSH connection.
    Ssh NetconfYang_Agent_Ssh

    // Session settings.
    Session NetconfYang_Agent_Session
}

func (agent *NetconfYang_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "netconf-yang"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("ssh", types.YChild{"Ssh", &agent.Ssh})
    agent.EntityData.Children.Append("session", types.YChild{"Session", &agent.Session})
    agent.EntityData.Leafs = types.NewOrderedMap()
    agent.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", agent.RateLimit})

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// NetconfYang_Agent_Ssh
// NETCONF YANG agent over SSH connection
type NetconfYang_Agent_Ssh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable NETCONF YANG agent over SSH connection. The type is interface{}.
    Enable interface{}
}

func (ssh *NetconfYang_Agent_Ssh) GetEntityData() *types.CommonEntityData {
    ssh.EntityData.YFilter = ssh.YFilter
    ssh.EntityData.YangName = "ssh"
    ssh.EntityData.BundleName = "cisco_ios_xr"
    ssh.EntityData.ParentYangName = "agent"
    ssh.EntityData.SegmentPath = "ssh"
    ssh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssh.EntityData.Children = types.NewOrderedMap()
    ssh.EntityData.Leafs = types.NewOrderedMap()
    ssh.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ssh.Enable})

    ssh.EntityData.YListKeys = []string {}

    return &(ssh.EntityData)
}

// NetconfYang_Agent_Session
// Session settings
type NetconfYang_Agent_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of allowable concurrent netconf-yang sessions. The type is
    // interface{} with range: 1..50. The default value is 50.
    Limit interface{}

    // Absolute timeout in minutes. The type is interface{} with range: 1..1440.
    // Units are minute.
    AbsoluteTimeout interface{}

    // Non-active session lifetime. The type is interface{} with range: 1..1440.
    // Units are minute.
    IdleTimeout interface{}
}

func (session *NetconfYang_Agent_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "agent"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("limit", types.YLeaf{"Limit", session.Limit})
    session.EntityData.Leafs.Append("absolute-timeout", types.YLeaf{"AbsoluteTimeout", session.AbsoluteTimeout})
    session.EntityData.Leafs.Append("idle-timeout", types.YLeaf{"IdleTimeout", session.IdleTimeout})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

