// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-locald package admin-plane configuration.
// 
// This module contains definitions
// for the following management objects:
//   aaa: Admin plane AAA configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    aaa.EntityData.AbsolutePath = aaa.EntityData.SegmentPath
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("usernames", types.YChild{"Usernames", &aaa.Usernames})
    aaa.EntityData.Leafs = types.NewOrderedMap()

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Aaa_Usernames
// Configure local username
type Aaa_Usernames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Admin Username. The type is slice of Aaa_Usernames_Username.
    Username []*Aaa_Usernames_Username
}

func (usernames *Aaa_Usernames) GetEntityData() *types.CommonEntityData {
    usernames.EntityData.YFilter = usernames.YFilter
    usernames.EntityData.YangName = "usernames"
    usernames.EntityData.BundleName = "cisco_ios_xr"
    usernames.EntityData.ParentYangName = "aaa"
    usernames.EntityData.SegmentPath = "usernames"
    usernames.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa/" + usernames.EntityData.SegmentPath
    usernames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernames.EntityData.Children = types.NewOrderedMap()
    usernames.EntityData.Children.Append("username", types.YChild{"Username", nil})
    for i := range usernames.Username {
        usernames.EntityData.Children.Append(types.GetSegmentPath(usernames.Username[i]), types.YChild{"Username", usernames.Username[i]})
    }
    usernames.EntityData.Leafs = types.NewOrderedMap()

    usernames.EntityData.YListKeys = []string {}

    return &(usernames.EntityData)
}

// Aaa_Usernames_Username
// Admin Username
type Aaa_Usernames_Username struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    username.EntityData.SegmentPath = "username" + types.AddKeyToken(username.Name, "name")
    username.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa/usernames/" + username.EntityData.SegmentPath
    username.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    username.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    username.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    username.EntityData.Children = types.NewOrderedMap()
    username.EntityData.Children.Append("usergroup-under-usernames", types.YChild{"UsergroupUnderUsernames", &username.UsergroupUnderUsernames})
    username.EntityData.Children.Append("secret", types.YChild{"Secret", &username.Secret})
    username.EntityData.Leafs = types.NewOrderedMap()
    username.EntityData.Leafs.Append("name", types.YLeaf{"Name", username.Name})

    username.EntityData.YListKeys = []string {"Name"}

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
    UsergroupUnderUsername []*Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
}

func (usergroupUnderUsernames *Aaa_Usernames_Username_UsergroupUnderUsernames) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsernames.EntityData.YFilter = usergroupUnderUsernames.YFilter
    usergroupUnderUsernames.EntityData.YangName = "usergroup-under-usernames"
    usergroupUnderUsernames.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsernames.EntityData.ParentYangName = "username"
    usergroupUnderUsernames.EntityData.SegmentPath = "usergroup-under-usernames"
    usergroupUnderUsernames.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa/usernames/username/" + usergroupUnderUsernames.EntityData.SegmentPath
    usergroupUnderUsernames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsernames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsernames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsernames.EntityData.Children = types.NewOrderedMap()
    usergroupUnderUsernames.EntityData.Children.Append("usergroup-under-username", types.YChild{"UsergroupUnderUsername", nil})
    for i := range usergroupUnderUsernames.UsergroupUnderUsername {
        usergroupUnderUsernames.EntityData.Children.Append(types.GetSegmentPath(usergroupUnderUsernames.UsergroupUnderUsername[i]), types.YChild{"UsergroupUnderUsername", usergroupUnderUsernames.UsergroupUnderUsername[i]})
    }
    usergroupUnderUsernames.EntityData.Leafs = types.NewOrderedMap()

    usergroupUnderUsernames.EntityData.YListKeys = []string {}

    return &(usergroupUnderUsernames.EntityData)
}

// Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername
// Name of the usergroup
type Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the usergroup. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsername.EntityData.YFilter = usergroupUnderUsername.YFilter
    usergroupUnderUsername.EntityData.YangName = "usergroup-under-username"
    usergroupUnderUsername.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsername.EntityData.ParentYangName = "usergroup-under-usernames"
    usergroupUnderUsername.EntityData.SegmentPath = "usergroup-under-username" + types.AddKeyToken(usergroupUnderUsername.Name, "name")
    usergroupUnderUsername.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa/usernames/username/usergroup-under-usernames/" + usergroupUnderUsername.EntityData.SegmentPath
    usergroupUnderUsername.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsername.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsername.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsername.EntityData.Children = types.NewOrderedMap()
    usergroupUnderUsername.EntityData.Leafs = types.NewOrderedMap()
    usergroupUnderUsername.EntityData.Leafs.Append("name", types.YLeaf{"Name", usergroupUnderUsername.Name})

    usergroupUnderUsername.EntityData.YListKeys = []string {"Name"}

    return &(usergroupUnderUsername.EntityData)
}

// Aaa_Usernames_Username_Secret
// Specify the secret for the admin user
type Aaa_Usernames_Username_Secret struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password type. The type is AaaAdminPassword.
    Type interface{}

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
    secret.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-locald-admin-cfg:aaa/usernames/username/" + secret.EntityData.SegmentPath
    secret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secret.EntityData.Children = types.NewOrderedMap()
    secret.EntityData.Leafs = types.NewOrderedMap()
    secret.EntityData.Leafs.Append("type", types.YLeaf{"Type", secret.Type})
    secret.EntityData.Leafs.Append("secret5", types.YLeaf{"Secret5", secret.Secret5})
    secret.EntityData.Leafs.Append("secret8", types.YLeaf{"Secret8", secret.Secret8})
    secret.EntityData.Leafs.Append("secret9", types.YLeaf{"Secret9", secret.Secret9})

    secret.EntityData.YListKeys = []string {}

    return &(secret.EntityData)
}

