// This MIB Module is a supplement to the
// CISCO-IETF-ATM2-PVCTRAP-MIB.
package cisco_ietf_atm2_pvctrap_mib_extn

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_atm2_pvctrap_mib_extn"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN}", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIBEXTN{}))
    ydk.RegisterEntity("CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN", reflect.TypeOf(CISCOIETFATM2PVCTRAPMIBEXTN{}))
}

// CISCOIETFATM2PVCTRAPMIBEXTN
type CISCOIETFATM2PVCTRAPMIBEXTN struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and atmVclOperStatus to
    // have changed in the last PVC notification interval.
    Atmcurrentstatuschangepvcltable CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed in the same direction
    // in the last PVC notification interval.
    Atmstatuschangepvclrangetable CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable
}

func (cISCOIETFATM2PVCTRAPMIBEXTN *CISCOIETFATM2PVCTRAPMIBEXTN) GetEntityData() *types.CommonEntityData {
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.YFilter = cISCOIETFATM2PVCTRAPMIBEXTN.YFilter
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.YangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.SegmentPath = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN:CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children["atmCurrentStatusChangePVclTable"] = types.YChild{"Atmcurrentstatuschangepvcltable", &cISCOIETFATM2PVCTRAPMIBEXTN.Atmcurrentstatuschangepvcltable}
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children["atmStatusChangePVclRangeTable"] = types.YChild{"Atmstatuschangepvclrangetable", &cISCOIETFATM2PVCTRAPMIBEXTN.Atmstatuschangepvclrangetable}
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFATM2PVCTRAPMIBEXTN.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last PVC notification interval. The
    // type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry.
    Atmcurrentstatuschangepvclentry []CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry
}

func (atmcurrentstatuschangepvcltable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable) GetEntityData() *types.CommonEntityData {
    atmcurrentstatuschangepvcltable.EntityData.YFilter = atmcurrentstatuschangepvcltable.YFilter
    atmcurrentstatuschangepvcltable.EntityData.YangName = "atmCurrentStatusChangePVclTable"
    atmcurrentstatuschangepvcltable.EntityData.BundleName = "cisco_ios_xe"
    atmcurrentstatuschangepvcltable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    atmcurrentstatuschangepvcltable.EntityData.SegmentPath = "atmCurrentStatusChangePVclTable"
    atmcurrentstatuschangepvcltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmcurrentstatuschangepvcltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmcurrentstatuschangepvcltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmcurrentstatuschangepvcltable.EntityData.Children = make(map[string]types.YChild)
    atmcurrentstatuschangepvcltable.EntityData.Children["atmCurrentStatusChangePVclEntry"] = types.YChild{"Atmcurrentstatuschangepvclentry", nil}
    for i := range atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry {
        atmcurrentstatuschangepvcltable.EntityData.Children[types.GetSegmentPath(&atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry[i])] = types.YChild{"Atmcurrentstatuschangepvclentry", &atmcurrentstatuschangepvcltable.Atmcurrentstatuschangepvclentry[i]}
    }
    atmcurrentstatuschangepvcltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmcurrentstatuschangepvcltable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry struct {
    EntityData types.CommonEntityData
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

    // The number of state transitions that has happened  on this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Atmpvclstatustransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclstatuschangestart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    Atmpvclstatuschangeend interface{}
}

func (atmcurrentstatuschangepvclentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmcurrentstatuschangepvcltable_Atmcurrentstatuschangepvclentry) GetEntityData() *types.CommonEntityData {
    atmcurrentstatuschangepvclentry.EntityData.YFilter = atmcurrentstatuschangepvclentry.YFilter
    atmcurrentstatuschangepvclentry.EntityData.YangName = "atmCurrentStatusChangePVclEntry"
    atmcurrentstatuschangepvclentry.EntityData.BundleName = "cisco_ios_xe"
    atmcurrentstatuschangepvclentry.EntityData.ParentYangName = "atmCurrentStatusChangePVclTable"
    atmcurrentstatuschangepvclentry.EntityData.SegmentPath = "atmCurrentStatusChangePVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmcurrentstatuschangepvclentry.Atmvclvci) + "']"
    atmcurrentstatuschangepvclentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmcurrentstatuschangepvclentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmcurrentstatuschangepvclentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmcurrentstatuschangepvclentry.EntityData.Children = make(map[string]types.YChild)
    atmcurrentstatuschangepvclentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmcurrentstatuschangepvclentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atmcurrentstatuschangepvclentry.Ifindex}
    atmcurrentstatuschangepvclentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", atmcurrentstatuschangepvclentry.Atmvclvpi}
    atmcurrentstatuschangepvclentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", atmcurrentstatuschangepvclentry.Atmvclvci}
    atmcurrentstatuschangepvclentry.EntityData.Leafs["atmPVclStatusTransition"] = types.YLeaf{"Atmpvclstatustransition", atmcurrentstatuschangepvclentry.Atmpvclstatustransition}
    atmcurrentstatuschangepvclentry.EntityData.Leafs["atmPVclStatusChangeStart"] = types.YLeaf{"Atmpvclstatuschangestart", atmcurrentstatuschangepvclentry.Atmpvclstatuschangestart}
    atmcurrentstatuschangepvclentry.EntityData.Leafs["atmPVclStatusChangeEnd"] = types.YLeaf{"Atmpvclstatuschangeend", atmcurrentstatuschangepvclentry.Atmpvclstatuschangeend}
    return &(atmcurrentstatuschangepvclentry.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry.
    Atmstatuschangepvclrangeentry []CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry
}

func (atmstatuschangepvclrangetable *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable) GetEntityData() *types.CommonEntityData {
    atmstatuschangepvclrangetable.EntityData.YFilter = atmstatuschangepvclrangetable.YFilter
    atmstatuschangepvclrangetable.EntityData.YangName = "atmStatusChangePVclRangeTable"
    atmstatuschangepvclrangetable.EntityData.BundleName = "cisco_ios_xe"
    atmstatuschangepvclrangetable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    atmstatuschangepvclrangetable.EntityData.SegmentPath = "atmStatusChangePVclRangeTable"
    atmstatuschangepvclrangetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmstatuschangepvclrangetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmstatuschangepvclrangetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmstatuschangepvclrangetable.EntityData.Children = make(map[string]types.YChild)
    atmstatuschangepvclrangetable.EntityData.Children["atmStatusChangePVclRangeEntry"] = types.YChild{"Atmstatuschangepvclrangeentry", nil}
    for i := range atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry {
        atmstatuschangepvclrangetable.EntityData.Children[types.GetSegmentPath(&atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry[i])] = types.YChild{"Atmstatuschangepvclrangeentry", &atmstatuschangepvclrangetable.Atmstatuschangepvclrangeentry[i]}
    }
    atmstatuschangepvclrangetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmstatuschangepvclrangetable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_Atmvcltable_Atmvclentry_Atmvclvpi
    Atmvclvpi interface{}

    // This attribute is a key. Index to represent different ranges, the first
    // range is  indexed by value 0, the second by 1 and so on... The type is
    // interface{} with range: 0..4294967295.
    Rangeindex interface{}

    // The first PVCL in a range of PVcls for which the  atmVclOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    Atmpvcllowerrangevalue interface{}

    // The last PVCL in a range of PVcls for which the  atmOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    Atmpvclhigherrangevalue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclrangestatuschangestart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    Atmpvclrangestatuschangeend interface{}
}

func (atmstatuschangepvclrangeentry *CISCOIETFATM2PVCTRAPMIBEXTN_Atmstatuschangepvclrangetable_Atmstatuschangepvclrangeentry) GetEntityData() *types.CommonEntityData {
    atmstatuschangepvclrangeentry.EntityData.YFilter = atmstatuschangepvclrangeentry.YFilter
    atmstatuschangepvclrangeentry.EntityData.YangName = "atmStatusChangePVclRangeEntry"
    atmstatuschangepvclrangeentry.EntityData.BundleName = "cisco_ios_xe"
    atmstatuschangepvclrangeentry.EntityData.ParentYangName = "atmStatusChangePVclRangeTable"
    atmstatuschangepvclrangeentry.EntityData.SegmentPath = "atmStatusChangePVclRangeEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Atmvclvpi) + "']" + "[rangeIndex='" + fmt.Sprintf("%v", atmstatuschangepvclrangeentry.Rangeindex) + "']"
    atmstatuschangepvclrangeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmstatuschangepvclrangeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmstatuschangepvclrangeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmstatuschangepvclrangeentry.EntityData.Children = make(map[string]types.YChild)
    atmstatuschangepvclrangeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmstatuschangepvclrangeentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atmstatuschangepvclrangeentry.Ifindex}
    atmstatuschangepvclrangeentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", atmstatuschangepvclrangeentry.Atmvclvpi}
    atmstatuschangepvclrangeentry.EntityData.Leafs["rangeIndex"] = types.YLeaf{"Rangeindex", atmstatuschangepvclrangeentry.Rangeindex}
    atmstatuschangepvclrangeentry.EntityData.Leafs["atmPVclLowerRangeValue"] = types.YLeaf{"Atmpvcllowerrangevalue", atmstatuschangepvclrangeentry.Atmpvcllowerrangevalue}
    atmstatuschangepvclrangeentry.EntityData.Leafs["atmPVclHigherRangeValue"] = types.YLeaf{"Atmpvclhigherrangevalue", atmstatuschangepvclrangeentry.Atmpvclhigherrangevalue}
    atmstatuschangepvclrangeentry.EntityData.Leafs["atmPVclRangeStatusChangeStart"] = types.YLeaf{"Atmpvclrangestatuschangestart", atmstatuschangepvclrangeentry.Atmpvclrangestatuschangestart}
    atmstatuschangepvclrangeentry.EntityData.Leafs["atmPVclRangeStatusChangeEnd"] = types.YLeaf{"Atmpvclrangestatuschangeend", atmstatuschangepvclrangeentry.Atmpvclrangestatuschangeend}
    return &(atmstatuschangepvclrangeentry.EntityData)
}

