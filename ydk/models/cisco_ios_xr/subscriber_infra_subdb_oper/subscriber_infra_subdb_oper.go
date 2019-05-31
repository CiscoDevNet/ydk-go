// This module contains a collection of YANG definitions
// for Cisco IOS-XR Subscriber-infra-subdb package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber-database: Subscriber database operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package subscriber_infra_subdb_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package subscriber_infra_subdb_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-Subscriber-infra-subdb-oper subscriber-database}", reflect.TypeOf(SubscriberDatabase{}))
    ydk.RegisterEntity("Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database", reflect.TypeOf(SubscriberDatabase{}))
}

// SessionState represents Session states
type SessionState string

const (
    // Initialisation
    SessionState_init SessionState = "init"

    // Destroy
    SessionState_destroy SessionState = "destroy"

    // Config generate
    SessionState_config_generate SessionState = "config-generate"

    // Feature registration wait
    SessionState_feature_registration_wait SessionState = "feature-registration-wait"

    // Config apply
    SessionState_config_apply SessionState = "config-apply"

    // Config done
    SessionState_config_done SessionState = "config-done"

    // Config removed
    SessionState_config_removed SessionState = "config-removed"

    // Config error
    SessionState_config_error SessionState = "config-error"

    // Error
    SessionState_error_ SessionState = "error"

    // Sync
    SessionState_sync SessionState = "sync"
)

// SubdbObjectTypeData represents Template types
type SubdbObjectTypeData string

const (
    // User profile
    SubdbObjectTypeData_user_profile SubdbObjectTypeData = "user-profile"

    // Service profile
    SubdbObjectTypeData_service_profile SubdbObjectTypeData = "service-profile"

    // Subscriber service template
    SubdbObjectTypeData_subscriber_service SubdbObjectTypeData = "subscriber-service"

    // PPP template
    SubdbObjectTypeData_ppp SubdbObjectTypeData = "ppp"

    // IP subscriber template
    SubdbObjectTypeData_ip_subscriber SubdbObjectTypeData = "ip-subscriber"
)

// SubscriberDatabase
// Subscriber database operational data
type SubscriberDatabase struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of nodes for which subscriber data is collected.
    Nodes SubscriberDatabase_Nodes
}

func (subscriberDatabase *SubscriberDatabase) GetEntityData() *types.CommonEntityData {
    subscriberDatabase.EntityData.YFilter = subscriberDatabase.YFilter
    subscriberDatabase.EntityData.YangName = "subscriber-database"
    subscriberDatabase.EntityData.BundleName = "cisco_ios_xr"
    subscriberDatabase.EntityData.ParentYangName = "Cisco-IOS-XR-Subscriber-infra-subdb-oper"
    subscriberDatabase.EntityData.SegmentPath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database"
    subscriberDatabase.EntityData.AbsolutePath = subscriberDatabase.EntityData.SegmentPath
    subscriberDatabase.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberDatabase.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberDatabase.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberDatabase.EntityData.Children = types.NewOrderedMap()
    subscriberDatabase.EntityData.Children.Append("nodes", types.YChild{"Nodes", &subscriberDatabase.Nodes})
    subscriberDatabase.EntityData.Leafs = types.NewOrderedMap()

    subscriberDatabase.EntityData.YListKeys = []string {}

    return &(subscriberDatabase.EntityData)
}

// SubscriberDatabase_Nodes
// List of nodes for which subscriber data is
// collected
type SubscriberDatabase_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber data for a particular node. The type is slice of
    // SubscriberDatabase_Nodes_Node.
    Node []*SubscriberDatabase_Nodes_Node
}

func (nodes *SubscriberDatabase_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "subscriber-database"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// SubscriberDatabase_Nodes_Node
// Subscriber data for a particular node
type SubscriberDatabase_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Subscriber data for associated templates.
    SubdbAssoc SubscriberDatabase_Nodes_Node_SubdbAssoc

    // Subscriber data for associated templates.
    Association SubscriberDatabase_Nodes_Node_Association

    // Subscriber data for associated templates.
    Summary SubscriberDatabase_Nodes_Node_Summary

    // Subscriber management session information.
    Session SubscriberDatabase_Nodes_Node_Session
}

func (node *SubscriberDatabase_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("subdb-assoc", types.YChild{"SubdbAssoc", &node.SubdbAssoc})
    node.EntityData.Children.Append("association", types.YChild{"Association", &node.Association})
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Children.Append("session", types.YChild{"Session", &node.Session})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// SubscriberDatabase_Nodes_Node_SubdbAssoc
// Subscriber data for associated templates
type SubscriberDatabase_Nodes_Node_SubdbAssoc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of associated subscriber labels.
    Labels SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels
}

func (subdbAssoc *SubscriberDatabase_Nodes_Node_SubdbAssoc) GetEntityData() *types.CommonEntityData {
    subdbAssoc.EntityData.YFilter = subdbAssoc.YFilter
    subdbAssoc.EntityData.YangName = "subdb-assoc"
    subdbAssoc.EntityData.BundleName = "cisco_ios_xr"
    subdbAssoc.EntityData.ParentYangName = "node"
    subdbAssoc.EntityData.SegmentPath = "subdb-assoc"
    subdbAssoc.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/" + subdbAssoc.EntityData.SegmentPath
    subdbAssoc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subdbAssoc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subdbAssoc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subdbAssoc.EntityData.Children = types.NewOrderedMap()
    subdbAssoc.EntityData.Children.Append("labels", types.YChild{"Labels", &subdbAssoc.Labels})
    subdbAssoc.EntityData.Leafs = types.NewOrderedMap()

    subdbAssoc.EntityData.YListKeys = []string {}

    return &(subdbAssoc.EntityData)
}

// SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels
// List of associated subscriber labels
type SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association for a given subscriber label. The type is slice of
    // SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label.
    Label []*SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label
}

func (labels *SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels) GetEntityData() *types.CommonEntityData {
    labels.EntityData.YFilter = labels.YFilter
    labels.EntityData.YangName = "labels"
    labels.EntityData.BundleName = "cisco_ios_xr"
    labels.EntityData.ParentYangName = "subdb-assoc"
    labels.EntityData.SegmentPath = "labels"
    labels.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/subdb-assoc/" + labels.EntityData.SegmentPath
    labels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labels.EntityData.Children = types.NewOrderedMap()
    labels.EntityData.Children.Append("label", types.YChild{"Label", nil})
    for i := range labels.Label {
        labels.EntityData.Children.Append(types.GetSegmentPath(labels.Label[i]), types.YChild{"Label", labels.Label[i]})
    }
    labels.EntityData.Leafs = types.NewOrderedMap()

    labels.EntityData.YListKeys = []string {}

    return &(labels.EntityData)
}

// SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label
// Association for a given subscriber label
type SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber label. The type is interface{} with
    // range: 0..4294967295.
    SubscriberLabel interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Association count which reflects number of entries in AssociatedTemplates.
    // The type is interface{} with range: 0..4294967295.
    Associations interface{}

    // Varlist Id. The type is interface{} with range: 0..4294967295.
    VarlistId interface{}

    // Subdb template.
    Template SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template
}

func (label *SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "labels"
    label.EntityData.SegmentPath = "label" + types.AddKeyToken(label.SubscriberLabel, "subscriber-label")
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/subdb-assoc/labels/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Children.Append("template", types.YChild{"Template", &label.Template})
    label.EntityData.Leafs = types.NewOrderedMap()
    label.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", label.SubscriberLabel})
    label.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", label.SessionId})
    label.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", label.InterfaceName})
    label.EntityData.Leafs.Append("associations", types.YLeaf{"Associations", label.Associations})
    label.EntityData.Leafs.Append("varlist-id", types.YLeaf{"VarlistId", label.VarlistId})

    label.EntityData.YListKeys = []string {"SubscriberLabel"}

    return &(label.EntityData)
}

// SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template
// Subdb template
type SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Associated templates. The type is slice of
    // SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template_AssociatedTemplate.
    AssociatedTemplate []*SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template_AssociatedTemplate
}

func (template *SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "label"
    template.EntityData.SegmentPath = "template"
    template.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/subdb-assoc/labels/label/" + template.EntityData.SegmentPath
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Children.Append("associated-template", types.YChild{"AssociatedTemplate", nil})
    for i := range template.AssociatedTemplate {
        types.SetYListKey(template.AssociatedTemplate[i], i)
        template.EntityData.Children.Append(types.GetSegmentPath(template.AssociatedTemplate[i]), types.YChild{"AssociatedTemplate", template.AssociatedTemplate[i]})
    }
    template.EntityData.Leafs = types.NewOrderedMap()

    template.EntityData.YListKeys = []string {}

    return &(template.EntityData)
}

// SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template_AssociatedTemplate
// Associated templates
type SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template_AssociatedTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Template type. The type is SubdbObjectTypeData.
    TemplateType interface{}

    // Template name. The type is string with length: 0..65.
    TemplateName interface{}

    // Varlist. The type is string with length: 0..1000.
    Varlist interface{}
}

func (associatedTemplate *SubscriberDatabase_Nodes_Node_SubdbAssoc_Labels_Label_Template_AssociatedTemplate) GetEntityData() *types.CommonEntityData {
    associatedTemplate.EntityData.YFilter = associatedTemplate.YFilter
    associatedTemplate.EntityData.YangName = "associated-template"
    associatedTemplate.EntityData.BundleName = "cisco_ios_xr"
    associatedTemplate.EntityData.ParentYangName = "template"
    associatedTemplate.EntityData.SegmentPath = "associated-template" + types.AddNoKeyToken(associatedTemplate)
    associatedTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/subdb-assoc/labels/label/template/" + associatedTemplate.EntityData.SegmentPath
    associatedTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associatedTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associatedTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associatedTemplate.EntityData.Children = types.NewOrderedMap()
    associatedTemplate.EntityData.Leafs = types.NewOrderedMap()
    associatedTemplate.EntityData.Leafs.Append("template-type", types.YLeaf{"TemplateType", associatedTemplate.TemplateType})
    associatedTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", associatedTemplate.TemplateName})
    associatedTemplate.EntityData.Leafs.Append("varlist", types.YLeaf{"Varlist", associatedTemplate.Varlist})

    associatedTemplate.EntityData.YListKeys = []string {}

    return &(associatedTemplate.EntityData)
}

// SubscriberDatabase_Nodes_Node_Association
// Subscriber data for associated templates
type SubscriberDatabase_Nodes_Node_Association struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of associated subscriber labels.
    Labels SubscriberDatabase_Nodes_Node_Association_Labels
}

func (association *SubscriberDatabase_Nodes_Node_Association) GetEntityData() *types.CommonEntityData {
    association.EntityData.YFilter = association.YFilter
    association.EntityData.YangName = "association"
    association.EntityData.BundleName = "cisco_ios_xr"
    association.EntityData.ParentYangName = "node"
    association.EntityData.SegmentPath = "association"
    association.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/" + association.EntityData.SegmentPath
    association.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    association.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    association.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    association.EntityData.Children = types.NewOrderedMap()
    association.EntityData.Children.Append("labels", types.YChild{"Labels", &association.Labels})
    association.EntityData.Leafs = types.NewOrderedMap()

    association.EntityData.YListKeys = []string {}

    return &(association.EntityData)
}

// SubscriberDatabase_Nodes_Node_Association_Labels
// List of associated subscriber labels
type SubscriberDatabase_Nodes_Node_Association_Labels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association for a given subscriber label. The type is slice of
    // SubscriberDatabase_Nodes_Node_Association_Labels_Label.
    Label []*SubscriberDatabase_Nodes_Node_Association_Labels_Label
}

func (labels *SubscriberDatabase_Nodes_Node_Association_Labels) GetEntityData() *types.CommonEntityData {
    labels.EntityData.YFilter = labels.YFilter
    labels.EntityData.YangName = "labels"
    labels.EntityData.BundleName = "cisco_ios_xr"
    labels.EntityData.ParentYangName = "association"
    labels.EntityData.SegmentPath = "labels"
    labels.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/association/" + labels.EntityData.SegmentPath
    labels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labels.EntityData.Children = types.NewOrderedMap()
    labels.EntityData.Children.Append("label", types.YChild{"Label", nil})
    for i := range labels.Label {
        labels.EntityData.Children.Append(types.GetSegmentPath(labels.Label[i]), types.YChild{"Label", labels.Label[i]})
    }
    labels.EntityData.Leafs = types.NewOrderedMap()

    labels.EntityData.YListKeys = []string {}

    return &(labels.EntityData)
}

// SubscriberDatabase_Nodes_Node_Association_Labels_Label
// Association for a given subscriber label
type SubscriberDatabase_Nodes_Node_Association_Labels_Label struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber label. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    SubscriberLabel interface{}

    // Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Association count which reflects number of entries in AssociatedTemplates.
    // The type is interface{} with range: 0..4294967295.
    Associations interface{}

    // Varlist Id. The type is interface{} with range: 0..4294967295.
    VarlistId interface{}

    // Subdb template.
    Template SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template
}

func (label *SubscriberDatabase_Nodes_Node_Association_Labels_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "labels"
    label.EntityData.SegmentPath = "label" + types.AddKeyToken(label.SubscriberLabel, "subscriber-label")
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/association/labels/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Children.Append("template", types.YChild{"Template", &label.Template})
    label.EntityData.Leafs = types.NewOrderedMap()
    label.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", label.SubscriberLabel})
    label.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", label.SessionId})
    label.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", label.InterfaceName})
    label.EntityData.Leafs.Append("associations", types.YLeaf{"Associations", label.Associations})
    label.EntityData.Leafs.Append("varlist-id", types.YLeaf{"VarlistId", label.VarlistId})

    label.EntityData.YListKeys = []string {"SubscriberLabel"}

    return &(label.EntityData)
}

// SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template
// Subdb template
type SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Associated templates. The type is slice of
    // SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template_AssociatedTemplate.
    AssociatedTemplate []*SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template_AssociatedTemplate
}

func (template *SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template) GetEntityData() *types.CommonEntityData {
    template.EntityData.YFilter = template.YFilter
    template.EntityData.YangName = "template"
    template.EntityData.BundleName = "cisco_ios_xr"
    template.EntityData.ParentYangName = "label"
    template.EntityData.SegmentPath = "template"
    template.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/association/labels/label/" + template.EntityData.SegmentPath
    template.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    template.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    template.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    template.EntityData.Children = types.NewOrderedMap()
    template.EntityData.Children.Append("associated-template", types.YChild{"AssociatedTemplate", nil})
    for i := range template.AssociatedTemplate {
        types.SetYListKey(template.AssociatedTemplate[i], i)
        template.EntityData.Children.Append(types.GetSegmentPath(template.AssociatedTemplate[i]), types.YChild{"AssociatedTemplate", template.AssociatedTemplate[i]})
    }
    template.EntityData.Leafs = types.NewOrderedMap()

    template.EntityData.YListKeys = []string {}

    return &(template.EntityData)
}

// SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template_AssociatedTemplate
// Associated templates
type SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template_AssociatedTemplate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Template type. The type is SubdbObjectTypeData.
    TemplateType interface{}

    // Template name. The type is string with length: 0..65.
    TemplateName interface{}

    // Varlist. The type is string with length: 0..1000.
    Varlist interface{}
}

func (associatedTemplate *SubscriberDatabase_Nodes_Node_Association_Labels_Label_Template_AssociatedTemplate) GetEntityData() *types.CommonEntityData {
    associatedTemplate.EntityData.YFilter = associatedTemplate.YFilter
    associatedTemplate.EntityData.YangName = "associated-template"
    associatedTemplate.EntityData.BundleName = "cisco_ios_xr"
    associatedTemplate.EntityData.ParentYangName = "template"
    associatedTemplate.EntityData.SegmentPath = "associated-template" + types.AddNoKeyToken(associatedTemplate)
    associatedTemplate.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/association/labels/label/template/" + associatedTemplate.EntityData.SegmentPath
    associatedTemplate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associatedTemplate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associatedTemplate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associatedTemplate.EntityData.Children = types.NewOrderedMap()
    associatedTemplate.EntityData.Leafs = types.NewOrderedMap()
    associatedTemplate.EntityData.Leafs.Append("template-type", types.YLeaf{"TemplateType", associatedTemplate.TemplateType})
    associatedTemplate.EntityData.Leafs.Append("template-name", types.YLeaf{"TemplateName", associatedTemplate.TemplateName})
    associatedTemplate.EntityData.Leafs.Append("varlist", types.YLeaf{"Varlist", associatedTemplate.Varlist})

    associatedTemplate.EntityData.YListKeys = []string {}

    return &(associatedTemplate.EntityData)
}

// SubscriberDatabase_Nodes_Node_Summary
// Subscriber data for associated templates
type SubscriberDatabase_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Entries in Association DB. The type is interface{} with range:
    // 0..4294967295.
    AssocDbEntries interface{}

    // Number of Associations in Association DB. The type is interface{} with
    // range: 0..4294967295.
    AssocDbAssociations interface{}

    // Number of Entries in Derived DB. The type is interface{} with range:
    // 0..4294967295.
    DerivedDbEntries interface{}

    // Number of Entries in Configuration DB. The type is interface{} with range:
    // 0..4294967295.
    ConfigDbEntries interface{}

    // Number of Entries in Interface DB. The type is interface{} with range:
    // 0..4294967295.
    InterfaceDbEntries interface{}

    // Number of IPSUB DHCP subscribers. The type is interface{} with range:
    // 0..4294967295.
    NumIpsubDhcp interface{}

    // Number of IPSUB Inband subscribers. The type is interface{} with range:
    // 0..4294967295.
    NumIpsubInband interface{}

    // Number of PPPOE subscribers. The type is interface{} with range:
    // 0..4294967295.
    NumPppoe interface{}

    // The count of the various configuration objects by type. The type is slice
    // of interface{} with range: 0..4294967295.
    SubdbObjCountsByType []interface{}

    // Number of subscribers in the various states. The type is slice of
    // interface{} with range: 0..4294967295.
    NumSubscribersInState []interface{}

    // Cumulative number of transitions through the various states. The type is
    // slice of interface{} with range: 0..4294967295.
    NumTransitionsThroughState []interface{}
}

func (summary *SubscriberDatabase_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("assoc-db-entries", types.YLeaf{"AssocDbEntries", summary.AssocDbEntries})
    summary.EntityData.Leafs.Append("assoc-db-associations", types.YLeaf{"AssocDbAssociations", summary.AssocDbAssociations})
    summary.EntityData.Leafs.Append("derived-db-entries", types.YLeaf{"DerivedDbEntries", summary.DerivedDbEntries})
    summary.EntityData.Leafs.Append("config-db-entries", types.YLeaf{"ConfigDbEntries", summary.ConfigDbEntries})
    summary.EntityData.Leafs.Append("interface-db-entries", types.YLeaf{"InterfaceDbEntries", summary.InterfaceDbEntries})
    summary.EntityData.Leafs.Append("num-ipsub-dhcp", types.YLeaf{"NumIpsubDhcp", summary.NumIpsubDhcp})
    summary.EntityData.Leafs.Append("num-ipsub-inband", types.YLeaf{"NumIpsubInband", summary.NumIpsubInband})
    summary.EntityData.Leafs.Append("num-pppoe", types.YLeaf{"NumPppoe", summary.NumPppoe})
    summary.EntityData.Leafs.Append("subdb-obj-counts-by-type", types.YLeaf{"SubdbObjCountsByType", summary.SubdbObjCountsByType})
    summary.EntityData.Leafs.Append("num-subscribers-in-state", types.YLeaf{"NumSubscribersInState", summary.NumSubscribersInState})
    summary.EntityData.Leafs.Append("num-transitions-through-state", types.YLeaf{"NumTransitionsThroughState", summary.NumTransitionsThroughState})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// SubscriberDatabase_Nodes_Node_Session
// Subscriber management session information
type SubscriberDatabase_Nodes_Node_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber management list of subscriber labels.
    Labels SubscriberDatabase_Nodes_Node_Session_Labels
}

func (session *SubscriberDatabase_Nodes_Node_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "node"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("labels", types.YChild{"Labels", &session.Labels})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// SubscriberDatabase_Nodes_Node_Session_Labels
// Subscriber management list of subscriber
// labels
type SubscriberDatabase_Nodes_Node_Session_Labels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session information for a subscriber label. The type is slice of
    // SubscriberDatabase_Nodes_Node_Session_Labels_Label.
    Label []*SubscriberDatabase_Nodes_Node_Session_Labels_Label
}

func (labels *SubscriberDatabase_Nodes_Node_Session_Labels) GetEntityData() *types.CommonEntityData {
    labels.EntityData.YFilter = labels.YFilter
    labels.EntityData.YangName = "labels"
    labels.EntityData.BundleName = "cisco_ios_xr"
    labels.EntityData.ParentYangName = "session"
    labels.EntityData.SegmentPath = "labels"
    labels.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/session/" + labels.EntityData.SegmentPath
    labels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labels.EntityData.Children = types.NewOrderedMap()
    labels.EntityData.Children.Append("label", types.YChild{"Label", nil})
    for i := range labels.Label {
        labels.EntityData.Children.Append(types.GetSegmentPath(labels.Label[i]), types.YChild{"Label", labels.Label[i]})
    }
    labels.EntityData.Leafs = types.NewOrderedMap()

    labels.EntityData.YListKeys = []string {}

    return &(labels.EntityData)
}

// SubscriberDatabase_Nodes_Node_Session_Labels_Label
// Session information for a subscriber label
type SubscriberDatabase_Nodes_Node_Session_Labels_Label struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Subscriber label. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    SubscriberLabel interface{}

    // Subscriber session state. The type is SessionState.
    SessionState interface{}

    // Activate request identifier. The type is interface{} with range:
    // 0..4294967295.
    ActivateRequestId interface{}

    // Transaction identifier associated with a particular 'produce_done' or
    // 'produce_all_done' request  default value is 0xffffffff which represents
    // 'None'. The type is interface{} with range: -2147483648..2147483647.
    TransactionId interface{}

    // Produce done request ID. The type is interface{} with range: 0..4294967295.
    ProduceDoneRequestId interface{}

    // Flags indicating if a destroy request is received. The type is bool.
    DestroyReqReceived interface{}

    // Destroy request ID. The type is interface{} with range: 0..4294967295.
    DestroyRequestId interface{}

    // Is true if configuration change due to template change only and not due to
    // a produce done request. The type is bool.
    IsConfigChanged interface{}

    // Is true if the creator of the connection is destroyed. The type is bool.
    IsCreatorGone interface{}

    // Is true if the deleted features have all been notified and all 'apply done'
    // ack messages have been received. The type is bool.
    IsDeleteNotifyDone interface{}

    // Is true if added/modified features have all been notified and all 'apply
    // done' ack messages have been received. The type is bool.
    AddModifyDone interface{}

    // Is true if the subscriber should be rolled back to the configuration prior
    // to this transaction when all the outstanding  backend programming interface
    // requests are replied. The type is bool.
    IsRollbackNeeded interface{}

    // Is true if subscriber's configuration is being rolled back. The type is
    // bool.
    IsRollbackInProgress interface{}

    // Is true if the subscriber's configuration is being applied due to
    // subscriber database server restart. The type is bool.
    IsServerRestartApply interface{}

    // Is true if rollback has previously been performed for this subscriber. The
    // type is bool.
    IsRollbackPerformed interface{}

    // Flags indicating if there is pending replication. The type is bool.
    ReplPending interface{}

    // Flags indicating if activate timer is running. The type is bool.
    ActivateTimerRunning interface{}

    // Flags indicating if apply timer is running. The type is bool.
    ApplyTimerRunning interface{}

    // the current size of the event queue. The type is bool.
    EventQueueSize interface{}

    // Restart vector to keep track of the restart state. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Restarts interface{}

    // Template Interface Identifier. The type is interface{} with range:
    // 0..4294967295.
    TemplateInterfaceId interface{}
}

func (label *SubscriberDatabase_Nodes_Node_Session_Labels_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "labels"
    label.EntityData.SegmentPath = "label" + types.AddKeyToken(label.SubscriberLabel, "subscriber-label")
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-Subscriber-infra-subdb-oper:subscriber-database/nodes/node/session/labels/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Leafs = types.NewOrderedMap()
    label.EntityData.Leafs.Append("subscriber-label", types.YLeaf{"SubscriberLabel", label.SubscriberLabel})
    label.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", label.SessionState})
    label.EntityData.Leafs.Append("activate-request-id", types.YLeaf{"ActivateRequestId", label.ActivateRequestId})
    label.EntityData.Leafs.Append("transaction-id", types.YLeaf{"TransactionId", label.TransactionId})
    label.EntityData.Leafs.Append("produce-done-request-id", types.YLeaf{"ProduceDoneRequestId", label.ProduceDoneRequestId})
    label.EntityData.Leafs.Append("destroy-req-received", types.YLeaf{"DestroyReqReceived", label.DestroyReqReceived})
    label.EntityData.Leafs.Append("destroy-request-id", types.YLeaf{"DestroyRequestId", label.DestroyRequestId})
    label.EntityData.Leafs.Append("is-config-changed", types.YLeaf{"IsConfigChanged", label.IsConfigChanged})
    label.EntityData.Leafs.Append("is-creator-gone", types.YLeaf{"IsCreatorGone", label.IsCreatorGone})
    label.EntityData.Leafs.Append("is-delete-notify-done", types.YLeaf{"IsDeleteNotifyDone", label.IsDeleteNotifyDone})
    label.EntityData.Leafs.Append("add-modify-done", types.YLeaf{"AddModifyDone", label.AddModifyDone})
    label.EntityData.Leafs.Append("is-rollback-needed", types.YLeaf{"IsRollbackNeeded", label.IsRollbackNeeded})
    label.EntityData.Leafs.Append("is-rollback-in-progress", types.YLeaf{"IsRollbackInProgress", label.IsRollbackInProgress})
    label.EntityData.Leafs.Append("is-server-restart-apply", types.YLeaf{"IsServerRestartApply", label.IsServerRestartApply})
    label.EntityData.Leafs.Append("is-rollback-performed", types.YLeaf{"IsRollbackPerformed", label.IsRollbackPerformed})
    label.EntityData.Leafs.Append("repl-pending", types.YLeaf{"ReplPending", label.ReplPending})
    label.EntityData.Leafs.Append("activate-timer-running", types.YLeaf{"ActivateTimerRunning", label.ActivateTimerRunning})
    label.EntityData.Leafs.Append("apply-timer-running", types.YLeaf{"ApplyTimerRunning", label.ApplyTimerRunning})
    label.EntityData.Leafs.Append("event-queue-size", types.YLeaf{"EventQueueSize", label.EventQueueSize})
    label.EntityData.Leafs.Append("restarts", types.YLeaf{"Restarts", label.Restarts})
    label.EntityData.Leafs.Append("template-interface-id", types.YLeaf{"TemplateInterfaceId", label.TemplateInterfaceId})

    label.EntityData.YListKeys = []string {"SubscriberLabel"}

    return &(label.EntityData)
}

