// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-management-cmd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   show-users: Show users statistics
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show users statistics.
    Sessions ShowUsers_Sessions
}

func (showUsers *ShowUsers) GetEntityData() *types.CommonEntityData {
    showUsers.EntityData.YFilter = showUsers.YFilter
    showUsers.EntityData.YangName = "show-users"
    showUsers.EntityData.BundleName = "cisco_ios_xr"
    showUsers.EntityData.ParentYangName = "Cisco-IOS-XR-tty-management-cmd-oper"
    showUsers.EntityData.SegmentPath = "Cisco-IOS-XR-tty-management-cmd-oper:show-users"
    showUsers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    showUsers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    showUsers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    showUsers.EntityData.Children = types.NewOrderedMap()
    showUsers.EntityData.Children.Append("sessions", types.YChild{"Sessions", &showUsers.Sessions})
    showUsers.EntityData.Leafs = types.NewOrderedMap()

    showUsers.EntityData.YListKeys = []string {}

    return &(showUsers.EntityData)
}

// ShowUsers_Sessions
// Show users statistics
type ShowUsers_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Show users statistics. The type is slice of ShowUsers_Sessions_Session.
    Session []*ShowUsers_Sessions_Session
}

func (sessions *ShowUsers_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "show-users"
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

// ShowUsers_Sessions_Session
// Show users statistics
type ShowUsers_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // 0..4294967295.
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

func (session *ShowUsers_Sessions_Session) GetEntityData() *types.CommonEntityData {
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
    session.EntityData.Leafs.Append("line", types.YLeaf{"Line", session.Line})
    session.EntityData.Leafs.Append("user", types.YLeaf{"User", session.User})
    session.EntityData.Leafs.Append("service", types.YLeaf{"Service", session.Service})
    session.EntityData.Leafs.Append("conns", types.YLeaf{"Conns", session.Conns})
    session.EntityData.Leafs.Append("idle-string", types.YLeaf{"IdleString", session.IdleString})
    session.EntityData.Leafs.Append("location", types.YLeaf{"Location", session.Location})

    session.EntityData.YListKeys = []string {"SessionId"}

    return &(session.EntityData)
}

