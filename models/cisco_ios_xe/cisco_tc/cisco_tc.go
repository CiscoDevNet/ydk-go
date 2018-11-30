// This module defines textual conventions used throughout
// cisco enterprise mibs.
package cisco_tc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_tc"))
}

// CiscoPortListRange represents syntax.
type CiscoPortListRange string

const (
    CiscoPortListRange_oneto2k CiscoPortListRange = "oneto2k"

    CiscoPortListRange_twoKto4K CiscoPortListRange = "twoKto4K"

    CiscoPortListRange_fourKto6K CiscoPortListRange = "fourKto6K"

    CiscoPortListRange_sixKto8K CiscoPortListRange = "sixKto8K"

    CiscoPortListRange_eightKto10K CiscoPortListRange = "eightKto10K"

    CiscoPortListRange_tenKto12K CiscoPortListRange = "tenKto12K"

    CiscoPortListRange_twelveKto14K CiscoPortListRange = "twelveKto14K"

    CiscoPortListRange_fourteenKto16K CiscoPortListRange = "fourteenKto16K"
)

// CiscoNetworkProtocol represents Represents the different types of network layer protocols.
type CiscoNetworkProtocol string

const (
    CiscoNetworkProtocol_ip CiscoNetworkProtocol = "ip"

    CiscoNetworkProtocol_decnet CiscoNetworkProtocol = "decnet"

    CiscoNetworkProtocol_pup CiscoNetworkProtocol = "pup"

    CiscoNetworkProtocol_chaos CiscoNetworkProtocol = "chaos"

    CiscoNetworkProtocol_xns CiscoNetworkProtocol = "xns"

    CiscoNetworkProtocol_x121 CiscoNetworkProtocol = "x121"

    CiscoNetworkProtocol_appletalk CiscoNetworkProtocol = "appletalk"

    CiscoNetworkProtocol_clns CiscoNetworkProtocol = "clns"

    CiscoNetworkProtocol_lat CiscoNetworkProtocol = "lat"

    CiscoNetworkProtocol_vines CiscoNetworkProtocol = "vines"

    CiscoNetworkProtocol_cons CiscoNetworkProtocol = "cons"

    CiscoNetworkProtocol_apollo CiscoNetworkProtocol = "apollo"

    CiscoNetworkProtocol_stun CiscoNetworkProtocol = "stun"

    CiscoNetworkProtocol_novell CiscoNetworkProtocol = "novell"

    CiscoNetworkProtocol_qllc CiscoNetworkProtocol = "qllc"

    CiscoNetworkProtocol_snapshot CiscoNetworkProtocol = "snapshot"

    CiscoNetworkProtocol_atmIlmi CiscoNetworkProtocol = "atmIlmi"

    CiscoNetworkProtocol_bstun CiscoNetworkProtocol = "bstun"

    CiscoNetworkProtocol_x25pvc CiscoNetworkProtocol = "x25pvc"

    CiscoNetworkProtocol_ipv6 CiscoNetworkProtocol = "ipv6"

    CiscoNetworkProtocol_cdm CiscoNetworkProtocol = "cdm"

    CiscoNetworkProtocol_nbf CiscoNetworkProtocol = "nbf"

    CiscoNetworkProtocol_bpxIgx CiscoNetworkProtocol = "bpxIgx"

    CiscoNetworkProtocol_clnsPfx CiscoNetworkProtocol = "clnsPfx"

    CiscoNetworkProtocol_http CiscoNetworkProtocol = "http"

    CiscoNetworkProtocol_unknown CiscoNetworkProtocol = "unknown"
)

// CiscoRowOperStatus represents     not exist in it's table.
type CiscoRowOperStatus string

const (
    CiscoRowOperStatus_active CiscoRowOperStatus = "active"

    CiscoRowOperStatus_activeDependencies CiscoRowOperStatus = "activeDependencies"

    CiscoRowOperStatus_inactiveDependency CiscoRowOperStatus = "inactiveDependency"

    CiscoRowOperStatus_missingDependency CiscoRowOperStatus = "missingDependency"
)

// CiscoLocationClass represents subChannel - a sub-channel on a logical interface.
type CiscoLocationClass string

const (
    CiscoLocationClass_chassis CiscoLocationClass = "chassis"

    CiscoLocationClass_shelf CiscoLocationClass = "shelf"

    CiscoLocationClass_slot CiscoLocationClass = "slot"

    CiscoLocationClass_subSlot CiscoLocationClass = "subSlot"

    CiscoLocationClass_port CiscoLocationClass = "port"

    CiscoLocationClass_subPort CiscoLocationClass = "subPort"

    CiscoLocationClass_channel CiscoLocationClass = "channel"

    CiscoLocationClass_subChannel CiscoLocationClass = "subChannel"
)

// IfOperStatusReason represents      'fabricNameInvalid (268)' - Fabric Name Invalid.
type IfOperStatusReason string

const (
    IfOperStatusReason_other IfOperStatusReason = "other"

    IfOperStatusReason_none IfOperStatusReason = "none"

    IfOperStatusReason_hwFailure IfOperStatusReason = "hwFailure"

    IfOperStatusReason_loopbackDiagFailure IfOperStatusReason = "loopbackDiagFailure"

    IfOperStatusReason_errorDisabled IfOperStatusReason = "errorDisabled"

    IfOperStatusReason_swFailure IfOperStatusReason = "swFailure"

    IfOperStatusReason_linkFailure IfOperStatusReason = "linkFailure"

    IfOperStatusReason_offline IfOperStatusReason = "offline"

    IfOperStatusReason_nonParticipating IfOperStatusReason = "nonParticipating"

    IfOperStatusReason_initializing IfOperStatusReason = "initializing"

    IfOperStatusReason_vsanInactive IfOperStatusReason = "vsanInactive"

    IfOperStatusReason_adminDown IfOperStatusReason = "adminDown"

    IfOperStatusReason_channelAdminDown IfOperStatusReason = "channelAdminDown"

    IfOperStatusReason_channelOperSuspended IfOperStatusReason = "channelOperSuspended"

    IfOperStatusReason_channelConfigurationInProgress IfOperStatusReason = "channelConfigurationInProgress"

    IfOperStatusReason_rcfInProgress IfOperStatusReason = "rcfInProgress"

    IfOperStatusReason_elpFailureIsolation IfOperStatusReason = "elpFailureIsolation"

    IfOperStatusReason_escFailureIsolation IfOperStatusReason = "escFailureIsolation"

    IfOperStatusReason_domainOverlapIsolation IfOperStatusReason = "domainOverlapIsolation"

    IfOperStatusReason_domainAddrAssignFailureIsolation IfOperStatusReason = "domainAddrAssignFailureIsolation"

    IfOperStatusReason_domainOtherSideEportIsolation IfOperStatusReason = "domainOtherSideEportIsolation"

    IfOperStatusReason_domainInvalidRcfReceived IfOperStatusReason = "domainInvalidRcfReceived"

    IfOperStatusReason_domainManagerDisabled IfOperStatusReason = "domainManagerDisabled"

    IfOperStatusReason_zoneMergeFailureIsolation IfOperStatusReason = "zoneMergeFailureIsolation"

    IfOperStatusReason_vsanMismatchIsolation IfOperStatusReason = "vsanMismatchIsolation"

    IfOperStatusReason_parentDown IfOperStatusReason = "parentDown"

    IfOperStatusReason_srcPortNotBound IfOperStatusReason = "srcPortNotBound"

    IfOperStatusReason_interfaceRemoved IfOperStatusReason = "interfaceRemoved"

    IfOperStatusReason_fcotNotPresent IfOperStatusReason = "fcotNotPresent"

    IfOperStatusReason_fcotVendorNotSupported IfOperStatusReason = "fcotVendorNotSupported"

    IfOperStatusReason_incompatibleAdminMode IfOperStatusReason = "incompatibleAdminMode"

    IfOperStatusReason_incompatibleAdminSpeed IfOperStatusReason = "incompatibleAdminSpeed"

    IfOperStatusReason_suspendedByMode IfOperStatusReason = "suspendedByMode"

    IfOperStatusReason_suspendedBySpeed IfOperStatusReason = "suspendedBySpeed"

    IfOperStatusReason_suspendedByWWN IfOperStatusReason = "suspendedByWWN"

    IfOperStatusReason_domainMaxReTxFailure IfOperStatusReason = "domainMaxReTxFailure"

    IfOperStatusReason_eppFailure IfOperStatusReason = "eppFailure"

    IfOperStatusReason_portVsanMismatchIsolation IfOperStatusReason = "portVsanMismatchIsolation"

    IfOperStatusReason_loopbackIsolation IfOperStatusReason = "loopbackIsolation"

    IfOperStatusReason_upgradeInProgress IfOperStatusReason = "upgradeInProgress"

    IfOperStatusReason_incompatibleAdminRxBbCredit IfOperStatusReason = "incompatibleAdminRxBbCredit"

    IfOperStatusReason_incompatibleAdminRxBufferSize IfOperStatusReason = "incompatibleAdminRxBufferSize"

    IfOperStatusReason_portChannelMembersDown IfOperStatusReason = "portChannelMembersDown"

    IfOperStatusReason_zoneRemoteNoRespIsolation IfOperStatusReason = "zoneRemoteNoRespIsolation"

    IfOperStatusReason_firstPortUpAsEport IfOperStatusReason = "firstPortUpAsEport"

    IfOperStatusReason_firstPortNotUp IfOperStatusReason = "firstPortNotUp"

    IfOperStatusReason_peerFCIPPortClosedConnection IfOperStatusReason = "peerFCIPPortClosedConnection"

    IfOperStatusReason_peerFCIPPortResetConnection IfOperStatusReason = "peerFCIPPortResetConnection"

    IfOperStatusReason_fcipPortMaxReTx IfOperStatusReason = "fcipPortMaxReTx"

    IfOperStatusReason_fcipPortKeepAliveTimerExpire IfOperStatusReason = "fcipPortKeepAliveTimerExpire"

    IfOperStatusReason_fcipPortPersistTimerExpire IfOperStatusReason = "fcipPortPersistTimerExpire"

    IfOperStatusReason_fcipPortSrcLinkDown IfOperStatusReason = "fcipPortSrcLinkDown"

    IfOperStatusReason_fcipPortSrcAdminDown IfOperStatusReason = "fcipPortSrcAdminDown"

    IfOperStatusReason_fcipPortAdminCfgChange IfOperStatusReason = "fcipPortAdminCfgChange"

    IfOperStatusReason_fcipSrcPortRemoved IfOperStatusReason = "fcipSrcPortRemoved"

    IfOperStatusReason_fcipSrcModuleNotOnline IfOperStatusReason = "fcipSrcModuleNotOnline"

    IfOperStatusReason_invalidConfig IfOperStatusReason = "invalidConfig"

    IfOperStatusReason_portBindFailure IfOperStatusReason = "portBindFailure"

    IfOperStatusReason_portFabricBindFailure IfOperStatusReason = "portFabricBindFailure"

    IfOperStatusReason_noCommonVsanIsolation IfOperStatusReason = "noCommonVsanIsolation"

    IfOperStatusReason_ficonVsanDown IfOperStatusReason = "ficonVsanDown"

    IfOperStatusReason_invalidAttachment IfOperStatusReason = "invalidAttachment"

    IfOperStatusReason_portBlocked IfOperStatusReason = "portBlocked"

    IfOperStatusReason_incomAdminRxBbCreditPerBuf IfOperStatusReason = "incomAdminRxBbCreditPerBuf"

    IfOperStatusReason_tooManyInvalidFlogis IfOperStatusReason = "tooManyInvalidFlogis"

    IfOperStatusReason_deniedDueToPortBinding IfOperStatusReason = "deniedDueToPortBinding"

    IfOperStatusReason_elpFailureRevMismatch IfOperStatusReason = "elpFailureRevMismatch"

    IfOperStatusReason_elpFailureClassFParamErr IfOperStatusReason = "elpFailureClassFParamErr"

    IfOperStatusReason_elpFailureClassNParamErr IfOperStatusReason = "elpFailureClassNParamErr"

    IfOperStatusReason_elpFailureUnknownFlowCtlCode IfOperStatusReason = "elpFailureUnknownFlowCtlCode"

    IfOperStatusReason_elpFailureInvalidFlowCtlParam IfOperStatusReason = "elpFailureInvalidFlowCtlParam"

    IfOperStatusReason_elpFailureInvalidPortName IfOperStatusReason = "elpFailureInvalidPortName"

    IfOperStatusReason_elpFailureInvalidSwitchName IfOperStatusReason = "elpFailureInvalidSwitchName"

    IfOperStatusReason_elpFailureRatovEdtovMismatch IfOperStatusReason = "elpFailureRatovEdtovMismatch"

    IfOperStatusReason_elpFailureLoopbackDetected IfOperStatusReason = "elpFailureLoopbackDetected"

    IfOperStatusReason_elpFailureInvalidTxBbCredit IfOperStatusReason = "elpFailureInvalidTxBbCredit"

    IfOperStatusReason_elpFailureInvalidPayloadSize IfOperStatusReason = "elpFailureInvalidPayloadSize"

    IfOperStatusReason_bundleMisCfg IfOperStatusReason = "bundleMisCfg"

    IfOperStatusReason_bitErrRuntimeThreshExceeded IfOperStatusReason = "bitErrRuntimeThreshExceeded"

    IfOperStatusReason_linkFailLinkReset IfOperStatusReason = "linkFailLinkReset"

    IfOperStatusReason_linkFailPortInitFail IfOperStatusReason = "linkFailPortInitFail"

    IfOperStatusReason_linkFailPortUnusable IfOperStatusReason = "linkFailPortUnusable"

    IfOperStatusReason_linkFailLossOfSignal IfOperStatusReason = "linkFailLossOfSignal"

    IfOperStatusReason_linkFailLossOfSync IfOperStatusReason = "linkFailLossOfSync"

    IfOperStatusReason_linkFailNosRcvd IfOperStatusReason = "linkFailNosRcvd"

    IfOperStatusReason_linkFailOLSRcvd IfOperStatusReason = "linkFailOLSRcvd"

    IfOperStatusReason_linkFailDebounceTimeout IfOperStatusReason = "linkFailDebounceTimeout"

    IfOperStatusReason_linkFailLrRcvd IfOperStatusReason = "linkFailLrRcvd"

    IfOperStatusReason_linkFailCreditLoss IfOperStatusReason = "linkFailCreditLoss"

    IfOperStatusReason_linkFailRxQOverflow IfOperStatusReason = "linkFailRxQOverflow"

    IfOperStatusReason_linkFailTooManyInterrupts IfOperStatusReason = "linkFailTooManyInterrupts"

    IfOperStatusReason_linkFailLipRcvdBb IfOperStatusReason = "linkFailLipRcvdBb"

    IfOperStatusReason_linkFailBbCreditLoss IfOperStatusReason = "linkFailBbCreditLoss"

    IfOperStatusReason_linkFailOpenPrimSignalTimeout IfOperStatusReason = "linkFailOpenPrimSignalTimeout"

    IfOperStatusReason_linkFailOpenPrimSignalReturned IfOperStatusReason = "linkFailOpenPrimSignalReturned"

    IfOperStatusReason_linkFailLipF8Rcvd IfOperStatusReason = "linkFailLipF8Rcvd"

    IfOperStatusReason_linkFailLineCardPortShutdown IfOperStatusReason = "linkFailLineCardPortShutdown"

    IfOperStatusReason_fcspAuthenfailure IfOperStatusReason = "fcspAuthenfailure"

    IfOperStatusReason_fcotChecksumError IfOperStatusReason = "fcotChecksumError"

    IfOperStatusReason_ohmsExtLoopbackTest IfOperStatusReason = "ohmsExtLoopbackTest"

    IfOperStatusReason_invalidFabricBindExchange IfOperStatusReason = "invalidFabricBindExchange"

    IfOperStatusReason_tovMismatch IfOperStatusReason = "tovMismatch"

    IfOperStatusReason_ficonNotEnabled IfOperStatusReason = "ficonNotEnabled"

    IfOperStatusReason_ficonNoPortNumber IfOperStatusReason = "ficonNoPortNumber"

    IfOperStatusReason_ficonBeingEnabled IfOperStatusReason = "ficonBeingEnabled"

    IfOperStatusReason_ePortProhibited IfOperStatusReason = "ePortProhibited"

    IfOperStatusReason_portGracefulShutdown IfOperStatusReason = "portGracefulShutdown"

    IfOperStatusReason_trunkNotFullyActive IfOperStatusReason = "trunkNotFullyActive"

    IfOperStatusReason_fabricBindingSwitchWwnNotFound IfOperStatusReason = "fabricBindingSwitchWwnNotFound"

    IfOperStatusReason_fabricBindingDomainInvalid IfOperStatusReason = "fabricBindingDomainInvalid"

    IfOperStatusReason_fabricBindingDbMismatch IfOperStatusReason = "fabricBindingDbMismatch"

    IfOperStatusReason_fabricBindingNoRspFromPeer IfOperStatusReason = "fabricBindingNoRspFromPeer"

    IfOperStatusReason_dpvmVsanSuspended IfOperStatusReason = "dpvmVsanSuspended"

    IfOperStatusReason_dpvmVsanNotFound IfOperStatusReason = "dpvmVsanNotFound"

    IfOperStatusReason_trackedPortDown IfOperStatusReason = "trackedPortDown"

    IfOperStatusReason_ecSuspendedOnLoop IfOperStatusReason = "ecSuspendedOnLoop"

    IfOperStatusReason_isolateBundleMisCfg IfOperStatusReason = "isolateBundleMisCfg"

    IfOperStatusReason_noPeerBundleSupport IfOperStatusReason = "noPeerBundleSupport"

    IfOperStatusReason_portBringupIsolation IfOperStatusReason = "portBringupIsolation"

    IfOperStatusReason_domainNotAllowedIsolated IfOperStatusReason = "domainNotAllowedIsolated"

    IfOperStatusReason_virtualIvrDomainOverlapIsolation IfOperStatusReason = "virtualIvrDomainOverlapIsolation"

    IfOperStatusReason_outOfService IfOperStatusReason = "outOfService"

    IfOperStatusReason_portAuthFailed IfOperStatusReason = "portAuthFailed"

    IfOperStatusReason_bundleStandby IfOperStatusReason = "bundleStandby"

    IfOperStatusReason_portConnectorTypeErr IfOperStatusReason = "portConnectorTypeErr"

    IfOperStatusReason_errorDisabledReInitLmtReached IfOperStatusReason = "errorDisabledReInitLmtReached"

    IfOperStatusReason_ficonDupPortNum IfOperStatusReason = "ficonDupPortNum"

    IfOperStatusReason_localRcf IfOperStatusReason = "localRcf"

    IfOperStatusReason_twoSwitchesWithSameWWN IfOperStatusReason = "twoSwitchesWithSameWWN"

    IfOperStatusReason_invalidOtherSidePrincEFPReqRecd IfOperStatusReason = "invalidOtherSidePrincEFPReqRecd"

    IfOperStatusReason_domainOther IfOperStatusReason = "domainOther"

    IfOperStatusReason_elpFailureAllZeroPeerWWNRcvd IfOperStatusReason = "elpFailureAllZeroPeerWWNRcvd"

    IfOperStatusReason_preferredPathIsolation IfOperStatusReason = "preferredPathIsolation"

    IfOperStatusReason_fcRedirectIsolation IfOperStatusReason = "fcRedirectIsolation"

    IfOperStatusReason_portActLicenseNotAvailable IfOperStatusReason = "portActLicenseNotAvailable"

    IfOperStatusReason_sdmIsolation IfOperStatusReason = "sdmIsolation"

    IfOperStatusReason_fcidAllocationFailed IfOperStatusReason = "fcidAllocationFailed"

    IfOperStatusReason_externallyDisabled IfOperStatusReason = "externallyDisabled"

    IfOperStatusReason_unavailableNPVUpstreamPort IfOperStatusReason = "unavailableNPVUpstreamPort"

    IfOperStatusReason_unavailableNPVPrefUpstreamPort IfOperStatusReason = "unavailableNPVPrefUpstreamPort"

    IfOperStatusReason_sfpReadError IfOperStatusReason = "sfpReadError"

    IfOperStatusReason_stickyDownOnLinkFailure IfOperStatusReason = "stickyDownOnLinkFailure"

    IfOperStatusReason_tooManyLinkFlapsErr IfOperStatusReason = "tooManyLinkFlapsErr"

    IfOperStatusReason_unidirectionalUDLD IfOperStatusReason = "unidirectionalUDLD"

    IfOperStatusReason_txRxLoopUDLD IfOperStatusReason = "txRxLoopUDLD"

    IfOperStatusReason_neighborMismatchUDLD IfOperStatusReason = "neighborMismatchUDLD"

    IfOperStatusReason_authzPending IfOperStatusReason = "authzPending"

    IfOperStatusReason_hotStdbyInBundle IfOperStatusReason = "hotStdbyInBundle"

    IfOperStatusReason_incompleteConfig IfOperStatusReason = "incompleteConfig"

    IfOperStatusReason_incompleteTunnelCfg IfOperStatusReason = "incompleteTunnelCfg"

    IfOperStatusReason_hwProgrammingFailed IfOperStatusReason = "hwProgrammingFailed"

    IfOperStatusReason_tunnelDstUnreachable IfOperStatusReason = "tunnelDstUnreachable"

    IfOperStatusReason_suspendByMtu IfOperStatusReason = "suspendByMtu"

    IfOperStatusReason_sfpInvalid IfOperStatusReason = "sfpInvalid"

    IfOperStatusReason_sfpAbsent IfOperStatusReason = "sfpAbsent"

    IfOperStatusReason_portCapabilitiesUnknown IfOperStatusReason = "portCapabilitiesUnknown"

    IfOperStatusReason_channelErrDisabled IfOperStatusReason = "channelErrDisabled"

    IfOperStatusReason_vrfMismatch IfOperStatusReason = "vrfMismatch"

    IfOperStatusReason_vrfForwardReferencing IfOperStatusReason = "vrfForwardReferencing"

    IfOperStatusReason_dupTunnelConfigDetected IfOperStatusReason = "dupTunnelConfigDetected"

    IfOperStatusReason_primaryVLANDown IfOperStatusReason = "primaryVLANDown"

    IfOperStatusReason_vrfUnusable IfOperStatusReason = "vrfUnusable"

    IfOperStatusReason_errDisableHandShkFailure IfOperStatusReason = "errDisableHandShkFailure"

    IfOperStatusReason_errDisabledBPDUGuard IfOperStatusReason = "errDisabledBPDUGuard"

    IfOperStatusReason_dot1xSecViolationErrDisabled IfOperStatusReason = "dot1xSecViolationErrDisabled"

    IfOperStatusReason_emptyEchoUDLD IfOperStatusReason = "emptyEchoUDLD"

    IfOperStatusReason_vfTaggingCapErr IfOperStatusReason = "vfTaggingCapErr"

    IfOperStatusReason_portDisabled IfOperStatusReason = "portDisabled"

    IfOperStatusReason_tunnelModeNotConfigured IfOperStatusReason = "tunnelModeNotConfigured"

    IfOperStatusReason_tunnelSrcNotConfigured IfOperStatusReason = "tunnelSrcNotConfigured"

    IfOperStatusReason_tunnelDstNotConfigured IfOperStatusReason = "tunnelDstNotConfigured"

    IfOperStatusReason_tunnelUnableToResolveSrcIP IfOperStatusReason = "tunnelUnableToResolveSrcIP"

    IfOperStatusReason_tunnelUnableToResolveDstIP IfOperStatusReason = "tunnelUnableToResolveDstIP"

    IfOperStatusReason_tunnelVrfDown IfOperStatusReason = "tunnelVrfDown"

    IfOperStatusReason_sifAdminDown IfOperStatusReason = "sifAdminDown"

    IfOperStatusReason_phyIntfDown IfOperStatusReason = "phyIntfDown"

    IfOperStatusReason_ifSifLimitExceeded IfOperStatusReason = "ifSifLimitExceeded"

    IfOperStatusReason_sifHoldDown IfOperStatusReason = "sifHoldDown"

    IfOperStatusReason_noFcoe IfOperStatusReason = "noFcoe"

    IfOperStatusReason_dcxCompatMismatch IfOperStatusReason = "dcxCompatMismatch"

    IfOperStatusReason_isolateBundleLimitExceeded IfOperStatusReason = "isolateBundleLimitExceeded"

    IfOperStatusReason_sifNotBound IfOperStatusReason = "sifNotBound"

    IfOperStatusReason_errDisabledLacpMiscfg IfOperStatusReason = "errDisabledLacpMiscfg"

    IfOperStatusReason_satFabricIfDown IfOperStatusReason = "satFabricIfDown"

    IfOperStatusReason_invalidSatFabricIf IfOperStatusReason = "invalidSatFabricIf"

    IfOperStatusReason_noRemoteChassis IfOperStatusReason = "noRemoteChassis"

    IfOperStatusReason_vicEnableNotReceived IfOperStatusReason = "vicEnableNotReceived"

    IfOperStatusReason_vicDisableReceived IfOperStatusReason = "vicDisableReceived"

    IfOperStatusReason_vlanVsanMappingNotEnabled IfOperStatusReason = "vlanVsanMappingNotEnabled"

    IfOperStatusReason_stpNotFwdingInFcoeMappedVlan IfOperStatusReason = "stpNotFwdingInFcoeMappedVlan"

    IfOperStatusReason_moduleOffline IfOperStatusReason = "moduleOffline"

    IfOperStatusReason_udldAggModeLinkFailure IfOperStatusReason = "udldAggModeLinkFailure"

    IfOperStatusReason_stpInconsVpcPeerDisabled IfOperStatusReason = "stpInconsVpcPeerDisabled"

    IfOperStatusReason_setPortStateFailed IfOperStatusReason = "setPortStateFailed"

    IfOperStatusReason_suspendedByVpc IfOperStatusReason = "suspendedByVpc"

    IfOperStatusReason_vpcCfgInProgress IfOperStatusReason = "vpcCfgInProgress"

    IfOperStatusReason_vpcPeerLinkDown IfOperStatusReason = "vpcPeerLinkDown"

    IfOperStatusReason_vpcNoRspFromPeer IfOperStatusReason = "vpcNoRspFromPeer"

    IfOperStatusReason_protoPortSuspend IfOperStatusReason = "protoPortSuspend"

    IfOperStatusReason_tunnelSrcDown IfOperStatusReason = "tunnelSrcDown"

    IfOperStatusReason_cdpInfoUnavailable IfOperStatusReason = "cdpInfoUnavailable"

    IfOperStatusReason_fexSfpInvalid IfOperStatusReason = "fexSfpInvalid"

    IfOperStatusReason_errDisabledIpConflict IfOperStatusReason = "errDisabledIpConflict"

    IfOperStatusReason_fcotClkRateMismatch IfOperStatusReason = "fcotClkRateMismatch"

    IfOperStatusReason_portGuardTrustSecViolation IfOperStatusReason = "portGuardTrustSecViolation"

    IfOperStatusReason_sdpTimeout IfOperStatusReason = "sdpTimeout"

    IfOperStatusReason_satIncompatTopo IfOperStatusReason = "satIncompatTopo"

    IfOperStatusReason_waitForFlogi IfOperStatusReason = "waitForFlogi"

    IfOperStatusReason_satNotConfigured IfOperStatusReason = "satNotConfigured"

    IfOperStatusReason_npivNotEnabledInUpstream IfOperStatusReason = "npivNotEnabledInUpstream"

    IfOperStatusReason_vsanMismatchWithUpstreamSwPort IfOperStatusReason = "vsanMismatchWithUpstreamSwPort"

    IfOperStatusReason_portGuardBitErrRate IfOperStatusReason = "portGuardBitErrRate"

    IfOperStatusReason_portGuardSigLoss IfOperStatusReason = "portGuardSigLoss"

    IfOperStatusReason_portGuardSyncLoss IfOperStatusReason = "portGuardSyncLoss"

    IfOperStatusReason_portGuardLinkReset IfOperStatusReason = "portGuardLinkReset"

    IfOperStatusReason_portGuardCreditLoss IfOperStatusReason = "portGuardCreditLoss"

    IfOperStatusReason_ipQosMgrPolicyAppFailure IfOperStatusReason = "ipQosMgrPolicyAppFailure"

    IfOperStatusReason_pauseRateLimitErrDisabled IfOperStatusReason = "pauseRateLimitErrDisabled"

    IfOperStatusReason_lstGrpUplinkDown IfOperStatusReason = "lstGrpUplinkDown"

    IfOperStatusReason_stickyDnLinkFailure IfOperStatusReason = "stickyDnLinkFailure"

    IfOperStatusReason_routerMacFailure IfOperStatusReason = "routerMacFailure"

    IfOperStatusReason_dcxMultipleMsapIds IfOperStatusReason = "dcxMultipleMsapIds"

    IfOperStatusReason_dcxHundredPdusRcvdNoAck IfOperStatusReason = "dcxHundredPdusRcvdNoAck"

    IfOperStatusReason_enmSatIncompatibleUplink IfOperStatusReason = "enmSatIncompatibleUplink"

    IfOperStatusReason_enmLoopDetected IfOperStatusReason = "enmLoopDetected"

    IfOperStatusReason_nonStickyExternallyDisabled IfOperStatusReason = "nonStickyExternallyDisabled"

    IfOperStatusReason_subGroupIdNotAssigned IfOperStatusReason = "subGroupIdNotAssigned"

    IfOperStatusReason_vemUnlicensed IfOperStatusReason = "vemUnlicensed"

    IfOperStatusReason_profileNotFound IfOperStatusReason = "profileNotFound"

    IfOperStatusReason_nonExistentVlan IfOperStatusReason = "nonExistentVlan"

    IfOperStatusReason_vlanInvalidType IfOperStatusReason = "vlanInvalidType"

    IfOperStatusReason_vlanDown IfOperStatusReason = "vlanDown"

    IfOperStatusReason_vpcPeerUpgrade IfOperStatusReason = "vpcPeerUpgrade"

    IfOperStatusReason_ipQosDcbxpCompatFailure IfOperStatusReason = "ipQosDcbxpCompatFailure"

    IfOperStatusReason_nonCiscoHbaVfTag IfOperStatusReason = "nonCiscoHbaVfTag"

    IfOperStatusReason_domainIdConfigMismatch IfOperStatusReason = "domainIdConfigMismatch"

    IfOperStatusReason_sfpSpeedMismatch IfOperStatusReason = "sfpSpeedMismatch"

    IfOperStatusReason_xcvrInitializing IfOperStatusReason = "xcvrInitializing"

    IfOperStatusReason_xcvrAbsent IfOperStatusReason = "xcvrAbsent"

    IfOperStatusReason_xcvrInvalid IfOperStatusReason = "xcvrInvalid"

    IfOperStatusReason_vfcBindingInvalid IfOperStatusReason = "vfcBindingInvalid"

    IfOperStatusReason_vlanNotFcoeEnabled IfOperStatusReason = "vlanNotFcoeEnabled"

    IfOperStatusReason_pvlanNativeVlanErr IfOperStatusReason = "pvlanNativeVlanErr"

    IfOperStatusReason_ethL2VlanDown IfOperStatusReason = "ethL2VlanDown"

    IfOperStatusReason_ethIntfInvalidBinding IfOperStatusReason = "ethIntfInvalidBinding"

    IfOperStatusReason_pmonFailure IfOperStatusReason = "pmonFailure"

    IfOperStatusReason_l3NotReady IfOperStatusReason = "l3NotReady"

    IfOperStatusReason_speedMismatch IfOperStatusReason = "speedMismatch"

    IfOperStatusReason_flowControlMismatch IfOperStatusReason = "flowControlMismatch"

    IfOperStatusReason_vdcMode IfOperStatusReason = "vdcMode"

    IfOperStatusReason_suspendedDueToMinLinks IfOperStatusReason = "suspendedDueToMinLinks"

    IfOperStatusReason_enmPinFailLinkDown IfOperStatusReason = "enmPinFailLinkDown"

    IfOperStatusReason_inactiveM1PortFpathActiveVlan IfOperStatusReason = "inactiveM1PortFpathActiveVlan"

    IfOperStatusReason_parentPortDown IfOperStatusReason = "parentPortDown"

    IfOperStatusReason_moduleRemoved IfOperStatusReason = "moduleRemoved"

    IfOperStatusReason_corePortMct IfOperStatusReason = "corePortMct"

    IfOperStatusReason_nonCorePortMct IfOperStatusReason = "nonCorePortMct"

    IfOperStatusReason_ficonInorderNotActive IfOperStatusReason = "ficonInorderNotActive"

    IfOperStatusReason_invalidEncapsulation IfOperStatusReason = "invalidEncapsulation"

    IfOperStatusReason_modemLineDeasserted IfOperStatusReason = "modemLineDeasserted"

    IfOperStatusReason_ipSecHndshkInProgress IfOperStatusReason = "ipSecHndshkInProgress"

    IfOperStatusReason_sfpEthCompliantErr IfOperStatusReason = "sfpEthCompliantErr"

    IfOperStatusReason_cellularModemUnattached IfOperStatusReason = "cellularModemUnattached"

    IfOperStatusReason_outOfGlblRxBuffers IfOperStatusReason = "outOfGlblRxBuffers"

    IfOperStatusReason_sfpFcCompliantErr IfOperStatusReason = "sfpFcCompliantErr"

    IfOperStatusReason_ethIntfNotLicensed IfOperStatusReason = "ethIntfNotLicensed"

    IfOperStatusReason_domainIdsInvalid IfOperStatusReason = "domainIdsInvalid"

    IfOperStatusReason_fabricNameInvalid IfOperStatusReason = "fabricNameInvalid"
)

// CiscoAlarmSeverity represents         informational events.
type CiscoAlarmSeverity string

const (
    CiscoAlarmSeverity_cleared CiscoAlarmSeverity = "cleared"

    CiscoAlarmSeverity_indeterminate CiscoAlarmSeverity = "indeterminate"

    CiscoAlarmSeverity_critical CiscoAlarmSeverity = "critical"

    CiscoAlarmSeverity_major CiscoAlarmSeverity = "major"

    CiscoAlarmSeverity_minor CiscoAlarmSeverity = "minor"

    CiscoAlarmSeverity_warning CiscoAlarmSeverity = "warning"

    CiscoAlarmSeverity_info CiscoAlarmSeverity = "info"
)

