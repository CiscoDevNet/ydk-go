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
    Node []FileSystem_Node
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

    fileSystem.EntityData.Children = make(map[string]types.YChild)
    fileSystem.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range fileSystem.Node {
        fileSystem.EntityData.Children[types.GetSegmentPath(&fileSystem.Node[i])] = types.YChild{"Node", &fileSystem.Node[i]}
    }
    fileSystem.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(fileSystem.EntityData)
}

// FileSystem_Node
// Node ID
type FileSystem_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Available file systems. The type is slice of FileSystem_Node_FileSystem.
    FileSystem []FileSystem_Node_FileSystem_
}

func (node *FileSystem_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "file-system"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["file-system"] = types.YChild{"FileSystem", nil}
    for i := range node.FileSystem {
        node.EntityData.Children[types.GetSegmentPath(&node.FileSystem[i])] = types.YChild{"FileSystem", &node.FileSystem[i]}
    }
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// FileSystem_Node_FileSystem_
// Available file systems
type FileSystem_Node_FileSystem_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size of the file system in bytes. The type is string. Units are byte.
    Size interface{}

    // Free space in the file system in bytes. The type is string. Units are byte.
    Free interface{}

    // Type of file system. The type is string.
    Type_ interface{}

    // Flags of file system. The type is string.
    Flags interface{}

    // Prefixes of file system. The type is string.
    Prefixes interface{}
}

func (fileSystem_ *FileSystem_Node_FileSystem_) GetEntityData() *types.CommonEntityData {
    fileSystem_.EntityData.YFilter = fileSystem_.YFilter
    fileSystem_.EntityData.YangName = "file-system"
    fileSystem_.EntityData.BundleName = "cisco_ios_xr"
    fileSystem_.EntityData.ParentYangName = "node"
    fileSystem_.EntityData.SegmentPath = "file-system"
    fileSystem_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fileSystem_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fileSystem_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fileSystem_.EntityData.Children = make(map[string]types.YChild)
    fileSystem_.EntityData.Leafs = make(map[string]types.YLeaf)
    fileSystem_.EntityData.Leafs["size"] = types.YLeaf{"Size", fileSystem_.Size}
    fileSystem_.EntityData.Leafs["free"] = types.YLeaf{"Free", fileSystem_.Free}
    fileSystem_.EntityData.Leafs["type"] = types.YLeaf{"Type_", fileSystem_.Type_}
    fileSystem_.EntityData.Leafs["flags"] = types.YLeaf{"Flags", fileSystem_.Flags}
    fileSystem_.EntityData.Leafs["prefixes"] = types.YLeaf{"Prefixes", fileSystem_.Prefixes}
    return &(fileSystem_.EntityData)
}

