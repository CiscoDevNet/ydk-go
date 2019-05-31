// This module contains definitions
// for the Calvados model objects.
// 
// This module holds a sample operational data store.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package opertest1

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package opertest1"))
    ydk.RegisterEntity("{http://www.cisco.com/panini/calvados/opertest1 oper}", reflect.TypeOf(Oper{}))
    ydk.RegisterEntity("opertest1:oper", reflect.TypeOf(Oper{}))
    ydk.RegisterEntity("{http://www.cisco.com/panini/calvados/opertest1 actions}", reflect.TypeOf(Actions{}))
    ydk.RegisterEntity("opertest1:actions", reflect.TypeOf(Actions{}))
}

// Oper
type Oper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is interface{} with range: 0..9.
    Fred interface{}

    // The type is interface{} with range: 0..9999.
    Barney interface{}

    // The type is interface{} with range: 0..9999.
    Wilma interface{}

    // The type is interface{} with range: 0..9999.
    Betty interface{}

    // The type is slice of Oper_Slates.
    Slates []*Oper_Slates

    
    Uname Oper_Uname

    
    Uptime Oper_Uptime

    
    W Oper_W
}

func (oper *Oper) GetEntityData() *types.CommonEntityData {
    oper.EntityData.YFilter = oper.YFilter
    oper.EntityData.YangName = "oper"
    oper.EntityData.BundleName = "cisco_ios_xr"
    oper.EntityData.ParentYangName = "opertest1"
    oper.EntityData.SegmentPath = "opertest1:oper"
    oper.EntityData.AbsolutePath = oper.EntityData.SegmentPath
    oper.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oper.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oper.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oper.EntityData.Children = types.NewOrderedMap()
    oper.EntityData.Children.Append("slates", types.YChild{"Slates", nil})
    for i := range oper.Slates {
        oper.EntityData.Children.Append(types.GetSegmentPath(oper.Slates[i]), types.YChild{"Slates", oper.Slates[i]})
    }
    oper.EntityData.Children.Append("uname", types.YChild{"Uname", &oper.Uname})
    oper.EntityData.Children.Append("uptime", types.YChild{"Uptime", &oper.Uptime})
    oper.EntityData.Children.Append("w", types.YChild{"W", &oper.W})
    oper.EntityData.Leafs = types.NewOrderedMap()
    oper.EntityData.Leafs.Append("fred", types.YLeaf{"Fred", oper.Fred})
    oper.EntityData.Leafs.Append("barney", types.YLeaf{"Barney", oper.Barney})
    oper.EntityData.Leafs.Append("wilma", types.YLeaf{"Wilma", oper.Wilma})
    oper.EntityData.Leafs.Append("betty", types.YLeaf{"Betty", oper.Betty})

    oper.EntityData.YListKeys = []string {}

    return &(oper.EntityData)
}

// Oper_Slates
type Oper_Slates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is interface{} with range: 0..999.
    Slatenum interface{}

    // The type is interface{} with range: 0..99999999.
    Mrslate interface{}
}

func (slates *Oper_Slates) GetEntityData() *types.CommonEntityData {
    slates.EntityData.YFilter = slates.YFilter
    slates.EntityData.YangName = "slates"
    slates.EntityData.BundleName = "cisco_ios_xr"
    slates.EntityData.ParentYangName = "oper"
    slates.EntityData.SegmentPath = "slates" + types.AddKeyToken(slates.Slatenum, "slatenum")
    slates.EntityData.AbsolutePath = "opertest1:oper/" + slates.EntityData.SegmentPath
    slates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slates.EntityData.Children = types.NewOrderedMap()
    slates.EntityData.Leafs = types.NewOrderedMap()
    slates.EntityData.Leafs.Append("slatenum", types.YLeaf{"Slatenum", slates.Slatenum})
    slates.EntityData.Leafs.Append("mrslate", types.YLeaf{"Mrslate", slates.Mrslate})

    slates.EntityData.YListKeys = []string {"Slatenum"}

    return &(slates.EntityData)
}

// Oper_Uname
type Oper_Uname struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The operating system name. The type is string.
    System interface{}

    // The hostname or nodename. The type is string.
    Nodename interface{}

    // The release level of the operating system implementation. The type is
    // string.
    Release interface{}

    // The operating system implmentation version level. The type is string.
    Version interface{}

    // The name of the hardware type the system is running on. The type is string.
    Machine interface{}

    // All the fields. The type is string.
    All interface{}
}

func (uname *Oper_Uname) GetEntityData() *types.CommonEntityData {
    uname.EntityData.YFilter = uname.YFilter
    uname.EntityData.YangName = "uname"
    uname.EntityData.BundleName = "cisco_ios_xr"
    uname.EntityData.ParentYangName = "oper"
    uname.EntityData.SegmentPath = "uname"
    uname.EntityData.AbsolutePath = "opertest1:oper/" + uname.EntityData.SegmentPath
    uname.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uname.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uname.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uname.EntityData.Children = types.NewOrderedMap()
    uname.EntityData.Leafs = types.NewOrderedMap()
    uname.EntityData.Leafs.Append("system", types.YLeaf{"System", uname.System})
    uname.EntityData.Leafs.Append("nodename", types.YLeaf{"Nodename", uname.Nodename})
    uname.EntityData.Leafs.Append("release", types.YLeaf{"Release", uname.Release})
    uname.EntityData.Leafs.Append("version", types.YLeaf{"Version", uname.Version})
    uname.EntityData.Leafs.Append("machine", types.YLeaf{"Machine", uname.Machine})
    uname.EntityData.Leafs.Append("all", types.YLeaf{"All", uname.All})

    uname.EntityData.YListKeys = []string {}

    return &(uname.EntityData)
}

// Oper_Uptime
type Oper_Uptime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current time. The type is string.
    Curtime interface{}

    // Time since boot. The type is string.
    Uptime interface{}

    // Weighted Exponential One-Minute Load Average. The type is string.
    Loadavg1min interface{}

    // Weighted Exponential Five-Minute Load Average. The type is string.
    Loadavg5min interface{}

    // Weighted Exponential Fifteen-Minute Load Average. The type is string.
    Loadavg15min interface{}
}

func (uptime *Oper_Uptime) GetEntityData() *types.CommonEntityData {
    uptime.EntityData.YFilter = uptime.YFilter
    uptime.EntityData.YangName = "uptime"
    uptime.EntityData.BundleName = "cisco_ios_xr"
    uptime.EntityData.ParentYangName = "oper"
    uptime.EntityData.SegmentPath = "uptime"
    uptime.EntityData.AbsolutePath = "opertest1:oper/" + uptime.EntityData.SegmentPath
    uptime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    uptime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    uptime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    uptime.EntityData.Children = types.NewOrderedMap()
    uptime.EntityData.Leafs = types.NewOrderedMap()
    uptime.EntityData.Leafs.Append("curtime", types.YLeaf{"Curtime", uptime.Curtime})
    uptime.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", uptime.Uptime})
    uptime.EntityData.Leafs.Append("loadavg1min", types.YLeaf{"Loadavg1min", uptime.Loadavg1min})
    uptime.EntityData.Leafs.Append("loadavg5min", types.YLeaf{"Loadavg5min", uptime.Loadavg5min})
    uptime.EntityData.Leafs.Append("loadavg15min", types.YLeaf{"Loadavg15min", uptime.Loadavg15min})

    uptime.EntityData.YListKeys = []string {}

    return &(uptime.EntityData)
}

// Oper_W
type Oper_W struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    HeaderInfo Oper_W_HeaderInfo

    
    Users Oper_W_Users
}

func (w *Oper_W) GetEntityData() *types.CommonEntityData {
    w.EntityData.YFilter = w.YFilter
    w.EntityData.YangName = "w"
    w.EntityData.BundleName = "cisco_ios_xr"
    w.EntityData.ParentYangName = "oper"
    w.EntityData.SegmentPath = "w"
    w.EntityData.AbsolutePath = "opertest1:oper/" + w.EntityData.SegmentPath
    w.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    w.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    w.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    w.EntityData.Children = types.NewOrderedMap()
    w.EntityData.Children.Append("header-info", types.YChild{"HeaderInfo", &w.HeaderInfo})
    w.EntityData.Children.Append("users", types.YChild{"Users", &w.Users})
    w.EntityData.Leafs = types.NewOrderedMap()

    w.EntityData.YListKeys = []string {}

    return &(w.EntityData)
}

// Oper_W_HeaderInfo
type Oper_W_HeaderInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current time. The type is string.
    Curtime interface{}

    // Time since boot. The type is string.
    Uptime interface{}

    // Weighted Exponential One-Minute Load Average. The type is string.
    Loadavg1min interface{}

    // Weighted Exponential Five-Minute Load Average. The type is string.
    Loadavg5min interface{}

    // Weighted Exponential Fifteen-Minute Load Average. The type is string.
    Loadavg15min interface{}
}

func (headerInfo *Oper_W_HeaderInfo) GetEntityData() *types.CommonEntityData {
    headerInfo.EntityData.YFilter = headerInfo.YFilter
    headerInfo.EntityData.YangName = "header-info"
    headerInfo.EntityData.BundleName = "cisco_ios_xr"
    headerInfo.EntityData.ParentYangName = "w"
    headerInfo.EntityData.SegmentPath = "header-info"
    headerInfo.EntityData.AbsolutePath = "opertest1:oper/w/" + headerInfo.EntityData.SegmentPath
    headerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    headerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    headerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    headerInfo.EntityData.Children = types.NewOrderedMap()
    headerInfo.EntityData.Leafs = types.NewOrderedMap()
    headerInfo.EntityData.Leafs.Append("curtime", types.YLeaf{"Curtime", headerInfo.Curtime})
    headerInfo.EntityData.Leafs.Append("uptime", types.YLeaf{"Uptime", headerInfo.Uptime})
    headerInfo.EntityData.Leafs.Append("loadavg1min", types.YLeaf{"Loadavg1min", headerInfo.Loadavg1min})
    headerInfo.EntityData.Leafs.Append("loadavg5min", types.YLeaf{"Loadavg5min", headerInfo.Loadavg5min})
    headerInfo.EntityData.Leafs.Append("loadavg15min", types.YLeaf{"Loadavg15min", headerInfo.Loadavg15min})

    headerInfo.EntityData.YListKeys = []string {}

    return &(headerInfo.EntityData)
}

// Oper_W_Users
type Oper_W_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Oper_W_Users_User.
    User []*Oper_W_Users_User
}

func (users *Oper_W_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "w"
    users.EntityData.SegmentPath = "users"
    users.EntityData.AbsolutePath = "opertest1:oper/w/" + users.EntityData.SegmentPath
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

// Oper_W_Users_User
type Oper_W_Users_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string.
    Tty interface{}

    // The type is string.
    User interface{}

    // The type is string.
    From interface{}

    // The type is string.
    LoginAt interface{}

    // The type is string.
    Idle interface{}

    // The type is string.
    Jcpu interface{}

    // The type is string.
    Pcpu interface{}

    // The type is string.
    What interface{}
}

func (user *Oper_W_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + types.AddKeyToken(user.Tty, "tty")
    user.EntityData.AbsolutePath = "opertest1:oper/w/users/" + user.EntityData.SegmentPath
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Leafs = types.NewOrderedMap()
    user.EntityData.Leafs.Append("tty", types.YLeaf{"Tty", user.Tty})
    user.EntityData.Leafs.Append("user", types.YLeaf{"User", user.User})
    user.EntityData.Leafs.Append("from", types.YLeaf{"From", user.From})
    user.EntityData.Leafs.Append("login-at", types.YLeaf{"LoginAt", user.LoginAt})
    user.EntityData.Leafs.Append("idle", types.YLeaf{"Idle", user.Idle})
    user.EntityData.Leafs.Append("jcpu", types.YLeaf{"Jcpu", user.Jcpu})
    user.EntityData.Leafs.Append("pcpu", types.YLeaf{"Pcpu", user.Pcpu})
    user.EntityData.Leafs.Append("what", types.YLeaf{"What", user.What})

    user.EntityData.YListKeys = []string {"Tty"}

    return &(user.EntityData)
}

// Actions
type Actions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (actions *Actions) GetEntityData() *types.CommonEntityData {
    actions.EntityData.YFilter = actions.YFilter
    actions.EntityData.YangName = "actions"
    actions.EntityData.BundleName = "cisco_ios_xr"
    actions.EntityData.ParentYangName = "opertest1"
    actions.EntityData.SegmentPath = "opertest1:actions"
    actions.EntityData.AbsolutePath = actions.EntityData.SegmentPath
    actions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    actions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    actions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    actions.EntityData.Children = types.NewOrderedMap()
    actions.EntityData.Leafs = types.NewOrderedMap()

    actions.EntityData.YListKeys = []string {}

    return &(actions.EntityData)
}

