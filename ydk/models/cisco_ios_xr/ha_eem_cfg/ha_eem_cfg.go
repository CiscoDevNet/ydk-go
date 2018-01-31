// This module contains a collection of YANG definitions
// for Cisco IOS-XR ha-eem package configuration.
// 
// This module contains definitions
// for the following management objects:
//   event-manager: Event manager configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ha_eem_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ha_eem_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ha-eem-cfg event-manager}", reflect.TypeOf(EventManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ha-eem-cfg:event-manager", reflect.TypeOf(EventManager{}))
}

// EventManagerPolicySec represents Event manager policy sec
type EventManagerPolicySec string

const (
    // Cisco Signature
    EventManagerPolicySec_rsa_2048 EventManagerPolicySec = "rsa-2048"

    // Trust Signature
    EventManagerPolicySec_trust EventManagerPolicySec = "trust"
)

// EventManagerPolicyMode represents Event manager policy mode
type EventManagerPolicyMode string

const (
    // Cisco Signature
    EventManagerPolicyMode_cisco EventManagerPolicyMode = "cisco"

    // Trust Signature
    EventManagerPolicyMode_trust EventManagerPolicyMode = "trust"
)

// EventManagerChecksum represents Event manager checksum
type EventManagerChecksum string

const (
    // Use SHA-1 checksum
    EventManagerChecksum_sha_1 EventManagerChecksum = "sha-1"

    // Use MD5 checksum
    EventManagerChecksum_md5 EventManagerChecksum = "md5"
)

// EventManagerPolicy represents Event manager policy
type EventManagerPolicy string

const (
    // Event manager system policy
    EventManagerPolicy_system EventManagerPolicy = "system"

    // Event manager user policy
    EventManagerPolicy_user EventManagerPolicy = "user"
)

// EventManager
// Event manager configuration
type EventManager struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set refresh time (in seconds) for policy username's AAA taskmap. The type
    // is interface{} with range: 10..4294967295. Units are second. The default
    // value is 1800.
    RefreshTime interface{}

    // Enable suspend policy scheduling. The type is bool.
    ScheduleSuspend interface{}

    // Set event manager user policy directory. The type is string.
    DirectoryUserPolicy interface{}

    // Path of the user policy library directory. The type is string.
    DirectoryUserLibrary interface{}

    // Register an event manager policy.
    Policies EventManager_Policies

    // scheduler classs type.
    SchedulerScript EventManager_SchedulerScript

    // Set an event manager global variable for event manager policies.
    Environments EventManager_Environments
}

func (eventManager *EventManager) GetFilter() yfilter.YFilter { return eventManager.YFilter }

func (eventManager *EventManager) SetFilter(yf yfilter.YFilter) { eventManager.YFilter = yf }

func (eventManager *EventManager) GetGoName(yname string) string {
    if yname == "refresh-time" { return "RefreshTime" }
    if yname == "schedule-suspend" { return "ScheduleSuspend" }
    if yname == "directory-user-policy" { return "DirectoryUserPolicy" }
    if yname == "directory-user-library" { return "DirectoryUserLibrary" }
    if yname == "policies" { return "Policies" }
    if yname == "scheduler-script" { return "SchedulerScript" }
    if yname == "environments" { return "Environments" }
    return ""
}

func (eventManager *EventManager) GetSegmentPath() string {
    return "Cisco-IOS-XR-ha-eem-cfg:event-manager"
}

func (eventManager *EventManager) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policies" {
        return &eventManager.Policies
    }
    if childYangName == "scheduler-script" {
        return &eventManager.SchedulerScript
    }
    if childYangName == "environments" {
        return &eventManager.Environments
    }
    return nil
}

func (eventManager *EventManager) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policies"] = &eventManager.Policies
    children["scheduler-script"] = &eventManager.SchedulerScript
    children["environments"] = &eventManager.Environments
    return children
}

func (eventManager *EventManager) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["refresh-time"] = eventManager.RefreshTime
    leafs["schedule-suspend"] = eventManager.ScheduleSuspend
    leafs["directory-user-policy"] = eventManager.DirectoryUserPolicy
    leafs["directory-user-library"] = eventManager.DirectoryUserLibrary
    return leafs
}

func (eventManager *EventManager) GetBundleName() string { return "cisco_ios_xr" }

func (eventManager *EventManager) GetYangName() string { return "event-manager" }

func (eventManager *EventManager) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventManager *EventManager) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventManager *EventManager) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventManager *EventManager) SetParent(parent types.Entity) { eventManager.parent = parent }

func (eventManager *EventManager) GetParent() types.Entity { return eventManager.parent }

func (eventManager *EventManager) GetParentYangName() string { return "Cisco-IOS-XR-ha-eem-cfg" }

// EventManager_Policies
// Register an event manager policy
type EventManager_Policies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the policy file. The type is slice of EventManager_Policies_Policy.
    Policy []EventManager_Policies_Policy
}

func (policies *EventManager_Policies) GetFilter() yfilter.YFilter { return policies.YFilter }

func (policies *EventManager_Policies) SetFilter(yf yfilter.YFilter) { policies.YFilter = yf }

func (policies *EventManager_Policies) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (policies *EventManager_Policies) GetSegmentPath() string {
    return "policies"
}

func (policies *EventManager_Policies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy" {
        for _, c := range policies.Policy {
            if policies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EventManager_Policies_Policy{}
        policies.Policy = append(policies.Policy, child)
        return &policies.Policy[len(policies.Policy)-1]
    }
    return nil
}

func (policies *EventManager_Policies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policies.Policy {
        children[policies.Policy[i].GetSegmentPath()] = &policies.Policy[i]
    }
    return children
}

func (policies *EventManager_Policies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (policies *EventManager_Policies) GetBundleName() string { return "cisco_ios_xr" }

func (policies *EventManager_Policies) GetYangName() string { return "policies" }

func (policies *EventManager_Policies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policies *EventManager_Policies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policies *EventManager_Policies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policies *EventManager_Policies) SetParent(parent types.Entity) { policies.parent = parent }

func (policies *EventManager_Policies) GetParent() types.Entity { return policies.parent }

func (policies *EventManager_Policies) GetParentYangName() string { return "event-manager" }

// EventManager_Policies_Policy
// Name of the policy file
type EventManager_Policies_Policy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the policy file. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PolicyName interface{}

    // A configured username. The type is string. This attribute is mandatory.
    Username interface{}

    // Time of validity (in seconds) for cached AAA taskmap of username (default
    // is 3600). The type is interface{} with range: 0..4294967295. This attribute
    // is mandatory. Units are second.
    PersistTime interface{}

    // Event manager type of this policy. The type is EventManagerPolicy.
    PolicyType interface{}

    // Specify Embedded Event Manager policy checksum. The type is
    // EventManagerChecksum.
    ChecksumType interface{}

    // CheckSum Value. The type is string.
    CheckSumValue interface{}

    // Specify Embedded Event Manager policy security mode. The type is
    // EventManagerPolicyMode.
    PolicySecurityMode interface{}

    // Event Manager policy security Level. The type is EventManagerPolicySec.
    PolicySecurityLevel interface{}
}

func (policy *EventManager_Policies_Policy) GetFilter() yfilter.YFilter { return policy.YFilter }

func (policy *EventManager_Policies_Policy) SetFilter(yf yfilter.YFilter) { policy.YFilter = yf }

func (policy *EventManager_Policies_Policy) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "username" { return "Username" }
    if yname == "persist-time" { return "PersistTime" }
    if yname == "policy-type" { return "PolicyType" }
    if yname == "checksum-type" { return "ChecksumType" }
    if yname == "check-sum-value" { return "CheckSumValue" }
    if yname == "policy-security-mode" { return "PolicySecurityMode" }
    if yname == "policy-security-level" { return "PolicySecurityLevel" }
    return ""
}

func (policy *EventManager_Policies_Policy) GetSegmentPath() string {
    return "policy" + "[policy-name='" + fmt.Sprintf("%v", policy.PolicyName) + "']"
}

func (policy *EventManager_Policies_Policy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policy *EventManager_Policies_Policy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policy *EventManager_Policies_Policy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = policy.PolicyName
    leafs["username"] = policy.Username
    leafs["persist-time"] = policy.PersistTime
    leafs["policy-type"] = policy.PolicyType
    leafs["checksum-type"] = policy.ChecksumType
    leafs["check-sum-value"] = policy.CheckSumValue
    leafs["policy-security-mode"] = policy.PolicySecurityMode
    leafs["policy-security-level"] = policy.PolicySecurityLevel
    return leafs
}

func (policy *EventManager_Policies_Policy) GetBundleName() string { return "cisco_ios_xr" }

func (policy *EventManager_Policies_Policy) GetYangName() string { return "policy" }

func (policy *EventManager_Policies_Policy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policy *EventManager_Policies_Policy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policy *EventManager_Policies_Policy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policy *EventManager_Policies_Policy) SetParent(parent types.Entity) { policy.parent = parent }

func (policy *EventManager_Policies_Policy) GetParent() types.Entity { return policy.parent }

func (policy *EventManager_Policies_Policy) GetParentYangName() string { return "policies" }

// EventManager_SchedulerScript
// scheduler classs type
type EventManager_SchedulerScript struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // scheduler thread classs .
    ThreadClasses EventManager_SchedulerScript_ThreadClasses
}

func (schedulerScript *EventManager_SchedulerScript) GetFilter() yfilter.YFilter { return schedulerScript.YFilter }

func (schedulerScript *EventManager_SchedulerScript) SetFilter(yf yfilter.YFilter) { schedulerScript.YFilter = yf }

func (schedulerScript *EventManager_SchedulerScript) GetGoName(yname string) string {
    if yname == "thread-classes" { return "ThreadClasses" }
    return ""
}

func (schedulerScript *EventManager_SchedulerScript) GetSegmentPath() string {
    return "scheduler-script"
}

func (schedulerScript *EventManager_SchedulerScript) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thread-classes" {
        return &schedulerScript.ThreadClasses
    }
    return nil
}

func (schedulerScript *EventManager_SchedulerScript) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["thread-classes"] = &schedulerScript.ThreadClasses
    return children
}

func (schedulerScript *EventManager_SchedulerScript) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (schedulerScript *EventManager_SchedulerScript) GetBundleName() string { return "cisco_ios_xr" }

func (schedulerScript *EventManager_SchedulerScript) GetYangName() string { return "scheduler-script" }

func (schedulerScript *EventManager_SchedulerScript) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (schedulerScript *EventManager_SchedulerScript) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (schedulerScript *EventManager_SchedulerScript) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (schedulerScript *EventManager_SchedulerScript) SetParent(parent types.Entity) { schedulerScript.parent = parent }

func (schedulerScript *EventManager_SchedulerScript) GetParent() types.Entity { return schedulerScript.parent }

func (schedulerScript *EventManager_SchedulerScript) GetParentYangName() string { return "event-manager" }

// EventManager_SchedulerScript_ThreadClasses
// scheduler thread classs 
type EventManager_SchedulerScript_ThreadClasses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // scheduler classs type argument. The type is slice of
    // EventManager_SchedulerScript_ThreadClasses_ThreadClass.
    ThreadClass []EventManager_SchedulerScript_ThreadClasses_ThreadClass
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetFilter() yfilter.YFilter { return threadClasses.YFilter }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) SetFilter(yf yfilter.YFilter) { threadClasses.YFilter = yf }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetGoName(yname string) string {
    if yname == "thread-class" { return "ThreadClass" }
    return ""
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetSegmentPath() string {
    return "thread-classes"
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "thread-class" {
        for _, c := range threadClasses.ThreadClass {
            if threadClasses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EventManager_SchedulerScript_ThreadClasses_ThreadClass{}
        threadClasses.ThreadClass = append(threadClasses.ThreadClass, child)
        return &threadClasses.ThreadClass[len(threadClasses.ThreadClass)-1]
    }
    return nil
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range threadClasses.ThreadClass {
        children[threadClasses.ThreadClass[i].GetSegmentPath()] = &threadClasses.ThreadClass[i]
    }
    return children
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetBundleName() string { return "cisco_ios_xr" }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetYangName() string { return "thread-classes" }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) SetParent(parent types.Entity) { threadClasses.parent = parent }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetParent() types.Entity { return threadClasses.parent }

func (threadClasses *EventManager_SchedulerScript_ThreadClasses) GetParentYangName() string { return "scheduler-script" }

// EventManager_SchedulerScript_ThreadClasses_ThreadClass
// scheduler classs type argument
type EventManager_SchedulerScript_ThreadClasses_ThreadClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the global variable. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ThreadClassName interface{}

    // number of scheduler threads. The type is interface{} with range: 1..5. This
    // attribute is mandatory.
    NumThreads interface{}
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetFilter() yfilter.YFilter { return threadClass.YFilter }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) SetFilter(yf yfilter.YFilter) { threadClass.YFilter = yf }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetGoName(yname string) string {
    if yname == "thread-class-name" { return "ThreadClassName" }
    if yname == "num-threads" { return "NumThreads" }
    return ""
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetSegmentPath() string {
    return "thread-class" + "[thread-class-name='" + fmt.Sprintf("%v", threadClass.ThreadClassName) + "']"
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["thread-class-name"] = threadClass.ThreadClassName
    leafs["num-threads"] = threadClass.NumThreads
    return leafs
}

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetBundleName() string { return "cisco_ios_xr" }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetYangName() string { return "thread-class" }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) SetParent(parent types.Entity) { threadClass.parent = parent }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetParent() types.Entity { return threadClass.parent }

func (threadClass *EventManager_SchedulerScript_ThreadClasses_ThreadClass) GetParentYangName() string { return "thread-classes" }

// EventManager_Environments
// Set an event manager global variable for event
// manager policies
type EventManager_Environments struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the global variable. The type is slice of
    // EventManager_Environments_Environment.
    Environment []EventManager_Environments_Environment
}

func (environments *EventManager_Environments) GetFilter() yfilter.YFilter { return environments.YFilter }

func (environments *EventManager_Environments) SetFilter(yf yfilter.YFilter) { environments.YFilter = yf }

func (environments *EventManager_Environments) GetGoName(yname string) string {
    if yname == "environment" { return "Environment" }
    return ""
}

func (environments *EventManager_Environments) GetSegmentPath() string {
    return "environments"
}

func (environments *EventManager_Environments) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "environment" {
        for _, c := range environments.Environment {
            if environments.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := EventManager_Environments_Environment{}
        environments.Environment = append(environments.Environment, child)
        return &environments.Environment[len(environments.Environment)-1]
    }
    return nil
}

func (environments *EventManager_Environments) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range environments.Environment {
        children[environments.Environment[i].GetSegmentPath()] = &environments.Environment[i]
    }
    return children
}

func (environments *EventManager_Environments) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (environments *EventManager_Environments) GetBundleName() string { return "cisco_ios_xr" }

func (environments *EventManager_Environments) GetYangName() string { return "environments" }

func (environments *EventManager_Environments) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (environments *EventManager_Environments) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (environments *EventManager_Environments) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (environments *EventManager_Environments) SetParent(parent types.Entity) { environments.parent = parent }

func (environments *EventManager_Environments) GetParent() types.Entity { return environments.parent }

func (environments *EventManager_Environments) GetParentYangName() string { return "event-manager" }

// EventManager_Environments_Environment
// Name of the global variable
type EventManager_Environments_Environment struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the global variable. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    EnvironmentName interface{}

    // Value of the global variable. The type is string. This attribute is
    // mandatory.
    EnvironmentValue interface{}
}

func (environment *EventManager_Environments_Environment) GetFilter() yfilter.YFilter { return environment.YFilter }

func (environment *EventManager_Environments_Environment) SetFilter(yf yfilter.YFilter) { environment.YFilter = yf }

func (environment *EventManager_Environments_Environment) GetGoName(yname string) string {
    if yname == "environment-name" { return "EnvironmentName" }
    if yname == "environment-value" { return "EnvironmentValue" }
    return ""
}

func (environment *EventManager_Environments_Environment) GetSegmentPath() string {
    return "environment" + "[environment-name='" + fmt.Sprintf("%v", environment.EnvironmentName) + "']"
}

func (environment *EventManager_Environments_Environment) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (environment *EventManager_Environments_Environment) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (environment *EventManager_Environments_Environment) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["environment-name"] = environment.EnvironmentName
    leafs["environment-value"] = environment.EnvironmentValue
    return leafs
}

func (environment *EventManager_Environments_Environment) GetBundleName() string { return "cisco_ios_xr" }

func (environment *EventManager_Environments_Environment) GetYangName() string { return "environment" }

func (environment *EventManager_Environments_Environment) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (environment *EventManager_Environments_Environment) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (environment *EventManager_Environments_Environment) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (environment *EventManager_Environments_Environment) SetParent(parent types.Entity) { environment.parent = parent }

func (environment *EventManager_Environments_Environment) GetParent() types.Entity { return environment.parent }

func (environment *EventManager_Environments_Environment) GetParentYangName() string { return "environments" }

