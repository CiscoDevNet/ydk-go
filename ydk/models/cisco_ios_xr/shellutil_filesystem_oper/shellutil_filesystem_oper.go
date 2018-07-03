// This module contains a collection of YANG definitions
// for Cisco IOS-XR shellutil-filesystem package operational data.
// 
// This module contains definitions
// for the following management objects:
//   file-system: List of filesystems
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package shellutil_filesystem_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package shellutil_filesystem_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-shellutil-filesystem-oper file-system}", reflect.TypeOf(FileSystem{}))
    ydk.RegisterEntity("Cisco-IOS-XR-shellutil-filesystem-oper:file-system", reflect.TypeOf(FileSystem{}))
}

// FileSystem
// List of filesystems
type FileSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node ID. The type is slice of FileSystem_Node.
    Node []*FileSystem_Node
}

func (fileSystem *FileSystem) GetEntityData() *types.CommonEntityData {
    fileSystem.EntityData.YFilter = fileSystem.YFilter
    fileSystem.EntityData.YangName = "file-system"
    fileSystem.EntityData.BundleName = "cisco_ios_xr"
    fileSystem.EntityData.ParentYangName = "Cisco-IOS-XR-shellutil-filesystem-oper"
    fileSystem.EntityData.SegmentPath = "Cisco-IOS-XR-shellutil-filesystem-oper:file-system"
    fileSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileSystem.EntityData.Children = types.NewOrderedMap()
    fileSystem.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range fileSystem.Node {
        fileSystem.EntityData.Children.Append(types.GetSegmentPath(fileSystem.Node[i]), types.YChild{"Node", fileSystem.Node[i]})
    }
    fileSystem.EntityData.Leafs = types.NewOrderedMap()

    fileSystem.EntityData.YListKeys = []string {}

    return &(fileSystem.EntityData)
}

// FileSystem_Node
// Node ID
type FileSystem_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Available file systems. The type is slice of FileSystem_Node_FileSystem.
    FileSystem []*FileSystem_Node_FileSystem
}

func (node *FileSystem_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "file-system"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("file-system", types.YChild{"FileSystem", nil})
    for i := range node.FileSystem {
        node.EntityData.Children.Append(types.GetSegmentPath(node.FileSystem[i]), types.YChild{"FileSystem", node.FileSystem[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// FileSystem_Node_FileSystem
// Available file systems
type FileSystem_Node_FileSystem struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size of the file system in bytes. The type is string. Units are byte.
    Size interface{}

    // Free space in the file system in bytes. The type is string. Units are byte.
    Free interface{}

    // Type of file system. The type is string.
    Type interface{}

    // Flags of file system. The type is string.
    Flags interface{}

    // Prefixes of file system. The type is string.
    Prefixes interface{}
}

func (fileSystem *FileSystem_Node_FileSystem) GetEntityData() *types.CommonEntityData {
    fileSystem.EntityData.YFilter = fileSystem.YFilter
    fileSystem.EntityData.YangName = "file-system"
    fileSystem.EntityData.BundleName = "cisco_ios_xr"
    fileSystem.EntityData.ParentYangName = "node"
    fileSystem.EntityData.SegmentPath = "file-system"
    fileSystem.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileSystem.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileSystem.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileSystem.EntityData.Children = types.NewOrderedMap()
    fileSystem.EntityData.Leafs = types.NewOrderedMap()
    fileSystem.EntityData.Leafs.Append("size", types.YLeaf{"Size", fileSystem.Size})
    fileSystem.EntityData.Leafs.Append("free", types.YLeaf{"Free", fileSystem.Free})
    fileSystem.EntityData.Leafs.Append("type", types.YLeaf{"Type", fileSystem.Type})
    fileSystem.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", fileSystem.Flags})
    fileSystem.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", fileSystem.Prefixes})

    fileSystem.EntityData.YListKeys = []string {}

    return &(fileSystem.EntityData)
}

