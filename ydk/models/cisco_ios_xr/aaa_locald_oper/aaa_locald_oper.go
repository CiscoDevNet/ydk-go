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
    parent types.Entity
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

func (aaa *Aaa) GetFilter() yfilter.YFilter { return aaa.YFilter }

func (aaa *Aaa) SetFilter(yf yfilter.YFilter) { aaa.YFilter = yf }

func (aaa *Aaa) GetGoName(yname string) string {
    if yname == "all-tasks" { return "AllTasks" }
    if yname == "currentuser-detail" { return "CurrentuserDetail" }
    if yname == "task-map" { return "TaskMap" }
    if yname == "taskgroups" { return "Taskgroups" }
    if yname == "users" { return "Users" }
    if yname == "usergroups" { return "Usergroups" }
    if yname == "authen-method" { return "AuthenMethod" }
    if yname == "current-usergroup" { return "CurrentUsergroup" }
    if yname == "Cisco-IOS-XR-aaa-diameter-oper:diameter" { return "Diameter" }
    if yname == "Cisco-IOS-XR-aaa-protocol-radius-oper:radius" { return "Radius" }
    if yname == "Cisco-IOS-XR-aaa-tacacs-oper:tacacs" { return "Tacacs" }
    return ""
}

func (aaa *Aaa) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-locald-oper:aaa"
}

func (aaa *Aaa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-tasks" {
        return &aaa.AllTasks
    }
    if childYangName == "currentuser-detail" {
        return &aaa.CurrentuserDetail
    }
    if childYangName == "task-map" {
        return &aaa.TaskMap
    }
    if childYangName == "taskgroups" {
        return &aaa.Taskgroups
    }
    if childYangName == "users" {
        return &aaa.Users
    }
    if childYangName == "usergroups" {
        return &aaa.Usergroups
    }
    if childYangName == "authen-method" {
        return &aaa.AuthenMethod
    }
    if childYangName == "current-usergroup" {
        return &aaa.CurrentUsergroup
    }
    if childYangName == "Cisco-IOS-XR-aaa-diameter-oper:diameter" {
        return &aaa.Diameter
    }
    if childYangName == "Cisco-IOS-XR-aaa-protocol-radius-oper:radius" {
        return &aaa.Radius
    }
    if childYangName == "Cisco-IOS-XR-aaa-tacacs-oper:tacacs" {
        return &aaa.Tacacs
    }
    return nil
}

func (aaa *Aaa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-tasks"] = &aaa.AllTasks
    children["currentuser-detail"] = &aaa.CurrentuserDetail
    children["task-map"] = &aaa.TaskMap
    children["taskgroups"] = &aaa.Taskgroups
    children["users"] = &aaa.Users
    children["usergroups"] = &aaa.Usergroups
    children["authen-method"] = &aaa.AuthenMethod
    children["current-usergroup"] = &aaa.CurrentUsergroup
    children["Cisco-IOS-XR-aaa-diameter-oper:diameter"] = &aaa.Diameter
    children["Cisco-IOS-XR-aaa-protocol-radius-oper:radius"] = &aaa.Radius
    children["Cisco-IOS-XR-aaa-tacacs-oper:tacacs"] = &aaa.Tacacs
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

func (aaa *Aaa) GetParentYangName() string { return "Cisco-IOS-XR-aaa-locald-oper" }

// Aaa_AllTasks
// All tasks supported by system
type Aaa_AllTasks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Names of available task-ids. The type is slice of string.
    TaskId []interface{}
}

func (allTasks *Aaa_AllTasks) GetFilter() yfilter.YFilter { return allTasks.YFilter }

func (allTasks *Aaa_AllTasks) SetFilter(yf yfilter.YFilter) { allTasks.YFilter = yf }

func (allTasks *Aaa_AllTasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    return ""
}

func (allTasks *Aaa_AllTasks) GetSegmentPath() string {
    return "all-tasks"
}

func (allTasks *Aaa_AllTasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allTasks *Aaa_AllTasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allTasks *Aaa_AllTasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = allTasks.TaskId
    return leafs
}

func (allTasks *Aaa_AllTasks) GetBundleName() string { return "cisco_ios_xr" }

func (allTasks *Aaa_AllTasks) GetYangName() string { return "all-tasks" }

func (allTasks *Aaa_AllTasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allTasks *Aaa_AllTasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allTasks *Aaa_AllTasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allTasks *Aaa_AllTasks) SetParent(parent types.Entity) { allTasks.parent = parent }

func (allTasks *Aaa_AllTasks) GetParent() types.Entity { return allTasks.parent }

func (allTasks *Aaa_AllTasks) GetParentYangName() string { return "aaa" }

// Aaa_CurrentuserDetail
// Current user specific details
type Aaa_CurrentuserDetail struct {
    parent types.Entity
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

func (currentuserDetail *Aaa_CurrentuserDetail) GetFilter() yfilter.YFilter { return currentuserDetail.YFilter }

func (currentuserDetail *Aaa_CurrentuserDetail) SetFilter(yf yfilter.YFilter) { currentuserDetail.YFilter = yf }

func (currentuserDetail *Aaa_CurrentuserDetail) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "authenmethod" { return "Authenmethod" }
    if yname == "usergroup" { return "Usergroup" }
    if yname == "taskmap" { return "Taskmap" }
    return ""
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetSegmentPath() string {
    return "currentuser-detail"
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = currentuserDetail.Name
    leafs["authenmethod"] = currentuserDetail.Authenmethod
    leafs["usergroup"] = currentuserDetail.Usergroup
    leafs["taskmap"] = currentuserDetail.Taskmap
    return leafs
}

func (currentuserDetail *Aaa_CurrentuserDetail) GetBundleName() string { return "cisco_ios_xr" }

func (currentuserDetail *Aaa_CurrentuserDetail) GetYangName() string { return "currentuser-detail" }

func (currentuserDetail *Aaa_CurrentuserDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentuserDetail *Aaa_CurrentuserDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentuserDetail *Aaa_CurrentuserDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentuserDetail *Aaa_CurrentuserDetail) SetParent(parent types.Entity) { currentuserDetail.parent = parent }

func (currentuserDetail *Aaa_CurrentuserDetail) GetParent() types.Entity { return currentuserDetail.parent }

func (currentuserDetail *Aaa_CurrentuserDetail) GetParentYangName() string { return "aaa" }

// Aaa_TaskMap
// Task map of current user
type Aaa_TaskMap struct {
    parent types.Entity
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

func (taskMap *Aaa_TaskMap) GetFilter() yfilter.YFilter { return taskMap.YFilter }

func (taskMap *Aaa_TaskMap) SetFilter(yf yfilter.YFilter) { taskMap.YFilter = yf }

func (taskMap *Aaa_TaskMap) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "authenmethod" { return "Authenmethod" }
    if yname == "usergroup" { return "Usergroup" }
    if yname == "taskmap" { return "Taskmap" }
    return ""
}

func (taskMap *Aaa_TaskMap) GetSegmentPath() string {
    return "task-map"
}

func (taskMap *Aaa_TaskMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (taskMap *Aaa_TaskMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (taskMap *Aaa_TaskMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = taskMap.Name
    leafs["authenmethod"] = taskMap.Authenmethod
    leafs["usergroup"] = taskMap.Usergroup
    leafs["taskmap"] = taskMap.Taskmap
    return leafs
}

func (taskMap *Aaa_TaskMap) GetBundleName() string { return "cisco_ios_xr" }

func (taskMap *Aaa_TaskMap) GetYangName() string { return "task-map" }

func (taskMap *Aaa_TaskMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskMap *Aaa_TaskMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskMap *Aaa_TaskMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskMap *Aaa_TaskMap) SetParent(parent types.Entity) { taskMap.parent = parent }

func (taskMap *Aaa_TaskMap) GetParent() types.Entity { return taskMap.parent }

func (taskMap *Aaa_TaskMap) GetParentYangName() string { return "aaa" }

// Aaa_Taskgroups
// Individual taskgroups container
type Aaa_Taskgroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specific Taskgroup Information. The type is slice of
    // Aaa_Taskgroups_Taskgroup.
    Taskgroup []Aaa_Taskgroups_Taskgroup
}

func (taskgroups *Aaa_Taskgroups) GetFilter() yfilter.YFilter { return taskgroups.YFilter }

func (taskgroups *Aaa_Taskgroups) SetFilter(yf yfilter.YFilter) { taskgroups.YFilter = yf }

func (taskgroups *Aaa_Taskgroups) GetGoName(yname string) string {
    if yname == "taskgroup" { return "Taskgroup" }
    return ""
}

func (taskgroups *Aaa_Taskgroups) GetSegmentPath() string {
    return "taskgroups"
}

func (taskgroups *Aaa_Taskgroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "taskgroup" {
        for _, c := range taskgroups.Taskgroup {
            if taskgroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup{}
        taskgroups.Taskgroup = append(taskgroups.Taskgroup, child)
        return &taskgroups.Taskgroup[len(taskgroups.Taskgroup)-1]
    }
    return nil
}

func (taskgroups *Aaa_Taskgroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskgroups.Taskgroup {
        children[taskgroups.Taskgroup[i].GetSegmentPath()] = &taskgroups.Taskgroup[i]
    }
    return children
}

func (taskgroups *Aaa_Taskgroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskgroups *Aaa_Taskgroups) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroups *Aaa_Taskgroups) GetYangName() string { return "taskgroups" }

func (taskgroups *Aaa_Taskgroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroups *Aaa_Taskgroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroups *Aaa_Taskgroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroups *Aaa_Taskgroups) SetParent(parent types.Entity) { taskgroups.parent = parent }

func (taskgroups *Aaa_Taskgroups) GetParent() types.Entity { return taskgroups.parent }

func (taskgroups *Aaa_Taskgroups) GetParentYangName() string { return "aaa" }

// Aaa_Taskgroups_Taskgroup
// Specific Taskgroup Information
type Aaa_Taskgroups_Taskgroup struct {
    parent types.Entity
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

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetFilter() yfilter.YFilter { return taskgroup.YFilter }

func (taskgroup *Aaa_Taskgroups_Taskgroup) SetFilter(yf yfilter.YFilter) { taskgroup.YFilter = yf }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "name-xr" { return "NameXr" }
    if yname == "included-task-ids" { return "IncludedTaskIds" }
    if yname == "task-map" { return "TaskMap" }
    return ""
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetSegmentPath() string {
    return "taskgroup" + "[name='" + fmt.Sprintf("%v", taskgroup.Name) + "']"
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "included-task-ids" {
        return &taskgroup.IncludedTaskIds
    }
    if childYangName == "task-map" {
        return &taskgroup.TaskMap
    }
    return nil
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["included-task-ids"] = &taskgroup.IncludedTaskIds
    children["task-map"] = &taskgroup.TaskMap
    return children
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = taskgroup.Name
    leafs["name-xr"] = taskgroup.NameXr
    return leafs
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetYangName() string { return "taskgroup" }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroup *Aaa_Taskgroups_Taskgroup) SetParent(parent types.Entity) { taskgroup.parent = parent }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetParent() types.Entity { return taskgroup.parent }

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetParentYangName() string { return "taskgroups" }

// Aaa_Taskgroups_Taskgroup_IncludedTaskIds
// Task-ids included
type Aaa_Taskgroups_Taskgroup_IncludedTaskIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks.
    Tasks []Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetFilter() yfilter.YFilter { return includedTaskIds.YFilter }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) SetFilter(yf yfilter.YFilter) { includedTaskIds.YFilter = yf }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetSegmentPath() string {
    return "included-task-ids"
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range includedTaskIds.Tasks {
            if includedTaskIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks{}
        includedTaskIds.Tasks = append(includedTaskIds.Tasks, child)
        return &includedTaskIds.Tasks[len(includedTaskIds.Tasks)-1]
    }
    return nil
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range includedTaskIds.Tasks {
        children[includedTaskIds.Tasks[i].GetSegmentPath()] = &includedTaskIds.Tasks[i]
    }
    return children
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetBundleName() string { return "cisco_ios_xr" }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetYangName() string { return "included-task-ids" }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) SetParent(parent types.Entity) { includedTaskIds.parent = parent }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetParent() types.Entity { return includedTaskIds.parent }

func (includedTaskIds *Aaa_Taskgroups_Taskgroup_IncludedTaskIds) GetParentYangName() string { return "taskgroup" }

// Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks
// List of permitted tasks
type Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Taskgroups_Taskgroup_IncludedTaskIds_Tasks) GetParentYangName() string { return "included-task-ids" }

// Aaa_Taskgroups_Taskgroup_TaskMap
// Computed task map
type Aaa_Taskgroups_Taskgroup_TaskMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Taskgroups_Taskgroup_TaskMap_Tasks.
    Tasks []Aaa_Taskgroups_Taskgroup_TaskMap_Tasks
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetFilter() yfilter.YFilter { return taskMap.YFilter }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) SetFilter(yf yfilter.YFilter) { taskMap.YFilter = yf }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetSegmentPath() string {
    return "task-map"
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range taskMap.Tasks {
            if taskMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Taskgroups_Taskgroup_TaskMap_Tasks{}
        taskMap.Tasks = append(taskMap.Tasks, child)
        return &taskMap.Tasks[len(taskMap.Tasks)-1]
    }
    return nil
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskMap.Tasks {
        children[taskMap.Tasks[i].GetSegmentPath()] = &taskMap.Tasks[i]
    }
    return children
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetBundleName() string { return "cisco_ios_xr" }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetYangName() string { return "task-map" }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) SetParent(parent types.Entity) { taskMap.parent = parent }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetParent() types.Entity { return taskMap.parent }

func (taskMap *Aaa_Taskgroups_Taskgroup_TaskMap) GetParentYangName() string { return "taskgroup" }

// Aaa_Taskgroups_Taskgroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Taskgroups_Taskgroup_TaskMap_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Taskgroups_Taskgroup_TaskMap_Tasks) GetParentYangName() string { return "task-map" }

// Aaa_Users
// Container for individual local user information
type Aaa_Users struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specific local user information. The type is slice of Aaa_Users_User.
    User []Aaa_Users_User
}

func (users *Aaa_Users) GetFilter() yfilter.YFilter { return users.YFilter }

func (users *Aaa_Users) SetFilter(yf yfilter.YFilter) { users.YFilter = yf }

func (users *Aaa_Users) GetGoName(yname string) string {
    if yname == "user" { return "User" }
    return ""
}

func (users *Aaa_Users) GetSegmentPath() string {
    return "users"
}

func (users *Aaa_Users) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user" {
        for _, c := range users.User {
            if users.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Users_User{}
        users.User = append(users.User, child)
        return &users.User[len(users.User)-1]
    }
    return nil
}

func (users *Aaa_Users) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range users.User {
        children[users.User[i].GetSegmentPath()] = &users.User[i]
    }
    return children
}

func (users *Aaa_Users) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (users *Aaa_Users) GetBundleName() string { return "cisco_ios_xr" }

func (users *Aaa_Users) GetYangName() string { return "users" }

func (users *Aaa_Users) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (users *Aaa_Users) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (users *Aaa_Users) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (users *Aaa_Users) SetParent(parent types.Entity) { users.parent = parent }

func (users *Aaa_Users) GetParent() types.Entity { return users.parent }

func (users *Aaa_Users) GetParentYangName() string { return "aaa" }

// Aaa_Users_User
// Specific local user information
type Aaa_Users_User struct {
    parent types.Entity
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

func (user *Aaa_Users_User) GetFilter() yfilter.YFilter { return user.YFilter }

func (user *Aaa_Users_User) SetFilter(yf yfilter.YFilter) { user.YFilter = yf }

func (user *Aaa_Users_User) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "name-xr" { return "NameXr" }
    if yname == "admin-user" { return "AdminUser" }
    if yname == "first-user" { return "FirstUser" }
    if yname == "usergroup" { return "Usergroup" }
    if yname == "task-map" { return "TaskMap" }
    return ""
}

func (user *Aaa_Users_User) GetSegmentPath() string {
    return "user" + "[name='" + fmt.Sprintf("%v", user.Name) + "']"
}

func (user *Aaa_Users_User) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "task-map" {
        return &user.TaskMap
    }
    return nil
}

func (user *Aaa_Users_User) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["task-map"] = &user.TaskMap
    return children
}

func (user *Aaa_Users_User) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = user.Name
    leafs["name-xr"] = user.NameXr
    leafs["admin-user"] = user.AdminUser
    leafs["first-user"] = user.FirstUser
    leafs["usergroup"] = user.Usergroup
    return leafs
}

func (user *Aaa_Users_User) GetBundleName() string { return "cisco_ios_xr" }

func (user *Aaa_Users_User) GetYangName() string { return "user" }

func (user *Aaa_Users_User) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (user *Aaa_Users_User) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (user *Aaa_Users_User) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (user *Aaa_Users_User) SetParent(parent types.Entity) { user.parent = parent }

func (user *Aaa_Users_User) GetParent() types.Entity { return user.parent }

func (user *Aaa_Users_User) GetParentYangName() string { return "users" }

// Aaa_Users_User_TaskMap
// Computed taskmap
type Aaa_Users_User_TaskMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of Aaa_Users_User_TaskMap_Tasks.
    Tasks []Aaa_Users_User_TaskMap_Tasks
}

func (taskMap *Aaa_Users_User_TaskMap) GetFilter() yfilter.YFilter { return taskMap.YFilter }

func (taskMap *Aaa_Users_User_TaskMap) SetFilter(yf yfilter.YFilter) { taskMap.YFilter = yf }

func (taskMap *Aaa_Users_User_TaskMap) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (taskMap *Aaa_Users_User_TaskMap) GetSegmentPath() string {
    return "task-map"
}

func (taskMap *Aaa_Users_User_TaskMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range taskMap.Tasks {
            if taskMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Users_User_TaskMap_Tasks{}
        taskMap.Tasks = append(taskMap.Tasks, child)
        return &taskMap.Tasks[len(taskMap.Tasks)-1]
    }
    return nil
}

func (taskMap *Aaa_Users_User_TaskMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskMap.Tasks {
        children[taskMap.Tasks[i].GetSegmentPath()] = &taskMap.Tasks[i]
    }
    return children
}

func (taskMap *Aaa_Users_User_TaskMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskMap *Aaa_Users_User_TaskMap) GetBundleName() string { return "cisco_ios_xr" }

func (taskMap *Aaa_Users_User_TaskMap) GetYangName() string { return "task-map" }

func (taskMap *Aaa_Users_User_TaskMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskMap *Aaa_Users_User_TaskMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskMap *Aaa_Users_User_TaskMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskMap *Aaa_Users_User_TaskMap) SetParent(parent types.Entity) { taskMap.parent = parent }

func (taskMap *Aaa_Users_User_TaskMap) GetParent() types.Entity { return taskMap.parent }

func (taskMap *Aaa_Users_User_TaskMap) GetParentYangName() string { return "user" }

// Aaa_Users_User_TaskMap_Tasks
// List of permitted tasks
type Aaa_Users_User_TaskMap_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Users_User_TaskMap_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Users_User_TaskMap_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Users_User_TaskMap_Tasks) GetParentYangName() string { return "task-map" }

// Aaa_Usergroups
// Container for individual usergroup Information
type Aaa_Usergroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specific Usergroup Information. The type is slice of
    // Aaa_Usergroups_Usergroup.
    Usergroup []Aaa_Usergroups_Usergroup
}

func (usergroups *Aaa_Usergroups) GetFilter() yfilter.YFilter { return usergroups.YFilter }

func (usergroups *Aaa_Usergroups) SetFilter(yf yfilter.YFilter) { usergroups.YFilter = yf }

func (usergroups *Aaa_Usergroups) GetGoName(yname string) string {
    if yname == "usergroup" { return "Usergroup" }
    return ""
}

func (usergroups *Aaa_Usergroups) GetSegmentPath() string {
    return "usergroups"
}

func (usergroups *Aaa_Usergroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "usergroup" {
        for _, c := range usergroups.Usergroup {
            if usergroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup{}
        usergroups.Usergroup = append(usergroups.Usergroup, child)
        return &usergroups.Usergroup[len(usergroups.Usergroup)-1]
    }
    return nil
}

func (usergroups *Aaa_Usergroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range usergroups.Usergroup {
        children[usergroups.Usergroup[i].GetSegmentPath()] = &usergroups.Usergroup[i]
    }
    return children
}

func (usergroups *Aaa_Usergroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (usergroups *Aaa_Usergroups) GetBundleName() string { return "cisco_ios_xr" }

func (usergroups *Aaa_Usergroups) GetYangName() string { return "usergroups" }

func (usergroups *Aaa_Usergroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroups *Aaa_Usergroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroups *Aaa_Usergroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroups *Aaa_Usergroups) SetParent(parent types.Entity) { usergroups.parent = parent }

func (usergroups *Aaa_Usergroups) GetParent() types.Entity { return usergroups.parent }

func (usergroups *Aaa_Usergroups) GetParentYangName() string { return "aaa" }

// Aaa_Usergroups_Usergroup
// Specific Usergroup Information
type Aaa_Usergroups_Usergroup struct {
    parent types.Entity
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

func (usergroup *Aaa_Usergroups_Usergroup) GetFilter() yfilter.YFilter { return usergroup.YFilter }

func (usergroup *Aaa_Usergroups_Usergroup) SetFilter(yf yfilter.YFilter) { usergroup.YFilter = yf }

func (usergroup *Aaa_Usergroups_Usergroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "name-xr" { return "NameXr" }
    if yname == "task-map" { return "TaskMap" }
    if yname == "taskgroup" { return "Taskgroup" }
    return ""
}

func (usergroup *Aaa_Usergroups_Usergroup) GetSegmentPath() string {
    return "usergroup" + "[name='" + fmt.Sprintf("%v", usergroup.Name) + "']"
}

func (usergroup *Aaa_Usergroups_Usergroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "task-map" {
        return &usergroup.TaskMap
    }
    if childYangName == "taskgroup" {
        for _, c := range usergroup.Taskgroup {
            if usergroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_Taskgroup{}
        usergroup.Taskgroup = append(usergroup.Taskgroup, child)
        return &usergroup.Taskgroup[len(usergroup.Taskgroup)-1]
    }
    return nil
}

func (usergroup *Aaa_Usergroups_Usergroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["task-map"] = &usergroup.TaskMap
    for i := range usergroup.Taskgroup {
        children[usergroup.Taskgroup[i].GetSegmentPath()] = &usergroup.Taskgroup[i]
    }
    return children
}

func (usergroup *Aaa_Usergroups_Usergroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = usergroup.Name
    leafs["name-xr"] = usergroup.NameXr
    return leafs
}

func (usergroup *Aaa_Usergroups_Usergroup) GetBundleName() string { return "cisco_ios_xr" }

func (usergroup *Aaa_Usergroups_Usergroup) GetYangName() string { return "usergroup" }

func (usergroup *Aaa_Usergroups_Usergroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (usergroup *Aaa_Usergroups_Usergroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (usergroup *Aaa_Usergroups_Usergroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (usergroup *Aaa_Usergroups_Usergroup) SetParent(parent types.Entity) { usergroup.parent = parent }

func (usergroup *Aaa_Usergroups_Usergroup) GetParent() types.Entity { return usergroup.parent }

func (usergroup *Aaa_Usergroups_Usergroup) GetParentYangName() string { return "usergroups" }

// Aaa_Usergroups_Usergroup_TaskMap
// Computed task map
type Aaa_Usergroups_Usergroup_TaskMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_TaskMap_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_TaskMap_Tasks
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetFilter() yfilter.YFilter { return taskMap.YFilter }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) SetFilter(yf yfilter.YFilter) { taskMap.YFilter = yf }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetSegmentPath() string {
    return "task-map"
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range taskMap.Tasks {
            if taskMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_TaskMap_Tasks{}
        taskMap.Tasks = append(taskMap.Tasks, child)
        return &taskMap.Tasks[len(taskMap.Tasks)-1]
    }
    return nil
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskMap.Tasks {
        children[taskMap.Tasks[i].GetSegmentPath()] = &taskMap.Tasks[i]
    }
    return children
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetBundleName() string { return "cisco_ios_xr" }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetYangName() string { return "task-map" }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) SetParent(parent types.Entity) { taskMap.parent = parent }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetParent() types.Entity { return taskMap.parent }

func (taskMap *Aaa_Usergroups_Usergroup_TaskMap) GetParentYangName() string { return "usergroup" }

// Aaa_Usergroups_Usergroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_TaskMap_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Usergroups_Usergroup_TaskMap_Tasks) GetParentYangName() string { return "task-map" }

// Aaa_Usergroups_Usergroup_Taskgroup
// Component taskgroups
type Aaa_Usergroups_Usergroup_Taskgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the taskgroup. The type is string.
    NameXr interface{}

    // Task-ids included.
    IncludedTaskIds Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds

    // Computed task map.
    TaskMap Aaa_Usergroups_Usergroup_Taskgroup_TaskMap
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetFilter() yfilter.YFilter { return taskgroup.YFilter }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) SetFilter(yf yfilter.YFilter) { taskgroup.YFilter = yf }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetGoName(yname string) string {
    if yname == "name-xr" { return "NameXr" }
    if yname == "included-task-ids" { return "IncludedTaskIds" }
    if yname == "task-map" { return "TaskMap" }
    return ""
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetSegmentPath() string {
    return "taskgroup"
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "included-task-ids" {
        return &taskgroup.IncludedTaskIds
    }
    if childYangName == "task-map" {
        return &taskgroup.TaskMap
    }
    return nil
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["included-task-ids"] = &taskgroup.IncludedTaskIds
    children["task-map"] = &taskgroup.TaskMap
    return children
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name-xr"] = taskgroup.NameXr
    return leafs
}

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetBundleName() string { return "cisco_ios_xr" }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetYangName() string { return "taskgroup" }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) SetParent(parent types.Entity) { taskgroup.parent = parent }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetParent() types.Entity { return taskgroup.parent }

func (taskgroup *Aaa_Usergroups_Usergroup_Taskgroup) GetParentYangName() string { return "usergroup" }

// Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds
// Task-ids included
type Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetFilter() yfilter.YFilter { return includedTaskIds.YFilter }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) SetFilter(yf yfilter.YFilter) { includedTaskIds.YFilter = yf }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetSegmentPath() string {
    return "included-task-ids"
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range includedTaskIds.Tasks {
            if includedTaskIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks{}
        includedTaskIds.Tasks = append(includedTaskIds.Tasks, child)
        return &includedTaskIds.Tasks[len(includedTaskIds.Tasks)-1]
    }
    return nil
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range includedTaskIds.Tasks {
        children[includedTaskIds.Tasks[i].GetSegmentPath()] = &includedTaskIds.Tasks[i]
    }
    return children
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetBundleName() string { return "cisco_ios_xr" }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetYangName() string { return "included-task-ids" }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) SetParent(parent types.Entity) { includedTaskIds.parent = parent }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetParent() types.Entity { return includedTaskIds.parent }

func (includedTaskIds *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds) GetParentYangName() string { return "taskgroup" }

// Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_IncludedTaskIds_Tasks) GetParentYangName() string { return "included-task-ids" }

// Aaa_Usergroups_Usergroup_Taskgroup_TaskMap
// Computed task map
type Aaa_Usergroups_Usergroup_Taskgroup_TaskMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of permitted tasks. The type is slice of
    // Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks.
    Tasks []Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetFilter() yfilter.YFilter { return taskMap.YFilter }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) SetFilter(yf yfilter.YFilter) { taskMap.YFilter = yf }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetGoName(yname string) string {
    if yname == "tasks" { return "Tasks" }
    return ""
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetSegmentPath() string {
    return "task-map"
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tasks" {
        for _, c := range taskMap.Tasks {
            if taskMap.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks{}
        taskMap.Tasks = append(taskMap.Tasks, child)
        return &taskMap.Tasks[len(taskMap.Tasks)-1]
    }
    return nil
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range taskMap.Tasks {
        children[taskMap.Tasks[i].GetSegmentPath()] = &taskMap.Tasks[i]
    }
    return children
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetBundleName() string { return "cisco_ios_xr" }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetYangName() string { return "task-map" }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) SetParent(parent types.Entity) { taskMap.parent = parent }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetParent() types.Entity { return taskMap.parent }

func (taskMap *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap) GetParentYangName() string { return "taskgroup" }

// Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks
// List of permitted tasks
type Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks struct {
    parent types.Entity
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

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetFilter() yfilter.YFilter { return tasks.YFilter }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) SetFilter(yf yfilter.YFilter) { tasks.YFilter = yf }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetGoName(yname string) string {
    if yname == "task-id" { return "TaskId" }
    if yname == "read" { return "Read" }
    if yname == "write" { return "Write" }
    if yname == "execute" { return "Execute" }
    if yname == "debug" { return "Debug" }
    return ""
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetSegmentPath() string {
    return "tasks"
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["task-id"] = tasks.TaskId
    leafs["read"] = tasks.Read
    leafs["write"] = tasks.Write
    leafs["execute"] = tasks.Execute
    leafs["debug"] = tasks.Debug
    return leafs
}

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetBundleName() string { return "cisco_ios_xr" }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetYangName() string { return "tasks" }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) SetParent(parent types.Entity) { tasks.parent = parent }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetParent() types.Entity { return tasks.parent }

func (tasks *Aaa_Usergroups_Usergroup_Taskgroup_TaskMap_Tasks) GetParentYangName() string { return "task-map" }

// Aaa_AuthenMethod
// Current users authentication method
type Aaa_AuthenMethod struct {
    parent types.Entity
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

func (authenMethod *Aaa_AuthenMethod) GetFilter() yfilter.YFilter { return authenMethod.YFilter }

func (authenMethod *Aaa_AuthenMethod) SetFilter(yf yfilter.YFilter) { authenMethod.YFilter = yf }

func (authenMethod *Aaa_AuthenMethod) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "authenmethod" { return "Authenmethod" }
    if yname == "usergroup" { return "Usergroup" }
    if yname == "taskmap" { return "Taskmap" }
    return ""
}

func (authenMethod *Aaa_AuthenMethod) GetSegmentPath() string {
    return "authen-method"
}

func (authenMethod *Aaa_AuthenMethod) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenMethod *Aaa_AuthenMethod) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenMethod *Aaa_AuthenMethod) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = authenMethod.Name
    leafs["authenmethod"] = authenMethod.Authenmethod
    leafs["usergroup"] = authenMethod.Usergroup
    leafs["taskmap"] = authenMethod.Taskmap
    return leafs
}

func (authenMethod *Aaa_AuthenMethod) GetBundleName() string { return "cisco_ios_xr" }

func (authenMethod *Aaa_AuthenMethod) GetYangName() string { return "authen-method" }

func (authenMethod *Aaa_AuthenMethod) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenMethod *Aaa_AuthenMethod) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenMethod *Aaa_AuthenMethod) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenMethod *Aaa_AuthenMethod) SetParent(parent types.Entity) { authenMethod.parent = parent }

func (authenMethod *Aaa_AuthenMethod) GetParent() types.Entity { return authenMethod.parent }

func (authenMethod *Aaa_AuthenMethod) GetParentYangName() string { return "aaa" }

// Aaa_CurrentUsergroup
// Specific Usergroup Information
type Aaa_CurrentUsergroup struct {
    parent types.Entity
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

func (currentUsergroup *Aaa_CurrentUsergroup) GetFilter() yfilter.YFilter { return currentUsergroup.YFilter }

func (currentUsergroup *Aaa_CurrentUsergroup) SetFilter(yf yfilter.YFilter) { currentUsergroup.YFilter = yf }

func (currentUsergroup *Aaa_CurrentUsergroup) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "authenmethod" { return "Authenmethod" }
    if yname == "usergroup" { return "Usergroup" }
    if yname == "taskmap" { return "Taskmap" }
    return ""
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetSegmentPath() string {
    return "current-usergroup"
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = currentUsergroup.Name
    leafs["authenmethod"] = currentUsergroup.Authenmethod
    leafs["usergroup"] = currentUsergroup.Usergroup
    leafs["taskmap"] = currentUsergroup.Taskmap
    return leafs
}

func (currentUsergroup *Aaa_CurrentUsergroup) GetBundleName() string { return "cisco_ios_xr" }

func (currentUsergroup *Aaa_CurrentUsergroup) GetYangName() string { return "current-usergroup" }

func (currentUsergroup *Aaa_CurrentUsergroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (currentUsergroup *Aaa_CurrentUsergroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (currentUsergroup *Aaa_CurrentUsergroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (currentUsergroup *Aaa_CurrentUsergroup) SetParent(parent types.Entity) { currentUsergroup.parent = parent }

func (currentUsergroup *Aaa_CurrentUsergroup) GetParent() types.Entity { return currentUsergroup.parent }

func (currentUsergroup *Aaa_CurrentUsergroup) GetParentYangName() string { return "aaa" }

// Aaa_Diameter
// Diameter operational data
type Aaa_Diameter struct {
    parent types.Entity
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

func (diameter *Aaa_Diameter) GetFilter() yfilter.YFilter { return diameter.YFilter }

func (diameter *Aaa_Diameter) SetFilter(yf yfilter.YFilter) { diameter.YFilter = yf }

func (diameter *Aaa_Diameter) GetGoName(yname string) string {
    if yname == "gy" { return "Gy" }
    if yname == "gx-statistics" { return "GxStatistics" }
    if yname == "gx" { return "Gx" }
    if yname == "peers" { return "Peers" }
    if yname == "nas" { return "Nas" }
    if yname == "nas-summary" { return "NasSummary" }
    if yname == "gy-session-ids" { return "GySessionIds" }
    if yname == "gy-statistics" { return "GyStatistics" }
    if yname == "gx-session-ids" { return "GxSessionIds" }
    if yname == "nas-session" { return "NasSession" }
    return ""
}

func (diameter *Aaa_Diameter) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-diameter-oper:diameter"
}

func (diameter *Aaa_Diameter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gy" {
        return &diameter.Gy
    }
    if childYangName == "gx-statistics" {
        return &diameter.GxStatistics
    }
    if childYangName == "gx" {
        return &diameter.Gx
    }
    if childYangName == "peers" {
        return &diameter.Peers
    }
    if childYangName == "nas" {
        return &diameter.Nas
    }
    if childYangName == "nas-summary" {
        return &diameter.NasSummary
    }
    if childYangName == "gy-session-ids" {
        return &diameter.GySessionIds
    }
    if childYangName == "gy-statistics" {
        return &diameter.GyStatistics
    }
    if childYangName == "gx-session-ids" {
        return &diameter.GxSessionIds
    }
    if childYangName == "nas-session" {
        return &diameter.NasSession
    }
    return nil
}

func (diameter *Aaa_Diameter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["gy"] = &diameter.Gy
    children["gx-statistics"] = &diameter.GxStatistics
    children["gx"] = &diameter.Gx
    children["peers"] = &diameter.Peers
    children["nas"] = &diameter.Nas
    children["nas-summary"] = &diameter.NasSummary
    children["gy-session-ids"] = &diameter.GySessionIds
    children["gy-statistics"] = &diameter.GyStatistics
    children["gx-session-ids"] = &diameter.GxSessionIds
    children["nas-session"] = &diameter.NasSession
    return children
}

func (diameter *Aaa_Diameter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diameter *Aaa_Diameter) GetBundleName() string { return "cisco_ios_xr" }

func (diameter *Aaa_Diameter) GetYangName() string { return "diameter" }

func (diameter *Aaa_Diameter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diameter *Aaa_Diameter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diameter *Aaa_Diameter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diameter *Aaa_Diameter) SetParent(parent types.Entity) { diameter.parent = parent }

func (diameter *Aaa_Diameter) GetParent() types.Entity { return diameter.parent }

func (diameter *Aaa_Diameter) GetParentYangName() string { return "aaa" }

// Aaa_Diameter_Gy
// Diameter global gy data
type Aaa_Diameter_Gy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Gy state. The type is bool.
    IsEnabled interface{}

    // Gy transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TxTimer interface{}

    // Gy retransmit count. The type is interface{} with range: 0..4294967295.
    Retransmits interface{}
}

func (gy *Aaa_Diameter_Gy) GetFilter() yfilter.YFilter { return gy.YFilter }

func (gy *Aaa_Diameter_Gy) SetFilter(yf yfilter.YFilter) { gy.YFilter = yf }

func (gy *Aaa_Diameter_Gy) GetGoName(yname string) string {
    if yname == "is-enabled" { return "IsEnabled" }
    if yname == "tx-timer" { return "TxTimer" }
    if yname == "retransmits" { return "Retransmits" }
    return ""
}

func (gy *Aaa_Diameter_Gy) GetSegmentPath() string {
    return "gy"
}

func (gy *Aaa_Diameter_Gy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gy *Aaa_Diameter_Gy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gy *Aaa_Diameter_Gy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-enabled"] = gy.IsEnabled
    leafs["tx-timer"] = gy.TxTimer
    leafs["retransmits"] = gy.Retransmits
    return leafs
}

func (gy *Aaa_Diameter_Gy) GetBundleName() string { return "cisco_ios_xr" }

func (gy *Aaa_Diameter_Gy) GetYangName() string { return "gy" }

func (gy *Aaa_Diameter_Gy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gy *Aaa_Diameter_Gy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gy *Aaa_Diameter_Gy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gy *Aaa_Diameter_Gy) SetParent(parent types.Entity) { gy.parent = parent }

func (gy *Aaa_Diameter_Gy) GetParent() types.Entity { return gy.parent }

func (gy *Aaa_Diameter_Gy) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_GxStatistics
// Diameter Gx Statistics data
type Aaa_Diameter_GxStatistics struct {
    parent types.Entity
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

func (gxStatistics *Aaa_Diameter_GxStatistics) GetFilter() yfilter.YFilter { return gxStatistics.YFilter }

func (gxStatistics *Aaa_Diameter_GxStatistics) SetFilter(yf yfilter.YFilter) { gxStatistics.YFilter = yf }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetGoName(yname string) string {
    if yname == "ccr-init-messages" { return "CcrInitMessages" }
    if yname == "ccr-init-failed-messages" { return "CcrInitFailedMessages" }
    if yname == "ccr-init-timed-out-messages" { return "CcrInitTimedOutMessages" }
    if yname == "ccr-init-retry-messages" { return "CcrInitRetryMessages" }
    if yname == "ccr-update-messages" { return "CcrUpdateMessages" }
    if yname == "ccr-update-failed-messages" { return "CcrUpdateFailedMessages" }
    if yname == "ccr-update-timed-out-messages" { return "CcrUpdateTimedOutMessages" }
    if yname == "ccr-update-retry-messages" { return "CcrUpdateRetryMessages" }
    if yname == "ccr-final-messages" { return "CcrFinalMessages" }
    if yname == "ccr-final-failed-messages" { return "CcrFinalFailedMessages" }
    if yname == "ccr-final-timed-out-messages" { return "CcrFinalTimedOutMessages" }
    if yname == "ccr-final-retry-messages" { return "CcrFinalRetryMessages" }
    if yname == "cca-init-messages" { return "CcaInitMessages" }
    if yname == "cca-init-error-messages" { return "CcaInitErrorMessages" }
    if yname == "cca-update-messages" { return "CcaUpdateMessages" }
    if yname == "cca-update-error-messages" { return "CcaUpdateErrorMessages" }
    if yname == "cca-final-messages" { return "CcaFinalMessages" }
    if yname == "cca-final-error-messages" { return "CcaFinalErrorMessages" }
    if yname == "rar-received-messages" { return "RarReceivedMessages" }
    if yname == "rar-received-error-messages" { return "RarReceivedErrorMessages" }
    if yname == "raa-sent-messages" { return "RaaSentMessages" }
    if yname == "raa-sent-error-messages" { return "RaaSentErrorMessages" }
    if yname == "asr-received-messages" { return "AsrReceivedMessages" }
    if yname == "asr-received-error-messages" { return "AsrReceivedErrorMessages" }
    if yname == "asa-sent-messsages" { return "AsaSentMesssages" }
    if yname == "asa-sent-error-messages" { return "AsaSentErrorMessages" }
    if yname == "session-termination-messages" { return "SessionTerminationMessages" }
    if yname == "unknown-request-messages" { return "UnknownRequestMessages" }
    if yname == "restore-sessions" { return "RestoreSessions" }
    if yname == "open-sessions" { return "OpenSessions" }
    if yname == "close-sessions" { return "CloseSessions" }
    if yname == "active-sessions" { return "ActiveSessions" }
    return ""
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetSegmentPath() string {
    return "gx-statistics"
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccr-init-messages"] = gxStatistics.CcrInitMessages
    leafs["ccr-init-failed-messages"] = gxStatistics.CcrInitFailedMessages
    leafs["ccr-init-timed-out-messages"] = gxStatistics.CcrInitTimedOutMessages
    leafs["ccr-init-retry-messages"] = gxStatistics.CcrInitRetryMessages
    leafs["ccr-update-messages"] = gxStatistics.CcrUpdateMessages
    leafs["ccr-update-failed-messages"] = gxStatistics.CcrUpdateFailedMessages
    leafs["ccr-update-timed-out-messages"] = gxStatistics.CcrUpdateTimedOutMessages
    leafs["ccr-update-retry-messages"] = gxStatistics.CcrUpdateRetryMessages
    leafs["ccr-final-messages"] = gxStatistics.CcrFinalMessages
    leafs["ccr-final-failed-messages"] = gxStatistics.CcrFinalFailedMessages
    leafs["ccr-final-timed-out-messages"] = gxStatistics.CcrFinalTimedOutMessages
    leafs["ccr-final-retry-messages"] = gxStatistics.CcrFinalRetryMessages
    leafs["cca-init-messages"] = gxStatistics.CcaInitMessages
    leafs["cca-init-error-messages"] = gxStatistics.CcaInitErrorMessages
    leafs["cca-update-messages"] = gxStatistics.CcaUpdateMessages
    leafs["cca-update-error-messages"] = gxStatistics.CcaUpdateErrorMessages
    leafs["cca-final-messages"] = gxStatistics.CcaFinalMessages
    leafs["cca-final-error-messages"] = gxStatistics.CcaFinalErrorMessages
    leafs["rar-received-messages"] = gxStatistics.RarReceivedMessages
    leafs["rar-received-error-messages"] = gxStatistics.RarReceivedErrorMessages
    leafs["raa-sent-messages"] = gxStatistics.RaaSentMessages
    leafs["raa-sent-error-messages"] = gxStatistics.RaaSentErrorMessages
    leafs["asr-received-messages"] = gxStatistics.AsrReceivedMessages
    leafs["asr-received-error-messages"] = gxStatistics.AsrReceivedErrorMessages
    leafs["asa-sent-messsages"] = gxStatistics.AsaSentMesssages
    leafs["asa-sent-error-messages"] = gxStatistics.AsaSentErrorMessages
    leafs["session-termination-messages"] = gxStatistics.SessionTerminationMessages
    leafs["unknown-request-messages"] = gxStatistics.UnknownRequestMessages
    leafs["restore-sessions"] = gxStatistics.RestoreSessions
    leafs["open-sessions"] = gxStatistics.OpenSessions
    leafs["close-sessions"] = gxStatistics.CloseSessions
    leafs["active-sessions"] = gxStatistics.ActiveSessions
    return leafs
}

func (gxStatistics *Aaa_Diameter_GxStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetYangName() string { return "gx-statistics" }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gxStatistics *Aaa_Diameter_GxStatistics) SetParent(parent types.Entity) { gxStatistics.parent = parent }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetParent() types.Entity { return gxStatistics.parent }

func (gxStatistics *Aaa_Diameter_GxStatistics) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Gx
// Diameter global gx data
type Aaa_Diameter_Gx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Gx state. The type is bool.
    IsEnabled interface{}

    // Gx transaction timer in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TxTimer interface{}

    // Gx retransmit count. The type is interface{} with range: 0..4294967295.
    Retransmits interface{}
}

func (gx *Aaa_Diameter_Gx) GetFilter() yfilter.YFilter { return gx.YFilter }

func (gx *Aaa_Diameter_Gx) SetFilter(yf yfilter.YFilter) { gx.YFilter = yf }

func (gx *Aaa_Diameter_Gx) GetGoName(yname string) string {
    if yname == "is-enabled" { return "IsEnabled" }
    if yname == "tx-timer" { return "TxTimer" }
    if yname == "retransmits" { return "Retransmits" }
    return ""
}

func (gx *Aaa_Diameter_Gx) GetSegmentPath() string {
    return "gx"
}

func (gx *Aaa_Diameter_Gx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gx *Aaa_Diameter_Gx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gx *Aaa_Diameter_Gx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-enabled"] = gx.IsEnabled
    leafs["tx-timer"] = gx.TxTimer
    leafs["retransmits"] = gx.Retransmits
    return leafs
}

func (gx *Aaa_Diameter_Gx) GetBundleName() string { return "cisco_ios_xr" }

func (gx *Aaa_Diameter_Gx) GetYangName() string { return "gx" }

func (gx *Aaa_Diameter_Gx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gx *Aaa_Diameter_Gx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gx *Aaa_Diameter_Gx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gx *Aaa_Diameter_Gx) SetParent(parent types.Entity) { gx.parent = parent }

func (gx *Aaa_Diameter_Gx) GetParent() types.Entity { return gx.parent }

func (gx *Aaa_Diameter_Gx) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Peers
// Diameter peer global data
type Aaa_Diameter_Peers struct {
    parent types.Entity
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

func (peers *Aaa_Diameter_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *Aaa_Diameter_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *Aaa_Diameter_Peers) GetGoName(yname string) string {
    if yname == "origin-host" { return "OriginHost" }
    if yname == "origin-realm" { return "OriginRealm" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "tls-trustpoint" { return "TlsTrustpoint" }
    if yname == "conn-retry-timer" { return "ConnRetryTimer" }
    if yname == "watchdog-timer" { return "WatchdogTimer" }
    if yname == "transaction-timer" { return "TransactionTimer" }
    if yname == "trans-total" { return "TransTotal" }
    if yname == "trans-max" { return "TransMax" }
    if yname == "peer" { return "Peer" }
    return ""
}

func (peers *Aaa_Diameter_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *Aaa_Diameter_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer" {
        for _, c := range peers.Peer {
            if peers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_Peers_Peer{}
        peers.Peer = append(peers.Peer, child)
        return &peers.Peer[len(peers.Peer)-1]
    }
    return nil
}

func (peers *Aaa_Diameter_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peers.Peer {
        children[peers.Peer[i].GetSegmentPath()] = &peers.Peer[i]
    }
    return children
}

func (peers *Aaa_Diameter_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["origin-host"] = peers.OriginHost
    leafs["origin-realm"] = peers.OriginRealm
    leafs["source-interface"] = peers.SourceInterface
    leafs["tls-trustpoint"] = peers.TlsTrustpoint
    leafs["conn-retry-timer"] = peers.ConnRetryTimer
    leafs["watchdog-timer"] = peers.WatchdogTimer
    leafs["transaction-timer"] = peers.TransactionTimer
    leafs["trans-total"] = peers.TransTotal
    leafs["trans-max"] = peers.TransMax
    return leafs
}

func (peers *Aaa_Diameter_Peers) GetBundleName() string { return "cisco_ios_xr" }

func (peers *Aaa_Diameter_Peers) GetYangName() string { return "peers" }

func (peers *Aaa_Diameter_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peers *Aaa_Diameter_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peers *Aaa_Diameter_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peers *Aaa_Diameter_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *Aaa_Diameter_Peers) GetParent() types.Entity { return peers.parent }

func (peers *Aaa_Diameter_Peers) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Peers_Peer
// Peer List
type Aaa_Diameter_Peers_Peer struct {
    parent types.Entity
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

func (peer *Aaa_Diameter_Peers_Peer) GetFilter() yfilter.YFilter { return peer.YFilter }

func (peer *Aaa_Diameter_Peers_Peer) SetFilter(yf yfilter.YFilter) { peer.YFilter = yf }

func (peer *Aaa_Diameter_Peers_Peer) GetGoName(yname string) string {
    if yname == "peer-name" { return "PeerName" }
    if yname == "peer-index" { return "PeerIndex" }
    if yname == "address" { return "Address" }
    if yname == "port" { return "Port" }
    if yname == "port-connect" { return "PortConnect" }
    if yname == "protocol-type" { return "ProtocolType" }
    if yname == "security-type" { return "SecurityType" }
    if yname == "conn-retry-timer" { return "ConnRetryTimer" }
    if yname == "watchdog-timer" { return "WatchdogTimer" }
    if yname == "transaction-timer" { return "TransactionTimer" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "destination-host" { return "DestinationHost" }
    if yname == "destination-realm" { return "DestinationRealm" }
    if yname == "peer-type" { return "PeerType" }
    if yname == "firmware-revision" { return "FirmwareRevision" }
    if yname == "state-duration" { return "StateDuration" }
    if yname == "last-disconnect-cause" { return "LastDisconnectCause" }
    if yname == "who-init-disconnect" { return "WhoInitDisconnect" }
    if yname == "in-as-rs" { return "InAsRs" }
    if yname == "out-as-rs" { return "OutAsRs" }
    if yname == "in-as-as" { return "InAsAs" }
    if yname == "out-as-as" { return "OutAsAs" }
    if yname == "in-ac-rs" { return "InAcRs" }
    if yname == "out-ac-rs" { return "OutAcRs" }
    if yname == "in-ac-as" { return "InAcAs" }
    if yname == "out-ac-as" { return "OutAcAs" }
    if yname == "in-ce-rs" { return "InCeRs" }
    if yname == "out-ce-rs" { return "OutCeRs" }
    if yname == "in-ce-as" { return "InCeAs" }
    if yname == "out-ce-as" { return "OutCeAs" }
    if yname == "in-dw-rs" { return "InDwRs" }
    if yname == "out-dw-rs" { return "OutDwRs" }
    if yname == "in-dw-as" { return "InDwAs" }
    if yname == "out-dw-as" { return "OutDwAs" }
    if yname == "in-dp-rs" { return "InDpRs" }
    if yname == "out-dp-rs" { return "OutDpRs" }
    if yname == "in-dp-as" { return "InDpAs" }
    if yname == "out-dp-as" { return "OutDpAs" }
    if yname == "in-ra-rs" { return "InRaRs" }
    if yname == "out-ra-rs" { return "OutRaRs" }
    if yname == "in-ra-as" { return "InRaAs" }
    if yname == "out-ra-as" { return "OutRaAs" }
    if yname == "in-st-rs" { return "InStRs" }
    if yname == "out-st-rs" { return "OutStRs" }
    if yname == "in-st-as" { return "InStAs" }
    if yname == "out-st-as" { return "OutStAs" }
    if yname == "in-cc-rs" { return "InCcRs" }
    if yname == "out-cc-rs" { return "OutCcRs" }
    if yname == "in-cc-as" { return "InCcAs" }
    if yname == "out-cc-as" { return "OutCcAs" }
    if yname == "out-aa-rs" { return "OutAaRs" }
    if yname == "in-aa-as" { return "InAaAs" }
    if yname == "malformed-requests" { return "MalformedRequests" }
    if yname == "received-proto-errors" { return "ReceivedProtoErrors" }
    if yname == "sent-proto-errors" { return "SentProtoErrors" }
    if yname == "received-transient-fails" { return "ReceivedTransientFails" }
    if yname == "sent-transient-fails" { return "SentTransientFails" }
    if yname == "received-permanent-fails" { return "ReceivedPermanentFails" }
    if yname == "sent-permanent-fails" { return "SentPermanentFails" }
    if yname == "transport-down" { return "TransportDown" }
    if yname == "state" { return "State" }
    return ""
}

func (peer *Aaa_Diameter_Peers_Peer) GetSegmentPath() string {
    return "peer"
}

func (peer *Aaa_Diameter_Peers_Peer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peer *Aaa_Diameter_Peers_Peer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peer *Aaa_Diameter_Peers_Peer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-name"] = peer.PeerName
    leafs["peer-index"] = peer.PeerIndex
    leafs["address"] = peer.Address
    leafs["port"] = peer.Port
    leafs["port-connect"] = peer.PortConnect
    leafs["protocol-type"] = peer.ProtocolType
    leafs["security-type"] = peer.SecurityType
    leafs["conn-retry-timer"] = peer.ConnRetryTimer
    leafs["watchdog-timer"] = peer.WatchdogTimer
    leafs["transaction-timer"] = peer.TransactionTimer
    leafs["vrf-name"] = peer.VrfName
    leafs["source-interface"] = peer.SourceInterface
    leafs["destination-host"] = peer.DestinationHost
    leafs["destination-realm"] = peer.DestinationRealm
    leafs["peer-type"] = peer.PeerType
    leafs["firmware-revision"] = peer.FirmwareRevision
    leafs["state-duration"] = peer.StateDuration
    leafs["last-disconnect-cause"] = peer.LastDisconnectCause
    leafs["who-init-disconnect"] = peer.WhoInitDisconnect
    leafs["in-as-rs"] = peer.InAsRs
    leafs["out-as-rs"] = peer.OutAsRs
    leafs["in-as-as"] = peer.InAsAs
    leafs["out-as-as"] = peer.OutAsAs
    leafs["in-ac-rs"] = peer.InAcRs
    leafs["out-ac-rs"] = peer.OutAcRs
    leafs["in-ac-as"] = peer.InAcAs
    leafs["out-ac-as"] = peer.OutAcAs
    leafs["in-ce-rs"] = peer.InCeRs
    leafs["out-ce-rs"] = peer.OutCeRs
    leafs["in-ce-as"] = peer.InCeAs
    leafs["out-ce-as"] = peer.OutCeAs
    leafs["in-dw-rs"] = peer.InDwRs
    leafs["out-dw-rs"] = peer.OutDwRs
    leafs["in-dw-as"] = peer.InDwAs
    leafs["out-dw-as"] = peer.OutDwAs
    leafs["in-dp-rs"] = peer.InDpRs
    leafs["out-dp-rs"] = peer.OutDpRs
    leafs["in-dp-as"] = peer.InDpAs
    leafs["out-dp-as"] = peer.OutDpAs
    leafs["in-ra-rs"] = peer.InRaRs
    leafs["out-ra-rs"] = peer.OutRaRs
    leafs["in-ra-as"] = peer.InRaAs
    leafs["out-ra-as"] = peer.OutRaAs
    leafs["in-st-rs"] = peer.InStRs
    leafs["out-st-rs"] = peer.OutStRs
    leafs["in-st-as"] = peer.InStAs
    leafs["out-st-as"] = peer.OutStAs
    leafs["in-cc-rs"] = peer.InCcRs
    leafs["out-cc-rs"] = peer.OutCcRs
    leafs["in-cc-as"] = peer.InCcAs
    leafs["out-cc-as"] = peer.OutCcAs
    leafs["out-aa-rs"] = peer.OutAaRs
    leafs["in-aa-as"] = peer.InAaAs
    leafs["malformed-requests"] = peer.MalformedRequests
    leafs["received-proto-errors"] = peer.ReceivedProtoErrors
    leafs["sent-proto-errors"] = peer.SentProtoErrors
    leafs["received-transient-fails"] = peer.ReceivedTransientFails
    leafs["sent-transient-fails"] = peer.SentTransientFails
    leafs["received-permanent-fails"] = peer.ReceivedPermanentFails
    leafs["sent-permanent-fails"] = peer.SentPermanentFails
    leafs["transport-down"] = peer.TransportDown
    leafs["state"] = peer.State
    return leafs
}

func (peer *Aaa_Diameter_Peers_Peer) GetBundleName() string { return "cisco_ios_xr" }

func (peer *Aaa_Diameter_Peers_Peer) GetYangName() string { return "peer" }

func (peer *Aaa_Diameter_Peers_Peer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peer *Aaa_Diameter_Peers_Peer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peer *Aaa_Diameter_Peers_Peer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peer *Aaa_Diameter_Peers_Peer) SetParent(parent types.Entity) { peer.parent = parent }

func (peer *Aaa_Diameter_Peers_Peer) GetParent() types.Entity { return peer.parent }

func (peer *Aaa_Diameter_Peers_Peer) GetParentYangName() string { return "peers" }

// Aaa_Diameter_Nas
// Diameter NAS data
type Aaa_Diameter_Nas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA NAS id. The type is string.
    AaanasId interface{}

    // List of NAS Entries. The type is slice of Aaa_Diameter_Nas_ListOfNas.
    ListOfNas []Aaa_Diameter_Nas_ListOfNas
}

func (nas *Aaa_Diameter_Nas) GetFilter() yfilter.YFilter { return nas.YFilter }

func (nas *Aaa_Diameter_Nas) SetFilter(yf yfilter.YFilter) { nas.YFilter = yf }

func (nas *Aaa_Diameter_Nas) GetGoName(yname string) string {
    if yname == "aaanas-id" { return "AaanasId" }
    if yname == "list-of-nas" { return "ListOfNas" }
    return ""
}

func (nas *Aaa_Diameter_Nas) GetSegmentPath() string {
    return "nas"
}

func (nas *Aaa_Diameter_Nas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "list-of-nas" {
        for _, c := range nas.ListOfNas {
            if nas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_Nas_ListOfNas{}
        nas.ListOfNas = append(nas.ListOfNas, child)
        return &nas.ListOfNas[len(nas.ListOfNas)-1]
    }
    return nil
}

func (nas *Aaa_Diameter_Nas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nas.ListOfNas {
        children[nas.ListOfNas[i].GetSegmentPath()] = &nas.ListOfNas[i]
    }
    return children
}

func (nas *Aaa_Diameter_Nas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaanas-id"] = nas.AaanasId
    return leafs
}

func (nas *Aaa_Diameter_Nas) GetBundleName() string { return "cisco_ios_xr" }

func (nas *Aaa_Diameter_Nas) GetYangName() string { return "nas" }

func (nas *Aaa_Diameter_Nas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nas *Aaa_Diameter_Nas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nas *Aaa_Diameter_Nas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nas *Aaa_Diameter_Nas) SetParent(parent types.Entity) { nas.parent = parent }

func (nas *Aaa_Diameter_Nas) GetParent() types.Entity { return nas.parent }

func (nas *Aaa_Diameter_Nas) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_Nas_ListOfNas
// List of NAS Entries
type Aaa_Diameter_Nas_ListOfNas struct {
    parent types.Entity
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

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetFilter() yfilter.YFilter { return listOfNas.YFilter }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) SetFilter(yf yfilter.YFilter) { listOfNas.YFilter = yf }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetGoName(yname string) string {
    if yname == "aaa-session-id" { return "AaaSessionId" }
    if yname == "diameter-session-id" { return "DiameterSessionId" }
    if yname == "authentication-status" { return "AuthenticationStatus" }
    if yname == "authorization-status" { return "AuthorizationStatus" }
    if yname == "accounting-status" { return "AccountingStatus" }
    if yname == "accounting-status-stop" { return "AccountingStatusStop" }
    if yname == "disconnect-status" { return "DisconnectStatus" }
    if yname == "accounting-intrim-in-packets" { return "AccountingIntrimInPackets" }
    if yname == "accounting-intrim-out-packets" { return "AccountingIntrimOutPackets" }
    if yname == "method-list" { return "MethodList" }
    if yname == "server-used-list" { return "ServerUsedList" }
    return ""
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetSegmentPath() string {
    return "list-of-nas"
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaa-session-id"] = listOfNas.AaaSessionId
    leafs["diameter-session-id"] = listOfNas.DiameterSessionId
    leafs["authentication-status"] = listOfNas.AuthenticationStatus
    leafs["authorization-status"] = listOfNas.AuthorizationStatus
    leafs["accounting-status"] = listOfNas.AccountingStatus
    leafs["accounting-status-stop"] = listOfNas.AccountingStatusStop
    leafs["disconnect-status"] = listOfNas.DisconnectStatus
    leafs["accounting-intrim-in-packets"] = listOfNas.AccountingIntrimInPackets
    leafs["accounting-intrim-out-packets"] = listOfNas.AccountingIntrimOutPackets
    leafs["method-list"] = listOfNas.MethodList
    leafs["server-used-list"] = listOfNas.ServerUsedList
    return leafs
}

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetBundleName() string { return "cisco_ios_xr" }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetYangName() string { return "list-of-nas" }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) SetParent(parent types.Entity) { listOfNas.parent = parent }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetParent() types.Entity { return listOfNas.parent }

func (listOfNas *Aaa_Diameter_Nas_ListOfNas) GetParentYangName() string { return "nas" }

// Aaa_Diameter_NasSummary
// Diameter NAS summary
type Aaa_Diameter_NasSummary struct {
    parent types.Entity
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

func (nasSummary *Aaa_Diameter_NasSummary) GetFilter() yfilter.YFilter { return nasSummary.YFilter }

func (nasSummary *Aaa_Diameter_NasSummary) SetFilter(yf yfilter.YFilter) { nasSummary.YFilter = yf }

func (nasSummary *Aaa_Diameter_NasSummary) GetGoName(yname string) string {
    if yname == "authen-response-in-packets" { return "AuthenResponseInPackets" }
    if yname == "authen-request-out-packets" { return "AuthenRequestOutPackets" }
    if yname == "authen-request-in-packets" { return "AuthenRequestInPackets" }
    if yname == "authen-response-out-packets" { return "AuthenResponseOutPackets" }
    if yname == "authen-success-packets" { return "AuthenSuccessPackets" }
    if yname == "authen-response-fail-packets" { return "AuthenResponseFailPackets" }
    if yname == "authorization-in-packets" { return "AuthorizationInPackets" }
    if yname == "authorization-out-packets" { return "AuthorizationOutPackets" }
    if yname == "authorization-request-in-packets" { return "AuthorizationRequestInPackets" }
    if yname == "authorization-response-out-packets" { return "AuthorizationResponseOutPackets" }
    if yname == "authorization-response-success-packets" { return "AuthorizationResponseSuccessPackets" }
    if yname == "authorization-response-fail-packets" { return "AuthorizationResponseFailPackets" }
    if yname == "accounting-response-in-packets" { return "AccountingResponseInPackets" }
    if yname == "accounting-request-out-packets" { return "AccountingRequestOutPackets" }
    if yname == "accounting-start-request-packets" { return "AccountingStartRequestPackets" }
    if yname == "accounting-start-response-packets" { return "AccountingStartResponsePackets" }
    if yname == "accounting-start-success-packets" { return "AccountingStartSuccessPackets" }
    if yname == "accounting-start-failed-packets" { return "AccountingStartFailedPackets" }
    if yname == "accounting-stop-response-in-packets" { return "AccountingStopResponseInPackets" }
    if yname == "accounting-stop-request-out-packets" { return "AccountingStopRequestOutPackets" }
    if yname == "accounting-stop-request-in-packets" { return "AccountingStopRequestInPackets" }
    if yname == "accounting-stop-response-out-packets" { return "AccountingStopResponseOutPackets" }
    if yname == "accounting-stop-success-response-packets" { return "AccountingStopSuccessResponsePackets" }
    if yname == "accounting-stop-failed-packets" { return "AccountingStopFailedPackets" }
    if yname == "accounting-intrim-response-in-packets" { return "AccountingIntrimResponseInPackets" }
    if yname == "accounting-interim-request-out-packets" { return "AccountingInterimRequestOutPackets" }
    if yname == "accounting-interim-request-in-packets" { return "AccountingInterimRequestInPackets" }
    if yname == "accounting-interim-response-out-packets" { return "AccountingInterimResponseOutPackets" }
    if yname == "accounting-interim-success-packets" { return "AccountingInterimSuccessPackets" }
    if yname == "accounting-interim-failed-packets" { return "AccountingInterimFailedPackets" }
    if yname == "disconnect-response-in-packets" { return "DisconnectResponseInPackets" }
    if yname == "disconnect-request-out-packets" { return "DisconnectRequestOutPackets" }
    if yname == "disconnect-request-in-packets" { return "DisconnectRequestInPackets" }
    if yname == "disconnect-response-out-packets" { return "DisconnectResponseOutPackets" }
    if yname == "disconnect-success-response-packets" { return "DisconnectSuccessResponsePackets" }
    if yname == "disconnect-failed-response-packets" { return "DisconnectFailedResponsePackets" }
    if yname == "coa-request-in-packets" { return "CoaRequestInPackets" }
    if yname == "coa-response-out-packets" { return "CoaResponseOutPackets" }
    if yname == "coa-request-packets" { return "CoaRequestPackets" }
    if yname == "coa-response-packets" { return "CoaResponsePackets" }
    if yname == "coa-success-packets" { return "CoaSuccessPackets" }
    if yname == "coa-failed-packets" { return "CoaFailedPackets" }
    if yname == "pod-in-packets" { return "PodInPackets" }
    if yname == "pod-out-packets" { return "PodOutPackets" }
    if yname == "pod-request-in-packets" { return "PodRequestInPackets" }
    if yname == "pod-response-out-packets" { return "PodResponseOutPackets" }
    if yname == "pod-success-packets" { return "PodSuccessPackets" }
    if yname == "pod-failed-packets" { return "PodFailedPackets" }
    return ""
}

func (nasSummary *Aaa_Diameter_NasSummary) GetSegmentPath() string {
    return "nas-summary"
}

func (nasSummary *Aaa_Diameter_NasSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nasSummary *Aaa_Diameter_NasSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nasSummary *Aaa_Diameter_NasSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["authen-response-in-packets"] = nasSummary.AuthenResponseInPackets
    leafs["authen-request-out-packets"] = nasSummary.AuthenRequestOutPackets
    leafs["authen-request-in-packets"] = nasSummary.AuthenRequestInPackets
    leafs["authen-response-out-packets"] = nasSummary.AuthenResponseOutPackets
    leafs["authen-success-packets"] = nasSummary.AuthenSuccessPackets
    leafs["authen-response-fail-packets"] = nasSummary.AuthenResponseFailPackets
    leafs["authorization-in-packets"] = nasSummary.AuthorizationInPackets
    leafs["authorization-out-packets"] = nasSummary.AuthorizationOutPackets
    leafs["authorization-request-in-packets"] = nasSummary.AuthorizationRequestInPackets
    leafs["authorization-response-out-packets"] = nasSummary.AuthorizationResponseOutPackets
    leafs["authorization-response-success-packets"] = nasSummary.AuthorizationResponseSuccessPackets
    leafs["authorization-response-fail-packets"] = nasSummary.AuthorizationResponseFailPackets
    leafs["accounting-response-in-packets"] = nasSummary.AccountingResponseInPackets
    leafs["accounting-request-out-packets"] = nasSummary.AccountingRequestOutPackets
    leafs["accounting-start-request-packets"] = nasSummary.AccountingStartRequestPackets
    leafs["accounting-start-response-packets"] = nasSummary.AccountingStartResponsePackets
    leafs["accounting-start-success-packets"] = nasSummary.AccountingStartSuccessPackets
    leafs["accounting-start-failed-packets"] = nasSummary.AccountingStartFailedPackets
    leafs["accounting-stop-response-in-packets"] = nasSummary.AccountingStopResponseInPackets
    leafs["accounting-stop-request-out-packets"] = nasSummary.AccountingStopRequestOutPackets
    leafs["accounting-stop-request-in-packets"] = nasSummary.AccountingStopRequestInPackets
    leafs["accounting-stop-response-out-packets"] = nasSummary.AccountingStopResponseOutPackets
    leafs["accounting-stop-success-response-packets"] = nasSummary.AccountingStopSuccessResponsePackets
    leafs["accounting-stop-failed-packets"] = nasSummary.AccountingStopFailedPackets
    leafs["accounting-intrim-response-in-packets"] = nasSummary.AccountingIntrimResponseInPackets
    leafs["accounting-interim-request-out-packets"] = nasSummary.AccountingInterimRequestOutPackets
    leafs["accounting-interim-request-in-packets"] = nasSummary.AccountingInterimRequestInPackets
    leafs["accounting-interim-response-out-packets"] = nasSummary.AccountingInterimResponseOutPackets
    leafs["accounting-interim-success-packets"] = nasSummary.AccountingInterimSuccessPackets
    leafs["accounting-interim-failed-packets"] = nasSummary.AccountingInterimFailedPackets
    leafs["disconnect-response-in-packets"] = nasSummary.DisconnectResponseInPackets
    leafs["disconnect-request-out-packets"] = nasSummary.DisconnectRequestOutPackets
    leafs["disconnect-request-in-packets"] = nasSummary.DisconnectRequestInPackets
    leafs["disconnect-response-out-packets"] = nasSummary.DisconnectResponseOutPackets
    leafs["disconnect-success-response-packets"] = nasSummary.DisconnectSuccessResponsePackets
    leafs["disconnect-failed-response-packets"] = nasSummary.DisconnectFailedResponsePackets
    leafs["coa-request-in-packets"] = nasSummary.CoaRequestInPackets
    leafs["coa-response-out-packets"] = nasSummary.CoaResponseOutPackets
    leafs["coa-request-packets"] = nasSummary.CoaRequestPackets
    leafs["coa-response-packets"] = nasSummary.CoaResponsePackets
    leafs["coa-success-packets"] = nasSummary.CoaSuccessPackets
    leafs["coa-failed-packets"] = nasSummary.CoaFailedPackets
    leafs["pod-in-packets"] = nasSummary.PodInPackets
    leafs["pod-out-packets"] = nasSummary.PodOutPackets
    leafs["pod-request-in-packets"] = nasSummary.PodRequestInPackets
    leafs["pod-response-out-packets"] = nasSummary.PodResponseOutPackets
    leafs["pod-success-packets"] = nasSummary.PodSuccessPackets
    leafs["pod-failed-packets"] = nasSummary.PodFailedPackets
    return leafs
}

func (nasSummary *Aaa_Diameter_NasSummary) GetBundleName() string { return "cisco_ios_xr" }

func (nasSummary *Aaa_Diameter_NasSummary) GetYangName() string { return "nas-summary" }

func (nasSummary *Aaa_Diameter_NasSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nasSummary *Aaa_Diameter_NasSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nasSummary *Aaa_Diameter_NasSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nasSummary *Aaa_Diameter_NasSummary) SetParent(parent types.Entity) { nasSummary.parent = parent }

func (nasSummary *Aaa_Diameter_NasSummary) GetParent() types.Entity { return nasSummary.parent }

func (nasSummary *Aaa_Diameter_NasSummary) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_GySessionIds
// Diameter Gy Session data list
type Aaa_Diameter_GySessionIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diameter Gy Session data. The type is slice of
    // Aaa_Diameter_GySessionIds_GySessionId.
    GySessionId []Aaa_Diameter_GySessionIds_GySessionId
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetFilter() yfilter.YFilter { return gySessionIds.YFilter }

func (gySessionIds *Aaa_Diameter_GySessionIds) SetFilter(yf yfilter.YFilter) { gySessionIds.YFilter = yf }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetGoName(yname string) string {
    if yname == "gy-session-id" { return "GySessionId" }
    return ""
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetSegmentPath() string {
    return "gy-session-ids"
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gy-session-id" {
        for _, c := range gySessionIds.GySessionId {
            if gySessionIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_GySessionIds_GySessionId{}
        gySessionIds.GySessionId = append(gySessionIds.GySessionId, child)
        return &gySessionIds.GySessionId[len(gySessionIds.GySessionId)-1]
    }
    return nil
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range gySessionIds.GySessionId {
        children[gySessionIds.GySessionId[i].GetSegmentPath()] = &gySessionIds.GySessionId[i]
    }
    return children
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gySessionIds *Aaa_Diameter_GySessionIds) GetBundleName() string { return "cisco_ios_xr" }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetYangName() string { return "gy-session-ids" }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gySessionIds *Aaa_Diameter_GySessionIds) SetParent(parent types.Entity) { gySessionIds.parent = parent }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetParent() types.Entity { return gySessionIds.parent }

func (gySessionIds *Aaa_Diameter_GySessionIds) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_GySessionIds_GySessionId
// Diameter Gy Session data
type Aaa_Diameter_GySessionIds_GySessionId struct {
    parent types.Entity
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

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetFilter() yfilter.YFilter { return gySessionId.YFilter }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) SetFilter(yf yfilter.YFilter) { gySessionId.YFilter = yf }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "aaa-session-id" { return "AaaSessionId" }
    if yname == "parent-aaa-session-id" { return "ParentAaaSessionId" }
    if yname == "diameter-session-id" { return "DiameterSessionId" }
    if yname == "request-number" { return "RequestNumber" }
    if yname == "session-state" { return "SessionState" }
    if yname == "request-type" { return "RequestType" }
    if yname == "retry-count" { return "RetryCount" }
    return ""
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetSegmentPath() string {
    return "gy-session-id" + "[session-id='" + fmt.Sprintf("%v", gySessionId.SessionId) + "']"
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = gySessionId.SessionId
    leafs["aaa-session-id"] = gySessionId.AaaSessionId
    leafs["parent-aaa-session-id"] = gySessionId.ParentAaaSessionId
    leafs["diameter-session-id"] = gySessionId.DiameterSessionId
    leafs["request-number"] = gySessionId.RequestNumber
    leafs["session-state"] = gySessionId.SessionState
    leafs["request-type"] = gySessionId.RequestType
    leafs["retry-count"] = gySessionId.RetryCount
    return leafs
}

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetBundleName() string { return "cisco_ios_xr" }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetYangName() string { return "gy-session-id" }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) SetParent(parent types.Entity) { gySessionId.parent = parent }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetParent() types.Entity { return gySessionId.parent }

func (gySessionId *Aaa_Diameter_GySessionIds_GySessionId) GetParentYangName() string { return "gy-session-ids" }

// Aaa_Diameter_GyStatistics
// Diameter Gy Statistics data
type Aaa_Diameter_GyStatistics struct {
    parent types.Entity
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

func (gyStatistics *Aaa_Diameter_GyStatistics) GetFilter() yfilter.YFilter { return gyStatistics.YFilter }

func (gyStatistics *Aaa_Diameter_GyStatistics) SetFilter(yf yfilter.YFilter) { gyStatistics.YFilter = yf }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetGoName(yname string) string {
    if yname == "ccr-init-messages" { return "CcrInitMessages" }
    if yname == "ccr-init-failed-messages" { return "CcrInitFailedMessages" }
    if yname == "ccr-init-timed-out-messages" { return "CcrInitTimedOutMessages" }
    if yname == "ccr-init-retry-messages" { return "CcrInitRetryMessages" }
    if yname == "ccr-update-messages" { return "CcrUpdateMessages" }
    if yname == "ccr-update-failed-messages" { return "CcrUpdateFailedMessages" }
    if yname == "ccr-update-timed-out-messages" { return "CcrUpdateTimedOutMessages" }
    if yname == "ccr-update-retry-messages" { return "CcrUpdateRetryMessages" }
    if yname == "ccr-final-messages" { return "CcrFinalMessages" }
    if yname == "ccr-final-failed-messages" { return "CcrFinalFailedMessages" }
    if yname == "ccr-final-timed-out-messages" { return "CcrFinalTimedOutMessages" }
    if yname == "ccr-final-retry-messages" { return "CcrFinalRetryMessages" }
    if yname == "cca-init-messages" { return "CcaInitMessages" }
    if yname == "cca-init-error-messages" { return "CcaInitErrorMessages" }
    if yname == "cca-update-messages" { return "CcaUpdateMessages" }
    if yname == "cca-update-error-messages" { return "CcaUpdateErrorMessages" }
    if yname == "cca-final-messages" { return "CcaFinalMessages" }
    if yname == "cca-final-error-messages" { return "CcaFinalErrorMessages" }
    if yname == "rar-received-messages" { return "RarReceivedMessages" }
    if yname == "rar-received-error-messages" { return "RarReceivedErrorMessages" }
    if yname == "raa-sent-messages" { return "RaaSentMessages" }
    if yname == "raa-sent-error-messages" { return "RaaSentErrorMessages" }
    if yname == "asr-received-messages" { return "AsrReceivedMessages" }
    if yname == "asr-received-error-messages" { return "AsrReceivedErrorMessages" }
    if yname == "asa-sent-messages" { return "AsaSentMessages" }
    if yname == "asa-sent-error-messages" { return "AsaSentErrorMessages" }
    if yname == "unknown-request-messages" { return "UnknownRequestMessages" }
    if yname == "restore-sessions" { return "RestoreSessions" }
    if yname == "open-sessions" { return "OpenSessions" }
    if yname == "close-sessions" { return "CloseSessions" }
    if yname == "active-sessions" { return "ActiveSessions" }
    return ""
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetSegmentPath() string {
    return "gy-statistics"
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ccr-init-messages"] = gyStatistics.CcrInitMessages
    leafs["ccr-init-failed-messages"] = gyStatistics.CcrInitFailedMessages
    leafs["ccr-init-timed-out-messages"] = gyStatistics.CcrInitTimedOutMessages
    leafs["ccr-init-retry-messages"] = gyStatistics.CcrInitRetryMessages
    leafs["ccr-update-messages"] = gyStatistics.CcrUpdateMessages
    leafs["ccr-update-failed-messages"] = gyStatistics.CcrUpdateFailedMessages
    leafs["ccr-update-timed-out-messages"] = gyStatistics.CcrUpdateTimedOutMessages
    leafs["ccr-update-retry-messages"] = gyStatistics.CcrUpdateRetryMessages
    leafs["ccr-final-messages"] = gyStatistics.CcrFinalMessages
    leafs["ccr-final-failed-messages"] = gyStatistics.CcrFinalFailedMessages
    leafs["ccr-final-timed-out-messages"] = gyStatistics.CcrFinalTimedOutMessages
    leafs["ccr-final-retry-messages"] = gyStatistics.CcrFinalRetryMessages
    leafs["cca-init-messages"] = gyStatistics.CcaInitMessages
    leafs["cca-init-error-messages"] = gyStatistics.CcaInitErrorMessages
    leafs["cca-update-messages"] = gyStatistics.CcaUpdateMessages
    leafs["cca-update-error-messages"] = gyStatistics.CcaUpdateErrorMessages
    leafs["cca-final-messages"] = gyStatistics.CcaFinalMessages
    leafs["cca-final-error-messages"] = gyStatistics.CcaFinalErrorMessages
    leafs["rar-received-messages"] = gyStatistics.RarReceivedMessages
    leafs["rar-received-error-messages"] = gyStatistics.RarReceivedErrorMessages
    leafs["raa-sent-messages"] = gyStatistics.RaaSentMessages
    leafs["raa-sent-error-messages"] = gyStatistics.RaaSentErrorMessages
    leafs["asr-received-messages"] = gyStatistics.AsrReceivedMessages
    leafs["asr-received-error-messages"] = gyStatistics.AsrReceivedErrorMessages
    leafs["asa-sent-messages"] = gyStatistics.AsaSentMessages
    leafs["asa-sent-error-messages"] = gyStatistics.AsaSentErrorMessages
    leafs["unknown-request-messages"] = gyStatistics.UnknownRequestMessages
    leafs["restore-sessions"] = gyStatistics.RestoreSessions
    leafs["open-sessions"] = gyStatistics.OpenSessions
    leafs["close-sessions"] = gyStatistics.CloseSessions
    leafs["active-sessions"] = gyStatistics.ActiveSessions
    return leafs
}

func (gyStatistics *Aaa_Diameter_GyStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetYangName() string { return "gy-statistics" }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gyStatistics *Aaa_Diameter_GyStatistics) SetParent(parent types.Entity) { gyStatistics.parent = parent }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetParent() types.Entity { return gyStatistics.parent }

func (gyStatistics *Aaa_Diameter_GyStatistics) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_GxSessionIds
// Diameter Gx Session data list
type Aaa_Diameter_GxSessionIds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Diameter Gx Session data. The type is slice of
    // Aaa_Diameter_GxSessionIds_GxSessionId.
    GxSessionId []Aaa_Diameter_GxSessionIds_GxSessionId
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetFilter() yfilter.YFilter { return gxSessionIds.YFilter }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) SetFilter(yf yfilter.YFilter) { gxSessionIds.YFilter = yf }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetGoName(yname string) string {
    if yname == "gx-session-id" { return "GxSessionId" }
    return ""
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetSegmentPath() string {
    return "gx-session-ids"
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gx-session-id" {
        for _, c := range gxSessionIds.GxSessionId {
            if gxSessionIds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_GxSessionIds_GxSessionId{}
        gxSessionIds.GxSessionId = append(gxSessionIds.GxSessionId, child)
        return &gxSessionIds.GxSessionId[len(gxSessionIds.GxSessionId)-1]
    }
    return nil
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range gxSessionIds.GxSessionId {
        children[gxSessionIds.GxSessionId[i].GetSegmentPath()] = &gxSessionIds.GxSessionId[i]
    }
    return children
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetBundleName() string { return "cisco_ios_xr" }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetYangName() string { return "gx-session-ids" }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) SetParent(parent types.Entity) { gxSessionIds.parent = parent }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetParent() types.Entity { return gxSessionIds.parent }

func (gxSessionIds *Aaa_Diameter_GxSessionIds) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_GxSessionIds_GxSessionId
// Diameter Gx Session data
type Aaa_Diameter_GxSessionIds_GxSessionId struct {
    parent types.Entity
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

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetFilter() yfilter.YFilter { return gxSessionId.YFilter }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) SetFilter(yf yfilter.YFilter) { gxSessionId.YFilter = yf }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "aaa-session-id" { return "AaaSessionId" }
    if yname == "diameter-session-id" { return "DiameterSessionId" }
    if yname == "request-number" { return "RequestNumber" }
    if yname == "session-state" { return "SessionState" }
    if yname == "request-type" { return "RequestType" }
    if yname == "retry-count" { return "RetryCount" }
    return ""
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetSegmentPath() string {
    return "gx-session-id" + "[session-id='" + fmt.Sprintf("%v", gxSessionId.SessionId) + "']"
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = gxSessionId.SessionId
    leafs["aaa-session-id"] = gxSessionId.AaaSessionId
    leafs["diameter-session-id"] = gxSessionId.DiameterSessionId
    leafs["request-number"] = gxSessionId.RequestNumber
    leafs["session-state"] = gxSessionId.SessionState
    leafs["request-type"] = gxSessionId.RequestType
    leafs["retry-count"] = gxSessionId.RetryCount
    return leafs
}

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetBundleName() string { return "cisco_ios_xr" }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetYangName() string { return "gx-session-id" }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) SetParent(parent types.Entity) { gxSessionId.parent = parent }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetParent() types.Entity { return gxSessionId.parent }

func (gxSessionId *Aaa_Diameter_GxSessionIds_GxSessionId) GetParentYangName() string { return "gx-session-ids" }

// Aaa_Diameter_NasSession
// Diameter Nas Session data
type Aaa_Diameter_NasSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AAA NAS id. The type is string.
    AaanasId interface{}

    // List of NAS Entries. The type is slice of
    // Aaa_Diameter_NasSession_ListOfNas.
    ListOfNas []Aaa_Diameter_NasSession_ListOfNas
}

func (nasSession *Aaa_Diameter_NasSession) GetFilter() yfilter.YFilter { return nasSession.YFilter }

func (nasSession *Aaa_Diameter_NasSession) SetFilter(yf yfilter.YFilter) { nasSession.YFilter = yf }

func (nasSession *Aaa_Diameter_NasSession) GetGoName(yname string) string {
    if yname == "aaanas-id" { return "AaanasId" }
    if yname == "list-of-nas" { return "ListOfNas" }
    return ""
}

func (nasSession *Aaa_Diameter_NasSession) GetSegmentPath() string {
    return "nas-session"
}

func (nasSession *Aaa_Diameter_NasSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "list-of-nas" {
        for _, c := range nasSession.ListOfNas {
            if nasSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Diameter_NasSession_ListOfNas{}
        nasSession.ListOfNas = append(nasSession.ListOfNas, child)
        return &nasSession.ListOfNas[len(nasSession.ListOfNas)-1]
    }
    return nil
}

func (nasSession *Aaa_Diameter_NasSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nasSession.ListOfNas {
        children[nasSession.ListOfNas[i].GetSegmentPath()] = &nasSession.ListOfNas[i]
    }
    return children
}

func (nasSession *Aaa_Diameter_NasSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaanas-id"] = nasSession.AaanasId
    return leafs
}

func (nasSession *Aaa_Diameter_NasSession) GetBundleName() string { return "cisco_ios_xr" }

func (nasSession *Aaa_Diameter_NasSession) GetYangName() string { return "nas-session" }

func (nasSession *Aaa_Diameter_NasSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nasSession *Aaa_Diameter_NasSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nasSession *Aaa_Diameter_NasSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nasSession *Aaa_Diameter_NasSession) SetParent(parent types.Entity) { nasSession.parent = parent }

func (nasSession *Aaa_Diameter_NasSession) GetParent() types.Entity { return nasSession.parent }

func (nasSession *Aaa_Diameter_NasSession) GetParentYangName() string { return "diameter" }

// Aaa_Diameter_NasSession_ListOfNas
// List of NAS Entries
type Aaa_Diameter_NasSession_ListOfNas struct {
    parent types.Entity
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

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetFilter() yfilter.YFilter { return listOfNas.YFilter }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) SetFilter(yf yfilter.YFilter) { listOfNas.YFilter = yf }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetGoName(yname string) string {
    if yname == "aaa-session-id" { return "AaaSessionId" }
    if yname == "diameter-session-id" { return "DiameterSessionId" }
    if yname == "authentication-status" { return "AuthenticationStatus" }
    if yname == "authorization-status" { return "AuthorizationStatus" }
    if yname == "accounting-status" { return "AccountingStatus" }
    if yname == "accounting-status-stop" { return "AccountingStatusStop" }
    if yname == "disconnect-status" { return "DisconnectStatus" }
    if yname == "accounting-intrim-in-packets" { return "AccountingIntrimInPackets" }
    if yname == "accounting-intrim-out-packets" { return "AccountingIntrimOutPackets" }
    if yname == "method-list" { return "MethodList" }
    if yname == "server-used-list" { return "ServerUsedList" }
    return ""
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetSegmentPath() string {
    return "list-of-nas"
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["aaa-session-id"] = listOfNas.AaaSessionId
    leafs["diameter-session-id"] = listOfNas.DiameterSessionId
    leafs["authentication-status"] = listOfNas.AuthenticationStatus
    leafs["authorization-status"] = listOfNas.AuthorizationStatus
    leafs["accounting-status"] = listOfNas.AccountingStatus
    leafs["accounting-status-stop"] = listOfNas.AccountingStatusStop
    leafs["disconnect-status"] = listOfNas.DisconnectStatus
    leafs["accounting-intrim-in-packets"] = listOfNas.AccountingIntrimInPackets
    leafs["accounting-intrim-out-packets"] = listOfNas.AccountingIntrimOutPackets
    leafs["method-list"] = listOfNas.MethodList
    leafs["server-used-list"] = listOfNas.ServerUsedList
    return leafs
}

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetBundleName() string { return "cisco_ios_xr" }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetYangName() string { return "list-of-nas" }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) SetParent(parent types.Entity) { listOfNas.parent = parent }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetParent() types.Entity { return listOfNas.parent }

func (listOfNas *Aaa_Diameter_NasSession_ListOfNas) GetParentYangName() string { return "nas-session" }

// Aaa_Radius
// RADIUS operational data
type Aaa_Radius struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of RADIUS servers configured.
    Servers Aaa_Radius_Servers

    // RADIUS source interfaces.
    RadiusSourceInterface Aaa_Radius_RadiusSourceInterface

    // RADIUS Client Information.
    Global Aaa_Radius_Global
}

func (radius *Aaa_Radius) GetFilter() yfilter.YFilter { return radius.YFilter }

func (radius *Aaa_Radius) SetFilter(yf yfilter.YFilter) { radius.YFilter = yf }

func (radius *Aaa_Radius) GetGoName(yname string) string {
    if yname == "servers" { return "Servers" }
    if yname == "radius-source-interface" { return "RadiusSourceInterface" }
    if yname == "global" { return "Global" }
    return ""
}

func (radius *Aaa_Radius) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-protocol-radius-oper:radius"
}

func (radius *Aaa_Radius) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "servers" {
        return &radius.Servers
    }
    if childYangName == "radius-source-interface" {
        return &radius.RadiusSourceInterface
    }
    if childYangName == "global" {
        return &radius.Global
    }
    return nil
}

func (radius *Aaa_Radius) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["servers"] = &radius.Servers
    children["radius-source-interface"] = &radius.RadiusSourceInterface
    children["global"] = &radius.Global
    return children
}

func (radius *Aaa_Radius) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radius *Aaa_Radius) GetBundleName() string { return "cisco_ios_xr" }

func (radius *Aaa_Radius) GetYangName() string { return "radius" }

func (radius *Aaa_Radius) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radius *Aaa_Radius) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radius *Aaa_Radius) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radius *Aaa_Radius) SetParent(parent types.Entity) { radius.parent = parent }

func (radius *Aaa_Radius) GetParent() types.Entity { return radius.parent }

func (radius *Aaa_Radius) GetParentYangName() string { return "aaa" }

// Aaa_Radius_Servers
// List of RADIUS servers configured
type Aaa_Radius_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RADIUS Server. The type is slice of Aaa_Radius_Servers_Server.
    Server []Aaa_Radius_Servers_Server
}

func (servers *Aaa_Radius_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *Aaa_Radius_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *Aaa_Radius_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *Aaa_Radius_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *Aaa_Radius_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *Aaa_Radius_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *Aaa_Radius_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *Aaa_Radius_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *Aaa_Radius_Servers) GetYangName() string { return "servers" }

func (servers *Aaa_Radius_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *Aaa_Radius_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *Aaa_Radius_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *Aaa_Radius_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *Aaa_Radius_Servers) GetParent() types.Entity { return servers.parent }

func (servers *Aaa_Radius_Servers) GetParentYangName() string { return "radius" }

// Aaa_Radius_Servers_Server
// RADIUS Server
type Aaa_Radius_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP address of RADIUS server. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Authentication Port number (standard port 1645). The type is interface{}
    // with range: 1..65535.
    AuthPortNumber interface{}

    // Accounting Port number (standard port 1646). The type is interface{} with
    // range: 1..65535.
    AcctPortNumber interface{}

    // IP address of RADIUS server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (server *Aaa_Radius_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_Radius_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_Radius_Servers_Server) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auth-port-number" { return "AuthPortNumber" }
    if yname == "acct-port-number" { return "AcctPortNumber" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "priority" { return "Priority" }
    if yname == "timeout-xr" { return "TimeoutXr" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "dead-time" { return "DeadTime" }
    if yname == "dead-detect-time" { return "DeadDetectTime" }
    if yname == "dead-detect-tries" { return "DeadDetectTries" }
    if yname == "authentication-port" { return "AuthenticationPort" }
    if yname == "accounting-port" { return "AccountingPort" }
    if yname == "state" { return "State" }
    if yname == "current-state-duration" { return "CurrentStateDuration" }
    if yname == "previous-state-duration" { return "PreviousStateDuration" }
    if yname == "packets-in" { return "PacketsIn" }
    if yname == "packets-out" { return "PacketsOut" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "aborts" { return "Aborts" }
    if yname == "replies-expected" { return "RepliesExpected" }
    if yname == "redirected-requests" { return "RedirectedRequests" }
    if yname == "authentication-rtt" { return "AuthenticationRtt" }
    if yname == "access-requests" { return "AccessRequests" }
    if yname == "access-request-retransmits" { return "AccessRequestRetransmits" }
    if yname == "access-accepts" { return "AccessAccepts" }
    if yname == "access-rejects" { return "AccessRejects" }
    if yname == "access-challenges" { return "AccessChallenges" }
    if yname == "bad-access-responses" { return "BadAccessResponses" }
    if yname == "bad-access-authenticators" { return "BadAccessAuthenticators" }
    if yname == "pending-access-requests" { return "PendingAccessRequests" }
    if yname == "access-timeouts" { return "AccessTimeouts" }
    if yname == "unknown-access-types" { return "UnknownAccessTypes" }
    if yname == "dropped-access-responses" { return "DroppedAccessResponses" }
    if yname == "throttled-access-reqs" { return "ThrottledAccessReqs" }
    if yname == "throttled-timed-out-reqs" { return "ThrottledTimedOutReqs" }
    if yname == "throttled-dropped-reqs" { return "ThrottledDroppedReqs" }
    if yname == "max-throttled-access-reqs" { return "MaxThrottledAccessReqs" }
    if yname == "currently-throttled-access-reqs" { return "CurrentlyThrottledAccessReqs" }
    if yname == "authen-response-time" { return "AuthenResponseTime" }
    if yname == "authen-transaction-successess" { return "AuthenTransactionSuccessess" }
    if yname == "authen-transaction-failure" { return "AuthenTransactionFailure" }
    if yname == "authen-unexpected-responses" { return "AuthenUnexpectedResponses" }
    if yname == "authen-server-error-responses" { return "AuthenServerErrorResponses" }
    if yname == "authen-incorrect-responses" { return "AuthenIncorrectResponses" }
    if yname == "author-requests" { return "AuthorRequests" }
    if yname == "author-request-timeouts" { return "AuthorRequestTimeouts" }
    if yname == "author-response-time" { return "AuthorResponseTime" }
    if yname == "author-transaction-successess" { return "AuthorTransactionSuccessess" }
    if yname == "author-transaction-failure" { return "AuthorTransactionFailure" }
    if yname == "author-unexpected-responses" { return "AuthorUnexpectedResponses" }
    if yname == "author-server-error-responses" { return "AuthorServerErrorResponses" }
    if yname == "author-incorrect-responses" { return "AuthorIncorrectResponses" }
    if yname == "accounting-rtt" { return "AccountingRtt" }
    if yname == "accounting-requests" { return "AccountingRequests" }
    if yname == "accounting-retransmits" { return "AccountingRetransmits" }
    if yname == "accounting-responses" { return "AccountingResponses" }
    if yname == "bad-accounting-responses" { return "BadAccountingResponses" }
    if yname == "bad-accounting-authenticators" { return "BadAccountingAuthenticators" }
    if yname == "pending-accounting-requets" { return "PendingAccountingRequets" }
    if yname == "accounting-timeouts" { return "AccountingTimeouts" }
    if yname == "unknown-accounting-types" { return "UnknownAccountingTypes" }
    if yname == "dropped-accounting-responses" { return "DroppedAccountingResponses" }
    if yname == "is-a-private-server" { return "IsAPrivateServer" }
    if yname == "total-test-auth-reqs" { return "TotalTestAuthReqs" }
    if yname == "total-test-auth-timeouts" { return "TotalTestAuthTimeouts" }
    if yname == "total-test-auth-response" { return "TotalTestAuthResponse" }
    if yname == "total-test-auth-pending" { return "TotalTestAuthPending" }
    if yname == "total-test-acct-reqs" { return "TotalTestAcctReqs" }
    if yname == "total-test-acct-timeouts" { return "TotalTestAcctTimeouts" }
    if yname == "total-test-acct-response" { return "TotalTestAcctResponse" }
    if yname == "total-test-acct-pending" { return "TotalTestAcctPending" }
    if yname == "throttled-acct-transactions" { return "ThrottledAcctTransactions" }
    if yname == "throttled-acct-timed-out-stats" { return "ThrottledAcctTimedOutStats" }
    if yname == "throttled-acct-failures-stats" { return "ThrottledAcctFailuresStats" }
    if yname == "max-acct-throttled" { return "MaxAcctThrottled" }
    if yname == "throttleda-acct-transactions" { return "ThrottledaAcctTransactions" }
    if yname == "acct-unexpected-responses" { return "AcctUnexpectedResponses" }
    if yname == "acct-server-error-responses" { return "AcctServerErrorResponses" }
    if yname == "acct-incorrect-responses" { return "AcctIncorrectResponses" }
    if yname == "acct-response-time" { return "AcctResponseTime" }
    if yname == "acct-transaction-successess" { return "AcctTransactionSuccessess" }
    if yname == "acct-transaction-failure" { return "AcctTransactionFailure" }
    if yname == "total-deadtime" { return "TotalDeadtime" }
    if yname == "last-deadtime" { return "LastDeadtime" }
    if yname == "is-quarantined" { return "IsQuarantined" }
    if yname == "group-name" { return "GroupName" }
    if yname == "ip-address-xr" { return "IpAddressXr" }
    if yname == "family" { return "Family" }
    return ""
}

func (server *Aaa_Radius_Servers_Server) GetSegmentPath() string {
    return "server"
}

func (server *Aaa_Radius_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_Radius_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_Radius_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = server.IpAddress
    leafs["auth-port-number"] = server.AuthPortNumber
    leafs["acct-port-number"] = server.AcctPortNumber
    leafs["ipv4-address"] = server.Ipv4Address
    leafs["priority"] = server.Priority
    leafs["timeout-xr"] = server.TimeoutXr
    leafs["retransmit"] = server.Retransmit
    leafs["dead-time"] = server.DeadTime
    leafs["dead-detect-time"] = server.DeadDetectTime
    leafs["dead-detect-tries"] = server.DeadDetectTries
    leafs["authentication-port"] = server.AuthenticationPort
    leafs["accounting-port"] = server.AccountingPort
    leafs["state"] = server.State
    leafs["current-state-duration"] = server.CurrentStateDuration
    leafs["previous-state-duration"] = server.PreviousStateDuration
    leafs["packets-in"] = server.PacketsIn
    leafs["packets-out"] = server.PacketsOut
    leafs["timeouts"] = server.Timeouts
    leafs["aborts"] = server.Aborts
    leafs["replies-expected"] = server.RepliesExpected
    leafs["redirected-requests"] = server.RedirectedRequests
    leafs["authentication-rtt"] = server.AuthenticationRtt
    leafs["access-requests"] = server.AccessRequests
    leafs["access-request-retransmits"] = server.AccessRequestRetransmits
    leafs["access-accepts"] = server.AccessAccepts
    leafs["access-rejects"] = server.AccessRejects
    leafs["access-challenges"] = server.AccessChallenges
    leafs["bad-access-responses"] = server.BadAccessResponses
    leafs["bad-access-authenticators"] = server.BadAccessAuthenticators
    leafs["pending-access-requests"] = server.PendingAccessRequests
    leafs["access-timeouts"] = server.AccessTimeouts
    leafs["unknown-access-types"] = server.UnknownAccessTypes
    leafs["dropped-access-responses"] = server.DroppedAccessResponses
    leafs["throttled-access-reqs"] = server.ThrottledAccessReqs
    leafs["throttled-timed-out-reqs"] = server.ThrottledTimedOutReqs
    leafs["throttled-dropped-reqs"] = server.ThrottledDroppedReqs
    leafs["max-throttled-access-reqs"] = server.MaxThrottledAccessReqs
    leafs["currently-throttled-access-reqs"] = server.CurrentlyThrottledAccessReqs
    leafs["authen-response-time"] = server.AuthenResponseTime
    leafs["authen-transaction-successess"] = server.AuthenTransactionSuccessess
    leafs["authen-transaction-failure"] = server.AuthenTransactionFailure
    leafs["authen-unexpected-responses"] = server.AuthenUnexpectedResponses
    leafs["authen-server-error-responses"] = server.AuthenServerErrorResponses
    leafs["authen-incorrect-responses"] = server.AuthenIncorrectResponses
    leafs["author-requests"] = server.AuthorRequests
    leafs["author-request-timeouts"] = server.AuthorRequestTimeouts
    leafs["author-response-time"] = server.AuthorResponseTime
    leafs["author-transaction-successess"] = server.AuthorTransactionSuccessess
    leafs["author-transaction-failure"] = server.AuthorTransactionFailure
    leafs["author-unexpected-responses"] = server.AuthorUnexpectedResponses
    leafs["author-server-error-responses"] = server.AuthorServerErrorResponses
    leafs["author-incorrect-responses"] = server.AuthorIncorrectResponses
    leafs["accounting-rtt"] = server.AccountingRtt
    leafs["accounting-requests"] = server.AccountingRequests
    leafs["accounting-retransmits"] = server.AccountingRetransmits
    leafs["accounting-responses"] = server.AccountingResponses
    leafs["bad-accounting-responses"] = server.BadAccountingResponses
    leafs["bad-accounting-authenticators"] = server.BadAccountingAuthenticators
    leafs["pending-accounting-requets"] = server.PendingAccountingRequets
    leafs["accounting-timeouts"] = server.AccountingTimeouts
    leafs["unknown-accounting-types"] = server.UnknownAccountingTypes
    leafs["dropped-accounting-responses"] = server.DroppedAccountingResponses
    leafs["is-a-private-server"] = server.IsAPrivateServer
    leafs["total-test-auth-reqs"] = server.TotalTestAuthReqs
    leafs["total-test-auth-timeouts"] = server.TotalTestAuthTimeouts
    leafs["total-test-auth-response"] = server.TotalTestAuthResponse
    leafs["total-test-auth-pending"] = server.TotalTestAuthPending
    leafs["total-test-acct-reqs"] = server.TotalTestAcctReqs
    leafs["total-test-acct-timeouts"] = server.TotalTestAcctTimeouts
    leafs["total-test-acct-response"] = server.TotalTestAcctResponse
    leafs["total-test-acct-pending"] = server.TotalTestAcctPending
    leafs["throttled-acct-transactions"] = server.ThrottledAcctTransactions
    leafs["throttled-acct-timed-out-stats"] = server.ThrottledAcctTimedOutStats
    leafs["throttled-acct-failures-stats"] = server.ThrottledAcctFailuresStats
    leafs["max-acct-throttled"] = server.MaxAcctThrottled
    leafs["throttleda-acct-transactions"] = server.ThrottledaAcctTransactions
    leafs["acct-unexpected-responses"] = server.AcctUnexpectedResponses
    leafs["acct-server-error-responses"] = server.AcctServerErrorResponses
    leafs["acct-incorrect-responses"] = server.AcctIncorrectResponses
    leafs["acct-response-time"] = server.AcctResponseTime
    leafs["acct-transaction-successess"] = server.AcctTransactionSuccessess
    leafs["acct-transaction-failure"] = server.AcctTransactionFailure
    leafs["total-deadtime"] = server.TotalDeadtime
    leafs["last-deadtime"] = server.LastDeadtime
    leafs["is-quarantined"] = server.IsQuarantined
    leafs["group-name"] = server.GroupName
    leafs["ip-address-xr"] = server.IpAddressXr
    leafs["family"] = server.Family
    return leafs
}

func (server *Aaa_Radius_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_Radius_Servers_Server) GetYangName() string { return "server" }

func (server *Aaa_Radius_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_Radius_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_Radius_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_Radius_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_Radius_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_Radius_Servers_Server) GetParentYangName() string { return "servers" }

// Aaa_Radius_RadiusSourceInterface
// RADIUS source interfaces
type Aaa_Radius_RadiusSourceInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of source interfaces. The type is slice of
    // Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface.
    ListOfSourceInterface []Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetFilter() yfilter.YFilter { return radiusSourceInterface.YFilter }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) SetFilter(yf yfilter.YFilter) { radiusSourceInterface.YFilter = yf }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetGoName(yname string) string {
    if yname == "list-of-source-interface" { return "ListOfSourceInterface" }
    return ""
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetSegmentPath() string {
    return "radius-source-interface"
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "list-of-source-interface" {
        for _, c := range radiusSourceInterface.ListOfSourceInterface {
            if radiusSourceInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface{}
        radiusSourceInterface.ListOfSourceInterface = append(radiusSourceInterface.ListOfSourceInterface, child)
        return &radiusSourceInterface.ListOfSourceInterface[len(radiusSourceInterface.ListOfSourceInterface)-1]
    }
    return nil
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range radiusSourceInterface.ListOfSourceInterface {
        children[radiusSourceInterface.ListOfSourceInterface[i].GetSegmentPath()] = &radiusSourceInterface.ListOfSourceInterface[i]
    }
    return children
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetBundleName() string { return "cisco_ios_xr" }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetYangName() string { return "radius-source-interface" }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) SetParent(parent types.Entity) { radiusSourceInterface.parent = parent }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetParent() types.Entity { return radiusSourceInterface.parent }

func (radiusSourceInterface *Aaa_Radius_RadiusSourceInterface) GetParentYangName() string { return "radius" }

// Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface
// List of source interfaces
type Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the source interface. The type is string.
    InterfaceName interface{}

    // IP address buffer. The type is string with length: 0..16.
    Ipaddrv4 interface{}

    // IP address buffer. The type is string with length: 0..46.
    Ipaddrv6 interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    Vrfid interface{}
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetFilter() yfilter.YFilter { return listOfSourceInterface.YFilter }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) SetFilter(yf yfilter.YFilter) { listOfSourceInterface.YFilter = yf }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ipaddrv4" { return "Ipaddrv4" }
    if yname == "ipaddrv6" { return "Ipaddrv6" }
    if yname == "vrfid" { return "Vrfid" }
    return ""
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetSegmentPath() string {
    return "list-of-source-interface"
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = listOfSourceInterface.InterfaceName
    leafs["ipaddrv4"] = listOfSourceInterface.Ipaddrv4
    leafs["ipaddrv6"] = listOfSourceInterface.Ipaddrv6
    leafs["vrfid"] = listOfSourceInterface.Vrfid
    return leafs
}

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetBundleName() string { return "cisco_ios_xr" }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetYangName() string { return "list-of-source-interface" }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) SetParent(parent types.Entity) { listOfSourceInterface.parent = parent }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetParent() types.Entity { return listOfSourceInterface.parent }

func (listOfSourceInterface *Aaa_Radius_RadiusSourceInterface_ListOfSourceInterface) GetParentYangName() string { return "radius-source-interface" }

// Aaa_Radius_Global
// RADIUS Client Information
type Aaa_Radius_Global struct {
    parent types.Entity
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

func (global *Aaa_Radius_Global) GetFilter() yfilter.YFilter { return global.YFilter }

func (global *Aaa_Radius_Global) SetFilter(yf yfilter.YFilter) { global.YFilter = yf }

func (global *Aaa_Radius_Global) GetGoName(yname string) string {
    if yname == "unknown-authentication-response" { return "UnknownAuthenticationResponse" }
    if yname == "authentication-nas-id" { return "AuthenticationNasId" }
    if yname == "unknown-accounting-response" { return "UnknownAccountingResponse" }
    if yname == "accounting-nas-id" { return "AccountingNasId" }
    return ""
}

func (global *Aaa_Radius_Global) GetSegmentPath() string {
    return "global"
}

func (global *Aaa_Radius_Global) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (global *Aaa_Radius_Global) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (global *Aaa_Radius_Global) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unknown-authentication-response"] = global.UnknownAuthenticationResponse
    leafs["authentication-nas-id"] = global.AuthenticationNasId
    leafs["unknown-accounting-response"] = global.UnknownAccountingResponse
    leafs["accounting-nas-id"] = global.AccountingNasId
    return leafs
}

func (global *Aaa_Radius_Global) GetBundleName() string { return "cisco_ios_xr" }

func (global *Aaa_Radius_Global) GetYangName() string { return "global" }

func (global *Aaa_Radius_Global) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (global *Aaa_Radius_Global) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (global *Aaa_Radius_Global) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (global *Aaa_Radius_Global) SetParent(parent types.Entity) { global.parent = parent }

func (global *Aaa_Radius_Global) GetParent() types.Entity { return global.parent }

func (global *Aaa_Radius_Global) GetParentYangName() string { return "radius" }

// Aaa_Tacacs
// TACACS operational data
type Aaa_Tacacs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TACACS Active Request List.
    Requests Aaa_Tacacs_Requests

    // TACACS server Information.
    Servers Aaa_Tacacs_Servers

    // TACACS sg Information.
    ServerGroups Aaa_Tacacs_ServerGroups
}

func (tacacs *Aaa_Tacacs) GetFilter() yfilter.YFilter { return tacacs.YFilter }

func (tacacs *Aaa_Tacacs) SetFilter(yf yfilter.YFilter) { tacacs.YFilter = yf }

func (tacacs *Aaa_Tacacs) GetGoName(yname string) string {
    if yname == "requests" { return "Requests" }
    if yname == "servers" { return "Servers" }
    if yname == "server-groups" { return "ServerGroups" }
    return ""
}

func (tacacs *Aaa_Tacacs) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-tacacs-oper:tacacs"
}

func (tacacs *Aaa_Tacacs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "requests" {
        return &tacacs.Requests
    }
    if childYangName == "servers" {
        return &tacacs.Servers
    }
    if childYangName == "server-groups" {
        return &tacacs.ServerGroups
    }
    return nil
}

func (tacacs *Aaa_Tacacs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["requests"] = &tacacs.Requests
    children["servers"] = &tacacs.Servers
    children["server-groups"] = &tacacs.ServerGroups
    return children
}

func (tacacs *Aaa_Tacacs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tacacs *Aaa_Tacacs) GetBundleName() string { return "cisco_ios_xr" }

func (tacacs *Aaa_Tacacs) GetYangName() string { return "tacacs" }

func (tacacs *Aaa_Tacacs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tacacs *Aaa_Tacacs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tacacs *Aaa_Tacacs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tacacs *Aaa_Tacacs) SetParent(parent types.Entity) { tacacs.parent = parent }

func (tacacs *Aaa_Tacacs) GetParent() types.Entity { return tacacs.parent }

func (tacacs *Aaa_Tacacs) GetParentYangName() string { return "aaa" }

// Aaa_Tacacs_Requests
// TACACS Active Request List
type Aaa_Tacacs_Requests struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // request. The type is slice of Aaa_Tacacs_Requests_Request.
    Request []Aaa_Tacacs_Requests_Request
}

func (requests *Aaa_Tacacs_Requests) GetFilter() yfilter.YFilter { return requests.YFilter }

func (requests *Aaa_Tacacs_Requests) SetFilter(yf yfilter.YFilter) { requests.YFilter = yf }

func (requests *Aaa_Tacacs_Requests) GetGoName(yname string) string {
    if yname == "request" { return "Request" }
    return ""
}

func (requests *Aaa_Tacacs_Requests) GetSegmentPath() string {
    return "requests"
}

func (requests *Aaa_Tacacs_Requests) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "request" {
        for _, c := range requests.Request {
            if requests.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_Requests_Request{}
        requests.Request = append(requests.Request, child)
        return &requests.Request[len(requests.Request)-1]
    }
    return nil
}

func (requests *Aaa_Tacacs_Requests) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range requests.Request {
        children[requests.Request[i].GetSegmentPath()] = &requests.Request[i]
    }
    return children
}

func (requests *Aaa_Tacacs_Requests) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (requests *Aaa_Tacacs_Requests) GetBundleName() string { return "cisco_ios_xr" }

func (requests *Aaa_Tacacs_Requests) GetYangName() string { return "requests" }

func (requests *Aaa_Tacacs_Requests) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requests *Aaa_Tacacs_Requests) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requests *Aaa_Tacacs_Requests) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requests *Aaa_Tacacs_Requests) SetParent(parent types.Entity) { requests.parent = parent }

func (requests *Aaa_Tacacs_Requests) GetParent() types.Entity { return requests.parent }

func (requests *Aaa_Tacacs_Requests) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Requests_Request
// request
type Aaa_Tacacs_Requests_Request struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // tacacs requestbag. The type is slice of
    // Aaa_Tacacs_Requests_Request_TacacsRequestbag.
    TacacsRequestbag []Aaa_Tacacs_Requests_Request_TacacsRequestbag
}

func (request *Aaa_Tacacs_Requests_Request) GetFilter() yfilter.YFilter { return request.YFilter }

func (request *Aaa_Tacacs_Requests_Request) SetFilter(yf yfilter.YFilter) { request.YFilter = yf }

func (request *Aaa_Tacacs_Requests_Request) GetGoName(yname string) string {
    if yname == "tacacs-requestbag" { return "TacacsRequestbag" }
    return ""
}

func (request *Aaa_Tacacs_Requests_Request) GetSegmentPath() string {
    return "request"
}

func (request *Aaa_Tacacs_Requests_Request) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tacacs-requestbag" {
        for _, c := range request.TacacsRequestbag {
            if request.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_Requests_Request_TacacsRequestbag{}
        request.TacacsRequestbag = append(request.TacacsRequestbag, child)
        return &request.TacacsRequestbag[len(request.TacacsRequestbag)-1]
    }
    return nil
}

func (request *Aaa_Tacacs_Requests_Request) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range request.TacacsRequestbag {
        children[request.TacacsRequestbag[i].GetSegmentPath()] = &request.TacacsRequestbag[i]
    }
    return children
}

func (request *Aaa_Tacacs_Requests_Request) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (request *Aaa_Tacacs_Requests_Request) GetBundleName() string { return "cisco_ios_xr" }

func (request *Aaa_Tacacs_Requests_Request) GetYangName() string { return "request" }

func (request *Aaa_Tacacs_Requests_Request) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (request *Aaa_Tacacs_Requests_Request) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (request *Aaa_Tacacs_Requests_Request) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (request *Aaa_Tacacs_Requests_Request) SetParent(parent types.Entity) { request.parent = parent }

func (request *Aaa_Tacacs_Requests_Request) GetParent() types.Entity { return request.parent }

func (request *Aaa_Tacacs_Requests_Request) GetParentYangName() string { return "requests" }

// Aaa_Tacacs_Requests_Request_TacacsRequestbag
// tacacs requestbag
type Aaa_Tacacs_Requests_Request_TacacsRequestbag struct {
    parent types.Entity
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

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetFilter() yfilter.YFilter { return tacacsRequestbag.YFilter }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) SetFilter(yf yfilter.YFilter) { tacacsRequestbag.YFilter = yf }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetGoName(yname string) string {
    if yname == "time-remaining" { return "TimeRemaining" }
    if yname == "bytes-out" { return "BytesOut" }
    if yname == "out-pak-size" { return "OutPakSize" }
    if yname == "bytes-in" { return "BytesIn" }
    if yname == "in-pak-size" { return "InPakSize" }
    if yname == "pak-type" { return "PakType" }
    if yname == "session-id" { return "SessionId" }
    if yname == "sock" { return "Sock" }
    return ""
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetSegmentPath() string {
    return "tacacs-requestbag"
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-remaining"] = tacacsRequestbag.TimeRemaining
    leafs["bytes-out"] = tacacsRequestbag.BytesOut
    leafs["out-pak-size"] = tacacsRequestbag.OutPakSize
    leafs["bytes-in"] = tacacsRequestbag.BytesIn
    leafs["in-pak-size"] = tacacsRequestbag.InPakSize
    leafs["pak-type"] = tacacsRequestbag.PakType
    leafs["session-id"] = tacacsRequestbag.SessionId
    leafs["sock"] = tacacsRequestbag.Sock
    return leafs
}

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetBundleName() string { return "cisco_ios_xr" }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetYangName() string { return "tacacs-requestbag" }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) SetParent(parent types.Entity) { tacacsRequestbag.parent = parent }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetParent() types.Entity { return tacacsRequestbag.parent }

func (tacacsRequestbag *Aaa_Tacacs_Requests_Request_TacacsRequestbag) GetParentYangName() string { return "request" }

// Aaa_Tacacs_Servers
// TACACS server Information
type Aaa_Tacacs_Servers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // server. The type is slice of Aaa_Tacacs_Servers_Server.
    Server []Aaa_Tacacs_Servers_Server
}

func (servers *Aaa_Tacacs_Servers) GetFilter() yfilter.YFilter { return servers.YFilter }

func (servers *Aaa_Tacacs_Servers) SetFilter(yf yfilter.YFilter) { servers.YFilter = yf }

func (servers *Aaa_Tacacs_Servers) GetGoName(yname string) string {
    if yname == "server" { return "Server" }
    return ""
}

func (servers *Aaa_Tacacs_Servers) GetSegmentPath() string {
    return "servers"
}

func (servers *Aaa_Tacacs_Servers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range servers.Server {
            if servers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_Servers_Server{}
        servers.Server = append(servers.Server, child)
        return &servers.Server[len(servers.Server)-1]
    }
    return nil
}

func (servers *Aaa_Tacacs_Servers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range servers.Server {
        children[servers.Server[i].GetSegmentPath()] = &servers.Server[i]
    }
    return children
}

func (servers *Aaa_Tacacs_Servers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (servers *Aaa_Tacacs_Servers) GetBundleName() string { return "cisco_ios_xr" }

func (servers *Aaa_Tacacs_Servers) GetYangName() string { return "servers" }

func (servers *Aaa_Tacacs_Servers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (servers *Aaa_Tacacs_Servers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (servers *Aaa_Tacacs_Servers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (servers *Aaa_Tacacs_Servers) SetParent(parent types.Entity) { servers.parent = parent }

func (servers *Aaa_Tacacs_Servers) GetParent() types.Entity { return servers.parent }

func (servers *Aaa_Tacacs_Servers) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_Servers_Server
// server
type Aaa_Tacacs_Servers_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // internet address of T+ server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (server *Aaa_Tacacs_Servers_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_Tacacs_Servers_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_Tacacs_Servers_Server) GetGoName(yname string) string {
    if yname == "addr" { return "Addr" }
    if yname == "timeout" { return "Timeout" }
    if yname == "port" { return "Port" }
    if yname == "bytes-in" { return "BytesIn" }
    if yname == "bytes-out" { return "BytesOut" }
    if yname == "closes" { return "Closes" }
    if yname == "opens" { return "Opens" }
    if yname == "errors" { return "Errors" }
    if yname == "aborts" { return "Aborts" }
    if yname == "paks-in" { return "PaksIn" }
    if yname == "paks-out" { return "PaksOut" }
    if yname == "replies-expected" { return "RepliesExpected" }
    if yname == "up" { return "Up" }
    if yname == "conn-up" { return "ConnUp" }
    if yname == "single-connect" { return "SingleConnect" }
    if yname == "is-private" { return "IsPrivate" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "addr-buf" { return "AddrBuf" }
    if yname == "family" { return "Family" }
    return ""
}

func (server *Aaa_Tacacs_Servers_Server) GetSegmentPath() string {
    return "server"
}

func (server *Aaa_Tacacs_Servers_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_Tacacs_Servers_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_Tacacs_Servers_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addr"] = server.Addr
    leafs["timeout"] = server.Timeout
    leafs["port"] = server.Port
    leafs["bytes-in"] = server.BytesIn
    leafs["bytes-out"] = server.BytesOut
    leafs["closes"] = server.Closes
    leafs["opens"] = server.Opens
    leafs["errors"] = server.Errors
    leafs["aborts"] = server.Aborts
    leafs["paks-in"] = server.PaksIn
    leafs["paks-out"] = server.PaksOut
    leafs["replies-expected"] = server.RepliesExpected
    leafs["up"] = server.Up
    leafs["conn-up"] = server.ConnUp
    leafs["single-connect"] = server.SingleConnect
    leafs["is-private"] = server.IsPrivate
    leafs["vrf-name"] = server.VrfName
    leafs["addr-buf"] = server.AddrBuf
    leafs["family"] = server.Family
    return leafs
}

func (server *Aaa_Tacacs_Servers_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_Tacacs_Servers_Server) GetYangName() string { return "server" }

func (server *Aaa_Tacacs_Servers_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_Tacacs_Servers_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_Tacacs_Servers_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_Tacacs_Servers_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_Tacacs_Servers_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_Tacacs_Servers_Server) GetParentYangName() string { return "servers" }

// Aaa_Tacacs_ServerGroups
// TACACS sg Information
type Aaa_Tacacs_ServerGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // server group. The type is slice of Aaa_Tacacs_ServerGroups_ServerGroup.
    ServerGroup []Aaa_Tacacs_ServerGroups_ServerGroup
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetFilter() yfilter.YFilter { return serverGroups.YFilter }

func (serverGroups *Aaa_Tacacs_ServerGroups) SetFilter(yf yfilter.YFilter) { serverGroups.YFilter = yf }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetGoName(yname string) string {
    if yname == "server-group" { return "ServerGroup" }
    return ""
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetSegmentPath() string {
    return "server-groups"
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-group" {
        for _, c := range serverGroups.ServerGroup {
            if serverGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_ServerGroups_ServerGroup{}
        serverGroups.ServerGroup = append(serverGroups.ServerGroup, child)
        return &serverGroups.ServerGroup[len(serverGroups.ServerGroup)-1]
    }
    return nil
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serverGroups.ServerGroup {
        children[serverGroups.ServerGroup[i].GetSegmentPath()] = &serverGroups.ServerGroup[i]
    }
    return children
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (serverGroups *Aaa_Tacacs_ServerGroups) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetYangName() string { return "server-groups" }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroups *Aaa_Tacacs_ServerGroups) SetParent(parent types.Entity) { serverGroups.parent = parent }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetParent() types.Entity { return serverGroups.parent }

func (serverGroups *Aaa_Tacacs_ServerGroups) GetParentYangName() string { return "tacacs" }

// Aaa_Tacacs_ServerGroups_ServerGroup
// server group
type Aaa_Tacacs_ServerGroups_ServerGroup struct {
    parent types.Entity
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

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetFilter() yfilter.YFilter { return serverGroup.YFilter }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) SetFilter(yf yfilter.YFilter) { serverGroup.YFilter = yf }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetGoName(yname string) string {
    if yname == "group-name" { return "GroupName" }
    if yname == "sg-map-num" { return "SgMapNum" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "server" { return "Server" }
    return ""
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetSegmentPath() string {
    return "server-group"
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server" {
        for _, c := range serverGroup.Server {
            if serverGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Aaa_Tacacs_ServerGroups_ServerGroup_Server{}
        serverGroup.Server = append(serverGroup.Server, child)
        return &serverGroup.Server[len(serverGroup.Server)-1]
    }
    return nil
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range serverGroup.Server {
        children[serverGroup.Server[i].GetSegmentPath()] = &serverGroup.Server[i]
    }
    return children
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-name"] = serverGroup.GroupName
    leafs["sg-map-num"] = serverGroup.SgMapNum
    leafs["vrf-name"] = serverGroup.VrfName
    return leafs
}

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetBundleName() string { return "cisco_ios_xr" }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetYangName() string { return "server-group" }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) SetParent(parent types.Entity) { serverGroup.parent = parent }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetParent() types.Entity { return serverGroup.parent }

func (serverGroup *Aaa_Tacacs_ServerGroups_ServerGroup) GetParentYangName() string { return "server-groups" }

// Aaa_Tacacs_ServerGroups_ServerGroup_Server
// list of servers in this group
type Aaa_Tacacs_ServerGroups_ServerGroup_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // internet address of T+ server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetGoName(yname string) string {
    if yname == "addr" { return "Addr" }
    if yname == "timeout" { return "Timeout" }
    if yname == "port" { return "Port" }
    if yname == "bytes-in" { return "BytesIn" }
    if yname == "bytes-out" { return "BytesOut" }
    if yname == "closes" { return "Closes" }
    if yname == "opens" { return "Opens" }
    if yname == "errors" { return "Errors" }
    if yname == "aborts" { return "Aborts" }
    if yname == "paks-in" { return "PaksIn" }
    if yname == "paks-out" { return "PaksOut" }
    if yname == "replies-expected" { return "RepliesExpected" }
    if yname == "up" { return "Up" }
    if yname == "conn-up" { return "ConnUp" }
    if yname == "single-connect" { return "SingleConnect" }
    if yname == "is-private" { return "IsPrivate" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "addr-buf" { return "AddrBuf" }
    if yname == "family" { return "Family" }
    return ""
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetSegmentPath() string {
    return "server"
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["addr"] = server.Addr
    leafs["timeout"] = server.Timeout
    leafs["port"] = server.Port
    leafs["bytes-in"] = server.BytesIn
    leafs["bytes-out"] = server.BytesOut
    leafs["closes"] = server.Closes
    leafs["opens"] = server.Opens
    leafs["errors"] = server.Errors
    leafs["aborts"] = server.Aborts
    leafs["paks-in"] = server.PaksIn
    leafs["paks-out"] = server.PaksOut
    leafs["replies-expected"] = server.RepliesExpected
    leafs["up"] = server.Up
    leafs["conn-up"] = server.ConnUp
    leafs["single-connect"] = server.SingleConnect
    leafs["is-private"] = server.IsPrivate
    leafs["vrf-name"] = server.VrfName
    leafs["addr-buf"] = server.AddrBuf
    leafs["family"] = server.Family
    return leafs
}

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetYangName() string { return "server" }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetParent() types.Entity { return server.parent }

func (server *Aaa_Tacacs_ServerGroups_ServerGroup_Server) GetParentYangName() string { return "server-group" }

