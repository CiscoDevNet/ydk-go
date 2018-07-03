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
    AtmCurrentStatusChangePVclTable CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable

    // A table indicating more than one VCLs in a consecutive  range and for each
    // VCL there is an active row in the  atmVclTable having an atmVclConnKind
    // value of `pvc' and atmVclOperStatus to have changed in the same direction
    // in the last PVC notification interval.
    AtmStatusChangePVclRangeTable CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable
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

    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children = types.NewOrderedMap()
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children.Append("atmCurrentStatusChangePVclTable", types.YChild{"AtmCurrentStatusChangePVclTable", &cISCOIETFATM2PVCTRAPMIBEXTN.AtmCurrentStatusChangePVclTable})
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Children.Append("atmStatusChangePVclRangeTable", types.YChild{"AtmStatusChangePVclRangeTable", &cISCOIETFATM2PVCTRAPMIBEXTN.AtmStatusChangePVclRangeTable})
    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFATM2PVCTRAPMIBEXTN.EntityData.YListKeys = []string {}

    return &(cISCOIETFATM2PVCTRAPMIBEXTN.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable
// A table indicating all VCLs for which there is an
// active row in the atmVclTable having an atmVclConnKind
// value of `pvc' and atmVclOperStatus to have changed in the
// last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in the table represents a VCL for which there is an active row
    // in the atmVclTable having an atmVclConnKind value of `pvc' and
    // atmVclOperStatus to have changed in the last PVC notification interval. The
    // type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable_AtmCurrentStatusChangePVclEntry.
    AtmCurrentStatusChangePVclEntry []*CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable_AtmCurrentStatusChangePVclEntry
}

func (atmCurrentStatusChangePVclTable *CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable) GetEntityData() *types.CommonEntityData {
    atmCurrentStatusChangePVclTable.EntityData.YFilter = atmCurrentStatusChangePVclTable.YFilter
    atmCurrentStatusChangePVclTable.EntityData.YangName = "atmCurrentStatusChangePVclTable"
    atmCurrentStatusChangePVclTable.EntityData.BundleName = "cisco_ios_xe"
    atmCurrentStatusChangePVclTable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    atmCurrentStatusChangePVclTable.EntityData.SegmentPath = "atmCurrentStatusChangePVclTable"
    atmCurrentStatusChangePVclTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmCurrentStatusChangePVclTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmCurrentStatusChangePVclTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmCurrentStatusChangePVclTable.EntityData.Children = types.NewOrderedMap()
    atmCurrentStatusChangePVclTable.EntityData.Children.Append("atmCurrentStatusChangePVclEntry", types.YChild{"AtmCurrentStatusChangePVclEntry", nil})
    for i := range atmCurrentStatusChangePVclTable.AtmCurrentStatusChangePVclEntry {
        atmCurrentStatusChangePVclTable.EntityData.Children.Append(types.GetSegmentPath(atmCurrentStatusChangePVclTable.AtmCurrentStatusChangePVclEntry[i]), types.YChild{"AtmCurrentStatusChangePVclEntry", atmCurrentStatusChangePVclTable.AtmCurrentStatusChangePVclEntry[i]})
    }
    atmCurrentStatusChangePVclTable.EntityData.Leafs = types.NewOrderedMap()

    atmCurrentStatusChangePVclTable.EntityData.YListKeys = []string {}

    return &(atmCurrentStatusChangePVclTable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable_AtmCurrentStatusChangePVclEntry
// Each entry in the table represents a VCL for which
// there is an active row in the atmVclTable having an
// atmVclConnKind value of `pvc' and atmVclOperStatus
// to have changed in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable_AtmCurrentStatusChangePVclEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. The type is string with range: 0..65535. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVci
    AtmVclVci interface{}

    // The number of state transitions that has happened  on this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    AtmPVclStatusTransition interface{}

    // The time stamp at which this PVCL changed state for the first time in  the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    AtmPVclStatusChangeStart interface{}

    // The time stamp of the last state change of this PVCL in the last
    // notification interval. The type is interface{} with range: 0..4294967295.
    AtmPVclStatusChangeEnd interface{}
}

func (atmCurrentStatusChangePVclEntry *CISCOIETFATM2PVCTRAPMIBEXTN_AtmCurrentStatusChangePVclTable_AtmCurrentStatusChangePVclEntry) GetEntityData() *types.CommonEntityData {
    atmCurrentStatusChangePVclEntry.EntityData.YFilter = atmCurrentStatusChangePVclEntry.YFilter
    atmCurrentStatusChangePVclEntry.EntityData.YangName = "atmCurrentStatusChangePVclEntry"
    atmCurrentStatusChangePVclEntry.EntityData.BundleName = "cisco_ios_xe"
    atmCurrentStatusChangePVclEntry.EntityData.ParentYangName = "atmCurrentStatusChangePVclTable"
    atmCurrentStatusChangePVclEntry.EntityData.SegmentPath = "atmCurrentStatusChangePVclEntry" + types.AddKeyToken(atmCurrentStatusChangePVclEntry.IfIndex, "ifIndex") + types.AddKeyToken(atmCurrentStatusChangePVclEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(atmCurrentStatusChangePVclEntry.AtmVclVci, "atmVclVci")
    atmCurrentStatusChangePVclEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmCurrentStatusChangePVclEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmCurrentStatusChangePVclEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmCurrentStatusChangePVclEntry.EntityData.Children = types.NewOrderedMap()
    atmCurrentStatusChangePVclEntry.EntityData.Leafs = types.NewOrderedMap()
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmCurrentStatusChangePVclEntry.IfIndex})
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", atmCurrentStatusChangePVclEntry.AtmVclVpi})
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("atmVclVci", types.YLeaf{"AtmVclVci", atmCurrentStatusChangePVclEntry.AtmVclVci})
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("atmPVclStatusTransition", types.YLeaf{"AtmPVclStatusTransition", atmCurrentStatusChangePVclEntry.AtmPVclStatusTransition})
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("atmPVclStatusChangeStart", types.YLeaf{"AtmPVclStatusChangeStart", atmCurrentStatusChangePVclEntry.AtmPVclStatusChangeStart})
    atmCurrentStatusChangePVclEntry.EntityData.Leafs.Append("atmPVclStatusChangeEnd", types.YLeaf{"AtmPVclStatusChangeEnd", atmCurrentStatusChangePVclEntry.AtmPVclStatusChangeEnd})

    atmCurrentStatusChangePVclEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "AtmVclVci"}

    return &(atmCurrentStatusChangePVclEntry.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable
// A table indicating more than one VCLs in a consecutive 
// range and for each VCL there is an active row in the 
// atmVclTable having an atmVclConnKind value of `pvc'
// and atmVclOperStatus to have changed in the same
// direction in the last PVC notification interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry in this table represents a range of VCLs and  for each VCL there
    // is an active row in the atmVclTable having an atmVclConnKind value of 'pvc'
    // and atmVclOperStatus to have changed in the same direction in the last
    // notification  interval. The type is slice of
    // CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable_AtmStatusChangePVclRangeEntry.
    AtmStatusChangePVclRangeEntry []*CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable_AtmStatusChangePVclRangeEntry
}

func (atmStatusChangePVclRangeTable *CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable) GetEntityData() *types.CommonEntityData {
    atmStatusChangePVclRangeTable.EntityData.YFilter = atmStatusChangePVclRangeTable.YFilter
    atmStatusChangePVclRangeTable.EntityData.YangName = "atmStatusChangePVclRangeTable"
    atmStatusChangePVclRangeTable.EntityData.BundleName = "cisco_ios_xe"
    atmStatusChangePVclRangeTable.EntityData.ParentYangName = "CISCO-IETF-ATM2-PVCTRAP-MIB-EXTN"
    atmStatusChangePVclRangeTable.EntityData.SegmentPath = "atmStatusChangePVclRangeTable"
    atmStatusChangePVclRangeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmStatusChangePVclRangeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmStatusChangePVclRangeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmStatusChangePVclRangeTable.EntityData.Children = types.NewOrderedMap()
    atmStatusChangePVclRangeTable.EntityData.Children.Append("atmStatusChangePVclRangeEntry", types.YChild{"AtmStatusChangePVclRangeEntry", nil})
    for i := range atmStatusChangePVclRangeTable.AtmStatusChangePVclRangeEntry {
        atmStatusChangePVclRangeTable.EntityData.Children.Append(types.GetSegmentPath(atmStatusChangePVclRangeTable.AtmStatusChangePVclRangeEntry[i]), types.YChild{"AtmStatusChangePVclRangeEntry", atmStatusChangePVclRangeTable.AtmStatusChangePVclRangeEntry[i]})
    }
    atmStatusChangePVclRangeTable.EntityData.Leafs = types.NewOrderedMap()

    atmStatusChangePVclRangeTable.EntityData.YListKeys = []string {}

    return &(atmStatusChangePVclRangeTable.EntityData)
}

// CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable_AtmStatusChangePVclRangeEntry
// Each entry in this table represents a range of VCLs and 
// for each VCL there is an active row in the atmVclTable having
// an atmVclConnKind value of 'pvc' and atmVclOperStatus to have
// changed in the same direction in the last notification 
// interval.
type CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable_AtmStatusChangePVclRangeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with range: 0..4095. Refers to
    // atm_mib.ATMMIB_AtmVclTable_AtmVclEntry_AtmVclVpi
    AtmVclVpi interface{}

    // This attribute is a key. Index to represent different ranges, the first
    // range is  indexed by value 0, the second by 1 and so on... The type is
    // interface{} with range: 0..4294967295.
    RangeIndex interface{}

    // The first PVCL in a range of PVcls for which the  atmVclOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    AtmPVclLowerRangeValue interface{}

    // The last PVCL in a range of PVcls for which the  atmOperStatus to have
    // changed in the last  notification interval. The type is interface{} with
    // range: 0..65536.
    AtmPVclHigherRangeValue interface{}

    // The time stamp at which the first PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    AtmPVclRangeStatusChangeStart interface{}

    // The time stamp at which the last PVCL in the range changed state in the
    // last notification interval. The type is interface{} with range:
    // 0..4294967295.
    AtmPVclRangeStatusChangeEnd interface{}
}

func (atmStatusChangePVclRangeEntry *CISCOIETFATM2PVCTRAPMIBEXTN_AtmStatusChangePVclRangeTable_AtmStatusChangePVclRangeEntry) GetEntityData() *types.CommonEntityData {
    atmStatusChangePVclRangeEntry.EntityData.YFilter = atmStatusChangePVclRangeEntry.YFilter
    atmStatusChangePVclRangeEntry.EntityData.YangName = "atmStatusChangePVclRangeEntry"
    atmStatusChangePVclRangeEntry.EntityData.BundleName = "cisco_ios_xe"
    atmStatusChangePVclRangeEntry.EntityData.ParentYangName = "atmStatusChangePVclRangeTable"
    atmStatusChangePVclRangeEntry.EntityData.SegmentPath = "atmStatusChangePVclRangeEntry" + types.AddKeyToken(atmStatusChangePVclRangeEntry.IfIndex, "ifIndex") + types.AddKeyToken(atmStatusChangePVclRangeEntry.AtmVclVpi, "atmVclVpi") + types.AddKeyToken(atmStatusChangePVclRangeEntry.RangeIndex, "rangeIndex")
    atmStatusChangePVclRangeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    atmStatusChangePVclRangeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    atmStatusChangePVclRangeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    atmStatusChangePVclRangeEntry.EntityData.Children = types.NewOrderedMap()
    atmStatusChangePVclRangeEntry.EntityData.Leafs = types.NewOrderedMap()
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", atmStatusChangePVclRangeEntry.IfIndex})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmVclVpi", types.YLeaf{"AtmVclVpi", atmStatusChangePVclRangeEntry.AtmVclVpi})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("rangeIndex", types.YLeaf{"RangeIndex", atmStatusChangePVclRangeEntry.RangeIndex})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmPVclLowerRangeValue", types.YLeaf{"AtmPVclLowerRangeValue", atmStatusChangePVclRangeEntry.AtmPVclLowerRangeValue})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmPVclHigherRangeValue", types.YLeaf{"AtmPVclHigherRangeValue", atmStatusChangePVclRangeEntry.AtmPVclHigherRangeValue})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmPVclRangeStatusChangeStart", types.YLeaf{"AtmPVclRangeStatusChangeStart", atmStatusChangePVclRangeEntry.AtmPVclRangeStatusChangeStart})
    atmStatusChangePVclRangeEntry.EntityData.Leafs.Append("atmPVclRangeStatusChangeEnd", types.YLeaf{"AtmPVclRangeStatusChangeEnd", atmStatusChangePVclRangeEntry.AtmPVclRangeStatusChangeEnd})

    atmStatusChangePVclRangeEntry.EntityData.YListKeys = []string {"IfIndex", "AtmVclVpi", "RangeIndex"}

    return &(atmStatusChangePVclRangeEntry.EntityData)
}

