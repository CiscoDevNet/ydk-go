// This module contains monitoring information about the YANG
// modules and submodules that are used within a YANG-based
// server.
// 
// Copyright (c) 2016 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 7895; see
// the RFC itself for full legal notices.
package yang_library

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package yang_library"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-yang-library modules-state}", reflect.TypeOf(ModulesState{}))
    ydk.RegisterEntity("ietf-yang-library:modules-state", reflect.TypeOf(ModulesState{}))
}

// ModulesState
// Contains YANG module monitoring information.
type ModulesState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Contains a server-specific identifier representing the current set of
    // modules and submodules.  The server MUST change the value of this leaf if
    // the information represented by the 'module' list instances has changed. The
    // type is string. This attribute is mandatory.
    ModuleSetId interface{}

    // Each entry represents one revision of one module currently supported by the
    // server. The type is slice of ModulesState_Module.
    Module []ModulesState_Module
}

func (modulesState *ModulesState) GetFilter() yfilter.YFilter { return modulesState.YFilter }

func (modulesState *ModulesState) SetFilter(yf yfilter.YFilter) { modulesState.YFilter = yf }

func (modulesState *ModulesState) GetGoName(yname string) string {
    if yname == "module-set-id" { return "ModuleSetId" }
    if yname == "module" { return "Module" }
    return ""
}

func (modulesState *ModulesState) GetSegmentPath() string {
    return "ietf-yang-library:modules-state"
}

func (modulesState *ModulesState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "module" {
        for _, c := range modulesState.Module {
            if modulesState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ModulesState_Module{}
        modulesState.Module = append(modulesState.Module, child)
        return &modulesState.Module[len(modulesState.Module)-1]
    }
    return nil
}

func (modulesState *ModulesState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range modulesState.Module {
        children[modulesState.Module[i].GetSegmentPath()] = &modulesState.Module[i]
    }
    return children
}

func (modulesState *ModulesState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["module-set-id"] = modulesState.ModuleSetId
    return leafs
}

func (modulesState *ModulesState) GetBundleName() string { return "ietf" }

func (modulesState *ModulesState) GetYangName() string { return "modules-state" }

func (modulesState *ModulesState) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (modulesState *ModulesState) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (modulesState *ModulesState) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (modulesState *ModulesState) SetParent(parent types.Entity) { modulesState.parent = parent }

func (modulesState *ModulesState) GetParent() types.Entity { return modulesState.parent }

func (modulesState *ModulesState) GetParentYangName() string { return "ietf-yang-library" }

// ModulesState_Module
// Each entry represents one revision of one module
// currently supported by the server.
type ModulesState_Module struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: \d{4}-\d{2}-\d{2}, or string with length: 0.
    Revision interface{}

    // Contains a URL that represents the YANG schema resource for this module or
    // submodule.  This leaf will only be present if there is a URL available for
    // retrieval of the schema for this entry. The type is string.
    Schema interface{}

    // The XML namespace identifier for this module. The type is string. This
    // attribute is mandatory.
    Namespace interface{}

    // List of YANG feature names from this module that are supported by the
    // server, regardless of whether they are defined in the module or any
    // included submodule. The type is slice of string with length:
    // 1..18446744073709551615.
    Feature []interface{}

    // Indicates the type of conformance the server is claiming for the YANG
    // module identified by this entry. The type is ConformanceType. This
    // attribute is mandatory.
    ConformanceType interface{}

    // List of YANG deviation module names and revisions used by this server to
    // modify the conformance of the module associated with this entry.  Note that
    // the same module can be used for deviations for multiple modules, so the
    // same entry MAY appear within multiple 'module' entries.  The deviation
    // module MUST be present in the 'module' list, with the same name and
    // revision values. The 'conformance-type' value will be 'implement' for the
    // deviation module. The type is slice of ModulesState_Module_Deviation.
    Deviation []ModulesState_Module_Deviation

    // Each entry represents one submodule within the parent module. The type is
    // slice of ModulesState_Module_Submodule.
    Submodule []ModulesState_Module_Submodule
}

func (module *ModulesState_Module) GetFilter() yfilter.YFilter { return module.YFilter }

func (module *ModulesState_Module) SetFilter(yf yfilter.YFilter) { module.YFilter = yf }

func (module *ModulesState_Module) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revision" { return "Revision" }
    if yname == "schema" { return "Schema" }
    if yname == "namespace" { return "Namespace" }
    if yname == "feature" { return "Feature" }
    if yname == "conformance-type" { return "ConformanceType" }
    if yname == "deviation" { return "Deviation" }
    if yname == "submodule" { return "Submodule" }
    return ""
}

func (module *ModulesState_Module) GetSegmentPath() string {
    return "module" + "[name='" + fmt.Sprintf("%v", module.Name) + "']" + "[revision='" + fmt.Sprintf("%v", module.Revision) + "']"
}

func (module *ModulesState_Module) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "deviation" {
        for _, c := range module.Deviation {
            if module.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ModulesState_Module_Deviation{}
        module.Deviation = append(module.Deviation, child)
        return &module.Deviation[len(module.Deviation)-1]
    }
    if childYangName == "submodule" {
        for _, c := range module.Submodule {
            if module.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := ModulesState_Module_Submodule{}
        module.Submodule = append(module.Submodule, child)
        return &module.Submodule[len(module.Submodule)-1]
    }
    return nil
}

func (module *ModulesState_Module) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range module.Deviation {
        children[module.Deviation[i].GetSegmentPath()] = &module.Deviation[i]
    }
    for i := range module.Submodule {
        children[module.Submodule[i].GetSegmentPath()] = &module.Submodule[i]
    }
    return children
}

func (module *ModulesState_Module) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = module.Name
    leafs["revision"] = module.Revision
    leafs["schema"] = module.Schema
    leafs["namespace"] = module.Namespace
    leafs["feature"] = module.Feature
    leafs["conformance-type"] = module.ConformanceType
    return leafs
}

func (module *ModulesState_Module) GetBundleName() string { return "ietf" }

func (module *ModulesState_Module) GetYangName() string { return "module" }

func (module *ModulesState_Module) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (module *ModulesState_Module) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (module *ModulesState_Module) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (module *ModulesState_Module) SetParent(parent types.Entity) { module.parent = parent }

func (module *ModulesState_Module) GetParent() types.Entity { return module.parent }

func (module *ModulesState_Module) GetParentYangName() string { return "modules-state" }

// ModulesState_Module_Deviation
// List of YANG deviation module names and revisions
// used by this server to modify the conformance of
// the module associated with this entry.  Note that
// the same module can be used for deviations for
// multiple modules, so the same entry MAY appear
// within multiple 'module' entries.
// 
// The deviation module MUST be present in the 'module'
// list, with the same name and revision values.
// The 'conformance-type' value will be 'implement' for
// the deviation module.
type ModulesState_Module_Deviation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: \d{4}-\d{2}-\d{2}, or string with length: 0.
    Revision interface{}
}

func (deviation *ModulesState_Module_Deviation) GetFilter() yfilter.YFilter { return deviation.YFilter }

func (deviation *ModulesState_Module_Deviation) SetFilter(yf yfilter.YFilter) { deviation.YFilter = yf }

func (deviation *ModulesState_Module_Deviation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revision" { return "Revision" }
    return ""
}

func (deviation *ModulesState_Module_Deviation) GetSegmentPath() string {
    return "deviation" + "[name='" + fmt.Sprintf("%v", deviation.Name) + "']" + "[revision='" + fmt.Sprintf("%v", deviation.Revision) + "']"
}

func (deviation *ModulesState_Module_Deviation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (deviation *ModulesState_Module_Deviation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (deviation *ModulesState_Module_Deviation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = deviation.Name
    leafs["revision"] = deviation.Revision
    return leafs
}

func (deviation *ModulesState_Module_Deviation) GetBundleName() string { return "ietf" }

func (deviation *ModulesState_Module_Deviation) GetYangName() string { return "deviation" }

func (deviation *ModulesState_Module_Deviation) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (deviation *ModulesState_Module_Deviation) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (deviation *ModulesState_Module_Deviation) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (deviation *ModulesState_Module_Deviation) SetParent(parent types.Entity) { deviation.parent = parent }

func (deviation *ModulesState_Module_Deviation) GetParent() types.Entity { return deviation.parent }

func (deviation *ModulesState_Module_Deviation) GetParentYangName() string { return "module" }

// ModulesState_Module_Submodule
// Each entry represents one submodule within the
// parent module.
type ModulesState_Module_Submodule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: \d{4}-\d{2}-\d{2}, or string with length: 0.
    Revision interface{}

    // Contains a URL that represents the YANG schema resource for this module or
    // submodule.  This leaf will only be present if there is a URL available for
    // retrieval of the schema for this entry. The type is string.
    Schema interface{}
}

func (submodule *ModulesState_Module_Submodule) GetFilter() yfilter.YFilter { return submodule.YFilter }

func (submodule *ModulesState_Module_Submodule) SetFilter(yf yfilter.YFilter) { submodule.YFilter = yf }

func (submodule *ModulesState_Module_Submodule) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "revision" { return "Revision" }
    if yname == "schema" { return "Schema" }
    return ""
}

func (submodule *ModulesState_Module_Submodule) GetSegmentPath() string {
    return "submodule" + "[name='" + fmt.Sprintf("%v", submodule.Name) + "']" + "[revision='" + fmt.Sprintf("%v", submodule.Revision) + "']"
}

func (submodule *ModulesState_Module_Submodule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (submodule *ModulesState_Module_Submodule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (submodule *ModulesState_Module_Submodule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = submodule.Name
    leafs["revision"] = submodule.Revision
    leafs["schema"] = submodule.Schema
    return leafs
}

func (submodule *ModulesState_Module_Submodule) GetBundleName() string { return "ietf" }

func (submodule *ModulesState_Module_Submodule) GetYangName() string { return "submodule" }

func (submodule *ModulesState_Module_Submodule) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (submodule *ModulesState_Module_Submodule) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (submodule *ModulesState_Module_Submodule) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (submodule *ModulesState_Module_Submodule) SetParent(parent types.Entity) { submodule.parent = parent }

func (submodule *ModulesState_Module_Submodule) GetParent() types.Entity { return submodule.parent }

func (submodule *ModulesState_Module_Submodule) GetParentYangName() string { return "module" }

// ModulesState_Module_ConformanceType represents for the YANG module identified by this entry.
type ModulesState_Module_ConformanceType string

const (
    // Indicates that the server implements one or more
    // protocol-accessible objects defined in the YANG module
    // identified in this entry.  This includes deviation
    // statements defined in the module.
    // 
    // For YANG version 1.1 modules, there is at most one
    // module entry with conformance type 'implement' for a
    // particular module name, since YANG 1.1 requires that,
    // at most, one revision of a module is implemented.
    // 
    // For YANG version 1 modules, there SHOULD NOT be more
    // than one module entry for a particular module name.
    ModulesState_Module_ConformanceType_implement ModulesState_Module_ConformanceType = "implement"

    // Indicates that the server imports reusable definitions
    // from the specified revision of the module but does
    // not implement any protocol-accessible objects from
    // this revision.
    // 
    // Multiple module entries for the same module name MAY
    // exist.  This can occur if multiple modules import the
    // same module but specify different revision dates in
    // the import statements.
    ModulesState_Module_ConformanceType_import_ ModulesState_Module_ConformanceType = "import"
)

