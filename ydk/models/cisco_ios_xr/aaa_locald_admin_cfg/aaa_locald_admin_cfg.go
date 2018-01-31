// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-locald package admin-plane configuration.
// 
// This module contains definitions
// for the following management objects:
//   aaa: Admin plane AAA configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_locald_admin_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_locald_admin_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-locald-admin-cfg aaa}", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-locald-admin-cfg:aaa", reflect.TypeOf(Aaa{}))
}

// AaaAdminPassword represents Aaa admin password
type AaaAdminPassword string

const (
    // Type 5 password
    AaaAdminPassword_type5 AaaAdminPassword = "type5"

    // Type 8 password
    AaaAdminPassword_type8 AaaAdminPassword = "type8"

    // Type 9 password
    AaaAdminPassword_type9 AaaAdminPassword = "type9"
)

// Aaa
// Admin plane AAA configuration
type Aaa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure local username.
    Usernames Aaa_Usernames
}

func (aaa *Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Aaa) GetGoName(yname string) string {
    if yname == "usernames" { return "Usernames" }
    return ""
}

func (aaa *Aaa) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa"
}

func (aaa *Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usernames" {
        return &aaa.Usernames
    }
    return nil
}

func (aaa *Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["usernames"] = &aaa.Usernames
    return children
}

func (aaa *Aaa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaa *Aaa) GetBundleName() string { return "cisco_ios_xr" }

func (aaa *Aaa) GetYangName() string { return "aaa" }

func (aaa *Aaa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaa *Aaa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaa *Aaa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaa *Aaa) SetParent(parent types.Entity) { aaa.parent = parent }

func (aaa *Aaa) GetParent() types.Entity { return aaa.parent }

func (aaa *Aaa) GetParentYangName() string { return "Cisco-IOS-XR-aaa-locald-admin-cfg" }

// Aaa_Usernames
// Configure local username
type Aaa_Usernames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin Username. The type is slice of Aaa_Usernames_Username.
    Username []Aaa_Usernames_Username
}

func (usernames *Aaa_Usernames) GetFilter() yfilter.YFilter { return usernames.YFilter }

func (usernames *Aaa_Usernames) SetFilter(yf yfilter.YFilter) { usernames.YFilter = yf }

func (usernames *Aaa_Usernames) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    return ""
}

func (usernames *Aaa_Usernames) GetSegmentPath() string {
    return "usernames"
}

func (usernames *Aaa_Usernames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "username" {
        for _, c := range usernames.Username {
            if usernames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usernames_Username{}
        usernames.Username = append(usernames.Username, child)
        return &usernames.Username[len(usernames.Username)-1]
    }
    return nil
}

func (usernames *Aaa_Usernames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usernames.Username {
        children[usernames.Username[i].GetSegmentPath()] = &usernames.Username[i]
    }
    return children
}

func (usernames *Aaa_Usernames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usernames *Aaa_Usernames) GetBundleName() string { return "cisco_ios_xr" }

func (usernames *Aaa_Usernames) GetYangName() string { return "usernames" }

func (usernames *Aaa_Usernames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usernames *Aaa_Usernames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usernames *Aaa_Usernames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usernames *Aaa_Usernames) SetParent(parent types.Entity) { usernames.parent = parent }

func (usernames *Aaa_Usernames) GetParent() types.Entity { return usernames.parent }

func (usernames *Aaa_Usernames) GetParentYangName() string { return "aaa" }

// Aaa_Usernames_Username
// Admin Username
type Aaa_Usernames_Username struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Username. The type is string.
    Name interface{}

    // Specify the usergroup to which this admin user belongs.
    UsergroupUnderUsernames Aaa_Usernames_Username_UsergroupUnderUsernames

    // Specify the secret for the admin user.
    Secret Aaa_Usernames_Username_Secret
}

func (username *Aaa_Usernames_Username) GetFilter() yfilter.YFilter { return username.YFilter }

func (username *Aaa_Usernames_Username) SetFilter(yf yfilter.YFilter) { username.YFilter = yf }

func (username *Aaa_Usernames_Username) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "usergroup-under-usernames" { return "UsergroupUnderUsernames" }
    if yname == "secret" { return "Secret" }
    return ""
}

func (username *Aaa_Usernames_Username) GetSegmentPath() string {
    return "username" + "[name='" + fmt.Sprintf("%v", username.Name) + "']"
}

func (username *Aaa_Usernames_Username) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup-under-usernames" {
        return &username.UsergroupUnderUsernames
    }
    if childYangName == "secret" {
        return &username.Secret
    }
    return nil
}

func (username *Aaa_Usernames_Username) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["usergroup-under-usernames"] = &username.UsergroupUnderUsernames
    children["secret"] = &username.Secret
    return children
}

func (username *Aaa_Usernames_Username) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = username.Name
    return leafs
}

func (username *Aaa_Usernames_Username) GetBundleName() string { return "cisco_ios_xr" }

func (username *Aaa_Usernames_Username) GetYangName() string { return "username" }

func (username *Aaa_Usernames_Username) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (username *Aaa_Usernames_Username) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (username *Aaa_Usernames_Username) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (username *Aaa_Usernames_Username) SetParent(parent types.Entity) { username.parent = parent }

func (username *Aaa_Usernames_Username) GetParent() types.Entity { return username.parent }

func (username *Aaa_Usernames_Username) GetParentYangName() string { return "usernames" }

// Aaa_Usernames_Username_UsergroupUnderUsernames
// Specify the usergroup to which this admin user
// belongs
type Aaa_Usernames_Username_UsergroupUnderUsernames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is slice of
    // Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername.
    UsergroupUnderUsername []Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetFilter() yfilter.YFilter { return usergroupUnderUsernames.YFilter }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) SetFilter(yf yfilter.YFilter) { usergroupUnderUsernames.YFilter = yf }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetGoName(yname string) string {
    if yname == "usergroup-under-username" { return "UsergroupUnderUsername" }
    return ""
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetSegmentPath() string {
    return "usergroup-under-usernames"
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup-under-username" {
        for _, c := range usergroupUnderUsernames.UsergroupUnderUsername {
            if usergroupUnderUsernames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername{}
        usergroupUnderUsernames.UsergroupUnderUsername = append(usergroupUnderUsernames.UsergroupUnderUsername, child)
        return &usergroupUnderUsernames.UsergroupUnderUsername[len(usergroupUnderUsernames.UsergroupUnderUsername)-1]
    }
    return nil
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usergroupUnderUsernames.UsergroupUnderUsername {
        children[usergroupUnderUsernames.UsergroupUnderUsername[i].GetSegmentPath()] = &usergroupUnderUsernames.UsergroupUnderUsername[i]
    }
    return children
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetYangName() string { return "usergroup-under-usernames" }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) SetParent(parent types.Entity) { usergroupUnderUsernames.parent = parent }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetParent() types.Entity { return usergroupUnderUsernames.parent }

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetParentYangName() string { return "username" }

// Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
// Name of the usergroup
type Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the usergroup. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetFilter() yfilter.YFilter { return usergroupUnderUsername.YFilter }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) SetFilter(yf yfilter.YFilter) { usergroupUnderUsername.YFilter = yf }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetSegmentPath() string {
    return "usergroup-under-username" + "[name='" + fmt.Sprintf("%v", usergroupUnderUsername.Name) + "']"
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = usergroupUnderUsername.Name
    return leafs
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetBundleName() string { return "cisco_ios_xr" }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetYangName() string { return "usergroup-under-username" }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) SetParent(parent types.Entity) { usergroupUnderUsername.parent = parent }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetParent() types.Entity { return usergroupUnderUsername.parent }

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetParentYangName() string { return "usergroup-under-usernames" }

// Aaa_Usernames_Username_Secret
// Specify the secret for the admin user
type Aaa_Usernames_Username_Secret struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Password type. The type is AaaAdminPassword.
    Type interface{}

    // The user's secret password. The type is string with pattern:
    // (!.+)|([^!].+).
    Secret5 interface{}

    // Type 8 password. The type is string with pattern: (!.+)|([^!].+).
    Secret8 interface{}

    // Type 9 password. The type is string with pattern: (!.+)|([^!].+).
    Secret9 interface{}
}

func (secret *Aaa_Usernames_Username_Secret) GetFilter() yfilter.YFilter { return secret.YFilter }

func (secret *Aaa_Usernames_Username_Secret) SetFilter(yf yfilter.YFilter) { secret.YFilter = yf }

func (secret *Aaa_Usernames_Username_Secret) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "secret5" { return "Secret5" }
    if yname == "secret8" { return "Secret8" }
    if yname == "secret9" { return "Secret9" }
    return ""
}

func (secret *Aaa_Usernames_Username_Secret) GetSegmentPath() string {
    return "secret"
}

func (secret *Aaa_Usernames_Username_Secret) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secret *Aaa_Usernames_Username_Secret) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secret *Aaa_Usernames_Username_Secret) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = secret.Type
    leafs["secret5"] = secret.Secret5
    leafs["secret8"] = secret.Secret8
    leafs["secret9"] = secret.Secret9
    return leafs
}

func (secret *Aaa_Usernames_Username_Secret) GetBundleName() string { return "cisco_ios_xr" }

func (secret *Aaa_Usernames_Username_Secret) GetYangName() string { return "secret" }

func (secret *Aaa_Usernames_Username_Secret) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secret *Aaa_Usernames_Username_Secret) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secret *Aaa_Usernames_Username_Secret) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secret *Aaa_Usernames_Username_Secret) SetParent(parent types.Entity) { secret.parent = parent }

func (secret *Aaa_Usernames_Username_Secret) GetParent() types.Entity { return secret.parent }

func (secret *Aaa_Usernames_Username_Secret) GetParentYangName() string { return "username" }

