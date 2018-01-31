// This module contains a collection of YANG definitions for
// monitoring the checkpoint archives in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package checkpoint_archive_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package checkpoint_archive_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-checkpoint-archive-oper checkpoint-archives}", reflect.TypeOf(CheckpointArchives{}))
    ydk.RegisterEntity("Cisco-IOS-XE-checkpoint-archive-oper:checkpoint-archives", reflect.TypeOf(CheckpointArchives{}))
}

// CheckpointArchives
// Contents of the checkpoint archive information base
type CheckpointArchives struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The maxium number of archives. The type is interface{} with range: 0..255.
    Max interface{}

    // The current number of archives. The type is interface{} with range: 0..255.
    Current interface{}

    // The most recent archive. The type is string.
    Recent interface{}

    // Archive information.
    Archives CheckpointArchives_Archives
}

func (checkpointArchives *CheckpointArchives) GetFilter() yfilter.YFilter { return checkpointArchives.YFilter }

func (checkpointArchives *CheckpointArchives) SetFilter(yf yfilter.YFilter) { checkpointArchives.YFilter = yf }

func (checkpointArchives *CheckpointArchives) GetGoName(yname string) string {
    if yname == "max" { return "Max" }
    if yname == "current" { return "Current" }
    if yname == "recent" { return "Recent" }
    if yname == "archives" { return "Archives" }
    return ""
}

func (checkpointArchives *CheckpointArchives) GetSegmentPath() string {
    return "Cisco-IOS-XE-checkpoint-archive-oper:checkpoint-archives"
}

func (checkpointArchives *CheckpointArchives) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "archives" {
        return &checkpointArchives.Archives
    }
    return nil
}

func (checkpointArchives *CheckpointArchives) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["archives"] = &checkpointArchives.Archives
    return children
}

func (checkpointArchives *CheckpointArchives) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max"] = checkpointArchives.Max
    leafs["current"] = checkpointArchives.Current
    leafs["recent"] = checkpointArchives.Recent
    return leafs
}

func (checkpointArchives *CheckpointArchives) GetBundleName() string { return "cisco_ios_xe" }

func (checkpointArchives *CheckpointArchives) GetYangName() string { return "checkpoint-archives" }

func (checkpointArchives *CheckpointArchives) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (checkpointArchives *CheckpointArchives) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (checkpointArchives *CheckpointArchives) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (checkpointArchives *CheckpointArchives) SetParent(parent types.Entity) { checkpointArchives.parent = parent }

func (checkpointArchives *CheckpointArchives) GetParent() types.Entity { return checkpointArchives.parent }

func (checkpointArchives *CheckpointArchives) GetParentYangName() string { return "Cisco-IOS-XE-checkpoint-archive-oper" }

// CheckpointArchives_Archives
// Archive information
type CheckpointArchives_Archives struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of archives. The type is slice of CheckpointArchives_Archives_Archive.
    Archive []CheckpointArchives_Archives_Archive
}

func (archives *CheckpointArchives_Archives) GetFilter() yfilter.YFilter { return archives.YFilter }

func (archives *CheckpointArchives_Archives) SetFilter(yf yfilter.YFilter) { archives.YFilter = yf }

func (archives *CheckpointArchives_Archives) GetGoName(yname string) string {
    if yname == "archive" { return "Archive" }
    return ""
}

func (archives *CheckpointArchives_Archives) GetSegmentPath() string {
    return "archives"
}

func (archives *CheckpointArchives_Archives) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "archive" {
        for _, c := range archives.Archive {
            if archives.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CheckpointArchives_Archives_Archive{}
        archives.Archive = append(archives.Archive, child)
        return &archives.Archive[len(archives.Archive)-1]
    }
    return nil
}

func (archives *CheckpointArchives_Archives) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range archives.Archive {
        children[archives.Archive[i].GetSegmentPath()] = &archives.Archive[i]
    }
    return children
}

func (archives *CheckpointArchives_Archives) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (archives *CheckpointArchives_Archives) GetBundleName() string { return "cisco_ios_xe" }

func (archives *CheckpointArchives_Archives) GetYangName() string { return "archives" }

func (archives *CheckpointArchives_Archives) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (archives *CheckpointArchives_Archives) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (archives *CheckpointArchives_Archives) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (archives *CheckpointArchives_Archives) SetParent(parent types.Entity) { archives.parent = parent }

func (archives *CheckpointArchives_Archives) GetParent() types.Entity { return archives.parent }

func (archives *CheckpointArchives_Archives) GetParentYangName() string { return "checkpoint-archives" }

// CheckpointArchives_Archives_Archive
// List of archives
type CheckpointArchives_Archives_Archive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The archive number. The type is interface{} with
    // range: 0..65535.
    Number interface{}

    // The name of the archive. The type is string.
    Name interface{}
}

func (archive *CheckpointArchives_Archives_Archive) GetFilter() yfilter.YFilter { return archive.YFilter }

func (archive *CheckpointArchives_Archives_Archive) SetFilter(yf yfilter.YFilter) { archive.YFilter = yf }

func (archive *CheckpointArchives_Archives_Archive) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "name" { return "Name" }
    return ""
}

func (archive *CheckpointArchives_Archives_Archive) GetSegmentPath() string {
    return "archive" + "[number='" + fmt.Sprintf("%v", archive.Number) + "']"
}

func (archive *CheckpointArchives_Archives_Archive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (archive *CheckpointArchives_Archives_Archive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (archive *CheckpointArchives_Archives_Archive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = archive.Number
    leafs["name"] = archive.Name
    return leafs
}

func (archive *CheckpointArchives_Archives_Archive) GetBundleName() string { return "cisco_ios_xe" }

func (archive *CheckpointArchives_Archives_Archive) GetYangName() string { return "archive" }

func (archive *CheckpointArchives_Archives_Archive) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (archive *CheckpointArchives_Archives_Archive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (archive *CheckpointArchives_Archives_Archive) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (archive *CheckpointArchives_Archives_Archive) SetParent(parent types.Entity) { archive.parent = parent }

func (archive *CheckpointArchives_Archives_Archive) GetParent() types.Entity { return archive.parent }

func (archive *CheckpointArchives_Archives_Archive) GetParentYangName() string { return "archives" }

