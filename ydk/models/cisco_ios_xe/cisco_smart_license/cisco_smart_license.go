// Smart licensing configuration, RPC, notification and operational data.
// 
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package cisco_smart_license

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_smart_license"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/cisco-smart-license register-id-token}", reflect.TypeOf(RegisterIdToken{}))
    ydk.RegisterEntity("cisco-smart-license:register-id-token", reflect.TypeOf(RegisterIdToken{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/cisco-smart-license de-register}", reflect.TypeOf(DeRegister{}))
    ydk.RegisterEntity("cisco-smart-license:de-register", reflect.TypeOf(DeRegister{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/cisco-smart-license renew-id}", reflect.TypeOf(RenewId{}))
    ydk.RegisterEntity("cisco-smart-license:renew-id", reflect.TypeOf(RenewId{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/cisco-smart-license renew-auth}", reflect.TypeOf(RenewAuth{}))
    ydk.RegisterEntity("cisco-smart-license:renew-auth", reflect.TypeOf(RenewAuth{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/cisco-smart-license licensing}", reflect.TypeOf(Licensing{}))
    ydk.RegisterEntity("cisco-smart-license:licensing", reflect.TypeOf(Licensing{}))
}

// TransportTypeEnum represents The type of transport in use by smart licensing.
type TransportTypeEnum string

const (
    // Smart Licensing is using callhome for communications.
    TransportTypeEnum_transport_type_callhome TransportTypeEnum = "transport-type-callhome"

    // Smart licensing is using the smart transport for 
    // communications.
    TransportTypeEnum_transport_type_smart TransportTypeEnum = "transport-type-smart"
)

// ErrorEnum represents Smart Licensing RPC calls
type ErrorEnum string

const (
    // Success
    ErrorEnum_success ErrorEnum = "success"

    // Malloc Error
    ErrorEnum_malloc ErrorEnum = "malloc"

    // Null pointer Error
    ErrorEnum_nullpointer ErrorEnum = "nullpointer"

    // deprecated DO NOT remove
    ErrorEnum_error3 ErrorEnum = "error3"

    // deprecated DO NOT remove
    ErrorEnum_error4 ErrorEnum = "error4"

    // deprecated DO NOT remove
    ErrorEnum_error5 ErrorEnum = "error5"

    // Bad input parameter
    ErrorEnum_BadInputParams ErrorEnum = "BadInputParams"

    // deprecated DO NOT remove
    ErrorEnum_error7 ErrorEnum = "error7"

    // Bad handle
    ErrorEnum_badhandle ErrorEnum = "badhandle"

    // The requested item was not found
    ErrorEnum_notfound ErrorEnum = "notfound"

    // The requested operation is not supported
    ErrorEnum_notsupported ErrorEnum = "notsupported"

    // Init failed because the agent is already initialized
    ErrorEnum_alreadyinit ErrorEnum = "alreadyinit"

    // API failed because the agent is not initialized
    ErrorEnum_notinit ErrorEnum = "notinit"

    // State machine creation failed
    ErrorEnum_smfailtocreate ErrorEnum = "smfailtocreate"

    // State machine not running
    ErrorEnum_smfailtorun ErrorEnum = "smfailtorun"

    // State machine failed to init
    ErrorEnum_smfailtoinit ErrorEnum = "smfailtoinit"

    // State machine failed to destroy
    ErrorEnum_smfailtodestroy ErrorEnum = "smfailtodestroy"

    // message parsing error
    ErrorEnum_msgparse ErrorEnum = "msgparse"

    // message building error
    ErrorEnum_msgbuild ErrorEnum = "msgbuild"

    // Smart Agent not enabled
    ErrorEnum_notenabled ErrorEnum = "notenabled"

    // Smart Agent request invalid
    ErrorEnum_invalidrequest ErrorEnum = "invalidrequest"

    // General initialization error. We call a number of 
    // system routines to 
    // initialize system resources and we can't translate their error 
    // codes to Smart Agent error codes. The log should have a detailed 
    // description of the error.
    ErrorEnum_init ErrorEnum = "init"

    // Smart Agent Fail to set state
    ErrorEnum_failtosetstate ErrorEnum = "failtosetstate"

    // Unsupported response type
    ErrorEnum_unsupportedresponse ErrorEnum = "unsupportedresponse"

    // Invalid response type
    ErrorEnum_invalidresponse ErrorEnum = "invalidresponse"

    // Smart Agent Trusted Storage failed to store*
    ErrorEnum_storagefailtostore ErrorEnum = "storagefailtostore"

    // Smart Agent Trusted Storage failed to retrieve*
    ErrorEnum_storagefailtoretrieve ErrorEnum = "storagefailtoretrieve"

    // Null CCOId and IdToken
    ErrorEnum_nullccoidtoken ErrorEnum = "nullccoidtoken"

    // Product Instance identifier failed to match
    ErrorEnum_matchidentifier ErrorEnum = "matchidentifier"

    // Vendor string failed to match
    ErrorEnum_matchvendor ErrorEnum = "matchvendor"

    // nonce failed to match
    ErrorEnum_matchnonce ErrorEnum = "matchnonce"

    // communication channel error. Comm layer (call home) is disabled 
    ErrorEnum_commdisabled ErrorEnum = "commdisabled"

    // Call Home message send error. probably a timeout so we
    // will retry the send. any other error from the comm send should
    // be a permanent failure
    ErrorEnum_commsend ErrorEnum = "commsend"

    // Call Home message send response error
    ErrorEnum_commresponse ErrorEnum = "commresponse"

    // Call Home unknown error
    ErrorEnum_communkown ErrorEnum = "communkown"

    // State machine Operation not permitted
    ErrorEnum_smpostnotallow ErrorEnum = "smpostnotallow"

    // Missing mandatory field in request message
    ErrorEnum_reqmsgmissingmandatoryfield ErrorEnum = "reqmsgmissingmandatoryfield"

    // We received a failure status in a response message.
    // The log will contain a error message
    ErrorEnum_responsefailed ErrorEnum = "responsefailed"

    // PI not initialized 
    ErrorEnum_pinotinit ErrorEnum = "pinotinit"

    // The agent cannot be enabled more than once
    ErrorEnum_alreadyenabled ErrorEnum = "alreadyenabled"

    // The agent is already registered
    ErrorEnum_alreadyregistered ErrorEnum = "alreadyregistered"

    // The certificate is invalid 
    ErrorEnum_certinvalid ErrorEnum = "certinvalid"

    // The certificate has expired 
    ErrorEnum_certexpired ErrorEnum = "certexpired"

    // The agent is not registered 
    ErrorEnum_notregistered ErrorEnum = "notregistered"

    // The CSR generation failed 
    ErrorEnum_csrgenerationfailed ErrorEnum = "csrgenerationfailed"

    // Signature Verification failed 
    ErrorEnum_verifysignaturefailed ErrorEnum = "verifysignaturefailed"

    // Signature Generation failed 
    ErrorEnum_generatesignaturefailed ErrorEnum = "generatesignaturefailed"

    // Signing Certificate Verification failed 
    ErrorEnum_signcertverificationfailed ErrorEnum = "signcertverificationfailed"

    // Node Certificate Verification failed 
    ErrorEnum_nodecertverificationfailed ErrorEnum = "nodecertverificationfailed"

    // Certificate Parsing failed 
    ErrorEnum_parsecertificatefailed ErrorEnum = "parsecertificatefailed"

    // Root Certificate Import failed 
    ErrorEnum_cryptorootcaimportfailed ErrorEnum = "cryptorootcaimportfailed"

    // The tag is invalid 
    ErrorEnum_taginvalid ErrorEnum = "taginvalid"

    // Smart agent is running on a standby RP 
    ErrorEnum_standby ErrorEnum = "standby"

    // Smart agent id token registration is in progress
    ErrorEnum_registrationinprogress ErrorEnum = "registrationinprogress"

    // Call Home is not ready because it is in restart ipc
    ErrorEnum_commretry ErrorEnum = "commretry"

    // Smart agent authorization renew is in progress
    ErrorEnum_authrenewinprogress ErrorEnum = "authrenewinprogress"

    // Smart agent id certificate renew is in progress
    ErrorEnum_idcertrenewinprogress ErrorEnum = "idcertrenewinprogress"

    // Udi List Has Not been changed 
    ErrorEnum_noudichange ErrorEnum = "noudichange"

    // Call home service cannot be turned on. 
    ErrorEnum_callhomeserviceoff ErrorEnum = "callhomeserviceoff"

    // message execution already in progress
    ErrorEnum_msgexecinprogress ErrorEnum = "msgexecinprogress"

    // message execution in progress flag is locked
    ErrorEnum_msgexecinproglocked ErrorEnum = "msgexecinproglocked"

    // The ID cert only matches some of the system Udi's
    ErrorEnum_certmatchessubsetudis ErrorEnum = "certmatchessubsetudis"

    // A Storage group has not all been changed
    ErrorEnum_storagegroupchangeincomplete ErrorEnum = "storagegroupchangeincomplete"

    // Storage Management is not Init
    ErrorEnum_storagemgmtnotinit ErrorEnum = "storagemgmtnotinit"

    // TS System Path list is not changed 
    ErrorEnum_tspathnotchanged ErrorEnum = "tspathnotchanged"

    // Crypto Initialization is not completed
    ErrorEnum_cryptoinitnotcompleted ErrorEnum = "cryptoinitnotcompleted"

    // The agent is not in unidentified state 
    ErrorEnum_notinunidentified ErrorEnum = "notinunidentified"

    // The platform provided path is invalid 
    ErrorEnum_platformpathinvalid ErrorEnum = "platformpathinvalid"

    // The platform provided UDI is invalid 
    ErrorEnum_platformudiinvalid ErrorEnum = "platformudiinvalid"

    // failed to create Trusted Store object 
    ErrorEnum_storageobjfailtocreate ErrorEnum = "storageobjfailtocreate"

    // failed to erase trusted store object 
    ErrorEnum_storageobjfailtoerase ErrorEnum = "storageobjfailtoerase"

    // trusted storage object/file does not exist
    ErrorEnum_storageobjdoesnotexist ErrorEnum = "storageobjdoesnotexist"

    // The message event is beyond the peer 
    ErrorEnum_messageeventexceedspeer ErrorEnum = "messageeventexceedspeer"

    // Validation of the authorization key failed.
    // It probably does not match the UDI. The device will go to
    // the unidentified state (not registered)
    ErrorEnum_codevalidationfailed ErrorEnum = "codevalidationfailed"

    // operation not supported because the agent is running in
    // permanent License reservation mode
    ErrorEnum_reserved ErrorEnum = "reserved"

    // No license reservation is in progress 
    ErrorEnum_noreservationinprogress ErrorEnum = "noreservationinprogress"

    // No authorization code instaled in device 
    ErrorEnum_noauthorizationinstalled ErrorEnum = "noauthorizationinstalled"

    // The reservation authorization code does not match the 
    // reservation request code
    ErrorEnum_reservationmismatch ErrorEnum = "reservationmismatch"

    // not in license reservation mode 
    ErrorEnum_notreservationmode ErrorEnum = "notreservationmode"

    //  General reservation error. This is used with the API 
    // functions that are called by the CLI. the API will return a 
    // very specific displayString that describes the error.
    ErrorEnum_reservationerror ErrorEnum = "reservationerror"

    // Sysmgr Init Failed 
    ErrorEnum_sysmgrinit ErrorEnum = "sysmgrinit"

    // Generic error for something already existing 
    ErrorEnum_alreadyexists ErrorEnum = "alreadyexists"

    // Error in object insert to xos list 
    ErrorEnum_listinsertfailed ErrorEnum = "listinsertfailed"

    // Session management not initialized
    ErrorEnum_sessionmgmtnotinit ErrorEnum = "sessionmgmtnotinit"

    // Error Creating Linked List
    ErrorEnum_listinitfailed ErrorEnum = "listinitfailed"

    // List in use
    ErrorEnum_listbusy ErrorEnum = "listbusy"

    // No Connected Clients 
    ErrorEnum_noclients ErrorEnum = "noclients"

    // Generic IPC layer error 
    ErrorEnum_ipc ErrorEnum = "ipc"

    // The IPC socket open error 
    ErrorEnum_ipcopen ErrorEnum = "ipcopen"

    // The IPC Initialization error 
    ErrorEnum_ipcinit ErrorEnum = "ipcinit"

    // The IPC Connection error 
    ErrorEnum_ipcconnect ErrorEnum = "ipcconnect"

    // The IPC Server Event error 
    ErrorEnum_ipcevents ErrorEnum = "ipcevents"

    // The IPC Management error 
    ErrorEnum_ipcmgmt ErrorEnum = "ipcmgmt"

    // The IPC Send error 
    ErrorEnum_ipcsend ErrorEnum = "ipcsend"

    // The IPC Recevive error 
    ErrorEnum_ipcreceive ErrorEnum = "ipcreceive"

    // The IPC Recevive error 
    ErrorEnum_ipctimeout ErrorEnum = "ipctimeout"

    // Failed to enqueue a message to the IPC Queue
    ErrorEnum_enqueuefailed ErrorEnum = "enqueuefailed"

    // Failed to dequeue a message from the IPC queue
    ErrorEnum_dequeuefailed ErrorEnum = "dequeuefailed"

    // Fail because we are about to shutdown and we need
    // to stop processing any more messages or responses
    ErrorEnum_shuttingdown ErrorEnum = "shuttingdown"

    // Could not validate Trust Chain
    ErrorEnum_couldnotvalidatetrustchain ErrorEnum = "couldnotvalidatetrustchain"

    // The reservation authorization code is already installed 
    ErrorEnum_reservationalreadyinstalled ErrorEnum = "reservationalreadyinstalled"

    // Failed to parse reservation authorization code 
    ErrorEnum_reservationinstallparsefail ErrorEnum = "reservationinstallparsefail"

    // Base64 encoding failed 
    ErrorEnum_base64encoding ErrorEnum = "base64encoding"

    // Base64 decoding failed 
    ErrorEnum_base64decoding ErrorEnum = "base64decoding"

    // Failed to find UUID inside software id tag 
    ErrorEnum_invalidsoftwareidtag ErrorEnum = "invalidsoftwareidtag"

    // Development certificates are being used with Production 
    // Root certificate 
    ErrorEnum_certificatemismatch ErrorEnum = "certificatemismatch"

    // No License Reservation 
    ErrorEnum_noreservation ErrorEnum = "noreservation"

    // the agent Daemon is unreachable 
    ErrorEnum_agentunreachable ErrorEnum = "agentunreachable"

    // the agent ignores event 
    ErrorEnum_ignoreevent ErrorEnum = "ignoreevent"

    // Base58 overflow, number too large. 
    ErrorEnum_b58overflow ErrorEnum = "b58overflow"

    // Base58 decode error. 
    ErrorEnum_b58decode ErrorEnum = "b58decode"

    // Bad base58 length. 
    ErrorEnum_b58badlen ErrorEnum = "b58badlen"

    // Invalid base58 digit. 
    ErrorEnum_b58invdigit ErrorEnum = "b58invdigit"

    // Overflow detected during base58 decode. 
    ErrorEnum_b58decodeoverflow ErrorEnum = "b58decodeoverflow"

    // Reservation version out of bound 
    ErrorEnum_reservationversionoutofbound ErrorEnum = "reservationversionoutofbound"

    // General Base58 encoding error 
    ErrorEnum_base58encode ErrorEnum = "base58encode"

    // General error code for adding item that already exists. 
    // Used in App HA setup when adding an HA peer device info that 
    // is already added 
    ErrorEnum_duplicatedentry ErrorEnum = "duplicatedentry"

    // General error code for trying to remove item that do 
    // not exist. Used in App HA setup when removing an HA peer device 
    // info that does not exist
    ErrorEnum_missingentry ErrorEnum = "missingentry"

    // The given peer info contain incorrect data format 
    ErrorEnum_badpeerinfoformat ErrorEnum = "badpeerinfoformat"

    // The given handle attribute list contains incomplete
    // application HA attributes
    ErrorEnum_badapplicationhaattributedataset ErrorEnum = "badapplicationhaattributedataset"

    // license reservation is in progress 
    ErrorEnum_reservationinprogress ErrorEnum = "reservationinprogress"

    // The xos_dm_create() failure causes this error code to be return 
    ErrorEnum_xdmcreatehandle ErrorEnum = "xdmcreatehandle"

    // Version in entitlement response message does not match 
    // with the
    // one already saved from last completed response message. We need
    // to send entitlemeent request again with proper version and data
    ErrorEnum_versionmismatchinentitlementrsp ErrorEnum = "versionmismatchinentitlementrsp"

    // Given valid HA Role is not supported by this operation. 
    // For Application HA valid roles are Active or Standby.
    ErrorEnum_harolenotsupported ErrorEnum = "harolenotsupported"

    // Application HA attribute contains invalid character.  
    // Character set supported is alphanumeric.
    ErrorEnum_apphainvalidcharacter ErrorEnum = "apphainvalidcharacter"

    // The peer info given is from the same device. 
    ErrorEnum_apphaaddpeerfromsamedevice ErrorEnum = "apphaaddpeerfromsamedevice"

    // When setting Application HA Attribute, a different 
    // handle with exactly
    // same tag, app ha name, app ha id exists in the agent.
    ErrorEnum_apphaappduplicatedinstance ErrorEnum = "apphaappduplicatedinstance"

    // Backend server does not support AppHA enabled message 
    // version. So
    // registration response (from backend) received with status set as 
    // FAILED or VERSION_TOO_HIGH.
    ErrorEnum_versionmismatchinregresponse ErrorEnum = "versionmismatchinregresponse"

    // Migration function was called but there are no i
    // call backs registered 
    ErrorEnum_conversionnocb ErrorEnum = "conversionnocb"

    // Migration was not enabled with the 
    // MIGRATION_ALLOWED property 
    ErrorEnum_conversionnotallowed ErrorEnum = "conversionnotallowed"

    // Migration is in progress 
    ErrorEnum_conversioninprogress ErrorEnum = "conversioninprogress"

    // Migration has already been started 
    ErrorEnum_conversionalreadystarted ErrorEnum = "conversionalreadystarted"

    // Migration is not enabled 
    ErrorEnum_conversionnotenabled ErrorEnum = "conversionnotenabled"

    // The backend derver this device is connected to does 
    // not support conversion
    ErrorEnum_versionconversionnotsupported ErrorEnum = "versionconversionnotsupported"

    // No conversion in progress
    ErrorEnum_noconversioninprogress ErrorEnum = "noconversioninprogress"

    // Loaded OpenSSL version is not the same as used for 
    // development
    ErrorEnum_cryptoversionmismatch ErrorEnum = "cryptoversionmismatch"

    // Some Smart Licensing Conversion jobs stopped successfully
    ErrorEnum_conversionstoppedpartially ErrorEnum = "conversionstoppedpartially"

    // Operation is not supported because Utility management 
    // is enabled
    ErrorEnum_utilityenabled ErrorEnum = "utilityenabled"

    // Utility is not enabled
    ErrorEnum_utilitynotenabled ErrorEnum = "utilitynotenabled"

    // Transport layer not abvailable
    ErrorEnum_transportnotavailable ErrorEnum = "transportnotavailable"

    // Unable to get FQDN
    ErrorEnum_fqdn ErrorEnum = "fqdn"

    // The current transport type does not support this feature
    ErrorEnum_thirdparty ErrorEnum = "thirdparty"

    // The transport type needed is not configured
    ErrorEnum_transporttype ErrorEnum = "transporttype"

    // max error code
    ErrorEnum_max ErrorEnum = "max"
)

// UtilityReportingTypeEnum represents What has triggered the system to start reporting utility usage.
type UtilityReportingTypeEnum string

const (
    // The system is not reporting utility usage data.
    UtilityReportingTypeEnum_utility_reporting_type_none UtilityReportingTypeEnum = "utility-reporting-type-none"

    // The system is reporting utility usage data because it has 
    // received subscription information from either the SSM or 
    // satellite.
    UtilityReportingTypeEnum_utility_reporting_type_subscription UtilityReportingTypeEnum = "utility-reporting-type-subscription"

    // The system is reporting utility usage data because it has 
    // received a utility certificate from a Third Party 
    // Billing Platform.
    UtilityReportingTypeEnum_utility_reporting_type_certificate UtilityReportingTypeEnum = "utility-reporting-type-certificate"
)

// EnforcementModeEnum represents  tells us how the license is being enforced.
type EnforcementModeEnum string

const (
    // The initial state after an entitlement request while we are waiting 
    // the Authorization request response. In this mode the device will 
    // have established communications with Cisco and successfully 
    // registered with the Cisco Licensing cloud.
    EnforcementModeEnum_enforcement_waiting EnforcementModeEnum = "enforcement-waiting"

    // Cisco Smart Software Manager (CSSM) has responded that
    // the entitlement requested is in compliance.
    EnforcementModeEnum_enforcement_in_compliance EnforcementModeEnum = "enforcement-in-compliance"

    // Cisco Smart Software Manager (CSSM) has responded that
    // the entitlement requested is out of compliance. 
    // either too many licenses /entitlements are in use or the license 
    // has not been purchased
    EnforcementModeEnum_enforcement_out_of_compliance EnforcementModeEnum = "enforcement-out-of-compliance"

    // more licenses are in use than were purchased but the customer
    //  is still within the terms of their contract
    EnforcementModeEnum_enforcement_overage EnforcementModeEnum = "enforcement-overage"

    // The evaluation period is in use.
    // It will remain in use until the following
    // two messages have been received by the product from the 
    // Cisco Smart Software Manager (CSSM):
    //  Successful response to a registration request,
    //  successful response to an entitlement authorization request
    EnforcementModeEnum_enforcement_evaluation EnforcementModeEnum = "enforcement-evaluation"

    // The evaluation period has expired
    EnforcementModeEnum_enforcement_evaluation_expired EnforcementModeEnum = "enforcement-evaluation-expired"

    // Authorization period has expired. This will occur if the product
    // has not been able to communicate with Cisco or a satellite 
    // for an extended period of time, usually 90 days.
    EnforcementModeEnum_enforcement_authorization_expired EnforcementModeEnum = "enforcement-authorization-expired"

    // The entitlement requested is in compliance because 
    // a reservation authorization code is installed and the product
    // is in Permanent License Reservation mode.
    EnforcementModeEnum_enforcement_reservation_in_compliance EnforcementModeEnum = "enforcement-reservation-in-compliance"

    // The entitlement tag is invalid.
    // The CSSM does not recognize the entitlement tag
    // because it is not in the database. This usually only occurs
    // during testing.
    EnforcementModeEnum_enforcement_invalid_tag EnforcementModeEnum = "enforcement-invalid-tag"

    // Smart licensing has been disabled. The feature using this license
    // should be disabled.
    EnforcementModeEnum_enforcement_disabled EnforcementModeEnum = "enforcement-disabled"
)

// AuthorizationStateEnum represents The smart licensing authorization state.
type AuthorizationStateEnum string

const (
    // No licenses are in use so there is no authorization 
    // state to report.
    AuthorizationStateEnum_auth_state_none AuthorizationStateEnum = "auth-state-none"

    // Evaluation period is in use and is counting down.
    AuthorizationStateEnum_auth_state_eval AuthorizationStateEnum = "auth-state-eval"

    // Evaluation period in use but it has expired.
    AuthorizationStateEnum_auth_state_eval_expired AuthorizationStateEnum = "auth-state-eval-expired"

    // All license usage is authorized and within terms 
    // of the customer's contract.
    AuthorizationStateEnum_auth_state_authorized AuthorizationStateEnum = "auth-state-authorized"

    // All license usage is authorized because a  
    // reservation authorization code is installed.
    AuthorizationStateEnum_auth_state_authorized_reservation AuthorizationStateEnum = "auth-state-authorized-reservation"

    // License usage is out of compliance with the 
    // terms of the contract. Either too many licenses are in
    // use or licenses that were not purchased are in use.
    AuthorizationStateEnum_auth_state_out_of_compliance AuthorizationStateEnum = "auth-state-out-of-compliance"

    // The authorization period has expired because this
    // product instance has not communicated with the
    // SSM or satellite in over 90 days.
    AuthorizationStateEnum_auth_state_authorization_expired AuthorizationStateEnum = "auth-state-authorization-expired"
)

// RegistrationStateEnum represents The smart licensing registration state.
type RegistrationStateEnum string

const (
    // This smart licensing instance is not registered.
    RegistrationStateEnum_reg_state_not_registered RegistrationStateEnum = "reg-state-not-registered"

    // Registration was successful and this smart licensing 
    // instance is registered.
    RegistrationStateEnum_reg_state_complete RegistrationStateEnum = "reg-state-complete"

    // Registration is in progress.
    RegistrationStateEnum_reg_state_in_progress RegistrationStateEnum = "reg-state-in-progress"

    // The initial registration attempt failed but
    // a retry is in progress.
    RegistrationStateEnum_reg_state_retry RegistrationStateEnum = "reg-state-retry"

    // Registration failed.
    RegistrationStateEnum_reg_state_failed RegistrationStateEnum = "reg-state-failed"
)

// NotifRegisterFailureEnum represents detailed failure message.
type NotifRegisterFailureEnum string

const (
    // General failure.
    NotifRegisterFailureEnum_general_failure NotifRegisterFailureEnum = "general-failure"

    // This smart licensing instance is already registered.
    NotifRegisterFailureEnum_already_registered_failure NotifRegisterFailureEnum = "already-registered-failure"

    // The de-register failed because this instance is not registered.
    NotifRegisterFailureEnum_de_register_failure NotifRegisterFailureEnum = "de-register-failure"
)

// RegisterIdToken
// Register with an ID token.
// This will begin the registration process. Since the registration 
// process will take somewhere between seconds and hours you must get the  
// registration-success or registration-fail notifications 
// or check the registration status in smart-license:state
// to know the status of the registration.
type RegisterIdToken struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RegisterIdToken_Input

    
    Output RegisterIdToken_Output
}

func (registerIdToken *RegisterIdToken) GetEntityData() *types.CommonEntityData {
    registerIdToken.EntityData.YFilter = registerIdToken.YFilter
    registerIdToken.EntityData.YangName = "register-id-token"
    registerIdToken.EntityData.BundleName = "cisco_ios_xe"
    registerIdToken.EntityData.ParentYangName = "cisco-smart-license"
    registerIdToken.EntityData.SegmentPath = "cisco-smart-license:register-id-token"
    registerIdToken.EntityData.AbsolutePath = registerIdToken.EntityData.SegmentPath
    registerIdToken.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registerIdToken.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registerIdToken.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registerIdToken.EntityData.Children = types.NewOrderedMap()
    registerIdToken.EntityData.Children.Append("input", types.YChild{"Input", &registerIdToken.Input})
    registerIdToken.EntityData.Children.Append("output", types.YChild{"Output", &registerIdToken.Output})
    registerIdToken.EntityData.Leafs = types.NewOrderedMap()

    registerIdToken.EntityData.YListKeys = []string {}

    return &(registerIdToken.EntityData)
}

// RegisterIdToken_Input
type RegisterIdToken_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The ID token used to register. The type is string with length: 1..255.
    IdToken interface{}

    // Force the registration if set. The type is bool. The default value is
    // false.
    Force interface{}
}

func (input *RegisterIdToken_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "register-id-token"
    input.EntityData.SegmentPath = "input"
    input.EntityData.AbsolutePath = "cisco-smart-license:register-id-token/" + input.EntityData.SegmentPath
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("id-token", types.YLeaf{"IdToken", input.IdToken})
    input.EntityData.Leafs.Append("force", types.YLeaf{"Force", input.Force})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RegisterIdToken_Output
type RegisterIdToken_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The return code.  If smart licensing is not enabled (status:enabled) then
    // the error will be error-enum:notenabled. On success the return code will be
    // error-enum:registrationinprogress. The type is ErrorEnum.
    ReturnCode interface{}
}

func (output *RegisterIdToken_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "register-id-token"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-smart-license:register-id-token/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("return-code", types.YLeaf{"ReturnCode", output.ReturnCode})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// DeRegister
// De-register. This will immediately de-register the device.
type DeRegister struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output DeRegister_Output
}

func (deRegister *DeRegister) GetEntityData() *types.CommonEntityData {
    deRegister.EntityData.YFilter = deRegister.YFilter
    deRegister.EntityData.YangName = "de-register"
    deRegister.EntityData.BundleName = "cisco_ios_xe"
    deRegister.EntityData.ParentYangName = "cisco-smart-license"
    deRegister.EntityData.SegmentPath = "cisco-smart-license:de-register"
    deRegister.EntityData.AbsolutePath = deRegister.EntityData.SegmentPath
    deRegister.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    deRegister.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    deRegister.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    deRegister.EntityData.Children = types.NewOrderedMap()
    deRegister.EntityData.Children.Append("output", types.YChild{"Output", &deRegister.Output})
    deRegister.EntityData.Leafs = types.NewOrderedMap()

    deRegister.EntityData.YListKeys = []string {}

    return &(deRegister.EntityData)
}

// DeRegister_Output
type DeRegister_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The return code for the de-register process. The type is ErrorEnum.
    ReturnCode interface{}
}

func (output *DeRegister_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "de-register"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-smart-license:de-register/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("return-code", types.YLeaf{"ReturnCode", output.ReturnCode})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// RenewId
// Under normal operations smart licensing will automatically renew 
// the ID certificates used for regsitration. This command can 
// be used if he customer wants to initiate a manual 
// registration renewal.
type RenewId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output RenewId_Output
}

func (renewId *RenewId) GetEntityData() *types.CommonEntityData {
    renewId.EntityData.YFilter = renewId.YFilter
    renewId.EntityData.YangName = "renew-id"
    renewId.EntityData.BundleName = "cisco_ios_xe"
    renewId.EntityData.ParentYangName = "cisco-smart-license"
    renewId.EntityData.SegmentPath = "cisco-smart-license:renew-id"
    renewId.EntityData.AbsolutePath = renewId.EntityData.SegmentPath
    renewId.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    renewId.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    renewId.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    renewId.EntityData.Children = types.NewOrderedMap()
    renewId.EntityData.Children.Append("output", types.YChild{"Output", &renewId.Output})
    renewId.EntityData.Leafs = types.NewOrderedMap()

    renewId.EntityData.YListKeys = []string {}

    return &(renewId.EntityData)
}

// RenewId_Output
type RenewId_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The return code. The type is ErrorEnum.
    ReturnCode interface{}
}

func (output *RenewId_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "renew-id"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-smart-license:renew-id/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("return-code", types.YLeaf{"ReturnCode", output.ReturnCode})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// RenewAuth
// Under normal operations smart licensing will automatically renew 
// the license authorization every 30 days. This command can 
// be used if he customer wants to iinitiate a manual 
// renewal.
type RenewAuth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Output RenewAuth_Output
}

func (renewAuth *RenewAuth) GetEntityData() *types.CommonEntityData {
    renewAuth.EntityData.YFilter = renewAuth.YFilter
    renewAuth.EntityData.YangName = "renew-auth"
    renewAuth.EntityData.BundleName = "cisco_ios_xe"
    renewAuth.EntityData.ParentYangName = "cisco-smart-license"
    renewAuth.EntityData.SegmentPath = "cisco-smart-license:renew-auth"
    renewAuth.EntityData.AbsolutePath = renewAuth.EntityData.SegmentPath
    renewAuth.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    renewAuth.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    renewAuth.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    renewAuth.EntityData.Children = types.NewOrderedMap()
    renewAuth.EntityData.Children.Append("output", types.YChild{"Output", &renewAuth.Output})
    renewAuth.EntityData.Leafs = types.NewOrderedMap()

    renewAuth.EntityData.YListKeys = []string {}

    return &(renewAuth.EntityData)
}

// RenewAuth_Output
type RenewAuth_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The return code. The type is ErrorEnum.
    ReturnCode interface{}
}

func (output *RenewAuth_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "renew-auth"
    output.EntityData.SegmentPath = "output"
    output.EntityData.AbsolutePath = "cisco-smart-license:renew-auth/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("return-code", types.YLeaf{"ReturnCode", output.ReturnCode})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// Licensing
// Container to hold config and state.
type Licensing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Smart licensing configuration.
    Config Licensing_Config

    // Smart licensing state.
    State Licensing_State
}

func (licensing *Licensing) GetEntityData() *types.CommonEntityData {
    licensing.EntityData.YFilter = licensing.YFilter
    licensing.EntityData.YangName = "licensing"
    licensing.EntityData.BundleName = "cisco_ios_xe"
    licensing.EntityData.ParentYangName = "cisco-smart-license"
    licensing.EntityData.SegmentPath = "cisco-smart-license:licensing"
    licensing.EntityData.AbsolutePath = licensing.EntityData.SegmentPath
    licensing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    licensing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    licensing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    licensing.EntityData.Children = types.NewOrderedMap()
    licensing.EntityData.Children.Append("config", types.YChild{"Config", &licensing.Config})
    licensing.EntityData.Children.Append("state", types.YChild{"State", &licensing.State})
    licensing.EntityData.Leafs = types.NewOrderedMap()

    licensing.EntityData.YListKeys = []string {}

    return &(licensing.EntityData)
}

// Licensing_Config
// Smart licensing configuration.
type Licensing_Config struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/disable smart licensing. If state:always-enabled is true then Smart
    // Licensing is always enabled and config:enable will be  silently ignored.
    // The type is bool.
    Enable interface{}

    // A free form ID set by the customer which will be included in the utility
    // usage (RUM) report and inthe message header. The type is string.
    CustomId interface{}

    // Controls whether or not some information is sent in messages to the   CSSM
    // or satellite.
    Privacy Licensing_Config_Privacy

    // Utility settings.
    Utility Licensing_Config_Utility

    // Transport layer settings.
    Transport Licensing_Config_Transport
}

func (config *Licensing_Config) GetEntityData() *types.CommonEntityData {
    config.EntityData.YFilter = config.YFilter
    config.EntityData.YangName = "config"
    config.EntityData.BundleName = "cisco_ios_xe"
    config.EntityData.ParentYangName = "licensing"
    config.EntityData.SegmentPath = "config"
    config.EntityData.AbsolutePath = "cisco-smart-license:licensing/" + config.EntityData.SegmentPath
    config.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    config.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    config.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    config.EntityData.Children = types.NewOrderedMap()
    config.EntityData.Children.Append("privacy", types.YChild{"Privacy", &config.Privacy})
    config.EntityData.Children.Append("utility", types.YChild{"Utility", &config.Utility})
    config.EntityData.Children.Append("transport", types.YChild{"Transport", &config.Transport})
    config.EntityData.Leafs = types.NewOrderedMap()
    config.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", config.Enable})
    config.EntityData.Leafs.Append("custom-id", types.YLeaf{"CustomId", config.CustomId})

    config.EntityData.YListKeys = []string {}

    return &(config.EntityData)
}

// Licensing_Config_Privacy
// Controls whether or not some information is sent in messages to the 
//  CSSM or satellite.
type Licensing_Config_Privacy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If true then the hostname will not be sent in  any messages. The type is
    // bool.
    Hostname interface{}

    // If true then the smart licensing version will not be sent in  any messages.
    // The type is bool.
    Version interface{}
}

func (privacy *Licensing_Config_Privacy) GetEntityData() *types.CommonEntityData {
    privacy.EntityData.YFilter = privacy.YFilter
    privacy.EntityData.YangName = "privacy"
    privacy.EntityData.BundleName = "cisco_ios_xe"
    privacy.EntityData.ParentYangName = "config"
    privacy.EntityData.SegmentPath = "privacy"
    privacy.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/" + privacy.EntityData.SegmentPath
    privacy.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    privacy.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    privacy.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    privacy.EntityData.Children = types.NewOrderedMap()
    privacy.EntityData.Leafs = types.NewOrderedMap()
    privacy.EntityData.Leafs.Append("hostname", types.YLeaf{"Hostname", privacy.Hostname})
    privacy.EntityData.Leafs.Append("version", types.YLeaf{"Version", privacy.Version})

    privacy.EntityData.YListKeys = []string {}

    return &(privacy.EntityData)
}

// Licensing_Config_Utility
// Utility settings.
type Licensing_Config_Utility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates the customer's intent to start reporting  utility usage
    // information. This alone does not enable  utility reporting. Either
    // subscription information will be  automatically downloaded from the SSM or
    // a utility certificate we be loaded if the system registers with a  Third
    // Party Billing Platform. The type is bool.
    UtilityEnable interface{}

    // Customer address information that will be included in  the utility usage
    // reports.
    CustomerInfo Licensing_Config_Utility_CustomerInfo
}

func (utility *Licensing_Config_Utility) GetEntityData() *types.CommonEntityData {
    utility.EntityData.YFilter = utility.YFilter
    utility.EntityData.YangName = "utility"
    utility.EntityData.BundleName = "cisco_ios_xe"
    utility.EntityData.ParentYangName = "config"
    utility.EntityData.SegmentPath = "utility"
    utility.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/" + utility.EntityData.SegmentPath
    utility.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utility.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utility.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utility.EntityData.Children = types.NewOrderedMap()
    utility.EntityData.Children.Append("customer-info", types.YChild{"CustomerInfo", &utility.CustomerInfo})
    utility.EntityData.Leafs = types.NewOrderedMap()
    utility.EntityData.Leafs.Append("utility-enable", types.YLeaf{"UtilityEnable", utility.UtilityEnable})

    utility.EntityData.YListKeys = []string {}

    return &(utility.EntityData)
}

// Licensing_Config_Utility_CustomerInfo
// Customer address information that will be included in 
// the utility usage reports.
type Licensing_Config_Utility_CustomerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The cisco issued customer id which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    Id interface{}

    // The customer company name which will be included in the utility usage (RUM)
    // report. The type is string with length: 1..250.
    Name interface{}

    // The customer company street address which will be included in the utility
    // usage (RUM) report. The type is string with length: 1..250.
    Street interface{}

    // The customer company city which will be included in the utility usage (RUM)
    // report. The type is string with length: 1..250.
    City interface{}

    // The customer company state which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    State interface{}

    // The customer company country which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    Country interface{}

    // The customer location specific postal code which will be included in the
    // utility usage (RUM) report. The type is string with length: 1..250.
    PostalCode interface{}
}

func (customerInfo *Licensing_Config_Utility_CustomerInfo) GetEntityData() *types.CommonEntityData {
    customerInfo.EntityData.YFilter = customerInfo.YFilter
    customerInfo.EntityData.YangName = "customer-info"
    customerInfo.EntityData.BundleName = "cisco_ios_xe"
    customerInfo.EntityData.ParentYangName = "utility"
    customerInfo.EntityData.SegmentPath = "customer-info"
    customerInfo.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/utility/" + customerInfo.EntityData.SegmentPath
    customerInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    customerInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    customerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    customerInfo.EntityData.Children = types.NewOrderedMap()
    customerInfo.EntityData.Leafs = types.NewOrderedMap()
    customerInfo.EntityData.Leafs.Append("id", types.YLeaf{"Id", customerInfo.Id})
    customerInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", customerInfo.Name})
    customerInfo.EntityData.Leafs.Append("street", types.YLeaf{"Street", customerInfo.Street})
    customerInfo.EntityData.Leafs.Append("city", types.YLeaf{"City", customerInfo.City})
    customerInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", customerInfo.State})
    customerInfo.EntityData.Leafs.Append("country", types.YLeaf{"Country", customerInfo.Country})
    customerInfo.EntityData.Leafs.Append("postal-code", types.YLeaf{"PostalCode", customerInfo.PostalCode})

    customerInfo.EntityData.YListKeys = []string {}

    return &(customerInfo.EntityData)
}

// Licensing_Config_Transport
// Transport layer settings.
type Licensing_Config_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The transport type. If transport-type is set to transport-type-callhome
    // then any additional transport settings must be done from the callhome CLI.
    // If the transport-type is set to transport-type-smart additional settings
    // are available below. The type is TransportTypeEnum.
    TransportType interface{}

    // Settings for the smart transport.
    TransportSmart Licensing_Config_Transport_TransportSmart
}

func (transport *Licensing_Config_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xe"
    transport.EntityData.ParentYangName = "config"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Children.Append("transport-smart", types.YChild{"TransportSmart", &transport.TransportSmart})
    transport.EntityData.Leafs = types.NewOrderedMap()
    transport.EntityData.Leafs.Append("transport-type", types.YLeaf{"TransportType", transport.TransportType})

    transport.EntityData.YListKeys = []string {}

    return &(transport.EntityData)
}

// Licensing_Config_Transport_TransportSmart
// Settings for the smart transport.
type Licensing_Config_Transport_TransportSmart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tell smart licensing to set the default URLs for both url-registration and
    // url-utility that point to  the Cisco SSM. The type is bool.
    UrlDefault interface{}

    // The URLS used for registration and utility reporting. These should be set
    // if the product instance will communicate with a satellite or Third Party 
    // Billing Platform.
    Urls Licensing_Config_Transport_TransportSmart_Urls
}

func (transportSmart *Licensing_Config_Transport_TransportSmart) GetEntityData() *types.CommonEntityData {
    transportSmart.EntityData.YFilter = transportSmart.YFilter
    transportSmart.EntityData.YangName = "transport-smart"
    transportSmart.EntityData.BundleName = "cisco_ios_xe"
    transportSmart.EntityData.ParentYangName = "transport"
    transportSmart.EntityData.SegmentPath = "transport-smart"
    transportSmart.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/transport/" + transportSmart.EntityData.SegmentPath
    transportSmart.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transportSmart.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transportSmart.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transportSmart.EntityData.Children = types.NewOrderedMap()
    transportSmart.EntityData.Children.Append("urls", types.YChild{"Urls", &transportSmart.Urls})
    transportSmart.EntityData.Leafs = types.NewOrderedMap()
    transportSmart.EntityData.Leafs.Append("url-default", types.YLeaf{"UrlDefault", transportSmart.UrlDefault})

    transportSmart.EntityData.YListKeys = []string {}

    return &(transportSmart.EntityData)
}

// Licensing_Config_Transport_TransportSmart_Urls
// The URLS used for registration and utility reporting.
// These should be set if the product instance will
// communicate with a satellite or Third Party 
// Billing Platform.
type Licensing_Config_Transport_TransportSmart_Urls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the URL used for registration, authorization and anything else not
    // related to utility usage reporting. The type is string.
    UrlRegistration interface{}

    // Set the URL to be used for sending utility usage  reports. This should be
    // the same as url-registration  if you are using a satellite. The type is
    // string.
    UrlUtility interface{}
}

func (urls *Licensing_Config_Transport_TransportSmart_Urls) GetEntityData() *types.CommonEntityData {
    urls.EntityData.YFilter = urls.YFilter
    urls.EntityData.YangName = "urls"
    urls.EntityData.BundleName = "cisco_ios_xe"
    urls.EntityData.ParentYangName = "transport-smart"
    urls.EntityData.SegmentPath = "urls"
    urls.EntityData.AbsolutePath = "cisco-smart-license:licensing/config/transport/transport-smart/" + urls.EntityData.SegmentPath
    urls.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    urls.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    urls.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    urls.EntityData.Children = types.NewOrderedMap()
    urls.EntityData.Leafs = types.NewOrderedMap()
    urls.EntityData.Leafs.Append("url-registration", types.YLeaf{"UrlRegistration", urls.UrlRegistration})
    urls.EntityData.Leafs.Append("url-utility", types.YLeaf{"UrlUtility", urls.UrlUtility})

    urls.EntityData.YListKeys = []string {}

    return &(urls.EntityData)
}

// Licensing_State
// Smart licensing state.
type Licensing_State struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Smart Licensing is always enabled. the smart-enabled leaf below  will
    // always be true. The config:enable setting above will be silently  ignored
    // by Smart licensing. The type is bool.
    AlwaysEnabled interface{}

    // Is smart licensing enabled. If always-enabled above is true then  smart
    // licensing is always enabled and can not be disabled. The type is bool.
    SmartEnabled interface{}

    // The smart licensing version in use. The type is string with length: 1..255.
    Version interface{}

    // Smart licensing state information.     This is only valid if smart-enabled
    // = true.
    StateInfo Licensing_State_StateInfo
}

func (state *Licensing_State) GetEntityData() *types.CommonEntityData {
    state.EntityData.YFilter = state.YFilter
    state.EntityData.YangName = "state"
    state.EntityData.BundleName = "cisco_ios_xe"
    state.EntityData.ParentYangName = "licensing"
    state.EntityData.SegmentPath = "state"
    state.EntityData.AbsolutePath = "cisco-smart-license:licensing/" + state.EntityData.SegmentPath
    state.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    state.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    state.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    state.EntityData.Children = types.NewOrderedMap()
    state.EntityData.Children.Append("state-info", types.YChild{"StateInfo", &state.StateInfo})
    state.EntityData.Leafs = types.NewOrderedMap()
    state.EntityData.Leafs.Append("always-enabled", types.YLeaf{"AlwaysEnabled", state.AlwaysEnabled})
    state.EntityData.Leafs.Append("smart-enabled", types.YLeaf{"SmartEnabled", state.SmartEnabled})
    state.EntityData.Leafs.Append("version", types.YLeaf{"Version", state.Version})

    state.EntityData.YListKeys = []string {}

    return &(state.EntityData)
}

// Licensing_State_StateInfo
// Smart licensing state information. 
//    This is only valid if smart-enabled = true.
type Licensing_State_StateInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The custom ID set by the customer that will be included in the utility
    // usage (RUM) report and in th emessage header. The type is string.
    CustomId interface{}

    // State of license registration.
    Registration Licensing_State_StateInfo_Registration

    // State of license authorization.
    Authorization Licensing_State_StateInfo_Authorization

    // State of utility reporting.
    Utility Licensing_State_StateInfo_Utility

    // State of the transport.
    Transport Licensing_State_StateInfo_Transport

    // State of the privacy settings.
    Privacy Licensing_State_StateInfo_Privacy

    // State of the evaluation period.
    Evaluation Licensing_State_StateInfo_Evaluation

    // UDI of the system.
    Udi Licensing_State_StateInfo_Udi

    // List of license (entitlement tag) usage information.  This only contains
    // the information for licenses that are in use. The type is slice of
    // Licensing_State_StateInfo_Usage.
    Usage []*Licensing_State_StateInfo_Usage
}

func (stateInfo *Licensing_State_StateInfo) GetEntityData() *types.CommonEntityData {
    stateInfo.EntityData.YFilter = stateInfo.YFilter
    stateInfo.EntityData.YangName = "state-info"
    stateInfo.EntityData.BundleName = "cisco_ios_xe"
    stateInfo.EntityData.ParentYangName = "state"
    stateInfo.EntityData.SegmentPath = "state-info"
    stateInfo.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/" + stateInfo.EntityData.SegmentPath
    stateInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    stateInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    stateInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    stateInfo.EntityData.Children = types.NewOrderedMap()
    stateInfo.EntityData.Children.Append("registration", types.YChild{"Registration", &stateInfo.Registration})
    stateInfo.EntityData.Children.Append("authorization", types.YChild{"Authorization", &stateInfo.Authorization})
    stateInfo.EntityData.Children.Append("utility", types.YChild{"Utility", &stateInfo.Utility})
    stateInfo.EntityData.Children.Append("transport", types.YChild{"Transport", &stateInfo.Transport})
    stateInfo.EntityData.Children.Append("privacy", types.YChild{"Privacy", &stateInfo.Privacy})
    stateInfo.EntityData.Children.Append("evaluation", types.YChild{"Evaluation", &stateInfo.Evaluation})
    stateInfo.EntityData.Children.Append("udi", types.YChild{"Udi", &stateInfo.Udi})
    stateInfo.EntityData.Children.Append("usage", types.YChild{"Usage", nil})
    for i := range stateInfo.Usage {
        stateInfo.EntityData.Children.Append(types.GetSegmentPath(stateInfo.Usage[i]), types.YChild{"Usage", stateInfo.Usage[i]})
    }
    stateInfo.EntityData.Leafs = types.NewOrderedMap()
    stateInfo.EntityData.Leafs.Append("custom-id", types.YLeaf{"CustomId", stateInfo.CustomId})

    stateInfo.EntityData.YListKeys = []string {}

    return &(stateInfo.EntityData)
}

// Licensing_State_StateInfo_Registration
// State of license registration.
type Licensing_State_StateInfo_Registration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current registration state. The type is RegistrationStateEnum.
    RegistrationState interface{}

    // Is the device allowed to enable export controlled features. The type is
    // bool.
    ExportControlAllowed interface{}

    // Registration is in progress.
    RegistrationInProgress Licensing_State_StateInfo_Registration_RegistrationInProgress

    // Registration failed.
    RegistrationFailed Licensing_State_StateInfo_Registration_RegistrationFailed

    // Registration failed and doing a retry.
    RegistrationRetry Licensing_State_StateInfo_Registration_RegistrationRetry

    // Registration success.
    RegistrationComplete Licensing_State_StateInfo_Registration_RegistrationComplete
}

func (registration *Licensing_State_StateInfo_Registration) GetEntityData() *types.CommonEntityData {
    registration.EntityData.YFilter = registration.YFilter
    registration.EntityData.YangName = "registration"
    registration.EntityData.BundleName = "cisco_ios_xe"
    registration.EntityData.ParentYangName = "state-info"
    registration.EntityData.SegmentPath = "registration"
    registration.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + registration.EntityData.SegmentPath
    registration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registration.EntityData.Children = types.NewOrderedMap()
    registration.EntityData.Children.Append("registration-in-progress", types.YChild{"RegistrationInProgress", &registration.RegistrationInProgress})
    registration.EntityData.Children.Append("registration-failed", types.YChild{"RegistrationFailed", &registration.RegistrationFailed})
    registration.EntityData.Children.Append("registration-retry", types.YChild{"RegistrationRetry", &registration.RegistrationRetry})
    registration.EntityData.Children.Append("registration-complete", types.YChild{"RegistrationComplete", &registration.RegistrationComplete})
    registration.EntityData.Leafs = types.NewOrderedMap()
    registration.EntityData.Leafs.Append("registration-state", types.YLeaf{"RegistrationState", registration.RegistrationState})
    registration.EntityData.Leafs.Append("export-control-allowed", types.YLeaf{"ExportControlAllowed", registration.ExportControlAllowed})

    registration.EntityData.YListKeys = []string {}

    return &(registration.EntityData)
}

// Licensing_State_StateInfo_Registration_RegistrationInProgress
// Registration is in progress.
type Licensing_State_StateInfo_Registration_RegistrationInProgress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the registration started. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartTime interface{}
}

func (registrationInProgress *Licensing_State_StateInfo_Registration_RegistrationInProgress) GetEntityData() *types.CommonEntityData {
    registrationInProgress.EntityData.YFilter = registrationInProgress.YFilter
    registrationInProgress.EntityData.YangName = "registration-in-progress"
    registrationInProgress.EntityData.BundleName = "cisco_ios_xe"
    registrationInProgress.EntityData.ParentYangName = "registration"
    registrationInProgress.EntityData.SegmentPath = "registration-in-progress"
    registrationInProgress.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/registration/" + registrationInProgress.EntityData.SegmentPath
    registrationInProgress.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registrationInProgress.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registrationInProgress.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registrationInProgress.EntityData.Children = types.NewOrderedMap()
    registrationInProgress.EntityData.Leafs = types.NewOrderedMap()
    registrationInProgress.EntityData.Leafs.Append("start-time", types.YLeaf{"StartTime", registrationInProgress.StartTime})

    registrationInProgress.EntityData.YListKeys = []string {}

    return &(registrationInProgress.EntityData)
}

// Licensing_State_StateInfo_Registration_RegistrationFailed
// Registration failed.
type Licensing_State_StateInfo_Registration_RegistrationFailed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the registration failed. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    FailTime interface{}

    // Failure message that can be displayed for the user.  This is not a parsable
    // message. The type is string with length: 0..255.
    FailMessage interface{}
}

func (registrationFailed *Licensing_State_StateInfo_Registration_RegistrationFailed) GetEntityData() *types.CommonEntityData {
    registrationFailed.EntityData.YFilter = registrationFailed.YFilter
    registrationFailed.EntityData.YangName = "registration-failed"
    registrationFailed.EntityData.BundleName = "cisco_ios_xe"
    registrationFailed.EntityData.ParentYangName = "registration"
    registrationFailed.EntityData.SegmentPath = "registration-failed"
    registrationFailed.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/registration/" + registrationFailed.EntityData.SegmentPath
    registrationFailed.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registrationFailed.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registrationFailed.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registrationFailed.EntityData.Children = types.NewOrderedMap()
    registrationFailed.EntityData.Leafs = types.NewOrderedMap()
    registrationFailed.EntityData.Leafs.Append("fail-time", types.YLeaf{"FailTime", registrationFailed.FailTime})
    registrationFailed.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", registrationFailed.FailMessage})

    registrationFailed.EntityData.YListKeys = []string {}

    return &(registrationFailed.EntityData)
}

// Licensing_State_StateInfo_Registration_RegistrationRetry
// Registration failed and doing a retry.
type Licensing_State_StateInfo_Registration_RegistrationRetry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the registration will be retried. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    RetryNextTime interface{}

    // Time the registration failed. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    FailTime interface{}

    // Failure message that can be displayed for the user.  This is not a parsable
    // message. The type is string with length: 0..255.
    FailMessage interface{}
}

func (registrationRetry *Licensing_State_StateInfo_Registration_RegistrationRetry) GetEntityData() *types.CommonEntityData {
    registrationRetry.EntityData.YFilter = registrationRetry.YFilter
    registrationRetry.EntityData.YangName = "registration-retry"
    registrationRetry.EntityData.BundleName = "cisco_ios_xe"
    registrationRetry.EntityData.ParentYangName = "registration"
    registrationRetry.EntityData.SegmentPath = "registration-retry"
    registrationRetry.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/registration/" + registrationRetry.EntityData.SegmentPath
    registrationRetry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registrationRetry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registrationRetry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registrationRetry.EntityData.Children = types.NewOrderedMap()
    registrationRetry.EntityData.Leafs = types.NewOrderedMap()
    registrationRetry.EntityData.Leafs.Append("retry-next-time", types.YLeaf{"RetryNextTime", registrationRetry.RetryNextTime})
    registrationRetry.EntityData.Leafs.Append("fail-time", types.YLeaf{"FailTime", registrationRetry.FailTime})
    registrationRetry.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", registrationRetry.FailMessage})

    registrationRetry.EntityData.YListKeys = []string {}

    return &(registrationRetry.EntityData)
}

// Licensing_State_StateInfo_Registration_RegistrationComplete
// Registration success.
type Licensing_State_StateInfo_Registration_RegistrationComplete struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the registration was successful. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CompleteTime interface{}

    // Time the last registration renewal occurred.  If empty then no renewal has
    // occurred. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastRenewTime interface{}

    // Time the registration will be automatically renewed. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextRenewTime interface{}

    // Time the registration will expire if it is not renewed. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ExpireTime interface{}

    // Was the last renewal attempt successful. The type is bool.
    LastRenewSuccess interface{}

    // If the last renewal failed then this is a failure message that can be
    // displayed for the user.  This is not a parsable string. Only present if the
    // last renewal failed. The type is string with length: 0..255.
    FailMessage interface{}

    // The smart account name for this registration. The type is string with
    // length: 1..255.
    SmartAccount interface{}

    // The virtual account name for this registration. The type is string with
    // length: 1..255.
    VirtualAccount interface{}
}

func (registrationComplete *Licensing_State_StateInfo_Registration_RegistrationComplete) GetEntityData() *types.CommonEntityData {
    registrationComplete.EntityData.YFilter = registrationComplete.YFilter
    registrationComplete.EntityData.YangName = "registration-complete"
    registrationComplete.EntityData.BundleName = "cisco_ios_xe"
    registrationComplete.EntityData.ParentYangName = "registration"
    registrationComplete.EntityData.SegmentPath = "registration-complete"
    registrationComplete.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/registration/" + registrationComplete.EntityData.SegmentPath
    registrationComplete.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    registrationComplete.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    registrationComplete.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    registrationComplete.EntityData.Children = types.NewOrderedMap()
    registrationComplete.EntityData.Leafs = types.NewOrderedMap()
    registrationComplete.EntityData.Leafs.Append("complete-time", types.YLeaf{"CompleteTime", registrationComplete.CompleteTime})
    registrationComplete.EntityData.Leafs.Append("last-renew-time", types.YLeaf{"LastRenewTime", registrationComplete.LastRenewTime})
    registrationComplete.EntityData.Leafs.Append("next-renew-time", types.YLeaf{"NextRenewTime", registrationComplete.NextRenewTime})
    registrationComplete.EntityData.Leafs.Append("expire-time", types.YLeaf{"ExpireTime", registrationComplete.ExpireTime})
    registrationComplete.EntityData.Leafs.Append("last-renew-success", types.YLeaf{"LastRenewSuccess", registrationComplete.LastRenewSuccess})
    registrationComplete.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", registrationComplete.FailMessage})
    registrationComplete.EntityData.Leafs.Append("smart-account", types.YLeaf{"SmartAccount", registrationComplete.SmartAccount})
    registrationComplete.EntityData.Leafs.Append("virtual-account", types.YLeaf{"VirtualAccount", registrationComplete.VirtualAccount})

    registrationComplete.EntityData.YListKeys = []string {}

    return &(registrationComplete.EntityData)
}

// Licensing_State_StateInfo_Authorization
// State of license authorization.
type Licensing_State_StateInfo_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The current authorization state. The type is AuthorizationStateEnum.
    AuthorizationState interface{}

    // No licenses in use. This empty container is not needed but is  a place
    // holder to show there is no data for this state.
    AuthorizationNone Licensing_State_StateInfo_Authorization_AuthorizationNone

    // Evaluation period is in use and counting down. The evaluation period only
    // counts down when licenses are in use.
    AuthorizationEval Licensing_State_StateInfo_Authorization_AuthorizationEval

    // Evaluation period is in use but has expired.
    AuthorizationEvalExpired Licensing_State_StateInfo_Authorization_AuthorizationEvalExpired

    // All license usage is authorized and within terms of the contract.
    AuthorizationAuthorized Licensing_State_StateInfo_Authorization_AuthorizationAuthorized

    // All license usage is authorized because a   reservation authorization code
    // is installed.
    AuthorizationAuthorizedReservation Licensing_State_StateInfo_Authorization_AuthorizationAuthorizedReservation

    // License usage is out of compliance with the terms of the  contract because
    // more licenses are in use than were purchased.
    AuthorizationOutOfCompliance Licensing_State_StateInfo_Authorization_AuthorizationOutOfCompliance

    // The authorization period has expired because the product  instance ahs not
    // communicated with the SSM or satellite in  over 90 days.
    AuthorizationAuthorizationExpired Licensing_State_StateInfo_Authorization_AuthorizationAuthorizationExpired
}

func (authorization *Licensing_State_StateInfo_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xe"
    authorization.EntityData.ParentYangName = "state-info"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + authorization.EntityData.SegmentPath
    authorization.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorization.EntityData.Children = types.NewOrderedMap()
    authorization.EntityData.Children.Append("authorization-none", types.YChild{"AuthorizationNone", &authorization.AuthorizationNone})
    authorization.EntityData.Children.Append("authorization-eval", types.YChild{"AuthorizationEval", &authorization.AuthorizationEval})
    authorization.EntityData.Children.Append("authorization-eval-expired", types.YChild{"AuthorizationEvalExpired", &authorization.AuthorizationEvalExpired})
    authorization.EntityData.Children.Append("authorization-authorized", types.YChild{"AuthorizationAuthorized", &authorization.AuthorizationAuthorized})
    authorization.EntityData.Children.Append("authorization-authorized-reservation", types.YChild{"AuthorizationAuthorizedReservation", &authorization.AuthorizationAuthorizedReservation})
    authorization.EntityData.Children.Append("authorization-out-of-compliance", types.YChild{"AuthorizationOutOfCompliance", &authorization.AuthorizationOutOfCompliance})
    authorization.EntityData.Children.Append("authorization-authorization-expired", types.YChild{"AuthorizationAuthorizationExpired", &authorization.AuthorizationAuthorizationExpired})
    authorization.EntityData.Leafs = types.NewOrderedMap()
    authorization.EntityData.Leafs.Append("authorization-state", types.YLeaf{"AuthorizationState", authorization.AuthorizationState})

    authorization.EntityData.YListKeys = []string {}

    return &(authorization.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationNone
// No licenses in use. This empty container is not needed but is 
// a place holder to show there is no data for this state.
type Licensing_State_StateInfo_Authorization_AuthorizationNone struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (authorizationNone *Licensing_State_StateInfo_Authorization_AuthorizationNone) GetEntityData() *types.CommonEntityData {
    authorizationNone.EntityData.YFilter = authorizationNone.YFilter
    authorizationNone.EntityData.YangName = "authorization-none"
    authorizationNone.EntityData.BundleName = "cisco_ios_xe"
    authorizationNone.EntityData.ParentYangName = "authorization"
    authorizationNone.EntityData.SegmentPath = "authorization-none"
    authorizationNone.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationNone.EntityData.SegmentPath
    authorizationNone.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationNone.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationNone.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationNone.EntityData.Children = types.NewOrderedMap()
    authorizationNone.EntityData.Leafs = types.NewOrderedMap()

    authorizationNone.EntityData.YListKeys = []string {}

    return &(authorizationNone.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationEval
// Evaluation period is in use and counting down.
// The evaluation period only counts down when licenses are
// in use.
type Licensing_State_StateInfo_Authorization_AuthorizationEval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds of license usage until the evaluation  period expires.
    // Note that this not a hard date and time because if no licenses are in use
    // the evaluation period stops  counting down. The type is interface{} with
    // range: 0..18446744073709551615.
    SecondsLeft interface{}
}

func (authorizationEval *Licensing_State_StateInfo_Authorization_AuthorizationEval) GetEntityData() *types.CommonEntityData {
    authorizationEval.EntityData.YFilter = authorizationEval.YFilter
    authorizationEval.EntityData.YangName = "authorization-eval"
    authorizationEval.EntityData.BundleName = "cisco_ios_xe"
    authorizationEval.EntityData.ParentYangName = "authorization"
    authorizationEval.EntityData.SegmentPath = "authorization-eval"
    authorizationEval.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationEval.EntityData.SegmentPath
    authorizationEval.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationEval.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationEval.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationEval.EntityData.Children = types.NewOrderedMap()
    authorizationEval.EntityData.Leafs = types.NewOrderedMap()
    authorizationEval.EntityData.Leafs.Append("seconds-left", types.YLeaf{"SecondsLeft", authorizationEval.SecondsLeft})

    authorizationEval.EntityData.YListKeys = []string {}

    return &(authorizationEval.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationEvalExpired
// Evaluation period is in use but has expired.
type Licensing_State_StateInfo_Authorization_AuthorizationEvalExpired struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the evaluation period expired. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ExpireTime interface{}
}

func (authorizationEvalExpired *Licensing_State_StateInfo_Authorization_AuthorizationEvalExpired) GetEntityData() *types.CommonEntityData {
    authorizationEvalExpired.EntityData.YFilter = authorizationEvalExpired.YFilter
    authorizationEvalExpired.EntityData.YangName = "authorization-eval-expired"
    authorizationEvalExpired.EntityData.BundleName = "cisco_ios_xe"
    authorizationEvalExpired.EntityData.ParentYangName = "authorization"
    authorizationEvalExpired.EntityData.SegmentPath = "authorization-eval-expired"
    authorizationEvalExpired.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationEvalExpired.EntityData.SegmentPath
    authorizationEvalExpired.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationEvalExpired.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationEvalExpired.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationEvalExpired.EntityData.Children = types.NewOrderedMap()
    authorizationEvalExpired.EntityData.Leafs = types.NewOrderedMap()
    authorizationEvalExpired.EntityData.Leafs.Append("expire-time", types.YLeaf{"ExpireTime", authorizationEvalExpired.ExpireTime})

    authorizationEvalExpired.EntityData.YListKeys = []string {}

    return &(authorizationEvalExpired.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationAuthorized
// All license usage is authorized and within terms of the contract.
type Licensing_State_StateInfo_Authorization_AuthorizationAuthorized struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last communication was successful or failed. The type is bool.
    LastCommStatusSuccess interface{}

    // Failure message if the last communications attempt failed. This can be
    // displayed for the user. It is not a parsable string. The type is string
    // with length: 0..255.
    FailMessage interface{}

    // Time the last communication attempt happened. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastCommTime interface{}

    // The next time communications will be attempted to the back end. This will
    // be zero if the initial communication has not completed. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextCommTime interface{}

    // If there are no communications between now and this time smart licensing
    // will enter the authorization expired state.  This may be zero indicating
    // there is no deadline. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CommDeadlineTime interface{}
}

func (authorizationAuthorized *Licensing_State_StateInfo_Authorization_AuthorizationAuthorized) GetEntityData() *types.CommonEntityData {
    authorizationAuthorized.EntityData.YFilter = authorizationAuthorized.YFilter
    authorizationAuthorized.EntityData.YangName = "authorization-authorized"
    authorizationAuthorized.EntityData.BundleName = "cisco_ios_xe"
    authorizationAuthorized.EntityData.ParentYangName = "authorization"
    authorizationAuthorized.EntityData.SegmentPath = "authorization-authorized"
    authorizationAuthorized.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationAuthorized.EntityData.SegmentPath
    authorizationAuthorized.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationAuthorized.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationAuthorized.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationAuthorized.EntityData.Children = types.NewOrderedMap()
    authorizationAuthorized.EntityData.Leafs = types.NewOrderedMap()
    authorizationAuthorized.EntityData.Leafs.Append("last-comm-status-success", types.YLeaf{"LastCommStatusSuccess", authorizationAuthorized.LastCommStatusSuccess})
    authorizationAuthorized.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", authorizationAuthorized.FailMessage})
    authorizationAuthorized.EntityData.Leafs.Append("last-comm-time", types.YLeaf{"LastCommTime", authorizationAuthorized.LastCommTime})
    authorizationAuthorized.EntityData.Leafs.Append("next-comm-time", types.YLeaf{"NextCommTime", authorizationAuthorized.NextCommTime})
    authorizationAuthorized.EntityData.Leafs.Append("comm-deadline-time", types.YLeaf{"CommDeadlineTime", authorizationAuthorized.CommDeadlineTime})

    authorizationAuthorized.EntityData.YListKeys = []string {}

    return &(authorizationAuthorized.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationAuthorizedReservation
// All license usage is authorized because a  
// reservation authorization code is installed.
type Licensing_State_StateInfo_Authorization_AuthorizationAuthorizedReservation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the reservation occurred. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ReservationTime interface{}
}

func (authorizationAuthorizedReservation *Licensing_State_StateInfo_Authorization_AuthorizationAuthorizedReservation) GetEntityData() *types.CommonEntityData {
    authorizationAuthorizedReservation.EntityData.YFilter = authorizationAuthorizedReservation.YFilter
    authorizationAuthorizedReservation.EntityData.YangName = "authorization-authorized-reservation"
    authorizationAuthorizedReservation.EntityData.BundleName = "cisco_ios_xe"
    authorizationAuthorizedReservation.EntityData.ParentYangName = "authorization"
    authorizationAuthorizedReservation.EntityData.SegmentPath = "authorization-authorized-reservation"
    authorizationAuthorizedReservation.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationAuthorizedReservation.EntityData.SegmentPath
    authorizationAuthorizedReservation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationAuthorizedReservation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationAuthorizedReservation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationAuthorizedReservation.EntityData.Children = types.NewOrderedMap()
    authorizationAuthorizedReservation.EntityData.Leafs = types.NewOrderedMap()
    authorizationAuthorizedReservation.EntityData.Leafs.Append("reservation-time", types.YLeaf{"ReservationTime", authorizationAuthorizedReservation.ReservationTime})

    authorizationAuthorizedReservation.EntityData.YListKeys = []string {}

    return &(authorizationAuthorizedReservation.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationOutOfCompliance
// License usage is out of compliance with the terms of the 
// contract because more licenses are in use than were purchased.
type Licensing_State_StateInfo_Authorization_AuthorizationOutOfCompliance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last communication was successful or failed. The type is bool.
    LastCommStatusSuccess interface{}

    // Failure message if the last communications attempt failed. This can be
    // displayed for the user. It is not a parsable string. The type is string
    // with length: 0..255.
    FailMessage interface{}

    // Time the last communication attempt happened. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastCommTime interface{}

    // The next time communications will be attempted to the back end. This will
    // be zero if the initial communication has not completed. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextCommTime interface{}

    // If there are no communications between now and this time smart licensing
    // will enter the authorization expired state.  This may be zero indicating
    // there is no deadline. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CommDeadlineTime interface{}

    // Time the product instance entered the out of compliance state. The type is
    // string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    OocTime interface{}
}

func (authorizationOutOfCompliance *Licensing_State_StateInfo_Authorization_AuthorizationOutOfCompliance) GetEntityData() *types.CommonEntityData {
    authorizationOutOfCompliance.EntityData.YFilter = authorizationOutOfCompliance.YFilter
    authorizationOutOfCompliance.EntityData.YangName = "authorization-out-of-compliance"
    authorizationOutOfCompliance.EntityData.BundleName = "cisco_ios_xe"
    authorizationOutOfCompliance.EntityData.ParentYangName = "authorization"
    authorizationOutOfCompliance.EntityData.SegmentPath = "authorization-out-of-compliance"
    authorizationOutOfCompliance.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationOutOfCompliance.EntityData.SegmentPath
    authorizationOutOfCompliance.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationOutOfCompliance.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationOutOfCompliance.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationOutOfCompliance.EntityData.Children = types.NewOrderedMap()
    authorizationOutOfCompliance.EntityData.Leafs = types.NewOrderedMap()
    authorizationOutOfCompliance.EntityData.Leafs.Append("last-comm-status-success", types.YLeaf{"LastCommStatusSuccess", authorizationOutOfCompliance.LastCommStatusSuccess})
    authorizationOutOfCompliance.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", authorizationOutOfCompliance.FailMessage})
    authorizationOutOfCompliance.EntityData.Leafs.Append("last-comm-time", types.YLeaf{"LastCommTime", authorizationOutOfCompliance.LastCommTime})
    authorizationOutOfCompliance.EntityData.Leafs.Append("next-comm-time", types.YLeaf{"NextCommTime", authorizationOutOfCompliance.NextCommTime})
    authorizationOutOfCompliance.EntityData.Leafs.Append("comm-deadline-time", types.YLeaf{"CommDeadlineTime", authorizationOutOfCompliance.CommDeadlineTime})
    authorizationOutOfCompliance.EntityData.Leafs.Append("ooc-time", types.YLeaf{"OocTime", authorizationOutOfCompliance.OocTime})

    authorizationOutOfCompliance.EntityData.YListKeys = []string {}

    return &(authorizationOutOfCompliance.EntityData)
}

// Licensing_State_StateInfo_Authorization_AuthorizationAuthorizationExpired
// The authorization period has expired because the product 
// instance ahs not communicated with the SSM or satellite in 
// over 90 days.
type Licensing_State_StateInfo_Authorization_AuthorizationAuthorizationExpired struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Last communication was successful or failed. The type is bool.
    LastCommStatusSuccess interface{}

    // Failure message if the last communications attempt failed. This can be
    // displayed for the user. It is not a parsable string. The type is string
    // with length: 0..255.
    FailMessage interface{}

    // Time the last communication attempt happened. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastCommTime interface{}

    // The next time communications will be attempted to the back end. This will
    // be zero if the initial communication has not completed. The type is string
    // with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextCommTime interface{}

    // If there are no communications between now and this time smart licensing
    // will enter the authorization expired state.  This may be zero indicating
    // there is no deadline. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    CommDeadlineTime interface{}
}

func (authorizationAuthorizationExpired *Licensing_State_StateInfo_Authorization_AuthorizationAuthorizationExpired) GetEntityData() *types.CommonEntityData {
    authorizationAuthorizationExpired.EntityData.YFilter = authorizationAuthorizationExpired.YFilter
    authorizationAuthorizationExpired.EntityData.YangName = "authorization-authorization-expired"
    authorizationAuthorizationExpired.EntityData.BundleName = "cisco_ios_xe"
    authorizationAuthorizationExpired.EntityData.ParentYangName = "authorization"
    authorizationAuthorizationExpired.EntityData.SegmentPath = "authorization-authorization-expired"
    authorizationAuthorizationExpired.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/authorization/" + authorizationAuthorizationExpired.EntityData.SegmentPath
    authorizationAuthorizationExpired.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    authorizationAuthorizationExpired.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    authorizationAuthorizationExpired.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    authorizationAuthorizationExpired.EntityData.Children = types.NewOrderedMap()
    authorizationAuthorizationExpired.EntityData.Leafs = types.NewOrderedMap()
    authorizationAuthorizationExpired.EntityData.Leafs.Append("last-comm-status-success", types.YLeaf{"LastCommStatusSuccess", authorizationAuthorizationExpired.LastCommStatusSuccess})
    authorizationAuthorizationExpired.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", authorizationAuthorizationExpired.FailMessage})
    authorizationAuthorizationExpired.EntityData.Leafs.Append("last-comm-time", types.YLeaf{"LastCommTime", authorizationAuthorizationExpired.LastCommTime})
    authorizationAuthorizationExpired.EntityData.Leafs.Append("next-comm-time", types.YLeaf{"NextCommTime", authorizationAuthorizationExpired.NextCommTime})
    authorizationAuthorizationExpired.EntityData.Leafs.Append("comm-deadline-time", types.YLeaf{"CommDeadlineTime", authorizationAuthorizationExpired.CommDeadlineTime})

    authorizationAuthorizationExpired.EntityData.YListKeys = []string {}

    return &(authorizationAuthorizationExpired.EntityData)
}

// Licensing_State_StateInfo_Utility
// State of utility reporting.
type Licensing_State_StateInfo_Utility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Utility reporting is enabled. The system still  needs a utility certificate
    // registration or a subscription  to begin reporting actual utility data. The
    // type is bool.
    Enabled interface{}

    // Is the system reporting utility data. If so then what triggered  it to
    // start reporting. If this is utility-reporting-none then all of the times
    // below will be zero. The type is UtilityReportingTypeEnum.
    Reporting interface{}

    // If the product instance is reporting utility data this will contain various
    // timing information about that reporting.
    ReportingTimes Licensing_State_StateInfo_Utility_ReportingTimes

    // Customer address information that will be sent  in the utility usage
    // reports.
    CustomerInfo Licensing_State_StateInfo_Utility_CustomerInfo
}

func (utility *Licensing_State_StateInfo_Utility) GetEntityData() *types.CommonEntityData {
    utility.EntityData.YFilter = utility.YFilter
    utility.EntityData.YangName = "utility"
    utility.EntityData.BundleName = "cisco_ios_xe"
    utility.EntityData.ParentYangName = "state-info"
    utility.EntityData.SegmentPath = "utility"
    utility.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + utility.EntityData.SegmentPath
    utility.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utility.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utility.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utility.EntityData.Children = types.NewOrderedMap()
    utility.EntityData.Children.Append("reporting-times", types.YChild{"ReportingTimes", &utility.ReportingTimes})
    utility.EntityData.Children.Append("customer-info", types.YChild{"CustomerInfo", &utility.CustomerInfo})
    utility.EntityData.Leafs = types.NewOrderedMap()
    utility.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", utility.Enabled})
    utility.EntityData.Leafs.Append("reporting", types.YLeaf{"Reporting", utility.Reporting})

    utility.EntityData.YListKeys = []string {}

    return &(utility.EntityData)
}

// Licensing_State_StateInfo_Utility_ReportingTimes
// If the product instance is reporting utility data this will
// contain various timing information about that reporting.
type Licensing_State_StateInfo_Utility_ReportingTimes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time the last report was sent. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    LastReportTime interface{}

    // Was the last report successfully sent?. The type is bool.
    LastReportSuccess interface{}

    // Failure message if the last report send failed. The type is string with
    // length: 0..255.
    FailMessage interface{}

    // Time the next report is scheduled to be sent. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    NextReportTime interface{}
}

func (reportingTimes *Licensing_State_StateInfo_Utility_ReportingTimes) GetEntityData() *types.CommonEntityData {
    reportingTimes.EntityData.YFilter = reportingTimes.YFilter
    reportingTimes.EntityData.YangName = "reporting-times"
    reportingTimes.EntityData.BundleName = "cisco_ios_xe"
    reportingTimes.EntityData.ParentYangName = "utility"
    reportingTimes.EntityData.SegmentPath = "reporting-times"
    reportingTimes.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/utility/" + reportingTimes.EntityData.SegmentPath
    reportingTimes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    reportingTimes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    reportingTimes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    reportingTimes.EntityData.Children = types.NewOrderedMap()
    reportingTimes.EntityData.Leafs = types.NewOrderedMap()
    reportingTimes.EntityData.Leafs.Append("last-report-time", types.YLeaf{"LastReportTime", reportingTimes.LastReportTime})
    reportingTimes.EntityData.Leafs.Append("last-report-success", types.YLeaf{"LastReportSuccess", reportingTimes.LastReportSuccess})
    reportingTimes.EntityData.Leafs.Append("fail-message", types.YLeaf{"FailMessage", reportingTimes.FailMessage})
    reportingTimes.EntityData.Leafs.Append("next-report-time", types.YLeaf{"NextReportTime", reportingTimes.NextReportTime})

    reportingTimes.EntityData.YListKeys = []string {}

    return &(reportingTimes.EntityData)
}

// Licensing_State_StateInfo_Utility_CustomerInfo
// Customer address information that will be sent 
// in the utility usage reports.
type Licensing_State_StateInfo_Utility_CustomerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The cisco issued customer id which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    Id interface{}

    // The customer company name which will be included in the utility usage (RUM)
    // report. The type is string with length: 1..250.
    Name interface{}

    // The customer company street address which will be included in the utility
    // usage (RUM) report. The type is string with length: 1..250.
    Street interface{}

    // The customer company city which will be included in the utility usage (RUM)
    // report. The type is string with length: 1..250.
    City interface{}

    // The customer company state which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    State interface{}

    // The customer company country which will be included in the utility usage
    // (RUM) report. The type is string with length: 1..250.
    Country interface{}

    // The customer location specific postal code which will be included in the
    // utility usage (RUM) report. The type is string with length: 1..250.
    PostalCode interface{}
}

func (customerInfo *Licensing_State_StateInfo_Utility_CustomerInfo) GetEntityData() *types.CommonEntityData {
    customerInfo.EntityData.YFilter = customerInfo.YFilter
    customerInfo.EntityData.YangName = "customer-info"
    customerInfo.EntityData.BundleName = "cisco_ios_xe"
    customerInfo.EntityData.ParentYangName = "utility"
    customerInfo.EntityData.SegmentPath = "customer-info"
    customerInfo.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/utility/" + customerInfo.EntityData.SegmentPath
    customerInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    customerInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    customerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    customerInfo.EntityData.Children = types.NewOrderedMap()
    customerInfo.EntityData.Leafs = types.NewOrderedMap()
    customerInfo.EntityData.Leafs.Append("id", types.YLeaf{"Id", customerInfo.Id})
    customerInfo.EntityData.Leafs.Append("name", types.YLeaf{"Name", customerInfo.Name})
    customerInfo.EntityData.Leafs.Append("street", types.YLeaf{"Street", customerInfo.Street})
    customerInfo.EntityData.Leafs.Append("city", types.YLeaf{"City", customerInfo.City})
    customerInfo.EntityData.Leafs.Append("state", types.YLeaf{"State", customerInfo.State})
    customerInfo.EntityData.Leafs.Append("country", types.YLeaf{"Country", customerInfo.Country})
    customerInfo.EntityData.Leafs.Append("postal-code", types.YLeaf{"PostalCode", customerInfo.PostalCode})

    customerInfo.EntityData.YListKeys = []string {}

    return &(customerInfo.EntityData)
}

// Licensing_State_StateInfo_Transport
// State of the transport.
type Licensing_State_StateInfo_Transport struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of communications transport smart licensing is using. The type is
    // TransportTypeEnum.
    TransportType interface{}

    // URLs in use if the transport type is smart.
    UrlSettings Licensing_State_StateInfo_Transport_UrlSettings
}

func (transport *Licensing_State_StateInfo_Transport) GetEntityData() *types.CommonEntityData {
    transport.EntityData.YFilter = transport.YFilter
    transport.EntityData.YangName = "transport"
    transport.EntityData.BundleName = "cisco_ios_xe"
    transport.EntityData.ParentYangName = "state-info"
    transport.EntityData.SegmentPath = "transport"
    transport.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + transport.EntityData.SegmentPath
    transport.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    transport.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    transport.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    transport.EntityData.Children = types.NewOrderedMap()
    transport.EntityData.Children.Append("url-settings", types.YChild{"UrlSettings", &transport.UrlSettings})
    transport.EntityData.Leafs = types.NewOrderedMap()
    transport.EntityData.Leafs.Append("transport-type", types.YLeaf{"TransportType", transport.TransportType})

    transport.EntityData.YListKeys = []string {}

    return &(transport.EntityData)
}

// Licensing_State_StateInfo_Transport_UrlSettings
// URLs in use if the transport type is smart.
type Licensing_State_StateInfo_Transport_UrlSettings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The URL used for registration, authorization and any  other messages not
    // related to utility. The type is string.
    UrlRegistration interface{}

    // The URL used for utility reporting. url-utility and  url-registration may
    // be the same or different. If a satellite is in use then they will probably
    // be the same. The type is string.
    UrlUtility interface{}
}

func (urlSettings *Licensing_State_StateInfo_Transport_UrlSettings) GetEntityData() *types.CommonEntityData {
    urlSettings.EntityData.YFilter = urlSettings.YFilter
    urlSettings.EntityData.YangName = "url-settings"
    urlSettings.EntityData.BundleName = "cisco_ios_xe"
    urlSettings.EntityData.ParentYangName = "transport"
    urlSettings.EntityData.SegmentPath = "url-settings"
    urlSettings.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/transport/" + urlSettings.EntityData.SegmentPath
    urlSettings.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    urlSettings.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    urlSettings.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    urlSettings.EntityData.Children = types.NewOrderedMap()
    urlSettings.EntityData.Leafs = types.NewOrderedMap()
    urlSettings.EntityData.Leafs.Append("url-registration", types.YLeaf{"UrlRegistration", urlSettings.UrlRegistration})
    urlSettings.EntityData.Leafs.Append("url-utility", types.YLeaf{"UrlUtility", urlSettings.UrlUtility})

    urlSettings.EntityData.YListKeys = []string {}

    return &(urlSettings.EntityData)
}

// Licensing_State_StateInfo_Privacy
// State of the privacy settings.
type Licensing_State_StateInfo_Privacy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If true then the hostname will not be sent in  any messages. The type is
    // bool.
    Hostname interface{}

    // If true then the smart licensing version will not be sent in  any messages.
    // The type is bool.
    Version interface{}
}

func (privacy *Licensing_State_StateInfo_Privacy) GetEntityData() *types.CommonEntityData {
    privacy.EntityData.YFilter = privacy.YFilter
    privacy.EntityData.YangName = "privacy"
    privacy.EntityData.BundleName = "cisco_ios_xe"
    privacy.EntityData.ParentYangName = "state-info"
    privacy.EntityData.SegmentPath = "privacy"
    privacy.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + privacy.EntityData.SegmentPath
    privacy.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    privacy.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    privacy.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    privacy.EntityData.Children = types.NewOrderedMap()
    privacy.EntityData.Leafs = types.NewOrderedMap()
    privacy.EntityData.Leafs.Append("hostname", types.YLeaf{"Hostname", privacy.Hostname})
    privacy.EntityData.Leafs.Append("version", types.YLeaf{"Version", privacy.Version})

    privacy.EntityData.YListKeys = []string {}

    return &(privacy.EntityData)
}

// Licensing_State_StateInfo_Evaluation
// State of the evaluation period.
type Licensing_State_StateInfo_Evaluation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is the evaluation period currently in use and counting down. The type is
    // bool.
    EvalInUse interface{}

    // Has the evaluation period expired. The type is bool.
    EvalExpired interface{}

    // If the evaluation period is not expired this is the number  of seconds left
    // in the period.
    EvalPeriodLeft Licensing_State_StateInfo_Evaluation_EvalPeriodLeft

    // If the evaluation period has expired then this is the  time of the
    // expiration.
    EvalExpireTime Licensing_State_StateInfo_Evaluation_EvalExpireTime
}

func (evaluation *Licensing_State_StateInfo_Evaluation) GetEntityData() *types.CommonEntityData {
    evaluation.EntityData.YFilter = evaluation.YFilter
    evaluation.EntityData.YangName = "evaluation"
    evaluation.EntityData.BundleName = "cisco_ios_xe"
    evaluation.EntityData.ParentYangName = "state-info"
    evaluation.EntityData.SegmentPath = "evaluation"
    evaluation.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + evaluation.EntityData.SegmentPath
    evaluation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    evaluation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    evaluation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    evaluation.EntityData.Children = types.NewOrderedMap()
    evaluation.EntityData.Children.Append("eval-period-left", types.YChild{"EvalPeriodLeft", &evaluation.EvalPeriodLeft})
    evaluation.EntityData.Children.Append("eval-expire-time", types.YChild{"EvalExpireTime", &evaluation.EvalExpireTime})
    evaluation.EntityData.Leafs = types.NewOrderedMap()
    evaluation.EntityData.Leafs.Append("eval-in-use", types.YLeaf{"EvalInUse", evaluation.EvalInUse})
    evaluation.EntityData.Leafs.Append("eval-expired", types.YLeaf{"EvalExpired", evaluation.EvalExpired})

    evaluation.EntityData.YListKeys = []string {}

    return &(evaluation.EntityData)
}

// Licensing_State_StateInfo_Evaluation_EvalPeriodLeft
// If the evaluation period is not expired this is the number 
// of seconds left in the period.
type Licensing_State_StateInfo_Evaluation_EvalPeriodLeft struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds of license usage until the evaluation  period expires.
    // The type is interface{} with range: 0..4294967295.
    TimeLeft interface{}
}

func (evalPeriodLeft *Licensing_State_StateInfo_Evaluation_EvalPeriodLeft) GetEntityData() *types.CommonEntityData {
    evalPeriodLeft.EntityData.YFilter = evalPeriodLeft.YFilter
    evalPeriodLeft.EntityData.YangName = "eval-period-left"
    evalPeriodLeft.EntityData.BundleName = "cisco_ios_xe"
    evalPeriodLeft.EntityData.ParentYangName = "evaluation"
    evalPeriodLeft.EntityData.SegmentPath = "eval-period-left"
    evalPeriodLeft.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/evaluation/" + evalPeriodLeft.EntityData.SegmentPath
    evalPeriodLeft.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    evalPeriodLeft.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    evalPeriodLeft.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    evalPeriodLeft.EntityData.Children = types.NewOrderedMap()
    evalPeriodLeft.EntityData.Leafs = types.NewOrderedMap()
    evalPeriodLeft.EntityData.Leafs.Append("time-left", types.YLeaf{"TimeLeft", evalPeriodLeft.TimeLeft})

    evalPeriodLeft.EntityData.YListKeys = []string {}

    return &(evalPeriodLeft.EntityData)
}

// Licensing_State_StateInfo_Evaluation_EvalExpireTime
// If the evaluation period has expired then this is the 
// time of the expiration.
type Licensing_State_StateInfo_Evaluation_EvalExpireTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Date and time the evaluation period expired. The type is string with
    // pattern: \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    ExpireTime interface{}
}

func (evalExpireTime *Licensing_State_StateInfo_Evaluation_EvalExpireTime) GetEntityData() *types.CommonEntityData {
    evalExpireTime.EntityData.YFilter = evalExpireTime.YFilter
    evalExpireTime.EntityData.YangName = "eval-expire-time"
    evalExpireTime.EntityData.BundleName = "cisco_ios_xe"
    evalExpireTime.EntityData.ParentYangName = "evaluation"
    evalExpireTime.EntityData.SegmentPath = "eval-expire-time"
    evalExpireTime.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/evaluation/" + evalExpireTime.EntityData.SegmentPath
    evalExpireTime.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    evalExpireTime.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    evalExpireTime.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    evalExpireTime.EntityData.Children = types.NewOrderedMap()
    evalExpireTime.EntityData.Leafs = types.NewOrderedMap()
    evalExpireTime.EntityData.Leafs.Append("expire-time", types.YLeaf{"ExpireTime", evalExpireTime.ExpireTime})

    evalExpireTime.EntityData.YListKeys = []string {}

    return &(evalExpireTime.EntityData)
}

// Licensing_State_StateInfo_Udi
// UDI of the system.
type Licensing_State_StateInfo_Udi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The Product Identifier. Always combined with sn. The type is string with
    // length: 1..255.
    Pid interface{}

    // The system serial number. Always combined with pid. The type is string with
    // length: 1..255.
    Sn interface{}

    // The version identifier. Usually combined with pid & sn. The type is string
    // with length: 1..255.
    Vid interface{}

    // A 32 byte hex value generated by the system.  This will be in proper UUID
    // format 8-4-4-4-12. Often used by VMs or other systems that do not have a
    // hardware identifier. The type is string with pattern:
    // [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}.
    Uuid interface{}

    // Free form virtual identifier often used by software  only devices like
    // software routers or VMs. The type is string with length: 1..255.
    Suvi interface{}

    // Host identifier available on some systems.  Typically 8 hex digits. The
    // type is string with length: 1..255.
    HostIdentifier interface{}

    // The MAC address of the system. This is usually only used if there  is
    // nothing else available to be used as an identifier. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}
}

func (udi *Licensing_State_StateInfo_Udi) GetEntityData() *types.CommonEntityData {
    udi.EntityData.YFilter = udi.YFilter
    udi.EntityData.YangName = "udi"
    udi.EntityData.BundleName = "cisco_ios_xe"
    udi.EntityData.ParentYangName = "state-info"
    udi.EntityData.SegmentPath = "udi"
    udi.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + udi.EntityData.SegmentPath
    udi.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    udi.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    udi.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    udi.EntityData.Children = types.NewOrderedMap()
    udi.EntityData.Leafs = types.NewOrderedMap()
    udi.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", udi.Pid})
    udi.EntityData.Leafs.Append("sn", types.YLeaf{"Sn", udi.Sn})
    udi.EntityData.Leafs.Append("vid", types.YLeaf{"Vid", udi.Vid})
    udi.EntityData.Leafs.Append("uuid", types.YLeaf{"Uuid", udi.Uuid})
    udi.EntityData.Leafs.Append("suvi", types.YLeaf{"Suvi", udi.Suvi})
    udi.EntityData.Leafs.Append("host-identifier", types.YLeaf{"HostIdentifier", udi.HostIdentifier})
    udi.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", udi.MacAddress})

    udi.EntityData.YListKeys = []string {}

    return &(udi.EntityData)
}

// Licensing_State_StateInfo_Usage
// List of license (entitlement tag) usage information. 
// This only contains the information for licenses that are in use.
type Licensing_State_StateInfo_Usage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The ISO 19770 entitlement tag used to define this
    // license. The type is string.
    EntitlementTag interface{}

    // The human readable license name from the entitlement tag. The type is
    // string.
    ShortName interface{}

    // The license name that can be seen in the CSSM portal or on  the satellite. 
    // This is only available after the product has registered. The type is
    // string.
    LicenseName interface{}

    // The long description of this license. This is only available after the
    // product has registered. The type is string.
    Description interface{}

    // The in-use count of this license. Note that licensing only reports  usage
    // for licenses that are in use (count of 1 or greater). The type is
    // interface{} with range: 0..4294967295.
    Count interface{}

    // The current enforcement mode of this license. The type is
    // EnforcementModeEnum.
    EnforcementMode interface{}

    // If true then this entitlement is being reported in a post paid mode with
    // utility usage reports. Otherwise it will be reported in  the regular
    // prepaid mode. The type is bool.
    PostPaid interface{}

    // If this entitlement tag is being reported in the post paid utility usage
    // mode and there is a subscription id in the customer's  virtual account this
    // will be that subscription id. The type is string.
    SubscriptionId interface{}
}

func (usage *Licensing_State_StateInfo_Usage) GetEntityData() *types.CommonEntityData {
    usage.EntityData.YFilter = usage.YFilter
    usage.EntityData.YangName = "usage"
    usage.EntityData.BundleName = "cisco_ios_xe"
    usage.EntityData.ParentYangName = "state-info"
    usage.EntityData.SegmentPath = "usage" + types.AddKeyToken(usage.EntitlementTag, "entitlement-tag")
    usage.EntityData.AbsolutePath = "cisco-smart-license:licensing/state/state-info/" + usage.EntityData.SegmentPath
    usage.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    usage.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    usage.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    usage.EntityData.Children = types.NewOrderedMap()
    usage.EntityData.Leafs = types.NewOrderedMap()
    usage.EntityData.Leafs.Append("entitlement-tag", types.YLeaf{"EntitlementTag", usage.EntitlementTag})
    usage.EntityData.Leafs.Append("short-name", types.YLeaf{"ShortName", usage.ShortName})
    usage.EntityData.Leafs.Append("license-name", types.YLeaf{"LicenseName", usage.LicenseName})
    usage.EntityData.Leafs.Append("description", types.YLeaf{"Description", usage.Description})
    usage.EntityData.Leafs.Append("count", types.YLeaf{"Count", usage.Count})
    usage.EntityData.Leafs.Append("enforcement-mode", types.YLeaf{"EnforcementMode", usage.EnforcementMode})
    usage.EntityData.Leafs.Append("post-paid", types.YLeaf{"PostPaid", usage.PostPaid})
    usage.EntityData.Leafs.Append("subscription-id", types.YLeaf{"SubscriptionId", usage.SubscriptionId})

    usage.EntityData.YListKeys = []string {"EntitlementTag"}

    return &(usage.EntityData)
}

