// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-server package configuration.
// 
// This module contains definitions
// for the following management objects:
//   tty: TTY Line Configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package tty_server_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tty_server_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tty-server-cfg tty}", reflect.TypeOf(Tty{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tty-server-cfg:tty", reflect.TypeOf(Tty{}))
}

// Tty
// TTY Line Configuration
type Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TTY templates.
    TtyLines Tty_TtyLines
}

func (tty *Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "Cisco-IOS-XR-tty-server-cfg"
    tty.EntityData.SegmentPath = "Cisco-IOS-XR-tty-server-cfg:tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("tty-lines", types.YChild{"TtyLines", &tty.TtyLines})
    tty.EntityData.Leafs = types.NewOrderedMap()

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// Tty_TtyLines
// TTY templates
type Tty_TtyLines struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TTY Line,Use string 'console' to configure a console line,Use string
    // 'default' to configure a default line,Use any string to configure a user
    // defined template. The type is slice of Tty_TtyLines_TtyLine.
    TtyLine []*Tty_TtyLines_TtyLine
}

func (ttyLines *Tty_TtyLines) GetEntityData() *types.CommonEntityData {
    ttyLines.EntityData.YFilter = ttyLines.YFilter
    ttyLines.EntityData.YangName = "tty-lines"
    ttyLines.EntityData.BundleName = "cisco_ios_xr"
    ttyLines.EntityData.ParentYangName = "tty"
    ttyLines.EntityData.SegmentPath = "tty-lines"
    ttyLines.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttyLines.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttyLines.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttyLines.EntityData.Children = types.NewOrderedMap()
    ttyLines.EntityData.Children.Append("tty-line", types.YChild{"TtyLine", nil})
    for i := range ttyLines.TtyLine {
        ttyLines.EntityData.Children.Append(types.GetSegmentPath(ttyLines.TtyLine[i]), types.YChild{"TtyLine", ttyLines.TtyLine[i]})
    }
    ttyLines.EntityData.Leafs = types.NewOrderedMap()

    ttyLines.EntityData.YListKeys = []string {}

    return &(ttyLines.EntityData)
}

// Tty_TtyLines_TtyLine
// TTY Line,Use string 'console' to configure a
// console line,Use string 'default' to configure
// a default line,Use any string to configure a
// user defined template
type Tty_TtyLines_TtyLine struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the template. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // TTY line general configuration.
    General Tty_TtyLines_TtyLine_General

    // Telnet protocol-specific configuration.
    Telnet Tty_TtyLines_TtyLine_Telnet

    // Container class for AAA related TTY configuration.
    Aaa Tty_TtyLines_TtyLine_Aaa

    // EXEC timeout and timestamp configurtion.
    Exec Tty_TtyLines_TtyLine_Exec

    // Management connection configuration.
    Connection Tty_TtyLines_TtyLine_Connection

    // Exec Mode Pager  configurtion.
    ExecMode Tty_TtyLines_TtyLine_ExecMode
}

func (ttyLine *Tty_TtyLines_TtyLine) GetEntityData() *types.CommonEntityData {
    ttyLine.EntityData.YFilter = ttyLine.YFilter
    ttyLine.EntityData.YangName = "tty-line"
    ttyLine.EntityData.BundleName = "cisco_ios_xr"
    ttyLine.EntityData.ParentYangName = "tty-lines"
    ttyLine.EntityData.SegmentPath = "tty-line" + types.AddKeyToken(ttyLine.Name, "name")
    ttyLine.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttyLine.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttyLine.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttyLine.EntityData.Children = types.NewOrderedMap()
    ttyLine.EntityData.Children.Append("general", types.YChild{"General", &ttyLine.General})
    ttyLine.EntityData.Children.Append("telnet", types.YChild{"Telnet", &ttyLine.Telnet})
    ttyLine.EntityData.Children.Append("aaa", types.YChild{"Aaa", &ttyLine.Aaa})
    ttyLine.EntityData.Children.Append("exec", types.YChild{"Exec", &ttyLine.Exec})
    ttyLine.EntityData.Children.Append("Cisco-IOS-XR-tty-management-cfg:connection", types.YChild{"Connection", &ttyLine.Connection})
    ttyLine.EntityData.Children.Append("Cisco-IOS-XR-tty-management-cfg:exec-mode", types.YChild{"ExecMode", &ttyLine.ExecMode})
    ttyLine.EntityData.Leafs = types.NewOrderedMap()
    ttyLine.EntityData.Leafs.Append("name", types.YLeaf{"Name", ttyLine.Name})

    ttyLine.EntityData.YListKeys = []string {"Name"}

    return &(ttyLine.EntityData)
}

// Tty_TtyLines_TtyLine_General
// TTY line general configuration
type Tty_TtyLines_TtyLine_General struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of lines on a screen. The type is interface{} with range: 0..512.
    // The default value is 24.
    Length interface{}

    // Absolute timeout for line disconnection. The type is interface{} with
    // range: 0..10000. Units are minute. The default value is 0.
    AbsoluteTimeout interface{}

    // Number of characters on a screen line. The type is interface{} with range:
    // 0..512. The default value is 80.
    Width interface{}
}

func (general *Tty_TtyLines_TtyLine_General) GetEntityData() *types.CommonEntityData {
    general.EntityData.YFilter = general.YFilter
    general.EntityData.YangName = "general"
    general.EntityData.BundleName = "cisco_ios_xr"
    general.EntityData.ParentYangName = "tty-line"
    general.EntityData.SegmentPath = "general"
    general.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    general.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    general.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    general.EntityData.Children = types.NewOrderedMap()
    general.EntityData.Leafs = types.NewOrderedMap()
    general.EntityData.Leafs.Append("length", types.YLeaf{"Length", general.Length})
    general.EntityData.Leafs.Append("absolute-timeout", types.YLeaf{"AbsoluteTimeout", general.AbsoluteTimeout})
    general.EntityData.Leafs.Append("width", types.YLeaf{"Width", general.Width})

    general.EntityData.YListKeys = []string {}

    return &(general.EntityData)
}

// Tty_TtyLines_TtyLine_Telnet
// Telnet protocol-specific configuration
type Tty_TtyLines_TtyLine_Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Send a CR as a CR followed by a NULL instead of a CRfollowed by a LF. The
    // type is interface{}.
    Transparent interface{}
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetEntityData() *types.CommonEntityData {
    telnet.EntityData.YFilter = telnet.YFilter
    telnet.EntityData.YangName = "telnet"
    telnet.EntityData.BundleName = "cisco_ios_xr"
    telnet.EntityData.ParentYangName = "tty-line"
    telnet.EntityData.SegmentPath = "telnet"
    telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnet.EntityData.Children = types.NewOrderedMap()
    telnet.EntityData.Leafs = types.NewOrderedMap()
    telnet.EntityData.Leafs.Append("transparent", types.YLeaf{"Transparent", telnet.Transparent})

    telnet.EntityData.YListKeys = []string {}

    return &(telnet.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa
// Container class for AAA related TTY
// configuration
type Tty_TtyLines_TtyLine_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeouts for any user input during login sequence. The type is interface{}
    // with range: 0..300. Units are second. The default value is 30.
    LoginTimeout interface{}

    // Configure a secure one way encrypted password. The type is string with
    // pattern: (!.+)|([^!].+).
    Secret interface{}

    // Configure the password for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Users characteristics.
    UserGroups Tty_TtyLines_TtyLine_Aaa_UserGroups

    // Authorization parameters.
    Authorization Tty_TtyLines_TtyLine_Aaa_Authorization

    // Authentication parameters.
    Authentication Tty_TtyLines_TtyLine_Aaa_Authentication

    // Accounting parameters.
    Accounting Tty_TtyLines_TtyLine_Aaa_Accounting
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "tty-line"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("user-groups", types.YChild{"UserGroups", &aaa.UserGroups})
    aaa.EntityData.Children.Append("authorization", types.YChild{"Authorization", &aaa.Authorization})
    aaa.EntityData.Children.Append("authentication", types.YChild{"Authentication", &aaa.Authentication})
    aaa.EntityData.Children.Append("accounting", types.YChild{"Accounting", &aaa.Accounting})
    aaa.EntityData.Leafs = types.NewOrderedMap()
    aaa.EntityData.Leafs.Append("login-timeout", types.YLeaf{"LoginTimeout", aaa.LoginTimeout})
    aaa.EntityData.Leafs.Append("secret", types.YLeaf{"Secret", aaa.Secret})
    aaa.EntityData.Leafs.Append("password", types.YLeaf{"Password", aaa.Password})

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa_UserGroups
// Users characteristics
type Tty_TtyLines_TtyLine_Aaa_UserGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Group to which the user will belong. The type is slice of
    // Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup.
    UserGroup []*Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetEntityData() *types.CommonEntityData {
    userGroups.EntityData.YFilter = userGroups.YFilter
    userGroups.EntityData.YangName = "user-groups"
    userGroups.EntityData.BundleName = "cisco_ios_xr"
    userGroups.EntityData.ParentYangName = "aaa"
    userGroups.EntityData.SegmentPath = "user-groups"
    userGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userGroups.EntityData.Children = types.NewOrderedMap()
    userGroups.EntityData.Children.Append("user-group", types.YChild{"UserGroup", nil})
    for i := range userGroups.UserGroup {
        userGroups.EntityData.Children.Append(types.GetSegmentPath(userGroups.UserGroup[i]), types.YChild{"UserGroup", userGroups.UserGroup[i]})
    }
    userGroups.EntityData.Leafs = types.NewOrderedMap()

    userGroups.EntityData.YListKeys = []string {}

    return &(userGroups.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup
// Group to which the user will belong
type Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Specify as 'root-system' for root-system group and 'other' for remaining
    // groups. The type is string. This attribute is mandatory.
    Category interface{}
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetEntityData() *types.CommonEntityData {
    userGroup.EntityData.YFilter = userGroup.YFilter
    userGroup.EntityData.YangName = "user-group"
    userGroup.EntityData.BundleName = "cisco_ios_xr"
    userGroup.EntityData.ParentYangName = "user-groups"
    userGroup.EntityData.SegmentPath = "user-group" + types.AddKeyToken(userGroup.Name, "name")
    userGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userGroup.EntityData.Children = types.NewOrderedMap()
    userGroup.EntityData.Leafs = types.NewOrderedMap()
    userGroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", userGroup.Name})
    userGroup.EntityData.Leafs.Append("category", types.YLeaf{"Category", userGroup.Category})

    userGroup.EntityData.YListKeys = []string {"Name"}

    return &(userGroup.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa_Authorization
// Authorization parameters
type Tty_TtyLines_TtyLine_Aaa_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // For starting an exec (shell). The type is string.
    Exec interface{}

    // Specify 'default' or use an authorization list with this name. The type is
    // string.
    EventManager interface{}

    // For exec (shell) configuration. The type is string.
    Commands interface{}
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "aaa"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("exec", types.YLeaf{"Exec", authorization.Exec})
    authorization.EntityData.Leafs.Append("event-manager", types.YLeaf{"EventManager", authorization.EventManager})
    authorization.EntityData.Leafs.Append("commands", types.YLeaf{"Commands", authorization.Commands})

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa_Authentication
// Authentication parameters
type Tty_TtyLines_TtyLine_Aaa_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication list name. The type is string.
    Login interface{}
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "aaa"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("login", types.YLeaf{"Login", authentication.Login})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Tty_TtyLines_TtyLine_Aaa_Accounting
// Accounting parameters
type Tty_TtyLines_TtyLine_Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // For starting an exec (shell). The type is string.
    Exec interface{}

    // For exec (shell) configuration. The type is string.
    Commands interface{}
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("exec", types.YLeaf{"Exec", accounting.Exec})
    accounting.EntityData.Leafs.Append("commands", types.YLeaf{"Commands", accounting.Commands})

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Tty_TtyLines_TtyLine_Exec
// EXEC timeout and timestamp configurtion
type Tty_TtyLines_TtyLine_Exec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // 'True' to Enable & 'False' to Disable time stamp. The type is bool.
    TimeStamp interface{}

    // EXEC Timeout.
    Timeout Tty_TtyLines_TtyLine_Exec_Timeout
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetEntityData() *types.CommonEntityData {
    exec.EntityData.YFilter = exec.YFilter
    exec.EntityData.YangName = "exec"
    exec.EntityData.BundleName = "cisco_ios_xr"
    exec.EntityData.ParentYangName = "tty-line"
    exec.EntityData.SegmentPath = "exec"
    exec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exec.EntityData.Children = types.NewOrderedMap()
    exec.EntityData.Children.Append("timeout", types.YChild{"Timeout", &exec.Timeout})
    exec.EntityData.Leafs = types.NewOrderedMap()
    exec.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", exec.TimeStamp})

    exec.EntityData.YListKeys = []string {}

    return &(exec.EntityData)
}

// Tty_TtyLines_TtyLine_Exec_Timeout
// EXEC Timeout
// This type is a presence type.
type Tty_TtyLines_TtyLine_Exec_Timeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Timeout in minutes. The type is interface{} with range: 0..35791. This
    // attribute is mandatory. Units are minute.
    Minutes interface{}

    // Timeout in seconds. The type is interface{} with range: 0..2147483. This
    // attribute is mandatory. Units are second.
    Seconds interface{}
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetEntityData() *types.CommonEntityData {
    timeout.EntityData.YFilter = timeout.YFilter
    timeout.EntityData.YangName = "timeout"
    timeout.EntityData.BundleName = "cisco_ios_xr"
    timeout.EntityData.ParentYangName = "exec"
    timeout.EntityData.SegmentPath = "timeout"
    timeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeout.EntityData.Children = types.NewOrderedMap()
    timeout.EntityData.Leafs = types.NewOrderedMap()
    timeout.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", timeout.Minutes})
    timeout.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", timeout.Seconds})

    timeout.EntityData.YListKeys = []string {}

    return &(timeout.EntityData)
}

// Tty_TtyLines_TtyLine_Connection
// Management connection configuration
type Tty_TtyLines_TtyLine_Connection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disconnect character's decimal equivalent value or Character . The type is
    // one of the following types: string with pattern:
    // (\p{IsBasicLatin}|\p{IsLatin-1Supplement})*, or int with range: 0..255.
    DisconnectCharacter interface{}

    // ACL to filter ingoing connections. The type is string.
    AclIn interface{}

    // ACL to filter outgoing connections. The type is string.
    AclOut interface{}

    // Command completion on whitespace. The type is interface{}.
    CliWhiteSpaceCompletion interface{}

    // The number of outgoing connections. The type is interface{} with range:
    // 0..20. The default value is 6.
    SessionLimit interface{}

    // Escape character or ASCII decimal equivalent value orspecial strings
    // NONE,DEFAULT,BREAK. The type is one of the following types: string with
    // pattern:
    // (\p{IsBasicLatin}|\p{IsLatin-1Supplement})|(DEFAULT)|(BREAK)|(NONE) The
    // default value is 30., or int with range: 0..255 The default value is 30..
    EscapeCharacter interface{}

    // The preferred protocol to use. The type is TtyTransportProtocol.
    TransportPreferred interface{}

    // Protocols to use when connecting to the terminal server.
    TransportInput Tty_TtyLines_TtyLine_Connection_TransportInput

    // Protocols to use for outgoing connections.
    TransportOutput Tty_TtyLines_TtyLine_Connection_TransportOutput

    // Interval for closing connection when there is no input traffic.
    SessionTimeout Tty_TtyLines_TtyLine_Connection_SessionTimeout
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetEntityData() *types.CommonEntityData {
    connection.EntityData.YFilter = connection.YFilter
    connection.EntityData.YangName = "connection"
    connection.EntityData.BundleName = "cisco_ios_xr"
    connection.EntityData.ParentYangName = "tty-line"
    connection.EntityData.SegmentPath = "Cisco-IOS-XR-tty-management-cfg:connection"
    connection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connection.EntityData.Children = types.NewOrderedMap()
    connection.EntityData.Children.Append("transport-input", types.YChild{"TransportInput", &connection.TransportInput})
    connection.EntityData.Children.Append("transport-output", types.YChild{"TransportOutput", &connection.TransportOutput})
    connection.EntityData.Children.Append("session-timeout", types.YChild{"SessionTimeout", &connection.SessionTimeout})
    connection.EntityData.Leafs = types.NewOrderedMap()
    connection.EntityData.Leafs.Append("disconnect-character", types.YLeaf{"DisconnectCharacter", connection.DisconnectCharacter})
    connection.EntityData.Leafs.Append("acl-in", types.YLeaf{"AclIn", connection.AclIn})
    connection.EntityData.Leafs.Append("acl-out", types.YLeaf{"AclOut", connection.AclOut})
    connection.EntityData.Leafs.Append("cli-white-space-completion", types.YLeaf{"CliWhiteSpaceCompletion", connection.CliWhiteSpaceCompletion})
    connection.EntityData.Leafs.Append("session-limit", types.YLeaf{"SessionLimit", connection.SessionLimit})
    connection.EntityData.Leafs.Append("escape-character", types.YLeaf{"EscapeCharacter", connection.EscapeCharacter})
    connection.EntityData.Leafs.Append("transport-preferred", types.YLeaf{"TransportPreferred", connection.TransportPreferred})

    connection.EntityData.YListKeys = []string {}

    return &(connection.EntityData)
}

// Tty_TtyLines_TtyLine_Connection_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_TtyLines_TtyLine_Connection_TransportInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: 0..4294967295.
    None interface{}
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetEntityData() *types.CommonEntityData {
    transportInput.EntityData.YFilter = transportInput.YFilter
    transportInput.EntityData.YangName = "transport-input"
    transportInput.EntityData.BundleName = "cisco_ios_xr"
    transportInput.EntityData.ParentYangName = "connection"
    transportInput.EntityData.SegmentPath = "transport-input"
    transportInput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportInput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportInput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportInput.EntityData.Children = types.NewOrderedMap()
    transportInput.EntityData.Leafs = types.NewOrderedMap()
    transportInput.EntityData.Leafs.Append("select", types.YLeaf{"Select", transportInput.Select})
    transportInput.EntityData.Leafs.Append("protocol1", types.YLeaf{"Protocol1", transportInput.Protocol1})
    transportInput.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", transportInput.Protocol2})
    transportInput.EntityData.Leafs.Append("none", types.YLeaf{"None", transportInput.None})

    transportInput.EntityData.YListKeys = []string {}

    return &(transportInput.EntityData)
}

// Tty_TtyLines_TtyLine_Connection_TransportOutput
// Protocols to use for outgoing connections
// This type is a presence type.
type Tty_TtyLines_TtyLine_Connection_TransportOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Choose transport protocols. The type is TtyTransportProtocolSelect. This
    // attribute is mandatory.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: 0..4294967295.
    None interface{}
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetEntityData() *types.CommonEntityData {
    transportOutput.EntityData.YFilter = transportOutput.YFilter
    transportOutput.EntityData.YangName = "transport-output"
    transportOutput.EntityData.BundleName = "cisco_ios_xr"
    transportOutput.EntityData.ParentYangName = "connection"
    transportOutput.EntityData.SegmentPath = "transport-output"
    transportOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportOutput.EntityData.Children = types.NewOrderedMap()
    transportOutput.EntityData.Leafs = types.NewOrderedMap()
    transportOutput.EntityData.Leafs.Append("select", types.YLeaf{"Select", transportOutput.Select})
    transportOutput.EntityData.Leafs.Append("protocol1", types.YLeaf{"Protocol1", transportOutput.Protocol1})
    transportOutput.EntityData.Leafs.Append("protocol2", types.YLeaf{"Protocol2", transportOutput.Protocol2})
    transportOutput.EntityData.Leafs.Append("none", types.YLeaf{"None", transportOutput.None})

    transportOutput.EntityData.YListKeys = []string {}

    return &(transportOutput.EntityData)
}

// Tty_TtyLines_TtyLine_Connection_SessionTimeout
// Interval for closing connection when there is
// no input traffic
// This type is a presence type.
type Tty_TtyLines_TtyLine_Connection_SessionTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Session timeout interval in minutes. The type is interface{} with range:
    // 0..35791. This attribute is mandatory.
    Timeout interface{}

    // Include output traffic as well as input traffic. The type is
    // TtySessionTimeoutDirection. This attribute is mandatory.
    Direction interface{}
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetEntityData() *types.CommonEntityData {
    sessionTimeout.EntityData.YFilter = sessionTimeout.YFilter
    sessionTimeout.EntityData.YangName = "session-timeout"
    sessionTimeout.EntityData.BundleName = "cisco_ios_xr"
    sessionTimeout.EntityData.ParentYangName = "connection"
    sessionTimeout.EntityData.SegmentPath = "session-timeout"
    sessionTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionTimeout.EntityData.Children = types.NewOrderedMap()
    sessionTimeout.EntityData.Leafs = types.NewOrderedMap()
    sessionTimeout.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", sessionTimeout.Timeout})
    sessionTimeout.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", sessionTimeout.Direction})

    sessionTimeout.EntityData.YListKeys = []string {}

    return &(sessionTimeout.EntityData)
}

// Tty_TtyLines_TtyLine_ExecMode
// Exec Mode Pager  configurtion
type Tty_TtyLines_TtyLine_ExecMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Preferred Paging Utility. The type is TtyPager. The default value is more.
    Pager interface{}
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetEntityData() *types.CommonEntityData {
    execMode.EntityData.YFilter = execMode.YFilter
    execMode.EntityData.YangName = "exec-mode"
    execMode.EntityData.BundleName = "cisco_ios_xr"
    execMode.EntityData.ParentYangName = "tty-line"
    execMode.EntityData.SegmentPath = "Cisco-IOS-XR-tty-management-cfg:exec-mode"
    execMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    execMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    execMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    execMode.EntityData.Children = types.NewOrderedMap()
    execMode.EntityData.Leafs = types.NewOrderedMap()
    execMode.EntityData.Leafs.Append("pager", types.YLeaf{"Pager", execMode.Pager})

    execMode.EntityData.YListKeys = []string {}

    return &(execMode.EntityData)
}

