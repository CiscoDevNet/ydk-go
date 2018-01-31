// This MIB Module is a supplement to the
// ATM-MIB.
package cisco_ietf_atm2_pvctrap_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_atm2_pvctrap_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-ATM2-PVCTRAP-MIB CISCO-IETF-ATM2-PVCTRAP-MIB}", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIB{}))
    ydk.RegisterEntity("CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIB{}))
}

// CISCOIETFATM2PVCTRAPMIB
type CISCOIETFATM2PVCTRAPMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and an atmVclOperStatus
    // with a value other than `up'.
    Atmcurrentlyfailingpvcltable CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetFilter() yfilter.YFilter { return cISCOIETFATM2PVCTRAPMIB.YFilter }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) SetFilter(yf yfilter.YFilter) { cISCOIETFATM2PVCTRAPMIB.YFilter = yf }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetGoName(yname string) string {
    if yname == "atmCurrentlyFailingPVclTable" { return "Atmcurrentlyfailingpvcltable" }
    return ""
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetSegmentPath() string {
    return "CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB"
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmCurrentlyFailingPVclTable" {
        return &cISCOIETFATM2PVCTRAPMIB.Atmcurrentlyfailingpvcltable
    }
    return nil
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["atmCurrentlyFailingPVclTable"] = &cISCOIETFATM2PVCTRAPMIB.Atmcurrentlyfailingpvcltable
    return children
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB" }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) SetParent(parent types.Entity) { cISCOIETFATM2PVCTRAPMIB.parent = parent }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetParent() types.Entity { return cISCOIETFATM2PVCTRAPMIB.parent }

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetParentYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB" }

// CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and an atmVclOperStatus with a value
// other than `up'.
type CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry in this table represents a VCL for which the atmVclRowStatus is
    // `active', the atmVclConnKind is `pvc', and the atmVclOperStatus is other
    // than `up'. The type is slice of
    // CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry.
    Atmcurrentlyfailingpvclentry []CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetFilter() yfilter.YFilter { return atmcurrentlyfailingpvcltable.YFilter }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) SetFilter(yf yfilter.YFilter) { atmcurrentlyfailingpvcltable.YFilter = yf }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetGoName(yname string) string {
    if yname == "atmCurrentlyFailingPVclEntry" { return "Atmcurrentlyfailingpvclentry" }
    return ""
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetSegmentPath() string {
    return "atmCurrentlyFailingPVclTable"
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "atmCurrentlyFailingPVclEntry" {
        for _, c := range atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry {
            if atmcurrentlyfailingpvcltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry{}
        atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry = append(atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry, child)
        return &atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry[len(atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry)-1]
    }
    return nil
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry {
        children[atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry[i].GetSegmentPath()] = &atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry[i]
    }
    return children
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetBundleName() string { return "cisco_ios_xe" }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetYangName() string { return "atmCurrentlyFailingPVclTable" }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) SetParent(parent types.Entity) { atmcurrentlyfailingpvcltable.parent = parent }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetParent() types.Entity { return atmcurrentlyfailingpvcltable.parent }

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetParentYangName() string { return "CISCO-IETF-ATM2-PVCTRAP-MIB" }

// CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry
// Each entry in this table represents a VCL for which
// the atmVclRowStatus is `active', the atmVclConnKind is
// `pvc', and the atmVclOperStatus is other than `up'.
type CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvci
    Atmvclvci interface{}

    // The time at which this PVCL began to fail. The type is interface{} with
    // range: 0..4294967295.
    Atmcurrentlyfailingpvcltimestamp interface{}

    // The time at which this PVCL began to fail during the PVC Notification
    // interval. The type is interface{} with range: 0..4294967295.
    Atmpreviouslyfailedpvcltimestamp interface{}
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetFilter() yfilter.YFilter { return atmcurrentlyfailingpvclentry.YFilter }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) SetFilter(yf yfilter.YFilter) { atmcurrentlyfailingpvclentry.YFilter = yf }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "atmVclVpi" { return "Atmvclvpi" }
    if yname == "atmVclVci" { return "Atmvclvci" }
    if yname == "atmCurrentlyFailingPVclTimeStamp" { return "Atmcurrentlyfailingpvcltimestamp" }
    if yname == "atmPreviouslyFailedPVclTimeStamp" { return "Atmpreviouslyfailedpvcltimestamp" }
    return ""
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetSegmentPath() string {
    return "atmCurrentlyFailingPVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Atmvclvci) + "']"
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = atmcurrentlyfailingpvclentry.Ifindex
    leafs["atmVclVpi"] = atmcurrentlyfailingpvclentry.Atmvclvpi
    leafs["atmVclVci"] = atmcurrentlyfailingpvclentry.Atmvclvci
    leafs["atmCurrentlyFailingPVclTimeStamp"] = atmcurrentlyfailingpvclentry.Atmcurrentlyfailingpvcltimestamp
    leafs["atmPreviouslyFailedPVclTimeStamp"] = atmcurrentlyfailingpvclentry.Atmpreviouslyfailedpvcltimestamp
    return leafs
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetBundleName() string { return "cisco_ios_xe" }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetYangName() string { return "atmCurrentlyFailingPVclEntry" }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) SetParent(parent types.Entity) { atmcurrentlyfailingpvclentry.parent = parent }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetParent() types.Entity { return atmcurrentlyfailingpvclentry.parent }

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetParentYangName() string { return "atmCurrentlyFailingPVclTable" }

