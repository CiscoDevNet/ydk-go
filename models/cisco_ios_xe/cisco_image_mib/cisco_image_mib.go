// Router image MIB which identify the capabilities
// and characteristics of the image
package cisco_image_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_image_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IMAGE-MIB CISCO-IMAGE-MIB}", reflect.TypeOf(CISCOIMAGEMIB{}))
    ydk.RegisterEntity("CISCO-IMAGE-MIB:CISCO-IMAGE-MIB", reflect.TypeOf(CISCOIMAGEMIB{}))
}

// CISCOIMAGEMIB
type CISCOIMAGEMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table provides content information describing the executing IOS image.
    CiscoImageTable CISCOIMAGEMIB_CiscoImageTable
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetEntityData() *types.CommonEntityData {
    cISCOIMAGEMIB.EntityData.YFilter = cISCOIMAGEMIB.YFilter
    cISCOIMAGEMIB.EntityData.YangName = "CISCO-IMAGE-MIB"
    cISCOIMAGEMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIMAGEMIB.EntityData.ParentYangName = "CISCO-IMAGE-MIB"
    cISCOIMAGEMIB.EntityData.SegmentPath = "CISCO-IMAGE-MIB:CISCO-IMAGE-MIB"
    cISCOIMAGEMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIMAGEMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIMAGEMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIMAGEMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIMAGEMIB.EntityData.Children.Append("ciscoImageTable", types.YChild{"CiscoImageTable", &cISCOIMAGEMIB.CiscoImageTable})
    cISCOIMAGEMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIMAGEMIB.EntityData.YListKeys = []string {}

    return &(cISCOIMAGEMIB.EntityData)
}

// CISCOIMAGEMIB_CiscoImageTable
// A table provides content information describing the
// executing IOS image.
type CISCOIMAGEMIB_CiscoImageTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A image characteristic string entry. The type is slice of
    // CISCOIMAGEMIB_CiscoImageTable_CiscoImageEntry.
    CiscoImageEntry []*CISCOIMAGEMIB_CiscoImageTable_CiscoImageEntry
}

func (ciscoImageTable *CISCOIMAGEMIB_CiscoImageTable) GetEntityData() *types.CommonEntityData {
    ciscoImageTable.EntityData.YFilter = ciscoImageTable.YFilter
    ciscoImageTable.EntityData.YangName = "ciscoImageTable"
    ciscoImageTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoImageTable.EntityData.ParentYangName = "CISCO-IMAGE-MIB"
    ciscoImageTable.EntityData.SegmentPath = "ciscoImageTable"
    ciscoImageTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoImageTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoImageTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoImageTable.EntityData.Children = types.NewOrderedMap()
    ciscoImageTable.EntityData.Children.Append("ciscoImageEntry", types.YChild{"CiscoImageEntry", nil})
    for i := range ciscoImageTable.CiscoImageEntry {
        ciscoImageTable.EntityData.Children.Append(types.GetSegmentPath(ciscoImageTable.CiscoImageEntry[i]), types.YChild{"CiscoImageEntry", ciscoImageTable.CiscoImageEntry[i]})
    }
    ciscoImageTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoImageTable.EntityData.YListKeys = []string {}

    return &(ciscoImageTable.EntityData)
}

// CISCOIMAGEMIB_CiscoImageTable_CiscoImageEntry
// A image characteristic string entry.
type CISCOIMAGEMIB_CiscoImageTable_CiscoImageEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A sequence number for each string stored in the
    // IOS image. The type is interface{} with range: 0..2147483647.
    CiscoImageIndex interface{}

    // The string of this entry. The type is string.
    CiscoImageString interface{}
}

func (ciscoImageEntry *CISCOIMAGEMIB_CiscoImageTable_CiscoImageEntry) GetEntityData() *types.CommonEntityData {
    ciscoImageEntry.EntityData.YFilter = ciscoImageEntry.YFilter
    ciscoImageEntry.EntityData.YangName = "ciscoImageEntry"
    ciscoImageEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoImageEntry.EntityData.ParentYangName = "ciscoImageTable"
    ciscoImageEntry.EntityData.SegmentPath = "ciscoImageEntry" + types.AddKeyToken(ciscoImageEntry.CiscoImageIndex, "ciscoImageIndex")
    ciscoImageEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoImageEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoImageEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoImageEntry.EntityData.Children = types.NewOrderedMap()
    ciscoImageEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoImageEntry.EntityData.Leafs.Append("ciscoImageIndex", types.YLeaf{"CiscoImageIndex", ciscoImageEntry.CiscoImageIndex})
    ciscoImageEntry.EntityData.Leafs.Append("ciscoImageString", types.YLeaf{"CiscoImageString", ciscoImageEntry.CiscoImageString})

    ciscoImageEntry.EntityData.YListKeys = []string {"CiscoImageIndex"}

    return &(ciscoImageEntry.EntityData)
}

