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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A table indicating all VCLs for which there is an active row in the
    // atmVclTable having an atmVclConnKind value of `pvc' and an atmVclOperStatus
    // with a value other than `up'.
    Atmcurrentlyfailingpvcltable CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFATM2PVCTRAPMIB.EntityData.YFilter = cISCOIETFATM2PVCTRAPMIB.YFilter
    cISCOIETFATM2PVCTRAPMIB.EntityData.YangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFATM2PVCTRAPMIB.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.SegmentPath = "CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFATM2PVCTRAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFATM2PVCTRAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFATM2PVCTRAPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIETFATM2PVCTRAPMIB.EntityData.Children["atmCurrentlyFailingPVclTable"] = types.YChild{"Atmcurrentlyfailingpvcltable", &cISCOIETFATM2PVCTRAPMIB.Atmcurrentlyfailingpvcltable}
    cISCOIETFATM2PVCTRAPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIETFATM2PVCTRAPMIB.EntityData)
}

// CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and an atmVclOperStatus with a value
// other than `up'.
type CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a VCL for which the atmVclRowStatus is
    // `active', the atmVclConnKind is `pvc', and the atmVclOperStatus is other
    // than `up'. The type is slice of
    // CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry.
    Atmcurrentlyfailingpvclentry []CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry
}

func (atmcurrentlyfailingpvcltable *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable) GetEntityData() *types.CommonEntityData {
    atmcurrentlyfailingpvcltable.EntityData.YFilter = atmcurrentlyfailingpvcltable.YFilter
    atmcurrentlyfailingpvcltable.EntityData.YangName = "atmCurrentlyFailingPVclTable"
    atmcurrentlyfailingpvcltable.EntityData.BundleName = "cisco_ios_xe"
    atmcurrentlyfailingpvcltable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    atmcurrentlyfailingpvcltable.EntityData.SegmentPath = "atmCurrentlyFailingPVclTable"
    atmcurrentlyfailingpvcltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmcurrentlyfailingpvcltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmcurrentlyfailingpvcltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmcurrentlyfailingpvcltable.EntityData.Children = make(map[string]types.YChild)
    atmcurrentlyfailingpvcltable.EntityData.Children["atmCurrentlyFailingPVclEntry"] = types.YChild{"Atmcurrentlyfailingpvclentry", nil}
    for i := range atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry {
        atmcurrentlyfailingpvcltable.EntityData.Children[types.GetSegmentPath(&atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry[i])] = types.YChild{"Atmcurrentlyfailingpvclentry", &atmcurrentlyfailingpvcltable.Atmcurrentlyfailingpvclentry[i]}
    }
    atmcurrentlyfailingpvcltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(atmcurrentlyfailingpvcltable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry
// Each entry in this table represents a VCL for which
// the atmVclRowStatus is `active', the atmVclConnKind is
// `pvc', and the atmVclOperStatus is other than `up'.
type CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry struct {
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

    // The time at which this PVCL began to fail. The type is interface{} with
    // range: 0..4294967295.
    Atmcurrentlyfailingpvcltimestamp interface{}

    // The time at which this PVCL began to fail during the PVC Notification
    // interval. The type is interface{} with range: 0..4294967295.
    Atmpreviouslyfailedpvcltimestamp interface{}
}

func (atmcurrentlyfailingpvclentry *CISCOIETFATM2PVCTRAPMIB_Atmcurrentlyfailingpvcltable_Atmcurrentlyfailingpvclentry) GetEntityData() *types.CommonEntityData {
    atmcurrentlyfailingpvclentry.EntityData.YFilter = atmcurrentlyfailingpvclentry.YFilter
    atmcurrentlyfailingpvclentry.EntityData.YangName = "atmCurrentlyFailingPVclEntry"
    atmcurrentlyfailingpvclentry.EntityData.BundleName = "cisco_ios_xe"
    atmcurrentlyfailingpvclentry.EntityData.ParentYangName = "atmCurrentlyFailingPVclTable"
    atmcurrentlyfailingpvclentry.EntityData.SegmentPath = "atmCurrentlyFailingPVclEntry" + "[ifIndex='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Ifindex) + "']" + "[atmVclVpi='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Atmvclvpi) + "']" + "[atmVclVci='" + fmt.Sprintf("%v", atmcurrentlyfailingpvclentry.Atmvclvci) + "']"
    atmcurrentlyfailingpvclentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmcurrentlyfailingpvclentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmcurrentlyfailingpvclentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmcurrentlyfailingpvclentry.EntityData.Children = make(map[string]types.YChild)
    atmcurrentlyfailingpvclentry.EntityData.Leafs = make(map[string]types.YLeaf)
    atmcurrentlyfailingpvclentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", atmcurrentlyfailingpvclentry.Ifindex}
    atmcurrentlyfailingpvclentry.EntityData.Leafs["atmVclVpi"] = types.YLeaf{"Atmvclvpi", atmcurrentlyfailingpvclentry.Atmvclvpi}
    atmcurrentlyfailingpvclentry.EntityData.Leafs["atmVclVci"] = types.YLeaf{"Atmvclvci", atmcurrentlyfailingpvclentry.Atmvclvci}
    atmcurrentlyfailingpvclentry.EntityData.Leafs["atmCurrentlyFailingPVclTimeStamp"] = types.YLeaf{"Atmcurrentlyfailingpvcltimestamp", atmcurrentlyfailingpvclentry.Atmcurrentlyfailingpvcltimestamp}
    atmcurrentlyfailingpvclentry.EntityData.Leafs["atmPreviouslyFailedPVclTimeStamp"] = types.YLeaf{"Atmpreviouslyfailedpvcltimestamp", atmcurrentlyfailingpvclentry.Atmpreviouslyfailedpvcltimestamp}
    return &(atmcurrentlyfailingpvclentry.EntityData)
}

