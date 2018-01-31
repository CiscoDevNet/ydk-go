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
    parent types.Entity
    YFilter yfilter.YFilter

    // A table provides content information describing the executing IOS image.
    Ciscoimagetable CISCOIMAGEMIB_Ciscoimagetable
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetFilter() yfilter.YFilter { return cISCOIMAGEMIB.YFilter }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) SetFilter(yf yfilter.YFilter) { cISCOIMAGEMIB.YFilter = yf }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetGoName(yname string) string {
    if yname == "ciscoImageTable" { return "Ciscoimagetable" }
    return ""
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetSegmentPath() string {
    return "CISCO-IMAGE-MIB:CISCO-IMAGE-MIB"
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoImageTable" {
        return &cISCOIMAGEMIB.Ciscoimagetable
    }
    return nil
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoImageTable"] = &cISCOIMAGEMIB.Ciscoimagetable
    return children
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetYangName() string { return "CISCO-IMAGE-MIB" }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) SetParent(parent types.Entity) { cISCOIMAGEMIB.parent = parent }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetParent() types.Entity { return cISCOIMAGEMIB.parent }

func (cISCOIMAGEMIB *CISCOIMAGEMIB) GetParentYangName() string { return "CISCO-IMAGE-MIB" }

// CISCOIMAGEMIB_Ciscoimagetable
// A table provides content information describing the
// executing IOS image.
type CISCOIMAGEMIB_Ciscoimagetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A image characteristic string entry. The type is slice of
    // CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry.
    Ciscoimageentry []CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetFilter() yfilter.YFilter { return ciscoimagetable.YFilter }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) SetFilter(yf yfilter.YFilter) { ciscoimagetable.YFilter = yf }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetGoName(yname string) string {
    if yname == "ciscoImageEntry" { return "Ciscoimageentry" }
    return ""
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetSegmentPath() string {
    return "ciscoImageTable"
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoImageEntry" {
        for _, c := range ciscoimagetable.Ciscoimageentry {
            if ciscoimagetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry{}
        ciscoimagetable.Ciscoimageentry = append(ciscoimagetable.Ciscoimageentry, child)
        return &ciscoimagetable.Ciscoimageentry[len(ciscoimagetable.Ciscoimageentry)-1]
    }
    return nil
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoimagetable.Ciscoimageentry {
        children[ciscoimagetable.Ciscoimageentry[i].GetSegmentPath()] = &ciscoimagetable.Ciscoimageentry[i]
    }
    return children
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetYangName() string { return "ciscoImageTable" }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) SetParent(parent types.Entity) { ciscoimagetable.parent = parent }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetParent() types.Entity { return ciscoimagetable.parent }

func (ciscoimagetable *CISCOIMAGEMIB_Ciscoimagetable) GetParentYangName() string { return "CISCO-IMAGE-MIB" }

// CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry
// A image characteristic string entry.
type CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A sequence number for each string stored in the
    // IOS image. The type is interface{} with range: 0..2147483647.
    Ciscoimageindex interface{}

    // The string of this entry. The type is string.
    Ciscoimagestring interface{}
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetFilter() yfilter.YFilter { return ciscoimageentry.YFilter }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) SetFilter(yf yfilter.YFilter) { ciscoimageentry.YFilter = yf }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetGoName(yname string) string {
    if yname == "ciscoImageIndex" { return "Ciscoimageindex" }
    if yname == "ciscoImageString" { return "Ciscoimagestring" }
    return ""
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetSegmentPath() string {
    return "ciscoImageEntry" + "[ciscoImageIndex='" + fmt.Sprintf("%v", ciscoimageentry.Ciscoimageindex) + "']"
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoImageIndex"] = ciscoimageentry.Ciscoimageindex
    leafs["ciscoImageString"] = ciscoimageentry.Ciscoimagestring
    return leafs
}

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetYangName() string { return "ciscoImageEntry" }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) SetParent(parent types.Entity) { ciscoimageentry.parent = parent }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetParent() types.Entity { return ciscoimageentry.parent }

func (ciscoimageentry *CISCOIMAGEMIB_Ciscoimagetable_Ciscoimageentry) GetParentYangName() string { return "ciscoImageTable" }

