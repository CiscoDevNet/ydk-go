// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-xml-ttyagent package operational data.
// 
// This module contains definitions
// for the following management objects:
//   netconf: NETCONF operational information
//   xr-xml: xr xml
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NETCONF agent operational information.
    Agent Netconf_Agent
}

func (netconf *Netconf) GetEntityData() *types.CommonEntityData {
    netconf.EntityData.YFilter = netconf.YFilter
    netconf.EntityData.YangName = "netconf"
    netconf.EntityData.BundleName = "cisco_ios_xr"
    netconf.EntityData.ParentYangName = "Cisco-IOS-XR-man-xml-ttyagent-oper"
    netconf.EntityData.SegmentPath = "Cisco-IOS-XR-man-xml-ttyagent-oper:netconf"
    netconf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconf.EntityData.Children = types.NewOrderedMap()
    netconf.EntityData.Children.Append("agent", types.YChild{"Agent", &netconf.Agent})
    netconf.EntityData.Leafs = types.NewOrderedMap()

    netconf.EntityData.YListKeys = []string {}

    return &(netconf.EntityData)
}

// Netconf_Agent
// NETCONF agent operational information
type Netconf_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NETCONF agent over TTY.
    Tty Netconf_Agent_Tty
}

func (agent *Netconf_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "netconf"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("tty", types.YChild{"Tty", &agent.Tty})
    agent.EntityData.Leafs = types.NewOrderedMap()

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// Netconf_Agent_Tty
// NETCONF agent over TTY
type Netconf_Agent_Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session information.
    Sessions Netconf_Agent_Tty_Sessions
}

func (tty *Netconf_Agent_Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "agent"
    tty.EntityData.SegmentPath = "tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("sessions", types.YChild{"Sessions", &tty.Sessions})
    tty.EntityData.Leafs = types.NewOrderedMap()

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// Netconf_Agent_Tty_Sessions
// Session information
type Netconf_Agent_Tty_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session information. The type is slice of
    // Netconf_Agent_Tty_Sessions_Session.
    Session []*Netconf_Agent_Tty_Sessions_Session
}

func (sessions *Netconf_Agent_Tty_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "tty"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Netconf_Agent_Tty_Sessions_Session
// Session information
type Netconf_Agent_Tty_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session ID. The type is interface{} with range:
    // 0..4294967295.
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

func (session *Netconf_Agent_Tty_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})
    session.EntityData.Leafs.Append("client-address", types.YLeaf{"ClientAddress", session.ClientAddress})
    session.EntityData.Leafs.Append("client-port", types.YLeaf{"ClientPort", session.ClientPort})
    session.EntityData.Leafs.Append("config-session-id", types.YLeaf{"ConfigSessionId", session.ConfigSessionId})
    session.EntityData.Leafs.Append("admin-config-session-id", types.YLeaf{"AdminConfigSessionId", session.AdminConfigSessionId})
    session.EntityData.Leafs.Append("alarm-notification", types.YLeaf{"AlarmNotification", session.AlarmNotification})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", session.StartTime})
    session.EntityData.Leafs.Append("elapsed-time", types.YLeaf{"ElapsedTime", session.ElapsedTime})
    session.EntityData.Leafs.Append("last-state-change", types.YLeaf{"LastStateChange", session.LastStateChange})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

// XrXml
// xr xml
type XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XML agents.
    Agent XrXml_Agent
}

func (xrXml *XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "Cisco-IOS-XR-man-xml-ttyagent-oper"
    xrXml.EntityData.SegmentPath = "Cisco-IOS-XR-man-xml-ttyagent-oper:xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("agent", types.YChild{"Agent", &xrXml.Agent})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// XrXml_Agent
// XML agents
type XrXml_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TTY sessions information.
    Tty XrXml_Agent_Tty

    // Default sessions information.
    Default XrXml_Agent_Default

    // SSL sessions information.
    Ssl XrXml_Agent_Ssl
}

func (agent *XrXml_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "xr-xml"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("tty", types.YChild{"Tty", &agent.Tty})
    agent.EntityData.Children.Append("default", types.YChild{"Default", &agent.Default})
    agent.EntityData.Children.Append("ssl", types.YChild{"Ssl", &agent.Ssl})
    agent.EntityData.Leafs = types.NewOrderedMap()

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// XrXml_Agent_Tty
// TTY sessions information
type XrXml_Agent_Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Tty_Sessions
}

func (tty *XrXml_Agent_Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "agent"
    tty.EntityData.SegmentPath = "tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("sessions", types.YChild{"Sessions", &tty.Sessions})
    tty.EntityData.Leafs = types.NewOrderedMap()

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// XrXml_Agent_Tty_Sessions
// sessions information
type XrXml_Agent_Tty_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Tty_Sessions_Session.
    Session []*XrXml_Agent_Tty_Sessions_Session
}

func (sessions *XrXml_Agent_Tty_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "tty"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// XrXml_Agent_Tty_Sessions_Session
// xml sessions information
type XrXml_Agent_Tty_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // 0..4294967295.
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

func (session *XrXml_Agent_Tty_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})
    session.EntityData.Leafs.Append("client-address", types.YLeaf{"ClientAddress", session.ClientAddress})
    session.EntityData.Leafs.Append("client-port", types.YLeaf{"ClientPort", session.ClientPort})
    session.EntityData.Leafs.Append("config-session-id", types.YLeaf{"ConfigSessionId", session.ConfigSessionId})
    session.EntityData.Leafs.Append("admin-config-session-id", types.YLeaf{"AdminConfigSessionId", session.AdminConfigSessionId})
    session.EntityData.Leafs.Append("alarm-notification", types.YLeaf{"AlarmNotification", session.AlarmNotification})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", session.StartTime})
    session.EntityData.Leafs.Append("elapsed-time", types.YLeaf{"ElapsedTime", session.ElapsedTime})
    session.EntityData.Leafs.Append("last-state-change", types.YLeaf{"LastStateChange", session.LastStateChange})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

// XrXml_Agent_Default
// Default sessions information
type XrXml_Agent_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Default_Sessions
}

func (self *XrXml_Agent_Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "agent"
    self.EntityData.SegmentPath = "default"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("sessions", types.YChild{"Sessions", &self.Sessions})
    self.EntityData.Leafs = types.NewOrderedMap()

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// XrXml_Agent_Default_Sessions
// sessions information
type XrXml_Agent_Default_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Default_Sessions_Session.
    Session []*XrXml_Agent_Default_Sessions_Session
}

func (sessions *XrXml_Agent_Default_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "default"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// XrXml_Agent_Default_Sessions_Session
// xml sessions information
type XrXml_Agent_Default_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // 0..4294967295.
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

func (session *XrXml_Agent_Default_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})
    session.EntityData.Leafs.Append("client-address", types.YLeaf{"ClientAddress", session.ClientAddress})
    session.EntityData.Leafs.Append("client-port", types.YLeaf{"ClientPort", session.ClientPort})
    session.EntityData.Leafs.Append("config-session-id", types.YLeaf{"ConfigSessionId", session.ConfigSessionId})
    session.EntityData.Leafs.Append("admin-config-session-id", types.YLeaf{"AdminConfigSessionId", session.AdminConfigSessionId})
    session.EntityData.Leafs.Append("alarm-notification", types.YLeaf{"AlarmNotification", session.AlarmNotification})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", session.StartTime})
    session.EntityData.Leafs.Append("elapsed-time", types.YLeaf{"ElapsedTime", session.ElapsedTime})
    session.EntityData.Leafs.Append("last-state-change", types.YLeaf{"LastStateChange", session.LastStateChange})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

// XrXml_Agent_Ssl
// SSL sessions information
type XrXml_Agent_Ssl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sessions information.
    Sessions XrXml_Agent_Ssl_Sessions
}

func (ssl *XrXml_Agent_Ssl) GetEntityData() *types.CommonEntityData {
    ssl.EntityData.YFilter = ssl.YFilter
    ssl.EntityData.YangName = "ssl"
    ssl.EntityData.BundleName = "cisco_ios_xr"
    ssl.EntityData.ParentYangName = "agent"
    ssl.EntityData.SegmentPath = "ssl"
    ssl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssl.EntityData.Children = types.NewOrderedMap()
    ssl.EntityData.Children.Append("sessions", types.YChild{"Sessions", &ssl.Sessions})
    ssl.EntityData.Leafs = types.NewOrderedMap()

    ssl.EntityData.YListKeys = []string {}

    return &(ssl.EntityData)
}

// XrXml_Agent_Ssl_Sessions
// sessions information
type XrXml_Agent_Ssl_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // xml sessions information. The type is slice of
    // XrXml_Agent_Ssl_Sessions_Session.
    Session []*XrXml_Agent_Ssl_Sessions_Session
}

func (sessions *XrXml_Agent_Ssl_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "ssl"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// XrXml_Agent_Ssl_Sessions_Session
// xml sessions information
type XrXml_Agent_Ssl_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // 0..4294967295.
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

func (session *XrXml_Agent_Ssl_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.SessionId, "session-id")
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", session.SessionId})
    session.EntityData.Leafs.Append("username", types.YLeaf{"Username", session.Username})
    session.EntityData.Leafs.Append("state", types.YLeaf{"State", session.State})
    session.EntityData.Leafs.Append("client-address", types.YLeaf{"ClientAddress", session.ClientAddress})
    session.EntityData.Leafs.Append("client-port", types.YLeaf{"ClientPort", session.ClientPort})
    session.EntityData.Leafs.Append("config-session-id", types.YLeaf{"ConfigSessionId", session.ConfigSessionId})
    session.EntityData.Leafs.Append("admin-config-session-id", types.YLeaf{"AdminConfigSessionId", session.AdminConfigSessionId})
    session.EntityData.Leafs.Append("alarm-notification", types.YLeaf{"AlarmNotification", session.AlarmNotification})
    session.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", session.VrfName})
    session.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", session.StartTime})
    session.EntityData.Leafs.Append("elapsed-time", types.YLeaf{"ElapsedTime", session.ElapsedTime})
    session.EntityData.Leafs.Append("last-state-change", types.YLeaf{"LastStateChange", session.LastStateChange})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

