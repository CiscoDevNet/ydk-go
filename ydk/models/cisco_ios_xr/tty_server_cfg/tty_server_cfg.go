// This module contains a collection of YANG definitions
// for Cisco IOS-XR tty-server package configuration.
// 
// This module contains definitions
// for the following management objects:
//   tty: TTY Line Configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // TTY templates.
    TtyLines Tty_TtyLines
}

func (tty *Tty) GetFilter() yfilter.YFilter { return tty.YFilter }

func (tty *Tty) SetFilter(yf yfilter.YFilter) { tty.YFilter = yf }

func (tty *Tty) GetGoName(yname string) string {
    if yname == "tty-lines" { return "TtyLines" }
    return ""
}

func (tty *Tty) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-server-cfg:tty"
}

func (tty *Tty) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tty-lines" {
        return &tty.TtyLines
    }
    return nil
}

func (tty *Tty) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tty-lines"] = &tty.TtyLines
    return children
}

func (tty *Tty) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tty *Tty) GetBundleName() string { return "cisco_ios_xr" }

func (tty *Tty) GetYangName() string { return "tty" }

func (tty *Tty) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tty *Tty) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tty *Tty) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tty *Tty) SetParent(parent types.Entity) { tty.parent = parent }

func (tty *Tty) GetParent() types.Entity { return tty.parent }

func (tty *Tty) GetParentYangName() string { return "Cisco-IOS-XR-tty-server-cfg" }

// Tty_TtyLines
// TTY templates
type Tty_TtyLines struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TTY Line,Use string 'console' to configure a console line,Use string
    // 'default' to configure a default line,Use any string to configure a user
    // defined template. The type is slice of Tty_TtyLines_TtyLine.
    TtyLine []Tty_TtyLines_TtyLine
}

func (ttyLines *Tty_TtyLines) GetFilter() yfilter.YFilter { return ttyLines.YFilter }

func (ttyLines *Tty_TtyLines) SetFilter(yf yfilter.YFilter) { ttyLines.YFilter = yf }

func (ttyLines *Tty_TtyLines) GetGoName(yname string) string {
    if yname == "tty-line" { return "TtyLine" }
    return ""
}

func (ttyLines *Tty_TtyLines) GetSegmentPath() string {
    return "tty-lines"
}

func (ttyLines *Tty_TtyLines) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tty-line" {
        for _, c := range ttyLines.TtyLine {
            if ttyLines.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_TtyLines_TtyLine{}
        ttyLines.TtyLine = append(ttyLines.TtyLine, child)
        return &ttyLines.TtyLine[len(ttyLines.TtyLine)-1]
    }
    return nil
}

func (ttyLines *Tty_TtyLines) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ttyLines.TtyLine {
        children[ttyLines.TtyLine[i].GetSegmentPath()] = &ttyLines.TtyLine[i]
    }
    return children
}

func (ttyLines *Tty_TtyLines) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ttyLines *Tty_TtyLines) GetBundleName() string { return "cisco_ios_xr" }

func (ttyLines *Tty_TtyLines) GetYangName() string { return "tty-lines" }

func (ttyLines *Tty_TtyLines) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ttyLines *Tty_TtyLines) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ttyLines *Tty_TtyLines) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ttyLines *Tty_TtyLines) SetParent(parent types.Entity) { ttyLines.parent = parent }

func (ttyLines *Tty_TtyLines) GetParent() types.Entity { return ttyLines.parent }

func (ttyLines *Tty_TtyLines) GetParentYangName() string { return "tty" }

// Tty_TtyLines_TtyLine
// TTY Line,Use string 'console' to configure a
// console line,Use string 'default' to configure
// a default line,Use any string to configure a
// user defined template
type Tty_TtyLines_TtyLine struct {
    parent types.Entity
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

func (ttyLine *Tty_TtyLines_TtyLine) GetFilter() yfilter.YFilter { return ttyLine.YFilter }

func (ttyLine *Tty_TtyLines_TtyLine) SetFilter(yf yfilter.YFilter) { ttyLine.YFilter = yf }

func (ttyLine *Tty_TtyLines_TtyLine) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "general" { return "General" }
    if yname == "telnet" { return "Telnet" }
    if yname == "aaa" { return "Aaa" }
    if yname == "exec" { return "Exec" }
    if yname == "Cisco-IOS-XR-tty-management-cfg:connection" { return "Connection" }
    if yname == "Cisco-IOS-XR-tty-management-cfg:exec-mode" { return "ExecMode" }
    return ""
}

func (ttyLine *Tty_TtyLines_TtyLine) GetSegmentPath() string {
    return "tty-line" + "[name='" + fmt.Sprintf("%v", ttyLine.Name) + "']"
}

func (ttyLine *Tty_TtyLines_TtyLine) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "general" {
        return &ttyLine.General
    }
    if childYangName == "telnet" {
        return &ttyLine.Telnet
    }
    if childYangName == "aaa" {
        return &ttyLine.Aaa
    }
    if childYangName == "exec" {
        return &ttyLine.Exec
    }
    if childYangName == "Cisco-IOS-XR-tty-management-cfg:connection" {
        return &ttyLine.Connection
    }
    if childYangName == "Cisco-IOS-XR-tty-management-cfg:exec-mode" {
        return &ttyLine.ExecMode
    }
    return nil
}

func (ttyLine *Tty_TtyLines_TtyLine) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["general"] = &ttyLine.General
    children["telnet"] = &ttyLine.Telnet
    children["aaa"] = &ttyLine.Aaa
    children["exec"] = &ttyLine.Exec
    children["Cisco-IOS-XR-tty-management-cfg:connection"] = &ttyLine.Connection
    children["Cisco-IOS-XR-tty-management-cfg:exec-mode"] = &ttyLine.ExecMode
    return children
}

func (ttyLine *Tty_TtyLines_TtyLine) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ttyLine.Name
    return leafs
}

func (ttyLine *Tty_TtyLines_TtyLine) GetBundleName() string { return "cisco_ios_xr" }

func (ttyLine *Tty_TtyLines_TtyLine) GetYangName() string { return "tty-line" }

func (ttyLine *Tty_TtyLines_TtyLine) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ttyLine *Tty_TtyLines_TtyLine) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ttyLine *Tty_TtyLines_TtyLine) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ttyLine *Tty_TtyLines_TtyLine) SetParent(parent types.Entity) { ttyLine.parent = parent }

func (ttyLine *Tty_TtyLines_TtyLine) GetParent() types.Entity { return ttyLine.parent }

func (ttyLine *Tty_TtyLines_TtyLine) GetParentYangName() string { return "tty-lines" }

// Tty_TtyLines_TtyLine_General
// TTY line general configuration
type Tty_TtyLines_TtyLine_General struct {
    parent types.Entity
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

func (general *Tty_TtyLines_TtyLine_General) GetFilter() yfilter.YFilter { return general.YFilter }

func (general *Tty_TtyLines_TtyLine_General) SetFilter(yf yfilter.YFilter) { general.YFilter = yf }

func (general *Tty_TtyLines_TtyLine_General) GetGoName(yname string) string {
    if yname == "length" { return "Length" }
    if yname == "absolute-timeout" { return "AbsoluteTimeout" }
    if yname == "width" { return "Width" }
    return ""
}

func (general *Tty_TtyLines_TtyLine_General) GetSegmentPath() string {
    return "general"
}

func (general *Tty_TtyLines_TtyLine_General) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (general *Tty_TtyLines_TtyLine_General) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (general *Tty_TtyLines_TtyLine_General) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["length"] = general.Length
    leafs["absolute-timeout"] = general.AbsoluteTimeout
    leafs["width"] = general.Width
    return leafs
}

func (general *Tty_TtyLines_TtyLine_General) GetBundleName() string { return "cisco_ios_xr" }

func (general *Tty_TtyLines_TtyLine_General) GetYangName() string { return "general" }

func (general *Tty_TtyLines_TtyLine_General) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (general *Tty_TtyLines_TtyLine_General) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (general *Tty_TtyLines_TtyLine_General) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (general *Tty_TtyLines_TtyLine_General) SetParent(parent types.Entity) { general.parent = parent }

func (general *Tty_TtyLines_TtyLine_General) GetParent() types.Entity { return general.parent }

func (general *Tty_TtyLines_TtyLine_General) GetParentYangName() string { return "tty-line" }

// Tty_TtyLines_TtyLine_Telnet
// Telnet protocol-specific configuration
type Tty_TtyLines_TtyLine_Telnet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Send a CR as a CR followed by a NULL instead of a CRfollowed by a LF. The
    // type is interface{}.
    Transparent interface{}
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetFilter() yfilter.YFilter { return telnet.YFilter }

func (telnet *Tty_TtyLines_TtyLine_Telnet) SetFilter(yf yfilter.YFilter) { telnet.YFilter = yf }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetGoName(yname string) string {
    if yname == "transparent" { return "Transparent" }
    return ""
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetSegmentPath() string {
    return "telnet"
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transparent"] = telnet.Transparent
    return leafs
}

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetYangName() string { return "telnet" }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (telnet *Tty_TtyLines_TtyLine_Telnet) SetParent(parent types.Entity) { telnet.parent = parent }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetParent() types.Entity { return telnet.parent }

func (telnet *Tty_TtyLines_TtyLine_Telnet) GetParentYangName() string { return "tty-line" }

// Tty_TtyLines_TtyLine_Aaa
// Container class for AAA related TTY
// configuration
type Tty_TtyLines_TtyLine_Aaa struct {
    parent types.Entity
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

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Tty_TtyLines_TtyLine_Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetGoName(yname string) string {
    if yname == "login-timeout" { return "LoginTimeout" }
    if yname == "secret" { return "Secret" }
    if yname == "password" { return "Password" }
    if yname == "user-groups" { return "UserGroups" }
    if yname == "authorization" { return "Authorization" }
    if yname == "authentication" { return "Authentication" }
    if yname == "accounting" { return "Accounting" }
    return ""
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetSegmentPath() string {
    return "aaa"
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user-groups" {
        return &aaa.UserGroups
    }
    if childYangName == "authorization" {
        return &aaa.Authorization
    }
    if childYangName == "authentication" {
        return &aaa.Authentication
    }
    if childYangName == "accounting" {
        return &aaa.Accounting
    }
    return nil
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["user-groups"] = &aaa.UserGroups
    children["authorization"] = &aaa.Authorization
    children["authentication"] = &aaa.Authentication
    children["accounting"] = &aaa.Accounting
    return children
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["login-timeout"] = aaa.LoginTimeout
    leafs["secret"] = aaa.Secret
    leafs["password"] = aaa.Password
    return leafs
}

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetYangName() string { return "aaa" }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Tty_TtyLines_TtyLine_Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Tty_TtyLines_TtyLine_Aaa) GetParentYangName() string { return "tty-line" }

// Tty_TtyLines_TtyLine_Aaa_UserGroups
// Users characteristics
type Tty_TtyLines_TtyLine_Aaa_UserGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Group to which the user will belong. The type is slice of
    // Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup.
    UserGroup []Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetFilter() yfilter.YFilter { return userGroups.YFilter }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) SetFilter(yf yfilter.YFilter) { userGroups.YFilter = yf }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetGoName(yname string) string {
    if yname == "user-group" { return "UserGroup" }
    return ""
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetSegmentPath() string {
    return "user-groups"
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user-group" {
        for _, c := range userGroups.UserGroup {
            if userGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup{}
        userGroups.UserGroup = append(userGroups.UserGroup, child)
        return &userGroups.UserGroup[len(userGroups.UserGroup)-1]
    }
    return nil
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range userGroups.UserGroup {
        children[userGroups.UserGroup[i].GetSegmentPath()] = &userGroups.UserGroup[i]
    }
    return children
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetBundleName() string { return "cisco_ios_xr" }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetYangName() string { return "user-groups" }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) SetParent(parent types.Entity) { userGroups.parent = parent }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetParent() types.Entity { return userGroups.parent }

func (userGroups *Tty_TtyLines_TtyLine_Aaa_UserGroups) GetParentYangName() string { return "aaa" }

// Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup
// Group to which the user will belong
type Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Specify as 'root-system' for root-system group and 'other' for remaining
    // groups. The type is string. This attribute is mandatory.
    Category interface{}
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetFilter() yfilter.YFilter { return userGroup.YFilter }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) SetFilter(yf yfilter.YFilter) { userGroup.YFilter = yf }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "category" { return "Category" }
    return ""
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetSegmentPath() string {
    return "user-group" + "[name='" + fmt.Sprintf("%v", userGroup.Name) + "']"
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = userGroup.Name
    leafs["category"] = userGroup.Category
    return leafs
}

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetBundleName() string { return "cisco_ios_xr" }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetYangName() string { return "user-group" }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) SetParent(parent types.Entity) { userGroup.parent = parent }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetParent() types.Entity { return userGroup.parent }

func (userGroup *Tty_TtyLines_TtyLine_Aaa_UserGroups_UserGroup) GetParentYangName() string { return "user-groups" }

// Tty_TtyLines_TtyLine_Aaa_Authorization
// Authorization parameters
type Tty_TtyLines_TtyLine_Aaa_Authorization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // For starting an exec (shell). The type is string.
    Exec interface{}

    // Specify 'default' or use an authorization list with this name. The type is
    // string.
    EventManager interface{}

    // For exec (shell) configuration. The type is string.
    Commands interface{}
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetFilter() yfilter.YFilter { return authorization.YFilter }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) SetFilter(yf yfilter.YFilter) { authorization.YFilter = yf }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetGoName(yname string) string {
    if yname == "exec" { return "Exec" }
    if yname == "event-manager" { return "EventManager" }
    if yname == "commands" { return "Commands" }
    return ""
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetSegmentPath() string {
    return "authorization"
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exec"] = authorization.Exec
    leafs["event-manager"] = authorization.EventManager
    leafs["commands"] = authorization.Commands
    return leafs
}

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetBundleName() string { return "cisco_ios_xr" }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetYangName() string { return "authorization" }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) SetParent(parent types.Entity) { authorization.parent = parent }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetParent() types.Entity { return authorization.parent }

func (authorization *Tty_TtyLines_TtyLine_Aaa_Authorization) GetParentYangName() string { return "aaa" }

// Tty_TtyLines_TtyLine_Aaa_Authentication
// Authentication parameters
type Tty_TtyLines_TtyLine_Aaa_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Authentication list name. The type is string.
    Login interface{}
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetGoName(yname string) string {
    if yname == "login" { return "Login" }
    return ""
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["login"] = authentication.Login
    return leafs
}

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetYangName() string { return "authentication" }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Tty_TtyLines_TtyLine_Aaa_Authentication) GetParentYangName() string { return "aaa" }

// Tty_TtyLines_TtyLine_Aaa_Accounting
// Accounting parameters
type Tty_TtyLines_TtyLine_Aaa_Accounting struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // For starting an exec (shell). The type is string.
    Exec interface{}

    // For exec (shell) configuration. The type is string.
    Commands interface{}
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetFilter() yfilter.YFilter { return accounting.YFilter }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) SetFilter(yf yfilter.YFilter) { accounting.YFilter = yf }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetGoName(yname string) string {
    if yname == "exec" { return "Exec" }
    if yname == "commands" { return "Commands" }
    return ""
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetSegmentPath() string {
    return "accounting"
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["exec"] = accounting.Exec
    leafs["commands"] = accounting.Commands
    return leafs
}

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetBundleName() string { return "cisco_ios_xr" }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetYangName() string { return "accounting" }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) SetParent(parent types.Entity) { accounting.parent = parent }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetParent() types.Entity { return accounting.parent }

func (accounting *Tty_TtyLines_TtyLine_Aaa_Accounting) GetParentYangName() string { return "aaa" }

// Tty_TtyLines_TtyLine_Exec
// EXEC timeout and timestamp configurtion
type Tty_TtyLines_TtyLine_Exec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // 'True' to Enable & 'False' to Disable time stamp. The type is bool.
    TimeStamp interface{}

    // EXEC Timeout.
    Timeout Tty_TtyLines_TtyLine_Exec_Timeout
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetFilter() yfilter.YFilter { return exec.YFilter }

func (exec *Tty_TtyLines_TtyLine_Exec) SetFilter(yf yfilter.YFilter) { exec.YFilter = yf }

func (exec *Tty_TtyLines_TtyLine_Exec) GetGoName(yname string) string {
    if yname == "time-stamp" { return "TimeStamp" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetSegmentPath() string {
    return "exec"
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timeout" {
        return &exec.Timeout
    }
    return nil
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timeout"] = &exec.Timeout
    return children
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-stamp"] = exec.TimeStamp
    return leafs
}

func (exec *Tty_TtyLines_TtyLine_Exec) GetBundleName() string { return "cisco_ios_xr" }

func (exec *Tty_TtyLines_TtyLine_Exec) GetYangName() string { return "exec" }

func (exec *Tty_TtyLines_TtyLine_Exec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exec *Tty_TtyLines_TtyLine_Exec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exec *Tty_TtyLines_TtyLine_Exec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exec *Tty_TtyLines_TtyLine_Exec) SetParent(parent types.Entity) { exec.parent = parent }

func (exec *Tty_TtyLines_TtyLine_Exec) GetParent() types.Entity { return exec.parent }

func (exec *Tty_TtyLines_TtyLine_Exec) GetParentYangName() string { return "tty-line" }

// Tty_TtyLines_TtyLine_Exec_Timeout
// EXEC Timeout
// This type is a presence type.
type Tty_TtyLines_TtyLine_Exec_Timeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 0..35791. This
    // attribute is mandatory. Units are minute.
    Minutes interface{}

    // Timeout in seconds. The type is interface{} with range: 0..2147483. This
    // attribute is mandatory. Units are second.
    Seconds interface{}
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetFilter() yfilter.YFilter { return timeout.YFilter }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) SetFilter(yf yfilter.YFilter) { timeout.YFilter = yf }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetGoName(yname string) string {
    if yname == "minutes" { return "Minutes" }
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetSegmentPath() string {
    return "timeout"
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minutes"] = timeout.Minutes
    leafs["seconds"] = timeout.Seconds
    return leafs
}

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetBundleName() string { return "cisco_ios_xr" }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetYangName() string { return "timeout" }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) SetParent(parent types.Entity) { timeout.parent = parent }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetParent() types.Entity { return timeout.parent }

func (timeout *Tty_TtyLines_TtyLine_Exec_Timeout) GetParentYangName() string { return "exec" }

// Tty_TtyLines_TtyLine_Connection
// Management connection configuration
type Tty_TtyLines_TtyLine_Connection struct {
    parent types.Entity
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
    // ((\p{IsBasicLatin}|\p{IsLatin-1Supplement})*)|(DEFAULT)|(BREAK)|(NONE) The
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

func (connection *Tty_TtyLines_TtyLine_Connection) GetFilter() yfilter.YFilter { return connection.YFilter }

func (connection *Tty_TtyLines_TtyLine_Connection) SetFilter(yf yfilter.YFilter) { connection.YFilter = yf }

func (connection *Tty_TtyLines_TtyLine_Connection) GetGoName(yname string) string {
    if yname == "disconnect-character" { return "DisconnectCharacter" }
    if yname == "acl-in" { return "AclIn" }
    if yname == "acl-out" { return "AclOut" }
    if yname == "cli-white-space-completion" { return "CliWhiteSpaceCompletion" }
    if yname == "session-limit" { return "SessionLimit" }
    if yname == "escape-character" { return "EscapeCharacter" }
    if yname == "transport-preferred" { return "TransportPreferred" }
    if yname == "transport-input" { return "TransportInput" }
    if yname == "transport-output" { return "TransportOutput" }
    if yname == "session-timeout" { return "SessionTimeout" }
    return ""
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-management-cfg:connection"
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transport-input" {
        return &connection.TransportInput
    }
    if childYangName == "transport-output" {
        return &connection.TransportOutput
    }
    if childYangName == "session-timeout" {
        return &connection.SessionTimeout
    }
    return nil
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transport-input"] = &connection.TransportInput
    children["transport-output"] = &connection.TransportOutput
    children["session-timeout"] = &connection.SessionTimeout
    return children
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disconnect-character"] = connection.DisconnectCharacter
    leafs["acl-in"] = connection.AclIn
    leafs["acl-out"] = connection.AclOut
    leafs["cli-white-space-completion"] = connection.CliWhiteSpaceCompletion
    leafs["session-limit"] = connection.SessionLimit
    leafs["escape-character"] = connection.EscapeCharacter
    leafs["transport-preferred"] = connection.TransportPreferred
    return leafs
}

func (connection *Tty_TtyLines_TtyLine_Connection) GetBundleName() string { return "cisco_ios_xr" }

func (connection *Tty_TtyLines_TtyLine_Connection) GetYangName() string { return "connection" }

func (connection *Tty_TtyLines_TtyLine_Connection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connection *Tty_TtyLines_TtyLine_Connection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connection *Tty_TtyLines_TtyLine_Connection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connection *Tty_TtyLines_TtyLine_Connection) SetParent(parent types.Entity) { connection.parent = parent }

func (connection *Tty_TtyLines_TtyLine_Connection) GetParent() types.Entity { return connection.parent }

func (connection *Tty_TtyLines_TtyLine_Connection) GetParentYangName() string { return "tty-line" }

// Tty_TtyLines_TtyLine_Connection_TransportInput
// Protocols to use when connecting to the
// terminal server
type Tty_TtyLines_TtyLine_Connection_TransportInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. The
    // default value is all.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: -2147483648..2147483647.
    None interface{}
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetFilter() yfilter.YFilter { return transportInput.YFilter }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) SetFilter(yf yfilter.YFilter) { transportInput.YFilter = yf }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    if yname == "protocol1" { return "Protocol1" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "none" { return "None" }
    return ""
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetSegmentPath() string {
    return "transport-input"
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["select"] = transportInput.Select
    leafs["protocol1"] = transportInput.Protocol1
    leafs["protocol2"] = transportInput.Protocol2
    leafs["none"] = transportInput.None
    return leafs
}

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetBundleName() string { return "cisco_ios_xr" }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetYangName() string { return "transport-input" }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) SetParent(parent types.Entity) { transportInput.parent = parent }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetParent() types.Entity { return transportInput.parent }

func (transportInput *Tty_TtyLines_TtyLine_Connection_TransportInput) GetParentYangName() string { return "connection" }

// Tty_TtyLines_TtyLine_Connection_TransportOutput
// Protocols to use for outgoing connections
// This type is a presence type.
type Tty_TtyLines_TtyLine_Connection_TransportOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Choose transport protocols. The type is TtyTransportProtocolSelect. This
    // attribute is mandatory.
    Select interface{}

    // Transport protocol1. The type is TtyTransportProtocol.
    Protocol1 interface{}

    // Transport protocol2. The type is TtyTransportProtocol.
    Protocol2 interface{}

    // Not used. The type is interface{} with range: -2147483648..2147483647.
    None interface{}
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetFilter() yfilter.YFilter { return transportOutput.YFilter }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) SetFilter(yf yfilter.YFilter) { transportOutput.YFilter = yf }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetGoName(yname string) string {
    if yname == "select" { return "Select" }
    if yname == "protocol1" { return "Protocol1" }
    if yname == "protocol2" { return "Protocol2" }
    if yname == "none" { return "None" }
    return ""
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetSegmentPath() string {
    return "transport-output"
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["select"] = transportOutput.Select
    leafs["protocol1"] = transportOutput.Protocol1
    leafs["protocol2"] = transportOutput.Protocol2
    leafs["none"] = transportOutput.None
    return leafs
}

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetBundleName() string { return "cisco_ios_xr" }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetYangName() string { return "transport-output" }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) SetParent(parent types.Entity) { transportOutput.parent = parent }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetParent() types.Entity { return transportOutput.parent }

func (transportOutput *Tty_TtyLines_TtyLine_Connection_TransportOutput) GetParentYangName() string { return "connection" }

// Tty_TtyLines_TtyLine_Connection_SessionTimeout
// Interval for closing connection when there is
// no input traffic
// This type is a presence type.
type Tty_TtyLines_TtyLine_Connection_SessionTimeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session timeout interval in minutes. The type is interface{} with range:
    // 0..35791. This attribute is mandatory.
    Timeout interface{}

    // Include output traffic as well as input traffic. The type is
    // TtySessionTimeoutDirection. This attribute is mandatory.
    Direction interface{}
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetFilter() yfilter.YFilter { return sessionTimeout.YFilter }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) SetFilter(yf yfilter.YFilter) { sessionTimeout.YFilter = yf }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    if yname == "direction" { return "Direction" }
    return ""
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetSegmentPath() string {
    return "session-timeout"
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = sessionTimeout.Timeout
    leafs["direction"] = sessionTimeout.Direction
    return leafs
}

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetYangName() string { return "session-timeout" }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) SetParent(parent types.Entity) { sessionTimeout.parent = parent }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetParent() types.Entity { return sessionTimeout.parent }

func (sessionTimeout *Tty_TtyLines_TtyLine_Connection_SessionTimeout) GetParentYangName() string { return "connection" }

// Tty_TtyLines_TtyLine_ExecMode
// Exec Mode Pager  configurtion
type Tty_TtyLines_TtyLine_ExecMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Preferred Paging Utility. The type is TtyPager. The default value is more.
    Pager interface{}
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetFilter() yfilter.YFilter { return execMode.YFilter }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) SetFilter(yf yfilter.YFilter) { execMode.YFilter = yf }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetGoName(yname string) string {
    if yname == "pager" { return "Pager" }
    return ""
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetSegmentPath() string {
    return "Cisco-IOS-XR-tty-management-cfg:exec-mode"
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pager"] = execMode.Pager
    return leafs
}

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetBundleName() string { return "cisco_ios_xr" }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetYangName() string { return "exec-mode" }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) SetParent(parent types.Entity) { execMode.parent = parent }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetParent() types.Entity { return execMode.parent }

func (execMode *Tty_TtyLines_TtyLine_ExecMode) GetParentYangName() string { return "tty-line" }

