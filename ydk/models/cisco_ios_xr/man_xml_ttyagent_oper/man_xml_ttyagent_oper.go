// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-xml-ttyagent package operational data.
// 
// This module contains definitions
// for the following management objects:
//   netconf: NETCONF operational information
//   xr-xml: xr xml
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package man_xml_ttyagent_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_xml_ttyagent_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-oper netconf}", reflect.TypeOf(Netconf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-oper:netconf", reflect.TypeOf(Netconf{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-oper xr-xml}", reflect.TypeOf(XrXml{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-oper:xr-xml", reflect.TypeOf(XrXml{}))
}

// XrXmlSessionAlarmRegister represents AlarmNotify
type XrXmlSessionAlarmRegister string

const (
    // Registered
    XrXmlSessionAlarmRegister_registered XrXmlSessionAlarmRegister = "registered"

    // NotRegistered
    XrXmlSessionAlarmRegister_not_registered XrXmlSessionAlarmRegister = "not-registered"
)

// XrXmlSessionState represents SessionState
type XrXmlSessionState string

const (
    // Idle
    XrXmlSessionState_idle XrXmlSessionState = "idle"

    // Busy
    XrXmlSessionState_busy XrXmlSessionState = "busy"
)

// Netconf
// NETCONF operational information
type Netconf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NETCONF agent operational information.
    Agent Netconf_Agent
}

func (netconf *Netconf) GetFilter() yfilter.YFilter { return netconf.YFilter }

func (netconf *Netconf) SetFilter(yf yfilter.YFilter) { netconf.YFilter = yf }

func (netconf *Netconf) GetGoName(yname string) string {
    if yname == "agent" { return "Agent" }
    return ""
}

func (netconf *Netconf) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-xml-ttyagent-oper:netconf"
}

func (netconf *Netconf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "agent" {
        return &netconf.Agent
    }
    return nil
}

func (netconf *Netconf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["agent"] = &netconf.Agent
    return children
}

func (netconf *Netconf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (netconf *Netconf) GetBundleName() string { return "cisco_ios_xr" }

func (netconf *Netconf) GetYangName() string { return "netconf" }

func (netconf *Netconf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netconf *Netconf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netconf *Netconf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netconf *Netconf) SetParent(parent types.Entity) { netconf.parent = parent }

func (netconf *Netconf) GetParent() types.Entity { return netconf.parent }

func (netconf *Netconf) GetParentYangName() string { return "Cisco-IOS-XR-man-xml-ttyagent-oper" }

// Netconf_Agent
// NETCONF agent operational information
type Netconf_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NETCONF agent over TTY.
    Tty Netconf_Agent_Tty
}

func (agent *Netconf_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *Netconf_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *Netconf_Agent) GetGoName(yname string) string {
    if yname == "tty" { return "Tty" }
    return ""
}

func (agent *Netconf_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *Netconf_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tty" {
        return &agent.Tty
    }
    return nil
}

func (agent *Netconf_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tty"] = &agent.Tty
    return children
}

func (agent *Netconf_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (agent *Netconf_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *Netconf_Agent) GetYangName() string { return "agent" }

func (agent *Netconf_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *Netconf_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *Netconf_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *Netconf_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *Netconf_Agent) GetParent() types.Entity { return agent.parent }

func (agent *Netconf_Agent) GetParentYangName() string { return "netconf" }

// Netconf_Agent_Tty
// NETCONF agent over TTY
type Netconf_Agent_Tty struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session information.
    Sessions Netconf_Agent_Tty_Sessions
}

func (tty *Netconf_Agent_Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *Netconf_Agent_Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *Netconf_Agent_Tty) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (tty *Netconf_Agent_Tty) GetSegmentPath() string {
    return "tty"
}

func (tty *Netconf_Agent_Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &tty.Sessions
    }
    return nil
}

func (tty *Netconf_Agent_Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &tty.Sessions
    return children
}

func (tty *Netconf_Agent_Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tty *Netconf_Agent_Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *Netconf_Agent_Tty) GetYangName() string { return "tty" }

func (tty *Netconf_Agent_Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *Netconf_Agent_Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *Netconf_Agent_Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *Netconf_Agent_Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *Netconf_Agent_Tty) GetParent() types.Entity { return tty.parent }

func (tty *Netconf_Agent_Tty) GetParentYangName() string { return "agent" }

// Netconf_Agent_Tty_Sessions
// Session information
type Netconf_Agent_Tty_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session information. The type is slice of
    // Netconf_Agent_Tty_Sessions_Session.
    Session []Netconf_Agent_Tty_Sessions_Session
}

func (sessions *Netconf_Agent_Tty_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Netconf_Agent_Tty_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Netconf_Agent_Tty_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *Netconf_Agent_Tty_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Netconf_Agent_Tty_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Netconf_Agent_Tty_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *Netconf_Agent_Tty_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *Netconf_Agent_Tty_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Netconf_Agent_Tty_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Netconf_Agent_Tty_Sessions) GetYangName() string { return "sessions" }

func (sessions *Netconf_Agent_Tty_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Netconf_Agent_Tty_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Netconf_Agent_Tty_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Netconf_Agent_Tty_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Netconf_Agent_Tty_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Netconf_Agent_Tty_Sessions) GetParentYangName() string { return "tty" }

// Netconf_Agent_Tty_Sessions_Session
// Session information
type Netconf_Agent_Tty_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session ID. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // Username. The type is string.
    Username interface{}

    // state of the session idle/busy. The type is XrXmlSessionState.
    State interface{}

    // ip address of the client. The type is string.
    ClientAddress interface{}

    // client's port. The type is interface{} with range: 0..4294967295.
    ClientPort interface{}

    // Config session ID. The type is string.
    ConfigSessionId interface{}

    // Admin config session ID. The type is string.
    AdminConfigSessionId interface{}

    // is the session registered for alarm notifications. The type is
    // XrXmlSessionAlarmRegister.
    AlarmNotification interface{}

    // VRF name . The type is string.
    VrfName interface{}

    // session start time in seconds since the Unix Epoch. The type is interface{}
    // with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Elapsed time(seconds) since a session is created. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ElapsedTime interface{}

    // Time(seconds) since last session state change happened . The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastStateChange interface{}
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Netconf_Agent_Tty_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Netconf_Agent_Tty_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "username" { return "Username" }
    if yname == "state" { return "State" }
    if yname == "client-address" { return "ClientAddress" }
    if yname == "client-port" { return "ClientPort" }
    if yname == "config-session-id" { return "ConfigSessionId" }
    if yname == "admin-config-session-id" { return "AdminConfigSessionId" }
    if yname == "alarm-notification" { return "AlarmNotification" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "start-time" { return "StartTime" }
    if yname == "elapsed-time" { return "ElapsedTime" }
    if yname == "last-state-change" { return "LastStateChange" }
    return ""
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["username"] = session.Username
    leafs["state"] = session.State
    leafs["client-address"] = session.ClientAddress
    leafs["client-port"] = session.ClientPort
    leafs["config-session-id"] = session.ConfigSessionId
    leafs["admin-config-session-id"] = session.AdminConfigSessionId
    leafs["alarm-notification"] = session.AlarmNotification
    leafs["vrf-name"] = session.VrfName
    leafs["start-time"] = session.StartTime
    leafs["elapsed-time"] = session.ElapsedTime
    leafs["last-state-change"] = session.LastStateChange
    return leafs
}

func (session *Netconf_Agent_Tty_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Netconf_Agent_Tty_Sessions_Session) GetYangName() string { return "session" }

func (session *Netconf_Agent_Tty_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Netconf_Agent_Tty_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Netconf_Agent_Tty_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Netconf_Agent_Tty_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Netconf_Agent_Tty_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *Netconf_Agent_Tty_Sessions_Session) GetParentYangName() string { return "sessions" }

// XrXml
// xr xml
type XrXml struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // XML agents.
    Agent XrXml_Agent
}

func (xrXml *XrXml) GetFilter() yfilter.YFilter { return xrXml.YFilter }

func (xrXml *XrXml) SetFilter(yf yfilter.YFilter) { xrXml.YFilter = yf }

func (xrXml *XrXml) GetGoName(yname string) string {
    if yname == "agent" { return "Agent" }
    return ""
}

func (xrXml *XrXml) GetSegmentPath() string {
    return "Cisco-IOS-XR-man-xml-ttyagent-oper:xr-xml"
}

func (xrXml *XrXml) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "agent" {
        return &xrXml.Agent
    }
    return nil
}

func (xrXml *XrXml) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["agent"] = &xrXml.Agent
    return children
}

func (xrXml *XrXml) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (xrXml *XrXml) GetBundleName() string { return "cisco_ios_xr" }

func (xrXml *XrXml) GetYangName() string { return "xr-xml" }

func (xrXml *XrXml) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (xrXml *XrXml) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (xrXml *XrXml) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (xrXml *XrXml) SetParent(parent types.Entity) { xrXml.parent = parent }

func (xrXml *XrXml) GetParent() types.Entity { return xrXml.parent }

func (xrXml *XrXml) GetParentYangName() string { return "Cisco-IOS-XR-man-xml-ttyagent-oper" }

// XrXml_Agent
// XML agents
type XrXml_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TTY sessions information.
    Tty XrXml_Agent_Tty

    // Default sessions information.
    Default XrXml_Agent_Default

    // SSL sessions information.
    Ssl XrXml_Agent_Ssl
}

func (agent *XrXml_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *XrXml_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *XrXml_Agent) GetGoName(yname string) string {
    if yname == "tty" { return "Tty" }
    if yname == "default" { return "Default" }
    if yname == "ssl" { return "Ssl" }
    return ""
}

func (agent *XrXml_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *XrXml_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tty" {
        return &agent.Tty
    }
    if childYangName == "default" {
        return &agent.Default
    }
    if childYangName == "ssl" {
        return &agent.Ssl
    }
    return nil
}

func (agent *XrXml_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tty"] = &agent.Tty
    children["default"] = &agent.Default
    children["ssl"] = &agent.Ssl
    return children
}

func (agent *XrXml_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (agent *XrXml_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *XrXml_Agent) GetYangName() string { return "agent" }

func (agent *XrXml_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *XrXml_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *XrXml_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *XrXml_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *XrXml_Agent) GetParent() types.Entity { return agent.parent }

func (agent *XrXml_Agent) GetParentYangName() string { return "xr-xml" }

// XrXml_Agent_Tty
// TTY sessions information
type XrXml_Agent_Tty struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Tty_Sessions
}

func (tty *XrXml_Agent_Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *XrXml_Agent_Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *XrXml_Agent_Tty) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (tty *XrXml_Agent_Tty) GetSegmentPath() string {
    return "tty"
}

func (tty *XrXml_Agent_Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &tty.Sessions
    }
    return nil
}

func (tty *XrXml_Agent_Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &tty.Sessions
    return children
}

func (tty *XrXml_Agent_Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tty *XrXml_Agent_Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *XrXml_Agent_Tty) GetYangName() string { return "tty" }

func (tty *XrXml_Agent_Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *XrXml_Agent_Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *XrXml_Agent_Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *XrXml_Agent_Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *XrXml_Agent_Tty) GetParent() types.Entity { return tty.parent }

func (tty *XrXml_Agent_Tty) GetParentYangName() string { return "agent" }

// XrXml_Agent_Tty_Sessions
// sessions information
type XrXml_Agent_Tty_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Tty_Sessions_Session.
    Session []XrXml_Agent_Tty_Sessions_Session
}

func (sessions *XrXml_Agent_Tty_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *XrXml_Agent_Tty_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *XrXml_Agent_Tty_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *XrXml_Agent_Tty_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *XrXml_Agent_Tty_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := XrXml_Agent_Tty_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *XrXml_Agent_Tty_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *XrXml_Agent_Tty_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *XrXml_Agent_Tty_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *XrXml_Agent_Tty_Sessions) GetYangName() string { return "sessions" }

func (sessions *XrXml_Agent_Tty_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *XrXml_Agent_Tty_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *XrXml_Agent_Tty_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *XrXml_Agent_Tty_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *XrXml_Agent_Tty_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *XrXml_Agent_Tty_Sessions) GetParentYangName() string { return "tty" }

// XrXml_Agent_Tty_Sessions_Session
// xml sessions information
type XrXml_Agent_Tty_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // Username. The type is string.
    Username interface{}

    // state of the session idle/busy. The type is XrXmlSessionState.
    State interface{}

    // ip address of the client. The type is string.
    ClientAddress interface{}

    // client's port. The type is interface{} with range: 0..4294967295.
    ClientPort interface{}

    // Config session ID. The type is string.
    ConfigSessionId interface{}

    // Admin config session ID. The type is string.
    AdminConfigSessionId interface{}

    // is the session registered for alarm notifications. The type is
    // XrXmlSessionAlarmRegister.
    AlarmNotification interface{}

    // VRF name . The type is string.
    VrfName interface{}

    // session start time in seconds since the Unix Epoch. The type is interface{}
    // with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Elapsed time(seconds) since a session is created. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ElapsedTime interface{}

    // Time(seconds) since last session state change happened . The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastStateChange interface{}
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Tty_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Tty_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "username" { return "Username" }
    if yname == "state" { return "State" }
    if yname == "client-address" { return "ClientAddress" }
    if yname == "client-port" { return "ClientPort" }
    if yname == "config-session-id" { return "ConfigSessionId" }
    if yname == "admin-config-session-id" { return "AdminConfigSessionId" }
    if yname == "alarm-notification" { return "AlarmNotification" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "start-time" { return "StartTime" }
    if yname == "elapsed-time" { return "ElapsedTime" }
    if yname == "last-state-change" { return "LastStateChange" }
    return ""
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["username"] = session.Username
    leafs["state"] = session.State
    leafs["client-address"] = session.ClientAddress
    leafs["client-port"] = session.ClientPort
    leafs["config-session-id"] = session.ConfigSessionId
    leafs["admin-config-session-id"] = session.AdminConfigSessionId
    leafs["alarm-notification"] = session.AlarmNotification
    leafs["vrf-name"] = session.VrfName
    leafs["start-time"] = session.StartTime
    leafs["elapsed-time"] = session.ElapsedTime
    leafs["last-state-change"] = session.LastStateChange
    return leafs
}

func (session *XrXml_Agent_Tty_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Tty_Sessions_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Tty_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Tty_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Tty_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Tty_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Tty_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Tty_Sessions_Session) GetParentYangName() string { return "sessions" }

// XrXml_Agent_Default
// Default sessions information
type XrXml_Agent_Default struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Default_Sessions
}

func (self *XrXml_Agent_Default) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *XrXml_Agent_Default) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *XrXml_Agent_Default) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (self *XrXml_Agent_Default) GetSegmentPath() string {
    return "default"
}

func (self *XrXml_Agent_Default) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &self.Sessions
    }
    return nil
}

func (self *XrXml_Agent_Default) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &self.Sessions
    return children
}

func (self *XrXml_Agent_Default) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *XrXml_Agent_Default) GetBundleName() string { return "cisco_ios_xr" }

func (self *XrXml_Agent_Default) GetYangName() string { return "default" }

func (self *XrXml_Agent_Default) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *XrXml_Agent_Default) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *XrXml_Agent_Default) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *XrXml_Agent_Default) SetParent(parent types.Entity) { self.parent = parent }

func (self *XrXml_Agent_Default) GetParent() types.Entity { return self.parent }

func (self *XrXml_Agent_Default) GetParentYangName() string { return "agent" }

// XrXml_Agent_Default_Sessions
// sessions information
type XrXml_Agent_Default_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Default_Sessions_Session.
    Session []XrXml_Agent_Default_Sessions_Session
}

func (sessions *XrXml_Agent_Default_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *XrXml_Agent_Default_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *XrXml_Agent_Default_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *XrXml_Agent_Default_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *XrXml_Agent_Default_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := XrXml_Agent_Default_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *XrXml_Agent_Default_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *XrXml_Agent_Default_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *XrXml_Agent_Default_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *XrXml_Agent_Default_Sessions) GetYangName() string { return "sessions" }

func (sessions *XrXml_Agent_Default_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *XrXml_Agent_Default_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *XrXml_Agent_Default_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *XrXml_Agent_Default_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *XrXml_Agent_Default_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *XrXml_Agent_Default_Sessions) GetParentYangName() string { return "default" }

// XrXml_Agent_Default_Sessions_Session
// xml sessions information
type XrXml_Agent_Default_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // Username. The type is string.
    Username interface{}

    // state of the session idle/busy. The type is XrXmlSessionState.
    State interface{}

    // ip address of the client. The type is string.
    ClientAddress interface{}

    // client's port. The type is interface{} with range: 0..4294967295.
    ClientPort interface{}

    // Config session ID. The type is string.
    ConfigSessionId interface{}

    // Admin config session ID. The type is string.
    AdminConfigSessionId interface{}

    // is the session registered for alarm notifications. The type is
    // XrXmlSessionAlarmRegister.
    AlarmNotification interface{}

    // VRF name . The type is string.
    VrfName interface{}

    // session start time in seconds since the Unix Epoch. The type is interface{}
    // with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Elapsed time(seconds) since a session is created. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ElapsedTime interface{}

    // Time(seconds) since last session state change happened . The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastStateChange interface{}
}

func (session *XrXml_Agent_Default_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Default_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Default_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "username" { return "Username" }
    if yname == "state" { return "State" }
    if yname == "client-address" { return "ClientAddress" }
    if yname == "client-port" { return "ClientPort" }
    if yname == "config-session-id" { return "ConfigSessionId" }
    if yname == "admin-config-session-id" { return "AdminConfigSessionId" }
    if yname == "alarm-notification" { return "AlarmNotification" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "start-time" { return "StartTime" }
    if yname == "elapsed-time" { return "ElapsedTime" }
    if yname == "last-state-change" { return "LastStateChange" }
    return ""
}

func (session *XrXml_Agent_Default_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *XrXml_Agent_Default_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Default_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Default_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["username"] = session.Username
    leafs["state"] = session.State
    leafs["client-address"] = session.ClientAddress
    leafs["client-port"] = session.ClientPort
    leafs["config-session-id"] = session.ConfigSessionId
    leafs["admin-config-session-id"] = session.AdminConfigSessionId
    leafs["alarm-notification"] = session.AlarmNotification
    leafs["vrf-name"] = session.VrfName
    leafs["start-time"] = session.StartTime
    leafs["elapsed-time"] = session.ElapsedTime
    leafs["last-state-change"] = session.LastStateChange
    return leafs
}

func (session *XrXml_Agent_Default_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Default_Sessions_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Default_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Default_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Default_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Default_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Default_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Default_Sessions_Session) GetParentYangName() string { return "sessions" }

// XrXml_Agent_Ssl
// SSL sessions information
type XrXml_Agent_Ssl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Ssl_Sessions
}

func (ssl *XrXml_Agent_Ssl) GetFilter() yfilter.YFilter { return ssl.YFilter }

func (ssl *XrXml_Agent_Ssl) SetFilter(yf yfilter.YFilter) { ssl.YFilter = yf }

func (ssl *XrXml_Agent_Ssl) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (ssl *XrXml_Agent_Ssl) GetSegmentPath() string {
    return "ssl"
}

func (ssl *XrXml_Agent_Ssl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &ssl.Sessions
    }
    return nil
}

func (ssl *XrXml_Agent_Ssl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &ssl.Sessions
    return children
}

func (ssl *XrXml_Agent_Ssl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssl *XrXml_Agent_Ssl) GetBundleName() string { return "cisco_ios_xr" }

func (ssl *XrXml_Agent_Ssl) GetYangName() string { return "ssl" }

func (ssl *XrXml_Agent_Ssl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssl *XrXml_Agent_Ssl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssl *XrXml_Agent_Ssl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssl *XrXml_Agent_Ssl) SetParent(parent types.Entity) { ssl.parent = parent }

func (ssl *XrXml_Agent_Ssl) GetParent() types.Entity { return ssl.parent }

func (ssl *XrXml_Agent_Ssl) GetParentYangName() string { return "agent" }

// XrXml_Agent_Ssl_Sessions
// sessions information
type XrXml_Agent_Ssl_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Ssl_Sessions_Session.
    Session []XrXml_Agent_Ssl_Sessions_Session
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *XrXml_Agent_Ssl_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *XrXml_Agent_Ssl_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := XrXml_Agent_Ssl_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *XrXml_Agent_Ssl_Sessions) GetYangName() string { return "sessions" }

func (sessions *XrXml_Agent_Ssl_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *XrXml_Agent_Ssl_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *XrXml_Agent_Ssl_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *XrXml_Agent_Ssl_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *XrXml_Agent_Ssl_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *XrXml_Agent_Ssl_Sessions) GetParentYangName() string { return "ssl" }

// XrXml_Agent_Ssl_Sessions_Session
// xml sessions information
type XrXml_Agent_Ssl_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // Username. The type is string.
    Username interface{}

    // state of the session idle/busy. The type is XrXmlSessionState.
    State interface{}

    // ip address of the client. The type is string.
    ClientAddress interface{}

    // client's port. The type is interface{} with range: 0..4294967295.
    ClientPort interface{}

    // Config session ID. The type is string.
    ConfigSessionId interface{}

    // Admin config session ID. The type is string.
    AdminConfigSessionId interface{}

    // is the session registered for alarm notifications. The type is
    // XrXmlSessionAlarmRegister.
    AlarmNotification interface{}

    // VRF name . The type is string.
    VrfName interface{}

    // session start time in seconds since the Unix Epoch. The type is interface{}
    // with range: 0..4294967295. Units are second.
    StartTime interface{}

    // Elapsed time(seconds) since a session is created. The type is interface{}
    // with range: 0..4294967295. Units are second.
    ElapsedTime interface{}

    // Time(seconds) since last session state change happened . The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastStateChange interface{}
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *XrXml_Agent_Ssl_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "username" { return "Username" }
    if yname == "state" { return "State" }
    if yname == "client-address" { return "ClientAddress" }
    if yname == "client-port" { return "ClientPort" }
    if yname == "config-session-id" { return "ConfigSessionId" }
    if yname == "admin-config-session-id" { return "AdminConfigSessionId" }
    if yname == "alarm-notification" { return "AlarmNotification" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "start-time" { return "StartTime" }
    if yname == "elapsed-time" { return "ElapsedTime" }
    if yname == "last-state-change" { return "LastStateChange" }
    return ""
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["username"] = session.Username
    leafs["state"] = session.State
    leafs["client-address"] = session.ClientAddress
    leafs["client-port"] = session.ClientPort
    leafs["config-session-id"] = session.ConfigSessionId
    leafs["admin-config-session-id"] = session.AdminConfigSessionId
    leafs["alarm-notification"] = session.AlarmNotification
    leafs["vrf-name"] = session.VrfName
    leafs["start-time"] = session.StartTime
    leafs["elapsed-time"] = session.ElapsedTime
    leafs["last-state-change"] = session.LastStateChange
    return leafs
}

func (session *XrXml_Agent_Ssl_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetYangName() string { return "session" }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *XrXml_Agent_Ssl_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *XrXml_Agent_Ssl_Sessions_Session) GetParentYangName() string { return "sessions" }

