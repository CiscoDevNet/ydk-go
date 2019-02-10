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
    AtmCurrentlyFailingPVclTable CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable
}

func (cISCOIETFATM2PVCTRAPMIB *CISCOIETFATM2PVCTRAPMIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFATM2PVCTRAPMIB.EntityData.YFilter = cISCOIETFATM2PVCTRAPMIB.YFilter
    cISCOIETFATM2PVCTRAPMIB.EntityData.YangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFATM2PVCTRAPMIB.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.SegmentPath = "CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB"
    cISCOIETFATM2PVCTRAPMIB.EntityData.AbsolutePath = cISCOIETFATM2PVCTRAPMIB.EntityData.SegmentPath
    cISCOIETFATM2PVCTRAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFATM2PVCTRAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFATM2PVCTRAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFATM2PVCTRAPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFATM2PVCTRAPMIB.EntityData.Children.Append("atmCurrentlyFailingPVclTable", types.YChild{"AtmCurrentlyFailingPVclTable", &cISCOIETFATM2PVCTRAPMIB.AtmCurrentlyFailingPVclTable})
    cISCOIETFATM2PVCTRAPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFATM2PVCTRAPMIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFATM2PVCTRAPMIB.EntityData)
}

// CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and an atmVclOperStatus with a value
// other than `up'.
type CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a VCL for which the atmVclRowStatus is
    // `active', the atmVclConnKind is `pvc', and the atmVclOperStatus is other
    // than `up'. The type is slice of
    // CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable_AtmCurrentlyFailingPVclEntry.
    AtmCurrentlyFailingPVclEntry []*CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable_AtmCurrentlyFailingPVclEntry
}

func (atmCurrentlyFailingPVclTable *CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable) GetEntityData() *types.CommonEntityData {
    atmCurrentlyFailingPVclTable.EntityData.YFilter = atmCurrentlyFailingPVclTable.YFilter
    atmCurrentlyFailingPVclTable.EntityData.YangName = "atmCurrentlyFailingPVclTable"
    atmCurrentlyFailingPVclTable.EntityData.BundleName = "cisco_ios_xe"
    atmCurrentlyFailingPVclTable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB"
    atmCurrentlyFailingPVclTable.EntityData.SegmentPath = "atmCurrentlyFailingPVclTable"
    atmCurrentlyFailingPVclTable.EntityData.AbsolutePath = "CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB/" + atmCurrentlyFailingPVclTable.EntityData.SegmentPath
    atmCurrentlyFailingPVclTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmCurrentlyFailingPVclTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmCurrentlyFailingPVclTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmCurrentlyFailingPVclTable.EntityData.Children = types.NewOrderedMap()
    atmCurrentlyFailingPVclTable.EntityData.Children.Append("atmCurrentlyFailingPVclEntry", types.YChild{"AtmCurrentlyFailingPVclEntry", nil})
    for i := range atmCurrentlyFailingPVclTable.AtmCurrentlyFailingPVclEntry {
        atmCurrentlyFailingPVclTable.EntityData.Children.Append(types.GetSegmentPath(atmCurrentlyFailingPVclTable.AtmCurrentlyFailingPVclEntry[i]), types.YChild{"AtmCurrentlyFailingPVclEntry", atmCurrentlyFailingPVclTable.AtmCurrentlyFailingPVclEntry[i]})
    }
    atmCurrentlyFailingPVclTable.EntityData.Leafs = types.NewOrderedMap()

    atmCurrentlyFailingPVclTable.EntityData.YListKeys = []string {}

    return &(atmCurrentlyFailingPVclTable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable_AtmCurrentlyFailingPVclEntry
// Each entry in this table represents a VCL for which
// the atmVclRowStatus is `active', the atmVclConnKind is
// `pvc', and the atmVclOperStatus is other than `up'.
type CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable_AtmCurrentlyFailingPVclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // The time at which this PVCL began to fail. The type is interface{} with
    // range: 0..4294967295.
    AtmCurrentlyFailingPVclTimeStamp interface{}

    // The time at which this PVCL began to fail during the PVC Notification
    // interval. The type is interface{} with range: 0..4294967295.
    AtmPreviouslyFailedPVclTimeStamp interface{}
}

func (atmCurrentlyFailingPVclEntry *CISCOIETFATM2PVCTRAPMIB_AtmCurrentlyFailingPVclTable_AtmCurrentlyFailingPVclEntry) GetEntityData() *types.CommonEntityData {
    atmCurrentlyFailingPVclEntry.EntityData.YFilter = atmCurrentlyFailingPVclEntry.YFilter
    atmCurrentlyFailingPVclEntry.EntityData.YangName = "atmCurrentlyFailingPVclEntry"
    atmCurrentlyFailingPVclEntry.EntityData.BundleName = "cisco_ios_xe"
    atmCurrentlyFailingPVclEntry.EntityData.ParentYangName = "atmCurrentlyFailingPVclTable"
    atmCurrentlyFailingPVclEntry.EntityData.SegmentPath = "atmCurrentlyFailingPVclEntry" + types.AddKeyToken(atmCurrentlyFailingPVclEntry.IfIndex, "ifIndex") + types.AddKeyToken(atmCurrentlyFailingPVclEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(atmCurrentlyFailingPVclEntry.AtmVclVci, "atmVclVci")
    atmCurrentlyFailingPVclEntry.EntityData.AbsolutePath = "CISCO-IETF-ATM2-PVCTRAP-MIB:CISCO-IETF-ATM2-PVCTRAP-MIB/atmCurrentlyFailingPVclTable/" + atmCurrentlyFailingPVclEntry.EntityData.SegmentPath
    atmCurrentlyFailingPVclEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmCurrentlyFailingPVclEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmCurrentlyFailingPVclEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmCurrentlyFailingPVclEntry.EntityData.Children = types.NewOrderedMap()
    atmCurrentlyFailingPVclEntry.EntityData.Leafs = types.NewOrderedMap()
    atmCurrentlyFailingPVclEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmCurrentlyFailingPVclEntry.IfIndex})
    atmCurrentlyFailingPVclEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", atmCurrentlyFailingPVclEntry.AtmVclVpi})
    atmCurrentlyFailingPVclEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", atmCurrentlyFailingPVclEntry.AtmVclVci})
    atmCurrentlyFailingPVclEntry.EntityData.Leafs.Append("atmCurrentlyFailingPVclTimeStamp", types.YLeaf{"AtmCurrentlyFailingPVclTimeStamp", atmCurrentlyFailingPVclEntry.AtmCurrentlyFailingPVclTimeStamp})
    atmCurrentlyFailingPVclEntry.EntityData.Leafs.Append("atmPreviouslyFailedPVclTimeStamp", types.YLeaf{"AtmPreviouslyFailedPVclTimeStamp", atmCurrentlyFailingPVclEntry.AtmPreviouslyFailedPVclTimeStamp})

    atmCurrentlyFailingPVclEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(atmCurrentlyFailingPVclEntry.EntityData)
}

