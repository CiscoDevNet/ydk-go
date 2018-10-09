// This module contains a collection of YANG definitions
// for Cisco IOS-XR cofo-infra package operational data.
// 
// This module contains definitions
// for the following management objects:
//   cofo: COFO operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package cofo_infra_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cofo_infra_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-cofo-infra-oper cofo}", reflect.TypeOf(Cofo{}))
    ydk.RegisterEntity("Cisco-IOS-XR-cofo-infra-oper:cofo", reflect.TypeOf(Cofo{}))
}

// Cofo
// COFO operational data
type Cofo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific COFO operational data.
    Nodes Cofo_Nodes
}

func (cofo *Cofo) GetEntityData() *types.CommonEntityData {
    cofo.EntityData.YFilter = cofo.YFilter
    cofo.EntityData.YangName = "cofo"
    cofo.EntityData.BundleName = "cisco_ios_xr"
    cofo.EntityData.ParentYangName = "Cisco-IOS-XR-cofo-infra-oper"
    cofo.EntityData.SegmentPath = "Cisco-IOS-XR-cofo-infra-oper:cofo"
    cofo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cofo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cofo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cofo.EntityData.Children = types.NewOrderedMap()
    cofo.EntityData.Children.Append("nodes", types.YChild{"Nodes", &cofo.Nodes})
    cofo.EntityData.Leafs = types.NewOrderedMap()

    cofo.EntityData.YListKeys = []string {}

    return &(cofo.EntityData)
}

// Cofo_Nodes
// Node-specific COFO operational data
type Cofo_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // COFO operational data for a particular node. The type is slice of
    // Cofo_Nodes_Node.
    Node []*Cofo_Nodes_Node
}

func (nodes *Cofo_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "cofo"
    nodes.EntityData.SegmentPath = "nodes"
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

// Cofo_Nodes_Node
// COFO operational data for a particular node
type Cofo_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // COFO Client data.
    ClientIds Cofo_Nodes_Node_ClientIds

    // COFO SDR database.
    TopicIds Cofo_Nodes_Node_TopicIds
}

func (node *Cofo_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("client-ids", types.YChild{"ClientIds", &node.ClientIds})
    node.EntityData.Children.Append("topic-ids", types.YChild{"TopicIds", &node.TopicIds})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Cofo_Nodes_Node_ClientIds
// COFO Client data
type Cofo_Nodes_Node_ClientIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // COFO client id. The type is slice of Cofo_Nodes_Node_ClientIds_ClientId.
    ClientId []*Cofo_Nodes_Node_ClientIds_ClientId
}

func (clientIds *Cofo_Nodes_Node_ClientIds) GetEntityData() *types.CommonEntityData {
    clientIds.EntityData.YFilter = clientIds.YFilter
    clientIds.EntityData.YangName = "client-ids"
    clientIds.EntityData.BundleName = "cisco_ios_xr"
    clientIds.EntityData.ParentYangName = "node"
    clientIds.EntityData.SegmentPath = "client-ids"
    clientIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientIds.EntityData.Children = types.NewOrderedMap()
    clientIds.EntityData.Children.Append("client-id", types.YChild{"ClientId", nil})
    for i := range clientIds.ClientId {
        clientIds.EntityData.Children.Append(types.GetSegmentPath(clientIds.ClientId[i]), types.YChild{"ClientId", clientIds.ClientId[i]})
    }
    clientIds.EntityData.Leafs = types.NewOrderedMap()

    clientIds.EntityData.YListKeys = []string {}

    return &(clientIds.EntityData)
}

// Cofo_Nodes_Node_ClientIds_ClientId
// COFO client id
type Cofo_Nodes_Node_ClientIds_ClientId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Client identifier. The type is
    // interface{} with range: 1..7.
    Id interface{}

    // Connection Handle. The type is interface{} with range: 0..4294967295.
    ConnectionHandle interface{}

    // Peer Handle. The type is interface{} with range: 0..4294967295.
    PeerHandle interface{}

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Purge timeout. The type is interface{} with range: 0..4294967295.
    PurgeTimeout interface{}

    // host client. The type is bool.
    HostClient interface{}

    // Connection state. The type is interface{} with range: 0..4294967295.
    ConnectionState interface{}

    // Topic Subscribed. The type is slice of interface{} with range:
    // 0..4294967295.
    TopicSubscribed []interface{}

    // Topic Published. The type is slice of interface{} with range:
    // 0..4294967295.
    TopicPublished []interface{}
}

func (clientId *Cofo_Nodes_Node_ClientIds_ClientId) GetEntityData() *types.CommonEntityData {
    clientId.EntityData.YFilter = clientId.YFilter
    clientId.EntityData.YangName = "client-id"
    clientId.EntityData.BundleName = "cisco_ios_xr"
    clientId.EntityData.ParentYangName = "client-ids"
    clientId.EntityData.SegmentPath = "client-id" + types.AddKeyToken(clientId.Id, "id")
    clientId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientId.EntityData.Children = types.NewOrderedMap()
    clientId.EntityData.Leafs = types.NewOrderedMap()
    clientId.EntityData.Leafs.Append("id", types.YLeaf{"Id", clientId.Id})
    clientId.EntityData.Leafs.Append("connection-handle", types.YLeaf{"ConnectionHandle", clientId.ConnectionHandle})
    clientId.EntityData.Leafs.Append("peer-handle", types.YLeaf{"PeerHandle", clientId.PeerHandle})
    clientId.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", clientId.ClientId})
    clientId.EntityData.Leafs.Append("purge-timeout", types.YLeaf{"PurgeTimeout", clientId.PurgeTimeout})
    clientId.EntityData.Leafs.Append("host-client", types.YLeaf{"HostClient", clientId.HostClient})
    clientId.EntityData.Leafs.Append("connection-state", types.YLeaf{"ConnectionState", clientId.ConnectionState})
    clientId.EntityData.Leafs.Append("topic-subscribed", types.YLeaf{"TopicSubscribed", clientId.TopicSubscribed})
    clientId.EntityData.Leafs.Append("topic-published", types.YLeaf{"TopicPublished", clientId.TopicPublished})

    clientId.EntityData.YListKeys = []string {"Id"}

    return &(clientId.EntityData)
}

// Cofo_Nodes_Node_TopicIds
// COFO SDR database
type Cofo_Nodes_Node_TopicIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // COFO topic id. The type is slice of Cofo_Nodes_Node_TopicIds_TopicId.
    TopicId []*Cofo_Nodes_Node_TopicIds_TopicId
}

func (topicIds *Cofo_Nodes_Node_TopicIds) GetEntityData() *types.CommonEntityData {
    topicIds.EntityData.YFilter = topicIds.YFilter
    topicIds.EntityData.YangName = "topic-ids"
    topicIds.EntityData.BundleName = "cisco_ios_xr"
    topicIds.EntityData.ParentYangName = "node"
    topicIds.EntityData.SegmentPath = "topic-ids"
    topicIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topicIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topicIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topicIds.EntityData.Children = types.NewOrderedMap()
    topicIds.EntityData.Children.Append("topic-id", types.YChild{"TopicId", nil})
    for i := range topicIds.TopicId {
        topicIds.EntityData.Children.Append(types.GetSegmentPath(topicIds.TopicId[i]), types.YChild{"TopicId", topicIds.TopicId[i]})
    }
    topicIds.EntityData.Leafs = types.NewOrderedMap()

    topicIds.EntityData.YListKeys = []string {}

    return &(topicIds.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId
// COFO topic id
type Cofo_Nodes_Node_TopicIds_TopicId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specific Topic identifier. The type is interface{}
    // with range: 1..8.
    Id interface{}

    // Topic ID. The type is interface{} with range: 0..4294967295.
    TopicId interface{}

    // Database info struct. The type is slice of
    // Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct.
    DatabaseInfoStruct []*Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct
}

func (topicId *Cofo_Nodes_Node_TopicIds_TopicId) GetEntityData() *types.CommonEntityData {
    topicId.EntityData.YFilter = topicId.YFilter
    topicId.EntityData.YangName = "topic-id"
    topicId.EntityData.BundleName = "cisco_ios_xr"
    topicId.EntityData.ParentYangName = "topic-ids"
    topicId.EntityData.SegmentPath = "topic-id" + types.AddKeyToken(topicId.Id, "id")
    topicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topicId.EntityData.Children = types.NewOrderedMap()
    topicId.EntityData.Children.Append("database-info-struct", types.YChild{"DatabaseInfoStruct", nil})
    for i := range topicId.DatabaseInfoStruct {
        topicId.EntityData.Children.Append(types.GetSegmentPath(topicId.DatabaseInfoStruct[i]), types.YChild{"DatabaseInfoStruct", topicId.DatabaseInfoStruct[i]})
    }
    topicId.EntityData.Leafs = types.NewOrderedMap()
    topicId.EntityData.Leafs.Append("id", types.YLeaf{"Id", topicId.Id})
    topicId.EntityData.Leafs.Append("topic-id", types.YLeaf{"TopicId", topicId.TopicId})

    topicId.EntityData.YListKeys = []string {"Id"}

    return &(topicId.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct
// Database info struct
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SDR ID. The type is interface{} with range: 0..4294967295.
    SdrId interface{}

    // Client db info struct. The type is slice of
    // Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct.
    ClientDbInfoStruct []*Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct
}

func (databaseInfoStruct *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct) GetEntityData() *types.CommonEntityData {
    databaseInfoStruct.EntityData.YFilter = databaseInfoStruct.YFilter
    databaseInfoStruct.EntityData.YangName = "database-info-struct"
    databaseInfoStruct.EntityData.BundleName = "cisco_ios_xr"
    databaseInfoStruct.EntityData.ParentYangName = "topic-id"
    databaseInfoStruct.EntityData.SegmentPath = "database-info-struct"
    databaseInfoStruct.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseInfoStruct.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseInfoStruct.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseInfoStruct.EntityData.Children = types.NewOrderedMap()
    databaseInfoStruct.EntityData.Children.Append("client-db-info-struct", types.YChild{"ClientDbInfoStruct", nil})
    for i := range databaseInfoStruct.ClientDbInfoStruct {
        databaseInfoStruct.EntityData.Children.Append(types.GetSegmentPath(databaseInfoStruct.ClientDbInfoStruct[i]), types.YChild{"ClientDbInfoStruct", databaseInfoStruct.ClientDbInfoStruct[i]})
    }
    databaseInfoStruct.EntityData.Leafs = types.NewOrderedMap()
    databaseInfoStruct.EntityData.Leafs.Append("sdr-id", types.YLeaf{"SdrId", databaseInfoStruct.SdrId})

    databaseInfoStruct.EntityData.YListKeys = []string {}

    return &(databaseInfoStruct.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct
// Client db info struct
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total objects. The type is interface{} with range: 0..4294967295.
    TotalObjects interface{}

    // Total valid items in db. The type is interface{} with range: 0..4294967295.
    TotalValidItemsInDb interface{}

    // Cofo object published array. The type is slice of
    // Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray.
    CofoObjectPublishedArray []*Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray
}

func (clientDbInfoStruct *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct) GetEntityData() *types.CommonEntityData {
    clientDbInfoStruct.EntityData.YFilter = clientDbInfoStruct.YFilter
    clientDbInfoStruct.EntityData.YangName = "client-db-info-struct"
    clientDbInfoStruct.EntityData.BundleName = "cisco_ios_xr"
    clientDbInfoStruct.EntityData.ParentYangName = "database-info-struct"
    clientDbInfoStruct.EntityData.SegmentPath = "client-db-info-struct"
    clientDbInfoStruct.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientDbInfoStruct.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientDbInfoStruct.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientDbInfoStruct.EntityData.Children = types.NewOrderedMap()
    clientDbInfoStruct.EntityData.Children.Append("cofo-object-published-array", types.YChild{"CofoObjectPublishedArray", nil})
    for i := range clientDbInfoStruct.CofoObjectPublishedArray {
        clientDbInfoStruct.EntityData.Children.Append(types.GetSegmentPath(clientDbInfoStruct.CofoObjectPublishedArray[i]), types.YChild{"CofoObjectPublishedArray", clientDbInfoStruct.CofoObjectPublishedArray[i]})
    }
    clientDbInfoStruct.EntityData.Leafs = types.NewOrderedMap()
    clientDbInfoStruct.EntityData.Leafs.Append("total-objects", types.YLeaf{"TotalObjects", clientDbInfoStruct.TotalObjects})
    clientDbInfoStruct.EntityData.Leafs.Append("total-valid-items-in-db", types.YLeaf{"TotalValidItemsInDb", clientDbInfoStruct.TotalValidItemsInDb})

    clientDbInfoStruct.EntityData.YListKeys = []string {}

    return &(clientDbInfoStruct.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray
// Cofo object published array
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client ID. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Object ID. The type is interface{} with range: 0..4294967295.
    ObjectId interface{}

    // Insert Count. The type is interface{} with range: 0..4294967295.
    InsertCount interface{}

    // Item State. The type is interface{} with range: 0..4294967295.
    ItemState interface{}

    // Cofo infra object key. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    CofoInfraObjectKey interface{}

    // Cofo infra object value. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    CofoInfraObjectValue interface{}

    // Cofo object add time.
    ObjectAddTime Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectAddTime

    // Cofo object delete time.
    ObjectDeleteTime Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectDeleteTime

    // Cofo object txl add time.
    ObjectTxlAddTime Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlAddTime

    // Cofo object txl encode time.
    ObjectTxlEncodeTime Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlEncodeTime
}

func (cofoObjectPublishedArray *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray) GetEntityData() *types.CommonEntityData {
    cofoObjectPublishedArray.EntityData.YFilter = cofoObjectPublishedArray.YFilter
    cofoObjectPublishedArray.EntityData.YangName = "cofo-object-published-array"
    cofoObjectPublishedArray.EntityData.BundleName = "cisco_ios_xr"
    cofoObjectPublishedArray.EntityData.ParentYangName = "client-db-info-struct"
    cofoObjectPublishedArray.EntityData.SegmentPath = "cofo-object-published-array"
    cofoObjectPublishedArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cofoObjectPublishedArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cofoObjectPublishedArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cofoObjectPublishedArray.EntityData.Children = types.NewOrderedMap()
    cofoObjectPublishedArray.EntityData.Children.Append("object-add-time", types.YChild{"ObjectAddTime", &cofoObjectPublishedArray.ObjectAddTime})
    cofoObjectPublishedArray.EntityData.Children.Append("object-delete-time", types.YChild{"ObjectDeleteTime", &cofoObjectPublishedArray.ObjectDeleteTime})
    cofoObjectPublishedArray.EntityData.Children.Append("object-txl-add-time", types.YChild{"ObjectTxlAddTime", &cofoObjectPublishedArray.ObjectTxlAddTime})
    cofoObjectPublishedArray.EntityData.Children.Append("object-txl-encode-time", types.YChild{"ObjectTxlEncodeTime", &cofoObjectPublishedArray.ObjectTxlEncodeTime})
    cofoObjectPublishedArray.EntityData.Leafs = types.NewOrderedMap()
    cofoObjectPublishedArray.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", cofoObjectPublishedArray.ClientId})
    cofoObjectPublishedArray.EntityData.Leafs.Append("object-id", types.YLeaf{"ObjectId", cofoObjectPublishedArray.ObjectId})
    cofoObjectPublishedArray.EntityData.Leafs.Append("insert-count", types.YLeaf{"InsertCount", cofoObjectPublishedArray.InsertCount})
    cofoObjectPublishedArray.EntityData.Leafs.Append("item-state", types.YLeaf{"ItemState", cofoObjectPublishedArray.ItemState})
    cofoObjectPublishedArray.EntityData.Leafs.Append("cofo-infra-object-key", types.YLeaf{"CofoInfraObjectKey", cofoObjectPublishedArray.CofoInfraObjectKey})
    cofoObjectPublishedArray.EntityData.Leafs.Append("cofo-infra-object-value", types.YLeaf{"CofoInfraObjectValue", cofoObjectPublishedArray.CofoInfraObjectValue})

    cofoObjectPublishedArray.EntityData.YListKeys = []string {}

    return &(cofoObjectPublishedArray.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectAddTime
// Cofo object add time
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectAddTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tv sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // tv nsec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (objectAddTime *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectAddTime) GetEntityData() *types.CommonEntityData {
    objectAddTime.EntityData.YFilter = objectAddTime.YFilter
    objectAddTime.EntityData.YangName = "object-add-time"
    objectAddTime.EntityData.BundleName = "cisco_ios_xr"
    objectAddTime.EntityData.ParentYangName = "cofo-object-published-array"
    objectAddTime.EntityData.SegmentPath = "object-add-time"
    objectAddTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectAddTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectAddTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectAddTime.EntityData.Children = types.NewOrderedMap()
    objectAddTime.EntityData.Leafs = types.NewOrderedMap()
    objectAddTime.EntityData.Leafs.Append("tv-sec", types.YLeaf{"TvSec", objectAddTime.TvSec})
    objectAddTime.EntityData.Leafs.Append("tv-nsec", types.YLeaf{"TvNsec", objectAddTime.TvNsec})

    objectAddTime.EntityData.YListKeys = []string {}

    return &(objectAddTime.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectDeleteTime
// Cofo object delete time
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectDeleteTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tv sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // tv nsec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (objectDeleteTime *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectDeleteTime) GetEntityData() *types.CommonEntityData {
    objectDeleteTime.EntityData.YFilter = objectDeleteTime.YFilter
    objectDeleteTime.EntityData.YangName = "object-delete-time"
    objectDeleteTime.EntityData.BundleName = "cisco_ios_xr"
    objectDeleteTime.EntityData.ParentYangName = "cofo-object-published-array"
    objectDeleteTime.EntityData.SegmentPath = "object-delete-time"
    objectDeleteTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectDeleteTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectDeleteTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectDeleteTime.EntityData.Children = types.NewOrderedMap()
    objectDeleteTime.EntityData.Leafs = types.NewOrderedMap()
    objectDeleteTime.EntityData.Leafs.Append("tv-sec", types.YLeaf{"TvSec", objectDeleteTime.TvSec})
    objectDeleteTime.EntityData.Leafs.Append("tv-nsec", types.YLeaf{"TvNsec", objectDeleteTime.TvNsec})

    objectDeleteTime.EntityData.YListKeys = []string {}

    return &(objectDeleteTime.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlAddTime
// Cofo object txl add time
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlAddTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tv sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // tv nsec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (objectTxlAddTime *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlAddTime) GetEntityData() *types.CommonEntityData {
    objectTxlAddTime.EntityData.YFilter = objectTxlAddTime.YFilter
    objectTxlAddTime.EntityData.YangName = "object-txl-add-time"
    objectTxlAddTime.EntityData.BundleName = "cisco_ios_xr"
    objectTxlAddTime.EntityData.ParentYangName = "cofo-object-published-array"
    objectTxlAddTime.EntityData.SegmentPath = "object-txl-add-time"
    objectTxlAddTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTxlAddTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTxlAddTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTxlAddTime.EntityData.Children = types.NewOrderedMap()
    objectTxlAddTime.EntityData.Leafs = types.NewOrderedMap()
    objectTxlAddTime.EntityData.Leafs.Append("tv-sec", types.YLeaf{"TvSec", objectTxlAddTime.TvSec})
    objectTxlAddTime.EntityData.Leafs.Append("tv-nsec", types.YLeaf{"TvNsec", objectTxlAddTime.TvNsec})

    objectTxlAddTime.EntityData.YListKeys = []string {}

    return &(objectTxlAddTime.EntityData)
}

// Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlEncodeTime
// Cofo object txl encode time
type Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlEncodeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // tv sec. The type is interface{} with range: 0..4294967295.
    TvSec interface{}

    // tv nsec. The type is interface{} with range: 0..4294967295.
    TvNsec interface{}
}

func (objectTxlEncodeTime *Cofo_Nodes_Node_TopicIds_TopicId_DatabaseInfoStruct_ClientDbInfoStruct_CofoObjectPublishedArray_ObjectTxlEncodeTime) GetEntityData() *types.CommonEntityData {
    objectTxlEncodeTime.EntityData.YFilter = objectTxlEncodeTime.YFilter
    objectTxlEncodeTime.EntityData.YangName = "object-txl-encode-time"
    objectTxlEncodeTime.EntityData.BundleName = "cisco_ios_xr"
    objectTxlEncodeTime.EntityData.ParentYangName = "cofo-object-published-array"
    objectTxlEncodeTime.EntityData.SegmentPath = "object-txl-encode-time"
    objectTxlEncodeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objectTxlEncodeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objectTxlEncodeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objectTxlEncodeTime.EntityData.Children = types.NewOrderedMap()
    objectTxlEncodeTime.EntityData.Leafs = types.NewOrderedMap()
    objectTxlEncodeTime.EntityData.Leafs.Append("tv-sec", types.YLeaf{"TvSec", objectTxlEncodeTime.TvSec})
    objectTxlEncodeTime.EntityData.Leafs.Append("tv-nsec", types.YLeaf{"TvNsec", objectTxlEncodeTime.TvNsec})

    objectTxlEncodeTime.EntityData.YListKeys = []string {}

    return &(objectTxlEncodeTime.EntityData)
}

