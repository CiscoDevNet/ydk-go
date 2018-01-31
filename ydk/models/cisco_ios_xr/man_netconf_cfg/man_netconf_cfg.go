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
    parent types.Entity
    YFilter yfilter.YFilter

    // NETCONF YANG agent configuration commands.
    Agent NetconfYang_Agent
}

func (netconfYang *NetconfYang) GetFilter() yfilter.YFilter { return netconfYang.YFilter }

func (netconfYang *NetconfYang) SetFilter(yf yfilter.YFilter) { netconfYang.YFilter = yf }

func (netconfYang *NetconfYang) GetGoName(yname string) string {
    if yname == "agent" { return "Agent" }
    return ""
}

func (netconfYang *NetconfYang) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-netconf-cfg:netconf-yang"
}

func (netconfYang *NetconfYang) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "agent" {
        return &netconfYang.Agent
    }
    return nil
}

func (netconfYang *NetconfYang) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["agent"] = &netconfYang.Agent
    return children
}

func (netconfYang *NetconfYang) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconfYang *NetconfYang) GetBundleName() string { return "cisco_ios_xr" }

func (netconfYang *NetconfYang) GetYangName() string { return "netconf-yang" }

func (netconfYang *NetconfYang) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconfYang *NetconfYang) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconfYang *NetconfYang) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconfYang *NetconfYang) SetParent(parent types.Entity) { netconfYang.parent = parent }

func (netconfYang *NetconfYang) GetParent() types.Entity { return netconfYang.parent }

func (netconfYang *NetconfYang) GetParentYangName() string { return "Cisco-IOS-XR-man-netconf-cfg" }

// NetconfYang_Agent
// NETCONF YANG agent configuration commands
type NetconfYang_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of bytes to process per sec. The type is interface{} with range:
    // 4096..4294967295. Units are byte.
    RateLimit interface{}

    // NETCONF YANG agent over SSH connection.
    Ssh NetconfYang_Agent_Ssh

    // Session settings.
    Session NetconfYang_Agent_Session
}

func (agent *NetconfYang_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *NetconfYang_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *NetconfYang_Agent) GetGoName(yname string) string {
    if yname == "rate-limit" { return "RateLimit" }
    if yname == "ssh" { return "Ssh" }
    if yname == "session" { return "Session" }
    return ""
}

func (agent *NetconfYang_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *NetconfYang_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ssh" {
        return &agent.Ssh
    }
    if childYangName == "session" {
        return &agent.Session
    }
    return nil
}

func (agent *NetconfYang_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ssh"] = &agent.Ssh
    children["session"] = &agent.Session
    return children
}

func (agent *NetconfYang_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rate-limit"] = agent.RateLimit
    return leafs
}

func (agent *NetconfYang_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *NetconfYang_Agent) GetYangName() string { return "agent" }

func (agent *NetconfYang_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *NetconfYang_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *NetconfYang_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *NetconfYang_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *NetconfYang_Agent) GetParent() types.Entity { return agent.parent }

func (agent *NetconfYang_Agent) GetParentYangName() string { return "netconf-yang" }

// NetconfYang_Agent_Ssh
// NETCONF YANG agent over SSH connection
type NetconfYang_Agent_Ssh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable NETCONF YANG agent over SSH connection. The type is interface{}.
    Enable interface{}
}

func (ssh *NetconfYang_Agent_Ssh) GetFilter() yfilter.YFilter { return ssh.YFilter }

func (ssh *NetconfYang_Agent_Ssh) SetFilter(yf yfilter.YFilter) { ssh.YFilter = yf }

func (ssh *NetconfYang_Agent_Ssh) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (ssh *NetconfYang_Agent_Ssh) GetSegmentPath() string {
    return "ssh"
}

func (ssh *NetconfYang_Agent_Ssh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ssh *NetconfYang_Agent_Ssh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ssh *NetconfYang_Agent_Ssh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ssh.Enable
    return leafs
}

func (ssh *NetconfYang_Agent_Ssh) GetBundleName() string { return "cisco_ios_xr" }

func (ssh *NetconfYang_Agent_Ssh) GetYangName() string { return "ssh" }

func (ssh *NetconfYang_Agent_Ssh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssh *NetconfYang_Agent_Ssh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssh *NetconfYang_Agent_Ssh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssh *NetconfYang_Agent_Ssh) SetParent(parent types.Entity) { ssh.parent = parent }

func (ssh *NetconfYang_Agent_Ssh) GetParent() types.Entity { return ssh.parent }

func (ssh *NetconfYang_Agent_Ssh) GetParentYangName() string { return "agent" }

// NetconfYang_Agent_Session
// Session settings
type NetconfYang_Agent_Session struct {
    parent types.Entity
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

func (session *NetconfYang_Agent_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *NetconfYang_Agent_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *NetconfYang_Agent_Session) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "idle-timeout" { return "IdleTimeout" }
    return ""
}

func (session *NetconfYang_Agent_Session) GetSegmentPath() string {
    return "session"
}

func (session *NetconfYang_Agent_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *NetconfYang_Agent_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *NetconfYang_Agent_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = session.Limit
    leafs["absolute-timeout"] = session.AbsoluteTimeout
    leafs["idle-timeout"] = session.IdleTimeout
    return leafs
}

func (session *NetconfYang_Agent_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *NetconfYang_Agent_Session) GetYangName() string { return "session" }

func (session *NetconfYang_Agent_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *NetconfYang_Agent_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *NetconfYang_Agent_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *NetconfYang_Agent_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *NetconfYang_Agent_Session) GetParent() types.Entity { return session.parent }

func (session *NetconfYang_Agent_Session) GetParentYangName() string { return "agent" }

