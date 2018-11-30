// This MIB module defines the namespace organization for the
// CableLabs enterprise OID registry.
// 
// Copyright 1999-2012 Cable Television Laboratories, Inc.
// All rights reserved.
package clab_def_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package clab_def_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CLAB-DEF-MIB CLAB-DEF-MIB}", reflect.TypeOf(CLABDEFMIB{}))
    ydk.RegisterEntity("CLAB-DEF-MIB:CLAB-DEF-MIB", reflect.TypeOf(CLABDEFMIB{}))
}

// CLABDEFMIB
type CLABDEFMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    ClabSecCertObject CLABDEFMIB_ClabSecCertObject
}

func (cLABDEFMIB *CLABDEFMIB) GetEntityData() *types.CommonEntityData {
    cLABDEFMIB.EntityData.YFilter = cLABDEFMIB.YFilter
    cLABDEFMIB.EntityData.YangName = "CLAB-DEF-MIB"
    cLABDEFMIB.EntityData.BundleName = "cisco_ios_xe"
    cLABDEFMIB.EntityData.ParentYangName = "CLAB-DEF-MIB"
    cLABDEFMIB.EntityData.SegmentPath = "CLAB-DEF-MIB:CLAB-DEF-MIB"
    cLABDEFMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cLABDEFMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cLABDEFMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cLABDEFMIB.EntityData.Children = types.NewOrderedMap()
    cLABDEFMIB.EntityData.Children.Append("clabSecCertObject", types.YChild{"ClabSecCertObject", &cLABDEFMIB.ClabSecCertObject})
    cLABDEFMIB.EntityData.Leafs = types.NewOrderedMap()

    cLABDEFMIB.EntityData.YListKeys = []string {}

    return &(cLABDEFMIB.EntityData)
}

// CLABDEFMIB_ClabSecCertObject
type CLABDEFMIB_ClabSecCertObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The X509 DER-encoded CableLabs Service Provider Root CA Certificate. The
    // type is string with length: 0..4096.
    ClabSrvcPrvdrRootCACert interface{}

    // The X509 DER-encoded CableLabs CVC Root CA Certificate. The type is string
    // with length: 0..4096.
    ClabCVCRootCACert interface{}

    // The X509 DER-encoded CableLabs CVC CA Certificate. The type is string with
    // length: 0..4096.
    ClabCVCCACert interface{}

    // The X509 DER-encoded Manufacturer CVC Certificate. The type is string with
    // length: 0..4096.
    ClabMfgCVCCert interface{}

    // The X509 DER-encoded Manufacturer CA Certificate. The type is string with
    // length: 0..4096.
    ClabMfgCACert interface{}
}

func (clabSecCertObject *CLABDEFMIB_ClabSecCertObject) GetEntityData() *types.CommonEntityData {
    clabSecCertObject.EntityData.YFilter = clabSecCertObject.YFilter
    clabSecCertObject.EntityData.YangName = "clabSecCertObject"
    clabSecCertObject.EntityData.BundleName = "cisco_ios_xe"
    clabSecCertObject.EntityData.ParentYangName = "CLAB-DEF-MIB"
    clabSecCertObject.EntityData.SegmentPath = "clabSecCertObject"
    clabSecCertObject.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clabSecCertObject.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clabSecCertObject.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clabSecCertObject.EntityData.Children = types.NewOrderedMap()
    clabSecCertObject.EntityData.Leafs = types.NewOrderedMap()
    clabSecCertObject.EntityData.Leafs.Append("clabSrvcPrvdrRootCACert", types.YLeaf{"ClabSrvcPrvdrRootCACert", clabSecCertObject.ClabSrvcPrvdrRootCACert})
    clabSecCertObject.EntityData.Leafs.Append("clabCVCRootCACert", types.YLeaf{"ClabCVCRootCACert", clabSecCertObject.ClabCVCRootCACert})
    clabSecCertObject.EntityData.Leafs.Append("clabCVCCACert", types.YLeaf{"ClabCVCCACert", clabSecCertObject.ClabCVCCACert})
    clabSecCertObject.EntityData.Leafs.Append("clabMfgCVCCert", types.YLeaf{"ClabMfgCVCCert", clabSecCertObject.ClabMfgCVCCert})
    clabSecCertObject.EntityData.Leafs.Append("clabMfgCACert", types.YLeaf{"ClabMfgCACert", clabSecCertObject.ClabMfgCACert})

    clabSecCertObject.EntityData.YListKeys = []string {}

    return &(clabSecCertObject.EntityData)
}

