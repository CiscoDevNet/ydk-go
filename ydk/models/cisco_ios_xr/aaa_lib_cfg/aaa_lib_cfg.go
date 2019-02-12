// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-lib package configuration.
// 
// This module contains definitions
// for the following management objects:
//   aaa: Authentication, Authorization and Accounting
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_lib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_lib_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-aaa-lib-cfg aaa}", reflect.TypeOf(Aaa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-aaa-lib-cfg:aaa", reflect.TypeOf(Aaa{}))
}

// Aaa
// Authentication, Authorization and Accounting
type Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This class is used for setting the default taskgroup to be used for remote
    // server authentication. The type is string.
    DefaultTaskgroup interface{}

    // Enable LI RADIUS Feature. The type is interface{}.
    Intercept interface{}

    // AAA accounting.
    Accountings Aaa_Accountings

    // AAA authorization.
    Authorizations Aaa_Authorizations

    // Configuration related to 'update' accounting.
    AccountingUpdate Aaa_AccountingUpdate

    // AAA banner.
    Banner Aaa_Banner

    // AAA authentication.
    Authentications Aaa_Authentications

    // Configure password-policy.
    PasswordPolicies Aaa_PasswordPolicies

    // AAA group definitions.
    ServerGroups Aaa_ServerGroups

    // Configure local usernames.
    Usernames Aaa_Usernames

    // Specify a taskgroup to inherit from.
    Taskgroups Aaa_Taskgroups

    // Specify a Usergroup to inherit from.
    Usergroups Aaa_Usergroups

    // Modify TACACS+ query parameters.
    Tacacs Aaa_Tacacs

    // AAA subscriber.
    AaaSubscriber Aaa_AaaSubscriber

    // AAA Mobile.
    AaaMobile Aaa_AaaMobile

    // AAA Dot1x.
    AaaDot1x Aaa_AaaDot1x

    // AAA RADIUS attribute configurations.
    RadiusAttribute Aaa_RadiusAttribute

    // Remote Access Dial-In User Service.
    Radius Aaa_Radius

    // Diameter base protocol.
    Diameter Aaa_Diameter
}

func (aaa *Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "Cisco-IOS-XR-aaa-lib-cfg"
    aaa.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-lib-cfg:aaa"
    aaa.EntityData.AbsolutePath = aaa.EntityData.SegmentPath
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = types.NewOrderedMap()
    aaa.EntityData.Children.Append("accountings", types.YChild{"Accountings", &aaa.Accountings})
    aaa.EntityData.Children.Append("authorizations", types.YChild{"Authorizations", &aaa.Authorizations})
    aaa.EntityData.Children.Append("accounting-update", types.YChild{"AccountingUpdate", &aaa.AccountingUpdate})
    aaa.EntityData.Children.Append("banner", types.YChild{"Banner", &aaa.Banner})
    aaa.EntityData.Children.Append("authentications", types.YChild{"Authentications", &aaa.Authentications})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-locald-cfg:password-policies", types.YChild{"PasswordPolicies", &aaa.PasswordPolicies})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-locald-cfg:server-groups", types.YChild{"ServerGroups", &aaa.ServerGroups})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-locald-cfg:usernames", types.YChild{"Usernames", &aaa.Usernames})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-locald-cfg:taskgroups", types.YChild{"Taskgroups", &aaa.Taskgroups})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-locald-cfg:usergroups", types.YChild{"Usergroups", &aaa.Usergroups})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-tacacs-cfg:tacacs", types.YChild{"Tacacs", &aaa.Tacacs})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber", types.YChild{"AaaSubscriber", &aaa.AaaSubscriber})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile", types.YChild{"AaaMobile", &aaa.AaaMobile})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x", types.YChild{"AaaDot1x", &aaa.AaaDot1x})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute", types.YChild{"RadiusAttribute", &aaa.RadiusAttribute})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-protocol-radius-cfg:radius", types.YChild{"Radius", &aaa.Radius})
    aaa.EntityData.Children.Append("Cisco-IOS-XR-aaa-diameter-cfg:diameter", types.YChild{"Diameter", &aaa.Diameter})
    aaa.EntityData.Leafs = types.NewOrderedMap()
    aaa.EntityData.Leafs.Append("default-taskgroup", types.YLeaf{"DefaultTaskgroup", aaa.DefaultTaskgroup})
    aaa.EntityData.Leafs.Append("intercept", types.YLeaf{"Intercept", aaa.Intercept})

    aaa.EntityData.YListKeys = []string {}

    return &(aaa.EntityData)
}

// Aaa_Accountings
// AAA accounting
type Aaa_Accountings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_Accountings_Accounting.
    Accounting []*Aaa_Accountings_Accounting
}

func (accountings *Aaa_Accountings) GetEntityData() *types.CommonEntityData {
    accountings.EntityData.YFilter = accountings.YFilter
    accountings.EntityData.YangName = "accountings"
    accountings.EntityData.BundleName = "cisco_ios_xr"
    accountings.EntityData.ParentYangName = "aaa"
    accountings.EntityData.SegmentPath = "accountings"
    accountings.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + accountings.EntityData.SegmentPath
    accountings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountings.EntityData.Children = types.NewOrderedMap()
    accountings.EntityData.Children.Append("accounting", types.YChild{"Accounting", nil})
    for i := range accountings.Accounting {
        accountings.EntityData.Children.Append(types.GetSegmentPath(accountings.Accounting[i]), types.YChild{"Accounting", accountings.Accounting[i]})
    }
    accountings.EntityData.Leafs = types.NewOrderedMap()

    accountings.EntityData.YListKeys = []string {}

    return &(accountings.EntityData)
}

// Aaa_Accountings_Accounting
// Configurations related to accounting
type Aaa_Accountings_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. exec:Account exec sessions, commands: Account CLI
    // commands. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // rpfailover. The type is AaaAccountingRpFailover.
    RpFailover interface{}

    // Broadcast. The type is AaaAccountingBroadcast.
    Broadcast interface{}

    // Stop only/Start Stop. The type is AaaAccounting.
    TypeXr interface{}

    // Method Type. The type is AaaMethodAccounting.
    Method1 interface{}

    // Method Type. The type is AaaMethodAccounting.
    Method2 interface{}

    // Method Type. The type is AaaMethodAccounting.
    Method3 interface{}

    // Method Type. The type is AaaMethodAccounting.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (accounting *Aaa_Accountings_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "accountings"
    accounting.EntityData.SegmentPath = "accounting" + types.AddKeyToken(accounting.Type, "type") + types.AddKeyToken(accounting.Listname, "listname")
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/accountings/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("type", types.YLeaf{"Type", accounting.Type})
    accounting.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", accounting.Listname})
    accounting.EntityData.Leafs.Append("rp-failover", types.YLeaf{"RpFailover", accounting.RpFailover})
    accounting.EntityData.Leafs.Append("broadcast", types.YLeaf{"Broadcast", accounting.Broadcast})
    accounting.EntityData.Leafs.Append("type-xr", types.YLeaf{"TypeXr", accounting.TypeXr})
    accounting.EntityData.Leafs.Append("method1", types.YLeaf{"Method1", accounting.Method1})
    accounting.EntityData.Leafs.Append("method2", types.YLeaf{"Method2", accounting.Method2})
    accounting.EntityData.Leafs.Append("method3", types.YLeaf{"Method3", accounting.Method3})
    accounting.EntityData.Leafs.Append("method4", types.YLeaf{"Method4", accounting.Method4})
    accounting.EntityData.Leafs.Append("server-group-name1", types.YLeaf{"ServerGroupName1", accounting.ServerGroupName1})
    accounting.EntityData.Leafs.Append("server-group-name2", types.YLeaf{"ServerGroupName2", accounting.ServerGroupName2})
    accounting.EntityData.Leafs.Append("server-group-name3", types.YLeaf{"ServerGroupName3", accounting.ServerGroupName3})
    accounting.EntityData.Leafs.Append("server-group-name4", types.YLeaf{"ServerGroupName4", accounting.ServerGroupName4})

    accounting.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(accounting.EntityData)
}

// Aaa_Authorizations
// AAA authorization
type Aaa_Authorizations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_Authorizations_Authorization.
    Authorization []*Aaa_Authorizations_Authorization
}

func (authorizations *Aaa_Authorizations) GetEntityData() *types.CommonEntityData {
    authorizations.EntityData.YFilter = authorizations.YFilter
    authorizations.EntityData.YangName = "authorizations"
    authorizations.EntityData.BundleName = "cisco_ios_xr"
    authorizations.EntityData.ParentYangName = "aaa"
    authorizations.EntityData.SegmentPath = "authorizations"
    authorizations.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + authorizations.EntityData.SegmentPath
    authorizations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizations.EntityData.Children = types.NewOrderedMap()
    authorizations.EntityData.Children.Append("authorization", types.YChild{"Authorization", nil})
    for i := range authorizations.Authorization {
        authorizations.EntityData.Children.Append(types.GetSegmentPath(authorizations.Authorization[i]), types.YChild{"Authorization", authorizations.Authorization[i]})
    }
    authorizations.EntityData.Leafs = types.NewOrderedMap()

    authorizations.EntityData.YListKeys = []string {}

    return &(authorizations.EntityData)
}

// Aaa_Authorizations_Authorization
// Configurations related to authorization
type Aaa_Authorizations_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. network: Authorize IKE requests, commands:
    // Authorize CLI commands. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. List name for AAA authorization. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Type. The type is AaaMethod.
    Method1 interface{}

    // Method Type. The type is AaaMethod.
    Method2 interface{}

    // Method Type. The type is AaaMethod.
    Method3 interface{}

    // Method Type. The type is AaaMethod.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (authorization *Aaa_Authorizations_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "authorizations"
    authorization.EntityData.SegmentPath = "authorization" + types.AddKeyToken(authorization.Type, "type") + types.AddKeyToken(authorization.Listname, "listname")
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/authorizations/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("type", types.YLeaf{"Type", authorization.Type})
    authorization.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", authorization.Listname})
    authorization.EntityData.Leafs.Append("method1", types.YLeaf{"Method1", authorization.Method1})
    authorization.EntityData.Leafs.Append("method2", types.YLeaf{"Method2", authorization.Method2})
    authorization.EntityData.Leafs.Append("method3", types.YLeaf{"Method3", authorization.Method3})
    authorization.EntityData.Leafs.Append("method4", types.YLeaf{"Method4", authorization.Method4})
    authorization.EntityData.Leafs.Append("server-group-name1", types.YLeaf{"ServerGroupName1", authorization.ServerGroupName1})
    authorization.EntityData.Leafs.Append("server-group-name2", types.YLeaf{"ServerGroupName2", authorization.ServerGroupName2})
    authorization.EntityData.Leafs.Append("server-group-name3", types.YLeaf{"ServerGroupName3", authorization.ServerGroupName3})
    authorization.EntityData.Leafs.Append("server-group-name4", types.YLeaf{"ServerGroupName4", authorization.ServerGroupName4})

    authorization.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(authorization.EntityData)
}

// Aaa_AccountingUpdate
// Configuration related to 'update' accounting
// This type is a presence type.
type Aaa_AccountingUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // newinfo/periodic. The type is AaaAccountingUpdate. This attribute is
    // mandatory.
    Type interface{}

    // Periodic update interval in minutes. The type is interface{} with range:
    // 0..35791394. Units are minute.
    PeriodicInterval interface{}
}

func (accountingUpdate *Aaa_AccountingUpdate) GetEntityData() *types.CommonEntityData {
    accountingUpdate.EntityData.YFilter = accountingUpdate.YFilter
    accountingUpdate.EntityData.YangName = "accounting-update"
    accountingUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountingUpdate.EntityData.ParentYangName = "aaa"
    accountingUpdate.EntityData.SegmentPath = "accounting-update"
    accountingUpdate.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + accountingUpdate.EntityData.SegmentPath
    accountingUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingUpdate.EntityData.Children = types.NewOrderedMap()
    accountingUpdate.EntityData.Leafs = types.NewOrderedMap()
    accountingUpdate.EntityData.Leafs.Append("type", types.YLeaf{"Type", accountingUpdate.Type})
    accountingUpdate.EntityData.Leafs.Append("periodic-interval", types.YLeaf{"PeriodicInterval", accountingUpdate.PeriodicInterval})

    accountingUpdate.EntityData.YListKeys = []string {}

    return &(accountingUpdate.EntityData)
}

// Aaa_Banner
// AAA banner
type Aaa_Banner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA login banner. The type is string.
    Login interface{}
}

func (banner *Aaa_Banner) GetEntityData() *types.CommonEntityData {
    banner.EntityData.YFilter = banner.YFilter
    banner.EntityData.YangName = "banner"
    banner.EntityData.BundleName = "cisco_ios_xr"
    banner.EntityData.ParentYangName = "aaa"
    banner.EntityData.SegmentPath = "banner"
    banner.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + banner.EntityData.SegmentPath
    banner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    banner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    banner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    banner.EntityData.Children = types.NewOrderedMap()
    banner.EntityData.Leafs = types.NewOrderedMap()
    banner.EntityData.Leafs.Append("login", types.YLeaf{"Login", banner.Login})

    banner.EntityData.YListKeys = []string {}

    return &(banner.EntityData)
}

// Aaa_Authentications
// AAA authentication
type Aaa_Authentications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_Authentications_Authentication.
    Authentication []*Aaa_Authentications_Authentication
}

func (authentications *Aaa_Authentications) GetEntityData() *types.CommonEntityData {
    authentications.EntityData.YFilter = authentications.YFilter
    authentications.EntityData.YangName = "authentications"
    authentications.EntityData.BundleName = "cisco_ios_xr"
    authentications.EntityData.ParentYangName = "aaa"
    authentications.EntityData.SegmentPath = "authentications"
    authentications.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + authentications.EntityData.SegmentPath
    authentications.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentications.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentications.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentications.EntityData.Children = types.NewOrderedMap()
    authentications.EntityData.Children.Append("authentication", types.YChild{"Authentication", nil})
    for i := range authentications.Authentication {
        authentications.EntityData.Children.Append(types.GetSegmentPath(authentications.Authentication[i]), types.YChild{"Authentication", authentications.Authentication[i]})
    }
    authentications.EntityData.Leafs = types.NewOrderedMap()

    authentications.EntityData.YListKeys = []string {}

    return &(authentications.EntityData)
}

// Aaa_Authentications_Authentication
// Configurations related to authentication
type Aaa_Authentications_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. login: Authenticate login sessions, ppp:
    // Authenticate ppp sessions. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Type interface{}

    // This attribute is a key. List name for AAA authentication. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Type. The type is AaaMethod.
    Method1 interface{}

    // Method Type. The type is AaaMethod.
    Method2 interface{}

    // Method Type. The type is AaaMethod.
    Method3 interface{}

    // Method Type. The type is AaaMethod.
    Method4 interface{}

    // Server group name. The type is string.
    ServerGroupName1 interface{}

    // Server group name. The type is string.
    ServerGroupName2 interface{}

    // Server group name. The type is string.
    ServerGroupName3 interface{}

    // Server group name. The type is string.
    ServerGroupName4 interface{}
}

func (authentication *Aaa_Authentications_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "authentications"
    authentication.EntityData.SegmentPath = "authentication" + types.AddKeyToken(authentication.Type, "type") + types.AddKeyToken(authentication.Listname, "listname")
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/authentications/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("type", types.YLeaf{"Type", authentication.Type})
    authentication.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", authentication.Listname})
    authentication.EntityData.Leafs.Append("method1", types.YLeaf{"Method1", authentication.Method1})
    authentication.EntityData.Leafs.Append("method2", types.YLeaf{"Method2", authentication.Method2})
    authentication.EntityData.Leafs.Append("method3", types.YLeaf{"Method3", authentication.Method3})
    authentication.EntityData.Leafs.Append("method4", types.YLeaf{"Method4", authentication.Method4})
    authentication.EntityData.Leafs.Append("server-group-name1", types.YLeaf{"ServerGroupName1", authentication.ServerGroupName1})
    authentication.EntityData.Leafs.Append("server-group-name2", types.YLeaf{"ServerGroupName2", authentication.ServerGroupName2})
    authentication.EntityData.Leafs.Append("server-group-name3", types.YLeaf{"ServerGroupName3", authentication.ServerGroupName3})
    authentication.EntityData.Leafs.Append("server-group-name4", types.YLeaf{"ServerGroupName4", authentication.ServerGroupName4})

    authentication.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(authentication.EntityData)
}

// Aaa_PasswordPolicies
// Configure password-policy
type Aaa_PasswordPolicies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password Policy name. The type is slice of
    // Aaa_PasswordPolicies_PasswordPolicy.
    PasswordPolicy []*Aaa_PasswordPolicies_PasswordPolicy
}

func (passwordPolicies *Aaa_PasswordPolicies) GetEntityData() *types.CommonEntityData {
    passwordPolicies.EntityData.YFilter = passwordPolicies.YFilter
    passwordPolicies.EntityData.YangName = "password-policies"
    passwordPolicies.EntityData.BundleName = "cisco_ios_xr"
    passwordPolicies.EntityData.ParentYangName = "aaa"
    passwordPolicies.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-cfg:password-policies"
    passwordPolicies.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + passwordPolicies.EntityData.SegmentPath
    passwordPolicies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passwordPolicies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passwordPolicies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passwordPolicies.EntityData.Children = types.NewOrderedMap()
    passwordPolicies.EntityData.Children.Append("password-policy", types.YChild{"PasswordPolicy", nil})
    for i := range passwordPolicies.PasswordPolicy {
        passwordPolicies.EntityData.Children.Append(types.GetSegmentPath(passwordPolicies.PasswordPolicy[i]), types.YChild{"PasswordPolicy", passwordPolicies.PasswordPolicy[i]})
    }
    passwordPolicies.EntityData.Leafs = types.NewOrderedMap()

    passwordPolicies.EntityData.YListKeys = []string {}

    return &(passwordPolicies.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy
// Password Policy name
type Aaa_PasswordPolicies_PasswordPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Password Policy name. The type is string with
    // length: 1..252.
    Name interface{}

    // Number of lower-case characters. The type is interface{} with range:
    // 0..253.
    LowerCase interface{}

    // Number of upper-case characters. The type is interface{} with range:
    // 0..253.
    UpperCase interface{}

    // Maximum length of the password. The type is interface{} with range: 2..253.
    MaxLength interface{}

    // Number of characters change required between old and new passwords. The
    // type is interface{} with range: 0..253.
    MinCharChange interface{}

    // Number of special characters. The type is interface{} with range: 0..253.
    SpecialChar interface{}

    // Number of numeric characters. The type is interface{} with range: 0..253.
    Numeric interface{}

    // Minimum length of the password. The type is interface{} with range: 2..253.
    MinLength interface{}

    // Number of maximum authentication attempts. The type is interface{} with
    // range: 1..24.
    AuthenMaxAttempts interface{}

    // Liftime of the password.
    Lifetime Aaa_PasswordPolicies_PasswordPolicy_Lifetime

    // Lockout time for the maximum authentication failures.
    LockoutTime Aaa_PasswordPolicies_PasswordPolicy_LockoutTime
}

func (passwordPolicy *Aaa_PasswordPolicies_PasswordPolicy) GetEntityData() *types.CommonEntityData {
    passwordPolicy.EntityData.YFilter = passwordPolicy.YFilter
    passwordPolicy.EntityData.YangName = "password-policy"
    passwordPolicy.EntityData.BundleName = "cisco_ios_xr"
    passwordPolicy.EntityData.ParentYangName = "password-policies"
    passwordPolicy.EntityData.SegmentPath = "password-policy" + types.AddKeyToken(passwordPolicy.Name, "name")
    passwordPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:password-policies/" + passwordPolicy.EntityData.SegmentPath
    passwordPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passwordPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passwordPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passwordPolicy.EntityData.Children = types.NewOrderedMap()
    passwordPolicy.EntityData.Children.Append("lifetime", types.YChild{"Lifetime", &passwordPolicy.Lifetime})
    passwordPolicy.EntityData.Children.Append("lockout-time", types.YChild{"LockoutTime", &passwordPolicy.LockoutTime})
    passwordPolicy.EntityData.Leafs = types.NewOrderedMap()
    passwordPolicy.EntityData.Leafs.Append("name", types.YLeaf{"Name", passwordPolicy.Name})
    passwordPolicy.EntityData.Leafs.Append("lower-case", types.YLeaf{"LowerCase", passwordPolicy.LowerCase})
    passwordPolicy.EntityData.Leafs.Append("upper-case", types.YLeaf{"UpperCase", passwordPolicy.UpperCase})
    passwordPolicy.EntityData.Leafs.Append("max-length", types.YLeaf{"MaxLength", passwordPolicy.MaxLength})
    passwordPolicy.EntityData.Leafs.Append("min-char-change", types.YLeaf{"MinCharChange", passwordPolicy.MinCharChange})
    passwordPolicy.EntityData.Leafs.Append("special-char", types.YLeaf{"SpecialChar", passwordPolicy.SpecialChar})
    passwordPolicy.EntityData.Leafs.Append("numeric", types.YLeaf{"Numeric", passwordPolicy.Numeric})
    passwordPolicy.EntityData.Leafs.Append("min-length", types.YLeaf{"MinLength", passwordPolicy.MinLength})
    passwordPolicy.EntityData.Leafs.Append("authen-max-attempts", types.YLeaf{"AuthenMaxAttempts", passwordPolicy.AuthenMaxAttempts})

    passwordPolicy.EntityData.YListKeys = []string {"Name"}

    return &(passwordPolicy.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy_Lifetime
// Liftime of the password
type Aaa_PasswordPolicies_PasswordPolicy_Lifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of years. The type is interface{} with range: 0..99.
    Years interface{}

    // Number of months. The type is interface{} with range: 0..11.
    Months interface{}

    // Number of hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Number of minutes. The type is interface{} with range: 0..59. Units are
    // minute.
    Minutes interface{}

    // Number of seconds. The type is interface{} with range: 30..59. Units are
    // second.
    Seconds interface{}

    // Number of days. The type is interface{} with range: 0..30. Units are day.
    Days interface{}
}

func (lifetime *Aaa_PasswordPolicies_PasswordPolicy_Lifetime) GetEntityData() *types.CommonEntityData {
    lifetime.EntityData.YFilter = lifetime.YFilter
    lifetime.EntityData.YangName = "lifetime"
    lifetime.EntityData.BundleName = "cisco_ios_xr"
    lifetime.EntityData.ParentYangName = "password-policy"
    lifetime.EntityData.SegmentPath = "lifetime"
    lifetime.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:password-policies/password-policy/" + lifetime.EntityData.SegmentPath
    lifetime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lifetime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lifetime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lifetime.EntityData.Children = types.NewOrderedMap()
    lifetime.EntityData.Leafs = types.NewOrderedMap()
    lifetime.EntityData.Leafs.Append("years", types.YLeaf{"Years", lifetime.Years})
    lifetime.EntityData.Leafs.Append("months", types.YLeaf{"Months", lifetime.Months})
    lifetime.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lifetime.Hours})
    lifetime.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lifetime.Minutes})
    lifetime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lifetime.Seconds})
    lifetime.EntityData.Leafs.Append("days", types.YLeaf{"Days", lifetime.Days})

    lifetime.EntityData.YListKeys = []string {}

    return &(lifetime.EntityData)
}

// Aaa_PasswordPolicies_PasswordPolicy_LockoutTime
// Lockout time for the maximum authentication
// failures
type Aaa_PasswordPolicies_PasswordPolicy_LockoutTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Number of minutes. The type is interface{} with range: 0..59. Units are
    // minute.
    Minutes interface{}

    // Number of seconds. The type is interface{} with range: 0..59. Units are
    // second.
    Seconds interface{}

    // Number of days. The type is interface{} with range: 0..255. Units are day.
    Days interface{}
}

func (lockoutTime *Aaa_PasswordPolicies_PasswordPolicy_LockoutTime) GetEntityData() *types.CommonEntityData {
    lockoutTime.EntityData.YFilter = lockoutTime.YFilter
    lockoutTime.EntityData.YangName = "lockout-time"
    lockoutTime.EntityData.BundleName = "cisco_ios_xr"
    lockoutTime.EntityData.ParentYangName = "password-policy"
    lockoutTime.EntityData.SegmentPath = "lockout-time"
    lockoutTime.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:password-policies/password-policy/" + lockoutTime.EntityData.SegmentPath
    lockoutTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lockoutTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lockoutTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lockoutTime.EntityData.Children = types.NewOrderedMap()
    lockoutTime.EntityData.Leafs = types.NewOrderedMap()
    lockoutTime.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lockoutTime.Hours})
    lockoutTime.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lockoutTime.Minutes})
    lockoutTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lockoutTime.Seconds})
    lockoutTime.EntityData.Leafs.Append("days", types.YLeaf{"Days", lockoutTime.Days})

    lockoutTime.EntityData.YListKeys = []string {}

    return &(lockoutTime.EntityData)
}

// Aaa_ServerGroups
// AAA group definitions
type Aaa_ServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TACACS+ server-group definition.
    TacacsServerGroups Aaa_ServerGroups_TacacsServerGroups

    // RADIUS server group definition.
    RadiusServerGroups Aaa_ServerGroups_RadiusServerGroups

    // DIAMETER server group definition.
    DiameterServerGroups Aaa_ServerGroups_DiameterServerGroups
}

func (serverGroups *Aaa_ServerGroups) GetEntityData() *types.CommonEntityData {
    serverGroups.EntityData.YFilter = serverGroups.YFilter
    serverGroups.EntityData.YangName = "server-groups"
    serverGroups.EntityData.BundleName = "cisco_ios_xr"
    serverGroups.EntityData.ParentYangName = "aaa"
    serverGroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-cfg:server-groups"
    serverGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + serverGroups.EntityData.SegmentPath
    serverGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroups.EntityData.Children = types.NewOrderedMap()
    serverGroups.EntityData.Children.Append("Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups", types.YChild{"TacacsServerGroups", &serverGroups.TacacsServerGroups})
    serverGroups.EntityData.Children.Append("Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups", types.YChild{"RadiusServerGroups", &serverGroups.RadiusServerGroups})
    serverGroups.EntityData.Children.Append("Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups", types.YChild{"DiameterServerGroups", &serverGroups.DiameterServerGroups})
    serverGroups.EntityData.Leafs = types.NewOrderedMap()

    serverGroups.EntityData.YListKeys = []string {}

    return &(serverGroups.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups
// TACACS+ server-group definition
type Aaa_ServerGroups_TacacsServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TACACS+ Server group name. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup.
    TacacsServerGroup []*Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup
}

func (tacacsServerGroups *Aaa_ServerGroups_TacacsServerGroups) GetEntityData() *types.CommonEntityData {
    tacacsServerGroups.EntityData.YFilter = tacacsServerGroups.YFilter
    tacacsServerGroups.EntityData.YangName = "tacacs-server-groups"
    tacacsServerGroups.EntityData.BundleName = "cisco_ios_xr"
    tacacsServerGroups.EntityData.ParentYangName = "server-groups"
    tacacsServerGroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups"
    tacacsServerGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/" + tacacsServerGroups.EntityData.SegmentPath
    tacacsServerGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacsServerGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacsServerGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacsServerGroups.EntityData.Children = types.NewOrderedMap()
    tacacsServerGroups.EntityData.Children.Append("tacacs-server-group", types.YChild{"TacacsServerGroup", nil})
    for i := range tacacsServerGroups.TacacsServerGroup {
        tacacsServerGroups.EntityData.Children.Append(types.GetSegmentPath(tacacsServerGroups.TacacsServerGroup[i]), types.YChild{"TacacsServerGroup", tacacsServerGroups.TacacsServerGroup[i]})
    }
    tacacsServerGroups.EntityData.Leafs = types.NewOrderedMap()

    tacacsServerGroups.EntityData.YListKeys = []string {}

    return &(tacacsServerGroups.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup
// TACACS+ Server group name
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. TACACS+ Server group name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // Specify VRF name of TACACS group. The type is string.
    Vrf interface{}

    // Specify a TACACS+ server.
    Servers Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers

    // List of private TACACS servers present in the group.
    PrivateServers Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers
}

func (tacacsServerGroup *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup) GetEntityData() *types.CommonEntityData {
    tacacsServerGroup.EntityData.YFilter = tacacsServerGroup.YFilter
    tacacsServerGroup.EntityData.YangName = "tacacs-server-group"
    tacacsServerGroup.EntityData.BundleName = "cisco_ios_xr"
    tacacsServerGroup.EntityData.ParentYangName = "tacacs-server-groups"
    tacacsServerGroup.EntityData.SegmentPath = "tacacs-server-group" + types.AddKeyToken(tacacsServerGroup.ServerGroupName, "server-group-name")
    tacacsServerGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups/" + tacacsServerGroup.EntityData.SegmentPath
    tacacsServerGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacsServerGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacsServerGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacsServerGroup.EntityData.Children = types.NewOrderedMap()
    tacacsServerGroup.EntityData.Children.Append("servers", types.YChild{"Servers", &tacacsServerGroup.Servers})
    tacacsServerGroup.EntityData.Children.Append("private-servers", types.YChild{"PrivateServers", &tacacsServerGroup.PrivateServers})
    tacacsServerGroup.EntityData.Leafs = types.NewOrderedMap()
    tacacsServerGroup.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", tacacsServerGroup.ServerGroupName})
    tacacsServerGroup.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", tacacsServerGroup.Vrf})

    tacacsServerGroup.EntityData.YListKeys = []string {"ServerGroupName"}

    return &(tacacsServerGroup.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers
// Specify a TACACS+ server
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server.
    Server []*Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "tacacs-server-group"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups/tacacs-server-group/" + servers.EntityData.SegmentPath
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = types.NewOrderedMap()
    servers.EntityData.Children.Append("server", types.YChild{"Server", nil})
    for i := range servers.Server {
        servers.EntityData.Children.Append(types.GetSegmentPath(servers.Server[i]), types.YChild{"Server", servers.Server[i]})
    }
    servers.EntityData.Leafs = types.NewOrderedMap()

    servers.EntityData.YListKeys = []string {}

    return &(servers.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (server *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server" + types.AddKeyToken(server.OrderingIndex, "ordering-index") + types.AddKeyToken(server.IpAddress, "ip-address")
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups/tacacs-server-group/servers/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", server.OrderingIndex})
    server.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", server.IpAddress})

    server.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress"}

    return &(server.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers
// List of private TACACS servers present in the
// group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A private server to include in the server group. The type is slice of
    // Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer.
    PrivateServer []*Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer
}

func (privateServers *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers) GetEntityData() *types.CommonEntityData {
    privateServers.EntityData.YFilter = privateServers.YFilter
    privateServers.EntityData.YangName = "private-servers"
    privateServers.EntityData.BundleName = "cisco_ios_xr"
    privateServers.EntityData.ParentYangName = "tacacs-server-group"
    privateServers.EntityData.SegmentPath = "private-servers"
    privateServers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups/tacacs-server-group/" + privateServers.EntityData.SegmentPath
    privateServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateServers.EntityData.Children = types.NewOrderedMap()
    privateServers.EntityData.Children.Append("private-server", types.YChild{"PrivateServer", nil})
    for i := range privateServers.PrivateServer {
        privateServers.EntityData.Children.Append(types.GetSegmentPath(privateServers.PrivateServer[i]), types.YChild{"PrivateServer", privateServers.PrivateServer[i]})
    }
    privateServers.EntityData.Leafs = types.NewOrderedMap()

    privateServers.EntityData.YListKeys = []string {}

    return &(privateServers.EntityData)
}

// Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer
// A private server to include in the server
// group
type Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port number (standard 49). The type is interface{}
    // with range: 1..65535.
    PortNumber interface{}

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}
}

func (privateServer *Aaa_ServerGroups_TacacsServerGroups_TacacsServerGroup_PrivateServers_PrivateServer) GetEntityData() *types.CommonEntityData {
    privateServer.EntityData.YFilter = privateServer.YFilter
    privateServer.EntityData.YangName = "private-server"
    privateServer.EntityData.BundleName = "cisco_ios_xr"
    privateServer.EntityData.ParentYangName = "private-servers"
    privateServer.EntityData.SegmentPath = "private-server" + types.AddKeyToken(privateServer.OrderingIndex, "ordering-index") + types.AddKeyToken(privateServer.IpAddress, "ip-address") + types.AddKeyToken(privateServer.PortNumber, "port-number")
    privateServer.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs-server-groups/tacacs-server-group/private-servers/" + privateServer.EntityData.SegmentPath
    privateServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateServer.EntityData.Children = types.NewOrderedMap()
    privateServer.EntityData.Leafs = types.NewOrderedMap()
    privateServer.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", privateServer.OrderingIndex})
    privateServer.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", privateServer.IpAddress})
    privateServer.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", privateServer.PortNumber})
    privateServer.EntityData.Leafs.Append("key", types.YLeaf{"Key", privateServer.Key})
    privateServer.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", privateServer.Timeout})

    privateServer.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress", "PortNumber"}

    return &(privateServer.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups
// RADIUS server group definition
type Aaa_ServerGroups_RadiusServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RADIUS server group name. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup.
    RadiusServerGroup []*Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup
}

func (radiusServerGroups *Aaa_ServerGroups_RadiusServerGroups) GetEntityData() *types.CommonEntityData {
    radiusServerGroups.EntityData.YFilter = radiusServerGroups.YFilter
    radiusServerGroups.EntityData.YangName = "radius-server-groups"
    radiusServerGroups.EntityData.BundleName = "cisco_ios_xr"
    radiusServerGroups.EntityData.ParentYangName = "server-groups"
    radiusServerGroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups"
    radiusServerGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/" + radiusServerGroups.EntityData.SegmentPath
    radiusServerGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radiusServerGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radiusServerGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radiusServerGroups.EntityData.Children = types.NewOrderedMap()
    radiusServerGroups.EntityData.Children.Append("radius-server-group", types.YChild{"RadiusServerGroup", nil})
    for i := range radiusServerGroups.RadiusServerGroup {
        radiusServerGroups.EntityData.Children.Append(types.GetSegmentPath(radiusServerGroups.RadiusServerGroup[i]), types.YChild{"RadiusServerGroup", radiusServerGroups.RadiusServerGroup[i]})
    }
    radiusServerGroups.EntityData.Leafs = types.NewOrderedMap()

    radiusServerGroups.EntityData.YListKeys = []string {}

    return &(radiusServerGroups.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup
// RADIUS server group name
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. RADIUS server group name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // This indicates the length of time (in minutes) for which RADIUS servers
    // present in this group remain marked as dead. The type is interface{} with
    // range: 1..1440. Units are minute.
    DeadTime interface{}

    // Specify interface for source address in RADIUS packets. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // Specify VRF name of RADIUS group. The type is string.
    Vrf interface{}

    // List of filters in server group.
    Accounting Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting

    // List of RADIUS servers present in the group.
    Servers Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers

    // List of private RADIUS servers present in the group.
    PrivateServers Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers

    // Radius throttling options.
    ServerGroupThrottle Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle

    // Radius load-balancing options.
    LoadBalance Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance

    // List of filters in server group.
    Authorization Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization
}

func (radiusServerGroup *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup) GetEntityData() *types.CommonEntityData {
    radiusServerGroup.EntityData.YFilter = radiusServerGroup.YFilter
    radiusServerGroup.EntityData.YangName = "radius-server-group"
    radiusServerGroup.EntityData.BundleName = "cisco_ios_xr"
    radiusServerGroup.EntityData.ParentYangName = "radius-server-groups"
    radiusServerGroup.EntityData.SegmentPath = "radius-server-group" + types.AddKeyToken(radiusServerGroup.ServerGroupName, "server-group-name")
    radiusServerGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/" + radiusServerGroup.EntityData.SegmentPath
    radiusServerGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radiusServerGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radiusServerGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radiusServerGroup.EntityData.Children = types.NewOrderedMap()
    radiusServerGroup.EntityData.Children.Append("accounting", types.YChild{"Accounting", &radiusServerGroup.Accounting})
    radiusServerGroup.EntityData.Children.Append("servers", types.YChild{"Servers", &radiusServerGroup.Servers})
    radiusServerGroup.EntityData.Children.Append("private-servers", types.YChild{"PrivateServers", &radiusServerGroup.PrivateServers})
    radiusServerGroup.EntityData.Children.Append("server-group-throttle", types.YChild{"ServerGroupThrottle", &radiusServerGroup.ServerGroupThrottle})
    radiusServerGroup.EntityData.Children.Append("load-balance", types.YChild{"LoadBalance", &radiusServerGroup.LoadBalance})
    radiusServerGroup.EntityData.Children.Append("authorization", types.YChild{"Authorization", &radiusServerGroup.Authorization})
    radiusServerGroup.EntityData.Leafs = types.NewOrderedMap()
    radiusServerGroup.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", radiusServerGroup.ServerGroupName})
    radiusServerGroup.EntityData.Leafs.Append("dead-time", types.YLeaf{"DeadTime", radiusServerGroup.DeadTime})
    radiusServerGroup.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", radiusServerGroup.SourceInterface})
    radiusServerGroup.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", radiusServerGroup.Vrf})

    radiusServerGroup.EntityData.YListKeys = []string {"ServerGroupName"}

    return &(radiusServerGroup.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting
// List of filters in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a filter in server group.
    Request Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request

    // Specify a filter in server group.
    Reply Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply
}

func (accounting *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "radius-server-group"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Children.Append("request", types.YChild{"Request", &accounting.Request})
    accounting.EntityData.Children.Append("reply", types.YChild{"Reply", &accounting.Reply})
    accounting.EntityData.Leafs = types.NewOrderedMap()

    accounting.EntityData.YListKeys = []string {}

    return &(accounting.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "accounting"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/accounting/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("action", types.YLeaf{"Action", request.Action})
    request.EntityData.Leafs.Append("attribute-list-name", types.YLeaf{"AttributeListName", request.AttributeListName})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Accounting_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "accounting"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/accounting/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("action", types.YLeaf{"Action", reply.Action})
    reply.EntityData.Leafs.Append("attribute-list-name", types.YLeaf{"AttributeListName", reply.AttributeListName})

    reply.EntityData.YListKeys = []string {}

    return &(reply.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers
// List of RADIUS servers present in the group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server.
    Server []*Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "radius-server-group"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + servers.EntityData.SegmentPath
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = types.NewOrderedMap()
    servers.EntityData.Children.Append("server", types.YChild{"Server", nil})
    for i := range servers.Server {
        servers.EntityData.Children.Append(types.GetSegmentPath(servers.Server[i]), types.YChild{"Server", servers.Server[i]})
    }
    servers.EntityData.Leafs = types.NewOrderedMap()

    servers.EntityData.YListKeys = []string {}

    return &(servers.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}
}

func (server *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server" + types.AddKeyToken(server.OrderingIndex, "ordering-index") + types.AddKeyToken(server.IpAddress, "ip-address") + types.AddKeyToken(server.AuthPortNumber, "auth-port-number") + types.AddKeyToken(server.AcctPortNumber, "acct-port-number")
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/servers/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", server.OrderingIndex})
    server.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", server.IpAddress})
    server.EntityData.Leafs.Append("auth-port-number", types.YLeaf{"AuthPortNumber", server.AuthPortNumber})
    server.EntityData.Leafs.Append("acct-port-number", types.YLeaf{"AcctPortNumber", server.AcctPortNumber})

    server.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress", "AuthPortNumber", "AcctPortNumber"}

    return &(server.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers
// List of private RADIUS servers present in the
// group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A private server to include in the server group. The type is slice of
    // Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer.
    PrivateServer []*Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer
}

func (privateServers *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers) GetEntityData() *types.CommonEntityData {
    privateServers.EntityData.YFilter = privateServers.YFilter
    privateServers.EntityData.YangName = "private-servers"
    privateServers.EntityData.BundleName = "cisco_ios_xr"
    privateServers.EntityData.ParentYangName = "radius-server-group"
    privateServers.EntityData.SegmentPath = "private-servers"
    privateServers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + privateServers.EntityData.SegmentPath
    privateServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateServers.EntityData.Children = types.NewOrderedMap()
    privateServers.EntityData.Children.Append("private-server", types.YChild{"PrivateServer", nil})
    for i := range privateServers.PrivateServer {
        privateServers.EntityData.Children.Append(types.GetSegmentPath(privateServers.PrivateServer[i]), types.YChild{"PrivateServer", privateServers.PrivateServer[i]})
    }
    privateServers.EntityData.Leafs = types.NewOrderedMap()

    privateServers.EntityData.YListKeys = []string {}

    return &(privateServers.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer
// A private server to include in the server
// group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    PrivateTimeout interface{}

    // Ignore Accounting port. The type is bool.
    IgnoreAccountingPort interface{}

    // Number of times to retransmit a request to the RADIUS server. The type is
    // interface{} with range: 1..100. The default value is 3.
    PrivateRetransmit interface{}

    // Idle time for the radius Server. The type is interface{} with range: 1..60.
    // The default value is 5.
    IdleTime interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    PrivateKey interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Ignore authentication Port. The type is bool.
    IgnoreAuthPort interface{}
}

func (privateServer *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_PrivateServers_PrivateServer) GetEntityData() *types.CommonEntityData {
    privateServer.EntityData.YFilter = privateServer.YFilter
    privateServer.EntityData.YangName = "private-server"
    privateServer.EntityData.BundleName = "cisco_ios_xr"
    privateServer.EntityData.ParentYangName = "private-servers"
    privateServer.EntityData.SegmentPath = "private-server" + types.AddKeyToken(privateServer.OrderingIndex, "ordering-index") + types.AddKeyToken(privateServer.IpAddress, "ip-address") + types.AddKeyToken(privateServer.AuthPortNumber, "auth-port-number") + types.AddKeyToken(privateServer.AcctPortNumber, "acct-port-number")
    privateServer.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/private-servers/" + privateServer.EntityData.SegmentPath
    privateServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateServer.EntityData.Children = types.NewOrderedMap()
    privateServer.EntityData.Leafs = types.NewOrderedMap()
    privateServer.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", privateServer.OrderingIndex})
    privateServer.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", privateServer.IpAddress})
    privateServer.EntityData.Leafs.Append("auth-port-number", types.YLeaf{"AuthPortNumber", privateServer.AuthPortNumber})
    privateServer.EntityData.Leafs.Append("acct-port-number", types.YLeaf{"AcctPortNumber", privateServer.AcctPortNumber})
    privateServer.EntityData.Leafs.Append("private-timeout", types.YLeaf{"PrivateTimeout", privateServer.PrivateTimeout})
    privateServer.EntityData.Leafs.Append("ignore-accounting-port", types.YLeaf{"IgnoreAccountingPort", privateServer.IgnoreAccountingPort})
    privateServer.EntityData.Leafs.Append("private-retransmit", types.YLeaf{"PrivateRetransmit", privateServer.PrivateRetransmit})
    privateServer.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", privateServer.IdleTime})
    privateServer.EntityData.Leafs.Append("private-key", types.YLeaf{"PrivateKey", privateServer.PrivateKey})
    privateServer.EntityData.Leafs.Append("username", types.YLeaf{"Username", privateServer.Username})
    privateServer.EntityData.Leafs.Append("ignore-auth-port", types.YLeaf{"IgnoreAuthPort", privateServer.IgnoreAuthPort})

    privateServer.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress", "AuthPortNumber", "AcctPortNumber"}

    return &(privateServer.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle
// Radius throttling options
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // To flow control the number of access requests sent to a radius server. The
    // type is interface{} with range: 0..65535. The default value is 0.
    Access interface{}

    // To flow control the number of accounting requests sent to a radius server.
    // The type is interface{} with range: 0..65535. The default value is 0.
    Accounting interface{}

    // Specify the number of timeouts exceeding which a throttled access request
    // is dropped. The type is interface{} with range: 1..10. The default value is
    // 3.
    AccessTimeout interface{}
}

func (serverGroupThrottle *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_ServerGroupThrottle) GetEntityData() *types.CommonEntityData {
    serverGroupThrottle.EntityData.YFilter = serverGroupThrottle.YFilter
    serverGroupThrottle.EntityData.YangName = "server-group-throttle"
    serverGroupThrottle.EntityData.BundleName = "cisco_ios_xr"
    serverGroupThrottle.EntityData.ParentYangName = "radius-server-group"
    serverGroupThrottle.EntityData.SegmentPath = "server-group-throttle"
    serverGroupThrottle.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + serverGroupThrottle.EntityData.SegmentPath
    serverGroupThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverGroupThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverGroupThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverGroupThrottle.EntityData.Children = types.NewOrderedMap()
    serverGroupThrottle.EntityData.Leafs = types.NewOrderedMap()
    serverGroupThrottle.EntityData.Leafs.Append("access", types.YLeaf{"Access", serverGroupThrottle.Access})
    serverGroupThrottle.EntityData.Leafs.Append("accounting", types.YLeaf{"Accounting", serverGroupThrottle.Accounting})
    serverGroupThrottle.EntityData.Leafs.Append("access-timeout", types.YLeaf{"AccessTimeout", serverGroupThrottle.AccessTimeout})

    serverGroupThrottle.EntityData.YListKeys = []string {}

    return &(serverGroupThrottle.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance
// Radius load-balancing options
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Method by which the next host will be picked.
    Method Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method
}

func (loadBalance *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance) GetEntityData() *types.CommonEntityData {
    loadBalance.EntityData.YFilter = loadBalance.YFilter
    loadBalance.EntityData.YangName = "load-balance"
    loadBalance.EntityData.BundleName = "cisco_ios_xr"
    loadBalance.EntityData.ParentYangName = "radius-server-group"
    loadBalance.EntityData.SegmentPath = "load-balance"
    loadBalance.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + loadBalance.EntityData.SegmentPath
    loadBalance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalance.EntityData.Children = types.NewOrderedMap()
    loadBalance.EntityData.Children.Append("method", types.YChild{"Method", &loadBalance.Method})
    loadBalance.EntityData.Leafs = types.NewOrderedMap()

    loadBalance.EntityData.YListKeys = []string {}

    return &(loadBalance.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method
// Method by which the next host will be picked
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Batch size for selection of the server.
    Name Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name
}

func (method *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method) GetEntityData() *types.CommonEntityData {
    method.EntityData.YFilter = method.YFilter
    method.EntityData.YangName = "method"
    method.EntityData.BundleName = "cisco_ios_xr"
    method.EntityData.ParentYangName = "load-balance"
    method.EntityData.SegmentPath = "method"
    method.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/load-balance/" + method.EntityData.SegmentPath
    method.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    method.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    method.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    method.EntityData.Children = types.NewOrderedMap()
    method.EntityData.Children.Append("name", types.YChild{"Name", &method.Name})
    method.EntityData.Leafs = types.NewOrderedMap()

    method.EntityData.YListKeys = []string {}

    return &(method.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name
// Batch size for selection of the server
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pick the server with the least transactions outstanding. The type is
    // interface{} with range: 0..4294967295. The default value is 4.
    LeastOutstanding interface{}

    // Batch size for selection of the server. The type is interface{} with range:
    // 1..1500. The default value is 25.
    BatchSize interface{}

    // Disable preferred server for this Server Group. The type is bool. The
    // default value is true.
    IgnorePreferredServer interface{}
}

func (name *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_LoadBalance_Method_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "method"
    name.EntityData.SegmentPath = "name"
    name.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/load-balance/method/" + name.EntityData.SegmentPath
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("least-outstanding", types.YLeaf{"LeastOutstanding", name.LeastOutstanding})
    name.EntityData.Leafs.Append("batch-size", types.YLeaf{"BatchSize", name.BatchSize})
    name.EntityData.Leafs.Append("ignore-preferred-server", types.YLeaf{"IgnorePreferredServer", name.IgnorePreferredServer})

    name.EntityData.YListKeys = []string {}

    return &(name.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization
// List of filters in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify a filter in server group.
    Request Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request

    // Specify a filter in server group.
    Reply Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply
}

func (authorization *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "radius-server-group"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Children.Append("request", types.YChild{"Request", &authorization.Request})
    authorization.EntityData.Children.Append("reply", types.YChild{"Reply", &authorization.Reply})
    authorization.EntityData.Leafs = types.NewOrderedMap()

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (request *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "authorization"
    request.EntityData.SegmentPath = "request"
    request.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/authorization/" + request.EntityData.SegmentPath
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = types.NewOrderedMap()
    request.EntityData.Leafs = types.NewOrderedMap()
    request.EntityData.Leafs.Append("action", types.YLeaf{"Action", request.Action})
    request.EntityData.Leafs.Append("attribute-list-name", types.YLeaf{"AttributeListName", request.AttributeListName})

    request.EntityData.YListKeys = []string {}

    return &(request.EntityData)
}

// Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply
// Specify a filter in server group
type Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the attribute list type accept or reject. The type is AaaAction.
    Action interface{}

    // Name of RADIUS attribute list. The type is string.
    AttributeListName interface{}
}

func (reply *Aaa_ServerGroups_RadiusServerGroups_RadiusServerGroup_Authorization_Reply) GetEntityData() *types.CommonEntityData {
    reply.EntityData.YFilter = reply.YFilter
    reply.EntityData.YangName = "reply"
    reply.EntityData.BundleName = "cisco_ios_xr"
    reply.EntityData.ParentYangName = "authorization"
    reply.EntityData.SegmentPath = "reply"
    reply.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius-server-groups/radius-server-group/authorization/" + reply.EntityData.SegmentPath
    reply.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reply.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reply.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reply.EntityData.Children = types.NewOrderedMap()
    reply.EntityData.Leafs = types.NewOrderedMap()
    reply.EntityData.Leafs.Append("action", types.YLeaf{"Action", reply.Action})
    reply.EntityData.Leafs.Append("attribute-list-name", types.YLeaf{"AttributeListName", reply.AttributeListName})

    reply.EntityData.YListKeys = []string {}

    return &(reply.EntityData)
}

// Aaa_ServerGroups_DiameterServerGroups
// DIAMETER server group definition
type Aaa_ServerGroups_DiameterServerGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DIAMETER server group name. The type is slice of
    // Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup.
    DiameterServerGroup []*Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup
}

func (diameterServerGroups *Aaa_ServerGroups_DiameterServerGroups) GetEntityData() *types.CommonEntityData {
    diameterServerGroups.EntityData.YFilter = diameterServerGroups.YFilter
    diameterServerGroups.EntityData.YangName = "diameter-server-groups"
    diameterServerGroups.EntityData.BundleName = "cisco_ios_xr"
    diameterServerGroups.EntityData.ParentYangName = "server-groups"
    diameterServerGroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups"
    diameterServerGroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/" + diameterServerGroups.EntityData.SegmentPath
    diameterServerGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameterServerGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameterServerGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameterServerGroups.EntityData.Children = types.NewOrderedMap()
    diameterServerGroups.EntityData.Children.Append("diameter-server-group", types.YChild{"DiameterServerGroup", nil})
    for i := range diameterServerGroups.DiameterServerGroup {
        diameterServerGroups.EntityData.Children.Append(types.GetSegmentPath(diameterServerGroups.DiameterServerGroup[i]), types.YChild{"DiameterServerGroup", diameterServerGroups.DiameterServerGroup[i]})
    }
    diameterServerGroups.EntityData.Leafs = types.NewOrderedMap()

    diameterServerGroups.EntityData.YListKeys = []string {}

    return &(diameterServerGroups.EntityData)
}

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup
// DIAMETER server group name
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. DIAMETER server group name. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServerGroupName interface{}

    // List of DIAMETER servers present in the group.
    Servers Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers
}

func (diameterServerGroup *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup) GetEntityData() *types.CommonEntityData {
    diameterServerGroup.EntityData.YFilter = diameterServerGroup.YFilter
    diameterServerGroup.EntityData.YangName = "diameter-server-group"
    diameterServerGroup.EntityData.BundleName = "cisco_ios_xr"
    diameterServerGroup.EntityData.ParentYangName = "diameter-server-groups"
    diameterServerGroup.EntityData.SegmentPath = "diameter-server-group" + types.AddKeyToken(diameterServerGroup.ServerGroupName, "server-group-name")
    diameterServerGroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups/" + diameterServerGroup.EntityData.SegmentPath
    diameterServerGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameterServerGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameterServerGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameterServerGroup.EntityData.Children = types.NewOrderedMap()
    diameterServerGroup.EntityData.Children.Append("servers", types.YChild{"Servers", &diameterServerGroup.Servers})
    diameterServerGroup.EntityData.Leafs = types.NewOrderedMap()
    diameterServerGroup.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", diameterServerGroup.ServerGroupName})

    diameterServerGroup.EntityData.YListKeys = []string {"ServerGroupName"}

    return &(diameterServerGroup.EntityData)
}

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers
// List of DIAMETER servers present in the group
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A server to include in the server group. The type is slice of
    // Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server.
    Server []*Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server
}

func (servers *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers) GetEntityData() *types.CommonEntityData {
    servers.EntityData.YFilter = servers.YFilter
    servers.EntityData.YangName = "servers"
    servers.EntityData.BundleName = "cisco_ios_xr"
    servers.EntityData.ParentYangName = "diameter-server-group"
    servers.EntityData.SegmentPath = "servers"
    servers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups/diameter-server-group/" + servers.EntityData.SegmentPath
    servers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    servers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    servers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    servers.EntityData.Children = types.NewOrderedMap()
    servers.EntityData.Children.Append("server", types.YChild{"Server", nil})
    for i := range servers.Server {
        servers.EntityData.Children.Append(types.GetSegmentPath(servers.Server[i]), types.YChild{"Server", servers.Server[i]})
    }
    servers.EntityData.Leafs = types.NewOrderedMap()

    servers.EntityData.YListKeys = []string {}

    return &(servers.EntityData)
}

// Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server
// A server to include in the server group
type Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. Name for the diameter peer configuration. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PeerName interface{}
}

func (server *Aaa_ServerGroups_DiameterServerGroups_DiameterServerGroup_Servers_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "servers"
    server.EntityData.SegmentPath = "server" + types.AddKeyToken(server.OrderingIndex, "ordering-index") + types.AddKeyToken(server.PeerName, "peer-name")
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:server-groups/Cisco-IOS-XR-aaa-diameter-cfg:diameter-server-groups/diameter-server-group/servers/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", server.OrderingIndex})
    server.EntityData.Leafs.Append("peer-name", types.YLeaf{"PeerName", server.PeerName})

    server.EntityData.YListKeys = []string {"OrderingIndex", "PeerName"}

    return &(server.EntityData)
}

// Aaa_Usernames
// Configure local usernames
type Aaa_Usernames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local username. The type is slice of Aaa_Usernames_Username.
    Username []*Aaa_Usernames_Username
}

func (usernames *Aaa_Usernames) GetEntityData() *types.CommonEntityData {
    usernames.EntityData.YFilter = usernames.YFilter
    usernames.EntityData.YangName = "usernames"
    usernames.EntityData.BundleName = "cisco_ios_xr"
    usernames.EntityData.ParentYangName = "aaa"
    usernames.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-cfg:usernames"
    usernames.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + usernames.EntityData.SegmentPath
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
// Local username
type Aaa_Usernames_Username struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the users in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. Username. The type is string.
    Name interface{}

    // Specify the secret for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Secret interface{}

    // Specify the password for the user. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Mention Password policy for the user.
    PasswordPolicy Aaa_Usernames_Username_PasswordPolicy

    // Specify the usergroup to which this user belongs.
    UsergroupUnderUsernames Aaa_Usernames_Username_UsergroupUnderUsernames
}

func (username *Aaa_Usernames_Username) GetEntityData() *types.CommonEntityData {
    username.EntityData.YFilter = username.YFilter
    username.EntityData.YangName = "username"
    username.EntityData.BundleName = "cisco_ios_xr"
    username.EntityData.ParentYangName = "usernames"
    username.EntityData.SegmentPath = "username" + types.AddKeyToken(username.OrderingIndex, "ordering-index") + types.AddKeyToken(username.Name, "name")
    username.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usernames/" + username.EntityData.SegmentPath
    username.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    username.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    username.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    username.EntityData.Children = types.NewOrderedMap()
    username.EntityData.Children.Append("password-policy", types.YChild{"PasswordPolicy", &username.PasswordPolicy})
    username.EntityData.Children.Append("usergroup-under-usernames", types.YChild{"UsergroupUnderUsernames", &username.UsergroupUnderUsernames})
    username.EntityData.Leafs = types.NewOrderedMap()
    username.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", username.OrderingIndex})
    username.EntityData.Leafs.Append("name", types.YLeaf{"Name", username.Name})
    username.EntityData.Leafs.Append("secret", types.YLeaf{"Secret", username.Secret})
    username.EntityData.Leafs.Append("password", types.YLeaf{"Password", username.Password})

    username.EntityData.YListKeys = []string {"OrderingIndex", "Name"}

    return &(username.EntityData)
}

// Aaa_Usernames_Username_PasswordPolicy
// Mention Password policy for the user
type Aaa_Usernames_Username_PasswordPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Password Policy name. The type is string with length: 1..252.
    Name interface{}

    // The user's password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}
}

func (passwordPolicy *Aaa_Usernames_Username_PasswordPolicy) GetEntityData() *types.CommonEntityData {
    passwordPolicy.EntityData.YFilter = passwordPolicy.YFilter
    passwordPolicy.EntityData.YangName = "password-policy"
    passwordPolicy.EntityData.BundleName = "cisco_ios_xr"
    passwordPolicy.EntityData.ParentYangName = "username"
    passwordPolicy.EntityData.SegmentPath = "password-policy"
    passwordPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usernames/username/" + passwordPolicy.EntityData.SegmentPath
    passwordPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passwordPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passwordPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passwordPolicy.EntityData.Children = types.NewOrderedMap()
    passwordPolicy.EntityData.Leafs = types.NewOrderedMap()
    passwordPolicy.EntityData.Leafs.Append("name", types.YLeaf{"Name", passwordPolicy.Name})
    passwordPolicy.EntityData.Leafs.Append("password", types.YLeaf{"Password", passwordPolicy.Password})

    passwordPolicy.EntityData.YListKeys = []string {}

    return &(passwordPolicy.EntityData)
}

// Aaa_Usernames_Username_UsergroupUnderUsernames
// Specify the usergroup to which this user
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
    usergroupUnderUsernames.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usernames/username/" + usergroupUnderUsernames.EntityData.SegmentPath
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
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (usergroupUnderUsername *Aaa_Usernames_Username_UsergroupUnderUsernames_UsergroupUnderUsername) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsername.EntityData.YFilter = usergroupUnderUsername.YFilter
    usergroupUnderUsername.EntityData.YangName = "usergroup-under-username"
    usergroupUnderUsername.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsername.EntityData.ParentYangName = "usergroup-under-usernames"
    usergroupUnderUsername.EntityData.SegmentPath = "usergroup-under-username" + types.AddKeyToken(usergroupUnderUsername.Name, "name")
    usergroupUnderUsername.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usernames/username/usergroup-under-usernames/" + usergroupUnderUsername.EntityData.SegmentPath
    usergroupUnderUsername.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsername.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsername.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsername.EntityData.Children = types.NewOrderedMap()
    usergroupUnderUsername.EntityData.Leafs = types.NewOrderedMap()
    usergroupUnderUsername.EntityData.Leafs.Append("name", types.YLeaf{"Name", usergroupUnderUsername.Name})

    usergroupUnderUsername.EntityData.YListKeys = []string {"Name"}

    return &(usergroupUnderUsername.EntityData)
}

// Aaa_Taskgroups
// Specify a taskgroup to inherit from
type Aaa_Taskgroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Taskgroup name. The type is slice of Aaa_Taskgroups_Taskgroup.
    Taskgroup []*Aaa_Taskgroups_Taskgroup
}

func (taskgroups *Aaa_Taskgroups) GetEntityData() *types.CommonEntityData {
    taskgroups.EntityData.YFilter = taskgroups.YFilter
    taskgroups.EntityData.YangName = "taskgroups"
    taskgroups.EntityData.BundleName = "cisco_ios_xr"
    taskgroups.EntityData.ParentYangName = "aaa"
    taskgroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-cfg:taskgroups"
    taskgroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + taskgroups.EntityData.SegmentPath
    taskgroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroups.EntityData.Children = types.NewOrderedMap()
    taskgroups.EntityData.Children.Append("taskgroup", types.YChild{"Taskgroup", nil})
    for i := range taskgroups.Taskgroup {
        taskgroups.EntityData.Children.Append(types.GetSegmentPath(taskgroups.Taskgroup[i]), types.YChild{"Taskgroup", taskgroups.Taskgroup[i]})
    }
    taskgroups.EntityData.Leafs = types.NewOrderedMap()

    taskgroups.EntityData.YListKeys = []string {}

    return &(taskgroups.EntityData)
}

// Aaa_Taskgroups_Taskgroup
// Taskgroup name
type Aaa_Taskgroups_Taskgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Taskgroup name. The type is string.
    Name interface{}

    // Description for the task group. The type is string.
    Description interface{}

    // Specify a taskgroup to inherit from.
    TaskgroupUnderTaskgroups Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups

    // Specify task IDs to be part of this group.
    Tasks Aaa_Taskgroups_Taskgroup_Tasks
}

func (taskgroup *Aaa_Taskgroups_Taskgroup) GetEntityData() *types.CommonEntityData {
    taskgroup.EntityData.YFilter = taskgroup.YFilter
    taskgroup.EntityData.YangName = "taskgroup"
    taskgroup.EntityData.BundleName = "cisco_ios_xr"
    taskgroup.EntityData.ParentYangName = "taskgroups"
    taskgroup.EntityData.SegmentPath = "taskgroup" + types.AddKeyToken(taskgroup.Name, "name")
    taskgroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:taskgroups/" + taskgroup.EntityData.SegmentPath
    taskgroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroup.EntityData.Children = types.NewOrderedMap()
    taskgroup.EntityData.Children.Append("taskgroup-under-taskgroups", types.YChild{"TaskgroupUnderTaskgroups", &taskgroup.TaskgroupUnderTaskgroups})
    taskgroup.EntityData.Children.Append("tasks", types.YChild{"Tasks", &taskgroup.Tasks})
    taskgroup.EntityData.Leafs = types.NewOrderedMap()
    taskgroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", taskgroup.Name})
    taskgroup.EntityData.Leafs.Append("description", types.YLeaf{"Description", taskgroup.Description})

    taskgroup.EntityData.YListKeys = []string {"Name"}

    return &(taskgroup.EntityData)
}

// Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups
// Specify a taskgroup to inherit from
type Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task group to include. The type is slice of
    // Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup.
    TaskgroupUnderTaskgroup []*Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup
}

func (taskgroupUnderTaskgroups *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups) GetEntityData() *types.CommonEntityData {
    taskgroupUnderTaskgroups.EntityData.YFilter = taskgroupUnderTaskgroups.YFilter
    taskgroupUnderTaskgroups.EntityData.YangName = "taskgroup-under-taskgroups"
    taskgroupUnderTaskgroups.EntityData.BundleName = "cisco_ios_xr"
    taskgroupUnderTaskgroups.EntityData.ParentYangName = "taskgroup"
    taskgroupUnderTaskgroups.EntityData.SegmentPath = "taskgroup-under-taskgroups"
    taskgroupUnderTaskgroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:taskgroups/taskgroup/" + taskgroupUnderTaskgroups.EntityData.SegmentPath
    taskgroupUnderTaskgroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroupUnderTaskgroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroupUnderTaskgroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroupUnderTaskgroups.EntityData.Children = types.NewOrderedMap()
    taskgroupUnderTaskgroups.EntityData.Children.Append("taskgroup-under-taskgroup", types.YChild{"TaskgroupUnderTaskgroup", nil})
    for i := range taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup {
        taskgroupUnderTaskgroups.EntityData.Children.Append(types.GetSegmentPath(taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup[i]), types.YChild{"TaskgroupUnderTaskgroup", taskgroupUnderTaskgroups.TaskgroupUnderTaskgroup[i]})
    }
    taskgroupUnderTaskgroups.EntityData.Leafs = types.NewOrderedMap()

    taskgroupUnderTaskgroups.EntityData.YListKeys = []string {}

    return &(taskgroupUnderTaskgroups.EntityData)
}

// Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup
// Name of the task group to include
type Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the task group to include. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (taskgroupUnderTaskgroup *Aaa_Taskgroups_Taskgroup_TaskgroupUnderTaskgroups_TaskgroupUnderTaskgroup) GetEntityData() *types.CommonEntityData {
    taskgroupUnderTaskgroup.EntityData.YFilter = taskgroupUnderTaskgroup.YFilter
    taskgroupUnderTaskgroup.EntityData.YangName = "taskgroup-under-taskgroup"
    taskgroupUnderTaskgroup.EntityData.BundleName = "cisco_ios_xr"
    taskgroupUnderTaskgroup.EntityData.ParentYangName = "taskgroup-under-taskgroups"
    taskgroupUnderTaskgroup.EntityData.SegmentPath = "taskgroup-under-taskgroup" + types.AddKeyToken(taskgroupUnderTaskgroup.Name, "name")
    taskgroupUnderTaskgroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:taskgroups/taskgroup/taskgroup-under-taskgroups/" + taskgroupUnderTaskgroup.EntityData.SegmentPath
    taskgroupUnderTaskgroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroupUnderTaskgroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroupUnderTaskgroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroupUnderTaskgroup.EntityData.Children = types.NewOrderedMap()
    taskgroupUnderTaskgroup.EntityData.Leafs = types.NewOrderedMap()
    taskgroupUnderTaskgroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", taskgroupUnderTaskgroup.Name})

    taskgroupUnderTaskgroup.EntityData.YListKeys = []string {"Name"}

    return &(taskgroupUnderTaskgroup.EntityData)
}

// Aaa_Taskgroups_Taskgroup_Tasks
// Specify task IDs to be part of this group
type Aaa_Taskgroups_Taskgroup_Tasks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Task ID to be included. The type is slice of
    // Aaa_Taskgroups_Taskgroup_Tasks_Task.
    Task []*Aaa_Taskgroups_Taskgroup_Tasks_Task
}

func (tasks *Aaa_Taskgroups_Taskgroup_Tasks) GetEntityData() *types.CommonEntityData {
    tasks.EntityData.YFilter = tasks.YFilter
    tasks.EntityData.YangName = "tasks"
    tasks.EntityData.BundleName = "cisco_ios_xr"
    tasks.EntityData.ParentYangName = "taskgroup"
    tasks.EntityData.SegmentPath = "tasks"
    tasks.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:taskgroups/taskgroup/" + tasks.EntityData.SegmentPath
    tasks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tasks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tasks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tasks.EntityData.Children = types.NewOrderedMap()
    tasks.EntityData.Children.Append("task", types.YChild{"Task", nil})
    for i := range tasks.Task {
        tasks.EntityData.Children.Append(types.GetSegmentPath(tasks.Task[i]), types.YChild{"Task", tasks.Task[i]})
    }
    tasks.EntityData.Leafs = types.NewOrderedMap()

    tasks.EntityData.YListKeys = []string {}

    return &(tasks.EntityData)
}

// Aaa_Taskgroups_Taskgroup_Tasks_Task
// Task ID to be included
type Aaa_Taskgroups_Taskgroup_Tasks_Task struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This specifies the operation permitted for this
    // task eg: read/write/execute/debug. The type is AaaLocaldTaskClass.
    Type interface{}

    // This attribute is a key. Task ID to which permission is to be granted
    // (please use class AllTasks to get a list of valid task IDs). The type is
    // string.
    TaskId interface{}
}

func (task *Aaa_Taskgroups_Taskgroup_Tasks_Task) GetEntityData() *types.CommonEntityData {
    task.EntityData.YFilter = task.YFilter
    task.EntityData.YangName = "task"
    task.EntityData.BundleName = "cisco_ios_xr"
    task.EntityData.ParentYangName = "tasks"
    task.EntityData.SegmentPath = "task" + types.AddKeyToken(task.Type, "type") + types.AddKeyToken(task.TaskId, "task-id")
    task.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:taskgroups/taskgroup/tasks/" + task.EntityData.SegmentPath
    task.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    task.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    task.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    task.EntityData.Children = types.NewOrderedMap()
    task.EntityData.Leafs = types.NewOrderedMap()
    task.EntityData.Leafs.Append("type", types.YLeaf{"Type", task.Type})
    task.EntityData.Leafs.Append("task-id", types.YLeaf{"TaskId", task.TaskId})

    task.EntityData.YListKeys = []string {"Type", "TaskId"}

    return &(task.EntityData)
}

// Aaa_Usergroups
// Specify a Usergroup to inherit from
type Aaa_Usergroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Usergroup name. The type is slice of Aaa_Usergroups_Usergroup.
    Usergroup []*Aaa_Usergroups_Usergroup
}

func (usergroups *Aaa_Usergroups) GetEntityData() *types.CommonEntityData {
    usergroups.EntityData.YFilter = usergroups.YFilter
    usergroups.EntityData.YangName = "usergroups"
    usergroups.EntityData.BundleName = "cisco_ios_xr"
    usergroups.EntityData.ParentYangName = "aaa"
    usergroups.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-locald-cfg:usergroups"
    usergroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + usergroups.EntityData.SegmentPath
    usergroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroups.EntityData.Children = types.NewOrderedMap()
    usergroups.EntityData.Children.Append("usergroup", types.YChild{"Usergroup", nil})
    for i := range usergroups.Usergroup {
        usergroups.EntityData.Children.Append(types.GetSegmentPath(usergroups.Usergroup[i]), types.YChild{"Usergroup", usergroups.Usergroup[i]})
    }
    usergroups.EntityData.Leafs = types.NewOrderedMap()

    usergroups.EntityData.YListKeys = []string {}

    return &(usergroups.EntityData)
}

// Aaa_Usergroups_Usergroup
// Usergroup name
type Aaa_Usergroups_Usergroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Usergroup name. The type is string.
    Name interface{}

    // Description for the user group. The type is string.
    Description interface{}

    // Task group associated with this group.
    TaskgroupUnderUsergroups Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups

    // User group to be inherited by this group.
    UsergroupUnderUsergroups Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups
}

func (usergroup *Aaa_Usergroups_Usergroup) GetEntityData() *types.CommonEntityData {
    usergroup.EntityData.YFilter = usergroup.YFilter
    usergroup.EntityData.YangName = "usergroup"
    usergroup.EntityData.BundleName = "cisco_ios_xr"
    usergroup.EntityData.ParentYangName = "usergroups"
    usergroup.EntityData.SegmentPath = "usergroup" + types.AddKeyToken(usergroup.Name, "name")
    usergroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usergroups/" + usergroup.EntityData.SegmentPath
    usergroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroup.EntityData.Children = types.NewOrderedMap()
    usergroup.EntityData.Children.Append("taskgroup-under-usergroups", types.YChild{"TaskgroupUnderUsergroups", &usergroup.TaskgroupUnderUsergroups})
    usergroup.EntityData.Children.Append("usergroup-under-usergroups", types.YChild{"UsergroupUnderUsergroups", &usergroup.UsergroupUnderUsergroups})
    usergroup.EntityData.Leafs = types.NewOrderedMap()
    usergroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", usergroup.Name})
    usergroup.EntityData.Leafs.Append("description", types.YLeaf{"Description", usergroup.Description})

    usergroup.EntityData.YListKeys = []string {"Name"}

    return &(usergroup.EntityData)
}

// Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups
// Task group associated with this group
type Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the task group. The type is slice of
    // Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup.
    TaskgroupUnderUsergroup []*Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup
}

func (taskgroupUnderUsergroups *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups) GetEntityData() *types.CommonEntityData {
    taskgroupUnderUsergroups.EntityData.YFilter = taskgroupUnderUsergroups.YFilter
    taskgroupUnderUsergroups.EntityData.YangName = "taskgroup-under-usergroups"
    taskgroupUnderUsergroups.EntityData.BundleName = "cisco_ios_xr"
    taskgroupUnderUsergroups.EntityData.ParentYangName = "usergroup"
    taskgroupUnderUsergroups.EntityData.SegmentPath = "taskgroup-under-usergroups"
    taskgroupUnderUsergroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usergroups/usergroup/" + taskgroupUnderUsergroups.EntityData.SegmentPath
    taskgroupUnderUsergroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroupUnderUsergroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroupUnderUsergroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroupUnderUsergroups.EntityData.Children = types.NewOrderedMap()
    taskgroupUnderUsergroups.EntityData.Children.Append("taskgroup-under-usergroup", types.YChild{"TaskgroupUnderUsergroup", nil})
    for i := range taskgroupUnderUsergroups.TaskgroupUnderUsergroup {
        taskgroupUnderUsergroups.EntityData.Children.Append(types.GetSegmentPath(taskgroupUnderUsergroups.TaskgroupUnderUsergroup[i]), types.YChild{"TaskgroupUnderUsergroup", taskgroupUnderUsergroups.TaskgroupUnderUsergroup[i]})
    }
    taskgroupUnderUsergroups.EntityData.Leafs = types.NewOrderedMap()

    taskgroupUnderUsergroups.EntityData.YListKeys = []string {}

    return &(taskgroupUnderUsergroups.EntityData)
}

// Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup
// Name of the task group
type Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the task group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (taskgroupUnderUsergroup *Aaa_Usergroups_Usergroup_TaskgroupUnderUsergroups_TaskgroupUnderUsergroup) GetEntityData() *types.CommonEntityData {
    taskgroupUnderUsergroup.EntityData.YFilter = taskgroupUnderUsergroup.YFilter
    taskgroupUnderUsergroup.EntityData.YangName = "taskgroup-under-usergroup"
    taskgroupUnderUsergroup.EntityData.BundleName = "cisco_ios_xr"
    taskgroupUnderUsergroup.EntityData.ParentYangName = "taskgroup-under-usergroups"
    taskgroupUnderUsergroup.EntityData.SegmentPath = "taskgroup-under-usergroup" + types.AddKeyToken(taskgroupUnderUsergroup.Name, "name")
    taskgroupUnderUsergroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usergroups/usergroup/taskgroup-under-usergroups/" + taskgroupUnderUsergroup.EntityData.SegmentPath
    taskgroupUnderUsergroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    taskgroupUnderUsergroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    taskgroupUnderUsergroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    taskgroupUnderUsergroup.EntityData.Children = types.NewOrderedMap()
    taskgroupUnderUsergroup.EntityData.Leafs = types.NewOrderedMap()
    taskgroupUnderUsergroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", taskgroupUnderUsergroup.Name})

    taskgroupUnderUsergroup.EntityData.YListKeys = []string {"Name"}

    return &(taskgroupUnderUsergroup.EntityData)
}

// Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups
// User group to be inherited by this group
type Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the user group. The type is slice of
    // Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup.
    UsergroupUnderUsergroup []*Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup
}

func (usergroupUnderUsergroups *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsergroups.EntityData.YFilter = usergroupUnderUsergroups.YFilter
    usergroupUnderUsergroups.EntityData.YangName = "usergroup-under-usergroups"
    usergroupUnderUsergroups.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsergroups.EntityData.ParentYangName = "usergroup"
    usergroupUnderUsergroups.EntityData.SegmentPath = "usergroup-under-usergroups"
    usergroupUnderUsergroups.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usergroups/usergroup/" + usergroupUnderUsergroups.EntityData.SegmentPath
    usergroupUnderUsergroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsergroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsergroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsergroups.EntityData.Children = types.NewOrderedMap()
    usergroupUnderUsergroups.EntityData.Children.Append("usergroup-under-usergroup", types.YChild{"UsergroupUnderUsergroup", nil})
    for i := range usergroupUnderUsergroups.UsergroupUnderUsergroup {
        usergroupUnderUsergroups.EntityData.Children.Append(types.GetSegmentPath(usergroupUnderUsergroups.UsergroupUnderUsergroup[i]), types.YChild{"UsergroupUnderUsergroup", usergroupUnderUsergroups.UsergroupUnderUsergroup[i]})
    }
    usergroupUnderUsergroups.EntityData.Leafs = types.NewOrderedMap()

    usergroupUnderUsergroups.EntityData.YListKeys = []string {}

    return &(usergroupUnderUsergroups.EntityData)
}

// Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup
// Name of the user group
type Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the user group. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (usergroupUnderUsergroup *Aaa_Usergroups_Usergroup_UsergroupUnderUsergroups_UsergroupUnderUsergroup) GetEntityData() *types.CommonEntityData {
    usergroupUnderUsergroup.EntityData.YFilter = usergroupUnderUsergroup.YFilter
    usergroupUnderUsergroup.EntityData.YangName = "usergroup-under-usergroup"
    usergroupUnderUsergroup.EntityData.BundleName = "cisco_ios_xr"
    usergroupUnderUsergroup.EntityData.ParentYangName = "usergroup-under-usergroups"
    usergroupUnderUsergroup.EntityData.SegmentPath = "usergroup-under-usergroup" + types.AddKeyToken(usergroupUnderUsergroup.Name, "name")
    usergroupUnderUsergroup.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-locald-cfg:usergroups/usergroup/usergroup-under-usergroups/" + usergroupUnderUsergroup.EntityData.SegmentPath
    usergroupUnderUsergroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usergroupUnderUsergroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usergroupUnderUsergroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usergroupUnderUsergroup.EntityData.Children = types.NewOrderedMap()
    usergroupUnderUsergroup.EntityData.Leafs = types.NewOrderedMap()
    usergroupUnderUsergroup.EntityData.Leafs.Append("name", types.YLeaf{"Name", usergroupUnderUsergroup.Name})

    usergroupUnderUsergroup.EntityData.YListKeys = []string {"Name"}

    return &(usergroupUnderUsergroup.EntityData)
}

// Aaa_Tacacs
// Modify TACACS+ query parameters
type Aaa_Tacacs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Use a single connection for all sessions for a given TACACS+ server. The
    // type is bool. The default value is false.
    SingleConnect interface{}

    // IPv6 configuration.
    Ipv6 Aaa_Tacacs_Ipv6

    // Specify a TACACS+ server.
    Hosts Aaa_Tacacs_Hosts

    // IPv4 configuration.
    Ipv4 Aaa_Tacacs_Ipv4

    // List of VRFs.
    Vrfs Aaa_Tacacs_Vrfs
}

func (tacacs *Aaa_Tacacs) GetEntityData() *types.CommonEntityData {
    tacacs.EntityData.YFilter = tacacs.YFilter
    tacacs.EntityData.YangName = "tacacs"
    tacacs.EntityData.BundleName = "cisco_ios_xr"
    tacacs.EntityData.ParentYangName = "aaa"
    tacacs.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-tacacs-cfg:tacacs"
    tacacs.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + tacacs.EntityData.SegmentPath
    tacacs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tacacs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tacacs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tacacs.EntityData.Children = types.NewOrderedMap()
    tacacs.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &tacacs.Ipv6})
    tacacs.EntityData.Children.Append("hosts", types.YChild{"Hosts", &tacacs.Hosts})
    tacacs.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &tacacs.Ipv4})
    tacacs.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &tacacs.Vrfs})
    tacacs.EntityData.Leafs = types.NewOrderedMap()
    tacacs.EntityData.Leafs.Append("key", types.YLeaf{"Key", tacacs.Key})
    tacacs.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", tacacs.Timeout})
    tacacs.EntityData.Leafs.Append("single-connect", types.YLeaf{"SingleConnect", tacacs.SingleConnect})

    tacacs.EntityData.YListKeys = []string {}

    return &(tacacs.EntityData)
}

// Aaa_Tacacs_Ipv6
// IPv6 configuration
type Aaa_Tacacs_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // TacacsDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv6 *Aaa_Tacacs_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "tacacs"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ipv6.Dscp})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Aaa_Tacacs_Hosts
// Specify a TACACS+ server
type Aaa_Tacacs_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // One of the TACACS+ servers. The type is slice of Aaa_Tacacs_Hosts_Host.
    Host []*Aaa_Tacacs_Hosts_Host
}

func (hosts *Aaa_Tacacs_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "tacacs"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/" + hosts.EntityData.SegmentPath
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Aaa_Tacacs_Hosts_Host
// One of the TACACS+ servers
type Aaa_Tacacs_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of TACACS+ server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port number (standard 49). The type is interface{}
    // with range: 1..65535.
    PortNumber interface{}

    // Set TACACS+ encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a TACACS+ server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Use a single connection for all sessions for a given TACACS+ server. The
    // type is bool. The default value is false.
    SingleConnect interface{}
}

func (host *Aaa_Tacacs_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.OrderingIndex, "ordering-index") + types.AddKeyToken(host.IpAddress, "ip-address") + types.AddKeyToken(host.PortNumber, "port-number")
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/hosts/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", host.OrderingIndex})
    host.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", host.IpAddress})
    host.EntityData.Leafs.Append("port-number", types.YLeaf{"PortNumber", host.PortNumber})
    host.EntityData.Leafs.Append("key", types.YLeaf{"Key", host.Key})
    host.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", host.Timeout})
    host.EntityData.Leafs.Append("single-connect", types.YLeaf{"SingleConnect", host.SingleConnect})

    host.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress", "PortNumber"}

    return &(host.EntityData)
}

// Aaa_Tacacs_Ipv4
// IPv4 configuration
type Aaa_Tacacs_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // TacacsDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Aaa_Tacacs_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "tacacs"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ipv4.Dscp})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Aaa_Tacacs_Vrfs
// List of VRFs
type Aaa_Tacacs_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRF. The type is slice of Aaa_Tacacs_Vrfs_Vrf.
    Vrf []*Aaa_Tacacs_Vrfs_Vrf
}

func (vrfs *Aaa_Tacacs_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "tacacs"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Aaa_Tacacs_Vrfs_Vrf
// A VRF
type Aaa_Tacacs_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. Specify 'default' for default VRF. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in TACACS+ packets. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}
}

func (vrf *Aaa_Tacacs_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-tacacs-cfg:tacacs/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", vrf.SourceInterface})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Aaa_AaaSubscriber
// AAA subscriber
type Aaa_AaaSubscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA authorization policy.
    PolicyIfAuthors Aaa_AaaSubscriber_PolicyIfAuthors

    // AAA accounting.
    Accountings Aaa_AaaSubscriber_Accountings

    // Set accounting parameters for Service.
    ServiceAccounting Aaa_AaaSubscriber_ServiceAccounting

    // AAA authorization prepaid.
    PrepaidAuthors Aaa_AaaSubscriber_PrepaidAuthors

    // AAA authorization.
    Authorizations Aaa_AaaSubscriber_Authorizations

    // AAA authentication.
    Authentications Aaa_AaaSubscriber_Authentications
}

func (aaaSubscriber *Aaa_AaaSubscriber) GetEntityData() *types.CommonEntityData {
    aaaSubscriber.EntityData.YFilter = aaaSubscriber.YFilter
    aaaSubscriber.EntityData.YangName = "aaa-subscriber"
    aaaSubscriber.EntityData.BundleName = "cisco_ios_xr"
    aaaSubscriber.EntityData.ParentYangName = "aaa"
    aaaSubscriber.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber"
    aaaSubscriber.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + aaaSubscriber.EntityData.SegmentPath
    aaaSubscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaSubscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaSubscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaSubscriber.EntityData.Children = types.NewOrderedMap()
    aaaSubscriber.EntityData.Children.Append("policy-if-authors", types.YChild{"PolicyIfAuthors", &aaaSubscriber.PolicyIfAuthors})
    aaaSubscriber.EntityData.Children.Append("accountings", types.YChild{"Accountings", &aaaSubscriber.Accountings})
    aaaSubscriber.EntityData.Children.Append("service-accounting", types.YChild{"ServiceAccounting", &aaaSubscriber.ServiceAccounting})
    aaaSubscriber.EntityData.Children.Append("prepaid-authors", types.YChild{"PrepaidAuthors", &aaaSubscriber.PrepaidAuthors})
    aaaSubscriber.EntityData.Children.Append("authorizations", types.YChild{"Authorizations", &aaaSubscriber.Authorizations})
    aaaSubscriber.EntityData.Children.Append("authentications", types.YChild{"Authentications", &aaaSubscriber.Authentications})
    aaaSubscriber.EntityData.Leafs = types.NewOrderedMap()

    aaaSubscriber.EntityData.YListKeys = []string {}

    return &(aaaSubscriber.EntityData)
}

// Aaa_AaaSubscriber_PolicyIfAuthors
// AAA authorization policy
type Aaa_AaaSubscriber_PolicyIfAuthors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor.
    PolicyIfAuthor []*Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor
}

func (policyIfAuthors *Aaa_AaaSubscriber_PolicyIfAuthors) GetEntityData() *types.CommonEntityData {
    policyIfAuthors.EntityData.YFilter = policyIfAuthors.YFilter
    policyIfAuthors.EntityData.YangName = "policy-if-authors"
    policyIfAuthors.EntityData.BundleName = "cisco_ios_xr"
    policyIfAuthors.EntityData.ParentYangName = "aaa-subscriber"
    policyIfAuthors.EntityData.SegmentPath = "policy-if-authors"
    policyIfAuthors.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + policyIfAuthors.EntityData.SegmentPath
    policyIfAuthors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIfAuthors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIfAuthors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIfAuthors.EntityData.Children = types.NewOrderedMap()
    policyIfAuthors.EntityData.Children.Append("policy-if-author", types.YChild{"PolicyIfAuthor", nil})
    for i := range policyIfAuthors.PolicyIfAuthor {
        policyIfAuthors.EntityData.Children.Append(types.GetSegmentPath(policyIfAuthors.PolicyIfAuthor[i]), types.YChild{"PolicyIfAuthor", policyIfAuthors.PolicyIfAuthor[i]})
    }
    policyIfAuthors.EntityData.Leafs = types.NewOrderedMap()

    policyIfAuthors.EntityData.YListKeys = []string {}

    return &(policyIfAuthors.EntityData)
}

// Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor
// Configurations related to authorization
type Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (policyIfAuthor *Aaa_AaaSubscriber_PolicyIfAuthors_PolicyIfAuthor) GetEntityData() *types.CommonEntityData {
    policyIfAuthor.EntityData.YFilter = policyIfAuthor.YFilter
    policyIfAuthor.EntityData.YangName = "policy-if-author"
    policyIfAuthor.EntityData.BundleName = "cisco_ios_xr"
    policyIfAuthor.EntityData.ParentYangName = "policy-if-authors"
    policyIfAuthor.EntityData.SegmentPath = "policy-if-author" + types.AddKeyToken(policyIfAuthor.Type, "type") + types.AddKeyToken(policyIfAuthor.Listname, "listname")
    policyIfAuthor.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/policy-if-authors/" + policyIfAuthor.EntityData.SegmentPath
    policyIfAuthor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyIfAuthor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyIfAuthor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyIfAuthor.EntityData.Children = types.NewOrderedMap()
    policyIfAuthor.EntityData.Leafs = types.NewOrderedMap()
    policyIfAuthor.EntityData.Leafs.Append("type", types.YLeaf{"Type", policyIfAuthor.Type})
    policyIfAuthor.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", policyIfAuthor.Listname})
    policyIfAuthor.EntityData.Leafs.Append("method", types.YLeaf{"Method", policyIfAuthor.Method})
    policyIfAuthor.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", policyIfAuthor.ServerGroupName})

    policyIfAuthor.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(policyIfAuthor.EntityData)
}

// Aaa_AaaSubscriber_Accountings
// AAA accounting
type Aaa_AaaSubscriber_Accountings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_AaaSubscriber_Accountings_Accounting.
    Accounting []*Aaa_AaaSubscriber_Accountings_Accounting
}

func (accountings *Aaa_AaaSubscriber_Accountings) GetEntityData() *types.CommonEntityData {
    accountings.EntityData.YFilter = accountings.YFilter
    accountings.EntityData.YangName = "accountings"
    accountings.EntityData.BundleName = "cisco_ios_xr"
    accountings.EntityData.ParentYangName = "aaa-subscriber"
    accountings.EntityData.SegmentPath = "accountings"
    accountings.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + accountings.EntityData.SegmentPath
    accountings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountings.EntityData.Children = types.NewOrderedMap()
    accountings.EntityData.Children.Append("accounting", types.YChild{"Accounting", nil})
    for i := range accountings.Accounting {
        accountings.EntityData.Children.Append(types.GetSegmentPath(accountings.Accounting[i]), types.YChild{"Accounting", accountings.Accounting[i]})
    }
    accountings.EntityData.Leafs = types.NewOrderedMap()

    accountings.EntityData.YListKeys = []string {}

    return &(accountings.EntityData)
}

// Aaa_AaaSubscriber_Accountings_Accounting
// Configurations related to accounting
type Aaa_AaaSubscriber_Accountings_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set accounting lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Broadcast. The type is AaaAccountingBroadcast. This attribute is mandatory.
    Broadcast interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (accounting *Aaa_AaaSubscriber_Accountings_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "accountings"
    accounting.EntityData.SegmentPath = "accounting" + types.AddKeyToken(accounting.Type, "type") + types.AddKeyToken(accounting.Listname, "listname")
    accounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/accountings/" + accounting.EntityData.SegmentPath
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = types.NewOrderedMap()
    accounting.EntityData.Leafs = types.NewOrderedMap()
    accounting.EntityData.Leafs.Append("type", types.YLeaf{"Type", accounting.Type})
    accounting.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", accounting.Listname})
    accounting.EntityData.Leafs.Append("broadcast", types.YLeaf{"Broadcast", accounting.Broadcast})
    accounting.EntityData.Leafs.Append("method", types.YLeaf{"Method", accounting.Method})
    accounting.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", accounting.ServerGroupName})

    accounting.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(accounting.EntityData)
}

// Aaa_AaaSubscriber_ServiceAccounting
// Set accounting parameters for Service
type Aaa_AaaSubscriber_ServiceAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Send extended/brief service accounting records. The type is
    // AaaServiceAccounting.
    Type interface{}
}

func (serviceAccounting *Aaa_AaaSubscriber_ServiceAccounting) GetEntityData() *types.CommonEntityData {
    serviceAccounting.EntityData.YFilter = serviceAccounting.YFilter
    serviceAccounting.EntityData.YangName = "service-accounting"
    serviceAccounting.EntityData.BundleName = "cisco_ios_xr"
    serviceAccounting.EntityData.ParentYangName = "aaa-subscriber"
    serviceAccounting.EntityData.SegmentPath = "service-accounting"
    serviceAccounting.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + serviceAccounting.EntityData.SegmentPath
    serviceAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceAccounting.EntityData.Children = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs = types.NewOrderedMap()
    serviceAccounting.EntityData.Leafs.Append("type", types.YLeaf{"Type", serviceAccounting.Type})

    serviceAccounting.EntityData.YListKeys = []string {}

    return &(serviceAccounting.EntityData)
}

// Aaa_AaaSubscriber_PrepaidAuthors
// AAA authorization prepaid
type Aaa_AaaSubscriber_PrepaidAuthors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor.
    PrepaidAuthor []*Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor
}

func (prepaidAuthors *Aaa_AaaSubscriber_PrepaidAuthors) GetEntityData() *types.CommonEntityData {
    prepaidAuthors.EntityData.YFilter = prepaidAuthors.YFilter
    prepaidAuthors.EntityData.YangName = "prepaid-authors"
    prepaidAuthors.EntityData.BundleName = "cisco_ios_xr"
    prepaidAuthors.EntityData.ParentYangName = "aaa-subscriber"
    prepaidAuthors.EntityData.SegmentPath = "prepaid-authors"
    prepaidAuthors.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + prepaidAuthors.EntityData.SegmentPath
    prepaidAuthors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepaidAuthors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepaidAuthors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepaidAuthors.EntityData.Children = types.NewOrderedMap()
    prepaidAuthors.EntityData.Children.Append("prepaid-author", types.YChild{"PrepaidAuthor", nil})
    for i := range prepaidAuthors.PrepaidAuthor {
        prepaidAuthors.EntityData.Children.Append(types.GetSegmentPath(prepaidAuthors.PrepaidAuthor[i]), types.YChild{"PrepaidAuthor", prepaidAuthors.PrepaidAuthor[i]})
    }
    prepaidAuthors.EntityData.Leafs = types.NewOrderedMap()

    prepaidAuthors.EntityData.YListKeys = []string {}

    return &(prepaidAuthors.EntityData)
}

// Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor
// Configurations related to authorization
type Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (prepaidAuthor *Aaa_AaaSubscriber_PrepaidAuthors_PrepaidAuthor) GetEntityData() *types.CommonEntityData {
    prepaidAuthor.EntityData.YFilter = prepaidAuthor.YFilter
    prepaidAuthor.EntityData.YangName = "prepaid-author"
    prepaidAuthor.EntityData.BundleName = "cisco_ios_xr"
    prepaidAuthor.EntityData.ParentYangName = "prepaid-authors"
    prepaidAuthor.EntityData.SegmentPath = "prepaid-author" + types.AddKeyToken(prepaidAuthor.Type, "type") + types.AddKeyToken(prepaidAuthor.Listname, "listname")
    prepaidAuthor.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/prepaid-authors/" + prepaidAuthor.EntityData.SegmentPath
    prepaidAuthor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepaidAuthor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepaidAuthor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepaidAuthor.EntityData.Children = types.NewOrderedMap()
    prepaidAuthor.EntityData.Leafs = types.NewOrderedMap()
    prepaidAuthor.EntityData.Leafs.Append("type", types.YLeaf{"Type", prepaidAuthor.Type})
    prepaidAuthor.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", prepaidAuthor.Listname})
    prepaidAuthor.EntityData.Leafs.Append("method", types.YLeaf{"Method", prepaidAuthor.Method})
    prepaidAuthor.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", prepaidAuthor.ServerGroupName})

    prepaidAuthor.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(prepaidAuthor.EntityData)
}

// Aaa_AaaSubscriber_Authorizations
// AAA authorization
type Aaa_AaaSubscriber_Authorizations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authorization. The type is slice of
    // Aaa_AaaSubscriber_Authorizations_Authorization.
    Authorization []*Aaa_AaaSubscriber_Authorizations_Authorization
}

func (authorizations *Aaa_AaaSubscriber_Authorizations) GetEntityData() *types.CommonEntityData {
    authorizations.EntityData.YFilter = authorizations.YFilter
    authorizations.EntityData.YangName = "authorizations"
    authorizations.EntityData.BundleName = "cisco_ios_xr"
    authorizations.EntityData.ParentYangName = "aaa-subscriber"
    authorizations.EntityData.SegmentPath = "authorizations"
    authorizations.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + authorizations.EntityData.SegmentPath
    authorizations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizations.EntityData.Children = types.NewOrderedMap()
    authorizations.EntityData.Children.Append("authorization", types.YChild{"Authorization", nil})
    for i := range authorizations.Authorization {
        authorizations.EntityData.Children.Append(types.GetSegmentPath(authorizations.Authorization[i]), types.YChild{"Authorization", authorizations.Authorization[i]})
    }
    authorizations.EntityData.Leafs = types.NewOrderedMap()

    authorizations.EntityData.YListKeys = []string {}

    return &(authorizations.EntityData)
}

// Aaa_AaaSubscriber_Authorizations_Authorization
// Configurations related to authorization
type Aaa_AaaSubscriber_Authorizations_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set authorization lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authorization list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authorization *Aaa_AaaSubscriber_Authorizations_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "authorizations"
    authorization.EntityData.SegmentPath = "authorization" + types.AddKeyToken(authorization.Type, "type") + types.AddKeyToken(authorization.Listname, "listname")
    authorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/authorizations/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("type", types.YLeaf{"Type", authorization.Type})
    authorization.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", authorization.Listname})
    authorization.EntityData.Leafs.Append("method", types.YLeaf{"Method", authorization.Method})
    authorization.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", authorization.ServerGroupName})

    authorization.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(authorization.EntityData)
}

// Aaa_AaaSubscriber_Authentications
// AAA authentication
type Aaa_AaaSubscriber_Authentications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_AaaSubscriber_Authentications_Authentication.
    Authentication []*Aaa_AaaSubscriber_Authentications_Authentication
}

func (authentications *Aaa_AaaSubscriber_Authentications) GetEntityData() *types.CommonEntityData {
    authentications.EntityData.YFilter = authentications.YFilter
    authentications.EntityData.YangName = "authentications"
    authentications.EntityData.BundleName = "cisco_ios_xr"
    authentications.EntityData.ParentYangName = "aaa-subscriber"
    authentications.EntityData.SegmentPath = "authentications"
    authentications.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/" + authentications.EntityData.SegmentPath
    authentications.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentications.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentications.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentications.EntityData.Children = types.NewOrderedMap()
    authentications.EntityData.Children.Append("authentication", types.YChild{"Authentication", nil})
    for i := range authentications.Authentication {
        authentications.EntityData.Children.Append(types.GetSegmentPath(authentications.Authentication[i]), types.YChild{"Authentication", authentications.Authentication[i]})
    }
    authentications.EntityData.Leafs = types.NewOrderedMap()

    authentications.EntityData.YListKeys = []string {}

    return &(authentications.EntityData)
}

// Aaa_AaaSubscriber_Authentications_Authentication
// Configurations related to authentication
type Aaa_AaaSubscriber_Authentications_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set authentication lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authentication list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authentication *Aaa_AaaSubscriber_Authentications_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "authentications"
    authentication.EntityData.SegmentPath = "authentication" + types.AddKeyToken(authentication.Type, "type") + types.AddKeyToken(authentication.Listname, "listname")
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-subscriber/authentications/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("type", types.YLeaf{"Type", authentication.Type})
    authentication.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", authentication.Listname})
    authentication.EntityData.Leafs.Append("method", types.YLeaf{"Method", authentication.Method})
    authentication.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", authentication.ServerGroupName})

    authentication.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(authentication.EntityData)
}

// Aaa_AaaMobile
// AAA Mobile
type Aaa_AaaMobile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA Mobile Accounting.
    Mobiles Aaa_AaaMobile_Mobiles
}

func (aaaMobile *Aaa_AaaMobile) GetEntityData() *types.CommonEntityData {
    aaaMobile.EntityData.YFilter = aaaMobile.YFilter
    aaaMobile.EntityData.YangName = "aaa-mobile"
    aaaMobile.EntityData.BundleName = "cisco_ios_xr"
    aaaMobile.EntityData.ParentYangName = "aaa"
    aaaMobile.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile"
    aaaMobile.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + aaaMobile.EntityData.SegmentPath
    aaaMobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaMobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaMobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaMobile.EntityData.Children = types.NewOrderedMap()
    aaaMobile.EntityData.Children.Append("mobiles", types.YChild{"Mobiles", &aaaMobile.Mobiles})
    aaaMobile.EntityData.Leafs = types.NewOrderedMap()

    aaaMobile.EntityData.YListKeys = []string {}

    return &(aaaMobile.EntityData)
}

// Aaa_AaaMobile_Mobiles
// AAA Mobile Accounting
type Aaa_AaaMobile_Mobiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to accounting. The type is slice of
    // Aaa_AaaMobile_Mobiles_Mobile.
    Mobile []*Aaa_AaaMobile_Mobiles_Mobile
}

func (mobiles *Aaa_AaaMobile_Mobiles) GetEntityData() *types.CommonEntityData {
    mobiles.EntityData.YFilter = mobiles.YFilter
    mobiles.EntityData.YangName = "mobiles"
    mobiles.EntityData.BundleName = "cisco_ios_xr"
    mobiles.EntityData.ParentYangName = "aaa-mobile"
    mobiles.EntityData.SegmentPath = "mobiles"
    mobiles.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile/" + mobiles.EntityData.SegmentPath
    mobiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobiles.EntityData.Children = types.NewOrderedMap()
    mobiles.EntityData.Children.Append("mobile", types.YChild{"Mobile", nil})
    for i := range mobiles.Mobile {
        mobiles.EntityData.Children.Append(types.GetSegmentPath(mobiles.Mobile[i]), types.YChild{"Mobile", mobiles.Mobile[i]})
    }
    mobiles.EntityData.Leafs = types.NewOrderedMap()

    mobiles.EntityData.YListKeys = []string {}

    return &(mobiles.EntityData)
}

// Aaa_AaaMobile_Mobiles_Mobile
// Configurations related to accounting
type Aaa_AaaMobile_Mobiles_Mobile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Named accounting list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Broadcast. The type is AaaAccountingBroadcast. This attribute is mandatory.
    Broadcast interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (mobile *Aaa_AaaMobile_Mobiles_Mobile) GetEntityData() *types.CommonEntityData {
    mobile.EntityData.YFilter = mobile.YFilter
    mobile.EntityData.YangName = "mobile"
    mobile.EntityData.BundleName = "cisco_ios_xr"
    mobile.EntityData.ParentYangName = "mobiles"
    mobile.EntityData.SegmentPath = "mobile" + types.AddKeyToken(mobile.Listname, "listname")
    mobile.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-mobile/mobiles/" + mobile.EntityData.SegmentPath
    mobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobile.EntityData.Children = types.NewOrderedMap()
    mobile.EntityData.Leafs = types.NewOrderedMap()
    mobile.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", mobile.Listname})
    mobile.EntityData.Leafs.Append("broadcast", types.YLeaf{"Broadcast", mobile.Broadcast})
    mobile.EntityData.Leafs.Append("method", types.YLeaf{"Method", mobile.Method})
    mobile.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", mobile.ServerGroupName})

    mobile.EntityData.YListKeys = []string {"Listname"}

    return &(mobile.EntityData)
}

// Aaa_AaaDot1x
// AAA Dot1x
type Aaa_AaaDot1x struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA authentication.
    Authentications Aaa_AaaDot1x_Authentications
}

func (aaaDot1x *Aaa_AaaDot1x) GetEntityData() *types.CommonEntityData {
    aaaDot1x.EntityData.YFilter = aaaDot1x.YFilter
    aaaDot1x.EntityData.YangName = "aaa-dot1x"
    aaaDot1x.EntityData.BundleName = "cisco_ios_xr"
    aaaDot1x.EntityData.ParentYangName = "aaa"
    aaaDot1x.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x"
    aaaDot1x.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + aaaDot1x.EntityData.SegmentPath
    aaaDot1x.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaDot1x.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaDot1x.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaDot1x.EntityData.Children = types.NewOrderedMap()
    aaaDot1x.EntityData.Children.Append("authentications", types.YChild{"Authentications", &aaaDot1x.Authentications})
    aaaDot1x.EntityData.Leafs = types.NewOrderedMap()

    aaaDot1x.EntityData.YListKeys = []string {}

    return &(aaaDot1x.EntityData)
}

// Aaa_AaaDot1x_Authentications
// AAA authentication
type Aaa_AaaDot1x_Authentications struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configurations related to authentication. The type is slice of
    // Aaa_AaaDot1x_Authentications_Authentication.
    Authentication []*Aaa_AaaDot1x_Authentications_Authentication
}

func (authentications *Aaa_AaaDot1x_Authentications) GetEntityData() *types.CommonEntityData {
    authentications.EntityData.YFilter = authentications.YFilter
    authentications.EntityData.YangName = "authentications"
    authentications.EntityData.BundleName = "cisco_ios_xr"
    authentications.EntityData.ParentYangName = "aaa-dot1x"
    authentications.EntityData.SegmentPath = "authentications"
    authentications.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x/" + authentications.EntityData.SegmentPath
    authentications.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentications.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentications.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentications.EntityData.Children = types.NewOrderedMap()
    authentications.EntityData.Children.Append("authentication", types.YChild{"Authentication", nil})
    for i := range authentications.Authentication {
        authentications.EntityData.Children.Append(types.GetSegmentPath(authentications.Authentication[i]), types.YChild{"Authentication", authentications.Authentication[i]})
    }
    authentications.EntityData.Leafs = types.NewOrderedMap()

    authentications.EntityData.YListKeys = []string {}

    return &(authentications.EntityData)
}

// Aaa_AaaDot1x_Authentications_Authentication
// Configurations related to authentication
type Aaa_AaaDot1x_Authentications_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Set authentication lists. The type is string with
    // pattern: (subscriber)|(service)|(policy-if)|(prepaid)|(dot1x).
    Type interface{}

    // This attribute is a key. Named authentication list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Listname interface{}

    // Method Types. The type is slice of AaaMethod.
    Method []interface{}

    // Server group names. The type is slice of string.
    ServerGroupName []interface{}
}

func (authentication *Aaa_AaaDot1x_Authentications_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "authentications"
    authentication.EntityData.SegmentPath = "authentication" + types.AddKeyToken(authentication.Type, "type") + types.AddKeyToken(authentication.Listname, "listname")
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:aaa-dot1x/authentications/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("type", types.YLeaf{"Type", authentication.Type})
    authentication.EntityData.Leafs.Append("listname", types.YLeaf{"Listname", authentication.Listname})
    authentication.EntityData.Leafs.Append("method", types.YLeaf{"Method", authentication.Method})
    authentication.EntityData.Leafs.Append("server-group-name", types.YLeaf{"ServerGroupName", authentication.ServerGroupName})

    authentication.EntityData.YListKeys = []string {"Type", "Listname"}

    return &(authentication.EntityData)
}

// Aaa_RadiusAttribute
// AAA RADIUS attribute configurations
type Aaa_RadiusAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute.
    NasPortId Aaa_RadiusAttribute_NasPortId

    // AAA calling station id attribute.
    CallingStation Aaa_RadiusAttribute_CallingStation

    // AAA called station id attribute.
    CalledStation Aaa_RadiusAttribute_CalledStation

    // AAA nas-port-id attribute.
    NasPort Aaa_RadiusAttribute_NasPort

    // AAA nas-port-id attribute format.
    FormatOthers Aaa_RadiusAttribute_FormatOthers
}

func (radiusAttribute *Aaa_RadiusAttribute) GetEntityData() *types.CommonEntityData {
    radiusAttribute.EntityData.YFilter = radiusAttribute.YFilter
    radiusAttribute.EntityData.YangName = "radius-attribute"
    radiusAttribute.EntityData.BundleName = "cisco_ios_xr"
    radiusAttribute.EntityData.ParentYangName = "aaa"
    radiusAttribute.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute"
    radiusAttribute.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + radiusAttribute.EntityData.SegmentPath
    radiusAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radiusAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radiusAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radiusAttribute.EntityData.Children = types.NewOrderedMap()
    radiusAttribute.EntityData.Children.Append("nas-port-id", types.YChild{"NasPortId", &radiusAttribute.NasPortId})
    radiusAttribute.EntityData.Children.Append("calling-station", types.YChild{"CallingStation", &radiusAttribute.CallingStation})
    radiusAttribute.EntityData.Children.Append("called-station", types.YChild{"CalledStation", &radiusAttribute.CalledStation})
    radiusAttribute.EntityData.Children.Append("nas-port", types.YChild{"NasPort", &radiusAttribute.NasPort})
    radiusAttribute.EntityData.Children.Append("format-others", types.YChild{"FormatOthers", &radiusAttribute.FormatOthers})
    radiusAttribute.EntityData.Leafs = types.NewOrderedMap()

    radiusAttribute.EntityData.YListKeys = []string {}

    return &(radiusAttribute.EntityData)
}

// Aaa_RadiusAttribute_NasPortId
// AAA nas-port-id attribute
type Aaa_RadiusAttribute_NasPortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_NasPortId_Formats
}

func (nasPortId *Aaa_RadiusAttribute_NasPortId) GetEntityData() *types.CommonEntityData {
    nasPortId.EntityData.YFilter = nasPortId.YFilter
    nasPortId.EntityData.YangName = "nas-port-id"
    nasPortId.EntityData.BundleName = "cisco_ios_xr"
    nasPortId.EntityData.ParentYangName = "radius-attribute"
    nasPortId.EntityData.SegmentPath = "nas-port-id"
    nasPortId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/" + nasPortId.EntityData.SegmentPath
    nasPortId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nasPortId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nasPortId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nasPortId.EntityData.Children = types.NewOrderedMap()
    nasPortId.EntityData.Children.Append("formats", types.YChild{"Formats", &nasPortId.Formats})
    nasPortId.EntityData.Leafs = types.NewOrderedMap()

    nasPortId.EntityData.YListKeys = []string {}

    return &(nasPortId.EntityData)
}

// Aaa_RadiusAttribute_NasPortId_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_NasPortId_Formats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_NasPortId_Formats_Format.
    Format []*Aaa_RadiusAttribute_NasPortId_Formats_Format
}

func (formats *Aaa_RadiusAttribute_NasPortId_Formats) GetEntityData() *types.CommonEntityData {
    formats.EntityData.YFilter = formats.YFilter
    formats.EntityData.YangName = "formats"
    formats.EntityData.BundleName = "cisco_ios_xr"
    formats.EntityData.ParentYangName = "nas-port-id"
    formats.EntityData.SegmentPath = "formats"
    formats.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/nas-port-id/" + formats.EntityData.SegmentPath
    formats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formats.EntityData.Children = types.NewOrderedMap()
    formats.EntityData.Children.Append("format", types.YChild{"Format", nil})
    for i := range formats.Format {
        formats.EntityData.Children.Append(types.GetSegmentPath(formats.Format[i]), types.YChild{"Format", formats.Format[i]})
    }
    formats.EntityData.Leafs = types.NewOrderedMap()

    formats.EntityData.YListKeys = []string {}

    return &(formats.EntityData)
}

// Aaa_RadiusAttribute_NasPortId_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_NasPortId_Formats_Format struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_NasPortId_Formats_Format) GetEntityData() *types.CommonEntityData {
    format.EntityData.YFilter = format.YFilter
    format.EntityData.YangName = "format"
    format.EntityData.BundleName = "cisco_ios_xr"
    format.EntityData.ParentYangName = "formats"
    format.EntityData.SegmentPath = "format" + types.AddKeyToken(format.Type, "type")
    format.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/nas-port-id/formats/" + format.EntityData.SegmentPath
    format.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    format.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    format.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    format.EntityData.Children = types.NewOrderedMap()
    format.EntityData.Leafs = types.NewOrderedMap()
    format.EntityData.Leafs.Append("type", types.YLeaf{"Type", format.Type})
    format.EntityData.Leafs.Append("format-name", types.YLeaf{"FormatName", format.FormatName})

    format.EntityData.YListKeys = []string {"Type"}

    return &(format.EntityData)
}

// Aaa_RadiusAttribute_CallingStation
// AAA calling station id attribute
type Aaa_RadiusAttribute_CallingStation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_CallingStation_Formats
}

func (callingStation *Aaa_RadiusAttribute_CallingStation) GetEntityData() *types.CommonEntityData {
    callingStation.EntityData.YFilter = callingStation.YFilter
    callingStation.EntityData.YangName = "calling-station"
    callingStation.EntityData.BundleName = "cisco_ios_xr"
    callingStation.EntityData.ParentYangName = "radius-attribute"
    callingStation.EntityData.SegmentPath = "calling-station"
    callingStation.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/" + callingStation.EntityData.SegmentPath
    callingStation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    callingStation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    callingStation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    callingStation.EntityData.Children = types.NewOrderedMap()
    callingStation.EntityData.Children.Append("formats", types.YChild{"Formats", &callingStation.Formats})
    callingStation.EntityData.Leafs = types.NewOrderedMap()

    callingStation.EntityData.YListKeys = []string {}

    return &(callingStation.EntityData)
}

// Aaa_RadiusAttribute_CallingStation_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_CallingStation_Formats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_CallingStation_Formats_Format.
    Format []*Aaa_RadiusAttribute_CallingStation_Formats_Format
}

func (formats *Aaa_RadiusAttribute_CallingStation_Formats) GetEntityData() *types.CommonEntityData {
    formats.EntityData.YFilter = formats.YFilter
    formats.EntityData.YangName = "formats"
    formats.EntityData.BundleName = "cisco_ios_xr"
    formats.EntityData.ParentYangName = "calling-station"
    formats.EntityData.SegmentPath = "formats"
    formats.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/calling-station/" + formats.EntityData.SegmentPath
    formats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formats.EntityData.Children = types.NewOrderedMap()
    formats.EntityData.Children.Append("format", types.YChild{"Format", nil})
    for i := range formats.Format {
        formats.EntityData.Children.Append(types.GetSegmentPath(formats.Format[i]), types.YChild{"Format", formats.Format[i]})
    }
    formats.EntityData.Leafs = types.NewOrderedMap()

    formats.EntityData.YListKeys = []string {}

    return &(formats.EntityData)
}

// Aaa_RadiusAttribute_CallingStation_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_CallingStation_Formats_Format struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_CallingStation_Formats_Format) GetEntityData() *types.CommonEntityData {
    format.EntityData.YFilter = format.YFilter
    format.EntityData.YangName = "format"
    format.EntityData.BundleName = "cisco_ios_xr"
    format.EntityData.ParentYangName = "formats"
    format.EntityData.SegmentPath = "format" + types.AddKeyToken(format.Type, "type")
    format.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/calling-station/formats/" + format.EntityData.SegmentPath
    format.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    format.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    format.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    format.EntityData.Children = types.NewOrderedMap()
    format.EntityData.Leafs = types.NewOrderedMap()
    format.EntityData.Leafs.Append("type", types.YLeaf{"Type", format.Type})
    format.EntityData.Leafs.Append("format-name", types.YLeaf{"FormatName", format.FormatName})

    format.EntityData.YListKeys = []string {"Type"}

    return &(format.EntityData)
}

// Aaa_RadiusAttribute_CalledStation
// AAA called station id attribute
type Aaa_RadiusAttribute_CalledStation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    Formats Aaa_RadiusAttribute_CalledStation_Formats
}

func (calledStation *Aaa_RadiusAttribute_CalledStation) GetEntityData() *types.CommonEntityData {
    calledStation.EntityData.YFilter = calledStation.YFilter
    calledStation.EntityData.YangName = "called-station"
    calledStation.EntityData.BundleName = "cisco_ios_xr"
    calledStation.EntityData.ParentYangName = "radius-attribute"
    calledStation.EntityData.SegmentPath = "called-station"
    calledStation.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/" + calledStation.EntityData.SegmentPath
    calledStation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    calledStation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    calledStation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    calledStation.EntityData.Children = types.NewOrderedMap()
    calledStation.EntityData.Children.Append("formats", types.YChild{"Formats", &calledStation.Formats})
    calledStation.EntityData.Leafs = types.NewOrderedMap()

    calledStation.EntityData.YListKeys = []string {}

    return &(calledStation.EntityData)
}

// Aaa_RadiusAttribute_CalledStation_Formats
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_CalledStation_Formats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nas-port-id attribute format. The type is slice of
    // Aaa_RadiusAttribute_CalledStation_Formats_Format.
    Format []*Aaa_RadiusAttribute_CalledStation_Formats_Format
}

func (formats *Aaa_RadiusAttribute_CalledStation_Formats) GetEntityData() *types.CommonEntityData {
    formats.EntityData.YFilter = formats.YFilter
    formats.EntityData.YangName = "formats"
    formats.EntityData.BundleName = "cisco_ios_xr"
    formats.EntityData.ParentYangName = "called-station"
    formats.EntityData.SegmentPath = "formats"
    formats.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/called-station/" + formats.EntityData.SegmentPath
    formats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formats.EntityData.Children = types.NewOrderedMap()
    formats.EntityData.Children.Append("format", types.YChild{"Format", nil})
    for i := range formats.Format {
        formats.EntityData.Children.Append(types.GetSegmentPath(formats.Format[i]), types.YChild{"Format", formats.Format[i]})
    }
    formats.EntityData.Leafs = types.NewOrderedMap()

    formats.EntityData.YListKeys = []string {}

    return &(formats.EntityData)
}

// Aaa_RadiusAttribute_CalledStation_Formats_Format
// nas-port-id attribute format
type Aaa_RadiusAttribute_CalledStation_Formats_Format struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is interface{} with range: 0..45.
    Type interface{}

    // AAA nas-port attribute format. The type is string. This attribute is
    // mandatory.
    FormatName interface{}
}

func (format *Aaa_RadiusAttribute_CalledStation_Formats_Format) GetEntityData() *types.CommonEntityData {
    format.EntityData.YFilter = format.YFilter
    format.EntityData.YangName = "format"
    format.EntityData.BundleName = "cisco_ios_xr"
    format.EntityData.ParentYangName = "formats"
    format.EntityData.SegmentPath = "format" + types.AddKeyToken(format.Type, "type")
    format.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/called-station/formats/" + format.EntityData.SegmentPath
    format.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    format.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    format.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    format.EntityData.Children = types.NewOrderedMap()
    format.EntityData.Leafs = types.NewOrderedMap()
    format.EntityData.Leafs.Append("type", types.YLeaf{"Type", format.Type})
    format.EntityData.Leafs.Append("format-name", types.YLeaf{"FormatName", format.FormatName})

    format.EntityData.YListKeys = []string {"Type"}

    return &(format.EntityData)
}

// Aaa_RadiusAttribute_NasPort
// AAA nas-port-id attribute
type Aaa_RadiusAttribute_NasPort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA nas-port-id attribute format.
    FormatExtendeds Aaa_RadiusAttribute_NasPort_FormatExtendeds
}

func (nasPort *Aaa_RadiusAttribute_NasPort) GetEntityData() *types.CommonEntityData {
    nasPort.EntityData.YFilter = nasPort.YFilter
    nasPort.EntityData.YangName = "nas-port"
    nasPort.EntityData.BundleName = "cisco_ios_xr"
    nasPort.EntityData.ParentYangName = "radius-attribute"
    nasPort.EntityData.SegmentPath = "nas-port"
    nasPort.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/" + nasPort.EntityData.SegmentPath
    nasPort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nasPort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nasPort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nasPort.EntityData.Children = types.NewOrderedMap()
    nasPort.EntityData.Children.Append("format-extendeds", types.YChild{"FormatExtendeds", &nasPort.FormatExtendeds})
    nasPort.EntityData.Leafs = types.NewOrderedMap()

    nasPort.EntityData.YListKeys = []string {}

    return &(nasPort.EntityData)
}

// Aaa_RadiusAttribute_NasPort_FormatExtendeds
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_NasPort_FormatExtendeds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // nas-port-id extended attribute. The type is slice of
    // Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended.
    FormatExtended []*Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended
}

func (formatExtendeds *Aaa_RadiusAttribute_NasPort_FormatExtendeds) GetEntityData() *types.CommonEntityData {
    formatExtendeds.EntityData.YFilter = formatExtendeds.YFilter
    formatExtendeds.EntityData.YangName = "format-extendeds"
    formatExtendeds.EntityData.BundleName = "cisco_ios_xr"
    formatExtendeds.EntityData.ParentYangName = "nas-port"
    formatExtendeds.EntityData.SegmentPath = "format-extendeds"
    formatExtendeds.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/nas-port/" + formatExtendeds.EntityData.SegmentPath
    formatExtendeds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formatExtendeds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formatExtendeds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formatExtendeds.EntityData.Children = types.NewOrderedMap()
    formatExtendeds.EntityData.Children.Append("format-extended", types.YChild{"FormatExtended", nil})
    for i := range formatExtendeds.FormatExtended {
        formatExtendeds.EntityData.Children.Append(types.GetSegmentPath(formatExtendeds.FormatExtended[i]), types.YChild{"FormatExtended", formatExtendeds.FormatExtended[i]})
    }
    formatExtendeds.EntityData.Leafs = types.NewOrderedMap()

    formatExtendeds.EntityData.YListKeys = []string {}

    return &(formatExtendeds.EntityData)
}

// Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended
// nas-port-id extended attribute
type Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. format type. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Value interface{}

    // This attribute is a key. AAA nas-port attribute format. The type is
    // interface{} with range: 0..45.
    Type interface{}

    // A 32 character string representing the format to be used. The type is
    // string with length: 1..32.
    FormatIdentifier interface{}
}

func (formatExtended *Aaa_RadiusAttribute_NasPort_FormatExtendeds_FormatExtended) GetEntityData() *types.CommonEntityData {
    formatExtended.EntityData.YFilter = formatExtended.YFilter
    formatExtended.EntityData.YangName = "format-extended"
    formatExtended.EntityData.BundleName = "cisco_ios_xr"
    formatExtended.EntityData.ParentYangName = "format-extendeds"
    formatExtended.EntityData.SegmentPath = "format-extended" + types.AddKeyToken(formatExtended.Value, "value") + types.AddKeyToken(formatExtended.Type, "type")
    formatExtended.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/nas-port/format-extendeds/" + formatExtended.EntityData.SegmentPath
    formatExtended.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formatExtended.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formatExtended.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formatExtended.EntityData.Children = types.NewOrderedMap()
    formatExtended.EntityData.Leafs = types.NewOrderedMap()
    formatExtended.EntityData.Leafs.Append("value", types.YLeaf{"Value", formatExtended.Value})
    formatExtended.EntityData.Leafs.Append("type", types.YLeaf{"Type", formatExtended.Type})
    formatExtended.EntityData.Leafs.Append("format-identifier", types.YLeaf{"FormatIdentifier", formatExtended.FormatIdentifier})

    formatExtended.EntityData.YListKeys = []string {"Value", "Type"}

    return &(formatExtended.EntityData)
}

// Aaa_RadiusAttribute_FormatOthers
// AAA nas-port-id attribute format
type Aaa_RadiusAttribute_FormatOthers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Other configs. The type is slice of
    // Aaa_RadiusAttribute_FormatOthers_FormatOther.
    FormatOther []*Aaa_RadiusAttribute_FormatOthers_FormatOther
}

func (formatOthers *Aaa_RadiusAttribute_FormatOthers) GetEntityData() *types.CommonEntityData {
    formatOthers.EntityData.YFilter = formatOthers.YFilter
    formatOthers.EntityData.YangName = "format-others"
    formatOthers.EntityData.BundleName = "cisco_ios_xr"
    formatOthers.EntityData.ParentYangName = "radius-attribute"
    formatOthers.EntityData.SegmentPath = "format-others"
    formatOthers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/" + formatOthers.EntityData.SegmentPath
    formatOthers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formatOthers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formatOthers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formatOthers.EntityData.Children = types.NewOrderedMap()
    formatOthers.EntityData.Children.Append("format-other", types.YChild{"FormatOther", nil})
    for i := range formatOthers.FormatOther {
        formatOthers.EntityData.Children.Append(types.GetSegmentPath(formatOthers.FormatOther[i]), types.YChild{"FormatOther", formatOthers.FormatOther[i]})
    }
    formatOthers.EntityData.Leafs = types.NewOrderedMap()

    formatOthers.EntityData.YListKeys = []string {}

    return &(formatOthers.EntityData)
}

// Aaa_RadiusAttribute_FormatOthers_FormatOther
// Other configs
type Aaa_RadiusAttribute_FormatOthers_FormatOther struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Nas-Port-Type value to apply format name on. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    NasPortTypeName interface{}

    // Argument1. The type is string.
    AttributeConfig1 interface{}

    // Argument2. The type is string.
    AttributeConfig2 interface{}

    // Argument3. The type is string.
    AttributeConfig3 interface{}

    // Argument4. The type is string.
    AttributeConfig4 interface{}

    // Argument5. The type is string.
    AttributeConfig5 interface{}

    // Argument6. The type is string.
    AttributeConfig6 interface{}

    // Argument7. The type is string.
    AttributeConfig7 interface{}

    // Argument8. The type is string.
    AttributeConfig8 interface{}

    // Argument9. The type is string.
    AttributeConfig9 interface{}

    // Argument10. The type is string.
    AttributeConfig10 interface{}

    // Argument11. The type is string.
    AttributeConfig11 interface{}

    // Argument12. The type is string.
    AttributeConfig12 interface{}

    // Argument13. The type is string.
    AttributeConfig13 interface{}

    // Argument14. The type is string.
    AttributeConfig14 interface{}

    // Argument15. The type is string.
    AttributeConfig15 interface{}

    // Argument16. The type is string.
    AttributeConfig16 interface{}

    // Argument17. The type is string.
    AttributeConfig17 interface{}

    // Argument18. The type is string.
    AttributeConfig18 interface{}

    // Argument19. The type is interface{} with range: 1..253.
    AttributeConfig19 interface{}
}

func (formatOther *Aaa_RadiusAttribute_FormatOthers_FormatOther) GetEntityData() *types.CommonEntityData {
    formatOther.EntityData.YFilter = formatOther.YFilter
    formatOther.EntityData.YangName = "format-other"
    formatOther.EntityData.BundleName = "cisco_ios_xr"
    formatOther.EntityData.ParentYangName = "format-others"
    formatOther.EntityData.SegmentPath = "format-other" + types.AddKeyToken(formatOther.NasPortTypeName, "nas-port-type-name")
    formatOther.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-aaacore-cfg:radius-attribute/format-others/" + formatOther.EntityData.SegmentPath
    formatOther.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    formatOther.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    formatOther.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    formatOther.EntityData.Children = types.NewOrderedMap()
    formatOther.EntityData.Leafs = types.NewOrderedMap()
    formatOther.EntityData.Leafs.Append("nas-port-type-name", types.YLeaf{"NasPortTypeName", formatOther.NasPortTypeName})
    formatOther.EntityData.Leafs.Append("attribute-config1", types.YLeaf{"AttributeConfig1", formatOther.AttributeConfig1})
    formatOther.EntityData.Leafs.Append("attribute-config2", types.YLeaf{"AttributeConfig2", formatOther.AttributeConfig2})
    formatOther.EntityData.Leafs.Append("attribute-config3", types.YLeaf{"AttributeConfig3", formatOther.AttributeConfig3})
    formatOther.EntityData.Leafs.Append("attribute-config4", types.YLeaf{"AttributeConfig4", formatOther.AttributeConfig4})
    formatOther.EntityData.Leafs.Append("attribute-config5", types.YLeaf{"AttributeConfig5", formatOther.AttributeConfig5})
    formatOther.EntityData.Leafs.Append("attribute-config6", types.YLeaf{"AttributeConfig6", formatOther.AttributeConfig6})
    formatOther.EntityData.Leafs.Append("attribute-config7", types.YLeaf{"AttributeConfig7", formatOther.AttributeConfig7})
    formatOther.EntityData.Leafs.Append("attribute-config8", types.YLeaf{"AttributeConfig8", formatOther.AttributeConfig8})
    formatOther.EntityData.Leafs.Append("attribute-config9", types.YLeaf{"AttributeConfig9", formatOther.AttributeConfig9})
    formatOther.EntityData.Leafs.Append("attribute-config10", types.YLeaf{"AttributeConfig10", formatOther.AttributeConfig10})
    formatOther.EntityData.Leafs.Append("attribute-config11", types.YLeaf{"AttributeConfig11", formatOther.AttributeConfig11})
    formatOther.EntityData.Leafs.Append("attribute-config12", types.YLeaf{"AttributeConfig12", formatOther.AttributeConfig12})
    formatOther.EntityData.Leafs.Append("attribute-config13", types.YLeaf{"AttributeConfig13", formatOther.AttributeConfig13})
    formatOther.EntityData.Leafs.Append("attribute-config14", types.YLeaf{"AttributeConfig14", formatOther.AttributeConfig14})
    formatOther.EntityData.Leafs.Append("attribute-config15", types.YLeaf{"AttributeConfig15", formatOther.AttributeConfig15})
    formatOther.EntityData.Leafs.Append("attribute-config16", types.YLeaf{"AttributeConfig16", formatOther.AttributeConfig16})
    formatOther.EntityData.Leafs.Append("attribute-config17", types.YLeaf{"AttributeConfig17", formatOther.AttributeConfig17})
    formatOther.EntityData.Leafs.Append("attribute-config18", types.YLeaf{"AttributeConfig18", formatOther.AttributeConfig18})
    formatOther.EntityData.Leafs.Append("attribute-config19", types.YLeaf{"AttributeConfig19", formatOther.AttributeConfig19})

    formatOther.EntityData.YListKeys = []string {"NasPortTypeName"}

    return &(formatOther.EntityData)
}

// Aaa_Radius
// Remote Access Dial-In User Service
type Aaa_Radius struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of times to retransmit a request to the RADIUS server(0-Disable).
    // The type is interface{} with range: 0..100. The default value is 3.
    Retransmit interface{}

    // This indicates the length of time (in minutes) for which a RADIUS server
    // remains marked as dead. The type is interface{} with range: 1..1440. Units
    // are minute.
    DeadTime interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    Key interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    Timeout interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAccountingPort interface{}

    // Idle time for RADIUS server. The type is interface{} with range: 1..1000.
    // The default value is 5.
    IdleTime interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAuthPort interface{}

    // List of RADIUS servers.
    Hosts Aaa_Radius_Hosts

    // RADIUS server dead criteria.
    DeadCriteria Aaa_Radius_DeadCriteria

    // disallow null-username.
    Disallow Aaa_Radius_Disallow

    // IPv6 configuration.
    Ipv6 Aaa_Radius_Ipv6

    // RADIUS dynamic authorization.
    DynamicAuthorization Aaa_Radius_DynamicAuthorization

    // Radius load-balancing options.
    LoadBalanceOptions Aaa_Radius_LoadBalanceOptions

    // List of VRFs.
    Vrfs Aaa_Radius_Vrfs

    // Radius throttling options.
    Throttle Aaa_Radius_Throttle

    // Unknown VSA (Vendor Specific Attribute) ignore configuration for RADIUS
    // server.
    Vsa Aaa_Radius_Vsa

    // IPv4 configuration.
    Ipv4 Aaa_Radius_Ipv4

    // attribute.
    RadiusAttribute Aaa_Radius_RadiusAttribute

    // Table of attribute list.
    Attributes Aaa_Radius_Attributes

    // Source port.
    SourcePort Aaa_Radius_SourcePort
}

func (radius *Aaa_Radius) GetEntityData() *types.CommonEntityData {
    radius.EntityData.YFilter = radius.YFilter
    radius.EntityData.YangName = "radius"
    radius.EntityData.BundleName = "cisco_ios_xr"
    radius.EntityData.ParentYangName = "aaa"
    radius.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-protocol-radius-cfg:radius"
    radius.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + radius.EntityData.SegmentPath
    radius.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radius.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radius.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radius.EntityData.Children = types.NewOrderedMap()
    radius.EntityData.Children.Append("hosts", types.YChild{"Hosts", &radius.Hosts})
    radius.EntityData.Children.Append("dead-criteria", types.YChild{"DeadCriteria", &radius.DeadCriteria})
    radius.EntityData.Children.Append("disallow", types.YChild{"Disallow", &radius.Disallow})
    radius.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &radius.Ipv6})
    radius.EntityData.Children.Append("dynamic-authorization", types.YChild{"DynamicAuthorization", &radius.DynamicAuthorization})
    radius.EntityData.Children.Append("load-balance-options", types.YChild{"LoadBalanceOptions", &radius.LoadBalanceOptions})
    radius.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &radius.Vrfs})
    radius.EntityData.Children.Append("throttle", types.YChild{"Throttle", &radius.Throttle})
    radius.EntityData.Children.Append("vsa", types.YChild{"Vsa", &radius.Vsa})
    radius.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &radius.Ipv4})
    radius.EntityData.Children.Append("radius-attribute", types.YChild{"RadiusAttribute", &radius.RadiusAttribute})
    radius.EntityData.Children.Append("attributes", types.YChild{"Attributes", &radius.Attributes})
    radius.EntityData.Children.Append("source-port", types.YChild{"SourcePort", &radius.SourcePort})
    radius.EntityData.Leafs = types.NewOrderedMap()
    radius.EntityData.Leafs.Append("retransmit", types.YLeaf{"Retransmit", radius.Retransmit})
    radius.EntityData.Leafs.Append("dead-time", types.YLeaf{"DeadTime", radius.DeadTime})
    radius.EntityData.Leafs.Append("key", types.YLeaf{"Key", radius.Key})
    radius.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", radius.Timeout})
    radius.EntityData.Leafs.Append("ignore-accounting-port", types.YLeaf{"IgnoreAccountingPort", radius.IgnoreAccountingPort})
    radius.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", radius.IdleTime})
    radius.EntityData.Leafs.Append("username", types.YLeaf{"Username", radius.Username})
    radius.EntityData.Leafs.Append("ignore-auth-port", types.YLeaf{"IgnoreAuthPort", radius.IgnoreAuthPort})

    radius.EntityData.YListKeys = []string {}

    return &(radius.EntityData)
}

// Aaa_Radius_Hosts
// List of RADIUS servers
type Aaa_Radius_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance of a RADIUS server. The type is slice of Aaa_Radius_Hosts_Host.
    Host []*Aaa_Radius_Hosts_Host
}

func (hosts *Aaa_Radius_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "radius"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + hosts.EntityData.SegmentPath
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Aaa_Radius_Hosts_Host
// Instance of a RADIUS server
type Aaa_Radius_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This is used to sort the servers in the order of
    // precedence. The type is interface{} with range: 0..4294967295.
    OrderingIndex interface{}

    // This attribute is a key. IP address of RADIUS server. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Authentication Port number (standard port 1645).
    // The type is interface{} with range: 0..65535.
    AuthPortNumber interface{}

    // This attribute is a key. Accounting Port number (standard port 1646). The
    // type is interface{} with range: 0..65535.
    AcctPortNumber interface{}

    // Number of times to retransmit a request to the RADIUS server. The type is
    // interface{} with range: 1..100. The default value is 3.
    HostRetransmit interface{}

    // Time to wait for a RADIUS server to reply. The type is interface{} with
    // range: 1..1000. The default value is 5.
    HostTimeout interface{}

    // RADIUS encryption key. The type is string with pattern: (!.+)|([^!].+).
    HostKey interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAccountingPort interface{}

    // Idle time for RADIUS server. The type is interface{} with range: 1..1000.
    // The default value is 5.
    IdleTime interface{}

    // Username to be tested for automated testing. The type is string.
    Username interface{}

    // Time to wait for a RADIUS server to reply. The type is bool.
    IgnoreAuthPort interface{}
}

func (host *Aaa_Radius_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.OrderingIndex, "ordering-index") + types.AddKeyToken(host.IpAddress, "ip-address") + types.AddKeyToken(host.AuthPortNumber, "auth-port-number") + types.AddKeyToken(host.AcctPortNumber, "acct-port-number")
    host.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/hosts/" + host.EntityData.SegmentPath
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ordering-index", types.YLeaf{"OrderingIndex", host.OrderingIndex})
    host.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", host.IpAddress})
    host.EntityData.Leafs.Append("auth-port-number", types.YLeaf{"AuthPortNumber", host.AuthPortNumber})
    host.EntityData.Leafs.Append("acct-port-number", types.YLeaf{"AcctPortNumber", host.AcctPortNumber})
    host.EntityData.Leafs.Append("host-retransmit", types.YLeaf{"HostRetransmit", host.HostRetransmit})
    host.EntityData.Leafs.Append("host-timeout", types.YLeaf{"HostTimeout", host.HostTimeout})
    host.EntityData.Leafs.Append("host-key", types.YLeaf{"HostKey", host.HostKey})
    host.EntityData.Leafs.Append("ignore-accounting-port", types.YLeaf{"IgnoreAccountingPort", host.IgnoreAccountingPort})
    host.EntityData.Leafs.Append("idle-time", types.YLeaf{"IdleTime", host.IdleTime})
    host.EntityData.Leafs.Append("username", types.YLeaf{"Username", host.Username})
    host.EntityData.Leafs.Append("ignore-auth-port", types.YLeaf{"IgnoreAuthPort", host.IgnoreAuthPort})

    host.EntityData.YListKeys = []string {"OrderingIndex", "IpAddress", "AuthPortNumber", "AcctPortNumber"}

    return &(host.EntityData)
}

// Aaa_Radius_DeadCriteria
// RADIUS server dead criteria
type Aaa_Radius_DeadCriteria struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of consecutive timeouts the router must experience in order to
    // mark the server as dead. All transmissions, including the initial transmit
    // and all retransmits, will be counted. The type is interface{} with range:
    // 1..100.
    Tries interface{}

    // The minimum amount of time which must elapse since the router last received
    // a valid RADIUS packet from the server prior to marking it dead. The type is
    // interface{} with range: 1..120.
    Time interface{}
}

func (deadCriteria *Aaa_Radius_DeadCriteria) GetEntityData() *types.CommonEntityData {
    deadCriteria.EntityData.YFilter = deadCriteria.YFilter
    deadCriteria.EntityData.YangName = "dead-criteria"
    deadCriteria.EntityData.BundleName = "cisco_ios_xr"
    deadCriteria.EntityData.ParentYangName = "radius"
    deadCriteria.EntityData.SegmentPath = "dead-criteria"
    deadCriteria.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + deadCriteria.EntityData.SegmentPath
    deadCriteria.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deadCriteria.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deadCriteria.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deadCriteria.EntityData.Children = types.NewOrderedMap()
    deadCriteria.EntityData.Leafs = types.NewOrderedMap()
    deadCriteria.EntityData.Leafs.Append("tries", types.YLeaf{"Tries", deadCriteria.Tries})
    deadCriteria.EntityData.Leafs.Append("time", types.YLeaf{"Time", deadCriteria.Time})

    deadCriteria.EntityData.YListKeys = []string {}

    return &(deadCriteria.EntityData)
}

// Aaa_Radius_Disallow
// disallow null-username
type Aaa_Radius_Disallow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disallow null-username. The type is interface{} with range: 0..4294967295.
    NullUsername interface{}
}

func (disallow *Aaa_Radius_Disallow) GetEntityData() *types.CommonEntityData {
    disallow.EntityData.YFilter = disallow.YFilter
    disallow.EntityData.YangName = "disallow"
    disallow.EntityData.BundleName = "cisco_ios_xr"
    disallow.EntityData.ParentYangName = "radius"
    disallow.EntityData.SegmentPath = "disallow"
    disallow.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + disallow.EntityData.SegmentPath
    disallow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disallow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disallow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disallow.EntityData.Children = types.NewOrderedMap()
    disallow.EntityData.Leafs = types.NewOrderedMap()
    disallow.EntityData.Leafs.Append("null-username", types.YLeaf{"NullUsername", disallow.NullUsername})

    disallow.EntityData.YListKeys = []string {}

    return &(disallow.EntityData)
}

// Aaa_Radius_Ipv6
// IPv6 configuration
type Aaa_Radius_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // AaaDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv6 *Aaa_Radius_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "radius"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ipv6.Dscp})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Aaa_Radius_DynamicAuthorization
// RADIUS dynamic authorization
type Aaa_Radius_DynamicAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore option for server key or session key. The type is AaaSelectKey.
    Ignore interface{}

    // Specify the COA server port to listen on. The type is interface{} with
    // range: 1000..5000.
    Port interface{}

    // RADIUS  dynamic  authorization  type. The type is AaaAuthentication.
    AuthenticationType interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}

    // Client data.
    Clients Aaa_Radius_DynamicAuthorization_Clients
}

func (dynamicAuthorization *Aaa_Radius_DynamicAuthorization) GetEntityData() *types.CommonEntityData {
    dynamicAuthorization.EntityData.YFilter = dynamicAuthorization.YFilter
    dynamicAuthorization.EntityData.YangName = "dynamic-authorization"
    dynamicAuthorization.EntityData.BundleName = "cisco_ios_xr"
    dynamicAuthorization.EntityData.ParentYangName = "radius"
    dynamicAuthorization.EntityData.SegmentPath = "dynamic-authorization"
    dynamicAuthorization.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + dynamicAuthorization.EntityData.SegmentPath
    dynamicAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicAuthorization.EntityData.Children = types.NewOrderedMap()
    dynamicAuthorization.EntityData.Children.Append("clients", types.YChild{"Clients", &dynamicAuthorization.Clients})
    dynamicAuthorization.EntityData.Leafs = types.NewOrderedMap()
    dynamicAuthorization.EntityData.Leafs.Append("ignore", types.YLeaf{"Ignore", dynamicAuthorization.Ignore})
    dynamicAuthorization.EntityData.Leafs.Append("port", types.YLeaf{"Port", dynamicAuthorization.Port})
    dynamicAuthorization.EntityData.Leafs.Append("authentication-type", types.YLeaf{"AuthenticationType", dynamicAuthorization.AuthenticationType})
    dynamicAuthorization.EntityData.Leafs.Append("server-key", types.YLeaf{"ServerKey", dynamicAuthorization.ServerKey})

    dynamicAuthorization.EntityData.YListKeys = []string {}

    return &(dynamicAuthorization.EntityData)
}

// Aaa_Radius_DynamicAuthorization_Clients
// Client data
type Aaa_Radius_DynamicAuthorization_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client data. The type is slice of
    // Aaa_Radius_DynamicAuthorization_Clients_Client.
    Client []*Aaa_Radius_DynamicAuthorization_Clients_Client

    // Client data. The type is slice of
    // Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName.
    ClientVrfName []*Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName
}

func (clients *Aaa_Radius_DynamicAuthorization_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "dynamic-authorization"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/dynamic-authorization/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Children.Append("client-vrf-name", types.YChild{"ClientVrfName", nil})
    for i := range clients.ClientVrfName {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.ClientVrfName[i]), types.YChild{"ClientVrfName", clients.ClientVrfName[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Aaa_Radius_DynamicAuthorization_Clients_Client
// Client data
type Aaa_Radius_DynamicAuthorization_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address of COA client. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}
}

func (client *Aaa_Radius_DynamicAuthorization_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.IpAddress, "ip-address")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/dynamic-authorization/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", client.IpAddress})
    client.EntityData.Leafs.Append("server-key", types.YLeaf{"ServerKey", client.ServerKey})

    client.EntityData.YListKeys = []string {"IpAddress"}

    return &(client.EntityData)
}

// Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName
// Client data
type Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // This attribute is a key. IP address of COA client. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // RADIUS CoA client encryption key. The type is string with pattern:
    // (!.+)|([^!].+).
    ServerKey interface{}
}

func (clientVrfName *Aaa_Radius_DynamicAuthorization_Clients_ClientVrfName) GetEntityData() *types.CommonEntityData {
    clientVrfName.EntityData.YFilter = clientVrfName.YFilter
    clientVrfName.EntityData.YangName = "client-vrf-name"
    clientVrfName.EntityData.BundleName = "cisco_ios_xr"
    clientVrfName.EntityData.ParentYangName = "clients"
    clientVrfName.EntityData.SegmentPath = "client-vrf-name" + types.AddKeyToken(clientVrfName.VrfName, "vrf-name") + types.AddKeyToken(clientVrfName.IpAddress, "ip-address")
    clientVrfName.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/dynamic-authorization/clients/" + clientVrfName.EntityData.SegmentPath
    clientVrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientVrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientVrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientVrfName.EntityData.Children = types.NewOrderedMap()
    clientVrfName.EntityData.Leafs = types.NewOrderedMap()
    clientVrfName.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", clientVrfName.VrfName})
    clientVrfName.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", clientVrfName.IpAddress})
    clientVrfName.EntityData.Leafs.Append("server-key", types.YLeaf{"ServerKey", clientVrfName.ServerKey})

    clientVrfName.EntityData.YListKeys = []string {"VrfName", "IpAddress"}

    return &(clientVrfName.EntityData)
}

// Aaa_Radius_LoadBalanceOptions
// Radius load-balancing options
type Aaa_Radius_LoadBalanceOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Method by which the next host will be picked.
    LoadBalanceMethod Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod
}

func (loadBalanceOptions *Aaa_Radius_LoadBalanceOptions) GetEntityData() *types.CommonEntityData {
    loadBalanceOptions.EntityData.YFilter = loadBalanceOptions.YFilter
    loadBalanceOptions.EntityData.YangName = "load-balance-options"
    loadBalanceOptions.EntityData.BundleName = "cisco_ios_xr"
    loadBalanceOptions.EntityData.ParentYangName = "radius"
    loadBalanceOptions.EntityData.SegmentPath = "load-balance-options"
    loadBalanceOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + loadBalanceOptions.EntityData.SegmentPath
    loadBalanceOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalanceOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalanceOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalanceOptions.EntityData.Children = types.NewOrderedMap()
    loadBalanceOptions.EntityData.Children.Append("load-balance-method", types.YChild{"LoadBalanceMethod", &loadBalanceOptions.LoadBalanceMethod})
    loadBalanceOptions.EntityData.Leafs = types.NewOrderedMap()

    loadBalanceOptions.EntityData.YListKeys = []string {}

    return &(loadBalanceOptions.EntityData)
}

// Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod
// Method by which the next host will be picked
type Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Batch size for selection of the server.
    BatchSize Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize
}

func (loadBalanceMethod *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod) GetEntityData() *types.CommonEntityData {
    loadBalanceMethod.EntityData.YFilter = loadBalanceMethod.YFilter
    loadBalanceMethod.EntityData.YangName = "load-balance-method"
    loadBalanceMethod.EntityData.BundleName = "cisco_ios_xr"
    loadBalanceMethod.EntityData.ParentYangName = "load-balance-options"
    loadBalanceMethod.EntityData.SegmentPath = "load-balance-method"
    loadBalanceMethod.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/load-balance-options/" + loadBalanceMethod.EntityData.SegmentPath
    loadBalanceMethod.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    loadBalanceMethod.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    loadBalanceMethod.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    loadBalanceMethod.EntityData.Children = types.NewOrderedMap()
    loadBalanceMethod.EntityData.Children.Append("batch-size", types.YChild{"BatchSize", &loadBalanceMethod.BatchSize})
    loadBalanceMethod.EntityData.Leafs = types.NewOrderedMap()

    loadBalanceMethod.EntityData.YListKeys = []string {}

    return &(loadBalanceMethod.EntityData)
}

// Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize
// Batch size for selection of the server
type Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Batch size for selection of the server. The type is interface{} with range:
    // 1..1500. The default value is 25.
    BatchSize interface{}

    // Disable preferred server for this Server Group. The type is bool. The
    // default value is true.
    IgnorePreferredServer interface{}
}

func (batchSize *Aaa_Radius_LoadBalanceOptions_LoadBalanceMethod_BatchSize) GetEntityData() *types.CommonEntityData {
    batchSize.EntityData.YFilter = batchSize.YFilter
    batchSize.EntityData.YangName = "batch-size"
    batchSize.EntityData.BundleName = "cisco_ios_xr"
    batchSize.EntityData.ParentYangName = "load-balance-method"
    batchSize.EntityData.SegmentPath = "batch-size"
    batchSize.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/load-balance-options/load-balance-method/" + batchSize.EntityData.SegmentPath
    batchSize.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    batchSize.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    batchSize.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    batchSize.EntityData.Children = types.NewOrderedMap()
    batchSize.EntityData.Leafs = types.NewOrderedMap()
    batchSize.EntityData.Leafs.Append("batch-size", types.YLeaf{"BatchSize", batchSize.BatchSize})
    batchSize.EntityData.Leafs.Append("ignore-preferred-server", types.YLeaf{"IgnorePreferredServer", batchSize.IgnorePreferredServer})

    batchSize.EntityData.YListKeys = []string {}

    return &(batchSize.EntityData)
}

// Aaa_Radius_Vrfs
// List of VRFs
type Aaa_Radius_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRF. The type is slice of Aaa_Radius_Vrfs_Vrf.
    Vrf []*Aaa_Radius_Vrfs_Vrf
}

func (vrfs *Aaa_Radius_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "radius"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Aaa_Radius_Vrfs_Vrf
// A VRF
type Aaa_Radius_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. Specify 'default' for defalut VRF. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in RADIUS packets. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}
}

func (vrf *Aaa_Radius_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", vrf.SourceInterface})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Aaa_Radius_Throttle
// Radius throttling options
type Aaa_Radius_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // To flow control the number of access requests sent to a radius server. The
    // type is interface{} with range: 0..65535. The default value is 0.
    Access interface{}

    // To flow control the number of accounting requests sent to a radius server.
    // The type is interface{} with range: 0..65535. The default value is 0.
    Accounting interface{}

    // Specify the number of timeouts exceeding which a throttled access request
    // is dropped. The type is interface{} with range: 1..10. The default value is
    // 3.
    AccessTimeout interface{}
}

func (throttle *Aaa_Radius_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "radius"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + throttle.EntityData.SegmentPath
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("access", types.YLeaf{"Access", throttle.Access})
    throttle.EntityData.Leafs.Append("accounting", types.YLeaf{"Accounting", throttle.Accounting})
    throttle.EntityData.Leafs.Append("access-timeout", types.YLeaf{"AccessTimeout", throttle.AccessTimeout})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Aaa_Radius_Vsa
// Unknown VSA (Vendor Specific Attribute) ignore
// configuration for RADIUS server
type Aaa_Radius_Vsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vendor Specific Attribute.
    Attribute Aaa_Radius_Vsa_Attribute
}

func (vsa *Aaa_Radius_Vsa) GetEntityData() *types.CommonEntityData {
    vsa.EntityData.YFilter = vsa.YFilter
    vsa.EntityData.YangName = "vsa"
    vsa.EntityData.BundleName = "cisco_ios_xr"
    vsa.EntityData.ParentYangName = "radius"
    vsa.EntityData.SegmentPath = "vsa"
    vsa.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + vsa.EntityData.SegmentPath
    vsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vsa.EntityData.Children = types.NewOrderedMap()
    vsa.EntityData.Children.Append("attribute", types.YChild{"Attribute", &vsa.Attribute})
    vsa.EntityData.Leafs = types.NewOrderedMap()

    vsa.EntityData.YListKeys = []string {}

    return &(vsa.EntityData)
}

// Aaa_Radius_Vsa_Attribute
// Vendor Specific Attribute
type Aaa_Radius_Vsa_Attribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore the VSA.
    Ignore Aaa_Radius_Vsa_Attribute_Ignore
}

func (attribute *Aaa_Radius_Vsa_Attribute) GetEntityData() *types.CommonEntityData {
    attribute.EntityData.YFilter = attribute.YFilter
    attribute.EntityData.YangName = "attribute"
    attribute.EntityData.BundleName = "cisco_ios_xr"
    attribute.EntityData.ParentYangName = "vsa"
    attribute.EntityData.SegmentPath = "attribute"
    attribute.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/vsa/" + attribute.EntityData.SegmentPath
    attribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attribute.EntityData.Children = types.NewOrderedMap()
    attribute.EntityData.Children.Append("ignore", types.YChild{"Ignore", &attribute.Ignore})
    attribute.EntityData.Leafs = types.NewOrderedMap()

    attribute.EntityData.YListKeys = []string {}

    return &(attribute.EntityData)
}

// Aaa_Radius_Vsa_Attribute_Ignore
// Ignore the VSA
type Aaa_Radius_Vsa_Attribute_Ignore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore the VSA and no prefix will reject the unknown VSA. The type is
    // interface{}.
    Unknown interface{}
}

func (ignore *Aaa_Radius_Vsa_Attribute_Ignore) GetEntityData() *types.CommonEntityData {
    ignore.EntityData.YFilter = ignore.YFilter
    ignore.EntityData.YangName = "ignore"
    ignore.EntityData.BundleName = "cisco_ios_xr"
    ignore.EntityData.ParentYangName = "attribute"
    ignore.EntityData.SegmentPath = "ignore"
    ignore.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/vsa/attribute/" + ignore.EntityData.SegmentPath
    ignore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ignore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ignore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ignore.EntityData.Children = types.NewOrderedMap()
    ignore.EntityData.Leafs = types.NewOrderedMap()
    ignore.EntityData.Leafs.Append("unknown", types.YLeaf{"Unknown", ignore.Unknown})

    ignore.EntityData.YListKeys = []string {}

    return &(ignore.EntityData)
}

// Aaa_Radius_Ipv4
// IPv4 configuration
type Aaa_Radius_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is one of the following types: enumeration
    // AaaDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Aaa_Radius_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "radius"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ipv4.Dscp})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Aaa_Radius_RadiusAttribute
// attribute
type Aaa_Radius_RadiusAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter-Id attribute configuration.
    FilterId11 Aaa_Radius_RadiusAttribute_FilterId11

    // Acct-Session-Id attribute(44).
    AcctMultiSessionId Aaa_Radius_RadiusAttribute_AcctMultiSessionId

    // Acct-Session-Id attribute(44).
    AcctSessionId Aaa_Radius_RadiusAttribute_AcctSessionId
}

func (radiusAttribute *Aaa_Radius_RadiusAttribute) GetEntityData() *types.CommonEntityData {
    radiusAttribute.EntityData.YFilter = radiusAttribute.YFilter
    radiusAttribute.EntityData.YangName = "radius-attribute"
    radiusAttribute.EntityData.BundleName = "cisco_ios_xr"
    radiusAttribute.EntityData.ParentYangName = "radius"
    radiusAttribute.EntityData.SegmentPath = "radius-attribute"
    radiusAttribute.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + radiusAttribute.EntityData.SegmentPath
    radiusAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    radiusAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    radiusAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    radiusAttribute.EntityData.Children = types.NewOrderedMap()
    radiusAttribute.EntityData.Children.Append("filter-id-11", types.YChild{"FilterId11", &radiusAttribute.FilterId11})
    radiusAttribute.EntityData.Children.Append("acct-multi-session-id", types.YChild{"AcctMultiSessionId", &radiusAttribute.AcctMultiSessionId})
    radiusAttribute.EntityData.Children.Append("acct-session-id", types.YChild{"AcctSessionId", &radiusAttribute.AcctSessionId})
    radiusAttribute.EntityData.Leafs = types.NewOrderedMap()

    radiusAttribute.EntityData.YListKeys = []string {}

    return &(radiusAttribute.EntityData)
}

// Aaa_Radius_RadiusAttribute_FilterId11
// Filter-Id attribute configuration
type Aaa_Radius_RadiusAttribute_FilterId11 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the attribute default direction.
    Defaults Aaa_Radius_RadiusAttribute_FilterId11_Defaults
}

func (filterId11 *Aaa_Radius_RadiusAttribute_FilterId11) GetEntityData() *types.CommonEntityData {
    filterId11.EntityData.YFilter = filterId11.YFilter
    filterId11.EntityData.YangName = "filter-id-11"
    filterId11.EntityData.BundleName = "cisco_ios_xr"
    filterId11.EntityData.ParentYangName = "radius-attribute"
    filterId11.EntityData.SegmentPath = "filter-id-11"
    filterId11.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/" + filterId11.EntityData.SegmentPath
    filterId11.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filterId11.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filterId11.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filterId11.EntityData.Children = types.NewOrderedMap()
    filterId11.EntityData.Children.Append("defaults", types.YChild{"Defaults", &filterId11.Defaults})
    filterId11.EntityData.Leafs = types.NewOrderedMap()

    filterId11.EntityData.YListKeys = []string {}

    return &(filterId11.EntityData)
}

// Aaa_Radius_RadiusAttribute_FilterId11_Defaults
// Set the attribute default direction
type Aaa_Radius_RadiusAttribute_FilterId11_Defaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filtering is applied to ingress(inbound)/egress(outbound) packets only. The
    // type is AaaDirection.
    Direction interface{}
}

func (defaults *Aaa_Radius_RadiusAttribute_FilterId11_Defaults) GetEntityData() *types.CommonEntityData {
    defaults.EntityData.YFilter = defaults.YFilter
    defaults.EntityData.YangName = "defaults"
    defaults.EntityData.BundleName = "cisco_ios_xr"
    defaults.EntityData.ParentYangName = "filter-id-11"
    defaults.EntityData.SegmentPath = "defaults"
    defaults.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/filter-id-11/" + defaults.EntityData.SegmentPath
    defaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaults.EntityData.Children = types.NewOrderedMap()
    defaults.EntityData.Leafs = types.NewOrderedMap()
    defaults.EntityData.Leafs.Append("direction", types.YLeaf{"Direction", defaults.Direction})

    defaults.EntityData.YListKeys = []string {}

    return &(defaults.EntityData)
}

// Aaa_Radius_RadiusAttribute_AcctMultiSessionId
// Acct-Session-Id attribute(44)
type Aaa_Radius_RadiusAttribute_AcctMultiSessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prepend Acct-Session-Id attribute with Nas-Port-Id.
    IncludeParentSessionId Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId
}

func (acctMultiSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId) GetEntityData() *types.CommonEntityData {
    acctMultiSessionId.EntityData.YFilter = acctMultiSessionId.YFilter
    acctMultiSessionId.EntityData.YangName = "acct-multi-session-id"
    acctMultiSessionId.EntityData.BundleName = "cisco_ios_xr"
    acctMultiSessionId.EntityData.ParentYangName = "radius-attribute"
    acctMultiSessionId.EntityData.SegmentPath = "acct-multi-session-id"
    acctMultiSessionId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/" + acctMultiSessionId.EntityData.SegmentPath
    acctMultiSessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acctMultiSessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acctMultiSessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acctMultiSessionId.EntityData.Children = types.NewOrderedMap()
    acctMultiSessionId.EntityData.Children.Append("include-parent-session-id", types.YChild{"IncludeParentSessionId", &acctMultiSessionId.IncludeParentSessionId})
    acctMultiSessionId.EntityData.Leafs = types.NewOrderedMap()

    acctMultiSessionId.EntityData.YListKeys = []string {}

    return &(acctMultiSessionId.EntityData)
}

// Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId
// Prepend Acct-Session-Id attribute with
// Nas-Port-Id
type Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // false/true. The type is AaaConfig.
    Config interface{}
}

func (includeParentSessionId *Aaa_Radius_RadiusAttribute_AcctMultiSessionId_IncludeParentSessionId) GetEntityData() *types.CommonEntityData {
    includeParentSessionId.EntityData.YFilter = includeParentSessionId.YFilter
    includeParentSessionId.EntityData.YangName = "include-parent-session-id"
    includeParentSessionId.EntityData.BundleName = "cisco_ios_xr"
    includeParentSessionId.EntityData.ParentYangName = "acct-multi-session-id"
    includeParentSessionId.EntityData.SegmentPath = "include-parent-session-id"
    includeParentSessionId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/acct-multi-session-id/" + includeParentSessionId.EntityData.SegmentPath
    includeParentSessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    includeParentSessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    includeParentSessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    includeParentSessionId.EntityData.Children = types.NewOrderedMap()
    includeParentSessionId.EntityData.Leafs = types.NewOrderedMap()
    includeParentSessionId.EntityData.Leafs.Append("config", types.YLeaf{"Config", includeParentSessionId.Config})

    includeParentSessionId.EntityData.YListKeys = []string {}

    return &(includeParentSessionId.EntityData)
}

// Aaa_Radius_RadiusAttribute_AcctSessionId
// Acct-Session-Id attribute(44)
type Aaa_Radius_RadiusAttribute_AcctSessionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prepend Acct-Session-Id attribute with Nas-Port-Id.
    PrependNasPortId Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId
}

func (acctSessionId *Aaa_Radius_RadiusAttribute_AcctSessionId) GetEntityData() *types.CommonEntityData {
    acctSessionId.EntityData.YFilter = acctSessionId.YFilter
    acctSessionId.EntityData.YangName = "acct-session-id"
    acctSessionId.EntityData.BundleName = "cisco_ios_xr"
    acctSessionId.EntityData.ParentYangName = "radius-attribute"
    acctSessionId.EntityData.SegmentPath = "acct-session-id"
    acctSessionId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/" + acctSessionId.EntityData.SegmentPath
    acctSessionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acctSessionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acctSessionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acctSessionId.EntityData.Children = types.NewOrderedMap()
    acctSessionId.EntityData.Children.Append("prepend-nas-port-id", types.YChild{"PrependNasPortId", &acctSessionId.PrependNasPortId})
    acctSessionId.EntityData.Leafs = types.NewOrderedMap()

    acctSessionId.EntityData.YListKeys = []string {}

    return &(acctSessionId.EntityData)
}

// Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId
// Prepend Acct-Session-Id attribute with
// Nas-Port-Id
type Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // false/true. The type is AaaConfig.
    Config interface{}
}

func (prependNasPortId *Aaa_Radius_RadiusAttribute_AcctSessionId_PrependNasPortId) GetEntityData() *types.CommonEntityData {
    prependNasPortId.EntityData.YFilter = prependNasPortId.YFilter
    prependNasPortId.EntityData.YangName = "prepend-nas-port-id"
    prependNasPortId.EntityData.BundleName = "cisco_ios_xr"
    prependNasPortId.EntityData.ParentYangName = "acct-session-id"
    prependNasPortId.EntityData.SegmentPath = "prepend-nas-port-id"
    prependNasPortId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/radius-attribute/acct-session-id/" + prependNasPortId.EntityData.SegmentPath
    prependNasPortId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prependNasPortId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prependNasPortId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prependNasPortId.EntityData.Children = types.NewOrderedMap()
    prependNasPortId.EntityData.Leafs = types.NewOrderedMap()
    prependNasPortId.EntityData.Leafs.Append("config", types.YLeaf{"Config", prependNasPortId.Config})

    prependNasPortId.EntityData.YListKeys = []string {}

    return &(prependNasPortId.EntityData)
}

// Aaa_Radius_Attributes
// Table of attribute list
type Aaa_Radius_Attributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attribute list name. The type is slice of Aaa_Radius_Attributes_Attribute.
    Attribute []*Aaa_Radius_Attributes_Attribute
}

func (attributes *Aaa_Radius_Attributes) GetEntityData() *types.CommonEntityData {
    attributes.EntityData.YFilter = attributes.YFilter
    attributes.EntityData.YangName = "attributes"
    attributes.EntityData.BundleName = "cisco_ios_xr"
    attributes.EntityData.ParentYangName = "radius"
    attributes.EntityData.SegmentPath = "attributes"
    attributes.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + attributes.EntityData.SegmentPath
    attributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributes.EntityData.Children = types.NewOrderedMap()
    attributes.EntityData.Children.Append("attribute", types.YChild{"Attribute", nil})
    for i := range attributes.Attribute {
        attributes.EntityData.Children.Append(types.GetSegmentPath(attributes.Attribute[i]), types.YChild{"Attribute", attributes.Attribute[i]})
    }
    attributes.EntityData.Leafs = types.NewOrderedMap()

    attributes.EntityData.YListKeys = []string {}

    return &(attributes.EntityData)
}

// Aaa_Radius_Attributes_Attribute
// Attribute list name
type Aaa_Radius_Attributes_Attribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Attribute list name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AttributeListName interface{}

    // Specify RADIUS attribute. The type is string.
    Attribute interface{}

    // Vendor Specific Attribute.
    VendorIds Aaa_Radius_Attributes_Attribute_VendorIds
}

func (attribute *Aaa_Radius_Attributes_Attribute) GetEntityData() *types.CommonEntityData {
    attribute.EntityData.YFilter = attribute.YFilter
    attribute.EntityData.YangName = "attribute"
    attribute.EntityData.BundleName = "cisco_ios_xr"
    attribute.EntityData.ParentYangName = "attributes"
    attribute.EntityData.SegmentPath = "attribute" + types.AddKeyToken(attribute.AttributeListName, "attribute-list-name")
    attribute.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/" + attribute.EntityData.SegmentPath
    attribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attribute.EntityData.Children = types.NewOrderedMap()
    attribute.EntityData.Children.Append("vendor-ids", types.YChild{"VendorIds", &attribute.VendorIds})
    attribute.EntityData.Leafs = types.NewOrderedMap()
    attribute.EntityData.Leafs.Append("attribute-list-name", types.YLeaf{"AttributeListName", attribute.AttributeListName})
    attribute.EntityData.Leafs.Append("attribute", types.YLeaf{"Attribute", attribute.Attribute})

    attribute.EntityData.YListKeys = []string {"AttributeListName"}

    return &(attribute.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds
// Vendor Specific Attribute
type Aaa_Radius_Attributes_Attribute_VendorIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Vendor ID of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId.
    VendorId []*Aaa_Radius_Attributes_Attribute_VendorIds_VendorId
}

func (vendorIds *Aaa_Radius_Attributes_Attribute_VendorIds) GetEntityData() *types.CommonEntityData {
    vendorIds.EntityData.YFilter = vendorIds.YFilter
    vendorIds.EntityData.YangName = "vendor-ids"
    vendorIds.EntityData.BundleName = "cisco_ios_xr"
    vendorIds.EntityData.ParentYangName = "attribute"
    vendorIds.EntityData.SegmentPath = "vendor-ids"
    vendorIds.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/" + vendorIds.EntityData.SegmentPath
    vendorIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendorIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendorIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendorIds.EntityData.Children = types.NewOrderedMap()
    vendorIds.EntityData.Children.Append("vendor-id", types.YChild{"VendorId", nil})
    for i := range vendorIds.VendorId {
        vendorIds.EntityData.Children.Append(types.GetSegmentPath(vendorIds.VendorId[i]), types.YChild{"VendorId", vendorIds.VendorId[i]})
    }
    vendorIds.EntityData.Leafs = types.NewOrderedMap()

    vendorIds.EntityData.YListKeys = []string {}

    return &(vendorIds.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId
// Vendor ID of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Vendor Id of vsa. The type is interface{} with
    // range: 0..4294967295.
    VendorId interface{}

    // Vendor Type of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType.
    VendorType []*Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType
}

func (vendorId *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId) GetEntityData() *types.CommonEntityData {
    vendorId.EntityData.YFilter = vendorId.YFilter
    vendorId.EntityData.YangName = "vendor-id"
    vendorId.EntityData.BundleName = "cisco_ios_xr"
    vendorId.EntityData.ParentYangName = "vendor-ids"
    vendorId.EntityData.SegmentPath = "vendor-id" + types.AddKeyToken(vendorId.VendorId, "vendor-id")
    vendorId.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/vendor-ids/" + vendorId.EntityData.SegmentPath
    vendorId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendorId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendorId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendorId.EntityData.Children = types.NewOrderedMap()
    vendorId.EntityData.Children.Append("vendor-type", types.YChild{"VendorType", nil})
    for i := range vendorId.VendorType {
        vendorId.EntityData.Children.Append(types.GetSegmentPath(vendorId.VendorType[i]), types.YChild{"VendorType", vendorId.VendorType[i]})
    }
    vendorId.EntityData.Leafs = types.NewOrderedMap()
    vendorId.EntityData.Leafs.Append("vendor-id", types.YLeaf{"VendorId", vendorId.VendorId})

    vendorId.EntityData.YListKeys = []string {"VendorId"}

    return &(vendorId.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType
// Vendor Type of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Vendor Type of vsa. The type is interface{} with
    // range: 0..4294967295.
    VendorType interface{}

    // Attribute Name of vsa. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName.
    AttributeName []*Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName
}

func (vendorType *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType) GetEntityData() *types.CommonEntityData {
    vendorType.EntityData.YFilter = vendorType.YFilter
    vendorType.EntityData.YangName = "vendor-type"
    vendorType.EntityData.BundleName = "cisco_ios_xr"
    vendorType.EntityData.ParentYangName = "vendor-id"
    vendorType.EntityData.SegmentPath = "vendor-type" + types.AddKeyToken(vendorType.VendorType, "vendor-type")
    vendorType.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/vendor-ids/vendor-id/" + vendorType.EntityData.SegmentPath
    vendorType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendorType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendorType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendorType.EntityData.Children = types.NewOrderedMap()
    vendorType.EntityData.Children.Append("attribute-name", types.YChild{"AttributeName", nil})
    for i := range vendorType.AttributeName {
        vendorType.EntityData.Children.Append(types.GetSegmentPath(vendorType.AttributeName[i]), types.YChild{"AttributeName", vendorType.AttributeName[i]})
    }
    vendorType.EntityData.Leafs = types.NewOrderedMap()
    vendorType.EntityData.Leafs.Append("vendor-type", types.YLeaf{"VendorType", vendorType.VendorType})

    vendorType.EntityData.YListKeys = []string {"VendorType"}

    return &(vendorType.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName
// Attribute Name of vsa
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Attribute Name of vsa. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    AttributeName interface{}

    // AttributeName of vsa is absent. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent.
    AttributeNameAbsent []*Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent
}

func (attributeName *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName) GetEntityData() *types.CommonEntityData {
    attributeName.EntityData.YFilter = attributeName.YFilter
    attributeName.EntityData.YangName = "attribute-name"
    attributeName.EntityData.BundleName = "cisco_ios_xr"
    attributeName.EntityData.ParentYangName = "vendor-type"
    attributeName.EntityData.SegmentPath = "attribute-name" + types.AddKeyToken(attributeName.AttributeName, "attribute-name")
    attributeName.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/vendor-ids/vendor-id/vendor-type/" + attributeName.EntityData.SegmentPath
    attributeName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeName.EntityData.Children = types.NewOrderedMap()
    attributeName.EntityData.Children.Append("attribute-name-absent", types.YChild{"AttributeNameAbsent", nil})
    for i := range attributeName.AttributeNameAbsent {
        attributeName.EntityData.Children.Append(types.GetSegmentPath(attributeName.AttributeNameAbsent[i]), types.YChild{"AttributeNameAbsent", attributeName.AttributeNameAbsent[i]})
    }
    attributeName.EntityData.Leafs = types.NewOrderedMap()
    attributeName.EntityData.Leafs.Append("attribute-name", types.YLeaf{"AttributeName", attributeName.AttributeName})

    attributeName.EntityData.YListKeys = []string {"AttributeName"}

    return &(attributeName.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent
// AttributeName of vsa is absent
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AttributeName of vsa is absent. The type is
    // interface{} with range: 0..4294967295.
    AttributeNameAbsent interface{}

    // AttributeName of vsa is present. The type is slice of
    // Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent.
    AttributeNamePresent []*Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent
}

func (attributeNameAbsent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent) GetEntityData() *types.CommonEntityData {
    attributeNameAbsent.EntityData.YFilter = attributeNameAbsent.YFilter
    attributeNameAbsent.EntityData.YangName = "attribute-name-absent"
    attributeNameAbsent.EntityData.BundleName = "cisco_ios_xr"
    attributeNameAbsent.EntityData.ParentYangName = "attribute-name"
    attributeNameAbsent.EntityData.SegmentPath = "attribute-name-absent" + types.AddKeyToken(attributeNameAbsent.AttributeNameAbsent, "attribute-name-absent")
    attributeNameAbsent.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/vendor-ids/vendor-id/vendor-type/attribute-name/" + attributeNameAbsent.EntityData.SegmentPath
    attributeNameAbsent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeNameAbsent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeNameAbsent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeNameAbsent.EntityData.Children = types.NewOrderedMap()
    attributeNameAbsent.EntityData.Children.Append("attribute-name-present", types.YChild{"AttributeNamePresent", nil})
    for i := range attributeNameAbsent.AttributeNamePresent {
        attributeNameAbsent.EntityData.Children.Append(types.GetSegmentPath(attributeNameAbsent.AttributeNamePresent[i]), types.YChild{"AttributeNamePresent", attributeNameAbsent.AttributeNamePresent[i]})
    }
    attributeNameAbsent.EntityData.Leafs = types.NewOrderedMap()
    attributeNameAbsent.EntityData.Leafs.Append("attribute-name-absent", types.YLeaf{"AttributeNameAbsent", attributeNameAbsent.AttributeNameAbsent})

    attributeNameAbsent.EntityData.YListKeys = []string {"AttributeNameAbsent"}

    return &(attributeNameAbsent.EntityData)
}

// Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent
// AttributeName of vsa is present
type Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AttributeName of vsa is present. The type is
    // interface{} with range: 0..4294967295.
    AttributeNamePresent interface{}
}

func (attributeNamePresent *Aaa_Radius_Attributes_Attribute_VendorIds_VendorId_VendorType_AttributeName_AttributeNameAbsent_AttributeNamePresent) GetEntityData() *types.CommonEntityData {
    attributeNamePresent.EntityData.YFilter = attributeNamePresent.YFilter
    attributeNamePresent.EntityData.YangName = "attribute-name-present"
    attributeNamePresent.EntityData.BundleName = "cisco_ios_xr"
    attributeNamePresent.EntityData.ParentYangName = "attribute-name-absent"
    attributeNamePresent.EntityData.SegmentPath = "attribute-name-present" + types.AddKeyToken(attributeNamePresent.AttributeNamePresent, "attribute-name-present")
    attributeNamePresent.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/attributes/attribute/vendor-ids/vendor-id/vendor-type/attribute-name/attribute-name-absent/" + attributeNamePresent.EntityData.SegmentPath
    attributeNamePresent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeNamePresent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeNamePresent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeNamePresent.EntityData.Children = types.NewOrderedMap()
    attributeNamePresent.EntityData.Leafs = types.NewOrderedMap()
    attributeNamePresent.EntityData.Leafs.Append("attribute-name-present", types.YLeaf{"AttributeNamePresent", attributeNamePresent.AttributeNamePresent})

    attributeNamePresent.EntityData.YListKeys = []string {"AttributeNamePresent"}

    return &(attributeNamePresent.EntityData)
}

// Aaa_Radius_SourcePort
// Source port
type Aaa_Radius_SourcePort struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable source-port extend . The type is interface{}.
    Extended interface{}
}

func (sourcePort *Aaa_Radius_SourcePort) GetEntityData() *types.CommonEntityData {
    sourcePort.EntityData.YFilter = sourcePort.YFilter
    sourcePort.EntityData.YangName = "source-port"
    sourcePort.EntityData.BundleName = "cisco_ios_xr"
    sourcePort.EntityData.ParentYangName = "radius"
    sourcePort.EntityData.SegmentPath = "source-port"
    sourcePort.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-protocol-radius-cfg:radius/" + sourcePort.EntityData.SegmentPath
    sourcePort.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourcePort.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourcePort.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourcePort.EntityData.Children = types.NewOrderedMap()
    sourcePort.EntityData.Leafs = types.NewOrderedMap()
    sourcePort.EntityData.Leafs.Append("extended", types.YLeaf{"Extended", sourcePort.Extended})

    sourcePort.EntityData.YListKeys = []string {}

    return &(sourcePort.EntityData)
}

// Aaa_Diameter
// Diameter base protocol
type Aaa_Diameter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify interface for source address in DIAMETER packets. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // Start diameter policy-if.
    Gy Aaa_Diameter_Gy

    // Origin sub commands.
    Origin Aaa_Diameter_Origin

    // Start diameter Nas.
    Nas Aaa_Diameter_Nas

    // TLS sub commands.
    DiameterTls Aaa_Diameter_DiameterTls

    // List of diameter peers.
    Peers Aaa_Diameter_Peers

    // Attribute list configuration for test command.
    Diams Aaa_Diameter_Diams

    // Start diameter policy-if.
    Gx Aaa_Diameter_Gx

    // Service sub commands.
    Services Aaa_Diameter_Services

    // Timers used for the peer.
    DiameterTimer Aaa_Diameter_DiameterTimer

    // Vendor specific.
    Vendor Aaa_Diameter_Vendor
}

func (diameter *Aaa_Diameter) GetEntityData() *types.CommonEntityData {
    diameter.EntityData.YFilter = diameter.YFilter
    diameter.EntityData.YangName = "diameter"
    diameter.EntityData.BundleName = "cisco_ios_xr"
    diameter.EntityData.ParentYangName = "aaa"
    diameter.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-diameter-cfg:diameter"
    diameter.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/" + diameter.EntityData.SegmentPath
    diameter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameter.EntityData.Children = types.NewOrderedMap()
    diameter.EntityData.Children.Append("gy", types.YChild{"Gy", &diameter.Gy})
    diameter.EntityData.Children.Append("origin", types.YChild{"Origin", &diameter.Origin})
    diameter.EntityData.Children.Append("nas", types.YChild{"Nas", &diameter.Nas})
    diameter.EntityData.Children.Append("diameter-tls", types.YChild{"DiameterTls", &diameter.DiameterTls})
    diameter.EntityData.Children.Append("peers", types.YChild{"Peers", &diameter.Peers})
    diameter.EntityData.Children.Append("diams", types.YChild{"Diams", &diameter.Diams})
    diameter.EntityData.Children.Append("gx", types.YChild{"Gx", &diameter.Gx})
    diameter.EntityData.Children.Append("services", types.YChild{"Services", &diameter.Services})
    diameter.EntityData.Children.Append("diameter-timer", types.YChild{"DiameterTimer", &diameter.DiameterTimer})
    diameter.EntityData.Children.Append("vendor", types.YChild{"Vendor", &diameter.Vendor})
    diameter.EntityData.Leafs = types.NewOrderedMap()
    diameter.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", diameter.SourceInterface})

    diameter.EntityData.YListKeys = []string {}

    return &(diameter.EntityData)
}

// Aaa_Diameter_Gy
// Start diameter policy-if
type Aaa_Diameter_Gy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set retransmit count. The type is interface{} with range: 1..10.
    Retransmit interface{}

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}

    // Transaction timer value. The type is interface{} with range: 6..1000.
    TxTimer interface{}
}

func (gy *Aaa_Diameter_Gy) GetEntityData() *types.CommonEntityData {
    gy.EntityData.YFilter = gy.YFilter
    gy.EntityData.YangName = "gy"
    gy.EntityData.BundleName = "cisco_ios_xr"
    gy.EntityData.ParentYangName = "diameter"
    gy.EntityData.SegmentPath = "gy"
    gy.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + gy.EntityData.SegmentPath
    gy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gy.EntityData.Children = types.NewOrderedMap()
    gy.EntityData.Leafs = types.NewOrderedMap()
    gy.EntityData.Leafs.Append("retransmit", types.YLeaf{"Retransmit", gy.Retransmit})
    gy.EntityData.Leafs.Append("dest-host", types.YLeaf{"DestHost", gy.DestHost})
    gy.EntityData.Leafs.Append("tx-timer", types.YLeaf{"TxTimer", gy.TxTimer})

    gy.EntityData.YListKeys = []string {}

    return &(gy.EntityData)
}

// Aaa_Diameter_Origin
// Origin sub commands
type Aaa_Diameter_Origin struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Origin Realm String. The type is string.
    Realm interface{}

    // Host name in FQDN format. The type is string.
    Host interface{}
}

func (origin *Aaa_Diameter_Origin) GetEntityData() *types.CommonEntityData {
    origin.EntityData.YFilter = origin.YFilter
    origin.EntityData.YangName = "origin"
    origin.EntityData.BundleName = "cisco_ios_xr"
    origin.EntityData.ParentYangName = "diameter"
    origin.EntityData.SegmentPath = "origin"
    origin.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + origin.EntityData.SegmentPath
    origin.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    origin.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    origin.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    origin.EntityData.Children = types.NewOrderedMap()
    origin.EntityData.Leafs = types.NewOrderedMap()
    origin.EntityData.Leafs.Append("realm", types.YLeaf{"Realm", origin.Realm})
    origin.EntityData.Leafs.Append("host", types.YLeaf{"Host", origin.Host})

    origin.EntityData.YListKeys = []string {}

    return &(origin.EntityData)
}

// Aaa_Diameter_Nas
// Start diameter Nas
type Aaa_Diameter_Nas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}
}

func (nas *Aaa_Diameter_Nas) GetEntityData() *types.CommonEntityData {
    nas.EntityData.YFilter = nas.YFilter
    nas.EntityData.YangName = "nas"
    nas.EntityData.BundleName = "cisco_ios_xr"
    nas.EntityData.ParentYangName = "diameter"
    nas.EntityData.SegmentPath = "nas"
    nas.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + nas.EntityData.SegmentPath
    nas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nas.EntityData.Children = types.NewOrderedMap()
    nas.EntityData.Leafs = types.NewOrderedMap()
    nas.EntityData.Leafs.Append("dest-host", types.YLeaf{"DestHost", nas.DestHost})

    nas.EntityData.YListKeys = []string {}

    return &(nas.EntityData)
}

// Aaa_Diameter_DiameterTls
// TLS sub commands
type Aaa_Diameter_DiameterTls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trustpoint label to be used. The type is string.
    Trustpoint interface{}
}

func (diameterTls *Aaa_Diameter_DiameterTls) GetEntityData() *types.CommonEntityData {
    diameterTls.EntityData.YFilter = diameterTls.YFilter
    diameterTls.EntityData.YangName = "diameter-tls"
    diameterTls.EntityData.BundleName = "cisco_ios_xr"
    diameterTls.EntityData.ParentYangName = "diameter"
    diameterTls.EntityData.SegmentPath = "diameter-tls"
    diameterTls.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + diameterTls.EntityData.SegmentPath
    diameterTls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameterTls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameterTls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameterTls.EntityData.Children = types.NewOrderedMap()
    diameterTls.EntityData.Leafs = types.NewOrderedMap()
    diameterTls.EntityData.Leafs.Append("trustpoint", types.YLeaf{"Trustpoint", diameterTls.Trustpoint})

    diameterTls.EntityData.YListKeys = []string {}

    return &(diameterTls.EntityData)
}

// Aaa_Diameter_Peers
// List of diameter peers
type Aaa_Diameter_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diameter peer instance. The type is slice of Aaa_Diameter_Peers_Peer.
    Peer []*Aaa_Diameter_Peers_Peer
}

func (peers *Aaa_Diameter_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "diameter"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + peers.EntityData.SegmentPath
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// Aaa_Diameter_Peers_Peer
// Diameter peer instance
type Aaa_Diameter_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name for the diameter peer configuration. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    PeerName interface{}

    // Destination host information. The type is string.
    HostDestination interface{}

    // IPv4 address of diameter server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // Realm to which the peer belongs to. The type is string.
    RealmDestination interface{}

    // Specify a Diameter transport protocol. The type is interface{} with range:
    // 1..65535.
    TcpTransport interface{}

    // Specify interface for source address in DIAMETER packets. The type is
    // string with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // IPv6 address of diameter server. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Specify a Diameter security type. The type is interface{} with range: 0..1.
    TlsTransport interface{}

    // VRF the peer belongs to. The type is string.
    VrfIp interface{}

    // Timers used for the peer.
    PeerTimer Aaa_Diameter_Peers_Peer_PeerTimer

    // Peer Type.
    PeerType Aaa_Diameter_Peers_Peer_PeerType
}

func (peer *Aaa_Diameter_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PeerName, "peer-name")
    peer.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/peers/" + peer.EntityData.SegmentPath
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("peer-timer", types.YChild{"PeerTimer", &peer.PeerTimer})
    peer.EntityData.Children.Append("peer-type", types.YChild{"PeerType", &peer.PeerType})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-name", types.YLeaf{"PeerName", peer.PeerName})
    peer.EntityData.Leafs.Append("host-destination", types.YLeaf{"HostDestination", peer.HostDestination})
    peer.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", peer.Ipv4Address})
    peer.EntityData.Leafs.Append("realm-destination", types.YLeaf{"RealmDestination", peer.RealmDestination})
    peer.EntityData.Leafs.Append("tcp-transport", types.YLeaf{"TcpTransport", peer.TcpTransport})
    peer.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", peer.SourceInterface})
    peer.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", peer.Ipv6Address})
    peer.EntityData.Leafs.Append("tls-transport", types.YLeaf{"TlsTransport", peer.TlsTransport})
    peer.EntityData.Leafs.Append("vrf-ip", types.YLeaf{"VrfIp", peer.VrfIp})

    peer.EntityData.YListKeys = []string {"PeerName"}

    return &(peer.EntityData)
}

// Aaa_Diameter_Peers_Peer_PeerTimer
// Timers used for the peer
type Aaa_Diameter_Peers_Peer_PeerTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Transaction interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Connection interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Watchdog interface{}
}

func (peerTimer *Aaa_Diameter_Peers_Peer_PeerTimer) GetEntityData() *types.CommonEntityData {
    peerTimer.EntityData.YFilter = peerTimer.YFilter
    peerTimer.EntityData.YangName = "peer-timer"
    peerTimer.EntityData.BundleName = "cisco_ios_xr"
    peerTimer.EntityData.ParentYangName = "peer"
    peerTimer.EntityData.SegmentPath = "peer-timer"
    peerTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/peers/peer/" + peerTimer.EntityData.SegmentPath
    peerTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerTimer.EntityData.Children = types.NewOrderedMap()
    peerTimer.EntityData.Leafs = types.NewOrderedMap()
    peerTimer.EntityData.Leafs.Append("transaction", types.YLeaf{"Transaction", peerTimer.Transaction})
    peerTimer.EntityData.Leafs.Append("connection", types.YLeaf{"Connection", peerTimer.Connection})
    peerTimer.EntityData.Leafs.Append("watchdog", types.YLeaf{"Watchdog", peerTimer.Watchdog})

    peerTimer.EntityData.YListKeys = []string {}

    return &(peerTimer.EntityData)
}

// Aaa_Diameter_Peers_Peer_PeerType
// Peer Type
type Aaa_Diameter_Peers_Peer_PeerType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enabled or disabled. The type is bool.
    Server interface{}
}

func (peerType *Aaa_Diameter_Peers_Peer_PeerType) GetEntityData() *types.CommonEntityData {
    peerType.EntityData.YFilter = peerType.YFilter
    peerType.EntityData.YangName = "peer-type"
    peerType.EntityData.BundleName = "cisco_ios_xr"
    peerType.EntityData.ParentYangName = "peer"
    peerType.EntityData.SegmentPath = "peer-type"
    peerType.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/peers/peer/" + peerType.EntityData.SegmentPath
    peerType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerType.EntityData.Children = types.NewOrderedMap()
    peerType.EntityData.Leafs = types.NewOrderedMap()
    peerType.EntityData.Leafs.Append("server", types.YLeaf{"Server", peerType.Server})

    peerType.EntityData.YListKeys = []string {}

    return &(peerType.EntityData)
}

// Aaa_Diameter_Diams
// Attribute list configuration for test command
type Aaa_Diameter_Diams struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // attribute list configuration. The type is slice of Aaa_Diameter_Diams_Diam.
    Diam []*Aaa_Diameter_Diams_Diam
}

func (diams *Aaa_Diameter_Diams) GetEntityData() *types.CommonEntityData {
    diams.EntityData.YFilter = diams.YFilter
    diams.EntityData.YangName = "diams"
    diams.EntityData.BundleName = "cisco_ios_xr"
    diams.EntityData.ParentYangName = "diameter"
    diams.EntityData.SegmentPath = "diams"
    diams.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + diams.EntityData.SegmentPath
    diams.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diams.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diams.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diams.EntityData.Children = types.NewOrderedMap()
    diams.EntityData.Children.Append("diam", types.YChild{"Diam", nil})
    for i := range diams.Diam {
        diams.EntityData.Children.Append(types.GetSegmentPath(diams.Diam[i]), types.YChild{"Diam", diams.Diam[i]})
    }
    diams.EntityData.Leafs = types.NewOrderedMap()

    diams.EntityData.YListKeys = []string {}

    return &(diams.EntityData)
}

// Aaa_Diameter_Diams_Diam
// attribute list configuration
type Aaa_Diameter_Diams_Diam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. attribute list number. The type is interface{}
    // with range: 0..99.
    ListId interface{}

    // Specify an attribute definition.
    DiamAttrDefs Aaa_Diameter_Diams_Diam_DiamAttrDefs
}

func (diam *Aaa_Diameter_Diams_Diam) GetEntityData() *types.CommonEntityData {
    diam.EntityData.YFilter = diam.YFilter
    diam.EntityData.YangName = "diam"
    diam.EntityData.BundleName = "cisco_ios_xr"
    diam.EntityData.ParentYangName = "diams"
    diam.EntityData.SegmentPath = "diam" + types.AddKeyToken(diam.ListId, "list-id")
    diam.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/diams/" + diam.EntityData.SegmentPath
    diam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diam.EntityData.Children = types.NewOrderedMap()
    diam.EntityData.Children.Append("diam-attr-defs", types.YChild{"DiamAttrDefs", &diam.DiamAttrDefs})
    diam.EntityData.Leafs = types.NewOrderedMap()
    diam.EntityData.Leafs.Append("list-id", types.YLeaf{"ListId", diam.ListId})

    diam.EntityData.YListKeys = []string {"ListId"}

    return &(diam.EntityData)
}

// Aaa_Diameter_Diams_Diam_DiamAttrDefs
// Specify an attribute definition
type Aaa_Diameter_Diams_Diam_DiamAttrDefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // vendor id. The type is slice of
    // Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef.
    DiamAttrDef []*Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef
}

func (diamAttrDefs *Aaa_Diameter_Diams_Diam_DiamAttrDefs) GetEntityData() *types.CommonEntityData {
    diamAttrDefs.EntityData.YFilter = diamAttrDefs.YFilter
    diamAttrDefs.EntityData.YangName = "diam-attr-defs"
    diamAttrDefs.EntityData.BundleName = "cisco_ios_xr"
    diamAttrDefs.EntityData.ParentYangName = "diam"
    diamAttrDefs.EntityData.SegmentPath = "diam-attr-defs"
    diamAttrDefs.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/diams/diam/" + diamAttrDefs.EntityData.SegmentPath
    diamAttrDefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diamAttrDefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diamAttrDefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diamAttrDefs.EntityData.Children = types.NewOrderedMap()
    diamAttrDefs.EntityData.Children.Append("diam-attr-def", types.YChild{"DiamAttrDef", nil})
    for i := range diamAttrDefs.DiamAttrDef {
        diamAttrDefs.EntityData.Children.Append(types.GetSegmentPath(diamAttrDefs.DiamAttrDef[i]), types.YChild{"DiamAttrDef", diamAttrDefs.DiamAttrDef[i]})
    }
    diamAttrDefs.EntityData.Leafs = types.NewOrderedMap()

    diamAttrDefs.EntityData.YListKeys = []string {}

    return &(diamAttrDefs.EntityData)
}

// Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef
// vendor id
type Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. value for vendor id. The type is interface{} with
    // range: 0..4294967295.
    VendorId interface{}

    // This attribute is a key. enter attribute id. The type is interface{} with
    // range: 1..65535.
    AttributeId interface{}

    // attr subcommands.
    DiamAttrValue Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue
}

func (diamAttrDef *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef) GetEntityData() *types.CommonEntityData {
    diamAttrDef.EntityData.YFilter = diamAttrDef.YFilter
    diamAttrDef.EntityData.YangName = "diam-attr-def"
    diamAttrDef.EntityData.BundleName = "cisco_ios_xr"
    diamAttrDef.EntityData.ParentYangName = "diam-attr-defs"
    diamAttrDef.EntityData.SegmentPath = "diam-attr-def" + types.AddKeyToken(diamAttrDef.VendorId, "vendor-id") + types.AddKeyToken(diamAttrDef.AttributeId, "attribute-id")
    diamAttrDef.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/diams/diam/diam-attr-defs/" + diamAttrDef.EntityData.SegmentPath
    diamAttrDef.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diamAttrDef.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diamAttrDef.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diamAttrDef.EntityData.Children = types.NewOrderedMap()
    diamAttrDef.EntityData.Children.Append("diam-attr-value", types.YChild{"DiamAttrValue", &diamAttrDef.DiamAttrValue})
    diamAttrDef.EntityData.Leafs = types.NewOrderedMap()
    diamAttrDef.EntityData.Leafs.Append("vendor-id", types.YLeaf{"VendorId", diamAttrDef.VendorId})
    diamAttrDef.EntityData.Leafs.Append("attribute-id", types.YLeaf{"AttributeId", diamAttrDef.AttributeId})

    diamAttrDef.EntityData.YListKeys = []string {"VendorId", "AttributeId"}

    return &(diamAttrDef.EntityData)
}

// Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue
// attr subcommands
type Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // String type. The type is string.
    TypeString interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    TypeIpv4Address interface{}

    // Binary type. The type is string.
    TypeBinary interface{}

    // Boolean type. The type is interface{} with range: 0..4294967295.
    TypeBoolean interface{}

    // Enumeration type. The type is interface{} with range: 0..4294967295.
    TypeEnum interface{}

    // Grouped attribute. The type is interface{} with range: 0..99.
    TypeGrouped interface{}

    // Numeric type. The type is interface{} with range: 0..4294967295.
    TypeUlong interface{}

    // Diameter-identity type. The type is string.
    TypeIdentity interface{}

    // Dataypte of attribute. The type is interface{} with range: 0..23.
    DataType interface{}

    // Is mandatory?. The type is interface{} with range: 0..1.
    Mandatory interface{}
}

func (diamAttrValue *Aaa_Diameter_Diams_Diam_DiamAttrDefs_DiamAttrDef_DiamAttrValue) GetEntityData() *types.CommonEntityData {
    diamAttrValue.EntityData.YFilter = diamAttrValue.YFilter
    diamAttrValue.EntityData.YangName = "diam-attr-value"
    diamAttrValue.EntityData.BundleName = "cisco_ios_xr"
    diamAttrValue.EntityData.ParentYangName = "diam-attr-def"
    diamAttrValue.EntityData.SegmentPath = "diam-attr-value"
    diamAttrValue.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/diams/diam/diam-attr-defs/diam-attr-def/" + diamAttrValue.EntityData.SegmentPath
    diamAttrValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diamAttrValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diamAttrValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diamAttrValue.EntityData.Children = types.NewOrderedMap()
    diamAttrValue.EntityData.Leafs = types.NewOrderedMap()
    diamAttrValue.EntityData.Leafs.Append("type-string", types.YLeaf{"TypeString", diamAttrValue.TypeString})
    diamAttrValue.EntityData.Leafs.Append("type-ipv4-address", types.YLeaf{"TypeIpv4Address", diamAttrValue.TypeIpv4Address})
    diamAttrValue.EntityData.Leafs.Append("type-binary", types.YLeaf{"TypeBinary", diamAttrValue.TypeBinary})
    diamAttrValue.EntityData.Leafs.Append("type-boolean", types.YLeaf{"TypeBoolean", diamAttrValue.TypeBoolean})
    diamAttrValue.EntityData.Leafs.Append("type-enum", types.YLeaf{"TypeEnum", diamAttrValue.TypeEnum})
    diamAttrValue.EntityData.Leafs.Append("type-grouped", types.YLeaf{"TypeGrouped", diamAttrValue.TypeGrouped})
    diamAttrValue.EntityData.Leafs.Append("type-ulong", types.YLeaf{"TypeUlong", diamAttrValue.TypeUlong})
    diamAttrValue.EntityData.Leafs.Append("type-identity", types.YLeaf{"TypeIdentity", diamAttrValue.TypeIdentity})
    diamAttrValue.EntityData.Leafs.Append("data-type", types.YLeaf{"DataType", diamAttrValue.DataType})
    diamAttrValue.EntityData.Leafs.Append("mandatory", types.YLeaf{"Mandatory", diamAttrValue.Mandatory})

    diamAttrValue.EntityData.YListKeys = []string {}

    return &(diamAttrValue.EntityData)
}

// Aaa_Diameter_Gx
// Start diameter policy-if
type Aaa_Diameter_Gx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set retransmit count. The type is interface{} with range: 1..10.
    Retransmit interface{}

    // Destination Host name in FQDN format. The type is string.
    DestHost interface{}

    // Transaction timer value. The type is interface{} with range: 6..1000.
    TxTimer interface{}
}

func (gx *Aaa_Diameter_Gx) GetEntityData() *types.CommonEntityData {
    gx.EntityData.YFilter = gx.YFilter
    gx.EntityData.YangName = "gx"
    gx.EntityData.BundleName = "cisco_ios_xr"
    gx.EntityData.ParentYangName = "diameter"
    gx.EntityData.SegmentPath = "gx"
    gx.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + gx.EntityData.SegmentPath
    gx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gx.EntityData.Children = types.NewOrderedMap()
    gx.EntityData.Leafs = types.NewOrderedMap()
    gx.EntityData.Leafs.Append("retransmit", types.YLeaf{"Retransmit", gx.Retransmit})
    gx.EntityData.Leafs.Append("dest-host", types.YLeaf{"DestHost", gx.DestHost})
    gx.EntityData.Leafs.Append("tx-timer", types.YLeaf{"TxTimer", gx.TxTimer})

    gx.EntityData.YListKeys = []string {}

    return &(gx.EntityData)
}

// Aaa_Diameter_Services
// Service sub commands
type Aaa_Diameter_Services struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diameter Service instance. The type is slice of
    // Aaa_Diameter_Services_Service.
    Service []*Aaa_Diameter_Services_Service
}

func (services *Aaa_Diameter_Services) GetEntityData() *types.CommonEntityData {
    services.EntityData.YFilter = services.YFilter
    services.EntityData.YangName = "services"
    services.EntityData.BundleName = "cisco_ios_xr"
    services.EntityData.ParentYangName = "diameter"
    services.EntityData.SegmentPath = "services"
    services.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + services.EntityData.SegmentPath
    services.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    services.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    services.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    services.EntityData.Children = types.NewOrderedMap()
    services.EntityData.Children.Append("service", types.YChild{"Service", nil})
    for i := range services.Service {
        services.EntityData.Children.Append(types.GetSegmentPath(services.Service[i]), types.YChild{"Service", services.Service[i]})
    }
    services.EntityData.Leafs = types.NewOrderedMap()

    services.EntityData.YListKeys = []string {}

    return &(services.EntityData)
}

// Aaa_Diameter_Services_Service
// Diameter Service instance
type Aaa_Diameter_Services_Service struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name for the diameter Service configuration. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ServiceName interface{}

    // Name for the diameter Service Monitoring configuration. The type is string.
    MonitoringKey interface{}
}

func (service *Aaa_Diameter_Services_Service) GetEntityData() *types.CommonEntityData {
    service.EntityData.YFilter = service.YFilter
    service.EntityData.YangName = "service"
    service.EntityData.BundleName = "cisco_ios_xr"
    service.EntityData.ParentYangName = "services"
    service.EntityData.SegmentPath = "service" + types.AddKeyToken(service.ServiceName, "service-name")
    service.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/services/" + service.EntityData.SegmentPath
    service.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    service.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    service.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    service.EntityData.Children = types.NewOrderedMap()
    service.EntityData.Leafs = types.NewOrderedMap()
    service.EntityData.Leafs.Append("service-name", types.YLeaf{"ServiceName", service.ServiceName})
    service.EntityData.Leafs.Append("monitoring-key", types.YLeaf{"MonitoringKey", service.MonitoringKey})

    service.EntityData.YListKeys = []string {"ServiceName"}

    return &(service.EntityData)
}

// Aaa_Diameter_DiameterTimer
// Timers used for the peer
type Aaa_Diameter_DiameterTimer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Transaction interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Connection interface{}

    // Timer value in seconds. The type is interface{} with range: 6..1000.
    Watchdog interface{}
}

func (diameterTimer *Aaa_Diameter_DiameterTimer) GetEntityData() *types.CommonEntityData {
    diameterTimer.EntityData.YFilter = diameterTimer.YFilter
    diameterTimer.EntityData.YangName = "diameter-timer"
    diameterTimer.EntityData.BundleName = "cisco_ios_xr"
    diameterTimer.EntityData.ParentYangName = "diameter"
    diameterTimer.EntityData.SegmentPath = "diameter-timer"
    diameterTimer.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + diameterTimer.EntityData.SegmentPath
    diameterTimer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diameterTimer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diameterTimer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diameterTimer.EntityData.Children = types.NewOrderedMap()
    diameterTimer.EntityData.Leafs = types.NewOrderedMap()
    diameterTimer.EntityData.Leafs.Append("transaction", types.YLeaf{"Transaction", diameterTimer.Transaction})
    diameterTimer.EntityData.Leafs.Append("connection", types.YLeaf{"Connection", diameterTimer.Connection})
    diameterTimer.EntityData.Leafs.Append("watchdog", types.YLeaf{"Watchdog", diameterTimer.Watchdog})

    diameterTimer.EntityData.YListKeys = []string {}

    return &(diameterTimer.EntityData)
}

// Aaa_Diameter_Vendor
// Vendor specific
type Aaa_Diameter_Vendor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Supported vendors.
    Supported Aaa_Diameter_Vendor_Supported
}

func (vendor *Aaa_Diameter_Vendor) GetEntityData() *types.CommonEntityData {
    vendor.EntityData.YFilter = vendor.YFilter
    vendor.EntityData.YangName = "vendor"
    vendor.EntityData.BundleName = "cisco_ios_xr"
    vendor.EntityData.ParentYangName = "diameter"
    vendor.EntityData.SegmentPath = "vendor"
    vendor.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/" + vendor.EntityData.SegmentPath
    vendor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vendor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vendor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vendor.EntityData.Children = types.NewOrderedMap()
    vendor.EntityData.Children.Append("supported", types.YChild{"Supported", &vendor.Supported})
    vendor.EntityData.Leafs = types.NewOrderedMap()

    vendor.EntityData.YListKeys = []string {}

    return &(vendor.EntityData)
}

// Aaa_Diameter_Vendor_Supported
// Supported vendors
type Aaa_Diameter_Vendor_Supported struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cisco attribute support. The type is bool.
    Cisco interface{}

    // 3GPP attribute support. The type is bool.
    Threegpp interface{}

    // Etsi attribute support. The type is bool.
    Etsi interface{}

    // Vodafone attribute support. The type is bool.
    Vodafone interface{}
}

func (supported *Aaa_Diameter_Vendor_Supported) GetEntityData() *types.CommonEntityData {
    supported.EntityData.YFilter = supported.YFilter
    supported.EntityData.YangName = "supported"
    supported.EntityData.BundleName = "cisco_ios_xr"
    supported.EntityData.ParentYangName = "vendor"
    supported.EntityData.SegmentPath = "supported"
    supported.EntityData.AbsolutePath = "Cisco-IOS-XR-aaa-lib-cfg:aaa/Cisco-IOS-XR-aaa-diameter-cfg:diameter/vendor/" + supported.EntityData.SegmentPath
    supported.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    supported.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    supported.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    supported.EntityData.Children = types.NewOrderedMap()
    supported.EntityData.Leafs = types.NewOrderedMap()
    supported.EntityData.Leafs.Append("cisco", types.YLeaf{"Cisco", supported.Cisco})
    supported.EntityData.Leafs.Append("threegpp", types.YLeaf{"Threegpp", supported.Threegpp})
    supported.EntityData.Leafs.Append("etsi", types.YLeaf{"Etsi", supported.Etsi})
    supported.EntityData.Leafs.Append("vodafone", types.YLeaf{"Vodafone", supported.Vodafone})

    supported.EntityData.YListKeys = []string {}

    return &(supported.EntityData)
}

