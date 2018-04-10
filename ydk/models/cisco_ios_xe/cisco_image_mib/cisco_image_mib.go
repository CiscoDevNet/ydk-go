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
    Ciscoimagetable CISCOIMAGEMIB_Ciscoimagetable
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

    cISCOIMAGEMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIMAGEMIB.EntityData.Children["ciscoImageTable"] = types.YChild{"Ciscoimagetable", &cISCOIMAGEMIB.Ciscoimagetable}
    cISCOIMAGEMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIMAGEMIB.EntityData)
}

// CISCOIMAGEMIB_Ciscoimagetable
// A table provides content information describing the
// executing IOS image.
type CISCOIMAGEMIB_Ciscoimagetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A image characteristic string entry. The type is slice of
    // CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry.
    Ciscoimageentry []CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetEntityData() *types.CommonEntityData {
    ciscoimagetable.EntityData.YFilter = ciscoimagetable.YFilter
    ciscoimagetable.EntityData.YangName = "ciscoImageTable"
    ciscoimagetable.EntityData.BundleName = "cisco_ios_xe"
    ciscoimagetable.EntityData.ParentYangName = "CISCO-IMAGE-MIB"
    ciscoimagetable.EntityData.SegmentPath = "ciscoImageTable"
    ciscoimagetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoimagetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoimagetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoimagetable.EntityData.Children = make(map[string]types.YChild)
    ciscoimagetable.EntityData.Children["ciscoImageEntry"] = types.YChild{"Ciscoimageentry", nil}
    for i := range ciscoimagetable.Ciscoimageentry {
        ciscoimagetable.EntityData.Children[types.GetSegmentPath(&ciscoimagetable.Ciscoimageentry[i])] = types.YChild{"Ciscoimageentry", &ciscoimagetable.Ciscoimageentry[i]}
    }
    ciscoimagetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoimagetable.EntityData)
}

// CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry
// A image characteristic string entry.
type CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A sequence number for each string stored in the
    // IOS image. The type is interface{} with range: 0..2147483647.
    Ciscoimageindex interface{}

    // The string of this entry. The type is string.
    Ciscoimagestring interface{}
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetEntityData() *types.CommonEntityData {
    ciscoimageentry.EntityData.YFilter = ciscoimageentry.YFilter
    ciscoimageentry.EntityData.YangName = "ciscoImageEntry"
    ciscoimageentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoimageentry.EntityData.ParentYangName = "ciscoImageTable"
    ciscoimageentry.EntityData.SegmentPath = "ciscoImageEntry" + "[ciscoImageIndex='" + fmt.Sprintf("%v", ciscoimageentry.Ciscoimageindex) + "']"
    ciscoimageentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoimageentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoimageentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoimageentry.EntityData.Children = make(map[string]types.YChild)
    ciscoimageentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoimageentry.EntityData.Leafs["ciscoImageIndex"] = types.YLeaf{"Ciscoimageindex", ciscoimageentry.Ciscoimageindex}
    ciscoimageentry.EntityData.Leafs["ciscoImageString"] = types.YLeaf{"Ciscoimagestring", ciscoimageentry.Ciscoimagestring}
    return &(ciscoimageentry.EntityData)
}

