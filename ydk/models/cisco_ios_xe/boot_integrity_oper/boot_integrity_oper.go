// This module contains a collection of YANG definitions
// for Cisco IOS XE boot integrity visibility.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package boot_integrity_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package boot_integrity_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-boot-integrity-oper boot-integrity-oper-data}", reflect.TypeOf(BootIntegrityOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data", reflect.TypeOf(BootIntegrityOperData{}))
}

// BootIntegrityOperData
// Enclosing container for the boot integrity 
// measurements of the system
type BootIntegrityOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of system boot integrity measurements for Boot,  Boot Loader, and OS
    // versions, hashes, and signatures as well as Platform Configuration
    // Registers (PCR) content. These measurements are captured utilizing Trusted
    // Anchor Module (TAM) services to communicate with system ACT2  (2nd Gen
    // Anti-Counterfeit Technology) hardware chip.
    BootIntegrity BootIntegrityOperData_BootIntegrity

    // List of system  certificate measurements for Cisco Root CA (CRCA), Cisco
    // Manufacturing CA (CMCA), and ACT2 RSA Secure Unique Device  Identity (SUDI)
    // CA PEM certifcates and SUDI generated signatures. These measurements are
    // captured utilizing Trusted Anchor Module (TAM) services to communicate with
    // system ACT2  (2nd Gen Anti-Counterfeit Technology) hardware chip.
    SudiCertificate BootIntegrityOperData_SudiCertificate
}

func (bootIntegrityOperData *BootIntegrityOperData) GetEntityData() *types.CommonEntityData {
    bootIntegrityOperData.EntityData.YFilter = bootIntegrityOperData.YFilter
    bootIntegrityOperData.EntityData.YangName = "boot-integrity-oper-data"
    bootIntegrityOperData.EntityData.BundleName = "cisco_ios_xe"
    bootIntegrityOperData.EntityData.ParentYangName = "Cisco-IOS-XE-boot-integrity-oper"
    bootIntegrityOperData.EntityData.SegmentPath = "Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data"
    bootIntegrityOperData.EntityData.AbsolutePath = bootIntegrityOperData.EntityData.SegmentPath
    bootIntegrityOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bootIntegrityOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bootIntegrityOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bootIntegrityOperData.EntityData.Children = types.NewOrderedMap()
    bootIntegrityOperData.EntityData.Children.Append("boot-integrity", types.YChild{"BootIntegrity", &bootIntegrityOperData.BootIntegrity})
    bootIntegrityOperData.EntityData.Children.Append("sudi-certificate", types.YChild{"SudiCertificate", &bootIntegrityOperData.SudiCertificate})
    bootIntegrityOperData.EntityData.Leafs = types.NewOrderedMap()

    bootIntegrityOperData.EntityData.YListKeys = []string {}

    return &(bootIntegrityOperData.EntityData)
}

// BootIntegrityOperData_BootIntegrity
// List of system boot integrity measurements for Boot, 
// Boot Loader, and OS versions, hashes, and signatures
// as well as Platform Configuration Registers (PCR) content.
// These measurements are captured utilizing Trusted Anchor Module (TAM)
// services to communicate with system ACT2 
// (2nd Gen Anti-Counterfeit Technology) hardware chip.
// This type is a presence type.
type BootIntegrityOperData_BootIntegrity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Product Identifier. The type is string.
    Platform interface{}

    // Boot 0 Version. The type is string.
    BootVer interface{}

    // Boot Loader Version. The type is string.
    BootLoaderVer interface{}

    // Operating System Version. The type is string.
    OsVersion interface{}

    // Boot 0 Hash. The type is string.
    BootHash interface{}

    // Boot Loader Hash. The type is string.
    BootLoaderHash interface{}

    // Operating System Hash. The type is string.
    OsHash interface{}

    // Number of Active Packages. The type is interface{} with range: 0..255.
    PackageCount interface{}

    // System ACT2 PCR Quote Signed Signature. The type is string.
    Signature interface{}

    // System ACT2 PCR Quote Signature Version. The type is interface{} with
    // range: 0..4294967295.
    SigVersion interface{}

    // System Platform Configuration Registers. The type is slice of
    // BootIntegrityOperData_BootIntegrity_PcrRegister.
    PcrRegister []*BootIntegrityOperData_BootIntegrity_PcrRegister

    // System Package Signatures. The type is slice of
    // BootIntegrityOperData_BootIntegrity_PackageSignature.
    PackageSignature []*BootIntegrityOperData_BootIntegrity_PackageSignature
}

func (bootIntegrity *BootIntegrityOperData_BootIntegrity) GetEntityData() *types.CommonEntityData {
    bootIntegrity.EntityData.YFilter = bootIntegrity.YFilter
    bootIntegrity.EntityData.YangName = "boot-integrity"
    bootIntegrity.EntityData.BundleName = "cisco_ios_xe"
    bootIntegrity.EntityData.ParentYangName = "boot-integrity-oper-data"
    bootIntegrity.EntityData.SegmentPath = "boot-integrity"
    bootIntegrity.EntityData.AbsolutePath = "Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data/" + bootIntegrity.EntityData.SegmentPath
    bootIntegrity.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bootIntegrity.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bootIntegrity.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bootIntegrity.EntityData.Children = types.NewOrderedMap()
    bootIntegrity.EntityData.Children.Append("pcr-register", types.YChild{"PcrRegister", nil})
    for i := range bootIntegrity.PcrRegister {
        bootIntegrity.EntityData.Children.Append(types.GetSegmentPath(bootIntegrity.PcrRegister[i]), types.YChild{"PcrRegister", bootIntegrity.PcrRegister[i]})
    }
    bootIntegrity.EntityData.Children.Append("package-signature", types.YChild{"PackageSignature", nil})
    for i := range bootIntegrity.PackageSignature {
        bootIntegrity.EntityData.Children.Append(types.GetSegmentPath(bootIntegrity.PackageSignature[i]), types.YChild{"PackageSignature", bootIntegrity.PackageSignature[i]})
    }
    bootIntegrity.EntityData.Leafs = types.NewOrderedMap()
    bootIntegrity.EntityData.Leafs.Append("platform", types.YLeaf{"Platform", bootIntegrity.Platform})
    bootIntegrity.EntityData.Leafs.Append("boot-ver", types.YLeaf{"BootVer", bootIntegrity.BootVer})
    bootIntegrity.EntityData.Leafs.Append("boot-loader-ver", types.YLeaf{"BootLoaderVer", bootIntegrity.BootLoaderVer})
    bootIntegrity.EntityData.Leafs.Append("os-version", types.YLeaf{"OsVersion", bootIntegrity.OsVersion})
    bootIntegrity.EntityData.Leafs.Append("boot-hash", types.YLeaf{"BootHash", bootIntegrity.BootHash})
    bootIntegrity.EntityData.Leafs.Append("boot-loader-hash", types.YLeaf{"BootLoaderHash", bootIntegrity.BootLoaderHash})
    bootIntegrity.EntityData.Leafs.Append("os-hash", types.YLeaf{"OsHash", bootIntegrity.OsHash})
    bootIntegrity.EntityData.Leafs.Append("package-count", types.YLeaf{"PackageCount", bootIntegrity.PackageCount})
    bootIntegrity.EntityData.Leafs.Append("signature", types.YLeaf{"Signature", bootIntegrity.Signature})
    bootIntegrity.EntityData.Leafs.Append("sig-version", types.YLeaf{"SigVersion", bootIntegrity.SigVersion})

    bootIntegrity.EntityData.YListKeys = []string {}

    return &(bootIntegrity.EntityData)
}

// BootIntegrityOperData_BootIntegrity_PcrRegister
// System Platform Configuration Registers
type BootIntegrityOperData_BootIntegrity_PcrRegister struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. References PCR Register Index. The type is
    // interface{} with range: 0..255.
    Index interface{}

    // References PCR Register Content. The type is string.
    PcrContent interface{}
}

func (pcrRegister *BootIntegrityOperData_BootIntegrity_PcrRegister) GetEntityData() *types.CommonEntityData {
    pcrRegister.EntityData.YFilter = pcrRegister.YFilter
    pcrRegister.EntityData.YangName = "pcr-register"
    pcrRegister.EntityData.BundleName = "cisco_ios_xe"
    pcrRegister.EntityData.ParentYangName = "boot-integrity"
    pcrRegister.EntityData.SegmentPath = "pcr-register" + types.AddKeyToken(pcrRegister.Index, "index")
    pcrRegister.EntityData.AbsolutePath = "Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data/boot-integrity/" + pcrRegister.EntityData.SegmentPath
    pcrRegister.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pcrRegister.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pcrRegister.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pcrRegister.EntityData.Children = types.NewOrderedMap()
    pcrRegister.EntityData.Leafs = types.NewOrderedMap()
    pcrRegister.EntityData.Leafs.Append("index", types.YLeaf{"Index", pcrRegister.Index})
    pcrRegister.EntityData.Leafs.Append("pcr-content", types.YLeaf{"PcrContent", pcrRegister.PcrContent})

    pcrRegister.EntityData.YListKeys = []string {"Index"}

    return &(pcrRegister.EntityData)
}

// BootIntegrityOperData_BootIntegrity_PackageSignature
// System Package Signatures
type BootIntegrityOperData_BootIntegrity_PackageSignature struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Package Name. The type is string.
    Name interface{}

    // Package Hash. The type is string.
    Hash interface{}
}

func (packageSignature *BootIntegrityOperData_BootIntegrity_PackageSignature) GetEntityData() *types.CommonEntityData {
    packageSignature.EntityData.YFilter = packageSignature.YFilter
    packageSignature.EntityData.YangName = "package-signature"
    packageSignature.EntityData.BundleName = "cisco_ios_xe"
    packageSignature.EntityData.ParentYangName = "boot-integrity"
    packageSignature.EntityData.SegmentPath = "package-signature" + types.AddKeyToken(packageSignature.Name, "name")
    packageSignature.EntityData.AbsolutePath = "Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data/boot-integrity/" + packageSignature.EntityData.SegmentPath
    packageSignature.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    packageSignature.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    packageSignature.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    packageSignature.EntityData.Children = types.NewOrderedMap()
    packageSignature.EntityData.Leafs = types.NewOrderedMap()
    packageSignature.EntityData.Leafs.Append("name", types.YLeaf{"Name", packageSignature.Name})
    packageSignature.EntityData.Leafs.Append("hash", types.YLeaf{"Hash", packageSignature.Hash})

    packageSignature.EntityData.YListKeys = []string {"Name"}

    return &(packageSignature.EntityData)
}

// BootIntegrityOperData_SudiCertificate
// List of system  certificate measurements for Cisco Root CA (CRCA),
// Cisco Manufacturing CA (CMCA), and ACT2 RSA Secure Unique Device 
// Identity (SUDI) CA PEM certifcates and SUDI generated signatures.
// These measurements are captured utilizing Trusted Anchor Module (TAM)
// services to communicate with system ACT2 
// (2nd Gen Anti-Counterfeit Technology) hardware chip.
// This type is a presence type.
type BootIntegrityOperData_SudiCertificate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Cisco Root CA PEM Certficate. The type is string.
    CrcaPem interface{}

    // Cisco Manufacturing CA PEM Certficate. The type is string.
    CmcaPem interface{}

    // ACT2 RSA SUDI CA PEM Certficate. The type is string.
    SudiPem interface{}

    // ACT2 RSA SUDI Certificate Generated Signature. The type is string.
    SudiSignature interface{}
}

func (sudiCertificate *BootIntegrityOperData_SudiCertificate) GetEntityData() *types.CommonEntityData {
    sudiCertificate.EntityData.YFilter = sudiCertificate.YFilter
    sudiCertificate.EntityData.YangName = "sudi-certificate"
    sudiCertificate.EntityData.BundleName = "cisco_ios_xe"
    sudiCertificate.EntityData.ParentYangName = "boot-integrity-oper-data"
    sudiCertificate.EntityData.SegmentPath = "sudi-certificate"
    sudiCertificate.EntityData.AbsolutePath = "Cisco-IOS-XE-boot-integrity-oper:boot-integrity-oper-data/" + sudiCertificate.EntityData.SegmentPath
    sudiCertificate.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sudiCertificate.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sudiCertificate.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sudiCertificate.EntityData.Children = types.NewOrderedMap()
    sudiCertificate.EntityData.Leafs = types.NewOrderedMap()
    sudiCertificate.EntityData.Leafs.Append("crca-pem", types.YLeaf{"CrcaPem", sudiCertificate.CrcaPem})
    sudiCertificate.EntityData.Leafs.Append("cmca-pem", types.YLeaf{"CmcaPem", sudiCertificate.CmcaPem})
    sudiCertificate.EntityData.Leafs.Append("sudi-pem", types.YLeaf{"SudiPem", sudiCertificate.SudiPem})
    sudiCertificate.EntityData.Leafs.Append("sudi-signature", types.YLeaf{"SudiSignature", sudiCertificate.SudiSignature})

    sudiCertificate.EntityData.YListKeys = []string {}

    return &(sudiCertificate.EntityData)
}

