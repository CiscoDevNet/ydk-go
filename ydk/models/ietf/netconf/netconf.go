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

// ErrorSeverityType represents NETCONF Error Severity
type ErrorSeverityType string

const (
    // Error severity
    ErrorSeverityType_error ErrorSeverityType = "error"

    // Warning severity
    ErrorSeverityType_warning ErrorSeverityType = "warning"
)

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
    EditOperationType_delete EditOperationType = "delete"

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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input GetConfig_Input

    
    Output GetConfig_Output
}

func (getConfig *GetConfig) GetFilter() yfilter.YFilter { return getConfig.YFilter }

func (getConfig *GetConfig) SetFilter(yf yfilter.YFilter) { getConfig.YFilter = yf }

func (getConfig *GetConfig) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (getConfig *GetConfig) GetSegmentPath() string {
    return "ietf-netconf:get-config"
}

func (getConfig *GetConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &getConfig.Input
    }
    if childYangName == "output" {
        return &getConfig.Output
    }
    return nil
}

func (getConfig *GetConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &getConfig.Input
    children["output"] = &getConfig.Output
    return children
}

func (getConfig *GetConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (getConfig *GetConfig) GetBundleName() string { return "ietf" }

func (getConfig *GetConfig) GetYangName() string { return "get-config" }

func (getConfig *GetConfig) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (getConfig *GetConfig) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (getConfig *GetConfig) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (getConfig *GetConfig) SetParent(parent types.Entity) { getConfig.parent = parent }

func (getConfig *GetConfig) GetParent() types.Entity { return getConfig.parent }

func (getConfig *GetConfig) GetParentYangName() string { return "ietf-netconf" }

// GetConfig_Input
type GetConfig_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subtree or XPath filter to use. The type is string.
    Filter interface{}

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}

    // Particular configuration to retrieve.
    Source GetConfig_Input_Source
}

func (input *GetConfig_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *GetConfig_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *GetConfig_Input) GetGoName(yname string) string {
    if yname == "filter" { return "Filter" }
    if yname == "with-defaults" { return "WithDefaults" }
    if yname == "source" { return "Source" }
    return ""
}

func (input *GetConfig_Input) GetSegmentPath() string {
    return "input"
}

func (input *GetConfig_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &input.Source
    }
    return nil
}

func (input *GetConfig_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &input.Source
    return children
}

func (input *GetConfig_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter"] = input.Filter
    leafs["with-defaults"] = input.WithDefaults
    return leafs
}

func (input *GetConfig_Input) GetBundleName() string { return "ietf" }

func (input *GetConfig_Input) GetYangName() string { return "input" }

func (input *GetConfig_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *GetConfig_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *GetConfig_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *GetConfig_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *GetConfig_Input) GetParent() types.Entity { return input.parent }

func (input *GetConfig_Input) GetParentYangName() string { return "get-config" }

// GetConfig_Input_Source
// Particular configuration to retrieve.
type GetConfig_Input_Source struct {
    parent types.Entity
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

func (source *GetConfig_Input_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *GetConfig_Input_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *GetConfig_Input_Source) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    return ""
}

func (source *GetConfig_Input_Source) GetSegmentPath() string {
    return "source"
}

func (source *GetConfig_Input_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *GetConfig_Input_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *GetConfig_Input_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = source.Candidate
    leafs["running"] = source.Running
    leafs["startup"] = source.Startup
    return leafs
}

func (source *GetConfig_Input_Source) GetBundleName() string { return "ietf" }

func (source *GetConfig_Input_Source) GetYangName() string { return "source" }

func (source *GetConfig_Input_Source) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (source *GetConfig_Input_Source) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (source *GetConfig_Input_Source) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (source *GetConfig_Input_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *GetConfig_Input_Source) GetParent() types.Entity { return source.parent }

func (source *GetConfig_Input_Source) GetParentYangName() string { return "input" }

// GetConfig_Output
type GetConfig_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Copy of the source datastore subset that matched the filter criteria (if
    // any).  An empty data container indicates that the request did not produce
    // any results. The type is string.
    Data interface{}
}

func (output *GetConfig_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *GetConfig_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *GetConfig_Output) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    return ""
}

func (output *GetConfig_Output) GetSegmentPath() string {
    return "output"
}

func (output *GetConfig_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *GetConfig_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *GetConfig_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data"] = output.Data
    return leafs
}

func (output *GetConfig_Output) GetBundleName() string { return "ietf" }

func (output *GetConfig_Output) GetYangName() string { return "output" }

func (output *GetConfig_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *GetConfig_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *GetConfig_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *GetConfig_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *GetConfig_Output) GetParent() types.Entity { return output.parent }

func (output *GetConfig_Output) GetParentYangName() string { return "get-config" }

// EditConfig
// The <edit-config> operation loads all or part of a specified
// configuration to the specified target configuration.
type EditConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input EditConfig_Input
}

func (editConfig *EditConfig) GetFilter() yfilter.YFilter { return editConfig.YFilter }

func (editConfig *EditConfig) SetFilter(yf yfilter.YFilter) { editConfig.YFilter = yf }

func (editConfig *EditConfig) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (editConfig *EditConfig) GetSegmentPath() string {
    return "ietf-netconf:edit-config"
}

func (editConfig *EditConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &editConfig.Input
    }
    return nil
}

func (editConfig *EditConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &editConfig.Input
    return children
}

func (editConfig *EditConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (editConfig *EditConfig) GetBundleName() string { return "ietf" }

func (editConfig *EditConfig) GetYangName() string { return "edit-config" }

func (editConfig *EditConfig) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (editConfig *EditConfig) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (editConfig *EditConfig) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (editConfig *EditConfig) SetParent(parent types.Entity) { editConfig.parent = parent }

func (editConfig *EditConfig) GetParent() types.Entity { return editConfig.parent }

func (editConfig *EditConfig) GetParentYangName() string { return "ietf-netconf" }

// EditConfig_Input
type EditConfig_Input struct {
    parent types.Entity
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

func (input *EditConfig_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *EditConfig_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *EditConfig_Input) GetGoName(yname string) string {
    if yname == "default-operation" { return "DefaultOperation" }
    if yname == "test-option" { return "TestOption" }
    if yname == "error-option" { return "ErrorOption" }
    if yname == "config" { return "Config" }
    if yname == "url" { return "Url" }
    if yname == "target" { return "Target" }
    return ""
}

func (input *EditConfig_Input) GetSegmentPath() string {
    return "input"
}

func (input *EditConfig_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    return nil
}

func (input *EditConfig_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    return children
}

func (input *EditConfig_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-operation"] = input.DefaultOperation
    leafs["test-option"] = input.TestOption
    leafs["error-option"] = input.ErrorOption
    leafs["config"] = input.Config
    leafs["url"] = input.Url
    return leafs
}

func (input *EditConfig_Input) GetBundleName() string { return "ietf" }

func (input *EditConfig_Input) GetYangName() string { return "input" }

func (input *EditConfig_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *EditConfig_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *EditConfig_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *EditConfig_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *EditConfig_Input) GetParent() types.Entity { return input.parent }

func (input *EditConfig_Input) GetParentYangName() string { return "edit-config" }

// EditConfig_Input_Target
// Particular configuration to edit.
type EditConfig_Input_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config source. The type is interface{}.
    Running interface{}
}

func (target *EditConfig_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *EditConfig_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *EditConfig_Input_Target) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    return ""
}

func (target *EditConfig_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *EditConfig_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *EditConfig_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *EditConfig_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = target.Candidate
    leafs["running"] = target.Running
    return leafs
}

func (target *EditConfig_Input_Target) GetBundleName() string { return "ietf" }

func (target *EditConfig_Input_Target) GetYangName() string { return "target" }

func (target *EditConfig_Input_Target) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (target *EditConfig_Input_Target) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (target *EditConfig_Input_Target) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (target *EditConfig_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *EditConfig_Input_Target) GetParent() types.Entity { return target.parent }

func (target *EditConfig_Input_Target) GetParentYangName() string { return "input" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CopyConfig_Input
}

func (copyConfig *CopyConfig) GetFilter() yfilter.YFilter { return copyConfig.YFilter }

func (copyConfig *CopyConfig) SetFilter(yf yfilter.YFilter) { copyConfig.YFilter = yf }

func (copyConfig *CopyConfig) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (copyConfig *CopyConfig) GetSegmentPath() string {
    return "ietf-netconf:copy-config"
}

func (copyConfig *CopyConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &copyConfig.Input
    }
    return nil
}

func (copyConfig *CopyConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &copyConfig.Input
    return children
}

func (copyConfig *CopyConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (copyConfig *CopyConfig) GetBundleName() string { return "ietf" }

func (copyConfig *CopyConfig) GetYangName() string { return "copy-config" }

func (copyConfig *CopyConfig) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (copyConfig *CopyConfig) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (copyConfig *CopyConfig) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (copyConfig *CopyConfig) SetParent(parent types.Entity) { copyConfig.parent = parent }

func (copyConfig *CopyConfig) GetParent() types.Entity { return copyConfig.parent }

func (copyConfig *CopyConfig) GetParentYangName() string { return "ietf-netconf" }

// CopyConfig_Input
type CopyConfig_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}

    // Particular configuration to copy to.
    Target CopyConfig_Input_Target

    // Particular configuration to copy from.
    Source CopyConfig_Input_Source
}

func (input *CopyConfig_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CopyConfig_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CopyConfig_Input) GetGoName(yname string) string {
    if yname == "with-defaults" { return "WithDefaults" }
    if yname == "target" { return "Target" }
    if yname == "source" { return "Source" }
    return ""
}

func (input *CopyConfig_Input) GetSegmentPath() string {
    return "input"
}

func (input *CopyConfig_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    if childYangName == "source" {
        return &input.Source
    }
    return nil
}

func (input *CopyConfig_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    children["source"] = &input.Source
    return children
}

func (input *CopyConfig_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["with-defaults"] = input.WithDefaults
    return leafs
}

func (input *CopyConfig_Input) GetBundleName() string { return "ietf" }

func (input *CopyConfig_Input) GetYangName() string { return "input" }

func (input *CopyConfig_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *CopyConfig_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *CopyConfig_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *CopyConfig_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CopyConfig_Input) GetParent() types.Entity { return input.parent }

func (input *CopyConfig_Input) GetParentYangName() string { return "copy-config" }

// CopyConfig_Input_Target
// Particular configuration to copy to.
type CopyConfig_Input_Target struct {
    parent types.Entity
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

func (target *CopyConfig_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *CopyConfig_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *CopyConfig_Input_Target) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    if yname == "url" { return "Url" }
    return ""
}

func (target *CopyConfig_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *CopyConfig_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *CopyConfig_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *CopyConfig_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = target.Candidate
    leafs["running"] = target.Running
    leafs["startup"] = target.Startup
    leafs["url"] = target.Url
    return leafs
}

func (target *CopyConfig_Input_Target) GetBundleName() string { return "ietf" }

func (target *CopyConfig_Input_Target) GetYangName() string { return "target" }

func (target *CopyConfig_Input_Target) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (target *CopyConfig_Input_Target) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (target *CopyConfig_Input_Target) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (target *CopyConfig_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *CopyConfig_Input_Target) GetParent() types.Entity { return target.parent }

func (target *CopyConfig_Input_Target) GetParentYangName() string { return "input" }

// CopyConfig_Input_Source
// Particular configuration to copy from.
type CopyConfig_Input_Source struct {
    parent types.Entity
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

func (source *CopyConfig_Input_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *CopyConfig_Input_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *CopyConfig_Input_Source) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    if yname == "url" { return "Url" }
    if yname == "config" { return "Config" }
    return ""
}

func (source *CopyConfig_Input_Source) GetSegmentPath() string {
    return "source"
}

func (source *CopyConfig_Input_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *CopyConfig_Input_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *CopyConfig_Input_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = source.Candidate
    leafs["running"] = source.Running
    leafs["startup"] = source.Startup
    leafs["url"] = source.Url
    leafs["config"] = source.Config
    return leafs
}

func (source *CopyConfig_Input_Source) GetBundleName() string { return "ietf" }

func (source *CopyConfig_Input_Source) GetYangName() string { return "source" }

func (source *CopyConfig_Input_Source) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (source *CopyConfig_Input_Source) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (source *CopyConfig_Input_Source) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (source *CopyConfig_Input_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *CopyConfig_Input_Source) GetParent() types.Entity { return source.parent }

func (source *CopyConfig_Input_Source) GetParentYangName() string { return "input" }

// DeleteConfig
// Delete a configuration datastore.
type DeleteConfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input DeleteConfig_Input
}

func (deleteConfig *DeleteConfig) GetFilter() yfilter.YFilter { return deleteConfig.YFilter }

func (deleteConfig *DeleteConfig) SetFilter(yf yfilter.YFilter) { deleteConfig.YFilter = yf }

func (deleteConfig *DeleteConfig) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (deleteConfig *DeleteConfig) GetSegmentPath() string {
    return "ietf-netconf:delete-config"
}

func (deleteConfig *DeleteConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &deleteConfig.Input
    }
    return nil
}

func (deleteConfig *DeleteConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &deleteConfig.Input
    return children
}

func (deleteConfig *DeleteConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (deleteConfig *DeleteConfig) GetBundleName() string { return "ietf" }

func (deleteConfig *DeleteConfig) GetYangName() string { return "delete-config" }

func (deleteConfig *DeleteConfig) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (deleteConfig *DeleteConfig) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (deleteConfig *DeleteConfig) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (deleteConfig *DeleteConfig) SetParent(parent types.Entity) { deleteConfig.parent = parent }

func (deleteConfig *DeleteConfig) GetParent() types.Entity { return deleteConfig.parent }

func (deleteConfig *DeleteConfig) GetParentYangName() string { return "ietf-netconf" }

// DeleteConfig_Input
type DeleteConfig_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular configuration to delete.
    Target DeleteConfig_Input_Target
}

func (input *DeleteConfig_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *DeleteConfig_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *DeleteConfig_Input) GetGoName(yname string) string {
    if yname == "target" { return "Target" }
    return ""
}

func (input *DeleteConfig_Input) GetSegmentPath() string {
    return "input"
}

func (input *DeleteConfig_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    return nil
}

func (input *DeleteConfig_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    return children
}

func (input *DeleteConfig_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *DeleteConfig_Input) GetBundleName() string { return "ietf" }

func (input *DeleteConfig_Input) GetYangName() string { return "input" }

func (input *DeleteConfig_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *DeleteConfig_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *DeleteConfig_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *DeleteConfig_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *DeleteConfig_Input) GetParent() types.Entity { return input.parent }

func (input *DeleteConfig_Input) GetParentYangName() string { return "delete-config" }

// DeleteConfig_Input_Target
// Particular configuration to delete.
type DeleteConfig_Input_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}

    // The URL-based configuration is the config target. The type is string.
    Url interface{}
}

func (target *DeleteConfig_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *DeleteConfig_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *DeleteConfig_Input_Target) GetGoName(yname string) string {
    if yname == "startup" { return "Startup" }
    if yname == "url" { return "Url" }
    return ""
}

func (target *DeleteConfig_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *DeleteConfig_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *DeleteConfig_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *DeleteConfig_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["startup"] = target.Startup
    leafs["url"] = target.Url
    return leafs
}

func (target *DeleteConfig_Input_Target) GetBundleName() string { return "ietf" }

func (target *DeleteConfig_Input_Target) GetYangName() string { return "target" }

func (target *DeleteConfig_Input_Target) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (target *DeleteConfig_Input_Target) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (target *DeleteConfig_Input_Target) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (target *DeleteConfig_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *DeleteConfig_Input_Target) GetParent() types.Entity { return target.parent }

func (target *DeleteConfig_Input_Target) GetParentYangName() string { return "input" }

// Lock
// The lock operation allows the client to lock the configuration
// system of a device.
type Lock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Lock_Input
}

func (lock *Lock) GetFilter() yfilter.YFilter { return lock.YFilter }

func (lock *Lock) SetFilter(yf yfilter.YFilter) { lock.YFilter = yf }

func (lock *Lock) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (lock *Lock) GetSegmentPath() string {
    return "ietf-netconf:lock"
}

func (lock *Lock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &lock.Input
    }
    return nil
}

func (lock *Lock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &lock.Input
    return children
}

func (lock *Lock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lock *Lock) GetBundleName() string { return "ietf" }

func (lock *Lock) GetYangName() string { return "lock" }

func (lock *Lock) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (lock *Lock) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (lock *Lock) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (lock *Lock) SetParent(parent types.Entity) { lock.parent = parent }

func (lock *Lock) GetParent() types.Entity { return lock.parent }

func (lock *Lock) GetParentYangName() string { return "ietf-netconf" }

// Lock_Input
type Lock_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular configuration to lock.
    Target Lock_Input_Target
}

func (input *Lock_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Lock_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Lock_Input) GetGoName(yname string) string {
    if yname == "target" { return "Target" }
    return ""
}

func (input *Lock_Input) GetSegmentPath() string {
    return "input"
}

func (input *Lock_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    return nil
}

func (input *Lock_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    return children
}

func (input *Lock_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *Lock_Input) GetBundleName() string { return "ietf" }

func (input *Lock_Input) GetYangName() string { return "input" }

func (input *Lock_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *Lock_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *Lock_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *Lock_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Lock_Input) GetParent() types.Entity { return input.parent }

func (input *Lock_Input) GetParentYangName() string { return "lock" }

// Lock_Input_Target
// Particular configuration to lock.
type Lock_Input_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config target. The type is interface{}.
    Running interface{}

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}
}

func (target *Lock_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *Lock_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *Lock_Input_Target) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    return ""
}

func (target *Lock_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *Lock_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *Lock_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *Lock_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = target.Candidate
    leafs["running"] = target.Running
    leafs["startup"] = target.Startup
    return leafs
}

func (target *Lock_Input_Target) GetBundleName() string { return "ietf" }

func (target *Lock_Input_Target) GetYangName() string { return "target" }

func (target *Lock_Input_Target) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (target *Lock_Input_Target) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (target *Lock_Input_Target) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (target *Lock_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *Lock_Input_Target) GetParent() types.Entity { return target.parent }

func (target *Lock_Input_Target) GetParentYangName() string { return "input" }

// Unlock
// The unlock operation is used to release a configuration lock,
// previously obtained with the 'lock' operation.
type Unlock struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Unlock_Input
}

func (unlock *Unlock) GetFilter() yfilter.YFilter { return unlock.YFilter }

func (unlock *Unlock) SetFilter(yf yfilter.YFilter) { unlock.YFilter = yf }

func (unlock *Unlock) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (unlock *Unlock) GetSegmentPath() string {
    return "ietf-netconf:unlock"
}

func (unlock *Unlock) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &unlock.Input
    }
    return nil
}

func (unlock *Unlock) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &unlock.Input
    return children
}

func (unlock *Unlock) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (unlock *Unlock) GetBundleName() string { return "ietf" }

func (unlock *Unlock) GetYangName() string { return "unlock" }

func (unlock *Unlock) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (unlock *Unlock) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (unlock *Unlock) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (unlock *Unlock) SetParent(parent types.Entity) { unlock.parent = parent }

func (unlock *Unlock) GetParent() types.Entity { return unlock.parent }

func (unlock *Unlock) GetParentYangName() string { return "ietf-netconf" }

// Unlock_Input
type Unlock_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular configuration to unlock.
    Target Unlock_Input_Target
}

func (input *Unlock_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Unlock_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Unlock_Input) GetGoName(yname string) string {
    if yname == "target" { return "Target" }
    return ""
}

func (input *Unlock_Input) GetSegmentPath() string {
    return "input"
}

func (input *Unlock_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        return &input.Target
    }
    return nil
}

func (input *Unlock_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target"] = &input.Target
    return children
}

func (input *Unlock_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *Unlock_Input) GetBundleName() string { return "ietf" }

func (input *Unlock_Input) GetYangName() string { return "input" }

func (input *Unlock_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *Unlock_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *Unlock_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *Unlock_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Unlock_Input) GetParent() types.Entity { return input.parent }

func (input *Unlock_Input) GetParentYangName() string { return "unlock" }

// Unlock_Input_Target
// Particular configuration to unlock.
type Unlock_Input_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The candidate configuration is the config target. The type is interface{}.
    Candidate interface{}

    // The running configuration is the config target. The type is interface{}.
    Running interface{}

    // The startup configuration is the config target. The type is interface{}.
    Startup interface{}
}

func (target *Unlock_Input_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *Unlock_Input_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *Unlock_Input_Target) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    return ""
}

func (target *Unlock_Input_Target) GetSegmentPath() string {
    return "target"
}

func (target *Unlock_Input_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (target *Unlock_Input_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (target *Unlock_Input_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = target.Candidate
    leafs["running"] = target.Running
    leafs["startup"] = target.Startup
    return leafs
}

func (target *Unlock_Input_Target) GetBundleName() string { return "ietf" }

func (target *Unlock_Input_Target) GetYangName() string { return "target" }

func (target *Unlock_Input_Target) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (target *Unlock_Input_Target) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (target *Unlock_Input_Target) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (target *Unlock_Input_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *Unlock_Input_Target) GetParent() types.Entity { return target.parent }

func (target *Unlock_Input_Target) GetParentYangName() string { return "input" }

// Get
// Retrieve running configuration and device state information.
type Get struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Get_Input

    
    Output Get_Output
}

func (get *Get) GetFilter() yfilter.YFilter { return get.YFilter }

func (get *Get) SetFilter(yf yfilter.YFilter) { get.YFilter = yf }

func (get *Get) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (get *Get) GetSegmentPath() string {
    return "ietf-netconf:get"
}

func (get *Get) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &get.Input
    }
    if childYangName == "output" {
        return &get.Output
    }
    return nil
}

func (get *Get) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &get.Input
    children["output"] = &get.Output
    return children
}

func (get *Get) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (get *Get) GetBundleName() string { return "ietf" }

func (get *Get) GetYangName() string { return "get" }

func (get *Get) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (get *Get) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (get *Get) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (get *Get) SetParent(parent types.Entity) { get.parent = parent }

func (get *Get) GetParent() types.Entity { return get.parent }

func (get *Get) GetParentYangName() string { return "ietf-netconf" }

// Get_Input
type Get_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This parameter specifies the portion of the system configuration and state
    // data to retrieve. The type is string.
    Filter interface{}

    // The explicit defaults processing mode requested. The type is
    // WithDefaultsMode.
    WithDefaults interface{}
}

func (input *Get_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Get_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Get_Input) GetGoName(yname string) string {
    if yname == "filter" { return "Filter" }
    if yname == "with-defaults" { return "WithDefaults" }
    return ""
}

func (input *Get_Input) GetSegmentPath() string {
    return "input"
}

func (input *Get_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Get_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Get_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter"] = input.Filter
    leafs["with-defaults"] = input.WithDefaults
    return leafs
}

func (input *Get_Input) GetBundleName() string { return "ietf" }

func (input *Get_Input) GetYangName() string { return "input" }

func (input *Get_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *Get_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *Get_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *Get_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Get_Input) GetParent() types.Entity { return input.parent }

func (input *Get_Input) GetParentYangName() string { return "get" }

// Get_Output
type Get_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Copy of the running datastore subset and/or state data that matched the
    // filter criteria (if any). An empty data container indicates that the
    // request did not produce any results. The type is string.
    Data interface{}
}

func (output *Get_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Get_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Get_Output) GetGoName(yname string) string {
    if yname == "data" { return "Data" }
    return ""
}

func (output *Get_Output) GetSegmentPath() string {
    return "output"
}

func (output *Get_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Get_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Get_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data"] = output.Data
    return leafs
}

func (output *Get_Output) GetBundleName() string { return "ietf" }

func (output *Get_Output) GetYangName() string { return "output" }

func (output *Get_Output) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (output *Get_Output) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (output *Get_Output) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (output *Get_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Get_Output) GetParent() types.Entity { return output.parent }

func (output *Get_Output) GetParentYangName() string { return "get" }

// CloseSession
// Request graceful termination of a NETCONF session.
type CloseSession struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (closeSession *CloseSession) GetFilter() yfilter.YFilter { return closeSession.YFilter }

func (closeSession *CloseSession) SetFilter(yf yfilter.YFilter) { closeSession.YFilter = yf }

func (closeSession *CloseSession) GetGoName(yname string) string {
    return ""
}

func (closeSession *CloseSession) GetSegmentPath() string {
    return "ietf-netconf:close-session"
}

func (closeSession *CloseSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (closeSession *CloseSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (closeSession *CloseSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (closeSession *CloseSession) GetBundleName() string { return "ietf" }

func (closeSession *CloseSession) GetYangName() string { return "close-session" }

func (closeSession *CloseSession) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (closeSession *CloseSession) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (closeSession *CloseSession) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (closeSession *CloseSession) SetParent(parent types.Entity) { closeSession.parent = parent }

func (closeSession *CloseSession) GetParent() types.Entity { return closeSession.parent }

func (closeSession *CloseSession) GetParentYangName() string { return "ietf-netconf" }

// KillSession
// Force the termination of a NETCONF session.
type KillSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input KillSession_Input
}

func (killSession *KillSession) GetFilter() yfilter.YFilter { return killSession.YFilter }

func (killSession *KillSession) SetFilter(yf yfilter.YFilter) { killSession.YFilter = yf }

func (killSession *KillSession) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (killSession *KillSession) GetSegmentPath() string {
    return "ietf-netconf:kill-session"
}

func (killSession *KillSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &killSession.Input
    }
    return nil
}

func (killSession *KillSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &killSession.Input
    return children
}

func (killSession *KillSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (killSession *KillSession) GetBundleName() string { return "ietf" }

func (killSession *KillSession) GetYangName() string { return "kill-session" }

func (killSession *KillSession) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (killSession *KillSession) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (killSession *KillSession) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (killSession *KillSession) SetParent(parent types.Entity) { killSession.parent = parent }

func (killSession *KillSession) GetParent() types.Entity { return killSession.parent }

func (killSession *KillSession) GetParentYangName() string { return "ietf-netconf" }

// KillSession_Input
type KillSession_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular session to kill. The type is interface{} with range:
    // 1..4294967295. This attribute is mandatory.
    SessionId interface{}
}

func (input *KillSession_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *KillSession_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *KillSession_Input) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    return ""
}

func (input *KillSession_Input) GetSegmentPath() string {
    return "input"
}

func (input *KillSession_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *KillSession_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *KillSession_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = input.SessionId
    return leafs
}

func (input *KillSession_Input) GetBundleName() string { return "ietf" }

func (input *KillSession_Input) GetYangName() string { return "input" }

func (input *KillSession_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *KillSession_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *KillSession_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *KillSession_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *KillSession_Input) GetParent() types.Entity { return input.parent }

func (input *KillSession_Input) GetParentYangName() string { return "kill-session" }

// Commit
// Commit the candidate configuration as the device's new
// current configuration.
type Commit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Commit_Input
}

func (commit *Commit) GetFilter() yfilter.YFilter { return commit.YFilter }

func (commit *Commit) SetFilter(yf yfilter.YFilter) { commit.YFilter = yf }

func (commit *Commit) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (commit *Commit) GetSegmentPath() string {
    return "ietf-netconf:commit"
}

func (commit *Commit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &commit.Input
    }
    return nil
}

func (commit *Commit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &commit.Input
    return children
}

func (commit *Commit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (commit *Commit) GetBundleName() string { return "ietf" }

func (commit *Commit) GetYangName() string { return "commit" }

func (commit *Commit) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (commit *Commit) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (commit *Commit) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (commit *Commit) SetParent(parent types.Entity) { commit.parent = parent }

func (commit *Commit) GetParent() types.Entity { return commit.parent }

func (commit *Commit) GetParentYangName() string { return "ietf-netconf" }

// Commit_Input
type Commit_Input struct {
    parent types.Entity
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

func (input *Commit_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Commit_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Commit_Input) GetGoName(yname string) string {
    if yname == "confirmed" { return "Confirmed" }
    if yname == "confirm-timeout" { return "ConfirmTimeout" }
    if yname == "persist" { return "Persist" }
    if yname == "persist-id" { return "PersistId" }
    return ""
}

func (input *Commit_Input) GetSegmentPath() string {
    return "input"
}

func (input *Commit_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *Commit_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *Commit_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["confirmed"] = input.Confirmed
    leafs["confirm-timeout"] = input.ConfirmTimeout
    leafs["persist"] = input.Persist
    leafs["persist-id"] = input.PersistId
    return leafs
}

func (input *Commit_Input) GetBundleName() string { return "ietf" }

func (input *Commit_Input) GetYangName() string { return "input" }

func (input *Commit_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *Commit_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *Commit_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *Commit_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Commit_Input) GetParent() types.Entity { return input.parent }

func (input *Commit_Input) GetParentYangName() string { return "commit" }

// DiscardChanges
// Revert the candidate configuration to the current
// running configuration.
type DiscardChanges struct {
    parent types.Entity
    YFilter yfilter.YFilter
}

func (discardChanges *DiscardChanges) GetFilter() yfilter.YFilter { return discardChanges.YFilter }

func (discardChanges *DiscardChanges) SetFilter(yf yfilter.YFilter) { discardChanges.YFilter = yf }

func (discardChanges *DiscardChanges) GetGoName(yname string) string {
    return ""
}

func (discardChanges *DiscardChanges) GetSegmentPath() string {
    return "ietf-netconf:discard-changes"
}

func (discardChanges *DiscardChanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discardChanges *DiscardChanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discardChanges *DiscardChanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (discardChanges *DiscardChanges) GetBundleName() string { return "ietf" }

func (discardChanges *DiscardChanges) GetYangName() string { return "discard-changes" }

func (discardChanges *DiscardChanges) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (discardChanges *DiscardChanges) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (discardChanges *DiscardChanges) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (discardChanges *DiscardChanges) SetParent(parent types.Entity) { discardChanges.parent = parent }

func (discardChanges *DiscardChanges) GetParent() types.Entity { return discardChanges.parent }

func (discardChanges *DiscardChanges) GetParentYangName() string { return "ietf-netconf" }

// CancelCommit
// This operation is used to cancel an ongoing confirmed commit.
// If the confirmed commit is persistent, the parameter
// 'persist-id' must be given, and it must match the value of the
// 'persist' parameter.
type CancelCommit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input CancelCommit_Input
}

func (cancelCommit *CancelCommit) GetFilter() yfilter.YFilter { return cancelCommit.YFilter }

func (cancelCommit *CancelCommit) SetFilter(yf yfilter.YFilter) { cancelCommit.YFilter = yf }

func (cancelCommit *CancelCommit) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (cancelCommit *CancelCommit) GetSegmentPath() string {
    return "ietf-netconf:cancel-commit"
}

func (cancelCommit *CancelCommit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &cancelCommit.Input
    }
    return nil
}

func (cancelCommit *CancelCommit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &cancelCommit.Input
    return children
}

func (cancelCommit *CancelCommit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cancelCommit *CancelCommit) GetBundleName() string { return "ietf" }

func (cancelCommit *CancelCommit) GetYangName() string { return "cancel-commit" }

func (cancelCommit *CancelCommit) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cancelCommit *CancelCommit) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cancelCommit *CancelCommit) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cancelCommit *CancelCommit) SetParent(parent types.Entity) { cancelCommit.parent = parent }

func (cancelCommit *CancelCommit) GetParent() types.Entity { return cancelCommit.parent }

func (cancelCommit *CancelCommit) GetParentYangName() string { return "ietf-netconf" }

// CancelCommit_Input
type CancelCommit_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This parameter is given in order to cancel a persistent confirmed commit. 
    // The value must be equal to the value given in the 'persist' parameter to
    // the <commit> operation. If it does not match, the operation fails with an
    // 'invalid-value' error. The type is string.
    PersistId interface{}
}

func (input *CancelCommit_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *CancelCommit_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *CancelCommit_Input) GetGoName(yname string) string {
    if yname == "persist-id" { return "PersistId" }
    return ""
}

func (input *CancelCommit_Input) GetSegmentPath() string {
    return "input"
}

func (input *CancelCommit_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *CancelCommit_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *CancelCommit_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["persist-id"] = input.PersistId
    return leafs
}

func (input *CancelCommit_Input) GetBundleName() string { return "ietf" }

func (input *CancelCommit_Input) GetYangName() string { return "input" }

func (input *CancelCommit_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *CancelCommit_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *CancelCommit_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *CancelCommit_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *CancelCommit_Input) GetParent() types.Entity { return input.parent }

func (input *CancelCommit_Input) GetParentYangName() string { return "cancel-commit" }

// Validate
// Validates the contents of the specified configuration.
type Validate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input Validate_Input
}

func (validate *Validate) GetFilter() yfilter.YFilter { return validate.YFilter }

func (validate *Validate) SetFilter(yf yfilter.YFilter) { validate.YFilter = yf }

func (validate *Validate) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    return ""
}

func (validate *Validate) GetSegmentPath() string {
    return "ietf-netconf:validate"
}

func (validate *Validate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &validate.Input
    }
    return nil
}

func (validate *Validate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &validate.Input
    return children
}

func (validate *Validate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (validate *Validate) GetBundleName() string { return "ietf" }

func (validate *Validate) GetYangName() string { return "validate" }

func (validate *Validate) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (validate *Validate) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (validate *Validate) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (validate *Validate) SetParent(parent types.Entity) { validate.parent = parent }

func (validate *Validate) GetParent() types.Entity { return validate.parent }

func (validate *Validate) GetParentYangName() string { return "ietf-netconf" }

// Validate_Input
type Validate_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Particular configuration to validate.
    Source Validate_Input_Source
}

func (input *Validate_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *Validate_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *Validate_Input) GetGoName(yname string) string {
    if yname == "source" { return "Source" }
    return ""
}

func (input *Validate_Input) GetSegmentPath() string {
    return "input"
}

func (input *Validate_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &input.Source
    }
    return nil
}

func (input *Validate_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &input.Source
    return children
}

func (input *Validate_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *Validate_Input) GetBundleName() string { return "ietf" }

func (input *Validate_Input) GetYangName() string { return "input" }

func (input *Validate_Input) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (input *Validate_Input) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (input *Validate_Input) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (input *Validate_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *Validate_Input) GetParent() types.Entity { return input.parent }

func (input *Validate_Input) GetParentYangName() string { return "validate" }

// Validate_Input_Source
// Particular configuration to validate.
type Validate_Input_Source struct {
    parent types.Entity
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

func (source *Validate_Input_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Validate_Input_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Validate_Input_Source) GetGoName(yname string) string {
    if yname == "candidate" { return "Candidate" }
    if yname == "running" { return "Running" }
    if yname == "startup" { return "Startup" }
    if yname == "url" { return "Url" }
    if yname == "config" { return "Config" }
    return ""
}

func (source *Validate_Input_Source) GetSegmentPath() string {
    return "source"
}

func (source *Validate_Input_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Validate_Input_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Validate_Input_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["candidate"] = source.Candidate
    leafs["running"] = source.Running
    leafs["startup"] = source.Startup
    leafs["url"] = source.Url
    leafs["config"] = source.Config
    return leafs
}

func (source *Validate_Input_Source) GetBundleName() string { return "ietf" }

func (source *Validate_Input_Source) GetYangName() string { return "source" }

func (source *Validate_Input_Source) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (source *Validate_Input_Source) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (source *Validate_Input_Source) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (source *Validate_Input_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Validate_Input_Source) GetParent() types.Entity { return source.parent }

func (source *Validate_Input_Source) GetParentYangName() string { return "input" }

