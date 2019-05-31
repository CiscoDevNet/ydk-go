// This module contains definitions
// for the Calvados model objects.
// 
// This module contains the YANG enumerated type
// definitions used by the Cisco IOS-XR SysAdmin
// Control Ethernet commands.
// 
// Copyright(c) 2011-2017 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_ethsw_esdma_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_ethsw_esdma_types"))
}

// EsdmaRackTypeEnum represents The valid chassis types supported supported in Sysadmin.
type EsdmaRackTypeEnum string

const (
    EsdmaRackTypeEnum_Unknown EsdmaRackTypeEnum = "Unknown"

    EsdmaRackTypeEnum_FCC EsdmaRackTypeEnum = "FCC"

    EsdmaRackTypeEnum_LCC EsdmaRackTypeEnum = "LCC"

    EsdmaRackTypeEnum_BSC EsdmaRackTypeEnum = "BSC"

    EsdmaRackTypeEnum_COMPUTE EsdmaRackTypeEnum = "COMPUTE"
)

// EsdmaRackNumEnum represents The valid rack numbers supported in Sysadmin.
type EsdmaRackNumEnum string

const (
    EsdmaRackNumEnum_Y_0 EsdmaRackNumEnum = "0"

    EsdmaRackNumEnum_Y_1 EsdmaRackNumEnum = "1"

    EsdmaRackNumEnum_Y_2 EsdmaRackNumEnum = "2"

    EsdmaRackNumEnum_Y_3 EsdmaRackNumEnum = "3"

    EsdmaRackNumEnum_Y_4 EsdmaRackNumEnum = "4"

    EsdmaRackNumEnum_Y_5 EsdmaRackNumEnum = "5"

    EsdmaRackNumEnum_Y_6 EsdmaRackNumEnum = "6"

    EsdmaRackNumEnum_Y_7 EsdmaRackNumEnum = "7"

    EsdmaRackNumEnum_Y_8 EsdmaRackNumEnum = "8"

    EsdmaRackNumEnum_Y_9 EsdmaRackNumEnum = "9"

    EsdmaRackNumEnum_Y_10 EsdmaRackNumEnum = "10"

    EsdmaRackNumEnum_Y_11 EsdmaRackNumEnum = "11"

    EsdmaRackNumEnum_Y_12 EsdmaRackNumEnum = "12"

    EsdmaRackNumEnum_Y_13 EsdmaRackNumEnum = "13"

    EsdmaRackNumEnum_Y_14 EsdmaRackNumEnum = "14"

    EsdmaRackNumEnum_Y_15 EsdmaRackNumEnum = "15"

    EsdmaRackNumEnum_F0 EsdmaRackNumEnum = "F0"

    EsdmaRackNumEnum_F1 EsdmaRackNumEnum = "F1"

    EsdmaRackNumEnum_F2 EsdmaRackNumEnum = "F2"

    EsdmaRackNumEnum_F3 EsdmaRackNumEnum = "F3"

    EsdmaRackNumEnum_B0 EsdmaRackNumEnum = "B0"

    EsdmaRackNumEnum_B1 EsdmaRackNumEnum = "B1"
)

// EsdmaCpu represents The set of CPU enumerations that have control plane Ethernet switches or run the MLAP protocol.
type EsdmaCpu string

const (
    EsdmaCpu_Unknown EsdmaCpu = "Unknown"

    EsdmaCpu_RP0 EsdmaCpu = "RP0"

    EsdmaCpu_RP1 EsdmaCpu = "RP1"

    EsdmaCpu_SC0 EsdmaCpu = "SC0"

    EsdmaCpu_SC1 EsdmaCpu = "SC1"

    EsdmaCpu_LC0 EsdmaCpu = "LC0"

    EsdmaCpu_LC1 EsdmaCpu = "LC1"

    EsdmaCpu_LC2 EsdmaCpu = "LC2"

    EsdmaCpu_LC3 EsdmaCpu = "LC3"

    EsdmaCpu_LC4 EsdmaCpu = "LC4"

    EsdmaCpu_LC5 EsdmaCpu = "LC5"

    EsdmaCpu_LC6 EsdmaCpu = "LC6"

    EsdmaCpu_LC7 EsdmaCpu = "LC7"

    EsdmaCpu_LC8 EsdmaCpu = "LC8"

    EsdmaCpu_LC9 EsdmaCpu = "LC9"

    EsdmaCpu_LC10 EsdmaCpu = "LC10"

    EsdmaCpu_LC11 EsdmaCpu = "LC11"

    EsdmaCpu_LC12 EsdmaCpu = "LC12"

    EsdmaCpu_LC13 EsdmaCpu = "LC13"

    EsdmaCpu_LC14 EsdmaCpu = "LC14"

    EsdmaCpu_LC15 EsdmaCpu = "LC15"

    EsdmaCpu_LC16 EsdmaCpu = "LC16"

    EsdmaCpu_LC17 EsdmaCpu = "LC17"

    EsdmaCpu_LC18 EsdmaCpu = "LC18"

    EsdmaCpu_LC19 EsdmaCpu = "LC19"

    EsdmaCpu_FC0 EsdmaCpu = "FC0"

    EsdmaCpu_FC1 EsdmaCpu = "FC1"

    EsdmaCpu_FC2 EsdmaCpu = "FC2"

    EsdmaCpu_FC3 EsdmaCpu = "FC3"

    EsdmaCpu_FC4 EsdmaCpu = "FC4"

    EsdmaCpu_FC5 EsdmaCpu = "FC5"

    EsdmaCpu_CB0 EsdmaCpu = "CB0"
)

// EsdmaSwitchTypeEnum represents The list of Ethernet switch types
type EsdmaSwitchTypeEnum string

const (
    EsdmaSwitchTypeEnum_RP_SW EsdmaSwitchTypeEnum = "RP-SW"

    EsdmaSwitchTypeEnum_SC_SW EsdmaSwitchTypeEnum = "SC-SW"

    EsdmaSwitchTypeEnum_LC_SW EsdmaSwitchTypeEnum = "LC-SW"

    EsdmaSwitchTypeEnum_F_SW0 EsdmaSwitchTypeEnum = "F-SW0"

    EsdmaSwitchTypeEnum_F_SW1 EsdmaSwitchTypeEnum = "F-SW1"

    EsdmaSwitchTypeEnum_FC_SW EsdmaSwitchTypeEnum = "FC-SW"

    EsdmaSwitchTypeEnum_EOBC_SW EsdmaSwitchTypeEnum = "EOBC-SW"

    EsdmaSwitchTypeEnum_EPC_SW EsdmaSwitchTypeEnum = "EPC-SW"

    EsdmaSwitchTypeEnum_CB_SW EsdmaSwitchTypeEnum = "CB-SW"

    EsdmaSwitchTypeEnum_Unknown EsdmaSwitchTypeEnum = "Unknown"

    EsdmaSwitchTypeEnum_RP_SW1 EsdmaSwitchTypeEnum = "RP-SW1"
)

// EsdmaSwitchYesNoEnum
type EsdmaSwitchYesNoEnum string

const (
    EsdmaSwitchYesNoEnum_Yes EsdmaSwitchYesNoEnum = "Yes"

    EsdmaSwitchYesNoEnum_No EsdmaSwitchYesNoEnum = "No"
)

// EsdmaSwitchSfpInsertedEnum
type EsdmaSwitchSfpInsertedEnum string

const (
    EsdmaSwitchSfpInsertedEnum_Unknown EsdmaSwitchSfpInsertedEnum = "Unknown"

    EsdmaSwitchSfpInsertedEnum_Yes EsdmaSwitchSfpInsertedEnum = "Yes"

    EsdmaSwitchSfpInsertedEnum_No EsdmaSwitchSfpInsertedEnum = "No"

    EsdmaSwitchSfpInsertedEnum_Failed EsdmaSwitchSfpInsertedEnum = "Failed"
)

// EsdmaSwitchSfpControllerEnum
type EsdmaSwitchSfpControllerEnum string

const (
    EsdmaSwitchSfpControllerEnum_Unknown EsdmaSwitchSfpControllerEnum = "Unknown"

    EsdmaSwitchSfpControllerEnum_Switch EsdmaSwitchSfpControllerEnum = "Switch"

    EsdmaSwitchSfpControllerEnum_PHY EsdmaSwitchSfpControllerEnum = "PHY"
)

// EsdmaSwitchSfpTranceiverTypeEnum
type EsdmaSwitchSfpTranceiverTypeEnum string

const (
    EsdmaSwitchSfpTranceiverTypeEnum_Unspecified EsdmaSwitchSfpTranceiverTypeEnum = "Unspecified"

    EsdmaSwitchSfpTranceiverTypeEnum_SFP EsdmaSwitchSfpTranceiverTypeEnum = "SFP"

    EsdmaSwitchSfpTranceiverTypeEnum_QSFP EsdmaSwitchSfpTranceiverTypeEnum = "QSFP"

    EsdmaSwitchSfpTranceiverTypeEnum_Unknown EsdmaSwitchSfpTranceiverTypeEnum = "Unknown"
)

// EsdmaSfpEncodingEnum
type EsdmaSfpEncodingEnum string

const (
    EsdmaSfpEncodingEnum_Unspecified EsdmaSfpEncodingEnum = "Unspecified"

    EsdmaSfpEncodingEnum_Y_8B__FWD_SLASH__10B EsdmaSfpEncodingEnum = "8B/10B"

    EsdmaSfpEncodingEnum_Y_4B__FWD_SLASH__5B EsdmaSfpEncodingEnum = "4B/5B"

    EsdmaSfpEncodingEnum_NRZ EsdmaSfpEncodingEnum = "NRZ"

    EsdmaSfpEncodingEnum_Manchester EsdmaSfpEncodingEnum = "Manchester"

    EsdmaSfpEncodingEnum_SONET_Scrambled EsdmaSfpEncodingEnum = "SONET Scrambled"

    EsdmaSfpEncodingEnum_Y_64B__FWD_SLASH__66B EsdmaSfpEncodingEnum = "64B/66B"

    EsdmaSfpEncodingEnum_Unknown EsdmaSfpEncodingEnum = "Unknown"
)

// EsdmaQsfpTransceiverEnum
type EsdmaQsfpTransceiverEnum string

const (
    EsdmaQsfpTransceiverEnum_QSFP_40G_LR4 EsdmaQsfpTransceiverEnum = "QSFP-40G-LR4"

    EsdmaQsfpTransceiverEnum_QSFP_40G_SR4 EsdmaQsfpTransceiverEnum = "QSFP-40G-SR4"

    EsdmaQsfpTransceiverEnum_QSFP_40G_CR4_Active EsdmaQsfpTransceiverEnum = "QSFP-40G-CR4-Active"

    EsdmaQsfpTransceiverEnum_QSFP_40G_CR4_Passive EsdmaQsfpTransceiverEnum = "QSFP-40G-CR4-Passive"

    EsdmaQsfpTransceiverEnum_Unknown EsdmaQsfpTransceiverEnum = "Unknown"
)

// EsdmaSwitchSfpTypeEnum
type EsdmaSwitchSfpTypeEnum string

const (
    EsdmaSwitchSfpTypeEnum_None EsdmaSwitchSfpTypeEnum = "None"

    EsdmaSwitchSfpTypeEnum_SFP_10G_SR EsdmaSwitchSfpTypeEnum = "SFP-10G-SR"

    EsdmaSwitchSfpTypeEnum_SFP_10G_LR EsdmaSwitchSfpTypeEnum = "SFP-10G-LR"

    EsdmaSwitchSfpTypeEnum_SFP_10G_LRM EsdmaSwitchSfpTypeEnum = "SFP-10G-LRM"

    EsdmaSwitchSfpTypeEnum_SFP_1G_SX EsdmaSwitchSfpTypeEnum = "SFP-1G-SX"

    EsdmaSwitchSfpTypeEnum_SFP_1G_LX EsdmaSwitchSfpTypeEnum = "SFP-1G-LX"

    EsdmaSwitchSfpTypeEnum_SFP_1000Base_T EsdmaSwitchSfpTypeEnum = "SFP-1000Base-T"

    EsdmaSwitchSfpTypeEnum_SFP_40G_SR4 EsdmaSwitchSfpTypeEnum = "SFP-40G-SR4"

    EsdmaSwitchSfpTypeEnum_SFP_40G_LR4 EsdmaSwitchSfpTypeEnum = "SFP-40G-LR4"

    EsdmaSwitchSfpTypeEnum_Unsupported EsdmaSwitchSfpTypeEnum = "Unsupported"
)

// EsdmaSwitchPortState represents The switch port up and down states
type EsdmaSwitchPortState string

const (
    EsdmaSwitchPortState_Unknown EsdmaSwitchPortState = "Unknown"

    EsdmaSwitchPortState_Test EsdmaSwitchPortState = "Test"

    EsdmaSwitchPortState_Down EsdmaSwitchPortState = "Down"

    EsdmaSwitchPortState_Up EsdmaSwitchPortState = "Up"
)

// SwitchForwardingState represents The set of switch port forwarding states
type SwitchForwardingState string

const (
    SwitchForwardingState_Unknown SwitchForwardingState = "Unknown"

    SwitchForwardingState_Y_ SwitchForwardingState = "-"

    SwitchForwardingState_Disabled SwitchForwardingState = "Disabled"

    SwitchForwardingState_Blocking SwitchForwardingState = "Blocking"

    SwitchForwardingState_Learning SwitchForwardingState = "Learning"

    SwitchForwardingState_Forwarding SwitchForwardingState = "Forwarding"
)

// EsdCirEirType
type EsdCirEirType string

const (
    EsdCirEirType_CIR EsdCirEirType = "CIR"

    EsdCirEirType_PIR EsdCirEirType = "PIR"
)

// MlapEpType
type MlapEpType string

const (
    MlapEpType_Unknown MlapEpType = "Unknown"

    MlapEpType_RP MlapEpType = "RP"

    MlapEpType_SC MlapEpType = "SC"

    MlapEpType_FC MlapEpType = "FC"

    MlapEpType_LC MlapEpType = "LC"

    MlapEpType_F_SW MlapEpType = "F-SW"

    MlapEpType_CB MlapEpType = "CB"
)

// MlapStateEnum represents The set of enumerated values that MLAP uses to manage a port's protocol state
type MlapStateEnum string

const (
    MlapStateEnum_Y_ MlapStateEnum = "-"

    MlapStateEnum_Down MlapStateEnum = "Down"

    MlapStateEnum_Up MlapStateEnum = "Up"

    MlapStateEnum_Admin_Down MlapStateEnum = "Admin Down"

    MlapStateEnum_Do_Not_Use MlapStateEnum = "Do Not Use"

    MlapStateEnum_Invalid MlapStateEnum = "Invalid"

    MlapStateEnum_Active MlapStateEnum = "Active"

    MlapStateEnum_Standby MlapStateEnum = "Standby"

    MlapStateEnum_Rem_Managed MlapStateEnum = "Rem Managed"
)

// MlapProtocolEnum represents The types of MLAP protocol
type MlapProtocolEnum string

const (
    MlapProtocolEnum_Internal MlapProtocolEnum = "Internal"

    MlapProtocolEnum_External MlapProtocolEnum = "External"
)

// MlapTraceVerbosity
type MlapTraceVerbosity string

const (
    MlapTraceVerbosity_Off MlapTraceVerbosity = "Off"

    MlapTraceVerbosity_Low MlapTraceVerbosity = "Low"

    MlapTraceVerbosity_Medium MlapTraceVerbosity = "Medium"

    MlapTraceVerbosity_High MlapTraceVerbosity = "High"
)

// SwitchDataDirectionEnum represents Switch data direction
type SwitchDataDirectionEnum string

const (
    SwitchDataDirectionEnum_Y_ SwitchDataDirectionEnum = "-"

    SwitchDataDirectionEnum_Both SwitchDataDirectionEnum = "Both"

    SwitchDataDirectionEnum_Rx SwitchDataDirectionEnum = "Rx"

    SwitchDataDirectionEnum_Tx SwitchDataDirectionEnum = "Tx"

    SwitchDataDirectionEnum_Unknown SwitchDataDirectionEnum = "Unknown"

    SwitchDataDirectionEnum_Invalid SwitchDataDirectionEnum = "Invalid"
)

// SwitchTableTypeEnum
type SwitchTableTypeEnum string

const (
    SwitchTableTypeEnum_Y_ SwitchTableTypeEnum = "-"

    SwitchTableTypeEnum_Port SwitchTableTypeEnum = "Port"

    SwitchTableTypeEnum_VLAN SwitchTableTypeEnum = "VLAN"

    SwitchTableTypeEnum_TCAM SwitchTableTypeEnum = "TCAM"

    SwitchTableTypeEnum_Unknown SwitchTableTypeEnum = "Unknown"
)

// SwitchMatchTypeEnum
type SwitchMatchTypeEnum string

const (
    SwitchMatchTypeEnum_Y_ SwitchMatchTypeEnum = "-"

    SwitchMatchTypeEnum_Any SwitchMatchTypeEnum = "Any"

    SwitchMatchTypeEnum_Tagged SwitchMatchTypeEnum = "Tagged"

    SwitchMatchTypeEnum_Untagged SwitchMatchTypeEnum = "Untagged"

    SwitchMatchTypeEnum_Unknown SwitchMatchTypeEnum = "Unknown"
)

// SwitchActionTypeEnum
type SwitchActionTypeEnum string

const (
    SwitchActionTypeEnum_Y_ SwitchActionTypeEnum = "-"

    SwitchActionTypeEnum_Translate SwitchActionTypeEnum = "Translate"

    SwitchActionTypeEnum_Remove_Outer SwitchActionTypeEnum = "Remove Outer"

    SwitchActionTypeEnum_Add_Outer SwitchActionTypeEnum = "Add Outer"

    SwitchActionTypeEnum_Drop SwitchActionTypeEnum = "Drop"

    SwitchActionTypeEnum_Forward SwitchActionTypeEnum = "Forward"

    SwitchActionTypeEnum_Unknown SwitchActionTypeEnum = "Unknown"
)

// EsdmaSdrTrafficType
type EsdmaSdrTrafficType string

const (
    EsdmaSdrTrafficType_Unknown EsdmaSdrTrafficType = "Unknown"

    EsdmaSdrTrafficType_IPC EsdmaSdrTrafficType = "IPC"

    EsdmaSdrTrafficType_MgmtEth EsdmaSdrTrafficType = "MgmtEth"

    EsdmaSdrTrafficType_All EsdmaSdrTrafficType = "All"

    EsdmaSdrTrafficType_Invalid EsdmaSdrTrafficType = "Invalid"
)

// EsdmaTrunkMemberStatus
type EsdmaTrunkMemberStatus string

const (
    EsdmaTrunkMemberStatus_Disabled EsdmaTrunkMemberStatus = "Disabled"

    EsdmaTrunkMemberStatus_Enabled EsdmaTrunkMemberStatus = "Enabled"

    EsdmaTrunkMemberStatus_Y_ EsdmaTrunkMemberStatus = "-"
)

