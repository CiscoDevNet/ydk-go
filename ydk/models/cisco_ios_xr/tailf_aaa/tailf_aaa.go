// This module defines the Tail-f AAA data model.
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

// Dataoperationtype
type Dataoperationtype string

const (
    Dataoperationtype_r Dataoperationtype = "r"

    Dataoperationtype_rx Dataoperationtype = "rx"

    Dataoperationtype_x Dataoperationtype = "x"

    Dataoperationtype_rw Dataoperationtype = "rw"

    Dataoperationtype_rwx Dataoperationtype = "rwx"

    Dataoperationtype_wx Dataoperationtype = "wx"

    Dataoperationtype_w Dataoperationtype = "w"

    Dataoperationtype_c Dataoperationtype = "c"

    Dataoperationtype_cr Dataoperationtype = "cr"

    Dataoperationtype_cu Dataoperationtype = "cu"

    Dataoperationtype_cd Dataoperationtype = "cd"

    Dataoperationtype_cx Dataoperationtype = "cx"

    Dataoperationtype_cru Dataoperationtype = "cru"

    Dataoperationtype_crd Dataoperationtype = "crd"

    Dataoperationtype_crx Dataoperationtype = "crx"

    Dataoperationtype_cud Dataoperationtype = "cud"

    Dataoperationtype_cux Dataoperationtype = "cux"

    Dataoperationtype_cdx Dataoperationtype = "cdx"

    Dataoperationtype_crud Dataoperationtype = "crud"

    Dataoperationtype_crux Dataoperationtype = "crux"

    Dataoperationtype_crdx Dataoperationtype = "crdx"

    Dataoperationtype_cudx Dataoperationtype = "cudx"

    Dataoperationtype_crudx Dataoperationtype = "crudx"

    Dataoperationtype_ru Dataoperationtype = "ru"

    Dataoperationtype_rd Dataoperationtype = "rd"

    Dataoperationtype_rud Dataoperationtype = "rud"

    Dataoperationtype_rux Dataoperationtype = "rux"

    Dataoperationtype_rdx Dataoperationtype = "rdx"

    Dataoperationtype_u Dataoperationtype = "u"

    Dataoperationtype_ud Dataoperationtype = "ud"

    Dataoperationtype_ux Dataoperationtype = "ux"

    Dataoperationtype_d Dataoperationtype = "d"

    Dataoperationtype_dx Dataoperationtype = "dx"
)

// Cmdoperationtype
type Cmdoperationtype string

const (
    Cmdoperationtype_r Cmdoperationtype = "r"

    Cmdoperationtype_rx Cmdoperationtype = "rx"

    Cmdoperationtype_x Cmdoperationtype = "x"
)

// Action
type Action string

const (
    Action_accept Action = "accept"

    Action_reject Action = "reject"

    Action_accept_log Action = "accept_log"
)

// Builtinmodes
type Builtinmodes string

const (
    Builtinmodes_exec Builtinmodes = "exec"

    Builtinmodes_configure Builtinmodes = "configure"
)

// Builtinmodes_
type Builtinmodes_ string

const (
    Builtinmodes__exec Builtinmodes_ = "exec"

    Builtinmodes__configure Builtinmodes_ = "configure"
)

// Aaa
type Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Authentication Aaa_Authentication

    
    Authorization Aaa_Authorization

    
    Ios Aaa_Ios

    
    PrivilegedAccess Aaa_PrivilegedAccess

    
    Accounting Aaa_Accounting

    
    UserGroup Aaa_UserGroup

    
    DisasterRecovery Aaa_DisasterRecovery
}

func (aaa *Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "tailf-aaa"
    aaa.EntityData.SegmentPath = "tailf-aaa:aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = make(map[string]types.YChild)
    aaa.EntityData.Children["authentication"] = types.YChild{"Authentication", &aaa.Authentication}
    aaa.EntityData.Children["authorization"] = types.YChild{"Authorization", &aaa.Authorization}
    aaa.EntityData.Children["ios"] = types.YChild{"Ios", &aaa.Ios}
    aaa.EntityData.Children["Cisco-IOS-XR-sysadmin-aaa-aaa-show:privileged-access"] = types.YChild{"PrivilegedAccess", &aaa.PrivilegedAccess}
    aaa.EntityData.Children["Cisco-IOS-XR-sysadmin-aaa-aaa-show:accounting"] = types.YChild{"Accounting", &aaa.Accounting}
    aaa.EntityData.Children["Cisco-IOS-XR-sysadmin-aaa-aaa-show:user-group"] = types.YChild{"UserGroup", &aaa.UserGroup}
    aaa.EntityData.Children["Cisco-IOS-XR-sysadmin-aaa-disaster-recovery:disaster-recovery"] = types.YChild{"DisasterRecovery", &aaa.DisasterRecovery}
    aaa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aaa.EntityData)
}

// Aaa_Authentication
type Aaa_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Users Aaa_Authentication_Users

    
    Groups Aaa_Authentication_Groups
}

func (authentication *Aaa_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "aaa"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Children["users"] = types.YChild{"Users", &authentication.Users}
    authentication.EntityData.Children["groups"] = types.YChild{"Groups", &authentication.Groups}
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authentication.EntityData)
}

// Aaa_Authentication_Users
type Aaa_Authentication_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authentication_Users_User.
    User []Aaa_Authentication_Users_User
}

func (users *Aaa_Authentication_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "authentication"
    users.EntityData.SegmentPath = "users"
    users.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    users.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    users.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    users.EntityData.Children = make(map[string]types.YChild)
    users.EntityData.Children["user"] = types.YChild{"User", nil}
    for i := range users.User {
        users.EntityData.Children[types.GetSegmentPath(&users.User[i])] = types.YChild{"User", &users.User[i]}
    }
    users.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(users.EntityData)
}

// Aaa_Authentication_Users_User
type Aaa_Authentication_Users_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: -2147483648..2147483647. This attribute
    // is mandatory.
    Uid interface{}

    // The type is interface{} with range: -2147483648..2147483647. This attribute
    // is mandatory.
    Gid interface{}

    // The type is string. This attribute is mandatory.
    Password interface{}

    // The type is string. This attribute is mandatory.
    SshKeydir interface{}

    // The type is string. This attribute is mandatory.
    Homedir interface{}
}

func (user *Aaa_Authentication_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + "[name='" + fmt.Sprintf("%v", user.Name) + "']"
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = make(map[string]types.YChild)
    user.EntityData.Leafs = make(map[string]types.YLeaf)
    user.EntityData.Leafs["name"] = types.YLeaf{"Name", user.Name}
    user.EntityData.Leafs["uid"] = types.YLeaf{"Uid", user.Uid}
    user.EntityData.Leafs["gid"] = types.YLeaf{"Gid", user.Gid}
    user.EntityData.Leafs["password"] = types.YLeaf{"Password", user.Password}
    user.EntityData.Leafs["ssh_keydir"] = types.YLeaf{"SshKeydir", user.SshKeydir}
    user.EntityData.Leafs["homedir"] = types.YLeaf{"Homedir", user.Homedir}
    return &(user.EntityData)
}

// Aaa_Authentication_Groups
type Aaa_Authentication_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authentication_Groups_Group.
    Group []Aaa_Authentication_Groups_Group
}

func (groups *Aaa_Authentication_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "authentication"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Aaa_Authentication_Groups_Group
type Aaa_Authentication_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is interface{} with range: -2147483648..2147483647.
    Gid interface{}

    // The type is string. This attribute is mandatory.
    Users interface{}
}

func (group *Aaa_Authentication_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[name='" + fmt.Sprintf("%v", group.Name) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["name"] = types.YLeaf{"Name", group.Name}
    group.EntityData.Leafs["gid"] = types.YLeaf{"Gid", group.Gid}
    group.EntityData.Leafs["users"] = types.YLeaf{"Users", group.Users}
    return &(group.EntityData)
}

// Aaa_Authorization
type Aaa_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cmdrules Aaa_Authorization_Cmdrules

    
    Datarules Aaa_Authorization_Datarules
}

func (authorization *Aaa_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "aaa"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = make(map[string]types.YChild)
    authorization.EntityData.Children["cmdrules"] = types.YChild{"Cmdrules", &authorization.Cmdrules}
    authorization.EntityData.Children["datarules"] = types.YChild{"Datarules", &authorization.Datarules}
    authorization.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authorization.EntityData)
}

// Aaa_Authorization_Cmdrules
type Aaa_Authorization_Cmdrules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authorization_Cmdrules_Cmdrule.
    Cmdrule []Aaa_Authorization_Cmdrules_Cmdrule
}

func (cmdrules *Aaa_Authorization_Cmdrules) GetEntityData() *types.CommonEntityData {
    cmdrules.EntityData.YFilter = cmdrules.YFilter
    cmdrules.EntityData.YangName = "cmdrules"
    cmdrules.EntityData.BundleName = "cisco_ios_xr"
    cmdrules.EntityData.ParentYangName = "authorization"
    cmdrules.EntityData.SegmentPath = "cmdrules"
    cmdrules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cmdrules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cmdrules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cmdrules.EntityData.Children = make(map[string]types.YChild)
    cmdrules.EntityData.Children["cmdrule"] = types.YChild{"Cmdrule", nil}
    for i := range cmdrules.Cmdrule {
        cmdrules.EntityData.Children[types.GetSegmentPath(&cmdrules.Cmdrule[i])] = types.YChild{"Cmdrule", &cmdrules.Cmdrule[i]}
    }
    cmdrules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cmdrules.EntityData)
}

// Aaa_Authorization_Cmdrules_Cmdrule
type Aaa_Authorization_Cmdrules_Cmdrule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // The type is string. This attribute is mandatory.
    Context interface{}

    // The type is string. This attribute is mandatory.
    Command interface{}

    // The type is string. This attribute is mandatory.
    Group interface{}

    // The type is Cmdoperationtype. This attribute is mandatory.
    Ops interface{}

    // The type is Action. This attribute is mandatory.
    Action interface{}
}

func (cmdrule *Aaa_Authorization_Cmdrules_Cmdrule) GetEntityData() *types.CommonEntityData {
    cmdrule.EntityData.YFilter = cmdrule.YFilter
    cmdrule.EntityData.YangName = "cmdrule"
    cmdrule.EntityData.BundleName = "cisco_ios_xr"
    cmdrule.EntityData.ParentYangName = "cmdrules"
    cmdrule.EntityData.SegmentPath = "cmdrule" + "[index='" + fmt.Sprintf("%v", cmdrule.Index) + "']"
    cmdrule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cmdrule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cmdrule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cmdrule.EntityData.Children = make(map[string]types.YChild)
    cmdrule.EntityData.Leafs = make(map[string]types.YLeaf)
    cmdrule.EntityData.Leafs["index"] = types.YLeaf{"Index", cmdrule.Index}
    cmdrule.EntityData.Leafs["context"] = types.YLeaf{"Context", cmdrule.Context}
    cmdrule.EntityData.Leafs["command"] = types.YLeaf{"Command", cmdrule.Command}
    cmdrule.EntityData.Leafs["group"] = types.YLeaf{"Group", cmdrule.Group}
    cmdrule.EntityData.Leafs["ops"] = types.YLeaf{"Ops", cmdrule.Ops}
    cmdrule.EntityData.Leafs["action"] = types.YLeaf{"Action", cmdrule.Action}
    return &(cmdrule.EntityData)
}

// Aaa_Authorization_Datarules
type Aaa_Authorization_Datarules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Authorization_Datarules_Datarule.
    Datarule []Aaa_Authorization_Datarules_Datarule
}

func (datarules *Aaa_Authorization_Datarules) GetEntityData() *types.CommonEntityData {
    datarules.EntityData.YFilter = datarules.YFilter
    datarules.EntityData.YangName = "datarules"
    datarules.EntityData.BundleName = "cisco_ios_xr"
    datarules.EntityData.ParentYangName = "authorization"
    datarules.EntityData.SegmentPath = "datarules"
    datarules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    datarules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    datarules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    datarules.EntityData.Children = make(map[string]types.YChild)
    datarules.EntityData.Children["datarule"] = types.YChild{"Datarule", nil}
    for i := range datarules.Datarule {
        datarules.EntityData.Children[types.GetSegmentPath(&datarules.Datarule[i])] = types.YChild{"Datarule", &datarules.Datarule[i]}
    }
    datarules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(datarules.EntityData)
}

// Aaa_Authorization_Datarules_Datarule
type Aaa_Authorization_Datarules_Datarule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // The type is Dataoperationtype. This attribute is mandatory.
    Ops interface{}

    // The type is Action. This attribute is mandatory.
    Action interface{}
}

func (datarule *Aaa_Authorization_Datarules_Datarule) GetEntityData() *types.CommonEntityData {
    datarule.EntityData.YFilter = datarule.YFilter
    datarule.EntityData.YangName = "datarule"
    datarule.EntityData.BundleName = "cisco_ios_xr"
    datarule.EntityData.ParentYangName = "datarules"
    datarule.EntityData.SegmentPath = "datarule" + "[index='" + fmt.Sprintf("%v", datarule.Index) + "']"
    datarule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    datarule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    datarule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    datarule.EntityData.Children = make(map[string]types.YChild)
    datarule.EntityData.Leafs = make(map[string]types.YLeaf)
    datarule.EntityData.Leafs["index"] = types.YLeaf{"Index", datarule.Index}
    datarule.EntityData.Leafs["namespace"] = types.YLeaf{"Namespace", datarule.Namespace}
    datarule.EntityData.Leafs["context"] = types.YLeaf{"Context", datarule.Context}
    datarule.EntityData.Leafs["keypath"] = types.YLeaf{"Keypath", datarule.Keypath}
    datarule.EntityData.Leafs["group"] = types.YLeaf{"Group", datarule.Group}
    datarule.EntityData.Leafs["ops"] = types.YLeaf{"Ops", datarule.Ops}
    datarule.EntityData.Leafs["action"] = types.YLeaf{"Action", datarule.Action}
    return &(datarule.EntityData)
}

// Aaa_Ios
// This type is a presence type.
type Aaa_Ios struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Aaa_Ios_Level.
    Level []Aaa_Ios_Level

    // The type is slice of Aaa_Ios_Privilege.
    Privilege []Aaa_Ios_Privilege
}

func (ios *Aaa_Ios) GetEntityData() *types.CommonEntityData {
    ios.EntityData.YFilter = ios.YFilter
    ios.EntityData.YangName = "ios"
    ios.EntityData.BundleName = "cisco_ios_xr"
    ios.EntityData.ParentYangName = "aaa"
    ios.EntityData.SegmentPath = "ios"
    ios.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ios.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ios.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ios.EntityData.Children = make(map[string]types.YChild)
    ios.EntityData.Children["level"] = types.YChild{"Level", nil}
    for i := range ios.Level {
        ios.EntityData.Children[types.GetSegmentPath(&ios.Level[i])] = types.YChild{"Level", &ios.Level[i]}
    }
    ios.EntityData.Children["privilege"] = types.YChild{"Privilege", nil}
    for i := range ios.Privilege {
        ios.EntityData.Children[types.GetSegmentPath(&ios.Privilege[i])] = types.YChild{"Privilege", &ios.Privilege[i]}
    }
    ios.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ios.EntityData)
}

// Aaa_Ios_Level
type Aaa_Ios_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    level.EntityData.SegmentPath = "level" + "[nr='" + fmt.Sprintf("%v", level.Nr) + "']"
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = make(map[string]types.YChild)
    level.EntityData.Leafs = make(map[string]types.YLeaf)
    level.EntityData.Leafs["nr"] = types.YLeaf{"Nr", level.Nr}
    level.EntityData.Leafs["secret"] = types.YLeaf{"Secret", level.Secret}
    level.EntityData.Leafs["password"] = types.YLeaf{"Password", level.Password}
    level.EntityData.Leafs["prompt"] = types.YLeaf{"Prompt", level.Prompt}
    return &(level.EntityData)
}

// Aaa_Ios_Privilege
type Aaa_Ios_Privilege struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following types: string, or
    // enumeration Builtinmodes_.
    Mode interface{}

    // The type is slice of Aaa_Ios_Privilege_Level.
    Level []Aaa_Ios_Privilege_Level
}

func (privilege *Aaa_Ios_Privilege) GetEntityData() *types.CommonEntityData {
    privilege.EntityData.YFilter = privilege.YFilter
    privilege.EntityData.YangName = "privilege"
    privilege.EntityData.BundleName = "cisco_ios_xr"
    privilege.EntityData.ParentYangName = "ios"
    privilege.EntityData.SegmentPath = "privilege" + "[mode='" + fmt.Sprintf("%v", privilege.Mode) + "']"
    privilege.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privilege.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privilege.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privilege.EntityData.Children = make(map[string]types.YChild)
    privilege.EntityData.Children["level"] = types.YChild{"Level", nil}
    for i := range privilege.Level {
        privilege.EntityData.Children[types.GetSegmentPath(&privilege.Level[i])] = types.YChild{"Level", &privilege.Level[i]}
    }
    privilege.EntityData.Leafs = make(map[string]types.YLeaf)
    privilege.EntityData.Leafs["mode"] = types.YLeaf{"Mode", privilege.Mode}
    return &(privilege.EntityData)
}

// Aaa_Ios_Privilege_Level
type Aaa_Ios_Privilege_Level struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is interface{} with range: 0..15.
    Nr interface{}

    // The type is slice of Aaa_Ios_Privilege_Level_Command.
    Command []Aaa_Ios_Privilege_Level_Command
}

func (level *Aaa_Ios_Privilege_Level) GetEntityData() *types.CommonEntityData {
    level.EntityData.YFilter = level.YFilter
    level.EntityData.YangName = "level"
    level.EntityData.BundleName = "cisco_ios_xr"
    level.EntityData.ParentYangName = "privilege"
    level.EntityData.SegmentPath = "level" + "[nr='" + fmt.Sprintf("%v", level.Nr) + "']"
    level.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    level.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    level.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    level.EntityData.Children = make(map[string]types.YChild)
    level.EntityData.Children["command"] = types.YChild{"Command", nil}
    for i := range level.Command {
        level.EntityData.Children[types.GetSegmentPath(&level.Command[i])] = types.YChild{"Command", &level.Command[i]}
    }
    level.EntityData.Leafs = make(map[string]types.YLeaf)
    level.EntityData.Leafs["nr"] = types.YLeaf{"Nr", level.Nr}
    return &(level.EntityData)
}

// Aaa_Ios_Privilege_Level_Command
type Aaa_Ios_Privilege_Level_Command struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}
}

func (command *Aaa_Ios_Privilege_Level_Command) GetEntityData() *types.CommonEntityData {
    command.EntityData.YFilter = command.YFilter
    command.EntityData.YangName = "command"
    command.EntityData.BundleName = "cisco_ios_xr"
    command.EntityData.ParentYangName = "level"
    command.EntityData.SegmentPath = "command" + "[name='" + fmt.Sprintf("%v", command.Name) + "']"
    command.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    command.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    command.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    command.EntityData.Children = make(map[string]types.YChild)
    command.EntityData.Leafs = make(map[string]types.YLeaf)
    command.EntityData.Leafs["name"] = types.YLeaf{"Name", command.Name}
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
    privilegedAccess.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privilegedAccess.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privilegedAccess.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privilegedAccess.EntityData.Children = make(map[string]types.YChild)
    privilegedAccess.EntityData.Leafs = make(map[string]types.YLeaf)
    privilegedAccess.EntityData.Leafs["shell-access"] = types.YLeaf{"ShellAccess", privilegedAccess.ShellAccess}
    privilegedAccess.EntityData.Leafs["first-user"] = types.YLeaf{"FirstUser", privilegedAccess.FirstUser}
    privilegedAccess.EntityData.Leafs["first-user-change"] = types.YLeaf{"FirstUserChange", privilegedAccess.FirstUserChange}
    privilegedAccess.EntityData.Leafs["current-disaster-recovery-user"] = types.YLeaf{"CurrentDisasterRecoveryUser", privilegedAccess.CurrentDisasterRecoveryUser}
    return &(privilegedAccess.EntityData)
}

// Aaa_Accounting
type Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    LogData interface{}
}

func (accounting *Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-aaa-aaa-show:accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    accounting.EntityData.Leafs["log-data"] = types.YLeaf{"LogData", accounting.LogData}
    return &(accounting.EntityData)
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
    userGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userGroup.EntityData.Children = make(map[string]types.YChild)
    userGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    userGroup.EntityData.Leafs["grp-data"] = types.YLeaf{"GrpData", userGroup.GrpData}
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
    disasterRecovery.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disasterRecovery.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disasterRecovery.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disasterRecovery.EntityData.Children = make(map[string]types.YChild)
    disasterRecovery.EntityData.Leafs = make(map[string]types.YLeaf)
    disasterRecovery.EntityData.Leafs["username"] = types.YLeaf{"Username", disasterRecovery.Username}
    disasterRecovery.EntityData.Leafs["password"] = types.YLeaf{"Password", disasterRecovery.Password}
    return &(disasterRecovery.EntityData)
}

// Alias
type Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    alias.EntityData.SegmentPath = "tailf-aaa:alias" + "[name='" + fmt.Sprintf("%v", alias.Name) + "']"
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = make(map[string]types.YChild)
    alias.EntityData.Leafs = make(map[string]types.YLeaf)
    alias.EntityData.Leafs["name"] = types.YLeaf{"Name", alias.Name}
    alias.EntityData.Leafs["expansion"] = types.YLeaf{"Expansion", alias.Expansion}
    return &(alias.EntityData)
}

// Session
// This type is a presence type.
type Session struct {
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

func (session *Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "tailf-aaa"
    session.EntityData.SegmentPath = "tailf-aaa:session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["complete-on-space"] = types.YLeaf{"CompleteOnSpace", session.CompleteOnSpace}
    session.EntityData.Leafs["ignore-leading-space"] = types.YLeaf{"IgnoreLeadingSpace", session.IgnoreLeadingSpace}
    session.EntityData.Leafs["idle-timeout"] = types.YLeaf{"IdleTimeout", session.IdleTimeout}
    session.EntityData.Leafs["paginate"] = types.YLeaf{"Paginate", session.Paginate}
    session.EntityData.Leafs["history"] = types.YLeaf{"History", session.History}
    session.EntityData.Leafs["autowizard"] = types.YLeaf{"Autowizard", session.Autowizard}
    session.EntityData.Leafs["show-defaults"] = types.YLeaf{"ShowDefaults", session.ShowDefaults}
    session.EntityData.Leafs["display-level"] = types.YLeaf{"DisplayLevel", session.DisplayLevel}
    session.EntityData.Leafs["prompt1"] = types.YLeaf{"Prompt1", session.Prompt1}
    session.EntityData.Leafs["prompt2"] = types.YLeaf{"Prompt2", session.Prompt2}
    return &(session.EntityData)
}

// User
type User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Name interface{}

    // The type is string.
    Description interface{}

    // The type is slice of User_Alias.
    Alias []User_Alias

    
    Session User_Session
}

func (user *User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "tailf-aaa"
    user.EntityData.SegmentPath = "tailf-aaa:user" + "[name='" + fmt.Sprintf("%v", user.Name) + "']"
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = make(map[string]types.YChild)
    user.EntityData.Children["alias"] = types.YChild{"Alias", nil}
    for i := range user.Alias {
        user.EntityData.Children[types.GetSegmentPath(&user.Alias[i])] = types.YChild{"Alias", &user.Alias[i]}
    }
    user.EntityData.Children["session"] = types.YChild{"Session", &user.Session}
    user.EntityData.Leafs = make(map[string]types.YLeaf)
    user.EntityData.Leafs["name"] = types.YLeaf{"Name", user.Name}
    user.EntityData.Leafs["description"] = types.YLeaf{"Description", user.Description}
    return &(user.EntityData)
}

// User_Alias
type User_Alias struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    alias.EntityData.SegmentPath = "alias" + "[name='" + fmt.Sprintf("%v", alias.Name) + "']"
    alias.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alias.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alias.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alias.EntityData.Children = make(map[string]types.YChild)
    alias.EntityData.Leafs = make(map[string]types.YLeaf)
    alias.EntityData.Leafs["name"] = types.YLeaf{"Name", alias.Name}
    alias.EntityData.Leafs["expansion"] = types.YLeaf{"Expansion", alias.Expansion}
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
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["complete-on-space"] = types.YLeaf{"CompleteOnSpace", session.CompleteOnSpace}
    session.EntityData.Leafs["ignore-leading-space"] = types.YLeaf{"IgnoreLeadingSpace", session.IgnoreLeadingSpace}
    session.EntityData.Leafs["idle-timeout"] = types.YLeaf{"IdleTimeout", session.IdleTimeout}
    session.EntityData.Leafs["paginate"] = types.YLeaf{"Paginate", session.Paginate}
    session.EntityData.Leafs["history"] = types.YLeaf{"History", session.History}
    session.EntityData.Leafs["autowizard"] = types.YLeaf{"Autowizard", session.Autowizard}
    session.EntityData.Leafs["show-defaults"] = types.YLeaf{"ShowDefaults", session.ShowDefaults}
    session.EntityData.Leafs["display-level"] = types.YLeaf{"DisplayLevel", session.DisplayLevel}
    session.EntityData.Leafs["prompt1"] = types.YLeaf{"Prompt1", session.Prompt1}
    session.EntityData.Leafs["prompt2"] = types.YLeaf{"Prompt2", session.Prompt2}
    return &(session.EntityData)
}

