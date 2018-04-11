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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure local username.
    Usernames Aaa_Usernames
}

func (aaa *Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-locald-admin-cfg"
    aaa.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = make(map[string]types.YChild)
    aaa.EntityData.Children["usernames"] = types.YChild{"Usernames", &aaa.Usernames}
    aaa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aaa.EntityData)
}

// Aaa_Usernames
// Configure local username
type Aaa_Usernames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin Username. The type is slice of Aaa_Usernames_Username.
    Username []Aaa_Usernames_Username
}

func (usernames *Aaa_Usernames) GetEntityData() *types.CommonEntityData {
    usernames.EntityData.YFilter = usernames.YFilter
    usernames.EntityData.YangName = "usernames"
    usernames.EntityData.BundleName = "cisco_ios_xr"
    usernames.EntityData.ParentYangName = "aaa"
    usernames.EntityData.SegmentPath = "usernames"
    usernames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernames.EntityData.Children = make(map[string]types.YChild)
    usernames.EntityData.Children["username"] = types.YChild{"Username", nil}
    for i := range usernames.Username {
        usernames.EntityData.Children[types.GetSegmentPath(&usernames.Username[i])] = types.YChild{"Username", &usernames.Username[i]}
    }
    usernames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usernames.EntityData)
}

// Aaa_Usernames_Username
// Admin Username
type Aaa_Usernames_Username struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Username. The type is string.
    Name interface{}

    // Specify the usergroup to which this admin user belongs.
    UsergroupUnderUsernames Aaa_Usernames_Username_UsergroupUnderUsernames

    // Specify the secret for the admin user.
    Secret Aaa_Usernames_Username_Secret
}

func (username *Aaa_Usernames_Username) GetEntityData() *types.CommonEntityData {
    username.EntityData.YFilter = username.YFilter
    username.EntityData.YangName = "username"
    username.EntityData.BundleName = "cisco_ios_xr"
    username.EntityData.ParentYangName = "usernames"
    username.EntityData.SegmentPath = "username" + "[name='" + fmt.Sprintf("%v", username.Name) + "']"
    username.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    username.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    username.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    username.EntityData.Children = make(map[string]types.YChild)
    username.EntityData.Children["usergroup-under-usernames"] = types.YChild{"UsergroupUnderUsernames", &username.UsergroupUnderUsernames}
    username.EntityData.Children["secret"] = types.YChild{"Secret", &username.Secret}
    username.EntityData.Leafs = make(map[string]types.YLeaf)
    username.EntityData.Leafs["name"] = types.YLeaf{"Name", username.Name}
    return &(username.EntityData)
}

// Aaa_Usernames_Username_UsergroupUnderUsernames
// Specify the usergroup to which this admin user
// belongs
type Aaa_Usernames_Username_UsergroupUnderUsernames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is slice of
    // Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername.
    UsergroupUnderUsername []Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsernames.EntityData.YFilter = usergroupUnderUsernames.YFilter
    usergroupUnderUsernames.EntityData.YangName = "usergroup-under-usernames"
    usergroupUnderUsernames.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsernames.EntityData.ParentYangName = "username"
    usergroupUnderUsernames.EntityData.SegmentPath = "usergroup-under-usernames"
    usergroupUnderUsernames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsernames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsernames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsernames.EntityData.Children = make(map[string]types.YChild)
    usergroupUnderUsernames.EntityData.Children["usergroup-under-username"] = types.YChild{"UsergroupUnderUsername", nil}
    for i := range usergroupUnderUsernames.UsergroupUnderUsername {
        usergroupUnderUsernames.EntityData.Children[types.GetSegmentPath(&usergroupUnderUsernames.UsergroupUnderUsername[i])] = types.YChild{"UsergroupUnderUsername", &usergroupUnderUsernames.UsergroupUnderUsername[i]}
    }
    usergroupUnderUsernames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usergroupUnderUsernames.EntityData)
}

// Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
// Name of the usergroup
type Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the usergroup. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsername.EntityData.YFilter = usergroupUnderUsername.YFilter
    usergroupUnderUsername.EntityData.YangName = "usergroup-under-username"
    usergroupUnderUsername.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsername.EntityData.ParentYangName = "usergroup-under-usernames"
    usergroupUnderUsername.EntityData.SegmentPath = "usergroup-under-username" + "[name='" + fmt.Sprintf("%v", usergroupUnderUsername.Name) + "']"
    usergroupUnderUsername.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsername.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsername.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsername.EntityData.Children = make(map[string]types.YChild)
    usergroupUnderUsername.EntityData.Leafs = make(map[string]types.YLeaf)
    usergroupUnderUsername.EntityData.Leafs["name"] = types.YLeaf{"Name", usergroupUnderUsername.Name}
    return &(usergroupUnderUsername.EntityData)
}

// Aaa_Usernames_Username_Secret
// Specify the secret for the admin user
type Aaa_Usernames_Username_Secret struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password type. The type is AaaAdminPassword.
    Type_ interface{}

    // The user's secret password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Secret5 interface{}

    // Type 8 password. The type is string with pattern: b'(!.+)|([^!].+)'.
    Secret8 interface{}

    // Type 9 password. The type is string with pattern: b'(!.+)|([^!].+)'.
    Secret9 interface{}
}

func (secret *Aaa_Usernames_Username_Secret) GetEntityData() *types.CommonEntityData {
    secret.EntityData.YFilter = secret.YFilter
    secret.EntityData.YangName = "secret"
    secret.EntityData.BundleName = "cisco_ios_xr"
    secret.EntityData.ParentYangName = "username"
    secret.EntityData.SegmentPath = "secret"
    secret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secret.EntityData.Children = make(map[string]types.YChild)
    secret.EntityData.Leafs = make(map[string]types.YLeaf)
    secret.EntityData.Leafs["type"] = types.YLeaf{"Type_", secret.Type_}
    secret.EntityData.Leafs["secret5"] = types.YLeaf{"Secret5", secret.Secret5}
    secret.EntityData.Leafs["secret8"] = types.YLeaf{"Secret8", secret.Secret8}
    secret.EntityData.Leafs["secret9"] = types.YLeaf{"Secret9", secret.Secret9}
    return &(secret.EntityData)
}

