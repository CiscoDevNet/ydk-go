// This module contains definitions
// for the Calvados model objects.
// 
// This module defines the Tail-f AAA data model.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package tailf_aaa

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tailf_aaa"))
    ydk.RegisterEntity("{http://tail-f.com/ns/aaa/1.1 aaa}", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("tailf-aaa:aaa", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/aaa/1.1 alias}", reflect.TypeOf(Alias{}))
    ydk.RegisterEntity("tailf-aaa:alias", reflect.TypeOf(Alias{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/aaa/1.1 session}", reflect.TypeOf(Session{}))
    ydk.RegisterEntity("tailf-aaa:session", reflect.TypeOf(Session{}))
    ydk.RegisterEntity("{http://tail-f.com/ns/aaa/1.1 user}", reflect.TypeOf(User{}))
    ydk.RegisterEntity("tailf-aaa:user", reflect.TypeOf(User{}))
}

// DataOperationType
type DataOperationType string

const (
    DataOperationType_r DataOperationType = "r"

    DataOperationType_rx DataOperationType = "rx"

    DataOperationType_x DataOperationType = "x"

    DataOperationType_rw DataOperationType = "rw"

    DataOperationType_rwx DataOperationType = "rwx"

    DataOperationType_wx DataOperationType = "wx"

    DataOperationType_w DataOperationType = "w"

    DataOperationType_c DataOperationType = "c"

    DataOperationType_cr DataOperationType = "cr"

    DataOperationType_cu DataOperationType = "cu"

    DataOperationType_cd DataOperationType = "cd"

    DataOperationType_cx DataOperationType = "cx"

    DataOperationType_cru DataOperationType = "cru"

    DataOperationType_crd DataOperationType = "crd"

    DataOperationType_crx DataOperationType = "crx"

    DataOperationType_cud DataOperationType = "cud"

    DataOperationType_cux DataOperationType = "cux"

    DataOperationType_cdx DataOperationType = "cdx"

    DataOperationType_crud DataOperationType = "crud"

    DataOperationType_crux DataOperationType = "crux"

    DataOperationType_crdx DataOperationType = "crdx"

    DataOperationType_cudx DataOperationType = "cudx"

    DataOperationType_crudx DataOperationType = "crudx"

    DataOperationType_ru DataOperationType = "ru"

    DataOperationType_rd DataOperationType = "rd"

    DataOperationType_rud DataOperationType = "rud"

    DataOperationType_rux DataOperationType = "rux"

    DataOperationType_rdx DataOperationType = "rdx"

    DataOperationType_u DataOperationType = "u"

    DataOperationType_ud DataOperationType = "ud"

    DataOperationType_ux DataOperationType = "ux"

    DataOperationType_d DataOperationType = "d"

    DataOperationType_dx DataOperationType = "dx"
)

// CmdOperationType
type CmdOperationType string

const (
    CmdOperationType_r CmdOperationType = "r"

    CmdOperationType_rx CmdOperationType = "rx"

    CmdOperationType_x CmdOperationType = "x"
)

// Action
type Action string

const (
    Action_accept Action = "accept"

    Action_reject Action = "reject"

    Action_accept_log Action = "accept_log"
)

// BuiltinModes
type BuiltinModes string

const (
    BuiltinModes_exec BuiltinModes = "exec"

    BuiltinModes_configure BuiltinModes = "configure"
)

// BuiltinModes_
type BuiltinModes_ string

const (
    BuiltinModes__exec BuiltinModes_ = "exec"

    BuiltinModes__configure BuiltinModes_ = "configure"
)

// Aaa
type Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Authentication Aaa_Authentication

    
    Authorization Aaa_Authorization

    
    Accounting Aaa_Accounting

    
    Ios Aaa_Ios

    
    PrivilegedAccess Aaa_PrivilegedAccess

    
    CiscoIOSXRSysadminAaaAaaShowAccounting Aaa_CiscoIOSXRSysadminAaaAaaShowAccounting

    
    UserGroup Aaa_UserGroup

    
    DisasterRecovery Aaa_DisasterRecovery
}

func (aaa *Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "tailf-aaa"
    aaa.EntityData.SegmentPath = "tailf-aaa:aaa"
    aaa.EntityData.AbsolutePath = aaa.EntityData.SegmentPath
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("authentication", types.YChild{"Authentication", &aaa.Authentication})
    aaa.EntityData.Children.Append("authorization", types.YChild{"Authorization", &aaa.Authorization})
    aaa.EntityData.Children.Append("accounting", types.YChild{"Accounting", &aaa.Accounting})
    aaa.EntityData.Children.Append("ios", types.YChild{"Ios", &aaa.Ios})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-aaa-aaa-show:privileged-access", types.YChild{"PrivilegedAccess", &aaa.PrivilegedAccess})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-aaa-aaa-show:accounting", types.YChild{"CiscoIOSXRSysadminAaaAaaShowAccounting", &aaa.CiscoIOSXRSysadminAaaAaaShowAccounting})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-aaa-aaa-show:user-group", types.YChild{"UserGroup", &aaa.UserGroup})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-sysadmin-aaa-disaster-recovery:disaster-recovery", types.YChild{"DisasterRecovery", &aaa.DisasterRecovery})
    aaa.EntityData.Leafs = types.NewOrderedMap()

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Aaa_Authentication
type Aaa_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Users Aaa_Authentication_Users

    
    Groups Aaa_Authentication_Groups

    
    Login Aaa_Authentication_Login
}

func (authentication *Aaa_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "aaa"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "tailf-aaa:aaa/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Children.Append("users", types.YChild{"Users", &authentication.Users})
    authentication.EntityData.Children.Append("groups", types.YChild{"Groups", &authentication.Groups})
    authentication.EntityData.Children.Append("login", types.YChild{"Login", &authentication.Login})
    authentication.EntityData.Leafs = types.NewOrderedMap()

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Aaa_Authentication_Users
type Aaa_Authentication_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authentication_Users_User.
    User []*Aaa_Authentication_Users_User
}

func (users *Aaa_Authentication_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "authentication"
    users.EntityData.SegmentPath = "users"
    users.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/" + users.EntityData.SegmentPath
    users.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    users.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    users.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    users.EntityData.Children = types.NewOrderedMap()
    users.EntityData.Children.Append("user", types.YChild{"User", nil})
    for i := range users.User {
        users.EntityData.Children.Append(types.GetSegmentPath(users.User[i]), types.YChild{"User", users.User[i]})
    }
    users.EntityData.Leafs = types.NewOrderedMap()

    users.EntityData.YListKeys = []string {}

    return &(users.EntityData)
}

// Aaa_Authentication_Users_User
type Aaa_Authentication_Users_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295. The default value is
    // 9000.
    Uid interface{}

    // The type is interface{} with range: 0..4294967295. The default value is
    // 100.
    Gid interface{}

    // The type is string. This attribute is mandatory.
    Password interface{}

    // The type is string. The default value is /home/sshdir.
    SshKeydir interface{}

    // The type is string. The default value is /home/homedir.
    Homedir interface{}
}

func (user *Aaa_Authentication_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + types.AddKeyToken(user.Name, "name")
    user.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/users/" + user.EntityData.SegmentPath
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Leafs = types.NewOrderedMap()
    user.EntityData.Leafs.Append("name", types.YLeaf{"Name", user.Name})
    user.EntityData.Leafs.Append("uid", types.YLeaf{"Uid", user.Uid})
    user.EntityData.Leafs.Append("gid", types.YLeaf{"Gid", user.Gid})
    user.EntityData.Leafs.Append("password", types.YLeaf{"Password", user.Password})
    user.EntityData.Leafs.Append("ssh_keydir", types.YLeaf{"SshKeydir", user.SshKeydir})
    user.EntityData.Leafs.Append("homedir", types.YLeaf{"Homedir", user.Homedir})

    user.EntityData.YListKeys = []string {"Name"}

    return &(user.EntityData)
}

// Aaa_Authentication_Groups
type Aaa_Authentication_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authentication_Groups_Group.
    Group []*Aaa_Authentication_Groups_Group
}

func (groups *Aaa_Authentication_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "authentication"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Aaa_Authentication_Groups_Group
type Aaa_Authentication_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: 0..4294967295. The default value is
    // 100.
    Gid interface{}

    // The type is string. The default value is %%__system_user__%%.
    Users interface{}
}

func (group *Aaa_Authentication_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Name, "name")
    group.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/groups/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("name", types.YLeaf{"Name", group.Name})
    group.EntityData.Leafs.Append("gid", types.YLeaf{"Gid", group.Gid})
    group.EntityData.Leafs.Append("users", types.YLeaf{"Users", group.Users})

    group.EntityData.YListKeys = []string {"Name"}

    return &(group.EntityData)
}

// Aaa_Authentication_Login
type Aaa_Authentication_Login struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Group Aaa_Authentication_Login_Group
}

func (login *Aaa_Authentication_Login) GetEntityData() *types.CommonEntityData {
    login.EntityData.YFilter = login.YFilter
    login.EntityData.YangName = "login"
    login.EntityData.BundleName = "cisco_ios_xr"
    login.EntityData.ParentYangName = "authentication"
    login.EntityData.SegmentPath = "login"
    login.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/" + login.EntityData.SegmentPath
    login.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    login.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    login.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    login.EntityData.Children = types.NewOrderedMap()
    login.EntityData.Children.Append("group", types.YChild{"Group", &login.Group})
    login.EntityData.Leafs = types.NewOrderedMap()

    login.EntityData.YListKeys = []string {}

    return &(login.EntityData)
}

// Aaa_Authentication_Login_Group
type Aaa_Authentication_Login_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Tacacs interface{}
}

func (group *Aaa_Authentication_Login_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "login"
    group.EntityData.SegmentPath = "group"
    group.EntityData.AbsolutePath = "tailf-aaa:aaa/authentication/login/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("tacacs", types.YLeaf{"Tacacs", group.Tacacs})

    group.EntityData.YListKeys = []string {}

    return &(group.EntityData)
}

// Aaa_Authorization
type Aaa_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cmdrules Aaa_Authorization_Cmdrules

    
    Datarules Aaa_Authorization_Datarules

    
    Commands Aaa_Authorization_Commands
}

func (authorization *Aaa_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "aaa"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.AbsolutePath = "tailf-aaa:aaa/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Children.Append("cmdrules", types.YChild{"Cmdrules", &authorization.Cmdrules})
    authorization.EntityData.Children.Append("datarules", types.YChild{"Datarules", &authorization.Datarules})
    authorization.EntityData.Children.Append("commands", types.YChild{"Commands", &authorization.Commands})
    authorization.EntityData.Leafs = types.NewOrderedMap()

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

// Aaa_Authorization_Cmdrules
type Aaa_Authorization_Cmdrules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authorization_Cmdrules_Cmdrule.
    Cmdrule []*Aaa_Authorization_Cmdrules_Cmdrule
}

func (cmdrules *Aaa_Authorization_Cmdrules) GetEntityData() *types.CommonEntityData {
    cmdrules.EntityData.YFilter = cmdrules.YFilter
    cmdrules.EntityData.YangName = "cmdrules"
    cmdrules.EntityData.BundleName = "cisco_ios_xr"
    cmdrules.EntityData.ParentYangName = "authorization"
    cmdrules.EntityData.SegmentPath = "cmdrules"
    cmdrules.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/" + cmdrules.EntityData.SegmentPath
    cmdrules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cmdrules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cmdrules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cmdrules.EntityData.Children = types.NewOrderedMap()
    cmdrules.EntityData.Children.Append("cmdrule", types.YChild{"Cmdrule", nil})
    for i := range cmdrules.Cmdrule {
        cmdrules.EntityData.Children.Append(types.GetSegmentPath(cmdrules.Cmdrule[i]), types.YChild{"Cmdrule", cmdrules.Cmdrule[i]})
    }
    cmdrules.EntityData.Leafs = types.NewOrderedMap()

    cmdrules.EntityData.YListKeys = []string {}

    return &(cmdrules.EntityData)
}

// Aaa_Authorization_Cmdrules_Cmdrule
type Aaa_Authorization_Cmdrules_Cmdrule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string. This attribute is mandatory.
    Context interface{}

    // The type is string. This attribute is mandatory.
    Command interface{}

    // The type is string. This attribute is mandatory.
    Group interface{}

    // The type is CmdOperationType. This attribute is mandatory.
    Ops interface{}

    // The type is Action. This attribute is mandatory.
    Action interface{}
}

func (cmdrule *Aaa_Authorization_Cmdrules_Cmdrule) GetEntityData() *types.CommonEntityData {
    cmdrule.EntityData.YFilter = cmdrule.YFilter
    cmdrule.EntityData.YangName = "cmdrule"
    cmdrule.EntityData.BundleName = "cisco_ios_xr"
    cmdrule.EntityData.ParentYangName = "cmdrules"
    cmdrule.EntityData.SegmentPath = "cmdrule" + types.AddKeyToken(cmdrule.Index, "index")
    cmdrule.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/cmdrules/" + cmdrule.EntityData.SegmentPath
    cmdrule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cmdrule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cmdrule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cmdrule.EntityData.Children = types.NewOrderedMap()
    cmdrule.EntityData.Leafs = types.NewOrderedMap()
    cmdrule.EntityData.Leafs.Append("index", types.YLeaf{"Index", cmdrule.Index})
    cmdrule.EntityData.Leafs.Append("context", types.YLeaf{"Context", cmdrule.Context})
    cmdrule.EntityData.Leafs.Append("command", types.YLeaf{"Command", cmdrule.Command})
    cmdrule.EntityData.Leafs.Append("group", types.YLeaf{"Group", cmdrule.Group})
    cmdrule.EntityData.Leafs.Append("ops", types.YLeaf{"Ops", cmdrule.Ops})
    cmdrule.EntityData.Leafs.Append("action", types.YLeaf{"Action", cmdrule.Action})

    cmdrule.EntityData.YListKeys = []string {"Index"}

    return &(cmdrule.EntityData)
}

// Aaa_Authorization_Datarules
type Aaa_Authorization_Datarules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authorization_Datarules_Datarule.
    Datarule []*Aaa_Authorization_Datarules_Datarule
}

func (datarules *Aaa_Authorization_Datarules) GetEntityData() *types.CommonEntityData {
    datarules.EntityData.YFilter = datarules.YFilter
    datarules.EntityData.YangName = "datarules"
    datarules.EntityData.BundleName = "cisco_ios_xr"
    datarules.EntityData.ParentYangName = "authorization"
    datarules.EntityData.SegmentPath = "datarules"
    datarules.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/" + datarules.EntityData.SegmentPath
    datarules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    datarules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    datarules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    datarules.EntityData.Children = types.NewOrderedMap()
    datarules.EntityData.Children.Append("datarule", types.YChild{"Datarule", nil})
    for i := range datarules.Datarule {
        datarules.EntityData.Children.Append(types.GetSegmentPath(datarules.Datarule[i]), types.YChild{"Datarule", datarules.Datarule[i]})
    }
    datarules.EntityData.Leafs = types.NewOrderedMap()

    datarules.EntityData.YListKeys = []string {}

    return &(datarules.EntityData)
}

// Aaa_Authorization_Datarules_Datarule
type Aaa_Authorization_Datarules_Datarule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string. This attribute is mandatory.
    Namespace interface{}

    // The type is string. The default value is *.
    Context interface{}

    // The type is string. This attribute is mandatory.
    Keypath interface{}

    // The type is string. This attribute is mandatory.
    Group interface{}

    // The type is DataOperationType. This attribute is mandatory.
    Ops interface{}

    // The type is Action. This attribute is mandatory.
    Action interface{}
}

func (datarule *Aaa_Authorization_Datarules_Datarule) GetEntityData() *types.CommonEntityData {
    datarule.EntityData.YFilter = datarule.YFilter
    datarule.EntityData.YangName = "datarule"
    datarule.EntityData.BundleName = "cisco_ios_xr"
    datarule.EntityData.ParentYangName = "datarules"
    datarule.EntityData.SegmentPath = "datarule" + types.AddKeyToken(datarule.Index, "index")
    datarule.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/datarules/" + datarule.EntityData.SegmentPath
    datarule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    datarule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    datarule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    datarule.EntityData.Children = types.NewOrderedMap()
    datarule.EntityData.Leafs = types.NewOrderedMap()
    datarule.EntityData.Leafs.Append("index", types.YLeaf{"Index", datarule.Index})
    datarule.EntityData.Leafs.Append("namespace", types.YLeaf{"Namespace", datarule.Namespace})
    datarule.EntityData.Leafs.Append("context", types.YLeaf{"Context", datarule.Context})
    datarule.EntityData.Leafs.Append("keypath", types.YLeaf{"Keypath", datarule.Keypath})
    datarule.EntityData.Leafs.Append("group", types.YLeaf{"Group", datarule.Group})
    datarule.EntityData.Leafs.Append("ops", types.YLeaf{"Ops", datarule.Ops})
    datarule.EntityData.Leafs.Append("action", types.YLeaf{"Action", datarule.Action})

    datarule.EntityData.YListKeys = []string {"Index"}

    return &(datarule.EntityData)
}

// Aaa_Authorization_Commands
type Aaa_Authorization_Commands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Group Aaa_Authorization_Commands_Group
}

func (commands *Aaa_Authorization_Commands) GetEntityData() *types.CommonEntityData {
    commands.EntityData.YFilter = commands.YFilter
    commands.EntityData.YangName = "commands"
    commands.EntityData.BundleName = "cisco_ios_xr"
    commands.EntityData.ParentYangName = "authorization"
    commands.EntityData.SegmentPath = "commands"
    commands.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/" + commands.EntityData.SegmentPath
    commands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commands.EntityData.Children = types.NewOrderedMap()
    commands.EntityData.Children.Append("group", types.YChild{"Group", &commands.Group})
    commands.EntityData.Leafs = types.NewOrderedMap()

    commands.EntityData.YListKeys = []string {}

    return &(commands.EntityData)
}

// Aaa_Authorization_Commands_Group
type Aaa_Authorization_Commands_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Tacacs interface{}

    // The type is interface{}.
    None interface{}
}

func (group *Aaa_Authorization_Commands_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "commands"
    group.EntityData.SegmentPath = "group"
    group.EntityData.AbsolutePath = "tailf-aaa:aaa/authorization/commands/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("tacacs", types.YLeaf{"Tacacs", group.Tacacs})
    group.EntityData.Leafs.Append("none", types.YLeaf{"None", group.None})

    group.EntityData.YListKeys = []string {}

    return &(group.EntityData)
}

// Aaa_Accounting
type Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Commands Aaa_Accounting_Commands
}

func (accounting *Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "tailf-aaa:aaa/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("commands", types.YChild{"Commands", &accounting.Commands})
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Aaa_Accounting_Commands
type Aaa_Accounting_Commands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Group Aaa_Accounting_Commands_Group
}

func (commands *Aaa_Accounting_Commands) GetEntityData() *types.CommonEntityData {
    commands.EntityData.YFilter = commands.YFilter
    commands.EntityData.YangName = "commands"
    commands.EntityData.BundleName = "cisco_ios_xr"
    commands.EntityData.ParentYangName = "accounting"
    commands.EntityData.SegmentPath = "commands"
    commands.EntityData.AbsolutePath = "tailf-aaa:aaa/accounting/" + commands.EntityData.SegmentPath
    commands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commands.EntityData.Children = types.NewOrderedMap()
    commands.EntityData.Children.Append("group", types.YChild{"Group", &commands.Group})
    commands.EntityData.Leafs = types.NewOrderedMap()

    commands.EntityData.YListKeys = []string {}

    return &(commands.EntityData)
}

// Aaa_Accounting_Commands_Group
type Aaa_Accounting_Commands_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{}.
    Tacacs interface{}
}

func (group *Aaa_Accounting_Commands_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "commands"
    group.EntityData.SegmentPath = "group"
    group.EntityData.AbsolutePath = "tailf-aaa:aaa/accounting/commands/" + group.EntityData.SegmentPath
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("tacacs", types.YLeaf{"Tacacs", group.Tacacs})

    group.EntityData.YListKeys = []string {}

    return &(group.EntityData)
}

// Aaa_Ios
// This type is a presence type.
type Aaa_Ios struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is slice of Aaa_Ios_Level.
    Level []*Aaa_Ios_Level

    // The type is slice of Aaa_Ios_Privilege.
    Privilege []*Aaa_Ios_Privilege
}

func (ios *Aaa_Ios) GetEntityData() *types.CommonEntityData {
    ios.EntityData.YFilter = ios.YFilter
    ios.EntityData.YangName = "ios"
    ios.EntityData.BundleName = "cisco_ios_xr"
    ios.EntityData.ParentYangName = "aaa"
    ios.EntityData.SegmentPath = "ios"
    ios.EntityData.AbsolutePath = "tailf-aaa:aaa/" + ios.EntityData.SegmentPath
    ios.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ios.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ios.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ios.EntityData.Children = types.NewOrderedMap()
    ios.EntityData.Children.Append("level", types.YChild{"Level", nil})
    for i := range ios.Level {
        ios.EntityData.Children.Append(types.GetSegmentPath(ios.Level[i]), types.YChild{"Level", ios.Level[i]})
    }
    ios.EntityData.Children.Append("privilege", types.YChild{"Privilege", nil})
    for i := range ios.Privilege {
        ios.EntityData.Children.Append(types.GetSegmentPath(ios.Privilege[i]), types.YChild{"Privilege", ios.Privilege[i]})
    }
    ios.EntityData.Leafs = types.NewOrderedMap()

    ios.EntityData.YListKeys = []string {}

    return &(ios.EntityData)
}

// Aaa_Ios_Level
type Aaa_Ios_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..15.
    Nr interface{}

    // The type is string.
    Secret interface{}

    // The type is string.
    Password interface{}

    // The type is string. The default value is \h# .
    Prompt interface{}
}

func (level *Aaa_Ios_Level) GetEntityData() *types.CommonEntityData {
    level.EntityData.YFilter = level.YFilter
    level.EntityData.YangName = "level"
    level.EntityData.BundleName = "cisco_ios_xr"
    level.EntityData.ParentYangName = "ios"
    level.EntityData.SegmentPath = "level" + types.AddKeyToken(level.Nr, "nr")
    level.EntityData.AbsolutePath = "tailf-aaa:aaa/ios/" + level.EntityData.SegmentPath
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = types.NewOrderedMap()
    level.EntityData.Leafs = types.NewOrderedMap()
    level.EntityData.Leafs.Append("nr", types.YLeaf{"Nr", level.Nr})
    level.EntityData.Leafs.Append("secret", types.YLeaf{"Secret", level.Secret})
    level.EntityData.Leafs.Append("password", types.YLeaf{"Password", level.Password})
    level.EntityData.Leafs.Append("prompt", types.YLeaf{"Prompt", level.Prompt})

    level.EntityData.YListKeys = []string {"Nr"}

    return &(level.EntityData)
}

// Aaa_Ios_Privilege
type Aaa_Ios_Privilege struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is one of the following types: string, or
    // enumeration BuiltinModes_.
    Mode interface{}

    // The type is slice of Aaa_Ios_Privilege_Level.
    Level []*Aaa_Ios_Privilege_Level
}

func (privilege *Aaa_Ios_Privilege) GetEntityData() *types.CommonEntityData {
    privilege.EntityData.YFilter = privilege.YFilter
    privilege.EntityData.YangName = "privilege"
    privilege.EntityData.BundleName = "cisco_ios_xr"
    privilege.EntityData.ParentYangName = "ios"
    privilege.EntityData.SegmentPath = "privilege" + types.AddKeyToken(privilege.Mode, "mode")
    privilege.EntityData.AbsolutePath = "tailf-aaa:aaa/ios/" + privilege.EntityData.SegmentPath
    privilege.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privilege.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privilege.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privilege.EntityData.Children = types.NewOrderedMap()
    privilege.EntityData.Children.Append("level", types.YChild{"Level", nil})
    for i := range privilege.Level {
        privilege.EntityData.Children.Append(types.GetSegmentPath(privilege.Level[i]), types.YChild{"Level", privilege.Level[i]})
    }
    privilege.EntityData.Leafs = types.NewOrderedMap()
    privilege.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", privilege.Mode})

    privilege.EntityData.YListKeys = []string {"Mode"}

    return &(privilege.EntityData)
}

// Aaa_Ios_Privilege_Level
type Aaa_Ios_Privilege_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..15.
    Nr interface{}

    // The type is slice of Aaa_Ios_Privilege_Level_Command.
    Command []*Aaa_Ios_Privilege_Level_Command
}

func (level *Aaa_Ios_Privilege_Level) GetEntityData() *types.CommonEntityData {
    level.EntityData.YFilter = level.YFilter
    level.EntityData.YangName = "level"
    level.EntityData.BundleName = "cisco_ios_xr"
    level.EntityData.ParentYangName = "privilege"
    level.EntityData.SegmentPath = "level" + types.AddKeyToken(level.Nr, "nr")
    level.EntityData.AbsolutePath = "tailf-aaa:aaa/ios/privilege/" + level.EntityData.SegmentPath
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = types.NewOrderedMap()
    level.EntityData.Children.Append("command", types.YChild{"Command", nil})
    for i := range level.Command {
        level.EntityData.Children.Append(types.GetSegmentPath(level.Command[i]), types.YChild{"Command", level.Command[i]})
    }
    level.EntityData.Leafs = types.NewOrderedMap()
    level.EntityData.Leafs.Append("nr", types.YLeaf{"Nr", level.Nr})

    level.EntityData.YListKeys = []string {"Nr"}

    return &(level.EntityData)
}

// Aaa_Ios_Privilege_Level_Command
type Aaa_Ios_Privilege_Level_Command struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}
}

func (command *Aaa_Ios_Privilege_Level_Command) GetEntityData() *types.CommonEntityData {
    command.EntityData.YFilter = command.YFilter
    command.EntityData.YangName = "command"
    command.EntityData.BundleName = "cisco_ios_xr"
    command.EntityData.ParentYangName = "level"
    command.EntityData.SegmentPath = "command" + types.AddKeyToken(command.Name, "name")
    command.EntityData.AbsolutePath = "tailf-aaa:aaa/ios/privilege/level/" + command.EntityData.SegmentPath
    command.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    command.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    command.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    command.EntityData.Children = types.NewOrderedMap()
    command.EntityData.Leafs = types.NewOrderedMap()
    command.EntityData.Leafs.Append("name", types.YLeaf{"Name", command.Name})

    command.EntityData.YListKeys = []string {"Name"}

    return &(command.EntityData)
}

// Aaa_PrivilegedAccess
type Aaa_PrivilegedAccess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    ShellAccess interface{}

    // The type is string.
    FirstUser interface{}

    // The type is string.
    FirstUserChange interface{}

    // The type is string.
    CurrentDisasterRecoveryUser interface{}
}

func (privilegedAccess *Aaa_PrivilegedAccess) GetEntityData() *types.CommonEntityData {
    privilegedAccess.EntityData.YFilter = privilegedAccess.YFilter
    privilegedAccess.EntityData.YangName = "privileged-access"
    privilegedAccess.EntityData.BundleName = "cisco_ios_xr"
    privilegedAccess.EntityData.ParentYangName = "aaa"
    privilegedAccess.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-aaa-aaa-show:privileged-access"
    privilegedAccess.EntityData.AbsolutePath = "tailf-aaa:aaa/" + privilegedAccess.EntityData.SegmentPath
    privilegedAccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privilegedAccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privilegedAccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privilegedAccess.EntityData.Children = types.NewOrderedMap()
    privilegedAccess.EntityData.Leafs = types.NewOrderedMap()
    privilegedAccess.EntityData.Leafs.Append("shell-access", types.YLeaf{"ShellAccess", privilegedAccess.ShellAccess})
    privilegedAccess.EntityData.Leafs.Append("first-user", types.YLeaf{"FirstUser", privilegedAccess.FirstUser})
    privilegedAccess.EntityData.Leafs.Append("first-user-change", types.YLeaf{"FirstUserChange", privilegedAccess.FirstUserChange})
    privilegedAccess.EntityData.Leafs.Append("current-disaster-recovery-user", types.YLeaf{"CurrentDisasterRecoveryUser", privilegedAccess.CurrentDisasterRecoveryUser})

    privilegedAccess.EntityData.YListKeys = []string {}

    return &(privilegedAccess.EntityData)
}

// Aaa_CiscoIOSXRSysadminAaaAaaShowAccounting
type Aaa_CiscoIOSXRSysadminAaaAaaShowAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogData interface{}
}

func (ciscoIOSXRSysadminAaaAaaShowAccounting *Aaa_CiscoIOSXRSysadminAaaAaaShowAccounting) GetEntityData() *types.CommonEntityData {
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.YFilter = ciscoIOSXRSysadminAaaAaaShowAccounting.YFilter
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.YangName = "accounting"
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.BundleName = "cisco_ios_xr"
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.ParentYangName = "aaa"
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-aaa-aaa-show:accounting"
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.AbsolutePath = "tailf-aaa:aaa/" + ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.SegmentPath
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.Children = types.NewOrderedMap()
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.Leafs = types.NewOrderedMap()
    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.Leafs.Append("log-data", types.YLeaf{"LogData", ciscoIOSXRSysadminAaaAaaShowAccounting.LogData})

    ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData.YListKeys = []string {}

    return &(ciscoIOSXRSysadminAaaAaaShowAccounting.EntityData)
}

// Aaa_UserGroup
type Aaa_UserGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    GrpData interface{}
}

func (userGroup *Aaa_UserGroup) GetEntityData() *types.CommonEntityData {
    userGroup.EntityData.YFilter = userGroup.YFilter
    userGroup.EntityData.YangName = "user-group"
    userGroup.EntityData.BundleName = "cisco_ios_xr"
    userGroup.EntityData.ParentYangName = "aaa"
    userGroup.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-aaa-aaa-show:user-group"
    userGroup.EntityData.AbsolutePath = "tailf-aaa:aaa/" + userGroup.EntityData.SegmentPath
    userGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userGroup.EntityData.Children = types.NewOrderedMap()
    userGroup.EntityData.Leafs = types.NewOrderedMap()
    userGroup.EntityData.Leafs.Append("grp-data", types.YLeaf{"GrpData", userGroup.GrpData})

    userGroup.EntityData.YListKeys = []string {}

    return &(userGroup.EntityData)
}

// Aaa_DisasterRecovery
type Aaa_DisasterRecovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string. Refers to tailf_aaa.Aaa_Authentication_Users_User_Name
    Username interface{}

    // The type is string.
    Password interface{}
}

func (disasterRecovery *Aaa_DisasterRecovery) GetEntityData() *types.CommonEntityData {
    disasterRecovery.EntityData.YFilter = disasterRecovery.YFilter
    disasterRecovery.EntityData.YangName = "disaster-recovery"
    disasterRecovery.EntityData.BundleName = "cisco_ios_xr"
    disasterRecovery.EntityData.ParentYangName = "aaa"
    disasterRecovery.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-aaa-disaster-recovery:disaster-recovery"
    disasterRecovery.EntityData.AbsolutePath = "tailf-aaa:aaa/" + disasterRecovery.EntityData.SegmentPath
    disasterRecovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disasterRecovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disasterRecovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disasterRecovery.EntityData.Children = types.NewOrderedMap()
    disasterRecovery.EntityData.Leafs = types.NewOrderedMap()
    disasterRecovery.EntityData.Leafs.Append("username", types.YLeaf{"Username", disasterRecovery.Username})
    disasterRecovery.EntityData.Leafs.Append("password", types.YLeaf{"Password", disasterRecovery.Password})

    disasterRecovery.EntityData.YListKeys = []string {}

    return &(disasterRecovery.EntityData)
}

// Alias
type Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is string. This attribute is mandatory.
    Expansion interface{}
}

func (alias *Alias) GetEntityData() *types.CommonEntityData {
    alias.EntityData.YFilter = alias.YFilter
    alias.EntityData.YangName = "alias"
    alias.EntityData.BundleName = "cisco_ios_xr"
    alias.EntityData.ParentYangName = "tailf-aaa"
    alias.EntityData.SegmentPath = "tailf-aaa:alias" + types.AddKeyToken(alias.Name, "name")
    alias.EntityData.AbsolutePath = alias.EntityData.SegmentPath
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = types.NewOrderedMap()
    alias.EntityData.Leafs = types.NewOrderedMap()
    alias.EntityData.Leafs.Append("name", types.YLeaf{"Name", alias.Name})
    alias.EntityData.Leafs.Append("expansion", types.YLeaf{"Expansion", alias.Expansion})

    alias.EntityData.YListKeys = []string {"Name"}

    return &(alias.EntityData)
}

// Session
// This type is a presence type.
type Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // The type is bool.
    CompleteOnSpace interface{}

    // The type is bool.
    IgnoreLeadingSpace interface{}

    // The type is interface{} with range: 0..8192.
    IdleTimeout interface{}

    // The type is bool.
    Paginate interface{}

    // The type is interface{} with range: 0..8192.
    History interface{}

    // The type is bool.
    Autowizard interface{}

    // The type is bool.
    ShowDefaults interface{}

    // The type is interface{} with range: 1..64.
    DisplayLevel interface{}

    // The type is string.
    Prompt1 interface{}

    // The type is string.
    Prompt2 interface{}
}

func (session *Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "tailf-aaa"
    session.EntityData.SegmentPath = "tailf-aaa:session"
    session.EntityData.AbsolutePath = session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("complete-on-space", types.YLeaf{"CompleteOnSpace", session.CompleteOnSpace})
    session.EntityData.Leafs.Append("ignore-leading-space", types.YLeaf{"IgnoreLeadingSpace", session.IgnoreLeadingSpace})
    session.EntityData.Leafs.Append("idle-timeout", types.YLeaf{"IdleTimeout", session.IdleTimeout})
    session.EntityData.Leafs.Append("paginate", types.YLeaf{"Paginate", session.Paginate})
    session.EntityData.Leafs.Append("history", types.YLeaf{"History", session.History})
    session.EntityData.Leafs.Append("autowizard", types.YLeaf{"Autowizard", session.Autowizard})
    session.EntityData.Leafs.Append("show-defaults", types.YLeaf{"ShowDefaults", session.ShowDefaults})
    session.EntityData.Leafs.Append("display-level", types.YLeaf{"DisplayLevel", session.DisplayLevel})
    session.EntityData.Leafs.Append("prompt1", types.YLeaf{"Prompt1", session.Prompt1})
    session.EntityData.Leafs.Append("prompt2", types.YLeaf{"Prompt2", session.Prompt2})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// User
type User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is string.
    Description interface{}

    // The type is slice of User_Alias.
    Alias []*User_Alias

    
    Session User_Session
}

func (user *User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "tailf-aaa"
    user.EntityData.SegmentPath = "tailf-aaa:user" + types.AddKeyToken(user.Name, "name")
    user.EntityData.AbsolutePath = user.EntityData.SegmentPath
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Children.Append("alias", types.YChild{"Alias", nil})
    for i := range user.Alias {
        user.EntityData.Children.Append(types.GetSegmentPath(user.Alias[i]), types.YChild{"Alias", user.Alias[i]})
    }
    user.EntityData.Children.Append("session", types.YChild{"Session", &user.Session})
    user.EntityData.Leafs = types.NewOrderedMap()
    user.EntityData.Leafs.Append("name", types.YLeaf{"Name", user.Name})
    user.EntityData.Leafs.Append("description", types.YLeaf{"Description", user.Description})

    user.EntityData.YListKeys = []string {"Name"}

    return &(user.EntityData)
}

// User_Alias
type User_Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is string. This attribute is mandatory.
    Expansion interface{}
}

func (alias *User_Alias) GetEntityData() *types.CommonEntityData {
    alias.EntityData.YFilter = alias.YFilter
    alias.EntityData.YangName = "alias"
    alias.EntityData.BundleName = "cisco_ios_xr"
    alias.EntityData.ParentYangName = "user"
    alias.EntityData.SegmentPath = "alias" + types.AddKeyToken(alias.Name, "name")
    alias.EntityData.AbsolutePath = "tailf-aaa:user/" + alias.EntityData.SegmentPath
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = types.NewOrderedMap()
    alias.EntityData.Leafs = types.NewOrderedMap()
    alias.EntityData.Leafs.Append("name", types.YLeaf{"Name", alias.Name})
    alias.EntityData.Leafs.Append("expansion", types.YLeaf{"Expansion", alias.Expansion})

    alias.EntityData.YListKeys = []string {"Name"}

    return &(alias.EntityData)
}

// User_Session
type User_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is bool.
    CompleteOnSpace interface{}

    // The type is bool.
    IgnoreLeadingSpace interface{}

    // The type is interface{} with range: 0..8192.
    IdleTimeout interface{}

    // The type is bool.
    Paginate interface{}

    // The type is interface{} with range: 0..8192.
    History interface{}

    // The type is bool.
    Autowizard interface{}

    // The type is bool.
    ShowDefaults interface{}

    // The type is interface{} with range: 1..64.
    DisplayLevel interface{}

    // The type is string.
    Prompt1 interface{}

    // The type is string.
    Prompt2 interface{}
}

func (session *User_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "user"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "tailf-aaa:user/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("complete-on-space", types.YLeaf{"CompleteOnSpace", session.CompleteOnSpace})
    session.EntityData.Leafs.Append("ignore-leading-space", types.YLeaf{"IgnoreLeadingSpace", session.IgnoreLeadingSpace})
    session.EntityData.Leafs.Append("idle-timeout", types.YLeaf{"IdleTimeout", session.IdleTimeout})
    session.EntityData.Leafs.Append("paginate", types.YLeaf{"Paginate", session.Paginate})
    session.EntityData.Leafs.Append("history", types.YLeaf{"History", session.History})
    session.EntityData.Leafs.Append("autowizard", types.YLeaf{"Autowizard", session.Autowizard})
    session.EntityData.Leafs.Append("show-defaults", types.YLeaf{"ShowDefaults", session.ShowDefaults})
    session.EntityData.Leafs.Append("display-level", types.YLeaf{"DisplayLevel", session.DisplayLevel})
    session.EntityData.Leafs.Append("prompt1", types.YLeaf{"Prompt1", session.Prompt1})
    session.EntityData.Leafs.Append("prompt2", types.YLeaf{"Prompt2", session.Prompt2})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

