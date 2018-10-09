// This is the MIB module for the DOCSIS Baseline
// Privacy Plus Interface (BPI+) at cable modems (CMs)
// and cable modem termination systems (CMTSs).
// 
// Copyright (C) The Internet Society (2004). This
// version of this MIB module is part of RFC XXXX; see
// the RFC itself for full legal notices.
package docs_ietf_bpi2_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_ietf_bpi2_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-IETF-BPI2-MIB DOCS-IETF-BPI2-MIB}", reflect.TypeOf(DOCSIETFBPI2MIB{}))
    ydk.RegisterEntity("DOCS-IETF-BPI2-MIB:DOCS-IETF-BPI2-MIB", reflect.TypeOf(DOCSIETFBPI2MIB{}))
}

// DocsBpkmDataEncryptAlg represents being referenced has no data encryption.
type DocsBpkmDataEncryptAlg string

const (
    DocsBpkmDataEncryptAlg_none DocsBpkmDataEncryptAlg = "none"

    DocsBpkmDataEncryptAlg_des56CbcMode DocsBpkmDataEncryptAlg = "des56CbcMode"

    DocsBpkmDataEncryptAlg_des40CbcMode DocsBpkmDataEncryptAlg = "des40CbcMode"

    DocsBpkmDataEncryptAlg_t3Des128CbcMode DocsBpkmDataEncryptAlg = "t3Des128CbcMode"

    DocsBpkmDataEncryptAlg_aes128CbcMode DocsBpkmDataEncryptAlg = "aes128CbcMode"

    DocsBpkmDataEncryptAlg_aes256CbcMode DocsBpkmDataEncryptAlg = "aes256CbcMode"
)

// DocsBpkmDataAuthentAlg represents the SAID being referenced.
type DocsBpkmDataAuthentAlg string

const (
    DocsBpkmDataAuthentAlg_none DocsBpkmDataAuthentAlg = "none"

    DocsBpkmDataAuthentAlg_hmacSha196 DocsBpkmDataAuthentAlg = "hmacSha196"
)

// DocsBpkmSAType represents to be determined.
type DocsBpkmSAType string

const (
    DocsBpkmSAType_none DocsBpkmSAType = "none"

    DocsBpkmSAType_primary DocsBpkmSAType = "primary"

    DocsBpkmSAType_static DocsBpkmSAType = "static"

    DocsBpkmSAType_dynamic DocsBpkmSAType = "dynamic"
)

// DOCSIETFBPI2MIB
type DOCSIETFBPI2MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DocsIetfBpi2CodeDownloadControl DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl

    // This table describes the basic and authorization related Baseline Privacy
    // Plus attributes of each CM MAC interface.
    DocsIetfBpi2CmBaseTable DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable

    // This table describes the attributes of each CM Traffic Encryption Key (TEK)
    // association. The CM maintains (no more than) one TEK association per SAID
    // per CM MAC interface.
    DocsIetfBpi2CmTEKTable DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable

    // This table maps multicast IP addresses to SAIDs per CM MAC Interface. It is
    // intended to map multicast IP addresses associated with SA MAP Request
    // messages.
    DocsIetfBpi2CmIpMulticastMapTable DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable

    // This table describes the Baseline Privacy Plus device certificates for each
    // CM MAC interface.
    DocsIetfBpi2CmDeviceCertTable DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable

    // This table describes the Baseline Privacy Plus cryptographic suite
    // capabilities for each CM MAC interface.
    DocsIetfBpi2CmCryptoSuiteTable DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable

    // This table describes the basic Baseline Privacy attributes of each CMTS MAC
    // interface.
    DocsIetfBpi2CmtsBaseTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable

    // This table describes the attributes of each CM authorization association.
    // The CMTS maintains one authorization association with each Baseline
    // Privacy- enabled CM, registered on each CMTS MAC interface, regardless of
    // whether the CM is authorized or rejected.
    DocsIetfBpi2CmtsAuthTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable

    // This table describes the attributes of each Traffic Encryption Key (TEK)
    // association. The CMTS Maintains one TEK association per SAID on each CMTS
    // MAC interface.
    DocsIetfBpi2CmtsTEKTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable

    // This table maps multicast IP addresses to SAIDs. If a multicast IP address
    // is mapped by multiple rows in the table, the row with the lowest
    // docsIetfBpi2CmtsIpMulticastIndex must be utilized for the mapping.
    DocsIetfBpi2CmtsIpMulticastMapTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable

    // This table describes the multicast SAID authorization for each CM on each
    // CMTS MAC interface.
    DocsIetfBpi2CmtsMulticastAuthTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable

    // A table of CM certificate trust entries provisioned to the CMTS.  The trust
    // object for a certificate in this table has an overriding effect on the
    // validity object of a certificate in the authorization table, as long as the
    // entire contents of the two certificates are identical.
    DocsIetfBpi2CmtsProvisionedCmCertTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable

    // The table of known Certificate Authority certificates acquired by this
    // device.
    DocsIetfBpi2CmtsCACertTable DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable
}

func (dOCSIETFBPI2MIB *DOCSIETFBPI2MIB) GetEntityData() *types.CommonEntityData {
    dOCSIETFBPI2MIB.EntityData.YFilter = dOCSIETFBPI2MIB.YFilter
    dOCSIETFBPI2MIB.EntityData.YangName = "DOCS-IETF-BPI2-MIB"
    dOCSIETFBPI2MIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSIETFBPI2MIB.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    dOCSIETFBPI2MIB.EntityData.SegmentPath = "DOCS-IETF-BPI2-MIB:DOCS-IETF-BPI2-MIB"
    dOCSIETFBPI2MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSIETFBPI2MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSIETFBPI2MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSIETFBPI2MIB.EntityData.Children = types.NewOrderedMap()
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CodeDownloadControl", types.YChild{"DocsIetfBpi2CodeDownloadControl", &dOCSIETFBPI2MIB.DocsIetfBpi2CodeDownloadControl})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmBaseTable", types.YChild{"DocsIetfBpi2CmBaseTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmBaseTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmTEKTable", types.YChild{"DocsIetfBpi2CmTEKTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmTEKTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmIpMulticastMapTable", types.YChild{"DocsIetfBpi2CmIpMulticastMapTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmIpMulticastMapTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmDeviceCertTable", types.YChild{"DocsIetfBpi2CmDeviceCertTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmDeviceCertTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmCryptoSuiteTable", types.YChild{"DocsIetfBpi2CmCryptoSuiteTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmCryptoSuiteTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsBaseTable", types.YChild{"DocsIetfBpi2CmtsBaseTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsBaseTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsAuthTable", types.YChild{"DocsIetfBpi2CmtsAuthTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsAuthTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsTEKTable", types.YChild{"DocsIetfBpi2CmtsTEKTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsTEKTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsIpMulticastMapTable", types.YChild{"DocsIetfBpi2CmtsIpMulticastMapTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsIpMulticastMapTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsMulticastAuthTable", types.YChild{"DocsIetfBpi2CmtsMulticastAuthTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsMulticastAuthTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsProvisionedCmCertTable", types.YChild{"DocsIetfBpi2CmtsProvisionedCmCertTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsProvisionedCmCertTable})
    dOCSIETFBPI2MIB.EntityData.Children.Append("docsIetfBpi2CmtsCACertTable", types.YChild{"DocsIetfBpi2CmtsCACertTable", &dOCSIETFBPI2MIB.DocsIetfBpi2CmtsCACertTable})
    dOCSIETFBPI2MIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSIETFBPI2MIB.EntityData.YListKeys = []string {}

    return &(dOCSIETFBPI2MIB.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl
type DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value indicates the result of the latest config file CVC verification,
    // SNMP CVC verification, or code file verification. The type is
    // DocsIetfBpi2CodeDownloadStatusCode.
    DocsIetfBpi2CodeDownloadStatusCode interface{}

    // The value of this object indicates the additional information to the status
    // code.  The value will include the error code and error description which
    // will be defined separately. The type is string.
    DocsIetfBpi2CodeDownloadStatusString interface{}

    // The value of this object is the device manufacturer's organizationName. The
    // type is string.
    DocsIetfBpi2CodeMfgOrgName interface{}

    // The value of this object is the device manufacturer's current
    // codeAccessStart value. This value always be referenced to Greenwich Mean
    // Time (GMT) and the value format must contain TimeZone information (fields
    // 8-10). The type is string with length: 11.
    DocsIetfBpi2CodeMfgCodeAccessStart interface{}

    // The value of this object is the device manufacturer's current
    // cvcAccessStart value. This value always be referenced to Greenwich Mean
    // Time (GMT) and the value format must contain TimeZone information (fields
    // 8-10). The type is string with length: 11.
    DocsIetfBpi2CodeMfgCvcAccessStart interface{}

    // The value of this object is the Co-Signer's organizationName.  The value is
    // a zero length string if the co-signer is not specified. The type is string.
    DocsIetfBpi2CodeCoSignerOrgName interface{}

    // The value of this object is the Co-Signer's current codeAccessStart value.
    // This value always be referenced to Greenwich Mean Time (GMT) and the value
    // format must contain TimeZone information (fields 8-10). If
    // docsIetfBpi2CodeCoSignerOrgName is a zero length string, the value of this
    // object is meaningless. The type is string with length: 11.
    DocsIetfBpi2CodeCoSignerCodeAccessStart interface{}

    // The value of this object is the Co-Signer's current cvcAccessStart value.
    // This value always be referenced to Greenwich Mean Time (GMT) and the value
    // format must contain TimeZone information (fields 8-10). If
    // docsIetfBpi2CodeCoSignerOrgName is a zero length string, the value of this
    // object is meaningless. The type is string with length: 11.
    DocsIetfBpi2CodeCoSignerCvcAccessStart interface{}

    // Setting a CVC to this object triggers the device to verify the CVC and
    // update the cvcAccessStart values, then the content of this object is
    // discarded.. If the device is not enabled to upgrade codefiles, or the CVC
    // verification fails, the CVC will be rejected. Reading this object always
    // returns the zero-length OCTET STRING. The type is string with length:
    // 0..4096.
    DocsIetfBpi2CodeCvcUpdate interface{}
}

func (docsIetfBpi2CodeDownloadControl *DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CodeDownloadControl.EntityData.YFilter = docsIetfBpi2CodeDownloadControl.YFilter
    docsIetfBpi2CodeDownloadControl.EntityData.YangName = "docsIetfBpi2CodeDownloadControl"
    docsIetfBpi2CodeDownloadControl.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CodeDownloadControl.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CodeDownloadControl.EntityData.SegmentPath = "docsIetfBpi2CodeDownloadControl"
    docsIetfBpi2CodeDownloadControl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CodeDownloadControl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CodeDownloadControl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CodeDownloadControl.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeDownloadStatusCode", types.YLeaf{"DocsIetfBpi2CodeDownloadStatusCode", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeDownloadStatusCode})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeDownloadStatusString", types.YLeaf{"DocsIetfBpi2CodeDownloadStatusString", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeDownloadStatusString})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeMfgOrgName", types.YLeaf{"DocsIetfBpi2CodeMfgOrgName", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeMfgOrgName})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeMfgCodeAccessStart", types.YLeaf{"DocsIetfBpi2CodeMfgCodeAccessStart", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeMfgCodeAccessStart})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeMfgCvcAccessStart", types.YLeaf{"DocsIetfBpi2CodeMfgCvcAccessStart", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeMfgCvcAccessStart})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeCoSignerOrgName", types.YLeaf{"DocsIetfBpi2CodeCoSignerOrgName", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeCoSignerOrgName})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeCoSignerCodeAccessStart", types.YLeaf{"DocsIetfBpi2CodeCoSignerCodeAccessStart", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeCoSignerCodeAccessStart})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeCoSignerCvcAccessStart", types.YLeaf{"DocsIetfBpi2CodeCoSignerCvcAccessStart", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeCoSignerCvcAccessStart})
    docsIetfBpi2CodeDownloadControl.EntityData.Leafs.Append("docsIetfBpi2CodeCvcUpdate", types.YLeaf{"DocsIetfBpi2CodeCvcUpdate", docsIetfBpi2CodeDownloadControl.DocsIetfBpi2CodeCvcUpdate})

    docsIetfBpi2CodeDownloadControl.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CodeDownloadControl.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode represents verification.
type DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_configFileCvcVerified DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "configFileCvcVerified"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_configFileCvcRejected DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "configFileCvcRejected"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_snmpCvcVerified DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "snmpCvcVerified"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_snmpCvcRejected DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "snmpCvcRejected"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_codeFileVerified DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "codeFileVerified"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_codeFileRejected DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "codeFileRejected"

    DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode_other DOCSIETFBPI2MIB_DocsIetfBpi2CodeDownloadControl_DocsIetfBpi2CodeDownloadStatusCode = "other"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable
// This table describes the basic and authorization
// related Baseline Privacy Plus attributes of each CM MAC
// interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing attributes of one CM MAC interface.
    // An entry in this table exists for each ifEntry with an ifType of
    // docsCableMaclayer(127). The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry.
    DocsIetfBpi2CmBaseEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry
}

func (docsIetfBpi2CmBaseTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmBaseTable.EntityData.YFilter = docsIetfBpi2CmBaseTable.YFilter
    docsIetfBpi2CmBaseTable.EntityData.YangName = "docsIetfBpi2CmBaseTable"
    docsIetfBpi2CmBaseTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmBaseTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmBaseTable.EntityData.SegmentPath = "docsIetfBpi2CmBaseTable"
    docsIetfBpi2CmBaseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmBaseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmBaseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmBaseTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmBaseTable.EntityData.Children.Append("docsIetfBpi2CmBaseEntry", types.YChild{"DocsIetfBpi2CmBaseEntry", nil})
    for i := range docsIetfBpi2CmBaseTable.DocsIetfBpi2CmBaseEntry {
        docsIetfBpi2CmBaseTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmBaseTable.DocsIetfBpi2CmBaseEntry[i]), types.YChild{"DocsIetfBpi2CmBaseEntry", docsIetfBpi2CmBaseTable.DocsIetfBpi2CmBaseEntry[i]})
    }
    docsIetfBpi2CmBaseTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmBaseTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmBaseTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry
// Each entry contains objects describing attributes of
// one CM MAC interface. An entry in this table exists for
// each ifEntry with an ifType of docsCableMaclayer(127).
type DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This object identifies whether this CM is provisioned to run Baseline
    // Privacy Plus. The type is bool.
    DocsIetfBpi2CmPrivacyEnable interface{}

    // The value of this object is a DER-encoded RSAPublicKey ASN.1 type string,
    // as defined in the RSA Encryption Standard (PKCS #1), corresponding to the
    // public key of the CM. The type is string with length: 0..524.
    DocsIetfBpi2CmPublicKey interface{}

    // The value of this object is the state of the CM authorization FSM.  The
    // start state indicates that FSM is in its initial state. The type is
    // DocsIetfBpi2CmAuthState.
    DocsIetfBpi2CmAuthState interface{}

    // The value of this object is the most recent authorization key sequence
    // number for this FSM. The type is interface{} with range: 0..15.
    DocsIetfBpi2CmAuthKeySequenceNumber interface{}

    // The value of this object is the actual clock time for expiration of the
    // immediate predecessor of the most recent authorization key for this FSM. 
    // If this FSM has only one authorization key, then the value is the time of
    // activation of this FSM. The type is string.
    DocsIetfBpi2CmAuthExpiresOld interface{}

    // The value of this object is the actual clock time for expiration of the
    // most recent authorization key for this FSM. The type is string.
    DocsIetfBpi2CmAuthExpiresNew interface{}

    // Setting this object to 'true' generates a Reauthorize event in the
    // authorization FSM. Reading this object always returns FALSE. This object is
    // for testing purposes only and therefore it does not require to be
    // associated with a last reset object. The type is bool.
    DocsIetfBpi2CmAuthReset interface{}

    // The value of this object is the grace time for an authorization key in
    // seconds.  A CM is expected to start trying to get a new authorization key
    // beginning AuthGraceTime seconds before the most recent authorization key
    // actually expires. The type is interface{} with range: 1..6047999. Units are
    // seconds.
    DocsIetfBpi2CmAuthGraceTime interface{}

    // The value of this object is the grace time for the TEK in seconds.  The CM
    // is expected to start trying to acquire a new TEK beginning TEK GraceTime
    // seconds before the expiration of the most recent TEK. The type is
    // interface{} with range: 1..302399. Units are seconds.
    DocsIetfBpi2CmTEKGraceTime interface{}

    // The value of this object is the Authorize Wait Timeout in second. The type
    // is interface{} with range: 1..30. Units are seconds.
    DocsIetfBpi2CmAuthWaitTimeout interface{}

    // The value of this object is the Reauthorize Wait Timeout in seconds. The
    // type is interface{} with range: 1..30. Units are seconds.
    DocsIetfBpi2CmReauthWaitTimeout interface{}

    // The value of this object is the Operational Wait Timeout in seconds. The
    // type is interface{} with range: 1..10. Units are seconds.
    DocsIetfBpi2CmOpWaitTimeout interface{}

    // The value of this object is the Rekey Wait Timeout in seconds. The type is
    // interface{} with range: 1..10. Units are seconds.
    DocsIetfBpi2CmRekeyWaitTimeout interface{}

    // The value of this object is the Authorization Reject Wait Timeout in
    // seconds. The type is interface{} with range: 1..600. Units are seconds.
    DocsIetfBpi2CmAuthRejectWaitTimeout interface{}

    // The value of this object is the retransmission interval, in seconds, of SA
    // Map Requests from the MAP Wait state. The type is interface{} with range:
    // 1..10. Units are seconds.
    DocsIetfBpi2CmSAMapWaitTimeout interface{}

    // The value of this object is the maximum number of Map Request retries
    // allowed. The type is interface{} with range: 0..10. Units are count.
    DocsIetfBpi2CmSAMapMaxRetries interface{}

    // The value of this object is the count of times the CM has transmitted an
    // Authentication Information message. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmAuthentInfos interface{}

    // The value of this object is the count of times the CM has transmitted an
    // Authorization Request message. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmAuthRequests interface{}

    // The value of this object is the count of times the CM has received an
    // Authorization Reply message. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmAuthReplies interface{}

    // The value of this object is the count of times the CM has received an
    // Authorization Reject message. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmAuthRejects interface{}

    // The value of this object is the count of times the CM has received an
    // Authorization Invalid message. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmAuthInvalids interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent Authorization Reject message received by the CM.  This has
    // value unknown(2) if the last Error-Code value was 0, and none(1) if no
    // Authorization Reject message has been received since reboot. The type is
    // DocsIetfBpi2CmAuthRejectErrorCode.
    DocsIetfBpi2CmAuthRejectErrorCode interface{}

    // The value of this object is the text string in most recent Authorization
    // Reject message received by the CM.  This is a zero length string if no
    // Authorization Reject message has been received since reboot. The type is
    // string with length: 0..128.
    DocsIetfBpi2CmAuthRejectErrorString interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent Authorization Invalid message received by the CM.  This has
    // value unknown(2) if the last Error-Code value was 0, and none(1) if no
    // Authorization Invalid message has been received since reboot. The type is
    // DocsIetfBpi2CmAuthInvalidErrorCode.
    DocsIetfBpi2CmAuthInvalidErrorCode interface{}

    // The value of this object is the text string in most recent Authorization
    // Invalid message received by the CM. This is a zero length string if no
    // Authorization Invalid message has been received since reboot. The type is
    // string with length: 0..128.
    DocsIetfBpi2CmAuthInvalidErrorString interface{}
}

func (docsIetfBpi2CmBaseEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmBaseEntry.EntityData.YFilter = docsIetfBpi2CmBaseEntry.YFilter
    docsIetfBpi2CmBaseEntry.EntityData.YangName = "docsIetfBpi2CmBaseEntry"
    docsIetfBpi2CmBaseEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmBaseEntry.EntityData.ParentYangName = "docsIetfBpi2CmBaseTable"
    docsIetfBpi2CmBaseEntry.EntityData.SegmentPath = "docsIetfBpi2CmBaseEntry" + types.AddKeyToken(docsIetfBpi2CmBaseEntry.IfIndex, "ifIndex")
    docsIetfBpi2CmBaseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmBaseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmBaseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmBaseEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmBaseEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmBaseEntry.IfIndex})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmPrivacyEnable", types.YLeaf{"DocsIetfBpi2CmPrivacyEnable", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmPrivacyEnable})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmPublicKey", types.YLeaf{"DocsIetfBpi2CmPublicKey", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmPublicKey})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthState", types.YLeaf{"DocsIetfBpi2CmAuthState", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthState})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthKeySequenceNumber", types.YLeaf{"DocsIetfBpi2CmAuthKeySequenceNumber", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthKeySequenceNumber})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthExpiresOld", types.YLeaf{"DocsIetfBpi2CmAuthExpiresOld", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthExpiresOld})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthExpiresNew", types.YLeaf{"DocsIetfBpi2CmAuthExpiresNew", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthExpiresNew})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthReset", types.YLeaf{"DocsIetfBpi2CmAuthReset", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthReset})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthGraceTime", types.YLeaf{"DocsIetfBpi2CmAuthGraceTime", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthGraceTime})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKGraceTime", types.YLeaf{"DocsIetfBpi2CmTEKGraceTime", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmTEKGraceTime})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthWaitTimeout", types.YLeaf{"DocsIetfBpi2CmAuthWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmReauthWaitTimeout", types.YLeaf{"DocsIetfBpi2CmReauthWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmReauthWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmOpWaitTimeout", types.YLeaf{"DocsIetfBpi2CmOpWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmOpWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmRekeyWaitTimeout", types.YLeaf{"DocsIetfBpi2CmRekeyWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmRekeyWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthRejectWaitTimeout", types.YLeaf{"DocsIetfBpi2CmAuthRejectWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthRejectWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmSAMapWaitTimeout", types.YLeaf{"DocsIetfBpi2CmSAMapWaitTimeout", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmSAMapWaitTimeout})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmSAMapMaxRetries", types.YLeaf{"DocsIetfBpi2CmSAMapMaxRetries", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmSAMapMaxRetries})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthentInfos", types.YLeaf{"DocsIetfBpi2CmAuthentInfos", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthentInfos})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthRequests", types.YLeaf{"DocsIetfBpi2CmAuthRequests", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthRequests})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthReplies", types.YLeaf{"DocsIetfBpi2CmAuthReplies", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthReplies})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthRejects", types.YLeaf{"DocsIetfBpi2CmAuthRejects", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthRejects})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthInvalids", types.YLeaf{"DocsIetfBpi2CmAuthInvalids", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthInvalids})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmAuthRejectErrorCode", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthRejectErrorCode})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthRejectErrorString", types.YLeaf{"DocsIetfBpi2CmAuthRejectErrorString", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthRejectErrorString})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthInvalidErrorCode", types.YLeaf{"DocsIetfBpi2CmAuthInvalidErrorCode", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthInvalidErrorCode})
    docsIetfBpi2CmBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmAuthInvalidErrorString", types.YLeaf{"DocsIetfBpi2CmAuthInvalidErrorString", docsIetfBpi2CmBaseEntry.DocsIetfBpi2CmAuthInvalidErrorString})

    docsIetfBpi2CmBaseEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIetfBpi2CmBaseEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode represents reboot.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_unauthorizedCm DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "unauthorizedCm"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_unsolicited DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "unsolicited"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_invalidKeySequence DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "invalidKeySequence"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode_keyRequestAuthenticationFailure DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthInvalidErrorCode = "keyRequestAuthenticationFailure"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode represents reboot.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_unauthorizedCm DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "unauthorizedCm"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_unauthorizedSaid DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "unauthorizedSaid"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_permanentAuthorizationFailure DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "permanentAuthorizationFailure"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode_timeOfDayNotAcquired DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthRejectErrorCode = "timeOfDayNotAcquired"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState represents in its initial state.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_start DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "start"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_authWait DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "authWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_authorized DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "authorized"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_reauthWait DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "reauthWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_authRejectWait DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "authRejectWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState_silent DOCSIETFBPI2MIB_DocsIetfBpi2CmBaseTable_DocsIetfBpi2CmBaseEntry_DocsIetfBpi2CmAuthState = "silent"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable
// This table describes the attributes of each CM
// Traffic Encryption Key (TEK) association. The CM maintains
// (no more than) one TEK association per SAID per CM MAC
// interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing the TEK association attributes of
    // one SAID. The CM MUST create one entry per SAID, regardless of whether the
    // SAID was obtained from a Registration Response message, from an
    // Authorization Reply message, or from any dynamic SAID establishment
    // mechanisms. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry.
    DocsIetfBpi2CmTEKEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry
}

func (docsIetfBpi2CmTEKTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmTEKTable.EntityData.YFilter = docsIetfBpi2CmTEKTable.YFilter
    docsIetfBpi2CmTEKTable.EntityData.YangName = "docsIetfBpi2CmTEKTable"
    docsIetfBpi2CmTEKTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmTEKTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmTEKTable.EntityData.SegmentPath = "docsIetfBpi2CmTEKTable"
    docsIetfBpi2CmTEKTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmTEKTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmTEKTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmTEKTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmTEKTable.EntityData.Children.Append("docsIetfBpi2CmTEKEntry", types.YChild{"DocsIetfBpi2CmTEKEntry", nil})
    for i := range docsIetfBpi2CmTEKTable.DocsIetfBpi2CmTEKEntry {
        docsIetfBpi2CmTEKTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmTEKTable.DocsIetfBpi2CmTEKEntry[i]), types.YChild{"DocsIetfBpi2CmTEKEntry", docsIetfBpi2CmTEKTable.DocsIetfBpi2CmTEKEntry[i]})
    }
    docsIetfBpi2CmTEKTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmTEKTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmTEKTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry
// Each entry contains objects describing the TEK
// association attributes of one SAID. The CM MUST create one
// entry per SAID, regardless of whether the SAID was obtained
// from a Registration Response message, from an Authorization
// Reply message, or from any dynamic SAID establishment
// mechanisms.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The value of this object is the DOCSIS Security
    // Association ID (SAID). The type is interface{} with range: 1..16383.
    DocsIetfBpi2CmTEKSAId interface{}

    // The value of this object is the type of security association. The type is
    // DocsBpkmSAType.
    DocsIetfBpi2CmTEKSAType interface{}

    // The value of this object is the data encryption algorithm for this SAID.
    // The type is DocsBpkmDataEncryptAlg.
    DocsIetfBpi2CmTEKDataEncryptAlg interface{}

    // The value of this object is the data authentication algorithm for this
    // SAID. The type is DocsBpkmDataAuthentAlg.
    DocsIetfBpi2CmTEKDataAuthentAlg interface{}

    // The value of this object is the state of the indicated TEK FSM.  The
    // start(1) state indicates that FSM is in its initial state. The type is
    // DocsIetfBpi2CmTEKState.
    DocsIetfBpi2CmTEKState interface{}

    // The value of this object is the most recent TEK key sequence number for
    // this TEK FSM. The type is interface{} with range: 0..15.
    DocsIetfBpi2CmTEKKeySequenceNumber interface{}

    // The value of this object is the actual clock time for expiration of the
    // immediate predecessor of the most recent TEK for this FSM.  If this FSM has
    // only one TEK, then the value is the time of activation of this FSM. The
    // type is string.
    DocsIetfBpi2CmTEKExpiresOld interface{}

    // The value of this object is the actual clock time for expiration of the
    // most recent TEK for this FSM. The type is string.
    DocsIetfBpi2CmTEKExpiresNew interface{}

    // The value of this object is the count of times the CM has transmitted a Key
    // Request message. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    DocsIetfBpi2CmTEKKeyRequests interface{}

    // The value of this object is the count of times the CM has received a Key
    // Reply message, including a message whose authentication failed.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    DocsIetfBpi2CmTEKKeyReplies interface{}

    // The value of this object is the count of times the CM has received a Key
    // Reject message, including a message whose authentication failed.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    DocsIetfBpi2CmTEKKeyRejects interface{}

    // The value of this object is the count of times the CM has received a TEK
    // Invalid message, including a message whose authentication failed.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    DocsIetfBpi2CmTEKInvalids interface{}

    // The value of this object is the count of times an Authorization Pending
    // (Auth Pend) event occurred in this FSM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmTEKAuthPends interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent Key Reject message received by the CM. This has value
    // unknown(2) if the last Error-Code value was 0, and none(1) if no Key Reject
    // message has been received since registration. The type is
    // DocsIetfBpi2CmTEKKeyRejectErrorCode.
    DocsIetfBpi2CmTEKKeyRejectErrorCode interface{}

    // The value of this object is the text string in most recent Key Reject
    // message received by the CM. This is a zero length string if no Key Reject
    // message has been received since registration. The type is string with
    // length: 0..128.
    DocsIetfBpi2CmTEKKeyRejectErrorString interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent TEK Invalid message received by the CM.  This has value
    // unknown(2) if the last Error-Code value was 0, and none(1) if no TEK
    // Invalid message has been received since registration. The type is
    // DocsIetfBpi2CmTEKInvalidErrorCode.
    DocsIetfBpi2CmTEKInvalidErrorCode interface{}

    // The value of this object is the text string in most recent TEK Invalid
    // message received by the CM. This is a zero length string if no TEK Invalid
    // message has been received since registration. The type is string with
    // length: 0..128.
    DocsIetfBpi2CmTEKInvalidErrorString interface{}
}

func (docsIetfBpi2CmTEKEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmTEKEntry.EntityData.YFilter = docsIetfBpi2CmTEKEntry.YFilter
    docsIetfBpi2CmTEKEntry.EntityData.YangName = "docsIetfBpi2CmTEKEntry"
    docsIetfBpi2CmTEKEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmTEKEntry.EntityData.ParentYangName = "docsIetfBpi2CmTEKTable"
    docsIetfBpi2CmTEKEntry.EntityData.SegmentPath = "docsIetfBpi2CmTEKEntry" + types.AddKeyToken(docsIetfBpi2CmTEKEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKSAId, "docsIetfBpi2CmTEKSAId")
    docsIetfBpi2CmTEKEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmTEKEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmTEKEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmTEKEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmTEKEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmTEKEntry.IfIndex})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKSAId", types.YLeaf{"DocsIetfBpi2CmTEKSAId", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKSAId})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKSAType", types.YLeaf{"DocsIetfBpi2CmTEKSAType", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKSAType})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKDataEncryptAlg", types.YLeaf{"DocsIetfBpi2CmTEKDataEncryptAlg", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKDataEncryptAlg})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKDataAuthentAlg", types.YLeaf{"DocsIetfBpi2CmTEKDataAuthentAlg", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKDataAuthentAlg})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKState", types.YLeaf{"DocsIetfBpi2CmTEKState", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKState})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeySequenceNumber", types.YLeaf{"DocsIetfBpi2CmTEKKeySequenceNumber", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeySequenceNumber})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKExpiresOld", types.YLeaf{"DocsIetfBpi2CmTEKExpiresOld", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKExpiresOld})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKExpiresNew", types.YLeaf{"DocsIetfBpi2CmTEKExpiresNew", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKExpiresNew})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeyRequests", types.YLeaf{"DocsIetfBpi2CmTEKKeyRequests", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeyRequests})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeyReplies", types.YLeaf{"DocsIetfBpi2CmTEKKeyReplies", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeyReplies})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeyRejects", types.YLeaf{"DocsIetfBpi2CmTEKKeyRejects", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeyRejects})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKInvalids", types.YLeaf{"DocsIetfBpi2CmTEKInvalids", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKInvalids})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKAuthPends", types.YLeaf{"DocsIetfBpi2CmTEKAuthPends", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKAuthPends})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeyRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmTEKKeyRejectErrorCode", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeyRejectErrorCode})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKKeyRejectErrorString", types.YLeaf{"DocsIetfBpi2CmTEKKeyRejectErrorString", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKKeyRejectErrorString})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKInvalidErrorCode", types.YLeaf{"DocsIetfBpi2CmTEKInvalidErrorCode", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKInvalidErrorCode})
    docsIetfBpi2CmTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmTEKInvalidErrorString", types.YLeaf{"DocsIetfBpi2CmTEKInvalidErrorString", docsIetfBpi2CmTEKEntry.DocsIetfBpi2CmTEKInvalidErrorString})

    docsIetfBpi2CmTEKEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmTEKSAId"}

    return &(docsIetfBpi2CmTEKEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode represents Invalid message has been received since registration.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode_invalidKeySequence DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKInvalidErrorCode = "invalidKeySequence"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode represents Reject message has been received since registration.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode_unauthorizedSaid DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKKeyRejectErrorCode = "unauthorizedSaid"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState represents is in its initial state.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_start DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "start"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_opWait DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "opWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_opReauthWait DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "opReauthWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_operational DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "operational"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_rekeyWait DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "rekeyWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState_rekeyReauthWait DOCSIETFBPI2MIB_DocsIetfBpi2CmTEKTable_DocsIetfBpi2CmTEKEntry_DocsIetfBpi2CmTEKState = "rekeyReauthWait"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable
// This table maps multicast IP addresses to SAIDs per
// CM MAC Interface.
// It is intended to map multicast IP addresses associated
// with SA MAP Request messages.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing the mapping of one multicast IP
    // address to one SAID, as well as associated state, message counters, and
    // error information.  An entry may be removed from this table upon the
    // reception of an SA Map Reject. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry.
    DocsIetfBpi2CmIpMulticastMapEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry
}

func (docsIetfBpi2CmIpMulticastMapTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmIpMulticastMapTable.EntityData.YFilter = docsIetfBpi2CmIpMulticastMapTable.YFilter
    docsIetfBpi2CmIpMulticastMapTable.EntityData.YangName = "docsIetfBpi2CmIpMulticastMapTable"
    docsIetfBpi2CmIpMulticastMapTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmIpMulticastMapTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmIpMulticastMapTable.EntityData.SegmentPath = "docsIetfBpi2CmIpMulticastMapTable"
    docsIetfBpi2CmIpMulticastMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmIpMulticastMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmIpMulticastMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmIpMulticastMapTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmIpMulticastMapTable.EntityData.Children.Append("docsIetfBpi2CmIpMulticastMapEntry", types.YChild{"DocsIetfBpi2CmIpMulticastMapEntry", nil})
    for i := range docsIetfBpi2CmIpMulticastMapTable.DocsIetfBpi2CmIpMulticastMapEntry {
        docsIetfBpi2CmIpMulticastMapTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmIpMulticastMapTable.DocsIetfBpi2CmIpMulticastMapEntry[i]), types.YChild{"DocsIetfBpi2CmIpMulticastMapEntry", docsIetfBpi2CmIpMulticastMapTable.DocsIetfBpi2CmIpMulticastMapEntry[i]})
    }
    docsIetfBpi2CmIpMulticastMapTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmIpMulticastMapTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmIpMulticastMapTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry
// Each entry contains objects describing the mapping of
// one multicast IP address to one SAID, as well as
// associated state, message counters, and error information.
// 
// An entry may be removed from this table upon the reception
// of an SA Map Reject.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The index of this row. The type is interface{}
    // with range: 1..4294967295.
    DocsIetfBpi2CmIpMulticastIndex interface{}

    // The type of internet address for docsIetfBpi2CmIpMulticastAddress. The type
    // is InetAddressType.
    DocsIetfBpi2CmIpMulticastAddressType interface{}

    // This object represents the IP multicast address to be mapped. The type of
    // this address is determined by the value of the
    // docsIetfBpi2CmIpMulticastAddressType object. The type is string with
    // length: 0..255.
    DocsIetfBpi2CmIpMulticastAddress interface{}

    // This object represents the SAID to which the IP multicast address has been
    // mapped.  If no SA Map Reply has been received for the IP address, this
    // object should have the value 0. The type is interface{} with range:
    // 0..16383.
    DocsIetfBpi2CmIpMulticastSAId interface{}

    // The value of this object is the state of the SA Mapping FSM for this IP.
    // The type is DocsIetfBpi2CmIpMulticastSAMapState.
    DocsIetfBpi2CmIpMulticastSAMapState interface{}

    // The value of this object is the count of times the CM has transmitted an SA
    // Map Request message for this IP. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmIpMulticastSAMapRequests interface{}

    // The value of this object is the count of times the CM has received an SA
    // Map Reply message for this IP. Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmIpMulticastSAMapReplies interface{}

    // The value of this object is the count of times the CM has received an SA
    // MAP Reject message for this IP. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmIpMulticastSAMapRejects interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // the most recent SA Map Reject message sent in response to an SA Map Request
    // for This IP.  It has the value none(1) if no SA MAP Reject message has been
    // received since entry creation. The type is
    // DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode.
    DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode interface{}

    // The value of this object is the text string in the most recent SA Map
    // Reject message sent in response to an SA Map Request for this IP.  It is a
    // zero length string if no SA Map Reject message has been received since
    // entry creation. The type is string with length: 0..128.
    DocsIetfBpi2CmIpMulticastSAMapRejectErrorString interface{}
}

func (docsIetfBpi2CmIpMulticastMapEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.YFilter = docsIetfBpi2CmIpMulticastMapEntry.YFilter
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.YangName = "docsIetfBpi2CmIpMulticastMapEntry"
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.ParentYangName = "docsIetfBpi2CmIpMulticastMapTable"
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.SegmentPath = "docsIetfBpi2CmIpMulticastMapEntry" + types.AddKeyToken(docsIetfBpi2CmIpMulticastMapEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastIndex, "docsIetfBpi2CmIpMulticastIndex")
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmIpMulticastMapEntry.IfIndex})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastIndex", types.YLeaf{"DocsIetfBpi2CmIpMulticastIndex", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastIndex})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastAddressType", types.YLeaf{"DocsIetfBpi2CmIpMulticastAddressType", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastAddressType})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastAddress", types.YLeaf{"DocsIetfBpi2CmIpMulticastAddress", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastAddress})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAId", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAId", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAId})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapState", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapState", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapState})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapRequests", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapRequests", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapRequests})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapReplies", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapReplies", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapReplies})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapRejects", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapRejects", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapRejects})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode})
    docsIetfBpi2CmIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmIpMulticastSAMapRejectErrorString", types.YLeaf{"DocsIetfBpi2CmIpMulticastSAMapRejectErrorString", docsIetfBpi2CmIpMulticastMapEntry.DocsIetfBpi2CmIpMulticastSAMapRejectErrorString})

    docsIetfBpi2CmIpMulticastMapEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmIpMulticastIndex"}

    return &(docsIetfBpi2CmIpMulticastMapEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode represents message has been received since entry creation.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode_noAuthForRequestedDSFlow DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode = "noAuthForRequestedDSFlow"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode_dsFlowNotMappedToSA DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapRejectErrorCode = "dsFlowNotMappedToSA"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState represents Mapping FSM for this IP.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState_start DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState = "start"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState_mapWait DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState = "mapWait"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState_mapped DOCSIETFBPI2MIB_DocsIetfBpi2CmIpMulticastMapTable_DocsIetfBpi2CmIpMulticastMapEntry_DocsIetfBpi2CmIpMulticastSAMapState = "mapped"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable
// This table describes the Baseline Privacy Plus
// device certificates for each CM MAC interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the device certificates of one CM MAC interface.  An
    // entry in this table exists for each ifEntry with an ifType of
    // docsCableMaclayer(127). The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable_DocsIetfBpi2CmDeviceCertEntry.
    DocsIetfBpi2CmDeviceCertEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable_DocsIetfBpi2CmDeviceCertEntry
}

func (docsIetfBpi2CmDeviceCertTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmDeviceCertTable.EntityData.YFilter = docsIetfBpi2CmDeviceCertTable.YFilter
    docsIetfBpi2CmDeviceCertTable.EntityData.YangName = "docsIetfBpi2CmDeviceCertTable"
    docsIetfBpi2CmDeviceCertTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmDeviceCertTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmDeviceCertTable.EntityData.SegmentPath = "docsIetfBpi2CmDeviceCertTable"
    docsIetfBpi2CmDeviceCertTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmDeviceCertTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmDeviceCertTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmDeviceCertTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmDeviceCertTable.EntityData.Children.Append("docsIetfBpi2CmDeviceCertEntry", types.YChild{"DocsIetfBpi2CmDeviceCertEntry", nil})
    for i := range docsIetfBpi2CmDeviceCertTable.DocsIetfBpi2CmDeviceCertEntry {
        docsIetfBpi2CmDeviceCertTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmDeviceCertTable.DocsIetfBpi2CmDeviceCertEntry[i]), types.YChild{"DocsIetfBpi2CmDeviceCertEntry", docsIetfBpi2CmDeviceCertTable.DocsIetfBpi2CmDeviceCertEntry[i]})
    }
    docsIetfBpi2CmDeviceCertTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmDeviceCertTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmDeviceCertTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable_DocsIetfBpi2CmDeviceCertEntry
// Each entry contains the device certificates of
// one CM MAC interface.  An entry in this table exists for
// each ifEntry with an ifType of docsCableMaclayer(127).
type DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable_DocsIetfBpi2CmDeviceCertEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The X509 DER-encoded cable modem certificate. Note:  This object can be set
    // only when the value is the zero-length OCTET STRING, otherwise an error
    // 'inconsistentValue' is returned.  Once the object contains  the
    // certificate, its access MUST be read-only and persists after
    // re-initialization of the managed system. The type is string with length:
    // 0..4096.
    DocsIetfBpi2CmDeviceCmCert interface{}

    // The X509 DER-encoded manufacturer certificate which signed the cable modem
    // certificate. The type is string with length: 0..4096.
    DocsIetfBpi2CmDeviceManufCert interface{}
}

func (docsIetfBpi2CmDeviceCertEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmDeviceCertTable_DocsIetfBpi2CmDeviceCertEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmDeviceCertEntry.EntityData.YFilter = docsIetfBpi2CmDeviceCertEntry.YFilter
    docsIetfBpi2CmDeviceCertEntry.EntityData.YangName = "docsIetfBpi2CmDeviceCertEntry"
    docsIetfBpi2CmDeviceCertEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmDeviceCertEntry.EntityData.ParentYangName = "docsIetfBpi2CmDeviceCertTable"
    docsIetfBpi2CmDeviceCertEntry.EntityData.SegmentPath = "docsIetfBpi2CmDeviceCertEntry" + types.AddKeyToken(docsIetfBpi2CmDeviceCertEntry.IfIndex, "ifIndex")
    docsIetfBpi2CmDeviceCertEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmDeviceCertEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmDeviceCertEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmDeviceCertEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmDeviceCertEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmDeviceCertEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmDeviceCertEntry.IfIndex})
    docsIetfBpi2CmDeviceCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmDeviceCmCert", types.YLeaf{"DocsIetfBpi2CmDeviceCmCert", docsIetfBpi2CmDeviceCertEntry.DocsIetfBpi2CmDeviceCmCert})
    docsIetfBpi2CmDeviceCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmDeviceManufCert", types.YLeaf{"DocsIetfBpi2CmDeviceManufCert", docsIetfBpi2CmDeviceCertEntry.DocsIetfBpi2CmDeviceManufCert})

    docsIetfBpi2CmDeviceCertEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIetfBpi2CmDeviceCertEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable
// This table describes the Baseline Privacy Plus
// cryptographic suite capabilities for each CM MAC
// interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains a cryptographic suite pair which this CM MAC supports.
    // The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable_DocsIetfBpi2CmCryptoSuiteEntry.
    DocsIetfBpi2CmCryptoSuiteEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable_DocsIetfBpi2CmCryptoSuiteEntry
}

func (docsIetfBpi2CmCryptoSuiteTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmCryptoSuiteTable.EntityData.YFilter = docsIetfBpi2CmCryptoSuiteTable.YFilter
    docsIetfBpi2CmCryptoSuiteTable.EntityData.YangName = "docsIetfBpi2CmCryptoSuiteTable"
    docsIetfBpi2CmCryptoSuiteTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmCryptoSuiteTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmCryptoSuiteTable.EntityData.SegmentPath = "docsIetfBpi2CmCryptoSuiteTable"
    docsIetfBpi2CmCryptoSuiteTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmCryptoSuiteTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmCryptoSuiteTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmCryptoSuiteTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmCryptoSuiteTable.EntityData.Children.Append("docsIetfBpi2CmCryptoSuiteEntry", types.YChild{"DocsIetfBpi2CmCryptoSuiteEntry", nil})
    for i := range docsIetfBpi2CmCryptoSuiteTable.DocsIetfBpi2CmCryptoSuiteEntry {
        docsIetfBpi2CmCryptoSuiteTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmCryptoSuiteTable.DocsIetfBpi2CmCryptoSuiteEntry[i]), types.YChild{"DocsIetfBpi2CmCryptoSuiteEntry", docsIetfBpi2CmCryptoSuiteTable.DocsIetfBpi2CmCryptoSuiteEntry[i]})
    }
    docsIetfBpi2CmCryptoSuiteTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmCryptoSuiteTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmCryptoSuiteTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable_DocsIetfBpi2CmCryptoSuiteEntry
// Each entry contains a cryptographic suite pair
// which this CM MAC supports.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable_DocsIetfBpi2CmCryptoSuiteEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The index for a cryptographic suite row. The type
    // is interface{} with range: 1..1000.
    DocsIetfBpi2CmCryptoSuiteIndex interface{}

    // The value of this object is the data encryption algorithm for this
    // cryptographic suite capability. The type is DocsBpkmDataEncryptAlg.
    DocsIetfBpi2CmCryptoSuiteDataEncryptAlg interface{}

    // The value of this object is the data authentication algorithm for this
    // cryptographic suite capability. The type is DocsBpkmDataAuthentAlg.
    DocsIetfBpi2CmCryptoSuiteDataAuthentAlg interface{}
}

func (docsIetfBpi2CmCryptoSuiteEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmCryptoSuiteTable_DocsIetfBpi2CmCryptoSuiteEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.YFilter = docsIetfBpi2CmCryptoSuiteEntry.YFilter
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.YangName = "docsIetfBpi2CmCryptoSuiteEntry"
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.ParentYangName = "docsIetfBpi2CmCryptoSuiteTable"
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.SegmentPath = "docsIetfBpi2CmCryptoSuiteEntry" + types.AddKeyToken(docsIetfBpi2CmCryptoSuiteEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmCryptoSuiteEntry.DocsIetfBpi2CmCryptoSuiteIndex, "docsIetfBpi2CmCryptoSuiteIndex")
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmCryptoSuiteEntry.IfIndex})
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Leafs.Append("docsIetfBpi2CmCryptoSuiteIndex", types.YLeaf{"DocsIetfBpi2CmCryptoSuiteIndex", docsIetfBpi2CmCryptoSuiteEntry.DocsIetfBpi2CmCryptoSuiteIndex})
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Leafs.Append("docsIetfBpi2CmCryptoSuiteDataEncryptAlg", types.YLeaf{"DocsIetfBpi2CmCryptoSuiteDataEncryptAlg", docsIetfBpi2CmCryptoSuiteEntry.DocsIetfBpi2CmCryptoSuiteDataEncryptAlg})
    docsIetfBpi2CmCryptoSuiteEntry.EntityData.Leafs.Append("docsIetfBpi2CmCryptoSuiteDataAuthentAlg", types.YLeaf{"DocsIetfBpi2CmCryptoSuiteDataAuthentAlg", docsIetfBpi2CmCryptoSuiteEntry.DocsIetfBpi2CmCryptoSuiteDataAuthentAlg})

    docsIetfBpi2CmCryptoSuiteEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmCryptoSuiteIndex"}

    return &(docsIetfBpi2CmCryptoSuiteEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable
// This table describes the basic Baseline Privacy
// attributes of each CMTS MAC interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing attributes of one CMTS MAC
    // interface.  An entry in this table exists for each ifEntry with an ifType
    // of docsCableMaclayer(127). The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry.
    DocsIetfBpi2CmtsBaseEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry
}

func (docsIetfBpi2CmtsBaseTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsBaseTable.EntityData.YFilter = docsIetfBpi2CmtsBaseTable.YFilter
    docsIetfBpi2CmtsBaseTable.EntityData.YangName = "docsIetfBpi2CmtsBaseTable"
    docsIetfBpi2CmtsBaseTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsBaseTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsBaseTable.EntityData.SegmentPath = "docsIetfBpi2CmtsBaseTable"
    docsIetfBpi2CmtsBaseTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsBaseTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsBaseTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsBaseTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsBaseTable.EntityData.Children.Append("docsIetfBpi2CmtsBaseEntry", types.YChild{"DocsIetfBpi2CmtsBaseEntry", nil})
    for i := range docsIetfBpi2CmtsBaseTable.DocsIetfBpi2CmtsBaseEntry {
        docsIetfBpi2CmtsBaseTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsBaseTable.DocsIetfBpi2CmtsBaseEntry[i]), types.YChild{"DocsIetfBpi2CmtsBaseEntry", docsIetfBpi2CmtsBaseTable.DocsIetfBpi2CmtsBaseEntry[i]})
    }
    docsIetfBpi2CmtsBaseTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsBaseTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsBaseTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry
// Each entry contains objects describing attributes of
// one CMTS MAC interface.  An entry in this table exists for
// each ifEntry with an ifType of docsCableMaclayer(127).
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // The value of this object is the default lifetime, in seconds, the CMTS
    // assigns to a new authorization key. This object value persist after
    // re-initialization of the managed system. The type is interface{} with
    // range: 1..6048000. Units are seconds.
    DocsIetfBpi2CmtsDefaultAuthLifetime interface{}

    // The value of this object is the default lifetime, in seconds, the CMTS
    // assigns to a new Traffic Encryption Key (TEK). This object value persist
    // after re-initialization of the managed system. The type is interface{} with
    // range: 1..604800. Units are seconds.
    DocsIetfBpi2CmtsDefaultTEKLifetime interface{}

    // This object determines the default trust of self-signed  manufacturer
    // certificate entries, contained in docsIetfBpi2CmtsCACertTable, created
    // after setting this object. This object needs not to persist after
    // re-initialization  of the managed system. The type is
    // DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust.
    DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust interface{}

    // Setting this object to 'true' causes all chained and root certificates in
    // the chain to have their validity periods checked against the current time
    // of day, when the CMTS receives an Authorization Request from the CM. A
    // 'false' setting causes all certificates in the chain not to have their
    // validity periods checked against the current time of day. This object needs
    // not to persist after re-initialization  of the managed system. The type is
    // bool.
    DocsIetfBpi2CmtsCheckCertValidityPeriods interface{}

    // The value of this object is the count of times the CMTS has received an
    // Authentication Information message from any CM. Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    DocsIetfBpi2CmtsAuthentInfos interface{}

    // The value of this object is the count of times the CMTS has received an
    // Authorization Request message from any CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthRequests interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Reply message to any CM. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthReplies interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Reject message to any CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthRejects interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Invalid message to any CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthInvalids interface{}

    // The value of this object is the count of times the CMTS has received an SA
    // Map Request message from any CM. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsSAMapRequests interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // SA Map Reply message to any CM. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsSAMapReplies interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // SA Map Reject message to any CM. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsSAMapRejects interface{}
}

func (docsIetfBpi2CmtsBaseEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsBaseEntry.EntityData.YFilter = docsIetfBpi2CmtsBaseEntry.YFilter
    docsIetfBpi2CmtsBaseEntry.EntityData.YangName = "docsIetfBpi2CmtsBaseEntry"
    docsIetfBpi2CmtsBaseEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsBaseEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsBaseTable"
    docsIetfBpi2CmtsBaseEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsBaseEntry" + types.AddKeyToken(docsIetfBpi2CmtsBaseEntry.IfIndex, "ifIndex")
    docsIetfBpi2CmtsBaseEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsBaseEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsBaseEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsBaseEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmtsBaseEntry.IfIndex})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsDefaultAuthLifetime", types.YLeaf{"DocsIetfBpi2CmtsDefaultAuthLifetime", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsDefaultAuthLifetime})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsDefaultTEKLifetime", types.YLeaf{"DocsIetfBpi2CmtsDefaultTEKLifetime", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsDefaultTEKLifetime})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsDefaultSelfSignedManufCertTrust", types.YLeaf{"DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCheckCertValidityPeriods", types.YLeaf{"DocsIetfBpi2CmtsCheckCertValidityPeriods", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsCheckCertValidityPeriods})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthentInfos", types.YLeaf{"DocsIetfBpi2CmtsAuthentInfos", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsAuthentInfos})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthRequests", types.YLeaf{"DocsIetfBpi2CmtsAuthRequests", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsAuthRequests})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthReplies", types.YLeaf{"DocsIetfBpi2CmtsAuthReplies", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsAuthReplies})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthRejects", types.YLeaf{"DocsIetfBpi2CmtsAuthRejects", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsAuthRejects})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthInvalids", types.YLeaf{"DocsIetfBpi2CmtsAuthInvalids", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsAuthInvalids})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsSAMapRequests", types.YLeaf{"DocsIetfBpi2CmtsSAMapRequests", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsSAMapRequests})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsSAMapReplies", types.YLeaf{"DocsIetfBpi2CmtsSAMapReplies", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsSAMapReplies})
    docsIetfBpi2CmtsBaseEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsSAMapRejects", types.YLeaf{"DocsIetfBpi2CmtsSAMapRejects", docsIetfBpi2CmtsBaseEntry.DocsIetfBpi2CmtsSAMapRejects})

    docsIetfBpi2CmtsBaseEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIetfBpi2CmtsBaseEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust represents  of the managed system.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust_trusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust = "trusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust_untrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsBaseTable_DocsIetfBpi2CmtsBaseEntry_DocsIetfBpi2CmtsDefaultSelfSignedManufCertTrust = "untrusted"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable
// This table describes the attributes of each CM
// authorization association. The CMTS maintains one
// authorization association with each Baseline Privacy-
// enabled CM, registered on each CMTS MAC interface,
// regardless of whether the CM is authorized or rejected.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing attributes of one authorization
    // association. The CMTS MUST create one entry per CM per MAC interface, based
    // on the receipt of an Authorization Request message, and MUST not delete the
    // entry until the CM loses registration. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry.
    DocsIetfBpi2CmtsAuthEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry
}

func (docsIetfBpi2CmtsAuthTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsAuthTable.EntityData.YFilter = docsIetfBpi2CmtsAuthTable.YFilter
    docsIetfBpi2CmtsAuthTable.EntityData.YangName = "docsIetfBpi2CmtsAuthTable"
    docsIetfBpi2CmtsAuthTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsAuthTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsAuthTable.EntityData.SegmentPath = "docsIetfBpi2CmtsAuthTable"
    docsIetfBpi2CmtsAuthTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsAuthTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsAuthTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsAuthTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsAuthTable.EntityData.Children.Append("docsIetfBpi2CmtsAuthEntry", types.YChild{"DocsIetfBpi2CmtsAuthEntry", nil})
    for i := range docsIetfBpi2CmtsAuthTable.DocsIetfBpi2CmtsAuthEntry {
        docsIetfBpi2CmtsAuthTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsAuthTable.DocsIetfBpi2CmtsAuthEntry[i]), types.YChild{"DocsIetfBpi2CmtsAuthEntry", docsIetfBpi2CmtsAuthTable.DocsIetfBpi2CmtsAuthEntry[i]})
    }
    docsIetfBpi2CmtsAuthTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsAuthTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsAuthTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry
// Each entry contains objects describing attributes of
// one authorization association. The CMTS MUST create one
// entry per CM per MAC interface, based on the receipt of an
// Authorization Request message, and MUST not delete the
// entry until the CM loses registration.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The value of this object is the physical address
    // of the CM to which the authorization association applies. The type is
    // string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIetfBpi2CmtsAuthCmMacAddress interface{}

    // The value of this object is the version of Baseline Privacy for which this
    // CM has registered. The value 'bpiplus' represents the value of BPI-Version
    // Attribute of the Baseline Privacy Key Management BPKM attribute BPI-Version
    // (1). The value 'bpi' is used to represent the CM registered using DOCSIS
    // 1.0 Baseline Privacy. The type is DocsIetfBpi2CmtsAuthCmBpiVersion.
    DocsIetfBpi2CmtsAuthCmBpiVersion interface{}

    // The value of this object is a DER-encoded RSAPublicKey ASN.1 type string,
    // as defined in the RSA Encryption Standard (PKCS #1), corresponding to the
    // public key of the CM.  This is the zero-length OCTET STRING if the CMTS
    // does not retain the public key. The type is string with length: 0..524.
    DocsIetfBpi2CmtsAuthCmPublicKey interface{}

    // The value of this object is the most recent authorization key sequence
    // number for this CM. The type is interface{} with range: 0..15.
    DocsIetfBpi2CmtsAuthCmKeySequenceNumber interface{}

    // The value of this object is the actual clock time for expiration of the
    // immediate predecessor of the most recent authorization key for this FSM. If
    // this FSM has only one authorization key, then the value is the time of
    // activation of this FSM. Note: This object has no meaning for CMs running in
    // BPI mode, therefore this object is not instantiated for entries associated
    // to those CMs. The type is string.
    DocsIetfBpi2CmtsAuthCmExpiresOld interface{}

    // The value of this object is the actual clock time for expiration of the
    // most recent authorization key for this FSM. The type is string.
    DocsIetfBpi2CmtsAuthCmExpiresNew interface{}

    // The value of this object is the lifetime, in seconds, the CMTS assigns to
    // an authorization key for this CM. The type is interface{} with range:
    // 1..6048000. Units are seconds.
    DocsIetfBpi2CmtsAuthCmLifetime interface{}

    // Setting this object to invalidateAuth(2) causes the CMTS to invalidate the
    // current CM authorization key(s), but not to transmit an Authorization
    // Invalid message nor to invalidate the primary SAID's TEKs.  Setting this
    // object to sendAuthInvalid(3) causes the CMTS to invalidate the current CM
    // authorization key(s), and to transmit an Authorization Invalid message to
    // the CM, but not to invalidate the primary SAID's TEKs.  Setting this object
    // to invalidateTeks(4) causes the CMTS to invalidate the current CM
    // authorization key(s), to transmit an Authorization Invalid message to the
    // CM, and to invalidate the TEKs associated with this CM's primary SAID. For
    // BPI mode, substitute all of the CM's unicast TEK(s) for the primary SAID's
    // TEKs in the previous paragraph. Reading this object returns the most
    // recently set value of this object, or returns noResetRequested(1) if the
    // object has not been set since entry creation. The type is
    // DocsIetfBpi2CmtsAuthCmReset.
    DocsIetfBpi2CmtsAuthCmReset interface{}

    // The value of this object is the count of times the CMTS has received an
    // Authentication Information message from this CM. Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime. The type is interface{} with range:
    // 0..4294967295.
    DocsIetfBpi2CmtsAuthCmInfos interface{}

    // The value of this object is the count of times the CMTS has received an
    // Authorization Request message from this CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthCmRequests interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Reply message to this CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthCmReplies interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Reject message to this CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthCmRejects interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // Authorization Invalid message to this CM. Discontinuities in the value of
    // this counter can occur at re-initialization of the management system, and
    // at other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthCmInvalids interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent Authorization Reject message transmitted to the CM.  This has
    // value unknown(2) if the last Error-Code value was 0, and none(1) if no
    // Authorization Reject message has been transmitted to the CM, since entry
    // creation. The type is DocsIetfBpi2CmtsAuthRejectErrorCode.
    DocsIetfBpi2CmtsAuthRejectErrorCode interface{}

    // The value of this object is the text string in most recent Authorization
    // Reject message transmitted to the CM.  This is a zero length string if no
    // Authorization Reject message has been transmitted to the CM, since entry
    // creation. The type is string with length: 0..128.
    DocsIetfBpi2CmtsAuthRejectErrorString interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // most recent Authorization Invalid message transmitted to the CM.  This has
    // value unknown(2) if the last Error-Code value was 0, and none(1) if no
    // Authorization Invalid message has been transmitted to the CM since entry
    // creation. The type is DocsIetfBpi2CmtsAuthInvalidErrorCode.
    DocsIetfBpi2CmtsAuthInvalidErrorCode interface{}

    // The value of this object is the text string in most recent Authorization
    // Invalid message transmitted to the CM.  This is a zero length string if no
    // Authorization Invalid message has been transmitted to the CM since entry
    // creation. The type is string with length: 0..128.
    DocsIetfBpi2CmtsAuthInvalidErrorString interface{}

    // The value of this object is the Primary Security Association identifier. 
    // For BPI mode, the value must be any unicast SID. The type is interface{}
    // with range: 0..16383.
    DocsIetfBpi2CmtsAuthPrimarySAId interface{}

    // Contains the reason why a CM's certificate is deemed valid or invalid.
    // Return unknown(0) if the CM is running BPI mode. ValidCmChained(1) means
    // the certificate is valid    because it chains to a valid certificate.
    // ValidCmTrusted(2) means the certificate is valid    because it has been
    // provisioned (in the    docsIetfBpi2CmtsProvisionedCmCert table) to be
    // trusted. InvalidCmUntrusted(3) means the certificate is invalid     
    // because it has been provisioned (in the     
    // docsIetfBpi2CmtsProvisionedCmCert table) to be untrusted.
    // InvalidCAUntrusted(4) means the certificate is invalid      because it
    // chains to an untrusted certificate. InvalidCmOther(5) and InvalidCAOther(6)
    // refer to      errors in parsing, validity periods, etc, which are     
    // attributable to the CM certificate or its chain      respectively;
    // additional information may be found      in
    // docsIetfBpi2AuthRejectErrorString for these types      of errors.
    // InvalidCmRevoked(7) means the certificate is    invalid as it was marked as
    // revoked.  InvalidCARevoked(8) means the CA certificate is    invalid as it
    // was marked as revoked. The type is DocsIetfBpi2CmtsAuthBpkmCmCertValid.
    DocsIetfBpi2CmtsAuthBpkmCmCertValid interface{}

    // The X509 CM Certificate sent as part of a BPKM Authorization Request. Note:
    // The zero-length OCTET STRING must be returned if the Entire certificate is
    // not retained in the CMTS. The type is string with length: 0..4096.
    DocsIetfBpi2CmtsAuthBpkmCmCert interface{}

    // A row index into docsIetfBpi2CmtsCACertTable. Returns the index in
    // docsIetfBpi2CmtsCACertTable which CA certificate this CM is chained to.  A
    // value of 0 means it could not be found or not applicable. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsAuthCACertIndexPtr interface{}
}

func (docsIetfBpi2CmtsAuthEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsAuthEntry.EntityData.YFilter = docsIetfBpi2CmtsAuthEntry.YFilter
    docsIetfBpi2CmtsAuthEntry.EntityData.YangName = "docsIetfBpi2CmtsAuthEntry"
    docsIetfBpi2CmtsAuthEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsAuthEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsAuthTable"
    docsIetfBpi2CmtsAuthEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsAuthEntry" + types.AddKeyToken(docsIetfBpi2CmtsAuthEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmMacAddress, "docsIetfBpi2CmtsAuthCmMacAddress")
    docsIetfBpi2CmtsAuthEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsAuthEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsAuthEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsAuthEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmtsAuthEntry.IfIndex})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmMacAddress", types.YLeaf{"DocsIetfBpi2CmtsAuthCmMacAddress", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmMacAddress})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmBpiVersion", types.YLeaf{"DocsIetfBpi2CmtsAuthCmBpiVersion", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmBpiVersion})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmPublicKey", types.YLeaf{"DocsIetfBpi2CmtsAuthCmPublicKey", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmPublicKey})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmKeySequenceNumber", types.YLeaf{"DocsIetfBpi2CmtsAuthCmKeySequenceNumber", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmKeySequenceNumber})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmExpiresOld", types.YLeaf{"DocsIetfBpi2CmtsAuthCmExpiresOld", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmExpiresOld})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmExpiresNew", types.YLeaf{"DocsIetfBpi2CmtsAuthCmExpiresNew", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmExpiresNew})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmLifetime", types.YLeaf{"DocsIetfBpi2CmtsAuthCmLifetime", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmLifetime})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmReset", types.YLeaf{"DocsIetfBpi2CmtsAuthCmReset", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmReset})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmInfos", types.YLeaf{"DocsIetfBpi2CmtsAuthCmInfos", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmInfos})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmRequests", types.YLeaf{"DocsIetfBpi2CmtsAuthCmRequests", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmRequests})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmReplies", types.YLeaf{"DocsIetfBpi2CmtsAuthCmReplies", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmReplies})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmRejects", types.YLeaf{"DocsIetfBpi2CmtsAuthCmRejects", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmRejects})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCmInvalids", types.YLeaf{"DocsIetfBpi2CmtsAuthCmInvalids", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCmInvalids})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmtsAuthRejectErrorCode", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthRejectErrorCode})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthRejectErrorString", types.YLeaf{"DocsIetfBpi2CmtsAuthRejectErrorString", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthRejectErrorString})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthInvalidErrorCode", types.YLeaf{"DocsIetfBpi2CmtsAuthInvalidErrorCode", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthInvalidErrorCode})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthInvalidErrorString", types.YLeaf{"DocsIetfBpi2CmtsAuthInvalidErrorString", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthInvalidErrorString})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthPrimarySAId", types.YLeaf{"DocsIetfBpi2CmtsAuthPrimarySAId", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthPrimarySAId})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthBpkmCmCertValid", types.YLeaf{"DocsIetfBpi2CmtsAuthBpkmCmCertValid", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthBpkmCmCertValid})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthBpkmCmCert", types.YLeaf{"DocsIetfBpi2CmtsAuthBpkmCmCert", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthBpkmCmCert})
    docsIetfBpi2CmtsAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsAuthCACertIndexPtr", types.YLeaf{"DocsIetfBpi2CmtsAuthCACertIndexPtr", docsIetfBpi2CmtsAuthEntry.DocsIetfBpi2CmtsAuthCACertIndexPtr})

    docsIetfBpi2CmtsAuthEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmtsAuthCmMacAddress"}

    return &(docsIetfBpi2CmtsAuthEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid represents    invalid as it was marked as revoked.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_validCmChained DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "validCmChained"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_validCmTrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "validCmTrusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCmUntrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCmUntrusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCAUntrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCAUntrusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCmOther DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCmOther"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCAOther DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCAOther"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCmRevoked DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCmRevoked"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid_invalidCARevoked DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthBpkmCmCertValid = "invalidCARevoked"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion represents CM registered using DOCSIS 1.0 Baseline Privacy.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion_bpi DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion = "bpi"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion_bpiPlus DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmBpiVersion = "bpiPlus"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset represents object has not been set since entry creation.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset_noResetRequested DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset = "noResetRequested"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset_invalidateAuth DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset = "invalidateAuth"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset_sendAuthInvalid DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset = "sendAuthInvalid"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset_invalidateTeks DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthCmReset = "invalidateTeks"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode represents the CM since entry creation.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_unauthorizedCm DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "unauthorizedCm"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_unsolicited DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "unsolicited"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_invalidKeySequence DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "invalidKeySequence"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode_keyRequestAuthenticationFailure DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthInvalidErrorCode = "keyRequestAuthenticationFailure"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode represents the CM, since entry creation.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_unauthorizedCm DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "unauthorizedCm"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_unauthorizedSaid DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "unauthorizedSaid"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_permanentAuthorizationFailure DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "permanentAuthorizationFailure"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode_timeOfDayNotAcquired DOCSIETFBPI2MIB_DocsIetfBpi2CmtsAuthTable_DocsIetfBpi2CmtsAuthEntry_DocsIetfBpi2CmtsAuthRejectErrorCode = "timeOfDayNotAcquired"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable
// This table describes the attributes of each
// Traffic Encryption Key (TEK) association. The CMTS
// Maintains one TEK association per SAID on each CMTS MAC
// interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing attributes of one TEK association on
    // a particular CMTS MAC interface. The CMTS MUST create one entry per SAID
    // per MAC interface, based on the receipt of a Key Request message, and MUST
    // not delete the entry before the CM authorization for the SAID permanently
    // expires. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry.
    DocsIetfBpi2CmtsTEKEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry
}

func (docsIetfBpi2CmtsTEKTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsTEKTable.EntityData.YFilter = docsIetfBpi2CmtsTEKTable.YFilter
    docsIetfBpi2CmtsTEKTable.EntityData.YangName = "docsIetfBpi2CmtsTEKTable"
    docsIetfBpi2CmtsTEKTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsTEKTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsTEKTable.EntityData.SegmentPath = "docsIetfBpi2CmtsTEKTable"
    docsIetfBpi2CmtsTEKTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsTEKTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsTEKTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsTEKTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsTEKTable.EntityData.Children.Append("docsIetfBpi2CmtsTEKEntry", types.YChild{"DocsIetfBpi2CmtsTEKEntry", nil})
    for i := range docsIetfBpi2CmtsTEKTable.DocsIetfBpi2CmtsTEKEntry {
        docsIetfBpi2CmtsTEKTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsTEKTable.DocsIetfBpi2CmtsTEKEntry[i]), types.YChild{"DocsIetfBpi2CmtsTEKEntry", docsIetfBpi2CmtsTEKTable.DocsIetfBpi2CmtsTEKEntry[i]})
    }
    docsIetfBpi2CmtsTEKTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsTEKTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsTEKTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry
// Each entry contains objects describing attributes of
// one TEK association on a particular CMTS MAC interface. The
// CMTS MUST create one entry per SAID per MAC interface,
// based on the receipt of a Key Request message, and MUST not
// delete the entry before the CM authorization for the SAID
// permanently expires.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The value of this object is the DOCSIS Security
    // Association ID (SAID). The type is interface{} with range: 1..16383.
    DocsIetfBpi2CmtsTEKSAId interface{}

    // The value of this object is the type of security association.  'dynamic'
    // does not apply to CMs running in BPI mode.  Unicast BPI TEKs must utilize
    // the 'primary' encoding and multicast BPI TEKs must utilize the 'static'
    // encoding. The type is DocsBpkmSAType.
    DocsIetfBpi2CmtsTEKSAType interface{}

    // The value of this object is the data encryption algorithm for this SAID.
    // The type is DocsBpkmDataEncryptAlg.
    DocsIetfBpi2CmtsTEKDataEncryptAlg interface{}

    // The value of this object is the data authentication algorithm for this
    // SAID. The type is DocsBpkmDataAuthentAlg.
    DocsIetfBpi2CmtsTEKDataAuthentAlg interface{}

    // The value of this object is the lifetime, in seconds, the CMTS assigns to
    // keys for this TEK association. The type is interface{} with range:
    // 1..604800. Units are seconds.
    DocsIetfBpi2CmtsTEKLifetime interface{}

    // The value of this object is the most recent TEK key sequence number for
    // this SAID. The type is interface{} with range: 0..15.
    DocsIetfBpi2CmtsTEKKeySequenceNumber interface{}

    // The value of this object is the actual clock time for expiration of the
    // immediate predecessor of the most recent TEK for this FSM. If this FSM has
    // only one TEK, then the value is the time of activation of this FSM. The
    // type is string.
    DocsIetfBpi2CmtsTEKExpiresOld interface{}

    // The value of this object is the actual clock time for expiration of the
    // most recent TEK for this FSM. The type is string.
    DocsIetfBpi2CmtsTEKExpiresNew interface{}

    // Setting this object to 'true' causes the CMTS to invalidate all currently
    // active TEK(s) and to generate new TEK(s) for the associated SAID; the CMTS
    // MAY also generate unsolicited TEK Invalid message(s), to optimize the TEK
    // synchronization between the CMTS and the CM(s). Reading this object always
    // returns FALSE. The type is bool.
    DocsIetfBpi2CmtsTEKReset interface{}

    // The value of this object is the count of times the CMTS has received a Key
    // Request message. Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime. The type is interface{} with
    // range: 0..4294967295.
    DocsIetfBpi2CmtsKeyRequests interface{}

    // The value of this object is the count of times the CMTS has transmitted a
    // Key Reply message. Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsKeyReplies interface{}

    // The value of this object is the count of times the CMTS has transmitted a
    // Key Reject message. Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsKeyRejects interface{}

    // The value of this object is the count of times the CMTS has transmitted a
    // TEK Invalid message. Discontinuities in the value of this counter can occur
    // at re-initialization of the management system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime. The type is
    // interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsTEKInvalids interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // the most recent Key Reject message sent in response to a Key Request for
    // this SAID. This has value unknown(2) if the last Error-Code value was 0,
    // and none(1) if no Key Reject message has been received since registration.
    // The type is DocsIetfBpi2CmtsKeyRejectErrorCode.
    DocsIetfBpi2CmtsKeyRejectErrorCode interface{}

    // The value of this object is the text string in the most recent Key Reject
    // message sent in response to a Key Request for this SAID. This is a zero
    // length string if no Key Reject message has been received since
    // registration. The type is string with length: 0..128.
    DocsIetfBpi2CmtsKeyRejectErrorString interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // the most recent TEK Invalid message sent in association with this SAID. 
    // This has value unknown(2) if the last Error-Code value was 0, and none(1)
    // if no TEK Invalid message has been received since registration. The type is
    // DocsIetfBpi2CmtsTEKInvalidErrorCode.
    DocsIetfBpi2CmtsTEKInvalidErrorCode interface{}

    // The value of this object is the text string in the most recent TEK Invalid
    // message sent in association with this SAID.  This is a zero length string
    // if no TEK Invalid message has been received since registration. The type is
    // string with length: 0..128.
    DocsIetfBpi2CmtsTEKInvalidErrorString interface{}
}

func (docsIetfBpi2CmtsTEKEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsTEKEntry.EntityData.YFilter = docsIetfBpi2CmtsTEKEntry.YFilter
    docsIetfBpi2CmtsTEKEntry.EntityData.YangName = "docsIetfBpi2CmtsTEKEntry"
    docsIetfBpi2CmtsTEKEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsTEKEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsTEKTable"
    docsIetfBpi2CmtsTEKEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsTEKEntry" + types.AddKeyToken(docsIetfBpi2CmtsTEKEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKSAId, "docsIetfBpi2CmtsTEKSAId")
    docsIetfBpi2CmtsTEKEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsTEKEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsTEKEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsTEKEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmtsTEKEntry.IfIndex})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKSAId", types.YLeaf{"DocsIetfBpi2CmtsTEKSAId", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKSAId})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKSAType", types.YLeaf{"DocsIetfBpi2CmtsTEKSAType", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKSAType})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKDataEncryptAlg", types.YLeaf{"DocsIetfBpi2CmtsTEKDataEncryptAlg", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKDataEncryptAlg})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKDataAuthentAlg", types.YLeaf{"DocsIetfBpi2CmtsTEKDataAuthentAlg", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKDataAuthentAlg})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKLifetime", types.YLeaf{"DocsIetfBpi2CmtsTEKLifetime", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKLifetime})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKKeySequenceNumber", types.YLeaf{"DocsIetfBpi2CmtsTEKKeySequenceNumber", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKKeySequenceNumber})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKExpiresOld", types.YLeaf{"DocsIetfBpi2CmtsTEKExpiresOld", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKExpiresOld})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKExpiresNew", types.YLeaf{"DocsIetfBpi2CmtsTEKExpiresNew", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKExpiresNew})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKReset", types.YLeaf{"DocsIetfBpi2CmtsTEKReset", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKReset})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsKeyRequests", types.YLeaf{"DocsIetfBpi2CmtsKeyRequests", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsKeyRequests})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsKeyReplies", types.YLeaf{"DocsIetfBpi2CmtsKeyReplies", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsKeyReplies})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsKeyRejects", types.YLeaf{"DocsIetfBpi2CmtsKeyRejects", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsKeyRejects})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKInvalids", types.YLeaf{"DocsIetfBpi2CmtsTEKInvalids", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKInvalids})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsKeyRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmtsKeyRejectErrorCode", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsKeyRejectErrorCode})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsKeyRejectErrorString", types.YLeaf{"DocsIetfBpi2CmtsKeyRejectErrorString", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsKeyRejectErrorString})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKInvalidErrorCode", types.YLeaf{"DocsIetfBpi2CmtsTEKInvalidErrorCode", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKInvalidErrorCode})
    docsIetfBpi2CmtsTEKEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsTEKInvalidErrorString", types.YLeaf{"DocsIetfBpi2CmtsTEKInvalidErrorString", docsIetfBpi2CmtsTEKEntry.DocsIetfBpi2CmtsTEKInvalidErrorString})

    docsIetfBpi2CmtsTEKEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmtsTEKSAId"}

    return &(docsIetfBpi2CmtsTEKEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode represents received since registration.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode_unauthorizedSaid DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsKeyRejectErrorCode = "unauthorizedSaid"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode represents since registration.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode_invalidKeySequence DOCSIETFBPI2MIB_DocsIetfBpi2CmtsTEKTable_DocsIetfBpi2CmtsTEKEntry_DocsIetfBpi2CmtsTEKInvalidErrorCode = "invalidKeySequence"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable
// This table maps multicast IP addresses to SAIDs.
// If a multicast IP address is mapped by multiple rows
// in the table, the row with the lowest
// docsIetfBpi2CmtsIpMulticastIndex must be utilized for the
// mapping.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing the mapping of a set of multicast IP
    // address and mask to one SAID associated to a CMTS MAC Interface, as well as
    // associated message counters and error information. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry.
    DocsIetfBpi2CmtsIpMulticastMapEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry
}

func (docsIetfBpi2CmtsIpMulticastMapTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.YFilter = docsIetfBpi2CmtsIpMulticastMapTable.YFilter
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.YangName = "docsIetfBpi2CmtsIpMulticastMapTable"
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.SegmentPath = "docsIetfBpi2CmtsIpMulticastMapTable"
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.Children.Append("docsIetfBpi2CmtsIpMulticastMapEntry", types.YChild{"DocsIetfBpi2CmtsIpMulticastMapEntry", nil})
    for i := range docsIetfBpi2CmtsIpMulticastMapTable.DocsIetfBpi2CmtsIpMulticastMapEntry {
        docsIetfBpi2CmtsIpMulticastMapTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsIpMulticastMapTable.DocsIetfBpi2CmtsIpMulticastMapEntry[i]), types.YChild{"DocsIetfBpi2CmtsIpMulticastMapEntry", docsIetfBpi2CmtsIpMulticastMapTable.DocsIetfBpi2CmtsIpMulticastMapEntry[i]})
    }
    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsIpMulticastMapTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsIpMulticastMapTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry
// Each entry contains objects describing the mapping of
// a set of multicast IP address and mask to one SAID
// associated to a CMTS MAC Interface, as well as associated
// message
// counters and error information.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The index of this row. Conceptual rows having the
    // value 'permanent' need not allow write-access to any columnar objects in
    // the row. The type is interface{} with range: 1..4294967295.
    DocsIetfBpi2CmtsIpMulticastIndex interface{}

    // The type of internet address for docsIetfBpi2CmtsIpMulticastAddress and
    // docsIetfBpi2CmtsIpMulticastMask. The type is InetAddressType.
    DocsIetfBpi2CmtsIpMulticastAddressType interface{}

    // This object represents the IP multicast address to be mapped, in
    // conjunction with docsIetfBpi2CmtsIpMulticastMask. The type of this address
    // is determined by the value of the object
    // docsIetfBpi2CmtsIpMulticastAddressType. The type is string with length:
    // 0..255.
    DocsIetfBpi2CmtsIpMulticastAddress interface{}

    // This object represents the IP multicast address mask for this row. An IP
    // multicast address matches this row if the logical AND of the address with
    // docsIetfBpi2CmtsIpMulticastMask is identical to the logical AND of
    // docsIetfBpi2CmtsIpMulticastAddr with docsIetfBpi2CmtsIpMulticastMask. The
    // type of this address is determined by the value of the object
    // docsIetfBpi2CmtsIpMulticastAddressType.  Note: For IPv6 this object needs
    // not to represent a  contiguous netmask, e.g. to associate an SAID to a 
    // multicast group matching 'any' multicast scope. The TC 
    // InetAddressPrefixLength is not used because it only  represents contiguous
    // netmask. The type is string with length: 0..255.
    DocsIetfBpi2CmtsIpMulticastMask interface{}

    // This object represents the multicast SAID to be used in this IP multicast
    // address mapping entry. The type is interface{} with range: 0..16383.
    DocsIetfBpi2CmtsIpMulticastSAId interface{}

    // The value of this object is the type of security association.  'dynamic'
    // does not apply to CMs running in BPI mode.  Unicast BPI TEKs must utilize
    // the 'primary' encoding and multicast BPI TEKs must utilize the 'static'
    // encoding.  SNMP created entries set this object by default to 'static' if
    // not set at row creation. The type is DocsBpkmSAType.
    DocsIetfBpi2CmtsIpMulticastSAType interface{}

    // The value of this object is the data encryption algorithm for this IP. The
    // type is DocsBpkmDataEncryptAlg.
    DocsIetfBpi2CmtsIpMulticastDataEncryptAlg interface{}

    // The value of this object is the data authentication algorithm for this IP.
    // The type is DocsBpkmDataAuthentAlg.
    DocsIetfBpi2CmtsIpMulticastDataAuthentAlg interface{}

    // The value of this object is the count of times the CMTS has received an SA
    // Map Request message for this IP. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsIpMulticastSAMapRequests interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // SA Map Reply message for this IP. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsIpMulticastSAMapReplies interface{}

    // The value of this object is the count of times the CMTS has transmitted an
    // SA Map Reject message for this IP. Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of ifCounterDiscontinuityTime. The
    // type is interface{} with range: 0..4294967295.
    DocsIetfBpi2CmtsIpMulticastSAMapRejects interface{}

    // The value of this object is the enumerated description of the Error-Code in
    // the most recent SA Map Reject message sent in response to a SA Map Request
    // for This IP.  It has value unknown(2) if the last Error-Code Value was 0,
    // and none(1) if no SA MAP Reject message has been received since entry
    // creation. The type is DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode.
    DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode interface{}

    // The value of this object is the text string in the most recent SA Map
    // Reject message sent in response to an SA Map Request for this IP.  It is a
    // zero length string if no SA Map Reject message has been received since
    // entry creation. The type is string with length: 0..128.
    DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorString interface{}

    // This object controls and reflects the IP multicast address mapping entry. 
    // There is no restriction on the ability to change values in this row while
    // the row is active. A created row can be set to active only after the
    // Corresponding instances of docsIetfBpi2CmtsIpMulticastAddress,
    // docsIetfBpi2CmtsIpMulticastMask, docsIetfBpi2CmtsIpMulticastSAId and
    // docsIetfBpi2CmtsIpMulticastSAType have all been set. The type is RowStatus.
    DocsIetfBpi2CmtsIpMulticastMapControl interface{}

    // The storage type for this conceptual row. Conceptual rows having the value
    // 'permanent' need not allow write-access to any columnar objects in the row.
    // The type is StorageType.
    DocsIetfBpi2CmtsIpMulticastMapStorageType interface{}
}

func (docsIetfBpi2CmtsIpMulticastMapEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.YFilter = docsIetfBpi2CmtsIpMulticastMapEntry.YFilter
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.YangName = "docsIetfBpi2CmtsIpMulticastMapEntry"
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsIpMulticastMapTable"
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsIpMulticastMapEntry" + types.AddKeyToken(docsIetfBpi2CmtsIpMulticastMapEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastIndex, "docsIetfBpi2CmtsIpMulticastIndex")
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmtsIpMulticastMapEntry.IfIndex})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastIndex", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastIndex", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastIndex})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastAddressType", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastAddressType", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastAddressType})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastAddress", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastAddress", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastAddress})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastMask", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastMask", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastMask})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAId", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAId", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAId})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAType", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAType", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAType})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastDataEncryptAlg", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastDataEncryptAlg", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastDataEncryptAlg})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastDataAuthentAlg", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastDataAuthentAlg", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastDataAuthentAlg})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAMapRequests", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAMapRequests", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAMapRequests})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAMapReplies", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAMapReplies", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAMapReplies})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAMapRejects", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAMapRejects", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAMapRejects})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastSAMapRejectErrorString", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorString", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorString})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastMapControl", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastMapControl", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastMapControl})
    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsIpMulticastMapStorageType", types.YLeaf{"DocsIetfBpi2CmtsIpMulticastMapStorageType", docsIetfBpi2CmtsIpMulticastMapEntry.DocsIetfBpi2CmtsIpMulticastMapStorageType})

    docsIetfBpi2CmtsIpMulticastMapEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmtsIpMulticastIndex"}

    return &(docsIetfBpi2CmtsIpMulticastMapEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode represents been received since entry creation.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode_none DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode = "none"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode_unknown DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode = "unknown"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode_noAuthForRequestedDSFlow DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode = "noAuthForRequestedDSFlow"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode_dsFlowNotMappedToSA DOCSIETFBPI2MIB_DocsIetfBpi2CmtsIpMulticastMapTable_DocsIetfBpi2CmtsIpMulticastMapEntry_DocsIetfBpi2CmtsIpMulticastSAMapRejectErrorCode = "dsFlowNotMappedToSA"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable
// This table describes the multicast SAID
// authorization for each CM on each CMTS MAC interface.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains objects describing the key authorization of one cable
    // modem for one multicast SAID for one CMTS MAC interface. Row entries
    // persist after re-initialization of the managed system. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable_DocsIetfBpi2CmtsMulticastAuthEntry.
    DocsIetfBpi2CmtsMulticastAuthEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable_DocsIetfBpi2CmtsMulticastAuthEntry
}

func (docsIetfBpi2CmtsMulticastAuthTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.YFilter = docsIetfBpi2CmtsMulticastAuthTable.YFilter
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.YangName = "docsIetfBpi2CmtsMulticastAuthTable"
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.SegmentPath = "docsIetfBpi2CmtsMulticastAuthTable"
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsMulticastAuthTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.Children.Append("docsIetfBpi2CmtsMulticastAuthEntry", types.YChild{"DocsIetfBpi2CmtsMulticastAuthEntry", nil})
    for i := range docsIetfBpi2CmtsMulticastAuthTable.DocsIetfBpi2CmtsMulticastAuthEntry {
        docsIetfBpi2CmtsMulticastAuthTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsMulticastAuthTable.DocsIetfBpi2CmtsMulticastAuthEntry[i]), types.YChild{"DocsIetfBpi2CmtsMulticastAuthEntry", docsIetfBpi2CmtsMulticastAuthTable.DocsIetfBpi2CmtsMulticastAuthEntry[i]})
    }
    docsIetfBpi2CmtsMulticastAuthTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsMulticastAuthTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsMulticastAuthTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable_DocsIetfBpi2CmtsMulticastAuthEntry
// Each entry contains objects describing the key
// authorization of one cable modem for one multicast SAID
// for one CMTS MAC interface.
// Row entries persist after re-initialization of
// the managed system.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable_DocsIetfBpi2CmtsMulticastAuthEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This object represents the multicast SAID for
    // authorization. The type is interface{} with range: 1..16383.
    DocsIetfBpi2CmtsMulticastAuthSAId interface{}

    // This attribute is a key. This object represents the MAC address of the CM
    // to which the multicast SAID authorization applies. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIetfBpi2CmtsMulticastAuthCmMacAddress interface{}

    // The status of this conceptual row for the authorization of  multicast SAIDs
    // to CMs. . The type is RowStatus.
    DocsIetfBpi2CmtsMulticastAuthControl interface{}
}

func (docsIetfBpi2CmtsMulticastAuthEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsMulticastAuthTable_DocsIetfBpi2CmtsMulticastAuthEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.YFilter = docsIetfBpi2CmtsMulticastAuthEntry.YFilter
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.YangName = "docsIetfBpi2CmtsMulticastAuthEntry"
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsMulticastAuthTable"
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsMulticastAuthEntry" + types.AddKeyToken(docsIetfBpi2CmtsMulticastAuthEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIetfBpi2CmtsMulticastAuthEntry.DocsIetfBpi2CmtsMulticastAuthSAId, "docsIetfBpi2CmtsMulticastAuthSAId") + types.AddKeyToken(docsIetfBpi2CmtsMulticastAuthEntry.DocsIetfBpi2CmtsMulticastAuthCmMacAddress, "docsIetfBpi2CmtsMulticastAuthCmMacAddress")
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIetfBpi2CmtsMulticastAuthEntry.IfIndex})
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsMulticastAuthSAId", types.YLeaf{"DocsIetfBpi2CmtsMulticastAuthSAId", docsIetfBpi2CmtsMulticastAuthEntry.DocsIetfBpi2CmtsMulticastAuthSAId})
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsMulticastAuthCmMacAddress", types.YLeaf{"DocsIetfBpi2CmtsMulticastAuthCmMacAddress", docsIetfBpi2CmtsMulticastAuthEntry.DocsIetfBpi2CmtsMulticastAuthCmMacAddress})
    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsMulticastAuthControl", types.YLeaf{"DocsIetfBpi2CmtsMulticastAuthControl", docsIetfBpi2CmtsMulticastAuthEntry.DocsIetfBpi2CmtsMulticastAuthControl})

    docsIetfBpi2CmtsMulticastAuthEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIetfBpi2CmtsMulticastAuthSAId", "DocsIetfBpi2CmtsMulticastAuthCmMacAddress"}

    return &(docsIetfBpi2CmtsMulticastAuthEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable
// A table of CM certificate trust entries provisioned
// to the CMTS.  The trust object for a certificate in this
// table has an overriding effect on the validity object of a
// certificate in the authorization table, as long as the
// entire contents of the two certificates are identical.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the CMTS's provisioned CM certificate table.  Row entries
    // persist after re-initialization of the managed system. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry.
    DocsIetfBpi2CmtsProvisionedCmCertEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry
}

func (docsIetfBpi2CmtsProvisionedCmCertTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.YFilter = docsIetfBpi2CmtsProvisionedCmCertTable.YFilter
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.YangName = "docsIetfBpi2CmtsProvisionedCmCertTable"
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.SegmentPath = "docsIetfBpi2CmtsProvisionedCmCertTable"
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.Children.Append("docsIetfBpi2CmtsProvisionedCmCertEntry", types.YChild{"DocsIetfBpi2CmtsProvisionedCmCertEntry", nil})
    for i := range docsIetfBpi2CmtsProvisionedCmCertTable.DocsIetfBpi2CmtsProvisionedCmCertEntry {
        docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsProvisionedCmCertTable.DocsIetfBpi2CmtsProvisionedCmCertEntry[i]), types.YChild{"DocsIetfBpi2CmtsProvisionedCmCertEntry", docsIetfBpi2CmtsProvisionedCmCertTable.DocsIetfBpi2CmtsProvisionedCmCertEntry[i]})
    }
    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsProvisionedCmCertTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsProvisionedCmCertTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry
// An entry in the CMTS's provisioned CM certificate
// table.  Row entries persist after re-initialization of
// the managed system.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of this row. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIetfBpi2CmtsProvisionedCmCertMacAddress interface{}

    // Trust state for the provisioned CM certificate entry. Note: Setting this
    // object need only override the validity of CM certificates sent in future
    // authorization requests; instantaneous effect need not occur. The type is
    // DocsIetfBpi2CmtsProvisionedCmCertTrust.
    DocsIetfBpi2CmtsProvisionedCmCertTrust interface{}

    // This object indicates how the certificate reached the CMTS.  Other(4) means
    // is originated from a source not identified above. The type is
    // DocsIetfBpi2CmtsProvisionedCmCertSource.
    DocsIetfBpi2CmtsProvisionedCmCertSource interface{}

    // The status of this conceptual row. Values in this row cannot be changed
    // while the row is 'active'. The type is RowStatus.
    DocsIetfBpi2CmtsProvisionedCmCertStatus interface{}

    // An X509 DER-encoded Certificate Authority certificate. Note: The
    // zero-length OCTET STRING must be returned, on reads, if the entire
    // certificate is not retained in the CMTS. The type is string with length:
    // 0..4096.
    DocsIetfBpi2CmtsProvisionedCmCert interface{}
}

func (docsIetfBpi2CmtsProvisionedCmCertEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.YFilter = docsIetfBpi2CmtsProvisionedCmCertEntry.YFilter
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.YangName = "docsIetfBpi2CmtsProvisionedCmCertEntry"
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsProvisionedCmCertTable"
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsProvisionedCmCertEntry" + types.AddKeyToken(docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCertMacAddress, "docsIetfBpi2CmtsProvisionedCmCertMacAddress")
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsProvisionedCmCertMacAddress", types.YLeaf{"DocsIetfBpi2CmtsProvisionedCmCertMacAddress", docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCertMacAddress})
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsProvisionedCmCertTrust", types.YLeaf{"DocsIetfBpi2CmtsProvisionedCmCertTrust", docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCertTrust})
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsProvisionedCmCertSource", types.YLeaf{"DocsIetfBpi2CmtsProvisionedCmCertSource", docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCertSource})
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsProvisionedCmCertStatus", types.YLeaf{"DocsIetfBpi2CmtsProvisionedCmCertStatus", docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCertStatus})
    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsProvisionedCmCert", types.YLeaf{"DocsIetfBpi2CmtsProvisionedCmCert", docsIetfBpi2CmtsProvisionedCmCertEntry.DocsIetfBpi2CmtsProvisionedCmCert})

    docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData.YListKeys = []string {"DocsIetfBpi2CmtsProvisionedCmCertMacAddress"}

    return &(docsIetfBpi2CmtsProvisionedCmCertEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource represents identified above.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource_snmp DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource = "snmp"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource_configurationFile DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource = "configurationFile"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource_externalDatabase DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource = "externalDatabase"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource_other DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertSource = "other"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust represents instantaneous effect need not occur.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust_trusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust = "trusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust_untrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsProvisionedCmCertTable_DocsIetfBpi2CmtsProvisionedCmCertEntry_DocsIetfBpi2CmtsProvisionedCmCertTrust = "untrusted"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable
// The table of known Certificate Authority certificates
// acquired by this device.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A row in the Certificate Authority certificate table.  Row entries with
    // trust status 'trusted', 'untrusted', or 'root' persist after
    // re-initialization  of the managed system. The type is slice of
    // DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry.
    DocsIetfBpi2CmtsCACertEntry []*DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry
}

func (docsIetfBpi2CmtsCACertTable *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsCACertTable.EntityData.YFilter = docsIetfBpi2CmtsCACertTable.YFilter
    docsIetfBpi2CmtsCACertTable.EntityData.YangName = "docsIetfBpi2CmtsCACertTable"
    docsIetfBpi2CmtsCACertTable.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsCACertTable.EntityData.ParentYangName = "DOCS-IETF-BPI2-MIB"
    docsIetfBpi2CmtsCACertTable.EntityData.SegmentPath = "docsIetfBpi2CmtsCACertTable"
    docsIetfBpi2CmtsCACertTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsCACertTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsCACertTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsCACertTable.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsCACertTable.EntityData.Children.Append("docsIetfBpi2CmtsCACertEntry", types.YChild{"DocsIetfBpi2CmtsCACertEntry", nil})
    for i := range docsIetfBpi2CmtsCACertTable.DocsIetfBpi2CmtsCACertEntry {
        docsIetfBpi2CmtsCACertTable.EntityData.Children.Append(types.GetSegmentPath(docsIetfBpi2CmtsCACertTable.DocsIetfBpi2CmtsCACertEntry[i]), types.YChild{"DocsIetfBpi2CmtsCACertEntry", docsIetfBpi2CmtsCACertTable.DocsIetfBpi2CmtsCACertEntry[i]})
    }
    docsIetfBpi2CmtsCACertTable.EntityData.Leafs = types.NewOrderedMap()

    docsIetfBpi2CmtsCACertTable.EntityData.YListKeys = []string {}

    return &(docsIetfBpi2CmtsCACertTable.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry
// A row in the Certificate Authority certificate
// table.  Row entries with trust status 'trusted',
// 'untrusted', or 'root' persist after re-initialization
//  of the managed system.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index for this row. The type is interface{}
    // with range: 1..4294967295.
    DocsIetfBpi2CmtsCACertIndex interface{}

    // The subject name exactly as it is encoded in the X509 certificate. The
    // organizationName portion of the certificate's subject name must be present.
    // All other fields are optional.  Any optional field present must be pre
    // pended with <CR> (carriage return, U+000D) <LF> (line feed, U+000A).
    // Ordering of fields present must conform to: organizationName <CR> <LF>
    // countryName <CR> <LF> stateOrProvinceName <CR> <LF> localityName <CR> <LF>
    // organizationalUnitName <CR> <LF> organizationalUnitName=<Manufacturing
    // Location> <CR> <LF> commonName. The type is string.
    DocsIetfBpi2CmtsCACertSubject interface{}

    // The issuer name exactly as it is encoded in the X509 certificate. The
    // commonName portion of the certificate's issuer name must be present.  All
    // other fields are optional.  Any optional field present must be pre pended
    // with <CR> (carriage return, U+000D) <LF> (line feed, U+000A). Ordering of
    // fields present must conform to:  CommonName <CR><LF> countryName <CR><LF>
    // stateOrProvinceName <CR><LF> localityName <CR><LF> organizationName
    // <CR><LF> organizationalUnitName <CR><LF>
    // organizationalUnitName=<Manufacturing Location>. The type is string.
    DocsIetfBpi2CmtsCACertIssuer interface{}

    // This CA certificate's serial number represented as an octet string. The
    // type is string with length: 1..32.
    DocsIetfBpi2CmtsCACertSerialNumber interface{}

    // This object controls the trust status of this certificate.  Root
    // certificates must be given root(4) trust; manufacturer certificates must
    // not be given root(4) trust.  Trust on root certificates must not change.
    // Note: Setting this object need only affect the validity of CM certificates
    // sent in future authorization requests; instantaneous effect need not occur.
    // The type is DocsIetfBpi2CmtsCACertTrust.
    DocsIetfBpi2CmtsCACertTrust interface{}

    // This object indicates how the certificate reached the CMTS.  Other(4) means
    // it originated from a source not identified above. The type is
    // DocsIetfBpi2CmtsCACertSource.
    DocsIetfBpi2CmtsCACertSource interface{}

    // The status of this conceptual row. An attempt to set writable columnar
    // values while this row is active behaves as follows: - Sets to the object
    // docsIetfBpi2CmtsCACertTrust are allowed. - Sets to the object
    // docsIetfBpi2CmtsCACert will return an   error inconsistentValue'. A newly
    // create entry cannot be set to active until the value of
    // docsIetfBpi2CmtsCACert is being set. The type is RowStatus.
    DocsIetfBpi2CmtsCACertStatus interface{}

    // An X509 DER-encoded Certificate Authority certificate. To help identify
    // certificates, either this object or docsIetfBpi2CmtsCACertThumbprint must
    // be returned by a CMTS for self-signed CA certificates.  Note: The
    // zero-length OCTET STRING must be returned, on reads, if the entire
    // certificate is not retained in the CMTS. The type is string with length:
    // 0..4096.
    DocsIetfBpi2CmtsCACert interface{}

    // The SHA-1 hash of a CA certificate. To help identify certificates, either
    // this object or docsIetfBpi2CmtsCACert must be returned by a CMTS for
    // self-signed CA certificates.  Note: The zero-length OCTET STRING must be
    // returned, on reads, if the CA certificate thumb print is not retained in
    // the CMTS. The type is string with length: 20.
    DocsIetfBpi2CmtsCACertThumbprint interface{}
}

func (docsIetfBpi2CmtsCACertEntry *DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry) GetEntityData() *types.CommonEntityData {
    docsIetfBpi2CmtsCACertEntry.EntityData.YFilter = docsIetfBpi2CmtsCACertEntry.YFilter
    docsIetfBpi2CmtsCACertEntry.EntityData.YangName = "docsIetfBpi2CmtsCACertEntry"
    docsIetfBpi2CmtsCACertEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIetfBpi2CmtsCACertEntry.EntityData.ParentYangName = "docsIetfBpi2CmtsCACertTable"
    docsIetfBpi2CmtsCACertEntry.EntityData.SegmentPath = "docsIetfBpi2CmtsCACertEntry" + types.AddKeyToken(docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertIndex, "docsIetfBpi2CmtsCACertIndex")
    docsIetfBpi2CmtsCACertEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIetfBpi2CmtsCACertEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIetfBpi2CmtsCACertEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIetfBpi2CmtsCACertEntry.EntityData.Children = types.NewOrderedMap()
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertIndex", types.YLeaf{"DocsIetfBpi2CmtsCACertIndex", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertIndex})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertSubject", types.YLeaf{"DocsIetfBpi2CmtsCACertSubject", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertSubject})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertIssuer", types.YLeaf{"DocsIetfBpi2CmtsCACertIssuer", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertIssuer})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertSerialNumber", types.YLeaf{"DocsIetfBpi2CmtsCACertSerialNumber", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertSerialNumber})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertTrust", types.YLeaf{"DocsIetfBpi2CmtsCACertTrust", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertTrust})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertSource", types.YLeaf{"DocsIetfBpi2CmtsCACertSource", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertSource})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertStatus", types.YLeaf{"DocsIetfBpi2CmtsCACertStatus", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertStatus})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACert", types.YLeaf{"DocsIetfBpi2CmtsCACert", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACert})
    docsIetfBpi2CmtsCACertEntry.EntityData.Leafs.Append("docsIetfBpi2CmtsCACertThumbprint", types.YLeaf{"DocsIetfBpi2CmtsCACertThumbprint", docsIetfBpi2CmtsCACertEntry.DocsIetfBpi2CmtsCACertThumbprint})

    docsIetfBpi2CmtsCACertEntry.EntityData.YListKeys = []string {"DocsIetfBpi2CmtsCACertIndex"}

    return &(docsIetfBpi2CmtsCACertEntry.EntityData)
}

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource represents identified above.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_snmp DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "snmp"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_configurationFile DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "configurationFile"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_externalDatabase DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "externalDatabase"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_other DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "other"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_authentInfo DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "authentInfo"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource_compiledIntoCode DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertSource = "compiledIntoCode"
)

// DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust represents instantaneous effect need not occur.
type DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust string

const (
    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust_trusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust = "trusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust_untrusted DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust = "untrusted"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust_chained DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust = "chained"

    DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust_root DOCSIETFBPI2MIB_DocsIetfBpi2CmtsCACertTable_DocsIetfBpi2CmtsCACertEntry_DocsIetfBpi2CmtsCACertTrust = "root"
)

