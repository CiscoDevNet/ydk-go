// Copyright (c) 2012 IETF Trust and the persons identified
// as the document authors.  All rights reserved.
// 
// This MIB module contains generic object definitions for
// MPLS Traffic Engineering in transport networks. This module is a
// cisco-ized version of the IETF draft:
// draft-ietf-mpls-tp-te-mib-03.
package cisco_ietf_mpls_id_std_03_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_mpls_id_std_03_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IETF-MPLS-ID-STD-03-MIB CISCO-IETF-MPLS-ID-STD-03-MIB}", reflect.TypeOf(CISCOIETFMPLSIDSTD03MIB{}))
    ydk.RegisterEntity("CISCO-IETF-MPLS-ID-STD-03-MIB:CISCO-IETF-MPLS-ID-STD-03-MIB", reflect.TypeOf(CISCOIETFMPLSIDSTD03MIB{}))
}

// CISCOIETFMPLSIDSTD03MIB
type CISCOIETFMPLSIDSTD03MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CmplsIdObjects CISCOIETFMPLSIDSTD03MIB_CmplsIdObjects
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetEntityData() *types.CommonEntityData {
    cISCOIETFMPLSIDSTD03MIB.EntityData.YFilter = cISCOIETFMPLSIDSTD03MIB.YFilter
    cISCOIETFMPLSIDSTD03MIB.EntityData.YangName = "CISCO-IETF-MPLS-ID-STD-03-MIB"
    cISCOIETFMPLSIDSTD03MIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIETFMPLSIDSTD03MIB.EntityData.ParentYangName = "CISCO-IETF-MPLS-ID-STD-03-MIB"
    cISCOIETFMPLSIDSTD03MIB.EntityData.SegmentPath = "CISCO-IETF-MPLS-ID-STD-03-MIB:CISCO-IETF-MPLS-ID-STD-03-MIB"
    cISCOIETFMPLSIDSTD03MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIETFMPLSIDSTD03MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIETFMPLSIDSTD03MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIETFMPLSIDSTD03MIB.EntityData.Children = types.NewOrderedMap()
    cISCOIETFMPLSIDSTD03MIB.EntityData.Children.Append("cmplsIdObjects", types.YChild{"CmplsIdObjects", &cISCOIETFMPLSIDSTD03MIB.CmplsIdObjects})
    cISCOIETFMPLSIDSTD03MIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIETFMPLSIDSTD03MIB.EntityData.YListKeys = []string {}

    return &(cISCOIETFMPLSIDSTD03MIB.EntityData)
}

// CISCOIETFMPLSIDSTD03MIB_CmplsIdObjects
type CISCOIETFMPLSIDSTD03MIB_CmplsIdObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object allows the administrator to assign a unique operator identifier
    // also called MPLS-TP Global_ID. The type is string with length: 4.
    CmplsGlobalId interface{}

    // This object allows the operator or service provider to assign a unique
    // MPLS-TP ITU-T Carrier Code (ICC) to a network. The type is string with
    // length: 1..6.
    CmplsIcc interface{}

    // This object allows the operator or service provider to assign a unique
    // MPLS-TP Node_ID.  The Node_ID is assigned within the scope of the
    // Global_ID. The type is interface{} with range: 0..4294967295.
    CmplsNodeId interface{}
}

func (cmplsIdObjects *CISCOIETFMPLSIDSTD03MIB_CmplsIdObjects) GetEntityData() *types.CommonEntityData {
    cmplsIdObjects.EntityData.YFilter = cmplsIdObjects.YFilter
    cmplsIdObjects.EntityData.YangName = "cmplsIdObjects"
    cmplsIdObjects.EntityData.BundleName = "cisco_ios_xe"
    cmplsIdObjects.EntityData.ParentYangName = "CISCO-IETF-MPLS-ID-STD-03-MIB"
    cmplsIdObjects.EntityData.SegmentPath = "cmplsIdObjects"
    cmplsIdObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cmplsIdObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cmplsIdObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cmplsIdObjects.EntityData.Children = types.NewOrderedMap()
    cmplsIdObjects.EntityData.Leafs = types.NewOrderedMap()
    cmplsIdObjects.EntityData.Leafs.Append("cmplsGlobalId", types.YLeaf{"CmplsGlobalId", cmplsIdObjects.CmplsGlobalId})
    cmplsIdObjects.EntityData.Leafs.Append("cmplsIcc", types.YLeaf{"CmplsIcc", cmplsIdObjects.CmplsIcc})
    cmplsIdObjects.EntityData.Leafs.Append("cmplsNodeId", types.YLeaf{"CmplsNodeId", cmplsIdObjects.CmplsNodeId})

    cmplsIdObjects.EntityData.YListKeys = []string {}

    return &(cmplsIdObjects.EntityData)
}

