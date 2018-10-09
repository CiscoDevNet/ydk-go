// This MIB module contains the management objects for the
// management of DOCSIS 3.0 features, primarily channel bonding,
// interface topology and enhanced signal quality monitoring.
// Copyright 1999-2016 Cable Television Laboratories, Inc.
// All rights reserved.
package docs_if3_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package docs_if3_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:DOCS-IF3-MIB DOCS-IF3-MIB}", reflect.TypeOf(DOCSIF3MIB{}))
    ydk.RegisterEntity("DOCS-IF3-MIB:DOCS-IF3-MIB", reflect.TypeOf(DOCSIF3MIB{}))
}

// CmtsCmRegState represents  in the CM-CTRL-REQ message from CMTS.
type CmtsCmRegState string

const (
    CmtsCmRegState_other CmtsCmRegState = "other"

    CmtsCmRegState_initialRanging CmtsCmRegState = "initialRanging"

    CmtsCmRegState_rangingAutoAdjComplete CmtsCmRegState = "rangingAutoAdjComplete"

    CmtsCmRegState_dhcpv4Complete CmtsCmRegState = "dhcpv4Complete"

    CmtsCmRegState_registrationComplete CmtsCmRegState = "registrationComplete"

    CmtsCmRegState_operational CmtsCmRegState = "operational"

    CmtsCmRegState_bpiInit CmtsCmRegState = "bpiInit"

    CmtsCmRegState_startEae CmtsCmRegState = "startEae"

    CmtsCmRegState_startDhcpv4 CmtsCmRegState = "startDhcpv4"

    CmtsCmRegState_startDhcpv6 CmtsCmRegState = "startDhcpv6"

    CmtsCmRegState_dhcpv6Complete CmtsCmRegState = "dhcpv6Complete"

    CmtsCmRegState_startConfigFileDownload CmtsCmRegState = "startConfigFileDownload"

    CmtsCmRegState_configFileDownloadComplete CmtsCmRegState = "configFileDownloadComplete"

    CmtsCmRegState_startRegistration CmtsCmRegState = "startRegistration"

    CmtsCmRegState_forwardingDisabled CmtsCmRegState = "forwardingDisabled"

    CmtsCmRegState_rfMuteAll CmtsCmRegState = "rfMuteAll"
)

// IfDirection represents Cable Modem Termination System.
type IfDirection string

const (
    IfDirection_downstream IfDirection = "downstream"

    IfDirection_upstream IfDirection = "upstream"
)

// SpectrumAnalysisWindowFunction represents be returned.
type SpectrumAnalysisWindowFunction string

const (
    SpectrumAnalysisWindowFunction_other SpectrumAnalysisWindowFunction = "other"

    SpectrumAnalysisWindowFunction_hann SpectrumAnalysisWindowFunction = "hann"

    SpectrumAnalysisWindowFunction_blackmanHarris SpectrumAnalysisWindowFunction = "blackmanHarris"

    SpectrumAnalysisWindowFunction_rectangular SpectrumAnalysisWindowFunction = "rectangular"

    SpectrumAnalysisWindowFunction_hamming SpectrumAnalysisWindowFunction = "hamming"

    SpectrumAnalysisWindowFunction_flatTop SpectrumAnalysisWindowFunction = "flatTop"

    SpectrumAnalysisWindowFunction_gaussian SpectrumAnalysisWindowFunction = "gaussian"

    SpectrumAnalysisWindowFunction_chebyshev SpectrumAnalysisWindowFunction = "chebyshev"
)

// CmRegState represents  in the CM-CTRL-REQ message from CMTS.
type CmRegState string

const (
    CmRegState_other CmRegState = "other"

    CmRegState_notReady CmRegState = "notReady"

    CmRegState_notSynchronized CmRegState = "notSynchronized"

    CmRegState_phySynchronized CmRegState = "phySynchronized"

    CmRegState_usParametersAcquired CmRegState = "usParametersAcquired"

    CmRegState_rangingComplete CmRegState = "rangingComplete"

    CmRegState_dhcpv4Complete CmRegState = "dhcpv4Complete"

    CmRegState_todEstablished CmRegState = "todEstablished"

    CmRegState_securityEstablished CmRegState = "securityEstablished"

    CmRegState_configFileDownloadComplete CmRegState = "configFileDownloadComplete"

    CmRegState_registrationComplete CmRegState = "registrationComplete"

    CmRegState_operational CmRegState = "operational"

    CmRegState_accessDenied CmRegState = "accessDenied"

    CmRegState_eaeInProgress CmRegState = "eaeInProgress"

    CmRegState_dhcpv4InProgress CmRegState = "dhcpv4InProgress"

    CmRegState_dhcpv6InProgress CmRegState = "dhcpv6InProgress"

    CmRegState_dhcpv6Complete CmRegState = "dhcpv6Complete"

    CmRegState_registrationInProgress CmRegState = "registrationInProgress"

    CmRegState_bpiInit CmRegState = "bpiInit"

    CmRegState_forwardingDisabled CmRegState = "forwardingDisabled"

    CmRegState_dsTopologyResolutionInProgress CmRegState = "dsTopologyResolutionInProgress"

    CmRegState_rangingInProgress CmRegState = "rangingInProgress"

    CmRegState_rfMuteAll CmRegState = "rfMuteAll"
)

// RangingState represents  indicates that the T4 timer expired on the CM.
type RangingState string

const (
    RangingState_other RangingState = "other"

    RangingState_aborted RangingState = "aborted"

    RangingState_retriesExceeded RangingState = "retriesExceeded"

    RangingState_success RangingState = "success"

    RangingState_continue_ RangingState = "continue"

    RangingState_timeoutT4 RangingState = "timeoutT4"
)

// DOCSIF3MIB
type DOCSIF3MIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    DocsIf3CmCapabilities DOCSIF3MIB_DocsIf3CmCapabilities

    
    DocsIf3CmtsCmCtrl DOCSIF3MIB_DocsIf3CmtsCmCtrl

    
    DocsIf3CmEnergyMgtCfg DOCSIF3MIB_DocsIf3CmEnergyMgtCfg

    
    DocsIf3CmSpectrumAnalysisCtrlCmd DOCSIF3MIB_DocsIf3CmSpectrumAnalysisCtrlCmd

    // This object defines attributes of the CM connectivity status. This object
    // provides CM connectivity status information of the CM previously available
    // in the SNMP table docsIfCmStatusTable.
    DocsIf3CmStatusTable DOCSIF3MIB_DocsIf3CmStatusTable

    // This object defines PHY and MAC information about the CM's upstream
    // channels operating in Multiple Transmit Channel (MTC) mode or in a Pre-3.0
    // DOSCIS transmit channel mode. This object provides per-CM Upstream channel
    // information previously available in the SNMP table docsIfCmStatusTable.
    DocsIf3CmStatusUsTable DOCSIF3MIB_DocsIf3CmStatusUsTable

    // This object defines attributes that represent the CM's registration status
    // as tracked by the CMTS. Refer to the individual attribute definitions for 
    // applicability to 3.0 and 3.1 Cable Modems.
    DocsIf3CmtsCmRegStatusTable DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable

    // This object defines status information of the CM currently in use Upstream
    // Logical Channels, as reported by the CMTS.
    DocsIf3CmtsCmUsStatusTable DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable

    // This object configures the association of downstream and upstream channels
    // to a particular MAC Domain (MD) on a CMTS.  The creation of channels and
    // MAC domain object interface instances is vendor-specific. In particular,
    // the assignment of the channel interface index is normally vendor-specific.
    // Therefore, this object is intended only for associating channels to a MAC
    // Domain and assumes that those channels were previously configured. The CMTS
    // may have restrictions on which channels can be configured in the same MAC
    // Domain.  For example, it could require the upstream channels to be from the
    // same line card. This object supports the creation and deletion of multiple
    // instances. Creation of a new instance of this object requires the ChId
    // attribute to be set.
    DocsIf3MdChCfgTable DOCSIF3MIB_DocsIf3MdChCfgTable

    // This object identifies the scope of the Receive Channel Configuration (RCC)
    // and provides a top level container for the Receive Module and Receive
    // Channel objects.  The CMTS selects an instance of this object to assign to
    // a CM when it registers. This object supports the creation and deletion of
    // multiple instances.
    DocsIf3RccCfgTable DOCSIF3MIB_DocsIf3RccCfgTable

    // The RCC Status object provides a read-only view of the
    // statically-configured (from the RccCfg object) and dynamically-created
    // RCCs. The CMTS creates an RCC Status instance for each unique MAC Domain
    // Cable Modem Service Group (MD-CM-SG) to which it signals an RCC to the CM.
    DocsIf3RccStatusTable DOCSIF3MIB_DocsIf3RccStatusTable

    // The Receive Channel Configuration object permits an operator to configure
    // how CMs registered with certain Receive Channel Profiles will configure the
    // Receive Channels within their profile. When a CM registers with an RCP for
    // which all Receive Channel Indices (RcIds) are configured in the Receive
    // Module object and all Receive Channels are configured within this object,
    // the CMTS should use the configuration within these objects to set the
    // Receive Channel Configuration returned to the CM in a REG-RSP message.  A
    // CMTS may require configuration of all pertinent Receive Module and Receive
    // Channel instances in order to register a CM that reports a Receive Channel
    // Profile (RCP), including any standard Receive Channel Profiles. If the CM
    // reports multiple RCPs, and Receive Module and Receive Channel objects have
    // instances for more than one RCP, the particular RCP selected by the CMTS is
    // not specified.  A CMTS is not restricted to assigning Receive Modules based
    // only on the contents of this object. This object supports the creation and
    // deletion of multiple instances. Creation of a new instance of this object
    // requires the ChIfIndex attribute to be set and a valid reference of a
    // RccCfg instance.
    DocsIf3RxChCfgTable DOCSIF3MIB_DocsIf3RxChCfgTable

    // The Receive Channel Status object reports the status of the
    // statically-configured and dynamically-created Receive Channels within an
    // RCC.
    DocsIf3RxChStatusTable DOCSIF3MIB_DocsIf3RxChStatusTable

    // The Receive Module Configuration object permits an operator to configure
    // how CMs with certain Receive Channel Profiles (RCPs) will configure the
    // Receive Modules within their profile upon CM registration.  When a CM
    // registers with an RCP for which all Receive Module Indices (RmIds) are
    // configured in this object and all Receive Channels are configured within
    // the Receive Channel (ReceiveChannel) object, the CMTS should use the
    // configuration within these objects to set the Receive Channel Configuration
    // assigned to the CM in a REG-RSP message.  A CMTS may require configuration
    // of all pertinent Receive Module and Receive Channel instances (i.e., MIB
    // table entries) in order to register a CM that reports a Receive Channel
    // Profile.  If the CM reports multiple RCPs, and Receive Module and Receive
    // Channel objects have instances (i.e., MIB table entries) for more than one
    // RCP reported by the CM, the particular RCP selected by the CMTS is not
    // specified.  A CMTS is not restricted to assigning Receive Modules based
    // only on the contents of this object.  This object supports the creation and
    // deletion of multiple instances. Creation of a new instance of this object
    // requires the reference of a valid RccCfg instance.
    DocsIf3RxModuleCfgTable DOCSIF3MIB_DocsIf3RxModuleCfgTable

    // The Receive Module Status object provides a read-only view of the
    // statically configured and dynamically created Receive Modules within an
    // RCC.
    DocsIf3RxModuleStatusTable DOCSIF3MIB_DocsIf3RxModuleStatusTable

    // This object reports the MD-DS-SG-ID and MD-US-SG-ID associated with a
    // MD-CM-SG-ID within a MAC Domain and the Fiber Nodes reached by the
    // MD-CM-SG.
    DocsIf3MdNodeStatusTable DOCSIF3MIB_DocsIf3MdNodeStatusTable

    // This object returns the list of downstream channel associated with a MAC
    // Domain MD-DS-SG-ID.
    DocsIf3MdDsSgStatusTable DOCSIF3MIB_DocsIf3MdDsSgStatusTable

    // This object returns the list of upstream channels associated with a MAC
    // Domain MD-US-SG-ID.
    DocsIf3MdUsSgStatusTable DOCSIF3MIB_DocsIf3MdUsSgStatusTable

    // This object returns the set of downstream channels that carry UCDs and MAPs
    // for a particular upstream channel in a MAC Domain.
    DocsIf3MdUsToDsChMappingTable DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable

    // This object contains MAC domain level control and configuration attributes.
    DocsIf3MdCfgTable DOCSIF3MIB_DocsIf3MdCfgTable

    // This object defines statically configured Downstream Bonding Groups and
    // Upstream Bonding Groups on the CMTS. This object supports the creation and
    // deletion of multiple instances. Creation of a new instance of this object
    // requires the ChList attribute to be set.
    DocsIf3BondingGrpCfgTable DOCSIF3MIB_DocsIf3BondingGrpCfgTable

    // This object returns administratively-configured and CMTS defined downstream
    // bonding groups.
    DocsIf3DsBondingGrpStatusTable DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable

    // This object returns administratively-configured and CMTS-defined upstream
    // bonding groups.
    DocsIf3UsBondingGrpStatusTable DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable

    // This object defines management extensions for upstream channels, in
    // particular SCDMA parameters.
    DocsIf3UsChExtTable DOCSIF3MIB_DocsIf3UsChExtTable

    // This object defines a set of upstream channels. These channel sets may be
    // associated with channel bonding groups, MD-US-SGs, MD-CM-SGs, or any other
    // channel set that the CMTS may derive from other CMTS processes.
    DocsIf3UsChSetTable DOCSIF3MIB_DocsIf3UsChSetTable

    // This object defines a set of downstream channels. These channel sets may be
    // associated with channel bonding groups, MD-DS-SGs, MD-CM-SGs, or any other
    // channel set that the CMTS may derive from other CMTS processes.
    DocsIf3DsChSetTable DOCSIF3MIB_DocsIf3DsChSetTable

    // This object provides an in-channel received modulation error ratio metric
    // for CM and CMTS.
    DocsIf3SignalQualityExtTable DOCSIF3MIB_DocsIf3SignalQualityExtTable

    // This object provides metrics and parameters associated with received
    // carrier, noise and interference power levels in the upstream channels of
    // the CMTS.
    DocsIf3CmtsSignalQualityExtTable DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable

    // This object is used to configure the logical upstream interfaces to perform
    // the spectrum measurements. This object supports creation and deletion of
    // instances.
    DocsIf3CmtsSpectrumAnalysisMeasTable DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable

    // This object represents the DOCSIS Path Verify Statistics collected in the
    // cable modem device. The CMTS controls the logging of DPV statistics in the
    // cable modem. Therefore the context and nature of the measurements are
    // governed by the CMTS and not self-descriptive when read from the CM.
    DocsIf3CmDpvStatsTable DOCSIF3MIB_DocsIf3CmDpvStatsTable

    // This object represents the control mechanism to enable the dispatching of
    // events based on the event Id. The following rules define the event control
    // behavior:  - The CmEventCtrl object has no instances or contains an   
    // instance with Event ID 0. All events matching the Local Log   settings of
    // docsDevEvReporting are sent to local log ONLY.  - The CmEventCtrl object
    // contains configured instances   Only events matching the Event Ids
    // configured in the object   are sent according to the settings of the
    // docsDevEvReporting   object.       The CM does not persist instances of
    // CmEventCtrl across    reinitializations.
    DocsIf3CmEventCtrlTable DOCSIF3MIB_DocsIf3CmEventCtrlTable

    // This object represents the control mechanism to enable the dispatching of 
    // events based on the event Id. The following rules define the event control
    // behavior:  - The CmtsEventCtrl object has no instances or contains an   
    // instance with Event ID 0. All events matching the Local Log    settings of
    // docsDevEvReporting are sent to local log ONLY.  - The CmtsEventCtrl object
    // contains configured instances   Only events matching the Event Ids
    // configured in the object   are sent according to the settings of the
    // docsDevEvReporting   object.     - The CMTS persists all instances of
    // CmtsEventCtrl across    reinitializations.
    DocsIf3CmtsEventCtrlTable DOCSIF3MIB_DocsIf3CmtsEventCtrlTable

    // This object contains CM MAC domain level control and configuration
    // attributes.
    DocsIf3CmMdCfgTable DOCSIF3MIB_DocsIf3CmMdCfgTable

    // This object provides configuration state information  on the CM for the
    // Energy Management 1x1 Mode feature.
    DocsIf3CmEnergyMgt1x1CfgTable DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable

    // This table provides a list of spectral analysis measurements  as performed
    // across a range of center frequencies. The table  is capable of representing
    // a full scan of the spectrum.
    DocsIf3CmSpectrumAnalysisMeasTable DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable

    // This table defines Energy Management mode statistics for the CM as reported
    // by the CMTS.  For example, such metrics can provide insight into
    // configuration of appropriate EM 1x1 Mode Activity Detection thresholds on
    // the CM and/or to get feedback on how/if the current thresholds are working
    // well or are  causing user experience issues.
    DocsIf3CmtsCmEmStatsTable DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable

    // This object defines Energy Management 1x1 mode statistics on the CM to
    // provide insight into configuration of appropriate EM 1x1 Mode Activity
    // Detection thresholds and/or to get feedback on how/if the current
    // thresholds are working well or are  causing user experience issues. These
    // statistics are only applicable/valid when the Energy  Management 1x1 mode
    // is enabled in the CM.
    DocsIf3CmEm1x1StatsTable DOCSIF3MIB_DocsIf3CmEm1x1StatsTable
}

func (dOCSIF3MIB *DOCSIF3MIB) GetEntityData() *types.CommonEntityData {
    dOCSIF3MIB.EntityData.YFilter = dOCSIF3MIB.YFilter
    dOCSIF3MIB.EntityData.YangName = "DOCS-IF3-MIB"
    dOCSIF3MIB.EntityData.BundleName = "cisco_ios_xe"
    dOCSIF3MIB.EntityData.ParentYangName = "DOCS-IF3-MIB"
    dOCSIF3MIB.EntityData.SegmentPath = "DOCS-IF3-MIB:DOCS-IF3-MIB"
    dOCSIF3MIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dOCSIF3MIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dOCSIF3MIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dOCSIF3MIB.EntityData.Children = types.NewOrderedMap()
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmCapabilities", types.YChild{"DocsIf3CmCapabilities", &dOCSIF3MIB.DocsIf3CmCapabilities})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsCmCtrl", types.YChild{"DocsIf3CmtsCmCtrl", &dOCSIF3MIB.DocsIf3CmtsCmCtrl})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmEnergyMgtCfg", types.YChild{"DocsIf3CmEnergyMgtCfg", &dOCSIF3MIB.DocsIf3CmEnergyMgtCfg})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmSpectrumAnalysisCtrlCmd", types.YChild{"DocsIf3CmSpectrumAnalysisCtrlCmd", &dOCSIF3MIB.DocsIf3CmSpectrumAnalysisCtrlCmd})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmStatusTable", types.YChild{"DocsIf3CmStatusTable", &dOCSIF3MIB.DocsIf3CmStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmStatusUsTable", types.YChild{"DocsIf3CmStatusUsTable", &dOCSIF3MIB.DocsIf3CmStatusUsTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsCmRegStatusTable", types.YChild{"DocsIf3CmtsCmRegStatusTable", &dOCSIF3MIB.DocsIf3CmtsCmRegStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsCmUsStatusTable", types.YChild{"DocsIf3CmtsCmUsStatusTable", &dOCSIF3MIB.DocsIf3CmtsCmUsStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdChCfgTable", types.YChild{"DocsIf3MdChCfgTable", &dOCSIF3MIB.DocsIf3MdChCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RccCfgTable", types.YChild{"DocsIf3RccCfgTable", &dOCSIF3MIB.DocsIf3RccCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RccStatusTable", types.YChild{"DocsIf3RccStatusTable", &dOCSIF3MIB.DocsIf3RccStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RxChCfgTable", types.YChild{"DocsIf3RxChCfgTable", &dOCSIF3MIB.DocsIf3RxChCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RxChStatusTable", types.YChild{"DocsIf3RxChStatusTable", &dOCSIF3MIB.DocsIf3RxChStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RxModuleCfgTable", types.YChild{"DocsIf3RxModuleCfgTable", &dOCSIF3MIB.DocsIf3RxModuleCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3RxModuleStatusTable", types.YChild{"DocsIf3RxModuleStatusTable", &dOCSIF3MIB.DocsIf3RxModuleStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdNodeStatusTable", types.YChild{"DocsIf3MdNodeStatusTable", &dOCSIF3MIB.DocsIf3MdNodeStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdDsSgStatusTable", types.YChild{"DocsIf3MdDsSgStatusTable", &dOCSIF3MIB.DocsIf3MdDsSgStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdUsSgStatusTable", types.YChild{"DocsIf3MdUsSgStatusTable", &dOCSIF3MIB.DocsIf3MdUsSgStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdUsToDsChMappingTable", types.YChild{"DocsIf3MdUsToDsChMappingTable", &dOCSIF3MIB.DocsIf3MdUsToDsChMappingTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3MdCfgTable", types.YChild{"DocsIf3MdCfgTable", &dOCSIF3MIB.DocsIf3MdCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3BondingGrpCfgTable", types.YChild{"DocsIf3BondingGrpCfgTable", &dOCSIF3MIB.DocsIf3BondingGrpCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3DsBondingGrpStatusTable", types.YChild{"DocsIf3DsBondingGrpStatusTable", &dOCSIF3MIB.DocsIf3DsBondingGrpStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3UsBondingGrpStatusTable", types.YChild{"DocsIf3UsBondingGrpStatusTable", &dOCSIF3MIB.DocsIf3UsBondingGrpStatusTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3UsChExtTable", types.YChild{"DocsIf3UsChExtTable", &dOCSIF3MIB.DocsIf3UsChExtTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3UsChSetTable", types.YChild{"DocsIf3UsChSetTable", &dOCSIF3MIB.DocsIf3UsChSetTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3DsChSetTable", types.YChild{"DocsIf3DsChSetTable", &dOCSIF3MIB.DocsIf3DsChSetTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3SignalQualityExtTable", types.YChild{"DocsIf3SignalQualityExtTable", &dOCSIF3MIB.DocsIf3SignalQualityExtTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsSignalQualityExtTable", types.YChild{"DocsIf3CmtsSignalQualityExtTable", &dOCSIF3MIB.DocsIf3CmtsSignalQualityExtTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsSpectrumAnalysisMeasTable", types.YChild{"DocsIf3CmtsSpectrumAnalysisMeasTable", &dOCSIF3MIB.DocsIf3CmtsSpectrumAnalysisMeasTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmDpvStatsTable", types.YChild{"DocsIf3CmDpvStatsTable", &dOCSIF3MIB.DocsIf3CmDpvStatsTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmEventCtrlTable", types.YChild{"DocsIf3CmEventCtrlTable", &dOCSIF3MIB.DocsIf3CmEventCtrlTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsEventCtrlTable", types.YChild{"DocsIf3CmtsEventCtrlTable", &dOCSIF3MIB.DocsIf3CmtsEventCtrlTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmMdCfgTable", types.YChild{"DocsIf3CmMdCfgTable", &dOCSIF3MIB.DocsIf3CmMdCfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmEnergyMgt1x1CfgTable", types.YChild{"DocsIf3CmEnergyMgt1x1CfgTable", &dOCSIF3MIB.DocsIf3CmEnergyMgt1x1CfgTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmSpectrumAnalysisMeasTable", types.YChild{"DocsIf3CmSpectrumAnalysisMeasTable", &dOCSIF3MIB.DocsIf3CmSpectrumAnalysisMeasTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmtsCmEmStatsTable", types.YChild{"DocsIf3CmtsCmEmStatsTable", &dOCSIF3MIB.DocsIf3CmtsCmEmStatsTable})
    dOCSIF3MIB.EntityData.Children.Append("docsIf3CmEm1x1StatsTable", types.YChild{"DocsIf3CmEm1x1StatsTable", &dOCSIF3MIB.DocsIf3CmEm1x1StatsTable})
    dOCSIF3MIB.EntityData.Leafs = types.NewOrderedMap()

    dOCSIF3MIB.EntityData.YListKeys = []string {}

    return &(dOCSIF3MIB.EntityData)
}

// DOCSIF3MIB_DocsIf3CmCapabilities
type DOCSIF3MIB_DocsIf3CmCapabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute contains the TLV encoding for TLV-5 sent in a REG-REQ.  The
    // first byte of this encoding is expected to be '05'H. The type is string
    // with length: 0 | 2..255.
    DocsIf3CmCapabilitiesReq interface{}

    // This attribute contains the TLV encoding for TLV-5 received in a REG-RSP.
    // The first byte of this encoding is expected to be '05'H. The type is string
    // with length: 0 | 2..255.
    DocsIf3CmCapabilitiesRsp interface{}
}

func (docsIf3CmCapabilities *DOCSIF3MIB_DocsIf3CmCapabilities) GetEntityData() *types.CommonEntityData {
    docsIf3CmCapabilities.EntityData.YFilter = docsIf3CmCapabilities.YFilter
    docsIf3CmCapabilities.EntityData.YangName = "docsIf3CmCapabilities"
    docsIf3CmCapabilities.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmCapabilities.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmCapabilities.EntityData.SegmentPath = "docsIf3CmCapabilities"
    docsIf3CmCapabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmCapabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmCapabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmCapabilities.EntityData.Children = types.NewOrderedMap()
    docsIf3CmCapabilities.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmCapabilities.EntityData.Leafs.Append("docsIf3CmCapabilitiesReq", types.YLeaf{"DocsIf3CmCapabilitiesReq", docsIf3CmCapabilities.DocsIf3CmCapabilitiesReq})
    docsIf3CmCapabilities.EntityData.Leafs.Append("docsIf3CmCapabilitiesRsp", types.YLeaf{"DocsIf3CmCapabilitiesRsp", docsIf3CmCapabilities.DocsIf3CmCapabilitiesRsp})

    docsIf3CmCapabilities.EntityData.YListKeys = []string {}

    return &(docsIf3CmCapabilities.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmCtrl
type DOCSIF3MIB_DocsIf3CmtsCmCtrl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute represents the MAC Address of the CM which the  CMTS is
    // instructed to send the CM-CTRL-REQ message. The type is string with
    // pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIf3CmtsCmCtrlCmdMacAddr interface{}

    // This attribute represents the Upstream Channel ID (UCID) to  mute or
    // unmute.  A value of zero indicates all upstream  channels.  This attribute
    // is only applicable when the  docsIf3CmtsCmCtrlCmdCommit attribute is set to
    // 'mute'. The type is interface{} with range: 0..255.
    DocsIf3CmtsCmCtrlCmdMuteUsChId interface{}

    // This attribute represents the length of time that the mute  operation is in
    // effect.  This attribute is only applicable  when the
    // docsIf3CmtsCmCtrlCmdCommit attribute is set to  'mute'.  A value of 0 is an
    // indication to unmute the channel referenced by the
    // docsIf3CmtsCmCtrlCmdMuteUsChId attribute while a value of 0xFFFFFFFF is
    // used to mute the channel referenced by the docsIf3CmtsCmCtrlCmdMuteUsChId
    // attribute indefinitely. The type is interface{} with range: 0..4294967295.
    // Units are milliseconds.
    DocsIf3CmtsCmCtrlCmdMuteInterval interface{}

    // When set to 'true', this attribute disables data forwarding  to the CMCI
    // when the docsIf3CmtsCmCtrlCmdCommit attribute is set to
    // 'disableForwarding'.  When set to 'false', this attribute enables data
    // forwarding to the CMCI when the docsIf3CmtsCmCtrlCmdCommit attribute is set
    // to 'disableForwarding'.  This attribute is only applicable when the 
    // docsIf3CmtsCmCtrlCmdCommit attribute is set to  'disableForwarding'. The
    // type is bool.
    DocsIf3CmtsCmCtrlCmdDisableForwarding interface{}

    // This attribute indicates the type of command for the  CMTS to trigger in
    // the CM-CTRL-REQ message. This attribute will return the value of the last
    // operation  performed or the default value if no operation has been 
    // performed. The type is DocsIf3CmtsCmCtrlCmdCommit.
    DocsIf3CmtsCmCtrlCmdCommit interface{}
}

func (docsIf3CmtsCmCtrl *DOCSIF3MIB_DocsIf3CmtsCmCtrl) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmCtrl.EntityData.YFilter = docsIf3CmtsCmCtrl.YFilter
    docsIf3CmtsCmCtrl.EntityData.YangName = "docsIf3CmtsCmCtrl"
    docsIf3CmtsCmCtrl.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmCtrl.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsCmCtrl.EntityData.SegmentPath = "docsIf3CmtsCmCtrl"
    docsIf3CmtsCmCtrl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmCtrl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmCtrl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmCtrl.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmCtrl.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsCmCtrl.EntityData.Leafs.Append("docsIf3CmtsCmCtrlCmdMacAddr", types.YLeaf{"DocsIf3CmtsCmCtrlCmdMacAddr", docsIf3CmtsCmCtrl.DocsIf3CmtsCmCtrlCmdMacAddr})
    docsIf3CmtsCmCtrl.EntityData.Leafs.Append("docsIf3CmtsCmCtrlCmdMuteUsChId", types.YLeaf{"DocsIf3CmtsCmCtrlCmdMuteUsChId", docsIf3CmtsCmCtrl.DocsIf3CmtsCmCtrlCmdMuteUsChId})
    docsIf3CmtsCmCtrl.EntityData.Leafs.Append("docsIf3CmtsCmCtrlCmdMuteInterval", types.YLeaf{"DocsIf3CmtsCmCtrlCmdMuteInterval", docsIf3CmtsCmCtrl.DocsIf3CmtsCmCtrlCmdMuteInterval})
    docsIf3CmtsCmCtrl.EntityData.Leafs.Append("docsIf3CmtsCmCtrlCmdDisableForwarding", types.YLeaf{"DocsIf3CmtsCmCtrlCmdDisableForwarding", docsIf3CmtsCmCtrl.DocsIf3CmtsCmCtrlCmdDisableForwarding})
    docsIf3CmtsCmCtrl.EntityData.Leafs.Append("docsIf3CmtsCmCtrlCmdCommit", types.YLeaf{"DocsIf3CmtsCmCtrlCmdCommit", docsIf3CmtsCmCtrl.DocsIf3CmtsCmCtrlCmdCommit})

    docsIf3CmtsCmCtrl.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsCmCtrl.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit represents performed.
type DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit string

const (
    DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit_mute DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit = "mute"

    DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit_cmReinit DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit = "cmReinit"

    DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit_disableForwarding DOCSIF3MIB_DocsIf3CmtsCmCtrl_DocsIf3CmtsCmCtrlCmdCommit = "disableForwarding"
)

// DOCSIF3MIB_DocsIf3CmEnergyMgtCfg
type DOCSIF3MIB_DocsIf3CmEnergyMgtCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute indicates which energy savings features have been enabled in
    // the Cable Modem. The CM enables use of Energy Management Features only if
    // both the Energy Management Feature Control TLV and Energy Management Modem
    // Capability Response from the CMTS indicate that the feature is enabled.  If
    // bit 0 is set,  the Energy Management 1x1 Mode feature is enabled. The type
    // is map[string]bool.
    DocsIf3CmEnergyMgtCfgFeatureEnabled interface{}

    // This attribute specifies a minimum time period (in seconds)  that must
    // elapse between EM-REQ transactions in certain  situations:  - In the case
    // of Energy Management 1x1 Mode, this attribute   sets the minimum cycle time
    // that a CM will use for sending   requests to enter Energy Management 1x1
    // Mode.  - In the case that the CM fails to receive an EM-RSP message   after
    // the maximum number of retries, this attribute sets    the minimum amount of
    // time to elapse before the CM can    attempt another EM-REQ transaction. The
    // type is interface{} with range: 0..65535. Units are seconds.
    DocsIf3CmEnergyMgtCfgCyclePeriod interface{}
}

func (docsIf3CmEnergyMgtCfg *DOCSIF3MIB_DocsIf3CmEnergyMgtCfg) GetEntityData() *types.CommonEntityData {
    docsIf3CmEnergyMgtCfg.EntityData.YFilter = docsIf3CmEnergyMgtCfg.YFilter
    docsIf3CmEnergyMgtCfg.EntityData.YangName = "docsIf3CmEnergyMgtCfg"
    docsIf3CmEnergyMgtCfg.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEnergyMgtCfg.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmEnergyMgtCfg.EntityData.SegmentPath = "docsIf3CmEnergyMgtCfg"
    docsIf3CmEnergyMgtCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEnergyMgtCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEnergyMgtCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEnergyMgtCfg.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEnergyMgtCfg.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmEnergyMgtCfg.EntityData.Leafs.Append("docsIf3CmEnergyMgtCfgFeatureEnabled", types.YLeaf{"DocsIf3CmEnergyMgtCfgFeatureEnabled", docsIf3CmEnergyMgtCfg.DocsIf3CmEnergyMgtCfgFeatureEnabled})
    docsIf3CmEnergyMgtCfg.EntityData.Leafs.Append("docsIf3CmEnergyMgtCfgCyclePeriod", types.YLeaf{"DocsIf3CmEnergyMgtCfgCyclePeriod", docsIf3CmEnergyMgtCfg.DocsIf3CmEnergyMgtCfgCyclePeriod})

    docsIf3CmEnergyMgtCfg.EntityData.YListKeys = []string {}

    return &(docsIf3CmEnergyMgtCfg.EntityData)
}

// DOCSIF3MIB_DocsIf3CmSpectrumAnalysisCtrlCmd
type DOCSIF3MIB_DocsIf3CmSpectrumAnalysisCtrlCmd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is used to enable or disable the spectrum  analyzer feature.
    // Setting this attribute to true triggers the CM to initiate measurements for
    // the spectrum analyzer feature based on the other configuration attributes
    // for the feature. By  default, the feature is disabled unless explicitly
    // enabled.  Note that the feature may be disabled by the system under 
    // certain circumstances if the spectrum analyzer would affect  critical
    // services. In such a case, the attribute will return 'false' when read, and
    // will reject sets to 'true' with an  error. Once the feature is enabled, any
    // Set operation on  the docsIf3CmSpectrumAnalysisCtrlCmd MIB group might not 
    // be effective until the feature is re-enabled again. The type is bool.
    DocsIf3CmSpectrumAnalysisCtrlCmdEnable interface{}

    // This attribute controls the length of time after the last  spectrum
    // analysis measurement before the feature is  automatically disabled. If set
    // to a value of 0, the feature will remain enabled until it is explicitly
    // disabled. The type is interface{} with range: 0..86400. Units are seconds.
    DocsIf3CmSpectrumAnalysisCtrlCmdInactivityTimeout interface{}

    // This attribute controls the center frequency of the first segment for the
    // spectrum analysis measurement. The frequency bins for this segment lie
    // symmetrically to the left and right of this center frequency.    If the
    // number of bins in a segment is odd, the segment center frequency lies
    // directly on the center bin.    If the number of bins in a segment is even,
    // the segment center frequency lies halfway between two bins.  Changing the
    // value of this attribute may result in  changes to the
    // docsIf3CmSpectrumAnalysisMeasTable. Note that if this attribute is set to
    // an invalid value,  the device may return an error, or may adjust the  value
    // of the attribute to the closest valid value. The type is interface{} with
    // range: 0..4294967295. Units are Hz.
    DocsIf3CmSpectrumAnalysisCtrlCmdFirstSegmentCenterFrequency interface{}

    // This attribute controls the center frequency of the last segment of the
    // spectrum analysis measurement.  The frequency bins for this segment lie
    // symmetrically to the left and right of this center frequency.   If the
    // number of bins in a segment is odd, the segment center frequency lies
    // directly on the center bin.   If the number of bins in a segment is even,
    // the segment center frequency lies halfway between two bins.  The value of
    // the LastSegmentCenterFrequency attribute  should be equal to the
    // FirstSegmentCenterFrequency  plus and integer number of segment spans as
    // determined  by the SegmentFrequencySpan.  Changing the value of this
    // attribute may result in changes to the docsIf3CmSpectrumAnalysisMeasTable. 
    // Note that if this attribute is set to an invalid value,  the device may
    // return an error, or may adjust the  value of the attribute to the closest
    // valid value. The type is interface{} with range: 0..4294967295. Units are
    // Hz.
    DocsIf3CmSpectrumAnalysisCtrlCmdLastSegmentCenterFrequency interface{}

    // This attribute controls the frequency span of each  segment (instance) in
    // the docsIf3CmSpectrumAnalysisMeasTable.    If set to a value of 0, then a
    // default span will be chosen based on the hardware capabilities of the 
    // device. Segments are contiguous from the  FirstSegementCenterFrequency to
    // the  LastSegmentCenterFrequency and the center frequency  for each
    // successive segment is incremented by the  SegmentFequencySpan.  The number
    // of segments is (LastSegmentCenterFrequency -
    // FirstSegmentCenterFrequency)/SegmentFrequencySpan + 1.  A segment is
    // equivalent to an instance in the  docsIf3CmSpectrumAnalysisMeasTable. The
    // chosen  SegmentFrequencySpan affects the number of entries in  the
    // docsIf3CmSpectrumAnalysisMeasTable. A more granular  SegmentFrequencySpan
    // may adversely affect the amount of  time needed to query the table entries
    // in addition to  possibly increasing the acquisition time.  Changing the
    // value of this attribute may result in  changes to the
    // docsIf3CmSpectrumAnalysisMeasTable.  Note that if this attribute is set to
    // an invalid value, the device may return an error, or may adjust the value 
    // of the object to the closest valid value. The type is interface{} with
    // range: 1000000..900000000. Units are Hz.
    DocsIf3CmSpectrumAnalysisCtrlCmdSegmentFrequencySpan interface{}

    // This attribute controls the number of bins collected  by the measurement
    // performed for each segment (instance)  in the
    // docsIf3CmSpectrumAnalysisMeasTable.    Note that if this attribute is set
    // to an invalid value,  the device may return an error, or may adjust the
    // value of the attribute to the closest valid value. The type is interface{}
    // with range: 2..2048. Units are bins-per-segment.
    DocsIf3CmSpectrumAnalysisCtrlCmdNumBinsPerSegment interface{}

    // This attribute allows the user to request an equivalent  noise bandwidth
    // for the resolution bandwidth filter used  in the spectrum analysis.  This
    // corresponds to the  spectral width of the window function used when
    // performing a discrete Fourier transform for the analysis.    The window
    // function which corresponds to a value written  to this object may be
    // obtained by reading the value of the WindowFunction attribute.  If an
    // unsupported value is requested, the device may  return an error, or choose
    // the closest valid value to the  one which is requested.   If the closest
    // value is chosen, then a subsequent read of  this attribute will return the
    // actual value which is in use. The type is interface{} with range: 50..500.
    // Units are hundredthsbin.
    DocsIf3CmSpectrumAnalysisCtrlCmdEquivalentNoiseBandwidth interface{}

    // This attribute controls or indicates the windowing function  which will be
    // used when performing the discrete Fourier transform for the analysis.  The
    // WindowFunction and the Equivalent Noise Bandwidth are related. If a
    // particular WindowFunction is selected, then the EquivalentNoiseBandwidth 
    // for the function which is in use, will be reported by the
    // EquivalentNoiseBandwidth attribute. Alternatively if an
    // EquivalentNoiseBandwidth value is chosen then if a WindowFunction function
    // representing that  EquivalentNoiseBandwidth is defined in the CM, that
    // value will be reported in the WindowFunction MIB object, or a value of 
    // 'other' will be reported. Use of 'modern' windowing functions not yet
    // defined will likely be reported as 'other'.  Note that all window functions
    // may not be supported by all  devices.  If an attempt is made to set the
    // object to an  unsupported window function, or if writing of the 
    // WindowFunction is not supported, an error will be returned. The type is
    // SpectrumAnalysisWindowFunction.
    DocsIf3CmSpectrumAnalysisCtrlCmdWindowFunction interface{}

    // This attribute controls the number of averages  that will be performed on
    // spectral bins.   The average will be computed using the 'leaky integrator'
    // method, where:   reported bin value = alpha*accumulated bin values +       
    // (1-alpha)*current bin value.    Alpha is one minus the reciprocal of the
    // number of averages. For example, if N=25, then alpha = 0.96.   A value of 1
    // indicates no averaging.   Re-writing the number of averages will restart
    // the averaging process.  If there are no accumulated values, the 
    // accumulators are made equal to the first measured bin  amplitudes.  If an
    // attempt is made to set the attribute to an unsupported  number of averages,
    // an error will be returned. The type is interface{} with range: 1..1000.
    DocsIf3CmSpectrumAnalysisCtrlCmdNumberOfAverages interface{}
}

func (docsIf3CmSpectrumAnalysisCtrlCmd *DOCSIF3MIB_DocsIf3CmSpectrumAnalysisCtrlCmd) GetEntityData() *types.CommonEntityData {
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.YFilter = docsIf3CmSpectrumAnalysisCtrlCmd.YFilter
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.YangName = "docsIf3CmSpectrumAnalysisCtrlCmd"
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.SegmentPath = "docsIf3CmSpectrumAnalysisCtrlCmd"
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Children = types.NewOrderedMap()
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdEnable", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdEnable", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdEnable})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdInactivityTimeout", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdInactivityTimeout", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdInactivityTimeout})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdFirstSegmentCenterFrequency", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdFirstSegmentCenterFrequency", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdFirstSegmentCenterFrequency})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdLastSegmentCenterFrequency", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdLastSegmentCenterFrequency", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdLastSegmentCenterFrequency})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdSegmentFrequencySpan", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdSegmentFrequencySpan", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdSegmentFrequencySpan})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdNumBinsPerSegment", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdNumBinsPerSegment", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdNumBinsPerSegment})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdEquivalentNoiseBandwidth", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdEquivalentNoiseBandwidth", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdEquivalentNoiseBandwidth})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdWindowFunction", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdWindowFunction", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdWindowFunction})
    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisCtrlCmdNumberOfAverages", types.YLeaf{"DocsIf3CmSpectrumAnalysisCtrlCmdNumberOfAverages", docsIf3CmSpectrumAnalysisCtrlCmd.DocsIf3CmSpectrumAnalysisCtrlCmdNumberOfAverages})

    docsIf3CmSpectrumAnalysisCtrlCmd.EntityData.YListKeys = []string {}

    return &(docsIf3CmSpectrumAnalysisCtrlCmd.EntityData)
}

// DOCSIF3MIB_DocsIf3CmStatusTable
// This object defines attributes of the CM connectivity
// status. This object provides CM connectivity status
// information of the CM previously available in
// the SNMP table docsIfCmStatusTable.
type DOCSIF3MIB_DocsIf3CmStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmStatusTable. An instance exist for the CM
    // MAC Domain Interface. The type is slice of
    // DOCSIF3MIB_DocsIf3CmStatusTable_DocsIf3CmStatusEntry.
    DocsIf3CmStatusEntry []*DOCSIF3MIB_DocsIf3CmStatusTable_DocsIf3CmStatusEntry
}

func (docsIf3CmStatusTable *DOCSIF3MIB_DocsIf3CmStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmStatusTable.EntityData.YFilter = docsIf3CmStatusTable.YFilter
    docsIf3CmStatusTable.EntityData.YangName = "docsIf3CmStatusTable"
    docsIf3CmStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmStatusTable.EntityData.SegmentPath = "docsIf3CmStatusTable"
    docsIf3CmStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmStatusTable.EntityData.Children.Append("docsIf3CmStatusEntry", types.YChild{"DocsIf3CmStatusEntry", nil})
    for i := range docsIf3CmStatusTable.DocsIf3CmStatusEntry {
        docsIf3CmStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmStatusTable.DocsIf3CmStatusEntry[i]), types.YChild{"DocsIf3CmStatusEntry", docsIf3CmStatusTable.DocsIf3CmStatusEntry[i]})
    }
    docsIf3CmStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmStatusTable_DocsIf3CmStatusEntry
// The conceptual row of docsIf3CmStatusTable.
// An instance exist for the CM MAC Domain Interface.
type DOCSIF3MIB_DocsIf3CmStatusTable_DocsIf3CmStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute denotes the current CM connectivity state. For the case of
    // IP acquisition related states, this attribute reflects states for the
    // current CM provisioning mode, not the other DHCP process associated with
    // dual stack operation. The type is CmRegState.
    DocsIf3CmStatusValue interface{}

    // This attribute denotes the status code for CM as defined in the OSSI
    // Specification. The status code consists of a single character indicating
    // error groups, followed by a two- or three-digit number indicating the
    // status condition, followed by a decimal. An example of a returned value
    // could be 'T101.0'. The zero-length hex string indicates no status code yet
    // registered. The type is string with length: 0 | 5..7.
    DocsIf3CmStatusCode interface{}

    // This attribute denotes the number of times the CM reset or initialized this
    // interface. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are resets.
    DocsIf3CmStatusResets interface{}

    // This attribute denotes the number of times the CM lost synchronization with
    // the downstream channel. Discontinuities in the value of this counter can
    // occur at re-initialization of the managed system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime for the CM MAC Domain
    // interface. The type is interface{} with range: 0..4294967295. Units are
    // messages.
    DocsIf3CmStatusLostSyncs interface{}

    // This attribute denotes the number of times the CM received invalid MAP
    // messages. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are maps.
    DocsIf3CmStatusInvalidMaps interface{}

    // This attribute denotes the number of times the CM received invalid UCD
    // messages. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are messages.
    DocsIf3CmStatusInvalidUcds interface{}

    // This attribute denotes the number of times the CM received invalid ranging
    // response messages. Discontinuities in the value of this counter can occur
    // at re-initialization of the managed system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are messages.
    DocsIf3CmStatusInvalidRangingRsps interface{}

    // This attribute denotes the number of times the CM received invalid
    // registration response messages. Discontinuities in the value of this
    // counter can occur at re-initialization of the managed system, and at other
    // times as indicated by the value of ifCounterDiscontinuityTime for the CM
    // MAC Domain interface. The type is interface{} with range: 0..4294967295.
    // Units are messages.
    DocsIf3CmStatusInvalidRegRsps interface{}

    // This attribute denotes the number of times counter T1 expired in the CM.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the managed system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime for the CM MAC Domain interface. The type is
    // interface{} with range: 0..4294967295. Units are timeouts.
    DocsIf3CmStatusT1Timeouts interface{}

    // This attribute denotes the number of times counter T2 expired in the CM.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the managed system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime for the CM MAC Domain interface. The type is
    // interface{} with range: 0..4294967295. Units are timeouts.
    DocsIf3CmStatusT2Timeouts interface{}

    // This attribute denotes the number of successful Upstream Channel Change
    // transactions. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime  for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are attempts.
    DocsIf3CmStatusUCCsSuccesses interface{}

    // This attribute denotes the number of failed Upstream Channel Change
    // transactions. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime for the CM MAC Domain interface.
    // The type is interface{} with range: 0..4294967295. Units are attempts.
    DocsIf3CmStatusUCCFails interface{}

    // This attribute indicates whether the CM is currently operating in Energy
    // Management 1x1 Mode. If this attribute returns 'true',  the CM is operating
    // in Energy Management 1x1 Mode. The type is bool.
    DocsIf3CmStatusEnergyMgt1x1OperStatus interface{}
}

func (docsIf3CmStatusEntry *DOCSIF3MIB_DocsIf3CmStatusTable_DocsIf3CmStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmStatusEntry.EntityData.YFilter = docsIf3CmStatusEntry.YFilter
    docsIf3CmStatusEntry.EntityData.YangName = "docsIf3CmStatusEntry"
    docsIf3CmStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmStatusEntry.EntityData.ParentYangName = "docsIf3CmStatusTable"
    docsIf3CmStatusEntry.EntityData.SegmentPath = "docsIf3CmStatusEntry" + types.AddKeyToken(docsIf3CmStatusEntry.IfIndex, "ifIndex")
    docsIf3CmStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmStatusEntry.IfIndex})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusValue", types.YLeaf{"DocsIf3CmStatusValue", docsIf3CmStatusEntry.DocsIf3CmStatusValue})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusCode", types.YLeaf{"DocsIf3CmStatusCode", docsIf3CmStatusEntry.DocsIf3CmStatusCode})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusResets", types.YLeaf{"DocsIf3CmStatusResets", docsIf3CmStatusEntry.DocsIf3CmStatusResets})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusLostSyncs", types.YLeaf{"DocsIf3CmStatusLostSyncs", docsIf3CmStatusEntry.DocsIf3CmStatusLostSyncs})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusInvalidMaps", types.YLeaf{"DocsIf3CmStatusInvalidMaps", docsIf3CmStatusEntry.DocsIf3CmStatusInvalidMaps})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusInvalidUcds", types.YLeaf{"DocsIf3CmStatusInvalidUcds", docsIf3CmStatusEntry.DocsIf3CmStatusInvalidUcds})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusInvalidRangingRsps", types.YLeaf{"DocsIf3CmStatusInvalidRangingRsps", docsIf3CmStatusEntry.DocsIf3CmStatusInvalidRangingRsps})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusInvalidRegRsps", types.YLeaf{"DocsIf3CmStatusInvalidRegRsps", docsIf3CmStatusEntry.DocsIf3CmStatusInvalidRegRsps})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusT1Timeouts", types.YLeaf{"DocsIf3CmStatusT1Timeouts", docsIf3CmStatusEntry.DocsIf3CmStatusT1Timeouts})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusT2Timeouts", types.YLeaf{"DocsIf3CmStatusT2Timeouts", docsIf3CmStatusEntry.DocsIf3CmStatusT2Timeouts})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusUCCsSuccesses", types.YLeaf{"DocsIf3CmStatusUCCsSuccesses", docsIf3CmStatusEntry.DocsIf3CmStatusUCCsSuccesses})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusUCCFails", types.YLeaf{"DocsIf3CmStatusUCCFails", docsIf3CmStatusEntry.DocsIf3CmStatusUCCFails})
    docsIf3CmStatusEntry.EntityData.Leafs.Append("docsIf3CmStatusEnergyMgt1x1OperStatus", types.YLeaf{"DocsIf3CmStatusEnergyMgt1x1OperStatus", docsIf3CmStatusEntry.DocsIf3CmStatusEnergyMgt1x1OperStatus})

    docsIf3CmStatusEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmStatusUsTable
// This object defines PHY and MAC information about
// the CM's upstream channels operating in Multiple Transmit
// Channel (MTC) mode or in a Pre-3.0 DOSCIS transmit
// channel mode. This object provides per-CM Upstream
// channel information previously available in the
// SNMP table docsIfCmStatusTable.
type DOCSIF3MIB_DocsIf3CmStatusUsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmStatusUsTable. An instance exist for each of
    // the CM's SC-QAM upstream channels  which are configured for data
    // transmission. The type is slice of
    // DOCSIF3MIB_DocsIf3CmStatusUsTable_DocsIf3CmStatusUsEntry.
    DocsIf3CmStatusUsEntry []*DOCSIF3MIB_DocsIf3CmStatusUsTable_DocsIf3CmStatusUsEntry
}

func (docsIf3CmStatusUsTable *DOCSIF3MIB_DocsIf3CmStatusUsTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmStatusUsTable.EntityData.YFilter = docsIf3CmStatusUsTable.YFilter
    docsIf3CmStatusUsTable.EntityData.YangName = "docsIf3CmStatusUsTable"
    docsIf3CmStatusUsTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmStatusUsTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmStatusUsTable.EntityData.SegmentPath = "docsIf3CmStatusUsTable"
    docsIf3CmStatusUsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmStatusUsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmStatusUsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmStatusUsTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmStatusUsTable.EntityData.Children.Append("docsIf3CmStatusUsEntry", types.YChild{"DocsIf3CmStatusUsEntry", nil})
    for i := range docsIf3CmStatusUsTable.DocsIf3CmStatusUsEntry {
        docsIf3CmStatusUsTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmStatusUsTable.DocsIf3CmStatusUsEntry[i]), types.YChild{"DocsIf3CmStatusUsEntry", docsIf3CmStatusUsTable.DocsIf3CmStatusUsEntry[i]})
    }
    docsIf3CmStatusUsTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmStatusUsTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmStatusUsTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmStatusUsTable_DocsIf3CmStatusUsEntry
// The conceptual row of docsIf3CmStatusUsTable.
// An instance exist for each of the CM's SC-QAM upstream channels 
// which are configured for data transmission.
type DOCSIF3MIB_DocsIf3CmStatusUsTable_DocsIf3CmStatusUsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute represents the operational CM transmit power for this SC-QAM
    // upstream channel. In order for this attribute to provide consistent
    // information  under all circumstances, a 3.1 CM will report the average
    // total  power for the SC-QAM channel the same as was done for  DOCSIS 3.0,
    // regardless of whether it is operating with a 3.1 or  a 3.0 CMTS. The value
    // that is reported was referred to as Pr in  the DOCSIS 3.0 PHY Spec. This
    // attribute is not applicable for  OFDMA channels. The type is interface{}
    // with range: -2147483648..2147483647. Units are TenthdBmV.
    DocsIf3CmStatusUsTxPower interface{}

    // This attribute denotes the number of times counter T3 expired in the CM for
    // this upstream channel. Discontinuities in the value of this counter can
    // occur at re-initialization of the managed system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime for the associated
    // upstream channel. The type is interface{} with range: 0..4294967295. Units
    // are timeouts.
    DocsIf3CmStatusUsT3Timeouts interface{}

    // This attribute denotes the number of times counter T4 expired in the CM for
    // this upstream channel. Discontinuities in the value of this counter can
    // occur at re-initialization of the managed system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime for the associated
    // upstream channel. The type is interface{} with range: 0..4294967295. Units
    // are timeouts.
    DocsIf3CmStatusUsT4Timeouts interface{}

    // This attribute denotes the number of times the ranging process was aborted
    // by the CMTS. Discontinuities in the value of this counter can occur at
    // re-initialization of the managed system, and at other times as indicated by
    // the value of ifCounterDiscontinuityTime ([RFC2863]) for the associated
    // upstream channel. The type is interface{} with range: 0..4294967295. Units
    // are attempts.
    DocsIf3CmStatusUsRangingAborteds interface{}

    // This attribute indicates modulation type status currently used by the CM
    // for this upstream channel. Since this object specifically identifies PHY
    // Layer mode, the shared upstream channel type 'tdmaAndAtdma' is not
    // permitted. The type is DocsisUpstreamType.
    DocsIf3CmStatusUsModulationType interface{}

    // This attribute indicates the pre-equalization data for the specified
    // upstream Channel on this CM after convolution with data indicated in the
    // RNG-RSP. This data is valid when docsIfUpChannelPreEqEnable RFC 4546 is set
    // to true. The type is string with length: 0 | 36..260.
    DocsIf3CmStatusUsEqData interface{}

    // This attribute denotes the number of times for excessive T3 timeouts.
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the managed system, and at other times as indicated by the value of
    // ifCounterDiscontinuityTime for the associated upstream channel. The type is
    // interface{} with range: 0..4294967295. Units are timeouts.
    DocsIf3CmStatusUsT3Exceededs interface{}

    // This attribute denotes whether the upstream channel is muted. The type is
    // bool.
    DocsIf3CmStatusUsIsMuted interface{}

    // This attribute denotes the ranging state of the CM. The type is
    // RangingState.
    DocsIf3CmStatusUsRangingStatus interface{}
}

func (docsIf3CmStatusUsEntry *DOCSIF3MIB_DocsIf3CmStatusUsTable_DocsIf3CmStatusUsEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmStatusUsEntry.EntityData.YFilter = docsIf3CmStatusUsEntry.YFilter
    docsIf3CmStatusUsEntry.EntityData.YangName = "docsIf3CmStatusUsEntry"
    docsIf3CmStatusUsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmStatusUsEntry.EntityData.ParentYangName = "docsIf3CmStatusUsTable"
    docsIf3CmStatusUsEntry.EntityData.SegmentPath = "docsIf3CmStatusUsEntry" + types.AddKeyToken(docsIf3CmStatusUsEntry.IfIndex, "ifIndex")
    docsIf3CmStatusUsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmStatusUsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmStatusUsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmStatusUsEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmStatusUsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmStatusUsEntry.IfIndex})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsTxPower", types.YLeaf{"DocsIf3CmStatusUsTxPower", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsTxPower})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsT3Timeouts", types.YLeaf{"DocsIf3CmStatusUsT3Timeouts", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsT3Timeouts})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsT4Timeouts", types.YLeaf{"DocsIf3CmStatusUsT4Timeouts", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsT4Timeouts})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsRangingAborteds", types.YLeaf{"DocsIf3CmStatusUsRangingAborteds", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsRangingAborteds})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsModulationType", types.YLeaf{"DocsIf3CmStatusUsModulationType", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsModulationType})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsEqData", types.YLeaf{"DocsIf3CmStatusUsEqData", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsEqData})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsT3Exceededs", types.YLeaf{"DocsIf3CmStatusUsT3Exceededs", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsT3Exceededs})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsIsMuted", types.YLeaf{"DocsIf3CmStatusUsIsMuted", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsIsMuted})
    docsIf3CmStatusUsEntry.EntityData.Leafs.Append("docsIf3CmStatusUsRangingStatus", types.YLeaf{"DocsIf3CmStatusUsRangingStatus", docsIf3CmStatusUsEntry.DocsIf3CmStatusUsRangingStatus})

    docsIf3CmStatusUsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmStatusUsEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable
// This object defines attributes that represent the CM's
// registration status as tracked by the CMTS.
// Refer to the individual attribute definitions for 
// applicability to 3.0 and 3.1 Cable Modems.
type DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsCmRegStatusTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry.
    DocsIf3CmtsCmRegStatusEntry []*DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry
}

func (docsIf3CmtsCmRegStatusTable *DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmRegStatusTable.EntityData.YFilter = docsIf3CmtsCmRegStatusTable.YFilter
    docsIf3CmtsCmRegStatusTable.EntityData.YangName = "docsIf3CmtsCmRegStatusTable"
    docsIf3CmtsCmRegStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmRegStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsCmRegStatusTable.EntityData.SegmentPath = "docsIf3CmtsCmRegStatusTable"
    docsIf3CmtsCmRegStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmRegStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmRegStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmRegStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmRegStatusTable.EntityData.Children.Append("docsIf3CmtsCmRegStatusEntry", types.YChild{"DocsIf3CmtsCmRegStatusEntry", nil})
    for i := range docsIf3CmtsCmRegStatusTable.DocsIf3CmtsCmRegStatusEntry {
        docsIf3CmtsCmRegStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsCmRegStatusTable.DocsIf3CmtsCmRegStatusEntry[i]), types.YChild{"DocsIf3CmtsCmRegStatusEntry", docsIf3CmtsCmRegStatusTable.DocsIf3CmtsCmRegStatusEntry[i]})
    }
    docsIf3CmtsCmRegStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsCmRegStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsCmRegStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry
// The conceptual row of docsIf3CmtsCmRegStatusTable.
type DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This attribute uniquely identifies a CM.  The CMTS
    // must assign a single id value for each CM MAC address seen by the CMTS. 
    // The CMTS should ensure that the association between an Id and MAC Address
    // remains constant during CMTS uptime. The type is interface{} with range:
    // 1..4294967295.
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute represents the MAC address of the CM. If the CM has multiple
    // MAC addresses, this is the MAC address associated with the MAC Domain
    // interface. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    DocsIf3CmtsCmRegStatusMacAddr interface{}

    // This attribute represents the IPv6 address of the CM. If the CM has no
    // Internet address assigned, or the Internet address is unknown, the value of
    // this attribute is the all zeros address. The type is string.
    DocsIf3CmtsCmRegStatusIPv6Addr interface{}

    // This attribute represents the IPv6 local scope address of the CM. If the CM
    // has no link local address assigned, or the Internet address is unknown, the
    // value of this attribute is the all zeros address. The type is string.
    DocsIf3CmtsCmRegStatusIPv6LinkLocal interface{}

    // This attribute represents the IPv4 address of this CM. If the CM has no IP
    // address assigned, or the IP address is unknown, this object returns
    // 0.0.0.0. The type is string.
    DocsIf3CmtsCmRegStatusIPv4Addr interface{}

    // This attribute represents the current CM connectivity state. The type is
    // CmtsCmRegState.
    DocsIf3CmtsCmRegStatusValue interface{}

    // This attribute represents the interface Index of the CMTS MAC Domain where
    // the CM is active. If the interface is unknown, the CMTS returns a value of
    // zero. The type is interface{} with range: 0..2147483647.
    DocsIf3CmtsCmRegStatusMdIfIndex interface{}

    // This attribute represents the ID of the MAC Domain CM Service Group Id
    // (MD-CM-SG-ID) in which the CM is registered. If the ID is unknown, the CMTS
    // returns a value of zero. The type is interface{} with range: 0..4294967295.
    DocsIf3CmtsCmRegStatusMdCmSgId interface{}

    // This attribute represents the RCP-ID associated with the CM if the CM is in
    // DOCSIS 3.0 mode. If the  RCP-ID is unknown or the CM is in DOCSIS 3.1 mode,
    // the CMTS returns a five octet long string of zeros. The type is string with
    // length: 5.
    DocsIf3CmtsCmRegStatusRcpId interface{}

    // This attribute represents the RCC Id the CMTS used to configure the CM
    // receive channel set during the registration process, if the CM is in DOCSIS
    // 3.0 mode. If unknown or the CM is in DOCSIS 3.1 mode, the CMTS returns the
    // value zero. The type is interface{} with range: 0..4294967295.
    DocsIf3CmtsCmRegStatusRccStatusId interface{}

    // This attribute represents the Receive Channel Set (RCS) that the CM is
    // currently using. If the RCS is unknown, the CMTS returns the value zero.
    // The type is interface{} with range: 0..4294967295.
    DocsIf3CmtsCmRegStatusRcsId interface{}

    // This attribute represents Transmit Channel Set (TCS) the CM is currently
    // using. If the TCS is unknown, the CMTS returns the value zero. The type is
    // interface{} with range: 0..4294967295.
    DocsIf3CmtsCmRegStatusTcsId interface{}

    // This attribute denotes the queuing services the CM  registered, either
    // DOCSIS 1.1 QoS or DOCSIS 1.0 CoS mode. The type is DocsisQosVersion.
    DocsIf3CmtsCmRegStatusQosVersion interface{}

    // This attribute represents the last time the CM registered. The type is
    // string.
    DocsIf3CmtsCmRegStatusLastRegTime interface{}

    // This attribute counts the number of upstream packets received on the SIDs
    // assigned to a CM that are any of the following: Upstream IPv4 ARP Requests
    // Upstream IPv6 Neighbor Solicitation Requests (For Routing CMTSs) Upstream
    // IPv4 or IPv6 packets to unresolved destinations in locally connected
    // downstream subnets in the HFC. Discontinuities in the value of this counter
    // can occur at re-initialization of the managed system, and at other times as
    // indicated by the value of ifCounterDiscontinuityTime for the associated MAC
    // Domain interface. The type is interface{} with range: 0..4294967295.
    DocsIf3CmtsCmRegStatusAddrResolutionReqs interface{}

    // This attribute indicates which, if any, of the Energy  Management Features
    // are enabled for this CM. If this attribute returns the em1x1Mode(0) bit
    // set, the CM is configured with the  Energy Management 1x1 Feature enabled.
    // If this attribute returns the dlsMode(1) bit set, the CM is configured with
    // the DLS Feature enabled. If this attribute returns all bits cleared, the CM
    // will  not request to operate in any Energy Management mode of operation.   
    // Note: This attribute only indicates if an Energy Management Feature  is
    // enabled/disabled via the CM config file and registration  request/response
    // exchange and does not indicate whether the CM is  actively operating in an
    // Energy Management Mode. The type is map[string]bool.
    DocsIf3CmtsCmRegStatusEnergyMgtEnabled interface{}

    // This attribute indicates whether the CM is currently operating in an Energy
    // Management Mode. If this attribute returns the em1x1Mode(0) bit set, the CM
    // is operating in Energy Management 1x1 Mode. If this attribute returns the
    // dlsMode(1) bit set, the CM is operating in DLS Mode. If this attribute
    // returns all bits  cleared, the CM is not operating in any Energy Management
    // Mode.  This attribute always returns 0x00 (no bits set) in the case when 
    // docsIf3CmtsCmRegStatusEnergyMgtEnabled is set to 0x00 (no Energy 
    // Management Features enabled).  Note: dlsMode(1) and em1x1Mode(0) are
    // mutually exclusive, thus  a return value where both of these bits are
    // 'true' is invalid. The type is map[string]bool.
    DocsIf3CmtsCmRegStatusEnergyMgtOperStatus interface{}
}

func (docsIf3CmtsCmRegStatusEntry *DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmRegStatusEntry.EntityData.YFilter = docsIf3CmtsCmRegStatusEntry.YFilter
    docsIf3CmtsCmRegStatusEntry.EntityData.YangName = "docsIf3CmtsCmRegStatusEntry"
    docsIf3CmtsCmRegStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmRegStatusEntry.EntityData.ParentYangName = "docsIf3CmtsCmRegStatusTable"
    docsIf3CmtsCmRegStatusEntry.EntityData.SegmentPath = "docsIf3CmtsCmRegStatusEntry" + types.AddKeyToken(docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId")
    docsIf3CmtsCmRegStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmRegStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmRegStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmRegStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusMacAddr", types.YLeaf{"DocsIf3CmtsCmRegStatusMacAddr", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusMacAddr})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusIPv6Addr", types.YLeaf{"DocsIf3CmtsCmRegStatusIPv6Addr", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusIPv6Addr})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusIPv6LinkLocal", types.YLeaf{"DocsIf3CmtsCmRegStatusIPv6LinkLocal", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusIPv6LinkLocal})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusIPv4Addr", types.YLeaf{"DocsIf3CmtsCmRegStatusIPv4Addr", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusIPv4Addr})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusValue", types.YLeaf{"DocsIf3CmtsCmRegStatusValue", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusValue})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusMdIfIndex", types.YLeaf{"DocsIf3CmtsCmRegStatusMdIfIndex", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusMdIfIndex})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusMdCmSgId", types.YLeaf{"DocsIf3CmtsCmRegStatusMdCmSgId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusMdCmSgId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusRcpId", types.YLeaf{"DocsIf3CmtsCmRegStatusRcpId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusRcpId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusRccStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusRccStatusId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusRccStatusId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusRcsId", types.YLeaf{"DocsIf3CmtsCmRegStatusRcsId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusRcsId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusTcsId", types.YLeaf{"DocsIf3CmtsCmRegStatusTcsId", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusTcsId})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusQosVersion", types.YLeaf{"DocsIf3CmtsCmRegStatusQosVersion", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusQosVersion})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusLastRegTime", types.YLeaf{"DocsIf3CmtsCmRegStatusLastRegTime", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusLastRegTime})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusAddrResolutionReqs", types.YLeaf{"DocsIf3CmtsCmRegStatusAddrResolutionReqs", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusAddrResolutionReqs})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusEnergyMgtEnabled", types.YLeaf{"DocsIf3CmtsCmRegStatusEnergyMgtEnabled", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusEnergyMgtEnabled})
    docsIf3CmtsCmRegStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusEnergyMgtOperStatus", types.YLeaf{"DocsIf3CmtsCmRegStatusEnergyMgtOperStatus", docsIf3CmtsCmRegStatusEntry.DocsIf3CmtsCmRegStatusEnergyMgtOperStatus})

    docsIf3CmtsCmRegStatusEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId"}

    return &(docsIf3CmtsCmRegStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable
// This object defines status information of the CM
// currently in use Upstream Logical Channels, as reported
// by the CMTS.
type DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsCmUsStatusTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable_DocsIf3CmtsCmUsStatusEntry.
    DocsIf3CmtsCmUsStatusEntry []*DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable_DocsIf3CmtsCmUsStatusEntry
}

func (docsIf3CmtsCmUsStatusTable *DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmUsStatusTable.EntityData.YFilter = docsIf3CmtsCmUsStatusTable.YFilter
    docsIf3CmtsCmUsStatusTable.EntityData.YangName = "docsIf3CmtsCmUsStatusTable"
    docsIf3CmtsCmUsStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmUsStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsCmUsStatusTable.EntityData.SegmentPath = "docsIf3CmtsCmUsStatusTable"
    docsIf3CmtsCmUsStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmUsStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmUsStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmUsStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmUsStatusTable.EntityData.Children.Append("docsIf3CmtsCmUsStatusEntry", types.YChild{"DocsIf3CmtsCmUsStatusEntry", nil})
    for i := range docsIf3CmtsCmUsStatusTable.DocsIf3CmtsCmUsStatusEntry {
        docsIf3CmtsCmUsStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsCmUsStatusTable.DocsIf3CmtsCmUsStatusEntry[i]), types.YChild{"DocsIf3CmtsCmUsStatusEntry", docsIf3CmtsCmUsStatusTable.DocsIf3CmtsCmUsStatusEntry[i]})
    }
    docsIf3CmtsCmUsStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsCmUsStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsCmUsStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable_DocsIf3CmtsCmUsStatusEntry
// The conceptual row of docsIf3CmtsCmUsStatusTable.
type DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable_DocsIf3CmtsCmUsStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry_DocsIf3CmtsCmRegStatusId
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute is a key. This attribute is a key that represents the
    // ifIndex of the upstream interface. The type is interface{} with range:
    // 1..2147483647.
    DocsIf3CmtsCmUsStatusChIfIndex interface{}

    // This attribute represents the modulation type currently used by this
    // upstream channel. The type is DocsisUpstreamType.
    DocsIf3CmtsCmUsStatusModulationType interface{}

    // This attribute represents the receive power of this upstream channel. The
    // reported value represents the total average power for the  channel
    // regardless of whether the CM is reporting Pr, total  average power, or
    // P1.6r, the power spectral density in an  equivalent 1.6 MHz spectrum. The
    // type is interface{} with range: -2147483648..2147483647. Units are
    // TenthdBmV.
    DocsIf3CmtsCmUsStatusRxPower interface{}

    // This attribute represents Signal/Noise ratio as perceived for upstream data
    // from the CM on this upstream channel. The type is interface{} with range:
    // -2147483648..2147483647. Units are TenthdB.
    DocsIf3CmtsCmUsStatusSignalNoise interface{}

    // This attribute represents microreflections received on this upstream
    // channel. The type is interface{} with range: 0..65535. Units are -dBc.
    DocsIf3CmtsCmUsStatusMicroreflections interface{}

    // This attribute represents the equalization data for the CM on this upstream
    // channel. The type is string with length: 0 | 36..260.
    DocsIf3CmtsCmUsStatusEqData interface{}

    // This attribute represents the codewords received without error from the CM
    // on this interface. Discontinuities in the value of this counter can occur
    // at re-initialization of the managed system, and at other times as indicated
    // by the value of ifCounterDiscontinuityTime for the associated upstream
    // channel. The type is interface{} with range: 0..4294967295.
    DocsIf3CmtsCmUsStatusUnerroreds interface{}

    // This attribute represents the codewords received with correctable errors
    // from the CM on this upstream channel. Discontinuities in the value of this
    // counter can occur at re-initialization of the managed system, and at other
    // times as indicated by the value of ifCounterDiscontinuityTime for the
    // associated upstream channel. The type is interface{} with range:
    // 0..4294967295.
    DocsIf3CmtsCmUsStatusCorrecteds interface{}

    // This attribute represents the codewords received with uncorrectable errors
    // from the CM on this upstream channel. Discontinuities in the value of this
    // counter can occur at re-initialization of the managed system, and at other
    // times as indicated by the value of ifCounterDiscontinuityTime for the
    // associated upstream channel. The type is interface{} with range:
    // 0..4294967295.
    DocsIf3CmtsCmUsStatusUncorrectables interface{}

    // This attribute represents the current measured round trip time on this CM's
    // upstream channel in units of (6.25 microseconds/(64*256)).  This attribute
    // returns zero if the value is unknown. The type is interface{} with range:
    // -2147483648..2147483647. Units are time tick/(64*256).
    DocsIf3CmtsCmUsStatusHighResolutionTimingOffset interface{}

    // This attribute has a value 'true' to indicate that  the CM's upstream
    // channel has been muted via  CM-CTRL-REQ/CM-CTRL-RSP message exchange. The
    // type is bool.
    DocsIf3CmtsCmUsStatusIsMuted interface{}

    // This attribute denotes the ranging state of the CM. The type is
    // RangingState.
    DocsIf3CmtsCmUsStatusRangingStatus interface{}
}

func (docsIf3CmtsCmUsStatusEntry *DOCSIF3MIB_DocsIf3CmtsCmUsStatusTable_DocsIf3CmtsCmUsStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmUsStatusEntry.EntityData.YFilter = docsIf3CmtsCmUsStatusEntry.YFilter
    docsIf3CmtsCmUsStatusEntry.EntityData.YangName = "docsIf3CmtsCmUsStatusEntry"
    docsIf3CmtsCmUsStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmUsStatusEntry.EntityData.ParentYangName = "docsIf3CmtsCmUsStatusTable"
    docsIf3CmtsCmUsStatusEntry.EntityData.SegmentPath = "docsIf3CmtsCmUsStatusEntry" + types.AddKeyToken(docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId") + types.AddKeyToken(docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusChIfIndex, "docsIf3CmtsCmUsStatusChIfIndex")
    docsIf3CmtsCmUsStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmUsStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmUsStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmUsStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmRegStatusId})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusChIfIndex", types.YLeaf{"DocsIf3CmtsCmUsStatusChIfIndex", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusChIfIndex})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusModulationType", types.YLeaf{"DocsIf3CmtsCmUsStatusModulationType", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusModulationType})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusRxPower", types.YLeaf{"DocsIf3CmtsCmUsStatusRxPower", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusRxPower})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusSignalNoise", types.YLeaf{"DocsIf3CmtsCmUsStatusSignalNoise", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusSignalNoise})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusMicroreflections", types.YLeaf{"DocsIf3CmtsCmUsStatusMicroreflections", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusMicroreflections})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusEqData", types.YLeaf{"DocsIf3CmtsCmUsStatusEqData", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusEqData})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusUnerroreds", types.YLeaf{"DocsIf3CmtsCmUsStatusUnerroreds", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusUnerroreds})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusCorrecteds", types.YLeaf{"DocsIf3CmtsCmUsStatusCorrecteds", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusCorrecteds})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusUncorrectables", types.YLeaf{"DocsIf3CmtsCmUsStatusUncorrectables", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusUncorrectables})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusHighResolutionTimingOffset", types.YLeaf{"DocsIf3CmtsCmUsStatusHighResolutionTimingOffset", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusHighResolutionTimingOffset})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusIsMuted", types.YLeaf{"DocsIf3CmtsCmUsStatusIsMuted", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusIsMuted})
    docsIf3CmtsCmUsStatusEntry.EntityData.Leafs.Append("docsIf3CmtsCmUsStatusRangingStatus", types.YLeaf{"DocsIf3CmtsCmUsStatusRangingStatus", docsIf3CmtsCmUsStatusEntry.DocsIf3CmtsCmUsStatusRangingStatus})

    docsIf3CmtsCmUsStatusEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId", "DocsIf3CmtsCmUsStatusChIfIndex"}

    return &(docsIf3CmtsCmUsStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdChCfgTable
// This object configures the association of downstream
// and upstream channels to a particular MAC Domain
// (MD) on a CMTS.  The creation of channels and MAC domain
// object interface instances is vendor-specific.
// In particular, the assignment of the channel interface
// index is normally vendor-specific. Therefore,
// this object is intended only for associating channels
// to a MAC Domain and assumes that those channels were
// previously configured.
// The CMTS may have restrictions on which channels can
// be configured in the same MAC Domain.  For example, it
// could require the upstream channels to be from the same
// line card.
// This object supports the creation and deletion of multiple
// instances.
// Creation of a new instance of this object requires the
// ChId attribute to be set.
type DOCSIF3MIB_DocsIf3MdChCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdChCfgTable. The ifIndex key corresponds to
    // the MAC Domain interface where the channel is configured. The CMTS persists
    // all instances of MdChCfg across reinitializations. The type is slice of
    // DOCSIF3MIB_DocsIf3MdChCfgTable_DocsIf3MdChCfgEntry.
    DocsIf3MdChCfgEntry []*DOCSIF3MIB_DocsIf3MdChCfgTable_DocsIf3MdChCfgEntry
}

func (docsIf3MdChCfgTable *DOCSIF3MIB_DocsIf3MdChCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdChCfgTable.EntityData.YFilter = docsIf3MdChCfgTable.YFilter
    docsIf3MdChCfgTable.EntityData.YangName = "docsIf3MdChCfgTable"
    docsIf3MdChCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdChCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdChCfgTable.EntityData.SegmentPath = "docsIf3MdChCfgTable"
    docsIf3MdChCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdChCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdChCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdChCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdChCfgTable.EntityData.Children.Append("docsIf3MdChCfgEntry", types.YChild{"DocsIf3MdChCfgEntry", nil})
    for i := range docsIf3MdChCfgTable.DocsIf3MdChCfgEntry {
        docsIf3MdChCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdChCfgTable.DocsIf3MdChCfgEntry[i]), types.YChild{"DocsIf3MdChCfgEntry", docsIf3MdChCfgTable.DocsIf3MdChCfgEntry[i]})
    }
    docsIf3MdChCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdChCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdChCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdChCfgTable_DocsIf3MdChCfgEntry
// The conceptual row of docsIf3MdChCfgTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the channel is configured.
// The CMTS persists all instances of MdChCfg across
// reinitializations.
type DOCSIF3MIB_DocsIf3MdChCfgTable_DocsIf3MdChCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the interface index of an
    // existing upstream or downstream channel that is configured to be part of
    // the MAC Domain. For the case of upstream interfaces the CMTS could reject
    // the assignment of upstream logical channels under the same physical
    // upstream interface to different MAC Domains. The type is interface{} with
    // range: 1..2147483647.
    DocsIf3MdChCfgChIfIndex interface{}

    // If set to 'true', this attribute configures the downstream channel as
    // Primary-Capable.  The default value for a downstream channel is 'true'.
    // This attribute is not relevant for upstream interfaces,  therefore it
    // reports the value 'false' for such interfaces. A CMTS may restrict the
    // permitted value of this attribute based upon physical channel capabilities.
    // The type is bool.
    DocsIf3MdChCfgIsPriCapableDs interface{}

    // This attribute contains the 8-bit Downstream Channel ID (DCID) or Upstream
    // Channel ID (UCID) configured for the channel in the MAC Domain. The type is
    // interface{} with range: 0..255.
    DocsIf3MdChCfgChId interface{}

    // This attribute contains Provisioned Attribute Mask of non-bonded service
    // flow assignment to this channel. The type is map[string]bool.
    DocsIf3MdChCfgSfProvAttrMask interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3MdChCfgRowStatus interface{}
}

func (docsIf3MdChCfgEntry *DOCSIF3MIB_DocsIf3MdChCfgTable_DocsIf3MdChCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdChCfgEntry.EntityData.YFilter = docsIf3MdChCfgEntry.YFilter
    docsIf3MdChCfgEntry.EntityData.YangName = "docsIf3MdChCfgEntry"
    docsIf3MdChCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdChCfgEntry.EntityData.ParentYangName = "docsIf3MdChCfgTable"
    docsIf3MdChCfgEntry.EntityData.SegmentPath = "docsIf3MdChCfgEntry" + types.AddKeyToken(docsIf3MdChCfgEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3MdChCfgEntry.DocsIf3MdChCfgChIfIndex, "docsIf3MdChCfgChIfIndex")
    docsIf3MdChCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdChCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdChCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdChCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdChCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3MdChCfgEntry.IfIndex})
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("docsIf3MdChCfgChIfIndex", types.YLeaf{"DocsIf3MdChCfgChIfIndex", docsIf3MdChCfgEntry.DocsIf3MdChCfgChIfIndex})
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("docsIf3MdChCfgIsPriCapableDs", types.YLeaf{"DocsIf3MdChCfgIsPriCapableDs", docsIf3MdChCfgEntry.DocsIf3MdChCfgIsPriCapableDs})
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("docsIf3MdChCfgChId", types.YLeaf{"DocsIf3MdChCfgChId", docsIf3MdChCfgEntry.DocsIf3MdChCfgChId})
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("docsIf3MdChCfgSfProvAttrMask", types.YLeaf{"DocsIf3MdChCfgSfProvAttrMask", docsIf3MdChCfgEntry.DocsIf3MdChCfgSfProvAttrMask})
    docsIf3MdChCfgEntry.EntityData.Leafs.Append("docsIf3MdChCfgRowStatus", types.YLeaf{"DocsIf3MdChCfgRowStatus", docsIf3MdChCfgEntry.DocsIf3MdChCfgRowStatus})

    docsIf3MdChCfgEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3MdChCfgChIfIndex"}

    return &(docsIf3MdChCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RccCfgTable
// This object identifies the scope of the Receive Channel
// Configuration (RCC) and provides a top level container
// for the Receive Module and Receive Channel
// objects.  The CMTS selects an instance of this object
// to assign to a CM when it registers.
// This object supports the creation and deletion of multiple
// instances.
type DOCSIF3MIB_DocsIf3RccCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RccCfgTable. The ifIndex key corresponds to
    // the MAC Domain interface where the RCC is configured.  The CMTS persists
    // all instances of RccCfg across  reinitializations. The type is slice of
    // DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry.
    DocsIf3RccCfgEntry []*DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry
}

func (docsIf3RccCfgTable *DOCSIF3MIB_DocsIf3RccCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3RccCfgTable.EntityData.YFilter = docsIf3RccCfgTable.YFilter
    docsIf3RccCfgTable.EntityData.YangName = "docsIf3RccCfgTable"
    docsIf3RccCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RccCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RccCfgTable.EntityData.SegmentPath = "docsIf3RccCfgTable"
    docsIf3RccCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RccCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RccCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RccCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RccCfgTable.EntityData.Children.Append("docsIf3RccCfgEntry", types.YChild{"DocsIf3RccCfgEntry", nil})
    for i := range docsIf3RccCfgTable.DocsIf3RccCfgEntry {
        docsIf3RccCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RccCfgTable.DocsIf3RccCfgEntry[i]), types.YChild{"DocsIf3RccCfgEntry", docsIf3RccCfgTable.DocsIf3RccCfgEntry[i]})
    }
    docsIf3RccCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RccCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3RccCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry
// The conceptual row of docsIf3RccCfgTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured.
//  The CMTS persists all instances of RccCfg across
//  reinitializations.
type DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the 'Receive Channel Profile
    // Identifier' (RCP-ID) configured for the MAC Domain indicated by this
    // instance. The type is string with length: 5.
    DocsIf3RccCfgRcpId interface{}

    // This attribute is a key. This key denotes an RCC combination assignment for
    // a particular RcpId and is unique per combination of MAC Domain and RcpId.
    // The type is interface{} with range: 1..4294967295.
    DocsIf3RccCfgRccCfgId interface{}

    // This attribute contains vendor-specific information of the CM Receive
    // Channel configuration. The type is string with length: 0..252.
    DocsIf3RccCfgVendorSpecific interface{}

    // This attribute contains a human-readable description of the CM RCP
    // Configuration. The type is string with length: 0..15.
    DocsIf3RccCfgDescription interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3RccCfgRowStatus interface{}
}

func (docsIf3RccCfgEntry *DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RccCfgEntry.EntityData.YFilter = docsIf3RccCfgEntry.YFilter
    docsIf3RccCfgEntry.EntityData.YangName = "docsIf3RccCfgEntry"
    docsIf3RccCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RccCfgEntry.EntityData.ParentYangName = "docsIf3RccCfgTable"
    docsIf3RccCfgEntry.EntityData.SegmentPath = "docsIf3RccCfgEntry" + types.AddKeyToken(docsIf3RccCfgEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RccCfgEntry.DocsIf3RccCfgRcpId, "docsIf3RccCfgRcpId") + types.AddKeyToken(docsIf3RccCfgEntry.DocsIf3RccCfgRccCfgId, "docsIf3RccCfgRccCfgId")
    docsIf3RccCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RccCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RccCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RccCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RccCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RccCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RccCfgEntry.IfIndex})
    docsIf3RccCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRcpId", types.YLeaf{"DocsIf3RccCfgRcpId", docsIf3RccCfgEntry.DocsIf3RccCfgRcpId})
    docsIf3RccCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRccCfgId", types.YLeaf{"DocsIf3RccCfgRccCfgId", docsIf3RccCfgEntry.DocsIf3RccCfgRccCfgId})
    docsIf3RccCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgVendorSpecific", types.YLeaf{"DocsIf3RccCfgVendorSpecific", docsIf3RccCfgEntry.DocsIf3RccCfgVendorSpecific})
    docsIf3RccCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgDescription", types.YLeaf{"DocsIf3RccCfgDescription", docsIf3RccCfgEntry.DocsIf3RccCfgDescription})
    docsIf3RccCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRowStatus", types.YLeaf{"DocsIf3RccCfgRowStatus", docsIf3RccCfgEntry.DocsIf3RccCfgRowStatus})

    docsIf3RccCfgEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccCfgRcpId", "DocsIf3RccCfgRccCfgId"}

    return &(docsIf3RccCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RccStatusTable
// The RCC Status object provides a read-only view of
// the statically-configured (from the RccCfg object)
// and dynamically-created RCCs.
// The CMTS creates an RCC Status instance for each unique
// MAC Domain Cable Modem Service Group (MD-CM-SG) to
// which it signals an RCC to the CM.
type DOCSIF3MIB_DocsIf3RccStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RccStatusTable. The ifIndex key corresponds to
    // the MAC Domain interface where the RCC is configured. The type is slice of
    // DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry.
    DocsIf3RccStatusEntry []*DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry
}

func (docsIf3RccStatusTable *DOCSIF3MIB_DocsIf3RccStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3RccStatusTable.EntityData.YFilter = docsIf3RccStatusTable.YFilter
    docsIf3RccStatusTable.EntityData.YangName = "docsIf3RccStatusTable"
    docsIf3RccStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RccStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RccStatusTable.EntityData.SegmentPath = "docsIf3RccStatusTable"
    docsIf3RccStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RccStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RccStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RccStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RccStatusTable.EntityData.Children.Append("docsIf3RccStatusEntry", types.YChild{"DocsIf3RccStatusEntry", nil})
    for i := range docsIf3RccStatusTable.DocsIf3RccStatusEntry {
        docsIf3RccStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RccStatusTable.DocsIf3RccStatusEntry[i]), types.YChild{"DocsIf3RccStatusEntry", docsIf3RccStatusTable.DocsIf3RccStatusEntry[i]})
    }
    docsIf3RccStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RccStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3RccStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry
// The conceptual row of docsIf3RccStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured.
type DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the RCP-ID to which this
    // instance applies. The type is string with length: 5.
    DocsIf3RccStatusRcpId interface{}

    // This attribute is a key. This key represents an RCC combination for a
    // particular RcpId either from an RCC configuration object or a
    // CMTS-determined RCC and is unique per combination of MAC Domain IfIndex and
    // RcpId. The type is interface{} with range: 1..4294967295.
    DocsIf3RccStatusRccStatusId interface{}

    // This attribute identifies an RCC-Configured combination from which this
    // instance was defined. If nonzero, it corresponds to the RccCfg instance
    // from which the RCC was created. Zero means that the  RCC was dynamically
    // created by the CMTS. The type is interface{} with range: 0..65535.
    DocsIf3RccStatusRccCfgId interface{}

    // This attribute indicates whether the RCC instance of this object is valid
    // or not. An RCC Status instance from a configured or a dynamic RCC could
    // become invalid, for example, due changes in the topology. The type is
    // DocsIf3RccStatusValidityCode.
    DocsIf3RccStatusValidityCode interface{}

    // This attribute contains the CMTS vendor-specific log information from the
    // Receive Channel Configuration Status encoding. The type is string.
    DocsIf3RccStatusValidityCodeText interface{}
}

func (docsIf3RccStatusEntry *DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RccStatusEntry.EntityData.YFilter = docsIf3RccStatusEntry.YFilter
    docsIf3RccStatusEntry.EntityData.YangName = "docsIf3RccStatusEntry"
    docsIf3RccStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RccStatusEntry.EntityData.ParentYangName = "docsIf3RccStatusTable"
    docsIf3RccStatusEntry.EntityData.SegmentPath = "docsIf3RccStatusEntry" + types.AddKeyToken(docsIf3RccStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RccStatusEntry.DocsIf3RccStatusRcpId, "docsIf3RccStatusRcpId") + types.AddKeyToken(docsIf3RccStatusEntry.DocsIf3RccStatusRccStatusId, "docsIf3RccStatusRccStatusId")
    docsIf3RccStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RccStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RccStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RccStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RccStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RccStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RccStatusEntry.IfIndex})
    docsIf3RccStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRcpId", types.YLeaf{"DocsIf3RccStatusRcpId", docsIf3RccStatusEntry.DocsIf3RccStatusRcpId})
    docsIf3RccStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRccStatusId", types.YLeaf{"DocsIf3RccStatusRccStatusId", docsIf3RccStatusEntry.DocsIf3RccStatusRccStatusId})
    docsIf3RccStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRccCfgId", types.YLeaf{"DocsIf3RccStatusRccCfgId", docsIf3RccStatusEntry.DocsIf3RccStatusRccCfgId})
    docsIf3RccStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusValidityCode", types.YLeaf{"DocsIf3RccStatusValidityCode", docsIf3RccStatusEntry.DocsIf3RccStatusValidityCode})
    docsIf3RccStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusValidityCodeText", types.YLeaf{"DocsIf3RccStatusValidityCodeText", docsIf3RccStatusEntry.DocsIf3RccStatusValidityCodeText})

    docsIf3RccStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccStatusRcpId", "DocsIf3RccStatusRccStatusId"}

    return &(docsIf3RccStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode represents for example, due changes in the topology.
type DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode string

const (
    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_other DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "other"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_valid DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "valid"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_invalid DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "invalid"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_wrongPrimaryDs DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "wrongPrimaryDs"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_missingPrimaryDs DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "missingPrimaryDs"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_multiplePrimaryDs DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "multiplePrimaryDs"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_duplicateDs DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "duplicateDs"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_wrongFrequencyRange DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "wrongFrequencyRange"

    DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode_wrongConnectivity DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusValidityCode = "wrongConnectivity"
)

// DOCSIF3MIB_DocsIf3RxChCfgTable
// The Receive Channel Configuration object permits
// an operator to configure how CMs registered with certain
// Receive Channel Profiles will configure the Receive
// Channels within their profile. When a CM registers
// with an RCP for which all Receive Channel Indices
// (RcIds) are configured in the Receive Module object
// and all Receive Channels are configured within this
// object, the CMTS should use the configuration within
// these objects to set the Receive Channel Configuration
// returned to the CM in a REG-RSP message.  A CMTS
// may require configuration of all pertinent Receive
// Module and Receive Channel instances in order to register
// a CM that reports a Receive Channel Profile (RCP),
// including any standard Receive Channel Profiles.
// If the CM reports multiple RCPs, and Receive Module
// and Receive Channel objects have instances for more
// than one RCP, the particular RCP selected by the CMTS
// is not specified.  A CMTS is not restricted to assigning
// Receive Modules based only on the contents of this
// object.
// This object supports the creation and deletion of multiple
// instances.
// Creation of a new instance of this object requires the
// ChIfIndex attribute to be set and a valid reference of
// a RccCfg instance.
type DOCSIF3MIB_DocsIf3RxChCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RxChCfgTable. The ifIndex key corresponds to
    // the MAC Domain interface where the RCC is configured. The CMTS persists all
    // instances of ReceiveChannelCfg across reinitializations. The type is slice
    // of DOCSIF3MIB_DocsIf3RxChCfgTable_DocsIf3RxChCfgEntry.
    DocsIf3RxChCfgEntry []*DOCSIF3MIB_DocsIf3RxChCfgTable_DocsIf3RxChCfgEntry
}

func (docsIf3RxChCfgTable *DOCSIF3MIB_DocsIf3RxChCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3RxChCfgTable.EntityData.YFilter = docsIf3RxChCfgTable.YFilter
    docsIf3RxChCfgTable.EntityData.YangName = "docsIf3RxChCfgTable"
    docsIf3RxChCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxChCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RxChCfgTable.EntityData.SegmentPath = "docsIf3RxChCfgTable"
    docsIf3RxChCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxChCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxChCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxChCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RxChCfgTable.EntityData.Children.Append("docsIf3RxChCfgEntry", types.YChild{"DocsIf3RxChCfgEntry", nil})
    for i := range docsIf3RxChCfgTable.DocsIf3RxChCfgEntry {
        docsIf3RxChCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RxChCfgTable.DocsIf3RxChCfgEntry[i]), types.YChild{"DocsIf3RxChCfgEntry", docsIf3RxChCfgTable.DocsIf3RxChCfgEntry[i]})
    }
    docsIf3RxChCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RxChCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3RxChCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RxChCfgTable_DocsIf3RxChCfgEntry
// The conceptual row of docsIf3RxChCfgTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured.
// The CMTS persists all instances of ReceiveChannelCfg across
// reinitializations.
type DOCSIF3MIB_DocsIf3RxChCfgTable_DocsIf3RxChCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with length: 5. Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry_DocsIf3RccCfgRcpId
    DocsIf3RccCfgRcpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry_DocsIf3RccCfgRccCfgId
    DocsIf3RccCfgRccCfgId interface{}

    // This attribute is a key. This key represents an identifier for the
    // parameters of the Receive Channel instance within the Receive Channel
    // Profile. The type is interface{} with range: 1..255.
    DocsIf3RxChCfgRcId interface{}

    // This attribute contains the interface index of a Downstream Channel that
    // this Receive Channel Instance defines. The type is interface{} with range:
    // 1..2147483647.
    DocsIf3RxChCfgChIfIndex interface{}

    // If set to 'true', this attribute indicates the Receive Channel is to be the
    // primary-capable downstream channel for the CM receiving this RCC.
    // Otherwise, the downstream channel is to be a non-primary-capable channel.
    // The type is bool.
    DocsIf3RxChCfgPrimaryDsIndicator interface{}

    // This attribute indicates the Receive Module (via the RmId from the
    // ReceiveModule object) to which this Receive Channel connects.  If this
    // object contains a zero value (and thus no Receive Channel Connectivity),
    // the Receive Channel Connectivity TLV is omitted from the RCC. The type is
    // interface{} with range: 0..255.
    DocsIf3RxChCfgRcRmConnectivityId interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3RxChCfgRowStatus interface{}
}

func (docsIf3RxChCfgEntry *DOCSIF3MIB_DocsIf3RxChCfgTable_DocsIf3RxChCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RxChCfgEntry.EntityData.YFilter = docsIf3RxChCfgEntry.YFilter
    docsIf3RxChCfgEntry.EntityData.YangName = "docsIf3RxChCfgEntry"
    docsIf3RxChCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxChCfgEntry.EntityData.ParentYangName = "docsIf3RxChCfgTable"
    docsIf3RxChCfgEntry.EntityData.SegmentPath = "docsIf3RxChCfgEntry" + types.AddKeyToken(docsIf3RxChCfgEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RxChCfgEntry.DocsIf3RccCfgRcpId, "docsIf3RccCfgRcpId") + types.AddKeyToken(docsIf3RxChCfgEntry.DocsIf3RccCfgRccCfgId, "docsIf3RccCfgRccCfgId") + types.AddKeyToken(docsIf3RxChCfgEntry.DocsIf3RxChCfgRcId, "docsIf3RxChCfgRcId")
    docsIf3RxChCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxChCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxChCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxChCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RxChCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RxChCfgEntry.IfIndex})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRcpId", types.YLeaf{"DocsIf3RccCfgRcpId", docsIf3RxChCfgEntry.DocsIf3RccCfgRcpId})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRccCfgId", types.YLeaf{"DocsIf3RccCfgRccCfgId", docsIf3RxChCfgEntry.DocsIf3RccCfgRccCfgId})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RxChCfgRcId", types.YLeaf{"DocsIf3RxChCfgRcId", docsIf3RxChCfgEntry.DocsIf3RxChCfgRcId})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RxChCfgChIfIndex", types.YLeaf{"DocsIf3RxChCfgChIfIndex", docsIf3RxChCfgEntry.DocsIf3RxChCfgChIfIndex})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RxChCfgPrimaryDsIndicator", types.YLeaf{"DocsIf3RxChCfgPrimaryDsIndicator", docsIf3RxChCfgEntry.DocsIf3RxChCfgPrimaryDsIndicator})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RxChCfgRcRmConnectivityId", types.YLeaf{"DocsIf3RxChCfgRcRmConnectivityId", docsIf3RxChCfgEntry.DocsIf3RxChCfgRcRmConnectivityId})
    docsIf3RxChCfgEntry.EntityData.Leafs.Append("docsIf3RxChCfgRowStatus", types.YLeaf{"DocsIf3RxChCfgRowStatus", docsIf3RxChCfgEntry.DocsIf3RxChCfgRowStatus})

    docsIf3RxChCfgEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccCfgRcpId", "DocsIf3RccCfgRccCfgId", "DocsIf3RxChCfgRcId"}

    return &(docsIf3RxChCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RxChStatusTable
// The Receive Channel Status object reports the status
// of the statically-configured and dynamically-created
// Receive Channels within an RCC.
type DOCSIF3MIB_DocsIf3RxChStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RxChStatusTable. The ifIndex key corresponds
    // to the MAC Domain interface where the RCC is configured. When this object
    // is defined on the CM, the value of RccStatusId is always 1. The type is
    // slice of DOCSIF3MIB_DocsIf3RxChStatusTable_DocsIf3RxChStatusEntry.
    DocsIf3RxChStatusEntry []*DOCSIF3MIB_DocsIf3RxChStatusTable_DocsIf3RxChStatusEntry
}

func (docsIf3RxChStatusTable *DOCSIF3MIB_DocsIf3RxChStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3RxChStatusTable.EntityData.YFilter = docsIf3RxChStatusTable.YFilter
    docsIf3RxChStatusTable.EntityData.YangName = "docsIf3RxChStatusTable"
    docsIf3RxChStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxChStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RxChStatusTable.EntityData.SegmentPath = "docsIf3RxChStatusTable"
    docsIf3RxChStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxChStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxChStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxChStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RxChStatusTable.EntityData.Children.Append("docsIf3RxChStatusEntry", types.YChild{"DocsIf3RxChStatusEntry", nil})
    for i := range docsIf3RxChStatusTable.DocsIf3RxChStatusEntry {
        docsIf3RxChStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RxChStatusTable.DocsIf3RxChStatusEntry[i]), types.YChild{"DocsIf3RxChStatusEntry", docsIf3RxChStatusTable.DocsIf3RxChStatusEntry[i]})
    }
    docsIf3RxChStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RxChStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3RxChStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RxChStatusTable_DocsIf3RxChStatusEntry
// The conceptual row of docsIf3RxChStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured. When this object is defined
// on the CM, the value of RccStatusId is always 1.
type DOCSIF3MIB_DocsIf3RxChStatusTable_DocsIf3RxChStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with length: 5. Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusRcpId
    DocsIf3RccStatusRcpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusRccStatusId
    DocsIf3RccStatusRccStatusId interface{}

    // This attribute is a key. This key represents an identifier for the
    // parameters of the Receive Channel instance within the Receive Channel
    // Profile. The type is interface{} with range: 1..255.
    DocsIf3RxChStatusRcId interface{}

    // This attribute contains the interface index of the Downstream Channel that
    // this Receive Channel Instance defines. The type is interface{} with range:
    // 0..2147483647.
    DocsIf3RxChStatusChIfIndex interface{}

    // If set to 'true', this attribute indicates the Receive Channel is to be the
    // primary-capable downstream channel for the CM receiving this RCC.
    // Otherwise, the downstream channel is to be a non-primary-capable channel.
    // The type is bool.
    DocsIf3RxChStatusPrimaryDsIndicator interface{}

    // This attribute identifies the Receive Module to which this Receive Channel
    // connects.  A value a zero indicates that the Receive Channel Connectivity
    // TLV is omitted from the RCC. The type is interface{} with range: 0..255.
    DocsIf3RxChStatusRcRmConnectivityId interface{}
}

func (docsIf3RxChStatusEntry *DOCSIF3MIB_DocsIf3RxChStatusTable_DocsIf3RxChStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RxChStatusEntry.EntityData.YFilter = docsIf3RxChStatusEntry.YFilter
    docsIf3RxChStatusEntry.EntityData.YangName = "docsIf3RxChStatusEntry"
    docsIf3RxChStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxChStatusEntry.EntityData.ParentYangName = "docsIf3RxChStatusTable"
    docsIf3RxChStatusEntry.EntityData.SegmentPath = "docsIf3RxChStatusEntry" + types.AddKeyToken(docsIf3RxChStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RxChStatusEntry.DocsIf3RccStatusRcpId, "docsIf3RccStatusRcpId") + types.AddKeyToken(docsIf3RxChStatusEntry.DocsIf3RccStatusRccStatusId, "docsIf3RccStatusRccStatusId") + types.AddKeyToken(docsIf3RxChStatusEntry.DocsIf3RxChStatusRcId, "docsIf3RxChStatusRcId")
    docsIf3RxChStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxChStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxChStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxChStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RxChStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RxChStatusEntry.IfIndex})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRcpId", types.YLeaf{"DocsIf3RccStatusRcpId", docsIf3RxChStatusEntry.DocsIf3RccStatusRcpId})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRccStatusId", types.YLeaf{"DocsIf3RccStatusRccStatusId", docsIf3RxChStatusEntry.DocsIf3RccStatusRccStatusId})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RxChStatusRcId", types.YLeaf{"DocsIf3RxChStatusRcId", docsIf3RxChStatusEntry.DocsIf3RxChStatusRcId})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RxChStatusChIfIndex", types.YLeaf{"DocsIf3RxChStatusChIfIndex", docsIf3RxChStatusEntry.DocsIf3RxChStatusChIfIndex})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RxChStatusPrimaryDsIndicator", types.YLeaf{"DocsIf3RxChStatusPrimaryDsIndicator", docsIf3RxChStatusEntry.DocsIf3RxChStatusPrimaryDsIndicator})
    docsIf3RxChStatusEntry.EntityData.Leafs.Append("docsIf3RxChStatusRcRmConnectivityId", types.YLeaf{"DocsIf3RxChStatusRcRmConnectivityId", docsIf3RxChStatusEntry.DocsIf3RxChStatusRcRmConnectivityId})

    docsIf3RxChStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccStatusRcpId", "DocsIf3RccStatusRccStatusId", "DocsIf3RxChStatusRcId"}

    return &(docsIf3RxChStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RxModuleCfgTable
// The Receive Module Configuration object permits
// an operator to configure how CMs with certain Receive
// Channel Profiles (RCPs) will configure the Receive
// Modules within their profile upon CM registration.
//  When a CM registers with an RCP for which all Receive
// Module Indices (RmIds) are configured in this object
// and all Receive Channels are configured within the
// Receive Channel (ReceiveChannel) object, the CMTS
// should use the configuration within these objects to
// set the Receive Channel Configuration assigned to
// the CM in a REG-RSP message.  A CMTS may require configuration
// of all pertinent Receive Module and Receive
// Channel instances (i.e., MIB table entries) in order
// to register a CM that reports a Receive Channel Profile.
//  If the CM reports multiple RCPs, and Receive Module
// and Receive Channel objects have instances (i.e.,
// MIB table entries) for more than one RCP reported by
// the CM, the particular RCP selected by the CMTS is not
// specified.  A CMTS is not restricted to assigning Receive
// Modules based only on the contents of this object.
// 
// This object supports the creation and deletion of multiple
// instances.
// Creation of a new instance of this object requires the
// reference of a valid RccCfg instance.
type DOCSIF3MIB_DocsIf3RxModuleCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RxModuleCfgTable. The ifIndex key corresponds
    // to the MAC Domain interface where the RCC is configured. The type is slice
    // of DOCSIF3MIB_DocsIf3RxModuleCfgTable_DocsIf3RxModuleCfgEntry.
    DocsIf3RxModuleCfgEntry []*DOCSIF3MIB_DocsIf3RxModuleCfgTable_DocsIf3RxModuleCfgEntry
}

func (docsIf3RxModuleCfgTable *DOCSIF3MIB_DocsIf3RxModuleCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3RxModuleCfgTable.EntityData.YFilter = docsIf3RxModuleCfgTable.YFilter
    docsIf3RxModuleCfgTable.EntityData.YangName = "docsIf3RxModuleCfgTable"
    docsIf3RxModuleCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxModuleCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RxModuleCfgTable.EntityData.SegmentPath = "docsIf3RxModuleCfgTable"
    docsIf3RxModuleCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxModuleCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxModuleCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxModuleCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RxModuleCfgTable.EntityData.Children.Append("docsIf3RxModuleCfgEntry", types.YChild{"DocsIf3RxModuleCfgEntry", nil})
    for i := range docsIf3RxModuleCfgTable.DocsIf3RxModuleCfgEntry {
        docsIf3RxModuleCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RxModuleCfgTable.DocsIf3RxModuleCfgEntry[i]), types.YChild{"DocsIf3RxModuleCfgEntry", docsIf3RxModuleCfgTable.DocsIf3RxModuleCfgEntry[i]})
    }
    docsIf3RxModuleCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RxModuleCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3RxModuleCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RxModuleCfgTable_DocsIf3RxModuleCfgEntry
// The conceptual row of docsIf3RxModuleCfgTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured.
type DOCSIF3MIB_DocsIf3RxModuleCfgTable_DocsIf3RxModuleCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with length: 5. Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry_DocsIf3RccCfgRcpId
    DocsIf3RccCfgRcpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccCfgTable_DocsIf3RccCfgEntry_DocsIf3RccCfgRccCfgId
    DocsIf3RccCfgRccCfgId interface{}

    // This attribute is a key. This key represents an identifier of a Receive
    // Module instance within the Receive Channel Profile. The type is interface{}
    // with range: 1..255.
    DocsIf3RxModuleCfgRmId interface{}

    // This attribute represents the higher level (i.e., closer to RF) Receive
    // Module to which this Receive Module connects.  If this object contains a
    // zero value (and thus no Receive Module Connectivity), the Receive Module
    // Connectivity TLV is omitted from the RCC. Within a single instance of the
    // ReceiveModule object, the RmRmConnectivityId attribute cannot contain the
    // same value as the RmId attribute.  The RmRmConnectivityId attribute points
    // to a separate ReceiveModule object instance with the same value of
    // RccCfgId. The type is interface{} with range: 0..255.
    DocsIf3RxModuleCfgRmRmConnectivityId interface{}

    // This attribute represents the center frequency, in Hz, and a multiple of
    // 62500, that indicates the lowest frequency channel of the Receive Module,
    // or 0 if not applicable to the Receive Module. The type is interface{} with
    // range: 0..4294967295. Units are Hz.
    DocsIf3RxModuleCfgFirstCenterFrequency interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3RxModuleCfgRowStatus interface{}
}

func (docsIf3RxModuleCfgEntry *DOCSIF3MIB_DocsIf3RxModuleCfgTable_DocsIf3RxModuleCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RxModuleCfgEntry.EntityData.YFilter = docsIf3RxModuleCfgEntry.YFilter
    docsIf3RxModuleCfgEntry.EntityData.YangName = "docsIf3RxModuleCfgEntry"
    docsIf3RxModuleCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxModuleCfgEntry.EntityData.ParentYangName = "docsIf3RxModuleCfgTable"
    docsIf3RxModuleCfgEntry.EntityData.SegmentPath = "docsIf3RxModuleCfgEntry" + types.AddKeyToken(docsIf3RxModuleCfgEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RxModuleCfgEntry.DocsIf3RccCfgRcpId, "docsIf3RccCfgRcpId") + types.AddKeyToken(docsIf3RxModuleCfgEntry.DocsIf3RccCfgRccCfgId, "docsIf3RccCfgRccCfgId") + types.AddKeyToken(docsIf3RxModuleCfgEntry.DocsIf3RxModuleCfgRmId, "docsIf3RxModuleCfgRmId")
    docsIf3RxModuleCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxModuleCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxModuleCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxModuleCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RxModuleCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RxModuleCfgEntry.IfIndex})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRcpId", types.YLeaf{"DocsIf3RccCfgRcpId", docsIf3RxModuleCfgEntry.DocsIf3RccCfgRcpId})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RccCfgRccCfgId", types.YLeaf{"DocsIf3RccCfgRccCfgId", docsIf3RxModuleCfgEntry.DocsIf3RccCfgRccCfgId})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RxModuleCfgRmId", types.YLeaf{"DocsIf3RxModuleCfgRmId", docsIf3RxModuleCfgEntry.DocsIf3RxModuleCfgRmId})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RxModuleCfgRmRmConnectivityId", types.YLeaf{"DocsIf3RxModuleCfgRmRmConnectivityId", docsIf3RxModuleCfgEntry.DocsIf3RxModuleCfgRmRmConnectivityId})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RxModuleCfgFirstCenterFrequency", types.YLeaf{"DocsIf3RxModuleCfgFirstCenterFrequency", docsIf3RxModuleCfgEntry.DocsIf3RxModuleCfgFirstCenterFrequency})
    docsIf3RxModuleCfgEntry.EntityData.Leafs.Append("docsIf3RxModuleCfgRowStatus", types.YLeaf{"DocsIf3RxModuleCfgRowStatus", docsIf3RxModuleCfgEntry.DocsIf3RxModuleCfgRowStatus})

    docsIf3RxModuleCfgEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccCfgRcpId", "DocsIf3RccCfgRccCfgId", "DocsIf3RxModuleCfgRmId"}

    return &(docsIf3RxModuleCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3RxModuleStatusTable
// The Receive Module Status object provides a read-only
// view of the statically configured and dynamically
// created Receive Modules within an RCC.
type DOCSIF3MIB_DocsIf3RxModuleStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3RxModuleStatusTable. The ifIndex key
    // corresponds to the MAC Domain interface where the RCC is configured. When
    // this object is defined on the CM, the value of RccStatusId is always 1. The
    // type is slice of
    // DOCSIF3MIB_DocsIf3RxModuleStatusTable_DocsIf3RxModuleStatusEntry.
    DocsIf3RxModuleStatusEntry []*DOCSIF3MIB_DocsIf3RxModuleStatusTable_DocsIf3RxModuleStatusEntry
}

func (docsIf3RxModuleStatusTable *DOCSIF3MIB_DocsIf3RxModuleStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3RxModuleStatusTable.EntityData.YFilter = docsIf3RxModuleStatusTable.YFilter
    docsIf3RxModuleStatusTable.EntityData.YangName = "docsIf3RxModuleStatusTable"
    docsIf3RxModuleStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxModuleStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3RxModuleStatusTable.EntityData.SegmentPath = "docsIf3RxModuleStatusTable"
    docsIf3RxModuleStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxModuleStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxModuleStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxModuleStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3RxModuleStatusTable.EntityData.Children.Append("docsIf3RxModuleStatusEntry", types.YChild{"DocsIf3RxModuleStatusEntry", nil})
    for i := range docsIf3RxModuleStatusTable.DocsIf3RxModuleStatusEntry {
        docsIf3RxModuleStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3RxModuleStatusTable.DocsIf3RxModuleStatusEntry[i]), types.YChild{"DocsIf3RxModuleStatusEntry", docsIf3RxModuleStatusTable.DocsIf3RxModuleStatusEntry[i]})
    }
    docsIf3RxModuleStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3RxModuleStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3RxModuleStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3RxModuleStatusTable_DocsIf3RxModuleStatusEntry
// The conceptual row of docsIf3RxModuleStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the RCC is configured. When this object is defined
// on the CM, the value of RccStatusId is always 1.
type DOCSIF3MIB_DocsIf3RxModuleStatusTable_DocsIf3RxModuleStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string with length: 5. Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusRcpId
    DocsIf3RccStatusRcpId interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3RccStatusTable_DocsIf3RccStatusEntry_DocsIf3RccStatusRccStatusId
    DocsIf3RccStatusRccStatusId interface{}

    // This attribute is a key. This key represents an identifier of a Receive
    // Module instance within the Receive Channel Profile. The type is interface{}
    // with range: 1..255.
    DocsIf3RxModuleStatusRmId interface{}

    // This attribute represents the Receive Module to which this Receive Module
    // connects. Requirements for module connectivity are detailed in the
    // RmRmConnectivityId of the RccCfg object. A value of zero indicates that the
    // Receive Module TLV is omitted from the RCC. The type is interface{} with
    // range: 0..255.
    DocsIf3RxModuleStatusRmRmConnectivityId interface{}

    // This attribute represents the low frequency channel of the Receive Module,
    // or 0 if not applicable to the Receive Module. The type is interface{} with
    // range: 0..4294967295. Units are Hz.
    DocsIf3RxModuleStatusFirstCenterFrequency interface{}
}

func (docsIf3RxModuleStatusEntry *DOCSIF3MIB_DocsIf3RxModuleStatusTable_DocsIf3RxModuleStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3RxModuleStatusEntry.EntityData.YFilter = docsIf3RxModuleStatusEntry.YFilter
    docsIf3RxModuleStatusEntry.EntityData.YangName = "docsIf3RxModuleStatusEntry"
    docsIf3RxModuleStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3RxModuleStatusEntry.EntityData.ParentYangName = "docsIf3RxModuleStatusTable"
    docsIf3RxModuleStatusEntry.EntityData.SegmentPath = "docsIf3RxModuleStatusEntry" + types.AddKeyToken(docsIf3RxModuleStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3RxModuleStatusEntry.DocsIf3RccStatusRcpId, "docsIf3RccStatusRcpId") + types.AddKeyToken(docsIf3RxModuleStatusEntry.DocsIf3RccStatusRccStatusId, "docsIf3RccStatusRccStatusId") + types.AddKeyToken(docsIf3RxModuleStatusEntry.DocsIf3RxModuleStatusRmId, "docsIf3RxModuleStatusRmId")
    docsIf3RxModuleStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3RxModuleStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3RxModuleStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3RxModuleStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3RxModuleStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3RxModuleStatusEntry.IfIndex})
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRcpId", types.YLeaf{"DocsIf3RccStatusRcpId", docsIf3RxModuleStatusEntry.DocsIf3RccStatusRcpId})
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("docsIf3RccStatusRccStatusId", types.YLeaf{"DocsIf3RccStatusRccStatusId", docsIf3RxModuleStatusEntry.DocsIf3RccStatusRccStatusId})
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("docsIf3RxModuleStatusRmId", types.YLeaf{"DocsIf3RxModuleStatusRmId", docsIf3RxModuleStatusEntry.DocsIf3RxModuleStatusRmId})
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("docsIf3RxModuleStatusRmRmConnectivityId", types.YLeaf{"DocsIf3RxModuleStatusRmRmConnectivityId", docsIf3RxModuleStatusEntry.DocsIf3RxModuleStatusRmRmConnectivityId})
    docsIf3RxModuleStatusEntry.EntityData.Leafs.Append("docsIf3RxModuleStatusFirstCenterFrequency", types.YLeaf{"DocsIf3RxModuleStatusFirstCenterFrequency", docsIf3RxModuleStatusEntry.DocsIf3RxModuleStatusFirstCenterFrequency})

    docsIf3RxModuleStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3RccStatusRcpId", "DocsIf3RccStatusRccStatusId", "DocsIf3RxModuleStatusRmId"}

    return &(docsIf3RxModuleStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdNodeStatusTable
// This object reports the MD-DS-SG-ID and MD-US-SG-ID
// associated with a MD-CM-SG-ID within a MAC Domain
// and the Fiber Nodes reached by the MD-CM-SG.
type DOCSIF3MIB_DocsIf3MdNodeStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdNodeStatusTable. The ifIndex key corresponds
    // to the MAC Domain interface where the MD-CM-SG-ID is configured. The type
    // is slice of DOCSIF3MIB_DocsIf3MdNodeStatusTable_DocsIf3MdNodeStatusEntry.
    DocsIf3MdNodeStatusEntry []*DOCSIF3MIB_DocsIf3MdNodeStatusTable_DocsIf3MdNodeStatusEntry
}

func (docsIf3MdNodeStatusTable *DOCSIF3MIB_DocsIf3MdNodeStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdNodeStatusTable.EntityData.YFilter = docsIf3MdNodeStatusTable.YFilter
    docsIf3MdNodeStatusTable.EntityData.YangName = "docsIf3MdNodeStatusTable"
    docsIf3MdNodeStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdNodeStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdNodeStatusTable.EntityData.SegmentPath = "docsIf3MdNodeStatusTable"
    docsIf3MdNodeStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdNodeStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdNodeStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdNodeStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdNodeStatusTable.EntityData.Children.Append("docsIf3MdNodeStatusEntry", types.YChild{"DocsIf3MdNodeStatusEntry", nil})
    for i := range docsIf3MdNodeStatusTable.DocsIf3MdNodeStatusEntry {
        docsIf3MdNodeStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdNodeStatusTable.DocsIf3MdNodeStatusEntry[i]), types.YChild{"DocsIf3MdNodeStatusEntry", docsIf3MdNodeStatusTable.DocsIf3MdNodeStatusEntry[i]})
    }
    docsIf3MdNodeStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdNodeStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdNodeStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdNodeStatusTable_DocsIf3MdNodeStatusEntry
// The conceptual row of docsIf3MdNodeStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the MD-CM-SG-ID is configured.
type DOCSIF3MIB_DocsIf3MdNodeStatusTable_DocsIf3MdNodeStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the name of a fiber node
    // associated with a MD-CM-SG of a MAC Domain. The type is string with length:
    // 0..16.
    DocsIf3MdNodeStatusNodeName interface{}

    // This attribute is a key. This attribute is a key and indicates the
    // MD-CM-SG-ID of this instance. A particular MdCmSgId in a MAC Domain is
    // associated with one or more Fiber Nodes. The type is interface{} with
    // range: 1..4294967295.
    DocsIf3MdNodeStatusMdCmSgId interface{}

    // This attribute corresponds to the MD-DS-SG-ID of the MD-CM-SG of this
    // object instance. The MdDsSgId values are unique within a MAC Domain. The
    // type is interface{} with range: 1..255.
    DocsIf3MdNodeStatusMdDsSgId interface{}

    // This attribute corresponds to the MD-US-SG-ID of the MD-CM-SG of this
    // object instance. The MdUsSgId values are unique within a MAC Domain. The
    // type is interface{} with range: 1..255.
    DocsIf3MdNodeStatusMdUsSgId interface{}
}

func (docsIf3MdNodeStatusEntry *DOCSIF3MIB_DocsIf3MdNodeStatusTable_DocsIf3MdNodeStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdNodeStatusEntry.EntityData.YFilter = docsIf3MdNodeStatusEntry.YFilter
    docsIf3MdNodeStatusEntry.EntityData.YangName = "docsIf3MdNodeStatusEntry"
    docsIf3MdNodeStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdNodeStatusEntry.EntityData.ParentYangName = "docsIf3MdNodeStatusTable"
    docsIf3MdNodeStatusEntry.EntityData.SegmentPath = "docsIf3MdNodeStatusEntry" + types.AddKeyToken(docsIf3MdNodeStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusNodeName, "docsIf3MdNodeStatusNodeName") + types.AddKeyToken(docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusMdCmSgId, "docsIf3MdNodeStatusMdCmSgId")
    docsIf3MdNodeStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdNodeStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdNodeStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdNodeStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdNodeStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdNodeStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3MdNodeStatusEntry.IfIndex})
    docsIf3MdNodeStatusEntry.EntityData.Leafs.Append("docsIf3MdNodeStatusNodeName", types.YLeaf{"DocsIf3MdNodeStatusNodeName", docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusNodeName})
    docsIf3MdNodeStatusEntry.EntityData.Leafs.Append("docsIf3MdNodeStatusMdCmSgId", types.YLeaf{"DocsIf3MdNodeStatusMdCmSgId", docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusMdCmSgId})
    docsIf3MdNodeStatusEntry.EntityData.Leafs.Append("docsIf3MdNodeStatusMdDsSgId", types.YLeaf{"DocsIf3MdNodeStatusMdDsSgId", docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusMdDsSgId})
    docsIf3MdNodeStatusEntry.EntityData.Leafs.Append("docsIf3MdNodeStatusMdUsSgId", types.YLeaf{"DocsIf3MdNodeStatusMdUsSgId", docsIf3MdNodeStatusEntry.DocsIf3MdNodeStatusMdUsSgId})

    docsIf3MdNodeStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3MdNodeStatusNodeName", "DocsIf3MdNodeStatusMdCmSgId"}

    return &(docsIf3MdNodeStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdDsSgStatusTable
// This object returns the list of downstream channel
// associated with a MAC Domain MD-DS-SG-ID.
type DOCSIF3MIB_DocsIf3MdDsSgStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdDsSgStatusTable. The ifIndex key corresponds
    // to the MAC Domain interface where the MD-DS-SG-ID is configured. The CMTS
    // is not required to persist instances of this object across
    // reinitializations. The type is slice of
    // DOCSIF3MIB_DocsIf3MdDsSgStatusTable_DocsIf3MdDsSgStatusEntry.
    DocsIf3MdDsSgStatusEntry []*DOCSIF3MIB_DocsIf3MdDsSgStatusTable_DocsIf3MdDsSgStatusEntry
}

func (docsIf3MdDsSgStatusTable *DOCSIF3MIB_DocsIf3MdDsSgStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdDsSgStatusTable.EntityData.YFilter = docsIf3MdDsSgStatusTable.YFilter
    docsIf3MdDsSgStatusTable.EntityData.YangName = "docsIf3MdDsSgStatusTable"
    docsIf3MdDsSgStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdDsSgStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdDsSgStatusTable.EntityData.SegmentPath = "docsIf3MdDsSgStatusTable"
    docsIf3MdDsSgStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdDsSgStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdDsSgStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdDsSgStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdDsSgStatusTable.EntityData.Children.Append("docsIf3MdDsSgStatusEntry", types.YChild{"DocsIf3MdDsSgStatusEntry", nil})
    for i := range docsIf3MdDsSgStatusTable.DocsIf3MdDsSgStatusEntry {
        docsIf3MdDsSgStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdDsSgStatusTable.DocsIf3MdDsSgStatusEntry[i]), types.YChild{"DocsIf3MdDsSgStatusEntry", docsIf3MdDsSgStatusTable.DocsIf3MdDsSgStatusEntry[i]})
    }
    docsIf3MdDsSgStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdDsSgStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdDsSgStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdDsSgStatusTable_DocsIf3MdDsSgStatusEntry
// The conceptual row of docsIf3MdDsSgStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the MD-DS-SG-ID is configured.
// The CMTS is not required to persist instances of this
// object across reinitializations.
type DOCSIF3MIB_DocsIf3MdDsSgStatusTable_DocsIf3MdDsSgStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents a MD-DS-SG-ID in a Mac Domain.
    // The type is interface{} with range: 1..255.
    DocsIf3MdDsSgStatusMdDsSgId interface{}

    // This attribute represents a reference to the list of downstream channels of
    // the MD-DS-SG-ID. The type is interface{} with range: 0..4294967295.
    DocsIf3MdDsSgStatusChSetId interface{}
}

func (docsIf3MdDsSgStatusEntry *DOCSIF3MIB_DocsIf3MdDsSgStatusTable_DocsIf3MdDsSgStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdDsSgStatusEntry.EntityData.YFilter = docsIf3MdDsSgStatusEntry.YFilter
    docsIf3MdDsSgStatusEntry.EntityData.YangName = "docsIf3MdDsSgStatusEntry"
    docsIf3MdDsSgStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdDsSgStatusEntry.EntityData.ParentYangName = "docsIf3MdDsSgStatusTable"
    docsIf3MdDsSgStatusEntry.EntityData.SegmentPath = "docsIf3MdDsSgStatusEntry" + types.AddKeyToken(docsIf3MdDsSgStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3MdDsSgStatusEntry.DocsIf3MdDsSgStatusMdDsSgId, "docsIf3MdDsSgStatusMdDsSgId")
    docsIf3MdDsSgStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdDsSgStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdDsSgStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdDsSgStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdDsSgStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdDsSgStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3MdDsSgStatusEntry.IfIndex})
    docsIf3MdDsSgStatusEntry.EntityData.Leafs.Append("docsIf3MdDsSgStatusMdDsSgId", types.YLeaf{"DocsIf3MdDsSgStatusMdDsSgId", docsIf3MdDsSgStatusEntry.DocsIf3MdDsSgStatusMdDsSgId})
    docsIf3MdDsSgStatusEntry.EntityData.Leafs.Append("docsIf3MdDsSgStatusChSetId", types.YLeaf{"DocsIf3MdDsSgStatusChSetId", docsIf3MdDsSgStatusEntry.DocsIf3MdDsSgStatusChSetId})

    docsIf3MdDsSgStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3MdDsSgStatusMdDsSgId"}

    return &(docsIf3MdDsSgStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdUsSgStatusTable
// This object returns the list of upstream channels
// associated with a MAC Domain MD-US-SG-ID.
type DOCSIF3MIB_DocsIf3MdUsSgStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdUsSgStatusTable. The ifIndex key corresponds
    // to the MAC Domain interface where the MD-DS-SG-ID is configured. The CMTS
    // is not required to persist instances of this object across
    // reinitializations. The type is slice of
    // DOCSIF3MIB_DocsIf3MdUsSgStatusTable_DocsIf3MdUsSgStatusEntry.
    DocsIf3MdUsSgStatusEntry []*DOCSIF3MIB_DocsIf3MdUsSgStatusTable_DocsIf3MdUsSgStatusEntry
}

func (docsIf3MdUsSgStatusTable *DOCSIF3MIB_DocsIf3MdUsSgStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdUsSgStatusTable.EntityData.YFilter = docsIf3MdUsSgStatusTable.YFilter
    docsIf3MdUsSgStatusTable.EntityData.YangName = "docsIf3MdUsSgStatusTable"
    docsIf3MdUsSgStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdUsSgStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdUsSgStatusTable.EntityData.SegmentPath = "docsIf3MdUsSgStatusTable"
    docsIf3MdUsSgStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdUsSgStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdUsSgStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdUsSgStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdUsSgStatusTable.EntityData.Children.Append("docsIf3MdUsSgStatusEntry", types.YChild{"DocsIf3MdUsSgStatusEntry", nil})
    for i := range docsIf3MdUsSgStatusTable.DocsIf3MdUsSgStatusEntry {
        docsIf3MdUsSgStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdUsSgStatusTable.DocsIf3MdUsSgStatusEntry[i]), types.YChild{"DocsIf3MdUsSgStatusEntry", docsIf3MdUsSgStatusTable.DocsIf3MdUsSgStatusEntry[i]})
    }
    docsIf3MdUsSgStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdUsSgStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdUsSgStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdUsSgStatusTable_DocsIf3MdUsSgStatusEntry
// The conceptual row of docsIf3MdUsSgStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the MD-DS-SG-ID is configured.
// The CMTS is not required to persist instances of this
// object across reinitializations.
type DOCSIF3MIB_DocsIf3MdUsSgStatusTable_DocsIf3MdUsSgStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents a MD-US-SG-ID in a Mac Domain.
    // The type is interface{} with range: 1..255.
    DocsIf3MdUsSgStatusMdUsSgId interface{}

    // This attribute represents a reference to the list of upstream channels of
    // the MD-US-SG-ID. The type is interface{} with range: 0..4294967295.
    DocsIf3MdUsSgStatusChSetId interface{}
}

func (docsIf3MdUsSgStatusEntry *DOCSIF3MIB_DocsIf3MdUsSgStatusTable_DocsIf3MdUsSgStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdUsSgStatusEntry.EntityData.YFilter = docsIf3MdUsSgStatusEntry.YFilter
    docsIf3MdUsSgStatusEntry.EntityData.YangName = "docsIf3MdUsSgStatusEntry"
    docsIf3MdUsSgStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdUsSgStatusEntry.EntityData.ParentYangName = "docsIf3MdUsSgStatusTable"
    docsIf3MdUsSgStatusEntry.EntityData.SegmentPath = "docsIf3MdUsSgStatusEntry" + types.AddKeyToken(docsIf3MdUsSgStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3MdUsSgStatusEntry.DocsIf3MdUsSgStatusMdUsSgId, "docsIf3MdUsSgStatusMdUsSgId")
    docsIf3MdUsSgStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdUsSgStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdUsSgStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdUsSgStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdUsSgStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdUsSgStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3MdUsSgStatusEntry.IfIndex})
    docsIf3MdUsSgStatusEntry.EntityData.Leafs.Append("docsIf3MdUsSgStatusMdUsSgId", types.YLeaf{"DocsIf3MdUsSgStatusMdUsSgId", docsIf3MdUsSgStatusEntry.DocsIf3MdUsSgStatusMdUsSgId})
    docsIf3MdUsSgStatusEntry.EntityData.Leafs.Append("docsIf3MdUsSgStatusChSetId", types.YLeaf{"DocsIf3MdUsSgStatusChSetId", docsIf3MdUsSgStatusEntry.DocsIf3MdUsSgStatusChSetId})

    docsIf3MdUsSgStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3MdUsSgStatusMdUsSgId"}

    return &(docsIf3MdUsSgStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable
// This object returns the set of downstream channels
// that carry UCDs and MAPs for a particular upstream channel
// in a MAC Domain.
type DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdUsToDsChMappingTable. The type is slice of
    // DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable_DocsIf3MdUsToDsChMappingEntry.
    DocsIf3MdUsToDsChMappingEntry []*DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable_DocsIf3MdUsToDsChMappingEntry
}

func (docsIf3MdUsToDsChMappingTable *DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdUsToDsChMappingTable.EntityData.YFilter = docsIf3MdUsToDsChMappingTable.YFilter
    docsIf3MdUsToDsChMappingTable.EntityData.YangName = "docsIf3MdUsToDsChMappingTable"
    docsIf3MdUsToDsChMappingTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdUsToDsChMappingTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdUsToDsChMappingTable.EntityData.SegmentPath = "docsIf3MdUsToDsChMappingTable"
    docsIf3MdUsToDsChMappingTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdUsToDsChMappingTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdUsToDsChMappingTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdUsToDsChMappingTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdUsToDsChMappingTable.EntityData.Children.Append("docsIf3MdUsToDsChMappingEntry", types.YChild{"DocsIf3MdUsToDsChMappingEntry", nil})
    for i := range docsIf3MdUsToDsChMappingTable.DocsIf3MdUsToDsChMappingEntry {
        docsIf3MdUsToDsChMappingTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdUsToDsChMappingTable.DocsIf3MdUsToDsChMappingEntry[i]), types.YChild{"DocsIf3MdUsToDsChMappingEntry", docsIf3MdUsToDsChMappingTable.DocsIf3MdUsToDsChMappingEntry[i]})
    }
    docsIf3MdUsToDsChMappingTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdUsToDsChMappingTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdUsToDsChMappingTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable_DocsIf3MdUsToDsChMappingEntry
// The conceptual row of docsIf3MdUsToDsChMappingTable.
type DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable_DocsIf3MdUsToDsChMappingEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This key represents the interface index of the
    // upstream channel to which this instance applies. The type is interface{}
    // with range: 1..2147483647.
    DocsIf3MdUsToDsChMappingUsIfIndex interface{}

    // This attribute is a key. This key represents the interface index of a
    // downstream channel carrying in UCDs and Maps associated with the upstream
    // channel defined by this instance. The type is interface{} with range:
    // 1..2147483647.
    DocsIf3MdUsToDsChMappingDsIfIndex interface{}

    // This attribute is a key. This key represents the MAC domain of the upstream
    // and downstream channels of this instance. The type is interface{} with
    // range: 1..2147483647.
    DocsIf3MdUsToDsChMappingMdIfIndex interface{}
}

func (docsIf3MdUsToDsChMappingEntry *DOCSIF3MIB_DocsIf3MdUsToDsChMappingTable_DocsIf3MdUsToDsChMappingEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdUsToDsChMappingEntry.EntityData.YFilter = docsIf3MdUsToDsChMappingEntry.YFilter
    docsIf3MdUsToDsChMappingEntry.EntityData.YangName = "docsIf3MdUsToDsChMappingEntry"
    docsIf3MdUsToDsChMappingEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdUsToDsChMappingEntry.EntityData.ParentYangName = "docsIf3MdUsToDsChMappingTable"
    docsIf3MdUsToDsChMappingEntry.EntityData.SegmentPath = "docsIf3MdUsToDsChMappingEntry" + types.AddKeyToken(docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingUsIfIndex, "docsIf3MdUsToDsChMappingUsIfIndex") + types.AddKeyToken(docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingDsIfIndex, "docsIf3MdUsToDsChMappingDsIfIndex") + types.AddKeyToken(docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingMdIfIndex, "docsIf3MdUsToDsChMappingMdIfIndex")
    docsIf3MdUsToDsChMappingEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdUsToDsChMappingEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdUsToDsChMappingEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdUsToDsChMappingEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdUsToDsChMappingEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdUsToDsChMappingEntry.EntityData.Leafs.Append("docsIf3MdUsToDsChMappingUsIfIndex", types.YLeaf{"DocsIf3MdUsToDsChMappingUsIfIndex", docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingUsIfIndex})
    docsIf3MdUsToDsChMappingEntry.EntityData.Leafs.Append("docsIf3MdUsToDsChMappingDsIfIndex", types.YLeaf{"DocsIf3MdUsToDsChMappingDsIfIndex", docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingDsIfIndex})
    docsIf3MdUsToDsChMappingEntry.EntityData.Leafs.Append("docsIf3MdUsToDsChMappingMdIfIndex", types.YLeaf{"DocsIf3MdUsToDsChMappingMdIfIndex", docsIf3MdUsToDsChMappingEntry.DocsIf3MdUsToDsChMappingMdIfIndex})

    docsIf3MdUsToDsChMappingEntry.EntityData.YListKeys = []string {"DocsIf3MdUsToDsChMappingUsIfIndex", "DocsIf3MdUsToDsChMappingDsIfIndex", "DocsIf3MdUsToDsChMappingMdIfIndex"}

    return &(docsIf3MdUsToDsChMappingEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdCfgTable
// This object contains MAC domain level control and
// configuration attributes.
type DOCSIF3MIB_DocsIf3MdCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3MdCfgTable. The CMTS persists all instances of
    // MdCfg across reinitializations. The ifIndex key corresponds to the MAC
    // Domain interface. The type is slice of
    // DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry.
    DocsIf3MdCfgEntry []*DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry
}

func (docsIf3MdCfgTable *DOCSIF3MIB_DocsIf3MdCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3MdCfgTable.EntityData.YFilter = docsIf3MdCfgTable.YFilter
    docsIf3MdCfgTable.EntityData.YangName = "docsIf3MdCfgTable"
    docsIf3MdCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3MdCfgTable.EntityData.SegmentPath = "docsIf3MdCfgTable"
    docsIf3MdCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3MdCfgTable.EntityData.Children.Append("docsIf3MdCfgEntry", types.YChild{"DocsIf3MdCfgEntry", nil})
    for i := range docsIf3MdCfgTable.DocsIf3MdCfgEntry {
        docsIf3MdCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3MdCfgTable.DocsIf3MdCfgEntry[i]), types.YChild{"DocsIf3MdCfgEntry", docsIf3MdCfgTable.DocsIf3MdCfgEntry[i]})
    }
    docsIf3MdCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3MdCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3MdCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry
// The conceptual row of docsIf3MdCfgTable.
// The CMTS persists all instances of MdCfg across
// reinitializations.
// The ifIndex key corresponds to the MAC Domain interface.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute configures the interval for the insertion of MDD messages in
    // each downstream channel of a MAC Domain. The type is interface{} with
    // range: 1..2000. Units are milliseconds.
    DocsIf3MdCfgMddInterval interface{}

    // This attribute configures the CMTS IP provisioning mode for a MAC Domain. 
    // When this attribute is set to 'ipv4Only' the CM will acquire a single IPv4
    // address for the CM management stack. When this attribute is set to
    // 'ipv6Only' the CM will acquire a single IPv6 address for the CM management
    // stack. When this attribute is set to 'alternate' the CM will acquire a
    // single IPv6 address for the CM management stack and, if failures occur, the
    // CM will fall back to provision and operation with an IPv4 address. When
    // this attribute is set to 'dualStack' the CM will acquire both an IPv6 and
    // IPv4 address for provisioning and operation. The type is
    // DocsIf3MdCfgIpProvMode.
    DocsIf3MdCfgIpProvMode interface{}

    // If set to 'true', this attribute enables the signaling of the CM-Status
    // Event reporting mechanism. The type is bool.
    DocsIf3MdCfgCmStatusEvCtlEnabled interface{}

    // This attribute indicates in MDD messages the upstream frequency upper band
    // edge of an upstream Channel.  A value 'standard' means Standard Frequency
    // Range and a value 'extended' means Extended Frequency Range. The type is
    // DocsIf3MdCfgUsFreqRange.
    DocsIf3MdCfgUsFreqRange interface{}

    // If set to 'true', this attribute enables the CMTS to use IP Multicast DSID
    // Forwarding (MDF) for the MAC domain. The type is bool.
    DocsIf3MdCfgMcastDsidFwdEnabled interface{}

    // If set to 'true', this attribute enables Downstream Channel Bonding for the
    // MAC Domain. The type is bool.
    DocsIf3MdCfgMultRxChModeEnabled interface{}

    // If set to 'true', this attribute enables Multiple Transmit Channel (MTC)
    // Mode for the MAC Domain. The type is bool.
    DocsIf3MdCfgMultTxChModeEnabled interface{}

    // This attribute enables or disables early authentication and encryption
    // (EAE) signaling for the MAC Domain. It also defines the type of EAE
    // enforcement in the case that EAE is enabled. If set to 'disableEAE', EAE is
    // disabled for the MAC Domain.  If set to 'enableEaeRangingBasedEnforcement',
    // 'enableEaeCapabilityBasedEnforcement' or 'enableEaeTotalEnforcement', EAE
    // is enabled for the MAC Domain.  The following EAE enforcement methods are
    // defined in the case where EAE signaling is enabled: The option
    // 'enableEaeRangingBasedEnforcement' indicates EAE is enforced on CMs that
    // perform ranging with a B-INIT-RNG-REQ message. The option
    // 'enableEaeCapabilityBasedEnforcement' indicates EAE is enforced on CMs that
    // perform ranging with a B-INIT-RNG-REQ message in which the EAE capability
    // flag is set. The option 'enableEaeTotalEnforcement' indicates EAE is
    // enforced on all CMs regardless of their EAE capabilities. The type is
    // DocsIf3MdCfgEarlyAuthEncrCtrl.
    DocsIf3MdCfgEarlyAuthEncrCtrl interface{}

    // If set to 'true', this attribute enables TFTP Proxy functionality for the
    // MAC Domain. The type is bool.
    DocsIf3MdCfgTftpProxyEnabled interface{}

    // If set to 'true', this attribute enables Source Address Verification (SAV)
    // functionality for the MAC Domain. The type is bool.
    DocsIf3MdCfgSrcAddrVerifEnabled interface{}

    // This attribute defines the ITU-J-83 Annex being used  for this MAC Domain. 
    // The value of this attribute  indicates the conformance of the
    // implementation to  important regional cable standards.   Valid enumerations
    // for the attribute are: unknown other annexA : Annex A from ITU-J83 is used.
    // annexB : Annex B from ITU-J83 is used. annexC : Annex C from ITU-J83 is
    // used.  Values 6-255 are reserved. The type is DocsIf3MdCfgDownChannelAnnex.
    DocsIf3MdCfgDownChannelAnnex interface{}

    // If set to 'true', this attribute instructs the CMTS MAC  Domain to enable
    // Upstream Drop Classifiers (UDC) for the  CMs attempting registration in
    // this MAC Domain. The type is bool.
    DocsIf3MdCfgCmUdcEnabled interface{}

    // If set to 'true' and when the CM signals to the CMTS 'Upstream Drop
    // Classifier Group ID' encodings, this attribute instructs the CMTS MAC
    // Domain to send the Subscriber  Management Filters rules associated with the
    // 'Upstream Drop Classifier Group ID' encodings to the CM in the form of UDCs
    // when the following conditions occurs: - The attribute CmUdcEnabled value
    // for this MAC Domain    is set to 'true', and  - The CM has the UDC
    // capability advertised as supported.  If there is no a single Subscriber
    // Management Filter  configured in the CMTS for the CM's signaled UDC Group
    // ID, the CMTS does not send UDC encodings to the CM.  It is vendor specific
    // whether the CMTS maintains enforcement of the CM signaled or default
    // Subscriber Management Filter  groups in the upstream direction. The type is
    // bool.
    DocsIf3MdCfgSendUdcRulesEnabled interface{}

    // This attribute indicates the list of Service Type IDs  associated with the
    // MAC Domain.  During the CM registration process the CMTS will attempt to
    // redirect the CM to a MAC Domain where the CM' Service Type TLV is contained
    // in this attribute. The type is string.
    DocsIf3MdCfgServiceTypeIdList interface{}

    // This attribute indicates the level of BPI+ enforcement policies with the
    // MAC Domain.  The following BPI+ enforcement policies are defined in the
    // case where BPI+ is enabled:  The option 'disable' indicates the CMTS does
    // not enforce BPI+.  The option
    // 'qosCfgFileWithBpi2AndCapabPrivSupportEnabled'  indicates the CMTS enforces
    // BPI+ on CMs that register with a  DOCSIS 1.1 style configuration file with
    // parameters indicating BPI+ is enabled (missing TLV 29 or containing TLV 29
    // set to  enable) and with a Modem Capabilities Privacy Support TLV (5.6) 
    // set to BPI+ support.  The option 'qosCfgFileWithBpi2Enabled' indicates the
    // CMTS  enforces BPI+ on CMs that register with a DOCSIS 1.1 style 
    // configuration file with parameters indicating BPI+ is enabled  (missing TLV
    // 29 or containing TLV 29 set to enable).  The option 'qosCfgFile' indicates
    // the CMTS enforces BPI+ on CMs that register with a DOCSIS 1.1 style
    // configuration file.  The option 'total' indicates the CMTS enforces BPI+ on
    // all CMs. The type is DocsIf3MdCfgBpi2EnforceCtrl.
    DocsIf3MdCfgBpi2EnforceCtrl interface{}

    // This attribute indicates whether the CMTS is configured for  1x1 Energy
    // Management Mode of operation on a per MAC Domain  basis. If this attribute
    // is set to 'true', the CMTS is  configured for 1x1 Energy Management Mode of
    // operation on this MAC Domain. If this attribute is set to 'false', the
    // Energy  Management 1x1 Mode of operation is disabled in the CMTS on  this
    // MAC Domain. The type is interface{} with range: 0..1.
    DocsIf3MdCfgEnergyMgt1x1Enabled interface{}
}

func (docsIf3MdCfgEntry *DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3MdCfgEntry.EntityData.YFilter = docsIf3MdCfgEntry.YFilter
    docsIf3MdCfgEntry.EntityData.YangName = "docsIf3MdCfgEntry"
    docsIf3MdCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3MdCfgEntry.EntityData.ParentYangName = "docsIf3MdCfgTable"
    docsIf3MdCfgEntry.EntityData.SegmentPath = "docsIf3MdCfgEntry" + types.AddKeyToken(docsIf3MdCfgEntry.IfIndex, "ifIndex")
    docsIf3MdCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3MdCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3MdCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3MdCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3MdCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3MdCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3MdCfgEntry.IfIndex})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgMddInterval", types.YLeaf{"DocsIf3MdCfgMddInterval", docsIf3MdCfgEntry.DocsIf3MdCfgMddInterval})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgIpProvMode", types.YLeaf{"DocsIf3MdCfgIpProvMode", docsIf3MdCfgEntry.DocsIf3MdCfgIpProvMode})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgCmStatusEvCtlEnabled", types.YLeaf{"DocsIf3MdCfgCmStatusEvCtlEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgCmStatusEvCtlEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgUsFreqRange", types.YLeaf{"DocsIf3MdCfgUsFreqRange", docsIf3MdCfgEntry.DocsIf3MdCfgUsFreqRange})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgMcastDsidFwdEnabled", types.YLeaf{"DocsIf3MdCfgMcastDsidFwdEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgMcastDsidFwdEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgMultRxChModeEnabled", types.YLeaf{"DocsIf3MdCfgMultRxChModeEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgMultRxChModeEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgMultTxChModeEnabled", types.YLeaf{"DocsIf3MdCfgMultTxChModeEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgMultTxChModeEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgEarlyAuthEncrCtrl", types.YLeaf{"DocsIf3MdCfgEarlyAuthEncrCtrl", docsIf3MdCfgEntry.DocsIf3MdCfgEarlyAuthEncrCtrl})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgTftpProxyEnabled", types.YLeaf{"DocsIf3MdCfgTftpProxyEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgTftpProxyEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgSrcAddrVerifEnabled", types.YLeaf{"DocsIf3MdCfgSrcAddrVerifEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgSrcAddrVerifEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgDownChannelAnnex", types.YLeaf{"DocsIf3MdCfgDownChannelAnnex", docsIf3MdCfgEntry.DocsIf3MdCfgDownChannelAnnex})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgCmUdcEnabled", types.YLeaf{"DocsIf3MdCfgCmUdcEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgCmUdcEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgSendUdcRulesEnabled", types.YLeaf{"DocsIf3MdCfgSendUdcRulesEnabled", docsIf3MdCfgEntry.DocsIf3MdCfgSendUdcRulesEnabled})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgServiceTypeIdList", types.YLeaf{"DocsIf3MdCfgServiceTypeIdList", docsIf3MdCfgEntry.DocsIf3MdCfgServiceTypeIdList})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgBpi2EnforceCtrl", types.YLeaf{"DocsIf3MdCfgBpi2EnforceCtrl", docsIf3MdCfgEntry.DocsIf3MdCfgBpi2EnforceCtrl})
    docsIf3MdCfgEntry.EntityData.Leafs.Append("docsIf3MdCfgEnergyMgt1x1Enabled", types.YLeaf{"DocsIf3MdCfgEnergyMgt1x1Enabled", docsIf3MdCfgEntry.DocsIf3MdCfgEnergyMgt1x1Enabled})

    docsIf3MdCfgEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3MdCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl represents The option 'total' indicates the CMTS enforces BPI+ on all CMs.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl string

const (
    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl_disable DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl = "disable"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl_qosCfgFileWithBpi2AndCapabPrivSupportEnabled DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl = "qosCfgFileWithBpi2AndCapabPrivSupportEnabled"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl_qosCfgFileWithBpi2Enabled DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl = "qosCfgFileWithBpi2Enabled"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl_qosCfgFile DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl = "qosCfgFile"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl_total DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgBpi2EnforceCtrl = "total"
)

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex represents Values 6-255 are reserved.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex string

const (
    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex_unknown DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex = "unknown"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex_other DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex = "other"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex_annexA DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex = "annexA"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex_annexB DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex = "annexB"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex_annexC DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgDownChannelAnnex = "annexC"
)

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl represents capabilities.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl string

const (
    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl_disableEae DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl = "disableEae"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl_enableEaeRangingBasedEnforcement DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl = "enableEaeRangingBasedEnforcement"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl_enableEaeCapabilityBasedEnforcement DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl = "enableEaeCapabilityBasedEnforcement"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl_enableEaeTotalEnforcement DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgEarlyAuthEncrCtrl = "enableEaeTotalEnforcement"
)

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode represents an IPv6 and IPv4 address for provisioning and operation.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode string

const (
    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode_ipv4Only DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode = "ipv4Only"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode_ipv6Only DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode = "ipv6Only"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode_alternate DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode = "alternate"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode_dualStack DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgIpProvMode = "dualStack"
)

// DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange represents a value 'extended' means Extended Frequency Range.
type DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange string

const (
    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange_standard DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange = "standard"

    DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange_extended DOCSIF3MIB_DocsIf3MdCfgTable_DocsIf3MdCfgEntry_DocsIf3MdCfgUsFreqRange = "extended"
)

// DOCSIF3MIB_DocsIf3BondingGrpCfgTable
// This object defines statically configured Downstream
// Bonding Groups and Upstream Bonding Groups on
// the CMTS.
// This object supports the creation and deletion of multiple
// instances.
// Creation of a new instance of this object requires the
// ChList attribute to be set.
type DOCSIF3MIB_DocsIf3BondingGrpCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3BondingGrpCfgTable. The ifIndex key
    // corresponds to the MAC Domain interface where the Bonding Group is
    // configured. The CMTS persists all instances of BondingGrpCfg across
    // reinitializations. The type is slice of
    // DOCSIF3MIB_DocsIf3BondingGrpCfgTable_DocsIf3BondingGrpCfgEntry.
    DocsIf3BondingGrpCfgEntry []*DOCSIF3MIB_DocsIf3BondingGrpCfgTable_DocsIf3BondingGrpCfgEntry
}

func (docsIf3BondingGrpCfgTable *DOCSIF3MIB_DocsIf3BondingGrpCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3BondingGrpCfgTable.EntityData.YFilter = docsIf3BondingGrpCfgTable.YFilter
    docsIf3BondingGrpCfgTable.EntityData.YangName = "docsIf3BondingGrpCfgTable"
    docsIf3BondingGrpCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3BondingGrpCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3BondingGrpCfgTable.EntityData.SegmentPath = "docsIf3BondingGrpCfgTable"
    docsIf3BondingGrpCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3BondingGrpCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3BondingGrpCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3BondingGrpCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3BondingGrpCfgTable.EntityData.Children.Append("docsIf3BondingGrpCfgEntry", types.YChild{"DocsIf3BondingGrpCfgEntry", nil})
    for i := range docsIf3BondingGrpCfgTable.DocsIf3BondingGrpCfgEntry {
        docsIf3BondingGrpCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3BondingGrpCfgTable.DocsIf3BondingGrpCfgEntry[i]), types.YChild{"DocsIf3BondingGrpCfgEntry", docsIf3BondingGrpCfgTable.DocsIf3BondingGrpCfgEntry[i]})
    }
    docsIf3BondingGrpCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3BondingGrpCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3BondingGrpCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3BondingGrpCfgTable_DocsIf3BondingGrpCfgEntry
// The conceptual row of docsIf3BondingGrpCfgTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the Bonding Group is configured.
// The CMTS persists all instances of BondingGrpCfg
// across reinitializations.
type DOCSIF3MIB_DocsIf3BondingGrpCfgTable_DocsIf3BondingGrpCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This attribute defines the ordered list of
    // channels that comprise the channel set. The type is IfDirection.
    DocsIf3BondingGrpCfgDir interface{}

    // This attribute is a key. This key represents the configured bonding group
    // identifier in the indicated direction for the MAC Domain. This attribute is
    // used for the sole purpose of tracking bonding groups defined by management
    // systems. The type is interface{} with range: 1..65535.
    DocsIf3BondingGrpCfgCfgId interface{}

    // This attribute contains the list of channels of the bonding group. The type
    // is string with length: 2..255.
    DocsIf3BondingGrpCfgChList interface{}

    // This attribute represents the Provisioned Attribute Mask encoding for the
    // bonding group. The type is string with length: 0..16.
    DocsIf3BondingGrpCfgSfProvAttrMask interface{}

    // For a Downstream Bonding Group, this attribute provides the DSID
    // Resequencing Wait Time that is to be used for all DSIDs associated with
    // this Downstream Bonding Group.  The value of 255 indicates that the DSID
    // Resequencing Wait Time is determined by the CMTS. The value zero in not
    // supported for downstream  bonding groups. For an Upstream Bonding Group,
    // this attribute has no meaning and returns the value 0. The type is
    // interface{} with range: 0..180 | 255..None. Units are hundredMicroseconds.
    DocsIf3BondingGrpCfgDsidReseqWaitTime interface{}

    // For a Downstream Bonding Group, this attribute provides the DSID
    // Resequencing Warning Threshold that is to be used for all DSIDs associated
    // with this Downstream Bonding Group. The value of 255 indicates that the
    // DSID Resequencing Warning Threshold is determined by the CMTS.  The value
    // of 0 indicates that the threshold warnings are disabled.  When  the value
    // of DsidReseqWaitTime is not equal to 0 or 255, the CMTS must ensure that
    // the value of this object is either  255 or less than the value of
    // DsidReseqWaitTime. The type is interface{} with range: 0..179 | 255..None.
    // Units are hundredMicroseconds.
    DocsIf3BondingGrpCfgDsidReseqWarnThrshld interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3BondingGrpCfgRowStatus interface{}
}

func (docsIf3BondingGrpCfgEntry *DOCSIF3MIB_DocsIf3BondingGrpCfgTable_DocsIf3BondingGrpCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3BondingGrpCfgEntry.EntityData.YFilter = docsIf3BondingGrpCfgEntry.YFilter
    docsIf3BondingGrpCfgEntry.EntityData.YangName = "docsIf3BondingGrpCfgEntry"
    docsIf3BondingGrpCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3BondingGrpCfgEntry.EntityData.ParentYangName = "docsIf3BondingGrpCfgTable"
    docsIf3BondingGrpCfgEntry.EntityData.SegmentPath = "docsIf3BondingGrpCfgEntry" + types.AddKeyToken(docsIf3BondingGrpCfgEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgDir, "docsIf3BondingGrpCfgDir") + types.AddKeyToken(docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgCfgId, "docsIf3BondingGrpCfgCfgId")
    docsIf3BondingGrpCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3BondingGrpCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3BondingGrpCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3BondingGrpCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3BondingGrpCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3BondingGrpCfgEntry.IfIndex})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgDir", types.YLeaf{"DocsIf3BondingGrpCfgDir", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgDir})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgCfgId", types.YLeaf{"DocsIf3BondingGrpCfgCfgId", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgCfgId})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgChList", types.YLeaf{"DocsIf3BondingGrpCfgChList", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgChList})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgSfProvAttrMask", types.YLeaf{"DocsIf3BondingGrpCfgSfProvAttrMask", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgSfProvAttrMask})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgDsidReseqWaitTime", types.YLeaf{"DocsIf3BondingGrpCfgDsidReseqWaitTime", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgDsidReseqWaitTime})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgDsidReseqWarnThrshld", types.YLeaf{"DocsIf3BondingGrpCfgDsidReseqWarnThrshld", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgDsidReseqWarnThrshld})
    docsIf3BondingGrpCfgEntry.EntityData.Leafs.Append("docsIf3BondingGrpCfgRowStatus", types.YLeaf{"DocsIf3BondingGrpCfgRowStatus", docsIf3BondingGrpCfgEntry.DocsIf3BondingGrpCfgRowStatus})

    docsIf3BondingGrpCfgEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3BondingGrpCfgDir", "DocsIf3BondingGrpCfgCfgId"}

    return &(docsIf3BondingGrpCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable
// This object returns administratively-configured
// and CMTS defined downstream bonding groups.
type DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3DsBondingGrpStatusTable. The ifIndex key
    // corresponds to the MAC Domain interface where the Bonding Group is
    // configured. The type is slice of
    // DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable_DocsIf3DsBondingGrpStatusEntry.
    DocsIf3DsBondingGrpStatusEntry []*DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable_DocsIf3DsBondingGrpStatusEntry
}

func (docsIf3DsBondingGrpStatusTable *DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3DsBondingGrpStatusTable.EntityData.YFilter = docsIf3DsBondingGrpStatusTable.YFilter
    docsIf3DsBondingGrpStatusTable.EntityData.YangName = "docsIf3DsBondingGrpStatusTable"
    docsIf3DsBondingGrpStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3DsBondingGrpStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3DsBondingGrpStatusTable.EntityData.SegmentPath = "docsIf3DsBondingGrpStatusTable"
    docsIf3DsBondingGrpStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3DsBondingGrpStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3DsBondingGrpStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3DsBondingGrpStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3DsBondingGrpStatusTable.EntityData.Children.Append("docsIf3DsBondingGrpStatusEntry", types.YChild{"DocsIf3DsBondingGrpStatusEntry", nil})
    for i := range docsIf3DsBondingGrpStatusTable.DocsIf3DsBondingGrpStatusEntry {
        docsIf3DsBondingGrpStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3DsBondingGrpStatusTable.DocsIf3DsBondingGrpStatusEntry[i]), types.YChild{"DocsIf3DsBondingGrpStatusEntry", docsIf3DsBondingGrpStatusTable.DocsIf3DsBondingGrpStatusEntry[i]})
    }
    docsIf3DsBondingGrpStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3DsBondingGrpStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3DsBondingGrpStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable_DocsIf3DsBondingGrpStatusEntry
// The conceptual row of docsIf3DsBondingGrpStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the Bonding Group is configured.
type DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable_DocsIf3DsBondingGrpStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the identifier for the
    // Downstream Bonding Group or the single-downstream channel of this instance.
    // The type is interface{} with range: 0..4294967295.
    DocsIf3DsBondingGrpStatusChSetId interface{}

    // This attribute corresponds to the MD-DS-SG-ID that includes all the
    // downstream channels of the Downstream Bonding Group. The value zero
    // indicates that the bonding group does not contain channels from a single
    // MD-DS-SG and therefore the bonding group is not valid and usable. The type
    // is interface{} with range: 0..255.
    DocsIf3DsBondingGrpStatusMdDsSgId interface{}

    // This attribute provides the BondingGrpCfgId for the downstream bonding
    // group if it was configured. Otherwise, the zero value indicates that the
    // CMTS will define the bonding group. The type is interface{} with range:
    // 0..65535.
    DocsIf3DsBondingGrpStatusCfgId interface{}
}

func (docsIf3DsBondingGrpStatusEntry *DOCSIF3MIB_DocsIf3DsBondingGrpStatusTable_DocsIf3DsBondingGrpStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3DsBondingGrpStatusEntry.EntityData.YFilter = docsIf3DsBondingGrpStatusEntry.YFilter
    docsIf3DsBondingGrpStatusEntry.EntityData.YangName = "docsIf3DsBondingGrpStatusEntry"
    docsIf3DsBondingGrpStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3DsBondingGrpStatusEntry.EntityData.ParentYangName = "docsIf3DsBondingGrpStatusTable"
    docsIf3DsBondingGrpStatusEntry.EntityData.SegmentPath = "docsIf3DsBondingGrpStatusEntry" + types.AddKeyToken(docsIf3DsBondingGrpStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3DsBondingGrpStatusEntry.DocsIf3DsBondingGrpStatusChSetId, "docsIf3DsBondingGrpStatusChSetId")
    docsIf3DsBondingGrpStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3DsBondingGrpStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3DsBondingGrpStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3DsBondingGrpStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3DsBondingGrpStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3DsBondingGrpStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3DsBondingGrpStatusEntry.IfIndex})
    docsIf3DsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3DsBondingGrpStatusChSetId", types.YLeaf{"DocsIf3DsBondingGrpStatusChSetId", docsIf3DsBondingGrpStatusEntry.DocsIf3DsBondingGrpStatusChSetId})
    docsIf3DsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3DsBondingGrpStatusMdDsSgId", types.YLeaf{"DocsIf3DsBondingGrpStatusMdDsSgId", docsIf3DsBondingGrpStatusEntry.DocsIf3DsBondingGrpStatusMdDsSgId})
    docsIf3DsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3DsBondingGrpStatusCfgId", types.YLeaf{"DocsIf3DsBondingGrpStatusCfgId", docsIf3DsBondingGrpStatusEntry.DocsIf3DsBondingGrpStatusCfgId})

    docsIf3DsBondingGrpStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3DsBondingGrpStatusChSetId"}

    return &(docsIf3DsBondingGrpStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable
// This object returns administratively-configured
// and CMTS-defined upstream bonding groups.
type DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3UsBondingGrpStatusTable. The ifIndex key
    // corresponds to the MAC Domain interface where the Bonding Group is
    // configured. The type is slice of
    // DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable_DocsIf3UsBondingGrpStatusEntry.
    DocsIf3UsBondingGrpStatusEntry []*DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable_DocsIf3UsBondingGrpStatusEntry
}

func (docsIf3UsBondingGrpStatusTable *DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable) GetEntityData() *types.CommonEntityData {
    docsIf3UsBondingGrpStatusTable.EntityData.YFilter = docsIf3UsBondingGrpStatusTable.YFilter
    docsIf3UsBondingGrpStatusTable.EntityData.YangName = "docsIf3UsBondingGrpStatusTable"
    docsIf3UsBondingGrpStatusTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsBondingGrpStatusTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3UsBondingGrpStatusTable.EntityData.SegmentPath = "docsIf3UsBondingGrpStatusTable"
    docsIf3UsBondingGrpStatusTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsBondingGrpStatusTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsBondingGrpStatusTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsBondingGrpStatusTable.EntityData.Children = types.NewOrderedMap()
    docsIf3UsBondingGrpStatusTable.EntityData.Children.Append("docsIf3UsBondingGrpStatusEntry", types.YChild{"DocsIf3UsBondingGrpStatusEntry", nil})
    for i := range docsIf3UsBondingGrpStatusTable.DocsIf3UsBondingGrpStatusEntry {
        docsIf3UsBondingGrpStatusTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3UsBondingGrpStatusTable.DocsIf3UsBondingGrpStatusEntry[i]), types.YChild{"DocsIf3UsBondingGrpStatusEntry", docsIf3UsBondingGrpStatusTable.DocsIf3UsBondingGrpStatusEntry[i]})
    }
    docsIf3UsBondingGrpStatusTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3UsBondingGrpStatusTable.EntityData.YListKeys = []string {}

    return &(docsIf3UsBondingGrpStatusTable.EntityData)
}

// DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable_DocsIf3UsBondingGrpStatusEntry
// The conceptual row of docsIf3UsBondingGrpStatusTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the Bonding Group is configured.
type DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable_DocsIf3UsBondingGrpStatusEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the identifier for the
    // Upstream Bonding Group or the single-upstream channel of this instance. The
    // type is interface{} with range: 0..4294967295.
    DocsIf3UsBondingGrpStatusChSetId interface{}

    // This attribute corresponds to the MD-US-SG-ID that includes all the
    // upstream channels of the Upstream Bonding Group. The value zero indicates
    // that the bonding group does not contain channels from a single MD-US-SG and
    // therefore the bonding group is not valid and usable. The type is
    // interface{} with range: 0..255.
    DocsIf3UsBondingGrpStatusMdUsSgId interface{}

    // This attribute provides the BondingGrpCfgId for the upstream bonding group
    // if it was configured. Otherwise, the zero value indicates that the CMTS
    // defines the bonding group. The type is interface{} with range:
    // 0..4294967295.
    DocsIf3UsBondingGrpStatusCfgId interface{}
}

func (docsIf3UsBondingGrpStatusEntry *DOCSIF3MIB_DocsIf3UsBondingGrpStatusTable_DocsIf3UsBondingGrpStatusEntry) GetEntityData() *types.CommonEntityData {
    docsIf3UsBondingGrpStatusEntry.EntityData.YFilter = docsIf3UsBondingGrpStatusEntry.YFilter
    docsIf3UsBondingGrpStatusEntry.EntityData.YangName = "docsIf3UsBondingGrpStatusEntry"
    docsIf3UsBondingGrpStatusEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsBondingGrpStatusEntry.EntityData.ParentYangName = "docsIf3UsBondingGrpStatusTable"
    docsIf3UsBondingGrpStatusEntry.EntityData.SegmentPath = "docsIf3UsBondingGrpStatusEntry" + types.AddKeyToken(docsIf3UsBondingGrpStatusEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3UsBondingGrpStatusEntry.DocsIf3UsBondingGrpStatusChSetId, "docsIf3UsBondingGrpStatusChSetId")
    docsIf3UsBondingGrpStatusEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsBondingGrpStatusEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsBondingGrpStatusEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsBondingGrpStatusEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3UsBondingGrpStatusEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3UsBondingGrpStatusEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3UsBondingGrpStatusEntry.IfIndex})
    docsIf3UsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3UsBondingGrpStatusChSetId", types.YLeaf{"DocsIf3UsBondingGrpStatusChSetId", docsIf3UsBondingGrpStatusEntry.DocsIf3UsBondingGrpStatusChSetId})
    docsIf3UsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3UsBondingGrpStatusMdUsSgId", types.YLeaf{"DocsIf3UsBondingGrpStatusMdUsSgId", docsIf3UsBondingGrpStatusEntry.DocsIf3UsBondingGrpStatusMdUsSgId})
    docsIf3UsBondingGrpStatusEntry.EntityData.Leafs.Append("docsIf3UsBondingGrpStatusCfgId", types.YLeaf{"DocsIf3UsBondingGrpStatusCfgId", docsIf3UsBondingGrpStatusEntry.DocsIf3UsBondingGrpStatusCfgId})

    docsIf3UsBondingGrpStatusEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3UsBondingGrpStatusChSetId"}

    return &(docsIf3UsBondingGrpStatusEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3UsChExtTable
// This object defines management extensions for upstream
// channels, in particular SCDMA parameters.
type DOCSIF3MIB_DocsIf3UsChExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3UsChExtTable. The ifIndex key corresponds to
    // each of the upstream channels. The type is slice of
    // DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry.
    DocsIf3UsChExtEntry []*DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry
}

func (docsIf3UsChExtTable *DOCSIF3MIB_DocsIf3UsChExtTable) GetEntityData() *types.CommonEntityData {
    docsIf3UsChExtTable.EntityData.YFilter = docsIf3UsChExtTable.YFilter
    docsIf3UsChExtTable.EntityData.YangName = "docsIf3UsChExtTable"
    docsIf3UsChExtTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsChExtTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3UsChExtTable.EntityData.SegmentPath = "docsIf3UsChExtTable"
    docsIf3UsChExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsChExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsChExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsChExtTable.EntityData.Children = types.NewOrderedMap()
    docsIf3UsChExtTable.EntityData.Children.Append("docsIf3UsChExtEntry", types.YChild{"DocsIf3UsChExtEntry", nil})
    for i := range docsIf3UsChExtTable.DocsIf3UsChExtEntry {
        docsIf3UsChExtTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3UsChExtTable.DocsIf3UsChExtEntry[i]), types.YChild{"DocsIf3UsChExtEntry", docsIf3UsChExtTable.DocsIf3UsChExtEntry[i]})
    }
    docsIf3UsChExtTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3UsChExtTable.EntityData.YListKeys = []string {}

    return &(docsIf3UsChExtTable.EntityData)
}

// DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry
// The conceptual row of docsIf3UsChExtTable.
// The ifIndex key corresponds to each of the upstream
// channels.
type DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute indicates the selection mode for active codes and code
    // hopping. 'none'    Non-SCDMA channel 'sac1NoCodeHopping'    Selectable
    // active codes mode 1 and code hopping disabled  'sac1CodeHoppingMode1'   
    // Selectable active codes mode 1 and code hopping mode 1
    // 'sac2CodeHoppingMode2'   Selectable active codes mode 2 and code hopping
    // mode 2 'sac2NoCodeHopping'   Selectable active codes mode 2 and code
    // hopping disabled. The type is DocsIf3UsChExtSacCodeHoppingSelectionMode.
    DocsIf3UsChExtSacCodeHoppingSelectionMode interface{}

    // This attribute represents the active codes of the upstream channel and it
    // is applicable only when SacCodeHoppingSelectionMode is
    // 'sac2CodeHoppingMode2. The type is string with length: 0 | 16.
    DocsIf3UsChExtScdmaSelectionStringActiveCodes interface{}
}

func (docsIf3UsChExtEntry *DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry) GetEntityData() *types.CommonEntityData {
    docsIf3UsChExtEntry.EntityData.YFilter = docsIf3UsChExtEntry.YFilter
    docsIf3UsChExtEntry.EntityData.YangName = "docsIf3UsChExtEntry"
    docsIf3UsChExtEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsChExtEntry.EntityData.ParentYangName = "docsIf3UsChExtTable"
    docsIf3UsChExtEntry.EntityData.SegmentPath = "docsIf3UsChExtEntry" + types.AddKeyToken(docsIf3UsChExtEntry.IfIndex, "ifIndex")
    docsIf3UsChExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsChExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsChExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsChExtEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3UsChExtEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3UsChExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3UsChExtEntry.IfIndex})
    docsIf3UsChExtEntry.EntityData.Leafs.Append("docsIf3UsChExtSacCodeHoppingSelectionMode", types.YLeaf{"DocsIf3UsChExtSacCodeHoppingSelectionMode", docsIf3UsChExtEntry.DocsIf3UsChExtSacCodeHoppingSelectionMode})
    docsIf3UsChExtEntry.EntityData.Leafs.Append("docsIf3UsChExtScdmaSelectionStringActiveCodes", types.YLeaf{"DocsIf3UsChExtScdmaSelectionStringActiveCodes", docsIf3UsChExtEntry.DocsIf3UsChExtScdmaSelectionStringActiveCodes})

    docsIf3UsChExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3UsChExtEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode represents   Selectable active codes mode 2 and code hopping disabled.
type DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode string

const (
    DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode_none DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode = "none"

    DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode_sac1NoCodeHopping DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode = "sac1NoCodeHopping"

    DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode_sac1CodeHoppingMode1 DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode = "sac1CodeHoppingMode1"

    DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode_sac2CodeHoppingMode2 DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode = "sac2CodeHoppingMode2"

    DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode_sac2NoCodeHopping DOCSIF3MIB_DocsIf3UsChExtTable_DocsIf3UsChExtEntry_DocsIf3UsChExtSacCodeHoppingSelectionMode = "sac2NoCodeHopping"
)

// DOCSIF3MIB_DocsIf3UsChSetTable
// This object defines a set of upstream channels. These
// channel sets may be associated with channel bonding
// groups, MD-US-SGs, MD-CM-SGs, or any other channel
// set that the CMTS may derive from other CMTS processes.
type DOCSIF3MIB_DocsIf3UsChSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3UsChSetTable. The ifIndex key corresponds to
    // the MAC Domain interface where the upstream channel set is defined. The
    // type is slice of DOCSIF3MIB_DocsIf3UsChSetTable_DocsIf3UsChSetEntry.
    DocsIf3UsChSetEntry []*DOCSIF3MIB_DocsIf3UsChSetTable_DocsIf3UsChSetEntry
}

func (docsIf3UsChSetTable *DOCSIF3MIB_DocsIf3UsChSetTable) GetEntityData() *types.CommonEntityData {
    docsIf3UsChSetTable.EntityData.YFilter = docsIf3UsChSetTable.YFilter
    docsIf3UsChSetTable.EntityData.YangName = "docsIf3UsChSetTable"
    docsIf3UsChSetTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsChSetTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3UsChSetTable.EntityData.SegmentPath = "docsIf3UsChSetTable"
    docsIf3UsChSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsChSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsChSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsChSetTable.EntityData.Children = types.NewOrderedMap()
    docsIf3UsChSetTable.EntityData.Children.Append("docsIf3UsChSetEntry", types.YChild{"DocsIf3UsChSetEntry", nil})
    for i := range docsIf3UsChSetTable.DocsIf3UsChSetEntry {
        docsIf3UsChSetTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3UsChSetTable.DocsIf3UsChSetEntry[i]), types.YChild{"DocsIf3UsChSetEntry", docsIf3UsChSetTable.DocsIf3UsChSetEntry[i]})
    }
    docsIf3UsChSetTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3UsChSetTable.EntityData.YListKeys = []string {}

    return &(docsIf3UsChSetTable.EntityData)
}

// DOCSIF3MIB_DocsIf3UsChSetTable_DocsIf3UsChSetEntry
// The conceptual row of docsIf3UsChSetTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the upstream channel set is defined.
type DOCSIF3MIB_DocsIf3UsChSetTable_DocsIf3UsChSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key defines a reference identifier for the
    // upstream channel set within the MAC Domain. The type is interface{} with
    // range: 0..4294967295.
    DocsIf3UsChSetId interface{}

    // This attribute defines the ordered list of channels that comprise the
    // upstream channel set. The type is string with length: 0 | 2..255.
    DocsIf3UsChSetChList interface{}
}

func (docsIf3UsChSetEntry *DOCSIF3MIB_DocsIf3UsChSetTable_DocsIf3UsChSetEntry) GetEntityData() *types.CommonEntityData {
    docsIf3UsChSetEntry.EntityData.YFilter = docsIf3UsChSetEntry.YFilter
    docsIf3UsChSetEntry.EntityData.YangName = "docsIf3UsChSetEntry"
    docsIf3UsChSetEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3UsChSetEntry.EntityData.ParentYangName = "docsIf3UsChSetTable"
    docsIf3UsChSetEntry.EntityData.SegmentPath = "docsIf3UsChSetEntry" + types.AddKeyToken(docsIf3UsChSetEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3UsChSetEntry.DocsIf3UsChSetId, "docsIf3UsChSetId")
    docsIf3UsChSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3UsChSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3UsChSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3UsChSetEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3UsChSetEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3UsChSetEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3UsChSetEntry.IfIndex})
    docsIf3UsChSetEntry.EntityData.Leafs.Append("docsIf3UsChSetId", types.YLeaf{"DocsIf3UsChSetId", docsIf3UsChSetEntry.DocsIf3UsChSetId})
    docsIf3UsChSetEntry.EntityData.Leafs.Append("docsIf3UsChSetChList", types.YLeaf{"DocsIf3UsChSetChList", docsIf3UsChSetEntry.DocsIf3UsChSetChList})

    docsIf3UsChSetEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3UsChSetId"}

    return &(docsIf3UsChSetEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3DsChSetTable
// This object defines a set of downstream channels.
// These channel sets may be associated with channel bonding
// groups, MD-DS-SGs, MD-CM-SGs, or any other channel
// set that the CMTS may derive from other CMTS processes.
type DOCSIF3MIB_DocsIf3DsChSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3DsChSetTable. The ifIndex key corresponds to
    // the MAC Domain interface where the downstream channel set is defined. The
    // type is slice of DOCSIF3MIB_DocsIf3DsChSetTable_DocsIf3DsChSetEntry.
    DocsIf3DsChSetEntry []*DOCSIF3MIB_DocsIf3DsChSetTable_DocsIf3DsChSetEntry
}

func (docsIf3DsChSetTable *DOCSIF3MIB_DocsIf3DsChSetTable) GetEntityData() *types.CommonEntityData {
    docsIf3DsChSetTable.EntityData.YFilter = docsIf3DsChSetTable.YFilter
    docsIf3DsChSetTable.EntityData.YangName = "docsIf3DsChSetTable"
    docsIf3DsChSetTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3DsChSetTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3DsChSetTable.EntityData.SegmentPath = "docsIf3DsChSetTable"
    docsIf3DsChSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3DsChSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3DsChSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3DsChSetTable.EntityData.Children = types.NewOrderedMap()
    docsIf3DsChSetTable.EntityData.Children.Append("docsIf3DsChSetEntry", types.YChild{"DocsIf3DsChSetEntry", nil})
    for i := range docsIf3DsChSetTable.DocsIf3DsChSetEntry {
        docsIf3DsChSetTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3DsChSetTable.DocsIf3DsChSetEntry[i]), types.YChild{"DocsIf3DsChSetEntry", docsIf3DsChSetTable.DocsIf3DsChSetEntry[i]})
    }
    docsIf3DsChSetTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3DsChSetTable.EntityData.YListKeys = []string {}

    return &(docsIf3DsChSetTable.EntityData)
}

// DOCSIF3MIB_DocsIf3DsChSetTable_DocsIf3DsChSetEntry
// The conceptual row of docsIf3DsChSetTable.
// The ifIndex key corresponds to the MAC Domain interface
// where the downstream channel set is defined.
type DOCSIF3MIB_DocsIf3DsChSetTable_DocsIf3DsChSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key defines a reference identifier for the
    // downstream channel set within the MAC Domain. The type is interface{} with
    // range: 0..4294967295.
    DocsIf3DsChSetId interface{}

    // This attribute defines the ordered list of channels that comprise the
    // downstream channel set. The type is string with length: 0 | 2..255.
    DocsIf3DsChSetChList interface{}
}

func (docsIf3DsChSetEntry *DOCSIF3MIB_DocsIf3DsChSetTable_DocsIf3DsChSetEntry) GetEntityData() *types.CommonEntityData {
    docsIf3DsChSetEntry.EntityData.YFilter = docsIf3DsChSetEntry.YFilter
    docsIf3DsChSetEntry.EntityData.YangName = "docsIf3DsChSetEntry"
    docsIf3DsChSetEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3DsChSetEntry.EntityData.ParentYangName = "docsIf3DsChSetTable"
    docsIf3DsChSetEntry.EntityData.SegmentPath = "docsIf3DsChSetEntry" + types.AddKeyToken(docsIf3DsChSetEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3DsChSetEntry.DocsIf3DsChSetId, "docsIf3DsChSetId")
    docsIf3DsChSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3DsChSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3DsChSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3DsChSetEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3DsChSetEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3DsChSetEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3DsChSetEntry.IfIndex})
    docsIf3DsChSetEntry.EntityData.Leafs.Append("docsIf3DsChSetId", types.YLeaf{"DocsIf3DsChSetId", docsIf3DsChSetEntry.DocsIf3DsChSetId})
    docsIf3DsChSetEntry.EntityData.Leafs.Append("docsIf3DsChSetChList", types.YLeaf{"DocsIf3DsChSetChList", docsIf3DsChSetEntry.DocsIf3DsChSetChList})

    docsIf3DsChSetEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3DsChSetId"}

    return &(docsIf3DsChSetEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3SignalQualityExtTable
// This object provides an in-channel received modulation
// error ratio metric for CM and CMTS.
type DOCSIF3MIB_DocsIf3SignalQualityExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3SignalQualityExtTable. At the CM, this object
    // describes the received modulation  error ratio of each downstream channel.
    // At the CMTS,  it describes the received modulation error ratio of each
    // logical upstream channel.  An entry in this table exists for each ifEntry
    // with an ifType of docsCableDownstream(128) for Cable Modems. For Cable
    // Modem Termination Systems, an entry exists for  each ifEntry with an ifType
    // of docsCableUpstreamChannel(205). The type is slice of
    // DOCSIF3MIB_DocsIf3SignalQualityExtTable_DocsIf3SignalQualityExtEntry.
    DocsIf3SignalQualityExtEntry []*DOCSIF3MIB_DocsIf3SignalQualityExtTable_DocsIf3SignalQualityExtEntry
}

func (docsIf3SignalQualityExtTable *DOCSIF3MIB_DocsIf3SignalQualityExtTable) GetEntityData() *types.CommonEntityData {
    docsIf3SignalQualityExtTable.EntityData.YFilter = docsIf3SignalQualityExtTable.YFilter
    docsIf3SignalQualityExtTable.EntityData.YangName = "docsIf3SignalQualityExtTable"
    docsIf3SignalQualityExtTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3SignalQualityExtTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3SignalQualityExtTable.EntityData.SegmentPath = "docsIf3SignalQualityExtTable"
    docsIf3SignalQualityExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3SignalQualityExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3SignalQualityExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3SignalQualityExtTable.EntityData.Children = types.NewOrderedMap()
    docsIf3SignalQualityExtTable.EntityData.Children.Append("docsIf3SignalQualityExtEntry", types.YChild{"DocsIf3SignalQualityExtEntry", nil})
    for i := range docsIf3SignalQualityExtTable.DocsIf3SignalQualityExtEntry {
        docsIf3SignalQualityExtTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3SignalQualityExtTable.DocsIf3SignalQualityExtEntry[i]), types.YChild{"DocsIf3SignalQualityExtEntry", docsIf3SignalQualityExtTable.DocsIf3SignalQualityExtEntry[i]})
    }
    docsIf3SignalQualityExtTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3SignalQualityExtTable.EntityData.YListKeys = []string {}

    return &(docsIf3SignalQualityExtTable.EntityData)
}

// DOCSIF3MIB_DocsIf3SignalQualityExtTable_DocsIf3SignalQualityExtEntry
// The conceptual row of docsIf3SignalQualityExtTable.
// At the CM, this object describes the received modulation 
// error ratio of each downstream channel. At the CMTS, 
// it describes the received modulation error ratio of each
// logical upstream channel. 
// An entry in this table exists for each ifEntry with an
// ifType of docsCableDownstream(128) for Cable Modems.
// For Cable Modem Termination Systems, an entry exists for 
// each ifEntry with an ifType of docsCableUpstreamChannel(205).
type DOCSIF3MIB_DocsIf3SignalQualityExtTable_DocsIf3SignalQualityExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // RxMER provides an in-channel received Modulation Error Ratio (MER). RxMER
    // is defined as an estimate, provided by the demodulator, of the ratio:
    // (average constellation energy with equally likely symbols) / (average
    // squared magnitude of error vector)  RxMER is measured just prior to FEC
    // (trellis/Reed-Solomon) decoding. RxMER includes the effects of the HFC
    // channel as well as implementation effects of the modulator and demodulator.
    // Error vector estimation may vary among demodulator implementations.  The
    // CMTS RxMER is averaged over a given number of bursts at the burst receiver,
    // which may correspond to transmissions from multiple users. In the case of
    // S-CDMA mode, RxMER is measured on the de-spread signal. The type is
    // interface{} with range: -2147483648..2147483647. Units are TenthdB.
    DocsIf3SignalQualityExtRxMER interface{}

    // RxMerSamples is a statistically significant number of symbols for the CM,
    // or bursts for the CMTS, processed to arrive at the RxMER value. For the
    // CMTS, the MER measurement includes only valid bursts that are not in
    // contention regions. The type is interface{} with range: 0..4294967295.
    DocsIf3SignalQualityExtRxMerSamples interface{}
}

func (docsIf3SignalQualityExtEntry *DOCSIF3MIB_DocsIf3SignalQualityExtTable_DocsIf3SignalQualityExtEntry) GetEntityData() *types.CommonEntityData {
    docsIf3SignalQualityExtEntry.EntityData.YFilter = docsIf3SignalQualityExtEntry.YFilter
    docsIf3SignalQualityExtEntry.EntityData.YangName = "docsIf3SignalQualityExtEntry"
    docsIf3SignalQualityExtEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3SignalQualityExtEntry.EntityData.ParentYangName = "docsIf3SignalQualityExtTable"
    docsIf3SignalQualityExtEntry.EntityData.SegmentPath = "docsIf3SignalQualityExtEntry" + types.AddKeyToken(docsIf3SignalQualityExtEntry.IfIndex, "ifIndex")
    docsIf3SignalQualityExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3SignalQualityExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3SignalQualityExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3SignalQualityExtEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3SignalQualityExtEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3SignalQualityExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3SignalQualityExtEntry.IfIndex})
    docsIf3SignalQualityExtEntry.EntityData.Leafs.Append("docsIf3SignalQualityExtRxMER", types.YLeaf{"DocsIf3SignalQualityExtRxMER", docsIf3SignalQualityExtEntry.DocsIf3SignalQualityExtRxMER})
    docsIf3SignalQualityExtEntry.EntityData.Leafs.Append("docsIf3SignalQualityExtRxMerSamples", types.YLeaf{"DocsIf3SignalQualityExtRxMerSamples", docsIf3SignalQualityExtEntry.DocsIf3SignalQualityExtRxMerSamples})

    docsIf3SignalQualityExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3SignalQualityExtEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable
// This object provides metrics and parameters associated
// with received carrier, noise and interference
// power levels in the upstream channels of the CMTS.
type DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsSignalQualityExtTable. The ifIndex key
    // corresponds to each of the upstream channels. The CMTS is not required to
    // persist the values of all  instances of CmtsSignalQualityExt across
    // reinitialization. The type is slice of
    // DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable_DocsIf3CmtsSignalQualityExtEntry.
    DocsIf3CmtsSignalQualityExtEntry []*DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable_DocsIf3CmtsSignalQualityExtEntry
}

func (docsIf3CmtsSignalQualityExtTable *DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsSignalQualityExtTable.EntityData.YFilter = docsIf3CmtsSignalQualityExtTable.YFilter
    docsIf3CmtsSignalQualityExtTable.EntityData.YangName = "docsIf3CmtsSignalQualityExtTable"
    docsIf3CmtsSignalQualityExtTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsSignalQualityExtTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsSignalQualityExtTable.EntityData.SegmentPath = "docsIf3CmtsSignalQualityExtTable"
    docsIf3CmtsSignalQualityExtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsSignalQualityExtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsSignalQualityExtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsSignalQualityExtTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsSignalQualityExtTable.EntityData.Children.Append("docsIf3CmtsSignalQualityExtEntry", types.YChild{"DocsIf3CmtsSignalQualityExtEntry", nil})
    for i := range docsIf3CmtsSignalQualityExtTable.DocsIf3CmtsSignalQualityExtEntry {
        docsIf3CmtsSignalQualityExtTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsSignalQualityExtTable.DocsIf3CmtsSignalQualityExtEntry[i]), types.YChild{"DocsIf3CmtsSignalQualityExtEntry", docsIf3CmtsSignalQualityExtTable.DocsIf3CmtsSignalQualityExtEntry[i]})
    }
    docsIf3CmtsSignalQualityExtTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsSignalQualityExtTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsSignalQualityExtTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable_DocsIf3CmtsSignalQualityExtEntry
// The conceptual row of docsIf3CmtsSignalQualityExtTable.
// The ifIndex key corresponds to each of the upstream
// channels.
// The CMTS is not required to persist the values of all 
// instances of CmtsSignalQualityExt across reinitialization.
type DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable_DocsIf3CmtsSignalQualityExtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute provides an upstream in-channel Carrier-to-Noise plus
    // Interference Ratio (CNIR). CNIR is defined as the ratio of the expected
    // commanded received signal power at the CMTS input, assuming QPSK0
    // modulation, to the noise plus interference in the channel. This measurement
    // occurs prior to the point at which  the desired CM signal, when present, is
    // demodulated. The measurement includes the effect of the receive matched
    // filter but does not include the effect of any ingress filtering. Both the
    // signal power and noise/interference power are referenced to the same point,
    // e.g., CMTS input. The type is interface{} with range:
    // -2147483648..2147483647.
    DocsIf3CmtsSignalQualityExtCNIR interface{}

    // ExpectedReceivedSignalPower is the power of the expected commanded received
    // signal in the channel, referenced to the CMTS input. The type is
    // interface{} with range: -2147483648..2147483647.
    DocsIf3CmtsSignalQualityExtExpectedRxSignalPower interface{}
}

func (docsIf3CmtsSignalQualityExtEntry *DOCSIF3MIB_DocsIf3CmtsSignalQualityExtTable_DocsIf3CmtsSignalQualityExtEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsSignalQualityExtEntry.EntityData.YFilter = docsIf3CmtsSignalQualityExtEntry.YFilter
    docsIf3CmtsSignalQualityExtEntry.EntityData.YangName = "docsIf3CmtsSignalQualityExtEntry"
    docsIf3CmtsSignalQualityExtEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsSignalQualityExtEntry.EntityData.ParentYangName = "docsIf3CmtsSignalQualityExtTable"
    docsIf3CmtsSignalQualityExtEntry.EntityData.SegmentPath = "docsIf3CmtsSignalQualityExtEntry" + types.AddKeyToken(docsIf3CmtsSignalQualityExtEntry.IfIndex, "ifIndex")
    docsIf3CmtsSignalQualityExtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsSignalQualityExtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsSignalQualityExtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsSignalQualityExtEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsSignalQualityExtEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsSignalQualityExtEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmtsSignalQualityExtEntry.IfIndex})
    docsIf3CmtsSignalQualityExtEntry.EntityData.Leafs.Append("docsIf3CmtsSignalQualityExtCNIR", types.YLeaf{"DocsIf3CmtsSignalQualityExtCNIR", docsIf3CmtsSignalQualityExtEntry.DocsIf3CmtsSignalQualityExtCNIR})
    docsIf3CmtsSignalQualityExtEntry.EntityData.Leafs.Append("docsIf3CmtsSignalQualityExtExpectedRxSignalPower", types.YLeaf{"DocsIf3CmtsSignalQualityExtExpectedRxSignalPower", docsIf3CmtsSignalQualityExtEntry.DocsIf3CmtsSignalQualityExtExpectedRxSignalPower})

    docsIf3CmtsSignalQualityExtEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmtsSignalQualityExtEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable
// This object is used to configure the logical upstream
// interfaces to perform the spectrum measurements.
// This object supports creation and deletion of instances.
type DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsSpectrumAnalysisMeasTable. The ifIndex key
    // corresponds to each of the upstream channels. The CMTS is not required to
    // persist instances of this object across reinitializations. The type is
    // slice of
    // DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable_DocsIf3CmtsSpectrumAnalysisMeasEntry.
    DocsIf3CmtsSpectrumAnalysisMeasEntry []*DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable_DocsIf3CmtsSpectrumAnalysisMeasEntry
}

func (docsIf3CmtsSpectrumAnalysisMeasTable *DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.YFilter = docsIf3CmtsSpectrumAnalysisMeasTable.YFilter
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.YangName = "docsIf3CmtsSpectrumAnalysisMeasTable"
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.SegmentPath = "docsIf3CmtsSpectrumAnalysisMeasTable"
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.Children.Append("docsIf3CmtsSpectrumAnalysisMeasEntry", types.YChild{"DocsIf3CmtsSpectrumAnalysisMeasEntry", nil})
    for i := range docsIf3CmtsSpectrumAnalysisMeasTable.DocsIf3CmtsSpectrumAnalysisMeasEntry {
        docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsSpectrumAnalysisMeasTable.DocsIf3CmtsSpectrumAnalysisMeasEntry[i]), types.YChild{"DocsIf3CmtsSpectrumAnalysisMeasEntry", docsIf3CmtsSpectrumAnalysisMeasTable.DocsIf3CmtsSpectrumAnalysisMeasEntry[i]})
    }
    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsSpectrumAnalysisMeasTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsSpectrumAnalysisMeasTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable_DocsIf3CmtsSpectrumAnalysisMeasEntry
// The conceptual row of docsIf3CmtsSpectrumAnalysisMeasTable.
// The ifIndex key corresponds to each of the upstream
// channels.
// The CMTS is not required to persist instances of this
// object across reinitializations.
type DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable_DocsIf3CmtsSpectrumAnalysisMeasEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute provides a list of the spectral amplitudes corresponding to
    // the frequency bins ordered from lowest to highest frequencies covering the
    // frequency span. Information about the center frequency, frequency span,
    // number of bins and resolution bandwidth are included to provide context to
    // the measurement point The CMTS must support the number of bins as an odd
    // number in order to provide a spectrum representation that is symmetric
    // about the middle data point or bin. The CMTS must support a number of bins
    // greater than or equal to 257 for frequency spans greater than or equal to
    // 6.4 MHz.  The CMTS must not exceed 25 kHz bin spacing for measurement of
    // frequency spans less than or equal to 6.4 MHz.  The bins measurements are
    // updated periodically at time intervals given by the TimeInterval attribute.
    // The type is string with length: 0 | 2..4116.
    DocsIf3CmtsSpectrumAnalysisMeasAmplitudeData interface{}

    // TimeInterval is the CMTS estimated average repetition period of
    // measurements. This attribute defines the average rate at which new spectra
    // can be retrieved. The type is interface{} with range: 0..4294967295. Units
    // are milliseconds.
    DocsIf3CmtsSpectrumAnalysisMeasTimeInterval interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3CmtsSpectrumAnalysisMeasRowStatus interface{}
}

func (docsIf3CmtsSpectrumAnalysisMeasEntry *DOCSIF3MIB_DocsIf3CmtsSpectrumAnalysisMeasTable_DocsIf3CmtsSpectrumAnalysisMeasEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.YFilter = docsIf3CmtsSpectrumAnalysisMeasEntry.YFilter
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.YangName = "docsIf3CmtsSpectrumAnalysisMeasEntry"
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.ParentYangName = "docsIf3CmtsSpectrumAnalysisMeasTable"
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.SegmentPath = "docsIf3CmtsSpectrumAnalysisMeasEntry" + types.AddKeyToken(docsIf3CmtsSpectrumAnalysisMeasEntry.IfIndex, "ifIndex")
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmtsSpectrumAnalysisMeasEntry.IfIndex})
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmtsSpectrumAnalysisMeasAmplitudeData", types.YLeaf{"DocsIf3CmtsSpectrumAnalysisMeasAmplitudeData", docsIf3CmtsSpectrumAnalysisMeasEntry.DocsIf3CmtsSpectrumAnalysisMeasAmplitudeData})
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmtsSpectrumAnalysisMeasTimeInterval", types.YLeaf{"DocsIf3CmtsSpectrumAnalysisMeasTimeInterval", docsIf3CmtsSpectrumAnalysisMeasEntry.DocsIf3CmtsSpectrumAnalysisMeasTimeInterval})
    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmtsSpectrumAnalysisMeasRowStatus", types.YLeaf{"DocsIf3CmtsSpectrumAnalysisMeasRowStatus", docsIf3CmtsSpectrumAnalysisMeasEntry.DocsIf3CmtsSpectrumAnalysisMeasRowStatus})

    docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmtsSpectrumAnalysisMeasEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmDpvStatsTable
// This object represents the DOCSIS Path Verify Statistics
// collected in the cable modem device.
// The CMTS controls the logging of DPV statistics in the
// cable modem. Therefore the context and nature of the
// measurements are governed by the CMTS and not self-descriptive
// when read from the CM.
type DOCSIF3MIB_DocsIf3CmDpvStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmDpvStatsTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmDpvStatsTable_DocsIf3CmDpvStatsEntry.
    DocsIf3CmDpvStatsEntry []*DOCSIF3MIB_DocsIf3CmDpvStatsTable_DocsIf3CmDpvStatsEntry
}

func (docsIf3CmDpvStatsTable *DOCSIF3MIB_DocsIf3CmDpvStatsTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmDpvStatsTable.EntityData.YFilter = docsIf3CmDpvStatsTable.YFilter
    docsIf3CmDpvStatsTable.EntityData.YangName = "docsIf3CmDpvStatsTable"
    docsIf3CmDpvStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmDpvStatsTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmDpvStatsTable.EntityData.SegmentPath = "docsIf3CmDpvStatsTable"
    docsIf3CmDpvStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmDpvStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmDpvStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmDpvStatsTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmDpvStatsTable.EntityData.Children.Append("docsIf3CmDpvStatsEntry", types.YChild{"DocsIf3CmDpvStatsEntry", nil})
    for i := range docsIf3CmDpvStatsTable.DocsIf3CmDpvStatsEntry {
        docsIf3CmDpvStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmDpvStatsTable.DocsIf3CmDpvStatsEntry[i]), types.YChild{"DocsIf3CmDpvStatsEntry", docsIf3CmDpvStatsTable.DocsIf3CmDpvStatsEntry[i]})
    }
    docsIf3CmDpvStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmDpvStatsTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmDpvStatsTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmDpvStatsTable_DocsIf3CmDpvStatsEntry
// The conceptual row of docsIf3CmDpvStatsTable.
type DOCSIF3MIB_DocsIf3CmDpvStatsTable_DocsIf3CmDpvStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. This key represents the DPV Group ID. The CM
    // reports two instance of DPV statistics per downstream normally referred as
    // Statistical Group 1 and Statistical Group 2. The type is interface{} with
    // range: 1..2.
    DocsIf3CmDpvStatsGrpId interface{}

    // This attribute represents the last latency measurement for this statistical
    // group. The type is interface{} with range: 0..4294967295. Units are
    // nanoseconds.
    DocsIf3CmDpvStatsLastMeasLatency interface{}

    // This attribute represents the last measurement time of the last latency
    // measurement for this statistical group. This attribute reports the EPOC
    // time value when no measurements are being reported or after the statistics
    // were cleared. The type is string.
    DocsIf3CmDpvStatsLastMeasTime interface{}

    // This attribute represents the minimum latency measurement for this
    // statistical group since the last time statistics were cleared. The type is
    // interface{} with range: 0..4294967295. Units are nanoseconds.
    DocsIf3CmDpvStatsMinLatency interface{}

    // This attribute represents the maximum latency measurement for this
    // statistical group since the last time statistics were cleared. The type is
    // interface{} with range: 0..4294967295. Units are nanoseconds.
    DocsIf3CmDpvStatsMaxLatency interface{}

    // This attribute represents the average latency measurement for this
    // statistical group since the last time statistics were cleared. The
    // averaging mechanism is controlled by the CMTS, and can be a simple average
    // (mean) or an exponential moving average. The type is interface{} with
    // range: 0..4294967295. Units are nanoseconds.
    DocsIf3CmDpvStatsAvgLatency interface{}

    // This attribute represents the number of latency measurements made for this
    // statistical group since the last time statistics were cleared. The type is
    // interface{} with range: 0..4294967295. Units are measurements.
    DocsIf3CmDpvStatsNumMeas interface{}

    // This attribute represents the last time statistics were cleared for this
    // statistical group. The type is string.
    DocsIf3CmDpvStatsLastClearTime interface{}
}

func (docsIf3CmDpvStatsEntry *DOCSIF3MIB_DocsIf3CmDpvStatsTable_DocsIf3CmDpvStatsEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmDpvStatsEntry.EntityData.YFilter = docsIf3CmDpvStatsEntry.YFilter
    docsIf3CmDpvStatsEntry.EntityData.YangName = "docsIf3CmDpvStatsEntry"
    docsIf3CmDpvStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmDpvStatsEntry.EntityData.ParentYangName = "docsIf3CmDpvStatsTable"
    docsIf3CmDpvStatsEntry.EntityData.SegmentPath = "docsIf3CmDpvStatsEntry" + types.AddKeyToken(docsIf3CmDpvStatsEntry.IfIndex, "ifIndex") + types.AddKeyToken(docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsGrpId, "docsIf3CmDpvStatsGrpId")
    docsIf3CmDpvStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmDpvStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmDpvStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmDpvStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmDpvStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmDpvStatsEntry.IfIndex})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsGrpId", types.YLeaf{"DocsIf3CmDpvStatsGrpId", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsGrpId})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsLastMeasLatency", types.YLeaf{"DocsIf3CmDpvStatsLastMeasLatency", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsLastMeasLatency})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsLastMeasTime", types.YLeaf{"DocsIf3CmDpvStatsLastMeasTime", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsLastMeasTime})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsMinLatency", types.YLeaf{"DocsIf3CmDpvStatsMinLatency", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsMinLatency})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsMaxLatency", types.YLeaf{"DocsIf3CmDpvStatsMaxLatency", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsMaxLatency})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsAvgLatency", types.YLeaf{"DocsIf3CmDpvStatsAvgLatency", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsAvgLatency})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsNumMeas", types.YLeaf{"DocsIf3CmDpvStatsNumMeas", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsNumMeas})
    docsIf3CmDpvStatsEntry.EntityData.Leafs.Append("docsIf3CmDpvStatsLastClearTime", types.YLeaf{"DocsIf3CmDpvStatsLastClearTime", docsIf3CmDpvStatsEntry.DocsIf3CmDpvStatsLastClearTime})

    docsIf3CmDpvStatsEntry.EntityData.YListKeys = []string {"IfIndex", "DocsIf3CmDpvStatsGrpId"}

    return &(docsIf3CmDpvStatsEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmEventCtrlTable
// This object represents the control mechanism to enable the
// dispatching of events based on the event Id. The following
// rules define the event control behavior:
// 
// - The CmEventCtrl object has no instances or contains an 
//   instance with Event ID 0. All events matching the Local Log
//   settings of docsDevEvReporting are sent to local log ONLY.
// 
// - The CmEventCtrl object contains configured instances
//   Only events matching the Event Ids configured in the object
//   are sent according to the settings of the docsDevEvReporting
//   object. 
//   
//   The CM does not persist instances of CmEventCtrl across 
//   reinitializations.
type DOCSIF3MIB_DocsIf3CmEventCtrlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmEventCtrlTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmEventCtrlTable_DocsIf3CmEventCtrlEntry.
    DocsIf3CmEventCtrlEntry []*DOCSIF3MIB_DocsIf3CmEventCtrlTable_DocsIf3CmEventCtrlEntry
}

func (docsIf3CmEventCtrlTable *DOCSIF3MIB_DocsIf3CmEventCtrlTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmEventCtrlTable.EntityData.YFilter = docsIf3CmEventCtrlTable.YFilter
    docsIf3CmEventCtrlTable.EntityData.YangName = "docsIf3CmEventCtrlTable"
    docsIf3CmEventCtrlTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEventCtrlTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmEventCtrlTable.EntityData.SegmentPath = "docsIf3CmEventCtrlTable"
    docsIf3CmEventCtrlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEventCtrlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEventCtrlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEventCtrlTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEventCtrlTable.EntityData.Children.Append("docsIf3CmEventCtrlEntry", types.YChild{"DocsIf3CmEventCtrlEntry", nil})
    for i := range docsIf3CmEventCtrlTable.DocsIf3CmEventCtrlEntry {
        docsIf3CmEventCtrlTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmEventCtrlTable.DocsIf3CmEventCtrlEntry[i]), types.YChild{"DocsIf3CmEventCtrlEntry", docsIf3CmEventCtrlTable.DocsIf3CmEventCtrlEntry[i]})
    }
    docsIf3CmEventCtrlTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmEventCtrlTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmEventCtrlTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmEventCtrlTable_DocsIf3CmEventCtrlEntry
// The conceptual row of docsIf3CmEventCtrlTable.
type DOCSIF3MIB_DocsIf3CmEventCtrlTable_DocsIf3CmEventCtrlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This key represents the Event ID of the event
    // being  enabled for delivery to a dispatch mechanism (e.g., syslog). The
    // type is interface{} with range: 0..4294967295.
    DocsIf3CmEventCtrlEventId interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3CmEventCtrlStatus interface{}
}

func (docsIf3CmEventCtrlEntry *DOCSIF3MIB_DocsIf3CmEventCtrlTable_DocsIf3CmEventCtrlEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmEventCtrlEntry.EntityData.YFilter = docsIf3CmEventCtrlEntry.YFilter
    docsIf3CmEventCtrlEntry.EntityData.YangName = "docsIf3CmEventCtrlEntry"
    docsIf3CmEventCtrlEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEventCtrlEntry.EntityData.ParentYangName = "docsIf3CmEventCtrlTable"
    docsIf3CmEventCtrlEntry.EntityData.SegmentPath = "docsIf3CmEventCtrlEntry" + types.AddKeyToken(docsIf3CmEventCtrlEntry.DocsIf3CmEventCtrlEventId, "docsIf3CmEventCtrlEventId")
    docsIf3CmEventCtrlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEventCtrlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEventCtrlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEventCtrlEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEventCtrlEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmEventCtrlEntry.EntityData.Leafs.Append("docsIf3CmEventCtrlEventId", types.YLeaf{"DocsIf3CmEventCtrlEventId", docsIf3CmEventCtrlEntry.DocsIf3CmEventCtrlEventId})
    docsIf3CmEventCtrlEntry.EntityData.Leafs.Append("docsIf3CmEventCtrlStatus", types.YLeaf{"DocsIf3CmEventCtrlStatus", docsIf3CmEventCtrlEntry.DocsIf3CmEventCtrlStatus})

    docsIf3CmEventCtrlEntry.EntityData.YListKeys = []string {"DocsIf3CmEventCtrlEventId"}

    return &(docsIf3CmEventCtrlEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsEventCtrlTable
// This object represents the control mechanism to enable the
// dispatching of  events based on the event Id. The following
// rules define the event control behavior:
// 
// - The CmtsEventCtrl object has no instances or contains an 
//   instance with Event ID 0. All events matching the Local Log 
//   settings of docsDevEvReporting are sent to local log ONLY.
// 
// - The CmtsEventCtrl object contains configured instances
//   Only events matching the Event Ids configured in the object
//   are sent according to the settings of the docsDevEvReporting
//   object. 
//   
// - The CMTS persists all instances of CmtsEventCtrl across 
//   reinitializations.
type DOCSIF3MIB_DocsIf3CmtsEventCtrlTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsEventCtrlTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmtsEventCtrlTable_DocsIf3CmtsEventCtrlEntry.
    DocsIf3CmtsEventCtrlEntry []*DOCSIF3MIB_DocsIf3CmtsEventCtrlTable_DocsIf3CmtsEventCtrlEntry
}

func (docsIf3CmtsEventCtrlTable *DOCSIF3MIB_DocsIf3CmtsEventCtrlTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsEventCtrlTable.EntityData.YFilter = docsIf3CmtsEventCtrlTable.YFilter
    docsIf3CmtsEventCtrlTable.EntityData.YangName = "docsIf3CmtsEventCtrlTable"
    docsIf3CmtsEventCtrlTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsEventCtrlTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsEventCtrlTable.EntityData.SegmentPath = "docsIf3CmtsEventCtrlTable"
    docsIf3CmtsEventCtrlTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsEventCtrlTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsEventCtrlTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsEventCtrlTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsEventCtrlTable.EntityData.Children.Append("docsIf3CmtsEventCtrlEntry", types.YChild{"DocsIf3CmtsEventCtrlEntry", nil})
    for i := range docsIf3CmtsEventCtrlTable.DocsIf3CmtsEventCtrlEntry {
        docsIf3CmtsEventCtrlTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsEventCtrlTable.DocsIf3CmtsEventCtrlEntry[i]), types.YChild{"DocsIf3CmtsEventCtrlEntry", docsIf3CmtsEventCtrlTable.DocsIf3CmtsEventCtrlEntry[i]})
    }
    docsIf3CmtsEventCtrlTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsEventCtrlTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsEventCtrlTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsEventCtrlTable_DocsIf3CmtsEventCtrlEntry
// The conceptual row of docsIf3CmtsEventCtrlTable.
type DOCSIF3MIB_DocsIf3CmtsEventCtrlTable_DocsIf3CmtsEventCtrlEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This key represents the Event ID of the event
    // being  enabled for delivery to a dispatch mechanism (e.g., syslog). The
    // type is interface{} with range: 0..4294967295.
    DocsIf3CmtsEventCtrlEventId interface{}

    // The status of this instance. The type is RowStatus.
    DocsIf3CmtsEventCtrlStatus interface{}
}

func (docsIf3CmtsEventCtrlEntry *DOCSIF3MIB_DocsIf3CmtsEventCtrlTable_DocsIf3CmtsEventCtrlEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsEventCtrlEntry.EntityData.YFilter = docsIf3CmtsEventCtrlEntry.YFilter
    docsIf3CmtsEventCtrlEntry.EntityData.YangName = "docsIf3CmtsEventCtrlEntry"
    docsIf3CmtsEventCtrlEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsEventCtrlEntry.EntityData.ParentYangName = "docsIf3CmtsEventCtrlTable"
    docsIf3CmtsEventCtrlEntry.EntityData.SegmentPath = "docsIf3CmtsEventCtrlEntry" + types.AddKeyToken(docsIf3CmtsEventCtrlEntry.DocsIf3CmtsEventCtrlEventId, "docsIf3CmtsEventCtrlEventId")
    docsIf3CmtsEventCtrlEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsEventCtrlEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsEventCtrlEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsEventCtrlEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsEventCtrlEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsEventCtrlEntry.EntityData.Leafs.Append("docsIf3CmtsEventCtrlEventId", types.YLeaf{"DocsIf3CmtsEventCtrlEventId", docsIf3CmtsEventCtrlEntry.DocsIf3CmtsEventCtrlEventId})
    docsIf3CmtsEventCtrlEntry.EntityData.Leafs.Append("docsIf3CmtsEventCtrlStatus", types.YLeaf{"DocsIf3CmtsEventCtrlStatus", docsIf3CmtsEventCtrlEntry.DocsIf3CmtsEventCtrlStatus})

    docsIf3CmtsEventCtrlEntry.EntityData.YListKeys = []string {"DocsIf3CmtsEventCtrlEventId"}

    return &(docsIf3CmtsEventCtrlEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmMdCfgTable
// This object contains CM MAC domain level control and
// configuration attributes.
type DOCSIF3MIB_DocsIf3CmMdCfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmMdCfgTable. The ifIndex key corresponds to
    // the MAC Domain interface. The type is slice of
    // DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry.
    DocsIf3CmMdCfgEntry []*DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry
}

func (docsIf3CmMdCfgTable *DOCSIF3MIB_DocsIf3CmMdCfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmMdCfgTable.EntityData.YFilter = docsIf3CmMdCfgTable.YFilter
    docsIf3CmMdCfgTable.EntityData.YangName = "docsIf3CmMdCfgTable"
    docsIf3CmMdCfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmMdCfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmMdCfgTable.EntityData.SegmentPath = "docsIf3CmMdCfgTable"
    docsIf3CmMdCfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmMdCfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmMdCfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmMdCfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmMdCfgTable.EntityData.Children.Append("docsIf3CmMdCfgEntry", types.YChild{"DocsIf3CmMdCfgEntry", nil})
    for i := range docsIf3CmMdCfgTable.DocsIf3CmMdCfgEntry {
        docsIf3CmMdCfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmMdCfgTable.DocsIf3CmMdCfgEntry[i]), types.YChild{"DocsIf3CmMdCfgEntry", docsIf3CmMdCfgTable.DocsIf3CmMdCfgEntry[i]})
    }
    docsIf3CmMdCfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmMdCfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmMdCfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry
// The conceptual row of docsIf3CmMdCfgTable.
// The ifIndex key corresponds to the MAC Domain interface.
type DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute specifies whether the CM honors or ignores the  CMTS MDD TLV
    // 5.1 setting in order to configure its IP provisioning  mode. The CM relies
    // upon the CMTS to facilitate the successful  IP address acquisition
    // independently of the MDD. When this attribute is set to 'ipv4Only' the CM
    // will initiate the acquisition of a single  IPv4 address for the CM
    // management stack. When this attribute is set to 'ipv6Only' the CM will
    // initiate the  acquisition of a single IPv6 address for the CM management
    // stack. When this attribute is set to 'honorMdd' the CM will initiate the 
    // acquisition of an IPv6 or IPv4 address as directed by the MDD  message for
    // provisioning and operation. The type is DocsIf3CmMdCfgIpProvMode.
    DocsIf3CmMdCfgIpProvMode interface{}

    // This attribute determines if the CM is to automatically reset upon change
    // of the IpProvMode attribute. The attribute has a default value of 'false'
    // (2) which means that the CM does not reset upon change to IpProvMode
    // attribute.  When this attribute is set to 'true' (1), the CM resets upon a
    // change to the IpProvMode attribute. The type is bool.
    DocsIf3CmMdCfgIpProvModeResetOnChange interface{}

    // This attribute determines how long a CM with  IpProvModeResetOnChange set
    // to 'true' waits to reset. When the  IpProvModeResetOnChange attribute is
    // set to 'true' (1), the CM will decrement from the configured timer value
    // before resetting. The default value of the
    // IpProvModeResetOnChangeHoldOffTimer is 0  seconds which is equivalent to an
    // immediate reset. The type is interface{} with range: 0..300. Units are
    // seconds.
    DocsIf3CmMdCfgIpProvModeResetOnChangeHoldOffTimer interface{}

    // This attribute determines if the CM persists the value of  IpProvMode
    // across a single reset or across all resets.   The  default value of this
    // attribute is 'nonVolatile' (3) which means that the CM persists the value
    // of IpProvMode across all resets.  The CM persists the value of IpProveMode
    // across only a single  reset when IpProvModeStorageType is set to
    // volatile(2). Other  StorageType values are not applicable for this object;
    // an attempt to set this object to a value of other(1), permanent(4), or
    // readOnly(5) will be rejected with an error code of inconsistentValue. The
    // type is StorageType.
    DocsIf3CmMdCfgIpProvModeStorageType interface{}
}

func (docsIf3CmMdCfgEntry *DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmMdCfgEntry.EntityData.YFilter = docsIf3CmMdCfgEntry.YFilter
    docsIf3CmMdCfgEntry.EntityData.YangName = "docsIf3CmMdCfgEntry"
    docsIf3CmMdCfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmMdCfgEntry.EntityData.ParentYangName = "docsIf3CmMdCfgTable"
    docsIf3CmMdCfgEntry.EntityData.SegmentPath = "docsIf3CmMdCfgEntry" + types.AddKeyToken(docsIf3CmMdCfgEntry.IfIndex, "ifIndex")
    docsIf3CmMdCfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmMdCfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmMdCfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmMdCfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmMdCfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmMdCfgEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmMdCfgEntry.IfIndex})
    docsIf3CmMdCfgEntry.EntityData.Leafs.Append("docsIf3CmMdCfgIpProvMode", types.YLeaf{"DocsIf3CmMdCfgIpProvMode", docsIf3CmMdCfgEntry.DocsIf3CmMdCfgIpProvMode})
    docsIf3CmMdCfgEntry.EntityData.Leafs.Append("docsIf3CmMdCfgIpProvModeResetOnChange", types.YLeaf{"DocsIf3CmMdCfgIpProvModeResetOnChange", docsIf3CmMdCfgEntry.DocsIf3CmMdCfgIpProvModeResetOnChange})
    docsIf3CmMdCfgEntry.EntityData.Leafs.Append("docsIf3CmMdCfgIpProvModeResetOnChangeHoldOffTimer", types.YLeaf{"DocsIf3CmMdCfgIpProvModeResetOnChangeHoldOffTimer", docsIf3CmMdCfgEntry.DocsIf3CmMdCfgIpProvModeResetOnChangeHoldOffTimer})
    docsIf3CmMdCfgEntry.EntityData.Leafs.Append("docsIf3CmMdCfgIpProvModeStorageType", types.YLeaf{"DocsIf3CmMdCfgIpProvModeStorageType", docsIf3CmMdCfgEntry.DocsIf3CmMdCfgIpProvModeStorageType})

    docsIf3CmMdCfgEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmMdCfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode represents message for provisioning and operation.
type DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode string

const (
    DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode_ipv4Only DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode = "ipv4Only"

    DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode_ipv6Only DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode = "ipv6Only"

    DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode_honorMdd DOCSIF3MIB_DocsIf3CmMdCfgTable_DocsIf3CmMdCfgEntry_DocsIf3CmMdCfgIpProvMode = "honorMdd"
)

// DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable
// This object provides configuration state information 
// on the CM for the Energy Management 1x1 Mode feature.
type DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmEnergyMgt1x1CfgTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable_DocsIf3CmEnergyMgt1x1CfgEntry.
    DocsIf3CmEnergyMgt1x1CfgEntry []*DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable_DocsIf3CmEnergyMgt1x1CfgEntry
}

func (docsIf3CmEnergyMgt1x1CfgTable *DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.YFilter = docsIf3CmEnergyMgt1x1CfgTable.YFilter
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.YangName = "docsIf3CmEnergyMgt1x1CfgTable"
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.SegmentPath = "docsIf3CmEnergyMgt1x1CfgTable"
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEnergyMgt1x1CfgTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.Children.Append("docsIf3CmEnergyMgt1x1CfgEntry", types.YChild{"DocsIf3CmEnergyMgt1x1CfgEntry", nil})
    for i := range docsIf3CmEnergyMgt1x1CfgTable.DocsIf3CmEnergyMgt1x1CfgEntry {
        docsIf3CmEnergyMgt1x1CfgTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmEnergyMgt1x1CfgTable.DocsIf3CmEnergyMgt1x1CfgEntry[i]), types.YChild{"DocsIf3CmEnergyMgt1x1CfgEntry", docsIf3CmEnergyMgt1x1CfgTable.DocsIf3CmEnergyMgt1x1CfgEntry[i]})
    }
    docsIf3CmEnergyMgt1x1CfgTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmEnergyMgt1x1CfgTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmEnergyMgt1x1CfgTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable_DocsIf3CmEnergyMgt1x1CfgEntry
// The conceptual row of docsIf3CmEnergyMgt1x1CfgTable.
type DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable_DocsIf3CmEnergyMgt1x1CfgEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This index indicates whether the threshold applies
    // to the  upstream or downstream. The type is IfDirection.
    DocsIf3CmEnergyMgt1x1CfgDirection interface{}

    // This attribute specifies the upstream or downstream bitrate threshold (in
    // bps) below which the CM will request to enter  Energy Management 1x1 Mode
    // operation. The type is interface{} with range: 0..4294967295. Units are
    // bps.
    DocsIf3CmEnergyMgt1x1CfgEntryBitrateThrshld interface{}

    // This attribute specifies the number of consecutive seconds  that the
    // upstream or downstream data rate needs to remain below the Upstream or
    // Downstream Entry Bitrate Threshold in order to  determine that a transition
    // to Energy Management 1x1 Mode is  required. The type is interface{} with
    // range: 1..65535. Units are seconds.
    DocsIf3CmEnergyMgt1x1CfgEntryTimeThrshld interface{}

    // This attribute specifies the upstream or downstream bitrate  threshold (in
    // bps) above which the CM will request to leave  Energy Management 1x1 Mode
    // operation. The type is interface{} with range: 0..4294967295. Units are
    // bps.
    DocsIf3CmEnergyMgt1x1CfgExitBitrateThrshld interface{}

    // This attribute specifies the number of consecutive seconds  that the
    // upstream or downstream data rate needs to remain above the Upstream or
    // Downstream Exit Bitrate Threshold in order to  determine that a transition
    // out of Energy Management 1x1 Mode  is required. The type is interface{}
    // with range: 1..65535. Units are seconds.
    DocsIf3CmEnergyMgt1x1CfgExitTimeThrshld interface{}
}

func (docsIf3CmEnergyMgt1x1CfgEntry *DOCSIF3MIB_DocsIf3CmEnergyMgt1x1CfgTable_DocsIf3CmEnergyMgt1x1CfgEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.YFilter = docsIf3CmEnergyMgt1x1CfgEntry.YFilter
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.YangName = "docsIf3CmEnergyMgt1x1CfgEntry"
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.ParentYangName = "docsIf3CmEnergyMgt1x1CfgTable"
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.SegmentPath = "docsIf3CmEnergyMgt1x1CfgEntry" + types.AddKeyToken(docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgDirection, "docsIf3CmEnergyMgt1x1CfgDirection")
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs.Append("docsIf3CmEnergyMgt1x1CfgDirection", types.YLeaf{"DocsIf3CmEnergyMgt1x1CfgDirection", docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgDirection})
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs.Append("docsIf3CmEnergyMgt1x1CfgEntryBitrateThrshld", types.YLeaf{"DocsIf3CmEnergyMgt1x1CfgEntryBitrateThrshld", docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgEntryBitrateThrshld})
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs.Append("docsIf3CmEnergyMgt1x1CfgEntryTimeThrshld", types.YLeaf{"DocsIf3CmEnergyMgt1x1CfgEntryTimeThrshld", docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgEntryTimeThrshld})
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs.Append("docsIf3CmEnergyMgt1x1CfgExitBitrateThrshld", types.YLeaf{"DocsIf3CmEnergyMgt1x1CfgExitBitrateThrshld", docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgExitBitrateThrshld})
    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.Leafs.Append("docsIf3CmEnergyMgt1x1CfgExitTimeThrshld", types.YLeaf{"DocsIf3CmEnergyMgt1x1CfgExitTimeThrshld", docsIf3CmEnergyMgt1x1CfgEntry.DocsIf3CmEnergyMgt1x1CfgExitTimeThrshld})

    docsIf3CmEnergyMgt1x1CfgEntry.EntityData.YListKeys = []string {"DocsIf3CmEnergyMgt1x1CfgDirection"}

    return &(docsIf3CmEnergyMgt1x1CfgEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable
// This table provides a list of spectral analysis measurements 
// as performed across a range of center frequencies. The table 
// is capable of representing a full scan of the spectrum.
type DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each row in the docsIf3CmSpectrumAnalysisMeasTable  represents the spectral
    // analysis around a single center  frequency point in the spectrum. The type
    // is slice of
    // DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable_DocsIf3CmSpectrumAnalysisMeasEntry.
    DocsIf3CmSpectrumAnalysisMeasEntry []*DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable_DocsIf3CmSpectrumAnalysisMeasEntry
}

func (docsIf3CmSpectrumAnalysisMeasTable *DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.YFilter = docsIf3CmSpectrumAnalysisMeasTable.YFilter
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.YangName = "docsIf3CmSpectrumAnalysisMeasTable"
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.SegmentPath = "docsIf3CmSpectrumAnalysisMeasTable"
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmSpectrumAnalysisMeasTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.Children.Append("docsIf3CmSpectrumAnalysisMeasEntry", types.YChild{"DocsIf3CmSpectrumAnalysisMeasEntry", nil})
    for i := range docsIf3CmSpectrumAnalysisMeasTable.DocsIf3CmSpectrumAnalysisMeasEntry {
        docsIf3CmSpectrumAnalysisMeasTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmSpectrumAnalysisMeasTable.DocsIf3CmSpectrumAnalysisMeasEntry[i]), types.YChild{"DocsIf3CmSpectrumAnalysisMeasEntry", docsIf3CmSpectrumAnalysisMeasTable.DocsIf3CmSpectrumAnalysisMeasEntry[i]})
    }
    docsIf3CmSpectrumAnalysisMeasTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmSpectrumAnalysisMeasTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmSpectrumAnalysisMeasTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable_DocsIf3CmSpectrumAnalysisMeasEntry
// Each row in the docsIf3CmSpectrumAnalysisMeasTable 
// represents the spectral analysis around a single center 
// frequency point in the spectrum.
type DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable_DocsIf3CmSpectrumAnalysisMeasEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This index indicates the center frequency of the
    // spectral analysis span which is represented by this instance. The type is
    // interface{} with range: -2147483648..2147483647.
    DocsIf3CmSpectrumAnalysisMeasFrequency interface{}

    // This attribute provides a list of the spectral amplitudes as  measured at
    // the center frequency specified by the Frequency index.  The frequency bins
    // are ordered from lowest to highest  frequencies covering the frequency
    // span. Information about the center frequency, frequency span, number of
    // bins and resolution bandwidth are included to provide context to the
    // measurement  point. The type is string with length: 0 | 2..4116.
    DocsIf3CmSpectrumAnalysisMeasAmplitudeData interface{}

    // This attribute provides the total RF power present in the  segment with the
    // center frequency equal to the Frequency index and the span equal to the
    // SegmentFrequencySpan. The value  represents the sum of the spectrum power
    // in all of the  associated bins. The value is computed by summing power (not
    // dB) values and converting the final sum to TenthdB. The type is interface{}
    // with range: -2147483648..2147483647.
    DocsIf3CmSpectrumAnalysisMeasTotalSegmentPower interface{}
}

func (docsIf3CmSpectrumAnalysisMeasEntry *DOCSIF3MIB_DocsIf3CmSpectrumAnalysisMeasTable_DocsIf3CmSpectrumAnalysisMeasEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.YFilter = docsIf3CmSpectrumAnalysisMeasEntry.YFilter
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.YangName = "docsIf3CmSpectrumAnalysisMeasEntry"
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.ParentYangName = "docsIf3CmSpectrumAnalysisMeasTable"
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.SegmentPath = "docsIf3CmSpectrumAnalysisMeasEntry" + types.AddKeyToken(docsIf3CmSpectrumAnalysisMeasEntry.DocsIf3CmSpectrumAnalysisMeasFrequency, "docsIf3CmSpectrumAnalysisMeasFrequency")
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisMeasFrequency", types.YLeaf{"DocsIf3CmSpectrumAnalysisMeasFrequency", docsIf3CmSpectrumAnalysisMeasEntry.DocsIf3CmSpectrumAnalysisMeasFrequency})
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisMeasAmplitudeData", types.YLeaf{"DocsIf3CmSpectrumAnalysisMeasAmplitudeData", docsIf3CmSpectrumAnalysisMeasEntry.DocsIf3CmSpectrumAnalysisMeasAmplitudeData})
    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.Leafs.Append("docsIf3CmSpectrumAnalysisMeasTotalSegmentPower", types.YLeaf{"DocsIf3CmSpectrumAnalysisMeasTotalSegmentPower", docsIf3CmSpectrumAnalysisMeasEntry.DocsIf3CmSpectrumAnalysisMeasTotalSegmentPower})

    docsIf3CmSpectrumAnalysisMeasEntry.EntityData.YListKeys = []string {"DocsIf3CmSpectrumAnalysisMeasFrequency"}

    return &(docsIf3CmSpectrumAnalysisMeasEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable
// This table defines Energy Management mode statistics for the
// CM as reported by the CMTS.  For example, such metrics can
// provide insight into configuration of appropriate EM 1x1 Mode
// Activity Detection thresholds on the CM and/or to get feedback
// on how/if the current thresholds are working well or are 
// causing user experience issues.
type DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmtsCmEmStatsTable. The type is slice of
    // DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable_DocsIf3CmtsCmEmStatsEntry.
    DocsIf3CmtsCmEmStatsEntry []*DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable_DocsIf3CmtsCmEmStatsEntry
}

func (docsIf3CmtsCmEmStatsTable *DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmEmStatsTable.EntityData.YFilter = docsIf3CmtsCmEmStatsTable.YFilter
    docsIf3CmtsCmEmStatsTable.EntityData.YangName = "docsIf3CmtsCmEmStatsTable"
    docsIf3CmtsCmEmStatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmEmStatsTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmtsCmEmStatsTable.EntityData.SegmentPath = "docsIf3CmtsCmEmStatsTable"
    docsIf3CmtsCmEmStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmEmStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmEmStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmEmStatsTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmEmStatsTable.EntityData.Children.Append("docsIf3CmtsCmEmStatsEntry", types.YChild{"DocsIf3CmtsCmEmStatsEntry", nil})
    for i := range docsIf3CmtsCmEmStatsTable.DocsIf3CmtsCmEmStatsEntry {
        docsIf3CmtsCmEmStatsTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmtsCmEmStatsTable.DocsIf3CmtsCmEmStatsEntry[i]), types.YChild{"DocsIf3CmtsCmEmStatsEntry", docsIf3CmtsCmEmStatsTable.DocsIf3CmtsCmEmStatsEntry[i]})
    }
    docsIf3CmtsCmEmStatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmtsCmEmStatsTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmtsCmEmStatsTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable_DocsIf3CmtsCmEmStatsEntry
// The conceptual row of docsIf3CmtsCmEmStatsTable.
type DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable_DocsIf3CmtsCmEmStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // docs_if3_mib.DOCSIF3MIB_DocsIf3CmtsCmRegStatusTable_DocsIf3CmtsCmRegStatusEntry_DocsIf3CmtsCmRegStatusId
    DocsIf3CmtsCmRegStatusId interface{}

    // This attribute indicates the total time duration, in seconds since
    // registration, the CM identified by  docsIf3CmtsCmRegStatusId has been in
    // Energy Management 1x1  mode, as controlled by the DBC-REQ Energy Management
    // 1x1  Mode Indicator TLV. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    DocsIf3CmtsCmEmStatsEm1x1ModeTotalDuration interface{}
}

func (docsIf3CmtsCmEmStatsEntry *DOCSIF3MIB_DocsIf3CmtsCmEmStatsTable_DocsIf3CmtsCmEmStatsEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmtsCmEmStatsEntry.EntityData.YFilter = docsIf3CmtsCmEmStatsEntry.YFilter
    docsIf3CmtsCmEmStatsEntry.EntityData.YangName = "docsIf3CmtsCmEmStatsEntry"
    docsIf3CmtsCmEmStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmtsCmEmStatsEntry.EntityData.ParentYangName = "docsIf3CmtsCmEmStatsTable"
    docsIf3CmtsCmEmStatsEntry.EntityData.SegmentPath = "docsIf3CmtsCmEmStatsEntry" + types.AddKeyToken(docsIf3CmtsCmEmStatsEntry.DocsIf3CmtsCmRegStatusId, "docsIf3CmtsCmRegStatusId")
    docsIf3CmtsCmEmStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmtsCmEmStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmtsCmEmStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmtsCmEmStatsEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmtsCmEmStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmtsCmEmStatsEntry.EntityData.Leafs.Append("docsIf3CmtsCmRegStatusId", types.YLeaf{"DocsIf3CmtsCmRegStatusId", docsIf3CmtsCmEmStatsEntry.DocsIf3CmtsCmRegStatusId})
    docsIf3CmtsCmEmStatsEntry.EntityData.Leafs.Append("docsIf3CmtsCmEmStatsEm1x1ModeTotalDuration", types.YLeaf{"DocsIf3CmtsCmEmStatsEm1x1ModeTotalDuration", docsIf3CmtsCmEmStatsEntry.DocsIf3CmtsCmEmStatsEm1x1ModeTotalDuration})

    docsIf3CmtsCmEmStatsEntry.EntityData.YListKeys = []string {"DocsIf3CmtsCmRegStatusId"}

    return &(docsIf3CmtsCmEmStatsEntry.EntityData)
}

// DOCSIF3MIB_DocsIf3CmEm1x1StatsTable
// This object defines Energy Management 1x1 mode statistics on
// the CM to provide insight into configuration of appropriate EM
// 1x1 Mode Activity Detection thresholds and/or to get feedback
// on how/if the current thresholds are working well or are 
// causing user experience issues.
// These statistics are only applicable/valid when the Energy 
// Management 1x1 mode is enabled in the CM.
type DOCSIF3MIB_DocsIf3CmEm1x1StatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The conceptual row of docsIf3CmEm1x1StatsTable. An instance exist for the
    // CM MAC Domain Interface. The type is slice of
    // DOCSIF3MIB_DocsIf3CmEm1x1StatsTable_DocsIf3CmEm1x1StatsEntry.
    DocsIf3CmEm1x1StatsEntry []*DOCSIF3MIB_DocsIf3CmEm1x1StatsTable_DocsIf3CmEm1x1StatsEntry
}

func (docsIf3CmEm1x1StatsTable *DOCSIF3MIB_DocsIf3CmEm1x1StatsTable) GetEntityData() *types.CommonEntityData {
    docsIf3CmEm1x1StatsTable.EntityData.YFilter = docsIf3CmEm1x1StatsTable.YFilter
    docsIf3CmEm1x1StatsTable.EntityData.YangName = "docsIf3CmEm1x1StatsTable"
    docsIf3CmEm1x1StatsTable.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEm1x1StatsTable.EntityData.ParentYangName = "DOCS-IF3-MIB"
    docsIf3CmEm1x1StatsTable.EntityData.SegmentPath = "docsIf3CmEm1x1StatsTable"
    docsIf3CmEm1x1StatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEm1x1StatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEm1x1StatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEm1x1StatsTable.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEm1x1StatsTable.EntityData.Children.Append("docsIf3CmEm1x1StatsEntry", types.YChild{"DocsIf3CmEm1x1StatsEntry", nil})
    for i := range docsIf3CmEm1x1StatsTable.DocsIf3CmEm1x1StatsEntry {
        docsIf3CmEm1x1StatsTable.EntityData.Children.Append(types.GetSegmentPath(docsIf3CmEm1x1StatsTable.DocsIf3CmEm1x1StatsEntry[i]), types.YChild{"DocsIf3CmEm1x1StatsEntry", docsIf3CmEm1x1StatsTable.DocsIf3CmEm1x1StatsEntry[i]})
    }
    docsIf3CmEm1x1StatsTable.EntityData.Leafs = types.NewOrderedMap()

    docsIf3CmEm1x1StatsTable.EntityData.YListKeys = []string {}

    return &(docsIf3CmEm1x1StatsTable.EntityData)
}

// DOCSIF3MIB_DocsIf3CmEm1x1StatsTable_DocsIf3CmEm1x1StatsEntry
// The conceptual row of docsIf3CmEm1x1StatsTable.
// An instance exist for the CM MAC Domain Interface.
type DOCSIF3MIB_DocsIf3CmEm1x1StatsTable_DocsIf3CmEm1x1StatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute indicates the number of times since  registration the CM
    // crossed below the upstream entry bitrate threshold for a number of
    // consecutive seconds equal to or  exceeding the upstream entry time
    // threshold. The type is interface{} with range: 0..4294967295.
    DocsIf3CmEm1x1StatsNumberTimesCrossedBelowUsEntryThrshlds interface{}

    // This attribute indicates the number of times since  registration the CM
    // crossed below the downstream entry  bitrate threshold for a number of
    // consecutive seconds equal to or exceeding the downstream entry time
    // threshold. The type is interface{} with range: 0..4294967295.
    DocsIf3CmEm1x1StatsNumberTimesCrossedBelowDsEntryThrshlds interface{}

    // This attribute indicates the total time duration, in seconds since
    // registration, the CM has been in Energy Management 1x1 mode, as controlled
    // by the DBC-REQ Energy Management 1x1 Mode Indicator TLV. This attribute
    // differs from  docsIf3CmEm1x1StatsTotalDurationBelowUsDsThrshlds because it
    // is dependent on effects of the Energy Management Cycle  Period, and
    // processing of EM-REQ/EM-RSP messages and DBC  messages that specifically
    // indicate entry into or exit from Energy Management 1x1 mode. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    DocsIf3CmEm1x1StatsTotalDuration interface{}

    // This attribute indicates the total time duration, in seconds since
    // registration, the CM satisfied upstream conditions for  entry into or
    // remaining in Energy Management 1x1 mode. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    DocsIf3CmEm1x1StatsTotalDurationBelowUsThrshlds interface{}

    // This attribute indicates the total time duration, in seconds since
    // registration, the CM satisfied downstream conditions for  entry into or
    // remaining in Energy Management 1x1 mode. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    DocsIf3CmEm1x1StatsTotalDurationBelowDsThrshlds interface{}

    // This attribute indicates the total time duration, in seconds  since
    // registration, the CM, with respect to both upstream and  downstream entry
    // and exit thresholds, satisfied conditions for  entry into and remaining in
    // Energy Management 1x1 mode.  This  attribute differs from
    // docsIf3CmEm1x1StatsTotalDuration  because it is not dependent on effects of
    // the Energy  Management Cycle Period or processing of EM-REQ/EM-RSP 
    // messages and DBC messages that specifically indicate entry into or exit
    // from Energy Management 1x1 mode. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    DocsIf3CmEm1x1StatsTotalDurationBelowUsDsThrshlds interface{}
}

func (docsIf3CmEm1x1StatsEntry *DOCSIF3MIB_DocsIf3CmEm1x1StatsTable_DocsIf3CmEm1x1StatsEntry) GetEntityData() *types.CommonEntityData {
    docsIf3CmEm1x1StatsEntry.EntityData.YFilter = docsIf3CmEm1x1StatsEntry.YFilter
    docsIf3CmEm1x1StatsEntry.EntityData.YangName = "docsIf3CmEm1x1StatsEntry"
    docsIf3CmEm1x1StatsEntry.EntityData.BundleName = "cisco_ios_xe"
    docsIf3CmEm1x1StatsEntry.EntityData.ParentYangName = "docsIf3CmEm1x1StatsTable"
    docsIf3CmEm1x1StatsEntry.EntityData.SegmentPath = "docsIf3CmEm1x1StatsEntry" + types.AddKeyToken(docsIf3CmEm1x1StatsEntry.IfIndex, "ifIndex")
    docsIf3CmEm1x1StatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    docsIf3CmEm1x1StatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    docsIf3CmEm1x1StatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    docsIf3CmEm1x1StatsEntry.EntityData.Children = types.NewOrderedMap()
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs = types.NewOrderedMap()
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", docsIf3CmEm1x1StatsEntry.IfIndex})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsNumberTimesCrossedBelowUsEntryThrshlds", types.YLeaf{"DocsIf3CmEm1x1StatsNumberTimesCrossedBelowUsEntryThrshlds", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsNumberTimesCrossedBelowUsEntryThrshlds})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsNumberTimesCrossedBelowDsEntryThrshlds", types.YLeaf{"DocsIf3CmEm1x1StatsNumberTimesCrossedBelowDsEntryThrshlds", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsNumberTimesCrossedBelowDsEntryThrshlds})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsTotalDuration", types.YLeaf{"DocsIf3CmEm1x1StatsTotalDuration", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsTotalDuration})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsTotalDurationBelowUsThrshlds", types.YLeaf{"DocsIf3CmEm1x1StatsTotalDurationBelowUsThrshlds", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsTotalDurationBelowUsThrshlds})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsTotalDurationBelowDsThrshlds", types.YLeaf{"DocsIf3CmEm1x1StatsTotalDurationBelowDsThrshlds", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsTotalDurationBelowDsThrshlds})
    docsIf3CmEm1x1StatsEntry.EntityData.Leafs.Append("docsIf3CmEm1x1StatsTotalDurationBelowUsDsThrshlds", types.YLeaf{"DocsIf3CmEm1x1StatsTotalDurationBelowUsDsThrshlds", docsIf3CmEm1x1StatsEntry.DocsIf3CmEm1x1StatsTotalDurationBelowUsDsThrshlds})

    docsIf3CmEm1x1StatsEntry.EntityData.YListKeys = []string {"IfIndex"}

    return &(docsIf3CmEm1x1StatsEntry.EntityData)
}

