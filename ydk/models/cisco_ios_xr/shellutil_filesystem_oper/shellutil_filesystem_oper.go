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
    parent types.Entity
    YFilter yfilter.YFilter

    // Node ID. The type is slice of FileSystem_Node.
    Node []FileSystem_Node
}

func (fileSystem *FileSystem) GetFilter() yfilter.YFilter { return fileSystem.YFilter }

func (fileSystem *FileSystem) SetFilter(yf yfilter.YFilter) { fileSystem.YFilter = yf }

func (fileSystem *FileSystem) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (fileSystem *FileSystem) GetSegmentPath() string {
    return "Cisco-IOS-XR-shellutil-filesystem-oper:file-system"
}

func (fileSystem *FileSystem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range fileSystem.Node {
            if fileSystem.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FileSystem_Node{}
        fileSystem.Node = append(fileSystem.Node, child)
        return &fileSystem.Node[len(fileSystem.Node)-1]
    }
    return nil
}

func (fileSystem *FileSystem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fileSystem.Node {
        children[fileSystem.Node[i].GetSegmentPath()] = &fileSystem.Node[i]
    }
    return children
}

func (fileSystem *FileSystem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fileSystem *FileSystem) GetBundleName() string { return "cisco_ios_xr" }

func (fileSystem *FileSystem) GetYangName() string { return "file-system" }

func (fileSystem *FileSystem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileSystem *FileSystem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileSystem *FileSystem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileSystem *FileSystem) SetParent(parent types.Entity) { fileSystem.parent = parent }

func (fileSystem *FileSystem) GetParent() types.Entity { return fileSystem.parent }

func (fileSystem *FileSystem) GetParentYangName() string { return "Cisco-IOS-XR-shellutil-filesystem-oper" }

// FileSystem_Node
// Node ID
type FileSystem_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Available file systems. The type is slice of FileSystem_Node_FileSystem.
    FileSystem []FileSystem_Node_FileSystem
}

func (node *FileSystem_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *FileSystem_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *FileSystem_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "file-system" { return "FileSystem" }
    return ""
}

func (node *FileSystem_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *FileSystem_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "file-system" {
        for _, c := range node.FileSystem {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := FileSystem_Node_FileSystem{}
        node.FileSystem = append(node.FileSystem, child)
        return &node.FileSystem[len(node.FileSystem)-1]
    }
    return nil
}

func (node *FileSystem_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.FileSystem {
        children[node.FileSystem[i].GetSegmentPath()] = &node.FileSystem[i]
    }
    return children
}

func (node *FileSystem_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *FileSystem_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *FileSystem_Node) GetYangName() string { return "node" }

func (node *FileSystem_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *FileSystem_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *FileSystem_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *FileSystem_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *FileSystem_Node) GetParent() types.Entity { return node.parent }

func (node *FileSystem_Node) GetParentYangName() string { return "file-system" }

// FileSystem_Node_FileSystem
// Available file systems
type FileSystem_Node_FileSystem struct {
    parent types.Entity
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

func (fileSystem *FileSystem_Node_FileSystem) GetFilter() yfilter.YFilter { return fileSystem.YFilter }

func (fileSystem *FileSystem_Node_FileSystem) SetFilter(yf yfilter.YFilter) { fileSystem.YFilter = yf }

func (fileSystem *FileSystem_Node_FileSystem) GetGoName(yname string) string {
    if yname == "size" { return "Size" }
    if yname == "free" { return "Free" }
    if yname == "type" { return "Type" }
    if yname == "flags" { return "Flags" }
    if yname == "prefixes" { return "Prefixes" }
    return ""
}

func (fileSystem *FileSystem_Node_FileSystem) GetSegmentPath() string {
    return "file-system"
}

func (fileSystem *FileSystem_Node_FileSystem) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fileSystem *FileSystem_Node_FileSystem) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fileSystem *FileSystem_Node_FileSystem) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["size"] = fileSystem.Size
    leafs["free"] = fileSystem.Free
    leafs["type"] = fileSystem.Type
    leafs["flags"] = fileSystem.Flags
    leafs["prefixes"] = fileSystem.Prefixes
    return leafs
}

func (fileSystem *FileSystem_Node_FileSystem) GetBundleName() string { return "cisco_ios_xr" }

func (fileSystem *FileSystem_Node_FileSystem) GetYangName() string { return "file-system" }

func (fileSystem *FileSystem_Node_FileSystem) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fileSystem *FileSystem_Node_FileSystem) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fileSystem *FileSystem_Node_FileSystem) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fileSystem *FileSystem_Node_FileSystem) SetParent(parent types.Entity) { fileSystem.parent = parent }

func (fileSystem *FileSystem_Node_FileSystem) GetParent() types.Entity { return fileSystem.parent }

func (fileSystem *FileSystem_Node_FileSystem) GetParentYangName() string { return "node" }

