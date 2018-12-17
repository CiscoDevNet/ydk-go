// This module contains a collection of YANG definitions
// for Cisco IOS-XR action package configuration.
// 
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package spirit_install_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package spirit_install_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-add}", reflect.TypeOf(InstallAdd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-add", reflect.TypeOf(InstallAdd{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-commit}", reflect.TypeOf(InstallCommit{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-commit", reflect.TypeOf(InstallCommit{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-remove}", reflect.TypeOf(InstallRemove{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-remove", reflect.TypeOf(InstallRemove{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-prepare}", reflect.TypeOf(InstallPrepare{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-prepare", reflect.TypeOf(InstallPrepare{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-activate}", reflect.TypeOf(InstallActivate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-activate", reflect.TypeOf(InstallActivate{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-deactivate}", reflect.TypeOf(InstallDeactivate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-deactivate", reflect.TypeOf(InstallDeactivate{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-extract}", reflect.TypeOf(InstallExtract{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-extract", reflect.TypeOf(InstallExtract{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-verify}", reflect.TypeOf(InstallVerify{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-verify", reflect.TypeOf(InstallVerify{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-spirit-install-act install-update}", reflect.TypeOf(InstallUpdate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-spirit-install-act:install-update", reflect.TypeOf(InstallUpdate{}))
}

// InstallAdd
// Cli-command install add source
type InstallAdd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallAdd_Input
}

func (installAdd *InstallAdd) GetEntityData() *types.CommonEntityData {
    installAdd.EntityData.YFilter = installAdd.YFilter
    installAdd.EntityData.YangName = "install-add"
    installAdd.EntityData.BundleName = "cisco_ios_xr"
    installAdd.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installAdd.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-add"
    installAdd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installAdd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installAdd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installAdd.EntityData.Children = types.NewOrderedMap()
    installAdd.EntityData.Children.Append("input", types.YChild{"Input", &installAdd.Input})
    installAdd.EntityData.Leafs = types.NewOrderedMap()

    installAdd.EntityData.YListKeys = []string {}

    return &(installAdd.EntityData)
}

// InstallAdd_Input
type InstallAdd_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (input *InstallAdd_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-add"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", input.Packagepath})
    input.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", input.Packagename})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallCommit
// Cli-command install commit
type InstallCommit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallCommit_Input
}

func (installCommit *InstallCommit) GetEntityData() *types.CommonEntityData {
    installCommit.EntityData.YFilter = installCommit.YFilter
    installCommit.EntityData.YangName = "install-commit"
    installCommit.EntityData.BundleName = "cisco_ios_xr"
    installCommit.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installCommit.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-commit"
    installCommit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installCommit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installCommit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installCommit.EntityData.Children = types.NewOrderedMap()
    installCommit.EntityData.Children.Append("input", types.YChild{"Input", &installCommit.Input})
    installCommit.EntityData.Leafs = types.NewOrderedMap()

    installCommit.EntityData.YListKeys = []string {}

    return &(installCommit.EntityData)
}

// InstallCommit_Input
type InstallCommit_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // commit packages in the system. The type is interface{}.
    Sdr interface{}
}

func (input *InstallCommit_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-commit"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("sdr", types.YLeaf{"Sdr", input.Sdr})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallRemove
// Cli-command remove
type InstallRemove struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallRemove_Input
}

func (installRemove *InstallRemove) GetEntityData() *types.CommonEntityData {
    installRemove.EntityData.YFilter = installRemove.YFilter
    installRemove.EntityData.YangName = "install-remove"
    installRemove.EntityData.BundleName = "cisco_ios_xr"
    installRemove.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installRemove.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-remove"
    installRemove.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installRemove.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installRemove.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installRemove.EntityData.Children = types.NewOrderedMap()
    installRemove.EntityData.Children.Append("input", types.YChild{"Input", &installRemove.Input})
    installRemove.EntityData.Leafs = types.NewOrderedMap()

    installRemove.EntityData.YListKeys = []string {}

    return &(installRemove.EntityData)
}

// InstallRemove_Input
type InstallRemove_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remove inactive packages from XR repo. The type is interface{}.
    Inactive interface{}

    // Remove inactive packages from Host,Sysadmin and XR repo. The type is
    // interface{}.
    Inactiveall interface{}

    
    Packages InstallRemove_Input_Packages

    
    Ids InstallRemove_Input_Ids
}

func (input *InstallRemove_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-remove"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("packages", types.YChild{"Packages", &input.Packages})
    input.EntityData.Children.Append("ids", types.YChild{"Ids", &input.Ids})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("inactive", types.YLeaf{"Inactive", input.Inactive})
    input.EntityData.Leafs.Append("inactiveall", types.YLeaf{"Inactiveall", input.Inactiveall})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallRemove_Input_Packages
type InstallRemove_Input_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}
}

func (packages *InstallRemove_Input_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "input"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Leafs = types.NewOrderedMap()
    packages.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", packages.Packagename})

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// InstallRemove_Input_Ids
type InstallRemove_Input_Ids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IdNo []interface{}
}

func (ids *InstallRemove_Input_Ids) GetEntityData() *types.CommonEntityData {
    ids.EntityData.YFilter = ids.YFilter
    ids.EntityData.YangName = "ids"
    ids.EntityData.BundleName = "cisco_ios_xr"
    ids.EntityData.ParentYangName = "input"
    ids.EntityData.SegmentPath = "ids"
    ids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ids.EntityData.Children = types.NewOrderedMap()
    ids.EntityData.Leafs = types.NewOrderedMap()
    ids.EntityData.Leafs.Append("id-no", types.YLeaf{"IdNo", ids.IdNo})

    ids.EntityData.YListKeys = []string {}

    return &(ids.EntityData)
}

// InstallPrepare
// Cli-command prepare
type InstallPrepare struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallPrepare_Input
}

func (installPrepare *InstallPrepare) GetEntityData() *types.CommonEntityData {
    installPrepare.EntityData.YFilter = installPrepare.YFilter
    installPrepare.EntityData.YangName = "install-prepare"
    installPrepare.EntityData.BundleName = "cisco_ios_xr"
    installPrepare.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installPrepare.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-prepare"
    installPrepare.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installPrepare.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installPrepare.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installPrepare.EntityData.Children = types.NewOrderedMap()
    installPrepare.EntityData.Children.Append("input", types.YChild{"Input", &installPrepare.Input})
    installPrepare.EntityData.Leafs = types.NewOrderedMap()

    installPrepare.EntityData.YListKeys = []string {}

    return &(installPrepare.EntityData)
}

// InstallPrepare_Input
type InstallPrepare_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clean the prepared packages. The type is interface{}.
    Clean interface{}

    
    Packages InstallPrepare_Input_Packages

    
    Ids InstallPrepare_Input_Ids

    
    PrepareForce InstallPrepare_Input_PrepareForce
}

func (input *InstallPrepare_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-prepare"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("packages", types.YChild{"Packages", &input.Packages})
    input.EntityData.Children.Append("ids", types.YChild{"Ids", &input.Ids})
    input.EntityData.Children.Append("prepare-force", types.YChild{"PrepareForce", &input.PrepareForce})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("clean", types.YLeaf{"Clean", input.Clean})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallPrepare_Input_Packages
type InstallPrepare_Input_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}
}

func (packages *InstallPrepare_Input_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "input"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Leafs = types.NewOrderedMap()
    packages.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", packages.Packagename})

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// InstallPrepare_Input_Ids
type InstallPrepare_Input_Ids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IdNo []interface{}
}

func (ids *InstallPrepare_Input_Ids) GetEntityData() *types.CommonEntityData {
    ids.EntityData.YFilter = ids.YFilter
    ids.EntityData.YangName = "ids"
    ids.EntityData.BundleName = "cisco_ios_xr"
    ids.EntityData.ParentYangName = "input"
    ids.EntityData.SegmentPath = "ids"
    ids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ids.EntityData.Children = types.NewOrderedMap()
    ids.EntityData.Leafs = types.NewOrderedMap()
    ids.EntityData.Leafs.Append("id-no", types.YLeaf{"IdNo", ids.IdNo})

    ids.EntityData.YListKeys = []string {}

    return &(ids.EntityData)
}

// InstallPrepare_Input_PrepareForce
type InstallPrepare_Input_PrepareForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (prepareForce *InstallPrepare_Input_PrepareForce) GetEntityData() *types.CommonEntityData {
    prepareForce.EntityData.YFilter = prepareForce.YFilter
    prepareForce.EntityData.YangName = "prepare-force"
    prepareForce.EntityData.BundleName = "cisco_ios_xr"
    prepareForce.EntityData.ParentYangName = "input"
    prepareForce.EntityData.SegmentPath = "prepare-force"
    prepareForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepareForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepareForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepareForce.EntityData.Children = types.NewOrderedMap()
    prepareForce.EntityData.Leafs = types.NewOrderedMap()
    prepareForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", prepareForce.Packagename})
    prepareForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", prepareForce.Ids})

    prepareForce.EntityData.YListKeys = []string {}

    return &(prepareForce.EntityData)
}

// InstallActivate
// Cli-command activate
type InstallActivate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallActivate_Input
}

func (installActivate *InstallActivate) GetEntityData() *types.CommonEntityData {
    installActivate.EntityData.YFilter = installActivate.YFilter
    installActivate.EntityData.YangName = "install-activate"
    installActivate.EntityData.BundleName = "cisco_ios_xr"
    installActivate.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installActivate.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-activate"
    installActivate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installActivate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installActivate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installActivate.EntityData.Children = types.NewOrderedMap()
    installActivate.EntityData.Children.Append("input", types.YChild{"Input", &installActivate.Input})
    installActivate.EntityData.Leafs = types.NewOrderedMap()

    installActivate.EntityData.YListKeys = []string {}

    return &(installActivate.EntityData)
}

// InstallActivate_Input
type InstallActivate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Activate the prepared package. The type is interface{}.
    ActivatePreparedPkg interface{}

    // Activate the prepared package with force option. The type is interface{}.
    Force interface{}

    // Activate the prepared package with reload option. The type is interface{}.
    ActivateReload interface{}

    // Activate the prepared package with reload force option. The type is
    // interface{}.
    ActivateReloadForce interface{}

    // Activate the prepared package with warm option. The type is interface{}.
    ActivateWarmPreparedPkg interface{}

    // Activate the prepared package with warm force option. The type is
    // interface{}.
    ActivateWarmForcePreparedPkg interface{}

    // Activate the prepared package with warm replace option. The type is
    // interface{}.
    ActivateWarmReplacePreparedPkg interface{}

    // Activate the prepared package with warm force replace option. The type is
    // interface{}.
    ActivateWarmForceReplacePreparedPkg interface{}

    
    Warm InstallActivate_Input_Warm

    
    WarmForce InstallActivate_Input_WarmForce

    
    WarmReplace InstallActivate_Input_WarmReplace

    
    WarmReplaceForce InstallActivate_Input_WarmReplaceForce

    
    Reload InstallActivate_Input_Reload

    
    ReloadForce InstallActivate_Input_ReloadForce

    
    Replace InstallActivate_Input_Replace

    
    ReplaceForce InstallActivate_Input_ReplaceForce

    
    ActivateForce InstallActivate_Input_ActivateForce

    
    Packages InstallActivate_Input_Packages

    
    Ids InstallActivate_Input_Ids
}

func (input *InstallActivate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-activate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("warm", types.YChild{"Warm", &input.Warm})
    input.EntityData.Children.Append("warm-force", types.YChild{"WarmForce", &input.WarmForce})
    input.EntityData.Children.Append("warm-replace", types.YChild{"WarmReplace", &input.WarmReplace})
    input.EntityData.Children.Append("warm-replace-force", types.YChild{"WarmReplaceForce", &input.WarmReplaceForce})
    input.EntityData.Children.Append("reload", types.YChild{"Reload", &input.Reload})
    input.EntityData.Children.Append("reload-force", types.YChild{"ReloadForce", &input.ReloadForce})
    input.EntityData.Children.Append("replace", types.YChild{"Replace", &input.Replace})
    input.EntityData.Children.Append("replace-force", types.YChild{"ReplaceForce", &input.ReplaceForce})
    input.EntityData.Children.Append("activate-force", types.YChild{"ActivateForce", &input.ActivateForce})
    input.EntityData.Children.Append("packages", types.YChild{"Packages", &input.Packages})
    input.EntityData.Children.Append("ids", types.YChild{"Ids", &input.Ids})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("activate_prepared_pkg", types.YLeaf{"ActivatePreparedPkg", input.ActivatePreparedPkg})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})
    input.EntityData.Leafs.Append("activate-reload", types.YLeaf{"ActivateReload", input.ActivateReload})
    input.EntityData.Leafs.Append("activate-reload-force", types.YLeaf{"ActivateReloadForce", input.ActivateReloadForce})
    input.EntityData.Leafs.Append("activate_warm_prepared_pkg", types.YLeaf{"ActivateWarmPreparedPkg", input.ActivateWarmPreparedPkg})
    input.EntityData.Leafs.Append("activate_warm_force_prepared_pkg", types.YLeaf{"ActivateWarmForcePreparedPkg", input.ActivateWarmForcePreparedPkg})
    input.EntityData.Leafs.Append("activate_warm_replace_prepared_pkg", types.YLeaf{"ActivateWarmReplacePreparedPkg", input.ActivateWarmReplacePreparedPkg})
    input.EntityData.Leafs.Append("activate_warm_force_replace_prepared_pkg", types.YLeaf{"ActivateWarmForceReplacePreparedPkg", input.ActivateWarmForceReplacePreparedPkg})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallActivate_Input_Warm
type InstallActivate_Input_Warm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (warm *InstallActivate_Input_Warm) GetEntityData() *types.CommonEntityData {
    warm.EntityData.YFilter = warm.YFilter
    warm.EntityData.YangName = "warm"
    warm.EntityData.BundleName = "cisco_ios_xr"
    warm.EntityData.ParentYangName = "input"
    warm.EntityData.SegmentPath = "warm"
    warm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warm.EntityData.Children = types.NewOrderedMap()
    warm.EntityData.Leafs = types.NewOrderedMap()
    warm.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warm.Packagename})
    warm.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", warm.Ids})

    warm.EntityData.YListKeys = []string {}

    return &(warm.EntityData)
}

// InstallActivate_Input_WarmForce
type InstallActivate_Input_WarmForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (warmForce *InstallActivate_Input_WarmForce) GetEntityData() *types.CommonEntityData {
    warmForce.EntityData.YFilter = warmForce.YFilter
    warmForce.EntityData.YangName = "warm-force"
    warmForce.EntityData.BundleName = "cisco_ios_xr"
    warmForce.EntityData.ParentYangName = "input"
    warmForce.EntityData.SegmentPath = "warm-force"
    warmForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmForce.EntityData.Children = types.NewOrderedMap()
    warmForce.EntityData.Leafs = types.NewOrderedMap()
    warmForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmForce.Packagename})
    warmForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", warmForce.Ids})

    warmForce.EntityData.YListKeys = []string {}

    return &(warmForce.EntityData)
}

// InstallActivate_Input_WarmReplace
type InstallActivate_Input_WarmReplace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (warmReplace *InstallActivate_Input_WarmReplace) GetEntityData() *types.CommonEntityData {
    warmReplace.EntityData.YFilter = warmReplace.YFilter
    warmReplace.EntityData.YangName = "warm-replace"
    warmReplace.EntityData.BundleName = "cisco_ios_xr"
    warmReplace.EntityData.ParentYangName = "input"
    warmReplace.EntityData.SegmentPath = "warm-replace"
    warmReplace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmReplace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmReplace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmReplace.EntityData.Children = types.NewOrderedMap()
    warmReplace.EntityData.Leafs = types.NewOrderedMap()
    warmReplace.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmReplace.Packagename})
    warmReplace.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", warmReplace.Ids})

    warmReplace.EntityData.YListKeys = []string {}

    return &(warmReplace.EntityData)
}

// InstallActivate_Input_WarmReplaceForce
type InstallActivate_Input_WarmReplaceForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (warmReplaceForce *InstallActivate_Input_WarmReplaceForce) GetEntityData() *types.CommonEntityData {
    warmReplaceForce.EntityData.YFilter = warmReplaceForce.YFilter
    warmReplaceForce.EntityData.YangName = "warm-replace-force"
    warmReplaceForce.EntityData.BundleName = "cisco_ios_xr"
    warmReplaceForce.EntityData.ParentYangName = "input"
    warmReplaceForce.EntityData.SegmentPath = "warm-replace-force"
    warmReplaceForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmReplaceForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmReplaceForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmReplaceForce.EntityData.Children = types.NewOrderedMap()
    warmReplaceForce.EntityData.Leafs = types.NewOrderedMap()
    warmReplaceForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmReplaceForce.Packagename})
    warmReplaceForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", warmReplaceForce.Ids})

    warmReplaceForce.EntityData.YListKeys = []string {}

    return &(warmReplaceForce.EntityData)
}

// InstallActivate_Input_Reload
type InstallActivate_Input_Reload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (reload *InstallActivate_Input_Reload) GetEntityData() *types.CommonEntityData {
    reload.EntityData.YFilter = reload.YFilter
    reload.EntityData.YangName = "reload"
    reload.EntityData.BundleName = "cisco_ios_xr"
    reload.EntityData.ParentYangName = "input"
    reload.EntityData.SegmentPath = "reload"
    reload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reload.EntityData.Children = types.NewOrderedMap()
    reload.EntityData.Leafs = types.NewOrderedMap()
    reload.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", reload.Packagename})
    reload.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", reload.Ids})

    reload.EntityData.YListKeys = []string {}

    return &(reload.EntityData)
}

// InstallActivate_Input_ReloadForce
type InstallActivate_Input_ReloadForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (reloadForce *InstallActivate_Input_ReloadForce) GetEntityData() *types.CommonEntityData {
    reloadForce.EntityData.YFilter = reloadForce.YFilter
    reloadForce.EntityData.YangName = "reload-force"
    reloadForce.EntityData.BundleName = "cisco_ios_xr"
    reloadForce.EntityData.ParentYangName = "input"
    reloadForce.EntityData.SegmentPath = "reload-force"
    reloadForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadForce.EntityData.Children = types.NewOrderedMap()
    reloadForce.EntityData.Leafs = types.NewOrderedMap()
    reloadForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", reloadForce.Packagename})
    reloadForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", reloadForce.Ids})

    reloadForce.EntityData.YListKeys = []string {}

    return &(reloadForce.EntityData)
}

// InstallActivate_Input_Replace
type InstallActivate_Input_Replace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (replace *InstallActivate_Input_Replace) GetEntityData() *types.CommonEntityData {
    replace.EntityData.YFilter = replace.YFilter
    replace.EntityData.YangName = "replace"
    replace.EntityData.BundleName = "cisco_ios_xr"
    replace.EntityData.ParentYangName = "input"
    replace.EntityData.SegmentPath = "replace"
    replace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replace.EntityData.Children = types.NewOrderedMap()
    replace.EntityData.Leafs = types.NewOrderedMap()
    replace.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replace.Packagename})
    replace.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", replace.Ids})

    replace.EntityData.YListKeys = []string {}

    return &(replace.EntityData)
}

// InstallActivate_Input_ReplaceForce
type InstallActivate_Input_ReplaceForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (replaceForce *InstallActivate_Input_ReplaceForce) GetEntityData() *types.CommonEntityData {
    replaceForce.EntityData.YFilter = replaceForce.YFilter
    replaceForce.EntityData.YangName = "replace-force"
    replaceForce.EntityData.BundleName = "cisco_ios_xr"
    replaceForce.EntityData.ParentYangName = "input"
    replaceForce.EntityData.SegmentPath = "replace-force"
    replaceForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaceForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaceForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaceForce.EntityData.Children = types.NewOrderedMap()
    replaceForce.EntityData.Leafs = types.NewOrderedMap()
    replaceForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replaceForce.Packagename})
    replaceForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", replaceForce.Ids})

    replaceForce.EntityData.YListKeys = []string {}

    return &(replaceForce.EntityData)
}

// InstallActivate_Input_ActivateForce
type InstallActivate_Input_ActivateForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (activateForce *InstallActivate_Input_ActivateForce) GetEntityData() *types.CommonEntityData {
    activateForce.EntityData.YFilter = activateForce.YFilter
    activateForce.EntityData.YangName = "activate-force"
    activateForce.EntityData.BundleName = "cisco_ios_xr"
    activateForce.EntityData.ParentYangName = "input"
    activateForce.EntityData.SegmentPath = "activate-force"
    activateForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activateForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activateForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activateForce.EntityData.Children = types.NewOrderedMap()
    activateForce.EntityData.Leafs = types.NewOrderedMap()
    activateForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", activateForce.Packagename})
    activateForce.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", activateForce.Ids})

    activateForce.EntityData.YListKeys = []string {}

    return &(activateForce.EntityData)
}

// InstallActivate_Input_Packages
type InstallActivate_Input_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}
}

func (packages *InstallActivate_Input_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "input"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Leafs = types.NewOrderedMap()
    packages.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", packages.Packagename})

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// InstallActivate_Input_Ids
type InstallActivate_Input_Ids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IdNo []interface{}
}

func (ids *InstallActivate_Input_Ids) GetEntityData() *types.CommonEntityData {
    ids.EntityData.YFilter = ids.YFilter
    ids.EntityData.YangName = "ids"
    ids.EntityData.BundleName = "cisco_ios_xr"
    ids.EntityData.ParentYangName = "input"
    ids.EntityData.SegmentPath = "ids"
    ids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ids.EntityData.Children = types.NewOrderedMap()
    ids.EntityData.Leafs = types.NewOrderedMap()
    ids.EntityData.Leafs.Append("id-no", types.YLeaf{"IdNo", ids.IdNo})

    ids.EntityData.YListKeys = []string {}

    return &(ids.EntityData)
}

// InstallDeactivate
// Cli-command deactivate
type InstallDeactivate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallDeactivate_Input
}

func (installDeactivate *InstallDeactivate) GetEntityData() *types.CommonEntityData {
    installDeactivate.EntityData.YFilter = installDeactivate.YFilter
    installDeactivate.EntityData.YangName = "install-deactivate"
    installDeactivate.EntityData.BundleName = "cisco_ios_xr"
    installDeactivate.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installDeactivate.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-deactivate"
    installDeactivate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installDeactivate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installDeactivate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installDeactivate.EntityData.Children = types.NewOrderedMap()
    installDeactivate.EntityData.Children.Append("input", types.YChild{"Input", &installDeactivate.Input})
    installDeactivate.EntityData.Leafs = types.NewOrderedMap()

    installDeactivate.EntityData.YListKeys = []string {}

    return &(installDeactivate.EntityData)
}

// InstallDeactivate_Input
type InstallDeactivate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Deactivate all superseded packages. The type is interface{}.
    Superseded interface{}

    
    Reload InstallDeactivate_Input_Reload

    
    Packages InstallDeactivate_Input_Packages

    
    Ids InstallDeactivate_Input_Ids
}

func (input *InstallDeactivate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-deactivate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("reload", types.YChild{"Reload", &input.Reload})
    input.EntityData.Children.Append("packages", types.YChild{"Packages", &input.Packages})
    input.EntityData.Children.Append("ids", types.YChild{"Ids", &input.Ids})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("superseded", types.YLeaf{"Superseded", input.Superseded})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallDeactivate_Input_Reload
type InstallDeactivate_Input_Reload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}

    // The type is slice of string.
    Ids []interface{}
}

func (reload *InstallDeactivate_Input_Reload) GetEntityData() *types.CommonEntityData {
    reload.EntityData.YFilter = reload.YFilter
    reload.EntityData.YangName = "reload"
    reload.EntityData.BundleName = "cisco_ios_xr"
    reload.EntityData.ParentYangName = "input"
    reload.EntityData.SegmentPath = "reload"
    reload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reload.EntityData.Children = types.NewOrderedMap()
    reload.EntityData.Leafs = types.NewOrderedMap()
    reload.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", reload.Packagename})
    reload.EntityData.Leafs.Append("ids", types.YLeaf{"Ids", reload.Ids})

    reload.EntityData.YListKeys = []string {}

    return &(reload.EntityData)
}

// InstallDeactivate_Input_Packages
type InstallDeactivate_Input_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}
}

func (packages *InstallDeactivate_Input_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "input"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Leafs = types.NewOrderedMap()
    packages.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", packages.Packagename})

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// InstallDeactivate_Input_Ids
type InstallDeactivate_Input_Ids struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    IdNo []interface{}
}

func (ids *InstallDeactivate_Input_Ids) GetEntityData() *types.CommonEntityData {
    ids.EntityData.YFilter = ids.YFilter
    ids.EntityData.YangName = "ids"
    ids.EntityData.BundleName = "cisco_ios_xr"
    ids.EntityData.ParentYangName = "input"
    ids.EntityData.SegmentPath = "ids"
    ids.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ids.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ids.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ids.EntityData.Children = types.NewOrderedMap()
    ids.EntityData.Leafs = types.NewOrderedMap()
    ids.EntityData.Leafs.Append("id-no", types.YLeaf{"IdNo", ids.IdNo})

    ids.EntityData.YListKeys = []string {}

    return &(ids.EntityData)
}

// InstallExtract
// Cli-command extract
type InstallExtract struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallExtract_Input
}

func (installExtract *InstallExtract) GetEntityData() *types.CommonEntityData {
    installExtract.EntityData.YFilter = installExtract.YFilter
    installExtract.EntityData.YangName = "install-extract"
    installExtract.EntityData.BundleName = "cisco_ios_xr"
    installExtract.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installExtract.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-extract"
    installExtract.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installExtract.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installExtract.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installExtract.EntityData.Children = types.NewOrderedMap()
    installExtract.EntityData.Children.Append("input", types.YChild{"Input", &installExtract.Input})
    installExtract.EntityData.Leafs = types.NewOrderedMap()

    installExtract.EntityData.YListKeys = []string {}

    return &(installExtract.EntityData)
}

// InstallExtract_Input
type InstallExtract_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Packages InstallExtract_Input_Packages
}

func (input *InstallExtract_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-extract"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("packages", types.YChild{"Packages", &input.Packages})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallExtract_Input_Packages
type InstallExtract_Input_Packages struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Packagename []interface{}
}

func (packages *InstallExtract_Input_Packages) GetEntityData() *types.CommonEntityData {
    packages.EntityData.YFilter = packages.YFilter
    packages.EntityData.YangName = "packages"
    packages.EntityData.BundleName = "cisco_ios_xr"
    packages.EntityData.ParentYangName = "input"
    packages.EntityData.SegmentPath = "packages"
    packages.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packages.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packages.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packages.EntityData.Children = types.NewOrderedMap()
    packages.EntityData.Leafs = types.NewOrderedMap()
    packages.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", packages.Packagename})

    packages.EntityData.YListKeys = []string {}

    return &(packages.EntityData)
}

// InstallVerify
// Cli-command install verify packages
type InstallVerify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallVerify_Input
}

func (installVerify *InstallVerify) GetEntityData() *types.CommonEntityData {
    installVerify.EntityData.YFilter = installVerify.YFilter
    installVerify.EntityData.YangName = "install-verify"
    installVerify.EntityData.BundleName = "cisco_ios_xr"
    installVerify.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installVerify.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-verify"
    installVerify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installVerify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installVerify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installVerify.EntityData.Children = types.NewOrderedMap()
    installVerify.EntityData.Children.Append("input", types.YChild{"Input", &installVerify.Input})
    installVerify.EntityData.Leafs = types.NewOrderedMap()

    installVerify.EntityData.YListKeys = []string {}

    return &(installVerify.EntityData)
}

// InstallVerify_Input
type InstallVerify_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of string.
    Location []interface{}
}

func (input *InstallVerify_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-verify"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallUpdate
// Cli-command install update
type InstallUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input InstallUpdate_Input
}

func (installUpdate *InstallUpdate) GetEntityData() *types.CommonEntityData {
    installUpdate.EntityData.YFilter = installUpdate.YFilter
    installUpdate.EntityData.YangName = "install-update"
    installUpdate.EntityData.BundleName = "cisco_ios_xr"
    installUpdate.EntityData.ParentYangName = "Cisco-IOS-XR-spirit-install-act"
    installUpdate.EntityData.SegmentPath = "Cisco-IOS-XR-spirit-install-act:install-update"
    installUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    installUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    installUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    installUpdate.EntityData.Children = types.NewOrderedMap()
    installUpdate.EntityData.Children.Append("input", types.YChild{"Input", &installUpdate.Input})
    installUpdate.EntityData.Leafs = types.NewOrderedMap()

    installUpdate.EntityData.YListKeys = []string {}

    return &(installUpdate.EntityData)
}

// InstallUpdate_Input
type InstallUpdate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}

    
    Warm InstallUpdate_Input_Warm

    
    WarmForce InstallUpdate_Input_WarmForce

    
    WarmReplace InstallUpdate_Input_WarmReplace

    
    WarmReplaceForce InstallUpdate_Input_WarmReplaceForce

    
    Force InstallUpdate_Input_Force

    
    Replace InstallUpdate_Input_Replace

    
    ReplaceForce InstallUpdate_Input_ReplaceForce

    
    ReplaceCommit InstallUpdate_Input_ReplaceCommit

    
    ReplaceCommitForce InstallUpdate_Input_ReplaceCommitForce
}

func (input *InstallUpdate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "install-update"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("warm", types.YChild{"Warm", &input.Warm})
    input.EntityData.Children.Append("warm-force", types.YChild{"WarmForce", &input.WarmForce})
    input.EntityData.Children.Append("warm-replace", types.YChild{"WarmReplace", &input.WarmReplace})
    input.EntityData.Children.Append("warm-replace-force", types.YChild{"WarmReplaceForce", &input.WarmReplaceForce})
    input.EntityData.Children.Append("force", types.YChild{"Force", &input.Force})
    input.EntityData.Children.Append("replace", types.YChild{"Replace", &input.Replace})
    input.EntityData.Children.Append("replace-force", types.YChild{"ReplaceForce", &input.ReplaceForce})
    input.EntityData.Children.Append("replace-commit", types.YChild{"ReplaceCommit", &input.ReplaceCommit})
    input.EntityData.Children.Append("replace-commit-force", types.YChild{"ReplaceCommitForce", &input.ReplaceCommitForce})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", input.Packagepath})
    input.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", input.Packagename})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// InstallUpdate_Input_Warm
type InstallUpdate_Input_Warm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (warm *InstallUpdate_Input_Warm) GetEntityData() *types.CommonEntityData {
    warm.EntityData.YFilter = warm.YFilter
    warm.EntityData.YangName = "warm"
    warm.EntityData.BundleName = "cisco_ios_xr"
    warm.EntityData.ParentYangName = "input"
    warm.EntityData.SegmentPath = "warm"
    warm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warm.EntityData.Children = types.NewOrderedMap()
    warm.EntityData.Leafs = types.NewOrderedMap()
    warm.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", warm.Packagepath})
    warm.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warm.Packagename})

    warm.EntityData.YListKeys = []string {}

    return &(warm.EntityData)
}

// InstallUpdate_Input_WarmForce
type InstallUpdate_Input_WarmForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (warmForce *InstallUpdate_Input_WarmForce) GetEntityData() *types.CommonEntityData {
    warmForce.EntityData.YFilter = warmForce.YFilter
    warmForce.EntityData.YangName = "warm-force"
    warmForce.EntityData.BundleName = "cisco_ios_xr"
    warmForce.EntityData.ParentYangName = "input"
    warmForce.EntityData.SegmentPath = "warm-force"
    warmForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmForce.EntityData.Children = types.NewOrderedMap()
    warmForce.EntityData.Leafs = types.NewOrderedMap()
    warmForce.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", warmForce.Packagepath})
    warmForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmForce.Packagename})

    warmForce.EntityData.YListKeys = []string {}

    return &(warmForce.EntityData)
}

// InstallUpdate_Input_WarmReplace
type InstallUpdate_Input_WarmReplace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (warmReplace *InstallUpdate_Input_WarmReplace) GetEntityData() *types.CommonEntityData {
    warmReplace.EntityData.YFilter = warmReplace.YFilter
    warmReplace.EntityData.YangName = "warm-replace"
    warmReplace.EntityData.BundleName = "cisco_ios_xr"
    warmReplace.EntityData.ParentYangName = "input"
    warmReplace.EntityData.SegmentPath = "warm-replace"
    warmReplace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmReplace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmReplace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmReplace.EntityData.Children = types.NewOrderedMap()
    warmReplace.EntityData.Leafs = types.NewOrderedMap()
    warmReplace.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", warmReplace.Packagepath})
    warmReplace.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmReplace.Packagename})

    warmReplace.EntityData.YListKeys = []string {}

    return &(warmReplace.EntityData)
}

// InstallUpdate_Input_WarmReplaceForce
type InstallUpdate_Input_WarmReplaceForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (warmReplaceForce *InstallUpdate_Input_WarmReplaceForce) GetEntityData() *types.CommonEntityData {
    warmReplaceForce.EntityData.YFilter = warmReplaceForce.YFilter
    warmReplaceForce.EntityData.YangName = "warm-replace-force"
    warmReplaceForce.EntityData.BundleName = "cisco_ios_xr"
    warmReplaceForce.EntityData.ParentYangName = "input"
    warmReplaceForce.EntityData.SegmentPath = "warm-replace-force"
    warmReplaceForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    warmReplaceForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    warmReplaceForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    warmReplaceForce.EntityData.Children = types.NewOrderedMap()
    warmReplaceForce.EntityData.Leafs = types.NewOrderedMap()
    warmReplaceForce.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", warmReplaceForce.Packagepath})
    warmReplaceForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", warmReplaceForce.Packagename})

    warmReplaceForce.EntityData.YListKeys = []string {}

    return &(warmReplaceForce.EntityData)
}

// InstallUpdate_Input_Force
type InstallUpdate_Input_Force struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (force *InstallUpdate_Input_Force) GetEntityData() *types.CommonEntityData {
    force.EntityData.YFilter = force.YFilter
    force.EntityData.YangName = "force"
    force.EntityData.BundleName = "cisco_ios_xr"
    force.EntityData.ParentYangName = "input"
    force.EntityData.SegmentPath = "force"
    force.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    force.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    force.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    force.EntityData.Children = types.NewOrderedMap()
    force.EntityData.Leafs = types.NewOrderedMap()
    force.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", force.Packagepath})
    force.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", force.Packagename})

    force.EntityData.YListKeys = []string {}

    return &(force.EntityData)
}

// InstallUpdate_Input_Replace
type InstallUpdate_Input_Replace struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (replace *InstallUpdate_Input_Replace) GetEntityData() *types.CommonEntityData {
    replace.EntityData.YFilter = replace.YFilter
    replace.EntityData.YangName = "replace"
    replace.EntityData.BundleName = "cisco_ios_xr"
    replace.EntityData.ParentYangName = "input"
    replace.EntityData.SegmentPath = "replace"
    replace.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replace.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replace.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replace.EntityData.Children = types.NewOrderedMap()
    replace.EntityData.Leafs = types.NewOrderedMap()
    replace.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", replace.Packagepath})
    replace.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replace.Packagename})

    replace.EntityData.YListKeys = []string {}

    return &(replace.EntityData)
}

// InstallUpdate_Input_ReplaceForce
type InstallUpdate_Input_ReplaceForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (replaceForce *InstallUpdate_Input_ReplaceForce) GetEntityData() *types.CommonEntityData {
    replaceForce.EntityData.YFilter = replaceForce.YFilter
    replaceForce.EntityData.YangName = "replace-force"
    replaceForce.EntityData.BundleName = "cisco_ios_xr"
    replaceForce.EntityData.ParentYangName = "input"
    replaceForce.EntityData.SegmentPath = "replace-force"
    replaceForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaceForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaceForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaceForce.EntityData.Children = types.NewOrderedMap()
    replaceForce.EntityData.Leafs = types.NewOrderedMap()
    replaceForce.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", replaceForce.Packagepath})
    replaceForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replaceForce.Packagename})

    replaceForce.EntityData.YListKeys = []string {}

    return &(replaceForce.EntityData)
}

// InstallUpdate_Input_ReplaceCommit
type InstallUpdate_Input_ReplaceCommit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (replaceCommit *InstallUpdate_Input_ReplaceCommit) GetEntityData() *types.CommonEntityData {
    replaceCommit.EntityData.YFilter = replaceCommit.YFilter
    replaceCommit.EntityData.YangName = "replace-commit"
    replaceCommit.EntityData.BundleName = "cisco_ios_xr"
    replaceCommit.EntityData.ParentYangName = "input"
    replaceCommit.EntityData.SegmentPath = "replace-commit"
    replaceCommit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaceCommit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaceCommit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaceCommit.EntityData.Children = types.NewOrderedMap()
    replaceCommit.EntityData.Leafs = types.NewOrderedMap()
    replaceCommit.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", replaceCommit.Packagepath})
    replaceCommit.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replaceCommit.Packagename})

    replaceCommit.EntityData.YListKeys = []string {}

    return &(replaceCommit.EntityData)
}

// InstallUpdate_Input_ReplaceCommitForce
type InstallUpdate_Input_ReplaceCommitForce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is string.
    Packagepath interface{}

    // The type is slice of string.
    Packagename []interface{}
}

func (replaceCommitForce *InstallUpdate_Input_ReplaceCommitForce) GetEntityData() *types.CommonEntityData {
    replaceCommitForce.EntityData.YFilter = replaceCommitForce.YFilter
    replaceCommitForce.EntityData.YangName = "replace-commit-force"
    replaceCommitForce.EntityData.BundleName = "cisco_ios_xr"
    replaceCommitForce.EntityData.ParentYangName = "input"
    replaceCommitForce.EntityData.SegmentPath = "replace-commit-force"
    replaceCommitForce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replaceCommitForce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replaceCommitForce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replaceCommitForce.EntityData.Children = types.NewOrderedMap()
    replaceCommitForce.EntityData.Leafs = types.NewOrderedMap()
    replaceCommitForce.EntityData.Leafs.Append("packagepath", types.YLeaf{"Packagepath", replaceCommitForce.Packagepath})
    replaceCommitForce.EntityData.Leafs.Append("packagename", types.YLeaf{"Packagename", replaceCommitForce.Packagename})

    replaceCommitForce.EntityData.YListKeys = []string {}

    return &(replaceCommitForce.EntityData)
}

