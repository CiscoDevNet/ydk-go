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
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cmplsidobjects CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetFilter() yfilter.YFilter { return cISCOIETFMPLSIDSTD03MIB.YFilter }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) SetFilter(yf yfilter.YFilter) { cISCOIETFMPLSIDSTD03MIB.YFilter = yf }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetGoName(yname string) string {
    if yname == "cmplsIdObjects" { return "Cmplsidobjects" }
    return ""
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetSegmentPath() string {
    return "CISCO-IETF-MPLS-ID-STD-03-MIB:CISCO-IETF-MPLS-ID-STD-03-MIB"
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cmplsIdObjects" {
        return &cISCOIETFMPLSIDSTD03MIB.Cmplsidobjects
    }
    return nil
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cmplsIdObjects"] = &cISCOIETFMPLSIDSTD03MIB.Cmplsidobjects
    return children
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetYangName() string { return "CISCO-IETF-MPLS-ID-STD-03-MIB" }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) SetParent(parent types.Entity) { cISCOIETFMPLSIDSTD03MIB.parent = parent }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetParent() types.Entity { return cISCOIETFMPLSIDSTD03MIB.parent }

func (cISCOIETFMPLSIDSTD03MIB *CISCOIETFMPLSIDSTD03MIB) GetParentYangName() string { return "CISCO-IETF-MPLS-ID-STD-03-MIB" }

// CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects
type CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object allows the administrator to assign a unique operator identifier
    // also called MPLS-TP Global_ID. The type is string with length: 4.
    Cmplsglobalid interface{}

    // This object allows the operator or service provider to assign a unique
    // MPLS-TP ITU-T Carrier Code (ICC) to a network. The type is string with
    // length: 1..6.
    Cmplsicc interface{}

    // This object allows the operator or service provider to assign a unique
    // MPLS-TP Node_ID.  The Node_ID is assigned within the scope of the
    // Global_ID. The type is interface{} with range: 0..4294967295.
    Cmplsnodeid interface{}
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetFilter() yfilter.YFilter { return cmplsidobjects.YFilter }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) SetFilter(yf yfilter.YFilter) { cmplsidobjects.YFilter = yf }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetGoName(yname string) string {
    if yname == "cmplsGlobalId" { return "Cmplsglobalid" }
    if yname == "cmplsIcc" { return "Cmplsicc" }
    if yname == "cmplsNodeId" { return "Cmplsnodeid" }
    return ""
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetSegmentPath() string {
    return "cmplsIdObjects"
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmplsGlobalId"] = cmplsidobjects.Cmplsglobalid
    leafs["cmplsIcc"] = cmplsidobjects.Cmplsicc
    leafs["cmplsNodeId"] = cmplsidobjects.Cmplsnodeid
    return leafs
}

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetBundleName() string { return "cisco_ios_xe" }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetYangName() string { return "cmplsIdObjects" }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) SetParent(parent types.Entity) { cmplsidobjects.parent = parent }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetParent() types.Entity { return cmplsidobjects.parent }

func (cmplsidobjects *CISCOIETFMPLSIDSTD03MIB_Cmplsidobjects) GetParentYangName() string { return "CISCO-IETF-MPLS-ID-STD-03-MIB" }

