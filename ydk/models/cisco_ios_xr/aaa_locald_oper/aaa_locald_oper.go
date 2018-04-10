// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-locald package operational data.
// 
// This module contains definitions
// for the following management objects:
//   aaa: AAA operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package aaa_locald_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_locald_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-locald-oper aaa}", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-locald-oper:aaa", reflect.TypeOf(Aaa{}))
}

// Aaa
// AAA operational data
type Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All tasks supported by system.
    AllTasks Aaa_AllTasks

    // Current user specific details.
    CurrentuserDetail Aaa_CurrentuserDetail

    // Task map of current user.
    TaskMap Aaa_TaskMap

    // Individual taskgroups container.
    Taskgroups Aaa_Taskgroups

    // Container for individual local user information.
    Users Aaa_Users

    // Container for individual password policy Information.
    PasswordPolicies Aaa_PasswordPolicies

    // Container for individual usergroup Information.
    Usergroups Aaa_Usergroups

    // Current users authentication method.
    AuthenMethod Aaa_AuthenMethod

    // Specific Usergroup Information.
    CurrentUsergroup Aaa_CurrentUsergroup

    // Diameter operational data.
    Diameter Aaa_Diameter

    // RADIUS operational data.
    Radius Aaa_Radius

    // TACACS operational data.
    Tacacs Aaa_Tacacs
}

func (aaa *Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-locald-oper"
    aaa.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-oper:aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = make(map[string]types.YChild)
    aaa.EntityData.Children["all-tasks"] = types.YChild{"AllTasks", &aaa.AllTasks}
    aaa.EntityData.Children["currentuser-detail"] = types.YChild{"CurrentuserDetail", &aaa.CurrentuserDetail}
    aaa.EntityData.Children["task-map"] = types.YChild{"TaskMap", &aaa.TaskMap}
    aaa.EntityData.Children["taskgroups"] = types.YChild{"Taskgroups", &aaa.Taskgroups}
    aaa.EntityData.Children["users"] = types.YChild{"Users", &aaa.Users}
    aaa.EntityData.Children["password-policies"] = types.YChild{"PasswordPolicies", &aaa.PasswordPolicies}
    aaa.EntityData.Children["usergroups"] = types.YChild{"Usergroups", &aaa.Usergroups}
    aaa.EntityData.Children["authen-method"] = types.YChild{"AuthenMethod", &aaa.AuthenMethod}
    aaa.EntityData.Children["current-usergroup"] = types.YChild{"CurrentUsergroup", &aaa.CurrentUsergroup}
    aaa.EntityData.Children["Cisco-IOS-XR-aaa-diameter-oper:diameter"] = types.YChild{"Diameter", &aaa.Diameter}
    aaa.EntityData.Children["Cisco-IOS-XR-aaa-protocol-radius-oper:radius"] = types.YChild{"Radius", &aaa.Radius}
    aaa.EntityData.Children["Cisco-IOS-XR-aaa-tacacs-oper:tacacs"] = types.YChild{"Tacacs", &aaa.Tacacs}
    aaa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aaa.EntityData)
}

// Aaa_AllTasks
// All tasks supported by system
type Aaa_AllTasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Names of available task-ids. The type is slice of string.
    TaskId []interface{}
}

func (allTasks *Aaa_AllTasks) GetEntityData() *types.CommonEntityData {
    allTasks.EntityData.YFilter = allTasks.YFilter
    allTasks.EntityData.YangName = "all-tasks"
    allTasks.EntityData.BundleName = "cisco_ios_xr"
    allTasks.EntityData.ParentYangName = "aaa"
    allTasks.EntityData.SegmentPath = "all-tasks"
    allTasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allTasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allTasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allTasks.EntityData.Children = make(map[string]types.YChild)
    allTasks.EntityData.Leafs = make(map[string]types.YLeaf)
    allTasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", allTasks.TaskId}
    return &(allTasks.EntityData)
}

// Aaa_CurrentuserDetail
// Current user specific details
type Aaa_CurrentuserDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is string.
    Name interface{}

    // Authentication method. The type is interface{} with range:
    // -2147483648..2147483647.
    Authenmethod interface{}

    // Component usergroups. The type is slice of string.
    Usergroup []interface{}

    // Task map details. The type is slice of string.
    Taskmap []interface{}
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetEntityData() *types.CommonEntityData {
    currentuserDetail.EntityData.YFilter = currentuserDetail.YFilter
    currentuserDetail.EntityData.YangName = "currentuser-detail"
    currentuserDetail.EntityData.BundleName = "cisco_ios_xr"
    currentuserDetail.EntityData.ParentYangName = "aaa"
    currentuserDetail.EntityData.SegmentPath = "currentuser-detail"
    currentuserDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentuserDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentuserDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentuserDetail.EntityData.Children = make(map[string]types.YChild)
    currentuserDetail.EntityData.Leafs = make(map[string]types.YLeaf)
    currentuserDetail.EntityData.Leafs["name"] = types.YLeaf{"Name", currentuserDetail.Name}
    currentuserDetail.EntityData.Leafs["authenmethod"] = types.YLeaf{"Authenmethod", currentuserDetail.Authenmethod}
    currentuserDetail.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", currentuserDetail.Usergroup}
    currentuserDetail.EntityData.Leafs["taskmap"] = types.YLeaf{"Taskmap", currentuserDetail.Taskmap}
    return &(currentuserDetail.EntityData)
}

// Aaa_TaskMap
// Task map of current user
type Aaa_TaskMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is string.
    Name interface{}

    // Authentication method. The type is interface{} with range:
    // -2147483648..2147483647.
    Authenmethod interface{}

    // Component usergroups. The type is slice of string.
    Usergroup []interface{}

    // Task map details. The type is slice of string.
    Taskmap []interface{}
}

func (taskMap *Aaa_TaskMap) GetEntityData() *types.CommonEntityData {
    taskMap.EntityData.YFilter = taskMap.YFilter
    taskMap.EntityData.YangName = "task-map"
    taskMap.EntityData.BundleName = "cisco_ios_xr"
    taskMap.EntityData.ParentYangName = "aaa"
    taskMap.EntityData.SegmentPath = "task-map"
    taskMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskMap.EntityData.Children = make(map[string]types.YChild)
    taskMap.EntityData.Leafs = make(map[string]types.YLeaf)
    taskMap.EntityData.Leafs["name"] = types.YLeaf{"Name", taskMap.Name}
    taskMap.EntityData.Leafs["authenmethod"] = types.YLeaf{"Authenmethod", taskMap.Authenmethod}
    taskMap.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", taskMap.Usergroup}
    taskMap.EntityData.Leafs["taskmap"] = types.YLeaf{"Taskmap", taskMap.Taskmap}
    return &(taskMap.EntityData)
}

// Aaa_Taskgroups
// Individual taskgroups container
type Aaa_Taskgroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific Taskgroup Information. The type is slice of
    // Aaa_Taskgroups_Taskgroup.
    Taskgroup []Aaa_Taskgroups_Taskgroup
}

func (taskgroups *Aaa_Taskgroups) GetEntityData() *types.CommonEntityData {
    taskgroups.EntityData.YFilter = taskgroups.YFilter
    taskgroups.EntityData.YangName = "taskgroups"
    taskgroups.EntityData.BundleName = "cisco_ios_xr"
    taskgroups.EntityData.ParentYangName = "aaa"
    taskgroups.EntityData.SegmentPath = "taskgroups"
    taskgroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroups.EntityData.Children = make(map[string]types.YChild)
    taskgroups.EntityData.Children["taskgroup"] = types.YChild{"Taskgroup", nil}
    for i := range taskgroups.Taskgroup {
        taskgroups.EntityData.Children[types.GetSegmentPath(&taskgroups.Taskgroup[i])] = types.YChild{"Taskgroup", &taskgroups.Taskgroup[i]}
    }
    taskgroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(taskgroups.EntityData)
}

// Aaa_Taskgroups_Taskgroup
// Specific Taskgroup Information
type Aaa_Taskgroups_Taskgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Taskgroup name. The type is string.
    Name interface{}

    // Name of the taskgroup. The type is string.
    NameXr interface{}

    // Task-ids included.
    IncludedTaskIds Aaa_Taskgroups_Taskgroup_IncludedTaskIds

    // Computed task map.
    TaskMap Aaa_Taskgroups_Taskgroup_TaskMap
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetEntityData() *types.CommonEntityData {
    taskgroup.EntityData.YFilter = taskgroup.YFilter
    taskgroup.EntityData.YangName = "taskgroup"
    taskgroup.EntityData.BundleName = "cisco_ios_xr"
    taskgroup.EntityData.ParentYangName = "taskgroups"
    taskgroup.EntityData.SegmentPath = "taskgroup" + "[name='" + fmt.Sprintf("%v", taskgroup.Name) + "']"
    taskgroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroup.EntityData.Children = make(map[string]types.YChild)
    taskgroup.EntityData.Children["included-task-ids"] = types.YChild{"IncludedTaskIds", &taskgroup.IncludedTaskIds}
    taskgroup.EntityData.Children["task-map"] = types.YChild{"TaskMap", &taskgroup.TaskMap}
    taskgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    taskgroup.EntityData.Leafs["name"] = types.YLeaf{"Name", taskgroup.Name}
    taskgroup.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", taskgroup.NameXr}
    return &(taskgroup.EntityData)
}

// Aaa_Taskgroups_Taskgroup_IncludedTaskIds
// Task-ids included
type Aaa_Taskgroups_Taskgroup_IncludedTaskIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks.
    Tasks []Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetEntityData() *types.CommonEntityData {
    includedTaskIds.EntityData.YFilter = includedTaskIds.YFilter
    includedTaskIds.EntityData.YangName = "included-task-ids"
    includedTaskIds.EntityData.BundleName = "cisco_ios_xr"
    includedTaskIds.EntityData.ParentYangName = "taskgroup"
    includedTaskIds.EntityData.SegmentPath = "included-task-ids"
    includedTaskIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includedTaskIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includedTaskIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includedTaskIds.EntityData.Children = make(map[string]types.YChild)
    includedTaskIds.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range includedTaskIds.Tasks {
        includedTaskIds.EntityData.Children[types.GetSegmentPath(&includedTaskIds.Tasks[i])] = types.YChild{"Tasks", &includedTaskIds.Tasks[i]}
    }
    includedTaskIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(includedTaskIds.EntityData)
}

// Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks
// List of permitted tasks
type Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "included-task-ids"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_Taskgroups_Taskgroup_TaskMap
// Computed task map
type Aaa_Taskgroups_Taskgroup_TaskMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Taskgroups_Taskgroup_TaskMap_Tasks.
    Tasks []Aaa_Taskgroups_Taskgroup_TaskMap_Tasks
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetEntityData() *types.CommonEntityData {
    taskMap.EntityData.YFilter = taskMap.YFilter
    taskMap.EntityData.YangName = "task-map"
    taskMap.EntityData.BundleName = "cisco_ios_xr"
    taskMap.EntityData.ParentYangName = "taskgroup"
    taskMap.EntityData.SegmentPath = "task-map"
    taskMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskMap.EntityData.Children = make(map[string]types.YChild)
    taskMap.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range taskMap.Tasks {
        taskMap.EntityData.Children[types.GetSegmentPath(&taskMap.Tasks[i])] = types.YChild{"Tasks", &taskMap.Tasks[i]}
    }
    taskMap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(taskMap.EntityData)
}

// Aaa_Taskgroups_Taskgroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Taskgroups_Taskgroup_TaskMap_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "task-map"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_Users
// Container for individual local user information
type Aaa_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific local user information. The type is slice of Aaa_Users_User.
    User []Aaa_Users_User
}

func (users *Aaa_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "aaa"
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

// Aaa_Users_User
// Specific local user information
type Aaa_Users_User struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Username. The type is string.
    Name interface{}

    // Username. The type is string.
    NameXr interface{}

    // Is admin plane user ?. The type is bool.
    AdminUser interface{}

    // Is first user ?. The type is bool.
    FirstUser interface{}

    // Member usergroups. The type is slice of string.
    Usergroup []interface{}

    // Computed taskmap.
    TaskMap Aaa_Users_User_TaskMap
}

func (user *Aaa_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + "[name='" + fmt.Sprintf("%v", user.Name) + "']"
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = make(map[string]types.YChild)
    user.EntityData.Children["task-map"] = types.YChild{"TaskMap", &user.TaskMap}
    user.EntityData.Leafs = make(map[string]types.YLeaf)
    user.EntityData.Leafs["name"] = types.YLeaf{"Name", user.Name}
    user.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", user.NameXr}
    user.EntityData.Leafs["admin-user"] = types.YLeaf{"AdminUser", user.AdminUser}
    user.EntityData.Leafs["first-user"] = types.YLeaf{"FirstUser", user.FirstUser}
    user.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", user.Usergroup}
    return &(user.EntityData)
}

// Aaa_Users_User_TaskMap
// Computed taskmap
type Aaa_Users_User_TaskMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of Aaa_Users_User_TaskMap_Tasks.
    Tasks []Aaa_Users_User_TaskMap_Tasks
}

func (taskMap *Aaa_Users_User_TaskMap) GetEntityData() *types.CommonEntityData {
    taskMap.EntityData.YFilter = taskMap.YFilter
    taskMap.EntityData.YangName = "task-map"
    taskMap.EntityData.BundleName = "cisco_ios_xr"
    taskMap.EntityData.ParentYangName = "user"
    taskMap.EntityData.SegmentPath = "task-map"
    taskMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskMap.EntityData.Children = make(map[string]types.YChild)
    taskMap.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range taskMap.Tasks {
        taskMap.EntityData.Children[types.GetSegmentPath(&taskMap.Tasks[i])] = types.YChild{"Tasks", &taskMap.Tasks[i]}
    }
    taskMap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(taskMap.EntityData)
}

// Aaa_Users_User_TaskMap_Tasks
// List of permitted tasks
type Aaa_Users_User_TaskMap_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "task-map"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_PasswordPolicies
// Container for individual password policy
// Information
type Aaa_PasswordPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password policy details. The type is slice of
    // Aaa_PasswordPolicies_PasswordPolicy.
    PasswordPolicy []Aaa_PasswordPolicies_PasswordPolicy
}

func (passwordPolicies *Aaa_PasswordPolicies) GetEntityData() *types.CommonEntityData {
    passwordPolicies.EntityData.YFilter = passwordPolicies.YFilter
    passwordPolicies.EntityData.YangName = "password-policies"
    passwordPolicies.EntityData.BundleName = "cisco_ios_xr"
    passwordPolicies.EntityData.ParentYangName = "aaa"
    passwordPolicies.EntityData.SegmentPath = "password-policies"
    passwordPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passwordPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passwordPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passwordPolicies.EntityData.Children = make(map[string]types.YChild)
    passwordPolicies.EntityData.Children["password-policy"] = types.YChild{"PasswordPolicy", nil}
    for i := range passwordPolicies.PasswordPolicy {
        passwordPolicies.EntityData.Children[types.GetSegmentPath(&passwordPolicies.PasswordPolicy[i])] = types.YChild{"PasswordPolicy", &passwordPolicies.PasswordPolicy[i]}
    }
    passwordPolicies.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(passwordPolicies.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy
// Password policy details
type Aaa_PasswordPolicies_PasswordPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Password policy name. The type is string.
    Name interface{}

    // Password Policy Name. The type is string.
    NameXr interface{}

    // Min Length. The type is interface{} with range: 0..255.
    MinLen interface{}

    // Max Length. The type is interface{} with range: 0..255.
    MaxLen interface{}

    // Special Character length. The type is interface{} with range: 0..255.
    SplChar interface{}

    // UpperCase Character length. The type is interface{} with range: 0..255.
    UpperCase interface{}

    // LowerCase Character length. The type is interface{} with range: 0..255.
    LowerCase interface{}

    // Numeric Character length. The type is interface{} with range: 0..255.
    Numeric interface{}

    // Number of different characters. The type is interface{} with range: 0..255.
    MinCharChange interface{}

    // Number of users with this policy. The type is interface{} with range:
    // 0..255.
    NumOfUsers interface{}

    // Maximum Failure Attempts allowed. The type is interface{} with range:
    // 0..4294967295.
    MaxFailAttempts interface{}

    // Count of users. The type is interface{} with range: 0..255.
    UsrCount interface{}

    // Error Count. The type is interface{} with range: 0..255.
    ErrCount interface{}

    // Lock Out Count. The type is interface{} with range: 0..255.
    LockOutCount interface{}

    // Lifetime of the policy.
    LifeTime Aaa_PasswordPolicies_PasswordPolicy_LifeTime

    // Lockout time of the policy.
    LockOutTime Aaa_PasswordPolicies_PasswordPolicy_LockOutTime
}

func (passwordPolicy *Aaa_PasswordPolicies_PasswordPolicy) GetEntityData() *types.CommonEntityData {
    passwordPolicy.EntityData.YFilter = passwordPolicy.YFilter
    passwordPolicy.EntityData.YangName = "password-policy"
    passwordPolicy.EntityData.BundleName = "cisco_ios_xr"
    passwordPolicy.EntityData.ParentYangName = "password-policies"
    passwordPolicy.EntityData.SegmentPath = "password-policy" + "[name='" + fmt.Sprintf("%v", passwordPolicy.Name) + "']"
    passwordPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passwordPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passwordPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passwordPolicy.EntityData.Children = make(map[string]types.YChild)
    passwordPolicy.EntityData.Children["life-time"] = types.YChild{"LifeTime", &passwordPolicy.LifeTime}
    passwordPolicy.EntityData.Children["lock-out-time"] = types.YChild{"LockOutTime", &passwordPolicy.LockOutTime}
    passwordPolicy.EntityData.Leafs = make(map[string]types.YLeaf)
    passwordPolicy.EntityData.Leafs["name"] = types.YLeaf{"Name", passwordPolicy.Name}
    passwordPolicy.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", passwordPolicy.NameXr}
    passwordPolicy.EntityData.Leafs["min-len"] = types.YLeaf{"MinLen", passwordPolicy.MinLen}
    passwordPolicy.EntityData.Leafs["max-len"] = types.YLeaf{"MaxLen", passwordPolicy.MaxLen}
    passwordPolicy.EntityData.Leafs["spl-char"] = types.YLeaf{"SplChar", passwordPolicy.SplChar}
    passwordPolicy.EntityData.Leafs["upper-case"] = types.YLeaf{"UpperCase", passwordPolicy.UpperCase}
    passwordPolicy.EntityData.Leafs["lower-case"] = types.YLeaf{"LowerCase", passwordPolicy.LowerCase}
    passwordPolicy.EntityData.Leafs["numeric"] = types.YLeaf{"Numeric", passwordPolicy.Numeric}
    passwordPolicy.EntityData.Leafs["min-char-change"] = types.YLeaf{"MinCharChange", passwordPolicy.MinCharChange}
    passwordPolicy.EntityData.Leafs["num-of-users"] = types.YLeaf{"NumOfUsers", passwordPolicy.NumOfUsers}
    passwordPolicy.EntityData.Leafs["max-fail-attempts"] = types.YLeaf{"MaxFailAttempts", passwordPolicy.MaxFailAttempts}
    passwordPolicy.EntityData.Leafs["usr-count"] = types.YLeaf{"UsrCount", passwordPolicy.UsrCount}
    passwordPolicy.EntityData.Leafs["err-count"] = types.YLeaf{"ErrCount", passwordPolicy.ErrCount}
    passwordPolicy.EntityData.Leafs["lock-out-count"] = types.YLeaf{"LockOutCount", passwordPolicy.LockOutCount}
    return &(passwordPolicy.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy_LifeTime
// Lifetime of the policy
type Aaa_PasswordPolicies_PasswordPolicy_LifeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // years. The type is interface{} with range: 0..255.
    Years interface{}

    // months. The type is interface{} with range: 0..255.
    Months interface{}

    // days. The type is interface{} with range: 0..255.
    Days interface{}

    // hours. The type is interface{} with range: 0..255.
    Hours interface{}

    // mins. The type is interface{} with range: 0..255.
    Mins interface{}

    // secs. The type is interface{} with range: 0..255.
    Secs interface{}
}

func (lifeTime *Aaa_PasswordPolicies_PasswordPolicy_LifeTime) GetEntityData() *types.CommonEntityData {
    lifeTime.EntityData.YFilter = lifeTime.YFilter
    lifeTime.EntityData.YangName = "life-time"
    lifeTime.EntityData.BundleName = "cisco_ios_xr"
    lifeTime.EntityData.ParentYangName = "password-policy"
    lifeTime.EntityData.SegmentPath = "life-time"
    lifeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lifeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lifeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lifeTime.EntityData.Children = make(map[string]types.YChild)
    lifeTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lifeTime.EntityData.Leafs["years"] = types.YLeaf{"Years", lifeTime.Years}
    lifeTime.EntityData.Leafs["months"] = types.YLeaf{"Months", lifeTime.Months}
    lifeTime.EntityData.Leafs["days"] = types.YLeaf{"Days", lifeTime.Days}
    lifeTime.EntityData.Leafs["hours"] = types.YLeaf{"Hours", lifeTime.Hours}
    lifeTime.EntityData.Leafs["mins"] = types.YLeaf{"Mins", lifeTime.Mins}
    lifeTime.EntityData.Leafs["secs"] = types.YLeaf{"Secs", lifeTime.Secs}
    return &(lifeTime.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy_LockOutTime
// Lockout time of the policy
type Aaa_PasswordPolicies_PasswordPolicy_LockOutTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // years. The type is interface{} with range: 0..255.
    Years interface{}

    // months. The type is interface{} with range: 0..255.
    Months interface{}

    // days. The type is interface{} with range: 0..255.
    Days interface{}

    // hours. The type is interface{} with range: 0..255.
    Hours interface{}

    // mins. The type is interface{} with range: 0..255.
    Mins interface{}

    // secs. The type is interface{} with range: 0..255.
    Secs interface{}
}

func (lockOutTime *Aaa_PasswordPolicies_PasswordPolicy_LockOutTime) GetEntityData() *types.CommonEntityData {
    lockOutTime.EntityData.YFilter = lockOutTime.YFilter
    lockOutTime.EntityData.YangName = "lock-out-time"
    lockOutTime.EntityData.BundleName = "cisco_ios_xr"
    lockOutTime.EntityData.ParentYangName = "password-policy"
    lockOutTime.EntityData.SegmentPath = "lock-out-time"
    lockOutTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lockOutTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lockOutTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lockOutTime.EntityData.Children = make(map[string]types.YChild)
    lockOutTime.EntityData.Leafs = make(map[string]types.YLeaf)
    lockOutTime.EntityData.Leafs["years"] = types.YLeaf{"Years", lockOutTime.Years}
    lockOutTime.EntityData.Leafs["months"] = types.YLeaf{"Months", lockOutTime.Months}
    lockOutTime.EntityData.Leafs["days"] = types.YLeaf{"Days", lockOutTime.Days}
    lockOutTime.EntityData.Leafs["hours"] = types.YLeaf{"Hours", lockOutTime.Hours}
    lockOutTime.EntityData.Leafs["mins"] = types.YLeaf{"Mins", lockOutTime.Mins}
    lockOutTime.EntityData.Leafs["secs"] = types.YLeaf{"Secs", lockOutTime.Secs}
    return &(lockOutTime.EntityData)
}

// Aaa_Usergroups
// Container for individual usergroup Information
type Aaa_Usergroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specific Usergroup Information. The type is slice of
    // Aaa_Usergroups_Usergroup.
    Usergroup []Aaa_Usergroups_Usergroup
}

func (usergroups *Aaa_Usergroups) GetEntityData() *types.CommonEntityData {
    usergroups.EntityData.YFilter = usergroups.YFilter
    usergroups.EntityData.YangName = "usergroups"
    usergroups.EntityData.BundleName = "cisco_ios_xr"
    usergroups.EntityData.ParentYangName = "aaa"
    usergroups.EntityData.SegmentPath = "usergroups"
    usergroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroups.EntityData.Children = make(map[string]types.YChild)
    usergroups.EntityData.Children["usergroup"] = types.YChild{"Usergroup", nil}
    for i := range usergroups.Usergroup {
        usergroups.EntityData.Children[types.GetSegmentPath(&usergroups.Usergroup[i])] = types.YChild{"Usergroup", &usergroups.Usergroup[i]}
    }
    usergroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usergroups.EntityData)
}

// Aaa_Usergroups_Usergroup
// Specific Usergroup Information
type Aaa_Usergroups_Usergroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Usergroup name. The type is string.
    Name interface{}

    // Name of the usergroup. The type is string.
    NameXr interface{}

    // Computed task map.
    TaskMap Aaa_Usergroups_Usergroup_TaskMap

    // Component taskgroups. The type is slice of
    // Aaa_Usergroups_Usergroup_Taskgroup.
    Taskgroup []Aaa_Usergroups_Usergroup_Taskgroup
}

func (usergroup *Aaa_Usergroups_Usergroup) GetEntityData() *types.CommonEntityData {
    usergroup.EntityData.YFilter = usergroup.YFilter
    usergroup.EntityData.YangName = "usergroup"
    usergroup.EntityData.BundleName = "cisco_ios_xr"
    usergroup.EntityData.ParentYangName = "usergroups"
    usergroup.EntityData.SegmentPath = "usergroup" + "[name='" + fmt.Sprintf("%v", usergroup.Name) + "']"
    usergroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroup.EntityData.Children = make(map[string]types.YChild)
    usergroup.EntityData.Children["task-map"] = types.YChild{"TaskMap", &usergroup.TaskMap}
    usergroup.EntityData.Children["taskgroup"] = types.YChild{"Taskgroup", nil}
    for i := range usergroup.Taskgroup {
        usergroup.EntityData.Children[types.GetSegmentPath(&usergroup.Taskgroup[i])] = types.YChild{"Taskgroup", &usergroup.Taskgroup[i]}
    }
    usergroup.EntityData.Leafs = make(map[string]types.YLeaf)
    usergroup.EntityData.Leafs["name"] = types.YLeaf{"Name", usergroup.Name}
    usergroup.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", usergroup.NameXr}
    return &(usergroup.EntityData)
}

// Aaa_Usergroups_Usergroup_TaskMap
// Computed task map
type Aaa_Usergroups_Usergroup_TaskMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_TaskMap_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_TaskMap_Tasks
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetEntityData() *types.CommonEntityData {
    taskMap.EntityData.YFilter = taskMap.YFilter
    taskMap.EntityData.YangName = "task-map"
    taskMap.EntityData.BundleName = "cisco_ios_xr"
    taskMap.EntityData.ParentYangName = "usergroup"
    taskMap.EntityData.SegmentPath = "task-map"
    taskMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskMap.EntityData.Children = make(map[string]types.YChild)
    taskMap.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range taskMap.Tasks {
        taskMap.EntityData.Children[types.GetSegmentPath(&taskMap.Tasks[i])] = types.YChild{"Tasks", &taskMap.Tasks[i]}
    }
    taskMap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(taskMap.EntityData)
}

// Aaa_Usergroups_Usergroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_TaskMap_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "task-map"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_Usergroups_Usergroup_Taskgroup
// Component taskgroups
type Aaa_Usergroups_Usergroup_Taskgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the taskgroup. The type is string.
    NameXr interface{}

    // Task-ids included.
    IncludedTaskIds Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds

    // Computed task map.
    TaskMap Aaa_Usergroups_Usergroup_Taskgroup_TaskMap
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetEntityData() *types.CommonEntityData {
    taskgroup.EntityData.YFilter = taskgroup.YFilter
    taskgroup.EntityData.YangName = "taskgroup"
    taskgroup.EntityData.BundleName = "cisco_ios_xr"
    taskgroup.EntityData.ParentYangName = "usergroup"
    taskgroup.EntityData.SegmentPath = "taskgroup"
    taskgroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroup.EntityData.Children = make(map[string]types.YChild)
    taskgroup.EntityData.Children["included-task-ids"] = types.YChild{"IncludedTaskIds", &taskgroup.IncludedTaskIds}
    taskgroup.EntityData.Children["task-map"] = types.YChild{"TaskMap", &taskgroup.TaskMap}
    taskgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    taskgroup.EntityData.Leafs["name-xr"] = types.YLeaf{"NameXr", taskgroup.NameXr}
    return &(taskgroup.EntityData)
}

// Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds
// Task-ids included
type Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetEntityData() *types.CommonEntityData {
    includedTaskIds.EntityData.YFilter = includedTaskIds.YFilter
    includedTaskIds.EntityData.YangName = "included-task-ids"
    includedTaskIds.EntityData.BundleName = "cisco_ios_xr"
    includedTaskIds.EntityData.ParentYangName = "taskgroup"
    includedTaskIds.EntityData.SegmentPath = "included-task-ids"
    includedTaskIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includedTaskIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includedTaskIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includedTaskIds.EntityData.Children = make(map[string]types.YChild)
    includedTaskIds.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range includedTaskIds.Tasks {
        includedTaskIds.EntityData.Children[types.GetSegmentPath(&includedTaskIds.Tasks[i])] = types.YChild{"Tasks", &includedTaskIds.Tasks[i]}
    }
    includedTaskIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(includedTaskIds.EntityData)
}

// Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "included-task-ids"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_Usergroups_Usergroup_Taskgroup_TaskMap
// Computed task map
type Aaa_Usergroups_Usergroup_Taskgroup_TaskMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetEntityData() *types.CommonEntityData {
    taskMap.EntityData.YFilter = taskMap.YFilter
    taskMap.EntityData.YangName = "task-map"
    taskMap.EntityData.BundleName = "cisco_ios_xr"
    taskMap.EntityData.ParentYangName = "taskgroup"
    taskMap.EntityData.SegmentPath = "task-map"
    taskMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskMap.EntityData.Children = make(map[string]types.YChild)
    taskMap.EntityData.Children["tasks"] = types.YChild{"Tasks", nil}
    for i := range taskMap.Tasks {
        taskMap.EntityData.Children[types.GetSegmentPath(&taskMap.Tasks[i])] = types.YChild{"Tasks", &taskMap.Tasks[i]}
    }
    taskMap.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(taskMap.EntityData)
}

// Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task-id. The type is string.
    TaskId interface{}

    // Is read permitted?. The type is bool.
    Read interface{}

    // Is write permitted?. The type is bool.
    Write interface{}

    // Is execute permitted?. The type is bool.
    Execute interface{}

    // Is debug permitted?. The type is bool.
    Debug interface{}
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "task-map"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = make(map[string]types.YChild)
    tasks.EntityData.Leafs = make(map[string]types.YLeaf)
    tasks.EntityData.Leafs["task-id"] = types.YLeaf{"TaskId", tasks.TaskId}
    tasks.EntityData.Leafs["read"] = types.YLeaf{"Read", tasks.Read}
    tasks.EntityData.Leafs["write"] = types.YLeaf{"Write", tasks.Write}
    tasks.EntityData.Leafs["execute"] = types.YLeaf{"Execute", tasks.Execute}
    tasks.EntityData.Leafs["debug"] = types.YLeaf{"Debug", tasks.Debug}
    return &(tasks.EntityData)
}

// Aaa_AuthenMethod
// Current users authentication method
type Aaa_AuthenMethod struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is string.
    Name interface{}

    // Authentication method. The type is interface{} with range:
    // -2147483648..2147483647.
    Authenmethod interface{}

    // Component usergroups. The type is slice of string.
    Usergroup []interface{}

    // Task map details. The type is slice of string.
    Taskmap []interface{}
}

func (authenMethod *Aaa_AuthenMethod) GetEntityData() *types.CommonEntityData {
    authenMethod.EntityData.YFilter = authenMethod.YFilter
    authenMethod.EntityData.YangName = "authen-method"
    authenMethod.EntityData.BundleName = "cisco_ios_xr"
    authenMethod.EntityData.ParentYangName = "aaa"
    authenMethod.EntityData.SegmentPath = "authen-method"
    authenMethod.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenMethod.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenMethod.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenMethod.EntityData.Children = make(map[string]types.YChild)
    authenMethod.EntityData.Leafs = make(map[string]types.YLeaf)
    authenMethod.EntityData.Leafs["name"] = types.YLeaf{"Name", authenMethod.Name}
    authenMethod.EntityData.Leafs["authenmethod"] = types.YLeaf{"Authenmethod", authenMethod.Authenmethod}
    authenMethod.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", authenMethod.Usergroup}
    authenMethod.EntityData.Leafs["taskmap"] = types.YLeaf{"Taskmap", authenMethod.Taskmap}
    return &(authenMethod.EntityData)
}

// Aaa_CurrentUsergroup
// Specific Usergroup Information
type Aaa_CurrentUsergroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the usergroup. The type is string.
    Name interface{}

    // Authentication method. The type is interface{} with range:
    // -2147483648..2147483647.
    Authenmethod interface{}

    // Component usergroups. The type is slice of string.
    Usergroup []interface{}

    // Task map details. The type is slice of string.
    Taskmap []interface{}
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetEntityData() *types.CommonEntityData {
    currentUsergroup.EntityData.YFilter = currentUsergroup.YFilter
    currentUsergroup.EntityData.YangName = "current-usergroup"
    currentUsergroup.EntityData.BundleName = "cisco_ios_xr"
    currentUsergroup.EntityData.ParentYangName = "aaa"
    currentUsergroup.EntityData.SegmentPath = "current-usergroup"
    currentUsergroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    currentUsergroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    currentUsergroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    currentUsergroup.EntityData.Children = make(map[string]types.YChild)
    currentUsergroup.EntityData.Leafs = make(map[string]types.YLeaf)
    currentUsergroup.EntityData.Leafs["name"] = types.YLeaf{"Name", currentUsergroup.Name}
    currentUsergroup.EntityData.Leafs["authenmethod"] = types.YLeaf{"Authenmethod", currentUsergroup.Authenmethod}
    currentUsergroup.EntityData.Leafs["usergroup"] = types.YLeaf{"Usergroup", currentUsergroup.Usergroup}
    currentUsergroup.EntityData.Leafs["taskmap"] = types.YLeaf{"Taskmap", currentUsergroup.Taskmap}
    return &(currentUsergroup.EntityData)
}

// Aaa_Diameter
// Diameter operational data
type Aaa_Diameter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diameter global gy data.
    Gy Aaa_Diameter_Gy

    // Diameter Gx Statistics data.
    GxStatistics Aaa_Diameter_GxStatistics

    // Diameter global gx data.
    Gx Aaa_Diameter_Gx

    // Diameter peer global data.
    Peers Aaa_Diameter_Peers

    // Diameter NAS data.
    Nas Aaa_Diameter_Nas

    // Diameter NAS summary.
    NasSummary Aaa_Diameter_NasSummary

    // Diameter Gy Session data list.
    GySessionIds Aaa_Diameter_GySessionIds

    // Diameter Gy Statistics data.
    GyStatistics Aaa_Diameter_GyStatistics

    // Diameter Gx Session data list.
    GxSessionIds Aaa_Diameter_GxSessionIds

    // Diameter Nas Session data.
    NasSession Aaa_Diameter_NasSession
}

func (diameter *Aaa_Diameter) GetEntityData() *types.CommonEntityData {
    diameter.EntityData.YFilter = diameter.YFilter
    diameter.EntityData.YangName = "diameter"
    diameter.EntityData.BundleName = "cisco_ios_xr"
    diameter.EntityData.ParentYangName = "aaa"
    diameter.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-diameter-oper:diameter"
    diameter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameter.EntityData.Children = make(map[string]types.YChild)
    diameter.EntityData.Children["gy"] = types.YChild{"Gy", &diameter.Gy}
    diameter.EntityData.Children["gx-statistics"] = types.YChild{"GxStatistics", &diameter.GxStatistics}
    diameter.EntityData.Children["gx"] = types.YChild{"Gx", &diameter.Gx}
    diameter.EntityData.Children["peers"] = types.YChild{"Peers", &diameter.Peers}
    diameter.EntityData.Children["nas"] = types.YChild{"Nas", &diameter.Nas}
    diameter.EntityData.Children["nas-summary"] = types.YChild{"NasSummary", &diameter.NasSummary}
    diameter.EntityData.Children["gy-session-ids"] = types.YChild{"GySessionIds", &diameter.GySessionIds}
    diameter.EntityData.Children["gy-statistics"] = types.YChild{"GyStatistics", &diameter.GyStatistics}
    diameter.EntityData.Children["gx-session-ids"] = types.YChild{"GxSessionIds", &diameter.GxSessionIds}
    diameter.EntityData.Children["nas-session"] = types.YChild{"NasSession", &diameter.NasSession}
    diameter.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(diameter.EntityData)
}

// Aaa_Diameter_Gy
// Diameter global gy data
type Aaa_Diameter_Gy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Gy state. The type is bool.
    IsEnabled interface{}

    // Gy transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TxTimer interface{}

    // Gy retransmit count. The type is interface{} with range: 0..4294967295.
    Retransmits interface{}
}

func (gy *Aaa_Diameter_Gy) GetEntityData() *types.CommonEntityData {
    gy.EntityData.YFilter = gy.YFilter
    gy.EntityData.YangName = "gy"
    gy.EntityData.BundleName = "cisco_ios_xr"
    gy.EntityData.ParentYangName = "diameter"
    gy.EntityData.SegmentPath = "gy"
    gy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gy.EntityData.Children = make(map[string]types.YChild)
    gy.EntityData.Leafs = make(map[string]types.YLeaf)
    gy.EntityData.Leafs["is-enabled"] = types.YLeaf{"IsEnabled", gy.IsEnabled}
    gy.EntityData.Leafs["tx-timer"] = types.YLeaf{"TxTimer", gy.TxTimer}
    gy.EntityData.Leafs["retransmits"] = types.YLeaf{"Retransmits", gy.Retransmits}
    return &(gy.EntityData)
}

// Aaa_Diameter_GxStatistics
// Diameter Gx Statistics data
type Aaa_Diameter_GxStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CCR Initial Messages. The type is interface{} with range: 0..4294967295.
    CcrInitMessages interface{}

    // CCR Initial Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrInitFailedMessages interface{}

    // CCR Initial Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrInitTimedOutMessages interface{}

    // CCR Initial Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrInitRetryMessages interface{}

    // CCR Update Messages. The type is interface{} with range: 0..4294967295.
    CcrUpdateMessages interface{}

    // CCR Update Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateFailedMessages interface{}

    // CCR Update Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateTimedOutMessages interface{}

    // CCR Update Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateRetryMessages interface{}

    // CCR Final Messages. The type is interface{} with range: 0..4294967295.
    CcrFinalMessages interface{}

    // CCR Final Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalFailedMessages interface{}

    // CCR Final Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalTimedOutMessages interface{}

    // CCR Final Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalRetryMessages interface{}

    // CCA Initial Messages. The type is interface{} with range: 0..4294967295.
    CcaInitMessages interface{}

    // CCA Initial Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaInitErrorMessages interface{}

    // CCA Update Messages. The type is interface{} with range: 0..4294967295.
    CcaUpdateMessages interface{}

    // CCA Update Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaUpdateErrorMessages interface{}

    // CCA Final Messages. The type is interface{} with range: 0..4294967295.
    CcaFinalMessages interface{}

    // CCA Final Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaFinalErrorMessages interface{}

    // RAR Received Messages. The type is interface{} with range: 0..4294967295.
    RarReceivedMessages interface{}

    // RAR Received Messages Error. The type is interface{} with range:
    // 0..4294967295.
    RarReceivedErrorMessages interface{}

    // RAA Sent Messages. The type is interface{} with range: 0..4294967295.
    RaaSentMessages interface{}

    // RAA Sent Messages Error. The type is interface{} with range: 0..4294967295.
    RaaSentErrorMessages interface{}

    // ASR Received Messages. The type is interface{} with range: 0..4294967295.
    AsrReceivedMessages interface{}

    // ASR Received Messages Error. The type is interface{} with range:
    // 0..4294967295.
    AsrReceivedErrorMessages interface{}

    // ASA Sent Messages. The type is interface{} with range: 0..4294967295.
    AsaSentMesssages interface{}

    // ASA Sent Messages Error. The type is interface{} with range: 0..4294967295.
    AsaSentErrorMessages interface{}

    // Session Termination from server. The type is interface{} with range:
    // 0..4294967295.
    SessionTerminationMessages interface{}

    // Unknown Request Messages. The type is interface{} with range:
    // 0..4294967295.
    UnknownRequestMessages interface{}

    // Restore Sessions. The type is interface{} with range: 0..4294967295.
    RestoreSessions interface{}

    // Total Opened Sessions. The type is interface{} with range: 0..4294967295.
    OpenSessions interface{}

    // Total Closed Sessions. The type is interface{} with range: 0..4294967295.
    CloseSessions interface{}

    // Total Active Sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetEntityData() *types.CommonEntityData {
    gxStatistics.EntityData.YFilter = gxStatistics.YFilter
    gxStatistics.EntityData.YangName = "gx-statistics"
    gxStatistics.EntityData.BundleName = "cisco_ios_xr"
    gxStatistics.EntityData.ParentYangName = "diameter"
    gxStatistics.EntityData.SegmentPath = "gx-statistics"
    gxStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gxStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gxStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gxStatistics.EntityData.Children = make(map[string]types.YChild)
    gxStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    gxStatistics.EntityData.Leafs["ccr-init-messages"] = types.YLeaf{"CcrInitMessages", gxStatistics.CcrInitMessages}
    gxStatistics.EntityData.Leafs["ccr-init-failed-messages"] = types.YLeaf{"CcrInitFailedMessages", gxStatistics.CcrInitFailedMessages}
    gxStatistics.EntityData.Leafs["ccr-init-timed-out-messages"] = types.YLeaf{"CcrInitTimedOutMessages", gxStatistics.CcrInitTimedOutMessages}
    gxStatistics.EntityData.Leafs["ccr-init-retry-messages"] = types.YLeaf{"CcrInitRetryMessages", gxStatistics.CcrInitRetryMessages}
    gxStatistics.EntityData.Leafs["ccr-update-messages"] = types.YLeaf{"CcrUpdateMessages", gxStatistics.CcrUpdateMessages}
    gxStatistics.EntityData.Leafs["ccr-update-failed-messages"] = types.YLeaf{"CcrUpdateFailedMessages", gxStatistics.CcrUpdateFailedMessages}
    gxStatistics.EntityData.Leafs["ccr-update-timed-out-messages"] = types.YLeaf{"CcrUpdateTimedOutMessages", gxStatistics.CcrUpdateTimedOutMessages}
    gxStatistics.EntityData.Leafs["ccr-update-retry-messages"] = types.YLeaf{"CcrUpdateRetryMessages", gxStatistics.CcrUpdateRetryMessages}
    gxStatistics.EntityData.Leafs["ccr-final-messages"] = types.YLeaf{"CcrFinalMessages", gxStatistics.CcrFinalMessages}
    gxStatistics.EntityData.Leafs["ccr-final-failed-messages"] = types.YLeaf{"CcrFinalFailedMessages", gxStatistics.CcrFinalFailedMessages}
    gxStatistics.EntityData.Leafs["ccr-final-timed-out-messages"] = types.YLeaf{"CcrFinalTimedOutMessages", gxStatistics.CcrFinalTimedOutMessages}
    gxStatistics.EntityData.Leafs["ccr-final-retry-messages"] = types.YLeaf{"CcrFinalRetryMessages", gxStatistics.CcrFinalRetryMessages}
    gxStatistics.EntityData.Leafs["cca-init-messages"] = types.YLeaf{"CcaInitMessages", gxStatistics.CcaInitMessages}
    gxStatistics.EntityData.Leafs["cca-init-error-messages"] = types.YLeaf{"CcaInitErrorMessages", gxStatistics.CcaInitErrorMessages}
    gxStatistics.EntityData.Leafs["cca-update-messages"] = types.YLeaf{"CcaUpdateMessages", gxStatistics.CcaUpdateMessages}
    gxStatistics.EntityData.Leafs["cca-update-error-messages"] = types.YLeaf{"CcaUpdateErrorMessages", gxStatistics.CcaUpdateErrorMessages}
    gxStatistics.EntityData.Leafs["cca-final-messages"] = types.YLeaf{"CcaFinalMessages", gxStatistics.CcaFinalMessages}
    gxStatistics.EntityData.Leafs["cca-final-error-messages"] = types.YLeaf{"CcaFinalErrorMessages", gxStatistics.CcaFinalErrorMessages}
    gxStatistics.EntityData.Leafs["rar-received-messages"] = types.YLeaf{"RarReceivedMessages", gxStatistics.RarReceivedMessages}
    gxStatistics.EntityData.Leafs["rar-received-error-messages"] = types.YLeaf{"RarReceivedErrorMessages", gxStatistics.RarReceivedErrorMessages}
    gxStatistics.EntityData.Leafs["raa-sent-messages"] = types.YLeaf{"RaaSentMessages", gxStatistics.RaaSentMessages}
    gxStatistics.EntityData.Leafs["raa-sent-error-messages"] = types.YLeaf{"RaaSentErrorMessages", gxStatistics.RaaSentErrorMessages}
    gxStatistics.EntityData.Leafs["asr-received-messages"] = types.YLeaf{"AsrReceivedMessages", gxStatistics.AsrReceivedMessages}
    gxStatistics.EntityData.Leafs["asr-received-error-messages"] = types.YLeaf{"AsrReceivedErrorMessages", gxStatistics.AsrReceivedErrorMessages}
    gxStatistics.EntityData.Leafs["asa-sent-messsages"] = types.YLeaf{"AsaSentMesssages", gxStatistics.AsaSentMesssages}
    gxStatistics.EntityData.Leafs["asa-sent-error-messages"] = types.YLeaf{"AsaSentErrorMessages", gxStatistics.AsaSentErrorMessages}
    gxStatistics.EntityData.Leafs["session-termination-messages"] = types.YLeaf{"SessionTerminationMessages", gxStatistics.SessionTerminationMessages}
    gxStatistics.EntityData.Leafs["unknown-request-messages"] = types.YLeaf{"UnknownRequestMessages", gxStatistics.UnknownRequestMessages}
    gxStatistics.EntityData.Leafs["restore-sessions"] = types.YLeaf{"RestoreSessions", gxStatistics.RestoreSessions}
    gxStatistics.EntityData.Leafs["open-sessions"] = types.YLeaf{"OpenSessions", gxStatistics.OpenSessions}
    gxStatistics.EntityData.Leafs["close-sessions"] = types.YLeaf{"CloseSessions", gxStatistics.CloseSessions}
    gxStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", gxStatistics.ActiveSessions}
    return &(gxStatistics.EntityData)
}

// Aaa_Diameter_Gx
// Diameter global gx data
type Aaa_Diameter_Gx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Gx state. The type is bool.
    IsEnabled interface{}

    // Gx transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TxTimer interface{}

    // Gx retransmit count. The type is interface{} with range: 0..4294967295.
    Retransmits interface{}
}

func (gx *Aaa_Diameter_Gx) GetEntityData() *types.CommonEntityData {
    gx.EntityData.YFilter = gx.YFilter
    gx.EntityData.YangName = "gx"
    gx.EntityData.BundleName = "cisco_ios_xr"
    gx.EntityData.ParentYangName = "diameter"
    gx.EntityData.SegmentPath = "gx"
    gx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gx.EntityData.Children = make(map[string]types.YChild)
    gx.EntityData.Leafs = make(map[string]types.YLeaf)
    gx.EntityData.Leafs["is-enabled"] = types.YLeaf{"IsEnabled", gx.IsEnabled}
    gx.EntityData.Leafs["tx-timer"] = types.YLeaf{"TxTimer", gx.TxTimer}
    gx.EntityData.Leafs["retransmits"] = types.YLeaf{"Retransmits", gx.Retransmits}
    return &(gx.EntityData)
}

// Aaa_Diameter_Peers
// Diameter peer global data
type Aaa_Diameter_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Origin Host. The type is string.
    OriginHost interface{}

    // Origin Realm. The type is string.
    OriginRealm interface{}

    // Source Interface. The type is string.
    SourceInterface interface{}

    // TLS Trustpoint. The type is string.
    TlsTrustpoint interface{}

    // Connection retry timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ConnRetryTimer interface{}

    // Watch dog timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    WatchdogTimer interface{}

    // Transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TransactionTimer interface{}

    // Total number of transactions. The type is interface{} with range:
    // 0..4294967295.
    TransTotal interface{}

    // Maximum number of transactions. The type is interface{} with range:
    // 0..4294967295.
    TransMax interface{}

    // Peer List. The type is slice of Aaa_Diameter_Peers_Peer.
    Peer []Aaa_Diameter_Peers_Peer
}

func (peers *Aaa_Diameter_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "diameter"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = make(map[string]types.YChild)
    peers.EntityData.Children["peer"] = types.YChild{"Peer", nil}
    for i := range peers.Peer {
        peers.EntityData.Children[types.GetSegmentPath(&peers.Peer[i])] = types.YChild{"Peer", &peers.Peer[i]}
    }
    peers.EntityData.Leafs = make(map[string]types.YLeaf)
    peers.EntityData.Leafs["origin-host"] = types.YLeaf{"OriginHost", peers.OriginHost}
    peers.EntityData.Leafs["origin-realm"] = types.YLeaf{"OriginRealm", peers.OriginRealm}
    peers.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", peers.SourceInterface}
    peers.EntityData.Leafs["tls-trustpoint"] = types.YLeaf{"TlsTrustpoint", peers.TlsTrustpoint}
    peers.EntityData.Leafs["conn-retry-timer"] = types.YLeaf{"ConnRetryTimer", peers.ConnRetryTimer}
    peers.EntityData.Leafs["watchdog-timer"] = types.YLeaf{"WatchdogTimer", peers.WatchdogTimer}
    peers.EntityData.Leafs["transaction-timer"] = types.YLeaf{"TransactionTimer", peers.TransactionTimer}
    peers.EntityData.Leafs["trans-total"] = types.YLeaf{"TransTotal", peers.TransTotal}
    peers.EntityData.Leafs["trans-max"] = types.YLeaf{"TransMax", peers.TransMax}
    return &(peers.EntityData)
}

// Aaa_Diameter_Peers_Peer
// Peer List
type Aaa_Diameter_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer Name. The type is string.
    PeerName interface{}

    // Peer Index. The type is interface{} with range: 0..4294967295.
    PeerIndex interface{}

    // IPv4 or IPv6 address of DIAMETER peer. The type is string.
    Address interface{}

    // Port number on which the peeris running. The type is interface{} with
    // range: 0..4294967295.
    Port interface{}

    // Local Connection port. The type is interface{} with range: 0..4294967295.
    PortConnect interface{}

    // Protocol Type. The type is ProtocolTypeValue.
    ProtocolType interface{}

    // Security type used to transport. The type is SecurityTypeValue.
    SecurityType interface{}

    // Connection retry timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    ConnRetryTimer interface{}

    // Watch dog timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    WatchdogTimer interface{}

    // Transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TransactionTimer interface{}

    // Vrf Name. The type is string.
    VrfName interface{}

    // Source Interface. The type is string.
    SourceInterface interface{}

    // Destination host name. The type is string.
    DestinationHost interface{}

    // Destination realm. The type is string.
    DestinationRealm interface{}

    // Peer Type. The type is Peer.
    PeerType interface{}

    // Firmware revision. The type is interface{} with range: 0..4294967295.
    FirmwareRevision interface{}

    // State Duration in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    StateDuration interface{}

    // Last Disconnect Reason. The type is DisconnectCause.
    LastDisconnectCause interface{}

    // Who Initiated Disconnect. The type is WhoInitiatedDisconnect.
    WhoInitDisconnect interface{}

    // Incoming ASRs. The type is interface{} with range: 0..4294967295.
    InAsRs interface{}

    // Outgoing ASRs. The type is interface{} with range: 0..4294967295.
    OutAsRs interface{}

    // Incoming ASAs. The type is interface{} with range: 0..4294967295.
    InAsAs interface{}

    // Outgoing ASAs. The type is interface{} with range: 0..4294967295.
    OutAsAs interface{}

    // Incoming ACRs. The type is interface{} with range: 0..4294967295.
    InAcRs interface{}

    // Outgoing ACRs. The type is interface{} with range: 0..4294967295.
    OutAcRs interface{}

    // Incoming ACAs. The type is interface{} with range: 0..4294967295.
    InAcAs interface{}

    // Outgoing ACAs. The type is interface{} with range: 0..4294967295.
    OutAcAs interface{}

    // Incoming CERs. The type is interface{} with range: 0..4294967295.
    InCeRs interface{}

    // Outgoing CERs. The type is interface{} with range: 0..4294967295.
    OutCeRs interface{}

    // Incoming CEAs. The type is interface{} with range: 0..4294967295.
    InCeAs interface{}

    // Outgoing CEAs. The type is interface{} with range: 0..4294967295.
    OutCeAs interface{}

    // Incoming DWRs. The type is interface{} with range: 0..4294967295.
    InDwRs interface{}

    // Outgoing DWRs. The type is interface{} with range: 0..4294967295.
    OutDwRs interface{}

    // Incoming DWAs. The type is interface{} with range: 0..4294967295.
    InDwAs interface{}

    // Outgoing DWAs. The type is interface{} with range: 0..4294967295.
    OutDwAs interface{}

    // Incoming DPRs. The type is interface{} with range: 0..4294967295.
    InDpRs interface{}

    // Outgoing DPRs. The type is interface{} with range: 0..4294967295.
    OutDpRs interface{}

    // Incoming DPAs. The type is interface{} with range: 0..4294967295.
    InDpAs interface{}

    // Outgoing DPAs. The type is interface{} with range: 0..4294967295.
    OutDpAs interface{}

    // Incoming RARs. The type is interface{} with range: 0..4294967295.
    InRaRs interface{}

    // Outgoing RARs. The type is interface{} with range: 0..4294967295.
    OutRaRs interface{}

    // Incoming RAAs. The type is interface{} with range: 0..4294967295.
    InRaAs interface{}

    // Outgoing RAAs. The type is interface{} with range: 0..4294967295.
    OutRaAs interface{}

    // Incoming STRs. The type is interface{} with range: 0..4294967295.
    InStRs interface{}

    // Outgoing STRs. The type is interface{} with range: 0..4294967295.
    OutStRs interface{}

    // Incoming STAs. The type is interface{} with range: 0..4294967295.
    InStAs interface{}

    // Outgoing STAs. The type is interface{} with range: 0..4294967295.
    OutStAs interface{}

    // Incoming CCRs. The type is interface{} with range: 0..4294967295.
    InCcRs interface{}

    // Outgoing CCRs. The type is interface{} with range: 0..4294967295.
    OutCcRs interface{}

    // Incoming CCAs. The type is interface{} with range: 0..4294967295.
    InCcAs interface{}

    // Outgoing CCAs. The type is interface{} with range: 0..4294967295.
    OutCcAs interface{}

    // Outgoing AARs. The type is interface{} with range: 0..4294967295.
    OutAaRs interface{}

    // Incoming AAAs. The type is interface{} with range: 0..4294967295.
    InAaAs interface{}

    // Malformed Requests. The type is interface{} with range: 0..4294967295.
    MalformedRequests interface{}

    // Protocol Error Received. The type is interface{} with range: 0..4294967295.
    ReceivedProtoErrors interface{}

    // Protocol Error Sent. The type is interface{} with range: 0..4294967295.
    SentProtoErrors interface{}

    // Transient failures Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedTransientFails interface{}

    // Transient failures Sent. The type is interface{} with range: 0..4294967295.
    SentTransientFails interface{}

    // Permanent Failures Received. The type is interface{} with range:
    // 0..4294967295.
    ReceivedPermanentFails interface{}

    // Permanent Failures Sent. The type is interface{} with range: 0..4294967295.
    SentPermanentFails interface{}

    // Transport Down. The type is interface{} with range: 0..4294967295.
    TransportDown interface{}

    // Peer Connection Status. The type is PeerStateValue.
    State interface{}
}

func (peer *Aaa_Diameter_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer"
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = make(map[string]types.YChild)
    peer.EntityData.Leafs = make(map[string]types.YLeaf)
    peer.EntityData.Leafs["peer-name"] = types.YLeaf{"PeerName", peer.PeerName}
    peer.EntityData.Leafs["peer-index"] = types.YLeaf{"PeerIndex", peer.PeerIndex}
    peer.EntityData.Leafs["address"] = types.YLeaf{"Address", peer.Address}
    peer.EntityData.Leafs["port"] = types.YLeaf{"Port", peer.Port}
    peer.EntityData.Leafs["port-connect"] = types.YLeaf{"PortConnect", peer.PortConnect}
    peer.EntityData.Leafs["protocol-type"] = types.YLeaf{"ProtocolType", peer.ProtocolType}
    peer.EntityData.Leafs["security-type"] = types.YLeaf{"SecurityType", peer.SecurityType}
    peer.EntityData.Leafs["conn-retry-timer"] = types.YLeaf{"ConnRetryTimer", peer.ConnRetryTimer}
    peer.EntityData.Leafs["watchdog-timer"] = types.YLeaf{"WatchdogTimer", peer.WatchdogTimer}
    peer.EntityData.Leafs["transaction-timer"] = types.YLeaf{"TransactionTimer", peer.TransactionTimer}
    peer.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", peer.VrfName}
    peer.EntityData.Leafs["source-interface"] = types.YLeaf{"SourceInterface", peer.SourceInterface}
    peer.EntityData.Leafs["destination-host"] = types.YLeaf{"DestinationHost", peer.DestinationHost}
    peer.EntityData.Leafs["destination-realm"] = types.YLeaf{"DestinationRealm", peer.DestinationRealm}
    peer.EntityData.Leafs["peer-type"] = types.YLeaf{"PeerType", peer.PeerType}
    peer.EntityData.Leafs["firmware-revision"] = types.YLeaf{"FirmwareRevision", peer.FirmwareRevision}
    peer.EntityData.Leafs["state-duration"] = types.YLeaf{"StateDuration", peer.StateDuration}
    peer.EntityData.Leafs["last-disconnect-cause"] = types.YLeaf{"LastDisconnectCause", peer.LastDisconnectCause}
    peer.EntityData.Leafs["who-init-disconnect"] = types.YLeaf{"WhoInitDisconnect", peer.WhoInitDisconnect}
    peer.EntityData.Leafs["in-as-rs"] = types.YLeaf{"InAsRs", peer.InAsRs}
    peer.EntityData.Leafs["out-as-rs"] = types.YLeaf{"OutAsRs", peer.OutAsRs}
    peer.EntityData.Leafs["in-as-as"] = types.YLeaf{"InAsAs", peer.InAsAs}
    peer.EntityData.Leafs["out-as-as"] = types.YLeaf{"OutAsAs", peer.OutAsAs}
    peer.EntityData.Leafs["in-ac-rs"] = types.YLeaf{"InAcRs", peer.InAcRs}
    peer.EntityData.Leafs["out-ac-rs"] = types.YLeaf{"OutAcRs", peer.OutAcRs}
    peer.EntityData.Leafs["in-ac-as"] = types.YLeaf{"InAcAs", peer.InAcAs}
    peer.EntityData.Leafs["out-ac-as"] = types.YLeaf{"OutAcAs", peer.OutAcAs}
    peer.EntityData.Leafs["in-ce-rs"] = types.YLeaf{"InCeRs", peer.InCeRs}
    peer.EntityData.Leafs["out-ce-rs"] = types.YLeaf{"OutCeRs", peer.OutCeRs}
    peer.EntityData.Leafs["in-ce-as"] = types.YLeaf{"InCeAs", peer.InCeAs}
    peer.EntityData.Leafs["out-ce-as"] = types.YLeaf{"OutCeAs", peer.OutCeAs}
    peer.EntityData.Leafs["in-dw-rs"] = types.YLeaf{"InDwRs", peer.InDwRs}
    peer.EntityData.Leafs["out-dw-rs"] = types.YLeaf{"OutDwRs", peer.OutDwRs}
    peer.EntityData.Leafs["in-dw-as"] = types.YLeaf{"InDwAs", peer.InDwAs}
    peer.EntityData.Leafs["out-dw-as"] = types.YLeaf{"OutDwAs", peer.OutDwAs}
    peer.EntityData.Leafs["in-dp-rs"] = types.YLeaf{"InDpRs", peer.InDpRs}
    peer.EntityData.Leafs["out-dp-rs"] = types.YLeaf{"OutDpRs", peer.OutDpRs}
    peer.EntityData.Leafs["in-dp-as"] = types.YLeaf{"InDpAs", peer.InDpAs}
    peer.EntityData.Leafs["out-dp-as"] = types.YLeaf{"OutDpAs", peer.OutDpAs}
    peer.EntityData.Leafs["in-ra-rs"] = types.YLeaf{"InRaRs", peer.InRaRs}
    peer.EntityData.Leafs["out-ra-rs"] = types.YLeaf{"OutRaRs", peer.OutRaRs}
    peer.EntityData.Leafs["in-ra-as"] = types.YLeaf{"InRaAs", peer.InRaAs}
    peer.EntityData.Leafs["out-ra-as"] = types.YLeaf{"OutRaAs", peer.OutRaAs}
    peer.EntityData.Leafs["in-st-rs"] = types.YLeaf{"InStRs", peer.InStRs}
    peer.EntityData.Leafs["out-st-rs"] = types.YLeaf{"OutStRs", peer.OutStRs}
    peer.EntityData.Leafs["in-st-as"] = types.YLeaf{"InStAs", peer.InStAs}
    peer.EntityData.Leafs["out-st-as"] = types.YLeaf{"OutStAs", peer.OutStAs}
    peer.EntityData.Leafs["in-cc-rs"] = types.YLeaf{"InCcRs", peer.InCcRs}
    peer.EntityData.Leafs["out-cc-rs"] = types.YLeaf{"OutCcRs", peer.OutCcRs}
    peer.EntityData.Leafs["in-cc-as"] = types.YLeaf{"InCcAs", peer.InCcAs}
    peer.EntityData.Leafs["out-cc-as"] = types.YLeaf{"OutCcAs", peer.OutCcAs}
    peer.EntityData.Leafs["out-aa-rs"] = types.YLeaf{"OutAaRs", peer.OutAaRs}
    peer.EntityData.Leafs["in-aa-as"] = types.YLeaf{"InAaAs", peer.InAaAs}
    peer.EntityData.Leafs["malformed-requests"] = types.YLeaf{"MalformedRequests", peer.MalformedRequests}
    peer.EntityData.Leafs["received-proto-errors"] = types.YLeaf{"ReceivedProtoErrors", peer.ReceivedProtoErrors}
    peer.EntityData.Leafs["sent-proto-errors"] = types.YLeaf{"SentProtoErrors", peer.SentProtoErrors}
    peer.EntityData.Leafs["received-transient-fails"] = types.YLeaf{"ReceivedTransientFails", peer.ReceivedTransientFails}
    peer.EntityData.Leafs["sent-transient-fails"] = types.YLeaf{"SentTransientFails", peer.SentTransientFails}
    peer.EntityData.Leafs["received-permanent-fails"] = types.YLeaf{"ReceivedPermanentFails", peer.ReceivedPermanentFails}
    peer.EntityData.Leafs["sent-permanent-fails"] = types.YLeaf{"SentPermanentFails", peer.SentPermanentFails}
    peer.EntityData.Leafs["transport-down"] = types.YLeaf{"TransportDown", peer.TransportDown}
    peer.EntityData.Leafs["state"] = types.YLeaf{"State", peer.State}
    return &(peer.EntityData)
}

// Aaa_Diameter_Nas
// Diameter NAS data
type Aaa_Diameter_Nas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NAS id. The type is string.
    AaanasId interface{}

    // List of NAS Entries. The type is slice of Aaa_Diameter_Nas_ListOfNas.
    ListOfNas []Aaa_Diameter_Nas_ListOfNas
}

func (nas *Aaa_Diameter_Nas) GetEntityData() *types.CommonEntityData {
    nas.EntityData.YFilter = nas.YFilter
    nas.EntityData.YangName = "nas"
    nas.EntityData.BundleName = "cisco_ios_xr"
    nas.EntityData.ParentYangName = "diameter"
    nas.EntityData.SegmentPath = "nas"
    nas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nas.EntityData.Children = make(map[string]types.YChild)
    nas.EntityData.Children["list-of-nas"] = types.YChild{"ListOfNas", nil}
    for i := range nas.ListOfNas {
        nas.EntityData.Children[types.GetSegmentPath(&nas.ListOfNas[i])] = types.YChild{"ListOfNas", &nas.ListOfNas[i]}
    }
    nas.EntityData.Leafs = make(map[string]types.YLeaf)
    nas.EntityData.Leafs["aaanas-id"] = types.YLeaf{"AaanasId", nas.AaanasId}
    return &(nas.EntityData)
}

// Aaa_Diameter_Nas_ListOfNas
// List of NAS Entries
type Aaa_Diameter_Nas_ListOfNas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA session id. The type is string.
    AaaSessionId interface{}

    // Diameter session id. The type is string.
    DiameterSessionId interface{}

    // Diameter AAR status. The type is interface{} with range: 0..4294967295.
    AuthenticationStatus interface{}

    // Diameter AAR status. The type is interface{} with range: 0..4294967295.
    AuthorizationStatus interface{}

    // Diameter ACR status start. The type is interface{} with range:
    // 0..4294967295.
    AccountingStatus interface{}

    // Diameter ACR status stop. The type is interface{} with range:
    // 0..4294967295.
    AccountingStatusStop interface{}

    // Diameter STR status. The type is interface{} with range: 0..4294967295.
    DisconnectStatus interface{}

    // Accounting intrim packet response in. The type is interface{} with range:
    // 0..4294967295.
    AccountingIntrimInPackets interface{}

    // Accounting intrim requests packets out. The type is interface{} with range:
    // 0..4294967295.
    AccountingIntrimOutPackets interface{}

    // Method list used for authentication. The type is string.
    MethodList interface{}

    // Server used for authentication. The type is string.
    ServerUsedList interface{}
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetEntityData() *types.CommonEntityData {
    listOfNas.EntityData.YFilter = listOfNas.YFilter
    listOfNas.EntityData.YangName = "list-of-nas"
    listOfNas.EntityData.BundleName = "cisco_ios_xr"
    listOfNas.EntityData.ParentYangName = "nas"
    listOfNas.EntityData.SegmentPath = "list-of-nas"
    listOfNas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    listOfNas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    listOfNas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    listOfNas.EntityData.Children = make(map[string]types.YChild)
    listOfNas.EntityData.Leafs = make(map[string]types.YLeaf)
    listOfNas.EntityData.Leafs["aaa-session-id"] = types.YLeaf{"AaaSessionId", listOfNas.AaaSessionId}
    listOfNas.EntityData.Leafs["diameter-session-id"] = types.YLeaf{"DiameterSessionId", listOfNas.DiameterSessionId}
    listOfNas.EntityData.Leafs["authentication-status"] = types.YLeaf{"AuthenticationStatus", listOfNas.AuthenticationStatus}
    listOfNas.EntityData.Leafs["authorization-status"] = types.YLeaf{"AuthorizationStatus", listOfNas.AuthorizationStatus}
    listOfNas.EntityData.Leafs["accounting-status"] = types.YLeaf{"AccountingStatus", listOfNas.AccountingStatus}
    listOfNas.EntityData.Leafs["accounting-status-stop"] = types.YLeaf{"AccountingStatusStop", listOfNas.AccountingStatusStop}
    listOfNas.EntityData.Leafs["disconnect-status"] = types.YLeaf{"DisconnectStatus", listOfNas.DisconnectStatus}
    listOfNas.EntityData.Leafs["accounting-intrim-in-packets"] = types.YLeaf{"AccountingIntrimInPackets", listOfNas.AccountingIntrimInPackets}
    listOfNas.EntityData.Leafs["accounting-intrim-out-packets"] = types.YLeaf{"AccountingIntrimOutPackets", listOfNas.AccountingIntrimOutPackets}
    listOfNas.EntityData.Leafs["method-list"] = types.YLeaf{"MethodList", listOfNas.MethodList}
    listOfNas.EntityData.Leafs["server-used-list"] = types.YLeaf{"ServerUsedList", listOfNas.ServerUsedList}
    return &(listOfNas.EntityData)
}

// Aaa_Diameter_NasSummary
// Diameter NAS summary
type Aaa_Diameter_NasSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authentication response pkt in. The type is interface{} with range:
    // 0..4294967295.
    AuthenResponseInPackets interface{}

    // Authentication request pkt out. The type is interface{} with range:
    // 0..4294967295.
    AuthenRequestOutPackets interface{}

    // Authentication request from client. The type is interface{} with range:
    // 0..4294967295.
    AuthenRequestInPackets interface{}

    // Authentication response frwd to client. The type is interface{} with range:
    // 0..4294967295.
    AuthenResponseOutPackets interface{}

    // Authentication response with success. The type is interface{} with range:
    // 0..4294967295.
    AuthenSuccessPackets interface{}

    // Authentication response with failure. The type is interface{} with range:
    // 0..4294967295.
    AuthenResponseFailPackets interface{}

    // Authorization response packet in. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationInPackets interface{}

    // Authorization request packet out. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationOutPackets interface{}

    // Authourization request from cleint. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationRequestInPackets interface{}

    // Authourization response frwd to client. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationResponseOutPackets interface{}

    // Authentication response with success. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationResponseSuccessPackets interface{}

    // Authentication response with failure. The type is interface{} with range:
    // 0..4294967295.
    AuthorizationResponseFailPackets interface{}

    // Accounting packet response in. The type is interface{} with range:
    // 0..4294967295.
    AccountingResponseInPackets interface{}

    // Accounting requests packets out. The type is interface{} with range:
    // 0..4294967295.
    AccountingRequestOutPackets interface{}

    // Accounting start request from cleint. The type is interface{} with range:
    // 0..4294967295.
    AccountingStartRequestPackets interface{}

    // Accounting start response forward to client. The type is interface{} with
    // range: 0..4294967295.
    AccountingStartResponsePackets interface{}

    // Accounting start response with success. The type is interface{} with range:
    // 0..4294967295.
    AccountingStartSuccessPackets interface{}

    // Accounting start response with failure. The type is interface{} with range:
    // 0..4294967295.
    AccountingStartFailedPackets interface{}

    // Accounting stop packet response in. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopResponseInPackets interface{}

    // Accounting stop requests packets out. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopRequestOutPackets interface{}

    // Acct stop request from cleint. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopRequestInPackets interface{}

    // Acct stop response forward to client. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopResponseOutPackets interface{}

    // Accounting stop response with success. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopSuccessResponsePackets interface{}

    // Accounting stop response with failure. The type is interface{} with range:
    // 0..4294967295.
    AccountingStopFailedPackets interface{}

    // Accounting interim packet response in. The type is interface{} with range:
    // 0..4294967295.
    AccountingIntrimResponseInPackets interface{}

    // Accounting interim requests packets out. The type is interface{} with
    // range: 0..4294967295.
    AccountingInterimRequestOutPackets interface{}

    // Accounting Interim request from cleint. The type is interface{} with range:
    // 0..4294967295.
    AccountingInterimRequestInPackets interface{}

    // Accounting interim response frwd to client. The type is interface{} with
    // range: 0..4294967295.
    AccountingInterimResponseOutPackets interface{}

    // Accounting interim response with success. The type is interface{} with
    // range: 0..4294967295.
    AccountingInterimSuccessPackets interface{}

    // Accounting interim response with failure. The type is interface{} with
    // range: 0..4294967295.
    AccountingInterimFailedPackets interface{}

    // Disconnect response packets in. The type is interface{} with range:
    // 0..4294967295.
    DisconnectResponseInPackets interface{}

    // Disconnect request pkt out. The type is interface{} with range:
    // 0..4294967295.
    DisconnectRequestOutPackets interface{}

    // Disconnect request from cleint. The type is interface{} with range:
    // 0..4294967295.
    DisconnectRequestInPackets interface{}

    // Disconnect response forward to client. The type is interface{} with range:
    // 0..4294967295.
    DisconnectResponseOutPackets interface{}

    // Disconnect response with success. The type is interface{} with range:
    // 0..4294967295.
    DisconnectSuccessResponsePackets interface{}

    // Disconnect response with failure. The type is interface{} with range:
    // 0..4294967295.
    DisconnectFailedResponsePackets interface{}

    // COA/RAR Request packet in. The type is interface{} with range:
    // 0..4294967295.
    CoaRequestInPackets interface{}

    // COA/RAR Response packet out. The type is interface{} with range:
    // 0..4294967295.
    CoaResponseOutPackets interface{}

    // COA request from client. The type is interface{} with range: 0..4294967295.
    CoaRequestPackets interface{}

    // COA response forward to client. The type is interface{} with range:
    // 0..4294967295.
    CoaResponsePackets interface{}

    // COA response with success. The type is interface{} with range:
    // 0..4294967295.
    CoaSuccessPackets interface{}

    // COA response with failure. The type is interface{} with range:
    // 0..4294967295.
    CoaFailedPackets interface{}

    // POD/RAR Request packets in. The type is interface{} with range:
    // 0..4294967295.
    PodInPackets interface{}

    // PAD/RAR Response packets out. The type is interface{} with range:
    // 0..4294967295.
    PodOutPackets interface{}

    // POD request from cleint. The type is interface{} with range: 0..4294967295.
    PodRequestInPackets interface{}

    // POD response forward to client. The type is interface{} with range:
    // 0..4294967295.
    PodResponseOutPackets interface{}

    // POD response with success. The type is interface{} with range:
    // 0..4294967295.
    PodSuccessPackets interface{}

    // POD response with failure. The type is interface{} with range:
    // 0..4294967295.
    PodFailedPackets interface{}
}

func (nasSummary *Aaa_Diameter_NasSummary) GetEntityData() *types.CommonEntityData {
    nasSummary.EntityData.YFilter = nasSummary.YFilter
    nasSummary.EntityData.YangName = "nas-summary"
    nasSummary.EntityData.BundleName = "cisco_ios_xr"
    nasSummary.EntityData.ParentYangName = "diameter"
    nasSummary.EntityData.SegmentPath = "nas-summary"
    nasSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nasSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nasSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nasSummary.EntityData.Children = make(map[string]types.YChild)
    nasSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    nasSummary.EntityData.Leafs["authen-response-in-packets"] = types.YLeaf{"AuthenResponseInPackets", nasSummary.AuthenResponseInPackets}
    nasSummary.EntityData.Leafs["authen-request-out-packets"] = types.YLeaf{"AuthenRequestOutPackets", nasSummary.AuthenRequestOutPackets}
    nasSummary.EntityData.Leafs["authen-request-in-packets"] = types.YLeaf{"AuthenRequestInPackets", nasSummary.AuthenRequestInPackets}
    nasSummary.EntityData.Leafs["authen-response-out-packets"] = types.YLeaf{"AuthenResponseOutPackets", nasSummary.AuthenResponseOutPackets}
    nasSummary.EntityData.Leafs["authen-success-packets"] = types.YLeaf{"AuthenSuccessPackets", nasSummary.AuthenSuccessPackets}
    nasSummary.EntityData.Leafs["authen-response-fail-packets"] = types.YLeaf{"AuthenResponseFailPackets", nasSummary.AuthenResponseFailPackets}
    nasSummary.EntityData.Leafs["authorization-in-packets"] = types.YLeaf{"AuthorizationInPackets", nasSummary.AuthorizationInPackets}
    nasSummary.EntityData.Leafs["authorization-out-packets"] = types.YLeaf{"AuthorizationOutPackets", nasSummary.AuthorizationOutPackets}
    nasSummary.EntityData.Leafs["authorization-request-in-packets"] = types.YLeaf{"AuthorizationRequestInPackets", nasSummary.AuthorizationRequestInPackets}
    nasSummary.EntityData.Leafs["authorization-response-out-packets"] = types.YLeaf{"AuthorizationResponseOutPackets", nasSummary.AuthorizationResponseOutPackets}
    nasSummary.EntityData.Leafs["authorization-response-success-packets"] = types.YLeaf{"AuthorizationResponseSuccessPackets", nasSummary.AuthorizationResponseSuccessPackets}
    nasSummary.EntityData.Leafs["authorization-response-fail-packets"] = types.YLeaf{"AuthorizationResponseFailPackets", nasSummary.AuthorizationResponseFailPackets}
    nasSummary.EntityData.Leafs["accounting-response-in-packets"] = types.YLeaf{"AccountingResponseInPackets", nasSummary.AccountingResponseInPackets}
    nasSummary.EntityData.Leafs["accounting-request-out-packets"] = types.YLeaf{"AccountingRequestOutPackets", nasSummary.AccountingRequestOutPackets}
    nasSummary.EntityData.Leafs["accounting-start-request-packets"] = types.YLeaf{"AccountingStartRequestPackets", nasSummary.AccountingStartRequestPackets}
    nasSummary.EntityData.Leafs["accounting-start-response-packets"] = types.YLeaf{"AccountingStartResponsePackets", nasSummary.AccountingStartResponsePackets}
    nasSummary.EntityData.Leafs["accounting-start-success-packets"] = types.YLeaf{"AccountingStartSuccessPackets", nasSummary.AccountingStartSuccessPackets}
    nasSummary.EntityData.Leafs["accounting-start-failed-packets"] = types.YLeaf{"AccountingStartFailedPackets", nasSummary.AccountingStartFailedPackets}
    nasSummary.EntityData.Leafs["accounting-stop-response-in-packets"] = types.YLeaf{"AccountingStopResponseInPackets", nasSummary.AccountingStopResponseInPackets}
    nasSummary.EntityData.Leafs["accounting-stop-request-out-packets"] = types.YLeaf{"AccountingStopRequestOutPackets", nasSummary.AccountingStopRequestOutPackets}
    nasSummary.EntityData.Leafs["accounting-stop-request-in-packets"] = types.YLeaf{"AccountingStopRequestInPackets", nasSummary.AccountingStopRequestInPackets}
    nasSummary.EntityData.Leafs["accounting-stop-response-out-packets"] = types.YLeaf{"AccountingStopResponseOutPackets", nasSummary.AccountingStopResponseOutPackets}
    nasSummary.EntityData.Leafs["accounting-stop-success-response-packets"] = types.YLeaf{"AccountingStopSuccessResponsePackets", nasSummary.AccountingStopSuccessResponsePackets}
    nasSummary.EntityData.Leafs["accounting-stop-failed-packets"] = types.YLeaf{"AccountingStopFailedPackets", nasSummary.AccountingStopFailedPackets}
    nasSummary.EntityData.Leafs["accounting-intrim-response-in-packets"] = types.YLeaf{"AccountingIntrimResponseInPackets", nasSummary.AccountingIntrimResponseInPackets}
    nasSummary.EntityData.Leafs["accounting-interim-request-out-packets"] = types.YLeaf{"AccountingInterimRequestOutPackets", nasSummary.AccountingInterimRequestOutPackets}
    nasSummary.EntityData.Leafs["accounting-interim-request-in-packets"] = types.YLeaf{"AccountingInterimRequestInPackets", nasSummary.AccountingInterimRequestInPackets}
    nasSummary.EntityData.Leafs["accounting-interim-response-out-packets"] = types.YLeaf{"AccountingInterimResponseOutPackets", nasSummary.AccountingInterimResponseOutPackets}
    nasSummary.EntityData.Leafs["accounting-interim-success-packets"] = types.YLeaf{"AccountingInterimSuccessPackets", nasSummary.AccountingInterimSuccessPackets}
    nasSummary.EntityData.Leafs["accounting-interim-failed-packets"] = types.YLeaf{"AccountingInterimFailedPackets", nasSummary.AccountingInterimFailedPackets}
    nasSummary.EntityData.Leafs["disconnect-response-in-packets"] = types.YLeaf{"DisconnectResponseInPackets", nasSummary.DisconnectResponseInPackets}
    nasSummary.EntityData.Leafs["disconnect-request-out-packets"] = types.YLeaf{"DisconnectRequestOutPackets", nasSummary.DisconnectRequestOutPackets}
    nasSummary.EntityData.Leafs["disconnect-request-in-packets"] = types.YLeaf{"DisconnectRequestInPackets", nasSummary.DisconnectRequestInPackets}
    nasSummary.EntityData.Leafs["disconnect-response-out-packets"] = types.YLeaf{"DisconnectResponseOutPackets", nasSummary.DisconnectResponseOutPackets}
    nasSummary.EntityData.Leafs["disconnect-success-response-packets"] = types.YLeaf{"DisconnectSuccessResponsePackets", nasSummary.DisconnectSuccessResponsePackets}
    nasSummary.EntityData.Leafs["disconnect-failed-response-packets"] = types.YLeaf{"DisconnectFailedResponsePackets", nasSummary.DisconnectFailedResponsePackets}
    nasSummary.EntityData.Leafs["coa-request-in-packets"] = types.YLeaf{"CoaRequestInPackets", nasSummary.CoaRequestInPackets}
    nasSummary.EntityData.Leafs["coa-response-out-packets"] = types.YLeaf{"CoaResponseOutPackets", nasSummary.CoaResponseOutPackets}
    nasSummary.EntityData.Leafs["coa-request-packets"] = types.YLeaf{"CoaRequestPackets", nasSummary.CoaRequestPackets}
    nasSummary.EntityData.Leafs["coa-response-packets"] = types.YLeaf{"CoaResponsePackets", nasSummary.CoaResponsePackets}
    nasSummary.EntityData.Leafs["coa-success-packets"] = types.YLeaf{"CoaSuccessPackets", nasSummary.CoaSuccessPackets}
    nasSummary.EntityData.Leafs["coa-failed-packets"] = types.YLeaf{"CoaFailedPackets", nasSummary.CoaFailedPackets}
    nasSummary.EntityData.Leafs["pod-in-packets"] = types.YLeaf{"PodInPackets", nasSummary.PodInPackets}
    nasSummary.EntityData.Leafs["pod-out-packets"] = types.YLeaf{"PodOutPackets", nasSummary.PodOutPackets}
    nasSummary.EntityData.Leafs["pod-request-in-packets"] = types.YLeaf{"PodRequestInPackets", nasSummary.PodRequestInPackets}
    nasSummary.EntityData.Leafs["pod-response-out-packets"] = types.YLeaf{"PodResponseOutPackets", nasSummary.PodResponseOutPackets}
    nasSummary.EntityData.Leafs["pod-success-packets"] = types.YLeaf{"PodSuccessPackets", nasSummary.PodSuccessPackets}
    nasSummary.EntityData.Leafs["pod-failed-packets"] = types.YLeaf{"PodFailedPackets", nasSummary.PodFailedPackets}
    return &(nasSummary.EntityData)
}

// Aaa_Diameter_GySessionIds
// Diameter Gy Session data list
type Aaa_Diameter_GySessionIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diameter Gy Session data. The type is slice of
    // Aaa_Diameter_GySessionIds_GySessionId.
    GySessionId []Aaa_Diameter_GySessionIds_GySessionId
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetEntityData() *types.CommonEntityData {
    gySessionIds.EntityData.YFilter = gySessionIds.YFilter
    gySessionIds.EntityData.YangName = "gy-session-ids"
    gySessionIds.EntityData.BundleName = "cisco_ios_xr"
    gySessionIds.EntityData.ParentYangName = "diameter"
    gySessionIds.EntityData.SegmentPath = "gy-session-ids"
    gySessionIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gySessionIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gySessionIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gySessionIds.EntityData.Children = make(map[string]types.YChild)
    gySessionIds.EntityData.Children["gy-session-id"] = types.YChild{"GySessionId", nil}
    for i := range gySessionIds.GySessionId {
        gySessionIds.EntityData.Children[types.GetSegmentPath(&gySessionIds.GySessionId[i])] = types.YChild{"GySessionId", &gySessionIds.GySessionId[i]}
    }
    gySessionIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(gySessionIds.EntityData)
}

// Aaa_Diameter_GySessionIds_GySessionId
// Diameter Gy Session data
type Aaa_Diameter_GySessionIds_GySessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // AAA session id. The type is interface{} with range: 0..4294967295.
    AaaSessionId interface{}

    // AAA Parent session id. The type is interface{} with range: 0..4294967295.
    ParentAaaSessionId interface{}

    // Diameter session id. The type is string.
    DiameterSessionId interface{}

    // Request Number. The type is interface{} with range: 0..4294967295.
    RequestNumber interface{}

    // Session State. The type is string.
    SessionState interface{}

    // Request Type. The type is string.
    RequestType interface{}

    // Gy Retry count. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetEntityData() *types.CommonEntityData {
    gySessionId.EntityData.YFilter = gySessionId.YFilter
    gySessionId.EntityData.YangName = "gy-session-id"
    gySessionId.EntityData.BundleName = "cisco_ios_xr"
    gySessionId.EntityData.ParentYangName = "gy-session-ids"
    gySessionId.EntityData.SegmentPath = "gy-session-id" + "[session-id='" + fmt.Sprintf("%v", gySessionId.SessionId) + "']"
    gySessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gySessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gySessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gySessionId.EntityData.Children = make(map[string]types.YChild)
    gySessionId.EntityData.Leafs = make(map[string]types.YLeaf)
    gySessionId.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", gySessionId.SessionId}
    gySessionId.EntityData.Leafs["aaa-session-id"] = types.YLeaf{"AaaSessionId", gySessionId.AaaSessionId}
    gySessionId.EntityData.Leafs["parent-aaa-session-id"] = types.YLeaf{"ParentAaaSessionId", gySessionId.ParentAaaSessionId}
    gySessionId.EntityData.Leafs["diameter-session-id"] = types.YLeaf{"DiameterSessionId", gySessionId.DiameterSessionId}
    gySessionId.EntityData.Leafs["request-number"] = types.YLeaf{"RequestNumber", gySessionId.RequestNumber}
    gySessionId.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", gySessionId.SessionState}
    gySessionId.EntityData.Leafs["request-type"] = types.YLeaf{"RequestType", gySessionId.RequestType}
    gySessionId.EntityData.Leafs["retry-count"] = types.YLeaf{"RetryCount", gySessionId.RetryCount}
    return &(gySessionId.EntityData)
}

// Aaa_Diameter_GyStatistics
// Diameter Gy Statistics data
type Aaa_Diameter_GyStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CCR Initial Messages. The type is interface{} with range: 0..4294967295.
    CcrInitMessages interface{}

    // CCR Initial Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrInitFailedMessages interface{}

    // CCR Initial Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrInitTimedOutMessages interface{}

    // CCR Initial Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrInitRetryMessages interface{}

    // CCR Update Messages. The type is interface{} with range: 0..4294967295.
    CcrUpdateMessages interface{}

    // CCR Update Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateFailedMessages interface{}

    // CCR Update Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateTimedOutMessages interface{}

    // CCR Update Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrUpdateRetryMessages interface{}

    // CCR Final Messages. The type is interface{} with range: 0..4294967295.
    CcrFinalMessages interface{}

    // CCR Final Messages Failed. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalFailedMessages interface{}

    // CCR Final Messages Timed Out. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalTimedOutMessages interface{}

    // CCR Final Messages retry. The type is interface{} with range:
    // 0..4294967295.
    CcrFinalRetryMessages interface{}

    // CCA Initial Messages. The type is interface{} with range: 0..4294967295.
    CcaInitMessages interface{}

    // CCA Initial Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaInitErrorMessages interface{}

    // CCA Update Messages. The type is interface{} with range: 0..4294967295.
    CcaUpdateMessages interface{}

    // CCA Update Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaUpdateErrorMessages interface{}

    // CCA Final Messages. The type is interface{} with range: 0..4294967295.
    CcaFinalMessages interface{}

    // CCA Final Messages Error. The type is interface{} with range:
    // 0..4294967295.
    CcaFinalErrorMessages interface{}

    // RAR Received Messages. The type is interface{} with range: 0..4294967295.
    RarReceivedMessages interface{}

    // RAR Received Messages Error. The type is interface{} with range:
    // 0..4294967295.
    RarReceivedErrorMessages interface{}

    // RAA Sent Messages. The type is interface{} with range: 0..4294967295.
    RaaSentMessages interface{}

    // RAA Sent Messages Error. The type is interface{} with range: 0..4294967295.
    RaaSentErrorMessages interface{}

    // ASR Received Messages. The type is interface{} with range: 0..4294967295.
    AsrReceivedMessages interface{}

    // ASR Received Messages Error. The type is interface{} with range:
    // 0..4294967295.
    AsrReceivedErrorMessages interface{}

    // ASA Sent Messages. The type is interface{} with range: 0..4294967295.
    AsaSentMessages interface{}

    // ASA Sent Messages Error. The type is interface{} with range: 0..4294967295.
    AsaSentErrorMessages interface{}

    // Unknown Request Messages. The type is interface{} with range:
    // 0..4294967295.
    UnknownRequestMessages interface{}

    // Restore Sessions. The type is interface{} with range: 0..4294967295.
    RestoreSessions interface{}

    // Total Opened Sessions. The type is interface{} with range: 0..4294967295.
    OpenSessions interface{}

    // Total Closed Sessions. The type is interface{} with range: 0..4294967295.
    CloseSessions interface{}

    // Total Active Sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetEntityData() *types.CommonEntityData {
    gyStatistics.EntityData.YFilter = gyStatistics.YFilter
    gyStatistics.EntityData.YangName = "gy-statistics"
    gyStatistics.EntityData.BundleName = "cisco_ios_xr"
    gyStatistics.EntityData.ParentYangName = "diameter"
    gyStatistics.EntityData.SegmentPath = "gy-statistics"
    gyStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gyStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gyStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gyStatistics.EntityData.Children = make(map[string]types.YChild)
    gyStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    gyStatistics.EntityData.Leafs["ccr-init-messages"] = types.YLeaf{"CcrInitMessages", gyStatistics.CcrInitMessages}
    gyStatistics.EntityData.Leafs["ccr-init-failed-messages"] = types.YLeaf{"CcrInitFailedMessages", gyStatistics.CcrInitFailedMessages}
    gyStatistics.EntityData.Leafs["ccr-init-timed-out-messages"] = types.YLeaf{"CcrInitTimedOutMessages", gyStatistics.CcrInitTimedOutMessages}
    gyStatistics.EntityData.Leafs["ccr-init-retry-messages"] = types.YLeaf{"CcrInitRetryMessages", gyStatistics.CcrInitRetryMessages}
    gyStatistics.EntityData.Leafs["ccr-update-messages"] = types.YLeaf{"CcrUpdateMessages", gyStatistics.CcrUpdateMessages}
    gyStatistics.EntityData.Leafs["ccr-update-failed-messages"] = types.YLeaf{"CcrUpdateFailedMessages", gyStatistics.CcrUpdateFailedMessages}
    gyStatistics.EntityData.Leafs["ccr-update-timed-out-messages"] = types.YLeaf{"CcrUpdateTimedOutMessages", gyStatistics.CcrUpdateTimedOutMessages}
    gyStatistics.EntityData.Leafs["ccr-update-retry-messages"] = types.YLeaf{"CcrUpdateRetryMessages", gyStatistics.CcrUpdateRetryMessages}
    gyStatistics.EntityData.Leafs["ccr-final-messages"] = types.YLeaf{"CcrFinalMessages", gyStatistics.CcrFinalMessages}
    gyStatistics.EntityData.Leafs["ccr-final-failed-messages"] = types.YLeaf{"CcrFinalFailedMessages", gyStatistics.CcrFinalFailedMessages}
    gyStatistics.EntityData.Leafs["ccr-final-timed-out-messages"] = types.YLeaf{"CcrFinalTimedOutMessages", gyStatistics.CcrFinalTimedOutMessages}
    gyStatistics.EntityData.Leafs["ccr-final-retry-messages"] = types.YLeaf{"CcrFinalRetryMessages", gyStatistics.CcrFinalRetryMessages}
    gyStatistics.EntityData.Leafs["cca-init-messages"] = types.YLeaf{"CcaInitMessages", gyStatistics.CcaInitMessages}
    gyStatistics.EntityData.Leafs["cca-init-error-messages"] = types.YLeaf{"CcaInitErrorMessages", gyStatistics.CcaInitErrorMessages}
    gyStatistics.EntityData.Leafs["cca-update-messages"] = types.YLeaf{"CcaUpdateMessages", gyStatistics.CcaUpdateMessages}
    gyStatistics.EntityData.Leafs["cca-update-error-messages"] = types.YLeaf{"CcaUpdateErrorMessages", gyStatistics.CcaUpdateErrorMessages}
    gyStatistics.EntityData.Leafs["cca-final-messages"] = types.YLeaf{"CcaFinalMessages", gyStatistics.CcaFinalMessages}
    gyStatistics.EntityData.Leafs["cca-final-error-messages"] = types.YLeaf{"CcaFinalErrorMessages", gyStatistics.CcaFinalErrorMessages}
    gyStatistics.EntityData.Leafs["rar-received-messages"] = types.YLeaf{"RarReceivedMessages", gyStatistics.RarReceivedMessages}
    gyStatistics.EntityData.Leafs["rar-received-error-messages"] = types.YLeaf{"RarReceivedErrorMessages", gyStatistics.RarReceivedErrorMessages}
    gyStatistics.EntityData.Leafs["raa-sent-messages"] = types.YLeaf{"RaaSentMessages", gyStatistics.RaaSentMessages}
    gyStatistics.EntityData.Leafs["raa-sent-error-messages"] = types.YLeaf{"RaaSentErrorMessages", gyStatistics.RaaSentErrorMessages}
    gyStatistics.EntityData.Leafs["asr-received-messages"] = types.YLeaf{"AsrReceivedMessages", gyStatistics.AsrReceivedMessages}
    gyStatistics.EntityData.Leafs["asr-received-error-messages"] = types.YLeaf{"AsrReceivedErrorMessages", gyStatistics.AsrReceivedErrorMessages}
    gyStatistics.EntityData.Leafs["asa-sent-messages"] = types.YLeaf{"AsaSentMessages", gyStatistics.AsaSentMessages}
    gyStatistics.EntityData.Leafs["asa-sent-error-messages"] = types.YLeaf{"AsaSentErrorMessages", gyStatistics.AsaSentErrorMessages}
    gyStatistics.EntityData.Leafs["unknown-request-messages"] = types.YLeaf{"UnknownRequestMessages", gyStatistics.UnknownRequestMessages}
    gyStatistics.EntityData.Leafs["restore-sessions"] = types.YLeaf{"RestoreSessions", gyStatistics.RestoreSessions}
    gyStatistics.EntityData.Leafs["open-sessions"] = types.YLeaf{"OpenSessions", gyStatistics.OpenSessions}
    gyStatistics.EntityData.Leafs["close-sessions"] = types.YLeaf{"CloseSessions", gyStatistics.CloseSessions}
    gyStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", gyStatistics.ActiveSessions}
    return &(gyStatistics.EntityData)
}

// Aaa_Diameter_GxSessionIds
// Diameter Gx Session data list
type Aaa_Diameter_GxSessionIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diameter Gx Session data. The type is slice of
    // Aaa_Diameter_GxSessionIds_GxSessionId.
    GxSessionId []Aaa_Diameter_GxSessionIds_GxSessionId
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetEntityData() *types.CommonEntityData {
    gxSessionIds.EntityData.YFilter = gxSessionIds.YFilter
    gxSessionIds.EntityData.YangName = "gx-session-ids"
    gxSessionIds.EntityData.BundleName = "cisco_ios_xr"
    gxSessionIds.EntityData.ParentYangName = "diameter"
    gxSessionIds.EntityData.SegmentPath = "gx-session-ids"
    gxSessionIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gxSessionIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gxSessionIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gxSessionIds.EntityData.Children = make(map[string]types.YChild)
    gxSessionIds.EntityData.Children["gx-session-id"] = types.YChild{"GxSessionId", nil}
    for i := range gxSessionIds.GxSessionId {
        gxSessionIds.EntityData.Children[types.GetSegmentPath(&gxSessionIds.GxSessionId[i])] = types.YChild{"GxSessionId", &gxSessionIds.GxSessionId[i]}
    }
    gxSessionIds.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(gxSessionIds.EntityData)
}

// Aaa_Diameter_GxSessionIds_GxSessionId
// Diameter Gx Session data
type Aaa_Diameter_GxSessionIds_GxSessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Id. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // AAA session id. The type is interface{} with range: 0..4294967295.
    AaaSessionId interface{}

    // Diameter session id. The type is string.
    DiameterSessionId interface{}

    // Request Number. The type is interface{} with range: 0..4294967295.
    RequestNumber interface{}

    // Session State. The type is string.
    SessionState interface{}

    // Request Type. The type is string.
    RequestType interface{}

    // Gx Retry count. The type is interface{} with range: 0..4294967295.
    RetryCount interface{}
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetEntityData() *types.CommonEntityData {
    gxSessionId.EntityData.YFilter = gxSessionId.YFilter
    gxSessionId.EntityData.YangName = "gx-session-id"
    gxSessionId.EntityData.BundleName = "cisco_ios_xr"
    gxSessionId.EntityData.ParentYangName = "gx-session-ids"
    gxSessionId.EntityData.SegmentPath = "gx-session-id" + "[session-id='" + fmt.Sprintf("%v", gxSessionId.SessionId) + "']"
    gxSessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gxSessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gxSessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gxSessionId.EntityData.Children = make(map[string]types.YChild)
    gxSessionId.EntityData.Leafs = make(map[string]types.YLeaf)
    gxSessionId.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", gxSessionId.SessionId}
    gxSessionId.EntityData.Leafs["aaa-session-id"] = types.YLeaf{"AaaSessionId", gxSessionId.AaaSessionId}
    gxSessionId.EntityData.Leafs["diameter-session-id"] = types.YLeaf{"DiameterSessionId", gxSessionId.DiameterSessionId}
    gxSessionId.EntityData.Leafs["request-number"] = types.YLeaf{"RequestNumber", gxSessionId.RequestNumber}
    gxSessionId.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", gxSessionId.SessionState}
    gxSessionId.EntityData.Leafs["request-type"] = types.YLeaf{"RequestType", gxSessionId.RequestType}
    gxSessionId.EntityData.Leafs["retry-count"] = types.YLeaf{"RetryCount", gxSessionId.RetryCount}
    return &(gxSessionId.EntityData)
}

// Aaa_Diameter_NasSession
// Diameter Nas Session data
type Aaa_Diameter_NasSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA NAS id. The type is string.
    AaanasId interface{}

    // List of NAS Entries. The type is slice of
    // Aaa_Diameter_NasSession_ListOfNas.
    ListOfNas []Aaa_Diameter_NasSession_ListOfNas
}

func (nasSession *Aaa_Diameter_NasSession) GetEntityData() *types.CommonEntityData {
    nasSession.EntityData.YFilter = nasSession.YFilter
    nasSession.EntityData.YangName = "nas-session"
    nasSession.EntityData.BundleName = "cisco_ios_xr"
    nasSession.EntityData.ParentYangName = "diameter"
    nasSession.EntityData.SegmentPath = "nas-session"
    nasSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nasSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nasSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nasSession.EntityData.Children = make(map[string]types.YChild)
    nasSession.EntityData.Children["list-of-nas"] = types.YChild{"ListOfNas", nil}
    for i := range nasSession.ListOfNas {
        nasSession.EntityData.Children[types.GetSegmentPath(&nasSession.ListOfNas[i])] = types.YChild{"ListOfNas", &nasSession.ListOfNas[i]}
    }
    nasSession.EntityData.Leafs = make(map[string]types.YLeaf)
    nasSession.EntityData.Leafs["aaanas-id"] = types.YLeaf{"AaanasId", nasSession.AaanasId}
    return &(nasSession.EntityData)
}

// Aaa_Diameter_NasSession_ListOfNas
// List of NAS Entries
type Aaa_Diameter_NasSession_ListOfNas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA session id. The type is string.
    AaaSessionId interface{}

    // Diameter session id. The type is string.
    DiameterSessionId interface{}

    // Diameter AAR status. The type is interface{} with range: 0..4294967295.
    AuthenticationStatus interface{}

    // Diameter AAR status. The type is interface{} with range: 0..4294967295.
    AuthorizationStatus interface{}

    // Diameter ACR status start. The type is interface{} with range:
    // 0..4294967295.
    AccountingStatus interface{}

    // Diameter ACR status stop. The type is interface{} with range:
    // 0..4294967295.
    AccountingStatusStop interface{}

    // Diameter STR status. The type is interface{} with range: 0..4294967295.
    DisconnectStatus interface{}

    // Accounting intrim packet response in. The type is interface{} with range:
    // 0..4294967295.
    AccountingIntrimInPackets interface{}

    // Accounting intrim requests packets out. The type is interface{} with range:
    // 0..4294967295.
    AccountingIntrimOutPackets interface{}

    // Method list used for authentication. The type is string.
    MethodList interface{}

    // Server used for authentication. The type is string.
    ServerUsedList interface{}
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetEntityData() *types.CommonEntityData {
    listOfNas.EntityData.YFilter = listOfNas.YFilter
    listOfNas.EntityData.YangName = "list-of-nas"
    listOfNas.EntityData.BundleName = "cisco_ios_xr"
    listOfNas.EntityData.ParentYangName = "nas-session"
    listOfNas.EntityData.SegmentPath = "list-of-nas"
    listOfNas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    listOfNas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    listOfNas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    listOfNas.EntityData.Children = make(map[string]types.YChild)
    listOfNas.EntityData.Leafs = make(map[string]types.YLeaf)
    listOfNas.EntityData.Leafs["aaa-session-id"] = types.YLeaf{"AaaSessionId", listOfNas.AaaSessionId}
    listOfNas.EntityData.Leafs["diameter-session-id"] = types.YLeaf{"DiameterSessionId", listOfNas.DiameterSessionId}
    listOfNas.EntityData.Leafs["authentication-status"] = types.YLeaf{"AuthenticationStatus", listOfNas.AuthenticationStatus}
    listOfNas.EntityData.Leafs["authorization-status"] = types.YLeaf{"AuthorizationStatus", listOfNas.AuthorizationStatus}
    listOfNas.EntityData.Leafs["accounting-status"] = types.YLeaf{"AccountingStatus", listOfNas.AccountingStatus}
    listOfNas.EntityData.Leafs["accounting-status-stop"] = types.YLeaf{"AccountingStatusStop", listOfNas.AccountingStatusStop}
    listOfNas.EntityData.Leafs["disconnect-status"] = types.YLeaf{"DisconnectStatus", listOfNas.DisconnectStatus}
    listOfNas.EntityData.Leafs["accounting-intrim-in-packets"] = types.YLeaf{"AccountingIntrimInPackets", listOfNas.AccountingIntrimInPackets}
    listOfNas.EntityData.Leafs["accounting-intrim-out-packets"] = types.YLeaf{"AccountingIntrimOutPackets", listOfNas.AccountingIntrimOutPackets}
    listOfNas.EntityData.Leafs["method-list"] = types.YLeaf{"MethodList", listOfNas.MethodList}
    listOfNas.EntityData.Leafs["server-used-list"] = types.YLeaf{"ServerUsedList", listOfNas.ServerUsedList}
    return &(listOfNas.EntityData)
}

// Aaa_Radius
// RADIUS operational data
type Aaa_Radius struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of RADIUS servers configured.
    Servers Aaa_Radius_Servers

    // RADIUS source interfaces.
    RadiusSourceInterface Aaa_Radius_RadiusSourceInterface

    // RADIUS Client Information.
    Global Aaa_Radius_Global
}

func (radius *Aaa_Radius) GetEntityData() *types.CommonEntityData {
    radius.EntityData.YFilter = radius.YFilter
    radius.EntityData.YangName = "radius"
    radius.EntityData.BundleName = "cisco_ios_xr"
    radius.EntityData.ParentYangName = "aaa"
    radius.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-protocol-radius-oper:radius"
    radius.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radius.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radius.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radius.EntityData.Children = make(map[string]types.YChild)
    radius.EntityData.Children["servers"] = types.YChild{"Servers", &radius.Servers}
    radius.EntityData.Children["radius-source-interface"] = types.YChild{"RadiusSourceInterface", &radius.RadiusSourceInterface}
    radius.EntityData.Children["global"] = types.YChild{"Global", &radius.Global}
    radius.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(radius.EntityData)
}

// Aaa_Radius_Servers
// List of RADIUS servers configured
type Aaa_Radius_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS Server. The type is slice of Aaa_Radius_Servers_Server.
    Server []Aaa_Radius_Servers_Server
}

func (servers *Aaa_Radius_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "radius"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = make(map[string]types.YChild)
    servers.EntityData.Children["server"] = types.YChild{"Server", nil}
    for i := range servers.Server {
        servers.EntityData.Children[types.GetSegmentPath(&servers.Server[i])] = types.YChild{"Server", &servers.Server[i]}
    }
    servers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(servers.EntityData)
}

// Aaa_Radius_Servers_Server
// RADIUS Server
type Aaa_Radius_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is one of the following types: string
    // with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}

    // Authentication Port number (standard port 1645). The type is interface{}
    // with range: 1..65535.
    AuthPortNumber interface{}

    // Accounting Port number (standard port 1646). The type is interface{} with
    // range: 1..65535.
    AcctPortNumber interface{}

    // IP address of RADIUS server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // A number that indicates the priority             of the server. The type is
    // interface{} with range: 0..4294967295.
    Priority interface{}

    // Per-server timeout in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TimeoutXr interface{}

    // Per-server retransmit. The type is interface{} with range: 0..4294967295.
    Retransmit interface{}

    // Per-server deadtime in minutes. The type is interface{} with range:
    // 0..4294967295. Units are minute.
    DeadTime interface{}

    // Per-server dead-detect time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    DeadDetectTime interface{}

    // Per-server dead-detect tries. The type is interface{} with range:
    // 0..4294967295.
    DeadDetectTries interface{}

    // Authentication port. The type is interface{} with range: 0..4294967295.
    AuthenticationPort interface{}

    // Accounting port. The type is interface{} with range: 0..4294967295.
    AccountingPort interface{}

    // State of the server UP/DOWN. The type is string.
    State interface{}

    // Elapsed time the server has been in              current state. The type is
    // interface{} with range: 0..4294967295.
    CurrentStateDuration interface{}

    // Elapsed time the server was been in              previous state. The type
    // is interface{} with range: 0..4294967295.
    PreviousStateDuration interface{}

    // Total number of incoming packets read. The type is interface{} with range:
    // 0..4294967295.
    PacketsIn interface{}

    // Total number of outgoing packets sent. The type is interface{} with range:
    // 0..4294967295.
    PacketsOut interface{}

    // Total number of packets timed-out. The type is interface{} with range:
    // 0..4294967295.
    Timeouts interface{}

    // Total number of requests aborted. The type is interface{} with range:
    // 0..4294967295.
    Aborts interface{}

    // Number of replies expected to arrive. The type is interface{} with range:
    // 0..4294967295.
    RepliesExpected interface{}

    // Number of requests redirected. The type is interface{} with range:
    // 0..4294967295.
    RedirectedRequests interface{}

    // Round-trip time for authentication               in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    AuthenticationRtt interface{}

    // Number of access requests. The type is interface{} with range:
    // 0..4294967295.
    AccessRequests interface{}

    // Number of retransmitted                          access requests. The type
    // is interface{} with range: 0..4294967295.
    AccessRequestRetransmits interface{}

    // Number of access accepts. The type is interface{} with range:
    // 0..4294967295.
    AccessAccepts interface{}

    // Number of access rejects. The type is interface{} with range:
    // 0..4294967295.
    AccessRejects interface{}

    // Number of access challenges. The type is interface{} with range:
    // 0..4294967295.
    AccessChallenges interface{}

    // Number of bad access responses. The type is interface{} with range:
    // 0..4294967295.
    BadAccessResponses interface{}

    // Number of bad access authenticators. The type is interface{} with range:
    // 0..4294967295.
    BadAccessAuthenticators interface{}

    // Number of pending access requests. The type is interface{} with range:
    // 0..4294967295.
    PendingAccessRequests interface{}

    // Number of access packets timed-out. The type is interface{} with range:
    // 0..4294967295.
    AccessTimeouts interface{}

    // Number of packets received with unknown          type from authentication
    // server. The type is interface{} with range: 0..4294967295.
    UnknownAccessTypes interface{}

    // Number of access responses dropped. The type is interface{} with range:
    // 0..4294967295.
    DroppedAccessResponses interface{}

    // No of throttled access reqs stats. The type is interface{} with range:
    // 0..4294967295.
    ThrottledAccessReqs interface{}

    // No of access reqs that is throttled is timedout. The type is interface{}
    // with range: 0..4294967295.
    ThrottledTimedOutReqs interface{}

    // No of discarded access reqs. The type is interface{} with range:
    // 0..4294967295.
    ThrottledDroppedReqs interface{}

    // Max throttled access reqs. The type is interface{} with range:
    // 0..4294967295.
    MaxThrottledAccessReqs interface{}

    // No of currently throttled access reqs. The type is interface{} with range:
    // 0..4294967295.
    CurrentlyThrottledAccessReqs interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AuthenResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthenTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthenTransactionFailure interface{}

    // Number of unexpected authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenUnexpectedResponses interface{}

    // Number of server error authentication responses. The type is interface{}
    // with range: 0..4294967295.
    AuthenServerErrorResponses interface{}

    // Number of incorrect authentication responses. The type is interface{} with
    // range: 0..4294967295.
    AuthenIncorrectResponses interface{}

    // Number of access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthorRequests interface{}

    // Number of access packets timed out. The type is interface{} with range:
    // 0..4294967295.
    AuthorRequestTimeouts interface{}

    // Average response time for authorization requests. The type is interface{}
    // with range: 0..4294967295.
    AuthorResponseTime interface{}

    // Number of succeeded authorization transactions. The type is interface{}
    // with range: 0..4294967295.
    AuthorTransactionSuccessess interface{}

    // Number of failed authorization transactions. The type is interface{} with
    // range: 0..4294967295.
    AuthorTransactionFailure interface{}

    // Number of unexpected authorization responses. The type is interface{} with
    // range: 0..4294967295.
    AuthorUnexpectedResponses interface{}

    // Number of server error authorization responses. The type is interface{}
    // with range: 0..4294967295.
    AuthorServerErrorResponses interface{}

    // Number of incorrect authorization responses. The type is interface{} with
    // range: 0..4294967295.
    AuthorIncorrectResponses interface{}

    // Round-trip time for accounting                   in milliseconds. The type
    // is interface{} with range: 0..4294967295. Units are millisecond.
    AccountingRtt interface{}

    // Number of accounting requests. The type is interface{} with range:
    // 0..4294967295.
    AccountingRequests interface{}

    // Number of retransmitted                          accounting requests. The
    // type is interface{} with range: 0..4294967295.
    AccountingRetransmits interface{}

    // Number of accounting responses. The type is interface{} with range:
    // 0..4294967295.
    AccountingResponses interface{}

    // Number of bad accounting responses. The type is interface{} with range:
    // 0..4294967295.
    BadAccountingResponses interface{}

    // Number of bad accounting                         authenticators. The type
    // is interface{} with range: 0..4294967295.
    BadAccountingAuthenticators interface{}

    // Number of pending accounting requests. The type is interface{} with range:
    // 0..4294967295.
    PendingAccountingRequets interface{}

    // Number of accounting packets                     timed-out. The type is
    // interface{} with range: 0..4294967295.
    AccountingTimeouts interface{}

    // Number of packets received with unknown          type from accounting
    // server. The type is interface{} with range: 0..4294967295.
    UnknownAccountingTypes interface{}

    // Number of accounting responses                   dropped. The type is
    // interface{} with range: 0..4294967295.
    DroppedAccountingResponses interface{}

    // Is a private server. The type is bool.
    IsAPrivateServer interface{}

    // Total auth test request. The type is interface{} with range: 0..4294967295.
    TotalTestAuthReqs interface{}

    // Total auth test timeouts. The type is interface{} with range:
    // 0..4294967295.
    TotalTestAuthTimeouts interface{}

    // Total auth test response. The type is interface{} with range:
    // 0..4294967295.
    TotalTestAuthResponse interface{}

    // Total auth test pending. The type is interface{} with range: 0..4294967295.
    TotalTestAuthPending interface{}

    // Total acct test req. The type is interface{} with range: 0..4294967295.
    TotalTestAcctReqs interface{}

    // Total acct test timeouts. The type is interface{} with range:
    // 0..4294967295.
    TotalTestAcctTimeouts interface{}

    // Total acct test response. The type is interface{} with range:
    // 0..4294967295.
    TotalTestAcctResponse interface{}

    // Total acct test pending. The type is interface{} with range: 0..4294967295.
    TotalTestAcctPending interface{}

    // No of throttled acct transactions stats. The type is interface{} with
    // range: 0..4294967295.
    ThrottledAcctTransactions interface{}

    // No of acct transaction that is throttled is timedout. The type is
    // interface{} with range: 0..4294967295.
    ThrottledAcctTimedOutStats interface{}

    // No of acct discarded transaction. The type is interface{} with range:
    // 0..4294967295.
    ThrottledAcctFailuresStats interface{}

    // Max throttled acct transactions. The type is interface{} with range:
    // 0..4294967295.
    MaxAcctThrottled interface{}

    // No of currently throttled acct transactions. The type is interface{} with
    // range: 0..4294967295.
    ThrottledaAcctTransactions interface{}

    // Number of unexpected accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctUnexpectedResponses interface{}

    // Number of server error accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctServerErrorResponses interface{}

    // Number of incorrect accounting responses. The type is interface{} with
    // range: 0..4294967295.
    AcctIncorrectResponses interface{}

    // Average response time for authentication requests. The type is interface{}
    // with range: 0..4294967295.
    AcctResponseTime interface{}

    // Number of succeeded authentication transactions. The type is interface{}
    // with range: 0..4294967295.
    AcctTransactionSuccessess interface{}

    // Number of failed authentication transactions. The type is interface{} with
    // range: 0..4294967295.
    AcctTransactionFailure interface{}

    // Total time of Server being in DEAD               state. The type is
    // interface{} with range: 0..4294967295.
    TotalDeadtime interface{}

    // Time of Server being in DEAD state,              after last UP. The type is
    // interface{} with range: 0..4294967295.
    LastDeadtime interface{}

    // flag to indicate Server is quarantined           or not (Automated TEST in
    // progress). The type is bool.
    IsQuarantined interface{}

    // Server group name for private server. The type is string.
    GroupName interface{}

    // IP address buffer. The type is string.
    IpAddressXr interface{}

    // IP address Family. The type is string.
    Family interface{}
}

func (server *Aaa_Radius_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", server.IpAddress}
    server.EntityData.Leafs["auth-port-number"] = types.YLeaf{"AuthPortNumber", server.AuthPortNumber}
    server.EntityData.Leafs["acct-port-number"] = types.YLeaf{"AcctPortNumber", server.AcctPortNumber}
    server.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", server.Ipv4Address}
    server.EntityData.Leafs["priority"] = types.YLeaf{"Priority", server.Priority}
    server.EntityData.Leafs["timeout-xr"] = types.YLeaf{"TimeoutXr", server.TimeoutXr}
    server.EntityData.Leafs["retransmit"] = types.YLeaf{"Retransmit", server.Retransmit}
    server.EntityData.Leafs["dead-time"] = types.YLeaf{"DeadTime", server.DeadTime}
    server.EntityData.Leafs["dead-detect-time"] = types.YLeaf{"DeadDetectTime", server.DeadDetectTime}
    server.EntityData.Leafs["dead-detect-tries"] = types.YLeaf{"DeadDetectTries", server.DeadDetectTries}
    server.EntityData.Leafs["authentication-port"] = types.YLeaf{"AuthenticationPort", server.AuthenticationPort}
    server.EntityData.Leafs["accounting-port"] = types.YLeaf{"AccountingPort", server.AccountingPort}
    server.EntityData.Leafs["state"] = types.YLeaf{"State", server.State}
    server.EntityData.Leafs["current-state-duration"] = types.YLeaf{"CurrentStateDuration", server.CurrentStateDuration}
    server.EntityData.Leafs["previous-state-duration"] = types.YLeaf{"PreviousStateDuration", server.PreviousStateDuration}
    server.EntityData.Leafs["packets-in"] = types.YLeaf{"PacketsIn", server.PacketsIn}
    server.EntityData.Leafs["packets-out"] = types.YLeaf{"PacketsOut", server.PacketsOut}
    server.EntityData.Leafs["timeouts"] = types.YLeaf{"Timeouts", server.Timeouts}
    server.EntityData.Leafs["aborts"] = types.YLeaf{"Aborts", server.Aborts}
    server.EntityData.Leafs["replies-expected"] = types.YLeaf{"RepliesExpected", server.RepliesExpected}
    server.EntityData.Leafs["redirected-requests"] = types.YLeaf{"RedirectedRequests", server.RedirectedRequests}
    server.EntityData.Leafs["authentication-rtt"] = types.YLeaf{"AuthenticationRtt", server.AuthenticationRtt}
    server.EntityData.Leafs["access-requests"] = types.YLeaf{"AccessRequests", server.AccessRequests}
    server.EntityData.Leafs["access-request-retransmits"] = types.YLeaf{"AccessRequestRetransmits", server.AccessRequestRetransmits}
    server.EntityData.Leafs["access-accepts"] = types.YLeaf{"AccessAccepts", server.AccessAccepts}
    server.EntityData.Leafs["access-rejects"] = types.YLeaf{"AccessRejects", server.AccessRejects}
    server.EntityData.Leafs["access-challenges"] = types.YLeaf{"AccessChallenges", server.AccessChallenges}
    server.EntityData.Leafs["bad-access-responses"] = types.YLeaf{"BadAccessResponses", server.BadAccessResponses}
    server.EntityData.Leafs["bad-access-authenticators"] = types.YLeaf{"BadAccessAuthenticators", server.BadAccessAuthenticators}
    server.EntityData.Leafs["pending-access-requests"] = types.YLeaf{"PendingAccessRequests", server.PendingAccessRequests}
    server.EntityData.Leafs["access-timeouts"] = types.YLeaf{"AccessTimeouts", server.AccessTimeouts}
    server.EntityData.Leafs["unknown-access-types"] = types.YLeaf{"UnknownAccessTypes", server.UnknownAccessTypes}
    server.EntityData.Leafs["dropped-access-responses"] = types.YLeaf{"DroppedAccessResponses", server.DroppedAccessResponses}
    server.EntityData.Leafs["throttled-access-reqs"] = types.YLeaf{"ThrottledAccessReqs", server.ThrottledAccessReqs}
    server.EntityData.Leafs["throttled-timed-out-reqs"] = types.YLeaf{"ThrottledTimedOutReqs", server.ThrottledTimedOutReqs}
    server.EntityData.Leafs["throttled-dropped-reqs"] = types.YLeaf{"ThrottledDroppedReqs", server.ThrottledDroppedReqs}
    server.EntityData.Leafs["max-throttled-access-reqs"] = types.YLeaf{"MaxThrottledAccessReqs", server.MaxThrottledAccessReqs}
    server.EntityData.Leafs["currently-throttled-access-reqs"] = types.YLeaf{"CurrentlyThrottledAccessReqs", server.CurrentlyThrottledAccessReqs}
    server.EntityData.Leafs["authen-response-time"] = types.YLeaf{"AuthenResponseTime", server.AuthenResponseTime}
    server.EntityData.Leafs["authen-transaction-successess"] = types.YLeaf{"AuthenTransactionSuccessess", server.AuthenTransactionSuccessess}
    server.EntityData.Leafs["authen-transaction-failure"] = types.YLeaf{"AuthenTransactionFailure", server.AuthenTransactionFailure}
    server.EntityData.Leafs["authen-unexpected-responses"] = types.YLeaf{"AuthenUnexpectedResponses", server.AuthenUnexpectedResponses}
    server.EntityData.Leafs["authen-server-error-responses"] = types.YLeaf{"AuthenServerErrorResponses", server.AuthenServerErrorResponses}
    server.EntityData.Leafs["authen-incorrect-responses"] = types.YLeaf{"AuthenIncorrectResponses", server.AuthenIncorrectResponses}
    server.EntityData.Leafs["author-requests"] = types.YLeaf{"AuthorRequests", server.AuthorRequests}
    server.EntityData.Leafs["author-request-timeouts"] = types.YLeaf{"AuthorRequestTimeouts", server.AuthorRequestTimeouts}
    server.EntityData.Leafs["author-response-time"] = types.YLeaf{"AuthorResponseTime", server.AuthorResponseTime}
    server.EntityData.Leafs["author-transaction-successess"] = types.YLeaf{"AuthorTransactionSuccessess", server.AuthorTransactionSuccessess}
    server.EntityData.Leafs["author-transaction-failure"] = types.YLeaf{"AuthorTransactionFailure", server.AuthorTransactionFailure}
    server.EntityData.Leafs["author-unexpected-responses"] = types.YLeaf{"AuthorUnexpectedResponses", server.AuthorUnexpectedResponses}
    server.EntityData.Leafs["author-server-error-responses"] = types.YLeaf{"AuthorServerErrorResponses", server.AuthorServerErrorResponses}
    server.EntityData.Leafs["author-incorrect-responses"] = types.YLeaf{"AuthorIncorrectResponses", server.AuthorIncorrectResponses}
    server.EntityData.Leafs["accounting-rtt"] = types.YLeaf{"AccountingRtt", server.AccountingRtt}
    server.EntityData.Leafs["accounting-requests"] = types.YLeaf{"AccountingRequests", server.AccountingRequests}
    server.EntityData.Leafs["accounting-retransmits"] = types.YLeaf{"AccountingRetransmits", server.AccountingRetransmits}
    server.EntityData.Leafs["accounting-responses"] = types.YLeaf{"AccountingResponses", server.AccountingResponses}
    server.EntityData.Leafs["bad-accounting-responses"] = types.YLeaf{"BadAccountingResponses", server.BadAccountingResponses}
    server.EntityData.Leafs["bad-accounting-authenticators"] = types.YLeaf{"BadAccountingAuthenticators", server.BadAccountingAuthenticators}
    server.EntityData.Leafs["pending-accounting-requets"] = types.YLeaf{"PendingAccountingRequets", server.PendingAccountingRequets}
    server.EntityData.Leafs["accounting-timeouts"] = types.YLeaf{"AccountingTimeouts", server.AccountingTimeouts}
    server.EntityData.Leafs["unknown-accounting-types"] = types.YLeaf{"UnknownAccountingTypes", server.UnknownAccountingTypes}
    server.EntityData.Leafs["dropped-accounting-responses"] = types.YLeaf{"DroppedAccountingResponses", server.DroppedAccountingResponses}
    server.EntityData.Leafs["is-a-private-server"] = types.YLeaf{"IsAPrivateServer", server.IsAPrivateServer}
    server.EntityData.Leafs["total-test-auth-reqs"] = types.YLeaf{"TotalTestAuthReqs", server.TotalTestAuthReqs}
    server.EntityData.Leafs["total-test-auth-timeouts"] = types.YLeaf{"TotalTestAuthTimeouts", server.TotalTestAuthTimeouts}
    server.EntityData.Leafs["total-test-auth-response"] = types.YLeaf{"TotalTestAuthResponse", server.TotalTestAuthResponse}
    server.EntityData.Leafs["total-test-auth-pending"] = types.YLeaf{"TotalTestAuthPending", server.TotalTestAuthPending}
    server.EntityData.Leafs["total-test-acct-reqs"] = types.YLeaf{"TotalTestAcctReqs", server.TotalTestAcctReqs}
    server.EntityData.Leafs["total-test-acct-timeouts"] = types.YLeaf{"TotalTestAcctTimeouts", server.TotalTestAcctTimeouts}
    server.EntityData.Leafs["total-test-acct-response"] = types.YLeaf{"TotalTestAcctResponse", server.TotalTestAcctResponse}
    server.EntityData.Leafs["total-test-acct-pending"] = types.YLeaf{"TotalTestAcctPending", server.TotalTestAcctPending}
    server.EntityData.Leafs["throttled-acct-transactions"] = types.YLeaf{"ThrottledAcctTransactions", server.ThrottledAcctTransactions}
    server.EntityData.Leafs["throttled-acct-timed-out-stats"] = types.YLeaf{"ThrottledAcctTimedOutStats", server.ThrottledAcctTimedOutStats}
    server.EntityData.Leafs["throttled-acct-failures-stats"] = types.YLeaf{"ThrottledAcctFailuresStats", server.ThrottledAcctFailuresStats}
    server.EntityData.Leafs["max-acct-throttled"] = types.YLeaf{"MaxAcctThrottled", server.MaxAcctThrottled}
    server.EntityData.Leafs["throttleda-acct-transactions"] = types.YLeaf{"ThrottledaAcctTransactions", server.ThrottledaAcctTransactions}
    server.EntityData.Leafs["acct-unexpected-responses"] = types.YLeaf{"AcctUnexpectedResponses", server.AcctUnexpectedResponses}
    server.EntityData.Leafs["acct-server-error-responses"] = types.YLeaf{"AcctServerErrorResponses", server.AcctServerErrorResponses}
    server.EntityData.Leafs["acct-incorrect-responses"] = types.YLeaf{"AcctIncorrectResponses", server.AcctIncorrectResponses}
    server.EntityData.Leafs["acct-response-time"] = types.YLeaf{"AcctResponseTime", server.AcctResponseTime}
    server.EntityData.Leafs["acct-transaction-successess"] = types.YLeaf{"AcctTransactionSuccessess", server.AcctTransactionSuccessess}
    server.EntityData.Leafs["acct-transaction-failure"] = types.YLeaf{"AcctTransactionFailure", server.AcctTransactionFailure}
    server.EntityData.Leafs["total-deadtime"] = types.YLeaf{"TotalDeadtime", server.TotalDeadtime}
    server.EntityData.Leafs["last-deadtime"] = types.YLeaf{"LastDeadtime", server.LastDeadtime}
    server.EntityData.Leafs["is-quarantined"] = types.YLeaf{"IsQuarantined", server.IsQuarantined}
    server.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", server.GroupName}
    server.EntityData.Leafs["ip-address-xr"] = types.YLeaf{"IpAddressXr", server.IpAddressXr}
    server.EntityData.Leafs["family"] = types.YLeaf{"Family", server.Family}
    return &(server.EntityData)
}

// Aaa_Radius_RadiusSourceInterface
// RADIUS source interfaces
type Aaa_Radius_RadiusSourceInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of source interfaces. The type is slice of
    // Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface.
    ListOfSourceInterface []Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetEntityData() *types.CommonEntityData {
    radiusSourceInterface.EntityData.YFilter = radiusSourceInterface.YFilter
    radiusSourceInterface.EntityData.YangName = "radius-source-interface"
    radiusSourceInterface.EntityData.BundleName = "cisco_ios_xr"
    radiusSourceInterface.EntityData.ParentYangName = "radius"
    radiusSourceInterface.EntityData.SegmentPath = "radius-source-interface"
    radiusSourceInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radiusSourceInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radiusSourceInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radiusSourceInterface.EntityData.Children = make(map[string]types.YChild)
    radiusSourceInterface.EntityData.Children["list-of-source-interface"] = types.YChild{"ListOfSourceInterface", nil}
    for i := range radiusSourceInterface.ListOfSourceInterface {
        radiusSourceInterface.EntityData.Children[types.GetSegmentPath(&radiusSourceInterface.ListOfSourceInterface[i])] = types.YChild{"ListOfSourceInterface", &radiusSourceInterface.ListOfSourceInterface[i]}
    }
    radiusSourceInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(radiusSourceInterface.EntityData)
}

// Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface
// List of source interfaces
type Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the source interface. The type is string.
    InterfaceName interface{}

    // IP address buffer. The type is string.
    Ipaddrv4 interface{}

    // IP address buffer. The type is string.
    Ipaddrv6 interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetEntityData() *types.CommonEntityData {
    listOfSourceInterface.EntityData.YFilter = listOfSourceInterface.YFilter
    listOfSourceInterface.EntityData.YangName = "list-of-source-interface"
    listOfSourceInterface.EntityData.BundleName = "cisco_ios_xr"
    listOfSourceInterface.EntityData.ParentYangName = "radius-source-interface"
    listOfSourceInterface.EntityData.SegmentPath = "list-of-source-interface"
    listOfSourceInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    listOfSourceInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    listOfSourceInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    listOfSourceInterface.EntityData.Children = make(map[string]types.YChild)
    listOfSourceInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    listOfSourceInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", listOfSourceInterface.InterfaceName}
    listOfSourceInterface.EntityData.Leafs["ipaddrv4"] = types.YLeaf{"Ipaddrv4", listOfSourceInterface.Ipaddrv4}
    listOfSourceInterface.EntityData.Leafs["ipaddrv6"] = types.YLeaf{"Ipaddrv6", listOfSourceInterface.Ipaddrv6}
    listOfSourceInterface.EntityData.Leafs["vrfid"] = types.YLeaf{"Vrfid", listOfSourceInterface.Vrfid}
    return &(listOfSourceInterface.EntityData)
}

// Aaa_Radius_Global
// RADIUS Client Information
type Aaa_Radius_Global struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of RADIUS Access-Responsepackets received from unknownaddresses. The
    // type is interface{} with range: 0..4294967295.
    UnknownAuthenticationResponse interface{}

    // NAS-Identifier of the RADIUSauthentication client. The type is string.
    AuthenticationNasId interface{}

    // Number of RADIUS Accounting-Responsepackets received from unknownaddresses.
    // The type is interface{} with range: 0..4294967295.
    UnknownAccountingResponse interface{}

    // NAS-Identifier of the RADIUSaccounting client. The type is string.
    AccountingNasId interface{}
}

func (global *Aaa_Radius_Global) GetEntityData() *types.CommonEntityData {
    global.EntityData.YFilter = global.YFilter
    global.EntityData.YangName = "global"
    global.EntityData.BundleName = "cisco_ios_xr"
    global.EntityData.ParentYangName = "radius"
    global.EntityData.SegmentPath = "global"
    global.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    global.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    global.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    global.EntityData.Children = make(map[string]types.YChild)
    global.EntityData.Leafs = make(map[string]types.YLeaf)
    global.EntityData.Leafs["unknown-authentication-response"] = types.YLeaf{"UnknownAuthenticationResponse", global.UnknownAuthenticationResponse}
    global.EntityData.Leafs["authentication-nas-id"] = types.YLeaf{"AuthenticationNasId", global.AuthenticationNasId}
    global.EntityData.Leafs["unknown-accounting-response"] = types.YLeaf{"UnknownAccountingResponse", global.UnknownAccountingResponse}
    global.EntityData.Leafs["accounting-nas-id"] = types.YLeaf{"AccountingNasId", global.AccountingNasId}
    return &(global.EntityData)
}

// Aaa_Tacacs
// TACACS operational data
type Aaa_Tacacs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TACACS Active Request List.
    Requests Aaa_Tacacs_Requests

    // TACACS server Information.
    Servers Aaa_Tacacs_Servers

    // TACACS sg Information.
    ServerGroups Aaa_Tacacs_ServerGroups
}

func (tacacs *Aaa_Tacacs) GetEntityData() *types.CommonEntityData {
    tacacs.EntityData.YFilter = tacacs.YFilter
    tacacs.EntityData.YangName = "tacacs"
    tacacs.EntityData.BundleName = "cisco_ios_xr"
    tacacs.EntityData.ParentYangName = "aaa"
    tacacs.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-tacacs-oper:tacacs"
    tacacs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacs.EntityData.Children = make(map[string]types.YChild)
    tacacs.EntityData.Children["requests"] = types.YChild{"Requests", &tacacs.Requests}
    tacacs.EntityData.Children["servers"] = types.YChild{"Servers", &tacacs.Servers}
    tacacs.EntityData.Children["server-groups"] = types.YChild{"ServerGroups", &tacacs.ServerGroups}
    tacacs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tacacs.EntityData)
}

// Aaa_Tacacs_Requests
// TACACS Active Request List
type Aaa_Tacacs_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // request. The type is slice of Aaa_Tacacs_Requests_Request.
    Request []Aaa_Tacacs_Requests_Request
}

func (requests *Aaa_Tacacs_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "tacacs"
    requests.EntityData.SegmentPath = "requests"
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = make(map[string]types.YChild)
    requests.EntityData.Children["request"] = types.YChild{"Request", nil}
    for i := range requests.Request {
        requests.EntityData.Children[types.GetSegmentPath(&requests.Request[i])] = types.YChild{"Request", &requests.Request[i]}
    }
    requests.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(requests.EntityData)
}

// Aaa_Tacacs_Requests_Request
// request
type Aaa_Tacacs_Requests_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tacacs requestbag. The type is slice of
    // Aaa_Tacacs_Requests_Request_TacacsRequestbag.
    TacacsRequestbag []Aaa_Tacacs_Requests_Request_TacacsRequestbag
}

func (request *Aaa_Tacacs_Requests_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "requests"
    request.EntityData.SegmentPath = "request"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Children["tacacs-requestbag"] = types.YChild{"TacacsRequestbag", nil}
    for i := range request.TacacsRequestbag {
        request.EntityData.Children[types.GetSegmentPath(&request.TacacsRequestbag[i])] = types.YChild{"TacacsRequestbag", &request.TacacsRequestbag[i]}
    }
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(request.EntityData)
}

// Aaa_Tacacs_Requests_Request_TacacsRequestbag
// tacacs requestbag
type Aaa_Tacacs_Requests_Request_TacacsRequestbag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // time remaining for this request. The type is interface{} with range:
    // 0..4294967295.
    TimeRemaining interface{}

    // bytes written. The type is interface{} with range: 0..4294967295. Units are
    // byte.
    BytesOut interface{}

    // size of the packet to be sent. The type is interface{} with range:
    // 0..4294967295.
    OutPakSize interface{}

    // bytes read from socket. The type is interface{} with range: 0..4294967295.
    // Units are byte.
    BytesIn interface{}

    // size of the packet to be received. The type is interface{} with range:
    // 0..4294967295.
    InPakSize interface{}

    // the type of packet. The type is string.
    PakType interface{}

    // same as in pkt hdr. The type is interface{} with range:
    // -2147483648..2147483647.
    SessionId interface{}

    // socket number. The type is interface{} with range: -2147483648..2147483647.
    Sock interface{}
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetEntityData() *types.CommonEntityData {
    tacacsRequestbag.EntityData.YFilter = tacacsRequestbag.YFilter
    tacacsRequestbag.EntityData.YangName = "tacacs-requestbag"
    tacacsRequestbag.EntityData.BundleName = "cisco_ios_xr"
    tacacsRequestbag.EntityData.ParentYangName = "request"
    tacacsRequestbag.EntityData.SegmentPath = "tacacs-requestbag"
    tacacsRequestbag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacsRequestbag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacsRequestbag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacsRequestbag.EntityData.Children = make(map[string]types.YChild)
    tacacsRequestbag.EntityData.Leafs = make(map[string]types.YLeaf)
    tacacsRequestbag.EntityData.Leafs["time-remaining"] = types.YLeaf{"TimeRemaining", tacacsRequestbag.TimeRemaining}
    tacacsRequestbag.EntityData.Leafs["bytes-out"] = types.YLeaf{"BytesOut", tacacsRequestbag.BytesOut}
    tacacsRequestbag.EntityData.Leafs["out-pak-size"] = types.YLeaf{"OutPakSize", tacacsRequestbag.OutPakSize}
    tacacsRequestbag.EntityData.Leafs["bytes-in"] = types.YLeaf{"BytesIn", tacacsRequestbag.BytesIn}
    tacacsRequestbag.EntityData.Leafs["in-pak-size"] = types.YLeaf{"InPakSize", tacacsRequestbag.InPakSize}
    tacacsRequestbag.EntityData.Leafs["pak-type"] = types.YLeaf{"PakType", tacacsRequestbag.PakType}
    tacacsRequestbag.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", tacacsRequestbag.SessionId}
    tacacsRequestbag.EntityData.Leafs["sock"] = types.YLeaf{"Sock", tacacsRequestbag.Sock}
    return &(tacacsRequestbag.EntityData)
}

// Aaa_Tacacs_Servers
// TACACS server Information
type Aaa_Tacacs_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // server. The type is slice of Aaa_Tacacs_Servers_Server.
    Server []Aaa_Tacacs_Servers_Server
}

func (servers *Aaa_Tacacs_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "tacacs"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = make(map[string]types.YChild)
    servers.EntityData.Children["server"] = types.YChild{"Server", nil}
    for i := range servers.Server {
        servers.EntityData.Children[types.GetSegmentPath(&servers.Server[i])] = types.YChild{"Server", &servers.Server[i]}
    }
    servers.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(servers.EntityData)
}

// Aaa_Tacacs_Servers_Server
// server
type Aaa_Tacacs_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // internet address of T+ server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr interface{}

    // per-server timeout. The type is interface{} with range: 0..4294967295.
    Timeout interface{}

    // per server port to use. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // # of bytes read. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    BytesIn interface{}

    // # of bytes out. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    BytesOut interface{}

    // socket closes. The type is interface{} with range: 0..4294967295.
    Closes interface{}

    // socket opens. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // error count. The type is interface{} with range: 0..4294967295.
    Errors interface{}

    // abort count. The type is interface{} with range: 0..4294967295.
    Aborts interface{}

    // # of incoming packets read. The type is interface{} with range:
    // 0..4294967295.
    PaksIn interface{}

    // # of outgoing packets sent. The type is interface{} with range:
    // 0..4294967295.
    PaksOut interface{}

    // # of replies expected to arrive. The type is interface{} with range:
    // 0..4294967295.
    RepliesExpected interface{}

    // is the server UP or down ?. The type is bool.
    Up interface{}

    // is the server connected ?. The type is bool.
    ConnUp interface{}

    // is this a single connect server ?. The type is bool.
    SingleConnect interface{}

    // is this a private server ?. The type is bool.
    IsPrivate interface{}

    // VRF in which server is reachable. The type is string with length: 0..33.
    VrfName interface{}

    // IP address buffer. The type is string with length: 0..46.
    AddrBuf interface{}

    // IP address Family. The type is string with length: 0..5.
    Family interface{}
}

func (server *Aaa_Tacacs_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["addr"] = types.YLeaf{"Addr", server.Addr}
    server.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", server.Timeout}
    server.EntityData.Leafs["port"] = types.YLeaf{"Port", server.Port}
    server.EntityData.Leafs["bytes-in"] = types.YLeaf{"BytesIn", server.BytesIn}
    server.EntityData.Leafs["bytes-out"] = types.YLeaf{"BytesOut", server.BytesOut}
    server.EntityData.Leafs["closes"] = types.YLeaf{"Closes", server.Closes}
    server.EntityData.Leafs["opens"] = types.YLeaf{"Opens", server.Opens}
    server.EntityData.Leafs["errors"] = types.YLeaf{"Errors", server.Errors}
    server.EntityData.Leafs["aborts"] = types.YLeaf{"Aborts", server.Aborts}
    server.EntityData.Leafs["paks-in"] = types.YLeaf{"PaksIn", server.PaksIn}
    server.EntityData.Leafs["paks-out"] = types.YLeaf{"PaksOut", server.PaksOut}
    server.EntityData.Leafs["replies-expected"] = types.YLeaf{"RepliesExpected", server.RepliesExpected}
    server.EntityData.Leafs["up"] = types.YLeaf{"Up", server.Up}
    server.EntityData.Leafs["conn-up"] = types.YLeaf{"ConnUp", server.ConnUp}
    server.EntityData.Leafs["single-connect"] = types.YLeaf{"SingleConnect", server.SingleConnect}
    server.EntityData.Leafs["is-private"] = types.YLeaf{"IsPrivate", server.IsPrivate}
    server.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", server.VrfName}
    server.EntityData.Leafs["addr-buf"] = types.YLeaf{"AddrBuf", server.AddrBuf}
    server.EntityData.Leafs["family"] = types.YLeaf{"Family", server.Family}
    return &(server.EntityData)
}

// Aaa_Tacacs_ServerGroups
// TACACS sg Information
type Aaa_Tacacs_ServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // server group. The type is slice of Aaa_Tacacs_ServerGroups_ServerGroup.
    ServerGroup []Aaa_Tacacs_ServerGroups_ServerGroup
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetEntityData() *types.CommonEntityData {
    serverGroups.EntityData.YFilter = serverGroups.YFilter
    serverGroups.EntityData.YangName = "server-groups"
    serverGroups.EntityData.BundleName = "cisco_ios_xr"
    serverGroups.EntityData.ParentYangName = "tacacs"
    serverGroups.EntityData.SegmentPath = "server-groups"
    serverGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroups.EntityData.Children = make(map[string]types.YChild)
    serverGroups.EntityData.Children["server-group"] = types.YChild{"ServerGroup", nil}
    for i := range serverGroups.ServerGroup {
        serverGroups.EntityData.Children[types.GetSegmentPath(&serverGroups.ServerGroup[i])] = types.YChild{"ServerGroup", &serverGroups.ServerGroup[i]}
    }
    serverGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(serverGroups.EntityData)
}

// Aaa_Tacacs_ServerGroups_ServerGroup
// server group
type Aaa_Tacacs_ServerGroups_ServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // name of the server group. The type is string.
    GroupName interface{}

    // server group mapped number. The type is interface{} with range:
    // 0..4294967295.
    SgMapNum interface{}

    // vrf of the group. The type is string with length: 0..33.
    VrfName interface{}

    // list of servers in this group. The type is slice of
    // Aaa_Tacacs_ServerGroups_ServerGroup_Server.
    Server []Aaa_Tacacs_ServerGroups_ServerGroup_Server
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetEntityData() *types.CommonEntityData {
    serverGroup.EntityData.YFilter = serverGroup.YFilter
    serverGroup.EntityData.YangName = "server-group"
    serverGroup.EntityData.BundleName = "cisco_ios_xr"
    serverGroup.EntityData.ParentYangName = "server-groups"
    serverGroup.EntityData.SegmentPath = "server-group"
    serverGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroup.EntityData.Children = make(map[string]types.YChild)
    serverGroup.EntityData.Children["server"] = types.YChild{"Server", nil}
    for i := range serverGroup.Server {
        serverGroup.EntityData.Children[types.GetSegmentPath(&serverGroup.Server[i])] = types.YChild{"Server", &serverGroup.Server[i]}
    }
    serverGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    serverGroup.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", serverGroup.GroupName}
    serverGroup.EntityData.Leafs["sg-map-num"] = types.YLeaf{"SgMapNum", serverGroup.SgMapNum}
    serverGroup.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", serverGroup.VrfName}
    return &(serverGroup.EntityData)
}

// Aaa_Tacacs_ServerGroups_ServerGroup_Server
// list of servers in this group
type Aaa_Tacacs_ServerGroups_ServerGroup_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // internet address of T+ server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Addr interface{}

    // per-server timeout. The type is interface{} with range: 0..4294967295.
    Timeout interface{}

    // per server port to use. The type is interface{} with range: 0..4294967295.
    Port interface{}

    // # of bytes read. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    BytesIn interface{}

    // # of bytes out. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    BytesOut interface{}

    // socket closes. The type is interface{} with range: 0..4294967295.
    Closes interface{}

    // socket opens. The type is interface{} with range: 0..4294967295.
    Opens interface{}

    // error count. The type is interface{} with range: 0..4294967295.
    Errors interface{}

    // abort count. The type is interface{} with range: 0..4294967295.
    Aborts interface{}

    // # of incoming packets read. The type is interface{} with range:
    // 0..4294967295.
    PaksIn interface{}

    // # of outgoing packets sent. The type is interface{} with range:
    // 0..4294967295.
    PaksOut interface{}

    // # of replies expected to arrive. The type is interface{} with range:
    // 0..4294967295.
    RepliesExpected interface{}

    // is the server UP or down ?. The type is bool.
    Up interface{}

    // is the server connected ?. The type is bool.
    ConnUp interface{}

    // is this a single connect server ?. The type is bool.
    SingleConnect interface{}

    // is this a private server ?. The type is bool.
    IsPrivate interface{}

    // VRF in which server is reachable. The type is string with length: 0..33.
    VrfName interface{}

    // IP address buffer. The type is string with length: 0..46.
    AddrBuf interface{}

    // IP address Family. The type is string with length: 0..5.
    Family interface{}
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "server-group"
    server.EntityData.SegmentPath = "server"
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = make(map[string]types.YChild)
    server.EntityData.Leafs = make(map[string]types.YLeaf)
    server.EntityData.Leafs["addr"] = types.YLeaf{"Addr", server.Addr}
    server.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", server.Timeout}
    server.EntityData.Leafs["port"] = types.YLeaf{"Port", server.Port}
    server.EntityData.Leafs["bytes-in"] = types.YLeaf{"BytesIn", server.BytesIn}
    server.EntityData.Leafs["bytes-out"] = types.YLeaf{"BytesOut", server.BytesOut}
    server.EntityData.Leafs["closes"] = types.YLeaf{"Closes", server.Closes}
    server.EntityData.Leafs["opens"] = types.YLeaf{"Opens", server.Opens}
    server.EntityData.Leafs["errors"] = types.YLeaf{"Errors", server.Errors}
    server.EntityData.Leafs["aborts"] = types.YLeaf{"Aborts", server.Aborts}
    server.EntityData.Leafs["paks-in"] = types.YLeaf{"PaksIn", server.PaksIn}
    server.EntityData.Leafs["paks-out"] = types.YLeaf{"PaksOut", server.PaksOut}
    server.EntityData.Leafs["replies-expected"] = types.YLeaf{"RepliesExpected", server.RepliesExpected}
    server.EntityData.Leafs["up"] = types.YLeaf{"Up", server.Up}
    server.EntityData.Leafs["conn-up"] = types.YLeaf{"ConnUp", server.ConnUp}
    server.EntityData.Leafs["single-connect"] = types.YLeaf{"SingleConnect", server.SingleConnect}
    server.EntityData.Leafs["is-private"] = types.YLeaf{"IsPrivate", server.IsPrivate}
    server.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", server.VrfName}
    server.EntityData.Leafs["addr-buf"] = types.YLeaf{"AddrBuf", server.AddrBuf}
    server.EntityData.Leafs["family"] = types.YLeaf{"Family", server.Family}
    return &(server.EntityData)
}

