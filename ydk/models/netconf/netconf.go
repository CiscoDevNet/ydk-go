// NETCONF Protocol Data Types and Protocol Operations.
// 
// Copyright (c) 2011 IETF Trust and the persons identified as
// the document authors.  All rights reserved.
// 
// Redistribution and use in source and binary forms, with or
// without modification, is permitted pursuant to, and subject
// to the license terms contained in, the Simplified BSD License
// set forth in Section 4.c of the IETF Trust's Legal Provisions
// Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC 6241; see
// the RFC itself for full legal notices.
package netconf

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package netconf"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 get-config}", reflect.TypeOf(GetConfig{}))
    ydk.RegisterEntity("ietf-netconf:get-config", reflect.TypeOf(GetConfig{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 edit-config}", reflect.TypeOf(EditConfig{}))
    ydk.RegisterEntity("ietf-netconf:edit-config", reflect.TypeOf(EditConfig{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 copy-config}", reflect.TypeOf(CopyConfig{}))
    ydk.RegisterEntity("ietf-netconf:copy-config", reflect.TypeOf(CopyConfig{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 delete-config}", reflect.TypeOf(DeleteConfig{}))
    ydk.RegisterEntity("ietf-netconf:delete-config", reflect.TypeOf(DeleteConfig{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 lock}", reflect.TypeOf(Lock{}))
    ydk.RegisterEntity("ietf-netconf:lock", reflect.TypeOf(Lock{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 unlock}", reflect.TypeOf(Unlock{}))
    ydk.RegisterEntity("ietf-netconf:unlock", reflect.TypeOf(Unlock{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 get}", reflect.TypeOf(Get{}))
    ydk.RegisterEntity("ietf-netconf:get", reflect.TypeOf(Get{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 close-session}", reflect.TypeOf(CloseSession{}))
    ydk.RegisterEntity("ietf-netconf:close-session", reflect.TypeOf(CloseSession{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 kill-session}", reflect.TypeOf(KillSession{}))
    ydk.RegisterEntity("ietf-netconf:kill-session", reflect.TypeOf(KillSession{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 commit}", reflect.TypeOf(Commit{}))
    ydk.RegisterEntity("ietf-netconf:commit", reflect.TypeOf(Commit{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 discard-changes}", reflect.TypeOf(DiscardChanges{}))
    ydk.RegisterEntity("ietf-netconf:discard-changes", reflect.TypeOf(DiscardChanges{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 cancel-commit}", reflect.TypeOf(CancelCommit{}))
    ydk.RegisterEntity("ietf-netconf:cancel-commit", reflect.TypeOf(CancelCommit{}))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:netconf:base:1.0 validate}", reflect.TypeOf(Validate{}))
    ydk.RegisterEntity("ietf-netconf:validate", reflect.TypeOf(Validate{}))
}

// ErrorTagType represents NETCONF Error Tag
type ErrorTagType string

const (
    // The request requires a resource that
    // already is in use.
    ErrorTagType_in_use ErrorTagType = "in-use"

    // The request specifies an unacceptable value for one
    // or more parameters.
    ErrorTagType_invalid_value ErrorTagType = "invalid-value"

    // The request or response (that would be generated) is
    // too large for the implementation to handle.
    ErrorTagType_too_big ErrorTagType = "too-big"

    // An expected attribute is missing.
    ErrorTagType_missing_attribute ErrorTagType = "missing-attribute"

    // An attribute value is not correct; e.g., wrong type,
    // out of range, pattern mismatch.
    ErrorTagType_bad_attribute ErrorTagType = "bad-attribute"

    // An unexpected attribute is present.
    ErrorTagType_unknown_attribute ErrorTagType = "unknown-attribute"

    // An expected element is missing.
    ErrorTagType_missing_element ErrorTagType = "missing-element"

    // An element value is not correct; e.g., wrong type,
    // out of range, pattern mismatch.
    ErrorTagType_bad_element ErrorTagType = "bad-element"

    // An unexpected element is present.
    ErrorTagType_unknown_element ErrorTagType = "unknown-element"

    // An unexpected namespace is present.
    ErrorTagType_unknown_namespace ErrorTagType = "unknown-namespace"

    // Access to the requested protocol operation or
    // data model is denied because authorization failed.
    ErrorTagType_access_denied ErrorTagType = "access-denied"

    // Access to the requested lock is denied because the
    // lock is currently held by another entity.
    ErrorTagType_lock_denied ErrorTagType = "lock-denied"

    // Request could not be completed because of
    // insufficient resources.
    ErrorTagType_resource_denied ErrorTagType = "resource-denied"

    // Request to roll back some configuration change (via
    // rollback-on-error or <discard-changes> operations)
    // was not completed for some reason.
    ErrorTagType_rollback_failed ErrorTagType = "rollback-failed"

    // Request could not be completed because the relevant
    // data model content already exists.  For example,
    // a 'create' operation was attempted on data that
    // already exists.
    ErrorTagType_data_exists ErrorTagType = "data-exists"

    // Request could not be completed because the relevant
    // data model content does not exist.  For example,
    // a 'delete' operation was attempted on
    // data that does not exist.
    ErrorTagType_data_missing ErrorTagType = "data-missing"

    // Request could not be completed because the requested
    // operation is not supported by this implementation.
    ErrorTagType_operation_not_supported ErrorTagType = "operation-not-supported"

    // Request could not be completed because the requested
    // operation failed for some reason not covered by
    // any other error condition.
    ErrorTagType_operation_failed ErrorTagType = "operation-failed"

    // This error-tag is obsolete, and SHOULD NOT be sent
    // by servers conforming to this document.
    ErrorTagType_partial_operation ErrorTagType = "partial-operation"

    // A message could not be handled because it failed to
    // be parsed correctly.  For example, the message is not
    // well-formed XML or it uses an invalid character set.
    ErrorTagType_malformed_message ErrorTagType = "malformed-message"
)

// ErrorSeverityType represents NETCONF Error Severity
type ErrorSeverityType string

const (
    // Error severity
    ErrorSeverityType_error_ ErrorSeverityType = "error"

    // Warning severity
    ErrorSeverityType_warning ErrorSeverityType = "warning"
)

// EditOperationType represents NETCONF 'operation' attribute values
type EditOperationType string

const (
    // The configuration data identified by the
    // element containing this attribute is merged
    // with the configuration at the corresponding
    // level in the configuration datastore identified
    // by the target parameter.
    EditOperationType_merge EditOperationType = "merge"

    // The configuration data identified by the element
    // containing this attribute replaces any related
    // configuration in the configuration datastore
    // identified by the target parameter.  If no such
    // configuration data exists in the configuration
    // datastore, it is created.  Unlike a
    // <copy-config> operation, which replaces the
    // entire target configuration, only the configuration
    // actually present in the config parameter is affected.
    EditOperationType_replace EditOperationType = "replace"

    // The configuration data identified by the element
    // containing this attribute is added to the
    // configuration if and only if the configuration
    // data does not already exist in the configuration
    // datastore.  If the configuration data exists, an
    // <rpc-error> element is returned with an
    // <error-tag> value of 'data-exists'.
    EditOperationType_create EditOperationType = "create"

    // The configuration data identified by the element
    // containing this attribute is deleted from the
    // configuration if and only if the configuration
    // data currently exists in the configuration
    // datastore.  If the configuration data does not
    // exist, an <rpc-error> element is returned with
    // an <error-tag> value of 'data-missing'.
    EditOperationType_delete_ EditOperationType = "delete"

    // The configuration data identified by the element
    // containing this attribute is deleted from the
    // configuration if the configuration
    // data currently exists in the configuration
    // datastore.  If the configuration data does not
    // exist, the 'remove' operation is silently ignored
    // by the server.
    EditOperationType_remove EditOperationType = "remove"
)

// GetConfig
// Retrieve all or part of a specified configuration.
type GetConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetConfig_Input

    
    Output GetConfig_Output
}

func (getConfig *GetConfig) GetEntityData() *types.CommonEntityData {
    getConfig.EntityData.YFilter = getConfig.YFilter
    getConfig.EntityData.YangName = "get-config"
    getConfig.EntityData.BundleName = "ietf"
    getConfig.EntityData.ParentYangName = "ietf-netconf"
    getConfig.EntityData.SegmentPath = "ietf-netconf:get-config"
    getConfig.EntityData.AbsolutePath = getConfig.EntityData.SegmentPath
    getConfig.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    getConfig.EntityData.NamespaceTable = ietf.GetNamespaces()
    getConfig.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    getConfig.EntityData.Children = types.NewOrderedMap()
    getConfig.EntityData.Children.Append("input", types.YChild{"Input", &getConfig.Input})
    getConfig.EntityData.Children.Append("output", types.YChild{"Output", &getConfig.Output})
    getConfig.EntityData.Leafs = types.NewOrderedMap()

    getConfig.EntityData.YListKeys = []string {}

    return &(getConfig.EntityData)
}

// GetConfig_Input
type GetConfig_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subtree or XPath filter to use. The type is string.
    Filter interface{}

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}

    // Particular configuration to retrieve.
    Source GetConfig_Input_Source
}

func (input *GetConfig_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "get-config"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:get-config/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("source", types.YChild{"Source", &input.Source})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("with-defaults", types.YLeaf{"WithDefaults", input.WithDefaults})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetConfig_Input_Source
// Particular configuration to retrieve.
type GetConfig_Input_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config source. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config source. The type is interface{}.
    Running interface{}

    // The startup configuration is the config source. This is
    // optional-to-implement on the server because not all servers will support
    // filtering for this datastore. The type is interface{}.
    Startup interface{}
}

func (source *GetConfig_Input_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "ietf"
    source.EntityData.ParentYangName = "input"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "ietf-netconf:get-config/input/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    source.EntityData.NamespaceTable = ietf.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", source.Candidate})
    source.EntityData.Leafs.Append("running", types.YLeaf{"Running", source.Running})
    source.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", source.Startup})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// GetConfig_Output
type GetConfig_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Copy of the source datastore subset that matched the filter criteria (if
    // any).  An empty data container indicates that the request did not produce
    // any results. The type is string.
    Data interface{}
}

func (output *GetConfig_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "get-config"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "ietf-netconf:get-config/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("data", types.YLeaf{"Data", output.Data})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// EditConfig
// The <edit-config> operation loads all or part of a specified
// configuration to the specified target configuration.
type EditConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input EditConfig_Input
}

func (editConfig *EditConfig) GetEntityData() *types.CommonEntityData {
    editConfig.EntityData.YFilter = editConfig.YFilter
    editConfig.EntityData.YangName = "edit-config"
    editConfig.EntityData.BundleName = "ietf"
    editConfig.EntityData.ParentYangName = "ietf-netconf"
    editConfig.EntityData.SegmentPath = "ietf-netconf:edit-config"
    editConfig.EntityData.AbsolutePath = editConfig.EntityData.SegmentPath
    editConfig.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    editConfig.EntityData.NamespaceTable = ietf.GetNamespaces()
    editConfig.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    editConfig.EntityData.Children = types.NewOrderedMap()
    editConfig.EntityData.Children.Append("input", types.YChild{"Input", &editConfig.Input})
    editConfig.EntityData.Leafs = types.NewOrderedMap()

    editConfig.EntityData.YListKeys = []string {}

    return &(editConfig.EntityData)
}

// EditConfig_Input
type EditConfig_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The default operation to use. The type is DefaultOperation. The default
    // value is merge.
    DefaultOperation interface{}

    // The test option to use. The type is TestOption. The default value is
    // test-then-set.
    TestOption interface{}

    // The error option to use. The type is ErrorOption. The default value is
    // stop-on-error.
    ErrorOption interface{}

    // Inline Config content. The type is string.
    Config interface{}

    // URL-based config content. The type is string.
    Url interface{}

    // Particular configuration to edit.
    Target EditConfig_Input_Target
}

func (input *EditConfig_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "edit-config"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:edit-config/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("default-operation", types.YLeaf{"DefaultOperation", input.DefaultOperation})
    input.EntityData.Leafs.Append("test-option", types.YLeaf{"TestOption", input.TestOption})
    input.EntityData.Leafs.Append("error-option", types.YLeaf{"ErrorOption", input.ErrorOption})
    input.EntityData.Leafs.Append("config", types.YLeaf{"Config", input.Config})
    input.EntityData.Leafs.Append("url", types.YLeaf{"Url", input.Url})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// EditConfig_Input_Target
// Particular configuration to edit.
type EditConfig_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config source. The type is interface{}.
    Running interface{}
}

func (target *EditConfig_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "ietf"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "ietf-netconf:edit-config/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    target.EntityData.NamespaceTable = ietf.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", target.Candidate})
    target.EntityData.Leafs.Append("running", types.YLeaf{"Running", target.Running})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// EditConfig_Input_DefaultOperation represents The default operation to use.
type EditConfig_Input_DefaultOperation string

const (
    // The default operation is merge.
    EditConfig_Input_DefaultOperation_merge EditConfig_Input_DefaultOperation = "merge"

    // The default operation is replace.
    EditConfig_Input_DefaultOperation_replace EditConfig_Input_DefaultOperation = "replace"

    // There is no default operation.
    EditConfig_Input_DefaultOperation_none EditConfig_Input_DefaultOperation = "none"
)

// EditConfig_Input_ErrorOption represents The error option to use.
type EditConfig_Input_ErrorOption string

const (
    // The server will stop on errors.
    EditConfig_Input_ErrorOption_stop_on_error EditConfig_Input_ErrorOption = "stop-on-error"

    // The server may continue on errors.
    EditConfig_Input_ErrorOption_continue_on_error EditConfig_Input_ErrorOption = "continue-on-error"

    // The server will roll back on errors.
    // This value can only be used if the 'rollback-on-error'
    // feature is supported.
    EditConfig_Input_ErrorOption_rollback_on_error EditConfig_Input_ErrorOption = "rollback-on-error"
)

// EditConfig_Input_TestOption represents The test option to use.
type EditConfig_Input_TestOption string

const (
    // The server will test and then set if no errors.
    EditConfig_Input_TestOption_test_then_set EditConfig_Input_TestOption = "test-then-set"

    // The server will set without a test first.
    EditConfig_Input_TestOption_set EditConfig_Input_TestOption = "set"

    // The server will only test and not set, even
    // if there are no errors.
    EditConfig_Input_TestOption_test_only EditConfig_Input_TestOption = "test-only"
)

// CopyConfig
// Create or replace an entire configuration datastore with the
// contents of another complete configuration datastore.
type CopyConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CopyConfig_Input
}

func (copyConfig *CopyConfig) GetEntityData() *types.CommonEntityData {
    copyConfig.EntityData.YFilter = copyConfig.YFilter
    copyConfig.EntityData.YangName = "copy-config"
    copyConfig.EntityData.BundleName = "ietf"
    copyConfig.EntityData.ParentYangName = "ietf-netconf"
    copyConfig.EntityData.SegmentPath = "ietf-netconf:copy-config"
    copyConfig.EntityData.AbsolutePath = copyConfig.EntityData.SegmentPath
    copyConfig.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    copyConfig.EntityData.NamespaceTable = ietf.GetNamespaces()
    copyConfig.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    copyConfig.EntityData.Children = types.NewOrderedMap()
    copyConfig.EntityData.Children.Append("input", types.YChild{"Input", &copyConfig.Input})
    copyConfig.EntityData.Leafs = types.NewOrderedMap()

    copyConfig.EntityData.YListKeys = []string {}

    return &(copyConfig.EntityData)
}

// CopyConfig_Input
type CopyConfig_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}

    // Particular configuration to copy to.
    Target CopyConfig_Input_Target

    // Particular configuration to copy from.
    Source CopyConfig_Input_Source
}

func (input *CopyConfig_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "copy-config"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:copy-config/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Children.Append("source", types.YChild{"Source", &input.Source})
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("with-defaults", types.YLeaf{"WithDefaults", input.WithDefaults})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// CopyConfig_Input_Target
// Particular configuration to copy to.
type CopyConfig_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config target. This is
    // optional-to-implement on the server. The type is interface{}.
    Running interface{}

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}

    // The URL-based configuration is the config target. The type is string.
    Url interface{}
}

func (target *CopyConfig_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "ietf"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "ietf-netconf:copy-config/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    target.EntityData.NamespaceTable = ietf.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", target.Candidate})
    target.EntityData.Leafs.Append("running", types.YLeaf{"Running", target.Running})
    target.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", target.Startup})
    target.EntityData.Leafs.Append("url", types.YLeaf{"Url", target.Url})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// CopyConfig_Input_Source
// Particular configuration to copy from.
type CopyConfig_Input_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config source. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config source. The type is interface{}.
    Running interface{}

    // The startup configuration is the config source. The type is interface{}.
    Startup interface{}

    // The URL-based configuration is the config source. The type is string.
    Url interface{}

    // Inline Config content: <config> element.  Represents an entire
    // configuration datastore, not a subset of the running datastore. The type is
    // string.
    Config interface{}
}

func (source *CopyConfig_Input_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "ietf"
    source.EntityData.ParentYangName = "input"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "ietf-netconf:copy-config/input/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    source.EntityData.NamespaceTable = ietf.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", source.Candidate})
    source.EntityData.Leafs.Append("running", types.YLeaf{"Running", source.Running})
    source.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", source.Startup})
    source.EntityData.Leafs.Append("url", types.YLeaf{"Url", source.Url})
    source.EntityData.Leafs.Append("config", types.YLeaf{"Config", source.Config})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// DeleteConfig
// Delete a configuration datastore.
type DeleteConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input DeleteConfig_Input
}

func (deleteConfig *DeleteConfig) GetEntityData() *types.CommonEntityData {
    deleteConfig.EntityData.YFilter = deleteConfig.YFilter
    deleteConfig.EntityData.YangName = "delete-config"
    deleteConfig.EntityData.BundleName = "ietf"
    deleteConfig.EntityData.ParentYangName = "ietf-netconf"
    deleteConfig.EntityData.SegmentPath = "ietf-netconf:delete-config"
    deleteConfig.EntityData.AbsolutePath = deleteConfig.EntityData.SegmentPath
    deleteConfig.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    deleteConfig.EntityData.NamespaceTable = ietf.GetNamespaces()
    deleteConfig.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    deleteConfig.EntityData.Children = types.NewOrderedMap()
    deleteConfig.EntityData.Children.Append("input", types.YChild{"Input", &deleteConfig.Input})
    deleteConfig.EntityData.Leafs = types.NewOrderedMap()

    deleteConfig.EntityData.YListKeys = []string {}

    return &(deleteConfig.EntityData)
}

// DeleteConfig_Input
type DeleteConfig_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular configuration to delete.
    Target DeleteConfig_Input_Target
}

func (input *DeleteConfig_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "delete-config"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:delete-config/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DeleteConfig_Input_Target
// Particular configuration to delete.
type DeleteConfig_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}

    // The URL-based configuration is the config target. The type is string.
    Url interface{}
}

func (target *DeleteConfig_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "ietf"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "ietf-netconf:delete-config/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    target.EntityData.NamespaceTable = ietf.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", target.Startup})
    target.EntityData.Leafs.Append("url", types.YLeaf{"Url", target.Url})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// Lock
// The lock operation allows the client to lock the configuration
// system of a device.
type Lock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Lock_Input
}

func (lock *Lock) GetEntityData() *types.CommonEntityData {
    lock.EntityData.YFilter = lock.YFilter
    lock.EntityData.YangName = "lock"
    lock.EntityData.BundleName = "ietf"
    lock.EntityData.ParentYangName = "ietf-netconf"
    lock.EntityData.SegmentPath = "ietf-netconf:lock"
    lock.EntityData.AbsolutePath = lock.EntityData.SegmentPath
    lock.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    lock.EntityData.NamespaceTable = ietf.GetNamespaces()
    lock.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    lock.EntityData.Children = types.NewOrderedMap()
    lock.EntityData.Children.Append("input", types.YChild{"Input", &lock.Input})
    lock.EntityData.Leafs = types.NewOrderedMap()

    lock.EntityData.YListKeys = []string {}

    return &(lock.EntityData)
}

// Lock_Input
type Lock_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular configuration to lock.
    Target Lock_Input_Target
}

func (input *Lock_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "lock"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:lock/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Lock_Input_Target
// Particular configuration to lock.
type Lock_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config target. The type is interface{}.
    Running interface{}

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}
}

func (target *Lock_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "ietf"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "ietf-netconf:lock/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    target.EntityData.NamespaceTable = ietf.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", target.Candidate})
    target.EntityData.Leafs.Append("running", types.YLeaf{"Running", target.Running})
    target.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", target.Startup})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// Unlock
// The unlock operation is used to release a configuration lock,
// previously obtained with the 'lock' operation.
type Unlock struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Unlock_Input
}

func (unlock *Unlock) GetEntityData() *types.CommonEntityData {
    unlock.EntityData.YFilter = unlock.YFilter
    unlock.EntityData.YangName = "unlock"
    unlock.EntityData.BundleName = "ietf"
    unlock.EntityData.ParentYangName = "ietf-netconf"
    unlock.EntityData.SegmentPath = "ietf-netconf:unlock"
    unlock.EntityData.AbsolutePath = unlock.EntityData.SegmentPath
    unlock.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    unlock.EntityData.NamespaceTable = ietf.GetNamespaces()
    unlock.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    unlock.EntityData.Children = types.NewOrderedMap()
    unlock.EntityData.Children.Append("input", types.YChild{"Input", &unlock.Input})
    unlock.EntityData.Leafs = types.NewOrderedMap()

    unlock.EntityData.YListKeys = []string {}

    return &(unlock.EntityData)
}

// Unlock_Input
type Unlock_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular configuration to unlock.
    Target Unlock_Input_Target
}

func (input *Unlock_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "unlock"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:unlock/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("target", types.YChild{"Target", &input.Target})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Unlock_Input_Target
// Particular configuration to unlock.
type Unlock_Input_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config target. The type is interface{}.
    Running interface{}

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}
}

func (target *Unlock_Input_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "ietf"
    target.EntityData.ParentYangName = "input"
    target.EntityData.SegmentPath = "target"
    target.EntityData.AbsolutePath = "ietf-netconf:unlock/input/" + target.EntityData.SegmentPath
    target.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    target.EntityData.NamespaceTable = ietf.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", target.Candidate})
    target.EntityData.Leafs.Append("running", types.YLeaf{"Running", target.Running})
    target.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", target.Startup})

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// Get
// Retrieve running configuration and device state information.
type Get struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Get_Input

    
    Output Get_Output
}

func (get *Get) GetEntityData() *types.CommonEntityData {
    get.EntityData.YFilter = get.YFilter
    get.EntityData.YangName = "get"
    get.EntityData.BundleName = "ietf"
    get.EntityData.ParentYangName = "ietf-netconf"
    get.EntityData.SegmentPath = "ietf-netconf:get"
    get.EntityData.AbsolutePath = get.EntityData.SegmentPath
    get.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    get.EntityData.NamespaceTable = ietf.GetNamespaces()
    get.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    get.EntityData.Children = types.NewOrderedMap()
    get.EntityData.Children.Append("input", types.YChild{"Input", &get.Input})
    get.EntityData.Children.Append("output", types.YChild{"Output", &get.Output})
    get.EntityData.Leafs = types.NewOrderedMap()

    get.EntityData.YListKeys = []string {}

    return &(get.EntityData)
}

// Get_Input
type Get_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This parameter specifies the portion of the system configuration and state
    // data to retrieve. The type is string.
    Filter interface{}

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}
}

func (input *Get_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "get"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:get/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", input.Filter})
    input.EntityData.Leafs.Append("with-defaults", types.YLeaf{"WithDefaults", input.WithDefaults})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Get_Output
type Get_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Copy of the running datastore subset and/or state data that matched the
    // filter criteria (if any). An empty data container indicates that the
    // request did not produce any results. The type is string.
    Data interface{}
}

func (output *Get_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "ietf"
    output.EntityData.ParentYangName = "get"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "ietf-netconf:get/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    output.EntityData.NamespaceTable = ietf.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("data", types.YLeaf{"Data", output.Data})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// CloseSession
// Request graceful termination of a NETCONF session.
type CloseSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (closeSession *CloseSession) GetEntityData() *types.CommonEntityData {
    closeSession.EntityData.YFilter = closeSession.YFilter
    closeSession.EntityData.YangName = "close-session"
    closeSession.EntityData.BundleName = "ietf"
    closeSession.EntityData.ParentYangName = "ietf-netconf"
    closeSession.EntityData.SegmentPath = "ietf-netconf:close-session"
    closeSession.EntityData.AbsolutePath = closeSession.EntityData.SegmentPath
    closeSession.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    closeSession.EntityData.NamespaceTable = ietf.GetNamespaces()
    closeSession.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    closeSession.EntityData.Children = types.NewOrderedMap()
    closeSession.EntityData.Leafs = types.NewOrderedMap()

    closeSession.EntityData.YListKeys = []string {}

    return &(closeSession.EntityData)
}

// KillSession
// Force the termination of a NETCONF session.
type KillSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input KillSession_Input
}

func (killSession *KillSession) GetEntityData() *types.CommonEntityData {
    killSession.EntityData.YFilter = killSession.YFilter
    killSession.EntityData.YangName = "kill-session"
    killSession.EntityData.BundleName = "ietf"
    killSession.EntityData.ParentYangName = "ietf-netconf"
    killSession.EntityData.SegmentPath = "ietf-netconf:kill-session"
    killSession.EntityData.AbsolutePath = killSession.EntityData.SegmentPath
    killSession.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    killSession.EntityData.NamespaceTable = ietf.GetNamespaces()
    killSession.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    killSession.EntityData.Children = types.NewOrderedMap()
    killSession.EntityData.Children.Append("input", types.YChild{"Input", &killSession.Input})
    killSession.EntityData.Leafs = types.NewOrderedMap()

    killSession.EntityData.YListKeys = []string {}

    return &(killSession.EntityData)
}

// KillSession_Input
type KillSession_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular session to kill. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    SessionId interface{}
}

func (input *KillSession_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "kill-session"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:kill-session/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", input.SessionId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Commit
// Commit the candidate configuration as the device's new
// current configuration.
type Commit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Commit_Input
}

func (commit *Commit) GetEntityData() *types.CommonEntityData {
    commit.EntityData.YFilter = commit.YFilter
    commit.EntityData.YangName = "commit"
    commit.EntityData.BundleName = "ietf"
    commit.EntityData.ParentYangName = "ietf-netconf"
    commit.EntityData.SegmentPath = "ietf-netconf:commit"
    commit.EntityData.AbsolutePath = commit.EntityData.SegmentPath
    commit.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    commit.EntityData.NamespaceTable = ietf.GetNamespaces()
    commit.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    commit.EntityData.Children = types.NewOrderedMap()
    commit.EntityData.Children.Append("input", types.YChild{"Input", &commit.Input})
    commit.EntityData.Leafs = types.NewOrderedMap()

    commit.EntityData.YListKeys = []string {}

    return &(commit.EntityData)
}

// Commit_Input
type Commit_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests a confirmed commit. The type is interface{}.
    Confirmed interface{}

    // The timeout interval for a confirmed commit. The type is interface{} with
    // range: 1..4294967295. Units are seconds. The default value is 600.
    ConfirmTimeout interface{}

    // This parameter is used to make a confirmed commit persistent.  A persistent
    // confirmed commit is not aborted if the NETCONF session terminates.  The
    // only way to abort a persistent confirmed commit is to let the timer expire,
    // or to use the <cancel-commit> operation.  The value of this parameter is a
    // token that must be given in the 'persist-id' parameter of <commit> or
    // <cancel-commit> operations in order to confirm or cancel the persistent
    // confirmed commit.  The token should be a random string. The type is string.
    Persist interface{}

    // This parameter is given in order to commit a persistent confirmed commit. 
    // The value must be equal to the value given in the 'persist' parameter to
    // the <commit> operation. If it does not match, the operation fails with an
    // 'invalid-value' error. The type is string.
    PersistId interface{}
}

func (input *Commit_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "commit"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:commit/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("confirmed", types.YLeaf{"Confirmed", input.Confirmed})
    input.EntityData.Leafs.Append("confirm-timeout", types.YLeaf{"ConfirmTimeout", input.ConfirmTimeout})
    input.EntityData.Leafs.Append("persist", types.YLeaf{"Persist", input.Persist})
    input.EntityData.Leafs.Append("persist-id", types.YLeaf{"PersistId", input.PersistId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// DiscardChanges
// Revert the candidate configuration to the current
// running configuration.
type DiscardChanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (discardChanges *DiscardChanges) GetEntityData() *types.CommonEntityData {
    discardChanges.EntityData.YFilter = discardChanges.YFilter
    discardChanges.EntityData.YangName = "discard-changes"
    discardChanges.EntityData.BundleName = "ietf"
    discardChanges.EntityData.ParentYangName = "ietf-netconf"
    discardChanges.EntityData.SegmentPath = "ietf-netconf:discard-changes"
    discardChanges.EntityData.AbsolutePath = discardChanges.EntityData.SegmentPath
    discardChanges.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    discardChanges.EntityData.NamespaceTable = ietf.GetNamespaces()
    discardChanges.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    discardChanges.EntityData.Children = types.NewOrderedMap()
    discardChanges.EntityData.Leafs = types.NewOrderedMap()

    discardChanges.EntityData.YListKeys = []string {}

    return &(discardChanges.EntityData)
}

// CancelCommit
// This operation is used to cancel an ongoing confirmed commit.
// If the confirmed commit is persistent, the parameter
// 'persist-id' must be given, and it must match the value of the
// 'persist' parameter.
type CancelCommit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input CancelCommit_Input
}

func (cancelCommit *CancelCommit) GetEntityData() *types.CommonEntityData {
    cancelCommit.EntityData.YFilter = cancelCommit.YFilter
    cancelCommit.EntityData.YangName = "cancel-commit"
    cancelCommit.EntityData.BundleName = "ietf"
    cancelCommit.EntityData.ParentYangName = "ietf-netconf"
    cancelCommit.EntityData.SegmentPath = "ietf-netconf:cancel-commit"
    cancelCommit.EntityData.AbsolutePath = cancelCommit.EntityData.SegmentPath
    cancelCommit.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cancelCommit.EntityData.NamespaceTable = ietf.GetNamespaces()
    cancelCommit.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cancelCommit.EntityData.Children = types.NewOrderedMap()
    cancelCommit.EntityData.Children.Append("input", types.YChild{"Input", &cancelCommit.Input})
    cancelCommit.EntityData.Leafs = types.NewOrderedMap()

    cancelCommit.EntityData.YListKeys = []string {}

    return &(cancelCommit.EntityData)
}

// CancelCommit_Input
type CancelCommit_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This parameter is given in order to cancel a persistent confirmed commit. 
    // The value must be equal to the value given in the 'persist' parameter to
    // the <commit> operation. If it does not match, the operation fails with an
    // 'invalid-value' error. The type is string.
    PersistId interface{}
}

func (input *CancelCommit_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "cancel-commit"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:cancel-commit/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("persist-id", types.YLeaf{"PersistId", input.PersistId})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Validate
// Validates the contents of the specified configuration.
type Validate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input Validate_Input
}

func (validate *Validate) GetEntityData() *types.CommonEntityData {
    validate.EntityData.YFilter = validate.YFilter
    validate.EntityData.YangName = "validate"
    validate.EntityData.BundleName = "ietf"
    validate.EntityData.ParentYangName = "ietf-netconf"
    validate.EntityData.SegmentPath = "ietf-netconf:validate"
    validate.EntityData.AbsolutePath = validate.EntityData.SegmentPath
    validate.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    validate.EntityData.NamespaceTable = ietf.GetNamespaces()
    validate.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    validate.EntityData.Children = types.NewOrderedMap()
    validate.EntityData.Children.Append("input", types.YChild{"Input", &validate.Input})
    validate.EntityData.Leafs = types.NewOrderedMap()

    validate.EntityData.YListKeys = []string {}

    return &(validate.EntityData)
}

// Validate_Input
type Validate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Particular configuration to validate.
    Source Validate_Input_Source
}

func (input *Validate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "ietf"
    input.EntityData.ParentYangName = "validate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "ietf-netconf:validate/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    input.EntityData.NamespaceTable = ietf.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Children.Append("source", types.YChild{"Source", &input.Source})
    input.EntityData.Leafs = types.NewOrderedMap()

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// Validate_Input_Source
// Particular configuration to validate.
type Validate_Input_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The candidate configuration is the config source. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config source. The type is interface{}.
    Running interface{}

    // The startup configuration is the config source. The type is interface{}.
    Startup interface{}

    // The URL-based configuration is the config source. The type is string.
    Url interface{}

    // Inline Config content: <config> element.  Represents an entire
    // configuration datastore, not a subset of the running datastore. The type is
    // string.
    Config interface{}
}

func (source *Validate_Input_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "ietf"
    source.EntityData.ParentYangName = "input"
    source.EntityData.SegmentPath = "source"
    source.EntityData.AbsolutePath = "ietf-netconf:validate/input/" + source.EntityData.SegmentPath
    source.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    source.EntityData.NamespaceTable = ietf.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("candidate", types.YLeaf{"Candidate", source.Candidate})
    source.EntityData.Leafs.Append("running", types.YLeaf{"Running", source.Running})
    source.EntityData.Leafs.Append("startup", types.YLeaf{"Startup", source.Startup})
    source.EntityData.Leafs.Append("url", types.YLeaf{"Url", source.Url})
    source.EntityData.Leafs.Append("config", types.YLeaf{"Config", source.Config})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

