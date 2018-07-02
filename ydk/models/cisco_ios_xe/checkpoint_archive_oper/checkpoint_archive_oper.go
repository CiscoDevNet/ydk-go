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
    EntityData types.CommonEntityData
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

func (checkpointArchives *CheckpointArchives) GetEntityData() *types.CommonEntityData {
    checkpointArchives.EntityData.YFilter = checkpointArchives.YFilter
    checkpointArchives.EntityData.YangName = "checkpoint-archives"
    checkpointArchives.EntityData.BundleName = "cisco_ios_xe"
    checkpointArchives.EntityData.ParentYangName = "Cisco-IOS-XE-checkpoint-archive-oper"
    checkpointArchives.EntityData.SegmentPath = "Cisco-IOS-XE-checkpoint-archive-oper:checkpoint-archives"
    checkpointArchives.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    checkpointArchives.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    checkpointArchives.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    checkpointArchives.EntityData.Children = types.NewOrderedMap()
    checkpointArchives.EntityData.Children.Append("archives", types.YChild{"Archives", &checkpointArchives.Archives})
    checkpointArchives.EntityData.Leafs = types.NewOrderedMap()
    checkpointArchives.EntityData.Leafs.Append("max", types.YLeaf{"Max", checkpointArchives.Max})
    checkpointArchives.EntityData.Leafs.Append("current", types.YLeaf{"Current", checkpointArchives.Current})
    checkpointArchives.EntityData.Leafs.Append("recent", types.YLeaf{"Recent", checkpointArchives.Recent})

    checkpointArchives.EntityData.YListKeys = []string {}

    return &(checkpointArchives.EntityData)
}

// CheckpointArchives_Archives
// Archive information
type CheckpointArchives_Archives struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of archives. The type is slice of CheckpointArchives_Archives_Archive.
    Archive []*CheckpointArchives_Archives_Archive
}

func (archives *CheckpointArchives_Archives) GetEntityData() *types.CommonEntityData {
    archives.EntityData.YFilter = archives.YFilter
    archives.EntityData.YangName = "archives"
    archives.EntityData.BundleName = "cisco_ios_xe"
    archives.EntityData.ParentYangName = "checkpoint-archives"
    archives.EntityData.SegmentPath = "archives"
    archives.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    archives.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    archives.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    archives.EntityData.Children = types.NewOrderedMap()
    archives.EntityData.Children.Append("archive", types.YChild{"Archive", nil})
    for i := range archives.Archive {
        archives.EntityData.Children.Append(types.GetSegmentPath(archives.Archive[i]), types.YChild{"Archive", archives.Archive[i]})
    }
    archives.EntityData.Leafs = types.NewOrderedMap()

    archives.EntityData.YListKeys = []string {}

    return &(archives.EntityData)
}

// CheckpointArchives_Archives_Archive
// List of archives
type CheckpointArchives_Archives_Archive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The archive number. The type is interface{} with
    // range: 0..65535.
    Number interface{}

    // The name of the archive. The type is string.
    Name interface{}
}

func (archive *CheckpointArchives_Archives_Archive) GetEntityData() *types.CommonEntityData {
    archive.EntityData.YFilter = archive.YFilter
    archive.EntityData.YangName = "archive"
    archive.EntityData.BundleName = "cisco_ios_xe"
    archive.EntityData.ParentYangName = "archives"
    archive.EntityData.SegmentPath = "archive" + types.AddKeyToken(archive.Number, "number")
    archive.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    archive.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    archive.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    archive.EntityData.Children = types.NewOrderedMap()
    archive.EntityData.Leafs = types.NewOrderedMap()
    archive.EntityData.Leafs.Append("number", types.YLeaf{"Number", archive.Number})
    archive.EntityData.Leafs.Append("name", types.YLeaf{"Name", archive.Name})

    archive.EntityData.YListKeys = []string {"Number"}

    return &(archive.EntityData)
}

