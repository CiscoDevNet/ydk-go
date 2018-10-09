// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-netconf package configuration.
// 
// This module contains definitions
// for the following management objects:
//   netconf-yang: NETCONF YANG configuration commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // Models to be disabled.
    Models NetconfYang_Agent_Models

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
    agent.EntityData.Children.Append("models", types.YChild{"Models", &agent.Models})
    agent.EntityData.Children.Append("ssh", types.YChild{"Ssh", &agent.Ssh})
    agent.EntityData.Children.Append("session", types.YChild{"Session", &agent.Session})
    agent.EntityData.Leafs = types.NewOrderedMap()
    agent.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", agent.RateLimit})

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// NetconfYang_Agent_Models
// Models to be disabled
type NetconfYang_Agent_Models struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of models: openconfig.
    Openconfig NetconfYang_Agent_Models_Openconfig
}

func (models *NetconfYang_Agent_Models) GetEntityData() *types.CommonEntityData {
    models.EntityData.YFilter = models.YFilter
    models.EntityData.YangName = "models"
    models.EntityData.BundleName = "cisco_ios_xr"
    models.EntityData.ParentYangName = "agent"
    models.EntityData.SegmentPath = "models"
    models.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    models.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    models.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    models.EntityData.Children = types.NewOrderedMap()
    models.EntityData.Children.Append("openconfig", types.YChild{"Openconfig", &models.Openconfig})
    models.EntityData.Leafs = types.NewOrderedMap()

    models.EntityData.YListKeys = []string {}

    return &(models.EntityData)
}

// NetconfYang_Agent_Models_Openconfig
// Type of models: openconfig
type NetconfYang_Agent_Models_Openconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable the specified model type. The type is interface{}.
    Disabled interface{}
}

func (openconfig *NetconfYang_Agent_Models_Openconfig) GetEntityData() *types.CommonEntityData {
    openconfig.EntityData.YFilter = openconfig.YFilter
    openconfig.EntityData.YangName = "openconfig"
    openconfig.EntityData.BundleName = "cisco_ios_xr"
    openconfig.EntityData.ParentYangName = "models"
    openconfig.EntityData.SegmentPath = "openconfig"
    openconfig.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    openconfig.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    openconfig.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    openconfig.EntityData.Children = types.NewOrderedMap()
    openconfig.EntityData.Leafs = types.NewOrderedMap()
    openconfig.EntityData.Leafs.Append("disabled", types.YLeaf{"Disabled", openconfig.Disabled})

    openconfig.EntityData.YListKeys = []string {}

    return &(openconfig.EntityData)
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

