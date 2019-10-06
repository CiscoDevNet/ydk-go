// This module contains a collection of YANG 
// definitions for configuration of streaming telemetry.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package mdt_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mdt_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mdt-cfg mdt-config-data}", reflect.TypeOf(MdtConfigData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mdt-cfg:mdt-config-data", reflect.TypeOf(MdtConfigData{}))
}

// MdtXfrmOperator represents Supported operator types
type MdtXfrmOperator string

const (
    // Default operator
    MdtXfrmOperator_operator_none MdtXfrmOperator = "operator-none"

    // Equal operator
    MdtXfrmOperator_eq MdtXfrmOperator = "eq"

    // Not equal operator
    MdtXfrmOperator_ne MdtXfrmOperator = "ne"

    // Greater than operator
    MdtXfrmOperator_gt MdtXfrmOperator = "gt"

    // Greater than or equal operator
    MdtXfrmOperator_ge MdtXfrmOperator = "ge"

    // Less than operator
    MdtXfrmOperator_lt MdtXfrmOperator = "lt"

    // Less than or equal operator
    MdtXfrmOperator_le MdtXfrmOperator = "le"
)

// MdtXfrmOpType represents Types of subscription transform operations.
type MdtXfrmOpType string

const (
    // Indicates that operation type is 
    // sub-record
    MdtXfrmOpType_sub_record MdtXfrmOpType = "sub-record"

    // Indicates that operation type is delta
    MdtXfrmOpType_delta MdtXfrmOpType = "delta"
)

// MdtXfrmLogicOp represents logical operations.
type MdtXfrmLogicOp string

const (
    // Indicates that no logical operation 
    // is selected
    MdtXfrmLogicOp_mdt_xfrm_lop_none MdtXfrmLogicOp = "mdt-xfrm-lop-none"

    // Indicates that logical operation 
    // is 'and' 
    MdtXfrmLogicOp_and MdtXfrmLogicOp = "and"

    // Indicates that logical operation is 'or' 
    MdtXfrmLogicOp_or MdtXfrmLogicOp = "or"
)

// MdtXfrmAttrType represents Types of subscription transform attribute type
type MdtXfrmAttrType string

const (
    // Indicates that no filter has been 
    // specified.
    MdtXfrmAttrType_mdt_xfrm_attr_none MdtXfrmAttrType = "mdt-xfrm-attr-none"

    // Indicates that mandatory filter is set.
    MdtXfrmAttrType_mandatory MdtXfrmAttrType = "mandatory"

    // Indicates that primary filter is set.
    MdtXfrmAttrType_primary MdtXfrmAttrType = "primary"
)

// MdtConfigData
// MDT configuration data
type MdtConfigData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of subscriptions. The type is slice of MdtConfigData_MdtSubscription.
    MdtSubscription []*MdtConfigData_MdtSubscription

    // List of subscription transforms. The type is slice of
    // MdtConfigData_MdtXfrm.
    MdtXfrm []*MdtConfigData_MdtXfrm
}

func (mdtConfigData *MdtConfigData) GetEntityData() *types.CommonEntityData {
    mdtConfigData.EntityData.YFilter = mdtConfigData.YFilter
    mdtConfigData.EntityData.YangName = "mdt-config-data"
    mdtConfigData.EntityData.BundleName = "cisco_ios_xe"
    mdtConfigData.EntityData.ParentYangName = "Cisco-IOS-XE-mdt-cfg"
    mdtConfigData.EntityData.SegmentPath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data"
    mdtConfigData.EntityData.AbsolutePath = mdtConfigData.EntityData.SegmentPath
    mdtConfigData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtConfigData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtConfigData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtConfigData.EntityData.Children = types.NewOrderedMap()
    mdtConfigData.EntityData.Children.Append("mdt-subscription", types.YChild{"MdtSubscription", nil})
    for i := range mdtConfigData.MdtSubscription {
        mdtConfigData.EntityData.Children.Append(types.GetSegmentPath(mdtConfigData.MdtSubscription[i]), types.YChild{"MdtSubscription", mdtConfigData.MdtSubscription[i]})
    }
    mdtConfigData.EntityData.Children.Append("mdt-xfrm", types.YChild{"MdtXfrm", nil})
    for i := range mdtConfigData.MdtXfrm {
        mdtConfigData.EntityData.Children.Append(types.GetSegmentPath(mdtConfigData.MdtXfrm[i]), types.YChild{"MdtXfrm", mdtConfigData.MdtXfrm[i]})
    }
    mdtConfigData.EntityData.Leafs = types.NewOrderedMap()

    mdtConfigData.EntityData.YListKeys = []string {}

    return &(mdtConfigData.EntityData)
}

// MdtConfigData_MdtSubscription
// List of subscriptions
type MdtConfigData_MdtSubscription struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique subscription identifier. The type is
    // interface{} with range: 0..2147483647.
    SubscriptionId interface{}

    // Common subscription information.
    Base MdtConfigData_MdtSubscription_Base

    // Configuration of receivers of configured  subscriptions. The type is slice
    // of MdtConfigData_MdtSubscription_MdtReceivers.
    MdtReceivers []*MdtConfigData_MdtSubscription_MdtReceivers
}

func (mdtSubscription *MdtConfigData_MdtSubscription) GetEntityData() *types.CommonEntityData {
    mdtSubscription.EntityData.YFilter = mdtSubscription.YFilter
    mdtSubscription.EntityData.YangName = "mdt-subscription"
    mdtSubscription.EntityData.BundleName = "cisco_ios_xe"
    mdtSubscription.EntityData.ParentYangName = "mdt-config-data"
    mdtSubscription.EntityData.SegmentPath = "mdt-subscription" + types.AddKeyToken(mdtSubscription.SubscriptionId, "subscription-id")
    mdtSubscription.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/" + mdtSubscription.EntityData.SegmentPath
    mdtSubscription.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtSubscription.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtSubscription.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtSubscription.EntityData.Children = types.NewOrderedMap()
    mdtSubscription.EntityData.Children.Append("base", types.YChild{"Base", &mdtSubscription.Base})
    mdtSubscription.EntityData.Children.Append("mdt-receivers", types.YChild{"MdtReceivers", nil})
    for i := range mdtSubscription.MdtReceivers {
        mdtSubscription.EntityData.Children.Append(types.GetSegmentPath(mdtSubscription.MdtReceivers[i]), types.YChild{"MdtReceivers", mdtSubscription.MdtReceivers[i]})
    }
    mdtSubscription.EntityData.Leafs = types.NewOrderedMap()
    mdtSubscription.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", mdtSubscription.SubscriptionId})

    mdtSubscription.EntityData.YListKeys = []string {"SubscriptionId"}

    return &(mdtSubscription.EntityData)
}

// MdtConfigData_MdtSubscription_Base
// Common subscription information.
type MdtConfigData_MdtSubscription_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the event stream being subscribed to. The type is string. The
    // default value is NETCONF.
    Stream interface{}

    // Update notification encoding. The type is string. The default value is
    // encode-xml.
    Encoding interface{}

    // Network instance name for the VRF. The type is string.
    SourceVrf interface{}

    // The source address for the notifications. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Placeholder for unset value. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    NoTrigger interface{}

    // Period of update notifications in 100ths of a second. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are centiseconds.
    Period interface{}

    // If true, there is no initial update notification with the current value of
    // all the data. NOT CURRENTLY SUPPORTED. If specified, must be false. The
    // type is bool.
    NoSynchOnStart interface{}

    // Placeholder for unset value. The type is interface{} with range:
    // 0..4294967295. The default value is 0.
    NoFilter interface{}

    // XPath expression describing the set of objects wanted as part of the
    // subscription. The type is string.
    Xpath interface{}

    // TDL-URI expression describing the set of objects wanted as part of the
    // subscription. The type is string.
    TdlUri interface{}

    // Transform name is the reference to  tdl transform scheme. The type is
    // string.
    TransformName interface{}
}

func (base *MdtConfigData_MdtSubscription_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xe"
    base.EntityData.ParentYangName = "mdt-subscription"
    base.EntityData.SegmentPath = "base"
    base.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription/" + base.EntityData.SegmentPath
    base.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Leafs = types.NewOrderedMap()
    base.EntityData.Leafs.Append("stream", types.YLeaf{"Stream", base.Stream})
    base.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", base.Encoding})
    base.EntityData.Leafs.Append("source-vrf", types.YLeaf{"SourceVrf", base.SourceVrf})
    base.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", base.SourceAddress})
    base.EntityData.Leafs.Append("no-trigger", types.YLeaf{"NoTrigger", base.NoTrigger})
    base.EntityData.Leafs.Append("period", types.YLeaf{"Period", base.Period})
    base.EntityData.Leafs.Append("no-synch-on-start", types.YLeaf{"NoSynchOnStart", base.NoSynchOnStart})
    base.EntityData.Leafs.Append("no-filter", types.YLeaf{"NoFilter", base.NoFilter})
    base.EntityData.Leafs.Append("xpath", types.YLeaf{"Xpath", base.Xpath})
    base.EntityData.Leafs.Append("tdl-uri", types.YLeaf{"TdlUri", base.TdlUri})
    base.EntityData.Leafs.Append("transform-name", types.YLeaf{"TransformName", base.TransformName})

    base.EntityData.YListKeys = []string {}

    return &(base.EntityData)
}

// MdtConfigData_MdtSubscription_MdtReceivers
// Configuration of receivers of configured 
// subscriptions.
type MdtConfigData_MdtSubscription_MdtReceivers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IP address of the receiver. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    Address interface{}

    // This attribute is a key. Network port of the receiver. The type is
    // interface{} with range: 0..65535. This attribute is mandatory.
    Port interface{}

    // Receiver transport protocol. The type is string. The default value is
    // netconf.
    Protocol interface{}

    // Name of protocol profile used by receiver. The type is string.
    Profile interface{}
}

func (mdtReceivers *MdtConfigData_MdtSubscription_MdtReceivers) GetEntityData() *types.CommonEntityData {
    mdtReceivers.EntityData.YFilter = mdtReceivers.YFilter
    mdtReceivers.EntityData.YangName = "mdt-receivers"
    mdtReceivers.EntityData.BundleName = "cisco_ios_xe"
    mdtReceivers.EntityData.ParentYangName = "mdt-subscription"
    mdtReceivers.EntityData.SegmentPath = "mdt-receivers" + types.AddKeyToken(mdtReceivers.Address, "address") + types.AddKeyToken(mdtReceivers.Port, "port")
    mdtReceivers.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-subscription/" + mdtReceivers.EntityData.SegmentPath
    mdtReceivers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtReceivers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtReceivers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtReceivers.EntityData.Children = types.NewOrderedMap()
    mdtReceivers.EntityData.Leafs = types.NewOrderedMap()
    mdtReceivers.EntityData.Leafs.Append("address", types.YLeaf{"Address", mdtReceivers.Address})
    mdtReceivers.EntityData.Leafs.Append("port", types.YLeaf{"Port", mdtReceivers.Port})
    mdtReceivers.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", mdtReceivers.Protocol})
    mdtReceivers.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", mdtReceivers.Profile})

    mdtReceivers.EntityData.YListKeys = []string {"Address", "Port"}

    return &(mdtReceivers.EntityData)
}

// MdtConfigData_MdtXfrm
// List of subscription transforms
type MdtConfigData_MdtXfrm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique transform identifier. The type is string.
    Name interface{}

    // When fully-specify is set,  fully-specify field identifier is sent  in the
    // response record along with field value. The type is bool.
    FullySpecify interface{}

    // Transform input information. The type is slice of
    // MdtConfigData_MdtXfrm_MdtXfrmInput.
    MdtXfrmInput []*MdtConfigData_MdtXfrm_MdtXfrmInput

    // Transform operations information. The type is slice of
    // MdtConfigData_MdtXfrm_MdtXfrmOp.
    MdtXfrmOp []*MdtConfigData_MdtXfrm_MdtXfrmOp
}

func (mdtXfrm *MdtConfigData_MdtXfrm) GetEntityData() *types.CommonEntityData {
    mdtXfrm.EntityData.YFilter = mdtXfrm.YFilter
    mdtXfrm.EntityData.YangName = "mdt-xfrm"
    mdtXfrm.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrm.EntityData.ParentYangName = "mdt-config-data"
    mdtXfrm.EntityData.SegmentPath = "mdt-xfrm" + types.AddKeyToken(mdtXfrm.Name, "name")
    mdtXfrm.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/" + mdtXfrm.EntityData.SegmentPath
    mdtXfrm.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrm.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrm.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrm.EntityData.Children = types.NewOrderedMap()
    mdtXfrm.EntityData.Children.Append("mdt-xfrm-input", types.YChild{"MdtXfrmInput", nil})
    for i := range mdtXfrm.MdtXfrmInput {
        mdtXfrm.EntityData.Children.Append(types.GetSegmentPath(mdtXfrm.MdtXfrmInput[i]), types.YChild{"MdtXfrmInput", mdtXfrm.MdtXfrmInput[i]})
    }
    mdtXfrm.EntityData.Children.Append("mdt-xfrm-op", types.YChild{"MdtXfrmOp", nil})
    for i := range mdtXfrm.MdtXfrmOp {
        mdtXfrm.EntityData.Children.Append(types.GetSegmentPath(mdtXfrm.MdtXfrmOp[i]), types.YChild{"MdtXfrmOp", mdtXfrm.MdtXfrmOp[i]})
    }
    mdtXfrm.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrm.EntityData.Leafs.Append("name", types.YLeaf{"Name", mdtXfrm.Name})
    mdtXfrm.EntityData.Leafs.Append("fully-specify", types.YLeaf{"FullySpecify", mdtXfrm.FullySpecify})

    mdtXfrm.EntityData.YListKeys = []string {"Name"}

    return &(mdtXfrm.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmInput
// Transform input information
type MdtConfigData_MdtXfrm_MdtXfrmInput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Transform input URI table name. The type is
    // string.
    TableName interface{}

    // Transform input URI full-path. The type is string.
    Uri interface{}

    // Transform input table join-key. The type is string.
    JoinKey interface{}

    // Transform input table attribute type  e.g. table is mandatory for join
    // record. The type is MdtXfrmAttrType.
    AttrType interface{}

    // Logical operation with next input table event. The type is MdtXfrmLogicOp.
    Lop interface{}

    // Transform input URI table fields. The type is slice of
    // MdtConfigData_MdtXfrm_MdtXfrmInput_MdtXfrmInputField.
    MdtXfrmInputField []*MdtConfigData_MdtXfrm_MdtXfrmInput_MdtXfrmInputField
}

func (mdtXfrmInput *MdtConfigData_MdtXfrm_MdtXfrmInput) GetEntityData() *types.CommonEntityData {
    mdtXfrmInput.EntityData.YFilter = mdtXfrmInput.YFilter
    mdtXfrmInput.EntityData.YangName = "mdt-xfrm-input"
    mdtXfrmInput.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrmInput.EntityData.ParentYangName = "mdt-xfrm"
    mdtXfrmInput.EntityData.SegmentPath = "mdt-xfrm-input" + types.AddKeyToken(mdtXfrmInput.TableName, "table-name")
    mdtXfrmInput.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/" + mdtXfrmInput.EntityData.SegmentPath
    mdtXfrmInput.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrmInput.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrmInput.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrmInput.EntityData.Children = types.NewOrderedMap()
    mdtXfrmInput.EntityData.Children.Append("mdt-xfrm-input-field", types.YChild{"MdtXfrmInputField", nil})
    for i := range mdtXfrmInput.MdtXfrmInputField {
        mdtXfrmInput.EntityData.Children.Append(types.GetSegmentPath(mdtXfrmInput.MdtXfrmInputField[i]), types.YChild{"MdtXfrmInputField", mdtXfrmInput.MdtXfrmInputField[i]})
    }
    mdtXfrmInput.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrmInput.EntityData.Leafs.Append("table-name", types.YLeaf{"TableName", mdtXfrmInput.TableName})
    mdtXfrmInput.EntityData.Leafs.Append("uri", types.YLeaf{"Uri", mdtXfrmInput.Uri})
    mdtXfrmInput.EntityData.Leafs.Append("join-key", types.YLeaf{"JoinKey", mdtXfrmInput.JoinKey})
    mdtXfrmInput.EntityData.Leafs.Append("attr-type", types.YLeaf{"AttrType", mdtXfrmInput.AttrType})
    mdtXfrmInput.EntityData.Leafs.Append("lop", types.YLeaf{"Lop", mdtXfrmInput.Lop})

    mdtXfrmInput.EntityData.YListKeys = []string {"TableName"}

    return &(mdtXfrmInput.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmInput_MdtXfrmInputField
// Transform input URI table fields
type MdtConfigData_MdtXfrm_MdtXfrmInput_MdtXfrmInputField struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Transform input URI table field name. The type is
    // string.
    Field interface{}
}

func (mdtXfrmInputField *MdtConfigData_MdtXfrm_MdtXfrmInput_MdtXfrmInputField) GetEntityData() *types.CommonEntityData {
    mdtXfrmInputField.EntityData.YFilter = mdtXfrmInputField.YFilter
    mdtXfrmInputField.EntityData.YangName = "mdt-xfrm-input-field"
    mdtXfrmInputField.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrmInputField.EntityData.ParentYangName = "mdt-xfrm-input"
    mdtXfrmInputField.EntityData.SegmentPath = "mdt-xfrm-input-field" + types.AddKeyToken(mdtXfrmInputField.Field, "field")
    mdtXfrmInputField.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/mdt-xfrm-input/" + mdtXfrmInputField.EntityData.SegmentPath
    mdtXfrmInputField.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrmInputField.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrmInputField.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrmInputField.EntityData.Children = types.NewOrderedMap()
    mdtXfrmInputField.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrmInputField.EntityData.Leafs.Append("field", types.YLeaf{"Field", mdtXfrmInputField.Field})

    mdtXfrmInputField.EntityData.YListKeys = []string {"Field"}

    return &(mdtXfrmInputField.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmOp
// Transform operations information
type MdtConfigData_MdtXfrm_MdtXfrmOp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique transform operation id. The type is
    // interface{} with range: 0..4294967295.
    Id interface{}

    // Transform operation filters.  These are evaluated before performing
    // transform action (e.g. subrecord)  on the response record. The type is
    // slice of MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters.
    MdtXfrmOpFilters []*MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters

    // Transform operation fields.  Default operation is subrecord. It is
    // performed on each field. The type is slice of
    // MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFields.
    MdtXfrmOpFields []*MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFields
}

func (mdtXfrmOp *MdtConfigData_MdtXfrm_MdtXfrmOp) GetEntityData() *types.CommonEntityData {
    mdtXfrmOp.EntityData.YFilter = mdtXfrmOp.YFilter
    mdtXfrmOp.EntityData.YangName = "mdt-xfrm-op"
    mdtXfrmOp.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrmOp.EntityData.ParentYangName = "mdt-xfrm"
    mdtXfrmOp.EntityData.SegmentPath = "mdt-xfrm-op" + types.AddKeyToken(mdtXfrmOp.Id, "id")
    mdtXfrmOp.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/" + mdtXfrmOp.EntityData.SegmentPath
    mdtXfrmOp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrmOp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrmOp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrmOp.EntityData.Children = types.NewOrderedMap()
    mdtXfrmOp.EntityData.Children.Append("mdt-xfrm-op-filters", types.YChild{"MdtXfrmOpFilters", nil})
    for i := range mdtXfrmOp.MdtXfrmOpFilters {
        mdtXfrmOp.EntityData.Children.Append(types.GetSegmentPath(mdtXfrmOp.MdtXfrmOpFilters[i]), types.YChild{"MdtXfrmOpFilters", mdtXfrmOp.MdtXfrmOpFilters[i]})
    }
    mdtXfrmOp.EntityData.Children.Append("mdt-xfrm-op-fields", types.YChild{"MdtXfrmOpFields", nil})
    for i := range mdtXfrmOp.MdtXfrmOpFields {
        mdtXfrmOp.EntityData.Children.Append(types.GetSegmentPath(mdtXfrmOp.MdtXfrmOpFields[i]), types.YChild{"MdtXfrmOpFields", mdtXfrmOp.MdtXfrmOpFields[i]})
    }
    mdtXfrmOp.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrmOp.EntityData.Leafs.Append("id", types.YLeaf{"Id", mdtXfrmOp.Id})

    mdtXfrmOp.EntityData.YListKeys = []string {"Id"}

    return &(mdtXfrmOp.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters
// Transform operation filters. 
// These are evaluated before performing
// transform action (e.g. subrecord) 
// on the response record
type MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. filters will be evaluated in sequence based  on
    // filter_id. The type is interface{} with range: 0..4294967295.
    FilterId interface{}

    // Transform operation filter field name. The type is string.
    Field interface{}

    // logical operation with condition. The type is MdtXfrmLogicOp.
    Lop interface{}

    // logical operation with next filter condition. The type is MdtXfrmLogicOp.
    NextLop interface{}

    // Transform operation event flag (e.g. onchange).
    OpEvent MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_OpEvent

    // Per field condition (e.g. f1 eq 'name').
    Condition MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_Condition
}

func (mdtXfrmOpFilters *MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters) GetEntityData() *types.CommonEntityData {
    mdtXfrmOpFilters.EntityData.YFilter = mdtXfrmOpFilters.YFilter
    mdtXfrmOpFilters.EntityData.YangName = "mdt-xfrm-op-filters"
    mdtXfrmOpFilters.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrmOpFilters.EntityData.ParentYangName = "mdt-xfrm-op"
    mdtXfrmOpFilters.EntityData.SegmentPath = "mdt-xfrm-op-filters" + types.AddKeyToken(mdtXfrmOpFilters.FilterId, "filter-id")
    mdtXfrmOpFilters.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/mdt-xfrm-op/" + mdtXfrmOpFilters.EntityData.SegmentPath
    mdtXfrmOpFilters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrmOpFilters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrmOpFilters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrmOpFilters.EntityData.Children = types.NewOrderedMap()
    mdtXfrmOpFilters.EntityData.Children.Append("op-event", types.YChild{"OpEvent", &mdtXfrmOpFilters.OpEvent})
    mdtXfrmOpFilters.EntityData.Children.Append("condition", types.YChild{"Condition", &mdtXfrmOpFilters.Condition})
    mdtXfrmOpFilters.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrmOpFilters.EntityData.Leafs.Append("filter-id", types.YLeaf{"FilterId", mdtXfrmOpFilters.FilterId})
    mdtXfrmOpFilters.EntityData.Leafs.Append("field", types.YLeaf{"Field", mdtXfrmOpFilters.Field})
    mdtXfrmOpFilters.EntityData.Leafs.Append("lop", types.YLeaf{"Lop", mdtXfrmOpFilters.Lop})
    mdtXfrmOpFilters.EntityData.Leafs.Append("next-lop", types.YLeaf{"NextLop", mdtXfrmOpFilters.NextLop})

    mdtXfrmOpFilters.EntityData.YListKeys = []string {"FilterId"}

    return &(mdtXfrmOpFilters.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_OpEvent
// Transform operation event flag (e.g. onchange)
type MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_OpEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates that onchange filter is set. The type is bool.
    Onchange interface{}
}

func (opEvent *MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_OpEvent) GetEntityData() *types.CommonEntityData {
    opEvent.EntityData.YFilter = opEvent.YFilter
    opEvent.EntityData.YangName = "op-event"
    opEvent.EntityData.BundleName = "cisco_ios_xe"
    opEvent.EntityData.ParentYangName = "mdt-xfrm-op-filters"
    opEvent.EntityData.SegmentPath = "op-event"
    opEvent.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/mdt-xfrm-op/mdt-xfrm-op-filters/" + opEvent.EntityData.SegmentPath
    opEvent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    opEvent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    opEvent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    opEvent.EntityData.Children = types.NewOrderedMap()
    opEvent.EntityData.Leafs = types.NewOrderedMap()
    opEvent.EntityData.Leafs.Append("onchange", types.YLeaf{"Onchange", opEvent.Onchange})

    opEvent.EntityData.YListKeys = []string {}

    return &(opEvent.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_Condition
// Per field condition (e.g. f1 eq 'name')
type MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_Condition struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of operator. The type is MdtXfrmOperator.
    Operator interface{}

    // Field value to operate on. The type is string.
    Value interface{}
}

func (condition *MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFilters_Condition) GetEntityData() *types.CommonEntityData {
    condition.EntityData.YFilter = condition.YFilter
    condition.EntityData.YangName = "condition"
    condition.EntityData.BundleName = "cisco_ios_xe"
    condition.EntityData.ParentYangName = "mdt-xfrm-op-filters"
    condition.EntityData.SegmentPath = "condition"
    condition.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/mdt-xfrm-op/mdt-xfrm-op-filters/" + condition.EntityData.SegmentPath
    condition.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    condition.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    condition.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    condition.EntityData.Children = types.NewOrderedMap()
    condition.EntityData.Leafs = types.NewOrderedMap()
    condition.EntityData.Leafs.Append("operator", types.YLeaf{"Operator", condition.Operator})
    condition.EntityData.Leafs.Append("value", types.YLeaf{"Value", condition.Value})

    condition.EntityData.YListKeys = []string {}

    return &(condition.EntityData)
}

// MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFields
// Transform operation fields. 
// Default operation is subrecord.
// It is performed on each field
type MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFields struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Transform operation response field-id, part of
    // response record. It is used to uniquely identify response record field. The
    // type is interface{} with range: 0..4294967295.
    FieldId interface{}

    // Subscription transform field name on  which transform operation is
    // performed. The type is string.
    Field interface{}

    // Subscription transform operation type. The type is MdtXfrmOpType.
    OpType interface{}
}

func (mdtXfrmOpFields *MdtConfigData_MdtXfrm_MdtXfrmOp_MdtXfrmOpFields) GetEntityData() *types.CommonEntityData {
    mdtXfrmOpFields.EntityData.YFilter = mdtXfrmOpFields.YFilter
    mdtXfrmOpFields.EntityData.YangName = "mdt-xfrm-op-fields"
    mdtXfrmOpFields.EntityData.BundleName = "cisco_ios_xe"
    mdtXfrmOpFields.EntityData.ParentYangName = "mdt-xfrm-op"
    mdtXfrmOpFields.EntityData.SegmentPath = "mdt-xfrm-op-fields" + types.AddKeyToken(mdtXfrmOpFields.FieldId, "field-id")
    mdtXfrmOpFields.EntityData.AbsolutePath = "Cisco-IOS-XE-mdt-cfg:mdt-config-data/mdt-xfrm/mdt-xfrm-op/" + mdtXfrmOpFields.EntityData.SegmentPath
    mdtXfrmOpFields.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mdtXfrmOpFields.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mdtXfrmOpFields.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mdtXfrmOpFields.EntityData.Children = types.NewOrderedMap()
    mdtXfrmOpFields.EntityData.Leafs = types.NewOrderedMap()
    mdtXfrmOpFields.EntityData.Leafs.Append("field-id", types.YLeaf{"FieldId", mdtXfrmOpFields.FieldId})
    mdtXfrmOpFields.EntityData.Leafs.Append("field", types.YLeaf{"Field", mdtXfrmOpFields.Field})
    mdtXfrmOpFields.EntityData.Leafs.Append("op-type", types.YLeaf{"OpType", mdtXfrmOpFields.OpType})

    mdtXfrmOpFields.EntityData.YListKeys = []string {"FieldId"}

    return &(mdtXfrmOpFields.EntityData)
}

