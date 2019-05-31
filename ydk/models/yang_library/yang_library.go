// This module contains monitoring information about the YANG
// modules and submodules that are used within a YANG-based
// server.
// Copyright (c) 2016 IETF Trust and the persons identified as
// authors of the code.  All rights reserved.
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Contains a server-specific identifier representing the current set of
    // modules and submodules.  The server MUST change the value of this leaf if
    // the information represented by the 'module' list instances has changed. The
    // type is string. This attribute is mandatory.
    ModuleSetId interface{}

    // Each entry represents one revision of one module currently supported by the
    // server. The type is slice of ModulesState_Module.
    Module []*ModulesState_Module
}

func (modulesState *ModulesState) GetEntityData() *types.CommonEntityData {
    modulesState.EntityData.YFilter = modulesState.YFilter
    modulesState.EntityData.YangName = "modules-state"
    modulesState.EntityData.BundleName = "ietf"
    modulesState.EntityData.ParentYangName = "ietf-yang-library"
    modulesState.EntityData.SegmentPath = "ietf-yang-library:modules-state"
    modulesState.EntityData.AbsolutePath = modulesState.EntityData.SegmentPath
    modulesState.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    modulesState.EntityData.NamespaceTable = ietf.GetNamespaces()
    modulesState.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    modulesState.EntityData.Children = types.NewOrderedMap()
    modulesState.EntityData.Children.Append("module", types.YChild{"Module", nil})
    for i := range modulesState.Module {
        modulesState.EntityData.Children.Append(types.GetSegmentPath(modulesState.Module[i]), types.YChild{"Module", modulesState.Module[i]})
    }
    modulesState.EntityData.Leafs = types.NewOrderedMap()
    modulesState.EntityData.Leafs.Append("module-set-id", types.YLeaf{"ModuleSetId", modulesState.ModuleSetId})

    modulesState.EntityData.YListKeys = []string {}

    return &(modulesState.EntityData)
}

// ModulesState_Module
// Each entry represents one revision of one module
// currently supported by the server.
type ModulesState_Module struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: b'\\d{4}-\\d{2}-\\d{2}', or string with length: 0..0.
    Revision interface{}

    // Contains a URL that represents the YANG schema resource for this module or
    // submodule. This leaf will only be present if there is a URL available for
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
    // same entry MAY appear within multiple 'module' entries. The deviation
    // module MUST be present in the 'module' list, with the same name and
    // revision values. The 'conformance-type' value will be 'implement' for the
    // deviation module. The type is slice of ModulesState_Module_Deviation.
    Deviation []*ModulesState_Module_Deviation

    // Each entry represents one submodule within the parent module. The type is
    // slice of ModulesState_Module_Submodule.
    Submodule []*ModulesState_Module_Submodule
}

func (module *ModulesState_Module) GetEntityData() *types.CommonEntityData {
    module.EntityData.YFilter = module.YFilter
    module.EntityData.YangName = "module"
    module.EntityData.BundleName = "ietf"
    module.EntityData.ParentYangName = "modules-state"
    module.EntityData.SegmentPath = "module" + types.AddKeyToken(module.Name, "name") + types.AddKeyToken(module.Revision, "revision")
    module.EntityData.AbsolutePath = "ietf-yang-library:modules-state/" + module.EntityData.SegmentPath
    module.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    module.EntityData.NamespaceTable = ietf.GetNamespaces()
    module.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    module.EntityData.Children = types.NewOrderedMap()
    module.EntityData.Children.Append("deviation", types.YChild{"Deviation", nil})
    for i := range module.Deviation {
        module.EntityData.Children.Append(types.GetSegmentPath(module.Deviation[i]), types.YChild{"Deviation", module.Deviation[i]})
    }
    module.EntityData.Children.Append("submodule", types.YChild{"Submodule", nil})
    for i := range module.Submodule {
        module.EntityData.Children.Append(types.GetSegmentPath(module.Submodule[i]), types.YChild{"Submodule", module.Submodule[i]})
    }
    module.EntityData.Leafs = types.NewOrderedMap()
    module.EntityData.Leafs.Append("name", types.YLeaf{"Name", module.Name})
    module.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", module.Revision})
    module.EntityData.Leafs.Append("schema", types.YLeaf{"Schema", module.Schema})
    module.EntityData.Leafs.Append("namespace", types.YLeaf{"Namespace", module.Namespace})
    module.EntityData.Leafs.Append("feature", types.YLeaf{"Feature", module.Feature})
    module.EntityData.Leafs.Append("conformance-type", types.YLeaf{"ConformanceType", module.ConformanceType})

    module.EntityData.YListKeys = []string {"Name", "Revision"}

    return &(module.EntityData)
}

// ModulesState_Module_Deviation
// List of YANG deviation module names and revisions
// used by this server to modify the conformance of
// the module associated with this entry.  Note that
// the same module can be used for deviations for
// multiple modules, so the same entry MAY appear
// within multiple 'module' entries.
// The deviation module MUST be present in the 'module'
// list, with the same name and revision values.
// The 'conformance-type' value will be 'implement' for
// the deviation module.
type ModulesState_Module_Deviation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: b'\\d{4}-\\d{2}-\\d{2}', or string with length: 0..0.
    Revision interface{}
}

func (deviation *ModulesState_Module_Deviation) GetEntityData() *types.CommonEntityData {
    deviation.EntityData.YFilter = deviation.YFilter
    deviation.EntityData.YangName = "deviation"
    deviation.EntityData.BundleName = "ietf"
    deviation.EntityData.ParentYangName = "module"
    deviation.EntityData.SegmentPath = "deviation" + types.AddKeyToken(deviation.Name, "name") + types.AddKeyToken(deviation.Revision, "revision")
    deviation.EntityData.AbsolutePath = "ietf-yang-library:modules-state/module/" + deviation.EntityData.SegmentPath
    deviation.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    deviation.EntityData.NamespaceTable = ietf.GetNamespaces()
    deviation.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    deviation.EntityData.Children = types.NewOrderedMap()
    deviation.EntityData.Leafs = types.NewOrderedMap()
    deviation.EntityData.Leafs.Append("name", types.YLeaf{"Name", deviation.Name})
    deviation.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", deviation.Revision})

    deviation.EntityData.YListKeys = []string {"Name", "Revision"}

    return &(deviation.EntityData)
}

// ModulesState_Module_Submodule
// Each entry represents one submodule within the
// parent module.
type ModulesState_Module_Submodule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The YANG module or submodule name. The type is
    // string with length: 1..18446744073709551615.
    Name interface{}

    // This attribute is a key. The YANG module or submodule revision date. A
    // zero-length string is used if no revision statement is present in the YANG
    // module or submodule. The type is one of the following types: string with
    // pattern: b'\\d{4}-\\d{2}-\\d{2}', or string with length: 0..0.
    Revision interface{}

    // Contains a URL that represents the YANG schema resource for this module or
    // submodule. This leaf will only be present if there is a URL available for
    // retrieval of the schema for this entry. The type is string.
    Schema interface{}
}

func (submodule *ModulesState_Module_Submodule) GetEntityData() *types.CommonEntityData {
    submodule.EntityData.YFilter = submodule.YFilter
    submodule.EntityData.YangName = "submodule"
    submodule.EntityData.BundleName = "ietf"
    submodule.EntityData.ParentYangName = "module"
    submodule.EntityData.SegmentPath = "submodule" + types.AddKeyToken(submodule.Name, "name") + types.AddKeyToken(submodule.Revision, "revision")
    submodule.EntityData.AbsolutePath = "ietf-yang-library:modules-state/module/" + submodule.EntityData.SegmentPath
    submodule.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    submodule.EntityData.NamespaceTable = ietf.GetNamespaces()
    submodule.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    submodule.EntityData.Children = types.NewOrderedMap()
    submodule.EntityData.Leafs = types.NewOrderedMap()
    submodule.EntityData.Leafs.Append("name", types.YLeaf{"Name", submodule.Name})
    submodule.EntityData.Leafs.Append("revision", types.YLeaf{"Revision", submodule.Revision})
    submodule.EntityData.Leafs.Append("schema", types.YLeaf{"Schema", submodule.Schema})

    submodule.EntityData.YListKeys = []string {"Name", "Revision"}

    return &(submodule.EntityData)
}

// ModulesState_Module_ConformanceType represents for the YANG module identified by this entry.
type ModulesState_Module_ConformanceType string

const (
    // Indicates that the server implements one or more
    // protocol-accessible objects defined in the YANG module
    // identified in this entry.  This includes deviation
    // statements defined in the module.
    // For YANG version 1.1 modules, there is at most one
    // module entry with conformance type 'implement' for a
    // particular module name, since YANG 1.1 requires that,
    // at most, one revision of a module is implemented.
    // For YANG version 1 modules, there SHOULD NOT be more
    // than one module entry for a particular module name.
    ModulesState_Module_ConformanceType_implement ModulesState_Module_ConformanceType = "implement"

    // Indicates that the server imports reusable definitions
    // from the specified revision of the module but does
    // not implement any protocol-accessible objects from
    // this revision.
    // Multiple module entries for the same module name MAY
    // exist.  This can occur if multiple modules import the
    // same module but specify different revision dates in
    // the import statements.
    ModulesState_Module_ConformanceType_import_ ModulesState_Module_ConformanceType = "import"
)

