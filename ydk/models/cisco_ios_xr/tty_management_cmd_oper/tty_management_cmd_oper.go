// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-management-cmd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   show-users: Show users statistics
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tty_management_cmd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_management_cmd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tty-management-cmd-oper show-users}", reflect.TypeOf(ShowUsers{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tty-management-cmd-oper:show-users", reflect.TypeOf(ShowUsers{}))
}

// ShowUsers
// Show users statistics
type ShowUsers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show users statistics.
    Sessions ShowUsers_Sessions
}

func (showUsers *ShowUsers) GetFilter() yfilter.YFilter { return showUsers.YFilter }

func (showUsers *ShowUsers) SetFilter(yf yfilter.YFilter) { showUsers.YFilter = yf }

func (showUsers *ShowUsers) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (showUsers *ShowUsers) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-management-cmd-oper:show-users"
}

func (showUsers *ShowUsers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &showUsers.Sessions
    }
    return nil
}

func (showUsers *ShowUsers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &showUsers.Sessions
    return children
}

func (showUsers *ShowUsers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (showUsers *ShowUsers) GetBundleName() string { return "cisco_ios_xr" }

func (showUsers *ShowUsers) GetYangName() string { return "show-users" }

func (showUsers *ShowUsers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (showUsers *ShowUsers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (showUsers *ShowUsers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (showUsers *ShowUsers) SetParent(parent types.Entity) { showUsers.parent = parent }

func (showUsers *ShowUsers) GetParent() types.Entity { return showUsers.parent }

func (showUsers *ShowUsers) GetParentYangName() string { return "Cisco-IOS-XR-tty-management-cmd-oper" }

// ShowUsers_Sessions
// Show users statistics
type ShowUsers_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Show users statistics. The type is slice of ShowUsers_Sessions_Session.
    Session []ShowUsers_Sessions_Session
}

func (sessions *ShowUsers_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *ShowUsers_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *ShowUsers_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *ShowUsers_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *ShowUsers_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ShowUsers_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *ShowUsers_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *ShowUsers_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *ShowUsers_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *ShowUsers_Sessions) GetYangName() string { return "sessions" }

func (sessions *ShowUsers_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *ShowUsers_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *ShowUsers_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *ShowUsers_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *ShowUsers_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *ShowUsers_Sessions) GetParentYangName() string { return "show-users" }

// ShowUsers_Sessions_Session
// Show users statistics
type ShowUsers_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // Line Number. The type is string.
    Line interface{}

    // User Name. The type is string.
    User interface{}

    // Service Name. The type is string.
    Service interface{}

    // No. of Connections. The type is string.
    Conns interface{}

    // Idle Time. The type is string.
    IdleString interface{}

    // location. The type is string.
    Location interface{}
}

func (session *ShowUsers_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *ShowUsers_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *ShowUsers_Sessions_Session) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "line" { return "Line" }
    if yname == "user" { return "User" }
    if yname == "service" { return "Service" }
    if yname == "conns" { return "Conns" }
    if yname == "idle-string" { return "IdleString" }
    if yname == "location" { return "Location" }
    return ""
}

func (session *ShowUsers_Sessions_Session) GetSegmentPath() string {
    return "session" + "[session-id='" + fmt.Sprintf("%v", session.SessionId) + "']"
}

func (session *ShowUsers_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *ShowUsers_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *ShowUsers_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = session.SessionId
    leafs["line"] = session.Line
    leafs["user"] = session.User
    leafs["service"] = session.Service
    leafs["conns"] = session.Conns
    leafs["idle-string"] = session.IdleString
    leafs["location"] = session.Location
    return leafs
}

func (session *ShowUsers_Sessions_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *ShowUsers_Sessions_Session) GetYangName() string { return "session" }

func (session *ShowUsers_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *ShowUsers_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *ShowUsers_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *ShowUsers_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *ShowUsers_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *ShowUsers_Sessions_Session) GetParentYangName() string { return "sessions" }

