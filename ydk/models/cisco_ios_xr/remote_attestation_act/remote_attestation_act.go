// This module defines procedure for remote attestation
//  of a network platform''s security posture.
//  This is useful to assess trustworthiness of
//  hardware and software of a network device.
// 
// Copyright (c) 2017 IETF Trust and the persons identified
// as authors of the code. All rights reserved.
// 
// Redistribution and use in source and binary forms, with
// or without modification, is permitted pursuant to, and
// subject to the license terms contained in, the Simplified
// BSD License set forth in Section 4.c of the IETF Trust''s
// Legal Provisions Relating to IETF Documents
// (http://trustee.ietf.org/license-info).
// 
// This version of this YANG module is part of RFC VVVV; see
// the RFC itself for full legal notices.
package remote_attestation_act

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package remote_attestation_act"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-remote-attestation-act get-certificate}", reflect.TypeOf(GetCertificate{}))
    ydk.RegisterEntity("Cisco-IOS-XR-remote-attestation-act:get-certificate", reflect.TypeOf(GetCertificate{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-remote-attestation-act attest-platform-config-registers}", reflect.TypeOf(AttestPlatformConfigRegisters{}))
    ydk.RegisterEntity("Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers", reflect.TypeOf(AttestPlatformConfigRegisters{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-remote-attestation-act get-platform-boot-integrity-event-logs}", reflect.TypeOf(GetPlatformBootIntegrityEventLogs{}))
    ydk.RegisterEntity("Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs", reflect.TypeOf(GetPlatformBootIntegrityEventLogs{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-remote-attestation-act get-platform-ima-event-logs}", reflect.TypeOf(GetPlatformImaEventLogs{}))
    ydk.RegisterEntity("Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs", reflect.TypeOf(GetPlatformImaEventLogs{}))
}

// GetCertificate
// Query certificate.
// Returns certificate chain
// associated with the queried certificate.
// An optional nonce can be provided, that is then used to
// return a signature over the certificate contents returned.
type GetCertificate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetCertificate_Input

    
    Output GetCertificate_Output
}

func (getCertificate *GetCertificate) GetEntityData() *types.CommonEntityData {
    getCertificate.EntityData.YFilter = getCertificate.YFilter
    getCertificate.EntityData.YangName = "get-certificate"
    getCertificate.EntityData.BundleName = "cisco_ios_xr"
    getCertificate.EntityData.ParentYangName = "Cisco-IOS-XR-remote-attestation-act"
    getCertificate.EntityData.SegmentPath = "Cisco-IOS-XR-remote-attestation-act:get-certificate"
    getCertificate.EntityData.AbsolutePath = getCertificate.EntityData.SegmentPath
    getCertificate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    getCertificate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    getCertificate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    getCertificate.EntityData.Children = types.NewOrderedMap()
    getCertificate.EntityData.Children.Append("input", types.YChild{"Input", &getCertificate.Input})
    getCertificate.EntityData.Children.Append("output", types.YChild{"Output", &getCertificate.Output})
    getCertificate.EntityData.Leafs = types.NewOrderedMap()

    getCertificate.EntityData.YListKeys = []string {}

    return &(getCertificate.EntityData)
}

// GetCertificate_Input
type GetCertificate_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce to be included in the attested output to prevent replay attacks. The
    // type is string with length: 0..64.
    Nonce interface{}

    // Certificate identifier. The type is string.
    CertificateIdentifier interface{}

    // In a distributed system get the data from a specific node identified by the
    // location. If this field is not specified data associated with each node
    // forming the system will be returned. The type is string.
    Location interface{}
}

func (input *GetCertificate_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "get-certificate"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("nonce", types.YLeaf{"Nonce", input.Nonce})
    input.EntityData.Leafs.Append("certificate-identifier", types.YLeaf{"CertificateIdentifier", input.CertificateIdentifier})
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetCertificate_Output
type GetCertificate_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    GetCertificateResponse GetCertificate_Output_GetCertificateResponse
}

func (output *GetCertificate_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "get-certificate"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("get-certificate-response", types.YChild{"GetCertificateResponse", &output.GetCertificateResponse})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// GetCertificate_Output_GetCertificateResponse
type GetCertificate_Output_GetCertificateResponse struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Certificate data of a node in a distributed system identified by the
    // location. The type is slice of
    // GetCertificate_Output_GetCertificateResponse_SystemCertificates.
    SystemCertificates []*GetCertificate_Output_GetCertificateResponse_SystemCertificates
}

func (getCertificateResponse *GetCertificate_Output_GetCertificateResponse) GetEntityData() *types.CommonEntityData {
    getCertificateResponse.EntityData.YFilter = getCertificateResponse.YFilter
    getCertificateResponse.EntityData.YangName = "get-certificate-response"
    getCertificateResponse.EntityData.BundleName = "cisco_ios_xr"
    getCertificateResponse.EntityData.ParentYangName = "output"
    getCertificateResponse.EntityData.SegmentPath = "get-certificate-response"
    getCertificateResponse.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/output/" + getCertificateResponse.EntityData.SegmentPath
    getCertificateResponse.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    getCertificateResponse.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    getCertificateResponse.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    getCertificateResponse.EntityData.Children = types.NewOrderedMap()
    getCertificateResponse.EntityData.Children.Append("system-certificates", types.YChild{"SystemCertificates", nil})
    for i := range getCertificateResponse.SystemCertificates {
        getCertificateResponse.EntityData.Children.Append(types.GetSegmentPath(getCertificateResponse.SystemCertificates[i]), types.YChild{"SystemCertificates", getCertificateResponse.SystemCertificates[i]})
    }
    getCertificateResponse.EntityData.Leafs = types.NewOrderedMap()

    getCertificateResponse.EntityData.YListKeys = []string {}

    return &(getCertificateResponse.EntityData)
}

// GetCertificate_Output_GetCertificateResponse_SystemCertificates
// Certificate data of a node in a distributed system
// identified by the location
type GetCertificate_Output_GetCertificateResponse_SystemCertificates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location of the node in the distributed system.
    // The type is string.
    NodeLocation interface{}

    // Nonce used for this output. The type is string with length: 0..64.
    Nonce interface{}

    // Signature version designates the format of the signed data. The type is
    // interface{} with range: 0..4294967295.
    SignatureVersion interface{}

    // The optional RSA or ECDSA signature across the certificates,the signature
    // version and the input nonce.Signed data format is: Nonce || UINT32
    // signature version || [Certificate included in the response in DER format].
    // The type is string.
    Signature interface{}

    // Certificates chain associated with the certificate being queried.
    Certificates GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates
}

func (systemCertificates *GetCertificate_Output_GetCertificateResponse_SystemCertificates) GetEntityData() *types.CommonEntityData {
    systemCertificates.EntityData.YFilter = systemCertificates.YFilter
    systemCertificates.EntityData.YangName = "system-certificates"
    systemCertificates.EntityData.BundleName = "cisco_ios_xr"
    systemCertificates.EntityData.ParentYangName = "get-certificate-response"
    systemCertificates.EntityData.SegmentPath = "system-certificates" + types.AddKeyToken(systemCertificates.NodeLocation, "node-location")
    systemCertificates.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/output/get-certificate-response/" + systemCertificates.EntityData.SegmentPath
    systemCertificates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemCertificates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemCertificates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemCertificates.EntityData.Children = types.NewOrderedMap()
    systemCertificates.EntityData.Children.Append("certificates", types.YChild{"Certificates", &systemCertificates.Certificates})
    systemCertificates.EntityData.Leafs = types.NewOrderedMap()
    systemCertificates.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", systemCertificates.NodeLocation})
    systemCertificates.EntityData.Leafs.Append("nonce", types.YLeaf{"Nonce", systemCertificates.Nonce})
    systemCertificates.EntityData.Leafs.Append("signature_version", types.YLeaf{"SignatureVersion", systemCertificates.SignatureVersion})
    systemCertificates.EntityData.Leafs.Append("signature", types.YLeaf{"Signature", systemCertificates.Signature})

    systemCertificates.EntityData.YListKeys = []string {"NodeLocation"}

    return &(systemCertificates.EntityData)
}

// GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates
// Certificates chain associated with the certificate
// being queried
type GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A X.509 certificate. The type is slice of
    // GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates_Certificate.
    Certificate []*GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates_Certificate
}

func (certificates *GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates) GetEntityData() *types.CommonEntityData {
    certificates.EntityData.YFilter = certificates.YFilter
    certificates.EntityData.YangName = "certificates"
    certificates.EntityData.BundleName = "cisco_ios_xr"
    certificates.EntityData.ParentYangName = "system-certificates"
    certificates.EntityData.SegmentPath = "certificates"
    certificates.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/output/get-certificate-response/system-certificates/" + certificates.EntityData.SegmentPath
    certificates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificates.EntityData.Children = types.NewOrderedMap()
    certificates.EntityData.Children.Append("certificate", types.YChild{"Certificate", nil})
    for i := range certificates.Certificate {
        certificates.EntityData.Children.Append(types.GetSegmentPath(certificates.Certificate[i]), types.YChild{"Certificate", certificates.Certificate[i]})
    }
    certificates.EntityData.Leafs = types.NewOrderedMap()

    certificates.EntityData.YListKeys = []string {}

    return &(certificates.EntityData)
}

// GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates_Certificate
// A X.509 certificate
type GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates_Certificate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. A node-unique certificate identifier. The type is
    // string.
    Name interface{}

    // Certificate content in DER format. The type is string.
    Value interface{}
}

func (certificate *GetCertificate_Output_GetCertificateResponse_SystemCertificates_Certificates_Certificate) GetEntityData() *types.CommonEntityData {
    certificate.EntityData.YFilter = certificate.YFilter
    certificate.EntityData.YangName = "certificate"
    certificate.EntityData.BundleName = "cisco_ios_xr"
    certificate.EntityData.ParentYangName = "certificates"
    certificate.EntityData.SegmentPath = "certificate" + types.AddKeyToken(certificate.Name, "name")
    certificate.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-certificate/output/get-certificate-response/system-certificates/certificates/" + certificate.EntityData.SegmentPath
    certificate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    certificate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    certificate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    certificate.EntityData.Children = types.NewOrderedMap()
    certificate.EntityData.Leafs = types.NewOrderedMap()
    certificate.EntityData.Leafs.Append("name", types.YLeaf{"Name", certificate.Name})
    certificate.EntityData.Leafs.Append("value", types.YLeaf{"Value", certificate.Value})

    certificate.EntityData.YListKeys = []string {"Name"}

    return &(certificate.EntityData)
}

// AttestPlatformConfigRegisters
// Attest Platform Configuration Register(PCRs)
type AttestPlatformConfigRegisters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input AttestPlatformConfigRegisters_Input

    
    Output AttestPlatformConfigRegisters_Output
}

func (attestPlatformConfigRegisters *AttestPlatformConfigRegisters) GetEntityData() *types.CommonEntityData {
    attestPlatformConfigRegisters.EntityData.YFilter = attestPlatformConfigRegisters.YFilter
    attestPlatformConfigRegisters.EntityData.YangName = "attest-platform-config-registers"
    attestPlatformConfigRegisters.EntityData.BundleName = "cisco_ios_xr"
    attestPlatformConfigRegisters.EntityData.ParentYangName = "Cisco-IOS-XR-remote-attestation-act"
    attestPlatformConfigRegisters.EntityData.SegmentPath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers"
    attestPlatformConfigRegisters.EntityData.AbsolutePath = attestPlatformConfigRegisters.EntityData.SegmentPath
    attestPlatformConfigRegisters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attestPlatformConfigRegisters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attestPlatformConfigRegisters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attestPlatformConfigRegisters.EntityData.Children = types.NewOrderedMap()
    attestPlatformConfigRegisters.EntityData.Children.Append("input", types.YChild{"Input", &attestPlatformConfigRegisters.Input})
    attestPlatformConfigRegisters.EntityData.Children.Append("output", types.YChild{"Output", &attestPlatformConfigRegisters.Output})
    attestPlatformConfigRegisters.EntityData.Leafs = types.NewOrderedMap()

    attestPlatformConfigRegisters.EntityData.YListKeys = []string {}

    return &(attestPlatformConfigRegisters.EntityData)
}

// AttestPlatformConfigRegisters_Input
type AttestPlatformConfigRegisters_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCR register indices to be included in the attested output. The type is
    // slice of interface{} with range: 0..65535.
    PcrIndex []interface{}

    // Nonce to be included in the attested output to prevent replay attacks. The
    // type is string with length: 0..64.
    Nonce interface{}

    // Identifier of the certificate to be used for attestation. The type is
    // string.
    AttestationCertificateIdentifier interface{}

    // In a distributed system get the data from a specific node identified by the
    // location. If this field is not specified data associated with each node
    // forming the system will be returned. The type is string.
    Location interface{}
}

func (input *AttestPlatformConfigRegisters_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "attest-platform-config-registers"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("pcr-index", types.YLeaf{"PcrIndex", input.PcrIndex})
    input.EntityData.Leafs.Append("nonce", types.YLeaf{"Nonce", input.Nonce})
    input.EntityData.Leafs.Append("attestation-certificate-identifier", types.YLeaf{"AttestationCertificateIdentifier", input.AttestationCertificateIdentifier})
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// AttestPlatformConfigRegisters_Output
type AttestPlatformConfigRegisters_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Attested Platform Config Register values.
    PlatformConfigRegisters AttestPlatformConfigRegisters_Output_PlatformConfigRegisters
}

func (output *AttestPlatformConfigRegisters_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "attest-platform-config-registers"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("platform-config-registers", types.YChild{"PlatformConfigRegisters", &output.PlatformConfigRegisters})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// AttestPlatformConfigRegisters_Output_PlatformConfigRegisters
// Attested Platform Config Register values
type AttestPlatformConfigRegisters_Output_PlatformConfigRegisters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Nonce used for this output. The type is string with length: 0..64.
    Nonce interface{}

    // Certificate data of a node in a distributed system identified by the
    // location. The type is slice of
    // AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData.
    NodeData []*AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData
}

func (platformConfigRegisters *AttestPlatformConfigRegisters_Output_PlatformConfigRegisters) GetEntityData() *types.CommonEntityData {
    platformConfigRegisters.EntityData.YFilter = platformConfigRegisters.YFilter
    platformConfigRegisters.EntityData.YangName = "platform-config-registers"
    platformConfigRegisters.EntityData.BundleName = "cisco_ios_xr"
    platformConfigRegisters.EntityData.ParentYangName = "output"
    platformConfigRegisters.EntityData.SegmentPath = "platform-config-registers"
    platformConfigRegisters.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers/output/" + platformConfigRegisters.EntityData.SegmentPath
    platformConfigRegisters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    platformConfigRegisters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    platformConfigRegisters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    platformConfigRegisters.EntityData.Children = types.NewOrderedMap()
    platformConfigRegisters.EntityData.Children.Append("node-data", types.YChild{"NodeData", nil})
    for i := range platformConfigRegisters.NodeData {
        platformConfigRegisters.EntityData.Children.Append(types.GetSegmentPath(platformConfigRegisters.NodeData[i]), types.YChild{"NodeData", platformConfigRegisters.NodeData[i]})
    }
    platformConfigRegisters.EntityData.Leafs = types.NewOrderedMap()
    platformConfigRegisters.EntityData.Leafs.Append("nonce", types.YLeaf{"Nonce", platformConfigRegisters.Nonce})

    platformConfigRegisters.EntityData.YListKeys = []string {}

    return &(platformConfigRegisters.EntityData)
}

// AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData
// Certificate data of a node in a distributed system
// identified by the location
type AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location of the node in the distributed system.
    // The type is string.
    NodeLocation interface{}

    // Uptime in seconds of this node reporting its data. The type is interface{}
    // with range: 0..4294967295.
    UpTime interface{}

    // TPM PCR Quote. The type is string.
    PcrQuote interface{}

    // PCR Quote signature using TPM-held ECC or RSA restricted key. The type is
    // string.
    PcrQuoteSignature interface{}

    // List of requested PCR contents. The type is slice of
    // AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData_PCR.
    PCR []*AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData_PCR
}

func (nodeData *AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData) GetEntityData() *types.CommonEntityData {
    nodeData.EntityData.YFilter = nodeData.YFilter
    nodeData.EntityData.YangName = "node-data"
    nodeData.EntityData.BundleName = "cisco_ios_xr"
    nodeData.EntityData.ParentYangName = "platform-config-registers"
    nodeData.EntityData.SegmentPath = "node-data" + types.AddKeyToken(nodeData.NodeLocation, "node-location")
    nodeData.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers/output/platform-config-registers/" + nodeData.EntityData.SegmentPath
    nodeData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeData.EntityData.Children = types.NewOrderedMap()
    nodeData.EntityData.Children.Append("PCR", types.YChild{"PCR", nil})
    for i := range nodeData.PCR {
        nodeData.EntityData.Children.Append(types.GetSegmentPath(nodeData.PCR[i]), types.YChild{"PCR", nodeData.PCR[i]})
    }
    nodeData.EntityData.Leafs = types.NewOrderedMap()
    nodeData.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", nodeData.NodeLocation})
    nodeData.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", nodeData.UpTime})
    nodeData.EntityData.Leafs.Append("pcr-quote", types.YLeaf{"PcrQuote", nodeData.PcrQuote})
    nodeData.EntityData.Leafs.Append("pcr-quote-signature", types.YLeaf{"PcrQuoteSignature", nodeData.PcrQuoteSignature})

    nodeData.EntityData.YListKeys = []string {"NodeLocation"}

    return &(nodeData.EntityData)
}

// AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData_PCR
// List of requested PCR contents
type AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData_PCR struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. PCR index. The type is interface{} with range:
    // 0..65535.
    Index interface{}

    // PCR register content. The type is slice of interface{} with range: 0..255.
    Value []interface{}
}

func (pCR *AttestPlatformConfigRegisters_Output_PlatformConfigRegisters_NodeData_PCR) GetEntityData() *types.CommonEntityData {
    pCR.EntityData.YFilter = pCR.YFilter
    pCR.EntityData.YangName = "PCR"
    pCR.EntityData.BundleName = "cisco_ios_xr"
    pCR.EntityData.ParentYangName = "node-data"
    pCR.EntityData.SegmentPath = "PCR" + types.AddKeyToken(pCR.Index, "index")
    pCR.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:attest-platform-config-registers/output/platform-config-registers/node-data/" + pCR.EntityData.SegmentPath
    pCR.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pCR.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pCR.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pCR.EntityData.Children = types.NewOrderedMap()
    pCR.EntityData.Leafs = types.NewOrderedMap()
    pCR.EntityData.Leafs.Append("index", types.YLeaf{"Index", pCR.Index})
    pCR.EntityData.Leafs.Append("value", types.YLeaf{"Value", pCR.Value})

    pCR.EntityData.YListKeys = []string {"Index"}

    return &(pCR.EntityData)
}

// GetPlatformBootIntegrityEventLogs
// Get platform's boot integrity
type GetPlatformBootIntegrityEventLogs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetPlatformBootIntegrityEventLogs_Input

    
    Output GetPlatformBootIntegrityEventLogs_Output
}

func (getPlatformBootIntegrityEventLogs *GetPlatformBootIntegrityEventLogs) GetEntityData() *types.CommonEntityData {
    getPlatformBootIntegrityEventLogs.EntityData.YFilter = getPlatformBootIntegrityEventLogs.YFilter
    getPlatformBootIntegrityEventLogs.EntityData.YangName = "get-platform-boot-integrity-event-logs"
    getPlatformBootIntegrityEventLogs.EntityData.BundleName = "cisco_ios_xr"
    getPlatformBootIntegrityEventLogs.EntityData.ParentYangName = "Cisco-IOS-XR-remote-attestation-act"
    getPlatformBootIntegrityEventLogs.EntityData.SegmentPath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs"
    getPlatformBootIntegrityEventLogs.EntityData.AbsolutePath = getPlatformBootIntegrityEventLogs.EntityData.SegmentPath
    getPlatformBootIntegrityEventLogs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    getPlatformBootIntegrityEventLogs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    getPlatformBootIntegrityEventLogs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    getPlatformBootIntegrityEventLogs.EntityData.Children = types.NewOrderedMap()
    getPlatformBootIntegrityEventLogs.EntityData.Children.Append("input", types.YChild{"Input", &getPlatformBootIntegrityEventLogs.Input})
    getPlatformBootIntegrityEventLogs.EntityData.Children.Append("output", types.YChild{"Output", &getPlatformBootIntegrityEventLogs.Output})
    getPlatformBootIntegrityEventLogs.EntityData.Leafs = types.NewOrderedMap()

    getPlatformBootIntegrityEventLogs.EntityData.YListKeys = []string {}

    return &(getPlatformBootIntegrityEventLogs.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Input
type GetPlatformBootIntegrityEventLogs_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // In a distributed system get the data from a specific node identified by the
    // location. If this field is not specified data associated with each node
    // forming the system will be returned. The type is string.
    Location interface{}

    // To filter event logs to be retrieved. - If set only events with sequence
    // number greater than that specified in this argument will be returned. The
    // type is interface{} with range: 0..18446744073709551615.
    StartEventNumber interface{}

    // To control event logs to be retrieved. - If set only events with sequence
    // number in the range of start-event-number to end-event-number will be
    // returned. The type is interface{} with range: 0..18446744073709551615.
    EndEventNumber interface{}
}

func (input *GetPlatformBootIntegrityEventLogs_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "get-platform-boot-integrity-event-logs"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})
    input.EntityData.Leafs.Append("start-event-number", types.YLeaf{"StartEventNumber", input.StartEventNumber})
    input.EntityData.Leafs.Append("end-event-number", types.YLeaf{"EndEventNumber", input.EndEventNumber})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output
type GetPlatformBootIntegrityEventLogs_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Boot integrity event logs.
    SystemBootIntegrity GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity
}

func (output *GetPlatformBootIntegrityEventLogs_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "get-platform-boot-integrity-event-logs"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("system-boot-integrity", types.YChild{"SystemBootIntegrity", &output.SystemBootIntegrity})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity
// Boot integrity event logs
type GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Boot integrity event logs of a node in a distributed system identified by
    // the location. The type is slice of
    // GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData.
    NodeData []*GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData
}

func (systemBootIntegrity *GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity) GetEntityData() *types.CommonEntityData {
    systemBootIntegrity.EntityData.YFilter = systemBootIntegrity.YFilter
    systemBootIntegrity.EntityData.YangName = "system-boot-integrity"
    systemBootIntegrity.EntityData.BundleName = "cisco_ios_xr"
    systemBootIntegrity.EntityData.ParentYangName = "output"
    systemBootIntegrity.EntityData.SegmentPath = "system-boot-integrity"
    systemBootIntegrity.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/output/" + systemBootIntegrity.EntityData.SegmentPath
    systemBootIntegrity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemBootIntegrity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemBootIntegrity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemBootIntegrity.EntityData.Children = types.NewOrderedMap()
    systemBootIntegrity.EntityData.Children.Append("node-data", types.YChild{"NodeData", nil})
    for i := range systemBootIntegrity.NodeData {
        systemBootIntegrity.EntityData.Children.Append(types.GetSegmentPath(systemBootIntegrity.NodeData[i]), types.YChild{"NodeData", systemBootIntegrity.NodeData[i]})
    }
    systemBootIntegrity.EntityData.Leafs = types.NewOrderedMap()

    systemBootIntegrity.EntityData.YListKeys = []string {}

    return &(systemBootIntegrity.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData
// Boot integrity event logs of a node in a distributed system
// identified by the location
type GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location of the node in the distributed system.
    // The type is string.
    NodeLocation interface{}

    // Uptime in seconds of this node reporting its data. The type is interface{}
    // with range: 0..4294967295.
    UpTime interface{}

    // Ordered list of TCG described event log that extended the PCRs in the order
    // they were logged. The type is slice of
    // GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog.
    EventLog []*GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog
}

func (nodeData *GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData) GetEntityData() *types.CommonEntityData {
    nodeData.EntityData.YFilter = nodeData.YFilter
    nodeData.EntityData.YangName = "node-data"
    nodeData.EntityData.BundleName = "cisco_ios_xr"
    nodeData.EntityData.ParentYangName = "system-boot-integrity"
    nodeData.EntityData.SegmentPath = "node-data" + types.AddKeyToken(nodeData.NodeLocation, "node-location")
    nodeData.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/output/system-boot-integrity/" + nodeData.EntityData.SegmentPath
    nodeData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeData.EntityData.Children = types.NewOrderedMap()
    nodeData.EntityData.Children.Append("event_log", types.YChild{"EventLog", nil})
    for i := range nodeData.EventLog {
        nodeData.EntityData.Children.Append(types.GetSegmentPath(nodeData.EventLog[i]), types.YChild{"EventLog", nodeData.EventLog[i]})
    }
    nodeData.EntityData.Leafs = types.NewOrderedMap()
    nodeData.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", nodeData.NodeLocation})
    nodeData.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", nodeData.UpTime})

    nodeData.EntityData.YListKeys = []string {"NodeLocation"}

    return &(nodeData.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog
// Ordered list of TCG described event log
// that extended the PCRs in the order they
// were logged
type GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique event number of this event. The type is
    // interface{} with range: 0..4294967295.
    EventNumber interface{}

    // log event type. The type is interface{} with range: 0..4294967295.
    EventType interface{}

    // Defines the PCR index that this event extended. The type is interface{}
    // with range: 0..65535.
    PcrIndex interface{}

    // Size of the event data. The type is interface{} with range: 0..4294967295.
    EventSize interface{}

    // the event data size determined by event-size. The type is string.
    EventData interface{}

    // Hash of event data. The type is slice of
    // GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList.
    DigestList []*GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList
}

func (eventLog *GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog) GetEntityData() *types.CommonEntityData {
    eventLog.EntityData.YFilter = eventLog.YFilter
    eventLog.EntityData.YangName = "event_log"
    eventLog.EntityData.BundleName = "cisco_ios_xr"
    eventLog.EntityData.ParentYangName = "node-data"
    eventLog.EntityData.SegmentPath = "event_log" + types.AddKeyToken(eventLog.EventNumber, "event-number")
    eventLog.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/output/system-boot-integrity/node-data/" + eventLog.EntityData.SegmentPath
    eventLog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventLog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventLog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventLog.EntityData.Children = types.NewOrderedMap()
    eventLog.EntityData.Children.Append("digest-list", types.YChild{"DigestList", nil})
    for i := range eventLog.DigestList {
        eventLog.EntityData.Children.Append(types.GetSegmentPath(eventLog.DigestList[i]), types.YChild{"DigestList", eventLog.DigestList[i]})
    }
    eventLog.EntityData.Leafs = types.NewOrderedMap()
    eventLog.EntityData.Leafs.Append("event-number", types.YLeaf{"EventNumber", eventLog.EventNumber})
    eventLog.EntityData.Leafs.Append("event-type", types.YLeaf{"EventType", eventLog.EventType})
    eventLog.EntityData.Leafs.Append("pcr-index", types.YLeaf{"PcrIndex", eventLog.PcrIndex})
    eventLog.EntityData.Leafs.Append("event-size", types.YLeaf{"EventSize", eventLog.EventSize})
    eventLog.EntityData.Leafs.Append("event-data", types.YLeaf{"EventData", eventLog.EventData})

    eventLog.EntityData.YListKeys = []string {"EventNumber"}

    return &(eventLog.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList
// Hash of event data
type GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Algorithm for this digest. The type is
    // DigestHashAlgorithm.
    DigestHashAlgorithm interface{}

    // The hash of the event data. The type is string.
    Digest interface{}
}

func (digestList *GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList) GetEntityData() *types.CommonEntityData {
    digestList.EntityData.YFilter = digestList.YFilter
    digestList.EntityData.YangName = "digest-list"
    digestList.EntityData.BundleName = "cisco_ios_xr"
    digestList.EntityData.ParentYangName = "event_log"
    digestList.EntityData.SegmentPath = "digest-list" + types.AddKeyToken(digestList.DigestHashAlgorithm, "digest-hash-algorithm")
    digestList.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-boot-integrity-event-logs/output/system-boot-integrity/node-data/event_log/" + digestList.EntityData.SegmentPath
    digestList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    digestList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    digestList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    digestList.EntityData.Children = types.NewOrderedMap()
    digestList.EntityData.Leafs = types.NewOrderedMap()
    digestList.EntityData.Leafs.Append("digest-hash-algorithm", types.YLeaf{"DigestHashAlgorithm", digestList.DigestHashAlgorithm})
    digestList.EntityData.Leafs.Append("digest", types.YLeaf{"Digest", digestList.Digest})

    digestList.EntityData.YListKeys = []string {"DigestHashAlgorithm"}

    return &(digestList.EntityData)
}

// GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm represents Algorithm for this digest
type GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm string

const (
    GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm_SHA1 GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm = "SHA1"

    GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm_SHA256 GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm = "SHA256"

    GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm_SHA384 GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm = "SHA384"

    GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm_SHA512 GetPlatformBootIntegrityEventLogs_Output_SystemBootIntegrity_NodeData_EventLog_DigestList_DigestHashAlgorithm = "SHA512"
)

// GetPlatformImaEventLogs
// Get platform IMA event logs
type GetPlatformImaEventLogs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input GetPlatformImaEventLogs_Input

    
    Output GetPlatformImaEventLogs_Output
}

func (getPlatformImaEventLogs *GetPlatformImaEventLogs) GetEntityData() *types.CommonEntityData {
    getPlatformImaEventLogs.EntityData.YFilter = getPlatformImaEventLogs.YFilter
    getPlatformImaEventLogs.EntityData.YangName = "get-platform-ima-event-logs"
    getPlatformImaEventLogs.EntityData.BundleName = "cisco_ios_xr"
    getPlatformImaEventLogs.EntityData.ParentYangName = "Cisco-IOS-XR-remote-attestation-act"
    getPlatformImaEventLogs.EntityData.SegmentPath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs"
    getPlatformImaEventLogs.EntityData.AbsolutePath = getPlatformImaEventLogs.EntityData.SegmentPath
    getPlatformImaEventLogs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    getPlatformImaEventLogs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    getPlatformImaEventLogs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    getPlatformImaEventLogs.EntityData.Children = types.NewOrderedMap()
    getPlatformImaEventLogs.EntityData.Children.Append("input", types.YChild{"Input", &getPlatformImaEventLogs.Input})
    getPlatformImaEventLogs.EntityData.Children.Append("output", types.YChild{"Output", &getPlatformImaEventLogs.Output})
    getPlatformImaEventLogs.EntityData.Leafs = types.NewOrderedMap()

    getPlatformImaEventLogs.EntityData.YListKeys = []string {}

    return &(getPlatformImaEventLogs.EntityData)
}

// GetPlatformImaEventLogs_Input
type GetPlatformImaEventLogs_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // In a distributed system get the data from a specific node identified by the
    // location. If this field is not specified data associated with each node
    // forming the system will be returned. The type is string.
    Location interface{}

    // To filter event logs to be retrieved. - If set only events with sequence
    // number greater than that specified in this argument will be returned. The
    // type is interface{} with range: 0..18446744073709551615.
    StartEventNumber interface{}

    // To control event logs to be retrieved. - If set only events with sequence
    // number in the range of start-event-number to end-event-number will be
    // returned. The type is interface{} with range: 0..18446744073709551615.
    EndEventNumber interface{}
}

func (input *GetPlatformImaEventLogs_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xr"
    input.EntityData.ParentYangName = "get-platform-ima-event-logs"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("location", types.YLeaf{"Location", input.Location})
    input.EntityData.Leafs.Append("start-event-number", types.YLeaf{"StartEventNumber", input.StartEventNumber})
    input.EntityData.Leafs.Append("end-event-number", types.YLeaf{"EndEventNumber", input.EndEventNumber})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// GetPlatformImaEventLogs_Output
type GetPlatformImaEventLogs_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Runtime integrity measurement event logs.
    SystemIma GetPlatformImaEventLogs_Output_SystemIma
}

func (output *GetPlatformImaEventLogs_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "get-platform-ima-event-logs"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Children.Append("system-ima", types.YChild{"SystemIma", &output.SystemIma})
    output.EntityData.Leafs = types.NewOrderedMap()

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// GetPlatformImaEventLogs_Output_SystemIma
// Runtime integrity measurement event logs
type GetPlatformImaEventLogs_Output_SystemIma struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IMA event logs of a node in a distributed system identified by the
    // location. The type is slice of
    // GetPlatformImaEventLogs_Output_SystemIma_NodeData.
    NodeData []*GetPlatformImaEventLogs_Output_SystemIma_NodeData
}

func (systemIma *GetPlatformImaEventLogs_Output_SystemIma) GetEntityData() *types.CommonEntityData {
    systemIma.EntityData.YFilter = systemIma.YFilter
    systemIma.EntityData.YangName = "system-ima"
    systemIma.EntityData.BundleName = "cisco_ios_xr"
    systemIma.EntityData.ParentYangName = "output"
    systemIma.EntityData.SegmentPath = "system-ima"
    systemIma.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs/output/" + systemIma.EntityData.SegmentPath
    systemIma.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    systemIma.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    systemIma.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    systemIma.EntityData.Children = types.NewOrderedMap()
    systemIma.EntityData.Children.Append("node-data", types.YChild{"NodeData", nil})
    for i := range systemIma.NodeData {
        systemIma.EntityData.Children.Append(types.GetSegmentPath(systemIma.NodeData[i]), types.YChild{"NodeData", systemIma.NodeData[i]})
    }
    systemIma.EntityData.Leafs = types.NewOrderedMap()

    systemIma.EntityData.YListKeys = []string {}

    return &(systemIma.EntityData)
}

// GetPlatformImaEventLogs_Output_SystemIma_NodeData
// IMA event logs of a node in a distributed system
// identified by the location
type GetPlatformImaEventLogs_Output_SystemIma_NodeData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Location of the node in the distributed system.
    // The type is string.
    NodeLocation interface{}

    // Uptime in seconds of this node reporting its data. The type is interface{}
    // with range: 0..4294967295.
    UpTime interface{}

    // Ordered list of ima event logs by event-number. The type is slice of
    // GetPlatformImaEventLogs_Output_SystemIma_NodeData_ImaEventLog.
    ImaEventLog []*GetPlatformImaEventLogs_Output_SystemIma_NodeData_ImaEventLog
}

func (nodeData *GetPlatformImaEventLogs_Output_SystemIma_NodeData) GetEntityData() *types.CommonEntityData {
    nodeData.EntityData.YFilter = nodeData.YFilter
    nodeData.EntityData.YangName = "node-data"
    nodeData.EntityData.BundleName = "cisco_ios_xr"
    nodeData.EntityData.ParentYangName = "system-ima"
    nodeData.EntityData.SegmentPath = "node-data" + types.AddKeyToken(nodeData.NodeLocation, "node-location")
    nodeData.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs/output/system-ima/" + nodeData.EntityData.SegmentPath
    nodeData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeData.EntityData.Children = types.NewOrderedMap()
    nodeData.EntityData.Children.Append("ima-event-log", types.YChild{"ImaEventLog", nil})
    for i := range nodeData.ImaEventLog {
        nodeData.EntityData.Children.Append(types.GetSegmentPath(nodeData.ImaEventLog[i]), types.YChild{"ImaEventLog", nodeData.ImaEventLog[i]})
    }
    nodeData.EntityData.Leafs = types.NewOrderedMap()
    nodeData.EntityData.Leafs.Append("node-location", types.YLeaf{"NodeLocation", nodeData.NodeLocation})
    nodeData.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", nodeData.UpTime})

    nodeData.EntityData.YListKeys = []string {"NodeLocation"}

    return &(nodeData.EntityData)
}

// GetPlatformImaEventLogs_Output_SystemIma_NodeData_ImaEventLog
// Ordered list of ima event logs by event-number
type GetPlatformImaEventLogs_Output_SystemIma_NodeData_ImaEventLog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unique number for this event for sequencing. The
    // type is interface{} with range: 0..18446744073709551615.
    EventNumber interface{}

    // Name of the template used for event for e.g. ima, ima-ng. The type is
    // string.
    ImaTemplate interface{}

    // File that was measured. The type is string.
    FilenameHint interface{}

    // Hash of filedata. The type is string.
    FiledataHash interface{}

    // Algorithm used for template-hash. The type is string.
    TemplateHashAlgorithm interface{}

    // hash(filedata-hash, filename-hint). The type is string.
    TemplateHash interface{}

    // Defines the PCR index that this event extended. The type is interface{}
    // with range: 0..65535.
    PcrIndex interface{}

    // The file signature. The type is string.
    Signature interface{}
}

func (imaEventLog *GetPlatformImaEventLogs_Output_SystemIma_NodeData_ImaEventLog) GetEntityData() *types.CommonEntityData {
    imaEventLog.EntityData.YFilter = imaEventLog.YFilter
    imaEventLog.EntityData.YangName = "ima-event-log"
    imaEventLog.EntityData.BundleName = "cisco_ios_xr"
    imaEventLog.EntityData.ParentYangName = "node-data"
    imaEventLog.EntityData.SegmentPath = "ima-event-log" + types.AddKeyToken(imaEventLog.EventNumber, "event-number")
    imaEventLog.EntityData.AbsolutePath = "Cisco-IOS-XR-remote-attestation-act:get-platform-ima-event-logs/output/system-ima/node-data/" + imaEventLog.EntityData.SegmentPath
    imaEventLog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    imaEventLog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    imaEventLog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    imaEventLog.EntityData.Children = types.NewOrderedMap()
    imaEventLog.EntityData.Leafs = types.NewOrderedMap()
    imaEventLog.EntityData.Leafs.Append("event-number", types.YLeaf{"EventNumber", imaEventLog.EventNumber})
    imaEventLog.EntityData.Leafs.Append("ima-template", types.YLeaf{"ImaTemplate", imaEventLog.ImaTemplate})
    imaEventLog.EntityData.Leafs.Append("filename-hint", types.YLeaf{"FilenameHint", imaEventLog.FilenameHint})
    imaEventLog.EntityData.Leafs.Append("filedata-hash", types.YLeaf{"FiledataHash", imaEventLog.FiledataHash})
    imaEventLog.EntityData.Leafs.Append("template-hash-algorithm", types.YLeaf{"TemplateHashAlgorithm", imaEventLog.TemplateHashAlgorithm})
    imaEventLog.EntityData.Leafs.Append("template-hash", types.YLeaf{"TemplateHash", imaEventLog.TemplateHash})
    imaEventLog.EntityData.Leafs.Append("pcr-index", types.YLeaf{"PcrIndex", imaEventLog.PcrIndex})
    imaEventLog.EntityData.Leafs.Append("signature", types.YLeaf{"Signature", imaEventLog.Signature})

    imaEventLog.EntityData.YListKeys = []string {"EventNumber"}

    return &(imaEventLog.EntityData)
}

